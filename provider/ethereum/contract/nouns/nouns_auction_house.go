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

// NounsAuctionHouseMetaData contains all meta data concerning the NounsAuctionHouse contract.
var NounsAuctionHouseMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"extended\",\"type\":\"bool\"}],\"name\":\"AuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"AuctionExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minBidIncrementPercentage\",\"type\":\"uint256\"}],\"name\":\"AuctionMinBidIncrementPercentageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reservePrice\",\"type\":\"uint256\"}],\"name\":\"AuctionReservePriceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AuctionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"}],\"name\":\"AuctionTimeBufferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nounId\",\"type\":\"uint256\"}],\"name\":\"createBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractINounsToken\",\"name\":\"_nouns\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minBidIncrementPercentage\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nouns\",\"outputs\":[{\"internalType\":\"contractINounsToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_minBidIncrementPercentage\",\"type\":\"uint8\"}],\"name\":\"setMinBidIncrementPercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_reservePrice\",\"type\":\"uint256\"}],\"name\":\"setReservePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeBuffer\",\"type\":\"uint256\"}],\"name\":\"setTimeBuffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settleCurrentAndCreateNewAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeBuffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NounsAuctionHouseABI is the input ABI used to generate the binding from.
// Deprecated: Use NounsAuctionHouseMetaData.ABI instead.
var NounsAuctionHouseABI = NounsAuctionHouseMetaData.ABI

// NounsAuctionHouse is an auto generated Go binding around an Ethereum contract.
type NounsAuctionHouse struct {
	NounsAuctionHouseCaller     // Read-only binding to the contract
	NounsAuctionHouseTransactor // Write-only binding to the contract
	NounsAuctionHouseFilterer   // Log filterer for contract events
}

// NounsAuctionHouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type NounsAuctionHouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsAuctionHouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NounsAuctionHouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsAuctionHouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NounsAuctionHouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsAuctionHouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NounsAuctionHouseSession struct {
	Contract     *NounsAuctionHouse // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NounsAuctionHouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NounsAuctionHouseCallerSession struct {
	Contract *NounsAuctionHouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// NounsAuctionHouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NounsAuctionHouseTransactorSession struct {
	Contract     *NounsAuctionHouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// NounsAuctionHouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type NounsAuctionHouseRaw struct {
	Contract *NounsAuctionHouse // Generic contract binding to access the raw methods on
}

// NounsAuctionHouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NounsAuctionHouseCallerRaw struct {
	Contract *NounsAuctionHouseCaller // Generic read-only contract binding to access the raw methods on
}

// NounsAuctionHouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NounsAuctionHouseTransactorRaw struct {
	Contract *NounsAuctionHouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNounsAuctionHouse creates a new instance of NounsAuctionHouse, bound to a specific deployed contract.
func NewNounsAuctionHouse(address common.Address, backend bind.ContractBackend) (*NounsAuctionHouse, error) {
	contract, err := bindNounsAuctionHouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouse{NounsAuctionHouseCaller: NounsAuctionHouseCaller{contract: contract}, NounsAuctionHouseTransactor: NounsAuctionHouseTransactor{contract: contract}, NounsAuctionHouseFilterer: NounsAuctionHouseFilterer{contract: contract}}, nil
}

// NewNounsAuctionHouseCaller creates a new read-only instance of NounsAuctionHouse, bound to a specific deployed contract.
func NewNounsAuctionHouseCaller(address common.Address, caller bind.ContractCaller) (*NounsAuctionHouseCaller, error) {
	contract, err := bindNounsAuctionHouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseCaller{contract: contract}, nil
}

// NewNounsAuctionHouseTransactor creates a new write-only instance of NounsAuctionHouse, bound to a specific deployed contract.
func NewNounsAuctionHouseTransactor(address common.Address, transactor bind.ContractTransactor) (*NounsAuctionHouseTransactor, error) {
	contract, err := bindNounsAuctionHouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseTransactor{contract: contract}, nil
}

// NewNounsAuctionHouseFilterer creates a new log filterer instance of NounsAuctionHouse, bound to a specific deployed contract.
func NewNounsAuctionHouseFilterer(address common.Address, filterer bind.ContractFilterer) (*NounsAuctionHouseFilterer, error) {
	contract, err := bindNounsAuctionHouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseFilterer{contract: contract}, nil
}

