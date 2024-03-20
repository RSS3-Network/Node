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

// SAVMBridgeMetaData contains all meta data concerning the SAVMBridge contract.
var SAVMBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dstRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"customData\",\"type\":\"bytes\"}],\"name\":\"bridgeOut\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"txUniqueIdentification\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dstRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"customData\",\"type\":\"bytes\"}],\"name\":\"getBridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SAVMBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use SAVMBridgeMetaData.ABI instead.
var SAVMBridgeABI = SAVMBridgeMetaData.ABI

// SAVMBridge is an auto generated Go binding around an Ethereum contract.
type SAVMBridge struct {
	SAVMBridgeCaller     // Read-only binding to the contract
	SAVMBridgeTransactor // Write-only binding to the contract
	SAVMBridgeFilterer   // Log filterer for contract events
}

// SAVMBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type SAVMBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SAVMBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SAVMBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SAVMBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SAVMBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SAVMBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SAVMBridgeSession struct {
	Contract     *SAVMBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SAVMBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SAVMBridgeCallerSession struct {
	Contract *SAVMBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SAVMBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SAVMBridgeTransactorSession struct {
	Contract     *SAVMBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SAVMBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type SAVMBridgeRaw struct {
	Contract *SAVMBridge // Generic contract binding to access the raw methods on
}

// SAVMBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SAVMBridgeCallerRaw struct {
	Contract *SAVMBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// SAVMBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SAVMBridgeTransactorRaw struct {
	Contract *SAVMBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSAVMBridge creates a new instance of SAVMBridge, bound to a specific deployed contract.
func NewSAVMBridge(address common.Address, backend bind.ContractBackend) (*SAVMBridge, error) {
	contract, err := bindSAVMBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SAVMBridge{SAVMBridgeCaller: SAVMBridgeCaller{contract: contract}, SAVMBridgeTransactor: SAVMBridgeTransactor{contract: contract}, SAVMBridgeFilterer: SAVMBridgeFilterer{contract: contract}}, nil
}

// NewSAVMBridgeCaller creates a new read-only instance of SAVMBridge, bound to a specific deployed contract.
func NewSAVMBridgeCaller(address common.Address, caller bind.ContractCaller) (*SAVMBridgeCaller, error) {
	contract, err := bindSAVMBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SAVMBridgeCaller{contract: contract}, nil
}

// NewSAVMBridgeTransactor creates a new write-only instance of SAVMBridge, bound to a specific deployed contract.
func NewSAVMBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*SAVMBridgeTransactor, error) {
	contract, err := bindSAVMBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SAVMBridgeTransactor{contract: contract}, nil
}

// NewSAVMBridgeFilterer creates a new log filterer instance of SAVMBridge, bound to a specific deployed contract.
func NewSAVMBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*SAVMBridgeFilterer, error) {
	contract, err := bindSAVMBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SAVMBridgeFilterer{contract: contract}, nil
}

// bindSAVMBridge binds a generic wrapper to an already deployed contract.
func bindSAVMBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SAVMBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SAVMBridge *SAVMBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SAVMBridge.Contract.SAVMBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SAVMBridge *SAVMBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SAVMBridge.Contract.SAVMBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SAVMBridge *SAVMBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SAVMBridge.Contract.SAVMBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SAVMBridge *SAVMBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SAVMBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SAVMBridge *SAVMBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SAVMBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SAVMBridge *SAVMBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SAVMBridge.Contract.contract.Transact(opts, method, params...)
}

// BridgeToken is a free data retrieval call binding the contract method 0xf4734b0c.
//
// Solidity: function bridgeToken() view returns(address)
func (_SAVMBridge *SAVMBridgeCaller) BridgeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SAVMBridge.contract.Call(opts, &out, "bridgeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeToken is a free data retrieval call binding the contract method 0xf4734b0c.
//
// Solidity: function bridgeToken() view returns(address)
func (_SAVMBridge *SAVMBridgeSession) BridgeToken() (common.Address, error) {
	return _SAVMBridge.Contract.BridgeToken(&_SAVMBridge.CallOpts)
}

// BridgeToken is a free data retrieval call binding the contract method 0xf4734b0c.
//
// Solidity: function bridgeToken() view returns(address)
func (_SAVMBridge *SAVMBridgeCallerSession) BridgeToken() (common.Address, error) {
	return _SAVMBridge.Contract.BridgeToken(&_SAVMBridge.CallOpts)
}

// GetBridgeFee is a free data retrieval call binding the contract method 0x95fa47b1.
//
// Solidity: function getBridgeFee(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) view returns(uint256 fee)
func (_SAVMBridge *SAVMBridgeCaller) GetBridgeFee(opts *bind.CallOpts, dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*big.Int, error) {
	var out []interface{}
	err := _SAVMBridge.contract.Call(opts, &out, "getBridgeFee", dstChainId, amount, dstRecipient, customData)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBridgeFee is a free data retrieval call binding the contract method 0x95fa47b1.
//
// Solidity: function getBridgeFee(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) view returns(uint256 fee)
func (_SAVMBridge *SAVMBridgeSession) GetBridgeFee(dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*big.Int, error) {
	return _SAVMBridge.Contract.GetBridgeFee(&_SAVMBridge.CallOpts, dstChainId, amount, dstRecipient, customData)
}

// GetBridgeFee is a free data retrieval call binding the contract method 0x95fa47b1.
//
// Solidity: function getBridgeFee(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) view returns(uint256 fee)
func (_SAVMBridge *SAVMBridgeCallerSession) GetBridgeFee(dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*big.Int, error) {
	return _SAVMBridge.Contract.GetBridgeFee(&_SAVMBridge.CallOpts, dstChainId, amount, dstRecipient, customData)
}

// BridgeOut is a paid mutator transaction binding the contract method 0xe43772fb.
//
// Solidity: function bridgeOut(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) payable returns(bytes32 txUniqueIdentification)
func (_SAVMBridge *SAVMBridgeTransactor) BridgeOut(opts *bind.TransactOpts, dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*types.Transaction, error) {
	return _SAVMBridge.contract.Transact(opts, "bridgeOut", dstChainId, amount, dstRecipient, customData)
}

// BridgeOut is a paid mutator transaction binding the contract method 0xe43772fb.
//
// Solidity: function bridgeOut(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) payable returns(bytes32 txUniqueIdentification)
func (_SAVMBridge *SAVMBridgeSession) BridgeOut(dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*types.Transaction, error) {
	return _SAVMBridge.Contract.BridgeOut(&_SAVMBridge.TransactOpts, dstChainId, amount, dstRecipient, customData)
}

// BridgeOut is a paid mutator transaction binding the contract method 0xe43772fb.
//
// Solidity: function bridgeOut(uint32 dstChainId, uint256 amount, address dstRecipient, bytes customData) payable returns(bytes32 txUniqueIdentification)
func (_SAVMBridge *SAVMBridgeTransactorSession) BridgeOut(dstChainId uint32, amount *big.Int, dstRecipient common.Address, customData []byte) (*types.Transaction, error) {
	return _SAVMBridge.Contract.BridgeOut(&_SAVMBridge.TransactOpts, dstChainId, amount, dstRecipient, customData)
}
