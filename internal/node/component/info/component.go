package info

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/constant"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/node/component"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/vsl"
	"github.com/rss3-network/node/v2/provider/httpx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Component struct {
	config              *config.File
	counter             metric.Int64Counter
	databaseClient      database.Client
	redisClient         rueidis.Client
	networkParamsCaller *vsl.NetworkParamsCaller
	httpClient          httpx.Client
}

const Name = "info"

func (c *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, _ *echo.Echo, config *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller) *Component {
	httpxClient, err := httpx.NewHTTPClient()
	if err != nil {
		return nil
	}

	c := &Component{
		config:              config,
		databaseClient:      databaseClient,
		redisClient:         redisClient,
		networkParamsCaller: networkParamsCaller,
		httpClient:          httpxClient,
	}

	if err := c.InitMeter(); err != nil {
		panic(err)
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

	_, span := otel.Tracer("").Start(ctx, "Info API Query", spanStartOptions...)
	defer span.End()
}
