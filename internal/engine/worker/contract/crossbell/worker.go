package crossbell

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell/character"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell/event"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell/periphery"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell/profile"
	"github.com/rss3-network/node/provider/ethereum/contract/crossbell/tips"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/endpoint"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/provider/ipfs"
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/filter"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// Worker is the worker for Crossbell.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	config            *config.Module
	ethereumClient    ethereum.Client
	ipfsClient        ipfs.HTTPClient
	tokenClient       token.Client
	eventFilterer     *event.EventFilterer
	profileFilterer   *profile.ProfileFilterer
	tipsFilterer      *tips.TipsFilterer
	erc20Filterer     *erc20.ERC20Filterer
	characterContract *character.CharacterCaller
	profileContract   *profile.ProfileCaller
	peripheryContract *periphery.PeripheryCaller
}

func (w *worker) Name() string {
	return filter.Crossbell.String()
}

// Filter crossbell contract address and event hash.
func (w *worker) Filter() engine.SourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			crossbell.AddressWeb3Entry,
			crossbell.AddressPeriphery,
			crossbell.AddressTips,
		},
		LogTopics: []common.Hash{
			crossbell.EventHashProfileCreated,
			crossbell.EventHashSetProfileURI,
			crossbell.EventHashCharacterCreated,
			crossbell.EventHashSetHandle,
			crossbell.EventHashSetCharacterURI,
			crossbell.EventHashPostNote,
			crossbell.EventHashSetNoteURI,
			crossbell.EventHashDeleteNote,
			crossbell.EventHashMintNote,
			crossbell.EventHashSetOperator,
			crossbell.EventHashAddOperator,
			crossbell.EventHashRemoveOperator,
			crossbell.EventHashGrantOperatorPermissions,
			crossbell.EventHashTipCharacterForNote,
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

	// Build default crossbell feed from task.
	feed, err := ethereumTask.BuildFeed(schema.WithFeedPlatform(filter.PlatformCrossbell))
	if err != nil {
		return nil, fmt.Errorf("build feed: %w", err)
	}

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		// Ignore anonymous logs.
		if len(log.Topics) == 0 {
			continue
		}

		var (
			actions []*schema.Action
			err     error
		)
		// Match crossbell core contract events
		switch {
		case w.matchProfileCreated(ethereumTask, log):
			actions, err = w.transformProfileCreated(ctx, ethereumTask, log)
		case w.matchSetProfileURI(ethereumTask, log):
			actions, err = w.transformSetProfileURI(ctx, ethereumTask, log)
		case w.matchCharacterCreated(ethereumTask, log):
			actions, err = w.transformCharacterCreated(ctx, ethereumTask, log)
		case w.matchCharacterSetHandle(ethereumTask, log):
			actions, err = w.transformCharacterSetHandle(ctx, ethereumTask, log)
		case w.matchSetCharacterURI(ethereumTask, log):
			actions, err = w.transformSetCharacterURI(ctx, ethereumTask, log)
		case w.matchPostCreated(ethereumTask, log):
			actions, err = w.transformPostCreated(ctx, ethereumTask, log)
		case w.matchSetNoteURI(ethereumTask, log):
			actions, err = w.transformSetNoteURI(ctx, ethereumTask, log)
		case w.matchDeleteNote(ethereumTask, log):
			actions, err = w.transformDeleteNote(ctx, ethereumTask, log)
		case w.matchMintNote(ethereumTask, log):
			actions, err = w.transformMintNote(ctx, ethereumTask, log)
		case w.matchSetOperator(ethereumTask, log):
			actions, err = w.transformSetOperator(ctx, ethereumTask, log)
		case w.matchAddOperator(ethereumTask, log):
			actions, err = w.transformAddOperator(ctx, ethereumTask, log)
		case w.matchRemoveOperator(ethereumTask, log):
			actions, err = w.transformRemoveOperator(ctx, ethereumTask, log)
		case w.matchGrantOperatorPermissions(ethereumTask, log):
			actions, err = w.transformGrantOperatorPermissions(ctx, ethereumTask, log)
		case w.matchTipsCharacterForNote(ethereumTask, log):
			actions, err = w.transformTipsCharacterForNote(ctx, ethereumTask, log)
		default:
			continue
		}

		if err != nil {
			return nil, err
		}

		// Change feed type to the first action type.
		for _, action := range actions {
			feed.Type = action.Type
		}

		feed.Actions = append(feed.Actions, actions...)
	}

	return feed, nil
}

