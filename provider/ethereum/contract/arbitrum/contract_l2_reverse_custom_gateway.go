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

// L2ReverseCustomGatewayMetaData contains all meta data concerning the L2ReverseCustomGateway contract.
var L2ReverseCustomGatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"}],\"name\":\"TokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TxToL1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_l2ToL1Id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalInitiated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1ERC20\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exitNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"outboundCalldata\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1ToL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"l1Address\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"l2Address\",\"type\":\"address[]\"}],\"name\":\"registerTokenFromL1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L2ReverseCustomGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L2ReverseCustomGatewayMetaData.ABI instead.
var L2ReverseCustomGatewayABI = L2ReverseCustomGatewayMetaData.ABI

// L2ReverseCustomGateway is an auto generated Go binding around an Ethereum contract.
type L2ReverseCustomGateway struct {
	L2ReverseCustomGatewayCaller     // Read-only binding to the contract
	L2ReverseCustomGatewayTransactor // Write-only binding to the contract
	L2ReverseCustomGatewayFilterer   // Log filterer for contract events
}

// L2ReverseCustomGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ReverseCustomGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ReverseCustomGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ReverseCustomGatewaySession struct {
	Contract     *L2ReverseCustomGateway // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2ReverseCustomGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ReverseCustomGatewayCallerSession struct {
	Contract *L2ReverseCustomGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// L2ReverseCustomGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ReverseCustomGatewayTransactorSession struct {
	Contract     *L2ReverseCustomGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// L2ReverseCustomGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ReverseCustomGatewayRaw struct {
	Contract *L2ReverseCustomGateway // Generic contract binding to access the raw methods on
}

// L2ReverseCustomGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayCallerRaw struct {
	Contract *L2ReverseCustomGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L2ReverseCustomGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ReverseCustomGatewayTransactorRaw struct {
	Contract *L2ReverseCustomGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ReverseCustomGateway creates a new instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGateway(address common.Address, backend bind.ContractBackend) (*L2ReverseCustomGateway, error) {
	contract, err := bindL2ReverseCustomGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGateway{L2ReverseCustomGatewayCaller: L2ReverseCustomGatewayCaller{contract: contract}, L2ReverseCustomGatewayTransactor: L2ReverseCustomGatewayTransactor{contract: contract}, L2ReverseCustomGatewayFilterer: L2ReverseCustomGatewayFilterer{contract: contract}}, nil
}

// NewL2ReverseCustomGatewayCaller creates a new read-only instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayCaller(address common.Address, caller bind.ContractCaller) (*L2ReverseCustomGatewayCaller, error) {
	contract, err := bindL2ReverseCustomGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayCaller{contract: contract}, nil
}

// NewL2ReverseCustomGatewayTransactor creates a new write-only instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L2ReverseCustomGatewayTransactor, error) {
	contract, err := bindL2ReverseCustomGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayTransactor{contract: contract}, nil
}

// NewL2ReverseCustomGatewayFilterer creates a new log filterer instance of L2ReverseCustomGateway, bound to a specific deployed contract.
func NewL2ReverseCustomGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L2ReverseCustomGatewayFilterer, error) {
	contract, err := bindL2ReverseCustomGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayFilterer{contract: contract}, nil
}

// bindL2ReverseCustomGateway binds a generic wrapper to an already deployed contract.
func bindL2ReverseCustomGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2ReverseCustomGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.L2ReverseCustomGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ReverseCustomGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1ERC20 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "calculateL2TokenAddress", l1ERC20)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.CalculateL2TokenAddress(&_L2ReverseCustomGateway.CallOpts, l1ERC20)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.CalculateL2TokenAddress(&_L2ReverseCustomGateway.CallOpts, l1ERC20)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) CounterpartGateway() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.CounterpartGateway(&_L2ReverseCustomGateway.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) CounterpartGateway() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.CounterpartGateway(&_L2ReverseCustomGateway.CallOpts)
}

// ExitNum is a free data retrieval call binding the contract method 0x015234ab.
//
// Solidity: function exitNum() view returns(uint256)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) ExitNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "exitNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitNum is a free data retrieval call binding the contract method 0x015234ab.
//
// Solidity: function exitNum() view returns(uint256)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) ExitNum() (*big.Int, error) {
	return _L2ReverseCustomGateway.Contract.ExitNum(&_L2ReverseCustomGateway.CallOpts)
}

