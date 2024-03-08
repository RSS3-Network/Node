package lido

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
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/contract/lido"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// Worker is the worker for Lido.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	// Base
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	// Specific Filters
	stakedETHFilterer              *lido.StakedETHFilterer
	wrappedStakedETHFilterer       *lido.WrappedStakedETHFilterer
	stakedETHWithdrawalNFTFilterer *lido.StakedETHWithdrawalNFTFilterer
	stakedMATICFilterer            *lido.StakedMATICFilterer
	// Common filters
	erc721Filterer *erc721.ERC721Filterer
}

func (w *worker) Name() string {
	return filter.Lido.String()
}

// Filter lido contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			lido.AddressStakedETH,
			lido.AddressWrappedStakedETH,
			lido.AddressStakedETHWithdrawalNFT,
			lido.AddressStakedMATIC,
			lido.AddressStakedMATICWithdrawalNFT,
			lido.AddressMATIC,
		},
		LogTopics: []common.Hash{
			lido.EventHashStakedETHSubmitted,
			lido.EventHashStakedETHWithdrawalNFTWithdrawalRequested,
			lido.EventHashStakedETHWithdrawalNFTWithdrawalClaimed,
			lido.EventHashStakedMATICSubmitEvent,
			lido.EventHashStakedMATICRequestWithdrawEvent,
			lido.EventHashStakedMATICClaimTokensEvent,
			lido.EventHashStakedETHTransferShares,
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

	// Build default lido feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformLido))
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

		// Match lido core contract events
		switch {
		case w.matchStakedETHSubmittedLog(ethereumTask, log):
			// Add ETH liquidity
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedETHSubmittedLog(ctx, ethereumTask, log)
		case w.matchStakedETHWithdrawalNFTWithdrawalRequestedLog(ethereumTask, log):
			// Mint ETH withdrawal NFT
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedETHWithdrawalNFTWithdrawalRequestedLog(ctx, ethereumTask, log)
		case w.matchStakedETHWithdrawalNFTWithdrawalClaimedLog(ethereumTask, log):
			// Remove ETH liquidity
			feed.Type = filter.TypeCollectibleBurn
			actions, err = w.transformStakedETHWithdrawalNFTWithdrawalClaimedLog(ctx, ethereumTask, log)
		case w.matchStakedMATICSubmitEventLog(ethereumTask, log):
			// Add MATIC liquidity
			feed.Type = filter.TypeExchangeLiquidity
			actions, err = w.transformStakedMATICSubmitEventLog(ctx, ethereumTask, log)
		case w.matchStakedMATICRequestWithdrawEventLog(ethereumTask, log):
			// Mint MATIC withdrawal NFT
			feed.Type = filter.TypeCollectibleMint
			actions, err = w.transformStakedMATICRequestWithdrawEventLog(ctx, ethereumTask, log)
		case w.matchStakedMATICClaimTokensEventLog(ethereumTask, log):
			// Remove MATIC liquidity
			feed.Type = filter.TypeCollectibleBurn
			actions, err = w.transformStakedMATICClaimTokensEventLog(ctx, ethereumTask, log)
		case w.matchStakedETHTransferSharesLog(ethereumTask, log):
			// Wrap or unwrap wstETH
			feed.Type = filter.TypeExchangeSwap
			actions, err = w.transformStakedETHTransferSharesLog(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// matchStakedETHSubmittedLog matches events that Add ETH liquidity
func (w *worker) matchStakedETHSubmittedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHSubmitted) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETH)
}

// matchStakedETHWithdrawalNFTWithdrawalRequestedLog matches events that Mint ETH withdrawal NFT
func (w *worker) matchStakedETHWithdrawalNFTWithdrawalRequestedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHWithdrawalNFTWithdrawalRequested) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETHWithdrawalNFT)
}

// matchStakedETHWithdrawalNFTWithdrawalClaimedLog matches events that Remove ETH liquidity
func (w *worker) matchStakedETHWithdrawalNFTWithdrawalClaimedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHWithdrawalNFTWithdrawalClaimed) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETHWithdrawalNFT)
}

// matchStakedMATICSubmitEventLog matches events that Add MATIC liquidity
func (w *worker) matchStakedMATICSubmitEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICSubmitEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

