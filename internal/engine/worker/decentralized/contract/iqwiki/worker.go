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
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
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
	return decentralized.IQWiki.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformIQWiki.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Polygon,
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

// Filter IQWiki contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
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
func matchEthereumIqWiki(task *source.Task) bool {
	return task.Transaction.From == iqwiki.AddressSig && iqwiki.AddressWiki == *task.Transaction.To
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
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

	// Build default IQWiki activity from task.
	activity, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
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

		activity.Actions = append(activity.Actions, action)
	}

	if len(activity.Actions) == 0 {
		return nil, fmt.Errorf("no actions in transaction: %s", ethereumTask.Transaction.Hash)
	}

	// Set activity type to the first action type & total actions.
	activity.Type = activity.Actions[0].Type
	activity.Tag = tag.Social
	activity.TotalActions = uint(len(activity.Actions))

	return activity, nil
}

// Parse action from Ethereum log.
func (w *worker) parseAction(ctx context.Context, log *ethereum.Log, ethereumTask *source.Task) (*activityx.Action, error) {
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

	action := &activityx.Action{
		Tag:      tag.Social,
		Type:     lo.If(iqwiki.StatusCreated == wikiResponse.ActivityByWikiIdAndBlock.Type, typex.SocialPost).Else(typex.SocialRevise),
		Platform: w.Platform(),
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
			Tags: lo.Map(wikiResponse.ActivityByWikiIdAndBlock.Content[0].Tags, func(x iqwiki.ActivityByWikiIdAndBlockActivityByWikiIdAndBlockActivityContentWikiTagsTag, _ int) string {
				return x.Id
			}),
			Media: lo.Map(wikiResponse.ActivityByWikiIdAndBlock.Content[0].Images, func(image iqwiki.ActivityByWikiIdAndBlockActivityByWikiIdAndBlockActivityContentWikiImagesImage, _ int) metadata.Media {
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
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
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
