// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package opensea

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AdditionalRecipient is an auto generated low-level Go binding around an user-defined struct.
type AdditionalRecipient struct {
	Amount    *big.Int
	Recipient common.Address
}

// AdvancedOrder is an auto generated low-level Go binding around an user-defined struct.
type AdvancedOrder struct {
	Parameters  OrderParameters
	Numerator   *big.Int
	Denominator *big.Int
	Signature   []byte
	ExtraData   []byte
}

// BasicOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type BasicOrderParameters struct {
	ConsiderationToken                common.Address
	ConsiderationIdentifier           *big.Int
	ConsiderationAmount               *big.Int
	Offerer                           common.Address
	Zone                              common.Address
	OfferToken                        common.Address
	OfferIdentifier                   *big.Int
	OfferAmount                       *big.Int
	BasicOrderType                    uint8
	StartTime                         *big.Int
	EndTime                           *big.Int
	ZoneHash                          [32]byte
	Salt                              *big.Int
	OffererConduitKey                 [32]byte
	FulfillerConduitKey               [32]byte
	TotalOriginalAdditionalRecipients *big.Int
	AdditionalRecipients              []AdditionalRecipient
	Signature                         []byte
}

// ConsiderationItem is an auto generated low-level Go binding around an user-defined struct.
type ConsiderationItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
	Recipient            common.Address
}

// CriteriaResolver is an auto generated low-level Go binding around an user-defined struct.
type CriteriaResolver struct {
	OrderIndex    *big.Int
	Side          uint8
	Index         *big.Int
	Identifier    *big.Int
	CriteriaProof [][32]byte
}

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Item       ReceivedItem
	Offerer    common.Address
	ConduitKey [32]byte
}

// Fulfillment is an auto generated low-level Go binding around an user-defined struct.
type Fulfillment struct {
	OfferComponents         []FulfillmentComponent
	ConsiderationComponents []FulfillmentComponent
}

// FulfillmentComponent is an auto generated low-level Go binding around an user-defined struct.
type FulfillmentComponent struct {
	OrderIndex *big.Int
	ItemIndex  *big.Int
}

// OfferItem is an auto generated low-level Go binding around an user-defined struct.
type OfferItem struct {
	ItemType             uint8
	Token                common.Address
	IdentifierOrCriteria *big.Int
	StartAmount          *big.Int
	EndAmount            *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Parameters OrderParameters
	Signature  []byte
}

// OrderComponents is an auto generated low-level Go binding around an user-defined struct.
type OrderComponents struct {
	Offerer       common.Address
	Zone          common.Address
	Offer         []OfferItem
	Consideration []ConsiderationItem
	OrderType     uint8
	StartTime     *big.Int
	EndTime       *big.Int
	ZoneHash      [32]byte
	Salt          *big.Int
	ConduitKey    [32]byte
	Counter       *big.Int
}

// OrderParameters is an auto generated low-level Go binding around an user-defined struct.
type OrderParameters struct {
	Offerer                         common.Address
	Zone                            common.Address
	Offer                           []OfferItem
	Consideration                   []ConsiderationItem
	OrderType                       uint8
	StartTime                       *big.Int
	EndTime                         *big.Int
	ZoneHash                        [32]byte
	Salt                            *big.Int
	ConduitKey                      [32]byte
	TotalOriginalConsiderationItems *big.Int
}

// ReceivedItem is an auto generated low-level Go binding around an user-defined struct.
type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

// SpentItem is an auto generated low-level Go binding around an user-defined struct.
type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

