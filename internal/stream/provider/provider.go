package provider

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/stream"
	"github.com/naturalselectionlabs/rss3-node/internal/stream/provider/kafka"
)

func New(ctx context.Context, config *config.Stream) (stream.Client, error) {
	switch config.Driver {
	case stream.DriverKafka:
		return kafka.New(ctx, config.URI, config.Topic)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}
