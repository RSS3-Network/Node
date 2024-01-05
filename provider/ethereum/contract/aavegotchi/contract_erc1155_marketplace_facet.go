// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavegotchi

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

// ERC1155Listing is an auto generated low-level Go binding around an user-defined struct.
type ERC1155Listing struct {
	ListingId           *big.Int
	Seller              common.Address
	Erc1155TokenAddress common.Address
	Erc1155TypeId       *big.Int
	Category            *big.Int
	Quantity            *big.Int
	PriceInWei          *big.Int
	TimeCreated         *big.Int
	TimeLastPurchased   *big.Int
	SourceListingId     *big.Int
	Sold                bool
	Cancelled           bool
	PrincipalSplit      [2]uint16
	Affiliate           common.Address
}

// ERC1155MarketplaceFacetCategory is an auto generated low-level Go binding around an user-defined struct.
type ERC1155MarketplaceFacetCategory struct {
	Erc1155TokenAddress common.Address
	Erc1155TypeId       *big.Int
	Category            *big.Int
}

// ERC1155MarketplaceMetaData contains all meta data concerning the ERC1155Marketplace contract.
var ERC1155MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingFeeInWei\",\"type\":\"uint256\"}],\"name\":\"ChangedListingFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERC1155ExecutedListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC1155ExecutedToRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERC1155ListingAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"ERC1155ListingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"ERC1155ListingSplit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"cancelERC1155Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_listingIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelERC1155Listings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"}],\"name\":\"executeERC1155Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_itemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"executeERC1155ListingToRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc1155TypeId\",\"type\":\"uint256\"}],\"name\":\"getERC1155Category\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"category_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"getERC1155Listing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeLastPurchased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceListingId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC1155Listing\",\"name\":\"listing_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC1155ListingFromToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeLastPurchased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceListingId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC1155Listing\",\"name\":\"listing_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getERC1155Listings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeLastPurchased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceListingId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC1155Listing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getOwnerERC1155Listings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeLastPurchased\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceListingId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC1155Listing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"}],\"internalType\":\"structERC1155MarketplaceFacet.Category[]\",\"name\":\"_categories\",\"type\":\"tuple[]\"}],\"name\":\"setERC1155Categories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"}],\"name\":\"setERC1155Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint16[2]\",\"name\":\"_principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"setERC1155ListingWithSplit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingFeeInWei\",\"type\":\"uint256\"}],\"name\":\"setListingFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_erc1155TypeIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"updateBatchERC1155Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc1155TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc1155TypeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"updateERC1155Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC1155MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1155MarketplaceMetaData.ABI instead.
var ERC1155MarketplaceABI = ERC1155MarketplaceMetaData.ABI

// ERC1155Marketplace is an auto generated Go binding around an Ethereum contract.
type ERC1155Marketplace struct {
	ERC1155MarketplaceCaller     // Read-only binding to the contract
	ERC1155MarketplaceTransactor // Write-only binding to the contract
	ERC1155MarketplaceFilterer   // Log filterer for contract events
}

// ERC1155MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1155MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1155MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1155MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1155MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1155MarketplaceSession struct {
	Contract     *ERC1155Marketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ERC1155MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1155MarketplaceCallerSession struct {
	Contract *ERC1155MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ERC1155MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1155MarketplaceTransactorSession struct {
	Contract     *ERC1155MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC1155MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1155MarketplaceRaw struct {
	Contract *ERC1155Marketplace // Generic contract binding to access the raw methods on
}

// ERC1155MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1155MarketplaceCallerRaw struct {
	Contract *ERC1155MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1155MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1155MarketplaceTransactorRaw struct {
	Contract *ERC1155MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1155Marketplace creates a new instance of ERC1155Marketplace, bound to a specific deployed contract.
func NewERC1155Marketplace(address common.Address, backend bind.ContractBackend) (*ERC1155Marketplace, error) {
	contract, err := bindERC1155Marketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1155Marketplace{ERC1155MarketplaceCaller: ERC1155MarketplaceCaller{contract: contract}, ERC1155MarketplaceTransactor: ERC1155MarketplaceTransactor{contract: contract}, ERC1155MarketplaceFilterer: ERC1155MarketplaceFilterer{contract: contract}}, nil
}

// NewERC1155MarketplaceCaller creates a new read-only instance of ERC1155Marketplace, bound to a specific deployed contract.
func NewERC1155MarketplaceCaller(address common.Address, caller bind.ContractCaller) (*ERC1155MarketplaceCaller, error) {
	contract, err := bindERC1155Marketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceCaller{contract: contract}, nil
}

// NewERC1155MarketplaceTransactor creates a new write-only instance of ERC1155Marketplace, bound to a specific deployed contract.
func NewERC1155MarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1155MarketplaceTransactor, error) {
	contract, err := bindERC1155Marketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceTransactor{contract: contract}, nil
}

// NewERC1155MarketplaceFilterer creates a new log filterer instance of ERC1155Marketplace, bound to a specific deployed contract.
func NewERC1155MarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1155MarketplaceFilterer, error) {
	contract, err := bindERC1155Marketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceFilterer{contract: contract}, nil
}

// bindERC1155Marketplace binds a generic wrapper to an already deployed contract.
func bindERC1155Marketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1155MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Marketplace *ERC1155MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Marketplace.Contract.ERC1155MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Marketplace *ERC1155MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ERC1155MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Marketplace *ERC1155MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ERC1155MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1155Marketplace *ERC1155MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1155Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.contract.Transact(opts, method, params...)
}

