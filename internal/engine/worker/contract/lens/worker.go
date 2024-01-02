package lens

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/arweave"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/lens"
	"github.com/naturalselectionlabs/rss3-node/provider/ipfs"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
)

// Worker is the worker for Lens.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config                         *engine.Config
	arweaveClient                  arweave.Client
	ipfsClient                     ipfs.HTTPClient
	lensHubV1                      *lens.V1LensHub
	lensHubV2                      *lens.V2LensHub
	lensHandleV2                   *lens.V2LensHandle
	handleRegistryV2               *lens.V2HandleRegistry
	eventsFiltererV1               *lens.V1EventsFilterer
	eventsFiltererV2               *lens.V2EventsFilterer
	eventsCollectPublicationAction *lens.V2CollectPublicationActionFilterer
}

func (w *worker) Name() string {
	return engine.Lens.String()
}

// Filter lens contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			lens.AddressLensProtocol,
			lens.AddressV1ProfileCreationProxy,
			lens.AddressV2LensHandle,
			lens.AddressV2ProfileCreationProxy,
			lens.AddressV2ProfileHandleRegistry,
			lens.AddressV2CollectPublicationAction,
			lens.AddressProxyAction[0],
			lens.AddressProxyAction[1],
			lens.AddressProxyAction[2],
			lens.AddressProxyAction[3],
			lens.AddressProxyAction[4],
		},
		LogTopics: []common.Hash{
			lens.EventHashV1PostCreated,
			lens.EventHashV1ProfileCreated,
			lens.EventHashV1CommentCreated,
			lens.EventHashV1MirrorCreated,
			lens.EventHashV1CollectNFTTransferred,
			lens.EventHashV2PostCreated,
			lens.EventHashV2CommentCreated,
			lens.EventHashV2MirrorCreated,
			lens.EventHashV2QuoteCreated,
			lens.EventHashV2Collected,
			lens.EventHashV2ProfileCreated,
		},
	}
}

func (w *worker) Match(_ context.Context, task engine.Task) (bool, error) {
	return task.GetNetwork().Source() == filter.NetworkEthereumSource, nil
}

// Transform Ethereum task to feed.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*schema.Feed, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default lens feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformLens))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions   []*schema.Action
			timestamp uint64
			err       error
		)
		// Match lens core contract events
		switch {
		case w.matchEthereumV1PostCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV1PostCreated(ctx, ethereumTask, log)
		case w.matchEthereumV1CommentCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV1CommentCreated(ctx, ethereumTask, log)
		case w.matchEthereumV1MirrorCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV1MirrorCreated(ctx, ethereumTask, log)
		case w.matchEthereumV1ProfileCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV1ProfileCreated(ctx, ethereumTask, log)
		case w.matchEthereumV1CollectNFTTransferred(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV1CollectNFTTransferred(ctx, ethereumTask, log)
		case w.matchEthereumV2PostCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2PostCreated(ctx, ethereumTask, log)
		case w.matchEthereumV2CommentCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2CommentCreated(ctx, ethereumTask, log)
		case w.matchEthereumV2MirrorCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2MirrorCreated(ctx, ethereumTask, log)
		case w.matchEthereumV2QuoteCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2QuoteCreated(ctx, ethereumTask, log)
		case w.matchEthereumV2Collected(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2Collected(ctx, ethereumTask, log)
		case w.matchEthereumV2ProfileCreated(ethereumTask, log):
			actions, timestamp, err = w.transformEthereumV2ProfileCreated(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		// Change feed type to the first action type.
		for _, action := range actions {
			feed.Type = action.Type
			feed.Timestamp = timestamp
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// matchEthereumV1PostCreated matches V1 PostCreated event.
func (w *worker) matchEthereumV1PostCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV1PostCreated)
}

// matchEthereumV1CommentCreated matches V1 CommentCreated event.
func (w *worker) matchEthereumV1CommentCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV1CommentCreated)
}

// matchEthereumV1MirrorCreated matches V1 MirrorCreated event.
func (w *worker) matchEthereumV1MirrorCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV1MirrorCreated)
}

