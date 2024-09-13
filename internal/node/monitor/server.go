package monitor

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

const (
	MonitorWorkerStatusJob = "worker_status"
	DatabaseMaintenanceJob = "database_maintenance"
)

type Monitor struct {
	config              *config.File
	databaseClient      database.Client
	redisClient         rueidis.Client
	networkParamsCaller *vsl.NetworkParamsCaller
	settlementCaller    *vsl.SettlementCaller
	clients             map[network.Network]Client
}

func (m *Monitor) Run(ctx context.Context) error {
	if m.databaseClient != nil && m.redisClient != nil {
		// Start the monitor cron job.
		monitorWorkerStatus, err := NewCronJob(m.redisClient, MonitorWorkerStatusJob, 10*time.Minute)
		if err != nil {
			return fmt.Errorf("new cron job: %w", err)
		}

		if err = monitorWorkerStatus.AddFunc(ctx, "@every 5m", func() {
			if err = parameter.CheckParamsTask(ctx, m.redisClient, m.networkParamsCaller); err != nil {
				return
			}

			if err = m.MonitorWorkerStatus(ctx); err != nil {
				return
			}
		}); err != nil {
			return fmt.Errorf("add heartbeat cron job: %w", err)
		}

		// Start the database maintenance cron job.
		databaseMaintenance, err := NewCronJob(m.redisClient, DatabaseMaintenanceJob, 5*24*time.Hour)
		if err != nil {
			return fmt.Errorf("new cron job: %w", err)
		}

		if err = databaseMaintenance.AddFunc(ctx, "0 0 0 * * *", func() {
			if err = m.MaintainCoveragePeriod(ctx); err != nil {
				zap.L().Error("maintain coverage period", zap.Error(err))

				return
			}
		}); err != nil {
			return fmt.Errorf("add database maintenance cron job: %w", err)
		}

		defer func() {
			monitorWorkerStatus.Stop()
			databaseMaintenance.Stop()
		}()

		monitorWorkerStatus.Start()
		databaseMaintenance.Start()
	}

	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopChan

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
	totalModules := len(configFile.Component.Decentralized) + len(configFile.Component.Federated)
	if configFile.Component.RSS != nil {
		totalModules++
	}

	modules := make([]*config.Module, 0, totalModules)

	modules = append(modules, configFile.Component.Decentralized...)

	modules = append(modules, configFile.Component.Federated...)

	if configFile.Component.RSS != nil {
		modules = append(modules, configFile.Component.RSS)
	}

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
		databaseClient:      databaseClient,
		redisClient:         redisClient,
		clients:             clients,
		networkParamsCaller: networkParamsCaller,
		settlementCaller:    settlementCaller,
	}

	// register router
	return instance, nil
}
