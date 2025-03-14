package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/v2/common/http/response"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"go.uber.org/zap"
)

type Response struct {
	Data []*activityx.Activity `json:"data"`
}

func (h *Component) Handler(ctx echo.Context) error {
	path := ctx.Param("*")
	zap.L().Debug("handling RSS request",
		zap.String("path", path),
		zap.String("request_uri", ctx.Request().RequestURI))

	go h.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, path)

	go h.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, path)

	addRecentRequest(ctx.Request().RequestURI)

	data, err := h.getActivities(ctx.Request().Context(), path, ctx.Request().URL)
	if err != nil {
		zap.L().Error("failed to get activities from RSS feed",
			zap.String("path", path),
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Info("successfully retrieved RSS activities",
		zap.String("path", path),
		zap.Int("activity_count", len(data)))

	return ctx.JSON(http.StatusOK, Response{
		Data: data,
	})
}
