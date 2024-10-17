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

// InboxMetaData contains all meta data concerning the Inbox contract.
var InboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDataLength\",\"type\":\"uint256\"}],\"name\":\"DataTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GasLimitTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientSubmissionCost\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L1Forked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin\",\"type\":\"address\"}],\"name\":\"NotAllowedOrigin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotForked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOrigin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"NotRollupOrOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"RetryableData\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"AllowListAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isEnabled\",\"type\":\"bool\"}],\"name\":\"AllowListEnabledUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"InboxMessageDelivered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNum\",\"type\":\"uint256\"}],\"name\":\"InboxMessageDeliveredFromOrigin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allowListEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dataLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseFee\",\"type\":\"uint256\"}],\"name\":\"calculateRetryableSubmissionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createRetryableTicketNoRefundAliasRewrite\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"_bridge\",\"type\":\"address\"},{\"internalType\":\"contractISequencerInbox\",\"name\":\"_sequencerInbox\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedContractTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendL1FundedUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2Message\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messageData\",\"type\":\"bytes\"}],\"name\":\"sendL2MessageFromOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sendUnsignedTransactionToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"withdrawTo\",\"type\":\"address\"}],\"name\":\"sendWithdrawEthToFork\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInbox\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setAllowList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_allowListEnabled\",\"type\":\"bool\"}],\"name\":\"setAllowListEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"l2CallValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"excessFeeRefundAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"callValueRefundAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unsafeCreateRetryableTicket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// InboxABI is the input ABI used to generate the binding from.
// Deprecated: Use InboxMetaData.ABI instead.
var InboxABI = InboxMetaData.ABI

// Inbox is an auto generated Go binding around an Ethereum contract.
type Inbox struct {
	InboxCaller     // Read-only binding to the contract
	InboxTransactor // Write-only binding to the contract
	InboxFilterer   // Log filterer for contract events
}

// InboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type InboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InboxSession struct {
	Contract     *Inbox            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InboxCallerSession struct {
	Contract *InboxCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// InboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InboxTransactorSession struct {
	Contract     *InboxTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type InboxRaw struct {
	Contract *Inbox // Generic contract binding to access the raw methods on
}

// InboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InboxCallerRaw struct {
	Contract *InboxCaller // Generic read-only contract binding to access the raw methods on
}

// InboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InboxTransactorRaw struct {
	Contract *InboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInbox creates a new instance of Inbox, bound to a specific deployed contract.
func NewInbox(address common.Address, backend bind.ContractBackend) (*Inbox, error) {
	contract, err := bindInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Inbox{InboxCaller: InboxCaller{contract: contract}, InboxTransactor: InboxTransactor{contract: contract}, InboxFilterer: InboxFilterer{contract: contract}}, nil
}

// NewInboxCaller creates a new read-only instance of Inbox, bound to a specific deployed contract.
func NewInboxCaller(address common.Address, caller bind.ContractCaller) (*InboxCaller, error) {
	contract, err := bindInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InboxCaller{contract: contract}, nil
}

// NewInboxTransactor creates a new write-only instance of Inbox, bound to a specific deployed contract.
func NewInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*InboxTransactor, error) {
	contract, err := bindInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InboxTransactor{contract: contract}, nil
}

// NewInboxFilterer creates a new log filterer instance of Inbox, bound to a specific deployed contract.
func NewInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*InboxFilterer, error) {
	contract, err := bindInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InboxFilterer{contract: contract}, nil
}

// bindInbox binds a generic wrapper to an already deployed contract.
func bindInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.InboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.InboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Inbox *InboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Inbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Inbox *InboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Inbox *InboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Inbox.Contract.contract.Transact(opts, method, params...)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxCaller) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "allowListEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxSession) AllowListEnabled() (bool, error) {
	return _Inbox.Contract.AllowListEnabled(&_Inbox.CallOpts)
}

// AllowListEnabled is a free data retrieval call binding the contract method 0x22bd5c1c.
//
// Solidity: function allowListEnabled() view returns(bool)
func (_Inbox *InboxCallerSession) AllowListEnabled() (bool, error) {
	return _Inbox.Contract.AllowListEnabled(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_Inbox *InboxCallerSession) Bridge() (common.Address, error) {
	return _Inbox.Contract.Bridge(&_Inbox.CallOpts)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxCaller) CalculateRetryableSubmissionFee(opts *bind.CallOpts, dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "calculateRetryableSubmissionFee", dataLength, baseFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxSession) CalculateRetryableSubmissionFee(dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	return _Inbox.Contract.CalculateRetryableSubmissionFee(&_Inbox.CallOpts, dataLength, baseFee)
}

// CalculateRetryableSubmissionFee is a free data retrieval call binding the contract method 0xa66b327d.
//
// Solidity: function calculateRetryableSubmissionFee(uint256 dataLength, uint256 baseFee) view returns(uint256)
func (_Inbox *InboxCallerSession) CalculateRetryableSubmissionFee(dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	return _Inbox.Contract.CalculateRetryableSubmissionFee(&_Inbox.CallOpts, dataLength, baseFee)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxCaller) IsAllowed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "isAllowed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _Inbox.Contract.IsAllowed(&_Inbox.CallOpts, arg0)
}

