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
	// Hub FIXME: update startup command to use `components` instead of `hub`
	Hub = "hub"
	// Indexer FIXME: do not use `indexer`, use `worker` instead
	Indexer     = "indexer"
	Broadcaster = "broadcaster"
	DefaultHost = "0.0.0.0"
	DefaultPort = "80"
)

// Node is logically formed by an API server and a list of components
type Node struct {
	apiServer  *echo.Echo
	components []*component.Component
}

func (s *Node) Run(_ context.Context) error {
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	return s.apiServer.Start(address)
}

func NewNode(ctx context.Context, config *config.File, databaseClient database.Client) *Node {
	apiServer := echo.New()

	node := Node{
		apiServer:  apiServer,
		components: []*component.Component{},
	}

	apiServer.HideBanner = true
	apiServer.HidePort = true
	apiServer.Validator = &component.Validator{
		Validator: validator.New(),
	}

	apiServer.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	if len(config.Node.RSS) > 0 {
		rssComponnet := rss.NewComponent(ctx, apiServer, config.Node.RSS)
		node.components = append(node.components, &rssComponnet)
	}

	if len(config.Node.Decentralized) > 0 {
		decentralizedComponent := decentralized.NewComponent(ctx, apiServer, databaseClient)
		node.components = append(node.components, &decentralizedComponent)
	}

	return &node
}
