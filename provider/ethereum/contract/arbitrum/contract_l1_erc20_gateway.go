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

// L1ERC20GatewayMetaData contains all meta data concerning the L1ERC20Gateway contract.
var L1ERC20GatewayMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_sequenceNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DepositInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_seqNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TxToL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"newData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"madeExternalCall\",\"type\":\"bool\"}],\"name\":\"WithdrawRedirected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"l1Token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalFinalized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"l1ERC20\",\"type\":\"address\"}],\"name\":\"calculateL2TokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cloneableProxyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"counterpartGateway\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"}],\"name\":\"encodeWithdrawal\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"finalizeInboundTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_initialData\",\"type\":\"bytes\"}],\"name\":\"getExternalCall\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"getOutboundCalldata\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"outboundCalldata\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inbox\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Counterpart\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_cloneableProxyHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_l2BeaconProxyFactory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2BeaconProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransfer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_refundTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_gasPriceBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"outboundTransferCustomRefund\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postUpgradeInit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"redirectedExits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isExit\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_newTo\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_newData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_exitNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_initialDestination\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newDestination\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_newData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"transferExitAndCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// L1ERC20GatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use L1ERC20GatewayMetaData.ABI instead.
var L1ERC20GatewayABI = L1ERC20GatewayMetaData.ABI

// L1ERC20Gateway is an auto generated Go binding around an Ethereum contract.
type L1ERC20Gateway struct {
	L1ERC20GatewayCaller     // Read-only binding to the contract
	L1ERC20GatewayTransactor // Write-only binding to the contract
	L1ERC20GatewayFilterer   // Log filterer for contract events
}

// L1ERC20GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type L1ERC20GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L1ERC20GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1ERC20GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1ERC20GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1ERC20GatewaySession struct {
	Contract     *L1ERC20Gateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1ERC20GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1ERC20GatewayCallerSession struct {
	Contract *L1ERC20GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// L1ERC20GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1ERC20GatewayTransactorSession struct {
	Contract     *L1ERC20GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// L1ERC20GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type L1ERC20GatewayRaw struct {
	Contract *L1ERC20Gateway // Generic contract binding to access the raw methods on
}

// L1ERC20GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1ERC20GatewayCallerRaw struct {
	Contract *L1ERC20GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// L1ERC20GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1ERC20GatewayTransactorRaw struct {
	Contract *L1ERC20GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL1ERC20Gateway creates a new instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20Gateway(address common.Address, backend bind.ContractBackend) (*L1ERC20Gateway, error) {
	contract, err := bindL1ERC20Gateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1ERC20Gateway{L1ERC20GatewayCaller: L1ERC20GatewayCaller{contract: contract}, L1ERC20GatewayTransactor: L1ERC20GatewayTransactor{contract: contract}, L1ERC20GatewayFilterer: L1ERC20GatewayFilterer{contract: contract}}, nil
}

// NewL1ERC20GatewayCaller creates a new read-only instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayCaller(address common.Address, caller bind.ContractCaller) (*L1ERC20GatewayCaller, error) {
	contract, err := bindL1ERC20Gateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayCaller{contract: contract}, nil
}

// NewL1ERC20GatewayTransactor creates a new write-only instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*L1ERC20GatewayTransactor, error) {
	contract, err := bindL1ERC20Gateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayTransactor{contract: contract}, nil
}

// NewL1ERC20GatewayFilterer creates a new log filterer instance of L1ERC20Gateway, bound to a specific deployed contract.
func NewL1ERC20GatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*L1ERC20GatewayFilterer, error) {
	contract, err := bindL1ERC20Gateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayFilterer{contract: contract}, nil
}

