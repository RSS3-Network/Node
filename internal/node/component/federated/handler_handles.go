package federated

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
)

var limit = 10

// GetHandles retrieves all active handles or updated handles based on the 'since' parameter
func (c *Component) GetHandles(ctx echo.Context) error {
	var request HandleRequest
	if err := ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	// Validate request
	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	var handles []string

	var paginatedHandles *model.PaginatedMastodonHandles

	// Fetch either all handles or updated handles based on 'since' value
	paginatedHandles, err := c.getUpdatedHandles(ctx.Request().Context(), request.Since, limit, request.Cursor)

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

type HandleRequest struct {
	Since  uint64 `query:"since" validate:"required"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Cursor string `query:"cursor"`
}

type PaginatedHandlesResponse struct {
	Handles    []string `json:"handles"`
	Cursor     string   `json:"cursor"`
	TotalCount int64    `json:"total_count"`
}
