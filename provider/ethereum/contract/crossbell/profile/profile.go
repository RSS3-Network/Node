// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package profile

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

// DataTypesCreateProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateProfileData struct {
	To                 common.Address
	Handle             string
	Uri                string
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypesERC721Struct is an auto generated low-level Go binding around an user-defined struct.
type DataTypesERC721Struct struct {
	TokenAddress  common.Address
	Erc721TokenId *big.Int
}

// DataTypesMintNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesMintNoteData struct {
	ProfileId      *big.Int
	NoteId         *big.Int
	To             common.Address
	MintModuleData []byte
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
	ProfileId *big.Int
	NoteId    *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	ProfileId          *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypesProfile is an auto generated low-level Go binding around an user-defined struct.
type DataTypesProfile struct {
	ProfileId   *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypescreateThenLinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkProfileData struct {
	FromProfileId *big.Int
	To            common.Address
	LinkType      [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromProfileId *big.Int
	EthAddress    common.Address
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	TokenId       *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromProfileId *big.Int
	ToLinkListId  *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypeslinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkProfileData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
	Data          []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ERC721Data struct {
	TokenAddress       common.Address
	TokenId            *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4LinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4LinklistData struct {
	LinklistId         *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4NoteData struct {
	ProfileId          *big.Int
	NoteId             *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4ProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4ProfileData struct {
	ProfileId          *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	ProfileId          *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAddressData struct {
	FromProfileId *big.Int
	EthAddress    common.Address
	LinkType      [32]byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	TokenId       *big.Int
	LinkType      [32]byte
}

// DataTypesunlinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkLinklistData struct {
	FromProfileId *big.Int
	ToLinkListId  *big.Int
	LinkType      [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
}

// DataTypesunlinkProfileData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkProfileData struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
}

// ProfileMetaData contains all meta data concerning the Profile contract.
var ProfileMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canCreate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Profile\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Profile\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfileUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryProfileId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_linklistContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_periphery\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryProfile\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"migrateNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"}],\"name\":\"postNote4Profile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setProfileUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4ProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Profile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4ERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4ERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryProfileId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setProfileUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"AttachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProfileCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"DeleteNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"DetachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"LinkProfileLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LinklistNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"LockNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MintNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MintNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PostNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetProfileUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"SetHandle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Profile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4ERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Linklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetMintModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetNoteUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldProfileId\",\"type\":\"uint256\"}],\"name\":\"SetPrimaryProfileId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"SetSocialToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkProfile\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromProfileeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkProfileLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toProfileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Web3EntryInitialized\",\"type\":\"event\"}]",
}

// ProfileABI is the input ABI used to generate the binding from.
// Deprecated: Use ProfileMetaData.ABI instead.
var ProfileABI = ProfileMetaData.ABI

// Profile is an auto generated Go binding around an Ethereum contract.
type Profile struct {
	ProfileCaller     // Read-only binding to the contract
	ProfileTransactor // Write-only binding to the contract
	ProfileFilterer   // Log filterer for contract events
}

// ProfileCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProfileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProfileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProfileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProfileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProfileSession struct {
	Contract     *Profile          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProfileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProfileCallerSession struct {
	Contract *ProfileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProfileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProfileTransactorSession struct {
	Contract     *ProfileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProfileRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProfileRaw struct {
	Contract *Profile // Generic contract binding to access the raw methods on
}

// ProfileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProfileCallerRaw struct {
	Contract *ProfileCaller // Generic read-only contract binding to access the raw methods on
}

// ProfileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProfileTransactorRaw struct {
	Contract *ProfileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProfile creates a new instance of Profile, bound to a specific deployed contract.
func NewProfile(address common.Address, backend bind.ContractBackend) (*Profile, error) {
	contract, err := bindProfile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Profile{ProfileCaller: ProfileCaller{contract: contract}, ProfileTransactor: ProfileTransactor{contract: contract}, ProfileFilterer: ProfileFilterer{contract: contract}}, nil
}

// NewProfileCaller creates a new read-only instance of Profile, bound to a specific deployed contract.
func NewProfileCaller(address common.Address, caller bind.ContractCaller) (*ProfileCaller, error) {
	contract, err := bindProfile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileCaller{contract: contract}, nil
}

// NewProfileTransactor creates a new write-only instance of Profile, bound to a specific deployed contract.
func NewProfileTransactor(address common.Address, transactor bind.ContractTransactor) (*ProfileTransactor, error) {
	contract, err := bindProfile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProfileTransactor{contract: contract}, nil
}

// NewProfileFilterer creates a new log filterer instance of Profile, bound to a specific deployed contract.
func NewProfileFilterer(address common.Address, filterer bind.ContractFilterer) (*ProfileFilterer, error) {
	contract, err := bindProfile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProfileFilterer{contract: contract}, nil
}

// bindProfile binds a generic wrapper to an already deployed contract.
func bindProfile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProfileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.ProfileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.ProfileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Profile *ProfileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Profile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Profile *ProfileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Profile *ProfileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Profile.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Profile *ProfileCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Profile *ProfileSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Profile.Contract.BalanceOf(&_Profile.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Profile *ProfileCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Profile.Contract.BalanceOf(&_Profile.CallOpts, owner)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Profile *ProfileCaller) CanCreate(opts *bind.CallOpts, handle string, account common.Address) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "canCreate", handle, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Profile *ProfileSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Profile.Contract.CanCreate(&_Profile.CallOpts, handle, account)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Profile *ProfileCallerSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Profile.Contract.CanCreate(&_Profile.CallOpts, handle, account)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Profile *ProfileCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Profile *ProfileSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetApproved(&_Profile.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Profile *ProfileCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetApproved(&_Profile.CallOpts, tokenId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Profile *ProfileCaller) GetHandle(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getHandle", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Profile *ProfileSession) GetHandle(profileId *big.Int) (string, error) {
	return _Profile.Contract.GetHandle(&_Profile.CallOpts, profileId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_Profile *ProfileCallerSession) GetHandle(profileId *big.Int) (string, error) {
	return _Profile.Contract.GetHandle(&_Profile.CallOpts, profileId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Profile *ProfileCaller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Profile *ProfileSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4Address(&_Profile.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Profile *ProfileCallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4Address(&_Profile.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Profile *ProfileCaller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Profile *ProfileSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4ERC721(&_Profile.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Profile *ProfileCallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4ERC721(&_Profile.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Profile *ProfileCaller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Profile *ProfileSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4Linklist(&_Profile.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Profile *ProfileCallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetLinkModule4Linklist(&_Profile.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Profile *ProfileCaller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Profile *ProfileSession) GetLinklistContract() (common.Address, error) {
	return _Profile.Contract.GetLinklistContract(&_Profile.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Profile *ProfileCallerSession) GetLinklistContract() (common.Address, error) {
	return _Profile.Contract.GetLinklistContract(&_Profile.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_Profile *ProfileCaller) GetLinklistId(opts *bind.CallOpts, profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinklistId", profileId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_Profile *ProfileSession) GetLinklistId(profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Profile.Contract.GetLinklistId(&_Profile.CallOpts, profileId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 profileId, bytes32 linkType) view returns(uint256)
func (_Profile *ProfileCallerSession) GetLinklistId(profileId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Profile.Contract.GetLinklistId(&_Profile.CallOpts, profileId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Profile *ProfileCaller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Profile *ProfileSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Profile.Contract.GetLinklistType(&_Profile.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Profile *ProfileCallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Profile.Contract.GetLinklistType(&_Profile.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Profile *ProfileCaller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Profile *ProfileSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Profile.Contract.GetLinklistUri(&_Profile.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Profile *ProfileCallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Profile.Contract.GetLinklistUri(&_Profile.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Profile *ProfileCaller) GetNote(opts *bind.CallOpts, profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getNote", profileId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Profile *ProfileSession) GetNote(profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Profile.Contract.GetNote(&_Profile.CallOpts, profileId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 profileId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Profile *ProfileCallerSession) GetNote(profileId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Profile.Contract.GetNote(&_Profile.CallOpts, profileId, noteId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_Profile *ProfileCaller) GetOperator(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getOperator", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_Profile *ProfileSession) GetOperator(profileId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetOperator(&_Profile.CallOpts, profileId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 profileId) view returns(address)
func (_Profile *ProfileCallerSession) GetOperator(profileId *big.Int) (common.Address, error) {
	return _Profile.Contract.GetOperator(&_Profile.CallOpts, profileId)
}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_Profile *ProfileCaller) GetPrimaryProfileId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getPrimaryProfileId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_Profile *ProfileSession) GetPrimaryProfileId(account common.Address) (*big.Int, error) {
	return _Profile.Contract.GetPrimaryProfileId(&_Profile.CallOpts, account)
}

// GetPrimaryProfileId is a free data retrieval call binding the contract method 0x1bb4d231.
//
// Solidity: function getPrimaryProfileId(address account) view returns(uint256)
func (_Profile *ProfileCallerSession) GetPrimaryProfileId(account common.Address) (*big.Int, error) {
	return _Profile.Contract.GetPrimaryProfileId(&_Profile.CallOpts, account)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileCaller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (DataTypesProfile, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(DataTypesProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfile)).(*DataTypesProfile)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileSession) GetProfile(profileId *big.Int) (DataTypesProfile, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileCallerSession) GetProfile(profileId *big.Int) (DataTypesProfile, error) {
	return _Profile.Contract.GetProfile(&_Profile.CallOpts, profileId)
}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileCaller) GetProfileByHandle(opts *bind.CallOpts, handle string) (DataTypesProfile, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfileByHandle", handle)

	if err != nil {
		return *new(DataTypesProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesProfile)).(*DataTypesProfile)

	return out0, err

}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileSession) GetProfileByHandle(handle string) (DataTypesProfile, error) {
	return _Profile.Contract.GetProfileByHandle(&_Profile.CallOpts, handle)
}

// GetProfileByHandle is a free data retrieval call binding the contract method 0x0c16de10.
//
// Solidity: function getProfileByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Profile *ProfileCallerSession) GetProfileByHandle(handle string) (DataTypesProfile, error) {
	return _Profile.Contract.GetProfileByHandle(&_Profile.CallOpts, handle)
}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_Profile *ProfileCaller) GetProfileUri(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getProfileUri", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_Profile *ProfileSession) GetProfileUri(profileId *big.Int) (string, error) {
	return _Profile.Contract.GetProfileUri(&_Profile.CallOpts, profileId)
}

// GetProfileUri is a free data retrieval call binding the contract method 0xcba4f5cc.
//
// Solidity: function getProfileUri(uint256 profileId) view returns(string)
func (_Profile *ProfileCallerSession) GetProfileUri(profileId *big.Int) (string, error) {
	return _Profile.Contract.GetProfileUri(&_Profile.CallOpts, profileId)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Profile *ProfileCaller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Profile *ProfileSession) GetRevision() (*big.Int, error) {
	return _Profile.Contract.GetRevision(&_Profile.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Profile *ProfileCallerSession) GetRevision() (*big.Int, error) {
	return _Profile.Contract.GetRevision(&_Profile.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Profile *ProfileCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Profile *ProfileSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Profile.Contract.IsApprovedForAll(&_Profile.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Profile *ProfileCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Profile.Contract.IsApprovedForAll(&_Profile.CallOpts, owner, operator)
}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_Profile *ProfileCaller) IsPrimaryProfile(opts *bind.CallOpts, profileId *big.Int) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "isPrimaryProfile", profileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_Profile *ProfileSession) IsPrimaryProfile(profileId *big.Int) (bool, error) {
	return _Profile.Contract.IsPrimaryProfile(&_Profile.CallOpts, profileId)
}

// IsPrimaryProfile is a free data retrieval call binding the contract method 0xc387c453.
//
// Solidity: function isPrimaryProfile(uint256 profileId) view returns(bool)
func (_Profile *ProfileCallerSession) IsPrimaryProfile(profileId *big.Int) (bool, error) {
	return _Profile.Contract.IsPrimaryProfile(&_Profile.CallOpts, profileId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Profile *ProfileCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Profile *ProfileSession) Name() (string, error) {
	return _Profile.Contract.Name(&_Profile.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Profile *ProfileCallerSession) Name() (string, error) {
	return _Profile.Contract.Name(&_Profile.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Profile *ProfileCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Profile *ProfileSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.OwnerOf(&_Profile.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Profile *ProfileCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Profile.Contract.OwnerOf(&_Profile.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Profile *ProfileCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Profile *ProfileSession) Resolver() (common.Address, error) {
	return _Profile.Contract.Resolver(&_Profile.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Profile *ProfileCallerSession) Resolver() (common.Address, error) {
	return _Profile.Contract.Resolver(&_Profile.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Profile *ProfileCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Profile *ProfileSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Profile.Contract.SupportsInterface(&_Profile.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Profile *ProfileCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Profile.Contract.SupportsInterface(&_Profile.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Profile *ProfileCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Profile *ProfileSession) Symbol() (string, error) {
	return _Profile.Contract.Symbol(&_Profile.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Profile *ProfileCallerSession) Symbol() (string, error) {
	return _Profile.Contract.Symbol(&_Profile.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Profile *ProfileCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Profile *ProfileSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Profile.Contract.TokenByIndex(&_Profile.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Profile *ProfileCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Profile.Contract.TokenByIndex(&_Profile.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Profile *ProfileCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Profile *ProfileSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Profile.Contract.TokenOfOwnerByIndex(&_Profile.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Profile *ProfileCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Profile.Contract.TokenOfOwnerByIndex(&_Profile.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_Profile *ProfileCaller) TokenURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "tokenURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_Profile *ProfileSession) TokenURI(profileId *big.Int) (string, error) {
	return _Profile.Contract.TokenURI(&_Profile.CallOpts, profileId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 profileId) view returns(string)
func (_Profile *ProfileCallerSession) TokenURI(profileId *big.Int) (string, error) {
	return _Profile.Contract.TokenURI(&_Profile.CallOpts, profileId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Profile *ProfileCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Profile.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Profile *ProfileSession) TotalSupply() (*big.Int, error) {
	return _Profile.Contract.TotalSupply(&_Profile.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Profile *ProfileCallerSession) TotalSupply() (*big.Int, error) {
	return _Profile.Contract.TotalSupply(&_Profile.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Profile *ProfileSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.Approve(&_Profile.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.Approve(&_Profile.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Profile *ProfileTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Profile *ProfileSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.Burn(&_Profile.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Profile *ProfileTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.Burn(&_Profile.TransactOpts, tokenId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) CreateProfile(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile(&_Profile.TransactOpts, vars)
}

// CreateProfile0 is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) CreateProfile0(opts *bind.TransactOpts, vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "createProfile0", vars)
}

// CreateProfile0 is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileSession) CreateProfile0(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile0(&_Profile.TransactOpts, vars)
}

// CreateProfile0 is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) CreateProfile0(vars DataTypesCreateProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateProfile0(&_Profile.TransactOpts, vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileTransactor) CreateThenLinkProfile(opts *bind.TransactOpts, vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "createThenLinkProfile", vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileSession) CreateThenLinkProfile(vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateThenLinkProfile(&_Profile.TransactOpts, vars)
}

// CreateThenLinkProfile is a paid mutator transaction binding the contract method 0x0ab6beba.
//
// Solidity: function createThenLinkProfile((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) CreateThenLinkProfile(vars DataTypescreateThenLinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.CreateThenLinkProfile(&_Profile.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileTransactor) DeleteNote(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "deleteNote", profileId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileSession) DeleteNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.DeleteNote(&_Profile.TransactOpts, profileId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileTransactorSession) DeleteNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.DeleteNote(&_Profile.TransactOpts, profileId, noteId)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Profile *ProfileTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "initialize", _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Profile *ProfileSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Profile *ProfileTransactorSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Profile.Contract.Initialize(&_Profile.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Profile.Contract.LinkAddress(&_Profile.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Profile.Contract.LinkAddress(&_Profile.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Profile.Contract.LinkAnyUri(&_Profile.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Profile.Contract.LinkAnyUri(&_Profile.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.LinkERC721(&_Profile.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.LinkERC721(&_Profile.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Profile.Contract.LinkLinklist(&_Profile.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Profile.Contract.LinkLinklist(&_Profile.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Profile.Contract.LinkNote(&_Profile.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Profile.Contract.LinkNote(&_Profile.TransactOpts, vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactor) LinkProfile(opts *bind.TransactOpts, vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "linkProfile", vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileSession) LinkProfile(vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.LinkProfile(&_Profile.TransactOpts, vars)
}

// LinkProfile is a paid mutator transaction binding the contract method 0xa914c76e.
//
// Solidity: function linkProfile((uint256,uint256,bytes32,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) LinkProfile(vars DataTypeslinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.LinkProfile(&_Profile.TransactOpts, vars)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileTransactor) LockNote(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "lockNote", profileId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileSession) LockNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.LockNote(&_Profile.TransactOpts, profileId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 profileId, uint256 noteId) returns()
func (_Profile *ProfileTransactorSession) LockNote(profileId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.LockNote(&_Profile.TransactOpts, profileId, noteId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 profileId) returns()
func (_Profile *ProfileTransactor) MigrateNote(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "migrateNote", profileId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 profileId) returns()
func (_Profile *ProfileSession) MigrateNote(profileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.MigrateNote(&_Profile.TransactOpts, profileId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 profileId) returns()
func (_Profile *ProfileTransactorSession) MigrateNote(profileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.MigrateNote(&_Profile.TransactOpts, profileId)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Profile *ProfileTransactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Profile *ProfileSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Profile.Contract.MintNote(&_Profile.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Profile *ProfileTransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Profile.Contract.MintNote(&_Profile.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Profile *ProfileTransactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Profile *ProfileSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Profile.Contract.PostNote(&_Profile.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Profile.Contract.PostNote(&_Profile.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4Address(opts *bind.TransactOpts, noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4Address", noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Profile *ProfileSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Address(&_Profile.TransactOpts, noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Address(&_Profile.TransactOpts, noteData, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4AnyUri(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4AnyUri", postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Profile *ProfileSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4AnyUri(&_Profile.TransactOpts, postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4AnyUri(&_Profile.TransactOpts, postNoteData, uri)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4ERC721(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4ERC721", postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Profile *ProfileSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4ERC721(&_Profile.TransactOpts, postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4ERC721(&_Profile.TransactOpts, postNoteData, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4Linklist(opts *bind.TransactOpts, noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4Linklist", noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Profile *ProfileSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Linklist(&_Profile.TransactOpts, noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Linklist(&_Profile.TransactOpts, noteData, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4Note(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4Note", postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Profile *ProfileSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Note(&_Profile.TransactOpts, postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Note(&_Profile.TransactOpts, postNoteData, note)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_Profile *ProfileTransactor) PostNote4Profile(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "postNote4Profile", postNoteData, toProfileId)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_Profile *ProfileSession) PostNote4Profile(postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Profile(&_Profile.TransactOpts, postNoteData, toProfileId)
}

// PostNote4Profile is a paid mutator transaction binding the contract method 0x2ff5b07e.
//
// Solidity: function postNote4Profile((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toProfileId) returns(uint256)
func (_Profile *ProfileTransactorSession) PostNote4Profile(postNoteData DataTypesPostNoteData, toProfileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.PostNote4Profile(&_Profile.TransactOpts, postNoteData, toProfileId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SafeTransferFrom(&_Profile.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SafeTransferFrom(&_Profile.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Profile *ProfileTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Profile *ProfileSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Profile.Contract.SafeTransferFrom0(&_Profile.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Profile *ProfileTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Profile.Contract.SafeTransferFrom0(&_Profile.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Profile *ProfileTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Profile *ProfileSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Profile.Contract.SetApprovalForAll(&_Profile.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Profile *ProfileTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Profile.Contract.SetApprovalForAll(&_Profile.TransactOpts, operator, approved)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_Profile *ProfileTransactor) SetHandle(opts *bind.TransactOpts, profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setHandle", profileId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_Profile *ProfileSession) SetHandle(profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Profile.Contract.SetHandle(&_Profile.TransactOpts, profileId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 profileId, string newHandle) returns()
func (_Profile *ProfileTransactorSession) SetHandle(profileId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Profile.Contract.SetHandle(&_Profile.TransactOpts, profileId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Address(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Address(&_Profile.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetLinkModule4ERC721(opts *bind.TransactOpts, vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinkModule4ERC721", vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4ERC721(&_Profile.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4ERC721(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Linklist(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Linklist(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetLinkModule4Note(opts *bind.TransactOpts, vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinkModule4Note", vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Note(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Note(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetLinkModule4Profile(opts *bind.TransactOpts, vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinkModule4Profile", vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetLinkModule4Profile(vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Profile(&_Profile.TransactOpts, vars)
}

// SetLinkModule4Profile is a paid mutator transaction binding the contract method 0x5b507cfd.
//
// Solidity: function setLinkModule4Profile((uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetLinkModule4Profile(vars DataTypessetLinkModule4ProfileData) (*types.Transaction, error) {
	return _Profile.Contract.SetLinkModule4Profile(&_Profile.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Profile *ProfileTransactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Profile *ProfileSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Profile.Contract.SetLinklistUri(&_Profile.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Profile *ProfileTransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Profile.Contract.SetLinklistUri(&_Profile.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Profile.Contract.SetMintModule4Note(&_Profile.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Profile *ProfileTransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Profile.Contract.SetMintModule4Note(&_Profile.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_Profile *ProfileTransactor) SetNoteUri(opts *bind.TransactOpts, profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setNoteUri", profileId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_Profile *ProfileSession) SetNoteUri(profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetNoteUri(&_Profile.TransactOpts, profileId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 profileId, uint256 noteId, string newUri) returns()
func (_Profile *ProfileTransactorSession) SetNoteUri(profileId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetNoteUri(&_Profile.TransactOpts, profileId, noteId, newUri)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_Profile *ProfileTransactor) SetOperator(opts *bind.TransactOpts, profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setOperator", profileId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_Profile *ProfileSession) SetOperator(profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetOperator(&_Profile.TransactOpts, profileId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 profileId, address operator) returns()
func (_Profile *ProfileTransactorSession) SetOperator(profileId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetOperator(&_Profile.TransactOpts, profileId, operator)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_Profile *ProfileTransactor) SetPrimaryProfileId(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setPrimaryProfileId", profileId)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_Profile *ProfileSession) SetPrimaryProfileId(profileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SetPrimaryProfileId(&_Profile.TransactOpts, profileId)
}

// SetPrimaryProfileId is a paid mutator transaction binding the contract method 0x295cb43e.
//
// Solidity: function setPrimaryProfileId(uint256 profileId) returns()
func (_Profile *ProfileTransactorSession) SetPrimaryProfileId(profileId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.SetPrimaryProfileId(&_Profile.TransactOpts, profileId)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileTransactor) SetProfileUri(opts *bind.TransactOpts, profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setProfileUri", profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetProfileUri(&_Profile.TransactOpts, profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileTransactorSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetProfileUri(&_Profile.TransactOpts, profileId, newUri)
}

// SetProfileUri0 is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileTransactor) SetProfileUri0(opts *bind.TransactOpts, profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setProfileUri0", profileId, newUri)
}

// SetProfileUri0 is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileSession) SetProfileUri0(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetProfileUri0(&_Profile.TransactOpts, profileId, newUri)
}

// SetProfileUri0 is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Profile *ProfileTransactorSession) SetProfileUri0(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Profile.Contract.SetProfileUri0(&_Profile.TransactOpts, profileId, newUri)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_Profile *ProfileTransactor) SetSocialToken(opts *bind.TransactOpts, profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "setSocialToken", profileId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_Profile *ProfileSession) SetSocialToken(profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetSocialToken(&_Profile.TransactOpts, profileId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 profileId, address tokenAddress) returns()
func (_Profile *ProfileTransactorSession) SetSocialToken(profileId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Profile.Contract.SetSocialToken(&_Profile.TransactOpts, profileId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.TransferFrom(&_Profile.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Profile *ProfileTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Profile.Contract.TransferFrom(&_Profile.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkAddress(&_Profile.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkAddress(&_Profile.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkAnyUri(&_Profile.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkAnyUri(&_Profile.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkERC721(&_Profile.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkERC721(&_Profile.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkLinklist(&_Profile.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkLinklist(&_Profile.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkNote(&_Profile.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkNote(&_Profile.TransactOpts, vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactor) UnlinkProfile(opts *bind.TransactOpts, vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _Profile.contract.Transact(opts, "unlinkProfile", vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileSession) UnlinkProfile(vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkProfile(&_Profile.TransactOpts, vars)
}

// UnlinkProfile is a paid mutator transaction binding the contract method 0x251a99cf.
//
// Solidity: function unlinkProfile((uint256,uint256,bytes32) vars) returns()
func (_Profile *ProfileTransactorSession) UnlinkProfile(vars DataTypesunlinkProfileData) (*types.Transaction, error) {
	return _Profile.Contract.UnlinkProfile(&_Profile.TransactOpts, vars)
}

// ProfileApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Profile contract.
type ProfileApprovalIterator struct {
	Event *ProfileApproval // Event containing the contract specifics and raw log

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
func (it *ProfileApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileApproval)
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
		it.Event = new(ProfileApproval)
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
func (it *ProfileApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileApproval represents a Approval event raised by the Profile contract.
type ProfileApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Profile *ProfileFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ProfileApprovalIterator, error) {

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

	logs, sub, err := _Profile.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileApprovalIterator{contract: _Profile.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Profile *ProfileFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProfileApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Profile.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileApproval)
				if err := _Profile.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseApproval(log types.Log) (*ProfileApproval, error) {
	event := new(ProfileApproval)
	if err := _Profile.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Profile contract.
type ProfileApprovalForAllIterator struct {
	Event *ProfileApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ProfileApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileApprovalForAll)
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
		it.Event = new(ProfileApprovalForAll)
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
func (it *ProfileApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileApprovalForAll represents a ApprovalForAll event raised by the Profile contract.
type ProfileApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Profile *ProfileFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ProfileApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProfileApprovalForAllIterator{contract: _Profile.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Profile *ProfileFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ProfileApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileApprovalForAll)
				if err := _Profile.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseApprovalForAll(log types.Log) (*ProfileApprovalForAll, error) {
	event := new(ProfileApprovalForAll)
	if err := _Profile.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileAttachLinklistIterator is returned from FilterAttachLinklist and is used to iterate over the raw logs and unpacked data for AttachLinklist events raised by the Profile contract.
type ProfileAttachLinklistIterator struct {
	Event *ProfileAttachLinklist // Event containing the contract specifics and raw log

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
func (it *ProfileAttachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileAttachLinklist)
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
		it.Event = new(ProfileAttachLinklist)
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
func (it *ProfileAttachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileAttachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileAttachLinklist represents a AttachLinklist event raised by the Profile contract.
type ProfileAttachLinklist struct {
	LinklistId *big.Int
	ProfileId  *big.Int
	LinkType   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAttachLinklist is a free log retrieval operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) FilterAttachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, profileId []*big.Int, linkType [][32]byte) (*ProfileAttachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "AttachLinklist", linklistIdRule, profileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &ProfileAttachLinklistIterator{contract: _Profile.contract, event: "AttachLinklist", logs: logs, sub: sub}, nil
}

// WatchAttachLinklist is a free log subscription operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) WatchAttachLinklist(opts *bind.WatchOpts, sink chan<- *ProfileAttachLinklist, linklistId []*big.Int, profileId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "AttachLinklist", linklistIdRule, profileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileAttachLinklist)
				if err := _Profile.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
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

// ParseAttachLinklist is a log parse operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) ParseAttachLinklist(log types.Log) (*ProfileAttachLinklist, error) {
	event := new(ProfileAttachLinklist)
	if err := _Profile.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the Profile contract.
type ProfileBaseInitializedIterator struct {
	Event *ProfileBaseInitialized // Event containing the contract specifics and raw log

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
func (it *ProfileBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileBaseInitialized)
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
		it.Event = new(ProfileBaseInitialized)
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
func (it *ProfileBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileBaseInitialized represents a BaseInitialized event raised by the Profile contract.
type ProfileBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*ProfileBaseInitializedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &ProfileBaseInitializedIterator{contract: _Profile.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *ProfileBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileBaseInitialized)
				if err := _Profile.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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

// ParseBaseInitialized is a log parse operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseBaseInitialized(log types.Log) (*ProfileBaseInitialized, error) {
	event := new(ProfileBaseInitialized)
	if err := _Profile.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileDeleteNoteIterator is returned from FilterDeleteNote and is used to iterate over the raw logs and unpacked data for DeleteNote events raised by the Profile contract.
type ProfileDeleteNoteIterator struct {
	Event *ProfileDeleteNote // Event containing the contract specifics and raw log

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
func (it *ProfileDeleteNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileDeleteNote)
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
		it.Event = new(ProfileDeleteNote)
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
func (it *ProfileDeleteNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileDeleteNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileDeleteNote represents a DeleteNote event raised by the Profile contract.
type ProfileDeleteNote struct {
	ProfileId *big.Int
	NoteId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeleteNote is a free log retrieval operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) FilterDeleteNote(opts *bind.FilterOpts, profileId []*big.Int) (*ProfileDeleteNoteIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "DeleteNote", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileDeleteNoteIterator{contract: _Profile.contract, event: "DeleteNote", logs: logs, sub: sub}, nil
}

// WatchDeleteNote is a free log subscription operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) WatchDeleteNote(opts *bind.WatchOpts, sink chan<- *ProfileDeleteNote, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "DeleteNote", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileDeleteNote)
				if err := _Profile.contract.UnpackLog(event, "DeleteNote", log); err != nil {
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

// ParseDeleteNote is a log parse operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) ParseDeleteNote(log types.Log) (*ProfileDeleteNote, error) {
	event := new(ProfileDeleteNote)
	if err := _Profile.contract.UnpackLog(event, "DeleteNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileDetachLinklistIterator is returned from FilterDetachLinklist and is used to iterate over the raw logs and unpacked data for DetachLinklist events raised by the Profile contract.
type ProfileDetachLinklistIterator struct {
	Event *ProfileDetachLinklist // Event containing the contract specifics and raw log

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
func (it *ProfileDetachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileDetachLinklist)
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
		it.Event = new(ProfileDetachLinklist)
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
func (it *ProfileDetachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileDetachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileDetachLinklist represents a DetachLinklist event raised by the Profile contract.
type ProfileDetachLinklist struct {
	LinklistId *big.Int
	ProfileId  *big.Int
	LinkType   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDetachLinklist is a free log retrieval operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) FilterDetachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, profileId []*big.Int, linkType [][32]byte) (*ProfileDetachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "DetachLinklist", linklistIdRule, profileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &ProfileDetachLinklistIterator{contract: _Profile.contract, event: "DetachLinklist", logs: logs, sub: sub}, nil
}

// WatchDetachLinklist is a free log subscription operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) WatchDetachLinklist(opts *bind.WatchOpts, sink chan<- *ProfileDetachLinklist, linklistId []*big.Int, profileId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "DetachLinklist", linklistIdRule, profileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileDetachLinklist)
				if err := _Profile.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
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

// ParseDetachLinklist is a log parse operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed profileId, bytes32 indexed linkType)
func (_Profile *ProfileFilterer) ParseDetachLinklist(log types.Log) (*ProfileDetachLinklist, error) {
	event := new(ProfileDetachLinklist)
	if err := _Profile.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkAddressIterator is returned from FilterLinkAddress and is used to iterate over the raw logs and unpacked data for LinkAddress events raised by the Profile contract.
type ProfileLinkAddressIterator struct {
	Event *ProfileLinkAddress // Event containing the contract specifics and raw log

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
func (it *ProfileLinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkAddress)
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
		it.Event = new(ProfileLinkAddress)
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
func (it *ProfileLinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkAddress represents a LinkAddress event raised by the Profile contract.
type ProfileLinkAddress struct {
	FromProfileId *big.Int
	EthAddress    common.Address
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkAddress is a free log retrieval operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterLinkAddress(opts *bind.FilterOpts, fromProfileId []*big.Int, ethAddress []common.Address) (*ProfileLinkAddressIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkAddress", fromProfileIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkAddressIterator{contract: _Profile.contract, event: "LinkAddress", logs: logs, sub: sub}, nil
}

// WatchLinkAddress is a free log subscription operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchLinkAddress(opts *bind.WatchOpts, sink chan<- *ProfileLinkAddress, fromProfileId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkAddress", fromProfileIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkAddress)
				if err := _Profile.contract.UnpackLog(event, "LinkAddress", log); err != nil {
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

// ParseLinkAddress is a log parse operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseLinkAddress(log types.Log) (*ProfileLinkAddress, error) {
	event := new(ProfileLinkAddress)
	if err := _Profile.contract.UnpackLog(event, "LinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkAnyUriIterator is returned from FilterLinkAnyUri and is used to iterate over the raw logs and unpacked data for LinkAnyUri events raised by the Profile contract.
type ProfileLinkAnyUriIterator struct {
	Event *ProfileLinkAnyUri // Event containing the contract specifics and raw log

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
func (it *ProfileLinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkAnyUri)
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
		it.Event = new(ProfileLinkAnyUri)
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
func (it *ProfileLinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkAnyUri represents a LinkAnyUri event raised by the Profile contract.
type ProfileLinkAnyUri struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkAnyUri is a free log retrieval operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterLinkAnyUri(opts *bind.FilterOpts, fromProfileId []*big.Int) (*ProfileLinkAnyUriIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkAnyUri", fromProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkAnyUriIterator{contract: _Profile.contract, event: "LinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchLinkAnyUri is a free log subscription operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchLinkAnyUri(opts *bind.WatchOpts, sink chan<- *ProfileLinkAnyUri, fromProfileId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkAnyUri", fromProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkAnyUri)
				if err := _Profile.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
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

// ParseLinkAnyUri is a log parse operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseLinkAnyUri(log types.Log) (*ProfileLinkAnyUri, error) {
	event := new(ProfileLinkAnyUri)
	if err := _Profile.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkERC721Iterator is returned from FilterLinkERC721 and is used to iterate over the raw logs and unpacked data for LinkERC721 events raised by the Profile contract.
type ProfileLinkERC721Iterator struct {
	Event *ProfileLinkERC721 // Event containing the contract specifics and raw log

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
func (it *ProfileLinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkERC721)
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
		it.Event = new(ProfileLinkERC721)
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
func (it *ProfileLinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkERC721 represents a LinkERC721 event raised by the Profile contract.
type ProfileLinkERC721 struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	ToNoteId      *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkERC721 is a free log retrieval operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterLinkERC721(opts *bind.FilterOpts, fromProfileId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*ProfileLinkERC721Iterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkERC721", fromProfileIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkERC721Iterator{contract: _Profile.contract, event: "LinkERC721", logs: logs, sub: sub}, nil
}

// WatchLinkERC721 is a free log subscription operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchLinkERC721(opts *bind.WatchOpts, sink chan<- *ProfileLinkERC721, fromProfileId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkERC721", fromProfileIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkERC721)
				if err := _Profile.contract.UnpackLog(event, "LinkERC721", log); err != nil {
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

// ParseLinkERC721 is a log parse operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseLinkERC721(log types.Log) (*ProfileLinkERC721, error) {
	event := new(ProfileLinkERC721)
	if err := _Profile.contract.UnpackLog(event, "LinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkLinklistIterator is returned from FilterLinkLinklist and is used to iterate over the raw logs and unpacked data for LinkLinklist events raised by the Profile contract.
type ProfileLinkLinklistIterator struct {
	Event *ProfileLinkLinklist // Event containing the contract specifics and raw log

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
func (it *ProfileLinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkLinklist)
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
		it.Event = new(ProfileLinkLinklist)
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
func (it *ProfileLinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkLinklist represents a LinkLinklist event raised by the Profile contract.
type ProfileLinkLinklist struct {
	FromProfileId *big.Int
	ToLinklistId  *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkLinklist is a free log retrieval operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) FilterLinkLinklist(opts *bind.FilterOpts, fromProfileId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*ProfileLinkLinklistIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkLinklist", fromProfileIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkLinklistIterator{contract: _Profile.contract, event: "LinkLinklist", logs: logs, sub: sub}, nil
}

// WatchLinkLinklist is a free log subscription operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) WatchLinkLinklist(opts *bind.WatchOpts, sink chan<- *ProfileLinkLinklist, fromProfileId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkLinklist", fromProfileIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkLinklist)
				if err := _Profile.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
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

// ParseLinkLinklist is a log parse operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) ParseLinkLinklist(log types.Log) (*ProfileLinkLinklist, error) {
	event := new(ProfileLinkLinklist)
	if err := _Profile.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkNoteIterator is returned from FilterLinkNote and is used to iterate over the raw logs and unpacked data for LinkNote events raised by the Profile contract.
type ProfileLinkNoteIterator struct {
	Event *ProfileLinkNote // Event containing the contract specifics and raw log

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
func (it *ProfileLinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkNote)
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
		it.Event = new(ProfileLinkNote)
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
func (it *ProfileLinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkNote represents a LinkNote event raised by the Profile contract.
type ProfileLinkNote struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkNote is a free log retrieval operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterLinkNote(opts *bind.FilterOpts, fromProfileId []*big.Int, toProfileId []*big.Int, toNoteId []*big.Int) (*ProfileLinkNoteIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkNote", fromProfileIdRule, toProfileIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkNoteIterator{contract: _Profile.contract, event: "LinkNote", logs: logs, sub: sub}, nil
}

// WatchLinkNote is a free log subscription operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchLinkNote(opts *bind.WatchOpts, sink chan<- *ProfileLinkNote, fromProfileId []*big.Int, toProfileId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkNote", fromProfileIdRule, toProfileIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkNote)
				if err := _Profile.contract.UnpackLog(event, "LinkNote", log); err != nil {
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

// ParseLinkNote is a log parse operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseLinkNote(log types.Log) (*ProfileLinkNote, error) {
	event := new(ProfileLinkNote)
	if err := _Profile.contract.UnpackLog(event, "LinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkProfileIterator is returned from FilterLinkProfile and is used to iterate over the raw logs and unpacked data for LinkProfile events raised by the Profile contract.
type ProfileLinkProfileIterator struct {
	Event *ProfileLinkProfile // Event containing the contract specifics and raw log

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
func (it *ProfileLinkProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkProfile)
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
		it.Event = new(ProfileLinkProfile)
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
func (it *ProfileLinkProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkProfile represents a LinkProfile event raised by the Profile contract.
type ProfileLinkProfile struct {
	Account       common.Address
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLinkProfile is a free log retrieval operation binding the contract event 0xbc914995d574dd9ef2df364e4eee2b85deda3ba35d054a62425fba1b97275716.
//
// Solidity: event LinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterLinkProfile(opts *bind.FilterOpts, account []common.Address, fromProfileId []*big.Int, toProfileId []*big.Int) (*ProfileLinkProfileIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkProfile", accountRule, fromProfileIdRule, toProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkProfileIterator{contract: _Profile.contract, event: "LinkProfile", logs: logs, sub: sub}, nil
}

// WatchLinkProfile is a free log subscription operation binding the contract event 0xbc914995d574dd9ef2df364e4eee2b85deda3ba35d054a62425fba1b97275716.
//
// Solidity: event LinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchLinkProfile(opts *bind.WatchOpts, sink chan<- *ProfileLinkProfile, account []common.Address, fromProfileId []*big.Int, toProfileId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkProfile", accountRule, fromProfileIdRule, toProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkProfile)
				if err := _Profile.contract.UnpackLog(event, "LinkProfile", log); err != nil {
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

// ParseLinkProfile is a log parse operation binding the contract event 0xbc914995d574dd9ef2df364e4eee2b85deda3ba35d054a62425fba1b97275716.
//
// Solidity: event LinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseLinkProfile(log types.Log) (*ProfileLinkProfile, error) {
	event := new(ProfileLinkProfile)
	if err := _Profile.contract.UnpackLog(event, "LinkProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinkProfileLinkIterator is returned from FilterLinkProfileLink and is used to iterate over the raw logs and unpacked data for LinkProfileLink events raised by the Profile contract.
type ProfileLinkProfileLinkIterator struct {
	Event *ProfileLinkProfileLink // Event containing the contract specifics and raw log

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
func (it *ProfileLinkProfileLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinkProfileLink)
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
		it.Event = new(ProfileLinkProfileLink)
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
func (it *ProfileLinkProfileLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinkProfileLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinkProfileLink represents a LinkProfileLink event raised by the Profile contract.
type ProfileLinkProfileLink struct {
	FromProfileId   *big.Int
	LinkType        [32]byte
	ClFromProfileId *big.Int
	ClToProfileId   *big.Int
	ClLinkType      [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkProfileLink is a free log retrieval operation binding the contract event 0xf0697acc3f194ceead566f2dffb676a7d4aaf14d0a2fad93923e778e4b40ae0d.
//
// Solidity: event LinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) FilterLinkProfileLink(opts *bind.FilterOpts, fromProfileId []*big.Int, linkType [][32]byte) (*ProfileLinkProfileLinkIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinkProfileLink", fromProfileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLinkProfileLinkIterator{contract: _Profile.contract, event: "LinkProfileLink", logs: logs, sub: sub}, nil
}

// WatchLinkProfileLink is a free log subscription operation binding the contract event 0xf0697acc3f194ceead566f2dffb676a7d4aaf14d0a2fad93923e778e4b40ae0d.
//
// Solidity: event LinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) WatchLinkProfileLink(opts *bind.WatchOpts, sink chan<- *ProfileLinkProfileLink, fromProfileId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinkProfileLink", fromProfileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinkProfileLink)
				if err := _Profile.contract.UnpackLog(event, "LinkProfileLink", log); err != nil {
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

// ParseLinkProfileLink is a log parse operation binding the contract event 0xf0697acc3f194ceead566f2dffb676a7d4aaf14d0a2fad93923e778e4b40ae0d.
//
// Solidity: event LinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) ParseLinkProfileLink(log types.Log) (*ProfileLinkProfileLink, error) {
	event := new(ProfileLinkProfileLink)
	if err := _Profile.contract.UnpackLog(event, "LinkProfileLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLinklistNFTInitializedIterator is returned from FilterLinklistNFTInitialized and is used to iterate over the raw logs and unpacked data for LinklistNFTInitialized events raised by the Profile contract.
type ProfileLinklistNFTInitializedIterator struct {
	Event *ProfileLinklistNFTInitialized // Event containing the contract specifics and raw log

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
func (it *ProfileLinklistNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLinklistNFTInitialized)
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
		it.Event = new(ProfileLinklistNFTInitialized)
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
func (it *ProfileLinklistNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLinklistNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLinklistNFTInitialized represents a LinklistNFTInitialized event raised by the Profile contract.
type ProfileLinklistNFTInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLinklistNFTInitialized is a free log retrieval operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) FilterLinklistNFTInitialized(opts *bind.FilterOpts) (*ProfileLinklistNFTInitializedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &ProfileLinklistNFTInitializedIterator{contract: _Profile.contract, event: "LinklistNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchLinklistNFTInitialized is a free log subscription operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) WatchLinklistNFTInitialized(opts *bind.WatchOpts, sink chan<- *ProfileLinklistNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLinklistNFTInitialized)
				if err := _Profile.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
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

// ParseLinklistNFTInitialized is a log parse operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) ParseLinklistNFTInitialized(log types.Log) (*ProfileLinklistNFTInitialized, error) {
	event := new(ProfileLinklistNFTInitialized)
	if err := _Profile.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileLockNoteIterator is returned from FilterLockNote and is used to iterate over the raw logs and unpacked data for LockNote events raised by the Profile contract.
type ProfileLockNoteIterator struct {
	Event *ProfileLockNote // Event containing the contract specifics and raw log

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
func (it *ProfileLockNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileLockNote)
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
		it.Event = new(ProfileLockNote)
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
func (it *ProfileLockNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileLockNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileLockNote represents a LockNote event raised by the Profile contract.
type ProfileLockNote struct {
	ProfileId *big.Int
	NoteId    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLockNote is a free log retrieval operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) FilterLockNote(opts *bind.FilterOpts, profileId []*big.Int) (*ProfileLockNoteIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "LockNote", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileLockNoteIterator{contract: _Profile.contract, event: "LockNote", logs: logs, sub: sub}, nil
}

// WatchLockNote is a free log subscription operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) WatchLockNote(opts *bind.WatchOpts, sink chan<- *ProfileLockNote, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "LockNote", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileLockNote)
				if err := _Profile.contract.UnpackLog(event, "LockNote", log); err != nil {
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

// ParseLockNote is a log parse operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed profileId, uint256 noteId)
func (_Profile *ProfileFilterer) ParseLockNote(log types.Log) (*ProfileLockNote, error) {
	event := new(ProfileLockNote)
	if err := _Profile.contract.UnpackLog(event, "LockNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileMintNFTInitializedIterator is returned from FilterMintNFTInitialized and is used to iterate over the raw logs and unpacked data for MintNFTInitialized events raised by the Profile contract.
type ProfileMintNFTInitializedIterator struct {
	Event *ProfileMintNFTInitialized // Event containing the contract specifics and raw log

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
func (it *ProfileMintNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileMintNFTInitialized)
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
		it.Event = new(ProfileMintNFTInitialized)
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
func (it *ProfileMintNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileMintNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileMintNFTInitialized represents a MintNFTInitialized event raised by the Profile contract.
type ProfileMintNFTInitialized struct {
	ProfileId *big.Int
	NoteId    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintNFTInitialized is a free log retrieval operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 profileId, uint256 noteId, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterMintNFTInitialized(opts *bind.FilterOpts) (*ProfileMintNFTInitializedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &ProfileMintNFTInitializedIterator{contract: _Profile.contract, event: "MintNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchMintNFTInitialized is a free log subscription operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 profileId, uint256 noteId, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchMintNFTInitialized(opts *bind.WatchOpts, sink chan<- *ProfileMintNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileMintNFTInitialized)
				if err := _Profile.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
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

// ParseMintNFTInitialized is a log parse operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 profileId, uint256 noteId, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseMintNFTInitialized(log types.Log) (*ProfileMintNFTInitialized, error) {
	event := new(ProfileMintNFTInitialized)
	if err := _Profile.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileMintNoteIterator is returned from FilterMintNote and is used to iterate over the raw logs and unpacked data for MintNote events raised by the Profile contract.
type ProfileMintNoteIterator struct {
	Event *ProfileMintNote // Event containing the contract specifics and raw log

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
func (it *ProfileMintNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileMintNote)
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
		it.Event = new(ProfileMintNote)
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
func (it *ProfileMintNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileMintNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileMintNote represents a MintNote event raised by the Profile contract.
type ProfileMintNote struct {
	To           common.Address
	ProfileId    *big.Int
	NoteId       *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintNote is a free log retrieval operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed profileId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Profile *ProfileFilterer) FilterMintNote(opts *bind.FilterOpts, to []common.Address, profileId []*big.Int, noteId []*big.Int) (*ProfileMintNoteIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "MintNote", toRule, profileIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileMintNoteIterator{contract: _Profile.contract, event: "MintNote", logs: logs, sub: sub}, nil
}

// WatchMintNote is a free log subscription operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed profileId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Profile *ProfileFilterer) WatchMintNote(opts *bind.WatchOpts, sink chan<- *ProfileMintNote, to []common.Address, profileId []*big.Int, noteId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "MintNote", toRule, profileIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileMintNote)
				if err := _Profile.contract.UnpackLog(event, "MintNote", log); err != nil {
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

// ParseMintNote is a log parse operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed profileId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Profile *ProfileFilterer) ParseMintNote(log types.Log) (*ProfileMintNote, error) {
	event := new(ProfileMintNote)
	if err := _Profile.contract.UnpackLog(event, "MintNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfilePostNoteIterator is returned from FilterPostNote and is used to iterate over the raw logs and unpacked data for PostNote events raised by the Profile contract.
type ProfilePostNoteIterator struct {
	Event *ProfilePostNote // Event containing the contract specifics and raw log

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
func (it *ProfilePostNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfilePostNote)
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
		it.Event = new(ProfilePostNote)
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
func (it *ProfilePostNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfilePostNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfilePostNote represents a PostNote event raised by the Profile contract.
type ProfilePostNote struct {
	ProfileId    *big.Int
	NoteId       *big.Int
	LinkKey      [32]byte
	LinkItemType [32]byte
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostNote is a free log retrieval operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed profileId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Profile *ProfileFilterer) FilterPostNote(opts *bind.FilterOpts, profileId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (*ProfilePostNoteIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "PostNote", profileIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return &ProfilePostNoteIterator{contract: _Profile.contract, event: "PostNote", logs: logs, sub: sub}, nil
}

// WatchPostNote is a free log subscription operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed profileId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Profile *ProfileFilterer) WatchPostNote(opts *bind.WatchOpts, sink chan<- *ProfilePostNote, profileId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "PostNote", profileIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfilePostNote)
				if err := _Profile.contract.UnpackLog(event, "PostNote", log); err != nil {
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

// ParsePostNote is a log parse operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed profileId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Profile *ProfileFilterer) ParsePostNote(log types.Log) (*ProfilePostNote, error) {
	event := new(ProfilePostNote)
	if err := _Profile.contract.UnpackLog(event, "PostNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileProfileCreatedIterator is returned from FilterProfileCreated and is used to iterate over the raw logs and unpacked data for ProfileCreated events raised by the Profile contract.
type ProfileProfileCreatedIterator struct {
	Event *ProfileProfileCreated // Event containing the contract specifics and raw log

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
func (it *ProfileProfileCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileProfileCreated)
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
		it.Event = new(ProfileProfileCreated)
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
func (it *ProfileProfileCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileProfileCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileProfileCreated represents a ProfileCreated event raised by the Profile contract.
type ProfileProfileCreated struct {
	ProfileId *big.Int
	Creator   common.Address
	To        common.Address
	Handle    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileCreated is a free log retrieval operation binding the contract event 0xa5802a04162552328d75eaac538a033704a7c3beab65d0a83e52da1c8c9b7cdf.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterProfileCreated(opts *bind.FilterOpts, profileId []*big.Int, creator []common.Address, to []common.Address) (*ProfileProfileCreatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ProfileProfileCreatedIterator{contract: _Profile.contract, event: "ProfileCreated", logs: logs, sub: sub}, nil
}

// WatchProfileCreated is a free log subscription operation binding the contract event 0xa5802a04162552328d75eaac538a033704a7c3beab65d0a83e52da1c8c9b7cdf.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchProfileCreated(opts *bind.WatchOpts, sink chan<- *ProfileProfileCreated, profileId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "ProfileCreated", profileIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileProfileCreated)
				if err := _Profile.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
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

// ParseProfileCreated is a log parse operation binding the contract event 0xa5802a04162552328d75eaac538a033704a7c3beab65d0a83e52da1c8c9b7cdf.
//
// Solidity: event ProfileCreated(uint256 indexed profileId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseProfileCreated(log types.Log) (*ProfileProfileCreated, error) {
	event := new(ProfileProfileCreated)
	if err := _Profile.contract.UnpackLog(event, "ProfileCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetHandleIterator is returned from FilterSetHandle and is used to iterate over the raw logs and unpacked data for SetHandle events raised by the Profile contract.
type ProfileSetHandleIterator struct {
	Event *ProfileSetHandle // Event containing the contract specifics and raw log

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
func (it *ProfileSetHandleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetHandle)
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
		it.Event = new(ProfileSetHandle)
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
func (it *ProfileSetHandleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetHandleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetHandle represents a SetHandle event raised by the Profile contract.
type ProfileSetHandle struct {
	Account   common.Address
	ProfileId *big.Int
	NewHandle string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetHandle is a free log retrieval operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed profileId, string newHandle)
func (_Profile *ProfileFilterer) FilterSetHandle(opts *bind.FilterOpts, account []common.Address, profileId []*big.Int) (*ProfileSetHandleIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetHandle", accountRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetHandleIterator{contract: _Profile.contract, event: "SetHandle", logs: logs, sub: sub}, nil
}

// WatchSetHandle is a free log subscription operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed profileId, string newHandle)
func (_Profile *ProfileFilterer) WatchSetHandle(opts *bind.WatchOpts, sink chan<- *ProfileSetHandle, account []common.Address, profileId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetHandle", accountRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetHandle)
				if err := _Profile.contract.UnpackLog(event, "SetHandle", log); err != nil {
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

// ParseSetHandle is a log parse operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed profileId, string newHandle)
func (_Profile *ProfileFilterer) ParseSetHandle(log types.Log) (*ProfileSetHandle, error) {
	event := new(ProfileSetHandle)
	if err := _Profile.contract.UnpackLog(event, "SetHandle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetLinkModule4AddressIterator is returned from FilterSetLinkModule4Address and is used to iterate over the raw logs and unpacked data for SetLinkModule4Address events raised by the Profile contract.
type ProfileSetLinkModule4AddressIterator struct {
	Event *ProfileSetLinkModule4Address // Event containing the contract specifics and raw log

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
func (it *ProfileSetLinkModule4AddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetLinkModule4Address)
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
		it.Event = new(ProfileSetLinkModule4Address)
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
func (it *ProfileSetLinkModule4AddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetLinkModule4AddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetLinkModule4Address represents a SetLinkModule4Address event raised by the Profile contract.
type ProfileSetLinkModule4Address struct {
	Account    common.Address
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Address is a free log retrieval operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetLinkModule4Address(opts *bind.FilterOpts, account []common.Address, linkModule []common.Address) (*ProfileSetLinkModule4AddressIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetLinkModule4AddressIterator{contract: _Profile.contract, event: "SetLinkModule4Address", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Address is a free log subscription operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetLinkModule4Address(opts *bind.WatchOpts, sink chan<- *ProfileSetLinkModule4Address, account []common.Address, linkModule []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetLinkModule4Address)
				if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
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

// ParseSetLinkModule4Address is a log parse operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetLinkModule4Address(log types.Log) (*ProfileSetLinkModule4Address, error) {
	event := new(ProfileSetLinkModule4Address)
	if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetLinkModule4ERC721Iterator is returned from FilterSetLinkModule4ERC721 and is used to iterate over the raw logs and unpacked data for SetLinkModule4ERC721 events raised by the Profile contract.
type ProfileSetLinkModule4ERC721Iterator struct {
	Event *ProfileSetLinkModule4ERC721 // Event containing the contract specifics and raw log

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
func (it *ProfileSetLinkModule4ERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetLinkModule4ERC721)
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
		it.Event = new(ProfileSetLinkModule4ERC721)
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
func (it *ProfileSetLinkModule4ERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetLinkModule4ERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetLinkModule4ERC721 represents a SetLinkModule4ERC721 event raised by the Profile contract.
type ProfileSetLinkModule4ERC721 struct {
	TokenAddress common.Address
	TokenId      *big.Int
	LinkModule   common.Address
	ReturnData   []byte
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4ERC721 is a free log retrieval operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetLinkModule4ERC721(opts *bind.FilterOpts, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (*ProfileSetLinkModule4ERC721Iterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetLinkModule4ERC721Iterator{contract: _Profile.contract, event: "SetLinkModule4ERC721", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4ERC721 is a free log subscription operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetLinkModule4ERC721(opts *bind.WatchOpts, sink chan<- *ProfileSetLinkModule4ERC721, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetLinkModule4ERC721)
				if err := _Profile.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
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

// ParseSetLinkModule4ERC721 is a log parse operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetLinkModule4ERC721(log types.Log) (*ProfileSetLinkModule4ERC721, error) {
	event := new(ProfileSetLinkModule4ERC721)
	if err := _Profile.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetLinkModule4LinklistIterator is returned from FilterSetLinkModule4Linklist and is used to iterate over the raw logs and unpacked data for SetLinkModule4Linklist events raised by the Profile contract.
type ProfileSetLinkModule4LinklistIterator struct {
	Event *ProfileSetLinkModule4Linklist // Event containing the contract specifics and raw log

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
func (it *ProfileSetLinkModule4LinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetLinkModule4Linklist)
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
		it.Event = new(ProfileSetLinkModule4Linklist)
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
func (it *ProfileSetLinkModule4LinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetLinkModule4LinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetLinkModule4Linklist represents a SetLinkModule4Linklist event raised by the Profile contract.
type ProfileSetLinkModule4Linklist struct {
	LinklistId *big.Int
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Linklist is a free log retrieval operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetLinkModule4Linklist(opts *bind.FilterOpts, linklistId []*big.Int, linkModule []common.Address) (*ProfileSetLinkModule4LinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetLinkModule4LinklistIterator{contract: _Profile.contract, event: "SetLinkModule4Linklist", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Linklist is a free log subscription operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetLinkModule4Linklist(opts *bind.WatchOpts, sink chan<- *ProfileSetLinkModule4Linklist, linklistId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetLinkModule4Linklist)
				if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
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

// ParseSetLinkModule4Linklist is a log parse operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetLinkModule4Linklist(log types.Log) (*ProfileSetLinkModule4Linklist, error) {
	event := new(ProfileSetLinkModule4Linklist)
	if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetLinkModule4NoteIterator is returned from FilterSetLinkModule4Note and is used to iterate over the raw logs and unpacked data for SetLinkModule4Note events raised by the Profile contract.
type ProfileSetLinkModule4NoteIterator struct {
	Event *ProfileSetLinkModule4Note // Event containing the contract specifics and raw log

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
func (it *ProfileSetLinkModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetLinkModule4Note)
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
		it.Event = new(ProfileSetLinkModule4Note)
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
func (it *ProfileSetLinkModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetLinkModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetLinkModule4Note represents a SetLinkModule4Note event raised by the Profile contract.
type ProfileSetLinkModule4Note struct {
	ProfileId  *big.Int
	NoteId     *big.Int
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Note is a free log retrieval operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetLinkModule4Note(opts *bind.FilterOpts, profileId []*big.Int, noteId []*big.Int, linkModule []common.Address) (*ProfileSetLinkModule4NoteIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetLinkModule4Note", profileIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetLinkModule4NoteIterator{contract: _Profile.contract, event: "SetLinkModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Note is a free log subscription operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetLinkModule4Note(opts *bind.WatchOpts, sink chan<- *ProfileSetLinkModule4Note, profileId []*big.Int, noteId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetLinkModule4Note", profileIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetLinkModule4Note)
				if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
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

// ParseSetLinkModule4Note is a log parse operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetLinkModule4Note(log types.Log) (*ProfileSetLinkModule4Note, error) {
	event := new(ProfileSetLinkModule4Note)
	if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetLinkModule4ProfileIterator is returned from FilterSetLinkModule4Profile and is used to iterate over the raw logs and unpacked data for SetLinkModule4Profile events raised by the Profile contract.
type ProfileSetLinkModule4ProfileIterator struct {
	Event *ProfileSetLinkModule4Profile // Event containing the contract specifics and raw log

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
func (it *ProfileSetLinkModule4ProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetLinkModule4Profile)
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
		it.Event = new(ProfileSetLinkModule4Profile)
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
func (it *ProfileSetLinkModule4ProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetLinkModule4ProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetLinkModule4Profile represents a SetLinkModule4Profile event raised by the Profile contract.
type ProfileSetLinkModule4Profile struct {
	ProfileId  *big.Int
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Profile is a free log retrieval operation binding the contract event 0x2852faa26318de2c5b2d5e319f7b3c29f0cd77f50c67559a4082bf5763c4d59f.
//
// Solidity: event SetLinkModule4Profile(uint256 indexed profileId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetLinkModule4Profile(opts *bind.FilterOpts, profileId []*big.Int, linkModule []common.Address) (*ProfileSetLinkModule4ProfileIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetLinkModule4Profile", profileIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetLinkModule4ProfileIterator{contract: _Profile.contract, event: "SetLinkModule4Profile", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Profile is a free log subscription operation binding the contract event 0x2852faa26318de2c5b2d5e319f7b3c29f0cd77f50c67559a4082bf5763c4d59f.
//
// Solidity: event SetLinkModule4Profile(uint256 indexed profileId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetLinkModule4Profile(opts *bind.WatchOpts, sink chan<- *ProfileSetLinkModule4Profile, profileId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetLinkModule4Profile", profileIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetLinkModule4Profile)
				if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Profile", log); err != nil {
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

// ParseSetLinkModule4Profile is a log parse operation binding the contract event 0x2852faa26318de2c5b2d5e319f7b3c29f0cd77f50c67559a4082bf5763c4d59f.
//
// Solidity: event SetLinkModule4Profile(uint256 indexed profileId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetLinkModule4Profile(log types.Log) (*ProfileSetLinkModule4Profile, error) {
	event := new(ProfileSetLinkModule4Profile)
	if err := _Profile.contract.UnpackLog(event, "SetLinkModule4Profile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetMintModule4NoteIterator is returned from FilterSetMintModule4Note and is used to iterate over the raw logs and unpacked data for SetMintModule4Note events raised by the Profile contract.
type ProfileSetMintModule4NoteIterator struct {
	Event *ProfileSetMintModule4Note // Event containing the contract specifics and raw log

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
func (it *ProfileSetMintModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetMintModule4Note)
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
		it.Event = new(ProfileSetMintModule4Note)
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
func (it *ProfileSetMintModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetMintModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetMintModule4Note represents a SetMintModule4Note event raised by the Profile contract.
type ProfileSetMintModule4Note struct {
	ProfileId  *big.Int
	NoteId     *big.Int
	MintModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMintModule4Note is a free log retrieval operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetMintModule4Note(opts *bind.FilterOpts, profileId []*big.Int, noteId []*big.Int, mintModule []common.Address) (*ProfileSetMintModule4NoteIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetMintModule4Note", profileIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetMintModule4NoteIterator{contract: _Profile.contract, event: "SetMintModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetMintModule4Note is a free log subscription operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetMintModule4Note(opts *bind.WatchOpts, sink chan<- *ProfileSetMintModule4Note, profileId []*big.Int, noteId []*big.Int, mintModule []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetMintModule4Note", profileIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetMintModule4Note)
				if err := _Profile.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
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

// ParseSetMintModule4Note is a log parse operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed profileId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetMintModule4Note(log types.Log) (*ProfileSetMintModule4Note, error) {
	event := new(ProfileSetMintModule4Note)
	if err := _Profile.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetNoteUriIterator is returned from FilterSetNoteUri and is used to iterate over the raw logs and unpacked data for SetNoteUri events raised by the Profile contract.
type ProfileSetNoteUriIterator struct {
	Event *ProfileSetNoteUri // Event containing the contract specifics and raw log

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
func (it *ProfileSetNoteUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetNoteUri)
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
		it.Event = new(ProfileSetNoteUri)
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
func (it *ProfileSetNoteUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetNoteUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetNoteUri represents a SetNoteUri event raised by the Profile contract.
type ProfileSetNoteUri struct {
	ProfileId *big.Int
	NoteId    *big.Int
	NewUri    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetNoteUri is a free log retrieval operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed profileId, uint256 noteId, string newUri)
func (_Profile *ProfileFilterer) FilterSetNoteUri(opts *bind.FilterOpts, profileId []*big.Int) (*ProfileSetNoteUriIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetNoteUri", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetNoteUriIterator{contract: _Profile.contract, event: "SetNoteUri", logs: logs, sub: sub}, nil
}

// WatchSetNoteUri is a free log subscription operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed profileId, uint256 noteId, string newUri)
func (_Profile *ProfileFilterer) WatchSetNoteUri(opts *bind.WatchOpts, sink chan<- *ProfileSetNoteUri, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetNoteUri", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetNoteUri)
				if err := _Profile.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
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

// ParseSetNoteUri is a log parse operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed profileId, uint256 noteId, string newUri)
func (_Profile *ProfileFilterer) ParseSetNoteUri(log types.Log) (*ProfileSetNoteUri, error) {
	event := new(ProfileSetNoteUri)
	if err := _Profile.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the Profile contract.
type ProfileSetOperatorIterator struct {
	Event *ProfileSetOperator // Event containing the contract specifics and raw log

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
func (it *ProfileSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetOperator)
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
		it.Event = new(ProfileSetOperator)
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
func (it *ProfileSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetOperator represents a SetOperator event raised by the Profile contract.
type ProfileSetOperator struct {
	ProfileId *big.Int
	Operator  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed profileId, address indexed operator, uint256 timestamp)
func (_Profile *ProfileFilterer) FilterSetOperator(opts *bind.FilterOpts, profileId []*big.Int, operator []common.Address) (*ProfileSetOperatorIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetOperator", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetOperatorIterator{contract: _Profile.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed profileId, address indexed operator, uint256 timestamp)
func (_Profile *ProfileFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *ProfileSetOperator, profileId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetOperator", profileIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetOperator)
				if err := _Profile.contract.UnpackLog(event, "SetOperator", log); err != nil {
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

// ParseSetOperator is a log parse operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed profileId, address indexed operator, uint256 timestamp)
func (_Profile *ProfileFilterer) ParseSetOperator(log types.Log) (*ProfileSetOperator, error) {
	event := new(ProfileSetOperator)
	if err := _Profile.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetPrimaryProfileIdIterator is returned from FilterSetPrimaryProfileId and is used to iterate over the raw logs and unpacked data for SetPrimaryProfileId events raised by the Profile contract.
type ProfileSetPrimaryProfileIdIterator struct {
	Event *ProfileSetPrimaryProfileId // Event containing the contract specifics and raw log

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
func (it *ProfileSetPrimaryProfileIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetPrimaryProfileId)
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
		it.Event = new(ProfileSetPrimaryProfileId)
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
func (it *ProfileSetPrimaryProfileIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetPrimaryProfileIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetPrimaryProfileId represents a SetPrimaryProfileId event raised by the Profile contract.
type ProfileSetPrimaryProfileId struct {
	Account      common.Address
	ProfileId    *big.Int
	OldProfileId *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetPrimaryProfileId is a free log retrieval operation binding the contract event 0xf12967ee1b3d3964e16093b221f5239f61fe0191c5ce5d5f543ae891d78f6603.
//
// Solidity: event SetPrimaryProfileId(address indexed account, uint256 indexed profileId, uint256 indexed oldProfileId)
func (_Profile *ProfileFilterer) FilterSetPrimaryProfileId(opts *bind.FilterOpts, account []common.Address, profileId []*big.Int, oldProfileId []*big.Int) (*ProfileSetPrimaryProfileIdIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var oldProfileIdRule []interface{}
	for _, oldProfileIdItem := range oldProfileId {
		oldProfileIdRule = append(oldProfileIdRule, oldProfileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetPrimaryProfileId", accountRule, profileIdRule, oldProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetPrimaryProfileIdIterator{contract: _Profile.contract, event: "SetPrimaryProfileId", logs: logs, sub: sub}, nil
}

// WatchSetPrimaryProfileId is a free log subscription operation binding the contract event 0xf12967ee1b3d3964e16093b221f5239f61fe0191c5ce5d5f543ae891d78f6603.
//
// Solidity: event SetPrimaryProfileId(address indexed account, uint256 indexed profileId, uint256 indexed oldProfileId)
func (_Profile *ProfileFilterer) WatchSetPrimaryProfileId(opts *bind.WatchOpts, sink chan<- *ProfileSetPrimaryProfileId, account []common.Address, profileId []*big.Int, oldProfileId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var oldProfileIdRule []interface{}
	for _, oldProfileIdItem := range oldProfileId {
		oldProfileIdRule = append(oldProfileIdRule, oldProfileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetPrimaryProfileId", accountRule, profileIdRule, oldProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetPrimaryProfileId)
				if err := _Profile.contract.UnpackLog(event, "SetPrimaryProfileId", log); err != nil {
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

// ParseSetPrimaryProfileId is a log parse operation binding the contract event 0xf12967ee1b3d3964e16093b221f5239f61fe0191c5ce5d5f543ae891d78f6603.
//
// Solidity: event SetPrimaryProfileId(address indexed account, uint256 indexed profileId, uint256 indexed oldProfileId)
func (_Profile *ProfileFilterer) ParseSetPrimaryProfileId(log types.Log) (*ProfileSetPrimaryProfileId, error) {
	event := new(ProfileSetPrimaryProfileId)
	if err := _Profile.contract.UnpackLog(event, "SetPrimaryProfileId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetProfileUriIterator is returned from FilterSetProfileUri and is used to iterate over the raw logs and unpacked data for SetProfileUri events raised by the Profile contract.
type ProfileSetProfileUriIterator struct {
	Event *ProfileSetProfileUri // Event containing the contract specifics and raw log

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
func (it *ProfileSetProfileUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetProfileUri)
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
		it.Event = new(ProfileSetProfileUri)
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
func (it *ProfileSetProfileUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetProfileUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetProfileUri represents a SetProfileUri event raised by the Profile contract.
type ProfileSetProfileUri struct {
	ProfileId *big.Int
	NewUri    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetProfileUri is a free log retrieval operation binding the contract event 0xc6b6b2c87fb0784186dd03398c7203c866d0ae59539fa3158aecbc86cb587a95.
//
// Solidity: event SetProfileUri(uint256 indexed profileId, string newUri)
func (_Profile *ProfileFilterer) FilterSetProfileUri(opts *bind.FilterOpts, profileId []*big.Int) (*ProfileSetProfileUriIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetProfileUri", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetProfileUriIterator{contract: _Profile.contract, event: "SetProfileUri", logs: logs, sub: sub}, nil
}

// WatchSetProfileUri is a free log subscription operation binding the contract event 0xc6b6b2c87fb0784186dd03398c7203c866d0ae59539fa3158aecbc86cb587a95.
//
// Solidity: event SetProfileUri(uint256 indexed profileId, string newUri)
func (_Profile *ProfileFilterer) WatchSetProfileUri(opts *bind.WatchOpts, sink chan<- *ProfileSetProfileUri, profileId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetProfileUri", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetProfileUri)
				if err := _Profile.contract.UnpackLog(event, "SetProfileUri", log); err != nil {
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

// ParseSetProfileUri is a log parse operation binding the contract event 0xc6b6b2c87fb0784186dd03398c7203c866d0ae59539fa3158aecbc86cb587a95.
//
// Solidity: event SetProfileUri(uint256 indexed profileId, string newUri)
func (_Profile *ProfileFilterer) ParseSetProfileUri(log types.Log) (*ProfileSetProfileUri, error) {
	event := new(ProfileSetProfileUri)
	if err := _Profile.contract.UnpackLog(event, "SetProfileUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileSetSocialTokenIterator is returned from FilterSetSocialToken and is used to iterate over the raw logs and unpacked data for SetSocialToken events raised by the Profile contract.
type ProfileSetSocialTokenIterator struct {
	Event *ProfileSetSocialToken // Event containing the contract specifics and raw log

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
func (it *ProfileSetSocialTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileSetSocialToken)
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
		it.Event = new(ProfileSetSocialToken)
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
func (it *ProfileSetSocialTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileSetSocialTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileSetSocialToken represents a SetSocialToken event raised by the Profile contract.
type ProfileSetSocialToken struct {
	Account      common.Address
	ProfileId    *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetSocialToken is a free log retrieval operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed profileId, address indexed tokenAddress)
func (_Profile *ProfileFilterer) FilterSetSocialToken(opts *bind.FilterOpts, account []common.Address, profileId []*big.Int, tokenAddress []common.Address) (*ProfileSetSocialTokenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "SetSocialToken", accountRule, profileIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProfileSetSocialTokenIterator{contract: _Profile.contract, event: "SetSocialToken", logs: logs, sub: sub}, nil
}

// WatchSetSocialToken is a free log subscription operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed profileId, address indexed tokenAddress)
func (_Profile *ProfileFilterer) WatchSetSocialToken(opts *bind.WatchOpts, sink chan<- *ProfileSetSocialToken, account []common.Address, profileId []*big.Int, tokenAddress []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "SetSocialToken", accountRule, profileIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileSetSocialToken)
				if err := _Profile.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
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

// ParseSetSocialToken is a log parse operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed profileId, address indexed tokenAddress)
func (_Profile *ProfileFilterer) ParseSetSocialToken(log types.Log) (*ProfileSetSocialToken, error) {
	event := new(ProfileSetSocialToken)
	if err := _Profile.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Profile contract.
type ProfileTransferIterator struct {
	Event *ProfileTransfer // Event containing the contract specifics and raw log

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
func (it *ProfileTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileTransfer)
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
		it.Event = new(ProfileTransfer)
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
func (it *ProfileTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileTransfer represents a Transfer event raised by the Profile contract.
type ProfileTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Profile *ProfileFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ProfileTransferIterator, error) {

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

	logs, sub, err := _Profile.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileTransferIterator{contract: _Profile.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Profile *ProfileFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProfileTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Profile.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileTransfer)
				if err := _Profile.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Profile *ProfileFilterer) ParseTransfer(log types.Log) (*ProfileTransfer, error) {
	event := new(ProfileTransfer)
	if err := _Profile.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkAddressIterator is returned from FilterUnlinkAddress and is used to iterate over the raw logs and unpacked data for UnlinkAddress events raised by the Profile contract.
type ProfileUnlinkAddressIterator struct {
	Event *ProfileUnlinkAddress // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkAddress)
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
		it.Event = new(ProfileUnlinkAddress)
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
func (it *ProfileUnlinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkAddress represents a UnlinkAddress event raised by the Profile contract.
type ProfileUnlinkAddress struct {
	FromProfileId *big.Int
	EthAddress    common.Address
	LinkType      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAddress is a free log retrieval operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType)
func (_Profile *ProfileFilterer) FilterUnlinkAddress(opts *bind.FilterOpts, fromProfileId []*big.Int, ethAddress []common.Address) (*ProfileUnlinkAddressIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkAddress", fromProfileIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkAddressIterator{contract: _Profile.contract, event: "UnlinkAddress", logs: logs, sub: sub}, nil
}

// WatchUnlinkAddress is a free log subscription operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType)
func (_Profile *ProfileFilterer) WatchUnlinkAddress(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkAddress, fromProfileId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkAddress", fromProfileIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkAddress)
				if err := _Profile.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
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

// ParseUnlinkAddress is a log parse operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromProfileId, address indexed ethAddress, bytes32 linkType)
func (_Profile *ProfileFilterer) ParseUnlinkAddress(log types.Log) (*ProfileUnlinkAddress, error) {
	event := new(ProfileUnlinkAddress)
	if err := _Profile.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkAnyUriIterator is returned from FilterUnlinkAnyUri and is used to iterate over the raw logs and unpacked data for UnlinkAnyUri events raised by the Profile contract.
type ProfileUnlinkAnyUriIterator struct {
	Event *ProfileUnlinkAnyUri // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkAnyUri)
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
		it.Event = new(ProfileUnlinkAnyUri)
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
func (it *ProfileUnlinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkAnyUri represents a UnlinkAnyUri event raised by the Profile contract.
type ProfileUnlinkAnyUri struct {
	FromProfileId *big.Int
	ToUri         string
	LinkType      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAnyUri is a free log retrieval operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType)
func (_Profile *ProfileFilterer) FilterUnlinkAnyUri(opts *bind.FilterOpts, fromProfileId []*big.Int) (*ProfileUnlinkAnyUriIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkAnyUri", fromProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkAnyUriIterator{contract: _Profile.contract, event: "UnlinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchUnlinkAnyUri is a free log subscription operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType)
func (_Profile *ProfileFilterer) WatchUnlinkAnyUri(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkAnyUri, fromProfileId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkAnyUri", fromProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkAnyUri)
				if err := _Profile.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
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

// ParseUnlinkAnyUri is a log parse operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromProfileId, string toUri, bytes32 linkType)
func (_Profile *ProfileFilterer) ParseUnlinkAnyUri(log types.Log) (*ProfileUnlinkAnyUri, error) {
	event := new(ProfileUnlinkAnyUri)
	if err := _Profile.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkERC721Iterator is returned from FilterUnlinkERC721 and is used to iterate over the raw logs and unpacked data for UnlinkERC721 events raised by the Profile contract.
type ProfileUnlinkERC721Iterator struct {
	Event *ProfileUnlinkERC721 // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkERC721)
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
		it.Event = new(ProfileUnlinkERC721)
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
func (it *ProfileUnlinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkERC721 represents a UnlinkERC721 event raised by the Profile contract.
type ProfileUnlinkERC721 struct {
	FromProfileId *big.Int
	TokenAddress  common.Address
	ToNoteId      *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkERC721 is a free log retrieval operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterUnlinkERC721(opts *bind.FilterOpts, fromProfileId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*ProfileUnlinkERC721Iterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkERC721", fromProfileIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkERC721Iterator{contract: _Profile.contract, event: "UnlinkERC721", logs: logs, sub: sub}, nil
}

// WatchUnlinkERC721 is a free log subscription operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchUnlinkERC721(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkERC721, fromProfileId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkERC721", fromProfileIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkERC721)
				if err := _Profile.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
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

// ParseUnlinkERC721 is a log parse operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromProfileId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseUnlinkERC721(log types.Log) (*ProfileUnlinkERC721, error) {
	event := new(ProfileUnlinkERC721)
	if err := _Profile.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkLinklistIterator is returned from FilterUnlinkLinklist and is used to iterate over the raw logs and unpacked data for UnlinkLinklist events raised by the Profile contract.
type ProfileUnlinkLinklistIterator struct {
	Event *ProfileUnlinkLinklist // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkLinklist)
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
		it.Event = new(ProfileUnlinkLinklist)
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
func (it *ProfileUnlinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkLinklist represents a UnlinkLinklist event raised by the Profile contract.
type ProfileUnlinkLinklist struct {
	FromProfileId *big.Int
	ToLinklistId  *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkLinklist is a free log retrieval operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) FilterUnlinkLinklist(opts *bind.FilterOpts, fromProfileId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*ProfileUnlinkLinklistIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkLinklist", fromProfileIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkLinklistIterator{contract: _Profile.contract, event: "UnlinkLinklist", logs: logs, sub: sub}, nil
}

// WatchUnlinkLinklist is a free log subscription operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) WatchUnlinkLinklist(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkLinklist, fromProfileId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkLinklist", fromProfileIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkLinklist)
				if err := _Profile.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
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

// ParseUnlinkLinklist is a log parse operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromProfileId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Profile *ProfileFilterer) ParseUnlinkLinklist(log types.Log) (*ProfileUnlinkLinklist, error) {
	event := new(ProfileUnlinkLinklist)
	if err := _Profile.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkNoteIterator is returned from FilterUnlinkNote and is used to iterate over the raw logs and unpacked data for UnlinkNote events raised by the Profile contract.
type ProfileUnlinkNoteIterator struct {
	Event *ProfileUnlinkNote // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkNote)
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
		it.Event = new(ProfileUnlinkNote)
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
func (it *ProfileUnlinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkNote represents a UnlinkNote event raised by the Profile contract.
type ProfileUnlinkNote struct {
	FromProfileId *big.Int
	ToProfileId   *big.Int
	ToNoteId      *big.Int
	LinkType      [32]byte
	LinklistId    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkNote is a free log retrieval operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) FilterUnlinkNote(opts *bind.FilterOpts, fromProfileId []*big.Int, toProfileId []*big.Int, toNoteId []*big.Int) (*ProfileUnlinkNoteIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkNote", fromProfileIdRule, toProfileIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkNoteIterator{contract: _Profile.contract, event: "UnlinkNote", logs: logs, sub: sub}, nil
}

// WatchUnlinkNote is a free log subscription operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) WatchUnlinkNote(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkNote, fromProfileId []*big.Int, toProfileId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkNote", fromProfileIdRule, toProfileIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkNote)
				if err := _Profile.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
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

// ParseUnlinkNote is a log parse operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromProfileId, uint256 indexed toProfileId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Profile *ProfileFilterer) ParseUnlinkNote(log types.Log) (*ProfileUnlinkNote, error) {
	event := new(ProfileUnlinkNote)
	if err := _Profile.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkProfileIterator is returned from FilterUnlinkProfile and is used to iterate over the raw logs and unpacked data for UnlinkProfile events raised by the Profile contract.
type ProfileUnlinkProfileIterator struct {
	Event *ProfileUnlinkProfile // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkProfileIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkProfile)
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
		it.Event = new(ProfileUnlinkProfile)
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
func (it *ProfileUnlinkProfileIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkProfileIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkProfile represents a UnlinkProfile event raised by the Profile contract.
type ProfileUnlinkProfile struct {
	Account       common.Address
	FromProfileId *big.Int
	ToProfileId   *big.Int
	LinkType      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnlinkProfile is a free log retrieval operation binding the contract event 0x08a14b775e42bf92e94f9d72aba57ad213485e0c20e373c9f5dd161fb81890a2.
//
// Solidity: event UnlinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType)
func (_Profile *ProfileFilterer) FilterUnlinkProfile(opts *bind.FilterOpts, account []common.Address, fromProfileId []*big.Int, toProfileId []*big.Int) (*ProfileUnlinkProfileIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkProfile", accountRule, fromProfileIdRule, toProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkProfileIterator{contract: _Profile.contract, event: "UnlinkProfile", logs: logs, sub: sub}, nil
}

// WatchUnlinkProfile is a free log subscription operation binding the contract event 0x08a14b775e42bf92e94f9d72aba57ad213485e0c20e373c9f5dd161fb81890a2.
//
// Solidity: event UnlinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType)
func (_Profile *ProfileFilterer) WatchUnlinkProfile(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkProfile, account []common.Address, fromProfileId []*big.Int, toProfileId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var toProfileIdRule []interface{}
	for _, toProfileIdItem := range toProfileId {
		toProfileIdRule = append(toProfileIdRule, toProfileIdItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkProfile", accountRule, fromProfileIdRule, toProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkProfile)
				if err := _Profile.contract.UnpackLog(event, "UnlinkProfile", log); err != nil {
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

// ParseUnlinkProfile is a log parse operation binding the contract event 0x08a14b775e42bf92e94f9d72aba57ad213485e0c20e373c9f5dd161fb81890a2.
//
// Solidity: event UnlinkProfile(address indexed account, uint256 indexed fromProfileId, uint256 indexed toProfileId, bytes32 linkType)
func (_Profile *ProfileFilterer) ParseUnlinkProfile(log types.Log) (*ProfileUnlinkProfile, error) {
	event := new(ProfileUnlinkProfile)
	if err := _Profile.contract.UnpackLog(event, "UnlinkProfile", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileUnlinkProfileLinkIterator is returned from FilterUnlinkProfileLink and is used to iterate over the raw logs and unpacked data for UnlinkProfileLink events raised by the Profile contract.
type ProfileUnlinkProfileLinkIterator struct {
	Event *ProfileUnlinkProfileLink // Event containing the contract specifics and raw log

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
func (it *ProfileUnlinkProfileLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileUnlinkProfileLink)
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
		it.Event = new(ProfileUnlinkProfileLink)
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
func (it *ProfileUnlinkProfileLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileUnlinkProfileLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileUnlinkProfileLink represents a UnlinkProfileLink event raised by the Profile contract.
type ProfileUnlinkProfileLink struct {
	FromProfileId    *big.Int
	LinkType         [32]byte
	ClFromProfileeId *big.Int
	ClToProfileId    *big.Int
	ClLinkType       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUnlinkProfileLink is a free log retrieval operation binding the contract event 0x90fe9f3bf41416d9a567460a40443ef7aea47220fa342bc0e07fd01c93ffc7e1.
//
// Solidity: event UnlinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileeId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) FilterUnlinkProfileLink(opts *bind.FilterOpts, fromProfileId []*big.Int, linkType [][32]byte) (*ProfileUnlinkProfileLinkIterator, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.FilterLogs(opts, "UnlinkProfileLink", fromProfileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &ProfileUnlinkProfileLinkIterator{contract: _Profile.contract, event: "UnlinkProfileLink", logs: logs, sub: sub}, nil
}

// WatchUnlinkProfileLink is a free log subscription operation binding the contract event 0x90fe9f3bf41416d9a567460a40443ef7aea47220fa342bc0e07fd01c93ffc7e1.
//
// Solidity: event UnlinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileeId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) WatchUnlinkProfileLink(opts *bind.WatchOpts, sink chan<- *ProfileUnlinkProfileLink, fromProfileId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromProfileIdRule []interface{}
	for _, fromProfileIdItem := range fromProfileId {
		fromProfileIdRule = append(fromProfileIdRule, fromProfileIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Profile.contract.WatchLogs(opts, "UnlinkProfileLink", fromProfileIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileUnlinkProfileLink)
				if err := _Profile.contract.UnpackLog(event, "UnlinkProfileLink", log); err != nil {
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

// ParseUnlinkProfileLink is a log parse operation binding the contract event 0x90fe9f3bf41416d9a567460a40443ef7aea47220fa342bc0e07fd01c93ffc7e1.
//
// Solidity: event UnlinkProfileLink(uint256 indexed fromProfileId, bytes32 indexed linkType, uint256 clFromProfileeId, uint256 clToProfileId, bytes32 clLinkType)
func (_Profile *ProfileFilterer) ParseUnlinkProfileLink(log types.Log) (*ProfileUnlinkProfileLink, error) {
	event := new(ProfileUnlinkProfileLink)
	if err := _Profile.contract.UnpackLog(event, "UnlinkProfileLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProfileWeb3EntryInitializedIterator is returned from FilterWeb3EntryInitialized and is used to iterate over the raw logs and unpacked data for Web3EntryInitialized events raised by the Profile contract.
type ProfileWeb3EntryInitializedIterator struct {
	Event *ProfileWeb3EntryInitialized // Event containing the contract specifics and raw log

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
func (it *ProfileWeb3EntryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProfileWeb3EntryInitialized)
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
		it.Event = new(ProfileWeb3EntryInitialized)
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
func (it *ProfileWeb3EntryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProfileWeb3EntryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProfileWeb3EntryInitialized represents a Web3EntryInitialized event raised by the Profile contract.
type ProfileWeb3EntryInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWeb3EntryInitialized is a free log retrieval operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) FilterWeb3EntryInitialized(opts *bind.FilterOpts) (*ProfileWeb3EntryInitializedIterator, error) {

	logs, sub, err := _Profile.contract.FilterLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return &ProfileWeb3EntryInitializedIterator{contract: _Profile.contract, event: "Web3EntryInitialized", logs: logs, sub: sub}, nil
}

// WatchWeb3EntryInitialized is a free log subscription operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) WatchWeb3EntryInitialized(opts *bind.WatchOpts, sink chan<- *ProfileWeb3EntryInitialized) (event.Subscription, error) {

	logs, sub, err := _Profile.contract.WatchLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProfileWeb3EntryInitialized)
				if err := _Profile.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
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

// ParseWeb3EntryInitialized is a log parse operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Profile *ProfileFilterer) ParseWeb3EntryInitialized(log types.Log) (*ProfileWeb3EntryInitialized, error) {
	event := new(ProfileWeb3EntryInitialized)
	if err := _Profile.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
