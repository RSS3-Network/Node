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

// CollectPublicationActionCollectData is an auto generated low-level Go binding around an user-defined struct.
type CollectPublicationActionCollectData struct {
	CollectModule common.Address
	CollectNFT    common.Address
}

// TypesProcessActionParams is an auto generated low-level Go binding around an user-defined struct.
type TypesProcessActionParams struct {
	PublicationActedProfileId *big.Int
	PublicationActedId        *big.Int
	ActorProfileId            *big.Int
	ActorProfileOwner         common.Address
	TransactionExecutor       common.Address
	ReferrerProfileIds        []*big.Int
	ReferrerPubIds            []*big.Int
	ReferrerPubTypes          []uint8
	ActionModuleData          []byte
}

// V2CollectPublicationActionMetaData contains all meta data concerning the V2CollectPublicationAction contract.
var V2CollectPublicationActionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"moduleOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CollectNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCollectModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotHub\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectModuleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectedProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectedPubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectorProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectActionData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectActionResult\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Collected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"COLLECT_NFT_IMPL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectData\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"}],\"internalType\":\"structCollectPublicationAction.CollectData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModuleMetadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"moduleOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"initializePublicationAction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"isCollectModuleRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationActedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationActedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"actorProfileOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"enumTypes.PublicationType[]\",\"name\":\"referrerPubTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes\",\"name\":\"actionModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.ProcessActionParams\",\"name\":\"processActionParams\",\"type\":\"tuple\"}],\"name\":\"processPublicationAction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"registerCollectModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_metadataURI\",\"type\":\"string\"}],\"name\":\"setModuleMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"verifyCollectModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V2CollectPublicationActionABI is the input ABI used to generate the binding from.
// Deprecated: Use V2CollectPublicationActionMetaData.ABI instead.
var V2CollectPublicationActionABI = V2CollectPublicationActionMetaData.ABI

// V2CollectPublicationAction is an auto generated Go binding around an Ethereum contract.
type V2CollectPublicationAction struct {
	V2CollectPublicationActionCaller     // Read-only binding to the contract
	V2CollectPublicationActionTransactor // Write-only binding to the contract
	V2CollectPublicationActionFilterer   // Log filterer for contract events
}

// V2CollectPublicationActionCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2CollectPublicationActionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2CollectPublicationActionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2CollectPublicationActionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2CollectPublicationActionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2CollectPublicationActionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2CollectPublicationActionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2CollectPublicationActionSession struct {
	Contract     *V2CollectPublicationAction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// V2CollectPublicationActionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2CollectPublicationActionCallerSession struct {
	Contract *V2CollectPublicationActionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// V2CollectPublicationActionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2CollectPublicationActionTransactorSession struct {
	Contract     *V2CollectPublicationActionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// V2CollectPublicationActionRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2CollectPublicationActionRaw struct {
	Contract *V2CollectPublicationAction // Generic contract binding to access the raw methods on
}

// V2CollectPublicationActionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2CollectPublicationActionCallerRaw struct {
	Contract *V2CollectPublicationActionCaller // Generic read-only contract binding to access the raw methods on
}

// V2CollectPublicationActionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2CollectPublicationActionTransactorRaw struct {
	Contract *V2CollectPublicationActionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2CollectPublicationAction creates a new instance of V2CollectPublicationAction, bound to a specific deployed contract.
func NewV2CollectPublicationAction(address common.Address, backend bind.ContractBackend) (*V2CollectPublicationAction, error) {
	contract, err := bindV2CollectPublicationAction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationAction{V2CollectPublicationActionCaller: V2CollectPublicationActionCaller{contract: contract}, V2CollectPublicationActionTransactor: V2CollectPublicationActionTransactor{contract: contract}, V2CollectPublicationActionFilterer: V2CollectPublicationActionFilterer{contract: contract}}, nil
}

