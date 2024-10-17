package zerion

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
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/zerion"
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
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

// worker represents the Zerion worker, which processes Zerion-related transactions.
type worker struct {
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	erc20Filterer  *erc20.ERC20Filterer
	routerFilterer *zerion.RouterFilterer
}

// Name returns the name of the worker.
func (w *worker) Name() string {
	return decentralized.Zerion.String()
}

// Platform returns the platform name for the Zerion worker.
func (w *worker) Platform() string {
	return decentralized.PlatformZerion.String()
}

// Network returns the list of networks supported by the Zerion worker.
func (w *worker) Network() []network.Network {
	return []network.Network{
		network.BinanceSmartChain,
		network.Polygon,
		network.Arbitrum,
		network.Avalanche,
		network.Optimism,
		network.Linea,
		network.Gnosis,
		network.XLayer,
		network.Base,
	}
}

// Tags returns the list of tags associated with the Zerion worker.
func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Exchange,
		tag.Transaction,
	}
}

// Types returns the list of activity types that the Zerion worker can process.
func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.ExchangeSwap,
		typex.TransactionTransfer,
	}
}

// Filter returns the data protocol filter for the Zerion worker.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			zerion.AddressRouter,
		},
		LogTopics: []common.Hash{
			zerion.EventHashExecuted,
		},
	}
}

// Transform processes the input task and transforms it into an activity.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type %T", task)
	}

	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))

	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Iterate through all logs in the transaction receipt
	for _, log := range ethereumTask.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchSwapLog(ethereumTask, log):
			actions, err := w.transformSwapLog(ctx, ethereumTask, log)
			if err != nil {
				zap.L().Warn("handle settlement trade log", zap.Error(err), zap.String("worker", w.Name()), zap.String("task", ethereumTask.ID()))
				continue
			}

			activity.Actions = append(activity.Actions, actions...)

		default:
			zap.L().Debug("unsupported log", zap.String("worker", w.Name()), zap.String("task", ethereumTask.ID()), zap.Stringer("topic", log.Topics[0]))
		}
	}

	if len(activity.Actions) == 0 {
		return nil, fmt.Errorf("no actions")
	}

	zap.L().Info("Processing task", zap.Any("task", ethereumTask))
	zap.L().Info("activity is: ", zap.Any("activity", activity))

	activity.Type = typex.ExchangeSwap

	return activity, nil
}

// matchSwapLog checks if the given log matches a Zerion swap event.
func (w *worker) matchSwapLog(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], zerion.EventHashExecuted) &&
		contract.MatchAddresses(log.Address, zerion.AddressRouter)
}

// transformSwapLog transforms a Zerion swap log into a list of actions.
func (w *worker) transformSwapLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	event, err := w.routerFilterer.ParseExecuted(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse Executed event: %w", err)
	}

	var actions []*activityx.Action

	// Handle protocol fee if present
	if event.ProtocolFeeAmount.Sign() == 1 {
		feeAction, err := w.buildTransactionTransferAction(ctx, task, event.Sender, zerion.AddressRouter, lo.Ternary(event.OutputToken == zerion.AddressNativeToken, nil, &event.OutputToken), event.ProtocolFeeAmount)
		if err != nil {
			return nil, fmt.Errorf("build transaction transfer action for fee: %w", err)
		}

		actions = append(actions, feeAction)
	}

	// Build the main swap action
	swapAction, err := w.buildExchangeSwapAction(ctx, task, event.Sender, event.Sender, event.InputToken, event.OutputToken, event.AbsoluteInputAmount, event.ReturnedAmount)
	if err != nil {
		return nil, fmt.Errorf("build exchange swap action: %w", err)
	}

	actions = append(actions, swapAction)

	return actions, nil
}

// buildTransactionTransferAction creates a TransactionTransfer action for a given transfer.
func (w *worker) buildTransactionTransferAction(ctx context.Context, task *source.Task, from, to common.Address, tokenAddress *common.Address, amount *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amount, 0))

	return &activityx.Action{
		Type:     typex.TransactionTransfer,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}, nil
}

// buildExchangeSwapAction creates an ExchangeSwap action for a given swap.
func (w *worker) buildExchangeSwapAction(ctx context.Context, task *source.Task, from, to, tokenIn, tokenOut common.Address, amountIn, amountOut *big.Int) (*activityx.Action, error) {
	tokenInAddress := lo.Ternary(tokenIn != zerion.AddressNativeToken, &tokenIn, nil)
	tokenOutAddress := lo.Ternary(tokenOut != zerion.AddressNativeToken, &tokenOut, nil)

	tokenInMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenInAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token in metadata: %w", err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0))

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, tokenOutAddress, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token out metadata: %w", err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountOut, 0))

	return &activityx.Action{
		Type:     typex.ExchangeSwap,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}, nil
}

// NewWorker creates and initializes a new Zerion worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var err error

	instance := worker{
		config: config,
	}

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize filterers.
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.routerFilterer = lo.Must(zerion.NewRouterFilterer(zerion.AddressRouter, nil))

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	return &instance, nil
}
