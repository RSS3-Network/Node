package stargate

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/protocol-go/schema/network"
)

// Factory
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Factory.abi --pkg stargate --type Factory --out contract_factory.go

// Pool
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Pool.abi --pkg stargate --type Pool --out contract_pool.go

// https://stargateprotocol.gitbook.io/stargate/developers/contract-addresses/mainnet
var (
	AddressRouterMainnet           = common.HexToAddress("0x8731d54E9D02c286767d56ac03e8037C07e01e98")
	AddressRouterBinanceSmartChain = common.HexToAddress("0x4a364f8c717cAAD9A442737Eb7b8A55cc6cf18D8")
	AddressRouterAvalanche         = common.HexToAddress("0x45A01E4e04F14f7A4a6702c74187c5F6222033cd")
	AddressRouterPolygon           = common.HexToAddress("0x45A01E4e04F14f7A4a6702c74187c5F6222033cd")
	AddressRouterArbitrumOne       = common.HexToAddress("0x53Bf833A5d6c4ddA888F69c22C88C9f356a41614")
	AddressRouterOptimism          = common.HexToAddress("0xB0D502E938ed5f4df2E681fE6E419ff29631d62b")
	AddressRouterBase              = common.HexToAddress("0x45f1A95A4D3f3836523F5c83673c797f4d4d263B")
	AddressRouterLinea             = common.HexToAddress("0x2F6F07CDcf3588944Bf4C42aC74ff24bF56e7590")

	AddressRouterETHMainnet     = common.HexToAddress("0x150f94B44927F078737562f0fcF3C95c01Cc2376")
	AddressRouterETHArbitrumOne = common.HexToAddress("0xbf22f0f184bCcbeA268dF387a49fF5238dD23E40")
	AddressRouterETHOptimism    = common.HexToAddress("0xB49c4e680174E331CB0A7fF3Ab58afC9738d5F8b")
	AddressRouterETHBase        = common.HexToAddress("0x50B6EbC2103BFEc165949CC946d739d5650d7ae4")
	AddressRouterETHLinea       = common.HexToAddress("0x8731d54E9D02c286767d56ac03e8037C07e01e98")

	AddressFactoryMainnet           = common.HexToAddress("0x06D538690AF257Da524f25D0CD52fD85b1c2173E")
	AddressFactoryBinanceSmartChain = common.HexToAddress("0xe7Ec689f432f29383f217e36e680B5C855051f25")
	AddressFactoryAvalanche         = common.HexToAddress("0x808d7c71ad2ba3FA531b068a2417C63106BC0949")
	AddressFactoryPolygon           = common.HexToAddress("0x808d7c71ad2ba3FA531b068a2417C63106BC0949")
	AddressFactoryArbitrumOne       = common.HexToAddress("0x55bDb4164D28FBaF0898e0eF14a589ac09Ac9970")
	AddressFactoryOptimism          = common.HexToAddress("0xE3B53AF74a4BF62Ae5511055290838050bf764Df")
	AddressFactoryBase              = common.HexToAddress("0xAf5191B0De278C7286d6C7CC6ab6BB8A73bA2Cd6")
	AddressFactoryLinea             = common.HexToAddress("0xaf54be5b6eec24d6bfacf1cce4eaf680a8239398")

	AddressSGETHMainnet     = common.HexToAddress("0x72E2F4830b9E45d52F80aC08CB2bEC0FeF72eD9c")
	AddressSGETHArbitrumOne = common.HexToAddress("0x82CbeCF39bEe528B5476FE6d1550af59a9dB6Fc0")
	AddressSGETHOptimism    = common.HexToAddress("0xb69c8CBCD90A39D8D3d3ccf0a3E968511C3856A0")
	AddressSGETHBase        = common.HexToAddress("0x224D8Fd7aB6AD4c6eb4611Ce56EF35Dec2277F03")
	AddressSGETHLinea       = common.HexToAddress("0xAad094F6A75A14417d39f04E690fC216f080A41a")

	EventHashPoolSwap            = contract.EventHash("Swap(uint16,uint256,address,uint256,uint256,uint256,uint256,uint256)")
	EventHashPoolSwapRemote      = contract.EventHash("SwapRemote(address,uint256,uint256,uint256)")
	EventHashPoolCreditChainPath = contract.EventHash("CreditChainPath(uint16,uint256,uint256,uint256)")
)

func EthereumChain(chainID uint16) (network.Network, bool) {
	var chain network.Network

	switch chainID {
	case 101:
		chain = network.Ethereum
	case 102:
		chain = network.BinanceSmartChain
	case 106:
		chain = network.Avalanche
	case 109:
		chain = network.Polygon
	case 110:
		chain = network.Arbitrum
	case 111:
		chain = network.Optimism
	case 184:
		chain = network.Base
	case 183:
		chain = network.Linea
	}

	return chain, chain != 0
}

func FactoryAddress(n network.Network) (common.Address, bool) {
	var address common.Address

	switch n {
	case network.Ethereum:
		address = AddressFactoryMainnet
	case network.BinanceSmartChain:
		address = AddressFactoryBinanceSmartChain
	case network.Avalanche:
		address = AddressFactoryAvalanche
	case network.Polygon:
		address = AddressFactoryPolygon
	case network.Arbitrum:
		address = AddressFactoryArbitrumOne
	case network.Optimism:
		address = AddressFactoryOptimism
	case network.Base:
		address = AddressFactoryBase
	case network.Linea:
		address = AddressFactoryLinea
	}

	return address, address != ethereum.AddressGenesis
}

func IsSGETH(address common.Address) bool {
	switch address {
	case
		AddressSGETHMainnet,
		AddressSGETHArbitrumOne,
		AddressSGETHOptimism,
		AddressSGETHBase,
		AddressSGETHLinea:
		return true
	default:
		return false
	}
}
