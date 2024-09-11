package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/grafana/pyroscope-go"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/flag"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/node"
	"github.com/rss3-network/node/internal/node/broadcaster"
	"github.com/rss3-network/node/internal/node/component/info"
	"github.com/rss3-network/node/internal/node/indexer"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/internal/stream"
	"github.com/rss3-network/node/internal/stream/provider"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/node/provider/redis"
	"github.com/rss3-network/node/provider/telemetry"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.uber.org/zap"
)

var flags *pflag.FlagSet

// the list of arguments to listen to
// each argument corresponds to a module
const (
	CoreServiceArg = "core"
	WorkerArg      = "worker"
	BroadcasterArg = "broadcaster"
	MonitorArg     = "monitor"
)

var command = cobra.Command{
	Use:           constant.Name,
	Version:       constant.BuildVersion(),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		flags = cmd.PersistentFlags()

		config, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))
		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		if config.Observability != nil {
			if err := setOpenTelemetry(config); err != nil {
				return fmt.Errorf("set open telemetry: %w", err)
			}
		}

		// Init stream client.
		var streamClient stream.Client

		if config.Stream != nil && *config.Stream.Enable {
			streamClient, err = provider.New(cmd.Context(), config.Stream)
			if err != nil {
				return fmt.Errorf("dial stream client: %w", err)
			}
		}

		var (
			redisClient    rueidis.Client
			databaseClient database.Client
		)

		module := lo.Must(flags.GetString(flag.KeyModule))

		var networkParamsCaller *vsl.NetworkParamsCaller

		var settlementCaller *vsl.SettlementCaller

		// Apply database migrations for all modules except the broadcaster.
		if module != BroadcasterArg && len(config.Component.Decentralized) > 0 {
			databaseClient, err = dialer.Dial(cmd.Context(), config.Database)
			if err != nil {
				return fmt.Errorf("dial database: %w", err)
			}

			if err := databaseClient.Migrate(cmd.Context()); err != nil {
				return fmt.Errorf("migrate database: %w", err)
			}

			// Init redis client
			redisClient, err = redis.NewClient(*config.Redis)
			if err != nil {
				return fmt.Errorf("new redis client: %w", err)
			}

			vslClient, err := parameter.InitVSLClient()
			if err != nil {
				return fmt.Errorf("init vsl client: %w", err)
			}

			chainID, err := vslClient.ChainID(context.Background())
			if err != nil {
				return fmt.Errorf("get chain id: %w", err)
			}

			networkParamsCaller, err = vsl.NewNetworkParamsCaller(vsl.AddressNetworkParams[chainID.Int64()], vslClient)
			if err != nil {
				return fmt.Errorf("new network params caller: %w", err)
			}

			settlementCaller, err = vsl.NewSettlementCaller(vsl.AddressSettlement[chainID.Int64()], vslClient)
			if err != nil {
				return fmt.Errorf("new settlement caller: %w", err)
			}

			epoch, err := parameter.GetCurrentEpochFromVSL(settlementCaller)
			if err != nil {
				return fmt.Errorf("get current epoch: %w", err)
			}

			// save epoch to redis cache
			err = parameter.UpdateCurrentEpoch(cmd.Context(), redisClient, epoch)
			if err != nil {
				return fmt.Errorf("update current epoch: %w", err)
			}

			// when start or restart the core, worker or monitor module, it will pull network parameters from VSL and record current epoch
			if _, err = parameter.PullNetworkParamsFromVSL(networkParamsCaller, uint64(epoch)); err != nil {
				zap.L().Error("pull network parameters from VSL", zap.Error(err))

				return fmt.Errorf("pull network parameters from VSL: %w", err)
			}

			for network, blockStart := range parameter.CurrentNetworkStartBlock {
				if blockStart == nil {
					continue // Skip if the start block is not defined.
				}

				// Convert big.Int to int64; safe as long as the value fits in int64.
				blockStartInt64 := blockStart.Block.Int64()

				// Update the current block start for the network in Redis.
				err := parameter.UpdateBlockStart(cmd.Context(), redisClient, network.String(), blockStartInt64)
				if err != nil {
					return fmt.Errorf("update current block start: %w", err)
				}
			}

			//	set first start time
			firstStartTime, err := info.GetFirstStartTime(cmd.Context(), redisClient)
			if err != nil {
				return fmt.Errorf("get first start time: %w", err)
			}

			if firstStartTime == 0 {
				//	update first start time to current timestamp in seconds
				err = info.UpdateFirstStartTime(cmd.Context(), redisClient, time.Now().Unix())
				if err != nil {
					return fmt.Errorf("update first start time: %w", err)
				}
			}

		}

		switch module {
		case CoreServiceArg:
			return runCoreService(cmd.Context(), config, databaseClient, redisClient, networkParamsCaller, settlementCaller)
		case WorkerArg:
			return runWorker(cmd.Context(), config, databaseClient, streamClient, redisClient)
		case BroadcasterArg:
			return runBroadcaster(cmd.Context(), config)
		case MonitorArg:
			return runMonitor(cmd.Context(), config, databaseClient, redisClient, networkParamsCaller, settlementCaller)
		}

		return fmt.Errorf("unsupported module %s", lo.Must(flags.GetString(flag.KeyModule)))
	},
}

