package rss3

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/rss3"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config          *config.Module
	ethereumClient  ethereum.Client
	tokenClient     token.Client
	stakingFilterer *rss3.StakingFilterer
}

func (w *worker) Name() string {
	return filter.RSS3.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			rss3.AddressStaking,
		},
		LogTopics: []common.Hash{
			rss3.EventHashStakingDeposited,
			rss3.EventHashStakingWithdrawn,
			rss3.EventHashRewardsClaimed,
		},
	}
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	switch task.GetNetwork() {
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

	// Match and handle logs.
	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*schema.Action
			err     error
		)

		switch {
		case w.matchStakingDeposited(ethereumTask, log):
			actions, err = w.transformStakingDeposited(ctx, ethereumTask, log)
		case w.matchStakingWithdrawn(ethereumTask, log):
			actions, err = w.transformStakingWithdrawn(ctx, ethereumTask, log)
		case w.matchStakingRewardsClaimed(ethereumTask, log):
			actions, err = w.transformStakingRewardsClaimed(ctx, ethereumTask, log)
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

func (w *worker) transformStakingDeposited(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
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

func (w *worker) transformStakingWithdrawn(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
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

func (w *worker) transformStakingRewardsClaimed(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
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
	// The Token always is $RSS3.
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &rss3.AddressToken, nil, task.Header.Number)
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

// NewWorker creates a new RSS3 worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
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

	// Initialize RSS3 staking filterer.
	instance.stakingFilterer = lo.Must(rss3.NewStakingFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
