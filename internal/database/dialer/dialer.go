package dialer

import (
	"context"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/dialer/postgres"
)

func Dial(ctx context.Context, config *config.Database) (database.Client, error) {
	return postgres.Dial(ctx, config.URI, true)
}
