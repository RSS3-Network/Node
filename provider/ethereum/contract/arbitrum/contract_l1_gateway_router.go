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

// L1GatewayRouterMetaData contains all meta data concerning the L1GatewayRouter contract.
var L1GatewayRouterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDefaultGateway\",\"type\":\"address\"}],\"name\":\"DefaultGatewayUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"GatewaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userFrom\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_userTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"name\":\"TransferRouted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_seqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TxToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSource\",\"type\":\"address\"}],\"name\":\"WhitelistSourceUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1ERC20\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"gateway\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_counterpartGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1TokenToGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransferCustomRefund\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newL1DefaultGateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"setDefaultGateway\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creditBackAddress\",\"type\":\"address\"}],\"name\":\"setGateway\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"setGateway\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_token\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateway\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"setGateways\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSource\",\"type\":\"address\"}],\"name\":\"updateWhitelistSource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1GatewayRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use L1GatewayRouterMetaData.ABI instead.
var L1GatewayRouterABI = L1GatewayRouterMetaData.ABI

// L1GatewayRouter is an auto generated Go binding around an Ethereum contract.
type L1GatewayRouter struct {
	L1GatewayRouterCaller     // Read-only binding to the contract
	L1GatewayRouterTransactor // Write-only binding to the contract
	L1GatewayRouterFilterer   // Log filterer for contract events
}

// L1GatewayRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1GatewayRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1GatewayRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1GatewayRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1GatewayRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1GatewayRouterSession struct {
	Contract     *L1GatewayRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1GatewayRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1GatewayRouterCallerSession struct {
	Contract *L1GatewayRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1GatewayRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1GatewayRouterTransactorSession struct {
	Contract     *L1GatewayRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1GatewayRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1GatewayRouterRaw struct {
	Contract *L1GatewayRouter // Generic contract binding to access the raw methods on
}

// L1GatewayRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1GatewayRouterCallerRaw struct {
	Contract *L1GatewayRouterCaller // Generic read-only contract binding to access the raw methods on
}

// L1GatewayRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1GatewayRouterTransactorRaw struct {
	Contract *L1GatewayRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1GatewayRouter creates a new instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouter(address common.Address, backend bind.ContractBackend) (*L1GatewayRouter, error) {
	contract, err := bindL1GatewayRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouter{L1GatewayRouterCaller: L1GatewayRouterCaller{contract: contract}, L1GatewayRouterTransactor: L1GatewayRouterTransactor{contract: contract}, L1GatewayRouterFilterer: L1GatewayRouterFilterer{contract: contract}}, nil
}

// NewL1GatewayRouterCaller creates a new read-only instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterCaller(address common.Address, caller bind.ContractCaller) (*L1GatewayRouterCaller, error) {
	contract, err := bindL1GatewayRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterCaller{contract: contract}, nil
}

// NewL1GatewayRouterTransactor creates a new write-only instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*L1GatewayRouterTransactor, error) {
	contract, err := bindL1GatewayRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterTransactor{contract: contract}, nil
}

// NewL1GatewayRouterFilterer creates a new log filterer instance of L1GatewayRouter, bound to a specific deployed contract.
func NewL1GatewayRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*L1GatewayRouterFilterer, error) {
	contract, err := bindL1GatewayRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterFilterer{contract: contract}, nil
}

// bindL1GatewayRouter binds a generic wrapper to an already deployed contract.
func bindL1GatewayRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1GatewayRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GatewayRouter *L1GatewayRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GatewayRouter.Contract.L1GatewayRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GatewayRouter *L1GatewayRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.L1GatewayRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GatewayRouter *L1GatewayRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.L1GatewayRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1GatewayRouter *L1GatewayRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1GatewayRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1GatewayRouter *L1GatewayRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1GatewayRouter *L1GatewayRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1ERC20 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "calculateL2TokenAddress", l1ERC20)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.CalculateL2TokenAddress(&_L1GatewayRouter.CallOpts, l1ERC20)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.CalculateL2TokenAddress(&_L1GatewayRouter.CallOpts, l1ERC20)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) CounterpartGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.CounterpartGateway(&_L1GatewayRouter.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) CounterpartGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.CounterpartGateway(&_L1GatewayRouter.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) DefaultGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "defaultGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) DefaultGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.DefaultGateway(&_L1GatewayRouter.CallOpts)
}

