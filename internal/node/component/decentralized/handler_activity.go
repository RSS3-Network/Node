package decentralized

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
)

type ActivityRequest struct {
	ID          string `param:"id"`
	ActionLimit int    `query:"action_limit"  validate:"min=1,max=20" default:"10"`
	ActionPage  int    `query:"action_page" validate:"min=1" default:"1"`
}

type AccountActivitiesRequest struct {
	Account        string               `param:"account" validate:"required"`
	Limit          int                  `query:"limit" validate:"min=1,max=100" default:"100"`
	ActionLimit    int                  `query:"action_limit" validate:"min=1,max=20" default:"10"`
	Cursor         *string              `query:"cursor"`
	SinceTimestamp *uint64              `query:"since_timestamp"`
	UntilTimestamp *uint64              `query:"until_timestamp"`
	Status         *bool                `query:"success"`
	Direction      *activityx.Direction `query:"direction"`
	Network        []network.Network    `query:"network"`
	Tag            []tag.Tag            `query:"tag"`
	Type           []schema.Type        `query:"-"`
	Platform       []string             `query:"platform"`
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

	go c.CollectMetric(ctx.Request().Context(), common.HexToAddress(request.Account).String())

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

func (c *Component) parseParams(params url.Values, tags []tag.Tag) ([]schema.Type, error) {
	if len(tags) == 0 {
		return nil, nil
	}

	types := make([]schema.Type, 0)

	for _, typex := range params["type"] {
		var (
			value schema.Type
			err   error
		)

		for _, tagx := range tags {
			value, err = schema.ParseTypeFromString(tagx, typex)
			if err == nil {
				types = append(types, value)

				break
			}
		}

		if err != nil {
			return nil, fmt.Errorf("invalid type: %s", typex)
		}
	}

	return types, nil
}
