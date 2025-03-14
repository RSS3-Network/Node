package node

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/rueidis"
	"github.com/robfig/cron/v3"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/config/parameter"
	"github.com/rss3-network/node/v2/docs"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/node/component"
	"github.com/rss3-network/node/v2/internal/node/component/aggregator"
	"github.com/rss3-network/node/v2/internal/node/component/ai"
	"github.com/rss3-network/node/v2/internal/node/component/decentralized"
	"github.com/rss3-network/node/v2/internal/node/component/federated"
	"github.com/rss3-network/node/v2/internal/node/component/info"
	"github.com/rss3-network/node/v2/internal/node/component/rss"
	"github.com/rss3-network/node/v2/internal/node/middlewarex"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/vsl"
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
		middlewarex.HeadToGetMiddleware,
	)

	aggComp := aggregator.Component{}

	infoComponent := info.NewComponent(ctx, apiServer, config, databaseClient, redisClient, networkParamsCaller)
	{
		var comp component.Component = infoComponent
		node.components = append(node.components, &comp)
		aggComp.Info = infoComponent
	}

	if config.Component.RSS != nil {
		rssComponent := rss.NewComponent(ctx, apiServer, config)
		{
			var comp component.Component = rssComponent
			node.components = append(node.components, &comp)
			aggComp.RSS = rssComponent
		}
	}

	if config.Component.AI != nil {
		aiComponent := ai.NewComponent(ctx, apiServer, config)
		{
			var comp component.Component = aiComponent
			node.components = append(node.components, &comp)
			aggComp.AI = aiComponent
		}
	}

	if len(config.Component.Decentralized) > 0 {
		decentralizedComponent := decentralized.NewComponent(ctx, apiServer, config, databaseClient, redisClient)
		{
			var comp component.Component = decentralizedComponent
			node.components = append(node.components, &comp)
			aggComp.Decentralized = decentralizedComponent
		}
	}

	if len(config.Component.Federated) > 0 {
		federatedComponent := federated.NewComponent(ctx, apiServer, config, databaseClient, redisClient)
		{
			var comp component.Component = federatedComponent
			node.components = append(node.components, &comp)
			aggComp.Federated = federatedComponent
		}
	}

	docs.RegisterHandlers(apiServer, aggComp)

	// Generate openapi.json
	apiServer.GET("/openapi.json", func(c echo.Context) error {
		swagger, err := docs.GetSwagger()
		swagger.Servers = append(swagger.Servers, &openapi3.Server{
			URL: config.Discovery.Server.Endpoint,
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, swagger)
	})

	zap.L().Info("Core service initialization completed")

	return &node
}

// CheckParams checks the network parameters and settlement tasks
func CheckParams(ctx context.Context, redisClient rueidis.Client, networkParamsCaller *vsl.NetworkParamsCaller, settlementCaller *vsl.SettlementCaller) error {
	zap.L().Debug("starting network parameters check service")

	checkParamsCron := cron.New()

	_, err := checkParamsCron.AddFunc("@every 5m", func() {
		zap.L().Debug("running scheduled network parameters check")

		localEpoch, err := parameter.GetCurrentEpoch(ctx, redisClient)
		if err != nil {
			zap.L().Error("Failed to get current epoch from local storage",
				zap.Error(err))
			return
		}

		remoteEpoch, err := parameter.GetCurrentEpochFromVSL(settlementCaller)
		if err != nil {
			zap.L().Error("Failed to get current epoch from VSL",
				zap.Error(err))
			return
		}

		if remoteEpoch > localEpoch {
			zap.L().Info("Updating local epoch to match remote",
				zap.Int64("local_epoch", localEpoch),
				zap.Int64("remote_epoch", remoteEpoch))

			err = parameter.UpdateCurrentEpoch(ctx, redisClient, remoteEpoch)
			if err != nil {
				zap.L().Error("Failed to update current epoch",
					zap.Error(err))
				return
			}
		}

		if err := parameter.CheckParamsTask(ctx, redisClient, networkParamsCaller); err != nil {
			zap.L().Error("Failed to check network parameters",
				zap.Error(err))
		}
	})
	if err != nil {
		return fmt.Errorf("add check param cron job: %w", err)
	}

	zap.L().Info("Starting network parameters check cron job")
	checkParamsCron.Start()

	<-ctx.Done()
	zap.L().Info("Stopping network parameters check cron job")
	checkParamsCron.Stop()

	return ctx.Err()
}
