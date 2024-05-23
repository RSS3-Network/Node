package info

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// GetNodeOperator returns the node information.
func (c *Component) GetNodeOperator(ctx echo.Context) error {
	go c.CollectMetric(ctx.Request().Context(), "info")

	zap.L().Debug("get node info", zap.String("request.ip", ctx.Request().RemoteAddr))

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", c.config.Discovery.Operator.EvmAddress)

	return ctx.JSON(http.StatusOK, response)
}
