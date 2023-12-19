package decentralized

import (
	"github.com/creasty/defaults"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-node/common/http/response"
	"github.com/naturalselectionlabs/rss3-node/internal/database/model"
	"github.com/samber/lo"
)

func (h *Hub) GetActivity(c echo.Context) error {
	request := ActivityRequest{}

	if err := c.Bind(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if err := c.Validate(&request); err != nil {
		return response.ValidateFailedError(c, err)
	}

	query := model.FeedQuery{
		ID:          lo.ToPtr(request.ID),
		ActionLimit: request.ActionLimit,
		ActionPage:  request.ActionPage,
	}

	activity, page, err := h.getFeed(c.Request().Context(), query)
	if err != nil {
		return response.InternalError(c, err)
	}

	return c.JSON(http.StatusOK, ActivityResponse{
		Data: activity,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	})
}

func (h *Hub) GetAccountActivities(c echo.Context) error {
	request := AccountActivitiesRequest{}

	if err := c.Bind(&request); err != nil {
		return response.BadRequestError(c, err)
	}

	if err := c.Validate(&request); err != nil {
		return response.ValidateFailedError(c, err)
	}

	if err := defaults.Set(&request); err != nil {
		return response.InternalError(c, err)
	}

	cursor, err := h.getCursor(c.Request().Context(), request.Cursor)
	if err != nil {
		return response.InternalError(c, err)
	}

	databaseRequest := model.FeedsQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Owner:          lo.ToPtr(common.HexToAddress(request.Account).String()),
		Limit:          request.Limit,
		ActionLimit:    request.ActionLimit,
		Status:         request.Status,
		Direction:      request.Direction,
		Tags:           lo.Uniq(request.Tag),
		Types:          lo.Uniq(request.Type),
		Platforms:      lo.Uniq(request.Platform),
	}

	activities, last, err := h.getFeeds(c.Request().Context(), databaseRequest)
	if err != nil {
		return response.InternalError(c, err)
	}

	return c.JSON(http.StatusOK, ActivitiesResponse{
		Data: activities,
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}
