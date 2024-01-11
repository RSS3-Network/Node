// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ens

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

// IPriceOraclePrice is an auto generated low-level Go binding around an user-defined struct.
type IPriceOraclePrice struct {
	Base    *big.Int
	Premium *big.Int
}

// ETHRegistrarControllerV2MetaData contains all meta data concerning the ETHRegistrarControllerV2 contract.
var ETHRegistrarControllerV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrarImplementation\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"contractIPriceOracle\",\"name\":\"_prices\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxCommitmentAge\",\"type\":\"uint256\"},{\"internalType\":\"contractReverseRegistrar\",\"name\":\"_reverseRegistrar\",\"type\":\"address\"},{\"internalType\":\"contractINameWrapper\",\"name\":\"_nameWrapper\",\"type\":\"address\"},{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooNew\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentTooOld\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"DurationTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxCommitmentAgeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"NameNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResolverRequiredWhenDataSupplied\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"UnexpiredCommitmentExists\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseCost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"commitment\",\"type\":\"bytes32\"}],\"name\":\"commit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"commitments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"makeCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minCommitmentAge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameWrapper\",\"outputs\":[{\"internalType\":\"contractINameWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prices\",\"outputs\":[{\"internalType\":\"contractIPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"},{\"internalType\":\"bool\",\"name\":\"reverseRecord\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"ownerControlledFuses\",\"type\":\"uint16\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"base\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"internalType\":\"structIPriceOracle.Price\",\"name\":\"price\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reverseRegistrar\",\"outputs\":[{\"internalType\":\"contractReverseRegistrar\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ETHRegistrarControllerV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHRegistrarControllerV2MetaData.ABI instead.
var ETHRegistrarControllerV2ABI = ETHRegistrarControllerV2MetaData.ABI

// ETHRegistrarControllerV2 is an auto generated Go binding around an Ethereum contract.
type ETHRegistrarControllerV2 struct {
	ETHRegistrarControllerV2Caller     // Read-only binding to the contract
	ETHRegistrarControllerV2Transactor // Write-only binding to the contract
	ETHRegistrarControllerV2Filterer   // Log filterer for contract events
}

// ETHRegistrarControllerV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ETHRegistrarControllerV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRegistrarControllerV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHRegistrarControllerV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRegistrarControllerV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHRegistrarControllerV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHRegistrarControllerV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHRegistrarControllerV2Session struct {
	Contract     *ETHRegistrarControllerV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ETHRegistrarControllerV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHRegistrarControllerV2CallerSession struct {
	Contract *ETHRegistrarControllerV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ETHRegistrarControllerV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHRegistrarControllerV2TransactorSession struct {
	Contract     *ETHRegistrarControllerV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ETHRegistrarControllerV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ETHRegistrarControllerV2Raw struct {
	Contract *ETHRegistrarControllerV2 // Generic contract binding to access the raw methods on
}

// ETHRegistrarControllerV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHRegistrarControllerV2CallerRaw struct {
	Contract *ETHRegistrarControllerV2Caller // Generic read-only contract binding to access the raw methods on
}

// ETHRegistrarControllerV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHRegistrarControllerV2TransactorRaw struct {
	Contract *ETHRegistrarControllerV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewETHRegistrarControllerV2 creates a new instance of ETHRegistrarControllerV2, bound to a specific deployed contract.
func NewETHRegistrarControllerV2(address common.Address, backend bind.ContractBackend) (*ETHRegistrarControllerV2, error) {
	contract, err := bindETHRegistrarControllerV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2{ETHRegistrarControllerV2Caller: ETHRegistrarControllerV2Caller{contract: contract}, ETHRegistrarControllerV2Transactor: ETHRegistrarControllerV2Transactor{contract: contract}, ETHRegistrarControllerV2Filterer: ETHRegistrarControllerV2Filterer{contract: contract}}, nil
}

// NewETHRegistrarControllerV2Caller creates a new read-only instance of ETHRegistrarControllerV2, bound to a specific deployed contract.
func NewETHRegistrarControllerV2Caller(address common.Address, caller bind.ContractCaller) (*ETHRegistrarControllerV2Caller, error) {
	contract, err := bindETHRegistrarControllerV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2Caller{contract: contract}, nil
}

// NewETHRegistrarControllerV2Transactor creates a new write-only instance of ETHRegistrarControllerV2, bound to a specific deployed contract.
func NewETHRegistrarControllerV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ETHRegistrarControllerV2Transactor, error) {
	contract, err := bindETHRegistrarControllerV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2Transactor{contract: contract}, nil
}

// NewETHRegistrarControllerV2Filterer creates a new log filterer instance of ETHRegistrarControllerV2, bound to a specific deployed contract.
func NewETHRegistrarControllerV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ETHRegistrarControllerV2Filterer, error) {
	contract, err := bindETHRegistrarControllerV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2Filterer{contract: contract}, nil
}

