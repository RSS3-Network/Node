package mastodon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/jaytaylor/html2text"
	"github.com/redis/rueidis"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/activitypub"
	"github.com/rss3-network/node/provider/activitypub"
	"github.com/rss3-network/node/provider/activitypub/mastodon"
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
	httpClient     httpx.Client
	databaseClient database.Client
	redisClient    rueidis.Client
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
		err = w.handleActivityPubCreate(ctx, activityPubTask.Message, activity)
	case mastodon.MessageTypeAnnounce.String():
		err = w.handleActivityPubAnnounce(ctx, activityPubTask.Message, activity)
	default:
		zap.L().Warn("unsupported type", zap.String("type", activityPubTask.Message.Type))
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("handle %s message: %w", activityPubTask.Message.Type, err)
	}

	return activity, nil
}

func (w *worker) handleActivityPubCreate(ctx context.Context, message activitypub.Object, activity *activityx.Activity) error {
	noteObjects, err := extractNoteObjects(message.Object)
	if err != nil {
		return fmt.Errorf("failed to extract Create objects: %w", err)
	}

	currentUserHandle := convertURLToHandle(message.ID)

	for _, currentNoteObject := range noteObjects {
		if err := w.handleSingleActivityPubCreate(ctx, message, currentNoteObject, activity, currentUserHandle); err != nil {
			return err
		}
	}

	return nil
}

// handleActivityPubCreate handles mastodon post message.
func (w *worker) handleSingleActivityPubCreate(ctx context.Context, message activitypub.Object, note activitypub.Note, activity *activityx.Activity, currentUserHandle string) error {
	noteObject, _ := message.Object.(map[string]interface{})
	timestamp := activity.Timestamp
	// Build the Activity SocialPost object from the Note
	post := w.buildPost(ctx, message, note, timestamp)

	// Set activity Type and Tag
	activity.Type = typex.SocialPost
	activity.Tag = tag.Social
	activity.Platform = w.Platform()

	// If the actor is from a relay server then we directly set it as the handle value
	if currentUserHandle == "" {
		currentUserHandle = message.Actor
	}

	activity.From = currentUserHandle
	toUserHandle := currentUserHandle

	// If the Note is a reply to another post, then the activity has a SocialComment object
	if replyToURLID, ok := noteObject[mastodon.InReplyTo].(string); ok {
		activity.Type = typex.SocialComment

		result, err := w.getParentStatusByParentID(ctx, replyToURLID)
		if err != nil {
			zap.L().Error("failed to get parent status and content by parent ID", zap.String("parentID", replyToURLID), zap.String("ID", message.ID))
		}

		var (
			targetContent string
			targetTime    uint64
		)

		if result != nil {
			targetContent, err = html2text.FromString(result.Content, html2text.Options{
				PrettyTables: true,
			})

			if err != nil {
				zap.L().Error("failed to convert HTML to text", zap.Error(err), zap.String("ID", replyToURLID))

				targetContent = result.Content
			}

			targetTime = w.parseTimestamp(result.Timestamp)
		}

		post.Target = w.buildTarget(replyToURLID, targetContent, targetTime)
		post.TargetURL = replyToURLID
		toUserHandle = convertURLToHandle(replyToURLID)

		if result != nil && result.Attachments != nil {
			w.buildPostMedia(post.Target, result.Attachments)
		}

		if result != nil && result.Tags != nil {
			w.buildPostTags(post.Target, result.Tags)
		}
	}

	activity.To = toUserHandle

	mainAction := w.createAction(activity.Type, activity.Tag, note.ID, currentUserHandle, toUserHandle, post)

	activity.Actions = append(activity.Actions, mainAction)

	activity.TotalActions = uint(len(activity.Actions))

	handles := []string{
		activity.To,
		activity.From,
	}

	// Store unique handles of this activity
	err := w.saveMastodonHandles(ctx, handles)
	if err != nil {
		zap.L().Error("failed to save mastodon handle", zap.Error(err), zap.String("currentUserHandle", currentUserHandle))
		return err
	}

	return nil
}

