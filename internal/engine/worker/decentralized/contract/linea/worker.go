package linea

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
	"github.com/rss3-network/node/provider/ethereum/contract/linea"
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
	lineaZKEVMV2Filter       *linea.ZKEVMV2Filterer
	lineaTokenBridgeFilterer *linea.TokenBridgeFilterer
	ethereumClient           ethereum.Client
	tokenClient              token.Client
}

func (w *worker) Name() string {
	return decentralized.Linea.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformLinea.String()
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

// Filter returns a filter for source.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			linea.AddressTokenBridge,
			linea.AddressZKEVMV2,
		},
		LogTopics: []common.Hash{
			linea.EventHashZKEVMV2MessageSent,
			linea.EventHashZKEVMV2MessageClaimed,
			linea.EventHashTokenBridgeBridgingInitiated,
			linea.EventHashTokenBridgeBridgingFinalized,
		},
	}
}

// Transform linea task to activity.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a linea task.
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*activityx.Action
			err     error
		)

		switch {
		case w.matchEthereumZKEVMV2MessageSentLog(*ethereumTask, log):
			actions, err = w.handleEthereumZKEVMV2MessageSentLog(ctx, *ethereumTask, *log, activity)
		case w.matchEthereumZKEVMV2MessageClaimedLog(*ethereumTask, log):
			actions, err = w.handleEthereumZKEVMV2MessageClaimedLog(ctx, *ethereumTask, *log, activity)
		case w.matchEthereumTokenBridgeBridgingInitiatedLog(*ethereumTask, log):
			actions, err = w.handleEthereumTokenBridgeBridgingInitiatedLog(ctx, *ethereumTask, *log, activity)
		case w.matchEthereumTokenBridgeBridgingFinalizedLog(*ethereumTask, log):
			actions, err = w.handleEthereumTokenBridgeBridgingFinalizedLog(ctx, *ethereumTask, *log, activity)
		default:
			zap.L().Debug("unsupported log", zap.String("task", task.ID()), zap.Uint("topic_index", log.Index))

			continue
		}

		if err != nil {
			zap.L().Warn("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))
			continue
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	activity.TotalActions = uint(len(activity.Actions))
	activity.Tag = tag.Transaction

	if len(activity.Actions) == 0 {
		return nil, nil
	}

	return activity, nil
}

func (w *worker) matchEthereumZKEVMV2MessageSentLog(task source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], linea.EventHashZKEVMV2MessageSent) {
		return false
	}

	return contract.MatchAddresses(log.Address, linea.AddressZKEVMV2) && contract.MatchAddresses(*task.Transaction.To, linea.AddressZKEVMV2)
}

func (w *worker) matchEthereumZKEVMV2MessageClaimedLog(_ source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], linea.EventHashZKEVMV2MessageClaimed) {
		return false
	}

	return contract.MatchAddresses(log.Address, linea.AddressZKEVMV2)
}

func (w *worker) matchEthereumTokenBridgeBridgingInitiatedLog(_ source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], linea.EventHashTokenBridgeBridgingInitiated) {
		return false
	}

	return contract.MatchAddresses(log.Address, linea.AddressTokenBridge)
}

func (w *worker) matchEthereumTokenBridgeBridgingFinalizedLog(_ source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], linea.EventHashTokenBridgeBridgingFinalized) {
		return false
	}

	return contract.MatchAddresses(log.Address, linea.AddressTokenBridge)
}

func (w *worker) handleEthereumZKEVMV2MessageSentLog(ctx context.Context, task source.Task, log ethereum.Log, activity *activityx.Activity) ([]*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.lineaZKEVMV2Filter.ParseMessageSent(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse MessageSent event: %w", err)
	}

	if event.Value.Sign() == 0 {
		return nil, nil
	}

	action, err := w.buildEthereumTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.Linea, metadata.ActionTransactionBridgeDeposit, nil, event.Value, log.BlockNumber)
	if err != nil {
		return nil, err
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleEthereumZKEVMV2MessageClaimedLog(ctx context.Context, task source.Task, log ethereum.Log, feed *activityx.Activity) ([]*activityx.Action, error) {
	feed.Type = typex.TransactionBridge

	if _, err := w.lineaZKEVMV2Filter.ParseMessageClaimed(log.Export()); err != nil {
		return nil, fmt.Errorf("parse MessageClaimed event: %w", err)
	}

	if !contract.MatchMethodIDs(task.Transaction.Input, linea.MethodIDZKEVMV2ClaimMessage) {
		return nil, fmt.Errorf("unmatched method id")
	}

	abi, err := linea.ZKEVMV2MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("load ZkEvmV2 ABI: %w", err)
	}

	method, err := abi.MethodById(linea.MethodIDZKEVMV2ClaimMessage[:])
	if err != nil {
		return nil, fmt.Errorf("load method by ID: %w", err)
	}

	values, err := method.Inputs.UnpackValues(task.Transaction.Input[4:])
	if err != nil {
		return nil, fmt.Errorf("unpack values: %w", err)
	}

	var input linea.ZKEVMV2ClaimMessageInput
	if err := method.Inputs.Copy(&input, values); err != nil {
		return nil, fmt.Errorf("copy input: %w", err)
	}

	if input.Value.Sign() == 0 {
		return nil, nil
	}

	action, err := w.buildEthereumTransactionBridgeAction(ctx, task.ChainID, input.From, input.To, network.Linea, network.Ethereum, metadata.ActionTransactionBridgeWithdraw, nil, input.Value, log.BlockNumber)
	if err != nil {
		return nil, err
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleEthereumTokenBridgeBridgingInitiatedLog(ctx context.Context, task source.Task, log ethereum.Log, activity *activityx.Activity) ([]*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.lineaTokenBridgeFilterer.ParseBridgingInitiatedV2(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse BridgingInitiated event: %w", err)
	}

	action, err := w.buildEthereumTransactionBridgeAction(ctx, task.ChainID, event.Sender, event.Recipient, network.Ethereum, network.Linea, metadata.ActionTransactionBridgeDeposit, &event.Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, err
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleEthereumTokenBridgeBridgingFinalizedLog(ctx context.Context, task source.Task, log ethereum.Log, activity *activityx.Activity) ([]*activityx.Action, error) {
	activity.Type = typex.TransactionBridge

	event, err := w.lineaTokenBridgeFilterer.ParseBridgingFinalizedV2(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Bridging event: %w", err)
	}

	action, err := w.buildEthereumTransactionBridgeAction(ctx, task.ChainID, task.Transaction.From, event.Recipient, network.Ethereum, network.Linea, metadata.ActionTransactionBridgeWithdraw, &event.NativeToken, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, err
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildEthereumTransactionBridgeAction(ctx context.Context, chainID uint64, sender, receiver common.Address, source, target network.Network, bridgeAction metadata.TransactionBridgeAction, tokenAddress *common.Address, tokenValue *big.Int, blockNumber *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token: %w", err)
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

	return &action, nil
}

// NewWorker returns a new linea worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		lineaZKEVMV2Filter:       lo.Must(linea.NewZKEVMV2Filterer(ethereum.AddressGenesis, nil)),
		lineaTokenBridgeFilterer: lo.Must(linea.NewTokenBridgeFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
