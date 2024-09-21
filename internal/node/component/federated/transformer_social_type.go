package decentralized

import (
	"context"
	"fmt"
	"strings"

	"github.com/rss3-network/node/schema/worker/decentralized"
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
func (c *Component) TransformSocialPost(ctx context.Context, network network.Network, platform string, action activity.Action) (activity.Action, error) {
	post, ok := action.Metadata.(*metadata.SocialPost)
	if !ok {
		zap.L().Error("invalid post metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid post metadata type: %T", action.Metadata)
	}

	if lo.IsEmpty(post.Handle) {
		post.Handle = action.From
	}

	post.AuthorURL = c.buildSocialAuthorURL(ctx, platform, post.Handle)

	if post.Target != nil && platform != "" && platform != decentralized.PlatformMirror.String() {
		post.Target.AuthorURL = c.buildSocialAuthorURL(ctx, platform, post.Target.Handle)
		post.TargetURL = c.buildSocialNoteURL(ctx, network, platform, lo.ToPtr(post.Target.Handle), lo.ToPtr(post.Target.ProfileID), lo.ToPtr(post.Target.PublicationID))
	}

	action.Metadata = post

	if lo.IsNotEmpty(platform) && platform == decentralized.PlatformMirror.String() {
		originID := post.PublicationID
		if post.Target != nil {
			originID = post.Target.PublicationID
		}

		if noteURL := c.buildSocialNoteURL(ctx, network, platform, lo.ToPtr(post.Handle), lo.ToPtr(post.ProfileID), lo.ToPtr(originID)); noteURL != "" {
			action.RelatedURLs = append(action.RelatedURLs, noteURL)
		}

		return action, nil
	}

	if noteURL := c.buildSocialNoteURL(ctx, network, platform, lo.ToPtr(post.Handle), lo.ToPtr(post.ProfileID), lo.ToPtr(post.PublicationID)); noteURL != "" {
		action.RelatedURLs = append(action.RelatedURLs, noteURL)
	}

	return action, nil
}

// TransformSocialProfile adds author url to social profile action
func (c *Component) TransformSocialProfile(ctx context.Context, platform string, action activity.Action) (activity.Action, error) {
	profile, ok := action.Metadata.(*metadata.SocialProfile)
	if !ok {
		zap.L().Error("invalid profile metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid profile metadata type: %T", action.Metadata)
	}

	action.RelatedURLs = append(action.RelatedURLs, c.buildSocialAuthorURL(ctx, platform, profile.Handle))

	return action, nil
}

// TransformSocialProxy adds author url to social proxy action
func (c *Component) TransformSocialProxy(ctx context.Context, platform string, action activity.Action) (activity.Action, error) {
	proxy, ok := action.Metadata.(*metadata.SocialProxy)
	if !ok {
		zap.L().Error("invalid proxy metadata type", zap.Any("metadata", action.Metadata))

		return activity.Action{}, fmt.Errorf("invalid proxy metadata type: %T", action.Metadata)
	}

	action.RelatedURLs = append(action.RelatedURLs, c.buildSocialAuthorURL(ctx, platform, proxy.Profile.Handle))

	return action, nil
}

// buildSocialAuthorURL returns author url based on platform and handle
func (c *Component) buildSocialAuthorURL(_ context.Context, platform string, handle string) string {
	if lo.IsEmpty(handle) || lo.IsEmpty(platform) {
		return ""
	}

	switch platform {
	case decentralized.PlatformCrossbell.String():
		return fmt.Sprintf("https://crossbell.io/@%s", strings.TrimSuffix(handle, ".csb"))
	case decentralized.PlatformLens.String():
		return fmt.Sprintf("https://hey.xyz/u/%s", strings.TrimSuffix(handle, ".lens"))
	case decentralized.PlatformMirror.String():
		return fmt.Sprintf("https://mirror.xyz/%s", handle)
	case decentralized.PlatformENS.String():
		return fmt.Sprintf("https://app.ens.domains/name/%s", handle)
	case decentralized.PlatformIQWiki.String():
		return fmt.Sprintf("https://iq.wiki/account/%s", handle)
	case decentralized.PlatformFarcaster.String():
		return fmt.Sprintf("https://warpcast.com/%s", handle)
	case decentralized.PlatformParagraph.String():
		return fmt.Sprintf("https://paragraph.xyz/@%s", handle)
	default:
		return ""
	}
}

// buildSocialNoteURL returns note url based on platform, handle, profileID and pubID
func (c *Component) buildSocialNoteURL(_ context.Context, n network.Network, platform string, handle, profileID, pubID *string) string {
	if platform == "" {
		return ""
	}

	switch platform {
	case decentralized.PlatformCrossbell.String():
		return lo.Ternary(lo.IsEmpty(profileID) || lo.IsEmpty(pubID), "", fmt.Sprintf("https://crossbell.io/notes/%s-%s", lo.FromPtr(profileID), lo.FromPtr(pubID)))
	case decentralized.PlatformLens.String():
		if n == network.Polygon {
			return lo.Ternary(lo.IsEmpty(profileID) || lo.IsEmpty(pubID), "", fmt.Sprintf("https://hey.xyz/posts/%s-%s", lo.FromPtr(profileID), lo.FromPtr(pubID)))
		}

		if n == network.Arweave {
			return lo.Ternary(lo.IsEmpty(pubID), "", fmt.Sprintf("https://hey.xyz/posts/%s", lo.FromPtr(pubID)))
		}
	case decentralized.PlatformMirror.String():
		return lo.Ternary(lo.IsEmpty(handle) || lo.IsEmpty(pubID), "", fmt.Sprintf("https://mirror.xyz/%s/%s", lo.FromPtr(handle), lo.FromPtr(pubID)))
	case decentralized.PlatformIQWiki.String():
		return lo.Ternary(lo.IsEmpty(pubID), "", fmt.Sprintf("https://iq.wiki/wiki/%s", lo.FromPtr(pubID)))
	case decentralized.PlatformFarcaster.String():
		return lo.Ternary(lo.IsEmpty(pubID) || lo.IsEmpty(handle), "", fmt.Sprintf("https://warpcast.com/%s/%s", lo.FromPtr(handle), lo.FromPtr(pubID)))
	case decentralized.PlatformParagraph.String():
		return lo.Ternary(lo.IsEmpty(pubID) || lo.IsEmpty(handle), "", fmt.Sprintf("https://paragraph.xyz/@%s/%s", lo.FromPtr(handle), lo.FromPtr(pubID)))
	}

	return ""
}
