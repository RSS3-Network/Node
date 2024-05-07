package rss

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/protocol-go/schema/network"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Component struct {
	httpClient *http.Client
	rsshub     *configx
	counter    metric.Int64Counter
}

type configx struct {
	endpoint  string
	accessKey string
}

const Name = "rss"

func (h *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config []*config.Module) component.Component {
	c := &Component{
		httpClient: http.DefaultClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	group.GET("/*", c.GetRSSHubHandler)

	if err := c.InitMeter(); err != nil {
		return nil
	}

	for _, conf := range config {
		if conf.Network == network.RSS {
			c.rsshub = &configx{
				endpoint: conf.Endpoint,
			}

			if err := c.setAccessKey(context.Background(), conf); err != nil {
				c.rsshub = nil
			}

			break
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

func (h *Component) CollectMetric(ctx context.Context, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", h.Name()),
		attribute.String("value", value),
	)

	h.counter.Add(ctx, int64(1), measurementOption)
}

// setAccessKey set the access code according to the RSSHub authentication specification.
func (h *Component) setAccessKey(_ context.Context, config *config.Module) error {
	if config.Parameters == nil {
		return nil
	}

	option, err := NewOption(config.Parameters)
	if err != nil {
		return fmt.Errorf("parse parmeters: %w", err)
	}

	if option.Authentication.AccessKey != "" {
		h.rsshub.accessKey = option.Authentication.AccessKey
	}

	return nil
}
