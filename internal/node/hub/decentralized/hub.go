package decentralized

import (
	"context"
	"fmt"

	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Hub struct {
	databaseClient            database.Client
	meterDecentralizedCounter metric.Int64Counter
}

// NewHub creates a new decentralized hub
func NewHub(_ context.Context, databaseClient database.Client) *Hub {
	hub := Hub{
		databaseClient: databaseClient,
	}

	if err := hub.initializeMeter(); err != nil {
		return nil
	}

	return &hub
}

func (h *Hub) initializeMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.meterDecentralizedCounter, err = meter.Int64Counter("rss3_node_decentralized"); err != nil {
		return fmt.Errorf("create meter of decentralized request: %w", err)
	}

	return nil
}

func (h *Hub) countAccount(ctx context.Context, account string) {
	meterDecentralizedCounterAttributes := metric.WithAttributes(
		attribute.String("hub", "decentralized"),
		attribute.String("account", account),
	)

	h.meterDecentralizedCounter.Add(ctx, int64(1), meterDecentralizedCounterAttributes)
}
