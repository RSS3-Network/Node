package near

import (
	"context"
	"fmt"
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/near"
	"github.com/rss3-network/node/provider/near"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config *config.Module
}

func (w *worker) Name() string {
	return workerx.Core.String()
}

func (w *worker) Platform() string {
	return ""
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Near,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Unknown,
		tag.Transaction,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.Unknown,
	}
}

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return nil
}

// Transform returns an activity  with the action of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to an Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity()
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// If the task is a native transfer transaction, handle it.
	if w.matchTransferTransaction(nearTask) {
		// Handle the native transfer transaction.
		action, err := w.handleNearTransferTransaction(ctx, nearTask)
		if err != nil {
			return nil, fmt.Errorf("handle native transfer transaction: %w", err)
		}

		activity.Type = action.Type
		activity.Actions = append(activity.Actions, action)
	}

	return activity, nil
}

// matchTransferTransaction returns true if the transaction is a transfer transaction.
func (w *worker) matchTransferTransaction(task *source.Task) bool {
	// Check if the transaction has actions
	if len(task.Transaction.Transaction.Actions) == 0 {
		return false
	}

	// Iterate through the actions to find a Transfer action
	for _, action := range task.Transaction.Transaction.Actions {
		// add logic to check deposit is larger than 0, use decimal
		if action.Transfer != nil {
			deposit, err := decimal.NewFromString(action.Transfer.Deposit)
			if err == nil && decimal.NewFromInt(0).LessThan(deposit) {
				return true
			}
		}
	}

	// If no Transfer action is found, it's not a native transfer transaction
	return false
}

// handleNearNativeTransferTransaction returns the action of the native transfer transaction.
func (w *worker) handleNearTransferTransaction(ctx context.Context, task *source.Task) (*activityx.Action, error) {
	// Find the Transfer action and get its deposit value
	var transferAction *near.TransferAction

	for _, action := range task.Transaction.Transaction.Actions {
		if action.Transfer != nil {
			transferAction = action.Transfer
			break
		}
	}

	if transferAction == nil {
		return nil, fmt.Errorf("no transfer action found in transaction")
	}

	value, err := decimal.NewFromString(transferAction.Deposit)
	if err != nil {
		return nil, fmt.Errorf("failed to parse deposit value: %w", err)
	}

	// Build the native transfer transaction action.
	return w.buildNearTransactionTransferAction(ctx, task.Transaction.Transaction.SignerID, task.Transaction.Transaction.ReceiverID, value.BigInt())
}

// buildNearTransactionTransferAction returns the native transfer transaction action.
func (w *worker) buildNearTransactionTransferAction(_ context.Context, from, to string, tokenValue *big.Int) (*activityx.Action, error) {
	action := activityx.Action{
		Type: typex.TransactionTransfer,
		From: from,
		To:   to,
		Metadata: metadata.TransactionTransfer{
			Name:     "NEAR",
			Symbol:   "NEAR",
			Decimals: 24,
			Value:    lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0)),
		},
	}

	return &action, nil
}

// NewWorker returns a new Arweave worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
