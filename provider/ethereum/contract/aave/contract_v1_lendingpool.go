// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// V1LendingPoolMetaData contains all meta data concerning the V1LendingPool contract.
var V1LendingPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_originationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accruedBorrowInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"OriginationFeeLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RedeemUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountMinusFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"LENDINGPOOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"internalType\":\"contractLendingPoolCore\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolDataProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsStable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parametersProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolParametersProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aTokenBalanceAfterRedeem\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V1LendingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use V1LendingPoolMetaData.ABI instead.
var V1LendingPoolABI = V1LendingPoolMetaData.ABI

// V1LendingPool is an auto generated Go binding around an Ethereum contract.
type V1LendingPool struct {
	V1LendingPoolCaller     // Read-only binding to the contract
	V1LendingPoolTransactor // Write-only binding to the contract
	V1LendingPoolFilterer   // Log filterer for contract events
}

// V1LendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1LendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1LendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1LendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1LendingPoolSession struct {
	Contract     *V1LendingPool    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1LendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1LendingPoolCallerSession struct {
	Contract *V1LendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// V1LendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1LendingPoolTransactorSession struct {
	Contract     *V1LendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// V1LendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1LendingPoolRaw struct {
	Contract *V1LendingPool // Generic contract binding to access the raw methods on
}

// V1LendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1LendingPoolCallerRaw struct {
	Contract *V1LendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// V1LendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1LendingPoolTransactorRaw struct {
	Contract *V1LendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1LendingPool creates a new instance of V1LendingPool, bound to a specific deployed contract.
func NewV1LendingPool(address common.Address, backend bind.ContractBackend) (*V1LendingPool, error) {
	contract, err := bindV1LendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1LendingPool{V1LendingPoolCaller: V1LendingPoolCaller{contract: contract}, V1LendingPoolTransactor: V1LendingPoolTransactor{contract: contract}, V1LendingPoolFilterer: V1LendingPoolFilterer{contract: contract}}, nil
}

// NewV1LendingPoolCaller creates a new read-only instance of V1LendingPool, bound to a specific deployed contract.
func NewV1LendingPoolCaller(address common.Address, caller bind.ContractCaller) (*V1LendingPoolCaller, error) {
	contract, err := bindV1LendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolCaller{contract: contract}, nil
}

// NewV1LendingPoolTransactor creates a new write-only instance of V1LendingPool, bound to a specific deployed contract.
func NewV1LendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*V1LendingPoolTransactor, error) {
	contract, err := bindV1LendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolTransactor{contract: contract}, nil
}

// NewV1LendingPoolFilterer creates a new log filterer instance of V1LendingPool, bound to a specific deployed contract.
func NewV1LendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*V1LendingPoolFilterer, error) {
	contract, err := bindV1LendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolFilterer{contract: contract}, nil
}

