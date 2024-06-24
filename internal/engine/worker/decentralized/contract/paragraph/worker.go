package paragraph

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/arweave"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/arweave/contract/paragraph"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config        *config.Module
	arweaveClient arweave.Client
	httpClient    httpx.Client
}

func (w *worker) Name() string {
	return decentralized.Paragraph.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformParagraph.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Arweave,
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
		typex.SocialRevise,
	}
}

// Filter returns a filter for source.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{OwnerAddresses: []string{paragraph.AddressParagraph}}
}

// Transform returns an activity  with the action of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to an Arweave task.
	arweaveTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	// Get actions and social content timestamp from the transaction.
	actions, err := w.transformParagraphAction(ctx, arweaveTask)
	if err != nil {
		return nil, fmt.Errorf("handle arweave paragraph transaction: %w", err)
	}

	activity.To = paragraph.AddressParagraph

	// Activity type should be inferred from the action (if it's revise)
	if actions[0] != nil {
		activity.Type = actions[0].Type
		activity.Actions = append(activity.Actions, actions...)
	}

	return activity, nil
}

// transformPostOrReviseAction Returns the actions of mirror post or revise.
func (w *worker) transformParagraphAction(ctx context.Context, task *source.Task) ([]*activityx.Action, error) {
	var (
		contributor     string
		publicationSlug string
		contentType     string
		postID          string
	)

	for _, tag := range task.Transaction.Tags {
		tagName, err := arweave.Base64Decode(tag.Name)
		if err != nil {
			return nil, fmt.Errorf("base64 decode tag name failed: %w", err)
		}

		tagValue, err := arweave.Base64Decode(tag.Value)
		if err != nil {
			return nil, err
		}

		switch string(tagName) {
		case "Contributor":
			contributor = string(tagValue)
		case "PublicationSlug":
			publicationSlug = strings.ReplaceAll(string(tagValue), "@", "")
		case "Content-Type":
			contentType = string(tagValue)
		case "PostId":
			postID = string(tagValue)
		}
	}

	transactionData, err := arweave.Base64Decode(task.Transaction.Data)
	if err != nil {
		return nil, fmt.Errorf("invalid foramt of transaction data: %w", err)
	}

	paragraphData := gjson.ParseBytes(transactionData)

	contentURI := fmt.Sprintf("https://arweave.net/%s", task.Transaction.ID)

	paragraphMetadata, err := w.buildParagraphMetadata(ctx, publicationSlug, contentURI, transactionData)
	if err != nil {
		return nil, fmt.Errorf("build arweave paragraph post metadata failed: %w", err)
	}

	var updated bool

	if paragraphData.Get("arweaveId").Exists() {
		updated = true
	}

	var action *activityx.Action

	if contributor != "" && postID != "" && contentType == "application/json" {
		// Build the post or revise action
		action = w.buildParagraphAction(ctx, contributor, paragraph.AddressParagraph, paragraphMetadata, updated)
	}

	actions := []*activityx.Action{
		action,
	}

	return actions, nil
}

// buildArweaveTransactionTransferAction Returns the native transfer transaction action.
func (w *worker) buildParagraphAction(_ context.Context, from, to string, paragraphMetadata *metadata.SocialPost, updated bool) *activityx.Action {
	// Default action type is post.
	filterType :=
		typex.SocialPost

	if updated {
		filterType =
			typex.SocialRevise
	}

	// Construct action
	action := activityx.Action{
		Type:     filterType,
		Tag:      tag.Social,
		Platform: w.Platform(),
		From:     from,
		To:       to,
		Metadata: paragraphMetadata,
	}

	return &action
}

// buildParagraphMetadata Returns the metadata of the paragraph post.
func (w *worker) buildParagraphMetadata(ctx context.Context, handle, contentURI string, contentData []byte) (*metadata.SocialPost, error) {
	paragraphData := gjson.ParseBytes(contentData)

	var media []metadata.Media

	if paragraphData.Get("cover_img").Exists() {
		address := paragraphData.Get("cover_img.img.src").String()
		if address != "" {
			mimeType, err := w.httpClient.GetContentType(ctx, address)
			if err != nil {
				return nil, fmt.Errorf("get content type failed: %w", err)
			}

			if mimeType != "" {
				media = append(media, metadata.Media{
					Address:  address,
					MimeType: mimeType,
				})
			}
		}
	}

	// Get media from collectibleImgUrl
	if paragraphData.Get("collectibleImgUrl").Exists() {
		address := paragraphData.Get("collectibleImgUrl").String()
		if address != "" {
			mimeType, err := w.httpClient.GetContentType(ctx, address)
			if err != nil {
				return nil, fmt.Errorf("get content type failed: %w", err)
			}

			if mimeType != "" {
				media = append(media, metadata.Media{
					Address:  address,
					MimeType: mimeType,
				})
			}
		}
	}

	paragraphTags := lo.Map(paragraphData.Get("categories").Array(), func(tag gjson.Result, _ int) string {
		return tag.String()
	})

	var profileID string

	if paragraphData.Get("authors").Exists() && len(paragraphData.Get("authors").Array()) > 0 {
		profileID = paragraphData.Get("authors").Array()[0].String()
	}

	return &metadata.SocialPost{
		Handle:        handle,
		Title:         paragraphData.Get("title").String(),
		Summary:       paragraphData.Get("post_preview").String(),
		Body:          paragraphData.Get("markdown").String(),
		ContentURI:    contentURI,
		ProfileID:     profileID,
		PublicationID: paragraphData.Get("slug").String(),
		Media:         media,
		Tags:          paragraphTags,
		Timestamp:     uint64(time.UnixMilli(paragraphData.Get("updatedAt").Int()).Unix()),
	}, nil
}

// NewWorker returns a new Arweave worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	var err error

	if instance.arweaveClient, err = arweave.NewClient(); err != nil {
		return nil, fmt.Errorf("new arweave client: %w", err)
	}

	if instance.httpClient, err = httpx.NewHTTPClient(); err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	return &instance, nil
}
