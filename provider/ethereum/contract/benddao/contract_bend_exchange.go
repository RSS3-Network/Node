// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benddao

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
	Maker              common.Address
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
	Interceptor        common.Address
	InterceptorExtra   []byte
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
	Interceptor        common.Address
	InterceptorExtra   []byte
}

// BendExchangeMetaData contains all meta data concerning the BendExchange contract.
var BendExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interceptorManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_transferManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_WETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinNonce\",\"type\":\"uint256\"}],\"name\":\"CancelAllOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"CancelMultipleOrders\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"authorizationManager\",\"type\":\"address\"}],\"name\":\"NewAuthorizationManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currencyManager\",\"type\":\"address\"}],\"name\":\"NewCurrencyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executionManager\",\"type\":\"address\"}],\"name\":\"NewExecutionManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"interceptorManager\",\"type\":\"address\"}],\"name\":\"NewInterceptorManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"NewProtocolFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"NewRoyaltyFeeManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"transferManager\",\"type\":\"address\"}],\"name\":\"NewTransferManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProtocolFeePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RoyaltyPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"makerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerAsk\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"makerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TakerBid\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"authorizationManager\",\"outputs\":[{\"internalType\":\"contractIAuthorizationManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minNonce\",\"type\":\"uint256\"}],\"name\":\"cancelAllOrdersForSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"orderNonces\",\"type\":\"uint256[]\"}],\"name\":\"cancelMultipleMakerOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currencyManager\",\"outputs\":[{\"internalType\":\"contractICurrencyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionManager\",\"outputs\":[{\"internalType\":\"contractIExecutionManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interceptorManager\",\"outputs\":[{\"internalType\":\"contractIInterceptorManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"orderNonce\",\"type\":\"uint256\"}],\"name\":\"isUserOrderNonceExecutedOrCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerBid\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerAsk\",\"type\":\"tuple\"}],\"name\":\"matchAskWithTakerBidUsingETHAndWETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"}],\"internalType\":\"structOrderTypes.TakerOrder\",\"name\":\"takerAsk\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isOrderAsk\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPercentageToAsk\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"interceptor\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"interceptorExtra\",\"type\":\"bytes\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structOrderTypes.MakerOrder\",\"name\":\"makerBid\",\"type\":\"tuple\"}],\"name\":\"matchBidWithTakerAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyFeeManager\",\"outputs\":[{\"internalType\":\"contractIRoyaltyFeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferManager\",\"outputs\":[{\"internalType\":\"contractITransferManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationManager\",\"type\":\"address\"}],\"name\":\"updateAuthorizationManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_currencyManager\",\"type\":\"address\"}],\"name\":\"updateCurrencyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executionManager\",\"type\":\"address\"}],\"name\":\"updateExecutionManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interceptorManager\",\"type\":\"address\"}],\"name\":\"updateInterceptorManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_protocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"updateProtocolFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyFeeManager\",\"type\":\"address\"}],\"name\":\"updateRoyaltyFeeManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_transferManager\",\"type\":\"address\"}],\"name\":\"updateTransferManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMinOrderNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BendExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use BendExchangeMetaData.ABI instead.
var BendExchangeABI = BendExchangeMetaData.ABI

// BendExchange is an auto generated Go binding around an Ethereum contract.
type BendExchange struct {
	BendExchangeCaller     // Read-only binding to the contract
	BendExchangeTransactor // Write-only binding to the contract
	BendExchangeFilterer   // Log filterer for contract events
}

// BendExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BendExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BendExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BendExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BendExchangeSession struct {
	Contract     *BendExchange     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BendExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BendExchangeCallerSession struct {
	Contract *BendExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BendExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BendExchangeTransactorSession struct {
	Contract     *BendExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BendExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BendExchangeRaw struct {
	Contract *BendExchange // Generic contract binding to access the raw methods on
}

// BendExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BendExchangeCallerRaw struct {
	Contract *BendExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// BendExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BendExchangeTransactorRaw struct {
	Contract *BendExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBendExchange creates a new instance of BendExchange, bound to a specific deployed contract.
func NewBendExchange(address common.Address, backend bind.ContractBackend) (*BendExchange, error) {
	contract, err := bindBendExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BendExchange{BendExchangeCaller: BendExchangeCaller{contract: contract}, BendExchangeTransactor: BendExchangeTransactor{contract: contract}, BendExchangeFilterer: BendExchangeFilterer{contract: contract}}, nil
}

// NewBendExchangeCaller creates a new read-only instance of BendExchange, bound to a specific deployed contract.
func NewBendExchangeCaller(address common.Address, caller bind.ContractCaller) (*BendExchangeCaller, error) {
	contract, err := bindBendExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BendExchangeCaller{contract: contract}, nil
}

// NewBendExchangeTransactor creates a new write-only instance of BendExchange, bound to a specific deployed contract.
func NewBendExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*BendExchangeTransactor, error) {
	contract, err := bindBendExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BendExchangeTransactor{contract: contract}, nil
}

// NewBendExchangeFilterer creates a new log filterer instance of BendExchange, bound to a specific deployed contract.
func NewBendExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*BendExchangeFilterer, error) {
	contract, err := bindBendExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BendExchangeFilterer{contract: contract}, nil
}