// matchStakedMATICRequestWithdrawEventLog matches events that Mint MATIC withdrawal NFT
func (w *worker) matchStakedMATICRequestWithdrawEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICRequestWithdrawEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

// matchStakedMATICClaimTokensEventLog matches events that Remove MATIC liquidity
func (w *worker) matchStakedMATICClaimTokensEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICClaimTokensEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

// matchStakedETHTransferSharesLog matches events that Wrap or unwrap wstETH
func (w *worker) matchStakedETHTransferSharesLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHTransferShares) {
		return false
	}

	if !contract.MatchAddresses(log.Address, lido.AddressStakedETH) {
		return false
	}

	event, err := w.stakedETHFilterer.ParseTransferShares(log.Export())
	if err != nil {
		return false
	}

	return event.From == lido.AddressWrappedStakedETH || event.To == lido.AddressWrappedStakedETH
}

// transformStakedETHWithdrawalNFTWithdrawalClaimedLog processes events that Remove ETH liquidity
func (w *worker) transformStakedETHWithdrawalNFTWithdrawalClaimedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakedETHWithdrawalNFTFilterer.ParseWithdrawalClaimed(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseWithdrawalClaimed event: %w", err)
	}

	// Burn NFT
	collectibleTransferAction, err := w.buildEthereumCollectibleTransferAction(ctx, task.Header.Number, task.ChainID, event.Owner, ethereum.AddressGenesis, event.Raw.Address, event.RequestId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	// Withdraw ETH
	transactionTransferAction, err := w.buildEthereumTransactionTransferAction(ctx, task.Header.Number, task.ChainID, event.Raw.Address, event.Receiver, nil, event.AmountOfETH)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	return []*schema.Action{
		collectibleTransferAction,
		transactionTransferAction,
	}, nil
}

// transformStakedETHWithdrawalNFTWithdrawalRequestedLog processes events that Mint ETH withdrawal NFT
func (w *worker) transformStakedETHWithdrawalNFTWithdrawalRequestedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakedETHWithdrawalNFTFilterer.ParseWithdrawalRequested(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseWithdrawalRequested event: %w", err)
	}

	exchangeLiquidityAction, err := w.buildEthereumExchangeLiquidityAction(ctx, task.Header.Number, task.ChainID, event.Owner, event.Raw.Address, nil, event.AmountOfStETH, metadata.ActionExchangeLiquidityRemove)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	collectibleTransferAction, err := w.buildEthereumCollectibleTransferAction(ctx, task.Header.Number, task.ChainID, ethereum.AddressGenesis, event.Owner, event.Raw.Address, event.RequestId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	return []*schema.Action{
		exchangeLiquidityAction,
		collectibleTransferAction,
	}, nil
}

// transformStakedETHSubmittedLog processes events that Add ETH liquidity
func (w *worker) transformStakedETHSubmittedLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakedETHFilterer.ParseSubmitted(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseSubmitted event: %w", err)
	}

	exchangeLiquidityAction, err := w.buildEthereumExchangeLiquidityAction(ctx, task.Header.Number, task.ChainID, event.Sender, event.Raw.Address, nil, event.Amount, metadata.ActionExchangeLiquidityAdd)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	transactionTransferAction, err := w.buildEthereumTransactionTransferAction(ctx, task.Header.Number, task.ChainID, ethereum.AddressGenesis, event.Sender, &event.Raw.Address, event.Amount)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	return []*schema.Action{
		exchangeLiquidityAction,
		transactionTransferAction,
	}, nil
}

// transformStakedMATICSubmitEventLog processes events that Add MATIC liquidity
func (w *worker) transformStakedMATICSubmitEventLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakedMATICFilterer.ParseSubmitEvent(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseSubmitEvent event: %w", err)
	}

	exchangeLiquidityAction, err := w.buildEthereumExchangeLiquidityAction(ctx, task.Header.Number, task.ChainID, event.From, event.Raw.Address, nil, event.Amount, metadata.ActionExchangeLiquidityAdd)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	transactionTransferAction, err := w.buildEthereumTransactionTransferAction(ctx, task.Header.Number, task.ChainID, ethereum.AddressGenesis, event.From, &event.Raw.Address, event.Amount)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	return []*schema.Action{
		exchangeLiquidityAction,
		transactionTransferAction,
	}, nil
}

