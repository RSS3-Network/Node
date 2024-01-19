package crossbell

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/config"
	"github.com/naturalselectionlabs/rss3-node/internal/engine"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/crossbell"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/crossbell/character"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/crossbell/event"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/crossbell/periphery"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/crossbell/profile"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/token"
	"github.com/naturalselectionlabs/rss3-node/provider/ipfs"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
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

// transformProfileCreated transforms ProfileCreated event.
func (w *worker) transformProfileCreated(ctx context.Context, _ *source.Task, log *ethereum.Log) ([]*schema.Action, error) {
	event, err := w.profileFilterer.ParseProfileCreated(log.Export())
	if err != nil {
		return nil, fmt.Errorf("parse profile created: %w", err)
	}

	profile, err := w.buildTransactionProfileMetadata(ctx, log.BlockNumber, metadata.ActionSocialProfileCreate, event.ProfileId, event.To, event.Handle, "")
	if err != nil {
		return nil, err
	}

	action := w.buildTransactionProfileAction(ctx, event.Creator, event.To, filter.TypeSocialProfile, *profile)

	return []*schema.Action{
		action,
	}, nil
}

// buildTransactionProfileAction builds profile action.
func (w *worker) buildTransactionProfileAction(_ context.Context, from, to common.Address, actionType filter.Type, profile metadata.SocialProfile) *schema.Action {
	return &schema.Action{
		From:     from.String(),
		To:       to.String(),
		Platform: filter.PlatformCrossbell.String(),
		Type:     actionType,
		Metadata: profile,
	}
}

// buildTransactionProfileMetadata builds profile metadata.
func (w *worker) buildTransactionProfileMetadata(
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

		profile.Address, err = w.getEthereumOwnerOf(ctx, blockNumber, profileID)
		if err != nil {
			return nil, err
		}
	}

	if lo.IsEmpty(handle) || lo.IsEmpty(uri) {
		profileHandle, profileURI, err := w.getEthereumProfile(ctx, blockNumber, profileID)
		if err != nil {
			return nil, err
		}

		profile.Handle = lo.If(lo.IsEmpty(handle), profileHandle).Else(handle)
		uri = lo.If(lo.IsEmpty(uri), profileURI).Else(uri)
	}

	profile.Handle = w.buildEthereumProfileHandleSuffix(ctx, profile.Handle)

	if lo.IsNotEmpty(uri) {
		content, err := w.getEthereumIPFSContent(ctx, uri)
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

	return profile, nil
}

// getEthereumOwnerOf gets owner of profile.
func (w *worker) getEthereumOwnerOf(_ context.Context, blockNumber *big.Int, profileID *big.Int) (common.Address, error) {
	owner, err := w.profileContract.OwnerOf(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return common.Address{}, fmt.Errorf("get owner of: %w", err)
	}

	return owner, nil
}

// getEthereumProfile gets profile handle and uri.
func (w *worker) getEthereumProfile(_ context.Context, blockNumber *big.Int, profileID *big.Int) (string, string, error) {
	profile, err := w.profileContract.GetProfile(&bind.CallOpts{BlockNumber: blockNumber}, profileID)
	if err != nil {
		return "", "", fmt.Errorf("get profile: %w", err)
	}

	return profile.Handle, profile.Uri, nil
}

// buildEthereumProfileHandleSuffix builds profile handle suffix.
func (w *worker) buildEthereumProfileHandleSuffix(_ context.Context, handle string) string {
	return fmt.Sprintf("%s.csb", handle)
}

// getEthereumIPFSContent gets IPFS content.
func (w *worker) getEthereumIPFSContent(ctx context.Context, contentURI string) (json.RawMessage, error) {
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

	return &instance, nil
}
