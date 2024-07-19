// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vsl

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

// SettlementMetaData contains all meta data concerning the Settlement contract.
var SettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EPOCH_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ORACLE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOTAL_REWARDS_PER_YEAR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributeRewards\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nodeAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"operationRewards\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"requestCounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"isFinal\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBonusInfo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"staking\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracleAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"operationRewardsPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTaxRateBasisPoints4PublicPool\",\"inputs\":[{\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashNodes\",\"inputs\":[{\"name\":\"nodeAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakingContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateRewardsRatio\",\"inputs\":[{\"name\":\"operationRewardsPercent\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AlphaWithdrawNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AmountTooSmall\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BatchSizeZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallerNotStaking\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChipNotAuthorized\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChipNotPublicGood\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ChipNotValid\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nodeAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ChipsIdOverflow\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChipsNotSameOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimIdNotExists\",\"inputs\":[{\"name\":\"claimId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ClaimTimeNotReady\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CreateNodeToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyChipsIds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyNodeList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExcessWithdrawalAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidArrayLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[{\"name\":\"expected\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidEpochNumber\",\"inputs\":[{\"name\":\"current\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"got\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTraitId\",\"inputs\":[{\"name\":\"traitId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"NodeAlreadyPublicGood\",\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NodeExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeNotExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NodeNotPublicGood\",\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NodeStakedOrDeposited\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OperationRewardsExceed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicGoodNodeNotDeposited\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PublicGoodNodeNotInAlphaPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardsAlreadyDistributed\",\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SettlementPhase\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StakeToPublicGoodNode\",\"inputs\":[{\"name\":\"nodeAddr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"StakingRewardsExceed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SubmissionIntervalNotElapsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaxRateBasisPointsTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TaxRateBasisPointsTooSmall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]}]",
}

// SettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use SettlementMetaData.ABI instead.
var SettlementABI = SettlementMetaData.ABI

// Settlement is an auto generated Go binding around an Ethereum contract.
type Settlement struct {
	SettlementCaller     // Read-only binding to the contract
	SettlementTransactor // Write-only binding to the contract
	SettlementFilterer   // Log filterer for contract events
}

// SettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettlementSession struct {
	Contract     *Settlement       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettlementCallerSession struct {
	Contract *SettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettlementTransactorSession struct {
	Contract     *SettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettlementRaw struct {
	Contract *Settlement // Generic contract binding to access the raw methods on
}

// SettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettlementCallerRaw struct {
	Contract *SettlementCaller // Generic read-only contract binding to access the raw methods on
}

// SettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettlementTransactorRaw struct {
	Contract *SettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSettlement creates a new instance of Settlement, bound to a specific deployed contract.
func NewSettlement(address common.Address, backend bind.ContractBackend) (*Settlement, error) {
	contract, err := bindSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Settlement{SettlementCaller: SettlementCaller{contract: contract}, SettlementTransactor: SettlementTransactor{contract: contract}, SettlementFilterer: SettlementFilterer{contract: contract}}, nil
}

// NewSettlementCaller creates a new read-only instance of Settlement, bound to a specific deployed contract.
func NewSettlementCaller(address common.Address, caller bind.ContractCaller) (*SettlementCaller, error) {
	contract, err := bindSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementCaller{contract: contract}, nil
}

// NewSettlementTransactor creates a new write-only instance of Settlement, bound to a specific deployed contract.
func NewSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*SettlementTransactor, error) {
	contract, err := bindSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettlementTransactor{contract: contract}, nil
}

// NewSettlementFilterer creates a new log filterer instance of Settlement, bound to a specific deployed contract.
func NewSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*SettlementFilterer, error) {
	contract, err := bindSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettlementFilterer{contract: contract}, nil
}

// bindSettlement binds a generic wrapper to an already deployed contract.
func bindSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SettlementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settlement *SettlementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settlement.Contract.SettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settlement *SettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.Contract.SettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settlement *SettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settlement.Contract.SettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Settlement *SettlementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Settlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Settlement *SettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Settlement *SettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Settlement.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Settlement *SettlementCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Settlement *SettlementSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Settlement.Contract.DEFAULTADMINROLE(&_Settlement.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Settlement *SettlementCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Settlement.Contract.DEFAULTADMINROLE(&_Settlement.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Settlement *SettlementCaller) EPOCHDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "EPOCH_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Settlement *SettlementSession) EPOCHDURATION() (*big.Int, error) {
	return _Settlement.Contract.EPOCHDURATION(&_Settlement.CallOpts)
}

