// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package looksrare

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

// ILooksRareProtocolNonceInvalidationParameters is an auto generated low-level Go binding around an user-defined struct.
type ILooksRareProtocolNonceInvalidationParameters struct {
	OrderHash          [32]byte
	OrderNonce         *big.Int
	IsNonceInvalidated bool
}

// OrderStructsMaker is an auto generated low-level Go binding around an user-defined struct.
type OrderStructsMaker struct {
	QuoteType            uint8
	GlobalNonce          *big.Int
	SubsetNonce          *big.Int
	OrderNonce           *big.Int
	StrategyId           *big.Int
	CollectionType       uint8
	Collection           common.Address
	Currency             common.Address
	Signer               common.Address
	StartTime            *big.Int
	EndTime              *big.Int
	Price                *big.Int
	ItemIds              []*big.Int
	Amounts              []*big.Int
	AdditionalParameters []byte
}

// OrderStructsMerkleTree is an auto generated low-level Go binding around an user-defined struct.
type OrderStructsMerkleTree struct {
	Root  [32]byte
	Proof []OrderStructsMerkleTreeNode
}

// OrderStructsMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type OrderStructsMerkleTreeNode struct {
	Value    [32]byte
	Position uint8
}

// OrderStructsTaker is an auto generated low-level Go binding around an user-defined struct.
type OrderStructsTaker struct {
	Recipient            common.Address
	AdditionalParameters []byte
}

// ExchangeV2MetaData contains all meta data concerning the ExchangeV2 contract.
var ExchangeV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_transferManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIdInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreatorFeeBpTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20TransferFromFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MerkleProofInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MerkleProofTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewGasLimitETHTransferTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NewProtocolFeeRecipientCannotBeNullAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOngoingTransferInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoSelectorForStrategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoncesInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAffiliateController\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotV2Strategy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullSignerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutsideOfTimeRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PercentageTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuoteTypeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RenouncementNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameDomainSeparator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureEOAInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureERC1271Invalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"SignatureLengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureParameterSInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"SignatureParameterVInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategyHasNoSelector\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"}],\"name\":\"StrategyNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategyNotUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StrategyProtocolFeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPotentialOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"affiliateFee\",\"type\":\"uint256\"}],\"name\":\"AffiliatePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CancelOwnershipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"CurrencyStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitiateOwnershipRenouncement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"potentialOwner\",\"type\":\"address\"}],\"name\":\"InitiateOwnershipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliateController\",\"type\":\"address\"}],\"name\":\"NewAffiliateController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"NewAffiliateProgramStatus\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"NewAffiliateRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"askNonce\",\"type\":\"uint256\"}],\"name\":\"NewBidAskNonces\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creatorFeeManager\",\"type\":\"address\"}],\"name\":\"NewCreatorFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"NewDomainSeparator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gasLimitETHTransfer\",\"type\":\"uint256\"}],\"name\":\"NewGasLimitETHTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxCreatorFeeBp\",\"type\":\"uint256\"}],\"name\":\"NewMaxCreatorFeeBp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"standardProtocolFeeBp\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minTotalFeeBp\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"maxProtocolFeeBp\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isMakerBid\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"NewStrategy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"OrderNoncesCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"standardProtocolFeeBp\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minTotalFeeBp\",\"type\":\"uint16\"}],\"name\":\"StrategyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"subsetNonces\",\"type\":\"uint256[]\"}],\"name\":\"SubsetNoncesCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNonceInvalidated\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structILooksRareProtocol.NonceInvalidationParameters\",\"name\":\"nonceInvalidationParameters\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"askUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[2]\",\"name\":\"feeRecipients\",\"type\":\"address[2]\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"feeAmounts\",\"type\":\"uint256[3]\"}],\"name\":\"TakerAsk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNonceInvalidated\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structILooksRareProtocol.NonceInvalidationParameters\",\"name\":\"nonceInvalidationParameters\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"address[2]\",\"name\":\"feeRecipients\",\"type\":\"address[2]\"},{\"indexed\":false,\"internalType\":\"uint256[3]\",\"name\":\"feeAmounts\",\"type\":\"uint256[3]\"}],\"name\":\"TakerBid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAGIC_VALUE_ORDER_NONCE_EXECUTED\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"standardProtocolFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minTotalFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxProtocolFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"isMakerBid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"addStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"affiliateController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"affiliateRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelOrderNonces\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"subsetNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelSubsetNonces\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOwnershipRenouncement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creatorFeeManager\",\"outputs\":[{\"internalType\":\"contractICreatorFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Taker[]\",\"name\":\"takerBids\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumQuoteType\",\"name\":\"quoteType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"globalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"enumCollectionType\",\"name\":\"collectionType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Maker[]\",\"name\":\"makerAsks\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"makerSignatures\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"enumOrderStructs.MerkleTreeNodePosition\",\"name\":\"position\",\"type\":\"uint8\"}],\"internalType\":\"structOrderStructs.MerkleTreeNode[]\",\"name\":\"proof\",\"type\":\"tuple[]\"}],\"internalType\":\"structOrderStructs.MerkleTree[]\",\"name\":\"merkleTrees\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAtomic\",\"type\":\"bool\"}],\"name\":\"executeMultipleTakerBids\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Taker\",\"name\":\"takerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumQuoteType\",\"name\":\"quoteType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"globalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"enumCollectionType\",\"name\":\"collectionType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Maker\",\"name\":\"makerBid\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"makerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"enumOrderStructs.MerkleTreeNodePosition\",\"name\":\"position\",\"type\":\"uint8\"}],\"internalType\":\"structOrderStructs.MerkleTreeNode[]\",\"name\":\"proof\",\"type\":\"tuple[]\"}],\"internalType\":\"structOrderStructs.MerkleTree\",\"name\":\"merkleTree\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"executeTakerAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Taker\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumQuoteType\",\"name\":\"quoteType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"globalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"enumCollectionType\",\"name\":\"collectionType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Maker\",\"name\":\"makerAsk\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"makerSignature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"enumOrderStructs.MerkleTreeNodePosition\",\"name\":\"position\",\"type\":\"uint8\"}],\"internalType\":\"structOrderStructs.MerkleTreeNode[]\",\"name\":\"proof\",\"type\":\"tuple[]\"}],\"internalType\":\"structOrderStructs.MerkleTree\",\"name\":\"merkleTree\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"executeTakerBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"proofLength\",\"type\":\"uint256\"}],\"name\":\"hashBatchOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"batchOrderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"bid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ask\",\"type\":\"bool\"}],\"name\":\"incrementBidAskNonces\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateOwnershipRenouncement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"initiateOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAffiliateProgramActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCurrencyAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCreatorFeeBp\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownershipStatus\",\"outputs\":[{\"internalType\":\"enumIOwnableTwoSteps.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Taker\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumQuoteType\",\"name\":\"quoteType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"globalNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"enumCollectionType\",\"name\":\"collectionType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"itemIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"additionalParameters\",\"type\":\"bytes\"}],\"internalType\":\"structOrderStructs.Maker\",\"name\":\"makerAsk\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"restrictedExecuteTakerBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"strategyInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"standardProtocolFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minTotalFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxProtocolFeeBp\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"internalType\":\"bool\",\"name\":\"isMakerBid\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferManager\",\"outputs\":[{\"internalType\":\"contractTransferManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAffiliateController\",\"type\":\"address\"}],\"name\":\"updateAffiliateController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"name\":\"updateAffiliateProgramStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bp\",\"type\":\"uint256\"}],\"name\":\"updateAffiliateRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newCreatorFeeManager\",\"type\":\"address\"}],\"name\":\"updateCreatorFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAllowed\",\"type\":\"bool\"}],\"name\":\"updateCurrencyStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateDomainSeparator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasLimitETHTransfer\",\"type\":\"uint256\"}],\"name\":\"updateETHGasLimitForTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMaxCreatorFeeBp\",\"type\":\"uint16\"}],\"name\":\"updateMaxCreatorFeeBp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateProtocolFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strategyId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"newStandardProtocolFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newMinTotalFee\",\"type\":\"uint16\"}],\"name\":\"updateStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBidAskNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bidNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"askNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userOrderNonce\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userSubsetNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExchangeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeV2MetaData.ABI instead.
var ExchangeV2ABI = ExchangeV2MetaData.ABI

// ExchangeV2 is an auto generated Go binding around an Ethereum contract.
type ExchangeV2 struct {
	ExchangeV2Caller     // Read-only binding to the contract
	ExchangeV2Transactor // Write-only binding to the contract
	ExchangeV2Filterer   // Log filterer for contract events
}

// ExchangeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeV2Session struct {
	Contract     *ExchangeV2       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeV2CallerSession struct {
	Contract *ExchangeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExchangeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeV2TransactorSession struct {
	Contract     *ExchangeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExchangeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeV2Raw struct {
	Contract *ExchangeV2 // Generic contract binding to access the raw methods on
}

// ExchangeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeV2CallerRaw struct {
	Contract *ExchangeV2Caller // Generic read-only contract binding to access the raw methods on
}

// ExchangeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeV2TransactorRaw struct {
	Contract *ExchangeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeV2 creates a new instance of ExchangeV2, bound to a specific deployed contract.
func NewExchangeV2(address common.Address, backend bind.ContractBackend) (*ExchangeV2, error) {
	contract, err := bindExchangeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeV2{ExchangeV2Caller: ExchangeV2Caller{contract: contract}, ExchangeV2Transactor: ExchangeV2Transactor{contract: contract}, ExchangeV2Filterer: ExchangeV2Filterer{contract: contract}}, nil
}

// NewExchangeV2Caller creates a new read-only instance of ExchangeV2, bound to a specific deployed contract.
func NewExchangeV2Caller(address common.Address, caller bind.ContractCaller) (*ExchangeV2Caller, error) {
	contract, err := bindExchangeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV2Caller{contract: contract}, nil
}

// NewExchangeV2Transactor creates a new write-only instance of ExchangeV2, bound to a specific deployed contract.
func NewExchangeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeV2Transactor, error) {
	contract, err := bindExchangeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeV2Transactor{contract: contract}, nil
}

// NewExchangeV2Filterer creates a new log filterer instance of ExchangeV2, bound to a specific deployed contract.
func NewExchangeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeV2Filterer, error) {
	contract, err := bindExchangeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeV2Filterer{contract: contract}, nil
}

// bindExchangeV2 binds a generic wrapper to an already deployed contract.
func bindExchangeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV2 *ExchangeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV2.Contract.ExchangeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV2 *ExchangeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExchangeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV2 *ExchangeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExchangeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeV2 *ExchangeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeV2 *ExchangeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeV2 *ExchangeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeV2.Contract.contract.Transact(opts, method, params...)
}

// MAGICVALUEORDERNONCEEXECUTED is a free data retrieval call binding the contract method 0x463357ec.
//
// Solidity: function MAGIC_VALUE_ORDER_NONCE_EXECUTED() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Caller) MAGICVALUEORDERNONCEEXECUTED(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "MAGIC_VALUE_ORDER_NONCE_EXECUTED")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAGICVALUEORDERNONCEEXECUTED is a free data retrieval call binding the contract method 0x463357ec.
//
// Solidity: function MAGIC_VALUE_ORDER_NONCE_EXECUTED() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Session) MAGICVALUEORDERNONCEEXECUTED() ([32]byte, error) {
	return _ExchangeV2.Contract.MAGICVALUEORDERNONCEEXECUTED(&_ExchangeV2.CallOpts)
}

// MAGICVALUEORDERNONCEEXECUTED is a free data retrieval call binding the contract method 0x463357ec.
//
// Solidity: function MAGIC_VALUE_ORDER_NONCE_EXECUTED() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2CallerSession) MAGICVALUEORDERNONCEEXECUTED() ([32]byte, error) {
	return _ExchangeV2.Contract.MAGICVALUEORDERNONCEEXECUTED(&_ExchangeV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) WETH() (common.Address, error) {
	return _ExchangeV2.Contract.WETH(&_ExchangeV2.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) WETH() (common.Address, error) {
	return _ExchangeV2.Contract.WETH(&_ExchangeV2.CallOpts)
}

// AffiliateController is a free data retrieval call binding the contract method 0x4cbac9dc.
//
// Solidity: function affiliateController() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) AffiliateController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "affiliateController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AffiliateController is a free data retrieval call binding the contract method 0x4cbac9dc.
//
// Solidity: function affiliateController() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) AffiliateController() (common.Address, error) {
	return _ExchangeV2.Contract.AffiliateController(&_ExchangeV2.CallOpts)
}

