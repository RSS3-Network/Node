package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRSSHubHandler get rsshub data from rsshub node
func (h *Hub) GetRSSHubHandler(c echo.Context) error {
	path := c.Param("*")
	rawQuery := c.Request().URL.RawQuery

	data, err := h.getRSSHubData(c.Request().Context(), path, rawQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSONBlob(http.StatusOK, data)
}
