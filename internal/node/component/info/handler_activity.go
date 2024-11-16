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

	zap.L().Debug("loading checkpoints from database")

	checkpoints, err := c.databaseClient.LoadCheckpoints(ctx, "", networkx.Unknown, "")
	if err != nil {
		return count, nil, fmt.Errorf("failed to find activity count: %w", err)
	}

	zap.L().Debug("processing checkpoints to calculate total activity count")

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
		zap.L().Debug("database client is not initialized, returning zero count")

		return ctx.JSON(http.StatusOK, StatisticResponse{
			Count: 0,
		})
	}

	zap.L().Debug("getting activity count from database")

	count, updateTime, err := c.getActivityCountFromDB(ctx.Request().Context())
	if err != nil {
		zap.L().Error("failed to get activity count from database",
			zap.Error(err))

		return response.InternalError(ctx)
	}

	zap.L().Debug("successfully retrieved activity count",
		zap.Int64("count", count),
		zap.Time("lastUpdate", *updateTime))

	return ctx.JSON(http.StatusOK, StatisticResponse{
		Count:      count,
		LastUpdate: updateTime,
	})
}
