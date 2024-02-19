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

// KiwiMetaData contains all meta data concerning the Kiwi contract.
var KiwiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ARRAY_LENGTH_MISMATCH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_DEPOSIT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_SIGNATURE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_WITHDRAW\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SIGNATURE_DEADLINE_EXPIRED\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TRANSFER_FAILED\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"reason\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"firstMinter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"zora\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"name\":\"RewardsDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"WITHDRAW_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"reason\",\"type\":\"bytes4\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"reasons\",\"type\":\"bytes4[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"depositBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"createReferral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"firstMinter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"zora\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"name\":\"depositRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdrawWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KiwiABI is the input ABI used to generate the binding from.
// Deprecated: Use KiwiMetaData.ABI instead.
var KiwiABI = KiwiMetaData.ABI

// Kiwi is an auto generated Go binding around an Ethereum contract.
type Kiwi struct {
	KiwiCaller     // Read-only binding to the contract
	KiwiTransactor // Write-only binding to the contract
	KiwiFilterer   // Log filterer for contract events
}

// KiwiCaller is an auto generated read-only Go binding around an Ethereum contract.
type KiwiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KiwiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KiwiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KiwiSession struct {
	Contract     *Kiwi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KiwiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KiwiCallerSession struct {
	Contract *KiwiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KiwiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KiwiTransactorSession struct {
	Contract     *KiwiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KiwiRaw is an auto generated low-level Go binding around an Ethereum contract.
type KiwiRaw struct {
	Contract *Kiwi // Generic contract binding to access the raw methods on
}

// KiwiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KiwiCallerRaw struct {
	Contract *KiwiCaller // Generic read-only contract binding to access the raw methods on
}

// KiwiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KiwiTransactorRaw struct {
	Contract *KiwiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKiwi creates a new instance of Kiwi, bound to a specific deployed contract.
func NewKiwi(address common.Address, backend bind.ContractBackend) (*Kiwi, error) {
	contract, err := bindKiwi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kiwi{KiwiCaller: KiwiCaller{contract: contract}, KiwiTransactor: KiwiTransactor{contract: contract}, KiwiFilterer: KiwiFilterer{contract: contract}}, nil
}

// NewKiwiCaller creates a new read-only instance of Kiwi, bound to a specific deployed contract.
func NewKiwiCaller(address common.Address, caller bind.ContractCaller) (*KiwiCaller, error) {
	contract, err := bindKiwi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KiwiCaller{contract: contract}, nil
}

// NewKiwiTransactor creates a new write-only instance of Kiwi, bound to a specific deployed contract.
func NewKiwiTransactor(address common.Address, transactor bind.ContractTransactor) (*KiwiTransactor, error) {
	contract, err := bindKiwi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KiwiTransactor{contract: contract}, nil
}

// NewKiwiFilterer creates a new log filterer instance of Kiwi, bound to a specific deployed contract.
func NewKiwiFilterer(address common.Address, filterer bind.ContractFilterer) (*KiwiFilterer, error) {
	contract, err := bindKiwi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KiwiFilterer{contract: contract}, nil
}

// bindKiwi binds a generic wrapper to an already deployed contract.
func bindKiwi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KiwiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kiwi *KiwiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kiwi.Contract.KiwiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kiwi *KiwiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.Contract.KiwiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kiwi *KiwiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kiwi.Contract.KiwiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kiwi *KiwiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kiwi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kiwi *KiwiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kiwi *KiwiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kiwi.Contract.contract.Transact(opts, method, params...)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Kiwi *KiwiCaller) WITHDRAWTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "WITHDRAW_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Kiwi *KiwiSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _Kiwi.Contract.WITHDRAWTYPEHASH(&_Kiwi.CallOpts)
}

