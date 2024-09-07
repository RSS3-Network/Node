package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/erc1155"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/contract/rss3"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/node/provider/ethereum/token"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"github.com/sourcegraph/conc/pool"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                           *config.Module
	ethereumClient                   ethereum.Client
	tokenClient                      token.Client
	erc20Filterer                    *erc20.ERC20Filterer
	erc721Filterer                   *erc721.ERC721Filterer
	erc1155Filterer                  *erc1155.ERC1155Filterer
	stakingVSLFilterer               *rss3.StakingVSLFilterer
	chipsFilterer                    *rss3.ChipsFilterer
	contractL2StandardBridgeFilterer *vsl.L2StandardBridgeFilterer
}

func (w *worker) Name() string {
	return workerx.Core.String()
}

func (w *worker) Platform() string {
	return ""
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Arbitrum,
		network.Avalanche,
		network.Base,
		network.BinanceSmartChain,
		network.Crossbell,
		network.Ethereum,
		network.Gnosis,
		network.Linea,
		network.Optimism,
		network.Polygon,
		network.SatoshiVM,
		network.VSL,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Unknown,
		tag.Collectible,
		tag.Transaction,
		tag.Exchange,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.TransactionTransfer,
		typex.TransactionApproval,
		typex.TransactionMint,
		typex.TransactionBurn,
		typex.TransactionBridge,
		typex.CollectibleTransfer,
		typex.CollectibleApproval,
		typex.CollectibleMint,
		typex.TransactionBurn,
		typex.ExchangeStaking,
	}
}

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return nil
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	tracerOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(engine.BuildTaskTraceAttributes(task)...),
	}

	ctx, span := otel.Tracer("").Start(ctx, "Worker transform", tracerOptions...)
	defer span.End()

	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := task.BuildActivity()
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	actions := make([][]*activityx.Action, len(ethereumTask.Receipt.Logs)+1)

	// If the transaction is failed, we will not process it.
	if w.matchFailedTransaction(ethereumTask) {
		return activity, nil
	}

	if w.matchNativeTransferTransaction(ethereumTask) {
		action, err := w.handleNativeTransferTransaction(ctx, ethereumTask)
		if err != nil {
			return nil, fmt.Errorf("handle native transfer transaction: %w", err)
		}

		activity.Type = action.Type
		activity.Actions = append(activity.Actions, action)
	}

	contextPool := pool.New().
		WithContext(ctx).
		WithFirstError().
		WithCancelOnError()

	// Handle ERC-20 token transfer logs.
	for index, log := range ethereumTask.Receipt.Logs {
		index, log := index, log

		// Ignore the log if it is removed or is anonymous event.
		if log.Removed || len(log.Topics) == 0 {
			continue
		}

		contextPool.Go(func(ctx context.Context) error {
			var (
				logActions []*activityx.Action
				err        error
			)

			switch {
			// VSL Bridge
			case w.matchL2StandardBridgeWithdrawalInitiatedLog(ethereumTask, log):
				logActions, err = w.transformL2StandardBridgeWithdrawalInitiatedLog(ctx, ethereumTask, log)
			case w.matchL2StandardBridgeDepositFinalizedLog(ethereumTask, log):
				logActions, err = w.transformL2StandardBridgeDepositFinalizedLog(ctx, ethereumTask, log)
			// VSL Staking
			case w.matchStakingVSLDeposited(ethereumTask, log):
				logActions, err = w.handleStakingVSLDeposited(ctx, ethereumTask, log)
			case w.matchStakingVSLStaked(ethereumTask, log):
				logActions, err = w.handleStakingVSLStaked(ctx, ethereumTask, log)
			case w.matchChipsTransfer(ethereumTask, log):
				logActions, err = w.handleChipsMint(ctx, ethereumTask, log)
			case w.matchERC20TransferLog(ethereumTask, log):
				logActions, err = w.handleERC20TransferLog(ctx, ethereumTask, log)
			case w.matchERC20ApprovalLog(ethereumTask, log):
				logActions, err = w.handleERC20ApproveLog(ctx, ethereumTask, log)
			case w.matchERC721TransferLog(ethereumTask, log):
				logActions, err = w.handleERC721TransferLog(ctx, ethereumTask, log)
			case w.matchERC721ApprovalLog(ethereumTask, log):
				logActions, err = w.handleERC721ApproveLog(ctx, ethereumTask, log)
			case w.matchERC1155TransferLog(ethereumTask, log):
				logActions, err = w.handleERC1155TransferLog(ctx, ethereumTask, log)
			case w.matchERC1155ApprovalLog(ethereumTask, log):
				logActions, err = w.handleERC1155ApproveLog(ctx, ethereumTask, log)
			}

			if err != nil {
				return err
			}

			actions[index] = logActions

			return nil
		})
	}

	if err := contextPool.Wait(); err != nil {
		if isInvalidTokenErr(err) {
			return activityx.NewUnknownActivity(activity), nil
		}

		return nil, fmt.Errorf("handle log: %w", err)
	}

	activity.Actions = append(activity.Actions, lo.Flatten(actions)...)

	for _, action := range activity.Actions {
		activity.Type = action.Type
	}

	return activity, nil
}

