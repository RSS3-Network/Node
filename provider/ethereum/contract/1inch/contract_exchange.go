// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package oneinch

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

// IZrxExchangeFillResults is an auto generated low-level Go binding around an user-defined struct.
type IZrxExchangeFillResults struct {
	MakerAssetFilledAmount *big.Int
	TakerAssetFilledAmount *big.Int
	MakerFeePaid           *big.Int
	TakerFeePaid           *big.Int
}

// IZrxExchangeOrder is an auto generated low-level Go binding around an user-defined struct.
type IZrxExchangeOrder struct {
	MakerAddress          common.Address
	TakerAddress          common.Address
	FeeRecipientAddress   common.Address
	SenderAddress         common.Address
	MakerAssetAmount      *big.Int
	TakerAssetAmount      *big.Int
	MakerFee              *big.Int
	TakerFee              *big.Int
	ExpirationTimeSeconds *big.Int
	Salt                  *big.Int
	MakerAssetData        []byte
	TakerAssetData        []byte
}

// IZrxExchangeOrderInfo is an auto generated low-level Go binding around an user-defined struct.
type IZrxExchangeOrderInfo struct {
	OrderStatus                 uint8
	OrderHash                   [32]byte
	OrderTakerAssetFilledAmount *big.Int
}

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIZrxExchange\",\"name\":\"zrx\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zrxTokenProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIZrxExchange.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"getOrdersInfoRespectingBalancesAndAllowances\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"orderStatus\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"orderTakerAssetFilledAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIZrxExchange.OrderInfo[]\",\"name\":\"ordersInfo\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenSell\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenBuy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zrxExchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zrxTokenProxy\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIZrxExchange.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"mul\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"div\",\"type\":\"uint256\"}],\"name\":\"marketSellOrdersProportion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdrawAllToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isShutdown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"infiniteApproveIfNeeded\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spender\",\"outputs\":[{\"internalType\":\"contractTokenSpender\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"callAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"callDataConcat\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"starts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasLimitsAndValues\",\"type\":\"uint256[]\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"shutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zrxExchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"zrxTokenProxy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFillAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipientAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimeSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"makerAssetData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"takerAssetData\",\"type\":\"bytes\"}],\"internalType\":\"structIZrxExchange.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"marketSellOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"makerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAssetFilledAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerFeePaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerFeePaid\",\"type\":\"uint256\"}],\"internalType\":\"structIZrxExchange.FillResults\",\"name\":\"totalFillResults\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIGST2\",\"name\":\"_gasToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"}],\"name\":\"History\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referrerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Swapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Shutdown\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]",
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

// GetOrdersInfoRespectingBalancesAndAllowances is a free data retrieval call binding the contract method 0x71a2039a.
//
// Solidity: function getOrdersInfoRespectingBalancesAndAllowances(address token, address zrx, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders) view returns((uint8,bytes32,uint256)[] ordersInfo)
func (_Exchange *ExchangeCaller) GetOrdersInfoRespectingBalancesAndAllowances(opts *bind.CallOpts, token common.Address, zrx common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder) ([]IZrxExchangeOrderInfo, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getOrdersInfoRespectingBalancesAndAllowances", token, zrx, zrxTokenProxy, orders)

	if err != nil {
		return *new([]IZrxExchangeOrderInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IZrxExchangeOrderInfo)).(*[]IZrxExchangeOrderInfo)

	return out0, err

}

// GetOrdersInfoRespectingBalancesAndAllowances is a free data retrieval call binding the contract method 0x71a2039a.
//
// Solidity: function getOrdersInfoRespectingBalancesAndAllowances(address token, address zrx, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders) view returns((uint8,bytes32,uint256)[] ordersInfo)
func (_Exchange *ExchangeSession) GetOrdersInfoRespectingBalancesAndAllowances(token common.Address, zrx common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder) ([]IZrxExchangeOrderInfo, error) {
	return _Exchange.Contract.GetOrdersInfoRespectingBalancesAndAllowances(&_Exchange.CallOpts, token, zrx, zrxTokenProxy, orders)
}