// SeaportV1MetaData contains all meta data concerning the SeaportV1 contract.
var SeaportV1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadContractSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BadReturnValueFromERC20OnTransfer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"BadSignatureV\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancelOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConsiderationLengthNotEqualToTotalOriginal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfallAmount\",\"type\":\"uint256\"}],\"name\":\"ConsiderationNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CriteriaNotEnabledForItem\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"identifiers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"ERC1155BatchTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InexactFraction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientNativeTokensSupplied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Invalid1155BatchTransferEncoding\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBasicOrderParameterEncoding\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidCallToConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"InvalidConduit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidContractOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"InvalidERC721TransferAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFulfillmentComponentData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNativeOfferItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"InvalidRestrictedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fulfillmentIndex\",\"type\":\"uint256\"}],\"name\":\"MismatchedFulfillmentOfferAndConsiderationComponents\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"MissingFulfillmentComponentOnAggregation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingItemAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MissingOriginalConsiderationItems\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeTokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NoContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReentrantCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSpecifiedOrdersAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferAndConsiderationRequiredOnFulfillment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OfferCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderAlreadyFilled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"}],\"name\":\"OrderCriteriaResolverOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderIsCancelled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderPartiallyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PartialFillsNotEnabledForOrder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenTransferGenericFailure\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedConsiderationCriteria\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerIndex\",\"type\":\"uint256\"}],\"name\":\"UnresolvedOfferCriteria\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnusedItemParameters\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"CounterIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structSpentItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structReceivedItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structOrderParameters\",\"name\":\"orderParameters\",\"type\":\"tuple\"}],\"name\":\"OrderValidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"fulfillAdvancedOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableAdvancedOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[][]\",\"name\":\"\",\"type\":\"tuple[][]\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maximumFulfilled\",\"type\":\"uint256\"}],\"name\":\"fulfillAvailableOrders\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"considerationToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"considerationIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"considerationAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerIdentifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"enumBasicOrderType\",\"name\":\"basicOrderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offererConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalAdditionalRecipients\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structAdditionalRecipient[]\",\"name\":\"additionalRecipients\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillBasicOrder_efficient_6GL6yc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"fulfillerConduitKey\",\"type\":\"bytes32\"}],\"name\":\"fulfillOrder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractOfferer\",\"type\":\"address\"}],\"name\":\"getContractOffererNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"}],\"name\":\"getCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"}],\"internalType\":\"structOrderComponents\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSize\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newCounter\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"information\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"conduitController\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"uint120\",\"name\":\"numerator\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"denominator\",\"type\":\"uint120\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structAdvancedOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"criteriaProof\",\"type\":\"bytes32[]\"}],\"internalType\":\"structCriteriaResolver[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"matchAdvancedOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"offerComponents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"orderIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIndex\",\"type\":\"uint256\"}],\"internalType\":\"structFulfillmentComponent[]\",\"name\":\"considerationComponents\",\"type\":\"tuple[]\"}],\"internalType\":\"structFulfillment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"matchOrders\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structReceivedItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"}],\"internalType\":\"structExecution[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zone\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOfferItem[]\",\"name\":\"offer\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumItemType\",\"name\":\"itemType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"identifierOrCriteria\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structConsiderationItem[]\",\"name\":\"consideration\",\"type\":\"tuple[]\"},{\"internalType\":\"enumOrderType\",\"name\":\"orderType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"zoneHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conduitKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"totalOriginalConsiderationItems\",\"type\":\"uint256\"}],\"internalType\":\"structOrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"name\":\"validate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SeaportV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use SeaportV1MetaData.ABI instead.
var SeaportV1ABI = SeaportV1MetaData.ABI

// SeaportV1 is an auto generated Go binding around an Ethereum contract.
type SeaportV1 struct {
	SeaportV1Caller     // Read-only binding to the contract
	SeaportV1Transactor // Write-only binding to the contract
	SeaportV1Filterer   // Log filterer for contract events
}

