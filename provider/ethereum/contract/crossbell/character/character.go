// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package character

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

// DataTypesCharacter is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCharacter struct {
	CharacterId *big.Int
	Handle      string
	Uri         string
	NoteCount   *big.Int
	SocialToken common.Address
	LinkModule  common.Address
}

// DataTypesCreateCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCreateCharacterData struct {
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
	CharacterId    *big.Int
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
	CharacterId *big.Int
	NoteId      *big.Int
}

// DataTypesPostNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesPostNoteData struct {
	CharacterId        *big.Int
	ContentUri         string
	LinkModule         common.Address
	LinkModuleInitData []byte
	MintModule         common.Address
	MintModuleInitData []byte
	Locked             bool
}

// DataTypescreateThenLinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypescreateThenLinkCharacterData struct {
	FromCharacterId *big.Int
	To              common.Address
	LinkType        [32]byte
}

// DataTypeslinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypeslinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypeslinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	Data            []byte
}

// DataTypessetLinkModule4AddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4AddressData struct {
	Account            common.Address
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetLinkModule4CharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetLinkModule4CharacterData struct {
	CharacterId        *big.Int
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
	CharacterId        *big.Int
	NoteId             *big.Int
	LinkModule         common.Address
	LinkModuleInitData []byte
}

// DataTypessetMintModule4NoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypessetMintModule4NoteData struct {
	CharacterId        *big.Int
	NoteId             *big.Int
	MintModule         common.Address
	MintModuleInitData []byte
}

// DataTypesunlinkAddressData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAddressData struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
}

// DataTypesunlinkAnyUriData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkAnyUriData struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
}

// DataTypesunlinkCharacterData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkCharacterData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkERC721Data is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkERC721Data struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	TokenId         *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkLinklistData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkLinklistData struct {
	FromCharacterId *big.Int
	ToLinkListId    *big.Int
	LinkType        [32]byte
}

// DataTypesunlinkNoteData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesunlinkNoteData struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
}

// CharacterMetaData contains all meta data concerning the Character contract.
var CharacterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canCreate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.CreateCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.createThenLinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createThenLinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"deleteNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacter\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getCharacterByHandle\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"noteCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"socialToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.Character\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getCharacterUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getLinkModule4Address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4ERC721\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinkModule4Linklist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLinklistContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"getLinklistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linkListId\",\"type\":\"uint256\"}],\"name\":\"getLinklistType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getLinklistUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"getNote\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mintNFT\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"deleted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.Note\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPrimaryCharacterId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevision\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_linklistContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mintNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_periphery\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_resolver\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"isPrimaryCharacter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.linkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"linkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"lockNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"migrateNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.MintNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mintNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postNote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"}],\"name\":\"postNote4Address\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"postNote4AnyUri\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"}],\"name\":\"postNote4Character\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ERC721Struct\",\"name\":\"erc721\",\"type\":\"tuple\"}],\"name\":\"postNote4ERC721\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"noteData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"}],\"name\":\"postNote4Linklist\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"}],\"internalType\":\"structDataTypes.PostNoteData\",\"name\":\"postNoteData\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NoteStruct\",\"name\":\"note\",\"type\":\"tuple\"}],\"name\":\"postNote4Note\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resolver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setCharacterUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"setHandle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4AddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Address\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4CharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Character\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4ERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4ERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4LinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Linklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"linkModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setLinkModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setLinkModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"setLinklistUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"mintModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structDataTypes.setMintModule4NoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setMintModule4Note\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setNoteUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryCharacterId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"setProfileUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"setSocialToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAddressData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkAnyUriData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkAnyUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkCharacterData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkCharacter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkERC721Data\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toLinkListId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkLinklistData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkLinklist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"internalType\":\"structDataTypes.unlinkNoteData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"unlinkNote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"AttachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BaseInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CharacterCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"DeleteNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"DetachLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkCharacter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"LinkCharacterLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"LinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LinklistNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"}],\"name\":\"LockNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MintNFTInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MintNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkItemType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"PostNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetCharacterUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newHandle\",\"type\":\"string\"}],\"name\":\"SetHandle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Address\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Character\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4ERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Linklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"linkModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetLinkModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetMintModule4Note\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newUri\",\"type\":\"string\"}],\"name\":\"SetNoteUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldCharacterId\",\"type\":\"uint256\"}],\"name\":\"SetPrimaryCharacterId\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"characterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"SetSocialToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ethAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkAnyUri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkCharacter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clFromCharactereId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"clToCharacterId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"clLinkType\",\"type\":\"bytes32\"}],\"name\":\"UnlinkCharacterLink\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkERC721\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toLinklistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkLinklist\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"fromCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toCharacterId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"toNoteId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"linkType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"linklistId\",\"type\":\"uint256\"}],\"name\":\"UnlinkNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Web3EntryInitialized\",\"type\":\"event\"}]",
}

// CharacterABI is the input ABI used to generate the binding from.
// Deprecated: Use CharacterMetaData.ABI instead.
var CharacterABI = CharacterMetaData.ABI

// Character is an auto generated Go binding around an Ethereum contract.
type Character struct {
	CharacterCaller     // Read-only binding to the contract
	CharacterTransactor // Write-only binding to the contract
	CharacterFilterer   // Log filterer for contract events
}

// CharacterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CharacterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CharacterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CharacterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharacterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CharacterSession struct {
	Contract     *Character        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CharacterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CharacterCallerSession struct {
	Contract *CharacterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CharacterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CharacterTransactorSession struct {
	Contract     *CharacterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CharacterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CharacterRaw struct {
	Contract *Character // Generic contract binding to access the raw methods on
}

// CharacterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CharacterCallerRaw struct {
	Contract *CharacterCaller // Generic read-only contract binding to access the raw methods on
}

// CharacterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CharacterTransactorRaw struct {
	Contract *CharacterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharacter creates a new instance of Character, bound to a specific deployed contract.
func NewCharacter(address common.Address, backend bind.ContractBackend) (*Character, error) {
	contract, err := bindCharacter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Character{CharacterCaller: CharacterCaller{contract: contract}, CharacterTransactor: CharacterTransactor{contract: contract}, CharacterFilterer: CharacterFilterer{contract: contract}}, nil
}

// NewCharacterCaller creates a new read-only instance of Character, bound to a specific deployed contract.
func NewCharacterCaller(address common.Address, caller bind.ContractCaller) (*CharacterCaller, error) {
	contract, err := bindCharacter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CharacterCaller{contract: contract}, nil
}

// NewCharacterTransactor creates a new write-only instance of Character, bound to a specific deployed contract.
func NewCharacterTransactor(address common.Address, transactor bind.ContractTransactor) (*CharacterTransactor, error) {
	contract, err := bindCharacter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CharacterTransactor{contract: contract}, nil
}

// NewCharacterFilterer creates a new log filterer instance of Character, bound to a specific deployed contract.
func NewCharacterFilterer(address common.Address, filterer bind.ContractFilterer) (*CharacterFilterer, error) {
	contract, err := bindCharacter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CharacterFilterer{contract: contract}, nil
}

// bindCharacter binds a generic wrapper to an already deployed contract.
func bindCharacter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CharacterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Character *CharacterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Character.Contract.CharacterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Character *CharacterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Character.Contract.CharacterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Character *CharacterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Character.Contract.CharacterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Character *CharacterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Character.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Character *CharacterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Character.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Character *CharacterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Character.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Character *CharacterCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Character *CharacterSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Character.Contract.BalanceOf(&_Character.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Character *CharacterCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Character.Contract.BalanceOf(&_Character.CallOpts, owner)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Character *CharacterCaller) CanCreate(opts *bind.CallOpts, handle string, account common.Address) (bool, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "canCreate", handle, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Character *CharacterSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Character.Contract.CanCreate(&_Character.CallOpts, handle, account)
}

// CanCreate is a free data retrieval call binding the contract method 0x7daca6d0.
//
// Solidity: function canCreate(string handle, address account) view returns(bool)
func (_Character *CharacterCallerSession) CanCreate(handle string, account common.Address) (bool, error) {
	return _Character.Contract.CanCreate(&_Character.CallOpts, handle, account)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Character *CharacterCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Character *CharacterSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetApproved(&_Character.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Character *CharacterCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetApproved(&_Character.CallOpts, tokenId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterCaller) GetCharacter(opts *bind.CallOpts, characterId *big.Int) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getCharacter", characterId)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Character.Contract.GetCharacter(&_Character.CallOpts, characterId)
}

// GetCharacter is a free data retrieval call binding the contract method 0xdabb0531.
//
// Solidity: function getCharacter(uint256 characterId) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterCallerSession) GetCharacter(characterId *big.Int) (DataTypesCharacter, error) {
	return _Character.Contract.GetCharacter(&_Character.CallOpts, characterId)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterCaller) GetCharacterByHandle(opts *bind.CallOpts, handle string) (DataTypesCharacter, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getCharacterByHandle", handle)

	if err != nil {
		return *new(DataTypesCharacter), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesCharacter)).(*DataTypesCharacter)

	return out0, err

}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Character.Contract.GetCharacterByHandle(&_Character.CallOpts, handle)
}

// GetCharacterByHandle is a free data retrieval call binding the contract method 0x9a50248d.
//
// Solidity: function getCharacterByHandle(string handle) view returns((uint256,string,string,uint256,address,address))
func (_Character *CharacterCallerSession) GetCharacterByHandle(handle string) (DataTypesCharacter, error) {
	return _Character.Contract.GetCharacterByHandle(&_Character.CallOpts, handle)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Character *CharacterCaller) GetCharacterUri(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getCharacterUri", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Character *CharacterSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Character.Contract.GetCharacterUri(&_Character.CallOpts, characterId)
}

