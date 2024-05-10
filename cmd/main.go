package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/grafana/pyroscope-go"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/flag"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer"
	"github.com/rss3-network/node/internal/node"
	"github.com/rss3-network/node/internal/node/broadcaster"
	"github.com/rss3-network/node/internal/node/indexer"
	"github.com/rss3-network/node/internal/node/monitor"
	"github.com/rss3-network/node/internal/stream"
	"github.com/rss3-network/node/internal/stream/provider"
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

		if err := setOpenTelemetry(config); err != nil {
			return fmt.Errorf("set open telemetry: %w", err)
		}

		// Init stream client.
		var streamClient stream.Client

		if *config.Stream.Enable {
			streamClient, err = provider.New(cmd.Context(), config.Stream)
			if err != nil {
				return fmt.Errorf("dial stream client: %w", err)
			}
		}

		// Init redis client
		redisClient, err := redis.NewClient(*config.Redis)
		if err != nil {
			return fmt.Errorf("new redis client: %w", err)
		}

		var databaseClient database.Client

		module := lo.Must(flags.GetString(flag.KeyModule))

		if module != node.Broadcaster {
			// Dial and migrate database.
			databaseClient, err = dialer.Dial(cmd.Context(), config.Database)
			if err != nil {
				return fmt.Errorf("dial database: %w", err)
			}

			if err := databaseClient.Migrate(cmd.Context()); err != nil {
				return fmt.Errorf("migrate database: %w", err)
			}
		}

		switch module {
		case node.CoreService:
			return runCoreService(cmd.Context(), config, databaseClient, redisClient)
		case node.Worker:
			return runWorker(cmd.Context(), config, databaseClient, streamClient, redisClient)
		case node.Broadcaster:
			return runBroadcaster(cmd.Context(), config)
		case node.Monitor:
			return runMonitor(cmd.Context(), config, databaseClient, redisClient)
		}

		return fmt.Errorf("unsupported module %s", lo.Must(flags.GetString(flag.KeyModule)))
	},
}

func runCoreService(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client) error {
	server := node.NewCoreService(ctx, config, databaseClient, redisClient)

	return server.Run(ctx)
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

func runMonitor(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client) error {
	server, err := monitor.NewMonitor(ctx, config, databaseClient, redisClient)
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
	command.PersistentFlags().String(flag.KeyModule, node.Worker, "module name")
	command.PersistentFlags().String(flag.KeyWorkerID, "", "indexer id")
}

func main() {
	// Flush the logs before the process exits.
	defer lo.Try(zap.L().Sync)

	if err := command.ExecuteContext(context.Background()); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}
