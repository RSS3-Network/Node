// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tips

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

// TipsMetaData contains all meta data concerning the Tips contract.
var TipsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"error\",\"name\":\"ErrCallerNotCharacterOwner\",\"inputs\":[]},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"internalType\":\"uint8\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TipCharacter\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"toCharacterId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TipCharacterForNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"toCharacterId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"toNoteId\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIERC1820Registry\"}],\"name\":\"ERC1820_REGISTRY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"TOKENS_RECIPIENT_INTERFACE_HASH\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getToken\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getWeb3Entry\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"web3Entry_\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"token_\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"tokensReceived\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"userData\",\"internalType\":\"bytes\"},{\"type\":\"bytes\",\"name\":\"operatorData\",\"internalType\":\"bytes\"}]}]",
}

// TipsABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsMetaData.ABI instead.
var TipsABI = TipsMetaData.ABI

// Tips is an auto generated Go binding around an Ethereum contract.
type Tips struct {
	TipsCaller     // Read-only binding to the contract
	TipsTransactor // Write-only binding to the contract
	TipsFilterer   // Log filterer for contract events
}

// TipsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsSession struct {
	Contract     *Tips             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsCallerSession struct {
	Contract *TipsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TipsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsTransactorSession struct {
	Contract     *TipsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsRaw struct {
	Contract *Tips // Generic contract binding to access the raw methods on
}

// TipsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsCallerRaw struct {
	Contract *TipsCaller // Generic read-only contract binding to access the raw methods on
}

// TipsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsTransactorRaw struct {
	Contract *TipsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTips creates a new instance of Tips, bound to a specific deployed contract.
func NewTips(address common.Address, backend bind.ContractBackend) (*Tips, error) {
	contract, err := bindTips(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tips{TipsCaller: TipsCaller{contract: contract}, TipsTransactor: TipsTransactor{contract: contract}, TipsFilterer: TipsFilterer{contract: contract}}, nil
}

// NewTipsCaller creates a new read-only instance of Tips, bound to a specific deployed contract.
func NewTipsCaller(address common.Address, caller bind.ContractCaller) (*TipsCaller, error) {
	contract, err := bindTips(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsCaller{contract: contract}, nil
}

// NewTipsTransactor creates a new write-only instance of Tips, bound to a specific deployed contract.
func NewTipsTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsTransactor, error) {
	contract, err := bindTips(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsTransactor{contract: contract}, nil
}

// NewTipsFilterer creates a new log filterer instance of Tips, bound to a specific deployed contract.
func NewTipsFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsFilterer, error) {
	contract, err := bindTips(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsFilterer{contract: contract}, nil
}

// bindTips binds a generic wrapper to an already deployed contract.
func bindTips(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TipsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.TipsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transact(opts, method, params...)
}

// ERC1820REGISTRY is a free data retrieval call binding the contract method 0x013eb177.
//
// Solidity: function ERC1820_REGISTRY() view returns(address)
func (_Tips *TipsCaller) ERC1820REGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tips.contract.Call(opts, &out, "ERC1820_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC1820REGISTRY is a free data retrieval call binding the contract method 0x013eb177.
//
// Solidity: function ERC1820_REGISTRY() view returns(address)
func (_Tips *TipsSession) ERC1820REGISTRY() (common.Address, error) {
	return _Tips.Contract.ERC1820REGISTRY(&_Tips.CallOpts)
}

// ERC1820REGISTRY is a free data retrieval call binding the contract method 0x013eb177.
//
// Solidity: function ERC1820_REGISTRY() view returns(address)
func (_Tips *TipsCallerSession) ERC1820REGISTRY() (common.Address, error) {
	return _Tips.Contract.ERC1820REGISTRY(&_Tips.CallOpts)
}

// TOKENSRECIPIENTINTERFACEHASH is a free data retrieval call binding the contract method 0x72581cc0.
//
// Solidity: function TOKENS_RECIPIENT_INTERFACE_HASH() view returns(bytes32)
func (_Tips *TipsCaller) TOKENSRECIPIENTINTERFACEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Tips.contract.Call(opts, &out, "TOKENS_RECIPIENT_INTERFACE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TOKENSRECIPIENTINTERFACEHASH is a free data retrieval call binding the contract method 0x72581cc0.
//
// Solidity: function TOKENS_RECIPIENT_INTERFACE_HASH() view returns(bytes32)
func (_Tips *TipsSession) TOKENSRECIPIENTINTERFACEHASH() ([32]byte, error) {
	return _Tips.Contract.TOKENSRECIPIENTINTERFACEHASH(&_Tips.CallOpts)
}

// TOKENSRECIPIENTINTERFACEHASH is a free data retrieval call binding the contract method 0x72581cc0.
//
// Solidity: function TOKENS_RECIPIENT_INTERFACE_HASH() view returns(bytes32)
func (_Tips *TipsCallerSession) TOKENSRECIPIENTINTERFACEHASH() ([32]byte, error) {
	return _Tips.Contract.TOKENSRECIPIENTINTERFACEHASH(&_Tips.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Tips *TipsCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tips.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Tips *TipsSession) GetToken() (common.Address, error) {
	return _Tips.Contract.GetToken(&_Tips.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Tips *TipsCallerSession) GetToken() (common.Address, error) {
	return _Tips.Contract.GetToken(&_Tips.CallOpts)
}

// GetWeb3Entry is a free data retrieval call binding the contract method 0xd1fc884c.
//
// Solidity: function getWeb3Entry() view returns(address)
func (_Tips *TipsCaller) GetWeb3Entry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tips.contract.Call(opts, &out, "getWeb3Entry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWeb3Entry is a free data retrieval call binding the contract method 0xd1fc884c.
//
// Solidity: function getWeb3Entry() view returns(address)
func (_Tips *TipsSession) GetWeb3Entry() (common.Address, error) {
	return _Tips.Contract.GetWeb3Entry(&_Tips.CallOpts)
}

// GetWeb3Entry is a free data retrieval call binding the contract method 0xd1fc884c.
//
// Solidity: function getWeb3Entry() view returns(address)
func (_Tips *TipsCallerSession) GetWeb3Entry() (common.Address, error) {
	return _Tips.Contract.GetWeb3Entry(&_Tips.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address web3Entry_, address token_) returns()
func (_Tips *TipsTransactor) Initialize(opts *bind.TransactOpts, web3Entry_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Tips.contract.Transact(opts, "initialize", web3Entry_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address web3Entry_, address token_) returns()
func (_Tips *TipsSession) Initialize(web3Entry_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Tips.Contract.Initialize(&_Tips.TransactOpts, web3Entry_, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address web3Entry_, address token_) returns()
func (_Tips *TipsTransactorSession) Initialize(web3Entry_ common.Address, token_ common.Address) (*types.Transaction, error) {
	return _Tips.Contract.Initialize(&_Tips.TransactOpts, web3Entry_, token_)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Tips *TipsTransactor) TokensReceived(opts *bind.TransactOpts, arg0 common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Tips.contract.Transact(opts, "tokensReceived", arg0, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Tips *TipsSession) TokensReceived(arg0 common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Tips.Contract.TokensReceived(&_Tips.TransactOpts, arg0, from, to, amount, userData, operatorData)
}

// TokensReceived is a paid mutator transaction binding the contract method 0x0023de29.
//
// Solidity: function tokensReceived(address , address from, address to, uint256 amount, bytes userData, bytes operatorData) returns()
func (_Tips *TipsTransactorSession) TokensReceived(arg0 common.Address, from common.Address, to common.Address, amount *big.Int, userData []byte, operatorData []byte) (*types.Transaction, error) {
	return _Tips.Contract.TokensReceived(&_Tips.TransactOpts, arg0, from, to, amount, userData, operatorData)
}

// TipsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tips contract.
type TipsInitializedIterator struct {
	Event *TipsInitialized // Event containing the contract specifics and raw log

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
func (it *TipsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TipsInitialized)
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
		it.Event = new(TipsInitialized)
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
func (it *TipsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TipsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TipsInitialized represents a Initialized event raised by the Tips contract.
type TipsInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tips *TipsFilterer) FilterInitialized(opts *bind.FilterOpts) (*TipsInitializedIterator, error) {

	logs, sub, err := _Tips.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TipsInitializedIterator{contract: _Tips.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tips *TipsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TipsInitialized) (event.Subscription, error) {

	logs, sub, err := _Tips.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TipsInitialized)
				if err := _Tips.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tips *TipsFilterer) ParseInitialized(log types.Log) (*TipsInitialized, error) {
	event := new(TipsInitialized)
	if err := _Tips.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsTipCharacterIterator is returned from FilterTipCharacter and is used to iterate over the raw logs and unpacked data for TipCharacter events raised by the Tips contract.
type TipsTipCharacterIterator struct {
	Event *TipsTipCharacter // Event containing the contract specifics and raw log

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
func (it *TipsTipCharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TipsTipCharacter)
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
		it.Event = new(TipsTipCharacter)
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
func (it *TipsTipCharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TipsTipCharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TipsTipCharacter represents a TipCharacter event raised by the Tips contract.
type TipsTipCharacter struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	Token           common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTipCharacter is a free log retrieval operation binding the contract event 0x9aa19e93d7c44ea227cf13c500c5e9bfcc7df95775b643011bf571ac13fa6038.
//
// Solidity: event TipCharacter(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, address token, uint256 amount)
func (_Tips *TipsFilterer) FilterTipCharacter(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int) (*TipsTipCharacterIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Tips.contract.FilterLogs(opts, "TipCharacter", fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &TipsTipCharacterIterator{contract: _Tips.contract, event: "TipCharacter", logs: logs, sub: sub}, nil
}

// WatchTipCharacter is a free log subscription operation binding the contract event 0x9aa19e93d7c44ea227cf13c500c5e9bfcc7df95775b643011bf571ac13fa6038.
//
// Solidity: event TipCharacter(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, address token, uint256 amount)
func (_Tips *TipsFilterer) WatchTipCharacter(opts *bind.WatchOpts, sink chan<- *TipsTipCharacter, fromCharacterId []*big.Int, toCharacterId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Tips.contract.WatchLogs(opts, "TipCharacter", fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TipsTipCharacter)
				if err := _Tips.contract.UnpackLog(event, "TipCharacter", log); err != nil {
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

// ParseTipCharacter is a log parse operation binding the contract event 0x9aa19e93d7c44ea227cf13c500c5e9bfcc7df95775b643011bf571ac13fa6038.
//
// Solidity: event TipCharacter(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, address token, uint256 amount)
func (_Tips *TipsFilterer) ParseTipCharacter(log types.Log) (*TipsTipCharacter, error) {
	event := new(TipsTipCharacter)
	if err := _Tips.contract.UnpackLog(event, "TipCharacter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TipsTipCharacterForNoteIterator is returned from FilterTipCharacterForNote and is used to iterate over the raw logs and unpacked data for TipCharacterForNote events raised by the Tips contract.
type TipsTipCharacterForNoteIterator struct {
	Event *TipsTipCharacterForNote // Event containing the contract specifics and raw log

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
func (it *TipsTipCharacterForNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TipsTipCharacterForNote)
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
		it.Event = new(TipsTipCharacterForNote)
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
func (it *TipsTipCharacterForNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TipsTipCharacterForNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TipsTipCharacterForNote represents a TipCharacterForNote event raised by the Tips contract.
type TipsTipCharacterForNote struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	Token           common.Address
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTipCharacterForNote is a free log retrieval operation binding the contract event 0x3cfc312a05bba2ef0c4efdf6a7a4ef883e9e35f403da1b4cbc810b3ba738dab5.
//
// Solidity: event TipCharacterForNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, address token, uint256 amount)
func (_Tips *TipsFilterer) FilterTipCharacterForNote(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (*TipsTipCharacterForNoteIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Tips.contract.FilterLogs(opts, "TipCharacterForNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &TipsTipCharacterForNoteIterator{contract: _Tips.contract, event: "TipCharacterForNote", logs: logs, sub: sub}, nil
}

// WatchTipCharacterForNote is a free log subscription operation binding the contract event 0x3cfc312a05bba2ef0c4efdf6a7a4ef883e9e35f403da1b4cbc810b3ba738dab5.
//
// Solidity: event TipCharacterForNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, address token, uint256 amount)
func (_Tips *TipsFilterer) WatchTipCharacterForNote(opts *bind.WatchOpts, sink chan<- *TipsTipCharacterForNote, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Tips.contract.WatchLogs(opts, "TipCharacterForNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TipsTipCharacterForNote)
				if err := _Tips.contract.UnpackLog(event, "TipCharacterForNote", log); err != nil {
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

// ParseTipCharacterForNote is a log parse operation binding the contract event 0x3cfc312a05bba2ef0c4efdf6a7a4ef883e9e35f403da1b4cbc810b3ba738dab5.
//
// Solidity: event TipCharacterForNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, address token, uint256 amount)
func (_Tips *TipsFilterer) ParseTipCharacterForNote(log types.Log) (*TipsTipCharacterForNote, error) {
	event := new(TipsTipCharacterForNote)
	if err := _Tips.contract.UnpackLog(event, "TipCharacterForNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
