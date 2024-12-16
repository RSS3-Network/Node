package federated

import (
	"context"
	"fmt"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// GetHandles retrieves all active handles or updated handles based on the 'since' parameter
func (c *Component) GetHandles(ctx echo.Context) error {
	var request HandleRequest
	if err := ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	zap.L().Debug("processing get handles request", zap.Any("request", request))

	// Validate request
	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	var (
		handles []string
		err     error
		cursor  string
	)

	switch request.Platform {
	case federated.PlatformMastodon:
		handles, err = c.getMastodonHandles(ctx.Request().Context(), request)
	case federated.PlatformBluesky:
		handles, err = c.getBlueskyHandles(ctx.Request().Context(), request)
	default:
		return response.BadRequestError(ctx, fmt.Errorf("unsupported platform"))
	}

	if err != nil {
		zap.L().Error("failed to get handles", zap.Error(err))

		return response.InternalError(ctx)
	}

	if len(handles) > 0 && len(handles) == request.Limit {
		cursor = handles[len(handles)-1]
	}

	return ctx.JSON(http.StatusOK, PaginatedHandlesResponse{
		Platform:   request.Platform,
		Handles:    handles,
		Cursor:     cursor,
		TotalCount: int64(len(handles)),
	})
}

func (c *Component) getMastodonHandles(ctx context.Context, request HandleRequest) ([]string, error) {
	query := model.QueryMastodonHandles{
		Limit:  lo.ToPtr(request.Limit),
		Since:  request.Since,
		Cursor: request.Cursor,
	}

	res, err := c.loadMastodonHandles(ctx, query)
	if err != nil {
		zap.L().Error("failed to get mastodon handles",
			zap.Error(err),
			zap.Any("query", query))

		return nil, err
	}

	handles := make([]string, 0, len(res))

	for _, handle := range res {
		handles = append(handles, handle.Handle)
	}

	return handles, nil
}

func (c *Component) getBlueskyHandles(ctx context.Context, request HandleRequest) ([]string, error) {
	query := model.QueryBlueskyProfiles{
		Limit:  lo.ToPtr(request.Limit),
		Since:  request.Since,
		Cursor: request.Cursor,
	}

	res, err := c.loadBlueskyProfiles(ctx, query)
	if err != nil {
		zap.L().Error("failed to get bluesky profiles",
			zap.Error(err),
			zap.Any("query", query))

		return nil, err
	}

	handles := make([]string, 0, len(res))

	for _, handle := range res {
		handles = append(handles, handle.Handle)
	}

	return handles, nil
}

type HandleRequest struct {
	Since    *uint64            `query:"since"`
	Limit    int                `query:"limit" default:"100" validate:"omitempty,min=1,max=500"`
	Cursor   *string            `query:"cursor"`
	Platform federated.Platform `query:"platform" default:"1" validate:"required"`
}

type PaginatedHandlesResponse struct {
	Platform   federated.Platform `json:"platform"`
	Handles    []string           `json:"handles"`
	Cursor     string             `json:"cursor,omitempty"`
	TotalCount int64              `json:"total_count"`
}
