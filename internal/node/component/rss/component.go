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
	httpClient      *http.Client
	rsshub          *config.Module
	meterRSSCounter metric.Int64Counter
}

const Name = "rss"

func (h *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, config []*config.Module) component.Component {
	component := &Component{
		httpClient: http.DefaultClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	group.GET("/*", component.GetRSSHubHandler)

	if err := component.InitMeter(); err != nil {
		return nil
	}

	for _, conf := range config {
		if conf.Network == network.RSS {
			component.rsshub = conf
			break
		}
	}

	return component
}

func (h *Component) InitMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.meterRSSCounter, err = meter.Int64Counter(h.Name()); err != nil {
		return fmt.Errorf("failed to init meter for component %s: %w", h.Name(), err)
	}

	return nil
}

func (h *Component) CollectMetric(ctx context.Context, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", h.Name()),
		attribute.String("value", value),
	)

	h.meterRSSCounter.Add(ctx, int64(1), measurementOption)
}
