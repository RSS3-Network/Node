// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lens

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

// TypesCommentParams is an auto generated low-level Go binding around an user-defined struct.
type TypesCommentParams struct {
	ProfileId               *big.Int
	ContentURI              string
	PointedProfileId        *big.Int
	PointedPubId            *big.Int
	ReferrerProfileIds      []*big.Int
	ReferrerPubIds          []*big.Int
	ReferenceModuleData     []byte
	ActionModules           []common.Address
	ActionModulesInitDatas  [][]byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// TypesMirrorParams is an auto generated low-level Go binding around an user-defined struct.
type TypesMirrorParams struct {
	ProfileId           *big.Int
	MetadataURI         string
	PointedProfileId    *big.Int
	PointedPubId        *big.Int
	ReferrerProfileIds  []*big.Int
	ReferrerPubIds      []*big.Int
	ReferenceModuleData []byte
}

// TypesPostParams is an auto generated low-level Go binding around an user-defined struct.
type TypesPostParams struct {
	ProfileId               *big.Int
	ContentURI              string
	ActionModules           []common.Address
	ActionModulesInitDatas  [][]byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// TypesPublicationActionParams is an auto generated low-level Go binding around an user-defined struct.
type TypesPublicationActionParams struct {
	PublicationActedProfileId *big.Int
	PublicationActedId        *big.Int
	ActorProfileId            *big.Int
	ReferrerProfileIds        []*big.Int
	ReferrerPubIds            []*big.Int
	ActionModuleAddress       common.Address
	ActionModuleData          []byte
}

// TypesQuoteParams is an auto generated low-level Go binding around an user-defined struct.
type TypesQuoteParams struct {
	ProfileId               *big.Int
	ContentURI              string
	PointedProfileId        *big.Int
	PointedPubId            *big.Int
	ReferrerProfileIds      []*big.Int
	ReferrerPubIds          []*big.Int
	ReferenceModuleData     []byte
	ActionModules           []common.Address
	ActionModulesInitDatas  [][]byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// V2EventsMetaData contains all meta data concerning the V2Events contract.
var V2EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationActedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationActedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"actionModuleAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"actionModuleData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTypes.PublicationActionParams\",\"name\":\"publicationActionParams\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"actionModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Acted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"byProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idOfProfileBlocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Blocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTypes.CommentParams\",\"name\":\"commentParams\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitReturnDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleInitReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelegatedExecutorsConfigApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"configNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"delegatedExecutors\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"approvals\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelegatedExecutorsConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"followerProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idOfProfileFollowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"followTokenIdAssigned\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"processFollowModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Followed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LegacyCollectNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTypes.MirrorParams\",\"name\":\"mirrorParams\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MirrorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NonceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTypes.PostParams\",\"name\":\"postParams\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitReturnDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleInitReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PostCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structTypes.QuoteParams\",\"name\":\"quoteParams\",\"type\":\"tuple\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitReturnDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleInitReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"QuoteCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumTypes.ProtocolState\",\"name\":\"prevState\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenGuardianDisablingTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenGuardianStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"prevTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasuryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"byProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idOfProfileUnblocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unblocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unfollowerProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idOfProfileUnfollowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unfollowed\",\"type\":\"event\"}]",
}

// V2EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use V2EventsMetaData.ABI instead.
var V2EventsABI = V2EventsMetaData.ABI

// V2Events is an auto generated Go binding around an Ethereum contract.
type V2Events struct {
	V2EventsCaller     // Read-only binding to the contract
	V2EventsTransactor // Write-only binding to the contract
	V2EventsFilterer   // Log filterer for contract events
}

// V2EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2EventsSession struct {
	Contract     *V2Events         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2EventsCallerSession struct {
	Contract *V2EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// V2EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2EventsTransactorSession struct {
	Contract     *V2EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// V2EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2EventsRaw struct {
	Contract *V2Events // Generic contract binding to access the raw methods on
}

// V2EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2EventsCallerRaw struct {
	Contract *V2EventsCaller // Generic read-only contract binding to access the raw methods on
}

// V2EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2EventsTransactorRaw struct {
	Contract *V2EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2Events creates a new instance of V2Events, bound to a specific deployed contract.
func NewV2Events(address common.Address, backend bind.ContractBackend) (*V2Events, error) {
	contract, err := bindV2Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2Events{V2EventsCaller: V2EventsCaller{contract: contract}, V2EventsTransactor: V2EventsTransactor{contract: contract}, V2EventsFilterer: V2EventsFilterer{contract: contract}}, nil
}

// NewV2EventsCaller creates a new read-only instance of V2Events, bound to a specific deployed contract.
func NewV2EventsCaller(address common.Address, caller bind.ContractCaller) (*V2EventsCaller, error) {
	contract, err := bindV2Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2EventsCaller{contract: contract}, nil
}

// NewV2EventsTransactor creates a new write-only instance of V2Events, bound to a specific deployed contract.
func NewV2EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*V2EventsTransactor, error) {
	contract, err := bindV2Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2EventsTransactor{contract: contract}, nil
}

// NewV2EventsFilterer creates a new log filterer instance of V2Events, bound to a specific deployed contract.
func NewV2EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*V2EventsFilterer, error) {
	contract, err := bindV2Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2EventsFilterer{contract: contract}, nil
}

// bindV2Events binds a generic wrapper to an already deployed contract.
func bindV2Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2Events *V2EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2Events.Contract.V2EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2Events *V2EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Events.Contract.V2EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2Events *V2EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2Events.Contract.V2EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2Events *V2EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2Events *V2EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2Events *V2EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2Events.Contract.contract.Transact(opts, method, params...)
}

