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

// RegistryTypesHandle is an auto generated low-level Go binding around an user-defined struct.
type RegistryTypesHandle struct {
	Id         *big.Int
	Collection common.Address
}

// RegistryTypesToken is an auto generated low-level Go binding around an user-defined struct.
type RegistryTypesToken struct {
	Id         *big.Int
	Collection common.Address
}

// V2HandleRegistryMetaData contains all meta data concerning the V2HandleRegistry contract.
var V2HandleRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lensHub\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lensHandles\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotHavePermissions\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HandleAndTokenNotInSameWallet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotHandleNorTokenOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLinked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyLensHub\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRegistryTypes.Handle\",\"name\":\"handle\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRegistryTypes.Token\",\"name\":\"token\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HandleLinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRegistryTypes.Handle\",\"name\":\"handle\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structRegistryTypes.Token\",\"name\":\"token\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HandleUnlinked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NonceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getDefaultHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"increment\",\"type\":\"uint8\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"link\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"linkWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"migrationLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"}],\"name\":\"resolve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"unlink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"handleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"unlinkWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V2HandleRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use V2HandleRegistryMetaData.ABI instead.
var V2HandleRegistryABI = V2HandleRegistryMetaData.ABI

// V2HandleRegistry is an auto generated Go binding around an Ethereum contract.
type V2HandleRegistry struct {
	V2HandleRegistryCaller     // Read-only binding to the contract
	V2HandleRegistryTransactor // Write-only binding to the contract
	V2HandleRegistryFilterer   // Log filterer for contract events
}

// V2HandleRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2HandleRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2HandleRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2HandleRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2HandleRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2HandleRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2HandleRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2HandleRegistrySession struct {
	Contract     *V2HandleRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2HandleRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2HandleRegistryCallerSession struct {
	Contract *V2HandleRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// V2HandleRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2HandleRegistryTransactorSession struct {
	Contract     *V2HandleRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// V2HandleRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2HandleRegistryRaw struct {
	Contract *V2HandleRegistry // Generic contract binding to access the raw methods on
}

// V2HandleRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2HandleRegistryCallerRaw struct {
	Contract *V2HandleRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// V2HandleRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2HandleRegistryTransactorRaw struct {
	Contract *V2HandleRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2HandleRegistry creates a new instance of V2HandleRegistry, bound to a specific deployed contract.
func NewV2HandleRegistry(address common.Address, backend bind.ContractBackend) (*V2HandleRegistry, error) {
	contract, err := bindV2HandleRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistry{V2HandleRegistryCaller: V2HandleRegistryCaller{contract: contract}, V2HandleRegistryTransactor: V2HandleRegistryTransactor{contract: contract}, V2HandleRegistryFilterer: V2HandleRegistryFilterer{contract: contract}}, nil
}

// NewV2HandleRegistryCaller creates a new read-only instance of V2HandleRegistry, bound to a specific deployed contract.
func NewV2HandleRegistryCaller(address common.Address, caller bind.ContractCaller) (*V2HandleRegistryCaller, error) {
	contract, err := bindV2HandleRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryCaller{contract: contract}, nil
}

// NewV2HandleRegistryTransactor creates a new write-only instance of V2HandleRegistry, bound to a specific deployed contract.
func NewV2HandleRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*V2HandleRegistryTransactor, error) {
	contract, err := bindV2HandleRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryTransactor{contract: contract}, nil
}

// NewV2HandleRegistryFilterer creates a new log filterer instance of V2HandleRegistry, bound to a specific deployed contract.
func NewV2HandleRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*V2HandleRegistryFilterer, error) {
	contract, err := bindV2HandleRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryFilterer{contract: contract}, nil
}

// bindV2HandleRegistry binds a generic wrapper to an already deployed contract.
func bindV2HandleRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2HandleRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2HandleRegistry *V2HandleRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2HandleRegistry.Contract.V2HandleRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2HandleRegistry *V2HandleRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.V2HandleRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2HandleRegistry *V2HandleRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.V2HandleRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2HandleRegistry *V2HandleRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2HandleRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2HandleRegistry *V2HandleRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2HandleRegistry *V2HandleRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetDefaultHandle is a free data retrieval call binding the contract method 0xe524488d.
//
// Solidity: function getDefaultHandle(uint256 profileId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistryCaller) GetDefaultHandle(opts *bind.CallOpts, profileId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V2HandleRegistry.contract.Call(opts, &out, "getDefaultHandle", profileId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDefaultHandle is a free data retrieval call binding the contract method 0xe524488d.
//
// Solidity: function getDefaultHandle(uint256 profileId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistrySession) GetDefaultHandle(profileId *big.Int) (*big.Int, error) {
	return _V2HandleRegistry.Contract.GetDefaultHandle(&_V2HandleRegistry.CallOpts, profileId)
}

