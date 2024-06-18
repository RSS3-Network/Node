package opensea

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/erc1155"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/contract/opensea"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// Worker is the worker for OpenSea.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                   *config.Module
	ethereumClient           ethereum.Client
	tokenClient              token.Client
	wyvernExchangeV1Filterer *opensea.WyvernExchangeV1Filterer
	wyvernExchangeV2Filterer *opensea.WyvernExchangeV2Filterer
	seaportV1Filterer        *opensea.SeaportV1Filterer
	erc20Filterer            *erc20.ERC20Filterer
	erc721Filterer           *erc721.ERC721Filterer
	erc1155Filterer          *erc1155.ERC1155Filterer
}

func (w *worker) Name() string {
	return decentralized.OpenSea.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformOpenSea.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Collectible,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.CollectibleTrade,
	}
}

// Filter opensea contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			opensea.AddressWyvernExchangeV1,
			opensea.AddressWyvernExchangeV2,
			opensea.AddressSeaportV1Dot1,
			opensea.AddressSeaportV1Dot2,
			opensea.AddressSeaportV1Dot3,
			opensea.AddressSeaportV1Dot4,
			opensea.AddressSeaportV1Dot5,
		},
		LogTopics: []common.Hash{
			opensea.EventHashWyvernExchangeV1OrdersMatched,
			opensea.EventHashWyvernExchangeV2OrdersMatched,
			opensea.EventHashSeaportV1OrderFulfilled,
		},
	}
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default opensea activity from task.
	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*activityx.Action
			err     error
		)

		// Match opensea core contract events
		switch {
		case w.matchWyvernExchangeV1Orders(ethereumTask, log):
			actions, err = w.transformWyvernExchangeV1Orders(ctx, ethereumTask, log)
		case w.matchWyvernExchangeV2Orders(ethereumTask, log):
			actions, err = w.transformWyvernExchangeV2Orders(ctx, ethereumTask, log)
		case w.matchSeaportV1OrderFulfilled(ethereumTask, log):
			actions, err = w.transformSeaportV1OrderFulfilled(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			if isInvalidTokenErr(err) {
				return activityx.NewUnknownActivity(activity), nil
			}

			return nil, err
		}

		// Change activity type to the first action type.
		for _, action := range actions {
			activity.Type = action.Type
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	return activity, nil
}

// matchWyvernExchangeV1Orders matches WyvernExchangeV1 OrdersMatched event.
func (w *worker) matchWyvernExchangeV1Orders(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == opensea.AddressWyvernExchangeV1 && contract.MatchEventHashes(log.Topics[0], opensea.EventHashWyvernExchangeV1OrdersMatched)
}

// matchWyvernExchangeV2Orders matches WyvernExchangeV2 OrdersMatched event.
func (w *worker) matchWyvernExchangeV2Orders(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == opensea.AddressWyvernExchangeV2 && contract.MatchEventHashes(log.Topics[0], opensea.EventHashWyvernExchangeV2OrdersMatched)
}

// matchSeaportV1OrderFulfilled matches SeaportV1 OrderFulfilled event.
func (w *worker) matchSeaportV1OrderFulfilled(_ *source.Task, log *ethereum.Log) bool {
	// Match SeaportV1 OrderFulfilled event.
	if !contract.MatchEventHashes(log.Topics[0], opensea.EventHashSeaportV1OrderFulfilled) {
		return false
	}

	// Match SeaportV1 contract address.
	return contract.MatchAddresses(
		log.Address,
		opensea.AddressSeaportV1Dot0, opensea.AddressSeaportV1Dot1, opensea.AddressSeaportV1Dot2,
		opensea.AddressSeaportV1Dot3, opensea.AddressSeaportV1Dot4, opensea.AddressSeaportV1Dot5,
	)
}

// transformWyvernExchangeV1Orders transforms WyvernExchangeV1 OrdersMatched event.
func (w *worker) transformWyvernExchangeV1Orders(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse OrdersMatched event.
	event, err := w.wyvernExchangeV1Filterer.ParseOrdersMatched(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse OrdersMatched event: %w", err)
	}

	var (
		nftAddress common.Address
		nftID      *big.Int
		nftValue   *big.Int
	)

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case log.Topics[0] == erc721.EventHashTransfer && len(log.Topics) == 4:
			// Parse erc721 transfer event.
			transferEvent, err := w.erc721Filterer.ParseTransfer(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse transfer event: %w", err)
			}

			// If the taker is not the recipient, swap the maker and taker.
			if event.Taker != transferEvent.To {
				event.Maker, event.Taker = event.Taker, event.Maker
			}

			nftAddress = transferEvent.Raw.Address
			nftID = transferEvent.TokenId
			nftValue = big.NewInt(1)
		case log.Topics[0] == erc1155.EventHashTransferSingle:
			// Parse erc1155 transfer event.
			transferEvent, err := w.erc1155Filterer.ParseTransferSingle(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse transfer event: %w", err)
			}

			// If the taker is not the recipient, swap the maker and taker.
			if event.Taker != transferEvent.To {
				event.Maker, event.Taker = event.Taker, event.Maker
			}

			nftAddress = transferEvent.Raw.Address
			nftID = transferEvent.Id
			nftValue = transferEvent.Value
		default:
			continue
		}
	}

	// Build collectible trade action.
	action, err := w.buildEthereumCollectibleTradeAction(ctx, task, event.Maker, event.Taker, nftAddress, nftID, nftValue, nil, event.Price)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		action,
	}, nil
}

