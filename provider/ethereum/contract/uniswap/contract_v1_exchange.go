// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// V1ExchangeMetaData contains all meta data concerning the V1Exchange contract.
var V1ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"TokenPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EthPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"setup\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":175875},{\"name\":\"addLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_liquidity\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":82616},{\"name\":\"removeLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116814},{\"name\":\"__default__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":true,\"type\":\"function\"},{\"name\":\"ethToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12757},{\"name\":\"ethToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12965},{\"name\":\"ethToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50463},{\"name\":\"ethToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50671},{\"name\":\"tokenToEthSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47503},{\"name\":\"tokenToEthTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47712},{\"name\":\"tokenToEthSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50175},{\"name\":\"tokenToEthTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50384},{\"name\":\"tokenToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51007},{\"name\":\"tokenToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51098},{\"name\":\"tokenToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":54928},{\"name\":\"tokenToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":55019},{\"name\":\"tokenToExchangeSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49342},{\"name\":\"tokenToExchangeTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49532},{\"name\":\"tokenToExchangeSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53233},{\"name\":\"tokenToExchangeTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53423},{\"name\":\"getEthToTokenInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5542},{\"name\":\"getEthToTokenOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6872},{\"name\":\"getTokenToEthInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5637},{\"name\":\"getTokenToEthOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6897},{\"name\":\"tokenAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1413},{\"name\":\"factoryAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1443},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1645},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75034},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110907},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38769},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1925},{\"name\":\"name\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1623},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1653},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1683},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1713}]",
}

// V1ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use V1ExchangeMetaData.ABI instead.
var V1ExchangeABI = V1ExchangeMetaData.ABI

// V1Exchange is an auto generated Go binding around an Ethereum contract.
type V1Exchange struct {
	V1ExchangeCaller     // Read-only binding to the contract
	V1ExchangeTransactor // Write-only binding to the contract
	V1ExchangeFilterer   // Log filterer for contract events
}

// V1ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1ExchangeSession struct {
	Contract     *V1Exchange       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1ExchangeCallerSession struct {
	Contract *V1ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// V1ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1ExchangeTransactorSession struct {
	Contract     *V1ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// V1ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1ExchangeRaw struct {
	Contract *V1Exchange // Generic contract binding to access the raw methods on
}

// V1ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1ExchangeCallerRaw struct {
	Contract *V1ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// V1ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1ExchangeTransactorRaw struct {
	Contract *V1ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1Exchange creates a new instance of V1Exchange, bound to a specific deployed contract.
func NewV1Exchange(address common.Address, backend bind.ContractBackend) (*V1Exchange, error) {
	contract, err := bindV1Exchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1Exchange{V1ExchangeCaller: V1ExchangeCaller{contract: contract}, V1ExchangeTransactor: V1ExchangeTransactor{contract: contract}, V1ExchangeFilterer: V1ExchangeFilterer{contract: contract}}, nil
}

// NewV1ExchangeCaller creates a new read-only instance of V1Exchange, bound to a specific deployed contract.
func NewV1ExchangeCaller(address common.Address, caller bind.ContractCaller) (*V1ExchangeCaller, error) {
	contract, err := bindV1Exchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeCaller{contract: contract}, nil
}

// NewV1ExchangeTransactor creates a new write-only instance of V1Exchange, bound to a specific deployed contract.
func NewV1ExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*V1ExchangeTransactor, error) {
	contract, err := bindV1Exchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeTransactor{contract: contract}, nil
}

// NewV1ExchangeFilterer creates a new log filterer instance of V1Exchange, bound to a specific deployed contract.
func NewV1ExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*V1ExchangeFilterer, error) {
	contract, err := bindV1Exchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeFilterer{contract: contract}, nil
}

