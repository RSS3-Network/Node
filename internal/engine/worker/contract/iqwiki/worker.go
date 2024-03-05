package iqwiki

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/iqwiki"
	"github.com/rss3-network/node/provider/ipfs"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
)

// Worker is the worker for IQWiki.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config         *config.Module
	ethereumClient ethereum.Client
	iqWikiClient   graphql.Client
	ipfsClient     ipfs.HTTPClient
	iqWikiFilterer *iqwiki.IqWikiFilterer
}

func (w *worker) Name() string {
	return filter.IQWiki.String()
}

// Filter IQWiki contract address and event hash.
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

// Match Ethereum task to IQWiki.
func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.GetNetwork().Source() == filter.NetworkEthereumSource, nil
}

// Match Ethereum task to IQWiki.
func matchEthereumIqWiki(task *source.Task) bool {
	return task.Transaction.From == iqwiki.AddressSig && iqwiki.AddressWiki == *task.Transaction.To
}

// Transform Ethereum task to feed.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	if ethereumTask.Transaction.To == nil {
		return nil, fmt.Errorf("invalid transaction to: %s", ethereumTask.Transaction.Hash)
	}

	if !matchEthereumIqWiki(ethereumTask) {
		return nil, fmt.Errorf("invalid iqwiki transaction: %s", ethereumTask.Transaction.Hash)
	}

	// Build default IQWiki feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformIQWiki))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Parse actions from task.
	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		action, err := w.parseAction(ctx, log, ethereumTask)
		if err != nil {
			continue
		}

		feed.Actions = append(feed.Actions, action)
	}

	if len(feed.Actions) == 0 {
		return nil, fmt.Errorf("no actions in transaction: %s", ethereumTask.Transaction.Hash)
	}

	// Set feed type to the first action type & total actions.
	feed.Type = feed.Actions[0].Type
	feed.Tag = filter.TagSocial
	feed.TotalActions = uint(len(feed.Actions))

	return feed, nil
}

// Parse action from Ethereum log.
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

// Get IPFS content from Ethereum.
func (w *worker) getEthereumIPFSContent(ctx context.Context, ipfsID string) ([]byte, error) {
	file, err := w.ipfsClient.Fetch(ctx, fmt.Sprintf("/ipfs/%s", ipfsID), ipfs.FetchModeQuick)
	if err != nil {
		return nil, fmt.Errorf("fetch ipfs: %w", err)
	}

	defer lo.Try(file.Close)

	return io.ReadAll(file)
}

// NewWorker creates a new IQWiki worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	// Initialize ethereum client.
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	// Initialize ipfs client.
	if instance.ipfsClient, err = ipfs.NewHTTPClient(ipfs.WithGateways(config.IPFSGateways)); err != nil {
		return nil, fmt.Errorf("new ipfs client: %w", err)
	}

	//	Initialize iqwiki client.
	if instance.iqWikiClient, err = iqwiki.NewClient(iqwiki.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize iqwiki client: %w", err)
	}

	// Initialize iqwiki filterer.
	instance.iqWikiFilterer = lo.Must(iqwiki.NewIqWikiFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
