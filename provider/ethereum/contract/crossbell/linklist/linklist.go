// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package linklist

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

// DataTypesNoteStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNoteStruct struct {
	ProfileId *big.Int
	NoteId    *big.Int
}

// DataTypesProfileLinkStruct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesProfileLinkStruct struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
}

// LinkListMetaData contains all meta data concerning the LinkList contract.
var LinkListMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"approved\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"Uri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"Web3Entry\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addLinkingAddress\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"ethAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"addLinkingAnyUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"toUri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"addLinkingERC721\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addLinkingLinklistId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"addLinkingNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toNoteId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addLinkingProfileId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"addLinkingProfileLink\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"linkData\",\"internalType\":\"structDataTypes.ProfileLinkStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"approve\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"balanceOf\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"burn\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"getApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}],\"name\":\"getCurrentTakeOver\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"\",\"internalType\":\"bytes32\"}],\"name\":\"getLinkType\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingAddressListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address[]\",\"name\":\"\",\"internalType\":\"address[]\"}],\"name\":\"getLinkingAddresses\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingAnyListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"getLinkingAnyUri\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes32[]\",\"name\":\"\",\"internalType\":\"bytes32[]\"}],\"name\":\"getLinkingAnyUriKeys\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string[]\",\"name\":\"results\",\"internalType\":\"string[]\"}],\"name\":\"getLinkingAnyUris\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.ERC721Struct\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingERC721\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingERC721ListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.ERC721Struct[]\",\"components\":[{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingERC721s\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"}],\"name\":\"getLinkingLinklistIds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingLinklistLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.NoteStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingNote\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingNoteListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.NoteStruct[]\",\"components\":[{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"noteId\",\"internalType\":\"uint256\"}]}],\"name\":\"getLinkingNotes\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256[]\",\"name\":\"\",\"internalType\":\"uint256[]\"}],\"name\":\"getLinkingProfileIds\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple\",\"name\":\"\",\"internalType\":\"structDataTypes.ProfileLinkStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}],\"name\":\"getLinkingProfileLink\",\"inputs\":[{\"type\":\"bytes32\",\"name\":\"linkKey\",\"internalType\":\"bytes32\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingProfileLinkListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"tuple[]\",\"name\":\"results\",\"internalType\":\"structDataTypes.ProfileLinkStruct[]\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}],\"name\":\"getLinkingProfileLinks\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"getLinkingProfileListLength\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"initialize\",\"inputs\":[{\"type\":\"string\",\"name\":\"_name\",\"internalType\":\"string\"},{\"type\":\"string\",\"name\":\"_symbol\",\"internalType\":\"string\"},{\"type\":\"address\",\"name\":\"_web3Entry\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isApprovedForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"mint\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"name\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"ownerOf\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingAddress\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"ethAddress\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingAnyUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"toUri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingERC721\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"tokenAddress\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"erc721TokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingLinklistId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"linklistId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingNote\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toNoteId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingProfileId\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"removeLinkingProfileLink\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"linkData\",\"internalType\":\"structDataTypes.ProfileLinkStruct\",\"components\":[{\"type\":\"uint256\",\"name\":\"fromProfileId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"toProfileId\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"linkType\",\"internalType\":\"bytes32\"}]}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"safeTransferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"_data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setApprovalForAll\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"approved\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setTakeOver\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"profileId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"setUri\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"string\",\"name\":\"_uri\",\"internalType\":\"string\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"symbol\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenByIndex\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"inputs\":[{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"index\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"string\",\"name\":\"\",\"internalType\":\"string\"}],\"name\":\"tokenURI\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"totalSupply\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferFrom\",\"inputs\":[{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]}]",
}

// LinkListABI is the input ABI used to generate the binding from.
// Deprecated: Use LinkListMetaData.ABI instead.
var LinkListABI = LinkListMetaData.ABI

// LinkList is an auto generated Go binding around an Ethereum contract.
type LinkList struct {
	LinkListCaller     // Read-only binding to the contract
	LinkListTransactor // Write-only binding to the contract
	LinkListFilterer   // Log filterer for contract events
}

