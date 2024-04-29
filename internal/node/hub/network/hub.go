package network

import (
	"context"
	"fmt"

	"github.com/rss3-network/node/internal/constant"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Hub struct {
	meterRSSCounter metric.Int64Counter
}

// NewHub creates a new rss hub
func NewHub(_ context.Context) *Hub {
	hub := &Hub{}

	if err := hub.initializeMeter(); err != nil {
		return nil
	}

	return hub
}

func (h *Hub) initializeMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.meterRSSCounter, err = meter.Int64Counter("rss3_node_network"); err != nil {
		return fmt.Errorf("create meter of network request: %w", err)
	}

	return nil
}

func (h *Hub) countRequest(ctx context.Context, network string) {
	meterRSSCounterAttributes := metric.WithAttributes(
		attribute.String("hub", "network"),
		attribute.String("network", network),
	)

	h.meterRSSCounter.Add(ctx, int64(1), meterRSSCounterAttributes)
}
