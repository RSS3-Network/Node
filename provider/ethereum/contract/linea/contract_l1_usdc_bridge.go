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

// L1USDCBridgeMetaData contains all meta data concerning the L1USDCBridge contract.
var L1USDCBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"AmountTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"NoBurnCapabilities\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remoteUSDCBridge\",\"type\":\"address\"}],\"name\":\"NotFromRemoteUSDCBridge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"messageService\",\"type\":\"address\"}],\"name\":\"NotMessageService\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"remoteUSDCBridge\",\"type\":\"address\"}],\"name\":\"RemoteUSDCBridgeAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteUSDCBridgeNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"messageService\",\"type\":\"address\"}],\"name\":\"SameMessageServiceAddr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"SenderBalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ZeroAmountNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"MessageServiceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedFromOtherLayer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRemoteUSDCBridge\",\"type\":\"address\"}],\"name\":\"RemoteUSDCBridgeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMessageService\",\"name\":\"newMessageService\",\"type\":\"address\"}],\"name\":\"changeMessageService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMessageService\",\"name\":\"_messageService\",\"type\":\"address\"},{\"internalType\":\"contractIUSDC\",\"name\":\"_usdc\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageService\",\"outputs\":[{\"internalType\":\"contractIMessageService\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"receiveFromOtherLayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteUSDCBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_remoteUSDCBridge\",\"type\":\"address\"}],\"name\":\"setRemoteUSDCBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIUSDC\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1USDCBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use L1USDCBridgeMetaData.ABI instead.
var L1USDCBridgeABI = L1USDCBridgeMetaData.ABI

// L1USDCBridge is an auto generated Go binding around an Ethereum contract.
type L1USDCBridge struct {
	L1USDCBridgeCaller     // Read-only binding to the contract
	L1USDCBridgeTransactor // Write-only binding to the contract
	L1USDCBridgeFilterer   // Log filterer for contract events
}

// L1USDCBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1USDCBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1USDCBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1USDCBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1USDCBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1USDCBridgeSession struct {
	Contract     *L1USDCBridge     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1USDCBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1USDCBridgeCallerSession struct {
	Contract *L1USDCBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L1USDCBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1USDCBridgeTransactorSession struct {
	Contract     *L1USDCBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L1USDCBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1USDCBridgeRaw struct {
	Contract *L1USDCBridge // Generic contract binding to access the raw methods on
}

// L1USDCBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1USDCBridgeCallerRaw struct {
	Contract *L1USDCBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// L1USDCBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1USDCBridgeTransactorRaw struct {
	Contract *L1USDCBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1USDCBridge creates a new instance of L1USDCBridge, bound to a specific deployed contract.
func NewL1USDCBridge(address common.Address, backend bind.ContractBackend) (*L1USDCBridge, error) {
	contract, err := bindL1USDCBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridge{L1USDCBridgeCaller: L1USDCBridgeCaller{contract: contract}, L1USDCBridgeTransactor: L1USDCBridgeTransactor{contract: contract}, L1USDCBridgeFilterer: L1USDCBridgeFilterer{contract: contract}}, nil
}

// NewL1USDCBridgeCaller creates a new read-only instance of L1USDCBridge, bound to a specific deployed contract.
func NewL1USDCBridgeCaller(address common.Address, caller bind.ContractCaller) (*L1USDCBridgeCaller, error) {
	contract, err := bindL1USDCBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeCaller{contract: contract}, nil
}

// NewL1USDCBridgeTransactor creates a new write-only instance of L1USDCBridge, bound to a specific deployed contract.
func NewL1USDCBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*L1USDCBridgeTransactor, error) {
	contract, err := bindL1USDCBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeTransactor{contract: contract}, nil
}

// NewL1USDCBridgeFilterer creates a new log filterer instance of L1USDCBridge, bound to a specific deployed contract.
func NewL1USDCBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*L1USDCBridgeFilterer, error) {
	contract, err := bindL1USDCBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeFilterer{contract: contract}, nil
}

