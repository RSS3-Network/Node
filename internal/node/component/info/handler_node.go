package info

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/internal/constant"
	"go.uber.org/zap"
)

type VersionResponse struct {
	Data Version `json:"data"`
}

type Version struct {
	Tag    string `json:"tag"`
	Commit string `json:"commit"`
}

// GetNodeOperator returns the node information.
func (c *Component) GetNodeOperator(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", c.config.Discovery.Operator.EvmAddress)

	return ctx.JSON(http.StatusOK, response)
}

// GetVersion returns the version of the network component.
func (c *Component) GetVersion(ctx echo.Context) error {
	tag, commit := constant.BuildVersionDetail()

	return ctx.JSON(http.StatusOK, VersionResponse{
		Data: Version{
			Tag:    tag,
			Commit: commit,
		},
	})
}
