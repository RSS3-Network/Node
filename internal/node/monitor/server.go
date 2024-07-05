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
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Monitor struct {
	config              *config.File
	cron                *cron.Cron
	databaseClient      database.Client
	redisClient         rueidis.Client
	networkParamsCaller *vsl.NetworkParamsCaller
	settlementCaller    *vsl.SettlementCaller
	clients             map[network.Network]Client
}

func (m *Monitor) Run(ctx context.Context) error {
	_, err := m.cron.AddFunc("@every 15m", func() {
		if err := parameter.CheckParamsTask(ctx, m.redisClient, m.networkParamsCaller, m.settlementCaller); err != nil {
			return
		}

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

// initNetworkClient initializes all network clients.
func initNetworkClient(m *config.Module) (Client, error) {
	var client Client

	var err error

	switch m.Network.Source() {
	case network.ArweaveSource:
		client, err = NewArweaveClient()
	case network.FarcasterSource:
		client, err = NewFarcasterClient()
	case network.RSSSource:
		client, err = NewRssClient(m.EndpointID, m.Parameters)
	case network.EthereumSource:
		client, err = NewEthereumClient(m.Endpoint)
	default:
		return nil, fmt.Errorf("unsupported network source: %s", m.Network)
	}

	if err != nil {
		return nil, err
	}

	return client, nil
}

// NewMonitor creates a new monitor instance.
func NewMonitor(_ context.Context, configFile *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) (*Monitor, error) {
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
		config:              configFile,
		cron:                cron.New(),
		databaseClient:      databaseClient,
		redisClient:         redisClient,
		clients:             clients,
		networkParamsCaller: networkParamsCaller,
		settlementCaller:    settlementCaller,
	}

	// register router
	return instance, nil
}