// matchProfileCreated matches ProfileCreated event.
func (w *worker) matchProfileCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashProfileCreated)
}

// matchSetProfileURI matches SetProfileURI event.
func (w *worker) matchSetProfileURI(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashSetProfileURI)
}

// matchCharacterCreated matches CharacterCreated event.
func (w *worker) matchCharacterCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashCharacterCreated)
}

// matchCharacterSetHandle matches SetCharacterHandle event.
func (w *worker) matchCharacterSetHandle(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashSetHandle)
}

// matchCSetCharacterURI matches SetCharacterURI event.
func (w *worker) matchSetCharacterURI(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashSetCharacterURI)
}

// matchPostCreated matches PostCreated event.
func (w *worker) matchPostCreated(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashPostNote)
}

// matchSetNoteURI matches SetNoteURI event.
func (w *worker) matchSetNoteURI(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashSetNoteURI)
}

// matchDeleteNote matches DeleteNote event.
func (w *worker) matchDeleteNote(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashDeleteNote)
}

// matchMintNote matches MintNote event.
func (w *worker) matchMintNote(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashMintNote)
}

// matchSetOperator matches SetOperator event.
func (w *worker) matchSetOperator(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashSetOperator)
}

// matchAddOperator matches AddOperator event.
func (w *worker) matchAddOperator(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashAddOperator)
}

// matchRemoveOperator matches RemoveOperator event.
func (w *worker) matchRemoveOperator(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashRemoveOperator)
}

// matchGrantOperatorPermissions matches GrantOperatorPermissions event.
func (w *worker) matchGrantOperatorPermissions(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressWeb3Entry && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashGrantOperatorPermissions)
}

// matchTipsCharacterForNote matches TipsCharacterForNote event.
func (w *worker) matchTipsCharacterForNote(_ *source.Task, log *ethereum.Log) bool {
	return log.Address == crossbell.AddressTips && contract.MatchEventHashes(log.Topics[0], crossbell.EventHashTipCharacterForNote)
}

// matchERC20TransferLog matches ERC20 Transfer log.
func (w *worker) matchERC20TransferLog(_ *source.Task, log *ethereum.Log) bool {
	return len(log.Topics) == 3 && contract.MatchEventHashes(log.Topics[0], erc20.EventHashTransfer)
}

