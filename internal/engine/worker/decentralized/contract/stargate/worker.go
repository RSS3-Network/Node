package stargate

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/layerzero"
	"github.com/rss3-network/node/provider/ethereum/contract/stargate"
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
	config               *config.Module
	ethereumClient       ethereum.Client
	tokenClient          token.Client
	stargatePoolFilterer *stargate.PoolFilterer
}

func (w *worker) Name() string {
	return decentralized.Stargate.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformStargate.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Arbitrum,
		network.Optimism,
		network.Base,
		network.Linea,
		network.Avalanche,
		network.BinanceSmartChain,
		network.Polygon,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Transaction,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.TransactionBridge,
		typex.TransactionTransfer,
	}
}

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	matched, err := w.matchStargateTransaction(ctx, task)
	if err != nil {
		zap.L().Error("match task", zap.String("task.id", task.ID()), zap.Error(err))

		return nil, nil
	}

	// If the task does not meet the filter conditions, it will be discarded.
	if !matched {
		zap.L().Warn("unmatched task", zap.String("task.id", task.ID()))

		return nil, nil
	}

	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*activityx.Action
			err     error
		)

		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchPoolSwapLog(ethereumTask, log):
			actions, err = w.transformLPoolSwapLog(ctx, ethereumTask, log)
		case w.matchPoolSwapRemoteLog(ethereumTask, log):
			actions, err = w.transformLPoolSwapRemoteLog(ctx, ethereumTask, log)
		default:
			zap.L().Warn("unsupported log", zap.String("task", task.ID()), zap.Uint("topic.index", log.Index))

			continue
		}

		if err != nil {
			zap.L().Warn("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			return nil, err
		}

		activity.Type = typex.TransactionBridge
		activity.Actions = append(activity.Actions, actions...)
	}

	return activity, nil
}

// matchStargateTransaction matches the Stargate contract.
func (w *worker) matchStargateTransaction(_ context.Context, task engine.Task) (bool, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return false, fmt.Errorf("invalid task type: %T", task)
	}

	if ethereumTask.Transaction.To == nil {
		return false, nil
	}

	if contract.MatchAddresses(*ethereumTask.Transaction.To, stargate.RouterETHAddresses()...) || contract.MatchAddresses(*ethereumTask.Transaction.To, stargate.RouterAddresses()...) {
		return true, nil
	}

	return lo.ContainsBy(ethereumTask.Receipt.Logs, func(log *ethereum.Log) bool {
		return len(log.Topics) > 0 && contract.MatchEventHashes(
			log.Topics[0],
			stargate.EventHashPoolCreditChainPath,
			stargate.EventHashPoolSwap,
			stargate.EventHashPoolSwapRemote,
		)
	}), nil
}

// matchPoolSwapLog the pool swap log.
func (w *worker) matchPoolSwapLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Topics[0] == stargate.EventHashPoolSwap
}

// matchPoolSwapRemoteLog the pool swap remote log.
func (w *worker) matchPoolSwapRemoteLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Topics[0] == stargate.EventHashPoolSwapRemote
}

