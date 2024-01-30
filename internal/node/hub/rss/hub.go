package rss

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/protocol-go/schema/filter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Hub struct {
	httpClient      *http.Client
	rsshub          *config.Module
	meterRSSCounter metric.Int64Counter
}

// NewHub creates a new rss hub
func NewHub(_ context.Context, config *config.File) *Hub {
	hub := &Hub{
		httpClient: http.DefaultClient,
	}

	if err := hub.initializeMeter(); err != nil {
		return nil
	}

	for _, conf := range config.Node.RSS {
		if conf.Network == filter.NetworkRSS {
			hub.rsshub = conf
		}
	}

	return hub
}

func (h *Hub) initializeMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.meterRSSCounter, err = meter.Int64Counter("rss3_node_rss"); err != nil {
		return fmt.Errorf("create meter of rss request: %w", err)
	}

	return nil
}

func (h *Hub) countRequest(ctx context.Context, path string) {
	meterRSSCounterAttributes := metric.WithAttributes(
		attribute.String("hub", "rss"),
		attribute.String("path", path),
	)

	h.meterRSSCounter.Add(ctx, int64(1), meterRSSCounterAttributes)
}
