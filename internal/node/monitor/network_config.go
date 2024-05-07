package monitor

import "github.com/rss3-network/protocol-go/schema/network"

var NetworkTorlerance = map[network.Network]uint64{
	network.Arweave:           100,
	network.Ethereum:          100,
	network.Polygon:           100,
	network.VSL:               100,
	network.Base:              100,
	network.Crossbell:         100,
	network.BinanceSmartChain: 100,
	network.Avalanche:         100,
	network.Gnosis:            100,
	network.Linea:             100,
	network.Optimism:          100,
	network.SatoshiVM:         100,
	network.Arbitrum:          100,
	network.Farcaster:         100,
}
