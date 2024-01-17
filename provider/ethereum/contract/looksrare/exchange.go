// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package looksrare

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

// OrderTypesMakerOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesMakerOrder struct {
	IsOrderAsk         bool
	Signer             common.Address
	Collection         common.Address
	Price              *big.Int
	TokenId            *big.Int
	Amount             *big.Int
	Strategy           common.Address
	Currency           common.Address
	Nonce              *big.Int
	StartTime          *big.Int
	EndTime            *big.Int
	MinPercentageToAsk *big.Int
	Params             []byte
	V                  uint8
	R                  [32]byte
	S                  [32]byte
}

// OrderTypesTakerOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderTypesTakerOrder struct {
	IsOrderAsk         bool
	Taker              common.Address
	Price              *big.Int
	TokenId            *big.Int
	MinPercentageToAsk *big.Int
	Params             []byte
}

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currencyManager\",\"type\":\"address\"}],\"name\":\"NewCurrencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executionManager\",\"type\":\"address\"}],\"name\":\"NewExecutionManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"NewRoyaltyFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferSelectorNFT\",\"type\":\"address\"}],\"name\":\"NewTransferSelectorNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltyPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerAsk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerBid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNonce\",\"type\":\"uint256\"}],\"name\":\"cancelAllOrdersForSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultipleMakerOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currencyManager\",\"outputs\":[{\"internalType\":\"contractICurrencyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionManager\",\"outputs\":[{\"internalType\":\"contractIExecutionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"isUserOrderNonceExecutedOrCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBidUsingETHAndWETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerBid\",\"type\":\"tuple\"}],\"name\":\"matchBidWithTakerAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyFeeManager\",\"outputs\":[{\"internalType\":\"contractIRoyaltyFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferSelectorNFT\",\"outputs\":[{\"internalType\":\"contractITransferSelectorNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"}],\"name\":\"updateCurrencyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"}],\"name\":\"updateExecutionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateProtocolFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"updateRoyaltyFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transferSelectorNFT\",\"type\":\"address\"}],\"name\":\"updateTransferSelectorNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMinOrderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Exchange *ExchangeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Exchange *ExchangeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Exchange.Contract.DOMAINSEPARATOR(&_Exchange.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Exchange.Contract.DOMAINSEPARATOR(&_Exchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeSession) WETH() (common.Address, error) {
	return _Exchange.Contract.WETH(&_Exchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Exchange *ExchangeCallerSession) WETH() (common.Address, error) {
	return _Exchange.Contract.WETH(&_Exchange.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Exchange *ExchangeCaller) CurrencyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "currencyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Exchange *ExchangeSession) CurrencyManager() (common.Address, error) {
	return _Exchange.Contract.CurrencyManager(&_Exchange.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_Exchange *ExchangeCallerSession) CurrencyManager() (common.Address, error) {
	return _Exchange.Contract.CurrencyManager(&_Exchange.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Exchange *ExchangeCaller) ExecutionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "executionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Exchange *ExchangeSession) ExecutionManager() (common.Address, error) {
	return _Exchange.Contract.ExecutionManager(&_Exchange.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_Exchange *ExchangeCallerSession) ExecutionManager() (common.Address, error) {
	return _Exchange.Contract.ExecutionManager(&_Exchange.CallOpts)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Exchange *ExchangeCaller) IsUserOrderNonceExecutedOrCancelled(opts *bind.CallOpts, user common.Address, orderNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isUserOrderNonceExecutedOrCancelled", user, orderNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Exchange *ExchangeSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _Exchange.Contract.IsUserOrderNonceExecutedOrCancelled(&_Exchange.CallOpts, user, orderNonce)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_Exchange *ExchangeCallerSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _Exchange.Contract.IsUserOrderNonceExecutedOrCancelled(&_Exchange.CallOpts, user, orderNonce)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Exchange.Contract.ProtocolFeeRecipient(&_Exchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_Exchange *ExchangeCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _Exchange.Contract.ProtocolFeeRecipient(&_Exchange.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Exchange *ExchangeCaller) RoyaltyFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "royaltyFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Exchange *ExchangeSession) RoyaltyFeeManager() (common.Address, error) {
	return _Exchange.Contract.RoyaltyFeeManager(&_Exchange.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_Exchange *ExchangeCallerSession) RoyaltyFeeManager() (common.Address, error) {
	return _Exchange.Contract.RoyaltyFeeManager(&_Exchange.CallOpts)
}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Exchange *ExchangeCaller) TransferSelectorNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "transferSelectorNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Exchange *ExchangeSession) TransferSelectorNFT() (common.Address, error) {
	return _Exchange.Contract.TransferSelectorNFT(&_Exchange.CallOpts)
}

// TransferSelectorNFT is a free data retrieval call binding the contract method 0x5e14f68e.
//
// Solidity: function transferSelectorNFT() view returns(address)
func (_Exchange *ExchangeCallerSession) TransferSelectorNFT() (common.Address, error) {
	return _Exchange.Contract.TransferSelectorNFT(&_Exchange.CallOpts)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Exchange *ExchangeCaller) UserMinOrderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "userMinOrderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Exchange *ExchangeSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.UserMinOrderNonce(&_Exchange.CallOpts, arg0)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_Exchange *ExchangeCallerSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _Exchange.Contract.UserMinOrderNonce(&_Exchange.CallOpts, arg0)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Exchange *ExchangeTransactor) CancelAllOrdersForSender(opts *bind.TransactOpts, minNonce *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelAllOrdersForSender", minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Exchange *ExchangeSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelAllOrdersForSender(&_Exchange.TransactOpts, minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_Exchange *ExchangeTransactorSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelAllOrdersForSender(&_Exchange.TransactOpts, minNonce)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Exchange *ExchangeTransactor) CancelMultipleMakerOrders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelMultipleMakerOrders", orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Exchange *ExchangeSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelMultipleMakerOrders(&_Exchange.TransactOpts, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_Exchange *ExchangeTransactorSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelMultipleMakerOrders(&_Exchange.TransactOpts, orderNonces)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Exchange *ExchangeTransactor) MatchAskWithTakerBid(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "matchAskWithTakerBid", takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Exchange *ExchangeSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchAskWithTakerBid(&_Exchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x38e29209.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_Exchange *ExchangeTransactorSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchAskWithTakerBid(&_Exchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Exchange *ExchangeTransactor) MatchAskWithTakerBidUsingETHAndWETH(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "matchAskWithTakerBidUsingETHAndWETH", takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Exchange *ExchangeSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_Exchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0xb4e4b296.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_Exchange *ExchangeTransactorSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_Exchange.TransactOpts, takerBid, makerAsk)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Exchange *ExchangeTransactor) MatchBidWithTakerAsk(opts *bind.TransactOpts, takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "matchBidWithTakerAsk", takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Exchange *ExchangeSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchBidWithTakerAsk(&_Exchange.TransactOpts, takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0x3b6d032e.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_Exchange *ExchangeTransactorSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _Exchange.Contract.MatchBidWithTakerAsk(&_Exchange.TransactOpts, takerAsk, makerBid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Exchange *ExchangeTransactor) UpdateCurrencyManager(opts *bind.TransactOpts, _currencyManager common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateCurrencyManager", _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Exchange *ExchangeSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateCurrencyManager(&_Exchange.TransactOpts, _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_Exchange *ExchangeTransactorSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateCurrencyManager(&_Exchange.TransactOpts, _currencyManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Exchange *ExchangeTransactor) UpdateExecutionManager(opts *bind.TransactOpts, _executionManager common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateExecutionManager", _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Exchange *ExchangeSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateExecutionManager(&_Exchange.TransactOpts, _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_Exchange *ExchangeTransactorSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateExecutionManager(&_Exchange.TransactOpts, _executionManager)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Exchange *ExchangeTransactor) UpdateProtocolFeeRecipient(opts *bind.TransactOpts, _protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateProtocolFeeRecipient", _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Exchange *ExchangeSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateProtocolFeeRecipient(&_Exchange.TransactOpts, _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_Exchange *ExchangeTransactorSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateProtocolFeeRecipient(&_Exchange.TransactOpts, _protocolFeeRecipient)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Exchange *ExchangeTransactor) UpdateRoyaltyFeeManager(opts *bind.TransactOpts, _royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateRoyaltyFeeManager", _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Exchange *ExchangeSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateRoyaltyFeeManager(&_Exchange.TransactOpts, _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_Exchange *ExchangeTransactorSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateRoyaltyFeeManager(&_Exchange.TransactOpts, _royaltyFeeManager)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Exchange *ExchangeTransactor) UpdateTransferSelectorNFT(opts *bind.TransactOpts, _transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "updateTransferSelectorNFT", _transferSelectorNFT)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Exchange *ExchangeSession) UpdateTransferSelectorNFT(_transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateTransferSelectorNFT(&_Exchange.TransactOpts, _transferSelectorNFT)
}

// UpdateTransferSelectorNFT is a paid mutator transaction binding the contract method 0xf75ff53f.
//
// Solidity: function updateTransferSelectorNFT(address _transferSelectorNFT) returns()
func (_Exchange *ExchangeTransactorSession) UpdateTransferSelectorNFT(_transferSelectorNFT common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.UpdateTransferSelectorNFT(&_Exchange.TransactOpts, _transferSelectorNFT)
}

// ExchangeCancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the Exchange contract.
type ExchangeCancelAllOrdersIterator struct {
	Event *ExchangeCancelAllOrders // Event containing the contract specifics and raw log

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
func (it *ExchangeCancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeCancelAllOrders)
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
		it.Event = new(ExchangeCancelAllOrders)
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
func (it *ExchangeCancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeCancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeCancelAllOrders represents a CancelAllOrders event raised by the Exchange contract.
type ExchangeCancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Exchange *ExchangeFilterer) FilterCancelAllOrders(opts *bind.FilterOpts, user []common.Address) (*ExchangeCancelAllOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeCancelAllOrdersIterator{contract: _Exchange.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Exchange *ExchangeFilterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *ExchangeCancelAllOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeCancelAllOrders)
				if err := _Exchange.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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

// ParseCancelAllOrders is a log parse operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_Exchange *ExchangeFilterer) ParseCancelAllOrders(log types.Log) (*ExchangeCancelAllOrders, error) {
	event := new(ExchangeCancelAllOrders)
	if err := _Exchange.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeCancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the Exchange contract.
type ExchangeCancelMultipleOrdersIterator struct {
	Event *ExchangeCancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *ExchangeCancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeCancelMultipleOrders)
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
		it.Event = new(ExchangeCancelMultipleOrders)
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
func (it *ExchangeCancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeCancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeCancelMultipleOrders represents a CancelMultipleOrders event raised by the Exchange contract.
type ExchangeCancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Exchange *ExchangeFilterer) FilterCancelMultipleOrders(opts *bind.FilterOpts, user []common.Address) (*ExchangeCancelMultipleOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeCancelMultipleOrdersIterator{contract: _Exchange.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Exchange *ExchangeFilterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *ExchangeCancelMultipleOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeCancelMultipleOrders)
				if err := _Exchange.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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

// ParseCancelMultipleOrders is a log parse operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_Exchange *ExchangeFilterer) ParseCancelMultipleOrders(log types.Log) (*ExchangeCancelMultipleOrders, error) {
	event := new(ExchangeCancelMultipleOrders)
	if err := _Exchange.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewCurrencyManagerIterator is returned from FilterNewCurrencyManager and is used to iterate over the raw logs and unpacked data for NewCurrencyManager events raised by the Exchange contract.
type ExchangeNewCurrencyManagerIterator struct {
	Event *ExchangeNewCurrencyManager // Event containing the contract specifics and raw log

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
func (it *ExchangeNewCurrencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewCurrencyManager)
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
		it.Event = new(ExchangeNewCurrencyManager)
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
func (it *ExchangeNewCurrencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewCurrencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewCurrencyManager represents a NewCurrencyManager event raised by the Exchange contract.
type ExchangeNewCurrencyManager struct {
	CurrencyManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCurrencyManager is a free log retrieval operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Exchange *ExchangeFilterer) FilterNewCurrencyManager(opts *bind.FilterOpts, currencyManager []common.Address) (*ExchangeNewCurrencyManagerIterator, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewCurrencyManagerIterator{contract: _Exchange.contract, event: "NewCurrencyManager", logs: logs, sub: sub}, nil
}

// WatchNewCurrencyManager is a free log subscription operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Exchange *ExchangeFilterer) WatchNewCurrencyManager(opts *bind.WatchOpts, sink chan<- *ExchangeNewCurrencyManager, currencyManager []common.Address) (event.Subscription, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewCurrencyManager)
				if err := _Exchange.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
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

// ParseNewCurrencyManager is a log parse operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_Exchange *ExchangeFilterer) ParseNewCurrencyManager(log types.Log) (*ExchangeNewCurrencyManager, error) {
	event := new(ExchangeNewCurrencyManager)
	if err := _Exchange.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewExecutionManagerIterator is returned from FilterNewExecutionManager and is used to iterate over the raw logs and unpacked data for NewExecutionManager events raised by the Exchange contract.
type ExchangeNewExecutionManagerIterator struct {
	Event *ExchangeNewExecutionManager // Event containing the contract specifics and raw log

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
func (it *ExchangeNewExecutionManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewExecutionManager)
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
		it.Event = new(ExchangeNewExecutionManager)
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
func (it *ExchangeNewExecutionManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewExecutionManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewExecutionManager represents a NewExecutionManager event raised by the Exchange contract.
type ExchangeNewExecutionManager struct {
	ExecutionManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionManager is a free log retrieval operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Exchange *ExchangeFilterer) FilterNewExecutionManager(opts *bind.FilterOpts, executionManager []common.Address) (*ExchangeNewExecutionManagerIterator, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewExecutionManagerIterator{contract: _Exchange.contract, event: "NewExecutionManager", logs: logs, sub: sub}, nil
}

// WatchNewExecutionManager is a free log subscription operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Exchange *ExchangeFilterer) WatchNewExecutionManager(opts *bind.WatchOpts, sink chan<- *ExchangeNewExecutionManager, executionManager []common.Address) (event.Subscription, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewExecutionManager)
				if err := _Exchange.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
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

// ParseNewExecutionManager is a log parse operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_Exchange *ExchangeFilterer) ParseNewExecutionManager(log types.Log) (*ExchangeNewExecutionManager, error) {
	event := new(ExchangeNewExecutionManager)
	if err := _Exchange.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the Exchange contract.
type ExchangeNewProtocolFeeRecipientIterator struct {
	Event *ExchangeNewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *ExchangeNewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewProtocolFeeRecipient)
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
		it.Event = new(ExchangeNewProtocolFeeRecipient)
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
func (it *ExchangeNewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the Exchange contract.
type ExchangeNewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Exchange *ExchangeFilterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts, protocolFeeRecipient []common.Address) (*ExchangeNewProtocolFeeRecipientIterator, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewProtocolFeeRecipientIterator{contract: _Exchange.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Exchange *ExchangeFilterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *ExchangeNewProtocolFeeRecipient, protocolFeeRecipient []common.Address) (event.Subscription, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewProtocolFeeRecipient)
				if err := _Exchange.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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

// ParseNewProtocolFeeRecipient is a log parse operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_Exchange *ExchangeFilterer) ParseNewProtocolFeeRecipient(log types.Log) (*ExchangeNewProtocolFeeRecipient, error) {
	event := new(ExchangeNewProtocolFeeRecipient)
	if err := _Exchange.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewRoyaltyFeeManagerIterator is returned from FilterNewRoyaltyFeeManager and is used to iterate over the raw logs and unpacked data for NewRoyaltyFeeManager events raised by the Exchange contract.
type ExchangeNewRoyaltyFeeManagerIterator struct {
	Event *ExchangeNewRoyaltyFeeManager // Event containing the contract specifics and raw log

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
func (it *ExchangeNewRoyaltyFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewRoyaltyFeeManager)
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
		it.Event = new(ExchangeNewRoyaltyFeeManager)
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
func (it *ExchangeNewRoyaltyFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewRoyaltyFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewRoyaltyFeeManager represents a NewRoyaltyFeeManager event raised by the Exchange contract.
type ExchangeNewRoyaltyFeeManager struct {
	RoyaltyFeeManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRoyaltyFeeManager is a free log retrieval operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Exchange *ExchangeFilterer) FilterNewRoyaltyFeeManager(opts *bind.FilterOpts, royaltyFeeManager []common.Address) (*ExchangeNewRoyaltyFeeManagerIterator, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewRoyaltyFeeManagerIterator{contract: _Exchange.contract, event: "NewRoyaltyFeeManager", logs: logs, sub: sub}, nil
}

// WatchNewRoyaltyFeeManager is a free log subscription operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Exchange *ExchangeFilterer) WatchNewRoyaltyFeeManager(opts *bind.WatchOpts, sink chan<- *ExchangeNewRoyaltyFeeManager, royaltyFeeManager []common.Address) (event.Subscription, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewRoyaltyFeeManager)
				if err := _Exchange.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
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

// ParseNewRoyaltyFeeManager is a log parse operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_Exchange *ExchangeFilterer) ParseNewRoyaltyFeeManager(log types.Log) (*ExchangeNewRoyaltyFeeManager, error) {
	event := new(ExchangeNewRoyaltyFeeManager)
	if err := _Exchange.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeNewTransferSelectorNFTIterator is returned from FilterNewTransferSelectorNFT and is used to iterate over the raw logs and unpacked data for NewTransferSelectorNFT events raised by the Exchange contract.
type ExchangeNewTransferSelectorNFTIterator struct {
	Event *ExchangeNewTransferSelectorNFT // Event containing the contract specifics and raw log

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
func (it *ExchangeNewTransferSelectorNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeNewTransferSelectorNFT)
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
		it.Event = new(ExchangeNewTransferSelectorNFT)
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
func (it *ExchangeNewTransferSelectorNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeNewTransferSelectorNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeNewTransferSelectorNFT represents a NewTransferSelectorNFT event raised by the Exchange contract.
type ExchangeNewTransferSelectorNFT struct {
	TransferSelectorNFT common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewTransferSelectorNFT is a free log retrieval operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Exchange *ExchangeFilterer) FilterNewTransferSelectorNFT(opts *bind.FilterOpts, transferSelectorNFT []common.Address) (*ExchangeNewTransferSelectorNFTIterator, error) {

	var transferSelectorNFTRule []interface{}
	for _, transferSelectorNFTItem := range transferSelectorNFT {
		transferSelectorNFTRule = append(transferSelectorNFTRule, transferSelectorNFTItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "NewTransferSelectorNFT", transferSelectorNFTRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeNewTransferSelectorNFTIterator{contract: _Exchange.contract, event: "NewTransferSelectorNFT", logs: logs, sub: sub}, nil
}

// WatchNewTransferSelectorNFT is a free log subscription operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Exchange *ExchangeFilterer) WatchNewTransferSelectorNFT(opts *bind.WatchOpts, sink chan<- *ExchangeNewTransferSelectorNFT, transferSelectorNFT []common.Address) (event.Subscription, error) {

	var transferSelectorNFTRule []interface{}
	for _, transferSelectorNFTItem := range transferSelectorNFT {
		transferSelectorNFTRule = append(transferSelectorNFTRule, transferSelectorNFTItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "NewTransferSelectorNFT", transferSelectorNFTRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeNewTransferSelectorNFT)
				if err := _Exchange.contract.UnpackLog(event, "NewTransferSelectorNFT", log); err != nil {
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

// ParseNewTransferSelectorNFT is a log parse operation binding the contract event 0x205d78ab41afe80bd6b6aaa5d7599d5300ff8690da3ab1302c1b552f7baf7d8c.
//
// Solidity: event NewTransferSelectorNFT(address indexed transferSelectorNFT)
func (_Exchange *ExchangeFilterer) ParseNewTransferSelectorNFT(log types.Log) (*ExchangeNewTransferSelectorNFT, error) {
	event := new(ExchangeNewTransferSelectorNFT)
	if err := _Exchange.contract.UnpackLog(event, "NewTransferSelectorNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeRoyaltyPaymentIterator is returned from FilterRoyaltyPayment and is used to iterate over the raw logs and unpacked data for RoyaltyPayment events raised by the Exchange contract.
type ExchangeRoyaltyPaymentIterator struct {
	Event *ExchangeRoyaltyPayment // Event containing the contract specifics and raw log

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
func (it *ExchangeRoyaltyPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRoyaltyPayment)
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
		it.Event = new(ExchangeRoyaltyPayment)
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
func (it *ExchangeRoyaltyPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRoyaltyPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRoyaltyPayment represents a RoyaltyPayment event raised by the Exchange contract.
type ExchangeRoyaltyPayment struct {
	Collection       common.Address
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	Currency         common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyPayment is a free log retrieval operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Exchange *ExchangeFilterer) FilterRoyaltyPayment(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (*ExchangeRoyaltyPaymentIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRoyaltyPaymentIterator{contract: _Exchange.contract, event: "RoyaltyPayment", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPayment is a free log subscription operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Exchange *ExchangeFilterer) WatchRoyaltyPayment(opts *bind.WatchOpts, sink chan<- *ExchangeRoyaltyPayment, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRoyaltyPayment)
				if err := _Exchange.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
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

// ParseRoyaltyPayment is a log parse operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_Exchange *ExchangeFilterer) ParseRoyaltyPayment(log types.Log) (*ExchangeRoyaltyPayment, error) {
	event := new(ExchangeRoyaltyPayment)
	if err := _Exchange.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeTakerAskIterator is returned from FilterTakerAsk and is used to iterate over the raw logs and unpacked data for TakerAsk events raised by the Exchange contract.
type ExchangeTakerAskIterator struct {
	Event *ExchangeTakerAsk // Event containing the contract specifics and raw log

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
func (it *ExchangeTakerAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTakerAsk)
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
		it.Event = new(ExchangeTakerAsk)
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
func (it *ExchangeTakerAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTakerAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTakerAsk represents a TakerAsk event raised by the Exchange contract.
type ExchangeTakerAsk struct {
	OrderHash  [32]byte
	OrderNonce *big.Int
	Taker      common.Address
	Maker      common.Address
	Strategy   common.Address
	Currency   common.Address
	Collection common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTakerAsk is a free log retrieval operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) FilterTakerAsk(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*ExchangeTakerAskIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTakerAskIterator{contract: _Exchange.contract, event: "TakerAsk", logs: logs, sub: sub}, nil
}

// WatchTakerAsk is a free log subscription operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) WatchTakerAsk(opts *bind.WatchOpts, sink chan<- *ExchangeTakerAsk, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTakerAsk)
				if err := _Exchange.contract.UnpackLog(event, "TakerAsk", log); err != nil {
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

// ParseTakerAsk is a log parse operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) ParseTakerAsk(log types.Log) (*ExchangeTakerAsk, error) {
	event := new(ExchangeTakerAsk)
	if err := _Exchange.contract.UnpackLog(event, "TakerAsk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeTakerBidIterator is returned from FilterTakerBid and is used to iterate over the raw logs and unpacked data for TakerBid events raised by the Exchange contract.
type ExchangeTakerBidIterator struct {
	Event *ExchangeTakerBid // Event containing the contract specifics and raw log

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
func (it *ExchangeTakerBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeTakerBid)
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
		it.Event = new(ExchangeTakerBid)
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
func (it *ExchangeTakerBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeTakerBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeTakerBid represents a TakerBid event raised by the Exchange contract.
type ExchangeTakerBid struct {
	OrderHash  [32]byte
	OrderNonce *big.Int
	Taker      common.Address
	Maker      common.Address
	Strategy   common.Address
	Currency   common.Address
	Collection common.Address
	TokenId    *big.Int
	Amount     *big.Int
	Price      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTakerBid is a free log retrieval operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) FilterTakerBid(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*ExchangeTakerBidIterator, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeTakerBidIterator{contract: _Exchange.contract, event: "TakerBid", logs: logs, sub: sub}, nil
}

// WatchTakerBid is a free log subscription operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) WatchTakerBid(opts *bind.WatchOpts, sink chan<- *ExchangeTakerBid, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var strategyRule []interface{}
	for _, strategyItem := range strategy {
		strategyRule = append(strategyRule, strategyItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeTakerBid)
				if err := _Exchange.contract.UnpackLog(event, "TakerBid", log); err != nil {
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

// ParseTakerBid is a log parse operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 orderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_Exchange *ExchangeFilterer) ParseTakerBid(log types.Log) (*ExchangeTakerBid, error) {
	event := new(ExchangeTakerBid)
	if err := _Exchange.contract.UnpackLog(event, "TakerBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
