// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// StableSwapNCoinsMetaData contains all meta data concerning the StableSwapNCoins contract.
var StableSwapNCoinsMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[2]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[2]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[3]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[3]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidityImbalance\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256[4]\",\"name\":\"token_amounts\",\"indexed\":false},{\"type\":\"uint256[4]\",\"name\":\"fees\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"invariant\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"token_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"}]",
}

// StableSwapNCoinsABI is the input ABI used to generate the binding from.
// Deprecated: Use StableSwapNCoinsMetaData.ABI instead.
var StableSwapNCoinsABI = StableSwapNCoinsMetaData.ABI

// StableSwapNCoins is an auto generated Go binding around an Ethereum contract.
type StableSwapNCoins struct {
	StableSwapNCoinsCaller     // Read-only binding to the contract
	StableSwapNCoinsTransactor // Write-only binding to the contract
	StableSwapNCoinsFilterer   // Log filterer for contract events
}

// StableSwapNCoinsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableSwapNCoinsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNCoinsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableSwapNCoinsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNCoinsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableSwapNCoinsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableSwapNCoinsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableSwapNCoinsSession struct {
	Contract     *StableSwapNCoins // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StableSwapNCoinsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableSwapNCoinsCallerSession struct {
	Contract *StableSwapNCoinsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StableSwapNCoinsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableSwapNCoinsTransactorSession struct {
	Contract     *StableSwapNCoinsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StableSwapNCoinsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableSwapNCoinsRaw struct {
	Contract *StableSwapNCoins // Generic contract binding to access the raw methods on
}

// StableSwapNCoinsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableSwapNCoinsCallerRaw struct {
	Contract *StableSwapNCoinsCaller // Generic read-only contract binding to access the raw methods on
}

// StableSwapNCoinsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableSwapNCoinsTransactorRaw struct {
	Contract *StableSwapNCoinsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableSwapNCoins creates a new instance of StableSwapNCoins, bound to a specific deployed contract.
