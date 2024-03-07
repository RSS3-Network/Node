package looksrare

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/looksrare"
	"github.com/rss3-network/node/provider/ethereum/contract/weth"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
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
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*schema.Action
			err     error
		)

		// Match looksrare core contract events
		switch {
		case w.matchExchangeAskMatched(ethereumTask, log):
			actions, err = w.transformExchangeAsk(ctx, ethereumTask, log)
		case w.matchExchangeBidMatched(ethereumTask, log):
			actions, err = w.transformExchangeBid(ctx, ethereumTask, log)
		case w.matchExchangeRoyaltyPaymentMatched(ethereumTask, log):
			actions, err = w.transformExchangeRoyaltyPayment(ctx, ethereumTask, log)
		case w.matchRoyaltyTransferMatched(ethereumTask, log):
			actions, err = w.transformRoyaltyTransfer(ctx, ethereumTask, log)
		case w.matchExchangeV2AskMatched(ethereumTask, log):
			actions, err = w.transformExchangeV2Ask(ctx, ethereumTask, log)
		case w.matchExchangeV2BidMatched(ethereumTask, log):
			actions, err = w.transformExchangeV2Bid(ctx, ethereumTask, log)
		case w.matchAggregatedBidMatched(ethereumTask, log):
			actions, err = w.transformV2AggregatedBid(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		feed.Type = filter.TypeCollectibleTrade
		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// matchExchangeAskMatched matches AddressExchange TakerAsk event.
func (w *worker) matchExchangeAskMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerAsk)
}

// matchExchangeBidMatched matches AddressExchange TakerBid event.
func (w *worker) matchExchangeBidMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerBid)
}

// matcExchangeRoyaltyPaymentMatched matches AddressExchange RoyaltyPayment event.
func (w *worker) matchExchangeRoyaltyPaymentMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchange && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashRoyaltyPayment)
}

// matchExchangeRoyaltyTransferMatched matches AddressExchangeV2 TakerAsk event.
func (w *worker) matchRoyaltyTransferMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == weth.AddressMainnet && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer)
}

// matchExchangeV2AskMatched matches AddressExchangeV2 TakerAsk event.
func (w *worker) matchExchangeV2AskMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchangeV2 && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerAskV2)
}

// matchExchangeV2BidMatched matches AddressExchangeV2 TakerBid event.
func (w *worker) matchExchangeV2BidMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressExchangeV2 && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashTakerBidV2)
}

// matchAggregatedBidMatched matches AddressAggregator AggregatorSweep event.
func (w *worker) matchAggregatedBidMatched(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == looksrare.AddressAggregator && contract.MatchEventHashes(log.Topics[0], looksrare.EventHashAggregatorSweep)
}

// matchEthereumAggregatorOrderFulfilled matches AddressAggregator OrderFulfilled event.
func (w *worker) matchAggregatorOrderFulfilled(_ *source.Task, log ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], looksrare.EventHashOrderFulfilled)
}

// transformExchangeAsk transforms AddressExchange TakerAsk event.
func (w *worker) transformExchangeAsk(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse TakerAsk event.
	event, err := w.exchangeFilterer.ParseTakerAsk(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TakerAsk event: %w", err)
	}

	action, err := w.buildCollectibleTradeAction(ctx, task, event.Maker, event.Taker, event.Collection, metadata.ActionCollectibleTradeSell, event.TokenId, event.Amount, event.Price, &event.Currency)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		action,
	}, nil
}

// transformExchangeBid transforms AddressExchange TakerBid event.
func (w *worker) transformExchangeBid(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse TakerAsk event.
	event, err := w.exchangeFilterer.ParseTakerBid(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TakerAsk event: %w", err)
	}

	to := event.Taker
	if event.Taker == looksrare.AddressAggregator {
		to = task.Transaction.From
	}

	action, err := w.buildCollectibleTradeAction(ctx, task, event.Maker, to, event.Collection, metadata.ActionCollectibleTradeBuy, event.TokenId, event.Amount, event.Price, &event.Currency)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		action,
	}, nil
}

// transformExchangeRoyaltyPayment transforms AddressExchange RoyaltyPayment event.
func (w *worker) transformExchangeRoyaltyPayment(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.exchangeFilterer.ParseRoyaltyPayment(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse RoyaltyPayment event: %w", err)
	}

	action, err := w.buildRoyaltyPaymentAction(ctx, task, event.Currency, event.Amount, event.RoyaltyRecipient)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		action,
	}, nil
}

