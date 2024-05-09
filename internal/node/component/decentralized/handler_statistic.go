package decentralized

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
)

type StatisticResponse struct {
	Count      int64      `json:"count"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
}

func (c *Component) GetActivitiesCount(ctx echo.Context) error {
	count, updateTime, err := c.getIndexCount(ctx.Request().Context())
	if err != nil {
		return response.InternalError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, StatisticResponse{
		Count:      count,
		LastUpdate: updateTime,
	})
}
