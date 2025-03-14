package dialer

import (
	"context"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/database/dialer/postgres"
)

func Dial(ctx context.Context, config *config.Database) (database.Client, error) {
	return postgres.Dial(ctx, config.URI, true)
}