// handleActivityPubAnnounce handles Announce activities (shares/boosts) in ActivityPub.
func (w *worker) handleActivityPubAnnounce(ctx context.Context, message activitypub.Object, activity *activityx.Activity) error {
	objects, err := extractAnnounceObjects(message.Object)
	if err != nil {
		return fmt.Errorf("failed to extract objects: %w", err)
	}

	currentUserHandle := convertURLToHandle(message.ID)

	// If the actor is from a relay server then we directly set it as the handle
	if currentUserHandle == "" {
		currentUserHandle = message.Actor
	}

	for _, obj := range objects {
		if err := w.handleSingleActivityPubAnnounce(ctx, message, obj, activity, currentUserHandle); err != nil {
			return err
		}
	}

	return nil
}

// handleSingleActivityPubAnnounce handles a single Announce activity.
func (w *worker) handleSingleActivityPubAnnounce(ctx context.Context, message activitypub.Object, obj activitypub.Object, activity *activityx.Activity, currentUserHandle string) error {
	activity.Type = typex.SocialShare
	activity.Tag = tag.Social
	activity.From = currentUserHandle

	// Extract the numeric part
	publicationID := ExtractPublicationID(message.ID)

	sharedID := obj.ID
	activity.To = convertURLToHandle(sharedID)
	toUserHandle := activity.To

	result, err := w.getParentStatusByParentID(ctx, sharedID)

	if err != nil {
		zap.L().Error("failed to get parent status and content by parent ID", zap.String("parentID", sharedID), zap.String("ID", message.ID))
	}

	var (
		targetContent string
		targetTime    uint64
	)

	if result != nil {
		targetContent, err = html2text.FromString(result.Content, html2text.Options{
			PrettyTables: true,
		})

		if err != nil {
			zap.L().Error("failed to convert HTML to text", zap.Error(err), zap.String("ID", sharedID))

			targetContent = result.Content
		}

		targetTime = w.parseTimestamp(result.Timestamp)
	}

	// Create a SocialPost object with the Announced ID
	post := &metadata.SocialPost{
		Handle:        currentUserHandle,
		PublicationID: publicationID,
		Timestamp:     w.parseTimestamp(message.Published),
		Target:        w.buildTarget(sharedID, targetContent, targetTime),
		TargetURL:     sharedID,
	}

	// set Attachment and Tags to target metadata
	if result != nil && result.Attachments != nil {
		w.buildPostMedia(post.Target, result.Attachments)
	}

	if result != nil && result.Tags != nil {
		w.buildPostTags(post.Target, result.Tags)
	}

	// Remove the "/activity" suffix from the message ID
	messageID := strings.TrimSuffix(message.ID, "/activity")

	// Create and add action to activity
	action := w.createAction(activity.Type, activity.Tag, messageID, currentUserHandle, toUserHandle, post)
	activity.Actions = append(activity.Actions, action)

	activity.TotalActions = uint(len(activity.Actions))

	handles := []string{
		activity.To,
		activity.From,
	}

	// Store the current user's unique handle
	err = w.saveMastodonHandles(ctx, handles)
	if err != nil {
		zap.L().Error("failed to save mastodon handle", zap.Error(err), zap.String("currentUserHandle", currentUserHandle))
		return err
	}

	return nil
}

