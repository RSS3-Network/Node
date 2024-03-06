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

// AggregationRouterV3SwapDescription is an auto generated low-level Go binding around an user-defined struct.
type AggregationRouterV3SwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// AggregationRouterV3MetaData contains all meta data concerning the AggregationRouterV3 contract.
var AggregationRouterV3MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"discountedSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chiSpent\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV3.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AggregationRouterV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregationRouterV3MetaData.ABI instead.
var AggregationRouterV3ABI = AggregationRouterV3MetaData.ABI

// AggregationRouterV3 is an auto generated Go binding around an Ethereum contract.
type AggregationRouterV3 struct {
	AggregationRouterV3Caller     // Read-only binding to the contract
	AggregationRouterV3Transactor // Write-only binding to the contract
	AggregationRouterV3Filterer   // Log filterer for contract events
}

// AggregationRouterV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AggregationRouterV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregationRouterV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregationRouterV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregationRouterV3Session struct {
	Contract     *AggregationRouterV3 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AggregationRouterV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregationRouterV3CallerSession struct {
	Contract *AggregationRouterV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AggregationRouterV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregationRouterV3TransactorSession struct {
	Contract     *AggregationRouterV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AggregationRouterV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AggregationRouterV3Raw struct {
	Contract *AggregationRouterV3 // Generic contract binding to access the raw methods on
}

// AggregationRouterV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregationRouterV3CallerRaw struct {
	Contract *AggregationRouterV3Caller // Generic read-only contract binding to access the raw methods on
}

// AggregationRouterV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregationRouterV3TransactorRaw struct {
	Contract *AggregationRouterV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregationRouterV3 creates a new instance of AggregationRouterV3, bound to a specific deployed contract.
func NewAggregationRouterV3(address common.Address, backend bind.ContractBackend) (*AggregationRouterV3, error) {
	contract, err := bindAggregationRouterV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3{AggregationRouterV3Caller: AggregationRouterV3Caller{contract: contract}, AggregationRouterV3Transactor: AggregationRouterV3Transactor{contract: contract}, AggregationRouterV3Filterer: AggregationRouterV3Filterer{contract: contract}}, nil
}

// NewAggregationRouterV3Caller creates a new read-only instance of AggregationRouterV3, bound to a specific deployed contract.
func NewAggregationRouterV3Caller(address common.Address, caller bind.ContractCaller) (*AggregationRouterV3Caller, error) {
	contract, err := bindAggregationRouterV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3Caller{contract: contract}, nil
}

// NewAggregationRouterV3Transactor creates a new write-only instance of AggregationRouterV3, bound to a specific deployed contract.
func NewAggregationRouterV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AggregationRouterV3Transactor, error) {
	contract, err := bindAggregationRouterV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3Transactor{contract: contract}, nil
}

// NewAggregationRouterV3Filterer creates a new log filterer instance of AggregationRouterV3, bound to a specific deployed contract.
func NewAggregationRouterV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AggregationRouterV3Filterer, error) {
	contract, err := bindAggregationRouterV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3Filterer{contract: contract}, nil
}

// bindAggregationRouterV3 binds a generic wrapper to an already deployed contract.
func bindAggregationRouterV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregationRouterV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV3 *AggregationRouterV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV3.Contract.AggregationRouterV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV3 *AggregationRouterV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.AggregationRouterV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV3 *AggregationRouterV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.AggregationRouterV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV3 *AggregationRouterV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV3 *AggregationRouterV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV3 *AggregationRouterV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV3 *AggregationRouterV3Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregationRouterV3.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV3 *AggregationRouterV3Session) Owner() (common.Address, error) {
	return _AggregationRouterV3.Contract.Owner(&_AggregationRouterV3.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV3 *AggregationRouterV3CallerSession) Owner() (common.Address, error) {
	return _AggregationRouterV3.Contract.Owner(&_AggregationRouterV3.CallOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV3 *AggregationRouterV3Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV3 *AggregationRouterV3Session) Destroy() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Destroy(&_AggregationRouterV3.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) Destroy() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Destroy(&_AggregationRouterV3.TransactOpts)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_AggregationRouterV3 *AggregationRouterV3Transactor) DiscountedSwap(opts *bind.TransactOpts, caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "discountedSwap", caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_AggregationRouterV3 *AggregationRouterV3Session) DiscountedSwap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.DiscountedSwap(&_AggregationRouterV3.TransactOpts, caller, desc, data)
}

