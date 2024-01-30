package meter

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rss3-network/node/config"
)

var (
	ErrorUnsupportedDriver = errors.New("unsupported driver")
)

type Driver string

const (
	DriverPrometheus Driver = "prometheus"
)

type Server interface {
	Run(config.OpenTelemetryMetricsConfig) error
}

var _ Server = (*server)(nil)

type server struct {
	httpServer *echo.Echo
}

func (s *server) Run(config config.OpenTelemetryMetricsConfig) error {
	return s.httpServer.Start(config.Endpoint)
}

func New(driver Driver) (Server, error) {
	if driver != DriverPrometheus {
		return nil, fmt.Errorf("%w: %s", ErrorUnsupportedDriver, driver)
	}

	instance := server{
		httpServer: echo.New(),
	}

	instance.httpServer.HideBanner = true
	instance.httpServer.HidePort = true

	instance.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	instance.httpServer.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	instance.httpServer.GET("/healthcheck", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})

	return &instance, nil
}
