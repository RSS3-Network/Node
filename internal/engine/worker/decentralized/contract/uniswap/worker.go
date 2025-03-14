package uniswap

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/uniswap"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/weth"
	"github.com/rss3-network/node/v2/provider/ethereum/token"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	ethereumClient ethereum.Client
	tokenClient    token.Client

	uniswapV1ExchangeFilterer                 *uniswap.V1ExchangeFilterer
	uniswapV2PairFilterer                     *uniswap.V2PairFilterer
	uniswapV3PoolFilterer                     *uniswap.V3PoolFilterer
	uniswapNonfungiblePositionManagerFilterer *uniswap.NonfungiblePositionManagerFilterer
	wethWETH9Filterer                         *weth.WETH9Filterer
}

func (w *worker) Name() string {
	return decentralized.Uniswap.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformUniswap.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Arbitrum,
		network.Optimism,
		network.Polygon,
		network.Base,
		network.BinanceSmartChain,
		network.Avalanche,
		network.Linea,
		network.SatoshiVM,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
		tag.Transaction,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeLiquidity,
		typex.ExchangeSwap,
		typex.TransactionMint,
	}
}

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		// LogAddresses: []common.Address{
		// 	uniswap.AddressV1Factory,
		// 	uniswap.AddressV2Migrator,
		// 	uniswap.AddressV2SwapRouter01,
		// 	uniswap.AddressV2SwapRouter02,
		// 	uniswap.AddressV3Migrator,
		// 	uniswap.AddressV3SwapRouter,
		// 	uniswap.AddressV3SwapRouterCelo,
		// 	uniswap.AddressV3SwapRouter02,
		// 	uniswap.AddressV3SwapRouter02Celo,
		// 	uniswap.AddressV3SwapRouter02BinanceSmartChain,
		// 	uniswap.AddressV2Factory,
		// 	uniswap.AddressV3Factory,
		// 	uniswap.AddressV3FactoryCelo,
		// 	uniswap.AddressV3FactoryBinanceSmartChain,
		// 	uniswap.AddressUniversalRouter,
		// },
		LogTopics: []common.Hash{
			uniswap.EventHashV3PoolSwap,
			uniswap.EventHashV3FactoryPoolCreated,
			uniswap.EventHashV1ExchangeTokenPurchase,
			uniswap.EventHashV1ExchangeEthPurchase,
			uniswap.EventHashV1ExchangeAddLiquidity,
			uniswap.EventHashV1ExchangeRemoveLiquidity,
			uniswap.EventHashV2FactoryPairCreated,
			uniswap.EventHashV3FactoryPoolCreated,
			uniswap.EventHashV2PairSwap,
			uniswap.EventHashV2PairMint,
			uniswap.EventHashV2PairBurn,
			uniswap.EventHashNonfungiblePositionManagerCollect,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	switch {
	case w.matchSwapTransaction(ethereumTask, ethereumTask.Transaction):
		activity.Type = typex.ExchangeSwap
		activity.Actions = w.transformSwapTransaction(ctx, ethereumTask, ethereumTask.Transaction)
	case w.matchLiquidityTransaction(ethereumTask, ethereumTask.Transaction):
		activity.Type = typex.ExchangeLiquidity
		activity.Actions = w.transformLiquidityTransaction(ctx, ethereumTask, ethereumTask.Transaction)
	default:
		return nil, nil
	}

	return activity, nil
}

func (w *worker) matchSwapTransaction(task *source.Task, transaction *ethereum.Transaction) bool {
	if transaction.To == nil {
		return false
	}

	switch *transaction.To {
	case // Uniswap V3
		uniswap.AddressV3SwapRouter,
		uniswap.AddressV3SwapRouter02,
		uniswap.AddressV3SwapRouter02Arbitrum,
		uniswap.AddressV3SwapRouter02Optimism,
		uniswap.AddressV3SwapRouter02Polygon,
		uniswap.AddressV3SwapRouter02Base,
		uniswap.AddressV3SwapRouter02BinanceSmartChain,
		uniswap.AddressV3SwapRouter02Avalanche,
		uniswap.AddressV3SwapRouter02Linea,
		uniswap.AddressUniversalRouter01,
		uniswap.AddressUniversalRouter01Arbitrum,
		uniswap.AddressUniversalRouter01Optimism,
		uniswap.AddressUniversalRouter01Polygon,
		uniswap.AddressUniversalRouter01Base,
		uniswap.AddressUniversalRouter01BinanceSmartChain,
		uniswap.AddressUniversalRouter01Avalanche,
		uniswap.AddressUniversalRouter02:
		return true
	case // Uniswap V2
		uniswap.AddressV2SwapRouter01,
		uniswap.AddressV2SwapRouter02,
		uniswap.AddressV2SwapRouterSAVM:
		return lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
			if len(log.Topics) == 0 {
				return false
			}

			return contract.MatchEventHashes(log.Topics[0], uniswap.EventHashV2PairSwap)
		})
	default: // Uniswap V1
		return lo.ContainsBy(task.Receipt.Logs, func(log *ethereum.Log) bool {
			if len(log.Topics) == 0 {
				return false
			}

			return contract.MatchEventHashes(
				log.Topics[0],
				uniswap.EventHashV1ExchangeTokenPurchase,
				uniswap.EventHashV1ExchangeEthPurchase,
			)
		})
	}
}

