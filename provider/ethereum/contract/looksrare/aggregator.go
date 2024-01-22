// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package looksrare

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

// BasicOrder is an auto generated low-level Go binding around an user-defined struct.
type BasicOrder struct {
	Signer         common.Address
	Collection     common.Address
	CollectionType uint8
	TokenIds       []*big.Int
	Amounts        []*big.Int
	Price          *big.Int
	Currency       common.Address
	StartTime      *big.Int
	EndTime        *big.Int
	Signature      []byte
}

// ILooksRareAggregatorTradeData is an auto generated low-level Go binding around an user-defined struct.
type ILooksRareAggregatorTradeData struct {
	Proxy           common.Address
	Selector        [4]byte
	Orders          []BasicOrder
	OrdersExtraData [][]byte
	ExtraData       []byte
}

// TokenTransfer is an auto generated low-level Go binding around an user-defined struct.
type TokenTransfer struct {
	Amount   *big.Int
	Currency common.Address
}

// AggregatorMetaData contains all meta data concerning the Aggregator contract.
var AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1155SafeBatchTransferFromFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20ApprovalFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC20TransferFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721TransferFromFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOrderLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoOngoingTransferInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RenouncementNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferAlreadyInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferNotInProgress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UseERC20EnabledLooksRareAggregator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPotentialOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"CancelOwnershipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"FunctionAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"FunctionRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"InitiateOwnershipRenouncement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"potentialOwner\",\"type\":\"address\"}],\"name\":\"InitiateOwnershipTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"NewOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sweeper\",\"type\":\"address\"}],\"name\":\"Sweep\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"addFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplace\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOwnershipRenouncement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"erc20EnabledLooksRareAggregator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structTokenTransfer[]\",\"name\":\"tokenTransfers\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"enumCollectionType\",\"name\":\"collectionType\",\"type\":\"uint8\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structBasicOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"ordersExtraData\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structILooksRareAggregator.TradeData[]\",\"name\":\"tradeData\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"originator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAtomic\",\"type\":\"bool\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initiateOwnershipRenouncement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPotentialOwner\",\"type\":\"address\"}],\"name\":\"initiateOwnershipTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownershipStatus\",\"outputs\":[{\"internalType\":\"enumIOwnableTwoSteps.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"removeFunction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"rescueERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"rescueERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20EnabledLooksRareAggregator\",\"type\":\"address\"}],\"name\":\"setERC20EnabledLooksRareAggregator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"supportsProxyFunction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorMetaData.ABI instead.
var AggregatorABI = AggregatorMetaData.ABI

// Aggregator is an auto generated Go binding around an Ethereum contract.
type Aggregator struct {
	AggregatorCaller     // Read-only binding to the contract
	AggregatorTransactor // Write-only binding to the contract
	AggregatorFilterer   // Log filterer for contract events
}

// AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorSession struct {
	Contract     *Aggregator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorCallerSession struct {
	Contract *AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorTransactorSession struct {
	Contract     *AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorRaw struct {
	Contract *Aggregator // Generic contract binding to access the raw methods on
}

// AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorCallerRaw struct {
	Contract *AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorTransactorRaw struct {
	Contract *AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregator creates a new instance of Aggregator, bound to a specific deployed contract.
func NewAggregator(address common.Address, backend bind.ContractBackend) (*Aggregator, error) {
	contract, err := bindAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// NewAggregatorCaller creates a new read-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AggregatorCaller, error) {
	contract, err := bindAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorCaller{contract: contract}, nil
}

// NewAggregatorTransactor creates a new write-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorTransactor, error) {
	contract, err := bindAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorTransactor{contract: contract}, nil
}

// NewAggregatorFilterer creates a new log filterer instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorFilterer, error) {
	contract, err := bindAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorFilterer{contract: contract}, nil
}

// bindAggregator binds a generic wrapper to an already deployed contract.
func bindAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transact(opts, method, params...)
}

