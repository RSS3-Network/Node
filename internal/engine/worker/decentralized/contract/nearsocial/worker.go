package nearsocial

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/near"
	"github.com/rss3-network/node/v2/provider/near"
	workerx "github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ engine.Worker = (*worker)(nil)

type worker struct {
	config *config.Module
}

func (w *worker) Name() string {
	return workerx.NearSocial.String()
}

func (w *worker) Platform() string {
	return workerx.PlatformNearSocial.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Near,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Social,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.SocialPost,
		typex.SocialComment,
		typex.SocialShare,
	}
}

const (
	socialNearReceiverID = "social.near"
)

func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		ReceiverIDs: []string{
			socialNearReceiverID,
		},
	}
}

// Transform processes a Near task and returns an activity.
// It is the main entry point for processing Near Social transactions.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity with the platform information.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Process the Near Social action in the transaction.
	action, err := w.handleNearSocialAction(ctx, nearTask, activity.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("handle near social action: %w", err)
	}

	if action != nil {
		activity.Type = action.Type
		activity.Actions = append(activity.Actions, action)
	} else {
		return nil, fmt.Errorf("no action found in transaction")
	}

	return activity, nil
}

// handleNearSocialAction processes the action in the Near Social transaction and returns an activityx.Action.
// This function is responsible for identifying and processing the 'set' function call in the transaction.
func (w *worker) handleNearSocialAction(_ context.Context, task *source.Task, timestamp uint64) (*activityx.Action, error) {
	for _, action := range task.Transaction.Transaction.Actions {
		if action.FunctionCall != nil && action.FunctionCall.MethodName == "set" {
			return w.processSetFunction(task.Transaction.Transaction.SignerID, action.FunctionCall, timestamp)
		}
	}

	return nil, nil
}

// processSetFunction handles the "set" function call and returns an activityx.Action.
// This function decodes and processes the arguments of the 'set' function call to build a social action.
func (w *worker) processSetFunction(signerID string, functionCall *near.FunctionCallAction, timestamp uint64) (*activityx.Action, error) {
	// Decode base64 args
	decodedArgs, err := base64.StdEncoding.DecodeString(functionCall.Args)
	if err != nil {
		return nil, fmt.Errorf("decode base64 args: %w", err)
	}

	// Unmarshal decoded args into FunctionCallArgs struct
	var args FunctionCallArgs
	if err := json.Unmarshal(decodedArgs, &args); err != nil {
		return nil, fmt.Errorf("unmarshal function call args: %w", err)
	}

	for _, userContent := range args.Data {
		if userContent.Post != nil {
			for path, postDataJSON := range userContent.Post {
				var postData PostData
				if err := json.Unmarshal([]byte(postDataJSON), &postData); err != nil {
					return nil, fmt.Errorf("unmarshal post data: %w", err)
				}

				return w.buildSocialAction(signerID, path, postData, args, timestamp)
			}
		}

		if userContent.Index != nil {
			return w.buildSocialAction(signerID, "", PostData{}, args, timestamp)
		}
	}

	return nil, nil
}

// buildSocialAction constructs an activityx.Action based on the provided social action data.
// This function is crucial for creating the appropriate action type (post, comment, or share) and populating its metadata.
func (w *worker) buildSocialAction(signerID, path string, postData PostData, args FunctionCallArgs, timestamp uint64) (*activityx.Action, error) {
	action := &activityx.Action{
		Type:     typex.SocialPost,
		Platform: w.Platform(),
		From:     signerID,
		To:       signerID,
		Metadata: metadata.SocialPost{
			Handle:    signerID,
			Body:      postData.Text,
			Timestamp: timestamp,
		},
	}

	if path == "comment" {
		action = w.handleComment(action, postData)
	}

	if userContent, ok := args.Data[signerID]; ok {
		action = w.processUserContent(action, userContent, signerID, timestamp)
	}

	return action, nil
}

func (w *worker) handleComment(action *activityx.Action, postData PostData) *activityx.Action {
	action.Type = typex.SocialComment

	if postData.Item != nil {
		target := &metadata.SocialPost{
			Handle:        strings.Split(postData.Item.Path, "/")[0],
			PublicationID: fmt.Sprintf("%s-%d", postData.Item.Path, postData.Item.BlockHeight),
		}

		if socialPost, ok := action.Metadata.(metadata.SocialPost); ok {
			socialPost.Target = target
			action.Metadata = socialPost
		}

		action.To = target.Handle
	}

	return action
}

func (w *worker) processUserContent(action *activityx.Action, userContent UserContent, signerID string, timestamp uint64) *activityx.Action {
	action = w.processHashtags(action, userContent)
	action = w.processCommentIndex(action, userContent)
	action = w.processReposts(action, userContent, signerID, timestamp)

	return action
}

func (w *worker) processHashtags(action *activityx.Action, userContent UserContent) *activityx.Action {
	if hashtagJSON, ok := userContent.Index["hashtag"]; ok {
		var hashtags []HashtagData

		if err := json.Unmarshal([]byte(hashtagJSON), &hashtags); err == nil {
			tags := make([]string, 0, len(hashtags))

			for _, hashtag := range hashtags {
				tags = append(tags, hashtag.Key)
			}

			if socialPost, ok := action.Metadata.(metadata.SocialPost); ok {
				socialPost.Tags = tags
				action.Metadata = socialPost
			}
		}
	}

	return action
}

func (w *worker) processCommentIndex(action *activityx.Action, userContent UserContent) *activityx.Action {
	if indexJSON, ok := userContent.Index["comment"]; ok {
		var indexData IndexData
		if err := json.Unmarshal([]byte(indexJSON), &indexData); err == nil {
			if socialPost, ok := action.Metadata.(metadata.SocialPost); ok {
				socialPost.PublicationID = fmt.Sprintf("%s-%d", indexData.Key.Path, indexData.Key.BlockHeight)
				action.Metadata = socialPost
			}
		}
	}

	return action
}

func (w *worker) processReposts(action *activityx.Action, userContent UserContent, signerID string, timestamp uint64) *activityx.Action {
	if repostJSON, ok := userContent.Index["repost"]; ok {
		var repostData RepostIndexData
		if err := json.Unmarshal([]byte(repostJSON), &repostData); err == nil {
			if len(repostData) > 0 {
				action.Type = typex.SocialShare

				if repostData[0].Value.Item != nil {
					pathParts := strings.Split(repostData[0].Value.Item.Path, "/")
					if len(pathParts) >= 2 {
						target := &metadata.SocialPost{
							Handle:        pathParts[0],
							PublicationID: fmt.Sprintf("%s-%d", repostData[0].Value.Item.Path, repostData[0].Value.Item.BlockHeight),
						}

						action.Metadata = metadata.SocialPost{
							Handle:    signerID,
							Timestamp: timestamp,
							Target:    target,
						}
						action.To = target.Handle
					}
				}
			}
		}
	}

	return action
}

// NewWorker returns a new near social worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	return &instance, nil
}
