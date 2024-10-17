package benddao

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/benddao"
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

var _ engine.Worker = (*worker)(nil)

type worker struct {
	tokenClient          token.Client
	bendExchangeFilterer *benddao.BendExchangeFilterer
	lendPoolFilterer     *benddao.LendPoolFilterer
}

func (w *worker) Name() string {
	return decentralized.BendDAO.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformBendDAO.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Collectible,
		tag.Exchange,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeLiquidity,
		typex.ExchangeLoan,
		typex.CollectibleAuction,
		typex.CollectibleTrade,
		typex.CollectibleTransfer,
	}
}

// Filter contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			benddao.AddressBendExchange,
			benddao.AddressLendPool,
		},
		LogTopics: []common.Hash{
			benddao.EventTakerAsk,
			benddao.EventTakerBid,
			benddao.EventDeposit,
			benddao.EventWithdraw,
			benddao.EventBorrow,
			benddao.EventRepay,
			benddao.EventAuction,
			benddao.EventRedeem,
			benddao.EventLiquidate,
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
		return nil, fmt.Errorf("build activity: %w", err)
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
		// Bend exchange
		case w.matchBendExchangeTakerAsk(log):
			actions, err = w.transformBendExchangeTakerAsk(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleTrade
		case w.matchBendExchangeTakerBid(log):
			actions, err = w.transformBendExchangeTakerBid(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleTrade

		// Lend pool
		case w.matchLendPoolDeposit(log):
			actions, err = w.transformLendPoolDeposit(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLiquidity
		case w.matchLendPoolWithdraw(log):
			actions, err = w.transformLendPoolWithdraw(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLiquidity
		case w.matchLendPoolBorrow(log):
			actions, err = w.transformLendPoolBorrow(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLoan
		case w.matchLendPoolRepay(log):
			actions, err = w.transformLendPoolRepay(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLoan
		case w.matchLendPoolAuction(log):
			actions, err = w.transformLendPoolAuction(ctx, ethereumTask, log)
			activity.Type = typex.CollectibleAuction
		case w.matchLendPoolRedeem(log):
			actions, err = w.transformLendPoolRedeem(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLoan
		case w.matchLendPoolLiquidate(log):
			actions, err = w.transformLendPoolLiquidate(ctx, ethereumTask, log)
			activity.Type = typex.ExchangeLoan
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	return activity, nil
}

// NewWorker creates a new worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	ethereumClient, err := ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...)
	if err != nil {
		return nil, fmt.Errorf("dial Ethereum: %w", err)
	}

	return &worker{
		tokenClient:          token.NewClient(ethereumClient),
		bendExchangeFilterer: lo.Must(benddao.NewBendExchangeFilterer(benddao.AddressBendExchange, nil)),
		lendPoolFilterer:     lo.Must(benddao.NewLendPoolFilterer(benddao.AddressLendPool, nil)),
	}, nil
}

func (w *worker) matchBendExchangeTakerAsk(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressBendExchange) && contract.MatchEventHashes(log.Topics[0], benddao.EventTakerAsk)
}

func (w *worker) matchBendExchangeTakerBid(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressBendExchange) && contract.MatchEventHashes(log.Topics[0], benddao.EventTakerBid)
}

func (w *worker) matchLendPoolDeposit(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventDeposit)
}

func (w *worker) matchLendPoolWithdraw(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventWithdraw)
}

func (w *worker) matchLendPoolBorrow(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventBorrow)
}

func (w *worker) matchLendPoolRepay(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventRepay)
}

func (w *worker) matchLendPoolAuction(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventAuction)
}

func (w *worker) matchLendPoolRedeem(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventRedeem)
}

func (w *worker) matchLendPoolLiquidate(log *ethereum.Log) bool {
	return contract.MatchAddresses(log.Address, benddao.AddressLendPool) && contract.MatchEventHashes(log.Topics[0], benddao.EventLiquidate)
}

