package oneinch

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/1inch"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/weth"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	oneinchExchangeFilterer            *oneinch.ExchangeFilterer
	oneinchAggregationRouterV2Filterer *oneinch.AggregationRouterV2Filterer
	oneinchAggregationRouterV3Filterer *oneinch.AggregationRouterV3Filterer
	oneinchAggregationRouterV4Filterer *oneinch.AggregationRouterV4Filterer
	oneinchAggregationRouterV5Filterer *oneinch.AggregationRouterV5Filterer
	erc20Filterer                      *erc20.ERC20Filterer
	weth9Filterer                      *weth.WETH9Filterer
	ethereumClient                     ethereum.Client
	tokenClient                        token.Client
}

func (w *worker) Name() string {
	//return filter.Oneinch.String()
	return filter.Oneinch.String()
}

// Filter returns a filter for source.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			oneinch.AddressAggregationRouterV2,
			oneinch.AddressAggregationRouterV3,
			oneinch.AddressAggregationRouterV4,
			oneinch.AddressAggregationRouterV5,
			oneinch.AddressExchange2,
			oneinch.AddressEther,
		},
		LogTopics: []common.Hash{
			oneinch.EventHashExchangeSwapped,
			oneinch.EventHashAggregationRouterV2Swapped,
			oneinch.EventHashAggregationRouterV3Swapped,
		},
	}
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.GetNetwork().Source() == filter.NetworkEthereumSource, nil
}

// Transform 1inch task to feed.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	// Cast the task to a 1inch task.
	oneinchTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the feed.
	feed, err := task.BuildFeed(schema.WithFeedPlatform(filter.Platform1inch))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	switch *oneinchTask.Transaction.To {
	case // Swap routers
		oneinch.AddressExchange2,
		oneinch.AddressAggregationRouterV2,
		oneinch.AddressAggregationRouterV3,
		oneinch.AddressAggregationRouterV4,
		oneinch.AddressAggregationRouterV5:
		err = w.handleEthereumExchangeSwapTransaction(ctx, oneinchTask, feed)
	default:
		return nil, fmt.Errorf("unknown transaction %s", task.ID())
	}

	if err != nil {
		zap.L().Warn("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

		return nil, err
	}

	return feed, nil
}

// handleEthereumExchangeSwapTransaction handles exchange swap transaction.
func (w *worker) handleEthereumExchangeSwapTransaction(ctx context.Context, task *source.Task, feed *schema.Feed) error {
	feed.Type = filter.TypeExchangeSwap

	var (
		actions []*schema.Action
		err     error
	)

	switch *task.Transaction.To {
	case
		oneinch.AddressExchange2,
		oneinch.AddressAggregationRouterV2,
		oneinch.AddressAggregationRouterV3:
		actions, err = w.handleEthereumExplicitAggregationRouterTransaction(ctx, task)
	case
		oneinch.AddressAggregationRouterV4,
		oneinch.AddressAggregationRouterV5:
		actions, err = w.handleEthereumImplicitAggregationRouterTransaction(ctx, task)
	default:
		return fmt.Errorf("unknown transaction %s", task.ID())
	}

	if err != nil {
		return err
	}

	feed.Actions = append(feed.Actions, actions...)

	return nil
}

