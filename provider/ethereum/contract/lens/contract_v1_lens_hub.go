// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lens

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

// IERC721TimeTokenData is an auto generated low-level Go binding around an user-defined struct.
type IERC721TimeTokenData struct {
	Owner         common.Address
	MintTimestamp *big.Int
}

// LenHubEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type LenHubEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// LensHubCollectWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubCollectWithSigData struct {
	Collector common.Address
	ProfileId *big.Int
	PubId     *big.Int
	Data      []byte
	Sig       LensHubEIP712Signature
}

// LensHubCommentData is an auto generated low-level Go binding around an user-defined struct.
type LensHubCommentData struct {
	ProfileId               *big.Int
	ContentURI              string
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// LensHubCommentWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubCommentWithSigData struct {
	ProfileId               *big.Int
	ContentURI              string
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     LensHubEIP712Signature
}

// LensHubCreateProfileData is an auto generated low-level Go binding around an user-defined struct.
type LensHubCreateProfileData struct {
	To                   common.Address
	Handle               string
	ImageURI             string
	FollowModule         common.Address
	FollowModuleInitData []byte
	FollowNFTURI         string
}

// LensHubEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type LensHubEIP712Signature struct {
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// LensHubFollowWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubFollowWithSigData struct {
	Follower   common.Address
	ProfileIds []*big.Int
	Datas      [][]byte
	Sig        LensHubEIP712Signature
}

// LensHubMirrorData is an auto generated low-level Go binding around an user-defined struct.
type LensHubMirrorData struct {
	ProfileId               *big.Int
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// LensHubMirrorWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubMirrorWithSigData struct {
	ProfileId               *big.Int
	ProfileIdPointed        *big.Int
	PubIdPointed            *big.Int
	ReferenceModuleData     []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     LensHubEIP712Signature
}

// LensHubPostData is an auto generated low-level Go binding around an user-defined struct.
type LensHubPostData struct {
	ProfileId               *big.Int
	ContentURI              string
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
}

// LensHubPostWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubPostWithSigData struct {
	ProfileId               *big.Int
	ContentURI              string
	CollectModule           common.Address
	CollectModuleInitData   []byte
	ReferenceModule         common.Address
	ReferenceModuleInitData []byte
	Sig                     LensHubEIP712Signature
}

// LensHubProfileStruct is an auto generated low-level Go binding around an user-defined struct.
type LensHubProfileStruct struct {
	PubCount     *big.Int
	FollowModule common.Address
	FollowNFT    common.Address
	Handle       string
	ImageURI     string
	FollowNFTURI string
}

// LensHubPublicationStruct is an auto generated low-level Go binding around an user-defined struct.
type LensHubPublicationStruct struct {
	ProfileIdPointed *big.Int
	PubIdPointed     *big.Int
	ContentURI       string
	ReferenceModule  common.Address
	CollectModule    common.Address
	CollectNFT       common.Address
}

// LensHubSetDefaultProfileWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubSetDefaultProfileWithSigData struct {
	Wallet    common.Address
	ProfileId *big.Int
	Sig       LensHubEIP712Signature
}

// LensHubSetDispatcherWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubSetDispatcherWithSigData struct {
	ProfileId  *big.Int
	Dispatcher common.Address
	Sig        LensHubEIP712Signature
}

// LensHubSetFollowModuleWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubSetFollowModuleWithSigData struct {
	ProfileId            *big.Int
	FollowModule         common.Address
	FollowModuleInitData []byte
	Sig                  LensHubEIP712Signature
}

// LensHubSetFollowNFTURIWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubSetFollowNFTURIWithSigData struct {
	ProfileId    *big.Int
	FollowNFTURI string
	Sig          LensHubEIP712Signature
}

// LensHubSetProfileImageURIWithSigData is an auto generated low-level Go binding around an user-defined struct.
type LensHubSetProfileImageURIWithSigData struct {
	ProfileId *big.Int
	ImageURI  string
	Sig       LensHubEIP712Signature
}

// V1LensHubMetaData contains all meta data concerning the V1LensHub contract.
var V1LensHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFTImpl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallerNotCollectNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotFollowNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInitImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmergencyAdminCannotUnpause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitParamsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernanceOrEmergencyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProfileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProfileOwnerOrDispatcher\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProfileCreatorNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProfileImageURILengthInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicationDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublishingPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLenHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"burnWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"collect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collector\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.CollectWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"collectWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structLensHub.CommentData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"comment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.CommentWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"commentWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structLensHub.CreateProfileData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"defaultProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitCollectNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"followNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitFollowNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"follow\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"follower\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.FollowWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"followWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getCollectNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollectNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getContentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getDispatcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFollowNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getFollowNFTURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getHandle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"internalType\":\"structLensHub.ProfileStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"handle\",\"type\":\"string\"}],\"name\":\"getProfileIdByHandle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPub\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFT\",\"type\":\"address\"}],\"internalType\":\"structLensHub.PublicationStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getPubCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubPointer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPubType\",\"outputs\":[{\"internalType\":\"enumLensHub.PubType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getReferenceModule\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumLensHub.ProtocolState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"}],\"name\":\"isCollectModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"}],\"name\":\"isFollowModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"}],\"name\":\"isProfileCreatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"}],\"name\":\"isReferenceModuleWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mintTimestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structLensHub.MirrorData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirror\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"profileIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubIdPointed\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.MirrorWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"mirrorWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"name\":\"permitForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structLensHub.PostData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"post\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleInitData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.PostWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"postWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"setDefaultProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.SetDefaultProfileWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDefaultProfileWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dispatcher\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.SetDispatcherWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setDispatcherWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"}],\"name\":\"setFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.SetFollowModuleWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowModuleWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"}],\"name\":\"setFollowNFTURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"followNFTURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.SetFollowNFTURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setFollowNFTURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"}],\"name\":\"setProfileImageURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"imageURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLensHub.EIP712Signature\",\"name\":\"sig\",\"type\":\"tuple\"}],\"internalType\":\"structLensHub.SetProfileImageURIWithSigData\",\"name\":\"vars\",\"type\":\"tuple\"}],\"name\":\"setProfileImageURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumLensHub.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sigNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenDataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"mintTimestamp\",\"type\":\"uint96\"}],\"internalType\":\"structIERC721Time.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistCollectModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistProfileCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistReferenceModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V1LensHubABI is the input ABI used to generate the binding from.
// Deprecated: Use V1LensHubMetaData.ABI instead.
var V1LensHubABI = V1LensHubMetaData.ABI