// GetOrdersInfoRespectingBalancesAndAllowances is a free data retrieval call binding the contract method 0x71a2039a.
//
// Solidity: function getOrdersInfoRespectingBalancesAndAllowances(address token, address zrx, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders) view returns((uint8,bytes32,uint256)[] ordersInfo)
func (_Exchange *ExchangeCallerSession) GetOrdersInfoRespectingBalancesAndAllowances(token common.Address, zrx common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder) ([]IZrxExchangeOrderInfo, error) {
	return _Exchange.Contract.GetOrdersInfoRespectingBalancesAndAllowances(&_Exchange.CallOpts, token, zrx, zrxTokenProxy, orders)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeCallerSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Exchange *ExchangeCaller) IsShutdown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isShutdown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Exchange *ExchangeSession) IsShutdown() (bool, error) {
	return _Exchange.Contract.IsShutdown(&_Exchange.CallOpts)
}

// IsShutdown is a free data retrieval call binding the contract method 0xbf86d690.
//
// Solidity: function isShutdown() view returns(bool)
func (_Exchange *ExchangeCallerSession) IsShutdown() (bool, error) {
	return _Exchange.Contract.IsShutdown(&_Exchange.CallOpts)
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

// Spender is a free data retrieval call binding the contract method 0xe8edc816.
//
// Solidity: function spender() view returns(address)
func (_Exchange *ExchangeCaller) Spender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "spender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Spender is a free data retrieval call binding the contract method 0xe8edc816.
//
// Solidity: function spender() view returns(address)
func (_Exchange *ExchangeSession) Spender() (common.Address, error) {
	return _Exchange.Contract.Spender(&_Exchange.CallOpts)
}

// Spender is a free data retrieval call binding the contract method 0xe8edc816.
//
// Solidity: function spender() view returns(address)
func (_Exchange *ExchangeCallerSession) Spender() (common.Address, error) {
	return _Exchange.Contract.Spender(&_Exchange.CallOpts)
}

// InfiniteApproveIfNeeded is a paid mutator transaction binding the contract method 0xc9b27359.
//
// Solidity: function infiniteApproveIfNeeded(address token, address to) returns()
func (_Exchange *ExchangeTransactor) InfiniteApproveIfNeeded(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "infiniteApproveIfNeeded", token, to)
}

// InfiniteApproveIfNeeded is a paid mutator transaction binding the contract method 0xc9b27359.
//
// Solidity: function infiniteApproveIfNeeded(address token, address to) returns()
func (_Exchange *ExchangeSession) InfiniteApproveIfNeeded(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.InfiniteApproveIfNeeded(&_Exchange.TransactOpts, token, to)
}

// InfiniteApproveIfNeeded is a paid mutator transaction binding the contract method 0xc9b27359.
//
// Solidity: function infiniteApproveIfNeeded(address token, address to) returns()
func (_Exchange *ExchangeTransactorSession) InfiniteApproveIfNeeded(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.InfiniteApproveIfNeeded(&_Exchange.TransactOpts, token, to)
}

// MarketSellOrders is a paid mutator transaction binding the contract method 0xfcc06f8e.
//
// Solidity: function marketSellOrders(address makerAsset, address zrxExchange, address zrxTokenProxy, uint256 takerAssetFillAmount, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures) returns((uint256,uint256,uint256,uint256) totalFillResults)
func (_Exchange *ExchangeTransactor) MarketSellOrders(opts *bind.TransactOpts, makerAsset common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, takerAssetFillAmount *big.Int, orders []IZrxExchangeOrder, signatures [][]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "marketSellOrders", makerAsset, zrxExchange, zrxTokenProxy, takerAssetFillAmount, orders, signatures)
}

