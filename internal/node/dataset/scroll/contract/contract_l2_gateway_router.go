// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// L2GatewayRouterMetaData contains all meta data concerning the L2GatewayRouter contract.
var L2GatewayRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"FinalizeDepositETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldDefaultERC20Gateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"SetDefaultERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGateway\",\"type\":\"address\"}],\"name\":\"SetERC20Gateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldETHGateway\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newEthGateway\",\"type\":\"address\"}],\"name\":\"SetETHGateway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawERC20\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"WithdrawETH\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"finalizeDepositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"getERC20Gateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Address\",\"type\":\"address\"}],\"name\":\"getL1ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getL2ERC20Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ethGateway\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_defaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDefaultERC20Gateway\",\"type\":\"address\"}],\"name\":\"setDefaultERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_gateways\",\"type\":\"address[]\"}],\"name\":\"setERC20Gateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newEthGateway\",\"type\":\"address\"}],\"name\":\"setETHGateway\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawERC20AndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"}],\"name\":\"withdrawETHAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// L2GatewayRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use L2GatewayRouterMetaData.ABI instead.
var L2GatewayRouterABI = L2GatewayRouterMetaData.ABI

// L2GatewayRouter is an auto generated Go binding around an Ethereum contract.
type L2GatewayRouter struct {
	L2GatewayRouterCaller     // Read-only binding to the contract
	L2GatewayRouterTransactor // Write-only binding to the contract
	L2GatewayRouterFilterer   // Log filterer for contract events
}

// L2GatewayRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2GatewayRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2GatewayRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2GatewayRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2GatewayRouterSession struct {
	Contract     *L2GatewayRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2GatewayRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2GatewayRouterCallerSession struct {
	Contract *L2GatewayRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L2GatewayRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2GatewayRouterTransactorSession struct {
	Contract     *L2GatewayRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L2GatewayRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2GatewayRouterRaw struct {
	Contract *L2GatewayRouter // Generic contract binding to access the raw methods on
}

// L2GatewayRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2GatewayRouterCallerRaw struct {
	Contract *L2GatewayRouterCaller // Generic read-only contract binding to access the raw methods on
}

// L2GatewayRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2GatewayRouterTransactorRaw struct {
	Contract *L2GatewayRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2GatewayRouter creates a new instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouter(address common.Address, backend bind.ContractBackend) (*L2GatewayRouter, error) {
	contract, err := bindL2GatewayRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouter{L2GatewayRouterCaller: L2GatewayRouterCaller{contract: contract}, L2GatewayRouterTransactor: L2GatewayRouterTransactor{contract: contract}, L2GatewayRouterFilterer: L2GatewayRouterFilterer{contract: contract}}, nil
}

// NewL2GatewayRouterCaller creates a new read-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterCaller(address common.Address, caller bind.ContractCaller) (*L2GatewayRouterCaller, error) {
	contract, err := bindL2GatewayRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterCaller{contract: contract}, nil
}

// NewL2GatewayRouterTransactor creates a new write-only instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*L2GatewayRouterTransactor, error) {
	contract, err := bindL2GatewayRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterTransactor{contract: contract}, nil
}

// NewL2GatewayRouterFilterer creates a new log filterer instance of L2GatewayRouter, bound to a specific deployed contract.
func NewL2GatewayRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*L2GatewayRouterFilterer, error) {
	contract, err := bindL2GatewayRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFilterer{contract: contract}, nil
}

