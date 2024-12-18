package decentralized

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/creasty/defaults"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/node/docs"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// BatchGetMetadataActivities returns the activities of specified decentralized metadata
func (c *Component) BatchGetMetadataActivities(ctx echo.Context) (err error) {
	request := struct {
		docs.PostDecentralizedMetadataJSONBody
		RawType     string          `json:"type"`
		RawMetadata json.RawMessage `json:"metadata"`
	}{}

	if err = ctx.Bind(&request); err != nil {
		return response.BadRequestError(ctx, err)
	}

	if len(request.RawType) == 0 || request.Tag == tag.Unknown {
		return response.BadRequestError(ctx, fmt.Errorf("empty tag or type"))
	}

	request.Type, err = schema.ParseTypeFromString(request.Tag, request.RawType)
	if err != nil {
		return response.BadRequestError(ctx, fmt.Errorf("invalid type: %s", request.RawType))
	}

	request.Metadata, err = metadata.Unmarshal(request.Type, request.RawMetadata)
	if err != nil {
		return response.BadRequestError(ctx, fmt.Errorf("failed to unmarshal metadata: %w", err))
	}

	for i := range request.Accounts {
		if common.IsHexAddress(request.Accounts[i]) {
			request.Accounts[i] = common.HexToAddress(request.Accounts[i]).String()
		}
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

		return response.BadRequestError(ctx, fmt.Errorf("invalid cursor: %w", err))
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
		Tags:           []tag.Tag{request.Tag},
		Types:          []schema.Type{request.Type},
		Platform:       request.Platform.String(),
		Metadata:       request.Metadata,
	}

	activities, last, err := c.getActivities(ctx.Request().Context(), databaseRequest)
	if err != nil {
		zap.L().Error("failed to get activities",
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved decentralized metadata activities",
		zap.Int("accounts_count", len(request.Accounts)),
		zap.Int("count", len(activities)))

	return ctx.JSON(http.StatusOK, ActivitiesResponse{
		Data: c.TransformActivities(ctx.Request().Context(), activities),
		Meta: lo.Ternary(len(activities) < databaseRequest.Limit, nil, &MetaCursor{
			Cursor: last,
		}),
	})
}