func (w *worker) transformBendExchangeTakerAsk(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.bendExchangeFilterer.ParseTakerAsk(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse takerAsk event: %w", err)
	}

	action, err := w.buildEthereumCollectibleTradeAction(ctx, task, event.Taker, event.Maker, event.Collection, event.TokenId, event.Amount, &event.Currency, event.Price)
	if err != nil {
		return nil, fmt.Errorf("build takerAsk action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformBendExchangeTakerBid(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.bendExchangeFilterer.ParseTakerBid(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse takerBid event: %w", err)
	}

	action, err := w.buildEthereumCollectibleTradeAction(ctx, task, event.Maker, event.Taker, event.Collection, event.TokenId, event.Amount, &event.Currency, event.Price)
	if err != nil {
		return nil, fmt.Errorf("build takerBid action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolDeposit(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse deposit event: %w", err)
	}

	action, err := w.buildEthereumLendPoolLiquidityAction(ctx, task, event.OnBehalfOf, log.Address, metadata.ActionExchangeLiquiditySupply, []metadata.Token{
		{
			Address: lo.ToPtr(event.Reserve.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount, 0)),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolWithdraw(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseWithdraw(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse withdraw event: %w", err)
	}

	action, err := w.buildEthereumLendPoolLiquidityAction(ctx, task, log.Address, event.To, metadata.ActionExchangeLiquidityWithdraw, []metadata.Token{
		{
			Address: lo.ToPtr(event.Reserve.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount, 0)),
		},
	})
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolBorrow(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseBorrow(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse borrow event: %w", err)
	}

	action, err := w.buildEthereumExchangeLoanAction(ctx, task, log.Address, event.OnBehalfOf, event.NftAsset, lo.ToPtr(event.Reserve), event.NftTokenId, big.NewInt(1), event.Amount, metadata.ActionExchangeLoanCreate)
	if err != nil {
		return nil, fmt.Errorf("build exchange loan action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolRepay(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseRepay(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse repay event: %w", err)
	}

	action, err := w.buildEthereumExchangeLoanAction(ctx, task, log.Address, event.Borrower, event.NftAsset, lo.ToPtr(event.Reserve), event.NftTokenId, big.NewInt(1), event.Amount, metadata.ActionExchangeLoanRepay)
	if err != nil {
		return nil, fmt.Errorf("build exchange loan action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolAuction(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseAuction(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse auction event: %w", err)
	}

	action, err := w.buildEthereumAuctionAction(ctx, task, event.OnBehalfOf, event.User, event.NftAsset, event.NftTokenId, big.NewInt(1), event.BidPrice, &event.Reserve, metadata.ActionCollectibleAuctionBid)
	if err != nil {
		return nil, fmt.Errorf("build collectible auction action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) transformLendPoolRedeem(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseRedeem(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse redeem event: %w", err)
	}

	actionTransfer, err := w.buildEthereumTransactionTransferAction(ctx, task, event.Borrower, log.Address, lo.ToPtr(event.Reserve), event.FineAmount)
	if err != nil {
		return nil, fmt.Errorf("build transfer action: %w", err)
	}

	actionLoan, err := w.buildEthereumExchangeLoanAction(ctx, task, log.Address, event.Borrower, event.NftAsset, lo.ToPtr(event.Reserve), event.NftTokenId, big.NewInt(1), event.BorrowAmount, metadata.ActionExchangeLoanRepay)
	if err != nil {
		return nil, fmt.Errorf("build exchange loan action: %w", err)
	}

	return []*activityx.Action{actionTransfer, actionLoan}, nil
}

func (w *worker) transformLendPoolLiquidate(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.lendPoolFilterer.ParseLiquidate(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse redeem event: %w", err)
	}

	actionTransfer, err := w.buildEthereumTransactionTransferAction(ctx, task, task.Transaction.From, event.Borrower, &event.Reserve, event.RemainAmount)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	actionLoan, err := w.buildEthereumExchangeLoanAction(ctx, task, log.Address, event.Borrower, event.NftAsset, lo.ToPtr(event.Reserve), event.NftTokenId, big.NewInt(1), event.RepayAmount, metadata.ActionExchangeLoanLiquidate)
	if err != nil {
		return nil, fmt.Errorf("build exchange loan action: %w", err)
	}

	actionCTransfer, err := w.buildEthereumCollectibleTransferAction(ctx, task, log.Address, task.Transaction.From, event.NftAsset, event.NftTokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	return []*activityx.Action{actionTransfer, actionLoan, actionCTransfer}, nil
}

func (w *worker) buildEthereumLendPoolLiquidityAction(ctx context.Context, task *source.Task, sender, receipt common.Address, action metadata.ExchangeLiquidityAction, tokens []metadata.Token) (*activityx.Action, error) {
	tokenMetadataSlice := make([]metadata.Token, 0, len(tokens))

	for _, t := range tokens {
		var tokenAddress *common.Address
		if t.Address != nil {
			tokenAddress = lo.ToPtr(common.HexToAddress(*t.Address))
		}

		tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenAddress, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup collectible token metadata: %w, chain: %v, token: %v", err, task.ChainID, tokenAddress)
		}

		tokenMetadata.Value = t.Value

		tokenMetadataSlice = append(tokenMetadataSlice, *tokenMetadata)
	}

	return &activityx.Action{
		Type:     typex.ExchangeLiquidity,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: action,
			Tokens: tokenMetadataSlice,
		},
	}, nil
}

func (w *worker) buildEthereumAuctionAction(ctx context.Context, task *source.Task, from, to, nft common.Address, nftID, nftValue, offerValue *big.Int, offerToken *common.Address, action metadata.CollectibleAuctionAction) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", nft, nftID, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	var (
		offerTokenMetadata *metadata.Token
	)

	if offerValue != nil {
		offerTokenMetadata, err = w.tokenClient.Lookup(ctx, task.ChainID, offerToken, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup token metadata %s: %w", offerToken, err)
		}

		offerTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))
	}

	return &activityx.Action{
		Type:     typex.CollectibleAuction,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleAuction{
			Action: action,
			Token:  *tokenMetadata,
			Cost:   offerTokenMetadata,
		},
	}, nil
}