// WITHDRAWTYPEHASH is a free data retrieval call binding the contract method 0x76c5d758.
//
// Solidity: function WITHDRAW_TYPEHASH() view returns(bytes32)
func (_Kiwi *KiwiCallerSession) WITHDRAWTYPEHASH() ([32]byte, error) {
	return _Kiwi.Contract.WITHDRAWTYPEHASH(&_Kiwi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Kiwi *KiwiCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Kiwi *KiwiSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.BalanceOf(&_Kiwi.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Kiwi *KiwiCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.BalanceOf(&_Kiwi.CallOpts, arg0)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kiwi *KiwiCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "eip712Domain")

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
func (_Kiwi *KiwiSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kiwi.Contract.Eip712Domain(&_Kiwi.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kiwi *KiwiCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kiwi.Contract.Eip712Domain(&_Kiwi.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Kiwi *KiwiCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Kiwi *KiwiSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.Nonces(&_Kiwi.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Kiwi *KiwiCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.Nonces(&_Kiwi.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiSession) TotalSupply() (*big.Int, error) {
	return _Kiwi.Contract.TotalSupply(&_Kiwi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiCallerSession) TotalSupply() (*big.Int, error) {
	return _Kiwi.Contract.TotalSupply(&_Kiwi.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_Kiwi *KiwiTransactor) Deposit(opts *bind.TransactOpts, to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "deposit", to, reason, comment)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_Kiwi *KiwiSession) Deposit(to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.Deposit(&_Kiwi.TransactOpts, to, reason, comment)
}

// Deposit is a paid mutator transaction binding the contract method 0xf7f8266f.
//
// Solidity: function deposit(address to, bytes4 reason, string comment) payable returns()
func (_Kiwi *KiwiTransactorSession) Deposit(to common.Address, reason [4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.Deposit(&_Kiwi.TransactOpts, to, reason, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_Kiwi *KiwiTransactor) DepositBatch(opts *bind.TransactOpts, recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "depositBatch", recipients, amounts, reasons, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_Kiwi *KiwiSession) DepositBatch(recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.DepositBatch(&_Kiwi.TransactOpts, recipients, amounts, reasons, comment)
}

// DepositBatch is a paid mutator transaction binding the contract method 0x300def95.
//
// Solidity: function depositBatch(address[] recipients, uint256[] amounts, bytes4[] reasons, string comment) payable returns()
func (_Kiwi *KiwiTransactorSession) DepositBatch(recipients []common.Address, amounts []*big.Int, reasons [][4]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.DepositBatch(&_Kiwi.TransactOpts, recipients, amounts, reasons, comment)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_Kiwi *KiwiTransactor) DepositRewards(opts *bind.TransactOpts, creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "depositRewards", creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_Kiwi *KiwiSession) DepositRewards(creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.DepositRewards(&_Kiwi.TransactOpts, creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// DepositRewards is a paid mutator transaction binding the contract method 0xfaa3516f.
//
// Solidity: function depositRewards(address creator, uint256 creatorReward, address createReferral, uint256 createReferralReward, address mintReferral, uint256 mintReferralReward, address firstMinter, uint256 firstMinterReward, address zora, uint256 zoraReward) payable returns()
func (_Kiwi *KiwiTransactorSession) DepositRewards(creator common.Address, creatorReward *big.Int, createReferral common.Address, createReferralReward *big.Int, mintReferral common.Address, mintReferralReward *big.Int, firstMinter common.Address, firstMinterReward *big.Int, zora common.Address, zoraReward *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.DepositRewards(&_Kiwi.TransactOpts, creator, creatorReward, createReferral, createReferralReward, mintReferral, mintReferralReward, firstMinter, firstMinterReward, zora, zoraReward)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Kiwi *KiwiSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Withdraw(&_Kiwi.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Withdraw(&_Kiwi.TransactOpts, to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactor) WithdrawFor(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "withdrawFor", to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_Kiwi *KiwiSession) WithdrawFor(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawFor(&_Kiwi.TransactOpts, to, amount)
}

// WithdrawFor is a paid mutator transaction binding the contract method 0xdb518db2.
//
// Solidity: function withdrawFor(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactorSession) WithdrawFor(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawFor(&_Kiwi.TransactOpts, to, amount)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Kiwi *KiwiTransactor) WithdrawWithSig(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "withdrawWithSig", from, to, amount, deadline, v, r, s)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Kiwi *KiwiSession) WithdrawWithSig(from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawWithSig(&_Kiwi.TransactOpts, from, to, amount, deadline, v, r, s)
}

// WithdrawWithSig is a paid mutator transaction binding the contract method 0xc27e9794.
//
// Solidity: function withdrawWithSig(address from, address to, uint256 amount, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Kiwi *KiwiTransactorSession) WithdrawWithSig(from common.Address, to common.Address, amount *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawWithSig(&_Kiwi.TransactOpts, from, to, amount, deadline, v, r, s)
}

// KiwiDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Kiwi contract.
type KiwiDepositIterator struct {
	Event *KiwiDeposit // Event containing the contract specifics and raw log

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
func (it *KiwiDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiDeposit)
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
		it.Event = new(KiwiDeposit)
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
func (it *KiwiDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiDeposit represents a Deposit event raised by the Kiwi contract.
type KiwiDeposit struct {
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
func (_Kiwi *KiwiFilterer) FilterDeposit(opts *bind.FilterOpts, from []common.Address, to []common.Address, reason [][4]byte) (*KiwiDepositIterator, error) {

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

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Deposit", fromRule, toRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &KiwiDepositIterator{contract: _Kiwi.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5b769452a2090142e059af5137c5b92a3e966cfb03b794cd01ac195d196c0001.
//
// Solidity: event Deposit(address indexed from, address indexed to, bytes4 indexed reason, uint256 amount, string comment)
func (_Kiwi *KiwiFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *KiwiDeposit, from []common.Address, to []common.Address, reason [][4]byte) (event.Subscription, error) {

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

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Deposit", fromRule, toRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiDeposit)
				if err := _Kiwi.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseDeposit(log types.Log) (*KiwiDeposit, error) {
	event := new(KiwiDeposit)
	if err := _Kiwi.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Kiwi contract.
type KiwiEIP712DomainChangedIterator struct {
	Event *KiwiEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *KiwiEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiEIP712DomainChanged)
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
		it.Event = new(KiwiEIP712DomainChanged)
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
func (it *KiwiEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiEIP712DomainChanged represents a EIP712DomainChanged event raised by the Kiwi contract.
type KiwiEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Kiwi *KiwiFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*KiwiEIP712DomainChangedIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &KiwiEIP712DomainChangedIterator{contract: _Kiwi.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Kiwi *KiwiFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *KiwiEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiEIP712DomainChanged)
				if err := _Kiwi.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseEIP712DomainChanged(log types.Log) (*KiwiEIP712DomainChanged, error) {
	event := new(KiwiEIP712DomainChanged)
	if err := _Kiwi.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiRewardsDepositIterator is returned from FilterRewardsDeposit and is used to iterate over the raw logs and unpacked data for RewardsDeposit events raised by the Kiwi contract.
type KiwiRewardsDepositIterator struct {
	Event *KiwiRewardsDeposit // Event containing the contract specifics and raw log

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
func (it *KiwiRewardsDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiRewardsDeposit)
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
		it.Event = new(KiwiRewardsDeposit)
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
func (it *KiwiRewardsDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiRewardsDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiRewardsDeposit represents a RewardsDeposit event raised by the Kiwi contract.
type KiwiRewardsDeposit struct {
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
func (_Kiwi *KiwiFilterer) FilterRewardsDeposit(opts *bind.FilterOpts, creator []common.Address, createReferral []common.Address, mintReferral []common.Address) (*KiwiRewardsDepositIterator, error) {

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

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "RewardsDeposit", creatorRule, createReferralRule, mintReferralRule)
	if err != nil {
		return nil, err
	}
	return &KiwiRewardsDepositIterator{contract: _Kiwi.contract, event: "RewardsDeposit", logs: logs, sub: sub}, nil
}

// WatchRewardsDeposit is a free log subscription operation binding the contract event 0x90e8cce6b15b450d1e56e9ef986d1cd376838a90944336c02886ca12b9e6ebd7.
//
// Solidity: event RewardsDeposit(address indexed creator, address indexed createReferral, address indexed mintReferral, address firstMinter, address zora, address from, uint256 creatorReward, uint256 createReferralReward, uint256 mintReferralReward, uint256 firstMinterReward, uint256 zoraReward)
func (_Kiwi *KiwiFilterer) WatchRewardsDeposit(opts *bind.WatchOpts, sink chan<- *KiwiRewardsDeposit, creator []common.Address, createReferral []common.Address, mintReferral []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "RewardsDeposit", creatorRule, createReferralRule, mintReferralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiRewardsDeposit)
				if err := _Kiwi.contract.UnpackLog(event, "RewardsDeposit", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseRewardsDeposit(log types.Log) (*KiwiRewardsDeposit, error) {
	event := new(KiwiRewardsDeposit)
	if err := _Kiwi.contract.UnpackLog(event, "RewardsDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Kiwi contract.
type KiwiWithdrawIterator struct {
	Event *KiwiWithdraw // Event containing the contract specifics and raw log

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
func (it *KiwiWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiWithdraw)
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
		it.Event = new(KiwiWithdraw)
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
func (it *KiwiWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiWithdraw represents a Withdraw event raised by the Kiwi contract.
type KiwiWithdraw struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed from, address indexed to, uint256 amount)
func (_Kiwi *KiwiFilterer) FilterWithdraw(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KiwiWithdrawIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KiwiWithdrawIterator{contract: _Kiwi.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed from, address indexed to, uint256 amount)
func (_Kiwi *KiwiFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *KiwiWithdraw, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Withdraw", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiWithdraw)
				if err := _Kiwi.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseWithdraw(log types.Log) (*KiwiWithdraw, error) {
	event := new(KiwiWithdraw)
	if err := _Kiwi.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