// matchEthereumV1ProfileCreated matches V1 ProfileCreated event.
func (w *worker) matchEthereumV1ProfileCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV1ProfileCreated)
}

// matchEthereumV1CollectNFTTransferred matches V1 CollectNFTTransferred event.
func (w *worker) matchEthereumV1CollectNFTTransferred(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV1CollectNFTTransferred)
}

// matchEthereumV2PostCreated matches V2 PostCreated event.
func (w *worker) matchEthereumV2PostCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2PostCreated)
}

// matchEthereumV2CommentCreated matches V2 CommentCreated event.
func (w *worker) matchEthereumV2CommentCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2CommentCreated)
}

// matchEthereumV2MirrorCreated matches V2 MirrorCreated event.
func (w *worker) matchEthereumV2MirrorCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2MirrorCreated)
}

// matchEthereumV2QuoteCreated matches V2 QuoteCreated event.
func (w *worker) matchEthereumV2QuoteCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2QuoteCreated)
}

// matchEthereumV2Collected matches V2 Collected event.
func (w *worker) matchEthereumV2Collected(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressV2CollectPublicationAction && len(log.Topics) == 4 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2Collected)
}

// matchEthereumV2ProfileCreated matches V2 ProfileCreated event.
func (w *worker) matchEthereumV2ProfileCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == lens.AddressLensProtocol && len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], lens.EventHashV2ProfileCreated)
}

// transformEthereumV1PostCreated transforms V1 PostCreated event.
func (w *worker) transformEthereumV1PostCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV1.ParsePostCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse post created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	post, platform, err := w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileId, event.PubId, event.ContentURI)
	if err != nil {
		return nil, 0, err
	}

	// Build post created action.
	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), *task.Transaction.To, platform, filter.TypeSocialPost, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV1CommentCreated transforms V1 CommentCreated event.
func (w *worker) transformEthereumV1CommentCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV1.ParseCommentCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse comment created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	actionTo, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.ProfileIdPointed)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of comment
	post, platform, err := w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileId, event.PubId, event.ContentURI)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	pointedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.ProfileIdPointed, event.PubIdPointed)
	if err != nil {
		return nil, 0, err
	}

	post.Target, _, err = w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileIdPointed, event.PubIdPointed, pointedContentURI)
	if err != nil {
		return nil, 0, err
	}

	// Build post created action.
	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), lo.FromPtr(actionTo), platform, filter.TypeSocialComment, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV1MirrorCreated transforms V1 MirrorCreated event.
func (w *worker) transformEthereumV1MirrorCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV1.ParseMirrorCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse mirror created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	//  get the metadata of mirror
	post, _, err := w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileId, event.PubId, "")
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	var platform string

	pointedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.ProfileIdPointed, event.PubIdPointed)
	if err != nil {
		return nil, 0, err
	}

	post.Target, platform, err = w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileIdPointed, event.PubIdPointed, pointedContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), *task.Transaction.To, platform, filter.TypeSocialShare, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV1ProfileCreated transforms V1 ProfileCreated event.
func (w *worker) transformEthereumV1ProfileCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV1.ParseProfileCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse profile created: %w", err)
	}

	profile := metadata.SocialProfile{
		Action:    metadata.ActionSocialProfileCreate,
		ProfileID: EncodeID(event.ProfileId),
		Address:   event.To,
		Handle:    event.Handle,
		ImageURI:  event.ImageURI,
	}

	action := w.buildEthereumTransactionProfileAction(ctx, event.Creator, event.To, filter.TypeSocialProfile, profile)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV1CollectNFTTransferred transforms V1 CollectNFTTransferred event.