// IsAllowed is a free data retrieval call binding the contract method 0xbabcc539.
//
// Solidity: function isAllowed(address ) view returns(bool)
func (_Inbox *InboxCallerSession) IsAllowed(arg0 common.Address) (bool, error) {
	return _Inbox.Contract.IsAllowed(&_Inbox.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Inbox *InboxCallerSession) Paused() (bool, error) {
	return _Inbox.Contract.Paused(&_Inbox.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxCaller) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Inbox.contract.Call(opts, &out, "sequencerInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxSession) SequencerInbox() (common.Address, error) {
	return _Inbox.Contract.SequencerInbox(&_Inbox.CallOpts)
}

// SequencerInbox is a free data retrieval call binding the contract method 0xee35f327.
//
// Solidity: function sequencerInbox() view returns(address)
func (_Inbox *InboxCallerSession) SequencerInbox() (common.Address, error) {
	return _Inbox.Contract.SequencerInbox(&_Inbox.CallOpts)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicket(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicket", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicket is a paid mutator transaction binding the contract method 0x679b6ded.
//
// Solidity: function createRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) CreateRetryableTicketNoRefundAliasRewrite(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "createRetryableTicketNoRefundAliasRewrite", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) CreateRetryableTicketNoRefundAliasRewrite(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// CreateRetryableTicketNoRefundAliasRewrite is a paid mutator transaction binding the contract method 0x1b871c8d.
//
// Solidity: function createRetryableTicketNoRefundAliasRewrite(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) CreateRetryableTicketNoRefundAliasRewrite(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.CreateRetryableTicketNoRefundAliasRewrite(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth", arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, arg0)
}

// DepositEth is a paid mutator transaction binding the contract method 0x0f4d14e9.
//
// Solidity: function depositEth(uint256 ) payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth(arg0 *big.Int) (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth(&_Inbox.TransactOpts, arg0)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxTransactor) DepositEth0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "depositEth0")
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxSession) DepositEth0() (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth0(&_Inbox.TransactOpts)
}

// DepositEth0 is a paid mutator transaction binding the contract method 0x439370b1.
//
// Solidity: function depositEth() payable returns(uint256)
func (_Inbox *InboxTransactorSession) DepositEth0() (*types.Transaction, error) {
	return _Inbox.Contract.DepositEth0(&_Inbox.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxTransactor) Initialize(opts *bind.TransactOpts, _bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "initialize", _bridge, _sequencerInbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxSession) Initialize(_bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _sequencerInbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _bridge, address _sequencerInbox) returns()
func (_Inbox *InboxTransactorSession) Initialize(_bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.Initialize(&_Inbox.TransactOpts, _bridge, _sequencerInbox)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Inbox *InboxTransactorSession) Pause() (*types.Transaction, error) {
	return _Inbox.Contract.Pause(&_Inbox.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address ) returns()
func (_Inbox *InboxTransactor) PostUpgradeInit(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "postUpgradeInit", arg0)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address ) returns()
func (_Inbox *InboxSession) PostUpgradeInit(arg0 common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.PostUpgradeInit(&_Inbox.TransactOpts, arg0)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0xc474d2c5.
//
// Solidity: function postUpgradeInit(address ) returns()
func (_Inbox *InboxTransactorSession) PostUpgradeInit(arg0 common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.PostUpgradeInit(&_Inbox.TransactOpts, arg0)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendContractTransaction", gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendContractTransaction is a paid mutator transaction binding the contract method 0x8a631aa6.
//
// Solidity: function sendContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, value, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedContractTransaction", gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedContractTransaction is a paid mutator transaction binding the contract method 0x5e916758.
//
// Solidity: function sendL1FundedContractTransaction(uint256 gasLimit, uint256 maxFeePerGas, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedContractTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedContractTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransaction is a paid mutator transaction binding the contract method 0x67ef3ab8.
//
// Solidity: function sendL1FundedUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) SendL1FundedUnsignedTransactionToFork(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL1FundedUnsignedTransactionToFork", gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) SendL1FundedUnsignedTransactionToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransactionToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL1FundedUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0xe6bd12cf.
//
// Solidity: function sendL1FundedUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) SendL1FundedUnsignedTransactionToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL1FundedUnsignedTransactionToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, data)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2Message", messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2Message is a paid mutator transaction binding the contract method 0xb75436bb.
//
// Solidity: function sendL2Message(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2Message(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2Message(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactor) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendL2MessageFromOrigin", messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendL2MessageFromOrigin is a paid mutator transaction binding the contract method 0x1fe927cf.
//
// Solidity: function sendL2MessageFromOrigin(bytes messageData) returns(uint256)
func (_Inbox *InboxTransactorSession) SendL2MessageFromOrigin(messageData []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendL2MessageFromOrigin(&_Inbox.TransactOpts, messageData)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransaction", gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransaction is a paid mutator transaction binding the contract method 0x5075788b.
//
// Solidity: function sendUnsignedTransaction(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendUnsignedTransaction(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransaction(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactor) SendUnsignedTransactionToFork(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendUnsignedTransactionToFork", gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxSession) SendUnsignedTransactionToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransactionToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendUnsignedTransactionToFork is a paid mutator transaction binding the contract method 0x00f72382.
//
// Solidity: function sendUnsignedTransactionToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, address to, uint256 value, bytes data) returns(uint256)
func (_Inbox *InboxTransactorSession) SendUnsignedTransactionToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.SendUnsignedTransactionToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, to, value, data)
}

// SendWithdrawEthToFork is a paid mutator transaction binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, uint256 value, address withdrawTo) returns(uint256)
func (_Inbox *InboxTransactor) SendWithdrawEthToFork(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, value *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "sendWithdrawEthToFork", gasLimit, maxFeePerGas, nonce, value, withdrawTo)
}

// SendWithdrawEthToFork is a paid mutator transaction binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, uint256 value, address withdrawTo) returns(uint256)
func (_Inbox *InboxSession) SendWithdrawEthToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, value *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SendWithdrawEthToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, value, withdrawTo)
}

// SendWithdrawEthToFork is a paid mutator transaction binding the contract method 0x70665f14.
//
// Solidity: function sendWithdrawEthToFork(uint256 gasLimit, uint256 maxFeePerGas, uint256 nonce, uint256 value, address withdrawTo) returns(uint256)
func (_Inbox *InboxTransactorSession) SendWithdrawEthToFork(gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, value *big.Int, withdrawTo common.Address) (*types.Transaction, error) {
	return _Inbox.Contract.SendWithdrawEthToFork(&_Inbox.TransactOpts, gasLimit, maxFeePerGas, nonce, value, withdrawTo)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxTransactor) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setAllowList", user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowList(&_Inbox.TransactOpts, user, val)
}