// handleEthereumExplicitAggregationRouterTransaction handles explicit aggregation router transaction.
func (w *worker) handleEthereumImplicitAggregationRouterTransaction(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	var (
		sender, receipt *common.Address
		err             error
	)

	switch *task.Transaction.To {
	case oneinch.AddressAggregationRouterV4:
		sender, receipt, err = w.parseEthereumAggregationRouterV4TransactionInput(ctx, *task)
	case oneinch.AddressAggregationRouterV5:
		sender, receipt, err = w.parseEthereumAggregationRouterV5TransactionInput(ctx, *task)
	default:
		return nil, fmt.Errorf("unsupported contract: %s", *task.Transaction.To)
	}

	if err != nil {
		return nil, fmt.Errorf("parse transaction input: %w", err)
	}

	valueMap := make(map[common.Address]*big.Int)

	// Pay with native token
	if task.Transaction.Value.Sign() != 0 {
		w.simulationTransfer(valueMap, true, oneinch.AddressEther, task.Transaction.Value)
	}

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch log.Topics[0] {
		case erc20.EventHashTransfer:
			// Filter ERC-721 transfer event
			if len(log.Topics) != 3 {
				continue
			}

			event, err := w.erc20Filterer.ParseTransfer(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse event: %w", err)
			}

			// Transfer in
			if event.From == *sender {
				w.simulationTransfer(valueMap, true, event.Raw.Address, event.Value)
			}

			// Transfer out
			if event.To == *receipt {
				w.simulationTransfer(valueMap, false, event.Raw.Address, event.Value)
			}
		case weth.EventHashWithdrawal:
			event, err := w.weth9Filterer.ParseWithdrawal(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse event: %w", err)
			}

			// Transfer out
			w.simulationTransfer(valueMap, false, oneinch.AddressEther, event.Wad)
		}
	}

	tokenIn, exists := lo.FindKeyBy(valueMap, func(token common.Address, value *big.Int) bool {
		return value.Sign() == -1 // value < 0
	})
	if !exists {
		return nil, fmt.Errorf("token in not match")
	}

	tokenOut, exists := lo.FindKeyBy(valueMap, func(token common.Address, value *big.Int) bool {
		return value.Sign() == 1 // value > 0
	})
	if !exists {
		return nil, fmt.Errorf("token out not match")
	}

	action, err := w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, *sender, *receipt, tokenIn, tokenOut, valueMap[tokenIn], valueMap[tokenOut])
	if err != nil {
		return nil, fmt.Errorf("build action: %w", err)
	}

	return []*schema.Action{
		action,
	}, nil
}

// simulationTransfer simulates the transfer of tokens.
func (w *worker) simulationTransfer(valueMap map[common.Address]*big.Int, transferIn bool, token common.Address, value *big.Int) {
	if valueMap[token] == nil {
		valueMap[token] = big.NewInt(0)
	}

	if transferIn {
		valueMap[token] = new(big.Int).Sub(valueMap[token], new(big.Int).Abs(value))
	} else {
		valueMap[token] = new(big.Int).Add(valueMap[token], new(big.Int).Abs(value))
	}
}

// parseEthereumAggregationRouterV5TransactionInput parses the transaction input of the aggregation router v5.
func (w *worker) parseEthereumAggregationRouterV5TransactionInput(ctx context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	switch {
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5ClipperSwapTo):
		return w.parseEthereumAggregationRouterV5TransactionClipperSwapToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5ClipperSwapToWithPermit):
		return w.parseEthereumAggregationRouterV5TransactionClipperSwapToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5FillOrderRFQTo):
		return w.parseEthereumAggregationRouterV5TransactionFillOrderRFQToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5FillOrderRFQToWithPermit):
		return w.parseEthereumAggregationRouterV5TransactionFillOrderRFQToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5FillOrderTo):
		return w.parseEthereumAggregationRouterV5TransactionFillOrderToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5FillOrderToWithPermit):
		return w.parseEthereumAggregationRouterV5TransactionFillOrderToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5Swap):
		return w.parseEthereumAggregationRouterV5TransactionSwapInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5UniswapV3SwapTo):
		return w.parseEthereumAggregationRouterV5TransactionUniswapV3SwapToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5UniswapV3SwapToWithPermit):
		return w.parseEthereumAggregationRouterV5TransactionUniswapV3SwapToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5UnoswapTo):
		return w.parseEthereumAggregationRouterV5TransactionUnoswapToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV5UnoswapToWithPermit):
		return w.parseEthereumAggregationRouterV5TransactionUnoswapToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(
		task.Transaction.Input,
		oneinch.MethodIDAggregationRouterV5ClipperSwap,
		oneinch.MethodIDAggregationRouterV5FillOrder,
		oneinch.MethodIDAggregationRouterV5FillOrderRFQ,
		oneinch.MethodIDAggregationRouterV5FillOrderRFQCompact,
		oneinch.MethodIDAggregationRouterV5UniswapV3Swap,
		oneinch.MethodIDAggregationRouterV5UniswapV3SwapCallback,
		oneinch.MethodIDAggregationRouterV5Unoswap,
	):
		return &task.Transaction.From, &task.Transaction.From, nil
	default:
		return nil, nil, fmt.Errorf("invalid transaction input")
	}
}

