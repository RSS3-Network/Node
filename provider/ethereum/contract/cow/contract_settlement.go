// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cow

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

// GPv2InteractionData is an auto generated low-level Go binding around an user-defined struct.
type GPv2InteractionData struct {
	Target   common.Address
	Value    *big.Int
	CallData []byte
}

// GPv2TradeData is an auto generated low-level Go binding around an user-defined struct.
type GPv2TradeData struct {
	SellTokenIndex *big.Int
	BuyTokenIndex  *big.Int
	Receiver       common.Address
	SellAmount     *big.Int
	BuyAmount      *big.Int
	ValidTo        uint32
	AppData        [32]byte
	FeeAmount      *big.Int
	Flags          *big.Int
	ExecutedAmount *big.Int
	Signature      []byte
}

// IVaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IVaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// SettlementMetaData contains all meta data concerning the Settlement contract.
var SettlementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractGPv2Authentication\",\"name\":\"authenticator_\",\"type\":\"address\"},{\"internalType\":\"contractIVault\",\"name\":\"vault_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"Interaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"OrderInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"signed\",\"type\":\"bool\"}],\"name\":\"PreSignature\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"sellToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"buyToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"authenticator\",\"outputs\":[{\"internalType\":\"contractGPv2Authentication\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"filledAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"orderUids\",\"type\":\"bytes[]\"}],\"name\":\"freeFilledAmountStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"orderUids\",\"type\":\"bytes[]\"}],\"name\":\"freePreSignatureStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"getStorageAt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"}],\"name\":\"invalidateOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"preSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"orderUid\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"signed\",\"type\":\"bool\"}],\"name\":\"setPreSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"clearingPrices\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sellTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validTo\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"appData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Trade.Data[]\",\"name\":\"trades\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Interaction.Data[][3]\",\"name\":\"interactions\",\"type\":\"tuple[][3]\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateDelegatecall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targetContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"calldataPayload\",\"type\":\"bytes\"}],\"name\":\"simulateDelegatecallInternal\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"response\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIVault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sellTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyTokenIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"validTo\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"appData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executedAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structGPv2Trade.Data\",\"name\":\"trade\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultRelayer\",\"outputs\":[{\"internalType\":\"contractGPv2VaultRelayer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_Settlement *SettlementCaller) Authenticator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "authenticator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_Settlement *SettlementSession) Authenticator() (common.Address, error) {
	return _Settlement.Contract.Authenticator(&_Settlement.CallOpts)
}

// Authenticator is a free data retrieval call binding the contract method 0x2335c76b.
//
// Solidity: function authenticator() view returns(address)
func (_Settlement *SettlementCallerSession) Authenticator() (common.Address, error) {
	return _Settlement.Contract.Authenticator(&_Settlement.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Settlement *SettlementCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Settlement *SettlementSession) DomainSeparator() ([32]byte, error) {
	return _Settlement.Contract.DomainSeparator(&_Settlement.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_Settlement *SettlementCallerSession) DomainSeparator() ([32]byte, error) {
	return _Settlement.Contract.DomainSeparator(&_Settlement.CallOpts)
}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_Settlement *SettlementCaller) FilledAmount(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "filledAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_Settlement *SettlementSession) FilledAmount(arg0 []byte) (*big.Int, error) {
	return _Settlement.Contract.FilledAmount(&_Settlement.CallOpts, arg0)
}

// FilledAmount is a free data retrieval call binding the contract method 0x2479fb6e.
//
// Solidity: function filledAmount(bytes ) view returns(uint256)
func (_Settlement *SettlementCallerSession) FilledAmount(arg0 []byte) (*big.Int, error) {
	return _Settlement.Contract.FilledAmount(&_Settlement.CallOpts, arg0)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Settlement *SettlementCaller) GetStorageAt(opts *bind.CallOpts, offset *big.Int, length *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "getStorageAt", offset, length)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Settlement *SettlementSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Settlement.Contract.GetStorageAt(&_Settlement.CallOpts, offset, length)
}

