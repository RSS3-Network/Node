// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paraswap

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

// IBalancerV2VaultBatchSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IBalancerV2VaultBatchSwapStep struct {
	PoolId        [32]byte
	AssetInIndex  *big.Int
	AssetOutIndex *big.Int
	Amount        *big.Int
	UserData      []byte
}

// IBalancerV2VaultFundManagement is an auto generated low-level Go binding around an user-defined struct.
type IBalancerV2VaultFundManagement struct {
	Sender              common.Address
	FromInternalBalance bool
	Recipient           common.Address
	ToInternalBalance   bool
}

// UtilsDirectBalancerV2 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectBalancerV2 struct {
	Swaps          []IBalancerV2VaultBatchSwapStep
	Assets         []common.Address
	Funds          IBalancerV2VaultFundManagement
	Limits         []*big.Int
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	Deadline       *big.Int
	FeePercent     *big.Int
	Vault          common.Address
	Partner        common.Address
	IsApproved     bool
	Beneficiary    common.Address
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectCurveV1 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectCurveV1 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	I              *big.Int
	J              *big.Int
	Partner        common.Address
	IsApproved     bool
	SwapType       uint8
	Beneficiary    common.Address
	NeedWrapNative bool
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectCurveV2 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectCurveV2 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	PoolAddress    common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	I              *big.Int
	J              *big.Int
	Partner        common.Address
	IsApproved     bool
	SwapType       uint8
	Beneficiary    common.Address
	NeedWrapNative bool
	Permit         []byte
	Uuid           [16]byte
}

// UtilsDirectUniV3 is an auto generated low-level Go binding around an user-defined struct.
type UtilsDirectUniV3 struct {
	FromToken      common.Address
	ToToken        common.Address
	Exchange       common.Address
	FromAmount     *big.Int
	ToAmount       *big.Int
	ExpectedAmount *big.Int
	FeePercent     *big.Int
	Deadline       *big.Int
	Partner        common.Address
	IsApproved     bool
	Beneficiary    common.Address
	Path           []byte
	Permit         []byte
	Uuid           [16]byte
}

// V5ParaSwapMetaData contains all meta data concerning the V5ParaSwap contract.
var V5ParaSwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_partnerSharePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxFeePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paraswapReferralShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_paraswapSlippageShare\",\"type\":\"uint256\"},{\"internalType\":\"contractIFeeClaimer\",\"name\":\"_feeClaimer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"BoughtV3\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDirectSwap.DirectSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"SwappedDirect\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receivedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"}],\"name\":\"SwappedV3\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WHITELISTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIBalancerV2Vault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIBalancerV2Vault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectBalancerV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directBalancerV2GivenInSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"poolId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetInIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetOutIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structIBalancerV2Vault.BatchSwapStep[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"fromInternalBalance\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"toInternalBalance\",\"type\":\"bool\"}],\"internalType\":\"structIBalancerV2Vault.FundManagement\",\"name\":\"funds\",\"type\":\"tuple\"},{\"internalType\":\"int256[]\",\"name\":\"limits\",\"type\":\"int256[]\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectBalancerV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directBalancerV2GivenOutSwap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"enumUtils.CurveSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"needWrapNative\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectCurveV1\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directCurveV1Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"j\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"enumUtils.CurveSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"needWrapNative\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectCurveV2\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directCurveV2Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectUniV3\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directUniV3Buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchange\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feePercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes16\",\"name\":\"uuid\",\"type\":\"bytes16\"}],\"internalType\":\"structUtils.DirectUniV3\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"directUniV3Swap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeClaimer\",\"outputs\":[{\"internalType\":\"contractIFeeClaimer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFeePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapReferralShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paraswapSlippageShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"partnerSharePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// V5ParaSwapABI is the input ABI used to generate the binding from.
// Deprecated: Use V5ParaSwapMetaData.ABI instead.
var V5ParaSwapABI = V5ParaSwapMetaData.ABI

// V5ParaSwap is an auto generated Go binding around an Ethereum contract.
type V5ParaSwap struct {
	V5ParaSwapCaller     // Read-only binding to the contract
	V5ParaSwapTransactor // Write-only binding to the contract
	V5ParaSwapFilterer   // Log filterer for contract events
}

// V5ParaSwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type V5ParaSwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V5ParaSwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V5ParaSwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V5ParaSwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V5ParaSwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V5ParaSwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V5ParaSwapSession struct {
	Contract     *V5ParaSwap       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V5ParaSwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V5ParaSwapCallerSession struct {
	Contract *V5ParaSwapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// V5ParaSwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V5ParaSwapTransactorSession struct {
	Contract     *V5ParaSwapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// V5ParaSwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type V5ParaSwapRaw struct {
	Contract *V5ParaSwap // Generic contract binding to access the raw methods on
}

// V5ParaSwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V5ParaSwapCallerRaw struct {
	Contract *V5ParaSwapCaller // Generic read-only contract binding to access the raw methods on
}

