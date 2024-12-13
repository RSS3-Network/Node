package federated

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"

	"github.com/rss3-network/node/internal/database/model"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	networkx "github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
)

func (c *Component) getActivity(ctx context.Context, request model.ActivityQuery) (*activityx.Activity, *int, error) {
	if request.Owner != nil {
		owner := c.transformHandler(ctx, []string{lo.FromPtr(request.Owner)})

		if len(owner) == 0 {
			return nil, nil, nil
		}

		request.Owner = lo.ToPtr(owner[0])
	}

	return c.databaseClient.FindActivity(ctx, request)
}

func (c *Component) getActivities(ctx context.Context, request model.ActivitiesQuery) ([]*activityx.Activity, string, error) {
	if request.Owner != nil {
		owner := c.transformHandler(ctx, []string{lo.FromPtr(request.Owner)})

		if len(owner) == 0 {
			return nil, "", nil
		}

		request.Owner = lo.ToPtr(owner[0])
	}

	if len(request.Owners) > 0 {
		request.Owners = c.transformHandler(ctx, request.Owners)

		if len(request.Owners) == 0 {
			return nil, "", nil
		}
	}

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

	cleanedCursor := *cursor
	prefix := "https://"

	if strings.HasPrefix(cleanedCursor, "https://") {
		cleanedCursor = cleanedCursor[8:]
	} else if strings.HasPrefix(cleanedCursor, "http://") {
		cleanedCursor = cleanedCursor[7:]
		prefix = "http://"
	}

	id, networkStr, found := strings.Cut(cleanedCursor, ":")
	if !found {
		return nil, fmt.Errorf("invalid cursor: missing network")
	}

	network, err := networkx.NetworkString(networkStr)
	if err != nil {
		return nil, fmt.Errorf("invalid cursor: %w", err)
	}

	data, _, err := c.getActivity(ctx, model.ActivityQuery{ID: lo.ToPtr(prefix + id), Network: lo.ToPtr(network)})
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

func (c *Component) transformHandler(ctx context.Context, owners []string) []string {
	var (
		blueskyHandlers = make([]string, 0, len(owners))
		results         = make([]string, 0, len(owners))
	)

	for _, owner := range owners {
		if strings.Contains(owner, ".bsky.social") {
			blueskyHandlers = append(blueskyHandlers, owner)
		} else {
			results = append(results, owner)
		}
	}

	blueskyProfiles, err := c.loadBlueskyProfiles(ctx, model.QueryBlueskyProfiles{Handles: blueskyHandlers})
	if err != nil {
		zap.L().Error("failed to load bluesky profiles", zap.Error(err))
	}

	if len(blueskyProfiles) > 0 {
		for _, profile := range blueskyProfiles {
			results = append(results, profile.DID)
		}
	}

	return results
}

// loadMastodonHandles retrieves the updated Mastodon handles from the database.
func (c *Component) loadMastodonHandles(ctx context.Context, query model.QueryMastodonHandles) ([]*model.MastodonHandle, error) {
	return c.databaseClient.GetUpdatedMastodonHandles(ctx, query)
}

// loadBlueskyProfiles retrieves the Bluesky profiles from the database.
func (c *Component) loadBlueskyProfiles(ctx context.Context, query model.QueryBlueskyProfiles) ([]*model.BlueskyProfile, error) {
	return c.databaseClient.LoadDatasetBlueskyProfiles(ctx, query)
}
