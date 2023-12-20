package rss

import (
	"context"
	"net/http"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type Hub struct {
	httpClient *http.Client
	rsshub     *engine.Config
}

// NewHub creates a new rss hub
func NewHub(_ context.Context, config *config.File) *Hub {
	hub := &Hub{
		httpClient: http.DefaultClient,
	}

	for _, conf := range config.Node.RSS {
		if conf.Network == filter.NetworkRSS {
			hub.rsshub = conf
		}
	}

	return hub
}