// V1LensHub is an auto generated Go binding around an Ethereum contract.
type V1LensHub struct {
	V1LensHubCaller     // Read-only binding to the contract
	V1LensHubTransactor // Write-only binding to the contract
	V1LensHubFilterer   // Log filterer for contract events
}

// V1LensHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type V1LensHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LensHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V1LensHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LensHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V1LensHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V1LensHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V1LensHubSession struct {
	Contract     *V1LensHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V1LensHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V1LensHubCallerSession struct {
	Contract *V1LensHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// V1LensHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V1LensHubTransactorSession struct {
	Contract     *V1LensHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// V1LensHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type V1LensHubRaw struct {
	Contract *V1LensHub // Generic contract binding to access the raw methods on
}

// V1LensHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V1LensHubCallerRaw struct {
	Contract *V1LensHubCaller // Generic read-only contract binding to access the raw methods on
}

// V1LensHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V1LensHubTransactorRaw struct {
	Contract *V1LensHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV1LensHub creates a new instance of V1LensHub, bound to a specific deployed contract.
func NewV1LensHub(address common.Address, backend bind.ContractBackend) (*V1LensHub, error) {
	contract, err := bindV1LensHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V1LensHub{V1LensHubCaller: V1LensHubCaller{contract: contract}, V1LensHubTransactor: V1LensHubTransactor{contract: contract}, V1LensHubFilterer: V1LensHubFilterer{contract: contract}}, nil
}

// NewV1LensHubCaller creates a new read-only instance of V1LensHub, bound to a specific deployed contract.
func NewV1LensHubCaller(address common.Address, caller bind.ContractCaller) (*V1LensHubCaller, error) {
	contract, err := bindV1LensHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V1LensHubCaller{contract: contract}, nil
}

// NewV1LensHubTransactor creates a new write-only instance of V1LensHub, bound to a specific deployed contract.
func NewV1LensHubTransactor(address common.Address, transactor bind.ContractTransactor) (*V1LensHubTransactor, error) {
	contract, err := bindV1LensHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V1LensHubTransactor{contract: contract}, nil
}

// NewV1LensHubFilterer creates a new log filterer instance of V1LensHub, bound to a specific deployed contract.
func NewV1LensHubFilterer(address common.Address, filterer bind.ContractFilterer) (*V1LensHubFilterer, error) {
	contract, err := bindV1LensHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V1LensHubFilterer{contract: contract}, nil
}

// bindV1LensHub binds a generic wrapper to an already deployed contract.
func bindV1LensHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V1LensHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1LensHub *V1LensHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1LensHub.Contract.V1LensHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1LensHub *V1LensHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1LensHub.Contract.V1LensHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1LensHub *V1LensHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1LensHub.Contract.V1LensHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V1LensHub *V1LensHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V1LensHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V1LensHub *V1LensHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V1LensHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V1LensHub *V1LensHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V1LensHub.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V1LensHub *V1LensHubSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.BalanceOf(&_V1LensHub.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.BalanceOf(&_V1LensHub.CallOpts, owner)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) DefaultProfile(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "defaultProfile", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_V1LensHub *V1LensHubSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.DefaultProfile(&_V1LensHub.CallOpts, wallet)
}