func (w *worker) transformEthereumV1CollectNFTTransferred(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	var from common.Address

	event, err := w.eventsFiltererV1.ParseCollectNFTTransferred(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse collect nft transferred: %w", err)
	}

	if event.From != ethereum.AddressGenesis {
		return nil, 0, nil
	}

	// get the metadata of the pointed post/comment
	mintContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.ProfileId, event.PubId)
	if err != nil {
		return nil, 0, err
	}

	post, platform, err := w.buildEthereumV1TransactionPostMetadata(ctx, log.BlockNumber, event.ProfileId, event.PubId, mintContentURI)
	if err != nil {
		return nil, 0, err
	}

	from = event.To

	for _, log := range task.Receipt.Logs {
		if log.Topics[0] != lens.EventHashV1CollectNFTTransferred {
			continue
		}

		event, err := w.eventsFiltererV1.ParseCollectNFTTransferred(log.Export())
		if err != nil {
			return nil, 0, fmt.Errorf("parse collect nft transferred: %w", err)
		}

		if event.From == ethereum.AddressGenesis {
			continue
		}

		if event.From == from {
			from = event.To
		}
	}

	action := w.buildEthereumTransactionPostAction(ctx, from, *task.Transaction.To, platform, filter.TypeSocialMint, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2PostCreated transforms V2 PostCreated event.
func (w *worker) transformEthereumV2PostCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV2.ParsePostCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse post created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.PostParams.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	post, platform, err := w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.PostParams.ProfileId, event.PubId, event.PostParams.ContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), *task.Transaction.To, platform, filter.TypeSocialPost, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2CommentCreated transforms V2 CommentCreated event.
func (w *worker) transformEthereumV2CommentCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV2.ParseCommentCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse comment created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.CommentParams.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	actionTo, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.CommentParams.PointedProfileId)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of comment
	post, platform, err := w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.CommentParams.ProfileId, event.PubId, event.CommentParams.ContentURI)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	pointedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.CommentParams.PointedProfileId, event.CommentParams.PointedPubId)
	if err != nil {
		return nil, 0, err
	}

	post.Target, _, err = w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.CommentParams.PointedProfileId, event.CommentParams.PointedPubId, pointedContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), lo.FromPtr(actionTo), platform, filter.TypeSocialComment, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2MirrorCreated transforms V2 MirrorCreated event.
func (w *worker) transformEthereumV2MirrorCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV2.ParseMirrorCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse mirror created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.MirrorParams.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	//  get the metadata of mirror
	post, _, err := w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.MirrorParams.ProfileId, event.PubId, "")
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	var platform string

	pointedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.MirrorParams.PointedProfileId, event.MirrorParams.PointedPubId)
	if err != nil {
		return nil, 0, err
	}

	post.Target, platform, err = w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.MirrorParams.PointedProfileId, event.MirrorParams.PointedPubId, pointedContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), *task.Transaction.To, platform, filter.TypeSocialShare, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2QuoteCreated transforms V2 QuoteCreated event.
func (w *worker) transformEthereumV2QuoteCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV2.ParseQuoteCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse quote created: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.QuoteParams.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	//  get the metadata of mirror
	post, platform, err := w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.QuoteParams.ProfileId, event.PubId, event.QuoteParams.ContentURI)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the pointed post/comment
	pointedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.QuoteParams.PointedProfileId, event.QuoteParams.PointedPubId)
	if err != nil {
		return nil, 0, err
	}

	post.Target, _, err = w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.QuoteParams.PointedProfileId, event.QuoteParams.PointedPubId, pointedContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), *task.Transaction.To, platform, filter.TypeSocialShare, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2Collected transforms V2 Collected event.