// bindV1LendingPool binds a generic wrapper to an already deployed contract.
func bindV1LendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V1LendingPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1LendingPool *V1LendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1LendingPool.Contract.V1LendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1LendingPool *V1LendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1LendingPool.Contract.V1LendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1LendingPool *V1LendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1LendingPool.Contract.V1LendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1LendingPool *V1LendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1LendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1LendingPool *V1LendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1LendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1LendingPool *V1LendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1LendingPool.Contract.contract.Transact(opts, method, params...)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_V1LendingPool *V1LendingPoolCaller) LENDINGPOOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "LENDINGPOOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_V1LendingPool *V1LendingPoolSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _V1LendingPool.Contract.LENDINGPOOLREVISION(&_V1LendingPool.CallOpts)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_V1LendingPool *V1LendingPoolCallerSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _V1LendingPool.Contract.LENDINGPOOLREVISION(&_V1LendingPool.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_V1LendingPool *V1LendingPoolCaller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "UINT_MAX_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_V1LendingPool *V1LendingPoolSession) UINTMAXVALUE() (*big.Int, error) {
	return _V1LendingPool.Contract.UINTMAXVALUE(&_V1LendingPool.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_V1LendingPool *V1LendingPoolCallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _V1LendingPool.Contract.UINTMAXVALUE(&_V1LendingPool.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolSession) AddressesProvider() (common.Address, error) {
	return _V1LendingPool.Contract.AddressesProvider(&_V1LendingPool.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCallerSession) AddressesProvider() (common.Address, error) {
	return _V1LendingPool.Contract.AddressesProvider(&_V1LendingPool.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_V1LendingPool *V1LendingPoolCaller) Core(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "core")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_V1LendingPool *V1LendingPoolSession) Core() (common.Address, error) {
	return _V1LendingPool.Contract.Core(&_V1LendingPool.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_V1LendingPool *V1LendingPoolCallerSession) Core() (common.Address, error) {
	return _V1LendingPool.Contract.Core(&_V1LendingPool.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCaller) DataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "dataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolSession) DataProvider() (common.Address, error) {
	return _V1LendingPool.Contract.DataProvider(&_V1LendingPool.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCallerSession) DataProvider() (common.Address, error) {
	return _V1LendingPool.Contract.DataProvider(&_V1LendingPool.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_V1LendingPool *V1LendingPoolCaller) GetReserveConfigurationData(opts *bind.CallOpts, _reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "getReserveConfigurationData", _reserve)

	outstruct := new(struct {
		Ltv                         *big.Int
		LiquidationThreshold        *big.Int
		LiquidationBonus            *big.Int
		InterestRateStrategyAddress common.Address
		UsageAsCollateralEnabled    bool
		BorrowingEnabled            bool
		StableBorrowRateEnabled     bool
		IsActive                    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ltv = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.InterestRateStrategyAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_V1LendingPool *V1LendingPoolSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _V1LendingPool.Contract.GetReserveConfigurationData(&_V1LendingPool.CallOpts, _reserve)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_V1LendingPool *V1LendingPoolCallerSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _V1LendingPool.Contract.GetReserveConfigurationData(&_V1LendingPool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_V1LendingPool *V1LendingPoolCaller) GetReserveData(opts *bind.CallOpts, _reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "getReserveData", _reserve)

	outstruct := new(struct {
		TotalLiquidity          *big.Int
		AvailableLiquidity      *big.Int
		TotalBorrowsStable      *big.Int
		TotalBorrowsVariable    *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		UtilizationRate         *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		ATokenAddress           common.Address
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalLiquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AvailableLiquidity = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsStable = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsVariable = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UtilizationRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.ATokenAddress = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[12], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_V1LendingPool *V1LendingPoolSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _V1LendingPool.Contract.GetReserveData(&_V1LendingPool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_V1LendingPool *V1LendingPoolCallerSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _V1LendingPool.Contract.GetReserveData(&_V1LendingPool.CallOpts, _reserve)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_V1LendingPool *V1LendingPoolCaller) GetReserves(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_V1LendingPool *V1LendingPoolSession) GetReserves() ([]common.Address, error) {
	return _V1LendingPool.Contract.GetReserves(&_V1LendingPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_V1LendingPool *V1LendingPoolCallerSession) GetReserves() ([]common.Address, error) {
	return _V1LendingPool.Contract.GetReserves(&_V1LendingPool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_V1LendingPool *V1LendingPoolCaller) GetUserAccountData(opts *bind.CallOpts, _user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "getUserAccountData", _user)

	outstruct := new(struct {
		TotalLiquidityETH           *big.Int
		TotalCollateralETH          *big.Int
		TotalBorrowsETH             *big.Int
		TotalFeesETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalLiquidityETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCollateralETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalBorrowsETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalFeesETH = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsETH = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_V1LendingPool *V1LendingPoolSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _V1LendingPool.Contract.GetUserAccountData(&_V1LendingPool.CallOpts, _user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_V1LendingPool *V1LendingPoolCallerSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _V1LendingPool.Contract.GetUserAccountData(&_V1LendingPool.CallOpts, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_V1LendingPool *V1LendingPoolCaller) GetUserReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "getUserReserveData", _reserve, _user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentBorrowBalance     *big.Int
		PrincipalBorrowBalance   *big.Int
		BorrowRateMode           *big.Int
		BorrowRate               *big.Int
		LiquidityRate            *big.Int
		OriginationFee           *big.Int
		VariableBorrowIndex      *big.Int
		LastUpdateTimestamp      *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentBorrowBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PrincipalBorrowBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BorrowRateMode = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BorrowRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OriginationFee = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_V1LendingPool *V1LendingPoolSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _V1LendingPool.Contract.GetUserReserveData(&_V1LendingPool.CallOpts, _reserve, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_V1LendingPool *V1LendingPoolCallerSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _V1LendingPool.Contract.GetUserReserveData(&_V1LendingPool.CallOpts, _reserve, _user)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCaller) ParametersProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LendingPool.contract.Call(opts, &out, "parametersProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolSession) ParametersProvider() (common.Address, error) {
	return _V1LendingPool.Contract.ParametersProvider(&_V1LendingPool.CallOpts)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_V1LendingPool *V1LendingPoolCallerSession) ParametersProvider() (common.Address, error) {
	return _V1LendingPool.Contract.ParametersProvider(&_V1LendingPool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_V1LendingPool *V1LendingPoolTransactor) Borrow(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "borrow", _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_V1LendingPool *V1LendingPoolSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Borrow(&_V1LendingPool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Borrow(&_V1LendingPool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_V1LendingPool *V1LendingPoolTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "deposit", _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_V1LendingPool *V1LendingPoolSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Deposit(&_V1LendingPool.TransactOpts, _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Deposit(&_V1LendingPool.TransactOpts, _reserve, _amount, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_V1LendingPool *V1LendingPoolTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "flashLoan", _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_V1LendingPool *V1LendingPoolSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _V1LendingPool.Contract.FlashLoan(&_V1LendingPool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _V1LendingPool.Contract.FlashLoan(&_V1LendingPool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_V1LendingPool *V1LendingPoolTransactor) Initialize(opts *bind.TransactOpts, _addressesProvider common.Address) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "initialize", _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_V1LendingPool *V1LendingPoolSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Initialize(&_V1LendingPool.TransactOpts, _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Initialize(&_V1LendingPool.TransactOpts, _addressesProvider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_V1LendingPool *V1LendingPoolTransactor) LiquidationCall(opts *bind.TransactOpts, _collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "liquidationCall", _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_V1LendingPool *V1LendingPoolSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _V1LendingPool.Contract.LiquidationCall(&_V1LendingPool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _V1LendingPool.Contract.LiquidationCall(&_V1LendingPool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_V1LendingPool *V1LendingPoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "rebalanceStableBorrowRate", _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_V1LendingPool *V1LendingPoolSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.RebalanceStableBorrowRate(&_V1LendingPool.TransactOpts, _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.RebalanceStableBorrowRate(&_V1LendingPool.TransactOpts, _reserve, _user)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_V1LendingPool *V1LendingPoolTransactor) RedeemUnderlying(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "redeemUnderlying", _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_V1LendingPool *V1LendingPoolSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _V1LendingPool.Contract.RedeemUnderlying(&_V1LendingPool.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _V1LendingPool.Contract.RedeemUnderlying(&_V1LendingPool.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_V1LendingPool *V1LendingPoolTransactor) Repay(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "repay", _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_V1LendingPool *V1LendingPoolSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Repay(&_V1LendingPool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.Repay(&_V1LendingPool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_V1LendingPool *V1LendingPoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_V1LendingPool *V1LendingPoolSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _V1LendingPool.Contract.SetUserUseReserveAsCollateral(&_V1LendingPool.TransactOpts, _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _V1LendingPool.Contract.SetUserUseReserveAsCollateral(&_V1LendingPool.TransactOpts, _reserve, _useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_V1LendingPool *V1LendingPoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _V1LendingPool.contract.Transact(opts, "swapBorrowRateMode", _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_V1LendingPool *V1LendingPoolSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.SwapBorrowRateMode(&_V1LendingPool.TransactOpts, _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_V1LendingPool *V1LendingPoolTransactorSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _V1LendingPool.Contract.SwapBorrowRateMode(&_V1LendingPool.TransactOpts, _reserve)
}

// V1LendingPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the V1LendingPool contract.
type V1LendingPoolBorrowIterator struct {
	Event *V1LendingPoolBorrow // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolBorrow)
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
		it.Event = new(V1LendingPoolBorrow)
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
func (it *V1LendingPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolBorrow represents a Borrow event raised by the V1LendingPool contract.
type V1LendingPoolBorrow struct {
	Reserve               common.Address
	User                  common.Address
	Amount                *big.Int
	BorrowRateMode        *big.Int
	BorrowRate            *big.Int
	OriginationFee        *big.Int
	BorrowBalanceIncrease *big.Int
	Referral              uint16
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterBorrow(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*V1LendingPoolBorrowIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolBorrowIterator{contract: _V1LendingPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *V1LendingPoolBorrow, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolBorrow)
				if err := _V1LendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseBorrow(log types.Log) (*V1LendingPoolBorrow, error) {
	event := new(V1LendingPoolBorrow)
	if err := _V1LendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the V1LendingPool contract.
type V1LendingPoolDepositIterator struct {
	Event *V1LendingPoolDeposit // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolDeposit)
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
		it.Event = new(V1LendingPoolDeposit)
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
func (it *V1LendingPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolDeposit represents a Deposit event raised by the V1LendingPool contract.
type V1LendingPoolDeposit struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Referral  uint16
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterDeposit(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*V1LendingPoolDepositIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolDepositIterator{contract: _V1LendingPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *V1LendingPoolDeposit, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolDeposit)
				if err := _V1LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseDeposit(log types.Log) (*V1LendingPoolDeposit, error) {
	event := new(V1LendingPoolDeposit)
	if err := _V1LendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the V1LendingPool contract.
type V1LendingPoolFlashLoanIterator struct {
	Event *V1LendingPoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolFlashLoan)
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
		it.Event = new(V1LendingPoolFlashLoan)
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
func (it *V1LendingPoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolFlashLoan represents a FlashLoan event raised by the V1LendingPool contract.
type V1LendingPoolFlashLoan struct {
	Target      common.Address
	Reserve     common.Address
	Amount      *big.Int
	TotalFee    *big.Int
	ProtocolFee *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, _target []common.Address, _reserve []common.Address) (*V1LendingPoolFlashLoanIterator, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolFlashLoanIterator{contract: _V1LendingPool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *V1LendingPoolFlashLoan, _target []common.Address, _reserve []common.Address) (event.Subscription, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolFlashLoan)
				if err := _V1LendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseFlashLoan(log types.Log) (*V1LendingPoolFlashLoan, error) {
	event := new(V1LendingPoolFlashLoan)
	if err := _V1LendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the V1LendingPool contract.
type V1LendingPoolLiquidationCallIterator struct {
	Event *V1LendingPoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolLiquidationCall)
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
		it.Event = new(V1LendingPoolLiquidationCall)
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
func (it *V1LendingPoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolLiquidationCall represents a LiquidationCall event raised by the V1LendingPool contract.
type V1LendingPoolLiquidationCall struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	PurchaseAmount             *big.Int
	LiquidatedCollateralAmount *big.Int
	AccruedBorrowInterest      *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*V1LendingPoolLiquidationCallIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolLiquidationCallIterator{contract: _V1LendingPool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *V1LendingPoolLiquidationCall, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolLiquidationCall)
				if err := _V1LendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseLiquidationCall(log types.Log) (*V1LendingPoolLiquidationCall, error) {
	event := new(V1LendingPoolLiquidationCall)
	if err := _V1LendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolOriginationFeeLiquidatedIterator is returned from FilterOriginationFeeLiquidated and is used to iterate over the raw logs and unpacked data for OriginationFeeLiquidated events raised by the V1LendingPool contract.
type V1LendingPoolOriginationFeeLiquidatedIterator struct {
	Event *V1LendingPoolOriginationFeeLiquidated // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolOriginationFeeLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolOriginationFeeLiquidated)
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
		it.Event = new(V1LendingPoolOriginationFeeLiquidated)
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
func (it *V1LendingPoolOriginationFeeLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolOriginationFeeLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolOriginationFeeLiquidated represents a OriginationFeeLiquidated event raised by the V1LendingPool contract.
type V1LendingPoolOriginationFeeLiquidated struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	FeeLiquidated              *big.Int
	LiquidatedCollateralForFee *big.Int
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterOriginationFeeLiquidated is a free log retrieval operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterOriginationFeeLiquidated(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*V1LendingPoolOriginationFeeLiquidatedIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolOriginationFeeLiquidatedIterator{contract: _V1LendingPool.contract, event: "OriginationFeeLiquidated", logs: logs, sub: sub}, nil
}

// WatchOriginationFeeLiquidated is a free log subscription operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchOriginationFeeLiquidated(opts *bind.WatchOpts, sink chan<- *V1LendingPoolOriginationFeeLiquidated, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolOriginationFeeLiquidated)
				if err := _V1LendingPool.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
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

// ParseOriginationFeeLiquidated is a log parse operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseOriginationFeeLiquidated(log types.Log) (*V1LendingPoolOriginationFeeLiquidated, error) {
	event := new(V1LendingPoolOriginationFeeLiquidated)
	if err := _V1LendingPool.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the V1LendingPool contract.
type V1LendingPoolRebalanceStableBorrowRateIterator struct {
	Event *V1LendingPoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolRebalanceStableBorrowRate)
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
		it.Event = new(V1LendingPoolRebalanceStableBorrowRate)
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
func (it *V1LendingPoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the V1LendingPool contract.
type V1LendingPoolRebalanceStableBorrowRate struct {
	Reserve               common.Address
	User                  common.Address
	NewStableRate         *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*V1LendingPoolRebalanceStableBorrowRateIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolRebalanceStableBorrowRateIterator{contract: _V1LendingPool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *V1LendingPoolRebalanceStableBorrowRate, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolRebalanceStableBorrowRate)
				if err := _V1LendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*V1LendingPoolRebalanceStableBorrowRate, error) {
	event := new(V1LendingPoolRebalanceStableBorrowRate)
	if err := _V1LendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolRedeemUnderlyingIterator is returned from FilterRedeemUnderlying and is used to iterate over the raw logs and unpacked data for RedeemUnderlying events raised by the V1LendingPool contract.
type V1LendingPoolRedeemUnderlyingIterator struct {
	Event *V1LendingPoolRedeemUnderlying // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolRedeemUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolRedeemUnderlying)
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
		it.Event = new(V1LendingPoolRedeemUnderlying)
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
func (it *V1LendingPoolRedeemUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolRedeemUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolRedeemUnderlying represents a RedeemUnderlying event raised by the V1LendingPool contract.
type V1LendingPoolRedeemUnderlying struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemUnderlying is a free log retrieval operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterRedeemUnderlying(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*V1LendingPoolRedeemUnderlyingIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolRedeemUnderlyingIterator{contract: _V1LendingPool.contract, event: "RedeemUnderlying", logs: logs, sub: sub}, nil
}

// WatchRedeemUnderlying is a free log subscription operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchRedeemUnderlying(opts *bind.WatchOpts, sink chan<- *V1LendingPoolRedeemUnderlying, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolRedeemUnderlying)
				if err := _V1LendingPool.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
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

// ParseRedeemUnderlying is a log parse operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseRedeemUnderlying(log types.Log) (*V1LendingPoolRedeemUnderlying, error) {
	event := new(V1LendingPoolRedeemUnderlying)
	if err := _V1LendingPool.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the V1LendingPool contract.
type V1LendingPoolRepayIterator struct {
	Event *V1LendingPoolRepay // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolRepay)
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
		it.Event = new(V1LendingPoolRepay)
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
func (it *V1LendingPoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolRepay represents a Repay event raised by the V1LendingPool contract.
type V1LendingPoolRepay struct {
	Reserve               common.Address
	User                  common.Address
	Repayer               common.Address
	AmountMinusFees       *big.Int
	Fees                  *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterRepay(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (*V1LendingPoolRepayIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolRepayIterator{contract: _V1LendingPool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *V1LendingPoolRepay, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolRepay)
				if err := _V1LendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseRepay(log types.Log) (*V1LendingPoolRepay, error) {
	event := new(V1LendingPoolRepay)
	if err := _V1LendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the V1LendingPool contract.
type V1LendingPoolReserveUsedAsCollateralDisabledIterator struct {
	Event *V1LendingPoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(V1LendingPoolReserveUsedAsCollateralDisabled)
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
func (it *V1LendingPoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the V1LendingPool contract.
type V1LendingPoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*V1LendingPoolReserveUsedAsCollateralDisabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolReserveUsedAsCollateralDisabledIterator{contract: _V1LendingPool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *V1LendingPoolReserveUsedAsCollateralDisabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolReserveUsedAsCollateralDisabled)
				if err := _V1LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*V1LendingPoolReserveUsedAsCollateralDisabled, error) {
	event := new(V1LendingPoolReserveUsedAsCollateralDisabled)
	if err := _V1LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the V1LendingPool contract.
type V1LendingPoolReserveUsedAsCollateralEnabledIterator struct {
	Event *V1LendingPoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(V1LendingPoolReserveUsedAsCollateralEnabled)
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
func (it *V1LendingPoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the V1LendingPool contract.
type V1LendingPoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*V1LendingPoolReserveUsedAsCollateralEnabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolReserveUsedAsCollateralEnabledIterator{contract: _V1LendingPool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *V1LendingPoolReserveUsedAsCollateralEnabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolReserveUsedAsCollateralEnabled)
				if err := _V1LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_V1LendingPool *V1LendingPoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*V1LendingPoolReserveUsedAsCollateralEnabled, error) {
	event := new(V1LendingPoolReserveUsedAsCollateralEnabled)
	if err := _V1LendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LendingPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the V1LendingPool contract.
type V1LendingPoolSwapIterator struct {
	Event *V1LendingPoolSwap // Event containing the contract specifics and raw log

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
func (it *V1LendingPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LendingPoolSwap)
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
		it.Event = new(V1LendingPoolSwap)
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
func (it *V1LendingPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LendingPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LendingPoolSwap represents a Swap event raised by the V1LendingPool contract.
type V1LendingPoolSwap struct {
	Reserve               common.Address
	User                  common.Address
	NewRateMode           *big.Int
	NewRate               *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) FilterSwap(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*V1LendingPoolSwapIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.FilterLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &V1LendingPoolSwapIterator{contract: _V1LendingPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *V1LendingPoolSwap, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _V1LendingPool.contract.WatchLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LendingPoolSwap)
				if err := _V1LendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_V1LendingPool *V1LendingPoolFilterer) ParseSwap(log types.Log) (*V1LendingPoolSwap, error) {
	event := new(V1LendingPoolSwap)
	if err := _V1LendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
