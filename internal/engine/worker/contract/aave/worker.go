package aave

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/aave"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
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
	return engine.Aave.String()
}

// Filter contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			aave.AddressV1LendingPool,
			aave.AddressV2LendingPoolMainnet,
			aave.AddressV2LendingPoolPolygon,
			aave.AddressV2LendingPoolAvalanche,
			aave.AddressV3PoolMainnet,
			aave.AddressV3PoolBase,
			aave.AddressV3PoolOthers,
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

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.GetNetwork().Source() == filter.NetworkEthereumSource, nil
}

// Transform Ethereum task to feed.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build feed base from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformAAVE))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	var actions []*schema.Action

	switch {
	case w.matchLiquidityV1Pool(ethereumTask):
		actions, err = w.handleV1LendingPool(ctx, ethereumTask)
	case w.matchLiquidityV2LendingPool(ethereumTask):
		actions, err = w.handleV2LendingPool(ctx, ethereumTask)
	case w.matchLiquidityV3Pool(ethereumTask):
		actions, err = w.handleV3Pool(ctx, ethereumTask)
	default:
		return nil, fmt.Errorf("unsupported transaction: %w", err)
	}

	if err != nil {
		return nil, err
	}

	feed.Type = filter.TypeExchangeLiquidity
	feed.Actions = append(feed.Actions, actions...)

	return feed, nil
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
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
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
