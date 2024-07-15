package mastodon

import (
	"context"
	"fmt"

	"github.com/rss3-network/node/internal/engine"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	httpClient httpx.Client
}

func (w *worker) Name() string {
	return federated.Mastodon.String()
}

func (w *worker) Platform() string {
	return federated.PlatformMastodon.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Mastodon,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Social,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{}
}

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return nil
}

func (w *worker) Transform(_ context.Context, _ engine.Task) (*activityx.Activity, error) {
	fmt.Println("[mastodon/worker.go] reached Transform().")
	return nil, nil
}

func NewWorker() (engine.Worker, error) {
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &worker{
		httpClient: httpClient,
	}, nil
}
