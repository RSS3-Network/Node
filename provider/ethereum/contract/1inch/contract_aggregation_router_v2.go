// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinch

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

// IOneInchCallerCallDescription is an auto generated low-level Go binding around an user-defined struct.
type IOneInchCallerCallDescription struct {
	TargetWithMandatory *big.Int
	GasLimit            *big.Int
	Value               *big.Int
	Data                []byte
}

// OneInchExchangeSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type OneInchExchangeSwapDescription struct {
	SrcToken         common.Address
	DstToken         common.Address
	SrcReceiver      common.Address
	DstReceiver      common.Address
	Amount           *big.Int
	MinReturnAmount  *big.Int
	GuaranteedAmount *big.Int
	Flags            *big.Int
	Referrer         common.Address
	Permit           []byte
}

// AggregationRouterV2MetaData contains all meta data concerning the AggregationRouterV2 contract.
var AggregationRouterV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIOneInchCaller\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOneInchExchange.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetWithMandatory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIOneInchCaller.CallDescription[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"discountedSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOneInchCaller\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structOneInchExchange.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetWithMandatory\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIOneInchCaller.CallDescription[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AggregationRouterV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregationRouterV2MetaData.ABI instead.
var AggregationRouterV2ABI = AggregationRouterV2MetaData.ABI

// AggregationRouterV2 is an auto generated Go binding around an Ethereum contract.
type AggregationRouterV2 struct {
	AggregationRouterV2Caller     // Read-only binding to the contract
	AggregationRouterV2Transactor // Write-only binding to the contract
	AggregationRouterV2Filterer   // Log filterer for contract events
}

// AggregationRouterV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AggregationRouterV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregationRouterV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregationRouterV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregationRouterV2Session struct {
	Contract     *AggregationRouterV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AggregationRouterV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregationRouterV2CallerSession struct {
	Contract *AggregationRouterV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AggregationRouterV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregationRouterV2TransactorSession struct {
	Contract     *AggregationRouterV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AggregationRouterV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AggregationRouterV2Raw struct {
	Contract *AggregationRouterV2 // Generic contract binding to access the raw methods on
}

// AggregationRouterV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregationRouterV2CallerRaw struct {
	Contract *AggregationRouterV2Caller // Generic read-only contract binding to access the raw methods on
}

// AggregationRouterV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregationRouterV2TransactorRaw struct {
	Contract *AggregationRouterV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregationRouterV2 creates a new instance of AggregationRouterV2, bound to a specific deployed contract.
func NewAggregationRouterV2(address common.Address, backend bind.ContractBackend) (*AggregationRouterV2, error) {
	contract, err := bindAggregationRouterV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2{AggregationRouterV2Caller: AggregationRouterV2Caller{contract: contract}, AggregationRouterV2Transactor: AggregationRouterV2Transactor{contract: contract}, AggregationRouterV2Filterer: AggregationRouterV2Filterer{contract: contract}}, nil
}

// NewAggregationRouterV2Caller creates a new read-only instance of AggregationRouterV2, bound to a specific deployed contract.
func NewAggregationRouterV2Caller(address common.Address, caller bind.ContractCaller) (*AggregationRouterV2Caller, error) {
	contract, err := bindAggregationRouterV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2Caller{contract: contract}, nil
}

// NewAggregationRouterV2Transactor creates a new write-only instance of AggregationRouterV2, bound to a specific deployed contract.
func NewAggregationRouterV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AggregationRouterV2Transactor, error) {
	contract, err := bindAggregationRouterV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2Transactor{contract: contract}, nil
}

// NewAggregationRouterV2Filterer creates a new log filterer instance of AggregationRouterV2, bound to a specific deployed contract.
func NewAggregationRouterV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AggregationRouterV2Filterer, error) {
	contract, err := bindAggregationRouterV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2Filterer{contract: contract}, nil
}

