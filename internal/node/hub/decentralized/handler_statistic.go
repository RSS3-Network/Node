package decentralized

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naturalselectionlabs/rss3-node/common/http/response"
)

func (h *Hub) GetActivitiesCount(c echo.Context) error {
	count, updateTime, err := h.getIndexCount(c.Request().Context())
	if err != nil {
		return response.InternalError(c, err)
	}

	return c.JSON(http.StatusOK, StatisticResponse{
		Count:      count,
		LastUpdate: updateTime,
	})
}