// MarketSellOrders is a paid mutator transaction binding the contract method 0xfcc06f8e.
//
// Solidity: function marketSellOrders(address makerAsset, address zrxExchange, address zrxTokenProxy, uint256 takerAssetFillAmount, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures) returns((uint256,uint256,uint256,uint256) totalFillResults)
func (_Exchange *ExchangeSession) MarketSellOrders(makerAsset common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, takerAssetFillAmount *big.Int, orders []IZrxExchangeOrder, signatures [][]byte) (*types.Transaction, error) {
	return _Exchange.Contract.MarketSellOrders(&_Exchange.TransactOpts, makerAsset, zrxExchange, zrxTokenProxy, takerAssetFillAmount, orders, signatures)
}

// MarketSellOrders is a paid mutator transaction binding the contract method 0xfcc06f8e.
//
// Solidity: function marketSellOrders(address makerAsset, address zrxExchange, address zrxTokenProxy, uint256 takerAssetFillAmount, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures) returns((uint256,uint256,uint256,uint256) totalFillResults)
func (_Exchange *ExchangeTransactorSession) MarketSellOrders(makerAsset common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, takerAssetFillAmount *big.Int, orders []IZrxExchangeOrder, signatures [][]byte) (*types.Transaction, error) {
	return _Exchange.Contract.MarketSellOrders(&_Exchange.TransactOpts, makerAsset, zrxExchange, zrxTokenProxy, takerAssetFillAmount, orders, signatures)
}

// MarketSellOrdersProportion is a paid mutator transaction binding the contract method 0xa96c400e.
//
// Solidity: function marketSellOrdersProportion(address tokenSell, address tokenBuy, address zrxExchange, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures, uint256 mul, uint256 div) returns()
func (_Exchange *ExchangeTransactor) MarketSellOrdersProportion(opts *bind.TransactOpts, tokenSell common.Address, tokenBuy common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder, signatures [][]byte, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "marketSellOrdersProportion", tokenSell, tokenBuy, zrxExchange, zrxTokenProxy, orders, signatures, mul, div)
}

// MarketSellOrdersProportion is a paid mutator transaction binding the contract method 0xa96c400e.
//
// Solidity: function marketSellOrdersProportion(address tokenSell, address tokenBuy, address zrxExchange, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures, uint256 mul, uint256 div) returns()
func (_Exchange *ExchangeSession) MarketSellOrdersProportion(tokenSell common.Address, tokenBuy common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder, signatures [][]byte, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.MarketSellOrdersProportion(&_Exchange.TransactOpts, tokenSell, tokenBuy, zrxExchange, zrxTokenProxy, orders, signatures, mul, div)
}

// MarketSellOrdersProportion is a paid mutator transaction binding the contract method 0xa96c400e.
//
// Solidity: function marketSellOrdersProportion(address tokenSell, address tokenBuy, address zrxExchange, address zrxTokenProxy, (address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes,bytes)[] orders, bytes[] signatures, uint256 mul, uint256 div) returns()
func (_Exchange *ExchangeTransactorSession) MarketSellOrdersProportion(tokenSell common.Address, tokenBuy common.Address, zrxExchange common.Address, zrxTokenProxy common.Address, orders []IZrxExchangeOrder, signatures [][]byte, mul *big.Int, div *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.MarketSellOrdersProportion(&_Exchange.TransactOpts, tokenSell, tokenBuy, zrxExchange, zrxTokenProxy, orders, signatures, mul, div)
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

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Exchange *ExchangeTransactor) Shutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "shutdown")
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Exchange *ExchangeSession) Shutdown() (*types.Transaction, error) {
	return _Exchange.Contract.Shutdown(&_Exchange.TransactOpts)
}

