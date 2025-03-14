package decentralized

import (
	"net/http"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/v2/common/http/response"
	"github.com/rss3-network/node/v2/docs"
	"github.com/rss3-network/node/v2/internal/database/model"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (c *Component) GetPlatformActivities(ctx echo.Context, plat decentralized.Platform, request docs.GetDecentralizedPlatformParams) (err error) {
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

	zap.L().Debug("processing decentralized platform activities request",
		zap.String("platform", plat.String()),
		zap.Int("limit", lo.FromPtr(request.Limit)),
		zap.String("cursor", lo.FromPtr(request.Cursor)))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get decentralized platform activities cursor",
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
		zap.L().Error("failed to get decentralized platform activities",
			zap.String("platform", plat.String()),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved decentralized platform activities",
		zap.String("platform", plat.String()),
		zap.Int("count", len(activities)))

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}

type PlatformActivitiesRequest struct {
	Platform decentralized.Platform `param:"platform" validate:"required"`

	Limit          int                  `query:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                  `query:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string              `query:"cursor"`
	SinceTimestamp *uint64              `query:"since_timestamp"`
	UntilTimestamp *uint64              `query:"until_timestamp"`
	Status         *bool                `query:"success"`
	Direction      *activityx.Direction `query:"direction"`
	Tag            []tag.Tag            `query:"tag"`
	Type           []schema.Type        `query:"-"`
	Network        []network.Network    `query:"network"`
}
