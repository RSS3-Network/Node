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

// TypesCreateProfileParams is an auto generated low-level Go binding around an user-defined struct.
type TypesCreateProfileParams struct {
	To                   common.Address
	FollowModule         common.Address
	FollowModuleInitData []byte
}

// TypesEIP712Signature is an auto generated low-level Go binding around an user-defined struct.
type TypesEIP712Signature struct {
	Signer   common.Address
	V        uint8
	R        [32]byte
	S        [32]byte
	Deadline *big.Int
}

// TypesLegacyCollectParams is an auto generated low-level Go binding around an user-defined struct.
type TypesLegacyCollectParams struct {
	PublicationCollectedProfileId *big.Int
	PublicationCollectedId        *big.Int
	CollectorProfileId            *big.Int
	ReferrerProfileId             *big.Int
	ReferrerPubId                 *big.Int
	CollectModuleData             []byte
}

// TypesMigrationParams is an auto generated low-level Go binding around an user-defined struct.
type TypesMigrationParams struct {
	LensHandlesAddress         common.Address
	TokenHandleRegistryAddress common.Address
	LegacyFeeFollowModule      common.Address
	LegacyProfileFollowModule  common.Address
	NewFeeFollowModule         common.Address
}

// TypesProfile is an auto generated low-level Go binding around an user-defined struct.
type TypesProfile struct {
	PubCount               *big.Int
	FollowModule           common.Address
	FollowNFT              common.Address
	DEPRECATEDHandle       string
	DEPRECATEDImageURI     string
	DEPRECATEDFollowNFTURI string
	MetadataURI            string
}

// TypesPublicationMemory is an auto generated low-level Go binding around an user-defined struct.
type TypesPublicationMemory struct {
	PointedProfileId        *big.Int
	PointedPubId            *big.Int
	ContentURI              string
	ReferenceModule         common.Address
	DEPRECATEDCollectModule common.Address
	DEPRECATEDCollectNFT    common.Address
	PubType                 uint8
	RootProfileId           *big.Int
	RootPubId               *big.Int
}

// TypesTokenData is an auto generated low-level Go binding around an user-defined struct.
type TypesTokenData struct {
	Owner         common.Address
	MintTimestamp *big.Int
}

// V2LensHubMetaData contains all meta data concerning the V2LensHub contract.
var V2LensHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"followNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collectNFTImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"moduleRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenGuardianCooldown\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"lensHandlesAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenHandleRegistryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"legacyFeeFollowModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"legacyProfileFollowModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeFollowModule\",\"type\":\"address\"}],\"internalType\":\"structTypes.MigrationParams\",\"name\":\"migrationParams\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotCollectNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotFollowNFT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotInitImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisablingAlreadyTriggered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutorInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GuardianEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InitParamsInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Initialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEOA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotHub\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotMigrationAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotProfileOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublishingPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CollectNFTTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes20\",\"name\":\"gitCommit\",\"type\":\"bytes20\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"LensUpgradeVersion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenGuardianDisablingTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TokenGuardianStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"prevTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasuryFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevTreasury\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TreasurySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"unfollowerProfileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idOfProfileUnfollowed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Unfollowed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DANGER__disableTokenGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_GUARDIAN_COOLDOWN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationActedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationActedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"actionModuleAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"actionModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PublicationActionParams\",\"name\":\"publicationActionParams\",\"type\":\"tuple\"}],\"name\":\"act\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationActedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationActedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"actionModuleAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"actionModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PublicationActionParams\",\"name\":\"publicationActionParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"actWithSig\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"name\":\"batchMigrateFollowModules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"followerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"idOfProfileFollowed\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"followTokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchMigrateFollowers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"followerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfileFollowed\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"followTokenIds\",\"type\":\"uint256[]\"}],\"name\":\"batchMigrateFollows\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"profileIds\",\"type\":\"uint256[]\"}],\"name\":\"batchMigrateProfiles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegatedExecutors\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"approvals\",\"type\":\"bool[]\"},{\"internalType\":\"uint64\",\"name\":\"configNumber\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"switchToGivenConfig\",\"type\":\"bool\"}],\"name\":\"changeDelegatedExecutorsConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegatedExecutors\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"approvals\",\"type\":\"bool[]\"}],\"name\":\"changeDelegatedExecutorsConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"delegatedExecutors\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"approvals\",\"type\":\"bool[]\"},{\"internalType\":\"uint64\",\"name\":\"configNumber\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"switchToGivenConfig\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"changeDelegatedExecutorsConfigWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationCollectedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationCollectedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referrerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referrerPubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.LegacyCollectParams\",\"name\":\"collectParams\",\"type\":\"tuple\"}],\"name\":\"collectLegacy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"publicationCollectedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicationCollectedId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referrerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referrerPubId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"collectModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.LegacyCollectParams\",\"name\":\"collectParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"collectLegacyWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CommentParams\",\"name\":\"commentParams\",\"type\":\"tuple\"}],\"name\":\"comment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CommentParams\",\"name\":\"commentParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"commentWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.CreateProfileParams\",\"name\":\"createProfileParams\",\"type\":\"tuple\"}],\"name\":\"createProfile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectNFTId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"emitCollectNFTTransferEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unfollowerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idOfProfileUnfollowed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"transactionExecutor\",\"type\":\"address\"}],\"name\":\"emitUnfollowedEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emitVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTokenGuardian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"followerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToFollow\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"followTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"follow\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"followerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToFollow\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"followTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"followWithSig\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getContentURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"}],\"name\":\"getDelegatedExecutorsConfigNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"}],\"name\":\"getDelegatedExecutorsMaxConfigNumberSet\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"}],\"name\":\"getDelegatedExecutorsPrevConfigNumber\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFollowNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGitCommit\",\"outputs\":[{\"internalType\":\"bytes20\",\"name\":\"\",\"type\":\"bytes20\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLegacyCollectNFTImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModuleRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"}],\"name\":\"getProfile\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pubCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"followNFT\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"__DEPRECATED__handle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__DEPRECATED__imageURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"__DEPRECATED__followNFTURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"internalType\":\"structTypes.Profile\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"handleHash\",\"type\":\"bytes32\"}],\"name\":\"getProfileIdByHandleHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPublication\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__DEPRECATED__collectModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"__DEPRECATED__collectNFT\",\"type\":\"address\"},{\"internalType\":\"enumTypes.PublicationType\",\"name\":\"pubType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"rootProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootPubId\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.PublicationMemory\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"}],\"name\":\"getPublicationType\",\"outputs\":[{\"internalType\":\"enumTypes.PublicationType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"internalType\":\"enumTypes.ProtocolState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"getTokenGuardianDisablingTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTreasuryFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"increment\",\"type\":\"uint8\"}],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pubId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"}],\"name\":\"isActionModuleEnabledInPublication\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"byProfileId\",\"type\":\"uint256\"}],\"name\":\"isBlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatedExecutor\",\"type\":\"address\"}],\"name\":\"isDelegatedExecutorApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegatorProfileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatedExecutor\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"configNumber\",\"type\":\"uint64\"}],\"name\":\"isDelegatedExecutorApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"followerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"followedProfileId\",\"type\":\"uint256\"}],\"name\":\"isFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"}],\"name\":\"isProfileCreatorWhitelisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mintTimestampOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MirrorParams\",\"name\":\"mirrorParams\",\"type\":\"tuple\"}],\"name\":\"mirror\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.MirrorParams\",\"name\":\"mirrorParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"mirrorWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PostParams\",\"name\":\"postParams\",\"type\":\"tuple\"}],\"name\":\"post\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.PostParams\",\"name\":\"postParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"postWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.QuoteParams\",\"name\":\"quoteParams\",\"type\":\"tuple\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"contentURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pointedProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pointedPubId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerProfileIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"referrerPubIds\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleData\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"actionModules\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"actionModulesInitDatas\",\"type\":\"bytes[]\"},{\"internalType\":\"address\",\"name\":\"referenceModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"referenceModuleInitData\",\"type\":\"bytes\"}],\"internalType\":\"structTypes.QuoteParams\",\"name\":\"quoteParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"quoteWithSig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"byProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToSetBlockStatus\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"blockStatus\",\"type\":\"bool[]\"}],\"name\":\"setBlockStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"byProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToSetBlockStatus\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"blockStatus\",\"type\":\"bool[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"setBlockStatusWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newEmergencyAdmin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"}],\"name\":\"setFollowModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"followModule\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"followModuleInitData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"setFollowModuleWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"migrationAdmins\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"setMigrationAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"setProfileMetadataURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"profileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"setProfileMetadataURIWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"royaltiesInBasisPoints\",\"type\":\"uint256\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumTypes.ProtocolState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"setState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newTreasuryFee\",\"type\":\"uint16\"}],\"name\":\"setTreasuryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenDataOf\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"mintTimestamp\",\"type\":\"uint96\"}],\"internalType\":\"structTypes.TokenData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unfollowerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToUnfollow\",\"type\":\"uint256[]\"}],\"name\":\"unfollow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"unfollowerProfileId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"idsOfProfilesToUnfollow\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structTypes.EIP712Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"name\":\"unfollowWithSig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"profileCreator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"whitelist\",\"type\":\"bool\"}],\"name\":\"whitelistProfileCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// V2LensHubABI is the input ABI used to generate the binding from.
// Deprecated: Use V2LensHubMetaData.ABI instead.
var V2LensHubABI = V2LensHubMetaData.ABI

// V2LensHub is an auto generated Go binding around an Ethereum contract.
type V2LensHub struct {
	V2LensHubCaller     // Read-only binding to the contract
	V2LensHubTransactor // Write-only binding to the contract
	V2LensHubFilterer   // Log filterer for contract events
}

// V2LensHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type V2LensHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2LensHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V2LensHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2LensHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V2LensHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V2LensHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V2LensHubSession struct {
	Contract     *V2LensHub        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V2LensHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V2LensHubCallerSession struct {
	Contract *V2LensHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// V2LensHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V2LensHubTransactorSession struct {
	Contract     *V2LensHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// V2LensHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type V2LensHubRaw struct {
	Contract *V2LensHub // Generic contract binding to access the raw methods on
}

// V2LensHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V2LensHubCallerRaw struct {
	Contract *V2LensHubCaller // Generic read-only contract binding to access the raw methods on
}

// V2LensHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V2LensHubTransactorRaw struct {
	Contract *V2LensHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV2LensHub creates a new instance of V2LensHub, bound to a specific deployed contract.
func NewV2LensHub(address common.Address, backend bind.ContractBackend) (*V2LensHub, error) {
	contract, err := bindV2LensHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V2LensHub{V2LensHubCaller: V2LensHubCaller{contract: contract}, V2LensHubTransactor: V2LensHubTransactor{contract: contract}, V2LensHubFilterer: V2LensHubFilterer{contract: contract}}, nil
}

// NewV2LensHubCaller creates a new read-only instance of V2LensHub, bound to a specific deployed contract.
func NewV2LensHubCaller(address common.Address, caller bind.ContractCaller) (*V2LensHubCaller, error) {
	contract, err := bindV2LensHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V2LensHubCaller{contract: contract}, nil
}

// NewV2LensHubTransactor creates a new write-only instance of V2LensHub, bound to a specific deployed contract.
func NewV2LensHubTransactor(address common.Address, transactor bind.ContractTransactor) (*V2LensHubTransactor, error) {
	contract, err := bindV2LensHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V2LensHubTransactor{contract: contract}, nil
}

// NewV2LensHubFilterer creates a new log filterer instance of V2LensHub, bound to a specific deployed contract.
func NewV2LensHubFilterer(address common.Address, filterer bind.ContractFilterer) (*V2LensHubFilterer, error) {
	contract, err := bindV2LensHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V2LensHubFilterer{contract: contract}, nil
}

// bindV2LensHub binds a generic wrapper to an already deployed contract.
func bindV2LensHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := V2LensHubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2LensHub *V2LensHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2LensHub.Contract.V2LensHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2LensHub *V2LensHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2LensHub.Contract.V2LensHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2LensHub *V2LensHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2LensHub.Contract.V2LensHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V2LensHub *V2LensHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V2LensHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V2LensHub *V2LensHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2LensHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V2LensHub *V2LensHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V2LensHub.Contract.contract.Transact(opts, method, params...)
}

