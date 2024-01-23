// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package matters

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

// MattersMetaData contains all meta data concerning the Matters contract.
var MattersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidURI\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SelfCuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Curation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Curation\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"curate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri_\",\"type\":\"string\"}],\"name\":\"curate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId_\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MattersABI is the input ABI used to generate the binding from.
// Deprecated: Use MattersMetaData.ABI instead.
var MattersABI = MattersMetaData.ABI

// Matters is an auto generated Go binding around an Ethereum contract.
type Matters struct {
	MattersCaller     // Read-only binding to the contract
	MattersTransactor // Write-only binding to the contract
	MattersFilterer   // Log filterer for contract events
}

// MattersCaller is an auto generated read-only Go binding around an Ethereum contract.
type MattersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MattersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MattersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MattersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MattersSession struct {
	Contract     *Matters          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MattersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MattersCallerSession struct {
	Contract *MattersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MattersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MattersTransactorSession struct {
	Contract     *MattersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MattersRaw is an auto generated low-level Go binding around an Ethereum contract.
type MattersRaw struct {
	Contract *Matters // Generic contract binding to access the raw methods on
}

// MattersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MattersCallerRaw struct {
	Contract *MattersCaller // Generic read-only contract binding to access the raw methods on
}

// MattersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MattersTransactorRaw struct {
	Contract *MattersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMatters creates a new instance of Matters, bound to a specific deployed contract.
func NewMatters(address common.Address, backend bind.ContractBackend) (*Matters, error) {
	contract, err := bindMatters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Matters{MattersCaller: MattersCaller{contract: contract}, MattersTransactor: MattersTransactor{contract: contract}, MattersFilterer: MattersFilterer{contract: contract}}, nil
}

// NewMattersCaller creates a new read-only instance of Matters, bound to a specific deployed contract.
func NewMattersCaller(address common.Address, caller bind.ContractCaller) (*MattersCaller, error) {
	contract, err := bindMatters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MattersCaller{contract: contract}, nil
}

// NewMattersTransactor creates a new write-only instance of Matters, bound to a specific deployed contract.
func NewMattersTransactor(address common.Address, transactor bind.ContractTransactor) (*MattersTransactor, error) {
	contract, err := bindMatters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MattersTransactor{contract: contract}, nil
}

// NewMattersFilterer creates a new log filterer instance of Matters, bound to a specific deployed contract.
func NewMattersFilterer(address common.Address, filterer bind.ContractFilterer) (*MattersFilterer, error) {
	contract, err := bindMatters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MattersFilterer{contract: contract}, nil
}

// bindMatters binds a generic wrapper to an already deployed contract.
func bindMatters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MattersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matters *MattersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matters.Contract.MattersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matters *MattersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matters.Contract.MattersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matters *MattersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matters.Contract.MattersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Matters *MattersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Matters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Matters *MattersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Matters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Matters *MattersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Matters.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Matters *MattersCaller) SupportsInterface(opts *bind.CallOpts, interfaceId_ [4]byte) (bool, error) {
	var out []interface{}
	err := _Matters.contract.Call(opts, &out, "supportsInterface", interfaceId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Matters *MattersSession) SupportsInterface(interfaceId_ [4]byte) (bool, error) {
	return _Matters.Contract.SupportsInterface(&_Matters.CallOpts, interfaceId_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId_) view returns(bool)
func (_Matters *MattersCallerSession) SupportsInterface(interfaceId_ [4]byte) (bool, error) {
	return _Matters.Contract.SupportsInterface(&_Matters.CallOpts, interfaceId_)
}

// Curate is a paid mutator transaction binding the contract method 0xdbcdaf5b.
//
// Solidity: function curate(address to_, address token_, uint256 amount_, string uri_) returns()
func (_Matters *MattersTransactor) Curate(opts *bind.TransactOpts, to_ common.Address, token_ common.Address, amount_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Matters.contract.Transact(opts, "curate", to_, token_, amount_, uri_)
}

// Curate is a paid mutator transaction binding the contract method 0xdbcdaf5b.
//
// Solidity: function curate(address to_, address token_, uint256 amount_, string uri_) returns()
func (_Matters *MattersSession) Curate(to_ common.Address, token_ common.Address, amount_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Matters.Contract.Curate(&_Matters.TransactOpts, to_, token_, amount_, uri_)
}

// Curate is a paid mutator transaction binding the contract method 0xdbcdaf5b.
//
// Solidity: function curate(address to_, address token_, uint256 amount_, string uri_) returns()
func (_Matters *MattersTransactorSession) Curate(to_ common.Address, token_ common.Address, amount_ *big.Int, uri_ string) (*types.Transaction, error) {
	return _Matters.Contract.Curate(&_Matters.TransactOpts, to_, token_, amount_, uri_)
}

// Curate0 is a paid mutator transaction binding the contract method 0xec1cccf8.
//
// Solidity: function curate(address to_, string uri_) payable returns()
func (_Matters *MattersTransactor) Curate0(opts *bind.TransactOpts, to_ common.Address, uri_ string) (*types.Transaction, error) {
	return _Matters.contract.Transact(opts, "curate0", to_, uri_)
}

// Curate0 is a paid mutator transaction binding the contract method 0xec1cccf8.
//
// Solidity: function curate(address to_, string uri_) payable returns()
func (_Matters *MattersSession) Curate0(to_ common.Address, uri_ string) (*types.Transaction, error) {
	return _Matters.Contract.Curate0(&_Matters.TransactOpts, to_, uri_)
}

// Curate0 is a paid mutator transaction binding the contract method 0xec1cccf8.
//
// Solidity: function curate(address to_, string uri_) payable returns()
func (_Matters *MattersTransactorSession) Curate0(to_ common.Address, uri_ string) (*types.Transaction, error) {
	return _Matters.Contract.Curate0(&_Matters.TransactOpts, to_, uri_)
}

// MattersCurationIterator is returned from FilterCuration and is used to iterate over the raw logs and unpacked data for Curation events raised by the Matters contract.
type MattersCurationIterator struct {
	Event *MattersCuration // Event containing the contract specifics and raw log

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
func (it *MattersCurationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattersCuration)
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
		it.Event = new(MattersCuration)
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
func (it *MattersCurationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattersCurationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattersCuration represents a Curation event raised by the Matters contract.
type MattersCuration struct {
	From   common.Address
	To     common.Address
	Token  common.Address
	Uri    string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCuration is a free log retrieval operation binding the contract event 0xc2e41b3d49bbccbac6ceb142bad6119608adf4f1ee1ca5cc6fc332e0ca2fc602.
//
// Solidity: event Curation(address indexed from, address indexed to, address indexed token, string uri, uint256 amount)
func (_Matters *MattersFilterer) FilterCuration(opts *bind.FilterOpts, from []common.Address, to []common.Address, token []common.Address) (*MattersCurationIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Matters.contract.FilterLogs(opts, "Curation", fromRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &MattersCurationIterator{contract: _Matters.contract, event: "Curation", logs: logs, sub: sub}, nil
}

// WatchCuration is a free log subscription operation binding the contract event 0xc2e41b3d49bbccbac6ceb142bad6119608adf4f1ee1ca5cc6fc332e0ca2fc602.
//
// Solidity: event Curation(address indexed from, address indexed to, address indexed token, string uri, uint256 amount)
func (_Matters *MattersFilterer) WatchCuration(opts *bind.WatchOpts, sink chan<- *MattersCuration, from []common.Address, to []common.Address, token []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Matters.contract.WatchLogs(opts, "Curation", fromRule, toRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattersCuration)
				if err := _Matters.contract.UnpackLog(event, "Curation", log); err != nil {
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

// ParseCuration is a log parse operation binding the contract event 0xc2e41b3d49bbccbac6ceb142bad6119608adf4f1ee1ca5cc6fc332e0ca2fc602.
//
// Solidity: event Curation(address indexed from, address indexed to, address indexed token, string uri, uint256 amount)
func (_Matters *MattersFilterer) ParseCuration(log types.Log) (*MattersCuration, error) {
	event := new(MattersCuration)
	if err := _Matters.contract.UnpackLog(event, "Curation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MattersCuration0Iterator is returned from FilterCuration0 and is used to iterate over the raw logs and unpacked data for Curation0 events raised by the Matters contract.
type MattersCuration0Iterator struct {
	Event *MattersCuration0 // Event containing the contract specifics and raw log

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
func (it *MattersCuration0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MattersCuration0)
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
		it.Event = new(MattersCuration0)
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
func (it *MattersCuration0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MattersCuration0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MattersCuration0 represents a Curation0 event raised by the Matters contract.
type MattersCuration0 struct {
	From   common.Address
	To     common.Address
	Uri    string
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCuration0 is a free log retrieval operation binding the contract event 0xd7dc0722e6bc9987e505da6eb6e28fb6cab480d622b478011168976231055694.
//
// Solidity: event Curation(address indexed from, address indexed to, string uri, uint256 amount)
func (_Matters *MattersFilterer) FilterCuration0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MattersCuration0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matters.contract.FilterLogs(opts, "Curation0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MattersCuration0Iterator{contract: _Matters.contract, event: "Curation0", logs: logs, sub: sub}, nil
}

// WatchCuration0 is a free log subscription operation binding the contract event 0xd7dc0722e6bc9987e505da6eb6e28fb6cab480d622b478011168976231055694.
//
// Solidity: event Curation(address indexed from, address indexed to, string uri, uint256 amount)
func (_Matters *MattersFilterer) WatchCuration0(opts *bind.WatchOpts, sink chan<- *MattersCuration0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Matters.contract.WatchLogs(opts, "Curation0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MattersCuration0)
				if err := _Matters.contract.UnpackLog(event, "Curation0", log); err != nil {
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

// ParseCuration0 is a log parse operation binding the contract event 0xd7dc0722e6bc9987e505da6eb6e28fb6cab480d622b478011168976231055694.
//
// Solidity: event Curation(address indexed from, address indexed to, string uri, uint256 amount)
func (_Matters *MattersFilterer) ParseCuration0(log types.Log) (*MattersCuration0, error) {
	event := new(MattersCuration0)
	if err := _Matters.contract.UnpackLog(event, "Curation0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
