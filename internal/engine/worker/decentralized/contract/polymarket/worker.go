package polymarket

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/polymarket"
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
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	ethereumClient ethereum.Client
	tokenClient    token.Client
	ctfExchange    *polymarket.CTFExchangeFilterer
}

func (w *worker) Name() string {
	return decentralized.Polymarket.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformPolymarket.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Polygon,
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

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			polymarket.AddressPolyMarketCTFExchange,
			polymarket.AddressPolyMarketNegRiskCTFExchange,
		},
		LogTopics: []common.Hash{
			polymarket.EventOrderFinalized,
		},
	}
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	polygonTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range polygonTask.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*activityx.Action
			err     error
		)

		switch {
		case w.matchOrderFinalizedLog(polygonTask, log):
			actions, err = w.transformOrderFinalizedLog(ctx, polygonTask, log)

		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	activity.Type = typex.CollectibleTrade

	return activity, nil
}

func (w *worker) matchOrderFinalizedLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], polymarket.EventOrderFinalized) &&
		contract.MatchAddresses(log.Address, polymarket.AddressPolyMarketCTFExchange, polymarket.AddressPolyMarketNegRiskCTFExchange)
}

func (w *worker) transformOrderFinalizedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	var err error

	// CTF and NegRiskCTF shares the same OrderFilled struct
	event, err := w.ctfExchange.ParseOrderFilled(log.Export())

	if err != nil {
		return nil, fmt.Errorf("parse OrderFilled event: %w", err)
	}

	buyAction, sellAction, err := w.buildMarketTradeAction(ctx, task, task.ChainID, event.Maker, event.Taker, event.MakerAssetId, event.TakerAssetId, event.OrderHash, event.MakerAmountFilled, event.TakerAmountFilled)
	if err != nil {
		return nil, fmt.Errorf("build market trade action: %w", err)
	}

	return []*activityx.Action{buyAction, sellAction}, nil
}

func (w *worker) buildMarketTradeAction(ctx context.Context, _ *source.Task, chainID uint64, maker, taker common.Address, makerAssetID, takerAssetID *big.Int, _ [32]byte, makerAmountFilled, takerAmountFilled *big.Int) (*activityx.Action, *activityx.Action, error) {
	makerAmountFilledDecimal := decimal.NewFromBigInt(utils.GetBigInt(makerAmountFilled), 0)
	takerAmountFilledDecimal := decimal.NewFromBigInt(utils.GetBigInt(takerAmountFilled), 0)

	var takerTokenAddress *common.Address
	if takerAssetID.Cmp(big.NewInt(0)) == 0 {
		takerTokenAddress = nil
	} else {
		address := common.HexToAddress(polymarket.AddressPolyMarketConditionTokens.String())
		takerTokenAddress = &address
	}

	takerToken, err := w.tokenClient.Lookup(ctx, chainID, takerTokenAddress, takerAssetID, nil)

	if err != nil {
		return nil, nil, fmt.Errorf("lookup taker token: %w", err)
	}

	takerToken.Value = &takerAmountFilledDecimal

	var makerTokenAddress *common.Address

	if makerAssetID.Cmp(big.NewInt(0)) == 0 {
		makerTokenAddress = nil
	} else {
		address := common.HexToAddress(polymarket.AddressPolyMarketConditionTokens.String())
		makerTokenAddress = &address
	}

	makerToken, err := w.tokenClient.Lookup(ctx, chainID, makerTokenAddress, makerAssetID, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("lookup maker token: %w", err)
	}

	makerToken.Value = &makerAmountFilledDecimal

	buyAction := &activityx.Action{
		Type:     typex.CollectibleTrade,
		Platform: w.Platform(),
		From:     taker.String(),
		To:       maker.String(),
		Metadata: metadata.CollectibleTrade{
			Action: metadata.ActionCollectibleTradeBuy,
			Token:  *takerToken,
			Cost:   makerToken,
		},
	}

	// Sell action (from the maker's perspective)
	sellAction := &activityx.Action{
		Type:     typex.CollectibleTrade,
		Platform: w.Platform(),
		From:     maker.String(),
		To:       taker.String(),
		Metadata: metadata.CollectibleTrade{
			Action: metadata.ActionCollectibleTradeSell,
			Token:  *takerToken,
			Cost:   makerToken,
		},
	}

	return buyAction, sellAction, nil
}

func NewWorker(config *config.Module) (engine.Worker, error) {
	instance := worker{
		ctfExchange: lo.Must(polymarket.NewCTFExchangeFilterer(ethereum.AddressGenesis, nil)),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
