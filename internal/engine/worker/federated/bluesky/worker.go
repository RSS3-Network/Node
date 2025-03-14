package bluesky

import (
	"context"
	"fmt"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/reiver/go-bsky/app/bsky/feed"
	"github.com/rss3-network/node/v2/internal/database"
	"github.com/rss3-network/node/v2/internal/database/model"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/atproto"
	at "github.com/rss3-network/node/v2/provider/atproto"
	"github.com/rss3-network/node/v2/schema/worker/federated"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

const (
	ActorProfile = "app.bsky.actor.profile"
)

type worker struct {
	databaseClient database.Client
}

func (w *worker) Name() string {
	return federated.Bluesky.String()
}

func (w *worker) Platform() string {
	return federated.PlatformBluesky.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Bluesky,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Social,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.SocialComment,
		typex.SocialPost,
		typex.SocialShare,
		typex.SocialLike,
		typex.SocialProfile,
	}
}

// Filter returns the data source filter for the worker.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		Type: []string{
			feed.PostTypeValue,
			feed.RepostTypeValue,
			feed.LikeTypeValue,
			ActorProfile,
		},
	}
}

// Transform processes the task and transforms it into an activity.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	atprotoTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle atproto message.
	switch atprotoTask.Message.Collection {
	case feed.PostTypeValue:
		if atprotoTask.Message.RefMessage == nil {
			w.transformPost(ctx, atprotoTask.Message, activity)
		} else {
			w.transformComment(ctx, atprotoTask.Message, activity)
		}
	case feed.RepostTypeValue:
		if atprotoTask.Message.RefMessage != nil {
			w.transformRepost(ctx, atprotoTask.Message, activity)
		}
	case feed.LikeTypeValue:
		if atprotoTask.Message.RefMessage != nil {
			w.transformLike(ctx, atprotoTask.Message, activity)
		}
	case ActorProfile:
		w.transformProfile(ctx, atprotoTask.Message, activity)
	default:
		zap.L().Warn("unsupported type", zap.String("type", atprotoTask.Message.Collection))

		return nil, nil
	}

	if len(activity.Actions) == 0 {
		return nil, nil
	}

	w.saveProfiles(ctx, atprotoTask)

	return activity, nil
}

// transformPost transforms a post message into an activity.
func (w *worker) transformPost(_ context.Context, message at.Message, activity *activityx.Activity) {
	activity.Type = typex.SocialPost
	activity.From = message.Did.String()
	activity.To = message.Did.String()

	activity.Actions = []*activityx.Action{
		w.buildPostAction(activity.From, activity.To, activity.Type, w.buildPostMetadata(message)),
	}
}

// transformComment transforms a comment message into an activity.
func (w *worker) transformComment(_ context.Context, message at.Message, activity *activityx.Activity) {
	activity.Type = typex.SocialComment
	activity.From = message.Did.String()
	activity.To = message.RefMessage.Did.String()

	post := w.buildPostMetadata(message)

	post.Target = w.buildPostMetadata(lo.FromPtr(message.RefMessage))

	activity.Actions = []*activityx.Action{
		w.buildPostAction(activity.From, activity.To, activity.Type, post),
	}
}

// transformRepost transforms a repost message into an activity.
func (w *worker) transformRepost(_ context.Context, message at.Message, activity *activityx.Activity) {
	activity.Type = typex.SocialShare
	activity.From = message.Did.String()
	activity.To = message.RefMessage.Did.String()

	activity.Actions = []*activityx.Action{
		w.buildPostAction(activity.From, activity.To, activity.Type, w.buildPostMetadata(lo.FromPtr(message.RefMessage))),
	}
}

// transformLike transforms a like message into an activity.
func (w *worker) transformLike(_ context.Context, message at.Message, activity *activityx.Activity) {
	activity.Type = typex.SocialLike
	activity.From = message.Did.String()
	activity.To = message.RefMessage.Did.String()

	activity.Actions = []*activityx.Action{
		w.buildPostAction(activity.From, activity.To, activity.Type, w.buildPostMetadata(lo.FromPtr(message.RefMessage))),
	}
}

