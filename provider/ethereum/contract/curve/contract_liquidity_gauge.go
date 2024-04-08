// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package curve

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

// LiquidityGaugeMetaData contains all meta data concerning the LiquidityGauge contract.
var LiquidityGaugeMetaData = &bind.MetaData{
	ABI: "[{\"name\":\"Deposit\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Withdraw\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"UpdateLiquidityLimit\",\"inputs\":[{\"type\":\"address\",\"name\":\"user\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"original_supply\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_balance\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"working_supply\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"lp_addr\"},{\"type\":\"address\",\"name\":\"_minter\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"name\":\"user_checkpoint\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2079152},{\"name\":\"claimable_tokens\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":1998318},{\"name\":\"kick\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2084532},{\"name\":\"set_approve_deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"addr\"},{\"type\":\"bool\",\"name\":\"can_deposit\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":35766},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"deposit\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"},{\"type\":\"address\",\"name\":\"addr\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"name\":\"withdraw\",\"outputs\":[],\"inputs\":[{\"type\":\"uint256\",\"name\":\"_value\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"gas\":2208318},{\"name\":\"integrate_checkpoint\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2297},{\"name\":\"minter\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1421},{\"name\":\"crv_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1451},{\"name\":\"lp_token\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1481},{\"name\":\"controller\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1511},{\"name\":\"voting_escrow\",\"outputs\":[{\"type\":\"address\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1541},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1725},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1601},{\"name\":\"future_epoch_time\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1631},{\"name\":\"approved_to_deposit\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"},{\"type\":\"address\",\"name\":\"arg1\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1969},{\"name\":\"working_balances\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1845},{\"name\":\"working_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1721},{\"name\":\"period\",\"outputs\":[{\"type\":\"int128\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1751},{\"name\":\"period_timestamp\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1890},{\"name\":\"integrate_inv_supply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1920},{\"name\":\"integrate_inv_supply_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1995},{\"name\":\"integrate_checkpoint_of\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2025},{\"name\":\"integrate_fraction\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"arg0\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":2055},{\"name\":\"inflation_rate\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\"}],\"inputs\":[],\"stateMutability\":\"view\",\"type\":\"function\",\"gas\":1931}]",
}

// LiquidityGaugeABI is the input ABI used to generate the binding from.
// Deprecated: Use LiquidityGaugeMetaData.ABI instead.
var LiquidityGaugeABI = LiquidityGaugeMetaData.ABI

// LiquidityGauge is an auto generated Go binding around an Ethereum contract.
type LiquidityGauge struct {
	LiquidityGaugeCaller     // Read-only binding to the contract
	LiquidityGaugeTransactor // Write-only binding to the contract
	LiquidityGaugeFilterer   // Log filterer for contract events
}

// LiquidityGaugeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LiquidityGaugeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityGaugeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LiquidityGaugeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityGaugeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LiquidityGaugeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LiquidityGaugeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LiquidityGaugeSession struct {
	Contract     *LiquidityGauge   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LiquidityGaugeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LiquidityGaugeCallerSession struct {
	Contract *LiquidityGaugeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LiquidityGaugeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LiquidityGaugeTransactorSession struct {
	Contract     *LiquidityGaugeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LiquidityGaugeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LiquidityGaugeRaw struct {
	Contract *LiquidityGauge // Generic contract binding to access the raw methods on
}

// LiquidityGaugeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LiquidityGaugeCallerRaw struct {
	Contract *LiquidityGaugeCaller // Generic read-only contract binding to access the raw methods on
}

// LiquidityGaugeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LiquidityGaugeTransactorRaw struct {
	Contract *LiquidityGaugeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLiquidityGauge creates a new instance of LiquidityGauge, bound to a specific deployed contract.
func NewLiquidityGauge(address common.Address, backend bind.ContractBackend) (*LiquidityGauge, error) {
	contract, err := bindLiquidityGauge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LiquidityGauge{LiquidityGaugeCaller: LiquidityGaugeCaller{contract: contract}, LiquidityGaugeTransactor: LiquidityGaugeTransactor{contract: contract}, LiquidityGaugeFilterer: LiquidityGaugeFilterer{contract: contract}}, nil
}

// NewLiquidityGaugeCaller creates a new read-only instance of LiquidityGauge, bound to a specific deployed contract.
func NewLiquidityGaugeCaller(address common.Address, caller bind.ContractCaller) (*LiquidityGaugeCaller, error) {
	contract, err := bindLiquidityGauge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeCaller{contract: contract}, nil
}

// NewLiquidityGaugeTransactor creates a new write-only instance of LiquidityGauge, bound to a specific deployed contract.
func NewLiquidityGaugeTransactor(address common.Address, transactor bind.ContractTransactor) (*LiquidityGaugeTransactor, error) {
	contract, err := bindLiquidityGauge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeTransactor{contract: contract}, nil
}

// NewLiquidityGaugeFilterer creates a new log filterer instance of LiquidityGauge, bound to a specific deployed contract.
func NewLiquidityGaugeFilterer(address common.Address, filterer bind.ContractFilterer) (*LiquidityGaugeFilterer, error) {
	contract, err := bindLiquidityGauge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeFilterer{contract: contract}, nil
}

// bindLiquidityGauge binds a generic wrapper to an already deployed contract.
func bindLiquidityGauge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LiquidityGaugeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityGauge *LiquidityGaugeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityGauge.Contract.LiquidityGaugeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityGauge *LiquidityGaugeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.LiquidityGaugeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityGauge *LiquidityGaugeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.LiquidityGaugeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LiquidityGauge *LiquidityGaugeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LiquidityGauge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LiquidityGauge *LiquidityGaugeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LiquidityGauge *LiquidityGaugeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.contract.Transact(opts, method, params...)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_LiquidityGauge *LiquidityGaugeCaller) ApprovedToDeposit(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "approved_to_deposit", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_LiquidityGauge *LiquidityGaugeSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _LiquidityGauge.Contract.ApprovedToDeposit(&_LiquidityGauge.CallOpts, arg0, arg1)
}

// ApprovedToDeposit is a free data retrieval call binding the contract method 0xe1522536.
//
// Solidity: function approved_to_deposit(address arg0, address arg1) view returns(bool)
func (_LiquidityGauge *LiquidityGaugeCallerSession) ApprovedToDeposit(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _LiquidityGauge.Contract.ApprovedToDeposit(&_LiquidityGauge.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.BalanceOf(&_LiquidityGauge.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.BalanceOf(&_LiquidityGauge.CallOpts, arg0)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidityGauge *LiquidityGaugeSession) Controller() (common.Address, error) {
	return _LiquidityGauge.Contract.Controller(&_LiquidityGauge.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCallerSession) Controller() (common.Address, error) {
	return _LiquidityGauge.Contract.Controller(&_LiquidityGauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCaller) CrvToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "crv_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeSession) CrvToken() (common.Address, error) {
	return _LiquidityGauge.Contract.CrvToken(&_LiquidityGauge.CallOpts)
}

// CrvToken is a free data retrieval call binding the contract method 0x76d8b117.
//
// Solidity: function crv_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCallerSession) CrvToken() (common.Address, error) {
	return _LiquidityGauge.Contract.CrvToken(&_LiquidityGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) FutureEpochTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "future_epoch_time")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) FutureEpochTime() (*big.Int, error) {
	return _LiquidityGauge.Contract.FutureEpochTime(&_LiquidityGauge.CallOpts)
}

// FutureEpochTime is a free data retrieval call binding the contract method 0xbe5d1be9.
//
// Solidity: function future_epoch_time() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) FutureEpochTime() (*big.Int, error) {
	return _LiquidityGauge.Contract.FutureEpochTime(&_LiquidityGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) InflationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "inflation_rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) InflationRate() (*big.Int, error) {
	return _LiquidityGauge.Contract.InflationRate(&_LiquidityGauge.CallOpts)
}

// InflationRate is a free data retrieval call binding the contract method 0x180692d0.
//
// Solidity: function inflation_rate() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) InflationRate() (*big.Int, error) {
	return _LiquidityGauge.Contract.InflationRate(&_LiquidityGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) IntegrateCheckpoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "integrate_checkpoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) IntegrateCheckpoint() (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateCheckpoint(&_LiquidityGauge.CallOpts)
}

