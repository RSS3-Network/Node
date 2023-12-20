package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/config/flag"
	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/internal/node"
	"github.com/naturalselectionlabs/rss3-node/internal/node/hub"
	"github.com/naturalselectionlabs/rss3-node/internal/node/indexer"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tdewolff/minify/v2/minify"
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

		// Dial and migrate database.
		databaseClient, err := dialer.Dial(cmd.Context(), config.Database)
		if err != nil {
			return fmt.Errorf("dial database: %w", err)
		}

		if err := databaseClient.Migrate(cmd.Context()); err != nil {
			return fmt.Errorf("migrate database: %w", err)
		}

		switch lo.Must(flags.GetString(flag.KeyModule)) {
		case node.Hub:
			return runHub(cmd.Context(), config, databaseClient)
		case node.Indexer:
			return runIndexer(cmd.Context(), config, databaseClient)
		}

		return fmt.Errorf("unsupported module %s", lo.Must(flags.GetString(flag.KeyModule)))
	},
}

func runHub(ctx context.Context, config *config.File, databaseClient database.Client) error {
	server, err := hub.NewServer(ctx, config, databaseClient)
	if err != nil {
		return fmt.Errorf("new server: %w", err)
	}

	return server.Run(ctx)
}

func runIndexer(ctx context.Context, config *config.File, databaseClient database.Client) error {
	parameters, err := minify.JSON(lo.Must(flags.GetString(flag.KeyIndexerParameters)))
	if err != nil {
		return fmt.Errorf("invalid indexer parameters: %w", err)
	}

	network, err := filter.NetworkString(lo.Must(flags.GetString(flag.KeyIndexerNetwork)))
	if err != nil {
		return fmt.Errorf("network string: %w", err)
	}

	worker, err := engine.NameString(lo.Must(flags.GetString(flag.KeyIndexerWorker)))
	if err != nil {
		return fmt.Errorf("worker string: %w", err)
	}

	for _, nodeConfig := range config.Node.Decentralized {
		if nodeConfig.Network == network && nodeConfig.Worker == worker && strings.EqualFold(nodeConfig.Parameters.String(), parameters) {
			server, err := indexer.NewServer(ctx, nodeConfig, databaseClient)
			if err != nil {
				return fmt.Errorf("new server: %w", err)
			}

			return server.Run(ctx)
		}
	}

	return fmt.Errorf("undefined indexer %s.%s.%s", network, worker, parameters)
}

func initializeLogger() {
	if os.Getenv(config.Environment) == config.EnvironmentDevelopment {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	}
}

func init() {
	initializeLogger()

	command.PersistentFlags().String(flag.KeyConfig, "./deploy/config.development.yaml", "config file path")
	command.PersistentFlags().String(flag.KeyModule, node.Indexer, "module name")
	command.PersistentFlags().String(flag.KeyIndexerNetwork, filter.NetworkEthereum.String(), "indexer network")
	command.PersistentFlags().String(flag.KeyIndexerWorker, engine.Fallback.String(), "indexer worker")
	command.PersistentFlags().String(flag.KeyIndexerParameters, "{}", "indexer parameters")
}

func main() {
	if err := command.ExecuteContext(context.Background()); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}