func (w *worker) matchFailedTransaction(task *source.Task) bool {
	return task.Receipt.Status == types.ReceiptStatusFailed
}

func (w *worker) matchNativeTransferTransaction(task *source.Task) bool {
	return task.Transaction.Value.Cmp(big.NewInt(0)) == 1 && len(task.Transaction.Input) == 0
}

func (w *worker) matchERC20TransferLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer)
}

func (w *worker) matchERC20ApprovalLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc20.EventHashApproval)
}

func (w *worker) matchERC721TransferLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc721.EventHashTransfer)
}

func (w *worker) matchERC721ApprovalLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc721.EventHashApproval) ||
		len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc721.EventHashApprovalForAll)
}

func (w *worker) matchERC1155TransferLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc1155.EventHashTransferSingle, erc1155.EventHashTransferBatch)
}

func (w *worker) matchERC1155ApprovalLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc1155.EventHashApprovalForAll)
}

// matchStakingVSLDeposited matches the staking VSL deposited event.
func (w *worker) matchStakingVSLDeposited(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressStakingVSL && contract.MatchEventHashes(log.Topics[0], rss3.EventHashStakingVSLDeposited)
}

// matchStakingVSLStaked matches the staking VSL staked event.
func (w *worker) matchStakingVSLStaked(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressStakingVSL && contract.MatchEventHashes(log.Topics[0], rss3.EventHashStakingVSLStaked)
}

// matchTransfer matches the transfer event.
func (w *worker) matchChipsTransfer(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == rss3.AddressChipsVSL && contract.MatchEventHashes(log.Topics[0], rss3.EventHashTransfer)
}

func (w *worker) matchL2StandardBridgeWithdrawalInitiatedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL2StandardBridge && log.Topics[0] == vsl.EventHashAddressL2StandardBridgeWithdrawalInitiated
}

func (w *worker) matchL2StandardBridgeDepositFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == vsl.AddressL2StandardBridge && log.Topics[0] == vsl.EventHashAddressL2StandardBridgeDepositFinalized
}

func (w *worker) handleNativeTransferTransaction(ctx context.Context, task *source.Task) (*activityx.Action, error) {
	transactionFrom := task.Transaction.From

	// If the transaction is a contract creation,
	// we will set the from address to the zero address.
	var transactionTo = ethereum.AddressGenesis
	if task.Transaction.To != nil {
		transactionTo = *task.Transaction.To
	}

	action, err := w.buildTransactionTransferAction(ctx, task, transactionFrom, transactionTo, nil, task.Transaction.Value)
	if err != nil {
		return nil, fmt.Errorf("build native transfer action: %w", err)
	}

	return action, nil
}

