package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
)

type Response struct {
	Data []*activityx.Activity `json:"data"`
}

func (h *Component) Handler(c echo.Context) error {
	path := c.Param("*")

	go h.CollectMetric(c.Request().Context(), path)

	data, err := h.getActivities(c.Request().Context(), path, c.Request().URL)
	if err != nil {
		return response.InternalError(c, err)
	}

	return c.JSON(http.StatusOK, Response{
		Data: data,
	})
}
