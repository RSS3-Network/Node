package agentdata

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

type Worker struct {
	config     *config.File
	httpClient *http.Client
}

func (w *Worker) Name() string {
	return "AgentData"
}

func (w *Worker) Platform() string {
	return ""
}

func (w *Worker) Network() []network.Network {
	return nil
}

func (w *Worker) Tags() []tag.Tag {
	return nil
}

func (w *Worker) Types() []schema.Type {
	return nil
}

func (w *Worker) Filter() engine.DataSourceFilter {
	return nil
}

func (w *Worker) Transform(_ context.Context, _ engine.Task) (*activityx.Activity, error) {
	return nil, nil
}

func NewWorker(config *config.File) (engine.Worker, error) {
	w := &Worker{
		config: config,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	if config.Component.AI == nil {
		return nil, fmt.Errorf("missing configuration for Component Agentdata")
	}

	return w, nil
}
