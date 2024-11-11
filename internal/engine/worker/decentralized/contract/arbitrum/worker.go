package arbitrum

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
	"github.com/rss3-network/node/provider/ethereum/contract/arbitrum"
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
	config                                 *config.Module
	ethereumClient                         ethereum.Client
	tokenClient                            token.Client
	contractBridgeFilterer                 *arbitrum.BridgeFilterer
	contractInboxFilterer                  *arbitrum.InboxFilterer
	contractL1CustomGatewayFilterer        *arbitrum.L1CustomGatewayFilterer
	contractL2ReverseCustomGatewayFilterer *arbitrum.L2ReverseCustomGatewayFilterer
	contractArbSysFilterer                 *arbitrum.ArbSysFilterer
}

func (w *worker) Name() string {
	return decentralized.Arbitrum.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformArbitrum.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Arbitrum,
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
			arbitrum.AddressBridgeOne,
			arbitrum.AddressInboxOne,
			arbitrum.AddressL1CustomGatewayOne,
			arbitrum.AddressL2ReverseCustomGateway,
			arbitrum.AddressArbSys,
		},
		LogTopics: []common.Hash{
			arbitrum.EventHashBridgeMessageDelivered,
			arbitrum.EventHashL1CustomGatewayDepositInitiated,
			arbitrum.EventHashL2ReverseCustomGatewayWithdrawalInitiated,
			arbitrum.EventHashArbSysL2ToL1Tx,
			arbitrum.EventHashL1CustomGatewayWithdrawalFinalized,
			arbitrum.EventHashL2ReverseCustomGatewayDepositFinalized,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	zap.L().Debug("transforming arbitrum task",
		zap.String("task_id", ethereumTask.ID()))

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
			zap.L().Debug("skipping anonymous log",
				zap.String("task_id", ethereumTask.ID()),
				zap.Uint("log_index", log.Index))

			continue
		}

		switch {
		case w.matchBridgeMessageDeliveredLog(ethereumTask, log):
			zap.L().Debug("handling bridge message delivered log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformBridgeMessageDeliveredLog(ctx, ethereumTask, log)
		case w.matchL1CustomGatewayDepositInitiatedLog(ethereumTask, log):
			zap.L().Debug("handling L1 custom gateway deposit initiated log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformL1CustomGatewayDepositInitiatedLog(ctx, ethereumTask, log)
		case w.matchL2ReverseCustomGatewayWithdrawalInitiatedLog(ethereumTask, log):
			zap.L().Debug("handling L2 reverse custom gateway withdrawal initiated log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformL2ReverseCustomGatewayWithdrawalInitiatedLog(ctx, ethereumTask, log)
		case w.matchArbSysL2ToL1TxLog(ethereumTask, log):
			zap.L().Debug("handling ArbSys L2 to L1 tx log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformArbSysL2ToL1TxLog(ctx, ethereumTask, log)
		case w.matchL1CustomGatewayWithdrawalFinalizedLog(ethereumTask, log):
			zap.L().Debug("handling L1 custom gateway withdrawal finalized log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformL1CustomGatewayWithdrawalFinalizedLog(ctx, ethereumTask, log)
		case w.matchL2ReverseCustomGatewayDepositFinalizedLog(ethereumTask, log):
			zap.L().Debug("handling L2 reverse custom gateway deposit finalized log",
				zap.String("task_id", ethereumTask.ID()))

			actions, err = w.transformL2ReverseCustomGatewayDepositFinalizedLog(ctx, ethereumTask, log)
		default:
			zap.L().Warn("unsupported log",
				zap.String("task_id", ethereumTask.ID()))

			continue
		}

		if err != nil {
			zap.L().Error("failed to handle ethereum log",
				zap.Error(err),
				zap.String("task_id", ethereumTask.ID()))

			continue
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	activity.Type = typex.TransactionBridge

	zap.L().Debug("successfully transformed arbitrum task",
		zap.String("task_id", ethereumTask.ID()))

	return activity, nil
}

func (w *worker) matchBridgeMessageDeliveredLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashBridgeMessageDelivered) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressBridgeOne)
}

func (w *worker) matchL1CustomGatewayDepositInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashL1CustomGatewayDepositInitiated) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressL1CustomGatewayOne)
}

func (w *worker) matchL1CustomGatewayWithdrawalFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashL1CustomGatewayWithdrawalFinalized) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressL1CustomGatewayOne)
}