// V5ParaSwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V5ParaSwapTransactorRaw struct {
	Contract *V5ParaSwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV5ParaSwap creates a new instance of V5ParaSwap, bound to a specific deployed contract.
func NewV5ParaSwap(address common.Address, backend bind.ContractBackend) (*V5ParaSwap, error) {
	contract, err := bindV5ParaSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwap{V5ParaSwapCaller: V5ParaSwapCaller{contract: contract}, V5ParaSwapTransactor: V5ParaSwapTransactor{contract: contract}, V5ParaSwapFilterer: V5ParaSwapFilterer{contract: contract}}, nil
}

// NewV5ParaSwapCaller creates a new read-only instance of V5ParaSwap, bound to a specific deployed contract.
func NewV5ParaSwapCaller(address common.Address, caller bind.ContractCaller) (*V5ParaSwapCaller, error) {
	contract, err := bindV5ParaSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapCaller{contract: contract}, nil
}

// NewV5ParaSwapTransactor creates a new write-only instance of V5ParaSwap, bound to a specific deployed contract.
func NewV5ParaSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*V5ParaSwapTransactor, error) {
	contract, err := bindV5ParaSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapTransactor{contract: contract}, nil
}

// NewV5ParaSwapFilterer creates a new log filterer instance of V5ParaSwap, bound to a specific deployed contract.
func NewV5ParaSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*V5ParaSwapFilterer, error) {
	contract, err := bindV5ParaSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapFilterer{contract: contract}, nil
}

// bindV5ParaSwap binds a generic wrapper to an already deployed contract.
func bindV5ParaSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V5ParaSwapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V5ParaSwap *V5ParaSwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V5ParaSwap.Contract.V5ParaSwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V5ParaSwap *V5ParaSwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.V5ParaSwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V5ParaSwap *V5ParaSwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.V5ParaSwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V5ParaSwap *V5ParaSwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V5ParaSwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V5ParaSwap *V5ParaSwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V5ParaSwap *V5ParaSwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.contract.Transact(opts, method, params...)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCaller) ROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapSession) ROUTERROLE() ([32]byte, error) {
	return _V5ParaSwap.Contract.ROUTERROLE(&_V5ParaSwap.CallOpts)
}

// ROUTERROLE is a free data retrieval call binding the contract method 0x30d643b5.
//
// Solidity: function ROUTER_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCallerSession) ROUTERROLE() ([32]byte, error) {
	return _V5ParaSwap.Contract.ROUTERROLE(&_V5ParaSwap.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCaller) WHITELISTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "WHITELISTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapSession) WHITELISTEDROLE() ([32]byte, error) {
	return _V5ParaSwap.Contract.WHITELISTEDROLE(&_V5ParaSwap.CallOpts)
}

// WHITELISTEDROLE is a free data retrieval call binding the contract method 0x7a3226ec.
//
// Solidity: function WHITELISTED_ROLE() view returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCallerSession) WHITELISTEDROLE() ([32]byte, error) {
	return _V5ParaSwap.Contract.WHITELISTEDROLE(&_V5ParaSwap.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_V5ParaSwap *V5ParaSwapCaller) FeeClaimer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "feeClaimer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_V5ParaSwap *V5ParaSwapSession) FeeClaimer() (common.Address, error) {
	return _V5ParaSwap.Contract.FeeClaimer(&_V5ParaSwap.CallOpts)
}

// FeeClaimer is a free data retrieval call binding the contract method 0x81cbd3ea.
//
// Solidity: function feeClaimer() view returns(address)
func (_V5ParaSwap *V5ParaSwapCallerSession) FeeClaimer() (common.Address, error) {
	return _V5ParaSwap.Contract.FeeClaimer(&_V5ParaSwap.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCaller) GetKey(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "getKey")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_V5ParaSwap *V5ParaSwapSession) GetKey() ([32]byte, error) {
	return _V5ParaSwap.Contract.GetKey(&_V5ParaSwap.CallOpts)
}