// GetDefaultHandle is a free data retrieval call binding the contract method 0xe524488d.
//
// Solidity: function getDefaultHandle(uint256 profileId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistryCallerSession) GetDefaultHandle(profileId *big.Int) (*big.Int, error) {
	return _V2HandleRegistry.Contract.GetDefaultHandle(&_V2HandleRegistry.CallOpts, profileId)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_V2HandleRegistry *V2HandleRegistryCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V2HandleRegistry.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_V2HandleRegistry *V2HandleRegistrySession) Nonces(signer common.Address) (*big.Int, error) {
	return _V2HandleRegistry.Contract.Nonces(&_V2HandleRegistry.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256 nonce)
func (_V2HandleRegistry *V2HandleRegistryCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _V2HandleRegistry.Contract.Nonces(&_V2HandleRegistry.CallOpts, signer)
}

// Resolve is a free data retrieval call binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 handleId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistryCaller) Resolve(opts *bind.CallOpts, handleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V2HandleRegistry.contract.Call(opts, &out, "resolve", handleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 handleId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistrySession) Resolve(handleId *big.Int) (*big.Int, error) {
	return _V2HandleRegistry.Contract.Resolve(&_V2HandleRegistry.CallOpts, handleId)
}

// Resolve is a free data retrieval call binding the contract method 0x4f896d4f.
//
// Solidity: function resolve(uint256 handleId) view returns(uint256)
func (_V2HandleRegistry *V2HandleRegistryCallerSession) Resolve(handleId *big.Int) (*big.Int, error) {
	return _V2HandleRegistry.Contract.Resolve(&_V2HandleRegistry.CallOpts, handleId)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) IncrementNonce(opts *bind.TransactOpts, increment uint8) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "incrementNonce", increment)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) IncrementNonce(increment uint8) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.IncrementNonce(&_V2HandleRegistry.TransactOpts, increment)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) IncrementNonce(increment uint8) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.IncrementNonce(&_V2HandleRegistry.TransactOpts, increment)
}

// Link is a paid mutator transaction binding the contract method 0x86cf48e7.
//
// Solidity: function link(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) Link(opts *bind.TransactOpts, handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "link", handleId, profileId)
}

// Link is a paid mutator transaction binding the contract method 0x86cf48e7.
//
// Solidity: function link(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) Link(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.Link(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// Link is a paid mutator transaction binding the contract method 0x86cf48e7.
//
// Solidity: function link(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) Link(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.Link(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// LinkWithSig is a paid mutator transaction binding the contract method 0x7e502fe0.
//
// Solidity: function linkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) LinkWithSig(opts *bind.TransactOpts, handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "linkWithSig", handleId, profileId, signature)
}

// LinkWithSig is a paid mutator transaction binding the contract method 0x7e502fe0.
//
// Solidity: function linkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) LinkWithSig(handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.LinkWithSig(&_V2HandleRegistry.TransactOpts, handleId, profileId, signature)
}

// LinkWithSig is a paid mutator transaction binding the contract method 0x7e502fe0.
//
// Solidity: function linkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) LinkWithSig(handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.LinkWithSig(&_V2HandleRegistry.TransactOpts, handleId, profileId, signature)
}

// MigrationLink is a paid mutator transaction binding the contract method 0xb418a9a3.
//
// Solidity: function migrationLink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) MigrationLink(opts *bind.TransactOpts, handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "migrationLink", handleId, profileId)
}

// MigrationLink is a paid mutator transaction binding the contract method 0xb418a9a3.
//
// Solidity: function migrationLink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) MigrationLink(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.MigrationLink(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// MigrationLink is a paid mutator transaction binding the contract method 0xb418a9a3.
//
// Solidity: function migrationLink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) MigrationLink(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.MigrationLink(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// Unlink is a paid mutator transaction binding the contract method 0x0c65b39d.
//
// Solidity: function unlink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) Unlink(opts *bind.TransactOpts, handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "unlink", handleId, profileId)
}