// NewV2CollectPublicationActionCaller creates a new read-only instance of V2CollectPublicationAction, bound to a specific deployed contract.
func NewV2CollectPublicationActionCaller(address common.Address, caller bind.ContractCaller) (*V2CollectPublicationActionCaller, error) {
	contract, err := bindV2CollectPublicationAction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionCaller{contract: contract}, nil
}

// NewV2CollectPublicationActionTransactor creates a new write-only instance of V2CollectPublicationAction, bound to a specific deployed contract.
func NewV2CollectPublicationActionTransactor(address common.Address, transactor bind.ContractTransactor) (*V2CollectPublicationActionTransactor, error) {
	contract, err := bindV2CollectPublicationAction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionTransactor{contract: contract}, nil
}

// NewV2CollectPublicationActionFilterer creates a new log filterer instance of V2CollectPublicationAction, bound to a specific deployed contract.
func NewV2CollectPublicationActionFilterer(address common.Address, filterer bind.ContractFilterer) (*V2CollectPublicationActionFilterer, error) {
	contract, err := bindV2CollectPublicationAction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionFilterer{contract: contract}, nil
}

// bindV2CollectPublicationAction binds a generic wrapper to an already deployed contract.
func bindV2CollectPublicationAction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2CollectPublicationActionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2CollectPublicationAction *V2CollectPublicationActionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2CollectPublicationAction.Contract.V2CollectPublicationActionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2CollectPublicationAction *V2CollectPublicationActionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.V2CollectPublicationActionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2CollectPublicationAction *V2CollectPublicationActionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.V2CollectPublicationActionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2CollectPublicationAction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.contract.Transact(opts, method, params...)
}

// COLLECTNFTIMPL is a free data retrieval call binding the contract method 0x00a7c89b.
//
// Solidity: function COLLECT_NFT_IMPL() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) COLLECTNFTIMPL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "COLLECT_NFT_IMPL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// COLLECTNFTIMPL is a free data retrieval call binding the contract method 0x00a7c89b.
//
// Solidity: function COLLECT_NFT_IMPL() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) COLLECTNFTIMPL() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.COLLECTNFTIMPL(&_V2CollectPublicationAction.CallOpts)
}

// COLLECTNFTIMPL is a free data retrieval call binding the contract method 0x00a7c89b.
//
// Solidity: function COLLECT_NFT_IMPL() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) COLLECTNFTIMPL() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.COLLECTNFTIMPL(&_V2CollectPublicationAction.CallOpts)
}

// HUB is a free data retrieval call binding the contract method 0xa4c52b86.
//
// Solidity: function HUB() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) HUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HUB is a free data retrieval call binding the contract method 0xa4c52b86.
//
// Solidity: function HUB() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) HUB() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.HUB(&_V2CollectPublicationAction.CallOpts)
}

// HUB is a free data retrieval call binding the contract method 0xa4c52b86.
//
// Solidity: function HUB() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) HUB() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.HUB(&_V2CollectPublicationAction.CallOpts)
}

// GetCollectData is a free data retrieval call binding the contract method 0xe8336893.
//
// Solidity: function getCollectData(uint256 profileId, uint256 pubId) view returns((address,address))
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) GetCollectData(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (CollectPublicationActionCollectData, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "getCollectData", profileId, pubId)

	if err != nil {
		return *new(CollectPublicationActionCollectData), err
	}

	out0 := *abi.ConvertType(out[0], new(CollectPublicationActionCollectData)).(*CollectPublicationActionCollectData)

	return out0, err

}

// GetCollectData is a free data retrieval call binding the contract method 0xe8336893.
//
// Solidity: function getCollectData(uint256 profileId, uint256 pubId) view returns((address,address))
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) GetCollectData(profileId *big.Int, pubId *big.Int) (CollectPublicationActionCollectData, error) {
	return _V2CollectPublicationAction.Contract.GetCollectData(&_V2CollectPublicationAction.CallOpts, profileId, pubId)
}