// TOKENGUARDIANCOOLDOWN is a free data retrieval call binding the contract method 0xa88fae83.
//
// Solidity: function TOKEN_GUARDIAN_COOLDOWN() view returns(uint256)
func (_V2LensHub *V2LensHubCaller) TOKENGUARDIANCOOLDOWN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "TOKEN_GUARDIAN_COOLDOWN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOKENGUARDIANCOOLDOWN is a free data retrieval call binding the contract method 0xa88fae83.
//
// Solidity: function TOKEN_GUARDIAN_COOLDOWN() view returns(uint256)
func (_V2LensHub *V2LensHubSession) TOKENGUARDIANCOOLDOWN() (*big.Int, error) {
	return _V2LensHub.Contract.TOKENGUARDIANCOOLDOWN(&_V2LensHub.CallOpts)
}

// TOKENGUARDIANCOOLDOWN is a free data retrieval call binding the contract method 0xa88fae83.
//
// Solidity: function TOKEN_GUARDIAN_COOLDOWN() view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) TOKENGUARDIANCOOLDOWN() (*big.Int, error) {
	return _V2LensHub.Contract.TOKENGUARDIANCOOLDOWN(&_V2LensHub.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V2LensHub *V2LensHubCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V2LensHub *V2LensHubSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.BalanceOf(&_V2LensHub.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.BalanceOf(&_V2LensHub.CallOpts, owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V2LensHub *V2LensHubCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V2LensHub *V2LensHubSession) Exists(tokenId *big.Int) (bool, error) {
	return _V2LensHub.Contract.Exists(&_V2LensHub.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _V2LensHub.Contract.Exists(&_V2LensHub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V2LensHub.Contract.GetApproved(&_V2LensHub.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _V2LensHub.Contract.GetApproved(&_V2LensHub.CallOpts, tokenId)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V2LensHub *V2LensHubCaller) GetContentURI(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (string, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getContentURI", profileId, pubId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V2LensHub *V2LensHubSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _V2LensHub.Contract.GetContentURI(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetContentURI is a free data retrieval call binding the contract method 0xb5a31496.
//
// Solidity: function getContentURI(uint256 profileId, uint256 pubId) view returns(string)
func (_V2LensHub *V2LensHubCallerSession) GetContentURI(profileId *big.Int, pubId *big.Int) (string, error) {
	return _V2LensHub.Contract.GetContentURI(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetDelegatedExecutorsConfigNumber is a free data retrieval call binding the contract method 0x86ee73d7.
//
// Solidity: function getDelegatedExecutorsConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCaller) GetDelegatedExecutorsConfigNumber(opts *bind.CallOpts, delegatorProfileId *big.Int) (uint64, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getDelegatedExecutorsConfigNumber", delegatorProfileId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDelegatedExecutorsConfigNumber is a free data retrieval call binding the contract method 0x86ee73d7.
//
// Solidity: function getDelegatedExecutorsConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubSession) GetDelegatedExecutorsConfigNumber(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsConfigNumber(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDelegatedExecutorsConfigNumber is a free data retrieval call binding the contract method 0x86ee73d7.
//
// Solidity: function getDelegatedExecutorsConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCallerSession) GetDelegatedExecutorsConfigNumber(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsConfigNumber(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDelegatedExecutorsMaxConfigNumberSet is a free data retrieval call binding the contract method 0xae56d75a.
//
// Solidity: function getDelegatedExecutorsMaxConfigNumberSet(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCaller) GetDelegatedExecutorsMaxConfigNumberSet(opts *bind.CallOpts, delegatorProfileId *big.Int) (uint64, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getDelegatedExecutorsMaxConfigNumberSet", delegatorProfileId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDelegatedExecutorsMaxConfigNumberSet is a free data retrieval call binding the contract method 0xae56d75a.
//
// Solidity: function getDelegatedExecutorsMaxConfigNumberSet(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubSession) GetDelegatedExecutorsMaxConfigNumberSet(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsMaxConfigNumberSet(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDelegatedExecutorsMaxConfigNumberSet is a free data retrieval call binding the contract method 0xae56d75a.
//
// Solidity: function getDelegatedExecutorsMaxConfigNumberSet(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCallerSession) GetDelegatedExecutorsMaxConfigNumberSet(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsMaxConfigNumberSet(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDelegatedExecutorsPrevConfigNumber is a free data retrieval call binding the contract method 0xd575b539.
//
// Solidity: function getDelegatedExecutorsPrevConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCaller) GetDelegatedExecutorsPrevConfigNumber(opts *bind.CallOpts, delegatorProfileId *big.Int) (uint64, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getDelegatedExecutorsPrevConfigNumber", delegatorProfileId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetDelegatedExecutorsPrevConfigNumber is a free data retrieval call binding the contract method 0xd575b539.
//
// Solidity: function getDelegatedExecutorsPrevConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubSession) GetDelegatedExecutorsPrevConfigNumber(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsPrevConfigNumber(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDelegatedExecutorsPrevConfigNumber is a free data retrieval call binding the contract method 0xd575b539.
//
// Solidity: function getDelegatedExecutorsPrevConfigNumber(uint256 delegatorProfileId) view returns(uint64)
func (_V2LensHub *V2LensHubCallerSession) GetDelegatedExecutorsPrevConfigNumber(delegatorProfileId *big.Int) (uint64, error) {
	return _V2LensHub.Contract.GetDelegatedExecutorsPrevConfigNumber(&_V2LensHub.CallOpts, delegatorProfileId)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V2LensHub *V2LensHubCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V2LensHub *V2LensHubSession) GetDomainSeparator() ([32]byte, error) {
	return _V2LensHub.Contract.GetDomainSeparator(&_V2LensHub.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_V2LensHub *V2LensHubCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _V2LensHub.Contract.GetDomainSeparator(&_V2LensHub.CallOpts)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubCaller) GetFollowNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getFollowNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubSession) GetFollowNFTImpl() (common.Address, error) {
	return _V2LensHub.Contract.GetFollowNFTImpl(&_V2LensHub.CallOpts)
}

// GetFollowNFTImpl is a free data retrieval call binding the contract method 0x3502ac4b.
//
// Solidity: function getFollowNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetFollowNFTImpl() (common.Address, error) {
	return _V2LensHub.Contract.GetFollowNFTImpl(&_V2LensHub.CallOpts)
}

// GetGitCommit is a free data retrieval call binding the contract method 0x2cf03735.
//
// Solidity: function getGitCommit() pure returns(bytes20)
func (_V2LensHub *V2LensHubCaller) GetGitCommit(opts *bind.CallOpts) ([20]byte, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getGitCommit")

	if err != nil {
		return *new([20]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([20]byte)).(*[20]byte)

	return out0, err

}

// GetGitCommit is a free data retrieval call binding the contract method 0x2cf03735.
//
// Solidity: function getGitCommit() pure returns(bytes20)
func (_V2LensHub *V2LensHubSession) GetGitCommit() ([20]byte, error) {
	return _V2LensHub.Contract.GetGitCommit(&_V2LensHub.CallOpts)
}

// GetGitCommit is a free data retrieval call binding the contract method 0x2cf03735.
//
// Solidity: function getGitCommit() pure returns(bytes20)
func (_V2LensHub *V2LensHubCallerSession) GetGitCommit() ([20]byte, error) {
	return _V2LensHub.Contract.GetGitCommit(&_V2LensHub.CallOpts)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V2LensHub *V2LensHubCaller) GetGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V2LensHub *V2LensHubSession) GetGovernance() (common.Address, error) {
	return _V2LensHub.Contract.GetGovernance(&_V2LensHub.CallOpts)
}

// GetGovernance is a free data retrieval call binding the contract method 0x289b3c0d.
//
// Solidity: function getGovernance() view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetGovernance() (common.Address, error) {
	return _V2LensHub.Contract.GetGovernance(&_V2LensHub.CallOpts)
}

// GetLegacyCollectNFTImpl is a free data retrieval call binding the contract method 0x8f5c291c.
//
// Solidity: function getLegacyCollectNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubCaller) GetLegacyCollectNFTImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getLegacyCollectNFTImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLegacyCollectNFTImpl is a free data retrieval call binding the contract method 0x8f5c291c.
//
// Solidity: function getLegacyCollectNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubSession) GetLegacyCollectNFTImpl() (common.Address, error) {
	return _V2LensHub.Contract.GetLegacyCollectNFTImpl(&_V2LensHub.CallOpts)
}

// GetLegacyCollectNFTImpl is a free data retrieval call binding the contract method 0x8f5c291c.
//
// Solidity: function getLegacyCollectNFTImpl() view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetLegacyCollectNFTImpl() (common.Address, error) {
	return _V2LensHub.Contract.GetLegacyCollectNFTImpl(&_V2LensHub.CallOpts)
}

// GetModuleRegistry is a free data retrieval call binding the contract method 0xedec7952.
//
// Solidity: function getModuleRegistry() view returns(address)
func (_V2LensHub *V2LensHubCaller) GetModuleRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getModuleRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetModuleRegistry is a free data retrieval call binding the contract method 0xedec7952.
//
// Solidity: function getModuleRegistry() view returns(address)
func (_V2LensHub *V2LensHubSession) GetModuleRegistry() (common.Address, error) {
	return _V2LensHub.Contract.GetModuleRegistry(&_V2LensHub.CallOpts)
}

// GetModuleRegistry is a free data retrieval call binding the contract method 0xedec7952.
//
// Solidity: function getModuleRegistry() view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetModuleRegistry() (common.Address, error) {
	return _V2LensHub.Contract.GetModuleRegistry(&_V2LensHub.CallOpts)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string,string))
func (_V2LensHub *V2LensHubCaller) GetProfile(opts *bind.CallOpts, profileId *big.Int) (TypesProfile, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getProfile", profileId)

	if err != nil {
		return *new(TypesProfile), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesProfile)).(*TypesProfile)

	return out0, err

}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string,string))
func (_V2LensHub *V2LensHubSession) GetProfile(profileId *big.Int) (TypesProfile, error) {
	return _V2LensHub.Contract.GetProfile(&_V2LensHub.CallOpts, profileId)
}

// GetProfile is a free data retrieval call binding the contract method 0xf08f4f64.
//
// Solidity: function getProfile(uint256 profileId) view returns((uint256,address,address,string,string,string,string))
func (_V2LensHub *V2LensHubCallerSession) GetProfile(profileId *big.Int) (TypesProfile, error) {
	return _V2LensHub.Contract.GetProfile(&_V2LensHub.CallOpts, profileId)
}

// GetProfileIdByHandleHash is a free data retrieval call binding the contract method 0x19e14070.
//
// Solidity: function getProfileIdByHandleHash(bytes32 handleHash) view returns(uint256)
func (_V2LensHub *V2LensHubCaller) GetProfileIdByHandleHash(opts *bind.CallOpts, handleHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getProfileIdByHandleHash", handleHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProfileIdByHandleHash is a free data retrieval call binding the contract method 0x19e14070.
//
// Solidity: function getProfileIdByHandleHash(bytes32 handleHash) view returns(uint256)
func (_V2LensHub *V2LensHubSession) GetProfileIdByHandleHash(handleHash [32]byte) (*big.Int, error) {
	return _V2LensHub.Contract.GetProfileIdByHandleHash(&_V2LensHub.CallOpts, handleHash)
}

// GetProfileIdByHandleHash is a free data retrieval call binding the contract method 0x19e14070.
//
// Solidity: function getProfileIdByHandleHash(bytes32 handleHash) view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) GetProfileIdByHandleHash(handleHash [32]byte) (*big.Int, error) {
	return _V2LensHub.Contract.GetProfileIdByHandleHash(&_V2LensHub.CallOpts, handleHash)
}

// GetPublication is a free data retrieval call binding the contract method 0x7385ebc9.
//
// Solidity: function getPublication(uint256 profileId, uint256 pubId) pure returns((uint256,uint256,string,address,address,address,uint8,uint256,uint256))
func (_V2LensHub *V2LensHubCaller) GetPublication(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (TypesPublicationMemory, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getPublication", profileId, pubId)

	if err != nil {
		return *new(TypesPublicationMemory), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesPublicationMemory)).(*TypesPublicationMemory)

	return out0, err

}