// transformRoyaltyTransfer transforms AddressExchangeV2 TakerAsk event.
func (w *worker) transformRoyaltyTransfer(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.erc20Filterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	if event.To != looksrare.AddressFeeSharingSetter {
		return nil, nil
	}

	action, err := w.buildRoyaltyTransferAction(ctx, task, event.To, event.Value)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		action,
	}, nil
}

// transformExchangeV2Ask transforms AddressExchangeV2 TakerAsk event.
func (w *worker) transformExchangeV2Ask(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse TakerAsk event.
	event, err := w.exchangeV2Filterer.ParseTakerAsk(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TakerAsk event: %w", err)
	}

	feeAmount := big.NewInt(0)
	for _, feeAmounts := range event.FeeAmounts {
		feeAmount = feeAmount.Add(feeAmount, feeAmounts)
	}

	actionAsk, err := w.buildCollectibleTradeAction(ctx, task, event.BidUser, event.AskUser, event.Collection, metadata.ActionCollectibleTradeSell, event.ItemIds[0], event.Amounts[0], feeAmount, &event.Currency)
	if err != nil {
		return nil, err
	}

	actions := []*schema.Action{
		actionAsk,
	}

	var creatorFee *schema.Action
	if event.FeeAmounts[1].Cmp(big.NewInt(0)) != 0 {
		creatorFee, err = w.buildV2RoyaltyFeeAction(ctx, task, event.BidUser, event.FeeRecipients[1], event.FeeAmounts[1], &event.Currency)
		if err != nil {
			return nil, err
		}

		actions = append(actions, creatorFee)
	}

	var actionRoyaltyFee *schema.Action
	if event.FeeAmounts[2].Cmp(big.NewInt(0)) != 0 {
		actionRoyaltyFee, err = w.buildV2RoyaltyFeeAction(ctx, task, event.BidUser, looksrare.AddressProtocolFeeRecipient, event.FeeAmounts[2], &event.Currency)
		if err != nil {
			return nil, err
		}

		actions = append(actions, actionRoyaltyFee)
	}

	return actions, nil
}

// transformExchangeV2Bid transforms AddressExchangeV2 TakerBid event.
func (w *worker) transformExchangeV2Bid(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse TakerBid event.
	event, err := w.exchangeV2Filterer.ParseTakerBid(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TakerBid event: %w", err)
	}

	feeAmount := big.NewInt(0)
	for _, feeAmounts := range event.FeeAmounts {
		feeAmount = feeAmount.Add(feeAmount, feeAmounts)
	}

	actionBid, err := w.buildCollectibleTradeAction(ctx, task, event.FeeRecipients[0], event.BidRecipient, event.Collection, metadata.ActionCollectibleTradeBuy, event.ItemIds[0], event.Amounts[0], feeAmount, &event.Currency)
	if err != nil {
		return nil, err
	}

	actions := []*schema.Action{
		actionBid,
	}

	var creatorFee *schema.Action
	if event.FeeAmounts[1].Cmp(big.NewInt(0)) != 0 {
		creatorFee, err = w.buildV2RoyaltyFeeAction(ctx, task, event.BidUser, event.FeeRecipients[1], event.FeeAmounts[1], &event.Currency)
		if err != nil {
			return nil, err
		}

		actions = append(actions, creatorFee)
	}

	var actionRoyaltyFee *schema.Action
	if event.FeeAmounts[2].Cmp(big.NewInt(0)) != 0 {
		actionRoyaltyFee, err = w.buildV2RoyaltyFeeAction(ctx, task, event.BidUser, looksrare.AddressProtocolFeeRecipient, event.FeeAmounts[2], &event.Currency)
		if err != nil {
			return nil, err
		}

		actions = append(actions, actionRoyaltyFee)
	}

	return actions, nil
}