// DefaultGateway is a free data retrieval call binding the contract method 0x03295802.
//
// Solidity: function defaultGateway() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) DefaultGateway() (common.Address, error) {
	return _L1GatewayRouter.Contract.DefaultGateway(&_L1GatewayRouter.CallOpts)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_L1GatewayRouter *L1GatewayRouterCaller) GetGateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "getGateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_L1GatewayRouter *L1GatewayRouterSession) GetGateway(_token common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetGateway(&_L1GatewayRouter.CallOpts, _token)
}

// GetGateway is a free data retrieval call binding the contract method 0xbda009fe.
//
// Solidity: function getGateway(address _token) view returns(address gateway)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) GetGateway(_token common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.GetGateway(&_L1GatewayRouter.CallOpts, _token)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterCaller) GetOutboundCalldata(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "getOutboundCalldata", _token, _from, _to, _amount, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1GatewayRouter.Contract.GetOutboundCalldata(&_L1GatewayRouter.CallOpts, _token, _from, _to, _amount, _data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1GatewayRouter.Contract.GetOutboundCalldata(&_L1GatewayRouter.CallOpts, _token, _from, _to, _amount, _data)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) Inbox() (common.Address, error) {
	return _L1GatewayRouter.Contract.Inbox(&_L1GatewayRouter.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) Inbox() (common.Address, error) {
	return _L1GatewayRouter.Contract.Inbox(&_L1GatewayRouter.CallOpts)
}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) L1TokenToGateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "l1TokenToGateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) L1TokenToGateway(arg0 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.L1TokenToGateway(&_L1GatewayRouter.CallOpts, arg0)
}

