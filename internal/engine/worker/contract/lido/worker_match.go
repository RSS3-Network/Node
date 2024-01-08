package lido

import (
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/lido"
)

func (w *worker) matchStakedETHSubmittedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHSubmitted) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETH)
}

func (w *worker) matchStakedETHWithdrawalNFTWithdrawalRequestedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHWithdrawalNFTWithdrawalRequested) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETHWithdrawalNFT)
}

func (w *worker) matchStakedETHWithdrawalNFTWithdrawalClaimedLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedETHWithdrawalNFTWithdrawalClaimed) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedETHWithdrawalNFT)
}

func (w *worker) matchStakedMATICSubmitEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICSubmitEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

func (w *worker) matchStakedMATICRequestWithdrawEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICRequestWithdrawEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

func (w *worker) matchStakedMATICClaimTokensEventLog(_ *source.Task, log *ethereum.Log) bool {
	if !contract.MatchEventHashes(log.Topics[0], lido.EventHashStakedMATICClaimTokensEvent) {
		return false
	}

	return contract.MatchAddresses(log.Address, lido.AddressStakedMATIC)
}

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
