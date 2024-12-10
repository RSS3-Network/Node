package atproto

import (
	"context"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/docs"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/schema/worker/atproto"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

func (c *Component) GetAccountActivities(ctx echo.Context, account string, request docs.GetAtprotoAccountParams) (err error) {
	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, account)

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, account)

	zap.L().Debug("processing atproto account activities request",
		zap.Any("request", request))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get atproto account activities cursor",
			zap.String("account", account),
			zap.String("cursor", lo.FromPtr(request.Cursor)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owner:          lo.ToPtr(account),
		Limit:          request.Limit,
		ActionLimit:    request.ActionLimit,
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms: lo.Uniq(lo.Map(request.Platform, func(platform atproto.Platform, _ int) string {
			return platform.String()
		})),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get atproto account activities",
			zap.String("account", account),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved atproto account activities",
		zap.String("account", account),
		zap.Int("count", len(activities)))

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}

func (c *Component) TransformActivities(ctx context.Context, activities []*activityx.Activity) []*activityx.Activity {
	results := make([]*activityx.Activity, len(activities))

	// iterate over the activities
	// 1. transform the activity such as adding related urls and filling the author url
	// 2. query etherface for the transaction to get parsed function name
	lop.ForEach(activities, func(_ *activityx.Activity, index int) {
		result, err := c.TransformActivity(ctx, activities[index])
		if err != nil {
			zap.L().Error("failed to transform atproto activity",
				zap.String("ID", activities[index].ID),
				zap.Error(err))

			return
		}

		results[index] = result
	})

	return results
}

type ActivityResponse struct {
	Data *activityx.Activity `json:"data"`
	Meta *MetaTotalPages     `json:"meta"`
}

type ActivitiesResponse struct {
	Data []*activityx.Activity `json:"data"`
	Meta *MetaCursor           `json:"meta,omitempty"`
}

type MetaTotalPages struct {
	TotalPages int `json:"totalPages"`
}

type MetaCursor struct {
	Cursor string `json:"cursor"`
}
