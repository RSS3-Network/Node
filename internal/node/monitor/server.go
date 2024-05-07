package monitor

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/rueidis"
	"github.com/robfig/cron/v3"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Monitor struct {
	config         *config.File
	cron           *cron.Cron
	databaseClient database.Client
	redisClient    rueidis.Client
	clients        map[network.Network]interface{}
}

func (m *Monitor) Run(ctx context.Context) error {
	_, err := m.cron.AddFunc("@every 10m", func() {
		if err := m.MonitorWorkerStatus(ctx); err != nil {
			return
		}
	})
	if err != nil {
		return fmt.Errorf("add heartbeat cron job: %w", err)
	}

	m.cron.Start()

	stopchan := make(chan os.Signal, 1)

	signal.Notify(stopchan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopchan

	return nil
}

// initNetworkClient initializes arweave, ethereum, and other network clients.
func initNetworkClient(n network.Network, endpoint string) (interface{}, error) {
	var client interface{}

	var err error

	switch n {
	case network.Arweave:
		client, err = arweave.NewClient()
	case network.Arbitrum, network.Avalanche, network.Base, network.BinanceSmartChain, network.Crossbell, network.Ethereum, network.Gnosis, network.Linea, network.Optimism, network.Polygon, network.SatoshiVM, network.VSL:
		client, err = ethereum.Dial(context.Background(), endpoint)
	case network.Farcaster:
		break
	default:
		err = fmt.Errorf("unsupported network: %s", n)
	}

	if err != nil {
		return nil, fmt.Errorf("initializing client for network %s: %w", n, err)
	}

	return client, nil
}

// NewMonitor creates a new monitor instance.
func NewMonitor(_ context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client) (*Monitor, error) {
	instance := &Monitor{
		config:         config,
		cron:           cron.New(),
		databaseClient: databaseClient,
		redisClient:    redisClient,
		clients:        make(map[network.Network]interface{}),
	}

	// init all clients
	for _, w := range config.Node.Decentralized {
		client, err := initNetworkClient(w.Network, w.Endpoint)
		if err != nil {
			return nil, fmt.Errorf("init network client: %w", err)
		}

		instance.clients[w.Network] = client
	}

	// register router
	return instance, nil
}
