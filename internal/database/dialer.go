package database

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database/internal"
)

func Dial(_ context.Context, config Config) (client *internal.Client, err error) {
	switch config.Driver {
	case DriverCockroach:
		client, err = internal.Open(string(DriverPostgres), config.URI)
	default:
		client, err = internal.Open(string(config.Driver), config.URI)
	}

	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	return client, nil
}