// transformProfileCreated transforms ProfileCreated event.
func (w *worker) transformProfileCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.profileFilterer.ParseProfileCreated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse profile created: %w", err)
	}

	profile, err := w.buildProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileCreate, event.ProfileId, event.To, event.Handle, "")
	if err != nil {
		return nil, err
	}

	action := w.buildProfileAction(ctx, event.Creator, event.To, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// transformSetProfileURI transforms SetProfileURI event.
func (w *worker) transformSetProfileURI(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.profileFilterer.ParseSetProfileUri(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse profile uri set: %w", err)
	}

	profile, err := w.buildProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileUpdate, event.ProfileId, common.Address{}, "", event.NewUri)
	if err != nil {
		return nil, err
	}

	action := w.buildProfileAction(ctx, profile.Address, *task.Transaction.To, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// transformCharacterCreated transforms CharacterCreated event.
func (w *worker) transformCharacterCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseCharacterCreated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse character created: %w", err)
	}

	profile, err := w.buildCharacterProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileCreate, event.CharacterId, event.To, event.Handle, "")
	if err != nil {
		return nil, err
	}

	action := w.buildProfileAction(ctx, event.Creator, profile.Address, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// transformCharacterSetHandle transforms SetCharacterHandle event.
func (w *worker) transformCharacterSetHandle(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseSetHandle(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse character handle set: %w", err)
	}

	profile, err := w.buildCharacterProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileUpdate, event.CharacterId, event.Account, event.NewHandle, "")
	if err != nil {
		return nil, err
	}

	action := w.buildProfileAction(ctx, profile.Address, *task.Transaction.To, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// transformSetCharacterURI transforms SetCharacterURI event.
func (w *worker) transformSetCharacterURI(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseSetCharacterUri(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse character uri set: %w", err)
	}

	profile, err := w.buildCharacterProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileUpdate, event.CharacterId, common.Address{}, "", event.NewUri)
	if err != nil {
		return nil, err
	}

	action := w.buildProfileAction(ctx, profile.Address, *task.Transaction.To, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// transformPostCreated transforms PostCreated event.
func (w *worker) transformPostCreated(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	feedType := filter.TypeSocialPost

	event, err := w.eventFilterer.ParsePostNote(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse post note: %w", err)
	}

	fromAddress, err := w.getOwnerOf(ctx, log.BlockNumber, event.CharacterId)
	if err != nil {
		return nil, err
	}

	post, linkKey, platform, err := w.buildPostMetadata(ctx, log.BlockNumber, event.CharacterId, event.NoteId, "")
	if err != nil {
		return nil, err
	}

	targetCharacterID, targetNoteID, err := w.getLinkingNote(ctx, log.BlockNumber, linkKey)
	if err != nil {
		return nil, err
	}

	if targetCharacterID != nil && targetCharacterID.Int64() != 0 {
		feedType = filter.TypeSocialComment

		post.Target, _, _, err = w.buildPostMetadata(ctx, log.BlockNumber, targetCharacterID, targetNoteID, "")
		if err != nil {
			return nil, err
		}
	}

	action := w.buildPostAction(ctx, fromAddress, *task.Transaction.To, platform, feedType, *post)

	return []*schema.Action{
		action,
	}, nil
}

// transformSetNoteURI transforms SetNoteURI event.
func (w *worker) transformSetNoteURI(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseSetNoteUri(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse set note uri: %w", err)
	}

	fromAddress, err := w.getOwnerOf(ctx, log.BlockNumber, event.CharacterId)
	if err != nil {
		return nil, err
	}

	post, _, platform, err := w.buildPostMetadata(ctx, log.BlockNumber, event.CharacterId, event.NoteId, event.NewUri)
	if err != nil {
		return nil, err
	}

	action := w.buildPostAction(ctx, fromAddress, *task.Transaction.To, platform, filter.TypeSocialRevise, *post)

	return []*schema.Action{
		action,
	}, nil
}

// transformDeleteNote transforms DeleteNote event.
func (w *worker) transformDeleteNote(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseDeleteNote(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse delete note: %w", err)
	}

	fromAddress, err := w.getOwnerOf(ctx, log.BlockNumber, event.CharacterId)
	if err != nil {
		return nil, err
	}

	post, _, platform, err := w.buildPostMetadata(ctx, log.BlockNumber, event.CharacterId, event.NoteId, "")
	if err != nil {
		return nil, err
	}

	action := w.buildPostAction(ctx, fromAddress, *task.Transaction.To, platform, filter.TypeSocialDelete, *post)

	return []*schema.Action{
		action,
	}, nil
}

// transformMintNote transforms MintNote event.
func (w *worker) transformMintNote(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseMintNote(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse mint note: %w", err)
	}

	post, _, platform, err := w.buildPostMetadata(ctx, log.BlockNumber, event.CharacterId, event.NoteId, "")
	if err != nil {
		return nil, err
	}

	action := w.buildPostAction(ctx, event.To, *task.Transaction.To, platform, filter.TypeSocialMint, *post)

	return []*schema.Action{
		action,
	}, nil
}

// transformTipsCharacterForNote transforms TipsCharacterForNote event.
func (w *worker) transformTipsCharacterForNote(ctx context.Context, task *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.tipsFilterer.ParseTipCharacterForNote(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse tips character for note: %w", err)
	}

	post, _, platform, err := w.buildPostMetadata(ctx, log.BlockNumber, event.ToCharacterId, event.ToNoteId, "")
	if err != nil {
		return nil, err
	}

	rewardTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, &event.Token, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", event.Token, err)
	}

	rewardTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(event.Amount, 0))

	post.Reward = rewardTokenMetadata

	var to common.Address

	for _, log := range task.Receipt.Logs {
		if len(log.Topics) == 0 {
			continue
		}

		switch {
		case w.matchERC20TransferLog(task, log):
			transferEvent, err := w.erc20Filterer.ParseTransfer(log.Export())
			if err != nil {
				return nil, fmt.Errorf("parse OrderFulfilled event: %w", err)
			}

			if transferEvent.From == crossbell.AddressTips {
				to = transferEvent.To
			}

		default:
			continue
		}

		if err != nil {
			zap.L().Debug("handle ethereum log", zap.Error(err), zap.String("task", task.ID()))
			continue
		}
	}

	action := w.buildPostRewardAction(ctx, task.Transaction.From, to, platform, filter.TypeSocialReward, *post)

	return []*schema.Action{
		action,
	}, nil
}

// transformSetOperator transforms SetOperator event.
func (w *worker) transformSetOperator(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseSetOperator(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse set operator: %w", err)
	}

	actionType := lo.If(event.Operator == ethereum.AddressGenesis, metadata.ActionSocialProxyRemove).Else(metadata.ActionSocialProxyAppoint)

	proxy, err := w.buildProxyMetadata(ctx, log.BlockNumber, lo.ToPtr(actionType), event.CharacterId, lo.ToPtr(event.Operator))
	if err != nil {
		return nil, err
	}

	action := w.buildProxyAction(ctx, proxy.Profile.Address, event.Operator, filter.TypeSocialProxy, *proxy)

	return []*schema.Action{
		action,
	}, nil
}

// transformAddOperator transforms AddOperator event.
func (w *worker) transformAddOperator(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseAddOperator(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse add operator: %w", err)
	}

	proxy, err := w.buildProxyMetadata(ctx, log.BlockNumber, lo.ToPtr(metadata.ActionSocialProxyAppoint), event.CharacterId, lo.ToPtr(event.Operator))
	if err != nil {
		return nil, err
	}

	action := w.buildProxyAction(ctx, proxy.Profile.Address, event.Operator, filter.TypeSocialProxy, *proxy)

	return []*schema.Action{
		action,
	}, nil
}

// transformRemoveOperator transforms RemoveOperator event.
func (w *worker) transformRemoveOperator(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseRemoveOperator(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse remove operator: %w", err)
	}

	proxy, err := w.buildProxyMetadata(ctx, log.BlockNumber, lo.ToPtr(metadata.ActionSocialProxyRemove), event.CharacterId, lo.ToPtr(event.Operator))
	if err != nil {
		return nil, err
	}

	action := w.buildProxyAction(ctx, proxy.Profile.Address, event.Operator, filter.TypeSocialProxy, *proxy)

	return []*schema.Action{
		action,
	}, nil
}

// transformGrantOperatorPermissions transforms GrantOperatorPermissions event.
func (w *worker) transformGrantOperatorPermissions(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.eventFilterer.ParseGrantOperatorPermissions(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse grant operator permissions: %w", err)
	}

	actionType := lo.If(lo.IsEmpty(event.PermissionBitMap.BitLen()), metadata.ActionSocialProxyRemove).Else(metadata.ActionSocialProxyAppoint)

	proxy, err := w.buildProxyMetadata(ctx, log.BlockNumber, lo.ToPtr(actionType), event.CharacterId, lo.ToPtr(event.Operator))
	if err != nil {
		return nil, err
	}

	action := w.buildProxyAction(ctx, proxy.Profile.Address, event.Operator, filter.TypeSocialProxy, *proxy)

	return []*schema.Action{
		action,
	}, nil
}

// buildProfileAction builds profile action.
func (w *worker) buildProfileAction(_ context.Context, from, to common.Address, actionType filter.Type, profile metadata.SocialProfile) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       to.String(),
		Platform: filter.PlatformCrossbell.String(),
		Type:     actionType,
		Metadata: profile,
	}
}