// transformProfile transforms a profile message into an activity.
func (w *worker) transformProfile(_ context.Context, message at.Message, activity *activityx.Activity) {
	activity.Type = typex.SocialProfile
	activity.From = message.Did.String()
	activity.To = message.Did.String()

	activity.Actions = []*activityx.Action{
		{
			Tag:      activity.Type.Tag(),
			Type:     activity.Type,
			Platform: w.Platform(),
			From:     activity.From,
			To:       activity.To,
			Metadata: w.buildProfileMetadata(message),
		},
	}
}

// buildPostMetadata constructs metadata for a post message.
func (w *worker) buildPostMetadata(message at.Message) *metadata.SocialPost {
	post := &metadata.SocialPost{
		Handle:        message.Handle,
		ProfileID:     message.Did.String(),
		PublicationID: message.Rkey,
		ContentURI:    message.URI,
		Timestamp:     uint64(message.CreatedAt.Unix()),
	}

	if message.Feed != nil {
		post.Body = message.Feed.Text
		post.Tags = message.Feed.Tags

		if message.Feed.Embed != nil {
			post.Media = w.buildPostMedia(message.Feed.Embed)
		}
	}

	return post
}

// buildProfileMetadata constructs metadata for a profile message.
func (w *worker) buildProfileMetadata(message at.Message) *metadata.SocialProfile {
	if message.Profile == nil {
		return nil
	}

	profile := &metadata.SocialProfile{
		Action: metadata.ActionSocialProfileCreate,
		Handle: message.Handle,
	}

	if message.Profile.Description != nil {
		profile.Bio = *message.Profile.Description
	}

	if message.Profile.Avatar != nil {
		profile.ImageURI = message.Profile.Avatar.Ref.String()
	}

	if message.Profile.DisplayName != nil {
		profile.Name = *message.Profile.DisplayName
	}

	return profile
}

// buildPostAction constructs an action for a post message.
func (w *worker) buildPostAction(from string, to string, typex schema.Type, post *metadata.SocialPost) *activityx.Action {
	return &activityx.Action{
		Tag:      typex.Tag(),
		Type:     typex,
		Platform: federated.PlatformBluesky.String(),
		From:     from,
		To:       to,
		Metadata: post,
	}
}

// buildPostMedia will build post media from embeds.
func (w *worker) buildPostMedia(embed *bsky.FeedPost_Embed) []metadata.Media {
	media := make([]metadata.Media, 0)

	if embed.EmbedImages != nil && embed.EmbedImages.Images != nil {
		for _, image := range embed.EmbedImages.Images {
			media = append(media, metadata.Media{
				Address:  image.Image.Ref.String(),
				MimeType: image.Image.MimeType,
			})
		}
	}

	if embed.EmbedRecord != nil && embed.EmbedRecord.Record != nil {
		media = append(media, metadata.Media{
			Address:  embed.EmbedRecord.Record.Uri,
			MimeType: "media/mp4",
		})
	}

	if embed.EmbedVideo != nil && embed.EmbedVideo.Video != nil {
		media = append(media, metadata.Media{
			Address:  embed.EmbedVideo.Video.Ref.String(),
			MimeType: embed.EmbedVideo.Video.MimeType,
		})
	}

	if embed.EmbedExternal != nil && embed.EmbedExternal.External != nil && embed.EmbedExternal.External.Thumb != nil {
		media = append(media, metadata.Media{
			Address:  embed.EmbedExternal.External.Thumb.Ref.String(),
			MimeType: embed.EmbedExternal.External.Thumb.MimeType,
		})
	}

	return media
}

// saveProfiles saves the profiles to the database.
func (w *worker) saveProfiles(ctx context.Context, task *source.Task) {
	profiles := []*model.BlueskyProfile{{
		DID:    task.Message.Did.String(),
		Handle: task.Message.Handle,
	}}

	if task.Message.RefMessage != nil {
		profiles = append(profiles, &model.BlueskyProfile{
			DID:    task.Message.RefMessage.Did.String(),
			Handle: task.Message.RefMessage.Handle,
		})
	}

	if err := w.databaseClient.SaveDatasetBlueskyProfiles(ctx, profiles); err != nil {
		zap.L().Error("save profiles", zap.Error(err))
	}
}

func NewWorker(databaseClient database.Client) (engine.Worker, error) {
	return &worker{
		databaseClient: databaseClient,
	}, nil
}
