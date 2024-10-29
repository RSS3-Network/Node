package activitypub

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
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
func NewOption(n network.Network, parameters *config.Parameters) (*Option, error) {
	var option Option

	zap.L().Info("parameters:", zap.Any("parameters", parameters))

	if parameters == nil {
		return &Option{
			TimestampStart: parameter.CurrentNetworkStartBlock[n].Timestamp,
		}, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	if option.TimestampStart == 0 {
		option.TimestampStart = parameter.CurrentNetworkStartBlock[n].Timestamp
	}

	zap.L().Info("option:", zap.Any("option", option))

	return &option, nil
}

// GetParams extracts the mastodon parameters
func GetParams(param *config.Parameters) ([]string, int, error) {
	// Extract relay URL list
	relayURLsRaw, ok := (*param)["relay_url_list"]
	if !ok {
		return nil, 0, fmt.Errorf("relay_url_list not found in parameters")
	}

	// Extract port, ensuring it's a string
	port, ok := (*param)["port"].(int)
	if !ok {
		return nil, 0, fmt.Errorf("port not found or is not a string in parameters")
	}

	// Parse relay URLs based on type
	switch urls := relayURLsRaw.(type) {
	case []string:
		return urls, port, nil
	case []interface{}:
		var relayURLs []string

		for _, u := range urls {
			strURL, ok := u.(string)
			if ok {
				relayURLs = append(relayURLs, strURL)
			}
		}

		if len(relayURLs) == 0 {
			return nil, 0, fmt.Errorf("no valid URLs found in relay_url_list")
		}

		return relayURLs, port, nil
	default:
		return nil, 0, fmt.Errorf("invalid relay_url_list type: %T", urls)
	}
}
