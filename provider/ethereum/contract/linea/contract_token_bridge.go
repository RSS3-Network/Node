// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package linea

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

// TokenBridgeMetaData contains all meta data concerning the TokenBridge contract.
var TokenBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AlreadyBridgedToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AlreadyBrigedToNativeTokenSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerIsNotMessageService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"DecimalsAreUnknown\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"permitData\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"permitSelector\",\"type\":\"bytes4\"}],\"name\":\"InvalidPermitData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NativeToBridgedTokenAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NotReserved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"PermitNotAllowingBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"PermitNotFromSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteTokenBridge\",\"type\":\"address\"}],\"name\":\"RemoteTokenBridgeAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ReservedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"StatusAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotDeployed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ZeroAmountNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"BridgingFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"BridgingFinalizedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgingInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgingInitiatedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"customContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"CustomContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"confirmedBy\",\"type\":\"address\"}],\"name\":\"DeploymentConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMessageService\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMessageService\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"MessageServiceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"NewToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nativeToken\",\"type\":\"address\"}],\"name\":\"NewTokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remoteTokenBridge\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"RemoteTokenBridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ReservationRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"bridgeToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_permitData\",\"type\":\"bytes\"}],\"name\":\"bridgeTokenWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"bridgedToNativeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_tokenMetadata\",\"type\":\"bytes\"}],\"name\":\"completeBridging\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"}],\"name\":\"confirmDeployment\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_messageService\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenBeacon\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetChainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_reservedTokens\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageService\",\"outputs\":[{\"internalType\":\"contractIMessageService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nativeToBridgedToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"removeReserved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_targetContract\",\"type\":\"address\"}],\"name\":\"setCustomContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nativeTokens\",\"type\":\"address[]\"}],\"name\":\"setDeployed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageService\",\"type\":\"address\"}],\"name\":\"setMessageService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteTokenBridge\",\"type\":\"address\"}],\"name\":\"setRemoteTokenBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setReserved\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sourceChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBeacon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenBridgeMetaData.ABI instead.
var TokenBridgeABI = TokenBridgeMetaData.ABI

// TokenBridge is an auto generated Go binding around an Ethereum contract.
type TokenBridge struct {
	TokenBridgeCaller     // Read-only binding to the contract
	TokenBridgeTransactor // Write-only binding to the contract
	TokenBridgeFilterer   // Log filterer for contract events
}

// TokenBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenBridgeSession struct {
	Contract     *TokenBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenBridgeCallerSession struct {
	Contract *TokenBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TokenBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenBridgeTransactorSession struct {
	Contract     *TokenBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TokenBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenBridgeRaw struct {
	Contract *TokenBridge // Generic contract binding to access the raw methods on
}

// TokenBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenBridgeCallerRaw struct {
	Contract *TokenBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// TokenBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenBridgeTransactorRaw struct {
	Contract *TokenBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenBridge creates a new instance of TokenBridge, bound to a specific deployed contract.
func NewTokenBridge(address common.Address, backend bind.ContractBackend) (*TokenBridge, error) {
	contract, err := bindTokenBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenBridge{TokenBridgeCaller: TokenBridgeCaller{contract: contract}, TokenBridgeTransactor: TokenBridgeTransactor{contract: contract}, TokenBridgeFilterer: TokenBridgeFilterer{contract: contract}}, nil
}

// NewTokenBridgeCaller creates a new read-only instance of TokenBridge, bound to a specific deployed contract.
func NewTokenBridgeCaller(address common.Address, caller bind.ContractCaller) (*TokenBridgeCaller, error) {
	contract, err := bindTokenBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeCaller{contract: contract}, nil
}

// NewTokenBridgeTransactor creates a new write-only instance of TokenBridge, bound to a specific deployed contract.
func NewTokenBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenBridgeTransactor, error) {
	contract, err := bindTokenBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeTransactor{contract: contract}, nil
}

// NewTokenBridgeFilterer creates a new log filterer instance of TokenBridge, bound to a specific deployed contract.
func NewTokenBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenBridgeFilterer, error) {
	contract, err := bindTokenBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeFilterer{contract: contract}, nil
}

// bindTokenBridge binds a generic wrapper to an already deployed contract.
func bindTokenBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBridge *TokenBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBridge.Contract.TokenBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBridge *TokenBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.Contract.TokenBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBridge *TokenBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBridge.Contract.TokenBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenBridge *TokenBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenBridge *TokenBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenBridge *TokenBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_TokenBridge *TokenBridgeCaller) BridgedToNativeToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "bridgedToNativeToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_TokenBridge *TokenBridgeSession) BridgedToNativeToken(arg0 common.Address) (common.Address, error) {
	return _TokenBridge.Contract.BridgedToNativeToken(&_TokenBridge.CallOpts, arg0)
}

// BridgedToNativeToken is a free data retrieval call binding the contract method 0xca41a247.
//
// Solidity: function bridgedToNativeToken(address ) view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) BridgedToNativeToken(arg0 common.Address) (common.Address, error) {
	return _TokenBridge.Contract.BridgedToNativeToken(&_TokenBridge.CallOpts, arg0)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_TokenBridge *TokenBridgeCaller) MessageService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "messageService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_TokenBridge *TokenBridgeSession) MessageService() (common.Address, error) {
	return _TokenBridge.Contract.MessageService(&_TokenBridge.CallOpts)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) MessageService() (common.Address, error) {
	return _TokenBridge.Contract.MessageService(&_TokenBridge.CallOpts)
}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0x0f6f86ec.
//
// Solidity: function nativeToBridgedToken(uint256 , address ) view returns(address)
func (_TokenBridge *TokenBridgeCaller) NativeToBridgedToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "nativeToBridgedToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0x0f6f86ec.
//
// Solidity: function nativeToBridgedToken(uint256 , address ) view returns(address)
func (_TokenBridge *TokenBridgeSession) NativeToBridgedToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _TokenBridge.Contract.NativeToBridgedToken(&_TokenBridge.CallOpts, arg0, arg1)
}

