package decentralized

import (
	"context"
	"net/http"
	"strconv"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/docs"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/utils"
	"github.com/rss3-network/node/schema/worker/decentralized"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

func (c *Component) GetActivity(ctx echo.Context, id string, request docs.GetDecentralizedTxIDParams) error {
	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, id)

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, id)

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing get decentralized activity request",
		zap.Any("request", request))

	query := model.ActivityQuery{
		ID:          lo.ToPtr(id),
		ActionLimit: lo.FromPtr(request.ActionLimit),
		ActionPage:  lo.FromPtr(request.ActionPage),
	}

	activity, page, err := c.getActivity(ctx.Request().Context(), query)
	if err != nil {
		zap.L().Error("failed to get decentralized activity",
			zap.String("id", id),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	// query etherface for the transaction
	if c.etherfaceClient != nil && activity != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
		zap.L().Debug("querying etherface for function hash",
			zap.String("function_hash", activity.Calldata.FunctionHash))

		activity.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx.Request().Context(), activity.Calldata.FunctionHash)
	}

	// transform the activity such as adding related urls
	result, err := c.TransformActivity(ctx.Request().Context(), activity)
	if err != nil {
		zap.L().Error("failed to transform decentralized activity",
			zap.String("id", id),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved decentralized activity",
		zap.String("id", id))

	return ctx.JSON(http.StatusOK, ActivityResponse{
		Data: result,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (c *Component) GetAccountActivities(ctx echo.Context, account string, request docs.GetDecentralizedAccountParams) (err error) {
	if request.Type, err = utils.ParseTypes(ctx.QueryParams()["type"], request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, common.HexToAddress(account).String())

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, common.HexToAddress(account).String())

	addRecentRequest(ctx.Request().RequestURI)

	zap.L().Debug("processing get decentralized account activities request",
		zap.Any("request", request))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get decentralized account activities cursor",
			zap.String("cursor", lo.FromPtr(request.Cursor)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owner:          lo.ToPtr(common.HexToAddress(account).String()),
		Limit:          lo.FromPtr(request.Limit),
		ActionLimit:    lo.FromPtr(request.ActionLimit),
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms: lo.Uniq(lo.Map(request.Platform, func(platform decentralized.Platform, _ int) string {
			return platform.String()
		})),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get decentralized account activities",
			zap.String("account", account),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved decentralized account activities",
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
	var request docs.PostDecentralizedAccountsJSONRequestBody

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	for i := range request.Accounts {
		if common.IsHexAddress(request.Accounts[i]) {
			request.Accounts[i] = common.HexToAddress(request.Accounts[i]).String()
		}
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

	zap.L().Debug("processing batch get decentralized accounts activities request",
		zap.Any("request", request))

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("failed to get decentralized accounts activities cursor",
			zap.String("cursor", lo.FromPtr(request.Cursor)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owners:         lo.Uniq(request.Accounts),
		Limit:          lo.FromPtr(request.Limit),
		ActionLimit:    lo.FromPtr(request.ActionLimit),
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(types),
		Platforms: lo.Uniq(lo.Map(request.Platform, func(platform decentralized.Platform, _ int) string {
			return platform.String()
		})),
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get activities",
			zap.Int("accounts_count", len(request.Accounts)),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved decentralized accounts activities",
		zap.Int("accounts_count", len(request.Accounts)),
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
			zap.L().Error("failed to transform decentralized activity",
				zap.String("id", activities[index].ID),
				zap.Error(err))

			return
		}

		// query etherface to get the parsed function name
		if c.etherfaceClient != nil && result.Type == typex.Unknown && result.Calldata != nil {
			zap.L().Debug("querying etherface for function hash",
				zap.String("function_hash", result.Calldata.FunctionHash))

			result.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx, result.Calldata.FunctionHash)
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
