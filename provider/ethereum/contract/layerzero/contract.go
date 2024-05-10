package layerzero

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/network"
)

// https://layerzero.gitbook.io/docs/technical-reference/mainnet/layerzero-labs-relayer.sol-addresses
var (
	AddressRelayerMainnet           = common.HexToAddress("0x902F09715B6303d4173037652FA7377e5b98089E")
	AddressRelayerBinanceSmartChain = common.HexToAddress("0xA27A2cA24DD28Ce14Fb5f5844b59851F03DCf182")
	AddressRelayerAvalanche         = common.HexToAddress("0xCD2E3622d483C7Dc855F72e5eafAdCD577ac78B4")
	AddressRelayerPolygon           = common.HexToAddress("0x75dC8e5F50C8221a82CA6aF64aF811caA983B65f")
	AddressRelayerArbitrum          = common.HexToAddress("0x177d36dBE2271A4DdB2Ad8304d82628eb921d790")
	AddressRelayerOptimism          = common.HexToAddress("0x81E792e5a9003CC1C8BF5569A00f34b65d75b017")
	AddressRelayerGnosis            = common.HexToAddress("0x5B19bd330A84c049b62D5B0FC2bA120217a18C1C")
	AddressRelayerBase              = common.HexToAddress("0xcb566e3B6934Fa77258d68ea18E931fa75e1aaAa")
	AddressRelayerLinea             = common.HexToAddress("0xA658742d33ebd2ce2F0bdFf73515Aa797Fd161D9")

	AddressUltraLightNodeMainnet           = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodeBinanceSmartChain = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodeAvalanche         = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodePolygon           = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodeArbitrumOne       = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodeOptimism          = common.HexToAddress("0x4D73AdB72bC3DD368966edD0f0b2148401A178E2")
	AddressUltraLightNodeBase              = common.HexToAddress("0x38dE71124f7a447a01D67945a51eDcE9FF491251")
	AddressUltraLightNodeLinea             = common.HexToAddress("0x38dE71124f7a447a01D67945a51eDcE9FF491251")
)

func UltraLightNodeAddress(n network.Network) (common.Address, bool) {
	var address common.Address

	switch n {
	case network.Ethereum:
		address = AddressUltraLightNodeMainnet
	case network.BinanceSmartChain:
		address = AddressUltraLightNodeBinanceSmartChain
	case network.Avalanche:
		address = AddressUltraLightNodeAvalanche
	case network.Polygon:
		address = AddressUltraLightNodePolygon
	case network.Arbitrum:
		address = AddressUltraLightNodeArbitrumOne
	case network.Optimism:
		address = AddressUltraLightNodeOptimism
	case network.Base:
		address = AddressUltraLightNodeBase
	case network.Linea:
		address = AddressUltraLightNodeLinea
	}

	return address, address != ethereum.AddressGenesis
}
