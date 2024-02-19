package kiwistand

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/kiwistand"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/samber/lo"
)

// Worker is the worker for KiwiStand.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                  *config.Module
	ethereumClient          ethereum.Client
	tokenClient             token.Client
	protocolRewardsFilterer *kiwistand.ProtocolRewardsFilterer
	kiwiFilterer            *kiwistand.KiwiFilterer
}

func (w *worker) Name() string {
	return filter.Highlight.String()
}

// Filter kiwistand contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			kiwistand.AddressKIWI,
			kiwistand.AddressProtocolRewards,
		},
		LogTopics: []common.Hash{
			kiwistand.EventHashRewardsDeposit,
			kiwistand.EventHashTransfer,
			kiwistand.EventHashSale,
		},
	}
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.GetNetwork().Source() == filter.NetworkEthereumSource, nil
}

// Transform Ethereum task to feed.
func (w *worker) Transform(_ context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default kiwistand feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformKiwiStand))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	//for _, log := range ethereumTask.Receipt.Logs {
	//	var (
	//		actions []*schema.Action
	//		err     error
	//	)
	//
	//	// Match kiwistand core contract events
	//	switch {
	//	case w.matchNumTokenMintMatched(ethereumTask, log):
	//		feed.Type = filter.TypeCollectibleMint
	//		actions, err = w.transformKiwiMint(ctx, ethereumTask, log)
	//	default:
	//		continue
	//	}
	//
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// Change feed type to the first action type.
	//	for _, action := range actions {
	//		feed.Type = action.Type
	//	}
	//
	//	feed.Actions = append(feed.Actions, actions...)
	//}

	return feed, nil
}

// NewWorker creates a new KiwiStand worker.
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

	// Initialize kiwistand filterers.
	instance.kiwiFilterer = lo.Must(kiwistand.NewKiwiFilterer(ethereum.AddressGenesis, nil))
	instance.protocolRewardsFilterer = lo.Must(kiwistand.NewProtocolRewardsFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
