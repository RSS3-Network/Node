package ens

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rss3-network/node/config"
	"github.com/rss3-network/node/internal/database"
	"github.com/rss3-network/node/internal/database/model"
	"github.com/rss3-network/node/internal/engine"
	source "github.com/rss3-network/node/internal/engine/source/ethereum"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/ens"
	"github.com/rss3-network/node/provider/ethereum/contract/erc1155"
	"github.com/rss3-network/node/provider/ethereum/contract/erc20"
	"github.com/rss3-network/node/provider/ethereum/contract/erc721"
	"github.com/rss3-network/node/provider/ethereum/token"
	"github.com/rss3-network/node/schema/worker/decentralized"
	"github.com/rss3-network/protocol-go/schema"
	activityx "github.com/rss3-network/protocol-go/schema/activity"
	"github.com/rss3-network/protocol-go/schema/metadata"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/rss3-network/protocol-go/schema/tag"
	"github.com/rss3-network/protocol-go/schema/typex"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
)

// Worker is the worker for ENS.
var _ engine.Worker = (*worker)(nil)

type worker struct {
	// Base
	config         *config.Module
	ethereumClient ethereum.Client
	tokenClient    token.Client
	databaseClient database.Client
	// Specific filters
	baseRegistrarImplementationFilterer *ens.BaseRegistrarImplementationFilterer
	ethRegistrarControllerV1Filterer    *ens.ETHRegistrarControllerV1Filterer
	ethRegistrarControllerV2Filterer    *ens.ETHRegistrarControllerV2Filterer
	publicResolverV1Filterer            *ens.PublicResolverV1Filterer
	publicResolverV2Filterer            *ens.PublicResolverV2Filterer
	nameWrapperFilterer                 *ens.NameWrapperFilterer
	// Common filters
	erc20Filterer   *erc20.ERC20Filterer
	erc721Filterer  *erc721.ERC721Filterer
	erc1155Filterer *erc1155.ERC1155Filterer
}

func (w *worker) Name() string {
	return decentralized.ENS.String()
}

func (w *worker) Platform() string {
	return decentralized.PlatformENS.String()
}

func (w *worker) Network() []network.Network {
	return []network.Network{
		network.Ethereum,
	}
}

func (w *worker) Tags() []tag.Tag {
	return []tag.Tag{
		tag.Social,
		tag.Collectible,
	}
}

func (w *worker) Types() []schema.Type {
	return []schema.Type{
		typex.CollectibleTrade,
		typex.SocialProfile,
	}
}

// Filter ens contract address and event hash.
func (w *worker) Filter() engine.DataSourceFilter {
	return &source.Filter{
		LogAddresses: []common.Address{
			ens.AddressBaseRegistrarImplementation,
			ens.AddressETHRegistrarControllerV1,
			ens.AddressETHRegistrarControllerV2,
			ens.AddressPublicResolverV1,
			ens.AddressPublicResolverV2,
			ens.AddressNameWrapper,
		},
		LogTopics: []common.Hash{
			ens.EventNameRegisteredControllerV1,
			ens.EventNameRegisteredControllerV2,
			ens.EventNameRenewed,
			ens.EventTextChanged,
			ens.EventTextChangedWithValue,
			ens.EventNameWrapped,
			ens.EventNameUnwrapped,

			ens.EventFusesSet,
			ens.EventAddressChanged,
			ens.EventContenthashChanged,
			ens.EventNameChanged,
			ens.EventPubkeyChanged,
		},
	}
}