// bindNounsAuctionHouse binds a generic wrapper to an already deployed contract.
func bindNounsAuctionHouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NounsAuctionHouseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsAuctionHouse *NounsAuctionHouseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsAuctionHouse.Contract.NounsAuctionHouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsAuctionHouse *NounsAuctionHouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.NounsAuctionHouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsAuctionHouse *NounsAuctionHouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.NounsAuctionHouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsAuctionHouse *NounsAuctionHouseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsAuctionHouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsAuctionHouse *NounsAuctionHouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsAuctionHouse *NounsAuctionHouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.contract.Transact(opts, method, params...)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Auction(opts *bind.CallOpts) (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "auction")

	outstruct := new(struct {
		NounId    *big.Int
		Amount    *big.Int
		StartTime *big.Int
		EndTime   *big.Int
		Bidder    common.Address
		Settled   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NounId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Settled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _NounsAuctionHouse.Contract.Auction(&_NounsAuctionHouse.CallOpts)
}

// Auction is a free data retrieval call binding the contract method 0x7d9f6db5.
//
// Solidity: function auction() view returns(uint256 nounId, uint256 amount, uint256 startTime, uint256 endTime, address bidder, bool settled)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Auction() (struct {
	NounId    *big.Int
	Amount    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Bidder    common.Address
	Settled   bool
}, error) {
	return _NounsAuctionHouse.Contract.Auction(&_NounsAuctionHouse.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Duration() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.Duration(&_NounsAuctionHouse.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Duration() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.Duration(&_NounsAuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) MinBidIncrementPercentage(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "minBidIncrementPercentage")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_NounsAuctionHouse *NounsAuctionHouseSession) MinBidIncrementPercentage() (uint8, error) {
	return _NounsAuctionHouse.Contract.MinBidIncrementPercentage(&_NounsAuctionHouse.CallOpts)
}

// MinBidIncrementPercentage is a free data retrieval call binding the contract method 0xb296024d.
//
// Solidity: function minBidIncrementPercentage() view returns(uint8)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) MinBidIncrementPercentage() (uint8, error) {
	return _NounsAuctionHouse.Contract.MinBidIncrementPercentage(&_NounsAuctionHouse.CallOpts)
}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Nouns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "nouns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Nouns() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Nouns(&_NounsAuctionHouse.CallOpts)
}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Nouns() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Nouns(&_NounsAuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Owner() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Owner(&_NounsAuctionHouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Owner() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Owner(&_NounsAuctionHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Paused() (bool, error) {
	return _NounsAuctionHouse.Contract.Paused(&_NounsAuctionHouse.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Paused() (bool, error) {
	return _NounsAuctionHouse.Contract.Paused(&_NounsAuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) ReservePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "reservePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseSession) ReservePrice() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.ReservePrice(&_NounsAuctionHouse.CallOpts)
}

// ReservePrice is a free data retrieval call binding the contract method 0xdb2e1eed.
//
// Solidity: function reservePrice() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) ReservePrice() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.ReservePrice(&_NounsAuctionHouse.CallOpts)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) TimeBuffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "timeBuffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseSession) TimeBuffer() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.TimeBuffer(&_NounsAuctionHouse.CallOpts)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint256)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) TimeBuffer() (*big.Int, error) {
	return _NounsAuctionHouse.Contract.TimeBuffer(&_NounsAuctionHouse.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsAuctionHouse.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseSession) Weth() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Weth(&_NounsAuctionHouse.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_NounsAuctionHouse *NounsAuctionHouseCallerSession) Weth() (common.Address, error) {
	return _NounsAuctionHouse.Contract.Weth(&_NounsAuctionHouse.CallOpts)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 nounId) payable returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) CreateBid(opts *bind.TransactOpts, nounId *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "createBid", nounId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 nounId) payable returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) CreateBid(nounId *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.CreateBid(&_NounsAuctionHouse.TransactOpts, nounId)
}

// CreateBid is a paid mutator transaction binding the contract method 0x659dd2b4.
//
// Solidity: function createBid(uint256 nounId) payable returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) CreateBid(nounId *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.CreateBid(&_NounsAuctionHouse.TransactOpts, nounId)
}

// Initialize is a paid mutator transaction binding the contract method 0x87f49f54.
//
// Solidity: function initialize(address _nouns, address _weth, uint256 _timeBuffer, uint256 _reservePrice, uint8 _minBidIncrementPercentage, uint256 _duration) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) Initialize(opts *bind.TransactOpts, _nouns common.Address, _weth common.Address, _timeBuffer *big.Int, _reservePrice *big.Int, _minBidIncrementPercentage uint8, _duration *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "initialize", _nouns, _weth, _timeBuffer, _reservePrice, _minBidIncrementPercentage, _duration)
}

// Initialize is a paid mutator transaction binding the contract method 0x87f49f54.
//
// Solidity: function initialize(address _nouns, address _weth, uint256 _timeBuffer, uint256 _reservePrice, uint8 _minBidIncrementPercentage, uint256 _duration) returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) Initialize(_nouns common.Address, _weth common.Address, _timeBuffer *big.Int, _reservePrice *big.Int, _minBidIncrementPercentage uint8, _duration *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Initialize(&_NounsAuctionHouse.TransactOpts, _nouns, _weth, _timeBuffer, _reservePrice, _minBidIncrementPercentage, _duration)
}

