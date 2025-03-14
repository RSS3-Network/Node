package matters

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-shiori/go-readability"
	"github.com/rss3-network/node/v2/config"
	"github.com/rss3-network/node/v2/internal/engine"
	source "github.com/rss3-network/node/v2/internal/engine/protocol/ethereum"
	"github.com/rss3-network/node/v2/internal/utils"
	"github.com/rss3-network/node/v2/provider/ethereum"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
	"github.com/rss3-network/node/v2/provider/ethereum/contract/matters"
	"github.com/rss3-network/node/v2/provider/ethereum/token"
	"github.com/rss3-network/node/v2/provider/httpx"
	"github.com/rss3-network/node/v2/provider/ipfs"
	"github.com/rss3-network/node/v2/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"golang.org/x/net/html"
)

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config            *config.Module
	mattersFilterer   *matters.MattersFilterer
	ipfsClient        ipfs.HTTPClient
	httpClient        httpx.Client
	ethereumClient    ethereum.Client
	tokenClient       token.Client
	readabilityParser readability.Parser
}

func (w *worker) Name() string {
	return decentralized.Matters.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformMatters.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Optimism,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Social,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.SocialReward,
	}
}

// Filter returns a filter for protocol.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			matters.AddressCuration,
		},
		LogTopics: []common.Hash{
			matters.EventCuration,
		},
	}
}

// Transform matters task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	// Cast the task to a matters task.
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build the activity.
	activity, err := task.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build activity: %w", err)
	}

	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*activityx.Action
			err     error
		)

		switch {
		case w.matchEthereumCurationTransaction(log):
			actions, err = w.handleEthereumCurationTransaction(ctx, *ethereumTask, *log, activity)
		default:
			continue
		}

		if err != nil {
			zap.L().Error("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))

			return nil, err
		}

		activity.Actions = append(activity.Actions, actions...)
	}

	activity.TotalActions = uint(len(activity.Actions))
	activity.Tag = tag.Social

	return activity, nil
}

func (w *worker) matchEthereumCurationTransaction(log *ethereum.Log) bool {
	return len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], matters.EventCuration)
}

func (w *worker) handleEthereumCurationTransaction(ctx context.Context, task source.Task, log ethereum.Log, activity *activityx.Activity) ([]*activityx.Action, error) {
	activity.Type = typex.SocialReward

	export := log.Export()

	event, err := w.mattersFilterer.ParseCuration(export)
	if err != nil {
		return nil, fmt.Errorf("ParseCuration event: %w", err)
	}

	action, err := w.buildEthereumCurationAction(ctx, task, event.From, event.To, event.Token, event.Amount, event.Uri)

	if err != nil {
		return nil, fmt.Errorf("build social reward action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

func (w *worker) fetchArticle(ctx context.Context, uri string) (*readability.Article, error) {
	_, path, err := ipfs.ParseURL(uri)

	var (
		readCloser io.ReadCloser
	)

	switch {
	case path != "":
		if err != nil {
			return nil, fmt.Errorf("parse IPFS URL: %w", err)
		}

		readCloser, err = w.ipfsClient.Fetch(ctx, path, ipfs.FetchModeQuick)
		if err != nil {
			return nil, fmt.Errorf("quick fetch %s: %w", path, err)
		}
	default:
		readCloser, err = w.httpClient.Fetch(ctx, uri)
		if err != nil {
			return nil, fmt.Errorf("fetch metadata from HTTP %s: %w", uri, err)
		}
	}

	node, err := html.Parse(readCloser)
	if err != nil {
		return nil, fmt.Errorf("parse html: %w", err)
	}

	w.removeUnusedHTMLTags(node)

	article, err := w.readabilityParser.ParseDocument(node, nil)
	if err != nil {
		return nil, fmt.Errorf("format article: %w", err)
	}

	converter := md.NewConverter("", true, nil)

	if article.Content, err = converter.ConvertString(article.Content); err != nil {
		return nil, fmt.Errorf("convert article: %w", err)
	}

	article.Byline = w.filterName(article.Byline)

	return &article, nil
}

func (w *worker) filterName(s string) string {
	i := strings.Index(s, "(")
	if i >= 0 {
		j := strings.Index(s[i:], ")")
		if j >= 0 {
			return s[i+1 : i+j]
		}
	}

	return ""
}

func (w *worker) removeUnusedHTMLTags(node *html.Node) {
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode && child.Data == "small" {
			child.Parent.RemoveChild(child)

			w.removeUnusedHTMLTags(node)
		} else {
			w.removeUnusedHTMLTags(child)
		}
	}
}

func (w *worker) buildEthereumCurationAction(ctx context.Context, task source.Task, trigger, recipient, token common.Address, amount *big.Int, uri string) (*activityx.Action, error) {
	article, err := w.fetchArticle(ctx, uri)
	if err != nil || article == nil {
		return nil, fmt.Errorf("fetch article: %w", err)
	}

	rewardTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, lo.ToPtr(token), nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", "", err)
	}

	rewardTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(utils.GetBigInt(amount), 0))

	return &activityx.Action{
		Type:     typex.SocialReward,
		Tag:      tag.Social,
		Platform: w.Platform(),
		From:     trigger.String(),
		To:       recipient.String(),
		Metadata: &metadata.SocialPost{
			Reward: rewardTokenMetadata,
			Target: &metadata.SocialPost{
				Handle:  article.Byline,
				Title:   article.Title,
				Summary: article.Excerpt,
				Body:    article.Content,
			},
		},
	}, nil
}

// NewWorker returns a new matters worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config:          config,
		mattersFilterer: lo.Must(matters.NewMattersFilterer(ethereum.AddressGenesis, nil)),
		ipfsClient:      lo.Must(ipfs.NewHTTPClient(ipfs.WithGateways(config.IPFSGateways))),
		httpClient:      lo.Must(httpx.NewHTTPClient()),
	}

	var err error

	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint.URL, config.Endpoint.BuildEthereumOptions()...); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	instance.tokenClient = token.NewClient(instance.ethereumClient)

	if instance.ipfsClient, err = ipfs.NewHTTPClient(ipfs.WithGateways(config.IPFSGateways)); err != nil {
		return nil, fmt.Errorf("new ipfs client: %w", err)
	}

	instance.readabilityParser = readability.NewParser()

	return &instance, nil
}