// GetPublication is a free data retrieval call binding the contract method 0x7385ebc9.
//
// Solidity: function getPublication(uint256 profileId, uint256 pubId) pure returns((uint256,uint256,string,address,address,address,uint8,uint256,uint256))
func (_V2LensHub *V2LensHubSession) GetPublication(profileId *big.Int, pubId *big.Int) (TypesPublicationMemory, error) {
	return _V2LensHub.Contract.GetPublication(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetPublication is a free data retrieval call binding the contract method 0x7385ebc9.
//
// Solidity: function getPublication(uint256 profileId, uint256 pubId) pure returns((uint256,uint256,string,address,address,address,uint8,uint256,uint256))
func (_V2LensHub *V2LensHubCallerSession) GetPublication(profileId *big.Int, pubId *big.Int) (TypesPublicationMemory, error) {
	return _V2LensHub.Contract.GetPublication(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetPublicationType is a free data retrieval call binding the contract method 0x08ed395c.
//
// Solidity: function getPublicationType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V2LensHub *V2LensHubCaller) GetPublicationType(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int) (uint8, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getPublicationType", profileId, pubId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPublicationType is a free data retrieval call binding the contract method 0x08ed395c.
//
// Solidity: function getPublicationType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V2LensHub *V2LensHubSession) GetPublicationType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _V2LensHub.Contract.GetPublicationType(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetPublicationType is a free data retrieval call binding the contract method 0x08ed395c.
//
// Solidity: function getPublicationType(uint256 profileId, uint256 pubId) view returns(uint8)
func (_V2LensHub *V2LensHubCallerSession) GetPublicationType(profileId *big.Int, pubId *big.Int) (uint8, error) {
	return _V2LensHub.Contract.GetPublicationType(&_V2LensHub.CallOpts, profileId, pubId)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V2LensHub *V2LensHubCaller) GetState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V2LensHub *V2LensHubSession) GetState() (uint8, error) {
	return _V2LensHub.Contract.GetState(&_V2LensHub.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns(uint8)
func (_V2LensHub *V2LensHubCallerSession) GetState() (uint8, error) {
	return _V2LensHub.Contract.GetState(&_V2LensHub.CallOpts)
}

// GetTokenGuardianDisablingTimestamp is a free data retrieval call binding the contract method 0xf3bc61f1.
//
// Solidity: function getTokenGuardianDisablingTimestamp(address wallet) view returns(uint256)
func (_V2LensHub *V2LensHubCaller) GetTokenGuardianDisablingTimestamp(opts *bind.CallOpts, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getTokenGuardianDisablingTimestamp", wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenGuardianDisablingTimestamp is a free data retrieval call binding the contract method 0xf3bc61f1.
//
// Solidity: function getTokenGuardianDisablingTimestamp(address wallet) view returns(uint256)
func (_V2LensHub *V2LensHubSession) GetTokenGuardianDisablingTimestamp(wallet common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.GetTokenGuardianDisablingTimestamp(&_V2LensHub.CallOpts, wallet)
}

// GetTokenGuardianDisablingTimestamp is a free data retrieval call binding the contract method 0xf3bc61f1.
//
// Solidity: function getTokenGuardianDisablingTimestamp(address wallet) view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) GetTokenGuardianDisablingTimestamp(wallet common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.GetTokenGuardianDisablingTimestamp(&_V2LensHub.CallOpts, wallet)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_V2LensHub *V2LensHubCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_V2LensHub *V2LensHubSession) GetTreasury() (common.Address, error) {
	return _V2LensHub.Contract.GetTreasury(&_V2LensHub.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_V2LensHub *V2LensHubCallerSession) GetTreasury() (common.Address, error) {
	return _V2LensHub.Contract.GetTreasury(&_V2LensHub.CallOpts)
}

// GetTreasuryData is a free data retrieval call binding the contract method 0x98f965d1.
//
// Solidity: function getTreasuryData() view returns(address, uint16)
func (_V2LensHub *V2LensHubCaller) GetTreasuryData(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getTreasuryData")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetTreasuryData is a free data retrieval call binding the contract method 0x98f965d1.
//
// Solidity: function getTreasuryData() view returns(address, uint16)
func (_V2LensHub *V2LensHubSession) GetTreasuryData() (common.Address, uint16, error) {
	return _V2LensHub.Contract.GetTreasuryData(&_V2LensHub.CallOpts)
}

// GetTreasuryData is a free data retrieval call binding the contract method 0x98f965d1.
//
// Solidity: function getTreasuryData() view returns(address, uint16)
func (_V2LensHub *V2LensHubCallerSession) GetTreasuryData() (common.Address, uint16, error) {
	return _V2LensHub.Contract.GetTreasuryData(&_V2LensHub.CallOpts)
}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint16)
func (_V2LensHub *V2LensHubCaller) GetTreasuryFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getTreasuryFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint16)
func (_V2LensHub *V2LensHubSession) GetTreasuryFee() (uint16, error) {
	return _V2LensHub.Contract.GetTreasuryFee(&_V2LensHub.CallOpts)
}

// GetTreasuryFee is a free data retrieval call binding the contract method 0x29070c6d.
//
// Solidity: function getTreasuryFee() view returns(uint16)
func (_V2LensHub *V2LensHubCallerSession) GetTreasuryFee() (uint16, error) {
	return _V2LensHub.Contract.GetTreasuryFee(&_V2LensHub.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_V2LensHub *V2LensHubCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_V2LensHub *V2LensHubSession) GetVersion() (string, error) {
	return _V2LensHub.Contract.GetVersion(&_V2LensHub.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_V2LensHub *V2LensHubCallerSession) GetVersion() (string, error) {
	return _V2LensHub.Contract.GetVersion(&_V2LensHub.CallOpts)
}

// IsActionModuleEnabledInPublication is a free data retrieval call binding the contract method 0x4bde5ec4.
//
// Solidity: function isActionModuleEnabledInPublication(uint256 profileId, uint256 pubId, address module) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsActionModuleEnabledInPublication(opts *bind.CallOpts, profileId *big.Int, pubId *big.Int, module common.Address) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isActionModuleEnabledInPublication", profileId, pubId, module)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActionModuleEnabledInPublication is a free data retrieval call binding the contract method 0x4bde5ec4.
//
// Solidity: function isActionModuleEnabledInPublication(uint256 profileId, uint256 pubId, address module) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsActionModuleEnabledInPublication(profileId *big.Int, pubId *big.Int, module common.Address) (bool, error) {
	return _V2LensHub.Contract.IsActionModuleEnabledInPublication(&_V2LensHub.CallOpts, profileId, pubId, module)
}

// IsActionModuleEnabledInPublication is a free data retrieval call binding the contract method 0x4bde5ec4.
//
// Solidity: function isActionModuleEnabledInPublication(uint256 profileId, uint256 pubId, address module) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsActionModuleEnabledInPublication(profileId *big.Int, pubId *big.Int, module common.Address) (bool, error) {
	return _V2LensHub.Contract.IsActionModuleEnabledInPublication(&_V2LensHub.CallOpts, profileId, pubId, module)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V2LensHub.Contract.IsApprovedForAll(&_V2LensHub.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _V2LensHub.Contract.IsApprovedForAll(&_V2LensHub.CallOpts, owner, operator)
}

// IsBlocked is a free data retrieval call binding the contract method 0x97e2adf2.
//
// Solidity: function isBlocked(uint256 profileId, uint256 byProfileId) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsBlocked(opts *bind.CallOpts, profileId *big.Int, byProfileId *big.Int) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isBlocked", profileId, byProfileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlocked is a free data retrieval call binding the contract method 0x97e2adf2.
//
// Solidity: function isBlocked(uint256 profileId, uint256 byProfileId) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsBlocked(profileId *big.Int, byProfileId *big.Int) (bool, error) {
	return _V2LensHub.Contract.IsBlocked(&_V2LensHub.CallOpts, profileId, byProfileId)
}

// IsBlocked is a free data retrieval call binding the contract method 0x97e2adf2.
//
// Solidity: function isBlocked(uint256 profileId, uint256 byProfileId) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsBlocked(profileId *big.Int, byProfileId *big.Int) (bool, error) {
	return _V2LensHub.Contract.IsBlocked(&_V2LensHub.CallOpts, profileId, byProfileId)
}

// IsDelegatedExecutorApproved is a free data retrieval call binding the contract method 0x7e341e1e.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsDelegatedExecutorApproved(opts *bind.CallOpts, delegatorProfileId *big.Int, delegatedExecutor common.Address) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isDelegatedExecutorApproved", delegatorProfileId, delegatedExecutor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegatedExecutorApproved is a free data retrieval call binding the contract method 0x7e341e1e.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsDelegatedExecutorApproved(delegatorProfileId *big.Int, delegatedExecutor common.Address) (bool, error) {
	return _V2LensHub.Contract.IsDelegatedExecutorApproved(&_V2LensHub.CallOpts, delegatorProfileId, delegatedExecutor)
}

// IsDelegatedExecutorApproved is a free data retrieval call binding the contract method 0x7e341e1e.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsDelegatedExecutorApproved(delegatorProfileId *big.Int, delegatedExecutor common.Address) (bool, error) {
	return _V2LensHub.Contract.IsDelegatedExecutorApproved(&_V2LensHub.CallOpts, delegatorProfileId, delegatedExecutor)
}

// IsDelegatedExecutorApproved0 is a free data retrieval call binding the contract method 0xe72dfde9.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor, uint64 configNumber) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsDelegatedExecutorApproved0(opts *bind.CallOpts, delegatorProfileId *big.Int, delegatedExecutor common.Address, configNumber uint64) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isDelegatedExecutorApproved0", delegatorProfileId, delegatedExecutor, configNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegatedExecutorApproved0 is a free data retrieval call binding the contract method 0xe72dfde9.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor, uint64 configNumber) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsDelegatedExecutorApproved0(delegatorProfileId *big.Int, delegatedExecutor common.Address, configNumber uint64) (bool, error) {
	return _V2LensHub.Contract.IsDelegatedExecutorApproved0(&_V2LensHub.CallOpts, delegatorProfileId, delegatedExecutor, configNumber)
}

// IsDelegatedExecutorApproved0 is a free data retrieval call binding the contract method 0xe72dfde9.
//
// Solidity: function isDelegatedExecutorApproved(uint256 delegatorProfileId, address delegatedExecutor, uint64 configNumber) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsDelegatedExecutorApproved0(delegatorProfileId *big.Int, delegatedExecutor common.Address, configNumber uint64) (bool, error) {
	return _V2LensHub.Contract.IsDelegatedExecutorApproved0(&_V2LensHub.CallOpts, delegatorProfileId, delegatedExecutor, configNumber)
}

// IsFollowing is a free data retrieval call binding the contract method 0x47720ebb.
//
// Solidity: function isFollowing(uint256 followerProfileId, uint256 followedProfileId) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsFollowing(opts *bind.CallOpts, followerProfileId *big.Int, followedProfileId *big.Int) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isFollowing", followerProfileId, followedProfileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFollowing is a free data retrieval call binding the contract method 0x47720ebb.
//
// Solidity: function isFollowing(uint256 followerProfileId, uint256 followedProfileId) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsFollowing(followerProfileId *big.Int, followedProfileId *big.Int) (bool, error) {
	return _V2LensHub.Contract.IsFollowing(&_V2LensHub.CallOpts, followerProfileId, followedProfileId)
}

// IsFollowing is a free data retrieval call binding the contract method 0x47720ebb.
//
// Solidity: function isFollowing(uint256 followerProfileId, uint256 followedProfileId) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsFollowing(followerProfileId *big.Int, followedProfileId *big.Int) (bool, error) {
	return _V2LensHub.Contract.IsFollowing(&_V2LensHub.CallOpts, followerProfileId, followedProfileId)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V2LensHub *V2LensHubCaller) IsProfileCreatorWhitelisted(opts *bind.CallOpts, profileCreator common.Address) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "isProfileCreatorWhitelisted", profileCreator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V2LensHub *V2LensHubSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _V2LensHub.Contract.IsProfileCreatorWhitelisted(&_V2LensHub.CallOpts, profileCreator)
}

// IsProfileCreatorWhitelisted is a free data retrieval call binding the contract method 0xaf05dd22.
//
// Solidity: function isProfileCreatorWhitelisted(address profileCreator) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) IsProfileCreatorWhitelisted(profileCreator common.Address) (bool, error) {
	return _V2LensHub.Contract.IsProfileCreatorWhitelisted(&_V2LensHub.CallOpts, profileCreator)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V2LensHub *V2LensHubCaller) MintTimestampOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "mintTimestampOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V2LensHub *V2LensHubSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _V2LensHub.Contract.MintTimestampOf(&_V2LensHub.CallOpts, tokenId)
}

