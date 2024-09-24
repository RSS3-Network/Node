package federated

import (
	"context"
	"fmt"

	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// TransformSocialType adds author url and note url to social actions based on type, network and platform
func (c *Component) TransformSocialType(ctx context.Context, network network.Network, platform string, action activity.Action) (activity.Action, error) {
	switch action.Type {
	case typex.SocialPost, typex.SocialComment, typex.SocialShare, typex.SocialRevise, typex.SocialMint, typex.SocialDelete:
		return c.TransformSocialPost(ctx, network, platform, action)
	case typex.SocialProfile:
		return c.TransformSocialProfile(ctx, platform, action)
	case typex.SocialProxy:
		return c.TransformSocialProxy(ctx, platform, action)
	}

	return action, nil
}

// TransformSocialPost adds author url and note url to social post action
func (c *Component) TransformSocialPost(_ context.Context, _ network.Network, _ string, action activity.Action) (activity.Action, error) {
	post, ok := action.Metadata.(*metadata.SocialPost)
	if !ok {
		zap.L().Error("invalid post metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid post metadata type: %T", action.Metadata)
	}

	if lo.IsEmpty(post.Handle) {
		post.Handle = action.From
	}

	action.Metadata = post

	return action, nil
}

// TransformSocialProfile adds author url to social profile action
func (c *Component) TransformSocialProfile(_ context.Context, _ string, action activity.Action) (activity.Action, error) {
	_, ok := action.Metadata.(*metadata.SocialProfile)
	if !ok {
		zap.L().Error("invalid profile metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid profile metadata type: %T", action.Metadata)
	}

	return action, nil
}

// TransformSocialProxy adds author url to social proxy action
func (c *Component) TransformSocialProxy(_ context.Context, _ string, action activity.Action) (activity.Action, error) {
	_, ok := action.Metadata.(*metadata.SocialProxy)
	if !ok {
		zap.L().Error("invalid proxy metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid proxy metadata type: %T", action.Metadata)
	}

	return action, nil
}
