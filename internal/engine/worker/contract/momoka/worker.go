package momoka

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/arweave"
	"github.com/rss3-network/node/provider/arweave"
	"github.com/rss3-network/node/provider/arweave/contract/momoka"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/lens"
	"github.com/rss3-network/node/provider/httpx"
	"github.com/rss3-network/node/provider/ipfs"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/tidwall/gjson"
)

// make sure worker implements engine.Worker
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config           *config.Module
	ethereumClient   ethereum.Client
	arweaveClient    arweave.Client
	ipfsClient       ipfs.HTTPClient
	httpClient       httpx.Client
	lensHubV1        *lens.V1LensHubCaller
	lensHubV2        *lens.V2LensHubCaller
	lensHandleV2     *lens.V2LensHandleCaller
	handleRegistryV2 *lens.V2HandleRegistryCaller
}

func (w *worker) Name() string {
	return filter.Momoka.String()
}

// Filter returns a filter for source.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{OwnerAddresses: momoka.AddressesBundlr}
}

// Match returns true if the task is an Arweave task.
func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	switch task.GetNetwork() {
	case filter.NetworkArweave:
		task := task.(*source.Task)

		owner, err := arweave.PublicKeyToAddress(task.Transaction.Owner)
		if err != nil {
			return false, fmt.Errorf("parse transaction owner: %w", err)
		}

		return owner != "" && lo.Contains(momoka.AddressesLens, owner), nil
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
	feed, err := task.BuildFeed(schema.WithFeedPlatform(filter.PlatformLens))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Get actions and social content timestamp from the transaction.
	actions, err := w.transformMomokaAction(ctx, arweaveTask)
	if err != nil {
		return nil, fmt.Errorf("handle arweave mirror transaction: %w", err)
	}

	feed.To = feed.From

	// Feed type should be inferred from the action (if it's revise)
	if actions[0] != nil {
		feed.Type = actions[0].Type
		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// transformPostOrReviseAction Returns the actions of mirror post or revise.
func (w *worker) transformMomokaAction(ctx context.Context, task *source.Task) ([]*schema.Action, error) {
	data, err := base64.RawURLEncoding.DecodeString(task.Transaction.Data)
	if err != nil {
		return nil, fmt.Errorf("decode transaction data: %w", err)
	}

	transactionData := gjson.ParseBytes(data)

	rawPublicationID := transactionData.Get("publicationId").String()

	timestamp := transactionData.Get("event.timestamp").Uint()

	// Polygon block number
	blockNumber := transactionData.Get("chainProofs.thisPublication.blockNumber").Uint()

	var socialType filter.SocialType

	var contentURI, rawProfileID, rawProfileIDPointed string

	switch transactionData.Get("type").String() {
	case "POST_CREATED":
		socialType = filter.TypeSocialPost

		if transactionData.Get("event.postParams").Exists() {
			contentURI = transactionData.Get("event.postParams.contentURI").String()
			rawProfileID = transactionData.Get("event.postParams.profileId").String()
		} else {
			contentURI = transactionData.Get("event.contentURI").String()
			rawProfileID = transactionData.Get("event.profileId").String()
		}
	case "COMMENT_CREATED":
		socialType = filter.TypeSocialComment

		if transactionData.Get("event.commentParams").Exists() {
			contentURI = transactionData.Get("event.commentParams.contentURI").String()
			rawProfileID = transactionData.Get("event.commentParams.profileId").String()
			rawProfileIDPointed = transactionData.Get("event.commentParams.pointedProfileId").String()
		} else {
			contentURI = transactionData.Get("event.contentURI").String()
			rawProfileID = transactionData.Get("event.profileId").String()
			rawProfileIDPointed = transactionData.Get("event.profileIdPointed").String()
		}
	case "MIRROR_CREATED":
		socialType = filter.TypeSocialShare

		if transactionData.Get("event.mirrorParams").Exists() {
			rawProfileID = transactionData.Get("event.mirrorParams.profileId").String()
			rawProfileIDPointed = transactionData.Get("event.mirrorParams.pointedProfileId").String()
		} else {
			rawProfileID = transactionData.Get("event.profileId").String()
			rawProfileIDPointed = transactionData.Get("event.profileIdPointed").String()
		}
	case "QUOTE_CREATED":
		socialType = filter.TypeSocialShare
		contentURI = transactionData.Get("event.quoteParams.contentURI").String()
		rawProfileID = transactionData.Get("event.quoteParams.profileId").String()
		rawProfileIDPointed = transactionData.Get("event.quoteParams.pointedProfileId").String()
	default:
		return nil, fmt.Errorf("unsupported transaction type: %s", transactionData.Get("type").String())
	}

	// Discard unsupported transaction type
	if rawProfileID == "" || rawPublicationID == "" {
		return nil, fmt.Errorf("missing profile id or publication id")
	}

	var profile, profilePointed string

	var profileIDInt *big.Int

	if rawProfileID != "" {
		profileIDInt, err = hexutil.DecodeBig(strings.Replace(rawProfileID, "0x0", "0x", 1))
		if err != nil {
			return nil, fmt.Errorf("decode profile id: %w", err)
		}

		profile, err = w.getLensHandle(ctx, new(big.Int).SetUint64(blockNumber), profileIDInt)
		if err != nil {
			return nil, fmt.Errorf("get lens handle: %w", err)
		}
	}

	momokaMetadata, err := w.buildArweaveMomokaPostMetadata(ctx, rawProfileID, profile, rawPublicationID, contentURI, false, timestamp)
	if err != nil {
		return nil, fmt.Errorf("build arweave momoka post metadata: %w", err)
	}

	if rawProfileIDPointed != "" {
		profileIDPointedInt, err := hexutil.DecodeBig(strings.Replace(rawProfileIDPointed, "0x0", "0x", 1))
		if err != nil {
			return nil, fmt.Errorf("decode profile id pointed: %w", err)
		}

		profilePointed, err = w.getLensHandle(ctx, new(big.Int).SetUint64(blockNumber), profileIDPointedInt)
		if err != nil {
			return nil, fmt.Errorf("get lens handle: %w", err)
		}

		location := transactionData.Get("chainProofs.pointer.location").String()

		body, err := w.getDataFromHTTP(ctx, location)
		if err != nil {
			return nil, fmt.Errorf("get publication from ipfs: %w", err)
		}

		defer lo.Try(body.Close)

		targetTransaction, err := io.ReadAll(body)
		if err != nil {
			return nil, fmt.Errorf("read all: %w", err)
		}

		targetData := gjson.ParseBytes(targetTransaction)

		var targetContentURI string

		if targetData.Get("event.postParams").Exists() {
			targetContentURI = targetData.Get("event.postParams.contentURI").String()
		} else if targetData.Get("event.commentParams").Exists() {
			targetContentURI = targetData.Get("event.commentParams.contentURI").String()
		} else if targetData.Get("event.quoteParams").Exists() {
			targetContentURI = targetData.Get("event.quoteParams.contentURI").String()
		} else {
			targetContentURI = targetData.Get("event.contentURI").String()
		}

		targetPublicationID := targetData.Get("publicationId").String()

		momokaMetadata.Target, err = w.buildArweaveMomokaPostMetadata(ctx, rawProfileIDPointed, profilePointed, targetPublicationID, targetContentURI, true, timestamp)
		if err != nil {
			return nil, fmt.Errorf("build arweave momoka post metadata: %w", err)
		}
	}

	from, err := w.getLensOwnerOf(ctx, new(big.Int).SetUint64(blockNumber), profileIDInt)
	if err != nil {
		return nil, fmt.Errorf("get lens owner of: %w", err)
	}

	feedFrom, err := arweave.PublicKeyToAddress(task.Transaction.Owner)
	if err != nil {
		return nil, fmt.Errorf("public key to address: %w", err)
	}

	action, err := w.buildArweaveMomokaAction(ctx, from.String(), feedFrom, socialType, momokaMetadata)
	if err != nil {
		return nil, fmt.Errorf("build arweave momoka action: %w", err)
	}

	actions := []*schema.Action{
		action,
	}

	return actions, nil
}

func (w *worker) buildArweaveMomokaAction(_ context.Context, from, to string, filterType filter.SocialType, momokaMetadata *metadata.SocialPost) (*schema.Action, error) {
	action := schema.Action{
		Type:     filterType,
		Tag:      filter.TagSocial,
		Platform: filter.PlatformLens.String(),
		From:     from,
		To:       to,
		Metadata: momokaMetadata,
	}

	return &action, nil
}

func (w *worker) buildArweaveMomokaPostMetadata(ctx context.Context, profileID, handle, pubID string, contentURI string, isTarget bool, timestamp uint64) (*metadata.SocialPost, error) {
	var contentData []byte

	if contentURI != "" {
		var err error

		body, err := w.getDataFromHTTP(ctx, contentURI)
		if err != nil {
			return nil, fmt.Errorf("get publication from ipfs: %w", err)
		}
		defer lo.Try(body.Close)

		contentData, err = io.ReadAll(body)
		if err != nil {
			return nil, fmt.Errorf("read all: %w", err)
		}
	}

	momokaData := gjson.ParseBytes(contentData)

	var momokaMedia []metadata.Media

	var momokaTags []string

	var content string

	if momokaData.Get("lens").Exists() {
		content = momokaData.Get("lens.content").String()

		momokaImages := lo.Map(momokaData.Get("lens.image").Array(), func(media gjson.Result, index int) metadata.Media {
			return metadata.Media{
				MimeType: media.Get("type").String(),
				Address:  media.Get("item").String(),
			}
		})

		momokaAttachments := lo.Map(momokaData.Get("lens.attachments").Array(), func(media gjson.Result, index int) metadata.Media {
			return metadata.Media{
				MimeType: media.Get("type").String(),
				Address:  media.Get("item").String(),
			}
		})

		momokaVideos := lo.Map(momokaData.Get("lens.video").Array(), func(media gjson.Result, index int) metadata.Media {
			return metadata.Media{
				MimeType: media.Get("type").String(),
				Address:  media.Get("item").String(),
			}
		})

		//append all three types of media
		momokaMedia = append(momokaImages, momokaVideos...)
		momokaMedia = append(momokaMedia, momokaAttachments...)

		momokaMedia = lo.Uniq(momokaMedia)

		momokaTags = lo.Map(momokaData.Get("lens.tags").Array(), func(tag gjson.Result, index int) string {
			return tag.String()
		})
	} else {
		content = momokaData.Get("content").String()

		momokaMedia = lo.Map(momokaData.Get("media").Array(), func(media gjson.Result, index int) metadata.Media {
			return metadata.Media{
				MimeType: media.Get("type").String(),
				Address:  media.Get("item").String(),
			}
		})

		momokaTags = lo.Map(momokaData.Get("tags").Array(), func(tag gjson.Result, index int) string {
			return tag.String()
		})
	}

	return &metadata.SocialPost{
		Handle:        handle,
		Title:         momokaData.Get("name").String(),
		Body:          content,
		ContentURI:    contentURI,
		ProfileID:     profileID,
		PublicationID: pubID,
		Media:         lo.Ternary(len(momokaMedia) == 0, nil, momokaMedia),
		Tags:          lo.Ternary(len(momokaTags) == 0, nil, momokaTags),
		Timestamp:     lo.If(isTarget, uint64(0)).Else(timestamp),
	}, nil
}

func (w *worker) getDataFromHTTP(ctx context.Context, contentURL string) (io.ReadCloser, error) {
	// get from ipfs
	if _, path, err := ipfs.ParseURL(contentURL); err == nil {
		resp, err := w.ipfsClient.Fetch(ctx, path, ipfs.FetchModeQuick)
		if err != nil {
			return nil, fmt.Errorf("quick fetch ipfs: %w", err)
		}

		return resp, nil
	}

	// get from arweave
	if strings.HasPrefix(contentURL, "ar://") {
		//	remove ar:// prefix
		contentURL = contentURL[5:]
	} else if strings.HasPrefix(contentURL, "https://arweave.net/") {
		//	 remove https://arweave.net/
		contentURL = contentURL[19:]
	}

	if strings.HasPrefix(contentURL, "https://") {
		return w.httpClient.Fetch(ctx, contentURL)
	}

	return w.arweaveClient.GetTransactionData(ctx, contentURL)
}

func (w *worker) getLensHandle(_ context.Context, blockNumber *big.Int, profileID *big.Int) (string, error) {
	if profileID == nil || profileID.Int64() == 0 {
		return "", nil
	}

	if blockNumber.Int64() < lens.BlockNumberLensV2 {
		profile, err := w.lensHubV1.GetProfile(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
		if err != nil {
			return "", fmt.Errorf("get ethereum profile v1: %w, profile id: %d", err, profileID)
		}

		return profile.Handle, nil
	}

	handleHash, err := w.handleRegistryV2.GetDefaultHandle(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return "", fmt.Errorf("get ethereum default handle v2: %w, profile id: %d", err, profileID)
	}

	name, err := w.lensHandleV2.GetLocalName(&bind.CallOpts{BlockNumber: blockNumber}, handleHash)
	if err != nil {
		return "", fmt.Errorf("get ethereum local name v2: %w, handle hash: %s", err, handleHash)
	}

	return fmt.Sprintf("%s.lens", name), nil
}

func (w *worker) getLensOwnerOf(_ context.Context, blockNumber *big.Int, profileID *big.Int) (*common.Address, error) {
	if blockNumber.Int64() < lens.BlockNumberLensV2 {
		address, err := w.lensHubV1.OwnerOf(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
		if err != nil {
			return nil, fmt.Errorf("get ethereum owner v1 of: %w, profile id: %d", err, profileID)
		}

		return &address, nil
	}

	address, err := w.lensHubV2.OwnerOf(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return nil, fmt.Errorf("get ethereum owner v2 of: %w, profile id: %d", err, profileID)
	}

	return &address, nil
}

// NewWorker returns a new Arweave worker.
func NewWorker(config *config.Module) (engine.Worker, error) {
	var instance = worker{
		config: config,
	}

	var err error

	// Initialize arweave client.
	if instance.arweaveClient, err = arweave.NewClient(); err != nil {
		return nil, fmt.Errorf("new arweave client: %w", err)
	}

	// Initialize ethereum client.
	if instance.ethereumClient, err = ethereum.Dial(context.Background(), config.Endpoint); err != nil {
		return nil, fmt.Errorf("initialize ethereum client: %w", err)
	}

	// Initialize ipfs client.
	if instance.ipfsClient, err = ipfs.NewHTTPClient(); err != nil {
		return nil, fmt.Errorf("new ipfs client: %w", err)
	}

	// Initialize http client.
	if instance.httpClient, err = httpx.NewHTTPClient(); err != nil {
		return nil, fmt.Errorf("new http client: %w", err)
	}

	lensHubV1, err := lens.NewV1LensHubCaller(lens.AddressLensProtocol, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create lens hub v1: %w", err)
	}

	lensHubV2, err := lens.NewV2LensHubCaller(lens.AddressLensProtocol, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create lens hub v2: %w", err)
	}

	lensHandleV2, err := lens.NewV2LensHandleCaller(lens.AddressV2LensHandle, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create lens handle v2: %w", err)
	}

	handleRegistry, err := lens.NewV2HandleRegistryCaller(lens.AddressV2ProfileHandleRegistry, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create handle registry v2: %w", err)
	}

	// Initialize lens filterers.
	instance.lensHubV1 = lensHubV1
	instance.lensHubV2 = lensHubV2
	instance.lensHandleV2 = lensHandleV2
	instance.handleRegistryV2 = handleRegistry

	return &instance, nil
}
