package base

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/base"
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

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	baseOptimismPortalFilterer   *base.OptimismPortalFilterer
	baseL1StandardBridgeFilterer *base.L1StandardBridgeFilterer
	ethereumClient               ethereum.Client
	tokenClient                  token.Client
}

func (w *worker) Name() string {
	return decentralized.Base.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformBase.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Base,
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

// Filter returns a filter for protocol.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			base.AddressOptimismPortal,
			base.AddressL1StandardBridge,
			base.AddressL2CrossDomainMessenger,
		},
		LogTopics: []common.Hash{
			base.EventHashOptimismPortalTransactionDeposited,
			base.EventHashOptimismPortalTransactionWithdrawalFinalized,
			base.EventHashL1StandardBridgeETHDepositInitiated,
			base.EventHashL1StandardBridgeERC20DepositInitiated,
			base.EventHashL1StandardBridgeETHWithdrawalFinalized,
			base.EventHashL1StandardBridgeERC20WithdrawalFinalized,
		},
	}
}

// Transform base task to activity.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a base task.
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	zap.L().Debug("transforming base task", zap.String("task_id", ethereumTask.ID()))

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			zap.L().Debug("skipping anonymous log")

			continue
		}

		var (
			action *activityx.Action
			err    error
		)

		zap.L().Debug("processing base log",
			zap.String("address", log.Address.String()),
			zap.String("topic", log.Topics[0].String()))

		switch {
		case w.matchEthereumOptimismPortalTransactionDepositedLog(log):
			zap.L().Debug("handling optimism portal transaction deposited log")

			action, err = w.handleEthereumOptimismPortalTransactionDepositedLog(ctx, *ethereumTask, activity, log)
		case w.matchEthereumL1StandardBridgeETHDepositInitiatedLog(log):
			zap.L().Debug("handling L1 standard bridge ETH deposit initiated log")

			action, err = w.handleEthereumL1StandardBridgeETHDepositInitiatedLog(ctx, *ethereumTask, activity, log)
		case w.matchEthereumL1StandardBridgeERC20DepositInitiatedLog(log):
			zap.L().Debug("handling L1 standard bridge ERC20 deposit initiated log")

			action, err = w.handleEthereumL1StandardBridgeERC20DepositInitiatedLog(ctx, *ethereumTask, activity, log)
		case w.matchEthereumL1StandardBridgeETHWithdrawalFinalizedLog(log):
			zap.L().Debug("handling L1 standard bridge ETH withdrawal finalized log")

			action, err = w.handleEthereumL1StandardBridgeETHWithdrawalFinalizedLog(ctx, *ethereumTask, activity, log)
		case w.matchEthereumL1StandardBridgeERC20WithdrawalFinalizedLog(log):
			zap.L().Debug("handling L1 standard bridge ERC20 withdrawal finalized log")

			action, err = w.handleEthereumL1StandardBridgeERC20WithdrawalFinalizedLog(ctx, *ethereumTask, activity, log)
		default:
			zap.L().Debug("skipping unsupported log")

			continue
		}

		if err != nil {
			zap.L().Error("failed to handle ethereum log",
				zap.Error(err))

			continue
		}

		activity.Actions = append(activity.Actions, action)
	}

	activity.TotalActions = uint(len(activity.Actions))
	activity.Tag = tag.Transaction

	if len(activity.Actions) == 0 {
		zap.L().Info("no actions found for task")

		return nil, nil
	}

	zap.L().Debug("successfully transformed base task")

	return activity, nil
}

func (w *worker) matchEthereumOptimismPortalTransactionDepositedLog(log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], base.EventHashOptimismPortalTransactionDeposited) {
		return false
	}

	return contract.MatchAddresses(log.Address, base.AddressOptimismPortal)
}

func (w *worker) matchEthereumL1StandardBridgeETHDepositInitiatedLog(log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], base.EventHashL1StandardBridgeETHDepositInitiated) {
		return false
	}

	return contract.MatchAddresses(log.Address, base.AddressL1StandardBridge)
}

func (w *worker) matchEthereumL1StandardBridgeERC20DepositInitiatedLog(log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], base.EventHashL1StandardBridgeERC20DepositInitiated) {
		return false
	}

	return contract.MatchAddresses(log.Address, base.AddressL1StandardBridge)
}