// GetStorageAt is a free data retrieval call binding the contract method 0x5624b25b.
//
// Solidity: function getStorageAt(uint256 offset, uint256 length) view returns(bytes)
func (_Settlement *SettlementCallerSession) GetStorageAt(offset *big.Int, length *big.Int) ([]byte, error) {
	return _Settlement.Contract.GetStorageAt(&_Settlement.CallOpts, offset, length)
}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_Settlement *SettlementCaller) PreSignature(opts *bind.CallOpts, arg0 []byte) (*big.Int, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "preSignature", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_Settlement *SettlementSession) PreSignature(arg0 []byte) (*big.Int, error) {
	return _Settlement.Contract.PreSignature(&_Settlement.CallOpts, arg0)
}

// PreSignature is a free data retrieval call binding the contract method 0xd08d33d1.
//
// Solidity: function preSignature(bytes ) view returns(uint256)
func (_Settlement *SettlementCallerSession) PreSignature(arg0 []byte) (*big.Int, error) {
	return _Settlement.Contract.PreSignature(&_Settlement.CallOpts, arg0)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Settlement *SettlementCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Settlement *SettlementSession) Vault() (common.Address, error) {
	return _Settlement.Contract.Vault(&_Settlement.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Settlement *SettlementCallerSession) Vault() (common.Address, error) {
	return _Settlement.Contract.Vault(&_Settlement.CallOpts)
}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_Settlement *SettlementCaller) VaultRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Settlement.contract.Call(opts, &out, "vaultRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_Settlement *SettlementSession) VaultRelayer() (common.Address, error) {
	return _Settlement.Contract.VaultRelayer(&_Settlement.CallOpts)
}

// VaultRelayer is a free data retrieval call binding the contract method 0x9b552cc2.
//
// Solidity: function vaultRelayer() view returns(address)
func (_Settlement *SettlementCallerSession) VaultRelayer() (common.Address, error) {
	return _Settlement.Contract.VaultRelayer(&_Settlement.CallOpts)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementTransactor) FreeFilledAmountStorage(opts *bind.TransactOpts, orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "freeFilledAmountStorage", orderUids)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementSession) FreeFilledAmountStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.Contract.FreeFilledAmountStorage(&_Settlement.TransactOpts, orderUids)
}

// FreeFilledAmountStorage is a paid mutator transaction binding the contract method 0xed9f35ce.
//
// Solidity: function freeFilledAmountStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementTransactorSession) FreeFilledAmountStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.Contract.FreeFilledAmountStorage(&_Settlement.TransactOpts, orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementTransactor) FreePreSignatureStorage(opts *bind.TransactOpts, orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "freePreSignatureStorage", orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementSession) FreePreSignatureStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.Contract.FreePreSignatureStorage(&_Settlement.TransactOpts, orderUids)
}

// FreePreSignatureStorage is a paid mutator transaction binding the contract method 0xa2a7d51b.
//
// Solidity: function freePreSignatureStorage(bytes[] orderUids) returns()
func (_Settlement *SettlementTransactorSession) FreePreSignatureStorage(orderUids [][]byte) (*types.Transaction, error) {
	return _Settlement.Contract.FreePreSignatureStorage(&_Settlement.TransactOpts, orderUids)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_Settlement *SettlementTransactor) InvalidateOrder(opts *bind.TransactOpts, orderUid []byte) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "invalidateOrder", orderUid)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_Settlement *SettlementSession) InvalidateOrder(orderUid []byte) (*types.Transaction, error) {
	return _Settlement.Contract.InvalidateOrder(&_Settlement.TransactOpts, orderUid)
}

// InvalidateOrder is a paid mutator transaction binding the contract method 0x15337bc0.
//
// Solidity: function invalidateOrder(bytes orderUid) returns()
func (_Settlement *SettlementTransactorSession) InvalidateOrder(orderUid []byte) (*types.Transaction, error) {
	return _Settlement.Contract.InvalidateOrder(&_Settlement.TransactOpts, orderUid)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_Settlement *SettlementTransactor) SetPreSignature(opts *bind.TransactOpts, orderUid []byte, signed bool) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "setPreSignature", orderUid, signed)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_Settlement *SettlementSession) SetPreSignature(orderUid []byte, signed bool) (*types.Transaction, error) {
	return _Settlement.Contract.SetPreSignature(&_Settlement.TransactOpts, orderUid, signed)
}

