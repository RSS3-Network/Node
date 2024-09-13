// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nouns

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

// INounsSeederSeed is an auto generated low-level Go binding around an user-defined struct.
type INounsSeederSeed struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}

// NounsTokenMetaData contains all meta data concerning the NounsToken contract.
var NounsTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_noundersDAO\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"},{\"internalType\":\"contractINounsDescriptor\",\"name\":\"_descriptor\",\"type\":\"address\"},{\"internalType\":\"contractINounsSeeder\",\"name\":\"_seeder\",\"type\":\"address\"},{\"internalType\":\"contractIProxyRegistry\",\"name\":\"_proxyRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"DescriptorLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractINounsDescriptor\",\"name\":\"descriptor\",\"type\":\"address\"}],\"name\":\"DescriptorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MinterLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"MinterUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NounBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint48\",\"name\":\"background\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"body\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"accessory\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"head\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"glasses\",\"type\":\"uint48\"}],\"indexed\":false,\"internalType\":\"structINounsSeeder.Seed\",\"name\":\"seed\",\"type\":\"tuple\"}],\"name\":\"NounCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"noundersDAO\",\"type\":\"address\"}],\"name\":\"NoundersDAOUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"SeederLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractINounsSeeder\",\"name\":\"seeder\",\"type\":\"address\"}],\"name\":\"SeederUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"dataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"descriptor\",\"outputs\":[{\"internalType\":\"contractINounsDescriptor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDescriptorLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMinterLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSeederLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockSeeder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noundersDAO\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyRegistry\",\"outputs\":[{\"internalType\":\"contractIProxyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seeder\",\"outputs\":[{\"internalType\":\"contractINounsSeeder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"seeds\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"background\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"body\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"accessory\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"head\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"glasses\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newContractURIHash\",\"type\":\"string\"}],\"name\":\"setContractURIHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINounsDescriptor\",\"name\":\"_descriptor\",\"type\":\"address\"}],\"name\":\"setDescriptor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_noundersDAO\",\"type\":\"address\"}],\"name\":\"setNoundersDAO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINounsSeeder\",\"name\":\"_seeder\",\"type\":\"address\"}],\"name\":\"setSeeder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"votesToDelegate\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NounsTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use NounsTokenMetaData.ABI instead.
var NounsTokenABI = NounsTokenMetaData.ABI

// NounsToken is an auto generated Go binding around an Ethereum contract.
type NounsToken struct {
	NounsTokenCaller     // Read-only binding to the contract
	NounsTokenTransactor // Write-only binding to the contract
	NounsTokenFilterer   // Log filterer for contract events
}

// NounsTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type NounsTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NounsTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NounsTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NounsTokenSession struct {
	Contract     *NounsToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NounsTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NounsTokenCallerSession struct {
	Contract *NounsTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NounsTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NounsTokenTransactorSession struct {
	Contract     *NounsTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NounsTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type NounsTokenRaw struct {
	Contract *NounsToken // Generic contract binding to access the raw methods on
}

// NounsTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NounsTokenCallerRaw struct {
	Contract *NounsTokenCaller // Generic read-only contract binding to access the raw methods on
}

// NounsTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NounsTokenTransactorRaw struct {
	Contract *NounsTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNounsToken creates a new instance of NounsToken, bound to a specific deployed contract.
func NewNounsToken(address common.Address, backend bind.ContractBackend) (*NounsToken, error) {
	contract, err := bindNounsToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NounsToken{NounsTokenCaller: NounsTokenCaller{contract: contract}, NounsTokenTransactor: NounsTokenTransactor{contract: contract}, NounsTokenFilterer: NounsTokenFilterer{contract: contract}}, nil
}

// NewNounsTokenCaller creates a new read-only instance of NounsToken, bound to a specific deployed contract.
func NewNounsTokenCaller(address common.Address, caller bind.ContractCaller) (*NounsTokenCaller, error) {
	contract, err := bindNounsToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NounsTokenCaller{contract: contract}, nil
}

// NewNounsTokenTransactor creates a new write-only instance of NounsToken, bound to a specific deployed contract.
func NewNounsTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*NounsTokenTransactor, error) {
	contract, err := bindNounsToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NounsTokenTransactor{contract: contract}, nil
}

// NewNounsTokenFilterer creates a new log filterer instance of NounsToken, bound to a specific deployed contract.
func NewNounsTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*NounsTokenFilterer, error) {
	contract, err := bindNounsToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NounsTokenFilterer{contract: contract}, nil
}