func (w *worker) transformWyvernExchangeV2Orders(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	return w.transformWyvernExchangeV1Orders(ctx, task, log)
}

func (w *worker) transformSeaportV1OrderFulfilled(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse OrderFulfilled event.
	event, err := w.seaportV1Filterer.ParseOrderFulfilled(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse OrderFulfilled event: %w", err)
	}

	// The seller is the maker, and the buyer is the taker.
	seller, buyer := event.Offerer, event.Recipient

	var (
		offerToken      *common.Address
		offerValue      *big.Int
		nft             common.Address
		nftID, nftValue *big.Int
	)

	// Go through event.Offer to get the offer token.
	// handle different item types.
	for _, offer := range event.Offer {
		switch offer.ItemType {
		case opensea.ItemTypeNative:
			offerToken, offerValue = nil, offer.Amount
		case opensea.ItemTypeERC20:
			offerToken, offerValue = lo.ToPtr(offer.Token), offer.Amount
		case opensea.ItemTypeERC721, opensea.ItemTypeERC1155:
			nft, nftID, nftValue = offer.Token, offer.Identifier, offer.Amount
		}
	}

	// Go through event.Consideration to get the offer token.
	// handle different item types.
	for _, consideration := range event.Consideration {
		switch consideration.ItemType {
		case opensea.ItemTypeNative:
			// Ignore token transfers where receipt is not seller.
			if consideration.Recipient != seller {
				continue
			}

			offerToken, offerValue = nil, consideration.Amount
		case opensea.ItemTypeERC20:
			// Ignore token transfers where receipt is not seller.
			if consideration.Recipient != seller {
				continue
			}

			offerToken, offerValue = lo.ToPtr(consideration.Token), consideration.Amount
		case opensea.ItemTypeERC721, opensea.ItemTypeERC1155:
			// The offerer is the seller, and the recipient is the buyer.
			seller, buyer = buyer, seller

			nft, nftID, nftValue = consideration.Token, consideration.Identifier, consideration.Amount
		}
	}

	// Build collectible trade action.
	action, err := w.buildEthereumCollectibleTradeAction(ctx, task, seller, buyer, nft, nftID, nftValue, offerToken, offerValue)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		action,
	}, nil
}

func (w *worker) buildEthereumCollectibleTradeAction(ctx context.Context, task *source.Task, seller, buyer, nft common.Address, nftID, nftValue *big.Int, offerToken *common.Address, offerValue *big.Int) (*activityx.Action, error) {
	if nftID == nil {
		return nil, fmt.Errorf("nft id is nil")
	}

	nftValue = lo.Ternary(nftValue == nil, big.NewInt(1), nftValue)

	// Get collectible token metadata.
	collectibleTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup collectible token metadata: %w", err)
	}

	collectibleTokenMetadata.ID = lo.ToPtr(decimal.NewFromBigInt(nftID, 0))
	collectibleTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	// Get offer token metadata.
	costTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, offerToken, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup collectible token metadata: %w", err)
	}

	if offerValue == nil {
		offerValue = big.NewInt(0)
	}

	costTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))

	action := activityx.Action{
		Type:     typex.CollectibleTrade,
		Platform: w.Platform(),
		From:     seller.String(),
		To:       buyer.String(),
		Metadata: metadata.CollectibleTrade{
			Action: lo.If(task.Transaction.From == seller,
				metadata.ActionCollectibleTradeSell).Else(metadata.ActionCollectibleTradeBuy),
			Token: *collectibleTokenMetadata,
			Cost:  costTokenMetadata,
		},
	}

	return &action, nil
}

// Handle invalid token error.
func isInvalidTokenErr(err error) bool {
	return err != nil && strings.HasPrefix(err.Error(), "lookup collectible token metadata")
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
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize opensea filterers.
	instance.wyvernExchangeV1Filterer = lo.Must(opensea.NewWyvernExchangeV1Filterer(ethereum.AddressGenesis, nil))
	instance.wyvernExchangeV2Filterer = lo.Must(opensea.NewWyvernExchangeV2Filterer(ethereum.AddressGenesis, nil))
	instance.seaportV1Filterer = lo.Must(opensea.NewSeaportV1Filterer(ethereum.AddressGenesis, nil))
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))
	instance.erc1155Filterer = lo.Must(erc1155.NewERC1155Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