func (w *worker) buildEthereumCollectibleTradeAction(ctx context.Context, task *source.Task, seller, buyer, nft common.Address, nftID, nftValue *big.Int, offerToken *common.Address, offerValue *big.Int) (*activityx.Action, error) {
	offerTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, offerToken, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", offerToken, err)
	}

	offerTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))

	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", nft, nftID, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	return &activityx.Action{
		Type:     typex.CollectibleTrade,
		Platform: w.Platform(),
		From:     seller.String(),
		To:       buyer.String(),
		Metadata: metadata.CollectibleTrade{
			Action: lo.Ternary(task.Transaction.From == seller, metadata.ActionCollectibleTradeSell, metadata.ActionCollectibleTradeBuy),
			Token:  *tokenMetadata,
			Cost:   offerTokenMetadata,
		},
	}, nil
}

func (w *worker) buildEthereumTransactionTransferAction(ctx context.Context, task *source.Task, from, to common.Address, token *common.Address, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, token, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", token, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	return &activityx.Action{
		Type:     typex.TransactionTransfer,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

func (w *worker) buildEthereumExchangeLoanAction(ctx context.Context, task *source.Task, lender, borrower, nft common.Address, offerToken *common.Address, nftID, nftValue, offerValue *big.Int, action metadata.ExchangeLoanAction) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", nft, nftID, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(nftValue, 0))

	var (
		offerTokenMetadata *metadata.Token
	)

	if offerValue != nil {
		offerTokenMetadata, err = w.tokenClient.Lookup(ctx, task.ChainID, offerToken, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup token metadata %s: %w", "", err)
		}

		offerTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(offerValue, 0))
	}

	return &activityx.Action{
		Type:     typex.ExchangeLoan,
		Platform: w.Platform(),
		From:     lender.String(),
		To:       borrower.String(),
		Metadata: metadata.ExchangeLoan{
			Action:     action,
			Collateral: *tokenMetadata,
			Amount:     offerTokenMetadata,
		},
	}, nil
}

func (w *worker) buildEthereumCollectibleTransferAction(ctx context.Context, task *source.Task, seller, buyer, nft common.Address, nftID, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &nft, nftID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", nft, nftID, err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	return &activityx.Action{
		Type:     typex.CollectibleTransfer,
		Platform: w.Platform(),
		From:     seller.String(),
		To:       buyer.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}, nil
}