func (w *worker) matchL2ReverseCustomGatewayWithdrawalInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashL2ReverseCustomGatewayWithdrawalInitiated) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressL2ReverseCustomGateway)
}

func (w *worker) matchL2ReverseCustomGatewayDepositFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashL2ReverseCustomGatewayDepositFinalized) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressL2ReverseCustomGateway)
}

func (w *worker) matchArbSysL2ToL1TxLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], arbitrum.EventHashArbSysL2ToL1Tx) {
		return false
	}

	return contract.MatchAddresses(log.Address, arbitrum.AddressArbSys)
}

func (w *worker) transformBridgeMessageDeliveredLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractBridgeFilterer.ParseMessageDelivered(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse MessageDelivered event: %w", err)
	}

	if event.Kind != arbitrum.L1MessageTypeETHDeposit && event.Kind != arbitrum.L1MessageTypeTransactCall {
		return nil, fmt.Errorf("unspoorted message kind: %d", event.Kind)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.Sender, event.Sender, network.Ethereum, network.Arbitrum, metadata.ActionTransactionBridgeDeposit, nil, task.Transaction.Value, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformL1CustomGatewayDepositInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1CustomGatewayFilterer.ParseDepositInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DepositInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.Arbitrum, metadata.ActionTransactionBridgeDeposit, &event.L1Token, task.Transaction.Value, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformL1CustomGatewayWithdrawalFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL1CustomGatewayFilterer.ParseWithdrawalFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse WithdrawalFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Arbitrum, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformL2ReverseCustomGatewayWithdrawalInitiatedLog(ctx context.Context, task *source.Task,
	log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL2ReverseCustomGatewayFilterer.ParseWithdrawalInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse WithdrawalInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Arbitrum,
		network.Ethereum, metadata.ActionTransactionBridgeWithdraw, &arbitrum.AddressARB, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformL2ReverseCustomGatewayDepositFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL2ReverseCustomGatewayFilterer.ParseDepositFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DepositFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.Arbitrum, metadata.ActionTransactionBridgeDeposit, &event.L1Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformArbSysL2ToL1TxLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractArbSysFilterer.ParseL2ToL1Tx(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse L2ToL1Tx event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.Caller, event.Destination, network.Arbitrum, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, nil, event.Callvalue, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) buildTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address,
	source, target network.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address,
	tokenValue *big.Int, blockNumber *big.Int) (*activityx.Action, error) {
	zap.L().Debug("building transaction bridge action",
		zap.String("sender", sender.String()),
		zap.String("receiver", receiver.String()),
		zap.String("source_network", source.String()),
		zap.String("target_network", target.String()),
		zap.String("action", bridgeAction.String()),
		zap.String("token_address", tokenAddress.String()),
		zap.String("token_value", tokenValue.String()))

	// If the chain is 'Arbitrum', then set blockNumber to be nil by default to use Lookup()
	if source == network.Arbitrum {
		blockNumber = nil
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := &activityx.Action{
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

	zap.L().Debug("successfully built transaction bridge action")

	return action, nil
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	zap.L().Debug("initializing arbitrum worker",
		zap.String("ID", config.ID),
		zap.String("network", config.Network.String()),
		zap.String("worker", config.Worker.Name()),
		zap.String("endpoint", config.Endpoint.URL),
		zap.Any("params", config.Parameters))

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

	instance.contractBridgeFilterer = lo.Must(arbitrum.NewBridgeFilterer(ethereum.AddressGenesis, nil))
	instance.contractInboxFilterer = lo.Must(arbitrum.NewInboxFilterer(ethereum.AddressGenesis, nil))
	instance.contractL1CustomGatewayFilterer = lo.Must(arbitrum.NewL1CustomGatewayFilterer(ethereum.AddressGenesis, nil))
	instance.contractL2ReverseCustomGatewayFilterer = lo.Must(arbitrum.NewL2ReverseCustomGatewayFilterer(ethereum.AddressGenesis, nil))
	instance.contractArbSysFilterer = lo.Must(arbitrum.NewArbSysFilterer(ethereum.AddressGenesis, nil))
	instance.contractL1CustomGatewayFilterer = lo.Must(arbitrum.NewL1CustomGatewayFilterer(ethereum.AddressGenesis, nil))
	instance.contractL2ReverseCustomGatewayFilterer = lo.Must(arbitrum.NewL2ReverseCustomGatewayFilterer(ethereum.AddressGenesis, nil))

	zap.L().Debug("arbitrum worker initialized successfully")

	return &instance, nil
}
