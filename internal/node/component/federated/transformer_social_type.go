package federated

import (
	"context"
	"fmt"
	"strings"

	"github.com/rss3-network/node/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
)

// TransformSocialType adds author url and note url to social actions based on type, network and platform
func (c *Component) TransformSocialType(ctx context.Context, network network.Network, platform string, action activity.Action) (activity.Action, error) {
	switch action.Type {
	case typex.SocialPost, typex.SocialComment, typex.SocialShare, typex.SocialLike:
		return c.TransformSocialPost(ctx, network, platform, action)
	}

	return action, nil
}

// TransformSocialPost adds author url and note url to social post action
func (c *Component) TransformSocialPost(ctx context.Context, network network.Network, platform string, action activity.Action) (activity.Action, error) {
	post, ok := action.Metadata.(*metadata.SocialPost)
	if !ok {
		return activity.Action{}, fmt.Errorf("invalid post metadata type: %T", action.Metadata)
	}

	if lo.IsEmpty(post.Handle) {
		post.Handle = action.From
	}

	post.AuthorURL = c.buildSocialAuthorURL(ctx, platform, post.Handle)

	if post.Target != nil && platform != "" {
		post.Target.AuthorURL = c.buildSocialAuthorURL(ctx, platform, post.Target.Handle)
		post.TargetURL = c.buildSocialNoteURL(ctx, network, platform, post.Target.Handle, post.Target.PublicationID)
	}

	action.Metadata = post

	if noteURL := c.buildSocialNoteURL(ctx, network, platform, post.Handle, post.PublicationID); noteURL != "" {
		action.RelatedURLs = append(action.RelatedURLs, noteURL)
	}

	return action, nil
}

// buildSocialAuthorURL returns author url based on domain and handle
func (c *Component) buildSocialAuthorURL(_ context.Context, platform, handle string) string {
	if lo.IsEmpty(handle) || lo.IsEmpty(platform) {
		return ""
	}

	switch platform {
	case federated.PlatformMastodon.String():
		parts := strings.SplitN(handle, "@", 3)
		if len(parts) != 3 {
			return ""
		}

		username, domain := parts[1], parts[2]

		return fmt.Sprintf("https://%s/users/%s", domain, username)
	case federated.PlatformBluesky.String():
		return fmt.Sprintf("https://bsky.app/profile/%s", handle)
	default:
		return ""
	}
}

// buildSocialNoteURL returns note url based on domain, handle, profileID and pubID
func (c *Component) buildSocialNoteURL(_ context.Context, _ network.Network, platform, handle, pubID string) string {
	if lo.IsEmpty(platform) || lo.IsEmpty(handle) || lo.IsEmpty(pubID) {
		return ""
	}

	switch platform {
	case federated.PlatformMastodon.String():
		parts := strings.SplitN(handle, "@", 3)
		if len(parts) != 3 {
			return ""
		}

		username, domain := parts[1], parts[2]

		return fmt.Sprintf("https://%s/users/%s/statuses/%s", domain, username, pubID)
	case federated.PlatformBluesky.String():
		return fmt.Sprintf("https://bsky.app/profile/%s/post/%s", handle, pubID)
	default:
		return ""
	}
}
