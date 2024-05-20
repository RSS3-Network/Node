package network

import (
	"testing"

	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimumResources(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		network  network.Network
		worker   worker.Worker
		expected MinimumResource
	}{
		// High-demand-workers + high-demand-networks
		{
			name:     "Arweave Momoka Resource",
			network:  network.Arweave,
			worker:   worker.Momoka,
			expected: baseResource.Mul(8),
		},
		{
			name:    "Polygon Core Resource",
			network: network.Polygon,
			worker:  worker.Core,
			expected: MinimumResource{
				CPUCore:    2.0,
				MemoryInGb: 2.0,
				SizeInGb:   parameter.NetworkCoreSizePerMonth[network.Polygon],
			},
		},

		// High-demand-workers + low-demand-networks
		{
			name:    "High Demand Worker - Low Demand Network",
			network: network.Linea,
			worker:  worker.Core,
			expected: MinimumResource{
				CPUCore:    1.0,
				MemoryInGb: 1.0,
				SizeInGb:   parameter.NetworkCoreSizePerMonth[network.Linea],
			},
		},
		{
			name:     "High Demand Worker - Low Demand Network (Different)",
			network:  network.Farcaster,
			worker:   worker.Core,
			expected: baseResource.Mul(4),
		},

		// Mid-demand-workers + high-demand-networks
		{
			name:     "Mid Demand Worker - High Demand Network",
			network:  network.Ethereum,
			worker:   worker.OpenSea,
			expected: baseResource.Mul(4),
		},
		{
			name:     "Mid Demand Worker - High Demand Network (Different)",
			network:  network.Arbitrum,
			worker:   worker.Curve,
			expected: baseResource.Mul(4),
		},

		// Mid-demand-workers + low-demand-networks
		{
			name:     "Mid Demand Worker - Low Demand Network",
			network:  network.SatoshiVM,
			worker:   worker.Uniswap,
			expected: baseResource.Mul(2),
		},
		{
			name:     "Mid Demand Worker - Low Demand Network (Different)",
			network:  network.Linea,
			worker:   worker.Stargate,
			expected: baseResource.Mul(2),
		},

		// Low-demand-workers + high-demand-networks
		{
			name:     "Low Demand Worker - High Demand Network",
			network:  network.Ethereum,
			worker:   worker.RSS3,
			expected: baseResource.Mul(2),
		},
		{
			name:     "Low Demand Worker - High Demand Network (Different)",
			network:  network.Polygon,
			worker:   worker.Lens,
			expected: baseResource.Mul(2),
		},

		// Low-demand-workers + low-demand-networks
		{
			name:     "Low Demand Worker - Low Demand Network",
			network:  network.Crossbell,
			worker:   worker.Crossbell,
			expected: baseResource,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := calculateMinimumResources(tc.network, tc.worker)
			assert.Equal(t, tc.expected, result)
		})
	}
}