// AffiliateController is a free data retrieval call binding the contract method 0x4cbac9dc.
//
// Solidity: function affiliateController() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) AffiliateController() (common.Address, error) {
	return _ExchangeV2.Contract.AffiliateController(&_ExchangeV2.CallOpts)
}

// AffiliateRates is a free data retrieval call binding the contract method 0x7a7d8851.
//
// Solidity: function affiliateRates(address ) view returns(uint256)
func (_ExchangeV2 *ExchangeV2Caller) AffiliateRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "affiliateRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AffiliateRates is a free data retrieval call binding the contract method 0x7a7d8851.
//
// Solidity: function affiliateRates(address ) view returns(uint256)
func (_ExchangeV2 *ExchangeV2Session) AffiliateRates(arg0 common.Address) (*big.Int, error) {
	return _ExchangeV2.Contract.AffiliateRates(&_ExchangeV2.CallOpts, arg0)
}

// AffiliateRates is a free data retrieval call binding the contract method 0x7a7d8851.
//
// Solidity: function affiliateRates(address ) view returns(uint256)
func (_ExchangeV2 *ExchangeV2CallerSession) AffiliateRates(arg0 common.Address) (*big.Int, error) {
	return _ExchangeV2.Contract.AffiliateRates(&_ExchangeV2.CallOpts, arg0)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ExchangeV2 *ExchangeV2Caller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ExchangeV2 *ExchangeV2Session) ChainId() (*big.Int, error) {
	return _ExchangeV2.Contract.ChainId(&_ExchangeV2.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_ExchangeV2 *ExchangeV2CallerSession) ChainId() (*big.Int, error) {
	return _ExchangeV2.Contract.ChainId(&_ExchangeV2.CallOpts)
}

// CreatorFeeManager is a free data retrieval call binding the contract method 0x838b8f5c.
//
// Solidity: function creatorFeeManager() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) CreatorFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "creatorFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreatorFeeManager is a free data retrieval call binding the contract method 0x838b8f5c.
//
// Solidity: function creatorFeeManager() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) CreatorFeeManager() (common.Address, error) {
	return _ExchangeV2.Contract.CreatorFeeManager(&_ExchangeV2.CallOpts)
}

// CreatorFeeManager is a free data retrieval call binding the contract method 0x838b8f5c.
//
// Solidity: function creatorFeeManager() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) CreatorFeeManager() (common.Address, error) {
	return _ExchangeV2.Contract.CreatorFeeManager(&_ExchangeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Caller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Session) DomainSeparator() ([32]byte, error) {
	return _ExchangeV2.Contract.DomainSeparator(&_ExchangeV2.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_ExchangeV2 *ExchangeV2CallerSession) DomainSeparator() ([32]byte, error) {
	return _ExchangeV2.Contract.DomainSeparator(&_ExchangeV2.CallOpts)
}

// HashBatchOrder is a free data retrieval call binding the contract method 0x86c07642.
//
// Solidity: function hashBatchOrder(bytes32 root, uint256 proofLength) pure returns(bytes32 batchOrderHash)
func (_ExchangeV2 *ExchangeV2Caller) HashBatchOrder(opts *bind.CallOpts, root [32]byte, proofLength *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "hashBatchOrder", root, proofLength)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashBatchOrder is a free data retrieval call binding the contract method 0x86c07642.
//
// Solidity: function hashBatchOrder(bytes32 root, uint256 proofLength) pure returns(bytes32 batchOrderHash)
func (_ExchangeV2 *ExchangeV2Session) HashBatchOrder(root [32]byte, proofLength *big.Int) ([32]byte, error) {
	return _ExchangeV2.Contract.HashBatchOrder(&_ExchangeV2.CallOpts, root, proofLength)
}

// HashBatchOrder is a free data retrieval call binding the contract method 0x86c07642.
//
// Solidity: function hashBatchOrder(bytes32 root, uint256 proofLength) pure returns(bytes32 batchOrderHash)
func (_ExchangeV2 *ExchangeV2CallerSession) HashBatchOrder(root [32]byte, proofLength *big.Int) ([32]byte, error) {
	return _ExchangeV2.Contract.HashBatchOrder(&_ExchangeV2.CallOpts, root, proofLength)
}

// IsAffiliateProgramActive is a free data retrieval call binding the contract method 0x67d9dd79.
//
// Solidity: function isAffiliateProgramActive() view returns(bool)
func (_ExchangeV2 *ExchangeV2Caller) IsAffiliateProgramActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "isAffiliateProgramActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAffiliateProgramActive is a free data retrieval call binding the contract method 0x67d9dd79.
//
// Solidity: function isAffiliateProgramActive() view returns(bool)
func (_ExchangeV2 *ExchangeV2Session) IsAffiliateProgramActive() (bool, error) {
	return _ExchangeV2.Contract.IsAffiliateProgramActive(&_ExchangeV2.CallOpts)
}

// IsAffiliateProgramActive is a free data retrieval call binding the contract method 0x67d9dd79.
//
// Solidity: function isAffiliateProgramActive() view returns(bool)
func (_ExchangeV2 *ExchangeV2CallerSession) IsAffiliateProgramActive() (bool, error) {
	return _ExchangeV2.Contract.IsAffiliateProgramActive(&_ExchangeV2.CallOpts)
}

// IsCurrencyAllowed is a free data retrieval call binding the contract method 0x5a195d19.
//
// Solidity: function isCurrencyAllowed(address ) view returns(bool)
func (_ExchangeV2 *ExchangeV2Caller) IsCurrencyAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "isCurrencyAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrencyAllowed is a free data retrieval call binding the contract method 0x5a195d19.
//
// Solidity: function isCurrencyAllowed(address ) view returns(bool)
func (_ExchangeV2 *ExchangeV2Session) IsCurrencyAllowed(arg0 common.Address) (bool, error) {
	return _ExchangeV2.Contract.IsCurrencyAllowed(&_ExchangeV2.CallOpts, arg0)
}

// IsCurrencyAllowed is a free data retrieval call binding the contract method 0x5a195d19.
//
// Solidity: function isCurrencyAllowed(address ) view returns(bool)
func (_ExchangeV2 *ExchangeV2CallerSession) IsCurrencyAllowed(arg0 common.Address) (bool, error) {
	return _ExchangeV2.Contract.IsCurrencyAllowed(&_ExchangeV2.CallOpts, arg0)
}

// MaxCreatorFeeBp is a free data retrieval call binding the contract method 0x6e90c014.
//
// Solidity: function maxCreatorFeeBp() view returns(uint16)
func (_ExchangeV2 *ExchangeV2Caller) MaxCreatorFeeBp(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "maxCreatorFeeBp")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MaxCreatorFeeBp is a free data retrieval call binding the contract method 0x6e90c014.
//
// Solidity: function maxCreatorFeeBp() view returns(uint16)
func (_ExchangeV2 *ExchangeV2Session) MaxCreatorFeeBp() (uint16, error) {
	return _ExchangeV2.Contract.MaxCreatorFeeBp(&_ExchangeV2.CallOpts)
}

// MaxCreatorFeeBp is a free data retrieval call binding the contract method 0x6e90c014.
//
// Solidity: function maxCreatorFeeBp() view returns(uint16)
func (_ExchangeV2 *ExchangeV2CallerSession) MaxCreatorFeeBp() (uint16, error) {
	return _ExchangeV2.Contract.MaxCreatorFeeBp(&_ExchangeV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) Owner() (common.Address, error) {
	return _ExchangeV2.Contract.Owner(&_ExchangeV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) Owner() (common.Address, error) {
	return _ExchangeV2.Contract.Owner(&_ExchangeV2.CallOpts)
}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_ExchangeV2 *ExchangeV2Caller) OwnershipStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "ownershipStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_ExchangeV2 *ExchangeV2Session) OwnershipStatus() (uint8, error) {
	return _ExchangeV2.Contract.OwnershipStatus(&_ExchangeV2.CallOpts)
}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_ExchangeV2 *ExchangeV2CallerSession) OwnershipStatus() (uint8, error) {
	return _ExchangeV2.Contract.OwnershipStatus(&_ExchangeV2.CallOpts)
}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) PotentialOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "potentialOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) PotentialOwner() (common.Address, error) {
	return _ExchangeV2.Contract.PotentialOwner(&_ExchangeV2.CallOpts)
}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) PotentialOwner() (common.Address, error) {
	return _ExchangeV2.Contract.PotentialOwner(&_ExchangeV2.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) ProtocolFeeRecipient() (common.Address, error) {
	return _ExchangeV2.Contract.ProtocolFeeRecipient(&_ExchangeV2.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _ExchangeV2.Contract.ProtocolFeeRecipient(&_ExchangeV2.CallOpts)
}

// StrategyInfo is a free data retrieval call binding the contract method 0xbb91c339.
//
// Solidity: function strategyInfo(uint256 ) view returns(bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2Caller) StrategyInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsActive              bool
	StandardProtocolFeeBp uint16
	MinTotalFeeBp         uint16
	MaxProtocolFeeBp      uint16
	Selector              [4]byte
	IsMakerBid            bool
	Implementation        common.Address
}, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "strategyInfo", arg0)

	outstruct := new(struct {
		IsActive              bool
		StandardProtocolFeeBp uint16
		MinTotalFeeBp         uint16
		MaxProtocolFeeBp      uint16
		Selector              [4]byte
		IsMakerBid            bool
		Implementation        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsActive = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.StandardProtocolFeeBp = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.MinTotalFeeBp = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.MaxProtocolFeeBp = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.Selector = *abi.ConvertType(out[4], new([4]byte)).(*[4]byte)
	outstruct.IsMakerBid = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Implementation = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StrategyInfo is a free data retrieval call binding the contract method 0xbb91c339.
//
// Solidity: function strategyInfo(uint256 ) view returns(bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2Session) StrategyInfo(arg0 *big.Int) (struct {
	IsActive              bool
	StandardProtocolFeeBp uint16
	MinTotalFeeBp         uint16
	MaxProtocolFeeBp      uint16
	Selector              [4]byte
	IsMakerBid            bool
	Implementation        common.Address
}, error) {
	return _ExchangeV2.Contract.StrategyInfo(&_ExchangeV2.CallOpts, arg0)
}

// StrategyInfo is a free data retrieval call binding the contract method 0xbb91c339.
//
// Solidity: function strategyInfo(uint256 ) view returns(bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2CallerSession) StrategyInfo(arg0 *big.Int) (struct {
	IsActive              bool
	StandardProtocolFeeBp uint16
	MinTotalFeeBp         uint16
	MaxProtocolFeeBp      uint16
	Selector              [4]byte
	IsMakerBid            bool
	Implementation        common.Address
}, error) {
	return _ExchangeV2.Contract.StrategyInfo(&_ExchangeV2.CallOpts, arg0)
}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_ExchangeV2 *ExchangeV2Caller) TransferManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "transferManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_ExchangeV2 *ExchangeV2Session) TransferManager() (common.Address, error) {
	return _ExchangeV2.Contract.TransferManager(&_ExchangeV2.CallOpts)
}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_ExchangeV2 *ExchangeV2CallerSession) TransferManager() (common.Address, error) {
	return _ExchangeV2.Contract.TransferManager(&_ExchangeV2.CallOpts)
}

