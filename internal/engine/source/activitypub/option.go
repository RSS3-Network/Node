package activitypub

import (
	"github.com/rss3-network/node/config"
)

type Option struct {
	KafkaTopic string `json:"mastodon_kafka_topic"`
}

func NewOption(parameters *config.Parameters) (*Option, error) {
	var option Option

	if parameters == nil {
		return &option, nil
	}

	if err := parameters.Decode(&option); err != nil {
		return nil, err
	}

	return &option, nil
}
