package decentralized

import (
	"context"
	"fmt"
	"strings"

	"github.com/rss3-network/node/internal/database/model"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

func (c *Component) getActivity(ctx context.Context, request model.ActivityQuery) (*activityx.Activity, *int, error) {
	return c.databaseClient.FindActivity(ctx, request)
}

func (c *Component) getActivities(ctx context.Context, request model.ActivitiesQuery) ([]*activityx.Activity, string, error) {
	activities, err := c.databaseClient.FindActivities(ctx, request)
	if err != nil {
		return nil, "", fmt.Errorf("failed to find activities: %w", err)
	}

	last, exist := lo.Last(activities)
	if exist {
		return activities, c.transformCursor(ctx, last), nil
	}

	return nil, "", nil
}

func (c *Component) getCursor(ctx context.Context, cursor *string) (*activityx.Activity, error) {
	if cursor == nil {
		return nil, nil
	}

	str := strings.Split(*cursor, ":")
	if len(str) != 2 {
		return nil, fmt.Errorf("invalid cursor")
	}

	network, err := networkx.NetworkString(str[1])
	if err != nil {
		return nil, fmt.Errorf("invalid cursor: %w", err)
	}

	data, _, err := c.getActivity(ctx, model.ActivityQuery{ID: lo.ToPtr(str[0]), Network: lo.ToPtr(network)})
	if err != nil {
		return nil, fmt.Errorf("failed to get cursor: %w", err)
	}

	return data, nil
}

func (c *Component) transformCursor(_ context.Context, activity *activityx.Activity) string {
	if activity == nil {
		return ""
	}

	return fmt.Sprintf("%s:%s", activity.ID, activity.Network)
}