// MintTimestampOf is a free data retrieval call binding the contract method 0x50ddf35c.
//
// Solidity: function mintTimestampOf(uint256 tokenId) view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) MintTimestampOf(tokenId *big.Int) (*big.Int, error) {
	return _V2LensHub.Contract.MintTimestampOf(&_V2LensHub.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V2LensHub *V2LensHubCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V2LensHub *V2LensHubSession) Name() (string, error) {
	return _V2LensHub.Contract.Name(&_V2LensHub.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_V2LensHub *V2LensHubCallerSession) Name() (string, error) {
	return _V2LensHub.Contract.Name(&_V2LensHub.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256)
func (_V2LensHub *V2LensHubCaller) Nonces(opts *bind.CallOpts, signer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "nonces", signer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256)
func (_V2LensHub *V2LensHubSession) Nonces(signer common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.Nonces(&_V2LensHub.CallOpts, signer)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address signer) view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) Nonces(signer common.Address) (*big.Int, error) {
	return _V2LensHub.Contract.Nonces(&_V2LensHub.CallOpts, signer)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V2LensHub.Contract.OwnerOf(&_V2LensHub.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_V2LensHub *V2LensHubCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _V2LensHub.Contract.OwnerOf(&_V2LensHub.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_V2LensHub *V2LensHubCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_V2LensHub *V2LensHubSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _V2LensHub.Contract.RoyaltyInfo(&_V2LensHub.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address, uint256)
func (_V2LensHub *V2LensHubCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (common.Address, *big.Int, error) {
	return _V2LensHub.Contract.RoyaltyInfo(&_V2LensHub.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V2LensHub *V2LensHubCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V2LensHub *V2LensHubSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V2LensHub.Contract.SupportsInterface(&_V2LensHub.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_V2LensHub *V2LensHubCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _V2LensHub.Contract.SupportsInterface(&_V2LensHub.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V2LensHub *V2LensHubCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V2LensHub *V2LensHubSession) Symbol() (string, error) {
	return _V2LensHub.Contract.Symbol(&_V2LensHub.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_V2LensHub *V2LensHubCallerSession) Symbol() (string, error) {
	return _V2LensHub.Contract.Symbol(&_V2LensHub.CallOpts)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V2LensHub *V2LensHubCaller) TokenDataOf(opts *bind.CallOpts, tokenId *big.Int) (TypesTokenData, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "tokenDataOf", tokenId)

	if err != nil {
		return *new(TypesTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new(TypesTokenData)).(*TypesTokenData)

	return out0, err

}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V2LensHub *V2LensHubSession) TokenDataOf(tokenId *big.Int) (TypesTokenData, error) {
	return _V2LensHub.Contract.TokenDataOf(&_V2LensHub.CallOpts, tokenId)
}

// TokenDataOf is a free data retrieval call binding the contract method 0xc0da9bcd.
//
// Solidity: function tokenDataOf(uint256 tokenId) view returns((address,uint96))
func (_V2LensHub *V2LensHubCallerSession) TokenDataOf(tokenId *big.Int) (TypesTokenData, error) {
	return _V2LensHub.Contract.TokenDataOf(&_V2LensHub.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V2LensHub *V2LensHubCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V2LensHub *V2LensHubSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V2LensHub.Contract.TokenURI(&_V2LensHub.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_V2LensHub *V2LensHubCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _V2LensHub.Contract.TokenURI(&_V2LensHub.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V2LensHub *V2LensHubCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _V2LensHub.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V2LensHub *V2LensHubSession) TotalSupply() (*big.Int, error) {
	return _V2LensHub.Contract.TotalSupply(&_V2LensHub.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_V2LensHub *V2LensHubCallerSession) TotalSupply() (*big.Int, error) {
	return _V2LensHub.Contract.TotalSupply(&_V2LensHub.CallOpts)
}

// DANGERDisableTokenGuardian is a paid mutator transaction binding the contract method 0x2248f76d.
//
// Solidity: function DANGER__disableTokenGuardian() returns()
func (_V2LensHub *V2LensHubTransactor) DANGERDisableTokenGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "DANGER__disableTokenGuardian")
}

// DANGERDisableTokenGuardian is a paid mutator transaction binding the contract method 0x2248f76d.
//
// Solidity: function DANGER__disableTokenGuardian() returns()
func (_V2LensHub *V2LensHubSession) DANGERDisableTokenGuardian() (*types.Transaction, error) {
	return _V2LensHub.Contract.DANGERDisableTokenGuardian(&_V2LensHub.TransactOpts)
}

// DANGERDisableTokenGuardian is a paid mutator transaction binding the contract method 0x2248f76d.
//
// Solidity: function DANGER__disableTokenGuardian() returns()
func (_V2LensHub *V2LensHubTransactorSession) DANGERDisableTokenGuardian() (*types.Transaction, error) {
	return _V2LensHub.Contract.DANGERDisableTokenGuardian(&_V2LensHub.TransactOpts)
}

// Act is a paid mutator transaction binding the contract method 0xc5d5d96a.
//
// Solidity: function act((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams) returns(bytes)
func (_V2LensHub *V2LensHubTransactor) Act(opts *bind.TransactOpts, publicationActionParams TypesPublicationActionParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "act", publicationActionParams)
}

// Act is a paid mutator transaction binding the contract method 0xc5d5d96a.
//
// Solidity: function act((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams) returns(bytes)
func (_V2LensHub *V2LensHubSession) Act(publicationActionParams TypesPublicationActionParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Act(&_V2LensHub.TransactOpts, publicationActionParams)
}

// Act is a paid mutator transaction binding the contract method 0xc5d5d96a.
//
// Solidity: function act((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams) returns(bytes)
func (_V2LensHub *V2LensHubTransactorSession) Act(publicationActionParams TypesPublicationActionParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Act(&_V2LensHub.TransactOpts, publicationActionParams)
}

// ActWithSig is a paid mutator transaction binding the contract method 0xc86642c7.
//
// Solidity: function actWithSig((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(bytes)
func (_V2LensHub *V2LensHubTransactor) ActWithSig(opts *bind.TransactOpts, publicationActionParams TypesPublicationActionParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "actWithSig", publicationActionParams, signature)
}

// ActWithSig is a paid mutator transaction binding the contract method 0xc86642c7.
//
// Solidity: function actWithSig((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(bytes)
func (_V2LensHub *V2LensHubSession) ActWithSig(publicationActionParams TypesPublicationActionParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.ActWithSig(&_V2LensHub.TransactOpts, publicationActionParams, signature)
}

// ActWithSig is a paid mutator transaction binding the contract method 0xc86642c7.
//
// Solidity: function actWithSig((uint256,uint256,uint256,uint256[],uint256[],address,bytes) publicationActionParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(bytes)
func (_V2LensHub *V2LensHubTransactorSession) ActWithSig(publicationActionParams TypesPublicationActionParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.ActWithSig(&_V2LensHub.TransactOpts, publicationActionParams, signature)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Approve(&_V2LensHub.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Approve(&_V2LensHub.TransactOpts, to, tokenId)
}

// BatchMigrateFollowModules is a paid mutator transaction binding the contract method 0x8c28cc1d.
//
// Solidity: function batchMigrateFollowModules(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubTransactor) BatchMigrateFollowModules(opts *bind.TransactOpts, profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "batchMigrateFollowModules", profileIds)
}

// BatchMigrateFollowModules is a paid mutator transaction binding the contract method 0x8c28cc1d.
//
// Solidity: function batchMigrateFollowModules(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubSession) BatchMigrateFollowModules(profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollowModules(&_V2LensHub.TransactOpts, profileIds)
}

// BatchMigrateFollowModules is a paid mutator transaction binding the contract method 0x8c28cc1d.
//
// Solidity: function batchMigrateFollowModules(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubTransactorSession) BatchMigrateFollowModules(profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollowModules(&_V2LensHub.TransactOpts, profileIds)
}

// BatchMigrateFollowers is a paid mutator transaction binding the contract method 0xae900c00.
//
// Solidity: function batchMigrateFollowers(uint256[] followerProfileIds, uint256 idOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubTransactor) BatchMigrateFollowers(opts *bind.TransactOpts, followerProfileIds []*big.Int, idOfProfileFollowed *big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "batchMigrateFollowers", followerProfileIds, idOfProfileFollowed, followTokenIds)
}

// BatchMigrateFollowers is a paid mutator transaction binding the contract method 0xae900c00.
//
// Solidity: function batchMigrateFollowers(uint256[] followerProfileIds, uint256 idOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubSession) BatchMigrateFollowers(followerProfileIds []*big.Int, idOfProfileFollowed *big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollowers(&_V2LensHub.TransactOpts, followerProfileIds, idOfProfileFollowed, followTokenIds)
}

// BatchMigrateFollowers is a paid mutator transaction binding the contract method 0xae900c00.
//
// Solidity: function batchMigrateFollowers(uint256[] followerProfileIds, uint256 idOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubTransactorSession) BatchMigrateFollowers(followerProfileIds []*big.Int, idOfProfileFollowed *big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollowers(&_V2LensHub.TransactOpts, followerProfileIds, idOfProfileFollowed, followTokenIds)
}

// BatchMigrateFollows is a paid mutator transaction binding the contract method 0xd9efd1d1.
//
// Solidity: function batchMigrateFollows(uint256 followerProfileId, uint256[] idsOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubTransactor) BatchMigrateFollows(opts *bind.TransactOpts, followerProfileId *big.Int, idsOfProfileFollowed []*big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "batchMigrateFollows", followerProfileId, idsOfProfileFollowed, followTokenIds)
}

// BatchMigrateFollows is a paid mutator transaction binding the contract method 0xd9efd1d1.
//
// Solidity: function batchMigrateFollows(uint256 followerProfileId, uint256[] idsOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubSession) BatchMigrateFollows(followerProfileId *big.Int, idsOfProfileFollowed []*big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollows(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfileFollowed, followTokenIds)
}

// BatchMigrateFollows is a paid mutator transaction binding the contract method 0xd9efd1d1.
//
// Solidity: function batchMigrateFollows(uint256 followerProfileId, uint256[] idsOfProfileFollowed, uint256[] followTokenIds) returns()
func (_V2LensHub *V2LensHubTransactorSession) BatchMigrateFollows(followerProfileId *big.Int, idsOfProfileFollowed []*big.Int, followTokenIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateFollows(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfileFollowed, followTokenIds)
}

// BatchMigrateProfiles is a paid mutator transaction binding the contract method 0xa7274274.
//
// Solidity: function batchMigrateProfiles(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubTransactor) BatchMigrateProfiles(opts *bind.TransactOpts, profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "batchMigrateProfiles", profileIds)
}

// BatchMigrateProfiles is a paid mutator transaction binding the contract method 0xa7274274.
//
// Solidity: function batchMigrateProfiles(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubSession) BatchMigrateProfiles(profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateProfiles(&_V2LensHub.TransactOpts, profileIds)
}

// BatchMigrateProfiles is a paid mutator transaction binding the contract method 0xa7274274.
//
// Solidity: function batchMigrateProfiles(uint256[] profileIds) returns()
func (_V2LensHub *V2LensHubTransactorSession) BatchMigrateProfiles(profileIds []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.BatchMigrateProfiles(&_V2LensHub.TransactOpts, profileIds)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V2LensHub *V2LensHubSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Burn(&_V2LensHub.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Burn(&_V2LensHub.TransactOpts, tokenId)
}

// ChangeDelegatedExecutorsConfig is a paid mutator transaction binding the contract method 0x51c3659c.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig) returns()
func (_V2LensHub *V2LensHubTransactor) ChangeDelegatedExecutorsConfig(opts *bind.TransactOpts, delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "changeDelegatedExecutorsConfig", delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig)
}

// ChangeDelegatedExecutorsConfig is a paid mutator transaction binding the contract method 0x51c3659c.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig) returns()
func (_V2LensHub *V2LensHubSession) ChangeDelegatedExecutorsConfig(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfig(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig)
}

// ChangeDelegatedExecutorsConfig is a paid mutator transaction binding the contract method 0x51c3659c.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig) returns()
func (_V2LensHub *V2LensHubTransactorSession) ChangeDelegatedExecutorsConfig(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfig(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig)
}

// ChangeDelegatedExecutorsConfig0 is a paid mutator transaction binding the contract method 0xc1f4b40a.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals) returns()
func (_V2LensHub *V2LensHubTransactor) ChangeDelegatedExecutorsConfig0(opts *bind.TransactOpts, delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "changeDelegatedExecutorsConfig0", delegatorProfileId, delegatedExecutors, approvals)
}

// ChangeDelegatedExecutorsConfig0 is a paid mutator transaction binding the contract method 0xc1f4b40a.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals) returns()
func (_V2LensHub *V2LensHubSession) ChangeDelegatedExecutorsConfig0(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfig0(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals)
}

