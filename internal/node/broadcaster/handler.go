package broadcaster

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func (b *Broadcaster) GetNodeInfo(c echo.Context) error {
	zap.L().Debug("get node info", zap.String("request.ip", c.Request().RemoteAddr))

	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", b.config.Discovery.Maintainer.EvmAddress)

	return c.JSON(http.StatusOK, response)
}
