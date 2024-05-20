package network

import (
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

type MinimumResource struct {
	CPUCore    float32 `json:"CPU_core"`
	MemoryInGb float32 `json:"memory_in_gb"`
	SizeInGb   float32 `json:"size_in_gb"`
}

// Mul multiplies the resource by a given multiplier
func (r MinimumResource) Mul(multiplier float32) MinimumResource {
	return MinimumResource{
		CPUCore:    r.CPUCore * multiplier,
		MemoryInGb: r.MemoryInGb * multiplier,
	}
}

// baseResource includes the base resource requirement for a worker.
var baseResource = MinimumResource{
	CPUCore:    0.25,
	MemoryInGb: 0.25,
}

// set a list of multipliers for network and worker
const (
	highDemandNetworkMultiplier = 2.0
	highDemandWorkerMultiplier  = 4.0
	midDemandWorkerMultiplier   = 2.0
)

// defines the demand level of a network
var (
	highDemandNetworks = []network.Network{network.Ethereum, network.Polygon, network.Arbitrum, network.Base, network.Gnosis, network.BinanceSmartChain, network.Optimism, network.Arweave}
	highDemandWorkers  = []worker.Worker{worker.Core, worker.Momoka}
	midDemandWorkers   = []worker.Worker{worker.Uniswap, worker.OpenSea, worker.Stargate, worker.Curve}
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

	return resource
}
