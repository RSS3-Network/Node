package mastodon

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/activitypub"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/schema/worker/federated"
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
	return federated.Mastodon.String()
}

func (w *worker) Platform() string {
	return federated.PlatformMastodon.String()
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
	zap.L().Info("[mastodon/worker.go] reached Transform()")

	activityPubTask, ok := task.(*source.Task)

	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))

	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle ActivityPub message.
	switch activityPubTask.Message.Type {
	case "Create":
		err := w.handleActivityPubCreate(ctx, activityPubTask.Message, activity)
		if err != nil {
			return nil, fmt.Errorf("error occured in handleActivityPubCreate")
		}
		// case "Announce":
	//	w.handleAnnounce(ctx, activityPubTask.Message, activity)
	// case "Like":
	//	w.handleLike(ctx, activityPubTask.Message, activity)
	default:
		zap.L().Debug("unsupported type", zap.String("type", activityPubTask.Message.Type))
	}

	return activity, nil
}

// handleActivityPubCreate handles mastodon post message.
// Currently, it only supports creation of Note object.
// Output activities are in type of 'SocialComment' and 'SocialPost'
func (w *worker) handleActivityPubCreate(ctx context.Context, message activitypub.Object, activity *activityx.Activity) error {
	noteObject, ok := message.Object.(map[string]interface{})
	if !ok || noteObject["type"] != "Note" {
		zap.L().Debug("unsupported object type for Create", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return fmt.Errorf("invalid object type for Create activity")
	}

	// Convert the map to a ActivityPub Note struct
	var note activitypub.Note
	if err := mapToStruct(noteObject, &note); err != nil {
		return fmt.Errorf("failed to convert note object: %w", err)
	}

	// Build the Activity SocialPost object from the Note
	post := w.buildPost(ctx, message, note)
	activity.Type = typex.SocialPost
	activity.From = message.Actor

	// Check if the Note is a reply to another post
	if parentID, ok := noteObject["inReplyTo"].(string); ok {
		activity.Type = typex.SocialComment
		post.Target = &metadata.SocialPost{
			PublicationID: parentID,
		}
	}

	// Build post actions and return the result
	return w.buildPostActions(ctx, []string{message.Actor}, activity, post, activity.Type)
}

// mapToStruct converts a map to a struct using JSON marshal and unmarshal
func mapToStruct(m map[string]interface{}, v interface{}) error {
	// Marshal the map to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		return fmt.Errorf("failed to marshal map: %w", err)
	}

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(jsonData, v); err != nil {
		return fmt.Errorf("failed to unmarshal into struct: %w", err)
	}

	return nil
}

// buildPost constructs a SocialPost object from ActivityPub object and note
func (w *worker) buildPost(ctx context.Context, obj activitypub.Object, note activitypub.Note) *metadata.SocialPost {
	// Create a new SocialPost with the content, profile ID, publication ID, and timestamp
	post := &metadata.SocialPost{
		Body:          note.Content,
		ProfileID:     obj.Actor,
		PublicationID: note.ID,
		Timestamp:     w.parseTimestamp(note.Published),
	}
	// Attach media to the post
	w.buildPostMedia(ctx, post, obj.Attachment)
	// w.buildPostTags(post, obj.Tag)

	return post
}

// buildPostActions adds actions to the activity based on the post details
func (w *worker) buildPostActions(_ context.Context, actors []string, activity *activityx.Activity, post *metadata.SocialPost, socialType schema.Type) error {
	// Iterate over the actors and create an action for each
	for _, actor := range actors {
		action := activityx.Action{
			Type:     socialType,
			Platform: w.Platform(),
			From:     actor,
			To:       "",
			Metadata: *post,
		}
		activity.Actions = append(activity.Actions, &action)
	}

	return nil
}

// buildPostMedia attaches media information to the post
func (w *worker) buildPostMedia(_ context.Context, post *metadata.SocialPost, attachments []activitypub.Attachment) {
	// Iterate over the attachments and add each media in attachment object to the post
	for _, attachment := range attachments {
		post.Media = append(post.Media, metadata.Media{
			Address:  attachment.URL,
			MimeType: "", // todo: MimeType is not provided in Attachement Yet
		})
	}
}

// Iterate over the attachments and add each media to the post
func (w *worker) parseTimestamp(timestamp string) uint64 {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return 0
	}

	return uint64(t.Unix())
}

// NewWorker creates a new Mastodon worker instance
func NewWorker() (engine.Worker, error) {
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &worker{
		httpClient: httpClient,
	}, nil
}
