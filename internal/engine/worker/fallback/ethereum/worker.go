package ethereum

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	ethereumx "github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config *engine.Config
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.Network() == filter.NetworkEthereum, nil
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*ethereum.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	feed, err := task.BuildFeed()
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	if w.matchEthereumFailedTransaction(ethereumTask) {
		return feed, nil
	}

	if w.matchEthereumNativeTransferTransaction(ethereumTask) {
		action, err := w.handleEthereumNativeTransferTransaction(ctx, ethereumTask)
		if err != nil {
			return nil, fmt.Errorf("handle native transfer transaction: %w", err)
		}

		feed.Type = action.Type
		feed.Actions = append(feed.Actions, action)
	}

	return feed, nil
}

func (w *worker) matchEthereumFailedTransaction(task *ethereum.Task) bool {
	return task.Receipt.Status == types.ReceiptStatusFailed
}

func (w *worker) matchEthereumNativeTransferTransaction(task *ethereum.Task) bool {
	return task.Transaction.Value().Cmp(big.NewInt(0)) == 1 && len(task.Transaction.Data()) == 0
}

func (w *worker) handleEthereumNativeTransferTransaction(ctx context.Context, task *ethereum.Task) (*schema.Action, error) {
	transactionFrom, err := types.LatestSignerForChainID(task.Transaction.ChainId()).Sender(task.Transaction)
	if err != nil {
		return nil, fmt.Errorf("recover transaction signer: %w", err)
	}

	var transactionTo = ethereumx.AddressGenesis
	if task.Transaction.To() != nil {
		transactionTo = *task.Transaction.To()
	}

	return w.buildEthereumTransactionTransferAction(ctx, task, transactionFrom, transactionTo, nil, task.Transaction.Value())
}

func (w *worker) buildEthereumTransactionTransferAction(_ context.Context, _ *ethereum.Task, from, to common.Address, _ *common.Address, tokenValue *big.Int) (*schema.Action, error) {
	// TODO Add support for ERC-20 and ERC-1577 token transfer transactions.
	action := schema.Action{
		Type: filter.TypeTransactionTransfer,
		From: from.String(),
		To:   to.String(),
		Metadata: metadata.TransactionTransfer{
			Value: lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0)),
		},
	}

	return &action, nil
}

func NewWorker(config *engine.Config) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
