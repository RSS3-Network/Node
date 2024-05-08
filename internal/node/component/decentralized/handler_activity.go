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

func (h *Component) GetActivity(c echo.Context) error {
	var request ActivityRequest

	if err := c.Bind(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if err := c.Validate(&request); err != nil {
		return response.ValidateFailedError(c, err)
	}

	query := model.ActivityQuery{
		ID:          lo.ToPtr(request.ID),
		ActionLimit: request.ActionLimit,
		ActionPage:  request.ActionPage,
	}

	activity, page, err := h.getActivity(c.Request().Context(), query)
	if err != nil {
		return response.InternalError(c, err)
	}

	// query etherface for the transaction
	if h.etherfaceClient != nil && activity != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
		activity.Calldata.ParsedFunction, _ = h.etherfaceClient.Lookup(c.Request().Context(), activity.Calldata.FunctionHash)
	}

	return c.JSON(http.StatusOK, ActivityResponse{
		Data: activity,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (h *Component) GetAccountActivities(c echo.Context) (err error) {
	request := AccountActivitiesRequest{}

	if err = c.Bind(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if request.Type, err = h.parseParams(c.QueryParams(), request.Tag); err != nil {
		return response.BadRequestError(c, err)
	}

	if err = defaults.Set(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if err = c.Validate(&request); err != nil {
		return response.ValidateFailedError(c, err)
	}

	go h.CollectMetric(context.TODO(), common.HexToAddress(request.Account).String())

	cursor, err := h.getCursor(c.Request().Context(), request.Cursor)
	if err != nil {
		return response.InternalError(c, err)
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

	activities, last, err := h.getActivities(c.Request().Context(), databaseRequest)
	if err != nil {
		return response.InternalError(c, err)
	}

	// iterate over the activities and query etherface for the transaction
	for _, activity := range activities {
		if h.etherfaceClient != nil && activity.Type == typex.Unknown && activity.Calldata != nil {
			activity.Calldata.ParsedFunction, _ = h.etherfaceClient.Lookup(c.Request().Context(), activity.Calldata.FunctionHash)
		}
	}

	return c.JSON(http.StatusOK, ActivitiesResponse{
		Data: activities,
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}
