package federated

import (
	"net/http"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/docs"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/utils"
	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (c *Component) GetPlatformActivities(ctx echo.Context, plat federated.Platform, request docs.GetFederatedPlatformParams) (err error) {
	if request.Type, err = utils.ParseTypes(ctx.QueryParams()["type"], request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, plat.String())

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, plat.String())

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing federated platform activities request",
		zap.String("platform", plat.String()),
		zap.Int("limit", lo.FromPtr(request.Limit)),
		zap.String("cursor", lo.FromPtr(request.Cursor)))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get federated platform activities cursor",
			zap.String("platform", plat.String()),
			zap.String("cursor", lo.FromPtr(request.Cursor)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Limit:          lo.FromPtr(request.Limit),
		ActionLimit:    lo.FromPtr(request.ActionLimit),
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms:      []string{plat.String()},
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get federated platform activities",
			zap.String("platform", plat.String()),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved federated platform activities",
		zap.String("platform", plat.String()),
		zap.Int("count", len(activities)))

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}
