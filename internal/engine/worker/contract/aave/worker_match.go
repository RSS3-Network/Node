package aave

import (
	"context"
	"fmt"

	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/aave"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

func (w *worker) matchLiquidityV1Pool(task *source.Task) bool {
	if task.Network != filter.NetworkEthereum {
		return false
	}

	if *task.Transaction.To != aave.AddressV1LendingPool {
		return false
	}

	return lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
		if len(log.Topics) == 0 {
			return false
		}

		return contract.MatchEventHashes(
			log.Topics[0],
			aave.EventHashV1LendingPoolDeposit,
			aave.EventHashV1LendingPoolBorrow,
			aave.EventHashV1LendingPoolRepay,
		)
	})
}

func (w *worker) matchLiquidityV2LendingPool(task *source.Task) bool {
	switch task.Network {
	case filter.NetworkEthereum:
		if *task.Transaction.To != aave.AddressV2LendingPoolMainnet {
			return false
		}
	case filter.NetworkPolygon:
		if *task.Transaction.To != aave.AddressV2LendingPoolPolygon {
			return false
		}
	case filter.NetworkAvalanche:
		if *task.Transaction.To != aave.AddressV2LendingPoolAvalanche {
			return false
		}
	default:
		return false
	}

	return lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
		if len(log.Topics) == 0 {
			return false
		}

		return contract.MatchEventHashes(
			log.Topics[0],
			aave.EventHashV2LendingPoolDeposit,
			aave.EventHashV2LendingPoolWithdraw,
			aave.EventHashV2LendingPoolBorrow,
			aave.EventHashV2LendingPoolRepay,
		)
	})
}

func (w *worker) matchLiquidityV3Pool(task *source.Task) bool {
	switch task.Network {
	case filter.NetworkEthereum:
		if *task.Transaction.To != aave.AddressV3PoolMainnet {
			return false
		}
	case filter.NetworkBase:
		if *task.Transaction.To != aave.AddressV3PoolBase {
			return false
		}
	case
		filter.NetworkOptimism,
		filter.NetworkArbitrumOne,
		filter.NetworkPolygon,
		filter.NetworkFantom,
		filter.NetworkAvalanche:
		if *task.Transaction.To != aave.AddressV3PoolOthers {
			return false
		}
	default:
		return false
	}

	return lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
		if len(log.Topics) == 0 {
			return false
		}

		return contract.MatchEventHashes(
			log.Topics[0],
			aave.EventHashV3PoolSupply,
			aave.EventHashV3PoolWithdraw,
			aave.EventHashV3PoolBorrow,
			aave.EventHashV3PoolRepay,
		)
	})
}

func (w *worker) handleV1LendingPool(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	actions := make([]*schema.Action, 0)

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			action *schema.Action
			err    error
		)

		switch {
		case w.matchEthereumV1LendingPoolDepositLog(log):
			action, err = w.transformV1LendingPoolDepositLog(ctx, task, log)
		case w.matchEthereumV1LendingPoolBorrowLog(log):
			action, err = w.transformV1LendingPoolBorrowLog(ctx, task, log)
		case w.matchEthereumV1LendingPoolRepayLog(log):
			action, err = w.transformV1LendingPoolRepayLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("handle ethereum v3 pool: %w", err)
		}

		if action != nil {
			actions = append(actions, action)
		}
	}

	return actions, nil
}

func (w *worker) handleV2LendingPool(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	actions := make([]*schema.Action, 0)

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			action *schema.Action
			err    error
		)

		switch {
		case w.matchEthereumV2LendingPoolDepositLog(log):
			action, err = w.transformV2LendingPoolDepositLog(ctx, task, log)
		case w.matchEthereumV2LendingPoolWithdrawLog(log):
			action, err = w.transformV2LendingPoolWithdrawLog(ctx, task, log)
		case w.matchEthereumV2LendingPoolBorrowLog(log):
			action, err = w.transformV2LendingPoolBorrowLog(ctx, task, log)
		case w.matchEthereumV2LendingPoolRepayLog(log):
			action, err = w.transformV2LendingPoolRepayLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("handle ethereum v3 pool: %w", err)
		}

		if action != nil {
			actions = append(actions, action)
		}
	}

	return actions, nil
}

func (w *worker) handleV3Pool(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	actions := make([]*schema.Action, 0)

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			action *schema.Action
			err    error
		)

		switch {
		case w.matchEthereumV3PoolSupplyLog(log):
			action, err = w.transformV3PoolSupplyLog(ctx, task, log)
		case w.matchEthereumV3PoolWithdrawLog(log):
			action, err = w.transformV3PoolWithdrawLog(ctx, task, log)
		case w.matchEthereumV3PoolBorrowLog(log):
			action, err = w.transformV3PoolBorrowLog(ctx, task, log)
		case w.matchEthereumV3PoolRepayLog(log):
			action, err = w.transformV3PoolRepayLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("handle ethereum v3 pool: %w", err)
		}

		if action != nil {
			actions = append(actions, action)
		}
	}

	return actions, nil
}

func (w *worker) matchEthereumV1LendingPoolDepositLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV1LendingPool) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV1LendingPoolDeposit)
}

func (w *worker) matchEthereumV1LendingPoolBorrowLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV1LendingPool) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV1LendingPoolBorrow)
}

func (w *worker) matchEthereumV1LendingPoolRepayLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV1LendingPool) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV1LendingPoolRepay)
}

func (w *worker) matchEthereumV2LendingPoolDepositLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV2LendingPoolMainnet, aave.AddressV2LendingPoolPolygon, aave.AddressV2LendingPoolAvalanche) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV2LendingPoolDeposit)
}

func (w *worker) matchEthereumV2LendingPoolWithdrawLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV2LendingPoolMainnet, aave.AddressV2LendingPoolPolygon, aave.AddressV2LendingPoolAvalanche) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV2LendingPoolWithdraw)
}

func (w *worker) matchEthereumV2LendingPoolBorrowLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV2LendingPoolMainnet, aave.AddressV2LendingPoolPolygon, aave.AddressV2LendingPoolAvalanche) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV2LendingPoolBorrow)
}

func (w *worker) matchEthereumV2LendingPoolRepayLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV2LendingPoolMainnet, aave.AddressV2LendingPoolPolygon, aave.AddressV2LendingPoolAvalanche) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV2LendingPoolRepay)
}

func (w *worker) matchEthereumV3PoolSupplyLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV3PoolMainnet, aave.AddressV3PoolBase, aave.AddressV3PoolOthers) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV3PoolSupply)
}

func (w *worker) matchEthereumV3PoolWithdrawLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV3PoolMainnet, aave.AddressV3PoolBase, aave.AddressV3PoolOthers) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV3PoolWithdraw)
}

func (w *worker) matchEthereumV3PoolBorrowLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV3PoolMainnet, aave.AddressV3PoolBase, aave.AddressV3PoolOthers) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV3PoolBorrow)
}

func (w *worker) matchEthereumV3PoolRepayLog(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, aave.AddressV3PoolMainnet, aave.AddressV3PoolBase, aave.AddressV3PoolOthers) && contract.MatchEventHashes(log.Topics[0], aave.EventHashV3PoolRepay)
}
