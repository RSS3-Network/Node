package mastodon

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/activitypub"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
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
	return decentralized.Mastodon.String()
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
	if !ok || noteObject[mastodon.Type] != mastodon.ObjectTypeNote {
		zap.L().Info("unsupported object type for Create", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return fmt.Errorf("invalid object type for Create activity")
	}

	// Convert the map to an ActivityPub Note struct
	var note activitypub.Note
	if err := mapToStruct(noteObject, &note); err != nil {
		return fmt.Errorf("failed to convert note object: %w", err)
	}

	timestamp := activity.Timestamp
	// Build the Activity SocialPost object from the Note
	post := w.buildPost(ctx, message, note, timestamp)
	activity.Type = typex.SocialPost

	currentUserHandle := convertURLToHandle(message.Actor)

	// If the actor is from a relay server then we directly set it as the handle
	if currentUserHandle == "" {
		currentUserHandle = message.Actor
	}

	activity.From = currentUserHandle

	activity.Platform = w.Platform()

	toUserHandle := currentUserHandle

	// Check if the Note is a reply to another post -  activity SocialComment object
	if parentID, ok := noteObject[mastodon.InReplyTo].(string); ok {
		activity.Type = typex.SocialComment
		post.Target = &metadata.SocialPost{
			PublicationID: parentID,
		}

		toUserHandle = convertURLToHandle(parentID)
	}

	activity.To = toUserHandle

	// Generate main action
	mainAction := w.createAction(activity.Type, currentUserHandle, toUserHandle, post)
	activity.Actions = append(activity.Actions, mainAction)

	// Generate additional actions for mentions
	mentionActions := w.createMentionActions(currentUserHandle, note, post)
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

	currentUserHandle := convertURLToHandle(message.Actor)

	// If the actor is from a relay server then we directly set it as the handle
	if currentUserHandle == "" {
		currentUserHandle = message.Actor
	}

	activity.From = currentUserHandle
	activity.To = currentUserHandle

	// Extract object IDs from the message
	objectIDs, err := extractObjectIDs(message.Object)
	if err != nil {
		zap.L().Info("unsupported object type for Announce", zap.String("type", fmt.Sprintf("%T", message.Object)))
		return err
	}

	// Iteratively create action for every announcement of the activity
	for _, announcedID := range objectIDs {
		toUserHandle := convertURLToHandle(announcedID)

		// Create a SocialPost object with the Announced ID
		post := &metadata.SocialPost{
			Handle:        toUserHandle,
			ProfileID:     message.Actor,
			PublicationID: message.ID,
			Timestamp:     w.parseTimestamp(message.Published),
			Target: &metadata.SocialPost{
				PublicationID: announcedID,
			},
		}

		// Create and add action to activity
		action := w.createAction(activity.Type, currentUserHandle, toUserHandle, post)
		activity.Actions = append(activity.Actions, action)
	}

	return nil
}

// handleActivityPubLike handles Like activities in ActivityPub.
func (w *worker) handleActivityPubLike(_ context.Context, message activitypub.Object, activity *activityx.Activity) error {
	activity.Type = typex.SocialComment

	currentUserHandle := convertURLToHandle(message.Actor)

	// If the actor is from a relay server then we directly set it as the handle
	if currentUserHandle == "" {
		currentUserHandle = message.Actor
	}

	activity.From = currentUserHandle

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
func (w *worker) createMentionActions(from string, note activitypub.Note, post *metadata.SocialPost) []*activityx.Action {
	var actions []*activityx.Action

	actionType := typex.SocialShare
	// Make mention actions for every tag in the activity
	for _, mention := range note.Tag {
		if mention.Type == mastodon.TagTypeMention {
			mentionUserHandle := convertURLToHandle(mention.Href)
			mentionAction := w.createAction(actionType, from, mentionUserHandle, post)
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
func (w *worker) buildPost(ctx context.Context, obj activitypub.Object, note activitypub.Note, timestamp uint64) *metadata.SocialPost {
	// Create a new SocialPost with the content, profile ID, publication ID, and timestamp
	post := &metadata.SocialPost{
		Body:          note.Content,
		ProfileID:     obj.Actor,
		PublicationID: note.ID,
		Timestamp:     timestamp,
		Handle:        convertURLToHandle(obj.Actor),
	}
	// Attach media to the post
	w.buildPostMedia(ctx, post, obj.Attachment)
	w.buildPostTags(post, note.Tag)

	return post
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
		if tag.Type == mastodon.TagTypeHashtag {
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

// convertURLToHandle converts a Mastodon URL into a user handle in the format '@username@domain'.
func convertURLToHandle(statusID string) string {
	parsedURL, err := url.Parse(statusID)
	if err != nil {
		return ""
	}

	// Handle ActivityPub actor URLs, in particular the URLs related to Relay servers
	if strings.HasSuffix(parsedURL.Path, "/actor") {
		username := "relay"
		domain := parsedURL.Host

		return fmt.Sprintf("@%s@%s", username, domain)
	}

	// Check if the path contains "@username" (e.g., "/@username/status")
	if strings.HasPrefix(parsedURL.Path, "/@") {
		parts := strings.Split(parsedURL.Path, "/")
		if len(parts) < 2 {
			return ""
		}

		username := strings.TrimPrefix(parts[1], "@")
		domain := parsedURL.Host

		return fmt.Sprintf("@%s@%s", username, domain)
	}

	// Fallback for other URL formats (like status URLs)
	parts := strings.Split(parsedURL.Path, "/")
	if len(parts) < 3 {
		return ""
	}

	username := parts[2]

	domain := parsedURL.Host

	return fmt.Sprintf("@%s@%s", username, domain)
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
