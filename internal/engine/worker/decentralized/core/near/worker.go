package near

import (
	"context"
	"fmt"
	"math/big"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/near"
	"github.com/rss3-network/node/internal/utils"
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
		typex.TransactionTransfer,
	}
}

// Filter returns a protocol filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return nil
}

// Transform returns an activity with the actions of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity()
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle all actions in the transaction
	actions, err := w.handleNearActions(ctx, nearTask)
	if err != nil {
		return nil, fmt.Errorf("handle near actions: %w", err)
	}

	if len(actions) > 0 {
		activity.Type = actions[0].Type
		activity.Actions = append(activity.Actions, actions...)
	} else {
		return nil, fmt.Errorf("no actions found in transaction")
	}

	return activity, nil
}

// handleNearActions processes all actions in the Near transaction and returns a slice of activityx.Action.
func (w *worker) handleNearActions(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	for _, action := range task.Transaction.Transaction.Actions {
		if action.Transfer != nil {
			transferAction, err := w.handleNearTransferAction(ctx, task.Transaction.Transaction.SignerID, task.Transaction.Transaction.ReceiverID, action.Transfer)
			if err != nil {
				return nil, fmt.Errorf("handle transfer action: %w", err)
			}

			if transferAction != nil {
				actions = append(actions, transferAction)
			}
		}
	}

	return actions, nil
}

// handleNearTransferAction processes a single Transfer action and returns an activityx.Action.
func (w *worker) handleNearTransferAction(ctx context.Context, from, to string, transferAction *near.TransferAction) (*activityx.Action, error) {
	value, err := decimal.NewFromString(transferAction.Deposit)
	if err != nil {
		return nil, fmt.Errorf("failed to parse deposit value: %w", err)
	}

	if value.IsZero() {
		return nil, nil // Skip zero-value transfers
	}

	return w.buildNearTransactionTransferAction(ctx, from, to, value.BigInt())
}

// buildNearTransactionTransferAction returns the native transfer transaction action.
func (w *worker) buildNearTransactionTransferAction(_ context.Context, from, to string, tokenValue *big.Int) (*activityx.Action, error) {
	return &activityx.Action{
		Type: typex.TransactionTransfer,
		From: from,
		To:   to,
		Metadata: metadata.TransactionTransfer{
			Name:     "NEAR",
			Symbol:   "NEAR",
			Decimals: 24,
			Value:    lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(tokenValue), 0)),
		},
	}, nil
}

// NewWorker returns a new Near worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
