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

	if len(request.RawType) == 0 || request.Tag == nil {
		return response.BadRequestError(ctx, fmt.Errorf("empty tag or type"))
	}

	typex, err := schema.ParseTypeFromString(lo.FromPtr(request.Tag), request.RawType)
	if err != nil {
		return response.BadRequestError(ctx, fmt.Errorf("invalid type: %s", request.RawType))
	}

	request.Type = lo.ToPtr(typex)

	meta, err := metadata.Unmarshal(typex, request.RawMetadata)
	if err != nil {
		return response.BadRequestError(ctx, fmt.Errorf("failed to unmarshal metadata: %w", err))
	}

	request.Metadata = lo.ToPtr(meta)

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

	databaseRequest := model.ActivitiesMetadataQuery{
		Cursor:         cursor,
		StartTimestamp: request.SinceTimestamp,
		EndTimestamp:   request.UntilTimestamp,
		Limit:          lo.FromPtr(request.Limit),
		ActionLimit:    lo.FromPtr(request.ActionLimit),
		Accounts:       request.Accounts,
		Platform:       request.Platform,
		Status:         request.Status,
		Network:        request.Network,
		Metadata:       request.Metadata,
	}

	activities, last, err := c.getActivitiesMetadata(ctx.Request().Context(), databaseRequest)
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