// SetAllowList is a paid mutator transaction binding the contract method 0xe3de72a5.
//
// Solidity: function setAllowList(address[] user, bool[] val) returns()
func (_Inbox *InboxTransactorSession) SetAllowList(user []common.Address, val []bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowList(&_Inbox.TransactOpts, user, val)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxTransactor) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "setAllowListEnabled", _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowListEnabled(&_Inbox.TransactOpts, _allowListEnabled)
}

// SetAllowListEnabled is a paid mutator transaction binding the contract method 0xefeadb6d.
//
// Solidity: function setAllowListEnabled(bool _allowListEnabled) returns()
func (_Inbox *InboxTransactorSession) SetAllowListEnabled(_allowListEnabled bool) (*types.Transaction, error) {
	return _Inbox.Contract.SetAllowListEnabled(&_Inbox.TransactOpts, _allowListEnabled)
}

// UniswapCreateRetryableTicket is a paid mutator transaction binding the contract method 0x1cb763d2.
//
// Solidity: function uniswapCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) UniswapCreateRetryableTicket(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "uniswapCreateRetryableTicket", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UniswapCreateRetryableTicket is a paid mutator transaction binding the contract method 0x1cb763d2.
//
// Solidity: function uniswapCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) UniswapCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UniswapCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UniswapCreateRetryableTicket is a paid mutator transaction binding the contract method 0x1cb763d2.
//
// Solidity: function uniswapCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) UniswapCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UniswapCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Inbox *InboxTransactorSession) Unpause() (*types.Transaction, error) {
	return _Inbox.Contract.Unpause(&_Inbox.TransactOpts)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactor) UnsafeCreateRetryableTicket(opts *bind.TransactOpts, to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.contract.Transact(opts, "unsafeCreateRetryableTicket", to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxSession) UnsafeCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UnsafeCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// UnsafeCreateRetryableTicket is a paid mutator transaction binding the contract method 0x6e6e8a6a.
//
// Solidity: function unsafeCreateRetryableTicket(address to, uint256 l2CallValue, uint256 maxSubmissionCost, address excessFeeRefundAddress, address callValueRefundAddress, uint256 gasLimit, uint256 maxFeePerGas, bytes data) payable returns(uint256)
func (_Inbox *InboxTransactorSession) UnsafeCreateRetryableTicket(to common.Address, l2CallValue *big.Int, maxSubmissionCost *big.Int, excessFeeRefundAddress common.Address, callValueRefundAddress common.Address, gasLimit *big.Int, maxFeePerGas *big.Int, data []byte) (*types.Transaction, error) {
	return _Inbox.Contract.UnsafeCreateRetryableTicket(&_Inbox.TransactOpts, to, l2CallValue, maxSubmissionCost, excessFeeRefundAddress, callValueRefundAddress, gasLimit, maxFeePerGas, data)
}

// InboxAllowListAddressSetIterator is returned from FilterAllowListAddressSet and is used to iterate over the raw logs and unpacked data for AllowListAddressSet events raised by the Inbox contract.
type InboxAllowListAddressSetIterator struct {
	Event *InboxAllowListAddressSet // Event containing the contract specifics and raw log

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
func (it *InboxAllowListAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxAllowListAddressSet)
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
		it.Event = new(InboxAllowListAddressSet)
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
func (it *InboxAllowListAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxAllowListAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxAllowListAddressSet represents a AllowListAddressSet event raised by the Inbox contract.
type InboxAllowListAddressSet struct {
	User common.Address
	Val  bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAllowListAddressSet is a free log retrieval operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) FilterAllowListAddressSet(opts *bind.FilterOpts, user []common.Address) (*InboxAllowListAddressSetIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return &InboxAllowListAddressSetIterator{contract: _Inbox.contract, event: "AllowListAddressSet", logs: logs, sub: sub}, nil
}

// WatchAllowListAddressSet is a free log subscription operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) WatchAllowListAddressSet(opts *bind.WatchOpts, sink chan<- *InboxAllowListAddressSet, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "AllowListAddressSet", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxAllowListAddressSet)
				if err := _Inbox.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
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

// ParseAllowListAddressSet is a log parse operation binding the contract event 0xd9739f45a01ce092c5cdb3d68f63d63d21676b1c6c0b4f9cbc6be4cf5449595a.
//
// Solidity: event AllowListAddressSet(address indexed user, bool val)
func (_Inbox *InboxFilterer) ParseAllowListAddressSet(log types.Log) (*InboxAllowListAddressSet, error) {
	event := new(InboxAllowListAddressSet)
	if err := _Inbox.contract.UnpackLog(event, "AllowListAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxAllowListEnabledUpdatedIterator is returned from FilterAllowListEnabledUpdated and is used to iterate over the raw logs and unpacked data for AllowListEnabledUpdated events raised by the Inbox contract.
type InboxAllowListEnabledUpdatedIterator struct {
	Event *InboxAllowListEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *InboxAllowListEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxAllowListEnabledUpdated)
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
		it.Event = new(InboxAllowListEnabledUpdated)
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
func (it *InboxAllowListEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxAllowListEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxAllowListEnabledUpdated represents a AllowListEnabledUpdated event raised by the Inbox contract.
type InboxAllowListEnabledUpdated struct {
	IsEnabled bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAllowListEnabledUpdated is a free log retrieval operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) FilterAllowListEnabledUpdated(opts *bind.FilterOpts) (*InboxAllowListEnabledUpdatedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &InboxAllowListEnabledUpdatedIterator{contract: _Inbox.contract, event: "AllowListEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowListEnabledUpdated is a free log subscription operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) WatchAllowListEnabledUpdated(opts *bind.WatchOpts, sink chan<- *InboxAllowListEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "AllowListEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxAllowListEnabledUpdated)
				if err := _Inbox.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
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

// ParseAllowListEnabledUpdated is a log parse operation binding the contract event 0x16435b45f7482047f839a6a19d291442627200f52cad2803c595150d0d440eb3.
//
// Solidity: event AllowListEnabledUpdated(bool isEnabled)
func (_Inbox *InboxFilterer) ParseAllowListEnabledUpdated(log types.Log) (*InboxAllowListEnabledUpdated, error) {
	event := new(InboxAllowListEnabledUpdated)
	if err := _Inbox.contract.UnpackLog(event, "AllowListEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredIterator is returned from FilterInboxMessageDelivered and is used to iterate over the raw logs and unpacked data for InboxMessageDelivered events raised by the Inbox contract.
type InboxInboxMessageDeliveredIterator struct {
	Event *InboxInboxMessageDelivered // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDelivered)
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
		it.Event = new(InboxInboxMessageDelivered)
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
func (it *InboxInboxMessageDeliveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDelivered represents a InboxMessageDelivered event raised by the Inbox contract.
type InboxInboxMessageDelivered struct {
	MessageNum *big.Int
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDelivered is a free log retrieval operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredIterator{contract: _Inbox.contract, event: "InboxMessageDelivered", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDelivered is a free log subscription operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDelivered", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDelivered)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
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

// ParseInboxMessageDelivered is a log parse operation binding the contract event 0xff64905f73a67fb594e0f940a8075a860db489ad991e032f48c81123eb52d60b.
//
// Solidity: event InboxMessageDelivered(uint256 indexed messageNum, bytes data)
func (_Inbox *InboxFilterer) ParseInboxMessageDelivered(log types.Log) (*InboxInboxMessageDelivered, error) {
	event := new(InboxInboxMessageDelivered)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDelivered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxInboxMessageDeliveredFromOriginIterator is returned from FilterInboxMessageDeliveredFromOrigin and is used to iterate over the raw logs and unpacked data for InboxMessageDeliveredFromOrigin events raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOriginIterator struct {
	Event *InboxInboxMessageDeliveredFromOrigin // Event containing the contract specifics and raw log

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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
		it.Event = new(InboxInboxMessageDeliveredFromOrigin)
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
func (it *InboxInboxMessageDeliveredFromOriginIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxInboxMessageDeliveredFromOriginIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxInboxMessageDeliveredFromOrigin represents a InboxMessageDeliveredFromOrigin event raised by the Inbox contract.
type InboxInboxMessageDeliveredFromOrigin struct {
	MessageNum *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInboxMessageDeliveredFromOrigin is a free log retrieval operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*InboxInboxMessageDeliveredFromOriginIterator, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return &InboxInboxMessageDeliveredFromOriginIterator{contract: _Inbox.contract, event: "InboxMessageDeliveredFromOrigin", logs: logs, sub: sub}, nil
}

// WatchInboxMessageDeliveredFromOrigin is a free log subscription operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *InboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {

	var messageNumRule []interface{}
	for _, messageNumItem := range messageNum {
		messageNumRule = append(messageNumRule, messageNumItem)
	}

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "InboxMessageDeliveredFromOrigin", messageNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxInboxMessageDeliveredFromOrigin)
				if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
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

// ParseInboxMessageDeliveredFromOrigin is a log parse operation binding the contract event 0xab532385be8f1005a4b6ba8fa20a2245facb346134ac739fe9a5198dc1580b9c.
//
// Solidity: event InboxMessageDeliveredFromOrigin(uint256 indexed messageNum)
func (_Inbox *InboxFilterer) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*InboxInboxMessageDeliveredFromOrigin, error) {
	event := new(InboxInboxMessageDeliveredFromOrigin)
	if err := _Inbox.contract.UnpackLog(event, "InboxMessageDeliveredFromOrigin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Inbox contract.
type InboxPausedIterator struct {
	Event *InboxPaused // Event containing the contract specifics and raw log

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
func (it *InboxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxPaused)
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
		it.Event = new(InboxPaused)
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
func (it *InboxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxPaused represents a Paused event raised by the Inbox contract.
type InboxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) FilterPaused(opts *bind.FilterOpts) (*InboxPausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &InboxPausedIterator{contract: _Inbox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *InboxPaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxPaused)
				if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Inbox *InboxFilterer) ParsePaused(log types.Log) (*InboxPaused, error) {
	event := new(InboxPaused)
	if err := _Inbox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InboxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Inbox contract.
type InboxUnpausedIterator struct {
	Event *InboxUnpaused // Event containing the contract specifics and raw log

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
func (it *InboxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InboxUnpaused)
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
		it.Event = new(InboxUnpaused)
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
func (it *InboxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InboxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InboxUnpaused represents a Unpaused event raised by the Inbox contract.
type InboxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*InboxUnpausedIterator, error) {

	logs, sub, err := _Inbox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &InboxUnpausedIterator{contract: _Inbox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *InboxUnpaused) (event.Subscription, error) {

	logs, sub, err := _Inbox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InboxUnpaused)
				if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Inbox *InboxFilterer) ParseUnpaused(log types.Log) (*InboxUnpaused, error) {
	event := new(InboxUnpaused)
	if err := _Inbox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
