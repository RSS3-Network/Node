// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polymarket

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

// ConditionTokensMetaData contains all meta data concerning the ConditionTokens contract.
var ConditionTokensMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"indexSets\",\"type\":\"uint256[]\"}],\"name\":\"redeemPositions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"payoutNumerators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"ids\",\"type\":\"uint256[]\"},{\"name\":\"values\",\"type\":\"uint256[]\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"collectionId\",\"type\":\"bytes32\"}],\"name\":\"getPositionId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owners\",\"type\":\"address[]\"},{\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"partition\",\"type\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"splitPosition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\"},{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"indexSet\",\"type\":\"uint256\"}],\"name\":\"getCollectionId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\"},{\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"name\":\"partition\",\"type\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mergePositions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"payouts\",\"type\":\"uint256[]\"}],\"name\":\"reportPayouts\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"getOutcomeSlotCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\"},{\"name\":\"questionId\",\"type\":\"bytes32\"},{\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"prepareCondition\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"payoutDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"id\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"}],\"name\":\"ConditionPreparation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"questionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"outcomeSlotCount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"payoutNumerators\",\"type\":\"uint256[]\"}],\"name\":\"ConditionResolution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionSplit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"stakeholder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"partition\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PositionsMerge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"collateralToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"parentCollectionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"conditionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"indexSets\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"payout\",\"type\":\"uint256\"}],\"name\":\"PayoutRedemption\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"}]",
}

// ConditionTokensABI is the input ABI used to generate the binding from.
// Deprecated: Use ConditionTokensMetaData.ABI instead.
var ConditionTokensABI = ConditionTokensMetaData.ABI

// ConditionTokens is an auto generated Go binding around an Ethereum contract.
type ConditionTokens struct {
	ConditionTokensCaller     // Read-only binding to the contract
	ConditionTokensTransactor // Write-only binding to the contract
	ConditionTokensFilterer   // Log filterer for contract events
}

// ConditionTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConditionTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConditionTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConditionTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConditionTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConditionTokensSession struct {
	Contract     *ConditionTokens  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConditionTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConditionTokensCallerSession struct {
	Contract *ConditionTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ConditionTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConditionTokensTransactorSession struct {
	Contract     *ConditionTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ConditionTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConditionTokensRaw struct {
	Contract *ConditionTokens // Generic contract binding to access the raw methods on
}

// ConditionTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConditionTokensCallerRaw struct {
	Contract *ConditionTokensCaller // Generic read-only contract binding to access the raw methods on
}

// ConditionTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConditionTokensTransactorRaw struct {
	Contract *ConditionTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConditionTokens creates a new instance of ConditionTokens, bound to a specific deployed contract.
func NewConditionTokens(address common.Address, backend bind.ContractBackend) (*ConditionTokens, error) {
	contract, err := bindConditionTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConditionTokens{ConditionTokensCaller: ConditionTokensCaller{contract: contract}, ConditionTokensTransactor: ConditionTokensTransactor{contract: contract}, ConditionTokensFilterer: ConditionTokensFilterer{contract: contract}}, nil
}

// NewConditionTokensCaller creates a new read-only instance of ConditionTokens, bound to a specific deployed contract.
func NewConditionTokensCaller(address common.Address, caller bind.ContractCaller) (*ConditionTokensCaller, error) {
	contract, err := bindConditionTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensCaller{contract: contract}, nil
}

// NewConditionTokensTransactor creates a new write-only instance of ConditionTokens, bound to a specific deployed contract.
func NewConditionTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*ConditionTokensTransactor, error) {
	contract, err := bindConditionTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensTransactor{contract: contract}, nil
}

// NewConditionTokensFilterer creates a new log filterer instance of ConditionTokens, bound to a specific deployed contract.
func NewConditionTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*ConditionTokensFilterer, error) {
	contract, err := bindConditionTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensFilterer{contract: contract}, nil
}

