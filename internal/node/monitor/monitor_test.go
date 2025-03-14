package monitor_test

import (
	"context"
	"testing"
	"time"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/config/parameter"
	"github.com/rss3-network/node/v2/internal/node/monitor"
	"github.com/rss3-network/node/v2/provider/ethereum/endpoint"
	redisx "github.com/rss3-network/node/v2/provider/redis"
	"github.com/rss3-network/node/v2/schema/worker"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/node/v2/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/stretchr/testify/require"
)

func TestMonitor(t *testing.T) {
	t.Parallel()

	mockNetworkTolerance()

	type arguments struct {
		config        *config.File
		currentState  monitor.CheckpointState
		targetState   uint64
		lastState     uint64
		latestState   uint64
		initialStatus worker.Status
	}

	testcases := []struct {
		name      string
		source    network.Protocol
		arguments arguments
		want      worker.Status
		wantError require.ErrorAssertionFunc
	}{
		// Ethereum Worker
		{
			name:   "Ethereum Worker Ready Status -> Indexing Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				latestState:   19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Ready Status -> Ready Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusReady,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Ready Status -> Ready Status (With Target Param)",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusReady,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				targetState: 19800000,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Indexing Status -> Ready Status (With Target Param)",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusIndexing,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				targetState: 19800000,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Indexing Status -> Indexing Status (With Target Param)",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusIndexing,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				targetState: 19800001,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Indexing Status -> Unhealthy Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusIndexing,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				lastState:   19800000,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Indexing Status -> Indexing Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusIndexing,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				lastState:   19800000 - 1,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Unhealthy Status -> Indexing Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusUnhealthy,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				lastState:   19800000 - 1,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Indexing Status -> Ready Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusIndexing,
				currentState: monitor.CheckpointState{
					BlockNumber: 19800000,
				},
				lastState:   19800000 - 1,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Ethereum Worker Unknown Status -> Indexing Status",
			source: network.EthereumProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "ethereum-core",
								Network: network.Ethereum,
								Worker:  decentralized.Core,
								Endpoint: config.Endpoint{
									URL: endpoint.MustGet(network.Ethereum),
								},
							},
						},
					},
				},
				initialStatus: worker.StatusUnknown,
				currentState: monitor.CheckpointState{
					BlockNumber: 1,
				},
				lastState:   0,
				latestState: 19800000 + parameter.CurrentNetworkTolerance[network.Ethereum] + 1,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		// Arweave Worker
		{
			name:   "Arweave Worker Ready Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Ready Status -> Ready Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Ready Status -> Ready Status (With Target Param)",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				targetState:   1420000,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Indexing Status -> Ready Status (With Target Param)",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				targetState:   1420000,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Indexing Status -> Indexing Status (With Target Param)",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				targetState:   1420001,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Indexing Status -> Unhealthy Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Indexing Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000 - 1,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Unhealthy Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000 - 1,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusUnhealthy,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Indexing Status -> Ready Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000 - 1,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Worker Unknown Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-mirror",
								Network: network.Arweave,
								Worker:  decentralized.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1,
				},
				lastState:     0,
				latestState:   1420000 + parameter.CurrentNetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusUnknown,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		{
			name:   "Arweave Momoka Worker Ready Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Ready Status -> Ready Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 - 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Indexing Status -> Unhealthy Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				lastState:     1637596118633,
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Indexing Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				lastState:     1637596118633 - 1,
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 + 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Unhealthy Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				lastState:     1637596118633 - 1,
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 + 1,
				initialStatus: worker.StatusUnhealthy,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Indexing Status -> Ready Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1637596118633,
				},
				lastState:     1637596118633 - 1,
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Arweave Momoka Worker Unknown Status -> Indexing Status",
			source: network.ArweaveProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "arweave-momoka",
								Network: network.Arweave,
								Worker:  decentralized.Momoka,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockTimestamp: 1,
				},
				lastState:     0,
				latestState:   1637596118633 + parameter.CurrentNetworkTolerance[network.Arweave]*120000 + 1,
				initialStatus: worker.StatusUnknown,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		// Farcaster Worker
		{
			name:   "Farcaster Worker Ready Status -> Indexing Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Ready Status -> Ready Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Indexing Status -> Unhealthy Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				lastState:     1714972833273,
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Indexing Status -> Indexing Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				lastState:     1714972833273 - 1,
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] + 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Unhealthy Status -> Indexing Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				lastState:     1714972833273 - 1,
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] + 1,
				initialStatus: worker.StatusUnhealthy,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Indexing Status -> Ready Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				lastState:     1714972833273 - 1,
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Farcaster Worker Unhealthy Status -> Indexing Status",
			source: network.FarcasterProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Decentralized: []*config.Module{
							{
								ID:      "farcaster-core",
								Network: network.Farcaster,
								Worker:  decentralized.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					FarcastState: monitor.FarcastState{
						EventID:           432183841886217,
						CastsBackfill:     true,
						ReactionsBackfill: true,
					},
				},
				lastState:     0,
				latestState:   1714972833273 + parameter.CurrentNetworkTolerance[network.Farcaster] + 1,
				initialStatus: worker.StatusUnknown,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		// ActivityPub (Mastodon)
		{
			name:   "Mastodon Worker Ready Status -> Ready Status",
			source: network.ActivityPubProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Federated: []*config.Module{
							{
								ID:      "mastodon-core",
								Network: network.Mastodon,
								Worker:  federated.Mastodon,
								Parameters: &config.Parameters{
									"relay_url_list": []string{
										"https://relay.fedi.buzz/instance/mastodon.social",
									},
									"port": 8181,
								},
							},
						},
					},
				},
				currentState:  monitor.CheckpointState{},
				latestState:   uint64(time.Now().Unix()),
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name:   "Mastodon Worker Ready Status -> Unhealthy Status",
			source: network.ActivityPubProtocol,
			arguments: arguments{
				config: &config.File{
					Component: &config.Component{
						Federated: []*config.Module{
							{
								ID:      "mastodon-core",
								Network: network.Mastodon,
								Worker:  federated.Mastodon,
								Parameters: &config.Parameters{
									"relay_url_list": []string{
										"https://relay.wszz/instance/mast.ocial",
									},
								},
							},
						},
					},
				},
				currentState:  monitor.CheckpointState{},
				latestState:   uint64(time.Now().Unix()),
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
	}

	// Start Redis container
	preset := redis.Preset(
		redis.WithVersion("6.0.9"),
	)

	container, err := gnomock.Start(preset)
	require.NoError(t, err)

	t.Cleanup(func() {
		require.NoError(t, gnomock.Stop(container))
	})

	// Connect to Redis
	redisClient, err := redisx.NewClient(config.Redis{
		Endpoint: container.DefaultAddress(),
	})
	require.NoError(t, err)

	// They share the same Redis container, so we can't run them in parallel
	// nolint:paralleltest
	for _, testcase := range testcases {
		testcase := testcase

		switch testcase.source {
		case network.FarcasterProtocol, network.ArweaveProtocol, network.EthereumProtocol, network.NearProtocol:
			t.Run(testcase.name, func(t *testing.T) {
				ctx := context.Background()

				instance, err := monitor.NewMonitor(ctx, testcase.arguments.config, nil, redisClient, nil, nil)
				require.NoError(t, err)

				// update worker status to initial status
				err = instance.UpdateWorkerStatusByID(ctx, testcase.arguments.config.Component.Decentralized[0].ID, testcase.arguments.initialStatus.String())
				require.NoError(t, err)

				// update worker progress
				err = instance.UpdateWorkerProgress(ctx, testcase.arguments.config.Component.Decentralized[0].ID, monitor.ConstructWorkerProgress(testcase.arguments.lastState, testcase.arguments.targetState, testcase.arguments.latestState, 0))
				require.NoError(t, err)

				// run monitor
				err = instance.MonitorMockWorkerStatus(ctx, testcase.arguments.currentState, testcase.arguments.targetState, testcase.arguments.latestState)
				require.NoError(t, err)

				// check final worker status
				status := instance.GetWorkerStatusByID(ctx, testcase.arguments.config.Component.Decentralized[0].ID)
				require.Equal(t, testcase.want, status)
			})
		case network.ActivityPubProtocol:
			t.Run(testcase.name, func(t *testing.T) {
				ctx := context.Background()

				instance, err := monitor.NewMonitor(ctx, testcase.arguments.config, nil, redisClient, nil, nil)
				require.NoError(t, err)

				// update worker status to initial status
				err = instance.UpdateWorkerStatusByID(ctx, testcase.arguments.config.Component.Federated[0].ID, testcase.arguments.initialStatus.String())
				require.NoError(t, err)

				// run monitor
				err = instance.MonitorMockWorkerStatus(ctx, testcase.arguments.currentState, testcase.arguments.targetState, testcase.arguments.latestState)
				require.NoError(t, err)

				// check final worker status
				status := instance.GetWorkerStatusByID(ctx, testcase.arguments.config.Component.Federated[0].ID)
				require.Equal(t, testcase.want, status)
			})
		default:
		}
	}
}

func mockNetworkTolerance() {
	parameter.CurrentNetworkTolerance = parameter.NetworkTolerance{
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
		network.Near:              1000,
		network.Mastodon:          7200000,
	}
}