// bindL1ERC20Gateway binds a generic wrapper to an already deployed contract.
func bindL1ERC20Gateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := L1ERC20GatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Gateway *L1ERC20GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.L1ERC20GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1ERC20Gateway *L1ERC20GatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1ERC20Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1ERC20Gateway *L1ERC20GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1ERC20Gateway *L1ERC20GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.contract.Transact(opts, method, params...)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) CalculateL2TokenAddress(opts *bind.CallOpts, l1ERC20 common.Address) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "calculateL2TokenAddress", l1ERC20)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1ERC20Gateway.Contract.CalculateL2TokenAddress(&_L1ERC20Gateway.CallOpts, l1ERC20)
}

// CalculateL2TokenAddress is a free data retrieval call binding the contract method 0xa7e28d48.
//
// Solidity: function calculateL2TokenAddress(address l1ERC20) view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) CalculateL2TokenAddress(l1ERC20 common.Address) (common.Address, error) {
	return _L1ERC20Gateway.Contract.CalculateL2TokenAddress(&_L1ERC20Gateway.CallOpts, l1ERC20)
}

// CloneableProxyHash is a free data retrieval call binding the contract method 0x97881f8d.
//
// Solidity: function cloneableProxyHash() view returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) CloneableProxyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "cloneableProxyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CloneableProxyHash is a free data retrieval call binding the contract method 0x97881f8d.
//
// Solidity: function cloneableProxyHash() view returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewaySession) CloneableProxyHash() ([32]byte, error) {
	return _L1ERC20Gateway.Contract.CloneableProxyHash(&_L1ERC20Gateway.CallOpts)
}

// CloneableProxyHash is a free data retrieval call binding the contract method 0x97881f8d.
//
// Solidity: function cloneableProxyHash() view returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) CloneableProxyHash() ([32]byte, error) {
	return _L1ERC20Gateway.Contract.CloneableProxyHash(&_L1ERC20Gateway.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) CounterpartGateway(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "counterpartGateway")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) CounterpartGateway() (common.Address, error) {
	return _L1ERC20Gateway.Contract.CounterpartGateway(&_L1ERC20Gateway.CallOpts)
}

// CounterpartGateway is a free data retrieval call binding the contract method 0x2db09c1c.
//
// Solidity: function counterpartGateway() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) CounterpartGateway() (common.Address, error) {
	return _L1ERC20Gateway.Contract.CounterpartGateway(&_L1ERC20Gateway.CallOpts)
}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) EncodeWithdrawal(opts *bind.CallOpts, _exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "encodeWithdrawal", _exitNum, _initialDestination)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewaySession) EncodeWithdrawal(_exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	return _L1ERC20Gateway.Contract.EncodeWithdrawal(&_L1ERC20Gateway.CallOpts, _exitNum, _initialDestination)
}

// EncodeWithdrawal is a free data retrieval call binding the contract method 0x020a6058.
//
// Solidity: function encodeWithdrawal(uint256 _exitNum, address _initialDestination) pure returns(bytes32)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) EncodeWithdrawal(_exitNum *big.Int, _initialDestination common.Address) ([32]byte, error) {
	return _L1ERC20Gateway.Contract.EncodeWithdrawal(&_L1ERC20Gateway.CallOpts, _exitNum, _initialDestination)
}

// GetExternalCall is a free data retrieval call binding the contract method 0xf68a9082.
//
// Solidity: function getExternalCall(uint256 _exitNum, address _initialDestination, bytes _initialData) view returns(address target, bytes data)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) GetExternalCall(opts *bind.CallOpts, _exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "getExternalCall", _exitNum, _initialDestination, _initialData)

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
func (_L1ERC20Gateway *L1ERC20GatewaySession) GetExternalCall(_exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	return _L1ERC20Gateway.Contract.GetExternalCall(&_L1ERC20Gateway.CallOpts, _exitNum, _initialDestination, _initialData)
}

