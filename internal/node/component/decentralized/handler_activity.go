package decentralized

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

func (c *Component) GetActivity(ctx echo.Context) error {
	var request ActivityRequest

	if err := ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err := ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, request.ID)

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, request.ID)

	addRecentRequest(ctx.Request().RequestURI)

	query := model.ActivityQuery{
		ID:          lo.ToPtr(request.ID),
		ActionLimit: request.ActionLimit,
		ActionPage:  request.ActionPage,
	}

	activity, page, err := c.getActivity(ctx.Request().Context(), query)
	if err != nil {
		zap.L().Error("GetActivity InternalError", zap.Error(err))
		return response.InternalError(ctx)
	}

	// query etherface for the transaction
	if c.etherfaceClient != nil && activity != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
		activity.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx.Request().Context(), activity.Calldata.FunctionHash)
	}

	// transform the activity such as adding related urls
	result, err := c.TransformActivity(ctx.Request().Context(), activity)
	if err != nil {
		zap.L().Error("TransformActivity InternalError", zap.Error(err))
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, ActivityResponse{
		Data: result,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (c *Component) GetAccountActivities(ctx echo.Context) (err error) {
	var request AccountActivitiesRequest

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if request.Type, err = c.parseTypes(ctx.QueryParams()["type"], request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidationFailedError(ctx, err)
	}

	go c.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, common.HexToAddress(request.Account).String())

	go c.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, common.HexToAddress(request.Account).String())

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
		Owner:          lo.ToPtr(common.HexToAddress(request.Account).String()),
		Limit:          request.Limit,
		ActionLimit:    request.ActionLimit,
		Status:         request.Status,
		Direction:      request.Direction,
		Network:        lo.Uniq(request.Network),
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

// BatchGetAccountsActivities returns the activities of multiple accounts in a single request
func (c *Component) BatchGetAccountsActivities(ctx echo.Context) (err error) {
	var request AccountsActivitiesRequest

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	types, err := c.parseTypes(request.Type, request.Tag)
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

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		zap.L().Error("getCursor InternalError", zap.Error(err))
		return response.InternalError(ctx)
	}

	databaseRequest := model.ActivitiesQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owners: lo.Uniq(lo.Map(request.Accounts, func(account string, _ int) string {
			return common.HexToAddress(account).String()
		})),
		Limit:       request.Limit,
		ActionLimit: request.ActionLimit,
		Status:      request.Status,
		Direction:   request.Direction,
		Network:     lo.Uniq(request.Network),
		Tags:        lo.Uniq(request.Tag),
		Types:       lo.Uniq(types),
		Platforms:   lo.Uniq(request.Platform),
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

func (c *Component) TransformActivities(ctx context.Context, activities []*activityx.Activity) []*activityx.Activity {
	results := make([]*activityx.Activity, len(activities))

	// iterate over the activities
	// 1. transform the activity such as adding related urls and filling the author url
	// 2. query etherface for the transaction to get parsed function name
	lop.ForEach(activities, func(_ *activityx.Activity, index int) {
		result, err := c.TransformActivity(ctx, activities[index])
		if err != nil {
			zap.L().Error("failed to load activity", zap.Error(err))

			return
		}

		// query etherface to get the parsed function name
		if c.etherfaceClient != nil && result.Type == typex.Unknown && result.Calldata != nil {
			result.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx, result.Calldata.FunctionHash)
		}

		results[index] = result
	})

	return results
}

func (c *Component) parseTypes(types []string, tags []tag.Tag) ([]schema.Type, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	schemaTypes := make([]schema.Type, 0)

	for _, typex := range types {
		var (
			value schema.Type
			err   error
		)

		for _, tagx := range tags {
			value, err = schema.ParseTypeFromString(tagx, typex)
			if err == nil {
				schemaTypes = append(schemaTypes, value)

				break
			}
		}

		if err != nil {
			return nil, fmt.Errorf("invalid type: %s", typex)
		}
	}

	return schemaTypes, nil
}

type ActivityRequest struct {
	ID          string `param:"id"`
	ActionLimit int    `query:"action_limit"  validate:"min=1,max=20" default:"10"`
	ActionPage  int    `query:"action_page" validate:"min=1" default:"1"`
}

type AccountActivitiesRequest struct {
	Account        string                   `param:"account" validate:"required"`
	Limit          int                      `query:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                      `query:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string                  `query:"cursor"`
	SinceTimestamp *uint64                  `query:"since_timestamp"`
	UntilTimestamp *uint64                  `query:"until_timestamp"`
	Status         *bool                    `query:"success"`
	Direction      *activityx.Direction     `query:"direction"`
	Network        []network.Network        `query:"network"`
	Tag            []tag.Tag                `query:"tag"`
	Type           []schema.Type            `query:"-"`
	Platform       []decentralized.Platform `query:"platform"`
}

type AccountsActivitiesRequest struct {
	Accounts       []string                 `json:"accounts" validate:"required"`
	Limit          int                      `json:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                      `json:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string                  `json:"cursor"`
	SinceTimestamp *uint64                  `json:"since_timestamp"`
	UntilTimestamp *uint64                  `json:"until_timestamp"`
	Status         *bool                    `json:"success"`
	Direction      *activityx.Direction     `json:"direction"`
	Network        []network.Network        `json:"network"`
	Tag            []tag.Tag                `json:"tag"`
	Type           []string                 `json:"type"`
	Platform       []decentralized.Platform `json:"platform"`
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
