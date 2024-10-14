// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zerion

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

// AbsoluteTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type AbsoluteTokenAmount struct {
	Token          common.Address
	AbsoluteAmount *big.Int
}

// AccountSignature is an auto generated low-level Go binding around an user-defined struct.
type AccountSignature struct {
	Salt      *big.Int
	Signature []byte
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Share       *big.Int
	Beneficiary common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	TokenAmount TokenAmount
	Permit      Permit
}

// Permit is an auto generated low-level Go binding around an user-defined struct.
type Permit struct {
	PermitType     uint8
	PermitCallData []byte
}

// ProtocolFeeSignature is an auto generated low-level Go binding around an user-defined struct.
type ProtocolFeeSignature struct {
	Deadline  *big.Int
	Signature []byte
}

// SwapDescription is an auto generated low-level Go binding around an user-defined struct.
type SwapDescription struct {
	SwapType       uint8
	ProtocolFee    Fee
	MarketplaceFee Fee
	Account        common.Address
	Caller         common.Address
	CallerCallData []byte
}

// TokenAmount is an auto generated low-level Go binding around an user-defined struct.
type TokenAmount struct {
	Token      common.Address
	Amount     *big.Int
	AmountType uint8
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedAccount\",\"type\":\"address\"}],\"name\":\"BadAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadAccountSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredAmount\",\"type\":\"uint256\"}],\"name\":\"BadAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumAmountType\",\"name\":\"amountType\",\"type\":\"uint8\"},{\"internalType\":\"enumAmountType\",\"name\":\"requiredAmountType\",\"type\":\"uint8\"}],\"name\":\"BadAmountType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedFeeAmount\",\"type\":\"uint256\"}],\"name\":\"BadFeeAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolFeeBanaficiary\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseProtocolFeeBeneficiary\",\"type\":\"address\"}],\"name\":\"BadFeeBeneficiary\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolFeeShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseProtocolFeeShare\",\"type\":\"uint256\"}],\"name\":\"BadFeeShare\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadFeeSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ExceedingDelimiterAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeLimit\",\"type\":\"uint256\"}],\"name\":\"ExceedingLimitFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inputBalanceChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredInputBalanceChange\",\"type\":\"uint256\"}],\"name\":\"HighInputBalanceChange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredAllowance\",\"type\":\"uint256\"}],\"name\":\"InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredMsgValue\",\"type\":\"uint256\"}],\"name\":\"InsufficientMsgValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actualOutputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requiredActualOutputAmount\",\"type\":\"uint256\"}],\"name\":\"LowActualOutputAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoneAmountType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonePermitType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoneSwapType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"PassedDeadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UsedHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeBeneficiary\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroReceiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSigner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"absoluteInputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inputBalanceChange\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"absoluteOutputAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"returnedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceFeeAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"marketplaceFee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callerCallData\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structSwapDescription\",\"name\":\"swapDescription\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldPendingOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProtocolFeeDefaultShare\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProtocolFeeDefaultBeneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFeeDefaultShare\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProtocolFeeDefaultBeneficiary\",\"type\":\"address\"}],\"name\":\"ProtocolFeeDefaultSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProtocolFeeSigner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProtocolFeeSigner\",\"type\":\"address\"}],\"name\":\"ProtocolFeeSignerSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumAmountType\",\"name\":\"amountType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPermitType\",\"name\":\"permitType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"permitCallData\",\"type\":\"bytes\"}],\"internalType\":\"structPermit\",\"name\":\"permit\",\"type\":\"tuple\"}],\"internalType\":\"structInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"absoluteAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAbsoluteTokenAmount\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"marketplaceFee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callerCallData\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescription\",\"name\":\"swapDescription\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structAccountSignature\",\"name\":\"accountSignature\",\"type\":\"tuple\"}],\"name\":\"cancelAccountSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumAmountType\",\"name\":\"amountType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPermitType\",\"name\":\"permitType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"permitCallData\",\"type\":\"bytes\"}],\"internalType\":\"structPermit\",\"name\":\"permit\",\"type\":\"tuple\"}],\"internalType\":\"structInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"absoluteAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAbsoluteTokenAmount\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"marketplaceFee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callerCallData\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescription\",\"name\":\"swapDescription\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structAccountSignature\",\"name\":\"accountSignature\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structProtocolFeeSignature\",\"name\":\"protocolFeeSignature\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"inputBalanceChange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualOutputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marketplaceFeeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeDefault\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFeeDefault\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"protocolFeeSigner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumAmountType\",\"name\":\"amountType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPermitType\",\"name\":\"permitType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"permitCallData\",\"type\":\"bytes\"}],\"internalType\":\"structPermit\",\"name\":\"permit\",\"type\":\"tuple\"}],\"internalType\":\"structInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"absoluteAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAbsoluteTokenAmount\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"marketplaceFee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callerCallData\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescription\",\"name\":\"swapDescription\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"name\":\"hashAccountSignatureData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedData\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"enumAmountType\",\"name\":\"amountType\",\"type\":\"uint8\"}],\"internalType\":\"structTokenAmount\",\"name\":\"tokenAmount\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumPermitType\",\"name\":\"permitType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"permitCallData\",\"type\":\"bytes\"}],\"internalType\":\"structPermit\",\"name\":\"permit\",\"type\":\"tuple\"}],\"internalType\":\"structInput\",\"name\":\"input\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"absoluteAmount\",\"type\":\"uint256\"}],\"internalType\":\"structAbsoluteTokenAmount\",\"name\":\"output\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumSwapType\",\"name\":\"swapType\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFee\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"marketplaceFee\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callerCallData\",\"type\":\"bytes\"}],\"internalType\":\"structSwapDescription\",\"name\":\"swapDescription\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"hashProtocolFeeSignatureData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashedData\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashToCheck\",\"type\":\"bytes32\"}],\"name\":\"isHashUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hashUsed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"returnLostTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingOwner\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"share\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"internalType\":\"structFee\",\"name\":\"protocolFeeDefault\",\"type\":\"tuple\"}],\"name\":\"setProtocolFeeDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"setProtocolFeeSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address owner)
func (_Router *RouterCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address owner)
func (_Router *RouterSession) GetOwner() (common.Address, error) {
	return _Router.Contract.GetOwner(&_Router.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address owner)
func (_Router *RouterCallerSession) GetOwner() (common.Address, error) {
	return _Router.Contract.GetOwner(&_Router.CallOpts)
}

// GetPendingOwner is a free data retrieval call binding the contract method 0x93596c7b.
//
// Solidity: function getPendingOwner() view returns(address pendingOwner)
func (_Router *RouterCaller) GetPendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getPendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPendingOwner is a free data retrieval call binding the contract method 0x93596c7b.
//
// Solidity: function getPendingOwner() view returns(address pendingOwner)
func (_Router *RouterSession) GetPendingOwner() (common.Address, error) {
	return _Router.Contract.GetPendingOwner(&_Router.CallOpts)
}

// GetPendingOwner is a free data retrieval call binding the contract method 0x93596c7b.
//
// Solidity: function getPendingOwner() view returns(address pendingOwner)
func (_Router *RouterCallerSession) GetPendingOwner() (common.Address, error) {
	return _Router.Contract.GetPendingOwner(&_Router.CallOpts)
}

// GetProtocolFeeDefault is a free data retrieval call binding the contract method 0x92ff2806.
//
// Solidity: function getProtocolFeeDefault() view returns((uint256,address) protocolFeeDefault)
func (_Router *RouterCaller) GetProtocolFeeDefault(opts *bind.CallOpts) (Fee, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getProtocolFeeDefault")

	if err != nil {
		return *new(Fee), err
	}

	out0 := *abi.ConvertType(out[0], new(Fee)).(*Fee)

	return out0, err

}

// GetProtocolFeeDefault is a free data retrieval call binding the contract method 0x92ff2806.
//
// Solidity: function getProtocolFeeDefault() view returns((uint256,address) protocolFeeDefault)
func (_Router *RouterSession) GetProtocolFeeDefault() (Fee, error) {
	return _Router.Contract.GetProtocolFeeDefault(&_Router.CallOpts)
}

// GetProtocolFeeDefault is a free data retrieval call binding the contract method 0x92ff2806.
//
// Solidity: function getProtocolFeeDefault() view returns((uint256,address) protocolFeeDefault)
func (_Router *RouterCallerSession) GetProtocolFeeDefault() (Fee, error) {
	return _Router.Contract.GetProtocolFeeDefault(&_Router.CallOpts)
}

// GetProtocolFeeSigner is a free data retrieval call binding the contract method 0xec5afea7.
//
// Solidity: function getProtocolFeeSigner() view returns(address protocolFeeSigner)
func (_Router *RouterCaller) GetProtocolFeeSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getProtocolFeeSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeeSigner is a free data retrieval call binding the contract method 0xec5afea7.
//
// Solidity: function getProtocolFeeSigner() view returns(address protocolFeeSigner)
func (_Router *RouterSession) GetProtocolFeeSigner() (common.Address, error) {
	return _Router.Contract.GetProtocolFeeSigner(&_Router.CallOpts)
}

// GetProtocolFeeSigner is a free data retrieval call binding the contract method 0xec5afea7.
//
// Solidity: function getProtocolFeeSigner() view returns(address protocolFeeSigner)
func (_Router *RouterCallerSession) GetProtocolFeeSigner() (common.Address, error) {
	return _Router.Contract.GetProtocolFeeSigner(&_Router.CallOpts)
}

// HashAccountSignatureData is a free data retrieval call binding the contract method 0xe1b92a9c.
//
// Solidity: function hashAccountSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 salt) view returns(bytes32 hashedData)
func (_Router *RouterCaller) HashAccountSignatureData(opts *bind.CallOpts, input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, salt *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "hashAccountSignatureData", input, output, swapDescription, salt)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashAccountSignatureData is a free data retrieval call binding the contract method 0xe1b92a9c.
//
// Solidity: function hashAccountSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 salt) view returns(bytes32 hashedData)
func (_Router *RouterSession) HashAccountSignatureData(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, salt *big.Int) ([32]byte, error) {
	return _Router.Contract.HashAccountSignatureData(&_Router.CallOpts, input, output, swapDescription, salt)
}

// HashAccountSignatureData is a free data retrieval call binding the contract method 0xe1b92a9c.
//
// Solidity: function hashAccountSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 salt) view returns(bytes32 hashedData)
func (_Router *RouterCallerSession) HashAccountSignatureData(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, salt *big.Int) ([32]byte, error) {
	return _Router.Contract.HashAccountSignatureData(&_Router.CallOpts, input, output, swapDescription, salt)
}

// HashProtocolFeeSignatureData is a free data retrieval call binding the contract method 0xebfa1a61.
//
// Solidity: function hashProtocolFeeSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 deadline) view returns(bytes32 hashedData)
func (_Router *RouterCaller) HashProtocolFeeSignatureData(opts *bind.CallOpts, input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, deadline *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "hashProtocolFeeSignatureData", input, output, swapDescription, deadline)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashProtocolFeeSignatureData is a free data retrieval call binding the contract method 0xebfa1a61.
//
// Solidity: function hashProtocolFeeSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 deadline) view returns(bytes32 hashedData)
func (_Router *RouterSession) HashProtocolFeeSignatureData(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, deadline *big.Int) ([32]byte, error) {
	return _Router.Contract.HashProtocolFeeSignatureData(&_Router.CallOpts, input, output, swapDescription, deadline)
}

// HashProtocolFeeSignatureData is a free data retrieval call binding the contract method 0xebfa1a61.
//
// Solidity: function hashProtocolFeeSignatureData(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, uint256 deadline) view returns(bytes32 hashedData)
func (_Router *RouterCallerSession) HashProtocolFeeSignatureData(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, deadline *big.Int) ([32]byte, error) {
	return _Router.Contract.HashProtocolFeeSignatureData(&_Router.CallOpts, input, output, swapDescription, deadline)
}

// IsHashUsed is a free data retrieval call binding the contract method 0x20021d8b.
//
// Solidity: function isHashUsed(bytes32 hashToCheck) view returns(bool hashUsed)
func (_Router *RouterCaller) IsHashUsed(opts *bind.CallOpts, hashToCheck [32]byte) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isHashUsed", hashToCheck)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHashUsed is a free data retrieval call binding the contract method 0x20021d8b.
//
// Solidity: function isHashUsed(bytes32 hashToCheck) view returns(bool hashUsed)
func (_Router *RouterSession) IsHashUsed(hashToCheck [32]byte) (bool, error) {
	return _Router.Contract.IsHashUsed(&_Router.CallOpts, hashToCheck)
}

// IsHashUsed is a free data retrieval call binding the contract method 0x20021d8b.
//
// Solidity: function isHashUsed(bytes32 hashToCheck) view returns(bool hashUsed)
func (_Router *RouterCallerSession) IsHashUsed(hashToCheck [32]byte) (bool, error) {
	return _Router.Contract.IsHashUsed(&_Router.CallOpts, hashToCheck)
}

// CancelAccountSignature is a paid mutator transaction binding the contract method 0xbba185de.
//
// Solidity: function cancelAccountSignature(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature) returns()
func (_Router *RouterTransactor) CancelAccountSignature(opts *bind.TransactOpts, input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "cancelAccountSignature", input, output, swapDescription, accountSignature)
}

