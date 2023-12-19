package uniswap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/uniswap"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	ethereumClient ethereum.Client
	tokenClient    token.Client

	uniswapV1ExchangeFilterer                 *uniswap.V1ExchangeFilterer
	uniswapV2PairFilterer                     *uniswap.V2PairFilterer
	uniswapV3PoolFilterer                     *uniswap.V3PoolFilterer
	uniswapNonfungiblePositionManagerFilterer *uniswap.NonfungiblePositionManagerFilterer
}

func (w *worker) Name() string {
	return engine.Uniswap.String()
}

func (w *worker) Filter() engine.SourceFilter {
	return nil
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	_, ok := task.(*source.Task)

	return ok, nil
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	return nil, nil
}

func (w *worker) handleV1TokenPurchase(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseTokenPurchase(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TokenPurchase event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	token, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token from factory: %w", err)
	}

	if token == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial v1 exhcange: %s", event.Raw.Address)
	}

	return w.buildExchangeSwapAction(ctx, task, task.Transaction.From, event.Buyer, nil, &token, event.EthSold, event.TokensBought)
}

func (w *worker) handleV1EthPurchaseLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseEthPurchase(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse EthPurchase event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	token, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token from factory: %w", err)
	}

	if token == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial v1 exhcange: %s", event.Raw.Address)
	}

	return w.buildExchangeSwapAction(ctx, task, task.Transaction.From, event.Buyer, &token, nil, event.TokensSold, event.EthBought)
}

func (w *worker) handleV1ExchangeAddLiquidityLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseAddLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse AddLiquidity event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	tokenAddress, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token address from factory: %w", err)
	}

	if tokenAddress == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial exhcange contract: %s", event.Raw.Address)
	}

	tokens := []metadata.Token{
		{ // Native token
			Value: lo.ToPtr(decimal.NewFromBigInt(event.EthAmount, 0)),
		},
		{ // ERC-20 token
			Address: lo.ToPtr(tokenAddress.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.TokenAmount, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Provider, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
}

func (w *worker) handleV1ExchangeRemoveLiquidityLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseRemoveLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse RemoveLiquidity event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	token, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token address from factory: %w", err)
	}

	if token == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial exhcange contract: %s", event.Raw.Address)
	}

	tokens := []metadata.Token{
		{
			// Native token
			Value: lo.ToPtr(decimal.NewFromBigInt(event.EthAmount, 0)),
		},
		{
			// ERC-20 token
			Address: lo.ToPtr(token.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.TokenAmount, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, event.Provider, metadata.ActionExchangeLiquidityRemove, tokens)
}

func (w *worker) handleV2SwapLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV2PairFilterer.ParseSwap(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Swap event: %w", err)
	}

	pair, err := uniswap.NewV2PairCaller(event.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load pair: %w", err)
	}

	tokenLeft, err := pair.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 0 from pair: %w", err)
	}

	tokenRight, err := pair.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 1 from pair: %w", err)
	}

	sender, receiver := event.Sender, event.To

	if event.Sender == *task.Transaction.To && event.To == *task.Transaction.To {
		isNotFirstSwapEvent := lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
			return len(log.Topics) != 0 && contract.MatchEventHashes(log.Topics[0], uniswap.EventHashV2PairSwap) && log.Index < event.Raw.Index
		})

		if !isNotFirstSwapEvent {
			sender = task.Transaction.From
		}
	}

	// nolint:revive // False positive
	if event.Amount1Out.Sign() == 1 {
		return w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenLeft, &tokenRight, event.Amount0In, event.Amount1Out)
	} else {
		return w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenRight, &tokenLeft, event.Amount1In, event.Amount0Out)
	}
}

func (w *worker) handleV2PairMintLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV2PairFilterer.ParseMint(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Mint event: %w", err)
	}

	pair, err := uniswap.NewV2PairCaller(event.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load pair: %w", err)
	}

	tokenLeft, err := pair.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 0 from pair: %w", err)
	}

	tokenRight, err := pair.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 1 from pair: %w", err)
	}

	factory, err := uniswap.NewV2FactoryCaller(uniswap.AddressV2Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	pairAddress, err := factory.GetPair(nil, tokenLeft, tokenRight)
	if err != nil {
		return nil, fmt.Errorf("get pair address from factory: %w", err)
	}

	if pairAddress != event.Raw.Address {
		return nil, fmt.Errorf("unofficial pair contract: %s", event.Raw.Address)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(tokenLeft.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount0, 0)),
		},
		{
			Address: lo.ToPtr(tokenRight.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount1, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Sender, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
}

func (w *worker) handleV2PairBurnLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV2PairFilterer.ParseBurn(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Burn event: %w", err)
	}

	pair, err := uniswap.NewV2PairCaller(event.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load pair: %w", err)
	}

	tokenLeft, err := pair.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 0 from pair: %w", err)
	}

	tokenRight, err := pair.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 1 from pair: %w", err)
	}

	factory, err := uniswap.NewV2FactoryCaller(uniswap.AddressV2Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	pairAddress, err := factory.GetPair(nil, tokenLeft, tokenRight)
	if err != nil {
		return nil, fmt.Errorf("get pair address from factory: %w", err)
	}

	if pairAddress != event.Raw.Address {
		return nil, fmt.Errorf("unofficial pair contract: %s", event.Raw.Address)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(tokenLeft.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount0, 0)),
		},
		{
			Address: lo.ToPtr(tokenRight.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount1, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityRemove, tokens)
}