// bindL1USDCBridge binds a generic wrapper to an already deployed contract.
func bindL1USDCBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1USDCBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1USDCBridge *L1USDCBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1USDCBridge.Contract.L1USDCBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1USDCBridge *L1USDCBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.L1USDCBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1USDCBridge *L1USDCBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.L1USDCBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1USDCBridge *L1USDCBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1USDCBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1USDCBridge *L1USDCBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1USDCBridge *L1USDCBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_L1USDCBridge *L1USDCBridgeCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_L1USDCBridge *L1USDCBridgeSession) Balance() (*big.Int, error) {
	return _L1USDCBridge.Contract.Balance(&_L1USDCBridge.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_L1USDCBridge *L1USDCBridgeCallerSession) Balance() (*big.Int, error) {
	return _L1USDCBridge.Contract.Balance(&_L1USDCBridge.CallOpts)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCaller) MessageService(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "messageService")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_L1USDCBridge *L1USDCBridgeSession) MessageService() (common.Address, error) {
	return _L1USDCBridge.Contract.MessageService(&_L1USDCBridge.CallOpts)
}

// MessageService is a free data retrieval call binding the contract method 0x8dae45dd.
//
// Solidity: function messageService() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCallerSession) MessageService() (common.Address, error) {
	return _L1USDCBridge.Contract.MessageService(&_L1USDCBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeSession) Owner() (common.Address, error) {
	return _L1USDCBridge.Contract.Owner(&_L1USDCBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCallerSession) Owner() (common.Address, error) {
	return _L1USDCBridge.Contract.Owner(&_L1USDCBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1USDCBridge *L1USDCBridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1USDCBridge *L1USDCBridgeSession) Paused() (bool, error) {
	return _L1USDCBridge.Contract.Paused(&_L1USDCBridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_L1USDCBridge *L1USDCBridgeCallerSession) Paused() (bool, error) {
	return _L1USDCBridge.Contract.Paused(&_L1USDCBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeSession) PendingOwner() (common.Address, error) {
	return _L1USDCBridge.Contract.PendingOwner(&_L1USDCBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCallerSession) PendingOwner() (common.Address, error) {
	return _L1USDCBridge.Contract.PendingOwner(&_L1USDCBridge.CallOpts)
}

// RemoteUSDCBridge is a free data retrieval call binding the contract method 0x66e6a855.
//
// Solidity: function remoteUSDCBridge() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCaller) RemoteUSDCBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "remoteUSDCBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteUSDCBridge is a free data retrieval call binding the contract method 0x66e6a855.
//
// Solidity: function remoteUSDCBridge() view returns(address)
func (_L1USDCBridge *L1USDCBridgeSession) RemoteUSDCBridge() (common.Address, error) {
	return _L1USDCBridge.Contract.RemoteUSDCBridge(&_L1USDCBridge.CallOpts)
}

// RemoteUSDCBridge is a free data retrieval call binding the contract method 0x66e6a855.
//
// Solidity: function remoteUSDCBridge() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCallerSession) RemoteUSDCBridge() (common.Address, error) {
	return _L1USDCBridge.Contract.RemoteUSDCBridge(&_L1USDCBridge.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1USDCBridge.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_L1USDCBridge *L1USDCBridgeSession) Usdc() (common.Address, error) {
	return _L1USDCBridge.Contract.Usdc(&_L1USDCBridge.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_L1USDCBridge *L1USDCBridgeCallerSession) Usdc() (common.Address, error) {
	return _L1USDCBridge.Contract.Usdc(&_L1USDCBridge.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeSession) AcceptOwnership() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.AcceptOwnership(&_L1USDCBridge.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.AcceptOwnership(&_L1USDCBridge.TransactOpts)
}

// ChangeMessageService is a paid mutator transaction binding the contract method 0x151fcebe.
//
// Solidity: function changeMessageService(address newMessageService) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) ChangeMessageService(opts *bind.TransactOpts, newMessageService common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "changeMessageService", newMessageService)
}

// ChangeMessageService is a paid mutator transaction binding the contract method 0x151fcebe.
//
// Solidity: function changeMessageService(address newMessageService) returns()
func (_L1USDCBridge *L1USDCBridgeSession) ChangeMessageService(newMessageService common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.ChangeMessageService(&_L1USDCBridge.TransactOpts, newMessageService)
}

// ChangeMessageService is a paid mutator transaction binding the contract method 0x151fcebe.
//
// Solidity: function changeMessageService(address newMessageService) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) ChangeMessageService(newMessageService common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.ChangeMessageService(&_L1USDCBridge.TransactOpts, newMessageService)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_L1USDCBridge *L1USDCBridgeSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Deposit(&_L1USDCBridge.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) payable returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Deposit(&_L1USDCBridge.TransactOpts, amount)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) payable returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) DepositTo(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "depositTo", amount, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) payable returns()
func (_L1USDCBridge *L1USDCBridgeSession) DepositTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.DepositTo(&_L1USDCBridge.TransactOpts, amount, to)
}

// DepositTo is a paid mutator transaction binding the contract method 0x70aff70f.
//
// Solidity: function depositTo(uint256 amount, address to) payable returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) DepositTo(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.DepositTo(&_L1USDCBridge.TransactOpts, amount, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageService, address _usdc) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) Initialize(opts *bind.TransactOpts, _messageService common.Address, _usdc common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "initialize", _messageService, _usdc)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageService, address _usdc) returns()
func (_L1USDCBridge *L1USDCBridgeSession) Initialize(_messageService common.Address, _usdc common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Initialize(&_L1USDCBridge.TransactOpts, _messageService, _usdc)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageService, address _usdc) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) Initialize(_messageService common.Address, _usdc common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Initialize(&_L1USDCBridge.TransactOpts, _messageService, _usdc)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1USDCBridge *L1USDCBridgeSession) Pause() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Pause(&_L1USDCBridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Pause(&_L1USDCBridge.TransactOpts)
}

// ReceiveFromOtherLayer is a paid mutator transaction binding the contract method 0x26dfbc20.
//
// Solidity: function receiveFromOtherLayer(address recipient, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) ReceiveFromOtherLayer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "receiveFromOtherLayer", recipient, amount)
}

// ReceiveFromOtherLayer is a paid mutator transaction binding the contract method 0x26dfbc20.
//
// Solidity: function receiveFromOtherLayer(address recipient, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeSession) ReceiveFromOtherLayer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.ReceiveFromOtherLayer(&_L1USDCBridge.TransactOpts, recipient, amount)
}

// ReceiveFromOtherLayer is a paid mutator transaction binding the contract method 0x26dfbc20.
//
// Solidity: function receiveFromOtherLayer(address recipient, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) ReceiveFromOtherLayer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.ReceiveFromOtherLayer(&_L1USDCBridge.TransactOpts, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.RenounceOwnership(&_L1USDCBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.RenounceOwnership(&_L1USDCBridge.TransactOpts)
}

// RescueTransfer is a paid mutator transaction binding the contract method 0xa7052cab.
//
// Solidity: function rescueTransfer(address to, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) RescueTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "rescueTransfer", to, amount)
}