// NativeToBridgedToken is a free data retrieval call binding the contract method 0x0f6f86ec.
//
// Solidity: function nativeToBridgedToken(uint256 , address ) view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) NativeToBridgedToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _TokenBridge.Contract.NativeToBridgedToken(&_TokenBridge.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridge *TokenBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridge *TokenBridgeSession) Owner() (common.Address, error) {
	return _TokenBridge.Contract.Owner(&_TokenBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) Owner() (common.Address, error) {
	return _TokenBridge.Contract.Owner(&_TokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenBridge *TokenBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenBridge *TokenBridgeSession) Paused() (bool, error) {
	return _TokenBridge.Contract.Paused(&_TokenBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TokenBridge *TokenBridgeCallerSession) Paused() (bool, error) {
	return _TokenBridge.Contract.Paused(&_TokenBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenBridge *TokenBridgeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenBridge *TokenBridgeSession) PendingOwner() (common.Address, error) {
	return _TokenBridge.Contract.PendingOwner(&_TokenBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) PendingOwner() (common.Address, error) {
	return _TokenBridge.Contract.PendingOwner(&_TokenBridge.CallOpts)
}

// RemoteSender is a free data retrieval call binding the contract method 0xa6ef995f.
//
// Solidity: function remoteSender() view returns(address)
func (_TokenBridge *TokenBridgeCaller) RemoteSender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "remoteSender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteSender is a free data retrieval call binding the contract method 0xa6ef995f.
//
// Solidity: function remoteSender() view returns(address)
func (_TokenBridge *TokenBridgeSession) RemoteSender() (common.Address, error) {
	return _TokenBridge.Contract.RemoteSender(&_TokenBridge.CallOpts)
}

// RemoteSender is a free data retrieval call binding the contract method 0xa6ef995f.
//
// Solidity: function remoteSender() view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) RemoteSender() (common.Address, error) {
	return _TokenBridge.Contract.RemoteSender(&_TokenBridge.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeCaller) SourceChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "sourceChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeSession) SourceChainId() (*big.Int, error) {
	return _TokenBridge.Contract.SourceChainId(&_TokenBridge.CallOpts)
}

// SourceChainId is a free data retrieval call binding the contract method 0x1544298e.
//
// Solidity: function sourceChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeCallerSession) SourceChainId() (*big.Int, error) {
	return _TokenBridge.Contract.SourceChainId(&_TokenBridge.CallOpts)
}

// TargetChainId is a free data retrieval call binding the contract method 0x146ffb26.
//
// Solidity: function targetChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeCaller) TargetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "targetChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetChainId is a free data retrieval call binding the contract method 0x146ffb26.
//
// Solidity: function targetChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeSession) TargetChainId() (*big.Int, error) {
	return _TokenBridge.Contract.TargetChainId(&_TokenBridge.CallOpts)
}

// TargetChainId is a free data retrieval call binding the contract method 0x146ffb26.
//
// Solidity: function targetChainId() view returns(uint256)
func (_TokenBridge *TokenBridgeCallerSession) TargetChainId() (*big.Int, error) {
	return _TokenBridge.Contract.TargetChainId(&_TokenBridge.CallOpts)
}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_TokenBridge *TokenBridgeCaller) TokenBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenBridge.contract.Call(opts, &out, "tokenBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_TokenBridge *TokenBridgeSession) TokenBeacon() (common.Address, error) {
	return _TokenBridge.Contract.TokenBeacon(&_TokenBridge.CallOpts)
}

// TokenBeacon is a free data retrieval call binding the contract method 0xccf5a77c.
//
// Solidity: function tokenBeacon() view returns(address)
func (_TokenBridge *TokenBridgeCallerSession) TokenBeacon() (common.Address, error) {
	return _TokenBridge.Contract.TokenBeacon(&_TokenBridge.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenBridge *TokenBridgeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenBridge *TokenBridgeSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenBridge.Contract.AcceptOwnership(&_TokenBridge.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenBridge *TokenBridgeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenBridge.Contract.AcceptOwnership(&_TokenBridge.TransactOpts)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_TokenBridge *TokenBridgeTransactor) BridgeToken(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "bridgeToken", _token, _amount, _recipient)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_TokenBridge *TokenBridgeSession) BridgeToken(_token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.BridgeToken(&_TokenBridge.TransactOpts, _token, _amount, _recipient)
}

// BridgeToken is a paid mutator transaction binding the contract method 0x522ea81a.
//
// Solidity: function bridgeToken(address _token, uint256 _amount, address _recipient) payable returns()
func (_TokenBridge *TokenBridgeTransactorSession) BridgeToken(_token common.Address, _amount *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.BridgeToken(&_TokenBridge.TransactOpts, _token, _amount, _recipient)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_TokenBridge *TokenBridgeTransactor) BridgeTokenWithPermit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "bridgeTokenWithPermit", _token, _amount, _recipient, _permitData)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_TokenBridge *TokenBridgeSession) BridgeTokenWithPermit(_token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _TokenBridge.Contract.BridgeTokenWithPermit(&_TokenBridge.TransactOpts, _token, _amount, _recipient, _permitData)
}

// BridgeTokenWithPermit is a paid mutator transaction binding the contract method 0xdfa96efb.
//
// Solidity: function bridgeTokenWithPermit(address _token, uint256 _amount, address _recipient, bytes _permitData) payable returns()
func (_TokenBridge *TokenBridgeTransactorSession) BridgeTokenWithPermit(_token common.Address, _amount *big.Int, _recipient common.Address, _permitData []byte) (*types.Transaction, error) {
	return _TokenBridge.Contract.BridgeTokenWithPermit(&_TokenBridge.TransactOpts, _token, _amount, _recipient, _permitData)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xe4d27451.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, uint256 _chainId, bytes _tokenMetadata) returns()
func (_TokenBridge *TokenBridgeTransactor) CompleteBridging(opts *bind.TransactOpts, _nativeToken common.Address, _amount *big.Int, _recipient common.Address, _chainId *big.Int, _tokenMetadata []byte) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "completeBridging", _nativeToken, _amount, _recipient, _chainId, _tokenMetadata)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xe4d27451.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, uint256 _chainId, bytes _tokenMetadata) returns()
func (_TokenBridge *TokenBridgeSession) CompleteBridging(_nativeToken common.Address, _amount *big.Int, _recipient common.Address, _chainId *big.Int, _tokenMetadata []byte) (*types.Transaction, error) {
	return _TokenBridge.Contract.CompleteBridging(&_TokenBridge.TransactOpts, _nativeToken, _amount, _recipient, _chainId, _tokenMetadata)
}

// CompleteBridging is a paid mutator transaction binding the contract method 0xe4d27451.
//
// Solidity: function completeBridging(address _nativeToken, uint256 _amount, address _recipient, uint256 _chainId, bytes _tokenMetadata) returns()
func (_TokenBridge *TokenBridgeTransactorSession) CompleteBridging(_nativeToken common.Address, _amount *big.Int, _recipient common.Address, _chainId *big.Int, _tokenMetadata []byte) (*types.Transaction, error) {
	return _TokenBridge.Contract.CompleteBridging(&_TokenBridge.TransactOpts, _nativeToken, _amount, _recipient, _chainId, _tokenMetadata)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_TokenBridge *TokenBridgeTransactor) ConfirmDeployment(opts *bind.TransactOpts, _tokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "confirmDeployment", _tokens)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_TokenBridge *TokenBridgeSession) ConfirmDeployment(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.ConfirmDeployment(&_TokenBridge.TransactOpts, _tokens)
}

