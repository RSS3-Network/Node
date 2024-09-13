package node

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/rueidis"
	"github.com/robfig/cron/v3"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/config/parameter"
	"github.com/rss3-network/node/docs"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/node/component"
	"github.com/rss3-network/node/internal/node/component/decentralized"
	"github.com/rss3-network/node/internal/node/component/info"
	"github.com/rss3-network/node/internal/node/component/rss"
	"github.com/rss3-network/node/internal/node/middlewarex"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"go.uber.org/zap"
)

// default host and port for the API server
const (
	DefaultHost = "0.0.0.0"
	DefaultPort = "80"
)

// Core is logically formed by an API server and a list of components
type Core struct {
	apiServer           *echo.Echo
	components          []*component.Component
	cron                *cron.Cron
	redisClient         rueidis.Client
	settlementCaller    *vsl.SettlementCaller
	networkParamsCaller *vsl.NetworkParamsCaller
}

func (s *Core) Run(_ context.Context) error {
	address := net.JoinHostPort(DefaultHost, DefaultPort)

	return s.apiServer.Start(address)
}

// NewCoreService initializes the core services required by the Core
func NewCoreService(ctx context.Context, config *config.File, databaseClient database.Client, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) *Core {
	apiServer := echo.New()

	node := Core{
		apiServer:           apiServer,
		components:          []*component.Component{},
		cron:                cron.New(),
		settlementCaller:    settlementCaller,
		networkParamsCaller: networkParamsCaller,
		redisClient:         redisClient,
	}

	apiServer.HideBanner = true
	apiServer.HidePort = true
	apiServer.Validator = &component.Validator{
		Validator: validator.New(),
	}

	apiServer.Use(
		middleware.CORSWithConfig(middleware.DefaultCORSConfig),
		middlewarex.DecodePathParamsMiddleware,
	)

	infoComponent := info.NewComponent(ctx, apiServer, config, databaseClient, redisClient, networkParamsCaller)
	node.components = append(node.components, &infoComponent)

	if config.Component.RSS != nil {
		rssComponent := rss.NewComponent(ctx, apiServer, config)
		node.components = append(node.components, &rssComponent)
	}

	if len(config.Component.Decentralized) > 0 {
		decentralizedComponent := decentralized.NewComponent(ctx, apiServer, config, databaseClient, redisClient)
		node.components = append(node.components, &decentralizedComponent)
	}

	// Generate openapi.json
	var endpoint string
	if config.Discovery != nil && config.Discovery.Server != nil {
		endpoint = config.Discovery.Server.Endpoint
	}

	content, err := docs.Generate(endpoint)
	if err != nil {
		zap.L().Error("docs generation error", zap.Error(err))
	}

	apiServer.GET("/openapi.json", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "application/json", content)
	})

	return &node
}

// CheckParams checks the network parameters and settlement tasks
func CheckParams(ctx context.Context, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	checkParamsCron := cron.New()

	_, err := checkParamsCron.AddFunc("@every 5m", func() {
		localEpoch, err := parameter.GetCurrentEpoch(ctx, redisClient)
		if err != nil {
			zap.L().Error("get current epoch", zap.Error(err))
			return
		}

		remoteEpoch, err := parameter.GetCurrentEpochFromVSL(settlementCaller)
		if err != nil {
			zap.L().Error("get current epoch", zap.Error(err))
			return
		}

		if remoteEpoch > localEpoch {
			err = parameter.UpdateCurrentEpoch(ctx, redisClient, remoteEpoch)
			if err != nil {
				zap.L().Error("update current epoch", zap.Error(err))
				return
			}
		}

		if err := parameter.CheckParamsTask(ctx, redisClient, networkParamsCaller); err != nil {
			zap.L().Error("check params tasks", zap.Error(err))
		}
	})
	if err != nil {
		return fmt.Errorf("add check param cron job: %w", err)
	}

	checkParamsCron.Start()

	<-ctx.Done()
	checkParamsCron.Stop()

	return ctx.Err()
}