// RescueTransfer is a paid mutator transaction binding the contract method 0xa7052cab.
//
// Solidity: function rescueTransfer(address to, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeSession) RescueTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.RescueTransfer(&_L1USDCBridge.TransactOpts, to, amount)
}

// RescueTransfer is a paid mutator transaction binding the contract method 0xa7052cab.
//
// Solidity: function rescueTransfer(address to, uint256 amount) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) RescueTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.RescueTransfer(&_L1USDCBridge.TransactOpts, to, amount)
}

// SetRemoteUSDCBridge is a paid mutator transaction binding the contract method 0x278e04cc.
//
// Solidity: function setRemoteUSDCBridge(address _remoteUSDCBridge) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) SetRemoteUSDCBridge(opts *bind.TransactOpts, _remoteUSDCBridge common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "setRemoteUSDCBridge", _remoteUSDCBridge)
}

// SetRemoteUSDCBridge is a paid mutator transaction binding the contract method 0x278e04cc.
//
// Solidity: function setRemoteUSDCBridge(address _remoteUSDCBridge) returns()
func (_L1USDCBridge *L1USDCBridgeSession) SetRemoteUSDCBridge(_remoteUSDCBridge common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.SetRemoteUSDCBridge(&_L1USDCBridge.TransactOpts, _remoteUSDCBridge)
}

// SetRemoteUSDCBridge is a paid mutator transaction binding the contract method 0x278e04cc.
//
// Solidity: function setRemoteUSDCBridge(address _remoteUSDCBridge) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) SetRemoteUSDCBridge(_remoteUSDCBridge common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.SetRemoteUSDCBridge(&_L1USDCBridge.TransactOpts, _remoteUSDCBridge)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCBridge *L1USDCBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.TransferOwnership(&_L1USDCBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1USDCBridge.Contract.TransferOwnership(&_L1USDCBridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1USDCBridge *L1USDCBridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1USDCBridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1USDCBridge *L1USDCBridgeSession) Unpause() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Unpause(&_L1USDCBridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_L1USDCBridge *L1USDCBridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _L1USDCBridge.Contract.Unpause(&_L1USDCBridge.TransactOpts)
}