// Transform Ethereum task to activityx.
func (w *worker) Transform(ctx context.Context, task engine.Task) (*activityx.Activity, error) {
	ethereumTask, ok := task.(*source.Task)
	if !ok {
		return nil, fmt.Errorf("invalid task type: %T", task)
	}

	// Build default ens _activities from task.
	_activities, err := ethereumTask.BuildActivity(activityx.WithActivityPlatform(w.Platform()))
	if err != nil {
		return nil, fmt.Errorf("build _activities: %w", err)
	}

	exist := lo.ContainsBy(ethereumTask.Receipt.Logs, func(log *ethereum.Log) bool {
		return w.matchEnsNameRegisteredV2(ctx, *log) || w.matchEnsNameRegisteredV1(ctx, *log)
	})

	// Match and handle ethereum logs.
	for _, log := range ethereumTask.Receipt.Logs {
		var (
			actions []*activityx.Action
			err     error
		)

		if len(log.Topics) == 0 {
			continue
		}

		if exist {
			_activities.Type = typex.CollectibleTrade

			switch {
			case w.matchEnsNameRegisteredV1(ctx, *log):
				actions, err = w.transformEnsNameRegisteredV1(ctx, *log, ethereumTask)
			case w.matchEnsNameRegisteredV2(ctx, *log):
				actions, err = w.transformEnsNameRegisteredV2(ctx, *log, ethereumTask)
			default:
				continue
			}
		} else {
			_activities.Type = typex.SocialProfile

			switch {
			case w.matchEnsNameRenewed(ctx, *log):
				actions, err = w.transformEnsNameRenewed(ctx, *log, ethereumTask)
			case w.matchEnsTextChanged(ctx, *log):
				actions, err = w.transformEnsTextChanged(ctx, *log, ethereumTask)
			case w.matchEnsTextChangedWithValue(ctx, *log):
				actions, err = w.transformEnsTextChangedWithValue(ctx, *log, ethereumTask)
			case w.matchEnsNameWrapped(ctx, *log):
				actions, err = w.transformEnsNameWrapped(ctx, *log, ethereumTask)
			case w.matchEnsNameUnwrapped(ctx, *log):
				actions, err = w.transformEnsNameUnwrapped(ctx, *log, ethereumTask)
			case w.matchEnsFusesSet(ctx, *log):
				actions, err = w.transformEnsFusesSet(ctx, *log, ethereumTask)
			case w.matchEnsContenthashChanged(ctx, *log):
				actions, err = w.transformEnsContenthashChanged(ctx, *log, ethereumTask)
			case w.matchEnsNameChanged(ctx, *log):
				actions, err = w.transformEnsNameChanged(ctx, *log, ethereumTask)
			case w.matchEnsAddressChanged(ctx, *log):
				actions, err = w.transformEnsAddressChanged(ctx, *log, ethereumTask)
			case w.matchEnsPubkeyChanged(ctx, *log):
				actions, err = w.transformEnsPubkeyChanged(ctx, *log, ethereumTask)
			default:
				continue
			}
		}

		if err != nil {
			return nil, err
		}

		// Change _activities type to the first action type.
		for _, action := range actions {
			_activities.Type = action.Type
		}

		_activities.Actions = append(_activities.Actions, actions...)
	}

	return _activities, nil
}

// matchEnsNameRegisteredV1 matches events that ENS name register through V1 contract
func (w *worker) matchEnsNameRegisteredV1(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressETHRegistrarControllerV1 && log.Topics[0] == ens.EventNameRegisteredControllerV1
}

// matchEnsNameRegisteredV2 matches events that ENS name register through V2 contract
func (w *worker) matchEnsNameRegisteredV2(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressETHRegistrarControllerV2 && log.Topics[0] == ens.EventNameRegisteredControllerV2
}

// matchEnsNameRenewed matches events that ENS name renew
func (w *worker) matchEnsNameRenewed(_ context.Context, log ethereum.Log) bool {
	return (log.Address == ens.AddressETHRegistrarControllerV1 || log.Address == ens.AddressETHRegistrarControllerV2) && log.Topics[0] == ens.EventNameRenewed
}

// matchEnsTextChanged matches events that ENS text change
func (w *worker) matchEnsTextChanged(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressPublicResolverV1 && log.Topics[0] == ens.EventTextChanged
}

// matchEnsTextChangedWithValue matches events that ENS text change with value
func (w *worker) matchEnsTextChangedWithValue(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressPublicResolverV2 && log.Topics[0] == ens.EventTextChangedWithValue
}

// matchEnsNameWrapped matches events that ENS name wrapped
func (w *worker) matchEnsNameWrapped(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventNameWrapped
}

// matchEnsNameUnwrapped matches events that ENS name unwrapped
func (w *worker) matchEnsNameUnwrapped(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventNameUnwrapped
}

// matchEnsFusesSet matches events that ENS fuses set
func (w *worker) matchEnsFusesSet(_ context.Context, log ethereum.Log) bool {
	return log.Address == ens.AddressNameWrapper && log.Topics[0] == ens.EventFusesSet
}

// matchEnsContenthashChanged matches events that ENS content hash change
func (w *worker) matchEnsContenthashChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventContenthashChanged
}