// bindBendExchange binds a generic wrapper to an already deployed contract.
func bindBendExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BendExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendExchange *BendExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendExchange.Contract.BendExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendExchange *BendExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendExchange.Contract.BendExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendExchange *BendExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendExchange.Contract.BendExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendExchange *BendExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendExchange *BendExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendExchange *BendExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendExchange.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendExchange *BendExchangeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendExchange *BendExchangeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BendExchange.Contract.DOMAINSEPARATOR(&_BendExchange.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_BendExchange *BendExchangeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _BendExchange.Contract.DOMAINSEPARATOR(&_BendExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BendExchange *BendExchangeCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BendExchange *BendExchangeSession) NAME() (string, error) {
	return _BendExchange.Contract.NAME(&_BendExchange.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_BendExchange *BendExchangeCallerSession) NAME() (string, error) {
	return _BendExchange.Contract.NAME(&_BendExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BendExchange *BendExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BendExchange *BendExchangeSession) VERSION() (string, error) {
	return _BendExchange.Contract.VERSION(&_BendExchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_BendExchange *BendExchangeCallerSession) VERSION() (string, error) {
	return _BendExchange.Contract.VERSION(&_BendExchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BendExchange *BendExchangeCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BendExchange *BendExchangeSession) WETH() (common.Address, error) {
	return _BendExchange.Contract.WETH(&_BendExchange.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_BendExchange *BendExchangeCallerSession) WETH() (common.Address, error) {
	return _BendExchange.Contract.WETH(&_BendExchange.CallOpts)
}

// AuthorizationManager is a free data retrieval call binding the contract method 0x8cfa228c.
//
// Solidity: function authorizationManager() view returns(address)
func (_BendExchange *BendExchangeCaller) AuthorizationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "authorizationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuthorizationManager is a free data retrieval call binding the contract method 0x8cfa228c.
//
// Solidity: function authorizationManager() view returns(address)
func (_BendExchange *BendExchangeSession) AuthorizationManager() (common.Address, error) {
	return _BendExchange.Contract.AuthorizationManager(&_BendExchange.CallOpts)
}

// AuthorizationManager is a free data retrieval call binding the contract method 0x8cfa228c.
//
// Solidity: function authorizationManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) AuthorizationManager() (common.Address, error) {
	return _BendExchange.Contract.AuthorizationManager(&_BendExchange.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_BendExchange *BendExchangeCaller) CurrencyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "currencyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_BendExchange *BendExchangeSession) CurrencyManager() (common.Address, error) {
	return _BendExchange.Contract.CurrencyManager(&_BendExchange.CallOpts)
}

// CurrencyManager is a free data retrieval call binding the contract method 0x0f747d74.
//
// Solidity: function currencyManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) CurrencyManager() (common.Address, error) {
	return _BendExchange.Contract.CurrencyManager(&_BendExchange.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_BendExchange *BendExchangeCaller) ExecutionManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "executionManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_BendExchange *BendExchangeSession) ExecutionManager() (common.Address, error) {
	return _BendExchange.Contract.ExecutionManager(&_BendExchange.CallOpts)
}

// ExecutionManager is a free data retrieval call binding the contract method 0x483abb9f.
//
// Solidity: function executionManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) ExecutionManager() (common.Address, error) {
	return _BendExchange.Contract.ExecutionManager(&_BendExchange.CallOpts)
}

// InterceptorManager is a free data retrieval call binding the contract method 0xf3e13610.
//
// Solidity: function interceptorManager() view returns(address)
func (_BendExchange *BendExchangeCaller) InterceptorManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "interceptorManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterceptorManager is a free data retrieval call binding the contract method 0xf3e13610.
//
// Solidity: function interceptorManager() view returns(address)
func (_BendExchange *BendExchangeSession) InterceptorManager() (common.Address, error) {
	return _BendExchange.Contract.InterceptorManager(&_BendExchange.CallOpts)
}

// InterceptorManager is a free data retrieval call binding the contract method 0xf3e13610.
//
// Solidity: function interceptorManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) InterceptorManager() (common.Address, error) {
	return _BendExchange.Contract.InterceptorManager(&_BendExchange.CallOpts)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_BendExchange *BendExchangeCaller) IsUserOrderNonceExecutedOrCancelled(opts *bind.CallOpts, user common.Address, orderNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "isUserOrderNonceExecutedOrCancelled", user, orderNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_BendExchange *BendExchangeSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _BendExchange.Contract.IsUserOrderNonceExecutedOrCancelled(&_BendExchange.CallOpts, user, orderNonce)
}

// IsUserOrderNonceExecutedOrCancelled is a free data retrieval call binding the contract method 0x31e27e27.
//
// Solidity: function isUserOrderNonceExecutedOrCancelled(address user, uint256 orderNonce) view returns(bool)
func (_BendExchange *BendExchangeCallerSession) IsUserOrderNonceExecutedOrCancelled(user common.Address, orderNonce *big.Int) (bool, error) {
	return _BendExchange.Contract.IsUserOrderNonceExecutedOrCancelled(&_BendExchange.CallOpts, user, orderNonce)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendExchange *BendExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendExchange *BendExchangeSession) Owner() (common.Address, error) {
	return _BendExchange.Contract.Owner(&_BendExchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendExchange *BendExchangeCallerSession) Owner() (common.Address, error) {
	return _BendExchange.Contract.Owner(&_BendExchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_BendExchange *BendExchangeCaller) ProtocolFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "protocolFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_BendExchange *BendExchangeSession) ProtocolFeeRecipient() (common.Address, error) {
	return _BendExchange.Contract.ProtocolFeeRecipient(&_BendExchange.CallOpts)
}

// ProtocolFeeRecipient is a free data retrieval call binding the contract method 0x64df049e.
//
// Solidity: function protocolFeeRecipient() view returns(address)
func (_BendExchange *BendExchangeCallerSession) ProtocolFeeRecipient() (common.Address, error) {
	return _BendExchange.Contract.ProtocolFeeRecipient(&_BendExchange.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_BendExchange *BendExchangeCaller) RoyaltyFeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "royaltyFeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_BendExchange *BendExchangeSession) RoyaltyFeeManager() (common.Address, error) {
	return _BendExchange.Contract.RoyaltyFeeManager(&_BendExchange.CallOpts)
}

// RoyaltyFeeManager is a free data retrieval call binding the contract method 0x87e4401f.
//
// Solidity: function royaltyFeeManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) RoyaltyFeeManager() (common.Address, error) {
	return _BendExchange.Contract.RoyaltyFeeManager(&_BendExchange.CallOpts)
}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_BendExchange *BendExchangeCaller) TransferManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "transferManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_BendExchange *BendExchangeSession) TransferManager() (common.Address, error) {
	return _BendExchange.Contract.TransferManager(&_BendExchange.CallOpts)
}

// TransferManager is a free data retrieval call binding the contract method 0x46ea2552.
//
// Solidity: function transferManager() view returns(address)
func (_BendExchange *BendExchangeCallerSession) TransferManager() (common.Address, error) {
	return _BendExchange.Contract.TransferManager(&_BendExchange.CallOpts)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_BendExchange *BendExchangeCaller) UserMinOrderNonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendExchange.contract.Call(opts, &out, "userMinOrderNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_BendExchange *BendExchangeSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _BendExchange.Contract.UserMinOrderNonce(&_BendExchange.CallOpts, arg0)
}

// UserMinOrderNonce is a free data retrieval call binding the contract method 0x4266581e.
//
// Solidity: function userMinOrderNonce(address ) view returns(uint256)
func (_BendExchange *BendExchangeCallerSession) UserMinOrderNonce(arg0 common.Address) (*big.Int, error) {
	return _BendExchange.Contract.UserMinOrderNonce(&_BendExchange.CallOpts, arg0)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_BendExchange *BendExchangeTransactor) CancelAllOrdersForSender(opts *bind.TransactOpts, minNonce *big.Int) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "cancelAllOrdersForSender", minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_BendExchange *BendExchangeSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _BendExchange.Contract.CancelAllOrdersForSender(&_BendExchange.TransactOpts, minNonce)
}

