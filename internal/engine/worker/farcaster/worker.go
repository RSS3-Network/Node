package farcaster

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/farcaster"
	"github.com/rss3-network/node/provider/farcaster"
	"github.com/rss3-network/node/provider/httpx"
	workerx "github.com/rss3-network/node/schema/worker"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	httpClient httpx.Client
}

func (w *worker) Name() string {
	return workerx.Farcaster.String()
}

func (w *worker) Platform() string {
	return workerx.Farcaster.Platform()
}

func (w *worker) Tag() tag.Tag {
	return tag.Social
}

func (w *worker) Types() []*schema.Type {
	panic("implement me")
}

// Filter returns a source filter.
func (w *worker) Filter() engine.SourceFilter {
	return nil
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	// The worker will handle all Farcaster message.
	return task.GetNetwork().Source() == network.FarcasterSource, nil
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	farcasterTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle Farcaster message.
	switch farcasterTask.Message.Data.Type {
	case farcaster.MessageTypeCastAdd.String():
		w.handleFarcasterAddCast(ctx, farcasterTask.Message, activity)
	case farcaster.MessageTypeReactionAdd.String():
		w.handleFarcasterRecastReaction(ctx, farcasterTask.Message, activity)
	default:
		zap.L().Debug("unsupported type", zap.String("type", farcasterTask.Message.Data.Type))
	}

	if err != nil {
		return nil, fmt.Errorf("handle farcaster message failed: %w", err)
	}

	if len(activity.Actions) == 0 {
		return nil, nil
	}

	return activity, nil
}

// handleFarcasterAddCast handles farcaster add cast message.
func (w *worker) handleFarcasterAddCast(ctx context.Context, message farcaster.Message, activity *activityx.Activity) {
	fid := int64(message.Data.Fid)

	post := w.buildPost(ctx, int64(message.Data.Fid), message.Hash, message.Data.CastAddBody)

	post.Handle = message.Data.Profile.Username
	activity.From = message.Data.Profile.CustodyAddress

	// this represents a reply post.
	if message.Data.CastAddBody.ParentCastID != nil {
		activity.Type = typex.SocialComment

		targetFid := int64(message.Data.CastAddBody.ParentCastID.Fid)

		targetMessage := message.Data.CastAddBody.ParentCast

		post.Target = w.buildPost(ctx, targetFid, targetMessage.Hash, targetMessage.Data.CastAddBody)
		// this represents a reply to self.
		if fid == targetFid {
			post.Target.Handle = post.Handle

			activity.To = activity.From

			w.buildPostActions(ctx, message.Data.Profile.EthAddresses, activity, post, activity.Type)

			return
		}

		// this represents a reply to others.
		post.Target.Handle = targetMessage.Data.Profile.Username
		activity.To = targetMessage.Data.Profile.CustodyAddress

		for _, from := range message.Data.Profile.EthAddresses {
			for _, to := range targetMessage.Data.Profile.EthAddresses {
				action := activityx.Action{
					Type:     typex.SocialComment,
					Platform: w.Platform(),
					From:     from,
					To:       to,
					Metadata: *post,
				}
				activity.Actions = append(activity.Actions, &action)
			}
		}

		return
	}

	activity.Type = typex.SocialPost
	activity.To = activity.From

	w.buildPostActions(ctx, message.Data.Profile.EthAddresses, activity, post, activity.Type)
}

// handleFarcasterRecastReaction handles farcaster recast reaction message.
func (w *worker) handleFarcasterRecastReaction(ctx context.Context, message farcaster.Message, activity *activityx.Activity) {
	fid := int64(message.Data.Fid)

	post := w.buildPost(ctx, int64(message.Data.Fid), message.Hash, nil)

	post.Handle = message.Data.Profile.Username
	activity.From = message.Data.Profile.CustodyAddress

	activity.Type = typex.SocialShare

	targetFid := int64(message.Data.ReactionBody.TargetCastID.Fid)

	targetMessage := message.Data.ReactionBody.TargetCast

	post.Target = w.buildPost(ctx, targetFid, targetMessage.Hash, targetMessage.Data.CastAddBody)

	if fid == targetFid {
		post.Target.Handle = post.Handle

		activity.To = activity.From

		w.buildPostActions(ctx, message.Data.Profile.EthAddresses, activity, post, activity.Type)

		return
	}

	post.Target.Handle = targetMessage.Data.Profile.Username
	activity.To = targetMessage.Data.Profile.CustodyAddress

	for _, from := range message.Data.Profile.EthAddresses {
		for _, to := range targetMessage.Data.Profile.EthAddresses {
			action := activityx.Action{
				Type:     typex.SocialShare,
				Platform: w.Platform(),
				From:     from,
				To:       to,
				Metadata: *post,
			}
			activity.Actions = append(activity.Actions, &action)
		}
	}
}

// buildPostActions builds post actions from message.
func (w *worker) buildPostActions(_ context.Context, ethAddresses []string, activity *activityx.Activity, post *metadata.SocialPost, socialType schema.Type) {
	for _, from := range ethAddresses {
		action := activityx.Action{
			Type:     socialType,
			Platform: w.Platform(),
			From:     from,
			To:       from,
			Metadata: *post,
		}
		activity.Actions = append(activity.Actions, &action)
	}
}

// buildPost builds post from message.
func (w *worker) buildPost(ctx context.Context, fid int64, hash string, body *farcaster.CastAddBody) *metadata.SocialPost {
	var (
		text   string
		embeds []string
	)

	if body != nil {
		text = w.castToString(body)
		embeds = lo.Map(body.Embeds, func(uri farcaster.Embed, _ int) string {
			return uri.URL
		})

		embeds = append(embeds, body.EmbedsDeprecated...)
	}

	post := &metadata.SocialPost{
		Body:          text,
		ProfileID:     strconv.FormatInt(fid, 10),
		PublicationID: common.HexToAddress(hash).String(),
	}

	w.buildPostMedia(ctx, post, embeds)

	return post
}

// castToString completes the body based on the username in mentions.
func (w *worker) castToString(cast *farcaster.CastAddBody) string {
	text := cast.Text
	mentions := cast.MentionsUsernames
	mentionsPositions := cast.MentionsPositions

	var (
		textWithMentions strings.Builder
		indexBytes       int32
	)

	for i := 0; i < len(mentions); i++ {
		textWithMentions.WriteString(text[indexBytes:mentionsPositions[i]])

		username := mentions[i]

		textWithMentions.WriteString("@" + username)

		indexBytes = mentionsPositions[i]
	}

	textWithMentions.WriteString(text[indexBytes:])

	return textWithMentions.String()
}

// buildPostMedia will build post media from embeds.
func (w *worker) buildPostMedia(ctx context.Context, post *metadata.SocialPost, embeds []string) {
	var locker sync.Mutex

	errorGroup, _ := errgroup.WithContext(ctx)

	for _, embed := range embeds {
		embed := embed

		errorGroup.Go(func() error {
			mimeType, _ := w.httpClient.GetContentType(ctx, embed)

			locker.Lock()
			defer locker.Unlock()

			if mimeType != "" {
				post.Media = append(post.Media, metadata.Media{
					Address: embed, MimeType: mimeType,
				})
			}

			return nil
		})
	}

	_ = errorGroup.Wait()
}

func NewWorker() (engine.Worker, error) {
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &worker{
		httpClient: httpClient,
	}, nil
}
