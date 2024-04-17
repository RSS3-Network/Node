package savm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/savm"
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
	config                     *config.Module
	ethereumClient             ethereum.Client
	tokenClient                token.Client
	contractSAVMBridgeFilterer *savm.SAVMBridgeFilterer
	contractBTCBridgeFilterer  *savm.BTCBridgeFilterer
	erc20Filterer              *erc20.ERC20Filterer
}

func (w *worker) Name() string {
	return filter.SAVM.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			savm.AddressSAVMBridge,
			savm.AddressBTCBridge,
		},
		LogTopics: []common.Hash{
			savm.EventHashBTCBridgeDeposit,
			savm.EventHashBTCBridgeWithdraw,
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

	feed, err := ethereumTask.BuildActivity(activity.WithActivityPlatform(filter.PlatformSAVM))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
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
		case w.matchBTCBridgeWithdrawLog(ethereumTask, log):
			actions, err = w.transformBTCBridgeWithdrawLog(ctx, ethereumTask, log)
		case w.matchBTCBridgeDepositLog(ethereumTask, log):
			actions, err = w.transformBTCBridgeDepositLog(ctx, ethereumTask, log)
		case w.matchSAVMTransferLog(ethereumTask, log):
			actions, err = w.transformSAVMBridgeLog(ctx, ethereumTask, log)
		default:
			zap.L().Warn("unsupported log", zap.String("task", task.ID()), zap.Uint("topic.index", log.Index))

			continue
		}

		if err != nil {
			zap.L().Warn("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			return nil, err
		}

		feed.Type = typex.TransactionBridge
		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

func (w *worker) matchBTCBridgeWithdrawLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == savm.AddressBTCBridge && log.Topics[0] == savm.EventHashBTCBridgeWithdraw
}

func (w *worker) matchBTCBridgeDepositLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == savm.AddressBTCBridge && log.Topics[0] == savm.EventHashBTCBridgeDeposit
}

func (w *worker) matchSAVMTransferLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == savm.AddressSAVMToken && log.Topics[0] == savm.EventHashSAVMTransfer
}

func (w *worker) transformBTCBridgeWithdrawLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractBTCBridgeFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Withdraw event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.Sender, event.Sender, network.SatoshiVM, network.Bitcoin, metadata.ActionTransactionBridgeWithdraw, nil, event.NormalizedAmount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformBTCBridgeDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.contractBTCBridgeFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Deposit event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.Recipient, event.Recipient, network.Bitcoin, network.SatoshiVM, metadata.ActionTransactionBridgeDeposit, nil, event.NormalizedAmount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activity.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformSAVMBridgeLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activity.Action, error) {
	event, err := w.erc20Filterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Withdraw event: %w", err)
	}

	var actions []*activity.Action

	switch event.To {
	case savm.AddressSAVMBridge:
		action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.From, network.SatoshiVM, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &savm.AddressSAVMToken, event.Value, log.BlockNumber)
		if err != nil {
			return nil, fmt.Errorf("build transaction bridge action: %w", err)
		}

		actions = append(actions, action)
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
		Platform: filter.PlatformSAVM.String(),
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

	// Initialize SAVM contract filterers.
	instance.contractSAVMBridgeFilterer = lo.Must(savm.NewSAVMBridgeFilterer(ethereum.AddressGenesis, nil))
	instance.contractBTCBridgeFilterer = lo.Must(savm.NewBTCBridgeFilterer(ethereum.AddressGenesis, nil))
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