func (w *worker) matchLiquidityTransaction(task *source.Task, transaction *ethereum.Transaction) bool {
	if transaction.To == nil {
		return false
	}

	switch *task.Transaction.To {
	case // Uniswap V3
		uniswap.AddressV3Migrator,
		uniswap.AddressV3MigratorArbitrum,
		uniswap.AddressV3MigratorPolygon,
		uniswap.AddressV3MigratorBase,
		uniswap.AddressV3MigratorBinanceSmartChain,
		uniswap.AddressV3MigratorAvalanche,
		uniswap.AddressV3MigratorLinea,
		uniswap.AddressNonfungiblePositionManager,
		uniswap.AddressNonfungiblePositionManagerArbitrum,
		uniswap.AddressNonfungiblePositionManagerPolygon,
		uniswap.AddressNonfungiblePositionManagerOptimism,
		uniswap.AddressNonfungiblePositionManagerBase,
		uniswap.AddressNonfungiblePositionManagerBinanceSmartChain,
		uniswap.AddressNonfungiblePositionManagerAvalanche,
		uniswap.AddressNonfungiblePositionManagerLinea:
		return true
	case // Uniswap V2
		uniswap.AddressV2SwapRouter01,
		uniswap.AddressV2SwapRouter02,
		uniswap.AddressV2SwapRouterSAVM,
		uniswap.AddressV2Migrator:
		return lo.ContainsBy(task.Receipt.Logs, func(item *ethereum.Log) bool {
			if len(item.Topics) == 0 {
				return false
			}

			return contract.MatchEventHashes(
				item.Topics[0],
				uniswap.EventHashV2PairMint,
				uniswap.EventHashV2PairBurn,
			)
		})
	default: // Uniswap V1
		return lo.ContainsBy(task.Receipt.Logs, func(item *ethereum.Log) bool {
			if len(item.Topics) == 0 {
				return false
			}

			return contract.MatchEventHashes(
				item.Topics[0],
				uniswap.EventHashV1ExchangeAddLiquidity,
				uniswap.EventHashV1ExchangeRemoveLiquidity,
			)
		})
	}
}

func (w *worker) matchNonfungiblePositionManagerTransferLog(task *source.Task, log *ethereum.Log) bool {
	managerAddress, err := w.getV3NonfungiblePositionManagerAddress(task.Network)
	if err != nil {
		return false
	}

	if log.Address != managerAddress || len(log.Topics) == 0 || log.Topics[0] != erc721.EventHashTransfer {
		return false
	}

	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseTransfer(log.Export())
	if err != nil {
		return false
	}

	// Mint from genesis address.
	return event.From == ethereum.AddressGenesis
}

func (w *worker) transformSwapTransaction(ctx context.Context, task *source.Task, _ *ethereum.Transaction) (actions []*activityx.Action) {
	var err error

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var buffer []*activityx.Action

		switch log.Export().Topics[0] {
		case uniswap.EventHashV1ExchangeTokenPurchase:
			buffer, err = w.transformV1TokenPurchaseLog(ctx, task, log)
		case uniswap.EventHashV1ExchangeEthPurchase:
			buffer, err = w.transformV1EthPurchaseLog(ctx, task, log)
		case uniswap.EventHashV2PairSwap:
			buffer, err = w.transformV2SwapLog(ctx, task, log)
		case uniswap.EventHashV3PoolSwap:
			buffer, err = w.transformV3SwapLog(ctx, task, log)
		case weth.EventHashDeposit:
			buffer, err = w.transformWETHDepositLog(ctx, task, log)
		case weth.EventHashWithdrawal:
			buffer, err = w.transformWETHWithdrawalLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			zap.L().Error("handle ethereum swap transaction", zap.Error(err), zap.String("task", task.ID()))

			continue
		}

		actions = append(actions, buffer...)
	}

	return actions
}

