package aggregator

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/v2/docs"
	"github.com/rss3-network/node/v2/internal/node/component/ai"
	"github.com/rss3-network/node/v2/internal/node/component/decentralized"
	"github.com/rss3-network/node/v2/internal/node/component/federated"
	"github.com/rss3-network/node/v2/internal/node/component/info"
	"github.com/rss3-network/node/v2/internal/node/component/rss"
	decentralizedx "github.com/rss3-network/node/v2/schema/worker/decentralized"
	federatedx "github.com/rss3-network/node/v2/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/network"
)

type Component struct {
	Info          *info.Component
	Decentralized *decentralized.Component
	Federated     *federated.Component
	RSS           *rss.Component
	AI            *ai.Component
}

var _ docs.ServerInterface = (*Component)(nil)

// emptyResponse returns an empty response with a 200 status code
func emptyResponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": []interface{}{},
		"meta": nil,
	})
}

// Decentralized Interface

func (c Component) PostDecentralizedAccounts(ctx echo.Context) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.BatchGetAccountsActivities(ctx)
}

func (c Component) PostDecentralizedMetadata(ctx echo.Context) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.BatchGetMetadataActivities(ctx)
}

func (c Component) GetDecentralizedNetwork(ctx echo.Context, network network.Network, params docs.GetDecentralizedNetworkParams) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.GetNetworkActivities(ctx, network, params)
}

func (c Component) GetDecentralizedPlatform(ctx echo.Context, platform decentralizedx.Platform, params docs.GetDecentralizedPlatformParams) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.GetPlatformActivities(ctx, platform, params)
}

func (c Component) GetDecentralizedTxID(ctx echo.Context, id string, params docs.GetDecentralizedTxIDParams) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.GetActivity(ctx, id, params)
}

func (c Component) GetDecentralizedAccount(ctx echo.Context, account string, params docs.GetDecentralizedAccountParams) error {
	if c.Decentralized == nil {
		return emptyResponse(ctx)
	}

	return c.Decentralized.GetAccountActivities(ctx, account, params)
}

// Federated Interface

func (c Component) PostFederatedAccounts(ctx echo.Context) error {
	if c.Federated == nil {
		return emptyResponse(ctx)
	}

	return c.Federated.BatchGetAccountsActivities(ctx)
}

func (c Component) GetFederatedNetwork(ctx echo.Context, network network.Network, params docs.GetFederatedNetworkParams) error {
	if c.Federated == nil {
		return emptyResponse(ctx)
	}

	return c.Federated.GetNetworkActivities(ctx, network, params)
}

func (c Component) GetFederatedPlatform(ctx echo.Context, platform federatedx.Platform, params docs.GetFederatedPlatformParams) error {
	if c.Federated == nil {
		return emptyResponse(ctx)
	}

	return c.Federated.GetPlatformActivities(ctx, platform, params)
}

func (c Component) GetFederatedTxID(ctx echo.Context, id string, params docs.GetFederatedTxIDParams) error {
	if c.Federated == nil {
		return emptyResponse(ctx)
	}

	return c.Federated.GetActivity(ctx, id, params)
}

func (c Component) GetFederatedAccount(ctx echo.Context, account string, params docs.GetFederatedAccountParams) error {
	if c.Federated == nil {
		return emptyResponse(ctx)
	}

	return c.Federated.GetAccountActivities(ctx, account, params)
}

// RSS Interface

// GetRSS ignore the second parameter
// manually create rss cause of wildcard routes are not part of the official open api specification
// https://github.com/oapi-codegen/oapi-codegen/issues/718
func (c Component) GetRSS(ctx echo.Context, _ string) error {
	if c.RSS == nil {
		return emptyResponse(ctx)
	}

	return c.RSS.Handler(ctx)
}

// AI Interface

// GetAgentData ignore the second parameter
func (c Component) GetAgentData(ctx echo.Context, _ string) error {
	if c.AI == nil {
		return emptyResponse(ctx)
	}

	return c.AI.Handler(ctx)
}

// Info Interface

func (c Component) GetNodeInfo(ctx echo.Context) error {
	return c.Info.GetNodeInfo(ctx)
}

func (c Component) GetWorkersStatus(ctx echo.Context) error {
	return c.Info.GetWorkersStatus(ctx)
}

func (c Component) GetNodeOperatorInfo(ctx echo.Context) error {
	return c.Info.GetNodeOperator(ctx)
}

func (c Component) GetActivityCount(ctx echo.Context) error {
	return c.Info.GetActivityCount(ctx)
}

func (c Component) GetNetworksConfig(ctx echo.Context) error {
	return c.Info.GetNetworkConfig(ctx)
}