// Shutdown is a paid mutator transaction binding the contract method 0xfc0e74d1.
//
// Solidity: function shutdown() returns()
func (_Exchange *ExchangeTransactorSession) Shutdown() (*types.Transaction, error) {
	return _Exchange.Contract.Shutdown(&_Exchange.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Exchange *ExchangeTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "swap", fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Exchange *ExchangeSession) Swap(fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Swap(&_Exchange.TransactOpts, fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Exchange *ExchangeTransactorSession) Swap(fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Swap(&_Exchange.TransactOpts, fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
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

// WithdrawAllToken is a paid mutator transaction binding the contract method 0xae4dd0fc.
//
// Solidity: function withdrawAllToken(address token) returns()
func (_Exchange *ExchangeTransactor) WithdrawAllToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "withdrawAllToken", token)
}

// WithdrawAllToken is a paid mutator transaction binding the contract method 0xae4dd0fc.
//
// Solidity: function withdrawAllToken(address token) returns()
func (_Exchange *ExchangeSession) WithdrawAllToken(token common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawAllToken(&_Exchange.TransactOpts, token)
}

// WithdrawAllToken is a paid mutator transaction binding the contract method 0xae4dd0fc.
//
// Solidity: function withdrawAllToken(address token) returns()
func (_Exchange *ExchangeTransactorSession) WithdrawAllToken(token common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.WithdrawAllToken(&_Exchange.TransactOpts, token)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Exchange.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Fallback(&_Exchange.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Exchange *ExchangeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Exchange.Contract.Fallback(&_Exchange.TransactOpts, calldata)
}

// ExchangeHistoryIterator is returned from FilterHistory and is used to iterate over the raw logs and unpacked data for History events raised by the Exchange contract.
type ExchangeHistoryIterator struct {
	Event *ExchangeHistory // Event containing the contract specifics and raw log

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
func (it *ExchangeHistoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeHistory)
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
		it.Event = new(ExchangeHistory)
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
func (it *ExchangeHistoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeHistoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeHistory represents a History event raised by the Exchange contract.
type ExchangeHistory struct {
	Sender     common.Address
	FromToken  common.Address
	ToToken    common.Address
	FromAmount *big.Int
	ToAmount   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHistory is a free log retrieval operation binding the contract event 0x894dbf1262199c24e1750298a384c709160f49d163422cc6cee694c73713f1d2.
//
// Solidity: event History(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount)
func (_Exchange *ExchangeFilterer) FilterHistory(opts *bind.FilterOpts, sender []common.Address) (*ExchangeHistoryIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "History", senderRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeHistoryIterator{contract: _Exchange.contract, event: "History", logs: logs, sub: sub}, nil
}

// WatchHistory is a free log subscription operation binding the contract event 0x894dbf1262199c24e1750298a384c709160f49d163422cc6cee694c73713f1d2.
//
// Solidity: event History(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount)
func (_Exchange *ExchangeFilterer) WatchHistory(opts *bind.WatchOpts, sink chan<- *ExchangeHistory, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "History", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeHistory)
				if err := _Exchange.contract.UnpackLog(event, "History", log); err != nil {
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

// ParseHistory is a log parse operation binding the contract event 0x894dbf1262199c24e1750298a384c709160f49d163422cc6cee694c73713f1d2.
//
// Solidity: event History(address indexed sender, address fromToken, address toToken, uint256 fromAmount, uint256 toAmount)
func (_Exchange *ExchangeFilterer) ParseHistory(log types.Log) (*ExchangeHistory, error) {
	event := new(ExchangeHistory)
	if err := _Exchange.contract.UnpackLog(event, "History", log); err != nil {
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

// ExchangeShutdownIterator is returned from FilterShutdown and is used to iterate over the raw logs and unpacked data for Shutdown events raised by the Exchange contract.
type ExchangeShutdownIterator struct {
	Event *ExchangeShutdown // Event containing the contract specifics and raw log

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
func (it *ExchangeShutdownIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeShutdown)
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
		it.Event = new(ExchangeShutdown)
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
func (it *ExchangeShutdownIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeShutdownIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeShutdown represents a Shutdown event raised by the Exchange contract.
type ExchangeShutdown struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterShutdown is a free log retrieval operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Exchange *ExchangeFilterer) FilterShutdown(opts *bind.FilterOpts) (*ExchangeShutdownIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return &ExchangeShutdownIterator{contract: _Exchange.contract, event: "Shutdown", logs: logs, sub: sub}, nil
}

// WatchShutdown is a free log subscription operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Exchange *ExchangeFilterer) WatchShutdown(opts *bind.WatchOpts, sink chan<- *ExchangeShutdown) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Shutdown")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeShutdown)
				if err := _Exchange.contract.UnpackLog(event, "Shutdown", log); err != nil {
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

// ParseShutdown is a log parse operation binding the contract event 0x4426aa1fb73e391071491fcfe21a88b5c38a0a0333a1f6e77161470439704cf8.
//
// Solidity: event Shutdown()
func (_Exchange *ExchangeFilterer) ParseShutdown(log types.Log) (*ExchangeShutdown, error) {
	event := new(ExchangeShutdown)
	if err := _Exchange.contract.UnpackLog(event, "Shutdown", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeSwappedIterator is returned from FilterSwapped and is used to iterate over the raw logs and unpacked data for Swapped events raised by the Exchange contract.
type ExchangeSwappedIterator struct {
	Event *ExchangeSwapped // Event containing the contract specifics and raw log

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
func (it *ExchangeSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeSwapped)
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
		it.Event = new(ExchangeSwapped)
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
func (it *ExchangeSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeSwapped represents a Swapped event raised by the Exchange contract.
type ExchangeSwapped struct {
	FromToken   common.Address
	ToToken     common.Address
	Referrer    common.Address
	FromAmount  *big.Int
	ToAmount    *big.Int
	ReferrerFee *big.Int
	Fee         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapped is a free log retrieval operation binding the contract event 0xe2cee3f6836059820b673943853afebd9b3026125dab0d774284e6f28a4855be.
//
// Solidity: event Swapped(address indexed fromToken, address indexed toToken, address indexed referrer, uint256 fromAmount, uint256 toAmount, uint256 referrerFee, uint256 fee)
func (_Exchange *ExchangeFilterer) FilterSwapped(opts *bind.FilterOpts, fromToken []common.Address, toToken []common.Address, referrer []common.Address) (*ExchangeSwappedIterator, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Swapped", fromTokenRule, toTokenRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeSwappedIterator{contract: _Exchange.contract, event: "Swapped", logs: logs, sub: sub}, nil
}

// WatchSwapped is a free log subscription operation binding the contract event 0xe2cee3f6836059820b673943853afebd9b3026125dab0d774284e6f28a4855be.
//
// Solidity: event Swapped(address indexed fromToken, address indexed toToken, address indexed referrer, uint256 fromAmount, uint256 toAmount, uint256 referrerFee, uint256 fee)
func (_Exchange *ExchangeFilterer) WatchSwapped(opts *bind.WatchOpts, sink chan<- *ExchangeSwapped, fromToken []common.Address, toToken []common.Address, referrer []common.Address) (event.Subscription, error) {

	var fromTokenRule []interface{}
	for _, fromTokenItem := range fromToken {
		fromTokenRule = append(fromTokenRule, fromTokenItem)
	}
	var toTokenRule []interface{}
	for _, toTokenItem := range toToken {
		toTokenRule = append(toTokenRule, toTokenItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Swapped", fromTokenRule, toTokenRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeSwapped)
				if err := _Exchange.contract.UnpackLog(event, "Swapped", log); err != nil {
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

// ParseSwapped is a log parse operation binding the contract event 0xe2cee3f6836059820b673943853afebd9b3026125dab0d774284e6f28a4855be.
//
// Solidity: event Swapped(address indexed fromToken, address indexed toToken, address indexed referrer, uint256 fromAmount, uint256 toAmount, uint256 referrerFee, uint256 fee)
func (_Exchange *ExchangeFilterer) ParseSwapped(log types.Log) (*ExchangeSwapped, error) {
	event := new(ExchangeSwapped)
	if err := _Exchange.contract.UnpackLog(event, "Swapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