// GetCollectData is a free data retrieval call binding the contract method 0xe8336893.
//
// Solidity: function getCollectData(uint256 profileId, uint256 pubId) view returns((address,address))
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) GetCollectData(profileId *big.Int, pubId *big.Int) (CollectPublicationActionCollectData, error) {
	return _V2CollectPublicationAction.Contract.GetCollectData(&_V2CollectPublicationAction.CallOpts, profileId, pubId)
}

// GetModuleMetadataURI is a free data retrieval call binding the contract method 0xce90d52e.
//
// Solidity: function getModuleMetadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) GetModuleMetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "getModuleMetadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetModuleMetadataURI is a free data retrieval call binding the contract method 0xce90d52e.
//
// Solidity: function getModuleMetadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) GetModuleMetadataURI() (string, error) {
	return _V2CollectPublicationAction.Contract.GetModuleMetadataURI(&_V2CollectPublicationAction.CallOpts)
}

// GetModuleMetadataURI is a free data retrieval call binding the contract method 0xce90d52e.
//
// Solidity: function getModuleMetadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) GetModuleMetadataURI() (string, error) {
	return _V2CollectPublicationAction.Contract.GetModuleMetadataURI(&_V2CollectPublicationAction.CallOpts)
}

// IsCollectModuleRegistered is a free data retrieval call binding the contract method 0xf32ffcbc.
//
// Solidity: function isCollectModuleRegistered(address collectModule) view returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) IsCollectModuleRegistered(opts *bind.CallOpts, collectModule common.Address) (bool, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "isCollectModuleRegistered", collectModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollectModuleRegistered is a free data retrieval call binding the contract method 0xf32ffcbc.
//
// Solidity: function isCollectModuleRegistered(address collectModule) view returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) IsCollectModuleRegistered(collectModule common.Address) (bool, error) {
	return _V2CollectPublicationAction.Contract.IsCollectModuleRegistered(&_V2CollectPublicationAction.CallOpts, collectModule)
}

// IsCollectModuleRegistered is a free data retrieval call binding the contract method 0xf32ffcbc.
//
// Solidity: function isCollectModuleRegistered(address collectModule) view returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) IsCollectModuleRegistered(collectModule common.Address) (bool, error) {
	return _V2CollectPublicationAction.Contract.IsCollectModuleRegistered(&_V2CollectPublicationAction.CallOpts, collectModule)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) MetadataURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "metadataURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) MetadataURI() (string, error) {
	return _V2CollectPublicationAction.Contract.MetadataURI(&_V2CollectPublicationAction.CallOpts)
}

// MetadataURI is a free data retrieval call binding the contract method 0x03ee438c.
//
// Solidity: function metadataURI() view returns(string)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) MetadataURI() (string, error) {
	return _V2CollectPublicationAction.Contract.MetadataURI(&_V2CollectPublicationAction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) Owner() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.Owner(&_V2CollectPublicationAction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) Owner() (common.Address, error) {
	return _V2CollectPublicationAction.Contract.Owner(&_V2CollectPublicationAction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _V2CollectPublicationAction.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _V2CollectPublicationAction.Contract.SupportsInterface(&_V2CollectPublicationAction.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _V2CollectPublicationAction.Contract.SupportsInterface(&_V2CollectPublicationAction.CallOpts, interfaceID)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address moduleOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) Initialize(opts *bind.TransactOpts, moduleOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "initialize", moduleOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address moduleOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) Initialize(moduleOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.Initialize(&_V2CollectPublicationAction.TransactOpts, moduleOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address moduleOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) Initialize(moduleOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.Initialize(&_V2CollectPublicationAction.TransactOpts, moduleOwner)
}

// InitializePublicationAction is a paid mutator transaction binding the contract method 0x0b960220.
//
// Solidity: function initializePublicationAction(uint256 profileId, uint256 pubId, address transactionExecutor, bytes data) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) InitializePublicationAction(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, transactionExecutor common.Address, data []byte) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "initializePublicationAction", profileId, pubId, transactionExecutor, data)
}

// InitializePublicationAction is a paid mutator transaction binding the contract method 0x0b960220.
//
// Solidity: function initializePublicationAction(uint256 profileId, uint256 pubId, address transactionExecutor, bytes data) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) InitializePublicationAction(profileId *big.Int, pubId *big.Int, transactionExecutor common.Address, data []byte) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.InitializePublicationAction(&_V2CollectPublicationAction.TransactOpts, profileId, pubId, transactionExecutor, data)
}

// InitializePublicationAction is a paid mutator transaction binding the contract method 0x0b960220.
//
// Solidity: function initializePublicationAction(uint256 profileId, uint256 pubId, address transactionExecutor, bytes data) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) InitializePublicationAction(profileId *big.Int, pubId *big.Int, transactionExecutor common.Address, data []byte) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.InitializePublicationAction(&_V2CollectPublicationAction.TransactOpts, profileId, pubId, transactionExecutor, data)
}