// ConfirmDeployment is a paid mutator transaction binding the contract method 0x4bf98dce.
//
// Solidity: function confirmDeployment(address[] _tokens) payable returns()
func (_TokenBridge *TokenBridgeTransactorSession) ConfirmDeployment(_tokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.ConfirmDeployment(&_TokenBridge.TransactOpts, _tokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0994f08.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, uint256 _sourceChainId, uint256 _targetChainId, address[] _reservedTokens) returns()
func (_TokenBridge *TokenBridgeTransactor) Initialize(opts *bind.TransactOpts, _securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _sourceChainId *big.Int, _targetChainId *big.Int, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "initialize", _securityCouncil, _messageService, _tokenBeacon, _sourceChainId, _targetChainId, _reservedTokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0994f08.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, uint256 _sourceChainId, uint256 _targetChainId, address[] _reservedTokens) returns()
func (_TokenBridge *TokenBridgeSession) Initialize(_securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _sourceChainId *big.Int, _targetChainId *big.Int, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.Initialize(&_TokenBridge.TransactOpts, _securityCouncil, _messageService, _tokenBeacon, _sourceChainId, _targetChainId, _reservedTokens)
}

// Initialize is a paid mutator transaction binding the contract method 0xb0994f08.
//
// Solidity: function initialize(address _securityCouncil, address _messageService, address _tokenBeacon, uint256 _sourceChainId, uint256 _targetChainId, address[] _reservedTokens) returns()
func (_TokenBridge *TokenBridgeTransactorSession) Initialize(_securityCouncil common.Address, _messageService common.Address, _tokenBeacon common.Address, _sourceChainId *big.Int, _targetChainId *big.Int, _reservedTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.Initialize(&_TokenBridge.TransactOpts, _securityCouncil, _messageService, _tokenBeacon, _sourceChainId, _targetChainId, _reservedTokens)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenBridge *TokenBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenBridge *TokenBridgeSession) Pause() (*types.Transaction, error) {
	return _TokenBridge.Contract.Pause(&_TokenBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TokenBridge *TokenBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _TokenBridge.Contract.Pause(&_TokenBridge.TransactOpts)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_TokenBridge *TokenBridgeTransactor) RemoveReserved(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "removeReserved", _token)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_TokenBridge *TokenBridgeSession) RemoveReserved(_token common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.RemoveReserved(&_TokenBridge.TransactOpts, _token)
}

// RemoveReserved is a paid mutator transaction binding the contract method 0xedc42a22.
//
// Solidity: function removeReserved(address _token) returns()
func (_TokenBridge *TokenBridgeTransactorSession) RemoveReserved(_token common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.RemoveReserved(&_TokenBridge.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridge *TokenBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridge *TokenBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBridge.Contract.RenounceOwnership(&_TokenBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenBridge *TokenBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenBridge.Contract.RenounceOwnership(&_TokenBridge.TransactOpts)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_TokenBridge *TokenBridgeTransactor) SetCustomContract(opts *bind.TransactOpts, _nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "setCustomContract", _nativeToken, _targetContract)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_TokenBridge *TokenBridgeSession) SetCustomContract(_nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetCustomContract(&_TokenBridge.TransactOpts, _nativeToken, _targetContract)
}

// SetCustomContract is a paid mutator transaction binding the contract method 0x1754f301.
//
// Solidity: function setCustomContract(address _nativeToken, address _targetContract) returns()
func (_TokenBridge *TokenBridgeTransactorSession) SetCustomContract(_nativeToken common.Address, _targetContract common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetCustomContract(&_TokenBridge.TransactOpts, _nativeToken, _targetContract)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_TokenBridge *TokenBridgeTransactor) SetDeployed(opts *bind.TransactOpts, _nativeTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "setDeployed", _nativeTokens)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_TokenBridge *TokenBridgeSession) SetDeployed(_nativeTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetDeployed(&_TokenBridge.TransactOpts, _nativeTokens)
}

// SetDeployed is a paid mutator transaction binding the contract method 0x2a564f34.
//
// Solidity: function setDeployed(address[] _nativeTokens) returns()
func (_TokenBridge *TokenBridgeTransactorSession) SetDeployed(_nativeTokens []common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetDeployed(&_TokenBridge.TransactOpts, _nativeTokens)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_TokenBridge *TokenBridgeTransactor) SetMessageService(opts *bind.TransactOpts, _messageService common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "setMessageService", _messageService)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_TokenBridge *TokenBridgeSession) SetMessageService(_messageService common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetMessageService(&_TokenBridge.TransactOpts, _messageService)
}

// SetMessageService is a paid mutator transaction binding the contract method 0xbe46096f.
//
// Solidity: function setMessageService(address _messageService) returns()
func (_TokenBridge *TokenBridgeTransactorSession) SetMessageService(_messageService common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetMessageService(&_TokenBridge.TransactOpts, _messageService)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_TokenBridge *TokenBridgeTransactor) SetRemoteTokenBridge(opts *bind.TransactOpts, _remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "setRemoteTokenBridge", _remoteTokenBridge)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_TokenBridge *TokenBridgeSession) SetRemoteTokenBridge(_remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetRemoteTokenBridge(&_TokenBridge.TransactOpts, _remoteTokenBridge)
}

// SetRemoteTokenBridge is a paid mutator transaction binding the contract method 0xa676e8ab.
//
// Solidity: function setRemoteTokenBridge(address _remoteTokenBridge) returns()
func (_TokenBridge *TokenBridgeTransactorSession) SetRemoteTokenBridge(_remoteTokenBridge common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetRemoteTokenBridge(&_TokenBridge.TransactOpts, _remoteTokenBridge)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_TokenBridge *TokenBridgeTransactor) SetReserved(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "setReserved", _token)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_TokenBridge *TokenBridgeSession) SetReserved(_token common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetReserved(&_TokenBridge.TransactOpts, _token)
}