func (w *worker) transformLiquidityTransaction(ctx context.Context, task *source.Task, _ *ethereum.Transaction) (actions []*activityx.Action) {
	var err error

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var buffer []*activityx.Action

		switch log.Export().Topics[0] {
		case uniswap.EventHashV1ExchangeAddLiquidity:
			buffer, err = w.transformV1ExchangeAddLiquidityLog(ctx, task, log)
		case uniswap.EventHashV1ExchangeRemoveLiquidity:
			buffer, err = w.transformV1ExchangeRemoveLiquidityLog(ctx, task, log)
		case uniswap.EventHashV2PairMint:
			buffer, err = w.transformV2PairMintLog(ctx, task, log)
		case uniswap.EventHashV2PairBurn:
			buffer, err = w.transformV2PairBurnLog(ctx, task, log)
		case uniswap.EventHashNonfungiblePositionManagerIncreaseLiquidity:
			buffer, err = w.transformNonfungiblePositionManagerIncreaseLiquidityLog(ctx, task, log)
		case uniswap.EventHashNonfungiblePositionManagerDecreaseLiquidity:
			buffer, err = w.transformNonfungiblePositionManagerDecreaseLiquidityLog(ctx, task, log)
		case uniswap.EventHashNonfungiblePositionManagerCollect:
			buffer, err = w.transformNonfungiblePositionManagerCollectLog(ctx, task, log)
		case erc721.EventHashTransfer:
			if !w.matchNonfungiblePositionManagerTransferLog(task, log) {
				continue
			}

			buffer, err = w.transformNonfungiblePositionManagerTransferLog(ctx, task, log)
		default:
			continue
		}

		if err != nil {
			zap.L().Error("handle ethereum liquidity transaction", zap.Error(err), zap.String("task", task.ID()))

			continue
		}

		actions = append(actions, buffer...)
	}

	return actions
}

func (w *worker) transformV1TokenPurchaseLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseTokenPurchase(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse TokenPurchase event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	tokenAddr, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token from factory: %w", err)
	}

	if tokenAddr == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial v1 exchange: %s", event.Raw.Address)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, task.Transaction.From, event.Buyer, nil, &tokenAddr, event.EthSold, event.TokensBought)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV1EthPurchaseLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseEthPurchase(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse EthPurchase event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	tokenAddr, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token from factory: %w", err)
	}

	if tokenAddr == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial v1 exhcange: %s", event.Raw.Address)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, task.Transaction.From, event.Buyer, &tokenAddr, nil, event.TokensSold, event.EthBought)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV1ExchangeAddLiquidityLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
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
			Value: lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.EthAmount), 0)),
		},
		{ // ERC-20 token
			Address: lo.ToPtr(tokenAddress.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.TokenAmount), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Provider, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV1ExchangeRemoveLiquidityLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapV1ExchangeFilterer.ParseRemoveLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse RemoveLiquidity event: %w", err)
	}

	factory, err := uniswap.NewV1FactoryCaller(uniswap.AddressV1Factory, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load factory: %w", err)
	}

	tokenAddr, err := factory.GetToken(nil, event.Raw.Address)
	if err != nil {
		return nil, fmt.Errorf("get token address from factory: %w", err)
	}

	if tokenAddr == ethereum.AddressGenesis {
		return nil, fmt.Errorf("unofficial exhcange contract: %s", event.Raw.Address)
	}

	tokens := []metadata.Token{
		{
			// Native token
			Value: lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.EthAmount), 0)),
		},
		{
			// ERC-20 token
			Address: lo.ToPtr(tokenAddr.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.TokenAmount), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, event.Provider, metadata.ActionExchangeLiquidityRemove, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV2SwapLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
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

	var actions []*activityx.Action

	// nolint:revive // False positive
	if event.Amount1Out.Sign() == 1 {
		action, err := w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenLeft, &tokenRight, event.Amount0In, event.Amount1Out)
		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	} else {
		action, err := w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenRight, &tokenLeft, event.Amount1In, event.Amount0Out)
		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	}

	return actions, nil
}