func (w *worker) handleEthereumV3SwapLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapV3PoolFilterer.ParseSwap(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Swap event: %w", err)
	}

	pool, err := uniswap.NewV3PoolCaller(event.Raw.Address, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load pool: %w", err)
	}

	tokenLeft, err := pool.Token0(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 0 from pool: %w", err)
	}

	tokenRight, err := pool.Token1(nil)
	if err != nil {
		return nil, fmt.Errorf("get token 1 from pool: %w", err)
	}

	fee, err := pool.Fee(nil)
	if err != nil {
		return nil, fmt.Errorf("get fee from pool: %w", err)
	}

	factory, err := uniswap.NewV3FactoryCaller(uniswap.AddressV3Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	poolAddress, err := factory.GetPool(nil, tokenLeft, tokenRight, fee)
	if err != nil {
		return nil, fmt.Errorf("get pool address from factory: %w", err)
	}

	if poolAddress != event.Raw.Address {
		return nil, fmt.Errorf("unofficial pool address: %s", event.Raw.Address)
	}

	sender, receiver := event.Sender, event.Recipient

	switch {
	case event.Sender == *task.Transaction.To && event.Recipient == *task.Transaction.To:
		isNotFirstSwapEvent := lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
			return len(log.Topics) != 0 && contract.MatchEventHashes(log.Topics[0], uniswap.EventHashV3PoolSwap) && log.Index < event.Raw.Index
		})

		if !isNotFirstSwapEvent {
			sender = task.Transaction.From
		}
	case event.Sender == *task.Transaction.To:
		sender = task.Transaction.From
	}

	// nolint:revive // False positive
	if event.Amount0.Sign() == 1 { // Amount0 is positive
		return w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenLeft, &tokenRight, event.Amount0, event.Amount1)
	} else {
		return w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenRight, &tokenLeft, event.Amount1, event.Amount0)
	}
}

func (w *worker) handleNonfungiblePositionManagerIncreaseLiquidityLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseIncreaseLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse IncreaseLiquidity event: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(uniswap.AddressNonfungiblePositionManager, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(nil, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount0, 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount1, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, task.Transaction.From, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
}

func (w *worker) handleNonfungiblePositionManagerDecreaseLiquidityLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseDecreaseLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DecreaseLiquidity event: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(uniswap.AddressNonfungiblePositionManager, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(nil, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount0, 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount1, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityRemove, tokens)
}

func (w *worker) handleNonfungiblePositionManagerCollectLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseCollect(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Collect event: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(uniswap.AddressNonfungiblePositionManager, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(nil, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount0, 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(event.Amount1, 0)),
		},
	}

	return w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityCollect, tokens)
}

func (w *worker) handleNonfungiblePositionManagerTransferLog(ctx context.Context, task source.Task, log *ethereum.Log) (*schema.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	return w.buildTransactionMintAction(ctx, task, event.From, event.To, event.Raw.Address, event.TokenId)
}

func (w *worker) buildExchangeSwapAction(ctx context.Context, task source.Task, sender, receipt common.Address, tokenIn, tokenOut *common.Address, amountIn, amountOut *big.Int) (*schema.Action, error) {
	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenIn, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenIn, err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenOut, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenOut, err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountOut, 0).Abs())

	action := schema.Action{
		Type:     filter.TypeExchangeSwap,
		Platform: filter.PlatformUniswap.String(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}

	return &action, nil
}

func (w *worker) buildExchangeLiquidityAction(ctx context.Context, task source.Task, sender, receipt common.Address, liquidityAction metadata.ExchangeLiquidityAction, tokens []metadata.Token) (*schema.Action, error) {
	tokenMetadataSlice := make([]metadata.Token, 0, len(tokens))

	for _, token := range tokens {
		var tokenAddress *common.Address
		if token.Address != nil {
			tokenAddress = lo.ToPtr(common.HexToAddress(*token.Address))
		}

		tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenAddress, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup token metadata %s: %w", tokenAddress, err)
		}

		tokenMetadata.Value = token.Value

		tokenMetadataSlice = append(tokenMetadataSlice, *tokenMetadata)
	}

	action := schema.Action{
		Type:     filter.TypeExchangeLiquidity,
		Platform: filter.PlatformUniswap.String(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: liquidityAction,
			Tokens: tokenMetadataSlice,
		},
	}

	return &action, nil
}

func (w *worker) buildTransactionMintAction(ctx context.Context, task source.Task, sender, receipt, tokenAddress common.Address, tokenID *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &tokenAddress, tokenID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %d: %w", tokenAddress, tokenID, err)
	}

	tokenMetadata.ID = lo.ToPtr(decimal.NewFromBigInt(tokenID, 0))
	tokenMetadata.Value = lo.ToPtr(decimal.NewFromInt(1))

	action := schema.Action{
		Type:     filter.TypeTransactionMint,
		Platform: filter.PlatformUniswap.String(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}

	return &action, nil
}

func NewWorker(config *engine.Config) (engine.Worker, error) {
	instance := worker{
		uniswapV1ExchangeFilterer:                 lo.Must(uniswap.NewV1ExchangeFilterer(ethereum.AddressGenesis, nil)),
		uniswapV2PairFilterer:                     lo.Must(uniswap.NewV2PairFilterer(ethereum.AddressGenesis, nil)),
		uniswapV3PoolFilterer:                     lo.Must(uniswap.NewV3PoolFilterer(ethereum.AddressGenesis, nil)),
		uniswapNonfungiblePositionManagerFilterer: lo.Must(uniswap.NewNonfungiblePositionManagerFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
