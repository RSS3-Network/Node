package looksrare

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/looksrare"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// Worker is the worker for OpenSea.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config             *config.Module
	ethereumClient     ethereum.Client
	exchangeFilterer   *looksrare.ExchangeFilterer
	exchangeV2Filterer *looksrare.ExchangeV2Filterer
	aggregatorFilterer *looksrare.AggregatorFilterer
	seaportFilterer    *looksrare.SeaportFilterer
	erc20Filterer      *erc20.ERC20Filterer
	tokenClient        token.Client
}

func (w *worker) Name() string {
	return filter.Looksrare.String()
}

// Filter opensea contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			looksrare.AddressExchange,
			looksrare.AddressAggregator,
			looksrare.AddressFeeSharingSetter,
			looksrare.AddressExchangeV2,
			looksrare.AddressProtocolFeeRecipient,
		},
		LogTopics: []common.Hash{
			looksrare.EventHashRoyaltyPayment,
			looksrare.EventHashTakerBid,
			looksrare.EventHashTakerAsk,
			looksrare.EventHashTakerBidV2,
			looksrare.EventHashTakerAskV2,
			looksrare.EventHashAggregatorSweep,
			looksrare.EventHashOrderFulfilled,
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

	// Build default looksrare feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformLooksRare))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*schema.Action
			err     error
		)

		// Match looksrare core contract events
		switch {
		case w.matchExchangeAskMatched(ethereumTask, log):
			actions, err = w.transformExchangeAsk(ctx, ethereumTask, log)
		default:
			continue
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

// matchExchangeAskMatched matches AddressExchange TakerAsk event.
func (w *worker) matchExchangeAskMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerAsk)
}

//// matchExchangeBidMatched matches AddressExchange TakerBid event.
//func (w *worker) matchExchangeBidMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerBid)
//}
//
//// matcExchangeRoyaltyPaymentMatched matches AddressExchange RoyaltyPayment event.
//func (w *worker) matchExchangeRoyaltyPaymentMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashRoyaltyPayment)
//}
//
//// matchExchangeV2AskMatched matches AddressExchangeV2 TakerAsk event.
//func (w *worker) matchRoyaltyTransferMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == weth.AddressMainnet && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer)
//}
//
//// matchExchangeV2AskMatched matches AddressExchangeV2 TakerAsk event.
//func (w *worker) matchExchangeV2AskMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressExchangeV2 && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerAskV2)
//}
//
//// matchExchangeV2BidMatched matches AddressExchangeV2 TakerBid event.
//func (w *worker) matchExchangeV2BidMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressExchangeV2 && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerBidV2)
//}
//
//// matchAggregatedBidMatched matches AddressAggregator AggregatorSweep event.
//func (w *worker) matchAggregatedBidMatched(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressAggregator && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashAggregatorSweep)
//}
//
//// matchAggregatorOrderFulfilled matches AddressAggregator OrderFulfilled event.
//func (w *worker) matchAggregatorOrderFulfilled(_ *source.Task, log *ethereum.Log) bool {
//	return log.Address == looksrare.AddressAggregator && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashOrderFulfilled)
//}

// transformExchangeAsk transforms AddressExchange TakerAsk event.
func (w *worker) transformExchangeAsk(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse TakerAsk event.
	event, err := w.exchangeFilterer.ParseTakerAsk(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TakerAsk event: %w", err)
	}

	action, err := w.buildEthereumCollectibleTradeAction(ctx, task, event.Maker, event.Taker, event.Collection, metadata.ActionCollectibleTradeSell, event.TokenId, event.Amount, event.Price, &event.Currency)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		action,
	}, nil
}

func (w *worker) buildEthereumCollectibleTradeAction(ctx context.Context, task *source.Task, maker, taker, nft common.Address, action metadata.CollectibleTradeAction, nftID, nftValue, offerValue *big.Int, currency *common.Address) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	tradeToken := metadata.CollectibleTrade{
		Action: action,
		Token:  *tokenMetadata,
	}

	if offerValue != nil {
		offerTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, currency, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup token metadata: %w", err)
		}

		offerTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))
		tradeToken.Cost = offerTokenMetadata
	}

	var from, to string
	if action == metadata.ActionCollectibleTradeSell {
		from = taker.String()
		to = maker.String()
	} else if action == metadata.ActionCollectibleTradeBuy {
		from = maker.String()
		to = taker.String()
	}

	return &schema.Action{
		Type:     filter.TypeCollectibleTrade,
		Platform: filter.PlatformLooksRare.String(),
		From:     from,
		To:       to,
		Metadata: tradeToken,
	}, nil
}

// NewWorker creates a new OpenSea worker.
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

	// Initialize opensea filterers.
	instance.exchangeFilterer = lo.Must(looksrare.NewExchangeFilterer(ethereum.AddressGenesis, nil))
	instance.exchangeV2Filterer = lo.Must(looksrare.NewExchangeV2Filterer(ethereum.AddressGenesis, nil))
	instance.aggregatorFilterer = lo.Must(looksrare.NewAggregatorFilterer(ethereum.AddressGenesis, nil))
	instance.seaportFilterer = lo.Must(looksrare.NewSeaportFilterer(ethereum.AddressGenesis, nil))

	// Initialize erc20 filterer.
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