// UserBidAskNonces is a free data retrieval call binding the contract method 0xa39bf29f.
//
// Solidity: function userBidAskNonces(address ) view returns(uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2Caller) UserBidAskNonces(opts *bind.CallOpts, arg0 common.Address) (struct {
	BidNonce *big.Int
	AskNonce *big.Int
}, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "userBidAskNonces", arg0)

	outstruct := new(struct {
		BidNonce *big.Int
		AskNonce *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AskNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserBidAskNonces is a free data retrieval call binding the contract method 0xa39bf29f.
//
// Solidity: function userBidAskNonces(address ) view returns(uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2Session) UserBidAskNonces(arg0 common.Address) (struct {
	BidNonce *big.Int
	AskNonce *big.Int
}, error) {
	return _ExchangeV2.Contract.UserBidAskNonces(&_ExchangeV2.CallOpts, arg0)
}

// UserBidAskNonces is a free data retrieval call binding the contract method 0xa39bf29f.
//
// Solidity: function userBidAskNonces(address ) view returns(uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2CallerSession) UserBidAskNonces(arg0 common.Address) (struct {
	BidNonce *big.Int
	AskNonce *big.Int
}, error) {
	return _ExchangeV2.Contract.UserBidAskNonces(&_ExchangeV2.CallOpts, arg0)
}

// UserOrderNonce is a free data retrieval call binding the contract method 0x20cd05c7.
//
// Solidity: function userOrderNonce(address , uint256 ) view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Caller) UserOrderNonce(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "userOrderNonce", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UserOrderNonce is a free data retrieval call binding the contract method 0x20cd05c7.
//
// Solidity: function userOrderNonce(address , uint256 ) view returns(bytes32)
func (_ExchangeV2 *ExchangeV2Session) UserOrderNonce(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ExchangeV2.Contract.UserOrderNonce(&_ExchangeV2.CallOpts, arg0, arg1)
}

// UserOrderNonce is a free data retrieval call binding the contract method 0x20cd05c7.
//
// Solidity: function userOrderNonce(address , uint256 ) view returns(bytes32)
func (_ExchangeV2 *ExchangeV2CallerSession) UserOrderNonce(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _ExchangeV2.Contract.UserOrderNonce(&_ExchangeV2.CallOpts, arg0, arg1)
}

// UserSubsetNonce is a free data retrieval call binding the contract method 0xea179b76.
//
// Solidity: function userSubsetNonce(address , uint256 ) view returns(bool)
func (_ExchangeV2 *ExchangeV2Caller) UserSubsetNonce(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _ExchangeV2.contract.Call(opts, &out, "userSubsetNonce", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserSubsetNonce is a free data retrieval call binding the contract method 0xea179b76.
//
// Solidity: function userSubsetNonce(address , uint256 ) view returns(bool)
func (_ExchangeV2 *ExchangeV2Session) UserSubsetNonce(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ExchangeV2.Contract.UserSubsetNonce(&_ExchangeV2.CallOpts, arg0, arg1)
}

// UserSubsetNonce is a free data retrieval call binding the contract method 0xea179b76.
//
// Solidity: function userSubsetNonce(address , uint256 ) view returns(bool)
func (_ExchangeV2 *ExchangeV2CallerSession) UserSubsetNonce(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ExchangeV2.Contract.UserSubsetNonce(&_ExchangeV2.CallOpts, arg0, arg1)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x91be1360.
//
// Solidity: function addStrategy(uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation) returns()
func (_ExchangeV2 *ExchangeV2Transactor) AddStrategy(opts *bind.TransactOpts, standardProtocolFeeBp uint16, minTotalFeeBp uint16, maxProtocolFeeBp uint16, selector [4]byte, isMakerBid bool, implementation common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "addStrategy", standardProtocolFeeBp, minTotalFeeBp, maxProtocolFeeBp, selector, isMakerBid, implementation)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x91be1360.
//
// Solidity: function addStrategy(uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation) returns()
func (_ExchangeV2 *ExchangeV2Session) AddStrategy(standardProtocolFeeBp uint16, minTotalFeeBp uint16, maxProtocolFeeBp uint16, selector [4]byte, isMakerBid bool, implementation common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.AddStrategy(&_ExchangeV2.TransactOpts, standardProtocolFeeBp, minTotalFeeBp, maxProtocolFeeBp, selector, isMakerBid, implementation)
}

// AddStrategy is a paid mutator transaction binding the contract method 0x91be1360.
//
// Solidity: function addStrategy(uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) AddStrategy(standardProtocolFeeBp uint16, minTotalFeeBp uint16, maxProtocolFeeBp uint16, selector [4]byte, isMakerBid bool, implementation common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.AddStrategy(&_ExchangeV2.TransactOpts, standardProtocolFeeBp, minTotalFeeBp, maxProtocolFeeBp, selector, isMakerBid, implementation)
}

// CancelOrderNonces is a paid mutator transaction binding the contract method 0x79ed31d4.
//
// Solidity: function cancelOrderNonces(uint256[] orderNonces) returns()
func (_ExchangeV2 *ExchangeV2Transactor) CancelOrderNonces(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "cancelOrderNonces", orderNonces)
}

// CancelOrderNonces is a paid mutator transaction binding the contract method 0x79ed31d4.
//
// Solidity: function cancelOrderNonces(uint256[] orderNonces) returns()
func (_ExchangeV2 *ExchangeV2Session) CancelOrderNonces(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelOrderNonces(&_ExchangeV2.TransactOpts, orderNonces)
}

// CancelOrderNonces is a paid mutator transaction binding the contract method 0x79ed31d4.
//
// Solidity: function cancelOrderNonces(uint256[] orderNonces) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) CancelOrderNonces(orderNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelOrderNonces(&_ExchangeV2.TransactOpts, orderNonces)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2Transactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2Session) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelOwnershipTransfer(&_ExchangeV2.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelOwnershipTransfer(&_ExchangeV2.TransactOpts)
}

// CancelSubsetNonces is a paid mutator transaction binding the contract method 0x134849a6.
//
// Solidity: function cancelSubsetNonces(uint256[] subsetNonces) returns()
func (_ExchangeV2 *ExchangeV2Transactor) CancelSubsetNonces(opts *bind.TransactOpts, subsetNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "cancelSubsetNonces", subsetNonces)
}

// CancelSubsetNonces is a paid mutator transaction binding the contract method 0x134849a6.
//
// Solidity: function cancelSubsetNonces(uint256[] subsetNonces) returns()
func (_ExchangeV2 *ExchangeV2Session) CancelSubsetNonces(subsetNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelSubsetNonces(&_ExchangeV2.TransactOpts, subsetNonces)
}

// CancelSubsetNonces is a paid mutator transaction binding the contract method 0x134849a6.
//
// Solidity: function cancelSubsetNonces(uint256[] subsetNonces) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) CancelSubsetNonces(subsetNonces []*big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.CancelSubsetNonces(&_ExchangeV2.TransactOpts, subsetNonces)
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2Transactor) ConfirmOwnershipRenouncement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "confirmOwnershipRenouncement")
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2Session) ConfirmOwnershipRenouncement() (*types.Transaction, error) {
	return _ExchangeV2.Contract.ConfirmOwnershipRenouncement(&_ExchangeV2.TransactOpts)
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) ConfirmOwnershipRenouncement() (*types.Transaction, error) {
	return _ExchangeV2.Contract.ConfirmOwnershipRenouncement(&_ExchangeV2.TransactOpts)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2Transactor) ConfirmOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "confirmOwnershipTransfer")
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2Session) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _ExchangeV2.Contract.ConfirmOwnershipTransfer(&_ExchangeV2.TransactOpts)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _ExchangeV2.Contract.ConfirmOwnershipTransfer(&_ExchangeV2.TransactOpts)
}

// ExecuteMultipleTakerBids is a paid mutator transaction binding the contract method 0xf4288a21.
//
// Solidity: function executeMultipleTakerBids((address,bytes)[] takerBids, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes)[] makerAsks, bytes[] makerSignatures, (bytes32,(bytes32,uint8)[])[] merkleTrees, address affiliate, bool isAtomic) payable returns()
func (_ExchangeV2 *ExchangeV2Transactor) ExecuteMultipleTakerBids(opts *bind.TransactOpts, takerBids []OrderStructsTaker, makerAsks []OrderStructsMaker, makerSignatures [][]byte, merkleTrees []OrderStructsMerkleTree, affiliate common.Address, isAtomic bool) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "executeMultipleTakerBids", takerBids, makerAsks, makerSignatures, merkleTrees, affiliate, isAtomic)
}

// ExecuteMultipleTakerBids is a paid mutator transaction binding the contract method 0xf4288a21.
//
// Solidity: function executeMultipleTakerBids((address,bytes)[] takerBids, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes)[] makerAsks, bytes[] makerSignatures, (bytes32,(bytes32,uint8)[])[] merkleTrees, address affiliate, bool isAtomic) payable returns()
func (_ExchangeV2 *ExchangeV2Session) ExecuteMultipleTakerBids(takerBids []OrderStructsTaker, makerAsks []OrderStructsMaker, makerSignatures [][]byte, merkleTrees []OrderStructsMerkleTree, affiliate common.Address, isAtomic bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteMultipleTakerBids(&_ExchangeV2.TransactOpts, takerBids, makerAsks, makerSignatures, merkleTrees, affiliate, isAtomic)
}

// ExecuteMultipleTakerBids is a paid mutator transaction binding the contract method 0xf4288a21.
//
// Solidity: function executeMultipleTakerBids((address,bytes)[] takerBids, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes)[] makerAsks, bytes[] makerSignatures, (bytes32,(bytes32,uint8)[])[] merkleTrees, address affiliate, bool isAtomic) payable returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) ExecuteMultipleTakerBids(takerBids []OrderStructsTaker, makerAsks []OrderStructsMaker, makerSignatures [][]byte, merkleTrees []OrderStructsMerkleTree, affiliate common.Address, isAtomic bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteMultipleTakerBids(&_ExchangeV2.TransactOpts, takerBids, makerAsks, makerSignatures, merkleTrees, affiliate, isAtomic)
}

// ExecuteTakerAsk is a paid mutator transaction binding the contract method 0xe72853e1.
//
// Solidity: function executeTakerAsk((address,bytes) takerAsk, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerBid, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) returns()
func (_ExchangeV2 *ExchangeV2Transactor) ExecuteTakerAsk(opts *bind.TransactOpts, takerAsk OrderStructsTaker, makerBid OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "executeTakerAsk", takerAsk, makerBid, makerSignature, merkleTree, affiliate)
}

// ExecuteTakerAsk is a paid mutator transaction binding the contract method 0xe72853e1.
//
// Solidity: function executeTakerAsk((address,bytes) takerAsk, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerBid, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) returns()
func (_ExchangeV2 *ExchangeV2Session) ExecuteTakerAsk(takerAsk OrderStructsTaker, makerBid OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteTakerAsk(&_ExchangeV2.TransactOpts, takerAsk, makerBid, makerSignature, merkleTree, affiliate)
}

// ExecuteTakerAsk is a paid mutator transaction binding the contract method 0xe72853e1.
//
// Solidity: function executeTakerAsk((address,bytes) takerAsk, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerBid, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) ExecuteTakerAsk(takerAsk OrderStructsTaker, makerBid OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteTakerAsk(&_ExchangeV2.TransactOpts, takerAsk, makerBid, makerSignature, merkleTree, affiliate)
}

// ExecuteTakerBid is a paid mutator transaction binding the contract method 0x8585ae03.
//
// Solidity: function executeTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) payable returns()
func (_ExchangeV2 *ExchangeV2Transactor) ExecuteTakerBid(opts *bind.TransactOpts, takerBid OrderStructsTaker, makerAsk OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "executeTakerBid", takerBid, makerAsk, makerSignature, merkleTree, affiliate)
}

// ExecuteTakerBid is a paid mutator transaction binding the contract method 0x8585ae03.
//
// Solidity: function executeTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) payable returns()
func (_ExchangeV2 *ExchangeV2Session) ExecuteTakerBid(takerBid OrderStructsTaker, makerAsk OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteTakerBid(&_ExchangeV2.TransactOpts, takerBid, makerAsk, makerSignature, merkleTree, affiliate)
}

