// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rss3

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

// StakingDeposit is an auto generated low-level Go binding around an user-defined struct.
type StakingDeposit struct {
	Amount *big.Int
	Start  uint64
	End    uint64
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsDistributed\",\"type\":\"uint256\"}],\"name\":\"RewardsDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundsWithdrawn\",\"type\":\"uint256\"}],\"name\":\"RewardsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REWARD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_LOCK_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POINTS_MULTIPLIER\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"cumulativeRewardsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getDepositsOf\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"start\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"end\",\"type\":\"uint64\"}],\"internalType\":\"structStaking.Deposit[]\",\"name\":\"_depositsOf\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDepositsOfLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStaked_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardReleased_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apr\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRewardTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pointsCorrection\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pointsPerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPerSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardReleased\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeWith180Days\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeWith270Days\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeWith360Days\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stakeWith90Days\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawableRewardsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdrawnRewardsOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingCaller) MAXLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MAX_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingSession) MAXLOCKDURATION() (*big.Int, error) {
	return _Staking.Contract.MAXLOCKDURATION(&_Staking.CallOpts)
}

// MAXLOCKDURATION is a free data retrieval call binding the contract method 0x4f1bfc9e.
//
// Solidity: function MAX_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingCallerSession) MAXLOCKDURATION() (*big.Int, error) {
	return _Staking.Contract.MAXLOCKDURATION(&_Staking.CallOpts)
}

// MAXREWARD is a free data retrieval call binding the contract method 0x0e1505e0.
//
// Solidity: function MAX_REWARD() view returns(uint256)
func (_Staking *StakingCaller) MAXREWARD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MAX_REWARD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREWARD is a free data retrieval call binding the contract method 0x0e1505e0.
//
// Solidity: function MAX_REWARD() view returns(uint256)
func (_Staking *StakingSession) MAXREWARD() (*big.Int, error) {
	return _Staking.Contract.MAXREWARD(&_Staking.CallOpts)
}

// MAXREWARD is a free data retrieval call binding the contract method 0x0e1505e0.
//
// Solidity: function MAX_REWARD() view returns(uint256)
func (_Staking *StakingCallerSession) MAXREWARD() (*big.Int, error) {
	return _Staking.Contract.MAXREWARD(&_Staking.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingCaller) MINLOCKDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MIN_LOCK_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingSession) MINLOCKDURATION() (*big.Int, error) {
	return _Staking.Contract.MINLOCKDURATION(&_Staking.CallOpts)
}

// MINLOCKDURATION is a free data retrieval call binding the contract method 0x78b4330f.
//
// Solidity: function MIN_LOCK_DURATION() view returns(uint256)
func (_Staking *StakingCallerSession) MINLOCKDURATION() (*big.Int, error) {
	return _Staking.Contract.MINLOCKDURATION(&_Staking.CallOpts)
}

// POINTSMULTIPLIER is a free data retrieval call binding the contract method 0x8f2203f6.
//
// Solidity: function POINTS_MULTIPLIER() view returns(uint128)
func (_Staking *StakingCaller) POINTSMULTIPLIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "POINTS_MULTIPLIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POINTSMULTIPLIER is a free data retrieval call binding the contract method 0x8f2203f6.
//
// Solidity: function POINTS_MULTIPLIER() view returns(uint128)
func (_Staking *StakingSession) POINTSMULTIPLIER() (*big.Int, error) {
	return _Staking.Contract.POINTSMULTIPLIER(&_Staking.CallOpts)
}