// bindL2GatewayRouter binds a generic wrapper to an already deployed contract.
func bindL2GatewayRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L2GatewayRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.L2GatewayRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.L2GatewayRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2GatewayRouter *L2GatewayRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2GatewayRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2GatewayRouter *L2GatewayRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.contract.Transact(opts, method, params...)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) ERC20Gateway(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ERC20Gateway", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// ERC20Gateway is a free data retrieval call binding the contract method 0x705b05b8.
//
// Solidity: function ERC20Gateway(address ) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) ERC20Gateway(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.ERC20Gateway(&_L2GatewayRouter.CallOpts, arg0)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) DefaultERC20Gateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "defaultERC20Gateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// DefaultERC20Gateway is a free data retrieval call binding the contract method 0xce8c3e06.
//
// Solidity: function defaultERC20Gateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) DefaultERC20Gateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.DefaultERC20Gateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) EthGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "ethGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// EthGateway is a free data retrieval call binding the contract method 0x8c00ce73.
//
// Solidity: function ethGateway() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) EthGateway() (common.Address, error) {
	return _L2GatewayRouter.Contract.EthGateway(&_L2GatewayRouter.CallOpts)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetERC20Gateway(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getERC20Gateway", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetERC20Gateway is a free data retrieval call binding the contract method 0x43c66741.
//
// Solidity: function getERC20Gateway(address _token) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetERC20Gateway(_token common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetERC20Gateway(&_L2GatewayRouter.CallOpts, _token)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL1ERC20Address(opts *bind.CallOpts, _l2Address common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL1ERC20Address", _l2Address)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL1ERC20Address is a free data retrieval call binding the contract method 0x54bbd59c.
//
// Solidity: function getL1ERC20Address(address _l2Address) view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL1ERC20Address(_l2Address common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL1ERC20Address(&_L2GatewayRouter.CallOpts, _l2Address)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) GetL2ERC20Address(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "getL2ERC20Address", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// GetL2ERC20Address is a free data retrieval call binding the contract method 0xc676ad29.
//
// Solidity: function getL2ERC20Address(address ) pure returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) GetL2ERC20Address(arg0 common.Address) (common.Address, error) {
	return _L2GatewayRouter.Contract.GetL2ERC20Address(&_L2GatewayRouter.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2GatewayRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2GatewayRouter *L2GatewayRouterCallerSession) Owner() (common.Address, error) {
	return _L2GatewayRouter.Contract.Owner(&_L2GatewayRouter.CallOpts)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositERC20(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositERC20", arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositERC20 is a paid mutator transaction binding the contract method 0x8431f5c1.
//
// Solidity: function finalizeDepositERC20(address , address , address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositERC20(arg0 common.Address, arg1 common.Address, arg2 common.Address, arg3 common.Address, arg4 *big.Int, arg5 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositERC20(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4, arg5)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) FinalizeDepositETH(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "finalizeDepositETH", arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// FinalizeDepositETH is a paid mutator transaction binding the contract method 0x232e8748.
//
// Solidity: function finalizeDepositETH(address , address , uint256 , bytes ) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) FinalizeDepositETH(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.FinalizeDepositETH(&_L2GatewayRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) Initialize(opts *bind.TransactOpts, _ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "initialize", _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _ethGateway, address _defaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) Initialize(_ethGateway common.Address, _defaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.Initialize(&_L2GatewayRouter.TransactOpts, _ethGateway, _defaultERC20Gateway)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.RenounceOwnership(&_L2GatewayRouter.TransactOpts)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetDefaultERC20Gateway(opts *bind.TransactOpts, _newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setDefaultERC20Gateway", _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetDefaultERC20Gateway is a paid mutator transaction binding the contract method 0x5dfd5b9a.
//
// Solidity: function setDefaultERC20Gateway(address _newDefaultERC20Gateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetDefaultERC20Gateway(_newDefaultERC20Gateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetDefaultERC20Gateway(&_L2GatewayRouter.TransactOpts, _newDefaultERC20Gateway)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetERC20Gateway(opts *bind.TransactOpts, _tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setERC20Gateway", _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetERC20Gateway is a paid mutator transaction binding the contract method 0x635c8637.
//
// Solidity: function setERC20Gateway(address[] _tokens, address[] _gateways) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetERC20Gateway(_tokens []common.Address, _gateways []common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetERC20Gateway(&_L2GatewayRouter.TransactOpts, _tokens, _gateways)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) SetETHGateway(opts *bind.TransactOpts, _newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "setETHGateway", _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// SetETHGateway is a paid mutator transaction binding the contract method 0x3d1d31c7.
//
// Solidity: function setETHGateway(address _newEthGateway) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) SetETHGateway(_newEthGateway common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.SetETHGateway(&_L2GatewayRouter.TransactOpts, _newEthGateway)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.TransferOwnership(&_L2GatewayRouter.TransactOpts, newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20", _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x6c07ea43.
//
// Solidity: function withdrawERC20(address _token, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20(_token common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20(&_L2GatewayRouter.TransactOpts, _token, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC200(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC200", _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC200 is a paid mutator transaction binding the contract method 0xa93a4af9.
//
// Solidity: function withdrawERC20(address _token, address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC200(_token common.Address, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC200(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawERC20AndCall(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawERC20AndCall", _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawERC20AndCall is a paid mutator transaction binding the contract method 0x575361b6.
//
// Solidity: function withdrawERC20AndCall(address _token, address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawERC20AndCall(_token common.Address, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawERC20AndCall(&_L2GatewayRouter.TransactOpts, _token, _to, _amount, _data, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH", _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x2fcc29fa.
//
// Solidity: function withdrawETH(address _to, uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH(_to common.Address, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH(&_L2GatewayRouter.TransactOpts, _to, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETH0(opts *bind.TransactOpts, _amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETH0", _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETH0 is a paid mutator transaction binding the contract method 0xc7cdea37.
//
// Solidity: function withdrawETH(uint256 _amount, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETH0(_amount *big.Int, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETH0(&_L2GatewayRouter.TransactOpts, _amount, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactor) WithdrawETHAndCall(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.contract.Transact(opts, "withdrawETHAndCall", _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// WithdrawETHAndCall is a paid mutator transaction binding the contract method 0x6dc24183.
//
// Solidity: function withdrawETHAndCall(address _to, uint256 _amount, bytes _data, uint256 _gasLimit) payable returns()
func (_L2GatewayRouter *L2GatewayRouterTransactorSession) WithdrawETHAndCall(_to common.Address, _amount *big.Int, _data []byte, _gasLimit *big.Int) (*types.Transaction, error) {
	return _L2GatewayRouter.Contract.WithdrawETHAndCall(&_L2GatewayRouter.TransactOpts, _to, _amount, _data, _gasLimit)
}

// L2GatewayRouterFinalizeDepositERC20Iterator is returned from FilterFinalizeDepositERC20 and is used to iterate over the raw logs and unpacked data for FinalizeDepositERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20Iterator struct {
	Event *L2GatewayRouterFinalizeDepositERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
		it.Event = new(L2GatewayRouterFinalizeDepositERC20)
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
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositERC20 represents a FinalizeDepositERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositERC20 is a free log retrieval operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterFinalizeDepositERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositERC20Iterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositERC20", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositERC20 is a free log subscription operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
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

// ParseFinalizeDepositERC20 is a log parse operation binding the contract event 0x165ba69f6ab40c50cade6f65431801e5f9c7d7830b7545391920db039133ba34.
//
// Solidity: event FinalizeDepositERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositERC20(log types.Log) (*L2GatewayRouterFinalizeDepositERC20, error) {
	event := new(L2GatewayRouterFinalizeDepositERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterFinalizeDepositETHIterator is returned from FilterFinalizeDepositETH and is used to iterate over the raw logs and unpacked data for FinalizeDepositETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETHIterator struct {
	Event *L2GatewayRouterFinalizeDepositETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
		it.Event = new(L2GatewayRouterFinalizeDepositETH)
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
func (it *L2GatewayRouterFinalizeDepositETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterFinalizeDepositETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterFinalizeDepositETH represents a FinalizeDepositETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterFinalizeDepositETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFinalizeDepositETH is a free log retrieval operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterFinalizeDepositETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterFinalizeDepositETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterFinalizeDepositETHIterator{contract: _L2GatewayRouter.contract, event: "FinalizeDepositETH", logs: logs, sub: sub}, nil
}

// WatchFinalizeDepositETH is a free log subscription operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchFinalizeDepositETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterFinalizeDepositETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "FinalizeDepositETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterFinalizeDepositETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
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

// ParseFinalizeDepositETH is a log parse operation binding the contract event 0x9e86c356e14e24e26e3ce769bf8b87de38e0faa0ed0ca946fa09659aa606bd2d.
//
// Solidity: event FinalizeDepositETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseFinalizeDepositETH(log types.Log) (*L2GatewayRouterFinalizeDepositETH, error) {
	event := new(L2GatewayRouterFinalizeDepositETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "FinalizeDepositETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the L2GatewayRouter contract.
type L2GatewayRouterInitializedIterator struct {
	Event *L2GatewayRouterInitialized // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterInitialized)
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
		it.Event = new(L2GatewayRouterInitialized)
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
func (it *L2GatewayRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterInitialized represents a Initialized event raised by the L2GatewayRouter contract.
type L2GatewayRouterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*L2GatewayRouterInitializedIterator, error) {

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterInitializedIterator{contract: _L2GatewayRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterInitialized)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseInitialized(log types.Log) (*L2GatewayRouterInitialized, error) {
	event := new(L2GatewayRouterInitialized)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferredIterator struct {
	Event *L2GatewayRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterOwnershipTransferred)
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
		it.Event = new(L2GatewayRouterOwnershipTransferred)
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
func (it *L2GatewayRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterOwnershipTransferred represents a OwnershipTransferred event raised by the L2GatewayRouter contract.
type L2GatewayRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L2GatewayRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterOwnershipTransferredIterator{contract: _L2GatewayRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterOwnershipTransferred)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseOwnershipTransferred(log types.Log) (*L2GatewayRouterOwnershipTransferred, error) {
	event := new(L2GatewayRouterOwnershipTransferred)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetDefaultERC20GatewayIterator is returned from FilterSetDefaultERC20Gateway and is used to iterate over the raw logs and unpacked data for SetDefaultERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20GatewayIterator struct {
	Event *L2GatewayRouterSetDefaultERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetDefaultERC20Gateway)
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
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetDefaultERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetDefaultERC20Gateway represents a SetDefaultERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetDefaultERC20Gateway struct {
	OldDefaultERC20Gateway common.Address
	NewDefaultERC20Gateway common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetDefaultERC20Gateway is a free log retrieval operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetDefaultERC20Gateway(opts *bind.FilterOpts, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (*L2GatewayRouterSetDefaultERC20GatewayIterator, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetDefaultERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetDefaultERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetDefaultERC20Gateway is a free log subscription operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetDefaultERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetDefaultERC20Gateway, oldDefaultERC20Gateway []common.Address, newDefaultERC20Gateway []common.Address) (event.Subscription, error) {

	var oldDefaultERC20GatewayRule []interface{}
	for _, oldDefaultERC20GatewayItem := range oldDefaultERC20Gateway {
		oldDefaultERC20GatewayRule = append(oldDefaultERC20GatewayRule, oldDefaultERC20GatewayItem)
	}
	var newDefaultERC20GatewayRule []interface{}
	for _, newDefaultERC20GatewayItem := range newDefaultERC20Gateway {
		newDefaultERC20GatewayRule = append(newDefaultERC20GatewayRule, newDefaultERC20GatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetDefaultERC20Gateway", oldDefaultERC20GatewayRule, newDefaultERC20GatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetDefaultERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
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

// ParseSetDefaultERC20Gateway is a log parse operation binding the contract event 0x2904fcae71038f87b116fd2875871e153722cabddd71de1b77473de263cd74d1.
//
// Solidity: event SetDefaultERC20Gateway(address indexed oldDefaultERC20Gateway, address indexed newDefaultERC20Gateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetDefaultERC20Gateway(log types.Log) (*L2GatewayRouterSetDefaultERC20Gateway, error) {
	event := new(L2GatewayRouterSetDefaultERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetDefaultERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetERC20GatewayIterator is returned from FilterSetERC20Gateway and is used to iterate over the raw logs and unpacked data for SetERC20Gateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20GatewayIterator struct {
	Event *L2GatewayRouterSetERC20Gateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetERC20GatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetERC20Gateway)
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
		it.Event = new(L2GatewayRouterSetERC20Gateway)
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
func (it *L2GatewayRouterSetERC20GatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetERC20GatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetERC20Gateway represents a SetERC20Gateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetERC20Gateway struct {
	Token      common.Address
	OldGateway common.Address
	NewGateway common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetERC20Gateway is a free log retrieval operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetERC20Gateway(opts *bind.FilterOpts, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (*L2GatewayRouterSetERC20GatewayIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetERC20GatewayIterator{contract: _L2GatewayRouter.contract, event: "SetERC20Gateway", logs: logs, sub: sub}, nil
}

// WatchSetERC20Gateway is a free log subscription operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetERC20Gateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetERC20Gateway, token []common.Address, oldGateway []common.Address, newGateway []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var oldGatewayRule []interface{}
	for _, oldGatewayItem := range oldGateway {
		oldGatewayRule = append(oldGatewayRule, oldGatewayItem)
	}
	var newGatewayRule []interface{}
	for _, newGatewayItem := range newGateway {
		newGatewayRule = append(newGatewayRule, newGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetERC20Gateway", tokenRule, oldGatewayRule, newGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetERC20Gateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
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

// ParseSetERC20Gateway is a log parse operation binding the contract event 0x0ead4808404683f66d413d788a768219ea9785c97889221193103841a5841eaf.
//
// Solidity: event SetERC20Gateway(address indexed token, address indexed oldGateway, address indexed newGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetERC20Gateway(log types.Log) (*L2GatewayRouterSetERC20Gateway, error) {
	event := new(L2GatewayRouterSetERC20Gateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetERC20Gateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterSetETHGatewayIterator is returned from FilterSetETHGateway and is used to iterate over the raw logs and unpacked data for SetETHGateway events raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGatewayIterator struct {
	Event *L2GatewayRouterSetETHGateway // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterSetETHGatewayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterSetETHGateway)
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
		it.Event = new(L2GatewayRouterSetETHGateway)
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
func (it *L2GatewayRouterSetETHGatewayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterSetETHGatewayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterSetETHGateway represents a SetETHGateway event raised by the L2GatewayRouter contract.
type L2GatewayRouterSetETHGateway struct {
	OldETHGateway common.Address
	NewEthGateway common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetETHGateway is a free log retrieval operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterSetETHGateway(opts *bind.FilterOpts, oldETHGateway []common.Address, newEthGateway []common.Address) (*L2GatewayRouterSetETHGatewayIterator, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterSetETHGatewayIterator{contract: _L2GatewayRouter.contract, event: "SetETHGateway", logs: logs, sub: sub}, nil
}

// WatchSetETHGateway is a free log subscription operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchSetETHGateway(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterSetETHGateway, oldETHGateway []common.Address, newEthGateway []common.Address) (event.Subscription, error) {

	var oldETHGatewayRule []interface{}
	for _, oldETHGatewayItem := range oldETHGateway {
		oldETHGatewayRule = append(oldETHGatewayRule, oldETHGatewayItem)
	}
	var newEthGatewayRule []interface{}
	for _, newEthGatewayItem := range newEthGateway {
		newEthGatewayRule = append(newEthGatewayRule, newEthGatewayItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "SetETHGateway", oldETHGatewayRule, newEthGatewayRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterSetETHGateway)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
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

// ParseSetETHGateway is a log parse operation binding the contract event 0xa1bfcc6dd729ad197a1180f44d5c12bcc630943df0874b9ed53da23165621b6a.
//
// Solidity: event SetETHGateway(address indexed oldETHGateway, address indexed newEthGateway)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseSetETHGateway(log types.Log) (*L2GatewayRouterSetETHGateway, error) {
	event := new(L2GatewayRouterSetETHGateway)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "SetETHGateway", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawERC20Iterator is returned from FilterWithdrawERC20 and is used to iterate over the raw logs and unpacked data for WithdrawERC20 events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20Iterator struct {
	Event *L2GatewayRouterWithdrawERC20 // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawERC20Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawERC20)
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
		it.Event = new(L2GatewayRouterWithdrawERC20)
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
func (it *L2GatewayRouterWithdrawERC20Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawERC20Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawERC20 represents a WithdrawERC20 event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawERC20 struct {
	L1Token common.Address
	L2Token common.Address
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Data    []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawERC20 is a free log retrieval operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawERC20(opts *bind.FilterOpts, l1Token []common.Address, l2Token []common.Address, from []common.Address) (*L2GatewayRouterWithdrawERC20Iterator, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawERC20Iterator{contract: _L2GatewayRouter.contract, event: "WithdrawERC20", logs: logs, sub: sub}, nil
}

// WatchWithdrawERC20 is a free log subscription operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawERC20(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawERC20, l1Token []common.Address, l2Token []common.Address, from []common.Address) (event.Subscription, error) {

	var l1TokenRule []interface{}
	for _, l1TokenItem := range l1Token {
		l1TokenRule = append(l1TokenRule, l1TokenItem)
	}
	var l2TokenRule []interface{}
	for _, l2TokenItem := range l2Token {
		l2TokenRule = append(l2TokenRule, l2TokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawERC20", l1TokenRule, l2TokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawERC20)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
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

// ParseWithdrawERC20 is a log parse operation binding the contract event 0xd8d3a3f4ab95694bef40475997598bcf8acd3ed9617a4c1013795429414c27e8.
//
// Solidity: event WithdrawERC20(address indexed l1Token, address indexed l2Token, address indexed from, address to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawERC20(log types.Log) (*L2GatewayRouterWithdrawERC20, error) {
	event := new(L2GatewayRouterWithdrawERC20)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawERC20", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2GatewayRouterWithdrawETHIterator is returned from FilterWithdrawETH and is used to iterate over the raw logs and unpacked data for WithdrawETH events raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETHIterator struct {
	Event *L2GatewayRouterWithdrawETH // Event containing the contract specifics and raw log

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
func (it *L2GatewayRouterWithdrawETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2GatewayRouterWithdrawETH)
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
		it.Event = new(L2GatewayRouterWithdrawETH)
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
func (it *L2GatewayRouterWithdrawETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2GatewayRouterWithdrawETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2GatewayRouterWithdrawETH represents a WithdrawETH event raised by the L2GatewayRouter contract.
type L2GatewayRouterWithdrawETH struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawETH is a free log retrieval operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) FilterWithdrawETH(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2GatewayRouterWithdrawETHIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.FilterLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2GatewayRouterWithdrawETHIterator{contract: _L2GatewayRouter.contract, event: "WithdrawETH", logs: logs, sub: sub}, nil
}

// WatchWithdrawETH is a free log subscription operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) WatchWithdrawETH(opts *bind.WatchOpts, sink chan<- *L2GatewayRouterWithdrawETH, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2GatewayRouter.contract.WatchLogs(opts, "WithdrawETH", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2GatewayRouterWithdrawETH)
				if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
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

// ParseWithdrawETH is a log parse operation binding the contract event 0xd8ed6eaa9a7a8980d7901e911fde6686810b989d3082182d1d3a3df6306ce20e.
//
// Solidity: event WithdrawETH(address indexed from, address indexed to, uint256 amount, bytes data)
func (_L2GatewayRouter *L2GatewayRouterFilterer) ParseWithdrawETH(log types.Log) (*L2GatewayRouterWithdrawETH, error) {
	event := new(L2GatewayRouterWithdrawETH)
	if err := _L2GatewayRouter.contract.UnpackLog(event, "WithdrawETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
