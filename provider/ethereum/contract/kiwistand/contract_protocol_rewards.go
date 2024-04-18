// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kiwistand

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

// ProtocolRewardsMetaData contains all meta data concerning the ProtocolRewards contract.
var ProtocolRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ARRAY_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_SIGNATURE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_WITHDRAW\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SIGNATURE_DEADLINE_EXPIRED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TRANSFER_FAILED\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"reason\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"firstMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zora\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"name\":\"RewardsDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WITHDRAW_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"reason\",\"type\":\"bytes4\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"reasons\",\"type\":\"bytes4[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"depositBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"firstMinter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zora\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"name\":\"depositRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdrawWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProtocolRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolRewardsMetaData.ABI instead.
var ProtocolRewardsABI = ProtocolRewardsMetaData.ABI

// ProtocolRewards is an auto generated Go binding around an Ethereum contract.
type ProtocolRewards struct {
	ProtocolRewardsCaller     // Read-only binding to the contract
	ProtocolRewardsTransactor // Write-only binding to the contract
	ProtocolRewardsFilterer   // Log filterer for contract events
}

// ProtocolRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolRewardsSession struct {
	Contract     *ProtocolRewards  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolRewardsCallerSession struct {
	Contract *ProtocolRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProtocolRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolRewardsTransactorSession struct {
	Contract     *ProtocolRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProtocolRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRewardsRaw struct {
	Contract *ProtocolRewards // Generic contract binding to access the raw methods on
}

// ProtocolRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolRewardsCallerRaw struct {
	Contract *ProtocolRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolRewardsTransactorRaw struct {
	Contract *ProtocolRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolRewards creates a new instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewards(address common.Address, backend bind.ContractBackend) (*ProtocolRewards, error) {
	contract, err := bindProtocolRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewards{ProtocolRewardsCaller: ProtocolRewardsCaller{contract: contract}, ProtocolRewardsTransactor: ProtocolRewardsTransactor{contract: contract}, ProtocolRewardsFilterer: ProtocolRewardsFilterer{contract: contract}}, nil
}

// NewProtocolRewardsCaller creates a new read-only instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsCaller(address common.Address, caller bind.ContractCaller) (*ProtocolRewardsCaller, error) {
	contract, err := bindProtocolRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsCaller{contract: contract}, nil
}

// NewProtocolRewardsTransactor creates a new write-only instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolRewardsTransactor, error) {
	contract, err := bindProtocolRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsTransactor{contract: contract}, nil
}

// NewProtocolRewardsFilterer creates a new log filterer instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolRewardsFilterer, error) {
	contract, err := bindProtocolRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsFilterer{contract: contract}, nil
}

