package rsshub

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/node/v2/schema/worker/rss"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
)

type Worker struct {
	config     *config.File
	httpClient *http.Client
	rsshub     *configx
}

type configx struct {
	id       string
	network  network.Network
	endpoint string
}

func (w *Worker) Name() string {
	return "RSSHub"
}

func (w *Worker) Platform() string {
	return rss.PlatformRSSHub.String()
}

func (w *Worker) Network() []network.Network {
	return []network.Network{w.rsshub.network}
}

func (w *Worker) Tags() []tag.Tag {
	return []tag.Tag{tag.Social}
}

func (w *Worker) Types() []schema.Type {
	return []schema.Type{typex.SocialPost}
}

func (w *Worker) Filter() engine.DataSourceFilter {
	return nil
}

func (w *Worker) Transform(_ context.Context, _ engine.Task) (*activityx.Activity, error) {
	// Not implemented
	return nil, fmt.Errorf("Transform method not implemented for RSSHub worker")
}

func NewWorker(config *config.File) (engine.Worker, error) {
	w := &Worker{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	if config.Component.RSS != nil && config.Component.RSS.Network == network.RSSHub {
		w.rsshub = &configx{
			id:       config.Component.RSS.ID,
			network:  config.Component.RSS.Network,
			endpoint: config.Component.RSS.Endpoint.URL,
		}
	}

	if w.rsshub == nil {
		return nil, fmt.Errorf("missing configuration for Component RSS")
	}

	return w, nil
}
