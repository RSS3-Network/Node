package rainbow

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/rainbow"
	"github.com/rss3-network/node/provider/ethereum/contract/weth"
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
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	erc20Filterer  *erc20.ERC20Filterer
	weth9Filterer  *weth.WETH9Filterer
}

// Name returns the name of the worker.
func (w *worker) Name() string {
	return decentralized.Rainbow.String()
}

// Platform returns the platform of the worker.
func (w *worker) Platform() string {
	return decentralized.PlatformRainbow.String()
}

// Network returns the supported networks of the worker.
func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.BinanceSmartChain,
		network.Base,
		network.Polygon,
		network.Arbitrum,
		network.Avalanche,
		network.Optimism,
	}
}

// Tags returns the tags of the worker.
func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
		tag.Transaction,
	}
}

// Types returns the types of the worker.
func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeSwap,
		typex.TransactionTransfer,
	}
}

// Filter returns the filter of the worker.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{}
}

// Transform transforms the task into an activity.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	if !w.matchSwapTransaction(ethereumTask) {
		return nil, nil
	}

	actions, err := w.transformSwapTransaction(ctx, ethereumTask)

	if err != nil {
		return nil, fmt.Errorf("handle ethereum swap transaction: %w", err)
	}

	activity.Actions = append(activity.Actions, actions...)

	if len(activity.Actions) == 0 {
		return nil, fmt.Errorf("no actions")
	}

	activity.Type = typex.ExchangeSwap

	return activity, nil
}

// matchSwapTransaction checks if the transaction matches a swap transaction.
func (w *worker) matchSwapTransaction(task *source.Task) bool {
	if task.Transaction == nil || task.Transaction.To == nil {
		return false
	}

	return *task.Transaction.To == rainbow.AddressRouter &&
		contract.MatchMethodIDs(
			task.Transaction.Input,
			rainbow.MethodIDRouterFillQuoteEthToToken,
			rainbow.MethodIDRouterFillQuoteTokenToEth,
			rainbow.MethodIDRouterFillQuoteTokenToToken,
		)
}

// transformSwapTransaction transforms a swap transaction into actions.
func (w *worker) transformSwapTransaction(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	switch {
	case bytes.HasPrefix(task.Transaction.Input, rainbow.MethodIDRouterFillQuoteEthToToken[:4]):
		ethToTokenActions, err := w.transformEthToTokenSwap(ctx, task)
		if err != nil {
			return nil, fmt.Errorf("transform ETH to token swap: %w", err)
		}

		actions = append(actions, ethToTokenActions...)
	case bytes.HasPrefix(task.Transaction.Input, rainbow.MethodIDRouterFillQuoteTokenToEth[:4]):
		tokenToEthActions, err := w.transformTokenToEthSwap(ctx, task)
		if err != nil {
			return nil, fmt.Errorf("transform token to ETH swap: %w", err)
		}

		actions = append(actions, tokenToEthActions...)
	case bytes.HasPrefix(task.Transaction.Input, rainbow.MethodIDRouterFillQuoteTokenToToken[:4]):
		tokenToTokenActions, err := w.transformTokenToTokenSwap(ctx, task)
		if err != nil {
			return nil, fmt.Errorf("transform token to token swap: %w", err)
		}

		actions = append(actions, tokenToTokenActions...)
	default:
		return nil, fmt.Errorf("unknown swap type")
	}

	return actions, nil
}

// transformEthToTokenSwap transforms an ETH to token swap transaction.
func (w *worker) transformEthToTokenSwap(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var (
		actions  []*activityx.Action
		valueMap = make(map[*common.Address]*big.Int)
	)

	abi, err := rainbow.RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("get abi: %w", err)
	}

	method, err := abi.MethodById(rainbow.MethodIDRouterFillQuoteEthToToken[:4])
	if err != nil {
		return nil, fmt.Errorf("get method by id: %w", err)
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack values: %w", err)
	}

	var input rainbow.RouterFillQuoteEthToTokenInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, fmt.Errorf("copy input: %w", err)
	}

	w.simulationTransfer(valueMap, true, nil, new(big.Int).Sub(task.Transaction.Value, input.FeeAmount))

	feeAction, err := w.buildTransactionTransferAction(ctx, task, task.Transaction.From, rainbow.AddressRouter, nil, input.FeeAmount)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	actions = append(actions, feeAction)

	return w.processSwapLogs(ctx, task, valueMap, actions)
}

// transformTokenToEthSwap transforms a token to ETH swap transaction.
func (w *worker) transformTokenToEthSwap(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var (
		actions  []*activityx.Action
		valueMap = make(map[*common.Address]*big.Int)
	)

	abi, err := rainbow.RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("get abi: %w", err)
	}

	method, err := abi.MethodById(rainbow.MethodIDRouterFillQuoteTokenToEth[:4])
	if err != nil {
		return nil, fmt.Errorf("get method by id: %w", err)
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack values: %w", err)
	}

	var input rainbow.RouterFillQuoteTokenToEthInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, fmt.Errorf("copy input: %w", err)
	}

	feePercentageBasisPoints := decimal.NewFromBigInt(input.FeePercentageBasisPoints, 0)

	return w.processSwapLogs(ctx, task, valueMap, actions, feePercentageBasisPoints)
}