// transformStakedMATICRequestWithdrawEventLog processes events that Mint MATIC withdrawal NFT
func (w *worker) transformStakedMATICRequestWithdrawEventLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	withdrawEvent, err := w.stakedMATICFilterer.ParseRequestWithdrawEvent(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseRequestWithdrawEvent event: %w", err)
	}

	transferLog, exists := lo.Find(task.Receipt.Logs, func(log *ethereum.Log) bool {
		return len(log.Topics) == 4 && log.Address == lido.AddressStakedMATICWithdrawalNFT && contract.MatchEventHashes(log.Topics[0], erc721.EventHashTransfer)
	})

	if !exists {
		return nil, fmt.Errorf("no collectible transfer event found")
	}

	transferEvent, err := w.erc721Filterer.ParseTransfer(transferLog.Export())
	if err != nil {
		return nil, fmt.Errorf("parse collectible transfer event: %w", err)
	}

	exchangeLiquidityAction, err := w.buildEthereumExchangeLiquidityAction(ctx, task.Header.Number, task.ChainID, withdrawEvent.From, lido.AddressStakedMATICWithdrawalNFT, &lido.AddressStakedMATIC, withdrawEvent.Amount, metadata.ActionExchangeLiquidityRemove)
	if err != nil {
		return nil, fmt.Errorf("build exchange liquidity action: %w", err)
	}

	collectibleTransferAction, err := w.buildEthereumCollectibleTransferAction(ctx, task.Header.Number, task.ChainID, transferEvent.From, transferEvent.To, transferEvent.Raw.Address, transferEvent.TokenId, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	return []*schema.Action{
		exchangeLiquidityAction,
		collectibleTransferAction,
	}, nil
}

// transformStakedMATICClaimTokensEventLog processes events that Remove MATIC liquidity
func (w *worker) transformStakedMATICClaimTokensEventLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.stakedMATICFilterer.ParseClaimTokensEvent(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseClaimTokensEvent event: %w", err)
	}

	// Burn NFT
	burnNFTAction, err := w.buildEthereumCollectibleTransferAction(ctx, task.Header.Number, task.ChainID, event.From, ethereum.AddressGenesis, lido.AddressStakedMATICWithdrawalNFT, event.Id, big.NewInt(1))
	if err != nil {
		return nil, fmt.Errorf("build collectible transfer action: %w", err)
	}

	// Withdraw MATIC
	withdrawMaticAction, err := w.buildEthereumTransactionTransferAction(ctx, task.Header.Number, task.ChainID, event.Raw.Address, event.From, &lido.AddressMATIC, event.AmountClaimed)
	if err != nil {
		return nil, fmt.Errorf("build transaction transfer action: %w", err)
	}

	return []*schema.Action{
		burnNFTAction,
		withdrawMaticAction,
	}, nil
}

// transformStakedETHTransferSharesLog processes events that Wrap or unwrap wstETH
func (w *worker) transformStakedETHTransferSharesLog(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	transferSharesEvent, err := w.stakedETHFilterer.ParseTransferShares(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseTransferShares event: %w", err)
	}

	wstETHTransferLog, exists := lo.Find(task.Receipt.Logs, func(log *ethereum.Log) bool {
		return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer) && log.Address == lido.AddressWrappedStakedETH
	})

	if !exists {
		return nil, fmt.Errorf("no wstETH transfer event found")
	}

	wstETHTransferEvent, err := w.wrappedStakedETHFilterer.ParseTransfer(wstETHTransferLog.Export())
	if err != nil {
		return nil, fmt.Errorf("parse wstETH transfer event: %w", err)
	}

	stETHTransferLog, exists := lo.Find(task.Receipt.Logs, func(log *ethereum.Log) bool {
		return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer) && log.Address == lido.AddressStakedETH
	})

	if !exists {
		return nil, fmt.Errorf("no wstETH transfer event found")
	}

	stETHTransferEvent, err := w.stakedETHFilterer.ParseTransfer(stETHTransferLog.Export())
	if err != nil {
		return nil, fmt.Errorf("parse wstETH transfer event: %w", err)
	}

	actions := make([]*schema.Action, 0)

	switch {
	case transferSharesEvent.To == lido.AddressWrappedStakedETH:
		// Swap stETH to wstETH
		action, err := w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, task.Transaction.From, task.Transaction.From, lido.AddressStakedETH, lido.AddressWrappedStakedETH, stETHTransferEvent.Value, wstETHTransferEvent.Value)
		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	case transferSharesEvent.From == lido.AddressWrappedStakedETH:
		// Swap wstETH to stETH
		action, err := w.buildEthereumExchangeSwapAction(ctx, task.Header.Number, task.ChainID, task.Transaction.From, task.Transaction.From, lido.AddressWrappedStakedETH, lido.AddressStakedETH, wstETHTransferEvent.Value, stETHTransferEvent.Value)
		if err != nil {
			return nil, fmt.Errorf("build exchange swap action: %w", err)
		}

		actions = append(actions, action)
	default:
		return nil, fmt.Errorf("unmatched TransferShares event")
	}

	return actions, nil
}