// CancelAllOrdersForSender is a paid mutator transaction binding the contract method 0xcbd2ec65.
//
// Solidity: function cancelAllOrdersForSender(uint256 minNonce) returns()
func (_BendExchange *BendExchangeTransactorSession) CancelAllOrdersForSender(minNonce *big.Int) (*types.Transaction, error) {
	return _BendExchange.Contract.CancelAllOrdersForSender(&_BendExchange.TransactOpts, minNonce)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_BendExchange *BendExchangeTransactor) CancelMultipleMakerOrders(opts *bind.TransactOpts, orderNonces []*big.Int) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "cancelMultipleMakerOrders", orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_BendExchange *BendExchangeSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _BendExchange.Contract.CancelMultipleMakerOrders(&_BendExchange.TransactOpts, orderNonces)
}

// CancelMultipleMakerOrders is a paid mutator transaction binding the contract method 0x9e53a69a.
//
// Solidity: function cancelMultipleMakerOrders(uint256[] orderNonces) returns()
func (_BendExchange *BendExchangeTransactorSession) CancelMultipleMakerOrders(orderNonces []*big.Int) (*types.Transaction, error) {
	return _BendExchange.Contract.CancelMultipleMakerOrders(&_BendExchange.TransactOpts, orderNonces)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x49bfa369.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_BendExchange *BendExchangeTransactor) MatchAskWithTakerBid(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "matchAskWithTakerBid", takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x49bfa369.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_BendExchange *BendExchangeSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchAskWithTakerBid(&_BendExchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBid is a paid mutator transaction binding the contract method 0x49bfa369.
//
// Solidity: function matchAskWithTakerBid((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) returns()
func (_BendExchange *BendExchangeTransactorSession) MatchAskWithTakerBid(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchAskWithTakerBid(&_BendExchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0x1d8fbc11.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_BendExchange *BendExchangeTransactor) MatchAskWithTakerBidUsingETHAndWETH(opts *bind.TransactOpts, takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "matchAskWithTakerBidUsingETHAndWETH", takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0x1d8fbc11.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_BendExchange *BendExchangeSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_BendExchange.TransactOpts, takerBid, makerAsk)
}

// MatchAskWithTakerBidUsingETHAndWETH is a paid mutator transaction binding the contract method 0x1d8fbc11.
//
// Solidity: function matchAskWithTakerBidUsingETHAndWETH((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerBid, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerAsk) payable returns()
func (_BendExchange *BendExchangeTransactorSession) MatchAskWithTakerBidUsingETHAndWETH(takerBid OrderTypesTakerOrder, makerAsk OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchAskWithTakerBidUsingETHAndWETH(&_BendExchange.TransactOpts, takerBid, makerAsk)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0xb9c2ca2c.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_BendExchange *BendExchangeTransactor) MatchBidWithTakerAsk(opts *bind.TransactOpts, takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "matchBidWithTakerAsk", takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0xb9c2ca2c.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_BendExchange *BendExchangeSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchBidWithTakerAsk(&_BendExchange.TransactOpts, takerAsk, makerBid)
}

// MatchBidWithTakerAsk is a paid mutator transaction binding the contract method 0xb9c2ca2c.
//
// Solidity: function matchBidWithTakerAsk((bool,address,uint256,uint256,uint256,bytes,address,bytes) takerAsk, (bool,address,address,uint256,uint256,uint256,address,address,uint256,uint256,uint256,uint256,bytes,address,bytes,uint8,bytes32,bytes32) makerBid) returns()
func (_BendExchange *BendExchangeTransactorSession) MatchBidWithTakerAsk(takerAsk OrderTypesTakerOrder, makerBid OrderTypesMakerOrder) (*types.Transaction, error) {
	return _BendExchange.Contract.MatchBidWithTakerAsk(&_BendExchange.TransactOpts, takerAsk, makerBid)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendExchange *BendExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendExchange *BendExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendExchange.Contract.RenounceOwnership(&_BendExchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendExchange *BendExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendExchange.Contract.RenounceOwnership(&_BendExchange.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendExchange *BendExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendExchange *BendExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.TransferOwnership(&_BendExchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendExchange *BendExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.TransferOwnership(&_BendExchange.TransactOpts, newOwner)
}

// UpdateAuthorizationManager is a paid mutator transaction binding the contract method 0x7d50fcfd.
//
// Solidity: function updateAuthorizationManager(address _authorizationManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateAuthorizationManager(opts *bind.TransactOpts, _authorizationManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateAuthorizationManager", _authorizationManager)
}

// UpdateAuthorizationManager is a paid mutator transaction binding the contract method 0x7d50fcfd.
//
// Solidity: function updateAuthorizationManager(address _authorizationManager) returns()
func (_BendExchange *BendExchangeSession) UpdateAuthorizationManager(_authorizationManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateAuthorizationManager(&_BendExchange.TransactOpts, _authorizationManager)
}

// UpdateAuthorizationManager is a paid mutator transaction binding the contract method 0x7d50fcfd.
//
// Solidity: function updateAuthorizationManager(address _authorizationManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateAuthorizationManager(_authorizationManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateAuthorizationManager(&_BendExchange.TransactOpts, _authorizationManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateCurrencyManager(opts *bind.TransactOpts, _currencyManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateCurrencyManager", _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_BendExchange *BendExchangeSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateCurrencyManager(&_BendExchange.TransactOpts, _currencyManager)
}

// UpdateCurrencyManager is a paid mutator transaction binding the contract method 0x5ce052d7.
//
// Solidity: function updateCurrencyManager(address _currencyManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateCurrencyManager(_currencyManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateCurrencyManager(&_BendExchange.TransactOpts, _currencyManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateExecutionManager(opts *bind.TransactOpts, _executionManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateExecutionManager", _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_BendExchange *BendExchangeSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateExecutionManager(&_BendExchange.TransactOpts, _executionManager)
}

// UpdateExecutionManager is a paid mutator transaction binding the contract method 0xd4ff41dc.
//
// Solidity: function updateExecutionManager(address _executionManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateExecutionManager(_executionManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateExecutionManager(&_BendExchange.TransactOpts, _executionManager)
}

// UpdateInterceptorManager is a paid mutator transaction binding the contract method 0x923f13ae.
//
// Solidity: function updateInterceptorManager(address _interceptorManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateInterceptorManager(opts *bind.TransactOpts, _interceptorManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateInterceptorManager", _interceptorManager)
}