func NewStableSwapNCoins(address common.Address, backend bind.ContractBackend) (*StableSwapNCoins, error) {
	contract, err := bindStableSwapNCoins(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoins{StableSwapNCoinsCaller: StableSwapNCoinsCaller{contract: contract}, StableSwapNCoinsTransactor: StableSwapNCoinsTransactor{contract: contract}, StableSwapNCoinsFilterer: StableSwapNCoinsFilterer{contract: contract}}, nil
}

// NewStableSwapNCoinsCaller creates a new read-only instance of StableSwapNCoins, bound to a specific deployed contract.
func NewStableSwapNCoinsCaller(address common.Address, caller bind.ContractCaller) (*StableSwapNCoinsCaller, error) {
	contract, err := bindStableSwapNCoins(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsCaller{contract: contract}, nil
}

// NewStableSwapNCoinsTransactor creates a new write-only instance of StableSwapNCoins, bound to a specific deployed contract.
func NewStableSwapNCoinsTransactor(address common.Address, transactor bind.ContractTransactor) (*StableSwapNCoinsTransactor, error) {
	contract, err := bindStableSwapNCoins(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsTransactor{contract: contract}, nil
}

// NewStableSwapNCoinsFilterer creates a new log filterer instance of StableSwapNCoins, bound to a specific deployed contract.
func NewStableSwapNCoinsFilterer(address common.Address, filterer bind.ContractFilterer) (*StableSwapNCoinsFilterer, error) {
	contract, err := bindStableSwapNCoins(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsFilterer{contract: contract}, nil
}

// bindStableSwapNCoins binds a generic wrapper to an already deployed contract.
func bindStableSwapNCoins(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StableSwapNCoinsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableSwapNCoins *StableSwapNCoinsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableSwapNCoins.Contract.StableSwapNCoinsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableSwapNCoins *StableSwapNCoinsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNCoins.Contract.StableSwapNCoinsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableSwapNCoins *StableSwapNCoinsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableSwapNCoins.Contract.StableSwapNCoinsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableSwapNCoins *StableSwapNCoinsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StableSwapNCoins.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableSwapNCoins *StableSwapNCoinsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableSwapNCoins.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableSwapNCoins *StableSwapNCoinsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableSwapNCoins.Contract.contract.Transact(opts, method, params...)
}

// StableSwapNCoinsAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidityIterator struct {
	Event *StableSwapNCoinsAddLiquidity // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsAddLiquidity)
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
		it.Event = new(StableSwapNCoinsAddLiquidity)
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
func (it *StableSwapNCoinsAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsAddLiquidity represents a AddLiquidity event raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsAddLiquidityIterator{contract: _StableSwapNCoins.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsAddLiquidity)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x26f55a85081d24974e85c6c00045d0f0453991e95873f52bff0d21af4079a768.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseAddLiquidity(log types.Log) (*StableSwapNCoinsAddLiquidity, error) {
	event := new(StableSwapNCoinsAddLiquidity)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsAddLiquidity0Iterator is returned from FilterAddLiquidity0 and is used to iterate over the raw logs and unpacked data for AddLiquidity0 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidity0Iterator struct {
	Event *StableSwapNCoinsAddLiquidity0 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsAddLiquidity0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsAddLiquidity0)
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
		it.Event = new(StableSwapNCoinsAddLiquidity0)
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
func (it *StableSwapNCoinsAddLiquidity0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsAddLiquidity0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsAddLiquidity0 represents a AddLiquidity0 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidity0 struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity0 is a free log retrieval operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterAddLiquidity0(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsAddLiquidity0Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "AddLiquidity0", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsAddLiquidity0Iterator{contract: _StableSwapNCoins.contract, event: "AddLiquidity0", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity0 is a free log subscription operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchAddLiquidity0(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsAddLiquidity0, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "AddLiquidity0", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsAddLiquidity0)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity0", log); err != nil {
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

// ParseAddLiquidity0 is a log parse operation binding the contract event 0x423f6495a08fc652425cf4ed0d1f9e37e571d9b9529b1c1c23cce780b2e7df0d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseAddLiquidity0(log types.Log) (*StableSwapNCoinsAddLiquidity0, error) {
	event := new(StableSwapNCoinsAddLiquidity0)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsAddLiquidity1Iterator is returned from FilterAddLiquidity1 and is used to iterate over the raw logs and unpacked data for AddLiquidity1 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidity1Iterator struct {
	Event *StableSwapNCoinsAddLiquidity1 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsAddLiquidity1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsAddLiquidity1)
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
		it.Event = new(StableSwapNCoinsAddLiquidity1)
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
func (it *StableSwapNCoinsAddLiquidity1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsAddLiquidity1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsAddLiquidity1 represents a AddLiquidity1 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsAddLiquidity1 struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity1 is a free log retrieval operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterAddLiquidity1(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsAddLiquidity1Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "AddLiquidity1", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsAddLiquidity1Iterator{contract: _StableSwapNCoins.contract, event: "AddLiquidity1", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity1 is a free log subscription operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchAddLiquidity1(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsAddLiquidity1, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "AddLiquidity1", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsAddLiquidity1)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity1", log); err != nil {
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

// ParseAddLiquidity1 is a log parse operation binding the contract event 0x3f1915775e0c9a38a57a7bb7f1f9005f486fb904e1f84aa215364d567319a58d.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseAddLiquidity1(log types.Log) (*StableSwapNCoinsAddLiquidity1, error) {
	event := new(StableSwapNCoinsAddLiquidity1)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "AddLiquidity1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityIterator struct {
	Event *StableSwapNCoinsRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidity)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidity)
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
func (it *StableSwapNCoinsRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidity represents a RemoveLiquidity event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidity struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidityIterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidity)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x7c363854ccf79623411f8995b362bce5eddff18c927edc6f5dbbb5e05819a82c.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidity(log types.Log) (*StableSwapNCoinsRemoveLiquidity, error) {
	event := new(StableSwapNCoinsRemoveLiquidity)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidity0Iterator is returned from FilterRemoveLiquidity0 and is used to iterate over the raw logs and unpacked data for RemoveLiquidity0 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidity0Iterator struct {
	Event *StableSwapNCoinsRemoveLiquidity0 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidity0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidity0)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidity0)
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
func (it *StableSwapNCoinsRemoveLiquidity0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidity0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidity0 represents a RemoveLiquidity0 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidity0 struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity0 is a free log retrieval operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidity0(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidity0Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidity0", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidity0Iterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidity0", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity0 is a free log subscription operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidity0(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidity0, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidity0", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidity0)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity0", log); err != nil {
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

// ParseRemoveLiquidity0 is a log parse operation binding the contract event 0xa49d4cf02656aebf8c771f5a8585638a2a15ee6c97cf7205d4208ed7c1df252d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidity0(log types.Log) (*StableSwapNCoinsRemoveLiquidity0, error) {
	event := new(StableSwapNCoinsRemoveLiquidity0)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidity1Iterator is returned from FilterRemoveLiquidity1 and is used to iterate over the raw logs and unpacked data for RemoveLiquidity1 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidity1Iterator struct {
	Event *StableSwapNCoinsRemoveLiquidity1 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidity1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidity1)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidity1)
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
func (it *StableSwapNCoinsRemoveLiquidity1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidity1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidity1 represents a RemoveLiquidity1 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidity1 struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity1 is a free log retrieval operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidity1(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidity1Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidity1", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidity1Iterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidity1", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity1 is a free log subscription operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidity1(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidity1, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidity1", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidity1)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity1", log); err != nil {
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

// ParseRemoveLiquidity1 is a log parse operation binding the contract event 0x9878ca375e106f2a43c3b599fc624568131c4c9a4ba66a14563715763be9d59d.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidity1(log types.Log) (*StableSwapNCoinsRemoveLiquidity1, error) {
	event := new(StableSwapNCoinsRemoveLiquidity1)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidity1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalanceIterator struct {
	Event *StableSwapNCoinsRemoveLiquidityImbalance // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance)
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
func (it *StableSwapNCoinsRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalance struct {
	Provider     common.Address
	TokenAmounts [2]*big.Int
	Fees         [2]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidityImbalanceIterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidityImbalance)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
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

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x2b5508378d7e19e0d5fa338419034731416c4f5b219a10379956f764317fd47e.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[2] token_amounts, uint256[2] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*StableSwapNCoinsRemoveLiquidityImbalance, error) {
	event := new(StableSwapNCoinsRemoveLiquidityImbalance)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidityImbalance0Iterator is returned from FilterRemoveLiquidityImbalance0 and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance0 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalance0Iterator struct {
	Event *StableSwapNCoinsRemoveLiquidityImbalance0 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidityImbalance0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance0)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance0)
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
func (it *StableSwapNCoinsRemoveLiquidityImbalance0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidityImbalance0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidityImbalance0 represents a RemoveLiquidityImbalance0 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalance0 struct {
	Provider     common.Address
	TokenAmounts [3]*big.Int
	Fees         [3]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance0 is a free log retrieval operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidityImbalance0(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidityImbalance0Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidityImbalance0", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidityImbalance0Iterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidityImbalance0", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance0 is a free log subscription operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidityImbalance0(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidityImbalance0, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidityImbalance0", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidityImbalance0)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance0", log); err != nil {
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

// ParseRemoveLiquidityImbalance0 is a log parse operation binding the contract event 0x173599dbf9c6ca6f7c3b590df07ae98a45d74ff54065505141e7de6c46a624c2.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[3] token_amounts, uint256[3] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidityImbalance0(log types.Log) (*StableSwapNCoinsRemoveLiquidityImbalance0, error) {
	event := new(StableSwapNCoinsRemoveLiquidityImbalance0)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StableSwapNCoinsRemoveLiquidityImbalance1Iterator is returned from FilterRemoveLiquidityImbalance1 and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance1 events raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalance1Iterator struct {
	Event *StableSwapNCoinsRemoveLiquidityImbalance1 // Event containing the contract specifics and raw log

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
func (it *StableSwapNCoinsRemoveLiquidityImbalance1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance1)
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
		it.Event = new(StableSwapNCoinsRemoveLiquidityImbalance1)
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
func (it *StableSwapNCoinsRemoveLiquidityImbalance1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableSwapNCoinsRemoveLiquidityImbalance1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableSwapNCoinsRemoveLiquidityImbalance1 represents a RemoveLiquidityImbalance1 event raised by the StableSwapNCoins contract.
type StableSwapNCoinsRemoveLiquidityImbalance1 struct {
	Provider     common.Address
	TokenAmounts [4]*big.Int
	Fees         [4]*big.Int
	Invariant    *big.Int
	TokenSupply  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance1 is a free log retrieval operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) FilterRemoveLiquidityImbalance1(opts *bind.FilterOpts, provider []common.Address) (*StableSwapNCoinsRemoveLiquidityImbalance1Iterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.FilterLogs(opts, "RemoveLiquidityImbalance1", providerRule)
	if err != nil {
		return nil, err
	}
	return &StableSwapNCoinsRemoveLiquidityImbalance1Iterator{contract: _StableSwapNCoins.contract, event: "RemoveLiquidityImbalance1", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance1 is a free log subscription operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) WatchRemoveLiquidityImbalance1(opts *bind.WatchOpts, sink chan<- *StableSwapNCoinsRemoveLiquidityImbalance1, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _StableSwapNCoins.contract.WatchLogs(opts, "RemoveLiquidityImbalance1", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableSwapNCoinsRemoveLiquidityImbalance1)
				if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance1", log); err != nil {
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

// ParseRemoveLiquidityImbalance1 is a log parse operation binding the contract event 0xb964b72f73f5ef5bf0fdc559b2fab9a7b12a39e47817a547f1f0aee47febd602.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[4] token_amounts, uint256[4] fees, uint256 invariant, uint256 token_supply)
func (_StableSwapNCoins *StableSwapNCoinsFilterer) ParseRemoveLiquidityImbalance1(log types.Log) (*StableSwapNCoinsRemoveLiquidityImbalance1, error) {
	event := new(StableSwapNCoinsRemoveLiquidityImbalance1)
	if err := _StableSwapNCoins.contract.UnpackLog(event, "RemoveLiquidityImbalance1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