// parseEthereumAggregationRouterV5TransactionClipperSwapToInput parses the transaction input of the aggregation router v5 clipper swap to.
func (w *worker) parseEthereumAggregationRouterV5TransactionClipperSwapToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5ClipperSwapTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5ClipperSwapToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV5TransactionClipperSwapToWithPermitInput parses the transaction input of the aggregation router v5 clipper swap to with permit.
func (w *worker) parseEthereumAggregationRouterV5TransactionClipperSwapToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5ClipperSwapToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5ClipperSwapToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV5TransactionFillOrderRFQToInput parses the transaction input of the aggregation router v5 fill order RFQ to.
func (w *worker) parseEthereumAggregationRouterV5TransactionFillOrderRFQToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5FillOrderRFQTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5FillOrderRFQToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV5TransactionFillOrderRFQToWithPermitInput parses the transaction input of the aggregation router v5 fill order RFQ to with permit.
func (w *worker) parseEthereumAggregationRouterV5TransactionFillOrderRFQToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5FillOrderRFQToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5FillOrderRFQToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV5TransactionFillOrderToInput parses the transaction input of the aggregation router v5 fill order to.
func (w *worker) parseEthereumAggregationRouterV5TransactionFillOrderToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5FillOrderTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5FillOrderToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV5TransactionFillOrderToWithPermitInput parses the transaction input of the aggregation router v5 fill order to with permit.
func (w *worker) parseEthereumAggregationRouterV5TransactionFillOrderToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5FillOrderToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5FillOrderToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV5TransactionSwapInput parses the transaction input of the aggregation router v5 swap.
func (w *worker) parseEthereumAggregationRouterV5TransactionSwapInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5Swap[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5SwapInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Desc.DstReceiver, nil
}

// parseEthereumAggregationRouterV5TransactionUniswapV3SwapToInput parses the transaction input of the aggregation router v5 uniswap v3 swap to.
func (w *worker) parseEthereumAggregationRouterV5TransactionUniswapV3SwapToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5UniswapV3SwapTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5UniswapV3SwapToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV5TransactionUniswapV3SwapToWithPermitInput parses the transaction input of the aggregation router v5 uniswap v3 swap to with permit.
func (w *worker) parseEthereumAggregationRouterV5TransactionUniswapV3SwapToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5UniswapV3SwapToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5UniswapV3SwapToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV5TransactionUnoswapToInput parses the transaction input of the aggregation router v5 unoswap to.
func (w *worker) parseEthereumAggregationRouterV5TransactionUnoswapToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5UnoswapTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5UnoswapToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV5TransactionUnoswapToWithPermitInput parses the transaction input of the aggregation router v5 unoswap to with permit.
func (w *worker) parseEthereumAggregationRouterV5TransactionUnoswapToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV5MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV5UnoswapToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV5UnoswapToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV4TransactionInput parses the transaction input of the aggregation router v4.
func (w *worker) parseEthereumAggregationRouterV4TransactionInput(ctx context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	switch {
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4ClipperSwapTo):
		return w.parseEthereumAggregationRouterV4TransactionClipperSwapToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4ClipperSwapToWithPermit):
		return w.parseEthereumAggregationRouterV4TransactionClipperSwapToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4FillOrderRFQTo):
		return w.parseEthereumAggregationRouterV4TransactionFillOrderRFQToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4FillOrderRFQToWithPermit):
		return w.parseEthereumAggregationRouterV4TransactionFillOrderRFQToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4Swap):
		return w.parseEthereumAggregationRouterV4TransactionSwapInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4UniswapV3SwapTo):
		return w.parseEthereumAggregationRouterV4TransactionUniswapV3SwapToInput(ctx, task)
	case contract.MatchMethodIDs(task.Transaction.Input, oneinch.MethodIDAggregationRouterV4UniswapV3SwapToWithPermit):
		return w.parseEthereumAggregationRouterV4TransactionUniswapV3SwapToWithPermitInput(ctx, task)
	case contract.MatchMethodIDs(
		task.Transaction.Input,
		oneinch.MethodIDAggregationRouterV4ClipperSwap,
		oneinch.MethodIDAggregationRouterV4FillOrderRFQ,
		oneinch.MethodIDAggregationRouterV4UniswapV3Swap,
		oneinch.MethodIDAggregationRouterV4UniswapV3SwapCallback,
		oneinch.MethodIDAggregationRouterV4Unoswap,
		oneinch.MethodIDAggregationRouterV4UnoswapWithPermit,
	):
		return &task.Transaction.From, &task.Transaction.From, nil
	default:
		return nil, nil, fmt.Errorf("invalid transaction input")
	}
}

