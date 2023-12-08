package rss

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRSSHubHandler get rsshub data from rsshub node
func (r *RSS) GetRSSHubHandler(c echo.Context) error {
	path := c.Param("*")
	rawQuery := c.Request().URL.RawQuery

	data, err := r.getRSSHubData(c.Request().Context(), path, rawQuery)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONBlob(http.StatusOK, data)
}
