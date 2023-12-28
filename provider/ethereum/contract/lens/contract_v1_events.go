// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lens

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

// V1EventsMetaData contains all meta data concerning the V1Events contract.
var V1EventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rootPubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Collected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CommentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DefaultProfileSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DispatcherSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EmergencyAdminSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"moduleGlobals\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FeeModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newPower\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDelegatedPowerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowNFTURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"followModuleDatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Followed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"approved\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool[]\",\"name\":\"enabled\",\"type\":\"bool[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"FollowsToggled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MirrorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleBaseConstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"prevWhitelisted\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsCurrencyWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsGovernanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"prevTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasuryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ModuleGlobalsTreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"collectModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"referenceModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PostCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"followModuleReturnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreatorWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileImageURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ReferenceModuleWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumEvents.ProtocolState\",\"name\":\"prevState\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumEvents.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StateSet\",\"type\":\"event\"}]",
}

// V1EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use V1EventsMetaData.ABI instead.
var V1EventsABI = V1EventsMetaData.ABI

// V1Events is an auto generated Go binding around an Ethereum contract.
type V1Events struct {
	V1EventsCaller     // Read-only binding to the contract
	V1EventsTransactor // Write-only binding to the contract
	V1EventsFilterer   // Log filterer for contract events
}

// V1EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1EventsSession struct {
	Contract     *V1Events         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1EventsCallerSession struct {
	Contract *V1EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// V1EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1EventsTransactorSession struct {
	Contract     *V1EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// V1EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1EventsRaw struct {
	Contract *V1Events // Generic contract binding to access the raw methods on
}

// V1EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1EventsCallerRaw struct {
	Contract *V1EventsCaller // Generic read-only contract binding to access the raw methods on
}

// V1EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1EventsTransactorRaw struct {
	Contract *V1EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1Events creates a new instance of V1Events, bound to a specific deployed contract.
