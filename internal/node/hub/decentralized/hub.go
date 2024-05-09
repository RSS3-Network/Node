package decentralized

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/provider/ethereum/etherface"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Hub struct {
	config                    *config.File
	databaseClient            database.Client
	redisClient               rueidis.Client
	meterDecentralizedCounter metric.Int64Counter
	etherfaceClient           etherface.Client
}

// NewHub creates a new decentralized hub
func NewHub(_ context.Context, databaseClient database.Client, config *config.File, redisClient rueidis.Client) *Hub {
	hub := Hub{
		config:         config,
		databaseClient: databaseClient,
		redisClient:    redisClient,
	}

	if err := hub.initializeMeter(); err != nil {
		return nil
	}

	// Initialize etherface client
	etherfaceClient, err := etherface.NewEtherfaceClient()
	if err != nil {
		log.Error("failed to initialize etherface client", "error", err)
	}

	hub.etherfaceClient = etherfaceClient

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