// builPostAction builds post action.
func (w *worker) buildPostAction(_ context.Context, from common.Address, to common.Address, platform string, actionType filter.Type, post metadata.SocialPost) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       to.String(),
		Platform: platform,
		Type:     actionType,
		Metadata: post,
	}
}

// buildPostRewardAction builds post reward action.
func (w *worker) buildPostRewardAction(_ context.Context, from common.Address, to common.Address, platform string, actionType filter.Type, post metadata.SocialPost) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       to.String(),
		Platform: platform,
		Type:     actionType,
		Metadata: &metadata.SocialPost{
			Target: &post,
		},
	}
}

// buildPostMetadata builds post metadata.
func (w *worker) buildPostMetadata(ctx context.Context, blockNumber, characterID, noteID *big.Int, contentURI string) (*metadata.SocialPost, common.Hash, string, error) {
	handle, err := w.getHandle(ctx, blockNumber, characterID)
	if err != nil {
		return nil, common.Hash{}, "", err
	}

	var linkKey common.Hash

	if contentURI == "" {
		contentURI, linkKey, err = w.getNote(ctx, blockNumber, characterID, noteID)
		if err != nil {
			return nil, common.Hash{}, "", err
		}
	}

	content, err := w.getIPFSContent(ctx, contentURI)
	if err != nil {
		return nil, common.Hash{}, "", err
	}

	var note NoteContent
	if err = json.Unmarshal(content, &note); err != nil {
		return nil, common.Hash{}, "", fmt.Errorf("unmarshal note content: %w", err)
	}

	platform := w.getNotePlatform(ctx, note.Sources)

	filteredMedia := []metadata.Media{}

	for _, attachment := range note.Attachments {
		if attachment.Address != "" {
			filteredMedia = append(filteredMedia, metadata.Media{
				MimeType: attachment.MimeType,
				Address:  attachment.Address,
			})
		}
	}

	return &metadata.SocialPost{
		ProfileID:     characterID.String(),
		PublicationID: noteID.String(),
		Handle:        w.buildProfileHandleSuffix(ctx, handle),
		Title:         note.Title,
		Body:          note.Content,
		Media:         filteredMedia,
		ContentURI:    contentURI,
	}, linkKey, platform, nil
}

