// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package periphery

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

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMigrateData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMigrateData struct {
	Account     common.Address
	Handle      string
	Uri         string
	ToAddresses []common.Address
	LinkType    [32]byte
}

// DataTypesNote is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNote struct {
	LinkItemType [32]byte
	LinkKey      [32]byte
	ContentUri   string
	LinkModule   common.Address
	MintModule   common.Address
	MintNFT      common.Address
	Deleted      bool
	Locked       bool
}

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	CharacterId *big.Int
	NoteId      *big.Int
}

// DataTypeslinkCharactersInBatchData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkCharactersInBatchData struct {
	FromCharacterId *big.Int
	ToCharacterIds  []*big.Int
	Data            [][]byte
	ToAddresses     []common.Address
	LinkType        [32]byte
}

// PeripheryMetaData contains all meta data concerning the Periphery contract.
var PeripheryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"version\",\"internalType\":\"uint8\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getLinkingAddress\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"address[]\"}],\"name\":\"getLinkingAddresses\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getLinkingAnyUri\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string[]\",\"name\":\"results\",\"internalType\":\"string[]\"}],\"name\":\"getLinkingAnyUris\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"characterId\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingCharacterId\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"results\",\"internalType\":\"uint256[]\"}],\"name\":\"getLinkingCharacterIds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.ERC721Struct\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingERC721\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.ERC721Struct[]\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingERC721s\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingLinklistId\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"linklistIds\",\"internalType\":\"uint256[]\"}],\"name\":\"getLinkingLinklistIds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.NoteStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"characterId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingNote\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.Note[]\",\"components\":[{\"type\":\"bytes32\",\"name\":\"linkItemType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintNFT\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"deleted\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getLinkingNotes\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.Note[]\",\"components\":[{\"type\":\"bytes32\",\"name\":\"linkItemType\",\"internalType\":\"bytes32\"},{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"},{\"type\":\"string\",\"name\":\"contentUri\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"linkModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintModule\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"mintNFT\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"deleted\",\"internalType\":\"bool\"},{\"type\":\"bool\",\"name\":\"locked\",\"internalType\":\"bool\"}]}],\"name\":\"getNotesByCharacterId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"characterId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"offset\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"limit\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"address\",\"name\":\"_web3Entry\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"_linklist\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"linkCharactersInBatch\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.linkCharactersInBatchData\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromCharacterId\",\"internalType\":\"uint256\"},{\"type\":\"uint256[]\",\"name\":\"toCharacterIds\",\"internalType\":\"uint256[]\"},{\"type\":\"bytes[]\",\"name\":\"data\",\"internalType\":\"bytes[]\"},{\"type\":\"address[]\",\"name\":\"toAddresses\",\"internalType\":\"address[]\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"linklist\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"migrate\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"vars\",\"internalType\":\"structDataTypes.MigrateData\",\"components\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"toAddresses\",\"internalType\":\"address[]\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"sync\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\"},{\"type\":\"string\",\"name\":\"handle\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"uri\",\"internalType\":\"string\"},{\"type\":\"address[]\",\"name\":\"toAddresses\",\"internalType\":\"address[]\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"web3Entry\",\"inputs\":[]}]",
}

// PeripheryABI is the input ABI used to generate the binding from.
// Deprecated: Use PeripheryMetaData.ABI instead.
var PeripheryABI = PeripheryMetaData.ABI

// Periphery is an auto generated Go binding around an Ethereum contract.
type Periphery struct {
	PeripheryCaller     // Read-only binding to the contract
	PeripheryTransactor // Write-only binding to the contract
	PeripheryFilterer   // Log filterer for contract events
}

// PeripheryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PeripheryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PeripheryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripheryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PeripheryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PeripherySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PeripherySession struct {
	Contract     *Periphery        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PeripheryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PeripheryCallerSession struct {
	Contract *PeripheryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PeripheryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PeripheryTransactorSession struct {
	Contract     *PeripheryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PeripheryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PeripheryRaw struct {
	Contract *Periphery // Generic contract binding to access the raw methods on
}

// PeripheryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PeripheryCallerRaw struct {
	Contract *PeripheryCaller // Generic read-only contract binding to access the raw methods on
}

// PeripheryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PeripheryTransactorRaw struct {
	Contract *PeripheryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPeriphery creates a new instance of Periphery, bound to a specific deployed contract.
func NewPeriphery(address common.Address, backend bind.ContractBackend) (*Periphery, error) {
	contract, err := bindPeriphery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Periphery{PeripheryCaller: PeripheryCaller{contract: contract}, PeripheryTransactor: PeripheryTransactor{contract: contract}, PeripheryFilterer: PeripheryFilterer{contract: contract}}, nil
}

// NewPeripheryCaller creates a new read-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryCaller(address common.Address, caller bind.ContractCaller) (*PeripheryCaller, error) {
	contract, err := bindPeriphery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryCaller{contract: contract}, nil
}