// IntegrateCheckpoint is a free data retrieval call binding the contract method 0xd31f3f6d.
//
// Solidity: function integrate_checkpoint() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) IntegrateCheckpoint() (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateCheckpoint(&_LiquidityGauge.CallOpts)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) IntegrateCheckpointOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "integrate_checkpoint_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateCheckpointOf(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateCheckpointOf is a free data retrieval call binding the contract method 0x9bd324f2.
//
// Solidity: function integrate_checkpoint_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) IntegrateCheckpointOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateCheckpointOf(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) IntegrateFraction(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "integrate_fraction", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateFraction(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateFraction is a free data retrieval call binding the contract method 0x09400707.
//
// Solidity: function integrate_fraction(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) IntegrateFraction(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateFraction(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) IntegrateInvSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "integrate_inv_supply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateInvSupply(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupply is a free data retrieval call binding the contract method 0xfec8ee0c.
//
// Solidity: function integrate_inv_supply(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) IntegrateInvSupply(arg0 *big.Int) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateInvSupply(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) IntegrateInvSupplyOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "integrate_inv_supply_of", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateInvSupplyOf(&_LiquidityGauge.CallOpts, arg0)
}

// IntegrateInvSupplyOf is a free data retrieval call binding the contract method 0xde263bfa.
//
// Solidity: function integrate_inv_supply_of(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) IntegrateInvSupplyOf(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.IntegrateInvSupplyOf(&_LiquidityGauge.CallOpts, arg0)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCaller) LpToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "lp_token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeSession) LpToken() (common.Address, error) {
	return _LiquidityGauge.Contract.LpToken(&_LiquidityGauge.CallOpts)
}

