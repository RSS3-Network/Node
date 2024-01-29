// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package event

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

// EventMetaData contains all meta data concerning the Event contract.
var EventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"AddOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"AttachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CharacterCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"DeleteNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"DetachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"GrantOperatorPermissions\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"permissionBitMap\",\"type\":\"uint256\"}],\"name\":\"GrantOperatorPermissions4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkCharacter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"LinkCharacterLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LinklistNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"LockNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MintNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MintNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PostNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RemoveOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetCharacterUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"SetHandle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Character\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4ERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Linklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetMintModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetNoteUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldCharacterId\",\"type\":\"uint256\"}],\"name\":\"SetPrimaryCharacterId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"SetSocialToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkCharacter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromCharactereId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkCharacterLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Web3EntryInitialized\",\"type\":\"event\"}]",
}

// EventABI is the input ABI used to generate the binding from.
// Deprecated: Use EventMetaData.ABI instead.
var EventABI = EventMetaData.ABI

// Event is an auto generated Go binding around an Ethereum contract.
type Event struct {
	EventCaller     // Read-only binding to the contract
	EventTransactor // Write-only binding to the contract
	EventFilterer   // Log filterer for contract events
}

// EventCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventSession struct {
	Contract     *Event            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventCallerSession struct {
	Contract *EventCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventTransactorSession struct {
	Contract     *EventTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventRaw struct {
	Contract *Event // Generic contract binding to access the raw methods on
}

// EventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventCallerRaw struct {
	Contract *EventCaller // Generic read-only contract binding to access the raw methods on
}

// EventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventTransactorRaw struct {
	Contract *EventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvent creates a new instance of Event, bound to a specific deployed contract.
func NewEvent(address common.Address, backend bind.ContractBackend) (*Event, error) {
	contract, err := bindEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Event{EventCaller: EventCaller{contract: contract}, EventTransactor: EventTransactor{contract: contract}, EventFilterer: EventFilterer{contract: contract}}, nil
}

// NewEventCaller creates a new read-only instance of Event, bound to a specific deployed contract.
func NewEventCaller(address common.Address, caller bind.ContractCaller) (*EventCaller, error) {
	contract, err := bindEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventCaller{contract: contract}, nil
}

// NewEventTransactor creates a new write-only instance of Event, bound to a specific deployed contract.
func NewEventTransactor(address common.Address, transactor bind.ContractTransactor) (*EventTransactor, error) {
	contract, err := bindEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventTransactor{contract: contract}, nil
}

// NewEventFilterer creates a new log filterer instance of Event, bound to a specific deployed contract.
func NewEventFilterer(address common.Address, filterer bind.ContractFilterer) (*EventFilterer, error) {
	contract, err := bindEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventFilterer{contract: contract}, nil
}

// bindEvent binds a generic wrapper to an already deployed contract.
func bindEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.EventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.EventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Event *EventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Event.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Event *EventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Event.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Event *EventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Event.Contract.contract.Transact(opts, method, params...)
}