// Erc20EnabledLooksRareAggregator is a free data retrieval call binding the contract method 0x9b110618.
//
// Solidity: function erc20EnabledLooksRareAggregator() view returns(address)
func (_Aggregator *AggregatorCaller) Erc20EnabledLooksRareAggregator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "erc20EnabledLooksRareAggregator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20EnabledLooksRareAggregator is a free data retrieval call binding the contract method 0x9b110618.
//
// Solidity: function erc20EnabledLooksRareAggregator() view returns(address)
func (_Aggregator *AggregatorSession) Erc20EnabledLooksRareAggregator() (common.Address, error) {
	return _Aggregator.Contract.Erc20EnabledLooksRareAggregator(&_Aggregator.CallOpts)
}

// Erc20EnabledLooksRareAggregator is a free data retrieval call binding the contract method 0x9b110618.
//
// Solidity: function erc20EnabledLooksRareAggregator() view returns(address)
func (_Aggregator *AggregatorCallerSession) Erc20EnabledLooksRareAggregator() (common.Address, error) {
	return _Aggregator.Contract.Erc20EnabledLooksRareAggregator(&_Aggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCallerSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_Aggregator *AggregatorCaller) OwnershipStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "ownershipStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_Aggregator *AggregatorSession) OwnershipStatus() (uint8, error) {
	return _Aggregator.Contract.OwnershipStatus(&_Aggregator.CallOpts)
}

// OwnershipStatus is a free data retrieval call binding the contract method 0x2bb5a9e6.
//
// Solidity: function ownershipStatus() view returns(uint8)
func (_Aggregator *AggregatorCallerSession) OwnershipStatus() (uint8, error) {
	return _Aggregator.Contract.OwnershipStatus(&_Aggregator.CallOpts)
}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_Aggregator *AggregatorCaller) PotentialOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "potentialOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_Aggregator *AggregatorSession) PotentialOwner() (common.Address, error) {
	return _Aggregator.Contract.PotentialOwner(&_Aggregator.CallOpts)
}

// PotentialOwner is a free data retrieval call binding the contract method 0x7762df25.
//
// Solidity: function potentialOwner() view returns(address)
func (_Aggregator *AggregatorCallerSession) PotentialOwner() (common.Address, error) {
	return _Aggregator.Contract.PotentialOwner(&_Aggregator.CallOpts)
}

// SupportsProxyFunction is a free data retrieval call binding the contract method 0x7bcf8bba.
//
// Solidity: function supportsProxyFunction(address proxy, bytes4 selector) view returns(bool isSupported)
func (_Aggregator *AggregatorCaller) SupportsProxyFunction(opts *bind.CallOpts, proxy common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "supportsProxyFunction", proxy, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsProxyFunction is a free data retrieval call binding the contract method 0x7bcf8bba.
//
// Solidity: function supportsProxyFunction(address proxy, bytes4 selector) view returns(bool isSupported)
func (_Aggregator *AggregatorSession) SupportsProxyFunction(proxy common.Address, selector [4]byte) (bool, error) {
	return _Aggregator.Contract.SupportsProxyFunction(&_Aggregator.CallOpts, proxy, selector)
}

// SupportsProxyFunction is a free data retrieval call binding the contract method 0x7bcf8bba.
//
// Solidity: function supportsProxyFunction(address proxy, bytes4 selector) view returns(bool isSupported)
func (_Aggregator *AggregatorCallerSession) SupportsProxyFunction(proxy common.Address, selector [4]byte) (bool, error) {
	return _Aggregator.Contract.SupportsProxyFunction(&_Aggregator.CallOpts, proxy, selector)
}

// AddFunction is a paid mutator transaction binding the contract method 0x0c9a50b3.
//
// Solidity: function addFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorTransactor) AddFunction(opts *bind.TransactOpts, proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "addFunction", proxy, selector)
}

// AddFunction is a paid mutator transaction binding the contract method 0x0c9a50b3.
//
// Solidity: function addFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorSession) AddFunction(proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.AddFunction(&_Aggregator.TransactOpts, proxy, selector)
}