// bindAggregationRouterV2 binds a generic wrapper to an already deployed contract.
func bindAggregationRouterV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregationRouterV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV2 *AggregationRouterV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV2.Contract.AggregationRouterV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV2 *AggregationRouterV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.AggregationRouterV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV2 *AggregationRouterV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.AggregationRouterV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV2 *AggregationRouterV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV2 *AggregationRouterV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV2 *AggregationRouterV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV2 *AggregationRouterV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregationRouterV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV2 *AggregationRouterV2Session) Owner() (common.Address, error) {
	return _AggregationRouterV2.Contract.Owner(&_AggregationRouterV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV2 *AggregationRouterV2CallerSession) Owner() (common.Address, error) {
	return _AggregationRouterV2.Contract.Owner(&_AggregationRouterV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AggregationRouterV2 *AggregationRouterV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AggregationRouterV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AggregationRouterV2 *AggregationRouterV2Session) Paused() (bool, error) {
	return _AggregationRouterV2.Contract.Paused(&_AggregationRouterV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AggregationRouterV2 *AggregationRouterV2CallerSession) Paused() (bool, error) {
	return _AggregationRouterV2.Contract.Paused(&_AggregationRouterV2.CallOpts)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x34b0793b.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2Transactor) DiscountedSwap(opts *bind.TransactOpts, caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "discountedSwap", caller, desc, calls)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x34b0793b.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2Session) DiscountedSwap(caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.DiscountedSwap(&_AggregationRouterV2.TransactOpts, caller, desc, calls)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x34b0793b.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) DiscountedSwap(caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.DiscountedSwap(&_AggregationRouterV2.TransactOpts, caller, desc, calls)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AggregationRouterV2 *AggregationRouterV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AggregationRouterV2 *AggregationRouterV2Session) Pause() (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.Pause(&_AggregationRouterV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) Pause() (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.Pause(&_AggregationRouterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV2 *AggregationRouterV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV2 *AggregationRouterV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.RenounceOwnership(&_AggregationRouterV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.RenounceOwnership(&_AggregationRouterV2.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV2 *AggregationRouterV2Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV2 *AggregationRouterV2Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.RescueFunds(&_AggregationRouterV2.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.RescueFunds(&_AggregationRouterV2.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2Transactor) Swap(opts *bind.TransactOpts, caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "swap", caller, desc, calls)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2Session) Swap(caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.Swap(&_AggregationRouterV2.TransactOpts, caller, desc, calls)
}

// Swap is a paid mutator transaction binding the contract method 0x90411a32.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,uint256,address,bytes) desc, (uint256,uint256,uint256,bytes)[] calls) payable returns(uint256 returnAmount)
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) Swap(caller common.Address, desc OneInchExchangeSwapDescription, calls []IOneInchCallerCallDescription) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.Swap(&_AggregationRouterV2.TransactOpts, caller, desc, calls)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV2 *AggregationRouterV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV2 *AggregationRouterV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.TransferOwnership(&_AggregationRouterV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV2 *AggregationRouterV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV2.Contract.TransferOwnership(&_AggregationRouterV2.TransactOpts, newOwner)
}

// AggregationRouterV2ErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the AggregationRouterV2 contract.
type AggregationRouterV2ErrorIterator struct {
	Event *AggregationRouterV2Error // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV2ErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV2Error)
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
		it.Event = new(AggregationRouterV2Error)
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
func (it *AggregationRouterV2ErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV2ErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV2Error represents a Error event raised by the AggregationRouterV2 contract.
type AggregationRouterV2Error struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) FilterError(opts *bind.FilterOpts) (*AggregationRouterV2ErrorIterator, error) {

	logs, sub, err := _AggregationRouterV2.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2ErrorIterator{contract: _AggregationRouterV2.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) WatchError(opts *bind.WatchOpts, sink chan<- *AggregationRouterV2Error) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV2.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV2Error)
				if err := _AggregationRouterV2.contract.UnpackLog(event, "Error", log); err != nil {
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

// ParseError is a log parse operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) ParseError(log types.Log) (*AggregationRouterV2Error, error) {
	event := new(AggregationRouterV2Error)
	if err := _AggregationRouterV2.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AggregationRouterV2 contract.
type AggregationRouterV2OwnershipTransferredIterator struct {
	Event *AggregationRouterV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV2OwnershipTransferred)
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
		it.Event = new(AggregationRouterV2OwnershipTransferred)
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
func (it *AggregationRouterV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV2OwnershipTransferred represents a OwnershipTransferred event raised by the AggregationRouterV2 contract.
type AggregationRouterV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregationRouterV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2OwnershipTransferredIterator{contract: _AggregationRouterV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregationRouterV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV2OwnershipTransferred)
				if err := _AggregationRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AggregationRouterV2 *AggregationRouterV2Filterer) ParseOwnershipTransferred(log types.Log) (*AggregationRouterV2OwnershipTransferred, error) {
	event := new(AggregationRouterV2OwnershipTransferred)
	if err := _AggregationRouterV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AggregationRouterV2 contract.
type AggregationRouterV2PausedIterator struct {
	Event *AggregationRouterV2Paused // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV2Paused)
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
		it.Event = new(AggregationRouterV2Paused)
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
func (it *AggregationRouterV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV2Paused represents a Paused event raised by the AggregationRouterV2 contract.
type AggregationRouterV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) FilterPaused(opts *bind.FilterOpts) (*AggregationRouterV2PausedIterator, error) {

	logs, sub, err := _AggregationRouterV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2PausedIterator{contract: _AggregationRouterV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AggregationRouterV2Paused) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV2Paused)
				if err := _AggregationRouterV2.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AggregationRouterV2 *AggregationRouterV2Filterer) ParsePaused(log types.Log) (*AggregationRouterV2Paused, error) {
	event := new(AggregationRouterV2Paused)
	if err := _AggregationRouterV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV2SwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the AggregationRouterV2 contract.
type AggregationRouterV2SwappedIterator struct {
	Event *AggregationRouterV2Swapped // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV2SwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV2Swapped)
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
		it.Event = new(AggregationRouterV2Swapped)
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
func (it *AggregationRouterV2SwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV2SwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV2Swapped represents a Swapped event raised by the AggregationRouterV2 contract.
type AggregationRouterV2Swapped struct {
	Sender           common.Address
	SrcToken         common.Address
	DstToken         common.Address
	DstReceiver      common.Address
	Amount           *big.Int
	SpentAmount      *big.Int
	ReturnAmount     *big.Int
	MinReturnAmount  *big.Int
	GuaranteedAmount *big.Int
	Referrer         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) FilterSwapped(opts *bind.FilterOpts, sender []common.Address, srcToken []common.Address, dstToken []common.Address) (*AggregationRouterV2SwappedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _AggregationRouterV2.contract.FilterLogs(opts, "Swapped", senderRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2SwappedIterator{contract: _AggregationRouterV2.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *AggregationRouterV2Swapped, sender []common.Address, srcToken []common.Address, dstToken []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var dstTokenRule []interface{}
	for _, dstTokenItem := range dstToken {
		dstTokenRule = append(dstTokenRule, dstTokenItem)
	}

	logs, sub, err := _AggregationRouterV2.contract.WatchLogs(opts, "Swapped", senderRule, srcTokenRule, dstTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV2Swapped)
				if err := _AggregationRouterV2.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0x76af224a143865a50b41496e1a73622698692c565c1214bc862f18e22d829c5e.
//
// Solidity: event Swapped(address indexed sender, address indexed srcToken, address indexed dstToken, address dstReceiver, uint256 amount, uint256 spentAmount, uint256 returnAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) ParseSwapped(log types.Log) (*AggregationRouterV2Swapped, error) {
	event := new(AggregationRouterV2Swapped)
	if err := _AggregationRouterV2.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AggregationRouterV2 contract.
type AggregationRouterV2UnpausedIterator struct {
	Event *AggregationRouterV2Unpaused // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV2Unpaused)
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
		it.Event = new(AggregationRouterV2Unpaused)
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
func (it *AggregationRouterV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV2Unpaused represents a Unpaused event raised by the AggregationRouterV2 contract.
type AggregationRouterV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*AggregationRouterV2UnpausedIterator, error) {

	logs, sub, err := _AggregationRouterV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV2UnpausedIterator{contract: _AggregationRouterV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AggregationRouterV2 *AggregationRouterV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AggregationRouterV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV2Unpaused)
				if err := _AggregationRouterV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AggregationRouterV2 *AggregationRouterV2Filterer) ParseUnpaused(log types.Log) (*AggregationRouterV2Unpaused, error) {
	event := new(AggregationRouterV2Unpaused)
	if err := _AggregationRouterV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
