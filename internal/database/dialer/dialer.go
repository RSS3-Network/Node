package dialer

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer/cockroach"
)

func Dial(ctx context.Context, config database.Config) (database.Client, error) {
	switch config.Driver {
	case database.DriverCockroach:
		return cockroach.Dial(ctx, config.URI, config.Mode)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}
