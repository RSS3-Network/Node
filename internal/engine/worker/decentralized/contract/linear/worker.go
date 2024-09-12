package linear

import (
	"context"
	"encoding/json"
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

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		ReceiverIDs: []string{
			liNEARReceiverID,
		},
	}
}

// Transform returns an activity with the actions of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to an Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

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
		activity.Type = actions[0].Type
		activity.Actions = append(activity.Actions, actions...)
	}

	return activity, nil
}

// handleLiNEARActions processes all actions in the LiNEAR transaction and returns a slice of activityx.Action.
func (w *worker) handleLiNEARActions(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	for _, action := range task.Transaction.Transaction.Actions {
		if action.FunctionCall != nil && (action.FunctionCall.MethodName == liNEARFTTransferMethod || action.FunctionCall.MethodName == liNEARFTTransferCallMethod) {
			functionCallActions, err := w.handleLiNEARFunctionCallAction(ctx, task.Transaction.Transaction.SignerID, task.Transaction.Transaction.ReceiverID, action.FunctionCall)
			if err != nil {
				return nil, fmt.Errorf("handle function call action: %w", err)
			}

			actions = append(actions, functionCallActions...)
		}
	}

	return actions, nil
}

// handleLiNEARFunctionCallAction processes a single FunctionCall action and returns an array of activityx.Action.
func (w *worker) handleLiNEARFunctionCallAction(_ context.Context, from, to string, functionCallAction *near.FunctionCallAction) ([]*activityx.Action, error) {
	// Decode function call args
	args, err := near.DecodeBase64(functionCallAction.Args)
	if err != nil {
		return nil, fmt.Errorf("decode function call args: %w", err)
	}

	var functionCallArgs FunctionCallArgs
	if err := json.Unmarshal(args, &functionCallArgs); err != nil {
		return nil, fmt.Errorf("unmarshal LiNEAR function call args: %w", err)
	}

	amount, ok := new(big.Int).SetString(functionCallArgs.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount: %s", functionCallArgs.Amount)
	}

	// Build transfer action
	transferAction := w.buildNearTransactionTransferAction(from, functionCallArgs.ReceiverID, amount)

	// unmarshal and parse msg
	var lineaMsg Msg

	if err := json.Unmarshal([]byte(functionCallArgs.Msg), &lineaMsg); err != nil {
		return nil, fmt.Errorf("unmarshal LiNEAR msg: %w", err)
	}

	actions := []*activityx.Action{transferAction}

	// If there are swap actions, build and add them
	if len(lineaMsg.Actions) > 0 {
		swapActions, err := w.buildLiNEARSwapActions(from, to, lineaMsg)
		if err != nil {
			return nil, fmt.Errorf("build swap actions: %w", err)
		}

		actions = append(actions, swapActions...)
	}

	return actions, nil
}

// buildNearTransactionTransferAction returns the native transfer transaction action.
func (w *worker) buildNearTransactionTransferAction(from, to string, tokenValue *big.Int) *activityx.Action {
	return &activityx.Action{
		Type: typex.TransactionTransfer,
		From: from,
		To:   to,
		Metadata: metadata.TransactionTransfer{
			Name:     "LNR",
			Symbol:   "LNR",
			Decimals: 18,
			Value:    lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0)),
		},
	}
}

// buildLiNEARSwapActions builds exchange swap actions for LiNEAR.
func (w *worker) buildLiNEARSwapActions(from, to string, linearMsg Msg) ([]*activityx.Action, error) {
	if len(linearMsg.Actions) == 0 {
		return nil, fmt.Errorf("no swap actions provided")
	}

	swapActions := make([]*activityx.Action, 0, len(linearMsg.Actions))

	for _, swapAction := range linearMsg.Actions {
		var amountIn *big.Int

		if swapAction.AmountIn != "" {
			var ok bool

			amountIn, ok = new(big.Int).SetString(swapAction.AmountIn, 10)
			if !ok {
				return nil, fmt.Errorf("invalid amount_in: %s", swapAction.AmountIn)
			}
		} else {
			amountIn = big.NewInt(0)
		}

		minAmountOut, ok := new(big.Int).SetString(swapAction.MinAmountOut, 10)
		if !ok {
			return nil, fmt.Errorf("invalid min_amount_out: %s", swapAction.MinAmountOut)
		}

		tokenIn := swapAction.TokenIn
		tokenOut := swapAction.TokenOut

		fromToken := &metadata.Token{
			Address:  &tokenIn,
			Value:    lo.ToPtr(decimal.NewFromBigInt(amountIn, 0)),
			Decimals: 0,
		}

		toToken := &metadata.Token{
			Address:  &tokenOut,
			Value:    lo.ToPtr(decimal.NewFromBigInt(minAmountOut, 0)),
			Decimals: 0,
		}

		action := &activityx.Action{
			Type:     typex.ExchangeSwap,
			Platform: w.Platform(),
			From:     from,
			To:       to,
			Metadata: metadata.ExchangeSwap{
				From: *fromToken,
				To:   *toToken,
			},
		}

		swapActions = append(swapActions, action)
	}

	return swapActions, nil
}

// NewWorker returns a new LiNEAR worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
