package federated

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
)

// GetHandles retrieves all active handles or updated handles based on the 'since' parameter
func (c *Component) GetHandles(ctx echo.Context) error {
	sinceStr := ctx.QueryParam("since")
	if sinceStr == "" {
		return response.BadRequestError(ctx, errors.New("'since' parameter is required"))
	}

	since, err := strconv.ParseUint(sinceStr, 10, 64)
	if err != nil {
		return response.BadRequestError(ctx, err)
	}

	var handles []string
	if since == 0 {
		handles, err = c.getAllHandles(ctx.Request().Context())
	} else {
		handles, err = c.getUpdatedHandles(ctx.Request().Context(), since)
	}

	if err != nil {
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{"handles": handles})
}
