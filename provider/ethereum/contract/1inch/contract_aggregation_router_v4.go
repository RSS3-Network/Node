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

// AggregationRouterV4SwapDescription is an auto generated low-level Go binding around an user-defined struct.
type AggregationRouterV4SwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
	Permit          []byte
}

// LimitOrderProtocolRFQOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type LimitOrderProtocolRFQOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// AggregationRouterV4MetaData contains all meta data concerning the AggregationRouterV4 contract.
var AggregationRouterV4MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"_clipperExchange\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIMIT_ORDER_RFQ_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"clipperSwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLimitOrderProtocolRFQ.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLimitOrderProtocolRFQ.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structLimitOrderProtocolRFQ.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"caller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationRouterV4.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLeft\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"pools\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AggregationRouterV4ABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregationRouterV4MetaData.ABI instead.
var AggregationRouterV4ABI = AggregationRouterV4MetaData.ABI

// AggregationRouterV4 is an auto generated Go binding around an Ethereum contract.
type AggregationRouterV4 struct {
	AggregationRouterV4Caller     // Read-only binding to the contract
	AggregationRouterV4Transactor // Write-only binding to the contract
	AggregationRouterV4Filterer   // Log filterer for contract events
}

// AggregationRouterV4Caller is an auto generated read-only Go binding around an Ethereum contract.
type AggregationRouterV4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregationRouterV4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregationRouterV4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationRouterV4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregationRouterV4Session struct {
	Contract     *AggregationRouterV4 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AggregationRouterV4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregationRouterV4CallerSession struct {
	Contract *AggregationRouterV4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AggregationRouterV4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregationRouterV4TransactorSession struct {
	Contract     *AggregationRouterV4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AggregationRouterV4Raw is an auto generated low-level Go binding around an Ethereum contract.
type AggregationRouterV4Raw struct {
	Contract *AggregationRouterV4 // Generic contract binding to access the raw methods on
}

// AggregationRouterV4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregationRouterV4CallerRaw struct {
	Contract *AggregationRouterV4Caller // Generic read-only contract binding to access the raw methods on
}

// AggregationRouterV4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregationRouterV4TransactorRaw struct {
	Contract *AggregationRouterV4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregationRouterV4 creates a new instance of AggregationRouterV4, bound to a specific deployed contract.
func NewAggregationRouterV4(address common.Address, backend bind.ContractBackend) (*AggregationRouterV4, error) {
	contract, err := bindAggregationRouterV4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4{AggregationRouterV4Caller: AggregationRouterV4Caller{contract: contract}, AggregationRouterV4Transactor: AggregationRouterV4Transactor{contract: contract}, AggregationRouterV4Filterer: AggregationRouterV4Filterer{contract: contract}}, nil
}

// NewAggregationRouterV4Caller creates a new read-only instance of AggregationRouterV4, bound to a specific deployed contract.
func NewAggregationRouterV4Caller(address common.Address, caller bind.ContractCaller) (*AggregationRouterV4Caller, error) {
	contract, err := bindAggregationRouterV4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4Caller{contract: contract}, nil
}

// NewAggregationRouterV4Transactor creates a new write-only instance of AggregationRouterV4, bound to a specific deployed contract.
func NewAggregationRouterV4Transactor(address common.Address, transactor bind.ContractTransactor) (*AggregationRouterV4Transactor, error) {
	contract, err := bindAggregationRouterV4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4Transactor{contract: contract}, nil
}

// NewAggregationRouterV4Filterer creates a new log filterer instance of AggregationRouterV4, bound to a specific deployed contract.
func NewAggregationRouterV4Filterer(address common.Address, filterer bind.ContractFilterer) (*AggregationRouterV4Filterer, error) {
	contract, err := bindAggregationRouterV4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4Filterer{contract: contract}, nil
}

// bindAggregationRouterV4 binds a generic wrapper to an already deployed contract.
func bindAggregationRouterV4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregationRouterV4MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV4 *AggregationRouterV4Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV4.Contract.AggregationRouterV4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV4 *AggregationRouterV4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.AggregationRouterV4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV4 *AggregationRouterV4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.AggregationRouterV4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationRouterV4 *AggregationRouterV4CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationRouterV4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationRouterV4 *AggregationRouterV4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationRouterV4 *AggregationRouterV4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4Caller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationRouterV4.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4Session) DOMAINSEPARATOR() ([32]byte, error) {
	return _AggregationRouterV4.Contract.DOMAINSEPARATOR(&_AggregationRouterV4.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4CallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AggregationRouterV4.Contract.DOMAINSEPARATOR(&_AggregationRouterV4.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4Caller) LIMITORDERRFQTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationRouterV4.contract.Call(opts, &out, "LIMIT_ORDER_RFQ_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4Session) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _AggregationRouterV4.Contract.LIMITORDERRFQTYPEHASH(&_AggregationRouterV4.CallOpts)
}

// LIMITORDERRFQTYPEHASH is a free data retrieval call binding the contract method 0x06bf53d0.
//
// Solidity: function LIMIT_ORDER_RFQ_TYPEHASH() view returns(bytes32)
func (_AggregationRouterV4 *AggregationRouterV4CallerSession) LIMITORDERRFQTYPEHASH() ([32]byte, error) {
	return _AggregationRouterV4.Contract.LIMITORDERRFQTYPEHASH(&_AggregationRouterV4.CallOpts)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_AggregationRouterV4 *AggregationRouterV4Caller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AggregationRouterV4.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_AggregationRouterV4 *AggregationRouterV4Session) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _AggregationRouterV4.Contract.InvalidatorForOrderRFQ(&_AggregationRouterV4.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_AggregationRouterV4 *AggregationRouterV4CallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _AggregationRouterV4.Contract.InvalidatorForOrderRFQ(&_AggregationRouterV4.CallOpts, maker, slot)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV4 *AggregationRouterV4Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregationRouterV4.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV4 *AggregationRouterV4Session) Owner() (common.Address, error) {
	return _AggregationRouterV4.Contract.Owner(&_AggregationRouterV4.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationRouterV4 *AggregationRouterV4CallerSession) Owner() (common.Address, error) {
	return _AggregationRouterV4.Contract.Owner(&_AggregationRouterV4.CallOpts)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.CancelOrderRFQ(&_AggregationRouterV4.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.CancelOrderRFQ(&_AggregationRouterV4.TransactOpts, orderInfo)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xb0431182.
//
// Solidity: function clipperSwap(address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) ClipperSwap(opts *bind.TransactOpts, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "clipperSwap", srcToken, dstToken, amount, minReturn)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xb0431182.
//
// Solidity: function clipperSwap(address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) ClipperSwap(srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwap(&_AggregationRouterV4.TransactOpts, srcToken, dstToken, amount, minReturn)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0xb0431182.
//
// Solidity: function clipperSwap(address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) ClipperSwap(srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwap(&_AggregationRouterV4.TransactOpts, srcToken, dstToken, amount, minReturn)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x9994dd15.
//
// Solidity: function clipperSwapTo(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) ClipperSwapTo(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "clipperSwapTo", recipient, srcToken, dstToken, amount, minReturn)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x9994dd15.
//
// Solidity: function clipperSwapTo(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) ClipperSwapTo(recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwapTo(&_AggregationRouterV4.TransactOpts, recipient, srcToken, dstToken, amount, minReturn)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x9994dd15.
//
// Solidity: function clipperSwapTo(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) ClipperSwapTo(recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwapTo(&_AggregationRouterV4.TransactOpts, recipient, srcToken, dstToken, amount, minReturn)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xd6a92a5d.
//
// Solidity: function clipperSwapToWithPermit(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) ClipperSwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "clipperSwapToWithPermit", recipient, srcToken, dstToken, amount, minReturn, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xd6a92a5d.
//
// Solidity: function clipperSwapToWithPermit(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) ClipperSwapToWithPermit(recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwapToWithPermit(&_AggregationRouterV4.TransactOpts, recipient, srcToken, dstToken, amount, minReturn, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xd6a92a5d.
//
// Solidity: function clipperSwapToWithPermit(address recipient, address srcToken, address dstToken, uint256 amount, uint256 minReturn, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) ClipperSwapToWithPermit(recipient common.Address, srcToken common.Address, dstToken common.Address, amount *big.Int, minReturn *big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.ClipperSwapToWithPermit(&_AggregationRouterV4.TransactOpts, recipient, srcToken, dstToken, amount, minReturn, permit)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) Destroy() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Destroy(&_AggregationRouterV4.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) Destroy() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Destroy(&_AggregationRouterV4.TransactOpts)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) FillOrderRFQ(opts *bind.TransactOpts, order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "fillOrderRFQ", order, signature, makingAmount, takingAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Session) FillOrderRFQ(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQ(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0xd0a3b665.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) FillOrderRFQ(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQ(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) FillOrderRFQTo(opts *bind.TransactOpts, order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "fillOrderRFQTo", order, signature, makingAmount, takingAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Session) FillOrderRFQTo(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQTo(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0xbaba5855.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target) payable returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) FillOrderRFQTo(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQTo(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount, target)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x4cc4a27b.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target, bytes permit) returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) FillOrderRFQToWithPermit(opts *bind.TransactOpts, order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "fillOrderRFQToWithPermit", order, signature, makingAmount, takingAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x4cc4a27b.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target, bytes permit) returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4Session) FillOrderRFQToWithPermit(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQToWithPermit(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x4cc4a27b.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 makingAmount, uint256 takingAmount, address target, bytes permit) returns(uint256, uint256)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) FillOrderRFQToWithPermit(order LimitOrderProtocolRFQOrderRFQ, signature []byte, makingAmount *big.Int, takingAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.FillOrderRFQToWithPermit(&_AggregationRouterV4.TransactOpts, order, signature, makingAmount, takingAmount, target, permit)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.RenounceOwnership(&_AggregationRouterV4.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.RenounceOwnership(&_AggregationRouterV4.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.RescueFunds(&_AggregationRouterV4.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.RescueFunds(&_AggregationRouterV4.TransactOpts, token, amount)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount, uint256 gasLeft)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) Swap(opts *bind.TransactOpts, caller common.Address, desc AggregationRouterV4SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "swap", caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount, uint256 gasLeft)
func (_AggregationRouterV4 *AggregationRouterV4Session) Swap(caller common.Address, desc AggregationRouterV4SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Swap(&_AggregationRouterV4.TransactOpts, caller, desc, data)
}

// Swap is a paid mutator transaction binding the contract method 0x7c025200.
//
// Solidity: function swap(address caller, (address,address,address,address,uint256,uint256,uint256,bytes) desc, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount, uint256 gasLeft)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) Swap(caller common.Address, desc AggregationRouterV4SwapDescription, data []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Swap(&_AggregationRouterV4.TransactOpts, caller, desc, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.TransferOwnership(&_AggregationRouterV4.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.TransferOwnership(&_AggregationRouterV4.TransactOpts, newOwner)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) UniswapV3Swap(opts *bind.TransactOpts, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "uniswapV3Swap", amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3Swap(&_AggregationRouterV4.TransactOpts, amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3Swap(&_AggregationRouterV4.TransactOpts, amount, minReturn, pools)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapCallback(&_AggregationRouterV4.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapCallback(&_AggregationRouterV4.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) UniswapV3SwapTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapTo(&_AggregationRouterV4.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapTo(&_AggregationRouterV4.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) UniswapV3SwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapToWithPermit(&_AggregationRouterV4.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UniswapV3SwapToWithPermit(&_AggregationRouterV4.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Unoswap(&_AggregationRouterV4.TransactOpts, srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x2e95b6c8.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools) payable returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Unoswap(&_AggregationRouterV4.TransactOpts, srcToken, amount, minReturn, pools)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Transactor) UnoswapWithPermit(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.Transact(opts, "unoswapWithPermit", srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4Session) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UnoswapWithPermit(&_AggregationRouterV4.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// UnoswapWithPermit is a paid mutator transaction binding the contract method 0xa1251d75.
//
// Solidity: function unoswapWithPermit(address srcToken, uint256 amount, uint256 minReturn, bytes32[] pools, bytes permit) returns(uint256 returnAmount)
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) UnoswapWithPermit(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools [][32]byte, permit []byte) (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.UnoswapWithPermit(&_AggregationRouterV4.TransactOpts, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV4 *AggregationRouterV4Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationRouterV4.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV4 *AggregationRouterV4Session) Receive() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Receive(&_AggregationRouterV4.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AggregationRouterV4 *AggregationRouterV4TransactorSession) Receive() (*types.Transaction, error) {
	return _AggregationRouterV4.Contract.Receive(&_AggregationRouterV4.TransactOpts)
}

// AggregationRouterV4OrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the AggregationRouterV4 contract.
type AggregationRouterV4OrderFilledRFQIterator struct {
	Event *AggregationRouterV4OrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV4OrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV4OrderFilledRFQ)
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
		it.Event = new(AggregationRouterV4OrderFilledRFQ)
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
func (it *AggregationRouterV4OrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV4OrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV4OrderFilledRFQ represents a OrderFilledRFQ event raised by the AggregationRouterV4 contract.
type AggregationRouterV4OrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_AggregationRouterV4 *AggregationRouterV4Filterer) FilterOrderFilledRFQ(opts *bind.FilterOpts) (*AggregationRouterV4OrderFilledRFQIterator, error) {

	logs, sub, err := _AggregationRouterV4.contract.FilterLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4OrderFilledRFQIterator{contract: _AggregationRouterV4.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_AggregationRouterV4 *AggregationRouterV4Filterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *AggregationRouterV4OrderFilledRFQ) (event.Subscription, error) {

	logs, sub, err := _AggregationRouterV4.contract.WatchLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV4OrderFilledRFQ)
				if err := _AggregationRouterV4.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_AggregationRouterV4 *AggregationRouterV4Filterer) ParseOrderFilledRFQ(log types.Log) (*AggregationRouterV4OrderFilledRFQ, error) {
	event := new(AggregationRouterV4OrderFilledRFQ)
	if err := _AggregationRouterV4.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationRouterV4OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AggregationRouterV4 contract.
type AggregationRouterV4OwnershipTransferredIterator struct {
	Event *AggregationRouterV4OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregationRouterV4OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationRouterV4OwnershipTransferred)
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
		it.Event = new(AggregationRouterV4OwnershipTransferred)
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
func (it *AggregationRouterV4OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationRouterV4OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationRouterV4OwnershipTransferred represents a OwnershipTransferred event raised by the AggregationRouterV4 contract.
type AggregationRouterV4OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV4 *AggregationRouterV4Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregationRouterV4OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV4.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregationRouterV4OwnershipTransferredIterator{contract: _AggregationRouterV4.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationRouterV4 *AggregationRouterV4Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregationRouterV4OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationRouterV4.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationRouterV4OwnershipTransferred)
				if err := _AggregationRouterV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AggregationRouterV4 *AggregationRouterV4Filterer) ParseOwnershipTransferred(log types.Log) (*AggregationRouterV4OwnershipTransferred, error) {
	event := new(AggregationRouterV4OwnershipTransferred)
	if err := _AggregationRouterV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
