package linear

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/near"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
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
	config *config.Module
}

func (w *worker) Name() string {
	return workerx.LiNEAR.String()
}

func (w *worker) Platform() string {
	return workerx.PlatformLiNEAR.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Near,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
		tag.Transaction,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeSwap,
		typex.TransactionTransfer,
	}
}

const (
	liNEARReceiverID           = "802d89b6e511b335f05024a65161bce7efc3f311.factory.bridge.near"
	liNEARFTTransferMethod     = "ft_transfer"
	liNEARFTTransferCallMethod = "ft_transfer_call"
)

// Filter returns a protocol filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		ReceiverIDs: []string{
			liNEARReceiverID,
		},
	}
}

// Transform returns an activity with the actions of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	zap.L().Debug("transforming linear task", zap.String("task_id", nearTask.ID()))

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle all actions in the transaction
	actions, err := w.handleLiNEARActions(ctx, nearTask)
	if err != nil {
		return nil, fmt.Errorf("handle LiNEAR actions: %w", err)
	}

	if len(actions) > 0 {
		// Check if there are any swap actions
		hasSwapAction := lo.ContainsBy(actions, func(action *activityx.Action) bool {
			return action.Type == typex.ExchangeSwap
		})

		if hasSwapAction {
			zap.L().Debug("detected swap action, setting activity type to ExchangeSwap")

			activity.Type = typex.ExchangeSwap
		} else {
			zap.L().Debug("no swap action detected, using first action type",
				zap.String("type", actions[0].Type.Name()))

			activity.Type = actions[0].Type
		}

		activity.Actions = append(activity.Actions, actions...)
	} else {
		return nil, fmt.Errorf("no actions found in transaction")
	}

	zap.L().Debug("successfully transformed linear task")

	return activity, nil
}

// handleLiNEARActions processes all actions in the LiNEAR transaction and returns a slice of activityx.Action.
func (w *worker) handleLiNEARActions(_ context.Context, task *source.Task) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	var events []Event

	for _, receipt := range task.Transaction.ReceiptsOutcome {
		for _, log := range receipt.Outcome.Logs {
			if strings.HasPrefix(log, "EVENT_JSON:") {
				eventJSON := strings.TrimPrefix(log, "EVENT_JSON:")

				var event Event

				if err := json.Unmarshal([]byte(eventJSON), &event); err != nil {
					return nil, fmt.Errorf("unmarshal event JSON: %w", err)
				}

				event.TokenAddress = receipt.Outcome.ExecutorID
				events = append(events, event)
			}
		}
	}

	if len(events) == 1 {
		// Transfer action
		action, err := w.buildTransferAction(&events[0])
		if err != nil {
			return nil, fmt.Errorf("build transfer action: %w", err)
		}

		actions = append(actions, action)
	} else if len(events) == 2 {
		// Swap action
		action, err := w.buildSwapAction(task.Transaction.Transaction.SignerID, &events[0], &events[1])
		if err != nil {
			return nil, fmt.Errorf("build swap action: %w", err)
		}

		actions = append(actions, action)
	}

	return actions, nil
}

func (w *worker) buildTransferAction(event *Event) (*activityx.Action, error) {
	if len(event.Data) == 0 {
		return nil, fmt.Errorf("no data in transfer event")
	}

	data := event.Data[0]

	zap.L().Debug("building transfer action",
		zap.String("old_owner", data.OldOwnerID),
		zap.String("new_owner", data.NewOwnerID),
		zap.String("amount", data.Amount))

	amount, ok := new(big.Int).SetString(data.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount: %s", data.Amount)
	}

	action := &activityx.Action{
		Type: typex.TransactionTransfer,
		From: data.OldOwnerID,
		To:   data.NewOwnerID,
		Metadata: metadata.TransactionTransfer{
			Name:     "LNR",
			Symbol:   "LNR",
			Decimals: 18,
			Value:    lo.ToPtr(decimal.NewFromBigInt(amount, 0)),
			Standard: getTokenStandard(event.Standard),
			Address:  lo.ToPtr(event.TokenAddress),
		},
	}

	zap.L().Debug("successfully built transfer action")

	return action, nil
}

func (w *worker) buildSwapAction(signerID string, event1, event2 *Event) (*activityx.Action, error) {
	zap.L().Debug("building swap action",
		zap.String("signer_id", signerID))

	var fromToken, toToken metadata.Token

	var toAddress string

	for _, event := range []*Event{event1, event2} {
		if len(event.Data) == 0 {
			zap.L().Warn("empty event data, skipping")
			continue
		}

		data := event.Data[0]

		amount, ok := new(big.Int).SetString(data.Amount, 10)
		if !ok {
			return nil, fmt.Errorf("invalid amount: %s", data.Amount)
		}

		token := metadata.Token{
			Value:    lo.ToPtr(decimal.NewFromBigInt(amount, 0)),
			Standard: getTokenStandard(event.Standard),
			Address:  lo.ToPtr(event.TokenAddress),
		}

		if data.OldOwnerID == signerID {
			fromToken = token
			toAddress = data.NewOwnerID

			zap.L().Debug("found from token",
				zap.String("old_owner", data.OldOwnerID),
				zap.String("amount", data.Amount))
		} else if data.NewOwnerID == signerID {
			toToken = token

			zap.L().Debug("found to token",
				zap.String("new_owner", data.NewOwnerID),
				zap.String("amount", data.Amount))
		}
	}

	action := &activityx.Action{
		Type:     typex.ExchangeSwap,
		Platform: w.Platform(),
		From:     signerID,
		To:       toAddress,
		Metadata: metadata.ExchangeSwap{
			From: fromToken,
			To:   toToken,
		},
	}

	zap.L().Debug("successfully built swap action")

	return action, nil
}

func getTokenStandard(standard string) metadata.Standard {
	switch standard {
	case "nep141":
		return metadata.StandardNEP141
	default:
		return metadata.StandardUnknown
	}
}

// NewWorker returns a new LiNEAR worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
