package decentralized

import (
	"context"
	"net/http"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
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
		return response.ValidateFailedError(ctx, err)
	}

	query := model.ActivityQuery{
		ID:          lo.ToPtr(request.ID),
		ActionLimit: request.ActionLimit,
		ActionPage:  request.ActionPage,
	}

	activity, page, err := c.getActivity(ctx.Request().Context(), query)
	if err != nil {
		return response.InternalError(ctx, err)
	}

	// query etherface for the transaction
	if c.etherfaceClient != nil && activity != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
		activity.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx.Request().Context(), activity.Calldata.FunctionHash)
	}

	return ctx.JSON(http.StatusOK, ActivityResponse{
		Data: activity,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (c *Component) GetAccountActivities(ctx echo.Context) (err error) {
	request := AccountActivitiesRequest{}

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if request.Type, err = c.parseParams(ctx.QueryParams(), request.Tag); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if err = ctx.Validate(&request); err != nil {
		return response.ValidateFailedError(ctx, err)
	}

	go c.CollectMetric(context.TODO(), common.HexToAddress(request.Account).String())

	cursor, err := c.getCursor(ctx.Request().Context(), request.Cursor)
	if err != nil {
		return response.InternalError(ctx, err)
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
		return response.InternalError(ctx, err)
	}

	// iterate over the activities and query etherface for the transaction
	for _, activity := range activities {
		if c.etherfaceClient != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
			activity.Calldata.ParsedFunction, _ = c.etherfaceClient.Lookup(ctx.Request().Context(), activity.Calldata.FunctionHash)
		}
	}

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: activities,
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}