// SeaportV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type SeaportV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SeaportV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeaportV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeaportV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeaportV1Session struct {
	Contract     *SeaportV1        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeaportV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeaportV1CallerSession struct {
	Contract *SeaportV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SeaportV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeaportV1TransactorSession struct {
	Contract     *SeaportV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SeaportV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type SeaportV1Raw struct {
	Contract *SeaportV1 // Generic contract binding to access the raw methods on
}

// SeaportV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeaportV1CallerRaw struct {
	Contract *SeaportV1Caller // Generic read-only contract binding to access the raw methods on
}

// SeaportV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeaportV1TransactorRaw struct {
	Contract *SeaportV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSeaportV1 creates a new instance of SeaportV1, bound to a specific deployed contract.
func NewSeaportV1(address common.Address, backend bind.ContractBackend) (*SeaportV1, error) {
	contract, err := bindSeaportV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SeaportV1{SeaportV1Caller: SeaportV1Caller{contract: contract}, SeaportV1Transactor: SeaportV1Transactor{contract: contract}, SeaportV1Filterer: SeaportV1Filterer{contract: contract}}, nil
}

// NewSeaportV1Caller creates a new read-only instance of SeaportV1, bound to a specific deployed contract.
func NewSeaportV1Caller(address common.Address, caller bind.ContractCaller) (*SeaportV1Caller, error) {
	contract, err := bindSeaportV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportV1Caller{contract: contract}, nil
}

// NewSeaportV1Transactor creates a new write-only instance of SeaportV1, bound to a specific deployed contract.
func NewSeaportV1Transactor(address common.Address, transactor bind.ContractTransactor) (*SeaportV1Transactor, error) {
	contract, err := bindSeaportV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeaportV1Transactor{contract: contract}, nil
}

// NewSeaportV1Filterer creates a new log filterer instance of SeaportV1, bound to a specific deployed contract.
func NewSeaportV1Filterer(address common.Address, filterer bind.ContractFilterer) (*SeaportV1Filterer, error) {
	contract, err := bindSeaportV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeaportV1Filterer{contract: contract}, nil
}

// bindSeaportV1 binds a generic wrapper to an already deployed contract.
func bindSeaportV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SeaportV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeaportV1 *SeaportV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SeaportV1.Contract.SeaportV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeaportV1 *SeaportV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeaportV1.Contract.SeaportV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeaportV1 *SeaportV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeaportV1.Contract.SeaportV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeaportV1 *SeaportV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SeaportV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeaportV1 *SeaportV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeaportV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeaportV1 *SeaportV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeaportV1.Contract.contract.Transact(opts, method, params...)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_SeaportV1 *SeaportV1Caller) GetContractOffererNonce(opts *bind.CallOpts, contractOfferer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "getContractOffererNonce", contractOfferer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_SeaportV1 *SeaportV1Session) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _SeaportV1.Contract.GetContractOffererNonce(&_SeaportV1.CallOpts, contractOfferer)
}

// GetContractOffererNonce is a free data retrieval call binding the contract method 0xa900866b.
//
// Solidity: function getContractOffererNonce(address contractOfferer) view returns(uint256 nonce)
func (_SeaportV1 *SeaportV1CallerSession) GetContractOffererNonce(contractOfferer common.Address) (*big.Int, error) {
	return _SeaportV1.Contract.GetContractOffererNonce(&_SeaportV1.CallOpts, contractOfferer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_SeaportV1 *SeaportV1Caller) GetCounter(opts *bind.CallOpts, offerer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "getCounter", offerer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_SeaportV1 *SeaportV1Session) GetCounter(offerer common.Address) (*big.Int, error) {
	return _SeaportV1.Contract.GetCounter(&_SeaportV1.CallOpts, offerer)
}

// GetCounter is a free data retrieval call binding the contract method 0xf07ec373.
//
// Solidity: function getCounter(address offerer) view returns(uint256 counter)
func (_SeaportV1 *SeaportV1CallerSession) GetCounter(offerer common.Address) (*big.Int, error) {
	return _SeaportV1.Contract.GetCounter(&_SeaportV1.CallOpts, offerer)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_SeaportV1 *SeaportV1Caller) GetOrderHash(opts *bind.CallOpts, arg0 OrderComponents) ([32]byte, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "getOrderHash", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_SeaportV1 *SeaportV1Session) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _SeaportV1.Contract.GetOrderHash(&_SeaportV1.CallOpts, arg0)
}

