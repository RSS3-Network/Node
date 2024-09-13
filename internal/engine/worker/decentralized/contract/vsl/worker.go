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
	config                           *config.Module
	ethereumClient                   ethereum.Client
	tokenClient                      token.Client
	contractL1StandardBridgeFilterer *vsl.L1StandardBridgeFilterer
}

func (w *worker) Name() string {
	return decentralized.VSL.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformVSL.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
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
	}
}

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			vsl.AddressL1StandardBridge,
		},
		LogTopics: []common.Hash{
			vsl.EventHashAddressL1StandardBridgeETHDepositInitiated,
			vsl.EventHashAddressL1StandardBridgeERC20DepositInitiated,
			vsl.EventHashAddressL1StandardBridgeETHWithdrawalFinalized,
			vsl.EventHashAddressL1StandardBridgeERC20WithdrawalFinalized,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
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
		case w.matchL1StandardBridgeETHDepositInitiatedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeETHDepositInitiatedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeERC20DepositInitiatedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeERC20DepositInitiatedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeETHWithdrawalFinalizedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeETHWithdrawalFinalizedLog(ctx, ethereumTask, log)
		case w.matchL1StandardBridgeERC20WithdrawalFinalizedLog(ethereumTask, log):
			actions, err = w.transformL1StandardBridgeERC20WithdrawalFinalizedLog(ctx, ethereumTask, log)
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

func (w *worker) transformL1StandardBridgeETHDepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHDepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHDepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeDeposit, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20DepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20DepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20DepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeDeposit, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeETHWithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseETHWithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHWithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, nil, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL1StandardBridgeERC20WithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1StandardBridgeFilterer.ParseERC20WithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20WithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

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

// NewWorker creates a new VSL worker.
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

	// Initialize VSL contract filterers.
	instance.contractL1StandardBridgeFilterer = lo.Must(vsl.NewL1StandardBridgeFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