// Initialize is a paid mutator transaction binding the contract method 0x87f49f54.
//
// Solidity: function initialize(address _nouns, address _weth, uint256 _timeBuffer, uint256 _reservePrice, uint8 _minBidIncrementPercentage, uint256 _duration) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) Initialize(_nouns common.Address, _weth common.Address, _timeBuffer *big.Int, _reservePrice *big.Int, _minBidIncrementPercentage uint8, _duration *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Initialize(&_NounsAuctionHouse.TransactOpts, _nouns, _weth, _timeBuffer, _reservePrice, _minBidIncrementPercentage, _duration)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) Pause() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Pause(&_NounsAuctionHouse.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) Pause() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Pause(&_NounsAuctionHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) RenounceOwnership() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.RenounceOwnership(&_NounsAuctionHouse.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.RenounceOwnership(&_NounsAuctionHouse.TransactOpts)
}

// SetMinBidIncrementPercentage is a paid mutator transaction binding the contract method 0x36ebdb38.
//
// Solidity: function setMinBidIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) SetMinBidIncrementPercentage(opts *bind.TransactOpts, _minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "setMinBidIncrementPercentage", _minBidIncrementPercentage)
}

// SetMinBidIncrementPercentage is a paid mutator transaction binding the contract method 0x36ebdb38.
//
// Solidity: function setMinBidIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) SetMinBidIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetMinBidIncrementPercentage(&_NounsAuctionHouse.TransactOpts, _minBidIncrementPercentage)
}

// SetMinBidIncrementPercentage is a paid mutator transaction binding the contract method 0x36ebdb38.
//
// Solidity: function setMinBidIncrementPercentage(uint8 _minBidIncrementPercentage) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) SetMinBidIncrementPercentage(_minBidIncrementPercentage uint8) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetMinBidIncrementPercentage(&_NounsAuctionHouse.TransactOpts, _minBidIncrementPercentage)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) SetReservePrice(opts *bind.TransactOpts, _reservePrice *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "setReservePrice", _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetReservePrice(&_NounsAuctionHouse.TransactOpts, _reservePrice)
}

// SetReservePrice is a paid mutator transaction binding the contract method 0xce9c7c0d.
//
// Solidity: function setReservePrice(uint256 _reservePrice) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) SetReservePrice(_reservePrice *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetReservePrice(&_NounsAuctionHouse.TransactOpts, _reservePrice)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) SetTimeBuffer(opts *bind.TransactOpts, _timeBuffer *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "setTimeBuffer", _timeBuffer)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) SetTimeBuffer(_timeBuffer *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetTimeBuffer(&_NounsAuctionHouse.TransactOpts, _timeBuffer)
}

// SetTimeBuffer is a paid mutator transaction binding the contract method 0x7120334b.
//
// Solidity: function setTimeBuffer(uint256 _timeBuffer) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) SetTimeBuffer(_timeBuffer *big.Int) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SetTimeBuffer(&_NounsAuctionHouse.TransactOpts, _timeBuffer)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) SettleAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "settleAuction")
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) SettleAuction() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SettleAuction(&_NounsAuctionHouse.TransactOpts)
}

// SettleAuction is a paid mutator transaction binding the contract method 0xa4d0a17e.
//
// Solidity: function settleAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) SettleAuction() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SettleAuction(&_NounsAuctionHouse.TransactOpts)
}

// SettleCurrentAndCreateNewAuction is a paid mutator transaction binding the contract method 0xf25efffc.
//
// Solidity: function settleCurrentAndCreateNewAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) SettleCurrentAndCreateNewAuction(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "settleCurrentAndCreateNewAuction")
}

// SettleCurrentAndCreateNewAuction is a paid mutator transaction binding the contract method 0xf25efffc.
//
// Solidity: function settleCurrentAndCreateNewAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) SettleCurrentAndCreateNewAuction() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SettleCurrentAndCreateNewAuction(&_NounsAuctionHouse.TransactOpts)
}