// L1USDCBridgeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the L1USDCBridge contract.
type L1USDCBridgeDepositedIterator struct {
	Event *L1USDCBridgeDeposited // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeDeposited)
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
		it.Event = new(L1USDCBridgeDeposited)
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
func (it *L1USDCBridgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeDeposited represents a Deposited event raised by the L1USDCBridge contract.
type L1USDCBridgeDeposited struct {
	Depositor common.Address
	Amount    *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0xb4e1304f97b5093610f51b33ddab6622388422e2dac138b0d32f93dcfbd39edf.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, address indexed to)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterDeposited(opts *bind.FilterOpts, depositor []common.Address, to []common.Address) (*L1USDCBridgeDepositedIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "Deposited", depositorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeDepositedIterator{contract: _L1USDCBridge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0xb4e1304f97b5093610f51b33ddab6622388422e2dac138b0d32f93dcfbd39edf.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, address indexed to)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeDeposited, depositor []common.Address, to []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "Deposited", depositorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeDeposited)
				if err := _L1USDCBridge.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0xb4e1304f97b5093610f51b33ddab6622388422e2dac138b0d32f93dcfbd39edf.
//
// Solidity: event Deposited(address indexed depositor, uint256 amount, address indexed to)
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseDeposited(log types.Log) (*L1USDCBridgeDeposited, error) {
	event := new(L1USDCBridgeDeposited)
	if err := _L1USDCBridge.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L1USDCBridge contract.
type L1USDCBridgeInitializedIterator struct {
	Event *L1USDCBridgeInitialized // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeInitialized)
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
		it.Event = new(L1USDCBridgeInitialized)
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
func (it *L1USDCBridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeInitialized represents a Initialized event raised by the L1USDCBridge contract.
type L1USDCBridgeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*L1USDCBridgeInitializedIterator, error) {

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeInitializedIterator{contract: _L1USDCBridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeInitialized)
				if err := _L1USDCBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseInitialized(log types.Log) (*L1USDCBridgeInitialized, error) {
	event := new(L1USDCBridgeInitialized)
	if err := _L1USDCBridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeMessageServiceUpdatedIterator is returned from FilterMessageServiceUpdated and is used to iterate over the raw logs and unpacked data for MessageServiceUpdated events raised by the L1USDCBridge contract.
type L1USDCBridgeMessageServiceUpdatedIterator struct {
	Event *L1USDCBridgeMessageServiceUpdated // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeMessageServiceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeMessageServiceUpdated)
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
		it.Event = new(L1USDCBridgeMessageServiceUpdated)
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
func (it *L1USDCBridgeMessageServiceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeMessageServiceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeMessageServiceUpdated represents a MessageServiceUpdated event raised by the L1USDCBridge contract.
type L1USDCBridgeMessageServiceUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageServiceUpdated is a free log retrieval operation binding the contract event 0x0d6365bdd6ffd3d8fb83691ea1eb4947dd043cceca1a957681a6d2b6050816e2.
//
// Solidity: event MessageServiceUpdated(address indexed oldAddress, address indexed newAddress)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterMessageServiceUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*L1USDCBridgeMessageServiceUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "MessageServiceUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeMessageServiceUpdatedIterator{contract: _L1USDCBridge.contract, event: "MessageServiceUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageServiceUpdated is a free log subscription operation binding the contract event 0x0d6365bdd6ffd3d8fb83691ea1eb4947dd043cceca1a957681a6d2b6050816e2.
//
// Solidity: event MessageServiceUpdated(address indexed oldAddress, address indexed newAddress)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchMessageServiceUpdated(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeMessageServiceUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "MessageServiceUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeMessageServiceUpdated)
				if err := _L1USDCBridge.contract.UnpackLog(event, "MessageServiceUpdated", log); err != nil {
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

// ParseMessageServiceUpdated is a log parse operation binding the contract event 0x0d6365bdd6ffd3d8fb83691ea1eb4947dd043cceca1a957681a6d2b6050816e2.
//
// Solidity: event MessageServiceUpdated(address indexed oldAddress, address indexed newAddress)
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseMessageServiceUpdated(log types.Log) (*L1USDCBridgeMessageServiceUpdated, error) {
	event := new(L1USDCBridgeMessageServiceUpdated)
	if err := _L1USDCBridge.contract.UnpackLog(event, "MessageServiceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the L1USDCBridge contract.
type L1USDCBridgeOwnershipTransferStartedIterator struct {
	Event *L1USDCBridgeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeOwnershipTransferStarted)
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
		it.Event = new(L1USDCBridgeOwnershipTransferStarted)
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
func (it *L1USDCBridgeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the L1USDCBridge contract.
type L1USDCBridgeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1USDCBridgeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeOwnershipTransferStartedIterator{contract: _L1USDCBridge.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeOwnershipTransferStarted)
				if err := _L1USDCBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseOwnershipTransferStarted(log types.Log) (*L1USDCBridgeOwnershipTransferStarted, error) {
	event := new(L1USDCBridgeOwnershipTransferStarted)
	if err := _L1USDCBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1USDCBridge contract.
type L1USDCBridgeOwnershipTransferredIterator struct {
	Event *L1USDCBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeOwnershipTransferred)
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
		it.Event = new(L1USDCBridgeOwnershipTransferred)
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
func (it *L1USDCBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the L1USDCBridge contract.
type L1USDCBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1USDCBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeOwnershipTransferredIterator{contract: _L1USDCBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeOwnershipTransferred)
				if err := _L1USDCBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*L1USDCBridgeOwnershipTransferred, error) {
	event := new(L1USDCBridgeOwnershipTransferred)
	if err := _L1USDCBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the L1USDCBridge contract.
type L1USDCBridgePausedIterator struct {
	Event *L1USDCBridgePaused // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgePaused)
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
		it.Event = new(L1USDCBridgePaused)
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
func (it *L1USDCBridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgePaused represents a Paused event raised by the L1USDCBridge contract.
type L1USDCBridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*L1USDCBridgePausedIterator, error) {

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgePausedIterator{contract: _L1USDCBridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *L1USDCBridgePaused) (event.Subscription, error) {

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgePaused)
				if err := _L1USDCBridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_L1USDCBridge *L1USDCBridgeFilterer) ParsePaused(log types.Log) (*L1USDCBridgePaused, error) {
	event := new(L1USDCBridgePaused)
	if err := _L1USDCBridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeReceivedFromOtherLayerIterator is returned from FilterReceivedFromOtherLayer and is used to iterate over the raw logs and unpacked data for ReceivedFromOtherLayer events raised by the L1USDCBridge contract.
type L1USDCBridgeReceivedFromOtherLayerIterator struct {
	Event *L1USDCBridgeReceivedFromOtherLayer // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeReceivedFromOtherLayerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeReceivedFromOtherLayer)
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
		it.Event = new(L1USDCBridgeReceivedFromOtherLayer)
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
func (it *L1USDCBridgeReceivedFromOtherLayerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeReceivedFromOtherLayerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeReceivedFromOtherLayer represents a ReceivedFromOtherLayer event raised by the L1USDCBridge contract.
type L1USDCBridgeReceivedFromOtherLayer struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReceivedFromOtherLayer is a free log retrieval operation binding the contract event 0xc79857a8951ecfecbabd2f50997eb17d43f2b30e775d4e471c351d5be1e63b93.
//
// Solidity: event ReceivedFromOtherLayer(address indexed recipient, uint256 indexed amount)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterReceivedFromOtherLayer(opts *bind.FilterOpts, recipient []common.Address, amount []*big.Int) (*L1USDCBridgeReceivedFromOtherLayerIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "ReceivedFromOtherLayer", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeReceivedFromOtherLayerIterator{contract: _L1USDCBridge.contract, event: "ReceivedFromOtherLayer", logs: logs, sub: sub}, nil
}

// WatchReceivedFromOtherLayer is a free log subscription operation binding the contract event 0xc79857a8951ecfecbabd2f50997eb17d43f2b30e775d4e471c351d5be1e63b93.
//
// Solidity: event ReceivedFromOtherLayer(address indexed recipient, uint256 indexed amount)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchReceivedFromOtherLayer(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeReceivedFromOtherLayer, recipient []common.Address, amount []*big.Int) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "ReceivedFromOtherLayer", recipientRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeReceivedFromOtherLayer)
				if err := _L1USDCBridge.contract.UnpackLog(event, "ReceivedFromOtherLayer", log); err != nil {
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

// ParseReceivedFromOtherLayer is a log parse operation binding the contract event 0xc79857a8951ecfecbabd2f50997eb17d43f2b30e775d4e471c351d5be1e63b93.
//
// Solidity: event ReceivedFromOtherLayer(address indexed recipient, uint256 indexed amount)
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseReceivedFromOtherLayer(log types.Log) (*L1USDCBridgeReceivedFromOtherLayer, error) {
	event := new(L1USDCBridgeReceivedFromOtherLayer)
	if err := _L1USDCBridge.contract.UnpackLog(event, "ReceivedFromOtherLayer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeRemoteUSDCBridgeSetIterator is returned from FilterRemoteUSDCBridgeSet and is used to iterate over the raw logs and unpacked data for RemoteUSDCBridgeSet events raised by the L1USDCBridge contract.
type L1USDCBridgeRemoteUSDCBridgeSetIterator struct {
	Event *L1USDCBridgeRemoteUSDCBridgeSet // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeRemoteUSDCBridgeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeRemoteUSDCBridgeSet)
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
		it.Event = new(L1USDCBridgeRemoteUSDCBridgeSet)
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
func (it *L1USDCBridgeRemoteUSDCBridgeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeRemoteUSDCBridgeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeRemoteUSDCBridgeSet represents a RemoteUSDCBridgeSet event raised by the L1USDCBridge contract.
type L1USDCBridgeRemoteUSDCBridgeSet struct {
	NewRemoteUSDCBridge common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRemoteUSDCBridgeSet is a free log retrieval operation binding the contract event 0x5303c9046615f202c2c30e2441fabee4ef1a3597170254aa4870a147c5cb7e34.
//
// Solidity: event RemoteUSDCBridgeSet(address indexed newRemoteUSDCBridge)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterRemoteUSDCBridgeSet(opts *bind.FilterOpts, newRemoteUSDCBridge []common.Address) (*L1USDCBridgeRemoteUSDCBridgeSetIterator, error) {

	var newRemoteUSDCBridgeRule []interface{}
	for _, newRemoteUSDCBridgeItem := range newRemoteUSDCBridge {
		newRemoteUSDCBridgeRule = append(newRemoteUSDCBridgeRule, newRemoteUSDCBridgeItem)
	}

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "RemoteUSDCBridgeSet", newRemoteUSDCBridgeRule)
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeRemoteUSDCBridgeSetIterator{contract: _L1USDCBridge.contract, event: "RemoteUSDCBridgeSet", logs: logs, sub: sub}, nil
}

// WatchRemoteUSDCBridgeSet is a free log subscription operation binding the contract event 0x5303c9046615f202c2c30e2441fabee4ef1a3597170254aa4870a147c5cb7e34.
//
// Solidity: event RemoteUSDCBridgeSet(address indexed newRemoteUSDCBridge)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchRemoteUSDCBridgeSet(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeRemoteUSDCBridgeSet, newRemoteUSDCBridge []common.Address) (event.Subscription, error) {

	var newRemoteUSDCBridgeRule []interface{}
	for _, newRemoteUSDCBridgeItem := range newRemoteUSDCBridge {
		newRemoteUSDCBridgeRule = append(newRemoteUSDCBridgeRule, newRemoteUSDCBridgeItem)
	}

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "RemoteUSDCBridgeSet", newRemoteUSDCBridgeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeRemoteUSDCBridgeSet)
				if err := _L1USDCBridge.contract.UnpackLog(event, "RemoteUSDCBridgeSet", log); err != nil {
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

// ParseRemoteUSDCBridgeSet is a log parse operation binding the contract event 0x5303c9046615f202c2c30e2441fabee4ef1a3597170254aa4870a147c5cb7e34.
//
// Solidity: event RemoteUSDCBridgeSet(address indexed newRemoteUSDCBridge)
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseRemoteUSDCBridgeSet(log types.Log) (*L1USDCBridgeRemoteUSDCBridgeSet, error) {
	event := new(L1USDCBridgeRemoteUSDCBridgeSet)
	if err := _L1USDCBridge.contract.UnpackLog(event, "RemoteUSDCBridgeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1USDCBridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the L1USDCBridge contract.
type L1USDCBridgeUnpausedIterator struct {
	Event *L1USDCBridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *L1USDCBridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1USDCBridgeUnpaused)
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
		it.Event = new(L1USDCBridgeUnpaused)
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
func (it *L1USDCBridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1USDCBridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1USDCBridgeUnpaused represents a Unpaused event raised by the L1USDCBridge contract.
type L1USDCBridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1USDCBridge *L1USDCBridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*L1USDCBridgeUnpausedIterator, error) {

	logs, sub, err := _L1USDCBridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &L1USDCBridgeUnpausedIterator{contract: _L1USDCBridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_L1USDCBridge *L1USDCBridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *L1USDCBridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _L1USDCBridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1USDCBridgeUnpaused)
				if err := _L1USDCBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_L1USDCBridge *L1USDCBridgeFilterer) ParseUnpaused(log types.Log) (*L1USDCBridgeUnpaused, error) {
	event := new(L1USDCBridgeUnpaused)
	if err := _L1USDCBridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