func (w *worker) transformEthereumV2Collected(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsCollectPublicationAction.ParseCollected(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse collected: %w", err)
	}

	actionFrom, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.CollectorProfileId)
	if err != nil {
		return nil, 0, err
	}

	actionTo, err := w.getLensOwnerOf(ctx, log.BlockNumber, event.CollectedProfileId)
	if err != nil {
		return nil, 0, err
	}

	// get the metadata of the collected post/comment
	collectedContentURI, err := w.getEthereumPublicationContentURI(ctx, log.BlockNumber, event.CollectedProfileId, event.CollectedPubId)
	if err != nil {
		return nil, 0, err
	}

	post, platform, err := w.buildEthereumV2TransactionPostMetadata(ctx, log.BlockNumber, event.CollectedProfileId, event.CollectedPubId, collectedContentURI)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionPostAction(ctx, lo.FromPtr(actionFrom), lo.FromPtr(actionTo), platform, filter.TypeSocialMint, *post)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

// transformEthereumV2ProfileCreated transforms V2 ProfileCreated event.
func (w *worker) transformEthereumV2ProfileCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, uint64, error) {
	event, err := w.eventsFiltererV2.ParseProfileCreated(log.Export())
	if err != nil {
		return nil, 0, fmt.Errorf("parse profile created: %w", err)
	}

	profile := metadata.SocialProfile{
		Action:    metadata.ActionSocialProfileCreate,
		ProfileID: EncodeID(event.ProfileId),
		Address:   event.To,
	}

	profile.Handle, err = w.getLensHandle(ctx, log.BlockNumber, event.ProfileId)
	if err != nil {
		return nil, 0, err
	}

	action := w.buildEthereumTransactionProfileAction(ctx, event.Creator, event.To, filter.TypeSocialProfile, profile)

	return []*schema.Action{
		action,
	}, event.Timestamp.Uint64(), nil
}

func (w *worker) buildEthereumTransactionPostAction(_ context.Context, from common.Address, to common.Address, platform string, socialType filter.Type, post metadata.SocialPost) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       lo.If(to == ethereum.AddressGenesis, "").Else(to.String()),
		Platform: platform,
		Type:     socialType,
		Metadata: post,
	}
}

func (w *worker) buildEthereumV1TransactionPostMetadata(ctx context.Context, blockNumber *big.Int, profileID, pubID *big.Int, contentURI string) (*metadata.SocialPost, string, error) {
	handle, err := w.getLensHandle(ctx, blockNumber, profileID)
	if err != nil {
		return nil, "", err
	}

	content, err := w.getEthereumPublication(ctx, contentURI)
	if err != nil {
		return nil, "", err
	}

	var publication PublicationV1
	if err = json.Unmarshal(content, &publication); err != nil {
		return nil, "", fmt.Errorf("unmarshal publication: %w", err)
	}

	return &metadata.SocialPost{
		Handle: handle,
		Body:   publication.Content,
		Media: lo.Map(publication.Media, func(media PublicationMedia, index int) metadata.Media {
			return metadata.Media{
				MimeType: media.Type,
				Address:  media.Item,
			}
		}),
		ProfileID:     EncodeID(profileID),
		PublicationID: EncodeID(pubID),
		ContentURI:    contentURI,
		Tags:          lo.If(len(publication.Tags) > 0, publication.Tags).Else(nil),
	}, publication.AppID, nil
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

func (w *worker) getEthereumPublication(ctx context.Context, contentURI string) (json.RawMessage, error) {
	if len(contentURI) == 0 {
		return []byte("{}"), nil
	}

	body, err := w.getPublicationFromHTTP(ctx, contentURI)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("read all body: %w", err)
	}

	return content, nil
}

func (w *worker) getPublicationFromHTTP(ctx context.Context, contentURL string) (io.ReadCloser, error) {
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

	// http request
	return w.arweaveClient.GetTransactionData(ctx, contentURL)
}

