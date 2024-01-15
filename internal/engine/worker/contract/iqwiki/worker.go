package iqwiki

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/iqwiki"
	"github.com/naturalselectionlabs/rss3-node/provider/ipfs"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
)

// Worker is the worker for OpenSea.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	iqWikiFilterer *iqwiki.IqWikiFilterer
	iqWikiClient   graphql.Client
	ipfsClient     ipfs.HTTPClient
}

func (w *worker) Name() string {
	return engine.IQWiki.String()
}

// Filter opensea contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			iqwiki.AddressWiki,
			iqwiki.AddressSig,
		},
		LogTopics: []common.Hash{
			iqwiki.EventPosted,
		},
	}
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	isEthereumNetwork := task.GetNetwork().Source() == filter.NetworkEthereumSource
	if !isEthereumNetwork {
		return false, nil
	}

	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return false, fmt.Errorf("invalid task type: %T", task)
	}

	isEmptyTo := ethereumTask.Transaction.To == nil
	if isEmptyTo {
		return false, nil
	}

	return matchEthereumIqWiki(ethereumTask)
}

func matchEthereumIqWiki(task *source.Task) (bool, error) {
	return task.Transaction.From == iqwiki.AddressSig && iqwiki.AddressWiki == *task.Transaction.To, nil
}

// Transform Ethereum task to feed.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default opensea feed from task.
	feed, _ := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformIQWiki))

	// Parse actions from task.
	w.parseActions(ctx, ethereumTask, feed)

	// Set feed type to the first action type & total actions.
	feed.Type = feed.Actions[0].Type
	feed.Tag = filter.TagSocial
	feed.TotalActions = uint(len(feed.Actions))

	return feed, nil
}

func (w *worker) parseActions(ctx context.Context, ethereumTask *source.Task, feed *schema.Feed) {
	for _, log := range ethereumTask.Receipt.Logs {
		action, err := w.parseAction(ctx, log, ethereumTask)
		if err != nil {
			continue
		}

		feed.Actions = append(feed.Actions, action)
	}
}

func (w *worker) parseAction(ctx context.Context, log *ethereum.Log, ethereumTask *source.Task) (*schema.Action, error) {
	wikiPosted, err := w.iqWikiFilterer.ParsePosted(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse posted: %w", err)
	}

	ipfsID := wikiPosted.Ipfs

	content, err := w.getEthereumIPFSContent(ctx, ipfsID)
	if err != nil {
		return nil, fmt.Errorf("get ipfs content: %w", err)
	}

	var wiki struct {
		ID string `json:"id"`
	}

	err = json.Unmarshal(content, &wiki)
	if err != nil {
		return nil, fmt.Errorf("unmarshal wiki error, hash %s, %w", ethereumTask.Transaction.Hash, err)
	}

	if wiki.ID == "" {
		return nil, fmt.Errorf("invalid wiki, hash %s, %w", ethereumTask.Transaction.Hash, err)
	}

	wikiResponse, err := iqwiki.ActivityByWikiIdAndBlock(context.Background(), w.iqWikiClient, int(ethereumTask.Header.Number.Int64()), wiki.ID)
	if err != nil {
		return nil, fmt.Errorf("fetch wiki %s, hash %s, %w", wiki.ID, ethereumTask.Transaction.Hash.String(), err)
	}

	action := &schema.Action{
		Tag:      filter.TagSocial,
		Type:     lo.If(iqwiki.StatusCreated == wikiResponse.ActivityByWikiIdAndBlock.Type, filter.TypeSocialPost).Else(filter.TypeSocialRevise),
		Platform: filter.PlatformIQWiki.String(),
		From:     wikiPosted.From.String(),
		To:       iqwiki.AddressWiki.String(),
		Metadata: metadata.SocialPost{
			Handle:        wikiResponse.ActivityByWikiIdAndBlock.User.Profile.Username,
			PublicationID: wikiResponse.ActivityByWikiIdAndBlock.WikiId,
			Body:          wikiResponse.ActivityByWikiIdAndBlock.Content[0].Content,
			Title:         wikiResponse.ActivityByWikiIdAndBlock.Content[0].Title,
			Summary:       wikiResponse.ActivityByWikiIdAndBlock.Content[0].Summary,
			ContentURI:    fmt.Sprintf("https://iq.wiki/wiki/%s", wikiResponse.ActivityByWikiIdAndBlock.WikiId),
			AuthorURL:     fmt.Sprintf("https://iq.wiki/account/%s", wikiResponse.ActivityByWikiIdAndBlock.User.Id),
			Tags: lo.Map(wikiResponse.ActivityByWikiIdAndBlock.Content[0].Tags, func(x iqwiki.ActivityByWikiIdAndBlockActivityByWikiIdAndBlockActivityContentWikiTagsTag, index int) string {
				return x.Id
			}),
			Media: lo.Map(wikiResponse.ActivityByWikiIdAndBlock.Content[0].Images, func(image iqwiki.ActivityByWikiIdAndBlockActivityByWikiIdAndBlockActivityContentWikiImagesImage, index int) metadata.Media {
				return metadata.Media{
					MimeType: "image/png",
					Address:  fmt.Sprintf("ipfs://%s", image.Id),
				}
			}),
		},
	}

	return action, nil
}

func (w *worker) getEthereumIPFSContent(ctx context.Context, ipfsID string) ([]byte, error) {
	file, err := w.ipfsClient.Fetch(ctx, fmt.Sprintf("/ipfs/%s", ipfsID), ipfs.FetchModeQuick)
	if err != nil {
		return nil, fmt.Errorf("fetch ipfs: %w", err)
	}

	defer lo.Try(file.Close)

	return io.ReadAll(file)
}

// NewWorker creates a new OpenSea worker.
func NewWorker() (engine.Worker, error) {
	instance := worker{
		iqWikiFilterer: lo.Must(iqwiki.NewIqWikiFilterer(ethereum.AddressGenesis, nil)),
		ipfsClient:     lo.Must(ipfs.NewHTTPClient()),
		iqWikiClient:   lo.Must(iqwiki.NewClient(iqwiki.Endpoint)),
	}

	return &instance, nil
}
