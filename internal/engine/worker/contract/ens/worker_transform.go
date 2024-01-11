package ens

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	source "github.com/naturalselectionlabs/rss3-node/internal/engine/source/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract/ens"
	"github.com/naturalselectionlabs/rss3-node/schema"
	"github.com/naturalselectionlabs/rss3-node/schema/filter"
	"github.com/naturalselectionlabs/rss3-node/schema/metadata"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

func (w *worker) transformEnsNameRegisteredV1(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.ethRegistrarControllerV1Filterer.ParseNameRegistered(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRegistered event: %w", err)
	}

	action, err := w.buildEthereumENSRegisterAction(ctx, task, event.Label, ethereum.AddressGenesis, event.Owner, event.Cost)

	if err != nil {
		return nil, fmt.Errorf("build collectible trade action: %w", err)
	}

	return []*schema.Action{action}, nil
}

func (w *worker) transformEnsNameRegisteredV2(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.ethRegistrarControllerV2Filterer.ParseNameRegistered(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRegistered event: %w", err)
	}

	action, err := w.buildEthereumENSRegisterAction(ctx, task, event.Label, ethereum.AddressGenesis, event.Owner, event.BaseCost)

	if err != nil {
		return nil, fmt.Errorf("build collectible trade action: %w", err)
	}

	return []*schema.Action{action}, nil
}

func (w *worker) transformEnsNameRenewed(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.ethRegistrarControllerV2Filterer.ParseNameRenewed(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameRenewed event: %w", err)
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, event.Expires,
			fmt.Sprintf("%s.eth", event.Name), "", "",
			metadata.ActionSocialProfileRenew),
	}, nil
}

func (w *worker) transformEnsTextChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV1Filterer.ParseTextChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseTextChanged event: %w", err)
	}

	nameHash := common.BytesToHash(event.Node[:])

	name, err := w.getEnsName(ctx, log.BlockNumber, nameHash)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, event.Key, "",
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsTextChangedWithValue(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseTextChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseTextChangedWithValue event: %w", err)
	}

	nameHash := common.BytesToHash(event.Node[:])

	name, err := w.getEnsName(ctx, log.BlockNumber, nameHash)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, event.Key, event.Value,
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsNameWrapped(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
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

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, "", "",
			metadata.ActionSocialProfileWrap),
	}, nil
}

func (w *worker) transformEnsNameUnwrapped(ctx context.Context, log ethereum.Log, _ *source.Task) ([]*schema.Action, error) {
	event, err := w.nameWrapperFilterer.ParseNameUnwrapped(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameUnwrapped event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, event.Owner, log.Address, nil,
			name, "", "",
			metadata.ActionSocialProfileUnwrap),
	}, nil
}

func (w *worker) transformEnsFusesSet(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.nameWrapperFilterer.ParseFusesSet(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseFusesSet event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.FusesSet.String(), strconv.Itoa(int(event.Fuses)),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsContenthashChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseContenthashChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseContenthashChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.ContentHashChanged.String(), common.BytesToHash(event.Hash).Hex(),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsNameChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseNameChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseNameChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.NameChanged.String(), event.Name,
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsAddressChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV2Filterer.ParseAddressChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParseAddressChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.AddressChanged.String(), common.BytesToAddress(event.NewAddress).Hex(),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) transformEnsPubkeyChanged(ctx context.Context, log ethereum.Log, task *source.Task) ([]*schema.Action, error) {
	event, err := w.publicResolverV2Filterer.ParsePubkeyChanged(log.Export())
	if err != nil {
		return nil, fmt.Errorf("ParsePubkeyChanged event: %w", err)
	}

	name, err := w.getEnsName(ctx, log.BlockNumber, event.Node)
	if err != nil {
		return nil, err
	}

	return []*schema.Action{
		w.buildEthereumENSProfileAction(ctx, task.Transaction.From, log.Address, nil,
			name, ens.PubkeyChanged.String(), common.Bytes2Hex(secp256k1.CompressPubkey(new(big.Int).SetBytes(event.X[:]), new(big.Int).SetBytes(event.Y[:]))),
			metadata.ActionSocialProfileUpdate),
	}, nil
}

func (w *worker) buildEthereumENSRegisterAction(ctx context.Context, task *source.Task, labels [32]byte, from, to common.Address, cost *big.Int) (*schema.Action, error) {
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

	return &schema.Action{
		Type:     filter.TypeCollectibleTrade,
		Platform: filter.PlatformENS.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: metadata.CollectibleTrade{
			Action: metadata.ActionCollectibleTradeBuy,
			Token:  *tokenMetadata,
			Cost:   costTokenMetadata,
		},
	}, nil
}

func (w *worker) buildEthereumENSProfileAction(_ context.Context, from, to common.Address, expires *big.Int, name, key, value string, action metadata.SocialProfileAction) *schema.Action {
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

	return &schema.Action{
		Type:     filter.TypeSocialProfile,
		Platform: filter.PlatformENS.String(),
		From:     from.String(),
		To:       to.String(),
		Metadata: socialProfile,
	}
}

// getEnsName get ens name by name hash (node)
func (w *worker) getEnsName(_ context.Context, blockNumber *big.Int, nameHash common.Hash) (string, error) {
	var (
		domains string
		err     error
	)

	// TODO: Optimize with multicall3

	// Try public resolver V1
	domains, err = w.publicResolverV1Caller.Name(&bind.CallOpts{BlockNumber: blockNumber}, nameHash)
	if err == nil && len(domains) > 0 {
		return domains, nil
	}

	// Try public resolver V2
	domains, err = w.publicResolverV2Caller.Name(&bind.CallOpts{BlockNumber: blockNumber}, nameHash)
	if err == nil && len(domains) > 0 {
		return domains, nil
	}

	// If still no luck
	return "", fmt.Errorf("fail to find ens name hash: %w, name hash: %v", err, nameHash.String())
}

// decodeName decode ens name
func (w *worker) decodeName(buf []byte) ([]string, error) {
	var (
		offset     = 0
		list       []string
		firstLabel string
	)

	// check first label length
	length := int(buf[offset])
	if length == 0 {
		return []string{firstLabel, "."}, nil
	}

	// decode every label
	for length > 0 {
		// extract label from buf
		label := string(buf[offset+1 : offset+1+length])

		// check label
		if !w.checkValidLabel(label) {
			return nil, fmt.Errorf("invalid label")
		}

		// Put dot between the gap if not the first label
		if offset > 1 {
			list = append(list, ".")
		} else {
			firstLabel = label
		}

		// label in list
		list = append(list, label)

		// update offset and len
		offset += length + 1
		length = int(buf[offset])
	}

	// return first label and complete domain string
	return []string{firstLabel, strings.Join(list, "")}, nil
}

// checkValidLabel: check label
func (w *worker) checkValidLabel(name string) bool {
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == 0 {
			zap.L().Error("Invalid label contained null byte. Skipping.", zap.String("name", name))
			return false
		} else if c == 46 {
			zap.L().Error("Invalid label contained separator char '.'. Skipping.", zap.String("name", name))
			return false
		}
	}

	return true
}