// bindETHRegistrarControllerV2 binds a generic wrapper to an already deployed contract.
func bindETHRegistrarControllerV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ETHRegistrarControllerV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHRegistrarControllerV2.Contract.ETHRegistrarControllerV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.ETHRegistrarControllerV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.ETHRegistrarControllerV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHRegistrarControllerV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MINREGISTRATIONDURATION(&_ETHRegistrarControllerV2.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MINREGISTRATIONDURATION(&_ETHRegistrarControllerV2.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Available(name string) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.Available(&_ETHRegistrarControllerV2.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) Available(name string) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.Available(&_ETHRegistrarControllerV2.CallOpts, name)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) Commitments(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "commitments", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.Commitments(&_ETHRegistrarControllerV2.CallOpts, arg0)
}

// Commitments is a free data retrieval call binding the contract method 0x839df945.
//
// Solidity: function commitments(bytes32 ) view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) Commitments(arg0 [32]byte) (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.Commitments(&_ETHRegistrarControllerV2.CallOpts, arg0)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) MakeCommitment(opts *bind.CallOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "makeCommitment", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _ETHRegistrarControllerV2.Contract.MakeCommitment(&_ETHRegistrarControllerV2.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MakeCommitment is a free data retrieval call binding the contract method 0x65a69dcf.
//
// Solidity: function makeCommitment(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) pure returns(bytes32)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) MakeCommitment(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) ([32]byte, error) {
	return _ETHRegistrarControllerV2.Contract.MakeCommitment(&_ETHRegistrarControllerV2.CallOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) MaxCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "maxCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) MaxCommitmentAge() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MaxCommitmentAge(&_ETHRegistrarControllerV2.CallOpts)
}

// MaxCommitmentAge is a free data retrieval call binding the contract method 0xce1e09c0.
//
// Solidity: function maxCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) MaxCommitmentAge() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MaxCommitmentAge(&_ETHRegistrarControllerV2.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) MinCommitmentAge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "minCommitmentAge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) MinCommitmentAge() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MinCommitmentAge(&_ETHRegistrarControllerV2.CallOpts)
}

// MinCommitmentAge is a free data retrieval call binding the contract method 0x8d839ffe.
//
// Solidity: function minCommitmentAge() view returns(uint256)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) MinCommitmentAge() (*big.Int, error) {
	return _ETHRegistrarControllerV2.Contract.MinCommitmentAge(&_ETHRegistrarControllerV2.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) NameWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "nameWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) NameWrapper() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.NameWrapper(&_ETHRegistrarControllerV2.CallOpts)
}

// NameWrapper is a free data retrieval call binding the contract method 0xa8e5fbc0.
//
// Solidity: function nameWrapper() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) NameWrapper() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.NameWrapper(&_ETHRegistrarControllerV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Owner() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.Owner(&_ETHRegistrarControllerV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) Owner() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.Owner(&_ETHRegistrarControllerV2.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) Prices(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "prices")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Prices() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.Prices(&_ETHRegistrarControllerV2.CallOpts)
}

// Prices is a free data retrieval call binding the contract method 0xd3419bf3.
//
// Solidity: function prices() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) Prices() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.Prices(&_ETHRegistrarControllerV2.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) RentPrice(opts *bind.CallOpts, name string, duration *big.Int) (IPriceOraclePrice, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "rentPrice", name, duration)

	if err != nil {
		return *new(IPriceOraclePrice), err
	}

	out0 := *abi.ConvertType(out[0], new(IPriceOraclePrice)).(*IPriceOraclePrice)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _ETHRegistrarControllerV2.Contract.RentPrice(&_ETHRegistrarControllerV2.CallOpts, name, duration)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 duration) view returns((uint256,uint256) price)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) RentPrice(name string, duration *big.Int) (IPriceOraclePrice, error) {
	return _ETHRegistrarControllerV2.Contract.RentPrice(&_ETHRegistrarControllerV2.CallOpts, name, duration)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) ReverseRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "reverseRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) ReverseRegistrar() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.ReverseRegistrar(&_ETHRegistrarControllerV2.CallOpts)
}