// AddFunction is a paid mutator transaction binding the contract method 0x0c9a50b3.
//
// Solidity: function addFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorTransactorSession) AddFunction(proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.AddFunction(&_Aggregator.TransactOpts, proxy, selector)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address currency, address marketplace, uint256 amount) returns()
func (_Aggregator *AggregatorTransactor) Approve(opts *bind.TransactOpts, currency common.Address, marketplace common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "approve", currency, marketplace, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address currency, address marketplace, uint256 amount) returns()
func (_Aggregator *AggregatorSession) Approve(currency common.Address, marketplace common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Approve(&_Aggregator.TransactOpts, currency, marketplace, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address currency, address marketplace, uint256 amount) returns()
func (_Aggregator *AggregatorTransactorSession) Approve(currency common.Address, marketplace common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Approve(&_Aggregator.TransactOpts, currency, marketplace, amount)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Aggregator *AggregatorTransactor) CancelOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "cancelOwnershipTransfer")
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Aggregator *AggregatorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Aggregator.Contract.CancelOwnershipTransfer(&_Aggregator.TransactOpts)
}

// CancelOwnershipTransfer is a paid mutator transaction binding the contract method 0x23452b9c.
//
// Solidity: function cancelOwnershipTransfer() returns()
func (_Aggregator *AggregatorTransactorSession) CancelOwnershipTransfer() (*types.Transaction, error) {
	return _Aggregator.Contract.CancelOwnershipTransfer(&_Aggregator.TransactOpts)
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_Aggregator *AggregatorTransactor) ConfirmOwnershipRenouncement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "confirmOwnershipRenouncement")
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_Aggregator *AggregatorSession) ConfirmOwnershipRenouncement() (*types.Transaction, error) {
	return _Aggregator.Contract.ConfirmOwnershipRenouncement(&_Aggregator.TransactOpts)
}

// ConfirmOwnershipRenouncement is a paid mutator transaction binding the contract method 0x3e567539.
//
// Solidity: function confirmOwnershipRenouncement() returns()
func (_Aggregator *AggregatorTransactorSession) ConfirmOwnershipRenouncement() (*types.Transaction, error) {
	return _Aggregator.Contract.ConfirmOwnershipRenouncement(&_Aggregator.TransactOpts)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_Aggregator *AggregatorTransactor) ConfirmOwnershipTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "confirmOwnershipTransfer")
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_Aggregator *AggregatorSession) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _Aggregator.Contract.ConfirmOwnershipTransfer(&_Aggregator.TransactOpts)
}

// ConfirmOwnershipTransfer is a paid mutator transaction binding the contract method 0x7200b829.
//
// Solidity: function confirmOwnershipTransfer() returns()
func (_Aggregator *AggregatorTransactorSession) ConfirmOwnershipTransfer() (*types.Transaction, error) {
	return _Aggregator.Contract.ConfirmOwnershipTransfer(&_Aggregator.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x12abee12.
//
// Solidity: function execute((uint256,address)[] tokenTransfers, (address,bytes4,(address,address,uint8,uint256[],uint256[],uint256,address,uint256,uint256,bytes)[],bytes[],bytes)[] tradeData, address originator, address recipient, bool isAtomic) payable returns()
func (_Aggregator *AggregatorTransactor) Execute(opts *bind.TransactOpts, tokenTransfers []TokenTransfer, tradeData []ILooksRareAggregatorTradeData, originator common.Address, recipient common.Address, isAtomic bool) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "execute", tokenTransfers, tradeData, originator, recipient, isAtomic)
}

// Execute is a paid mutator transaction binding the contract method 0x12abee12.
//
// Solidity: function execute((uint256,address)[] tokenTransfers, (address,bytes4,(address,address,uint8,uint256[],uint256[],uint256,address,uint256,uint256,bytes)[],bytes[],bytes)[] tradeData, address originator, address recipient, bool isAtomic) payable returns()
func (_Aggregator *AggregatorSession) Execute(tokenTransfers []TokenTransfer, tradeData []ILooksRareAggregatorTradeData, originator common.Address, recipient common.Address, isAtomic bool) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute(&_Aggregator.TransactOpts, tokenTransfers, tradeData, originator, recipient, isAtomic)
}