// GetKey is a free data retrieval call binding the contract method 0x82678dd6.
//
// Solidity: function getKey() pure returns(bytes32)
func (_V5ParaSwap *V5ParaSwapCallerSession) GetKey() ([32]byte, error) {
	return _V5ParaSwap.Contract.GetKey(&_V5ParaSwap.CallOpts)
}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_V5ParaSwap *V5ParaSwapCaller) Initialize(opts *bind.CallOpts, arg0 []byte) error {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "initialize", arg0)

	if err != nil {
		return err
	}

	return err

}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_V5ParaSwap *V5ParaSwapSession) Initialize(arg0 []byte) error {
	return _V5ParaSwap.Contract.Initialize(&_V5ParaSwap.CallOpts, arg0)
}

// Initialize is a free data retrieval call binding the contract method 0x439fab91.
//
// Solidity: function initialize(bytes ) pure returns()
func (_V5ParaSwap *V5ParaSwapCallerSession) Initialize(arg0 []byte) error {
	return _V5ParaSwap.Contract.Initialize(&_V5ParaSwap.CallOpts, arg0)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCaller) MaxFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "maxFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapSession) MaxFeePercent() (*big.Int, error) {
	return _V5ParaSwap.Contract.MaxFeePercent(&_V5ParaSwap.CallOpts)
}

// MaxFeePercent is a free data retrieval call binding the contract method 0xd830a05b.
//
// Solidity: function maxFeePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCallerSession) MaxFeePercent() (*big.Int, error) {
	return _V5ParaSwap.Contract.MaxFeePercent(&_V5ParaSwap.CallOpts)
}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCaller) ParaswapReferralShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "paraswapReferralShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapSession) ParaswapReferralShare() (*big.Int, error) {
	return _V5ParaSwap.Contract.ParaswapReferralShare(&_V5ParaSwap.CallOpts)
}

// ParaswapReferralShare is a free data retrieval call binding the contract method 0xd555d4f9.
//
// Solidity: function paraswapReferralShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCallerSession) ParaswapReferralShare() (*big.Int, error) {
	return _V5ParaSwap.Contract.ParaswapReferralShare(&_V5ParaSwap.CallOpts)
}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCaller) ParaswapSlippageShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "paraswapSlippageShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapSession) ParaswapSlippageShare() (*big.Int, error) {
	return _V5ParaSwap.Contract.ParaswapSlippageShare(&_V5ParaSwap.CallOpts)
}

// ParaswapSlippageShare is a free data retrieval call binding the contract method 0xc25ff026.
//
// Solidity: function paraswapSlippageShare() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCallerSession) ParaswapSlippageShare() (*big.Int, error) {
	return _V5ParaSwap.Contract.ParaswapSlippageShare(&_V5ParaSwap.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCaller) PartnerSharePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "partnerSharePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapSession) PartnerSharePercent() (*big.Int, error) {
	return _V5ParaSwap.Contract.PartnerSharePercent(&_V5ParaSwap.CallOpts)
}

// PartnerSharePercent is a free data retrieval call binding the contract method 0x12070a41.
//
// Solidity: function partnerSharePercent() view returns(uint256)
func (_V5ParaSwap *V5ParaSwapCallerSession) PartnerSharePercent() (*big.Int, error) {
	return _V5ParaSwap.Contract.PartnerSharePercent(&_V5ParaSwap.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_V5ParaSwap *V5ParaSwapCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V5ParaSwap.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_V5ParaSwap *V5ParaSwapSession) Weth() (common.Address, error) {
	return _V5ParaSwap.Contract.Weth(&_V5ParaSwap.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_V5ParaSwap *V5ParaSwapCallerSession) Weth() (common.Address, error) {
	return _V5ParaSwap.Contract.Weth(&_V5ParaSwap.CallOpts)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectBalancerV2GivenInSwap(opts *bind.TransactOpts, data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directBalancerV2GivenInSwap", data)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectBalancerV2GivenInSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectBalancerV2GivenInSwap(&_V5ParaSwap.TransactOpts, data)
}

// DirectBalancerV2GivenInSwap is a paid mutator transaction binding the contract method 0xb22f4db8.
//
// Solidity: function directBalancerV2GivenInSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectBalancerV2GivenInSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectBalancerV2GivenInSwap(&_V5ParaSwap.TransactOpts, data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectBalancerV2GivenOutSwap(opts *bind.TransactOpts, data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directBalancerV2GivenOutSwap", data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectBalancerV2GivenOutSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectBalancerV2GivenOutSwap(&_V5ParaSwap.TransactOpts, data)
}

// DirectBalancerV2GivenOutSwap is a paid mutator transaction binding the contract method 0x19fc5be0.
//
// Solidity: function directBalancerV2GivenOutSwap(((bytes32,uint256,uint256,uint256,bytes)[],address[],(address,bool,address,bool),int256[],uint256,uint256,uint256,uint256,uint256,address,address,bool,address,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectBalancerV2GivenOutSwap(data UtilsDirectBalancerV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectBalancerV2GivenOutSwap(&_V5ParaSwap.TransactOpts, data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectCurveV1Swap(opts *bind.TransactOpts, data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directCurveV1Swap", data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectCurveV1Swap(data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectCurveV1Swap(&_V5ParaSwap.TransactOpts, data)
}

// DirectCurveV1Swap is a paid mutator transaction binding the contract method 0x3865bde6.
//
// Solidity: function directCurveV1Swap((address,address,address,uint256,uint256,uint256,uint256,int128,int128,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectCurveV1Swap(data UtilsDirectCurveV1) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectCurveV1Swap(&_V5ParaSwap.TransactOpts, data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectCurveV2Swap(opts *bind.TransactOpts, data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directCurveV2Swap", data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectCurveV2Swap(data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectCurveV2Swap(&_V5ParaSwap.TransactOpts, data)
}

