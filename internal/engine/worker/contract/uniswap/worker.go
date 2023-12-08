package uniswap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
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

func (w *worker) Match(ctx context.Context, task engine.Task) (bool, error) {
	return false, nil
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

	return w.buildEthereumExchangeLiquidityAction(ctx, task, event.Provider, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
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

	return w.buildEthereumExchangeLiquidityAction(ctx, task, event.Raw.Address, event.Provider, metadata.ActionExchangeLiquidityRemove, tokens)
}

func (w *worker) buildExchangeSwapAction(ctx context.Context, task source.Task, sender, receipt common.Address, tokenIn, tokenOut *common.Address, amountIn, amountOut *big.Int) (*schema.Action, error) {
	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.Chain, tokenIn, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenIn, err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.Chain, tokenOut, nil, task.Header.Number)
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

func (w *worker) buildEthereumExchangeLiquidityAction(ctx context.Context, task source.Task, sender, receipt common.Address, liquidityAction metadata.ExchangeLiquidityAction, tokens []metadata.Token) (*schema.Action, error) {
	tokenMetadataSlice := make([]metadata.Token, 0, len(tokens))

	for _, token := range tokens {
		var tokenAddress *common.Address
		if token.Address != nil {
			tokenAddress = lo.ToPtr(common.HexToAddress(*token.Address))
		}

		tokenMetadata, err := w.tokenClient.Lookup(ctx, task.Chain, tokenAddress, nil, task.Header.Number)
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
