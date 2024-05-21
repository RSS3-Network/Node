package info

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// GetNodeInfo returns the node information.
func (c *Component) GetNodeInfo(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", c.config.Discovery.Maintainer.EvmAddress)

	return ctx.JSON(http.StatusOK, response)
}