// L1TokenToGateway is a free data retrieval call binding the contract method 0xed08fdc6.
//
// Solidity: function l1TokenToGateway(address ) view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) L1TokenToGateway(arg0 common.Address) (common.Address, error) {
	return _L1GatewayRouter.Contract.L1TokenToGateway(&_L1GatewayRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) Owner() (common.Address, error) {
	return _L1GatewayRouter.Contract.Owner(&_L1GatewayRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) Owner() (common.Address, error) {
	return _L1GatewayRouter.Contract.Owner(&_L1GatewayRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) Router() (common.Address, error) {
	return _L1GatewayRouter.Contract.Router(&_L1GatewayRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) Router() (common.Address, error) {
	return _L1GatewayRouter.Contract.Router(&_L1GatewayRouter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1GatewayRouter *L1GatewayRouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1GatewayRouter *L1GatewayRouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1GatewayRouter.Contract.SupportsInterface(&_L1GatewayRouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1GatewayRouter.Contract.SupportsInterface(&_L1GatewayRouter.CallOpts, interfaceId)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1GatewayRouter.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterSession) Whitelist() (common.Address, error) {
	return _L1GatewayRouter.Contract.Whitelist(&_L1GatewayRouter.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1GatewayRouter *L1GatewayRouterCallerSession) Whitelist() (common.Address, error) {
	return _L1GatewayRouter.Contract.Whitelist(&_L1GatewayRouter.CallOpts)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "finalizeInboundTransfer", arg0, arg1, arg2, arg3, arg4)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterSession) FinalizeInboundTransfer(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeInboundTransfer(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address , address , address , uint256 , bytes ) payable returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) FinalizeInboundTransfer(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.FinalizeInboundTransfer(&_L1GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _defaultGateway, address , address _counterpartGateway, address _inbox) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _defaultGateway common.Address, arg2 common.Address, _counterpartGateway common.Address, _inbox common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "initialize", _owner, _defaultGateway, arg2, _counterpartGateway, _inbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _defaultGateway, address , address _counterpartGateway, address _inbox) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) Initialize(_owner common.Address, _defaultGateway common.Address, arg2 common.Address, _counterpartGateway common.Address, _inbox common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.Initialize(&_L1GatewayRouter.TransactOpts, _owner, _defaultGateway, arg2, _counterpartGateway, _inbox)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address _owner, address _defaultGateway, address , address _counterpartGateway, address _inbox) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) Initialize(_owner common.Address, _defaultGateway common.Address, arg2 common.Address, _counterpartGateway common.Address, _inbox common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.Initialize(&_L1GatewayRouter.TransactOpts, _owner, _defaultGateway, arg2, _counterpartGateway, _inbox)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterTransactor) OutboundTransfer(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "outboundTransfer", _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterSession) OutboundTransfer(_token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.OutboundTransfer(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) OutboundTransfer(_token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.OutboundTransfer(&_L1GatewayRouter.TransactOpts, _token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterTransactor) OutboundTransferCustomRefund(opts *bind.TransactOpts, _token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "outboundTransferCustomRefund", _token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterSession) OutboundTransferCustomRefund(_token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.OutboundTransferCustomRefund(&_L1GatewayRouter.TransactOpts, _token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) OutboundTransferCustomRefund(_token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.OutboundTransferCustomRefund(&_L1GatewayRouter.TransactOpts, _token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1GatewayRouter *L1GatewayRouterSession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.PostUpgradeInit(&_L1GatewayRouter.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.PostUpgradeInit(&_L1GatewayRouter.TransactOpts)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0x5625a952.
//
// Solidity: function setDefaultGateway(address newL1DefaultGateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetDefaultGateway(opts *bind.TransactOpts, newL1DefaultGateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setDefaultGateway", newL1DefaultGateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0x5625a952.
//
// Solidity: function setDefaultGateway(address newL1DefaultGateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterSession) SetDefaultGateway(newL1DefaultGateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetDefaultGateway(&_L1GatewayRouter.TransactOpts, newL1DefaultGateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetDefaultGateway is a paid mutator transaction binding the contract method 0x5625a952.
//
// Solidity: function setDefaultGateway(address newL1DefaultGateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetDefaultGateway(newL1DefaultGateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetDefaultGateway(&_L1GatewayRouter.TransactOpts, newL1DefaultGateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateway is a paid mutator transaction binding the contract method 0x2d67b72d.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetGateway(opts *bind.TransactOpts, _gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setGateway", _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// SetGateway is a paid mutator transaction binding the contract method 0x2d67b72d.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterSession) SetGateway(_gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateway(&_L1GatewayRouter.TransactOpts, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// SetGateway is a paid mutator transaction binding the contract method 0x2d67b72d.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetGateway(_gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateway(&_L1GatewayRouter.TransactOpts, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// SetGateway0 is a paid mutator transaction binding the contract method 0xdd614569.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetGateway0(opts *bind.TransactOpts, _gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setGateway0", _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateway0 is a paid mutator transaction binding the contract method 0xdd614569.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterSession) SetGateway0(_gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateway0(&_L1GatewayRouter.TransactOpts, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateway0 is a paid mutator transaction binding the contract method 0xdd614569.
//
// Solidity: function setGateway(address _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetGateway0(_gateway common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateway0(&_L1GatewayRouter.TransactOpts, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateways is a paid mutator transaction binding the contract method 0x658b53f4.
//
// Solidity: function setGateways(address[] _token, address[] _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetGateways(opts *bind.TransactOpts, _token []common.Address, _gateway []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setGateways", _token, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateways is a paid mutator transaction binding the contract method 0x658b53f4.
//
// Solidity: function setGateways(address[] _token, address[] _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterSession) SetGateways(_token []common.Address, _gateway []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateways(&_L1GatewayRouter.TransactOpts, _token, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetGateways is a paid mutator transaction binding the contract method 0x658b53f4.
//
// Solidity: function setGateways(address[] _token, address[] _gateway, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetGateways(_token []common.Address, _gateway []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetGateways(&_L1GatewayRouter.TransactOpts, _token, _gateway, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetOwner(&_L1GatewayRouter.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.SetOwner(&_L1GatewayRouter.TransactOpts, newOwner)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactor) UpdateWhitelistSource(opts *bind.TransactOpts, newSource common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.contract.Transact(opts, "updateWhitelistSource", newSource)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_L1GatewayRouter *L1GatewayRouterSession) UpdateWhitelistSource(newSource common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.UpdateWhitelistSource(&_L1GatewayRouter.TransactOpts, newSource)
}

// UpdateWhitelistSource is a paid mutator transaction binding the contract method 0x47466f98.
//
// Solidity: function updateWhitelistSource(address newSource) returns()
func (_L1GatewayRouter *L1GatewayRouterTransactorSession) UpdateWhitelistSource(newSource common.Address) (*types.Transaction, error) {
	return _L1GatewayRouter.Contract.UpdateWhitelistSource(&_L1GatewayRouter.TransactOpts, newSource)
}

// L1GatewayRouterDefaultGatewayUpdatedIterator is returned from FilterDefaultGatewayUpdated and is used to iterate over the raw logs and unpacked data for DefaultGatewayUpdated events raised by the L1GatewayRouter contract.
type L1GatewayRouterDefaultGatewayUpdatedIterator struct {
	Event *L1GatewayRouterDefaultGatewayUpdated // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterDefaultGatewayUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterDefaultGatewayUpdated)
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
		it.Event = new(L1GatewayRouterDefaultGatewayUpdated)
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
func (it *L1GatewayRouterDefaultGatewayUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterDefaultGatewayUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterDefaultGatewayUpdated represents a DefaultGatewayUpdated event raised by the L1GatewayRouter contract.
type L1GatewayRouterDefaultGatewayUpdated struct {
	NewDefaultGateway common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDefaultGatewayUpdated is a free log retrieval operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterDefaultGatewayUpdated(opts *bind.FilterOpts) (*L1GatewayRouterDefaultGatewayUpdatedIterator, error) {

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "DefaultGatewayUpdated")
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterDefaultGatewayUpdatedIterator{contract: _L1GatewayRouter.contract, event: "DefaultGatewayUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultGatewayUpdated is a free log subscription operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchDefaultGatewayUpdated(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterDefaultGatewayUpdated) (event.Subscription, error) {

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "DefaultGatewayUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterDefaultGatewayUpdated)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "DefaultGatewayUpdated", log); err != nil {
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

// ParseDefaultGatewayUpdated is a log parse operation binding the contract event 0x3a8f8eb961383a94d41d193e16a3af73eaddfd5764a4c640257323a1603ac331.
//
// Solidity: event DefaultGatewayUpdated(address newDefaultGateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseDefaultGatewayUpdated(log types.Log) (*L1GatewayRouterDefaultGatewayUpdated, error) {
	event := new(L1GatewayRouterDefaultGatewayUpdated)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "DefaultGatewayUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterGatewaySetIterator is returned from FilterGatewaySet and is used to iterate over the raw logs and unpacked data for GatewaySet events raised by the L1GatewayRouter contract.
type L1GatewayRouterGatewaySetIterator struct {
	Event *L1GatewayRouterGatewaySet // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterGatewaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterGatewaySet)
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
		it.Event = new(L1GatewayRouterGatewaySet)
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
func (it *L1GatewayRouterGatewaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterGatewaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterGatewaySet represents a GatewaySet event raised by the L1GatewayRouter contract.
type L1GatewayRouterGatewaySet struct {
	L1Token common.Address
	Gateway common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGatewaySet is a free log retrieval operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterGatewaySet(opts *bind.FilterOpts, l1Token []common.Address, gateway []common.Address) (*L1GatewayRouterGatewaySetIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "GatewaySet", l1TokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterGatewaySetIterator{contract: _L1GatewayRouter.contract, event: "GatewaySet", logs: logs, sub: sub}, nil
}

// WatchGatewaySet is a free log subscription operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchGatewaySet(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterGatewaySet, l1Token []common.Address, gateway []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var gatewayRule []interface{}
	for _, gatewayItem := range gateway {
		gatewayRule = append(gatewayRule, gatewayItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "GatewaySet", l1TokenRule, gatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterGatewaySet)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "GatewaySet", log); err != nil {
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

// ParseGatewaySet is a log parse operation binding the contract event 0x812ca95fe4492a9e2d1f2723c2c40c03a60a27b059581ae20ac4e4d73bfba354.
//
// Solidity: event GatewaySet(address indexed l1Token, address indexed gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseGatewaySet(log types.Log) (*L1GatewayRouterGatewaySet, error) {
	event := new(L1GatewayRouterGatewaySet)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "GatewaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterTransferRoutedIterator is returned from FilterTransferRouted and is used to iterate over the raw logs and unpacked data for TransferRouted events raised by the L1GatewayRouter contract.
type L1GatewayRouterTransferRoutedIterator struct {
	Event *L1GatewayRouterTransferRouted // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterTransferRoutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterTransferRouted)
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
		it.Event = new(L1GatewayRouterTransferRouted)
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
func (it *L1GatewayRouterTransferRoutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterTransferRoutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterTransferRouted represents a TransferRouted event raised by the L1GatewayRouter contract.
type L1GatewayRouterTransferRouted struct {
	Token    common.Address
	UserFrom common.Address
	UserTo   common.Address
	Gateway  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferRouted is a free log retrieval operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterTransferRouted(opts *bind.FilterOpts, token []common.Address, _userFrom []common.Address, _userTo []common.Address) (*L1GatewayRouterTransferRoutedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _userFromRule []interface{}
	for _, _userFromItem := range _userFrom {
		_userFromRule = append(_userFromRule, _userFromItem)
	}
	var _userToRule []interface{}
	for _, _userToItem := range _userTo {
		_userToRule = append(_userToRule, _userToItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "TransferRouted", tokenRule, _userFromRule, _userToRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterTransferRoutedIterator{contract: _L1GatewayRouter.contract, event: "TransferRouted", logs: logs, sub: sub}, nil
}

// WatchTransferRouted is a free log subscription operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchTransferRouted(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterTransferRouted, token []common.Address, _userFrom []common.Address, _userTo []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var _userFromRule []interface{}
	for _, _userFromItem := range _userFrom {
		_userFromRule = append(_userFromRule, _userFromItem)
	}
	var _userToRule []interface{}
	for _, _userToItem := range _userTo {
		_userToRule = append(_userToRule, _userToItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "TransferRouted", tokenRule, _userFromRule, _userToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterTransferRouted)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "TransferRouted", log); err != nil {
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

// ParseTransferRouted is a log parse operation binding the contract event 0x85291dff2161a93c2f12c819d31889c96c63042116f5bc5a205aa701c2c429f5.
//
// Solidity: event TransferRouted(address indexed token, address indexed _userFrom, address indexed _userTo, address gateway)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseTransferRouted(log types.Log) (*L1GatewayRouterTransferRouted, error) {
	event := new(L1GatewayRouterTransferRouted)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "TransferRouted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterTxToL2Iterator is returned from FilterTxToL2 and is used to iterate over the raw logs and unpacked data for TxToL2 events raised by the L1GatewayRouter contract.
type L1GatewayRouterTxToL2Iterator struct {
	Event *L1GatewayRouterTxToL2 // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterTxToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterTxToL2)
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
		it.Event = new(L1GatewayRouterTxToL2)
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
func (it *L1GatewayRouterTxToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterTxToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterTxToL2 represents a TxToL2 event raised by the L1GatewayRouter contract.
type L1GatewayRouterTxToL2 struct {
	From   common.Address
	To     common.Address
	SeqNum *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxToL2 is a free log retrieval operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterTxToL2(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (*L1GatewayRouterTxToL2Iterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _seqNumRule []interface{}
	for _, _seqNumItem := range _seqNum {
		_seqNumRule = append(_seqNumRule, _seqNumItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterTxToL2Iterator{contract: _L1GatewayRouter.contract, event: "TxToL2", logs: logs, sub: sub}, nil
}

// WatchTxToL2 is a free log subscription operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchTxToL2(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterTxToL2, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _seqNumRule []interface{}
	for _, _seqNumItem := range _seqNum {
		_seqNumRule = append(_seqNumRule, _seqNumItem)
	}

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterTxToL2)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "TxToL2", log); err != nil {
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

// ParseTxToL2 is a log parse operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseTxToL2(log types.Log) (*L1GatewayRouterTxToL2, error) {
	event := new(L1GatewayRouterTxToL2)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "TxToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1GatewayRouterWhitelistSourceUpdatedIterator is returned from FilterWhitelistSourceUpdated and is used to iterate over the raw logs and unpacked data for WhitelistSourceUpdated events raised by the L1GatewayRouter contract.
type L1GatewayRouterWhitelistSourceUpdatedIterator struct {
	Event *L1GatewayRouterWhitelistSourceUpdated // Event containing the contract specifics and raw log

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
func (it *L1GatewayRouterWhitelistSourceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1GatewayRouterWhitelistSourceUpdated)
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
		it.Event = new(L1GatewayRouterWhitelistSourceUpdated)
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
func (it *L1GatewayRouterWhitelistSourceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1GatewayRouterWhitelistSourceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1GatewayRouterWhitelistSourceUpdated represents a WhitelistSourceUpdated event raised by the L1GatewayRouter contract.
type L1GatewayRouterWhitelistSourceUpdated struct {
	NewSource common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWhitelistSourceUpdated is a free log retrieval operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_L1GatewayRouter *L1GatewayRouterFilterer) FilterWhitelistSourceUpdated(opts *bind.FilterOpts) (*L1GatewayRouterWhitelistSourceUpdatedIterator, error) {

	logs, sub, err := _L1GatewayRouter.contract.FilterLogs(opts, "WhitelistSourceUpdated")
	if err != nil {
		return nil, err
	}
	return &L1GatewayRouterWhitelistSourceUpdatedIterator{contract: _L1GatewayRouter.contract, event: "WhitelistSourceUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistSourceUpdated is a free log subscription operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_L1GatewayRouter *L1GatewayRouterFilterer) WatchWhitelistSourceUpdated(opts *bind.WatchOpts, sink chan<- *L1GatewayRouterWhitelistSourceUpdated) (event.Subscription, error) {

	logs, sub, err := _L1GatewayRouter.contract.WatchLogs(opts, "WhitelistSourceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1GatewayRouterWhitelistSourceUpdated)
				if err := _L1GatewayRouter.contract.UnpackLog(event, "WhitelistSourceUpdated", log); err != nil {
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

// ParseWhitelistSourceUpdated is a log parse operation binding the contract event 0x37389c47920d5cc3229678a0205d0455002c07541a4139ebdce91ac227465777.
//
// Solidity: event WhitelistSourceUpdated(address newSource)
func (_L1GatewayRouter *L1GatewayRouterFilterer) ParseWhitelistSourceUpdated(log types.Log) (*L1GatewayRouterWhitelistSourceUpdated, error) {
	event := new(L1GatewayRouterWhitelistSourceUpdated)
	if err := _L1GatewayRouter.contract.UnpackLog(event, "WhitelistSourceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
