package federated

import (
	"context"
	"net/http"
	"strconv"

	"github.com/creasty/defaults"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/v2/common/http/response"
	"github.com/rss3-network/node/v2/docs"
	"github.com/rss3-network/node/v2/internal/database/model"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/schema/worker/federated"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

func (c *Component) GetActivity(ctx echo.Context, id string, request docs.GetFederatedTxIDParams) error {
	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, id)

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, id)

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing federated activity request",
		zap.Any("request", request))

	query := model.ActivityQuery{
		ID:          lo.ToPtr(id),
		ActionLimit: lo.FromPtr(request.ActionLimit),
		ActionPage:  lo.FromPtr(request.ActionPage),
	}

	activity, page, err := c.getActivity(ctx.Request().Context(), query)
	if err != nil {
		zap.L().Error("failed to get federated activity",
			zap.String("id", id),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	// transform the activity such as adding related urls
	result, err := c.TransformActivity(ctx.Request().Context(), activity)
	if err != nil {
		zap.L().Error("failed to transform federated activity",
			zap.String("id", id),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved federated activity",
		zap.String("id", id))

	return ctx.JSON(http.StatusOK, ActivityResponse{
		Data: result,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (c *Component) GetAccountActivities(ctx echo.Context, account string, request docs.GetFederatedAccountParams) (err error) {
	if request.Type, err = utils.ParseTypes(ctx.QueryParams()["type"], request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, account)

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, account)

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing federated account activities request",
		zap.Any("request", request))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get federated account activities cursor",
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
		Limit:          lo.FromPtr(request.Limit),
		ActionLimit:    lo.FromPtr(request.ActionLimit),
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms: lo.Uniq(lo.Map(request.Platform, func(platform federated.Platform, _ int) string {
			return platform.String()
		})),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get federated account activities",
			zap.String("account", account),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved federated account activities",
		zap.String("account", account),
		zap.Int("count", len(activities)))

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}

// BatchGetAccountsActivities returns the activities of multiple accounts in a single request
func (c *Component) BatchGetAccountsActivities(ctx echo.Context) (err error) {
	var request docs.PostFederatedAccountsJSONBody

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	types, err := utils.ParseTypes(request.Type, request.Tag)
	if err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, strconv.Itoa(len(request.Accounts)))

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, strconv.Itoa(len(request.Accounts)))

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing federated batch accounts activities request",
		zap.Any("accounts_count", request.Accounts),
		zap.Int("limit", lo.FromPtr(request.Limit)))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get federated batch accounts activities cursor",
			zap.String("cursor", lo.FromPtr(request.Cursor)),
			zap.Int("accounts_count", len(request.Accounts)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owners: lo.Uniq(lo.Map(request.Accounts, func(account string, _ int) string {
			return account
		})),
		Limit:       lo.FromPtr(request.Limit),
		ActionLimit: lo.FromPtr(request.ActionLimit),
		Status:      request.Status,
		Direction:   request.Direction,
		Network:     lo.Uniq(request.Network),
		Tags:        lo.Uniq(request.Tag),
		Types:       lo.Uniq(types),
		Platforms: lo.Uniq(lo.Map(request.Platform, func(platform federated.Platform, _ int) string {
			return platform.String()
		})),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get federated batch accounts activities",
			zap.Int("accounts_count", len(request.Accounts)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Debug("successfully retrieved federated batch accounts activities",
		zap.Int("accounts_count", len(request.Accounts)),
		zap.Int("activities_count", len(activities)))

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
			zap.L().Error("failed to transform federated activity",
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
