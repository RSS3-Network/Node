// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vsl

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

// NetworkParamsMetaData contains all meta data concerning the NetworkParams contract.
var NetworkParamsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParams\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"adminAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setParams\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"params\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ParamsSet\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"params\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]}]",
}

// NetworkParamsABI is the input ABI used to generate the binding from.
// Deprecated: Use NetworkParamsMetaData.ABI instead.
var NetworkParamsABI = NetworkParamsMetaData.ABI

// NetworkParams is an auto generated Go binding around an Ethereum contract.
type NetworkParams struct {
	NetworkParamsCaller     // Read-only binding to the contract
	NetworkParamsTransactor // Write-only binding to the contract
	NetworkParamsFilterer   // Log filterer for contract events
}

// NetworkParamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NetworkParamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkParamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NetworkParamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkParamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NetworkParamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NetworkParamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NetworkParamsSession struct {
	Contract     *NetworkParams    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NetworkParamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NetworkParamsCallerSession struct {
	Contract *NetworkParamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NetworkParamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NetworkParamsTransactorSession struct {
	Contract     *NetworkParamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NetworkParamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NetworkParamsRaw struct {
	Contract *NetworkParams // Generic contract binding to access the raw methods on
}

// NetworkParamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NetworkParamsCallerRaw struct {
	Contract *NetworkParamsCaller // Generic read-only contract binding to access the raw methods on
}

// NetworkParamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NetworkParamsTransactorRaw struct {
	Contract *NetworkParamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNetworkParams creates a new instance of NetworkParams, bound to a specific deployed contract.
func NewNetworkParams(address common.Address, backend bind.ContractBackend) (*NetworkParams, error) {
	contract, err := bindNetworkParams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NetworkParams{NetworkParamsCaller: NetworkParamsCaller{contract: contract}, NetworkParamsTransactor: NetworkParamsTransactor{contract: contract}, NetworkParamsFilterer: NetworkParamsFilterer{contract: contract}}, nil
}

// NewNetworkParamsCaller creates a new read-only instance of NetworkParams, bound to a specific deployed contract.
func NewNetworkParamsCaller(address common.Address, caller bind.ContractCaller) (*NetworkParamsCaller, error) {
	contract, err := bindNetworkParams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsCaller{contract: contract}, nil
}

// NewNetworkParamsTransactor creates a new write-only instance of NetworkParams, bound to a specific deployed contract.
func NewNetworkParamsTransactor(address common.Address, transactor bind.ContractTransactor) (*NetworkParamsTransactor, error) {
	contract, err := bindNetworkParams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsTransactor{contract: contract}, nil
}

// NewNetworkParamsFilterer creates a new log filterer instance of NetworkParams, bound to a specific deployed contract.
func NewNetworkParamsFilterer(address common.Address, filterer bind.ContractFilterer) (*NetworkParamsFilterer, error) {
	contract, err := bindNetworkParams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsFilterer{contract: contract}, nil
}

// bindNetworkParams binds a generic wrapper to an already deployed contract.
func bindNetworkParams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NetworkParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkParams *NetworkParamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkParams.Contract.NetworkParamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkParams *NetworkParamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkParams.Contract.NetworkParamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkParams *NetworkParamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkParams.Contract.NetworkParamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NetworkParams *NetworkParamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NetworkParams.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NetworkParams *NetworkParamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NetworkParams.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NetworkParams *NetworkParamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NetworkParams.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsSession) ADMINROLE() ([32]byte, error) {
	return _NetworkParams.Contract.ADMINROLE(&_NetworkParams.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsCallerSession) ADMINROLE() ([32]byte, error) {
	return _NetworkParams.Contract.ADMINROLE(&_NetworkParams.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NetworkParams.Contract.DEFAULTADMINROLE(&_NetworkParams.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NetworkParams *NetworkParamsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NetworkParams.Contract.DEFAULTADMINROLE(&_NetworkParams.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0xace30dbc.
//
// Solidity: function getParams(uint64 epoch) view returns(string)
func (_NetworkParams *NetworkParamsCaller) GetParams(opts *bind.CallOpts, epoch uint64) (string, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "getParams", epoch)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetParams is a free data retrieval call binding the contract method 0xace30dbc.
//
// Solidity: function getParams(uint64 epoch) view returns(string)
func (_NetworkParams *NetworkParamsSession) GetParams(epoch uint64) (string, error) {
	return _NetworkParams.Contract.GetParams(&_NetworkParams.CallOpts, epoch)
}

// GetParams is a free data retrieval call binding the contract method 0xace30dbc.
//
// Solidity: function getParams(uint64 epoch) view returns(string)
func (_NetworkParams *NetworkParamsCallerSession) GetParams(epoch uint64) (string, error) {
	return _NetworkParams.Contract.GetParams(&_NetworkParams.CallOpts, epoch)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkParams *NetworkParamsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkParams *NetworkParamsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NetworkParams.Contract.GetRoleAdmin(&_NetworkParams.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NetworkParams *NetworkParamsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NetworkParams.Contract.GetRoleAdmin(&_NetworkParams.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NetworkParams *NetworkParamsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NetworkParams *NetworkParamsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NetworkParams.Contract.GetRoleMember(&_NetworkParams.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NetworkParams *NetworkParamsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NetworkParams.Contract.GetRoleMember(&_NetworkParams.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NetworkParams *NetworkParamsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NetworkParams *NetworkParamsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NetworkParams.Contract.GetRoleMemberCount(&_NetworkParams.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NetworkParams *NetworkParamsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NetworkParams.Contract.GetRoleMemberCount(&_NetworkParams.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkParams *NetworkParamsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkParams *NetworkParamsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NetworkParams.Contract.HasRole(&_NetworkParams.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NetworkParams *NetworkParamsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NetworkParams.Contract.HasRole(&_NetworkParams.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkParams *NetworkParamsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NetworkParams.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkParams *NetworkParamsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NetworkParams.Contract.SupportsInterface(&_NetworkParams.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NetworkParams *NetworkParamsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NetworkParams.Contract.SupportsInterface(&_NetworkParams.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.GrantRole(&_NetworkParams.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.GrantRole(&_NetworkParams.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address adminAccount) returns()
func (_NetworkParams *NetworkParamsTransactor) Initialize(opts *bind.TransactOpts, adminAccount common.Address) (*types.Transaction, error) {
	return _NetworkParams.contract.Transact(opts, "initialize", adminAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address adminAccount) returns()
func (_NetworkParams *NetworkParamsSession) Initialize(adminAccount common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.Initialize(&_NetworkParams.TransactOpts, adminAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address adminAccount) returns()
func (_NetworkParams *NetworkParamsTransactorSession) Initialize(adminAccount common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.Initialize(&_NetworkParams.TransactOpts, adminAccount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkParams *NetworkParamsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkParams.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkParams *NetworkParamsSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.RenounceRole(&_NetworkParams.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_NetworkParams *NetworkParamsTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.RenounceRole(&_NetworkParams.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.RevokeRole(&_NetworkParams.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NetworkParams *NetworkParamsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NetworkParams.Contract.RevokeRole(&_NetworkParams.TransactOpts, role, account)
}

// SetParams is a paid mutator transaction binding the contract method 0x28bc7935.
//
// Solidity: function setParams(uint64 epoch, string params) returns()
func (_NetworkParams *NetworkParamsTransactor) SetParams(opts *bind.TransactOpts, epoch uint64, params string) (*types.Transaction, error) {
	return _NetworkParams.contract.Transact(opts, "setParams", epoch, params)
}

// SetParams is a paid mutator transaction binding the contract method 0x28bc7935.
//
// Solidity: function setParams(uint64 epoch, string params) returns()
func (_NetworkParams *NetworkParamsSession) SetParams(epoch uint64, params string) (*types.Transaction, error) {
	return _NetworkParams.Contract.SetParams(&_NetworkParams.TransactOpts, epoch, params)
}

// SetParams is a paid mutator transaction binding the contract method 0x28bc7935.
//
// Solidity: function setParams(uint64 epoch, string params) returns()
func (_NetworkParams *NetworkParamsTransactorSession) SetParams(epoch uint64, params string) (*types.Transaction, error) {
	return _NetworkParams.Contract.SetParams(&_NetworkParams.TransactOpts, epoch, params)
}

// NetworkParamsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NetworkParams contract.
type NetworkParamsInitializedIterator struct {
	Event *NetworkParamsInitialized // Event containing the contract specifics and raw log

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
func (it *NetworkParamsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkParamsInitialized)
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
		it.Event = new(NetworkParamsInitialized)
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
func (it *NetworkParamsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkParamsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkParamsInitialized represents a Initialized event raised by the NetworkParams contract.
type NetworkParamsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkParams *NetworkParamsFilterer) FilterInitialized(opts *bind.FilterOpts) (*NetworkParamsInitializedIterator, error) {

	logs, sub, err := _NetworkParams.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NetworkParamsInitializedIterator{contract: _NetworkParams.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkParams *NetworkParamsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NetworkParamsInitialized) (event.Subscription, error) {

	logs, sub, err := _NetworkParams.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkParamsInitialized)
				if err := _NetworkParams.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NetworkParams *NetworkParamsFilterer) ParseInitialized(log types.Log) (*NetworkParamsInitialized, error) {
	event := new(NetworkParamsInitialized)
	if err := _NetworkParams.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkParamsParamsSetIterator is returned from FilterParamsSet and is used to iterate over the raw logs and unpacked data for ParamsSet events raised by the NetworkParams contract.
type NetworkParamsParamsSetIterator struct {
	Event *NetworkParamsParamsSet // Event containing the contract specifics and raw log

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
func (it *NetworkParamsParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkParamsParamsSet)
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
		it.Event = new(NetworkParamsParamsSet)
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
func (it *NetworkParamsParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkParamsParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkParamsParamsSet represents a ParamsSet event raised by the NetworkParams contract.
type NetworkParamsParamsSet struct {
	Epoch  uint64
	Params string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterParamsSet is a free log retrieval operation binding the contract event 0xcb46857e4fc83e0b5ff8a04732e1bba8fb60164f0f93baf3f27af5e24be8339a.
//
// Solidity: event ParamsSet(uint64 indexed epoch, string params)
func (_NetworkParams *NetworkParamsFilterer) FilterParamsSet(opts *bind.FilterOpts, epoch []uint64) (*NetworkParamsParamsSetIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _NetworkParams.contract.FilterLogs(opts, "ParamsSet", epochRule)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsParamsSetIterator{contract: _NetworkParams.contract, event: "ParamsSet", logs: logs, sub: sub}, nil
}

// WatchParamsSet is a free log subscription operation binding the contract event 0xcb46857e4fc83e0b5ff8a04732e1bba8fb60164f0f93baf3f27af5e24be8339a.
//
// Solidity: event ParamsSet(uint64 indexed epoch, string params)
func (_NetworkParams *NetworkParamsFilterer) WatchParamsSet(opts *bind.WatchOpts, sink chan<- *NetworkParamsParamsSet, epoch []uint64) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _NetworkParams.contract.WatchLogs(opts, "ParamsSet", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkParamsParamsSet)
				if err := _NetworkParams.contract.UnpackLog(event, "ParamsSet", log); err != nil {
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

// ParseParamsSet is a log parse operation binding the contract event 0xcb46857e4fc83e0b5ff8a04732e1bba8fb60164f0f93baf3f27af5e24be8339a.
//
// Solidity: event ParamsSet(uint64 indexed epoch, string params)
func (_NetworkParams *NetworkParamsFilterer) ParseParamsSet(log types.Log) (*NetworkParamsParamsSet, error) {
	event := new(NetworkParamsParamsSet)
	if err := _NetworkParams.contract.UnpackLog(event, "ParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkParamsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NetworkParams contract.
type NetworkParamsRoleAdminChangedIterator struct {
	Event *NetworkParamsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NetworkParamsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkParamsRoleAdminChanged)
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
		it.Event = new(NetworkParamsRoleAdminChanged)
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
func (it *NetworkParamsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkParamsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkParamsRoleAdminChanged represents a RoleAdminChanged event raised by the NetworkParams contract.
type NetworkParamsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkParams *NetworkParamsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NetworkParamsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NetworkParams.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsRoleAdminChangedIterator{contract: _NetworkParams.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkParams *NetworkParamsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NetworkParamsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _NetworkParams.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkParamsRoleAdminChanged)
				if err := _NetworkParams.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NetworkParams *NetworkParamsFilterer) ParseRoleAdminChanged(log types.Log) (*NetworkParamsRoleAdminChanged, error) {
	event := new(NetworkParamsRoleAdminChanged)
	if err := _NetworkParams.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkParamsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NetworkParams contract.
type NetworkParamsRoleGrantedIterator struct {
	Event *NetworkParamsRoleGranted // Event containing the contract specifics and raw log

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
func (it *NetworkParamsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkParamsRoleGranted)
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
		it.Event = new(NetworkParamsRoleGranted)
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
func (it *NetworkParamsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkParamsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkParamsRoleGranted represents a RoleGranted event raised by the NetworkParams contract.
type NetworkParamsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetworkParamsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkParams.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsRoleGrantedIterator{contract: _NetworkParams.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NetworkParamsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkParams.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkParamsRoleGranted)
				if err := _NetworkParams.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) ParseRoleGranted(log types.Log) (*NetworkParamsRoleGranted, error) {
	event := new(NetworkParamsRoleGranted)
	if err := _NetworkParams.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NetworkParamsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NetworkParams contract.
type NetworkParamsRoleRevokedIterator struct {
	Event *NetworkParamsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NetworkParamsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NetworkParamsRoleRevoked)
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
		it.Event = new(NetworkParamsRoleRevoked)
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
func (it *NetworkParamsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NetworkParamsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NetworkParamsRoleRevoked represents a RoleRevoked event raised by the NetworkParams contract.
type NetworkParamsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NetworkParamsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkParams.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NetworkParamsRoleRevokedIterator{contract: _NetworkParams.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NetworkParamsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _NetworkParams.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NetworkParamsRoleRevoked)
				if err := _NetworkParams.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NetworkParams *NetworkParamsFilterer) ParseRoleRevoked(log types.Log) (*NetworkParamsRoleRevoked, error) {
	event := new(NetworkParamsRoleRevoked)
	if err := _NetworkParams.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
