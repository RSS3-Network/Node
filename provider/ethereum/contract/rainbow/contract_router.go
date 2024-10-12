// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rainbow

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

// PermitHelperPermit is an auto generated low-level Go binding around an user-defined struct.
type PermitHelperPermit struct {
	Value            *big.Int
	Nonce            *big.Int
	Deadline         *big.Int
	IsDaiStylePermit bool
	V                uint8
	R                [32]byte
	S                [32]byte
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EthWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"}],\"name\":\"OwnerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"SwapTargetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"SwapTargetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"fillQuoteEthToToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercentageBasisPoints\",\"type\":\"uint256\"}],\"name\":\"fillQuoteTokenToEth\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercentageBasisPoints\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDaiStylePermit\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPermitHelper.Permit\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"fillQuoteTokenToEthWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"fillQuoteTokenToToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sellTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyTokenAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"swapCallData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDaiStylePermit\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structPermitHelper.Permit\",\"name\":\"permitData\",\"type\":\"tuple\"}],\"name\":\"fillQuoteTokenToTokenWithPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapTargets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"add\",\"type\":\"bool\"}],\"name\":\"updateSwapTargets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawEth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// SwapTargets is a free data retrieval call binding the contract method 0x83c4a19d.
//
// Solidity: function swapTargets(address ) view returns(bool)
func (_Router *RouterCaller) SwapTargets(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "swapTargets", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapTargets is a free data retrieval call binding the contract method 0x83c4a19d.
//
// Solidity: function swapTargets(address ) view returns(bool)
func (_Router *RouterSession) SwapTargets(arg0 common.Address) (bool, error) {
	return _Router.Contract.SwapTargets(&_Router.CallOpts, arg0)
}

// SwapTargets is a free data retrieval call binding the contract method 0x83c4a19d.
//
// Solidity: function swapTargets(address ) view returns(bool)
func (_Router *RouterCallerSession) SwapTargets(arg0 common.Address) (bool, error) {
	return _Router.Contract.SwapTargets(&_Router.CallOpts, arg0)
}

// FillQuoteEthToToken is a paid mutator transaction binding the contract method 0x3c2b9a7d.
//
// Solidity: function fillQuoteEthToToken(address buyTokenAddress, address target, bytes swapCallData, uint256 feeAmount) payable returns()
func (_Router *RouterTransactor) FillQuoteEthToToken(opts *bind.TransactOpts, buyTokenAddress common.Address, target common.Address, swapCallData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "fillQuoteEthToToken", buyTokenAddress, target, swapCallData, feeAmount)
}

// FillQuoteEthToToken is a paid mutator transaction binding the contract method 0x3c2b9a7d.
//
// Solidity: function fillQuoteEthToToken(address buyTokenAddress, address target, bytes swapCallData, uint256 feeAmount) payable returns()
func (_Router *RouterSession) FillQuoteEthToToken(buyTokenAddress common.Address, target common.Address, swapCallData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteEthToToken(&_Router.TransactOpts, buyTokenAddress, target, swapCallData, feeAmount)
}

// FillQuoteEthToToken is a paid mutator transaction binding the contract method 0x3c2b9a7d.
//
// Solidity: function fillQuoteEthToToken(address buyTokenAddress, address target, bytes swapCallData, uint256 feeAmount) payable returns()
func (_Router *RouterTransactorSession) FillQuoteEthToToken(buyTokenAddress common.Address, target common.Address, swapCallData []byte, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteEthToToken(&_Router.TransactOpts, buyTokenAddress, target, swapCallData, feeAmount)
}

// FillQuoteTokenToEth is a paid mutator transaction binding the contract method 0x999b6464.
//
// Solidity: function fillQuoteTokenToEth(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints) payable returns()
func (_Router *RouterTransactor) FillQuoteTokenToEth(opts *bind.TransactOpts, sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "fillQuoteTokenToEth", sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints)
}

// FillQuoteTokenToEth is a paid mutator transaction binding the contract method 0x999b6464.
//
// Solidity: function fillQuoteTokenToEth(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints) payable returns()
func (_Router *RouterSession) FillQuoteTokenToEth(sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToEth(&_Router.TransactOpts, sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints)
}

// FillQuoteTokenToEth is a paid mutator transaction binding the contract method 0x999b6464.
//
// Solidity: function fillQuoteTokenToEth(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints) payable returns()
func (_Router *RouterTransactorSession) FillQuoteTokenToEth(sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToEth(&_Router.TransactOpts, sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints)
}

// FillQuoteTokenToEthWithPermit is a paid mutator transaction binding the contract method 0xb3093838.
//
// Solidity: function fillQuoteTokenToEthWithPermit(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterTransactor) FillQuoteTokenToEthWithPermit(opts *bind.TransactOpts, sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "fillQuoteTokenToEthWithPermit", sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints, permitData)
}

