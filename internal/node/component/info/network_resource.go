package info

import (
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

type MinimumResource struct {
	CPUCore       float32 `json:"cpu_core"`
	MemoryInGb    float32 `json:"memory_in_gb"`
	DiskSpaceInGb uint    `json:"disk_space_in_gb"`
}

// Mul multiplies the resource by a given multiplier
func (r MinimumResource) Mul(multiplier float32) MinimumResource {
	return MinimumResource{
		CPUCore:       r.CPUCore * multiplier,
		MemoryInGb:    r.MemoryInGb * multiplier,
		DiskSpaceInGb: r.DiskSpaceInGb,
	}
}

// baseResource includes the base resource requirement for a worker.
var baseResource = MinimumResource{
	CPUCore:    0.25,
	MemoryInGb: 0.25,
	// 1GB is the default disk space for a non-core worker for a month
	// this is a ballpark figure as it's impossible to calculate the exact disk space required for all non-core workers
	DiskSpaceInGb: 1 * parameter.NumberOfMonthsToCover,
}

// set a list of multipliers for network and worker
const (
	highDemandNetworkMultiplier = 2.0
	highDemandWorkerMultiplier  = 4.0
	midDemandWorkerMultiplier   = 2.0
)

// defines the demand level of a network
var (
	highDemandNetworks = []network.Network{network.Ethereum, network.Polygon, network.Arbitrum, network.Base, network.Gnosis, network.BinanceSmartChain, network.Optimism, network.Arweave, network.Farcaster}
	highDemandWorkers  = []worker.Worker{decentralized.Core, decentralized.Momoka}
	midDemandWorkers   = []worker.Worker{decentralized.Uniswap, decentralized.OpenSea, decentralized.Stargate, decentralized.Curve}
)

// calculateMinimumResources calculates the minimum resources based on network and worker.
func calculateMinimumResources(n network.Network, w worker.Worker) MinimumResource {
	resource := baseResource

	if lo.Contains(highDemandNetworks, n) {
		resource = resource.Mul(highDemandNetworkMultiplier)
	}

	switch {
	case lo.Contains(highDemandWorkers, w):
		resource = resource.Mul(highDemandWorkerMultiplier)
	case lo.Contains(midDemandWorkers, w):
		resource = resource.Mul(midDemandWorkerMultiplier)
	}

	if w.Name() == decentralized.Core.Name() {
		resource.DiskSpaceInGb = parameter.CurrentNetworkCoreWorkerDiskSpacePerMonth[n] * parameter.NumberOfMonthsToCover
	}

	return resource
}