// transformV2AggregatedBid transforms AddressAggregator AggregatorSweep event.
func (w *worker) transformV2AggregatedBid(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	// Parse AggregatorSweep event.
	event, err := w.aggregatorFilterer.ParseSweep(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Sweep event: %w", err)
	}

	var orderEvents []looksrare.SeaportOrderFulfilled

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchAggregatorOrderFulfilled(task, *log):
			orderEvent, err := w.seaportFilterer.ParseOrderFulfilled(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse OrderFulfilled event: %w", err)
			}

			orderEvents = append(orderEvents, *orderEvent)
		default:
			continue
		}

		if err != nil {
			zap.L().Debug("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			continue
		}
	}

	var actions []*schema.Action

	var tokenAddress common.Address

	for _, orderEvent := range orderEvents {
		cost := big.NewInt(0)

		for _, fee := range orderEvent.Consideration {
			fee := fee

			if event.Sweeper == fee.Recipient {
				continue
			}

			actionFee, err := w.buildV2RoyaltyFeeAction(ctx, task, event.Sweeper, fee.Recipient, fee.Amount, &fee.Token)
			if err != nil {
				return nil, err
			}

			tokenAddress = fee.Token
			cost = cost.Add(cost, fee.Amount)

			actions = append(actions, actionFee)
		}

		for _, offer := range orderEvent.Offer {
			collection := offer.Token
			tokenID := offer.Identifier
			amount := offer.Amount

			actionBid, err := w.buildCollectibleTradeAction(ctx, task, orderEvent.Offerer, orderEvent.Recipient, collection, metadata.ActionCollectibleTradeBuy, tokenID, amount, cost, &tokenAddress)
			if err != nil {
				return nil, err
			}

			actions = append(actions, actionBid)
		}
	}

	return actions, nil
}

// buildEthereumCollectibleTradeAction builds collectible trade action.
func (w *worker) buildCollectibleTradeAction(ctx context.Context, task *source.Task, maker, taker, nft common.Address, action metadata.CollectibleTradeAction, nftID, nftValue, offerValue *big.Int, currency *common.Address) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	tradeToken := metadata.CollectibleTrade{
		Action: action,
		Token:  *tokenMetadata,
	}

	if currency.String() == "0x0000000000000000000000000000000000000000" {
		currency = nil
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

// buildRoyaltyPaymentAction builds royalty payment action.
func (w *worker) buildRoyaltyPaymentAction(ctx context.Context, task *source.Task, currency common.Address, amount *big.Int, receipt common.Address) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &currency, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	var from common.Address

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchExchangeAskMatched(task, log):
			event, err := w.exchangeFilterer.ParseTakerAsk(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse taker ask event: %w", err)
			}

			from = event.Maker
		case w.matchExchangeBidMatched(task, log): // Deposit and withdraw ETH
			event, err := w.exchangeFilterer.ParseTakerBid(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse taker bid event: %w", err)
			}

			from = event.Maker
		default:
			continue
		}

		if err != nil {
			zap.L().Debug("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			continue
		}
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amount, 0))

	return &schema.Action{
		Type:     filter.TypeTransactionTransfer,
		Platform: filter.PlatformLooksRare.String(),
		From:     from.String(),
		To:       receipt.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

// buildRoyaltyTransferAction builds royalty transfer action.
func (w *worker) buildRoyaltyTransferAction(ctx context.Context, task *source.Task, to common.Address, wad *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &weth.AddressMainnet, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	var from common.Address

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchExchangeAskMatched(task, log):
			event, err := w.exchangeFilterer.ParseTakerAsk(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse taker ask event: %w", err)
			}

			from = event.Maker
		case w.matchExchangeBidMatched(task, log): // Deposit and withdraw ETH
			event, err := w.exchangeFilterer.ParseTakerBid(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse taker bid event: %w", err)
			}

			from = event.Maker
		default:
			continue
		}

		if err != nil {
			zap.L().Debug("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			continue
		}
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(wad, 0))

	return &schema.Action{
		Type:     filter.TypeTransactionTransfer,
		Platform: filter.PlatformLooksRare.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

// buildV2RoyaltyFeeAction builds royalty fee action.
func (w *worker) buildV2RoyaltyFeeAction(ctx context.Context, task *source.Task, from common.Address, to common.Address, amount *big.Int, currency *common.Address) (*schema.Action, error) {
	if currency.String() == "0x0000000000000000000000000000000000000000" {
		currency = nil
	}

	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, currency, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amount, 0))

	return &schema.Action{
		Type:     filter.TypeTransactionTransfer,
		Platform: filter.PlatformLooksRare.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
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
