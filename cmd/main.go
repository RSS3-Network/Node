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

		configFile, err := config.Setup(lo.Must(flags.GetString(flag.KeyConfig)))

		if err != nil {
			return fmt.Errorf("setup config file: %w", err)
		}

		zap.L().Debug("config file loaded", zap.Any("config", configFile))

		// set the logger level based on the environment
		// if the environment variable is not set, use the environment from the config file
		environment := os.Getenv(config.Environment)
		if environment == "" {
			environment = configFile.Environment
			if environment == config.EnvironmentDevelopment {
				zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
			} else {
				zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
			}
		}

		zap.L().Info("logger initialized", zap.String("environment", environment))

		if err = config.HasOneWorker(configFile); err != nil {
			return err
		}

		if configFile.Observability != nil {
			if err := setOpenTelemetry(configFile); err != nil {
				return fmt.Errorf("set open telemetry: %w", err)
			}

			zap.L().Info("open telemetry configured successfully")
		}

		// Init stream client.
		var streamClient stream.Client

		if configFile.Stream != nil && *configFile.Stream.Enable {
			streamClient, err = provider.New(cmd.Context(), configFile.Stream)
			if err != nil {
				return fmt.Errorf("dial stream client: %w", err)
			}

			zap.L().Info("stream client initialized successfully")
		}

		var (
			redisClient    rueidis.Client
			databaseClient database.Client
		)

		module := lo.Must(flags.GetString(flag.KeyModule))

		var networkParamsCaller *vsl.NetworkParamsCaller

		var settlementCaller *vsl.SettlementCaller

		// Broadcaster and RSS Only Node does not need Redis, DB and Network Params Client
		if module != BroadcasterArg && !config.IsRSSComponentOnly(configFile) {
			// Init a Redis client.
			if configFile.Redis == nil {
				return fmt.Errorf("redis configFile is missing")
			}

			redisClient, err = redis.NewClient(*configFile.Redis)
			if err != nil {
				return fmt.Errorf("new redis client: %w", err)
			}

			zap.L().Info("redis client initialized successfully")

			databaseClient, err = dialer.Dial(cmd.Context(), configFile.Database)
			if err != nil {
				return fmt.Errorf("dial database: %w", err)
			}

			zap.L().Info("database client initialized successfully")

			tx, err := databaseClient.Begin(cmd.Context())
			if err != nil {
				return fmt.Errorf("begin transaction: %w", err)
			}

			if err := tx.Migrate(cmd.Context()); err != nil {
				zap.L().Error("database migration failed", zap.Error(err))
				err := tx.Rollback()
				if err != nil {
					return fmt.Errorf("rollback database: %w", err)
				}

				return fmt.Errorf("migrate database: %w", err)
			}

			zap.L().Info("database migration completed successfully")

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

			zap.L().Info("current epoch updated successfully", zap.Int64("epoch", epoch))

			// when start or restart the core, worker or monitor module, it will pull network parameters from VSL and record current epoch
			if _, err = parameter.PullNetworkParamsFromVSL(networkParamsCaller, uint64(epoch)); err != nil {
				return fmt.Errorf("pull network parameters from VSL: %w", err)
			}

			zap.L().Info("network parameters pulled successfully")

			zap.L().Debug("network parameters pulled successfully", zap.Any("parameters", parameter.CurrentNetworkStartBlock))

			for network, blockStart := range parameter.CurrentNetworkStartBlock {
				if blockStart == nil {
					zap.L().Debug("skipping network with nil block start", zap.String("network", network.String()))

					continue // Skip if the start block is not defined.
				}

				// Convert big.Int to int64; safe as long as the value fits in int64.
				blockStartInt64 := blockStart.Block.Int64()

				// Update the current block start for the network in Redis.
				err := parameter.UpdateBlockStart(cmd.Context(), redisClient, network.String(), blockStartInt64)
				if err != nil {
					return fmt.Errorf("update current block start: %w", err)
				}

				zap.L().Debug("block start updated", zap.String("network", network.String()), zap.Int64("block", blockStartInt64))
			}

			if err := tx.Commit(); err != nil {
				return fmt.Errorf("commit transaction: %w", err)
			}

			zap.L().Info("database transaction committed successfully")
		}

		switch module {
		case CoreServiceArg:
			if err := recordFirstStartTime(); err != nil {
				return fmt.Errorf("record first start time: %w", err)
			}

			return runCoreService(cmd.Context(), configFile, databaseClient, redisClient, networkParamsCaller, settlementCaller)
		case WorkerArg:
			return runWorker(cmd.Context(), configFile, databaseClient, streamClient, redisClient)
		case BroadcasterArg:
			return runBroadcaster(cmd.Context(), configFile)
		case MonitorArg:
			return runMonitor(cmd.Context(), configFile, databaseClient, redisClient, networkParamsCaller, settlementCaller)
		}

		return fmt.Errorf("unsupported module %s", lo.Must(flags.GetString(flag.KeyModule)))
	},
}

// recordFirstStartTime set the first start time to the current timestamp in seconds if it is not set.
func recordFirstStartTime() error {
	//	set first start time
	firstStartTime, err := info.GetFirstStartTime()
	if err != nil {
		return fmt.Errorf("get first start time: %w", err)
	}

	if firstStartTime == 0 {
		zap.L().Info("first start time not set, updating to current time")
		//	update first start time to current timestamp in seconds
		err = info.UpdateFirstStartTime(time.Now().Unix())
		if err != nil {
			return fmt.Errorf("update first start time: %w", err)
		}

		zap.L().Info("first start time updated successfully")
	}

	zap.L().Debug("first start time already set", zap.Int64("timestamp", firstStartTime))

	return nil
}

