package atproto

import (
	"context"
	"fmt"
	"sync"

	cb "github.com/emirpasic/gods/queues/circularbuffer"
	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/engine/protocol/atproto"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/middleware"
	"github.com/rss3-network/node/provider/atproto/bluesky"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Component struct {
	config         *config.File
	counter        metric.Int64Counter
	databaseClient database.Client
	redisClient    rueidis.Client
	blueskyClient  *bluesky.Client
}

const Name = "atproto"

var (
	RecentRequests      *cb.Queue
	recentRequestsMutex sync.RWMutex
)

func (c *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(ctx context.Context, apiServer *echo.Echo, config *config.File, databaseClient database.Client, redisClient rueidis.Client) (*Component, error) {
	option, err := atproto.NewOption(config.Component.Atproto[0].Parameters)
	if err != nil {
		return nil, fmt.Errorf("create option failed: %w", err)
	}

	blueskyClient, err := bluesky.NewClient(ctx, option.Username, option.Password)
	if err != nil {
		return nil, fmt.Errorf("create bluesky client failed: %w", err)
	}

	c := &Component{
		config:         config,
		databaseClient: databaseClient,
		redisClient:    redisClient,
		blueskyClient:  blueskyClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	// Add middleware for bearer token authentication
	group.Use(middleware.BearerAuth(config.Discovery.Server.AccessToken))

	if err := c.InitMeter(); err != nil {
		panic(err)
	}

	return c, nil
}

func (c *Component) InitMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if c.counter, err = meter.Int64Counter(c.Name()); err != nil {
		return fmt.Errorf("failed to init meter for component %s: %w", c.Name(), err)
	}

	return nil
}

func (c *Component) CollectMetric(ctx context.Context, path, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", c.Name()),
		attribute.String("path", path),
		attribute.String("value", value),
	)

	c.counter.Add(ctx, int64(1), measurementOption)
}

func (c *Component) CollectTrace(ctx context.Context, path, value string) {
	spanStartOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(
			attribute.String("path", path),
			attribute.String("value", value),
		),
	}

	_, span := otel.Tracer("").Start(ctx, "atproto API Query", spanStartOptions...)
	defer span.End()
}

// GetRecentRequest returns the filtered recent requests.
func GetRecentRequest() []string {
	recentRequestsMutex.RLock()
	defer recentRequestsMutex.RUnlock()

	// Convert queue to slice
	values := RecentRequests.Values()

	// Filter out empty strings and convert to []string
	filteredRequests := lo.FilterMap(values, func(item interface{}, _ int) (string, bool) {
		str, ok := item.(string)
		return str, ok && str != ""
	})

	return filteredRequests
}