func NewV1Events(address common.Address, backend bind.ContractBackend) (*V1Events, error) {
	contract, err := bindV1Events(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1Events{V1EventsCaller: V1EventsCaller{contract: contract}, V1EventsTransactor: V1EventsTransactor{contract: contract}, V1EventsFilterer: V1EventsFilterer{contract: contract}}, nil
}

// NewV1EventsCaller creates a new read-only instance of V1Events, bound to a specific deployed contract.
func NewV1EventsCaller(address common.Address, caller bind.ContractCaller) (*V1EventsCaller, error) {
	contract, err := bindV1Events(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1EventsCaller{contract: contract}, nil
}

// NewV1EventsTransactor creates a new write-only instance of V1Events, bound to a specific deployed contract.
func NewV1EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*V1EventsTransactor, error) {
	contract, err := bindV1Events(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1EventsTransactor{contract: contract}, nil
}

// NewV1EventsFilterer creates a new log filterer instance of V1Events, bound to a specific deployed contract.
func NewV1EventsFilterer(address common.Address, filterer bind.ContractFilterer) (*V1EventsFilterer, error) {
	contract, err := bindV1Events(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1EventsFilterer{contract: contract}, nil
}

// bindV1Events binds a generic wrapper to an already deployed contract.
func bindV1Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V1EventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Events *V1EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Events.Contract.V1EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Events *V1EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Events.Contract.V1EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Events *V1EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Events.Contract.V1EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Events *V1EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Events *V1EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Events *V1EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Events.Contract.contract.Transact(opts, method, params...)
}

// V1EventsBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the V1Events contract.
type V1EventsBaseInitializedIterator struct {
	Event *V1EventsBaseInitialized // Event containing the contract specifics and raw log

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
func (it *V1EventsBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsBaseInitialized)
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
		it.Event = new(V1EventsBaseInitialized)
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
func (it *V1EventsBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsBaseInitialized represents a BaseInitialized event raised by the V1Events contract.
type V1EventsBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*V1EventsBaseInitializedIterator, error) {

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &V1EventsBaseInitializedIterator{contract: _V1Events.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *V1EventsBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsBaseInitialized)
				if err := _V1Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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
func (_V1Events *V1EventsFilterer) ParseBaseInitialized(log types.Log) (*V1EventsBaseInitialized, error) {
	event := new(V1EventsBaseInitialized)
	if err := _V1Events.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCollectModuleWhitelistedIterator is returned from FilterCollectModuleWhitelisted and is used to iterate over the raw logs and unpacked data for CollectModuleWhitelisted events raised by the V1Events contract.
type V1EventsCollectModuleWhitelistedIterator struct {
	Event *V1EventsCollectModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *V1EventsCollectModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCollectModuleWhitelisted)
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
		it.Event = new(V1EventsCollectModuleWhitelisted)
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
func (it *V1EventsCollectModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCollectModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCollectModuleWhitelisted represents a CollectModuleWhitelisted event raised by the V1Events contract.
type V1EventsCollectModuleWhitelisted struct {
	CollectModule common.Address
	Whitelisted   bool
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCollectModuleWhitelisted is a free log retrieval operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCollectModuleWhitelisted(opts *bind.FilterOpts, collectModule []common.Address, whitelisted []bool) (*V1EventsCollectModuleWhitelistedIterator, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCollectModuleWhitelistedIterator{contract: _V1Events.contract, event: "CollectModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchCollectModuleWhitelisted is a free log subscription operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCollectModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *V1EventsCollectModuleWhitelisted, collectModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var collectModuleRule []interface{}
	for _, collectModuleItem := range collectModule {
		collectModuleRule = append(collectModuleRule, collectModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "CollectModuleWhitelisted", collectModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCollectModuleWhitelisted)
				if err := _V1Events.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
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

// ParseCollectModuleWhitelisted is a log parse operation binding the contract event 0x6cc19a794d6a439023150cd58748eed4353190c0bb060d2e6250e2df4a68b673.
//
// Solidity: event CollectModuleWhitelisted(address indexed collectModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCollectModuleWhitelisted(log types.Log) (*V1EventsCollectModuleWhitelisted, error) {
	event := new(V1EventsCollectModuleWhitelisted)
	if err := _V1Events.contract.UnpackLog(event, "CollectModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCollectNFTDeployedIterator is returned from FilterCollectNFTDeployed and is used to iterate over the raw logs and unpacked data for CollectNFTDeployed events raised by the V1Events contract.
type V1EventsCollectNFTDeployedIterator struct {
	Event *V1EventsCollectNFTDeployed // Event containing the contract specifics and raw log

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
func (it *V1EventsCollectNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCollectNFTDeployed)
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
		it.Event = new(V1EventsCollectNFTDeployed)
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
func (it *V1EventsCollectNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCollectNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCollectNFTDeployed represents a CollectNFTDeployed event raised by the V1Events contract.
type V1EventsCollectNFTDeployed struct {
	ProfileId  *big.Int
	PubId      *big.Int
	CollectNFT common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTDeployed is a free log retrieval operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCollectNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (*V1EventsCollectNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCollectNFTDeployedIterator{contract: _V1Events.contract, event: "CollectNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchCollectNFTDeployed is a free log subscription operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCollectNFTDeployed(opts *bind.WatchOpts, sink chan<- *V1EventsCollectNFTDeployed, profileId []*big.Int, pubId []*big.Int, collectNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTRule []interface{}
	for _, collectNFTItem := range collectNFT {
		collectNFTRule = append(collectNFTRule, collectNFTItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "CollectNFTDeployed", profileIdRule, pubIdRule, collectNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCollectNFTDeployed)
				if err := _V1Events.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
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

// ParseCollectNFTDeployed is a log parse operation binding the contract event 0x0b227b550ffed48af813b32e246f787e99581ee13206ba8f9d90d63615269b3f.
//
// Solidity: event CollectNFTDeployed(uint256 indexed profileId, uint256 indexed pubId, address indexed collectNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCollectNFTDeployed(log types.Log) (*V1EventsCollectNFTDeployed, error) {
	event := new(V1EventsCollectNFTDeployed)
	if err := _V1Events.contract.UnpackLog(event, "CollectNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCollectNFTInitializedIterator is returned from FilterCollectNFTInitialized and is used to iterate over the raw logs and unpacked data for CollectNFTInitialized events raised by the V1Events contract.
type V1EventsCollectNFTInitializedIterator struct {
	Event *V1EventsCollectNFTInitialized // Event containing the contract specifics and raw log

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
func (it *V1EventsCollectNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCollectNFTInitialized)
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
		it.Event = new(V1EventsCollectNFTInitialized)
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
func (it *V1EventsCollectNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCollectNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCollectNFTInitialized represents a CollectNFTInitialized event raised by the V1Events contract.
type V1EventsCollectNFTInitialized struct {
	ProfileId *big.Int
	PubId     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTInitialized is a free log retrieval operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCollectNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*V1EventsCollectNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCollectNFTInitializedIterator{contract: _V1Events.contract, event: "CollectNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchCollectNFTInitialized is a free log subscription operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCollectNFTInitialized(opts *bind.WatchOpts, sink chan<- *V1EventsCollectNFTInitialized, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "CollectNFTInitialized", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCollectNFTInitialized)
				if err := _V1Events.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
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

// ParseCollectNFTInitialized is a log parse operation binding the contract event 0x898a2dec95856255977a0fb48cebc30051d50c0d8d33f93dea1e3ddb2e342442.
//
// Solidity: event CollectNFTInitialized(uint256 indexed profileId, uint256 indexed pubId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCollectNFTInitialized(log types.Log) (*V1EventsCollectNFTInitialized, error) {
	event := new(V1EventsCollectNFTInitialized)
	if err := _V1Events.contract.UnpackLog(event, "CollectNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCollectNFTTransferredIterator is returned from FilterCollectNFTTransferred and is used to iterate over the raw logs and unpacked data for CollectNFTTransferred events raised by the V1Events contract.
type V1EventsCollectNFTTransferredIterator struct {
	Event *V1EventsCollectNFTTransferred // Event containing the contract specifics and raw log

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
func (it *V1EventsCollectNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCollectNFTTransferred)
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
		it.Event = new(V1EventsCollectNFTTransferred)
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
func (it *V1EventsCollectNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCollectNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCollectNFTTransferred represents a CollectNFTTransferred event raised by the V1Events contract.
type V1EventsCollectNFTTransferred struct {
	ProfileId    *big.Int
	PubId        *big.Int
	CollectNFTId *big.Int
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTTransferred is a free log retrieval operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCollectNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (*V1EventsCollectNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCollectNFTTransferredIterator{contract: _V1Events.contract, event: "CollectNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchCollectNFTTransferred is a free log subscription operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCollectNFTTransferred(opts *bind.WatchOpts, sink chan<- *V1EventsCollectNFTTransferred, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCollectNFTTransferred)
				if err := _V1Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
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

// ParseCollectNFTTransferred is a log parse operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCollectNFTTransferred(log types.Log) (*V1EventsCollectNFTTransferred, error) {
	event := new(V1EventsCollectNFTTransferred)
	if err := _V1Events.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCollectedIterator is returned from FilterCollected and is used to iterate over the raw logs and unpacked data for Collected events raised by the V1Events contract.
type V1EventsCollectedIterator struct {
	Event *V1EventsCollected // Event containing the contract specifics and raw log

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
func (it *V1EventsCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCollected)
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
		it.Event = new(V1EventsCollected)
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
func (it *V1EventsCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCollected represents a Collected event raised by the V1Events contract.
type V1EventsCollected struct {
	Collector         common.Address
	ProfileId         *big.Int
	PubId             *big.Int
	RootProfileId     *big.Int
	RootPubId         *big.Int
	CollectModuleData []byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollected is a free log retrieval operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCollected(opts *bind.FilterOpts, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (*V1EventsCollectedIterator, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCollectedIterator{contract: _V1Events.contract, event: "Collected", logs: logs, sub: sub}, nil
}

// WatchCollected is a free log subscription operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCollected(opts *bind.WatchOpts, sink chan<- *V1EventsCollected, collector []common.Address, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var collectorRule []interface{}
	for _, collectorItem := range collector {
		collectorRule = append(collectorRule, collectorItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "Collected", collectorRule, profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCollected)
				if err := _V1Events.contract.UnpackLog(event, "Collected", log); err != nil {
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

// ParseCollected is a log parse operation binding the contract event 0xed39bf0d9afa849610b901c9d7f4d00751ba20de2db023428065bec153833218.
//
// Solidity: event Collected(address indexed collector, uint256 indexed profileId, uint256 indexed pubId, uint256 rootProfileId, uint256 rootPubId, bytes collectModuleData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCollected(log types.Log) (*V1EventsCollected, error) {
	event := new(V1EventsCollected)
	if err := _V1Events.contract.UnpackLog(event, "Collected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsCommentCreatedIterator is returned from FilterCommentCreated and is used to iterate over the raw logs and unpacked data for CommentCreated events raised by the V1Events contract.
type V1EventsCommentCreatedIterator struct {
	Event *V1EventsCommentCreated // Event containing the contract specifics and raw log

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
func (it *V1EventsCommentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsCommentCreated)
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
		it.Event = new(V1EventsCommentCreated)
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
func (it *V1EventsCommentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsCommentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsCommentCreated represents a CommentCreated event raised by the V1Events contract.
type V1EventsCommentCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCommentCreated is a free log retrieval operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterCommentCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*V1EventsCommentCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsCommentCreatedIterator{contract: _V1Events.contract, event: "CommentCreated", logs: logs, sub: sub}, nil
}

// WatchCommentCreated is a free log subscription operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchCommentCreated(opts *bind.WatchOpts, sink chan<- *V1EventsCommentCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "CommentCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsCommentCreated)
				if err := _V1Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
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

// ParseCommentCreated is a log parse operation binding the contract event 0x7b4d1aa33773161799847429e4fbf29f56dbf1a3fe815f5070231cbfba402c37.
//
// Solidity: event CommentCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseCommentCreated(log types.Log) (*V1EventsCommentCreated, error) {
	event := new(V1EventsCommentCreated)
	if err := _V1Events.contract.UnpackLog(event, "CommentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsDefaultProfileSetIterator is returned from FilterDefaultProfileSet and is used to iterate over the raw logs and unpacked data for DefaultProfileSet events raised by the V1Events contract.
type V1EventsDefaultProfileSetIterator struct {
	Event *V1EventsDefaultProfileSet // Event containing the contract specifics and raw log

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
func (it *V1EventsDefaultProfileSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsDefaultProfileSet)
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
		it.Event = new(V1EventsDefaultProfileSet)
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
func (it *V1EventsDefaultProfileSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsDefaultProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsDefaultProfileSet represents a DefaultProfileSet event raised by the V1Events contract.
type V1EventsDefaultProfileSet struct {
	Wallet    common.Address
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDefaultProfileSet is a free log retrieval operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterDefaultProfileSet(opts *bind.FilterOpts, wallet []common.Address, profileId []*big.Int) (*V1EventsDefaultProfileSetIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsDefaultProfileSetIterator{contract: _V1Events.contract, event: "DefaultProfileSet", logs: logs, sub: sub}, nil
}

// WatchDefaultProfileSet is a free log subscription operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchDefaultProfileSet(opts *bind.WatchOpts, sink chan<- *V1EventsDefaultProfileSet, wallet []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "DefaultProfileSet", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsDefaultProfileSet)
				if err := _V1Events.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
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

// ParseDefaultProfileSet is a log parse operation binding the contract event 0x0afd7c479e8bc7dcdb856b3cc27d2332dfe1f018fde574ea124919ddcae8a933.
//
// Solidity: event DefaultProfileSet(address indexed wallet, uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseDefaultProfileSet(log types.Log) (*V1EventsDefaultProfileSet, error) {
	event := new(V1EventsDefaultProfileSet)
	if err := _V1Events.contract.UnpackLog(event, "DefaultProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsDispatcherSetIterator is returned from FilterDispatcherSet and is used to iterate over the raw logs and unpacked data for DispatcherSet events raised by the V1Events contract.
type V1EventsDispatcherSetIterator struct {
	Event *V1EventsDispatcherSet // Event containing the contract specifics and raw log

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
func (it *V1EventsDispatcherSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsDispatcherSet)
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
		it.Event = new(V1EventsDispatcherSet)
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
func (it *V1EventsDispatcherSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsDispatcherSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsDispatcherSet represents a DispatcherSet event raised by the V1Events contract.
type V1EventsDispatcherSet struct {
	ProfileId  *big.Int
	Dispatcher common.Address
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDispatcherSet is a free log retrieval operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterDispatcherSet(opts *bind.FilterOpts, profileId []*big.Int, dispatcher []common.Address) (*V1EventsDispatcherSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsDispatcherSetIterator{contract: _V1Events.contract, event: "DispatcherSet", logs: logs, sub: sub}, nil
}

// WatchDispatcherSet is a free log subscription operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchDispatcherSet(opts *bind.WatchOpts, sink chan<- *V1EventsDispatcherSet, profileId []*big.Int, dispatcher []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var dispatcherRule []interface{}
	for _, dispatcherItem := range dispatcher {
		dispatcherRule = append(dispatcherRule, dispatcherItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "DispatcherSet", profileIdRule, dispatcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsDispatcherSet)
				if err := _V1Events.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
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

// ParseDispatcherSet is a log parse operation binding the contract event 0x22baaec4952f35f59e45bd2ddb287e1ccc6d319375770c09428eb8f8d604e065.
//
// Solidity: event DispatcherSet(uint256 indexed profileId, address indexed dispatcher, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseDispatcherSet(log types.Log) (*V1EventsDispatcherSet, error) {
	event := new(V1EventsDispatcherSet)
	if err := _V1Events.contract.UnpackLog(event, "DispatcherSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsEmergencyAdminSetIterator is returned from FilterEmergencyAdminSet and is used to iterate over the raw logs and unpacked data for EmergencyAdminSet events raised by the V1Events contract.
type V1EventsEmergencyAdminSetIterator struct {
	Event *V1EventsEmergencyAdminSet // Event containing the contract specifics and raw log

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
func (it *V1EventsEmergencyAdminSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsEmergencyAdminSet)
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
		it.Event = new(V1EventsEmergencyAdminSet)
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
func (it *V1EventsEmergencyAdminSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsEmergencyAdminSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsEmergencyAdminSet represents a EmergencyAdminSet event raised by the V1Events contract.
type V1EventsEmergencyAdminSet struct {
	Caller            common.Address
	OldEmergencyAdmin common.Address
	NewEmergencyAdmin common.Address
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminSet is a free log retrieval operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterEmergencyAdminSet(opts *bind.FilterOpts, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (*V1EventsEmergencyAdminSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsEmergencyAdminSetIterator{contract: _V1Events.contract, event: "EmergencyAdminSet", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminSet is a free log subscription operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchEmergencyAdminSet(opts *bind.WatchOpts, sink chan<- *V1EventsEmergencyAdminSet, caller []common.Address, oldEmergencyAdmin []common.Address, newEmergencyAdmin []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var oldEmergencyAdminRule []interface{}
	for _, oldEmergencyAdminItem := range oldEmergencyAdmin {
		oldEmergencyAdminRule = append(oldEmergencyAdminRule, oldEmergencyAdminItem)
	}
	var newEmergencyAdminRule []interface{}
	for _, newEmergencyAdminItem := range newEmergencyAdmin {
		newEmergencyAdminRule = append(newEmergencyAdminRule, newEmergencyAdminItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "EmergencyAdminSet", callerRule, oldEmergencyAdminRule, newEmergencyAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsEmergencyAdminSet)
				if err := _V1Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
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

// ParseEmergencyAdminSet is a log parse operation binding the contract event 0x676c0801b0f400762e958ee31cfbb10870e70786f6761f57c8647e766b0db3d9.
//
// Solidity: event EmergencyAdminSet(address indexed caller, address indexed oldEmergencyAdmin, address indexed newEmergencyAdmin, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseEmergencyAdminSet(log types.Log) (*V1EventsEmergencyAdminSet, error) {
	event := new(V1EventsEmergencyAdminSet)
	if err := _V1Events.contract.UnpackLog(event, "EmergencyAdminSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFeeModuleBaseConstructedIterator is returned from FilterFeeModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for FeeModuleBaseConstructed events raised by the V1Events contract.
type V1EventsFeeModuleBaseConstructedIterator struct {
	Event *V1EventsFeeModuleBaseConstructed // Event containing the contract specifics and raw log

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
func (it *V1EventsFeeModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFeeModuleBaseConstructed)
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
		it.Event = new(V1EventsFeeModuleBaseConstructed)
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
func (it *V1EventsFeeModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFeeModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFeeModuleBaseConstructed represents a FeeModuleBaseConstructed event raised by the V1Events contract.
type V1EventsFeeModuleBaseConstructed struct {
	ModuleGlobals common.Address
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFeeModuleBaseConstructed is a free log retrieval operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFeeModuleBaseConstructed(opts *bind.FilterOpts, moduleGlobals []common.Address) (*V1EventsFeeModuleBaseConstructedIterator, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFeeModuleBaseConstructedIterator{contract: _V1Events.contract, event: "FeeModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchFeeModuleBaseConstructed is a free log subscription operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFeeModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *V1EventsFeeModuleBaseConstructed, moduleGlobals []common.Address) (event.Subscription, error) {

	var moduleGlobalsRule []interface{}
	for _, moduleGlobalsItem := range moduleGlobals {
		moduleGlobalsRule = append(moduleGlobalsRule, moduleGlobalsItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FeeModuleBaseConstructed", moduleGlobalsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFeeModuleBaseConstructed)
				if err := _V1Events.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
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

// ParseFeeModuleBaseConstructed is a log parse operation binding the contract event 0x4e84a529f4c627b5e787037d117873af1018768804cca3c7f0d47041fe2c89ed.
//
// Solidity: event FeeModuleBaseConstructed(address indexed moduleGlobals, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFeeModuleBaseConstructed(log types.Log) (*V1EventsFeeModuleBaseConstructed, error) {
	event := new(V1EventsFeeModuleBaseConstructed)
	if err := _V1Events.contract.UnpackLog(event, "FeeModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowModuleSetIterator is returned from FilterFollowModuleSet and is used to iterate over the raw logs and unpacked data for FollowModuleSet events raised by the V1Events contract.
type V1EventsFollowModuleSetIterator struct {
	Event *V1EventsFollowModuleSet // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowModuleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowModuleSet)
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
		it.Event = new(V1EventsFollowModuleSet)
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
func (it *V1EventsFollowModuleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowModuleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowModuleSet represents a FollowModuleSet event raised by the V1Events contract.
type V1EventsFollowModuleSet struct {
	ProfileId              *big.Int
	FollowModule           common.Address
	FollowModuleReturnData []byte
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleSet is a free log retrieval operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowModuleSet(opts *bind.FilterOpts, profileId []*big.Int) (*V1EventsFollowModuleSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowModuleSetIterator{contract: _V1Events.contract, event: "FollowModuleSet", logs: logs, sub: sub}, nil
}

// WatchFollowModuleSet is a free log subscription operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowModuleSet(opts *bind.WatchOpts, sink chan<- *V1EventsFollowModuleSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowModuleSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowModuleSet)
				if err := _V1Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
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

// ParseFollowModuleSet is a log parse operation binding the contract event 0x92d95e400932d129885e627b38b169cbb28443ffaaa282d0fba0cf8797721359.
//
// Solidity: event FollowModuleSet(uint256 indexed profileId, address followModule, bytes followModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowModuleSet(log types.Log) (*V1EventsFollowModuleSet, error) {
	event := new(V1EventsFollowModuleSet)
	if err := _V1Events.contract.UnpackLog(event, "FollowModuleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowModuleWhitelistedIterator is returned from FilterFollowModuleWhitelisted and is used to iterate over the raw logs and unpacked data for FollowModuleWhitelisted events raised by the V1Events contract.
type V1EventsFollowModuleWhitelistedIterator struct {
	Event *V1EventsFollowModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowModuleWhitelisted)
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
		it.Event = new(V1EventsFollowModuleWhitelisted)
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
func (it *V1EventsFollowModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowModuleWhitelisted represents a FollowModuleWhitelisted event raised by the V1Events contract.
type V1EventsFollowModuleWhitelisted struct {
	FollowModule common.Address
	Whitelisted  bool
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowModuleWhitelisted is a free log retrieval operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowModuleWhitelisted(opts *bind.FilterOpts, followModule []common.Address, whitelisted []bool) (*V1EventsFollowModuleWhitelistedIterator, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowModuleWhitelistedIterator{contract: _V1Events.contract, event: "FollowModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchFollowModuleWhitelisted is a free log subscription operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *V1EventsFollowModuleWhitelisted, followModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var followModuleRule []interface{}
	for _, followModuleItem := range followModule {
		followModuleRule = append(followModuleRule, followModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowModuleWhitelisted", followModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowModuleWhitelisted)
				if err := _V1Events.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
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

// ParseFollowModuleWhitelisted is a log parse operation binding the contract event 0x52c5b7889df9f12f84ec3da051e854e5876678370d8357959c23ef59dd6486df.
//
// Solidity: event FollowModuleWhitelisted(address indexed followModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowModuleWhitelisted(log types.Log) (*V1EventsFollowModuleWhitelisted, error) {
	event := new(V1EventsFollowModuleWhitelisted)
	if err := _V1Events.contract.UnpackLog(event, "FollowModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowNFTDelegatedPowerChangedIterator is returned from FilterFollowNFTDelegatedPowerChanged and is used to iterate over the raw logs and unpacked data for FollowNFTDelegatedPowerChanged events raised by the V1Events contract.
type V1EventsFollowNFTDelegatedPowerChangedIterator struct {
	Event *V1EventsFollowNFTDelegatedPowerChanged // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowNFTDelegatedPowerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowNFTDelegatedPowerChanged)
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
		it.Event = new(V1EventsFollowNFTDelegatedPowerChanged)
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
func (it *V1EventsFollowNFTDelegatedPowerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowNFTDelegatedPowerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowNFTDelegatedPowerChanged represents a FollowNFTDelegatedPowerChanged event raised by the V1Events contract.
type V1EventsFollowNFTDelegatedPowerChanged struct {
	Delegate  common.Address
	NewPower  *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDelegatedPowerChanged is a free log retrieval operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowNFTDelegatedPowerChanged(opts *bind.FilterOpts, delegate []common.Address, newPower []*big.Int) (*V1EventsFollowNFTDelegatedPowerChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowNFTDelegatedPowerChangedIterator{contract: _V1Events.contract, event: "FollowNFTDelegatedPowerChanged", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDelegatedPowerChanged is a free log subscription operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowNFTDelegatedPowerChanged(opts *bind.WatchOpts, sink chan<- *V1EventsFollowNFTDelegatedPowerChanged, delegate []common.Address, newPower []*big.Int) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}
	var newPowerRule []interface{}
	for _, newPowerItem := range newPower {
		newPowerRule = append(newPowerRule, newPowerItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowNFTDelegatedPowerChanged", delegateRule, newPowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowNFTDelegatedPowerChanged)
				if err := _V1Events.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
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

// ParseFollowNFTDelegatedPowerChanged is a log parse operation binding the contract event 0xd9a6070174f4ccca76ed4896432e9a090b16e07e8fe27f275f50b33500b98e52.
//
// Solidity: event FollowNFTDelegatedPowerChanged(address indexed delegate, uint256 indexed newPower, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowNFTDelegatedPowerChanged(log types.Log) (*V1EventsFollowNFTDelegatedPowerChanged, error) {
	event := new(V1EventsFollowNFTDelegatedPowerChanged)
	if err := _V1Events.contract.UnpackLog(event, "FollowNFTDelegatedPowerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowNFTDeployedIterator is returned from FilterFollowNFTDeployed and is used to iterate over the raw logs and unpacked data for FollowNFTDeployed events raised by the V1Events contract.
type V1EventsFollowNFTDeployedIterator struct {
	Event *V1EventsFollowNFTDeployed // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowNFTDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowNFTDeployed)
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
		it.Event = new(V1EventsFollowNFTDeployed)
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
func (it *V1EventsFollowNFTDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowNFTDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowNFTDeployed represents a FollowNFTDeployed event raised by the V1Events contract.
type V1EventsFollowNFTDeployed struct {
	ProfileId *big.Int
	FollowNFT common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTDeployed is a free log retrieval operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowNFTDeployed(opts *bind.FilterOpts, profileId []*big.Int, followNFT []common.Address) (*V1EventsFollowNFTDeployedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowNFTDeployedIterator{contract: _V1Events.contract, event: "FollowNFTDeployed", logs: logs, sub: sub}, nil
}

// WatchFollowNFTDeployed is a free log subscription operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowNFTDeployed(opts *bind.WatchOpts, sink chan<- *V1EventsFollowNFTDeployed, profileId []*big.Int, followNFT []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTRule []interface{}
	for _, followNFTItem := range followNFT {
		followNFTRule = append(followNFTRule, followNFTItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowNFTDeployed", profileIdRule, followNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowNFTDeployed)
				if err := _V1Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
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

// ParseFollowNFTDeployed is a log parse operation binding the contract event 0x44403e38baed5e40df7f64ff8708b076c75a0dfda8380e75df5c36f11a476743.
//
// Solidity: event FollowNFTDeployed(uint256 indexed profileId, address indexed followNFT, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowNFTDeployed(log types.Log) (*V1EventsFollowNFTDeployed, error) {
	event := new(V1EventsFollowNFTDeployed)
	if err := _V1Events.contract.UnpackLog(event, "FollowNFTDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowNFTInitializedIterator is returned from FilterFollowNFTInitialized and is used to iterate over the raw logs and unpacked data for FollowNFTInitialized events raised by the V1Events contract.
type V1EventsFollowNFTInitializedIterator struct {
	Event *V1EventsFollowNFTInitialized // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowNFTInitialized)
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
		it.Event = new(V1EventsFollowNFTInitialized)
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
func (it *V1EventsFollowNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowNFTInitialized represents a FollowNFTInitialized event raised by the V1Events contract.
type V1EventsFollowNFTInitialized struct {
	ProfileId *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTInitialized is a free log retrieval operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowNFTInitialized(opts *bind.FilterOpts, profileId []*big.Int) (*V1EventsFollowNFTInitializedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowNFTInitializedIterator{contract: _V1Events.contract, event: "FollowNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchFollowNFTInitialized is a free log subscription operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowNFTInitialized(opts *bind.WatchOpts, sink chan<- *V1EventsFollowNFTInitialized, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowNFTInitialized", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowNFTInitialized)
				if err := _V1Events.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
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

// ParseFollowNFTInitialized is a log parse operation binding the contract event 0xaec15127df11a6b562c87d31bcb8f4cd2f0cf57fb9b663d6334abf41fea94d95.
//
// Solidity: event FollowNFTInitialized(uint256 indexed profileId, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowNFTInitialized(log types.Log) (*V1EventsFollowNFTInitialized, error) {
	event := new(V1EventsFollowNFTInitialized)
	if err := _V1Events.contract.UnpackLog(event, "FollowNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowNFTTransferredIterator is returned from FilterFollowNFTTransferred and is used to iterate over the raw logs and unpacked data for FollowNFTTransferred events raised by the V1Events contract.
type V1EventsFollowNFTTransferredIterator struct {
	Event *V1EventsFollowNFTTransferred // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowNFTTransferred)
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
		it.Event = new(V1EventsFollowNFTTransferred)
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
func (it *V1EventsFollowNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowNFTTransferred represents a FollowNFTTransferred event raised by the V1Events contract.
type V1EventsFollowNFTTransferred struct {
	ProfileId   *big.Int
	FollowNFTId *big.Int
	From        common.Address
	To          common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTTransferred is a free log retrieval operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, followNFTId []*big.Int) (*V1EventsFollowNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowNFTTransferredIterator{contract: _V1Events.contract, event: "FollowNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchFollowNFTTransferred is a free log subscription operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowNFTTransferred(opts *bind.WatchOpts, sink chan<- *V1EventsFollowNFTTransferred, profileId []*big.Int, followNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var followNFTIdRule []interface{}
	for _, followNFTIdItem := range followNFTId {
		followNFTIdRule = append(followNFTIdRule, followNFTIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowNFTTransferred", profileIdRule, followNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowNFTTransferred)
				if err := _V1Events.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
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

// ParseFollowNFTTransferred is a log parse operation binding the contract event 0x4996ad2257e7db44908136c43128cc10ca988096f67dc6bb0bcee11d151368fb.
//
// Solidity: event FollowNFTTransferred(uint256 indexed profileId, uint256 indexed followNFTId, address from, address to, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowNFTTransferred(log types.Log) (*V1EventsFollowNFTTransferred, error) {
	event := new(V1EventsFollowNFTTransferred)
	if err := _V1Events.contract.UnpackLog(event, "FollowNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowNFTURISetIterator is returned from FilterFollowNFTURISet and is used to iterate over the raw logs and unpacked data for FollowNFTURISet events raised by the V1Events contract.
type V1EventsFollowNFTURISetIterator struct {
	Event *V1EventsFollowNFTURISet // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowNFTURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowNFTURISet)
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
		it.Event = new(V1EventsFollowNFTURISet)
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
func (it *V1EventsFollowNFTURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowNFTURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowNFTURISet represents a FollowNFTURISet event raised by the V1Events contract.
type V1EventsFollowNFTURISet struct {
	ProfileId    *big.Int
	FollowNFTURI string
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFollowNFTURISet is a free log retrieval operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowNFTURISet(opts *bind.FilterOpts, profileId []*big.Int) (*V1EventsFollowNFTURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowNFTURISetIterator{contract: _V1Events.contract, event: "FollowNFTURISet", logs: logs, sub: sub}, nil
}

// WatchFollowNFTURISet is a free log subscription operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowNFTURISet(opts *bind.WatchOpts, sink chan<- *V1EventsFollowNFTURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowNFTURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowNFTURISet)
				if err := _V1Events.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
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

// ParseFollowNFTURISet is a log parse operation binding the contract event 0xe82886e1af6fcab5caef13815b22f51384e970c367a785f265d13860a7d6966d.
//
// Solidity: event FollowNFTURISet(uint256 indexed profileId, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowNFTURISet(log types.Log) (*V1EventsFollowNFTURISet, error) {
	event := new(V1EventsFollowNFTURISet)
	if err := _V1Events.contract.UnpackLog(event, "FollowNFTURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowedIterator is returned from FilterFollowed and is used to iterate over the raw logs and unpacked data for Followed events raised by the V1Events contract.
type V1EventsFollowedIterator struct {
	Event *V1EventsFollowed // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowed)
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
		it.Event = new(V1EventsFollowed)
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
func (it *V1EventsFollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowed represents a Followed event raised by the V1Events contract.
type V1EventsFollowed struct {
	Follower          common.Address
	ProfileIds        []*big.Int
	FollowModuleDatas [][]byte
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFollowed is a free log retrieval operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowed(opts *bind.FilterOpts, follower []common.Address) (*V1EventsFollowedIterator, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowedIterator{contract: _V1Events.contract, event: "Followed", logs: logs, sub: sub}, nil
}

// WatchFollowed is a free log subscription operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowed(opts *bind.WatchOpts, sink chan<- *V1EventsFollowed, follower []common.Address) (event.Subscription, error) {

	var followerRule []interface{}
	for _, followerItem := range follower {
		followerRule = append(followerRule, followerItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "Followed", followerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowed)
				if err := _V1Events.contract.UnpackLog(event, "Followed", log); err != nil {
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

// ParseFollowed is a log parse operation binding the contract event 0x40487072dc56f384287d26fbe090f404143c2737d54632177451d1f74bd82c76.
//
// Solidity: event Followed(address indexed follower, uint256[] profileIds, bytes[] followModuleDatas, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowed(log types.Log) (*V1EventsFollowed, error) {
	event := new(V1EventsFollowed)
	if err := _V1Events.contract.UnpackLog(event, "Followed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowsApprovedIterator is returned from FilterFollowsApproved and is used to iterate over the raw logs and unpacked data for FollowsApproved events raised by the V1Events contract.
type V1EventsFollowsApprovedIterator struct {
	Event *V1EventsFollowsApproved // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowsApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowsApproved)
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
		it.Event = new(V1EventsFollowsApproved)
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
func (it *V1EventsFollowsApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowsApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowsApproved represents a FollowsApproved event raised by the V1Events contract.
type V1EventsFollowsApproved struct {
	Owner     common.Address
	ProfileId *big.Int
	Addresses []common.Address
	Approved  []bool
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFollowsApproved is a free log retrieval operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowsApproved(opts *bind.FilterOpts, owner []common.Address, profileId []*big.Int) (*V1EventsFollowsApprovedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowsApprovedIterator{contract: _V1Events.contract, event: "FollowsApproved", logs: logs, sub: sub}, nil
}

// WatchFollowsApproved is a free log subscription operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowsApproved(opts *bind.WatchOpts, sink chan<- *V1EventsFollowsApproved, owner []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowsApproved", ownerRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowsApproved)
				if err := _V1Events.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
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

// ParseFollowsApproved is a log parse operation binding the contract event 0xc67fc3972da5d6434ab7b796ba133c240d40ee4e69129963c5aa0f2a6f7c3ad6.
//
// Solidity: event FollowsApproved(address indexed owner, uint256 indexed profileId, address[] addresses, bool[] approved, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowsApproved(log types.Log) (*V1EventsFollowsApproved, error) {
	event := new(V1EventsFollowsApproved)
	if err := _V1Events.contract.UnpackLog(event, "FollowsApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsFollowsToggledIterator is returned from FilterFollowsToggled and is used to iterate over the raw logs and unpacked data for FollowsToggled events raised by the V1Events contract.
type V1EventsFollowsToggledIterator struct {
	Event *V1EventsFollowsToggled // Event containing the contract specifics and raw log

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
func (it *V1EventsFollowsToggledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsFollowsToggled)
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
		it.Event = new(V1EventsFollowsToggled)
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
func (it *V1EventsFollowsToggledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsFollowsToggledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsFollowsToggled represents a FollowsToggled event raised by the V1Events contract.
type V1EventsFollowsToggled struct {
	Owner      common.Address
	ProfileIds []*big.Int
	Enabled    []bool
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFollowsToggled is a free log retrieval operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterFollowsToggled(opts *bind.FilterOpts, owner []common.Address) (*V1EventsFollowsToggledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsFollowsToggledIterator{contract: _V1Events.contract, event: "FollowsToggled", logs: logs, sub: sub}, nil
}

// WatchFollowsToggled is a free log subscription operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchFollowsToggled(opts *bind.WatchOpts, sink chan<- *V1EventsFollowsToggled, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "FollowsToggled", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsFollowsToggled)
				if err := _V1Events.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
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

// ParseFollowsToggled is a log parse operation binding the contract event 0x5538c80c8d3bee397d87a7d153f7f085bb12adf2fe25a026c7cc4e83d8c5f1d7.
//
// Solidity: event FollowsToggled(address indexed owner, uint256[] profileIds, bool[] enabled, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseFollowsToggled(log types.Log) (*V1EventsFollowsToggled, error) {
	event := new(V1EventsFollowsToggled)
	if err := _V1Events.contract.UnpackLog(event, "FollowsToggled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsGovernanceSetIterator is returned from FilterGovernanceSet and is used to iterate over the raw logs and unpacked data for GovernanceSet events raised by the V1Events contract.
type V1EventsGovernanceSetIterator struct {
	Event *V1EventsGovernanceSet // Event containing the contract specifics and raw log

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
func (it *V1EventsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsGovernanceSet)
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
		it.Event = new(V1EventsGovernanceSet)
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
func (it *V1EventsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsGovernanceSet represents a GovernanceSet event raised by the V1Events contract.
type V1EventsGovernanceSet struct {
	Caller         common.Address
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterGovernanceSet is a free log retrieval operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterGovernanceSet(opts *bind.FilterOpts, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (*V1EventsGovernanceSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsGovernanceSetIterator{contract: _V1Events.contract, event: "GovernanceSet", logs: logs, sub: sub}, nil
}

// WatchGovernanceSet is a free log subscription operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchGovernanceSet(opts *bind.WatchOpts, sink chan<- *V1EventsGovernanceSet, caller []common.Address, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "GovernanceSet", callerRule, prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsGovernanceSet)
				if err := _V1Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
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

// ParseGovernanceSet is a log parse operation binding the contract event 0xe552a55455b740845a5c07ed445d1724142fc997b389835495a29b30cddc1ccd.
//
// Solidity: event GovernanceSet(address indexed caller, address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseGovernanceSet(log types.Log) (*V1EventsGovernanceSet, error) {
	event := new(V1EventsGovernanceSet)
	if err := _V1Events.contract.UnpackLog(event, "GovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsMirrorCreatedIterator is returned from FilterMirrorCreated and is used to iterate over the raw logs and unpacked data for MirrorCreated events raised by the V1Events contract.
type V1EventsMirrorCreatedIterator struct {
	Event *V1EventsMirrorCreated // Event containing the contract specifics and raw log

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
func (it *V1EventsMirrorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsMirrorCreated)
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
		it.Event = new(V1EventsMirrorCreated)
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
func (it *V1EventsMirrorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsMirrorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsMirrorCreated represents a MirrorCreated event raised by the V1Events contract.
type V1EventsMirrorCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ProfileIdPointed          *big.Int
	PubIdPointed              *big.Int
	ReferenceModuleData       []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterMirrorCreated is a free log retrieval operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterMirrorCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*V1EventsMirrorCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsMirrorCreatedIterator{contract: _V1Events.contract, event: "MirrorCreated", logs: logs, sub: sub}, nil
}

// WatchMirrorCreated is a free log subscription operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchMirrorCreated(opts *bind.WatchOpts, sink chan<- *V1EventsMirrorCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "MirrorCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsMirrorCreated)
				if err := _V1Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
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

// ParseMirrorCreated is a log parse operation binding the contract event 0x9ea5dedb85bd9da4e264ee5a39b7ba0982e5d4d035d55edfa98a36b00e770b5a.
//
// Solidity: event MirrorCreated(uint256 indexed profileId, uint256 indexed pubId, uint256 profileIdPointed, uint256 pubIdPointed, bytes referenceModuleData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseMirrorCreated(log types.Log) (*V1EventsMirrorCreated, error) {
	event := new(V1EventsMirrorCreated)
	if err := _V1Events.contract.UnpackLog(event, "MirrorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsModuleBaseConstructedIterator is returned from FilterModuleBaseConstructed and is used to iterate over the raw logs and unpacked data for ModuleBaseConstructed events raised by the V1Events contract.
type V1EventsModuleBaseConstructedIterator struct {
	Event *V1EventsModuleBaseConstructed // Event containing the contract specifics and raw log

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
func (it *V1EventsModuleBaseConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsModuleBaseConstructed)
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
		it.Event = new(V1EventsModuleBaseConstructed)
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
func (it *V1EventsModuleBaseConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsModuleBaseConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsModuleBaseConstructed represents a ModuleBaseConstructed event raised by the V1Events contract.
type V1EventsModuleBaseConstructed struct {
	Hub       common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModuleBaseConstructed is a free log retrieval operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterModuleBaseConstructed(opts *bind.FilterOpts, hub []common.Address) (*V1EventsModuleBaseConstructedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsModuleBaseConstructedIterator{contract: _V1Events.contract, event: "ModuleBaseConstructed", logs: logs, sub: sub}, nil
}

// WatchModuleBaseConstructed is a free log subscription operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchModuleBaseConstructed(opts *bind.WatchOpts, sink chan<- *V1EventsModuleBaseConstructed, hub []common.Address) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ModuleBaseConstructed", hubRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsModuleBaseConstructed)
				if err := _V1Events.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
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

// ParseModuleBaseConstructed is a log parse operation binding the contract event 0xf1a1fa6b64aa95186f5a1285e76198d0da80d9c5a88062641d447f1d7c54e56c.
//
// Solidity: event ModuleBaseConstructed(address indexed hub, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseModuleBaseConstructed(log types.Log) (*V1EventsModuleBaseConstructed, error) {
	event := new(V1EventsModuleBaseConstructed)
	if err := _V1Events.contract.UnpackLog(event, "ModuleBaseConstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsModuleGlobalsCurrencyWhitelistedIterator is returned from FilterModuleGlobalsCurrencyWhitelisted and is used to iterate over the raw logs and unpacked data for ModuleGlobalsCurrencyWhitelisted events raised by the V1Events contract.
type V1EventsModuleGlobalsCurrencyWhitelistedIterator struct {
	Event *V1EventsModuleGlobalsCurrencyWhitelisted // Event containing the contract specifics and raw log

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
func (it *V1EventsModuleGlobalsCurrencyWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsModuleGlobalsCurrencyWhitelisted)
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
		it.Event = new(V1EventsModuleGlobalsCurrencyWhitelisted)
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
func (it *V1EventsModuleGlobalsCurrencyWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsModuleGlobalsCurrencyWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsModuleGlobalsCurrencyWhitelisted represents a ModuleGlobalsCurrencyWhitelisted event raised by the V1Events contract.
type V1EventsModuleGlobalsCurrencyWhitelisted struct {
	Currency        common.Address
	PrevWhitelisted bool
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsCurrencyWhitelisted is a free log retrieval operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterModuleGlobalsCurrencyWhitelisted(opts *bind.FilterOpts, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (*V1EventsModuleGlobalsCurrencyWhitelistedIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsModuleGlobalsCurrencyWhitelistedIterator{contract: _V1Events.contract, event: "ModuleGlobalsCurrencyWhitelisted", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsCurrencyWhitelisted is a free log subscription operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchModuleGlobalsCurrencyWhitelisted(opts *bind.WatchOpts, sink chan<- *V1EventsModuleGlobalsCurrencyWhitelisted, currency []common.Address, prevWhitelisted []bool, whitelisted []bool) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var prevWhitelistedRule []interface{}
	for _, prevWhitelistedItem := range prevWhitelisted {
		prevWhitelistedRule = append(prevWhitelistedRule, prevWhitelistedItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ModuleGlobalsCurrencyWhitelisted", currencyRule, prevWhitelistedRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsModuleGlobalsCurrencyWhitelisted)
				if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
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

// ParseModuleGlobalsCurrencyWhitelisted is a log parse operation binding the contract event 0x79c3cefc851fd6040f06af202c542818d9fb39bcddcb7a7e3f563b15300a2743.
//
// Solidity: event ModuleGlobalsCurrencyWhitelisted(address indexed currency, bool indexed prevWhitelisted, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseModuleGlobalsCurrencyWhitelisted(log types.Log) (*V1EventsModuleGlobalsCurrencyWhitelisted, error) {
	event := new(V1EventsModuleGlobalsCurrencyWhitelisted)
	if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsCurrencyWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsModuleGlobalsGovernanceSetIterator is returned from FilterModuleGlobalsGovernanceSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsGovernanceSet events raised by the V1Events contract.
type V1EventsModuleGlobalsGovernanceSetIterator struct {
	Event *V1EventsModuleGlobalsGovernanceSet // Event containing the contract specifics and raw log

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
func (it *V1EventsModuleGlobalsGovernanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsModuleGlobalsGovernanceSet)
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
		it.Event = new(V1EventsModuleGlobalsGovernanceSet)
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
func (it *V1EventsModuleGlobalsGovernanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsModuleGlobalsGovernanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsModuleGlobalsGovernanceSet represents a ModuleGlobalsGovernanceSet event raised by the V1Events contract.
type V1EventsModuleGlobalsGovernanceSet struct {
	PrevGovernance common.Address
	NewGovernance  common.Address
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsGovernanceSet is a free log retrieval operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterModuleGlobalsGovernanceSet(opts *bind.FilterOpts, prevGovernance []common.Address, newGovernance []common.Address) (*V1EventsModuleGlobalsGovernanceSetIterator, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsModuleGlobalsGovernanceSetIterator{contract: _V1Events.contract, event: "ModuleGlobalsGovernanceSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsGovernanceSet is a free log subscription operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchModuleGlobalsGovernanceSet(opts *bind.WatchOpts, sink chan<- *V1EventsModuleGlobalsGovernanceSet, prevGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var prevGovernanceRule []interface{}
	for _, prevGovernanceItem := range prevGovernance {
		prevGovernanceRule = append(prevGovernanceRule, prevGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ModuleGlobalsGovernanceSet", prevGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsModuleGlobalsGovernanceSet)
				if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
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

// ParseModuleGlobalsGovernanceSet is a log parse operation binding the contract event 0xbf538a2c0db3d440906b8179dd0394f68a65b0b1481da70ffee24e19dccee84c.
//
// Solidity: event ModuleGlobalsGovernanceSet(address indexed prevGovernance, address indexed newGovernance, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseModuleGlobalsGovernanceSet(log types.Log) (*V1EventsModuleGlobalsGovernanceSet, error) {
	event := new(V1EventsModuleGlobalsGovernanceSet)
	if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsGovernanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsModuleGlobalsTreasuryFeeSetIterator is returned from FilterModuleGlobalsTreasuryFeeSet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasuryFeeSet events raised by the V1Events contract.
type V1EventsModuleGlobalsTreasuryFeeSetIterator struct {
	Event *V1EventsModuleGlobalsTreasuryFeeSet // Event containing the contract specifics and raw log

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
func (it *V1EventsModuleGlobalsTreasuryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsModuleGlobalsTreasuryFeeSet)
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
		it.Event = new(V1EventsModuleGlobalsTreasuryFeeSet)
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
func (it *V1EventsModuleGlobalsTreasuryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsModuleGlobalsTreasuryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsModuleGlobalsTreasuryFeeSet represents a ModuleGlobalsTreasuryFeeSet event raised by the V1Events contract.
type V1EventsModuleGlobalsTreasuryFeeSet struct {
	PrevTreasuryFee uint16
	NewTreasuryFee  uint16
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasuryFeeSet is a free log retrieval operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterModuleGlobalsTreasuryFeeSet(opts *bind.FilterOpts, prevTreasuryFee []uint16, newTreasuryFee []uint16) (*V1EventsModuleGlobalsTreasuryFeeSetIterator, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsModuleGlobalsTreasuryFeeSetIterator{contract: _V1Events.contract, event: "ModuleGlobalsTreasuryFeeSet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasuryFeeSet is a free log subscription operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchModuleGlobalsTreasuryFeeSet(opts *bind.WatchOpts, sink chan<- *V1EventsModuleGlobalsTreasuryFeeSet, prevTreasuryFee []uint16, newTreasuryFee []uint16) (event.Subscription, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ModuleGlobalsTreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsModuleGlobalsTreasuryFeeSet)
				if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
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

// ParseModuleGlobalsTreasuryFeeSet is a log parse operation binding the contract event 0xec936862e6bb897cd711a5f31825057583128c2a482f0a4c9a4e6c3fd7c023f4.
//
// Solidity: event ModuleGlobalsTreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseModuleGlobalsTreasuryFeeSet(log types.Log) (*V1EventsModuleGlobalsTreasuryFeeSet, error) {
	event := new(V1EventsModuleGlobalsTreasuryFeeSet)
	if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsTreasuryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsModuleGlobalsTreasurySetIterator is returned from FilterModuleGlobalsTreasurySet and is used to iterate over the raw logs and unpacked data for ModuleGlobalsTreasurySet events raised by the V1Events contract.
type V1EventsModuleGlobalsTreasurySetIterator struct {
	Event *V1EventsModuleGlobalsTreasurySet // Event containing the contract specifics and raw log

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
func (it *V1EventsModuleGlobalsTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsModuleGlobalsTreasurySet)
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
		it.Event = new(V1EventsModuleGlobalsTreasurySet)
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
func (it *V1EventsModuleGlobalsTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsModuleGlobalsTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsModuleGlobalsTreasurySet represents a ModuleGlobalsTreasurySet event raised by the V1Events contract.
type V1EventsModuleGlobalsTreasurySet struct {
	PrevTreasury common.Address
	NewTreasury  common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterModuleGlobalsTreasurySet is a free log retrieval operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterModuleGlobalsTreasurySet(opts *bind.FilterOpts, prevTreasury []common.Address, newTreasury []common.Address) (*V1EventsModuleGlobalsTreasurySetIterator, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsModuleGlobalsTreasurySetIterator{contract: _V1Events.contract, event: "ModuleGlobalsTreasurySet", logs: logs, sub: sub}, nil
}

// WatchModuleGlobalsTreasurySet is a free log subscription operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchModuleGlobalsTreasurySet(opts *bind.WatchOpts, sink chan<- *V1EventsModuleGlobalsTreasurySet, prevTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ModuleGlobalsTreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsModuleGlobalsTreasurySet)
				if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
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

// ParseModuleGlobalsTreasurySet is a log parse operation binding the contract event 0x3dfc53d6b49bfbc932b215ba515f0d0ab0e17aac17726fba48075f0c16c7ffe3.
//
// Solidity: event ModuleGlobalsTreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseModuleGlobalsTreasurySet(log types.Log) (*V1EventsModuleGlobalsTreasurySet, error) {
	event := new(V1EventsModuleGlobalsTreasurySet)
	if err := _V1Events.contract.UnpackLog(event, "ModuleGlobalsTreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsPostCreatedIterator is returned from FilterPostCreated and is used to iterate over the raw logs and unpacked data for PostCreated events raised by the V1Events contract.
type V1EventsPostCreatedIterator struct {
	Event *V1EventsPostCreated // Event containing the contract specifics and raw log

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
func (it *V1EventsPostCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsPostCreated)
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
		it.Event = new(V1EventsPostCreated)
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
func (it *V1EventsPostCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsPostCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsPostCreated represents a PostCreated event raised by the V1Events contract.
type V1EventsPostCreated struct {
	ProfileId                 *big.Int
	PubId                     *big.Int
	ContentURI                string
	CollectModule             common.Address
	CollectModuleReturnData   []byte
	ReferenceModule           common.Address
	ReferenceModuleReturnData []byte
	Timestamp                 *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterPostCreated is a free log retrieval operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterPostCreated(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int) (*V1EventsPostCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsPostCreatedIterator{contract: _V1Events.contract, event: "PostCreated", logs: logs, sub: sub}, nil
}

// WatchPostCreated is a free log subscription operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchPostCreated(opts *bind.WatchOpts, sink chan<- *V1EventsPostCreated, profileId []*big.Int, pubId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "PostCreated", profileIdRule, pubIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsPostCreated)
				if err := _V1Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
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

// ParsePostCreated is a log parse operation binding the contract event 0xc672c38b4d26c3c978228e99164105280410b144af24dd3ed8e4f9d211d96a50.
//
// Solidity: event PostCreated(uint256 indexed profileId, uint256 indexed pubId, string contentURI, address collectModule, bytes collectModuleReturnData, address referenceModule, bytes referenceModuleReturnData, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParsePostCreated(log types.Log) (*V1EventsPostCreated, error) {
	event := new(V1EventsPostCreated)
	if err := _V1Events.contract.UnpackLog(event, "PostCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the V1Events contract.
type V1EventsProfileCreatedIterator struct {
	Event *V1EventsProfileCreated // Event containing the contract specifics and raw log

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
func (it *V1EventsProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsProfileCreated)
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
		it.Event = new(V1EventsProfileCreated)
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
func (it *V1EventsProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsProfileCreated represents a ProfileCreated event raised by the V1Events contract.
type V1EventsProfileCreated struct {
	ProfileId              *big.Int
	Creator                common.Address
	To                     common.Address
	Handle                 string
	ImageURI               string
	FollowModule           common.Address
	FollowModuleReturnData []byte
	FollowNFTURI           string
	Timestamp              *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterProfileCreated(opts *bind.FilterOpts, profileId []*big.Int, creator []common.Address, to []common.Address) (*V1EventsProfileCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsProfileCreatedIterator{contract: _V1Events.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *V1EventsProfileCreated, profileId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsProfileCreated)
				if err := _V1Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0x4e14f57cff7910416f2ef43cf05019b5a97a313de71fec9344be11b9b88fed12.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, string imageURI, address followModule, bytes followModuleReturnData, string followNFTURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseProfileCreated(log types.Log) (*V1EventsProfileCreated, error) {
	event := new(V1EventsProfileCreated)
	if err := _V1Events.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsProfileCreatorWhitelistedIterator is returned from FilterProfileCreatorWhitelisted and is used to iterate over the raw logs and unpacked data for ProfileCreatorWhitelisted events raised by the V1Events contract.
type V1EventsProfileCreatorWhitelistedIterator struct {
	Event *V1EventsProfileCreatorWhitelisted // Event containing the contract specifics and raw log

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
func (it *V1EventsProfileCreatorWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsProfileCreatorWhitelisted)
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
		it.Event = new(V1EventsProfileCreatorWhitelisted)
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
func (it *V1EventsProfileCreatorWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsProfileCreatorWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsProfileCreatorWhitelisted represents a ProfileCreatorWhitelisted event raised by the V1Events contract.
type V1EventsProfileCreatorWhitelisted struct {
	ProfileCreator common.Address
	Whitelisted    bool
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfileCreatorWhitelisted is a free log retrieval operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterProfileCreatorWhitelisted(opts *bind.FilterOpts, profileCreator []common.Address, whitelisted []bool) (*V1EventsProfileCreatorWhitelistedIterator, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsProfileCreatorWhitelistedIterator{contract: _V1Events.contract, event: "ProfileCreatorWhitelisted", logs: logs, sub: sub}, nil
}

// WatchProfileCreatorWhitelisted is a free log subscription operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchProfileCreatorWhitelisted(opts *bind.WatchOpts, sink chan<- *V1EventsProfileCreatorWhitelisted, profileCreator []common.Address, whitelisted []bool) (event.Subscription, error) {

	var profileCreatorRule []interface{}
	for _, profileCreatorItem := range profileCreator {
		profileCreatorRule = append(profileCreatorRule, profileCreatorItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ProfileCreatorWhitelisted", profileCreatorRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsProfileCreatorWhitelisted)
				if err := _V1Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
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

// ParseProfileCreatorWhitelisted is a log parse operation binding the contract event 0x8f617843889b94892bd44852d36ca6a7f49ecf4350a01e7b68e22d80f4ed95bc.
//
// Solidity: event ProfileCreatorWhitelisted(address indexed profileCreator, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseProfileCreatorWhitelisted(log types.Log) (*V1EventsProfileCreatorWhitelisted, error) {
	event := new(V1EventsProfileCreatorWhitelisted)
	if err := _V1Events.contract.UnpackLog(event, "ProfileCreatorWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsProfileImageURISetIterator is returned from FilterProfileImageURISet and is used to iterate over the raw logs and unpacked data for ProfileImageURISet events raised by the V1Events contract.
type V1EventsProfileImageURISetIterator struct {
	Event *V1EventsProfileImageURISet // Event containing the contract specifics and raw log

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
func (it *V1EventsProfileImageURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsProfileImageURISet)
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
		it.Event = new(V1EventsProfileImageURISet)
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
func (it *V1EventsProfileImageURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsProfileImageURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsProfileImageURISet represents a ProfileImageURISet event raised by the V1Events contract.
type V1EventsProfileImageURISet struct {
	ProfileId *big.Int
	ImageURI  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileImageURISet is a free log retrieval operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterProfileImageURISet(opts *bind.FilterOpts, profileId []*big.Int) (*V1EventsProfileImageURISetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsProfileImageURISetIterator{contract: _V1Events.contract, event: "ProfileImageURISet", logs: logs, sub: sub}, nil
}

// WatchProfileImageURISet is a free log subscription operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchProfileImageURISet(opts *bind.WatchOpts, sink chan<- *V1EventsProfileImageURISet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ProfileImageURISet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsProfileImageURISet)
				if err := _V1Events.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
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

// ParseProfileImageURISet is a log parse operation binding the contract event 0xd5a5879cad33c830cc1432c1850107029a09c80c60e9bce1ecd08d24880bd46c.
//
// Solidity: event ProfileImageURISet(uint256 indexed profileId, string imageURI, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseProfileImageURISet(log types.Log) (*V1EventsProfileImageURISet, error) {
	event := new(V1EventsProfileImageURISet)
	if err := _V1Events.contract.UnpackLog(event, "ProfileImageURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsProfileMetadataSetIterator is returned from FilterProfileMetadataSet and is used to iterate over the raw logs and unpacked data for ProfileMetadataSet events raised by the V1Events contract.
type V1EventsProfileMetadataSetIterator struct {
	Event *V1EventsProfileMetadataSet // Event containing the contract specifics and raw log

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
func (it *V1EventsProfileMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsProfileMetadataSet)
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
		it.Event = new(V1EventsProfileMetadataSet)
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
func (it *V1EventsProfileMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsProfileMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsProfileMetadataSet represents a ProfileMetadataSet event raised by the V1Events contract.
type V1EventsProfileMetadataSet struct {
	ProfileId *big.Int
	Metadata  string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileMetadataSet is a free log retrieval operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterProfileMetadataSet(opts *bind.FilterOpts, profileId []*big.Int) (*V1EventsProfileMetadataSetIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsProfileMetadataSetIterator{contract: _V1Events.contract, event: "ProfileMetadataSet", logs: logs, sub: sub}, nil
}

// WatchProfileMetadataSet is a free log subscription operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchProfileMetadataSet(opts *bind.WatchOpts, sink chan<- *V1EventsProfileMetadataSet, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ProfileMetadataSet", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsProfileMetadataSet)
				if err := _V1Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
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

// ParseProfileMetadataSet is a log parse operation binding the contract event 0xf901a8b3832914a45999dd4c425fbe42eb4182724d394100401e633d9f0b286a.
//
// Solidity: event ProfileMetadataSet(uint256 indexed profileId, string metadata, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseProfileMetadataSet(log types.Log) (*V1EventsProfileMetadataSet, error) {
	event := new(V1EventsProfileMetadataSet)
	if err := _V1Events.contract.UnpackLog(event, "ProfileMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsReferenceModuleWhitelistedIterator is returned from FilterReferenceModuleWhitelisted and is used to iterate over the raw logs and unpacked data for ReferenceModuleWhitelisted events raised by the V1Events contract.
type V1EventsReferenceModuleWhitelistedIterator struct {
	Event *V1EventsReferenceModuleWhitelisted // Event containing the contract specifics and raw log

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
func (it *V1EventsReferenceModuleWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsReferenceModuleWhitelisted)
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
		it.Event = new(V1EventsReferenceModuleWhitelisted)
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
func (it *V1EventsReferenceModuleWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsReferenceModuleWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsReferenceModuleWhitelisted represents a ReferenceModuleWhitelisted event raised by the V1Events contract.
type V1EventsReferenceModuleWhitelisted struct {
	ReferenceModule common.Address
	Whitelisted     bool
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferenceModuleWhitelisted is a free log retrieval operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterReferenceModuleWhitelisted(opts *bind.FilterOpts, referenceModule []common.Address, whitelisted []bool) (*V1EventsReferenceModuleWhitelistedIterator, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsReferenceModuleWhitelistedIterator{contract: _V1Events.contract, event: "ReferenceModuleWhitelisted", logs: logs, sub: sub}, nil
}

// WatchReferenceModuleWhitelisted is a free log subscription operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchReferenceModuleWhitelisted(opts *bind.WatchOpts, sink chan<- *V1EventsReferenceModuleWhitelisted, referenceModule []common.Address, whitelisted []bool) (event.Subscription, error) {

	var referenceModuleRule []interface{}
	for _, referenceModuleItem := range referenceModule {
		referenceModuleRule = append(referenceModuleRule, referenceModuleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "ReferenceModuleWhitelisted", referenceModuleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsReferenceModuleWhitelisted)
				if err := _V1Events.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
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

// ParseReferenceModuleWhitelisted is a log parse operation binding the contract event 0x37872a053ef20cb52defb7c9ec20e1a87cb8dd5098ac9e76a144be263dfef572.
//
// Solidity: event ReferenceModuleWhitelisted(address indexed referenceModule, bool indexed whitelisted, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseReferenceModuleWhitelisted(log types.Log) (*V1EventsReferenceModuleWhitelisted, error) {
	event := new(V1EventsReferenceModuleWhitelisted)
	if err := _V1Events.contract.UnpackLog(event, "ReferenceModuleWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1EventsStateSetIterator is returned from FilterStateSet and is used to iterate over the raw logs and unpacked data for StateSet events raised by the V1Events contract.
type V1EventsStateSetIterator struct {
	Event *V1EventsStateSet // Event containing the contract specifics and raw log

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
func (it *V1EventsStateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1EventsStateSet)
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
		it.Event = new(V1EventsStateSet)
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
func (it *V1EventsStateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1EventsStateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1EventsStateSet represents a StateSet event raised by the V1Events contract.
type V1EventsStateSet struct {
	Caller    common.Address
	PrevState uint8
	NewState  uint8
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStateSet is a free log retrieval operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V1Events *V1EventsFilterer) FilterStateSet(opts *bind.FilterOpts, caller []common.Address, prevState []uint8, newState []uint8) (*V1EventsStateSetIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _V1Events.contract.FilterLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &V1EventsStateSetIterator{contract: _V1Events.contract, event: "StateSet", logs: logs, sub: sub}, nil
}

// WatchStateSet is a free log subscription operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V1Events *V1EventsFilterer) WatchStateSet(opts *bind.WatchOpts, sink chan<- *V1EventsStateSet, caller []common.Address, prevState []uint8, newState []uint8) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var prevStateRule []interface{}
	for _, prevStateItem := range prevState {
		prevStateRule = append(prevStateRule, prevStateItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _V1Events.contract.WatchLogs(opts, "StateSet", callerRule, prevStateRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1EventsStateSet)
				if err := _V1Events.contract.UnpackLog(event, "StateSet", log); err != nil {
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

// ParseStateSet is a log parse operation binding the contract event 0xa2f9a1499fc1f9b7796d21fe5761290ccb7e0ef6ccf35fa58b668f304a62a1ca.
//
// Solidity: event StateSet(address indexed caller, uint8 indexed prevState, uint8 indexed newState, uint256 timestamp)
func (_V1Events *V1EventsFilterer) ParseStateSet(log types.Log) (*V1EventsStateSet, error) {
	event := new(V1EventsStateSet)
	if err := _V1Events.contract.UnpackLog(event, "StateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