func runCoreService(ctx context.Context, configFile *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	zap.L().Info("initializing core service")

	server := node.NewCoreService(ctx, configFile, databaseClient, redisClient, networkParamsCaller, settlementCaller)

	checkCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	if !config.IsRSSComponentOnly(configFile) {
		go func() {
			if err := node.CheckParams(checkCtx, redisClient, networkParamsCaller, settlementCaller); err != nil {
				fmt.Printf("Error checking parameters: %v\n", err)
			}
		}()
	}

	apiErrChan := make(chan error, 1)
	go func() {
		zap.L().Info("starting core service")
		apiErrChan <- server.Run(ctx)
	}()

	// Set up signal handling
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-stopChan:
		zap.L().Info("shutdown signal received", zap.String("signal", sig.String()))
		fmt.Printf("Shutdown signal received: %v.\n", sig)
	case err := <-apiErrChan:
		zap.L().Error("core service encountered an error", zap.Error(err))
		cancel() // signal all goroutines to stop on error

		return err
	}

	cancel() // ensure cancellation if exiting normally
	zap.L().Info("core service shutdown complete")

	return nil
}

// findModuleByID find and returns the specified worker ID in all components.
func findModuleByID(configFile *config.File, workerID string) (*config.Module, error) {
	zap.L().Debug("searching for module", zap.String("workerID", workerID))

	// find the module in a specific component list
	findInComponent := func(components []*config.Module) (*config.Module, bool) {
		for _, module := range components {
			if strings.EqualFold(module.ID, workerID) {
				zap.L().Debug("module found", zap.String("moduleID", module.ID))
				return module, true
			}
		}

		return nil, false
	}

	// Search in decentralized components
	if module, found := findInComponent(configFile.Component.Decentralized); found {
		return module, nil
	}

	// Search in federated components
	if module, found := findInComponent(configFile.Component.Federated); found {
		return module, nil
	}

	if module, found := findInComponent([]*config.Module{configFile.Component.RSS}); found {
		return module, nil
	}

	return nil, fmt.Errorf("undefined module %s", workerID)
}

func runWorker(ctx context.Context, configFile *config.File, databaseClient database.Client, streamClient stream.Client, redisClient rueidis.Client) error {
	workerID, err := flags.GetString(flag.KeyWorkerID)
	if err != nil {
		return fmt.Errorf("invalid worker id: %w", err)
	}

	zap.L().Info("starting worker", zap.String("workerID", workerID))

	module, err := findModuleByID(configFile, workerID)
	if err != nil {
		return fmt.Errorf("find module by id: %w", err)
	}

	server, err := indexer.NewServer(ctx, module, databaseClient, streamClient, redisClient)

	if err != nil {
		return fmt.Errorf("new indexer server: %w", err)
	}

	zap.L().Info("worker initialized successfully", zap.String("workerID", workerID))

	return server.Run(ctx)
}

func runBroadcaster(ctx context.Context, config *config.File) error {
	zap.L().Info("initializing broadcaster")

	server, err := broadcaster.NewBroadcaster(ctx, config)

	if err != nil {
		return fmt.Errorf("new broadcaster: %w", err)
	}

	zap.L().Info("broadcaster initialized successfully")

	return server.Run(ctx)
}

func runMonitor(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	zap.L().Info("initializing monitor")

	server, err := monitor.NewMonitor(ctx, config, databaseClient, redisClient, networkParamsCaller, settlementCaller)

	if err != nil {
		return fmt.Errorf("new monitor: %w", err)
	}

	zap.L().Info("monitor initialized successfully")

	return server.Run(ctx)
}

func setOpenTelemetry(config *config.File) error {
	zap.L().Debug("setting up OpenTelemetry")
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
			zap.L().Info("starting OpenTelemetry meter server")

			if err := meterServer.Run(*observabilityConfig.Metrics); err != nil {
				zap.L().Error("failed to run telemetry meter server", zap.Error(err))
			}
		}()
	}

	zap.L().Debug("OpenTelemetry setup completed")

	return nil
}

func initializeLogger() {
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
		zap.L().Info("initialized development logger")
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
		zap.L().Info("initialized production logger")
	}
}

func initializePyroscope() {
	// Only use Pyroscope in development environment.
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		zap.L().Debug("checking for Pyroscope configuration")
		// Start Pyroscope agent if the environment variable is set.
		if serverAddress := os.Getenv(config.EnvironmentPyroscopeEndpoint); serverAddress != "" {
			zap.L().Info("initializing Pyroscope", zap.String("serverAddress", serverAddress))
			_, _ = pyroscope.Start(pyroscope.Config{
				ApplicationName: constant.Name,
				ServerAddress:   serverAddress,
				Logger:          zap.L().Sugar(),
				ProfileTypes:    append(pyroscope.DefaultProfileTypes, pyroscope.ProfileGoroutines),
			})
		} else {
			zap.L().Debug("Pyroscope endpoint not configured, skipping initialization")
		}
	}
}

func init() {
	initializeLogger()
	initializePyroscope()

	command.PersistentFlags().String(flag.KeyConfig, "config.yaml", "config file name")
	command.PersistentFlags().String(flag.KeyModule, WorkerArg, "module name")
	command.PersistentFlags().String(flag.KeyWorkerID, "", "worker id")
	zap.L().Debug("command flags initialized")
}

func main() {
	// Flush the logs before the process exits.
	defer lo.Try(zap.L().Sync)

	if err := command.ExecuteContext(context.Background()); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}
