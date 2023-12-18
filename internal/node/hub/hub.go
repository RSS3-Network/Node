package hub

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/node/hub/decentralized"
	"github.com/naturalselectionlabs/rss3-node/internal/node/hub/rss"
)

type Hub struct {
	RSS           *rss.Hub
	Decentralized *decentralized.Hub
}

func NewHub(ctx context.Context, config *config.File, database database.Client) *Hub {
	return &Hub{
		RSS:           rss.NewHub(ctx, config),
		Decentralized: decentralized.NewHub(ctx, database),
	}
}