// bindV1Exchange binds a generic wrapper to an already deployed contract.
func bindV1Exchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V1ExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Exchange *V1ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Exchange.Contract.V1ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Exchange *V1ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Exchange.Contract.V1ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Exchange *V1ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Exchange.Contract.V1ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1Exchange *V1ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1Exchange *V1ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1Exchange *V1ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1Exchange.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _V1Exchange.Contract.Allowance(&_V1Exchange.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _V1Exchange.Contract.Allowance(&_V1Exchange.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _V1Exchange.Contract.BalanceOf(&_V1Exchange.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _V1Exchange.Contract.BalanceOf(&_V1Exchange.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) Decimals() (*big.Int, error) {
	return _V1Exchange.Contract.Decimals(&_V1Exchange.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) Decimals() (*big.Int, error) {
	return _V1Exchange.Contract.Decimals(&_V1Exchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_V1Exchange *V1ExchangeCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_V1Exchange *V1ExchangeSession) FactoryAddress() (common.Address, error) {
	return _V1Exchange.Contract.FactoryAddress(&_V1Exchange.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_V1Exchange *V1ExchangeCallerSession) FactoryAddress() (common.Address, error) {
	return _V1Exchange.Contract.FactoryAddress(&_V1Exchange.CallOpts)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) GetEthToTokenInputPrice(opts *bind.CallOpts, eth_sold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "getEthToTokenInputPrice", eth_sold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetEthToTokenInputPrice(&_V1Exchange.CallOpts, eth_sold)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetEthToTokenInputPrice(&_V1Exchange.CallOpts, eth_sold)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) GetEthToTokenOutputPrice(opts *bind.CallOpts, tokens_bought *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "getEthToTokenOutputPrice", tokens_bought)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetEthToTokenOutputPrice(&_V1Exchange.CallOpts, tokens_bought)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetEthToTokenOutputPrice(&_V1Exchange.CallOpts, tokens_bought)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) GetTokenToEthInputPrice(opts *bind.CallOpts, tokens_sold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "getTokenToEthInputPrice", tokens_sold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetTokenToEthInputPrice(&_V1Exchange.CallOpts, tokens_sold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetTokenToEthInputPrice(&_V1Exchange.CallOpts, tokens_sold)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) GetTokenToEthOutputPrice(opts *bind.CallOpts, eth_bought *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "getTokenToEthOutputPrice", eth_bought)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetTokenToEthOutputPrice(&_V1Exchange.CallOpts, eth_bought)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _V1Exchange.Contract.GetTokenToEthOutputPrice(&_V1Exchange.CallOpts, eth_bought)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_V1Exchange *V1ExchangeCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_V1Exchange *V1ExchangeSession) Name() ([32]byte, error) {
	return _V1Exchange.Contract.Name(&_V1Exchange.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_V1Exchange *V1ExchangeCallerSession) Name() ([32]byte, error) {
	return _V1Exchange.Contract.Name(&_V1Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_V1Exchange *V1ExchangeCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_V1Exchange *V1ExchangeSession) Symbol() ([32]byte, error) {
	return _V1Exchange.Contract.Symbol(&_V1Exchange.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_V1Exchange *V1ExchangeCallerSession) Symbol() ([32]byte, error) {
	return _V1Exchange.Contract.Symbol(&_V1Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_V1Exchange *V1ExchangeCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_V1Exchange *V1ExchangeSession) TokenAddress() (common.Address, error) {
	return _V1Exchange.Contract.TokenAddress(&_V1Exchange.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_V1Exchange *V1ExchangeCallerSession) TokenAddress() (common.Address, error) {
	return _V1Exchange.Contract.TokenAddress(&_V1Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_V1Exchange *V1ExchangeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1Exchange.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TotalSupply() (*big.Int, error) {
	return _V1Exchange.Contract.TotalSupply(&_V1Exchange.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_V1Exchange *V1ExchangeCallerSession) TotalSupply() (*big.Int, error) {
	return _V1Exchange.Contract.TotalSupply(&_V1Exchange.CallOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_V1Exchange *V1ExchangeTransactor) Default(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "__default__")
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_V1Exchange *V1ExchangeSession) Default() (*types.Transaction, error) {
	return _V1Exchange.Contract.Default(&_V1Exchange.TransactOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_V1Exchange *V1ExchangeTransactorSession) Default() (*types.Transaction, error) {
	return _V1Exchange.Contract.Default(&_V1Exchange.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) AddLiquidity(opts *bind.TransactOpts, min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "addLiquidity", min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.AddLiquidity(&_V1Exchange.TransactOpts, min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.AddLiquidity(&_V1Exchange.TransactOpts, min_liquidity, max_tokens, deadline)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.Approve(&_V1Exchange.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.Approve(&_V1Exchange.TransactOpts, _spender, _value)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) EthToTokenSwapInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "ethToTokenSwapInput", min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenSwapInput(&_V1Exchange.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenSwapInput(&_V1Exchange.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) EthToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "ethToTokenSwapOutput", tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, deadline)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) EthToTokenTransferInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "ethToTokenTransferInput", min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenTransferInput(&_V1Exchange.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenTransferInput(&_V1Exchange.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) EthToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "ethToTokenTransferOutput", tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.EthToTokenTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, deadline, recipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_V1Exchange *V1ExchangeTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "removeLiquidity", amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_V1Exchange *V1ExchangeSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.RemoveLiquidity(&_V1Exchange.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.RemoveLiquidity(&_V1Exchange.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_V1Exchange *V1ExchangeTransactor) Setup(opts *bind.TransactOpts, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "setup", token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_V1Exchange *V1ExchangeSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.Setup(&_V1Exchange.TransactOpts, token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_V1Exchange *V1ExchangeTransactorSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.Setup(&_V1Exchange.TransactOpts, token_addr)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToEthSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToEthSwapInput", tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToEthSwapOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToEthSwapOutput", eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthSwapOutput(&_V1Exchange.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthSwapOutput(&_V1Exchange.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToEthTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToEthTransferInput", tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToEthTransferOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToEthTransferOutput", eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthTransferOutput(&_V1Exchange.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToEthTransferOutput(&_V1Exchange.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToExchangeSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToExchangeSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToExchangeSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToExchangeSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToExchangeTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToExchangeTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToExchangeTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToExchangeTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToExchangeTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToTokenSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToTokenSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenSwapInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToTokenSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenSwapOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToTokenTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToTokenTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenTransferInput(&_V1Exchange.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactor) TokenToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "tokenToTokenTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_V1Exchange *V1ExchangeTransactorSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _V1Exchange.Contract.TokenToTokenTransferOutput(&_V1Exchange.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.Transfer(&_V1Exchange.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.Transfer(&_V1Exchange.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TransferFrom(&_V1Exchange.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_V1Exchange *V1ExchangeTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _V1Exchange.Contract.TransferFrom(&_V1Exchange.TransactOpts, _from, _to, _value)
}

// V1ExchangeAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the V1Exchange contract.
type V1ExchangeAddLiquidityIterator struct {
	Event *V1ExchangeAddLiquidity // Event containing the contract specifics and raw log

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
func (it *V1ExchangeAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeAddLiquidity)
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
		it.Event = new(V1ExchangeAddLiquidity)
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
func (it *V1ExchangeAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeAddLiquidity represents a AddLiquidity event raised by the V1Exchange contract.
type V1ExchangeAddLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*V1ExchangeAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeAddLiquidityIterator{contract: _V1Exchange.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *V1ExchangeAddLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeAddLiquidity)
				if err := _V1Exchange.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) ParseAddLiquidity(log types.Log) (*V1ExchangeAddLiquidity, error) {
	event := new(V1ExchangeAddLiquidity)
	if err := _V1Exchange.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1ExchangeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the V1Exchange contract.
type V1ExchangeApprovalIterator struct {
	Event *V1ExchangeApproval // Event containing the contract specifics and raw log

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
func (it *V1ExchangeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeApproval)
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
		it.Event = new(V1ExchangeApproval)
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
func (it *V1ExchangeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeApproval represents a Approval event raised by the V1Exchange contract.
type V1ExchangeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*V1ExchangeApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeApprovalIterator{contract: _V1Exchange.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *V1ExchangeApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeApproval)
				if err := _V1Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) ParseApproval(log types.Log) (*V1ExchangeApproval, error) {
	event := new(V1ExchangeApproval)
	if err := _V1Exchange.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1ExchangeEthPurchaseIterator is returned from FilterEthPurchase and is used to iterate over the raw logs and unpacked data for EthPurchase events raised by the V1Exchange contract.
type V1ExchangeEthPurchaseIterator struct {
	Event *V1ExchangeEthPurchase // Event containing the contract specifics and raw log

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
func (it *V1ExchangeEthPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeEthPurchase)
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
		it.Event = new(V1ExchangeEthPurchase)
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
func (it *V1ExchangeEthPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeEthPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeEthPurchase represents a EthPurchase event raised by the V1Exchange contract.
type V1ExchangeEthPurchase struct {
	Buyer      common.Address
	TokensSold *big.Int
	EthBought  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthPurchase is a free log retrieval operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_V1Exchange *V1ExchangeFilterer) FilterEthPurchase(opts *bind.FilterOpts, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (*V1ExchangeEthPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeEthPurchaseIterator{contract: _V1Exchange.contract, event: "EthPurchase", logs: logs, sub: sub}, nil
}

// WatchEthPurchase is a free log subscription operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_V1Exchange *V1ExchangeFilterer) WatchEthPurchase(opts *bind.WatchOpts, sink chan<- *V1ExchangeEthPurchase, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeEthPurchase)
				if err := _V1Exchange.contract.UnpackLog(event, "EthPurchase", log); err != nil {
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

// ParseEthPurchase is a log parse operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_V1Exchange *V1ExchangeFilterer) ParseEthPurchase(log types.Log) (*V1ExchangeEthPurchase, error) {
	event := new(V1ExchangeEthPurchase)
	if err := _V1Exchange.contract.UnpackLog(event, "EthPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1ExchangeRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the V1Exchange contract.
type V1ExchangeRemoveLiquidityIterator struct {
	Event *V1ExchangeRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *V1ExchangeRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeRemoveLiquidity)
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
		it.Event = new(V1ExchangeRemoveLiquidity)
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
func (it *V1ExchangeRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeRemoveLiquidity represents a RemoveLiquidity event raised by the V1Exchange contract.
type V1ExchangeRemoveLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*V1ExchangeRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeRemoveLiquidityIterator{contract: _V1Exchange.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *V1ExchangeRemoveLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeRemoveLiquidity)
				if err := _V1Exchange.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_V1Exchange *V1ExchangeFilterer) ParseRemoveLiquidity(log types.Log) (*V1ExchangeRemoveLiquidity, error) {
	event := new(V1ExchangeRemoveLiquidity)
	if err := _V1Exchange.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1ExchangeTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the V1Exchange contract.
type V1ExchangeTokenPurchaseIterator struct {
	Event *V1ExchangeTokenPurchase // Event containing the contract specifics and raw log

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
func (it *V1ExchangeTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeTokenPurchase)
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
		it.Event = new(V1ExchangeTokenPurchase)
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
func (it *V1ExchangeTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeTokenPurchase represents a TokenPurchase event raised by the V1Exchange contract.
type V1ExchangeTokenPurchase struct {
	Buyer        common.Address
	EthSold      *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_V1Exchange *V1ExchangeFilterer) FilterTokenPurchase(opts *bind.FilterOpts, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (*V1ExchangeTokenPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeTokenPurchaseIterator{contract: _V1Exchange.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_V1Exchange *V1ExchangeFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *V1ExchangeTokenPurchase, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeTokenPurchase)
				if err := _V1Exchange.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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

// ParseTokenPurchase is a log parse operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_V1Exchange *V1ExchangeFilterer) ParseTokenPurchase(log types.Log) (*V1ExchangeTokenPurchase, error) {
	event := new(V1ExchangeTokenPurchase)
	if err := _V1Exchange.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1ExchangeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the V1Exchange contract.
type V1ExchangeTransferIterator struct {
	Event *V1ExchangeTransfer // Event containing the contract specifics and raw log

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
func (it *V1ExchangeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1ExchangeTransfer)
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
		it.Event = new(V1ExchangeTransfer)
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
func (it *V1ExchangeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1ExchangeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1ExchangeTransfer represents a Transfer event raised by the V1Exchange contract.
type V1ExchangeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*V1ExchangeTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _V1Exchange.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &V1ExchangeTransferIterator{contract: _V1Exchange.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *V1ExchangeTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _V1Exchange.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1ExchangeTransfer)
				if err := _V1Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_V1Exchange *V1ExchangeFilterer) ParseTransfer(log types.Log) (*V1ExchangeTransfer, error) {
	event := new(V1ExchangeTransfer)
	if err := _V1Exchange.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
