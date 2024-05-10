package broadcaster

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (b *Broadcaster) GetNodeInfo(c echo.Context) error {
	response := fmt.Sprintf("This is an RSS3 Node operated by %s.", b.config.Discovery.Maintainer.EvmAddress)
	return c.JSON(http.StatusOK, response)
}
