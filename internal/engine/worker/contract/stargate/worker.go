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
	"github.com/rss3-network/node/provider/ethereum/contract/layerzero"
	"github.com/rss3-network/node/provider/ethereum/contract/stargate"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
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
	return filter.Stargate.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			// RouterETHAddresses
			stargate.AddressRouterETHMainnet,
			stargate.AddressRouterETHArbitrumOne,
			stargate.AddressRouterETHOptimism,
			stargate.AddressRouterETHBase,
			stargate.AddressRouterETHLinea,

			// RouterAddresses
			stargate.AddressRouterMainnet,
			stargate.AddressRouterBinanceSmartChain,
			stargate.AddressRouterAvalanche,
			stargate.AddressRouterPolygon,
			stargate.AddressRouterArbitrumOne,
			stargate.AddressRouterOptimism,
			stargate.AddressRouterBase,
			stargate.AddressRouterLinea,

			// RelayerAddresses
			layerzero.AddressRelayerMainnet,
			layerzero.AddressRelayerBinanceSmartChain,
			layerzero.AddressRelayerAvalanche,
			layerzero.AddressRelayerPolygon,
			layerzero.AddressRelayerArbitrum,
			layerzero.AddressRelayerOptimism,
			layerzero.AddressRelayerGnosis,
			layerzero.AddressRelayerBase,
			layerzero.AddressRelayerLinea,
		},
		LogTopics: []common.Hash{
			stargate.EventHashPoolCreditChainPath,
			stargate.EventHashPoolSwap,
			stargate.EventHashPoolSwapRemote,
		},
	}
}

func (w *worker) Match(_ context.Context, _ engine.Task) (bool, error) {
	return true, nil // TODO Remove this function.
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformStargate))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*schema.Action
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

		feed.Type = filter.TypeTransactionBridge
		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
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
func (w *worker) transformLPoolSwapLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
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

	actions := []*schema.Action{
		layerzeroFeeAction,
		stargateFeeAction,
		bridgeAction,
	}

	return actions, nil
}

// transformLPoolSwapRemoteLog transforms the pool swap remote log.
func (w *worker) transformLPoolSwapRemoteLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
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

	actions := []*schema.Action{
		bridgeAction,
	}

	return actions, nil
}

// buildTransactionBridgeAction builds the transaction bridge action.
func (w *worker) buildTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address, source, target filter.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int, blockNumber *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := schema.Action{
		Type:     filter.TypeTransactionBridge,
		Platform: filter.PlatformStargate.String(),
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
func (w *worker) buildTransactionTransferAction(ctx context.Context, task *source.Task, sender, receipt common.Address, token *common.Address, value *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, token, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	action := schema.Action{
		Type:     filter.TypeTransactionTransfer,
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

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize Stargate contract filterers.
	instance.stargatePoolFilterer = lo.Must(stargate.NewPoolFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