// transformLPoolSwapLog transforms the pool swap log.
func (w *worker) transformLPoolSwapLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.stargatePoolFilterer.ParseSwap(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Swap event: %w", err)
	}

	pool, err := stargate.NewPoolCaller(event.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new poll caller: %w", err)
	}

	callOptions := bind.CallOpts{
		BlockNumber: task.Header.Number,
		Context:     ctx,
	}

	factoryAddress, exists := stargate.FactoryAddress(task.Network)
	if !exists {
		return nil, fmt.Errorf("unsupport chain: %s", task.Network)
	}

	factoryCaller, err := stargate.NewFactoryCaller(factoryAddress, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new factory caller: %w", err)
	}

	poolID, err := pool.PoolId(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get pool id: %w", err)
	}

	poolAddress, err := factoryCaller.GetPool(&callOptions, poolID)
	if err != nil {
		return nil, fmt.Errorf("get pool by id: %w", err)
	}

	if event.Raw.Address != poolAddress {
		return nil, fmt.Errorf("invalid pool address: %s", event.Raw.Address)
	}

	tokenAddress, err := pool.Token(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get token address: %w", err)
	}

	targetChain, exists := stargate.EthereumChain(event.ChainId)
	if !exists {
		return nil, fmt.Errorf("invalid chain id: %d", event.ChainId)
	}

	sender := lo.Ternary(stargate.IsSGETH(tokenAddress), task.Transaction.From, event.From)

	bridgeAction, err := w.buildTransactionBridgeAction(ctx, task.ChainID, sender, sender, task.Network, targetChain, metadata.ActionTransactionBridgeDeposit, lo.Ternary(stargate.IsSGETH(tokenAddress), nil, &tokenAddress), event.AmountSD, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	layerzeroFeeValue := decimal.NewFromBigInt(task.Transaction.Value, 0)
	if stargate.IsSGETH(tokenAddress) {
		layerzeroFeeValue = layerzeroFeeValue.
			Sub(decimal.NewFromBigInt(event.AmountSD, 0)).
			Sub(decimal.NewFromBigInt(event.EqFee, 0)).
			Sub(decimal.NewFromBigInt(event.ProtocolFee, 0)).
			Sub(decimal.NewFromBigInt(event.LpFee, 0))
	}

	layerzeroUltraLightNodeAddress, exists := layerzero.UltraLightNodeAddress(task.Network)
	if !exists {
		return nil, fmt.Errorf("unsupport chain: %s", task.Network)
	}

	layerzeroFeeAction, err := w.buildTransactionTransferAction(ctx, task, task.Transaction.From, layerzeroUltraLightNodeAddress, nil, layerzeroFeeValue.BigInt())
	if err != nil {
		return nil, fmt.Errorf("build transction transfer action: %w", err)
	}

	stargateFee := decimal.Zero.
		Add(decimal.NewFromBigInt(event.EqFee, 0)).
		Add(decimal.NewFromBigInt(event.ProtocolFee, 0)).
		Add(decimal.NewFromBigInt(event.LpFee, 0))

	stargateFeeAction, err := w.buildTransactionTransferAction(ctx, task, task.Transaction.From, event.Raw.Address, lo.Ternary(stargate.IsSGETH(tokenAddress), nil, &tokenAddress), stargateFee.BigInt())
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	actions := []*activityx.Action{
		layerzeroFeeAction,
		stargateFeeAction,
		bridgeAction,
	}

	return actions, nil
}

// transformLPoolSwapRemoteLog transforms the pool swap remote log.
func (w *worker) transformLPoolSwapRemoteLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	swapRemoteEvent, err := w.stargatePoolFilterer.ParseSwapRemote(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Swap event: %w", err)
	}

	poolCreditChainPathLog, exists := lo.Find(task.Receipt.Logs, func(log *ethereum.Log) bool {
		return len(log.Topics) > 0 && log.Topics[0] == stargate.EventHashPoolCreditChainPath
	})
	if !exists {
		return nil, fmt.Errorf("unmatched pool CreditChainPath log")
	}

	creditChainPathEvent, err := w.stargatePoolFilterer.ParseCreditChainPath(poolCreditChainPathLog.Export())
	if err != nil {
		return nil, fmt.Errorf("parse CreditChainPath event: %w", err)
	}

	pool, err := stargate.NewPoolCaller(swapRemoteEvent.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("new poll caller: %w", err)
	}

	callOptions := bind.CallOpts{
		BlockNumber: task.Header.Number,
		Context:     ctx,
	}

	tokenAddress, err := pool.Token(&callOptions)
	if err != nil {
		return nil, fmt.Errorf("get token: %w", err)
	}

	sourceChain, exists := stargate.EthereumChain(creditChainPathEvent.ChainId)
	if !exists {
		return nil, fmt.Errorf("invalid chain id: %d", creditChainPathEvent.ChainId)
	}

	bridgeAction, err := w.buildTransactionBridgeAction(ctx, task.ChainID, swapRemoteEvent.Raw.Address, swapRemoteEvent.To, sourceChain, task.Network, metadata.ActionTransactionBridgeWithdraw, lo.Ternary(stargate.IsSGETH(tokenAddress), nil, &tokenAddress), swapRemoteEvent.AmountSD, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		bridgeAction,
	}

	return actions, nil
}

// buildTransactionBridgeAction builds the transaction bridge action.
func (w *worker) buildTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address, source, target network.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int, blockNumber *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := activityx.Action{
		Type:     typex.TransactionBridge,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.TransactionBridge{
			Action:        bridgeAction,
			SourceNetwork: source,
			TargetNetwork: target,
			Token:         *tokenMetadata,
		},
	}

	return &action, nil
}

// buildTransactionTransferAction builds the Ethereum transaction transfer action.
func (w *worker) buildTransactionTransferAction(ctx context.Context, task *source.Task, sender, receipt common.Address, token *common.Address, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, token, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	action := activityx.Action{
		Type:     typex.TransactionTransfer,
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}

	return &action, nil
}

// NewWorker creates a new Stargate worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize Stargate contract filterers.
	instance.stargatePoolFilterer = lo.Must(stargate.NewPoolFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