func (w *worker) transformV2PairMintLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
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

	factoryAddress, err := w.getV2FactoryAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get factory address: %w", err)
	}

	factory, err := uniswap.NewV2FactoryCaller(factoryAddress, w.ethereumClient)
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
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount0), 0)),
		},
		{
			Address: lo.ToPtr(tokenRight.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount1), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Sender, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV2PairBurnLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
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

	factoryAddress, err := w.getV2FactoryAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get factory address: %w", err)
	}

	factory, err := uniswap.NewV2FactoryCaller(factoryAddress, w.ethereumClient)
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
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount0), 0)),
		},
		{
			Address: lo.ToPtr(tokenRight.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount1), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityRemove, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformV3SwapLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
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

	factoryAddress, err := w.getV3FactoryAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get factory address: %w", err)
	}

	factory, err := uniswap.NewV3FactoryCaller(factoryAddress, w.ethereumClient)
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

	var actions []*activityx.Action

	// nolint:revive // False positive
	if event.Amount0.Sign() == 1 { // Amount0 is positive
		action, err := w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenLeft, &tokenRight, event.Amount0, event.Amount1)
		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	} else {
		action, err := w.buildExchangeSwapAction(ctx, task, sender, receiver, &tokenRight, &tokenLeft, event.Amount1, event.Amount0)

		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	}

	return actions, nil
}

func (w *worker) transformNonfungiblePositionManagerIncreaseLiquidityLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseIncreaseLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse IncreaseLiquidity event: %w", err)
	}

	managerAddress, err := w.getV3NonfungiblePositionManagerAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get nonfungible position manager address: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(managerAddress, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(&bind.CallOpts{BlockNumber: log.BlockNumber}, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount0), 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount1), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, task.Transaction.From, event.Raw.Address, metadata.ActionExchangeLiquidityAdd, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformNonfungiblePositionManagerDecreaseLiquidityLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseDecreaseLiquidity(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse DecreaseLiquidity event: %w", err)
	}

	managerAddress, err := w.getV3NonfungiblePositionManagerAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get nonfungible position manager address: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(managerAddress, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(&bind.CallOpts{BlockNumber: log.BlockNumber}, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount0), 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount1), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityRemove, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformNonfungiblePositionManagerCollectLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseCollect(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Collect event: %w", err)
	}

	managerAddress, err := w.getV3NonfungiblePositionManagerAddress(task.Network)
	if err != nil {
		return nil, fmt.Errorf("get nonfungible position manager address: %w", err)
	}

	nonfungiblePositionManager, err := uniswap.NewNonfungiblePositionManagerCaller(managerAddress, w.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("load nonfungible position manager: %w", err)
	}

	positions, err := nonfungiblePositionManager.Positions(&bind.CallOpts{BlockNumber: log.BlockNumber}, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("get positions from nonfungible position manager: %w", err)
	}

	tokens := []metadata.Token{
		{
			Address: lo.ToPtr(positions.Token0.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount0), 0)),
		},
		{
			Address: lo.ToPtr(positions.Token1.String()),
			Value:   lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(event.Amount1), 0)),
		},
	}

	action, err := w.buildExchangeLiquidityAction(ctx, task, event.Raw.Address, task.Transaction.From, metadata.ActionExchangeLiquidityCollect, tokens)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformNonfungiblePositionManagerTransferLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.uniswapNonfungiblePositionManagerFilterer.ParseTransfer(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Transfer event: %w", err)
	}

	action, err := w.buildTransactionMintAction(ctx, task, event.From, event.To, event.Raw.Address, event.TokenId)
	if err != nil {
		return nil, fmt.Errorf("build transaction mint action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildExchangeSwapAction(ctx context.Context, task *source.Task, sender, receipt common.Address, tokenIn, tokenOut *common.Address, amountIn, amountOut *big.Int) (*activityx.Action, error) {
	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenIn, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenIn, err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(amountIn), 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenOut, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenOut, err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(amountOut), 0).Abs())

	action := activityx.Action{
		Type:     typex.ExchangeSwap,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}

	return &action, nil
}

func (w *worker) transformWETHDepositLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	if !weth.IsWETH(task.ChainID, log.Address) {
		return nil, fmt.Errorf("unofficial weth contract: %s", log.Address)
	}

	event, err := w.wethWETH9Filterer.ParseDeposit(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Deposit event: %w", err)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, task.Transaction.From, event.Dst, nil, &log.Address, event.Wad, event.Wad)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) transformWETHWithdrawalLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	if !weth.IsWETH(task.ChainID, log.Address) {
		return nil, fmt.Errorf("unofficial weth contract: %s", log.Address)
	}

	event, err := w.wethWETH9Filterer.ParseWithdrawal(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Withdrawal event: %w", err)
	}

	action, err := w.buildExchangeSwapAction(ctx, task, event.Src, task.Transaction.From, &log.Address, nil, event.Wad, event.Wad)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildExchangeLiquidityAction(ctx context.Context, task *source.Task, sender, receipt common.Address, liquidityAction metadata.ExchangeLiquidityAction, tokens []metadata.Token) (*activityx.Action, error) {
	tokenMetadataSlice := make([]metadata.Token, 0, len(tokens))

	for _, t := range tokens {
		var tokenAddress *common.Address
		if t.Address != nil {
			tokenAddress = lo.ToPtr(common.HexToAddress(*t.Address))
		}

		tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenAddress, nil, task.Header.Number)
		if err != nil {
			return nil, fmt.Errorf("lookup token metadata %s: %w", tokenAddress, err)
		}

		tokenMetadata.Value = t.Value

		tokenMetadataSlice = append(tokenMetadataSlice, *tokenMetadata)
	}

	action := activityx.Action{
		Type:     typex.ExchangeLiquidity,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: liquidityAction,
			Tokens: tokenMetadataSlice,
		},
	}

	return &action, nil
}