// GetERC1155Category is a free data retrieval call binding the contract method 0x446978f0.
//
// Solidity: function getERC1155Category(address _erc1155TokenAddress, uint256 _erc1155TypeId) view returns(uint256 category_)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetERC1155Category(opts *bind.CallOpts, _erc1155TokenAddress common.Address, _erc1155TypeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getERC1155Category", _erc1155TokenAddress, _erc1155TypeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERC1155Category is a free data retrieval call binding the contract method 0x446978f0.
//
// Solidity: function getERC1155Category(address _erc1155TokenAddress, uint256 _erc1155TypeId) view returns(uint256 category_)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetERC1155Category(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int) (*big.Int, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Category(&_ERC1155Marketplace.CallOpts, _erc1155TokenAddress, _erc1155TypeId)
}

// GetERC1155Category is a free data retrieval call binding the contract method 0x446978f0.
//
// Solidity: function getERC1155Category(address _erc1155TokenAddress, uint256 _erc1155TypeId) view returns(uint256 category_)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetERC1155Category(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int) (*big.Int, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Category(&_ERC1155Marketplace.CallOpts, _erc1155TokenAddress, _erc1155TypeId)
}

// GetERC1155Listing is a free data retrieval call binding the contract method 0x09c415af.
//
// Solidity: function getERC1155Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetERC1155Listing(opts *bind.CallOpts, _listingId *big.Int) (ERC1155Listing, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getERC1155Listing", _listingId)

	if err != nil {
		return *new(ERC1155Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC1155Listing)).(*ERC1155Listing)

	return out0, err

}

// GetERC1155Listing is a free data retrieval call binding the contract method 0x09c415af.
//
// Solidity: function getERC1155Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetERC1155Listing(_listingId *big.Int) (ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Listing(&_ERC1155Marketplace.CallOpts, _listingId)
}

// GetERC1155Listing is a free data retrieval call binding the contract method 0x09c415af.
//
// Solidity: function getERC1155Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetERC1155Listing(_listingId *big.Int) (ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Listing(&_ERC1155Marketplace.CallOpts, _listingId)
}

// GetERC1155ListingFromToken is a free data retrieval call binding the contract method 0x519f62ed.
//
// Solidity: function getERC1155ListingFromToken(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetERC1155ListingFromToken(opts *bind.CallOpts, _erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (ERC1155Listing, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getERC1155ListingFromToken", _erc1155TokenAddress, _erc1155TypeId, _owner)

	if err != nil {
		return *new(ERC1155Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC1155Listing)).(*ERC1155Listing)

	return out0, err

}

// GetERC1155ListingFromToken is a free data retrieval call binding the contract method 0x519f62ed.
//
// Solidity: function getERC1155ListingFromToken(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetERC1155ListingFromToken(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155ListingFromToken(&_ERC1155Marketplace.CallOpts, _erc1155TokenAddress, _erc1155TypeId, _owner)
}

// GetERC1155ListingFromToken is a free data retrieval call binding the contract method 0x519f62ed.
//
// Solidity: function getERC1155ListingFromToken(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address) listing_)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetERC1155ListingFromToken(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155ListingFromToken(&_ERC1155Marketplace.CallOpts, _erc1155TokenAddress, _erc1155TypeId, _owner)
}

// GetERC1155Listings is a free data retrieval call binding the contract method 0x12542188.
//
// Solidity: function getERC1155Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetERC1155Listings(opts *bind.CallOpts, _category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getERC1155Listings", _category, _sort, _length)

	if err != nil {
		return *new([]ERC1155Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC1155Listing)).(*[]ERC1155Listing)

	return out0, err

}

