// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arbitrum

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

// L1CustomGatewayMetaData contains all meta data concerning the L1CustomGateway contract.
var L1CustomGatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l1Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"l2Address\",\"type\":\"address\"}],\"name\":\"TokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_seqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TxToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"madeExternalCall\",\"type\":\"bool\"}],\"name\":\"WithdrawRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1ERC20\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"}],\"name\":\"encodeWithdrawal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_l1Addresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_l2Addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"forceRegisterTokenToL2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initialData\",\"type\":\"bytes\"}],\"name\":\"getExternalCall\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"outboundCalldata\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"l1ToL2Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransferCustomRefund\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"redirectedExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_newTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_newData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creditBackAddress\",\"type\":\"address\"}],\"name\":\"registerTokenToL2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSubmissionCost\",\"type\":\"uint256\"}],\"name\":\"registerTokenToL2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_newData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferExitAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1CustomGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1CustomGatewayMetaData.ABI instead.
var L1CustomGatewayABI = L1CustomGatewayMetaData.ABI

// L1CustomGateway is an auto generated Go binding around an Ethereum contract.
type L1CustomGateway struct {
	L1CustomGatewayCaller     // Read-only binding to the contract
	L1CustomGatewayTransactor // Write-only binding to the contract
	L1CustomGatewayFilterer   // Log filterer for contract events
}

// L1CustomGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1CustomGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1CustomGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1CustomGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1CustomGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1CustomGatewaySession struct {
	Contract     *L1CustomGateway  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1CustomGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1CustomGatewayCallerSession struct {
	Contract *L1CustomGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// L1CustomGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1CustomGatewayTransactorSession struct {
	Contract     *L1CustomGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// L1CustomGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1CustomGatewayRaw struct {
	Contract *L1CustomGateway // Generic contract binding to access the raw methods on
}

// L1CustomGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1CustomGatewayCallerRaw struct {
	Contract *L1CustomGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1CustomGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1CustomGatewayTransactorRaw struct {
	Contract *L1CustomGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1CustomGateway creates a new instance of L1CustomGateway, bound to a specific deployed contract.
func NewL1CustomGateway(address common.Address, backend bind.ContractBackend) (*L1CustomGateway, error) {
	contract, err := bindL1CustomGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1CustomGateway{L1CustomGatewayCaller: L1CustomGatewayCaller{contract: contract}, L1CustomGatewayTransactor: L1CustomGatewayTransactor{contract: contract}, L1CustomGatewayFilterer: L1CustomGatewayFilterer{contract: contract}}, nil
}

// NewL1CustomGatewayCaller creates a new read-only instance of L1CustomGateway, bound to a specific deployed contract.
func NewL1CustomGatewayCaller(address common.Address, caller bind.ContractCaller) (*L1CustomGatewayCaller, error) {
	contract, err := bindL1CustomGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayCaller{contract: contract}, nil
}

// NewL1CustomGatewayTransactor creates a new write-only instance of L1CustomGateway, bound to a specific deployed contract.
func NewL1CustomGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1CustomGatewayTransactor, error) {
	contract, err := bindL1CustomGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayTransactor{contract: contract}, nil
}

// NewL1CustomGatewayFilterer creates a new log filterer instance of L1CustomGateway, bound to a specific deployed contract.
func NewL1CustomGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1CustomGatewayFilterer, error) {
	contract, err := bindL1CustomGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayFilterer{contract: contract}, nil
}

