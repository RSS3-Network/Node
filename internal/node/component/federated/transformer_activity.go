package federated

import (
	"context"

	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/tag"
	lop "github.com/samber/lo/parallel"
	"go.uber.org/zap"
)

// TransformActivity should add related URLs to the activity based on action tag, network and platform
func (c *Component) TransformActivity(ctx context.Context, activity *activityx.Activity) (*activityx.Activity, error) {
	if activity == nil {
		return nil, nil
	}

	// iterate over actions and transform them based on tag, network and platform
	lop.ForEach(activity.Actions, func(actionPtr *activityx.Action, index int) {
		action := *actionPtr

		var err error

		switch action.Tag {
		case tag.Social:
			*activity.Actions[index], err = c.TransformSocialType(ctx, activity.Network, activity.Platform, *actionPtr)
		default:
			activity.Actions[index] = actionPtr
		}

		if err != nil {
			zap.L().Error("failed to transform action", zap.Error(err), zap.String("id", activity.ID))
		}
	})

	return activity, nil
}
