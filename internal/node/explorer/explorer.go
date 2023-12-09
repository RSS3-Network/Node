package explorer

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/node/explorer/rss"
)

type Explorer struct {
	RSS *rss.RSS
}

func NewExplorer(ctx context.Context, config *config.File, _ database.Client) *Explorer {
	return &Explorer{
		RSS: rss.NewRSS(ctx, config),
	}
}
