package provider

import (
	"context"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/stream"
	"github.com/rss3-network/node/v2/internal/stream/provider/kafka"
)

const kafkaTopic = "rss3.node.activities"

func New(ctx context.Context, config *config.Stream) (stream.Client, error) {
	return kafka.New(ctx, config.URI, kafkaTopic)
}