func (w *worker) buildEthereumExchangeLiquidityAction(ctx context.Context, blockNumber *big.Int, chainID uint64, sender, receiver common.Address, tokenAddress *common.Address, tokenValue *big.Int, liquidityAction metadata.ExchangeLiquidityAction) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	action := schema.Action{
		Type:     filter.TypeExchangeLiquidity,
		Platform: filter.PlatformLido.String(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.ExchangeLiquidity{
			Action: liquidityAction,
			Tokens: []metadata.Token{
				*tokenMetadata,
			},
		},
	}

	return &action, nil
}

func (w *worker) buildEthereumTransactionTransferAction(ctx context.Context, blockNumber *big.Int, chainID uint64, sender, receiver common.Address, tokenAddress *common.Address, tokenValue *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, tokenAddress, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	actionType := filter.TypeTransactionTransfer

	if sender == ethereum.AddressGenesis {
		actionType = filter.TypeTransactionMint
	}

	if receiver == ethereum.AddressGenesis {
		actionType = filter.TypeTransactionBurn
	}

	action := schema.Action{
		Type:     actionType,
		Platform: filter.PlatformLido.String(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.TransactionTransfer(*tokenMetadata),
	}

	return &action, nil
}

func (w *worker) buildEthereumCollectibleTransferAction(ctx context.Context, blockNumber *big.Int, chainID uint64, sender, receiver, tokenAddress common.Address, tokenID, tokenValue *big.Int) (*schema.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, chainID, &tokenAddress, tokenID, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token: %w", err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(tokenValue, 0))

	actionType := filter.TypeCollectibleTransfer

	if sender == ethereum.AddressGenesis {
		actionType = filter.TypeCollectibleMint
	}

	if receiver == ethereum.AddressGenesis {
		actionType = filter.TypeCollectibleBurn
	}

	action := schema.Action{
		Type:     actionType,
		Platform: filter.PlatformLido.String(),
		From:     sender.String(),
		To:       receiver.String(),
		Metadata: metadata.CollectibleTransfer(*tokenMetadata),
	}

	return &action, nil
}

func (w *worker) buildEthereumExchangeSwapAction(ctx context.Context, blockNumber *big.Int, chainID uint64, from, to, tokenIn, tokenOut common.Address, amountIn, amountOut *big.Int) (*schema.Action, error) {
	tokenInMetadata, err := w.tokenClient.Lookup(ctx, chainID, &tokenIn, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenIn, err)
	}

	tokenInMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountIn, 0).Abs())

	tokenOutMetadata, err := w.tokenClient.Lookup(ctx, chainID, &tokenOut, nil, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", tokenOut, err)
	}

	tokenOutMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(amountOut, 0).Abs())

	action := schema.Action{
		Type:     filter.TypeExchangeSwap,
		Platform: filter.PlatformLido.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.ExchangeSwap{
			From: *tokenInMetadata,
			To:   *tokenOutMetadata,
		},
	}

	return &action, nil
}

// NewWorker creates a new Lido worker.
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

	// Initialize lido filterers.
	instance.stakedETHFilterer = lo.Must(lido.NewStakedETHFilterer(ethereum.AddressGenesis, nil))
	instance.wrappedStakedETHFilterer = lo.Must(lido.NewWrappedStakedETHFilterer(ethereum.AddressGenesis, nil))
	instance.stakedETHWithdrawalNFTFilterer = lo.Must(lido.NewStakedETHWithdrawalNFTFilterer(ethereum.AddressGenesis, nil))
	instance.stakedMATICFilterer = lo.Must(lido.NewStakedMATICFilterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