// LpToken is a free data retrieval call binding the contract method 0x82c63066.
//
// Solidity: function lp_token() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCallerSession) LpToken() (common.Address, error) {
	return _LiquidityGauge.Contract.LpToken(&_LiquidityGauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_LiquidityGauge *LiquidityGaugeSession) Minter() (common.Address, error) {
	return _LiquidityGauge.Contract.Minter(&_LiquidityGauge.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCallerSession) Minter() (common.Address, error) {
	return _LiquidityGauge.Contract.Minter(&_LiquidityGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_LiquidityGauge *LiquidityGaugeCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_LiquidityGauge *LiquidityGaugeSession) Period() (*big.Int, error) {
	return _LiquidityGauge.Contract.Period(&_LiquidityGauge.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(int128)
func (_LiquidityGauge *LiquidityGaugeCallerSession) Period() (*big.Int, error) {
	return _LiquidityGauge.Contract.Period(&_LiquidityGauge.CallOpts)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) PeriodTimestamp(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "period_timestamp", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _LiquidityGauge.Contract.PeriodTimestamp(&_LiquidityGauge.CallOpts, arg0)
}

// PeriodTimestamp is a free data retrieval call binding the contract method 0x7598108c.
//
// Solidity: function period_timestamp(uint256 arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) PeriodTimestamp(arg0 *big.Int) (*big.Int, error) {
	return _LiquidityGauge.Contract.PeriodTimestamp(&_LiquidityGauge.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) TotalSupply() (*big.Int, error) {
	return _LiquidityGauge.Contract.TotalSupply(&_LiquidityGauge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) TotalSupply() (*big.Int, error) {
	return _LiquidityGauge.Contract.TotalSupply(&_LiquidityGauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCaller) VotingEscrow(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "voting_escrow")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_LiquidityGauge *LiquidityGaugeSession) VotingEscrow() (common.Address, error) {
	return _LiquidityGauge.Contract.VotingEscrow(&_LiquidityGauge.CallOpts)
}

// VotingEscrow is a free data retrieval call binding the contract method 0xdfe05031.
//
// Solidity: function voting_escrow() view returns(address)
func (_LiquidityGauge *LiquidityGaugeCallerSession) VotingEscrow() (common.Address, error) {
	return _LiquidityGauge.Contract.VotingEscrow(&_LiquidityGauge.CallOpts)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) WorkingBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "working_balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.WorkingBalances(&_LiquidityGauge.CallOpts, arg0)
}

// WorkingBalances is a free data retrieval call binding the contract method 0x13ecb1ca.
//
// Solidity: function working_balances(address arg0) view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) WorkingBalances(arg0 common.Address) (*big.Int, error) {
	return _LiquidityGauge.Contract.WorkingBalances(&_LiquidityGauge.CallOpts, arg0)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCaller) WorkingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LiquidityGauge.contract.Call(opts, &out, "working_supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) WorkingSupply() (*big.Int, error) {
	return _LiquidityGauge.Contract.WorkingSupply(&_LiquidityGauge.CallOpts)
}

// WorkingSupply is a free data retrieval call binding the contract method 0x17e28089.
//
// Solidity: function working_supply() view returns(uint256)
func (_LiquidityGauge *LiquidityGaugeCallerSession) WorkingSupply() (*big.Int, error) {
	return _LiquidityGauge.Contract.WorkingSupply(&_LiquidityGauge.CallOpts)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_LiquidityGauge *LiquidityGaugeTransactor) ClaimableTokens(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "claimable_tokens", addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_LiquidityGauge *LiquidityGaugeSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.ClaimableTokens(&_LiquidityGauge.TransactOpts, addr)
}

// ClaimableTokens is a paid mutator transaction binding the contract method 0x33134583.
//
// Solidity: function claimable_tokens(address addr) returns(uint256)
func (_LiquidityGauge *LiquidityGaugeTransactorSession) ClaimableTokens(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.ClaimableTokens(&_LiquidityGauge.TransactOpts, addr)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeTransactor) Deposit(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "deposit", _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Deposit(&_LiquidityGauge.TransactOpts, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeTransactorSession) Deposit(_value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Deposit(&_LiquidityGauge.TransactOpts, _value)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_LiquidityGauge *LiquidityGaugeTransactor) Deposit0(opts *bind.TransactOpts, _value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "deposit0", _value, addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_LiquidityGauge *LiquidityGaugeSession) Deposit0(_value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Deposit0(&_LiquidityGauge.TransactOpts, _value, addr)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 _value, address addr) returns()
func (_LiquidityGauge *LiquidityGaugeTransactorSession) Deposit0(_value *big.Int, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Deposit0(&_LiquidityGauge.TransactOpts, _value, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_LiquidityGauge *LiquidityGaugeTransactor) Kick(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "kick", addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_LiquidityGauge *LiquidityGaugeSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Kick(&_LiquidityGauge.TransactOpts, addr)
}

// Kick is a paid mutator transaction binding the contract method 0x96c55175.
//
// Solidity: function kick(address addr) returns()
func (_LiquidityGauge *LiquidityGaugeTransactorSession) Kick(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Kick(&_LiquidityGauge.TransactOpts, addr)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_LiquidityGauge *LiquidityGaugeTransactor) SetApproveDeposit(opts *bind.TransactOpts, addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "set_approve_deposit", addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_LiquidityGauge *LiquidityGaugeSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.SetApproveDeposit(&_LiquidityGauge.TransactOpts, addr, can_deposit)
}

// SetApproveDeposit is a paid mutator transaction binding the contract method 0x1d2747d4.
//
// Solidity: function set_approve_deposit(address addr, bool can_deposit) returns()
func (_LiquidityGauge *LiquidityGaugeTransactorSession) SetApproveDeposit(addr common.Address, can_deposit bool) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.SetApproveDeposit(&_LiquidityGauge.TransactOpts, addr, can_deposit)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_LiquidityGauge *LiquidityGaugeTransactor) UserCheckpoint(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "user_checkpoint", addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_LiquidityGauge *LiquidityGaugeSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.UserCheckpoint(&_LiquidityGauge.TransactOpts, addr)
}

// UserCheckpoint is a paid mutator transaction binding the contract method 0x4b820093.
//
// Solidity: function user_checkpoint(address addr) returns(bool)
func (_LiquidityGauge *LiquidityGaugeTransactorSession) UserCheckpoint(addr common.Address) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.UserCheckpoint(&_LiquidityGauge.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Withdraw(&_LiquidityGauge.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_LiquidityGauge *LiquidityGaugeTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _LiquidityGauge.Contract.Withdraw(&_LiquidityGauge.TransactOpts, _value)
}

// LiquidityGaugeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LiquidityGauge contract.
type LiquidityGaugeDepositIterator struct {
	Event *LiquidityGaugeDeposit // Event containing the contract specifics and raw log

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
func (it *LiquidityGaugeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityGaugeDeposit)
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
		it.Event = new(LiquidityGaugeDeposit)
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
func (it *LiquidityGaugeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityGaugeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityGaugeDeposit represents a Deposit event raised by the LiquidityGauge contract.
type LiquidityGaugeDeposit struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) FilterDeposit(opts *bind.FilterOpts, provider []common.Address) (*LiquidityGaugeDepositIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityGauge.contract.FilterLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeDepositIterator{contract: _LiquidityGauge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LiquidityGaugeDeposit, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityGauge.contract.WatchLogs(opts, "Deposit", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityGaugeDeposit)
				if err := _LiquidityGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) ParseDeposit(log types.Log) (*LiquidityGaugeDeposit, error) {
	event := new(LiquidityGaugeDeposit)
	if err := _LiquidityGauge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityGaugeUpdateLiquidityLimitIterator is returned from FilterUpdateLiquidityLimit and is used to iterate over the raw logs and unpacked data for UpdateLiquidityLimit events raised by the LiquidityGauge contract.
type LiquidityGaugeUpdateLiquidityLimitIterator struct {
	Event *LiquidityGaugeUpdateLiquidityLimit // Event containing the contract specifics and raw log

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
func (it *LiquidityGaugeUpdateLiquidityLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityGaugeUpdateLiquidityLimit)
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
		it.Event = new(LiquidityGaugeUpdateLiquidityLimit)
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
func (it *LiquidityGaugeUpdateLiquidityLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityGaugeUpdateLiquidityLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityGaugeUpdateLiquidityLimit represents a UpdateLiquidityLimit event raised by the LiquidityGauge contract.
type LiquidityGaugeUpdateLiquidityLimit struct {
	User            common.Address
	OriginalBalance *big.Int
	OriginalSupply  *big.Int
	WorkingBalance  *big.Int
	WorkingSupply   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityLimit is a free log retrieval operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_LiquidityGauge *LiquidityGaugeFilterer) FilterUpdateLiquidityLimit(opts *bind.FilterOpts) (*LiquidityGaugeUpdateLiquidityLimitIterator, error) {

	logs, sub, err := _LiquidityGauge.contract.FilterLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeUpdateLiquidityLimitIterator{contract: _LiquidityGauge.contract, event: "UpdateLiquidityLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityLimit is a free log subscription operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_LiquidityGauge *LiquidityGaugeFilterer) WatchUpdateLiquidityLimit(opts *bind.WatchOpts, sink chan<- *LiquidityGaugeUpdateLiquidityLimit) (event.Subscription, error) {

	logs, sub, err := _LiquidityGauge.contract.WatchLogs(opts, "UpdateLiquidityLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityGaugeUpdateLiquidityLimit)
				if err := _LiquidityGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
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

// ParseUpdateLiquidityLimit is a log parse operation binding the contract event 0x7ecd84343f76a23d2227290e0288da3251b045541698e575a5515af4f04197a3.
//
// Solidity: event UpdateLiquidityLimit(address user, uint256 original_balance, uint256 original_supply, uint256 working_balance, uint256 working_supply)
func (_LiquidityGauge *LiquidityGaugeFilterer) ParseUpdateLiquidityLimit(log types.Log) (*LiquidityGaugeUpdateLiquidityLimit, error) {
	event := new(LiquidityGaugeUpdateLiquidityLimit)
	if err := _LiquidityGauge.contract.UnpackLog(event, "UpdateLiquidityLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LiquidityGaugeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LiquidityGauge contract.
type LiquidityGaugeWithdrawIterator struct {
	Event *LiquidityGaugeWithdraw // Event containing the contract specifics and raw log

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
func (it *LiquidityGaugeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LiquidityGaugeWithdraw)
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
		it.Event = new(LiquidityGaugeWithdraw)
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
func (it *LiquidityGaugeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LiquidityGaugeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LiquidityGaugeWithdraw represents a Withdraw event raised by the LiquidityGauge contract.
type LiquidityGaugeWithdraw struct {
	Provider common.Address
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) FilterWithdraw(opts *bind.FilterOpts, provider []common.Address) (*LiquidityGaugeWithdrawIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityGauge.contract.FilterLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return &LiquidityGaugeWithdrawIterator{contract: _LiquidityGauge.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LiquidityGaugeWithdraw, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _LiquidityGauge.contract.WatchLogs(opts, "Withdraw", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LiquidityGaugeWithdraw)
				if err := _LiquidityGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed provider, uint256 value)
func (_LiquidityGauge *LiquidityGaugeFilterer) ParseWithdraw(log types.Log) (*LiquidityGaugeWithdraw, error) {
	event := new(LiquidityGaugeWithdraw)
	if err := _LiquidityGauge.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