// ChangeDelegatedExecutorsConfig0 is a paid mutator transaction binding the contract method 0xc1f4b40a.
//
// Solidity: function changeDelegatedExecutorsConfig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals) returns()
func (_V2LensHub *V2LensHubTransactorSession) ChangeDelegatedExecutorsConfig0(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfig0(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals)
}

// ChangeDelegatedExecutorsConfigWithSig is a paid mutator transaction binding the contract method 0x4926c4ed.
//
// Solidity: function changeDelegatedExecutorsConfigWithSig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactor) ChangeDelegatedExecutorsConfigWithSig(opts *bind.TransactOpts, delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "changeDelegatedExecutorsConfigWithSig", delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig, signature)
}

// ChangeDelegatedExecutorsConfigWithSig is a paid mutator transaction binding the contract method 0x4926c4ed.
//
// Solidity: function changeDelegatedExecutorsConfigWithSig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubSession) ChangeDelegatedExecutorsConfigWithSig(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfigWithSig(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig, signature)
}

// ChangeDelegatedExecutorsConfigWithSig is a paid mutator transaction binding the contract method 0x4926c4ed.
//
// Solidity: function changeDelegatedExecutorsConfigWithSig(uint256 delegatorProfileId, address[] delegatedExecutors, bool[] approvals, uint64 configNumber, bool switchToGivenConfig, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactorSession) ChangeDelegatedExecutorsConfigWithSig(delegatorProfileId *big.Int, delegatedExecutors []common.Address, approvals []bool, configNumber uint64, switchToGivenConfig bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.ChangeDelegatedExecutorsConfigWithSig(&_V2LensHub.TransactOpts, delegatorProfileId, delegatedExecutors, approvals, configNumber, switchToGivenConfig, signature)
}

// CollectLegacy is a paid mutator transaction binding the contract method 0x4727ce3a.
//
// Solidity: function collectLegacy((uint256,uint256,uint256,uint256,uint256,bytes) collectParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) CollectLegacy(opts *bind.TransactOpts, collectParams TypesLegacyCollectParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "collectLegacy", collectParams)
}

// CollectLegacy is a paid mutator transaction binding the contract method 0x4727ce3a.
//
// Solidity: function collectLegacy((uint256,uint256,uint256,uint256,uint256,bytes) collectParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) CollectLegacy(collectParams TypesLegacyCollectParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.CollectLegacy(&_V2LensHub.TransactOpts, collectParams)
}

// CollectLegacy is a paid mutator transaction binding the contract method 0x4727ce3a.
//
// Solidity: function collectLegacy((uint256,uint256,uint256,uint256,uint256,bytes) collectParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) CollectLegacy(collectParams TypesLegacyCollectParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.CollectLegacy(&_V2LensHub.TransactOpts, collectParams)
}

// CollectLegacyWithSig is a paid mutator transaction binding the contract method 0xb7902e73.
//
// Solidity: function collectLegacyWithSig((uint256,uint256,uint256,uint256,uint256,bytes) collectParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) CollectLegacyWithSig(opts *bind.TransactOpts, collectParams TypesLegacyCollectParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "collectLegacyWithSig", collectParams, signature)
}

// CollectLegacyWithSig is a paid mutator transaction binding the contract method 0xb7902e73.
//
// Solidity: function collectLegacyWithSig((uint256,uint256,uint256,uint256,uint256,bytes) collectParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubSession) CollectLegacyWithSig(collectParams TypesLegacyCollectParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.CollectLegacyWithSig(&_V2LensHub.TransactOpts, collectParams, signature)
}

// CollectLegacyWithSig is a paid mutator transaction binding the contract method 0xb7902e73.
//
// Solidity: function collectLegacyWithSig((uint256,uint256,uint256,uint256,uint256,bytes) collectParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) CollectLegacyWithSig(collectParams TypesLegacyCollectParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.CollectLegacyWithSig(&_V2LensHub.TransactOpts, collectParams, signature)
}

// Comment is a paid mutator transaction binding the contract method 0xb273b4a7.
//
// Solidity: function comment((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) Comment(opts *bind.TransactOpts, commentParams TypesCommentParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "comment", commentParams)
}

// Comment is a paid mutator transaction binding the contract method 0xb273b4a7.
//
// Solidity: function comment((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) Comment(commentParams TypesCommentParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Comment(&_V2LensHub.TransactOpts, commentParams)
}

// Comment is a paid mutator transaction binding the contract method 0xb273b4a7.
//
// Solidity: function comment((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) Comment(commentParams TypesCommentParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Comment(&_V2LensHub.TransactOpts, commentParams)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0xb42df51a.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) CommentWithSig(opts *bind.TransactOpts, commentParams TypesCommentParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "commentWithSig", commentParams, signature)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0xb42df51a.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubSession) CommentWithSig(commentParams TypesCommentParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.CommentWithSig(&_V2LensHub.TransactOpts, commentParams, signature)
}

// CommentWithSig is a paid mutator transaction binding the contract method 0xb42df51a.
//
// Solidity: function commentWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) commentParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) CommentWithSig(commentParams TypesCommentParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.CommentWithSig(&_V2LensHub.TransactOpts, commentParams, signature)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x560a4db1.
//
// Solidity: function createProfile((address,address,bytes) createProfileParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) CreateProfile(opts *bind.TransactOpts, createProfileParams TypesCreateProfileParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "createProfile", createProfileParams)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x560a4db1.
//
// Solidity: function createProfile((address,address,bytes) createProfileParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) CreateProfile(createProfileParams TypesCreateProfileParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.CreateProfile(&_V2LensHub.TransactOpts, createProfileParams)
}

// CreateProfile is a paid mutator transaction binding the contract method 0x560a4db1.
//
// Solidity: function createProfile((address,address,bytes) createProfileParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) CreateProfile(createProfileParams TypesCreateProfileParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.CreateProfile(&_V2LensHub.TransactOpts, createProfileParams)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V2LensHub *V2LensHubTransactor) EmitCollectNFTTransferEvent(opts *bind.TransactOpts, profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "emitCollectNFTTransferEvent", profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V2LensHub *V2LensHubSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitCollectNFTTransferEvent(&_V2LensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitCollectNFTTransferEvent is a paid mutator transaction binding the contract method 0x86e2947b.
//
// Solidity: function emitCollectNFTTransferEvent(uint256 profileId, uint256 pubId, uint256 collectNFTId, address from, address to) returns()
func (_V2LensHub *V2LensHubTransactorSession) EmitCollectNFTTransferEvent(profileId *big.Int, pubId *big.Int, collectNFTId *big.Int, from common.Address, to common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitCollectNFTTransferEvent(&_V2LensHub.TransactOpts, profileId, pubId, collectNFTId, from, to)
}

// EmitUnfollowedEvent is a paid mutator transaction binding the contract method 0x28ab6fb4.
//
// Solidity: function emitUnfollowedEvent(uint256 unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor) returns()
func (_V2LensHub *V2LensHubTransactor) EmitUnfollowedEvent(opts *bind.TransactOpts, unfollowerProfileId *big.Int, idOfProfileUnfollowed *big.Int, transactionExecutor common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "emitUnfollowedEvent", unfollowerProfileId, idOfProfileUnfollowed, transactionExecutor)
}

// EmitUnfollowedEvent is a paid mutator transaction binding the contract method 0x28ab6fb4.
//
// Solidity: function emitUnfollowedEvent(uint256 unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor) returns()
func (_V2LensHub *V2LensHubSession) EmitUnfollowedEvent(unfollowerProfileId *big.Int, idOfProfileUnfollowed *big.Int, transactionExecutor common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitUnfollowedEvent(&_V2LensHub.TransactOpts, unfollowerProfileId, idOfProfileUnfollowed, transactionExecutor)
}

// EmitUnfollowedEvent is a paid mutator transaction binding the contract method 0x28ab6fb4.
//
// Solidity: function emitUnfollowedEvent(uint256 unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor) returns()
func (_V2LensHub *V2LensHubTransactorSession) EmitUnfollowedEvent(unfollowerProfileId *big.Int, idOfProfileUnfollowed *big.Int, transactionExecutor common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitUnfollowedEvent(&_V2LensHub.TransactOpts, unfollowerProfileId, idOfProfileUnfollowed, transactionExecutor)
}

// EmitVersion is a paid mutator transaction binding the contract method 0xac11641d.
//
// Solidity: function emitVersion() returns()
func (_V2LensHub *V2LensHubTransactor) EmitVersion(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "emitVersion")
}

// EmitVersion is a paid mutator transaction binding the contract method 0xac11641d.
//
// Solidity: function emitVersion() returns()
func (_V2LensHub *V2LensHubSession) EmitVersion() (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitVersion(&_V2LensHub.TransactOpts)
}

// EmitVersion is a paid mutator transaction binding the contract method 0xac11641d.
//
// Solidity: function emitVersion() returns()
func (_V2LensHub *V2LensHubTransactorSession) EmitVersion() (*types.Transaction, error) {
	return _V2LensHub.Contract.EmitVersion(&_V2LensHub.TransactOpts)
}

// EnableTokenGuardian is a paid mutator transaction binding the contract method 0x1e9df673.
//
// Solidity: function enableTokenGuardian() returns()
func (_V2LensHub *V2LensHubTransactor) EnableTokenGuardian(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "enableTokenGuardian")
}

// EnableTokenGuardian is a paid mutator transaction binding the contract method 0x1e9df673.
//
// Solidity: function enableTokenGuardian() returns()
func (_V2LensHub *V2LensHubSession) EnableTokenGuardian() (*types.Transaction, error) {
	return _V2LensHub.Contract.EnableTokenGuardian(&_V2LensHub.TransactOpts)
}

// EnableTokenGuardian is a paid mutator transaction binding the contract method 0x1e9df673.
//
// Solidity: function enableTokenGuardian() returns()
func (_V2LensHub *V2LensHubTransactorSession) EnableTokenGuardian() (*types.Transaction, error) {
	return _V2LensHub.Contract.EnableTokenGuardian(&_V2LensHub.TransactOpts)
}

// Follow is a paid mutator transaction binding the contract method 0x4b7312a9.
//
// Solidity: function follow(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas) returns(uint256[])
func (_V2LensHub *V2LensHubTransactor) Follow(opts *bind.TransactOpts, followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "follow", followerProfileId, idsOfProfilesToFollow, followTokenIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0x4b7312a9.
//
// Solidity: function follow(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas) returns(uint256[])
func (_V2LensHub *V2LensHubSession) Follow(followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.Follow(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfilesToFollow, followTokenIds, datas)
}

// Follow is a paid mutator transaction binding the contract method 0x4b7312a9.
//
// Solidity: function follow(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas) returns(uint256[])
func (_V2LensHub *V2LensHubTransactorSession) Follow(followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.Follow(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfilesToFollow, followTokenIds, datas)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x4d7b35a0.
//
// Solidity: function followWithSig(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256[])
func (_V2LensHub *V2LensHubTransactor) FollowWithSig(opts *bind.TransactOpts, followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "followWithSig", followerProfileId, idsOfProfilesToFollow, followTokenIds, datas, signature)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x4d7b35a0.
//
// Solidity: function followWithSig(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256[])
func (_V2LensHub *V2LensHubSession) FollowWithSig(followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.FollowWithSig(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfilesToFollow, followTokenIds, datas, signature)
}

// FollowWithSig is a paid mutator transaction binding the contract method 0x4d7b35a0.
//
// Solidity: function followWithSig(uint256 followerProfileId, uint256[] idsOfProfilesToFollow, uint256[] followTokenIds, bytes[] datas, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256[])
func (_V2LensHub *V2LensHubTransactorSession) FollowWithSig(followerProfileId *big.Int, idsOfProfilesToFollow []*big.Int, followTokenIds []*big.Int, datas [][]byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.FollowWithSig(&_V2LensHub.TransactOpts, followerProfileId, idsOfProfilesToFollow, followTokenIds, datas, signature)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2LensHub *V2LensHubTransactor) IncrementNonce(opts *bind.TransactOpts, increment uint8) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "incrementNonce", increment)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2LensHub *V2LensHubSession) IncrementNonce(increment uint8) (*types.Transaction, error) {
	return _V2LensHub.Contract.IncrementNonce(&_V2LensHub.TransactOpts, increment)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x365ae23a.
//
// Solidity: function incrementNonce(uint8 increment) returns()
func (_V2LensHub *V2LensHubTransactorSession) IncrementNonce(increment uint8) (*types.Transaction, error) {
	return _V2LensHub.Contract.IncrementNonce(&_V2LensHub.TransactOpts, increment)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V2LensHub *V2LensHubTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "initialize", name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V2LensHub *V2LensHubSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.Initialize(&_V2LensHub.TransactOpts, name, symbol, newGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address newGovernance) returns()
func (_V2LensHub *V2LensHubTransactorSession) Initialize(name string, symbol string, newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.Initialize(&_V2LensHub.TransactOpts, name, symbol, newGovernance)
}

// Mirror is a paid mutator transaction binding the contract method 0xf90604d1.
//
// Solidity: function mirror((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) Mirror(opts *bind.TransactOpts, mirrorParams TypesMirrorParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "mirror", mirrorParams)
}

// Mirror is a paid mutator transaction binding the contract method 0xf90604d1.
//
// Solidity: function mirror((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) Mirror(mirrorParams TypesMirrorParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Mirror(&_V2LensHub.TransactOpts, mirrorParams)
}