// matchEnsNameChanged matches events that ENS name change
func (w *worker) matchEnsNameChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventNameChanged
}

// matchEnsAddressChanged matches events that ENS address change
func (w *worker) matchEnsAddressChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventAddressChanged
}

// matchEnsPubkeyChanged matches events that ENS pubkey change
func (w *worker) matchEnsPubkeyChanged(_ context.Context, log ethereum.Log) bool {
	return log.Topics[0] == ens.EventPubkeyChanged
}

// transformEnsNameRegisteredV1 processes events that ENS name register through V1 contract
func (w *worker) transformEnsNameRegisteredV1(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.ethRegistrarControllerV1Filterer.ParseNameRegistered(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRegistered event: %w", err)
	}

	action, err := w.buildEthereumENSRegisterAction(ctx, task, event.Label, ethereum.AddressGenesis, event.Owner, event.Cost, event.Name)

	if err != nil {
		return nil, fmt.Errorf("build collectible trade action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

// transformEnsNameRegisteredV2 processes events that ENS name register through V2 contract
func (w *worker) transformEnsNameRegisteredV2(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.ethRegistrarControllerV2Filterer.ParseNameRegistered(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRegistered event: %w", err)
	}

	action, err := w.buildEthereumENSRegisterAction(ctx, task, event.Label, ethereum.AddressGenesis, event.Owner, event.BaseCost, event.Name)

	if err != nil {
		return nil, fmt.Errorf("build collectible trade action: %w", err)
	}

	return []*activityx.Action{action}, nil
}

// transformEnsNameRenewed processes events that ENS name renew
func (w *worker) transformEnsNameRenewed(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.ethRegistrarControllerV2Filterer.ParseNameRenewed(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRenewed event: %w", err)
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, event.Expires,
			fmt.Sprintf("%s.eth", event.Name), "", "",
			metadata.ActionSocialProfileRenew),
	}, nil
}

// transformEnsTextChanged processes events that ENS text change
func (w *worker) transformEnsTextChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV1Filterer.ParseTextChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseTextChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, event.Key, "",
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsTextChangedWithValue processes events that ENS text change with value
func (w *worker) transformEnsTextChangedWithValue(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseTextChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseTextChangedWithValue event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, event.Key, event.Value,
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsNameWrapped processes events that ENS name wrapped
func (w *worker) transformEnsNameWrapped(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.nameWrapperFilterer.ParseNameWrapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameWrapped event: %w", err)
	}

	var name string

	decodeEnsName, err := w.decodeName(event.Name)
	if err != nil {
		return nil, err
	}

	if len(decodeEnsName) != 2 {
		return nil, nil
	}

	name = decodeEnsName[1]

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, "", "",
			metadata.ActionSocialProfileWrap),
	}, nil
}

// transformEnsNameUnwrapped processes events that ENS name unwrapped
func (w *worker) transformEnsNameUnwrapped(ctx context.Context, log ethereum.Log, _ *source.Task) ([]*activityx.Action, error) {
	event, err := w.nameWrapperFilterer.ParseNameUnwrapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameUnwrapped event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, event.Owner, log.Address, nil,
			name, "", "",
			metadata.ActionSocialProfileUnwrap),
	}, nil
}