// CancelAccountSignature is a paid mutator transaction binding the contract method 0xbba185de.
//
// Solidity: function cancelAccountSignature(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature) returns()
func (_Router *RouterSession) CancelAccountSignature(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature) (*types.Transaction, error) {
	return _Router.Contract.CancelAccountSignature(&_Router.TransactOpts, input, output, swapDescription, accountSignature)
}

// CancelAccountSignature is a paid mutator transaction binding the contract method 0xbba185de.
//
// Solidity: function cancelAccountSignature(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature) returns()
func (_Router *RouterTransactorSession) CancelAccountSignature(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature) (*types.Transaction, error) {
	return _Router.Contract.CancelAccountSignature(&_Router.TransactOpts, input, output, swapDescription, accountSignature)
}

// Execute is a paid mutator transaction binding the contract method 0x83d13e01.
//
// Solidity: function execute(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature, (uint256,bytes) protocolFeeSignature) payable returns(uint256 inputBalanceChange, uint256 actualOutputAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount)
func (_Router *RouterTransactor) Execute(opts *bind.TransactOpts, input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature, protocolFeeSignature ProtocolFeeSignature) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "execute", input, output, swapDescription, accountSignature, protocolFeeSignature)
}

// Execute is a paid mutator transaction binding the contract method 0x83d13e01.
//
// Solidity: function execute(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature, (uint256,bytes) protocolFeeSignature) payable returns(uint256 inputBalanceChange, uint256 actualOutputAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount)
func (_Router *RouterSession) Execute(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature, protocolFeeSignature ProtocolFeeSignature) (*types.Transaction, error) {
	return _Router.Contract.Execute(&_Router.TransactOpts, input, output, swapDescription, accountSignature, protocolFeeSignature)
}