// Unlink is a paid mutator transaction binding the contract method 0x0c65b39d.
//
// Solidity: function unlink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) Unlink(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.Unlink(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// Unlink is a paid mutator transaction binding the contract method 0x0c65b39d.
//
// Solidity: function unlink(uint256 handleId, uint256 profileId) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) Unlink(handleId *big.Int, profileId *big.Int) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.Unlink(&_V2HandleRegistry.TransactOpts, handleId, profileId)
}

// UnlinkWithSig is a paid mutator transaction binding the contract method 0x825ab164.
//
// Solidity: function unlinkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactor) UnlinkWithSig(opts *bind.TransactOpts, handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.contract.Transact(opts, "unlinkWithSig", handleId, profileId, signature)
}

// UnlinkWithSig is a paid mutator transaction binding the contract method 0x825ab164.
//
// Solidity: function unlinkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistrySession) UnlinkWithSig(handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.UnlinkWithSig(&_V2HandleRegistry.TransactOpts, handleId, profileId, signature)
}

// UnlinkWithSig is a paid mutator transaction binding the contract method 0x825ab164.
//
// Solidity: function unlinkWithSig(uint256 handleId, uint256 profileId, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2HandleRegistry *V2HandleRegistryTransactorSession) UnlinkWithSig(handleId *big.Int, profileId *big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2HandleRegistry.Contract.UnlinkWithSig(&_V2HandleRegistry.TransactOpts, handleId, profileId, signature)
}

