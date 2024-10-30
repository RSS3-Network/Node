package activitypub

import (
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
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
	var option Option

	zap.L().Info("parameters:", zap.Any("parameters", parameters))

	if parameters == nil {
		if isMonitor {
			return &Option{
				RelayURLList: mastodon.DefaultRelayURLList,
				Port:         mastodon.DefaultServerPort,
			}, nil
		}

		return &Option{
			RelayURLList:   mastodon.DefaultRelayURLList,
			Port:           mastodon.DefaultServerPort,
			TimestampStart: parameter.CurrentNetworkStartBlock[n].Timestamp,
		}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	// Apply defaults if RelayURLList or Port are not set
	if len(option.RelayURLList) == 0 {
		option.RelayURLList = mastodon.DefaultRelayURLList
		zap.L().Info("RelayURLList not specified, using default", zap.Strings("defaultRelayURLList", mastodon.DefaultRelayURLList))
	}

	if option.Port == 0 {
		option.Port = mastodon.DefaultServerPort
		zap.L().Info("Port not specified, using default", zap.Int64("defaultPort", mastodon.DefaultServerPort))
	}

	if !isMonitor {
		if option.TimestampStart == 0 {
			option.TimestampStart = parameter.CurrentNetworkStartBlock[n].Timestamp
		}
	}

	zap.L().Info("option:", zap.Any("option", option))

	return &option, nil
}
