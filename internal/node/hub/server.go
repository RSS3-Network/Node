package hub

import (
	"context"
	"net"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
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
	instance.httpServer.Validator = &Validator{
		validate: validator.New(),
	}

	instance.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// register router
	instance.httpServer.GET("/rss/*", instance.hub.RSS.GetRSSHubHandler)
	instance.httpServer.GET("/decentralized/tx/:id", instance.hub.Decentralized.GetActivity)
	instance.httpServer.GET("/decentralized/:account", instance.hub.Decentralized.GetAccountActivities)
	instance.httpServer.GET("/decentralized/count", instance.hub.Decentralized.GetActivitiesCount)

	return &instance, nil
}
