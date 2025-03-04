package agentdata

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	cb "github.com/emirpasic/gods/queues/circularbuffer"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/middleware"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Component struct {
	config     *config.File
	httpClient *http.Client
	counter    metric.Int64Counter
}

const Name = "agentdata"

const MaxRecentRequests = 10

var (
	RecentRequests      *cb.Queue
	recentRequestsMutex sync.RWMutex
)

func (h *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config *config.File) *Component {
	RecentRequests = cb.New(MaxRecentRequests)

	c := &Component{
		config:     config,
		httpClient: http.DefaultClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	// Add middleware for bearer token authentication
	group.Use(middleware.BearerAuth(config.Discovery.Server.AccessToken))

	group.GET("/*", c.Handler)

	if err := c.InitMeter(); err != nil {
		panic(err)
	}

	return c
}

func (h *Component) InitMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.counter, err = meter.Int64Counter(h.Name()); err != nil {
		return fmt.Errorf("failed to init meter for component %s: %w", h.Name(), err)
	}

	return nil
}

func (h *Component) CollectMetric(ctx context.Context, path, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", h.Name()),
		attribute.String("path", path),
		attribute.String("value", value),
	)

	h.counter.Add(ctx, int64(1), measurementOption)
}

func (h *Component) CollectTrace(ctx context.Context, path, value string) {
	spanStartOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindServer),
		trace.WithAttributes(
			attribute.String("path", path),
			attribute.String("value", value),
		),
	}

	_, span := otel.Tracer("").Start(ctx, "RSS API Query", spanStartOptions...)
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
