package mastodon

import (
	"context"
	"fmt"

	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/activityPub"
	// source "github.com/rss3-network/node/internal/engine/source/activitypub"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"go.uber.org/zap"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	httpClient httpx.Client
}

func (w *worker) Name() string {
	return decentralized.Core.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformMastodon.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Mastodon,
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
	}
}

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return nil
}

func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	fmt.Println("[mastodon/worker.go] Transform reached")

	mastodonTask, ok := task.(*source.Task)

	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	taskMessage := mastodonTask.Message

	// Handle Mastodon message.
	switch mastodonTask.Message.Type {
	case "Create":
		if note, ok := mastodonTask.Message.Object.(map[string]interface{}); ok {
			if noteType, exists := note["type"].(string); exists && noteType == "Note" {
				w.handleMastodonCreateNote(ctx, taskMessage, activity)
			}
		}
	// case "Follow":
	//	w.handleMastodonFollow(ctx, taskMessage, activity)
	// case "Like":
	//	w.handleMastodonLike(ctx, taskMessage, activity)
	default:
		zap.L().Debug("unsupported type", zap.String("type", mastodonTask.Message.Type))
	}

	// if len(activity.Actions) == 0 {
	//	return nil, nil
	// }

	// Print the activity object
	fmt.Printf("[mastodon/worker.go] Transformed Activity: %+v\n", activity)

	return activity, nil
}

// function to handle Create Mastodon Note messages
func (w *worker) handleMastodonCreateNote(_ context.Context, message activitypub.Object, activity *activityx.Activity) {
	fmt.Println("[handleMastodonCreateNote] Handling Create Note")

	post := &metadata.SocialPost{
		Body:          "...",
		ProfileID:     message.ID,
		PublicationID: "...",
		Timestamp:     0,
	}

	// note := &activitypub.Note{
	//	ID:        message.ID,
	//	Type:      "Note",
	//	Content:   "default content",
	//	Published: message.Published,
	//	To:        message.To,
	//	CC:        message.CC,
	//}

	action := &activityx.Action{
		Type:     typex.SocialPost,
		Platform: decentralized.PlatformMastodon.String(),
		From:     message.Actor,
		To:       message.To[0],
		Metadata: *post,
	}
	activity.Actions = append(activity.Actions, action)
}

// func (w *worker) buildActivityAsAction(_ activitypub.Object) (metadata.Metadata, error) {
//	// switch msg.Type {
//	// case "Note":
//	//	return buildNoteMetadata(msg), nil
//	// case "Follow":
//	//	return buildFollowMetadata(msg), nil
//	// case "Like":
//	//	return buildLikeMetadata(msg), nil
//	// }
//	return nil, nil
// }

// func buildNoteMetadata(msg activitypub.Object) activitypub.Note {
//	return activitypub.Note{
//		ID:        msg.ID,
//		Type:      msg.Type,
//		Content:   msg.Attributes["content"].(string),
//		Published: msg.Published,
//		To:        msg.To,
//		CC:        msg.CC,
//	}
// }

// function to handle Follow Mastodon messages
// func (w *worker) handleMastodonFollow(_ context.Context, _ activitypub.Object, _ *activityx.Activity) {
//	// Implement your logic to handle Follow messages
//
//	fmt.Println("[handleMastodonFollow] Handling Follow")
//
//	// followActivity := &activitypub.Follow{
//	//	ID:        message.ID,
//	//	Type:      "Follow",
//	//	Actor:     message.Actor,
//	//	Target:    message.Object,
//	//	Published: message.Published,
//	//	To:        message.To,
//	//	CC:        message.CC,
//	//}
//	//
//	// action := &activityx.Action{
//	//	Type:     typex.SocialAction, // Example type, adjust based on your schema
//	//	Platform: decentralized.PlatformMastodon.String(),
//	//	From:     msg.Actor,
//	//	To:       msg.To[0], // Assuming a single recipient for simplicity
//	//	Metadata: *followActivity,
//	//}
//}

// function to handle Like Mastodon messages
// func (w *worker) handleMastodonLike(_ context.Context, _ activitypub.Object, _ *activityx.Activity) {
//	// Implement your logic to handle Like messages
//	fmt.Println("[handleMastodonLike] Handling Like")
// }

func NewWorker() (engine.Worker, error) {
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &worker{
		httpClient: httpClient,
	}, nil
}