func (w *worker) buildTransactionMintAction(ctx context.Context, task *source.Task, sender, receipt, tokenAddress common.Address, tokenID *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &tokenAddress, tokenID, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %d: %w", tokenAddress, tokenID, err)
	}

	tokenMetadata.ID = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(tokenID), 0))
	tokenMetadata.Value = lo.ToPtr(decimal.NewFromInt(1))

	action := activityx.Action{
		Type:     typex.TransactionMint,
		Platform: w.Platform(),
		From:     sender.String(),
		To:       receipt.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}

	return &action, nil
}

func (w *worker) getV3NonfungiblePositionManagerAddress(n network.Network) (common.Address, error) {
	switch n {
	case network.Ethereum:
		return uniswap.AddressNonfungiblePositionManager, nil
	case network.Arbitrum:
		return uniswap.AddressNonfungiblePositionManagerArbitrum, nil
	case network.Optimism:
		return uniswap.AddressNonfungiblePositionManagerOptimism, nil
	case network.Polygon:
		return uniswap.AddressNonfungiblePositionManagerPolygon, nil
	case network.Base:
		return uniswap.AddressNonfungiblePositionManagerBase, nil
	case network.BinanceSmartChain:
		return uniswap.AddressNonfungiblePositionManagerBinanceSmartChain, nil
	case network.Avalanche:
		return uniswap.AddressNonfungiblePositionManagerAvalanche, nil
	case network.Linea:
		return uniswap.AddressNonfungiblePositionManagerLinea, nil
	default:
		return ethereum.AddressGenesis, fmt.Errorf("unsupported network: %s", n)
	}
}

func (w *worker) getV3FactoryAddress(n network.Network) (common.Address, error) {
	switch n {
	case network.Ethereum:
		return uniswap.AddressV3Factory, nil
	case network.Arbitrum:
		return uniswap.AddressV3FactoryArbitrum, nil
	case network.Optimism:
		return uniswap.AddressV3FactoryOptimism, nil
	case network.Polygon:
		return uniswap.AddressV3FactoryPolygon, nil
	case network.Base:
		return uniswap.AddressV3FactoryBase, nil
	case network.BinanceSmartChain:
		return uniswap.AddressV3FactoryBinanceSmartChain, nil
	case network.Avalanche:
		return uniswap.AddressV3FactoryAvalanche, nil
	case network.Linea:
		return uniswap.AddressV3FactoryLinea, nil
	default:
		return ethereum.AddressGenesis, fmt.Errorf("unsupported network: %s", n)
	}
}

func (w *worker) getV2FactoryAddress(n network.Network) (common.Address, error) {
	switch n {
	case network.Ethereum:
		return uniswap.AddressV2Factory, nil
	case network.SatoshiVM:
		return uniswap.AddressV2FactorySAVM, nil
	default:
		return ethereum.AddressGenesis, fmt.Errorf("unsupported network: %s", n)
	}
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		uniswapV1ExchangeFilterer:                 lo.Must(uniswap.NewV1ExchangeFilterer(ethereum.AddressGenesis, nil)),
		uniswapV2PairFilterer:                     lo.Must(uniswap.NewV2PairFilterer(ethereum.AddressGenesis, nil)),
		uniswapV3PoolFilterer:                     lo.Must(uniswap.NewV3PoolFilterer(ethereum.AddressGenesis, nil)),
		uniswapNonfungiblePositionManagerFilterer: lo.Must(uniswap.NewNonfungiblePositionManagerFilterer(ethereum.AddressGenesis, nil)),
		wethWETH9Filterer:                         lo.Must(weth.NewWETH9Filterer(ethereum.AddressGenesis, nil)),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