// GetERC1155Listings is a free data retrieval call binding the contract method 0x12542188.
//
// Solidity: function getERC1155Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetERC1155Listings(_category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Listings(&_ERC1155Marketplace.CallOpts, _category, _sort, _length)
}

// GetERC1155Listings is a free data retrieval call binding the contract method 0x12542188.
//
// Solidity: function getERC1155Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetERC1155Listings(_category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetERC1155Listings(&_ERC1155Marketplace.CallOpts, _category, _sort, _length)
}

// GetListingFeeInWei is a free data retrieval call binding the contract method 0x962095a1.
//
// Solidity: function getListingFeeInWei() view returns(uint256)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetListingFeeInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getListingFeeInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingFeeInWei is a free data retrieval call binding the contract method 0x962095a1.
//
// Solidity: function getListingFeeInWei() view returns(uint256)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetListingFeeInWei() (*big.Int, error) {
	return _ERC1155Marketplace.Contract.GetListingFeeInWei(&_ERC1155Marketplace.CallOpts)
}

// GetListingFeeInWei is a free data retrieval call binding the contract method 0x962095a1.
//
// Solidity: function getListingFeeInWei() view returns(uint256)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetListingFeeInWei() (*big.Int, error) {
	return _ERC1155Marketplace.Contract.GetListingFeeInWei(&_ERC1155Marketplace.CallOpts)
}

// GetOwnerERC1155Listings is a free data retrieval call binding the contract method 0x781f11d4.
//
// Solidity: function getOwnerERC1155Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceCaller) GetOwnerERC1155Listings(opts *bind.CallOpts, _owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	var out []interface{}
	err := _ERC1155Marketplace.contract.Call(opts, &out, "getOwnerERC1155Listings", _owner, _category, _sort, _length)

	if err != nil {
		return *new([]ERC1155Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC1155Listing)).(*[]ERC1155Listing)

	return out0, err

}

// GetOwnerERC1155Listings is a free data retrieval call binding the contract method 0x781f11d4.
//
// Solidity: function getOwnerERC1155Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceSession) GetOwnerERC1155Listings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetOwnerERC1155Listings(&_ERC1155Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// GetOwnerERC1155Listings is a free data retrieval call binding the contract method 0x781f11d4.
//
// Solidity: function getOwnerERC1155Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,uint16[2],address)[] listings_)
func (_ERC1155Marketplace *ERC1155MarketplaceCallerSession) GetOwnerERC1155Listings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC1155Listing, error) {
	return _ERC1155Marketplace.Contract.GetOwnerERC1155Listings(&_ERC1155Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// CancelERC1155Listing is a paid mutator transaction binding the contract method 0x08642e24.
//
// Solidity: function cancelERC1155Listing(uint256 _listingId) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) CancelERC1155Listing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "cancelERC1155Listing", _listingId)
}

// CancelERC1155Listing is a paid mutator transaction binding the contract method 0x08642e24.
//
// Solidity: function cancelERC1155Listing(uint256 _listingId) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) CancelERC1155Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.CancelERC1155Listing(&_ERC1155Marketplace.TransactOpts, _listingId)
}

// CancelERC1155Listing is a paid mutator transaction binding the contract method 0x08642e24.
//
// Solidity: function cancelERC1155Listing(uint256 _listingId) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) CancelERC1155Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.CancelERC1155Listing(&_ERC1155Marketplace.TransactOpts, _listingId)
}

// CancelERC1155Listings is a paid mutator transaction binding the contract method 0x17ca3681.
//
// Solidity: function cancelERC1155Listings(uint256[] _listingIds) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) CancelERC1155Listings(opts *bind.TransactOpts, _listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "cancelERC1155Listings", _listingIds)
}

// CancelERC1155Listings is a paid mutator transaction binding the contract method 0x17ca3681.
//
// Solidity: function cancelERC1155Listings(uint256[] _listingIds) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) CancelERC1155Listings(_listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.CancelERC1155Listings(&_ERC1155Marketplace.TransactOpts, _listingIds)
}

// CancelERC1155Listings is a paid mutator transaction binding the contract method 0x17ca3681.
//
// Solidity: function cancelERC1155Listings(uint256[] _listingIds) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) CancelERC1155Listings(_listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.CancelERC1155Listings(&_ERC1155Marketplace.TransactOpts, _listingIds)
}

