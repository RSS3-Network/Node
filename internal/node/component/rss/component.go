package rss

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/middleware"
	"github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Component struct {
	config     *config.File
	httpClient *http.Client
	rsshub     *configx
	counter    metric.Int64Counter
}

type configx struct {
	id        string
	network   network.Network
	worker    worker.Worker
	endpoint  string
	accessKey string
}

const Name = "rss"

func (h *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config *config.File) component.Component {
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

	if config.Component.RSS != nil && config.Component.RSS.Network == network.RSS {
		c.rsshub = &configx{
			id:       config.Component.RSS.ID,
			network:  config.Component.RSS.Network,
			worker:   config.Component.RSS.Worker,
			endpoint: config.Component.RSS.Endpoint.URL,
		}

		if err := c.setAccessKey(config.Component.RSS); err != nil {
			c.rsshub = nil
		}
	}

	if c.rsshub == nil {
		panic("Missing configuration for Component RSS")
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

// setAccessKey set the access code according to the RSSHub authentication specification.
func (h *Component) setAccessKey(config *config.Module) error {
	if config.Parameters == nil {
		return nil
	}

	option, err := NewOption(config.Parameters)
	if err != nil {
		return fmt.Errorf("parse config parmeters: %w", err)
	}

	if option.Authentication.AccessKey != "" {
		h.rsshub.accessKey = option.Authentication.AccessKey
	}

	return nil
}