// buildProxyAction builds proxy action.
func (w *worker) buildProxyAction(_ context.Context, from common.Address, to common.Address, actionType filter.Type, proxy metadata.SocialProxy) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       to.String(),
		Platform: filter.PlatformCrossbell.String(),
		Type:     actionType,
		Metadata: proxy,
	}
}

// buildProfileMetadata builds profile metadata.
func (w *worker) buildProfileMetadata(
	ctx context.Context,
	blockNumber *big.Int,
	action metadata.SocialProfileAction,
	profileID *big.Int,
	address common.Address,
	handle string,
	uri string,
) (*metadata.SocialProfile, error) {
	profile := &metadata.SocialProfile{
		Action:    action,
		ProfileID: profileID.String(),
		Address:   address,
		Handle:    handle,
	}

	if lo.IsEmpty(address) {
		var err error

		profile.Address, err = w.getOwnerOf(ctx, blockNumber, profileID)
		if err != nil {
			return nil, err
		}
	}

	if lo.IsEmpty(handle) || lo.IsEmpty(uri) {
		profileHandle, profileURI, err := w.getProfile(ctx, blockNumber, profileID)
		if err != nil {
			return nil, err
		}

		profile.Handle = lo.If(lo.IsEmpty(handle), profileHandle).Else(handle)
		uri = lo.If(lo.IsEmpty(uri), profileURI).Else(uri)
	}

	profile.Handle = w.buildProfileHandleSuffix(ctx, profile.Handle)

	if lo.IsNotEmpty(uri) {
		content, err := w.getIPFSContent(ctx, uri)
		if err != nil {
			return nil, err
		}

		mimeType := http.DetectContentType(content)
		if strings.HasPrefix(mimeType, "image") {
			profile.ImageURI = uri

			return profile, nil
		}

		var profileURI ProfileURIContent
		if err = json.Unmarshal(content, &profileURI); err != nil {
			return nil, fmt.Errorf("unmarshal profile uri content: %w", err)
		}

		if len(profileURI.ConnectedAvatars) > 0 {
			profile.ImageURI = profileURI.ConnectedAvatars[0]
		}

		profile.Bio = profileURI.Bio
		profile.Name = profileURI.Name
	}

	if strings.Contains(profile.ImageURI, "csb://asset:") {
		profile.ImageURI, _ = w.getAssetImageURI(ctx, profile.ImageURI)
	}

	return profile, nil
}