// DiscountedSwap is a paid mutator transaction binding the contract method 0x6c4a483e.
//
// Solidity: function discountedSwap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft, uint256 chiSpent)
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) DiscountedSwap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.DiscountedSwap(&_AggregationRouterV3.TransactOpts, caller, desc, data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV3 *AggregationRouterV3Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV3 *AggregationRouterV3Session) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.RenounceOwnership(&_AggregationRouterV3.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.RenounceOwnership(&_AggregationRouterV3.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV3 *AggregationRouterV3Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV3 *AggregationRouterV3Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.RescueFunds(&_AggregationRouterV3.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.RescueFunds(&_AggregationRouterV3.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_AggregationRouterV3 *AggregationRouterV3Transactor) Swap(opts *bind.TransactOpts, caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "swap", caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_AggregationRouterV3 *AggregationRouterV3Session) Swap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Swap(&_AggregationRouterV3.TransactOpts, caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 gasLeft)
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) Swap(caller common.Address, desc AggregationRouterV3SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Swap(&_AggregationRouterV3.TransactOpts, caller, desc, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV3 *AggregationRouterV3Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV3 *AggregationRouterV3Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.TransferOwnership(&_AggregationRouterV3.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.TransferOwnership(&_AggregationRouterV3.TransactOpts, newOwner)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Transactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Session) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Unoswap(&_AggregationRouterV3.TransactOpts, srcToken, amount, minReturn, arg3)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] ) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, arg3 [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Unoswap(&_AggregationRouterV3.TransactOpts, srcToken, amount, minReturn, arg3)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Transactor) UnoswapWithPermit(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.Transact(opts, "unoswapWithPermit", srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Session) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.UnoswapWithPermit(&_AggregationRouterV3.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) payable returns(uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.UnoswapWithPermit(&_AggregationRouterV3.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV3 *AggregationRouterV3Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV3.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV3 *AggregationRouterV3Session) Receive() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Receive(&_AggregationRouterV3.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV3 *AggregationRouterV3TransactorSession) Receive() (*types.Transaction, error) {
	return _AggregationRouterV3.Contract.Receive(&_AggregationRouterV3.TransactOpts)
}

// AggregationRouterV3ErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the AggregationRouterV3 contract.
type AggregationRouterV3ErrorIterator struct {
	Event *AggregationRouterV3Error // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV3ErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV3Error)
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
		it.Event = new(AggregationRouterV3Error)
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
func (it *AggregationRouterV3ErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV3ErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV3Error represents a Error event raised by the AggregationRouterV3 contract.
type AggregationRouterV3Error struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) FilterError(opts *bind.FilterOpts) (*AggregationRouterV3ErrorIterator, error) {

	logs, sub, err := _AggregationRouterV3.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3ErrorIterator{contract: _AggregationRouterV3.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string reason)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) WatchError(opts *bind.WatchOpts, sink chan<- *AggregationRouterV3Error) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV3.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV3Error)
				if err := _AggregationRouterV3.contract.UnpackLog(event, "Error", log); err != nil {
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
func (_AggregationRouterV3 *AggregationRouterV3Filterer) ParseError(log types.Log) (*AggregationRouterV3Error, error) {
	event := new(AggregationRouterV3Error)
	if err := _AggregationRouterV3.contract.UnpackLog(event, "Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV3OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AggregationRouterV3 contract.
type AggregationRouterV3OwnershipTransferredIterator struct {
	Event *AggregationRouterV3OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV3OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV3OwnershipTransferred)
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
		it.Event = new(AggregationRouterV3OwnershipTransferred)
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
func (it *AggregationRouterV3OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV3OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV3OwnershipTransferred represents a OwnershipTransferred event raised by the AggregationRouterV3 contract.
type AggregationRouterV3OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregationRouterV3OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV3.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3OwnershipTransferredIterator{contract: _AggregationRouterV3.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregationRouterV3OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV3.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV3OwnershipTransferred)
				if err := _AggregationRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AggregationRouterV3 *AggregationRouterV3Filterer) ParseOwnershipTransferred(log types.Log) (*AggregationRouterV3OwnershipTransferred, error) {
	event := new(AggregationRouterV3OwnershipTransferred)
	if err := _AggregationRouterV3.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV3SwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the AggregationRouterV3 contract.
type AggregationRouterV3SwappedIterator struct {
	Event *AggregationRouterV3Swapped // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV3SwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV3Swapped)
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
		it.Event = new(AggregationRouterV3Swapped)
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
func (it *AggregationRouterV3SwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV3SwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV3Swapped represents a Swapped event raised by the AggregationRouterV3 contract.
type AggregationRouterV3Swapped struct {
	Sender       common.Address
	SrcToken     common.Address
	DstToken     common.Address
	DstReceiver  common.Address
	SpentAmount  *big.Int
	ReturnAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) FilterSwapped(opts *bind.FilterOpts) (*AggregationRouterV3SwappedIterator, error) {

	logs, sub, err := _AggregationRouterV3.contract.FilterLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV3SwappedIterator{contract: _AggregationRouterV3.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *AggregationRouterV3Swapped) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV3.contract.WatchLogs(opts, "Swapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV3Swapped)
				if err := _AggregationRouterV3.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xd6d4f5681c246c9f42c203e287975af1601f8df8035a9251f79aab5c8f09e2f8.
//
// Solidity: event Swapped(address sender, address srcToken, address dstToken, address dstReceiver, uint256 spentAmount, uint256 returnAmount)
func (_AggregationRouterV3 *AggregationRouterV3Filterer) ParseSwapped(log types.Log) (*AggregationRouterV3Swapped, error) {
	event := new(AggregationRouterV3Swapped)
	if err := _AggregationRouterV3.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
