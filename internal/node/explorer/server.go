package explorer

import (
	"context"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-node/internal/config"
	"github.com/naturalselectionlabs/rss3-node/internal/database"
	"github.com/naturalselectionlabs/rss3-node/internal/node/explorer/handler"
)

type Server struct {
	httpServer *echo.Echo
	handler    *handler.Handler
}

func (s *Server) Run(_ context.Context) error {
	address := net.JoinHostPort("127.0.0.1", "8000")

	return s.httpServer.Start(address)
}

func NewServer(ctx context.Context, config *config.File, databaseClient database.Client) (*Server, error) {
	instance := Server{
		httpServer: echo.New(),
		handler:    handler.NewHandler(ctx, config, databaseClient),
	}

	instance.httpServer.HideBanner = true
	instance.httpServer.HidePort = true
	instance.httpServer.GET("/rss/rsshub/*", instance.handler.RSSHandler.GetRSSHub)

	return &instance, nil
}
