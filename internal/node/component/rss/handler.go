package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"go.uber.org/zap"
)

type Response struct {
	Data []*activityx.Activity `json:"data"`
}

func (h *Component) Handler(ctx echo.Context) error {
	path := ctx.Param("*")

	go h.CollectTrace(ctx.Request().Context(), ctx.Request().RequestURI, path)

	go h.CollectMetric(ctx.Request().Context(), ctx.Request().RequestURI, path)

	data, err := h.getActivities(ctx.Request().Context(), path, ctx.Request().URL)
	if err != nil {
		zap.L().Error("getActivities InternalError", zap.Error(err))
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, Response{
		Data: data,
	})
}
