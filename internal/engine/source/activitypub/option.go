package activitypub

import (
	"fmt"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/protocol-go/schema/network"
)

// Option represents the configuration options for the ActivityPub client.
type Option struct {
	KafkaTopic     string `json:"kafka_topic"`
	TimestampStart int64  `json:"timestamp_start" mapstructure:"timestamp_start"`
}

// NewOption creates a new Option instance from the provided parameters.
func NewOption(n network.Network, parameters *config.Parameters) (*Option, error) {
	var option Option

	fmt.Println("parameters is:", parameters)

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

	return &option, nil
}