// DirectCurveV2Swap is a paid mutator transaction binding the contract method 0x58f15100.
//
// Solidity: function directCurveV2Swap((address,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,address,bool,uint8,address,bool,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectCurveV2Swap(data UtilsDirectCurveV2) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectCurveV2Swap(&_V5ParaSwap.TransactOpts, data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectUniV3Buy(opts *bind.TransactOpts, data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directUniV3Buy", data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectUniV3Buy(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectUniV3Buy(&_V5ParaSwap.TransactOpts, data)
}

// DirectUniV3Buy is a paid mutator transaction binding the contract method 0x87a63926.
//
// Solidity: function directUniV3Buy((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectUniV3Buy(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectUniV3Buy(&_V5ParaSwap.TransactOpts, data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactor) DirectUniV3Swap(opts *bind.TransactOpts, data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.contract.Transact(opts, "directUniV3Swap", data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapSession) DirectUniV3Swap(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectUniV3Swap(&_V5ParaSwap.TransactOpts, data)
}

// DirectUniV3Swap is a paid mutator transaction binding the contract method 0xa6886da9.
//
// Solidity: function directUniV3Swap((address,address,address,uint256,uint256,uint256,uint256,uint256,address,bool,address,bytes,bytes,bytes16) data) payable returns()
func (_V5ParaSwap *V5ParaSwapTransactorSession) DirectUniV3Swap(data UtilsDirectUniV3) (*types.Transaction, error) {
	return _V5ParaSwap.Contract.DirectUniV3Swap(&_V5ParaSwap.TransactOpts, data)
}

// V5ParaSwapBoughtV3Iterator is returned from FilterBoughtV3 and is used to iterate over the raw logs and unpacked data for BoughtV3 events raised by the V5ParaSwap contract.
type V5ParaSwapBoughtV3Iterator struct {
	Event *V5ParaSwapBoughtV3 // Event containing the contract specifics and raw log

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
func (it *V5ParaSwapBoughtV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V5ParaSwapBoughtV3)
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
		it.Event = new(V5ParaSwapBoughtV3)
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
func (it *V5ParaSwapBoughtV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V5ParaSwapBoughtV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V5ParaSwapBoughtV3 represents a BoughtV3 event raised by the V5ParaSwap contract.
type V5ParaSwapBoughtV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBoughtV3 is a free log retrieval operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) FilterBoughtV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*V5ParaSwapBoughtV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.FilterLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapBoughtV3Iterator{contract: _V5ParaSwap.contract, event: "BoughtV3", logs: logs, sub: sub}, nil
}