// ExitNum is a free data retrieval call binding the contract method 0x015234ab.
//
// Solidity: function exitNum() view returns(uint256)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) ExitNum() (*big.Int, error) {
	return _L2ReverseCustomGateway.Contract.ExitNum(&_L2ReverseCustomGateway.CallOpts)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) GetOutboundCalldata(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "getOutboundCalldata", _token, _from, _to, _amount, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L2ReverseCustomGateway.Contract.GetOutboundCalldata(&_L2ReverseCustomGateway.CallOpts, _token, _from, _to, _amount, _data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L2ReverseCustomGateway.Contract.GetOutboundCalldata(&_L2ReverseCustomGateway.CallOpts, _token, _from, _to, _amount, _data)
}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) L1ToL2Token(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "l1ToL2Token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) L1ToL2Token(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.L1ToL2Token(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) L1ToL2Token(arg0 common.Address) (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.L1ToL2Token(&_L2ReverseCustomGateway.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2ReverseCustomGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Router() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Router(&_L2ReverseCustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayCallerSession) Router() (common.Address, error) {
	return _L2ReverseCustomGateway.Contract.Router(&_L2ReverseCustomGateway.CallOpts)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "finalizeInboundTransfer", _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.FinalizeInboundTransfer(&_L2ReverseCustomGateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.FinalizeInboundTransfer(&_L2ReverseCustomGateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l1Counterpart, address _router) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) Initialize(opts *bind.TransactOpts, _l1Counterpart common.Address, _router common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "initialize", _l1Counterpart, _router)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l1Counterpart, address _router) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) Initialize(_l1Counterpart common.Address, _router common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.Initialize(&_L2ReverseCustomGateway.TransactOpts, _l1Counterpart, _router)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _l1Counterpart, address _router) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) Initialize(_l1Counterpart common.Address, _router common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.Initialize(&_L2ReverseCustomGateway.TransactOpts, _l1Counterpart, _router)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) OutboundTransfer(opts *bind.TransactOpts, _l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "outboundTransfer", _l1Token, _to, _amount, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.OutboundTransfer(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _to, _amount, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0x7b3a3c8b.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, bytes _data) payable returns(bytes)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.OutboundTransfer(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _to, _amount, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 , uint256 , bytes _data) payable returns(bytes res)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) OutboundTransfer0(opts *bind.TransactOpts, _l1Token common.Address, _to common.Address, _amount *big.Int, arg3 *big.Int, arg4 *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "outboundTransfer0", _l1Token, _to, _amount, arg3, arg4, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 , uint256 , bytes _data) payable returns(bytes res)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) OutboundTransfer0(_l1Token common.Address, _to common.Address, _amount *big.Int, arg3 *big.Int, arg4 *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.OutboundTransfer0(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _to, _amount, arg3, arg4, _data)
}

// OutboundTransfer0 is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 , uint256 , bytes _data) payable returns(bytes res)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) OutboundTransfer0(_l1Token common.Address, _to common.Address, _amount *big.Int, arg3 *big.Int, arg4 *big.Int, _data []byte) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.OutboundTransfer0(&_L2ReverseCustomGateway.TransactOpts, _l1Token, _to, _amount, arg3, arg4, _data)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) PostUpgradeInit() (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.PostUpgradeInit(&_L2ReverseCustomGateway.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.PostUpgradeInit(&_L2ReverseCustomGateway.TransactOpts)
}

// RegisterTokenFromL1 is a paid mutator transaction binding the contract method 0xd4f5532f.
//
// Solidity: function registerTokenFromL1(address[] l1Address, address[] l2Address) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactor) RegisterTokenFromL1(opts *bind.TransactOpts, l1Address []common.Address, l2Address []common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.contract.Transact(opts, "registerTokenFromL1", l1Address, l2Address)
}

// RegisterTokenFromL1 is a paid mutator transaction binding the contract method 0xd4f5532f.
//
// Solidity: function registerTokenFromL1(address[] l1Address, address[] l2Address) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewaySession) RegisterTokenFromL1(l1Address []common.Address, l2Address []common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.RegisterTokenFromL1(&_L2ReverseCustomGateway.TransactOpts, l1Address, l2Address)
}

// RegisterTokenFromL1 is a paid mutator transaction binding the contract method 0xd4f5532f.
//
// Solidity: function registerTokenFromL1(address[] l1Address, address[] l2Address) returns()
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayTransactorSession) RegisterTokenFromL1(l1Address []common.Address, l2Address []common.Address) (*types.Transaction, error) {
	return _L2ReverseCustomGateway.Contract.RegisterTokenFromL1(&_L2ReverseCustomGateway.TransactOpts, l1Address, l2Address)
}