// EPOCHDURATION is a free data retrieval call binding the contract method 0xa70b9f0c.
//
// Solidity: function EPOCH_DURATION() view returns(uint256)
func (_Settlement *SettlementCallerSession) EPOCHDURATION() (*big.Int, error) {
	return _Settlement.Contract.EPOCHDURATION(&_Settlement.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Settlement *SettlementCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Settlement *SettlementSession) ORACLEROLE() ([32]byte, error) {
	return _Settlement.Contract.ORACLEROLE(&_Settlement.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_Settlement *SettlementCallerSession) ORACLEROLE() ([32]byte, error) {
	return _Settlement.Contract.ORACLEROLE(&_Settlement.CallOpts)
}

// TOTALREWARDSPERYEAR is a free data retrieval call binding the contract method 0xd2adf460.
//
// Solidity: function TOTAL_REWARDS_PER_YEAR() view returns(uint256)
func (_Settlement *SettlementCaller) TOTALREWARDSPERYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "TOTAL_REWARDS_PER_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALREWARDSPERYEAR is a free data retrieval call binding the contract method 0xd2adf460.
//
// Solidity: function TOTAL_REWARDS_PER_YEAR() view returns(uint256)
func (_Settlement *SettlementSession) TOTALREWARDSPERYEAR() (*big.Int, error) {
	return _Settlement.Contract.TOTALREWARDSPERYEAR(&_Settlement.CallOpts)
}

// TOTALREWARDSPERYEAR is a free data retrieval call binding the contract method 0xd2adf460.
//
// Solidity: function TOTAL_REWARDS_PER_YEAR() view returns(uint256)
func (_Settlement *SettlementCallerSession) TOTALREWARDSPERYEAR() (*big.Int, error) {
	return _Settlement.Contract.TOTALREWARDSPERYEAR(&_Settlement.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Settlement *SettlementCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Settlement *SettlementSession) CurrentEpoch() (*big.Int, error) {
	return _Settlement.Contract.CurrentEpoch(&_Settlement.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Settlement *SettlementCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Settlement.Contract.CurrentEpoch(&_Settlement.CallOpts)
}

// GetBonusInfo is a free data retrieval call binding the contract method 0xfc5567d0.
//
// Solidity: function getBonusInfo() view returns(uint256, uint256)
func (_Settlement *SettlementCaller) GetBonusInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "getBonusInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBonusInfo is a free data retrieval call binding the contract method 0xfc5567d0.
//
// Solidity: function getBonusInfo() view returns(uint256, uint256)
func (_Settlement *SettlementSession) GetBonusInfo() (*big.Int, *big.Int, error) {
	return _Settlement.Contract.GetBonusInfo(&_Settlement.CallOpts)
}

// GetBonusInfo is a free data retrieval call binding the contract method 0xfc5567d0.
//
// Solidity: function getBonusInfo() view returns(uint256, uint256)
func (_Settlement *SettlementCallerSession) GetBonusInfo() (*big.Int, *big.Int, error) {
	return _Settlement.Contract.GetBonusInfo(&_Settlement.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Settlement *SettlementCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Settlement *SettlementSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Settlement.Contract.GetRoleAdmin(&_Settlement.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Settlement *SettlementCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Settlement.Contract.GetRoleAdmin(&_Settlement.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Settlement *SettlementCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Settlement *SettlementSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Settlement.Contract.GetRoleMember(&_Settlement.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Settlement *SettlementCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Settlement.Contract.GetRoleMember(&_Settlement.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Settlement *SettlementCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Settlement *SettlementSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Settlement.Contract.GetRoleMemberCount(&_Settlement.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Settlement *SettlementCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Settlement.Contract.GetRoleMemberCount(&_Settlement.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Settlement *SettlementCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Settlement *SettlementSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Settlement.Contract.HasRole(&_Settlement.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Settlement *SettlementCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Settlement.Contract.HasRole(&_Settlement.CallOpts, role, account)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Settlement *SettlementCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Settlement *SettlementSession) StakingContract() (common.Address, error) {
	return _Settlement.Contract.StakingContract(&_Settlement.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Settlement *SettlementCallerSession) StakingContract() (common.Address, error) {
	return _Settlement.Contract.StakingContract(&_Settlement.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Settlement *SettlementCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Settlement *SettlementSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Settlement.Contract.SupportsInterface(&_Settlement.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Settlement *SettlementCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Settlement.Contract.SupportsInterface(&_Settlement.CallOpts, interfaceId)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5c3d3078.
//
// Solidity: function distributeRewards(uint256 epoch, address[] nodeAddrs, uint256[] operationRewards, uint256[] requestCounts, bool isFinal) returns()
func (_Settlement *SettlementTransactor) DistributeRewards(opts *bind.TransactOpts, epoch *big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, requestCounts []*big.Int, isFinal bool) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "distributeRewards", epoch, nodeAddrs, operationRewards, requestCounts, isFinal)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5c3d3078.
//
// Solidity: function distributeRewards(uint256 epoch, address[] nodeAddrs, uint256[] operationRewards, uint256[] requestCounts, bool isFinal) returns()
func (_Settlement *SettlementSession) DistributeRewards(epoch *big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, requestCounts []*big.Int, isFinal bool) (*types.Transaction, error) {
	return _Settlement.Contract.DistributeRewards(&_Settlement.TransactOpts, epoch, nodeAddrs, operationRewards, requestCounts, isFinal)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x5c3d3078.
//
// Solidity: function distributeRewards(uint256 epoch, address[] nodeAddrs, uint256[] operationRewards, uint256[] requestCounts, bool isFinal) returns()
func (_Settlement *SettlementTransactorSession) DistributeRewards(epoch *big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, requestCounts []*big.Int, isFinal bool) (*types.Transaction, error) {
	return _Settlement.Contract.DistributeRewards(&_Settlement.TransactOpts, epoch, nodeAddrs, operationRewards, requestCounts, isFinal)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Settlement *SettlementTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Settlement *SettlementSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.GrantRole(&_Settlement.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Settlement *SettlementTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.GrantRole(&_Settlement.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address oracleAccount, uint256 startTime, uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementTransactor) Initialize(opts *bind.TransactOpts, staking common.Address, oracleAccount common.Address, startTime *big.Int, operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "initialize", staking, oracleAccount, startTime, operationRewardsPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address oracleAccount, uint256 startTime, uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementSession) Initialize(staking common.Address, oracleAccount common.Address, startTime *big.Int, operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.Initialize(&_Settlement.TransactOpts, staking, oracleAccount, startTime, operationRewardsPercent)
}

// Initialize is a paid mutator transaction binding the contract method 0xeb990c59.
//
// Solidity: function initialize(address staking, address oracleAccount, uint256 startTime, uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementTransactorSession) Initialize(staking common.Address, oracleAccount common.Address, startTime *big.Int, operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.Initialize(&_Settlement.TransactOpts, staking, oracleAccount, startTime, operationRewardsPercent)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Settlement *SettlementTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Settlement *SettlementSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.RenounceRole(&_Settlement.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Settlement *SettlementTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.RenounceRole(&_Settlement.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Settlement *SettlementTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Settlement *SettlementSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.RevokeRole(&_Settlement.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Settlement *SettlementTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.RevokeRole(&_Settlement.TransactOpts, role, account)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_Settlement *SettlementTransactor) SetTaxRateBasisPoints4PublicPool(opts *bind.TransactOpts, taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "setTaxRateBasisPoints4PublicPool", taxRateBasisPoints)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_Settlement *SettlementSession) SetTaxRateBasisPoints4PublicPool(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _Settlement.Contract.SetTaxRateBasisPoints4PublicPool(&_Settlement.TransactOpts, taxRateBasisPoints)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_Settlement *SettlementTransactorSession) SetTaxRateBasisPoints4PublicPool(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _Settlement.Contract.SetTaxRateBasisPoints4PublicPool(&_Settlement.TransactOpts, taxRateBasisPoints)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_Settlement *SettlementTransactor) SlashNodes(opts *bind.TransactOpts, nodeAddrs []common.Address) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "slashNodes", nodeAddrs)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_Settlement *SettlementSession) SlashNodes(nodeAddrs []common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.SlashNodes(&_Settlement.TransactOpts, nodeAddrs)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_Settlement *SettlementTransactorSession) SlashNodes(nodeAddrs []common.Address) (*types.Transaction, error) {
	return _Settlement.Contract.SlashNodes(&_Settlement.TransactOpts, nodeAddrs)
}

// UpdateRewardsRatio is a paid mutator transaction binding the contract method 0xc8afeea5.
//
// Solidity: function updateRewardsRatio(uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementTransactor) UpdateRewardsRatio(opts *bind.TransactOpts, operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "updateRewardsRatio", operationRewardsPercent)
}

// UpdateRewardsRatio is a paid mutator transaction binding the contract method 0xc8afeea5.
//
// Solidity: function updateRewardsRatio(uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementSession) UpdateRewardsRatio(operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.UpdateRewardsRatio(&_Settlement.TransactOpts, operationRewardsPercent)
}

// UpdateRewardsRatio is a paid mutator transaction binding the contract method 0xc8afeea5.
//
// Solidity: function updateRewardsRatio(uint256 operationRewardsPercent) returns()
func (_Settlement *SettlementTransactorSession) UpdateRewardsRatio(operationRewardsPercent *big.Int) (*types.Transaction, error) {
	return _Settlement.Contract.UpdateRewardsRatio(&_Settlement.TransactOpts, operationRewardsPercent)
}

// SettlementInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Settlement contract.
type SettlementInitializedIterator struct {
	Event *SettlementInitialized // Event containing the contract specifics and raw log

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
func (it *SettlementInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementInitialized)
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
		it.Event = new(SettlementInitialized)
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
func (it *SettlementInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementInitialized represents a Initialized event raised by the Settlement contract.
type SettlementInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Settlement *SettlementFilterer) FilterInitialized(opts *bind.FilterOpts) (*SettlementInitializedIterator, error) {

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SettlementInitializedIterator{contract: _Settlement.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Settlement *SettlementFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SettlementInitialized) (event.Subscription, error) {

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementInitialized)
				if err := _Settlement.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Settlement *SettlementFilterer) ParseInitialized(log types.Log) (*SettlementInitialized, error) {
	event := new(SettlementInitialized)
	if err := _Settlement.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Settlement contract.
type SettlementRoleAdminChangedIterator struct {
	Event *SettlementRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SettlementRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementRoleAdminChanged)
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
		it.Event = new(SettlementRoleAdminChanged)
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
func (it *SettlementRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementRoleAdminChanged represents a RoleAdminChanged event raised by the Settlement contract.
type SettlementRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Settlement *SettlementFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SettlementRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SettlementRoleAdminChangedIterator{contract: _Settlement.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Settlement *SettlementFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SettlementRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementRoleAdminChanged)
				if err := _Settlement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Settlement *SettlementFilterer) ParseRoleAdminChanged(log types.Log) (*SettlementRoleAdminChanged, error) {
	event := new(SettlementRoleAdminChanged)
	if err := _Settlement.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Settlement contract.
type SettlementRoleGrantedIterator struct {
	Event *SettlementRoleGranted // Event containing the contract specifics and raw log

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
func (it *SettlementRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementRoleGranted)
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
		it.Event = new(SettlementRoleGranted)
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
func (it *SettlementRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementRoleGranted represents a RoleGranted event raised by the Settlement contract.
type SettlementRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SettlementRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SettlementRoleGrantedIterator{contract: _Settlement.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SettlementRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementRoleGranted)
				if err := _Settlement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) ParseRoleGranted(log types.Log) (*SettlementRoleGranted, error) {
	event := new(SettlementRoleGranted)
	if err := _Settlement.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Settlement contract.
type SettlementRoleRevokedIterator struct {
	Event *SettlementRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SettlementRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementRoleRevoked)
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
		it.Event = new(SettlementRoleRevoked)
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
func (it *SettlementRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementRoleRevoked represents a RoleRevoked event raised by the Settlement contract.
type SettlementRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SettlementRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SettlementRoleRevokedIterator{contract: _Settlement.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SettlementRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementRoleRevoked)
				if err := _Settlement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Settlement *SettlementFilterer) ParseRoleRevoked(log types.Log) (*SettlementRoleRevoked, error) {
	event := new(SettlementRoleRevoked)
	if err := _Settlement.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