// saveMastodonHandles store the unique handles into the relevant DB table
func (w *worker) saveMastodonHandles(ctx context.Context, handles []string) error {
	now := uint64(time.Now().Unix())

	handleSlice := make([]*model.MastodonHandle, 0, len(handles))

	for _, handleString := range handles {
		mastodonHandle := &model.MastodonHandle{
			Handle:      handleString,
			LastUpdated: now,
		}
		handleSlice = append(handleSlice, mastodonHandle)
	}

	if err := w.databaseClient.SaveRecentMastodonHandles(ctx, handleSlice); err != nil {
		return fmt.Errorf("failed to update recent handles: %w", err)
	}

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

// createAction creates an activity action.
func (w *worker) createAction(actionType schema.Type, actionTag tag.Tag, relatedID string, from, to string, metadata metadata.Metadata) *activityx.Action {
	return &activityx.Action{
		Type:        actionType,
		Tag:         actionTag,
		Platform:    w.Platform(),
		From:        from,
		To:          to,
		Metadata:    metadata,
		RelatedURLs: []string{relatedID},
	}
}

// buildTarget constructs a SocialComment Target
func (w *worker) buildTarget(parentID string, content string, timeStamp uint64) *metadata.SocialPost {
	handle := convertURLToHandle(parentID)

	// Extract the numeric part at the end of the string
	publicationID := ExtractPublicationID(parentID)

	post := &metadata.SocialPost{
		PublicationID: publicationID,
		Handle:        handle,
		Body:          content,
		Timestamp:     timeStamp,
	}

	return post
}

// buildPost constructs a SocialPost object from ActivityPub object and note
func (w *worker) buildPost(_ context.Context, obj activitypub.Object, note activitypub.Note, timestamp uint64) *metadata.SocialPost {
	// Create a new SocialPost with the content, profile ID, publication ID, and timestamp
	handle := convertURLToHandle(obj.ID)

	// Extract the numeric part at the end of the string
	publicationID := ExtractPublicationID(note.ID)

	post := &metadata.SocialPost{
		PublicationID: publicationID,
		Timestamp:     timestamp,
		Handle:        handle,
	}

	var err error

	post.Body, err = html2text.FromString(note.Content, html2text.Options{
		PrettyTables: true,
	})

	if err != nil {
		zap.L().Error("failed to convert HTML to text", zap.Error(err), zap.String("ID", note.ID))

		post.Body = note.Content
	}

	// Attach media to the post
	if objMap, ok := obj.Object.(map[string]interface{}); ok {
		w.buildPostMedia(post, objMap[mastodon.Attachment])
		w.buildPostTags(post, objMap[mastodon.Tag])
	}

	return post
}

// buildPostMedia attaches media information to the post
func (w *worker) buildPostMedia(post *metadata.SocialPost, attachments interface{}) {
	processAttachment := func(attachment map[string]interface{}) {
		if currentURL, ok := attachment[mastodon.AttachmentURL].(string); ok {
			media := metadata.Media{Address: currentURL}
			if mediaType, ok := attachment[mastodon.AttachmentMediaType].(string); ok {
				media.MimeType = mediaType
			}

			post.Media = append(post.Media, media)
		}
	}

	switch v := attachments.(type) {
	case []map[string]interface{}:
		for _, attachment := range v {
			processAttachment(attachment)
		}
	case []interface{}:
		for _, a := range v {
			if attachment, ok := a.(map[string]interface{}); ok {
				processAttachment(attachment)
			}
		}
	case []activitypub.Attachment:
		for _, attachment := range v {
			media := metadata.Media{
				Address:  attachment.URL,
				MimeType: attachment.MediaType,
			}
			post.Media = append(post.Media, media)
		}
	default:
		zap.L().Debug("Unexpected attachments type", zap.String("type", fmt.Sprintf("%T", attachments)))
	}
}

// buildPostTags builds the Tags field in the metadata
func (w *worker) buildPostTags(post *metadata.SocialPost, tags interface{}) {
	processTag := func(tagType, name string) {
		switch tagType {
		case mastodon.TagTypeHashtag:
			post.Tags = append(post.Tags, name)
		case mastodon.TagTypeMention:
			post.Tags = append(post.Tags, "@"+strings.TrimPrefix(name, "@"))
		}
	}

	switch v := tags.(type) {
	case []map[string]interface{}:
		for _, tag := range v {
			if tagType, ok := tag[mastodon.TagType].(string); ok {
				if name, ok := tag[mastodon.TagName].(string); ok {
					processTag(tagType, name)
				}
			}
		}
	case []interface{}:
		for _, t := range v {
			if tag, ok := t.(map[string]interface{}); ok {
				if tagType, ok := tag[mastodon.TagType].(string); ok {
					if name, ok := tag[mastodon.TagName].(string); ok {
						processTag(tagType, name)
					}
				}
			}
		}
	case []activitypub.Tag:
		for _, t := range v {
			processTag(t.Type, t.Name)
		}
	default:
		zap.L().Debug("Unexpected tags type", zap.String("type", fmt.Sprintf("%T", tags)))
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

// extractNoteObjects used to extract Note Objects from the ActivityPub object
func extractNoteObjects(object interface{}) ([]activitypub.Note, error) {
	var objects []activitypub.Note

	switch obj := object.(type) {
	case map[string]interface{}:
		var o activitypub.Note
		if err := mapToStruct(obj, &o); err != nil {
			return nil, fmt.Errorf("failed to convert object: %w", err)
		}

		objects = append(objects, o)
	case []interface{}:
		for _, item := range obj {
			switch item := item.(type) {
			case map[string]interface{}:
				var o activitypub.Note
				if err := mapToStruct(item, &o); err != nil {
					return nil, fmt.Errorf("failed to convert object: %w", err)
				}

				objects = append(objects, o)

			default:
				return nil, fmt.Errorf("unsupported object type in array: %T", item)
			}
		}
	default:
		return nil, fmt.Errorf("unsupported object type: %T", obj)
	}

	return objects, nil
}

// extractAnnounceObjects used to extract Objects from the ActivityPub object
func extractAnnounceObjects(object interface{}) ([]activitypub.Object, error) {
	var objects []activitypub.Object

	switch obj := object.(type) {
	case map[string]interface{}:
		var o activitypub.Object
		if err := mapToStruct(obj, &o); err != nil {
			return nil, fmt.Errorf("failed to convert object: %w", err)
		}

		objects = append(objects, o)
	case []interface{}:
		for _, item := range obj {
			switch item := item.(type) {
			case map[string]interface{}:
				var o activitypub.Object
				if err := mapToStruct(item, &o); err != nil {
					return nil, fmt.Errorf("failed to convert object: %w", err)
				}

				objects = append(objects, o)
			default:
				return nil, fmt.Errorf("unsupported object type in array: %T", item)
			}
		}
	case string:
		objects = append(objects, activitypub.Object{
			ID: obj,
		})
	default:
		return nil, fmt.Errorf("unsupported object type: %T", obj)
	}

	return objects, nil
}

// convertURLToHandle converts the statusID to a handle string
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

	// Handle the new ID format (https://domain.com/users/username/statuses/statusid)
	parts := strings.Split(parsedURL.Path, "/")
	if len(parts) >= 4 && parts[1] == "users" {
		username := parts[2]
		domain := parsedURL.Host

		return fmt.Sprintf("@%s@%s", username, domain)
	}

	// Fallback for other URL formats
	if len(parts) >= 3 {
		username := parts[2]
		domain := parsedURL.Host

		return fmt.Sprintf("@%s@%s", username, domain)
	}

	return ""
}

// getParentStatusByParentID retrieves the parent status and content by its parent ID
func (w *worker) getParentStatusByParentID(ctx context.Context, parentID string) (*activitypub.StatusResult, error) {
	apiURL := fmt.Sprintf("%s/activity", parentID)

	// Create a new HTTP GET request with the given API URL and context
	req, err := w.httpClient.Fetch(ctx, apiURL)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	body, err := io.ReadAll(req)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	// Unmarshal the JSON response into a StatusResponse struct
	var status activitypub.StatusResponse
	if err := json.Unmarshal(body, &status); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	// Return the status content and timestamp as a StatusResult
	return &activitypub.StatusResult{
		Content:     status.Object.Content,
		Timestamp:   status.Object.Published.Format(time.RFC3339),
		Attachments: status.Object.Attachment,
		Tags:        status.Object.Tag,
	}, nil
}

// NewWorker creates a new Mastodon worker instance
func NewWorker(databaseClient database.Client, redisClient rueidis.Client) (engine.Worker, error) {
	httpClient, err := httpx.NewHTTPClient()

	if err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &worker{
		httpClient:     httpClient,
		databaseClient: databaseClient,
		redisClient:    redisClient,
	}, nil
}

// ExtractPublicationID extracts the numeric publication ID from the end of a URL string.
// It returns an empty string if no numeric ID is found.
func ExtractPublicationID(url string) string {
	// Split the URL string on the "/" character
	parts := strings.Split(url, "/")

	// Iterate through the parts in reverse order
	for i := len(parts) - 1; i >= 0; i-- {
		// Check if the current part is a number
		if isNumeric(parts[i]) {
			return parts[i]
		}
	}

	return ""
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