// parseEthereumAggregationRouterV4TransactionUniswapV3SwapToWithPermitInput parses the transaction input of the aggregation router v4 uniswap v3 swap to with permit.
func (w *worker) parseEthereumAggregationRouterV4TransactionUniswapV3SwapToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4UniswapV3SwapToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4UniswapV3SwapToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV4TransactionUniswapV3SwapToInput parses the transaction input of the aggregation router v4 uniswap v3 swap to.
func (w *worker) parseEthereumAggregationRouterV4TransactionUniswapV3SwapToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4UniswapV3SwapTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4UniswapV3SwapToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV4TransactionSwapInput parses the transaction input of the aggregation router v4 swap.
func (w *worker) parseEthereumAggregationRouterV4TransactionSwapInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4Swap[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4SwapInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Desc.DstReceiver, nil
}

// parseEthereumAggregationRouterV4TransactionFillOrderRFQToWithPermitInput parses the transaction input of the aggregation router v4 fill order RFQ to with permit.
func (w *worker) parseEthereumAggregationRouterV4TransactionFillOrderRFQToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4FillOrderRFQToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4FillOrderRFQToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV4TransactionFillOrderRFQToInput parses the transaction input of the aggregation router v4 fill order RFQ to.
func (w *worker) parseEthereumAggregationRouterV4TransactionFillOrderRFQToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4FillOrderRFQTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4FillOrderRFQToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Target, nil
}

// parseEthereumAggregationRouterV4TransactionClipperSwapToWithPermitInput parses the transaction input of the aggregation router v4 clipper swap to with permit.
func (w *worker) parseEthereumAggregationRouterV4TransactionClipperSwapToWithPermitInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4ClipperSwapToWithPermit[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4ClipperSwapToWithPermitInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// parseEthereumAggregationRouterV4TransactionClipperSwapToInput parses the transaction input of the aggregation router v4 clipper swap to.
func (w *worker) parseEthereumAggregationRouterV4TransactionClipperSwapToInput(_ context.Context, task source.Task) (sender, receiver *common.Address, err error) {
	abi, err := oneinch.AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, nil, err
	}

	method, err := abi.MethodById(oneinch.MethodIDAggregationRouterV4ClipperSwapTo[:])
	if err != nil {
		return nil, nil, err
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, nil, fmt.Errorf("unpack values: %w", err)
	}

	var input oneinch.AggregationRouterV4ClipperSwapToInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, nil, fmt.Errorf("copy values: %w", err)
	}

	return &task.Transaction.From, &input.Recipient, nil
}

// handleEthereumExplicitAggregationRouterTransaction handles the explicit aggregation router transaction.
func (w *worker) handleEthereumExplicitAggregationRouterTransaction(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	actions := make([]*schema.Action, 0)

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			action *schema.Action
			err    error
		)

		switch log.Topics[0] {
		case oneinch.EventHashExchangeSwapped:
			action, err = w.handleEthereumExchangeSwappedLog(ctx, task, log)
		case oneinch.EventHashAggregationRouterV2Swapped:
			action, err = w.handleEthereumAggregationRouterV2SwappedLog(ctx, task, log)
		case oneinch.EventHashAggregationRouterV3Swapped:
			action, err = w.handleEthereumAggregationRouterV3SwappedLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			zap.L().Warn("handle ethereum swap transaction", zap.Error(err), zap.String("worker", w.Name()), zap.String("task", task.ID()))

			continue
		}

		actions = append(actions, action)
	}

	return actions, nil
}