// bindL1CustomGateway binds a generic wrapper to an already deployed contract.
func bindL1CustomGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1CustomGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1CustomGateway *L1CustomGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1CustomGateway.Contract.L1CustomGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1CustomGateway *L1CustomGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.L1CustomGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1CustomGateway *L1CustomGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.L1CustomGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1CustomGateway *L1CustomGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1CustomGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1CustomGateway *L1CustomGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1CustomGateway *L1CustomGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1ERC20 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "calculateL2TokenAddress", l1ERC20)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1CustomGateway.Contract.CalculateL2TokenAddress(&_L1CustomGateway.CallOpts, l1ERC20)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1CustomGateway.Contract.CalculateL2TokenAddress(&_L1CustomGateway.CallOpts, l1ERC20)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) CounterpartGateway() (common.Address, error) {
	return _L1CustomGateway.Contract.CounterpartGateway(&_L1CustomGateway.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) CounterpartGateway() (common.Address, error) {
	return _L1CustomGateway.Contract.CounterpartGateway(&_L1CustomGateway.CallOpts)
}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1CustomGateway *L1CustomGatewayCaller) EncodeWithdrawal(opts *bind.CallOpts, _exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "encodeWithdrawal", _exitNum, _initialDestination)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1CustomGateway *L1CustomGatewaySession) EncodeWithdrawal(_exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	return _L1CustomGateway.Contract.EncodeWithdrawal(&_L1CustomGateway.CallOpts, _exitNum, _initialDestination)
}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1CustomGateway *L1CustomGatewayCallerSession) EncodeWithdrawal(_exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	return _L1CustomGateway.Contract.EncodeWithdrawal(&_L1CustomGateway.CallOpts, _exitNum, _initialDestination)
}