// NewPeripheryTransactor creates a new write-only instance of Periphery, bound to a specific deployed contract.
func NewPeripheryTransactor(address common.Address, transactor bind.ContractTransactor) (*PeripheryTransactor, error) {
	contract, err := bindPeriphery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PeripheryTransactor{contract: contract}, nil
}

// NewPeripheryFilterer creates a new log filterer instance of Periphery, bound to a specific deployed contract.
func NewPeripheryFilterer(address common.Address, filterer bind.ContractFilterer) (*PeripheryFilterer, error) {
	contract, err := bindPeriphery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PeripheryFilterer{contract: contract}, nil
}

// bindPeriphery binds a generic wrapper to an already deployed contract.
func bindPeriphery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PeripheryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.PeripheryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.PeripheryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Periphery *PeripheryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Periphery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Periphery *PeripheryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Periphery *PeripheryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Periphery.Contract.contract.Transact(opts, method, params...)
}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) pure returns(address)
func (_Periphery *PeripheryCaller) GetLinkingAddress(opts *bind.CallOpts, linkKey [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAddress", linkKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) pure returns(address)
func (_Periphery *PeripherySession) GetLinkingAddress(linkKey [32]byte) (common.Address, error) {
	return _Periphery.Contract.GetLinkingAddress(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAddress is a free data retrieval call binding the contract method 0x3e7cb539.
//
// Solidity: function getLinkingAddress(bytes32 linkKey) pure returns(address)
func (_Periphery *PeripheryCallerSession) GetLinkingAddress(linkKey [32]byte) (common.Address, error) {
	return _Periphery.Contract.GetLinkingAddress(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripheryCaller) GetLinkingAddresses(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAddresses", fromCharacterId, linkType)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripherySession) GetLinkingAddresses(fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	return _Periphery.Contract.GetLinkingAddresses(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x76ae90ef.
//
// Solidity: function getLinkingAddresses(uint256 fromCharacterId, bytes32 linkType) view returns(address[])
func (_Periphery *PeripheryCallerSession) GetLinkingAddresses(fromCharacterId *big.Int, linkType [32]byte) ([]common.Address, error) {
	return _Periphery.Contract.GetLinkingAddresses(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripheryCaller) GetLinkingAnyUri(opts *bind.CallOpts, linkKey [32]byte) (string, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAnyUri", linkKey)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripherySession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _Periphery.Contract.GetLinkingAnyUri(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_Periphery *PeripheryCallerSession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _Periphery.Contract.GetLinkingAnyUri(&_Periphery.CallOpts, linkKey)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripheryCaller) GetLinkingAnyUris(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingAnyUris", fromCharacterId, linkType)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripherySession) GetLinkingAnyUris(fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	return _Periphery.Contract.GetLinkingAnyUris(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0x1964a421.
//
// Solidity: function getLinkingAnyUris(uint256 fromCharacterId, bytes32 linkType) view returns(string[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingAnyUris(fromCharacterId *big.Int, linkType [32]byte) ([]string, error) {
	return _Periphery.Contract.GetLinkingAnyUris(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) pure returns(uint256 characterId)
func (_Periphery *PeripheryCaller) GetLinkingCharacterId(opts *bind.CallOpts, linkKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingCharacterId", linkKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) pure returns(uint256 characterId)
func (_Periphery *PeripherySession) GetLinkingCharacterId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingCharacterId is a free data retrieval call binding the contract method 0xd85d04f4.
//
// Solidity: function getLinkingCharacterId(bytes32 linkKey) pure returns(uint256 characterId)
func (_Periphery *PeripheryCallerSession) GetLinkingCharacterId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripheryCaller) GetLinkingCharacterIds(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingCharacterIds", fromCharacterId, linkType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripherySession) GetLinkingCharacterIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingCharacterIds is a free data retrieval call binding the contract method 0xdfc1e2bf.
//
// Solidity: function getLinkingCharacterIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingCharacterIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingCharacterIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripheryCaller) GetLinkingERC721(opts *bind.CallOpts, linkKey [32]byte) (DataTypesERC721Struct, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingERC721", linkKey)

	if err != nil {
		return *new(DataTypesERC721Struct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesERC721Struct)).(*DataTypesERC721Struct)

	return out0, err

}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripherySession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721(&_Periphery.CallOpts, linkKey)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_Periphery *PeripheryCallerSession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721(&_Periphery.CallOpts, linkKey)
}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x62c3ed30.
//
// Solidity: function getLinkingERC721s(uint256 fromCharacterId, bytes32 linkType) view returns((address,uint256)[] results)
func (_Periphery *PeripheryCaller) GetLinkingERC721s(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesERC721Struct, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingERC721s", fromCharacterId, linkType)

	if err != nil {
		return *new([]DataTypesERC721Struct), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesERC721Struct)).(*[]DataTypesERC721Struct)

	return out0, err

}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x62c3ed30.
//
// Solidity: function getLinkingERC721s(uint256 fromCharacterId, bytes32 linkType) view returns((address,uint256)[] results)
func (_Periphery *PeripherySession) GetLinkingERC721s(fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721s(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x62c3ed30.
//
// Solidity: function getLinkingERC721s(uint256 fromCharacterId, bytes32 linkType) view returns((address,uint256)[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingERC721s(fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesERC721Struct, error) {
	return _Periphery.Contract.GetLinkingERC721s(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingLinklistId is a free data retrieval call binding the contract method 0x8654af27.
//
// Solidity: function getLinkingLinklistId(bytes32 linkKey) pure returns(uint256 linklistId)
func (_Periphery *PeripheryCaller) GetLinkingLinklistId(opts *bind.CallOpts, linkKey [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingLinklistId", linkKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingLinklistId is a free data retrieval call binding the contract method 0x8654af27.
//
// Solidity: function getLinkingLinklistId(bytes32 linkKey) pure returns(uint256 linklistId)
func (_Periphery *PeripherySession) GetLinkingLinklistId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingLinklistId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingLinklistId is a free data retrieval call binding the contract method 0x8654af27.
//
// Solidity: function getLinkingLinklistId(bytes32 linkKey) pure returns(uint256 linklistId)
func (_Periphery *PeripheryCallerSession) GetLinkingLinklistId(linkKey [32]byte) (*big.Int, error) {
	return _Periphery.Contract.GetLinkingLinklistId(&_Periphery.CallOpts, linkKey)
}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0x2260c7d2.
//
// Solidity: function getLinkingLinklistIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] linklistIds)
func (_Periphery *PeripheryCaller) GetLinkingLinklistIds(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingLinklistIds", fromCharacterId, linkType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0x2260c7d2.
//
// Solidity: function getLinkingLinklistIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] linklistIds)
func (_Periphery *PeripherySession) GetLinkingLinklistIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingLinklistIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0x2260c7d2.
//
// Solidity: function getLinkingLinklistIds(uint256 fromCharacterId, bytes32 linkType) view returns(uint256[] linklistIds)
func (_Periphery *PeripheryCallerSession) GetLinkingLinklistIds(fromCharacterId *big.Int, linkType [32]byte) ([]*big.Int, error) {
	return _Periphery.Contract.GetLinkingLinklistIds(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_Periphery *PeripheryCaller) GetLinkingNote(opts *bind.CallOpts, linkKey [32]byte) (DataTypesNoteStruct, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingNote", linkKey)

	if err != nil {
		return *new(DataTypesNoteStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNoteStruct)).(*DataTypesNoteStruct)

	return out0, err

}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_Periphery *PeripherySession) GetLinkingNote(linkKey [32]byte) (DataTypesNoteStruct, error) {
	return _Periphery.Contract.GetLinkingNote(&_Periphery.CallOpts, linkKey)
}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_Periphery *PeripheryCallerSession) GetLinkingNote(linkKey [32]byte) (DataTypesNoteStruct, error) {
	return _Periphery.Contract.GetLinkingNote(&_Periphery.CallOpts, linkKey)
}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x2c422599.
//
// Solidity: function getLinkingNotes(uint256 fromCharacterId, bytes32 linkType) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripheryCaller) GetLinkingNotes(opts *bind.CallOpts, fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesNote, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getLinkingNotes", fromCharacterId, linkType)

	if err != nil {
		return *new([]DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesNote)).(*[]DataTypesNote)

	return out0, err

}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x2c422599.
//
// Solidity: function getLinkingNotes(uint256 fromCharacterId, bytes32 linkType) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripherySession) GetLinkingNotes(fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesNote, error) {
	return _Periphery.Contract.GetLinkingNotes(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x2c422599.
//
// Solidity: function getLinkingNotes(uint256 fromCharacterId, bytes32 linkType) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripheryCallerSession) GetLinkingNotes(fromCharacterId *big.Int, linkType [32]byte) ([]DataTypesNote, error) {
	return _Periphery.Contract.GetLinkingNotes(&_Periphery.CallOpts, fromCharacterId, linkType)
}

// GetNotesByCharacterId is a free data retrieval call binding the contract method 0x0f9b729c.
//
// Solidity: function getNotesByCharacterId(uint256 characterId, uint256 offset, uint256 limit) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripheryCaller) GetNotesByCharacterId(opts *bind.CallOpts, characterId *big.Int, offset *big.Int, limit *big.Int) ([]DataTypesNote, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "getNotesByCharacterId", characterId, offset, limit)

	if err != nil {
		return *new([]DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesNote)).(*[]DataTypesNote)

	return out0, err

}

// GetNotesByCharacterId is a free data retrieval call binding the contract method 0x0f9b729c.
//
// Solidity: function getNotesByCharacterId(uint256 characterId, uint256 offset, uint256 limit) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripherySession) GetNotesByCharacterId(characterId *big.Int, offset *big.Int, limit *big.Int) ([]DataTypesNote, error) {
	return _Periphery.Contract.GetNotesByCharacterId(&_Periphery.CallOpts, characterId, offset, limit)
}

// GetNotesByCharacterId is a free data retrieval call binding the contract method 0x0f9b729c.
//
// Solidity: function getNotesByCharacterId(uint256 characterId, uint256 offset, uint256 limit) view returns((bytes32,bytes32,string,address,address,address,bool,bool)[] results)
func (_Periphery *PeripheryCallerSession) GetNotesByCharacterId(characterId *big.Int, offset *big.Int, limit *big.Int) ([]DataTypesNote, error) {
	return _Periphery.Contract.GetNotesByCharacterId(&_Periphery.CallOpts, characterId, offset, limit)
}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripheryCaller) Linklist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "linklist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripherySession) Linklist() (common.Address, error) {
	return _Periphery.Contract.Linklist(&_Periphery.CallOpts)
}

// Linklist is a free data retrieval call binding the contract method 0x06bee8d4.
//
// Solidity: function linklist() view returns(address)
func (_Periphery *PeripheryCallerSession) Linklist() (common.Address, error) {
	return _Periphery.Contract.Linklist(&_Periphery.CallOpts)
}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripheryCaller) Web3Entry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Periphery.contract.Call(opts, &out, "web3Entry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripherySession) Web3Entry() (common.Address, error) {
	return _Periphery.Contract.Web3Entry(&_Periphery.CallOpts)
}

// Web3Entry is a free data retrieval call binding the contract method 0xe1332272.
//
// Solidity: function web3Entry() view returns(address)
func (_Periphery *PeripheryCallerSession) Web3Entry() (common.Address, error) {
	return _Periphery.Contract.Web3Entry(&_Periphery.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _web3Entry, address _linklist) returns()
func (_Periphery *PeripheryTransactor) Initialize(opts *bind.TransactOpts, _web3Entry common.Address, _linklist common.Address) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "initialize", _web3Entry, _linklist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _web3Entry, address _linklist) returns()
func (_Periphery *PeripherySession) Initialize(_web3Entry common.Address, _linklist common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.Initialize(&_Periphery.TransactOpts, _web3Entry, _linklist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _web3Entry, address _linklist) returns()
func (_Periphery *PeripheryTransactorSession) Initialize(_web3Entry common.Address, _linklist common.Address) (*types.Transaction, error) {
	return _Periphery.Contract.Initialize(&_Periphery.TransactOpts, _web3Entry, _linklist)
}

// LinkCharactersInBatch is a paid mutator transaction binding the contract method 0xd55a0060.
//
// Solidity: function linkCharactersInBatch((uint256,uint256[],bytes[],address[],bytes32) vars) returns()
func (_Periphery *PeripheryTransactor) LinkCharactersInBatch(opts *bind.TransactOpts, vars DataTypeslinkCharactersInBatchData) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "linkCharactersInBatch", vars)
}

// LinkCharactersInBatch is a paid mutator transaction binding the contract method 0xd55a0060.
//
// Solidity: function linkCharactersInBatch((uint256,uint256[],bytes[],address[],bytes32) vars) returns()
func (_Periphery *PeripherySession) LinkCharactersInBatch(vars DataTypeslinkCharactersInBatchData) (*types.Transaction, error) {
	return _Periphery.Contract.LinkCharactersInBatch(&_Periphery.TransactOpts, vars)
}

// LinkCharactersInBatch is a paid mutator transaction binding the contract method 0xd55a0060.
//
// Solidity: function linkCharactersInBatch((uint256,uint256[],bytes[],address[],bytes32) vars) returns()
func (_Periphery *PeripheryTransactorSession) LinkCharactersInBatch(vars DataTypeslinkCharactersInBatchData) (*types.Transaction, error) {
	return _Periphery.Contract.LinkCharactersInBatch(&_Periphery.TransactOpts, vars)
}

// Migrate is a paid mutator transaction binding the contract method 0xe3d24428.
//
// Solidity: function migrate((address,string,string,address[],bytes32) vars) returns()
func (_Periphery *PeripheryTransactor) Migrate(opts *bind.TransactOpts, vars DataTypesMigrateData) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "migrate", vars)
}

// Migrate is a paid mutator transaction binding the contract method 0xe3d24428.
//
// Solidity: function migrate((address,string,string,address[],bytes32) vars) returns()
func (_Periphery *PeripherySession) Migrate(vars DataTypesMigrateData) (*types.Transaction, error) {
	return _Periphery.Contract.Migrate(&_Periphery.TransactOpts, vars)
}

// Migrate is a paid mutator transaction binding the contract method 0xe3d24428.
//
// Solidity: function migrate((address,string,string,address[],bytes32) vars) returns()
func (_Periphery *PeripheryTransactorSession) Migrate(vars DataTypesMigrateData) (*types.Transaction, error) {
	return _Periphery.Contract.Migrate(&_Periphery.TransactOpts, vars)
}

// Sync is a paid mutator transaction binding the contract method 0x5fb8652e.
//
// Solidity: function sync(address account, string handle, string uri, address[] toAddresses, bytes32 linkType) returns()
func (_Periphery *PeripheryTransactor) Sync(opts *bind.TransactOpts, account common.Address, handle string, uri string, toAddresses []common.Address, linkType [32]byte) (*types.Transaction, error) {
	return _Periphery.contract.Transact(opts, "sync", account, handle, uri, toAddresses, linkType)
}

// Sync is a paid mutator transaction binding the contract method 0x5fb8652e.
//
// Solidity: function sync(address account, string handle, string uri, address[] toAddresses, bytes32 linkType) returns()
func (_Periphery *PeripherySession) Sync(account common.Address, handle string, uri string, toAddresses []common.Address, linkType [32]byte) (*types.Transaction, error) {
	return _Periphery.Contract.Sync(&_Periphery.TransactOpts, account, handle, uri, toAddresses, linkType)
}

// Sync is a paid mutator transaction binding the contract method 0x5fb8652e.
//
// Solidity: function sync(address account, string handle, string uri, address[] toAddresses, bytes32 linkType) returns()
func (_Periphery *PeripheryTransactorSession) Sync(account common.Address, handle string, uri string, toAddresses []common.Address, linkType [32]byte) (*types.Transaction, error) {
	return _Periphery.Contract.Sync(&_Periphery.TransactOpts, account, handle, uri, toAddresses, linkType)
}

// PeripheryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Periphery contract.
type PeripheryInitializedIterator struct {
	Event *PeripheryInitialized // Event containing the contract specifics and raw log

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
func (it *PeripheryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PeripheryInitialized)
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
		it.Event = new(PeripheryInitialized)
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
func (it *PeripheryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PeripheryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PeripheryInitialized represents a Initialized event raised by the Periphery contract.
type PeripheryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Periphery *PeripheryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PeripheryInitializedIterator, error) {

	logs, sub, err := _Periphery.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PeripheryInitializedIterator{contract: _Periphery.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Periphery *PeripheryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PeripheryInitialized) (event.Subscription, error) {

	logs, sub, err := _Periphery.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PeripheryInitialized)
				if err := _Periphery.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Periphery *PeripheryFilterer) ParseInitialized(log types.Log) (*PeripheryInitialized, error) {
	event := new(PeripheryInitialized)
	if err := _Periphery.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
