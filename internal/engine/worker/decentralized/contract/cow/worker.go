package cow

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/internal/utils"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/cow"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                *config.Module
	ethereumClient        ethereum.Client
	tokenClient           token.Client
	cowSettlementFilterer *cow.SettlementFilterer
}

func (w *worker) Name() string {
	return decentralized.Cow.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformCow.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Arbitrum,
		network.Base,
		network.Gnosis,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeSwap,
	}
}

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			cow.AddressSettlement,
		},
		LogTopics: []common.Hash{
			cow.EventHashSettlementTrade,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchSettlementTradeLog(ethereumTask, log):
			actions, err := w.transformSettlementTradeLog(ctx, ethereumTask, log)
			if err != nil {
				zap.L().Error("failed to handle settlement trade log", zap.Error(err))

				continue
			}

			activity.Actions = append(activity.Actions, actions...)
		default:
			continue
		}
	}

	if len(activity.Actions) == 0 {
		return nil, fmt.Errorf("no actions")
	}

	activity.Type = typex.ExchangeSwap

	return activity, nil
}

func (w *worker) matchSettlementTradeLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], cow.EventHashSettlementTrade) &&
		contract.MatchAddresses(log.Address, cow.AddressSettlement)
}

func (w *worker) transformSettlementTradeLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.cowSettlementFilterer.ParseTrade(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Trade event: %w", err)
	}

	actions := make([]*activityx.Action, 0)

	swapAction, err := w.buildExchangeSwapAction(ctx, task, event.Owner, event.Owner, event.SellToken, event.BuyToken, new(big.Int).Sub(event.SellAmount, event.FeeAmount), event.BuyAmount)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions = append(actions, swapAction)

	return actions, nil
}

func (w *worker) buildExchangeSwapAction(ctx context.Context, task *source.Task, from, to, tokenIn, tokenOut common.Address, amountIn, amountOut *big.Int) (*activityx.Action, error) {
	tokenInAddress := lo.Ternary(tokenIn != cow.AddressETH, &tokenIn, nil)
	tokenOutAddress := lo.Ternary(tokenOut != cow.AddressETH, &tokenOut, nil)

	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenInAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token in metadata: %w", err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(amountIn), 0))

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenOutAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token out metadata: %w", err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(amountOut), 0))

	return &activityx.Action{
		Type:     typex.ExchangeSwap,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}, nil
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		config:                config,
		cowSettlementFilterer: lo.Must(cow.NewSettlementFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