// GetExternalCall is a free data retrieval call binding the contract method 0xf68a9082.
//
// Solidity: function getExternalCall(uint256 _exitNum, address _initialDestination, bytes _initialData) view returns(address target, bytes data)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) GetExternalCall(_exitNum *big.Int, _initialDestination common.Address, _initialData []byte) (struct {
	Target common.Address
	Data   []byte
}, error) {
	return _L1ERC20Gateway.Contract.GetExternalCall(&_L1ERC20Gateway.CallOpts, _exitNum, _initialDestination, _initialData)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) GetOutboundCalldata(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "getOutboundCalldata", _token, _from, _to, _amount, _data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1ERC20Gateway *L1ERC20GatewaySession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1ERC20Gateway.Contract.GetOutboundCalldata(&_L1ERC20Gateway.CallOpts, _token, _from, _to, _amount, _data)
}

// GetOutboundCalldata is a free data retrieval call binding the contract method 0xa0c76a96.
//
// Solidity: function getOutboundCalldata(address _token, address _from, address _to, uint256 _amount, bytes _data) view returns(bytes outboundCalldata)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) GetOutboundCalldata(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) ([]byte, error) {
	return _L1ERC20Gateway.Contract.GetOutboundCalldata(&_L1ERC20Gateway.CallOpts, _token, _from, _to, _amount, _data)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) Inbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "inbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) Inbox() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Inbox(&_L1ERC20Gateway.CallOpts)
}

// Inbox is a free data retrieval call binding the contract method 0xfb0e722b.
//
// Solidity: function inbox() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) Inbox() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Inbox(&_L1ERC20Gateway.CallOpts)
}

// L2BeaconProxyFactory is a free data retrieval call binding the contract method 0x70fc045f.
//
// Solidity: function l2BeaconProxyFactory() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) L2BeaconProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "l2BeaconProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2BeaconProxyFactory is a free data retrieval call binding the contract method 0x70fc045f.
//
// Solidity: function l2BeaconProxyFactory() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) L2BeaconProxyFactory() (common.Address, error) {
	return _L1ERC20Gateway.Contract.L2BeaconProxyFactory(&_L1ERC20Gateway.CallOpts)
}

// L2BeaconProxyFactory is a free data retrieval call binding the contract method 0x70fc045f.
//
// Solidity: function l2BeaconProxyFactory() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) L2BeaconProxyFactory() (common.Address, error) {
	return _L1ERC20Gateway.Contract.L2BeaconProxyFactory(&_L1ERC20Gateway.CallOpts)
}

// RedirectedExits is a free data retrieval call binding the contract method 0xbcf2e6eb.
//
// Solidity: function redirectedExits(bytes32 ) view returns(bool isExit, address _newTo, bytes _newData)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) RedirectedExits(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "redirectedExits", arg0)

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
func (_L1ERC20Gateway *L1ERC20GatewaySession) RedirectedExits(arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	return _L1ERC20Gateway.Contract.RedirectedExits(&_L1ERC20Gateway.CallOpts, arg0)
}

