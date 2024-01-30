package optimism

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/optimism"
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
	config                           *config.Module
	ethereumClient                   ethereum.Client
	tokenClient                      token.Client
	contractL1StandardBridgeFilterer *optimism.L1StandardBridgeFilterer
	contractL2StandardBridgeFilterer *optimism.L2StandardBridgeFilterer
}

func (w *worker) Name() string {
	return filter.Optimism.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			optimism.AddressL1StandardBridge,
			optimism.AddressL2StandardBridge,
		},
		LogTopics: []common.Hash{
			optimism.EventHashAddressL1StandardBridgeETHDepositInitiated,
			optimism.EventHashAddressL1StandardBridgeERC20DepositInitiated,
			optimism.EventHashAddressL1StandardBridgeETHWithdrawalFinalized,
			optimism.EventHashAddressL1StandardBridgeERC20WithdrawalFinalized,
			optimism.EventHashAddressL2StandardBridgeWithdrawalInitiated,
			optimism.EventHashAddressL2StandardBridgeDepositFinalized,
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

	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformOptimism))
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
		case w.matchL1StandardBridgeETHDepositInitiatedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeETHDepositInitiatedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeERC20DepositInitiatedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeERC20DepositInitiatedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeETHWithdrawalFinalizedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeETHWithdrawalFinalizedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeERC20WithdrawalFinalizedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeERC20WithdrawalFinalizedLog(ctx, ethereumTask, log)
		case w.matchL2StandardBridgeWithdrawalInitiatedLog(ethereumTask, log):
			actions, err = w.transformL2StandardBridgeWithdrawalInitiatedLog(ctx, ethereumTask, log)
		case w.matchL2StandardBridgeDepositFinalizedLog(ethereumTask, log):
			actions, err = w.transformL2StandardBridgeDepositFinalizedLog(ctx, ethereumTask, log)
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

func (w *worker) matchL1StandardBridgeETHDepositInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL1StandardBridge && log.Topics[0] == optimism.EventHashAddressL1StandardBridgeETHDepositInitiated
}

func (w *worker) matchL1StandardBridgeERC20DepositInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL1StandardBridge && log.Topics[0] == optimism.EventHashAddressL1StandardBridgeERC20DepositInitiated
}

func (w *worker) matchL1StandardBridgeETHWithdrawalFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL1StandardBridge && log.Topics[0] == optimism.EventHashAddressL1StandardBridgeETHWithdrawalFinalized
}

func (w *worker) matchL1StandardBridgeERC20WithdrawalFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL1StandardBridge && log.Topics[0] == optimism.EventHashAddressL1StandardBridgeERC20WithdrawalFinalized
}

func (w *worker) matchL2StandardBridgeWithdrawalInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL2StandardBridge && log.Topics[0] == optimism.EventHashAddressL2StandardBridgeWithdrawalInitiated
}

func (w *worker) matchL2StandardBridgeDepositFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == optimism.AddressL2StandardBridge && log.Topics[0] == optimism.EventHashAddressL2StandardBridgeDepositFinalized
}

func (w *worker) transformL1StandardBridgeETHDepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHDepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHDepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkEthereum, filter.NetworkOptimism, metadata.ActionTransactionBridgeDeposit, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20DepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20DepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20DepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkEthereum, filter.NetworkOptimism, metadata.ActionTransactionBridgeDeposit, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeETHWithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHWithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHWithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkOptimism, filter.NetworkEthereum, metadata.ActionTransactionBridgeWithdraw, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20WithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20WithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20WithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkOptimism, filter.NetworkEthereum, metadata.ActionTransactionBridgeWithdraw, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL2StandardBridgeWithdrawalInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseWithdrawalInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse WithdrawalInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkOptimism, filter.NetworkEthereum, metadata.ActionTransactionBridgeDeposit, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL2StandardBridgeDepositFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseDepositFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DepositFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, filter.NetworkEthereum, filter.NetworkOptimism, metadata.ActionTransactionBridgeWithdraw, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address, source, target filter.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int, blockNumber *big.Int) (*schema.Action, error) {
	// Ignore L2 ETH token address.
	if tokenAddress != nil && (*tokenAddress == optimism.AddressL1ETH || *tokenAddress == optimism.AddressL2ETH) {
		tokenAddress = nil
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := schema.Action{
		Type:     filter.TypeTransactionBridge,
		Platform: filter.PlatformOptimism.String(),
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

// NewWorker creates a new Optimism worker.
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

	// Initialize Optimism contract filterers.
	instance.contractL1StandardBridgeFilterer = lo.Must(optimism.NewL1StandardBridgeFilterer(ethereum.AddressGenesis, nil))
	instance.contractL2StandardBridgeFilterer = lo.Must(optimism.NewL2StandardBridgeFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