// Mirror is a paid mutator transaction binding the contract method 0xf90604d1.
//
// Solidity: function mirror((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) Mirror(mirrorParams TypesMirrorParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Mirror(&_V2LensHub.TransactOpts, mirrorParams)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xe6a402b5.
//
// Solidity: function mirrorWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) MirrorWithSig(opts *bind.TransactOpts, mirrorParams TypesMirrorParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "mirrorWithSig", mirrorParams, signature)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xe6a402b5.
//
// Solidity: function mirrorWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubSession) MirrorWithSig(mirrorParams TypesMirrorParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.MirrorWithSig(&_V2LensHub.TransactOpts, mirrorParams, signature)
}

// MirrorWithSig is a paid mutator transaction binding the contract method 0xe6a402b5.
//
// Solidity: function mirrorWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes) mirrorParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) MirrorWithSig(mirrorParams TypesMirrorParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.MirrorWithSig(&_V2LensHub.TransactOpts, mirrorParams, signature)
}

// Post is a paid mutator transaction binding the contract method 0x66b0dcd3.
//
// Solidity: function post((uint256,string,address[],bytes[],address,bytes) postParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) Post(opts *bind.TransactOpts, postParams TypesPostParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "post", postParams)
}

// Post is a paid mutator transaction binding the contract method 0x66b0dcd3.
//
// Solidity: function post((uint256,string,address[],bytes[],address,bytes) postParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) Post(postParams TypesPostParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Post(&_V2LensHub.TransactOpts, postParams)
}

// Post is a paid mutator transaction binding the contract method 0x66b0dcd3.
//
// Solidity: function post((uint256,string,address[],bytes[],address,bytes) postParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) Post(postParams TypesPostParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Post(&_V2LensHub.TransactOpts, postParams)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x907cd7d2.
//
// Solidity: function postWithSig((uint256,string,address[],bytes[],address,bytes) postParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) PostWithSig(opts *bind.TransactOpts, postParams TypesPostParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "postWithSig", postParams, signature)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x907cd7d2.
//
// Solidity: function postWithSig((uint256,string,address[],bytes[],address,bytes) postParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubSession) PostWithSig(postParams TypesPostParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.PostWithSig(&_V2LensHub.TransactOpts, postParams, signature)
}

// PostWithSig is a paid mutator transaction binding the contract method 0x907cd7d2.
//
// Solidity: function postWithSig((uint256,string,address[],bytes[],address,bytes) postParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) PostWithSig(postParams TypesPostParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.PostWithSig(&_V2LensHub.TransactOpts, postParams, signature)
}

// Quote is a paid mutator transaction binding the contract method 0xf0ba35f0.
//
// Solidity: function quote((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) Quote(opts *bind.TransactOpts, quoteParams TypesQuoteParams) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "quote", quoteParams)
}

// Quote is a paid mutator transaction binding the contract method 0xf0ba35f0.
//
// Solidity: function quote((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams) returns(uint256)
func (_V2LensHub *V2LensHubSession) Quote(quoteParams TypesQuoteParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Quote(&_V2LensHub.TransactOpts, quoteParams)
}

// Quote is a paid mutator transaction binding the contract method 0xf0ba35f0.
//
// Solidity: function quote((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) Quote(quoteParams TypesQuoteParams) (*types.Transaction, error) {
	return _V2LensHub.Contract.Quote(&_V2LensHub.TransactOpts, quoteParams)
}

// QuoteWithSig is a paid mutator transaction binding the contract method 0x65f29f27.
//
// Solidity: function quoteWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactor) QuoteWithSig(opts *bind.TransactOpts, quoteParams TypesQuoteParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "quoteWithSig", quoteParams, signature)
}

// QuoteWithSig is a paid mutator transaction binding the contract method 0x65f29f27.
//
// Solidity: function quoteWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubSession) QuoteWithSig(quoteParams TypesQuoteParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.QuoteWithSig(&_V2LensHub.TransactOpts, quoteParams, signature)
}

// QuoteWithSig is a paid mutator transaction binding the contract method 0x65f29f27.
//
// Solidity: function quoteWithSig((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes) quoteParams, (address,uint8,bytes32,bytes32,uint256) signature) returns(uint256)
func (_V2LensHub *V2LensHubTransactorSession) QuoteWithSig(quoteParams TypesQuoteParams, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.QuoteWithSig(&_V2LensHub.TransactOpts, quoteParams, signature)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.SafeTransferFrom(&_V2LensHub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.SafeTransferFrom(&_V2LensHub.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V2LensHub *V2LensHubTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V2LensHub *V2LensHubSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.SafeTransferFrom0(&_V2LensHub.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_V2LensHub *V2LensHubTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.SafeTransferFrom0(&_V2LensHub.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V2LensHub *V2LensHubTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V2LensHub *V2LensHubSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetApprovalForAll(&_V2LensHub.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetApprovalForAll(&_V2LensHub.TransactOpts, operator, approved)
}

// SetBlockStatus is a paid mutator transaction binding the contract method 0xfb827e82.
//
// Solidity: function setBlockStatus(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus) returns()
func (_V2LensHub *V2LensHubTransactor) SetBlockStatus(opts *bind.TransactOpts, byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setBlockStatus", byProfileId, idsOfProfilesToSetBlockStatus, blockStatus)
}

// SetBlockStatus is a paid mutator transaction binding the contract method 0xfb827e82.
//
// Solidity: function setBlockStatus(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus) returns()
func (_V2LensHub *V2LensHubSession) SetBlockStatus(byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetBlockStatus(&_V2LensHub.TransactOpts, byProfileId, idsOfProfilesToSetBlockStatus, blockStatus)
}

// SetBlockStatus is a paid mutator transaction binding the contract method 0xfb827e82.
//
// Solidity: function setBlockStatus(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetBlockStatus(byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetBlockStatus(&_V2LensHub.TransactOpts, byProfileId, idsOfProfilesToSetBlockStatus, blockStatus)
}

// SetBlockStatusWithSig is a paid mutator transaction binding the contract method 0xc6d45944.
//
// Solidity: function setBlockStatusWithSig(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactor) SetBlockStatusWithSig(opts *bind.TransactOpts, byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setBlockStatusWithSig", byProfileId, idsOfProfilesToSetBlockStatus, blockStatus, signature)
}

// SetBlockStatusWithSig is a paid mutator transaction binding the contract method 0xc6d45944.
//
// Solidity: function setBlockStatusWithSig(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubSession) SetBlockStatusWithSig(byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetBlockStatusWithSig(&_V2LensHub.TransactOpts, byProfileId, idsOfProfilesToSetBlockStatus, blockStatus, signature)
}

// SetBlockStatusWithSig is a paid mutator transaction binding the contract method 0xc6d45944.
//
// Solidity: function setBlockStatusWithSig(uint256 byProfileId, uint256[] idsOfProfilesToSetBlockStatus, bool[] blockStatus, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetBlockStatusWithSig(byProfileId *big.Int, idsOfProfilesToSetBlockStatus []*big.Int, blockStatus []bool, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetBlockStatusWithSig(&_V2LensHub.TransactOpts, byProfileId, idsOfProfilesToSetBlockStatus, blockStatus, signature)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V2LensHub *V2LensHubTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setEmergencyAdmin", newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V2LensHub *V2LensHubSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetEmergencyAdmin(&_V2LensHub.TransactOpts, newEmergencyAdmin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address newEmergencyAdmin) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetEmergencyAdmin(newEmergencyAdmin common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetEmergencyAdmin(&_V2LensHub.TransactOpts, newEmergencyAdmin)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V2LensHub *V2LensHubTransactor) SetFollowModule(opts *bind.TransactOpts, profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setFollowModule", profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V2LensHub *V2LensHubSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetFollowModule(&_V2LensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModule is a paid mutator transaction binding the contract method 0x6dea40b3.
//
// Solidity: function setFollowModule(uint256 profileId, address followModule, bytes followModuleInitData) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetFollowModule(profileId *big.Int, followModule common.Address, followModuleInitData []byte) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetFollowModule(&_V2LensHub.TransactOpts, profileId, followModule, followModuleInitData)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0xce30bb4f.
//
// Solidity: function setFollowModuleWithSig(uint256 profileId, address followModule, bytes followModuleInitData, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactor) SetFollowModuleWithSig(opts *bind.TransactOpts, profileId *big.Int, followModule common.Address, followModuleInitData []byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setFollowModuleWithSig", profileId, followModule, followModuleInitData, signature)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0xce30bb4f.
//
// Solidity: function setFollowModuleWithSig(uint256 profileId, address followModule, bytes followModuleInitData, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubSession) SetFollowModuleWithSig(profileId *big.Int, followModule common.Address, followModuleInitData []byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetFollowModuleWithSig(&_V2LensHub.TransactOpts, profileId, followModule, followModuleInitData, signature)
}

// SetFollowModuleWithSig is a paid mutator transaction binding the contract method 0xce30bb4f.
//
// Solidity: function setFollowModuleWithSig(uint256 profileId, address followModule, bytes followModuleInitData, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetFollowModuleWithSig(profileId *big.Int, followModule common.Address, followModuleInitData []byte, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetFollowModuleWithSig(&_V2LensHub.TransactOpts, profileId, followModule, followModuleInitData, signature)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V2LensHub *V2LensHubTransactor) SetGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setGovernance", newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V2LensHub *V2LensHubSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetGovernance(&_V2LensHub.TransactOpts, newGovernance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address newGovernance) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetGovernance(&_V2LensHub.TransactOpts, newGovernance)
}

// SetMigrationAdmins is a paid mutator transaction binding the contract method 0x8e15512a.
//
// Solidity: function setMigrationAdmins(address[] migrationAdmins, bool whitelisted) returns()
func (_V2LensHub *V2LensHubTransactor) SetMigrationAdmins(opts *bind.TransactOpts, migrationAdmins []common.Address, whitelisted bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setMigrationAdmins", migrationAdmins, whitelisted)
}

// SetMigrationAdmins is a paid mutator transaction binding the contract method 0x8e15512a.
//
// Solidity: function setMigrationAdmins(address[] migrationAdmins, bool whitelisted) returns()
func (_V2LensHub *V2LensHubSession) SetMigrationAdmins(migrationAdmins []common.Address, whitelisted bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetMigrationAdmins(&_V2LensHub.TransactOpts, migrationAdmins, whitelisted)
}

// SetMigrationAdmins is a paid mutator transaction binding the contract method 0x8e15512a.
//
// Solidity: function setMigrationAdmins(address[] migrationAdmins, bool whitelisted) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetMigrationAdmins(migrationAdmins []common.Address, whitelisted bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetMigrationAdmins(&_V2LensHub.TransactOpts, migrationAdmins, whitelisted)
}

// SetProfileMetadataURI is a paid mutator transaction binding the contract method 0xefe4fd83.
//
// Solidity: function setProfileMetadataURI(uint256 profileId, string metadataURI) returns()
func (_V2LensHub *V2LensHubTransactor) SetProfileMetadataURI(opts *bind.TransactOpts, profileId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setProfileMetadataURI", profileId, metadataURI)
}

// SetProfileMetadataURI is a paid mutator transaction binding the contract method 0xefe4fd83.
//
// Solidity: function setProfileMetadataURI(uint256 profileId, string metadataURI) returns()
func (_V2LensHub *V2LensHubSession) SetProfileMetadataURI(profileId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetProfileMetadataURI(&_V2LensHub.TransactOpts, profileId, metadataURI)
}

// SetProfileMetadataURI is a paid mutator transaction binding the contract method 0xefe4fd83.
//
// Solidity: function setProfileMetadataURI(uint256 profileId, string metadataURI) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetProfileMetadataURI(profileId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetProfileMetadataURI(&_V2LensHub.TransactOpts, profileId, metadataURI)
}