// ExecuteTakerBid is a paid mutator transaction binding the contract method 0x8585ae03.
//
// Solidity: function executeTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, bytes makerSignature, (bytes32,(bytes32,uint8)[]) merkleTree, address affiliate) payable returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) ExecuteTakerBid(takerBid OrderStructsTaker, makerAsk OrderStructsMaker, makerSignature []byte, merkleTree OrderStructsMerkleTree, affiliate common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.ExecuteTakerBid(&_ExchangeV2.TransactOpts, takerBid, makerAsk, makerSignature, merkleTree, affiliate)
}

// IncrementBidAskNonces is a paid mutator transaction binding the contract method 0xd5b010f5.
//
// Solidity: function incrementBidAskNonces(bool bid, bool ask) returns()
func (_ExchangeV2 *ExchangeV2Transactor) IncrementBidAskNonces(opts *bind.TransactOpts, bid bool, ask bool) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "incrementBidAskNonces", bid, ask)
}

// IncrementBidAskNonces is a paid mutator transaction binding the contract method 0xd5b010f5.
//
// Solidity: function incrementBidAskNonces(bool bid, bool ask) returns()
func (_ExchangeV2 *ExchangeV2Session) IncrementBidAskNonces(bid bool, ask bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.IncrementBidAskNonces(&_ExchangeV2.TransactOpts, bid, ask)
}

// IncrementBidAskNonces is a paid mutator transaction binding the contract method 0xd5b010f5.
//
// Solidity: function incrementBidAskNonces(bool bid, bool ask) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) IncrementBidAskNonces(bid bool, ask bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.IncrementBidAskNonces(&_ExchangeV2.TransactOpts, bid, ask)
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2Transactor) InitiateOwnershipRenouncement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "initiateOwnershipRenouncement")
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2Session) InitiateOwnershipRenouncement() (*types.Transaction, error) {
	return _ExchangeV2.Contract.InitiateOwnershipRenouncement(&_ExchangeV2.TransactOpts)
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) InitiateOwnershipRenouncement() (*types.Transaction, error) {
	return _ExchangeV2.Contract.InitiateOwnershipRenouncement(&_ExchangeV2.TransactOpts)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_ExchangeV2 *ExchangeV2Transactor) InitiateOwnershipTransfer(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "initiateOwnershipTransfer", newPotentialOwner)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_ExchangeV2 *ExchangeV2Session) InitiateOwnershipTransfer(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.InitiateOwnershipTransfer(&_ExchangeV2.TransactOpts, newPotentialOwner)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) InitiateOwnershipTransfer(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.InitiateOwnershipTransfer(&_ExchangeV2.TransactOpts, newPotentialOwner)
}

// RestrictedExecuteTakerBid is a paid mutator transaction binding the contract method 0xd5b7f065.
//
// Solidity: function restrictedExecuteTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, address sender, bytes32 orderHash) returns(uint256 protocolFeeAmount)
func (_ExchangeV2 *ExchangeV2Transactor) RestrictedExecuteTakerBid(opts *bind.TransactOpts, takerBid OrderStructsTaker, makerAsk OrderStructsMaker, sender common.Address, orderHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "restrictedExecuteTakerBid", takerBid, makerAsk, sender, orderHash)
}

// RestrictedExecuteTakerBid is a paid mutator transaction binding the contract method 0xd5b7f065.
//
// Solidity: function restrictedExecuteTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, address sender, bytes32 orderHash) returns(uint256 protocolFeeAmount)
func (_ExchangeV2 *ExchangeV2Session) RestrictedExecuteTakerBid(takerBid OrderStructsTaker, makerAsk OrderStructsMaker, sender common.Address, orderHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV2.Contract.RestrictedExecuteTakerBid(&_ExchangeV2.TransactOpts, takerBid, makerAsk, sender, orderHash)
}

// RestrictedExecuteTakerBid is a paid mutator transaction binding the contract method 0xd5b7f065.
//
// Solidity: function restrictedExecuteTakerBid((address,bytes) takerBid, (uint8,uint256,uint256,uint256,uint256,uint8,address,address,address,uint256,uint256,uint256,uint256[],uint256[],bytes) makerAsk, address sender, bytes32 orderHash) returns(uint256 protocolFeeAmount)
func (_ExchangeV2 *ExchangeV2TransactorSession) RestrictedExecuteTakerBid(takerBid OrderStructsTaker, makerAsk OrderStructsMaker, sender common.Address, orderHash [32]byte) (*types.Transaction, error) {
	return _ExchangeV2.Contract.RestrictedExecuteTakerBid(&_ExchangeV2.TransactOpts, takerBid, makerAsk, sender, orderHash)
}

// UpdateAffiliateController is a paid mutator transaction binding the contract method 0xa02bab57.
//
// Solidity: function updateAffiliateController(address newAffiliateController) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateAffiliateController(opts *bind.TransactOpts, newAffiliateController common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateAffiliateController", newAffiliateController)
}

// UpdateAffiliateController is a paid mutator transaction binding the contract method 0xa02bab57.
//
// Solidity: function updateAffiliateController(address newAffiliateController) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateAffiliateController(newAffiliateController common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateController(&_ExchangeV2.TransactOpts, newAffiliateController)
}

// UpdateAffiliateController is a paid mutator transaction binding the contract method 0xa02bab57.
//
// Solidity: function updateAffiliateController(address newAffiliateController) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateAffiliateController(newAffiliateController common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateController(&_ExchangeV2.TransactOpts, newAffiliateController)
}

// UpdateAffiliateProgramStatus is a paid mutator transaction binding the contract method 0x54878876.
//
// Solidity: function updateAffiliateProgramStatus(bool isActive) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateAffiliateProgramStatus(opts *bind.TransactOpts, isActive bool) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateAffiliateProgramStatus", isActive)
}

// UpdateAffiliateProgramStatus is a paid mutator transaction binding the contract method 0x54878876.
//
// Solidity: function updateAffiliateProgramStatus(bool isActive) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateAffiliateProgramStatus(isActive bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateProgramStatus(&_ExchangeV2.TransactOpts, isActive)
}

// UpdateAffiliateProgramStatus is a paid mutator transaction binding the contract method 0x54878876.
//
// Solidity: function updateAffiliateProgramStatus(bool isActive) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateAffiliateProgramStatus(isActive bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateProgramStatus(&_ExchangeV2.TransactOpts, isActive)
}

// UpdateAffiliateRate is a paid mutator transaction binding the contract method 0xae1cce5a.
//
// Solidity: function updateAffiliateRate(address affiliate, uint256 bp) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateAffiliateRate(opts *bind.TransactOpts, affiliate common.Address, bp *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateAffiliateRate", affiliate, bp)
}

// UpdateAffiliateRate is a paid mutator transaction binding the contract method 0xae1cce5a.
//
// Solidity: function updateAffiliateRate(address affiliate, uint256 bp) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateAffiliateRate(affiliate common.Address, bp *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateRate(&_ExchangeV2.TransactOpts, affiliate, bp)
}

// UpdateAffiliateRate is a paid mutator transaction binding the contract method 0xae1cce5a.
//
// Solidity: function updateAffiliateRate(address affiliate, uint256 bp) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateAffiliateRate(affiliate common.Address, bp *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateAffiliateRate(&_ExchangeV2.TransactOpts, affiliate, bp)
}

// UpdateCreatorFeeManager is a paid mutator transaction binding the contract method 0xb647a404.
//
// Solidity: function updateCreatorFeeManager(address newCreatorFeeManager) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateCreatorFeeManager(opts *bind.TransactOpts, newCreatorFeeManager common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateCreatorFeeManager", newCreatorFeeManager)
}

// UpdateCreatorFeeManager is a paid mutator transaction binding the contract method 0xb647a404.
//
// Solidity: function updateCreatorFeeManager(address newCreatorFeeManager) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateCreatorFeeManager(newCreatorFeeManager common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateCreatorFeeManager(&_ExchangeV2.TransactOpts, newCreatorFeeManager)
}

// UpdateCreatorFeeManager is a paid mutator transaction binding the contract method 0xb647a404.
//
// Solidity: function updateCreatorFeeManager(address newCreatorFeeManager) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateCreatorFeeManager(newCreatorFeeManager common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateCreatorFeeManager(&_ExchangeV2.TransactOpts, newCreatorFeeManager)
}

// UpdateCurrencyStatus is a paid mutator transaction binding the contract method 0x1d3c4268.
//
// Solidity: function updateCurrencyStatus(address currency, bool isAllowed) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateCurrencyStatus(opts *bind.TransactOpts, currency common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateCurrencyStatus", currency, isAllowed)
}

// UpdateCurrencyStatus is a paid mutator transaction binding the contract method 0x1d3c4268.
//
// Solidity: function updateCurrencyStatus(address currency, bool isAllowed) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateCurrencyStatus(currency common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateCurrencyStatus(&_ExchangeV2.TransactOpts, currency, isAllowed)
}

// UpdateCurrencyStatus is a paid mutator transaction binding the contract method 0x1d3c4268.
//
// Solidity: function updateCurrencyStatus(address currency, bool isAllowed) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateCurrencyStatus(currency common.Address, isAllowed bool) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateCurrencyStatus(&_ExchangeV2.TransactOpts, currency, isAllowed)
}

// UpdateDomainSeparator is a paid mutator transaction binding the contract method 0x89ccfe89.
//
// Solidity: function updateDomainSeparator() returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateDomainSeparator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateDomainSeparator")
}

// UpdateDomainSeparator is a paid mutator transaction binding the contract method 0x89ccfe89.
//
// Solidity: function updateDomainSeparator() returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateDomainSeparator() (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateDomainSeparator(&_ExchangeV2.TransactOpts)
}

// UpdateDomainSeparator is a paid mutator transaction binding the contract method 0x89ccfe89.
//
// Solidity: function updateDomainSeparator() returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateDomainSeparator() (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateDomainSeparator(&_ExchangeV2.TransactOpts)
}

// UpdateETHGasLimitForTransfer is a paid mutator transaction binding the contract method 0x974e7c9f.
//
// Solidity: function updateETHGasLimitForTransfer(uint256 newGasLimitETHTransfer) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateETHGasLimitForTransfer(opts *bind.TransactOpts, newGasLimitETHTransfer *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateETHGasLimitForTransfer", newGasLimitETHTransfer)
}

// UpdateETHGasLimitForTransfer is a paid mutator transaction binding the contract method 0x974e7c9f.
//
// Solidity: function updateETHGasLimitForTransfer(uint256 newGasLimitETHTransfer) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateETHGasLimitForTransfer(newGasLimitETHTransfer *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateETHGasLimitForTransfer(&_ExchangeV2.TransactOpts, newGasLimitETHTransfer)
}

// UpdateETHGasLimitForTransfer is a paid mutator transaction binding the contract method 0x974e7c9f.
//
// Solidity: function updateETHGasLimitForTransfer(uint256 newGasLimitETHTransfer) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateETHGasLimitForTransfer(newGasLimitETHTransfer *big.Int) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateETHGasLimitForTransfer(&_ExchangeV2.TransactOpts, newGasLimitETHTransfer)
}

// UpdateMaxCreatorFeeBp is a paid mutator transaction binding the contract method 0x46b625bc.
//
// Solidity: function updateMaxCreatorFeeBp(uint16 newMaxCreatorFeeBp) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateMaxCreatorFeeBp(opts *bind.TransactOpts, newMaxCreatorFeeBp uint16) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateMaxCreatorFeeBp", newMaxCreatorFeeBp)
}

// UpdateMaxCreatorFeeBp is a paid mutator transaction binding the contract method 0x46b625bc.
//
// Solidity: function updateMaxCreatorFeeBp(uint16 newMaxCreatorFeeBp) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateMaxCreatorFeeBp(newMaxCreatorFeeBp uint16) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateMaxCreatorFeeBp(&_ExchangeV2.TransactOpts, newMaxCreatorFeeBp)
}

