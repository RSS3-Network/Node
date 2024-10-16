package federated

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/samber/lo"
)

var defaultLimit = 100

// GetHandles retrieves all active handles or updated handles based on the 'since' parameter
func (c *Component) GetHandles(ctx echo.Context) error {
	var request HandleRequest
	if err := ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if request.Limit == 0 {
		request.Limit = defaultLimit
	}
	// Validate request
	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	query := model.QueryMastodonHandles{
		Limit:  lo.ToPtr(request.Limit),
		Since:  lo.Ternary(request.Since > 0, lo.ToPtr(request.Since), nil),
		Cursor: lo.Ternary(request.Cursor != "", lo.ToPtr(request.Cursor), nil),
	}

	res, err := c.getUpdatedHandles(ctx.Request().Context(), query)
	if err != nil {
		return response.InternalError(ctx)
	}

	handles := make([]string, 0, len(res))

	var cursor string

	for i, handle := range res {
		handles = append(handles, handle.Handle)

		if i == len(res)-1 && len(res) == request.Limit {
			cursor = handle.Handle
		}
	}

	return ctx.JSON(http.StatusOK, PaginatedHandlesResponse{
		Handles:    handles,
		Cursor:     cursor,
		TotalCount: int64(len(handles)),
	})
}

type HandleRequest struct {
	Since  uint64 `query:"since"`
	Limit  int    `query:"limit" validate:"omitempty,min=1,max=500"`
	Cursor string `query:"cursor"`
}

type PaginatedHandlesResponse struct {
	Handles    []string `json:"handles"`
	Cursor     string   `json:"cursor,omitempty"`
	TotalCount int64    `json:"total_count"`
}
