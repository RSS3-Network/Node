package aave

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/aave"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/erc1155"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/v2/provider/ethereum/token"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
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
	// Base
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	// Specific filters
	v1LendingPoolFilterer *aave.V1LendingPoolFilterer
	v2LendingPoolFilterer *aave.V2LendingPoolFilterer
	v3PoolFilterer        *aave.V3PoolFilterer
	// Common filters
	erc20Filterer   *erc20.ERC20Filterer
	erc721Filterer  *erc721.ERC721Filterer
	erc1155Filterer *erc1155.ERC1155Filterer
}

func (w *worker) Name() string {
	return decentralized.Aave.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformAAVE.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Polygon,
		network.Avalanche,
		network.Base,
		network.Optimism,
		network.Arbitrum,
		network.Gnosis,
		network.BinanceSmartChain,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeLiquidity,
	}
}

// Filter contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	var aaveV2LendingAddress, aaveV3PoolAddress common.Address

	switch w.config.Network {
	case network.Ethereum:
		aaveV2LendingAddress = aave.AddressV2LendingPoolMainnet
		aaveV3PoolAddress = aave.AddressV3PoolMainnet
	case network.Polygon:
		aaveV2LendingAddress = aave.AddressV2LendingPoolPolygon
		aaveV3PoolAddress = aave.AddressV3PoolOthers
	case network.Avalanche:
		aaveV2LendingAddress = aave.AddressV2LendingPoolAvalanche
		aaveV3PoolAddress = aave.AddressV3PoolOthers
	case network.Base:
		aaveV3PoolAddress = aave.AddressV3PoolBase
	case network.Gnosis:
		aaveV3PoolAddress = aave.AddressV3PoolGnosis
	case network.BinanceSmartChain:
		aaveV3PoolAddress = aave.AddressV3PoolBinanceSmartChain
	case network.Optimism, network.Arbitrum:
		aaveV3PoolAddress = aave.AddressV3PoolOthers
	default:
		aaveV2LendingAddress = aave.AddressV2LendingPoolMainnet
		aaveV3PoolAddress = aave.AddressV3PoolMainnet
	}

	return &source.Filter{
		LogAddresses: []common.Address{
			aave.AddressV1LendingPool,
			aaveV2LendingAddress,
			aaveV3PoolAddress,
		},
		LogTopics: []common.Hash{
			aave.EventHashV1LendingPoolDeposit,
			aave.EventHashV1LendingPoolBorrow,
			aave.EventHashV1LendingPoolRepay,
			aave.EventHashV2LendingPoolDeposit,
			aave.EventHashV2LendingPoolWithdraw,
			aave.EventHashV2LendingPoolBorrow,
			aave.EventHashV2LendingPoolRepay,
			aave.EventHashV3PoolSupply,
			aave.EventHashV3PoolWithdraw,
			aave.EventHashV3PoolBorrow,
			aave.EventHashV3PoolRepay,
		},
	}
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	if ethereumTask.Transaction.To == nil {
		return nil, fmt.Errorf("invalid transaction to: %s", ethereumTask.Transaction.Hash)
	}

	// Build activity base from task.
	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w, task_id: %s", err, ethereumTask.ID())
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
		case w.matchLiquidityV1Pool(ethereumTask, log):
			actions, err = w.handleV1LendingPool(ctx, ethereumTask, log)
		case w.matchLiquidityV2LendingPool(ethereumTask, log):
			actions, err = w.handleV2LendingPool(ctx, ethereumTask, log)
		case w.matchLiquidityV3Pool(ethereumTask, log):
			actions, err = w.handleV3Pool(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("handle pool log: %w", err)
		}

		activity.Type = typex.ExchangeLiquidity

		if actions != nil {
			activity.Actions = append(activity.Actions, actions...)
		}
	}

	return activity, nil
}

func (w *worker) matchLiquidityV1Pool(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(
		log.Topics[0],
		aave.EventHashV1LendingPoolDeposit,
		aave.EventHashV1LendingPoolBorrow,
		aave.EventHashV1LendingPoolRepay,
	)
}

