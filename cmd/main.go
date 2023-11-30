package main

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/config/flag"
	"github.com/naturalselectionlabs/rss3-node/internal/constant"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer"
	"github.com/naturalselectionlabs/rss3-node/internal/node"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var command = cobra.Command{
	Use:           constant.Name,
	Version:       constant.BuildVersion(),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		config, err := config.Setup(lo.Must(cmd.PersistentFlags().GetString(flag.KeyConfig)))
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

		for _, nodeConfig := range config.Node {
			if nodeConfig.Name != lo.Must(cmd.PersistentFlags().GetString(flag.KeyNodeName)) {
				continue
			}

			server, err := node.NewServer(nodeConfig, databaseClient)
			if err != nil {
				return fmt.Errorf("build node server: %w", err)
			}

			return server.Run(cmd.Context())
		}

		return fmt.Errorf("unsupported node %s", viper.GetString(flag.KeyNodeName))
	},
}

func initializeLogger() {
	if viper.GetString(config.Environment) == config.EnvironmentDevelopment {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	}
}

func init() {
	initializeLogger()

	command.PersistentFlags().String(flag.KeyConfig, "./deploy/config.development.yaml", "config file path")
	command.PersistentFlags().String(flag.KeyNodeName, "ethereum.mainnet.fallback:default", "node name")
}

func main() {
	if err := command.ExecuteContext(context.Background()); err != nil {
		zap.L().Fatal("execute command", zap.Error(err))
	}
}