// transformEnsFusesSet processes events that ENS fuses set
func (w *worker) transformEnsFusesSet(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.nameWrapperFilterer.ParseFusesSet(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseFusesSet event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.FusesSet.String(), strconv.Itoa(int(event.Fuses)),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsContenthashChanged processes events that ENS content hash change
func (w *worker) transformEnsContenthashChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseContenthashChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseContenthashChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.ContentHashChanged.String(), common.BytesToHash(event.Hash).Hex(),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsNameChanged processes events that ENS name change
func (w *worker) transformEnsNameChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseNameChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.NameChanged.String(), event.Name,
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsAddressChanged processes events that ENS address change
func (w *worker) transformEnsAddressChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseAddressChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseAddressChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.AddressChanged.String(), common.BytesToAddress(event.NewAddress).Hex(),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

// transformEnsPubkeyChanged processes events that ENS pubkey change
func (w *worker) transformEnsPubkeyChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*activityx.Action, error) {
	event, err := w.publicResolverV2Filterer.ParsePubkeyChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParsePubkeyChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*activityx.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.PubkeyChanged.String(), common.Bytes2Hex(CompressPubkey(new(big.Int).SetBytes(event.X[:]), new(big.Int).SetBytes(event.Y[:]))),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) buildEthereumENSRegisterAction(ctx context.Context, task *source.Task, labels [32]byte, from, to common.Address, cost *big.Int, name string) (*activityx.Action, error) {
	tokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, lo.ToPtr(ens.AddressBaseRegistrarImplementation), new(big.Int).SetBytes(labels[:]), task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s %s: %w", ens.AddressBaseRegistrarImplementation.String(), new(big.Int).SetBytes(labels[:]), err)
	}

	tokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(big.NewInt(1), 0))

	costTokenMetadata, err := w.tokenClient.Lookup(ctx, task.ChainID, nil, nil, task.Header.Number)
	if err != nil {
		return nil, fmt.Errorf("lookup token metadata %s: %w", "", err)
	}

	costTokenMetadata.Value = lo.ToPtr(decimal.NewFromBigInt(cost, 0))

	// Save namehash into database for further query requirements
	fullName := fmt.Sprintf("%s.%s", name, "eth")
	if err = w.databaseClient.SaveDatasetENSNamehash(ctx, &model.ENSNamehash{
		Name: fullName,
		Hash: ens.NameHash(fullName),
	}); err != nil {
		return nil, fmt.Errorf("save dataset ens namehash: %w", err)
	}

	return &activityx.Action{
		Type:     typex.CollectibleTrade,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTrade{
			Action: metadata.ActionCollectibleTradeBuy,
			Token:  *tokenMetadata,
			Cost:   costTokenMetadata,
		},
	}, nil
}

func (w *worker) buildEthereumENSProfileAction(_ context.Context, from, to common.Address, expires *big.Int, name, key, value string, action metadata.SocialProfileAction) *activityx.Action {
	label := strings.Split(name, ".eth")[0]
	tokenID := common.BytesToHash(crypto.Keccak256([]byte(label))).Big()

	socialProfile := metadata.SocialProfile{
		Action:    action,
		ProfileID: tokenID.String(),
		Address:   ens.AddressBaseRegistrarImplementation,
		Handle:    name,
		ImageURI:  fmt.Sprintf("https://metadata.ens.domains/mainnet/%s/%s/image", ens.AddressBaseRegistrarImplementation.String(), tokenID.String()),
		Key:       key,
		Value:     value,
	}

	if expires != nil {
		socialProfile.Expiry = lo.ToPtr(time.Unix(expires.Int64(), 0))
	}

	return &activityx.Action{
		Type:     typex.SocialProfile,
		Platform: w.Platform(),
		From:     from.String(),
		To:       to.String(),
		Metadata: socialProfile,
	}
}

// NewWorker creates a new ENS worker.
func NewWorker(config *config.Module, databaseClient database.Client) (engine.Worker, error) {
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

	// Initialize token client.
	instance.tokenClient = token.NewClient(instance.ethereumClient)

	// Initialize database client.
	instance.databaseClient = databaseClient

	// Initialize ens filterers and callers.
	instance.baseRegistrarImplementationFilterer = lo.Must(ens.NewBaseRegistrarImplementationFilterer(ethereum.AddressGenesis, nil))
	instance.ethRegistrarControllerV1Filterer = lo.Must(ens.NewETHRegistrarControllerV1Filterer(ethereum.AddressGenesis, nil))
	instance.ethRegistrarControllerV2Filterer = lo.Must(ens.NewETHRegistrarControllerV2Filterer(ethereum.AddressGenesis, nil))
	instance.publicResolverV1Filterer = lo.Must(ens.NewPublicResolverV1Filterer(ethereum.AddressGenesis, nil))
	instance.publicResolverV2Filterer = lo.Must(ens.NewPublicResolverV2Filterer(ethereum.AddressGenesis, nil))
	instance.nameWrapperFilterer = lo.Must(ens.NewNameWrapperFilterer(ethereum.AddressGenesis, nil))

	instance.erc20Filterer = lo.Must(erc20.NewERC20Filterer(ethereum.AddressGenesis, nil))
	instance.erc721Filterer = lo.Must(erc721.NewERC721Filterer(ethereum.AddressGenesis, nil))
	instance.erc1155Filterer = lo.Must(erc1155.NewERC1155Filterer(ethereum.AddressGenesis, nil))

	return &instance, nil
}
