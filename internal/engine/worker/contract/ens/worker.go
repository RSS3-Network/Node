package ens

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/ens"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc1155"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

// Worker is the worker for ENS.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	// Base
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	databaseClient database.Client
	// Specific filters
	baseRegistrarImplementationFilterer *ens.BaseRegistrarImplementationFilterer
	ethRegistrarControllerV1Filterer    *ens.ETHRegistrarControllerV1Filterer
	ethRegistrarControllerV2Filterer    *ens.ETHRegistrarControllerV2Filterer
	publicResolverV1Filterer            *ens.PublicResolverV1Filterer
	publicResolverV2Filterer            *ens.PublicResolverV2Filterer
	nameWrapperFilterer                 *ens.NameWrapperFilterer
	// Common filters
	erc20Filterer   *erc20.ERC20Filterer
	erc721Filterer  *erc721.ERC721Filterer
	erc1155Filterer *erc1155.ERC1155Filterer
	// Contract caller
	nameWrapper *ens.NameWrapperCaller
}

func (w *worker) Name() string {
	return engine.ENS.String()
}

// Filter ens contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			ens.AddressBaseRegistrarImplementation,
			ens.AddressETHRegistrarControllerV1,
			ens.AddressETHRegistrarControllerV2,
			ens.AddressPublicResolverV1,
			ens.AddressPublicResolverV2,
			ens.AddressNameWrapper,
		},
		LogTopics: []common.Hash{
			ens.EventNameRegisteredControllerV1,
			ens.EventNameRegisteredControllerV2,
			ens.EventNameRenewed,
			ens.EventTextChanged,
			ens.EventTextChangedWithValue,
			ens.EventNameWrapped,
			ens.EventNameUnwrapped,

			ens.EventFusesSet,
			ens.EventAddressChanged,
			ens.EventContenthashChanged,
			ens.EventNameChanged,
			ens.EventPubkeyChanged,
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

	// Build default ens feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformENS))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	exist := lo.ContainsBy(ethereumTask.Receipt.Logs, func(log *ethereum.Log) bool {
		return w.matchEnsNameRegisteredV2(ctx, *log) || w.matchEnsNameRegisteredV1(ctx, *log)
	})

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*schema.Action
			err     error
		)

		if len(log.Topics) == 0 {
			continue
		}

		if exist {
			switch {
			case w.matchEnsNameRegisteredV1(ctx, *log):
				feed.Type = filter.TypeCollectibleTrade
				actions, err = w.transformEnsNameRegisteredV1(ctx, *log, ethereumTask)
			case w.matchEnsNameRegisteredV2(ctx, *log):
				feed.Type = filter.TypeCollectibleTrade
				actions, err = w.transformEnsNameRegisteredV2(ctx, *log, ethereumTask)
			default:
				continue
			}
		} else {
			switch {
			case w.matchEnsNameRenewed(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsNameRenewed(ctx, *log, ethereumTask)
			case w.matchEnsTextChanged(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsTextChanged(ctx, *log, ethereumTask)
			case w.matchEnsTextChangedWithValue(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsTextChangedWithValue(ctx, *log, ethereumTask)
			case w.matchEnsNameWrapped(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsNameWrapped(ctx, *log, ethereumTask)
			case w.matchEnsNameUnwrapped(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsNameUnwrapped(ctx, *log, ethereumTask)
			case w.matchEnsFusesSet(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsFusesSet(ctx, *log, ethereumTask)
			case w.matchEnsContenthashChanged(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsContenthashChanged(ctx, *log, ethereumTask)
			case w.matchEnsNameChanged(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsNameChanged(ctx, *log, ethereumTask)
			case w.matchEnsAddressChanged(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsAddressChanged(ctx, *log, ethereumTask)
			case w.matchEnsPubkeyChanged(ctx, *log):
				feed.Type = filter.TypeSocialProfile
				actions, err = w.transformEnsPubkeyChanged(ctx, *log, ethereumTask)
			default:
				continue
			}
		}

		if err != nil {
			return nil, err
		}

		// Change feed type to the first action type.
		for _, action := range actions {
			feed.Type = action.Type
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// NewWorker creates a new ENS worker.
func NewWorker(config *config.Module, databaseClient database.Client) (engine.Worker, error) {
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

	// Initialize database client.
	instance.databaseClient = databaseClient

	// Initialize ens filterers and callers.
	instance.baseRegistrarImplementationFilterer = lo.Must(ens.NewBaseRegistrarImplementationFilterer(ethereum.AddressGenesis, nil))
	instance.ethRegistrarControllerV1Filterer = lo.Must(ens.NewETHRegistrarControllerV1Filterer(ethereum.AddressGenesis, nil))
	instance.ethRegistrarControllerV2Filterer = lo.Must(ens.NewETHRegistrarControllerV2Filterer(ethereum.AddressGenesis, nil))
	instance.publicResolverV1Filterer = lo.Must(ens.NewPublicResolverV1Filterer(ethereum.AddressGenesis, nil))
	instance.publicResolverV2Filterer = lo.Must(ens.NewPublicResolverV2Filterer(ethereum.AddressGenesis, nil))
	instance.nameWrapperFilterer = lo.Must(ens.NewNameWrapperFilterer(ethereum.AddressGenesis, nil))

	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))
	instance.erc1155Filterer = lo.Must(erc1155.NewERC1155Filterer(ethereum.AddressGenesis, nil))

	instance.nameWrapper = lo.Must(ens.NewNameWrapperCaller(ens.AddressNameWrapper, instance.ethereumClient))

	return &instance, nil
}
