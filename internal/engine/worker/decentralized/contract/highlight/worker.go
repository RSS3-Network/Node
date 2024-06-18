package highlight

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
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/contract/highlight"
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

// Worker is the worker for Highlight.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config              *config.Module
	ethereumClient      ethereum.Client
	tokenClient         token.Client
	mintManagerFilterer *highlight.MintManagerFilterer
	erc721Filterer      *erc721.ERC721Filterer
}

func (w *worker) Name() string {
	return decentralized.Highlight.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformHighlight.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
		network.Polygon,
		network.Optimism,
		network.Arbitrum,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Collectible,
		tag.Transaction,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.CollectibleMint,
		typex.TransactionTransfer,
	}
}

// Filter highlight contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	var hightlightAddress common.Address

	switch w.config.Network {
	case network.Ethereum:
		hightlightAddress = highlight.AddressMintManagerMainnet
	case network.Polygon:
		hightlightAddress = highlight.AddressMintManagerPolygon
	case network.Optimism:
		hightlightAddress = highlight.AddressMintManagerOptimism
	case network.Arbitrum:
		hightlightAddress = highlight.AddressMintManagerArbitrum
	default:
		hightlightAddress = highlight.AddressMintManagerMainnet
	}

	return &source.Filter{
		LogAddresses: []common.Address{
			hightlightAddress,
		},
		LogTopics: []common.Hash{
			highlight.EventHashNativeGasTokenPayment,
			highlight.EventHashNumTokenMint,
		},
	}
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default highlight activity from task.
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

		// Match highlight core contract events
		switch {
		case w.matchNativeGasTokenPaymentMatched(ethereumTask, log):
			actions, err = w.transformNativeGasTokenPayment(ctx, ethereumTask, log)
		case w.matchNumTokenMintMatched(ethereumTask, log):
			activity.Type = typex.CollectibleMint
			actions, err = w.transformNumTokenMint(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
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

// matchNativeGasTokenPaymentMatched matches NativeGasTokenPayment event.
func (w *worker) matchNativeGasTokenPaymentMatched(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], highlight.EventHashNativeGasTokenPayment)
}

// matchNumTokenMintMatched matches NumTokenMint event.
func (w *worker) matchNumTokenMintMatched(_ *source.Task, log *ethereum.Log) bool {
	return contract.MatchEventHashes(log.Topics[0], highlight.EventHashNumTokenMint)
}

// transformNativeGasTokenPayment transforms NativeGasTokenPayment event.
func (w *worker) transformNativeGasTokenPayment(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse NativeGasTokenPayment event.
	event, err := w.mintManagerFilterer.ParseNativeGasTokenPayment(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse NativeGasTokenPayment event: %w", err)
	}

	creatorFeeAction, err := w.buildTransferAction(ctx, task, task.Transaction.From, event.PaymentRecipient, event.AmountToCreator)
	if err != nil {
		return nil, err
	}

	transactionFee := new(big.Int).Mul(event.AmountToCreator, big.NewInt(int64(event.PercentageBPSOfTotal)))
	transactionFee.Div(transactionFee, big.NewInt(100000))

	transactionFeeAction, err := w.buildTransferAction(ctx, task, task.Transaction.From, *task.Transaction.To, transactionFee)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		creatorFeeAction,
		transactionFeeAction,
	}, nil
}

// transformNumTokenMint transforms NumTokenMint event.
func (w *worker) transformNumTokenMint(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*activityx.Action, error) {
	// Parse NumTokenMint event.
	event, err := w.mintManagerFilterer.ParseNumTokenMint(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse NumTokenMint event: %w", err)
	}

	var tokenIDs []*big.Int

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		if log.Address == event.ContractAddress && contract.MatchEventHashes(log.Topics[0], erc721.EventHashTransfer) {
			transferEvent, err := w.erc721Filterer.ParseTransfer(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse transfer event: %w", err)
			}

			tokenIDs = append(tokenIDs, transferEvent.TokenId)
		}
	}

	actions := make([]*activityx.Action, 0, len(tokenIDs))

	for _, tokenID := range tokenIDs {
		action, err := w.buildHighlightMintAction(ctx, task, ethereum.AddressGenesis, task.Transaction.From, event.ContractAddress, tokenID, big.NewInt(1))
		if err != nil {
			return nil, fmt.Errorf("build collectible mint action: %w", err)
		}

		actions = append(actions, action)
	}

	return actions, nil
}

// buildTransferAction builds transfer action.
func (w *worker) buildTransferAction(ctx context.Context, task *source.Task, from common.Address, to common.Address, amount *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, nil, nil, task.Header.Number)
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

// buildHighlightMintAction builds highlight mint action.
func (w *worker) buildHighlightMintAction(ctx context.Context, task *source.Task, from, to common.Address, contract common.Address, id *big.Int, value *big.Int) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &contract, id, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(value, 0))

	return &activityx.Action{
		Type:     typex.CollectibleMint,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}, nil
}

// NewWorker creates a new Highlight worker.
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

	// Initialize highlight filterers.
	instance.mintManagerFilterer = lo.Must(highlight.NewMintManagerFilterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