// V2HandleRegistryHandleLinkedIterator is returned from FilterHandleLinked and is used to iterate over the raw logs and unpacked data for HandleLinked events raised by the V2HandleRegistry contract.
type V2HandleRegistryHandleLinkedIterator struct {
	Event *V2HandleRegistryHandleLinked // Event containing the contract specifics and raw log

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
func (it *V2HandleRegistryHandleLinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2HandleRegistryHandleLinked)
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
		it.Event = new(V2HandleRegistryHandleLinked)
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
func (it *V2HandleRegistryHandleLinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2HandleRegistryHandleLinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2HandleRegistryHandleLinked represents a HandleLinked event raised by the V2HandleRegistry contract.
type V2HandleRegistryHandleLinked struct {
	Handle              RegistryTypesHandle
	Token               RegistryTypesToken
	TransactionExecutor common.Address
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHandleLinked is a free log retrieval operation binding the contract event 0xc33717b10968a153b6fa59edb9356cc86846bd3be6f416f4410338e25bb28d3c.
//
// Solidity: event HandleLinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) FilterHandleLinked(opts *bind.FilterOpts) (*V2HandleRegistryHandleLinkedIterator, error) {

	logs, sub, err := _V2HandleRegistry.contract.FilterLogs(opts, "HandleLinked")
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryHandleLinkedIterator{contract: _V2HandleRegistry.contract, event: "HandleLinked", logs: logs, sub: sub}, nil
}

// WatchHandleLinked is a free log subscription operation binding the contract event 0xc33717b10968a153b6fa59edb9356cc86846bd3be6f416f4410338e25bb28d3c.
//
// Solidity: event HandleLinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) WatchHandleLinked(opts *bind.WatchOpts, sink chan<- *V2HandleRegistryHandleLinked) (event.Subscription, error) {

	logs, sub, err := _V2HandleRegistry.contract.WatchLogs(opts, "HandleLinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2HandleRegistryHandleLinked)
				if err := _V2HandleRegistry.contract.UnpackLog(event, "HandleLinked", log); err != nil {
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

// ParseHandleLinked is a log parse operation binding the contract event 0xc33717b10968a153b6fa59edb9356cc86846bd3be6f416f4410338e25bb28d3c.
//
// Solidity: event HandleLinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) ParseHandleLinked(log types.Log) (*V2HandleRegistryHandleLinked, error) {
	event := new(V2HandleRegistryHandleLinked)
	if err := _V2HandleRegistry.contract.UnpackLog(event, "HandleLinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2HandleRegistryHandleUnlinkedIterator is returned from FilterHandleUnlinked and is used to iterate over the raw logs and unpacked data for HandleUnlinked events raised by the V2HandleRegistry contract.
type V2HandleRegistryHandleUnlinkedIterator struct {
	Event *V2HandleRegistryHandleUnlinked // Event containing the contract specifics and raw log

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
func (it *V2HandleRegistryHandleUnlinkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2HandleRegistryHandleUnlinked)
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
		it.Event = new(V2HandleRegistryHandleUnlinked)
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
func (it *V2HandleRegistryHandleUnlinkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2HandleRegistryHandleUnlinkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2HandleRegistryHandleUnlinked represents a HandleUnlinked event raised by the V2HandleRegistry contract.
type V2HandleRegistryHandleUnlinked struct {
	Handle              RegistryTypesHandle
	Token               RegistryTypesToken
	TransactionExecutor common.Address
	Timestamp           *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterHandleUnlinked is a free log retrieval operation binding the contract event 0xc5b450c435c821270ac9ca0c97936e8b0a38575bb06b6250cbf5be0c3f44daa3.
//
// Solidity: event HandleUnlinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) FilterHandleUnlinked(opts *bind.FilterOpts) (*V2HandleRegistryHandleUnlinkedIterator, error) {

	logs, sub, err := _V2HandleRegistry.contract.FilterLogs(opts, "HandleUnlinked")
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryHandleUnlinkedIterator{contract: _V2HandleRegistry.contract, event: "HandleUnlinked", logs: logs, sub: sub}, nil
}

// WatchHandleUnlinked is a free log subscription operation binding the contract event 0xc5b450c435c821270ac9ca0c97936e8b0a38575bb06b6250cbf5be0c3f44daa3.
//
// Solidity: event HandleUnlinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) WatchHandleUnlinked(opts *bind.WatchOpts, sink chan<- *V2HandleRegistryHandleUnlinked) (event.Subscription, error) {

	logs, sub, err := _V2HandleRegistry.contract.WatchLogs(opts, "HandleUnlinked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2HandleRegistryHandleUnlinked)
				if err := _V2HandleRegistry.contract.UnpackLog(event, "HandleUnlinked", log); err != nil {
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

// ParseHandleUnlinked is a log parse operation binding the contract event 0xc5b450c435c821270ac9ca0c97936e8b0a38575bb06b6250cbf5be0c3f44daa3.
//
// Solidity: event HandleUnlinked((uint256,address) handle, (uint256,address) token, address transactionExecutor, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) ParseHandleUnlinked(log types.Log) (*V2HandleRegistryHandleUnlinked, error) {
	event := new(V2HandleRegistryHandleUnlinked)
	if err := _V2HandleRegistry.contract.UnpackLog(event, "HandleUnlinked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2HandleRegistryNonceUpdatedIterator is returned from FilterNonceUpdated and is used to iterate over the raw logs and unpacked data for NonceUpdated events raised by the V2HandleRegistry contract.
type V2HandleRegistryNonceUpdatedIterator struct {
	Event *V2HandleRegistryNonceUpdated // Event containing the contract specifics and raw log

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
func (it *V2HandleRegistryNonceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2HandleRegistryNonceUpdated)
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
		it.Event = new(V2HandleRegistryNonceUpdated)
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
func (it *V2HandleRegistryNonceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2HandleRegistryNonceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2HandleRegistryNonceUpdated represents a NonceUpdated event raised by the V2HandleRegistry contract.
type V2HandleRegistryNonceUpdated struct {
	Signer    common.Address
	Nonce     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNonceUpdated is a free log retrieval operation binding the contract event 0xc906270cebe7667882104effe64262a73c422ab9176a111e05ea837b021065fc.
//
// Solidity: event NonceUpdated(address indexed signer, uint256 nonce, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) FilterNonceUpdated(opts *bind.FilterOpts, signer []common.Address) (*V2HandleRegistryNonceUpdatedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _V2HandleRegistry.contract.FilterLogs(opts, "NonceUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return &V2HandleRegistryNonceUpdatedIterator{contract: _V2HandleRegistry.contract, event: "NonceUpdated", logs: logs, sub: sub}, nil
}

// WatchNonceUpdated is a free log subscription operation binding the contract event 0xc906270cebe7667882104effe64262a73c422ab9176a111e05ea837b021065fc.
//
// Solidity: event NonceUpdated(address indexed signer, uint256 nonce, uint256 timestamp)
func (_V2HandleRegistry *V2HandleRegistryFilterer) WatchNonceUpdated(opts *bind.WatchOpts, sink chan<- *V2HandleRegistryNonceUpdated, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _V2HandleRegistry.contract.WatchLogs(opts, "NonceUpdated", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2HandleRegistryNonceUpdated)
				if err := _V2HandleRegistry.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
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
func (_V2HandleRegistry *V2HandleRegistryFilterer) ParseNonceUpdated(log types.Log) (*V2HandleRegistryNonceUpdated, error) {
	event := new(V2HandleRegistryNonceUpdated)
	if err := _V2HandleRegistry.contract.UnpackLog(event, "NonceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