// POINTSMULTIPLIER is a free data retrieval call binding the contract method 0x8f2203f6.
//
// Solidity: function POINTS_MULTIPLIER() view returns(uint128)
func (_Staking *StakingCallerSession) POINTSMULTIPLIER() (*big.Int, error) {
	return _Staking.Contract.POINTSMULTIPLIER(&_Staking.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Staking *StakingCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Staking *StakingSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Staking.Contract.Allowance(&_Staking.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Staking *StakingCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Staking.Contract.Allowance(&_Staking.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Staking *StakingCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Staking *StakingSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Staking *StakingCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.BalanceOf(&_Staking.CallOpts, account)
}

// ClaimableTime is a free data retrieval call binding the contract method 0x5dc252b7.
//
// Solidity: function claimableTime(address ) view returns(uint256)
func (_Staking *StakingCaller) ClaimableTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "claimableTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTime is a free data retrieval call binding the contract method 0x5dc252b7.
//
// Solidity: function claimableTime(address ) view returns(uint256)
func (_Staking *StakingSession) ClaimableTime(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.ClaimableTime(&_Staking.CallOpts, arg0)
}

// ClaimableTime is a free data retrieval call binding the contract method 0x5dc252b7.
//
// Solidity: function claimableTime(address ) view returns(uint256)
func (_Staking *StakingCallerSession) ClaimableTime(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.ClaimableTime(&_Staking.CallOpts, arg0)
}

// CumulativeRewardsOf is a free data retrieval call binding the contract method 0x10accecc.
//
// Solidity: function cumulativeRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCaller) CumulativeRewardsOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "cumulativeRewardsOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardsOf is a free data retrieval call binding the contract method 0x10accecc.
//
// Solidity: function cumulativeRewardsOf(address account) view returns(uint256)
func (_Staking *StakingSession) CumulativeRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.CumulativeRewardsOf(&_Staking.CallOpts, account)
}

// CumulativeRewardsOf is a free data retrieval call binding the contract method 0x10accecc.
//
// Solidity: function cumulativeRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCallerSession) CumulativeRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.CumulativeRewardsOf(&_Staking.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staking *StakingCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staking *StakingSession) Decimals() (uint8, error) {
	return _Staking.Contract.Decimals(&_Staking.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Staking *StakingCallerSession) Decimals() (uint8, error) {
	return _Staking.Contract.Decimals(&_Staking.CallOpts)
}

// DepositsOf is a free data retrieval call binding the contract method 0xae22192e.
//
// Solidity: function depositsOf(address , uint256 ) view returns(uint256 amount, uint64 start, uint64 end)
func (_Staking *StakingCaller) DepositsOf(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Start  uint64
	End    uint64
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "depositsOf", arg0, arg1)

	outstruct := new(struct {
		Amount *big.Int
		Start  uint64
		End    uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Start = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.End = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// DepositsOf is a free data retrieval call binding the contract method 0xae22192e.
//
// Solidity: function depositsOf(address , uint256 ) view returns(uint256 amount, uint64 start, uint64 end)
func (_Staking *StakingSession) DepositsOf(arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Start  uint64
	End    uint64
}, error) {
	return _Staking.Contract.DepositsOf(&_Staking.CallOpts, arg0, arg1)
}

// DepositsOf is a free data retrieval call binding the contract method 0xae22192e.
//
// Solidity: function depositsOf(address , uint256 ) view returns(uint256 amount, uint64 start, uint64 end)
func (_Staking *StakingCallerSession) DepositsOf(arg0 common.Address, arg1 *big.Int) (struct {
	Amount *big.Int
	Start  uint64
	End    uint64
}, error) {
	return _Staking.Contract.DepositsOf(&_Staking.CallOpts, arg0, arg1)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(uint256)
func (_Staking *StakingCaller) End(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "end")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(uint256)
func (_Staking *StakingSession) End() (*big.Int, error) {
	return _Staking.Contract.End(&_Staking.CallOpts)
}

// End is a free data retrieval call binding the contract method 0xefbe1c1c.
//
// Solidity: function end() view returns(uint256)
func (_Staking *StakingCallerSession) End() (*big.Int, error) {
	return _Staking.Contract.End(&_Staking.CallOpts)
}

// GetDepositsOf is a free data retrieval call binding the contract method 0x2bb14fd2.
//
// Solidity: function getDepositsOf(address account, uint256 offset, uint256 limit) view returns((uint256,uint64,uint64)[] _depositsOf)
func (_Staking *StakingCaller) GetDepositsOf(opts *bind.CallOpts, account common.Address, offset *big.Int, limit *big.Int) ([]StakingDeposit, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDepositsOf", account, offset, limit)

	if err != nil {
		return *new([]StakingDeposit), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingDeposit)).(*[]StakingDeposit)

	return out0, err

}

// GetDepositsOf is a free data retrieval call binding the contract method 0x2bb14fd2.
//
// Solidity: function getDepositsOf(address account, uint256 offset, uint256 limit) view returns((uint256,uint64,uint64)[] _depositsOf)
func (_Staking *StakingSession) GetDepositsOf(account common.Address, offset *big.Int, limit *big.Int) ([]StakingDeposit, error) {
	return _Staking.Contract.GetDepositsOf(&_Staking.CallOpts, account, offset, limit)
}

// GetDepositsOf is a free data retrieval call binding the contract method 0x2bb14fd2.
//
// Solidity: function getDepositsOf(address account, uint256 offset, uint256 limit) view returns((uint256,uint64,uint64)[] _depositsOf)
func (_Staking *StakingCallerSession) GetDepositsOf(account common.Address, offset *big.Int, limit *big.Int) ([]StakingDeposit, error) {
	return _Staking.Contract.GetDepositsOf(&_Staking.CallOpts, account, offset, limit)
}

// GetDepositsOfLength is a free data retrieval call binding the contract method 0xb8162dd2.
//
// Solidity: function getDepositsOfLength(address account) view returns(uint256)
func (_Staking *StakingCaller) GetDepositsOfLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getDepositsOfLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositsOfLength is a free data retrieval call binding the contract method 0xb8162dd2.
//
// Solidity: function getDepositsOfLength(address account) view returns(uint256)
func (_Staking *StakingSession) GetDepositsOfLength(account common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDepositsOfLength(&_Staking.CallOpts, account)
}

// GetDepositsOfLength is a free data retrieval call binding the contract method 0xb8162dd2.
//
// Solidity: function getDepositsOfLength(address account) view returns(uint256)
func (_Staking *StakingCallerSession) GetDepositsOfLength(account common.Address) (*big.Int, error) {
	return _Staking.Contract.GetDepositsOfLength(&_Staking.CallOpts, account)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(uint256 startTime, uint256 endTime, uint256 totalStaked_, uint256 rewardReleased_, uint256 apr)
func (_Staking *StakingCaller) GetInfo(opts *bind.CallOpts) (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	TotalStaked    *big.Int
	RewardReleased *big.Int
	Apr            *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getInfo")

	outstruct := new(struct {
		StartTime      *big.Int
		EndTime        *big.Int
		TotalStaked    *big.Int
		RewardReleased *big.Int
		Apr            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardReleased = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Apr = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(uint256 startTime, uint256 endTime, uint256 totalStaked_, uint256 rewardReleased_, uint256 apr)
func (_Staking *StakingSession) GetInfo() (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	TotalStaked    *big.Int
	RewardReleased *big.Int
	Apr            *big.Int
}, error) {
	return _Staking.Contract.GetInfo(&_Staking.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() view returns(uint256 startTime, uint256 endTime, uint256 totalStaked_, uint256 rewardReleased_, uint256 apr)
func (_Staking *StakingCallerSession) GetInfo() (struct {
	StartTime      *big.Int
	EndTime        *big.Int
	TotalStaked    *big.Int
	RewardReleased *big.Int
	Apr            *big.Int
}, error) {
	return _Staking.Contract.GetInfo(&_Staking.CallOpts)
}

// LastRewardTime is a free data retrieval call binding the contract method 0x9231cf74.
//
// Solidity: function lastRewardTime() view returns(uint256)
func (_Staking *StakingCaller) LastRewardTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "lastRewardTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRewardTime is a free data retrieval call binding the contract method 0x9231cf74.
//
// Solidity: function lastRewardTime() view returns(uint256)
func (_Staking *StakingSession) LastRewardTime() (*big.Int, error) {
	return _Staking.Contract.LastRewardTime(&_Staking.CallOpts)
}

// LastRewardTime is a free data retrieval call binding the contract method 0x9231cf74.
//
// Solidity: function lastRewardTime() view returns(uint256)
func (_Staking *StakingCallerSession) LastRewardTime() (*big.Int, error) {
	return _Staking.Contract.LastRewardTime(&_Staking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staking *StakingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staking *StakingSession) Name() (string, error) {
	return _Staking.Contract.Name(&_Staking.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Staking *StakingCallerSession) Name() (string, error) {
	return _Staking.Contract.Name(&_Staking.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address account) view returns(uint256)
func (_Staking *StakingCaller) PendingRewards(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pendingRewards", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address account) view returns(uint256)
func (_Staking *StakingSession) PendingRewards(account common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingRewards(&_Staking.CallOpts, account)
}

// PendingRewards is a free data retrieval call binding the contract method 0x31d7a262.
//
// Solidity: function pendingRewards(address account) view returns(uint256)
func (_Staking *StakingCallerSession) PendingRewards(account common.Address) (*big.Int, error) {
	return _Staking.Contract.PendingRewards(&_Staking.CallOpts, account)
}

// PointsCorrection is a free data retrieval call binding the contract method 0xb182eb91.
//
// Solidity: function pointsCorrection(address ) view returns(int256)
func (_Staking *StakingCaller) PointsCorrection(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pointsCorrection", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointsCorrection is a free data retrieval call binding the contract method 0xb182eb91.
//
// Solidity: function pointsCorrection(address ) view returns(int256)
func (_Staking *StakingSession) PointsCorrection(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.PointsCorrection(&_Staking.CallOpts, arg0)
}

// PointsCorrection is a free data retrieval call binding the contract method 0xb182eb91.
//
// Solidity: function pointsCorrection(address ) view returns(int256)
func (_Staking *StakingCallerSession) PointsCorrection(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.PointsCorrection(&_Staking.CallOpts, arg0)
}

// PointsPerShare is a free data retrieval call binding the contract method 0x7e245d79.
//
// Solidity: function pointsPerShare() view returns(uint256)
func (_Staking *StakingCaller) PointsPerShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pointsPerShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PointsPerShare is a free data retrieval call binding the contract method 0x7e245d79.
//
// Solidity: function pointsPerShare() view returns(uint256)
func (_Staking *StakingSession) PointsPerShare() (*big.Int, error) {
	return _Staking.Contract.PointsPerShare(&_Staking.CallOpts)
}

// PointsPerShare is a free data retrieval call binding the contract method 0x7e245d79.
//
// Solidity: function pointsPerShare() view returns(uint256)
func (_Staking *StakingCallerSession) PointsPerShare() (*big.Int, error) {
	return _Staking.Contract.PointsPerShare(&_Staking.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_Staking *StakingCaller) RewardPerSecond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardPerSecond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_Staking *StakingSession) RewardPerSecond() (*big.Int, error) {
	return _Staking.Contract.RewardPerSecond(&_Staking.CallOpts)
}

// RewardPerSecond is a free data retrieval call binding the contract method 0x8f10369a.
//
// Solidity: function rewardPerSecond() view returns(uint256)
func (_Staking *StakingCallerSession) RewardPerSecond() (*big.Int, error) {
	return _Staking.Contract.RewardPerSecond(&_Staking.CallOpts)
}

// RewardReleased is a free data retrieval call binding the contract method 0x09dbf795.
//
// Solidity: function rewardReleased() view returns(uint256)
func (_Staking *StakingCaller) RewardReleased(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "rewardReleased")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardReleased is a free data retrieval call binding the contract method 0x09dbf795.
//
// Solidity: function rewardReleased() view returns(uint256)
func (_Staking *StakingSession) RewardReleased() (*big.Int, error) {
	return _Staking.Contract.RewardReleased(&_Staking.CallOpts)
}

// RewardReleased is a free data retrieval call binding the contract method 0x09dbf795.
//
// Solidity: function rewardReleased() view returns(uint256)
func (_Staking *StakingCallerSession) RewardReleased() (*big.Int, error) {
	return _Staking.Contract.RewardReleased(&_Staking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingSession) StakingToken() (common.Address, error) {
	return _Staking.Contract.StakingToken(&_Staking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingCallerSession) StakingToken() (common.Address, error) {
	return _Staking.Contract.StakingToken(&_Staking.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Staking *StakingCaller) Start(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "start")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Staking *StakingSession) Start() (*big.Int, error) {
	return _Staking.Contract.Start(&_Staking.CallOpts)
}

// Start is a free data retrieval call binding the contract method 0xbe9a6555.
//
// Solidity: function start() view returns(uint256)
func (_Staking *StakingCallerSession) Start() (*big.Int, error) {
	return _Staking.Contract.Start(&_Staking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staking *StakingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staking *StakingSession) Symbol() (string, error) {
	return _Staking.Contract.Symbol(&_Staking.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Staking *StakingCallerSession) Symbol() (string, error) {
	return _Staking.Contract.Symbol(&_Staking.CallOpts)
}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address ) view returns(uint256)
func (_Staking *StakingCaller) TotalDepositOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalDepositOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address ) view returns(uint256)
func (_Staking *StakingSession) TotalDepositOf(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalDepositOf(&_Staking.CallOpts, arg0)
}

// TotalDepositOf is a free data retrieval call binding the contract method 0xc5d511e1.
//
// Solidity: function totalDepositOf(address ) view returns(uint256)
func (_Staking *StakingCallerSession) TotalDepositOf(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.TotalDepositOf(&_Staking.CallOpts, arg0)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Staking *StakingCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Staking *StakingSession) TotalStaked() (*big.Int, error) {
	return _Staking.Contract.TotalStaked(&_Staking.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Staking *StakingCallerSession) TotalStaked() (*big.Int, error) {
	return _Staking.Contract.TotalStaked(&_Staking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staking *StakingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staking *StakingSession) TotalSupply() (*big.Int, error) {
	return _Staking.Contract.TotalSupply(&_Staking.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Staking *StakingCallerSession) TotalSupply() (*big.Int, error) {
	return _Staking.Contract.TotalSupply(&_Staking.CallOpts)
}

// WithdrawableRewardsOf is a free data retrieval call binding the contract method 0x7cd0b5c7.
//
// Solidity: function withdrawableRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCaller) WithdrawableRewardsOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawableRewardsOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableRewardsOf is a free data retrieval call binding the contract method 0x7cd0b5c7.
//
// Solidity: function withdrawableRewardsOf(address account) view returns(uint256)
func (_Staking *StakingSession) WithdrawableRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableRewardsOf(&_Staking.CallOpts, account)
}

// WithdrawableRewardsOf is a free data retrieval call binding the contract method 0x7cd0b5c7.
//
// Solidity: function withdrawableRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawableRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawableRewardsOf(&_Staking.CallOpts, account)
}

// WithdrawnRewards is a free data retrieval call binding the contract method 0xdd6624e4.
//
// Solidity: function withdrawnRewards(address ) view returns(uint256)
func (_Staking *StakingCaller) WithdrawnRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawnRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnRewards is a free data retrieval call binding the contract method 0xdd6624e4.
//
// Solidity: function withdrawnRewards(address ) view returns(uint256)
func (_Staking *StakingSession) WithdrawnRewards(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnRewards(&_Staking.CallOpts, arg0)
}

// WithdrawnRewards is a free data retrieval call binding the contract method 0xdd6624e4.
//
// Solidity: function withdrawnRewards(address ) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawnRewards(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnRewards(&_Staking.CallOpts, arg0)
}

// WithdrawnRewardsOf is a free data retrieval call binding the contract method 0x18f9e291.
//
// Solidity: function withdrawnRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCaller) WithdrawnRewardsOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawnRewardsOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnRewardsOf is a free data retrieval call binding the contract method 0x18f9e291.
//
// Solidity: function withdrawnRewardsOf(address account) view returns(uint256)
func (_Staking *StakingSession) WithdrawnRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnRewardsOf(&_Staking.CallOpts, account)
}

// WithdrawnRewardsOf is a free data retrieval call binding the contract method 0x18f9e291.
//
// Solidity: function withdrawnRewardsOf(address account) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawnRewardsOf(account common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnRewardsOf(&_Staking.CallOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Staking *StakingTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Staking *StakingSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Approve(&_Staking.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Staking *StakingTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Approve(&_Staking.TransactOpts, spender, amount)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _receiver) returns()
func (_Staking *StakingTransactor) ClaimRewards(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "claimRewards", _receiver)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _receiver) returns()
func (_Staking *StakingSession) ClaimRewards(_receiver common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards(&_Staking.TransactOpts, _receiver)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _receiver) returns()
func (_Staking *StakingTransactorSession) ClaimRewards(_receiver common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ClaimRewards(&_Staking.TransactOpts, _receiver)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Staking *StakingTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Staking *StakingSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DecreaseAllowance(&_Staking.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Staking *StakingTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DecreaseAllowance(&_Staking.TransactOpts, spender, subtractedValue)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Staking *StakingTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Staking *StakingSession) DistributeRewards() (*types.Transaction, error) {
	return _Staking.Contract.DistributeRewards(&_Staking.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_Staking *StakingTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _Staking.Contract.DistributeRewards(&_Staking.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Staking *StakingTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Staking *StakingSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.IncreaseAllowance(&_Staking.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Staking *StakingTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.IncreaseAllowance(&_Staking.TransactOpts, spender, addedValue)
}

// StakeWith180Days is a paid mutator transaction binding the contract method 0x616869be.
//
// Solidity: function stakeWith180Days(uint256 amount) returns()
func (_Staking *StakingTransactor) StakeWith180Days(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeWith180Days", amount)
}

// StakeWith180Days is a paid mutator transaction binding the contract method 0x616869be.
//
// Solidity: function stakeWith180Days(uint256 amount) returns()
func (_Staking *StakingSession) StakeWith180Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith180Days(&_Staking.TransactOpts, amount)
}

// StakeWith180Days is a paid mutator transaction binding the contract method 0x616869be.
//
// Solidity: function stakeWith180Days(uint256 amount) returns()
func (_Staking *StakingTransactorSession) StakeWith180Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith180Days(&_Staking.TransactOpts, amount)
}

// StakeWith270Days is a paid mutator transaction binding the contract method 0xc812e633.
//
// Solidity: function stakeWith270Days(uint256 amount) returns()
func (_Staking *StakingTransactor) StakeWith270Days(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeWith270Days", amount)
}

// StakeWith270Days is a paid mutator transaction binding the contract method 0xc812e633.
//
// Solidity: function stakeWith270Days(uint256 amount) returns()
func (_Staking *StakingSession) StakeWith270Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith270Days(&_Staking.TransactOpts, amount)
}

// StakeWith270Days is a paid mutator transaction binding the contract method 0xc812e633.
//
// Solidity: function stakeWith270Days(uint256 amount) returns()
func (_Staking *StakingTransactorSession) StakeWith270Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith270Days(&_Staking.TransactOpts, amount)
}

// StakeWith360Days is a paid mutator transaction binding the contract method 0x278dc969.
//
// Solidity: function stakeWith360Days(uint256 amount) returns()
func (_Staking *StakingTransactor) StakeWith360Days(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeWith360Days", amount)
}

// StakeWith360Days is a paid mutator transaction binding the contract method 0x278dc969.
//
// Solidity: function stakeWith360Days(uint256 amount) returns()
func (_Staking *StakingSession) StakeWith360Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith360Days(&_Staking.TransactOpts, amount)
}

// StakeWith360Days is a paid mutator transaction binding the contract method 0x278dc969.
//
// Solidity: function stakeWith360Days(uint256 amount) returns()
func (_Staking *StakingTransactorSession) StakeWith360Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith360Days(&_Staking.TransactOpts, amount)
}

// StakeWith90Days is a paid mutator transaction binding the contract method 0x383c7d87.
//
// Solidity: function stakeWith90Days(uint256 amount) returns()
func (_Staking *StakingTransactor) StakeWith90Days(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stakeWith90Days", amount)
}

// StakeWith90Days is a paid mutator transaction binding the contract method 0x383c7d87.
//
// Solidity: function stakeWith90Days(uint256 amount) returns()
func (_Staking *StakingSession) StakeWith90Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith90Days(&_Staking.TransactOpts, amount)
}

// StakeWith90Days is a paid mutator transaction binding the contract method 0x383c7d87.
//
// Solidity: function stakeWith90Days(uint256 amount) returns()
func (_Staking *StakingTransactorSession) StakeWith90Days(amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StakeWith90Days(&_Staking.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Staking *StakingTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Staking *StakingSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Transfer(&_Staking.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Staking *StakingTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Transfer(&_Staking.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Staking *StakingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Staking *StakingSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.TransferFrom(&_Staking.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Staking *StakingTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.TransferFrom(&_Staking.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 depositId, address receiver) returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, depositId *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", depositId, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 depositId, address receiver) returns()
func (_Staking *StakingSession) Withdraw(depositId *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, depositId, receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x00f714ce.
//
// Solidity: function withdraw(uint256 depositId, address receiver) returns()
func (_Staking *StakingTransactorSession) Withdraw(depositId *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, depositId, receiver)
}

// StakingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Staking contract.
type StakingApprovalIterator struct {
	Event *StakingApproval // Event containing the contract specifics and raw log

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
func (it *StakingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingApproval)
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
		it.Event = new(StakingApproval)
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
func (it *StakingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingApproval represents a Approval event raised by the Staking contract.
type StakingApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Staking *StakingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakingApprovalIterator{contract: _Staking.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Staking *StakingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakingApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingApproval)
				if err := _Staking.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Staking *StakingFilterer) ParseApproval(log types.Log) (*StakingApproval, error) {
	event := new(StakingApproval)
	if err := _Staking.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Staking contract.
type StakingDepositedIterator struct {
	Event *StakingDeposited // Event containing the contract specifics and raw log

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
func (it *StakingDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDeposited)
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
		it.Event = new(StakingDeposited)
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
func (it *StakingDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDeposited represents a Deposited event raised by the Staking contract.
type StakingDeposited struct {
	Staker   common.Address
	Amount   *big.Int
	Duration *big.Int
	Start    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address indexed staker, uint256 indexed amount, uint256 indexed duration, uint256 start)
func (_Staking *StakingFilterer) FilterDeposited(opts *bind.FilterOpts, staker []common.Address, amount []*big.Int, duration []*big.Int) (*StakingDepositedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Deposited", stakerRule, amountRule, durationRule)
	if err != nil {
		return nil, err
	}
	return &StakingDepositedIterator{contract: _Staking.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address indexed staker, uint256 indexed amount, uint256 indexed duration, uint256 start)
func (_Staking *StakingFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StakingDeposited, staker []common.Address, amount []*big.Int, duration []*big.Int) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var durationRule []interface{}
	for _, durationItem := range duration {
		durationRule = append(durationRule, durationItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Deposited", stakerRule, amountRule, durationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDeposited)
				if err := _Staking.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x91ede45f04a37a7c170f5c1207df3b6bc748dc1e04ad5e917a241d0f52feada3.
//
// Solidity: event Deposited(address indexed staker, uint256 indexed amount, uint256 indexed duration, uint256 start)
func (_Staking *StakingFilterer) ParseDeposited(log types.Log) (*StakingDeposited, error) {
	event := new(StakingDeposited)
	if err := _Staking.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Staking contract.
type StakingRewardsClaimedIterator struct {
	Event *StakingRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *StakingRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsClaimed)
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
		it.Event = new(StakingRewardsClaimed)
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
func (it *StakingRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsClaimed represents a RewardsClaimed event raised by the Staking contract.
type StakingRewardsClaimed struct {
	From         common.Address
	Receiver     common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed _from, address indexed _receiver, uint256 indexed rewardAmount)
func (_Staking *StakingFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, _from []common.Address, _receiver []common.Address, rewardAmount []*big.Int) (*StakingRewardsClaimedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}
	var rewardAmountRule []interface{}
	for _, rewardAmountItem := range rewardAmount {
		rewardAmountRule = append(rewardAmountRule, rewardAmountItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsClaimed", _fromRule, _receiverRule, rewardAmountRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsClaimedIterator{contract: _Staking.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed _from, address indexed _receiver, uint256 indexed rewardAmount)
func (_Staking *StakingFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *StakingRewardsClaimed, _from []common.Address, _receiver []common.Address, rewardAmount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _receiverRule []interface{}
	for _, _receiverItem := range _receiver {
		_receiverRule = append(_receiverRule, _receiverItem)
	}
	var rewardAmountRule []interface{}
	for _, rewardAmountItem := range rewardAmount {
		rewardAmountRule = append(rewardAmountRule, rewardAmountItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsClaimed", _fromRule, _receiverRule, rewardAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsClaimed)
				if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x9310ccfcb8de723f578a9e4282ea9f521f05ae40dc08f3068dfad528a65ee3c7.
//
// Solidity: event RewardsClaimed(address indexed _from, address indexed _receiver, uint256 indexed rewardAmount)
func (_Staking *StakingFilterer) ParseRewardsClaimed(log types.Log) (*StakingRewardsClaimed, error) {
	event := new(StakingRewardsClaimed)
	if err := _Staking.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsDistributedIterator is returned from FilterRewardsDistributed and is used to iterate over the raw logs and unpacked data for RewardsDistributed events raised by the Staking contract.
type StakingRewardsDistributedIterator struct {
	Event *StakingRewardsDistributed // Event containing the contract specifics and raw log

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
func (it *StakingRewardsDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsDistributed)
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
		it.Event = new(StakingRewardsDistributed)
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
func (it *StakingRewardsDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsDistributed represents a RewardsDistributed event raised by the Staking contract.
type StakingRewardsDistributed struct {
	By                 common.Address
	RewardsDistributed *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributed is a free log retrieval operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed by, uint256 rewardsDistributed)
func (_Staking *StakingFilterer) FilterRewardsDistributed(opts *bind.FilterOpts, by []common.Address) (*StakingRewardsDistributedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsDistributedIterator{contract: _Staking.contract, event: "RewardsDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributed is a free log subscription operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed by, uint256 rewardsDistributed)
func (_Staking *StakingFilterer) WatchRewardsDistributed(opts *bind.WatchOpts, sink chan<- *StakingRewardsDistributed, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsDistributed", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsDistributed)
				if err := _Staking.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
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

// ParseRewardsDistributed is a log parse operation binding the contract event 0xdf29796aad820e4bb192f3a8d631b76519bcd2cbe77cc85af20e9df53cece086.
//
// Solidity: event RewardsDistributed(address indexed by, uint256 rewardsDistributed)
func (_Staking *StakingFilterer) ParseRewardsDistributed(log types.Log) (*StakingRewardsDistributed, error) {
	event := new(StakingRewardsDistributed)
	if err := _Staking.contract.UnpackLog(event, "RewardsDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRewardsWithdrawnIterator is returned from FilterRewardsWithdrawn and is used to iterate over the raw logs and unpacked data for RewardsWithdrawn events raised by the Staking contract.
type StakingRewardsWithdrawnIterator struct {
	Event *StakingRewardsWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingRewardsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRewardsWithdrawn)
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
		it.Event = new(StakingRewardsWithdrawn)
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
func (it *StakingRewardsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRewardsWithdrawn represents a RewardsWithdrawn event raised by the Staking contract.
type StakingRewardsWithdrawn struct {
	By             common.Address
	FundsWithdrawn *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsWithdrawn is a free log retrieval operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed by, uint256 fundsWithdrawn)
func (_Staking *StakingFilterer) FilterRewardsWithdrawn(opts *bind.FilterOpts, by []common.Address) (*StakingRewardsWithdrawnIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RewardsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return &StakingRewardsWithdrawnIterator{contract: _Staking.contract, event: "RewardsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchRewardsWithdrawn is a free log subscription operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed by, uint256 fundsWithdrawn)
func (_Staking *StakingFilterer) WatchRewardsWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingRewardsWithdrawn, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RewardsWithdrawn", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRewardsWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
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

// ParseRewardsWithdrawn is a log parse operation binding the contract event 0x8a43c4352486ec339f487f64af78ca5cbf06cd47833f073d3baf3a193e503161.
//
// Solidity: event RewardsWithdrawn(address indexed by, uint256 fundsWithdrawn)
func (_Staking *StakingFilterer) ParseRewardsWithdrawn(log types.Log) (*StakingRewardsWithdrawn, error) {
	event := new(StakingRewardsWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "RewardsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Staking contract.
type StakingTransferIterator struct {
	Event *StakingTransfer // Event containing the contract specifics and raw log

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
func (it *StakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTransfer)
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
		it.Event = new(StakingTransfer)
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
func (it *StakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTransfer represents a Transfer event raised by the Staking contract.
type StakingTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakingTransferIterator{contract: _Staking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTransfer)
				if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Staking *StakingFilterer) ParseTransfer(log types.Log) (*StakingTransfer, error) {
	event := new(StakingTransfer)
	if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Staking contract.
type StakingWithdrawnIterator struct {
	Event *StakingWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawn)
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
		it.Event = new(StakingWithdrawn)
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
func (it *StakingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawn represents a Withdrawn event raised by the Staking contract.
type StakingWithdrawn struct {
	DepositId *big.Int
	Receiver  common.Address
	From      common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0xe5df19de43c8c04fd192bc68e484b2593570925fbb6ad8c07ccafbc2aa5c37a1.
//
// Solidity: event Withdrawn(uint256 indexed depositId, address indexed receiver, address indexed from, uint256 amount)
func (_Staking *StakingFilterer) FilterWithdrawn(opts *bind.FilterOpts, depositId []*big.Int, receiver []common.Address, from []common.Address) (*StakingWithdrawnIterator, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdrawn", depositIdRule, receiverRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawnIterator{contract: _Staking.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0xe5df19de43c8c04fd192bc68e484b2593570925fbb6ad8c07ccafbc2aa5c37a1.
//
// Solidity: event Withdrawn(uint256 indexed depositId, address indexed receiver, address indexed from, uint256 amount)
func (_Staking *StakingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingWithdrawn, depositId []*big.Int, receiver []common.Address, from []common.Address) (event.Subscription, error) {

	var depositIdRule []interface{}
	for _, depositIdItem := range depositId {
		depositIdRule = append(depositIdRule, depositIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdrawn", depositIdRule, receiverRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0xe5df19de43c8c04fd192bc68e484b2593570925fbb6ad8c07ccafbc2aa5c37a1.
//
// Solidity: event Withdrawn(uint256 indexed depositId, address indexed receiver, address indexed from, uint256 amount)
func (_Staking *StakingFilterer) ParseWithdrawn(log types.Log) (*StakingWithdrawn, error) {
	event := new(StakingWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