// bindNounsToken binds a generic wrapper to an already deployed contract.
func bindNounsToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NounsTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsToken *NounsTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsToken.Contract.NounsTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsToken *NounsTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.Contract.NounsTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsToken *NounsTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsToken.Contract.NounsTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsToken *NounsTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsToken *NounsTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsToken *NounsTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsToken.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _NounsToken.Contract.DELEGATIONTYPEHASH(&_NounsToken.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _NounsToken.Contract.DELEGATIONTYPEHASH(&_NounsToken.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NounsToken.Contract.DOMAINTYPEHASH(&_NounsToken.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsToken *NounsTokenCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NounsToken.Contract.DOMAINTYPEHASH(&_NounsToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NounsToken *NounsTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NounsToken *NounsTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NounsToken.Contract.BalanceOf(&_NounsToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NounsToken *NounsTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NounsToken.Contract.BalanceOf(&_NounsToken.CallOpts, owner)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_NounsToken *NounsTokenCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_NounsToken *NounsTokenSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _NounsToken.Contract.Checkpoints(&_NounsToken.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint96 votes)
func (_NounsToken *NounsTokenCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _NounsToken.Contract.Checkpoints(&_NounsToken.CallOpts, arg0, arg1)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NounsToken *NounsTokenCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NounsToken *NounsTokenSession) ContractURI() (string, error) {
	return _NounsToken.Contract.ContractURI(&_NounsToken.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NounsToken *NounsTokenCallerSession) ContractURI() (string, error) {
	return _NounsToken.Contract.ContractURI(&_NounsToken.CallOpts)
}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenCaller) DataURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "dataURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenSession) DataURI(tokenId *big.Int) (string, error) {
	return _NounsToken.Contract.DataURI(&_NounsToken.CallOpts, tokenId)
}

// DataURI is a free data retrieval call binding the contract method 0x5ac1e3bb.
//
// Solidity: function dataURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenCallerSession) DataURI(tokenId *big.Int) (string, error) {
	return _NounsToken.Contract.DataURI(&_NounsToken.CallOpts, tokenId)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NounsToken *NounsTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NounsToken *NounsTokenSession) Decimals() (uint8, error) {
	return _NounsToken.Contract.Decimals(&_NounsToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_NounsToken *NounsTokenCallerSession) Decimals() (uint8, error) {
	return _NounsToken.Contract.Decimals(&_NounsToken.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_NounsToken *NounsTokenCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_NounsToken *NounsTokenSession) Delegates(delegator common.Address) (common.Address, error) {
	return _NounsToken.Contract.Delegates(&_NounsToken.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_NounsToken *NounsTokenCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _NounsToken.Contract.Delegates(&_NounsToken.CallOpts, delegator)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsToken *NounsTokenCaller) Descriptor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "descriptor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsToken *NounsTokenSession) Descriptor() (common.Address, error) {
	return _NounsToken.Contract.Descriptor(&_NounsToken.CallOpts)
}

// Descriptor is a free data retrieval call binding the contract method 0x303e74df.
//
// Solidity: function descriptor() view returns(address)
func (_NounsToken *NounsTokenCallerSession) Descriptor() (common.Address, error) {
	return _NounsToken.Contract.Descriptor(&_NounsToken.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NounsToken.Contract.GetApproved(&_NounsToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NounsToken.Contract.GetApproved(&_NounsToken.CallOpts, tokenId)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_NounsToken *NounsTokenCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_NounsToken *NounsTokenSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _NounsToken.Contract.GetCurrentVotes(&_NounsToken.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint96)
func (_NounsToken *NounsTokenCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _NounsToken.Contract.GetCurrentVotes(&_NounsToken.CallOpts, account)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_NounsToken *NounsTokenCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_NounsToken *NounsTokenSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.GetPriorVotes(&_NounsToken.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint96)
func (_NounsToken *NounsTokenCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.GetPriorVotes(&_NounsToken.CallOpts, account, blockNumber)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NounsToken *NounsTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NounsToken *NounsTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NounsToken.Contract.IsApprovedForAll(&_NounsToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NounsToken *NounsTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NounsToken.Contract.IsApprovedForAll(&_NounsToken.CallOpts, owner, operator)
}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_NounsToken *NounsTokenCaller) IsDescriptorLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "isDescriptorLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_NounsToken *NounsTokenSession) IsDescriptorLocked() (bool, error) {
	return _NounsToken.Contract.IsDescriptorLocked(&_NounsToken.CallOpts)
}

// IsDescriptorLocked is a free data retrieval call binding the contract method 0xc1b8e4e1.
//
// Solidity: function isDescriptorLocked() view returns(bool)
func (_NounsToken *NounsTokenCallerSession) IsDescriptorLocked() (bool, error) {
	return _NounsToken.Contract.IsDescriptorLocked(&_NounsToken.CallOpts)
}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_NounsToken *NounsTokenCaller) IsMinterLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "isMinterLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_NounsToken *NounsTokenSession) IsMinterLocked() (bool, error) {
	return _NounsToken.Contract.IsMinterLocked(&_NounsToken.CallOpts)
}

// IsMinterLocked is a free data retrieval call binding the contract method 0x1e688e10.
//
// Solidity: function isMinterLocked() view returns(bool)
func (_NounsToken *NounsTokenCallerSession) IsMinterLocked() (bool, error) {
	return _NounsToken.Contract.IsMinterLocked(&_NounsToken.CallOpts)
}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_NounsToken *NounsTokenCaller) IsSeederLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "isSeederLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_NounsToken *NounsTokenSession) IsSeederLocked() (bool, error) {
	return _NounsToken.Contract.IsSeederLocked(&_NounsToken.CallOpts)
}

// IsSeederLocked is a free data retrieval call binding the contract method 0xc8fc0c23.
//
// Solidity: function isSeederLocked() view returns(bool)
func (_NounsToken *NounsTokenCallerSession) IsSeederLocked() (bool, error) {
	return _NounsToken.Contract.IsSeederLocked(&_NounsToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NounsToken *NounsTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NounsToken *NounsTokenSession) Minter() (common.Address, error) {
	return _NounsToken.Contract.Minter(&_NounsToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_NounsToken *NounsTokenCallerSession) Minter() (common.Address, error) {
	return _NounsToken.Contract.Minter(&_NounsToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsToken *NounsTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsToken *NounsTokenSession) Name() (string, error) {
	return _NounsToken.Contract.Name(&_NounsToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsToken *NounsTokenCallerSession) Name() (string, error) {
	return _NounsToken.Contract.Name(&_NounsToken.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NounsToken *NounsTokenCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NounsToken *NounsTokenSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NounsToken.Contract.Nonces(&_NounsToken.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NounsToken *NounsTokenCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NounsToken.Contract.Nonces(&_NounsToken.CallOpts, arg0)
}

// NoundersDAO is a free data retrieval call binding the contract method 0x655932a4.
//
// Solidity: function noundersDAO() view returns(address)
func (_NounsToken *NounsTokenCaller) NoundersDAO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "noundersDAO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NoundersDAO is a free data retrieval call binding the contract method 0x655932a4.
//
// Solidity: function noundersDAO() view returns(address)
func (_NounsToken *NounsTokenSession) NoundersDAO() (common.Address, error) {
	return _NounsToken.Contract.NoundersDAO(&_NounsToken.CallOpts)
}

// NoundersDAO is a free data retrieval call binding the contract method 0x655932a4.
//
// Solidity: function noundersDAO() view returns(address)
func (_NounsToken *NounsTokenCallerSession) NoundersDAO() (common.Address, error) {
	return _NounsToken.Contract.NoundersDAO(&_NounsToken.CallOpts)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_NounsToken *NounsTokenCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_NounsToken *NounsTokenSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _NounsToken.Contract.NumCheckpoints(&_NounsToken.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_NounsToken *NounsTokenCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _NounsToken.Contract.NumCheckpoints(&_NounsToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsToken *NounsTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsToken *NounsTokenSession) Owner() (common.Address, error) {
	return _NounsToken.Contract.Owner(&_NounsToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsToken *NounsTokenCallerSession) Owner() (common.Address, error) {
	return _NounsToken.Contract.Owner(&_NounsToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NounsToken.Contract.OwnerOf(&_NounsToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NounsToken *NounsTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NounsToken.Contract.OwnerOf(&_NounsToken.CallOpts, tokenId)
}

// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
//
// Solidity: function proxyRegistry() view returns(address)
func (_NounsToken *NounsTokenCaller) ProxyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "proxyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
//
// Solidity: function proxyRegistry() view returns(address)
func (_NounsToken *NounsTokenSession) ProxyRegistry() (common.Address, error) {
	return _NounsToken.Contract.ProxyRegistry(&_NounsToken.CallOpts)
}

// ProxyRegistry is a free data retrieval call binding the contract method 0xb50cbd9f.
//
// Solidity: function proxyRegistry() view returns(address)
func (_NounsToken *NounsTokenCallerSession) ProxyRegistry() (common.Address, error) {
	return _NounsToken.Contract.ProxyRegistry(&_NounsToken.CallOpts)
}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_NounsToken *NounsTokenCaller) Seeder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "seeder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_NounsToken *NounsTokenSession) Seeder() (common.Address, error) {
	return _NounsToken.Contract.Seeder(&_NounsToken.CallOpts)
}

// Seeder is a free data retrieval call binding the contract method 0x684931ed.
//
// Solidity: function seeder() view returns(address)
func (_NounsToken *NounsTokenCallerSession) Seeder() (common.Address, error) {
	return _NounsToken.Contract.Seeder(&_NounsToken.CallOpts)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_NounsToken *NounsTokenCaller) Seeds(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "seeds", arg0)

	outstruct := new(struct {
		Background *big.Int
		Body       *big.Int
		Accessory  *big.Int
		Head       *big.Int
		Glasses    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Background = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Body = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Accessory = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Head = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Glasses = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_NounsToken *NounsTokenSession) Seeds(arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	return _NounsToken.Contract.Seeds(&_NounsToken.CallOpts, arg0)
}

// Seeds is a free data retrieval call binding the contract method 0xf0503e80.
//
// Solidity: function seeds(uint256 ) view returns(uint48 background, uint48 body, uint48 accessory, uint48 head, uint48 glasses)
func (_NounsToken *NounsTokenCallerSession) Seeds(arg0 *big.Int) (struct {
	Background *big.Int
	Body       *big.Int
	Accessory  *big.Int
	Head       *big.Int
	Glasses    *big.Int
}, error) {
	return _NounsToken.Contract.Seeds(&_NounsToken.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NounsToken *NounsTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NounsToken *NounsTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NounsToken.Contract.SupportsInterface(&_NounsToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NounsToken *NounsTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NounsToken.Contract.SupportsInterface(&_NounsToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NounsToken *NounsTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NounsToken *NounsTokenSession) Symbol() (string, error) {
	return _NounsToken.Contract.Symbol(&_NounsToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NounsToken *NounsTokenCallerSession) Symbol() (string, error) {
	return _NounsToken.Contract.Symbol(&_NounsToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.TokenByIndex(&_NounsToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.TokenByIndex(&_NounsToken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.TokenOfOwnerByIndex(&_NounsToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NounsToken *NounsTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NounsToken.Contract.TokenOfOwnerByIndex(&_NounsToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NounsToken.Contract.TokenURI(&_NounsToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NounsToken *NounsTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NounsToken.Contract.TokenURI(&_NounsToken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NounsToken *NounsTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NounsToken *NounsTokenSession) TotalSupply() (*big.Int, error) {
	return _NounsToken.Contract.TotalSupply(&_NounsToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NounsToken *NounsTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _NounsToken.Contract.TotalSupply(&_NounsToken.CallOpts)
}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_NounsToken *NounsTokenCaller) VotesToDelegate(opts *bind.CallOpts, delegator common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NounsToken.contract.Call(opts, &out, "votesToDelegate", delegator)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_NounsToken *NounsTokenSession) VotesToDelegate(delegator common.Address) (*big.Int, error) {
	return _NounsToken.Contract.VotesToDelegate(&_NounsToken.CallOpts, delegator)
}

// VotesToDelegate is a free data retrieval call binding the contract method 0xe9580e91.
//
// Solidity: function votesToDelegate(address delegator) view returns(uint96)
func (_NounsToken *NounsTokenCallerSession) VotesToDelegate(delegator common.Address) (*big.Int, error) {
	return _NounsToken.Contract.VotesToDelegate(&_NounsToken.CallOpts, delegator)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.Approve(&_NounsToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.Approve(&_NounsToken.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_NounsToken *NounsTokenTransactor) Burn(opts *bind.TransactOpts, nounId *big.Int) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "burn", nounId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_NounsToken *NounsTokenSession) Burn(nounId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.Burn(&_NounsToken.TransactOpts, nounId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 nounId) returns()
func (_NounsToken *NounsTokenTransactorSession) Burn(nounId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.Burn(&_NounsToken.TransactOpts, nounId)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_NounsToken *NounsTokenTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_NounsToken *NounsTokenSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.Delegate(&_NounsToken.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_NounsToken *NounsTokenTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.Delegate(&_NounsToken.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsToken *NounsTokenTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsToken *NounsTokenSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsToken.Contract.DelegateBySig(&_NounsToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsToken *NounsTokenTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsToken.Contract.DelegateBySig(&_NounsToken.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_NounsToken *NounsTokenTransactor) LockDescriptor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "lockDescriptor")
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_NounsToken *NounsTokenSession) LockDescriptor() (*types.Transaction, error) {
	return _NounsToken.Contract.LockDescriptor(&_NounsToken.TransactOpts)
}

// LockDescriptor is a paid mutator transaction binding the contract method 0x41b5d0de.
//
// Solidity: function lockDescriptor() returns()
func (_NounsToken *NounsTokenTransactorSession) LockDescriptor() (*types.Transaction, error) {
	return _NounsToken.Contract.LockDescriptor(&_NounsToken.TransactOpts)
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_NounsToken *NounsTokenTransactor) LockMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "lockMinter")
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_NounsToken *NounsTokenSession) LockMinter() (*types.Transaction, error) {
	return _NounsToken.Contract.LockMinter(&_NounsToken.TransactOpts)
}

// LockMinter is a paid mutator transaction binding the contract method 0x76daebe1.
//
// Solidity: function lockMinter() returns()
func (_NounsToken *NounsTokenTransactorSession) LockMinter() (*types.Transaction, error) {
	return _NounsToken.Contract.LockMinter(&_NounsToken.TransactOpts)
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_NounsToken *NounsTokenTransactor) LockSeeder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "lockSeeder")
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_NounsToken *NounsTokenSession) LockSeeder() (*types.Transaction, error) {
	return _NounsToken.Contract.LockSeeder(&_NounsToken.TransactOpts)
}

// LockSeeder is a paid mutator transaction binding the contract method 0x5f295a67.
//
// Solidity: function lockSeeder() returns()
func (_NounsToken *NounsTokenTransactorSession) LockSeeder() (*types.Transaction, error) {
	return _NounsToken.Contract.LockSeeder(&_NounsToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_NounsToken *NounsTokenTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_NounsToken *NounsTokenSession) Mint() (*types.Transaction, error) {
	return _NounsToken.Contract.Mint(&_NounsToken.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() returns(uint256)
func (_NounsToken *NounsTokenTransactorSession) Mint() (*types.Transaction, error) {
	return _NounsToken.Contract.Mint(&_NounsToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsToken *NounsTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsToken *NounsTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _NounsToken.Contract.RenounceOwnership(&_NounsToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsToken *NounsTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NounsToken.Contract.RenounceOwnership(&_NounsToken.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.SafeTransferFrom(&_NounsToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.SafeTransferFrom(&_NounsToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NounsToken *NounsTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NounsToken *NounsTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NounsToken.Contract.SafeTransferFrom0(&_NounsToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NounsToken *NounsTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NounsToken.Contract.SafeTransferFrom0(&_NounsToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NounsToken *NounsTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NounsToken *NounsTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NounsToken.Contract.SetApprovalForAll(&_NounsToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NounsToken *NounsTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NounsToken.Contract.SetApprovalForAll(&_NounsToken.TransactOpts, operator, approved)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_NounsToken *NounsTokenTransactor) SetContractURIHash(opts *bind.TransactOpts, newContractURIHash string) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setContractURIHash", newContractURIHash)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_NounsToken *NounsTokenSession) SetContractURIHash(newContractURIHash string) (*types.Transaction, error) {
	return _NounsToken.Contract.SetContractURIHash(&_NounsToken.TransactOpts, newContractURIHash)
}

// SetContractURIHash is a paid mutator transaction binding the contract method 0xbaedc1c4.
//
// Solidity: function setContractURIHash(string newContractURIHash) returns()
func (_NounsToken *NounsTokenTransactorSession) SetContractURIHash(newContractURIHash string) (*types.Transaction, error) {
	return _NounsToken.Contract.SetContractURIHash(&_NounsToken.TransactOpts, newContractURIHash)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsToken *NounsTokenTransactor) SetDescriptor(opts *bind.TransactOpts, _descriptor common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setDescriptor", _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsToken *NounsTokenSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetDescriptor(&_NounsToken.TransactOpts, _descriptor)
}

// SetDescriptor is a paid mutator transaction binding the contract method 0x01b9a397.
//
// Solidity: function setDescriptor(address _descriptor) returns()
func (_NounsToken *NounsTokenTransactorSession) SetDescriptor(_descriptor common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetDescriptor(&_NounsToken.TransactOpts, _descriptor)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_NounsToken *NounsTokenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_NounsToken *NounsTokenSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetMinter(&_NounsToken.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_NounsToken *NounsTokenTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetMinter(&_NounsToken.TransactOpts, _minter)
}

// SetNoundersDAO is a paid mutator transaction binding the contract method 0x058df0ab.
//
// Solidity: function setNoundersDAO(address _noundersDAO) returns()
func (_NounsToken *NounsTokenTransactor) SetNoundersDAO(opts *bind.TransactOpts, _noundersDAO common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setNoundersDAO", _noundersDAO)
}

// SetNoundersDAO is a paid mutator transaction binding the contract method 0x058df0ab.
//
// Solidity: function setNoundersDAO(address _noundersDAO) returns()
func (_NounsToken *NounsTokenSession) SetNoundersDAO(_noundersDAO common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetNoundersDAO(&_NounsToken.TransactOpts, _noundersDAO)
}

// SetNoundersDAO is a paid mutator transaction binding the contract method 0x058df0ab.
//
// Solidity: function setNoundersDAO(address _noundersDAO) returns()
func (_NounsToken *NounsTokenTransactorSession) SetNoundersDAO(_noundersDAO common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetNoundersDAO(&_NounsToken.TransactOpts, _noundersDAO)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_NounsToken *NounsTokenTransactor) SetSeeder(opts *bind.TransactOpts, _seeder common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "setSeeder", _seeder)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_NounsToken *NounsTokenSession) SetSeeder(_seeder common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetSeeder(&_NounsToken.TransactOpts, _seeder)
}

// SetSeeder is a paid mutator transaction binding the contract method 0xd50b31eb.
//
// Solidity: function setSeeder(address _seeder) returns()
func (_NounsToken *NounsTokenTransactorSession) SetSeeder(_seeder common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.SetSeeder(&_NounsToken.TransactOpts, _seeder)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.TransferFrom(&_NounsToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NounsToken *NounsTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NounsToken.Contract.TransferFrom(&_NounsToken.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsToken *NounsTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NounsToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsToken *NounsTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.TransferOwnership(&_NounsToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsToken *NounsTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NounsToken.Contract.TransferOwnership(&_NounsToken.TransactOpts, newOwner)
}

// NounsTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NounsToken contract.
type NounsTokenApprovalIterator struct {
	Event *NounsTokenApproval // Event containing the contract specifics and raw log

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
func (it *NounsTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenApproval)
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
		it.Event = new(NounsTokenApproval)
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
func (it *NounsTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenApproval represents a Approval event raised by the NounsToken contract.
type NounsTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NounsTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenApprovalIterator{contract: _NounsToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NounsTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenApproval)
				if err := _NounsToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) ParseApproval(log types.Log) (*NounsTokenApproval, error) {
	event := new(NounsTokenApproval)
	if err := _NounsToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NounsToken contract.
type NounsTokenApprovalForAllIterator struct {
	Event *NounsTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NounsTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenApprovalForAll)
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
		it.Event = new(NounsTokenApprovalForAll)
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
func (it *NounsTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenApprovalForAll represents a ApprovalForAll event raised by the NounsToken contract.
type NounsTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NounsToken *NounsTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NounsTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenApprovalForAllIterator{contract: _NounsToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NounsToken *NounsTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NounsTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenApprovalForAll)
				if err := _NounsToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NounsToken *NounsTokenFilterer) ParseApprovalForAll(log types.Log) (*NounsTokenApprovalForAll, error) {
	event := new(NounsTokenApprovalForAll)
	if err := _NounsToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the NounsToken contract.
type NounsTokenDelegateChangedIterator struct {
	Event *NounsTokenDelegateChanged // Event containing the contract specifics and raw log

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
func (it *NounsTokenDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenDelegateChanged)
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
		it.Event = new(NounsTokenDelegateChanged)
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
func (it *NounsTokenDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenDelegateChanged represents a DelegateChanged event raised by the NounsToken contract.
type NounsTokenDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_NounsToken *NounsTokenFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*NounsTokenDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenDelegateChangedIterator{contract: _NounsToken.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_NounsToken *NounsTokenFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *NounsTokenDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenDelegateChanged)
				if err := _NounsToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_NounsToken *NounsTokenFilterer) ParseDelegateChanged(log types.Log) (*NounsTokenDelegateChanged, error) {
	event := new(NounsTokenDelegateChanged)
	if err := _NounsToken.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the NounsToken contract.
type NounsTokenDelegateVotesChangedIterator struct {
	Event *NounsTokenDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *NounsTokenDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenDelegateVotesChanged)
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
		it.Event = new(NounsTokenDelegateVotesChanged)
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
func (it *NounsTokenDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenDelegateVotesChanged represents a DelegateVotesChanged event raised by the NounsToken contract.
type NounsTokenDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_NounsToken *NounsTokenFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*NounsTokenDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenDelegateVotesChangedIterator{contract: _NounsToken.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_NounsToken *NounsTokenFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *NounsTokenDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenDelegateVotesChanged)
				if err := _NounsToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_NounsToken *NounsTokenFilterer) ParseDelegateVotesChanged(log types.Log) (*NounsTokenDelegateVotesChanged, error) {
	event := new(NounsTokenDelegateVotesChanged)
	if err := _NounsToken.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenDescriptorLockedIterator is returned from FilterDescriptorLocked and is used to iterate over the raw logs and unpacked data for DescriptorLocked events raised by the NounsToken contract.
type NounsTokenDescriptorLockedIterator struct {
	Event *NounsTokenDescriptorLocked // Event containing the contract specifics and raw log

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
func (it *NounsTokenDescriptorLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenDescriptorLocked)
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
		it.Event = new(NounsTokenDescriptorLocked)
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
func (it *NounsTokenDescriptorLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenDescriptorLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenDescriptorLocked represents a DescriptorLocked event raised by the NounsToken contract.
type NounsTokenDescriptorLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDescriptorLocked is a free log retrieval operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_NounsToken *NounsTokenFilterer) FilterDescriptorLocked(opts *bind.FilterOpts) (*NounsTokenDescriptorLockedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "DescriptorLocked")
	if err != nil {
		return nil, err
	}
	return &NounsTokenDescriptorLockedIterator{contract: _NounsToken.contract, event: "DescriptorLocked", logs: logs, sub: sub}, nil
}

// WatchDescriptorLocked is a free log subscription operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_NounsToken *NounsTokenFilterer) WatchDescriptorLocked(opts *bind.WatchOpts, sink chan<- *NounsTokenDescriptorLocked) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "DescriptorLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenDescriptorLocked)
				if err := _NounsToken.contract.UnpackLog(event, "DescriptorLocked", log); err != nil {
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

// ParseDescriptorLocked is a log parse operation binding the contract event 0x593e31e306c198bef259d839f7c6dc4ff7fc10c07f76fab193a210b03704105f.
//
// Solidity: event DescriptorLocked()
func (_NounsToken *NounsTokenFilterer) ParseDescriptorLocked(log types.Log) (*NounsTokenDescriptorLocked, error) {
	event := new(NounsTokenDescriptorLocked)
	if err := _NounsToken.contract.UnpackLog(event, "DescriptorLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenDescriptorUpdatedIterator is returned from FilterDescriptorUpdated and is used to iterate over the raw logs and unpacked data for DescriptorUpdated events raised by the NounsToken contract.
type NounsTokenDescriptorUpdatedIterator struct {
	Event *NounsTokenDescriptorUpdated // Event containing the contract specifics and raw log

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
func (it *NounsTokenDescriptorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenDescriptorUpdated)
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
		it.Event = new(NounsTokenDescriptorUpdated)
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
func (it *NounsTokenDescriptorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenDescriptorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenDescriptorUpdated represents a DescriptorUpdated event raised by the NounsToken contract.
type NounsTokenDescriptorUpdated struct {
	Descriptor common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDescriptorUpdated is a free log retrieval operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_NounsToken *NounsTokenFilterer) FilterDescriptorUpdated(opts *bind.FilterOpts) (*NounsTokenDescriptorUpdatedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsTokenDescriptorUpdatedIterator{contract: _NounsToken.contract, event: "DescriptorUpdated", logs: logs, sub: sub}, nil
}

// WatchDescriptorUpdated is a free log subscription operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_NounsToken *NounsTokenFilterer) WatchDescriptorUpdated(opts *bind.WatchOpts, sink chan<- *NounsTokenDescriptorUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "DescriptorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenDescriptorUpdated)
				if err := _NounsToken.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
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

// ParseDescriptorUpdated is a log parse operation binding the contract event 0x6e66ab22238a5471005895947c8f57db923c2a9c9c73180eff80864c21295c1b.
//
// Solidity: event DescriptorUpdated(address descriptor)
func (_NounsToken *NounsTokenFilterer) ParseDescriptorUpdated(log types.Log) (*NounsTokenDescriptorUpdated, error) {
	event := new(NounsTokenDescriptorUpdated)
	if err := _NounsToken.contract.UnpackLog(event, "DescriptorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenMinterLockedIterator is returned from FilterMinterLocked and is used to iterate over the raw logs and unpacked data for MinterLocked events raised by the NounsToken contract.
type NounsTokenMinterLockedIterator struct {
	Event *NounsTokenMinterLocked // Event containing the contract specifics and raw log

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
func (it *NounsTokenMinterLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenMinterLocked)
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
		it.Event = new(NounsTokenMinterLocked)
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
func (it *NounsTokenMinterLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenMinterLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenMinterLocked represents a MinterLocked event raised by the NounsToken contract.
type NounsTokenMinterLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMinterLocked is a free log retrieval operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_NounsToken *NounsTokenFilterer) FilterMinterLocked(opts *bind.FilterOpts) (*NounsTokenMinterLockedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "MinterLocked")
	if err != nil {
		return nil, err
	}
	return &NounsTokenMinterLockedIterator{contract: _NounsToken.contract, event: "MinterLocked", logs: logs, sub: sub}, nil
}

// WatchMinterLocked is a free log subscription operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_NounsToken *NounsTokenFilterer) WatchMinterLocked(opts *bind.WatchOpts, sink chan<- *NounsTokenMinterLocked) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "MinterLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenMinterLocked)
				if err := _NounsToken.contract.UnpackLog(event, "MinterLocked", log); err != nil {
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

// ParseMinterLocked is a log parse operation binding the contract event 0x192417b3f16b1ce69e0c59b0376549666650245ffc05e4b2569089dda8589b66.
//
// Solidity: event MinterLocked()
func (_NounsToken *NounsTokenFilterer) ParseMinterLocked(log types.Log) (*NounsTokenMinterLocked, error) {
	event := new(NounsTokenMinterLocked)
	if err := _NounsToken.contract.UnpackLog(event, "MinterLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenMinterUpdatedIterator is returned from FilterMinterUpdated and is used to iterate over the raw logs and unpacked data for MinterUpdated events raised by the NounsToken contract.
type NounsTokenMinterUpdatedIterator struct {
	Event *NounsTokenMinterUpdated // Event containing the contract specifics and raw log

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
func (it *NounsTokenMinterUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenMinterUpdated)
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
		it.Event = new(NounsTokenMinterUpdated)
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
func (it *NounsTokenMinterUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenMinterUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenMinterUpdated represents a MinterUpdated event raised by the NounsToken contract.
type NounsTokenMinterUpdated struct {
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinterUpdated is a free log retrieval operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_NounsToken *NounsTokenFilterer) FilterMinterUpdated(opts *bind.FilterOpts) (*NounsTokenMinterUpdatedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "MinterUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsTokenMinterUpdatedIterator{contract: _NounsToken.contract, event: "MinterUpdated", logs: logs, sub: sub}, nil
}

// WatchMinterUpdated is a free log subscription operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_NounsToken *NounsTokenFilterer) WatchMinterUpdated(opts *bind.WatchOpts, sink chan<- *NounsTokenMinterUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "MinterUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenMinterUpdated)
				if err := _NounsToken.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
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

// ParseMinterUpdated is a log parse operation binding the contract event 0xad0f299ec81a386c98df0ac27dae11dd020ed1b56963c53a7292e7a3a314539a.
//
// Solidity: event MinterUpdated(address minter)
func (_NounsToken *NounsTokenFilterer) ParseMinterUpdated(log types.Log) (*NounsTokenMinterUpdated, error) {
	event := new(NounsTokenMinterUpdated)
	if err := _NounsToken.contract.UnpackLog(event, "MinterUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenNounBurnedIterator is returned from FilterNounBurned and is used to iterate over the raw logs and unpacked data for NounBurned events raised by the NounsToken contract.
type NounsTokenNounBurnedIterator struct {
	Event *NounsTokenNounBurned // Event containing the contract specifics and raw log

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
func (it *NounsTokenNounBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenNounBurned)
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
		it.Event = new(NounsTokenNounBurned)
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
func (it *NounsTokenNounBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenNounBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenNounBurned represents a NounBurned event raised by the NounsToken contract.
type NounsTokenNounBurned struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNounBurned is a free log retrieval operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) FilterNounBurned(opts *bind.FilterOpts, tokenId []*big.Int) (*NounsTokenNounBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "NounBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenNounBurnedIterator{contract: _NounsToken.contract, event: "NounBurned", logs: logs, sub: sub}, nil
}

// WatchNounBurned is a free log subscription operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) WatchNounBurned(opts *bind.WatchOpts, sink chan<- *NounsTokenNounBurned, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "NounBurned", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenNounBurned)
				if err := _NounsToken.contract.UnpackLog(event, "NounBurned", log); err != nil {
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

// ParseNounBurned is a log parse operation binding the contract event 0x806be94a2ac8b92d74e99aa8add5a8e54528a01ec914a9e00d201a6480ed9863.
//
// Solidity: event NounBurned(uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) ParseNounBurned(log types.Log) (*NounsTokenNounBurned, error) {
	event := new(NounsTokenNounBurned)
	if err := _NounsToken.contract.UnpackLog(event, "NounBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenNounCreatedIterator is returned from FilterNounCreated and is used to iterate over the raw logs and unpacked data for NounCreated events raised by the NounsToken contract.
type NounsTokenNounCreatedIterator struct {
	Event *NounsTokenNounCreated // Event containing the contract specifics and raw log

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
func (it *NounsTokenNounCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenNounCreated)
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
		it.Event = new(NounsTokenNounCreated)
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
func (it *NounsTokenNounCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenNounCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenNounCreated represents a NounCreated event raised by the NounsToken contract.
type NounsTokenNounCreated struct {
	TokenId *big.Int
	Seed    INounsSeederSeed
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNounCreated is a free log retrieval operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_NounsToken *NounsTokenFilterer) FilterNounCreated(opts *bind.FilterOpts, tokenId []*big.Int) (*NounsTokenNounCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "NounCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenNounCreatedIterator{contract: _NounsToken.contract, event: "NounCreated", logs: logs, sub: sub}, nil
}

// WatchNounCreated is a free log subscription operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_NounsToken *NounsTokenFilterer) WatchNounCreated(opts *bind.WatchOpts, sink chan<- *NounsTokenNounCreated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "NounCreated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenNounCreated)
				if err := _NounsToken.contract.UnpackLog(event, "NounCreated", log); err != nil {
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

// ParseNounCreated is a log parse operation binding the contract event 0x1106ee9d020bfbb5ee34cf5535a5fbf024a011bd130078088cbf124ab3092478.
//
// Solidity: event NounCreated(uint256 indexed tokenId, (uint48,uint48,uint48,uint48,uint48) seed)
func (_NounsToken *NounsTokenFilterer) ParseNounCreated(log types.Log) (*NounsTokenNounCreated, error) {
	event := new(NounsTokenNounCreated)
	if err := _NounsToken.contract.UnpackLog(event, "NounCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenNoundersDAOUpdatedIterator is returned from FilterNoundersDAOUpdated and is used to iterate over the raw logs and unpacked data for NoundersDAOUpdated events raised by the NounsToken contract.
type NounsTokenNoundersDAOUpdatedIterator struct {
	Event *NounsTokenNoundersDAOUpdated // Event containing the contract specifics and raw log

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
func (it *NounsTokenNoundersDAOUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenNoundersDAOUpdated)
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
		it.Event = new(NounsTokenNoundersDAOUpdated)
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
func (it *NounsTokenNoundersDAOUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenNoundersDAOUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenNoundersDAOUpdated represents a NoundersDAOUpdated event raised by the NounsToken contract.
type NounsTokenNoundersDAOUpdated struct {
	NoundersDAO common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNoundersDAOUpdated is a free log retrieval operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_NounsToken *NounsTokenFilterer) FilterNoundersDAOUpdated(opts *bind.FilterOpts) (*NounsTokenNoundersDAOUpdatedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "NoundersDAOUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsTokenNoundersDAOUpdatedIterator{contract: _NounsToken.contract, event: "NoundersDAOUpdated", logs: logs, sub: sub}, nil
}

// WatchNoundersDAOUpdated is a free log subscription operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_NounsToken *NounsTokenFilterer) WatchNoundersDAOUpdated(opts *bind.WatchOpts, sink chan<- *NounsTokenNoundersDAOUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "NoundersDAOUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenNoundersDAOUpdated)
				if err := _NounsToken.contract.UnpackLog(event, "NoundersDAOUpdated", log); err != nil {
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

// ParseNoundersDAOUpdated is a log parse operation binding the contract event 0x3a0b923617f180781f3530e464cb4a8b9393e69f47607e4eb28d61cd87ce968c.
//
// Solidity: event NoundersDAOUpdated(address noundersDAO)
func (_NounsToken *NounsTokenFilterer) ParseNoundersDAOUpdated(log types.Log) (*NounsTokenNoundersDAOUpdated, error) {
	event := new(NounsTokenNoundersDAOUpdated)
	if err := _NounsToken.contract.UnpackLog(event, "NoundersDAOUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NounsToken contract.
type NounsTokenOwnershipTransferredIterator struct {
	Event *NounsTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NounsTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenOwnershipTransferred)
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
		it.Event = new(NounsTokenOwnershipTransferred)
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
func (it *NounsTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenOwnershipTransferred represents a OwnershipTransferred event raised by the NounsToken contract.
type NounsTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NounsToken *NounsTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NounsTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenOwnershipTransferredIterator{contract: _NounsToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NounsToken *NounsTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NounsTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenOwnershipTransferred)
				if err := _NounsToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NounsToken *NounsTokenFilterer) ParseOwnershipTransferred(log types.Log) (*NounsTokenOwnershipTransferred, error) {
	event := new(NounsTokenOwnershipTransferred)
	if err := _NounsToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenSeederLockedIterator is returned from FilterSeederLocked and is used to iterate over the raw logs and unpacked data for SeederLocked events raised by the NounsToken contract.
type NounsTokenSeederLockedIterator struct {
	Event *NounsTokenSeederLocked // Event containing the contract specifics and raw log

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
func (it *NounsTokenSeederLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenSeederLocked)
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
		it.Event = new(NounsTokenSeederLocked)
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
func (it *NounsTokenSeederLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenSeederLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenSeederLocked represents a SeederLocked event raised by the NounsToken contract.
type NounsTokenSeederLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterSeederLocked is a free log retrieval operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_NounsToken *NounsTokenFilterer) FilterSeederLocked(opts *bind.FilterOpts) (*NounsTokenSeederLockedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "SeederLocked")
	if err != nil {
		return nil, err
	}
	return &NounsTokenSeederLockedIterator{contract: _NounsToken.contract, event: "SeederLocked", logs: logs, sub: sub}, nil
}

// WatchSeederLocked is a free log subscription operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_NounsToken *NounsTokenFilterer) WatchSeederLocked(opts *bind.WatchOpts, sink chan<- *NounsTokenSeederLocked) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "SeederLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenSeederLocked)
				if err := _NounsToken.contract.UnpackLog(event, "SeederLocked", log); err != nil {
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

// ParseSeederLocked is a log parse operation binding the contract event 0xf59561f22794afcfb1e6be5c4733f5449fd167252a96b74bb06d341fb0dac7ed.
//
// Solidity: event SeederLocked()
func (_NounsToken *NounsTokenFilterer) ParseSeederLocked(log types.Log) (*NounsTokenSeederLocked, error) {
	event := new(NounsTokenSeederLocked)
	if err := _NounsToken.contract.UnpackLog(event, "SeederLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenSeederUpdatedIterator is returned from FilterSeederUpdated and is used to iterate over the raw logs and unpacked data for SeederUpdated events raised by the NounsToken contract.
type NounsTokenSeederUpdatedIterator struct {
	Event *NounsTokenSeederUpdated // Event containing the contract specifics and raw log

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
func (it *NounsTokenSeederUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenSeederUpdated)
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
		it.Event = new(NounsTokenSeederUpdated)
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
func (it *NounsTokenSeederUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenSeederUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenSeederUpdated represents a SeederUpdated event raised by the NounsToken contract.
type NounsTokenSeederUpdated struct {
	Seeder common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSeederUpdated is a free log retrieval operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_NounsToken *NounsTokenFilterer) FilterSeederUpdated(opts *bind.FilterOpts) (*NounsTokenSeederUpdatedIterator, error) {

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "SeederUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsTokenSeederUpdatedIterator{contract: _NounsToken.contract, event: "SeederUpdated", logs: logs, sub: sub}, nil
}

// WatchSeederUpdated is a free log subscription operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_NounsToken *NounsTokenFilterer) WatchSeederUpdated(opts *bind.WatchOpts, sink chan<- *NounsTokenSeederUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "SeederUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenSeederUpdated)
				if err := _NounsToken.contract.UnpackLog(event, "SeederUpdated", log); err != nil {
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

// ParseSeederUpdated is a log parse operation binding the contract event 0xb3025222d01ce9a26c7f9d52bc3bfd0352366bd90a793c273fbfe1c81e0e288e.
//
// Solidity: event SeederUpdated(address seeder)
func (_NounsToken *NounsTokenFilterer) ParseSeederUpdated(log types.Log) (*NounsTokenSeederUpdated, error) {
	event := new(NounsTokenSeederUpdated)
	if err := _NounsToken.contract.UnpackLog(event, "SeederUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NounsToken contract.
type NounsTokenTransferIterator struct {
	Event *NounsTokenTransfer // Event containing the contract specifics and raw log

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
func (it *NounsTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsTokenTransfer)
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
		it.Event = new(NounsTokenTransfer)
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
func (it *NounsTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsTokenTransfer represents a Transfer event raised by the NounsToken contract.
type NounsTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NounsTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsTokenTransferIterator{contract: _NounsToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NounsTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NounsToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsTokenTransfer)
				if err := _NounsToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NounsToken *NounsTokenFilterer) ParseTransfer(log types.Log) (*NounsTokenTransfer, error) {
	event := new(NounsTokenTransfer)
	if err := _NounsToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
