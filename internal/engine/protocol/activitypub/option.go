package activitypub

import (
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/config/parameter"
	"github.com/rss3-network/node/v2/provider/activitypub/mastodon"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.uber.org/zap"
)

// Option represents the configuration options for the ActivityPub client.
type Option struct {
	RelayURLList   []string `json:"relay_url_list"`
	Port           int64    `json:"port"`
	TimestampStart int64    `json:"timestamp_start" mapstructure:"timestamp_start"`
}

// NewOption creates a new Option instance from the provided parameters.
func NewOption(n network.Network, parameters *config.Parameters, isMonitor bool) (*Option, error) {
	zap.L().Info("parameters:", zap.Any("parameters", parameters))

	// Initialize with default values
	option := Option{
		RelayURLList: mastodon.DefaultRelayURLList,
		Port:         mastodon.DefaultServerPort,
	}

	// Set timestamp for non-monitor mode
	if !isMonitor {
		option.TimestampStart = parameter.CurrentNetworkStartBlock[n].Timestamp
	}

	// If no parameters provided, return defaults
	if parameters == nil {
		return &option, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	// Apply defaults if RelayURLList or Port are not set
	if len(option.RelayURLList) == 0 {
		option.RelayURLList = mastodon.DefaultRelayURLList
		zap.L().Info("relay URL list not specified, using default", zap.Strings("defaultRelayURLList", mastodon.DefaultRelayURLList))
	}

	if option.Port == 0 {
		option.Port = mastodon.DefaultServerPort
		zap.L().Info("port not specified, using default", zap.Int64("defaultPort", mastodon.DefaultServerPort))
	}

	zap.L().Info("option:", zap.Any("option", option))

	return &option, nil
}