// Execute is a paid mutator transaction binding the contract method 0x83d13e01.
//
// Solidity: function execute(((address,uint256,uint8),(uint8,bytes)) input, (address,uint256) output, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, (uint256,bytes) accountSignature, (uint256,bytes) protocolFeeSignature) payable returns(uint256 inputBalanceChange, uint256 actualOutputAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount)
func (_Router *RouterTransactorSession) Execute(input Input, output AbsoluteTokenAmount, swapDescription SwapDescription, accountSignature AccountSignature, protocolFeeSignature ProtocolFeeSignature) (*types.Transaction, error) {
	return _Router.Contract.Execute(&_Router.TransactOpts, input, output, swapDescription, accountSignature, protocolFeeSignature)
}

// ReturnLostTokens is a paid mutator transaction binding the contract method 0xe89db671.
//
// Solidity: function returnLostTokens(address token, address beneficiary) returns()
func (_Router *RouterTransactor) ReturnLostTokens(opts *bind.TransactOpts, token common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "returnLostTokens", token, beneficiary)
}

// ReturnLostTokens is a paid mutator transaction binding the contract method 0xe89db671.
//
// Solidity: function returnLostTokens(address token, address beneficiary) returns()
func (_Router *RouterSession) ReturnLostTokens(token common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _Router.Contract.ReturnLostTokens(&_Router.TransactOpts, token, beneficiary)
}