// bindProtocolRewards binds a generic wrapper to an already deployed contract.
func bindProtocolRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRewards *ProtocolRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRewards.Contract.ProtocolRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRewards *ProtocolRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ProtocolRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRewards *ProtocolRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ProtocolRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRewards *ProtocolRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRewards *ProtocolRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRewards *ProtocolRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.contract.Transact(opts, method, params...)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) WITHDRAWTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "WITHDRAW_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _ProtocolRewards.Contract.WITHDRAWTYPEHASH(&_ProtocolRewards.CallOpts)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _ProtocolRewards.Contract.WITHDRAWTYPEHASH(&_ProtocolRewards.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.BalanceOf(&_ProtocolRewards.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.BalanceOf(&_ProtocolRewards.CallOpts, arg0)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ProtocolRewards *ProtocolRewardsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ProtocolRewards *ProtocolRewardsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ProtocolRewards.Contract.Eip712Domain(&_ProtocolRewards.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _ProtocolRewards.Contract.Eip712Domain(&_ProtocolRewards.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.Nonces(&_ProtocolRewards.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.Nonces(&_ProtocolRewards.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) TotalSupply() (*big.Int, error) {
	return _ProtocolRewards.Contract.TotalSupply(&_ProtocolRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _ProtocolRewards.Contract.TotalSupply(&_ProtocolRewards.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Deposit(opts *bind.TransactOpts, to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "deposit", to, reason, comment)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsSession) Deposit(to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Deposit(&_ProtocolRewards.TransactOpts, to, reason, comment)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Deposit(to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Deposit(&_ProtocolRewards.TransactOpts, to, reason, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) DepositBatch(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "depositBatch", recipients, amounts, reasons, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsSession) DepositBatch(recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.DepositBatch(&_ProtocolRewards.TransactOpts, recipients, amounts, reasons, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) DepositBatch(recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.DepositBatch(&_ProtocolRewards.TransactOpts, recipients, amounts, reasons, comment)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) DepositRewards(opts *bind.TransactOpts, creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "depositRewards", creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_ProtocolRewards *ProtocolRewardsSession) DepositRewards(creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.DepositRewards(&_ProtocolRewards.TransactOpts, creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) DepositRewards(creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.DepositRewards(&_ProtocolRewards.TransactOpts, creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Withdraw(&_ProtocolRewards.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Withdraw(&_ProtocolRewards.TransactOpts, to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) WithdrawFor(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "withdrawFor", to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsSession) WithdrawFor(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawFor(&_ProtocolRewards.TransactOpts, to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) WithdrawFor(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawFor(&_ProtocolRewards.TransactOpts, to, amount)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) WithdrawWithSig(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "withdrawWithSig", from, to, amount, deadline, v, r, s)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ProtocolRewards *ProtocolRewardsSession) WithdrawWithSig(from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawWithSig(&_ProtocolRewards.TransactOpts, from, to, amount, deadline, v, r, s)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) WithdrawWithSig(from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawWithSig(&_ProtocolRewards.TransactOpts, from, to, amount, deadline, v, r, s)
}

// ProtocolRewardsDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ProtocolRewards contract.
type ProtocolRewardsDepositIterator struct {
	Event *ProtocolRewardsDeposit // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsDeposit)
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
		it.Event = new(ProtocolRewardsDeposit)
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
func (it *ProtocolRewardsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsDeposit represents a Deposit event raised by the ProtocolRewards contract.
type ProtocolRewardsDeposit struct {
	From    common.Address
	To      common.Address
	Reason  [4]byte
	Amount  *big.Int
	Comment string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5b769452a2090142e059af5137c5b92a3e966cfb03b794cd01ac195d196c0001.
//
// Solidity: event Deposit(address indexed from, address indexed to, bytes4 indexed reason, uint256 amount, string comment)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, to []common.Address, reason [][4]byte) (*ProtocolRewardsDepositIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Deposit", fromRule, toRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsDepositIterator{contract: _ProtocolRewards.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5b769452a2090142e059af5137c5b92a3e966cfb03b794cd01ac195d196c0001.
//
// Solidity: event Deposit(address indexed from, address indexed to, bytes4 indexed reason, uint256 amount, string comment)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsDeposit, from []common.Address, to []common.Address, reason [][4]byte) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Deposit", fromRule, toRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsDeposit)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5b769452a2090142e059af5137c5b92a3e966cfb03b794cd01ac195d196c0001.
//
// Solidity: event Deposit(address indexed from, address indexed to, bytes4 indexed reason, uint256 amount, string comment)
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseDeposit(log types.Log) (*ProtocolRewardsDeposit, error) {
	event := new(ProtocolRewardsDeposit)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the ProtocolRewards contract.
type ProtocolRewardsEIP712DomainChangedIterator struct {
	Event *ProtocolRewardsEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsEIP712DomainChanged)
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
		it.Event = new(ProtocolRewardsEIP712DomainChanged)
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
func (it *ProtocolRewardsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsEIP712DomainChanged represents a EIP712DomainChanged event raised by the ProtocolRewards contract.
type ProtocolRewardsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ProtocolRewardsEIP712DomainChangedIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsEIP712DomainChangedIterator{contract: _ProtocolRewards.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsEIP712DomainChanged)
				if err := _ProtocolRewards.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseEIP712DomainChanged(log types.Log) (*ProtocolRewardsEIP712DomainChanged, error) {
	event := new(ProtocolRewardsEIP712DomainChanged)
	if err := _ProtocolRewards.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsRewardsDepositIterator is returned from FilterRewardsDeposit and is used to iterate over the raw logs and unpacked data for RewardsDeposit events raised by the ProtocolRewards contract.
type ProtocolRewardsRewardsDepositIterator struct {
	Event *ProtocolRewardsRewardsDeposit // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsRewardsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsRewardsDeposit)
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
		it.Event = new(ProtocolRewardsRewardsDeposit)
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
func (it *ProtocolRewardsRewardsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsRewardsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsRewardsDeposit represents a RewardsDeposit event raised by the ProtocolRewards contract.
type ProtocolRewardsRewardsDeposit struct {
	Creator              common.Address
	CreateReferral       common.Address
	MintReferral         common.Address
	FirstMinter          common.Address
	Zora                 common.Address
	From                 common.Address
	CreatorReward        *big.Int
	CreateReferralReward *big.Int
	MintReferralReward   *big.Int
	FirstMinterReward    *big.Int
	ZoraReward           *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRewardsDeposit is a free log retrieval operation binding the contract event 0x90e8cce6b15b450d1e56e9ef986d1cd376838a90944336c02886ca12b9e6ebd7.
//
// Solidity: event RewardsDeposit(address indexed creator, address indexed createReferral, address indexed mintReferral, address firstMinter, address zora, address from, uint256 creatorReward, uint256 createReferralReward, uint256 mintReferralReward, uint256 firstMinterReward, uint256 zoraReward)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterRewardsDeposit(opts *bind.FilterOpts, creator []common.Address, createReferral []common.Address, mintReferral []common.Address) (*ProtocolRewardsRewardsDepositIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var createReferralRule []interface{}
	for _, createReferralItem := range createReferral {
		createReferralRule = append(createReferralRule, createReferralItem)
	}
	var mintReferralRule []interface{}
	for _, mintReferralItem := range mintReferral {
		mintReferralRule = append(mintReferralRule, mintReferralItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "RewardsDeposit", creatorRule, createReferralRule, mintReferralRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsRewardsDepositIterator{contract: _ProtocolRewards.contract, event: "RewardsDeposit", logs: logs, sub: sub}, nil
}

// WatchRewardsDeposit is a free log subscription operation binding the contract event 0x90e8cce6b15b450d1e56e9ef986d1cd376838a90944336c02886ca12b9e6ebd7.
//
// Solidity: event RewardsDeposit(address indexed creator, address indexed createReferral, address indexed mintReferral, address firstMinter, address zora, address from, uint256 creatorReward, uint256 createReferralReward, uint256 mintReferralReward, uint256 firstMinterReward, uint256 zoraReward)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchRewardsDeposit(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsRewardsDeposit, creator []common.Address, createReferral []common.Address, mintReferral []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var createReferralRule []interface{}
	for _, createReferralItem := range createReferral {
		createReferralRule = append(createReferralRule, createReferralItem)
	}
	var mintReferralRule []interface{}
	for _, mintReferralItem := range mintReferral {
		mintReferralRule = append(mintReferralRule, mintReferralItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "RewardsDeposit", creatorRule, createReferralRule, mintReferralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsRewardsDeposit)
				if err := _ProtocolRewards.contract.UnpackLog(event, "RewardsDeposit", log); err != nil {
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

// ParseRewardsDeposit is a log parse operation binding the contract event 0x90e8cce6b15b450d1e56e9ef986d1cd376838a90944336c02886ca12b9e6ebd7.
//
// Solidity: event RewardsDeposit(address indexed creator, address indexed createReferral, address indexed mintReferral, address firstMinter, address zora, address from, uint256 creatorReward, uint256 createReferralReward, uint256 mintReferralReward, uint256 firstMinterReward, uint256 zoraReward)
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseRewardsDeposit(log types.Log) (*ProtocolRewardsRewardsDeposit, error) {
	event := new(ProtocolRewardsRewardsDeposit)
	if err := _ProtocolRewards.contract.UnpackLog(event, "RewardsDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ProtocolRewards contract.
type ProtocolRewardsWithdrawIterator struct {
	Event *ProtocolRewardsWithdraw // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsWithdraw)
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
		it.Event = new(ProtocolRewardsWithdraw)
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
func (it *ProtocolRewardsWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsWithdraw represents a Withdraw event raised by the ProtocolRewards contract.
type ProtocolRewardsWithdraw struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed from, address indexed to, uint256 amount)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ProtocolRewardsWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsWithdrawIterator{contract: _ProtocolRewards.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed from, address indexed to, uint256 amount)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsWithdraw, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsWithdraw)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed from, address indexed to, uint256 amount)
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseWithdraw(log types.Log) (*ProtocolRewardsWithdraw, error) {
	event := new(ProtocolRewardsWithdraw)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