// handleEthereumAggregationRouterV3SwappedLog handles the aggregation router v3 swapped log.
func (w *worker) handleEthereumAggregationRouterV3SwappedLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	if log.Address != oneinch.AddressAggregationRouterV3 {
		return nil, fmt.Errorf("unofficial aggregation router v3 contract %s", log.Address)
	}

	event, err := w.oneinchAggregationRouterV3Filterer.ParseSwapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseSwapped event: %w", err)
	}

	return w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, event.Sender, event.DstReceiver, event.SrcToken, event.DstToken, event.SpentAmount, event.ReturnAmount)
}

// handleEthereumAggregationRouterV2SwappedLog handles the aggregation router v2 swapped log.
func (w *worker) handleEthereumAggregationRouterV2SwappedLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	if log.Address != oneinch.AddressAggregationRouterV2 {
		return nil, fmt.Errorf("unofficial aggregation router v2 contract %s", log.Address)
	}

	event, err := w.oneinchAggregationRouterV2Filterer.ParseSwapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseSwapped event: %w", err)
	}

	return w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, event.Sender, event.DstReceiver, event.SrcToken, event.DstToken, event.SpentAmount, event.ReturnAmount)
}

// handleEthereumExchangeSwappedLog handles the exchange swapped log.
func (w *worker) handleEthereumExchangeSwappedLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	if log.Address != oneinch.AddressExchange2 {
		return nil, fmt.Errorf("unofficial exchange contract %s", log.Address)
	}

	event, err := w.oneinchExchangeFilterer.ParseSwapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseSwapped event: %w", err)
	}

	return w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, task.Transaction.From, task.Transaction.From, event.FromToken, event.ToToken, event.FromAmount, event.ToAmount)
}

// buildEthereumExchangeSwapAction builds the exchange swap action.
func (w *worker) buildEthereumExchangeSwapAction(ctx context.Context, blockNumber *big.Int, chain uint64, from, to, tokenIn, tokenOut common.Address, amountIn, amountOut *big.Int) (*schema.Action, error) {
	var (
		tokenInAddress  = lo.Ternary(tokenIn != oneinch.AddressEther, lo.ToPtr(tokenIn), nil)
		tokenOutAddress = lo.Ternary(tokenOut != oneinch.AddressEther, lo.ToPtr(tokenOut), nil)
	)

	tokenInMetadata, err := w.tokenClient.Lookup(ctx, chain, tokenInAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenIn, err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, chain, tokenOutAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenOut, err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountOut, 0).Abs())

	action := schema.Action{
		Tag:      filter.TagExchange,
		Type:     filter.TypeExchangeSwap,
		Platform: filter.Platform1inch.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}

	return &action, nil
}

// NewWorker returns a new 1inch worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		oneinchExchangeFilterer:            lo.Must(oneinch.NewExchangeFilterer(ethereum.AddressGenesis, nil)),
		oneinchAggregationRouterV2Filterer: lo.Must(oneinch.NewAggregationRouterV2Filterer(ethereum.AddressGenesis, nil)),
		oneinchAggregationRouterV3Filterer: lo.Must(oneinch.NewAggregationRouterV3Filterer(ethereum.AddressGenesis, nil)),
		oneinchAggregationRouterV4Filterer: lo.Must(oneinch.NewAggregationRouterV4Filterer(ethereum.AddressGenesis, nil)),
		oneinchAggregationRouterV5Filterer: lo.Must(oneinch.NewAggregationRouterV5Filterer(ethereum.AddressGenesis, nil)),
		erc20Filterer:                      lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil)),
		weth9Filterer:                      lo.Must(weth.NewWETH9Filterer(ethereum.AddressGenesis, nil)),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