// ReturnLostTokens is a paid mutator transaction binding the contract method 0xe89db671.
//
// Solidity: function returnLostTokens(address token, address beneficiary) returns()
func (_Router *RouterTransactorSession) ReturnLostTokens(token common.Address, beneficiary common.Address) (*types.Transaction, error) {
	return _Router.Contract.ReturnLostTokens(&_Router.TransactOpts, token, beneficiary)
}

// SetOwner is a paid mutator transaction binding the contract method 0x40caae06.
//
// Solidity: function setOwner() returns()
func (_Router *RouterTransactor) SetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setOwner")
}

// SetOwner is a paid mutator transaction binding the contract method 0x40caae06.
//
// Solidity: function setOwner() returns()
func (_Router *RouterSession) SetOwner() (*types.Transaction, error) {
	return _Router.Contract.SetOwner(&_Router.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x40caae06.
//
// Solidity: function setOwner() returns()
func (_Router *RouterTransactorSession) SetOwner() (*types.Transaction, error) {
	return _Router.Contract.SetOwner(&_Router.TransactOpts)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Router *RouterTransactor) SetPendingOwner(opts *bind.TransactOpts, newPendingOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setPendingOwner", newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Router *RouterSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetPendingOwner(&_Router.TransactOpts, newPendingOwner)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address newPendingOwner) returns()
func (_Router *RouterTransactorSession) SetPendingOwner(newPendingOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetPendingOwner(&_Router.TransactOpts, newPendingOwner)
}

// SetProtocolFeeDefault is a paid mutator transaction binding the contract method 0xb0998bf7.
//
// Solidity: function setProtocolFeeDefault((uint256,address) protocolFeeDefault) returns()
func (_Router *RouterTransactor) SetProtocolFeeDefault(opts *bind.TransactOpts, protocolFeeDefault Fee) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setProtocolFeeDefault", protocolFeeDefault)
}

// SetProtocolFeeDefault is a paid mutator transaction binding the contract method 0xb0998bf7.
//
// Solidity: function setProtocolFeeDefault((uint256,address) protocolFeeDefault) returns()
func (_Router *RouterSession) SetProtocolFeeDefault(protocolFeeDefault Fee) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeDefault(&_Router.TransactOpts, protocolFeeDefault)
}

// SetProtocolFeeDefault is a paid mutator transaction binding the contract method 0xb0998bf7.
//
// Solidity: function setProtocolFeeDefault((uint256,address) protocolFeeDefault) returns()
func (_Router *RouterTransactorSession) SetProtocolFeeDefault(protocolFeeDefault Fee) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeDefault(&_Router.TransactOpts, protocolFeeDefault)
}

