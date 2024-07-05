package parameter

import (
	"math/big"

	"github.com/rss3-network/protocol-go/schema/network"
)

// NumberOfMonthsToCover the number of months that a Node should cover data for
const NumberOfMonthsToCover = 3

type NetworkTolerance map[string]uint64
type NetworkStartBlock map[string]*big.Int
type NetworkCoreWorkerDiskSpacePerMonth map[string]uint

// CurrentNetworkTolerance should be updated each epoch from vsl
var CurrentNetworkTolerance = NetworkTolerance{
	network.Arbitrum.String():          1000,
	network.Arweave.String():           100,
	network.Avalanche.String():         100,
	network.Base.String():              100,
	network.BinanceSmartChain.String(): 100,
	network.Crossbell.String():         500,
	network.Ethereum.String():          100,
	network.Farcaster.String():         3600000,
	network.Gnosis.String():            100,
	network.Linea.String():             100,
	network.Optimism.String():          100,
	network.Polygon.String():           100,
	network.SatoshiVM.String():         100,
	network.VSL.String():               100,
}

// CurrentNetworkStartBlock should be updated each epoch from vsl
var CurrentNetworkStartBlock = NetworkStartBlock{
	network.Arbitrum.String():          big.NewInt(185724972),
	network.Arweave.String():           big.NewInt(1374360),
	network.Avalanche.String():         big.NewInt(42301570),
	network.Base.String():              big.NewInt(11216527),
	network.BinanceSmartChain.String(): big.NewInt(36563564),
	network.Crossbell.String():         big.NewInt(58846671),
	network.Ethereum.String():          big.NewInt(19334220),
	network.Gnosis.String():            big.NewInt(32695982),
	network.Linea.String():             big.NewInt(2591120),
	network.Optimism.String():          big.NewInt(116811812),
	network.Polygon.String():           big.NewInt(54103805),
	network.SatoshiVM.String():         big.NewInt(60741),
	network.VSL.String():               big.NewInt(14192),
}

// CurrentNetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
// The data is calculated based on the average disk space usage during 2024 Q1.
// Actually usage may vary depending on the network's activity.
var CurrentNetworkCoreWorkerDiskSpacePerMonth = NetworkCoreWorkerDiskSpacePerMonth{
	network.Arbitrum.String():          26,
	network.Arweave.String():           0,
	network.Avalanche.String():         0,
	network.Base.String():              10,
	network.BinanceSmartChain.String(): 117,
	network.Crossbell.String():         0,
	network.Ethereum.String():          51,
	network.Gnosis.String():            9,
	network.Linea.String():             31,
	network.Optimism.String():          25,
	network.Polygon.String():           153,
	network.SatoshiVM.String():         1,
	network.VSL.String():               1,
	network.Farcaster.String():         50,
}