// LinkListCaller is an auto generated read-only Go binding around an Ethereum contract.
type LinkListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LinkListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LinkListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LinkListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LinkListSession struct {
	Contract     *LinkList         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LinkListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LinkListCallerSession struct {
	Contract *LinkListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LinkListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LinkListTransactorSession struct {
	Contract     *LinkListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LinkListRaw is an auto generated low-level Go binding around an Ethereum contract.
type LinkListRaw struct {
	Contract *LinkList // Generic contract binding to access the raw methods on
}

// LinkListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LinkListCallerRaw struct {
	Contract *LinkListCaller // Generic read-only contract binding to access the raw methods on
}

// LinkListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LinkListTransactorRaw struct {
	Contract *LinkListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLinkList creates a new instance of LinkList, bound to a specific deployed contract.
func NewLinkList(address common.Address, backend bind.ContractBackend) (*LinkList, error) {
	contract, err := bindLinkList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LinkList{LinkListCaller: LinkListCaller{contract: contract}, LinkListTransactor: LinkListTransactor{contract: contract}, LinkListFilterer: LinkListFilterer{contract: contract}}, nil
}

// NewLinkListCaller creates a new read-only instance of LinkList, bound to a specific deployed contract.
func NewLinkListCaller(address common.Address, caller bind.ContractCaller) (*LinkListCaller, error) {
	contract, err := bindLinkList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LinkListCaller{contract: contract}, nil
}

// NewLinkListTransactor creates a new write-only instance of LinkList, bound to a specific deployed contract.
func NewLinkListTransactor(address common.Address, transactor bind.ContractTransactor) (*LinkListTransactor, error) {
	contract, err := bindLinkList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LinkListTransactor{contract: contract}, nil
}

// NewLinkListFilterer creates a new log filterer instance of LinkList, bound to a specific deployed contract.
func NewLinkListFilterer(address common.Address, filterer bind.ContractFilterer) (*LinkListFilterer, error) {
	contract, err := bindLinkList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LinkListFilterer{contract: contract}, nil
}

// bindLinkList binds a generic wrapper to an already deployed contract.
func bindLinkList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LinkListMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkList *LinkListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkList.Contract.LinkListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkList *LinkListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkList.Contract.LinkListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkList *LinkListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkList.Contract.LinkListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LinkList *LinkListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LinkList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LinkList *LinkListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LinkList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LinkList *LinkListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LinkList.Contract.contract.Transact(opts, method, params...)
}

// Uri is a free data retrieval call binding the contract method 0xac150a54.
//
// Solidity: function Uri(uint256 tokenId) view returns(string)
func (_LinkList *LinkListCaller) Uri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "Uri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0xac150a54.
//
// Solidity: function Uri(uint256 tokenId) view returns(string)
func (_LinkList *LinkListSession) Uri(tokenId *big.Int) (string, error) {
	return _LinkList.Contract.Uri(&_LinkList.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0xac150a54.
//
// Solidity: function Uri(uint256 tokenId) view returns(string)
func (_LinkList *LinkListCallerSession) Uri(tokenId *big.Int) (string, error) {
	return _LinkList.Contract.Uri(&_LinkList.CallOpts, tokenId)
}

// Web3Entry is a free data retrieval call binding the contract method 0xc5a5ed51.
//
// Solidity: function Web3Entry() view returns(address)
func (_LinkList *LinkListCaller) Web3Entry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "Web3Entry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Web3Entry is a free data retrieval call binding the contract method 0xc5a5ed51.
//
// Solidity: function Web3Entry() view returns(address)
func (_LinkList *LinkListSession) Web3Entry() (common.Address, error) {
	return _LinkList.Contract.Web3Entry(&_LinkList.CallOpts)
}

// Web3Entry is a free data retrieval call binding the contract method 0xc5a5ed51.
//
// Solidity: function Web3Entry() view returns(address)
func (_LinkList *LinkListCallerSession) Web3Entry() (common.Address, error) {
	return _LinkList.Contract.Web3Entry(&_LinkList.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LinkList *LinkListCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LinkList *LinkListSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkList.Contract.BalanceOf(&_LinkList.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LinkList *LinkListCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LinkList.Contract.BalanceOf(&_LinkList.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LinkList *LinkListCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LinkList *LinkListSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LinkList.Contract.GetApproved(&_LinkList.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LinkList *LinkListCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LinkList.Contract.GetApproved(&_LinkList.CallOpts, tokenId)
}

// GetCurrentTakeOver is a free data retrieval call binding the contract method 0x493fa4dc.
//
// Solidity: function getCurrentTakeOver(uint256 tokenId) view returns(uint256 profileId)
func (_LinkList *LinkListCaller) GetCurrentTakeOver(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getCurrentTakeOver", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTakeOver is a free data retrieval call binding the contract method 0x493fa4dc.
//
// Solidity: function getCurrentTakeOver(uint256 tokenId) view returns(uint256 profileId)
func (_LinkList *LinkListSession) GetCurrentTakeOver(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetCurrentTakeOver(&_LinkList.CallOpts, tokenId)
}

// GetCurrentTakeOver is a free data retrieval call binding the contract method 0x493fa4dc.
//
// Solidity: function getCurrentTakeOver(uint256 tokenId) view returns(uint256 profileId)
func (_LinkList *LinkListCallerSession) GetCurrentTakeOver(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetCurrentTakeOver(&_LinkList.CallOpts, tokenId)
}

// GetLinkType is a free data retrieval call binding the contract method 0x00fba027.
//
// Solidity: function getLinkType(uint256 tokenId) view returns(bytes32)
func (_LinkList *LinkListCaller) GetLinkType(opts *bind.CallOpts, tokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkType", tokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinkType is a free data retrieval call binding the contract method 0x00fba027.
//
// Solidity: function getLinkType(uint256 tokenId) view returns(bytes32)
func (_LinkList *LinkListSession) GetLinkType(tokenId *big.Int) ([32]byte, error) {
	return _LinkList.Contract.GetLinkType(&_LinkList.CallOpts, tokenId)
}

// GetLinkType is a free data retrieval call binding the contract method 0x00fba027.
//
// Solidity: function getLinkType(uint256 tokenId) view returns(bytes32)
func (_LinkList *LinkListCallerSession) GetLinkType(tokenId *big.Int) ([32]byte, error) {
	return _LinkList.Contract.GetLinkType(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAddressListLength is a free data retrieval call binding the contract method 0x97ce7df6.
//
// Solidity: function getLinkingAddressListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingAddressListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAddressListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingAddressListLength is a free data retrieval call binding the contract method 0x97ce7df6.
//
// Solidity: function getLinkingAddressListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingAddressListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingAddressListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAddressListLength is a free data retrieval call binding the contract method 0x97ce7df6.
//
// Solidity: function getLinkingAddressListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingAddressListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingAddressListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x5c369ec3.
//
// Solidity: function getLinkingAddresses(uint256 tokenId) view returns(address[])
func (_LinkList *LinkListCaller) GetLinkingAddresses(opts *bind.CallOpts, tokenId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAddresses", tokenId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x5c369ec3.
//
// Solidity: function getLinkingAddresses(uint256 tokenId) view returns(address[])
func (_LinkList *LinkListSession) GetLinkingAddresses(tokenId *big.Int) ([]common.Address, error) {
	return _LinkList.Contract.GetLinkingAddresses(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAddresses is a free data retrieval call binding the contract method 0x5c369ec3.
//
// Solidity: function getLinkingAddresses(uint256 tokenId) view returns(address[])
func (_LinkList *LinkListCallerSession) GetLinkingAddresses(tokenId *big.Int) ([]common.Address, error) {
	return _LinkList.Contract.GetLinkingAddresses(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyListLength is a free data retrieval call binding the contract method 0x4aa7b87a.
//
// Solidity: function getLinkingAnyListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingAnyListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAnyListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingAnyListLength is a free data retrieval call binding the contract method 0x4aa7b87a.
//
// Solidity: function getLinkingAnyListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingAnyListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingAnyListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyListLength is a free data retrieval call binding the contract method 0x4aa7b87a.
//
// Solidity: function getLinkingAnyListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingAnyListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingAnyListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_LinkList *LinkListCaller) GetLinkingAnyUri(opts *bind.CallOpts, linkKey [32]byte) (string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAnyUri", linkKey)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_LinkList *LinkListSession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _LinkList.Contract.GetLinkingAnyUri(&_LinkList.CallOpts, linkKey)
}

// GetLinkingAnyUri is a free data retrieval call binding the contract method 0x5f799cc6.
//
// Solidity: function getLinkingAnyUri(bytes32 linkKey) view returns(string)
func (_LinkList *LinkListCallerSession) GetLinkingAnyUri(linkKey [32]byte) (string, error) {
	return _LinkList.Contract.GetLinkingAnyUri(&_LinkList.CallOpts, linkKey)
}

// GetLinkingAnyUriKeys is a free data retrieval call binding the contract method 0x76bac025.
//
// Solidity: function getLinkingAnyUriKeys(uint256 tokenId) view returns(bytes32[])
func (_LinkList *LinkListCaller) GetLinkingAnyUriKeys(opts *bind.CallOpts, tokenId *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAnyUriKeys", tokenId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetLinkingAnyUriKeys is a free data retrieval call binding the contract method 0x76bac025.
//
// Solidity: function getLinkingAnyUriKeys(uint256 tokenId) view returns(bytes32[])
func (_LinkList *LinkListSession) GetLinkingAnyUriKeys(tokenId *big.Int) ([][32]byte, error) {
	return _LinkList.Contract.GetLinkingAnyUriKeys(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyUriKeys is a free data retrieval call binding the contract method 0x76bac025.
//
// Solidity: function getLinkingAnyUriKeys(uint256 tokenId) view returns(bytes32[])
func (_LinkList *LinkListCallerSession) GetLinkingAnyUriKeys(tokenId *big.Int) ([][32]byte, error) {
	return _LinkList.Contract.GetLinkingAnyUriKeys(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0xe64f2baa.
//
// Solidity: function getLinkingAnyUris(uint256 tokenId) view returns(string[] results)
func (_LinkList *LinkListCaller) GetLinkingAnyUris(opts *bind.CallOpts, tokenId *big.Int) ([]string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingAnyUris", tokenId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0xe64f2baa.
//
// Solidity: function getLinkingAnyUris(uint256 tokenId) view returns(string[] results)
func (_LinkList *LinkListSession) GetLinkingAnyUris(tokenId *big.Int) ([]string, error) {
	return _LinkList.Contract.GetLinkingAnyUris(&_LinkList.CallOpts, tokenId)
}

// GetLinkingAnyUris is a free data retrieval call binding the contract method 0xe64f2baa.
//
// Solidity: function getLinkingAnyUris(uint256 tokenId) view returns(string[] results)
func (_LinkList *LinkListCallerSession) GetLinkingAnyUris(tokenId *big.Int) ([]string, error) {
	return _LinkList.Contract.GetLinkingAnyUris(&_LinkList.CallOpts, tokenId)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_LinkList *LinkListCaller) GetLinkingERC721(opts *bind.CallOpts, linkKey [32]byte) (DataTypesERC721Struct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingERC721", linkKey)

	if err != nil {
		return *new(DataTypesERC721Struct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesERC721Struct)).(*DataTypesERC721Struct)

	return out0, err

}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_LinkList *LinkListSession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _LinkList.Contract.GetLinkingERC721(&_LinkList.CallOpts, linkKey)
}

// GetLinkingERC721 is a free data retrieval call binding the contract method 0x5440522b.
//
// Solidity: function getLinkingERC721(bytes32 linkKey) view returns((address,uint256))
func (_LinkList *LinkListCallerSession) GetLinkingERC721(linkKey [32]byte) (DataTypesERC721Struct, error) {
	return _LinkList.Contract.GetLinkingERC721(&_LinkList.CallOpts, linkKey)
}

// GetLinkingERC721ListLength is a free data retrieval call binding the contract method 0x6d37d274.
//
// Solidity: function getLinkingERC721ListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingERC721ListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingERC721ListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingERC721ListLength is a free data retrieval call binding the contract method 0x6d37d274.
//
// Solidity: function getLinkingERC721ListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingERC721ListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingERC721ListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingERC721ListLength is a free data retrieval call binding the contract method 0x6d37d274.
//
// Solidity: function getLinkingERC721ListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingERC721ListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingERC721ListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x9037c0f9.
//
// Solidity: function getLinkingERC721s(uint256 tokenId) view returns((address,uint256)[] results)
func (_LinkList *LinkListCaller) GetLinkingERC721s(opts *bind.CallOpts, tokenId *big.Int) ([]DataTypesERC721Struct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingERC721s", tokenId)

	if err != nil {
		return *new([]DataTypesERC721Struct), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesERC721Struct)).(*[]DataTypesERC721Struct)

	return out0, err

}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x9037c0f9.
//
// Solidity: function getLinkingERC721s(uint256 tokenId) view returns((address,uint256)[] results)
func (_LinkList *LinkListSession) GetLinkingERC721s(tokenId *big.Int) ([]DataTypesERC721Struct, error) {
	return _LinkList.Contract.GetLinkingERC721s(&_LinkList.CallOpts, tokenId)
}

// GetLinkingERC721s is a free data retrieval call binding the contract method 0x9037c0f9.
//
// Solidity: function getLinkingERC721s(uint256 tokenId) view returns((address,uint256)[] results)
func (_LinkList *LinkListCallerSession) GetLinkingERC721s(tokenId *big.Int) ([]DataTypesERC721Struct, error) {
	return _LinkList.Contract.GetLinkingERC721s(&_LinkList.CallOpts, tokenId)
}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0xed6223cf.
//
// Solidity: function getLinkingLinklistIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListCaller) GetLinkingLinklistIds(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingLinklistIds", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0xed6223cf.
//
// Solidity: function getLinkingLinklistIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListSession) GetLinkingLinklistIds(tokenId *big.Int) ([]*big.Int, error) {
	return _LinkList.Contract.GetLinkingLinklistIds(&_LinkList.CallOpts, tokenId)
}

// GetLinkingLinklistIds is a free data retrieval call binding the contract method 0xed6223cf.
//
// Solidity: function getLinkingLinklistIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListCallerSession) GetLinkingLinklistIds(tokenId *big.Int) ([]*big.Int, error) {
	return _LinkList.Contract.GetLinkingLinklistIds(&_LinkList.CallOpts, tokenId)
}

// GetLinkingLinklistLength is a free data retrieval call binding the contract method 0x8fd6dc39.
//
// Solidity: function getLinkingLinklistLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingLinklistLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingLinklistLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingLinklistLength is a free data retrieval call binding the contract method 0x8fd6dc39.
//
// Solidity: function getLinkingLinklistLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingLinklistLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingLinklistLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingLinklistLength is a free data retrieval call binding the contract method 0x8fd6dc39.
//
// Solidity: function getLinkingLinklistLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingLinklistLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingLinklistLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_LinkList *LinkListCaller) GetLinkingNote(opts *bind.CallOpts, linkKey [32]byte) (DataTypesNoteStruct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingNote", linkKey)

	if err != nil {
		return *new(DataTypesNoteStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNoteStruct)).(*DataTypesNoteStruct)

	return out0, err

}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_LinkList *LinkListSession) GetLinkingNote(linkKey [32]byte) (DataTypesNoteStruct, error) {
	return _LinkList.Contract.GetLinkingNote(&_LinkList.CallOpts, linkKey)
}

// GetLinkingNote is a free data retrieval call binding the contract method 0x4a7906b9.
//
// Solidity: function getLinkingNote(bytes32 linkKey) view returns((uint256,uint256))
func (_LinkList *LinkListCallerSession) GetLinkingNote(linkKey [32]byte) (DataTypesNoteStruct, error) {
	return _LinkList.Contract.GetLinkingNote(&_LinkList.CallOpts, linkKey)
}

// GetLinkingNoteListLength is a free data retrieval call binding the contract method 0xf61087f6.
//
// Solidity: function getLinkingNoteListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingNoteListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingNoteListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingNoteListLength is a free data retrieval call binding the contract method 0xf61087f6.
//
// Solidity: function getLinkingNoteListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingNoteListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingNoteListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingNoteListLength is a free data retrieval call binding the contract method 0xf61087f6.
//
// Solidity: function getLinkingNoteListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingNoteListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingNoteListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x5e9f678b.
//
// Solidity: function getLinkingNotes(uint256 tokenId) view returns((uint256,uint256)[] results)
func (_LinkList *LinkListCaller) GetLinkingNotes(opts *bind.CallOpts, tokenId *big.Int) ([]DataTypesNoteStruct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingNotes", tokenId)

	if err != nil {
		return *new([]DataTypesNoteStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesNoteStruct)).(*[]DataTypesNoteStruct)

	return out0, err

}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x5e9f678b.
//
// Solidity: function getLinkingNotes(uint256 tokenId) view returns((uint256,uint256)[] results)
func (_LinkList *LinkListSession) GetLinkingNotes(tokenId *big.Int) ([]DataTypesNoteStruct, error) {
	return _LinkList.Contract.GetLinkingNotes(&_LinkList.CallOpts, tokenId)
}

// GetLinkingNotes is a free data retrieval call binding the contract method 0x5e9f678b.
//
// Solidity: function getLinkingNotes(uint256 tokenId) view returns((uint256,uint256)[] results)
func (_LinkList *LinkListCallerSession) GetLinkingNotes(tokenId *big.Int) ([]DataTypesNoteStruct, error) {
	return _LinkList.Contract.GetLinkingNotes(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileIds is a free data retrieval call binding the contract method 0xa72d2317.
//
// Solidity: function getLinkingProfileIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListCaller) GetLinkingProfileIds(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingProfileIds", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLinkingProfileIds is a free data retrieval call binding the contract method 0xa72d2317.
//
// Solidity: function getLinkingProfileIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListSession) GetLinkingProfileIds(tokenId *big.Int) ([]*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileIds(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileIds is a free data retrieval call binding the contract method 0xa72d2317.
//
// Solidity: function getLinkingProfileIds(uint256 tokenId) view returns(uint256[])
func (_LinkList *LinkListCallerSession) GetLinkingProfileIds(tokenId *big.Int) ([]*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileIds(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileLink is a free data retrieval call binding the contract method 0xa9ec321f.
//
// Solidity: function getLinkingProfileLink(bytes32 linkKey) view returns((uint256,uint256,bytes32))
func (_LinkList *LinkListCaller) GetLinkingProfileLink(opts *bind.CallOpts, linkKey [32]byte) (DataTypesProfileLinkStruct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingProfileLink", linkKey)

	if err != nil {
		return *new(DataTypesProfileLinkStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfileLinkStruct)).(*DataTypesProfileLinkStruct)

	return out0, err

}

// GetLinkingProfileLink is a free data retrieval call binding the contract method 0xa9ec321f.
//
// Solidity: function getLinkingProfileLink(bytes32 linkKey) view returns((uint256,uint256,bytes32))
func (_LinkList *LinkListSession) GetLinkingProfileLink(linkKey [32]byte) (DataTypesProfileLinkStruct, error) {
	return _LinkList.Contract.GetLinkingProfileLink(&_LinkList.CallOpts, linkKey)
}

// GetLinkingProfileLink is a free data retrieval call binding the contract method 0xa9ec321f.
//
// Solidity: function getLinkingProfileLink(bytes32 linkKey) view returns((uint256,uint256,bytes32))
func (_LinkList *LinkListCallerSession) GetLinkingProfileLink(linkKey [32]byte) (DataTypesProfileLinkStruct, error) {
	return _LinkList.Contract.GetLinkingProfileLink(&_LinkList.CallOpts, linkKey)
}

// GetLinkingProfileLinkListLength is a free data retrieval call binding the contract method 0xb34bcb96.
//
// Solidity: function getLinkingProfileLinkListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingProfileLinkListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingProfileLinkListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingProfileLinkListLength is a free data retrieval call binding the contract method 0xb34bcb96.
//
// Solidity: function getLinkingProfileLinkListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingProfileLinkListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileLinkListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileLinkListLength is a free data retrieval call binding the contract method 0xb34bcb96.
//
// Solidity: function getLinkingProfileLinkListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingProfileLinkListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileLinkListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileLinks is a free data retrieval call binding the contract method 0x0b341885.
//
// Solidity: function getLinkingProfileLinks(uint256 tokenId) view returns((uint256,uint256,bytes32)[] results)
func (_LinkList *LinkListCaller) GetLinkingProfileLinks(opts *bind.CallOpts, tokenId *big.Int) ([]DataTypesProfileLinkStruct, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingProfileLinks", tokenId)

	if err != nil {
		return *new([]DataTypesProfileLinkStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesProfileLinkStruct)).(*[]DataTypesProfileLinkStruct)

	return out0, err

}

// GetLinkingProfileLinks is a free data retrieval call binding the contract method 0x0b341885.
//
// Solidity: function getLinkingProfileLinks(uint256 tokenId) view returns((uint256,uint256,bytes32)[] results)
func (_LinkList *LinkListSession) GetLinkingProfileLinks(tokenId *big.Int) ([]DataTypesProfileLinkStruct, error) {
	return _LinkList.Contract.GetLinkingProfileLinks(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileLinks is a free data retrieval call binding the contract method 0x0b341885.
//
// Solidity: function getLinkingProfileLinks(uint256 tokenId) view returns((uint256,uint256,bytes32)[] results)
func (_LinkList *LinkListCallerSession) GetLinkingProfileLinks(tokenId *big.Int) ([]DataTypesProfileLinkStruct, error) {
	return _LinkList.Contract.GetLinkingProfileLinks(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileListLength is a free data retrieval call binding the contract method 0x6d46df26.
//
// Solidity: function getLinkingProfileListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCaller) GetLinkingProfileListLength(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "getLinkingProfileListLength", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinkingProfileListLength is a free data retrieval call binding the contract method 0x6d46df26.
//
// Solidity: function getLinkingProfileListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListSession) GetLinkingProfileListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileListLength(&_LinkList.CallOpts, tokenId)
}

// GetLinkingProfileListLength is a free data retrieval call binding the contract method 0x6d46df26.
//
// Solidity: function getLinkingProfileListLength(uint256 tokenId) view returns(uint256)
func (_LinkList *LinkListCallerSession) GetLinkingProfileListLength(tokenId *big.Int) (*big.Int, error) {
	return _LinkList.Contract.GetLinkingProfileListLength(&_LinkList.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LinkList *LinkListCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LinkList *LinkListSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LinkList.Contract.IsApprovedForAll(&_LinkList.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LinkList *LinkListCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LinkList.Contract.IsApprovedForAll(&_LinkList.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LinkList *LinkListCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LinkList *LinkListSession) Name() (string, error) {
	return _LinkList.Contract.Name(&_LinkList.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LinkList *LinkListCallerSession) Name() (string, error) {
	return _LinkList.Contract.Name(&_LinkList.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LinkList *LinkListCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LinkList *LinkListSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LinkList.Contract.OwnerOf(&_LinkList.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LinkList *LinkListCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LinkList.Contract.OwnerOf(&_LinkList.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LinkList *LinkListCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LinkList *LinkListSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LinkList.Contract.SupportsInterface(&_LinkList.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LinkList *LinkListCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LinkList.Contract.SupportsInterface(&_LinkList.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LinkList *LinkListCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LinkList *LinkListSession) Symbol() (string, error) {
	return _LinkList.Contract.Symbol(&_LinkList.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LinkList *LinkListCallerSession) Symbol() (string, error) {
	return _LinkList.Contract.Symbol(&_LinkList.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LinkList *LinkListCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LinkList *LinkListSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LinkList.Contract.TokenByIndex(&_LinkList.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_LinkList *LinkListCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _LinkList.Contract.TokenByIndex(&_LinkList.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LinkList *LinkListCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LinkList *LinkListSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LinkList.Contract.TokenOfOwnerByIndex(&_LinkList.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_LinkList *LinkListCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _LinkList.Contract.TokenOfOwnerByIndex(&_LinkList.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LinkList *LinkListCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LinkList *LinkListSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LinkList.Contract.TokenURI(&_LinkList.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LinkList *LinkListCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LinkList.Contract.TokenURI(&_LinkList.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LinkList *LinkListCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LinkList.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LinkList *LinkListSession) TotalSupply() (*big.Int, error) {
	return _LinkList.Contract.TotalSupply(&_LinkList.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LinkList *LinkListCallerSession) TotalSupply() (*big.Int, error) {
	return _LinkList.Contract.TotalSupply(&_LinkList.CallOpts)
}

// AddLinkingAddress is a paid mutator transaction binding the contract method 0xed386e65.
//
// Solidity: function addLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListTransactor) AddLinkingAddress(opts *bind.TransactOpts, tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingAddress", tokenId, ethAddress)
}

// AddLinkingAddress is a paid mutator transaction binding the contract method 0xed386e65.
//
// Solidity: function addLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListSession) AddLinkingAddress(tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingAddress(&_LinkList.TransactOpts, tokenId, ethAddress)
}

// AddLinkingAddress is a paid mutator transaction binding the contract method 0xed386e65.
//
// Solidity: function addLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListTransactorSession) AddLinkingAddress(tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingAddress(&_LinkList.TransactOpts, tokenId, ethAddress)
}

// AddLinkingAnyUri is a paid mutator transaction binding the contract method 0x6c217e12.
//
// Solidity: function addLinkingAnyUri(uint256 tokenId, string toUri) returns(bytes32)
func (_LinkList *LinkListTransactor) AddLinkingAnyUri(opts *bind.TransactOpts, tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingAnyUri", tokenId, toUri)
}

// AddLinkingAnyUri is a paid mutator transaction binding the contract method 0x6c217e12.
//
// Solidity: function addLinkingAnyUri(uint256 tokenId, string toUri) returns(bytes32)
func (_LinkList *LinkListSession) AddLinkingAnyUri(tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingAnyUri(&_LinkList.TransactOpts, tokenId, toUri)
}

// AddLinkingAnyUri is a paid mutator transaction binding the contract method 0x6c217e12.
//
// Solidity: function addLinkingAnyUri(uint256 tokenId, string toUri) returns(bytes32)
func (_LinkList *LinkListTransactorSession) AddLinkingAnyUri(tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingAnyUri(&_LinkList.TransactOpts, tokenId, toUri)
}

// AddLinkingERC721 is a paid mutator transaction binding the contract method 0x2ea24efc.
//
// Solidity: function addLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns(bytes32)
func (_LinkList *LinkListTransactor) AddLinkingERC721(opts *bind.TransactOpts, tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingERC721", tokenId, tokenAddress, erc721TokenId)
}

// AddLinkingERC721 is a paid mutator transaction binding the contract method 0x2ea24efc.
//
// Solidity: function addLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns(bytes32)
func (_LinkList *LinkListSession) AddLinkingERC721(tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingERC721(&_LinkList.TransactOpts, tokenId, tokenAddress, erc721TokenId)
}

// AddLinkingERC721 is a paid mutator transaction binding the contract method 0x2ea24efc.
//
// Solidity: function addLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns(bytes32)
func (_LinkList *LinkListTransactorSession) AddLinkingERC721(tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingERC721(&_LinkList.TransactOpts, tokenId, tokenAddress, erc721TokenId)
}

// AddLinkingLinklistId is a paid mutator transaction binding the contract method 0xc1cd4f82.
//
// Solidity: function addLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListTransactor) AddLinkingLinklistId(opts *bind.TransactOpts, tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingLinklistId", tokenId, linklistId)
}

// AddLinkingLinklistId is a paid mutator transaction binding the contract method 0xc1cd4f82.
//
// Solidity: function addLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListSession) AddLinkingLinklistId(tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingLinklistId(&_LinkList.TransactOpts, tokenId, linklistId)
}

// AddLinkingLinklistId is a paid mutator transaction binding the contract method 0xc1cd4f82.
//
// Solidity: function addLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListTransactorSession) AddLinkingLinklistId(tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingLinklistId(&_LinkList.TransactOpts, tokenId, linklistId)
}

// AddLinkingNote is a paid mutator transaction binding the contract method 0x5cb46be7.
//
// Solidity: function addLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns(bytes32)
func (_LinkList *LinkListTransactor) AddLinkingNote(opts *bind.TransactOpts, tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingNote", tokenId, toProfileId, toNoteId)
}

// AddLinkingNote is a paid mutator transaction binding the contract method 0x5cb46be7.
//
// Solidity: function addLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns(bytes32)
func (_LinkList *LinkListSession) AddLinkingNote(tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingNote(&_LinkList.TransactOpts, tokenId, toProfileId, toNoteId)
}

// AddLinkingNote is a paid mutator transaction binding the contract method 0x5cb46be7.
//
// Solidity: function addLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns(bytes32)
func (_LinkList *LinkListTransactorSession) AddLinkingNote(tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingNote(&_LinkList.TransactOpts, tokenId, toProfileId, toNoteId)
}

// AddLinkingProfileId is a paid mutator transaction binding the contract method 0xf7e1711a.
//
// Solidity: function addLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListTransactor) AddLinkingProfileId(opts *bind.TransactOpts, tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingProfileId", tokenId, toProfileId)
}

// AddLinkingProfileId is a paid mutator transaction binding the contract method 0xf7e1711a.
//
// Solidity: function addLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListSession) AddLinkingProfileId(tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingProfileId(&_LinkList.TransactOpts, tokenId, toProfileId)
}

// AddLinkingProfileId is a paid mutator transaction binding the contract method 0xf7e1711a.
//
// Solidity: function addLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListTransactorSession) AddLinkingProfileId(tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingProfileId(&_LinkList.TransactOpts, tokenId, toProfileId)
}

// AddLinkingProfileLink is a paid mutator transaction binding the contract method 0x494c4d99.
//
// Solidity: function addLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListTransactor) AddLinkingProfileLink(opts *bind.TransactOpts, tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "addLinkingProfileLink", tokenId, linkData)
}

// AddLinkingProfileLink is a paid mutator transaction binding the contract method 0x494c4d99.
//
// Solidity: function addLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListSession) AddLinkingProfileLink(tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingProfileLink(&_LinkList.TransactOpts, tokenId, linkData)
}

// AddLinkingProfileLink is a paid mutator transaction binding the contract method 0x494c4d99.
//
// Solidity: function addLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListTransactorSession) AddLinkingProfileLink(tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.Contract.AddLinkingProfileLink(&_LinkList.TransactOpts, tokenId, linkData)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LinkList *LinkListSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Approve(&_LinkList.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Approve(&_LinkList.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LinkList *LinkListTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LinkList *LinkListSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Burn(&_LinkList.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_LinkList *LinkListTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Burn(&_LinkList.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _web3Entry) returns()
func (_LinkList *LinkListTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _web3Entry common.Address) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "initialize", _name, _symbol, _web3Entry)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _web3Entry) returns()
func (_LinkList *LinkListSession) Initialize(_name string, _symbol string, _web3Entry common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.Initialize(&_LinkList.TransactOpts, _name, _symbol, _web3Entry)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string _name, string _symbol, address _web3Entry) returns()
func (_LinkList *LinkListTransactorSession) Initialize(_name string, _symbol string, _web3Entry common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.Initialize(&_LinkList.TransactOpts, _name, _symbol, _web3Entry)
}

// Mint is a paid mutator transaction binding the contract method 0xc89dcfce.
//
// Solidity: function mint(address to, bytes32 linkType, uint256 tokenId) returns()
func (_LinkList *LinkListTransactor) Mint(opts *bind.TransactOpts, to common.Address, linkType [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "mint", to, linkType, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xc89dcfce.
//
// Solidity: function mint(address to, bytes32 linkType, uint256 tokenId) returns()
func (_LinkList *LinkListSession) Mint(to common.Address, linkType [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Mint(&_LinkList.TransactOpts, to, linkType, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0xc89dcfce.
//
// Solidity: function mint(address to, bytes32 linkType, uint256 tokenId) returns()
func (_LinkList *LinkListTransactorSession) Mint(to common.Address, linkType [32]byte, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.Mint(&_LinkList.TransactOpts, to, linkType, tokenId)
}

// RemoveLinkingAddress is a paid mutator transaction binding the contract method 0x5956da73.
//
// Solidity: function removeLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingAddress(opts *bind.TransactOpts, tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingAddress", tokenId, ethAddress)
}

// RemoveLinkingAddress is a paid mutator transaction binding the contract method 0x5956da73.
//
// Solidity: function removeLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListSession) RemoveLinkingAddress(tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingAddress(&_LinkList.TransactOpts, tokenId, ethAddress)
}

// RemoveLinkingAddress is a paid mutator transaction binding the contract method 0x5956da73.
//
// Solidity: function removeLinkingAddress(uint256 tokenId, address ethAddress) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingAddress(tokenId *big.Int, ethAddress common.Address) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingAddress(&_LinkList.TransactOpts, tokenId, ethAddress)
}

// RemoveLinkingAnyUri is a paid mutator transaction binding the contract method 0x42c1721f.
//
// Solidity: function removeLinkingAnyUri(uint256 tokenId, string toUri) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingAnyUri(opts *bind.TransactOpts, tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingAnyUri", tokenId, toUri)
}

// RemoveLinkingAnyUri is a paid mutator transaction binding the contract method 0x42c1721f.
//
// Solidity: function removeLinkingAnyUri(uint256 tokenId, string toUri) returns()
func (_LinkList *LinkListSession) RemoveLinkingAnyUri(tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingAnyUri(&_LinkList.TransactOpts, tokenId, toUri)
}

// RemoveLinkingAnyUri is a paid mutator transaction binding the contract method 0x42c1721f.
//
// Solidity: function removeLinkingAnyUri(uint256 tokenId, string toUri) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingAnyUri(tokenId *big.Int, toUri string) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingAnyUri(&_LinkList.TransactOpts, tokenId, toUri)
}

// RemoveLinkingERC721 is a paid mutator transaction binding the contract method 0x0370a161.
//
// Solidity: function removeLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingERC721(opts *bind.TransactOpts, tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingERC721", tokenId, tokenAddress, erc721TokenId)
}

// RemoveLinkingERC721 is a paid mutator transaction binding the contract method 0x0370a161.
//
// Solidity: function removeLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns()
func (_LinkList *LinkListSession) RemoveLinkingERC721(tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingERC721(&_LinkList.TransactOpts, tokenId, tokenAddress, erc721TokenId)
}

// RemoveLinkingERC721 is a paid mutator transaction binding the contract method 0x0370a161.
//
// Solidity: function removeLinkingERC721(uint256 tokenId, address tokenAddress, uint256 erc721TokenId) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingERC721(tokenId *big.Int, tokenAddress common.Address, erc721TokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingERC721(&_LinkList.TransactOpts, tokenId, tokenAddress, erc721TokenId)
}

// RemoveLinkingLinklistId is a paid mutator transaction binding the contract method 0xc70881c5.
//
// Solidity: function removeLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingLinklistId(opts *bind.TransactOpts, tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingLinklistId", tokenId, linklistId)
}

// RemoveLinkingLinklistId is a paid mutator transaction binding the contract method 0xc70881c5.
//
// Solidity: function removeLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListSession) RemoveLinkingLinklistId(tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingLinklistId(&_LinkList.TransactOpts, tokenId, linklistId)
}

// RemoveLinkingLinklistId is a paid mutator transaction binding the contract method 0xc70881c5.
//
// Solidity: function removeLinkingLinklistId(uint256 tokenId, uint256 linklistId) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingLinklistId(tokenId *big.Int, linklistId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingLinklistId(&_LinkList.TransactOpts, tokenId, linklistId)
}

// RemoveLinkingNote is a paid mutator transaction binding the contract method 0x040f7618.
//
// Solidity: function removeLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingNote(opts *bind.TransactOpts, tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingNote", tokenId, toProfileId, toNoteId)
}

// RemoveLinkingNote is a paid mutator transaction binding the contract method 0x040f7618.
//
// Solidity: function removeLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns()
func (_LinkList *LinkListSession) RemoveLinkingNote(tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingNote(&_LinkList.TransactOpts, tokenId, toProfileId, toNoteId)
}

// RemoveLinkingNote is a paid mutator transaction binding the contract method 0x040f7618.
//
// Solidity: function removeLinkingNote(uint256 tokenId, uint256 toProfileId, uint256 toNoteId) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingNote(tokenId *big.Int, toProfileId *big.Int, toNoteId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingNote(&_LinkList.TransactOpts, tokenId, toProfileId, toNoteId)
}

// RemoveLinkingProfileId is a paid mutator transaction binding the contract method 0x91e5454b.
//
// Solidity: function removeLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingProfileId(opts *bind.TransactOpts, tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingProfileId", tokenId, toProfileId)
}

// RemoveLinkingProfileId is a paid mutator transaction binding the contract method 0x91e5454b.
//
// Solidity: function removeLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListSession) RemoveLinkingProfileId(tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingProfileId(&_LinkList.TransactOpts, tokenId, toProfileId)
}

// RemoveLinkingProfileId is a paid mutator transaction binding the contract method 0x91e5454b.
//
// Solidity: function removeLinkingProfileId(uint256 tokenId, uint256 toProfileId) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingProfileId(tokenId *big.Int, toProfileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingProfileId(&_LinkList.TransactOpts, tokenId, toProfileId)
}

// RemoveLinkingProfileLink is a paid mutator transaction binding the contract method 0x70a5dba8.
//
// Solidity: function removeLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListTransactor) RemoveLinkingProfileLink(opts *bind.TransactOpts, tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "removeLinkingProfileLink", tokenId, linkData)
}

// RemoveLinkingProfileLink is a paid mutator transaction binding the contract method 0x70a5dba8.
//
// Solidity: function removeLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListSession) RemoveLinkingProfileLink(tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingProfileLink(&_LinkList.TransactOpts, tokenId, linkData)
}

// RemoveLinkingProfileLink is a paid mutator transaction binding the contract method 0x70a5dba8.
//
// Solidity: function removeLinkingProfileLink(uint256 tokenId, (uint256,uint256,bytes32) linkData) returns()
func (_LinkList *LinkListTransactorSession) RemoveLinkingProfileLink(tokenId *big.Int, linkData DataTypesProfileLinkStruct) (*types.Transaction, error) {
	return _LinkList.Contract.RemoveLinkingProfileLink(&_LinkList.TransactOpts, tokenId, linkData)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.SafeTransferFrom(&_LinkList.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.SafeTransferFrom(&_LinkList.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LinkList *LinkListTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LinkList *LinkListSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LinkList.Contract.SafeTransferFrom0(&_LinkList.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_LinkList *LinkListTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _LinkList.Contract.SafeTransferFrom0(&_LinkList.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LinkList *LinkListTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LinkList *LinkListSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LinkList.Contract.SetApprovalForAll(&_LinkList.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LinkList *LinkListTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LinkList.Contract.SetApprovalForAll(&_LinkList.TransactOpts, operator, approved)
}

// SetTakeOver is a paid mutator transaction binding the contract method 0x716c99a4.
//
// Solidity: function setTakeOver(uint256 tokenId, address to, uint256 profileId) returns()
func (_LinkList *LinkListTransactor) SetTakeOver(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, profileId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "setTakeOver", tokenId, to, profileId)
}

// SetTakeOver is a paid mutator transaction binding the contract method 0x716c99a4.
//
// Solidity: function setTakeOver(uint256 tokenId, address to, uint256 profileId) returns()
func (_LinkList *LinkListSession) SetTakeOver(tokenId *big.Int, to common.Address, profileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.SetTakeOver(&_LinkList.TransactOpts, tokenId, to, profileId)
}

// SetTakeOver is a paid mutator transaction binding the contract method 0x716c99a4.
//
// Solidity: function setTakeOver(uint256 tokenId, address to, uint256 profileId) returns()
func (_LinkList *LinkListTransactorSession) SetTakeOver(tokenId *big.Int, to common.Address, profileId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.SetTakeOver(&_LinkList.TransactOpts, tokenId, to, profileId)
}

// SetUri is a paid mutator transaction binding the contract method 0x782f08ae.
//
// Solidity: function setUri(uint256 tokenId, string _uri) returns()
func (_LinkList *LinkListTransactor) SetUri(opts *bind.TransactOpts, tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "setUri", tokenId, _uri)
}

// SetUri is a paid mutator transaction binding the contract method 0x782f08ae.
//
// Solidity: function setUri(uint256 tokenId, string _uri) returns()
func (_LinkList *LinkListSession) SetUri(tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _LinkList.Contract.SetUri(&_LinkList.TransactOpts, tokenId, _uri)
}

// SetUri is a paid mutator transaction binding the contract method 0x782f08ae.
//
// Solidity: function setUri(uint256 tokenId, string _uri) returns()
func (_LinkList *LinkListTransactorSession) SetUri(tokenId *big.Int, _uri string) (*types.Transaction, error) {
	return _LinkList.Contract.SetUri(&_LinkList.TransactOpts, tokenId, _uri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.TransferFrom(&_LinkList.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LinkList *LinkListTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LinkList.Contract.TransferFrom(&_LinkList.TransactOpts, from, to, tokenId)
}

// LinkListApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LinkList contract.
type LinkListApprovalIterator struct {
	Event *LinkListApproval // Event containing the contract specifics and raw log

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
func (it *LinkListApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkListApproval)
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
		it.Event = new(LinkListApproval)
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
func (it *LinkListApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkListApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkListApproval represents a Approval event raised by the LinkList contract.
type LinkListApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LinkList *LinkListFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LinkListApprovalIterator, error) {

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

	logs, sub, err := _LinkList.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LinkListApprovalIterator{contract: _LinkList.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LinkList *LinkListFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LinkListApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LinkList.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkListApproval)
				if err := _LinkList.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_LinkList *LinkListFilterer) ParseApproval(log types.Log) (*LinkListApproval, error) {
	event := new(LinkListApproval)
	if err := _LinkList.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LinkListApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LinkList contract.
type LinkListApprovalForAllIterator struct {
	Event *LinkListApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LinkListApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkListApprovalForAll)
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
		it.Event = new(LinkListApprovalForAll)
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
func (it *LinkListApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkListApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkListApprovalForAll represents a ApprovalForAll event raised by the LinkList contract.
type LinkListApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LinkList *LinkListFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LinkListApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LinkList.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LinkListApprovalForAllIterator{contract: _LinkList.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LinkList *LinkListFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LinkListApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LinkList.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkListApprovalForAll)
				if err := _LinkList.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_LinkList *LinkListFilterer) ParseApprovalForAll(log types.Log) (*LinkListApprovalForAll, error) {
	event := new(LinkListApprovalForAll)
	if err := _LinkList.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LinkListTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LinkList contract.
type LinkListTransferIterator struct {
	Event *LinkListTransfer // Event containing the contract specifics and raw log

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
func (it *LinkListTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LinkListTransfer)
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
		it.Event = new(LinkListTransfer)
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
func (it *LinkListTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LinkListTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LinkListTransfer represents a Transfer event raised by the LinkList contract.
type LinkListTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LinkList *LinkListFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LinkListTransferIterator, error) {

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

	logs, sub, err := _LinkList.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LinkListTransferIterator{contract: _LinkList.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LinkList *LinkListFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LinkListTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _LinkList.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LinkListTransfer)
				if err := _LinkList.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_LinkList *LinkListFilterer) ParseTransfer(log types.Log) (*LinkListTransfer, error) {
	event := new(LinkListTransfer)
	if err := _LinkList.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
