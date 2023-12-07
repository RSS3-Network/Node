package rss

import (
	"context"
	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"net/http"
)

type Handler struct {
	httpClient *http.Client
	rsshub     *engine.Config
}

// NewHandler creates a new rss Handler
func NewHandler(_ context.Context, config *config.File) *Handler {
	handler := &Handler{
		httpClient: http.DefaultClient,
	}

	for _, rss := range config.Node.RSS {
		if rss.Chain == filter.ChainRSSRSSHub {
			handler.rsshub = rss
		}
	}

	return handler
}