// FillQuoteTokenToEthWithPermit is a paid mutator transaction binding the contract method 0xb3093838.
//
// Solidity: function fillQuoteTokenToEthWithPermit(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterSession) FillQuoteTokenToEthWithPermit(sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToEthWithPermit(&_Router.TransactOpts, sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints, permitData)
}

// FillQuoteTokenToEthWithPermit is a paid mutator transaction binding the contract method 0xb3093838.
//
// Solidity: function fillQuoteTokenToEthWithPermit(address sellTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feePercentageBasisPoints, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterTransactorSession) FillQuoteTokenToEthWithPermit(sellTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feePercentageBasisPoints *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToEthWithPermit(&_Router.TransactOpts, sellTokenAddress, target, swapCallData, sellAmount, feePercentageBasisPoints, permitData)
}

// FillQuoteTokenToToken is a paid mutator transaction binding the contract method 0x55e4b7be.
//
// Solidity: function fillQuoteTokenToToken(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount) payable returns()
func (_Router *RouterTransactor) FillQuoteTokenToToken(opts *bind.TransactOpts, sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "fillQuoteTokenToToken", sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount)
}

// FillQuoteTokenToToken is a paid mutator transaction binding the contract method 0x55e4b7be.
//
// Solidity: function fillQuoteTokenToToken(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount) payable returns()
func (_Router *RouterSession) FillQuoteTokenToToken(sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToToken(&_Router.TransactOpts, sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount)
}

// FillQuoteTokenToToken is a paid mutator transaction binding the contract method 0x55e4b7be.
//
// Solidity: function fillQuoteTokenToToken(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount) payable returns()
func (_Router *RouterTransactorSession) FillQuoteTokenToToken(sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToToken(&_Router.TransactOpts, sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount)
}

// FillQuoteTokenToTokenWithPermit is a paid mutator transaction binding the contract method 0xb0480bbd.
//
// Solidity: function fillQuoteTokenToTokenWithPermit(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterTransactor) FillQuoteTokenToTokenWithPermit(opts *bind.TransactOpts, sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "fillQuoteTokenToTokenWithPermit", sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount, permitData)
}

// FillQuoteTokenToTokenWithPermit is a paid mutator transaction binding the contract method 0xb0480bbd.
//
// Solidity: function fillQuoteTokenToTokenWithPermit(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterSession) FillQuoteTokenToTokenWithPermit(sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToTokenWithPermit(&_Router.TransactOpts, sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount, permitData)
}

