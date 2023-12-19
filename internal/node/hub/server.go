package hub

import (
	"context"
	"net"
	"net/http"

	"github.com/NaturalSelectionLabs/goapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "80"
)

type Server struct {
	httpServer *echo.Echo
	hub        *Hub
}

func (s *Server) Run(_ context.Context) error {
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	return s.httpServer.Start(address)
}

func NewServer(ctx context.Context, config *config.File, databaseClient database.Client) (*Server, error) {
	instance := Server{
		httpServer: echo.New(),
		hub:        NewHub(ctx, config, databaseClient),
	}

	instance.httpServer.HideBanner = true
	instance.httpServer.HidePort = true

	instance.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// register router
	group := goapi.NewRouter().Group("")
	{
		group.GET("/healthcheck", HealthCheck)
		group.GET("/rss/rsshub/*", instance.hub.RSS.GetRSSHubHandler)
		group.GET("/activities/{id}", instance.hub.Decentralized.GetActivity)
		group.GET("/accounts/{account}/activities", instance.hub.Decentralized.GetAccountActivities)
	}

	instance.httpServer.Use(echo.WrapMiddleware(group.Handler))

	instance.httpServer.Routers()

	return &instance, nil
}

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