// GetOrderHash is a free data retrieval call binding the contract method 0x79df72bd.
//
// Solidity: function getOrderHash((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) ) view returns(bytes32 orderHash)
func (_SeaportV1 *SeaportV1CallerSession) GetOrderHash(arg0 OrderComponents) ([32]byte, error) {
	return _SeaportV1.Contract.GetOrderHash(&_SeaportV1.CallOpts, arg0)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_SeaportV1 *SeaportV1Caller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "getOrderStatus", orderHash)

	outstruct := new(struct {
		IsValidated bool
		IsCancelled bool
		TotalFilled *big.Int
		TotalSize   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsValidated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsCancelled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.TotalFilled = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalSize = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_SeaportV1 *SeaportV1Session) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _SeaportV1.Contract.GetOrderStatus(&_SeaportV1.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns(bool isValidated, bool isCancelled, uint256 totalFilled, uint256 totalSize)
func (_SeaportV1 *SeaportV1CallerSession) GetOrderStatus(orderHash [32]byte) (struct {
	IsValidated bool
	IsCancelled bool
	TotalFilled *big.Int
	TotalSize   *big.Int
}, error) {
	return _SeaportV1.Contract.GetOrderStatus(&_SeaportV1.CallOpts, orderHash)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_SeaportV1 *SeaportV1Caller) Information(opts *bind.CallOpts) (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "information")

	outstruct := new(struct {
		Version           string
		DomainSeparator   [32]byte
		ConduitController common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Version = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DomainSeparator = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.ConduitController = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_SeaportV1 *SeaportV1Session) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _SeaportV1.Contract.Information(&_SeaportV1.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xf47b7740.
//
// Solidity: function information() view returns(string version, bytes32 domainSeparator, address conduitController)
func (_SeaportV1 *SeaportV1CallerSession) Information() (struct {
	Version           string
	DomainSeparator   [32]byte
	ConduitController common.Address
}, error) {
	return _SeaportV1.Contract.Information(&_SeaportV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SeaportV1 *SeaportV1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SeaportV1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SeaportV1 *SeaportV1Session) Name() (string, error) {
	return _SeaportV1.Contract.Name(&_SeaportV1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SeaportV1 *SeaportV1CallerSession) Name() (string, error) {
	return _SeaportV1.Contract.Name(&_SeaportV1.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_SeaportV1 *SeaportV1Transactor) Cancel(opts *bind.TransactOpts, orders []OrderComponents) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "cancel", orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_SeaportV1 *SeaportV1Session) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _SeaportV1.Contract.Cancel(&_SeaportV1.TransactOpts, orders)
}

// Cancel is a paid mutator transaction binding the contract method 0xfd9f1e10.
//
// Solidity: function cancel((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256)[] orders) returns(bool cancelled)
func (_SeaportV1 *SeaportV1TransactorSession) Cancel(orders []OrderComponents) (*types.Transaction, error) {
	return _SeaportV1.Contract.Cancel(&_SeaportV1.TransactOpts, orders)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Transactor) FulfillAdvancedOrder(opts *bind.TransactOpts, arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillAdvancedOrder", arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Session) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAdvancedOrder(&_SeaportV1.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAdvancedOrder is a paid mutator transaction binding the contract method 0xe7acab24.
//
// Solidity: function fulfillAdvancedOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes) , (uint256,uint8,uint256,uint256,bytes32[])[] , bytes32 fulfillerConduitKey, address recipient) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1TransactorSession) FulfillAdvancedOrder(arg0 AdvancedOrder, arg1 []CriteriaResolver, fulfillerConduitKey [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAdvancedOrder(&_SeaportV1.TransactOpts, arg0, arg1, fulfillerConduitKey, recipient)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Transactor) FulfillAvailableAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillAvailableAdvancedOrders", arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Session) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAvailableAdvancedOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableAdvancedOrders is a paid mutator transaction binding the contract method 0x87201b41.
//
// Solidity: function fulfillAvailableAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, address recipient, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1TransactorSession) FulfillAvailableAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 [][]FulfillmentComponent, arg3 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, recipient common.Address, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAvailableAdvancedOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, arg3, fulfillerConduitKey, recipient, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Transactor) FulfillAvailableOrders(opts *bind.TransactOpts, arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillAvailableOrders", arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Session) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAvailableOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillAvailableOrders is a paid mutator transaction binding the contract method 0xed98a574.
//
// Solidity: function fulfillAvailableOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , (uint256,uint256)[][] , (uint256,uint256)[][] , bytes32 fulfillerConduitKey, uint256 maximumFulfilled) payable returns(bool[], ((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1TransactorSession) FulfillAvailableOrders(arg0 []Order, arg1 [][]FulfillmentComponent, arg2 [][]FulfillmentComponent, fulfillerConduitKey [32]byte, maximumFulfilled *big.Int) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillAvailableOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, fulfillerConduitKey, maximumFulfilled)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Transactor) FulfillBasicOrder(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillBasicOrder", parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Session) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillBasicOrder(&_SeaportV1.TransactOpts, parameters)
}

// FulfillBasicOrder is a paid mutator transaction binding the contract method 0xfb0f3ee1.
//
// Solidity: function fulfillBasicOrder((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1TransactorSession) FulfillBasicOrder(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillBasicOrder(&_SeaportV1.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Transactor) FulfillBasicOrderEfficient6GL6yc(opts *bind.TransactOpts, parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillBasicOrder_efficient_6GL6yc", parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Session) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillBasicOrderEfficient6GL6yc(&_SeaportV1.TransactOpts, parameters)
}

// FulfillBasicOrderEfficient6GL6yc is a paid mutator transaction binding the contract method 0x00000000.
//
// Solidity: function fulfillBasicOrder_efficient_6GL6yc((address,uint256,uint256,address,address,address,uint256,uint256,uint8,uint256,uint256,bytes32,uint256,bytes32,bytes32,uint256,(uint256,address)[],bytes) parameters) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1TransactorSession) FulfillBasicOrderEfficient6GL6yc(parameters BasicOrderParameters) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillBasicOrderEfficient6GL6yc(&_SeaportV1.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Transactor) FulfillOrder(opts *bind.TransactOpts, arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "fulfillOrder", arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1Session) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillOrder(&_SeaportV1.TransactOpts, arg0, fulfillerConduitKey)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xb3a34c4c.
//
// Solidity: function fulfillOrder(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes) , bytes32 fulfillerConduitKey) payable returns(bool fulfilled)
func (_SeaportV1 *SeaportV1TransactorSession) FulfillOrder(arg0 Order, fulfillerConduitKey [32]byte) (*types.Transaction, error) {
	return _SeaportV1.Contract.FulfillOrder(&_SeaportV1.TransactOpts, arg0, fulfillerConduitKey)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_SeaportV1 *SeaportV1Transactor) IncrementCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "incrementCounter")
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_SeaportV1 *SeaportV1Session) IncrementCounter() (*types.Transaction, error) {
	return _SeaportV1.Contract.IncrementCounter(&_SeaportV1.TransactOpts)
}

// IncrementCounter is a paid mutator transaction binding the contract method 0x5b34b966.
//
// Solidity: function incrementCounter() returns(uint256 newCounter)
func (_SeaportV1 *SeaportV1TransactorSession) IncrementCounter() (*types.Transaction, error) {
	return _SeaportV1.Contract.IncrementCounter(&_SeaportV1.TransactOpts)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Transactor) MatchAdvancedOrders(opts *bind.TransactOpts, arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "matchAdvancedOrders", arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Session) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.Contract.MatchAdvancedOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchAdvancedOrders is a paid mutator transaction binding the contract method 0xf2d12b12.
//
// Solidity: function matchAdvancedOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),uint120,uint120,bytes,bytes)[] , (uint256,uint8,uint256,uint256,bytes32[])[] , ((uint256,uint256)[],(uint256,uint256)[])[] , address recipient) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1TransactorSession) MatchAdvancedOrders(arg0 []AdvancedOrder, arg1 []CriteriaResolver, arg2 []Fulfillment, recipient common.Address) (*types.Transaction, error) {
	return _SeaportV1.Contract.MatchAdvancedOrders(&_SeaportV1.TransactOpts, arg0, arg1, arg2, recipient)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Transactor) MatchOrders(opts *bind.TransactOpts, arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "matchOrders", arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1Session) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _SeaportV1.Contract.MatchOrders(&_SeaportV1.TransactOpts, arg0, arg1)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xa8174404.
//
// Solidity: function matchOrders(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] , ((uint256,uint256)[],(uint256,uint256)[])[] ) payable returns(((uint8,address,uint256,uint256,address),address,bytes32)[])
func (_SeaportV1 *SeaportV1TransactorSession) MatchOrders(arg0 []Order, arg1 []Fulfillment) (*types.Transaction, error) {
	return _SeaportV1.Contract.MatchOrders(&_SeaportV1.TransactOpts, arg0, arg1)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_SeaportV1 *SeaportV1Transactor) Validate(opts *bind.TransactOpts, arg0 []Order) (*types.Transaction, error) {
	return _SeaportV1.contract.Transact(opts, "validate", arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_SeaportV1 *SeaportV1Session) Validate(arg0 []Order) (*types.Transaction, error) {
	return _SeaportV1.Contract.Validate(&_SeaportV1.TransactOpts, arg0)
}

// Validate is a paid mutator transaction binding the contract method 0x88147732.
//
// Solidity: function validate(((address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256),bytes)[] ) returns(bool)
func (_SeaportV1 *SeaportV1TransactorSession) Validate(arg0 []Order) (*types.Transaction, error) {
	return _SeaportV1.Contract.Validate(&_SeaportV1.TransactOpts, arg0)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SeaportV1 *SeaportV1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeaportV1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SeaportV1 *SeaportV1Session) Receive() (*types.Transaction, error) {
	return _SeaportV1.Contract.Receive(&_SeaportV1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SeaportV1 *SeaportV1TransactorSession) Receive() (*types.Transaction, error) {
	return _SeaportV1.Contract.Receive(&_SeaportV1.TransactOpts)
}

// SeaportV1CounterIncrementedIterator is returned from FilterCounterIncremented and is used to iterate over the raw logs and unpacked data for CounterIncremented events raised by the SeaportV1 contract.
type SeaportV1CounterIncrementedIterator struct {
	Event *SeaportV1CounterIncremented // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportV1CounterIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportV1CounterIncremented)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportV1CounterIncremented)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportV1CounterIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportV1CounterIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportV1CounterIncremented represents a CounterIncremented event raised by the SeaportV1 contract.
type SeaportV1CounterIncremented struct {
	NewCounter *big.Int
	Offerer    common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCounterIncremented is a free log retrieval operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_SeaportV1 *SeaportV1Filterer) FilterCounterIncremented(opts *bind.FilterOpts, offerer []common.Address) (*SeaportV1CounterIncrementedIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _SeaportV1.contract.FilterLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return &SeaportV1CounterIncrementedIterator{contract: _SeaportV1.contract, event: "CounterIncremented", logs: logs, sub: sub}, nil
}

// WatchCounterIncremented is a free log subscription operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_SeaportV1 *SeaportV1Filterer) WatchCounterIncremented(opts *bind.WatchOpts, sink chan<- *SeaportV1CounterIncremented, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _SeaportV1.contract.WatchLogs(opts, "CounterIncremented", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportV1CounterIncremented)
				if err := _SeaportV1.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCounterIncremented is a log parse operation binding the contract event 0x721c20121297512b72821b97f5326877ea8ecf4bb9948fea5bfcb6453074d37f.
//
// Solidity: event CounterIncremented(uint256 newCounter, address indexed offerer)
func (_SeaportV1 *SeaportV1Filterer) ParseCounterIncremented(log types.Log) (*SeaportV1CounterIncremented, error) {
	event := new(SeaportV1CounterIncremented)
	if err := _SeaportV1.contract.UnpackLog(event, "CounterIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportV1OrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the SeaportV1 contract.
type SeaportV1OrderCancelledIterator struct {
	Event *SeaportV1OrderCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportV1OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportV1OrderCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportV1OrderCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportV1OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportV1OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportV1OrderCancelled represents a OrderCancelled event raised by the SeaportV1 contract.
type SeaportV1OrderCancelled struct {
	OrderHash [32]byte
	Offerer   common.Address
	Zone      common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_SeaportV1 *SeaportV1Filterer) FilterOrderCancelled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportV1OrderCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _SeaportV1.contract.FilterLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportV1OrderCancelledIterator{contract: _SeaportV1.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_SeaportV1 *SeaportV1Filterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *SeaportV1OrderCancelled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _SeaportV1.contract.WatchLogs(opts, "OrderCancelled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportV1OrderCancelled)
				if err := _SeaportV1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderCancelled is a log parse operation binding the contract event 0x6bacc01dbe442496068f7d234edd811f1a5f833243e0aec824f86ab861f3c90d.
//
// Solidity: event OrderCancelled(bytes32 orderHash, address indexed offerer, address indexed zone)
func (_SeaportV1 *SeaportV1Filterer) ParseOrderCancelled(log types.Log) (*SeaportV1OrderCancelled, error) {
	event := new(SeaportV1OrderCancelled)
	if err := _SeaportV1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportV1OrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the SeaportV1 contract.
type SeaportV1OrderFulfilledIterator struct {
	Event *SeaportV1OrderFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportV1OrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportV1OrderFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportV1OrderFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportV1OrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportV1OrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportV1OrderFulfilled represents a OrderFulfilled event raised by the SeaportV1 contract.
type SeaportV1OrderFulfilled struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_SeaportV1 *SeaportV1Filterer) FilterOrderFulfilled(opts *bind.FilterOpts, offerer []common.Address, zone []common.Address) (*SeaportV1OrderFulfilledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _SeaportV1.contract.FilterLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return &SeaportV1OrderFulfilledIterator{contract: _SeaportV1.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_SeaportV1 *SeaportV1Filterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *SeaportV1OrderFulfilled, offerer []common.Address, zone []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}
	var zoneRule []interface{}
	for _, zoneItem := range zone {
		zoneRule = append(zoneRule, zoneItem)
	}

	logs, sub, err := _SeaportV1.contract.WatchLogs(opts, "OrderFulfilled", offererRule, zoneRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportV1OrderFulfilled)
				if err := _SeaportV1.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderFulfilled is a log parse operation binding the contract event 0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31.
//
// Solidity: event OrderFulfilled(bytes32 orderHash, address indexed offerer, address indexed zone, address recipient, (uint8,address,uint256,uint256)[] offer, (uint8,address,uint256,uint256,address)[] consideration)
func (_SeaportV1 *SeaportV1Filterer) ParseOrderFulfilled(log types.Log) (*SeaportV1OrderFulfilled, error) {
	event := new(SeaportV1OrderFulfilled)
	if err := _SeaportV1.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportV1OrderValidatedIterator is returned from FilterOrderValidated and is used to iterate over the raw logs and unpacked data for OrderValidated events raised by the SeaportV1 contract.
type SeaportV1OrderValidatedIterator struct {
	Event *SeaportV1OrderValidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportV1OrderValidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportV1OrderValidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportV1OrderValidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportV1OrderValidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportV1OrderValidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportV1OrderValidated represents a OrderValidated event raised by the SeaportV1 contract.
type SeaportV1OrderValidated struct {
	OrderHash       [32]byte
	OrderParameters OrderParameters
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterOrderValidated is a free log retrieval operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_SeaportV1 *SeaportV1Filterer) FilterOrderValidated(opts *bind.FilterOpts) (*SeaportV1OrderValidatedIterator, error) {

	logs, sub, err := _SeaportV1.contract.FilterLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return &SeaportV1OrderValidatedIterator{contract: _SeaportV1.contract, event: "OrderValidated", logs: logs, sub: sub}, nil
}

// WatchOrderValidated is a free log subscription operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_SeaportV1 *SeaportV1Filterer) WatchOrderValidated(opts *bind.WatchOpts, sink chan<- *SeaportV1OrderValidated) (event.Subscription, error) {

	logs, sub, err := _SeaportV1.contract.WatchLogs(opts, "OrderValidated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportV1OrderValidated)
				if err := _SeaportV1.contract.UnpackLog(event, "OrderValidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrderValidated is a log parse operation binding the contract event 0xf280791efe782edcf06ce15c8f4dff17601db3b88eb3805a0db7d77faf757f04.
//
// Solidity: event OrderValidated(bytes32 orderHash, (address,address,(uint8,address,uint256,uint256,uint256)[],(uint8,address,uint256,uint256,uint256,address)[],uint8,uint256,uint256,bytes32,uint256,bytes32,uint256) orderParameters)
func (_SeaportV1 *SeaportV1Filterer) ParseOrderValidated(log types.Log) (*SeaportV1OrderValidated, error) {
	event := new(SeaportV1OrderValidated)
	if err := _SeaportV1.contract.UnpackLog(event, "OrderValidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeaportV1OrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the SeaportV1 contract.
type SeaportV1OrdersMatchedIterator struct {
	Event *SeaportV1OrdersMatched // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SeaportV1OrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeaportV1OrdersMatched)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SeaportV1OrdersMatched)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SeaportV1OrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeaportV1OrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeaportV1OrdersMatched represents a OrdersMatched event raised by the SeaportV1 contract.
type SeaportV1OrdersMatched struct {
	OrderHashes [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_SeaportV1 *SeaportV1Filterer) FilterOrdersMatched(opts *bind.FilterOpts) (*SeaportV1OrdersMatchedIterator, error) {

	logs, sub, err := _SeaportV1.contract.FilterLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return &SeaportV1OrdersMatchedIterator{contract: _SeaportV1.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_SeaportV1 *SeaportV1Filterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *SeaportV1OrdersMatched) (event.Subscription, error) {

	logs, sub, err := _SeaportV1.contract.WatchLogs(opts, "OrdersMatched")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeaportV1OrdersMatched)
				if err := _SeaportV1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOrdersMatched is a log parse operation binding the contract event 0x4b9f2d36e1b4c93de62cc077b00b1a91d84b6c31b4a14e012718dcca230689e7.
//
// Solidity: event OrdersMatched(bytes32[] orderHashes)
func (_SeaportV1 *SeaportV1Filterer) ParseOrdersMatched(log types.Log) (*SeaportV1OrdersMatched, error) {
	event := new(SeaportV1OrdersMatched)
	if err := _SeaportV1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
