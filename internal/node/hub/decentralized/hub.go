package decentralized

import (
	"context"

	"github.com/naturalselectionlabs/rss3-node/internal/database"
)

type Hub struct {
	databaseClient database.Client
}

// NewHub creates a new decentralized hub
func NewHub(_ context.Context, databaseClient database.Client) *Hub {
	return &Hub{
		databaseClient: databaseClient,
	}
}