func (w *worker) handleERC20TransferLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.erc20Filterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC-20 Transfer event: %w", err)
	}

	action, err := w.buildTransactionTransferAction(ctx, task, event.From, event.To, &event.Raw.Address, event.Value)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleERC20ApproveLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.erc20Filterer.ParseApproval(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC-20 Approval event: %w", err)
	}

	action, err := w.buildTransactionApprovalAction(ctx, task, event.Owner, event.Spender, &event.Raw.Address, event.Value)
	if err != nil {
		return nil, fmt.Errorf("build transaction approval action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleERC721TransferLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.erc721Filterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC-721 Transfer event: %w", err)
	}

	action, err := w.buildCollectibleTransferAction(ctx, task, event.From, event.To, event.Raw.Address, event.TokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleERC721ApproveLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	var action *activityx.Action

	switch {
	case len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc721.EventHashApproval):
		event, err := w.erc721Filterer.ParseApproval(log.Export())
		if err != nil {
			return nil, fmt.Errorf("parse ERC-721 Approval event: %w", err)
		}

		if action, err = w.buildCollectibleApprovalAction(ctx, task, event.Owner, event.Approved, event.Raw.Address, event.TokenId, nil); err != nil {
			return nil, fmt.Errorf("build collectible approval action: %w", err)
		}
	case len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc721.EventHashApprovalForAll):
		event, err := w.erc721Filterer.ParseApprovalForAll(log.Export())
		if err != nil {
			return nil, fmt.Errorf("parse ERC-721 ApprovalForAll event: %w", err)
		}

		if action, err = w.buildCollectibleApprovalAction(ctx, task, event.Owner, event.Operator, event.Raw.Address, nil, &event.Approved); err != nil {
			return nil, fmt.Errorf("build collectible approval action: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported ERC-721 approval event %s", log.Topics[0])
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) handleERC1155TransferLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	actions := make([]*activityx.Action, 0)

	switch {
	case len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc1155.EventHashTransferSingle):
		event, err := w.erc1155Filterer.ParseTransferSingle(log.Export())
		if err != nil {
			return nil, fmt.Errorf("parse ERC-1577 TransferSingle event: %w", err)
		}

		action, err := w.buildCollectibleTransferAction(ctx, task, event.From, event.To, event.Raw.Address, event.Id, event.Value)
		if err != nil {
			return nil, fmt.Errorf("build collectible transfer action: %w", err)
		}

		actions = append(actions, action)
	case len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], erc1155.EventHashTransferBatch):
		event, err := w.erc1155Filterer.ParseTransferBatch(log.Export())
		if err != nil {
			return nil, fmt.Errorf("parse ERC-1155 TransferBatch event: %w", err)
		}

		if len(event.Ids) != len(event.Values) || len(event.Ids) == 0 {
			return nil, fmt.Errorf("invalid ERC-1155 TransferBatch event %s", log.Topics[0])
		}

		for i := range event.Ids {
			id, value := event.Ids[i], event.Values[i]

			action, err := w.buildCollectibleTransferAction(ctx, task, event.From, event.To, event.Raw.Address, id, value)
			if err != nil {
				return nil, fmt.Errorf("build collectible transfer action: %w", err)
			}

			actions = append(actions, action)
		}
	default:
		return nil, fmt.Errorf("invalid ERC-1155 transfer event %s", log.Topics[0])
	}

	return actions, nil
}

