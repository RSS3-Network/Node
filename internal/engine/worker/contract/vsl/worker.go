package vsl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                           *config.Module
	ethereumClient                   ethereum.Client
	tokenClient                      token.Client
	contractL1StandardBridgeFilterer *vsl.L1StandardBridgeFilterer
	contractL2StandardBridgeFilterer *vsl.L2StandardBridgeFilterer
}

func (w *worker) Name() string {
	return filter.VSL.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			vsl.AddressL1StandardBridge,
			vsl.AddressL2StandardBridge,
		},
		LogTopics: []common.Hash{
			vsl.EventHashAddressL1StandardBridgeETHDepositInitiated,
			vsl.EventHashAddressL1StandardBridgeERC20DepositInitiated,
			vsl.EventHashAddressL1StandardBridgeETHWithdrawalFinalized,
			vsl.EventHashAddressL1StandardBridgeERC20WithdrawalFinalized,
			vsl.EventHashAddressL2StandardBridgeWithdrawalInitiated,
			vsl.EventHashAddressL2StandardBridgeDepositFinalized,
		},
	}
}

func (w *worker) Match(_ context.Context, _ engine.Task) (bool, error) {
	return true, nil // TODO Remove this function.
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activity.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	_activity, err := ethereumTask.BuildActivity(activity.WithActivityPlatform(filter.PlatformVSL))
	if err != nil {
		return nil, fmt.Errorf("build _activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*activity.Action
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

		_activity.Type = typex.TransactionBridge
		_activity.Actions = append(_activity.Actions, actions...)
	}

	return _activity, nil
}

func (w *worker) matchL1StandardBridgeETHDepositInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL1StandardBridge && log.Topics[0] == vsl.EventHashAddressL1StandardBridgeETHDepositInitiated
}

func (w *worker) matchL1StandardBridgeERC20DepositInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL1StandardBridge && log.Topics[0] == vsl.EventHashAddressL1StandardBridgeERC20DepositInitiated
}

func (w *worker) matchL1StandardBridgeETHWithdrawalFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL1StandardBridge && log.Topics[0] == vsl.EventHashAddressL1StandardBridgeETHWithdrawalFinalized
}

func (w *worker) matchL1StandardBridgeERC20WithdrawalFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL1StandardBridge && log.Topics[0] == vsl.EventHashAddressL1StandardBridgeERC20WithdrawalFinalized
}

func (w *worker) matchL2StandardBridgeWithdrawalInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL2StandardBridge && log.Topics[0] == vsl.EventHashAddressL2StandardBridgeWithdrawalInitiated
}

func (w *worker) matchL2StandardBridgeDepositFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL2StandardBridge && log.Topics[0] == vsl.EventHashAddressL2StandardBridgeDepositFinalized
}

func (w *worker) transformL1StandardBridgeETHDepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHDepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHDepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeDeposit, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20DepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20DepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20DepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeDeposit, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeETHWithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHWithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHWithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20WithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20WithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20WithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL2StandardBridgeWithdrawalInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseWithdrawalInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse WithdrawalInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeDeposit, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL2StandardBridgeDepositFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseDepositFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DepositFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeWithdraw, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address, source, target network.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int, blockNumber *big.Int) (*activity.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := activity.Action{
		Type:     typex.TransactionBridge,
		Platform: filter.PlatformVSL.String(),
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

// NewWorker creates a new VSL worker.
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

	// Initialize VSL contract filterers.
	instance.contractL1StandardBridgeFilterer = lo.Must(vsl.NewL1StandardBridgeFilterer(ethereum.AddressGenesis, nil))
	instance.contractL2StandardBridgeFilterer = lo.Must(vsl.NewL2StandardBridgeFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