// EventAddOperatorIterator is returned from FilterAddOperator and is used to iterate over the raw logs and unpacked data for AddOperator events raised by the Event contract.
type EventAddOperatorIterator struct {
	Event *EventAddOperator // Event containing the contract specifics and raw log

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
func (it *EventAddOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventAddOperator)
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
		it.Event = new(EventAddOperator)
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
func (it *EventAddOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventAddOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventAddOperator represents a AddOperator event raised by the Event contract.
type EventAddOperator struct {
	CharacterId *big.Int
	Operator    common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddOperator is a free log retrieval operation binding the contract event 0x58f51b5bb567864de85c6a422b33491f2418924a44613d7b9341f58657bdd833.
//
// Solidity: event AddOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) FilterAddOperator(opts *bind.FilterOpts, characterId []*big.Int, operator []common.Address) (*EventAddOperatorIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "AddOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EventAddOperatorIterator{contract: _Event.contract, event: "AddOperator", logs: logs, sub: sub}, nil
}

// WatchAddOperator is a free log subscription operation binding the contract event 0x58f51b5bb567864de85c6a422b33491f2418924a44613d7b9341f58657bdd833.
//
// Solidity: event AddOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) WatchAddOperator(opts *bind.WatchOpts, sink chan<- *EventAddOperator, characterId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "AddOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventAddOperator)
				if err := _Event.contract.UnpackLog(event, "AddOperator", log); err != nil {
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

// ParseAddOperator is a log parse operation binding the contract event 0x58f51b5bb567864de85c6a422b33491f2418924a44613d7b9341f58657bdd833.
//
// Solidity: event AddOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) ParseAddOperator(log types.Log) (*EventAddOperator, error) {
	event := new(EventAddOperator)
	if err := _Event.contract.UnpackLog(event, "AddOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventAttachLinklistIterator is returned from FilterAttachLinklist and is used to iterate over the raw logs and unpacked data for AttachLinklist events raised by the Event contract.
type EventAttachLinklistIterator struct {
	Event *EventAttachLinklist // Event containing the contract specifics and raw log

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
func (it *EventAttachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventAttachLinklist)
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
		it.Event = new(EventAttachLinklist)
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
func (it *EventAttachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventAttachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventAttachLinklist represents a AttachLinklist event raised by the Event contract.
type EventAttachLinklist struct {
	LinklistId  *big.Int
	CharacterId *big.Int
	LinkType    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttachLinklist is a free log retrieval operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) FilterAttachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (*EventAttachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "AttachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &EventAttachLinklistIterator{contract: _Event.contract, event: "AttachLinklist", logs: logs, sub: sub}, nil
}

// WatchAttachLinklist is a free log subscription operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) WatchAttachLinklist(opts *bind.WatchOpts, sink chan<- *EventAttachLinklist, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "AttachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventAttachLinklist)
				if err := _Event.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
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

// ParseAttachLinklist is a log parse operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) ParseAttachLinklist(log types.Log) (*EventAttachLinklist, error) {
	event := new(EventAttachLinklist)
	if err := _Event.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the Event contract.
type EventBaseInitializedIterator struct {
	Event *EventBaseInitialized // Event containing the contract specifics and raw log

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
func (it *EventBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventBaseInitialized)
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
		it.Event = new(EventBaseInitialized)
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
func (it *EventBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventBaseInitialized represents a BaseInitialized event raised by the Event contract.
type EventBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Event *EventFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*EventBaseInitializedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &EventBaseInitializedIterator{contract: _Event.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Event *EventFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *EventBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventBaseInitialized)
				if err := _Event.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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

// ParseBaseInitialized is a log parse operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Event *EventFilterer) ParseBaseInitialized(log types.Log) (*EventBaseInitialized, error) {
	event := new(EventBaseInitialized)
	if err := _Event.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventCharacterCreatedIterator is returned from FilterCharacterCreated and is used to iterate over the raw logs and unpacked data for CharacterCreated events raised by the Event contract.
type EventCharacterCreatedIterator struct {
	Event *EventCharacterCreated // Event containing the contract specifics and raw log

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
func (it *EventCharacterCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventCharacterCreated)
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
		it.Event = new(EventCharacterCreated)
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
func (it *EventCharacterCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventCharacterCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventCharacterCreated represents a CharacterCreated event raised by the Event contract.
type EventCharacterCreated struct {
	CharacterId *big.Int
	Creator     common.Address
	To          common.Address
	Handle      string
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCharacterCreated is a free log retrieval operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Event *EventFilterer) FilterCharacterCreated(opts *bind.FilterOpts, characterId []*big.Int, creator []common.Address, to []common.Address) (*EventCharacterCreatedIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "CharacterCreated", characterIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EventCharacterCreatedIterator{contract: _Event.contract, event: "CharacterCreated", logs: logs, sub: sub}, nil
}

// WatchCharacterCreated is a free log subscription operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Event *EventFilterer) WatchCharacterCreated(opts *bind.WatchOpts, sink chan<- *EventCharacterCreated, characterId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "CharacterCreated", characterIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventCharacterCreated)
				if err := _Event.contract.UnpackLog(event, "CharacterCreated", log); err != nil {
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

// ParseCharacterCreated is a log parse operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Event *EventFilterer) ParseCharacterCreated(log types.Log) (*EventCharacterCreated, error) {
	event := new(EventCharacterCreated)
	if err := _Event.contract.UnpackLog(event, "CharacterCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventDeleteNoteIterator is returned from FilterDeleteNote and is used to iterate over the raw logs and unpacked data for DeleteNote events raised by the Event contract.
type EventDeleteNoteIterator struct {
	Event *EventDeleteNote // Event containing the contract specifics and raw log

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
func (it *EventDeleteNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventDeleteNote)
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
		it.Event = new(EventDeleteNote)
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
func (it *EventDeleteNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventDeleteNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventDeleteNote represents a DeleteNote event raised by the Event contract.
type EventDeleteNote struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteNote is a free log retrieval operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) FilterDeleteNote(opts *bind.FilterOpts, characterId []*big.Int) (*EventDeleteNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "DeleteNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventDeleteNoteIterator{contract: _Event.contract, event: "DeleteNote", logs: logs, sub: sub}, nil
}

// WatchDeleteNote is a free log subscription operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) WatchDeleteNote(opts *bind.WatchOpts, sink chan<- *EventDeleteNote, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "DeleteNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventDeleteNote)
				if err := _Event.contract.UnpackLog(event, "DeleteNote", log); err != nil {
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

// ParseDeleteNote is a log parse operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) ParseDeleteNote(log types.Log) (*EventDeleteNote, error) {
	event := new(EventDeleteNote)
	if err := _Event.contract.UnpackLog(event, "DeleteNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventDetachLinklistIterator is returned from FilterDetachLinklist and is used to iterate over the raw logs and unpacked data for DetachLinklist events raised by the Event contract.
type EventDetachLinklistIterator struct {
	Event *EventDetachLinklist // Event containing the contract specifics and raw log

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
func (it *EventDetachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventDetachLinklist)
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
		it.Event = new(EventDetachLinklist)
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
func (it *EventDetachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventDetachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventDetachLinklist represents a DetachLinklist event raised by the Event contract.
type EventDetachLinklist struct {
	LinklistId  *big.Int
	CharacterId *big.Int
	LinkType    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDetachLinklist is a free log retrieval operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) FilterDetachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (*EventDetachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "DetachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &EventDetachLinklistIterator{contract: _Event.contract, event: "DetachLinklist", logs: logs, sub: sub}, nil
}

// WatchDetachLinklist is a free log subscription operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) WatchDetachLinklist(opts *bind.WatchOpts, sink chan<- *EventDetachLinklist, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "DetachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventDetachLinklist)
				if err := _Event.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
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

// ParseDetachLinklist is a log parse operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Event *EventFilterer) ParseDetachLinklist(log types.Log) (*EventDetachLinklist, error) {
	event := new(EventDetachLinklist)
	if err := _Event.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventGrantOperatorPermissionsIterator is returned from FilterGrantOperatorPermissions and is used to iterate over the raw logs and unpacked data for GrantOperatorPermissions events raised by the Event contract.
type EventGrantOperatorPermissionsIterator struct {
	Event *EventGrantOperatorPermissions // Event containing the contract specifics and raw log

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
func (it *EventGrantOperatorPermissionsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventGrantOperatorPermissions)
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
		it.Event = new(EventGrantOperatorPermissions)
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
func (it *EventGrantOperatorPermissionsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventGrantOperatorPermissionsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventGrantOperatorPermissions represents a GrantOperatorPermissions event raised by the Event contract.
type EventGrantOperatorPermissions struct {
	CharacterId      *big.Int
	Operator         common.Address
	PermissionBitMap *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGrantOperatorPermissions is a free log retrieval operation binding the contract event 0x4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd27.
//
// Solidity: event GrantOperatorPermissions(uint256 indexed characterId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) FilterGrantOperatorPermissions(opts *bind.FilterOpts, characterId []*big.Int, operator []common.Address) (*EventGrantOperatorPermissionsIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "GrantOperatorPermissions", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EventGrantOperatorPermissionsIterator{contract: _Event.contract, event: "GrantOperatorPermissions", logs: logs, sub: sub}, nil
}

// WatchGrantOperatorPermissions is a free log subscription operation binding the contract event 0x4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd27.
//
// Solidity: event GrantOperatorPermissions(uint256 indexed characterId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) WatchGrantOperatorPermissions(opts *bind.WatchOpts, sink chan<- *EventGrantOperatorPermissions, characterId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "GrantOperatorPermissions", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventGrantOperatorPermissions)
				if err := _Event.contract.UnpackLog(event, "GrantOperatorPermissions", log); err != nil {
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

// ParseGrantOperatorPermissions is a log parse operation binding the contract event 0x4b0dcce4c30a5691ff14b7d4a8a27e2738b66289dc90120ebbc092812941bd27.
//
// Solidity: event GrantOperatorPermissions(uint256 indexed characterId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) ParseGrantOperatorPermissions(log types.Log) (*EventGrantOperatorPermissions, error) {
	event := new(EventGrantOperatorPermissions)
	if err := _Event.contract.UnpackLog(event, "GrantOperatorPermissions", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventGrantOperatorPermissions4NoteIterator is returned from FilterGrantOperatorPermissions4Note and is used to iterate over the raw logs and unpacked data for GrantOperatorPermissions4Note events raised by the Event contract.
type EventGrantOperatorPermissions4NoteIterator struct {
	Event *EventGrantOperatorPermissions4Note // Event containing the contract specifics and raw log

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
func (it *EventGrantOperatorPermissions4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventGrantOperatorPermissions4Note)
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
		it.Event = new(EventGrantOperatorPermissions4Note)
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
func (it *EventGrantOperatorPermissions4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventGrantOperatorPermissions4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventGrantOperatorPermissions4Note represents a GrantOperatorPermissions4Note event raised by the Event contract.
type EventGrantOperatorPermissions4Note struct {
	CharacterId      *big.Int
	NoteId           *big.Int
	Operator         common.Address
	PermissionBitMap *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterGrantOperatorPermissions4Note is a free log retrieval operation binding the contract event 0x488a41148fb8f04fec9e8e1f936444b9b70c0f084a5242092d8b4b01f6163d80.
//
// Solidity: event GrantOperatorPermissions4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) FilterGrantOperatorPermissions4Note(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, operator []common.Address) (*EventGrantOperatorPermissions4NoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "GrantOperatorPermissions4Note", characterIdRule, noteIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EventGrantOperatorPermissions4NoteIterator{contract: _Event.contract, event: "GrantOperatorPermissions4Note", logs: logs, sub: sub}, nil
}

// WatchGrantOperatorPermissions4Note is a free log subscription operation binding the contract event 0x488a41148fb8f04fec9e8e1f936444b9b70c0f084a5242092d8b4b01f6163d80.
//
// Solidity: event GrantOperatorPermissions4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) WatchGrantOperatorPermissions4Note(opts *bind.WatchOpts, sink chan<- *EventGrantOperatorPermissions4Note, characterId []*big.Int, noteId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "GrantOperatorPermissions4Note", characterIdRule, noteIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventGrantOperatorPermissions4Note)
				if err := _Event.contract.UnpackLog(event, "GrantOperatorPermissions4Note", log); err != nil {
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

// ParseGrantOperatorPermissions4Note is a log parse operation binding the contract event 0x488a41148fb8f04fec9e8e1f936444b9b70c0f084a5242092d8b4b01f6163d80.
//
// Solidity: event GrantOperatorPermissions4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed operator, uint256 permissionBitMap)
func (_Event *EventFilterer) ParseGrantOperatorPermissions4Note(log types.Log) (*EventGrantOperatorPermissions4Note, error) {
	event := new(EventGrantOperatorPermissions4Note)
	if err := _Event.contract.UnpackLog(event, "GrantOperatorPermissions4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkAddressIterator is returned from FilterLinkAddress and is used to iterate over the raw logs and unpacked data for LinkAddress events raised by the Event contract.
type EventLinkAddressIterator struct {
	Event *EventLinkAddress // Event containing the contract specifics and raw log

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
func (it *EventLinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkAddress)
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
		it.Event = new(EventLinkAddress)
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
func (it *EventLinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkAddress represents a LinkAddress event raised by the Event contract.
type EventLinkAddress struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkAddress is a free log retrieval operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterLinkAddress(opts *bind.FilterOpts, fromCharacterId []*big.Int, ethAddress []common.Address) (*EventLinkAddressIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkAddressIterator{contract: _Event.contract, event: "LinkAddress", logs: logs, sub: sub}, nil
}

// WatchLinkAddress is a free log subscription operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchLinkAddress(opts *bind.WatchOpts, sink chan<- *EventLinkAddress, fromCharacterId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkAddress)
				if err := _Event.contract.UnpackLog(event, "LinkAddress", log); err != nil {
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

// ParseLinkAddress is a log parse operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseLinkAddress(log types.Log) (*EventLinkAddress, error) {
	event := new(EventLinkAddress)
	if err := _Event.contract.UnpackLog(event, "LinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkAnyUriIterator is returned from FilterLinkAnyUri and is used to iterate over the raw logs and unpacked data for LinkAnyUri events raised by the Event contract.
type EventLinkAnyUriIterator struct {
	Event *EventLinkAnyUri // Event containing the contract specifics and raw log

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
func (it *EventLinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkAnyUri)
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
		it.Event = new(EventLinkAnyUri)
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
func (it *EventLinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkAnyUri represents a LinkAnyUri event raised by the Event contract.
type EventLinkAnyUri struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkAnyUri is a free log retrieval operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterLinkAnyUri(opts *bind.FilterOpts, fromCharacterId []*big.Int) (*EventLinkAnyUriIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkAnyUriIterator{contract: _Event.contract, event: "LinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchLinkAnyUri is a free log subscription operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchLinkAnyUri(opts *bind.WatchOpts, sink chan<- *EventLinkAnyUri, fromCharacterId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkAnyUri)
				if err := _Event.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
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

// ParseLinkAnyUri is a log parse operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseLinkAnyUri(log types.Log) (*EventLinkAnyUri, error) {
	event := new(EventLinkAnyUri)
	if err := _Event.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkCharacterIterator is returned from FilterLinkCharacter and is used to iterate over the raw logs and unpacked data for LinkCharacter events raised by the Event contract.
type EventLinkCharacterIterator struct {
	Event *EventLinkCharacter // Event containing the contract specifics and raw log

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
func (it *EventLinkCharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkCharacter)
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
		it.Event = new(EventLinkCharacter)
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
func (it *EventLinkCharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkCharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkCharacter represents a LinkCharacter event raised by the Event contract.
type EventLinkCharacter struct {
	Account         common.Address
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkCharacter is a free log retrieval operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterLinkCharacter(opts *bind.FilterOpts, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (*EventLinkCharacterIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkCharacterIterator{contract: _Event.contract, event: "LinkCharacter", logs: logs, sub: sub}, nil
}

// WatchLinkCharacter is a free log subscription operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchLinkCharacter(opts *bind.WatchOpts, sink chan<- *EventLinkCharacter, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkCharacter)
				if err := _Event.contract.UnpackLog(event, "LinkCharacter", log); err != nil {
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

// ParseLinkCharacter is a log parse operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseLinkCharacter(log types.Log) (*EventLinkCharacter, error) {
	event := new(EventLinkCharacter)
	if err := _Event.contract.UnpackLog(event, "LinkCharacter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkCharacterLinkIterator is returned from FilterLinkCharacterLink and is used to iterate over the raw logs and unpacked data for LinkCharacterLink events raised by the Event contract.
type EventLinkCharacterLinkIterator struct {
	Event *EventLinkCharacterLink // Event containing the contract specifics and raw log

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
func (it *EventLinkCharacterLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkCharacterLink)
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
		it.Event = new(EventLinkCharacterLink)
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
func (it *EventLinkCharacterLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkCharacterLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkCharacterLink represents a LinkCharacterLink event raised by the Event contract.
type EventLinkCharacterLink struct {
	FromCharacterId   *big.Int
	LinkType          [32]byte
	ClFromCharacterId *big.Int
	ClToCharacterId   *big.Int
	ClLinkType        [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLinkCharacterLink is a free log retrieval operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) FilterLinkCharacterLink(opts *bind.FilterOpts, fromCharacterId []*big.Int, linkType [][32]byte) (*EventLinkCharacterLinkIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkCharacterLinkIterator{contract: _Event.contract, event: "LinkCharacterLink", logs: logs, sub: sub}, nil
}

// WatchLinkCharacterLink is a free log subscription operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) WatchLinkCharacterLink(opts *bind.WatchOpts, sink chan<- *EventLinkCharacterLink, fromCharacterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkCharacterLink)
				if err := _Event.contract.UnpackLog(event, "LinkCharacterLink", log); err != nil {
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

// ParseLinkCharacterLink is a log parse operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) ParseLinkCharacterLink(log types.Log) (*EventLinkCharacterLink, error) {
	event := new(EventLinkCharacterLink)
	if err := _Event.contract.UnpackLog(event, "LinkCharacterLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkERC721Iterator is returned from FilterLinkERC721 and is used to iterate over the raw logs and unpacked data for LinkERC721 events raised by the Event contract.
type EventLinkERC721Iterator struct {
	Event *EventLinkERC721 // Event containing the contract specifics and raw log

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
func (it *EventLinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkERC721)
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
		it.Event = new(EventLinkERC721)
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
func (it *EventLinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkERC721 represents a LinkERC721 event raised by the Event contract.
type EventLinkERC721 struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkERC721 is a free log retrieval operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterLinkERC721(opts *bind.FilterOpts, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*EventLinkERC721Iterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkERC721Iterator{contract: _Event.contract, event: "LinkERC721", logs: logs, sub: sub}, nil
}

// WatchLinkERC721 is a free log subscription operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchLinkERC721(opts *bind.WatchOpts, sink chan<- *EventLinkERC721, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkERC721)
				if err := _Event.contract.UnpackLog(event, "LinkERC721", log); err != nil {
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

// ParseLinkERC721 is a log parse operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseLinkERC721(log types.Log) (*EventLinkERC721, error) {
	event := new(EventLinkERC721)
	if err := _Event.contract.UnpackLog(event, "LinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkLinklistIterator is returned from FilterLinkLinklist and is used to iterate over the raw logs and unpacked data for LinkLinklist events raised by the Event contract.
type EventLinkLinklistIterator struct {
	Event *EventLinkLinklist // Event containing the contract specifics and raw log

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
func (it *EventLinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkLinklist)
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
		it.Event = new(EventLinkLinklist)
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
func (it *EventLinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkLinklist represents a LinkLinklist event raised by the Event contract.
type EventLinkLinklist struct {
	FromCharacterId *big.Int
	ToLinklistId    *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkLinklist is a free log retrieval operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) FilterLinkLinklist(opts *bind.FilterOpts, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*EventLinkLinklistIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkLinklistIterator{contract: _Event.contract, event: "LinkLinklist", logs: logs, sub: sub}, nil
}

// WatchLinkLinklist is a free log subscription operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) WatchLinkLinklist(opts *bind.WatchOpts, sink chan<- *EventLinkLinklist, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkLinklist)
				if err := _Event.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
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

// ParseLinkLinklist is a log parse operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) ParseLinkLinklist(log types.Log) (*EventLinkLinklist, error) {
	event := new(EventLinkLinklist)
	if err := _Event.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinkNoteIterator is returned from FilterLinkNote and is used to iterate over the raw logs and unpacked data for LinkNote events raised by the Event contract.
type EventLinkNoteIterator struct {
	Event *EventLinkNote // Event containing the contract specifics and raw log

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
func (it *EventLinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinkNote)
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
		it.Event = new(EventLinkNote)
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
func (it *EventLinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinkNote represents a LinkNote event raised by the Event contract.
type EventLinkNote struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkNote is a free log retrieval operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterLinkNote(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (*EventLinkNoteIterator, error) {

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

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLinkNoteIterator{contract: _Event.contract, event: "LinkNote", logs: logs, sub: sub}, nil
}

// WatchLinkNote is a free log subscription operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchLinkNote(opts *bind.WatchOpts, sink chan<- *EventLinkNote, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinkNote)
				if err := _Event.contract.UnpackLog(event, "LinkNote", log); err != nil {
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

// ParseLinkNote is a log parse operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseLinkNote(log types.Log) (*EventLinkNote, error) {
	event := new(EventLinkNote)
	if err := _Event.contract.UnpackLog(event, "LinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLinklistNFTInitializedIterator is returned from FilterLinklistNFTInitialized and is used to iterate over the raw logs and unpacked data for LinklistNFTInitialized events raised by the Event contract.
type EventLinklistNFTInitializedIterator struct {
	Event *EventLinklistNFTInitialized // Event containing the contract specifics and raw log

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
func (it *EventLinklistNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLinklistNFTInitialized)
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
		it.Event = new(EventLinklistNFTInitialized)
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
func (it *EventLinklistNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLinklistNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLinklistNFTInitialized represents a LinklistNFTInitialized event raised by the Event contract.
type EventLinklistNFTInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLinklistNFTInitialized is a free log retrieval operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Event *EventFilterer) FilterLinklistNFTInitialized(opts *bind.FilterOpts) (*EventLinklistNFTInitializedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &EventLinklistNFTInitializedIterator{contract: _Event.contract, event: "LinklistNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchLinklistNFTInitialized is a free log subscription operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Event *EventFilterer) WatchLinklistNFTInitialized(opts *bind.WatchOpts, sink chan<- *EventLinklistNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLinklistNFTInitialized)
				if err := _Event.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
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

// ParseLinklistNFTInitialized is a log parse operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Event *EventFilterer) ParseLinklistNFTInitialized(log types.Log) (*EventLinklistNFTInitialized, error) {
	event := new(EventLinklistNFTInitialized)
	if err := _Event.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventLockNoteIterator is returned from FilterLockNote and is used to iterate over the raw logs and unpacked data for LockNote events raised by the Event contract.
type EventLockNoteIterator struct {
	Event *EventLockNote // Event containing the contract specifics and raw log

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
func (it *EventLockNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventLockNote)
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
		it.Event = new(EventLockNote)
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
func (it *EventLockNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventLockNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventLockNote represents a LockNote event raised by the Event contract.
type EventLockNote struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockNote is a free log retrieval operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) FilterLockNote(opts *bind.FilterOpts, characterId []*big.Int) (*EventLockNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "LockNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventLockNoteIterator{contract: _Event.contract, event: "LockNote", logs: logs, sub: sub}, nil
}

// WatchLockNote is a free log subscription operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) WatchLockNote(opts *bind.WatchOpts, sink chan<- *EventLockNote, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "LockNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventLockNote)
				if err := _Event.contract.UnpackLog(event, "LockNote", log); err != nil {
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

// ParseLockNote is a log parse operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Event *EventFilterer) ParseLockNote(log types.Log) (*EventLockNote, error) {
	event := new(EventLockNote)
	if err := _Event.contract.UnpackLog(event, "LockNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventMintNFTInitializedIterator is returned from FilterMintNFTInitialized and is used to iterate over the raw logs and unpacked data for MintNFTInitialized events raised by the Event contract.
type EventMintNFTInitializedIterator struct {
	Event *EventMintNFTInitialized // Event containing the contract specifics and raw log

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
func (it *EventMintNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventMintNFTInitialized)
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
		it.Event = new(EventMintNFTInitialized)
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
func (it *EventMintNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventMintNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventMintNFTInitialized represents a MintNFTInitialized event raised by the Event contract.
type EventMintNFTInitialized struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintNFTInitialized is a free log retrieval operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Event *EventFilterer) FilterMintNFTInitialized(opts *bind.FilterOpts) (*EventMintNFTInitializedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &EventMintNFTInitializedIterator{contract: _Event.contract, event: "MintNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchMintNFTInitialized is a free log subscription operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Event *EventFilterer) WatchMintNFTInitialized(opts *bind.WatchOpts, sink chan<- *EventMintNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventMintNFTInitialized)
				if err := _Event.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
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

// ParseMintNFTInitialized is a log parse operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Event *EventFilterer) ParseMintNFTInitialized(log types.Log) (*EventMintNFTInitialized, error) {
	event := new(EventMintNFTInitialized)
	if err := _Event.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventMintNoteIterator is returned from FilterMintNote and is used to iterate over the raw logs and unpacked data for MintNote events raised by the Event contract.
type EventMintNoteIterator struct {
	Event *EventMintNote // Event containing the contract specifics and raw log

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
func (it *EventMintNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventMintNote)
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
		it.Event = new(EventMintNote)
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
func (it *EventMintNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventMintNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventMintNote represents a MintNote event raised by the Event contract.
type EventMintNote struct {
	To           common.Address
	CharacterId  *big.Int
	NoteId       *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintNote is a free log retrieval operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Event *EventFilterer) FilterMintNote(opts *bind.FilterOpts, to []common.Address, characterId []*big.Int, noteId []*big.Int) (*EventMintNoteIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "MintNote", toRule, characterIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return &EventMintNoteIterator{contract: _Event.contract, event: "MintNote", logs: logs, sub: sub}, nil
}

// WatchMintNote is a free log subscription operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Event *EventFilterer) WatchMintNote(opts *bind.WatchOpts, sink chan<- *EventMintNote, to []common.Address, characterId []*big.Int, noteId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "MintNote", toRule, characterIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventMintNote)
				if err := _Event.contract.UnpackLog(event, "MintNote", log); err != nil {
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

// ParseMintNote is a log parse operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Event *EventFilterer) ParseMintNote(log types.Log) (*EventMintNote, error) {
	event := new(EventMintNote)
	if err := _Event.contract.UnpackLog(event, "MintNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventPostNoteIterator is returned from FilterPostNote and is used to iterate over the raw logs and unpacked data for PostNote events raised by the Event contract.
type EventPostNoteIterator struct {
	Event *EventPostNote // Event containing the contract specifics and raw log

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
func (it *EventPostNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventPostNote)
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
		it.Event = new(EventPostNote)
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
func (it *EventPostNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventPostNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventPostNote represents a PostNote event raised by the Event contract.
type EventPostNote struct {
	CharacterId  *big.Int
	NoteId       *big.Int
	LinkKey      [32]byte
	LinkItemType [32]byte
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostNote is a free log retrieval operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Event *EventFilterer) FilterPostNote(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (*EventPostNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "PostNote", characterIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return &EventPostNoteIterator{contract: _Event.contract, event: "PostNote", logs: logs, sub: sub}, nil
}

// WatchPostNote is a free log subscription operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Event *EventFilterer) WatchPostNote(opts *bind.WatchOpts, sink chan<- *EventPostNote, characterId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "PostNote", characterIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventPostNote)
				if err := _Event.contract.UnpackLog(event, "PostNote", log); err != nil {
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

// ParsePostNote is a log parse operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Event *EventFilterer) ParsePostNote(log types.Log) (*EventPostNote, error) {
	event := new(EventPostNote)
	if err := _Event.contract.UnpackLog(event, "PostNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventRemoveOperatorIterator is returned from FilterRemoveOperator and is used to iterate over the raw logs and unpacked data for RemoveOperator events raised by the Event contract.
type EventRemoveOperatorIterator struct {
	Event *EventRemoveOperator // Event containing the contract specifics and raw log

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
func (it *EventRemoveOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventRemoveOperator)
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
		it.Event = new(EventRemoveOperator)
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
func (it *EventRemoveOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventRemoveOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventRemoveOperator represents a RemoveOperator event raised by the Event contract.
type EventRemoveOperator struct {
	CharacterId *big.Int
	Operator    common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveOperator is a free log retrieval operation binding the contract event 0xaa9506a57073a80893a2d2fdd53f4bd49d281a8548f483ad230f2d5da7f6710c.
//
// Solidity: event RemoveOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) FilterRemoveOperator(opts *bind.FilterOpts, characterId []*big.Int, operator []common.Address) (*EventRemoveOperatorIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "RemoveOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EventRemoveOperatorIterator{contract: _Event.contract, event: "RemoveOperator", logs: logs, sub: sub}, nil
}

// WatchRemoveOperator is a free log subscription operation binding the contract event 0xaa9506a57073a80893a2d2fdd53f4bd49d281a8548f483ad230f2d5da7f6710c.
//
// Solidity: event RemoveOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) WatchRemoveOperator(opts *bind.WatchOpts, sink chan<- *EventRemoveOperator, characterId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "RemoveOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventRemoveOperator)
				if err := _Event.contract.UnpackLog(event, "RemoveOperator", log); err != nil {
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

// ParseRemoveOperator is a log parse operation binding the contract event 0xaa9506a57073a80893a2d2fdd53f4bd49d281a8548f483ad230f2d5da7f6710c.
//
// Solidity: event RemoveOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) ParseRemoveOperator(log types.Log) (*EventRemoveOperator, error) {
	event := new(EventRemoveOperator)
	if err := _Event.contract.UnpackLog(event, "RemoveOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetCharacterUriIterator is returned from FilterSetCharacterUri and is used to iterate over the raw logs and unpacked data for SetCharacterUri events raised by the Event contract.
type EventSetCharacterUriIterator struct {
	Event *EventSetCharacterUri // Event containing the contract specifics and raw log

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
func (it *EventSetCharacterUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetCharacterUri)
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
		it.Event = new(EventSetCharacterUri)
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
func (it *EventSetCharacterUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetCharacterUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetCharacterUri represents a SetCharacterUri event raised by the Event contract.
type EventSetCharacterUri struct {
	CharacterId *big.Int
	NewUri      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetCharacterUri is a free log retrieval operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Event *EventFilterer) FilterSetCharacterUri(opts *bind.FilterOpts, characterId []*big.Int) (*EventSetCharacterUriIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetCharacterUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventSetCharacterUriIterator{contract: _Event.contract, event: "SetCharacterUri", logs: logs, sub: sub}, nil
}

// WatchSetCharacterUri is a free log subscription operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Event *EventFilterer) WatchSetCharacterUri(opts *bind.WatchOpts, sink chan<- *EventSetCharacterUri, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetCharacterUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetCharacterUri)
				if err := _Event.contract.UnpackLog(event, "SetCharacterUri", log); err != nil {
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

// ParseSetCharacterUri is a log parse operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Event *EventFilterer) ParseSetCharacterUri(log types.Log) (*EventSetCharacterUri, error) {
	event := new(EventSetCharacterUri)
	if err := _Event.contract.UnpackLog(event, "SetCharacterUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetHandleIterator is returned from FilterSetHandle and is used to iterate over the raw logs and unpacked data for SetHandle events raised by the Event contract.
type EventSetHandleIterator struct {
	Event *EventSetHandle // Event containing the contract specifics and raw log

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
func (it *EventSetHandleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetHandle)
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
		it.Event = new(EventSetHandle)
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
func (it *EventSetHandleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetHandleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetHandle represents a SetHandle event raised by the Event contract.
type EventSetHandle struct {
	Account     common.Address
	CharacterId *big.Int
	NewHandle   string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetHandle is a free log retrieval operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Event *EventFilterer) FilterSetHandle(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int) (*EventSetHandleIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetHandle", accountRule, characterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventSetHandleIterator{contract: _Event.contract, event: "SetHandle", logs: logs, sub: sub}, nil
}

// WatchSetHandle is a free log subscription operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Event *EventFilterer) WatchSetHandle(opts *bind.WatchOpts, sink chan<- *EventSetHandle, account []common.Address, characterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetHandle", accountRule, characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetHandle)
				if err := _Event.contract.UnpackLog(event, "SetHandle", log); err != nil {
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

// ParseSetHandle is a log parse operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Event *EventFilterer) ParseSetHandle(log types.Log) (*EventSetHandle, error) {
	event := new(EventSetHandle)
	if err := _Event.contract.UnpackLog(event, "SetHandle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetLinkModule4AddressIterator is returned from FilterSetLinkModule4Address and is used to iterate over the raw logs and unpacked data for SetLinkModule4Address events raised by the Event contract.
type EventSetLinkModule4AddressIterator struct {
	Event *EventSetLinkModule4Address // Event containing the contract specifics and raw log

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
func (it *EventSetLinkModule4AddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetLinkModule4Address)
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
		it.Event = new(EventSetLinkModule4Address)
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
func (it *EventSetLinkModule4AddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetLinkModule4AddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetLinkModule4Address represents a SetLinkModule4Address event raised by the Event contract.
type EventSetLinkModule4Address struct {
	Account    common.Address
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Address is a free log retrieval operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetLinkModule4Address(opts *bind.FilterOpts, account []common.Address, linkModule []common.Address) (*EventSetLinkModule4AddressIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetLinkModule4AddressIterator{contract: _Event.contract, event: "SetLinkModule4Address", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Address is a free log subscription operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetLinkModule4Address(opts *bind.WatchOpts, sink chan<- *EventSetLinkModule4Address, account []common.Address, linkModule []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetLinkModule4Address)
				if err := _Event.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
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

// ParseSetLinkModule4Address is a log parse operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetLinkModule4Address(log types.Log) (*EventSetLinkModule4Address, error) {
	event := new(EventSetLinkModule4Address)
	if err := _Event.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetLinkModule4CharacterIterator is returned from FilterSetLinkModule4Character and is used to iterate over the raw logs and unpacked data for SetLinkModule4Character events raised by the Event contract.
type EventSetLinkModule4CharacterIterator struct {
	Event *EventSetLinkModule4Character // Event containing the contract specifics and raw log

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
func (it *EventSetLinkModule4CharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetLinkModule4Character)
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
		it.Event = new(EventSetLinkModule4Character)
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
func (it *EventSetLinkModule4CharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetLinkModule4CharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetLinkModule4Character represents a SetLinkModule4Character event raised by the Event contract.
type EventSetLinkModule4Character struct {
	CharacterId *big.Int
	LinkModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Character is a free log retrieval operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetLinkModule4Character(opts *bind.FilterOpts, characterId []*big.Int, linkModule []common.Address) (*EventSetLinkModule4CharacterIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetLinkModule4Character", characterIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetLinkModule4CharacterIterator{contract: _Event.contract, event: "SetLinkModule4Character", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Character is a free log subscription operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetLinkModule4Character(opts *bind.WatchOpts, sink chan<- *EventSetLinkModule4Character, characterId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetLinkModule4Character", characterIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetLinkModule4Character)
				if err := _Event.contract.UnpackLog(event, "SetLinkModule4Character", log); err != nil {
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

// ParseSetLinkModule4Character is a log parse operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetLinkModule4Character(log types.Log) (*EventSetLinkModule4Character, error) {
	event := new(EventSetLinkModule4Character)
	if err := _Event.contract.UnpackLog(event, "SetLinkModule4Character", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetLinkModule4ERC721Iterator is returned from FilterSetLinkModule4ERC721 and is used to iterate over the raw logs and unpacked data for SetLinkModule4ERC721 events raised by the Event contract.
type EventSetLinkModule4ERC721Iterator struct {
	Event *EventSetLinkModule4ERC721 // Event containing the contract specifics and raw log

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
func (it *EventSetLinkModule4ERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetLinkModule4ERC721)
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
		it.Event = new(EventSetLinkModule4ERC721)
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
func (it *EventSetLinkModule4ERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetLinkModule4ERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetLinkModule4ERC721 represents a SetLinkModule4ERC721 event raised by the Event contract.
type EventSetLinkModule4ERC721 struct {
	TokenAddress common.Address
	TokenId      *big.Int
	LinkModule   common.Address
	ReturnData   []byte
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4ERC721 is a free log retrieval operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetLinkModule4ERC721(opts *bind.FilterOpts, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (*EventSetLinkModule4ERC721Iterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetLinkModule4ERC721Iterator{contract: _Event.contract, event: "SetLinkModule4ERC721", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4ERC721 is a free log subscription operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetLinkModule4ERC721(opts *bind.WatchOpts, sink chan<- *EventSetLinkModule4ERC721, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetLinkModule4ERC721)
				if err := _Event.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
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

// ParseSetLinkModule4ERC721 is a log parse operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetLinkModule4ERC721(log types.Log) (*EventSetLinkModule4ERC721, error) {
	event := new(EventSetLinkModule4ERC721)
	if err := _Event.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetLinkModule4LinklistIterator is returned from FilterSetLinkModule4Linklist and is used to iterate over the raw logs and unpacked data for SetLinkModule4Linklist events raised by the Event contract.
type EventSetLinkModule4LinklistIterator struct {
	Event *EventSetLinkModule4Linklist // Event containing the contract specifics and raw log

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
func (it *EventSetLinkModule4LinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetLinkModule4Linklist)
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
		it.Event = new(EventSetLinkModule4Linklist)
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
func (it *EventSetLinkModule4LinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetLinkModule4LinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetLinkModule4Linklist represents a SetLinkModule4Linklist event raised by the Event contract.
type EventSetLinkModule4Linklist struct {
	LinklistId *big.Int
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Linklist is a free log retrieval operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetLinkModule4Linklist(opts *bind.FilterOpts, linklistId []*big.Int, linkModule []common.Address) (*EventSetLinkModule4LinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetLinkModule4LinklistIterator{contract: _Event.contract, event: "SetLinkModule4Linklist", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Linklist is a free log subscription operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetLinkModule4Linklist(opts *bind.WatchOpts, sink chan<- *EventSetLinkModule4Linklist, linklistId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetLinkModule4Linklist)
				if err := _Event.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
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

// ParseSetLinkModule4Linklist is a log parse operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetLinkModule4Linklist(log types.Log) (*EventSetLinkModule4Linklist, error) {
	event := new(EventSetLinkModule4Linklist)
	if err := _Event.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetLinkModule4NoteIterator is returned from FilterSetLinkModule4Note and is used to iterate over the raw logs and unpacked data for SetLinkModule4Note events raised by the Event contract.
type EventSetLinkModule4NoteIterator struct {
	Event *EventSetLinkModule4Note // Event containing the contract specifics and raw log

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
func (it *EventSetLinkModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetLinkModule4Note)
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
		it.Event = new(EventSetLinkModule4Note)
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
func (it *EventSetLinkModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetLinkModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetLinkModule4Note represents a SetLinkModule4Note event raised by the Event contract.
type EventSetLinkModule4Note struct {
	CharacterId *big.Int
	NoteId      *big.Int
	LinkModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Note is a free log retrieval operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetLinkModule4Note(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, linkModule []common.Address) (*EventSetLinkModule4NoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetLinkModule4Note", characterIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetLinkModule4NoteIterator{contract: _Event.contract, event: "SetLinkModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Note is a free log subscription operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetLinkModule4Note(opts *bind.WatchOpts, sink chan<- *EventSetLinkModule4Note, characterId []*big.Int, noteId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetLinkModule4Note", characterIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetLinkModule4Note)
				if err := _Event.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
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

// ParseSetLinkModule4Note is a log parse operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetLinkModule4Note(log types.Log) (*EventSetLinkModule4Note, error) {
	event := new(EventSetLinkModule4Note)
	if err := _Event.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetMintModule4NoteIterator is returned from FilterSetMintModule4Note and is used to iterate over the raw logs and unpacked data for SetMintModule4Note events raised by the Event contract.
type EventSetMintModule4NoteIterator struct {
	Event *EventSetMintModule4Note // Event containing the contract specifics and raw log

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
func (it *EventSetMintModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetMintModule4Note)
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
		it.Event = new(EventSetMintModule4Note)
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
func (it *EventSetMintModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetMintModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetMintModule4Note represents a SetMintModule4Note event raised by the Event contract.
type EventSetMintModule4Note struct {
	CharacterId *big.Int
	NoteId      *big.Int
	MintModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMintModule4Note is a free log retrieval operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) FilterSetMintModule4Note(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, mintModule []common.Address) (*EventSetMintModule4NoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetMintModule4Note", characterIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return &EventSetMintModule4NoteIterator{contract: _Event.contract, event: "SetMintModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetMintModule4Note is a free log subscription operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) WatchSetMintModule4Note(opts *bind.WatchOpts, sink chan<- *EventSetMintModule4Note, characterId []*big.Int, noteId []*big.Int, mintModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetMintModule4Note", characterIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetMintModule4Note)
				if err := _Event.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
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

// ParseSetMintModule4Note is a log parse operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Event *EventFilterer) ParseSetMintModule4Note(log types.Log) (*EventSetMintModule4Note, error) {
	event := new(EventSetMintModule4Note)
	if err := _Event.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetNoteUriIterator is returned from FilterSetNoteUri and is used to iterate over the raw logs and unpacked data for SetNoteUri events raised by the Event contract.
type EventSetNoteUriIterator struct {
	Event *EventSetNoteUri // Event containing the contract specifics and raw log

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
func (it *EventSetNoteUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetNoteUri)
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
		it.Event = new(EventSetNoteUri)
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
func (it *EventSetNoteUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetNoteUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetNoteUri represents a SetNoteUri event raised by the Event contract.
type EventSetNoteUri struct {
	CharacterId *big.Int
	NoteId      *big.Int
	NewUri      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNoteUri is a free log retrieval operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Event *EventFilterer) FilterSetNoteUri(opts *bind.FilterOpts, characterId []*big.Int) (*EventSetNoteUriIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetNoteUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventSetNoteUriIterator{contract: _Event.contract, event: "SetNoteUri", logs: logs, sub: sub}, nil
}

// WatchSetNoteUri is a free log subscription operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Event *EventFilterer) WatchSetNoteUri(opts *bind.WatchOpts, sink chan<- *EventSetNoteUri, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetNoteUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetNoteUri)
				if err := _Event.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
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

// ParseSetNoteUri is a log parse operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Event *EventFilterer) ParseSetNoteUri(log types.Log) (*EventSetNoteUri, error) {
	event := new(EventSetNoteUri)
	if err := _Event.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the Event contract.
type EventSetOperatorIterator struct {
	Event *EventSetOperator // Event containing the contract specifics and raw log

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
func (it *EventSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetOperator)
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
		it.Event = new(EventSetOperator)
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
func (it *EventSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetOperator represents a SetOperator event raised by the Event contract.
type EventSetOperator struct {
	CharacterId *big.Int
	Operator    common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) FilterSetOperator(opts *bind.FilterOpts, characterId []*big.Int, operator []common.Address) (*EventSetOperatorIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &EventSetOperatorIterator{contract: _Event.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *EventSetOperator, characterId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetOperator)
				if err := _Event.contract.UnpackLog(event, "SetOperator", log); err != nil {
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

// ParseSetOperator is a log parse operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Event *EventFilterer) ParseSetOperator(log types.Log) (*EventSetOperator, error) {
	event := new(EventSetOperator)
	if err := _Event.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetPrimaryCharacterIdIterator is returned from FilterSetPrimaryCharacterId and is used to iterate over the raw logs and unpacked data for SetPrimaryCharacterId events raised by the Event contract.
type EventSetPrimaryCharacterIdIterator struct {
	Event *EventSetPrimaryCharacterId // Event containing the contract specifics and raw log

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
func (it *EventSetPrimaryCharacterIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetPrimaryCharacterId)
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
		it.Event = new(EventSetPrimaryCharacterId)
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
func (it *EventSetPrimaryCharacterIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetPrimaryCharacterIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetPrimaryCharacterId represents a SetPrimaryCharacterId event raised by the Event contract.
type EventSetPrimaryCharacterId struct {
	Account        common.Address
	CharacterId    *big.Int
	OldCharacterId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetPrimaryCharacterId is a free log retrieval operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Event *EventFilterer) FilterSetPrimaryCharacterId(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int, oldCharacterId []*big.Int) (*EventSetPrimaryCharacterIdIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var oldCharacterIdRule []interface{}
	for _, oldCharacterIdItem := range oldCharacterId {
		oldCharacterIdRule = append(oldCharacterIdRule, oldCharacterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetPrimaryCharacterId", accountRule, characterIdRule, oldCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventSetPrimaryCharacterIdIterator{contract: _Event.contract, event: "SetPrimaryCharacterId", logs: logs, sub: sub}, nil
}

// WatchSetPrimaryCharacterId is a free log subscription operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Event *EventFilterer) WatchSetPrimaryCharacterId(opts *bind.WatchOpts, sink chan<- *EventSetPrimaryCharacterId, account []common.Address, characterId []*big.Int, oldCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var oldCharacterIdRule []interface{}
	for _, oldCharacterIdItem := range oldCharacterId {
		oldCharacterIdRule = append(oldCharacterIdRule, oldCharacterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetPrimaryCharacterId", accountRule, characterIdRule, oldCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetPrimaryCharacterId)
				if err := _Event.contract.UnpackLog(event, "SetPrimaryCharacterId", log); err != nil {
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

// ParseSetPrimaryCharacterId is a log parse operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Event *EventFilterer) ParseSetPrimaryCharacterId(log types.Log) (*EventSetPrimaryCharacterId, error) {
	event := new(EventSetPrimaryCharacterId)
	if err := _Event.contract.UnpackLog(event, "SetPrimaryCharacterId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventSetSocialTokenIterator is returned from FilterSetSocialToken and is used to iterate over the raw logs and unpacked data for SetSocialToken events raised by the Event contract.
type EventSetSocialTokenIterator struct {
	Event *EventSetSocialToken // Event containing the contract specifics and raw log

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
func (it *EventSetSocialTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventSetSocialToken)
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
		it.Event = new(EventSetSocialToken)
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
func (it *EventSetSocialTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventSetSocialTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventSetSocialToken represents a SetSocialToken event raised by the Event contract.
type EventSetSocialToken struct {
	Account      common.Address
	CharacterId  *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetSocialToken is a free log retrieval operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Event *EventFilterer) FilterSetSocialToken(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int, tokenAddress []common.Address) (*EventSetSocialTokenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "SetSocialToken", accountRule, characterIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventSetSocialTokenIterator{contract: _Event.contract, event: "SetSocialToken", logs: logs, sub: sub}, nil
}

// WatchSetSocialToken is a free log subscription operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Event *EventFilterer) WatchSetSocialToken(opts *bind.WatchOpts, sink chan<- *EventSetSocialToken, account []common.Address, characterId []*big.Int, tokenAddress []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "SetSocialToken", accountRule, characterIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventSetSocialToken)
				if err := _Event.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
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

// ParseSetSocialToken is a log parse operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Event *EventFilterer) ParseSetSocialToken(log types.Log) (*EventSetSocialToken, error) {
	event := new(EventSetSocialToken)
	if err := _Event.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkAddressIterator is returned from FilterUnlinkAddress and is used to iterate over the raw logs and unpacked data for UnlinkAddress events raised by the Event contract.
type EventUnlinkAddressIterator struct {
	Event *EventUnlinkAddress // Event containing the contract specifics and raw log

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
func (it *EventUnlinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkAddress)
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
		it.Event = new(EventUnlinkAddress)
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
func (it *EventUnlinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkAddress represents a UnlinkAddress event raised by the Event contract.
type EventUnlinkAddress struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAddress is a free log retrieval operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Event *EventFilterer) FilterUnlinkAddress(opts *bind.FilterOpts, fromCharacterId []*big.Int, ethAddress []common.Address) (*EventUnlinkAddressIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkAddressIterator{contract: _Event.contract, event: "UnlinkAddress", logs: logs, sub: sub}, nil
}

// WatchUnlinkAddress is a free log subscription operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Event *EventFilterer) WatchUnlinkAddress(opts *bind.WatchOpts, sink chan<- *EventUnlinkAddress, fromCharacterId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkAddress)
				if err := _Event.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
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

// ParseUnlinkAddress is a log parse operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Event *EventFilterer) ParseUnlinkAddress(log types.Log) (*EventUnlinkAddress, error) {
	event := new(EventUnlinkAddress)
	if err := _Event.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkAnyUriIterator is returned from FilterUnlinkAnyUri and is used to iterate over the raw logs and unpacked data for UnlinkAnyUri events raised by the Event contract.
type EventUnlinkAnyUriIterator struct {
	Event *EventUnlinkAnyUri // Event containing the contract specifics and raw log

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
func (it *EventUnlinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkAnyUri)
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
		it.Event = new(EventUnlinkAnyUri)
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
func (it *EventUnlinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkAnyUri represents a UnlinkAnyUri event raised by the Event contract.
type EventUnlinkAnyUri struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAnyUri is a free log retrieval operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Event *EventFilterer) FilterUnlinkAnyUri(opts *bind.FilterOpts, fromCharacterId []*big.Int) (*EventUnlinkAnyUriIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkAnyUriIterator{contract: _Event.contract, event: "UnlinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchUnlinkAnyUri is a free log subscription operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Event *EventFilterer) WatchUnlinkAnyUri(opts *bind.WatchOpts, sink chan<- *EventUnlinkAnyUri, fromCharacterId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkAnyUri)
				if err := _Event.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
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

// ParseUnlinkAnyUri is a log parse operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Event *EventFilterer) ParseUnlinkAnyUri(log types.Log) (*EventUnlinkAnyUri, error) {
	event := new(EventUnlinkAnyUri)
	if err := _Event.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkCharacterIterator is returned from FilterUnlinkCharacter and is used to iterate over the raw logs and unpacked data for UnlinkCharacter events raised by the Event contract.
type EventUnlinkCharacterIterator struct {
	Event *EventUnlinkCharacter // Event containing the contract specifics and raw log

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
func (it *EventUnlinkCharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkCharacter)
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
		it.Event = new(EventUnlinkCharacter)
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
func (it *EventUnlinkCharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkCharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkCharacter represents a UnlinkCharacter event raised by the Event contract.
type EventUnlinkCharacter struct {
	Account         common.Address
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkCharacter is a free log retrieval operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Event *EventFilterer) FilterUnlinkCharacter(opts *bind.FilterOpts, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (*EventUnlinkCharacterIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkCharacterIterator{contract: _Event.contract, event: "UnlinkCharacter", logs: logs, sub: sub}, nil
}

// WatchUnlinkCharacter is a free log subscription operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Event *EventFilterer) WatchUnlinkCharacter(opts *bind.WatchOpts, sink chan<- *EventUnlinkCharacter, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkCharacter)
				if err := _Event.contract.UnpackLog(event, "UnlinkCharacter", log); err != nil {
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

// ParseUnlinkCharacter is a log parse operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Event *EventFilterer) ParseUnlinkCharacter(log types.Log) (*EventUnlinkCharacter, error) {
	event := new(EventUnlinkCharacter)
	if err := _Event.contract.UnpackLog(event, "UnlinkCharacter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkCharacterLinkIterator is returned from FilterUnlinkCharacterLink and is used to iterate over the raw logs and unpacked data for UnlinkCharacterLink events raised by the Event contract.
type EventUnlinkCharacterLinkIterator struct {
	Event *EventUnlinkCharacterLink // Event containing the contract specifics and raw log

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
func (it *EventUnlinkCharacterLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkCharacterLink)
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
		it.Event = new(EventUnlinkCharacterLink)
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
func (it *EventUnlinkCharacterLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkCharacterLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkCharacterLink represents a UnlinkCharacterLink event raised by the Event contract.
type EventUnlinkCharacterLink struct {
	FromCharacterId    *big.Int
	LinkType           [32]byte
	ClFromCharactereId *big.Int
	ClToCharacterId    *big.Int
	ClLinkType         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUnlinkCharacterLink is a free log retrieval operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) FilterUnlinkCharacterLink(opts *bind.FilterOpts, fromCharacterId []*big.Int, linkType [][32]byte) (*EventUnlinkCharacterLinkIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkCharacterLinkIterator{contract: _Event.contract, event: "UnlinkCharacterLink", logs: logs, sub: sub}, nil
}

// WatchUnlinkCharacterLink is a free log subscription operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) WatchUnlinkCharacterLink(opts *bind.WatchOpts, sink chan<- *EventUnlinkCharacterLink, fromCharacterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkCharacterLink)
				if err := _Event.contract.UnpackLog(event, "UnlinkCharacterLink", log); err != nil {
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

// ParseUnlinkCharacterLink is a log parse operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Event *EventFilterer) ParseUnlinkCharacterLink(log types.Log) (*EventUnlinkCharacterLink, error) {
	event := new(EventUnlinkCharacterLink)
	if err := _Event.contract.UnpackLog(event, "UnlinkCharacterLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkERC721Iterator is returned from FilterUnlinkERC721 and is used to iterate over the raw logs and unpacked data for UnlinkERC721 events raised by the Event contract.
type EventUnlinkERC721Iterator struct {
	Event *EventUnlinkERC721 // Event containing the contract specifics and raw log

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
func (it *EventUnlinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkERC721)
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
		it.Event = new(EventUnlinkERC721)
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
func (it *EventUnlinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkERC721 represents a UnlinkERC721 event raised by the Event contract.
type EventUnlinkERC721 struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkERC721 is a free log retrieval operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterUnlinkERC721(opts *bind.FilterOpts, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*EventUnlinkERC721Iterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkERC721Iterator{contract: _Event.contract, event: "UnlinkERC721", logs: logs, sub: sub}, nil
}

// WatchUnlinkERC721 is a free log subscription operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchUnlinkERC721(opts *bind.WatchOpts, sink chan<- *EventUnlinkERC721, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkERC721)
				if err := _Event.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
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

// ParseUnlinkERC721 is a log parse operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseUnlinkERC721(log types.Log) (*EventUnlinkERC721, error) {
	event := new(EventUnlinkERC721)
	if err := _Event.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkLinklistIterator is returned from FilterUnlinkLinklist and is used to iterate over the raw logs and unpacked data for UnlinkLinklist events raised by the Event contract.
type EventUnlinkLinklistIterator struct {
	Event *EventUnlinkLinklist // Event containing the contract specifics and raw log

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
func (it *EventUnlinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkLinklist)
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
		it.Event = new(EventUnlinkLinklist)
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
func (it *EventUnlinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkLinklist represents a UnlinkLinklist event raised by the Event contract.
type EventUnlinkLinklist struct {
	FromCharacterId *big.Int
	ToLinklistId    *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkLinklist is a free log retrieval operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) FilterUnlinkLinklist(opts *bind.FilterOpts, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*EventUnlinkLinklistIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkLinklistIterator{contract: _Event.contract, event: "UnlinkLinklist", logs: logs, sub: sub}, nil
}

// WatchUnlinkLinklist is a free log subscription operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) WatchUnlinkLinklist(opts *bind.WatchOpts, sink chan<- *EventUnlinkLinklist, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkLinklist)
				if err := _Event.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
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

// ParseUnlinkLinklist is a log parse operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Event *EventFilterer) ParseUnlinkLinklist(log types.Log) (*EventUnlinkLinklist, error) {
	event := new(EventUnlinkLinklist)
	if err := _Event.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventUnlinkNoteIterator is returned from FilterUnlinkNote and is used to iterate over the raw logs and unpacked data for UnlinkNote events raised by the Event contract.
type EventUnlinkNoteIterator struct {
	Event *EventUnlinkNote // Event containing the contract specifics and raw log

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
func (it *EventUnlinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventUnlinkNote)
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
		it.Event = new(EventUnlinkNote)
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
func (it *EventUnlinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventUnlinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventUnlinkNote represents a UnlinkNote event raised by the Event contract.
type EventUnlinkNote struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkNote is a free log retrieval operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) FilterUnlinkNote(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (*EventUnlinkNoteIterator, error) {

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

	logs, sub, err := _Event.contract.FilterLogs(opts, "UnlinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &EventUnlinkNoteIterator{contract: _Event.contract, event: "UnlinkNote", logs: logs, sub: sub}, nil
}

// WatchUnlinkNote is a free log subscription operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) WatchUnlinkNote(opts *bind.WatchOpts, sink chan<- *EventUnlinkNote, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Event.contract.WatchLogs(opts, "UnlinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventUnlinkNote)
				if err := _Event.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
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

// ParseUnlinkNote is a log parse operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Event *EventFilterer) ParseUnlinkNote(log types.Log) (*EventUnlinkNote, error) {
	event := new(EventUnlinkNote)
	if err := _Event.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventWeb3EntryInitializedIterator is returned from FilterWeb3EntryInitialized and is used to iterate over the raw logs and unpacked data for Web3EntryInitialized events raised by the Event contract.
type EventWeb3EntryInitializedIterator struct {
	Event *EventWeb3EntryInitialized // Event containing the contract specifics and raw log

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
func (it *EventWeb3EntryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventWeb3EntryInitialized)
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
		it.Event = new(EventWeb3EntryInitialized)
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
func (it *EventWeb3EntryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventWeb3EntryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventWeb3EntryInitialized represents a Web3EntryInitialized event raised by the Event contract.
type EventWeb3EntryInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWeb3EntryInitialized is a free log retrieval operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Event *EventFilterer) FilterWeb3EntryInitialized(opts *bind.FilterOpts) (*EventWeb3EntryInitializedIterator, error) {

	logs, sub, err := _Event.contract.FilterLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return &EventWeb3EntryInitializedIterator{contract: _Event.contract, event: "Web3EntryInitialized", logs: logs, sub: sub}, nil
}

// WatchWeb3EntryInitialized is a free log subscription operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Event *EventFilterer) WatchWeb3EntryInitialized(opts *bind.WatchOpts, sink chan<- *EventWeb3EntryInitialized) (event.Subscription, error) {

	logs, sub, err := _Event.contract.WatchLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventWeb3EntryInitialized)
				if err := _Event.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
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

// ParseWeb3EntryInitialized is a log parse operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Event *EventFilterer) ParseWeb3EntryInitialized(log types.Log) (*EventWeb3EntryInitialized, error) {
	event := new(EventWeb3EntryInitialized)
	if err := _Event.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