// ProcessPublicationAction is a paid mutator transaction binding the contract method 0x90ce6e08.
//
// Solidity: function processPublicationAction((uint256,uint256,uint256,address,address,uint256[],uint256[],uint8[],bytes) processActionParams) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) ProcessPublicationAction(opts *bind.TransactOpts, processActionParams TypesProcessActionParams) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "processPublicationAction", processActionParams)
}

// ProcessPublicationAction is a paid mutator transaction binding the contract method 0x90ce6e08.
//
// Solidity: function processPublicationAction((uint256,uint256,uint256,address,address,uint256[],uint256[],uint8[],bytes) processActionParams) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) ProcessPublicationAction(processActionParams TypesProcessActionParams) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.ProcessPublicationAction(&_V2CollectPublicationAction.TransactOpts, processActionParams)
}

// ProcessPublicationAction is a paid mutator transaction binding the contract method 0x90ce6e08.
//
// Solidity: function processPublicationAction((uint256,uint256,uint256,address,address,uint256[],uint256[],uint8[],bytes) processActionParams) returns(bytes)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) ProcessPublicationAction(processActionParams TypesProcessActionParams) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.ProcessPublicationAction(&_V2CollectPublicationAction.TransactOpts, processActionParams)
}

// RegisterCollectModule is a paid mutator transaction binding the contract method 0x8e8b8c97.
//
// Solidity: function registerCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) RegisterCollectModule(opts *bind.TransactOpts, collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "registerCollectModule", collectModule)
}

// RegisterCollectModule is a paid mutator transaction binding the contract method 0x8e8b8c97.
//
// Solidity: function registerCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) RegisterCollectModule(collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.RegisterCollectModule(&_V2CollectPublicationAction.TransactOpts, collectModule)
}

// RegisterCollectModule is a paid mutator transaction binding the contract method 0x8e8b8c97.
//
// Solidity: function registerCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) RegisterCollectModule(collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.RegisterCollectModule(&_V2CollectPublicationAction.TransactOpts, collectModule)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) RenounceOwnership() (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.RenounceOwnership(&_V2CollectPublicationAction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.RenounceOwnership(&_V2CollectPublicationAction.TransactOpts)
}

// SetModuleMetadataURI is a paid mutator transaction binding the contract method 0x681591c1.
//
// Solidity: function setModuleMetadataURI(string _metadataURI) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) SetModuleMetadataURI(opts *bind.TransactOpts, _metadataURI string) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "setModuleMetadataURI", _metadataURI)
}

// SetModuleMetadataURI is a paid mutator transaction binding the contract method 0x681591c1.
//
// Solidity: function setModuleMetadataURI(string _metadataURI) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) SetModuleMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.SetModuleMetadataURI(&_V2CollectPublicationAction.TransactOpts, _metadataURI)
}