// V2EventsActedIterator is returned from FilterActed and is used to iterate over the raw logs and unpacked data for Acted events raised by the V2Events contract.
type V2EventsActedIterator struct {
	Event *V2EventsActed // Event containing the contract specifics and raw log

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
func (it *V2EventsActedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsActed)
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
		it.Event = new(V2EventsActed)
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
func (it *V2EventsActedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsActedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsActed represents a Acted event raised by the V2Events contract.
type V2EventsActed struct {
	PublicationActionParams TypesPublicationActionParams
	ActionModuleReturnData  []byte
	TransactionExecutor     common.Address
	Timestamp               *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterActed is a free log retrieval operation binding the contract event 0x61f8aa74c55cf20b1d5e4f2f6531f66747a0bbbc7696cbb2844738feb8300aad.
//
// Solidity: event Acted((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, bytes actionModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterActed(opts *bind.FilterOpts) (*V2EventsActedIterator, error) {

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "Acted")
	if err != nil {
		return nil, err
	}
	return &V2EventsActedIterator{contract: _V2Events.contract, event: "Acted", logs: logs, sub: sub}, nil
}

// WatchActed is a free log subscription operation binding the contract event 0x61f8aa74c55cf20b1d5e4f2f6531f66747a0bbbc7696cbb2844738feb8300aad.
//
// Solidity: event Acted((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, bytes actionModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchActed(opts *bind.WatchOpts, sink chan<- *V2EventsActed) (event.Subscription, error) {

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "Acted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsActed)
				if err := _V2Events.contract.UnpackLog(event, "Acted", log); err != nil {
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

// ParseActed is a log parse operation binding the contract event 0x61f8aa74c55cf20b1d5e4f2f6531f66747a0bbbc7696cbb2844738feb8300aad.
//
// Solidity: event Acted((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, bytes actionModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseActed(log types.Log) (*V2EventsActed, error) {
	event := new(V2EventsActed)
	if err := _V2Events.contract.UnpackLog(event, "Acted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the V2Events contract.
type V2EventsBaseInitializedIterator struct {
	Event *V2EventsBaseInitialized // Event containing the contract specifics and raw log

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
func (it *V2EventsBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsBaseInitialized)
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
		it.Event = new(V2EventsBaseInitialized)
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
func (it *V2EventsBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsBaseInitialized represents a BaseInitialized event raised by the V2Events contract.
type V2EventsBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*V2EventsBaseInitializedIterator, error) {

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &V2EventsBaseInitializedIterator{contract: _V2Events.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *V2EventsBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsBaseInitialized)
				if err := _V2Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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

// ParseBaseInitialized is a log parse operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseBaseInitialized(log types.Log) (*V2EventsBaseInitialized, error) {
	event := new(V2EventsBaseInitialized)
	if err := _V2Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsBlockedIterator is returned from FilterBlocked and is used to iterate over the raw logs and unpacked data for Blocked events raised by the V2Events contract.
type V2EventsBlockedIterator struct {
	Event *V2EventsBlocked // Event containing the contract specifics and raw log

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
func (it *V2EventsBlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsBlocked)
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
		it.Event = new(V2EventsBlocked)
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
func (it *V2EventsBlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsBlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsBlocked represents a Blocked event raised by the V2Events contract.
type V2EventsBlocked struct {
	ByProfileId         *big.Int
	IdOfProfileBlocked  *big.Int
	TransactionExecutor common.Address
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterBlocked is a free log retrieval operation binding the contract event 0xf516352169f43b6951b54bf4ca4dd3e4632ba04a38956d8af42c44cfde4a0ecb.
//
// Solidity: event Blocked(uint256 indexed byProfileId, uint256 idOfProfileBlocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterBlocked(opts *bind.FilterOpts, byProfileId []*big.Int) (*V2EventsBlockedIterator, error) {

	var byProfileIdRule []interface{}
	for _, byProfileIdItem := range byProfileId {
		byProfileIdRule = append(byProfileIdRule, byProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "Blocked", byProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsBlockedIterator{contract: _V2Events.contract, event: "Blocked", logs: logs, sub: sub}, nil
}

// WatchBlocked is a free log subscription operation binding the contract event 0xf516352169f43b6951b54bf4ca4dd3e4632ba04a38956d8af42c44cfde4a0ecb.
//
// Solidity: event Blocked(uint256 indexed byProfileId, uint256 idOfProfileBlocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchBlocked(opts *bind.WatchOpts, sink chan<- *V2EventsBlocked, byProfileId []*big.Int) (event.Subscription, error) {

	var byProfileIdRule []interface{}
	for _, byProfileIdItem := range byProfileId {
		byProfileIdRule = append(byProfileIdRule, byProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "Blocked", byProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsBlocked)
				if err := _V2Events.contract.UnpackLog(event, "Blocked", log); err != nil {
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

// ParseBlocked is a log parse operation binding the contract event 0xf516352169f43b6951b54bf4ca4dd3e4632ba04a38956d8af42c44cfde4a0ecb.
//
// Solidity: event Blocked(uint256 indexed byProfileId, uint256 idOfProfileBlocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseBlocked(log types.Log) (*V2EventsBlocked, error) {
	event := new(V2EventsBlocked)
	if err := _V2Events.contract.UnpackLog(event, "Blocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsCollectNFTTransferredIterator is returned from FilterCollectNFTTransferred and is used to iterate over the raw logs and unpacked data for CollectNFTTransferred events raised by the V2Events contract.
type V2EventsCollectNFTTransferredIterator struct {
	Event *V2EventsCollectNFTTransferred // Event containing the contract specifics and raw log

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
func (it *V2EventsCollectNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsCollectNFTTransferred)
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
		it.Event = new(V2EventsCollectNFTTransferred)
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
func (it *V2EventsCollectNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsCollectNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsCollectNFTTransferred represents a CollectNFTTransferred event raised by the V2Events contract.
type V2EventsCollectNFTTransferred struct {
	ProfileId    *big.Int
	PubId        *big.Int
	CollectNFTId *big.Int
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTTransferred is a free log retrieval operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterCollectNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (*V2EventsCollectNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsCollectNFTTransferredIterator{contract: _V2Events.contract, event: "CollectNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchCollectNFTTransferred is a free log subscription operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchCollectNFTTransferred(opts *bind.WatchOpts, sink chan<- *V2EventsCollectNFTTransferred, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsCollectNFTTransferred)
				if err := _V2Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
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

// ParseCollectNFTTransferred is a log parse operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseCollectNFTTransferred(log types.Log) (*V2EventsCollectNFTTransferred, error) {
	event := new(V2EventsCollectNFTTransferred)
	if err := _V2Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsCommentCreatedIterator is returned from FilterCommentCreated and is used to iterate over the raw logs and unpacked data for CommentCreated events raised by the V2Events contract.
type V2EventsCommentCreatedIterator struct {
	Event *V2EventsCommentCreated // Event containing the contract specifics and raw log

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
func (it *V2EventsCommentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsCommentCreated)
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
		it.Event = new(V2EventsCommentCreated)
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
func (it *V2EventsCommentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsCommentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsCommentCreated represents a CommentCreated event raised by the V2Events contract.
type V2EventsCommentCreated struct {
	CommentParams                 TypesCommentParams
	PubId                         *big.Int
	ReferenceModuleReturnData     []byte
	ActionModulesInitReturnDatas  [][]byte
	ReferenceModuleInitReturnData []byte
	TransactionExecutor           common.Address
	Timestamp                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterCommentCreated is a free log retrieval operation binding the contract event 0x6730c5edd594025e9d1057522801addabbb26fe8ec0acb70a658002f75684388.
//
// Solidity: event CommentCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterCommentCreated(opts *bind.FilterOpts, pubId []*big.Int) (*V2EventsCommentCreatedIterator, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "CommentCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsCommentCreatedIterator{contract: _V2Events.contract, event: "CommentCreated", logs: logs, sub: sub}, nil
}

// WatchCommentCreated is a free log subscription operation binding the contract event 0x6730c5edd594025e9d1057522801addabbb26fe8ec0acb70a658002f75684388.
//
// Solidity: event CommentCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchCommentCreated(opts *bind.WatchOpts, sink chan<- *V2EventsCommentCreated, pubId []*big.Int) (event.Subscription, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "CommentCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsCommentCreated)
				if err := _V2Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
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

// ParseCommentCreated is a log parse operation binding the contract event 0x6730c5edd594025e9d1057522801addabbb26fe8ec0acb70a658002f75684388.
//
// Solidity: event CommentCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseCommentCreated(log types.Log) (*V2EventsCommentCreated, error) {
	event := new(V2EventsCommentCreated)
	if err := _V2Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsDelegatedExecutorsConfigAppliedIterator is returned from FilterDelegatedExecutorsConfigApplied and is used to iterate over the raw logs and unpacked data for DelegatedExecutorsConfigApplied events raised by the V2Events contract.
type V2EventsDelegatedExecutorsConfigAppliedIterator struct {
	Event *V2EventsDelegatedExecutorsConfigApplied // Event containing the contract specifics and raw log

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
func (it *V2EventsDelegatedExecutorsConfigAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsDelegatedExecutorsConfigApplied)
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
		it.Event = new(V2EventsDelegatedExecutorsConfigApplied)
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
func (it *V2EventsDelegatedExecutorsConfigAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsDelegatedExecutorsConfigAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsDelegatedExecutorsConfigApplied represents a DelegatedExecutorsConfigApplied event raised by the V2Events contract.
type V2EventsDelegatedExecutorsConfigApplied struct {
	DelegatorProfileId *big.Int
	ConfigNumber       *big.Int
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatedExecutorsConfigApplied is a free log retrieval operation binding the contract event 0x47f050f9e2779af58ee28eb99af69c3168632fbf7a45f97317137e2098e538f3.
//
// Solidity: event DelegatedExecutorsConfigApplied(uint256 indexed delegatorProfileId, uint256 indexed configNumber, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterDelegatedExecutorsConfigApplied(opts *bind.FilterOpts, delegatorProfileId []*big.Int, configNumber []*big.Int) (*V2EventsDelegatedExecutorsConfigAppliedIterator, error) {

	var delegatorProfileIdRule []interface{}
	for _, delegatorProfileIdItem := range delegatorProfileId {
		delegatorProfileIdRule = append(delegatorProfileIdRule, delegatorProfileIdItem)
	}
	var configNumberRule []interface{}
	for _, configNumberItem := range configNumber {
		configNumberRule = append(configNumberRule, configNumberItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "DelegatedExecutorsConfigApplied", delegatorProfileIdRule, configNumberRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsDelegatedExecutorsConfigAppliedIterator{contract: _V2Events.contract, event: "DelegatedExecutorsConfigApplied", logs: logs, sub: sub}, nil
}

// WatchDelegatedExecutorsConfigApplied is a free log subscription operation binding the contract event 0x47f050f9e2779af58ee28eb99af69c3168632fbf7a45f97317137e2098e538f3.
//
// Solidity: event DelegatedExecutorsConfigApplied(uint256 indexed delegatorProfileId, uint256 indexed configNumber, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchDelegatedExecutorsConfigApplied(opts *bind.WatchOpts, sink chan<- *V2EventsDelegatedExecutorsConfigApplied, delegatorProfileId []*big.Int, configNumber []*big.Int) (event.Subscription, error) {

	var delegatorProfileIdRule []interface{}
	for _, delegatorProfileIdItem := range delegatorProfileId {
		delegatorProfileIdRule = append(delegatorProfileIdRule, delegatorProfileIdItem)
	}
	var configNumberRule []interface{}
	for _, configNumberItem := range configNumber {
		configNumberRule = append(configNumberRule, configNumberItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "DelegatedExecutorsConfigApplied", delegatorProfileIdRule, configNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsDelegatedExecutorsConfigApplied)
				if err := _V2Events.contract.UnpackLog(event, "DelegatedExecutorsConfigApplied", log); err != nil {
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

// ParseDelegatedExecutorsConfigApplied is a log parse operation binding the contract event 0x47f050f9e2779af58ee28eb99af69c3168632fbf7a45f97317137e2098e538f3.
//
// Solidity: event DelegatedExecutorsConfigApplied(uint256 indexed delegatorProfileId, uint256 indexed configNumber, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseDelegatedExecutorsConfigApplied(log types.Log) (*V2EventsDelegatedExecutorsConfigApplied, error) {
	event := new(V2EventsDelegatedExecutorsConfigApplied)
	if err := _V2Events.contract.UnpackLog(event, "DelegatedExecutorsConfigApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsDelegatedExecutorsConfigChangedIterator is returned from FilterDelegatedExecutorsConfigChanged and is used to iterate over the raw logs and unpacked data for DelegatedExecutorsConfigChanged events raised by the V2Events contract.
type V2EventsDelegatedExecutorsConfigChangedIterator struct {
	Event *V2EventsDelegatedExecutorsConfigChanged // Event containing the contract specifics and raw log

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
func (it *V2EventsDelegatedExecutorsConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsDelegatedExecutorsConfigChanged)
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
		it.Event = new(V2EventsDelegatedExecutorsConfigChanged)
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
func (it *V2EventsDelegatedExecutorsConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsDelegatedExecutorsConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsDelegatedExecutorsConfigChanged represents a DelegatedExecutorsConfigChanged event raised by the V2Events contract.
type V2EventsDelegatedExecutorsConfigChanged struct {
	DelegatorProfileId *big.Int
	ConfigNumber       *big.Int
	DelegatedExecutors []common.Address
	Approvals          []bool
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDelegatedExecutorsConfigChanged is a free log retrieval operation binding the contract event 0xfd73bf07ef75af85f5dec85a987a10fc4a54c330dd4d13ae0f619cf59d96e506.
//
// Solidity: event DelegatedExecutorsConfigChanged(uint256 indexed delegatorProfileId, uint256 indexed configNumber, address[] delegatedExecutors, bool[] approvals, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterDelegatedExecutorsConfigChanged(opts *bind.FilterOpts, delegatorProfileId []*big.Int, configNumber []*big.Int) (*V2EventsDelegatedExecutorsConfigChangedIterator, error) {

	var delegatorProfileIdRule []interface{}
	for _, delegatorProfileIdItem := range delegatorProfileId {
		delegatorProfileIdRule = append(delegatorProfileIdRule, delegatorProfileIdItem)
	}
	var configNumberRule []interface{}
	for _, configNumberItem := range configNumber {
		configNumberRule = append(configNumberRule, configNumberItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "DelegatedExecutorsConfigChanged", delegatorProfileIdRule, configNumberRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsDelegatedExecutorsConfigChangedIterator{contract: _V2Events.contract, event: "DelegatedExecutorsConfigChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatedExecutorsConfigChanged is a free log subscription operation binding the contract event 0xfd73bf07ef75af85f5dec85a987a10fc4a54c330dd4d13ae0f619cf59d96e506.
//
// Solidity: event DelegatedExecutorsConfigChanged(uint256 indexed delegatorProfileId, uint256 indexed configNumber, address[] delegatedExecutors, bool[] approvals, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchDelegatedExecutorsConfigChanged(opts *bind.WatchOpts, sink chan<- *V2EventsDelegatedExecutorsConfigChanged, delegatorProfileId []*big.Int, configNumber []*big.Int) (event.Subscription, error) {

	var delegatorProfileIdRule []interface{}
	for _, delegatorProfileIdItem := range delegatorProfileId {
		delegatorProfileIdRule = append(delegatorProfileIdRule, delegatorProfileIdItem)
	}
	var configNumberRule []interface{}
	for _, configNumberItem := range configNumber {
		configNumberRule = append(configNumberRule, configNumberItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "DelegatedExecutorsConfigChanged", delegatorProfileIdRule, configNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsDelegatedExecutorsConfigChanged)
				if err := _V2Events.contract.UnpackLog(event, "DelegatedExecutorsConfigChanged", log); err != nil {
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

// ParseDelegatedExecutorsConfigChanged is a log parse operation binding the contract event 0xfd73bf07ef75af85f5dec85a987a10fc4a54c330dd4d13ae0f619cf59d96e506.
//
// Solidity: event DelegatedExecutorsConfigChanged(uint256 indexed delegatorProfileId, uint256 indexed configNumber, address[] delegatedExecutors, bool[] approvals, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseDelegatedExecutorsConfigChanged(log types.Log) (*V2EventsDelegatedExecutorsConfigChanged, error) {
	event := new(V2EventsDelegatedExecutorsConfigChanged)
	if err := _V2Events.contract.UnpackLog(event, "DelegatedExecutorsConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsEmergencyAdminSetIterator is returned from FilterEmergencyAdminSet and is used to iterate over the raw logs and unpacked data for EmergencyAdminSet events raised by the V2Events contract.
type V2EventsEmergencyAdminSetIterator struct {
	Event *V2EventsEmergencyAdminSet // Event containing the contract specifics and raw log

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
func (it *V2EventsEmergencyAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsEmergencyAdminSet)
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
		it.Event = new(V2EventsEmergencyAdminSet)
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
func (it *V2EventsEmergencyAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsEmergencyAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsEmergencyAdminSet represents a EmergencyAdminSet event raised by the V2Events contract.
type V2EventsEmergencyAdminSet struct {
	Caller            common.Address
	OldEmergencyAdmin common.Address
	NewEmergencyAdmin common.Address
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminSet is a free log retrieval operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterEmergencyAdminSet(opts *bind.FilterOpts, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (*V2EventsEmergencyAdminSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsEmergencyAdminSetIterator{contract: _V2Events.contract, event: "EmergencyAdminSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminSet is a free log subscription operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchEmergencyAdminSet(opts *bind.WatchOpts, sink chan<- *V2EventsEmergencyAdminSet, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsEmergencyAdminSet)
				if err := _V2Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
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

// ParseEmergencyAdminSet is a log parse operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseEmergencyAdminSet(log types.Log) (*V2EventsEmergencyAdminSet, error) {
	event := new(V2EventsEmergencyAdminSet)
	if err := _V2Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsFollowModuleSetIterator is returned from FilterFollowModuleSet and is used to iterate over the raw logs and unpacked data for FollowModuleSet events raised by the V2Events contract.
type V2EventsFollowModuleSetIterator struct {
	Event *V2EventsFollowModuleSet // Event containing the contract specifics and raw log

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
func (it *V2EventsFollowModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsFollowModuleSet)
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
		it.Event = new(V2EventsFollowModuleSet)
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
func (it *V2EventsFollowModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsFollowModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsFollowModuleSet represents a FollowModuleSet event raised by the V2Events contract.
type V2EventsFollowModuleSet struct {
	ProfileId              *big.Int
	FollowModule           common.Address
	FollowModuleInitData   []byte
	FollowModuleReturnData []byte
	TransactionExecutor    common.Address
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleSet is a free log retrieval operation binding the contract event 0x160458ccae41f9f8e4bc14a37e1fc3d900f768c5c12991c5140490aae8292600.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleInitData, bytes followModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterFollowModuleSet(opts *bind.FilterOpts, profileId []*big.Int) (*V2EventsFollowModuleSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsFollowModuleSetIterator{contract: _V2Events.contract, event: "FollowModuleSet", logs: logs, sub: sub}, nil
}

// WatchFollowModuleSet is a free log subscription operation binding the contract event 0x160458ccae41f9f8e4bc14a37e1fc3d900f768c5c12991c5140490aae8292600.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleInitData, bytes followModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchFollowModuleSet(opts *bind.WatchOpts, sink chan<- *V2EventsFollowModuleSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsFollowModuleSet)
				if err := _V2Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
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

// ParseFollowModuleSet is a log parse operation binding the contract event 0x160458ccae41f9f8e4bc14a37e1fc3d900f768c5c12991c5140490aae8292600.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleInitData, bytes followModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseFollowModuleSet(log types.Log) (*V2EventsFollowModuleSet, error) {
	event := new(V2EventsFollowModuleSet)
	if err := _V2Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsFollowNFTDeployedIterator is returned from FilterFollowNFTDeployed and is used to iterate over the raw logs and unpacked data for FollowNFTDeployed events raised by the V2Events contract.
type V2EventsFollowNFTDeployedIterator struct {
	Event *V2EventsFollowNFTDeployed // Event containing the contract specifics and raw log

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
func (it *V2EventsFollowNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsFollowNFTDeployed)
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
		it.Event = new(V2EventsFollowNFTDeployed)
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
func (it *V2EventsFollowNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsFollowNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsFollowNFTDeployed represents a FollowNFTDeployed event raised by the V2Events contract.
type V2EventsFollowNFTDeployed struct {
	ProfileId *big.Int
	FollowNFT common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDeployed is a free log retrieval operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterFollowNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, followNFT []common.Address) (*V2EventsFollowNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsFollowNFTDeployedIterator{contract: _V2Events.contract, event: "FollowNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDeployed is a free log subscription operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchFollowNFTDeployed(opts *bind.WatchOpts, sink chan<- *V2EventsFollowNFTDeployed, profileId []*big.Int, followNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsFollowNFTDeployed)
				if err := _V2Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
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

// ParseFollowNFTDeployed is a log parse operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseFollowNFTDeployed(log types.Log) (*V2EventsFollowNFTDeployed, error) {
	event := new(V2EventsFollowNFTDeployed)
	if err := _V2Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsFollowedIterator is returned from FilterFollowed and is used to iterate over the raw logs and unpacked data for Followed events raised by the V2Events contract.
type V2EventsFollowedIterator struct {
	Event *V2EventsFollowed // Event containing the contract specifics and raw log

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
func (it *V2EventsFollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsFollowed)
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
		it.Event = new(V2EventsFollowed)
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
func (it *V2EventsFollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsFollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsFollowed represents a Followed event raised by the V2Events contract.
type V2EventsFollowed struct {
	FollowerProfileId             *big.Int
	IdOfProfileFollowed           *big.Int
	FollowTokenIdAssigned         *big.Int
	FollowModuleData              []byte
	ProcessFollowModuleReturnData []byte
	TransactionExecutor           common.Address
	Timestamp                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterFollowed is a free log retrieval operation binding the contract event 0x817d2c71a3ec35dc50f2e4b0d890943c89f2a7ab9d96eff233eda4932b506d0b.
//
// Solidity: event Followed(uint256 indexed followerProfileId, uint256 idOfProfileFollowed, uint256 followTokenIdAssigned, bytes followModuleData, bytes processFollowModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterFollowed(opts *bind.FilterOpts, followerProfileId []*big.Int) (*V2EventsFollowedIterator, error) {

	var followerProfileIdRule []interface{}
	for _, followerProfileIdItem := range followerProfileId {
		followerProfileIdRule = append(followerProfileIdRule, followerProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "Followed", followerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsFollowedIterator{contract: _V2Events.contract, event: "Followed", logs: logs, sub: sub}, nil
}

// WatchFollowed is a free log subscription operation binding the contract event 0x817d2c71a3ec35dc50f2e4b0d890943c89f2a7ab9d96eff233eda4932b506d0b.
//
// Solidity: event Followed(uint256 indexed followerProfileId, uint256 idOfProfileFollowed, uint256 followTokenIdAssigned, bytes followModuleData, bytes processFollowModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchFollowed(opts *bind.WatchOpts, sink chan<- *V2EventsFollowed, followerProfileId []*big.Int) (event.Subscription, error) {

	var followerProfileIdRule []interface{}
	for _, followerProfileIdItem := range followerProfileId {
		followerProfileIdRule = append(followerProfileIdRule, followerProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "Followed", followerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsFollowed)
				if err := _V2Events.contract.UnpackLog(event, "Followed", log); err != nil {
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

// ParseFollowed is a log parse operation binding the contract event 0x817d2c71a3ec35dc50f2e4b0d890943c89f2a7ab9d96eff233eda4932b506d0b.
//
// Solidity: event Followed(uint256 indexed followerProfileId, uint256 idOfProfileFollowed, uint256 followTokenIdAssigned, bytes followModuleData, bytes processFollowModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseFollowed(log types.Log) (*V2EventsFollowed, error) {
	event := new(V2EventsFollowed)
	if err := _V2Events.contract.UnpackLog(event, "Followed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsGovernanceSetIterator is returned from FilterGovernanceSet and is used to iterate over the raw logs and unpacked data for GovernanceSet events raised by the V2Events contract.
type V2EventsGovernanceSetIterator struct {
	Event *V2EventsGovernanceSet // Event containing the contract specifics and raw log

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
func (it *V2EventsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsGovernanceSet)
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
		it.Event = new(V2EventsGovernanceSet)
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
func (it *V2EventsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsGovernanceSet represents a GovernanceSet event raised by the V2Events contract.
type V2EventsGovernanceSet struct {
	Caller         common.Address
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSet is a free log retrieval operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterGovernanceSet(opts *bind.FilterOpts, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (*V2EventsGovernanceSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsGovernanceSetIterator{contract: _V2Events.contract, event: "GovernanceSet", logs: logs, sub: sub}, nil
}

// WatchGovernanceSet is a free log subscription operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchGovernanceSet(opts *bind.WatchOpts, sink chan<- *V2EventsGovernanceSet, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsGovernanceSet)
				if err := _V2Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
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

// ParseGovernanceSet is a log parse operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseGovernanceSet(log types.Log) (*V2EventsGovernanceSet, error) {
	event := new(V2EventsGovernanceSet)
	if err := _V2Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsLegacyCollectNFTDeployedIterator is returned from FilterLegacyCollectNFTDeployed and is used to iterate over the raw logs and unpacked data for LegacyCollectNFTDeployed events raised by the V2Events contract.
type V2EventsLegacyCollectNFTDeployedIterator struct {
	Event *V2EventsLegacyCollectNFTDeployed // Event containing the contract specifics and raw log

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
func (it *V2EventsLegacyCollectNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsLegacyCollectNFTDeployed)
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
		it.Event = new(V2EventsLegacyCollectNFTDeployed)
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
func (it *V2EventsLegacyCollectNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsLegacyCollectNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsLegacyCollectNFTDeployed represents a LegacyCollectNFTDeployed event raised by the V2Events contract.
type V2EventsLegacyCollectNFTDeployed struct {
	ProfileId  *big.Int
	PubId      *big.Int
	CollectNFT common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLegacyCollectNFTDeployed is a free log retrieval operation binding the contract event 0xe9493626ff4ace96801910d37796801b339b1d5cf29e49bdc3cb870ec5ad0770.
//
// Solidity: event LegacyCollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterLegacyCollectNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (*V2EventsLegacyCollectNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "LegacyCollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsLegacyCollectNFTDeployedIterator{contract: _V2Events.contract, event: "LegacyCollectNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchLegacyCollectNFTDeployed is a free log subscription operation binding the contract event 0xe9493626ff4ace96801910d37796801b339b1d5cf29e49bdc3cb870ec5ad0770.
//
// Solidity: event LegacyCollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchLegacyCollectNFTDeployed(opts *bind.WatchOpts, sink chan<- *V2EventsLegacyCollectNFTDeployed, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "LegacyCollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsLegacyCollectNFTDeployed)
				if err := _V2Events.contract.UnpackLog(event, "LegacyCollectNFTDeployed", log); err != nil {
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

// ParseLegacyCollectNFTDeployed is a log parse operation binding the contract event 0xe9493626ff4ace96801910d37796801b339b1d5cf29e49bdc3cb870ec5ad0770.
//
// Solidity: event LegacyCollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseLegacyCollectNFTDeployed(log types.Log) (*V2EventsLegacyCollectNFTDeployed, error) {
	event := new(V2EventsLegacyCollectNFTDeployed)
	if err := _V2Events.contract.UnpackLog(event, "LegacyCollectNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsMirrorCreatedIterator is returned from FilterMirrorCreated and is used to iterate over the raw logs and unpacked data for MirrorCreated events raised by the V2Events contract.
type V2EventsMirrorCreatedIterator struct {
	Event *V2EventsMirrorCreated // Event containing the contract specifics and raw log

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
func (it *V2EventsMirrorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsMirrorCreated)
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
		it.Event = new(V2EventsMirrorCreated)
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
func (it *V2EventsMirrorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsMirrorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsMirrorCreated represents a MirrorCreated event raised by the V2Events contract.
type V2EventsMirrorCreated struct {
	MirrorParams              TypesMirrorParams
	PubId                     *big.Int
	ReferenceModuleReturnData []byte
	TransactionExecutor       common.Address
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMirrorCreated is a free log retrieval operation binding the contract event 0x19822529a03d77bbe525763dd7064f5c182a5ede1bdd88e73a07221d3f3feb6d.
//
// Solidity: event MirrorCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, uint256 indexed pubId, bytes referenceModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterMirrorCreated(opts *bind.FilterOpts, pubId []*big.Int) (*V2EventsMirrorCreatedIterator, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "MirrorCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsMirrorCreatedIterator{contract: _V2Events.contract, event: "MirrorCreated", logs: logs, sub: sub}, nil
}

// WatchMirrorCreated is a free log subscription operation binding the contract event 0x19822529a03d77bbe525763dd7064f5c182a5ede1bdd88e73a07221d3f3feb6d.
//
// Solidity: event MirrorCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, uint256 indexed pubId, bytes referenceModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchMirrorCreated(opts *bind.WatchOpts, sink chan<- *V2EventsMirrorCreated, pubId []*big.Int) (event.Subscription, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "MirrorCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsMirrorCreated)
				if err := _V2Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
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

// ParseMirrorCreated is a log parse operation binding the contract event 0x19822529a03d77bbe525763dd7064f5c182a5ede1bdd88e73a07221d3f3feb6d.
//
// Solidity: event MirrorCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, uint256 indexed pubId, bytes referenceModuleReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseMirrorCreated(log types.Log) (*V2EventsMirrorCreated, error) {
	event := new(V2EventsMirrorCreated)
	if err := _V2Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsNonceUpdatedIterator is returned from FilterNonceUpdated and is used to iterate over the raw logs and unpacked data for NonceUpdated events raised by the V2Events contract.
type V2EventsNonceUpdatedIterator struct {
	Event *V2EventsNonceUpdated // Event containing the contract specifics and raw log

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
func (it *V2EventsNonceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsNonceUpdated)
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
		it.Event = new(V2EventsNonceUpdated)
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
func (it *V2EventsNonceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsNonceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsNonceUpdated represents a NonceUpdated event raised by the V2Events contract.
type V2EventsNonceUpdated struct {
	Signer    common.Address
	Nonce     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNonceUpdated is a free log retrieval operation binding the contract event 0xc906270cebe7667882104effe64262a73c422ab9176a111e05ea837b021065fc.
//
// Solidity: event NonceUpdated(address indexed signer, uint256 nonce, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterNonceUpdated(opts *bind.FilterOpts, signer []common.Address) (*V2EventsNonceUpdatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "NonceUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsNonceUpdatedIterator{contract: _V2Events.contract, event: "NonceUpdated", logs: logs, sub: sub}, nil
}

// WatchNonceUpdated is a free log subscription operation binding the contract event 0xc906270cebe7667882104effe64262a73c422ab9176a111e05ea837b021065fc.
//
// Solidity: event NonceUpdated(address indexed signer, uint256 nonce, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchNonceUpdated(opts *bind.WatchOpts, sink chan<- *V2EventsNonceUpdated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "NonceUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsNonceUpdated)
				if err := _V2Events.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
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

// ParseNonceUpdated is a log parse operation binding the contract event 0xc906270cebe7667882104effe64262a73c422ab9176a111e05ea837b021065fc.
//
// Solidity: event NonceUpdated(address indexed signer, uint256 nonce, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseNonceUpdated(log types.Log) (*V2EventsNonceUpdated, error) {
	event := new(V2EventsNonceUpdated)
	if err := _V2Events.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the V2Events contract.
type V2EventsPostCreatedIterator struct {
	Event *V2EventsPostCreated // Event containing the contract specifics and raw log

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
func (it *V2EventsPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsPostCreated)
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
		it.Event = new(V2EventsPostCreated)
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
func (it *V2EventsPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsPostCreated represents a PostCreated event raised by the V2Events contract.
type V2EventsPostCreated struct {
	PostParams                    TypesPostParams
	PubId                         *big.Int
	ActionModulesInitReturnDatas  [][]byte
	ReferenceModuleInitReturnData []byte
	TransactionExecutor           common.Address
	Timestamp                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0xe18912378f90aa372fc9ab7ab5ff7e4744182bdef133ccad56d5a18864456742.
//
// Solidity: event PostCreated((uint256,string,address[],bytes[],address,bytes) postParams, uint256 indexed pubId, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterPostCreated(opts *bind.FilterOpts, pubId []*big.Int) (*V2EventsPostCreatedIterator, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "PostCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsPostCreatedIterator{contract: _V2Events.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0xe18912378f90aa372fc9ab7ab5ff7e4744182bdef133ccad56d5a18864456742.
//
// Solidity: event PostCreated((uint256,string,address[],bytes[],address,bytes) postParams, uint256 indexed pubId, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *V2EventsPostCreated, pubId []*big.Int) (event.Subscription, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "PostCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsPostCreated)
				if err := _V2Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
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

// ParsePostCreated is a log parse operation binding the contract event 0xe18912378f90aa372fc9ab7ab5ff7e4744182bdef133ccad56d5a18864456742.
//
// Solidity: event PostCreated((uint256,string,address[],bytes[],address,bytes) postParams, uint256 indexed pubId, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParsePostCreated(log types.Log) (*V2EventsPostCreated, error) {
	event := new(V2EventsPostCreated)
	if err := _V2Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the V2Events contract.
type V2EventsProfileCreatedIterator struct {
	Event *V2EventsProfileCreated // Event containing the contract specifics and raw log

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
func (it *V2EventsProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsProfileCreated)
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
		it.Event = new(V2EventsProfileCreated)
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
func (it *V2EventsProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsProfileCreated represents a ProfileCreated event raised by the V2Events contract.
type V2EventsProfileCreated struct {
	ProfileId *big.Int
	Creator   common.Address
	To        common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0xf642d82f9bf073e3403d88853e8ee1a91d4fff05e11bcdf593f09ce442c6b247.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterProfileCreated(opts *bind.FilterOpts, profileId []*big.Int, creator []common.Address, to []common.Address) (*V2EventsProfileCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsProfileCreatedIterator{contract: _V2Events.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0xf642d82f9bf073e3403d88853e8ee1a91d4fff05e11bcdf593f09ce442c6b247.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *V2EventsProfileCreated, profileId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsProfileCreated)
				if err := _V2Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0xf642d82f9bf073e3403d88853e8ee1a91d4fff05e11bcdf593f09ce442c6b247.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseProfileCreated(log types.Log) (*V2EventsProfileCreated, error) {
	event := new(V2EventsProfileCreated)
	if err := _V2Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsProfileCreatorWhitelistedIterator is returned from FilterProfileCreatorWhitelisted and is used to iterate over the raw logs and unpacked data for ProfileCreatorWhitelisted events raised by the V2Events contract.
type V2EventsProfileCreatorWhitelistedIterator struct {
	Event *V2EventsProfileCreatorWhitelisted // Event containing the contract specifics and raw log

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
func (it *V2EventsProfileCreatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsProfileCreatorWhitelisted)
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
		it.Event = new(V2EventsProfileCreatorWhitelisted)
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
func (it *V2EventsProfileCreatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsProfileCreatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsProfileCreatorWhitelisted represents a ProfileCreatorWhitelisted event raised by the V2Events contract.
type V2EventsProfileCreatorWhitelisted struct {
	ProfileCreator common.Address
	Whitelisted    bool
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfileCreatorWhitelisted is a free log retrieval operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterProfileCreatorWhitelisted(opts *bind.FilterOpts, profileCreator []common.Address, whitelisted []bool) (*V2EventsProfileCreatorWhitelistedIterator, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsProfileCreatorWhitelistedIterator{contract: _V2Events.contract, event: "ProfileCreatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchProfileCreatorWhitelisted is a free log subscription operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchProfileCreatorWhitelisted(opts *bind.WatchOpts, sink chan<- *V2EventsProfileCreatorWhitelisted, profileCreator []common.Address, whitelisted []bool) (event.Subscription, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsProfileCreatorWhitelisted)
				if err := _V2Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
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

// ParseProfileCreatorWhitelisted is a log parse operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseProfileCreatorWhitelisted(log types.Log) (*V2EventsProfileCreatorWhitelisted, error) {
	event := new(V2EventsProfileCreatorWhitelisted)
	if err := _V2Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsProfileMetadataSetIterator is returned from FilterProfileMetadataSet and is used to iterate over the raw logs and unpacked data for ProfileMetadataSet events raised by the V2Events contract.
type V2EventsProfileMetadataSetIterator struct {
	Event *V2EventsProfileMetadataSet // Event containing the contract specifics and raw log

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
func (it *V2EventsProfileMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsProfileMetadataSet)
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
		it.Event = new(V2EventsProfileMetadataSet)
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
func (it *V2EventsProfileMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsProfileMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsProfileMetadataSet represents a ProfileMetadataSet event raised by the V2Events contract.
type V2EventsProfileMetadataSet struct {
	ProfileId           *big.Int
	Metadata            string
	TransactionExecutor common.Address
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterProfileMetadataSet is a free log retrieval operation binding the contract event 0x1202cc90de8ecfb21f883e5a5257eb6cd0deabf48ef14dbf146c1225b3c50abb.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterProfileMetadataSet(opts *bind.FilterOpts, profileId []*big.Int) (*V2EventsProfileMetadataSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsProfileMetadataSetIterator{contract: _V2Events.contract, event: "ProfileMetadataSet", logs: logs, sub: sub}, nil
}

// WatchProfileMetadataSet is a free log subscription operation binding the contract event 0x1202cc90de8ecfb21f883e5a5257eb6cd0deabf48ef14dbf146c1225b3c50abb.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchProfileMetadataSet(opts *bind.WatchOpts, sink chan<- *V2EventsProfileMetadataSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsProfileMetadataSet)
				if err := _V2Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
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

// ParseProfileMetadataSet is a log parse operation binding the contract event 0x1202cc90de8ecfb21f883e5a5257eb6cd0deabf48ef14dbf146c1225b3c50abb.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseProfileMetadataSet(log types.Log) (*V2EventsProfileMetadataSet, error) {
	event := new(V2EventsProfileMetadataSet)
	if err := _V2Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsQuoteCreatedIterator is returned from FilterQuoteCreated and is used to iterate over the raw logs and unpacked data for QuoteCreated events raised by the V2Events contract.
type V2EventsQuoteCreatedIterator struct {
	Event *V2EventsQuoteCreated // Event containing the contract specifics and raw log

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
func (it *V2EventsQuoteCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsQuoteCreated)
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
		it.Event = new(V2EventsQuoteCreated)
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
func (it *V2EventsQuoteCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsQuoteCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsQuoteCreated represents a QuoteCreated event raised by the V2Events contract.
type V2EventsQuoteCreated struct {
	QuoteParams                   TypesQuoteParams
	PubId                         *big.Int
	ReferenceModuleReturnData     []byte
	ActionModulesInitReturnDatas  [][]byte
	ReferenceModuleInitReturnData []byte
	TransactionExecutor           common.Address
	Timestamp                     *big.Int
	Raw                           types.Log // Blockchain specific contextual infos
}

// FilterQuoteCreated is a free log retrieval operation binding the contract event 0x90991d5410a24294cd806880f6e27b11ebe48fbb0cea07cdd566a073bc6ff71d.
//
// Solidity: event QuoteCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterQuoteCreated(opts *bind.FilterOpts, pubId []*big.Int) (*V2EventsQuoteCreatedIterator, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "QuoteCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsQuoteCreatedIterator{contract: _V2Events.contract, event: "QuoteCreated", logs: logs, sub: sub}, nil
}

// WatchQuoteCreated is a free log subscription operation binding the contract event 0x90991d5410a24294cd806880f6e27b11ebe48fbb0cea07cdd566a073bc6ff71d.
//
// Solidity: event QuoteCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchQuoteCreated(opts *bind.WatchOpts, sink chan<- *V2EventsQuoteCreated, pubId []*big.Int) (event.Subscription, error) {

	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "QuoteCreated", pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsQuoteCreated)
				if err := _V2Events.contract.UnpackLog(event, "QuoteCreated", log); err != nil {
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

// ParseQuoteCreated is a log parse operation binding the contract event 0x90991d5410a24294cd806880f6e27b11ebe48fbb0cea07cdd566a073bc6ff71d.
//
// Solidity: event QuoteCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, uint256 indexed pubId, bytes referenceModuleReturnData, bytes[] actionModulesInitReturnDatas, bytes referenceModuleInitReturnData, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseQuoteCreated(log types.Log) (*V2EventsQuoteCreated, error) {
	event := new(V2EventsQuoteCreated)
	if err := _V2Events.contract.UnpackLog(event, "QuoteCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the V2Events contract.
type V2EventsStateSetIterator struct {
	Event *V2EventsStateSet // Event containing the contract specifics and raw log

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
func (it *V2EventsStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsStateSet)
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
		it.Event = new(V2EventsStateSet)
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
func (it *V2EventsStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsStateSet represents a StateSet event raised by the V2Events contract.
type V2EventsStateSet struct {
	Caller    common.Address
	PrevState uint8
	NewState  uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterStateSet(opts *bind.FilterOpts, caller []common.Address, prevState []uint8, newState []uint8) (*V2EventsStateSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsStateSetIterator{contract: _V2Events.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *V2EventsStateSet, caller []common.Address, prevState []uint8, newState []uint8) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsStateSet)
				if err := _V2Events.contract.UnpackLog(event, "StateSet", log); err != nil {
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

// ParseStateSet is a log parse operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseStateSet(log types.Log) (*V2EventsStateSet, error) {
	event := new(V2EventsStateSet)
	if err := _V2Events.contract.UnpackLog(event, "StateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsTokenGuardianStateChangedIterator is returned from FilterTokenGuardianStateChanged and is used to iterate over the raw logs and unpacked data for TokenGuardianStateChanged events raised by the V2Events contract.
type V2EventsTokenGuardianStateChangedIterator struct {
	Event *V2EventsTokenGuardianStateChanged // Event containing the contract specifics and raw log

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
func (it *V2EventsTokenGuardianStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsTokenGuardianStateChanged)
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
		it.Event = new(V2EventsTokenGuardianStateChanged)
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
func (it *V2EventsTokenGuardianStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsTokenGuardianStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsTokenGuardianStateChanged represents a TokenGuardianStateChanged event raised by the V2Events contract.
type V2EventsTokenGuardianStateChanged struct {
	Wallet                          common.Address
	Enabled                         bool
	TokenGuardianDisablingTimestamp *big.Int
	Timestamp                       *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTokenGuardianStateChanged is a free log retrieval operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterTokenGuardianStateChanged(opts *bind.FilterOpts, wallet []common.Address, enabled []bool) (*V2EventsTokenGuardianStateChangedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "TokenGuardianStateChanged", walletRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsTokenGuardianStateChangedIterator{contract: _V2Events.contract, event: "TokenGuardianStateChanged", logs: logs, sub: sub}, nil
}

// WatchTokenGuardianStateChanged is a free log subscription operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchTokenGuardianStateChanged(opts *bind.WatchOpts, sink chan<- *V2EventsTokenGuardianStateChanged, wallet []common.Address, enabled []bool) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "TokenGuardianStateChanged", walletRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsTokenGuardianStateChanged)
				if err := _V2Events.contract.UnpackLog(event, "TokenGuardianStateChanged", log); err != nil {
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

// ParseTokenGuardianStateChanged is a log parse operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseTokenGuardianStateChanged(log types.Log) (*V2EventsTokenGuardianStateChanged, error) {
	event := new(V2EventsTokenGuardianStateChanged)
	if err := _V2Events.contract.UnpackLog(event, "TokenGuardianStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsTreasuryFeeSetIterator is returned from FilterTreasuryFeeSet and is used to iterate over the raw logs and unpacked data for TreasuryFeeSet events raised by the V2Events contract.
type V2EventsTreasuryFeeSetIterator struct {
	Event *V2EventsTreasuryFeeSet // Event containing the contract specifics and raw log

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
func (it *V2EventsTreasuryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsTreasuryFeeSet)
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
		it.Event = new(V2EventsTreasuryFeeSet)
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
func (it *V2EventsTreasuryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsTreasuryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsTreasuryFeeSet represents a TreasuryFeeSet event raised by the V2Events contract.
type V2EventsTreasuryFeeSet struct {
	PrevTreasuryFee uint16
	NewTreasuryFee  uint16
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTreasuryFeeSet is a free log retrieval operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterTreasuryFeeSet(opts *bind.FilterOpts, prevTreasuryFee []uint16, newTreasuryFee []uint16) (*V2EventsTreasuryFeeSetIterator, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "TreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsTreasuryFeeSetIterator{contract: _V2Events.contract, event: "TreasuryFeeSet", logs: logs, sub: sub}, nil
}

// WatchTreasuryFeeSet is a free log subscription operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchTreasuryFeeSet(opts *bind.WatchOpts, sink chan<- *V2EventsTreasuryFeeSet, prevTreasuryFee []uint16, newTreasuryFee []uint16) (event.Subscription, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "TreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsTreasuryFeeSet)
				if err := _V2Events.contract.UnpackLog(event, "TreasuryFeeSet", log); err != nil {
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

// ParseTreasuryFeeSet is a log parse operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseTreasuryFeeSet(log types.Log) (*V2EventsTreasuryFeeSet, error) {
	event := new(V2EventsTreasuryFeeSet)
	if err := _V2Events.contract.UnpackLog(event, "TreasuryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsTreasurySetIterator is returned from FilterTreasurySet and is used to iterate over the raw logs and unpacked data for TreasurySet events raised by the V2Events contract.
type V2EventsTreasurySetIterator struct {
	Event *V2EventsTreasurySet // Event containing the contract specifics and raw log

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
func (it *V2EventsTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsTreasurySet)
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
		it.Event = new(V2EventsTreasurySet)
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
func (it *V2EventsTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsTreasurySet represents a TreasurySet event raised by the V2Events contract.
type V2EventsTreasurySet struct {
	PrevTreasury common.Address
	NewTreasury  common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTreasurySet is a free log retrieval operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterTreasurySet(opts *bind.FilterOpts, prevTreasury []common.Address, newTreasury []common.Address) (*V2EventsTreasurySetIterator, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "TreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsTreasurySetIterator{contract: _V2Events.contract, event: "TreasurySet", logs: logs, sub: sub}, nil
}

// WatchTreasurySet is a free log subscription operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchTreasurySet(opts *bind.WatchOpts, sink chan<- *V2EventsTreasurySet, prevTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "TreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsTreasurySet)
				if err := _V2Events.contract.UnpackLog(event, "TreasurySet", log); err != nil {
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

// ParseTreasurySet is a log parse operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseTreasurySet(log types.Log) (*V2EventsTreasurySet, error) {
	event := new(V2EventsTreasurySet)
	if err := _V2Events.contract.UnpackLog(event, "TreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsUnblockedIterator is returned from FilterUnblocked and is used to iterate over the raw logs and unpacked data for Unblocked events raised by the V2Events contract.
type V2EventsUnblockedIterator struct {
	Event *V2EventsUnblocked // Event containing the contract specifics and raw log

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
func (it *V2EventsUnblockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsUnblocked)
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
		it.Event = new(V2EventsUnblocked)
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
func (it *V2EventsUnblockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsUnblockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsUnblocked represents a Unblocked event raised by the V2Events contract.
type V2EventsUnblocked struct {
	ByProfileId          *big.Int
	IdOfProfileUnblocked *big.Int
	TransactionExecutor  common.Address
	Timestamp            *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUnblocked is a free log retrieval operation binding the contract event 0x8beee0a04ecf0572a57773511663b9baa3f7f6fce0e01112a1bf1a354cddc6a6.
//
// Solidity: event Unblocked(uint256 indexed byProfileId, uint256 idOfProfileUnblocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterUnblocked(opts *bind.FilterOpts, byProfileId []*big.Int) (*V2EventsUnblockedIterator, error) {

	var byProfileIdRule []interface{}
	for _, byProfileIdItem := range byProfileId {
		byProfileIdRule = append(byProfileIdRule, byProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "Unblocked", byProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsUnblockedIterator{contract: _V2Events.contract, event: "Unblocked", logs: logs, sub: sub}, nil
}

// WatchUnblocked is a free log subscription operation binding the contract event 0x8beee0a04ecf0572a57773511663b9baa3f7f6fce0e01112a1bf1a354cddc6a6.
//
// Solidity: event Unblocked(uint256 indexed byProfileId, uint256 idOfProfileUnblocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchUnblocked(opts *bind.WatchOpts, sink chan<- *V2EventsUnblocked, byProfileId []*big.Int) (event.Subscription, error) {

	var byProfileIdRule []interface{}
	for _, byProfileIdItem := range byProfileId {
		byProfileIdRule = append(byProfileIdRule, byProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "Unblocked", byProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsUnblocked)
				if err := _V2Events.contract.UnpackLog(event, "Unblocked", log); err != nil {
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

// ParseUnblocked is a log parse operation binding the contract event 0x8beee0a04ecf0572a57773511663b9baa3f7f6fce0e01112a1bf1a354cddc6a6.
//
// Solidity: event Unblocked(uint256 indexed byProfileId, uint256 idOfProfileUnblocked, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseUnblocked(log types.Log) (*V2EventsUnblocked, error) {
	event := new(V2EventsUnblocked)
	if err := _V2Events.contract.UnpackLog(event, "Unblocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2EventsUnfollowedIterator is returned from FilterUnfollowed and is used to iterate over the raw logs and unpacked data for Unfollowed events raised by the V2Events contract.
type V2EventsUnfollowedIterator struct {
	Event *V2EventsUnfollowed // Event containing the contract specifics and raw log

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
func (it *V2EventsUnfollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2EventsUnfollowed)
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
		it.Event = new(V2EventsUnfollowed)
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
func (it *V2EventsUnfollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2EventsUnfollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2EventsUnfollowed represents a Unfollowed event raised by the V2Events contract.
type V2EventsUnfollowed struct {
	UnfollowerProfileId   *big.Int
	IdOfProfileUnfollowed *big.Int
	TransactionExecutor   common.Address
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUnfollowed is a free log retrieval operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) FilterUnfollowed(opts *bind.FilterOpts, unfollowerProfileId []*big.Int) (*V2EventsUnfollowedIterator, error) {

	var unfollowerProfileIdRule []interface{}
	for _, unfollowerProfileIdItem := range unfollowerProfileId {
		unfollowerProfileIdRule = append(unfollowerProfileIdRule, unfollowerProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.FilterLogs(opts, "Unfollowed", unfollowerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2EventsUnfollowedIterator{contract: _V2Events.contract, event: "Unfollowed", logs: logs, sub: sub}, nil
}

// WatchUnfollowed is a free log subscription operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) WatchUnfollowed(opts *bind.WatchOpts, sink chan<- *V2EventsUnfollowed, unfollowerProfileId []*big.Int) (event.Subscription, error) {

	var unfollowerProfileIdRule []interface{}
	for _, unfollowerProfileIdItem := range unfollowerProfileId {
		unfollowerProfileIdRule = append(unfollowerProfileIdRule, unfollowerProfileIdItem)
	}

	logs, sub, err := _V2Events.contract.WatchLogs(opts, "Unfollowed", unfollowerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2EventsUnfollowed)
				if err := _V2Events.contract.UnpackLog(event, "Unfollowed", log); err != nil {
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

// ParseUnfollowed is a log parse operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2Events *V2EventsFilterer) ParseUnfollowed(log types.Log) (*V2EventsUnfollowed, error) {
	event := new(V2EventsUnfollowed)
	if err := _V2Events.contract.UnpackLog(event, "Unfollowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