// RedirectedExits is a free data retrieval call binding the contract method 0xbcf2e6eb.
//
// Solidity: function redirectedExits(bytes32 ) view returns(bool isExit, address _newTo, bytes _newData)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) RedirectedExits(arg0 [32]byte) (struct {
	IsExit  bool
	NewTo   common.Address
	NewData []byte
}, error) {
	return _L1ERC20Gateway.Contract.RedirectedExits(&_L1ERC20Gateway.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) Router() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Router(&_L1ERC20Gateway.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) Router() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Router(&_L1ERC20Gateway.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1ERC20Gateway *L1ERC20GatewaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1ERC20Gateway.Contract.SupportsInterface(&_L1ERC20Gateway.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _L1ERC20Gateway.Contract.SupportsInterface(&_L1ERC20Gateway.CallOpts, interfaceId)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCaller) Whitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1ERC20Gateway.contract.Call(opts, &out, "whitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewaySession) Whitelist() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Whitelist(&_L1ERC20Gateway.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x93e59dc1.
//
// Solidity: function whitelist() view returns(address)
func (_L1ERC20Gateway *L1ERC20GatewayCallerSession) Whitelist() (common.Address, error) {
	return _L1ERC20Gateway.Contract.Whitelist(&_L1ERC20Gateway.CallOpts)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) FinalizeInboundTransfer(opts *bind.TransactOpts, _token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "finalizeInboundTransfer", _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.FinalizeInboundTransfer(&_L1ERC20Gateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// FinalizeInboundTransfer is a paid mutator transaction binding the contract method 0x2e567b36.
//
// Solidity: function finalizeInboundTransfer(address _token, address _from, address _to, uint256 _amount, bytes _data) payable returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) FinalizeInboundTransfer(_token common.Address, _from common.Address, _to common.Address, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.FinalizeInboundTransfer(&_L1ERC20Gateway.TransactOpts, _token, _from, _to, _amount, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01893bf.
//
// Solidity: function initialize(address _l2Counterpart, address _router, address _inbox, bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) Initialize(opts *bind.TransactOpts, _l2Counterpart common.Address, _router common.Address, _inbox common.Address, _cloneableProxyHash [32]byte, _l2BeaconProxyFactory common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "initialize", _l2Counterpart, _router, _inbox, _cloneableProxyHash, _l2BeaconProxyFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01893bf.
//
// Solidity: function initialize(address _l2Counterpart, address _router, address _inbox, bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory) returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) Initialize(_l2Counterpart common.Address, _router common.Address, _inbox common.Address, _cloneableProxyHash [32]byte, _l2BeaconProxyFactory common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Initialize(&_L1ERC20Gateway.TransactOpts, _l2Counterpart, _router, _inbox, _cloneableProxyHash, _l2BeaconProxyFactory)
}

// Initialize is a paid mutator transaction binding the contract method 0xa01893bf.
//
// Solidity: function initialize(address _l2Counterpart, address _router, address _inbox, bytes32 _cloneableProxyHash, address _l2BeaconProxyFactory) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) Initialize(_l2Counterpart common.Address, _router common.Address, _inbox common.Address, _cloneableProxyHash [32]byte, _l2BeaconProxyFactory common.Address) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.Initialize(&_L1ERC20Gateway.TransactOpts, _l2Counterpart, _router, _inbox, _cloneableProxyHash, _l2BeaconProxyFactory)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) OutboundTransfer(opts *bind.TransactOpts, _l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "outboundTransfer", _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewaySession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.OutboundTransfer(&_L1ERC20Gateway.TransactOpts, _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransfer is a paid mutator transaction binding the contract method 0xd2ce7d65.
//
// Solidity: function outboundTransfer(address _l1Token, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) OutboundTransfer(_l1Token common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.OutboundTransfer(&_L1ERC20Gateway.TransactOpts, _l1Token, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) OutboundTransferCustomRefund(opts *bind.TransactOpts, _l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "outboundTransferCustomRefund", _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewaySession) OutboundTransferCustomRefund(_l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.OutboundTransferCustomRefund(&_L1ERC20Gateway.TransactOpts, _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// OutboundTransferCustomRefund is a paid mutator transaction binding the contract method 0x4fb1a07b.
//
// Solidity: function outboundTransferCustomRefund(address _l1Token, address _refundTo, address _to, uint256 _amount, uint256 _maxGas, uint256 _gasPriceBid, bytes _data) payable returns(bytes res)
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) OutboundTransferCustomRefund(_l1Token common.Address, _refundTo common.Address, _to common.Address, _amount *big.Int, _maxGas *big.Int, _gasPriceBid *big.Int, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.OutboundTransferCustomRefund(&_L1ERC20Gateway.TransactOpts, _l1Token, _refundTo, _to, _amount, _maxGas, _gasPriceBid, _data)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) PostUpgradeInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "postUpgradeInit")
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.PostUpgradeInit(&_L1ERC20Gateway.TransactOpts)
}

// PostUpgradeInit is a paid mutator transaction binding the contract method 0x95fcea78.
//
// Solidity: function postUpgradeInit() returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) PostUpgradeInit() (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.PostUpgradeInit(&_L1ERC20Gateway.TransactOpts)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactor) TransferExitAndCall(opts *bind.TransactOpts, _exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.contract.Transact(opts, "transferExitAndCall", _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1ERC20Gateway *L1ERC20GatewaySession) TransferExitAndCall(_exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.TransferExitAndCall(&_L1ERC20Gateway.TransactOpts, _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// TransferExitAndCall is a paid mutator transaction binding the contract method 0xbd5f3e7d.
//
// Solidity: function transferExitAndCall(uint256 _exitNum, address _initialDestination, address _newDestination, bytes _newData, bytes _data) returns()
func (_L1ERC20Gateway *L1ERC20GatewayTransactorSession) TransferExitAndCall(_exitNum *big.Int, _initialDestination common.Address, _newDestination common.Address, _newData []byte, _data []byte) (*types.Transaction, error) {
	return _L1ERC20Gateway.Contract.TransferExitAndCall(&_L1ERC20Gateway.TransactOpts, _exitNum, _initialDestination, _newDestination, _newData, _data)
}

// L1ERC20GatewayDepositInitiatedIterator is returned from FilterDepositInitiated and is used to iterate over the raw logs and unpacked data for DepositInitiated events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayDepositInitiatedIterator struct {
	Event *L1ERC20GatewayDepositInitiated // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayDepositInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayDepositInitiated)
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
		it.Event = new(L1ERC20GatewayDepositInitiated)
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
func (it *L1ERC20GatewayDepositInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayDepositInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayDepositInitiated represents a DepositInitiated event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayDepositInitiated struct {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterDepositInitiated(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _sequenceNumber []*big.Int) (*L1ERC20GatewayDepositInitiatedIterator, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "DepositInitiated", _fromRule, _toRule, _sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayDepositInitiatedIterator{contract: _L1ERC20Gateway.contract, event: "DepositInitiated", logs: logs, sub: sub}, nil
}

// WatchDepositInitiated is a free log subscription operation binding the contract event 0xb8910b9960c443aac3240b98585384e3a6f109fbf6969e264c3f183d69aba7e1.
//
// Solidity: event DepositInitiated(address l1Token, address indexed _from, address indexed _to, uint256 indexed _sequenceNumber, uint256 _amount)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchDepositInitiated(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayDepositInitiated, _from []common.Address, _to []common.Address, _sequenceNumber []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "DepositInitiated", _fromRule, _toRule, _sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayDepositInitiated)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseDepositInitiated(log types.Log) (*L1ERC20GatewayDepositInitiated, error) {
	event := new(L1ERC20GatewayDepositInitiated)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "DepositInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20GatewayTxToL2Iterator is returned from FilterTxToL2 and is used to iterate over the raw logs and unpacked data for TxToL2 events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayTxToL2Iterator struct {
	Event *L1ERC20GatewayTxToL2 // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayTxToL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayTxToL2)
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
		it.Event = new(L1ERC20GatewayTxToL2)
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
func (it *L1ERC20GatewayTxToL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayTxToL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayTxToL2 represents a TxToL2 event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayTxToL2 struct {
	From   common.Address
	To     common.Address
	SeqNum *big.Int
	Data   []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTxToL2 is a free log retrieval operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterTxToL2(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (*L1ERC20GatewayTxToL2Iterator, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayTxToL2Iterator{contract: _L1ERC20Gateway.contract, event: "TxToL2", logs: logs, sub: sub}, nil
}

// WatchTxToL2 is a free log subscription operation binding the contract event 0xc1d1490cf25c3b40d600dfb27c7680340ed1ab901b7e8f3551280968a3b372b0.
//
// Solidity: event TxToL2(address indexed _from, address indexed _to, uint256 indexed _seqNum, bytes _data)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchTxToL2(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayTxToL2, _from []common.Address, _to []common.Address, _seqNum []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "TxToL2", _fromRule, _toRule, _seqNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayTxToL2)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "TxToL2", log); err != nil {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseTxToL2(log types.Log) (*L1ERC20GatewayTxToL2, error) {
	event := new(L1ERC20GatewayTxToL2)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "TxToL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20GatewayWithdrawRedirectedIterator is returned from FilterWithdrawRedirected and is used to iterate over the raw logs and unpacked data for WithdrawRedirected events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayWithdrawRedirectedIterator struct {
	Event *L1ERC20GatewayWithdrawRedirected // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayWithdrawRedirectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayWithdrawRedirected)
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
		it.Event = new(L1ERC20GatewayWithdrawRedirected)
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
func (it *L1ERC20GatewayWithdrawRedirectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayWithdrawRedirectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayWithdrawRedirected represents a WithdrawRedirected event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayWithdrawRedirected struct {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterWithdrawRedirected(opts *bind.FilterOpts, from []common.Address, to []common.Address, exitNum []*big.Int) (*L1ERC20GatewayWithdrawRedirectedIterator, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "WithdrawRedirected", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayWithdrawRedirectedIterator{contract: _L1ERC20Gateway.contract, event: "WithdrawRedirected", logs: logs, sub: sub}, nil
}

// WatchWithdrawRedirected is a free log subscription operation binding the contract event 0x56735ccb9dc7d2222ce4177fc3aea44c33882e2a2c73e0fb1c6b93c9c13a04d4.
//
// Solidity: event WithdrawRedirected(address indexed from, address indexed to, uint256 indexed exitNum, bytes newData, bytes data, bool madeExternalCall)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchWithdrawRedirected(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayWithdrawRedirected, from []common.Address, to []common.Address, exitNum []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "WithdrawRedirected", fromRule, toRule, exitNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayWithdrawRedirected)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "WithdrawRedirected", log); err != nil {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseWithdrawRedirected(log types.Log) (*L1ERC20GatewayWithdrawRedirected, error) {
	event := new(L1ERC20GatewayWithdrawRedirected)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "WithdrawRedirected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1ERC20GatewayWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the L1ERC20Gateway contract.
type L1ERC20GatewayWithdrawalFinalizedIterator struct {
	Event *L1ERC20GatewayWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *L1ERC20GatewayWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1ERC20GatewayWithdrawalFinalized)
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
		it.Event = new(L1ERC20GatewayWithdrawalFinalized)
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
func (it *L1ERC20GatewayWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ERC20GatewayWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1ERC20GatewayWithdrawalFinalized represents a WithdrawalFinalized event raised by the L1ERC20Gateway contract.
type L1ERC20GatewayWithdrawalFinalized struct {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _exitNum []*big.Int) (*L1ERC20GatewayWithdrawalFinalizedIterator, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.FilterLogs(opts, "WithdrawalFinalized", _fromRule, _toRule, _exitNumRule)
	if err != nil {
		return nil, err
	}
	return &L1ERC20GatewayWithdrawalFinalizedIterator{contract: _L1ERC20Gateway.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0x891afe029c75c4f8c5855fc3480598bc5a53739344f6ae575bdb7ea2a79f56b3.
//
// Solidity: event WithdrawalFinalized(address l1Token, address indexed _from, address indexed _to, uint256 indexed _exitNum, uint256 _amount)
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *L1ERC20GatewayWithdrawalFinalized, _from []common.Address, _to []common.Address, _exitNum []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _L1ERC20Gateway.contract.WatchLogs(opts, "WithdrawalFinalized", _fromRule, _toRule, _exitNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1ERC20GatewayWithdrawalFinalized)
				if err := _L1ERC20Gateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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
func (_L1ERC20Gateway *L1ERC20GatewayFilterer) ParseWithdrawalFinalized(log types.Log) (*L1ERC20GatewayWithdrawalFinalized, error) {
	event := new(L1ERC20GatewayWithdrawalFinalized)
	if err := _L1ERC20Gateway.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