// SetModuleMetadataURI is a paid mutator transaction binding the contract method 0x681591c1.
//
// Solidity: function setModuleMetadataURI(string _metadataURI) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) SetModuleMetadataURI(_metadataURI string) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.SetModuleMetadataURI(&_V2CollectPublicationAction.TransactOpts, _metadataURI)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.TransferOwnership(&_V2CollectPublicationAction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.TransferOwnership(&_V2CollectPublicationAction.TransactOpts, newOwner)
}

// VerifyCollectModule is a paid mutator transaction binding the contract method 0x895e9c16.
//
// Solidity: function verifyCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactor) VerifyCollectModule(opts *bind.TransactOpts, collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.contract.Transact(opts, "verifyCollectModule", collectModule)
}

// VerifyCollectModule is a paid mutator transaction binding the contract method 0x895e9c16.
//
// Solidity: function verifyCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionSession) VerifyCollectModule(collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.VerifyCollectModule(&_V2CollectPublicationAction.TransactOpts, collectModule)
}

// VerifyCollectModule is a paid mutator transaction binding the contract method 0x895e9c16.
//
// Solidity: function verifyCollectModule(address collectModule) returns(bool)
func (_V2CollectPublicationAction *V2CollectPublicationActionTransactorSession) VerifyCollectModule(collectModule common.Address) (*types.Transaction, error) {
	return _V2CollectPublicationAction.Contract.VerifyCollectModule(&_V2CollectPublicationAction.TransactOpts, collectModule)
}

