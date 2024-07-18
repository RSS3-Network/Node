package info

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/internal/constant"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type VersionResponse struct {
	Data Version `json:"data"`
}

type Version struct {
	Tag    string `json:"tag"`
	Commit string `json:"commit"`
}

type NodeInfoResponse struct {
	Data NodeInfo `json:"data"`
}

type NodeInfo struct {
	Operator common.Address `json:"operator"`
	Version  Version        `json:"version"`
	Uptime   string         `json:"uptime"`
}

// GetNodeInfo returns the node information.
func (c *Component) GetNodeInfo(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	version := c.buildVersion()

	evmAddress := common.Address{}

	if !lo.IsEmpty(c.config.Discovery.Operator.EvmAddress) {
		evmAddress = c.config.Discovery.Operator.EvmAddress
	}

	return ctx.JSON(http.StatusOK, NodeInfoResponse{
		Data: NodeInfo{
			Version:  version,
			Operator: evmAddress,
		},
	})
}

// GetVersion returns the version of the network component.
func (c *Component) GetVersion(ctx echo.Context) error {
	version := c.buildVersion()

	return ctx.JSON(http.StatusOK, VersionResponse{
		Data: version,
	})
}

func (c *Component) buildVersion() Version {
	tag, commit := constant.BuildVersionDetail()

	return Version{
		Tag:    tag,
		Commit: commit,
	}
}
