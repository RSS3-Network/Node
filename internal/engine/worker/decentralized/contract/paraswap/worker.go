package paraswap

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
	"github.com/rss3-network/node/provider/ethereum/contract/paraswap"
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
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	ethereumClient   ethereum.Client
	tokenClient      token.Client
	paraswapV5Filter *paraswap.V5ParaSwapFilterer
}

func (w *worker) Name() string {
	return decentralized.Paraswap.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformParaswap.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Optimism,
		network.BinanceSmartChain,
		network.Avalanche,
		network.Arbitrum,
		network.Polygon,
		network.Base,
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
			paraswap.AddressV5ParaSwap,
			paraswap.AddressV5ParaSwapBase,
			paraswap.AddressV5ParaSwapBinanceSmartChain,
			paraswap.AddressV5ParaSwapOptimism,
			paraswap.AddressV5ParaSwapPolygon,
			paraswap.AddressV5ParaSwapArbitrum,
			paraswap.AddressV5ParaSwapAvalanche,
		},
		LogTopics: []common.Hash{
			paraswap.EventHashV3Swapped,
			paraswap.EventHashV3Bought,
			paraswap.EventHashSwappedDirect,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	activity.Type = typex.ExchangeSwap
	activity.Actions = w.transformSwapTransaction(ctx, ethereumTask)

	return activity, nil
}

func (w *worker) transformSwapTransaction(ctx context.Context, ethereumTask *source.Task) (actions []*activityx.Action) {
	for _, log := range ethereumTask.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			buffer []*activityx.Action
			err    error
		)

		switch {
		case w.matchV3SwappedLog(ethereumTask, log):
			buffer, err = w.transformV3SwappedLog(ctx, ethereumTask, log)
		case w.matchV3BoughtLog(ethereumTask, log):
			buffer, err = w.transformV3BoughtLog(ctx, ethereumTask, log)
		case w.matchSwappedDirectLog(ethereumTask, log):
			buffer, err = w.transformSwappedDirectLog(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			continue
		}

		actions = append(actions, buffer...)
	}

	return actions
}

func (w *worker) matchV3SwappedLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], paraswap.EventHashV3Swapped) &&
		contract.MatchAddresses(log.Address, paraswap.AddressV5ParaSwap, paraswap.AddressV5ParaSwapBase)
}

func (w *worker) matchV3BoughtLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], paraswap.EventHashV3Bought) &&
		contract.MatchAddresses(log.Address, paraswap.AddressV5ParaSwap, paraswap.AddressV5ParaSwapBase)
}

func (w *worker) matchSwappedDirectLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], paraswap.EventHashSwappedDirect) &&
		contract.MatchAddresses(log.Address, paraswap.AddressV5ParaSwap, paraswap.AddressV5ParaSwapBase)
}

func (w *worker) transformV3SwappedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.paraswapV5Filter.ParseSwappedV3(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse SwappedV3 event: %w", err)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, event.Beneficiary, event.Beneficiary, event.SrcToken, event.DestToken, event.SrcAmount, event.ReceivedAmount)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformV3BoughtLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.paraswapV5Filter.ParseBoughtV3(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse BoughtV3 event: %w", err)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, event.Beneficiary, event.Beneficiary, event.SrcToken, event.DestToken, event.SrcAmount, event.ReceivedAmount)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformSwappedDirectLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.paraswapV5Filter.ParseSwappedDirect(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse SwappedDirect event: %w", err)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, event.Beneficiary, event.Beneficiary, event.SrcToken, event.DestToken, event.SrcAmount, event.ReceivedAmount)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) buildExchangeSwapAction(ctx context.Context, task *source.Task, sender, receiver common.Address, tokenIn, tokenOut common.Address, amountIn, amountOut *big.Int) (*activityx.Action, error) {
	tokenInAddress := lo.Ternary(tokenIn != paraswap.AddressETH, &tokenIn, nil)
	tokenOutAddress := lo.Ternary(tokenOut != paraswap.AddressETH, &tokenOut, nil)

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

	action := activityx.Action{
		Type:     typex.ExchangeSwap,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}

	return &action, nil
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		paraswapV5Filter: lo.Must(paraswap.NewV5ParaSwapFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