// UpdateMaxCreatorFeeBp is a paid mutator transaction binding the contract method 0x46b625bc.
//
// Solidity: function updateMaxCreatorFeeBp(uint16 newMaxCreatorFeeBp) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateMaxCreatorFeeBp(newMaxCreatorFeeBp uint16) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateMaxCreatorFeeBp(&_ExchangeV2.TransactOpts, newMaxCreatorFeeBp)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateProtocolFeeRecipient(opts *bind.TransactOpts, newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateProtocolFeeRecipient", newProtocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateProtocolFeeRecipient(&_ExchangeV2.TransactOpts, newProtocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address newProtocolFeeRecipient) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateProtocolFeeRecipient(newProtocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateProtocolFeeRecipient(&_ExchangeV2.TransactOpts, newProtocolFeeRecipient)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0xd5a06adf.
//
// Solidity: function updateStrategy(uint256 strategyId, bool isActive, uint16 newStandardProtocolFee, uint16 newMinTotalFee) returns()
func (_ExchangeV2 *ExchangeV2Transactor) UpdateStrategy(opts *bind.TransactOpts, strategyId *big.Int, isActive bool, newStandardProtocolFee uint16, newMinTotalFee uint16) (*types.Transaction, error) {
	return _ExchangeV2.contract.Transact(opts, "updateStrategy", strategyId, isActive, newStandardProtocolFee, newMinTotalFee)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0xd5a06adf.
//
// Solidity: function updateStrategy(uint256 strategyId, bool isActive, uint16 newStandardProtocolFee, uint16 newMinTotalFee) returns()
func (_ExchangeV2 *ExchangeV2Session) UpdateStrategy(strategyId *big.Int, isActive bool, newStandardProtocolFee uint16, newMinTotalFee uint16) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateStrategy(&_ExchangeV2.TransactOpts, strategyId, isActive, newStandardProtocolFee, newMinTotalFee)
}

// UpdateStrategy is a paid mutator transaction binding the contract method 0xd5a06adf.
//
// Solidity: function updateStrategy(uint256 strategyId, bool isActive, uint16 newStandardProtocolFee, uint16 newMinTotalFee) returns()
func (_ExchangeV2 *ExchangeV2TransactorSession) UpdateStrategy(strategyId *big.Int, isActive bool, newStandardProtocolFee uint16, newMinTotalFee uint16) (*types.Transaction, error) {
	return _ExchangeV2.Contract.UpdateStrategy(&_ExchangeV2.TransactOpts, strategyId, isActive, newStandardProtocolFee, newMinTotalFee)
}