// Execute is a paid mutator transaction binding the contract method 0x12abee12.
//
// Solidity: function execute((uint256,address)[] tokenTransfers, (address,bytes4,(address,address,uint8,uint256[],uint256[],uint256,address,uint256,uint256,bytes)[],bytes[],bytes)[] tradeData, address originator, address recipient, bool isAtomic) payable returns()
func (_Aggregator *AggregatorTransactorSession) Execute(tokenTransfers []TokenTransfer, tradeData []ILooksRareAggregatorTradeData, originator common.Address, recipient common.Address, isAtomic bool) (*types.Transaction, error) {
	return _Aggregator.Contract.Execute(&_Aggregator.TransactOpts, tokenTransfers, tradeData, originator, recipient, isAtomic)
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_Aggregator *AggregatorTransactor) InitiateOwnershipRenouncement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "initiateOwnershipRenouncement")
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_Aggregator *AggregatorSession) InitiateOwnershipRenouncement() (*types.Transaction, error) {
	return _Aggregator.Contract.InitiateOwnershipRenouncement(&_Aggregator.TransactOpts)
}

// InitiateOwnershipRenouncement is a paid mutator transaction binding the contract method 0x5b6ac011.
//
// Solidity: function initiateOwnershipRenouncement() returns()
func (_Aggregator *AggregatorTransactorSession) InitiateOwnershipRenouncement() (*types.Transaction, error) {
	return _Aggregator.Contract.InitiateOwnershipRenouncement(&_Aggregator.TransactOpts)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_Aggregator *AggregatorTransactor) InitiateOwnershipTransfer(opts *bind.TransactOpts, newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "initiateOwnershipTransfer", newPotentialOwner)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_Aggregator *AggregatorSession) InitiateOwnershipTransfer(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.InitiateOwnershipTransfer(&_Aggregator.TransactOpts, newPotentialOwner)
}

