package monitor

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/rueidis"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/config/parameter"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/vsl"
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
		zap.L().Info("starting monitor service")

		// Start the monitor cron job.
		monitorWorkerStatus, err := NewCronJob(m.redisClient, MonitorWorkerStatusJob, 10*time.Minute)
		if err != nil {
			return fmt.Errorf("new cron job: %w", err)
		}

		if err = monitorWorkerStatus.AddFunc(ctx, "@every 5m", func() {
			zap.L().Debug("running monitor worker status check")
			if err = parameter.CheckParamsTask(ctx, m.redisClient, m.networkParamsCaller); err != nil {
				zap.L().Error("failed to check network parameters", zap.Error(err))
				return
			}

			if err = m.MonitorWorkerStatus(ctx); err != nil {
				zap.L().Error("failed to monitor worker status", zap.Error(err))
				return
			}
			zap.L().Debug("completed monitor worker status check")
		}); err != nil {
			return fmt.Errorf("add heartbeat cron job: %w", err)
		}

		// Start the database maintenance cron job.
		databaseMaintenance, err := NewCronJob(m.redisClient, DatabaseMaintenanceJob, 5*24*time.Hour)
		if err != nil {
			return fmt.Errorf("new cron job: %w", err)
		}

		if err = databaseMaintenance.AddFunc(ctx, "0 0 0 * * *", func() {
			zap.L().Info("starting database maintenance")
			if err = m.MaintainCoveragePeriod(ctx); err != nil {
				zap.L().Error("maintain coverage period", zap.Error(err))
				return
			}
			zap.L().Info("completed database maintenance")
		}); err != nil {
			return fmt.Errorf("add database maintenance cron job: %w", err)
		}

		defer func() {
			monitorWorkerStatus.Stop()
			databaseMaintenance.Stop()
		}()

		monitorWorkerStatus.Start()
		databaseMaintenance.Start()
		zap.L().Info("monitor service started successfully")
	}

	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopChan

	return nil
}

// initNetworkClient initializes all network clients.
func initNetworkClient(m *config.Module) (Client, error) {
	zap.L().Debug("initializing network client", zap.String("network", m.Network.String()))

	var client Client

	var err error

	switch m.Network.Protocol() {
	case network.ActivityPubProtocol:
		client, err = NewActivityPubClient(m.Network, m.Parameters)
	case network.ArweaveProtocol:
		client, err = NewArweaveClient()
	case network.FarcasterProtocol:
		client, err = NewFarcasterClient()
	case network.EthereumProtocol:
		client, err = NewEthereumClient(m.Endpoint)
	case network.NearProtocol:
		client, err = NewNearClient(m.Endpoint)
	case network.ATProtocol:
		client, err = NewAtprotoClient()
	default:
		return nil, fmt.Errorf("unsupported network protocol: %s", m.Network)
	}

	if err != nil {
		return nil, err
	}

	zap.L().Debug("network client initialized successfully", zap.String("network", m.Network.String()))

	return client, nil
}

// NewMonitor creates a new monitor instance.
func NewMonitor(_ context.Context, configFile *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) (*Monitor, error) {
	zap.L().Debug("creating new monitor instance")

	totalModules := len(configFile.Component.Decentralized) + len(configFile.Component.Federated)
	zap.L().Debug("initializing modules", zap.Int("total_modules", totalModules))

	modules := make([]*config.Module, 0, totalModules)

	modules = append(modules, configFile.Component.Decentralized...)
	modules = append(modules, configFile.Component.Federated...)

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

	zap.L().Info("monitor instance created successfully")

	return instance, nil
}
