package aave

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

func (w *worker) transformV1LendingPoolDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV1LendingPoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV1LendingPoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.AmountMinusFees)
}

func (w *worker) transformV2LendingPoolDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolWithdrawLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse withdraw event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityWithdraw, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse borrow event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse repay event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolSupplyLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v3PoolFilterer.ParseSupply(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse supply event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolWithdrawLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v3PoolFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse withdraw event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, pool, lender, metadata.ActionExchangeLiquidityWithdraw, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v3PoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse borrow event: %w", err)
	}

	lender, borrower := log.Address, event.User

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, borrower, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.v3PoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse repay event: %w", err)
	}

	lender, borrower := log.Address, event.User

	return w.buildEthereumExchangeLiquidityAction(ctx, task, borrower, lender, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.Amount)
}

func (w *worker) buildEthereumExchangeLiquidityAction(ctx context.Context, task *source.Task, from, to common.Address, exchangeLiquidityAction metadata.ExchangeLiquidityAction, tokenAddress common.Address, tokenValue *big.Int) (*schema.Action, error) {
	token, err := w.tokenClient.Lookup(ctx, task.ChainID, &tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenAddress, err)
	}

	token.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := schema.Action{
		Type:     filter.TypeExchangeLiquidity,
		Platform: filter.PlatformAAVE.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: exchangeLiquidityAction,
			Tokens: []metadata.Token{
				*token,
			},
		},
	}

	return &action, nil
}
