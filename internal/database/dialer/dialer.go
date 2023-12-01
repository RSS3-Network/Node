package dialer

import (
	"context"
	"fmt"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/database/dialer/cockroachdb"
)

func Dial(ctx context.Context, config *database.Config) (database.Client, error) {
	switch config.Driver {
	case database.DriverCockroachDB:
		return cockroachdb.Dial(ctx, config.URI, config.Partition)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", config.Driver)
	}
}
