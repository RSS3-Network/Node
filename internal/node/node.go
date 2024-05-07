package node

import (
	"context"
	"net"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/decentralized"
	"github.com/rss3-network/node/internal/node/component/rss"
)

const (
	// FIXME: update startup command to use `core` instead of `hub`
	CoreService = "hub"
	// FIXME: do not use `indexer`, use `worker` instead
	Worker      = "indexer"
	Broadcaster = "broadcaster"
	DefaultHost = "0.0.0.0"
	DefaultPort = "80"
)

// Core is logically formed by an API server and a list of components
type Core struct {
	apiServer  *echo.Echo
	components []*component.Component
}

func (s *Core) Run(_ context.Context) error {
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	return s.apiServer.Start(address)
}

// NewCoreService initializes the core services required by the Core
func NewCoreService(ctx context.Context, config *config.File, databaseClient database.Client) *Core {
	apiServer := echo.New()

	node := Core{
		apiServer:  apiServer,
		components: []*component.Component{},
	}

	apiServer.HideBanner = true
	apiServer.HidePort = true
	apiServer.Validator = &component.Validator{
		Validator: validator.New(),
	}

	apiServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	if len(config.Component.RSS) > 0 {
		rssComponent := rss.NewComponent(ctx, apiServer, config.Component.RSS)
		node.components = append(node.components, &rssComponent)
	}

	if len(config.Component.Decentralized) > 0 {
		decentralizedComponent := decentralized.NewComponent(ctx, apiServer, databaseClient)
		node.components = append(node.components, &decentralizedComponent)
	}

	return &node
}
