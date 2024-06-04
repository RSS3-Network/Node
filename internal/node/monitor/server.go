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
	"github.com/rss3-network/protocol-go/schema/network"
)

type Monitor struct {
	config         *config.File
	cron           *cron.Cron
	databaseClient database.Client
	redisClient    rueidis.Client
	clients        map[network.Network]Client
}

func (m *Monitor) Run(ctx context.Context) error {
	_, err := m.cron.AddFunc("@every 15m", func() {
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
func initNetworkClient(m *config.Module) (Client, error) {
	var client Client

	var err error

	switch m.Network.Source() {
	case network.ArweaveSource:
		client, err = NewArweaveClient()
	case network.FarcasterSource:
		client, err = NewFarcasterClient()
	case network.RSSSource:
		client, err = NewRssClient(m.Endpoint, m.Parameters)
	default:
		client, err = NewEthereumClient(m.Endpoint)
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewMonitor creates a new monitor instance.
func NewMonitor(_ context.Context, configFile *config.File, databaseClient database.Client, redisClient rueidis.Client) (*Monitor, error) {
	modules := append(append(configFile.Component.Decentralized, configFile.Component.RSS...), configFile.Component.Federated...)

	clients := make(map[network.Network]Client)

	// init all clients
	for _, m := range modules {
		client, err := initNetworkClient(m)
		if err != nil {
			return nil, fmt.Errorf("init network client: %w", err)
		}

		clients[m.Network] = client
	}

	instance := &Monitor{
		config:         configFile,
		cron:           cron.New(),
		databaseClient: databaseClient,
		redisClient:    redisClient,
		clients:        clients,
	}

	// register router
	return instance, nil
}