// GetCharacterUri is a free data retrieval call binding the contract method 0x144a3e83.
//
// Solidity: function getCharacterUri(uint256 characterId) view returns(string)
func (_Character *CharacterCallerSession) GetCharacterUri(characterId *big.Int) (string, error) {
	return _Character.Contract.GetCharacterUri(&_Character.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Character *CharacterCaller) GetHandle(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getHandle", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Character *CharacterSession) GetHandle(characterId *big.Int) (string, error) {
	return _Character.Contract.GetHandle(&_Character.CallOpts, characterId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 characterId) view returns(string)
func (_Character *CharacterCallerSession) GetHandle(characterId *big.Int) (string, error) {
	return _Character.Contract.GetHandle(&_Character.CallOpts, characterId)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Character *CharacterCaller) GetLinkModule4Address(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinkModule4Address", account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Character *CharacterSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Character.Contract.GetLinkModule4Address(&_Character.CallOpts, account)
}

// GetLinkModule4Address is a free data retrieval call binding the contract method 0x31b9d08c.
//
// Solidity: function getLinkModule4Address(address account) view returns(address)
func (_Character *CharacterCallerSession) GetLinkModule4Address(account common.Address) (common.Address, error) {
	return _Character.Contract.GetLinkModule4Address(&_Character.CallOpts, account)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Character *CharacterCaller) GetLinkModule4ERC721(opts *bind.CallOpts, tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinkModule4ERC721", tokenAddress, tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Character *CharacterSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetLinkModule4ERC721(&_Character.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4ERC721 is a free data retrieval call binding the contract method 0x2209d145.
//
// Solidity: function getLinkModule4ERC721(address tokenAddress, uint256 tokenId) view returns(address)
func (_Character *CharacterCallerSession) GetLinkModule4ERC721(tokenAddress common.Address, tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetLinkModule4ERC721(&_Character.CallOpts, tokenAddress, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Character *CharacterCaller) GetLinkModule4Linklist(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinkModule4Linklist", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Character *CharacterSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetLinkModule4Linklist(&_Character.CallOpts, tokenId)
}

// GetLinkModule4Linklist is a free data retrieval call binding the contract method 0xfe9299fb.
//
// Solidity: function getLinkModule4Linklist(uint256 tokenId) view returns(address)
func (_Character *CharacterCallerSession) GetLinkModule4Linklist(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.GetLinkModule4Linklist(&_Character.CallOpts, tokenId)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Character *CharacterCaller) GetLinklistContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinklistContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Character *CharacterSession) GetLinklistContract() (common.Address, error) {
	return _Character.Contract.GetLinklistContract(&_Character.CallOpts)
}

// GetLinklistContract is a free data retrieval call binding the contract method 0xc053f6b8.
//
// Solidity: function getLinklistContract() view returns(address)
func (_Character *CharacterCallerSession) GetLinklistContract() (common.Address, error) {
	return _Character.Contract.GetLinklistContract(&_Character.CallOpts)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Character *CharacterCaller) GetLinklistId(opts *bind.CallOpts, characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinklistId", characterId, linkType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Character *CharacterSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Character.Contract.GetLinklistId(&_Character.CallOpts, characterId, linkType)
}

// GetLinklistId is a free data retrieval call binding the contract method 0xd70e10c6.
//
// Solidity: function getLinklistId(uint256 characterId, bytes32 linkType) view returns(uint256)
func (_Character *CharacterCallerSession) GetLinklistId(characterId *big.Int, linkType [32]byte) (*big.Int, error) {
	return _Character.Contract.GetLinklistId(&_Character.CallOpts, characterId, linkType)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Character *CharacterCaller) GetLinklistType(opts *bind.CallOpts, linkListId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinklistType", linkListId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Character *CharacterSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Character.Contract.GetLinklistType(&_Character.CallOpts, linkListId)
}

// GetLinklistType is a free data retrieval call binding the contract method 0x8b4ca06a.
//
// Solidity: function getLinklistType(uint256 linkListId) view returns(bytes32)
func (_Character *CharacterCallerSession) GetLinklistType(linkListId *big.Int) ([32]byte, error) {
	return _Character.Contract.GetLinklistType(&_Character.CallOpts, linkListId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Character *CharacterCaller) GetLinklistUri(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getLinklistUri", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Character *CharacterSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Character.Contract.GetLinklistUri(&_Character.CallOpts, tokenId)
}

// GetLinklistUri is a free data retrieval call binding the contract method 0xdca27135.
//
// Solidity: function getLinklistUri(uint256 tokenId) view returns(string)
func (_Character *CharacterCallerSession) GetLinklistUri(tokenId *big.Int) (string, error) {
	return _Character.Contract.GetLinklistUri(&_Character.CallOpts, tokenId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Character *CharacterCaller) GetNote(opts *bind.CallOpts, characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getNote", characterId, noteId)

	if err != nil {
		return *new(DataTypesNote), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNote)).(*DataTypesNote)

	return out0, err

}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Character *CharacterSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Character.Contract.GetNote(&_Character.CallOpts, characterId, noteId)
}

// GetNote is a free data retrieval call binding the contract method 0xdb491e80.
//
// Solidity: function getNote(uint256 characterId, uint256 noteId) view returns((bytes32,bytes32,string,address,address,address,bool,bool))
func (_Character *CharacterCallerSession) GetNote(characterId *big.Int, noteId *big.Int) (DataTypesNote, error) {
	return _Character.Contract.GetNote(&_Character.CallOpts, characterId, noteId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Character *CharacterCaller) GetOperator(opts *bind.CallOpts, characterId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getOperator", characterId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Character *CharacterSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Character.Contract.GetOperator(&_Character.CallOpts, characterId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 characterId) view returns(address)
func (_Character *CharacterCallerSession) GetOperator(characterId *big.Int) (common.Address, error) {
	return _Character.Contract.GetOperator(&_Character.CallOpts, characterId)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Character *CharacterCaller) GetPrimaryCharacterId(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getPrimaryCharacterId", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Character *CharacterSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Character.Contract.GetPrimaryCharacterId(&_Character.CallOpts, account)
}

// GetPrimaryCharacterId is a free data retrieval call binding the contract method 0x2abc6bf6.
//
// Solidity: function getPrimaryCharacterId(address account) view returns(uint256)
func (_Character *CharacterCallerSession) GetPrimaryCharacterId(account common.Address) (*big.Int, error) {
	return _Character.Contract.GetPrimaryCharacterId(&_Character.CallOpts, account)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Character *CharacterCaller) GetRevision(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "getRevision")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Character *CharacterSession) GetRevision() (*big.Int, error) {
	return _Character.Contract.GetRevision(&_Character.CallOpts)
}

// GetRevision is a free data retrieval call binding the contract method 0x1316529d.
//
// Solidity: function getRevision() pure returns(uint256)
func (_Character *CharacterCallerSession) GetRevision() (*big.Int, error) {
	return _Character.Contract.GetRevision(&_Character.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Character *CharacterCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Character *CharacterSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Character.Contract.IsApprovedForAll(&_Character.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Character *CharacterCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Character.Contract.IsApprovedForAll(&_Character.CallOpts, owner, operator)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Character *CharacterCaller) IsPrimaryCharacter(opts *bind.CallOpts, characterId *big.Int) (bool, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "isPrimaryCharacter", characterId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Character *CharacterSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Character.Contract.IsPrimaryCharacter(&_Character.CallOpts, characterId)
}

// IsPrimaryCharacter is a free data retrieval call binding the contract method 0x8734bbfc.
//
// Solidity: function isPrimaryCharacter(uint256 characterId) view returns(bool)
func (_Character *CharacterCallerSession) IsPrimaryCharacter(characterId *big.Int) (bool, error) {
	return _Character.Contract.IsPrimaryCharacter(&_Character.CallOpts, characterId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Character *CharacterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Character *CharacterSession) Name() (string, error) {
	return _Character.Contract.Name(&_Character.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Character *CharacterCallerSession) Name() (string, error) {
	return _Character.Contract.Name(&_Character.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Character *CharacterCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Character *CharacterSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.OwnerOf(&_Character.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Character *CharacterCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Character.Contract.OwnerOf(&_Character.CallOpts, tokenId)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Character *CharacterCaller) Resolver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "resolver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Character *CharacterSession) Resolver() (common.Address, error) {
	return _Character.Contract.Resolver(&_Character.CallOpts)
}

// Resolver is a free data retrieval call binding the contract method 0x04f3bcec.
//
// Solidity: function resolver() view returns(address)
func (_Character *CharacterCallerSession) Resolver() (common.Address, error) {
	return _Character.Contract.Resolver(&_Character.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Character *CharacterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Character *CharacterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Character.Contract.SupportsInterface(&_Character.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Character *CharacterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Character.Contract.SupportsInterface(&_Character.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Character *CharacterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Character *CharacterSession) Symbol() (string, error) {
	return _Character.Contract.Symbol(&_Character.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Character *CharacterCallerSession) Symbol() (string, error) {
	return _Character.Contract.Symbol(&_Character.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Character *CharacterCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Character *CharacterSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Character.Contract.TokenByIndex(&_Character.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Character *CharacterCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Character.Contract.TokenByIndex(&_Character.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Character *CharacterCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Character *CharacterSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Character.Contract.TokenOfOwnerByIndex(&_Character.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Character *CharacterCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Character.Contract.TokenOfOwnerByIndex(&_Character.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Character *CharacterCaller) TokenURI(opts *bind.CallOpts, characterId *big.Int) (string, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "tokenURI", characterId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Character *CharacterSession) TokenURI(characterId *big.Int) (string, error) {
	return _Character.Contract.TokenURI(&_Character.CallOpts, characterId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 characterId) view returns(string)
func (_Character *CharacterCallerSession) TokenURI(characterId *big.Int) (string, error) {
	return _Character.Contract.TokenURI(&_Character.CallOpts, characterId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Character *CharacterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Character.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Character *CharacterSession) TotalSupply() (*big.Int, error) {
	return _Character.Contract.TotalSupply(&_Character.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Character *CharacterCallerSession) TotalSupply() (*big.Int, error) {
	return _Character.Contract.TotalSupply(&_Character.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Character *CharacterTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Character *CharacterSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.Approve(&_Character.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Character *CharacterTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.Approve(&_Character.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Character *CharacterTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Character *CharacterSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.Burn(&_Character.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Character *CharacterTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.Burn(&_Character.TransactOpts, tokenId)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterTransactor) CreateCharacter(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "createCharacter", vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateCharacter(&_Character.TransactOpts, vars)
}

// CreateCharacter is a paid mutator transaction binding the contract method 0xcd69fe61.
//
// Solidity: function createCharacter((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) CreateCharacter(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateCharacter(&_Character.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterTransactor) CreateProfile(opts *bind.TransactOpts, vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterSession) CreateProfile(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateProfile(&_Character.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xbd5f69cb.
//
// Solidity: function createProfile((address,string,string,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) CreateProfile(vars DataTypesCreateCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateProfile(&_Character.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Character *CharacterTransactor) CreateThenLinkCharacter(opts *bind.TransactOpts, vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "createThenLinkCharacter", vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Character *CharacterSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateThenLinkCharacter(&_Character.TransactOpts, vars)
}

// CreateThenLinkCharacter is a paid mutator transaction binding the contract method 0xf6479d77.
//
// Solidity: function createThenLinkCharacter((uint256,address,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) CreateThenLinkCharacter(vars DataTypescreateThenLinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.CreateThenLinkCharacter(&_Character.TransactOpts, vars)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterTransactor) DeleteNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "deleteNote", characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.DeleteNote(&_Character.TransactOpts, characterId, noteId)
}

// DeleteNote is a paid mutator transaction binding the contract method 0xc2a6fe3b.
//
// Solidity: function deleteNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterTransactorSession) DeleteNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.DeleteNote(&_Character.TransactOpts, characterId, noteId)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Character *CharacterTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "initialize", _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Character *CharacterSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Character.Contract.Initialize(&_Character.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// Initialize is a paid mutator transaction binding the contract method 0xe56f2fe4.
//
// Solidity: function initialize(string _name, string _symbol, address _linklistContract, address _mintNFTImpl, address _periphery, address _resolver) returns()
func (_Character *CharacterTransactorSession) Initialize(_name string, _symbol string, _linklistContract common.Address, _mintNFTImpl common.Address, _periphery common.Address, _resolver common.Address) (*types.Transaction, error) {
	return _Character.Contract.Initialize(&_Character.TransactOpts, _name, _symbol, _linklistContract, _mintNFTImpl, _periphery, _resolver)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkAddress(opts *bind.TransactOpts, vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkAddress", vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Character.Contract.LinkAddress(&_Character.TransactOpts, vars)
}

// LinkAddress is a paid mutator transaction binding the contract method 0x388f5083.
//
// Solidity: function linkAddress((uint256,address,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkAddress(vars DataTypeslinkAddressData) (*types.Transaction, error) {
	return _Character.Contract.LinkAddress(&_Character.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkAnyUri(opts *bind.TransactOpts, vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkAnyUri", vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Character.Contract.LinkAnyUri(&_Character.TransactOpts, vars)
}

// LinkAnyUri is a paid mutator transaction binding the contract method 0x5fb88183.
//
// Solidity: function linkAnyUri((uint256,string,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkAnyUri(vars DataTypeslinkAnyUriData) (*types.Transaction, error) {
	return _Character.Contract.LinkAnyUri(&_Character.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkCharacter(opts *bind.TransactOpts, vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkCharacter", vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.LinkCharacter(&_Character.TransactOpts, vars)
}

// LinkCharacter is a paid mutator transaction binding the contract method 0x188b04b3.
//
// Solidity: function linkCharacter((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkCharacter(vars DataTypeslinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.LinkCharacter(&_Character.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkERC721(opts *bind.TransactOpts, vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkERC721", vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Character.Contract.LinkERC721(&_Character.TransactOpts, vars)
}

// LinkERC721 is a paid mutator transaction binding the contract method 0xcb8e757e.
//
// Solidity: function linkERC721((uint256,address,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkERC721(vars DataTypeslinkERC721Data) (*types.Transaction, error) {
	return _Character.Contract.LinkERC721(&_Character.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkLinklist(opts *bind.TransactOpts, vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkLinklist", vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Character.Contract.LinkLinklist(&_Character.TransactOpts, vars)
}

// LinkLinklist is a paid mutator transaction binding the contract method 0x9864c307.
//
// Solidity: function linkLinklist((uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkLinklist(vars DataTypeslinkLinklistData) (*types.Transaction, error) {
	return _Character.Contract.LinkLinklist(&_Character.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactor) LinkNote(opts *bind.TransactOpts, vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "linkNote", vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Character.Contract.LinkNote(&_Character.TransactOpts, vars)
}

// LinkNote is a paid mutator transaction binding the contract method 0xb9d32845.
//
// Solidity: function linkNote((uint256,uint256,uint256,bytes32,bytes) vars) returns()
func (_Character *CharacterTransactorSession) LinkNote(vars DataTypeslinkNoteData) (*types.Transaction, error) {
	return _Character.Contract.LinkNote(&_Character.TransactOpts, vars)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterTransactor) LockNote(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "lockNote", characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.LockNote(&_Character.TransactOpts, characterId, noteId)
}

// LockNote is a paid mutator transaction binding the contract method 0x74f345cf.
//
// Solidity: function lockNote(uint256 characterId, uint256 noteId) returns()
func (_Character *CharacterTransactorSession) LockNote(characterId *big.Int, noteId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.LockNote(&_Character.TransactOpts, characterId, noteId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 characterId) returns()
func (_Character *CharacterTransactor) MigrateNote(opts *bind.TransactOpts, characterId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "migrateNote", characterId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 characterId) returns()
func (_Character *CharacterSession) MigrateNote(characterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.MigrateNote(&_Character.TransactOpts, characterId)
}

// MigrateNote is a paid mutator transaction binding the contract method 0x14bd433c.
//
// Solidity: function migrateNote(uint256 characterId) returns()
func (_Character *CharacterTransactorSession) MigrateNote(characterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.MigrateNote(&_Character.TransactOpts, characterId)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Character *CharacterTransactor) MintNote(opts *bind.TransactOpts, vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "mintNote", vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Character *CharacterSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Character.Contract.MintNote(&_Character.TransactOpts, vars)
}

// MintNote is a paid mutator transaction binding the contract method 0xa7ccb4bf.
//
// Solidity: function mintNote((uint256,uint256,address,bytes) vars) returns(uint256)
func (_Character *CharacterTransactorSession) MintNote(vars DataTypesMintNoteData) (*types.Transaction, error) {
	return _Character.Contract.MintNote(&_Character.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Character *CharacterTransactor) PostNote(opts *bind.TransactOpts, vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote", vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Character *CharacterSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Character.Contract.PostNote(&_Character.TransactOpts, vars)
}

// PostNote is a paid mutator transaction binding the contract method 0x29c301c2.
//
// Solidity: function postNote((uint256,string,address,bytes,address,bytes,bool) vars) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote(vars DataTypesPostNoteData) (*types.Transaction, error) {
	return _Character.Contract.PostNote(&_Character.TransactOpts, vars)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Character *CharacterTransactor) PostNote4Address(opts *bind.TransactOpts, noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4Address", noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Character *CharacterSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Address(&_Character.TransactOpts, noteData, ethAddress)
}

// PostNote4Address is a paid mutator transaction binding the contract method 0x92f7070b.
//
// Solidity: function postNote4Address((uint256,string,address,bytes,address,bytes,bool) noteData, address ethAddress) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4Address(noteData DataTypesPostNoteData, ethAddress common.Address) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Address(&_Character.TransactOpts, noteData, ethAddress)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Character *CharacterTransactor) PostNote4AnyUri(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4AnyUri", postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Character *CharacterSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Character.Contract.PostNote4AnyUri(&_Character.TransactOpts, postNoteData, uri)
}

// PostNote4AnyUri is a paid mutator transaction binding the contract method 0xf316bacd.
//
// Solidity: function postNote4AnyUri((uint256,string,address,bytes,address,bytes,bool) postNoteData, string uri) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4AnyUri(postNoteData DataTypesPostNoteData, uri string) (*types.Transaction, error) {
	return _Character.Contract.PostNote4AnyUri(&_Character.TransactOpts, postNoteData, uri)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Character *CharacterTransactor) PostNote4Character(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4Character", postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Character *CharacterSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Character(&_Character.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4Character is a paid mutator transaction binding the contract method 0xaf90b112.
//
// Solidity: function postNote4Character((uint256,string,address,bytes,address,bytes,bool) postNoteData, uint256 toCharacterId) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4Character(postNoteData DataTypesPostNoteData, toCharacterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Character(&_Character.TransactOpts, postNoteData, toCharacterId)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Character *CharacterTransactor) PostNote4ERC721(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4ERC721", postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Character *CharacterSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Character.Contract.PostNote4ERC721(&_Character.TransactOpts, postNoteData, erc721)
}

// PostNote4ERC721 is a paid mutator transaction binding the contract method 0x327b2a03.
//
// Solidity: function postNote4ERC721((uint256,string,address,bytes,address,bytes,bool) postNoteData, (address,uint256) erc721) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4ERC721(postNoteData DataTypesPostNoteData, erc721 DataTypesERC721Struct) (*types.Transaction, error) {
	return _Character.Contract.PostNote4ERC721(&_Character.TransactOpts, postNoteData, erc721)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Character *CharacterTransactor) PostNote4Linklist(opts *bind.TransactOpts, noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4Linklist", noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Character *CharacterSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Linklist(&_Character.TransactOpts, noteData, toLinklistId)
}

// PostNote4Linklist is a paid mutator transaction binding the contract method 0x44b82a24.
//
// Solidity: function postNote4Linklist((uint256,string,address,bytes,address,bytes,bool) noteData, uint256 toLinklistId) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4Linklist(noteData DataTypesPostNoteData, toLinklistId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Linklist(&_Character.TransactOpts, noteData, toLinklistId)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Character *CharacterTransactor) PostNote4Note(opts *bind.TransactOpts, postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "postNote4Note", postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Character *CharacterSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Note(&_Character.TransactOpts, postNoteData, note)
}

// PostNote4Note is a paid mutator transaction binding the contract method 0x9a4dec18.
//
// Solidity: function postNote4Note((uint256,string,address,bytes,address,bytes,bool) postNoteData, (uint256,uint256) note) returns(uint256)
func (_Character *CharacterTransactorSession) PostNote4Note(postNoteData DataTypesPostNoteData, note DataTypesNoteStruct) (*types.Transaction, error) {
	return _Character.Contract.PostNote4Note(&_Character.TransactOpts, postNoteData, note)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.SafeTransferFrom(&_Character.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.SafeTransferFrom(&_Character.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Character *CharacterTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Character *CharacterSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Character.Contract.SafeTransferFrom0(&_Character.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Character *CharacterTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Character.Contract.SafeTransferFrom0(&_Character.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Character *CharacterTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Character *CharacterSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Character.Contract.SetApprovalForAll(&_Character.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Character *CharacterTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Character.Contract.SetApprovalForAll(&_Character.TransactOpts, operator, approved)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Character *CharacterTransactor) SetCharacterUri(opts *bind.TransactOpts, characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setCharacterUri", characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Character *CharacterSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetCharacterUri(&_Character.TransactOpts, characterId, newUri)
}

// SetCharacterUri is a paid mutator transaction binding the contract method 0x47f94de7.
//
// Solidity: function setCharacterUri(uint256 characterId, string newUri) returns()
func (_Character *CharacterTransactorSession) SetCharacterUri(characterId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetCharacterUri(&_Character.TransactOpts, characterId, newUri)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Character *CharacterTransactor) SetHandle(opts *bind.TransactOpts, characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setHandle", characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Character *CharacterSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Character.Contract.SetHandle(&_Character.TransactOpts, characterId, newHandle)
}

// SetHandle is a paid mutator transaction binding the contract method 0xa6e6178d.
//
// Solidity: function setHandle(uint256 characterId, string newHandle) returns()
func (_Character *CharacterTransactorSession) SetHandle(characterId *big.Int, newHandle string) (*types.Transaction, error) {
	return _Character.Contract.SetHandle(&_Character.TransactOpts, characterId, newHandle)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetLinkModule4Address(opts *bind.TransactOpts, vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinkModule4Address", vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Character *CharacterSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Address(&_Character.TransactOpts, vars)
}

// SetLinkModule4Address is a paid mutator transaction binding the contract method 0x08cb68ff.
//
// Solidity: function setLinkModule4Address((address,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetLinkModule4Address(vars DataTypessetLinkModule4AddressData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Address(&_Character.TransactOpts, vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetLinkModule4Character(opts *bind.TransactOpts, vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinkModule4Character", vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Character *CharacterSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Character(&_Character.TransactOpts, vars)
}

// SetLinkModule4Character is a paid mutator transaction binding the contract method 0xfd2d866f.
//
// Solidity: function setLinkModule4Character((uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetLinkModule4Character(vars DataTypessetLinkModule4CharacterData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Character(&_Character.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetLinkModule4ERC721(opts *bind.TransactOpts, vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinkModule4ERC721", vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Character *CharacterSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4ERC721(&_Character.TransactOpts, vars)
}

// SetLinkModule4ERC721 is a paid mutator transaction binding the contract method 0x69492c97.
//
// Solidity: function setLinkModule4ERC721((address,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetLinkModule4ERC721(vars DataTypessetLinkModule4ERC721Data) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4ERC721(&_Character.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetLinkModule4Linklist(opts *bind.TransactOpts, vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinkModule4Linklist", vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Character *CharacterSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Linklist(&_Character.TransactOpts, vars)
}

// SetLinkModule4Linklist is a paid mutator transaction binding the contract method 0x0c4dd5f2.
//
// Solidity: function setLinkModule4Linklist((uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetLinkModule4Linklist(vars DataTypessetLinkModule4LinklistData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Linklist(&_Character.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetLinkModule4Note(opts *bind.TransactOpts, vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinkModule4Note", vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Note(&_Character.TransactOpts, vars)
}

// SetLinkModule4Note is a paid mutator transaction binding the contract method 0xdb8c198d.
//
// Solidity: function setLinkModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetLinkModule4Note(vars DataTypessetLinkModule4NoteData) (*types.Transaction, error) {
	return _Character.Contract.SetLinkModule4Note(&_Character.TransactOpts, vars)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Character *CharacterTransactor) SetLinklistUri(opts *bind.TransactOpts, linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setLinklistUri", linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Character *CharacterSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Character.Contract.SetLinklistUri(&_Character.TransactOpts, linklistId, uri)
}

// SetLinklistUri is a paid mutator transaction binding the contract method 0x33f06ee6.
//
// Solidity: function setLinklistUri(uint256 linklistId, string uri) returns()
func (_Character *CharacterTransactorSession) SetLinklistUri(linklistId *big.Int, uri string) (*types.Transaction, error) {
	return _Character.Contract.SetLinklistUri(&_Character.TransactOpts, linklistId, uri)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactor) SetMintModule4Note(opts *bind.TransactOpts, vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setMintModule4Note", vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Character.Contract.SetMintModule4Note(&_Character.TransactOpts, vars)
}

// SetMintModule4Note is a paid mutator transaction binding the contract method 0xd23b320b.
//
// Solidity: function setMintModule4Note((uint256,uint256,address,bytes) vars) returns()
func (_Character *CharacterTransactorSession) SetMintModule4Note(vars DataTypessetMintModule4NoteData) (*types.Transaction, error) {
	return _Character.Contract.SetMintModule4Note(&_Character.TransactOpts, vars)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Character *CharacterTransactor) SetNoteUri(opts *bind.TransactOpts, characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setNoteUri", characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Character *CharacterSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetNoteUri(&_Character.TransactOpts, characterId, noteId, newUri)
}

// SetNoteUri is a paid mutator transaction binding the contract method 0x628b644a.
//
// Solidity: function setNoteUri(uint256 characterId, uint256 noteId, string newUri) returns()
func (_Character *CharacterTransactorSession) SetNoteUri(characterId *big.Int, noteId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetNoteUri(&_Character.TransactOpts, characterId, noteId, newUri)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Character *CharacterTransactor) SetOperator(opts *bind.TransactOpts, characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setOperator", characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Character *CharacterSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Character.Contract.SetOperator(&_Character.TransactOpts, characterId, operator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xe7a1c1c0.
//
// Solidity: function setOperator(uint256 characterId, address operator) returns()
func (_Character *CharacterTransactorSession) SetOperator(characterId *big.Int, operator common.Address) (*types.Transaction, error) {
	return _Character.Contract.SetOperator(&_Character.TransactOpts, characterId, operator)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Character *CharacterTransactor) SetPrimaryCharacterId(opts *bind.TransactOpts, characterId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setPrimaryCharacterId", characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Character *CharacterSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.SetPrimaryCharacterId(&_Character.TransactOpts, characterId)
}

// SetPrimaryCharacterId is a paid mutator transaction binding the contract method 0xf2ad8075.
//
// Solidity: function setPrimaryCharacterId(uint256 characterId) returns()
func (_Character *CharacterTransactorSession) SetPrimaryCharacterId(characterId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.SetPrimaryCharacterId(&_Character.TransactOpts, characterId)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Character *CharacterTransactor) SetProfileUri(opts *bind.TransactOpts, profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setProfileUri", profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Character *CharacterSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetProfileUri(&_Character.TransactOpts, profileId, newUri)
}

// SetProfileUri is a paid mutator transaction binding the contract method 0x7c392b51.
//
// Solidity: function setProfileUri(uint256 profileId, string newUri) returns()
func (_Character *CharacterTransactorSession) SetProfileUri(profileId *big.Int, newUri string) (*types.Transaction, error) {
	return _Character.Contract.SetProfileUri(&_Character.TransactOpts, profileId, newUri)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Character *CharacterTransactor) SetSocialToken(opts *bind.TransactOpts, characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "setSocialToken", characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Character *CharacterSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Character.Contract.SetSocialToken(&_Character.TransactOpts, characterId, tokenAddress)
}

// SetSocialToken is a paid mutator transaction binding the contract method 0x95d9fa7d.
//
// Solidity: function setSocialToken(uint256 characterId, address tokenAddress) returns()
func (_Character *CharacterTransactorSession) SetSocialToken(characterId *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Character.Contract.SetSocialToken(&_Character.TransactOpts, characterId, tokenAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.TransferFrom(&_Character.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Character *CharacterTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Character.Contract.TransferFrom(&_Character.TransactOpts, from, to, tokenId)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkAddress(opts *bind.TransactOpts, vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkAddress", vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkAddress(&_Character.TransactOpts, vars)
}

// UnlinkAddress is a paid mutator transaction binding the contract method 0x93f057e5.
//
// Solidity: function unlinkAddress((uint256,address,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkAddress(vars DataTypesunlinkAddressData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkAddress(&_Character.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkAnyUri(opts *bind.TransactOpts, vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkAnyUri", vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkAnyUri(&_Character.TransactOpts, vars)
}

// UnlinkAnyUri is a paid mutator transaction binding the contract method 0xef0828ab.
//
// Solidity: function unlinkAnyUri((uint256,string,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkAnyUri(vars DataTypesunlinkAnyUriData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkAnyUri(&_Character.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkCharacter(opts *bind.TransactOpts, vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkCharacter", vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkCharacter(&_Character.TransactOpts, vars)
}

// UnlinkCharacter is a paid mutator transaction binding the contract method 0x0ff98244.
//
// Solidity: function unlinkCharacter((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkCharacter(vars DataTypesunlinkCharacterData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkCharacter(&_Character.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkERC721(opts *bind.TransactOpts, vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkERC721", vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Character.Contract.UnlinkERC721(&_Character.TransactOpts, vars)
}

// UnlinkERC721 is a paid mutator transaction binding the contract method 0x867884e6.
//
// Solidity: function unlinkERC721((uint256,address,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkERC721(vars DataTypesunlinkERC721Data) (*types.Transaction, error) {
	return _Character.Contract.UnlinkERC721(&_Character.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkLinklist(opts *bind.TransactOpts, vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkLinklist", vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkLinklist(&_Character.TransactOpts, vars)
}

// UnlinkLinklist is a paid mutator transaction binding the contract method 0x5a936d10.
//
// Solidity: function unlinkLinklist((uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkLinklist(vars DataTypesunlinkLinklistData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkLinklist(&_Character.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactor) UnlinkNote(opts *bind.TransactOpts, vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Character.contract.Transact(opts, "unlinkNote", vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkNote(&_Character.TransactOpts, vars)
}

// UnlinkNote is a paid mutator transaction binding the contract method 0x40ad34d8.
//
// Solidity: function unlinkNote((uint256,uint256,uint256,bytes32) vars) returns()
func (_Character *CharacterTransactorSession) UnlinkNote(vars DataTypesunlinkNoteData) (*types.Transaction, error) {
	return _Character.Contract.UnlinkNote(&_Character.TransactOpts, vars)
}

// CharacterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Character contract.
type CharacterApprovalIterator struct {
	Event *CharacterApproval // Event containing the contract specifics and raw log

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
func (it *CharacterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterApproval)
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
		it.Event = new(CharacterApproval)
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
func (it *CharacterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterApproval represents a Approval event raised by the Character contract.
type CharacterApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Character *CharacterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CharacterApprovalIterator, error) {

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

	logs, sub, err := _Character.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterApprovalIterator{contract: _Character.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Character *CharacterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CharacterApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Character.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterApproval)
				if err := _Character.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Character *CharacterFilterer) ParseApproval(log types.Log) (*CharacterApproval, error) {
	event := new(CharacterApproval)
	if err := _Character.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Character contract.
type CharacterApprovalForAllIterator struct {
	Event *CharacterApprovalForAll // Event containing the contract specifics and raw log

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
func (it *CharacterApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterApprovalForAll)
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
		it.Event = new(CharacterApprovalForAll)
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
func (it *CharacterApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterApprovalForAll represents a ApprovalForAll event raised by the Character contract.
type CharacterApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Character *CharacterFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CharacterApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CharacterApprovalForAllIterator{contract: _Character.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Character *CharacterFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CharacterApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterApprovalForAll)
				if err := _Character.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Character *CharacterFilterer) ParseApprovalForAll(log types.Log) (*CharacterApprovalForAll, error) {
	event := new(CharacterApprovalForAll)
	if err := _Character.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterAttachLinklistIterator is returned from FilterAttachLinklist and is used to iterate over the raw logs and unpacked data for AttachLinklist events raised by the Character contract.
type CharacterAttachLinklistIterator struct {
	Event *CharacterAttachLinklist // Event containing the contract specifics and raw log

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
func (it *CharacterAttachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterAttachLinklist)
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
		it.Event = new(CharacterAttachLinklist)
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
func (it *CharacterAttachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterAttachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterAttachLinklist represents a AttachLinklist event raised by the Character contract.
type CharacterAttachLinklist struct {
	LinklistId  *big.Int
	CharacterId *big.Int
	LinkType    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttachLinklist is a free log retrieval operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) FilterAttachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (*CharacterAttachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "AttachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &CharacterAttachLinklistIterator{contract: _Character.contract, event: "AttachLinklist", logs: logs, sub: sub}, nil
}

// WatchAttachLinklist is a free log subscription operation binding the contract event 0x94703ec1dd639b589d05fa7c0c14ee1e83b906bfb5f50642cc7ea415a8172288.
//
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) WatchAttachLinklist(opts *bind.WatchOpts, sink chan<- *CharacterAttachLinklist, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "AttachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterAttachLinklist)
				if err := _Character.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
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
// Solidity: event AttachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) ParseAttachLinklist(log types.Log) (*CharacterAttachLinklist, error) {
	event := new(CharacterAttachLinklist)
	if err := _Character.contract.UnpackLog(event, "AttachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterBaseInitializedIterator is returned from FilterBaseInitialized and is used to iterate over the raw logs and unpacked data for BaseInitialized events raised by the Character contract.
type CharacterBaseInitializedIterator struct {
	Event *CharacterBaseInitialized // Event containing the contract specifics and raw log

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
func (it *CharacterBaseInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterBaseInitialized)
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
		it.Event = new(CharacterBaseInitialized)
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
func (it *CharacterBaseInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterBaseInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterBaseInitialized represents a BaseInitialized event raised by the Character contract.
type CharacterBaseInitialized struct {
	Name      string
	Symbol    string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBaseInitialized is a free log retrieval operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Character *CharacterFilterer) FilterBaseInitialized(opts *bind.FilterOpts) (*CharacterBaseInitializedIterator, error) {

	logs, sub, err := _Character.contract.FilterLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return &CharacterBaseInitializedIterator{contract: _Character.contract, event: "BaseInitialized", logs: logs, sub: sub}, nil
}

// WatchBaseInitialized is a free log subscription operation binding the contract event 0x414cd0b34676984f09a5f76ce9718d4062e50283abe0e7e274a9a5b4e0c99c30.
//
// Solidity: event BaseInitialized(string name, string symbol, uint256 timestamp)
func (_Character *CharacterFilterer) WatchBaseInitialized(opts *bind.WatchOpts, sink chan<- *CharacterBaseInitialized) (event.Subscription, error) {

	logs, sub, err := _Character.contract.WatchLogs(opts, "BaseInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterBaseInitialized)
				if err := _Character.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
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
func (_Character *CharacterFilterer) ParseBaseInitialized(log types.Log) (*CharacterBaseInitialized, error) {
	event := new(CharacterBaseInitialized)
	if err := _Character.contract.UnpackLog(event, "BaseInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterCharacterCreatedIterator is returned from FilterCharacterCreated and is used to iterate over the raw logs and unpacked data for CharacterCreated events raised by the Character contract.
type CharacterCharacterCreatedIterator struct {
	Event *CharacterCharacterCreated // Event containing the contract specifics and raw log

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
func (it *CharacterCharacterCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterCharacterCreated)
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
		it.Event = new(CharacterCharacterCreated)
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
func (it *CharacterCharacterCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterCharacterCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterCharacterCreated represents a CharacterCreated event raised by the Character contract.
type CharacterCharacterCreated struct {
	CharacterId *big.Int
	Creator     common.Address
	To          common.Address
	Handle      string
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCharacterCreated is a free log retrieval operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Character *CharacterFilterer) FilterCharacterCreated(opts *bind.FilterOpts, characterId []*big.Int, creator []common.Address, to []common.Address) (*CharacterCharacterCreatedIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "CharacterCreated", characterIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CharacterCharacterCreatedIterator{contract: _Character.contract, event: "CharacterCreated", logs: logs, sub: sub}, nil
}

// WatchCharacterCreated is a free log subscription operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Character *CharacterFilterer) WatchCharacterCreated(opts *bind.WatchOpts, sink chan<- *CharacterCharacterCreated, characterId []*big.Int, creator []common.Address, to []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "CharacterCreated", characterIdRule, creatorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterCharacterCreated)
				if err := _Character.contract.UnpackLog(event, "CharacterCreated", log); err != nil {
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

// ParseCharacterCreated is a log parse operation binding the contract event 0x71c6a413fae531e87669ea74b7c28e2afa90047f8a36be6690c7ff625b18afa6.
//
// Solidity: event CharacterCreated(uint256 indexed characterId, address indexed creator, address indexed to, string handle, uint256 timestamp)
func (_Character *CharacterFilterer) ParseCharacterCreated(log types.Log) (*CharacterCharacterCreated, error) {
	event := new(CharacterCharacterCreated)
	if err := _Character.contract.UnpackLog(event, "CharacterCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterDeleteNoteIterator is returned from FilterDeleteNote and is used to iterate over the raw logs and unpacked data for DeleteNote events raised by the Character contract.
type CharacterDeleteNoteIterator struct {
	Event *CharacterDeleteNote // Event containing the contract specifics and raw log

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
func (it *CharacterDeleteNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterDeleteNote)
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
		it.Event = new(CharacterDeleteNote)
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
func (it *CharacterDeleteNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterDeleteNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterDeleteNote represents a DeleteNote event raised by the Character contract.
type CharacterDeleteNote struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeleteNote is a free log retrieval operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) FilterDeleteNote(opts *bind.FilterOpts, characterId []*big.Int) (*CharacterDeleteNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "DeleteNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterDeleteNoteIterator{contract: _Character.contract, event: "DeleteNote", logs: logs, sub: sub}, nil
}

// WatchDeleteNote is a free log subscription operation binding the contract event 0x4f1db9708b537c1d26a7af4b235fd079bf2342d92a276e27eb6c8717e8bbcf93.
//
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) WatchDeleteNote(opts *bind.WatchOpts, sink chan<- *CharacterDeleteNote, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "DeleteNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterDeleteNote)
				if err := _Character.contract.UnpackLog(event, "DeleteNote", log); err != nil {
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
// Solidity: event DeleteNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) ParseDeleteNote(log types.Log) (*CharacterDeleteNote, error) {
	event := new(CharacterDeleteNote)
	if err := _Character.contract.UnpackLog(event, "DeleteNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterDetachLinklistIterator is returned from FilterDetachLinklist and is used to iterate over the raw logs and unpacked data for DetachLinklist events raised by the Character contract.
type CharacterDetachLinklistIterator struct {
	Event *CharacterDetachLinklist // Event containing the contract specifics and raw log

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
func (it *CharacterDetachLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterDetachLinklist)
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
		it.Event = new(CharacterDetachLinklist)
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
func (it *CharacterDetachLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterDetachLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterDetachLinklist represents a DetachLinklist event raised by the Character contract.
type CharacterDetachLinklist struct {
	LinklistId  *big.Int
	CharacterId *big.Int
	LinkType    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDetachLinklist is a free log retrieval operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) FilterDetachLinklist(opts *bind.FilterOpts, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (*CharacterDetachLinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "DetachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &CharacterDetachLinklistIterator{contract: _Character.contract, event: "DetachLinklist", logs: logs, sub: sub}, nil
}

// WatchDetachLinklist is a free log subscription operation binding the contract event 0x5751ba9aebec2dcd504f8865b0e0c98a772403773fe528fce2fce580c9a8a165.
//
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) WatchDetachLinklist(opts *bind.WatchOpts, sink chan<- *CharacterDetachLinklist, linklistId []*big.Int, characterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "DetachLinklist", linklistIdRule, characterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterDetachLinklist)
				if err := _Character.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
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
// Solidity: event DetachLinklist(uint256 indexed linklistId, uint256 indexed characterId, bytes32 indexed linkType)
func (_Character *CharacterFilterer) ParseDetachLinklist(log types.Log) (*CharacterDetachLinklist, error) {
	event := new(CharacterDetachLinklist)
	if err := _Character.contract.UnpackLog(event, "DetachLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkAddressIterator is returned from FilterLinkAddress and is used to iterate over the raw logs and unpacked data for LinkAddress events raised by the Character contract.
type CharacterLinkAddressIterator struct {
	Event *CharacterLinkAddress // Event containing the contract specifics and raw log

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
func (it *CharacterLinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkAddress)
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
		it.Event = new(CharacterLinkAddress)
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
func (it *CharacterLinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkAddress represents a LinkAddress event raised by the Character contract.
type CharacterLinkAddress struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkAddress is a free log retrieval operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterLinkAddress(opts *bind.FilterOpts, fromCharacterId []*big.Int, ethAddress []common.Address) (*CharacterLinkAddressIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkAddressIterator{contract: _Character.contract, event: "LinkAddress", logs: logs, sub: sub}, nil
}

// WatchLinkAddress is a free log subscription operation binding the contract event 0x3dad60a88f1d8ee170dfbeb8c65c705bd47922f205e774e10e49e4e53d93a8bd.
//
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchLinkAddress(opts *bind.WatchOpts, sink chan<- *CharacterLinkAddress, fromCharacterId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkAddress)
				if err := _Character.contract.UnpackLog(event, "LinkAddress", log); err != nil {
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
// Solidity: event LinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseLinkAddress(log types.Log) (*CharacterLinkAddress, error) {
	event := new(CharacterLinkAddress)
	if err := _Character.contract.UnpackLog(event, "LinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkAnyUriIterator is returned from FilterLinkAnyUri and is used to iterate over the raw logs and unpacked data for LinkAnyUri events raised by the Character contract.
type CharacterLinkAnyUriIterator struct {
	Event *CharacterLinkAnyUri // Event containing the contract specifics and raw log

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
func (it *CharacterLinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkAnyUri)
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
		it.Event = new(CharacterLinkAnyUri)
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
func (it *CharacterLinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkAnyUri represents a LinkAnyUri event raised by the Character contract.
type CharacterLinkAnyUri struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkAnyUri is a free log retrieval operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterLinkAnyUri(opts *bind.FilterOpts, fromCharacterId []*big.Int) (*CharacterLinkAnyUriIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkAnyUriIterator{contract: _Character.contract, event: "LinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchLinkAnyUri is a free log subscription operation binding the contract event 0x2585d212014e5edb24ec129ea2e571eb4b8ac8265156ceb119abdcfa591b746d.
//
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchLinkAnyUri(opts *bind.WatchOpts, sink chan<- *CharacterLinkAnyUri, fromCharacterId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkAnyUri)
				if err := _Character.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
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
// Solidity: event LinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseLinkAnyUri(log types.Log) (*CharacterLinkAnyUri, error) {
	event := new(CharacterLinkAnyUri)
	if err := _Character.contract.UnpackLog(event, "LinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkCharacterIterator is returned from FilterLinkCharacter and is used to iterate over the raw logs and unpacked data for LinkCharacter events raised by the Character contract.
type CharacterLinkCharacterIterator struct {
	Event *CharacterLinkCharacter // Event containing the contract specifics and raw log

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
func (it *CharacterLinkCharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkCharacter)
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
		it.Event = new(CharacterLinkCharacter)
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
func (it *CharacterLinkCharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkCharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkCharacter represents a LinkCharacter event raised by the Character contract.
type CharacterLinkCharacter struct {
	Account         common.Address
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkCharacter is a free log retrieval operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterLinkCharacter(opts *bind.FilterOpts, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (*CharacterLinkCharacterIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkCharacterIterator{contract: _Character.contract, event: "LinkCharacter", logs: logs, sub: sub}, nil
}

// WatchLinkCharacter is a free log subscription operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchLinkCharacter(opts *bind.WatchOpts, sink chan<- *CharacterLinkCharacter, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkCharacter)
				if err := _Character.contract.UnpackLog(event, "LinkCharacter", log); err != nil {
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

// ParseLinkCharacter is a log parse operation binding the contract event 0x90a3a33689f2c64e961b6b2dc225d86e878360f111f66aaa290d36c2e13b0180.
//
// Solidity: event LinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseLinkCharacter(log types.Log) (*CharacterLinkCharacter, error) {
	event := new(CharacterLinkCharacter)
	if err := _Character.contract.UnpackLog(event, "LinkCharacter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkCharacterLinkIterator is returned from FilterLinkCharacterLink and is used to iterate over the raw logs and unpacked data for LinkCharacterLink events raised by the Character contract.
type CharacterLinkCharacterLinkIterator struct {
	Event *CharacterLinkCharacterLink // Event containing the contract specifics and raw log

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
func (it *CharacterLinkCharacterLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkCharacterLink)
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
		it.Event = new(CharacterLinkCharacterLink)
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
func (it *CharacterLinkCharacterLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkCharacterLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkCharacterLink represents a LinkCharacterLink event raised by the Character contract.
type CharacterLinkCharacterLink struct {
	FromCharacterId   *big.Int
	LinkType          [32]byte
	ClFromCharacterId *big.Int
	ClToCharacterId   *big.Int
	ClLinkType        [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLinkCharacterLink is a free log retrieval operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) FilterLinkCharacterLink(opts *bind.FilterOpts, fromCharacterId []*big.Int, linkType [][32]byte) (*CharacterLinkCharacterLinkIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkCharacterLinkIterator{contract: _Character.contract, event: "LinkCharacterLink", logs: logs, sub: sub}, nil
}

// WatchLinkCharacterLink is a free log subscription operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) WatchLinkCharacterLink(opts *bind.WatchOpts, sink chan<- *CharacterLinkCharacterLink, fromCharacterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkCharacterLink)
				if err := _Character.contract.UnpackLog(event, "LinkCharacterLink", log); err != nil {
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

// ParseLinkCharacterLink is a log parse operation binding the contract event 0x68023d5cb8e71265d7242a843a21afe933830debc738c00c5c3dc82666b4f1ae.
//
// Solidity: event LinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharacterId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) ParseLinkCharacterLink(log types.Log) (*CharacterLinkCharacterLink, error) {
	event := new(CharacterLinkCharacterLink)
	if err := _Character.contract.UnpackLog(event, "LinkCharacterLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkERC721Iterator is returned from FilterLinkERC721 and is used to iterate over the raw logs and unpacked data for LinkERC721 events raised by the Character contract.
type CharacterLinkERC721Iterator struct {
	Event *CharacterLinkERC721 // Event containing the contract specifics and raw log

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
func (it *CharacterLinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkERC721)
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
		it.Event = new(CharacterLinkERC721)
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
func (it *CharacterLinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkERC721 represents a LinkERC721 event raised by the Character contract.
type CharacterLinkERC721 struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkERC721 is a free log retrieval operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterLinkERC721(opts *bind.FilterOpts, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*CharacterLinkERC721Iterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkERC721Iterator{contract: _Character.contract, event: "LinkERC721", logs: logs, sub: sub}, nil
}

// WatchLinkERC721 is a free log subscription operation binding the contract event 0x72413a1a5c43b255ad065653bf49db3c79ff7853ddaa4f4962ccd05e5368ad77.
//
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchLinkERC721(opts *bind.WatchOpts, sink chan<- *CharacterLinkERC721, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkERC721)
				if err := _Character.contract.UnpackLog(event, "LinkERC721", log); err != nil {
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
// Solidity: event LinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseLinkERC721(log types.Log) (*CharacterLinkERC721, error) {
	event := new(CharacterLinkERC721)
	if err := _Character.contract.UnpackLog(event, "LinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkLinklistIterator is returned from FilterLinkLinklist and is used to iterate over the raw logs and unpacked data for LinkLinklist events raised by the Character contract.
type CharacterLinkLinklistIterator struct {
	Event *CharacterLinkLinklist // Event containing the contract specifics and raw log

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
func (it *CharacterLinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkLinklist)
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
		it.Event = new(CharacterLinkLinklist)
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
func (it *CharacterLinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkLinklist represents a LinkLinklist event raised by the Character contract.
type CharacterLinkLinklist struct {
	FromCharacterId *big.Int
	ToLinklistId    *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkLinklist is a free log retrieval operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) FilterLinkLinklist(opts *bind.FilterOpts, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*CharacterLinkLinklistIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkLinklistIterator{contract: _Character.contract, event: "LinkLinklist", logs: logs, sub: sub}, nil
}

// WatchLinkLinklist is a free log subscription operation binding the contract event 0x2e398bc911c0eb636118bc7906bdfa98b2ccf6ef7ee76e3f07068d0eee488e3f.
//
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) WatchLinkLinklist(opts *bind.WatchOpts, sink chan<- *CharacterLinkLinklist, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkLinklist)
				if err := _Character.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
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
// Solidity: event LinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) ParseLinkLinklist(log types.Log) (*CharacterLinkLinklist, error) {
	event := new(CharacterLinkLinklist)
	if err := _Character.contract.UnpackLog(event, "LinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinkNoteIterator is returned from FilterLinkNote and is used to iterate over the raw logs and unpacked data for LinkNote events raised by the Character contract.
type CharacterLinkNoteIterator struct {
	Event *CharacterLinkNote // Event containing the contract specifics and raw log

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
func (it *CharacterLinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinkNote)
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
		it.Event = new(CharacterLinkNote)
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
func (it *CharacterLinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinkNote represents a LinkNote event raised by the Character contract.
type CharacterLinkNote struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLinkNote is a free log retrieval operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterLinkNote(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (*CharacterLinkNoteIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLinkNoteIterator{contract: _Character.contract, event: "LinkNote", logs: logs, sub: sub}, nil
}

// WatchLinkNote is a free log subscription operation binding the contract event 0x3cae5b1196087b560b6735377bbf745e5754f8082348a432b806afa01686ef48.
//
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchLinkNote(opts *bind.WatchOpts, sink chan<- *CharacterLinkNote, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinkNote)
				if err := _Character.contract.UnpackLog(event, "LinkNote", log); err != nil {
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
// Solidity: event LinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseLinkNote(log types.Log) (*CharacterLinkNote, error) {
	event := new(CharacterLinkNote)
	if err := _Character.contract.UnpackLog(event, "LinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLinklistNFTInitializedIterator is returned from FilterLinklistNFTInitialized and is used to iterate over the raw logs and unpacked data for LinklistNFTInitialized events raised by the Character contract.
type CharacterLinklistNFTInitializedIterator struct {
	Event *CharacterLinklistNFTInitialized // Event containing the contract specifics and raw log

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
func (it *CharacterLinklistNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLinklistNFTInitialized)
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
		it.Event = new(CharacterLinklistNFTInitialized)
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
func (it *CharacterLinklistNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLinklistNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLinklistNFTInitialized represents a LinklistNFTInitialized event raised by the Character contract.
type CharacterLinklistNFTInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLinklistNFTInitialized is a free log retrieval operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Character *CharacterFilterer) FilterLinklistNFTInitialized(opts *bind.FilterOpts) (*CharacterLinklistNFTInitializedIterator, error) {

	logs, sub, err := _Character.contract.FilterLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &CharacterLinklistNFTInitializedIterator{contract: _Character.contract, event: "LinklistNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchLinklistNFTInitialized is a free log subscription operation binding the contract event 0xcfdec2ffedf2f5ec02de6f351c5f9b6150601f657926e9e87b16390d562af4e7.
//
// Solidity: event LinklistNFTInitialized(uint256 timestamp)
func (_Character *CharacterFilterer) WatchLinklistNFTInitialized(opts *bind.WatchOpts, sink chan<- *CharacterLinklistNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Character.contract.WatchLogs(opts, "LinklistNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLinklistNFTInitialized)
				if err := _Character.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
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
func (_Character *CharacterFilterer) ParseLinklistNFTInitialized(log types.Log) (*CharacterLinklistNFTInitialized, error) {
	event := new(CharacterLinklistNFTInitialized)
	if err := _Character.contract.UnpackLog(event, "LinklistNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterLockNoteIterator is returned from FilterLockNote and is used to iterate over the raw logs and unpacked data for LockNote events raised by the Character contract.
type CharacterLockNoteIterator struct {
	Event *CharacterLockNote // Event containing the contract specifics and raw log

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
func (it *CharacterLockNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterLockNote)
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
		it.Event = new(CharacterLockNote)
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
func (it *CharacterLockNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterLockNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterLockNote represents a LockNote event raised by the Character contract.
type CharacterLockNote struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLockNote is a free log retrieval operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) FilterLockNote(opts *bind.FilterOpts, characterId []*big.Int) (*CharacterLockNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "LockNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterLockNoteIterator{contract: _Character.contract, event: "LockNote", logs: logs, sub: sub}, nil
}

// WatchLockNote is a free log subscription operation binding the contract event 0x036469f3e73c83520cdefa197d7a9c854c2f8bc0164b82e9f2bd4aa7e150fd30.
//
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) WatchLockNote(opts *bind.WatchOpts, sink chan<- *CharacterLockNote, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "LockNote", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterLockNote)
				if err := _Character.contract.UnpackLog(event, "LockNote", log); err != nil {
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
// Solidity: event LockNote(uint256 indexed characterId, uint256 noteId)
func (_Character *CharacterFilterer) ParseLockNote(log types.Log) (*CharacterLockNote, error) {
	event := new(CharacterLockNote)
	if err := _Character.contract.UnpackLog(event, "LockNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterMintNFTInitializedIterator is returned from FilterMintNFTInitialized and is used to iterate over the raw logs and unpacked data for MintNFTInitialized events raised by the Character contract.
type CharacterMintNFTInitializedIterator struct {
	Event *CharacterMintNFTInitialized // Event containing the contract specifics and raw log

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
func (it *CharacterMintNFTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterMintNFTInitialized)
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
		it.Event = new(CharacterMintNFTInitialized)
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
func (it *CharacterMintNFTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterMintNFTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterMintNFTInitialized represents a MintNFTInitialized event raised by the Character contract.
type CharacterMintNFTInitialized struct {
	CharacterId *big.Int
	NoteId      *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMintNFTInitialized is a free log retrieval operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Character *CharacterFilterer) FilterMintNFTInitialized(opts *bind.FilterOpts) (*CharacterMintNFTInitializedIterator, error) {

	logs, sub, err := _Character.contract.FilterLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return &CharacterMintNFTInitializedIterator{contract: _Character.contract, event: "MintNFTInitialized", logs: logs, sub: sub}, nil
}

// WatchMintNFTInitialized is a free log subscription operation binding the contract event 0x6f606e3871b6f24ddc0daa88a8ed8e8761eb9d1474e64322940f858df287de50.
//
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Character *CharacterFilterer) WatchMintNFTInitialized(opts *bind.WatchOpts, sink chan<- *CharacterMintNFTInitialized) (event.Subscription, error) {

	logs, sub, err := _Character.contract.WatchLogs(opts, "MintNFTInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterMintNFTInitialized)
				if err := _Character.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
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
// Solidity: event MintNFTInitialized(uint256 characterId, uint256 noteId, uint256 timestamp)
func (_Character *CharacterFilterer) ParseMintNFTInitialized(log types.Log) (*CharacterMintNFTInitialized, error) {
	event := new(CharacterMintNFTInitialized)
	if err := _Character.contract.UnpackLog(event, "MintNFTInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterMintNoteIterator is returned from FilterMintNote and is used to iterate over the raw logs and unpacked data for MintNote events raised by the Character contract.
type CharacterMintNoteIterator struct {
	Event *CharacterMintNote // Event containing the contract specifics and raw log

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
func (it *CharacterMintNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterMintNote)
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
		it.Event = new(CharacterMintNote)
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
func (it *CharacterMintNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterMintNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterMintNote represents a MintNote event raised by the Character contract.
type CharacterMintNote struct {
	To           common.Address
	CharacterId  *big.Int
	NoteId       *big.Int
	TokenAddress common.Address
	TokenId      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintNote is a free log retrieval operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Character *CharacterFilterer) FilterMintNote(opts *bind.FilterOpts, to []common.Address, characterId []*big.Int, noteId []*big.Int) (*CharacterMintNoteIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "MintNote", toRule, characterIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterMintNoteIterator{contract: _Character.contract, event: "MintNote", logs: logs, sub: sub}, nil
}

// WatchMintNote is a free log subscription operation binding the contract event 0x6f1704cf3bc1cfc1bc2284eee0ba541a208bc80989f559ed39cc6567d77cf212.
//
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Character *CharacterFilterer) WatchMintNote(opts *bind.WatchOpts, sink chan<- *CharacterMintNote, to []common.Address, characterId []*big.Int, noteId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "MintNote", toRule, characterIdRule, noteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterMintNote)
				if err := _Character.contract.UnpackLog(event, "MintNote", log); err != nil {
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
// Solidity: event MintNote(address indexed to, uint256 indexed characterId, uint256 indexed noteId, address tokenAddress, uint256 tokenId)
func (_Character *CharacterFilterer) ParseMintNote(log types.Log) (*CharacterMintNote, error) {
	event := new(CharacterMintNote)
	if err := _Character.contract.UnpackLog(event, "MintNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterPostNoteIterator is returned from FilterPostNote and is used to iterate over the raw logs and unpacked data for PostNote events raised by the Character contract.
type CharacterPostNoteIterator struct {
	Event *CharacterPostNote // Event containing the contract specifics and raw log

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
func (it *CharacterPostNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterPostNote)
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
		it.Event = new(CharacterPostNote)
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
func (it *CharacterPostNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterPostNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterPostNote represents a PostNote event raised by the Character contract.
type CharacterPostNote struct {
	CharacterId  *big.Int
	NoteId       *big.Int
	LinkKey      [32]byte
	LinkItemType [32]byte
	Data         []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPostNote is a free log retrieval operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Character *CharacterFilterer) FilterPostNote(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (*CharacterPostNoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "PostNote", characterIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return &CharacterPostNoteIterator{contract: _Character.contract, event: "PostNote", logs: logs, sub: sub}, nil
}

// WatchPostNote is a free log subscription operation binding the contract event 0x6ea6daa2449ded342bb92186e54e576ec7c6a729d4ccbf6f364e1bd1f1a52740.
//
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Character *CharacterFilterer) WatchPostNote(opts *bind.WatchOpts, sink chan<- *CharacterPostNote, characterId []*big.Int, noteId []*big.Int, linkKey [][32]byte) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkKeyRule []interface{}
	for _, linkKeyItem := range linkKey {
		linkKeyRule = append(linkKeyRule, linkKeyItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "PostNote", characterIdRule, noteIdRule, linkKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterPostNote)
				if err := _Character.contract.UnpackLog(event, "PostNote", log); err != nil {
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
// Solidity: event PostNote(uint256 indexed characterId, uint256 indexed noteId, bytes32 indexed linkKey, bytes32 linkItemType, bytes data)
func (_Character *CharacterFilterer) ParsePostNote(log types.Log) (*CharacterPostNote, error) {
	event := new(CharacterPostNote)
	if err := _Character.contract.UnpackLog(event, "PostNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetCharacterUriIterator is returned from FilterSetCharacterUri and is used to iterate over the raw logs and unpacked data for SetCharacterUri events raised by the Character contract.
type CharacterSetCharacterUriIterator struct {
	Event *CharacterSetCharacterUri // Event containing the contract specifics and raw log

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
func (it *CharacterSetCharacterUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetCharacterUri)
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
		it.Event = new(CharacterSetCharacterUri)
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
func (it *CharacterSetCharacterUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetCharacterUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetCharacterUri represents a SetCharacterUri event raised by the Character contract.
type CharacterSetCharacterUri struct {
	CharacterId *big.Int
	NewUri      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetCharacterUri is a free log retrieval operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Character *CharacterFilterer) FilterSetCharacterUri(opts *bind.FilterOpts, characterId []*big.Int) (*CharacterSetCharacterUriIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetCharacterUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetCharacterUriIterator{contract: _Character.contract, event: "SetCharacterUri", logs: logs, sub: sub}, nil
}

// WatchSetCharacterUri is a free log subscription operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Character *CharacterFilterer) WatchSetCharacterUri(opts *bind.WatchOpts, sink chan<- *CharacterSetCharacterUri, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetCharacterUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetCharacterUri)
				if err := _Character.contract.UnpackLog(event, "SetCharacterUri", log); err != nil {
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

// ParseSetCharacterUri is a log parse operation binding the contract event 0x17d7c9f69270ba135480ef16837f38b9d37d3ab291cbd3ba03982290c6631997.
//
// Solidity: event SetCharacterUri(uint256 indexed characterId, string newUri)
func (_Character *CharacterFilterer) ParseSetCharacterUri(log types.Log) (*CharacterSetCharacterUri, error) {
	event := new(CharacterSetCharacterUri)
	if err := _Character.contract.UnpackLog(event, "SetCharacterUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetHandleIterator is returned from FilterSetHandle and is used to iterate over the raw logs and unpacked data for SetHandle events raised by the Character contract.
type CharacterSetHandleIterator struct {
	Event *CharacterSetHandle // Event containing the contract specifics and raw log

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
func (it *CharacterSetHandleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetHandle)
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
		it.Event = new(CharacterSetHandle)
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
func (it *CharacterSetHandleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetHandleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetHandle represents a SetHandle event raised by the Character contract.
type CharacterSetHandle struct {
	Account     common.Address
	CharacterId *big.Int
	NewHandle   string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetHandle is a free log retrieval operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Character *CharacterFilterer) FilterSetHandle(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int) (*CharacterSetHandleIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetHandle", accountRule, characterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetHandleIterator{contract: _Character.contract, event: "SetHandle", logs: logs, sub: sub}, nil
}

// WatchSetHandle is a free log subscription operation binding the contract event 0x9d175e377ed0c2f6df0cefe4efe069e62eaa3dad046bb8c88e293a776c3cf96e.
//
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Character *CharacterFilterer) WatchSetHandle(opts *bind.WatchOpts, sink chan<- *CharacterSetHandle, account []common.Address, characterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetHandle", accountRule, characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetHandle)
				if err := _Character.contract.UnpackLog(event, "SetHandle", log); err != nil {
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
// Solidity: event SetHandle(address indexed account, uint256 indexed characterId, string newHandle)
func (_Character *CharacterFilterer) ParseSetHandle(log types.Log) (*CharacterSetHandle, error) {
	event := new(CharacterSetHandle)
	if err := _Character.contract.UnpackLog(event, "SetHandle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetLinkModule4AddressIterator is returned from FilterSetLinkModule4Address and is used to iterate over the raw logs and unpacked data for SetLinkModule4Address events raised by the Character contract.
type CharacterSetLinkModule4AddressIterator struct {
	Event *CharacterSetLinkModule4Address // Event containing the contract specifics and raw log

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
func (it *CharacterSetLinkModule4AddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetLinkModule4Address)
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
		it.Event = new(CharacterSetLinkModule4Address)
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
func (it *CharacterSetLinkModule4AddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetLinkModule4AddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetLinkModule4Address represents a SetLinkModule4Address event raised by the Character contract.
type CharacterSetLinkModule4Address struct {
	Account    common.Address
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Address is a free log retrieval operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetLinkModule4Address(opts *bind.FilterOpts, account []common.Address, linkModule []common.Address) (*CharacterSetLinkModule4AddressIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetLinkModule4AddressIterator{contract: _Character.contract, event: "SetLinkModule4Address", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Address is a free log subscription operation binding the contract event 0x7685796f9ddd10aa092582edf9c2a9ea78db9685e6f997342b6ab22268a730d8.
//
// Solidity: event SetLinkModule4Address(address indexed account, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetLinkModule4Address(opts *bind.WatchOpts, sink chan<- *CharacterSetLinkModule4Address, account []common.Address, linkModule []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetLinkModule4Address", accountRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetLinkModule4Address)
				if err := _Character.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
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
func (_Character *CharacterFilterer) ParseSetLinkModule4Address(log types.Log) (*CharacterSetLinkModule4Address, error) {
	event := new(CharacterSetLinkModule4Address)
	if err := _Character.contract.UnpackLog(event, "SetLinkModule4Address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetLinkModule4CharacterIterator is returned from FilterSetLinkModule4Character and is used to iterate over the raw logs and unpacked data for SetLinkModule4Character events raised by the Character contract.
type CharacterSetLinkModule4CharacterIterator struct {
	Event *CharacterSetLinkModule4Character // Event containing the contract specifics and raw log

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
func (it *CharacterSetLinkModule4CharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetLinkModule4Character)
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
		it.Event = new(CharacterSetLinkModule4Character)
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
func (it *CharacterSetLinkModule4CharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetLinkModule4CharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetLinkModule4Character represents a SetLinkModule4Character event raised by the Character contract.
type CharacterSetLinkModule4Character struct {
	CharacterId *big.Int
	LinkModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Character is a free log retrieval operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetLinkModule4Character(opts *bind.FilterOpts, characterId []*big.Int, linkModule []common.Address) (*CharacterSetLinkModule4CharacterIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetLinkModule4Character", characterIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetLinkModule4CharacterIterator{contract: _Character.contract, event: "SetLinkModule4Character", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Character is a free log subscription operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetLinkModule4Character(opts *bind.WatchOpts, sink chan<- *CharacterSetLinkModule4Character, characterId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetLinkModule4Character", characterIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetLinkModule4Character)
				if err := _Character.contract.UnpackLog(event, "SetLinkModule4Character", log); err != nil {
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

// ParseSetLinkModule4Character is a log parse operation binding the contract event 0x0b73469075e92f46cac48635c7f660883558cc7886309d409da73dea2f56b61a.
//
// Solidity: event SetLinkModule4Character(uint256 indexed characterId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) ParseSetLinkModule4Character(log types.Log) (*CharacterSetLinkModule4Character, error) {
	event := new(CharacterSetLinkModule4Character)
	if err := _Character.contract.UnpackLog(event, "SetLinkModule4Character", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetLinkModule4ERC721Iterator is returned from FilterSetLinkModule4ERC721 and is used to iterate over the raw logs and unpacked data for SetLinkModule4ERC721 events raised by the Character contract.
type CharacterSetLinkModule4ERC721Iterator struct {
	Event *CharacterSetLinkModule4ERC721 // Event containing the contract specifics and raw log

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
func (it *CharacterSetLinkModule4ERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetLinkModule4ERC721)
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
		it.Event = new(CharacterSetLinkModule4ERC721)
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
func (it *CharacterSetLinkModule4ERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetLinkModule4ERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetLinkModule4ERC721 represents a SetLinkModule4ERC721 event raised by the Character contract.
type CharacterSetLinkModule4ERC721 struct {
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
func (_Character *CharacterFilterer) FilterSetLinkModule4ERC721(opts *bind.FilterOpts, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (*CharacterSetLinkModule4ERC721Iterator, error) {

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

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetLinkModule4ERC721Iterator{contract: _Character.contract, event: "SetLinkModule4ERC721", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4ERC721 is a free log subscription operation binding the contract event 0xd0411ae508eec872740a07b3a8da69f9a925547a40bbbdb612971a050c61e19e.
//
// Solidity: event SetLinkModule4ERC721(address indexed tokenAddress, uint256 indexed tokenId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetLinkModule4ERC721(opts *bind.WatchOpts, sink chan<- *CharacterSetLinkModule4ERC721, tokenAddress []common.Address, tokenId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetLinkModule4ERC721", tokenAddressRule, tokenIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetLinkModule4ERC721)
				if err := _Character.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
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
func (_Character *CharacterFilterer) ParseSetLinkModule4ERC721(log types.Log) (*CharacterSetLinkModule4ERC721, error) {
	event := new(CharacterSetLinkModule4ERC721)
	if err := _Character.contract.UnpackLog(event, "SetLinkModule4ERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetLinkModule4LinklistIterator is returned from FilterSetLinkModule4Linklist and is used to iterate over the raw logs and unpacked data for SetLinkModule4Linklist events raised by the Character contract.
type CharacterSetLinkModule4LinklistIterator struct {
	Event *CharacterSetLinkModule4Linklist // Event containing the contract specifics and raw log

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
func (it *CharacterSetLinkModule4LinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetLinkModule4Linklist)
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
		it.Event = new(CharacterSetLinkModule4Linklist)
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
func (it *CharacterSetLinkModule4LinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetLinkModule4LinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetLinkModule4Linklist represents a SetLinkModule4Linklist event raised by the Character contract.
type CharacterSetLinkModule4Linklist struct {
	LinklistId *big.Int
	LinkModule common.Address
	ReturnData []byte
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Linklist is a free log retrieval operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetLinkModule4Linklist(opts *bind.FilterOpts, linklistId []*big.Int, linkModule []common.Address) (*CharacterSetLinkModule4LinklistIterator, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetLinkModule4LinklistIterator{contract: _Character.contract, event: "SetLinkModule4Linklist", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Linklist is a free log subscription operation binding the contract event 0x63dbee1d4ec714c8d35de5e060e27c372b6a409081cdb917c86ea48fdad89b4b.
//
// Solidity: event SetLinkModule4Linklist(uint256 indexed linklistId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetLinkModule4Linklist(opts *bind.WatchOpts, sink chan<- *CharacterSetLinkModule4Linklist, linklistId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetLinkModule4Linklist", linklistIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetLinkModule4Linklist)
				if err := _Character.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
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
func (_Character *CharacterFilterer) ParseSetLinkModule4Linklist(log types.Log) (*CharacterSetLinkModule4Linklist, error) {
	event := new(CharacterSetLinkModule4Linklist)
	if err := _Character.contract.UnpackLog(event, "SetLinkModule4Linklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetLinkModule4NoteIterator is returned from FilterSetLinkModule4Note and is used to iterate over the raw logs and unpacked data for SetLinkModule4Note events raised by the Character contract.
type CharacterSetLinkModule4NoteIterator struct {
	Event *CharacterSetLinkModule4Note // Event containing the contract specifics and raw log

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
func (it *CharacterSetLinkModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetLinkModule4Note)
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
		it.Event = new(CharacterSetLinkModule4Note)
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
func (it *CharacterSetLinkModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetLinkModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetLinkModule4Note represents a SetLinkModule4Note event raised by the Character contract.
type CharacterSetLinkModule4Note struct {
	CharacterId *big.Int
	NoteId      *big.Int
	LinkModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetLinkModule4Note is a free log retrieval operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetLinkModule4Note(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, linkModule []common.Address) (*CharacterSetLinkModule4NoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetLinkModule4Note", characterIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetLinkModule4NoteIterator{contract: _Character.contract, event: "SetLinkModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetLinkModule4Note is a free log subscription operation binding the contract event 0x889c6317f76b8527935ed434226767d05f8b7c664d99f6f787e62efd558f6f01.
//
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetLinkModule4Note(opts *bind.WatchOpts, sink chan<- *CharacterSetLinkModule4Note, characterId []*big.Int, noteId []*big.Int, linkModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var linkModuleRule []interface{}
	for _, linkModuleItem := range linkModule {
		linkModuleRule = append(linkModuleRule, linkModuleItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetLinkModule4Note", characterIdRule, noteIdRule, linkModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetLinkModule4Note)
				if err := _Character.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
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
// Solidity: event SetLinkModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed linkModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) ParseSetLinkModule4Note(log types.Log) (*CharacterSetLinkModule4Note, error) {
	event := new(CharacterSetLinkModule4Note)
	if err := _Character.contract.UnpackLog(event, "SetLinkModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetMintModule4NoteIterator is returned from FilterSetMintModule4Note and is used to iterate over the raw logs and unpacked data for SetMintModule4Note events raised by the Character contract.
type CharacterSetMintModule4NoteIterator struct {
	Event *CharacterSetMintModule4Note // Event containing the contract specifics and raw log

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
func (it *CharacterSetMintModule4NoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetMintModule4Note)
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
		it.Event = new(CharacterSetMintModule4Note)
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
func (it *CharacterSetMintModule4NoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetMintModule4NoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetMintModule4Note represents a SetMintModule4Note event raised by the Character contract.
type CharacterSetMintModule4Note struct {
	CharacterId *big.Int
	NoteId      *big.Int
	MintModule  common.Address
	ReturnData  []byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetMintModule4Note is a free log retrieval operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetMintModule4Note(opts *bind.FilterOpts, characterId []*big.Int, noteId []*big.Int, mintModule []common.Address) (*CharacterSetMintModule4NoteIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetMintModule4Note", characterIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetMintModule4NoteIterator{contract: _Character.contract, event: "SetMintModule4Note", logs: logs, sub: sub}, nil
}

// WatchSetMintModule4Note is a free log subscription operation binding the contract event 0x36e973ebf2a1c9c4006aaad244866e6dea9a0e85770deea599b193a9eb51b8f7.
//
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetMintModule4Note(opts *bind.WatchOpts, sink chan<- *CharacterSetMintModule4Note, characterId []*big.Int, noteId []*big.Int, mintModule []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var noteIdRule []interface{}
	for _, noteIdItem := range noteId {
		noteIdRule = append(noteIdRule, noteIdItem)
	}
	var mintModuleRule []interface{}
	for _, mintModuleItem := range mintModule {
		mintModuleRule = append(mintModuleRule, mintModuleItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetMintModule4Note", characterIdRule, noteIdRule, mintModuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetMintModule4Note)
				if err := _Character.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
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
// Solidity: event SetMintModule4Note(uint256 indexed characterId, uint256 indexed noteId, address indexed mintModule, bytes returnData, uint256 timestamp)
func (_Character *CharacterFilterer) ParseSetMintModule4Note(log types.Log) (*CharacterSetMintModule4Note, error) {
	event := new(CharacterSetMintModule4Note)
	if err := _Character.contract.UnpackLog(event, "SetMintModule4Note", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetNoteUriIterator is returned from FilterSetNoteUri and is used to iterate over the raw logs and unpacked data for SetNoteUri events raised by the Character contract.
type CharacterSetNoteUriIterator struct {
	Event *CharacterSetNoteUri // Event containing the contract specifics and raw log

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
func (it *CharacterSetNoteUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetNoteUri)
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
		it.Event = new(CharacterSetNoteUri)
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
func (it *CharacterSetNoteUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetNoteUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetNoteUri represents a SetNoteUri event raised by the Character contract.
type CharacterSetNoteUri struct {
	CharacterId *big.Int
	NoteId      *big.Int
	NewUri      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetNoteUri is a free log retrieval operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Character *CharacterFilterer) FilterSetNoteUri(opts *bind.FilterOpts, characterId []*big.Int) (*CharacterSetNoteUriIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetNoteUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetNoteUriIterator{contract: _Character.contract, event: "SetNoteUri", logs: logs, sub: sub}, nil
}

// WatchSetNoteUri is a free log subscription operation binding the contract event 0x179143dd0d35a50f482efc003aa5b84a0c3a40d74e9cc2d7ea3053b9e8c37697.
//
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Character *CharacterFilterer) WatchSetNoteUri(opts *bind.WatchOpts, sink chan<- *CharacterSetNoteUri, characterId []*big.Int) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetNoteUri", characterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetNoteUri)
				if err := _Character.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
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
// Solidity: event SetNoteUri(uint256 indexed characterId, uint256 noteId, string newUri)
func (_Character *CharacterFilterer) ParseSetNoteUri(log types.Log) (*CharacterSetNoteUri, error) {
	event := new(CharacterSetNoteUri)
	if err := _Character.contract.UnpackLog(event, "SetNoteUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the Character contract.
type CharacterSetOperatorIterator struct {
	Event *CharacterSetOperator // Event containing the contract specifics and raw log

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
func (it *CharacterSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetOperator)
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
		it.Event = new(CharacterSetOperator)
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
func (it *CharacterSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetOperator represents a SetOperator event raised by the Character contract.
type CharacterSetOperator struct {
	CharacterId *big.Int
	Operator    common.Address
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Character *CharacterFilterer) FilterSetOperator(opts *bind.FilterOpts, characterId []*big.Int, operator []common.Address) (*CharacterSetOperatorIterator, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetOperatorIterator{contract: _Character.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0x691b92a93c526c4f5a357e56a33c33f6250f94e19e6c49be805800615c414931.
//
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Character *CharacterFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *CharacterSetOperator, characterId []*big.Int, operator []common.Address) (event.Subscription, error) {

	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetOperator", characterIdRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetOperator)
				if err := _Character.contract.UnpackLog(event, "SetOperator", log); err != nil {
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
// Solidity: event SetOperator(uint256 indexed characterId, address indexed operator, uint256 timestamp)
func (_Character *CharacterFilterer) ParseSetOperator(log types.Log) (*CharacterSetOperator, error) {
	event := new(CharacterSetOperator)
	if err := _Character.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetPrimaryCharacterIdIterator is returned from FilterSetPrimaryCharacterId and is used to iterate over the raw logs and unpacked data for SetPrimaryCharacterId events raised by the Character contract.
type CharacterSetPrimaryCharacterIdIterator struct {
	Event *CharacterSetPrimaryCharacterId // Event containing the contract specifics and raw log

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
func (it *CharacterSetPrimaryCharacterIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetPrimaryCharacterId)
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
		it.Event = new(CharacterSetPrimaryCharacterId)
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
func (it *CharacterSetPrimaryCharacterIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetPrimaryCharacterIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetPrimaryCharacterId represents a SetPrimaryCharacterId event raised by the Character contract.
type CharacterSetPrimaryCharacterId struct {
	Account        common.Address
	CharacterId    *big.Int
	OldCharacterId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetPrimaryCharacterId is a free log retrieval operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Character *CharacterFilterer) FilterSetPrimaryCharacterId(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int, oldCharacterId []*big.Int) (*CharacterSetPrimaryCharacterIdIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var oldCharacterIdRule []interface{}
	for _, oldCharacterIdItem := range oldCharacterId {
		oldCharacterIdRule = append(oldCharacterIdRule, oldCharacterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetPrimaryCharacterId", accountRule, characterIdRule, oldCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetPrimaryCharacterIdIterator{contract: _Character.contract, event: "SetPrimaryCharacterId", logs: logs, sub: sub}, nil
}

// WatchSetPrimaryCharacterId is a free log subscription operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Character *CharacterFilterer) WatchSetPrimaryCharacterId(opts *bind.WatchOpts, sink chan<- *CharacterSetPrimaryCharacterId, account []common.Address, characterId []*big.Int, oldCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var oldCharacterIdRule []interface{}
	for _, oldCharacterIdItem := range oldCharacterId {
		oldCharacterIdRule = append(oldCharacterIdRule, oldCharacterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetPrimaryCharacterId", accountRule, characterIdRule, oldCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetPrimaryCharacterId)
				if err := _Character.contract.UnpackLog(event, "SetPrimaryCharacterId", log); err != nil {
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

// ParseSetPrimaryCharacterId is a log parse operation binding the contract event 0xce95332e6082aebeb8058a7b56d1a109f67d6550552ed04d36aca4a6acd4d7de.
//
// Solidity: event SetPrimaryCharacterId(address indexed account, uint256 indexed characterId, uint256 indexed oldCharacterId)
func (_Character *CharacterFilterer) ParseSetPrimaryCharacterId(log types.Log) (*CharacterSetPrimaryCharacterId, error) {
	event := new(CharacterSetPrimaryCharacterId)
	if err := _Character.contract.UnpackLog(event, "SetPrimaryCharacterId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterSetSocialTokenIterator is returned from FilterSetSocialToken and is used to iterate over the raw logs and unpacked data for SetSocialToken events raised by the Character contract.
type CharacterSetSocialTokenIterator struct {
	Event *CharacterSetSocialToken // Event containing the contract specifics and raw log

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
func (it *CharacterSetSocialTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterSetSocialToken)
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
		it.Event = new(CharacterSetSocialToken)
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
func (it *CharacterSetSocialTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterSetSocialTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterSetSocialToken represents a SetSocialToken event raised by the Character contract.
type CharacterSetSocialToken struct {
	Account      common.Address
	CharacterId  *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetSocialToken is a free log retrieval operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Character *CharacterFilterer) FilterSetSocialToken(opts *bind.FilterOpts, account []common.Address, characterId []*big.Int, tokenAddress []common.Address) (*CharacterSetSocialTokenIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "SetSocialToken", accountRule, characterIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharacterSetSocialTokenIterator{contract: _Character.contract, event: "SetSocialToken", logs: logs, sub: sub}, nil
}

// WatchSetSocialToken is a free log subscription operation binding the contract event 0x276c2d4b7f7beaa1847ff9c41b9d2e0e57482efe8051eea98eea653bfca3ac45.
//
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Character *CharacterFilterer) WatchSetSocialToken(opts *bind.WatchOpts, sink chan<- *CharacterSetSocialToken, account []common.Address, characterId []*big.Int, tokenAddress []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var characterIdRule []interface{}
	for _, characterIdItem := range characterId {
		characterIdRule = append(characterIdRule, characterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "SetSocialToken", accountRule, characterIdRule, tokenAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterSetSocialToken)
				if err := _Character.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
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
// Solidity: event SetSocialToken(address indexed account, uint256 indexed characterId, address indexed tokenAddress)
func (_Character *CharacterFilterer) ParseSetSocialToken(log types.Log) (*CharacterSetSocialToken, error) {
	event := new(CharacterSetSocialToken)
	if err := _Character.contract.UnpackLog(event, "SetSocialToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Character contract.
type CharacterTransferIterator struct {
	Event *CharacterTransfer // Event containing the contract specifics and raw log

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
func (it *CharacterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterTransfer)
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
		it.Event = new(CharacterTransfer)
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
func (it *CharacterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterTransfer represents a Transfer event raised by the Character contract.
type CharacterTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Character *CharacterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CharacterTransferIterator, error) {

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

	logs, sub, err := _Character.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterTransferIterator{contract: _Character.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Character *CharacterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CharacterTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Character.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterTransfer)
				if err := _Character.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Character *CharacterFilterer) ParseTransfer(log types.Log) (*CharacterTransfer, error) {
	event := new(CharacterTransfer)
	if err := _Character.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkAddressIterator is returned from FilterUnlinkAddress and is used to iterate over the raw logs and unpacked data for UnlinkAddress events raised by the Character contract.
type CharacterUnlinkAddressIterator struct {
	Event *CharacterUnlinkAddress // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkAddress)
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
		it.Event = new(CharacterUnlinkAddress)
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
func (it *CharacterUnlinkAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkAddress represents a UnlinkAddress event raised by the Character contract.
type CharacterUnlinkAddress struct {
	FromCharacterId *big.Int
	EthAddress      common.Address
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAddress is a free log retrieval operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Character *CharacterFilterer) FilterUnlinkAddress(opts *bind.FilterOpts, fromCharacterId []*big.Int, ethAddress []common.Address) (*CharacterUnlinkAddressIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkAddressIterator{contract: _Character.contract, event: "UnlinkAddress", logs: logs, sub: sub}, nil
}

// WatchUnlinkAddress is a free log subscription operation binding the contract event 0x93453451dd1d041ffa18c6c1f3f2e21a6d73c3d8d32deaf595b53a14914c8715.
//
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Character *CharacterFilterer) WatchUnlinkAddress(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkAddress, fromCharacterId []*big.Int, ethAddress []common.Address) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var ethAddressRule []interface{}
	for _, ethAddressItem := range ethAddress {
		ethAddressRule = append(ethAddressRule, ethAddressItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkAddress", fromCharacterIdRule, ethAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkAddress)
				if err := _Character.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
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
// Solidity: event UnlinkAddress(uint256 indexed fromCharacterId, address indexed ethAddress, bytes32 linkType)
func (_Character *CharacterFilterer) ParseUnlinkAddress(log types.Log) (*CharacterUnlinkAddress, error) {
	event := new(CharacterUnlinkAddress)
	if err := _Character.contract.UnpackLog(event, "UnlinkAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkAnyUriIterator is returned from FilterUnlinkAnyUri and is used to iterate over the raw logs and unpacked data for UnlinkAnyUri events raised by the Character contract.
type CharacterUnlinkAnyUriIterator struct {
	Event *CharacterUnlinkAnyUri // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkAnyUriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkAnyUri)
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
		it.Event = new(CharacterUnlinkAnyUri)
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
func (it *CharacterUnlinkAnyUriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkAnyUriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkAnyUri represents a UnlinkAnyUri event raised by the Character contract.
type CharacterUnlinkAnyUri struct {
	FromCharacterId *big.Int
	ToUri           string
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkAnyUri is a free log retrieval operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Character *CharacterFilterer) FilterUnlinkAnyUri(opts *bind.FilterOpts, fromCharacterId []*big.Int) (*CharacterUnlinkAnyUriIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkAnyUriIterator{contract: _Character.contract, event: "UnlinkAnyUri", logs: logs, sub: sub}, nil
}

// WatchUnlinkAnyUri is a free log subscription operation binding the contract event 0x1be04341b458d762379a09acac0df7b19945e7c12a63405d218df896e9bf7887.
//
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Character *CharacterFilterer) WatchUnlinkAnyUri(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkAnyUri, fromCharacterId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkAnyUri", fromCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkAnyUri)
				if err := _Character.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
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
// Solidity: event UnlinkAnyUri(uint256 indexed fromCharacterId, string toUri, bytes32 linkType)
func (_Character *CharacterFilterer) ParseUnlinkAnyUri(log types.Log) (*CharacterUnlinkAnyUri, error) {
	event := new(CharacterUnlinkAnyUri)
	if err := _Character.contract.UnpackLog(event, "UnlinkAnyUri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkCharacterIterator is returned from FilterUnlinkCharacter and is used to iterate over the raw logs and unpacked data for UnlinkCharacter events raised by the Character contract.
type CharacterUnlinkCharacterIterator struct {
	Event *CharacterUnlinkCharacter // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkCharacterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkCharacter)
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
		it.Event = new(CharacterUnlinkCharacter)
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
func (it *CharacterUnlinkCharacterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkCharacterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkCharacter represents a UnlinkCharacter event raised by the Character contract.
type CharacterUnlinkCharacter struct {
	Account         common.Address
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	LinkType        [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkCharacter is a free log retrieval operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Character *CharacterFilterer) FilterUnlinkCharacter(opts *bind.FilterOpts, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (*CharacterUnlinkCharacterIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkCharacterIterator{contract: _Character.contract, event: "UnlinkCharacter", logs: logs, sub: sub}, nil
}

// WatchUnlinkCharacter is a free log subscription operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Character *CharacterFilterer) WatchUnlinkCharacter(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkCharacter, account []common.Address, fromCharacterId []*big.Int, toCharacterId []*big.Int) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkCharacter", accountRule, fromCharacterIdRule, toCharacterIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkCharacter)
				if err := _Character.contract.UnpackLog(event, "UnlinkCharacter", log); err != nil {
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

// ParseUnlinkCharacter is a log parse operation binding the contract event 0x056af39d4e7ea448ab8c74ce7a5ccb3940d3df73fc086846fd57acd615ed9e94.
//
// Solidity: event UnlinkCharacter(address indexed account, uint256 indexed fromCharacterId, uint256 indexed toCharacterId, bytes32 linkType)
func (_Character *CharacterFilterer) ParseUnlinkCharacter(log types.Log) (*CharacterUnlinkCharacter, error) {
	event := new(CharacterUnlinkCharacter)
	if err := _Character.contract.UnpackLog(event, "UnlinkCharacter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkCharacterLinkIterator is returned from FilterUnlinkCharacterLink and is used to iterate over the raw logs and unpacked data for UnlinkCharacterLink events raised by the Character contract.
type CharacterUnlinkCharacterLinkIterator struct {
	Event *CharacterUnlinkCharacterLink // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkCharacterLinkIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkCharacterLink)
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
		it.Event = new(CharacterUnlinkCharacterLink)
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
func (it *CharacterUnlinkCharacterLinkIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkCharacterLinkIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkCharacterLink represents a UnlinkCharacterLink event raised by the Character contract.
type CharacterUnlinkCharacterLink struct {
	FromCharacterId    *big.Int
	LinkType           [32]byte
	ClFromCharactereId *big.Int
	ClToCharacterId    *big.Int
	ClLinkType         [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterUnlinkCharacterLink is a free log retrieval operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) FilterUnlinkCharacterLink(opts *bind.FilterOpts, fromCharacterId []*big.Int, linkType [][32]byte) (*CharacterUnlinkCharacterLinkIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkCharacterLinkIterator{contract: _Character.contract, event: "UnlinkCharacterLink", logs: logs, sub: sub}, nil
}

// WatchUnlinkCharacterLink is a free log subscription operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) WatchUnlinkCharacterLink(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkCharacterLink, fromCharacterId []*big.Int, linkType [][32]byte) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var linkTypeRule []interface{}
	for _, linkTypeItem := range linkType {
		linkTypeRule = append(linkTypeRule, linkTypeItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkCharacterLink", fromCharacterIdRule, linkTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkCharacterLink)
				if err := _Character.contract.UnpackLog(event, "UnlinkCharacterLink", log); err != nil {
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

// ParseUnlinkCharacterLink is a log parse operation binding the contract event 0xb027fe4b36accc4e88bb9af4c96eac3e299e963e941a62415b15e0e070620686.
//
// Solidity: event UnlinkCharacterLink(uint256 indexed fromCharacterId, bytes32 indexed linkType, uint256 clFromCharactereId, uint256 clToCharacterId, bytes32 clLinkType)
func (_Character *CharacterFilterer) ParseUnlinkCharacterLink(log types.Log) (*CharacterUnlinkCharacterLink, error) {
	event := new(CharacterUnlinkCharacterLink)
	if err := _Character.contract.UnpackLog(event, "UnlinkCharacterLink", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkERC721Iterator is returned from FilterUnlinkERC721 and is used to iterate over the raw logs and unpacked data for UnlinkERC721 events raised by the Character contract.
type CharacterUnlinkERC721Iterator struct {
	Event *CharacterUnlinkERC721 // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkERC721Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkERC721)
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
		it.Event = new(CharacterUnlinkERC721)
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
func (it *CharacterUnlinkERC721Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkERC721Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkERC721 represents a UnlinkERC721 event raised by the Character contract.
type CharacterUnlinkERC721 struct {
	FromCharacterId *big.Int
	TokenAddress    common.Address
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkERC721 is a free log retrieval operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterUnlinkERC721(opts *bind.FilterOpts, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (*CharacterUnlinkERC721Iterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkERC721Iterator{contract: _Character.contract, event: "UnlinkERC721", logs: logs, sub: sub}, nil
}

// WatchUnlinkERC721 is a free log subscription operation binding the contract event 0xd87f9606a19988b6cf42d250d484940673ce551ab5f80289051cc343ff13121c.
//
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchUnlinkERC721(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkERC721, fromCharacterId []*big.Int, tokenAddress []common.Address, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkERC721", fromCharacterIdRule, tokenAddressRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkERC721)
				if err := _Character.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
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
// Solidity: event UnlinkERC721(uint256 indexed fromCharacterId, address indexed tokenAddress, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseUnlinkERC721(log types.Log) (*CharacterUnlinkERC721, error) {
	event := new(CharacterUnlinkERC721)
	if err := _Character.contract.UnpackLog(event, "UnlinkERC721", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkLinklistIterator is returned from FilterUnlinkLinklist and is used to iterate over the raw logs and unpacked data for UnlinkLinklist events raised by the Character contract.
type CharacterUnlinkLinklistIterator struct {
	Event *CharacterUnlinkLinklist // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkLinklistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkLinklist)
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
		it.Event = new(CharacterUnlinkLinklist)
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
func (it *CharacterUnlinkLinklistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkLinklistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkLinklist represents a UnlinkLinklist event raised by the Character contract.
type CharacterUnlinkLinklist struct {
	FromCharacterId *big.Int
	ToLinklistId    *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkLinklist is a free log retrieval operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) FilterUnlinkLinklist(opts *bind.FilterOpts, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (*CharacterUnlinkLinklistIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkLinklistIterator{contract: _Character.contract, event: "UnlinkLinklist", logs: logs, sub: sub}, nil
}

// WatchUnlinkLinklist is a free log subscription operation binding the contract event 0x42b4ce79acc0bdbfa79f30ba8758f3a465824adff1ea290b6d5aeeeef05eb14f.
//
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) WatchUnlinkLinklist(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkLinklist, fromCharacterId []*big.Int, toLinklistId []*big.Int, linklistId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toLinklistIdRule []interface{}
	for _, toLinklistIdItem := range toLinklistId {
		toLinklistIdRule = append(toLinklistIdRule, toLinklistIdItem)
	}

	var linklistIdRule []interface{}
	for _, linklistIdItem := range linklistId {
		linklistIdRule = append(linklistIdRule, linklistIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkLinklist", fromCharacterIdRule, toLinklistIdRule, linklistIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkLinklist)
				if err := _Character.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
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
// Solidity: event UnlinkLinklist(uint256 indexed fromCharacterId, uint256 indexed toLinklistId, bytes32 linkType, uint256 indexed linklistId)
func (_Character *CharacterFilterer) ParseUnlinkLinklist(log types.Log) (*CharacterUnlinkLinklist, error) {
	event := new(CharacterUnlinkLinklist)
	if err := _Character.contract.UnpackLog(event, "UnlinkLinklist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterUnlinkNoteIterator is returned from FilterUnlinkNote and is used to iterate over the raw logs and unpacked data for UnlinkNote events raised by the Character contract.
type CharacterUnlinkNoteIterator struct {
	Event *CharacterUnlinkNote // Event containing the contract specifics and raw log

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
func (it *CharacterUnlinkNoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterUnlinkNote)
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
		it.Event = new(CharacterUnlinkNote)
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
func (it *CharacterUnlinkNoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterUnlinkNoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterUnlinkNote represents a UnlinkNote event raised by the Character contract.
type CharacterUnlinkNote struct {
	FromCharacterId *big.Int
	ToCharacterId   *big.Int
	ToNoteId        *big.Int
	LinkType        [32]byte
	LinklistId      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUnlinkNote is a free log retrieval operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) FilterUnlinkNote(opts *bind.FilterOpts, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (*CharacterUnlinkNoteIterator, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.FilterLogs(opts, "UnlinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return &CharacterUnlinkNoteIterator{contract: _Character.contract, event: "UnlinkNote", logs: logs, sub: sub}, nil
}

// WatchUnlinkNote is a free log subscription operation binding the contract event 0xd27a71fc88ac85c4657b81c4d24a9cd9a034971683620f179a19d179fe0a956d.
//
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) WatchUnlinkNote(opts *bind.WatchOpts, sink chan<- *CharacterUnlinkNote, fromCharacterId []*big.Int, toCharacterId []*big.Int, toNoteId []*big.Int) (event.Subscription, error) {

	var fromCharacterIdRule []interface{}
	for _, fromCharacterIdItem := range fromCharacterId {
		fromCharacterIdRule = append(fromCharacterIdRule, fromCharacterIdItem)
	}
	var toCharacterIdRule []interface{}
	for _, toCharacterIdItem := range toCharacterId {
		toCharacterIdRule = append(toCharacterIdRule, toCharacterIdItem)
	}
	var toNoteIdRule []interface{}
	for _, toNoteIdItem := range toNoteId {
		toNoteIdRule = append(toNoteIdRule, toNoteIdItem)
	}

	logs, sub, err := _Character.contract.WatchLogs(opts, "UnlinkNote", fromCharacterIdRule, toCharacterIdRule, toNoteIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterUnlinkNote)
				if err := _Character.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
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
// Solidity: event UnlinkNote(uint256 indexed fromCharacterId, uint256 indexed toCharacterId, uint256 indexed toNoteId, bytes32 linkType, uint256 linklistId)
func (_Character *CharacterFilterer) ParseUnlinkNote(log types.Log) (*CharacterUnlinkNote, error) {
	event := new(CharacterUnlinkNote)
	if err := _Character.contract.UnpackLog(event, "UnlinkNote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharacterWeb3EntryInitializedIterator is returned from FilterWeb3EntryInitialized and is used to iterate over the raw logs and unpacked data for Web3EntryInitialized events raised by the Character contract.
type CharacterWeb3EntryInitializedIterator struct {
	Event *CharacterWeb3EntryInitialized // Event containing the contract specifics and raw log

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
func (it *CharacterWeb3EntryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharacterWeb3EntryInitialized)
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
		it.Event = new(CharacterWeb3EntryInitialized)
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
func (it *CharacterWeb3EntryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharacterWeb3EntryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharacterWeb3EntryInitialized represents a Web3EntryInitialized event raised by the Character contract.
type CharacterWeb3EntryInitialized struct {
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWeb3EntryInitialized is a free log retrieval operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Character *CharacterFilterer) FilterWeb3EntryInitialized(opts *bind.FilterOpts) (*CharacterWeb3EntryInitializedIterator, error) {

	logs, sub, err := _Character.contract.FilterLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return &CharacterWeb3EntryInitializedIterator{contract: _Character.contract, event: "Web3EntryInitialized", logs: logs, sub: sub}, nil
}

// WatchWeb3EntryInitialized is a free log subscription operation binding the contract event 0x400175a56dd3710794078f7b9dbe8296ac94c5a248dfd51bb22ed4ab9eaa9fbf.
//
// Solidity: event Web3EntryInitialized(uint256 timestamp)
func (_Character *CharacterFilterer) WatchWeb3EntryInitialized(opts *bind.WatchOpts, sink chan<- *CharacterWeb3EntryInitialized) (event.Subscription, error) {

	logs, sub, err := _Character.contract.WatchLogs(opts, "Web3EntryInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharacterWeb3EntryInitialized)
				if err := _Character.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
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
func (_Character *CharacterFilterer) ParseWeb3EntryInitialized(log types.Log) (*CharacterWeb3EntryInitialized, error) {
	event := new(CharacterWeb3EntryInitialized)
	if err := _Character.contract.UnpackLog(event, "Web3EntryInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
