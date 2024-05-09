package monitor_test

import (
	"context"
	"testing"

	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/redis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	redisx "github.com/rss3-network/node/provider/redis"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/stretchr/testify/require"
)

func TestMonitor(t *testing.T) {
	t.Parallel()

	type arguments struct {
		config        *config.File
		currentState  monitor.CheckpointState
		lastState     uint64
		latestState   uint64
		initialStatus worker.Status
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      worker.Status
		wantError require.ErrorAssertionFunc
	}{
		// Ethereum Worker
		{
			name: "Ethereum Worker Ready Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Ethereum,
								Worker:  worker.Core,
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
				latestState:   19800000 + monitor.NetworkTolerance[network.Ethereum] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Ethereum Worker Ready Status -> Ready Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Ethereum,
								Worker:  worker.Core,
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
				latestState: 19800000 + monitor.NetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name: "Ethereum Worker Indexing Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Ethereum,
								Worker:  worker.Core,
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
				latestState: 19800000 + monitor.NetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Ethereum Worker Indexing Status -> Indexing Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Ethereum,
								Worker:  worker.Core,
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
				latestState: 19800000 + monitor.NetworkTolerance[network.Ethereum] - 1,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		// Arweave Worker
		{
			name: "Arweave Worker Ready Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Arweave,
								Worker:  worker.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				latestState:   1420000 + monitor.NetworkTolerance[network.Arweave] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Arweave Worker Ready Status -> Ready Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Arweave,
								Worker:  worker.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				latestState:   1420000 + monitor.NetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name: "Arweave Worker Indexing Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Arweave,
								Worker:  worker.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000,
				latestState:   1420000 + monitor.NetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Arweave Worker Indexing Status -> Indexing Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Arweave,
								Worker:  worker.Mirror,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					BlockHeight: 1420000,
				},
				lastState:     1420000 - 1,
				latestState:   1420000 + monitor.NetworkTolerance[network.Arweave] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
			wantError: require.NoError,
		},

		// Farcaster Worker
		{
			name: "Farcaster Worker Ready Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Farcaster,
								Worker:  worker.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					EventID:          432183841886217,
					CastsBackfill:    true,
					ReactionBackfill: true,
				},
				latestState:   1714972833273 + monitor.NetworkTolerance[network.Farcaster] + 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Farcaster Worker Ready Status -> Ready Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Farcaster,
								Worker:  worker.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					EventID:          432183841886217,
					CastsBackfill:    true,
					ReactionBackfill: true,
				},
				latestState:   1714972833273 + monitor.NetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusReady,
			},
			want:      worker.StatusReady,
			wantError: require.NoError,
		},
		{
			name: "Farcaster Worker Indexing Status -> Unhealthy Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Farcaster,
								Worker:  worker.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					EventID:          432183841886217,
					CastsBackfill:    true,
					ReactionBackfill: true,
				},
				lastState:     1714972833273,
				latestState:   1714972833273 + monitor.NetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusUnhealthy,
			wantError: require.NoError,
		},
		{
			name: "Farcaster Worker Indexing Status -> Indexing Status",
			arguments: arguments{
				config: &config.File{
					Node: &config.Node{
						Decentralized: []*config.Module{
							{
								Network: network.Farcaster,
								Worker:  worker.Core,
							},
						},
					},
				},
				currentState: monitor.CheckpointState{
					EventID:          432183841886217,
					CastsBackfill:    true,
					ReactionBackfill: true,
				},
				lastState:     1714972833273 - 1,
				latestState:   1714972833273 + monitor.NetworkTolerance[network.Farcaster] - 1,
				initialStatus: worker.StatusIndexing,
			},
			want:      worker.StatusIndexing,
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
		Endpoints: []string{
			container.DefaultAddress(),
		},
	})
	require.NoError(t, err)

	// They share the same Redis container, so we can't run them in parallel
	// nolint:paralleltest
	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			ctx := context.Background()

			instance, err := monitor.NewMonitor(ctx, testcase.arguments.config, nil, redisClient)
			require.NoError(t, err)

			// update worker status to initial status
			err = instance.UpdateWorkerStatus(ctx, testcase.arguments.config.Node.Decentralized[0].Network.String(), testcase.arguments.config.Node.Decentralized[0].Worker.String(), testcase.arguments.initialStatus.String())
			require.NoError(t, err)

			// update worker progress
			err = instance.UpdateWorkerProgress(ctx, testcase.arguments.config.Node.Decentralized[0].ID(), testcase.arguments.lastState)
			require.NoError(t, err)

			// run monitor
			err = instance.MonitorMockWorkerStatus(ctx, testcase.arguments.currentState, testcase.arguments.latestState)
			require.NoError(t, err)

			// check final worker status
			status := instance.GetWorkerStatus(ctx, testcase.arguments.config.Node.Decentralized[0].Network.String(), testcase.arguments.config.Node.Decentralized[0].Worker.String())
			require.Equal(t, testcase.want, status)
		})
	}
}