func runCoreService(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	server := node.NewCoreService(ctx, config, databaseClient, redisClient, networkParamsCaller, settlementCaller)

	checkCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		if err := node.CheckParams(checkCtx, redisClient, networkParamsCaller, settlementCaller); err != nil {
			fmt.Printf("Error checking parameters: %v\n", err)
		}
	}()

	apiErrChan := make(chan error, 1)
	go func() {
		apiErrChan <- server.Run(ctx)
	}()

	// Set up signal handling
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stopChan:
		fmt.Printf("Shutdown signal received: %v.\n", sig)
	case err := <-apiErrChan:
		cancel() // signal all goroutines to stop on error
		return err
	}

	cancel() // ensure cancellation if exiting normally

	return nil
}

func runWorker(ctx context.Context, configFile *config.File, databaseClient database.Client, streamClient stream.Client, redisClient rueidis.Client) error {
	workerID, err := flags.GetString(flag.KeyWorkerID)
	if err != nil {
		return fmt.Errorf("invalid worker id: %w", err)
	}

	module, found := lo.Find(configFile.Component.Decentralized, func(module *config.Module) bool {
		return strings.EqualFold(module.ID, workerID)
	})

	if !found {
		return fmt.Errorf("undefined module %s", workerID)
	}

	server, err := indexer.NewServer(ctx, module, databaseClient, streamClient, redisClient)
	if err != nil {
		return fmt.Errorf("new indexer server: %w", err)
	}

	return server.Run(ctx)
}

func runBroadcaster(ctx context.Context, config *config.File) error {
	server, err := broadcaster.NewBroadcaster(ctx, config)
	if err != nil {
		return fmt.Errorf("new broadcaster: %w", err)
	}

	return server.Run(ctx)
}

func runMonitor(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	server, err := monitor.NewMonitor(ctx, config, databaseClient, redisClient, networkParamsCaller, settlementCaller)
	if err != nil {
		return fmt.Errorf("new monitor: %w", err)
	}

	return server.Run(ctx)
}

func setOpenTelemetry(config *config.File) error {
	// Set OpenTelemetry global tracer and meter provider.
	observabilityConfig := config.Observability.OpenTelemetry

	if observabilityConfig.Traces.Enable {
		tracerProvider, err := telemetry.OpenTelemetryTracer(observabilityConfig)
		if err != nil {
			return fmt.Errorf("open telemetry tracer: %w", err)
		}

		otel.SetTracerProvider(tracerProvider)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))
	}

	if observabilityConfig.Metrics.Enable {
		meterProvider, err := telemetry.OpenTelemetryMeter()
		if err != nil {
			return fmt.Errorf("open telemetry meter: %w", err)
		}

		otel.SetMeterProvider(meterProvider)

		meterServer, err := telemetry.OpenTelemetryMeterServer()
		if err != nil {
			return fmt.Errorf("open telemetry meter server: %w", err)
		}

		go func() {
			if err := meterServer.Run(*observabilityConfig.Metrics); err != nil {
				zap.L().Error("failed to run telemetry meter server", zap.Error(err))
			}
		}()
	}

	return nil
}

func initializeLogger() {
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	}
}

func initializePyroscope() {
	// Only use Pyroscope in development environment.
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		// Start Pyroscope agent if the environment variable is set.
		if serverAddress := os.Getenv(config.EnvironmentPyroscopeEndpoint); serverAddress != "" {
			_, _ = pyroscope.Start(pyroscope.Config{
				ApplicationName: constant.Name,
				ServerAddress:   serverAddress,
				Logger:          zap.L().Sugar(),
				ProfileTypes:    append(pyroscope.DefaultProfileTypes, pyroscope.ProfileGoroutines),
			})
		}
	}
}

func init() {
	initializeLogger()
	initializePyroscope()

	command.PersistentFlags().String(flag.KeyConfig, "config.yaml", "config file name")
	command.PersistentFlags().String(flag.KeyModule, WorkerArg, "module name")
	command.PersistentFlags().String(flag.KeyWorkerID, "", "worker id")
}

func main() {
	// Flush the logs before the process exits.
	defer lo.Try(zap.L().Sync)

	if err := command.ExecuteContext(context.Background()); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}
