package network

import (
	"testing"

	"github.com/rss3-network/node/config/parameter"
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
			name:     "Arweave Momoka",
			network:  network.Arweave,
			worker:   worker.Momoka,
			expected: baseResource.Mul(8),
		},
		{
			name:    "Polygon Core",
			network: network.Polygon,
			worker:  worker.Core,
			expected: MinimumResource{
				CPUCore:       2.0,
				MemoryInGb:    2.0,
				DiskSpaceInGb: parameter.NetworkCoreWorkerDiskSpacePerMonth[network.Polygon],
			},
		},

		// High-demand-workers + low-demand-networks
		{
			name:    "Linea Core",
			network: network.Linea,
			worker:  worker.Core,
			expected: MinimumResource{
				CPUCore:       1.0,
				MemoryInGb:    1.0,
				DiskSpaceInGb: parameter.NetworkCoreWorkerDiskSpacePerMonth[network.Linea],
			},
		},
		// FIXME: the disk space required for Facaster core worker is unknown
		// {
		//	name:     "High Demand Worker - Low Demand Network (Different)",
		//	network:  network.Farcaster,
		//	worker:   worker.Core,
		//	expected: baseResource.Mul(4),
		// },

		// Mid-demand-workers + high-demand-networks
		{
			name:     "Ethereum OpenSea",
			network:  network.Ethereum,
			worker:   worker.OpenSea,
			expected: baseResource.Mul(4),
		},
		{
			name:     "Arbitrum Curve",
			network:  network.Arbitrum,
			worker:   worker.Curve,
			expected: baseResource.Mul(4),
		},

		// Mid-demand-workers + low-demand-networks
		{
			name:     "SatoshiVM Uniswap",
			network:  network.SatoshiVM,
			worker:   worker.Uniswap,
			expected: baseResource.Mul(2),
		},
		{
			name:     "Linea Stargate",
			network:  network.Linea,
			worker:   worker.Stargate,
			expected: baseResource.Mul(2),
		},

		// Low-demand-workers + high-demand-networks
		{
			name:     "Ethereum RSS3",
			network:  network.Ethereum,
			worker:   worker.RSS3,
			expected: baseResource.Mul(2),
		},
		{
			name:     "Polygon Lens",
			network:  network.Polygon,
			worker:   worker.Lens,
			expected: baseResource.Mul(2),
		},

		// Low-demand-workers + low-demand-networks
		{
			name:     "Crossbell",
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