// transformTokenToTokenSwap transforms a token to token swap transaction.
func (w *worker) transformTokenToTokenSwap(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var (
		actions  []*activityx.Action
		valueMap = make(map[*common.Address]*big.Int)
	)

	abi, err := rainbow.RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("get abi: %w", err)
	}

	method, err := abi.MethodById(rainbow.MethodIDRouterFillQuoteTokenToToken[:4])
	if err != nil {
		return nil, fmt.Errorf("get method by id: %w", err)
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack values: %w", err)
	}

	var input rainbow.RouterFillQuoteTokenToTokenInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, fmt.Errorf("copy input: %w", err)
	}

	feeAction, err := w.buildTransactionTransferAction(ctx, task, task.Transaction.From, rainbow.AddressRouter, &input.SellTokenAddress, input.FeeAmount)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	actions = append(actions, feeAction)

	return w.processSwapLogs(ctx, task, valueMap, actions)
}

// processSwapLogs processes the logs of a swap transaction.
func (w *worker) processSwapLogs(ctx context.Context, task *source.Task, valueMap map[*common.Address]*big.Int, actions []*activityx.Action, feePercentageBasisPoints ...decimal.Decimal) ([]*activityx.Action, error) {
	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch log.Topics[0] {
		case erc20.EventHashTransfer:
			if len(log.Topics) != 3 {
				continue
			}

			event, err := w.erc20Filterer.ParseTransfer(log.Export())
			if err != nil {
				zap.L().Error("parse event", zap.Error(err))
				continue
			}

			if event.From == task.Transaction.From {
				w.simulationTransfer(valueMap, true, &event.Raw.Address, event.Value)
			}

			if event.To == task.Transaction.From {
				w.simulationTransfer(valueMap, false, &event.Raw.Address, event.Value)
			}
		case weth.EventHashWithdrawal:
			event, err := w.weth9Filterer.ParseWithdrawal(log.Export())
			if err != nil {
				zap.L().Error("parse event", zap.Error(err))
				continue
			}

			w.simulationTransfer(valueMap, false, nil, event.Wad)
		}
	}

	tokenIn, tokenOut, err := w.findTokens(valueMap)
	if err != nil {
		return nil, err
	}

	if len(feePercentageBasisPoints) > 0 && !feePercentageBasisPoints[0].IsZero() {
		feeAction, err := w.buildFeeAction(ctx, task, tokenOut, valueMap[tokenOut], feePercentageBasisPoints[0])
		if err != nil {
			return nil, err
		}

		actions = append(actions, feeAction)
	}

	swapAction, err := w.buildExchangeSwapAction(ctx, task, task.Transaction.From, task.Transaction.From, tokenIn, tokenOut, valueMap[tokenIn], valueMap[tokenOut])
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions = append(actions, swapAction)

	return actions, nil
}

// findTokens finds the input and output tokens from the value map.
func (w *worker) findTokens(valueMap map[*common.Address]*big.Int) (*common.Address, *common.Address, error) {
	tokenIn, exists := lo.FindKeyBy(valueMap, func(_ *common.Address, value *big.Int) bool {
		return value.Sign() == -1
	})
	if !exists {
		return nil, nil, fmt.Errorf("token in not match")
	}

	tokenOut, exists := lo.FindKeyBy(valueMap, func(_ *common.Address, value *big.Int) bool {
		return value.Sign() == 1
	})
	if !exists {
		return nil, nil, fmt.Errorf("token out not match")
	}

	return tokenIn, tokenOut, nil
}

// buildFeeAction builds a fee action for the swap.
func (w *worker) buildFeeAction(ctx context.Context, task *source.Task, _ *common.Address, amountOut *big.Int, feePercentageBasisPoints decimal.Decimal) (*activityx.Action, error) {
	ethDiff := decimal.NewFromBigInt(amountOut, 0)
	fees := ethDiff.Mul(feePercentageBasisPoints).DivRound(decimal.NewFromInt(1e18), -1)

	return w.buildTransactionTransferAction(ctx, task, task.Transaction.From, rainbow.AddressRouter, nil, fees.BigInt())
}

// simulationTransfer simulates a token transfer in the value map.
func (w *worker) simulationTransfer(valueMap map[*common.Address]*big.Int, transferIn bool, token *common.Address, value *big.Int) {
	if valueMap[token] == nil {
		valueMap[token] = big.NewInt(0)
	}

	if transferIn {
		valueMap[token] = new(big.Int).Sub(valueMap[token], value)
	} else {
		valueMap[token] = new(big.Int).Add(valueMap[token], value)
	}
}

// buildExchangeSwapAction builds an exchange swap action.
func (w *worker) buildExchangeSwapAction(ctx context.Context, task *source.Task, from, to common.Address, tokenIn, tokenOut *common.Address, amountIn, amountOut *big.Int) (*activityx.Action, error) {
	tokenInAddress := lo.Ternary(tokenIn != nil && *tokenIn != ethereum.AddressGenesis, tokenIn, nil)
	tokenOutAddress := lo.Ternary(tokenOut != nil && *tokenOut != ethereum.AddressGenesis, tokenOut, nil)

	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenInAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token in metadata: %w", err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenOutAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token out metadata: %w", err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountOut, 0).Abs())

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

// buildTransactionTransferAction builds a transaction transfer action.
func (w *worker) buildTransactionTransferAction(ctx context.Context, task *source.Task, from, to common.Address, tokenAddress *common.Address, amount *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amount, 0))

	return &activityx.Action{
		Type:     typex.TransactionTransfer,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	var err error

	instance := worker{
		config: config,
	}

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize filterers.
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.weth9Filterer = lo.Must(weth.NewWETH9Filterer(ethereum.AddressGenesis, nil))

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