// SetReserved is a paid mutator transaction binding the contract method 0xcdd914c5.
//
// Solidity: function setReserved(address _token) returns()
func (_TokenBridge *TokenBridgeTransactorSession) SetReserved(_token common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.SetReserved(&_TokenBridge.TransactOpts, _token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridge *TokenBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridge *TokenBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.TransferOwnership(&_TokenBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenBridge *TokenBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenBridge.Contract.TransferOwnership(&_TokenBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenBridge *TokenBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenBridge *TokenBridgeSession) Unpause() (*types.Transaction, error) {
	return _TokenBridge.Contract.Unpause(&_TokenBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TokenBridge *TokenBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _TokenBridge.Contract.Unpause(&_TokenBridge.TransactOpts)
}

// TokenBridgeBridgingFinalizedIterator is returned from FilterBridgingFinalized and is used to iterate over the raw logs and unpacked data for BridgingFinalized events raised by the TokenBridge contract.
type TokenBridgeBridgingFinalizedIterator struct {
	Event *TokenBridgeBridgingFinalized // Event containing the contract specifics and raw log

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
func (it *TokenBridgeBridgingFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeBridgingFinalized)
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
		it.Event = new(TokenBridgeBridgingFinalized)
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
func (it *TokenBridgeBridgingFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeBridgingFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeBridgingFinalized represents a BridgingFinalized event raised by the TokenBridge contract.
type TokenBridgeBridgingFinalized struct {
	NativeToken  common.Address
	BridgedToken common.Address
	Amount       *big.Int
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgingFinalized is a free log retrieval operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address indexed nativeToken, address indexed bridgedToken, uint256 indexed amount, address recipient)
func (_TokenBridge *TokenBridgeFilterer) FilterBridgingFinalized(opts *bind.FilterOpts, nativeToken []common.Address, bridgedToken []common.Address, amount []*big.Int) (*TokenBridgeBridgingFinalizedIterator, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "BridgingFinalized", nativeTokenRule, bridgedTokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeBridgingFinalizedIterator{contract: _TokenBridge.contract, event: "BridgingFinalized", logs: logs, sub: sub}, nil
}

// WatchBridgingFinalized is a free log subscription operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address indexed nativeToken, address indexed bridgedToken, uint256 indexed amount, address recipient)
func (_TokenBridge *TokenBridgeFilterer) WatchBridgingFinalized(opts *bind.WatchOpts, sink chan<- *TokenBridgeBridgingFinalized, nativeToken []common.Address, bridgedToken []common.Address, amount []*big.Int) (event.Subscription, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "BridgingFinalized", nativeTokenRule, bridgedTokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeBridgingFinalized)
				if err := _TokenBridge.contract.UnpackLog(event, "BridgingFinalized", log); err != nil {
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

// ParseBridgingFinalized is a log parse operation binding the contract event 0xd28a2d30314c6a2f46b657c15ee4d7ffc33b2817e78f341a260e216cebfbdbef.
//
// Solidity: event BridgingFinalized(address indexed nativeToken, address indexed bridgedToken, uint256 indexed amount, address recipient)
func (_TokenBridge *TokenBridgeFilterer) ParseBridgingFinalized(log types.Log) (*TokenBridgeBridgingFinalized, error) {
	event := new(TokenBridgeBridgingFinalized)
	if err := _TokenBridge.contract.UnpackLog(event, "BridgingFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeBridgingFinalizedV2Iterator is returned from FilterBridgingFinalizedV2 and is used to iterate over the raw logs and unpacked data for BridgingFinalizedV2 events raised by the TokenBridge contract.
type TokenBridgeBridgingFinalizedV2Iterator struct {
	Event *TokenBridgeBridgingFinalizedV2 // Event containing the contract specifics and raw log

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
func (it *TokenBridgeBridgingFinalizedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeBridgingFinalizedV2)
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
		it.Event = new(TokenBridgeBridgingFinalizedV2)
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
func (it *TokenBridgeBridgingFinalizedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeBridgingFinalizedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeBridgingFinalizedV2 represents a BridgingFinalizedV2 event raised by the TokenBridge contract.
type TokenBridgeBridgingFinalizedV2 struct {
	NativeToken  common.Address
	BridgedToken common.Address
	Amount       *big.Int
	Recipient    common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBridgingFinalizedV2 is a free log retrieval operation binding the contract event 0x6ed06519caca659cdefa71015c79a561928d3cf8cc4a3e9739fde9fb5fb38d64.
//
// Solidity: event BridgingFinalizedV2(address indexed nativeToken, address indexed bridgedToken, uint256 amount, address indexed recipient)
func (_TokenBridge *TokenBridgeFilterer) FilterBridgingFinalizedV2(opts *bind.FilterOpts, nativeToken []common.Address, bridgedToken []common.Address, recipient []common.Address) (*TokenBridgeBridgingFinalizedV2Iterator, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "BridgingFinalizedV2", nativeTokenRule, bridgedTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeBridgingFinalizedV2Iterator{contract: _TokenBridge.contract, event: "BridgingFinalizedV2", logs: logs, sub: sub}, nil
}

// WatchBridgingFinalizedV2 is a free log subscription operation binding the contract event 0x6ed06519caca659cdefa71015c79a561928d3cf8cc4a3e9739fde9fb5fb38d64.
//
// Solidity: event BridgingFinalizedV2(address indexed nativeToken, address indexed bridgedToken, uint256 amount, address indexed recipient)
func (_TokenBridge *TokenBridgeFilterer) WatchBridgingFinalizedV2(opts *bind.WatchOpts, sink chan<- *TokenBridgeBridgingFinalizedV2, nativeToken []common.Address, bridgedToken []common.Address, recipient []common.Address) (event.Subscription, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "BridgingFinalizedV2", nativeTokenRule, bridgedTokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeBridgingFinalizedV2)
				if err := _TokenBridge.contract.UnpackLog(event, "BridgingFinalizedV2", log); err != nil {
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

// ParseBridgingFinalizedV2 is a log parse operation binding the contract event 0x6ed06519caca659cdefa71015c79a561928d3cf8cc4a3e9739fde9fb5fb38d64.
//
// Solidity: event BridgingFinalizedV2(address indexed nativeToken, address indexed bridgedToken, uint256 amount, address indexed recipient)
func (_TokenBridge *TokenBridgeFilterer) ParseBridgingFinalizedV2(log types.Log) (*TokenBridgeBridgingFinalizedV2, error) {
	event := new(TokenBridgeBridgingFinalizedV2)
	if err := _TokenBridge.contract.UnpackLog(event, "BridgingFinalizedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeBridgingInitiatedIterator is returned from FilterBridgingInitiated and is used to iterate over the raw logs and unpacked data for BridgingInitiated events raised by the TokenBridge contract.
type TokenBridgeBridgingInitiatedIterator struct {
	Event *TokenBridgeBridgingInitiated // Event containing the contract specifics and raw log

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
func (it *TokenBridgeBridgingInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeBridgingInitiated)
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
		it.Event = new(TokenBridgeBridgingInitiated)
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
func (it *TokenBridgeBridgingInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeBridgingInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeBridgingInitiated represents a BridgingInitiated event raised by the TokenBridge contract.
type TokenBridgeBridgingInitiated struct {
	Sender    common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgingInitiated is a free log retrieval operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address indexed sender, address recipient, address indexed token, uint256 indexed amount)
func (_TokenBridge *TokenBridgeFilterer) FilterBridgingInitiated(opts *bind.FilterOpts, sender []common.Address, token []common.Address, amount []*big.Int) (*TokenBridgeBridgingInitiatedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "BridgingInitiated", senderRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeBridgingInitiatedIterator{contract: _TokenBridge.contract, event: "BridgingInitiated", logs: logs, sub: sub}, nil
}

// WatchBridgingInitiated is a free log subscription operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address indexed sender, address recipient, address indexed token, uint256 indexed amount)
func (_TokenBridge *TokenBridgeFilterer) WatchBridgingInitiated(opts *bind.WatchOpts, sink chan<- *TokenBridgeBridgingInitiated, sender []common.Address, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "BridgingInitiated", senderRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeBridgingInitiated)
				if err := _TokenBridge.contract.UnpackLog(event, "BridgingInitiated", log); err != nil {
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

// ParseBridgingInitiated is a log parse operation binding the contract event 0xde5fcf0a1aebed387067eb25655de732ccfc43fe5b5a3d91d367c26e773fcd1c.
//
// Solidity: event BridgingInitiated(address indexed sender, address recipient, address indexed token, uint256 indexed amount)
func (_TokenBridge *TokenBridgeFilterer) ParseBridgingInitiated(log types.Log) (*TokenBridgeBridgingInitiated, error) {
	event := new(TokenBridgeBridgingInitiated)
	if err := _TokenBridge.contract.UnpackLog(event, "BridgingInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeBridgingInitiatedV2Iterator is returned from FilterBridgingInitiatedV2 and is used to iterate over the raw logs and unpacked data for BridgingInitiatedV2 events raised by the TokenBridge contract.
type TokenBridgeBridgingInitiatedV2Iterator struct {
	Event *TokenBridgeBridgingInitiatedV2 // Event containing the contract specifics and raw log

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
func (it *TokenBridgeBridgingInitiatedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeBridgingInitiatedV2)
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
		it.Event = new(TokenBridgeBridgingInitiatedV2)
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
func (it *TokenBridgeBridgingInitiatedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeBridgingInitiatedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeBridgingInitiatedV2 represents a BridgingInitiatedV2 event raised by the TokenBridge contract.
type TokenBridgeBridgingInitiatedV2 struct {
	Sender    common.Address
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBridgingInitiatedV2 is a free log retrieval operation binding the contract event 0x8780a94875b70464f8ac6c28851501d32e7fd4ee574e4b94beb28923a3c42d9c.
//
// Solidity: event BridgingInitiatedV2(address indexed sender, address indexed recipient, address indexed token, uint256 amount)
func (_TokenBridge *TokenBridgeFilterer) FilterBridgingInitiatedV2(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, token []common.Address) (*TokenBridgeBridgingInitiatedV2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "BridgingInitiatedV2", senderRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeBridgingInitiatedV2Iterator{contract: _TokenBridge.contract, event: "BridgingInitiatedV2", logs: logs, sub: sub}, nil
}

// WatchBridgingInitiatedV2 is a free log subscription operation binding the contract event 0x8780a94875b70464f8ac6c28851501d32e7fd4ee574e4b94beb28923a3c42d9c.
//
// Solidity: event BridgingInitiatedV2(address indexed sender, address indexed recipient, address indexed token, uint256 amount)
func (_TokenBridge *TokenBridgeFilterer) WatchBridgingInitiatedV2(opts *bind.WatchOpts, sink chan<- *TokenBridgeBridgingInitiatedV2, sender []common.Address, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "BridgingInitiatedV2", senderRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeBridgingInitiatedV2)
				if err := _TokenBridge.contract.UnpackLog(event, "BridgingInitiatedV2", log); err != nil {
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

// ParseBridgingInitiatedV2 is a log parse operation binding the contract event 0x8780a94875b70464f8ac6c28851501d32e7fd4ee574e4b94beb28923a3c42d9c.
//
// Solidity: event BridgingInitiatedV2(address indexed sender, address indexed recipient, address indexed token, uint256 amount)
func (_TokenBridge *TokenBridgeFilterer) ParseBridgingInitiatedV2(log types.Log) (*TokenBridgeBridgingInitiatedV2, error) {
	event := new(TokenBridgeBridgingInitiatedV2)
	if err := _TokenBridge.contract.UnpackLog(event, "BridgingInitiatedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeCustomContractSetIterator is returned from FilterCustomContractSet and is used to iterate over the raw logs and unpacked data for CustomContractSet events raised by the TokenBridge contract.
type TokenBridgeCustomContractSetIterator struct {
	Event *TokenBridgeCustomContractSet // Event containing the contract specifics and raw log

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
func (it *TokenBridgeCustomContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeCustomContractSet)
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
		it.Event = new(TokenBridgeCustomContractSet)
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
func (it *TokenBridgeCustomContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeCustomContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeCustomContractSet represents a CustomContractSet event raised by the TokenBridge contract.
type TokenBridgeCustomContractSet struct {
	NativeToken    common.Address
	CustomContract common.Address
	SetBy          common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCustomContractSet is a free log retrieval operation binding the contract event 0x844cb5c635052898ad92bea4ece14519111765d835105e76aa1f77ad0d0aa81f.
//
// Solidity: event CustomContractSet(address indexed nativeToken, address indexed customContract, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) FilterCustomContractSet(opts *bind.FilterOpts, nativeToken []common.Address, customContract []common.Address, setBy []common.Address) (*TokenBridgeCustomContractSetIterator, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var customContractRule []interface{}
	for _, customContractItem := range customContract {
		customContractRule = append(customContractRule, customContractItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "CustomContractSet", nativeTokenRule, customContractRule, setByRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeCustomContractSetIterator{contract: _TokenBridge.contract, event: "CustomContractSet", logs: logs, sub: sub}, nil
}

// WatchCustomContractSet is a free log subscription operation binding the contract event 0x844cb5c635052898ad92bea4ece14519111765d835105e76aa1f77ad0d0aa81f.
//
// Solidity: event CustomContractSet(address indexed nativeToken, address indexed customContract, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) WatchCustomContractSet(opts *bind.WatchOpts, sink chan<- *TokenBridgeCustomContractSet, nativeToken []common.Address, customContract []common.Address, setBy []common.Address) (event.Subscription, error) {

	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}
	var customContractRule []interface{}
	for _, customContractItem := range customContract {
		customContractRule = append(customContractRule, customContractItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "CustomContractSet", nativeTokenRule, customContractRule, setByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeCustomContractSet)
				if err := _TokenBridge.contract.UnpackLog(event, "CustomContractSet", log); err != nil {
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

// ParseCustomContractSet is a log parse operation binding the contract event 0x844cb5c635052898ad92bea4ece14519111765d835105e76aa1f77ad0d0aa81f.
//
// Solidity: event CustomContractSet(address indexed nativeToken, address indexed customContract, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) ParseCustomContractSet(log types.Log) (*TokenBridgeCustomContractSet, error) {
	event := new(TokenBridgeCustomContractSet)
	if err := _TokenBridge.contract.UnpackLog(event, "CustomContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeDeploymentConfirmedIterator is returned from FilterDeploymentConfirmed and is used to iterate over the raw logs and unpacked data for DeploymentConfirmed events raised by the TokenBridge contract.
type TokenBridgeDeploymentConfirmedIterator struct {
	Event *TokenBridgeDeploymentConfirmed // Event containing the contract specifics and raw log

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
func (it *TokenBridgeDeploymentConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeDeploymentConfirmed)
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
		it.Event = new(TokenBridgeDeploymentConfirmed)
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
func (it *TokenBridgeDeploymentConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeDeploymentConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeDeploymentConfirmed represents a DeploymentConfirmed event raised by the TokenBridge contract.
type TokenBridgeDeploymentConfirmed struct {
	Tokens      []common.Address
	ConfirmedBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeploymentConfirmed is a free log retrieval operation binding the contract event 0x59eab5b5f813ac9e0c10035dfb55b5e3419eff53c0f7a869fb3c22400ea036d6.
//
// Solidity: event DeploymentConfirmed(address[] tokens, address indexed confirmedBy)
func (_TokenBridge *TokenBridgeFilterer) FilterDeploymentConfirmed(opts *bind.FilterOpts, confirmedBy []common.Address) (*TokenBridgeDeploymentConfirmedIterator, error) {

	var confirmedByRule []interface{}
	for _, confirmedByItem := range confirmedBy {
		confirmedByRule = append(confirmedByRule, confirmedByItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "DeploymentConfirmed", confirmedByRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeDeploymentConfirmedIterator{contract: _TokenBridge.contract, event: "DeploymentConfirmed", logs: logs, sub: sub}, nil
}

// WatchDeploymentConfirmed is a free log subscription operation binding the contract event 0x59eab5b5f813ac9e0c10035dfb55b5e3419eff53c0f7a869fb3c22400ea036d6.
//
// Solidity: event DeploymentConfirmed(address[] tokens, address indexed confirmedBy)
func (_TokenBridge *TokenBridgeFilterer) WatchDeploymentConfirmed(opts *bind.WatchOpts, sink chan<- *TokenBridgeDeploymentConfirmed, confirmedBy []common.Address) (event.Subscription, error) {

	var confirmedByRule []interface{}
	for _, confirmedByItem := range confirmedBy {
		confirmedByRule = append(confirmedByRule, confirmedByItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "DeploymentConfirmed", confirmedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeDeploymentConfirmed)
				if err := _TokenBridge.contract.UnpackLog(event, "DeploymentConfirmed", log); err != nil {
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

// ParseDeploymentConfirmed is a log parse operation binding the contract event 0x59eab5b5f813ac9e0c10035dfb55b5e3419eff53c0f7a869fb3c22400ea036d6.
//
// Solidity: event DeploymentConfirmed(address[] tokens, address indexed confirmedBy)
func (_TokenBridge *TokenBridgeFilterer) ParseDeploymentConfirmed(log types.Log) (*TokenBridgeDeploymentConfirmed, error) {
	event := new(TokenBridgeDeploymentConfirmed)
	if err := _TokenBridge.contract.UnpackLog(event, "DeploymentConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TokenBridge contract.
type TokenBridgeInitializedIterator struct {
	Event *TokenBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *TokenBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeInitialized)
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
		it.Event = new(TokenBridgeInitialized)
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
func (it *TokenBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeInitialized represents a Initialized event raised by the TokenBridge contract.
type TokenBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenBridge *TokenBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenBridgeInitializedIterator, error) {

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenBridgeInitializedIterator{contract: _TokenBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenBridge *TokenBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeInitialized)
				if err := _TokenBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TokenBridge *TokenBridgeFilterer) ParseInitialized(log types.Log) (*TokenBridgeInitialized, error) {
	event := new(TokenBridgeInitialized)
	if err := _TokenBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeMessageServiceUpdatedIterator is returned from FilterMessageServiceUpdated and is used to iterate over the raw logs and unpacked data for MessageServiceUpdated events raised by the TokenBridge contract.
type TokenBridgeMessageServiceUpdatedIterator struct {
	Event *TokenBridgeMessageServiceUpdated // Event containing the contract specifics and raw log

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
func (it *TokenBridgeMessageServiceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeMessageServiceUpdated)
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
		it.Event = new(TokenBridgeMessageServiceUpdated)
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
func (it *TokenBridgeMessageServiceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeMessageServiceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeMessageServiceUpdated represents a MessageServiceUpdated event raised by the TokenBridge contract.
type TokenBridgeMessageServiceUpdated struct {
	NewMessageService common.Address
	OldMessageService common.Address
	SetBy             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMessageServiceUpdated is a free log retrieval operation binding the contract event 0xc96d462e42a71473da49a1d58c1754b9b2d319786692d621dc7f921331c517e9.
//
// Solidity: event MessageServiceUpdated(address indexed newMessageService, address indexed oldMessageService, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) FilterMessageServiceUpdated(opts *bind.FilterOpts, newMessageService []common.Address, oldMessageService []common.Address, setBy []common.Address) (*TokenBridgeMessageServiceUpdatedIterator, error) {

	var newMessageServiceRule []interface{}
	for _, newMessageServiceItem := range newMessageService {
		newMessageServiceRule = append(newMessageServiceRule, newMessageServiceItem)
	}
	var oldMessageServiceRule []interface{}
	for _, oldMessageServiceItem := range oldMessageService {
		oldMessageServiceRule = append(oldMessageServiceRule, oldMessageServiceItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "MessageServiceUpdated", newMessageServiceRule, oldMessageServiceRule, setByRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeMessageServiceUpdatedIterator{contract: _TokenBridge.contract, event: "MessageServiceUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageServiceUpdated is a free log subscription operation binding the contract event 0xc96d462e42a71473da49a1d58c1754b9b2d319786692d621dc7f921331c517e9.
//
// Solidity: event MessageServiceUpdated(address indexed newMessageService, address indexed oldMessageService, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) WatchMessageServiceUpdated(opts *bind.WatchOpts, sink chan<- *TokenBridgeMessageServiceUpdated, newMessageService []common.Address, oldMessageService []common.Address, setBy []common.Address) (event.Subscription, error) {

	var newMessageServiceRule []interface{}
	for _, newMessageServiceItem := range newMessageService {
		newMessageServiceRule = append(newMessageServiceRule, newMessageServiceItem)
	}
	var oldMessageServiceRule []interface{}
	for _, oldMessageServiceItem := range oldMessageService {
		oldMessageServiceRule = append(oldMessageServiceRule, oldMessageServiceItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "MessageServiceUpdated", newMessageServiceRule, oldMessageServiceRule, setByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeMessageServiceUpdated)
				if err := _TokenBridge.contract.UnpackLog(event, "MessageServiceUpdated", log); err != nil {
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

// ParseMessageServiceUpdated is a log parse operation binding the contract event 0xc96d462e42a71473da49a1d58c1754b9b2d319786692d621dc7f921331c517e9.
//
// Solidity: event MessageServiceUpdated(address indexed newMessageService, address indexed oldMessageService, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) ParseMessageServiceUpdated(log types.Log) (*TokenBridgeMessageServiceUpdated, error) {
	event := new(TokenBridgeMessageServiceUpdated)
	if err := _TokenBridge.contract.UnpackLog(event, "MessageServiceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeNewTokenIterator is returned from FilterNewToken and is used to iterate over the raw logs and unpacked data for NewToken events raised by the TokenBridge contract.
type TokenBridgeNewTokenIterator struct {
	Event *TokenBridgeNewToken // Event containing the contract specifics and raw log

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
func (it *TokenBridgeNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeNewToken)
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
		it.Event = new(TokenBridgeNewToken)
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
func (it *TokenBridgeNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeNewToken represents a NewToken event raised by the TokenBridge contract.
type TokenBridgeNewToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewToken is a free log retrieval operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) FilterNewToken(opts *bind.FilterOpts, token []common.Address) (*TokenBridgeNewTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "NewToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeNewTokenIterator{contract: _TokenBridge.contract, event: "NewToken", logs: logs, sub: sub}, nil
}

// WatchNewToken is a free log subscription operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) WatchNewToken(opts *bind.WatchOpts, sink chan<- *TokenBridgeNewToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "NewToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeNewToken)
				if err := _TokenBridge.contract.UnpackLog(event, "NewToken", log); err != nil {
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

// ParseNewToken is a log parse operation binding the contract event 0x0f53e2a811b6fd2d6cd965fd6c27b44fb924ca39f7a7f321115705c22366d623.
//
// Solidity: event NewToken(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) ParseNewToken(log types.Log) (*TokenBridgeNewToken, error) {
	event := new(TokenBridgeNewToken)
	if err := _TokenBridge.contract.UnpackLog(event, "NewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeNewTokenDeployedIterator is returned from FilterNewTokenDeployed and is used to iterate over the raw logs and unpacked data for NewTokenDeployed events raised by the TokenBridge contract.
type TokenBridgeNewTokenDeployedIterator struct {
	Event *TokenBridgeNewTokenDeployed // Event containing the contract specifics and raw log

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
func (it *TokenBridgeNewTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeNewTokenDeployed)
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
		it.Event = new(TokenBridgeNewTokenDeployed)
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
func (it *TokenBridgeNewTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeNewTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeNewTokenDeployed represents a NewTokenDeployed event raised by the TokenBridge contract.
type TokenBridgeNewTokenDeployed struct {
	BridgedToken common.Address
	NativeToken  common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewTokenDeployed is a free log retrieval operation binding the contract event 0xd5d4920bb61e6141c8499d50a7bd617dae2b1818c9d6b995d3f2ba4975e32ea4.
//
// Solidity: event NewTokenDeployed(address indexed bridgedToken, address indexed nativeToken)
func (_TokenBridge *TokenBridgeFilterer) FilterNewTokenDeployed(opts *bind.FilterOpts, bridgedToken []common.Address, nativeToken []common.Address) (*TokenBridgeNewTokenDeployedIterator, error) {

	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}
	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "NewTokenDeployed", bridgedTokenRule, nativeTokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeNewTokenDeployedIterator{contract: _TokenBridge.contract, event: "NewTokenDeployed", logs: logs, sub: sub}, nil
}

// WatchNewTokenDeployed is a free log subscription operation binding the contract event 0xd5d4920bb61e6141c8499d50a7bd617dae2b1818c9d6b995d3f2ba4975e32ea4.
//
// Solidity: event NewTokenDeployed(address indexed bridgedToken, address indexed nativeToken)
func (_TokenBridge *TokenBridgeFilterer) WatchNewTokenDeployed(opts *bind.WatchOpts, sink chan<- *TokenBridgeNewTokenDeployed, bridgedToken []common.Address, nativeToken []common.Address) (event.Subscription, error) {

	var bridgedTokenRule []interface{}
	for _, bridgedTokenItem := range bridgedToken {
		bridgedTokenRule = append(bridgedTokenRule, bridgedTokenItem)
	}
	var nativeTokenRule []interface{}
	for _, nativeTokenItem := range nativeToken {
		nativeTokenRule = append(nativeTokenRule, nativeTokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "NewTokenDeployed", bridgedTokenRule, nativeTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeNewTokenDeployed)
				if err := _TokenBridge.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
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

// ParseNewTokenDeployed is a log parse operation binding the contract event 0xd5d4920bb61e6141c8499d50a7bd617dae2b1818c9d6b995d3f2ba4975e32ea4.
//
// Solidity: event NewTokenDeployed(address indexed bridgedToken, address indexed nativeToken)
func (_TokenBridge *TokenBridgeFilterer) ParseNewTokenDeployed(log types.Log) (*TokenBridgeNewTokenDeployed, error) {
	event := new(TokenBridgeNewTokenDeployed)
	if err := _TokenBridge.contract.UnpackLog(event, "NewTokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the TokenBridge contract.
type TokenBridgeOwnershipTransferStartedIterator struct {
	Event *TokenBridgeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *TokenBridgeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeOwnershipTransferStarted)
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
		it.Event = new(TokenBridgeOwnershipTransferStarted)
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
func (it *TokenBridgeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the TokenBridge contract.
type TokenBridgeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenBridge *TokenBridgeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenBridgeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeOwnershipTransferStartedIterator{contract: _TokenBridge.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenBridge *TokenBridgeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *TokenBridgeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeOwnershipTransferStarted)
				if err := _TokenBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenBridge *TokenBridgeFilterer) ParseOwnershipTransferStarted(log types.Log) (*TokenBridgeOwnershipTransferStarted, error) {
	event := new(TokenBridgeOwnershipTransferStarted)
	if err := _TokenBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenBridge contract.
type TokenBridgeOwnershipTransferredIterator struct {
	Event *TokenBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeOwnershipTransferred)
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
		it.Event = new(TokenBridgeOwnershipTransferred)
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
func (it *TokenBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the TokenBridge contract.
type TokenBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBridge *TokenBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeOwnershipTransferredIterator{contract: _TokenBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenBridge *TokenBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeOwnershipTransferred)
				if err := _TokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenBridge *TokenBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*TokenBridgeOwnershipTransferred, error) {
	event := new(TokenBridgeOwnershipTransferred)
	if err := _TokenBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TokenBridge contract.
type TokenBridgePausedIterator struct {
	Event *TokenBridgePaused // Event containing the contract specifics and raw log

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
func (it *TokenBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgePaused)
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
		it.Event = new(TokenBridgePaused)
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
func (it *TokenBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgePaused represents a Paused event raised by the TokenBridge contract.
type TokenBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenBridge *TokenBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*TokenBridgePausedIterator, error) {

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TokenBridgePausedIterator{contract: _TokenBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenBridge *TokenBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TokenBridgePaused) (event.Subscription, error) {

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgePaused)
				if err := _TokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TokenBridge *TokenBridgeFilterer) ParsePaused(log types.Log) (*TokenBridgePaused, error) {
	event := new(TokenBridgePaused)
	if err := _TokenBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeRemoteTokenBridgeSetIterator is returned from FilterRemoteTokenBridgeSet and is used to iterate over the raw logs and unpacked data for RemoteTokenBridgeSet events raised by the TokenBridge contract.
type TokenBridgeRemoteTokenBridgeSetIterator struct {
	Event *TokenBridgeRemoteTokenBridgeSet // Event containing the contract specifics and raw log

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
func (it *TokenBridgeRemoteTokenBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeRemoteTokenBridgeSet)
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
		it.Event = new(TokenBridgeRemoteTokenBridgeSet)
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
func (it *TokenBridgeRemoteTokenBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeRemoteTokenBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeRemoteTokenBridgeSet represents a RemoteTokenBridgeSet event raised by the TokenBridge contract.
type TokenBridgeRemoteTokenBridgeSet struct {
	RemoteTokenBridge common.Address
	SetBy             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRemoteTokenBridgeSet is a free log retrieval operation binding the contract event 0xb044c1a1a05a729c402def784b4e4cb01612ff03eee6f0beb3eba0f0606260a1.
//
// Solidity: event RemoteTokenBridgeSet(address indexed remoteTokenBridge, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) FilterRemoteTokenBridgeSet(opts *bind.FilterOpts, remoteTokenBridge []common.Address, setBy []common.Address) (*TokenBridgeRemoteTokenBridgeSetIterator, error) {

	var remoteTokenBridgeRule []interface{}
	for _, remoteTokenBridgeItem := range remoteTokenBridge {
		remoteTokenBridgeRule = append(remoteTokenBridgeRule, remoteTokenBridgeItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "RemoteTokenBridgeSet", remoteTokenBridgeRule, setByRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeRemoteTokenBridgeSetIterator{contract: _TokenBridge.contract, event: "RemoteTokenBridgeSet", logs: logs, sub: sub}, nil
}

// WatchRemoteTokenBridgeSet is a free log subscription operation binding the contract event 0xb044c1a1a05a729c402def784b4e4cb01612ff03eee6f0beb3eba0f0606260a1.
//
// Solidity: event RemoteTokenBridgeSet(address indexed remoteTokenBridge, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) WatchRemoteTokenBridgeSet(opts *bind.WatchOpts, sink chan<- *TokenBridgeRemoteTokenBridgeSet, remoteTokenBridge []common.Address, setBy []common.Address) (event.Subscription, error) {

	var remoteTokenBridgeRule []interface{}
	for _, remoteTokenBridgeItem := range remoteTokenBridge {
		remoteTokenBridgeRule = append(remoteTokenBridgeRule, remoteTokenBridgeItem)
	}
	var setByRule []interface{}
	for _, setByItem := range setBy {
		setByRule = append(setByRule, setByItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "RemoteTokenBridgeSet", remoteTokenBridgeRule, setByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeRemoteTokenBridgeSet)
				if err := _TokenBridge.contract.UnpackLog(event, "RemoteTokenBridgeSet", log); err != nil {
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

// ParseRemoteTokenBridgeSet is a log parse operation binding the contract event 0xb044c1a1a05a729c402def784b4e4cb01612ff03eee6f0beb3eba0f0606260a1.
//
// Solidity: event RemoteTokenBridgeSet(address indexed remoteTokenBridge, address indexed setBy)
func (_TokenBridge *TokenBridgeFilterer) ParseRemoteTokenBridgeSet(log types.Log) (*TokenBridgeRemoteTokenBridgeSet, error) {
	event := new(TokenBridgeRemoteTokenBridgeSet)
	if err := _TokenBridge.contract.UnpackLog(event, "RemoteTokenBridgeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeReservationRemovedIterator is returned from FilterReservationRemoved and is used to iterate over the raw logs and unpacked data for ReservationRemoved events raised by the TokenBridge contract.
type TokenBridgeReservationRemovedIterator struct {
	Event *TokenBridgeReservationRemoved // Event containing the contract specifics and raw log

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
func (it *TokenBridgeReservationRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeReservationRemoved)
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
		it.Event = new(TokenBridgeReservationRemoved)
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
func (it *TokenBridgeReservationRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeReservationRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeReservationRemoved represents a ReservationRemoved event raised by the TokenBridge contract.
type TokenBridgeReservationRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReservationRemoved is a free log retrieval operation binding the contract event 0x0145163d8d460d1ab21463758d147fdfe79d4b57c81ca3d1439996104ae68959.
//
// Solidity: event ReservationRemoved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) FilterReservationRemoved(opts *bind.FilterOpts, token []common.Address) (*TokenBridgeReservationRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "ReservationRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeReservationRemovedIterator{contract: _TokenBridge.contract, event: "ReservationRemoved", logs: logs, sub: sub}, nil
}

// WatchReservationRemoved is a free log subscription operation binding the contract event 0x0145163d8d460d1ab21463758d147fdfe79d4b57c81ca3d1439996104ae68959.
//
// Solidity: event ReservationRemoved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) WatchReservationRemoved(opts *bind.WatchOpts, sink chan<- *TokenBridgeReservationRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "ReservationRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeReservationRemoved)
				if err := _TokenBridge.contract.UnpackLog(event, "ReservationRemoved", log); err != nil {
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

// ParseReservationRemoved is a log parse operation binding the contract event 0x0145163d8d460d1ab21463758d147fdfe79d4b57c81ca3d1439996104ae68959.
//
// Solidity: event ReservationRemoved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) ParseReservationRemoved(log types.Log) (*TokenBridgeReservationRemoved, error) {
	event := new(TokenBridgeReservationRemoved)
	if err := _TokenBridge.contract.UnpackLog(event, "ReservationRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeTokenDeployedIterator is returned from FilterTokenDeployed and is used to iterate over the raw logs and unpacked data for TokenDeployed events raised by the TokenBridge contract.
type TokenBridgeTokenDeployedIterator struct {
	Event *TokenBridgeTokenDeployed // Event containing the contract specifics and raw log

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
func (it *TokenBridgeTokenDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeTokenDeployed)
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
		it.Event = new(TokenBridgeTokenDeployed)
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
func (it *TokenBridgeTokenDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeTokenDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeTokenDeployed represents a TokenDeployed event raised by the TokenBridge contract.
type TokenBridgeTokenDeployed struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenDeployed is a free log retrieval operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) FilterTokenDeployed(opts *bind.FilterOpts, token []common.Address) (*TokenBridgeTokenDeployedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "TokenDeployed", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeTokenDeployedIterator{contract: _TokenBridge.contract, event: "TokenDeployed", logs: logs, sub: sub}, nil
}

// WatchTokenDeployed is a free log subscription operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) WatchTokenDeployed(opts *bind.WatchOpts, sink chan<- *TokenBridgeTokenDeployed, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "TokenDeployed", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeTokenDeployed)
				if err := _TokenBridge.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
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

// ParseTokenDeployed is a log parse operation binding the contract event 0x91d24864a084ab70b268a1f865e757ca12006cf298d763b6be697302ef86498c.
//
// Solidity: event TokenDeployed(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) ParseTokenDeployed(log types.Log) (*TokenBridgeTokenDeployed, error) {
	event := new(TokenBridgeTokenDeployed)
	if err := _TokenBridge.contract.UnpackLog(event, "TokenDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeTokenReservedIterator is returned from FilterTokenReserved and is used to iterate over the raw logs and unpacked data for TokenReserved events raised by the TokenBridge contract.
type TokenBridgeTokenReservedIterator struct {
	Event *TokenBridgeTokenReserved // Event containing the contract specifics and raw log

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
func (it *TokenBridgeTokenReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeTokenReserved)
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
		it.Event = new(TokenBridgeTokenReserved)
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
func (it *TokenBridgeTokenReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeTokenReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeTokenReserved represents a TokenReserved event raised by the TokenBridge contract.
type TokenBridgeTokenReserved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenReserved is a free log retrieval operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) FilterTokenReserved(opts *bind.FilterOpts, token []common.Address) (*TokenBridgeTokenReservedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "TokenReserved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenBridgeTokenReservedIterator{contract: _TokenBridge.contract, event: "TokenReserved", logs: logs, sub: sub}, nil
}

// WatchTokenReserved is a free log subscription operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) WatchTokenReserved(opts *bind.WatchOpts, sink chan<- *TokenBridgeTokenReserved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "TokenReserved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeTokenReserved)
				if err := _TokenBridge.contract.UnpackLog(event, "TokenReserved", log); err != nil {
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

// ParseTokenReserved is a log parse operation binding the contract event 0x5e023c7a09fa0534ce3199f65fc3e635a5e851c5adc88ebda3b9d332ae07cbe9.
//
// Solidity: event TokenReserved(address indexed token)
func (_TokenBridge *TokenBridgeFilterer) ParseTokenReserved(log types.Log) (*TokenBridgeTokenReserved, error) {
	event := new(TokenBridgeTokenReserved)
	if err := _TokenBridge.contract.UnpackLog(event, "TokenReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TokenBridge contract.
type TokenBridgeUnpausedIterator struct {
	Event *TokenBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *TokenBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenBridgeUnpaused)
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
		it.Event = new(TokenBridgeUnpaused)
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
func (it *TokenBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenBridgeUnpaused represents a Unpaused event raised by the TokenBridge contract.
type TokenBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenBridge *TokenBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TokenBridgeUnpausedIterator, error) {

	logs, sub, err := _TokenBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TokenBridgeUnpausedIterator{contract: _TokenBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenBridge *TokenBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TokenBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _TokenBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenBridgeUnpaused)
				if err := _TokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TokenBridge *TokenBridgeFilterer) ParseUnpaused(log types.Log) (*TokenBridgeUnpaused, error) {
	event := new(TokenBridgeUnpaused)
	if err := _TokenBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