func (w *worker) matchLiquidityV2LendingPool(task *source.Task, log *ethereum.Log) bool {
	switch task.Network {
	case network.Ethereum:
		if *task.Transaction.To != aave.AddressV2LendingPoolMainnet {
			return false
		}
	case network.Polygon:
		if *task.Transaction.To != aave.AddressV2LendingPoolPolygon {
			return false
		}
	case network.Avalanche:
		if *task.Transaction.To != aave.AddressV2LendingPoolAvalanche {
			return false
		}
	default:
		return false
	}

	return contract.MatchEventHashes(
		log.Topics[0],
		aave.EventHashV2LendingPoolDeposit,
		aave.EventHashV2LendingPoolWithdraw,
		aave.EventHashV2LendingPoolBorrow,
		aave.EventHashV2LendingPoolRepay,
	)
}

func (w *worker) matchLiquidityV3Pool(task *source.Task, log *ethereum.Log) bool {
	switch task.Network {
	case network.Ethereum:
		if *task.Transaction.To != aave.AddressV3PoolMainnet {
			return false
		}
	case network.Base:
		if *task.Transaction.To != aave.AddressV3PoolBase {
			return false
		}
	case
		network.Optimism,
		network.Arbitrum,
		network.Polygon,
		network.Avalanche:
		if *task.Transaction.To != aave.AddressV3PoolOthers {
			return false
		}
	default:
		return false
	}

	return contract.MatchEventHashes(
		log.Topics[0],
		aave.EventHashV3PoolSupply,
		aave.EventHashV3PoolWithdraw,
		aave.EventHashV3PoolBorrow,
		aave.EventHashV3PoolRepay,
	)
}

func (w *worker) handleV1LendingPool(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	var (
		action *activityx.Action
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
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("handle v1 pool: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) handleV2LendingPool(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	var (
		action *activityx.Action
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
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("handle v2 pool: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) handleV3Pool(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	var (
		action *activityx.Action
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
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("handle v3 pool: %w", err)
	}

	return []*activityx.Action{action}, nil
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

func (w *worker) transformV1LendingPoolDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV1LendingPoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV1LendingPoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v1LendingPoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.AmountMinusFees)
}

func (w *worker) transformV2LendingPoolDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolWithdrawLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse withdraw event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityWithdraw, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse borrow event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV2LendingPoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v2LendingPoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse repay event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolSupplyLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v3PoolFilterer.ParseSupply(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse supply event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, pool, metadata.ActionExchangeLiquiditySupply, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolWithdrawLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v3PoolFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse withdraw event: %w", err)
	}

	lender, pool := event.User, log.Address

	return w.buildEthereumExchangeLiquidityAction(ctx, task, pool, lender, metadata.ActionExchangeLiquidityWithdraw, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolBorrowLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v3PoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse borrow event: %w", err)
	}

	lender, borrower := log.Address, event.User

	return w.buildEthereumExchangeLiquidityAction(ctx, task, lender, borrower, metadata.ActionExchangeLiquidityBorrow, event.Reserve, event.Amount)
}

func (w *worker) transformV3PoolRepayLog(ctx context.Context, task *source.Task, log *ethereum.Log) (*activityx.Action, error) {
	event, err := w.v3PoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse repay event: %w", err)
	}

	lender, borrower := log.Address, event.User

	return w.buildEthereumExchangeLiquidityAction(ctx, task, borrower, lender, metadata.ActionExchangeLiquidityRepay, event.Reserve, event.Amount)
}

func (w *worker) buildEthereumExchangeLiquidityAction(ctx context.Context, task *source.Task, from, to common.Address, exchangeLiquidityAction metadata.ExchangeLiquidityAction, tokenAddress common.Address, tokenValue *big.Int) (*activityx.Action, error) {
	targetToken, err := w.tokenClient.Lookup(ctx, task.ChainID, &tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenAddress, err)
	}

	targetToken.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(tokenValue), 0))

	action := activityx.Action{
		Type:     typex.ExchangeLiquidity,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: exchangeLiquidityAction,
			Tokens: []metadata.Token{
				*targetToken,
			},
		},
	}

	return &action, nil
}

// NewWorker creates a new worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	// Initialize ethereum client.
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize filterers.
	instance.v1LendingPoolFilterer = lo.Must(aave.NewV1LendingPoolFilterer(ethereum.AddressGenesis, nil))
	instance.v2LendingPoolFilterer = lo.Must(aave.NewV2LendingPoolFilterer(ethereum.AddressGenesis, nil))
	instance.v3PoolFilterer = lo.Must(aave.NewV3PoolFilterer(ethereum.AddressGenesis, nil))

	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))
	instance.erc1155Filterer = lo.Must(erc1155.NewERC1155Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