// SetProtocolFeeSigner is a paid mutator transaction binding the contract method 0xb297ccfd.
//
// Solidity: function setProtocolFeeSigner(address signer) returns()
func (_Router *RouterTransactor) SetProtocolFeeSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setProtocolFeeSigner", signer)
}

// SetProtocolFeeSigner is a paid mutator transaction binding the contract method 0xb297ccfd.
//
// Solidity: function setProtocolFeeSigner(address signer) returns()
func (_Router *RouterSession) SetProtocolFeeSigner(signer common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeSigner(&_Router.TransactOpts, signer)
}

// SetProtocolFeeSigner is a paid mutator transaction binding the contract method 0xb297ccfd.
//
// Solidity: function setProtocolFeeSigner(address signer) returns()
func (_Router *RouterTransactorSession) SetProtocolFeeSigner(signer common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetProtocolFeeSigner(&_Router.TransactOpts, signer)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// RouterExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the Router contract.
type RouterExecutedIterator struct {
	Event *RouterExecuted // Event containing the contract specifics and raw log

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
func (it *RouterExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterExecuted)
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
		it.Event = new(RouterExecuted)
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
func (it *RouterExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterExecuted represents a Executed event raised by the Router contract.
type RouterExecuted struct {
	InputToken           common.Address
	AbsoluteInputAmount  *big.Int
	InputBalanceChange   *big.Int
	OutputToken          common.Address
	AbsoluteOutputAmount *big.Int
	ReturnedAmount       *big.Int
	ProtocolFeeAmount    *big.Int
	MarketplaceFeeAmount *big.Int
	SwapDescription      SwapDescription
	Sender               common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x5b99554095b82fa9d0725a9ff1f1db8bc3e4d461e7d61911c752c349a47b987f.
//
// Solidity: event Executed(address indexed inputToken, uint256 absoluteInputAmount, uint256 inputBalanceChange, address indexed outputToken, uint256 absoluteOutputAmount, uint256 returnedAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, address sender)
func (_Router *RouterFilterer) FilterExecuted(opts *bind.FilterOpts, inputToken []common.Address, outputToken []common.Address) (*RouterExecutedIterator, error) {

	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}

	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "Executed", inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return &RouterExecutedIterator{contract: _Router.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x5b99554095b82fa9d0725a9ff1f1db8bc3e4d461e7d61911c752c349a47b987f.
//
// Solidity: event Executed(address indexed inputToken, uint256 absoluteInputAmount, uint256 inputBalanceChange, address indexed outputToken, uint256 absoluteOutputAmount, uint256 returnedAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, address sender)
func (_Router *RouterFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *RouterExecuted, inputToken []common.Address, outputToken []common.Address) (event.Subscription, error) {

	var inputTokenRule []interface{}
	for _, inputTokenItem := range inputToken {
		inputTokenRule = append(inputTokenRule, inputTokenItem)
	}

	var outputTokenRule []interface{}
	for _, outputTokenItem := range outputToken {
		outputTokenRule = append(outputTokenRule, outputTokenItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "Executed", inputTokenRule, outputTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterExecuted)
				if err := _Router.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x5b99554095b82fa9d0725a9ff1f1db8bc3e4d461e7d61911c752c349a47b987f.
//
// Solidity: event Executed(address indexed inputToken, uint256 absoluteInputAmount, uint256 inputBalanceChange, address indexed outputToken, uint256 absoluteOutputAmount, uint256 returnedAmount, uint256 protocolFeeAmount, uint256 marketplaceFeeAmount, (uint8,(uint256,address),(uint256,address),address,address,bytes) swapDescription, address sender)
func (_Router *RouterFilterer) ParseExecuted(log types.Log) (*RouterExecuted, error) {
	event := new(RouterExecuted)
	if err := _Router.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Router contract.
type RouterOwnerSetIterator struct {
	Event *RouterOwnerSet // Event containing the contract specifics and raw log

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
func (it *RouterOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnerSet)
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
		it.Event = new(RouterOwnerSet)
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
func (it *RouterOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnerSet represents a OwnerSet event raised by the Router contract.
type RouterOwnerSet struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnerSet(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*RouterOwnerSetIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnerSetIterator{contract: _Router.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *RouterOwnerSet, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnerSet", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnerSet)
				if err := _Router.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x342827c97908e5e2f71151c08502a66d44b6f758e3ac2f1de95f02eb95f0a735.
//
// Solidity: event OwnerSet(address indexed oldOwner, address indexed newOwner)
func (_Router *RouterFilterer) ParseOwnerSet(log types.Log) (*RouterOwnerSet, error) {
	event := new(RouterOwnerSet)
	if err := _Router.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterPendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the Router contract.
type RouterPendingOwnerSetIterator struct {
	Event *RouterPendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *RouterPendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterPendingOwnerSet)
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
		it.Event = new(RouterPendingOwnerSet)
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
func (it *RouterPendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterPendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterPendingOwnerSet represents a PendingOwnerSet event raised by the Router contract.
type RouterPendingOwnerSet struct {
	OldPendingOwner common.Address
	NewPendingOwner common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Router *RouterFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, oldPendingOwner []common.Address, newPendingOwner []common.Address) (*RouterPendingOwnerSetIterator, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "PendingOwnerSet", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterPendingOwnerSetIterator{contract: _Router.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Router *RouterFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *RouterPendingOwnerSet, oldPendingOwner []common.Address, newPendingOwner []common.Address) (event.Subscription, error) {

	var oldPendingOwnerRule []interface{}
	for _, oldPendingOwnerItem := range oldPendingOwner {
		oldPendingOwnerRule = append(oldPendingOwnerRule, oldPendingOwnerItem)
	}
	var newPendingOwnerRule []interface{}
	for _, newPendingOwnerItem := range newPendingOwner {
		newPendingOwnerRule = append(newPendingOwnerRule, newPendingOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "PendingOwnerSet", oldPendingOwnerRule, newPendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterPendingOwnerSet)
				if err := _Router.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0xa86864fa6b65f969d5ac8391ddaac6a0eba3f41386cbf6e78c3e4d6c59eb115f.
//
// Solidity: event PendingOwnerSet(address indexed oldPendingOwner, address indexed newPendingOwner)
func (_Router *RouterFilterer) ParsePendingOwnerSet(log types.Log) (*RouterPendingOwnerSet, error) {
	event := new(RouterPendingOwnerSet)
	if err := _Router.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterProtocolFeeDefaultSetIterator is returned from FilterProtocolFeeDefaultSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeDefaultSet events raised by the Router contract.
type RouterProtocolFeeDefaultSetIterator struct {
	Event *RouterProtocolFeeDefaultSet // Event containing the contract specifics and raw log

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
func (it *RouterProtocolFeeDefaultSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterProtocolFeeDefaultSet)
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
		it.Event = new(RouterProtocolFeeDefaultSet)
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
func (it *RouterProtocolFeeDefaultSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterProtocolFeeDefaultSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterProtocolFeeDefaultSet represents a ProtocolFeeDefaultSet event raised by the Router contract.
type RouterProtocolFeeDefaultSet struct {
	OldProtocolFeeDefaultShare       *big.Int
	OldProtocolFeeDefaultBeneficiary common.Address
	NewProtocolFeeDefaultShare       *big.Int
	NewProtocolFeeDefaultBeneficiary common.Address
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeDefaultSet is a free log retrieval operation binding the contract event 0xbed3189b824b5183bdd720fd47ce4f7747e55802778e0c0010d61647d4b60360.
//
// Solidity: event ProtocolFeeDefaultSet(uint256 oldProtocolFeeDefaultShare, address indexed oldProtocolFeeDefaultBeneficiary, uint256 newProtocolFeeDefaultShare, address indexed newProtocolFeeDefaultBeneficiary)
func (_Router *RouterFilterer) FilterProtocolFeeDefaultSet(opts *bind.FilterOpts, oldProtocolFeeDefaultBeneficiary []common.Address, newProtocolFeeDefaultBeneficiary []common.Address) (*RouterProtocolFeeDefaultSetIterator, error) {

	var oldProtocolFeeDefaultBeneficiaryRule []interface{}
	for _, oldProtocolFeeDefaultBeneficiaryItem := range oldProtocolFeeDefaultBeneficiary {
		oldProtocolFeeDefaultBeneficiaryRule = append(oldProtocolFeeDefaultBeneficiaryRule, oldProtocolFeeDefaultBeneficiaryItem)
	}

	var newProtocolFeeDefaultBeneficiaryRule []interface{}
	for _, newProtocolFeeDefaultBeneficiaryItem := range newProtocolFeeDefaultBeneficiary {
		newProtocolFeeDefaultBeneficiaryRule = append(newProtocolFeeDefaultBeneficiaryRule, newProtocolFeeDefaultBeneficiaryItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ProtocolFeeDefaultSet", oldProtocolFeeDefaultBeneficiaryRule, newProtocolFeeDefaultBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &RouterProtocolFeeDefaultSetIterator{contract: _Router.contract, event: "ProtocolFeeDefaultSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeDefaultSet is a free log subscription operation binding the contract event 0xbed3189b824b5183bdd720fd47ce4f7747e55802778e0c0010d61647d4b60360.
//
// Solidity: event ProtocolFeeDefaultSet(uint256 oldProtocolFeeDefaultShare, address indexed oldProtocolFeeDefaultBeneficiary, uint256 newProtocolFeeDefaultShare, address indexed newProtocolFeeDefaultBeneficiary)
func (_Router *RouterFilterer) WatchProtocolFeeDefaultSet(opts *bind.WatchOpts, sink chan<- *RouterProtocolFeeDefaultSet, oldProtocolFeeDefaultBeneficiary []common.Address, newProtocolFeeDefaultBeneficiary []common.Address) (event.Subscription, error) {

	var oldProtocolFeeDefaultBeneficiaryRule []interface{}
	for _, oldProtocolFeeDefaultBeneficiaryItem := range oldProtocolFeeDefaultBeneficiary {
		oldProtocolFeeDefaultBeneficiaryRule = append(oldProtocolFeeDefaultBeneficiaryRule, oldProtocolFeeDefaultBeneficiaryItem)
	}

	var newProtocolFeeDefaultBeneficiaryRule []interface{}
	for _, newProtocolFeeDefaultBeneficiaryItem := range newProtocolFeeDefaultBeneficiary {
		newProtocolFeeDefaultBeneficiaryRule = append(newProtocolFeeDefaultBeneficiaryRule, newProtocolFeeDefaultBeneficiaryItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ProtocolFeeDefaultSet", oldProtocolFeeDefaultBeneficiaryRule, newProtocolFeeDefaultBeneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterProtocolFeeDefaultSet)
				if err := _Router.contract.UnpackLog(event, "ProtocolFeeDefaultSet", log); err != nil {
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

// ParseProtocolFeeDefaultSet is a log parse operation binding the contract event 0xbed3189b824b5183bdd720fd47ce4f7747e55802778e0c0010d61647d4b60360.
//
// Solidity: event ProtocolFeeDefaultSet(uint256 oldProtocolFeeDefaultShare, address indexed oldProtocolFeeDefaultBeneficiary, uint256 newProtocolFeeDefaultShare, address indexed newProtocolFeeDefaultBeneficiary)
func (_Router *RouterFilterer) ParseProtocolFeeDefaultSet(log types.Log) (*RouterProtocolFeeDefaultSet, error) {
	event := new(RouterProtocolFeeDefaultSet)
	if err := _Router.contract.UnpackLog(event, "ProtocolFeeDefaultSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RouterProtocolFeeSignerSetIterator is returned from FilterProtocolFeeSignerSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeSignerSet events raised by the Router contract.
type RouterProtocolFeeSignerSetIterator struct {
	Event *RouterProtocolFeeSignerSet // Event containing the contract specifics and raw log

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
func (it *RouterProtocolFeeSignerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterProtocolFeeSignerSet)
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
		it.Event = new(RouterProtocolFeeSignerSet)
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
func (it *RouterProtocolFeeSignerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterProtocolFeeSignerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterProtocolFeeSignerSet represents a ProtocolFeeSignerSet event raised by the Router contract.
type RouterProtocolFeeSignerSet struct {
	OldProtocolFeeSigner common.Address
	NewProtocolFeeSigner common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeSignerSet is a free log retrieval operation binding the contract event 0x401823c5aef6f718747b02330ff8612a8f1734ad582a37101b7f74aedf88a3cf.
//
// Solidity: event ProtocolFeeSignerSet(address indexed oldProtocolFeeSigner, address indexed newProtocolFeeSigner)
func (_Router *RouterFilterer) FilterProtocolFeeSignerSet(opts *bind.FilterOpts, oldProtocolFeeSigner []common.Address, newProtocolFeeSigner []common.Address) (*RouterProtocolFeeSignerSetIterator, error) {

	var oldProtocolFeeSignerRule []interface{}
	for _, oldProtocolFeeSignerItem := range oldProtocolFeeSigner {
		oldProtocolFeeSignerRule = append(oldProtocolFeeSignerRule, oldProtocolFeeSignerItem)
	}
	var newProtocolFeeSignerRule []interface{}
	for _, newProtocolFeeSignerItem := range newProtocolFeeSigner {
		newProtocolFeeSignerRule = append(newProtocolFeeSignerRule, newProtocolFeeSignerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "ProtocolFeeSignerSet", oldProtocolFeeSignerRule, newProtocolFeeSignerRule)
	if err != nil {
		return nil, err
	}
	return &RouterProtocolFeeSignerSetIterator{contract: _Router.contract, event: "ProtocolFeeSignerSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSignerSet is a free log subscription operation binding the contract event 0x401823c5aef6f718747b02330ff8612a8f1734ad582a37101b7f74aedf88a3cf.
//
// Solidity: event ProtocolFeeSignerSet(address indexed oldProtocolFeeSigner, address indexed newProtocolFeeSigner)
func (_Router *RouterFilterer) WatchProtocolFeeSignerSet(opts *bind.WatchOpts, sink chan<- *RouterProtocolFeeSignerSet, oldProtocolFeeSigner []common.Address, newProtocolFeeSigner []common.Address) (event.Subscription, error) {

	var oldProtocolFeeSignerRule []interface{}
	for _, oldProtocolFeeSignerItem := range oldProtocolFeeSigner {
		oldProtocolFeeSignerRule = append(oldProtocolFeeSignerRule, oldProtocolFeeSignerItem)
	}
	var newProtocolFeeSignerRule []interface{}
	for _, newProtocolFeeSignerItem := range newProtocolFeeSigner {
		newProtocolFeeSignerRule = append(newProtocolFeeSignerRule, newProtocolFeeSignerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "ProtocolFeeSignerSet", oldProtocolFeeSignerRule, newProtocolFeeSignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterProtocolFeeSignerSet)
				if err := _Router.contract.UnpackLog(event, "ProtocolFeeSignerSet", log); err != nil {
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

// ParseProtocolFeeSignerSet is a log parse operation binding the contract event 0x401823c5aef6f718747b02330ff8612a8f1734ad582a37101b7f74aedf88a3cf.
//
// Solidity: event ProtocolFeeSignerSet(address indexed oldProtocolFeeSigner, address indexed newProtocolFeeSigner)
func (_Router *RouterFilterer) ParseProtocolFeeSignerSet(log types.Log) (*RouterProtocolFeeSignerSet, error) {
	event := new(RouterProtocolFeeSignerSet)
	if err := _Router.contract.UnpackLog(event, "ProtocolFeeSignerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
