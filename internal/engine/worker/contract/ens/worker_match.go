package ens

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/ens"
)

func (w *worker) matchEnsNameRegisteredV1(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressETHRegistrarControllerV1 && log.Topics[0] == ens.EventNameRegisteredControllerV1
}

func (w *worker) matchEnsNameRegisteredV2(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressETHRegistrarControllerV2 && log.Topics[0] == ens.EventNameRegisteredControllerV2
}

func (w *worker) matchEnsNameRenewed(_ context.Context, log ethereum.Log) bool {
	return (log.Address == ens.AddressETHRegistrarControllerV1 || log.Address == ens.AddressETHRegistrarControllerV2) && log.Topics[0] == ens.EventNameRenewed
}

func (w *worker) matchEnsTextChanged(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressPublicResolverV1 && log.Topics[0] == ens.EventTextChanged
}

func (w *worker) matchEnsTextChangedWithValue(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressPublicResolverV2 && log.Topics[0] == ens.EventTextChangedWithValue
}

func (w *worker) matchEnsNameWrapped(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventNameWrapped
}

func (w *worker) matchEnsNameUnwrapped(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventNameUnwrapped
}

func (w *worker) matchEnsFusesSet(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventFusesSet
}

func (w *worker) matchEnsContenthashChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventContenthashChanged
}

func (w *worker) matchEnsNameChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventNameChanged
}

func (w *worker) matchEnsAddressChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventAddressChanged
}

func (w *worker) matchEnsPubkeyChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventPubkeyChanged
}
