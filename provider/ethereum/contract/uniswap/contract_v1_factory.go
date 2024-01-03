// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// V1FactoryMetaData contains all meta data concerning the V1Factory contract.
var V1FactoryMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"NewExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"exchange\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initializeFactory\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"template\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35725},{\"name\":\"createExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":187911},{\"name\":\"getExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"getToken\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"exchange\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":745},{\"name\":\"getTokenWithId\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"token_id\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":736},{\"name\":\"exchangeTemplate\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"tokenCount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]",
}

// V1FactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use V1FactoryMetaData.ABI instead.
var V1FactoryABI = V1FactoryMetaData.ABI

// V1Factory is an auto generated Go binding around an Ethereum contract.
type V1Factory struct {
	V1FactoryCaller     // Read-only binding to the contract
	V1FactoryTransactor // Write-only binding to the contract
	V1FactoryFilterer   // Log filterer for contract events
}

// V1FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1FactorySession struct {
	Contract     *V1Factory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1FactoryCallerSession struct {
	Contract *V1FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// V1FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1FactoryTransactorSession struct {
	Contract     *V1FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// V1FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1FactoryRaw struct {
	Contract *V1Factory // Generic contract binding to access the raw methods on
}

// V1FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1FactoryCallerRaw struct {
	Contract *V1FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// V1FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1FactoryTransactorRaw struct {
	Contract *V1FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1Factory creates a new instance of V1Factory, bound to a specific deployed contract.
func NewV1Factory(address common.Address, backend bind.ContractBackend) (*V1Factory, error) {
	contract, err := bindV1Factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1Factory{V1FactoryCaller: V1FactoryCaller{contract: contract}, V1FactoryTransactor: V1FactoryTransactor{contract: contract}, V1FactoryFilterer: V1FactoryFilterer{contract: contract}}, nil
}

// NewV1FactoryCaller creates a new read-only instance of V1Factory, bound to a specific deployed contract.
func NewV1FactoryCaller(address common.Address, caller bind.ContractCaller) (*V1FactoryCaller, error) {
	contract, err := bindV1Factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1FactoryCaller{contract: contract}, nil
}

// NewV1FactoryTransactor creates a new write-only instance of V1Factory, bound to a specific deployed contract.
func NewV1FactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*V1FactoryTransactor, error) {
	contract, err := bindV1Factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1FactoryTransactor{contract: contract}, nil
}

// NewV1FactoryFilterer creates a new log filterer instance of V1Factory, bound to a specific deployed contract.
func NewV1FactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*V1FactoryFilterer, error) {
	contract, err := bindV1Factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1FactoryFilterer{contract: contract}, nil
}

// bindV1Factory binds a generic wrapper to an already deployed contract.
func bindV1Factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V1FactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Factory *V1FactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Factory.Contract.V1FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Factory *V1FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Factory.Contract.V1FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Factory *V1FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Factory.Contract.V1FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Factory *V1FactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Factory *V1FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Factory *V1FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Factory.Contract.contract.Transact(opts, method, params...)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_V1Factory *V1FactoryCaller) ExchangeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1Factory.contract.Call(opts, &out, "exchangeTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_V1Factory *V1FactorySession) ExchangeTemplate() (common.Address, error) {
	return _V1Factory.Contract.ExchangeTemplate(&_V1Factory.CallOpts)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_V1Factory *V1FactoryCallerSession) ExchangeTemplate() (common.Address, error) {
	return _V1Factory.Contract.ExchangeTemplate(&_V1Factory.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_V1Factory *V1FactoryCaller) GetExchange(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _V1Factory.contract.Call(opts, &out, "getExchange", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_V1Factory *V1FactorySession) GetExchange(token common.Address) (common.Address, error) {
	return _V1Factory.Contract.GetExchange(&_V1Factory.CallOpts, token)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_V1Factory *V1FactoryCallerSession) GetExchange(token common.Address) (common.Address, error) {
	return _V1Factory.Contract.GetExchange(&_V1Factory.CallOpts, token)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_V1Factory *V1FactoryCaller) GetToken(opts *bind.CallOpts, exchange common.Address) (common.Address, error) {
	var out []interface{}
	err := _V1Factory.contract.Call(opts, &out, "getToken", exchange)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_V1Factory *V1FactorySession) GetToken(exchange common.Address) (common.Address, error) {
	return _V1Factory.Contract.GetToken(&_V1Factory.CallOpts, exchange)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_V1Factory *V1FactoryCallerSession) GetToken(exchange common.Address) (common.Address, error) {
	return _V1Factory.Contract.GetToken(&_V1Factory.CallOpts, exchange)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_V1Factory *V1FactoryCaller) GetTokenWithId(opts *bind.CallOpts, token_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1Factory.contract.Call(opts, &out, "getTokenWithId", token_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_V1Factory *V1FactorySession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _V1Factory.Contract.GetTokenWithId(&_V1Factory.CallOpts, token_id)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_V1Factory *V1FactoryCallerSession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _V1Factory.Contract.GetTokenWithId(&_V1Factory.CallOpts, token_id)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_V1Factory *V1FactoryCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1Factory.contract.Call(opts, &out, "tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_V1Factory *V1FactorySession) TokenCount() (*big.Int, error) {
	return _V1Factory.Contract.TokenCount(&_V1Factory.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_V1Factory *V1FactoryCallerSession) TokenCount() (*big.Int, error) {
	return _V1Factory.Contract.TokenCount(&_V1Factory.CallOpts)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_V1Factory *V1FactoryTransactor) CreateExchange(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _V1Factory.contract.Transact(opts, "createExchange", token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_V1Factory *V1FactorySession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _V1Factory.Contract.CreateExchange(&_V1Factory.TransactOpts, token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_V1Factory *V1FactoryTransactorSession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _V1Factory.Contract.CreateExchange(&_V1Factory.TransactOpts, token)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_V1Factory *V1FactoryTransactor) InitializeFactory(opts *bind.TransactOpts, template common.Address) (*types.Transaction, error) {
	return _V1Factory.contract.Transact(opts, "initializeFactory", template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_V1Factory *V1FactorySession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _V1Factory.Contract.InitializeFactory(&_V1Factory.TransactOpts, template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_V1Factory *V1FactoryTransactorSession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _V1Factory.Contract.InitializeFactory(&_V1Factory.TransactOpts, template)
}

// V1FactoryNewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the V1Factory contract.
type V1FactoryNewExchangeIterator struct {
	Event *V1FactoryNewExchange // Event containing the contract specifics and raw log

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
func (it *V1FactoryNewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1FactoryNewExchange)
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
		it.Event = new(V1FactoryNewExchange)
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
func (it *V1FactoryNewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1FactoryNewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1FactoryNewExchange represents a NewExchange event raised by the V1Factory contract.
type V1FactoryNewExchange struct {
	Token    common.Address
	Exchange common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_V1Factory *V1FactoryFilterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, exchange []common.Address) (*V1FactoryNewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _V1Factory.contract.FilterLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return &V1FactoryNewExchangeIterator{contract: _V1Factory.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_V1Factory *V1FactoryFilterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *V1FactoryNewExchange, token []common.Address, exchange []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _V1Factory.contract.WatchLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1FactoryNewExchange)
				if err := _V1Factory.contract.UnpackLog(event, "NewExchange", log); err != nil {
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

// ParseNewExchange is a log parse operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_V1Factory *V1FactoryFilterer) ParseNewExchange(log types.Log) (*V1FactoryNewExchange, error) {
	event := new(V1FactoryNewExchange)
	if err := _V1Factory.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