// InitiateOwnershipTransfer is a paid mutator transaction binding the contract method 0xc0b6f561.
//
// Solidity: function initiateOwnershipTransfer(address newPotentialOwner) returns()
func (_Aggregator *AggregatorTransactorSession) InitiateOwnershipTransfer(newPotentialOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.InitiateOwnershipTransfer(&_Aggregator.TransactOpts, newPotentialOwner)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Aggregator *AggregatorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC1155BatchReceived(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC1155BatchReceived(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC1155Received(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC1155Received(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC721Received(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Aggregator *AggregatorTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnERC721Received(&_Aggregator.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveFunction is a paid mutator transaction binding the contract method 0xb9495800.
//
// Solidity: function removeFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorTransactor) RemoveFunction(opts *bind.TransactOpts, proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "removeFunction", proxy, selector)
}

// RemoveFunction is a paid mutator transaction binding the contract method 0xb9495800.
//
// Solidity: function removeFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorSession) RemoveFunction(proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.RemoveFunction(&_Aggregator.TransactOpts, proxy, selector)
}

// RemoveFunction is a paid mutator transaction binding the contract method 0xb9495800.
//
// Solidity: function removeFunction(address proxy, bytes4 selector) returns()
func (_Aggregator *AggregatorTransactorSession) RemoveFunction(proxy common.Address, selector [4]byte) (*types.Transaction, error) {
	return _Aggregator.Contract.RemoveFunction(&_Aggregator.TransactOpts, proxy, selector)
}

// RescueERC1155 is a paid mutator transaction binding the contract method 0xedf9d5c8.
//
// Solidity: function rescueERC1155(address collection, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Aggregator *AggregatorTransactor) RescueERC1155(opts *bind.TransactOpts, collection common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "rescueERC1155", collection, to, tokenIds, amounts)
}

// RescueERC1155 is a paid mutator transaction binding the contract method 0xedf9d5c8.
//
// Solidity: function rescueERC1155(address collection, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Aggregator *AggregatorSession) RescueERC1155(collection common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.RescueERC1155(&_Aggregator.TransactOpts, collection, to, tokenIds, amounts)
}

// RescueERC1155 is a paid mutator transaction binding the contract method 0xedf9d5c8.
//
// Solidity: function rescueERC1155(address collection, address to, uint256[] tokenIds, uint256[] amounts) returns()
func (_Aggregator *AggregatorTransactorSession) RescueERC1155(collection common.Address, to common.Address, tokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.RescueERC1155(&_Aggregator.TransactOpts, collection, to, tokenIds, amounts)
}

// RescueERC721 is a paid mutator transaction binding the contract method 0x7df325e1.
//
// Solidity: function rescueERC721(address collection, address to, uint256 tokenId) returns()
func (_Aggregator *AggregatorTransactor) RescueERC721(opts *bind.TransactOpts, collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "rescueERC721", collection, to, tokenId)
}

// RescueERC721 is a paid mutator transaction binding the contract method 0x7df325e1.
//
// Solidity: function rescueERC721(address collection, address to, uint256 tokenId) returns()
func (_Aggregator *AggregatorSession) RescueERC721(collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.RescueERC721(&_Aggregator.TransactOpts, collection, to, tokenId)
}

// RescueERC721 is a paid mutator transaction binding the contract method 0x7df325e1.
//
// Solidity: function rescueERC721(address collection, address to, uint256 tokenId) returns()
func (_Aggregator *AggregatorTransactorSession) RescueERC721(collection common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.RescueERC721(&_Aggregator.TransactOpts, collection, to, tokenId)
}

// SetERC20EnabledLooksRareAggregator is a paid mutator transaction binding the contract method 0xca53d0b7.
//
// Solidity: function setERC20EnabledLooksRareAggregator(address _erc20EnabledLooksRareAggregator) returns()
func (_Aggregator *AggregatorTransactor) SetERC20EnabledLooksRareAggregator(opts *bind.TransactOpts, _erc20EnabledLooksRareAggregator common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setERC20EnabledLooksRareAggregator", _erc20EnabledLooksRareAggregator)
}

// SetERC20EnabledLooksRareAggregator is a paid mutator transaction binding the contract method 0xca53d0b7.
//
// Solidity: function setERC20EnabledLooksRareAggregator(address _erc20EnabledLooksRareAggregator) returns()
func (_Aggregator *AggregatorSession) SetERC20EnabledLooksRareAggregator(_erc20EnabledLooksRareAggregator common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetERC20EnabledLooksRareAggregator(&_Aggregator.TransactOpts, _erc20EnabledLooksRareAggregator)
}

// SetERC20EnabledLooksRareAggregator is a paid mutator transaction binding the contract method 0xca53d0b7.
//
// Solidity: function setERC20EnabledLooksRareAggregator(address _erc20EnabledLooksRareAggregator) returns()
func (_Aggregator *AggregatorTransactorSession) SetERC20EnabledLooksRareAggregator(_erc20EnabledLooksRareAggregator common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetERC20EnabledLooksRareAggregator(&_Aggregator.TransactOpts, _erc20EnabledLooksRareAggregator)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregator *AggregatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregator *AggregatorSession) Receive() (*types.Transaction, error) {
	return _Aggregator.Contract.Receive(&_Aggregator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Aggregator *AggregatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Aggregator.Contract.Receive(&_Aggregator.TransactOpts)
}

// AggregatorCancelOwnershipTransferIterator is returned from FilterCancelOwnershipTransfer and is used to iterate over the raw logs and unpacked data for CancelOwnershipTransfer events raised by the Aggregator contract.
type AggregatorCancelOwnershipTransferIterator struct {
	Event *AggregatorCancelOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *AggregatorCancelOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorCancelOwnershipTransfer)
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
		it.Event = new(AggregatorCancelOwnershipTransfer)
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
func (it *AggregatorCancelOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorCancelOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorCancelOwnershipTransfer represents a CancelOwnershipTransfer event raised by the Aggregator contract.
type AggregatorCancelOwnershipTransfer struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCancelOwnershipTransfer is a free log retrieval operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_Aggregator *AggregatorFilterer) FilterCancelOwnershipTransfer(opts *bind.FilterOpts) (*AggregatorCancelOwnershipTransferIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "CancelOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return &AggregatorCancelOwnershipTransferIterator{contract: _Aggregator.contract, event: "CancelOwnershipTransfer", logs: logs, sub: sub}, nil
}

// WatchCancelOwnershipTransfer is a free log subscription operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_Aggregator *AggregatorFilterer) WatchCancelOwnershipTransfer(opts *bind.WatchOpts, sink chan<- *AggregatorCancelOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "CancelOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorCancelOwnershipTransfer)
				if err := _Aggregator.contract.UnpackLog(event, "CancelOwnershipTransfer", log); err != nil {
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

// ParseCancelOwnershipTransfer is a log parse operation binding the contract event 0x8eca980489e87f7dba4f26917aa4bfc906eb3f2b4f7b4b9fd0ff2b8bb3e21ae3.
//
// Solidity: event CancelOwnershipTransfer()
func (_Aggregator *AggregatorFilterer) ParseCancelOwnershipTransfer(log types.Log) (*AggregatorCancelOwnershipTransfer, error) {
	event := new(AggregatorCancelOwnershipTransfer)
	if err := _Aggregator.contract.UnpackLog(event, "CancelOwnershipTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorFunctionAddedIterator is returned from FilterFunctionAdded and is used to iterate over the raw logs and unpacked data for FunctionAdded events raised by the Aggregator contract.
type AggregatorFunctionAddedIterator struct {
	Event *AggregatorFunctionAdded // Event containing the contract specifics and raw log

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
func (it *AggregatorFunctionAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorFunctionAdded)
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
		it.Event = new(AggregatorFunctionAdded)
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
func (it *AggregatorFunctionAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorFunctionAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorFunctionAdded represents a FunctionAdded event raised by the Aggregator contract.
type AggregatorFunctionAdded struct {
	Proxy    common.Address
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFunctionAdded is a free log retrieval operation binding the contract event 0xee43eda4dc4728f3b2c1a9b5dfdda27f5e3345bb35f639415984ffb909f1c953.
//
// Solidity: event FunctionAdded(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) FilterFunctionAdded(opts *bind.FilterOpts) (*AggregatorFunctionAddedIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "FunctionAdded")
	if err != nil {
		return nil, err
	}
	return &AggregatorFunctionAddedIterator{contract: _Aggregator.contract, event: "FunctionAdded", logs: logs, sub: sub}, nil
}

// WatchFunctionAdded is a free log subscription operation binding the contract event 0xee43eda4dc4728f3b2c1a9b5dfdda27f5e3345bb35f639415984ffb909f1c953.
//
// Solidity: event FunctionAdded(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) WatchFunctionAdded(opts *bind.WatchOpts, sink chan<- *AggregatorFunctionAdded) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "FunctionAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorFunctionAdded)
				if err := _Aggregator.contract.UnpackLog(event, "FunctionAdded", log); err != nil {
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

// ParseFunctionAdded is a log parse operation binding the contract event 0xee43eda4dc4728f3b2c1a9b5dfdda27f5e3345bb35f639415984ffb909f1c953.
//
// Solidity: event FunctionAdded(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) ParseFunctionAdded(log types.Log) (*AggregatorFunctionAdded, error) {
	event := new(AggregatorFunctionAdded)
	if err := _Aggregator.contract.UnpackLog(event, "FunctionAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorFunctionRemovedIterator is returned from FilterFunctionRemoved and is used to iterate over the raw logs and unpacked data for FunctionRemoved events raised by the Aggregator contract.
type AggregatorFunctionRemovedIterator struct {
	Event *AggregatorFunctionRemoved // Event containing the contract specifics and raw log

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
func (it *AggregatorFunctionRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorFunctionRemoved)
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
		it.Event = new(AggregatorFunctionRemoved)
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
func (it *AggregatorFunctionRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorFunctionRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorFunctionRemoved represents a FunctionRemoved event raised by the Aggregator contract.
type AggregatorFunctionRemoved struct {
	Proxy    common.Address
	Selector [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFunctionRemoved is a free log retrieval operation binding the contract event 0x5abbfd8975596a613736baed6c3c864ee660459894ad98c7e6492b4c3c40f325.
//
// Solidity: event FunctionRemoved(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) FilterFunctionRemoved(opts *bind.FilterOpts) (*AggregatorFunctionRemovedIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "FunctionRemoved")
	if err != nil {
		return nil, err
	}
	return &AggregatorFunctionRemovedIterator{contract: _Aggregator.contract, event: "FunctionRemoved", logs: logs, sub: sub}, nil
}

// WatchFunctionRemoved is a free log subscription operation binding the contract event 0x5abbfd8975596a613736baed6c3c864ee660459894ad98c7e6492b4c3c40f325.
//
// Solidity: event FunctionRemoved(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) WatchFunctionRemoved(opts *bind.WatchOpts, sink chan<- *AggregatorFunctionRemoved) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "FunctionRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorFunctionRemoved)
				if err := _Aggregator.contract.UnpackLog(event, "FunctionRemoved", log); err != nil {
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

// ParseFunctionRemoved is a log parse operation binding the contract event 0x5abbfd8975596a613736baed6c3c864ee660459894ad98c7e6492b4c3c40f325.
//
// Solidity: event FunctionRemoved(address proxy, bytes4 selector)
func (_Aggregator *AggregatorFilterer) ParseFunctionRemoved(log types.Log) (*AggregatorFunctionRemoved, error) {
	event := new(AggregatorFunctionRemoved)
	if err := _Aggregator.contract.UnpackLog(event, "FunctionRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorInitiateOwnershipRenouncementIterator is returned from FilterInitiateOwnershipRenouncement and is used to iterate over the raw logs and unpacked data for InitiateOwnershipRenouncement events raised by the Aggregator contract.
type AggregatorInitiateOwnershipRenouncementIterator struct {
	Event *AggregatorInitiateOwnershipRenouncement // Event containing the contract specifics and raw log

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
func (it *AggregatorInitiateOwnershipRenouncementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorInitiateOwnershipRenouncement)
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
		it.Event = new(AggregatorInitiateOwnershipRenouncement)
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
func (it *AggregatorInitiateOwnershipRenouncementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorInitiateOwnershipRenouncementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorInitiateOwnershipRenouncement represents a InitiateOwnershipRenouncement event raised by the Aggregator contract.
type AggregatorInitiateOwnershipRenouncement struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterInitiateOwnershipRenouncement is a free log retrieval operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_Aggregator *AggregatorFilterer) FilterInitiateOwnershipRenouncement(opts *bind.FilterOpts) (*AggregatorInitiateOwnershipRenouncementIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "InitiateOwnershipRenouncement")
	if err != nil {
		return nil, err
	}
	return &AggregatorInitiateOwnershipRenouncementIterator{contract: _Aggregator.contract, event: "InitiateOwnershipRenouncement", logs: logs, sub: sub}, nil
}

// WatchInitiateOwnershipRenouncement is a free log subscription operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_Aggregator *AggregatorFilterer) WatchInitiateOwnershipRenouncement(opts *bind.WatchOpts, sink chan<- *AggregatorInitiateOwnershipRenouncement) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "InitiateOwnershipRenouncement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorInitiateOwnershipRenouncement)
				if err := _Aggregator.contract.UnpackLog(event, "InitiateOwnershipRenouncement", log); err != nil {
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

// ParseInitiateOwnershipRenouncement is a log parse operation binding the contract event 0x3ff05a45e46337fa1cbf20996d2eeb927280bce099f37252bcca1040609604ec.
//
// Solidity: event InitiateOwnershipRenouncement()
func (_Aggregator *AggregatorFilterer) ParseInitiateOwnershipRenouncement(log types.Log) (*AggregatorInitiateOwnershipRenouncement, error) {
	event := new(AggregatorInitiateOwnershipRenouncement)
	if err := _Aggregator.contract.UnpackLog(event, "InitiateOwnershipRenouncement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorInitiateOwnershipTransferIterator is returned from FilterInitiateOwnershipTransfer and is used to iterate over the raw logs and unpacked data for InitiateOwnershipTransfer events raised by the Aggregator contract.
type AggregatorInitiateOwnershipTransferIterator struct {
	Event *AggregatorInitiateOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *AggregatorInitiateOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorInitiateOwnershipTransfer)
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
		it.Event = new(AggregatorInitiateOwnershipTransfer)
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
func (it *AggregatorInitiateOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorInitiateOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorInitiateOwnershipTransfer represents a InitiateOwnershipTransfer event raised by the Aggregator contract.
type AggregatorInitiateOwnershipTransfer struct {
	PreviousOwner  common.Address
	PotentialOwner common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitiateOwnershipTransfer is a free log retrieval operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_Aggregator *AggregatorFilterer) FilterInitiateOwnershipTransfer(opts *bind.FilterOpts) (*AggregatorInitiateOwnershipTransferIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "InitiateOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return &AggregatorInitiateOwnershipTransferIterator{contract: _Aggregator.contract, event: "InitiateOwnershipTransfer", logs: logs, sub: sub}, nil
}

// WatchInitiateOwnershipTransfer is a free log subscription operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_Aggregator *AggregatorFilterer) WatchInitiateOwnershipTransfer(opts *bind.WatchOpts, sink chan<- *AggregatorInitiateOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "InitiateOwnershipTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorInitiateOwnershipTransfer)
				if err := _Aggregator.contract.UnpackLog(event, "InitiateOwnershipTransfer", log); err != nil {
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

// ParseInitiateOwnershipTransfer is a log parse operation binding the contract event 0xb86c75c9bffca616b2d314cc914f7c3f1d174255b16b941c3f3ededee276d5ef.
//
// Solidity: event InitiateOwnershipTransfer(address previousOwner, address potentialOwner)
func (_Aggregator *AggregatorFilterer) ParseInitiateOwnershipTransfer(log types.Log) (*AggregatorInitiateOwnershipTransfer, error) {
	event := new(AggregatorInitiateOwnershipTransfer)
	if err := _Aggregator.contract.UnpackLog(event, "InitiateOwnershipTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorNewOwnerIterator is returned from FilterNewOwner and is used to iterate over the raw logs and unpacked data for NewOwner events raised by the Aggregator contract.
type AggregatorNewOwnerIterator struct {
	Event *AggregatorNewOwner // Event containing the contract specifics and raw log

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
func (it *AggregatorNewOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorNewOwner)
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
		it.Event = new(AggregatorNewOwner)
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
func (it *AggregatorNewOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorNewOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorNewOwner represents a NewOwner event raised by the Aggregator contract.
type AggregatorNewOwner struct {
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewOwner is a free log retrieval operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Aggregator *AggregatorFilterer) FilterNewOwner(opts *bind.FilterOpts) (*AggregatorNewOwnerIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return &AggregatorNewOwnerIterator{contract: _Aggregator.contract, event: "NewOwner", logs: logs, sub: sub}, nil
}

// WatchNewOwner is a free log subscription operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Aggregator *AggregatorFilterer) WatchNewOwner(opts *bind.WatchOpts, sink chan<- *AggregatorNewOwner) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "NewOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorNewOwner)
				if err := _Aggregator.contract.UnpackLog(event, "NewOwner", log); err != nil {
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

// ParseNewOwner is a log parse operation binding the contract event 0x3edd90e7770f06fafde38004653b33870066c33bfc923ff6102acd601f85dfbc.
//
// Solidity: event NewOwner(address newOwner)
func (_Aggregator *AggregatorFilterer) ParseNewOwner(log types.Log) (*AggregatorNewOwner, error) {
	event := new(AggregatorNewOwner)
	if err := _Aggregator.contract.UnpackLog(event, "NewOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorSweepIterator is returned from FilterSweep and is used to iterate over the raw logs and unpacked data for Sweep events raised by the Aggregator contract.
type AggregatorSweepIterator struct {
	Event *AggregatorSweep // Event containing the contract specifics and raw log

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
func (it *AggregatorSweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorSweep)
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
		it.Event = new(AggregatorSweep)
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
func (it *AggregatorSweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorSweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorSweep represents a Sweep event raised by the Aggregator contract.
type AggregatorSweep struct {
	Sweeper common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSweep is a free log retrieval operation binding the contract event 0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73.
//
// Solidity: event Sweep(address sweeper)
func (_Aggregator *AggregatorFilterer) FilterSweep(opts *bind.FilterOpts) (*AggregatorSweepIterator, error) {

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "Sweep")
	if err != nil {
		return nil, err
	}
	return &AggregatorSweepIterator{contract: _Aggregator.contract, event: "Sweep", logs: logs, sub: sub}, nil
}

// WatchSweep is a free log subscription operation binding the contract event 0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73.
//
// Solidity: event Sweep(address sweeper)
func (_Aggregator *AggregatorFilterer) WatchSweep(opts *bind.WatchOpts, sink chan<- *AggregatorSweep) (event.Subscription, error) {

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "Sweep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorSweep)
				if err := _Aggregator.contract.UnpackLog(event, "Sweep", log); err != nil {
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

// ParseSweep is a log parse operation binding the contract event 0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73.
//
// Solidity: event Sweep(address sweeper)
func (_Aggregator *AggregatorFilterer) ParseSweep(log types.Log) (*AggregatorSweep, error) {
	event := new(AggregatorSweep)
	if err := _Aggregator.contract.UnpackLog(event, "Sweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