// FillQuoteTokenToTokenWithPermit is a paid mutator transaction binding the contract method 0xb0480bbd.
//
// Solidity: function fillQuoteTokenToTokenWithPermit(address sellTokenAddress, address buyTokenAddress, address target, bytes swapCallData, uint256 sellAmount, uint256 feeAmount, (uint256,uint256,uint256,bool,uint8,bytes32,bytes32) permitData) payable returns()
func (_Router *RouterTransactorSession) FillQuoteTokenToTokenWithPermit(sellTokenAddress common.Address, buyTokenAddress common.Address, target common.Address, swapCallData []byte, sellAmount *big.Int, feeAmount *big.Int, permitData PermitHelperPermit) (*types.Transaction, error) {
	return _Router.Contract.FillQuoteTokenToTokenWithPermit(&_Router.TransactOpts, sellTokenAddress, buyTokenAddress, target, swapCallData, sellAmount, feeAmount, permitData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// UpdateSwapTargets is a paid mutator transaction binding the contract method 0x97bbda0e.
//
// Solidity: function updateSwapTargets(address target, bool add) returns()
func (_Router *RouterTransactor) UpdateSwapTargets(opts *bind.TransactOpts, target common.Address, add bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "updateSwapTargets", target, add)
}

// UpdateSwapTargets is a paid mutator transaction binding the contract method 0x97bbda0e.
//
// Solidity: function updateSwapTargets(address target, bool add) returns()
func (_Router *RouterSession) UpdateSwapTargets(target common.Address, add bool) (*types.Transaction, error) {
	return _Router.Contract.UpdateSwapTargets(&_Router.TransactOpts, target, add)
}

// UpdateSwapTargets is a paid mutator transaction binding the contract method 0x97bbda0e.
//
// Solidity: function updateSwapTargets(address target, bool add) returns()
func (_Router *RouterTransactorSession) UpdateSwapTargets(target common.Address, add bool) (*types.Transaction, error) {
	return _Router.Contract.UpdateSwapTargets(&_Router.TransactOpts, target, add)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_Router *RouterTransactor) WithdrawEth(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawEth", to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_Router *RouterSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawEth(&_Router.TransactOpts, to, amount)
}

// WithdrawEth is a paid mutator transaction binding the contract method 0x1b9a91a4.
//
// Solidity: function withdrawEth(address to, uint256 amount) returns()
func (_Router *RouterTransactorSession) WithdrawEth(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawEth(&_Router.TransactOpts, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Router *RouterTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "withdrawToken", token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Router *RouterSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawToken(&_Router.TransactOpts, token, to, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address token, address to, uint256 amount) returns()
func (_Router *RouterTransactorSession) WithdrawToken(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.WithdrawToken(&_Router.TransactOpts, token, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// RouterEthWithdrawnIterator is returned from FilterEthWithdrawn and is used to iterate over the raw logs and unpacked data for EthWithdrawn events raised by the Router contract.
type RouterEthWithdrawnIterator struct {
	Event *RouterEthWithdrawn // Event containing the contract specifics and raw log

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
func (it *RouterEthWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterEthWithdrawn)
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
		it.Event = new(RouterEthWithdrawn)
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
func (it *RouterEthWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterEthWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterEthWithdrawn represents a EthWithdrawn event raised by the Router contract.
type RouterEthWithdrawn struct {
	Target common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEthWithdrawn is a free log retrieval operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed target, uint256 amount)
func (_Router *RouterFilterer) FilterEthWithdrawn(opts *bind.FilterOpts, target []common.Address) (*RouterEthWithdrawnIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "EthWithdrawn", targetRule)
	if err != nil {
		return nil, err
	}
	return &RouterEthWithdrawnIterator{contract: _Router.contract, event: "EthWithdrawn", logs: logs, sub: sub}, nil
}

// WatchEthWithdrawn is a free log subscription operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed target, uint256 amount)
func (_Router *RouterFilterer) WatchEthWithdrawn(opts *bind.WatchOpts, sink chan<- *RouterEthWithdrawn, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "EthWithdrawn", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterEthWithdrawn)
				if err := _Router.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
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

// ParseEthWithdrawn is a log parse operation binding the contract event 0x8455ae6be5d92f1df1c3c1484388e247a36c7e60d72055ae216dbc258f257d4b.
//
// Solidity: event EthWithdrawn(address indexed target, uint256 amount)
func (_Router *RouterFilterer) ParseEthWithdrawn(log types.Log) (*RouterEthWithdrawn, error) {
	event := new(RouterEthWithdrawn)
	if err := _Router.contract.UnpackLog(event, "EthWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnerChangedIterator is returned from FilterOwnerChanged and is used to iterate over the raw logs and unpacked data for OwnerChanged events raised by the Router contract.
type RouterOwnerChangedIterator struct {
	Event *RouterOwnerChanged // Event containing the contract specifics and raw log

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
func (it *RouterOwnerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnerChanged)
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
		it.Event = new(RouterOwnerChanged)
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
func (it *RouterOwnerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnerChanged represents a OwnerChanged event raised by the Router contract.
type RouterOwnerChanged struct {
	NewOwner common.Address
	OldOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerChanged is a free log retrieval operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed newOwner, address indexed oldOwner)
func (_Router *RouterFilterer) FilterOwnerChanged(opts *bind.FilterOpts, newOwner []common.Address, oldOwner []common.Address) (*RouterOwnerChangedIterator, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnerChanged", newOwnerRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnerChangedIterator{contract: _Router.contract, event: "OwnerChanged", logs: logs, sub: sub}, nil
}

// WatchOwnerChanged is a free log subscription operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed newOwner, address indexed oldOwner)
func (_Router *RouterFilterer) WatchOwnerChanged(opts *bind.WatchOpts, sink chan<- *RouterOwnerChanged, newOwner []common.Address, oldOwner []common.Address) (event.Subscription, error) {

	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}
	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnerChanged", newOwnerRule, oldOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnerChanged)
				if err := _Router.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
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

// ParseOwnerChanged is a log parse operation binding the contract event 0xb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c.
//
// Solidity: event OwnerChanged(address indexed newOwner, address indexed oldOwner)
func (_Router *RouterFilterer) ParseOwnerChanged(log types.Log) (*RouterOwnerChanged, error) {
	event := new(RouterOwnerChanged)
	if err := _Router.contract.UnpackLog(event, "OwnerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterSwapTargetAddedIterator is returned from FilterSwapTargetAdded and is used to iterate over the raw logs and unpacked data for SwapTargetAdded events raised by the Router contract.
type RouterSwapTargetAddedIterator struct {
	Event *RouterSwapTargetAdded // Event containing the contract specifics and raw log

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
func (it *RouterSwapTargetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterSwapTargetAdded)
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
		it.Event = new(RouterSwapTargetAdded)
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
func (it *RouterSwapTargetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterSwapTargetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterSwapTargetAdded represents a SwapTargetAdded event raised by the Router contract.
type RouterSwapTargetAdded struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapTargetAdded is a free log retrieval operation binding the contract event 0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf.
//
// Solidity: event SwapTargetAdded(address indexed target)
func (_Router *RouterFilterer) FilterSwapTargetAdded(opts *bind.FilterOpts, target []common.Address) (*RouterSwapTargetAddedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "SwapTargetAdded", targetRule)
	if err != nil {
		return nil, err
	}
	return &RouterSwapTargetAddedIterator{contract: _Router.contract, event: "SwapTargetAdded", logs: logs, sub: sub}, nil
}

// WatchSwapTargetAdded is a free log subscription operation binding the contract event 0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf.
//
// Solidity: event SwapTargetAdded(address indexed target)
func (_Router *RouterFilterer) WatchSwapTargetAdded(opts *bind.WatchOpts, sink chan<- *RouterSwapTargetAdded, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "SwapTargetAdded", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterSwapTargetAdded)
				if err := _Router.contract.UnpackLog(event, "SwapTargetAdded", log); err != nil {
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

// ParseSwapTargetAdded is a log parse operation binding the contract event 0xb907822409611d127ab6a64611591b98e03a6a85ade4f258bae26b7c1efdfeaf.
//
// Solidity: event SwapTargetAdded(address indexed target)
func (_Router *RouterFilterer) ParseSwapTargetAdded(log types.Log) (*RouterSwapTargetAdded, error) {
	event := new(RouterSwapTargetAdded)
	if err := _Router.contract.UnpackLog(event, "SwapTargetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterSwapTargetRemovedIterator is returned from FilterSwapTargetRemoved and is used to iterate over the raw logs and unpacked data for SwapTargetRemoved events raised by the Router contract.
type RouterSwapTargetRemovedIterator struct {
	Event *RouterSwapTargetRemoved // Event containing the contract specifics and raw log

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
func (it *RouterSwapTargetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterSwapTargetRemoved)
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
		it.Event = new(RouterSwapTargetRemoved)
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
func (it *RouterSwapTargetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterSwapTargetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterSwapTargetRemoved represents a SwapTargetRemoved event raised by the Router contract.
type RouterSwapTargetRemoved struct {
	Target common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapTargetRemoved is a free log retrieval operation binding the contract event 0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c.
//
// Solidity: event SwapTargetRemoved(address indexed target)
func (_Router *RouterFilterer) FilterSwapTargetRemoved(opts *bind.FilterOpts, target []common.Address) (*RouterSwapTargetRemovedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "SwapTargetRemoved", targetRule)
	if err != nil {
		return nil, err
	}
	return &RouterSwapTargetRemovedIterator{contract: _Router.contract, event: "SwapTargetRemoved", logs: logs, sub: sub}, nil
}

// WatchSwapTargetRemoved is a free log subscription operation binding the contract event 0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c.
//
// Solidity: event SwapTargetRemoved(address indexed target)
func (_Router *RouterFilterer) WatchSwapTargetRemoved(opts *bind.WatchOpts, sink chan<- *RouterSwapTargetRemoved, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "SwapTargetRemoved", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterSwapTargetRemoved)
				if err := _Router.contract.UnpackLog(event, "SwapTargetRemoved", log); err != nil {
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

// ParseSwapTargetRemoved is a log parse operation binding the contract event 0x393b8be3e26787f19285ecd039dfd80bc6507828750f4d50367e6efe2524695c.
//
// Solidity: event SwapTargetRemoved(address indexed target)
func (_Router *RouterFilterer) ParseSwapTargetRemoved(log types.Log) (*RouterSwapTargetRemoved, error) {
	event := new(RouterSwapTargetRemoved)
	if err := _Router.contract.UnpackLog(event, "SwapTargetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the Router contract.
type RouterTokenWithdrawnIterator struct {
	Event *RouterTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *RouterTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterTokenWithdrawn)
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
		it.Event = new(RouterTokenWithdrawn)
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
func (it *RouterTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterTokenWithdrawn represents a TokenWithdrawn event raised by the Router contract.
type RouterTokenWithdrawn struct {
	Token  common.Address
	Target common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed target, uint256 amount)
func (_Router *RouterFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, token []common.Address, target []common.Address) (*RouterTokenWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "TokenWithdrawn", tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &RouterTokenWithdrawnIterator{contract: _Router.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed target, uint256 amount)
func (_Router *RouterFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *RouterTokenWithdrawn, token []common.Address, target []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "TokenWithdrawn", tokenRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterTokenWithdrawn)
				if err := _Router.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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

// ParseTokenWithdrawn is a log parse operation binding the contract event 0x8210728e7c071f615b840ee026032693858fbcd5e5359e67e438c890f59e5620.
//
// Solidity: event TokenWithdrawn(address indexed token, address indexed target, uint256 amount)
func (_Router *RouterFilterer) ParseTokenWithdrawn(log types.Log) (*RouterTokenWithdrawn, error) {
	event := new(RouterTokenWithdrawn)
	if err := _Router.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
