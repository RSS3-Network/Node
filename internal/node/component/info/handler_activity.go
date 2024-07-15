package info

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rss3-network/node/common/http/response"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type StatisticResponse struct {
	Count      int64      `json:"count"`
	LastUpdate *time.Time `json:"last_update,omitempty"`
}

// getActivityCountFromDB returns the total number of activities stored in the database.
func (c *Component) getActivityCountFromDB(ctx context.Context) (int64, *time.Time, error) {
	var (
		count      int64
		updateTime *time.Time
	)

	checkpoints, err := c.databaseClient.LoadCheckpoints(ctx, "", networkx.Unknown, "")
	if err != nil {
		return count, nil, fmt.Errorf("failed to find activity count: %w", err)
	}

	for _, checkpoint := range checkpoints {
		count += checkpoint.IndexCount

		if updateTime == nil || checkpoint.UpdatedAt.After(*updateTime) {
			updateTime = lo.ToPtr(checkpoint.UpdatedAt)
		}
	}

	return count, updateTime, nil
}

// GetActivityCount returns the total number of activities indexed by this Node.
func (c *Component) GetActivityCount(ctx echo.Context) error {
	if c.databaseClient == nil {
		return ctx.JSON(http.StatusOK, StatisticResponse{
			Count: 0,
		})
	}

	count, updateTime, err := c.getActivityCountFromDB(ctx.Request().Context())
	if err != nil {
		zap.L().Error("getActivityCountFromDB InternalError", zap.Error(err))
		return response.InternalError(ctx)
	}

	return ctx.JSON(http.StatusOK, StatisticResponse{
		Count:      count,
		LastUpdate: updateTime,
	})
}