// ExchangeV2AffiliatePaymentIterator is returned from FilterAffiliatePayment and is used to iterate over the raw logs and unpacked data for AffiliatePayment events raised by the ExchangeV2 contract.
type ExchangeV2AffiliatePaymentIterator struct {
	Event *ExchangeV2AffiliatePayment // Event containing the contract specifics and raw log

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
func (it *ExchangeV2AffiliatePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2AffiliatePayment)
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
		it.Event = new(ExchangeV2AffiliatePayment)
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
func (it *ExchangeV2AffiliatePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2AffiliatePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2AffiliatePayment represents a AffiliatePayment event raised by the ExchangeV2 contract.
type ExchangeV2AffiliatePayment struct {
	Affiliate    common.Address
	Currency     common.Address
	AffiliateFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAffiliatePayment is a free log retrieval operation binding the contract event 0x49e92b6b3114f7e128555cd58c568f7c2e0e56fe9b4c46b6125bc308184623b3.
//
// Solidity: event AffiliatePayment(address affiliate, address currency, uint256 affiliateFee)
func (_ExchangeV2 *ExchangeV2Filterer) FilterAffiliatePayment(opts *bind.FilterOpts) (*ExchangeV2AffiliatePaymentIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "AffiliatePayment")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2AffiliatePaymentIterator{contract: _ExchangeV2.contract, event: "AffiliatePayment", logs: logs, sub: sub}, nil
}

// WatchAffiliatePayment is a free log subscription operation binding the contract event 0x49e92b6b3114f7e128555cd58c568f7c2e0e56fe9b4c46b6125bc308184623b3.
//
// Solidity: event AffiliatePayment(address affiliate, address currency, uint256 affiliateFee)
func (_ExchangeV2 *ExchangeV2Filterer) WatchAffiliatePayment(opts *bind.WatchOpts, sink chan<- *ExchangeV2AffiliatePayment) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "AffiliatePayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2AffiliatePayment)
				if err := _ExchangeV2.contract.UnpackLog(event, "AffiliatePayment", log); err != nil {
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

// ParseAffiliatePayment is a log parse operation binding the contract event 0x49e92b6b3114f7e128555cd58c568f7c2e0e56fe9b4c46b6125bc308184623b3.
//
// Solidity: event AffiliatePayment(address affiliate, address currency, uint256 affiliateFee)
func (_ExchangeV2 *ExchangeV2Filterer) ParseAffiliatePayment(log types.Log) (*ExchangeV2AffiliatePayment, error) {
	event := new(ExchangeV2AffiliatePayment)
	if err := _ExchangeV2.contract.UnpackLog(event, "AffiliatePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2CancelOwnershipTransferIterator is returned from FilterCancelOwnershipTransfer and is used to iterate over the raw logs and unpacked data for CancelOwnershipTransfer events raised by the ExchangeV2 contract.
type ExchangeV2CancelOwnershipTransferIterator struct {
	Event *ExchangeV2CancelOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeV2CancelOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2CancelOwnershipTransfer)
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
		it.Event = new(ExchangeV2CancelOwnershipTransfer)
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
func (it *ExchangeV2CancelOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2CancelOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2CancelOwnershipTransfer represents a CancelOwnershipTransfer event raised by the ExchangeV2 contract.
type ExchangeV2CancelOwnershipTransfer struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelOwnershipTransfer is a free log retrieval operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_ExchangeV2 *ExchangeV2Filterer) FilterCancelOwnershipTransfer(opts *bind.FilterOpts) (*ExchangeV2CancelOwnershipTransferIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "CancelOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2CancelOwnershipTransferIterator{contract: _ExchangeV2.contract, event: "CancelOwnershipTransfer", logs: logs, sub: sub}, nil
}

// WatchCancelOwnershipTransfer is a free log subscription operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_ExchangeV2 *ExchangeV2Filterer) WatchCancelOwnershipTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeV2CancelOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "CancelOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2CancelOwnershipTransfer)
				if err := _ExchangeV2.contract.UnpackLog(event, "CancelOwnershipTransfer", log); err != nil {
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

// ParseCancelOwnershipTransfer is a log parse operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_ExchangeV2 *ExchangeV2Filterer) ParseCancelOwnershipTransfer(log types.Log) (*ExchangeV2CancelOwnershipTransfer, error) {
	event := new(ExchangeV2CancelOwnershipTransfer)
	if err := _ExchangeV2.contract.UnpackLog(event, "CancelOwnershipTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2CurrencyStatusUpdatedIterator is returned from FilterCurrencyStatusUpdated and is used to iterate over the raw logs and unpacked data for CurrencyStatusUpdated events raised by the ExchangeV2 contract.
type ExchangeV2CurrencyStatusUpdatedIterator struct {
	Event *ExchangeV2CurrencyStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangeV2CurrencyStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2CurrencyStatusUpdated)
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
		it.Event = new(ExchangeV2CurrencyStatusUpdated)
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
func (it *ExchangeV2CurrencyStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2CurrencyStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2CurrencyStatusUpdated represents a CurrencyStatusUpdated event raised by the ExchangeV2 contract.
type ExchangeV2CurrencyStatusUpdated struct {
	Currency  common.Address
	IsAllowed bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCurrencyStatusUpdated is a free log retrieval operation binding the contract event 0xba28eda47a2e15b1dd3269e6d82f66730d20a5661aa40e9faf9f311c7872a543.
//
// Solidity: event CurrencyStatusUpdated(address currency, bool isAllowed)
func (_ExchangeV2 *ExchangeV2Filterer) FilterCurrencyStatusUpdated(opts *bind.FilterOpts) (*ExchangeV2CurrencyStatusUpdatedIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "CurrencyStatusUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2CurrencyStatusUpdatedIterator{contract: _ExchangeV2.contract, event: "CurrencyStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchCurrencyStatusUpdated is a free log subscription operation binding the contract event 0xba28eda47a2e15b1dd3269e6d82f66730d20a5661aa40e9faf9f311c7872a543.
//
// Solidity: event CurrencyStatusUpdated(address currency, bool isAllowed)
func (_ExchangeV2 *ExchangeV2Filterer) WatchCurrencyStatusUpdated(opts *bind.WatchOpts, sink chan<- *ExchangeV2CurrencyStatusUpdated) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "CurrencyStatusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2CurrencyStatusUpdated)
				if err := _ExchangeV2.contract.UnpackLog(event, "CurrencyStatusUpdated", log); err != nil {
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

// ParseCurrencyStatusUpdated is a log parse operation binding the contract event 0xba28eda47a2e15b1dd3269e6d82f66730d20a5661aa40e9faf9f311c7872a543.
//
// Solidity: event CurrencyStatusUpdated(address currency, bool isAllowed)
func (_ExchangeV2 *ExchangeV2Filterer) ParseCurrencyStatusUpdated(log types.Log) (*ExchangeV2CurrencyStatusUpdated, error) {
	event := new(ExchangeV2CurrencyStatusUpdated)
	if err := _ExchangeV2.contract.UnpackLog(event, "CurrencyStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2InitiateOwnershipRenouncementIterator is returned from FilterInitiateOwnershipRenouncement and is used to iterate over the raw logs and unpacked data for InitiateOwnershipRenouncement events raised by the ExchangeV2 contract.
type ExchangeV2InitiateOwnershipRenouncementIterator struct {
	Event *ExchangeV2InitiateOwnershipRenouncement // Event containing the contract specifics and raw log

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
func (it *ExchangeV2InitiateOwnershipRenouncementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2InitiateOwnershipRenouncement)
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
		it.Event = new(ExchangeV2InitiateOwnershipRenouncement)
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
func (it *ExchangeV2InitiateOwnershipRenouncementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2InitiateOwnershipRenouncementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2InitiateOwnershipRenouncement represents a InitiateOwnershipRenouncement event raised by the ExchangeV2 contract.
type ExchangeV2InitiateOwnershipRenouncement struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitiateOwnershipRenouncement is a free log retrieval operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_ExchangeV2 *ExchangeV2Filterer) FilterInitiateOwnershipRenouncement(opts *bind.FilterOpts) (*ExchangeV2InitiateOwnershipRenouncementIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "InitiateOwnershipRenouncement")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2InitiateOwnershipRenouncementIterator{contract: _ExchangeV2.contract, event: "InitiateOwnershipRenouncement", logs: logs, sub: sub}, nil
}

// WatchInitiateOwnershipRenouncement is a free log subscription operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_ExchangeV2 *ExchangeV2Filterer) WatchInitiateOwnershipRenouncement(opts *bind.WatchOpts, sink chan<- *ExchangeV2InitiateOwnershipRenouncement) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "InitiateOwnershipRenouncement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2InitiateOwnershipRenouncement)
				if err := _ExchangeV2.contract.UnpackLog(event, "InitiateOwnershipRenouncement", log); err != nil {
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

// ParseInitiateOwnershipRenouncement is a log parse operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_ExchangeV2 *ExchangeV2Filterer) ParseInitiateOwnershipRenouncement(log types.Log) (*ExchangeV2InitiateOwnershipRenouncement, error) {
	event := new(ExchangeV2InitiateOwnershipRenouncement)
	if err := _ExchangeV2.contract.UnpackLog(event, "InitiateOwnershipRenouncement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2InitiateOwnershipTransferIterator is returned from FilterInitiateOwnershipTransfer and is used to iterate over the raw logs and unpacked data for InitiateOwnershipTransfer events raised by the ExchangeV2 contract.
type ExchangeV2InitiateOwnershipTransferIterator struct {
	Event *ExchangeV2InitiateOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeV2InitiateOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2InitiateOwnershipTransfer)
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
		it.Event = new(ExchangeV2InitiateOwnershipTransfer)
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
func (it *ExchangeV2InitiateOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2InitiateOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2InitiateOwnershipTransfer represents a InitiateOwnershipTransfer event raised by the ExchangeV2 contract.
type ExchangeV2InitiateOwnershipTransfer struct {
	PreviousOwner  common.Address
	PotentialOwner common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiateOwnershipTransfer is a free log retrieval operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_ExchangeV2 *ExchangeV2Filterer) FilterInitiateOwnershipTransfer(opts *bind.FilterOpts) (*ExchangeV2InitiateOwnershipTransferIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "InitiateOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2InitiateOwnershipTransferIterator{contract: _ExchangeV2.contract, event: "InitiateOwnershipTransfer", logs: logs, sub: sub}, nil
}

// WatchInitiateOwnershipTransfer is a free log subscription operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_ExchangeV2 *ExchangeV2Filterer) WatchInitiateOwnershipTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeV2InitiateOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "InitiateOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2InitiateOwnershipTransfer)
				if err := _ExchangeV2.contract.UnpackLog(event, "InitiateOwnershipTransfer", log); err != nil {
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

// ParseInitiateOwnershipTransfer is a log parse operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_ExchangeV2 *ExchangeV2Filterer) ParseInitiateOwnershipTransfer(log types.Log) (*ExchangeV2InitiateOwnershipTransfer, error) {
	event := new(ExchangeV2InitiateOwnershipTransfer)
	if err := _ExchangeV2.contract.UnpackLog(event, "InitiateOwnershipTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewAffiliateControllerIterator is returned from FilterNewAffiliateController and is used to iterate over the raw logs and unpacked data for NewAffiliateController events raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateControllerIterator struct {
	Event *ExchangeV2NewAffiliateController // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewAffiliateControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewAffiliateController)
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
		it.Event = new(ExchangeV2NewAffiliateController)
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
func (it *ExchangeV2NewAffiliateControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewAffiliateControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewAffiliateController represents a NewAffiliateController event raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateController struct {
	AffiliateController common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewAffiliateController is a free log retrieval operation binding the contract event 0xa92d85531a006d07fd0df4c61259b6dc18e4f492857e2454e5a20ee5e55cddcc.
//
// Solidity: event NewAffiliateController(address affiliateController)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewAffiliateController(opts *bind.FilterOpts) (*ExchangeV2NewAffiliateControllerIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewAffiliateController")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewAffiliateControllerIterator{contract: _ExchangeV2.contract, event: "NewAffiliateController", logs: logs, sub: sub}, nil
}

// WatchNewAffiliateController is a free log subscription operation binding the contract event 0xa92d85531a006d07fd0df4c61259b6dc18e4f492857e2454e5a20ee5e55cddcc.
//
// Solidity: event NewAffiliateController(address affiliateController)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewAffiliateController(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewAffiliateController) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewAffiliateController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewAffiliateController)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateController", log); err != nil {
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

// ParseNewAffiliateController is a log parse operation binding the contract event 0xa92d85531a006d07fd0df4c61259b6dc18e4f492857e2454e5a20ee5e55cddcc.
//
// Solidity: event NewAffiliateController(address affiliateController)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewAffiliateController(log types.Log) (*ExchangeV2NewAffiliateController, error) {
	event := new(ExchangeV2NewAffiliateController)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewAffiliateProgramStatusIterator is returned from FilterNewAffiliateProgramStatus and is used to iterate over the raw logs and unpacked data for NewAffiliateProgramStatus events raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateProgramStatusIterator struct {
	Event *ExchangeV2NewAffiliateProgramStatus // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewAffiliateProgramStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewAffiliateProgramStatus)
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
		it.Event = new(ExchangeV2NewAffiliateProgramStatus)
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
func (it *ExchangeV2NewAffiliateProgramStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewAffiliateProgramStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewAffiliateProgramStatus represents a NewAffiliateProgramStatus event raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateProgramStatus struct {
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAffiliateProgramStatus is a free log retrieval operation binding the contract event 0xdca612ba3556b7c2603089071be1feb2404df55dcabceee9d5fd852fdb39bc34.
//
// Solidity: event NewAffiliateProgramStatus(bool isActive)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewAffiliateProgramStatus(opts *bind.FilterOpts) (*ExchangeV2NewAffiliateProgramStatusIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewAffiliateProgramStatus")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewAffiliateProgramStatusIterator{contract: _ExchangeV2.contract, event: "NewAffiliateProgramStatus", logs: logs, sub: sub}, nil
}

// WatchNewAffiliateProgramStatus is a free log subscription operation binding the contract event 0xdca612ba3556b7c2603089071be1feb2404df55dcabceee9d5fd852fdb39bc34.
//
// Solidity: event NewAffiliateProgramStatus(bool isActive)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewAffiliateProgramStatus(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewAffiliateProgramStatus) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewAffiliateProgramStatus")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewAffiliateProgramStatus)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateProgramStatus", log); err != nil {
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

// ParseNewAffiliateProgramStatus is a log parse operation binding the contract event 0xdca612ba3556b7c2603089071be1feb2404df55dcabceee9d5fd852fdb39bc34.
//
// Solidity: event NewAffiliateProgramStatus(bool isActive)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewAffiliateProgramStatus(log types.Log) (*ExchangeV2NewAffiliateProgramStatus, error) {
	event := new(ExchangeV2NewAffiliateProgramStatus)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateProgramStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewAffiliateRateIterator is returned from FilterNewAffiliateRate and is used to iterate over the raw logs and unpacked data for NewAffiliateRate events raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateRateIterator struct {
	Event *ExchangeV2NewAffiliateRate // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewAffiliateRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewAffiliateRate)
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
		it.Event = new(ExchangeV2NewAffiliateRate)
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
func (it *ExchangeV2NewAffiliateRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewAffiliateRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewAffiliateRate represents a NewAffiliateRate event raised by the ExchangeV2 contract.
type ExchangeV2NewAffiliateRate struct {
	Affiliate common.Address
	Rate      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewAffiliateRate is a free log retrieval operation binding the contract event 0xa16126d9473196242b0784325b30642b89f34442dd178a852f5b88ee483a30d8.
//
// Solidity: event NewAffiliateRate(address affiliate, uint256 rate)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewAffiliateRate(opts *bind.FilterOpts) (*ExchangeV2NewAffiliateRateIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewAffiliateRate")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewAffiliateRateIterator{contract: _ExchangeV2.contract, event: "NewAffiliateRate", logs: logs, sub: sub}, nil
}

// WatchNewAffiliateRate is a free log subscription operation binding the contract event 0xa16126d9473196242b0784325b30642b89f34442dd178a852f5b88ee483a30d8.
//
// Solidity: event NewAffiliateRate(address affiliate, uint256 rate)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewAffiliateRate(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewAffiliateRate) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewAffiliateRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewAffiliateRate)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateRate", log); err != nil {
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

// ParseNewAffiliateRate is a log parse operation binding the contract event 0xa16126d9473196242b0784325b30642b89f34442dd178a852f5b88ee483a30d8.
//
// Solidity: event NewAffiliateRate(address affiliate, uint256 rate)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewAffiliateRate(log types.Log) (*ExchangeV2NewAffiliateRate, error) {
	event := new(ExchangeV2NewAffiliateRate)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewAffiliateRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewBidAskNoncesIterator is returned from FilterNewBidAskNonces and is used to iterate over the raw logs and unpacked data for NewBidAskNonces events raised by the ExchangeV2 contract.
type ExchangeV2NewBidAskNoncesIterator struct {
	Event *ExchangeV2NewBidAskNonces // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewBidAskNoncesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewBidAskNonces)
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
		it.Event = new(ExchangeV2NewBidAskNonces)
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
func (it *ExchangeV2NewBidAskNoncesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewBidAskNoncesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewBidAskNonces represents a NewBidAskNonces event raised by the ExchangeV2 contract.
type ExchangeV2NewBidAskNonces struct {
	User     common.Address
	BidNonce *big.Int
	AskNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewBidAskNonces is a free log retrieval operation binding the contract event 0xb738dd6073fae1a7128e3fcc6b4ca6e1356b7232f87cc98f8a2857bcd83dfc44.
//
// Solidity: event NewBidAskNonces(address user, uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewBidAskNonces(opts *bind.FilterOpts) (*ExchangeV2NewBidAskNoncesIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewBidAskNonces")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewBidAskNoncesIterator{contract: _ExchangeV2.contract, event: "NewBidAskNonces", logs: logs, sub: sub}, nil
}

// WatchNewBidAskNonces is a free log subscription operation binding the contract event 0xb738dd6073fae1a7128e3fcc6b4ca6e1356b7232f87cc98f8a2857bcd83dfc44.
//
// Solidity: event NewBidAskNonces(address user, uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewBidAskNonces(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewBidAskNonces) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewBidAskNonces")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewBidAskNonces)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewBidAskNonces", log); err != nil {
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

// ParseNewBidAskNonces is a log parse operation binding the contract event 0xb738dd6073fae1a7128e3fcc6b4ca6e1356b7232f87cc98f8a2857bcd83dfc44.
//
// Solidity: event NewBidAskNonces(address user, uint256 bidNonce, uint256 askNonce)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewBidAskNonces(log types.Log) (*ExchangeV2NewBidAskNonces, error) {
	event := new(ExchangeV2NewBidAskNonces)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewBidAskNonces", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewCreatorFeeManagerIterator is returned from FilterNewCreatorFeeManager and is used to iterate over the raw logs and unpacked data for NewCreatorFeeManager events raised by the ExchangeV2 contract.
type ExchangeV2NewCreatorFeeManagerIterator struct {
	Event *ExchangeV2NewCreatorFeeManager // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewCreatorFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewCreatorFeeManager)
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
		it.Event = new(ExchangeV2NewCreatorFeeManager)
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
func (it *ExchangeV2NewCreatorFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewCreatorFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewCreatorFeeManager represents a NewCreatorFeeManager event raised by the ExchangeV2 contract.
type ExchangeV2NewCreatorFeeManager struct {
	CreatorFeeManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewCreatorFeeManager is a free log retrieval operation binding the contract event 0xa7014d98341b07b23615cb6b4da7cca2a381932b46fb39ca4b8c3875c53aa764.
//
// Solidity: event NewCreatorFeeManager(address creatorFeeManager)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewCreatorFeeManager(opts *bind.FilterOpts) (*ExchangeV2NewCreatorFeeManagerIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewCreatorFeeManager")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewCreatorFeeManagerIterator{contract: _ExchangeV2.contract, event: "NewCreatorFeeManager", logs: logs, sub: sub}, nil
}

// WatchNewCreatorFeeManager is a free log subscription operation binding the contract event 0xa7014d98341b07b23615cb6b4da7cca2a381932b46fb39ca4b8c3875c53aa764.
//
// Solidity: event NewCreatorFeeManager(address creatorFeeManager)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewCreatorFeeManager(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewCreatorFeeManager) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewCreatorFeeManager")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewCreatorFeeManager)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewCreatorFeeManager", log); err != nil {
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

// ParseNewCreatorFeeManager is a log parse operation binding the contract event 0xa7014d98341b07b23615cb6b4da7cca2a381932b46fb39ca4b8c3875c53aa764.
//
// Solidity: event NewCreatorFeeManager(address creatorFeeManager)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewCreatorFeeManager(log types.Log) (*ExchangeV2NewCreatorFeeManager, error) {
	event := new(ExchangeV2NewCreatorFeeManager)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewCreatorFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewDomainSeparatorIterator is returned from FilterNewDomainSeparator and is used to iterate over the raw logs and unpacked data for NewDomainSeparator events raised by the ExchangeV2 contract.
type ExchangeV2NewDomainSeparatorIterator struct {
	Event *ExchangeV2NewDomainSeparator // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewDomainSeparatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewDomainSeparator)
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
		it.Event = new(ExchangeV2NewDomainSeparator)
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
func (it *ExchangeV2NewDomainSeparatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewDomainSeparatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewDomainSeparator represents a NewDomainSeparator event raised by the ExchangeV2 contract.
type ExchangeV2NewDomainSeparator struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewDomainSeparator is a free log retrieval operation binding the contract event 0x5b2d1f36cd3ec425baab07b99853532e2ba6387a472ddbff437c5cc96f2f20ca.
//
// Solidity: event NewDomainSeparator()
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewDomainSeparator(opts *bind.FilterOpts) (*ExchangeV2NewDomainSeparatorIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewDomainSeparator")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewDomainSeparatorIterator{contract: _ExchangeV2.contract, event: "NewDomainSeparator", logs: logs, sub: sub}, nil
}

// WatchNewDomainSeparator is a free log subscription operation binding the contract event 0x5b2d1f36cd3ec425baab07b99853532e2ba6387a472ddbff437c5cc96f2f20ca.
//
// Solidity: event NewDomainSeparator()
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewDomainSeparator(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewDomainSeparator) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewDomainSeparator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewDomainSeparator)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewDomainSeparator", log); err != nil {
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

// ParseNewDomainSeparator is a log parse operation binding the contract event 0x5b2d1f36cd3ec425baab07b99853532e2ba6387a472ddbff437c5cc96f2f20ca.
//
// Solidity: event NewDomainSeparator()
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewDomainSeparator(log types.Log) (*ExchangeV2NewDomainSeparator, error) {
	event := new(ExchangeV2NewDomainSeparator)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewDomainSeparator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewGasLimitETHTransferIterator is returned from FilterNewGasLimitETHTransfer and is used to iterate over the raw logs and unpacked data for NewGasLimitETHTransfer events raised by the ExchangeV2 contract.
type ExchangeV2NewGasLimitETHTransferIterator struct {
	Event *ExchangeV2NewGasLimitETHTransfer // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewGasLimitETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewGasLimitETHTransfer)
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
		it.Event = new(ExchangeV2NewGasLimitETHTransfer)
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
func (it *ExchangeV2NewGasLimitETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewGasLimitETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewGasLimitETHTransfer represents a NewGasLimitETHTransfer event raised by the ExchangeV2 contract.
type ExchangeV2NewGasLimitETHTransfer struct {
	GasLimitETHTransfer *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewGasLimitETHTransfer is a free log retrieval operation binding the contract event 0xfc3a20d07f3d5bcc0b01a52011f630765611323fa9afa69f63ba2aa19f7364b6.
//
// Solidity: event NewGasLimitETHTransfer(uint256 gasLimitETHTransfer)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewGasLimitETHTransfer(opts *bind.FilterOpts) (*ExchangeV2NewGasLimitETHTransferIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewGasLimitETHTransfer")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewGasLimitETHTransferIterator{contract: _ExchangeV2.contract, event: "NewGasLimitETHTransfer", logs: logs, sub: sub}, nil
}

// WatchNewGasLimitETHTransfer is a free log subscription operation binding the contract event 0xfc3a20d07f3d5bcc0b01a52011f630765611323fa9afa69f63ba2aa19f7364b6.
//
// Solidity: event NewGasLimitETHTransfer(uint256 gasLimitETHTransfer)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewGasLimitETHTransfer(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewGasLimitETHTransfer) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewGasLimitETHTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewGasLimitETHTransfer)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewGasLimitETHTransfer", log); err != nil {
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

// ParseNewGasLimitETHTransfer is a log parse operation binding the contract event 0xfc3a20d07f3d5bcc0b01a52011f630765611323fa9afa69f63ba2aa19f7364b6.
//
// Solidity: event NewGasLimitETHTransfer(uint256 gasLimitETHTransfer)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewGasLimitETHTransfer(log types.Log) (*ExchangeV2NewGasLimitETHTransfer, error) {
	event := new(ExchangeV2NewGasLimitETHTransfer)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewGasLimitETHTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewMaxCreatorFeeBpIterator is returned from FilterNewMaxCreatorFeeBp and is used to iterate over the raw logs and unpacked data for NewMaxCreatorFeeBp events raised by the ExchangeV2 contract.
type ExchangeV2NewMaxCreatorFeeBpIterator struct {
	Event *ExchangeV2NewMaxCreatorFeeBp // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewMaxCreatorFeeBpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewMaxCreatorFeeBp)
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
		it.Event = new(ExchangeV2NewMaxCreatorFeeBp)
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
func (it *ExchangeV2NewMaxCreatorFeeBpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewMaxCreatorFeeBpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewMaxCreatorFeeBp represents a NewMaxCreatorFeeBp event raised by the ExchangeV2 contract.
type ExchangeV2NewMaxCreatorFeeBp struct {
	MaxCreatorFeeBp *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewMaxCreatorFeeBp is a free log retrieval operation binding the contract event 0xcaba16bb74e50491b14ebf4755083b43aaf56a765134681af613a2ef8d732f4f.
//
// Solidity: event NewMaxCreatorFeeBp(uint256 maxCreatorFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewMaxCreatorFeeBp(opts *bind.FilterOpts) (*ExchangeV2NewMaxCreatorFeeBpIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewMaxCreatorFeeBp")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewMaxCreatorFeeBpIterator{contract: _ExchangeV2.contract, event: "NewMaxCreatorFeeBp", logs: logs, sub: sub}, nil
}

// WatchNewMaxCreatorFeeBp is a free log subscription operation binding the contract event 0xcaba16bb74e50491b14ebf4755083b43aaf56a765134681af613a2ef8d732f4f.
//
// Solidity: event NewMaxCreatorFeeBp(uint256 maxCreatorFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewMaxCreatorFeeBp(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewMaxCreatorFeeBp) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewMaxCreatorFeeBp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewMaxCreatorFeeBp)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewMaxCreatorFeeBp", log); err != nil {
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

// ParseNewMaxCreatorFeeBp is a log parse operation binding the contract event 0xcaba16bb74e50491b14ebf4755083b43aaf56a765134681af613a2ef8d732f4f.
//
// Solidity: event NewMaxCreatorFeeBp(uint256 maxCreatorFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewMaxCreatorFeeBp(log types.Log) (*ExchangeV2NewMaxCreatorFeeBp, error) {
	event := new(ExchangeV2NewMaxCreatorFeeBp)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewMaxCreatorFeeBp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the ExchangeV2 contract.
type ExchangeV2NewOwnerIterator struct {
	Event *ExchangeV2NewOwner // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewOwner)
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
		it.Event = new(ExchangeV2NewOwner)
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
func (it *ExchangeV2NewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewOwner represents a NewOwner event raised by the ExchangeV2 contract.
type ExchangeV2NewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewOwner(opts *bind.FilterOpts) (*ExchangeV2NewOwnerIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewOwnerIterator{contract: _ExchangeV2.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewOwner) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewOwner)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewOwner(log types.Log) (*ExchangeV2NewOwner, error) {
	event := new(ExchangeV2NewOwner)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the ExchangeV2 contract.
type ExchangeV2NewProtocolFeeRecipientIterator struct {
	Event *ExchangeV2NewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewProtocolFeeRecipient)
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
		it.Event = new(ExchangeV2NewProtocolFeeRecipient)
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
func (it *ExchangeV2NewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the ExchangeV2 contract.
type ExchangeV2NewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts) (*ExchangeV2NewProtocolFeeRecipientIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewProtocolFeeRecipientIterator{contract: _ExchangeV2.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewProtocolFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewProtocolFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewProtocolFeeRecipient)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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

// ParseNewProtocolFeeRecipient is a log parse operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address protocolFeeRecipient)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewProtocolFeeRecipient(log types.Log) (*ExchangeV2NewProtocolFeeRecipient, error) {
	event := new(ExchangeV2NewProtocolFeeRecipient)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2NewStrategyIterator is returned from FilterNewStrategy and is used to iterate over the raw logs and unpacked data for NewStrategy events raised by the ExchangeV2 contract.
type ExchangeV2NewStrategyIterator struct {
	Event *ExchangeV2NewStrategy // Event containing the contract specifics and raw log

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
func (it *ExchangeV2NewStrategyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2NewStrategy)
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
		it.Event = new(ExchangeV2NewStrategy)
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
func (it *ExchangeV2NewStrategyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2NewStrategyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2NewStrategy represents a NewStrategy event raised by the ExchangeV2 contract.
type ExchangeV2NewStrategy struct {
	StrategyId            *big.Int
	StandardProtocolFeeBp uint16
	MinTotalFeeBp         uint16
	MaxProtocolFeeBp      uint16
	Selector              [4]byte
	IsMakerBid            bool
	Implementation        common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterNewStrategy is a free log retrieval operation binding the contract event 0x5290475107686ff8d28cd104943b127d453b23622ac55346373fa25c0c8957a2.
//
// Solidity: event NewStrategy(uint256 strategyId, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2Filterer) FilterNewStrategy(opts *bind.FilterOpts) (*ExchangeV2NewStrategyIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "NewStrategy")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2NewStrategyIterator{contract: _ExchangeV2.contract, event: "NewStrategy", logs: logs, sub: sub}, nil
}

// WatchNewStrategy is a free log subscription operation binding the contract event 0x5290475107686ff8d28cd104943b127d453b23622ac55346373fa25c0c8957a2.
//
// Solidity: event NewStrategy(uint256 strategyId, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2Filterer) WatchNewStrategy(opts *bind.WatchOpts, sink chan<- *ExchangeV2NewStrategy) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "NewStrategy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2NewStrategy)
				if err := _ExchangeV2.contract.UnpackLog(event, "NewStrategy", log); err != nil {
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

// ParseNewStrategy is a log parse operation binding the contract event 0x5290475107686ff8d28cd104943b127d453b23622ac55346373fa25c0c8957a2.
//
// Solidity: event NewStrategy(uint256 strategyId, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp, uint16 maxProtocolFeeBp, bytes4 selector, bool isMakerBid, address implementation)
func (_ExchangeV2 *ExchangeV2Filterer) ParseNewStrategy(log types.Log) (*ExchangeV2NewStrategy, error) {
	event := new(ExchangeV2NewStrategy)
	if err := _ExchangeV2.contract.UnpackLog(event, "NewStrategy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2OrderNoncesCancelledIterator is returned from FilterOrderNoncesCancelled and is used to iterate over the raw logs and unpacked data for OrderNoncesCancelled events raised by the ExchangeV2 contract.
type ExchangeV2OrderNoncesCancelledIterator struct {
	Event *ExchangeV2OrderNoncesCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeV2OrderNoncesCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2OrderNoncesCancelled)
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
		it.Event = new(ExchangeV2OrderNoncesCancelled)
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
func (it *ExchangeV2OrderNoncesCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2OrderNoncesCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2OrderNoncesCancelled represents a OrderNoncesCancelled event raised by the ExchangeV2 contract.
type ExchangeV2OrderNoncesCancelled struct {
	User        common.Address
	OrderNonces []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderNoncesCancelled is a free log retrieval operation binding the contract event 0x0560c6093fba8a508d0e6ea3b4d7260d7afa9b152731f03a2d05dfe39b0ec425.
//
// Solidity: event OrderNoncesCancelled(address user, uint256[] orderNonces)
func (_ExchangeV2 *ExchangeV2Filterer) FilterOrderNoncesCancelled(opts *bind.FilterOpts) (*ExchangeV2OrderNoncesCancelledIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "OrderNoncesCancelled")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2OrderNoncesCancelledIterator{contract: _ExchangeV2.contract, event: "OrderNoncesCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderNoncesCancelled is a free log subscription operation binding the contract event 0x0560c6093fba8a508d0e6ea3b4d7260d7afa9b152731f03a2d05dfe39b0ec425.
//
// Solidity: event OrderNoncesCancelled(address user, uint256[] orderNonces)
func (_ExchangeV2 *ExchangeV2Filterer) WatchOrderNoncesCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeV2OrderNoncesCancelled) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "OrderNoncesCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2OrderNoncesCancelled)
				if err := _ExchangeV2.contract.UnpackLog(event, "OrderNoncesCancelled", log); err != nil {
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

// ParseOrderNoncesCancelled is a log parse operation binding the contract event 0x0560c6093fba8a508d0e6ea3b4d7260d7afa9b152731f03a2d05dfe39b0ec425.
//
// Solidity: event OrderNoncesCancelled(address user, uint256[] orderNonces)
func (_ExchangeV2 *ExchangeV2Filterer) ParseOrderNoncesCancelled(log types.Log) (*ExchangeV2OrderNoncesCancelled, error) {
	event := new(ExchangeV2OrderNoncesCancelled)
	if err := _ExchangeV2.contract.UnpackLog(event, "OrderNoncesCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2StrategyUpdatedIterator is returned from FilterStrategyUpdated and is used to iterate over the raw logs and unpacked data for StrategyUpdated events raised by the ExchangeV2 contract.
type ExchangeV2StrategyUpdatedIterator struct {
	Event *ExchangeV2StrategyUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangeV2StrategyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2StrategyUpdated)
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
		it.Event = new(ExchangeV2StrategyUpdated)
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
func (it *ExchangeV2StrategyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2StrategyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2StrategyUpdated represents a StrategyUpdated event raised by the ExchangeV2 contract.
type ExchangeV2StrategyUpdated struct {
	StrategyId            *big.Int
	IsActive              bool
	StandardProtocolFeeBp uint16
	MinTotalFeeBp         uint16
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterStrategyUpdated is a free log retrieval operation binding the contract event 0x3b25bf77fa688236b850bf89c87e353098688237aa18dc42593aff0f6387aea9.
//
// Solidity: event StrategyUpdated(uint256 strategyId, bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) FilterStrategyUpdated(opts *bind.FilterOpts) (*ExchangeV2StrategyUpdatedIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "StrategyUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2StrategyUpdatedIterator{contract: _ExchangeV2.contract, event: "StrategyUpdated", logs: logs, sub: sub}, nil
}

// WatchStrategyUpdated is a free log subscription operation binding the contract event 0x3b25bf77fa688236b850bf89c87e353098688237aa18dc42593aff0f6387aea9.
//
// Solidity: event StrategyUpdated(uint256 strategyId, bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) WatchStrategyUpdated(opts *bind.WatchOpts, sink chan<- *ExchangeV2StrategyUpdated) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "StrategyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2StrategyUpdated)
				if err := _ExchangeV2.contract.UnpackLog(event, "StrategyUpdated", log); err != nil {
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

// ParseStrategyUpdated is a log parse operation binding the contract event 0x3b25bf77fa688236b850bf89c87e353098688237aa18dc42593aff0f6387aea9.
//
// Solidity: event StrategyUpdated(uint256 strategyId, bool isActive, uint16 standardProtocolFeeBp, uint16 minTotalFeeBp)
func (_ExchangeV2 *ExchangeV2Filterer) ParseStrategyUpdated(log types.Log) (*ExchangeV2StrategyUpdated, error) {
	event := new(ExchangeV2StrategyUpdated)
	if err := _ExchangeV2.contract.UnpackLog(event, "StrategyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2SubsetNoncesCancelledIterator is returned from FilterSubsetNoncesCancelled and is used to iterate over the raw logs and unpacked data for SubsetNoncesCancelled events raised by the ExchangeV2 contract.
type ExchangeV2SubsetNoncesCancelledIterator struct {
	Event *ExchangeV2SubsetNoncesCancelled // Event containing the contract specifics and raw log

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
func (it *ExchangeV2SubsetNoncesCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2SubsetNoncesCancelled)
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
		it.Event = new(ExchangeV2SubsetNoncesCancelled)
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
func (it *ExchangeV2SubsetNoncesCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2SubsetNoncesCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2SubsetNoncesCancelled represents a SubsetNoncesCancelled event raised by the ExchangeV2 contract.
type ExchangeV2SubsetNoncesCancelled struct {
	User         common.Address
	SubsetNonces []*big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubsetNoncesCancelled is a free log retrieval operation binding the contract event 0xe8036d6fb143373f3ff63e551373f5fffe4267f6809bf6d3934014a18a9b38f6.
//
// Solidity: event SubsetNoncesCancelled(address user, uint256[] subsetNonces)
func (_ExchangeV2 *ExchangeV2Filterer) FilterSubsetNoncesCancelled(opts *bind.FilterOpts) (*ExchangeV2SubsetNoncesCancelledIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "SubsetNoncesCancelled")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2SubsetNoncesCancelledIterator{contract: _ExchangeV2.contract, event: "SubsetNoncesCancelled", logs: logs, sub: sub}, nil
}

// WatchSubsetNoncesCancelled is a free log subscription operation binding the contract event 0xe8036d6fb143373f3ff63e551373f5fffe4267f6809bf6d3934014a18a9b38f6.
//
// Solidity: event SubsetNoncesCancelled(address user, uint256[] subsetNonces)
func (_ExchangeV2 *ExchangeV2Filterer) WatchSubsetNoncesCancelled(opts *bind.WatchOpts, sink chan<- *ExchangeV2SubsetNoncesCancelled) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "SubsetNoncesCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2SubsetNoncesCancelled)
				if err := _ExchangeV2.contract.UnpackLog(event, "SubsetNoncesCancelled", log); err != nil {
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

// ParseSubsetNoncesCancelled is a log parse operation binding the contract event 0xe8036d6fb143373f3ff63e551373f5fffe4267f6809bf6d3934014a18a9b38f6.
//
// Solidity: event SubsetNoncesCancelled(address user, uint256[] subsetNonces)
func (_ExchangeV2 *ExchangeV2Filterer) ParseSubsetNoncesCancelled(log types.Log) (*ExchangeV2SubsetNoncesCancelled, error) {
	event := new(ExchangeV2SubsetNoncesCancelled)
	if err := _ExchangeV2.contract.UnpackLog(event, "SubsetNoncesCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2TakerAskIterator is returned from FilterTakerAsk and is used to iterate over the raw logs and unpacked data for TakerAsk events raised by the ExchangeV2 contract.
type ExchangeV2TakerAskIterator struct {
	Event *ExchangeV2TakerAsk // Event containing the contract specifics and raw log

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
func (it *ExchangeV2TakerAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2TakerAsk)
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
		it.Event = new(ExchangeV2TakerAsk)
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
func (it *ExchangeV2TakerAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2TakerAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2TakerAsk represents a TakerAsk event raised by the ExchangeV2 contract.
type ExchangeV2TakerAsk struct {
	NonceInvalidationParameters ILooksRareProtocolNonceInvalidationParameters
	AskUser                     common.Address
	BidUser                     common.Address
	StrategyId                  *big.Int
	Currency                    common.Address
	Collection                  common.Address
	ItemIds                     []*big.Int
	Amounts                     []*big.Int
	FeeRecipients               [2]common.Address
	FeeAmounts                  [3]*big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTakerAsk is a free log retrieval operation binding the contract event 0x9aaa45d6db2ef74ead0751ea9113263d1dec1b50cea05f0ca2002cb8063564a4.
//
// Solidity: event TakerAsk((bytes32,uint256,bool) nonceInvalidationParameters, address askUser, address bidUser, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) FilterTakerAsk(opts *bind.FilterOpts) (*ExchangeV2TakerAskIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "TakerAsk")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2TakerAskIterator{contract: _ExchangeV2.contract, event: "TakerAsk", logs: logs, sub: sub}, nil
}

// WatchTakerAsk is a free log subscription operation binding the contract event 0x9aaa45d6db2ef74ead0751ea9113263d1dec1b50cea05f0ca2002cb8063564a4.
//
// Solidity: event TakerAsk((bytes32,uint256,bool) nonceInvalidationParameters, address askUser, address bidUser, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) WatchTakerAsk(opts *bind.WatchOpts, sink chan<- *ExchangeV2TakerAsk) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "TakerAsk")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2TakerAsk)
				if err := _ExchangeV2.contract.UnpackLog(event, "TakerAsk", log); err != nil {
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

// ParseTakerAsk is a log parse operation binding the contract event 0x9aaa45d6db2ef74ead0751ea9113263d1dec1b50cea05f0ca2002cb8063564a4.
//
// Solidity: event TakerAsk((bytes32,uint256,bool) nonceInvalidationParameters, address askUser, address bidUser, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) ParseTakerAsk(log types.Log) (*ExchangeV2TakerAsk, error) {
	event := new(ExchangeV2TakerAsk)
	if err := _ExchangeV2.contract.UnpackLog(event, "TakerAsk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeV2TakerBidIterator is returned from FilterTakerBid and is used to iterate over the raw logs and unpacked data for TakerBid events raised by the ExchangeV2 contract.
type ExchangeV2TakerBidIterator struct {
	Event *ExchangeV2TakerBid // Event containing the contract specifics and raw log

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
func (it *ExchangeV2TakerBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeV2TakerBid)
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
		it.Event = new(ExchangeV2TakerBid)
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
func (it *ExchangeV2TakerBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeV2TakerBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeV2TakerBid represents a TakerBid event raised by the ExchangeV2 contract.
type ExchangeV2TakerBid struct {
	NonceInvalidationParameters ILooksRareProtocolNonceInvalidationParameters
	BidUser                     common.Address
	BidRecipient                common.Address
	StrategyId                  *big.Int
	Currency                    common.Address
	Collection                  common.Address
	ItemIds                     []*big.Int
	Amounts                     []*big.Int
	FeeRecipients               [2]common.Address
	FeeAmounts                  [3]*big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterTakerBid is a free log retrieval operation binding the contract event 0x3ee3de4684413690dee6fff1a0a4f92916a1b97d1c5a83cdf24671844306b2e3.
//
// Solidity: event TakerBid((bytes32,uint256,bool) nonceInvalidationParameters, address bidUser, address bidRecipient, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) FilterTakerBid(opts *bind.FilterOpts) (*ExchangeV2TakerBidIterator, error) {

	logs, sub, err := _ExchangeV2.contract.FilterLogs(opts, "TakerBid")
	if err != nil {
		return nil, err
	}
	return &ExchangeV2TakerBidIterator{contract: _ExchangeV2.contract, event: "TakerBid", logs: logs, sub: sub}, nil
}

// WatchTakerBid is a free log subscription operation binding the contract event 0x3ee3de4684413690dee6fff1a0a4f92916a1b97d1c5a83cdf24671844306b2e3.
//
// Solidity: event TakerBid((bytes32,uint256,bool) nonceInvalidationParameters, address bidUser, address bidRecipient, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) WatchTakerBid(opts *bind.WatchOpts, sink chan<- *ExchangeV2TakerBid) (event.Subscription, error) {

	logs, sub, err := _ExchangeV2.contract.WatchLogs(opts, "TakerBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeV2TakerBid)
				if err := _ExchangeV2.contract.UnpackLog(event, "TakerBid", log); err != nil {
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

// ParseTakerBid is a log parse operation binding the contract event 0x3ee3de4684413690dee6fff1a0a4f92916a1b97d1c5a83cdf24671844306b2e3.
//
// Solidity: event TakerBid((bytes32,uint256,bool) nonceInvalidationParameters, address bidUser, address bidRecipient, uint256 strategyId, address currency, address collection, uint256[] itemIds, uint256[] amounts, address[2] feeRecipients, uint256[3] feeAmounts)
func (_ExchangeV2 *ExchangeV2Filterer) ParseTakerBid(log types.Log) (*ExchangeV2TakerBid, error) {
	event := new(ExchangeV2TakerBid)
	if err := _ExchangeV2.contract.UnpackLog(event, "TakerBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