// ExecuteERC1155Listing is a paid mutator transaction binding the contract method 0x575ae876.
//
// Solidity: function executeERC1155Listing(uint256 _listingId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) ExecuteERC1155Listing(opts *bind.TransactOpts, _listingId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "executeERC1155Listing", _listingId, _quantity, _priceInWei)
}

// ExecuteERC1155Listing is a paid mutator transaction binding the contract method 0x575ae876.
//
// Solidity: function executeERC1155Listing(uint256 _listingId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) ExecuteERC1155Listing(_listingId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ExecuteERC1155Listing(&_ERC1155Marketplace.TransactOpts, _listingId, _quantity, _priceInWei)
}

// ExecuteERC1155Listing is a paid mutator transaction binding the contract method 0x575ae876.
//
// Solidity: function executeERC1155Listing(uint256 _listingId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) ExecuteERC1155Listing(_listingId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ExecuteERC1155Listing(&_ERC1155Marketplace.TransactOpts, _listingId, _quantity, _priceInWei)
}

// ExecuteERC1155ListingToRecipient is a paid mutator transaction binding the contract method 0x68edd211.
//
// Solidity: function executeERC1155ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _itemId, uint256 _quantity, uint256 _priceInWei, address _recipient) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) ExecuteERC1155ListingToRecipient(opts *bind.TransactOpts, _listingId *big.Int, _contractAddress common.Address, _itemId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "executeERC1155ListingToRecipient", _listingId, _contractAddress, _itemId, _quantity, _priceInWei, _recipient)
}

// ExecuteERC1155ListingToRecipient is a paid mutator transaction binding the contract method 0x68edd211.
//
// Solidity: function executeERC1155ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _itemId, uint256 _quantity, uint256 _priceInWei, address _recipient) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) ExecuteERC1155ListingToRecipient(_listingId *big.Int, _contractAddress common.Address, _itemId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ExecuteERC1155ListingToRecipient(&_ERC1155Marketplace.TransactOpts, _listingId, _contractAddress, _itemId, _quantity, _priceInWei, _recipient)
}

// ExecuteERC1155ListingToRecipient is a paid mutator transaction binding the contract method 0x68edd211.
//
// Solidity: function executeERC1155ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _itemId, uint256 _quantity, uint256 _priceInWei, address _recipient) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) ExecuteERC1155ListingToRecipient(_listingId *big.Int, _contractAddress common.Address, _itemId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.ExecuteERC1155ListingToRecipient(&_ERC1155Marketplace.TransactOpts, _listingId, _contractAddress, _itemId, _quantity, _priceInWei, _recipient)
}

// SetERC1155Categories is a paid mutator transaction binding the contract method 0xdc1265f8.
//
// Solidity: function setERC1155Categories((address,uint256,uint256)[] _categories) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) SetERC1155Categories(opts *bind.TransactOpts, _categories []ERC1155MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "setERC1155Categories", _categories)
}

// SetERC1155Categories is a paid mutator transaction binding the contract method 0xdc1265f8.
//
// Solidity: function setERC1155Categories((address,uint256,uint256)[] _categories) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) SetERC1155Categories(_categories []ERC1155MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155Categories(&_ERC1155Marketplace.TransactOpts, _categories)
}

// SetERC1155Categories is a paid mutator transaction binding the contract method 0xdc1265f8.
//
// Solidity: function setERC1155Categories((address,uint256,uint256)[] _categories) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) SetERC1155Categories(_categories []ERC1155MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155Categories(&_ERC1155Marketplace.TransactOpts, _categories)
}

// SetERC1155Listing is a paid mutator transaction binding the contract method 0xe8da2bfa.
//
// Solidity: function setERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) SetERC1155Listing(opts *bind.TransactOpts, _erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "setERC1155Listing", _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei)
}

// SetERC1155Listing is a paid mutator transaction binding the contract method 0xe8da2bfa.
//
// Solidity: function setERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) SetERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei)
}

// SetERC1155Listing is a paid mutator transaction binding the contract method 0xe8da2bfa.
//
// Solidity: function setERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) SetERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei)
}

// SetERC1155ListingWithSplit is a paid mutator transaction binding the contract method 0x92d718f6.
//
// Solidity: function setERC1155ListingWithSplit(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) SetERC1155ListingWithSplit(opts *bind.TransactOpts, _erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "setERC1155ListingWithSplit", _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei, _principalSplit, _affiliate)
}

// SetERC1155ListingWithSplit is a paid mutator transaction binding the contract method 0x92d718f6.
//
// Solidity: function setERC1155ListingWithSplit(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) SetERC1155ListingWithSplit(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155ListingWithSplit(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei, _principalSplit, _affiliate)
}

