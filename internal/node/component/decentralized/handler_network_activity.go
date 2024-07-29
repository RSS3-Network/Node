package decentralized

import (
	"net/http"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

func (c *Component) GetNetworkActivities(ctx echo.Context) (err error) {
	var request NetworkActivitiesRequest

	if err := ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if request.Type, err = c.parseTypes(ctx.QueryParams()["type"], request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, request.Network.String())

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, request.Network.String())

	addRecentRequest(ctx.Request().RequestURI)

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("getCursor InternalError", zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Limit:          request.Limit,
		ActionLimit:    request.ActionLimit,
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        []network.Network{request.Network},
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms:      lo.Uniq(request.Platform),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("getActivities InternalError", zap.Error(err))

		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}

type NetworkActivitiesRequest struct {
	Network network.Network `param:"network" validate:"required"`

	Limit          int                      `query:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                      `query:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string                  `query:"cursor"`
	SinceTimestamp *uint64                  `query:"since_timestamp"`
	UntilTimestamp *uint64                  `query:"until_timestamp"`
	Status         *bool                    `query:"success"`
	Direction      *activityx.Direction     `query:"direction"`
	Tag            []tag.Tag                `query:"tag"`
	Type           []schema.Type            `query:"-"`
	Platform       []decentralized.Platform `query:"platform"`
}
