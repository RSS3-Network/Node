package decentralized

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/common/http/response"
	"github.com/naturalselectionlabs/rss3-node/internal/database/model"
	"github.com/samber/lo"
)

func (h *Hub) GetActivity(ctx context.Context, request ActivityRequest) ActivityResponseInterface {
	query := model.FeedQuery{
		ID:          lo.ToPtr(request.ID),
		ActionLimit: request.ActionLimit,
		ActionPage:  request.ActionPage,
	}

	activity, page, err := h.getFeed(ctx, query)
	if err != nil {
		return response.InternalErrorResponse()
	}

	return ActivityResponse{
		Data: activity,
		Meta: lo.Ternary(page == nil, nil, &MetaTotalPages{
			TotalPages: lo.FromPtr(page),
		}),
	}
}

func (h *Hub) GetAccountActivities(ctx context.Context, request AccountActivitiesRequest) ActivitiesResponseInterface {
	// TODO domain resolve
	cursor, err := h.getCursor(ctx, request.Cursor)
	if err != nil {
		return response.BadRequestErrorResponse()
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

	activities, err := h.getFeeds(ctx, databaseRequest)
	if err != nil {
		return response.InternalErrorResponse()
	}

	return ActivitiesResponse{
		Data: activities,
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: activities[len(activities)-1].ID,
		}),
	}
}
