package ens

import (
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/BaseRegistrarImplementation.abi --pkg ens --type BaseRegistrarImplementation --out contract_base_registrar_implementation.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/ETHRegistrarControllerV1.abi --pkg ens --type ETHRegistrarControllerV1 --out contract_eth_registrar_controller_v1.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/ETHRegistrarControllerV2.abi --pkg ens --type ETHRegistrarControllerV2 --out contract_eth_registrar_controller_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/PublicResolverV1.abi --pkg ens --type PublicResolverV1 --out contract_public_resolver_v1.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/PublicResolverV2.abi --pkg ens --type PublicResolverV2 --out contract_public_resolver_v2.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/NameWrapper.abi --pkg ens --type NameWrapper --out contract_name_wrapper.go

var (
	AddressBaseRegistrarImplementation = common.HexToAddress("0x57f1887a8BF19b14fC0dF6Fd9B2acc9Af147eA85")
	AddressETHRegistrarControllerV1    = common.HexToAddress("0x283Af0B28c62C092C9727F1Ee09c02CA627EB7F5")
	AddressETHRegistrarControllerV2    = common.HexToAddress("0x253553366Da8546fC250F225fe3d25d0C782303b")
	AddressPublicResolverV1            = common.HexToAddress("0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41")
	AddressPublicResolverV2            = common.HexToAddress("0x231b0ee14048e9dccd1d247744d114a4eb5e8e63")
	AddressNameWrapper                 = common.HexToAddress("0xD4416b13d2b3a9aBae7AcD5D6C2BbDBE25686401")

	EventNameRegisteredControllerV1 = contract.EventHash("NameRegistered(string,bytes32,address,uint256,uint256)")
	EventNameRegisteredControllerV2 = contract.EventHash("NameRegistered(string,bytes32,address,uint256,uint256,uint256)")
	EventNameRenewed                = contract.EventHash("NameRenewed(string,bytes32,uint256,uint256)")
	EventTextChanged                = contract.EventHash("TextChanged(bytes32,string,string)")
	EventTextChangedWithValue       = contract.EventHash("TextChanged(bytes32,string,string,string)")
	EventNameWrapped                = contract.EventHash("NameWrapped(bytes32,bytes,address,uint32,uint64)")
	EventNameUnwrapped              = contract.EventHash("NameUnwrapped(bytes32,address)")

	EventFusesSet           = contract.EventHash("FusesSet(bytes32,uint32)")
	EventAddressChanged     = contract.EventHash("AddressChanged(bytes32,uint256,bytes)")
	EventContenthashChanged = contract.EventHash("ContenthashChanged(bytes32,bytes)")
	EventNameChanged        = contract.EventHash("NameChanged(bytes32,string)")
	EventPubkeyChanged      = contract.EventHash("PubkeyChanged(bytes32,bytes32,bytes32)")
	// EventVersionChanged     = contract.EventHash("VersionChanged(bytes32,uint64)")
	// EventInterfaceChanged   = contract.EventHash("InterfaceChanged(bytes32,bytes4,address)")
)

// NameHash is a recursive process that can generate a unique hash for any valid domain name.
// https://docs.ens.domains/contract-api-reference/name-processing#algorithm
func NameHash(name string) (node common.Hash) {
	labels := strings.Split(name, ".")

	// Reverse domain name notation. https://en.wikipedia.org/wiki/Reverse_domain_name_notation
	sort.SliceStable(labels, func(_, _ int) bool {
		return true
	})

	for _, label := range labels {
		node = crypto.Keccak256Hash(node.Bytes(), crypto.Keccak256Hash([]byte(label)).Bytes())
	}

	return node
}
