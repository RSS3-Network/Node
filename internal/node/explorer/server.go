package explorer

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
	explorer   *Explorer
}

func (s *Server) Run(_ context.Context) error {
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	return s.httpServer.Start(address)
}

func NewServer(ctx context.Context, config *config.File, databaseClient database.Client) (*Server, error) {
	instance := Server{
		httpServer: echo.New(),
		explorer:   NewExplorer(ctx, config, databaseClient),
	}

	instance.httpServer.HideBanner = true
	instance.httpServer.HidePort = true
	instance.httpServer.GET("/rss/rsshub/*", instance.explorer.RSS.GetRSSHubHandler)

	return &instance, nil
}
