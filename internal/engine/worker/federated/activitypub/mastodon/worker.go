package mastodon

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

// Transform processes the task and converts it into an activity.
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
	case mastodon.MessageTypeCreate.String():
		err := w.handleActivityPubCreate(ctx, activityPubTask.Message, activity)
		if err != nil {
			return nil, fmt.Errorf("error occurred in handleActivityPubCreate: %w", err)
		}
	case mastodon.MessageTypeAnnounce.String():
		err := w.handleActivityPubAnnounce(ctx, activityPubTask.Message, activity)
		if err != nil {
			return nil, fmt.Errorf("error occurred in handleActivityPubAnnounce: %w", err)
		}
	case mastodon.MessageTypeLike.String():
		err := w.handleActivityPubLike(ctx, activityPubTask.Message, activity)
		if err != nil {
			return nil, fmt.Errorf("error occurred in handleActivityPubLike: %w", err)
		}
	default:
		zap.L().Info("unsupported type", zap.String("type", activityPubTask.Message.Type))
	}

	zap.L().Info("unsupported type", zap.String("type", activityPubTask.Message.Type))

	return activity, nil
}

// handleActivityPubCreate handles mastodon post message.
func (w *worker) handleActivityPubCreate(ctx context.Context, message activitypub.Object, activity *activityx.Activity) error {
	noteObject, ok := message.Object.(map[string]interface{})
	if !ok || noteObject["type"] != "Note" {
		zap.L().Info("unsupported object type for Create", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return fmt.Errorf("invalid object type for Create activity")
	}

	// Convert the map to an ActivityPub Note struct
	var note activitypub.Note
	if err := mapToStruct(noteObject, &note); err != nil {
		return fmt.Errorf("failed to convert note object: %w", err)
	}

	// Build the Activity SocialPost object from the Note
	post := w.buildPost(ctx, message, note)
	activity.Type = typex.SocialPost
	activity.From = message.Actor
	activity.Platform = w.Platform()

	// Check if the Note is a reply to another post
	// If true, then make it an activity SocialComment object
	if parentID, ok := noteObject["inReplyTo"].(string); ok {
		activity.Type = typex.SocialComment
		post.Target = &metadata.SocialPost{
			PublicationID: parentID,
		}
	}

	// Determine the main recipient of this Post
	// recipient := ""
	// if len(note.To) > 0 {
	//	recipient = note.To[0]
	// }

	// Generate main action
	mainAction := w.createAction(activity.Type, message.Actor, "", post)
	activity.Actions = append(activity.Actions, mainAction)

	// Generate additional actions for mentions
	mentionActions := w.createMentionActions(activity.Type, message.Actor, note, post)
	activity.Actions = append(activity.Actions, mentionActions...)

	return nil
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

// handleActivityPubAnnounce handles Announce activities (shares/boosts) in ActivityPub.
func (w *worker) handleActivityPubAnnounce(_ context.Context, message activitypub.Object, activity *activityx.Activity) error {
	activity.Type = typex.SocialShare
	activity.From = message.Actor

	// Extract object IDs from the message
	objectIDs, err := extractObjectIDs(message.Object)
	if err != nil {
		zap.L().Info("unsupported object type for Announce", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return err
	}

	// Iteratively create action for every announcement of the activity
	for _, announcedID := range objectIDs {
		// Create a SocialPost object with the Announced ID
		post := &metadata.SocialPost{
			ProfileID:     message.Actor,
			PublicationID: message.ID,
			Timestamp:     w.parseTimestamp(message.Published),
			Target: &metadata.SocialPost{
				PublicationID: announcedID,
			},
		}

		// Create and add action to activity
		action := w.createAction(activity.Type, message.Actor, "", post)
		activity.Actions = append(activity.Actions, action)
	}

	return nil
}

// handleActivityPubLike handles Like activities in ActivityPub.
func (w *worker) handleActivityPubLike(_ context.Context, message activitypub.Object, activity *activityx.Activity) error {
	activity.Type = typex.SocialComment
	activity.From = message.Actor

	// Extract object IDs from the message
	objectIDs, err := extractObjectIDs(message.Object)
	if err != nil {
		zap.L().Info("unsupported object type for Like", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return err
	}

	for _, likedID := range objectIDs {
		// Create a SocialPost object with the Liked ID
		post := &metadata.SocialPost{
			ProfileID:     message.Actor,
			PublicationID: message.ID,
			Timestamp:     w.parseTimestamp(message.Published),
			Target: &metadata.SocialPost{
				PublicationID: likedID,
			},
		}

		// Create and add action to activity
		action := w.createAction(activity.Type, message.Actor, "", post)
		activity.Actions = append(activity.Actions, action)
	}

	return nil
}

// createMentionActions generates actions for mentions within a note.
func (w *worker) createMentionActions(actionType schema.Type, from string, note activitypub.Note, post *metadata.SocialPost) []*activityx.Action {
	var actions []*activityx.Action

	// Make mention actions for every tag in the activity
	for _, mention := range note.Tag {
		if mention.Type == "Mention" {
			mentionAction := w.createAction(actionType, from, mention.Href, post)
			actions = append(actions, mentionAction)
		}
	}

	return actions
}

// createAction creates an activity action.
func (w *worker) createAction(actionType schema.Type, from, to string, metadata metadata.Metadata) *activityx.Action {
	return &activityx.Action{
		Type:     actionType,
		Platform: w.Platform(),
		From:     from,
		To:       to,
		Metadata: metadata,
	}
}

// buildPost constructs a SocialPost object from ActivityPub object and note
func (w *worker) buildPost(ctx context.Context, obj activitypub.Object, note activitypub.Note) *metadata.SocialPost {
	// Create a new SocialPost with the content, profile ID, publication ID, and timestamp
	post := &metadata.SocialPost{
		Body:          note.Content,
		ProfileID:     obj.Actor,
		PublicationID: note.ID,
		Timestamp:     w.parseTimestamp(note.Published),
		Handle:        w.extractHandle(obj.Actor),
	}
	// Attach media to the post
	w.buildPostMedia(ctx, post, obj.Attachment)
	w.buildPostTags(post, note.Tag)

	return post
}

// extractHandle parses the username out of the actor string
func (w *worker) extractHandle(actor string) string {
	// Extract the last part of the URL after the final slash
	parts := strings.Split(actor, "/")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}

	return actor
}

// buildPostMedia attaches media information to the post
func (w *worker) buildPostMedia(_ context.Context, post *metadata.SocialPost, attachments []activitypub.Attachment) {
	// Iterate over the attachments and add each media in attachment object to the post
	for _, attachment := range attachments {
		post.Media = append(post.Media, metadata.Media{
			Address:  attachment.URL,
			MimeType: attachment.Type,
		})
	}
}

// buildPostTags builds the Tags field in the metatdata
func (w *worker) buildPostTags(post *metadata.SocialPost, tags []activitypub.Tag) {
	for _, tag := range tags {
		if tag.Type == "Hashtag" {
			post.Tags = append(post.Tags, tag.Name)
		}
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

// extractObjectIDs used to extract Object IDs for Announce and Like ActivityPub object
func extractObjectIDs(object interface{}) ([]string, error) {
	var ids []string

	switch obj := object.(type) {
	case string:
		ids = append(ids, obj)
	case map[string]interface{}:
		var nestedObject activitypub.Object
		if err := mapToStruct(obj, &nestedObject); err != nil {
			return nil, fmt.Errorf("failed to convert nested object: %w", err)
		}

		ids = append(ids, nestedObject.ID)
	case []interface{}:
		for _, item := range obj {
			switch item := item.(type) {
			case string:
				ids = append(ids, item)
			case map[string]interface{}:
				var nestedObject activitypub.Object
				if err := mapToStruct(item, &nestedObject); err != nil {
					return nil, fmt.Errorf("failed to convert nested object: %w", err)
				}

				ids = append(ids, nestedObject.ID)
			default:
				return nil, fmt.Errorf("unsupported object type in array: %T", item)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported object type: %T", obj)
	}

	return ids, nil
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