// getEthereumPublicationContentURI gets content uri from ethereum.
func (w *worker) getEthereumPublicationContentURI(_ context.Context, blockNumber *big.Int, profileID, pubID *big.Int) (string, error) {
	contentURI, err := w.lensHubV1.GetContentURI(&bind.CallOpts{BlockNumber: blockNumber}, profileID, pubID)
	if err != nil {
		return "", fmt.Errorf("get ethereum content uri v1: %w, profile id: %d, pub id: %d", err, profileID, pubID)
	}

	return contentURI, nil
}

// buildEthereumTransactionProfileAction builds profile action.
func (w *worker) buildEthereumTransactionProfileAction(_ context.Context, from common.Address, to common.Address, socialType filter.Type, profile metadata.SocialProfile) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       lo.If(to == ethereum.AddressGenesis, "").Else(to.String()),
		Platform: filter.PlatformLens.String(),
		Type:     socialType,
		Metadata: profile,
	}
}

// buildEthereumV2TransactionPostMetadata builds post metadata.
func (w *worker) buildEthereumV2TransactionPostMetadata(ctx context.Context, blockNumber *big.Int, profileID, pubID *big.Int, contentURI string) (*metadata.SocialPost, string, error) {
	handle, err := w.getLensHandle(ctx, blockNumber, profileID)
	if err != nil {
		return nil, "", err
	}

	content, err := w.getEthereumPublication(ctx, contentURI)
	if err != nil {
		return nil, "", err
	}

	var publication PublicationV2
	if err = json.Unmarshal(content, &publication); err != nil {
		return nil, "", fmt.Errorf("unmarshal publication: %w", err)
	}

	body := publication.Content
	if len(publication.Content) == 0 {
		body = publication.Lens.Content
	}

	media := make([]metadata.Media, 0)

	if len(publication.Lens.Image.Item) > 0 {
		media = append(media, metadata.Media{
			MimeType: publication.Lens.Image.Type,
			Address:  publication.Lens.Image.Item,
		})
	}

	if len(publication.Lens.Video.Item) > 0 {
		media = append(media, metadata.Media{
			MimeType: publication.Lens.Video.Type,
			Address:  publication.Lens.Video.Item,
		})
	}

	return &metadata.SocialPost{
		Handle:        handle,
		Body:          body,
		Title:         publication.Title,
		Media:         media,
		ProfileID:     EncodeID(profileID),
		PublicationID: EncodeID(pubID),
		ContentURI:    contentURI,
		Tags:          lo.If(len(publication.Lens.Tags) > 0, publication.Lens.Tags).Else(nil),
	}, publication.Lens.AppID, nil
}

// NewWorker creates a new Lens worker.
func NewWorker(config *engine.Config) (engine.Worker, error) {
	var (
		err      error
		instance = worker{
			config: config,
		}
	)

	// Initialize arweave client.
	if instance.arweaveClient, err = arweave.NewClient(); err != nil {
		return nil, fmt.Errorf("new arweave client: %w", err)
	}

	// Initialize ipfs client.
	if instance.ipfsClient, err = ipfs.NewHTTPClient(); err != nil {
		return nil, fmt.Errorf("new ipfs client: %w", err)
	}

	// Initialize lens filterers.
	instance.lensHubV1 = lo.Must(lens.NewV1LensHub(ethereum.AddressGenesis, nil))
	instance.lensHubV2 = lo.Must(lens.NewV2LensHub(ethereum.AddressGenesis, nil))
	instance.lensHandleV2 = lo.Must(lens.NewV2LensHandle(ethereum.AddressGenesis, nil))
	instance.handleRegistryV2 = lo.Must(lens.NewV2HandleRegistry(ethereum.AddressGenesis, nil))
	instance.eventsFiltererV1 = lo.Must(lens.NewV1EventsFilterer(ethereum.AddressGenesis, nil))
	instance.eventsFiltererV2 = lo.Must(lens.NewV2EventsFilterer(ethereum.AddressGenesis, nil))
	instance.eventsCollectPublicationAction = lo.Must(lens.NewV2CollectPublicationActionFilterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
