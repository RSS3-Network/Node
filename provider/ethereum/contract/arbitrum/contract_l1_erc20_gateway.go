// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum

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

// L1ERC20GatewayMetaData contains all meta data concerning the L1ERC20Gateway contract.
var L1ERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// L1ERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC20GatewayMetaData.ABI instead.
var L1ERC20GatewayABI = L1ERC20GatewayMetaData.ABI

// L1ERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L1ERC20Gateway struct {
	L1ERC20GatewayCaller     // Read-only binding to the contract
	L1ERC20GatewayTransactor // Write-only binding to the contract
	L1ERC20GatewayFilterer   // Log filterer for contract events
}

// L1ERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC20GatewaySession struct {
	Contract     *L1ERC20Gateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC20GatewayCallerSession struct {
	Contract *L1ERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L1ERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC20GatewayTransactorSession struct {
	Contract     *L1ERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L1ERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC20GatewayRaw struct {
	Contract *L1ERC20Gateway // Generic contract binding to access the raw methods on
}

// L1ERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC20GatewayCallerRaw struct {
	Contract *L1ERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC20GatewayTransactorRaw struct {
	Contract *L1ERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC20Gateway creates a new instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20Gateway(address common.Address, backend bind.ContractBackend) (*L1ERC20Gateway, error) {
	contract, err := bindL1ERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Gateway{L1ERC20GatewayCaller: L1ERC20GatewayCaller{contract: contract}, L1ERC20GatewayTransactor: L1ERC20GatewayTransactor{contract: contract}, L1ERC20GatewayFilterer: L1ERC20GatewayFilterer{contract: contract}}, nil
}

// NewL1ERC20GatewayCaller creates a new read-only instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ERC20GatewayCaller, error) {
	contract, err := bindL1ERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayCaller{contract: contract}, nil
}

// NewL1ERC20GatewayTransactor creates a new write-only instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC20GatewayTransactor, error) {
	contract, err := bindL1ERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayTransactor{contract: contract}, nil
}

// NewL1ERC20GatewayFilterer creates a new log filterer instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC20GatewayFilterer, error) {
	contract, err := bindL1ERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayFilterer{contract: contract}, nil
}

// bindL1ERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL1ERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1ERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Gateway *L1ERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Gateway *L1ERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Gateway *L1ERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) Admin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "admin")
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_L1ERC20Gateway *L1ERC20GatewaySession) Admin() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Admin(&_L1ERC20Gateway.TransactOpts)
}

// Admin is a paid mutator transaction binding the contract method 0xf851a440.
//
// Solidity: function admin() returns(address admin_)
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) Admin() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Admin(&_L1ERC20Gateway.TransactOpts)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.ChangeAdmin(&_L1ERC20Gateway.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.ChangeAdmin(&_L1ERC20Gateway.TransactOpts, newAdmin)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) Implementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "implementation")
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_L1ERC20Gateway *L1ERC20GatewaySession) Implementation() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Implementation(&_L1ERC20Gateway.TransactOpts)
}

// Implementation is a paid mutator transaction binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() returns(address implementation_)
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) Implementation() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Implementation(&_L1ERC20Gateway.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.UpgradeTo(&_L1ERC20Gateway.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.UpgradeTo(&_L1ERC20Gateway.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.UpgradeToAndCall(&_L1ERC20Gateway.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.UpgradeToAndCall(&_L1ERC20Gateway.TransactOpts, newImplementation, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Fallback(&_L1ERC20Gateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Fallback(&_L1ERC20Gateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) Receive() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Receive(&_L1ERC20Gateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Receive(&_L1ERC20Gateway.TransactOpts)
}

// L1ERC20GatewayAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayAdminChangedIterator struct {
	Event *L1ERC20GatewayAdminChanged // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayAdminChanged)
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
		it.Event = new(L1ERC20GatewayAdminChanged)
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
func (it *L1ERC20GatewayAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayAdminChanged represents a AdminChanged event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*L1ERC20GatewayAdminChangedIterator, error) {

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayAdminChangedIterator{contract: _L1ERC20Gateway.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayAdminChanged) (event.Subscription, error) {

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayAdminChanged)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseAdminChanged(log types.Log) (*L1ERC20GatewayAdminChanged, error) {
	event := new(L1ERC20GatewayAdminChanged)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20GatewayUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayUpgradedIterator struct {
	Event *L1ERC20GatewayUpgraded // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayUpgraded)
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
		it.Event = new(L1ERC20GatewayUpgraded)
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
func (it *L1ERC20GatewayUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayUpgraded represents a Upgraded event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*L1ERC20GatewayUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayUpgradedIterator{contract: _L1ERC20Gateway.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayUpgraded)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseUpgraded(log types.Log) (*L1ERC20GatewayUpgraded, error) {
	event := new(L1ERC20GatewayUpgraded)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