// UpdateInterceptorManager is a paid mutator transaction binding the contract method 0x923f13ae.
//
// Solidity: function updateInterceptorManager(address _interceptorManager) returns()
func (_BendExchange *BendExchangeSession) UpdateInterceptorManager(_interceptorManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateInterceptorManager(&_BendExchange.TransactOpts, _interceptorManager)
}

// UpdateInterceptorManager is a paid mutator transaction binding the contract method 0x923f13ae.
//
// Solidity: function updateInterceptorManager(address _interceptorManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateInterceptorManager(_interceptorManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateInterceptorManager(&_BendExchange.TransactOpts, _interceptorManager)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_BendExchange *BendExchangeTransactor) UpdateProtocolFeeRecipient(opts *bind.TransactOpts, _protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateProtocolFeeRecipient", _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_BendExchange *BendExchangeSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateProtocolFeeRecipient(&_BendExchange.TransactOpts, _protocolFeeRecipient)
}

// UpdateProtocolFeeRecipient is a paid mutator transaction binding the contract method 0x1df47f80.
//
// Solidity: function updateProtocolFeeRecipient(address _protocolFeeRecipient) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateProtocolFeeRecipient(_protocolFeeRecipient common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateProtocolFeeRecipient(&_BendExchange.TransactOpts, _protocolFeeRecipient)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateRoyaltyFeeManager(opts *bind.TransactOpts, _royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateRoyaltyFeeManager", _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_BendExchange *BendExchangeSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateRoyaltyFeeManager(&_BendExchange.TransactOpts, _royaltyFeeManager)
}

// UpdateRoyaltyFeeManager is a paid mutator transaction binding the contract method 0xc5498769.
//
// Solidity: function updateRoyaltyFeeManager(address _royaltyFeeManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateRoyaltyFeeManager(_royaltyFeeManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateRoyaltyFeeManager(&_BendExchange.TransactOpts, _royaltyFeeManager)
}

// UpdateTransferManager is a paid mutator transaction binding the contract method 0x2dd85ad9.
//
// Solidity: function updateTransferManager(address _transferManager) returns()
func (_BendExchange *BendExchangeTransactor) UpdateTransferManager(opts *bind.TransactOpts, _transferManager common.Address) (*types.Transaction, error) {
	return _BendExchange.contract.Transact(opts, "updateTransferManager", _transferManager)
}

// UpdateTransferManager is a paid mutator transaction binding the contract method 0x2dd85ad9.
//
// Solidity: function updateTransferManager(address _transferManager) returns()
func (_BendExchange *BendExchangeSession) UpdateTransferManager(_transferManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateTransferManager(&_BendExchange.TransactOpts, _transferManager)
}

// UpdateTransferManager is a paid mutator transaction binding the contract method 0x2dd85ad9.
//
// Solidity: function updateTransferManager(address _transferManager) returns()
func (_BendExchange *BendExchangeTransactorSession) UpdateTransferManager(_transferManager common.Address) (*types.Transaction, error) {
	return _BendExchange.Contract.UpdateTransferManager(&_BendExchange.TransactOpts, _transferManager)
}