// SetProfileMetadataURIWithSig is a paid mutator transaction binding the contract method 0x6d0fd658.
//
// Solidity: function setProfileMetadataURIWithSig(uint256 profileId, string metadataURI, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactor) SetProfileMetadataURIWithSig(opts *bind.TransactOpts, profileId *big.Int, metadataURI string, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setProfileMetadataURIWithSig", profileId, metadataURI, signature)
}

// SetProfileMetadataURIWithSig is a paid mutator transaction binding the contract method 0x6d0fd658.
//
// Solidity: function setProfileMetadataURIWithSig(uint256 profileId, string metadataURI, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubSession) SetProfileMetadataURIWithSig(profileId *big.Int, metadataURI string, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetProfileMetadataURIWithSig(&_V2LensHub.TransactOpts, profileId, metadataURI, signature)
}

// SetProfileMetadataURIWithSig is a paid mutator transaction binding the contract method 0x6d0fd658.
//
// Solidity: function setProfileMetadataURIWithSig(uint256 profileId, string metadataURI, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetProfileMetadataURIWithSig(profileId *big.Int, metadataURI string, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetProfileMetadataURIWithSig(&_V2LensHub.TransactOpts, profileId, metadataURI, signature)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x4209a2e1.
//
// Solidity: function setRoyalty(uint256 royaltiesInBasisPoints) returns()
func (_V2LensHub *V2LensHubTransactor) SetRoyalty(opts *bind.TransactOpts, royaltiesInBasisPoints *big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setRoyalty", royaltiesInBasisPoints)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x4209a2e1.
//
// Solidity: function setRoyalty(uint256 royaltiesInBasisPoints) returns()
func (_V2LensHub *V2LensHubSession) SetRoyalty(royaltiesInBasisPoints *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetRoyalty(&_V2LensHub.TransactOpts, royaltiesInBasisPoints)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x4209a2e1.
//
// Solidity: function setRoyalty(uint256 royaltiesInBasisPoints) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetRoyalty(royaltiesInBasisPoints *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetRoyalty(&_V2LensHub.TransactOpts, royaltiesInBasisPoints)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V2LensHub *V2LensHubTransactor) SetState(opts *bind.TransactOpts, newState uint8) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setState", newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V2LensHub *V2LensHubSession) SetState(newState uint8) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetState(&_V2LensHub.TransactOpts, newState)
}

// SetState is a paid mutator transaction binding the contract method 0x56de96db.
//
// Solidity: function setState(uint8 newState) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetState(newState uint8) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetState(&_V2LensHub.TransactOpts, newState)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_V2LensHub *V2LensHubTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_V2LensHub *V2LensHubSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetTreasury(&_V2LensHub.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetTreasury(&_V2LensHub.TransactOpts, newTreasury)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0xa652db49.
//
// Solidity: function setTreasuryFee(uint16 newTreasuryFee) returns()
func (_V2LensHub *V2LensHubTransactor) SetTreasuryFee(opts *bind.TransactOpts, newTreasuryFee uint16) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "setTreasuryFee", newTreasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0xa652db49.
//
// Solidity: function setTreasuryFee(uint16 newTreasuryFee) returns()
func (_V2LensHub *V2LensHubSession) SetTreasuryFee(newTreasuryFee uint16) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetTreasuryFee(&_V2LensHub.TransactOpts, newTreasuryFee)
}

// SetTreasuryFee is a paid mutator transaction binding the contract method 0xa652db49.
//
// Solidity: function setTreasuryFee(uint16 newTreasuryFee) returns()
func (_V2LensHub *V2LensHubTransactorSession) SetTreasuryFee(newTreasuryFee uint16) (*types.Transaction, error) {
	return _V2LensHub.Contract.SetTreasuryFee(&_V2LensHub.TransactOpts, newTreasuryFee)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.TransferFrom(&_V2LensHub.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_V2LensHub *V2LensHubTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.TransferFrom(&_V2LensHub.TransactOpts, from, to, tokenId)
}

// Unfollow is a paid mutator transaction binding the contract method 0x815ed04a.
//
// Solidity: function unfollow(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow) returns()
func (_V2LensHub *V2LensHubTransactor) Unfollow(opts *bind.TransactOpts, unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "unfollow", unfollowerProfileId, idsOfProfilesToUnfollow)
}

// Unfollow is a paid mutator transaction binding the contract method 0x815ed04a.
//
// Solidity: function unfollow(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow) returns()
func (_V2LensHub *V2LensHubSession) Unfollow(unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Unfollow(&_V2LensHub.TransactOpts, unfollowerProfileId, idsOfProfilesToUnfollow)
}

// Unfollow is a paid mutator transaction binding the contract method 0x815ed04a.
//
// Solidity: function unfollow(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow) returns()
func (_V2LensHub *V2LensHubTransactorSession) Unfollow(unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int) (*types.Transaction, error) {
	return _V2LensHub.Contract.Unfollow(&_V2LensHub.TransactOpts, unfollowerProfileId, idsOfProfilesToUnfollow)
}

// UnfollowWithSig is a paid mutator transaction binding the contract method 0x809d8947.
//
// Solidity: function unfollowWithSig(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactor) UnfollowWithSig(opts *bind.TransactOpts, unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "unfollowWithSig", unfollowerProfileId, idsOfProfilesToUnfollow, signature)
}

// UnfollowWithSig is a paid mutator transaction binding the contract method 0x809d8947.
//
// Solidity: function unfollowWithSig(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubSession) UnfollowWithSig(unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.UnfollowWithSig(&_V2LensHub.TransactOpts, unfollowerProfileId, idsOfProfilesToUnfollow, signature)
}