// bindConditionTokens binds a generic wrapper to an already deployed contract.
func bindConditionTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ConditionTokensMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionTokens *ConditionTokensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionTokens.Contract.ConditionTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionTokens *ConditionTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionTokens.Contract.ConditionTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionTokens *ConditionTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionTokens.Contract.ConditionTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConditionTokens *ConditionTokensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConditionTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConditionTokens *ConditionTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConditionTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConditionTokens *ConditionTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConditionTokens.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionTokens *ConditionTokensCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionTokens *ConditionTokensSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ConditionTokens.Contract.BalanceOf(&_ConditionTokens.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256)
func (_ConditionTokens *ConditionTokensCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _ConditionTokens.Contract.BalanceOf(&_ConditionTokens.CallOpts, owner, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionTokens *ConditionTokensCaller) BalanceOfBatch(opts *bind.CallOpts, owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "balanceOfBatch", owners, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionTokens *ConditionTokensSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ConditionTokens.Contract.BalanceOfBatch(&_ConditionTokens.CallOpts, owners, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] owners, uint256[] ids) view returns(uint256[])
func (_ConditionTokens *ConditionTokensCallerSession) BalanceOfBatch(owners []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ConditionTokens.Contract.BalanceOfBatch(&_ConditionTokens.CallOpts, owners, ids)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionTokens *ConditionTokensCaller) GetCollectionId(opts *bind.CallOpts, parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "getCollectionId", parentCollectionId, conditionId, indexSet)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionTokens *ConditionTokensSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _ConditionTokens.Contract.GetCollectionId(&_ConditionTokens.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetCollectionId is a free data retrieval call binding the contract method 0x856296f7.
//
// Solidity: function getCollectionId(bytes32 parentCollectionId, bytes32 conditionId, uint256 indexSet) view returns(bytes32)
func (_ConditionTokens *ConditionTokensCallerSession) GetCollectionId(parentCollectionId [32]byte, conditionId [32]byte, indexSet *big.Int) ([32]byte, error) {
	return _ConditionTokens.Contract.GetCollectionId(&_ConditionTokens.CallOpts, parentCollectionId, conditionId, indexSet)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionTokens *ConditionTokensCaller) GetConditionId(opts *bind.CallOpts, oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "getConditionId", oracle, questionId, outcomeSlotCount)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionTokens *ConditionTokensSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _ConditionTokens.Contract.GetConditionId(&_ConditionTokens.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetConditionId is a free data retrieval call binding the contract method 0x852c6ae2.
//
// Solidity: function getConditionId(address oracle, bytes32 questionId, uint256 outcomeSlotCount) pure returns(bytes32)
func (_ConditionTokens *ConditionTokensCallerSession) GetConditionId(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) ([32]byte, error) {
	return _ConditionTokens.Contract.GetConditionId(&_ConditionTokens.CallOpts, oracle, questionId, outcomeSlotCount)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionTokens *ConditionTokensCaller) GetOutcomeSlotCount(opts *bind.CallOpts, conditionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "getOutcomeSlotCount", conditionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionTokens *ConditionTokensSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.GetOutcomeSlotCount(&_ConditionTokens.CallOpts, conditionId)
}

// GetOutcomeSlotCount is a free data retrieval call binding the contract method 0xd42dc0c2.
//
// Solidity: function getOutcomeSlotCount(bytes32 conditionId) view returns(uint256)
func (_ConditionTokens *ConditionTokensCallerSession) GetOutcomeSlotCount(conditionId [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.GetOutcomeSlotCount(&_ConditionTokens.CallOpts, conditionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionTokens *ConditionTokensCaller) GetPositionId(opts *bind.CallOpts, collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "getPositionId", collateralToken, collectionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionTokens *ConditionTokensSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.GetPositionId(&_ConditionTokens.CallOpts, collateralToken, collectionId)
}

// GetPositionId is a free data retrieval call binding the contract method 0x39dd7530.
//
// Solidity: function getPositionId(address collateralToken, bytes32 collectionId) pure returns(uint256)
func (_ConditionTokens *ConditionTokensCallerSession) GetPositionId(collateralToken common.Address, collectionId [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.GetPositionId(&_ConditionTokens.CallOpts, collateralToken, collectionId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionTokens *ConditionTokensCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionTokens *ConditionTokensSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ConditionTokens.Contract.IsApprovedForAll(&_ConditionTokens.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ConditionTokens *ConditionTokensCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ConditionTokens.Contract.IsApprovedForAll(&_ConditionTokens.CallOpts, owner, operator)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensCaller) PayoutDenominator(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "payoutDenominator", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.PayoutDenominator(&_ConditionTokens.CallOpts, arg0)
}

// PayoutDenominator is a free data retrieval call binding the contract method 0xdd34de67.
//
// Solidity: function payoutDenominator(bytes32 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensCallerSession) PayoutDenominator(arg0 [32]byte) (*big.Int, error) {
	return _ConditionTokens.Contract.PayoutDenominator(&_ConditionTokens.CallOpts, arg0)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensCaller) PayoutNumerators(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "payoutNumerators", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _ConditionTokens.Contract.PayoutNumerators(&_ConditionTokens.CallOpts, arg0, arg1)
}

// PayoutNumerators is a free data retrieval call binding the contract method 0x0504c814.
//
// Solidity: function payoutNumerators(bytes32 , uint256 ) view returns(uint256)
func (_ConditionTokens *ConditionTokensCallerSession) PayoutNumerators(arg0 [32]byte, arg1 *big.Int) (*big.Int, error) {
	return _ConditionTokens.Contract.PayoutNumerators(&_ConditionTokens.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionTokens *ConditionTokensCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ConditionTokens.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionTokens *ConditionTokensSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConditionTokens.Contract.SupportsInterface(&_ConditionTokens.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ConditionTokens *ConditionTokensCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ConditionTokens.Contract.SupportsInterface(&_ConditionTokens.CallOpts, interfaceId)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensTransactor) MergePositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "mergePositions", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.MergePositions(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// MergePositions is a paid mutator transaction binding the contract method 0x9e7212ad.
//
// Solidity: function mergePositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) MergePositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.MergePositions(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionTokens *ConditionTokensTransactor) PrepareCondition(opts *bind.TransactOpts, oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "prepareCondition", oracle, questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionTokens *ConditionTokensSession) PrepareCondition(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.PrepareCondition(&_ConditionTokens.TransactOpts, oracle, questionId, outcomeSlotCount)
}

// PrepareCondition is a paid mutator transaction binding the contract method 0xd96ee754.
//
// Solidity: function prepareCondition(address oracle, bytes32 questionId, uint256 outcomeSlotCount) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) PrepareCondition(oracle common.Address, questionId [32]byte, outcomeSlotCount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.PrepareCondition(&_ConditionTokens.TransactOpts, oracle, questionId, outcomeSlotCount)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionTokens *ConditionTokensTransactor) RedeemPositions(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "redeemPositions", collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionTokens *ConditionTokensSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.RedeemPositions(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// RedeemPositions is a paid mutator transaction binding the contract method 0x01b7037c.
//
// Solidity: function redeemPositions(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] indexSets) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) RedeemPositions(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, indexSets []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.RedeemPositions(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, indexSets)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionTokens *ConditionTokensTransactor) ReportPayouts(opts *bind.TransactOpts, questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "reportPayouts", questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionTokens *ConditionTokensSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.ReportPayouts(&_ConditionTokens.TransactOpts, questionId, payouts)
}

// ReportPayouts is a paid mutator transaction binding the contract method 0xc49298ac.
//
// Solidity: function reportPayouts(bytes32 questionId, uint256[] payouts) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) ReportPayouts(questionId [32]byte, payouts []*big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.ReportPayouts(&_ConditionTokens.TransactOpts, questionId, payouts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionTokens *ConditionTokensTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionTokens *ConditionTokensSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SafeBatchTransferFrom(&_ConditionTokens.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SafeBatchTransferFrom(&_ConditionTokens.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionTokens *ConditionTokensTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionTokens *ConditionTokensSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SafeTransferFrom(&_ConditionTokens.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SafeTransferFrom(&_ConditionTokens.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionTokens *ConditionTokensTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionTokens *ConditionTokensSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SetApprovalForAll(&_ConditionTokens.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SetApprovalForAll(&_ConditionTokens.TransactOpts, operator, approved)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensTransactor) SplitPosition(opts *bind.TransactOpts, collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.contract.Transact(opts, "splitPosition", collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SplitPosition(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// SplitPosition is a paid mutator transaction binding the contract method 0x72ce4275.
//
// Solidity: function splitPosition(address collateralToken, bytes32 parentCollectionId, bytes32 conditionId, uint256[] partition, uint256 amount) returns()
func (_ConditionTokens *ConditionTokensTransactorSession) SplitPosition(collateralToken common.Address, parentCollectionId [32]byte, conditionId [32]byte, partition []*big.Int, amount *big.Int) (*types.Transaction, error) {
	return _ConditionTokens.Contract.SplitPosition(&_ConditionTokens.TransactOpts, collateralToken, parentCollectionId, conditionId, partition, amount)
}

// ConditionTokensApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ConditionTokens contract.
type ConditionTokensApprovalForAllIterator struct {
	Event *ConditionTokensApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ConditionTokensApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensApprovalForAll)
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
		it.Event = new(ConditionTokensApprovalForAll)
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
func (it *ConditionTokensApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensApprovalForAll represents a ApprovalForAll event raised by the ConditionTokens contract.
type ConditionTokensApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionTokens *ConditionTokensFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ConditionTokensApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensApprovalForAllIterator{contract: _ConditionTokens.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionTokens *ConditionTokensFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ConditionTokensApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensApprovalForAll)
				if err := _ConditionTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ConditionTokens *ConditionTokensFilterer) ParseApprovalForAll(log types.Log) (*ConditionTokensApprovalForAll, error) {
	event := new(ConditionTokensApprovalForAll)
	if err := _ConditionTokens.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensConditionPreparationIterator is returned from FilterConditionPreparation and is used to iterate over the raw logs and unpacked data for ConditionPreparation events raised by the ConditionTokens contract.
type ConditionTokensConditionPreparationIterator struct {
	Event *ConditionTokensConditionPreparation // Event containing the contract specifics and raw log

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
func (it *ConditionTokensConditionPreparationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensConditionPreparation)
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
		it.Event = new(ConditionTokensConditionPreparation)
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
func (it *ConditionTokensConditionPreparationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensConditionPreparationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensConditionPreparation represents a ConditionPreparation event raised by the ConditionTokens contract.
type ConditionTokensConditionPreparation struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionPreparation is a free log retrieval operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionTokens *ConditionTokensFilterer) FilterConditionPreparation(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*ConditionTokensConditionPreparationIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensConditionPreparationIterator{contract: _ConditionTokens.contract, event: "ConditionPreparation", logs: logs, sub: sub}, nil
}

// WatchConditionPreparation is a free log subscription operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionTokens *ConditionTokensFilterer) WatchConditionPreparation(opts *bind.WatchOpts, sink chan<- *ConditionTokensConditionPreparation, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "ConditionPreparation", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensConditionPreparation)
				if err := _ConditionTokens.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
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

// ParseConditionPreparation is a log parse operation binding the contract event 0xab3760c3bd2bb38b5bcf54dc79802ed67338b4cf29f3054ded67ed24661e4177.
//
// Solidity: event ConditionPreparation(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount)
func (_ConditionTokens *ConditionTokensFilterer) ParseConditionPreparation(log types.Log) (*ConditionTokensConditionPreparation, error) {
	event := new(ConditionTokensConditionPreparation)
	if err := _ConditionTokens.contract.UnpackLog(event, "ConditionPreparation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensConditionResolutionIterator is returned from FilterConditionResolution and is used to iterate over the raw logs and unpacked data for ConditionResolution events raised by the ConditionTokens contract.
type ConditionTokensConditionResolutionIterator struct {
	Event *ConditionTokensConditionResolution // Event containing the contract specifics and raw log

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
func (it *ConditionTokensConditionResolutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensConditionResolution)
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
		it.Event = new(ConditionTokensConditionResolution)
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
func (it *ConditionTokensConditionResolutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensConditionResolutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensConditionResolution represents a ConditionResolution event raised by the ConditionTokens contract.
type ConditionTokensConditionResolution struct {
	ConditionId      [32]byte
	Oracle           common.Address
	QuestionId       [32]byte
	OutcomeSlotCount *big.Int
	PayoutNumerators []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterConditionResolution is a free log retrieval operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionTokens *ConditionTokensFilterer) FilterConditionResolution(opts *bind.FilterOpts, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (*ConditionTokensConditionResolutionIterator, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensConditionResolutionIterator{contract: _ConditionTokens.contract, event: "ConditionResolution", logs: logs, sub: sub}, nil
}

// WatchConditionResolution is a free log subscription operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionTokens *ConditionTokensFilterer) WatchConditionResolution(opts *bind.WatchOpts, sink chan<- *ConditionTokensConditionResolution, conditionId [][32]byte, oracle []common.Address, questionId [][32]byte) (event.Subscription, error) {

	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var questionIdRule []interface{}
	for _, questionIdItem := range questionId {
		questionIdRule = append(questionIdRule, questionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "ConditionResolution", conditionIdRule, oracleRule, questionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensConditionResolution)
				if err := _ConditionTokens.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
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

// ParseConditionResolution is a log parse operation binding the contract event 0xb44d84d3289691f71497564b85d4233648d9dbae8cbdbb4329f301c3a0185894.
//
// Solidity: event ConditionResolution(bytes32 indexed conditionId, address indexed oracle, bytes32 indexed questionId, uint256 outcomeSlotCount, uint256[] payoutNumerators)
func (_ConditionTokens *ConditionTokensFilterer) ParseConditionResolution(log types.Log) (*ConditionTokensConditionResolution, error) {
	event := new(ConditionTokensConditionResolution)
	if err := _ConditionTokens.contract.UnpackLog(event, "ConditionResolution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensPayoutRedemptionIterator is returned from FilterPayoutRedemption and is used to iterate over the raw logs and unpacked data for PayoutRedemption events raised by the ConditionTokens contract.
type ConditionTokensPayoutRedemptionIterator struct {
	Event *ConditionTokensPayoutRedemption // Event containing the contract specifics and raw log

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
func (it *ConditionTokensPayoutRedemptionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensPayoutRedemption)
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
		it.Event = new(ConditionTokensPayoutRedemption)
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
func (it *ConditionTokensPayoutRedemptionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensPayoutRedemptionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensPayoutRedemption represents a PayoutRedemption event raised by the ConditionTokens contract.
type ConditionTokensPayoutRedemption struct {
	Redeemer           common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	IndexSets          []*big.Int
	Payout             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPayoutRedemption is a free log retrieval operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionTokens *ConditionTokensFilterer) FilterPayoutRedemption(opts *bind.FilterOpts, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (*ConditionTokensPayoutRedemptionIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensPayoutRedemptionIterator{contract: _ConditionTokens.contract, event: "PayoutRedemption", logs: logs, sub: sub}, nil
}

// WatchPayoutRedemption is a free log subscription operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionTokens *ConditionTokensFilterer) WatchPayoutRedemption(opts *bind.WatchOpts, sink chan<- *ConditionTokensPayoutRedemption, redeemer []common.Address, collateralToken []common.Address, parentCollectionId [][32]byte) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var collateralTokenRule []interface{}
	for _, collateralTokenItem := range collateralToken {
		collateralTokenRule = append(collateralTokenRule, collateralTokenItem)
	}
	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "PayoutRedemption", redeemerRule, collateralTokenRule, parentCollectionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensPayoutRedemption)
				if err := _ConditionTokens.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
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

// ParsePayoutRedemption is a log parse operation binding the contract event 0x2682012a4a4f1973119f1c9b90745d1bd91fa2bab387344f044cb3586864d18d.
//
// Solidity: event PayoutRedemption(address indexed redeemer, address indexed collateralToken, bytes32 indexed parentCollectionId, bytes32 conditionId, uint256[] indexSets, uint256 payout)
func (_ConditionTokens *ConditionTokensFilterer) ParsePayoutRedemption(log types.Log) (*ConditionTokensPayoutRedemption, error) {
	event := new(ConditionTokensPayoutRedemption)
	if err := _ConditionTokens.contract.UnpackLog(event, "PayoutRedemption", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensPositionSplitIterator is returned from FilterPositionSplit and is used to iterate over the raw logs and unpacked data for PositionSplit events raised by the ConditionTokens contract.
type ConditionTokensPositionSplitIterator struct {
	Event *ConditionTokensPositionSplit // Event containing the contract specifics and raw log

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
func (it *ConditionTokensPositionSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensPositionSplit)
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
		it.Event = new(ConditionTokensPositionSplit)
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
func (it *ConditionTokensPositionSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensPositionSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensPositionSplit represents a PositionSplit event raised by the ConditionTokens contract.
type ConditionTokensPositionSplit struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionSplit is a free log retrieval operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) FilterPositionSplit(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*ConditionTokensPositionSplitIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensPositionSplitIterator{contract: _ConditionTokens.contract, event: "PositionSplit", logs: logs, sub: sub}, nil
}

// WatchPositionSplit is a free log subscription operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) WatchPositionSplit(opts *bind.WatchOpts, sink chan<- *ConditionTokensPositionSplit, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "PositionSplit", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensPositionSplit)
				if err := _ConditionTokens.contract.UnpackLog(event, "PositionSplit", log); err != nil {
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

// ParsePositionSplit is a log parse operation binding the contract event 0x2e6bb91f8cbcda0c93623c54d0403a43514fabc40084ec96b6d5379a74786298.
//
// Solidity: event PositionSplit(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) ParsePositionSplit(log types.Log) (*ConditionTokensPositionSplit, error) {
	event := new(ConditionTokensPositionSplit)
	if err := _ConditionTokens.contract.UnpackLog(event, "PositionSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensPositionsMergeIterator is returned from FilterPositionsMerge and is used to iterate over the raw logs and unpacked data for PositionsMerge events raised by the ConditionTokens contract.
type ConditionTokensPositionsMergeIterator struct {
	Event *ConditionTokensPositionsMerge // Event containing the contract specifics and raw log

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
func (it *ConditionTokensPositionsMergeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensPositionsMerge)
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
		it.Event = new(ConditionTokensPositionsMerge)
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
func (it *ConditionTokensPositionsMergeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensPositionsMergeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensPositionsMerge represents a PositionsMerge event raised by the ConditionTokens contract.
type ConditionTokensPositionsMerge struct {
	Stakeholder        common.Address
	CollateralToken    common.Address
	ParentCollectionId [32]byte
	ConditionId        [32]byte
	Partition          []*big.Int
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPositionsMerge is a free log retrieval operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) FilterPositionsMerge(opts *bind.FilterOpts, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (*ConditionTokensPositionsMergeIterator, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensPositionsMergeIterator{contract: _ConditionTokens.contract, event: "PositionsMerge", logs: logs, sub: sub}, nil
}

// WatchPositionsMerge is a free log subscription operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) WatchPositionsMerge(opts *bind.WatchOpts, sink chan<- *ConditionTokensPositionsMerge, stakeholder []common.Address, parentCollectionId [][32]byte, conditionId [][32]byte) (event.Subscription, error) {

	var stakeholderRule []interface{}
	for _, stakeholderItem := range stakeholder {
		stakeholderRule = append(stakeholderRule, stakeholderItem)
	}

	var parentCollectionIdRule []interface{}
	for _, parentCollectionIdItem := range parentCollectionId {
		parentCollectionIdRule = append(parentCollectionIdRule, parentCollectionIdItem)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "PositionsMerge", stakeholderRule, parentCollectionIdRule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensPositionsMerge)
				if err := _ConditionTokens.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
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

// ParsePositionsMerge is a log parse operation binding the contract event 0x6f13ca62553fcc2bcd2372180a43949c1e4cebba603901ede2f4e14f36b282ca.
//
// Solidity: event PositionsMerge(address indexed stakeholder, address collateralToken, bytes32 indexed parentCollectionId, bytes32 indexed conditionId, uint256[] partition, uint256 amount)
func (_ConditionTokens *ConditionTokensFilterer) ParsePositionsMerge(log types.Log) (*ConditionTokensPositionsMerge, error) {
	event := new(ConditionTokensPositionsMerge)
	if err := _ConditionTokens.contract.UnpackLog(event, "PositionsMerge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ConditionTokens contract.
type ConditionTokensTransferBatchIterator struct {
	Event *ConditionTokensTransferBatch // Event containing the contract specifics and raw log

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
func (it *ConditionTokensTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensTransferBatch)
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
		it.Event = new(ConditionTokensTransferBatch)
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
func (it *ConditionTokensTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensTransferBatch represents a TransferBatch event raised by the ConditionTokens contract.
type ConditionTokensTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionTokens *ConditionTokensFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ConditionTokensTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensTransferBatchIterator{contract: _ConditionTokens.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionTokens *ConditionTokensFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ConditionTokensTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensTransferBatch)
				if err := _ConditionTokens.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_ConditionTokens *ConditionTokensFilterer) ParseTransferBatch(log types.Log) (*ConditionTokensTransferBatch, error) {
	event := new(ConditionTokensTransferBatch)
	if err := _ConditionTokens.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the ConditionTokens contract.
type ConditionTokensTransferSingleIterator struct {
	Event *ConditionTokensTransferSingle // Event containing the contract specifics and raw log

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
func (it *ConditionTokensTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensTransferSingle)
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
		it.Event = new(ConditionTokensTransferSingle)
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
func (it *ConditionTokensTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensTransferSingle represents a TransferSingle event raised by the ConditionTokens contract.
type ConditionTokensTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionTokens *ConditionTokensFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*ConditionTokensTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensTransferSingleIterator{contract: _ConditionTokens.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionTokens *ConditionTokensFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *ConditionTokensTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensTransferSingle)
				if err := _ConditionTokens.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_ConditionTokens *ConditionTokensFilterer) ParseTransferSingle(log types.Log) (*ConditionTokensTransferSingle, error) {
	event := new(ConditionTokensTransferSingle)
	if err := _ConditionTokens.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConditionTokensURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the ConditionTokens contract.
type ConditionTokensURIIterator struct {
	Event *ConditionTokensURI // Event containing the contract specifics and raw log

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
func (it *ConditionTokensURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConditionTokensURI)
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
		it.Event = new(ConditionTokensURI)
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
func (it *ConditionTokensURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConditionTokensURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConditionTokensURI represents a URI event raised by the ConditionTokens contract.
type ConditionTokensURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionTokens *ConditionTokensFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*ConditionTokensURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConditionTokens.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &ConditionTokensURIIterator{contract: _ConditionTokens.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionTokens *ConditionTokensFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *ConditionTokensURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConditionTokens.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConditionTokensURI)
				if err := _ConditionTokens.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_ConditionTokens *ConditionTokensFilterer) ParseURI(log types.Log) (*ConditionTokensURI, error) {
	event := new(ConditionTokensURI)
	if err := _ConditionTokens.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
