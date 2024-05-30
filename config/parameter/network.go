package parameter

import (
	"math/big"

	"github.com/rss3-network/protocol-go/schema/network"
)

// NumberOfMonthsToCover the number of months that a Node should cover data for
const NumberOfMonthsToCover = 3

var NetworkTolerance = map[network.Network]uint64{
	network.Arbitrum:          1000,
	network.Arweave:           100,
	network.Avalanche:         100,
	network.Base:              100,
	network.BinanceSmartChain: 100,
	network.Crossbell:         500,
	network.Ethereum:          100,
	network.Farcaster:         3600000,
	network.Gnosis:            100,
	network.Linea:             100,
	network.Optimism:          100,
	network.Polygon:           100,
	network.SatoshiVM:         100,
	network.VSL:               100,
}

// NetworkStartBlock FIXME: provide a script to generate blocks given a data
var NetworkStartBlock = map[network.Network]*big.Int{
	network.Arbitrum:          big.NewInt(185724972),
	network.Arweave:           big.NewInt(1374360),
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

// NetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
// The data is calculated based on the average disk space usage during 2024 Q1.
// Actually usage may vary depending on the network's activity.
var NetworkCoreWorkerDiskSpacePerMonth = map[network.Network]uint{
	network.Arbitrum:          26,
	network.Arweave:           0,
	network.Avalanche:         0,
	network.Base:              10,
	network.BinanceSmartChain: 117,
	network.Crossbell:         0,
	network.Ethereum:          51,
	network.Gnosis:            9,
	network.Linea:             31,
	network.Optimism:          25,
	network.Polygon:           153,
	network.SatoshiVM:         1,
	network.VSL:               1,
}