// SettleCurrentAndCreateNewAuction is a paid mutator transaction binding the contract method 0xf25efffc.
//
// Solidity: function settleCurrentAndCreateNewAuction() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) SettleCurrentAndCreateNewAuction() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.SettleCurrentAndCreateNewAuction(&_NounsAuctionHouse.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.TransferOwnership(&_NounsAuctionHouse.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.TransferOwnership(&_NounsAuctionHouse.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsAuctionHouse.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseSession) Unpause() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Unpause(&_NounsAuctionHouse.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NounsAuctionHouse *NounsAuctionHouseTransactorSession) Unpause() (*types.Transaction, error) {
	return _NounsAuctionHouse.Contract.Unpause(&_NounsAuctionHouse.TransactOpts)
}

// NounsAuctionHouseAuctionBidIterator is returned from FilterAuctionBid and is used to iterate over the raw logs and unpacked data for AuctionBid events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionBidIterator struct {
	Event *NounsAuctionHouseAuctionBid // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionBid)
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
		it.Event = new(NounsAuctionHouseAuctionBid)
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
func (it *NounsAuctionHouseAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionBid represents a AuctionBid event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionBid struct {
	NounId   *big.Int
	Sender   common.Address
	Value    *big.Int
	Extended bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAuctionBid is a free log retrieval operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value, bool extended)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionBid(opts *bind.FilterOpts, nounId []*big.Int) (*NounsAuctionHouseAuctionBidIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionBidIterator{contract: _NounsAuctionHouse.contract, event: "AuctionBid", logs: logs, sub: sub}, nil
}

// WatchAuctionBid is a free log subscription operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value, bool extended)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionBid(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionBid, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionBid", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionBid)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionBid", log); err != nil {
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

// ParseAuctionBid is a log parse operation binding the contract event 0x1159164c56f277e6fc99c11731bd380e0347deb969b75523398734c252706ea3.
//
// Solidity: event AuctionBid(uint256 indexed nounId, address sender, uint256 value, bool extended)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionBid(log types.Log) (*NounsAuctionHouseAuctionBid, error) {
	event := new(NounsAuctionHouseAuctionBid)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionCreatedIterator struct {
	Event *NounsAuctionHouseAuctionCreated // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionCreated)
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
		it.Event = new(NounsAuctionHouseAuctionCreated)
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
func (it *NounsAuctionHouseAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionCreated represents a AuctionCreated event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionCreated struct {
	NounId    *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionCreated(opts *bind.FilterOpts, nounId []*big.Int) (*NounsAuctionHouseAuctionCreatedIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionCreatedIterator{contract: _NounsAuctionHouse.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionCreated, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionCreated", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionCreated)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xd6eddd1118d71820909c1197aa966dbc15ed6f508554252169cc3d5ccac756ca.
//
// Solidity: event AuctionCreated(uint256 indexed nounId, uint256 startTime, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionCreated(log types.Log) (*NounsAuctionHouseAuctionCreated, error) {
	event := new(NounsAuctionHouseAuctionCreated)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionExtendedIterator is returned from FilterAuctionExtended and is used to iterate over the raw logs and unpacked data for AuctionExtended events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionExtendedIterator struct {
	Event *NounsAuctionHouseAuctionExtended // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionExtended)
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
		it.Event = new(NounsAuctionHouseAuctionExtended)
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
func (it *NounsAuctionHouseAuctionExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionExtended represents a AuctionExtended event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionExtended struct {
	NounId  *big.Int
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionExtended is a free log retrieval operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionExtended(opts *bind.FilterOpts, nounId []*big.Int) (*NounsAuctionHouseAuctionExtendedIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionExtended", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionExtendedIterator{contract: _NounsAuctionHouse.contract, event: "AuctionExtended", logs: logs, sub: sub}, nil
}

// WatchAuctionExtended is a free log subscription operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionExtended(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionExtended, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionExtended", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionExtended)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
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

// ParseAuctionExtended is a log parse operation binding the contract event 0x6e912a3a9105bdd2af817ba5adc14e6c127c1035b5b648faa29ca0d58ab8ff4e.
//
// Solidity: event AuctionExtended(uint256 indexed nounId, uint256 endTime)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionExtended(log types.Log) (*NounsAuctionHouseAuctionExtended, error) {
	event := new(NounsAuctionHouseAuctionExtended)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator is returned from FilterAuctionMinBidIncrementPercentageUpdated and is used to iterate over the raw logs and unpacked data for AuctionMinBidIncrementPercentageUpdated events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator struct {
	Event *NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
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
		it.Event = new(NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
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
func (it *NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated represents a AuctionMinBidIncrementPercentageUpdated event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated struct {
	MinBidIncrementPercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterAuctionMinBidIncrementPercentageUpdated is a free log retrieval operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionMinBidIncrementPercentageUpdated(opts *bind.FilterOpts) (*NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator, error) {

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionMinBidIncrementPercentageUpdatedIterator{contract: _NounsAuctionHouse.contract, event: "AuctionMinBidIncrementPercentageUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionMinBidIncrementPercentageUpdated is a free log subscription operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionMinBidIncrementPercentageUpdated(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionMinBidIncrementPercentageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
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

// ParseAuctionMinBidIncrementPercentageUpdated is a log parse operation binding the contract event 0xec5ccd96cc77b6219e9d44143df916af68fc169339ea7de5008ff15eae13450d.
//
// Solidity: event AuctionMinBidIncrementPercentageUpdated(uint256 minBidIncrementPercentage)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionMinBidIncrementPercentageUpdated(log types.Log) (*NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated, error) {
	event := new(NounsAuctionHouseAuctionMinBidIncrementPercentageUpdated)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionMinBidIncrementPercentageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionReservePriceUpdatedIterator is returned from FilterAuctionReservePriceUpdated and is used to iterate over the raw logs and unpacked data for AuctionReservePriceUpdated events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionReservePriceUpdatedIterator struct {
	Event *NounsAuctionHouseAuctionReservePriceUpdated // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionReservePriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionReservePriceUpdated)
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
		it.Event = new(NounsAuctionHouseAuctionReservePriceUpdated)
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
func (it *NounsAuctionHouseAuctionReservePriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionReservePriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionReservePriceUpdated represents a AuctionReservePriceUpdated event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionReservePriceUpdated struct {
	ReservePrice *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionReservePriceUpdated is a free log retrieval operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionReservePriceUpdated(opts *bind.FilterOpts) (*NounsAuctionHouseAuctionReservePriceUpdatedIterator, error) {

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionReservePriceUpdatedIterator{contract: _NounsAuctionHouse.contract, event: "AuctionReservePriceUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionReservePriceUpdated is a free log subscription operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionReservePriceUpdated(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionReservePriceUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionReservePriceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionReservePriceUpdated)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
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

// ParseAuctionReservePriceUpdated is a log parse operation binding the contract event 0x6ab2e127d7fdf53b8f304e59d3aab5bfe97979f52a85479691a6fab27a28a6b2.
//
// Solidity: event AuctionReservePriceUpdated(uint256 reservePrice)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionReservePriceUpdated(log types.Log) (*NounsAuctionHouseAuctionReservePriceUpdated, error) {
	event := new(NounsAuctionHouseAuctionReservePriceUpdated)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionReservePriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionSettledIterator is returned from FilterAuctionSettled and is used to iterate over the raw logs and unpacked data for AuctionSettled events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionSettledIterator struct {
	Event *NounsAuctionHouseAuctionSettled // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionSettled)
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
		it.Event = new(NounsAuctionHouseAuctionSettled)
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
func (it *NounsAuctionHouseAuctionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionSettled represents a AuctionSettled event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionSettled struct {
	NounId *big.Int
	Winner common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAuctionSettled is a free log retrieval operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionSettled(opts *bind.FilterOpts, nounId []*big.Int) (*NounsAuctionHouseAuctionSettledIterator, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionSettledIterator{contract: _NounsAuctionHouse.contract, event: "AuctionSettled", logs: logs, sub: sub}, nil
}

// WatchAuctionSettled is a free log subscription operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionSettled(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionSettled, nounId []*big.Int) (event.Subscription, error) {

	var nounIdRule []interface{}
	for _, nounIdItem := range nounId {
		nounIdRule = append(nounIdRule, nounIdItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionSettled", nounIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionSettled)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
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

// ParseAuctionSettled is a log parse operation binding the contract event 0xc9f72b276a388619c6d185d146697036241880c36654b1a3ffdad07c24038d99.
//
// Solidity: event AuctionSettled(uint256 indexed nounId, address winner, uint256 amount)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionSettled(log types.Log) (*NounsAuctionHouseAuctionSettled, error) {
	event := new(NounsAuctionHouseAuctionSettled)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseAuctionTimeBufferUpdatedIterator is returned from FilterAuctionTimeBufferUpdated and is used to iterate over the raw logs and unpacked data for AuctionTimeBufferUpdated events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionTimeBufferUpdatedIterator struct {
	Event *NounsAuctionHouseAuctionTimeBufferUpdated // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseAuctionTimeBufferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseAuctionTimeBufferUpdated)
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
		it.Event = new(NounsAuctionHouseAuctionTimeBufferUpdated)
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
func (it *NounsAuctionHouseAuctionTimeBufferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseAuctionTimeBufferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseAuctionTimeBufferUpdated represents a AuctionTimeBufferUpdated event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseAuctionTimeBufferUpdated struct {
	TimeBuffer *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionTimeBufferUpdated is a free log retrieval operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterAuctionTimeBufferUpdated(opts *bind.FilterOpts) (*NounsAuctionHouseAuctionTimeBufferUpdatedIterator, error) {

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseAuctionTimeBufferUpdatedIterator{contract: _NounsAuctionHouse.contract, event: "AuctionTimeBufferUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionTimeBufferUpdated is a free log subscription operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchAuctionTimeBufferUpdated(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseAuctionTimeBufferUpdated) (event.Subscription, error) {

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "AuctionTimeBufferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseAuctionTimeBufferUpdated)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
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

// ParseAuctionTimeBufferUpdated is a log parse operation binding the contract event 0x1b55d9f7002bda4490f467e326f22a4a847629c0f2d1ed421607d318d25b410d.
//
// Solidity: event AuctionTimeBufferUpdated(uint256 timeBuffer)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseAuctionTimeBufferUpdated(log types.Log) (*NounsAuctionHouseAuctionTimeBufferUpdated, error) {
	event := new(NounsAuctionHouseAuctionTimeBufferUpdated)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "AuctionTimeBufferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseOwnershipTransferredIterator struct {
	Event *NounsAuctionHouseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseOwnershipTransferred)
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
		it.Event = new(NounsAuctionHouseOwnershipTransferred)
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
func (it *NounsAuctionHouseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseOwnershipTransferred represents a OwnershipTransferred event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NounsAuctionHouseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseOwnershipTransferredIterator{contract: _NounsAuctionHouse.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseOwnershipTransferred)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseOwnershipTransferred(log types.Log) (*NounsAuctionHouseOwnershipTransferred, error) {
	event := new(NounsAuctionHouseOwnershipTransferred)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHousePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NounsAuctionHouse contract.
type NounsAuctionHousePausedIterator struct {
	Event *NounsAuctionHousePaused // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHousePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHousePaused)
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
		it.Event = new(NounsAuctionHousePaused)
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
func (it *NounsAuctionHousePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHousePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHousePaused represents a Paused event raised by the NounsAuctionHouse contract.
type NounsAuctionHousePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterPaused(opts *bind.FilterOpts) (*NounsAuctionHousePausedIterator, error) {

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHousePausedIterator{contract: _NounsAuctionHouse.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NounsAuctionHousePaused) (event.Subscription, error) {

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHousePaused)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParsePaused(log types.Log) (*NounsAuctionHousePaused, error) {
	event := new(NounsAuctionHousePaused)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsAuctionHouseUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NounsAuctionHouse contract.
type NounsAuctionHouseUnpausedIterator struct {
	Event *NounsAuctionHouseUnpaused // Event containing the contract specifics and raw log

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
func (it *NounsAuctionHouseUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsAuctionHouseUnpaused)
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
		it.Event = new(NounsAuctionHouseUnpaused)
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
func (it *NounsAuctionHouseUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsAuctionHouseUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsAuctionHouseUnpaused represents a Unpaused event raised by the NounsAuctionHouse contract.
type NounsAuctionHouseUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NounsAuctionHouseUnpausedIterator, error) {

	logs, sub, err := _NounsAuctionHouse.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NounsAuctionHouseUnpausedIterator{contract: _NounsAuctionHouse.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NounsAuctionHouseUnpaused) (event.Subscription, error) {

	logs, sub, err := _NounsAuctionHouse.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsAuctionHouseUnpaused)
				if err := _NounsAuctionHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NounsAuctionHouse *NounsAuctionHouseFilterer) ParseUnpaused(log types.Log) (*NounsAuctionHouseUnpaused, error) {
	event := new(NounsAuctionHouseUnpaused)
	if err := _NounsAuctionHouse.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