func (w *worker) handleERC1155ApproveLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.erc1155Filterer.ParseApprovalForAll(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse ERC-1155 ApprovalForAll event: %w", err)
	}

	action, err := w.buildCollectibleApprovalAction(ctx, task, event.Account, event.Operator, event.Raw.Address, nil, &event.Approved)
	if err != nil {
		return nil, fmt.Errorf("build collectible approval action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

// handleStakingVSLDeposited  handles the staking VSL deposited event.
func (w *worker) handleStakingVSLDeposited(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.stakingVSLFilterer.ParseDeposited(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Deposited event: %w", err)
	}

	action, err := w.buildExchangeStakingVSLAction(ctx, task, task.Transaction.From, event.NodeAddr, event.Amount, metadata.ActionExchangeStakingStake)
	if err != nil {
		return nil, fmt.Errorf("build exchange staking action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

// handleStakingVSLStaked  handles the staking VSL staked event.
func (w *worker) handleStakingVSLStaked(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.stakingVSLFilterer.ParseStaked(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Deposited event: %w", err)
	}

	stakingAction, err := w.buildExchangeStakingVSLAction(ctx, task, task.Transaction.From, event.NodeAddr, event.Amount, metadata.ActionExchangeStakingStake)
	if err != nil {
		return nil, fmt.Errorf("build exchange staking action: %w", err)
	}

	actions := []*activityx.Action{
		stakingAction,
	}

	return actions, nil
}

// handleChipsMint  handles the transfer event.
func (w *worker) handleChipsMint(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse Transfer event.
	event, err := w.chipsFilterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	action, err := w.buildChipsMintAction(ctx, task, event.From, event.To, log.Address, event.TokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build ChipsMint action: %w", err)
	}

	return []*activityx.Action{
		action,
	}, nil
}

func (w *worker) transformL2StandardBridgeWithdrawalInitiatedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseWithdrawalInitiated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse WithdrawalInitiated event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.VSL, network.Ethereum, metadata.ActionTransactionBridgeDeposit, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformL2StandardBridgeDepositFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.contractL2StandardBridgeFilterer.ParseDepositFinalized(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DepositFinalized event: %w", err)
	}

	action, err := w.buildTransactionBridgeAction(ctx, task.ChainID, event.From, event.To, network.Ethereum, network.VSL, metadata.ActionTransactionBridgeWithdraw, &event.L2Token, event.Amount, log.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("build transaction bridge action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildTransactionTransferAction(ctx context.Context, task *source.Task, from, to common.Address, tokenAddress *common.Address, tokenValue *big.Int) (*activityx.Action, error) {
	chainID, err := network.EthereumChainIDString(task.GetNetwork().String())
	if err != nil {
		return nil, fmt.Errorf("invalid chain id: %w", err)
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, uint64(chainID), tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	var actionType typex.TransactionType

	switch {
	case ethereum.IsBurnAddress(to):
		actionType = typex.TransactionBurn
	case ethereum.AddressGenesis == from:
		actionType = typex.TransactionMint
	default:
		actionType = typex.TransactionTransfer
	}

	action := activityx.Action{
		Type:     actionType,
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}

	return &action, nil
}

func (w *worker) buildTransactionApprovalAction(ctx context.Context, task *source.Task, from, to common.Address, tokenAddress *common.Address, tokenValue *big.Int) (*activityx.Action, error) {
	chainID, err := network.EthereumChainIDString(task.GetNetwork().String())
	if err != nil {
		return nil, fmt.Errorf("invalid chain id: %w", err)
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, uint64(chainID), tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	// Use the token value to determine the action type.
	metadataAction := metadata.ActionTransactionApprove
	if tokenValue.Cmp(big.NewInt(0)) != 1 {
		metadataAction = metadata.ActionTransactionRevoke
	}

	action := activityx.Action{
		Type: typex.TransactionApproval,
		From: from.String(),
		To:   to.String(),
		Metadata: metadata.TransactionApproval{
			Action: metadataAction,
			Token:  *tokenMetadata,
		},
	}

	return &action, nil
}

func (w *worker) buildCollectibleTransferAction(ctx context.Context, task *source.Task, from, to common.Address, tokenAddress common.Address, tokenID *big.Int, tokenValue *big.Int) (*activityx.Action, error) {
	chainID, err := network.EthereumChainIDString(task.GetNetwork().String())
	if err != nil {
		return nil, fmt.Errorf("invalid chain id: %w", err)
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, uint64(chainID), &tokenAddress, tokenID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	var actionType typex.CollectibleType

	switch {
	case ethereum.IsBurnAddress(to):
		actionType = typex.CollectibleBurn
	case ethereum.AddressGenesis == from:
		actionType = typex.CollectibleMint
	default:
		actionType = typex.CollectibleTransfer
	}

	action := activityx.Action{
		Type:     actionType,
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}

	return &action, nil
}

func (w *worker) buildCollectibleApprovalAction(ctx context.Context, task *source.Task, from common.Address, to common.Address, tokenAddress common.Address, id *big.Int, approved *bool) (*activityx.Action, error) {
	chainID, err := network.EthereumChainIDString(task.GetNetwork().String())
	if err != nil {
		return nil, fmt.Errorf("invalid chain id: %w", err)
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, uint64(chainID), &tokenAddress, id, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token %s: %w", tokenAddress, err)
	}

	var metadataAction metadata.CollectibleApprovalAction

	switch {
	case approved == nil && to == ethereum.AddressGenesis:
		metadataAction = metadata.ActionCollectibleApprovalRevoke
	case approved == nil:
		metadataAction = metadata.ActionCollectibleApprovalApprove
	case *approved:
		metadataAction = metadata.ActionCollectibleApprovalApprove
	default:
		metadataAction = metadata.ActionCollectibleApprovalRevoke
	}

	action := activityx.Action{
		Type: typex.CollectibleApproval,
		From: from.String(),
		To:   to.String(),
		Metadata: metadata.CollectibleApproval{
			Action: metadataAction,
			Token:  *tokenMetadata,
		},
	}

	return &action, nil
}

// buildExchangeStakingVSLAction builds the exchange staking VSL action.
func (w *worker) buildExchangeStakingVSLAction(ctx context.Context, task *source.Task, from, to common.Address, tokenValue *big.Int, stakingAction metadata.ExchangeStakingAction) (*activityx.Action, error) {
	// The Token always is $RSS3.
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, nil, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := activityx.Action{
		Type:     typex.ExchangeStaking,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeStaking{
			Action: stakingAction,
			Token:  *tokenMetadata,
		},
	}

	return &action, nil
}

// buildChipsMintAction builds the ChipsMint action.
func (w *worker) buildChipsMintAction(ctx context.Context, task *source.Task, from, to common.Address, contract common.Address, id *big.Int, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &contract, id, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	return &activityx.Action{
		Type:     typex.CollectibleMint,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}, nil
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

// Handle invalid token error.
func isInvalidTokenErr(err error) bool {
	return err != nil && strings.Contains(err.Error(), "unsupported NFT standard")
}

func NewWorker(config *config.Module, redisClient rueidis.Client) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient, token.WithRueidisClient(redisClient))

	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))
	instance.erc1155Filterer = lo.Must(erc1155.NewERC1155Filterer(ethereum.AddressGenesis, nil))

	// Initialize RSS3 contract filterers.
	instance.stakingVSLFilterer = lo.Must(rss3.NewStakingVSLFilterer(ethereum.AddressGenesis, nil))
	instance.chipsFilterer = lo.Must(rss3.NewChipsFilterer(ethereum.AddressGenesis, nil))

	// Initialize VSL contract filterers.
	instance.contractL2StandardBridgeFilterer = lo.Must(vsl.NewL2StandardBridgeFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