// getAssetImageURI gets asset image uri.
func (w *worker) getAssetImageURI(ctx context.Context, assetURI string) (string, error) {
	asset := strings.FieldsFunc(strings.TrimPrefix(assetURI, "csb://asset:"), func(r rune) bool {
		return r == '-' || r == '@'
	})

	if len(asset) != 3 {
		return "", fmt.Errorf("invalid asset uri: %s", assetURI)
	}

	network, chainID := filter.NetworkAndChainID(asset[2])
	if chainID <= 0 {
		return "", fmt.Errorf("invalid chain id: %s", asset[2])
	}

	tokenID, err := decimal.NewFromString(asset[1])
	if err != nil {
		return "", fmt.Errorf("invalid token id: %s", asset[1])
	}

	var tokenClient token.Client

	if chainID != filter.EthereumChainIDCrossbell {
		var ethereumClient ethereum.Client

		// Initialize ethereum client.
		endpoints, exists := endpoint.Get(network)
		if !exists {
			return "", fmt.Errorf("get endpoint: %w", err)
		}

		for _, endpoint := range endpoints {
			ethereumClient, err = ethereum.Dial(context.Background(), endpoint)
			if err == nil {
				break
			}
		}

		if ethereumClient == nil {
			return "", fmt.Errorf("initialize ethereum client: %w", err)
		}

		// Initialize token client.
		tokenClient = token.NewClient(ethereumClient)
	} else {
		tokenClient = w.tokenClient
	}

	// parse token metadata
	tokenMetadata, err := tokenClient.Lookup(ctx, uint64(chainID), lo.ToPtr(common.HexToAddress(asset[0])), tokenID.BigInt(), nil, token.WithParseTokenMetadata(true))
	if err != nil {
		return "", fmt.Errorf("lookup token metadata %s: %w", asset[0], err)
	}

	return tokenMetadata.ParsedImageURL, nil
}

// buildProxyMetadata
func (w *worker) buildProxyMetadata(
	ctx context.Context,
	blockNumber *big.Int,
	action *metadata.SocialProxyAction,
	characterID *big.Int,
	proxyAddress *common.Address,
) (*metadata.SocialProxy, error) {
	proxy := &metadata.SocialProxy{
		Action:       *action,
		ProxyAddress: *proxyAddress,
		Profile: metadata.SocialProfile{
			ProfileID: characterID.String(),
		},
	}

	var err error

	proxy.Profile.Address, err = w.getOwnerOf(ctx, blockNumber, characterID)
	if err != nil {
		return nil, err
	}

	proxy.Profile.Handle, err = w.getHandle(ctx, blockNumber, characterID)
	if err != nil {
		return nil, err
	}

	proxy.Profile.Handle = w.buildProfileHandleSuffix(ctx, proxy.Profile.Handle)

	return proxy, nil
}

// buildCharacterProfileMetadata builds character profile metadata.
func (w *worker) buildCharacterProfileMetadata(
	ctx context.Context,
	blockNumber *big.Int,
	action metadata.SocialProfileAction,
	characterID *big.Int,
	address common.Address,
	handle string,
	uri string,
) (*metadata.SocialProfile, error) {
	profile := &metadata.SocialProfile{
		Action:    action,
		ProfileID: characterID.String(),
		Address:   address,
		Handle:    handle,
	}

	if lo.IsEmpty(address) {
		var err error

		profile.Address, err = w.getOwnerOf(ctx, blockNumber, characterID)
		if err != nil {
			return nil, err
		}
	}

	if lo.IsEmpty(handle) || lo.IsEmpty(uri) {
		characterHandle, characterURI, err := w.getCharacter(ctx, blockNumber, characterID)
		if err != nil {
			return nil, err
		}

		profile.Handle = lo.If(lo.IsEmpty(handle), characterHandle).Else(handle)
		uri = lo.If(lo.IsEmpty(uri), characterURI).Else(uri)
	}

	if !lo.IsEmpty(uri) {
		content, err := w.getIPFSContent(ctx, uri)
		if err != nil {
			return nil, err
		}

		var characterURI CharacterURIContent
		if err = json.Unmarshal(content, &characterURI); err != nil {
			return nil, fmt.Errorf("unmarshal character uri content: %w", err)
		}

		if len(characterURI.Avatars) > 0 {
			profile.ImageURI = characterURI.Avatars[0]
		}

		profile.Bio = characterURI.Bio
		profile.Name = characterURI.Name
	}

	if strings.Contains(profile.ImageURI, "csb://asset:") {
		profile.ImageURI, _ = w.getAssetImageURI(ctx, profile.ImageURI)
	}

	profile.Handle = w.buildProfileHandleSuffix(ctx, profile.Handle)

	return profile, nil
}

// getHandle gets handle of profile.
func (w *worker) getHandle(_ context.Context, blockNumber, characterID *big.Int) (string, error) {
	handle, err := w.characterContract.GetHandle(&bind.CallOpts{BlockNumber: blockNumber}, characterID)
	if err != nil {
		return "", fmt.Errorf("get handle: %w", err)
	}

	return handle, nil
}

