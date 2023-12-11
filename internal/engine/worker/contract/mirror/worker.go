package mirror

import (
	"context"
	"fmt"

	"github.com/gabriel-vasile/mimetype"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/arweave"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave/contract/mirror"
	"github.com/naturalselectionlabs/rss3-node/provider/ipfs"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config        *engine.Config
	arweaveClient arweave.Client
	ipfsClient    ipfs.HTTPClient
}

func (w *worker) Name() string {
	return engine.Mirror.String()
}

// Match returns true if the task is an Arweave task.
func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	switch task.Network() {
	case filter.NetworkArweave:
		task := task.(*source.Task)

		// Check if the transaction belongs to the mirror contract.
		owner, err := arweave.PublicKeyToAddress(task.Transaction.Owner)
		if err != nil {
			return false, fmt.Errorf("parse transaction owner: %w", err)
		}

		return owner == mirror.AddressMirror, nil
	default:
		return false, nil
	}
}

// Transform returns a feed with the action of the task.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	// Cast the task to an Arweave task.
	arweaveTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the feed.
	feed, err := task.BuildFeed(schema.WithFeedPlatform(filter.PlatformMirror))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Get actions and social content timestamp from the transaction.
	actions, timestamp, err := w.transformPostOrReviseAction(ctx, arweaveTask)
	if err != nil {
		return nil, fmt.Errorf("handle arweave mirror transaction: %w", err)
	}

	feed.To = mirror.AddressMirror
	feed.Timestamp = timestamp

	// Feed type should be inferred from the action (if it's revise)
	if actions != nil {
		feed.Type = actions[0].Type
		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// transformPostOrReviseAction Returns the actions of mirror post or revise.
func (w *worker) transformPostOrReviseAction(ctx context.Context, task *source.Task) ([]*schema.Action, uint64, error) {
	var (
		contentDigest       string
		originContentDigest string
		emptyOriginDigest   bool
	)

	for _, tag := range task.Transaction.Tags {
		tagName, err := arweave.Base64Decode(tag.Name)
		if err != nil {
			return nil, 0, fmt.Errorf("base64 decode tag name: %w", err)
		}

		tagValue, err := arweave.Base64Decode(tag.Value)
		if err != nil {
			return nil, 0, err
		}

		switch string(tagName) {
		case "Content-Digest":
			contentDigest = string(tagValue)
		case "Original-Content-Digest":
			originContentDigest = string(tagValue)

			if len(string(tagValue)) == 0 {
				emptyOriginDigest = true
			}
		}
	}

	// Construct content URI from tx id
	contentURI := fmt.Sprintf("ar://%s", task.Transaction.ID)

	// Get detailed post info from transaction data
	transactionData, err := arweave.Base64Decode(task.Transaction.Data)
	if err != nil {
		return nil, 0, fmt.Errorf("invalid foramt of transaction data: %w", err)
	}

	mirrorData := gjson.ParseBytes(transactionData)

	author := mirrorData.Get("authorship.contributor").String()
	timestamp := mirrorData.Get("content.timestamp").Uint()

	var media []metadata.Media

	// Get mirror nft as media
	address := mirrorData.Get("wnft.imageURI").String()
	if address != "" {
		file, err := w.ipfsClient.Fetch(ctx, fmt.Sprintf("/ipfs/%s", address), ipfs.FetchModeQuick)
		if err != nil {
			return nil, 0, fmt.Errorf("fetch ipfs: %w", err)
		}

		defer lo.Try(file.Close)

		// Get nft mimetype
		result, err := mimetype.DetectReader(file)
		if err != nil {
			return nil, 0, fmt.Errorf("detect mimetype: %w", err)
		}

		if result == nil {
			return nil, 0, fmt.Errorf("empty result")
		}

		media = append(media, metadata.Media{
			Address:  fmt.Sprintf("ipfs://%s", address),
			MimeType: result.String(),
		})
	}

	var publicationID string

	if contentDigest == "" {
		publicationID = mirrorData.Get("digest").String()
	} else {
		publicationID = contentDigest
	}

	if originContentDigest != "" {
		publicationID = originContentDigest
	}

	// Construct mirror Metadata
	mirrorMetadata := &metadata.SocialPost{
		Title:         mirrorData.Get("content.title").String(),
		Body:          mirrorData.Get("content.body").String(),
		ContentURI:    contentURI,
		PublicationID: publicationID,
		Media:         media,
	}

	// Build the post or revise action
	action, err := w.buildPostOrReviseAction(ctx, task.Transaction.ID, author, mirror.AddressMirror, mirrorMetadata, emptyOriginDigest)
	if err != nil {
		return nil, 0, fmt.Errorf("build post action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, timestamp, nil
}

// buildArweaveTransactionTransferAction Returns the native transfer transaction action.
func (w *worker) buildPostOrReviseAction(_ context.Context, _, from, to string, mirrorMetadata *metadata.SocialPost, emptyOriginDigest bool) (*schema.Action, error) {
	// Default action type is post.
	filterType := filter.TypeSocialPost

	// if the origin digest is empty, the action type should be revise.
	if emptyOriginDigest {
		filterType = filter.TypeSocialRevise
	}

	// TODO Should also use dataset to identify the action type.

	// Construct action
	action := schema.Action{
		Type:     filterType,
		Tag:      filter.TagSocial,
		Platform: filter.PlatformMirror.String(),
		From:     from,
		To:       to,
		Metadata: mirrorMetadata,
	}

	return &action, nil
}

// NewWorker returns a new Arweave worker.
func NewWorker(config *engine.Config) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	var err error

	if instance.arweaveClient, err = arweave.NewClient(); err != nil {
		return nil, fmt.Errorf("new arweave client: %w", err)
	}

	if instance.ipfsClient, err = ipfs.NewHTTPClient(); err != nil {
		return nil, fmt.Errorf("new ipfs client: %w", err)
	}

	return &instance, nil
}