// UnfollowWithSig is a paid mutator transaction binding the contract method 0x809d8947.
//
// Solidity: function unfollowWithSig(uint256 unfollowerProfileId, uint256[] idsOfProfilesToUnfollow, (address,uint8,bytes32,bytes32,uint256) signature) returns()
func (_V2LensHub *V2LensHubTransactorSession) UnfollowWithSig(unfollowerProfileId *big.Int, idsOfProfilesToUnfollow []*big.Int, signature TypesEIP712Signature) (*types.Transaction, error) {
	return _V2LensHub.Contract.UnfollowWithSig(&_V2LensHub.TransactOpts, unfollowerProfileId, idsOfProfilesToUnfollow, signature)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V2LensHub *V2LensHubTransactor) WhitelistProfileCreator(opts *bind.TransactOpts, profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V2LensHub.contract.Transact(opts, "whitelistProfileCreator", profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V2LensHub *V2LensHubSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.WhitelistProfileCreator(&_V2LensHub.TransactOpts, profileCreator, whitelist)
}

// WhitelistProfileCreator is a paid mutator transaction binding the contract method 0x20905506.
//
// Solidity: function whitelistProfileCreator(address profileCreator, bool whitelist) returns()
func (_V2LensHub *V2LensHubTransactorSession) WhitelistProfileCreator(profileCreator common.Address, whitelist bool) (*types.Transaction, error) {
	return _V2LensHub.Contract.WhitelistProfileCreator(&_V2LensHub.TransactOpts, profileCreator, whitelist)
}

// V2LensHubApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the V2LensHub contract.
type V2LensHubApprovalIterator struct {
	Event *V2LensHubApproval // Event containing the contract specifics and raw log

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
func (it *V2LensHubApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubApproval)
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
		it.Event = new(V2LensHubApproval)
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
func (it *V2LensHubApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubApproval represents a Approval event raised by the V2LensHub contract.
type V2LensHubApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V2LensHub *V2LensHubFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*V2LensHubApprovalIterator, error) {

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

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubApprovalIterator{contract: _V2LensHub.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_V2LensHub *V2LensHubFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *V2LensHubApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubApproval)
				if err := _V2LensHub.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_V2LensHub *V2LensHubFilterer) ParseApproval(log types.Log) (*V2LensHubApproval, error) {
	event := new(V2LensHubApproval)
	if err := _V2LensHub.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the V2LensHub contract.
type V2LensHubApprovalForAllIterator struct {
	Event *V2LensHubApprovalForAll // Event containing the contract specifics and raw log

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
func (it *V2LensHubApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubApprovalForAll)
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
		it.Event = new(V2LensHubApprovalForAll)
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
func (it *V2LensHubApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubApprovalForAll represents a ApprovalForAll event raised by the V2LensHub contract.
type V2LensHubApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V2LensHub *V2LensHubFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*V2LensHubApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubApprovalForAllIterator{contract: _V2LensHub.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_V2LensHub *V2LensHubFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *V2LensHubApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubApprovalForAll)
				if err := _V2LensHub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_V2LensHub *V2LensHubFilterer) ParseApprovalForAll(log types.Log) (*V2LensHubApprovalForAll, error) {
	event := new(V2LensHubApprovalForAll)
	if err := _V2LensHub.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubCollectNFTTransferredIterator is returned from FilterCollectNFTTransferred and is used to iterate over the raw logs and unpacked data for CollectNFTTransferred events raised by the V2LensHub contract.
type V2LensHubCollectNFTTransferredIterator struct {
	Event *V2LensHubCollectNFTTransferred // Event containing the contract specifics and raw log

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
func (it *V2LensHubCollectNFTTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubCollectNFTTransferred)
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
		it.Event = new(V2LensHubCollectNFTTransferred)
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
func (it *V2LensHubCollectNFTTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubCollectNFTTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubCollectNFTTransferred represents a CollectNFTTransferred event raised by the V2LensHub contract.
type V2LensHubCollectNFTTransferred struct {
	ProfileId    *big.Int
	PubId        *big.Int
	CollectNFTId *big.Int
	From         common.Address
	To           common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectNFTTransferred is a free log retrieval operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterCollectNFTTransferred(opts *bind.FilterOpts, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (*V2LensHubCollectNFTTransferredIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubCollectNFTTransferredIterator{contract: _V2LensHub.contract, event: "CollectNFTTransferred", logs: logs, sub: sub}, nil
}

// WatchCollectNFTTransferred is a free log subscription operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchCollectNFTTransferred(opts *bind.WatchOpts, sink chan<- *V2LensHubCollectNFTTransferred, profileId []*big.Int, pubId []*big.Int, collectNFTId []*big.Int) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var pubIdRule []interface{}
	for _, pubIdItem := range pubId {
		pubIdRule = append(pubIdRule, pubIdItem)
	}
	var collectNFTIdRule []interface{}
	for _, collectNFTIdItem := range collectNFTId {
		collectNFTIdRule = append(collectNFTIdRule, collectNFTIdItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "CollectNFTTransferred", profileIdRule, pubIdRule, collectNFTIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubCollectNFTTransferred)
				if err := _V2LensHub.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
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

// ParseCollectNFTTransferred is a log parse operation binding the contract event 0x68edb7ec2c37d21b3b72233960b487f2966f4ac82b7430d39f24d1f8d6f99106.
//
// Solidity: event CollectNFTTransferred(uint256 indexed profileId, uint256 indexed pubId, uint256 indexed collectNFTId, address from, address to, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseCollectNFTTransferred(log types.Log) (*V2LensHubCollectNFTTransferred, error) {
	event := new(V2LensHubCollectNFTTransferred)
	if err := _V2LensHub.contract.UnpackLog(event, "CollectNFTTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubLensUpgradeVersionIterator is returned from FilterLensUpgradeVersion and is used to iterate over the raw logs and unpacked data for LensUpgradeVersion events raised by the V2LensHub contract.
type V2LensHubLensUpgradeVersionIterator struct {
	Event *V2LensHubLensUpgradeVersion // Event containing the contract specifics and raw log

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
func (it *V2LensHubLensUpgradeVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubLensUpgradeVersion)
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
		it.Event = new(V2LensHubLensUpgradeVersion)
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
func (it *V2LensHubLensUpgradeVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubLensUpgradeVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubLensUpgradeVersion represents a LensUpgradeVersion event raised by the V2LensHub contract.
type V2LensHubLensUpgradeVersion struct {
	Implementation common.Address
	Version        string
	GitCommit      [20]byte
	Timestamp      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLensUpgradeVersion is a free log retrieval operation binding the contract event 0x1ac53e0fe87e82f3d54f8bb8042dee65a6f997a3a15e75730153b8c546c0d280.
//
// Solidity: event LensUpgradeVersion(address implementation, string version, bytes20 gitCommit, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterLensUpgradeVersion(opts *bind.FilterOpts) (*V2LensHubLensUpgradeVersionIterator, error) {

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "LensUpgradeVersion")
	if err != nil {
		return nil, err
	}
	return &V2LensHubLensUpgradeVersionIterator{contract: _V2LensHub.contract, event: "LensUpgradeVersion", logs: logs, sub: sub}, nil
}

// WatchLensUpgradeVersion is a free log subscription operation binding the contract event 0x1ac53e0fe87e82f3d54f8bb8042dee65a6f997a3a15e75730153b8c546c0d280.
//
// Solidity: event LensUpgradeVersion(address implementation, string version, bytes20 gitCommit, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchLensUpgradeVersion(opts *bind.WatchOpts, sink chan<- *V2LensHubLensUpgradeVersion) (event.Subscription, error) {

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "LensUpgradeVersion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubLensUpgradeVersion)
				if err := _V2LensHub.contract.UnpackLog(event, "LensUpgradeVersion", log); err != nil {
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

// ParseLensUpgradeVersion is a log parse operation binding the contract event 0x1ac53e0fe87e82f3d54f8bb8042dee65a6f997a3a15e75730153b8c546c0d280.
//
// Solidity: event LensUpgradeVersion(address implementation, string version, bytes20 gitCommit, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseLensUpgradeVersion(log types.Log) (*V2LensHubLensUpgradeVersion, error) {
	event := new(V2LensHubLensUpgradeVersion)
	if err := _V2LensHub.contract.UnpackLog(event, "LensUpgradeVersion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubTokenGuardianStateChangedIterator is returned from FilterTokenGuardianStateChanged and is used to iterate over the raw logs and unpacked data for TokenGuardianStateChanged events raised by the V2LensHub contract.
type V2LensHubTokenGuardianStateChangedIterator struct {
	Event *V2LensHubTokenGuardianStateChanged // Event containing the contract specifics and raw log

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
func (it *V2LensHubTokenGuardianStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubTokenGuardianStateChanged)
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
		it.Event = new(V2LensHubTokenGuardianStateChanged)
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
func (it *V2LensHubTokenGuardianStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubTokenGuardianStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubTokenGuardianStateChanged represents a TokenGuardianStateChanged event raised by the V2LensHub contract.
type V2LensHubTokenGuardianStateChanged struct {
	Wallet                          common.Address
	Enabled                         bool
	TokenGuardianDisablingTimestamp *big.Int
	Timestamp                       *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterTokenGuardianStateChanged is a free log retrieval operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterTokenGuardianStateChanged(opts *bind.FilterOpts, wallet []common.Address, enabled []bool) (*V2LensHubTokenGuardianStateChangedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "TokenGuardianStateChanged", walletRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubTokenGuardianStateChangedIterator{contract: _V2LensHub.contract, event: "TokenGuardianStateChanged", logs: logs, sub: sub}, nil
}

// WatchTokenGuardianStateChanged is a free log subscription operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchTokenGuardianStateChanged(opts *bind.WatchOpts, sink chan<- *V2LensHubTokenGuardianStateChanged, wallet []common.Address, enabled []bool) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var enabledRule []interface{}
	for _, enabledItem := range enabled {
		enabledRule = append(enabledRule, enabledItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "TokenGuardianStateChanged", walletRule, enabledRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubTokenGuardianStateChanged)
				if err := _V2LensHub.contract.UnpackLog(event, "TokenGuardianStateChanged", log); err != nil {
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

// ParseTokenGuardianStateChanged is a log parse operation binding the contract event 0x035adf3bbe16b317cf4a3e05c966ea6571d1af00147c5f121bd1514b1e322a06.
//
// Solidity: event TokenGuardianStateChanged(address indexed wallet, bool indexed enabled, uint256 tokenGuardianDisablingTimestamp, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseTokenGuardianStateChanged(log types.Log) (*V2LensHubTokenGuardianStateChanged, error) {
	event := new(V2LensHubTokenGuardianStateChanged)
	if err := _V2LensHub.contract.UnpackLog(event, "TokenGuardianStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the V2LensHub contract.
type V2LensHubTransferIterator struct {
	Event *V2LensHubTransfer // Event containing the contract specifics and raw log

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
func (it *V2LensHubTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubTransfer)
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
		it.Event = new(V2LensHubTransfer)
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
func (it *V2LensHubTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubTransfer represents a Transfer event raised by the V2LensHub contract.
type V2LensHubTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V2LensHub *V2LensHubFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*V2LensHubTransferIterator, error) {

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

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubTransferIterator{contract: _V2LensHub.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_V2LensHub *V2LensHubFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *V2LensHubTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubTransfer)
				if err := _V2LensHub.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_V2LensHub *V2LensHubFilterer) ParseTransfer(log types.Log) (*V2LensHubTransfer, error) {
	event := new(V2LensHubTransfer)
	if err := _V2LensHub.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubTreasuryFeeSetIterator is returned from FilterTreasuryFeeSet and is used to iterate over the raw logs and unpacked data for TreasuryFeeSet events raised by the V2LensHub contract.
type V2LensHubTreasuryFeeSetIterator struct {
	Event *V2LensHubTreasuryFeeSet // Event containing the contract specifics and raw log

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
func (it *V2LensHubTreasuryFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubTreasuryFeeSet)
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
		it.Event = new(V2LensHubTreasuryFeeSet)
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
func (it *V2LensHubTreasuryFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubTreasuryFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubTreasuryFeeSet represents a TreasuryFeeSet event raised by the V2LensHub contract.
type V2LensHubTreasuryFeeSet struct {
	PrevTreasuryFee uint16
	NewTreasuryFee  uint16
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTreasuryFeeSet is a free log retrieval operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterTreasuryFeeSet(opts *bind.FilterOpts, prevTreasuryFee []uint16, newTreasuryFee []uint16) (*V2LensHubTreasuryFeeSetIterator, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "TreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubTreasuryFeeSetIterator{contract: _V2LensHub.contract, event: "TreasuryFeeSet", logs: logs, sub: sub}, nil
}

// WatchTreasuryFeeSet is a free log subscription operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchTreasuryFeeSet(opts *bind.WatchOpts, sink chan<- *V2LensHubTreasuryFeeSet, prevTreasuryFee []uint16, newTreasuryFee []uint16) (event.Subscription, error) {

	var prevTreasuryFeeRule []interface{}
	for _, prevTreasuryFeeItem := range prevTreasuryFee {
		prevTreasuryFeeRule = append(prevTreasuryFeeRule, prevTreasuryFeeItem)
	}
	var newTreasuryFeeRule []interface{}
	for _, newTreasuryFeeItem := range newTreasuryFee {
		newTreasuryFeeRule = append(newTreasuryFeeRule, newTreasuryFeeItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "TreasuryFeeSet", prevTreasuryFeeRule, newTreasuryFeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubTreasuryFeeSet)
				if err := _V2LensHub.contract.UnpackLog(event, "TreasuryFeeSet", log); err != nil {
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

// ParseTreasuryFeeSet is a log parse operation binding the contract event 0x6076e665d0cd29a9fb0391c62a1c3c1c6d18531bf470fce88abbc7f33b855f7f.
//
// Solidity: event TreasuryFeeSet(uint16 indexed prevTreasuryFee, uint16 indexed newTreasuryFee, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseTreasuryFeeSet(log types.Log) (*V2LensHubTreasuryFeeSet, error) {
	event := new(V2LensHubTreasuryFeeSet)
	if err := _V2LensHub.contract.UnpackLog(event, "TreasuryFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubTreasurySetIterator is returned from FilterTreasurySet and is used to iterate over the raw logs and unpacked data for TreasurySet events raised by the V2LensHub contract.
type V2LensHubTreasurySetIterator struct {
	Event *V2LensHubTreasurySet // Event containing the contract specifics and raw log

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
func (it *V2LensHubTreasurySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubTreasurySet)
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
		it.Event = new(V2LensHubTreasurySet)
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
func (it *V2LensHubTreasurySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubTreasurySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubTreasurySet represents a TreasurySet event raised by the V2LensHub contract.
type V2LensHubTreasurySet struct {
	PrevTreasury common.Address
	NewTreasury  common.Address
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTreasurySet is a free log retrieval operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterTreasurySet(opts *bind.FilterOpts, prevTreasury []common.Address, newTreasury []common.Address) (*V2LensHubTreasurySetIterator, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "TreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubTreasurySetIterator{contract: _V2LensHub.contract, event: "TreasurySet", logs: logs, sub: sub}, nil
}

// WatchTreasurySet is a free log subscription operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchTreasurySet(opts *bind.WatchOpts, sink chan<- *V2LensHubTreasurySet, prevTreasury []common.Address, newTreasury []common.Address) (event.Subscription, error) {

	var prevTreasuryRule []interface{}
	for _, prevTreasuryItem := range prevTreasury {
		prevTreasuryRule = append(prevTreasuryRule, prevTreasuryItem)
	}
	var newTreasuryRule []interface{}
	for _, newTreasuryItem := range newTreasury {
		newTreasuryRule = append(newTreasuryRule, newTreasuryItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "TreasurySet", prevTreasuryRule, newTreasuryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubTreasurySet)
				if err := _V2LensHub.contract.UnpackLog(event, "TreasurySet", log); err != nil {
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

// ParseTreasurySet is a log parse operation binding the contract event 0x8df20312a19939ae72e29e4500356a05238ef5e6794a3323c184d97bb615d4fe.
//
// Solidity: event TreasurySet(address indexed prevTreasury, address indexed newTreasury, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseTreasurySet(log types.Log) (*V2LensHubTreasurySet, error) {
	event := new(V2LensHubTreasurySet)
	if err := _V2LensHub.contract.UnpackLog(event, "TreasurySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// V2LensHubUnfollowedIterator is returned from FilterUnfollowed and is used to iterate over the raw logs and unpacked data for Unfollowed events raised by the V2LensHub contract.
type V2LensHubUnfollowedIterator struct {
	Event *V2LensHubUnfollowed // Event containing the contract specifics and raw log

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
func (it *V2LensHubUnfollowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(V2LensHubUnfollowed)
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
		it.Event = new(V2LensHubUnfollowed)
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
func (it *V2LensHubUnfollowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *V2LensHubUnfollowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// V2LensHubUnfollowed represents a Unfollowed event raised by the V2LensHub contract.
type V2LensHubUnfollowed struct {
	UnfollowerProfileId   *big.Int
	IdOfProfileUnfollowed *big.Int
	TransactionExecutor   common.Address
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUnfollowed is a free log retrieval operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) FilterUnfollowed(opts *bind.FilterOpts, unfollowerProfileId []*big.Int) (*V2LensHubUnfollowedIterator, error) {

	var unfollowerProfileIdRule []interface{}
	for _, unfollowerProfileIdItem := range unfollowerProfileId {
		unfollowerProfileIdRule = append(unfollowerProfileIdRule, unfollowerProfileIdItem)
	}

	logs, sub, err := _V2LensHub.contract.FilterLogs(opts, "Unfollowed", unfollowerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return &V2LensHubUnfollowedIterator{contract: _V2LensHub.contract, event: "Unfollowed", logs: logs, sub: sub}, nil
}

// WatchUnfollowed is a free log subscription operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) WatchUnfollowed(opts *bind.WatchOpts, sink chan<- *V2LensHubUnfollowed, unfollowerProfileId []*big.Int) (event.Subscription, error) {

	var unfollowerProfileIdRule []interface{}
	for _, unfollowerProfileIdItem := range unfollowerProfileId {
		unfollowerProfileIdRule = append(unfollowerProfileIdRule, unfollowerProfileIdItem)
	}

	logs, sub, err := _V2LensHub.contract.WatchLogs(opts, "Unfollowed", unfollowerProfileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(V2LensHubUnfollowed)
				if err := _V2LensHub.contract.UnpackLog(event, "Unfollowed", log); err != nil {
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

// ParseUnfollowed is a log parse operation binding the contract event 0x9bbadc4d29f8416b3b1ed6fe7b42cc3588aaca742ac8c1661b3bb0a4c5ab1673.
//
// Solidity: event Unfollowed(uint256 indexed unfollowerProfileId, uint256 idOfProfileUnfollowed, address transactionExecutor, uint256 timestamp)
func (_V2LensHub *V2LensHubFilterer) ParseUnfollowed(log types.Log) (*V2LensHubUnfollowed, error) {
	event := new(V2LensHubUnfollowed)
	if err := _V2LensHub.contract.UnpackLog(event, "Unfollowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