// WatchBoughtV3 is a free log subscription operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) WatchBoughtV3(opts *bind.WatchOpts, sink chan<- *V5ParaSwapBoughtV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.WatchLogs(opts, "BoughtV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V5ParaSwapBoughtV3)
				if err := _V5ParaSwap.contract.UnpackLog(event, "BoughtV3", log); err != nil {
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

// ParseBoughtV3 is a log parse operation binding the contract event 0x4cc7e95e48af62690313a0733e93308ac9a73326bc3c29f1788b1191c376d5b6.
//
// Solidity: event BoughtV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) ParseBoughtV3(log types.Log) (*V5ParaSwapBoughtV3, error) {
	event := new(V5ParaSwapBoughtV3)
	if err := _V5ParaSwap.contract.UnpackLog(event, "BoughtV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V5ParaSwapSwappedDirectIterator is returned from FilterSwappedDirect and is used to iterate over the raw logs and unpacked data for SwappedDirect events raised by the V5ParaSwap contract.
type V5ParaSwapSwappedDirectIterator struct {
	Event *V5ParaSwapSwappedDirect // Event containing the contract specifics and raw log

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
func (it *V5ParaSwapSwappedDirectIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V5ParaSwapSwappedDirect)
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
		it.Event = new(V5ParaSwapSwappedDirect)
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
func (it *V5ParaSwapSwappedDirectIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V5ParaSwapSwappedDirectIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V5ParaSwapSwappedDirect represents a SwappedDirect event raised by the V5ParaSwap contract.
type V5ParaSwapSwappedDirect struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Kind           uint8
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwappedDirect is a free log retrieval operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) FilterSwappedDirect(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*V5ParaSwapSwappedDirectIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.FilterLogs(opts, "SwappedDirect", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapSwappedDirectIterator{contract: _V5ParaSwap.contract, event: "SwappedDirect", logs: logs, sub: sub}, nil
}

// WatchSwappedDirect is a free log subscription operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) WatchSwappedDirect(opts *bind.WatchOpts, sink chan<- *V5ParaSwapSwappedDirect, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.WatchLogs(opts, "SwappedDirect", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V5ParaSwapSwappedDirect)
				if err := _V5ParaSwap.contract.UnpackLog(event, "SwappedDirect", log); err != nil {
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

// ParseSwappedDirect is a log parse operation binding the contract event 0xd2d73da2b5fd52cd654d8fd1b514ad57355bad741de639e3a1c3a20dd9f17347.
//
// Solidity: event SwappedDirect(bytes16 uuid, address partner, uint256 feePercent, address initiator, uint8 kind, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) ParseSwappedDirect(log types.Log) (*V5ParaSwapSwappedDirect, error) {
	event := new(V5ParaSwapSwappedDirect)
	if err := _V5ParaSwap.contract.UnpackLog(event, "SwappedDirect", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V5ParaSwapSwappedV3Iterator is returned from FilterSwappedV3 and is used to iterate over the raw logs and unpacked data for SwappedV3 events raised by the V5ParaSwap contract.
type V5ParaSwapSwappedV3Iterator struct {
	Event *V5ParaSwapSwappedV3 // Event containing the contract specifics and raw log

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
func (it *V5ParaSwapSwappedV3Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V5ParaSwapSwappedV3)
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
		it.Event = new(V5ParaSwapSwappedV3)
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
func (it *V5ParaSwapSwappedV3Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V5ParaSwapSwappedV3Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V5ParaSwapSwappedV3 represents a SwappedV3 event raised by the V5ParaSwap contract.
type V5ParaSwapSwappedV3 struct {
	Uuid           [16]byte
	Partner        common.Address
	FeePercent     *big.Int
	Initiator      common.Address
	Beneficiary    common.Address
	SrcToken       common.Address
	DestToken      common.Address
	SrcAmount      *big.Int
	ReceivedAmount *big.Int
	ExpectedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSwappedV3 is a free log retrieval operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) FilterSwappedV3(opts *bind.FilterOpts, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (*V5ParaSwapSwappedV3Iterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.FilterLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return &V5ParaSwapSwappedV3Iterator{contract: _V5ParaSwap.contract, event: "SwappedV3", logs: logs, sub: sub}, nil
}

// WatchSwappedV3 is a free log subscription operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) WatchSwappedV3(opts *bind.WatchOpts, sink chan<- *V5ParaSwapSwappedV3, beneficiary []common.Address, srcToken []common.Address, destToken []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var srcTokenRule []interface{}
	for _, srcTokenItem := range srcToken {
		srcTokenRule = append(srcTokenRule, srcTokenItem)
	}
	var destTokenRule []interface{}
	for _, destTokenItem := range destToken {
		destTokenRule = append(destTokenRule, destTokenItem)
	}

	logs, sub, err := _V5ParaSwap.contract.WatchLogs(opts, "SwappedV3", beneficiaryRule, srcTokenRule, destTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V5ParaSwapSwappedV3)
				if err := _V5ParaSwap.contract.UnpackLog(event, "SwappedV3", log); err != nil {
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

// ParseSwappedV3 is a log parse operation binding the contract event 0xe00361d207b252a464323eb23d45d42583e391f2031acdd2e9fa36efddd43cb0.
//
// Solidity: event SwappedV3(bytes16 uuid, address partner, uint256 feePercent, address initiator, address indexed beneficiary, address indexed srcToken, address indexed destToken, uint256 srcAmount, uint256 receivedAmount, uint256 expectedAmount)
func (_V5ParaSwap *V5ParaSwapFilterer) ParseSwappedV3(log types.Log) (*V5ParaSwapSwappedV3, error) {
	event := new(V5ParaSwapSwappedV3)
	if err := _V5ParaSwap.contract.UnpackLog(event, "SwappedV3", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