// L2ReverseCustomGatewayDepositFinalizedIterator is returned from FilterDepositFinalized and is used to iterate over the raw logs and unpacked data for DepositFinalized events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayDepositFinalizedIterator struct {
	Event *L2ReverseCustomGatewayDepositFinalized // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayDepositFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayDepositFinalized)
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
		it.Event = new(L2ReverseCustomGatewayDepositFinalized)
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
func (it *L2ReverseCustomGatewayDepositFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayDepositFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayDepositFinalized represents a DepositFinalized event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayDepositFinalized struct {
	L1Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositFinalized is a free log retrieval operation binding the contract event 0xc7f2e9c55c40a50fbc217dfc70cd39a222940dfa62145aa0ca49eb9535d4fcb2.
//
// Solidity: event DepositFinalized(address indexed l1Token, address indexed _from, address indexed _to, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterDepositFinalized(opts *bind.FilterOpts, l1Token []common.Address, _from []common.Address, _to []common.Address) (*L2ReverseCustomGatewayDepositFinalizedIterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "DepositFinalized", l1TokenRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayDepositFinalizedIterator{contract: _L2ReverseCustomGateway.contract, event: "DepositFinalized", logs: logs, sub: sub}, nil
}

// WatchDepositFinalized is a free log subscription operation binding the contract event 0xc7f2e9c55c40a50fbc217dfc70cd39a222940dfa62145aa0ca49eb9535d4fcb2.
//
// Solidity: event DepositFinalized(address indexed l1Token, address indexed _from, address indexed _to, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchDepositFinalized(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayDepositFinalized, l1Token []common.Address, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "DepositFinalized", l1TokenRule, _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayDepositFinalized)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
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

// ParseDepositFinalized is a log parse operation binding the contract event 0xc7f2e9c55c40a50fbc217dfc70cd39a222940dfa62145aa0ca49eb9535d4fcb2.
//
// Solidity: event DepositFinalized(address indexed l1Token, address indexed _from, address indexed _to, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseDepositFinalized(log types.Log) (*L2ReverseCustomGatewayDepositFinalized, error) {
	event := new(L2ReverseCustomGatewayDepositFinalized)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "DepositFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayTokenSetIterator is returned from FilterTokenSet and is used to iterate over the raw logs and unpacked data for TokenSet events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayTokenSetIterator struct {
	Event *L2ReverseCustomGatewayTokenSet // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayTokenSet)
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
		it.Event = new(L2ReverseCustomGatewayTokenSet)
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
func (it *L2ReverseCustomGatewayTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayTokenSet represents a TokenSet event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayTokenSet struct {
	L1Address common.Address
	L2Address common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSet is a free log retrieval operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterTokenSet(opts *bind.FilterOpts, l1Address []common.Address, l2Address []common.Address) (*L2ReverseCustomGatewayTokenSetIterator, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "TokenSet", l1AddressRule, l2AddressRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayTokenSetIterator{contract: _L2ReverseCustomGateway.contract, event: "TokenSet", logs: logs, sub: sub}, nil
}

// WatchTokenSet is a free log subscription operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchTokenSet(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayTokenSet, l1Address []common.Address, l2Address []common.Address) (event.Subscription, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "TokenSet", l1AddressRule, l2AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayTokenSet)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "TokenSet", log); err != nil {
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

// ParseTokenSet is a log parse operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseTokenSet(log types.Log) (*L2ReverseCustomGatewayTokenSet, error) {
	event := new(L2ReverseCustomGatewayTokenSet)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "TokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayTxToL1Iterator is returned from FilterTxToL1 and is used to iterate over the raw logs and unpacked data for TxToL1 events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayTxToL1Iterator struct {
	Event *L2ReverseCustomGatewayTxToL1 // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayTxToL1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayTxToL1)
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
		it.Event = new(L2ReverseCustomGatewayTxToL1)
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
func (it *L2ReverseCustomGatewayTxToL1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayTxToL1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayTxToL1 represents a TxToL1 event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayTxToL1 struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTxToL1 is a free log retrieval operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterTxToL1(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _id []*big.Int) (*L2ReverseCustomGatewayTxToL1Iterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "TxToL1", _fromRule, _toRule, _idRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayTxToL1Iterator{contract: _L2ReverseCustomGateway.contract, event: "TxToL1", logs: logs, sub: sub}, nil
}

// WatchTxToL1 is a free log subscription operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchTxToL1(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayTxToL1, _from []common.Address, _to []common.Address, _id []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "TxToL1", _fromRule, _toRule, _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayTxToL1)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "TxToL1", log); err != nil {
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

// ParseTxToL1 is a log parse operation binding the contract event 0x2b986d32a0536b7e19baa48ab949fec7b903b7fad7730820b20632d100cc3a68.
//
// Solidity: event TxToL1(address indexed _from, address indexed _to, uint256 indexed _id, bytes _data)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseTxToL1(log types.Log) (*L2ReverseCustomGatewayTxToL1, error) {
	event := new(L2ReverseCustomGatewayTxToL1)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "TxToL1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ReverseCustomGatewayWithdrawalInitiatedIterator is returned from FilterWithdrawalInitiated and is used to iterate over the raw logs and unpacked data for WithdrawalInitiated events raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayWithdrawalInitiatedIterator struct {
	Event *L2ReverseCustomGatewayWithdrawalInitiated // Event containing the contract specifics and raw log

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
func (it *L2ReverseCustomGatewayWithdrawalInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ReverseCustomGatewayWithdrawalInitiated)
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
		it.Event = new(L2ReverseCustomGatewayWithdrawalInitiated)
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
func (it *L2ReverseCustomGatewayWithdrawalInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ReverseCustomGatewayWithdrawalInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ReverseCustomGatewayWithdrawalInitiated represents a WithdrawalInitiated event raised by the L2ReverseCustomGateway contract.
type L2ReverseCustomGatewayWithdrawalInitiated struct {
	L1Token  common.Address
	From     common.Address
	To       common.Address
	L2ToL1Id *big.Int
	ExitNum  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalInitiated is a free log retrieval operation binding the contract event 0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73.
//
// Solidity: event WithdrawalInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _l2ToL1Id, uint256 _exitNum, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) FilterWithdrawalInitiated(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _l2ToL1Id []*big.Int) (*L2ReverseCustomGatewayWithdrawalInitiatedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _l2ToL1IdRule []interface{}
	for _, _l2ToL1IdItem := range _l2ToL1Id {
		_l2ToL1IdRule = append(_l2ToL1IdRule, _l2ToL1IdItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.FilterLogs(opts, "WithdrawalInitiated", _fromRule, _toRule, _l2ToL1IdRule)
	if err != nil {
		return nil, err
	}
	return &L2ReverseCustomGatewayWithdrawalInitiatedIterator{contract: _L2ReverseCustomGateway.contract, event: "WithdrawalInitiated", logs: logs, sub: sub}, nil
}

// WatchWithdrawalInitiated is a free log subscription operation binding the contract event 0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73.
//
// Solidity: event WithdrawalInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _l2ToL1Id, uint256 _exitNum, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) WatchWithdrawalInitiated(opts *bind.WatchOpts, sink chan<- *L2ReverseCustomGatewayWithdrawalInitiated, _from []common.Address, _to []common.Address, _l2ToL1Id []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _l2ToL1IdRule []interface{}
	for _, _l2ToL1IdItem := range _l2ToL1Id {
		_l2ToL1IdRule = append(_l2ToL1IdRule, _l2ToL1IdItem)
	}

	logs, sub, err := _L2ReverseCustomGateway.contract.WatchLogs(opts, "WithdrawalInitiated", _fromRule, _toRule, _l2ToL1IdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ReverseCustomGatewayWithdrawalInitiated)
				if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
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

// ParseWithdrawalInitiated is a log parse operation binding the contract event 0x3073a74ecb728d10be779fe19a74a1428e20468f5b4d167bf9c73d9067847d73.
//
// Solidity: event WithdrawalInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _l2ToL1Id, uint256 _exitNum, uint256 _amount)
func (_L2ReverseCustomGateway *L2ReverseCustomGatewayFilterer) ParseWithdrawalInitiated(log types.Log) (*L2ReverseCustomGatewayWithdrawalInitiated, error) {
	event := new(L2ReverseCustomGatewayWithdrawalInitiated)
	if err := _L2ReverseCustomGateway.contract.UnpackLog(event, "WithdrawalInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
