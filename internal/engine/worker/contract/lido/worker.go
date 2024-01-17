package lido

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/lido"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
)

// Worker is the worker for Lido.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	// Base
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	// Specific Filters
	stakedETHFilterer              *lido.StakedETHFilterer
	wrappedStakedETHFilterer       *lido.WrappedStakedETHFilterer
	stakedETHWithdrawalNFTFilterer *lido.StakedETHWithdrawalNFTFilterer
	stakedMATICFilterer            *lido.StakedMATICFilterer
	// Common filters
	erc721Filterer *erc721.ERC721Filterer
}

func (w *worker) Name() string {
	return filter.Lido.String()
}

// Filter lido contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			lido.AddressStakedETH,
			lido.AddressWrappedStakedETH,
			lido.AddressStakedETHWithdrawalNFT,
			lido.AddressStakedMATIC,
			lido.AddressStakedMATICWithdrawalNFT,
			lido.AddressMATIC,
		},
		LogTopics: []common.Hash{
			lido.EventHashStakedETHSubmitted,
			lido.EventHashStakedETHWithdrawalNFTWithdrawalRequested,
			lido.EventHashStakedETHWithdrawalNFTWithdrawalClaimed,
			lido.EventHashStakedMATICSubmitEvent,
			lido.EventHashStakedMATICRequestWithdrawEvent,
			lido.EventHashStakedMATICClaimTokensEvent,
			lido.EventHashStakedETHTransferShares,
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

	// Build default lido feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformLido))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*schema.Action
			err     error
		)

		// Match lido core contract events
		switch {
		case w.matchStakedETHSubmittedLog(ethereumTask, log):
			// Add ETH liquidity
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedETHSubmittedLog(ctx, ethereumTask, log)
		case w.matchStakedETHWithdrawalNFTWithdrawalRequestedLog(ethereumTask, log):
			// Mint ETH withdrawal NFT
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedETHWithdrawalNFTWithdrawalRequestedLog(ctx, ethereumTask, log)
		case w.matchStakedETHWithdrawalNFTWithdrawalClaimedLog(ethereumTask, log):
			// Remove ETH liquidity
			feed.Type = filter.TypeCollectibleBurn
			actions, err = w.transformStakedETHWithdrawalNFTWithdrawalClaimedLog(ctx, ethereumTask, log)
		case w.matchStakedMATICSubmitEventLog(ethereumTask, log):
			// Add MATIC liquidity
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedMATICSubmitEventLog(ctx, ethereumTask, log)
		case w.matchStakedMATICRequestWithdrawEventLog(ethereumTask, log):
			// Mint MATIC withdrawal NFT
			feed.Type = filter.TypeCollectibleMint
			actions, err = w.transformStakedMATICRequestWithdrawEventLog(ctx, ethereumTask, log)
		case w.matchStakedMATICClaimTokensEventLog(ethereumTask, log):
			// Remove MATIC liquidity
			feed.Type = filter.TypeCollectibleBurn
			actions, err = w.transformStakedMATICClaimTokensEventLog(ctx, ethereumTask, log)
		case w.matchStakedETHTransferSharesLog(ethereumTask, log):
			// Wrap or unwrap wstETH
			feed.Type = filter.TypeExchangeSwap
			actions, err = w.transformStakedETHTransferSharesLog(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// NewWorker creates a new Lido worker.
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

	// Initialize lido filterers.
	instance.stakedETHFilterer = lo.Must(lido.NewStakedETHFilterer(ethereum.AddressGenesis, nil))
	instance.wrappedStakedETHFilterer = lo.Must(lido.NewWrappedStakedETHFilterer(ethereum.AddressGenesis, nil))
	instance.stakedETHWithdrawalNFTFilterer = lo.Must(lido.NewStakedETHWithdrawalNFTFilterer(ethereum.AddressGenesis, nil))
	instance.stakedMATICFilterer = lo.Must(lido.NewStakedMATICFilterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
