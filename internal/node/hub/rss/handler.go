package rss

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	"github.com/rss3-network/protocol-go/schema"
)

type Response struct {
	Data []*schema.Feed `json:"data"`
}

// GetRSSHubHandler get rsshub data from rsshub node
func (h *Hub) GetRSSHubHandler(c echo.Context) error {
	path := c.Param("*")
	rawQuery := c.Request().URL.RawQuery

	go h.countRequest(context.TODO(), path)

	data, err := h.getRSSHubData(c.Request().Context(), path, rawQuery)
	if err != nil {
		return response.InternalError(c, err)
	}

	return c.JSON(http.StatusOK, Response{
		Data: data,
	})
}