// SetERC1155ListingWithSplit is a paid mutator transaction binding the contract method 0x92d718f6.
//
// Solidity: function setERC1155ListingWithSplit(address _erc1155TokenAddress, uint256 _erc1155TypeId, uint256 _quantity, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) SetERC1155ListingWithSplit(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _quantity *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetERC1155ListingWithSplit(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _quantity, _priceInWei, _principalSplit, _affiliate)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _listingFeeInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) SetListingFee(opts *bind.TransactOpts, _listingFeeInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "setListingFee", _listingFeeInWei)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _listingFeeInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) SetListingFee(_listingFeeInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetListingFee(&_ERC1155Marketplace.TransactOpts, _listingFeeInWei)
}

// SetListingFee is a paid mutator transaction binding the contract method 0x131dbd09.
//
// Solidity: function setListingFee(uint256 _listingFeeInWei) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) SetListingFee(_listingFeeInWei *big.Int) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.SetListingFee(&_ERC1155Marketplace.TransactOpts, _listingFeeInWei)
}

// UpdateBatchERC1155Listing is a paid mutator transaction binding the contract method 0xfac43c85.
//
// Solidity: function updateBatchERC1155Listing(address _erc1155TokenAddress, uint256[] _erc1155TypeIds, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) UpdateBatchERC1155Listing(opts *bind.TransactOpts, _erc1155TokenAddress common.Address, _erc1155TypeIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "updateBatchERC1155Listing", _erc1155TokenAddress, _erc1155TypeIds, _owner)
}

// UpdateBatchERC1155Listing is a paid mutator transaction binding the contract method 0xfac43c85.
//
// Solidity: function updateBatchERC1155Listing(address _erc1155TokenAddress, uint256[] _erc1155TypeIds, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) UpdateBatchERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.UpdateBatchERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeIds, _owner)
}

// UpdateBatchERC1155Listing is a paid mutator transaction binding the contract method 0xfac43c85.
//
// Solidity: function updateBatchERC1155Listing(address _erc1155TokenAddress, uint256[] _erc1155TypeIds, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) UpdateBatchERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeIds []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.UpdateBatchERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeIds, _owner)
}

// UpdateERC1155Listing is a paid mutator transaction binding the contract method 0xe9a8887f.
//
// Solidity: function updateERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactor) UpdateERC1155Listing(opts *bind.TransactOpts, _erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.contract.Transact(opts, "updateERC1155Listing", _erc1155TokenAddress, _erc1155TypeId, _owner)
}

// UpdateERC1155Listing is a paid mutator transaction binding the contract method 0xe9a8887f.
//
// Solidity: function updateERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceSession) UpdateERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.UpdateERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _owner)
}

// UpdateERC1155Listing is a paid mutator transaction binding the contract method 0xe9a8887f.
//
// Solidity: function updateERC1155Listing(address _erc1155TokenAddress, uint256 _erc1155TypeId, address _owner) returns()
func (_ERC1155Marketplace *ERC1155MarketplaceTransactorSession) UpdateERC1155Listing(_erc1155TokenAddress common.Address, _erc1155TypeId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC1155Marketplace.Contract.UpdateERC1155Listing(&_ERC1155Marketplace.TransactOpts, _erc1155TokenAddress, _erc1155TypeId, _owner)
}

// ERC1155MarketplaceChangedListingFeeIterator is returned from FilterChangedListingFee and is used to iterate over the raw logs and unpacked data for ChangedListingFee events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceChangedListingFeeIterator struct {
	Event *ERC1155MarketplaceChangedListingFee // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceChangedListingFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceChangedListingFee)
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
		it.Event = new(ERC1155MarketplaceChangedListingFee)
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
func (it *ERC1155MarketplaceChangedListingFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceChangedListingFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceChangedListingFee represents a ChangedListingFee event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceChangedListingFee struct {
	ListingFeeInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChangedListingFee is a free log retrieval operation binding the contract event 0x47110aa6379903a1c93053b7d8b6e8bdbbadd893e457a1e10588c4e2a9575dc7.
//
// Solidity: event ChangedListingFee(uint256 listingFeeInWei)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterChangedListingFee(opts *bind.FilterOpts) (*ERC1155MarketplaceChangedListingFeeIterator, error) {

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ChangedListingFee")
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceChangedListingFeeIterator{contract: _ERC1155Marketplace.contract, event: "ChangedListingFee", logs: logs, sub: sub}, nil
}

// WatchChangedListingFee is a free log subscription operation binding the contract event 0x47110aa6379903a1c93053b7d8b6e8bdbbadd893e457a1e10588c4e2a9575dc7.
//
// Solidity: event ChangedListingFee(uint256 listingFeeInWei)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchChangedListingFee(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceChangedListingFee) (event.Subscription, error) {

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ChangedListingFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceChangedListingFee)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ChangedListingFee", log); err != nil {
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