// getLinkingNote gets linking note of profile.
func (w *worker) getLinkingNote(_ context.Context, blockNumber *big.Int, linkKey common.Hash) (*big.Int, *big.Int, error) {
	linking, err := w.peripheryContract.GetLinkingNote(&bind.CallOpts{BlockNumber: blockNumber}, linkKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get linking note: %w", err)
	}

	return linking.CharacterId, linking.NoteId, nil
}

// getNote gets note of profile.
func (w *worker) getNote(_ context.Context, blockNumber, characterID, noteID *big.Int) (string, common.Hash, error) {
	note, err := w.characterContract.GetNote(&bind.CallOpts{BlockNumber: blockNumber}, characterID, noteID)
	if err != nil {
		return "", common.Hash{}, fmt.Errorf("get note: %w", err)
	}

	return note.ContentUri, note.LinkKey, nil
}

// getNotePlatform gets note platform of profile.
func (w *worker) getNotePlatform(_ context.Context, sources []string) string {
	filtered := lo.Filter(sources, func(s string, _ int) bool {
		return regexp.MustCompile(`^(x|X|Cross|cross)([a-zA-Z][a-z]*)$`).MatchString(s)
	})

	if len(filtered) == 0 {
		return filter.PlatformCrossbell.String()
	}

	return filtered[0]
}

// getOwnerOf gets owner of profile.
func (w *worker) getOwnerOf(_ context.Context, blockNumber *big.Int, profileID *big.Int) (common.Address, error) {
	owner, err := w.profileContract.OwnerOf(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return common.Address{}, fmt.Errorf("get owner of: %w", err)
	}

	return owner, nil
}

// getProfile gets profile handle and uri.
func (w *worker) getProfile(_ context.Context, blockNumber *big.Int, profileID *big.Int) (string, string, error) {
	profile, err := w.profileContract.GetProfile(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return "", "", fmt.Errorf("get profile: %w", err)
	}

	return profile.Handle, profile.Uri, nil
}

// getCharacter gets asset image uri.
func (w *worker) getCharacter(_ context.Context, blockNumber *big.Int, characterID *big.Int) (string, string, error) {
	character, err := w.characterContract.GetCharacter(&bind.CallOpts{BlockNumber: blockNumber}, characterID)
	if err != nil {
		return "", "", fmt.Errorf("get character: %w", err)
	}

	return character.Handle, character.Uri, nil
}

// buildProfileHandleSuffix builds profile handle suffix.
func (w *worker) buildProfileHandleSuffix(_ context.Context, handle string) string {
	return fmt.Sprintf("%s.csb", handle)
}

// getIPFSContent gets IPFS content.
func (w *worker) getIPFSContent(ctx context.Context, contentURI string) (json.RawMessage, error) {
	_, path, err := ipfs.ParseURL(contentURI)
	if err != nil {
		return nil, fmt.Errorf("parse ipfs url: %w", err)
	}

	body, err := w.ipfsClient.Fetch(ctx, path, ipfs.FetchModeQuick)
	if err != nil {
		return nil, fmt.Errorf("quick fetch ipfs: %w", err)
	}

	return io.ReadAll(body)
}

// NewWorker creates a new Lens worker.
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

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize crossbell callers.
	characterContract, err := character.NewCharacterCaller(crossbell.AddressWeb3Entry, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get character contract: %w", err)
	}

	profileContract, err := profile.NewProfileCaller(crossbell.AddressWeb3Entry, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get profile contract: %w", err)
	}

	peripheryContract, err := periphery.NewPeripheryCaller(crossbell.AddressPeriphery, instance.ethereumClient)
	if err != nil {
		return nil, fmt.Errorf("failed to get periphery contract: %w", err)
	}

	instance.characterContract = characterContract
	instance.profileContract = profileContract
	instance.peripheryContract = peripheryContract

	// Initialize crossbell filterers.
	instance.eventFilterer = lo.Must(event.NewEventFilterer(ethereum.AddressGenesis, nil))
	instance.profileFilterer = lo.Must(profile.NewProfileFilterer(ethereum.AddressGenesis, nil))
	instance.tipsFilterer = lo.Must(tips.NewTipsFilterer(ethereum.AddressGenesis, nil))
	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
