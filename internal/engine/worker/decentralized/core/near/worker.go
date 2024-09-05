package near

import (
	"context"
	"fmt"
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/near"
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
	if w.matchNearNativeTransferTransaction(nearTask) {
		// Handle the native transfer transaction.
		action, err := w.handleNearNativeTransferTransaction(ctx, nearTask)
		if err != nil {
			return nil, fmt.Errorf("handle native transfer transaction: %w", err)
		}

		activity.Type = action.Type
		activity.Actions = append(activity.Actions, action)
	}

	return activity, nil
}

// matchNearNativeTransferTransaction returns true if the transaction is a native transfer transaction.
func (w *worker) matchNearNativeTransferTransaction(task *source.Task) bool {
	return task.Transaction.Transaction.PriorityFee > 0
}

// handleNearNativeTransferTransaction returns the action of the native transfer transaction.
func (w *worker) handleNearNativeTransferTransaction(ctx context.Context, task *source.Task) (*activityx.Action, error) {
	value := decimal.NewFromInt(int64(task.Transaction.Transaction.PriorityFee))

	// Build the native transfer transaction action.
	return w.buildNearTransactionTransferAction(ctx, task.Transaction.Transaction.PublicKey, task.Transaction.Transaction.ReceiverID, value.BigInt())
}

// buildNearTransactionTransferAction returns the native transfer transaction action.
func (w *worker) buildNearTransactionTransferAction(_ context.Context, from, to string, tokenValue *big.Int) (*activityx.Action, error) {
	action := activityx.Action{
		Type: typex.TransactionTransfer,
		From: from,
		To:   to,
		Metadata: metadata.TransactionTransfer{
			Value: lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0)),
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
