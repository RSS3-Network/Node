package rss

import (
	"context"
	"net/http"

	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
)

type RSS struct {
	httpClient *http.Client
	rsshub     *engine.Config
}

// NewRSS creates a new rss explorer
func NewRSS(_ context.Context, config *config.File) *RSS {
	rss := &RSS{
		httpClient: http.DefaultClient,
	}

	for _, conf := range config.Node.RSS {
		if conf.Chain == filter.ChainRSSRSSHub {
			rss.rsshub = conf
		}
	}

	return rss
}