// SetPreSignature is a paid mutator transaction binding the contract method 0xec6cb13f.
//
// Solidity: function setPreSignature(bytes orderUid, bool signed) returns()
func (_Settlement *SettlementTransactorSession) SetPreSignature(orderUid []byte, signed bool) (*types.Transaction, error) {
	return _Settlement.Contract.SetPreSignature(&_Settlement.TransactOpts, orderUid, signed)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_Settlement *SettlementTransactor) Settle(opts *bind.TransactOpts, tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "settle", tokens, clearingPrices, trades, interactions)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_Settlement *SettlementSession) Settle(tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _Settlement.Contract.Settle(&_Settlement.TransactOpts, tokens, clearingPrices, trades, interactions)
}

// Settle is a paid mutator transaction binding the contract method 0x13d79a0b.
//
// Solidity: function settle(address[] tokens, uint256[] clearingPrices, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes)[] trades, (address,uint256,bytes)[][3] interactions) returns()
func (_Settlement *SettlementTransactorSession) Settle(tokens []common.Address, clearingPrices []*big.Int, trades []GPv2TradeData, interactions [3][]GPv2InteractionData) (*types.Transaction, error) {
	return _Settlement.Contract.Settle(&_Settlement.TransactOpts, tokens, clearingPrices, trades, interactions)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementTransactor) SimulateDelegatecall(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "simulateDelegatecall", targetContract, calldataPayload)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementSession) SimulateDelegatecall(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.Contract.SimulateDelegatecall(&_Settlement.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecall is a paid mutator transaction binding the contract method 0xf84436bd.
//
// Solidity: function simulateDelegatecall(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementTransactorSession) SimulateDelegatecall(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.Contract.SimulateDelegatecall(&_Settlement.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementTransactor) SimulateDelegatecallInternal(opts *bind.TransactOpts, targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "simulateDelegatecallInternal", targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementSession) SimulateDelegatecallInternal(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.Contract.SimulateDelegatecallInternal(&_Settlement.TransactOpts, targetContract, calldataPayload)
}

// SimulateDelegatecallInternal is a paid mutator transaction binding the contract method 0x43218e19.
//
// Solidity: function simulateDelegatecallInternal(address targetContract, bytes calldataPayload) returns(bytes response)
func (_Settlement *SettlementTransactorSession) SimulateDelegatecallInternal(targetContract common.Address, calldataPayload []byte) (*types.Transaction, error) {
	return _Settlement.Contract.SimulateDelegatecallInternal(&_Settlement.TransactOpts, targetContract, calldataPayload)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_Settlement *SettlementTransactor) Swap(opts *bind.TransactOpts, swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _Settlement.contract.Transact(opts, "swap", swaps, tokens, trade)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_Settlement *SettlementSession) Swap(swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _Settlement.Contract.Swap(&_Settlement.TransactOpts, swaps, tokens, trade)
}

// Swap is a paid mutator transaction binding the contract method 0x845a101f.
//
// Solidity: function swap((bytes32,uint256,uint256,uint256,bytes)[] swaps, address[] tokens, (uint256,uint256,address,uint256,uint256,uint32,bytes32,uint256,uint256,uint256,bytes) trade) returns()
func (_Settlement *SettlementTransactorSession) Swap(swaps []IVaultBatchSwapStep, tokens []common.Address, trade GPv2TradeData) (*types.Transaction, error) {
	return _Settlement.Contract.Swap(&_Settlement.TransactOpts, swaps, tokens, trade)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Settlement *SettlementTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Settlement.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Settlement *SettlementSession) Receive() (*types.Transaction, error) {
	return _Settlement.Contract.Receive(&_Settlement.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Settlement *SettlementTransactorSession) Receive() (*types.Transaction, error) {
	return _Settlement.Contract.Receive(&_Settlement.TransactOpts)
}

// SettlementInteractionIterator is returned from FilterInteraction and is used to iterate over the raw logs and unpacked data for Interaction events raised by the Settlement contract.
type SettlementInteractionIterator struct {
	Event *SettlementInteraction // Event containing the contract specifics and raw log

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
func (it *SettlementInteractionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementInteraction)
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
		it.Event = new(SettlementInteraction)
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
func (it *SettlementInteractionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementInteractionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementInteraction represents a Interaction event raised by the Settlement contract.
type SettlementInteraction struct {
	Target   common.Address
	Value    *big.Int
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInteraction is a free log retrieval operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_Settlement *SettlementFilterer) FilterInteraction(opts *bind.FilterOpts, target []common.Address) (*SettlementInteractionIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "Interaction", targetRule)
	if err != nil {
		return nil, err
	}
	return &SettlementInteractionIterator{contract: _Settlement.contract, event: "Interaction", logs: logs, sub: sub}, nil
}

// WatchInteraction is a free log subscription operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_Settlement *SettlementFilterer) WatchInteraction(opts *bind.WatchOpts, sink chan<- *SettlementInteraction, target []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "Interaction", targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementInteraction)
				if err := _Settlement.contract.UnpackLog(event, "Interaction", log); err != nil {
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

// ParseInteraction is a log parse operation binding the contract event 0xed99827efb37016f2275f98c4bcf71c7551c75d59e9b450f79fa32e60be672c2.
//
// Solidity: event Interaction(address indexed target, uint256 value, bytes4 selector)
func (_Settlement *SettlementFilterer) ParseInteraction(log types.Log) (*SettlementInteraction, error) {
	event := new(SettlementInteraction)
	if err := _Settlement.contract.UnpackLog(event, "Interaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementOrderInvalidatedIterator is returned from FilterOrderInvalidated and is used to iterate over the raw logs and unpacked data for OrderInvalidated events raised by the Settlement contract.
type SettlementOrderInvalidatedIterator struct {
	Event *SettlementOrderInvalidated // Event containing the contract specifics and raw log

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
func (it *SettlementOrderInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementOrderInvalidated)
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
		it.Event = new(SettlementOrderInvalidated)
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
func (it *SettlementOrderInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementOrderInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementOrderInvalidated represents a OrderInvalidated event raised by the Settlement contract.
type SettlementOrderInvalidated struct {
	Owner    common.Address
	OrderUid []byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrderInvalidated is a free log retrieval operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_Settlement *SettlementFilterer) FilterOrderInvalidated(opts *bind.FilterOpts, owner []common.Address) (*SettlementOrderInvalidatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "OrderInvalidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SettlementOrderInvalidatedIterator{contract: _Settlement.contract, event: "OrderInvalidated", logs: logs, sub: sub}, nil
}

// WatchOrderInvalidated is a free log subscription operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_Settlement *SettlementFilterer) WatchOrderInvalidated(opts *bind.WatchOpts, sink chan<- *SettlementOrderInvalidated, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "OrderInvalidated", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementOrderInvalidated)
				if err := _Settlement.contract.UnpackLog(event, "OrderInvalidated", log); err != nil {
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

// ParseOrderInvalidated is a log parse operation binding the contract event 0x875b6cb035bbd4ac6500fabc6d1e4ca5bdc58a3e2b424ccb5c24cdbebeb009a9.
//
// Solidity: event OrderInvalidated(address indexed owner, bytes orderUid)
func (_Settlement *SettlementFilterer) ParseOrderInvalidated(log types.Log) (*SettlementOrderInvalidated, error) {
	event := new(SettlementOrderInvalidated)
	if err := _Settlement.contract.UnpackLog(event, "OrderInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementPreSignatureIterator is returned from FilterPreSignature and is used to iterate over the raw logs and unpacked data for PreSignature events raised by the Settlement contract.
type SettlementPreSignatureIterator struct {
	Event *SettlementPreSignature // Event containing the contract specifics and raw log

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
func (it *SettlementPreSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementPreSignature)
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
		it.Event = new(SettlementPreSignature)
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
func (it *SettlementPreSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementPreSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementPreSignature represents a PreSignature event raised by the Settlement contract.
type SettlementPreSignature struct {
	Owner    common.Address
	OrderUid []byte
	Signed   bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPreSignature is a free log retrieval operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_Settlement *SettlementFilterer) FilterPreSignature(opts *bind.FilterOpts, owner []common.Address) (*SettlementPreSignatureIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "PreSignature", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SettlementPreSignatureIterator{contract: _Settlement.contract, event: "PreSignature", logs: logs, sub: sub}, nil
}

// WatchPreSignature is a free log subscription operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_Settlement *SettlementFilterer) WatchPreSignature(opts *bind.WatchOpts, sink chan<- *SettlementPreSignature, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "PreSignature", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementPreSignature)
				if err := _Settlement.contract.UnpackLog(event, "PreSignature", log); err != nil {
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

// ParsePreSignature is a log parse operation binding the contract event 0x01bf7c8b0ca55deecbea89d7e58295b7ffbf685fd0d96801034ba8c6ffe1c68d.
//
// Solidity: event PreSignature(address indexed owner, bytes orderUid, bool signed)
func (_Settlement *SettlementFilterer) ParsePreSignature(log types.Log) (*SettlementPreSignature, error) {
	event := new(SettlementPreSignature)
	if err := _Settlement.contract.UnpackLog(event, "PreSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the Settlement contract.
type SettlementSettlementIterator struct {
	Event *SettlementSettlement // Event containing the contract specifics and raw log

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
func (it *SettlementSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementSettlement)
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
		it.Event = new(SettlementSettlement)
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
func (it *SettlementSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementSettlement represents a Settlement event raised by the Settlement contract.
type SettlementSettlement struct {
	Solver common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_Settlement *SettlementFilterer) FilterSettlement(opts *bind.FilterOpts, solver []common.Address) (*SettlementSettlementIterator, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "Settlement", solverRule)
	if err != nil {
		return nil, err
	}
	return &SettlementSettlementIterator{contract: _Settlement.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_Settlement *SettlementFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *SettlementSettlement, solver []common.Address) (event.Subscription, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "Settlement", solverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementSettlement)
				if err := _Settlement.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// ParseSettlement is a log parse operation binding the contract event 0x40338ce1a7c49204f0099533b1e9a7ee0a3d261f84974ab7af36105b8c4e9db4.
//
// Solidity: event Settlement(address indexed solver)
func (_Settlement *SettlementFilterer) ParseSettlement(log types.Log) (*SettlementSettlement, error) {
	event := new(SettlementSettlement)
	if err := _Settlement.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SettlementTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Settlement contract.
type SettlementTradeIterator struct {
	Event *SettlementTrade // Event containing the contract specifics and raw log

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
func (it *SettlementTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettlementTrade)
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
		it.Event = new(SettlementTrade)
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
func (it *SettlementTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettlementTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettlementTrade represents a Trade event raised by the Settlement contract.
type SettlementTrade struct {
	Owner      common.Address
	SellToken  common.Address
	BuyToken   common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	FeeAmount  *big.Int
	OrderUid   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_Settlement *SettlementFilterer) FilterTrade(opts *bind.FilterOpts, owner []common.Address) (*SettlementTradeIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.FilterLogs(opts, "Trade", ownerRule)
	if err != nil {
		return nil, err
	}
	return &SettlementTradeIterator{contract: _Settlement.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_Settlement *SettlementFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *SettlementTrade, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Settlement.contract.WatchLogs(opts, "Trade", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettlementTrade)
				if err := _Settlement.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xa07a543ab8a018198e99ca0184c93fe9050a79400a0a723441f84de1d972cc17.
//
// Solidity: event Trade(address indexed owner, address sellToken, address buyToken, uint256 sellAmount, uint256 buyAmount, uint256 feeAmount, bytes orderUid)
func (_Settlement *SettlementFilterer) ParseTrade(log types.Log) (*SettlementTrade, error) {
	event := new(SettlementTrade)
	if err := _Settlement.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