// V2CollectPublicationActionCollectModuleRegisteredIterator is returned from FilterCollectModuleRegistered and is used to iterate over the raw logs and unpacked data for CollectModuleRegistered events raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollectModuleRegisteredIterator struct {
	Event *V2CollectPublicationActionCollectModuleRegistered // Event containing the contract specifics and raw log

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
func (it *V2CollectPublicationActionCollectModuleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2CollectPublicationActionCollectModuleRegistered)
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
		it.Event = new(V2CollectPublicationActionCollectModuleRegistered)
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
func (it *V2CollectPublicationActionCollectModuleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2CollectPublicationActionCollectModuleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2CollectPublicationActionCollectModuleRegistered represents a CollectModuleRegistered event raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollectModuleRegistered struct {
	CollectModule common.Address
	Metadata      string
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollectModuleRegistered is a free log retrieval operation binding the contract event 0xf6b57d3a49fa5be1fc25346ee0b64048a6a4734e50280e64080e3de8dca50267.
//
// Solidity: event CollectModuleRegistered(address collectModule, string metadata, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) FilterCollectModuleRegistered(opts *bind.FilterOpts) (*V2CollectPublicationActionCollectModuleRegisteredIterator, error) {

	logs, sub, err := _V2CollectPublicationAction.contract.FilterLogs(opts, "CollectModuleRegistered")
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionCollectModuleRegisteredIterator{contract: _V2CollectPublicationAction.contract, event: "CollectModuleRegistered", logs: logs, sub: sub}, nil
}

// WatchCollectModuleRegistered is a free log subscription operation binding the contract event 0xf6b57d3a49fa5be1fc25346ee0b64048a6a4734e50280e64080e3de8dca50267.
//
// Solidity: event CollectModuleRegistered(address collectModule, string metadata, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) WatchCollectModuleRegistered(opts *bind.WatchOpts, sink chan<- *V2CollectPublicationActionCollectModuleRegistered) (event.Subscription, error) {

	logs, sub, err := _V2CollectPublicationAction.contract.WatchLogs(opts, "CollectModuleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2CollectPublicationActionCollectModuleRegistered)
				if err := _V2CollectPublicationAction.contract.UnpackLog(event, "CollectModuleRegistered", log); err != nil {
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

// ParseCollectModuleRegistered is a log parse operation binding the contract event 0xf6b57d3a49fa5be1fc25346ee0b64048a6a4734e50280e64080e3de8dca50267.
//
// Solidity: event CollectModuleRegistered(address collectModule, string metadata, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) ParseCollectModuleRegistered(log types.Log) (*V2CollectPublicationActionCollectModuleRegistered, error) {
	event := new(V2CollectPublicationActionCollectModuleRegistered)
	if err := _V2CollectPublicationAction.contract.UnpackLog(event, "CollectModuleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2CollectPublicationActionCollectNFTDeployedIterator is returned from FilterCollectNFTDeployed and is used to iterate over the raw logs and unpacked data for CollectNFTDeployed events raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollectNFTDeployedIterator struct {
	Event *V2CollectPublicationActionCollectNFTDeployed // Event containing the contract specifics and raw log

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
func (it *V2CollectPublicationActionCollectNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2CollectPublicationActionCollectNFTDeployed)
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
		it.Event = new(V2CollectPublicationActionCollectNFTDeployed)
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
func (it *V2CollectPublicationActionCollectNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2CollectPublicationActionCollectNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2CollectPublicationActionCollectNFTDeployed represents a CollectNFTDeployed event raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollectNFTDeployed struct {
	ProfileId  *big.Int
	PubId      *big.Int
	CollectNFT common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTDeployed is a free log retrieval operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) FilterCollectNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (*V2CollectPublicationActionCollectNFTDeployedIterator, error) {

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

	logs, sub, err := _V2CollectPublicationAction.contract.FilterLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionCollectNFTDeployedIterator{contract: _V2CollectPublicationAction.contract, event: "CollectNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchCollectNFTDeployed is a free log subscription operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) WatchCollectNFTDeployed(opts *bind.WatchOpts, sink chan<- *V2CollectPublicationActionCollectNFTDeployed, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _V2CollectPublicationAction.contract.WatchLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2CollectPublicationActionCollectNFTDeployed)
				if err := _V2CollectPublicationAction.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
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

// ParseCollectNFTDeployed is a log parse operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) ParseCollectNFTDeployed(log types.Log) (*V2CollectPublicationActionCollectNFTDeployed, error) {
	event := new(V2CollectPublicationActionCollectNFTDeployed)
	if err := _V2CollectPublicationAction.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2CollectPublicationActionCollectedIterator is returned from FilterCollected and is used to iterate over the raw logs and unpacked data for Collected events raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollectedIterator struct {
	Event *V2CollectPublicationActionCollected // Event containing the contract specifics and raw log

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
func (it *V2CollectPublicationActionCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2CollectPublicationActionCollected)
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
		it.Event = new(V2CollectPublicationActionCollected)
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
func (it *V2CollectPublicationActionCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2CollectPublicationActionCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2CollectPublicationActionCollected represents a Collected event raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionCollected struct {
	CollectedProfileId  *big.Int
	CollectedPubId      *big.Int
	CollectorProfileId  *big.Int
	NftRecipient        common.Address
	CollectActionData   []byte
	CollectActionResult []byte
	CollectNFT          common.Address
	TokenId             *big.Int
	TransactionExecutor common.Address
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterCollected is a free log retrieval operation binding the contract event 0x4b220a4de1946418b442a3659c114b9f74ed61e23509c34f97bbe8f2f1d645e0.
//
// Solidity: event Collected(uint256 indexed collectedProfileId, uint256 indexed collectedPubId, uint256 indexed collectorProfileId, address nftRecipient, bytes collectActionData, bytes collectActionResult, address collectNFT, uint256 tokenId, address transactionExecutor, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) FilterCollected(opts *bind.FilterOpts, collectedProfileId []*big.Int, collectedPubId []*big.Int, collectorProfileId []*big.Int) (*V2CollectPublicationActionCollectedIterator, error) {

	var collectedProfileIdRule []interface{}
	for _, collectedProfileIdItem := range collectedProfileId {
		collectedProfileIdRule = append(collectedProfileIdRule, collectedProfileIdItem)
	}
	var collectedPubIdRule []interface{}
	for _, collectedPubIdItem := range collectedPubId {
		collectedPubIdRule = append(collectedPubIdRule, collectedPubIdItem)
	}
	var collectorProfileIdRule []interface{}
	for _, collectorProfileIdItem := range collectorProfileId {
		collectorProfileIdRule = append(collectorProfileIdRule, collectorProfileIdItem)
	}

	logs, sub, err := _V2CollectPublicationAction.contract.FilterLogs(opts, "Collected", collectedProfileIdRule, collectedPubIdRule, collectorProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionCollectedIterator{contract: _V2CollectPublicationAction.contract, event: "Collected", logs: logs, sub: sub}, nil
}

// WatchCollected is a free log subscription operation binding the contract event 0x4b220a4de1946418b442a3659c114b9f74ed61e23509c34f97bbe8f2f1d645e0.
//
// Solidity: event Collected(uint256 indexed collectedProfileId, uint256 indexed collectedPubId, uint256 indexed collectorProfileId, address nftRecipient, bytes collectActionData, bytes collectActionResult, address collectNFT, uint256 tokenId, address transactionExecutor, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) WatchCollected(opts *bind.WatchOpts, sink chan<- *V2CollectPublicationActionCollected, collectedProfileId []*big.Int, collectedPubId []*big.Int, collectorProfileId []*big.Int) (event.Subscription, error) {

	var collectedProfileIdRule []interface{}
	for _, collectedProfileIdItem := range collectedProfileId {
		collectedProfileIdRule = append(collectedProfileIdRule, collectedProfileIdItem)
	}
	var collectedPubIdRule []interface{}
	for _, collectedPubIdItem := range collectedPubId {
		collectedPubIdRule = append(collectedPubIdRule, collectedPubIdItem)
	}
	var collectorProfileIdRule []interface{}
	for _, collectorProfileIdItem := range collectorProfileId {
		collectorProfileIdRule = append(collectorProfileIdRule, collectorProfileIdItem)
	}

	logs, sub, err := _V2CollectPublicationAction.contract.WatchLogs(opts, "Collected", collectedProfileIdRule, collectedPubIdRule, collectorProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2CollectPublicationActionCollected)
				if err := _V2CollectPublicationAction.contract.UnpackLog(event, "Collected", log); err != nil {
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

// ParseCollected is a log parse operation binding the contract event 0x4b220a4de1946418b442a3659c114b9f74ed61e23509c34f97bbe8f2f1d645e0.
//
// Solidity: event Collected(uint256 indexed collectedProfileId, uint256 indexed collectedPubId, uint256 indexed collectorProfileId, address nftRecipient, bytes collectActionData, bytes collectActionResult, address collectNFT, uint256 tokenId, address transactionExecutor, uint256 timestamp)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) ParseCollected(log types.Log) (*V2CollectPublicationActionCollected, error) {
	event := new(V2CollectPublicationActionCollected)
	if err := _V2CollectPublicationAction.contract.UnpackLog(event, "Collected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2CollectPublicationActionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionOwnershipTransferredIterator struct {
	Event *V2CollectPublicationActionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *V2CollectPublicationActionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2CollectPublicationActionOwnershipTransferred)
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
		it.Event = new(V2CollectPublicationActionOwnershipTransferred)
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
func (it *V2CollectPublicationActionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2CollectPublicationActionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2CollectPublicationActionOwnershipTransferred represents a OwnershipTransferred event raised by the V2CollectPublicationAction contract.
type V2CollectPublicationActionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*V2CollectPublicationActionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V2CollectPublicationAction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &V2CollectPublicationActionOwnershipTransferredIterator{contract: _V2CollectPublicationAction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *V2CollectPublicationActionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _V2CollectPublicationAction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2CollectPublicationActionOwnershipTransferred)
				if err := _V2CollectPublicationAction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_V2CollectPublicationAction *V2CollectPublicationActionFilterer) ParseOwnershipTransferred(log types.Log) (*V2CollectPublicationActionOwnershipTransferred, error) {
	event := new(V2CollectPublicationActionOwnershipTransferred)
	if err := _V2CollectPublicationAction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
