package lido

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc20"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/erc721"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/lido"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

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
