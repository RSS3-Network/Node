package broadcaster

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/robfig/cron/v3"
	"github.com/rss3-network/node/config"
	"go.uber.org/zap"
)

const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "80"
)

type Broadcaster struct {
	config     *config.File
	cron       *cron.Cron
	httpClient *http.Client
	httpServer *echo.Echo
}

func (b *Broadcaster) Run(ctx context.Context) error {
	// run api server
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	go func() {
		if err := b.httpServer.Start(address); err != nil {
			zap.L().Error("failed to run http server", zap.Error(err))
		}
	}()

	// run register cron job
	if err := b.Register(ctx); err != nil {
		return fmt.Errorf("register: %w", err)
	}

	_, err := b.cron.AddFunc("@every 1m", func() {
		if err := b.Heartbeat(ctx); err != nil {
			return
		}
	})
	if err != nil {
		return fmt.Errorf("add heartbeat cron job: %w", err)
	}

	b.cron.Start()

	stopchan := make(chan os.Signal, 1)

	signal.Notify(stopchan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	<-stopchan

	return nil
}

func NewBroadcaster(_ context.Context, config *config.File) (*Broadcaster, error) {
	instance := &Broadcaster{
		config:     config,
		cron:       cron.New(),
		httpClient: http.DefaultClient,
		httpServer: echo.New(),
	}

	instance.httpServer.HideBanner = true
	instance.httpServer.HidePort = true

	instance.httpServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	// register router
	instance.httpServer.GET("/", instance.GetNodeInfo)

	return instance, nil
}