// DefaultProfile is a free data retrieval call binding the contract method 0x92254a62.
//
// Solidity: function defaultProfile(address wallet) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) DefaultProfile(wallet common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.DefaultProfile(&_V1LensHub.CallOpts, wallet)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V1LensHub *V1LensHubCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V1LensHub *V1LensHubSession) Exists(tokenId *big.Int) (bool, error) {
	return _V1LensHub.Contract.Exists(&_V1LensHub.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _V1LensHub.Contract.Exists(&_V1LensHub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetApproved(&_V1LensHub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetApproved(&_V1LensHub.CallOpts, tokenId)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetCollectModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getCollectModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetCollectModule(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetCollectModule is a free data retrieval call binding the contract method 0x57ff49f6.
//
// Solidity: function getCollectModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetCollectModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetCollectModule(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetCollectNFT(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getCollectNFT", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetCollectNFT(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetCollectNFT is a free data retrieval call binding the contract method 0x52aaef55.
//
// Solidity: function getCollectNFT(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetCollectNFT(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetCollectNFT(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubCaller) GetCollectNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getCollectNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubSession) GetCollectNFTImpl() (common.Address, error) {
	return _V1LensHub.Contract.GetCollectNFTImpl(&_V1LensHub.CallOpts)
}

// GetCollectNFTImpl is a free data retrieval call binding the contract method 0x77349a5f.
//
// Solidity: function getCollectNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetCollectNFTImpl() (common.Address, error) {
	return _V1LensHub.Contract.GetCollectNFTImpl(&_V1LensHub.CallOpts)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V1LensHub *V1LensHubCaller) GetContentURI(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getContentURI", profileId, pubId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V1LensHub *V1LensHubSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetContentURI(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V1LensHub *V1LensHubCallerSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetContentURI(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetDispatcher(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getDispatcher", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetDispatcher(&_V1LensHub.CallOpts, profileId)
}

// GetDispatcher is a free data retrieval call binding the contract method 0x540528b9.
//
// Solidity: function getDispatcher(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetDispatcher(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetDispatcher(&_V1LensHub.CallOpts, profileId)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V1LensHub *V1LensHubCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V1LensHub *V1LensHubSession) GetDomainSeparator() ([32]byte, error) {
	return _V1LensHub.Contract.GetDomainSeparator(&_V1LensHub.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V1LensHub *V1LensHubCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _V1LensHub.Contract.GetDomainSeparator(&_V1LensHub.CallOpts)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetFollowModule(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getFollowModule", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetFollowModule(&_V1LensHub.CallOpts, profileId)
}

// GetFollowModule is a free data retrieval call binding the contract method 0xd923d20c.
//
// Solidity: function getFollowModule(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetFollowModule(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetFollowModule(&_V1LensHub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetFollowNFT(opts *bind.CallOpts, profileId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getFollowNFT", profileId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetFollowNFT(&_V1LensHub.CallOpts, profileId)
}

// GetFollowNFT is a free data retrieval call binding the contract method 0xa9ec6563.
//
// Solidity: function getFollowNFT(uint256 profileId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetFollowNFT(profileId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetFollowNFT(&_V1LensHub.CallOpts, profileId)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubCaller) GetFollowNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getFollowNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubSession) GetFollowNFTImpl() (common.Address, error) {
	return _V1LensHub.Contract.GetFollowNFTImpl(&_V1LensHub.CallOpts)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetFollowNFTImpl() (common.Address, error) {
	return _V1LensHub.Contract.GetFollowNFTImpl(&_V1LensHub.CallOpts)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubCaller) GetFollowNFTURI(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getFollowNFTURI", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetFollowNFTURI(&_V1LensHub.CallOpts, profileId)
}

// GetFollowNFTURI is a free data retrieval call binding the contract method 0x374c9473.
//
// Solidity: function getFollowNFTURI(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubCallerSession) GetFollowNFTURI(profileId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetFollowNFTURI(&_V1LensHub.CallOpts, profileId)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V1LensHub *V1LensHubCaller) GetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V1LensHub *V1LensHubSession) GetGovernance() (common.Address, error) {
	return _V1LensHub.Contract.GetGovernance(&_V1LensHub.CallOpts)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetGovernance() (common.Address, error) {
	return _V1LensHub.Contract.GetGovernance(&_V1LensHub.CallOpts)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubCaller) GetHandle(opts *bind.CallOpts, profileId *big.Int) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getHandle", profileId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubSession) GetHandle(profileId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetHandle(&_V1LensHub.CallOpts, profileId)
}

// GetHandle is a free data retrieval call binding the contract method 0xec81d194.
//
// Solidity: function getHandle(uint256 profileId) view returns(string)
func (_V1LensHub *V1LensHubCallerSession) GetHandle(profileId *big.Int) (string, error) {
	return _V1LensHub.Contract.GetHandle(&_V1LensHub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_V1LensHub *V1LensHubCaller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (LensHubProfileStruct, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(LensHubProfileStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(LensHubProfileStruct)).(*LensHubProfileStruct)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_V1LensHub *V1LensHubSession) GetProfile(profileId *big.Int) (LensHubProfileStruct, error) {
	return _V1LensHub.Contract.GetProfile(&_V1LensHub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string))
func (_V1LensHub *V1LensHubCallerSession) GetProfile(profileId *big.Int) (LensHubProfileStruct, error) {
	return _V1LensHub.Contract.GetProfile(&_V1LensHub.CallOpts, profileId)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) GetProfileIdByHandle(opts *bind.CallOpts, handle string) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getProfileIdByHandle", handle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_V1LensHub *V1LensHubSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _V1LensHub.Contract.GetProfileIdByHandle(&_V1LensHub.CallOpts, handle)
}

// GetProfileIdByHandle is a free data retrieval call binding the contract method 0x20fa728a.
//
// Solidity: function getProfileIdByHandle(string handle) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) GetProfileIdByHandle(handle string) (*big.Int, error) {
	return _V1LensHub.Contract.GetProfileIdByHandle(&_V1LensHub.CallOpts, handle)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_V1LensHub *V1LensHubCaller) GetPub(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (LensHubPublicationStruct, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getPub", profileId, pubId)

	if err != nil {
		return *new(LensHubPublicationStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(LensHubPublicationStruct)).(*LensHubPublicationStruct)

	return out0, err

}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_V1LensHub *V1LensHubSession) GetPub(profileId *big.Int, pubId *big.Int) (LensHubPublicationStruct, error) {
	return _V1LensHub.Contract.GetPub(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetPub is a free data retrieval call binding the contract method 0xc736c990.
//
// Solidity: function getPub(uint256 profileId, uint256 pubId) view returns((uint256,uint256,string,address,address,address))
func (_V1LensHub *V1LensHubCallerSession) GetPub(profileId *big.Int, pubId *big.Int) (LensHubPublicationStruct, error) {
	return _V1LensHub.Contract.GetPub(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) GetPubCount(opts *bind.CallOpts, profileId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getPubCount", profileId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_V1LensHub *V1LensHubSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.GetPubCount(&_V1LensHub.CallOpts, profileId)
}

// GetPubCount is a free data retrieval call binding the contract method 0x3a15ff07.
//
// Solidity: function getPubCount(uint256 profileId) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) GetPubCount(profileId *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.GetPubCount(&_V1LensHub.CallOpts, profileId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_V1LensHub *V1LensHubCaller) GetPubPointer(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getPubPointer", profileId, pubId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_V1LensHub *V1LensHubSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _V1LensHub.Contract.GetPubPointer(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetPubPointer is a free data retrieval call binding the contract method 0x5ca3eebf.
//
// Solidity: function getPubPointer(uint256 profileId, uint256 pubId) view returns(uint256, uint256)
func (_V1LensHub *V1LensHubCallerSession) GetPubPointer(profileId *big.Int, pubId *big.Int) (*big.Int, *big.Int, error) {
	return _V1LensHub.Contract.GetPubPointer(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V1LensHub *V1LensHubCaller) GetPubType(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (uint8, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getPubType", profileId, pubId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V1LensHub *V1LensHubSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _V1LensHub.Contract.GetPubType(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetPubType is a free data retrieval call binding the contract method 0x31fff07c.
//
// Solidity: function getPubType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V1LensHub *V1LensHubCallerSession) GetPubType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _V1LensHub.Contract.GetPubType(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCaller) GetReferenceModule(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getReferenceModule", profileId, pubId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetReferenceModule(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetReferenceModule is a free data retrieval call binding the contract method 0xb7984c05.
//
// Solidity: function getReferenceModule(uint256 profileId, uint256 pubId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) GetReferenceModule(profileId *big.Int, pubId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.GetReferenceModule(&_V1LensHub.CallOpts, profileId, pubId)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V1LensHub *V1LensHubCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V1LensHub *V1LensHubSession) GetState() (uint8, error) {
	return _V1LensHub.Contract.GetState(&_V1LensHub.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V1LensHub *V1LensHubCallerSession) GetState() (uint8, error) {
	return _V1LensHub.Contract.GetState(&_V1LensHub.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V1LensHub *V1LensHubCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V1LensHub *V1LensHubSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V1LensHub.Contract.IsApprovedForAll(&_V1LensHub.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V1LensHub.Contract.IsApprovedForAll(&_V1LensHub.CallOpts, owner, operator)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_V1LensHub *V1LensHubCaller) IsCollectModuleWhitelisted(opts *bind.CallOpts, collectModule common.Address) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "isCollectModuleWhitelisted", collectModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_V1LensHub *V1LensHubSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsCollectModuleWhitelisted(&_V1LensHub.CallOpts, collectModule)
}

// IsCollectModuleWhitelisted is a free data retrieval call binding the contract method 0xad3e72bf.
//
// Solidity: function isCollectModuleWhitelisted(address collectModule) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) IsCollectModuleWhitelisted(collectModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsCollectModuleWhitelisted(&_V1LensHub.CallOpts, collectModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_V1LensHub *V1LensHubCaller) IsFollowModuleWhitelisted(opts *bind.CallOpts, followModule common.Address) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "isFollowModuleWhitelisted", followModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_V1LensHub *V1LensHubSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsFollowModuleWhitelisted(&_V1LensHub.CallOpts, followModule)
}

// IsFollowModuleWhitelisted is a free data retrieval call binding the contract method 0x1cbb2620.
//
// Solidity: function isFollowModuleWhitelisted(address followModule) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) IsFollowModuleWhitelisted(followModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsFollowModuleWhitelisted(&_V1LensHub.CallOpts, followModule)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V1LensHub *V1LensHubCaller) IsProfileCreatorWhitelisted(opts *bind.CallOpts, profileCreator common.Address) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "isProfileCreatorWhitelisted", profileCreator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V1LensHub *V1LensHubSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _V1LensHub.Contract.IsProfileCreatorWhitelisted(&_V1LensHub.CallOpts, profileCreator)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _V1LensHub.Contract.IsProfileCreatorWhitelisted(&_V1LensHub.CallOpts, profileCreator)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_V1LensHub *V1LensHubCaller) IsReferenceModuleWhitelisted(opts *bind.CallOpts, referenceModule common.Address) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "isReferenceModuleWhitelisted", referenceModule)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_V1LensHub *V1LensHubSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsReferenceModuleWhitelisted(&_V1LensHub.CallOpts, referenceModule)
}

// IsReferenceModuleWhitelisted is a free data retrieval call binding the contract method 0x8e204fb4.
//
// Solidity: function isReferenceModuleWhitelisted(address referenceModule) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) IsReferenceModuleWhitelisted(referenceModule common.Address) (bool, error) {
	return _V1LensHub.Contract.IsReferenceModuleWhitelisted(&_V1LensHub.CallOpts, referenceModule)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) MintTimestampOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "mintTimestampOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V1LensHub *V1LensHubSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.MintTimestampOf(&_V1LensHub.CallOpts, tokenId)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.MintTimestampOf(&_V1LensHub.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V1LensHub *V1LensHubCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V1LensHub *V1LensHubSession) Name() (string, error) {
	return _V1LensHub.Contract.Name(&_V1LensHub.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V1LensHub *V1LensHubCallerSession) Name() (string, error) {
	return _V1LensHub.Contract.Name(&_V1LensHub.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.OwnerOf(&_V1LensHub.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V1LensHub *V1LensHubCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V1LensHub.Contract.OwnerOf(&_V1LensHub.CallOpts, tokenId)
}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) SigNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "sigNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_V1LensHub *V1LensHubSession) SigNonces(arg0 common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.SigNonces(&_V1LensHub.CallOpts, arg0)
}

// SigNonces is a free data retrieval call binding the contract method 0xf990ccd7.
//
// Solidity: function sigNonces(address ) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) SigNonces(arg0 common.Address) (*big.Int, error) {
	return _V1LensHub.Contract.SigNonces(&_V1LensHub.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V1LensHub *V1LensHubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V1LensHub *V1LensHubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V1LensHub.Contract.SupportsInterface(&_V1LensHub.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V1LensHub *V1LensHubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V1LensHub.Contract.SupportsInterface(&_V1LensHub.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V1LensHub *V1LensHubCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V1LensHub *V1LensHubSession) Symbol() (string, error) {
	return _V1LensHub.Contract.Symbol(&_V1LensHub.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V1LensHub *V1LensHubCallerSession) Symbol() (string, error) {
	return _V1LensHub.Contract.Symbol(&_V1LensHub.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.TokenByIndex(&_V1LensHub.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.TokenByIndex(&_V1LensHub.CallOpts, index)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V1LensHub *V1LensHubCaller) TokenDataOf(opts *bind.CallOpts, tokenId *big.Int) (IERC721TimeTokenData, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "tokenDataOf", tokenId)

	if err != nil {
		return *new(IERC721TimeTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721TimeTokenData)).(*IERC721TimeTokenData)

	return out0, err

}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V1LensHub *V1LensHubSession) TokenDataOf(tokenId *big.Int) (IERC721TimeTokenData, error) {
	return _V1LensHub.Contract.TokenDataOf(&_V1LensHub.CallOpts, tokenId)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V1LensHub *V1LensHubCallerSession) TokenDataOf(tokenId *big.Int) (IERC721TimeTokenData, error) {
	return _V1LensHub.Contract.TokenDataOf(&_V1LensHub.CallOpts, tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.TokenOfOwnerByIndex(&_V1LensHub.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _V1LensHub.Contract.TokenOfOwnerByIndex(&_V1LensHub.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V1LensHub *V1LensHubCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V1LensHub *V1LensHubSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V1LensHub.Contract.TokenURI(&_V1LensHub.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V1LensHub *V1LensHubCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V1LensHub.Contract.TokenURI(&_V1LensHub.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V1LensHub *V1LensHubCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V1LensHub.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V1LensHub *V1LensHubSession) TotalSupply() (*big.Int, error) {
	return _V1LensHub.Contract.TotalSupply(&_V1LensHub.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V1LensHub *V1LensHubCallerSession) TotalSupply() (*big.Int, error) {
	return _V1LensHub.Contract.TotalSupply(&_V1LensHub.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.Approve(&_V1LensHub.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.Approve(&_V1LensHub.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V1LensHub *V1LensHubSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.Burn(&_V1LensHub.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.Burn(&_V1LensHub.TransactOpts, tokenId)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactor) BurnWithSig(opts *bind.TransactOpts, tokenId *big.Int, sig LenHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "burnWithSig", tokenId, sig)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubSession) BurnWithSig(tokenId *big.Int, sig LenHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.BurnWithSig(&_V1LensHub.TransactOpts, tokenId, sig)
}

// BurnWithSig is a paid mutator transaction binding the contract method 0xdd69cdb1.
//
// Solidity: function burnWithSig(uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactorSession) BurnWithSig(tokenId *big.Int, sig LenHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.BurnWithSig(&_V1LensHub.TransactOpts, tokenId, sig)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) Collect(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "collect", profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_V1LensHub *V1LensHubSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.Collect(&_V1LensHub.TransactOpts, profileId, pubId, data)
}

// Collect is a paid mutator transaction binding the contract method 0x84114ad4.
//
// Solidity: function collect(uint256 profileId, uint256 pubId, bytes data) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) Collect(profileId *big.Int, pubId *big.Int, data []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.Collect(&_V1LensHub.TransactOpts, profileId, pubId, data)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) CollectWithSig(opts *bind.TransactOpts, vars LensHubCollectWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "collectWithSig", vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) CollectWithSig(vars LensHubCollectWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CollectWithSig(&_V1LensHub.TransactOpts, vars)
}

// CollectWithSig is a paid mutator transaction binding the contract method 0xb48951e4.
//
// Solidity: function collectWithSig((address,uint256,uint256,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) CollectWithSig(vars LensHubCollectWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CollectWithSig(&_V1LensHub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) Comment(opts *bind.TransactOpts, vars LensHubCommentData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "comment", vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) Comment(vars LensHubCommentData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Comment(&_V1LensHub.TransactOpts, vars)
}

// Comment is a paid mutator transaction binding the contract method 0xb6f32d2b.
//
// Solidity: function comment((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) Comment(vars LensHubCommentData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Comment(&_V1LensHub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) CommentWithSig(opts *bind.TransactOpts, vars LensHubCommentWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "commentWithSig", vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) CommentWithSig(vars LensHubCommentWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CommentWithSig(&_V1LensHub.TransactOpts, vars)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0x7a375716.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) CommentWithSig(vars LensHubCommentWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CommentWithSig(&_V1LensHub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) CreateProfile(opts *bind.TransactOpts, vars LensHubCreateProfileData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "createProfile", vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) CreateProfile(vars LensHubCreateProfileData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CreateProfile(&_V1LensHub.TransactOpts, vars)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xffea138e.
//
// Solidity: function createProfile((address,string,string,address,bytes,string) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) CreateProfile(vars LensHubCreateProfileData) (*types.Transaction, error) {
	return _V1LensHub.Contract.CreateProfile(&_V1LensHub.TransactOpts, vars)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubTransactor) EmitCollectNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "emitCollectNFTTransferEvent", profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.EmitCollectNFTTransferEvent(&_V1LensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubTransactorSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.EmitCollectNFTTransferEvent(&_V1LensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubTransactor) EmitFollowNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "emitFollowNFTTransferEvent", profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.EmitFollowNFTTransferEvent(&_V1LensHub.TransactOpts, profileId, followNFTId, from, to)
}

// EmitFollowNFTTransferEvent is a paid mutator transaction binding the contract method 0xbdfdd4bc.
//
// Solidity: function emitFollowNFTTransferEvent(uint256 profileId, uint256 followNFTId, address from, address to) returns()
func (_V1LensHub *V1LensHubTransactorSession) EmitFollowNFTTransferEvent(profileId *big.Int, followNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.EmitFollowNFTTransferEvent(&_V1LensHub.TransactOpts, profileId, followNFTId, from, to)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_V1LensHub *V1LensHubTransactor) Follow(opts *bind.TransactOpts, profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "follow", profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_V1LensHub *V1LensHubSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.Follow(&_V1LensHub.TransactOpts, profileIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0xfb78ae6c.
//
// Solidity: function follow(uint256[] profileIds, bytes[] datas) returns(uint256[])
func (_V1LensHub *V1LensHubTransactorSession) Follow(profileIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.Follow(&_V1LensHub.TransactOpts, profileIds, datas)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_V1LensHub *V1LensHubTransactor) FollowWithSig(opts *bind.TransactOpts, vars LensHubFollowWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "followWithSig", vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_V1LensHub *V1LensHubSession) FollowWithSig(vars LensHubFollowWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.FollowWithSig(&_V1LensHub.TransactOpts, vars)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x8e4fd6a9.
//
// Solidity: function followWithSig((address,uint256[],bytes[],(uint8,bytes32,bytes32,uint256)) vars) returns(uint256[])
func (_V1LensHub *V1LensHubTransactorSession) FollowWithSig(vars LensHubFollowWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.FollowWithSig(&_V1LensHub.TransactOpts, vars)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V1LensHub *V1LensHubTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "initialize", name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V1LensHub *V1LensHubSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.Initialize(&_V1LensHub.TransactOpts, name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V1LensHub *V1LensHubTransactorSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.Initialize(&_V1LensHub.TransactOpts, name, symbol, newGovernance)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) Mirror(opts *bind.TransactOpts, vars LensHubMirrorData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "mirror", vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) Mirror(vars LensHubMirrorData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Mirror(&_V1LensHub.TransactOpts, vars)
}

// Mirror is a paid mutator transaction binding the contract method 0x2faeed81.
//
// Solidity: function mirror((uint256,uint256,uint256,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) Mirror(vars LensHubMirrorData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Mirror(&_V1LensHub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) MirrorWithSig(opts *bind.TransactOpts, vars LensHubMirrorWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "mirrorWithSig", vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) MirrorWithSig(vars LensHubMirrorWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.MirrorWithSig(&_V1LensHub.TransactOpts, vars)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xdf457c34.
//
// Solidity: function mirrorWithSig((uint256,uint256,uint256,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) MirrorWithSig(vars LensHubMirrorWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.MirrorWithSig(&_V1LensHub.TransactOpts, vars)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactor) Permit(opts *bind.TransactOpts, spender common.Address, tokenId *big.Int, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "permit", spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubSession) Permit(spender common.Address, tokenId *big.Int, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.Permit(&_V1LensHub.TransactOpts, spender, tokenId, sig)
}

// Permit is a paid mutator transaction binding the contract method 0x7ef67f99.
//
// Solidity: function permit(address spender, uint256 tokenId, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactorSession) Permit(spender common.Address, tokenId *big.Int, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.Permit(&_V1LensHub.TransactOpts, spender, tokenId, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactor) PermitForAll(opts *bind.TransactOpts, owner common.Address, operator common.Address, approved bool, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "permitForAll", owner, operator, approved, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubSession) PermitForAll(owner common.Address, operator common.Address, approved bool, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.PermitForAll(&_V1LensHub.TransactOpts, owner, operator, approved, sig)
}

// PermitForAll is a paid mutator transaction binding the contract method 0x89028a13.
//
// Solidity: function permitForAll(address owner, address operator, bool approved, (uint8,bytes32,bytes32,uint256) sig) returns()
func (_V1LensHub *V1LensHubTransactorSession) PermitForAll(owner common.Address, operator common.Address, approved bool, sig LensHubEIP712Signature) (*types.Transaction, error) {
	return _V1LensHub.Contract.PermitForAll(&_V1LensHub.TransactOpts, owner, operator, approved, sig)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) Post(opts *bind.TransactOpts, vars LensHubPostData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "post", vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) Post(vars LensHubPostData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Post(&_V1LensHub.TransactOpts, vars)
}

// Post is a paid mutator transaction binding the contract method 0x963ff141.
//
// Solidity: function post((uint256,string,address,bytes,address,bytes) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) Post(vars LensHubPostData) (*types.Transaction, error) {
	return _V1LensHub.Contract.Post(&_V1LensHub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactor) PostWithSig(opts *bind.TransactOpts, vars LensHubPostWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "postWithSig", vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubSession) PostWithSig(vars LensHubPostWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.PostWithSig(&_V1LensHub.TransactOpts, vars)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x3b508132.
//
// Solidity: function postWithSig((uint256,string,address,bytes,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns(uint256)
func (_V1LensHub *V1LensHubTransactorSession) PostWithSig(vars LensHubPostWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.PostWithSig(&_V1LensHub.TransactOpts, vars)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.SafeTransferFrom(&_V1LensHub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.SafeTransferFrom(&_V1LensHub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V1LensHub *V1LensHubTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V1LensHub *V1LensHubSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.SafeTransferFrom0(&_V1LensHub.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V1LensHub *V1LensHubTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.SafeTransferFrom0(&_V1LensHub.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V1LensHub *V1LensHubTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V1LensHub *V1LensHubSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetApprovalForAll(&_V1LensHub.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetApprovalForAll(&_V1LensHub.TransactOpts, operator, approved)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_V1LensHub *V1LensHubTransactor) SetDefaultProfile(opts *bind.TransactOpts, profileId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setDefaultProfile", profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_V1LensHub *V1LensHubSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDefaultProfile(&_V1LensHub.TransactOpts, profileId)
}

// SetDefaultProfile is a paid mutator transaction binding the contract method 0xf1b2f8bc.
//
// Solidity: function setDefaultProfile(uint256 profileId) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetDefaultProfile(profileId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDefaultProfile(&_V1LensHub.TransactOpts, profileId)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactor) SetDefaultProfileWithSig(opts *bind.TransactOpts, vars LensHubSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setDefaultProfileWithSig", vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubSession) SetDefaultProfileWithSig(vars LensHubSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDefaultProfileWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetDefaultProfileWithSig is a paid mutator transaction binding the contract method 0xdc217253.
//
// Solidity: function setDefaultProfileWithSig((address,uint256,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetDefaultProfileWithSig(vars LensHubSetDefaultProfileWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDefaultProfileWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_V1LensHub *V1LensHubTransactor) SetDispatcher(opts *bind.TransactOpts, profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setDispatcher", profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_V1LensHub *V1LensHubSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDispatcher(&_V1LensHub.TransactOpts, profileId, dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xbfd24f47.
//
// Solidity: function setDispatcher(uint256 profileId, address dispatcher) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetDispatcher(profileId *big.Int, dispatcher common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDispatcher(&_V1LensHub.TransactOpts, profileId, dispatcher)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactor) SetDispatcherWithSig(opts *bind.TransactOpts, vars LensHubSetDispatcherWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setDispatcherWithSig", vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubSession) SetDispatcherWithSig(vars LensHubSetDispatcherWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDispatcherWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetDispatcherWithSig is a paid mutator transaction binding the contract method 0xbfbf0b4b.
//
// Solidity: function setDispatcherWithSig((uint256,address,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetDispatcherWithSig(vars LensHubSetDispatcherWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetDispatcherWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V1LensHub *V1LensHubTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setEmergencyAdmin", newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V1LensHub *V1LensHubSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetEmergencyAdmin(&_V1LensHub.TransactOpts, newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetEmergencyAdmin(&_V1LensHub.TransactOpts, newEmergencyAdmin)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V1LensHub *V1LensHubTransactor) SetFollowModule(opts *bind.TransactOpts, profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setFollowModule", profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V1LensHub *V1LensHubSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowModule(&_V1LensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowModule(&_V1LensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactor) SetFollowModuleWithSig(opts *bind.TransactOpts, vars LensHubSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setFollowModuleWithSig", vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubSession) SetFollowModuleWithSig(vars LensHubSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowModuleWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0x3b28b89f.
//
// Solidity: function setFollowModuleWithSig((uint256,address,bytes,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetFollowModuleWithSig(vars LensHubSetFollowModuleWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowModuleWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_V1LensHub *V1LensHubTransactor) SetFollowNFTURI(opts *bind.TransactOpts, profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setFollowNFTURI", profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_V1LensHub *V1LensHubSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowNFTURI(&_V1LensHub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURI is a paid mutator transaction binding the contract method 0xc6b5d06c.
//
// Solidity: function setFollowNFTURI(uint256 profileId, string followNFTURI) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetFollowNFTURI(profileId *big.Int, followNFTURI string) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowNFTURI(&_V1LensHub.TransactOpts, profileId, followNFTURI)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactor) SetFollowNFTURIWithSig(opts *bind.TransactOpts, vars LensHubSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setFollowNFTURIWithSig", vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubSession) SetFollowNFTURIWithSig(vars LensHubSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowNFTURIWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetFollowNFTURIWithSig is a paid mutator transaction binding the contract method 0xbd12d3f0.
//
// Solidity: function setFollowNFTURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetFollowNFTURIWithSig(vars LensHubSetFollowNFTURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetFollowNFTURIWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V1LensHub *V1LensHubTransactor) SetGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setGovernance", newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V1LensHub *V1LensHubSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetGovernance(&_V1LensHub.TransactOpts, newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetGovernance(&_V1LensHub.TransactOpts, newGovernance)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_V1LensHub *V1LensHubTransactor) SetProfileImageURI(opts *bind.TransactOpts, profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setProfileImageURI", profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_V1LensHub *V1LensHubSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetProfileImageURI(&_V1LensHub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURI is a paid mutator transaction binding the contract method 0x054871b8.
//
// Solidity: function setProfileImageURI(uint256 profileId, string imageURI) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetProfileImageURI(profileId *big.Int, imageURI string) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetProfileImageURI(&_V1LensHub.TransactOpts, profileId, imageURI)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactor) SetProfileImageURIWithSig(opts *bind.TransactOpts, vars LensHubSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setProfileImageURIWithSig", vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubSession) SetProfileImageURIWithSig(vars LensHubSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetProfileImageURIWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetProfileImageURIWithSig is a paid mutator transaction binding the contract method 0x9b84a14c.
//
// Solidity: function setProfileImageURIWithSig((uint256,string,(uint8,bytes32,bytes32,uint256)) vars) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetProfileImageURIWithSig(vars LensHubSetProfileImageURIWithSigData) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetProfileImageURIWithSig(&_V1LensHub.TransactOpts, vars)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V1LensHub *V1LensHubTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V1LensHub *V1LensHubSession) SetState(newState uint8) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetState(&_V1LensHub.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V1LensHub *V1LensHubTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _V1LensHub.Contract.SetState(&_V1LensHub.TransactOpts, newState)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.TransferFrom(&_V1LensHub.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V1LensHub *V1LensHubTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V1LensHub.Contract.TransferFrom(&_V1LensHub.TransactOpts, from, to, tokenId)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactor) WhitelistCollectModule(opts *bind.TransactOpts, collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "whitelistCollectModule", collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistCollectModule(&_V1LensHub.TransactOpts, collectModule, whitelist)
}

// WhitelistCollectModule is a paid mutator transaction binding the contract method 0x31dcebe3.
//
// Solidity: function whitelistCollectModule(address collectModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactorSession) WhitelistCollectModule(collectModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistCollectModule(&_V1LensHub.TransactOpts, collectModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactor) WhitelistFollowModule(opts *bind.TransactOpts, followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "whitelistFollowModule", followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistFollowModule(&_V1LensHub.TransactOpts, followModule, whitelist)
}

// WhitelistFollowModule is a paid mutator transaction binding the contract method 0xa6d8e1e4.
//
// Solidity: function whitelistFollowModule(address followModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactorSession) WhitelistFollowModule(followModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistFollowModule(&_V1LensHub.TransactOpts, followModule, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactor) WhitelistProfileCreator(opts *bind.TransactOpts, profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "whitelistProfileCreator", profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V1LensHub *V1LensHubSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistProfileCreator(&_V1LensHub.TransactOpts, profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactorSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistProfileCreator(&_V1LensHub.TransactOpts, profileCreator, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactor) WhitelistReferenceModule(opts *bind.TransactOpts, referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.contract.Transact(opts, "whitelistReferenceModule", referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistReferenceModule(&_V1LensHub.TransactOpts, referenceModule, whitelist)
}

// WhitelistReferenceModule is a paid mutator transaction binding the contract method 0x4187e4c5.
//
// Solidity: function whitelistReferenceModule(address referenceModule, bool whitelist) returns()
func (_V1LensHub *V1LensHubTransactorSession) WhitelistReferenceModule(referenceModule common.Address, whitelist bool) (*types.Transaction, error) {
	return _V1LensHub.Contract.WhitelistReferenceModule(&_V1LensHub.TransactOpts, referenceModule, whitelist)
}

// V1LensHubApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the V1LensHub contract.
type V1LensHubApprovalIterator struct {
	Event *V1LensHubApproval // Event containing the contract specifics and raw log

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
func (it *V1LensHubApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LensHubApproval)
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
		it.Event = new(V1LensHubApproval)
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
func (it *V1LensHubApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LensHubApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LensHubApproval represents a Approval event raised by the V1LensHub contract.
type V1LensHubApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V1LensHub *V1LensHubFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*V1LensHubApprovalIterator, error) {

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

	logs, sub, err := _V1LensHub.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V1LensHubApprovalIterator{contract: _V1LensHub.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V1LensHub *V1LensHubFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *V1LensHubApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _V1LensHub.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LensHubApproval)
				if err := _V1LensHub.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_V1LensHub *V1LensHubFilterer) ParseApproval(log types.Log) (*V1LensHubApproval, error) {
	event := new(V1LensHubApproval)
	if err := _V1LensHub.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LensHubApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the V1LensHub contract.
type V1LensHubApprovalForAllIterator struct {
	Event *V1LensHubApprovalForAll // Event containing the contract specifics and raw log

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
func (it *V1LensHubApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LensHubApprovalForAll)
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
		it.Event = new(V1LensHubApprovalForAll)
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
func (it *V1LensHubApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LensHubApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LensHubApprovalForAll represents a ApprovalForAll event raised by the V1LensHub contract.
type V1LensHubApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V1LensHub *V1LensHubFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*V1LensHubApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V1LensHub.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &V1LensHubApprovalForAllIterator{contract: _V1LensHub.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V1LensHub *V1LensHubFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *V1LensHubApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V1LensHub.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LensHubApprovalForAll)
				if err := _V1LensHub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_V1LensHub *V1LensHubFilterer) ParseApprovalForAll(log types.Log) (*V1LensHubApprovalForAll, error) {
	event := new(V1LensHubApprovalForAll)
	if err := _V1LensHub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V1LensHubTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the V1LensHub contract.
type V1LensHubTransferIterator struct {
	Event *V1LensHubTransfer // Event containing the contract specifics and raw log

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
func (it *V1LensHubTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V1LensHubTransfer)
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
		it.Event = new(V1LensHubTransfer)
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
func (it *V1LensHubTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V1LensHubTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V1LensHubTransfer represents a Transfer event raised by the V1LensHub contract.
type V1LensHubTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V1LensHub *V1LensHubFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*V1LensHubTransferIterator, error) {

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

	logs, sub, err := _V1LensHub.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V1LensHubTransferIterator{contract: _V1LensHub.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V1LensHub *V1LensHubFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *V1LensHubTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _V1LensHub.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V1LensHubTransfer)
				if err := _V1LensHub.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_V1LensHub *V1LensHubFilterer) ParseTransfer(log types.Log) (*V1LensHubTransfer, error) {
	event := new(V1LensHubTransfer)
	if err := _V1LensHub.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
