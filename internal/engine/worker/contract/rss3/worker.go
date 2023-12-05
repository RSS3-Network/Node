package rss3

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/rss3"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config          *engine.Config
	ethereumClient  ethereum.Client
	tokenClient     token.Client
	stakingFilterer *rss3.StakingFilterer
}

func (w *worker) Name() string {
	return engine.RSS3.String()
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	switch task.Network() {
	case filter.NetworkEthereum:
		task := task.(*source.Task)
		return task.Transaction.To != nil && *task.Transaction.To == rss3.AddressStaking, nil
	default:
		return false, nil
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformRSS3))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*schema.Action
			err     error
		)

		switch {
		case w.matchStakingDeposited(ethereumTask, log):
			actions, err = w.handleStakingDeposited(ctx, ethereumTask, log)
		case w.matchStakingWithdrawn(ethereumTask, log):
			actions, err = w.handleStakingWithdrawn(ctx, ethereumTask, log)
		case w.matchStakingRewardsClaimed(ethereumTask, log):
			actions, err = w.handleStakingRewardsClaimed(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		// Overwrite the type for feed.
		for _, action := range actions {
			feed.Type = action.Type
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

func (w *worker) matchStakingDeposited(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressStaking && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], rss3.EventHashStakingDeposited)
}

func (w *worker) matchStakingWithdrawn(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressStaking && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], rss3.EventHashStakingWithdrawn)
}

func (w *worker) matchStakingRewardsClaimed(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressStaking && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], rss3.EventHashRewardsClaimed)
}

func (w *worker) handleStakingDeposited(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakingFilterer.ParseDeposited(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Deposited event: %w", err)
	}

	period := metadata.ExchangeStakingPeriod{
		Start: time.Unix(event.Start.Int64(), 0),
		End:   time.Unix(event.Start.Int64(), 0).Add(time.Duration(event.Duration.Int64()) * time.Second),
	}

	action, err := w.buildExchangeStakingAction(ctx, task, event.Staker, event.Staker, event.Amount, metadata.ActionExchangeStakingStake, &period)
	if err != nil {
		return nil, fmt.Errorf("build exchange staking action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleStakingWithdrawn(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakingFilterer.ParseWithdrawn(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Withdrawn event: %w", err)
	}

	action, err := w.buildExchangeStakingAction(ctx, task, event.From, event.Receiver, event.Amount, metadata.ActionExchangeStakingUnstake, nil)
	if err != nil {
		return nil, fmt.Errorf("build action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleStakingRewardsClaimed(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakingFilterer.ParseRewardsClaimed(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse RewardsClaimed event: %w", err)
	}

	action, err := w.buildExchangeStakingAction(ctx, task, event.From, event.Receiver, event.RewardAmount, metadata.ActionExchangeStakingClaim, nil)
	if err != nil {
		return nil, fmt.Errorf("build action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildExchangeStakingAction(ctx context.Context, task *source.Task, from, to common.Address, tokenValue *big.Int, stakingAction metadata.ExchangeStakingAction, period *metadata.ExchangeStakingPeriod) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.Chain, &rss3.AddressToken, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", rss3.AddressToken, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := schema.Action{
		Type:     filter.TypeExchangeStaking,
		Platform: filter.PlatformRSS3.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeStaking{
			Action: stakingAction,
			Token:  *tokenMetadata,
			Period: period,
		},
	}

	return &action, nil
}

func NewWorker(config *engine.Config) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	instance.stakingFilterer = lo.Must(rss3.NewStakingFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
