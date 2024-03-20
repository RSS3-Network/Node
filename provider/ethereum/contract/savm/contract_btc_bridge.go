// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package savm

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

// BTCBridgeMetaData contains all meta data concerning the BTCBridge contract.
var BTCBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"normalizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDepositFee\",\"type\":\"uint256\"}],\"name\":\"SetDepositFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeRecipient\",\"type\":\"address\"}],\"name\":\"SetFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"SetMinWithdrawAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newWithdrawFee\",\"type\":\"uint256\"}],\"name\":\"SetWithdrawFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"normalizedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"depositFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"getBridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"getTotalWithdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"newMinWithdrawAmount\",\"type\":\"uint128\"}],\"name\":\"setMinWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"recipient\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BTCBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BTCBridgeMetaData.ABI instead.
var BTCBridgeABI = BTCBridgeMetaData.ABI

// BTCBridge is an auto generated Go binding around an Ethereum contract.
type BTCBridge struct {
	BTCBridgeCaller     // Read-only binding to the contract
	BTCBridgeTransactor // Write-only binding to the contract
	BTCBridgeFilterer   // Log filterer for contract events
}

// BTCBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BTCBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BTCBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BTCBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BTCBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BTCBridgeSession struct {
	Contract     *BTCBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BTCBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BTCBridgeCallerSession struct {
	Contract *BTCBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BTCBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BTCBridgeTransactorSession struct {
	Contract     *BTCBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BTCBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BTCBridgeRaw struct {
	Contract *BTCBridge // Generic contract binding to access the raw methods on
}

// BTCBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BTCBridgeCallerRaw struct {
	Contract *BTCBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BTCBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BTCBridgeTransactorRaw struct {
	Contract *BTCBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBTCBridge creates a new instance of BTCBridge, bound to a specific deployed contract.
func NewBTCBridge(address common.Address, backend bind.ContractBackend) (*BTCBridge, error) {
	contract, err := bindBTCBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BTCBridge{BTCBridgeCaller: BTCBridgeCaller{contract: contract}, BTCBridgeTransactor: BTCBridgeTransactor{contract: contract}, BTCBridgeFilterer: BTCBridgeFilterer{contract: contract}}, nil
}

// NewBTCBridgeCaller creates a new read-only instance of BTCBridge, bound to a specific deployed contract.
func NewBTCBridgeCaller(address common.Address, caller bind.ContractCaller) (*BTCBridgeCaller, error) {
	contract, err := bindBTCBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BTCBridgeCaller{contract: contract}, nil
}

// NewBTCBridgeTransactor creates a new write-only instance of BTCBridge, bound to a specific deployed contract.
func NewBTCBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BTCBridgeTransactor, error) {
	contract, err := bindBTCBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BTCBridgeTransactor{contract: contract}, nil
}

// NewBTCBridgeFilterer creates a new log filterer instance of BTCBridge, bound to a specific deployed contract.
func NewBTCBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BTCBridgeFilterer, error) {
	contract, err := bindBTCBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BTCBridgeFilterer{contract: contract}, nil
}

// bindBTCBridge binds a generic wrapper to an already deployed contract.
func bindBTCBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BTCBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCBridge *BTCBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCBridge.Contract.BTCBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCBridge *BTCBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCBridge.Contract.BTCBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCBridge *BTCBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCBridge.Contract.BTCBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BTCBridge *BTCBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BTCBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BTCBridge *BTCBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BTCBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BTCBridge *BTCBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BTCBridge.Contract.contract.Transact(opts, method, params...)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BTCBridge *BTCBridgeCaller) DepositFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "depositFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BTCBridge *BTCBridgeSession) DepositFee() (*big.Int, error) {
	return _BTCBridge.Contract.DepositFee(&_BTCBridge.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_BTCBridge *BTCBridgeCallerSession) DepositFee() (*big.Int, error) {
	return _BTCBridge.Contract.DepositFee(&_BTCBridge.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BTCBridge *BTCBridgeCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BTCBridge *BTCBridgeSession) FeeRecipient() (common.Address, error) {
	return _BTCBridge.Contract.FeeRecipient(&_BTCBridge.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BTCBridge *BTCBridgeCallerSession) FeeRecipient() (common.Address, error) {
	return _BTCBridge.Contract.FeeRecipient(&_BTCBridge.CallOpts)
}

// GetBridgeFee is a free data retrieval call binding the contract method 0x5c12ba59.
//
// Solidity: function getBridgeFee(bytes payload) view returns(uint256 bridgeFee)
func (_BTCBridge *BTCBridgeCaller) GetBridgeFee(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "getBridgeFee", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBridgeFee is a free data retrieval call binding the contract method 0x5c12ba59.
//
// Solidity: function getBridgeFee(bytes payload) view returns(uint256 bridgeFee)
func (_BTCBridge *BTCBridgeSession) GetBridgeFee(payload []byte) (*big.Int, error) {
	return _BTCBridge.Contract.GetBridgeFee(&_BTCBridge.CallOpts, payload)
}

// GetBridgeFee is a free data retrieval call binding the contract method 0x5c12ba59.
//
// Solidity: function getBridgeFee(bytes payload) view returns(uint256 bridgeFee)
func (_BTCBridge *BTCBridgeCallerSession) GetBridgeFee(payload []byte) (*big.Int, error) {
	return _BTCBridge.Contract.GetBridgeFee(&_BTCBridge.CallOpts, payload)
}

// GetTotalWithdrawFee is a free data retrieval call binding the contract method 0xe33f61f0.
//
// Solidity: function getTotalWithdrawFee(uint256 amount, bytes recipient) view returns(uint256 fee)
func (_BTCBridge *BTCBridgeCaller) GetTotalWithdrawFee(opts *bind.CallOpts, amount *big.Int, recipient []byte) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "getTotalWithdrawFee", amount, recipient)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalWithdrawFee is a free data retrieval call binding the contract method 0xe33f61f0.
//
// Solidity: function getTotalWithdrawFee(uint256 amount, bytes recipient) view returns(uint256 fee)
func (_BTCBridge *BTCBridgeSession) GetTotalWithdrawFee(amount *big.Int, recipient []byte) (*big.Int, error) {
	return _BTCBridge.Contract.GetTotalWithdrawFee(&_BTCBridge.CallOpts, amount, recipient)
}

// GetTotalWithdrawFee is a free data retrieval call binding the contract method 0xe33f61f0.
//
// Solidity: function getTotalWithdrawFee(uint256 amount, bytes recipient) view returns(uint256 fee)
func (_BTCBridge *BTCBridgeCallerSession) GetTotalWithdrawFee(amount *big.Int, recipient []byte) (*big.Int, error) {
	return _BTCBridge.Contract.GetTotalWithdrawFee(&_BTCBridge.CallOpts, amount, recipient)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_BTCBridge *BTCBridgeCaller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_BTCBridge *BTCBridgeSession) MinWithdrawAmount() (*big.Int, error) {
	return _BTCBridge.Contract.MinWithdrawAmount(&_BTCBridge.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_BTCBridge *BTCBridgeCallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _BTCBridge.Contract.MinWithdrawAmount(&_BTCBridge.CallOpts)
}

// SubSupply is a free data retrieval call binding the contract method 0x5f8ead51.
//
// Solidity: function subSupply() view returns(uint256)
func (_BTCBridge *BTCBridgeCaller) SubSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "subSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubSupply is a free data retrieval call binding the contract method 0x5f8ead51.
//
// Solidity: function subSupply() view returns(uint256)
func (_BTCBridge *BTCBridgeSession) SubSupply() (*big.Int, error) {
	return _BTCBridge.Contract.SubSupply(&_BTCBridge.CallOpts)
}

// SubSupply is a free data retrieval call binding the contract method 0x5f8ead51.
//
// Solidity: function subSupply() view returns(uint256)
func (_BTCBridge *BTCBridgeCallerSession) SubSupply() (*big.Int, error) {
	return _BTCBridge.Contract.SubSupply(&_BTCBridge.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_BTCBridge *BTCBridgeCaller) WithdrawFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BTCBridge.contract.Call(opts, &out, "withdrawFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_BTCBridge *BTCBridgeSession) WithdrawFee() (*big.Int, error) {
	return _BTCBridge.Contract.WithdrawFee(&_BTCBridge.CallOpts)
}

// WithdrawFee is a free data retrieval call binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() view returns(uint256)
func (_BTCBridge *BTCBridgeCallerSession) WithdrawFee() (*big.Int, error) {
	return _BTCBridge.Contract.WithdrawFee(&_BTCBridge.CallOpts)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x1f945d9a.
//
// Solidity: function setMinWithdrawAmount(uint128 newMinWithdrawAmount) returns()
func (_BTCBridge *BTCBridgeTransactor) SetMinWithdrawAmount(opts *bind.TransactOpts, newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _BTCBridge.contract.Transact(opts, "setMinWithdrawAmount", newMinWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x1f945d9a.
//
// Solidity: function setMinWithdrawAmount(uint128 newMinWithdrawAmount) returns()
func (_BTCBridge *BTCBridgeSession) SetMinWithdrawAmount(newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _BTCBridge.Contract.SetMinWithdrawAmount(&_BTCBridge.TransactOpts, newMinWithdrawAmount)
}

// SetMinWithdrawAmount is a paid mutator transaction binding the contract method 0x1f945d9a.
//
// Solidity: function setMinWithdrawAmount(uint128 newMinWithdrawAmount) returns()
func (_BTCBridge *BTCBridgeTransactorSession) SetMinWithdrawAmount(newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _BTCBridge.Contract.SetMinWithdrawAmount(&_BTCBridge.TransactOpts, newMinWithdrawAmount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes recipient) payable returns()
func (_BTCBridge *BTCBridgeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, recipient []byte) (*types.Transaction, error) {
	return _BTCBridge.contract.Transact(opts, "withdraw", amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes recipient) payable returns()
func (_BTCBridge *BTCBridgeSession) Withdraw(amount *big.Int, recipient []byte) (*types.Transaction, error) {
	return _BTCBridge.Contract.Withdraw(&_BTCBridge.TransactOpts, amount, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x030ba25d.
//
// Solidity: function withdraw(uint256 amount, bytes recipient) payable returns()
func (_BTCBridge *BTCBridgeTransactorSession) Withdraw(amount *big.Int, recipient []byte) (*types.Transaction, error) {
	return _BTCBridge.Contract.Withdraw(&_BTCBridge.TransactOpts, amount, recipient)
}

// BTCBridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BTCBridge contract.
type BTCBridgeDepositIterator struct {
	Event *BTCBridgeDeposit // Event containing the contract specifics and raw log

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
func (it *BTCBridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeDeposit)
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
		it.Event = new(BTCBridgeDeposit)
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
func (it *BTCBridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeDeposit represents a Deposit event raised by the BTCBridge contract.
type BTCBridgeDeposit struct {
	Id               [32]byte
	NormalizedAmount *big.Int
	Recipient        common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdd6949a56c21c08b78f6a73eb8bebc859eff308638c0bc5eb78c017148dd0ee0.
//
// Solidity: event Deposit(bytes32 id, uint256 normalizedAmount, address recipient)
func (_BTCBridge *BTCBridgeFilterer) FilterDeposit(opts *bind.FilterOpts) (*BTCBridgeDepositIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeDepositIterator{contract: _BTCBridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdd6949a56c21c08b78f6a73eb8bebc859eff308638c0bc5eb78c017148dd0ee0.
//
// Solidity: event Deposit(bytes32 id, uint256 normalizedAmount, address recipient)
func (_BTCBridge *BTCBridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BTCBridgeDeposit) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeDeposit)
				if err := _BTCBridge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdd6949a56c21c08b78f6a73eb8bebc859eff308638c0bc5eb78c017148dd0ee0.
//
// Solidity: event Deposit(bytes32 id, uint256 normalizedAmount, address recipient)
func (_BTCBridge *BTCBridgeFilterer) ParseDeposit(log types.Log) (*BTCBridgeDeposit, error) {
	event := new(BTCBridgeDeposit)
	if err := _BTCBridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCBridgeSetDepositFeeIterator is returned from FilterSetDepositFee and is used to iterate over the raw logs and unpacked data for SetDepositFee events raised by the BTCBridge contract.
type BTCBridgeSetDepositFeeIterator struct {
	Event *BTCBridgeSetDepositFee // Event containing the contract specifics and raw log

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
func (it *BTCBridgeSetDepositFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeSetDepositFee)
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
		it.Event = new(BTCBridgeSetDepositFee)
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
func (it *BTCBridgeSetDepositFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeSetDepositFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeSetDepositFee represents a SetDepositFee event raised by the BTCBridge contract.
type BTCBridgeSetDepositFee struct {
	NewDepositFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDepositFee is a free log retrieval operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 newDepositFee)
func (_BTCBridge *BTCBridgeFilterer) FilterSetDepositFee(opts *bind.FilterOpts) (*BTCBridgeSetDepositFeeIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeSetDepositFeeIterator{contract: _BTCBridge.contract, event: "SetDepositFee", logs: logs, sub: sub}, nil
}

// WatchSetDepositFee is a free log subscription operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 newDepositFee)
func (_BTCBridge *BTCBridgeFilterer) WatchSetDepositFee(opts *bind.WatchOpts, sink chan<- *BTCBridgeSetDepositFee) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeSetDepositFee)
				if err := _BTCBridge.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
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

// ParseSetDepositFee is a log parse operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 newDepositFee)
func (_BTCBridge *BTCBridgeFilterer) ParseSetDepositFee(log types.Log) (*BTCBridgeSetDepositFee, error) {
	event := new(BTCBridgeSetDepositFee)
	if err := _BTCBridge.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCBridgeSetFeeRecipientIterator is returned from FilterSetFeeRecipient and is used to iterate over the raw logs and unpacked data for SetFeeRecipient events raised by the BTCBridge contract.
type BTCBridgeSetFeeRecipientIterator struct {
	Event *BTCBridgeSetFeeRecipient // Event containing the contract specifics and raw log

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
func (it *BTCBridgeSetFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeSetFeeRecipient)
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
		it.Event = new(BTCBridgeSetFeeRecipient)
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
func (it *BTCBridgeSetFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeSetFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeSetFeeRecipient represents a SetFeeRecipient event raised by the BTCBridge contract.
type BTCBridgeSetFeeRecipient struct {
	NewFeeRecipient common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetFeeRecipient is a free log retrieval operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address newFeeRecipient)
func (_BTCBridge *BTCBridgeFilterer) FilterSetFeeRecipient(opts *bind.FilterOpts) (*BTCBridgeSetFeeRecipientIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "SetFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeSetFeeRecipientIterator{contract: _BTCBridge.contract, event: "SetFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchSetFeeRecipient is a free log subscription operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address newFeeRecipient)
func (_BTCBridge *BTCBridgeFilterer) WatchSetFeeRecipient(opts *bind.WatchOpts, sink chan<- *BTCBridgeSetFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "SetFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeSetFeeRecipient)
				if err := _BTCBridge.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
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

// ParseSetFeeRecipient is a log parse operation binding the contract event 0x2e979f80fe4d43055c584cf4a8467c55875ea36728fc37176c05acd784eb7a73.
//
// Solidity: event SetFeeRecipient(address newFeeRecipient)
func (_BTCBridge *BTCBridgeFilterer) ParseSetFeeRecipient(log types.Log) (*BTCBridgeSetFeeRecipient, error) {
	event := new(BTCBridgeSetFeeRecipient)
	if err := _BTCBridge.contract.UnpackLog(event, "SetFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCBridgeSetMinWithdrawAmountIterator is returned from FilterSetMinWithdrawAmount and is used to iterate over the raw logs and unpacked data for SetMinWithdrawAmount events raised by the BTCBridge contract.
type BTCBridgeSetMinWithdrawAmountIterator struct {
	Event *BTCBridgeSetMinWithdrawAmount // Event containing the contract specifics and raw log

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
func (it *BTCBridgeSetMinWithdrawAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeSetMinWithdrawAmount)
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
		it.Event = new(BTCBridgeSetMinWithdrawAmount)
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
func (it *BTCBridgeSetMinWithdrawAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeSetMinWithdrawAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeSetMinWithdrawAmount represents a SetMinWithdrawAmount event raised by the BTCBridge contract.
type BTCBridgeSetMinWithdrawAmount struct {
	NewMinWithdrawAmount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetMinWithdrawAmount is a free log retrieval operation binding the contract event 0x562d97cf82396452eea1987eacabb4af63827265ee67447bb8fe9d292f7a4508.
//
// Solidity: event SetMinWithdrawAmount(uint256 newMinWithdrawAmount)
func (_BTCBridge *BTCBridgeFilterer) FilterSetMinWithdrawAmount(opts *bind.FilterOpts) (*BTCBridgeSetMinWithdrawAmountIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "SetMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeSetMinWithdrawAmountIterator{contract: _BTCBridge.contract, event: "SetMinWithdrawAmount", logs: logs, sub: sub}, nil
}

// WatchSetMinWithdrawAmount is a free log subscription operation binding the contract event 0x562d97cf82396452eea1987eacabb4af63827265ee67447bb8fe9d292f7a4508.
//
// Solidity: event SetMinWithdrawAmount(uint256 newMinWithdrawAmount)
func (_BTCBridge *BTCBridgeFilterer) WatchSetMinWithdrawAmount(opts *bind.WatchOpts, sink chan<- *BTCBridgeSetMinWithdrawAmount) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "SetMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeSetMinWithdrawAmount)
				if err := _BTCBridge.contract.UnpackLog(event, "SetMinWithdrawAmount", log); err != nil {
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

// ParseSetMinWithdrawAmount is a log parse operation binding the contract event 0x562d97cf82396452eea1987eacabb4af63827265ee67447bb8fe9d292f7a4508.
//
// Solidity: event SetMinWithdrawAmount(uint256 newMinWithdrawAmount)
func (_BTCBridge *BTCBridgeFilterer) ParseSetMinWithdrawAmount(log types.Log) (*BTCBridgeSetMinWithdrawAmount, error) {
	event := new(BTCBridgeSetMinWithdrawAmount)
	if err := _BTCBridge.contract.UnpackLog(event, "SetMinWithdrawAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCBridgeSetWithdrawFeeIterator is returned from FilterSetWithdrawFee and is used to iterate over the raw logs and unpacked data for SetWithdrawFee events raised by the BTCBridge contract.
type BTCBridgeSetWithdrawFeeIterator struct {
	Event *BTCBridgeSetWithdrawFee // Event containing the contract specifics and raw log

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
func (it *BTCBridgeSetWithdrawFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeSetWithdrawFee)
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
		it.Event = new(BTCBridgeSetWithdrawFee)
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
func (it *BTCBridgeSetWithdrawFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeSetWithdrawFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeSetWithdrawFee represents a SetWithdrawFee event raised by the BTCBridge contract.
type BTCBridgeSetWithdrawFee struct {
	NewWithdrawFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetWithdrawFee is a free log retrieval operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 newWithdrawFee)
func (_BTCBridge *BTCBridgeFilterer) FilterSetWithdrawFee(opts *bind.FilterOpts) (*BTCBridgeSetWithdrawFeeIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "SetWithdrawFee")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeSetWithdrawFeeIterator{contract: _BTCBridge.contract, event: "SetWithdrawFee", logs: logs, sub: sub}, nil
}

// WatchSetWithdrawFee is a free log subscription operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 newWithdrawFee)
func (_BTCBridge *BTCBridgeFilterer) WatchSetWithdrawFee(opts *bind.WatchOpts, sink chan<- *BTCBridgeSetWithdrawFee) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "SetWithdrawFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeSetWithdrawFee)
				if err := _BTCBridge.contract.UnpackLog(event, "SetWithdrawFee", log); err != nil {
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

// ParseSetWithdrawFee is a log parse operation binding the contract event 0x7be0a744e4d6f887e4fd578978ae62cb2568d860f0f2eb0a54fd0de804b16440.
//
// Solidity: event SetWithdrawFee(uint256 newWithdrawFee)
func (_BTCBridge *BTCBridgeFilterer) ParseSetWithdrawFee(log types.Log) (*BTCBridgeSetWithdrawFee, error) {
	event := new(BTCBridgeSetWithdrawFee)
	if err := _BTCBridge.contract.UnpackLog(event, "SetWithdrawFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BTCBridgeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the BTCBridge contract.
type BTCBridgeWithdrawIterator struct {
	Event *BTCBridgeWithdraw // Event containing the contract specifics and raw log

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
func (it *BTCBridgeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BTCBridgeWithdraw)
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
		it.Event = new(BTCBridgeWithdraw)
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
func (it *BTCBridgeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BTCBridgeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BTCBridgeWithdraw represents a Withdraw event raised by the BTCBridge contract.
type BTCBridgeWithdraw struct {
	Id               [32]byte
	NormalizedAmount *big.Int
	Sender           common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x181177e6e9bbf7d7a45604cf4fdb0444641e20905d8d4daf2a7863036e56dc15.
//
// Solidity: event Withdraw(bytes32 id, uint256 normalizedAmount, address sender)
func (_BTCBridge *BTCBridgeFilterer) FilterWithdraw(opts *bind.FilterOpts) (*BTCBridgeWithdrawIterator, error) {

	logs, sub, err := _BTCBridge.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &BTCBridgeWithdrawIterator{contract: _BTCBridge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x181177e6e9bbf7d7a45604cf4fdb0444641e20905d8d4daf2a7863036e56dc15.
//
// Solidity: event Withdraw(bytes32 id, uint256 normalizedAmount, address sender)
func (_BTCBridge *BTCBridgeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *BTCBridgeWithdraw) (event.Subscription, error) {

	logs, sub, err := _BTCBridge.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BTCBridgeWithdraw)
				if err := _BTCBridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x181177e6e9bbf7d7a45604cf4fdb0444641e20905d8d4daf2a7863036e56dc15.
//
// Solidity: event Withdraw(bytes32 id, uint256 normalizedAmount, address sender)
func (_BTCBridge *BTCBridgeFilterer) ParseWithdraw(log types.Log) (*BTCBridgeWithdraw, error) {
	event := new(BTCBridgeWithdraw)
	if err := _BTCBridge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