// ReverseRegistrar is a free data retrieval call binding the contract method 0x80869853.
//
// Solidity: function reverseRegistrar() view returns(address)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) ReverseRegistrar() (common.Address, error) {
	return _ETHRegistrarControllerV2.Contract.ReverseRegistrar(&_ETHRegistrarControllerV2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.SupportsInterface(&_ETHRegistrarControllerV2.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.SupportsInterface(&_ETHRegistrarControllerV2.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Caller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ETHRegistrarControllerV2.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Valid(name string) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.Valid(&_ETHRegistrarControllerV2.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2CallerSession) Valid(name string) (bool, error) {
	return _ETHRegistrarControllerV2.Contract.Valid(&_ETHRegistrarControllerV2.CallOpts, name)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) Commit(opts *bind.TransactOpts, commitment [32]byte) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "commit", commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Commit(&_ETHRegistrarControllerV2.TransactOpts, commitment)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 commitment) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) Commit(commitment [32]byte) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Commit(&_ETHRegistrarControllerV2.TransactOpts, commitment)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) RecoverFunds(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "recoverFunds", _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.RecoverFunds(&_ETHRegistrarControllerV2.TransactOpts, _token, _to, _amount)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0x5d3590d5.
//
// Solidity: function recoverFunds(address _token, address _to, uint256 _amount) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) RecoverFunds(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.RecoverFunds(&_ETHRegistrarControllerV2.TransactOpts, _token, _to, _amount)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) Register(opts *bind.TransactOpts, name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "register", name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Register(&_ETHRegistrarControllerV2.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Register is a paid mutator transaction binding the contract method 0x74694a2b.
//
// Solidity: function register(string name, address owner, uint256 duration, bytes32 secret, address resolver, bytes[] data, bool reverseRecord, uint16 ownerControlledFuses) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) Register(name string, owner common.Address, duration *big.Int, secret [32]byte, resolver common.Address, data [][]byte, reverseRecord bool, ownerControlledFuses uint16) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Register(&_ETHRegistrarControllerV2.TransactOpts, name, owner, duration, secret, resolver, data, reverseRecord, ownerControlledFuses)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) Renew(opts *bind.TransactOpts, name string, duration *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "renew", name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Renew(&_ETHRegistrarControllerV2.TransactOpts, name, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 duration) payable returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) Renew(name string, duration *big.Int) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Renew(&_ETHRegistrarControllerV2.TransactOpts, name, duration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.RenounceOwnership(&_ETHRegistrarControllerV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.RenounceOwnership(&_ETHRegistrarControllerV2.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.TransferOwnership(&_ETHRegistrarControllerV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.TransferOwnership(&_ETHRegistrarControllerV2.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Session) Withdraw() (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Withdraw(&_ETHRegistrarControllerV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _ETHRegistrarControllerV2.Contract.Withdraw(&_ETHRegistrarControllerV2.TransactOpts)
}

// ETHRegistrarControllerV2NameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2NameRegisteredIterator struct {
	Event *ETHRegistrarControllerV2NameRegistered // Event containing the contract specifics and raw log

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
func (it *ETHRegistrarControllerV2NameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRegistrarControllerV2NameRegistered)
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
		it.Event = new(ETHRegistrarControllerV2NameRegistered)
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
func (it *ETHRegistrarControllerV2NameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRegistrarControllerV2NameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRegistrarControllerV2NameRegistered represents a NameRegistered event raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2NameRegistered struct {
	Name     string
	Label    [32]byte
	Owner    common.Address
	BaseCost *big.Int
	Premium  *big.Int
	Expires  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*ETHRegistrarControllerV2NameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2NameRegisteredIterator{contract: _ETHRegistrarControllerV2.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *ETHRegistrarControllerV2NameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRegistrarControllerV2NameRegistered)
				if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0x69e37f151eb98a09618ddaa80c8cfaf1ce5996867c489f45b555b412271ebf27.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 baseCost, uint256 premium, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) ParseNameRegistered(log types.Log) (*ETHRegistrarControllerV2NameRegistered, error) {
	event := new(ETHRegistrarControllerV2NameRegistered)
	if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHRegistrarControllerV2NameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2NameRenewedIterator struct {
	Event *ETHRegistrarControllerV2NameRenewed // Event containing the contract specifics and raw log

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
func (it *ETHRegistrarControllerV2NameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRegistrarControllerV2NameRenewed)
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
		it.Event = new(ETHRegistrarControllerV2NameRenewed)
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
func (it *ETHRegistrarControllerV2NameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRegistrarControllerV2NameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRegistrarControllerV2NameRenewed represents a NameRenewed event raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2NameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*ETHRegistrarControllerV2NameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2NameRenewedIterator{contract: _ETHRegistrarControllerV2.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *ETHRegistrarControllerV2NameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRegistrarControllerV2NameRenewed)
				if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) ParseNameRenewed(log types.Log) (*ETHRegistrarControllerV2NameRenewed, error) {
	event := new(ETHRegistrarControllerV2NameRenewed)
	if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHRegistrarControllerV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2OwnershipTransferredIterator struct {
	Event *ETHRegistrarControllerV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ETHRegistrarControllerV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHRegistrarControllerV2OwnershipTransferred)
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
		it.Event = new(ETHRegistrarControllerV2OwnershipTransferred)
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
func (it *ETHRegistrarControllerV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHRegistrarControllerV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHRegistrarControllerV2OwnershipTransferred represents a OwnershipTransferred event raised by the ETHRegistrarControllerV2 contract.
type ETHRegistrarControllerV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ETHRegistrarControllerV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ETHRegistrarControllerV2OwnershipTransferredIterator{contract: _ETHRegistrarControllerV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ETHRegistrarControllerV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ETHRegistrarControllerV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHRegistrarControllerV2OwnershipTransferred)
				if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ETHRegistrarControllerV2 *ETHRegistrarControllerV2Filterer) ParseOwnershipTransferred(log types.Log) (*ETHRegistrarControllerV2OwnershipTransferred, error) {
	event := new(ETHRegistrarControllerV2OwnershipTransferred)
	if err := _ETHRegistrarControllerV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