func (w *worker) matchEthereumL1StandardBridgeETHWithdrawalFinalizedLog(log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], base.EventHashL1StandardBridgeETHWithdrawalFinalized) {
		return false
	}

	return contract.MatchAddresses(log.Address, base.AddressL1StandardBridge)
}

func (w *worker) matchEthereumL1StandardBridgeERC20WithdrawalFinalizedLog(log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], base.EventHashL1StandardBridgeERC20WithdrawalFinalized) {
		return false
	}

	return contract.MatchAddresses(log.Address, base.AddressL1StandardBridge)
}

func (w *worker) handleEthereumOptimismPortalTransactionDepositedLog(ctx context.Context, task source.Task, activity *activityx.Activity, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.baseOptimismPortalFilterer.ParseTransactionDeposited(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TransactionDeposited event: %w", err)
	}

	if event.To == base.AddressL2CrossDomainMessenger {
		return nil, fmt.Errorf("unsupported cross domain message")
	}

	activity.Type = typex.TransactionBridge

	return w.buildEthereumTransactionBridgeAction(ctx, task.Header.Number, task.ChainID, event.From, event.To, network.Ethereum, network.Base, metadata.ActionTransactionBridgeDeposit, nil, task.Transaction.Value)
}

func (w *worker) handleEthereumL1StandardBridgeETHDepositInitiatedLog(ctx context.Context, task source.Task, activity *activityx.Activity, log *ethereum.Log) (*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.baseL1StandardBridgeFilterer.ParseETHDepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHDepositInitiated event: %w", err)
	}

	return w.buildEthereumTransactionBridgeAction(ctx, task.Header.Number, task.ChainID, event.From, event.To, network.Ethereum, network.Base, metadata.ActionTransactionBridgeDeposit, nil, event.Amount)
}

func (w *worker) handleEthereumL1StandardBridgeERC20DepositInitiatedLog(ctx context.Context, task source.Task, activity *activityx.Activity, log *ethereum.Log) (*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.baseL1StandardBridgeFilterer.ParseERC20DepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20DepositInitiated event: %w", err)
	}

	return w.buildEthereumTransactionBridgeAction(ctx, task.Header.Number, task.ChainID, event.From, event.To, network.Ethereum, network.Base, metadata.ActionTransactionBridgeDeposit, &event.L1Token, event.Amount)
}

func (w *worker) handleEthereumL1StandardBridgeETHWithdrawalFinalizedLog(ctx context.Context, task source.Task, activity *activityx.Activity, log *ethereum.Log) (*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.baseL1StandardBridgeFilterer.ParseETHWithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ETHWithdrawalFinalized event: %w", err)
	}

	return w.buildEthereumTransactionBridgeAction(ctx, task.Header.Number, task.ChainID, event.From, event.To, network.Base, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, nil, event.Amount)
}
func (w *worker) handleEthereumL1StandardBridgeERC20WithdrawalFinalizedLog(ctx context.Context, task source.Task, activity *activityx.Activity, log *ethereum.Log) (*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.baseL1StandardBridgeFilterer.ParseERC20WithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC20WithdrawalFinalized event: %w", err)
	}

	return w.buildEthereumTransactionBridgeAction(ctx, task.Header.Number, task.ChainID, event.From, event.To, network.Base, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &event.L1Token, event.Amount)
}

func (w *worker) buildEthereumTransactionBridgeAction(ctx context.Context, blockNumber *big.Int, chainID uint64, sender, receiver common.Address, source, target network.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int) (*activityx.Action, error) {
	zap.L().Debug("building ethereum transaction bridge action",
		zap.String("sender", sender.String()),
		zap.String("receiver", receiver.String()),
		zap.String("source_network", source.String()),
		zap.String("target_network", target.String()),
		zap.String("action", bridgeAction.String()),
		zap.Any("token_address", tokenAddress),
		zap.Any("token_value", tokenValue))

	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := activityx.Action{
		Type:     typex.TransactionBridge,
		Tag:      tag.Transaction,
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

	zap.L().Debug("successfully built ethereum transaction bridge action")

	return &action, nil
}

// NewWorker returns a new base worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		baseOptimismPortalFilterer:   lo.Must(base.NewOptimismPortalFilterer(ethereum.AddressGenesis, nil)),
		baseL1StandardBridgeFilterer: lo.Must(base.NewL1StandardBridgeFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
