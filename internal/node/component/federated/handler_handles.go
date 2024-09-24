package federated

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
)

// GetAllActiveHandles retrieves all active handles
func (c *Component) GetAllActiveHandles(ctx echo.Context) error {
	handles, err := c.getAllHandles(ctx.Request().Context())
	if err != nil {
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"handles": handles})
}

// GetUpdatedHandles retrieves handles updated since a given timestamp
func (c *Component) GetUpdatedHandles(ctx echo.Context) error {
	sinceStr := ctx.QueryParam("since")
	since, err := strconv.ParseUint(sinceStr, 10, 64)

	if err != nil {
		return response.BadRequestError(ctx, err)
	}

	handles, err := c.getUpdatedHandles(ctx.Request().Context(), since)
	if err != nil {
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"handles": handles})
}
