package hub

import (
	"context"
	"net"

	"github.com/labstack/echo/v4"
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

	// register router
	instance.httpServer.GET("/rss/rsshub/*", instance.hub.RSS.GetRSSHubHandler)
	instance.httpServer.GET("/activities/:id", instance.hub.Decentralized.GetActivity)
	instance.httpServer.GET("/accounts/:account/activities", instance.hub.Decentralized.GetAccountActivities)

	return &instance, nil
}