// ParseChangedListingFee is a log parse operation binding the contract event 0x47110aa6379903a1c93053b7d8b6e8bdbbadd893e457a1e10588c4e2a9575dc7.
//
// Solidity: event ChangedListingFee(uint256 listingFeeInWei)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseChangedListingFee(log types.Log) (*ERC1155MarketplaceChangedListingFee, error) {
	event := new(ERC1155MarketplaceChangedListingFee)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ChangedListingFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MarketplaceERC1155ExecutedListingIterator is returned from FilterERC1155ExecutedListing and is used to iterate over the raw logs and unpacked data for ERC1155ExecutedListing events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ExecutedListingIterator struct {
	Event *ERC1155MarketplaceERC1155ExecutedListing // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceERC1155ExecutedListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceERC1155ExecutedListing)
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
		it.Event = new(ERC1155MarketplaceERC1155ExecutedListing)
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
func (it *ERC1155MarketplaceERC1155ExecutedListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceERC1155ExecutedListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceERC1155ExecutedListing represents a ERC1155ExecutedListing event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ExecutedListing struct {
	ListingId           *big.Int
	Seller              common.Address
	Buyer               common.Address
	Erc1155TokenAddress common.Address
	Erc1155TypeId       *big.Int
	Category            *big.Int
	Quantity            *big.Int
	PriceInWei          *big.Int
	Time                *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC1155ExecutedListing is a free log retrieval operation binding the contract event 0x4004185f45cf5f331fe63afc9aa5aa0a0d8cfa1d6bc5d8c6cb0304136c49515c.
//
// Solidity: event ERC1155ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 _quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterERC1155ExecutedListing(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, category []*big.Int) (*ERC1155MarketplaceERC1155ExecutedListingIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ERC1155ExecutedListing", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceERC1155ExecutedListingIterator{contract: _ERC1155Marketplace.contract, event: "ERC1155ExecutedListing", logs: logs, sub: sub}, nil
}

// WatchERC1155ExecutedListing is a free log subscription operation binding the contract event 0x4004185f45cf5f331fe63afc9aa5aa0a0d8cfa1d6bc5d8c6cb0304136c49515c.
//
// Solidity: event ERC1155ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 _quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchERC1155ExecutedListing(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceERC1155ExecutedListing, listingId []*big.Int, seller []common.Address, category []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ERC1155ExecutedListing", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceERC1155ExecutedListing)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ExecutedListing", log); err != nil {
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

// ParseERC1155ExecutedListing is a log parse operation binding the contract event 0x4004185f45cf5f331fe63afc9aa5aa0a0d8cfa1d6bc5d8c6cb0304136c49515c.
//
// Solidity: event ERC1155ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 _quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseERC1155ExecutedListing(log types.Log) (*ERC1155MarketplaceERC1155ExecutedListing, error) {
	event := new(ERC1155MarketplaceERC1155ExecutedListing)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ExecutedListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MarketplaceERC1155ExecutedToRecipientIterator is returned from FilterERC1155ExecutedToRecipient and is used to iterate over the raw logs and unpacked data for ERC1155ExecutedToRecipient events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ExecutedToRecipientIterator struct {
	Event *ERC1155MarketplaceERC1155ExecutedToRecipient // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceERC1155ExecutedToRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceERC1155ExecutedToRecipient)
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
		it.Event = new(ERC1155MarketplaceERC1155ExecutedToRecipient)
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
func (it *ERC1155MarketplaceERC1155ExecutedToRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceERC1155ExecutedToRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceERC1155ExecutedToRecipient represents a ERC1155ExecutedToRecipient event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ExecutedToRecipient struct {
	ListingId *big.Int
	Buyer     common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155ExecutedToRecipient is a free log retrieval operation binding the contract event 0x6d10597eaea3b6f10739896c93a2cdebf87810d37226b40bc3e433745f55ce9c.
//
// Solidity: event ERC1155ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterERC1155ExecutedToRecipient(opts *bind.FilterOpts, listingId []*big.Int, buyer []common.Address, recipient []common.Address) (*ERC1155MarketplaceERC1155ExecutedToRecipientIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ERC1155ExecutedToRecipient", listingIdRule, buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceERC1155ExecutedToRecipientIterator{contract: _ERC1155Marketplace.contract, event: "ERC1155ExecutedToRecipient", logs: logs, sub: sub}, nil
}

// WatchERC1155ExecutedToRecipient is a free log subscription operation binding the contract event 0x6d10597eaea3b6f10739896c93a2cdebf87810d37226b40bc3e433745f55ce9c.
//
// Solidity: event ERC1155ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchERC1155ExecutedToRecipient(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceERC1155ExecutedToRecipient, listingId []*big.Int, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ERC1155ExecutedToRecipient", listingIdRule, buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceERC1155ExecutedToRecipient)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ExecutedToRecipient", log); err != nil {
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

// ParseERC1155ExecutedToRecipient is a log parse operation binding the contract event 0x6d10597eaea3b6f10739896c93a2cdebf87810d37226b40bc3e433745f55ce9c.
//
// Solidity: event ERC1155ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseERC1155ExecutedToRecipient(log types.Log) (*ERC1155MarketplaceERC1155ExecutedToRecipient, error) {
	event := new(ERC1155MarketplaceERC1155ExecutedToRecipient)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ExecutedToRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MarketplaceERC1155ListingAddIterator is returned from FilterERC1155ListingAdd and is used to iterate over the raw logs and unpacked data for ERC1155ListingAdd events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingAddIterator struct {
	Event *ERC1155MarketplaceERC1155ListingAdd // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceERC1155ListingAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceERC1155ListingAdd)
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
		it.Event = new(ERC1155MarketplaceERC1155ListingAdd)
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
func (it *ERC1155MarketplaceERC1155ListingAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceERC1155ListingAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceERC1155ListingAdd represents a ERC1155ListingAdd event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingAdd struct {
	ListingId           *big.Int
	Seller              common.Address
	Erc1155TokenAddress common.Address
	Erc1155TypeId       *big.Int
	Category            *big.Int
	Quantity            *big.Int
	PriceInWei          *big.Int
	Time                *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC1155ListingAdd is a free log retrieval operation binding the contract event 0xf99a07edc22c88f73b19df09fffa9b11c9ab5c5dd6dfc2f6e103a95f5b40e2fc.
//
// Solidity: event ERC1155ListingAdd(uint256 indexed listingId, address indexed seller, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterERC1155ListingAdd(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, category []*big.Int) (*ERC1155MarketplaceERC1155ListingAddIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ERC1155ListingAdd", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceERC1155ListingAddIterator{contract: _ERC1155Marketplace.contract, event: "ERC1155ListingAdd", logs: logs, sub: sub}, nil
}

// WatchERC1155ListingAdd is a free log subscription operation binding the contract event 0xf99a07edc22c88f73b19df09fffa9b11c9ab5c5dd6dfc2f6e103a95f5b40e2fc.
//
// Solidity: event ERC1155ListingAdd(uint256 indexed listingId, address indexed seller, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchERC1155ListingAdd(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceERC1155ListingAdd, listingId []*big.Int, seller []common.Address, category []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var categoryRule []interface{}
	for _, categoryItem := range category {
		categoryRule = append(categoryRule, categoryItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ERC1155ListingAdd", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceERC1155ListingAdd)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingAdd", log); err != nil {
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

// ParseERC1155ListingAdd is a log parse operation binding the contract event 0xf99a07edc22c88f73b19df09fffa9b11c9ab5c5dd6dfc2f6e103a95f5b40e2fc.
//
// Solidity: event ERC1155ListingAdd(uint256 indexed listingId, address indexed seller, address erc1155TokenAddress, uint256 erc1155TypeId, uint256 indexed category, uint256 quantity, uint256 priceInWei, uint256 time)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseERC1155ListingAdd(log types.Log) (*ERC1155MarketplaceERC1155ListingAdd, error) {
	event := new(ERC1155MarketplaceERC1155ListingAdd)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MarketplaceERC1155ListingCancelledIterator is returned from FilterERC1155ListingCancelled and is used to iterate over the raw logs and unpacked data for ERC1155ListingCancelled events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingCancelledIterator struct {
	Event *ERC1155MarketplaceERC1155ListingCancelled // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceERC1155ListingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceERC1155ListingCancelled)
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
		it.Event = new(ERC1155MarketplaceERC1155ListingCancelled)
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
func (it *ERC1155MarketplaceERC1155ListingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceERC1155ListingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceERC1155ListingCancelled represents a ERC1155ListingCancelled event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingCancelled struct {
	ListingId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155ListingCancelled is a free log retrieval operation binding the contract event 0xfcaf6dd2bd3a0f4e96727088cce0c84773cc4445af7e6625fb74db10fdbf56bf.
//
// Solidity: event ERC1155ListingCancelled(uint256 indexed listingId)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterERC1155ListingCancelled(opts *bind.FilterOpts, listingId []*big.Int) (*ERC1155MarketplaceERC1155ListingCancelledIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ERC1155ListingCancelled", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceERC1155ListingCancelledIterator{contract: _ERC1155Marketplace.contract, event: "ERC1155ListingCancelled", logs: logs, sub: sub}, nil
}

// WatchERC1155ListingCancelled is a free log subscription operation binding the contract event 0xfcaf6dd2bd3a0f4e96727088cce0c84773cc4445af7e6625fb74db10fdbf56bf.
//
// Solidity: event ERC1155ListingCancelled(uint256 indexed listingId)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchERC1155ListingCancelled(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceERC1155ListingCancelled, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ERC1155ListingCancelled", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceERC1155ListingCancelled)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingCancelled", log); err != nil {
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

// ParseERC1155ListingCancelled is a log parse operation binding the contract event 0xfcaf6dd2bd3a0f4e96727088cce0c84773cc4445af7e6625fb74db10fdbf56bf.
//
// Solidity: event ERC1155ListingCancelled(uint256 indexed listingId)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseERC1155ListingCancelled(log types.Log) (*ERC1155MarketplaceERC1155ListingCancelled, error) {
	event := new(ERC1155MarketplaceERC1155ListingCancelled)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1155MarketplaceERC1155ListingSplitIterator is returned from FilterERC1155ListingSplit and is used to iterate over the raw logs and unpacked data for ERC1155ListingSplit events raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingSplitIterator struct {
	Event *ERC1155MarketplaceERC1155ListingSplit // Event containing the contract specifics and raw log

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
func (it *ERC1155MarketplaceERC1155ListingSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1155MarketplaceERC1155ListingSplit)
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
		it.Event = new(ERC1155MarketplaceERC1155ListingSplit)
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
func (it *ERC1155MarketplaceERC1155ListingSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1155MarketplaceERC1155ListingSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1155MarketplaceERC1155ListingSplit represents a ERC1155ListingSplit event raised by the ERC1155Marketplace contract.
type ERC1155MarketplaceERC1155ListingSplit struct {
	ListingId      *big.Int
	PrincipalSplit [2]uint16
	Affiliate      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterERC1155ListingSplit is a free log retrieval operation binding the contract event 0xabfb70fdd014e075c73fed3aa8dea5b5d2a0ee1a60595b9d0b3e4a58b1d8a2e0.
//
// Solidity: event ERC1155ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) FilterERC1155ListingSplit(opts *bind.FilterOpts, listingId []*big.Int) (*ERC1155MarketplaceERC1155ListingSplitIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.FilterLogs(opts, "ERC1155ListingSplit", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC1155MarketplaceERC1155ListingSplitIterator{contract: _ERC1155Marketplace.contract, event: "ERC1155ListingSplit", logs: logs, sub: sub}, nil
}

// WatchERC1155ListingSplit is a free log subscription operation binding the contract event 0xabfb70fdd014e075c73fed3aa8dea5b5d2a0ee1a60595b9d0b3e4a58b1d8a2e0.
//
// Solidity: event ERC1155ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) WatchERC1155ListingSplit(opts *bind.WatchOpts, sink chan<- *ERC1155MarketplaceERC1155ListingSplit, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC1155Marketplace.contract.WatchLogs(opts, "ERC1155ListingSplit", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1155MarketplaceERC1155ListingSplit)
				if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingSplit", log); err != nil {
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

// ParseERC1155ListingSplit is a log parse operation binding the contract event 0xabfb70fdd014e075c73fed3aa8dea5b5d2a0ee1a60595b9d0b3e4a58b1d8a2e0.
//
// Solidity: event ERC1155ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC1155Marketplace *ERC1155MarketplaceFilterer) ParseERC1155ListingSplit(log types.Log) (*ERC1155MarketplaceERC1155ListingSplit, error) {
	event := new(ERC1155MarketplaceERC1155ListingSplit)
	if err := _ERC1155Marketplace.contract.UnpackLog(event, "ERC1155ListingSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
