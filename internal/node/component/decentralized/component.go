package decentralized

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/internal/constant"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/node/component"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type Component struct {
	databaseClient            database.Client
	meterDecentralizedCounter metric.Int64Counter
}

const Name = "decentralized"

func (h *Component) Name() string {
	return Name
}

var _ component.Component = (*Component)(nil)

func NewComponent(_ context.Context, apiServer *echo.Echo, databaseClient database.Client) component.Component {
	component := &Component{
		databaseClient: databaseClient,
	}

	group := apiServer.Group(fmt.Sprintf("/%s", Name))

	group.GET("/tx/:id", component.GetActivity)
	group.GET("/:account", component.GetAccountActivities)
	group.GET("/count", component.GetActivitiesCount)

	if err := component.InitMeter(); err != nil {
		return nil
	}

	return component
}

func (h *Component) InitMeter() (err error) {
	meter := otel.GetMeterProvider().Meter(constant.Name)

	if h.meterDecentralizedCounter, err = meter.Int64Counter(h.Name()); err != nil {
		return fmt.Errorf("failed to init meter for component %s: %w", h.Name(), err)
	}

	return nil
}

func (h *Component) CollectMetric(ctx context.Context, value string) {
	measurementOption := metric.WithAttributes(
		attribute.String("component", h.Name()),
		attribute.String("path", value),
	)

	h.meterDecentralizedCounter.Add(ctx, int64(1), measurementOption)
}
