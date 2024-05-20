package parameter

import (
	"math/big"

	"github.com/rss3-network/protocol-go/schema/network"
)

var NetworkTolerance = map[network.Network]uint64{
	network.Arbitrum:          100,
	network.Arweave:           100,
	network.Avalanche:         100,
	network.Base:              100,
	network.BinanceSmartChain: 100,
	network.Crossbell:         100,
	network.Ethereum:          100,
	network.Farcaster:         3600000,
	network.Gnosis:            100,
	network.Linea:             100,
	network.Optimism:          100,
	network.Polygon:           100,
	network.SatoshiVM:         100,
	network.VSL:               100,
}

var NetworkStartBlock = map[network.Network]*big.Int{
	network.Arbitrum:          big.NewInt(185724972),
	network.Arweave:           big.NewInt(1374361),
	network.Avalanche:         big.NewInt(42301570),
	network.Base:              big.NewInt(11216527),
	network.BinanceSmartChain: big.NewInt(36563564),
	network.Crossbell:         big.NewInt(58846671),
	network.Ethereum:          big.NewInt(19334220),
	network.Gnosis:            big.NewInt(32695982),
	network.Linea:             big.NewInt(2591120),
	network.Optimism:          big.NewInt(116811812),
	network.Polygon:           big.NewInt(54103805),
	network.SatoshiVM:         big.NewInt(60741),
	network.VSL:               big.NewInt(14192),
}

var NetworkCoreSizePerMonth = map[network.Network]float32{
	network.Arbitrum:          25.86,
	network.Arweave:           0,
	network.Avalanche:         0,
	network.Base:              10.08,
	network.BinanceSmartChain: 116.54,
	network.Crossbell:         0,
	network.Ethereum:          50.57,
	network.Gnosis:            8.46,
	network.Linea:             30.70,
	network.Optimism:          24.29,
	network.Polygon:           153.01,
	network.SatoshiVM:         0.04,
	network.VSL:               0.37,
}
