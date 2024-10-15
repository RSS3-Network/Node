package nearsocial

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/near"
	"github.com/rss3-network/node/provider/near"
	workerx "github.com/rss3-network/node/schema/worker/decentralized"
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

// Filter returns a source filter.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		ReceiverIDs: []string{
			socialNearReceiverID,
		},
	}
}

// Transform returns an activity with the actions of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a Near task.
	nearTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Handle all actions in the transaction
	actions, err := w.handleNearSocialActions(ctx, nearTask, activity.Timestamp)
	if err != nil {
		return nil, fmt.Errorf("handle near social actions: %w", err)
	}

	if len(actions) > 0 {
		activity.Type = actions[0].Type
		activity.Actions = append(activity.Actions, actions...)
	} else {
		return nil, fmt.Errorf("no actions found in transaction")
	}

	return activity, nil
}

// handleNearSocialActions processes all actions in the near social transaction and returns a slice of activityx.Action.
func (w *worker) handleNearSocialActions(_ context.Context, task *source.Task, timestamp uint64) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	for _, action := range task.Transaction.Transaction.Actions {
		if action.FunctionCall != nil && action.FunctionCall.MethodName == "set" {
			newActions, err := w.processSetFunction(task.Transaction.Transaction.SignerID, action.FunctionCall, timestamp)
			if err != nil {
				return nil, err
			}

			actions = append(actions, newActions...)
		}
	}

	return actions, nil
}

// processSetFunction handles the "set" function call and returns a slice of activityx.Action.
func (w *worker) processSetFunction(signerID string, functionCall *near.FunctionCallAction, timestamp uint64) ([]*activityx.Action, error) {
	var actions []*activityx.Action

	// Decode base64 args
	decodedArgs, err := base64.StdEncoding.DecodeString(functionCall.Args)
	if err != nil {
		return nil, fmt.Errorf("decode base64 args: %w", err)
	}

	// Unmarshal decoded args
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

				action := w.buildSocialPostAction(signerID, path, postData, args, timestamp)
				actions = append(actions, action)
			}
		}
	}

	return actions, nil
}

func (w *worker) buildSocialPostAction(signerID, path string, postData PostData, args FunctionCallArgs, timestamp uint64) *activityx.Action {
	action := &activityx.Action{
		Type:     typex.SocialPost,
		Platform: w.Platform(),
		From:     signerID,
		To:       socialNearReceiverID,
		Metadata: metadata.SocialPost{
			Handle:    signerID,
			Body:      postData.Text,
			Timestamp: timestamp,
		},
	}

	// Handle comment
	if path == "comment" {
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
		}
	}

	// Extract hashtags and mentions from args
	if userContent, ok := args.Data[signerID]; ok {
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

		if indexJSON, ok := userContent.Index["comment"]; ok {
			var indexData IndexData
			if err := json.Unmarshal([]byte(indexJSON), &indexData); err == nil {
				if socialPost, ok := action.Metadata.(metadata.SocialPost); ok {
					socialPost.PublicationID = fmt.Sprintf("%s-%d", indexData.Key.Path, indexData.Key.BlockHeight)
					action.Metadata = socialPost
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
