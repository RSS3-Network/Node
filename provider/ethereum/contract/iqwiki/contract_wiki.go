// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iqwiki

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

// IqWikiMetaData contains all meta data concerning the IqWiki contract.
var IqWikiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"DeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WikiNotValid\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"}],\"name\":\"Posted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfs\",\"type\":\"string\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfs\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"postBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidator\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IqWikiABI is the input ABI used to generate the binding from.
// Deprecated: Use IqWikiMetaData.ABI instead.
var IqWikiABI = IqWikiMetaData.ABI

// IqWiki is an auto generated Go binding around an Ethereum contract.
type IqWiki struct {
	IqWikiCaller     // Read-only binding to the contract
	IqWikiTransactor // Write-only binding to the contract
	IqWikiFilterer   // Log filterer for contract events
}

// IqWikiCaller is an auto generated read-only Go binding around an Ethereum contract.
type IqWikiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IqWikiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IqWikiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IqWikiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IqWikiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IqWikiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IqWikiSession struct {
	Contract     *IqWiki           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IqWikiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IqWikiCallerSession struct {
	Contract *IqWikiCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IqWikiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IqWikiTransactorSession struct {
	Contract     *IqWikiTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IqWikiRaw is an auto generated low-level Go binding around an Ethereum contract.
type IqWikiRaw struct {
	Contract *IqWiki // Generic contract binding to access the raw methods on
}

// IqWikiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IqWikiCallerRaw struct {
	Contract *IqWikiCaller // Generic read-only contract binding to access the raw methods on
}

// IqWikiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IqWikiTransactorRaw struct {
	Contract *IqWikiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIqWiki creates a new instance of IqWiki, bound to a specific deployed contract.
func NewIqWiki(address common.Address, backend bind.ContractBackend) (*IqWiki, error) {
	contract, err := bindIqWiki(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IqWiki{IqWikiCaller: IqWikiCaller{contract: contract}, IqWikiTransactor: IqWikiTransactor{contract: contract}, IqWikiFilterer: IqWikiFilterer{contract: contract}}, nil
}

// NewIqWikiCaller creates a new read-only instance of IqWiki, bound to a specific deployed contract.
func NewIqWikiCaller(address common.Address, caller bind.ContractCaller) (*IqWikiCaller, error) {
	contract, err := bindIqWiki(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IqWikiCaller{contract: contract}, nil
}

// NewIqWikiTransactor creates a new write-only instance of IqWiki, bound to a specific deployed contract.
func NewIqWikiTransactor(address common.Address, transactor bind.ContractTransactor) (*IqWikiTransactor, error) {
	contract, err := bindIqWiki(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IqWikiTransactor{contract: contract}, nil
}

// NewIqWikiFilterer creates a new log filterer instance of IqWiki, bound to a specific deployed contract.
func NewIqWikiFilterer(address common.Address, filterer bind.ContractFilterer) (*IqWikiFilterer, error) {
	contract, err := bindIqWiki(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IqWikiFilterer{contract: contract}, nil
}

// bindIqWiki binds a generic wrapper to an already deployed contract.
func bindIqWiki(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IqWikiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IqWiki *IqWikiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IqWiki.Contract.IqWikiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IqWiki *IqWikiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IqWiki.Contract.IqWikiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IqWiki *IqWikiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IqWiki.Contract.IqWikiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IqWiki *IqWikiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IqWiki.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IqWiki *IqWikiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IqWiki.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IqWiki *IqWikiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IqWiki.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IqWiki *IqWikiCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IqWiki.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IqWiki *IqWikiSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IqWiki.Contract.DOMAINSEPARATOR(&_IqWiki.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_IqWiki *IqWikiCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _IqWiki.Contract.DOMAINSEPARATOR(&_IqWiki.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IqWiki *IqWikiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IqWiki.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IqWiki *IqWikiSession) Owner() (common.Address, error) {
	return _IqWiki.Contract.Owner(&_IqWiki.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IqWiki *IqWikiCallerSession) Owner() (common.Address, error) {
	return _IqWiki.Contract.Owner(&_IqWiki.CallOpts)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string ipfs) returns()
func (_IqWiki *IqWikiTransactor) Post(opts *bind.TransactOpts, ipfs string) (*types.Transaction, error) {
	return _IqWiki.contract.Transact(opts, "post", ipfs)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string ipfs) returns()
func (_IqWiki *IqWikiSession) Post(ipfs string) (*types.Transaction, error) {
	return _IqWiki.Contract.Post(&_IqWiki.TransactOpts, ipfs)
}

// Post is a paid mutator transaction binding the contract method 0x8ee93cf3.
//
// Solidity: function post(string ipfs) returns()
func (_IqWiki *IqWikiTransactorSession) Post(ipfs string) (*types.Transaction, error) {
	return _IqWiki.Contract.Post(&_IqWiki.TransactOpts, ipfs)
}

// PostBySig is a paid mutator transaction binding the contract method 0xed53ddb9.
//
// Solidity: function postBySig(string ipfs, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IqWiki *IqWikiTransactor) PostBySig(opts *bind.TransactOpts, ipfs string, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IqWiki.contract.Transact(opts, "postBySig", ipfs, _user, _deadline, _v, _r, _s)
}

// PostBySig is a paid mutator transaction binding the contract method 0xed53ddb9.
//
// Solidity: function postBySig(string ipfs, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IqWiki *IqWikiSession) PostBySig(ipfs string, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IqWiki.Contract.PostBySig(&_IqWiki.TransactOpts, ipfs, _user, _deadline, _v, _r, _s)
}

// PostBySig is a paid mutator transaction binding the contract method 0xed53ddb9.
//
// Solidity: function postBySig(string ipfs, address _user, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_IqWiki *IqWikiTransactorSession) PostBySig(ipfs string, _user common.Address, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _IqWiki.Contract.PostBySig(&_IqWiki.TransactOpts, ipfs, _user, _deadline, _v, _r, _s)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_IqWiki *IqWikiTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IqWiki.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_IqWiki *IqWikiSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _IqWiki.Contract.SetOwner(&_IqWiki.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_IqWiki *IqWikiTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _IqWiki.Contract.SetOwner(&_IqWiki.TransactOpts, newOwner)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_IqWiki *IqWikiTransactor) SetValidator(opts *bind.TransactOpts, _validator common.Address) (*types.Transaction, error) {
	return _IqWiki.contract.Transact(opts, "setValidator", _validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_IqWiki *IqWikiSession) SetValidator(_validator common.Address) (*types.Transaction, error) {
	return _IqWiki.Contract.SetValidator(&_IqWiki.TransactOpts, _validator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _validator) returns()
func (_IqWiki *IqWikiTransactorSession) SetValidator(_validator common.Address) (*types.Transaction, error) {
	return _IqWiki.Contract.SetValidator(&_IqWiki.TransactOpts, _validator)
}

// IqWikiOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the IqWiki contract.
type IqWikiOwnerUpdatedIterator struct {
	Event *IqWikiOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *IqWikiOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IqWikiOwnerUpdated)
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
		it.Event = new(IqWikiOwnerUpdated)
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
func (it *IqWikiOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IqWikiOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IqWikiOwnerUpdated represents a OwnerUpdated event raised by the IqWiki contract.
type IqWikiOwnerUpdated struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed user, address indexed newOwner)
func (_IqWiki *IqWikiFilterer) FilterOwnerUpdated(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*IqWikiOwnerUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IqWiki.contract.FilterLogs(opts, "OwnerUpdated", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IqWikiOwnerUpdatedIterator{contract: _IqWiki.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed user, address indexed newOwner)
func (_IqWiki *IqWikiFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *IqWikiOwnerUpdated, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IqWiki.contract.WatchLogs(opts, "OwnerUpdated", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IqWikiOwnerUpdated)
				if err := _IqWiki.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed user, address indexed newOwner)
func (_IqWiki *IqWikiFilterer) ParseOwnerUpdated(log types.Log) (*IqWikiOwnerUpdated, error) {
	event := new(IqWikiOwnerUpdated)
	if err := _IqWiki.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IqWikiPostedIterator is returned from FilterPosted and is used to iterate over the raw logs and unpacked data for Posted events raised by the IqWiki contract.
type IqWikiPostedIterator struct {
	Event *IqWikiPosted // Event containing the contract specifics and raw log

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
func (it *IqWikiPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IqWikiPosted)
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
		it.Event = new(IqWikiPosted)
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
func (it *IqWikiPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IqWikiPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IqWikiPosted represents a Posted event raised by the IqWiki contract.
type IqWikiPosted struct {
	From common.Address
	Ipfs string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPosted is a free log retrieval operation binding the contract event 0x0fe877471cb763db81561ce83d01be57b6699361a3febbc7a0cdfb18df66246b.
//
// Solidity: event Posted(address indexed _from, string _ipfs)
func (_IqWiki *IqWikiFilterer) FilterPosted(opts *bind.FilterOpts, _from []common.Address) (*IqWikiPostedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IqWiki.contract.FilterLogs(opts, "Posted", _fromRule)
	if err != nil {
		return nil, err
	}
	return &IqWikiPostedIterator{contract: _IqWiki.contract, event: "Posted", logs: logs, sub: sub}, nil
}

// WatchPosted is a free log subscription operation binding the contract event 0x0fe877471cb763db81561ce83d01be57b6699361a3febbc7a0cdfb18df66246b.
//
// Solidity: event Posted(address indexed _from, string _ipfs)
func (_IqWiki *IqWikiFilterer) WatchPosted(opts *bind.WatchOpts, sink chan<- *IqWikiPosted, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _IqWiki.contract.WatchLogs(opts, "Posted", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IqWikiPosted)
				if err := _IqWiki.contract.UnpackLog(event, "Posted", log); err != nil {
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

// ParsePosted is a log parse operation binding the contract event 0x0fe877471cb763db81561ce83d01be57b6699361a3febbc7a0cdfb18df66246b.
//
// Solidity: event Posted(address indexed _from, string _ipfs)
func (_IqWiki *IqWikiFilterer) ParsePosted(log types.Log) (*IqWikiPosted, error) {
	event := new(IqWikiPosted)
	if err := _IqWiki.contract.UnpackLog(event, "Posted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