// BendExchangeCancelAllOrdersIterator is returned from FilterCancelAllOrders and is used to iterate over the raw logs and unpacked data for CancelAllOrders events raised by the BendExchange contract.
type BendExchangeCancelAllOrdersIterator struct {
	Event *BendExchangeCancelAllOrders // Event containing the contract specifics and raw log

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
func (it *BendExchangeCancelAllOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeCancelAllOrders)
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
		it.Event = new(BendExchangeCancelAllOrders)
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
func (it *BendExchangeCancelAllOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeCancelAllOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeCancelAllOrders represents a CancelAllOrders event raised by the BendExchange contract.
type BendExchangeCancelAllOrders struct {
	User        common.Address
	NewMinNonce *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelAllOrders is a free log retrieval operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_BendExchange *BendExchangeFilterer) FilterCancelAllOrders(opts *bind.FilterOpts, user []common.Address) (*BendExchangeCancelAllOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeCancelAllOrdersIterator{contract: _BendExchange.contract, event: "CancelAllOrders", logs: logs, sub: sub}, nil
}

// WatchCancelAllOrders is a free log subscription operation binding the contract event 0x1e7178d84f0b0825c65795cd62e7972809ad3aac6917843aaec596161b2c0a97.
//
// Solidity: event CancelAllOrders(address indexed user, uint256 newMinNonce)
func (_BendExchange *BendExchangeFilterer) WatchCancelAllOrders(opts *bind.WatchOpts, sink chan<- *BendExchangeCancelAllOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "CancelAllOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeCancelAllOrders)
				if err := _BendExchange.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseCancelAllOrders(log types.Log) (*BendExchangeCancelAllOrders, error) {
	event := new(BendExchangeCancelAllOrders)
	if err := _BendExchange.contract.UnpackLog(event, "CancelAllOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeCancelMultipleOrdersIterator is returned from FilterCancelMultipleOrders and is used to iterate over the raw logs and unpacked data for CancelMultipleOrders events raised by the BendExchange contract.
type BendExchangeCancelMultipleOrdersIterator struct {
	Event *BendExchangeCancelMultipleOrders // Event containing the contract specifics and raw log

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
func (it *BendExchangeCancelMultipleOrdersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeCancelMultipleOrders)
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
		it.Event = new(BendExchangeCancelMultipleOrders)
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
func (it *BendExchangeCancelMultipleOrdersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeCancelMultipleOrdersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeCancelMultipleOrders represents a CancelMultipleOrders event raised by the BendExchange contract.
type BendExchangeCancelMultipleOrders struct {
	User        common.Address
	OrderNonces []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCancelMultipleOrders is a free log retrieval operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_BendExchange *BendExchangeFilterer) FilterCancelMultipleOrders(opts *bind.FilterOpts, user []common.Address) (*BendExchangeCancelMultipleOrdersIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeCancelMultipleOrdersIterator{contract: _BendExchange.contract, event: "CancelMultipleOrders", logs: logs, sub: sub}, nil
}

// WatchCancelMultipleOrders is a free log subscription operation binding the contract event 0xfa0ae5d80fe3763c880a3839fab0294171a6f730d1f82c4cd5392c6f67b41732.
//
// Solidity: event CancelMultipleOrders(address indexed user, uint256[] orderNonces)
func (_BendExchange *BendExchangeFilterer) WatchCancelMultipleOrders(opts *bind.WatchOpts, sink chan<- *BendExchangeCancelMultipleOrders, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "CancelMultipleOrders", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeCancelMultipleOrders)
				if err := _BendExchange.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseCancelMultipleOrders(log types.Log) (*BendExchangeCancelMultipleOrders, error) {
	event := new(BendExchangeCancelMultipleOrders)
	if err := _BendExchange.contract.UnpackLog(event, "CancelMultipleOrders", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewAuthorizationManagerIterator is returned from FilterNewAuthorizationManager and is used to iterate over the raw logs and unpacked data for NewAuthorizationManager events raised by the BendExchange contract.
type BendExchangeNewAuthorizationManagerIterator struct {
	Event *BendExchangeNewAuthorizationManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewAuthorizationManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewAuthorizationManager)
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
		it.Event = new(BendExchangeNewAuthorizationManager)
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
func (it *BendExchangeNewAuthorizationManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewAuthorizationManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewAuthorizationManager represents a NewAuthorizationManager event raised by the BendExchange contract.
type BendExchangeNewAuthorizationManager struct {
	AuthorizationManager common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewAuthorizationManager is a free log retrieval operation binding the contract event 0xcef74152554b7ba4c68515844f0d3e2f59b8560ff0e910f13c3a9b7413ec9939.
//
// Solidity: event NewAuthorizationManager(address indexed authorizationManager)
func (_BendExchange *BendExchangeFilterer) FilterNewAuthorizationManager(opts *bind.FilterOpts, authorizationManager []common.Address) (*BendExchangeNewAuthorizationManagerIterator, error) {

	var authorizationManagerRule []interface{}
	for _, authorizationManagerItem := range authorizationManager {
		authorizationManagerRule = append(authorizationManagerRule, authorizationManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewAuthorizationManager", authorizationManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewAuthorizationManagerIterator{contract: _BendExchange.contract, event: "NewAuthorizationManager", logs: logs, sub: sub}, nil
}

// WatchNewAuthorizationManager is a free log subscription operation binding the contract event 0xcef74152554b7ba4c68515844f0d3e2f59b8560ff0e910f13c3a9b7413ec9939.
//
// Solidity: event NewAuthorizationManager(address indexed authorizationManager)
func (_BendExchange *BendExchangeFilterer) WatchNewAuthorizationManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewAuthorizationManager, authorizationManager []common.Address) (event.Subscription, error) {

	var authorizationManagerRule []interface{}
	for _, authorizationManagerItem := range authorizationManager {
		authorizationManagerRule = append(authorizationManagerRule, authorizationManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewAuthorizationManager", authorizationManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewAuthorizationManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewAuthorizationManager", log); err != nil {
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

// ParseNewAuthorizationManager is a log parse operation binding the contract event 0xcef74152554b7ba4c68515844f0d3e2f59b8560ff0e910f13c3a9b7413ec9939.
//
// Solidity: event NewAuthorizationManager(address indexed authorizationManager)
func (_BendExchange *BendExchangeFilterer) ParseNewAuthorizationManager(log types.Log) (*BendExchangeNewAuthorizationManager, error) {
	event := new(BendExchangeNewAuthorizationManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewAuthorizationManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewCurrencyManagerIterator is returned from FilterNewCurrencyManager and is used to iterate over the raw logs and unpacked data for NewCurrencyManager events raised by the BendExchange contract.
type BendExchangeNewCurrencyManagerIterator struct {
	Event *BendExchangeNewCurrencyManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewCurrencyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewCurrencyManager)
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
		it.Event = new(BendExchangeNewCurrencyManager)
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
func (it *BendExchangeNewCurrencyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewCurrencyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewCurrencyManager represents a NewCurrencyManager event raised by the BendExchange contract.
type BendExchangeNewCurrencyManager struct {
	CurrencyManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewCurrencyManager is a free log retrieval operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_BendExchange *BendExchangeFilterer) FilterNewCurrencyManager(opts *bind.FilterOpts, currencyManager []common.Address) (*BendExchangeNewCurrencyManagerIterator, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewCurrencyManagerIterator{contract: _BendExchange.contract, event: "NewCurrencyManager", logs: logs, sub: sub}, nil
}

// WatchNewCurrencyManager is a free log subscription operation binding the contract event 0xb4f5db40df3aced29e88a4babbc3b46e305e07d34098525d18b1497056e63838.
//
// Solidity: event NewCurrencyManager(address indexed currencyManager)
func (_BendExchange *BendExchangeFilterer) WatchNewCurrencyManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewCurrencyManager, currencyManager []common.Address) (event.Subscription, error) {

	var currencyManagerRule []interface{}
	for _, currencyManagerItem := range currencyManager {
		currencyManagerRule = append(currencyManagerRule, currencyManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewCurrencyManager", currencyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewCurrencyManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseNewCurrencyManager(log types.Log) (*BendExchangeNewCurrencyManager, error) {
	event := new(BendExchangeNewCurrencyManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewCurrencyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewExecutionManagerIterator is returned from FilterNewExecutionManager and is used to iterate over the raw logs and unpacked data for NewExecutionManager events raised by the BendExchange contract.
type BendExchangeNewExecutionManagerIterator struct {
	Event *BendExchangeNewExecutionManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewExecutionManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewExecutionManager)
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
		it.Event = new(BendExchangeNewExecutionManager)
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
func (it *BendExchangeNewExecutionManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewExecutionManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewExecutionManager represents a NewExecutionManager event raised by the BendExchange contract.
type BendExchangeNewExecutionManager struct {
	ExecutionManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionManager is a free log retrieval operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_BendExchange *BendExchangeFilterer) FilterNewExecutionManager(opts *bind.FilterOpts, executionManager []common.Address) (*BendExchangeNewExecutionManagerIterator, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewExecutionManagerIterator{contract: _BendExchange.contract, event: "NewExecutionManager", logs: logs, sub: sub}, nil
}

// WatchNewExecutionManager is a free log subscription operation binding the contract event 0x36e2a376eabc3bc60cb88f29c288f53e36874a95a7f407330ab4f166b0905698.
//
// Solidity: event NewExecutionManager(address indexed executionManager)
func (_BendExchange *BendExchangeFilterer) WatchNewExecutionManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewExecutionManager, executionManager []common.Address) (event.Subscription, error) {

	var executionManagerRule []interface{}
	for _, executionManagerItem := range executionManager {
		executionManagerRule = append(executionManagerRule, executionManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewExecutionManager", executionManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewExecutionManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseNewExecutionManager(log types.Log) (*BendExchangeNewExecutionManager, error) {
	event := new(BendExchangeNewExecutionManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewExecutionManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewInterceptorManagerIterator is returned from FilterNewInterceptorManager and is used to iterate over the raw logs and unpacked data for NewInterceptorManager events raised by the BendExchange contract.
type BendExchangeNewInterceptorManagerIterator struct {
	Event *BendExchangeNewInterceptorManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewInterceptorManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewInterceptorManager)
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
		it.Event = new(BendExchangeNewInterceptorManager)
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
func (it *BendExchangeNewInterceptorManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewInterceptorManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewInterceptorManager represents a NewInterceptorManager event raised by the BendExchange contract.
type BendExchangeNewInterceptorManager struct {
	InterceptorManager common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewInterceptorManager is a free log retrieval operation binding the contract event 0x6a562813c3ab4feba79168d6fbb51d07c7e92461417fb2aa078ec811e49068d0.
//
// Solidity: event NewInterceptorManager(address indexed interceptorManager)
func (_BendExchange *BendExchangeFilterer) FilterNewInterceptorManager(opts *bind.FilterOpts, interceptorManager []common.Address) (*BendExchangeNewInterceptorManagerIterator, error) {

	var interceptorManagerRule []interface{}
	for _, interceptorManagerItem := range interceptorManager {
		interceptorManagerRule = append(interceptorManagerRule, interceptorManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewInterceptorManager", interceptorManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewInterceptorManagerIterator{contract: _BendExchange.contract, event: "NewInterceptorManager", logs: logs, sub: sub}, nil
}

// WatchNewInterceptorManager is a free log subscription operation binding the contract event 0x6a562813c3ab4feba79168d6fbb51d07c7e92461417fb2aa078ec811e49068d0.
//
// Solidity: event NewInterceptorManager(address indexed interceptorManager)
func (_BendExchange *BendExchangeFilterer) WatchNewInterceptorManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewInterceptorManager, interceptorManager []common.Address) (event.Subscription, error) {

	var interceptorManagerRule []interface{}
	for _, interceptorManagerItem := range interceptorManager {
		interceptorManagerRule = append(interceptorManagerRule, interceptorManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewInterceptorManager", interceptorManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewInterceptorManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewInterceptorManager", log); err != nil {
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

// ParseNewInterceptorManager is a log parse operation binding the contract event 0x6a562813c3ab4feba79168d6fbb51d07c7e92461417fb2aa078ec811e49068d0.
//
// Solidity: event NewInterceptorManager(address indexed interceptorManager)
func (_BendExchange *BendExchangeFilterer) ParseNewInterceptorManager(log types.Log) (*BendExchangeNewInterceptorManager, error) {
	event := new(BendExchangeNewInterceptorManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewInterceptorManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewProtocolFeeRecipientIterator is returned from FilterNewProtocolFeeRecipient and is used to iterate over the raw logs and unpacked data for NewProtocolFeeRecipient events raised by the BendExchange contract.
type BendExchangeNewProtocolFeeRecipientIterator struct {
	Event *BendExchangeNewProtocolFeeRecipient // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewProtocolFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewProtocolFeeRecipient)
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
		it.Event = new(BendExchangeNewProtocolFeeRecipient)
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
func (it *BendExchangeNewProtocolFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewProtocolFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewProtocolFeeRecipient represents a NewProtocolFeeRecipient event raised by the BendExchange contract.
type BendExchangeNewProtocolFeeRecipient struct {
	ProtocolFeeRecipient common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolFeeRecipient is a free log retrieval operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_BendExchange *BendExchangeFilterer) FilterNewProtocolFeeRecipient(opts *bind.FilterOpts, protocolFeeRecipient []common.Address) (*BendExchangeNewProtocolFeeRecipientIterator, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewProtocolFeeRecipientIterator{contract: _BendExchange.contract, event: "NewProtocolFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewProtocolFeeRecipient is a free log subscription operation binding the contract event 0x8cffb07faa2874440346743bdc0a86b06c3335cc47dc49b327d10e77b73ceb10.
//
// Solidity: event NewProtocolFeeRecipient(address indexed protocolFeeRecipient)
func (_BendExchange *BendExchangeFilterer) WatchNewProtocolFeeRecipient(opts *bind.WatchOpts, sink chan<- *BendExchangeNewProtocolFeeRecipient, protocolFeeRecipient []common.Address) (event.Subscription, error) {

	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewProtocolFeeRecipient", protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewProtocolFeeRecipient)
				if err := _BendExchange.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseNewProtocolFeeRecipient(log types.Log) (*BendExchangeNewProtocolFeeRecipient, error) {
	event := new(BendExchangeNewProtocolFeeRecipient)
	if err := _BendExchange.contract.UnpackLog(event, "NewProtocolFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewRoyaltyFeeManagerIterator is returned from FilterNewRoyaltyFeeManager and is used to iterate over the raw logs and unpacked data for NewRoyaltyFeeManager events raised by the BendExchange contract.
type BendExchangeNewRoyaltyFeeManagerIterator struct {
	Event *BendExchangeNewRoyaltyFeeManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewRoyaltyFeeManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewRoyaltyFeeManager)
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
		it.Event = new(BendExchangeNewRoyaltyFeeManager)
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
func (it *BendExchangeNewRoyaltyFeeManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewRoyaltyFeeManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewRoyaltyFeeManager represents a NewRoyaltyFeeManager event raised by the BendExchange contract.
type BendExchangeNewRoyaltyFeeManager struct {
	RoyaltyFeeManager common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewRoyaltyFeeManager is a free log retrieval operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_BendExchange *BendExchangeFilterer) FilterNewRoyaltyFeeManager(opts *bind.FilterOpts, royaltyFeeManager []common.Address) (*BendExchangeNewRoyaltyFeeManagerIterator, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewRoyaltyFeeManagerIterator{contract: _BendExchange.contract, event: "NewRoyaltyFeeManager", logs: logs, sub: sub}, nil
}

// WatchNewRoyaltyFeeManager is a free log subscription operation binding the contract event 0x80e3874461ebbd918ac3e81da0a92e5e51387d70f337237c9123e48d20e5a508.
//
// Solidity: event NewRoyaltyFeeManager(address indexed royaltyFeeManager)
func (_BendExchange *BendExchangeFilterer) WatchNewRoyaltyFeeManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewRoyaltyFeeManager, royaltyFeeManager []common.Address) (event.Subscription, error) {

	var royaltyFeeManagerRule []interface{}
	for _, royaltyFeeManagerItem := range royaltyFeeManager {
		royaltyFeeManagerRule = append(royaltyFeeManagerRule, royaltyFeeManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewRoyaltyFeeManager", royaltyFeeManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewRoyaltyFeeManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseNewRoyaltyFeeManager(log types.Log) (*BendExchangeNewRoyaltyFeeManager, error) {
	event := new(BendExchangeNewRoyaltyFeeManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewRoyaltyFeeManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeNewTransferManagerIterator is returned from FilterNewTransferManager and is used to iterate over the raw logs and unpacked data for NewTransferManager events raised by the BendExchange contract.
type BendExchangeNewTransferManagerIterator struct {
	Event *BendExchangeNewTransferManager // Event containing the contract specifics and raw log

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
func (it *BendExchangeNewTransferManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeNewTransferManager)
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
		it.Event = new(BendExchangeNewTransferManager)
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
func (it *BendExchangeNewTransferManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeNewTransferManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeNewTransferManager represents a NewTransferManager event raised by the BendExchange contract.
type BendExchangeNewTransferManager struct {
	TransferManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewTransferManager is a free log retrieval operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_BendExchange *BendExchangeFilterer) FilterNewTransferManager(opts *bind.FilterOpts, transferManager []common.Address) (*BendExchangeNewTransferManagerIterator, error) {

	var transferManagerRule []interface{}
	for _, transferManagerItem := range transferManager {
		transferManagerRule = append(transferManagerRule, transferManagerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "NewTransferManager", transferManagerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeNewTransferManagerIterator{contract: _BendExchange.contract, event: "NewTransferManager", logs: logs, sub: sub}, nil
}

// WatchNewTransferManager is a free log subscription operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_BendExchange *BendExchangeFilterer) WatchNewTransferManager(opts *bind.WatchOpts, sink chan<- *BendExchangeNewTransferManager, transferManager []common.Address) (event.Subscription, error) {

	var transferManagerRule []interface{}
	for _, transferManagerItem := range transferManager {
		transferManagerRule = append(transferManagerRule, transferManagerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "NewTransferManager", transferManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeNewTransferManager)
				if err := _BendExchange.contract.UnpackLog(event, "NewTransferManager", log); err != nil {
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

// ParseNewTransferManager is a log parse operation binding the contract event 0xbb34fee0d3bb7c602ecfe1c8bed62b67a720f718e0eaf72246d3457fcb25be17.
//
// Solidity: event NewTransferManager(address indexed transferManager)
func (_BendExchange *BendExchangeFilterer) ParseNewTransferManager(log types.Log) (*BendExchangeNewTransferManager, error) {
	event := new(BendExchangeNewTransferManager)
	if err := _BendExchange.contract.UnpackLog(event, "NewTransferManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BendExchange contract.
type BendExchangeOwnershipTransferredIterator struct {
	Event *BendExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BendExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeOwnershipTransferred)
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
		it.Event = new(BendExchangeOwnershipTransferred)
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
func (it *BendExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the BendExchange contract.
type BendExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendExchange *BendExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BendExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeOwnershipTransferredIterator{contract: _BendExchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendExchange *BendExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BendExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeOwnershipTransferred)
				if err := _BendExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*BendExchangeOwnershipTransferred, error) {
	event := new(BendExchangeOwnershipTransferred)
	if err := _BendExchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeProtocolFeePaymentIterator is returned from FilterProtocolFeePayment and is used to iterate over the raw logs and unpacked data for ProtocolFeePayment events raised by the BendExchange contract.
type BendExchangeProtocolFeePaymentIterator struct {
	Event *BendExchangeProtocolFeePayment // Event containing the contract specifics and raw log

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
func (it *BendExchangeProtocolFeePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeProtocolFeePayment)
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
		it.Event = new(BendExchangeProtocolFeePayment)
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
func (it *BendExchangeProtocolFeePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeProtocolFeePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeProtocolFeePayment represents a ProtocolFeePayment event raised by the BendExchange contract.
type BendExchangeProtocolFeePayment struct {
	Collection           common.Address
	TokenId              *big.Int
	ProtocolFeeRecipient common.Address
	Currency             common.Address
	Amount               *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeePayment is a free log retrieval operation binding the contract event 0x4e2800017b302274fac4f45ad76bf09bfddec9f66525450dff12dcca91344f86.
//
// Solidity: event ProtocolFeePayment(address indexed collection, uint256 indexed tokenId, address indexed protocolFeeRecipient, address currency, uint256 amount)
func (_BendExchange *BendExchangeFilterer) FilterProtocolFeePayment(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, protocolFeeRecipient []common.Address) (*BendExchangeProtocolFeePaymentIterator, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "ProtocolFeePayment", collectionRule, tokenIdRule, protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeProtocolFeePaymentIterator{contract: _BendExchange.contract, event: "ProtocolFeePayment", logs: logs, sub: sub}, nil
}

// WatchProtocolFeePayment is a free log subscription operation binding the contract event 0x4e2800017b302274fac4f45ad76bf09bfddec9f66525450dff12dcca91344f86.
//
// Solidity: event ProtocolFeePayment(address indexed collection, uint256 indexed tokenId, address indexed protocolFeeRecipient, address currency, uint256 amount)
func (_BendExchange *BendExchangeFilterer) WatchProtocolFeePayment(opts *bind.WatchOpts, sink chan<- *BendExchangeProtocolFeePayment, collection []common.Address, tokenId []*big.Int, protocolFeeRecipient []common.Address) (event.Subscription, error) {

	var collectionRule []interface{}
	for _, collectionItem := range collection {
		collectionRule = append(collectionRule, collectionItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var protocolFeeRecipientRule []interface{}
	for _, protocolFeeRecipientItem := range protocolFeeRecipient {
		protocolFeeRecipientRule = append(protocolFeeRecipientRule, protocolFeeRecipientItem)
	}

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "ProtocolFeePayment", collectionRule, tokenIdRule, protocolFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeProtocolFeePayment)
				if err := _BendExchange.contract.UnpackLog(event, "ProtocolFeePayment", log); err != nil {
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

// ParseProtocolFeePayment is a log parse operation binding the contract event 0x4e2800017b302274fac4f45ad76bf09bfddec9f66525450dff12dcca91344f86.
//
// Solidity: event ProtocolFeePayment(address indexed collection, uint256 indexed tokenId, address indexed protocolFeeRecipient, address currency, uint256 amount)
func (_BendExchange *BendExchangeFilterer) ParseProtocolFeePayment(log types.Log) (*BendExchangeProtocolFeePayment, error) {
	event := new(BendExchangeProtocolFeePayment)
	if err := _BendExchange.contract.UnpackLog(event, "ProtocolFeePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeRoyaltyPaymentIterator is returned from FilterRoyaltyPayment and is used to iterate over the raw logs and unpacked data for RoyaltyPayment events raised by the BendExchange contract.
type BendExchangeRoyaltyPaymentIterator struct {
	Event *BendExchangeRoyaltyPayment // Event containing the contract specifics and raw log

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
func (it *BendExchangeRoyaltyPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeRoyaltyPayment)
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
		it.Event = new(BendExchangeRoyaltyPayment)
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
func (it *BendExchangeRoyaltyPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeRoyaltyPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeRoyaltyPayment represents a RoyaltyPayment event raised by the BendExchange contract.
type BendExchangeRoyaltyPayment struct {
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
func (_BendExchange *BendExchangeFilterer) FilterRoyaltyPayment(opts *bind.FilterOpts, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (*BendExchangeRoyaltyPaymentIterator, error) {

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

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeRoyaltyPaymentIterator{contract: _BendExchange.contract, event: "RoyaltyPayment", logs: logs, sub: sub}, nil
}

// WatchRoyaltyPayment is a free log subscription operation binding the contract event 0x27c4f0403323142b599832f26acd21c74a9e5b809f2215726e244a4ac588cd7d.
//
// Solidity: event RoyaltyPayment(address indexed collection, uint256 indexed tokenId, address indexed royaltyRecipient, address currency, uint256 amount)
func (_BendExchange *BendExchangeFilterer) WatchRoyaltyPayment(opts *bind.WatchOpts, sink chan<- *BendExchangeRoyaltyPayment, collection []common.Address, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "RoyaltyPayment", collectionRule, tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeRoyaltyPayment)
				if err := _BendExchange.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
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
func (_BendExchange *BendExchangeFilterer) ParseRoyaltyPayment(log types.Log) (*BendExchangeRoyaltyPayment, error) {
	event := new(BendExchangeRoyaltyPayment)
	if err := _BendExchange.contract.UnpackLog(event, "RoyaltyPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeTakerAskIterator is returned from FilterTakerAsk and is used to iterate over the raw logs and unpacked data for TakerAsk events raised by the BendExchange contract.
type BendExchangeTakerAskIterator struct {
	Event *BendExchangeTakerAsk // Event containing the contract specifics and raw log

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
func (it *BendExchangeTakerAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeTakerAsk)
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
		it.Event = new(BendExchangeTakerAsk)
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
func (it *BendExchangeTakerAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeTakerAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeTakerAsk represents a TakerAsk event raised by the BendExchange contract.
type BendExchangeTakerAsk struct {
	MakerOrderHash [32]byte
	OrderNonce     *big.Int
	Taker          common.Address
	Maker          common.Address
	Strategy       common.Address
	Currency       common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTakerAsk is a free log retrieval operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) FilterTakerAsk(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*BendExchangeTakerAskIterator, error) {

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

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeTakerAskIterator{contract: _BendExchange.contract, event: "TakerAsk", logs: logs, sub: sub}, nil
}

// WatchTakerAsk is a free log subscription operation binding the contract event 0x68cd251d4d267c6e2034ff0088b990352b97b2002c0476587d0c4da889c11330.
//
// Solidity: event TakerAsk(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) WatchTakerAsk(opts *bind.WatchOpts, sink chan<- *BendExchangeTakerAsk, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "TakerAsk", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeTakerAsk)
				if err := _BendExchange.contract.UnpackLog(event, "TakerAsk", log); err != nil {
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
// Solidity: event TakerAsk(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) ParseTakerAsk(log types.Log) (*BendExchangeTakerAsk, error) {
	event := new(BendExchangeTakerAsk)
	if err := _BendExchange.contract.UnpackLog(event, "TakerAsk", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendExchangeTakerBidIterator is returned from FilterTakerBid and is used to iterate over the raw logs and unpacked data for TakerBid events raised by the BendExchange contract.
type BendExchangeTakerBidIterator struct {
	Event *BendExchangeTakerBid // Event containing the contract specifics and raw log

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
func (it *BendExchangeTakerBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendExchangeTakerBid)
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
		it.Event = new(BendExchangeTakerBid)
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
func (it *BendExchangeTakerBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendExchangeTakerBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendExchangeTakerBid represents a TakerBid event raised by the BendExchange contract.
type BendExchangeTakerBid struct {
	MakerOrderHash [32]byte
	OrderNonce     *big.Int
	Taker          common.Address
	Maker          common.Address
	Strategy       common.Address
	Currency       common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	Price          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTakerBid is a free log retrieval operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) FilterTakerBid(opts *bind.FilterOpts, taker []common.Address, maker []common.Address, strategy []common.Address) (*BendExchangeTakerBidIterator, error) {

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

	logs, sub, err := _BendExchange.contract.FilterLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return &BendExchangeTakerBidIterator{contract: _BendExchange.contract, event: "TakerBid", logs: logs, sub: sub}, nil
}

// WatchTakerBid is a free log subscription operation binding the contract event 0x95fb6205e23ff6bda16a2d1dba56b9ad7c783f67c96fa149785052f47696f2be.
//
// Solidity: event TakerBid(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) WatchTakerBid(opts *bind.WatchOpts, sink chan<- *BendExchangeTakerBid, taker []common.Address, maker []common.Address, strategy []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BendExchange.contract.WatchLogs(opts, "TakerBid", takerRule, makerRule, strategyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendExchangeTakerBid)
				if err := _BendExchange.contract.UnpackLog(event, "TakerBid", log); err != nil {
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
// Solidity: event TakerBid(bytes32 makerOrderHash, uint256 orderNonce, address indexed taker, address indexed maker, address indexed strategy, address currency, address collection, uint256 tokenId, uint256 amount, uint256 price)
func (_BendExchange *BendExchangeFilterer) ParseTakerBid(log types.Log) (*BendExchangeTakerBid, error) {
	event := new(BendExchangeTakerBid)
	if err := _BendExchange.contract.UnpackLog(event, "TakerBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
