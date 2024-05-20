package network

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Component struct {
	config  *config.File
	counter metric.Int64Counter
}

const Name = "networks"

func (c *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config *config.File) component.Component {
	c := &Component{
		config: config,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	group.GET("", c.GetNetworksHandler)
	group.GET("/get_endpoint_config", c.GetEndpointConfig)
	group.GET("/:network/list_workers", c.GetWorkersByNetwork)
	group.GET("/:network/workers/:worker", c.GetWorkerConfig)

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

func (c *Component) CollectMetric(ctx context.Context, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", c.Name()),
		attribute.String("path", value),
	)

	c.counter.Add(ctx, int64(1), measurementOption)
}
