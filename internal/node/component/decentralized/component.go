package decentralized

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
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/middleware"
	"github.com/rss3-network/node/provider/ethereum/etherface"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Component struct {
	config          *config.File
	counter         metric.Int64Counter
	databaseClient  database.Client
	etherfaceClient etherface.Client
	redisClient     rueidis.Client
}

const Name = "decentralized"

const MaxRecentRequests = 10

var (
	RecentRequests      *cb.Queue
	recentRequestsMutex sync.RWMutex
)

func (c *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config *config.File, databaseClient database.Client, redisClient rueidis.Client) component.Component {
	RecentRequests = cb.New(MaxRecentRequests)

	c := &Component{
		config:         config,
		databaseClient: databaseClient,
		redisClient:    redisClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	// Add middleware for bearer token authentication
	group.Use(middleware.BearerAuth(config.Discovery.Server.AccessToken))

	group.GET("/tx/:id", c.GetActivity)
	group.GET("/:account", c.GetAccountActivities)
	group.GET("/network/:network", c.GetNetworkActivities)
	group.GET("/platform/:platform", c.GetPlatformActivities)

	group.POST("/accounts", c.BatchGetAccountsActivities)

	if err := c.InitMeter(); err != nil {
		panic(err)
	}

	// Initialize etherface client, an optional dependency
	etherfaceClient, err := etherface.NewEtherfaceClient()
	if err != nil {
		zap.L().Warn("failed to initialize etherface client", zap.Any("error", err))
	} else {
		c.etherfaceClient = etherfaceClient
	}

	return c
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

	_, span := otel.Tracer("").Start(ctx, "Decentralized API Query", spanStartOptions...)
	defer span.End()
}

func addRecentRequest(path string) {
	recentRequestsMutex.Lock()
	defer recentRequestsMutex.Unlock()

	RecentRequests.Enqueue(path)
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
