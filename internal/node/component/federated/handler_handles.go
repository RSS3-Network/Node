package federated

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
)

var limit = 10

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

	// Get 'limit' and 'cursor' query parameters for pagination
	limitStr := ctx.QueryParam("limit")
	cursor := ctx.QueryParam("cursor")

	// Default limit if not provided
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			return response.BadRequestError(ctx, err)
		}
	}

	var handles []string

	var paginatedHandles *model.PaginatedMastodonHandles

	// Fetch either all handles or updated handles based on 'since' value
	paginatedHandles, err = c.getUpdatedHandles(ctx.Request().Context(), since, limit, cursor)

	if err != nil {
		return response.InternalError(ctx)
	}

	// Extract handles from paginated result
	handles = paginatedHandles.Handles

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"handles":     handles,
		"cursor":      paginatedHandles.NextCursor,
		"total_count": paginatedHandles.TotalCount,
	})
}