// GetExternalCall is a free data retrieval call binding the contract method 0xf68a9082.
//
// Solidity: function getExternalCall(uint256 _exitNum, address _initialDestination, bytes _initialData) view returns(address target, bytes data)
func (_L1CustomGateway *L1CustomGatewayCaller) GetExternalCall(opts *bind.CallOpts, _exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "getExternalCall", _exitNum, _initialDestination, _initialData)

	outstruct := new(struct {
		Target common.Address
		Data   []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Target = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetExternalCall is a free data retrieval call binding the contract method 0xf68a9082.
//
// Solidity: function getExternalCall(uint256 _exitNum, address _initialDestination, bytes _initialData) view returns(address target, bytes data)
func (_L1CustomGateway *L1CustomGatewaySession) GetExternalCall(_exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	return _L1CustomGateway.Contract.GetExternalCall(&_L1CustomGateway.CallOpts, _exitNum, _initialDestination, _initialData)
}

// GetExternalCall is a free data retrieval call binding the contract method 0xf68a9082.
//
// Solidity: function getExternalCall(uint256 _exitNum, address _initialDestination, bytes _initialData) view returns(address target, bytes data)
func (_L1CustomGateway *L1CustomGatewayCallerSession) GetExternalCall(_exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	return _L1CustomGateway.Contract.GetExternalCall(&_L1CustomGateway.CallOpts, _exitNum, _initialDestination, _initialData)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _l1Token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1CustomGateway *L1CustomGatewayCaller) GetOutboundCalldata(opts *bind.CallOpts, _l1Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "getOutboundCalldata", _l1Token, _from, _to, _amount, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _l1Token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1CustomGateway *L1CustomGatewaySession) GetOutboundCalldata(_l1Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1CustomGateway.Contract.GetOutboundCalldata(&_L1CustomGateway.CallOpts, _l1Token, _from, _to, _amount, _data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _l1Token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1CustomGateway *L1CustomGatewayCallerSession) GetOutboundCalldata(_l1Token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1CustomGateway.Contract.GetOutboundCalldata(&_L1CustomGateway.CallOpts, _l1Token, _from, _to, _amount, _data)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) Inbox() (common.Address, error) {
	return _L1CustomGateway.Contract.Inbox(&_L1CustomGateway.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) Inbox() (common.Address, error) {
	return _L1CustomGateway.Contract.Inbox(&_L1CustomGateway.CallOpts)
}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) L1ToL2Token(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "l1ToL2Token", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) L1ToL2Token(arg0 common.Address) (common.Address, error) {
	return _L1CustomGateway.Contract.L1ToL2Token(&_L1CustomGateway.CallOpts, arg0)
}

// L1ToL2Token is a free data retrieval call binding the contract method 0x8a2dc014.
//
// Solidity: function l1ToL2Token(address ) view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) L1ToL2Token(arg0 common.Address) (common.Address, error) {
	return _L1CustomGateway.Contract.L1ToL2Token(&_L1CustomGateway.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) Owner() (common.Address, error) {
	return _L1CustomGateway.Contract.Owner(&_L1CustomGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) Owner() (common.Address, error) {
	return _L1CustomGateway.Contract.Owner(&_L1CustomGateway.CallOpts)
}

// RedirectedExits is a free data retrieval call binding the contract method 0xbcf2e6eb.
//
// Solidity: function redirectedExits(bytes32 ) view returns(bool isExit, address _newTo, bytes _newData)
func (_L1CustomGateway *L1CustomGatewayCaller) RedirectedExits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "redirectedExits", arg0)

	outstruct := new(struct {
		IsExit  bool
		NewTo   common.Address
		NewData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsExit = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.NewTo = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NewData = *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return *outstruct, err

}

// RedirectedExits is a free data retrieval call binding the contract method 0xbcf2e6eb.
//
// Solidity: function redirectedExits(bytes32 ) view returns(bool isExit, address _newTo, bytes _newData)
func (_L1CustomGateway *L1CustomGatewaySession) RedirectedExits(arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	return _L1CustomGateway.Contract.RedirectedExits(&_L1CustomGateway.CallOpts, arg0)
}

// RedirectedExits is a free data retrieval call binding the contract method 0xbcf2e6eb.
//
// Solidity: function redirectedExits(bytes32 ) view returns(bool isExit, address _newTo, bytes _newData)
func (_L1CustomGateway *L1CustomGatewayCallerSession) RedirectedExits(arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	return _L1CustomGateway.Contract.RedirectedExits(&_L1CustomGateway.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) Router() (common.Address, error) {
	return _L1CustomGateway.Contract.Router(&_L1CustomGateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) Router() (common.Address, error) {
	return _L1CustomGateway.Contract.Router(&_L1CustomGateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1CustomGateway *L1CustomGatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1CustomGateway *L1CustomGatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1CustomGateway.Contract.SupportsInterface(&_L1CustomGateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1CustomGateway *L1CustomGatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1CustomGateway.Contract.SupportsInterface(&_L1CustomGateway.CallOpts, interfaceId)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1CustomGateway.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1CustomGateway *L1CustomGatewaySession) Whitelist() (common.Address, error) {
	return _L1CustomGateway.Contract.Whitelist(&_L1CustomGateway.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1CustomGateway *L1CustomGatewayCallerSession) Whitelist() (common.Address, error) {
	return _L1CustomGateway.Contract.Whitelist(&_L1CustomGateway.CallOpts)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1CustomGateway *L1CustomGatewayTransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "finalizeInboundTransfer", _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1CustomGateway *L1CustomGatewaySession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.FinalizeInboundTransfer(&_L1CustomGateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1CustomGateway *L1CustomGatewayTransactorSession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.FinalizeInboundTransfer(&_L1CustomGateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// ForceRegisterTokenToL2 is a paid mutator transaction binding the contract method 0x1d3a689f.
//
// Solidity: function forceRegisterTokenToL2(address[] _l1Addresses, address[] _l2Addresses, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactor) ForceRegisterTokenToL2(opts *bind.TransactOpts, _l1Addresses []common.Address, _l2Addresses []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "forceRegisterTokenToL2", _l1Addresses, _l2Addresses, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// ForceRegisterTokenToL2 is a paid mutator transaction binding the contract method 0x1d3a689f.
//
// Solidity: function forceRegisterTokenToL2(address[] _l1Addresses, address[] _l2Addresses, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewaySession) ForceRegisterTokenToL2(_l1Addresses []common.Address, _l2Addresses []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.ForceRegisterTokenToL2(&_L1CustomGateway.TransactOpts, _l1Addresses, _l2Addresses, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// ForceRegisterTokenToL2 is a paid mutator transaction binding the contract method 0x1d3a689f.
//
// Solidity: function forceRegisterTokenToL2(address[] _l1Addresses, address[] _l2Addresses, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactorSession) ForceRegisterTokenToL2(_l1Addresses []common.Address, _l2Addresses []common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.ForceRegisterTokenToL2(&_L1CustomGateway.TransactOpts, _l1Addresses, _l2Addresses, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _l1Counterpart, address _l1Router, address _inbox, address _owner) returns()
func (_L1CustomGateway *L1CustomGatewayTransactor) Initialize(opts *bind.TransactOpts, _l1Counterpart common.Address, _l1Router common.Address, _inbox common.Address, _owner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "initialize", _l1Counterpart, _l1Router, _inbox, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _l1Counterpart, address _l1Router, address _inbox, address _owner) returns()
func (_L1CustomGateway *L1CustomGatewaySession) Initialize(_l1Counterpart common.Address, _l1Router common.Address, _inbox common.Address, _owner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.Initialize(&_L1CustomGateway.TransactOpts, _l1Counterpart, _l1Router, _inbox, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _l1Counterpart, address _l1Router, address _inbox, address _owner) returns()
func (_L1CustomGateway *L1CustomGatewayTransactorSession) Initialize(_l1Counterpart common.Address, _l1Router common.Address, _inbox common.Address, _owner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.Initialize(&_L1CustomGateway.TransactOpts, _l1Counterpart, _l1Router, _inbox, _owner)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewayTransactor) OutboundTransfer(opts *bind.TransactOpts, _l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "outboundTransfer", _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewaySession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.OutboundTransfer(&_L1CustomGateway.TransactOpts, _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewayTransactorSession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.OutboundTransfer(&_L1CustomGateway.TransactOpts, _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewayTransactor) OutboundTransferCustomRefund(opts *bind.TransactOpts, _l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "outboundTransferCustomRefund", _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewaySession) OutboundTransferCustomRefund(_l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.OutboundTransferCustomRefund(&_L1CustomGateway.TransactOpts, _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1CustomGateway *L1CustomGatewayTransactorSession) OutboundTransferCustomRefund(_l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.OutboundTransferCustomRefund(&_L1CustomGateway.TransactOpts, _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1CustomGateway *L1CustomGatewayTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1CustomGateway *L1CustomGatewaySession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1CustomGateway.Contract.PostUpgradeInit(&_L1CustomGateway.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1CustomGateway *L1CustomGatewayTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1CustomGateway.Contract.PostUpgradeInit(&_L1CustomGateway.TransactOpts)
}

// RegisterTokenToL2 is a paid mutator transaction binding the contract method 0xca346d4a.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactor) RegisterTokenToL2(opts *bind.TransactOpts, _l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "registerTokenToL2", _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// RegisterTokenToL2 is a paid mutator transaction binding the contract method 0xca346d4a.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewaySession) RegisterTokenToL2(_l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.RegisterTokenToL2(&_L1CustomGateway.TransactOpts, _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// RegisterTokenToL2 is a paid mutator transaction binding the contract method 0xca346d4a.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost, address _creditBackAddress) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactorSession) RegisterTokenToL2(_l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int, _creditBackAddress common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.RegisterTokenToL2(&_L1CustomGateway.TransactOpts, _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost, _creditBackAddress)
}

// RegisterTokenToL20 is a paid mutator transaction binding the contract method 0xf26bdead.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactor) RegisterTokenToL20(opts *bind.TransactOpts, _l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "registerTokenToL20", _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// RegisterTokenToL20 is a paid mutator transaction binding the contract method 0xf26bdead.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewaySession) RegisterTokenToL20(_l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.RegisterTokenToL20(&_L1CustomGateway.TransactOpts, _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// RegisterTokenToL20 is a paid mutator transaction binding the contract method 0xf26bdead.
//
// Solidity: function registerTokenToL2(address _l2Address, uint256 _maxGas, uint256 _gasPriceBid, uint256 _maxSubmissionCost) payable returns(uint256)
func (_L1CustomGateway *L1CustomGatewayTransactorSession) RegisterTokenToL20(_l2Address common.Address, _maxGas *big.Int, _gasPriceBid *big.Int, _maxSubmissionCost *big.Int) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.RegisterTokenToL20(&_L1CustomGateway.TransactOpts, _l2Address, _maxGas, _gasPriceBid, _maxSubmissionCost)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1CustomGateway *L1CustomGatewayTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1CustomGateway *L1CustomGatewaySession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.SetOwner(&_L1CustomGateway.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_L1CustomGateway *L1CustomGatewayTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.SetOwner(&_L1CustomGateway.TransactOpts, newOwner)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1CustomGateway *L1CustomGatewayTransactor) TransferExitAndCall(opts *bind.TransactOpts, _exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.contract.Transact(opts, "transferExitAndCall", _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1CustomGateway *L1CustomGatewaySession) TransferExitAndCall(_exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.TransferExitAndCall(&_L1CustomGateway.TransactOpts, _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1CustomGateway *L1CustomGatewayTransactorSession) TransferExitAndCall(_exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1CustomGateway.Contract.TransferExitAndCall(&_L1CustomGateway.TransactOpts, _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// L1CustomGatewayDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the L1CustomGateway contract.
type L1CustomGatewayDepositInitiatedIterator struct {
	Event *L1CustomGatewayDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1CustomGatewayDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomGatewayDepositInitiated)
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
		it.Event = new(L1CustomGatewayDepositInitiated)
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
func (it *L1CustomGatewayDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomGatewayDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomGatewayDepositInitiated represents a DepositInitiated event raised by the L1CustomGateway contract.
type L1CustomGatewayDepositInitiated struct {
	L1Token        common.Address
	From           common.Address
	To             common.Address
	SequenceNumber *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDepositInitiated is a free log retrieval operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _sequenceNumber, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) FilterDepositInitiated(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _sequenceNumber []*big.Int) (*L1CustomGatewayDepositInitiatedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _sequenceNumberRule []interface{}
	for _, _sequenceNumberItem := range _sequenceNumber {
		_sequenceNumberRule = append(_sequenceNumberRule, _sequenceNumberItem)
	}

	logs, sub, err := _L1CustomGateway.contract.FilterLogs(opts, "DepositInitiated", _fromRule, _toRule, _sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayDepositInitiatedIterator{contract: _L1CustomGateway.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _sequenceNumber, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1CustomGatewayDepositInitiated, _from []common.Address, _to []common.Address, _sequenceNumber []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _sequenceNumberRule []interface{}
	for _, _sequenceNumberItem := range _sequenceNumber {
		_sequenceNumberRule = append(_sequenceNumberRule, _sequenceNumberItem)
	}

	logs, sub, err := _L1CustomGateway.contract.WatchLogs(opts, "DepositInitiated", _fromRule, _toRule, _sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomGatewayDepositInitiated)
				if err := _L1CustomGateway.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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

// ParseDepositInitiated is a log parse operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _sequenceNumber, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) ParseDepositInitiated(log types.Log) (*L1CustomGatewayDepositInitiated, error) {
	event := new(L1CustomGatewayDepositInitiated)
	if err := _L1CustomGateway.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomGatewayTokenSetIterator is returned from FilterTokenSet and is used to iterate over the raw logs and unpacked data for TokenSet events raised by the L1CustomGateway contract.
type L1CustomGatewayTokenSetIterator struct {
	Event *L1CustomGatewayTokenSet // Event containing the contract specifics and raw log

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
func (it *L1CustomGatewayTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomGatewayTokenSet)
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
		it.Event = new(L1CustomGatewayTokenSet)
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
func (it *L1CustomGatewayTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomGatewayTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomGatewayTokenSet represents a TokenSet event raised by the L1CustomGateway contract.
type L1CustomGatewayTokenSet struct {
	L1Address common.Address
	L2Address common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSet is a free log retrieval operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L1CustomGateway *L1CustomGatewayFilterer) FilterTokenSet(opts *bind.FilterOpts, l1Address []common.Address, l2Address []common.Address) (*L1CustomGatewayTokenSetIterator, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}

	logs, sub, err := _L1CustomGateway.contract.FilterLogs(opts, "TokenSet", l1AddressRule, l2AddressRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayTokenSetIterator{contract: _L1CustomGateway.contract, event: "TokenSet", logs: logs, sub: sub}, nil
}

// WatchTokenSet is a free log subscription operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L1CustomGateway *L1CustomGatewayFilterer) WatchTokenSet(opts *bind.WatchOpts, sink chan<- *L1CustomGatewayTokenSet, l1Address []common.Address, l2Address []common.Address) (event.Subscription, error) {

	var l1AddressRule []interface{}
	for _, l1AddressItem := range l1Address {
		l1AddressRule = append(l1AddressRule, l1AddressItem)
	}
	var l2AddressRule []interface{}
	for _, l2AddressItem := range l2Address {
		l2AddressRule = append(l2AddressRule, l2AddressItem)
	}

	logs, sub, err := _L1CustomGateway.contract.WatchLogs(opts, "TokenSet", l1AddressRule, l2AddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomGatewayTokenSet)
				if err := _L1CustomGateway.contract.UnpackLog(event, "TokenSet", log); err != nil {
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

// ParseTokenSet is a log parse operation binding the contract event 0x0dd664a155dd89526bb019e22b00291bb7ca9d07ba3ec4a1a76b410da9797ceb.
//
// Solidity: event TokenSet(address indexed l1Address, address indexed l2Address)
func (_L1CustomGateway *L1CustomGatewayFilterer) ParseTokenSet(log types.Log) (*L1CustomGatewayTokenSet, error) {
	event := new(L1CustomGatewayTokenSet)
	if err := _L1CustomGateway.contract.UnpackLog(event, "TokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomGatewayTxToL2Iterator is returned from FilterTxToL2 and is used to iterate over the raw logs and unpacked data for TxToL2 events raised by the L1CustomGateway contract.
type L1CustomGatewayTxToL2Iterator struct {
	Event *L1CustomGatewayTxToL2 // Event containing the contract specifics and raw log

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
func (it *L1CustomGatewayTxToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomGatewayTxToL2)
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
		it.Event = new(L1CustomGatewayTxToL2)
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
func (it *L1CustomGatewayTxToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomGatewayTxToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomGatewayTxToL2 represents a TxToL2 event raised by the L1CustomGateway contract.
type L1CustomGatewayTxToL2 struct {
	From   common.Address
	To     common.Address
	SeqNum *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxToL2 is a free log retrieval operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1CustomGateway *L1CustomGatewayFilterer) FilterTxToL2(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (*L1CustomGatewayTxToL2Iterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _seqNumRule []interface{}
	for _, _seqNumItem := range _seqNum {
		_seqNumRule = append(_seqNumRule, _seqNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.FilterLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayTxToL2Iterator{contract: _L1CustomGateway.contract, event: "TxToL2", logs: logs, sub: sub}, nil
}

// WatchTxToL2 is a free log subscription operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1CustomGateway *L1CustomGatewayFilterer) WatchTxToL2(opts *bind.WatchOpts, sink chan<- *L1CustomGatewayTxToL2, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _seqNumRule []interface{}
	for _, _seqNumItem := range _seqNum {
		_seqNumRule = append(_seqNumRule, _seqNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.WatchLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomGatewayTxToL2)
				if err := _L1CustomGateway.contract.UnpackLog(event, "TxToL2", log); err != nil {
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

// ParseTxToL2 is a log parse operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1CustomGateway *L1CustomGatewayFilterer) ParseTxToL2(log types.Log) (*L1CustomGatewayTxToL2, error) {
	event := new(L1CustomGatewayTxToL2)
	if err := _L1CustomGateway.contract.UnpackLog(event, "TxToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomGatewayWithdrawRedirectedIterator is returned from FilterWithdrawRedirected and is used to iterate over the raw logs and unpacked data for WithdrawRedirected events raised by the L1CustomGateway contract.
type L1CustomGatewayWithdrawRedirectedIterator struct {
	Event *L1CustomGatewayWithdrawRedirected // Event containing the contract specifics and raw log

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
func (it *L1CustomGatewayWithdrawRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomGatewayWithdrawRedirected)
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
		it.Event = new(L1CustomGatewayWithdrawRedirected)
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
func (it *L1CustomGatewayWithdrawRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomGatewayWithdrawRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomGatewayWithdrawRedirected represents a WithdrawRedirected event raised by the L1CustomGateway contract.
type L1CustomGatewayWithdrawRedirected struct {
	From             common.Address
	To               common.Address
	ExitNum          *big.Int
	NewData          []byte
	Data             []byte
	MadeExternalCall bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRedirected is a free log retrieval operation binding the contract event 0x56735ccb9dc7d2222ce4177fc3aea44c33882e2a2c73e0fb1c6b93c9c13a04d4.
//
// Solidity: event WithdrawRedirected(address indexed from, address indexed to, uint256 indexed exitNum, bytes newData, bytes data, bool madeExternalCall)
func (_L1CustomGateway *L1CustomGatewayFilterer) FilterWithdrawRedirected(opts *bind.FilterOpts, from []common.Address, to []common.Address, exitNum []*big.Int) (*L1CustomGatewayWithdrawRedirectedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var exitNumRule []interface{}
	for _, exitNumItem := range exitNum {
		exitNumRule = append(exitNumRule, exitNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.FilterLogs(opts, "WithdrawRedirected", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayWithdrawRedirectedIterator{contract: _L1CustomGateway.contract, event: "WithdrawRedirected", logs: logs, sub: sub}, nil
}

// WatchWithdrawRedirected is a free log subscription operation binding the contract event 0x56735ccb9dc7d2222ce4177fc3aea44c33882e2a2c73e0fb1c6b93c9c13a04d4.
//
// Solidity: event WithdrawRedirected(address indexed from, address indexed to, uint256 indexed exitNum, bytes newData, bytes data, bool madeExternalCall)
func (_L1CustomGateway *L1CustomGatewayFilterer) WatchWithdrawRedirected(opts *bind.WatchOpts, sink chan<- *L1CustomGatewayWithdrawRedirected, from []common.Address, to []common.Address, exitNum []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var exitNumRule []interface{}
	for _, exitNumItem := range exitNum {
		exitNumRule = append(exitNumRule, exitNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.WatchLogs(opts, "WithdrawRedirected", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomGatewayWithdrawRedirected)
				if err := _L1CustomGateway.contract.UnpackLog(event, "WithdrawRedirected", log); err != nil {
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

// ParseWithdrawRedirected is a log parse operation binding the contract event 0x56735ccb9dc7d2222ce4177fc3aea44c33882e2a2c73e0fb1c6b93c9c13a04d4.
//
// Solidity: event WithdrawRedirected(address indexed from, address indexed to, uint256 indexed exitNum, bytes newData, bytes data, bool madeExternalCall)
func (_L1CustomGateway *L1CustomGatewayFilterer) ParseWithdrawRedirected(log types.Log) (*L1CustomGatewayWithdrawRedirected, error) {
	event := new(L1CustomGatewayWithdrawRedirected)
	if err := _L1CustomGateway.contract.UnpackLog(event, "WithdrawRedirected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1CustomGatewayWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the L1CustomGateway contract.
type L1CustomGatewayWithdrawalFinalizedIterator struct {
	Event *L1CustomGatewayWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1CustomGatewayWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1CustomGatewayWithdrawalFinalized)
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
		it.Event = new(L1CustomGatewayWithdrawalFinalized)
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
func (it *L1CustomGatewayWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1CustomGatewayWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1CustomGatewayWithdrawalFinalized represents a WithdrawalFinalized event raised by the L1CustomGateway contract.
type L1CustomGatewayWithdrawalFinalized struct {
	L1Token common.Address
	From    common.Address
	To      common.Address
	ExitNum *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed _from, address indexed _to, uint256 indexed _exitNum, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _exitNum []*big.Int) (*L1CustomGatewayWithdrawalFinalizedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _exitNumRule []interface{}
	for _, _exitNumItem := range _exitNum {
		_exitNumRule = append(_exitNumRule, _exitNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.FilterLogs(opts, "WithdrawalFinalized", _fromRule, _toRule, _exitNumRule)
	if err != nil {
		return nil, err
	}
	return &L1CustomGatewayWithdrawalFinalizedIterator{contract: _L1CustomGateway.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed _from, address indexed _to, uint256 indexed _exitNum, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1CustomGatewayWithdrawalFinalized, _from []common.Address, _to []common.Address, _exitNum []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _exitNumRule []interface{}
	for _, _exitNumItem := range _exitNum {
		_exitNumRule = append(_exitNumRule, _exitNumItem)
	}

	logs, sub, err := _L1CustomGateway.contract.WatchLogs(opts, "WithdrawalFinalized", _fromRule, _toRule, _exitNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1CustomGatewayWithdrawalFinalized)
				if err := _L1CustomGateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed _from, address indexed _to, uint256 indexed _exitNum, uint256 _amount)
func (_L1CustomGateway *L1CustomGatewayFilterer) ParseWithdrawalFinalized(log types.Log) (*L1CustomGatewayWithdrawalFinalized, error) {
	event := new(L1CustomGatewayWithdrawalFinalized)
	if err := _L1CustomGateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
