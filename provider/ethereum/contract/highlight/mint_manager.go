// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package highlight

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

// IAbridgedMintVectorAbridgedVector is an auto generated low-level Go binding around an user-defined struct.
type IAbridgedMintVectorAbridgedVector struct {
	ContractAddress            common.Address
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	PaymentRecipient           common.Address
	MaxTotalClaimableViaVector *big.Int
	TotalClaimedViaVector      *big.Int
	Currency                   common.Address
	TokenLimitPerTx            *big.Int
	MaxUserClaimableViaVector  *big.Int
	PricePerToken              *big.Int
	EditionId                  *big.Int
	EditionBasedCollection     bool
	RequireDirectEOA           bool
	AllowlistRoot              [32]byte
}

// IAbridgedMintVectorAbridgedVectorData is an auto generated low-level Go binding around an user-defined struct.
type IAbridgedMintVectorAbridgedVectorData struct {
	ContractAddress            *big.Int
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	PaymentRecipient           *big.Int
	MaxTotalClaimableViaVector *big.Int
	TotalClaimedViaVector      *big.Int
	Currency                   *big.Int
	TokenLimitPerTx            *big.Int
	MaxUserClaimableViaVector  *big.Int
	PricePerToken              *big.Int
	EditionId                  *big.Int
	EditionBasedCollection     bool
	RequireDirectEOA           bool
	AllowlistRoot              [32]byte
}

// IAbridgedMintVectorUpdateAbridgedVectorConfig is an auto generated low-level Go binding around an user-defined struct.
type IAbridgedMintVectorUpdateAbridgedVectorConfig struct {
	UpdateStartTimestamp             uint16
	UpdateEndTimestamp               uint16
	UpdatePaymentRecipient           uint16
	UpdateMaxTotalClaimableViaVector uint16
	UpdateTokenLimitPerTx            uint16
	UpdateMaxUserClaimableViaVector  uint16
	UpdatePricePerToken              uint8
	UpdateAllowlistRoot              uint8
	UpdateRequireDirectEOA           uint8
	UpdateMetadata                   uint8
}

// MintManagerClaim is an auto generated low-level Go binding around an user-defined struct.
type MintManagerClaim struct {
	Currency              common.Address
	ContractAddress       common.Address
	Claimer               common.Address
	PaymentRecipient      common.Address
	PricePerToken         *big.Int
	NumTokensToMint       uint64
	MaxClaimableViaVector *big.Int
	MaxClaimablePerUser   *big.Int
	EditionId             *big.Int
	ClaimExpiryTimestamp  *big.Int
	ClaimNonce            [32]byte
	OffchainVectorId      [32]byte
}

// MintManagerClaimWithMetaTxPacket is an auto generated low-level Go binding around an user-defined struct.
type MintManagerClaimWithMetaTxPacket struct {
	Currency                 common.Address
	ContractAddress          common.Address
	Claimer                  common.Address
	PricePerToken            *big.Int
	NumTokensToMint          uint64
	PurchaseToCreatorPacket  MintManagerPurchaserMetaTxPacket
	PurchaseToPlatformPacket MintManagerPurchaserMetaTxPacket
	MaxClaimableViaVector    *big.Int
	MaxClaimablePerUser      *big.Int
	EditionId                *big.Int
	ClaimExpiryTimestamp     *big.Int
	ClaimNonce               [32]byte
	OffchainVectorId         [32]byte
}

// MintManagerPurchaserMetaTxPacket is an auto generated low-level Go binding around an user-defined struct.
type MintManagerPurchaserMetaTxPacket struct {
	FunctionSignature []byte
	SigR              [32]byte
	SigS              [32]byte
	SigV              uint8
}

// MintManagerSeriesClaim is an auto generated low-level Go binding around an user-defined struct.
type MintManagerSeriesClaim struct {
	Currency              common.Address
	ContractAddress       common.Address
	Claimer               common.Address
	PaymentRecipient      common.Address
	PricePerToken         *big.Int
	MaxPerTxn             uint64
	MaxClaimableViaVector uint64
	MaxClaimablePerUser   uint64
	ClaimExpiryTimestamp  uint64
	ClaimNonce            [32]byte
	OffchainVectorId      [32]byte
}

// MintManagerMetaData contains all meta data concerning the MintManager contract.
var MintManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AllowlistInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyTypeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EtherSendFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExecutorChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTotalClaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintFeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnchainVectorMintGuardFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotClaimer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotDirectEOA\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeMintRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VectorUpdateActionFrozen\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VectorWrongCollectionType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"onChainVector\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"ChooseTokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToCreator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"percentageBPSOfTotal\",\"type\":\"uint32\"}],\"name\":\"ERC20Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMintManager.PurchaserMetaTxPacket\",\"name\":\"purchaseToCreatorPacket\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMintManager.PurchaserMetaTxPacket\",\"name\":\"purchaseToPlatformPacket\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20PaymentMetaTxPackets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint48\",\"name\":\"editionId\",\"type\":\"uint48\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"EditionVectorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToCreator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"percentageBPSOfTotal\",\"type\":\"uint32\"}],\"name\":\"NativeGasTokenPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"onChainVector\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numMinted\",\"type\":\"uint256\"}],\"name\":\"NumTokenMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"added\",\"type\":\"bool\"}],\"name\":\"PlatformExecutorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"SeriesVectorCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"}],\"name\":\"VectorDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"flexibleData\",\"type\":\"uint128\"}],\"name\":\"VectorMetadataSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"}],\"name\":\"VectorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"addPlatformExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint160\",\"name\":\"contractAddress\",\"type\":\"uint160\"},{\"internalType\":\"uint48\",\"name\":\"startTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"paymentRecipient\",\"type\":\"uint160\"},{\"internalType\":\"uint48\",\"name\":\"maxTotalClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"totalClaimedViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint160\",\"name\":\"currency\",\"type\":\"uint160\"},{\"internalType\":\"uint48\",\"name\":\"tokenLimitPerTx\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"maxUserClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint192\",\"name\":\"pricePerToken\",\"type\":\"uint192\"},{\"internalType\":\"uint48\",\"name\":\"editionId\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"editionBasedCollection\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireDirectEOA\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"allowlistRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAbridgedMintVector.AbridgedVectorData\",\"name\":\"_vector\",\"type\":\"tuple\"}],\"name\":\"createAbridgedVector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"}],\"name\":\"deleteAbridgedVector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"deprecatePlatformExecutor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"numTokensToMint\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.Claim\",\"name\":\"_claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"gatedMintEdition721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"numTokensToMint\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.Claim\",\"name\":\"claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"claimSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"}],\"name\":\"gatedSeriesMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"maxPerTxn\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.SeriesClaim\",\"name\":\"claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"claimSignature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"gatedSeriesMintChooseToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"}],\"name\":\"getAbridgedVector\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"startTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"maxTotalClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"totalClaimedViaVector\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tokenLimitPerTx\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"maxUserClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint192\",\"name\":\"pricePerToken\",\"type\":\"uint192\"},{\"internalType\":\"uint48\",\"name\":\"editionId\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"editionBasedCollection\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireDirectEOA\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"allowlistRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAbridgedMintVector.AbridgedVector\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"}],\"name\":\"getAbridgedVectorMetadata\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"}],\"name\":\"getClaimNoncesUsedForOffchainVector\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNumClaimedPerUserOffchainVector\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"platform\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialExecutor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialPlatformMintFee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"vectorId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nonce\",\"type\":\"bytes32\"}],\"name\":\"isNonceUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"offchainVectorsClaimState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numClaimed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformExecutors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"flexibleData\",\"type\":\"uint128\"}],\"name\":\"setAbridgedVectorMetadata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"startTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"endTimestamp\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"maxTotalClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"totalClaimedViaVector\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint48\",\"name\":\"tokenLimitPerTx\",\"type\":\"uint48\"},{\"internalType\":\"uint48\",\"name\":\"maxUserClaimableViaVector\",\"type\":\"uint48\"},{\"internalType\":\"uint192\",\"name\":\"pricePerToken\",\"type\":\"uint192\"},{\"internalType\":\"uint48\",\"name\":\"editionId\",\"type\":\"uint48\"},{\"internalType\":\"bool\",\"name\":\"editionBasedCollection\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireDirectEOA\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"allowlistRoot\",\"type\":\"bytes32\"}],\"internalType\":\"structIAbridgedMintVector.AbridgedVector\",\"name\":\"_newVector\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"updateStartTimestamp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"updateEndTimestamp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"updatePaymentRecipient\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"updateMaxTotalClaimableViaVector\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"updateTokenLimitPerTx\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"updateMaxUserClaimableViaVector\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"updatePricePerToken\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"updateAllowlistRoot\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"updateRequireDirectEOA\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"updateMetadata\",\"type\":\"uint8\"}],\"internalType\":\"structIAbridgedMintVector.UpdateAbridgedVectorConfig\",\"name\":\"updateConfig\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"pause\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"flexibleData\",\"type\":\"uint128\"}],\"name\":\"updateAbridgedVector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPlatformMintFee\",\"type\":\"uint256\"}],\"name\":\"updatePlatformMintFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userClaims\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"numTokensToMint\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"}],\"name\":\"vectorMintEdition721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"numTokensToMint\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"vectorMintEdition721WithAllowlist\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"numTokensToMint\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"}],\"name\":\"vectorMintSeries721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"vectorId\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"numTokensToMint\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"vectorMintSeries721WithAllowlist\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vectorMutabilities\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"updatesFrozen\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"deleteFrozen\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"pausesFrozen\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vectorToEditionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vectors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"tokenLimitPerTx\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxTotalClaimableViaVector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxUserClaimableViaVector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"totalClaimedViaVector\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"allowlistRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"paused\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"numTokensToMint\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.Claim\",\"name\":\"claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedMsgSender\",\"type\":\"address\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"numTokensToMint\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"internalType\":\"structMintManager.PurchaserMetaTxPacket\",\"name\":\"purchaseToCreatorPacket\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"internalType\":\"structMintManager.PurchaserMetaTxPacket\",\"name\":\"purchaseToPlatformPacket\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"editionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.ClaimWithMetaTxPacket\",\"name\":\"claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedMsgSender\",\"type\":\"address\"}],\"name\":\"verifyClaimWithMetaTxPacket\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"paymentRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"maxPerTxn\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxClaimableViaVector\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxClaimablePerUser\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"claimExpiryTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"claimNonce\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"offchainVectorId\",\"type\":\"bytes32\"}],\"internalType\":\"structMintManager.SeriesClaim\",\"name\":\"claim\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"expectedMsgSender\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"verifySeriesClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawNativeGasToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MintManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use MintManagerMetaData.ABI instead.
var MintManagerABI = MintManagerMetaData.ABI

// MintManager is an auto generated Go binding around an Ethereum contract.
type MintManager struct {
	MintManagerCaller     // Read-only binding to the contract
	MintManagerTransactor // Write-only binding to the contract
	MintManagerFilterer   // Log filterer for contract events
}

// MintManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintManagerSession struct {
	Contract     *MintManager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintManagerCallerSession struct {
	Contract *MintManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MintManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintManagerTransactorSession struct {
	Contract     *MintManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MintManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintManagerRaw struct {
	Contract *MintManager // Generic contract binding to access the raw methods on
}

// MintManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintManagerCallerRaw struct {
	Contract *MintManagerCaller // Generic read-only contract binding to access the raw methods on
}

// MintManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintManagerTransactorRaw struct {
	Contract *MintManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintManager creates a new instance of MintManager, bound to a specific deployed contract.
func NewMintManager(address common.Address, backend bind.ContractBackend) (*MintManager, error) {
	contract, err := bindMintManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintManager{MintManagerCaller: MintManagerCaller{contract: contract}, MintManagerTransactor: MintManagerTransactor{contract: contract}, MintManagerFilterer: MintManagerFilterer{contract: contract}}, nil
}

// NewMintManagerCaller creates a new read-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerCaller(address common.Address, caller bind.ContractCaller) (*MintManagerCaller, error) {
	contract, err := bindMintManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerCaller{contract: contract}, nil
}

// NewMintManagerTransactor creates a new write-only instance of MintManager, bound to a specific deployed contract.
func NewMintManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*MintManagerTransactor, error) {
	contract, err := bindMintManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintManagerTransactor{contract: contract}, nil
}

// NewMintManagerFilterer creates a new log filterer instance of MintManager, bound to a specific deployed contract.
func NewMintManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*MintManagerFilterer, error) {
	contract, err := bindMintManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintManagerFilterer{contract: contract}, nil
}

// bindMintManager binds a generic wrapper to an already deployed contract.
func bindMintManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.MintManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.MintManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintManager *MintManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintManager *MintManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintManager *MintManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintManager.Contract.contract.Transact(opts, method, params...)
}

// GetAbridgedVector is a free data retrieval call binding the contract method 0x8a320d64.
//
// Solidity: function getAbridgedVector(uint256 vectorId) view returns((address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32))
func (_MintManager *MintManagerCaller) GetAbridgedVector(opts *bind.CallOpts, vectorId *big.Int) (IAbridgedMintVectorAbridgedVector, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "getAbridgedVector", vectorId)

	if err != nil {
		return *new(IAbridgedMintVectorAbridgedVector), err
	}

	out0 := *abi.ConvertType(out[0], new(IAbridgedMintVectorAbridgedVector)).(*IAbridgedMintVectorAbridgedVector)

	return out0, err

}

// GetAbridgedVector is a free data retrieval call binding the contract method 0x8a320d64.
//
// Solidity: function getAbridgedVector(uint256 vectorId) view returns((address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32))
func (_MintManager *MintManagerSession) GetAbridgedVector(vectorId *big.Int) (IAbridgedMintVectorAbridgedVector, error) {
	return _MintManager.Contract.GetAbridgedVector(&_MintManager.CallOpts, vectorId)
}

// GetAbridgedVector is a free data retrieval call binding the contract method 0x8a320d64.
//
// Solidity: function getAbridgedVector(uint256 vectorId) view returns((address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32))
func (_MintManager *MintManagerCallerSession) GetAbridgedVector(vectorId *big.Int) (IAbridgedMintVectorAbridgedVector, error) {
	return _MintManager.Contract.GetAbridgedVector(&_MintManager.CallOpts, vectorId)
}

// GetAbridgedVectorMetadata is a free data retrieval call binding the contract method 0xa0f76523.
//
// Solidity: function getAbridgedVectorMetadata(uint256 vectorId) view returns(bool, uint128)
func (_MintManager *MintManagerCaller) GetAbridgedVectorMetadata(opts *bind.CallOpts, vectorId *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "getAbridgedVectorMetadata", vectorId)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAbridgedVectorMetadata is a free data retrieval call binding the contract method 0xa0f76523.
//
// Solidity: function getAbridgedVectorMetadata(uint256 vectorId) view returns(bool, uint128)
func (_MintManager *MintManagerSession) GetAbridgedVectorMetadata(vectorId *big.Int) (bool, *big.Int, error) {
	return _MintManager.Contract.GetAbridgedVectorMetadata(&_MintManager.CallOpts, vectorId)
}

// GetAbridgedVectorMetadata is a free data retrieval call binding the contract method 0xa0f76523.
//
// Solidity: function getAbridgedVectorMetadata(uint256 vectorId) view returns(bool, uint128)
func (_MintManager *MintManagerCallerSession) GetAbridgedVectorMetadata(vectorId *big.Int) (bool, *big.Int, error) {
	return _MintManager.Contract.GetAbridgedVectorMetadata(&_MintManager.CallOpts, vectorId)
}

// GetClaimNoncesUsedForOffchainVector is a free data retrieval call binding the contract method 0xae709ae3.
//
// Solidity: function getClaimNoncesUsedForOffchainVector(bytes32 vectorId) view returns(bytes32[])
func (_MintManager *MintManagerCaller) GetClaimNoncesUsedForOffchainVector(opts *bind.CallOpts, vectorId [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "getClaimNoncesUsedForOffchainVector", vectorId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetClaimNoncesUsedForOffchainVector is a free data retrieval call binding the contract method 0xae709ae3.
//
// Solidity: function getClaimNoncesUsedForOffchainVector(bytes32 vectorId) view returns(bytes32[])
func (_MintManager *MintManagerSession) GetClaimNoncesUsedForOffchainVector(vectorId [32]byte) ([][32]byte, error) {
	return _MintManager.Contract.GetClaimNoncesUsedForOffchainVector(&_MintManager.CallOpts, vectorId)
}

// GetClaimNoncesUsedForOffchainVector is a free data retrieval call binding the contract method 0xae709ae3.
//
// Solidity: function getClaimNoncesUsedForOffchainVector(bytes32 vectorId) view returns(bytes32[])
func (_MintManager *MintManagerCallerSession) GetClaimNoncesUsedForOffchainVector(vectorId [32]byte) ([][32]byte, error) {
	return _MintManager.Contract.GetClaimNoncesUsedForOffchainVector(&_MintManager.CallOpts, vectorId)
}

// GetNumClaimedPerUserOffchainVector is a free data retrieval call binding the contract method 0x53274246.
//
// Solidity: function getNumClaimedPerUserOffchainVector(bytes32 vectorId, address user) view returns(uint256)
func (_MintManager *MintManagerCaller) GetNumClaimedPerUserOffchainVector(opts *bind.CallOpts, vectorId [32]byte, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "getNumClaimedPerUserOffchainVector", vectorId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumClaimedPerUserOffchainVector is a free data retrieval call binding the contract method 0x53274246.
//
// Solidity: function getNumClaimedPerUserOffchainVector(bytes32 vectorId, address user) view returns(uint256)
func (_MintManager *MintManagerSession) GetNumClaimedPerUserOffchainVector(vectorId [32]byte, user common.Address) (*big.Int, error) {
	return _MintManager.Contract.GetNumClaimedPerUserOffchainVector(&_MintManager.CallOpts, vectorId, user)
}

// GetNumClaimedPerUserOffchainVector is a free data retrieval call binding the contract method 0x53274246.
//
// Solidity: function getNumClaimedPerUserOffchainVector(bytes32 vectorId, address user) view returns(uint256)
func (_MintManager *MintManagerCallerSession) GetNumClaimedPerUserOffchainVector(vectorId [32]byte, user common.Address) (*big.Int, error) {
	return _MintManager.Contract.GetNumClaimedPerUserOffchainVector(&_MintManager.CallOpts, vectorId, user)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3716e284.
//
// Solidity: function isNonceUsed(bytes32 vectorId, bytes32 nonce) view returns(bool)
func (_MintManager *MintManagerCaller) IsNonceUsed(opts *bind.CallOpts, vectorId [32]byte, nonce [32]byte) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "isNonceUsed", vectorId, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3716e284.
//
// Solidity: function isNonceUsed(bytes32 vectorId, bytes32 nonce) view returns(bool)
func (_MintManager *MintManagerSession) IsNonceUsed(vectorId [32]byte, nonce [32]byte) (bool, error) {
	return _MintManager.Contract.IsNonceUsed(&_MintManager.CallOpts, vectorId, nonce)
}

// IsNonceUsed is a free data retrieval call binding the contract method 0x3716e284.
//
// Solidity: function isNonceUsed(bytes32 vectorId, bytes32 nonce) view returns(bool)
func (_MintManager *MintManagerCallerSession) IsNonceUsed(vectorId [32]byte, nonce [32]byte) (bool, error) {
	return _MintManager.Contract.IsNonceUsed(&_MintManager.CallOpts, vectorId, nonce)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintManager *MintManagerCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintManager *MintManagerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MintManager.Contract.IsTrustedForwarder(&_MintManager.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_MintManager *MintManagerCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _MintManager.Contract.IsTrustedForwarder(&_MintManager.CallOpts, forwarder)
}

// OffchainVectorsClaimState is a free data retrieval call binding the contract method 0x0dbb18a1.
//
// Solidity: function offchainVectorsClaimState(bytes32 ) view returns(uint256 numClaimed)
func (_MintManager *MintManagerCaller) OffchainVectorsClaimState(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "offchainVectorsClaimState", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffchainVectorsClaimState is a free data retrieval call binding the contract method 0x0dbb18a1.
//
// Solidity: function offchainVectorsClaimState(bytes32 ) view returns(uint256 numClaimed)
func (_MintManager *MintManagerSession) OffchainVectorsClaimState(arg0 [32]byte) (*big.Int, error) {
	return _MintManager.Contract.OffchainVectorsClaimState(&_MintManager.CallOpts, arg0)
}

// OffchainVectorsClaimState is a free data retrieval call binding the contract method 0x0dbb18a1.
//
// Solidity: function offchainVectorsClaimState(bytes32 ) view returns(uint256 numClaimed)
func (_MintManager *MintManagerCallerSession) OffchainVectorsClaimState(arg0 [32]byte) (*big.Int, error) {
	return _MintManager.Contract.OffchainVectorsClaimState(&_MintManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MintManager *MintManagerCallerSession) Owner() (common.Address, error) {
	return _MintManager.Contract.Owner(&_MintManager.CallOpts)
}

// PlatformExecutors is a free data retrieval call binding the contract method 0x9c5c0492.
//
// Solidity: function platformExecutors() view returns(address[])
func (_MintManager *MintManagerCaller) PlatformExecutors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "platformExecutors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// PlatformExecutors is a free data retrieval call binding the contract method 0x9c5c0492.
//
// Solidity: function platformExecutors() view returns(address[])
func (_MintManager *MintManagerSession) PlatformExecutors() ([]common.Address, error) {
	return _MintManager.Contract.PlatformExecutors(&_MintManager.CallOpts)
}

// PlatformExecutors is a free data retrieval call binding the contract method 0x9c5c0492.
//
// Solidity: function platformExecutors() view returns(address[])
func (_MintManager *MintManagerCallerSession) PlatformExecutors() ([]common.Address, error) {
	return _MintManager.Contract.PlatformExecutors(&_MintManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintManager *MintManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintManager *MintManagerSession) ProxiableUUID() ([32]byte, error) {
	return _MintManager.Contract.ProxiableUUID(&_MintManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_MintManager *MintManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _MintManager.Contract.ProxiableUUID(&_MintManager.CallOpts)
}

// UserClaims is a free data retrieval call binding the contract method 0x9e2dc500.
//
// Solidity: function userClaims(uint256 , address ) view returns(uint64)
func (_MintManager *MintManagerCaller) UserClaims(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "userClaims", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// UserClaims is a free data retrieval call binding the contract method 0x9e2dc500.
//
// Solidity: function userClaims(uint256 , address ) view returns(uint64)
func (_MintManager *MintManagerSession) UserClaims(arg0 *big.Int, arg1 common.Address) (uint64, error) {
	return _MintManager.Contract.UserClaims(&_MintManager.CallOpts, arg0, arg1)
}

// UserClaims is a free data retrieval call binding the contract method 0x9e2dc500.
//
// Solidity: function userClaims(uint256 , address ) view returns(uint64)
func (_MintManager *MintManagerCallerSession) UserClaims(arg0 *big.Int, arg1 common.Address) (uint64, error) {
	return _MintManager.Contract.UserClaims(&_MintManager.CallOpts, arg0, arg1)
}

// VectorMutabilities is a free data retrieval call binding the contract method 0xc462507e.
//
// Solidity: function vectorMutabilities(uint256 ) view returns(uint8 updatesFrozen, uint8 deleteFrozen, uint8 pausesFrozen)
func (_MintManager *MintManagerCaller) VectorMutabilities(opts *bind.CallOpts, arg0 *big.Int) (struct {
	UpdatesFrozen uint8
	DeleteFrozen  uint8
	PausesFrozen  uint8
}, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "vectorMutabilities", arg0)

	outstruct := new(struct {
		UpdatesFrozen uint8
		DeleteFrozen  uint8
		PausesFrozen  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UpdatesFrozen = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.DeleteFrozen = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.PausesFrozen = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// VectorMutabilities is a free data retrieval call binding the contract method 0xc462507e.
//
// Solidity: function vectorMutabilities(uint256 ) view returns(uint8 updatesFrozen, uint8 deleteFrozen, uint8 pausesFrozen)
func (_MintManager *MintManagerSession) VectorMutabilities(arg0 *big.Int) (struct {
	UpdatesFrozen uint8
	DeleteFrozen  uint8
	PausesFrozen  uint8
}, error) {
	return _MintManager.Contract.VectorMutabilities(&_MintManager.CallOpts, arg0)
}

// VectorMutabilities is a free data retrieval call binding the contract method 0xc462507e.
//
// Solidity: function vectorMutabilities(uint256 ) view returns(uint8 updatesFrozen, uint8 deleteFrozen, uint8 pausesFrozen)
func (_MintManager *MintManagerCallerSession) VectorMutabilities(arg0 *big.Int) (struct {
	UpdatesFrozen uint8
	DeleteFrozen  uint8
	PausesFrozen  uint8
}, error) {
	return _MintManager.Contract.VectorMutabilities(&_MintManager.CallOpts, arg0)
}

// VectorToEditionId is a free data retrieval call binding the contract method 0xb414ae2f.
//
// Solidity: function vectorToEditionId(uint256 ) view returns(uint256)
func (_MintManager *MintManagerCaller) VectorToEditionId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "vectorToEditionId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VectorToEditionId is a free data retrieval call binding the contract method 0xb414ae2f.
//
// Solidity: function vectorToEditionId(uint256 ) view returns(uint256)
func (_MintManager *MintManagerSession) VectorToEditionId(arg0 *big.Int) (*big.Int, error) {
	return _MintManager.Contract.VectorToEditionId(&_MintManager.CallOpts, arg0)
}

// VectorToEditionId is a free data retrieval call binding the contract method 0xb414ae2f.
//
// Solidity: function vectorToEditionId(uint256 ) view returns(uint256)
func (_MintManager *MintManagerCallerSession) VectorToEditionId(arg0 *big.Int) (*big.Int, error) {
	return _MintManager.Contract.VectorToEditionId(&_MintManager.CallOpts, arg0)
}

// Vectors is a free data retrieval call binding the contract method 0x619b8589.
//
// Solidity: function vectors(uint256 ) view returns(address contractAddress, address currency, address paymentRecipient, uint256 startTimestamp, uint256 endTimestamp, uint256 pricePerToken, uint64 tokenLimitPerTx, uint64 maxTotalClaimableViaVector, uint64 maxUserClaimableViaVector, uint64 totalClaimedViaVector, bytes32 allowlistRoot, uint8 paused)
func (_MintManager *MintManagerCaller) Vectors(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ContractAddress            common.Address
	Currency                   common.Address
	PaymentRecipient           common.Address
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	PricePerToken              *big.Int
	TokenLimitPerTx            uint64
	MaxTotalClaimableViaVector uint64
	MaxUserClaimableViaVector  uint64
	TotalClaimedViaVector      uint64
	AllowlistRoot              [32]byte
	Paused                     uint8
}, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "vectors", arg0)

	outstruct := new(struct {
		ContractAddress            common.Address
		Currency                   common.Address
		PaymentRecipient           common.Address
		StartTimestamp             *big.Int
		EndTimestamp               *big.Int
		PricePerToken              *big.Int
		TokenLimitPerTx            uint64
		MaxTotalClaimableViaVector uint64
		MaxUserClaimableViaVector  uint64
		TotalClaimedViaVector      uint64
		AllowlistRoot              [32]byte
		Paused                     uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ContractAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Currency = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PaymentRecipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.StartTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EndTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PricePerToken = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TokenLimitPerTx = *abi.ConvertType(out[6], new(uint64)).(*uint64)
	outstruct.MaxTotalClaimableViaVector = *abi.ConvertType(out[7], new(uint64)).(*uint64)
	outstruct.MaxUserClaimableViaVector = *abi.ConvertType(out[8], new(uint64)).(*uint64)
	outstruct.TotalClaimedViaVector = *abi.ConvertType(out[9], new(uint64)).(*uint64)
	outstruct.AllowlistRoot = *abi.ConvertType(out[10], new([32]byte)).(*[32]byte)
	outstruct.Paused = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// Vectors is a free data retrieval call binding the contract method 0x619b8589.
//
// Solidity: function vectors(uint256 ) view returns(address contractAddress, address currency, address paymentRecipient, uint256 startTimestamp, uint256 endTimestamp, uint256 pricePerToken, uint64 tokenLimitPerTx, uint64 maxTotalClaimableViaVector, uint64 maxUserClaimableViaVector, uint64 totalClaimedViaVector, bytes32 allowlistRoot, uint8 paused)
func (_MintManager *MintManagerSession) Vectors(arg0 *big.Int) (struct {
	ContractAddress            common.Address
	Currency                   common.Address
	PaymentRecipient           common.Address
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	PricePerToken              *big.Int
	TokenLimitPerTx            uint64
	MaxTotalClaimableViaVector uint64
	MaxUserClaimableViaVector  uint64
	TotalClaimedViaVector      uint64
	AllowlistRoot              [32]byte
	Paused                     uint8
}, error) {
	return _MintManager.Contract.Vectors(&_MintManager.CallOpts, arg0)
}

// Vectors is a free data retrieval call binding the contract method 0x619b8589.
//
// Solidity: function vectors(uint256 ) view returns(address contractAddress, address currency, address paymentRecipient, uint256 startTimestamp, uint256 endTimestamp, uint256 pricePerToken, uint64 tokenLimitPerTx, uint64 maxTotalClaimableViaVector, uint64 maxUserClaimableViaVector, uint64 totalClaimedViaVector, bytes32 allowlistRoot, uint8 paused)
func (_MintManager *MintManagerCallerSession) Vectors(arg0 *big.Int) (struct {
	ContractAddress            common.Address
	Currency                   common.Address
	PaymentRecipient           common.Address
	StartTimestamp             *big.Int
	EndTimestamp               *big.Int
	PricePerToken              *big.Int
	TokenLimitPerTx            uint64
	MaxTotalClaimableViaVector uint64
	MaxUserClaimableViaVector  uint64
	TotalClaimedViaVector      uint64
	AllowlistRoot              [32]byte
	Paused                     uint8
}, error) {
	return _MintManager.Contract.Vectors(&_MintManager.CallOpts, arg0)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xf4a40345.
//
// Solidity: function verifyClaim((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerCaller) VerifyClaim(opts *bind.CallOpts, claim MintManagerClaim, signature []byte, expectedMsgSender common.Address) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "verifyClaim", claim, signature, expectedMsgSender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0xf4a40345.
//
// Solidity: function verifyClaim((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerSession) VerifyClaim(claim MintManagerClaim, signature []byte, expectedMsgSender common.Address) (bool, error) {
	return _MintManager.Contract.VerifyClaim(&_MintManager.CallOpts, claim, signature, expectedMsgSender)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xf4a40345.
//
// Solidity: function verifyClaim((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerCallerSession) VerifyClaim(claim MintManagerClaim, signature []byte, expectedMsgSender common.Address) (bool, error) {
	return _MintManager.Contract.VerifyClaim(&_MintManager.CallOpts, claim, signature, expectedMsgSender)
}

// VerifyClaimWithMetaTxPacket is a free data retrieval call binding the contract method 0x5dcd547e.
//
// Solidity: function verifyClaimWithMetaTxPacket((address,address,address,uint256,uint64,(bytes,bytes32,bytes32,uint8),(bytes,bytes32,bytes32,uint8),uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerCaller) VerifyClaimWithMetaTxPacket(opts *bind.CallOpts, claim MintManagerClaimWithMetaTxPacket, signature []byte, expectedMsgSender common.Address) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "verifyClaimWithMetaTxPacket", claim, signature, expectedMsgSender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaimWithMetaTxPacket is a free data retrieval call binding the contract method 0x5dcd547e.
//
// Solidity: function verifyClaimWithMetaTxPacket((address,address,address,uint256,uint64,(bytes,bytes32,bytes32,uint8),(bytes,bytes32,bytes32,uint8),uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerSession) VerifyClaimWithMetaTxPacket(claim MintManagerClaimWithMetaTxPacket, signature []byte, expectedMsgSender common.Address) (bool, error) {
	return _MintManager.Contract.VerifyClaimWithMetaTxPacket(&_MintManager.CallOpts, claim, signature, expectedMsgSender)
}

// VerifyClaimWithMetaTxPacket is a free data retrieval call binding the contract method 0x5dcd547e.
//
// Solidity: function verifyClaimWithMetaTxPacket((address,address,address,uint256,uint64,(bytes,bytes32,bytes32,uint8),(bytes,bytes32,bytes32,uint8),uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes signature, address expectedMsgSender) view returns(bool)
func (_MintManager *MintManagerCallerSession) VerifyClaimWithMetaTxPacket(claim MintManagerClaimWithMetaTxPacket, signature []byte, expectedMsgSender common.Address) (bool, error) {
	return _MintManager.Contract.VerifyClaimWithMetaTxPacket(&_MintManager.CallOpts, claim, signature, expectedMsgSender)
}

// VerifySeriesClaim is a free data retrieval call binding the contract method 0xfc425875.
//
// Solidity: function verifySeriesClaim((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes signature, address expectedMsgSender, uint256[] tokenIds) view returns(bool)
func (_MintManager *MintManagerCaller) VerifySeriesClaim(opts *bind.CallOpts, claim MintManagerSeriesClaim, signature []byte, expectedMsgSender common.Address, tokenIds []*big.Int) (bool, error) {
	var out []interface{}
	err := _MintManager.contract.Call(opts, &out, "verifySeriesClaim", claim, signature, expectedMsgSender, tokenIds)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySeriesClaim is a free data retrieval call binding the contract method 0xfc425875.
//
// Solidity: function verifySeriesClaim((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes signature, address expectedMsgSender, uint256[] tokenIds) view returns(bool)
func (_MintManager *MintManagerSession) VerifySeriesClaim(claim MintManagerSeriesClaim, signature []byte, expectedMsgSender common.Address, tokenIds []*big.Int) (bool, error) {
	return _MintManager.Contract.VerifySeriesClaim(&_MintManager.CallOpts, claim, signature, expectedMsgSender, tokenIds)
}

// VerifySeriesClaim is a free data retrieval call binding the contract method 0xfc425875.
//
// Solidity: function verifySeriesClaim((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes signature, address expectedMsgSender, uint256[] tokenIds) view returns(bool)
func (_MintManager *MintManagerCallerSession) VerifySeriesClaim(claim MintManagerSeriesClaim, signature []byte, expectedMsgSender common.Address, tokenIds []*big.Int) (bool, error) {
	return _MintManager.Contract.VerifySeriesClaim(&_MintManager.CallOpts, claim, signature, expectedMsgSender, tokenIds)
}

// AddPlatformExecutor is a paid mutator transaction binding the contract method 0xf73bc2da.
//
// Solidity: function addPlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerTransactor) AddPlatformExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "addPlatformExecutor", _executor)
}

// AddPlatformExecutor is a paid mutator transaction binding the contract method 0xf73bc2da.
//
// Solidity: function addPlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerSession) AddPlatformExecutor(_executor common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.AddPlatformExecutor(&_MintManager.TransactOpts, _executor)
}

// AddPlatformExecutor is a paid mutator transaction binding the contract method 0xf73bc2da.
//
// Solidity: function addPlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerTransactorSession) AddPlatformExecutor(_executor common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.AddPlatformExecutor(&_MintManager.TransactOpts, _executor)
}

// CreateAbridgedVector is a paid mutator transaction binding the contract method 0x77a856ea.
//
// Solidity: function createAbridgedVector((uint160,uint48,uint48,uint160,uint48,uint48,uint160,uint48,uint48,uint192,uint48,bool,bool,bytes32) _vector) returns()
func (_MintManager *MintManagerTransactor) CreateAbridgedVector(opts *bind.TransactOpts, _vector IAbridgedMintVectorAbridgedVectorData) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "createAbridgedVector", _vector)
}

// CreateAbridgedVector is a paid mutator transaction binding the contract method 0x77a856ea.
//
// Solidity: function createAbridgedVector((uint160,uint48,uint48,uint160,uint48,uint48,uint160,uint48,uint48,uint192,uint48,bool,bool,bytes32) _vector) returns()
func (_MintManager *MintManagerSession) CreateAbridgedVector(_vector IAbridgedMintVectorAbridgedVectorData) (*types.Transaction, error) {
	return _MintManager.Contract.CreateAbridgedVector(&_MintManager.TransactOpts, _vector)
}

// CreateAbridgedVector is a paid mutator transaction binding the contract method 0x77a856ea.
//
// Solidity: function createAbridgedVector((uint160,uint48,uint48,uint160,uint48,uint48,uint160,uint48,uint48,uint192,uint48,bool,bool,bytes32) _vector) returns()
func (_MintManager *MintManagerTransactorSession) CreateAbridgedVector(_vector IAbridgedMintVectorAbridgedVectorData) (*types.Transaction, error) {
	return _MintManager.Contract.CreateAbridgedVector(&_MintManager.TransactOpts, _vector)
}

// DeleteAbridgedVector is a paid mutator transaction binding the contract method 0x4d66747b.
//
// Solidity: function deleteAbridgedVector(uint256 vectorId) returns()
func (_MintManager *MintManagerTransactor) DeleteAbridgedVector(opts *bind.TransactOpts, vectorId *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "deleteAbridgedVector", vectorId)
}

// DeleteAbridgedVector is a paid mutator transaction binding the contract method 0x4d66747b.
//
// Solidity: function deleteAbridgedVector(uint256 vectorId) returns()
func (_MintManager *MintManagerSession) DeleteAbridgedVector(vectorId *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.DeleteAbridgedVector(&_MintManager.TransactOpts, vectorId)
}

// DeleteAbridgedVector is a paid mutator transaction binding the contract method 0x4d66747b.
//
// Solidity: function deleteAbridgedVector(uint256 vectorId) returns()
func (_MintManager *MintManagerTransactorSession) DeleteAbridgedVector(vectorId *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.DeleteAbridgedVector(&_MintManager.TransactOpts, vectorId)
}

// DeprecatePlatformExecutor is a paid mutator transaction binding the contract method 0x46b060c3.
//
// Solidity: function deprecatePlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerTransactor) DeprecatePlatformExecutor(opts *bind.TransactOpts, _executor common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "deprecatePlatformExecutor", _executor)
}

// DeprecatePlatformExecutor is a paid mutator transaction binding the contract method 0x46b060c3.
//
// Solidity: function deprecatePlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerSession) DeprecatePlatformExecutor(_executor common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.DeprecatePlatformExecutor(&_MintManager.TransactOpts, _executor)
}

// DeprecatePlatformExecutor is a paid mutator transaction binding the contract method 0x46b060c3.
//
// Solidity: function deprecatePlatformExecutor(address _executor) returns()
func (_MintManager *MintManagerTransactorSession) DeprecatePlatformExecutor(_executor common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.DeprecatePlatformExecutor(&_MintManager.TransactOpts, _executor)
}

// GatedMintEdition721 is a paid mutator transaction binding the contract method 0x0680475b.
//
// Solidity: function gatedMintEdition721((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) _claim, bytes _signature, address _recipient) payable returns()
func (_MintManager *MintManagerTransactor) GatedMintEdition721(opts *bind.TransactOpts, _claim MintManagerClaim, _signature []byte, _recipient common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "gatedMintEdition721", _claim, _signature, _recipient)
}

// GatedMintEdition721 is a paid mutator transaction binding the contract method 0x0680475b.
//
// Solidity: function gatedMintEdition721((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) _claim, bytes _signature, address _recipient) payable returns()
func (_MintManager *MintManagerSession) GatedMintEdition721(_claim MintManagerClaim, _signature []byte, _recipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.GatedMintEdition721(&_MintManager.TransactOpts, _claim, _signature, _recipient)
}

// GatedMintEdition721 is a paid mutator transaction binding the contract method 0x0680475b.
//
// Solidity: function gatedMintEdition721((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) _claim, bytes _signature, address _recipient) payable returns()
func (_MintManager *MintManagerTransactorSession) GatedMintEdition721(_claim MintManagerClaim, _signature []byte, _recipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.GatedMintEdition721(&_MintManager.TransactOpts, _claim, _signature, _recipient)
}

// GatedSeriesMint is a paid mutator transaction binding the contract method 0x99bcbf21.
//
// Solidity: function gatedSeriesMint((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactor) GatedSeriesMint(opts *bind.TransactOpts, claim MintManagerClaim, claimSignature []byte, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "gatedSeriesMint", claim, claimSignature, mintRecipient)
}

// GatedSeriesMint is a paid mutator transaction binding the contract method 0x99bcbf21.
//
// Solidity: function gatedSeriesMint((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient) payable returns()
func (_MintManager *MintManagerSession) GatedSeriesMint(claim MintManagerClaim, claimSignature []byte, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.GatedSeriesMint(&_MintManager.TransactOpts, claim, claimSignature, mintRecipient)
}

// GatedSeriesMint is a paid mutator transaction binding the contract method 0x99bcbf21.
//
// Solidity: function gatedSeriesMint((address,address,address,address,uint256,uint64,uint256,uint256,uint256,uint256,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactorSession) GatedSeriesMint(claim MintManagerClaim, claimSignature []byte, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.GatedSeriesMint(&_MintManager.TransactOpts, claim, claimSignature, mintRecipient)
}

// GatedSeriesMintChooseToken is a paid mutator transaction binding the contract method 0x6c1b7abd.
//
// Solidity: function gatedSeriesMintChooseToken((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient, uint256[] tokenIds) payable returns()
func (_MintManager *MintManagerTransactor) GatedSeriesMintChooseToken(opts *bind.TransactOpts, claim MintManagerSeriesClaim, claimSignature []byte, mintRecipient common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "gatedSeriesMintChooseToken", claim, claimSignature, mintRecipient, tokenIds)
}

// GatedSeriesMintChooseToken is a paid mutator transaction binding the contract method 0x6c1b7abd.
//
// Solidity: function gatedSeriesMintChooseToken((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient, uint256[] tokenIds) payable returns()
func (_MintManager *MintManagerSession) GatedSeriesMintChooseToken(claim MintManagerSeriesClaim, claimSignature []byte, mintRecipient common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.GatedSeriesMintChooseToken(&_MintManager.TransactOpts, claim, claimSignature, mintRecipient, tokenIds)
}

// GatedSeriesMintChooseToken is a paid mutator transaction binding the contract method 0x6c1b7abd.
//
// Solidity: function gatedSeriesMintChooseToken((address,address,address,address,uint256,uint64,uint64,uint64,uint64,bytes32,bytes32) claim, bytes claimSignature, address mintRecipient, uint256[] tokenIds) payable returns()
func (_MintManager *MintManagerTransactorSession) GatedSeriesMintChooseToken(claim MintManagerSeriesClaim, claimSignature []byte, mintRecipient common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.GatedSeriesMintChooseToken(&_MintManager.TransactOpts, claim, claimSignature, mintRecipient, tokenIds)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address platform, address _owner, address trustedForwarder, address initialExecutor, uint256 initialPlatformMintFee) returns()
func (_MintManager *MintManagerTransactor) Initialize(opts *bind.TransactOpts, platform common.Address, _owner common.Address, trustedForwarder common.Address, initialExecutor common.Address, initialPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "initialize", platform, _owner, trustedForwarder, initialExecutor, initialPlatformMintFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address platform, address _owner, address trustedForwarder, address initialExecutor, uint256 initialPlatformMintFee) returns()
func (_MintManager *MintManagerSession) Initialize(platform common.Address, _owner common.Address, trustedForwarder common.Address, initialExecutor common.Address, initialPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Initialize(&_MintManager.TransactOpts, platform, _owner, trustedForwarder, initialExecutor, initialPlatformMintFee)
}

// Initialize is a paid mutator transaction binding the contract method 0xf7013ef6.
//
// Solidity: function initialize(address platform, address _owner, address trustedForwarder, address initialExecutor, uint256 initialPlatformMintFee) returns()
func (_MintManager *MintManagerTransactorSession) Initialize(platform common.Address, _owner common.Address, trustedForwarder common.Address, initialExecutor common.Address, initialPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.Initialize(&_MintManager.TransactOpts, platform, _owner, trustedForwarder, initialExecutor, initialPlatformMintFee)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MintManager *MintManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MintManager.Contract.RenounceOwnership(&_MintManager.TransactOpts)
}

// SetAbridgedVectorMetadata is a paid mutator transaction binding the contract method 0x391a6d28.
//
// Solidity: function setAbridgedVectorMetadata(uint256 vectorId, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerTransactor) SetAbridgedVectorMetadata(opts *bind.TransactOpts, vectorId *big.Int, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "setAbridgedVectorMetadata", vectorId, pause, flexibleData)
}

// SetAbridgedVectorMetadata is a paid mutator transaction binding the contract method 0x391a6d28.
//
// Solidity: function setAbridgedVectorMetadata(uint256 vectorId, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerSession) SetAbridgedVectorMetadata(vectorId *big.Int, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.SetAbridgedVectorMetadata(&_MintManager.TransactOpts, vectorId, pause, flexibleData)
}

// SetAbridgedVectorMetadata is a paid mutator transaction binding the contract method 0x391a6d28.
//
// Solidity: function setAbridgedVectorMetadata(uint256 vectorId, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerTransactorSession) SetAbridgedVectorMetadata(vectorId *big.Int, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.SetAbridgedVectorMetadata(&_MintManager.TransactOpts, vectorId, pause, flexibleData)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MintManager *MintManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.TransferOwnership(&_MintManager.TransactOpts, newOwner)
}

// UpdateAbridgedVector is a paid mutator transaction binding the contract method 0x2cf9adc8.
//
// Solidity: function updateAbridgedVector(uint256 vectorId, (address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32) _newVector, (uint16,uint16,uint16,uint16,uint16,uint16,uint8,uint8,uint8,uint8) updateConfig, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerTransactor) UpdateAbridgedVector(opts *bind.TransactOpts, vectorId *big.Int, _newVector IAbridgedMintVectorAbridgedVector, updateConfig IAbridgedMintVectorUpdateAbridgedVectorConfig, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "updateAbridgedVector", vectorId, _newVector, updateConfig, pause, flexibleData)
}

// UpdateAbridgedVector is a paid mutator transaction binding the contract method 0x2cf9adc8.
//
// Solidity: function updateAbridgedVector(uint256 vectorId, (address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32) _newVector, (uint16,uint16,uint16,uint16,uint16,uint16,uint8,uint8,uint8,uint8) updateConfig, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerSession) UpdateAbridgedVector(vectorId *big.Int, _newVector IAbridgedMintVectorAbridgedVector, updateConfig IAbridgedMintVectorUpdateAbridgedVectorConfig, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.UpdateAbridgedVector(&_MintManager.TransactOpts, vectorId, _newVector, updateConfig, pause, flexibleData)
}

// UpdateAbridgedVector is a paid mutator transaction binding the contract method 0x2cf9adc8.
//
// Solidity: function updateAbridgedVector(uint256 vectorId, (address,uint48,uint48,address,uint48,uint48,address,uint48,uint48,uint192,uint48,bool,bool,bytes32) _newVector, (uint16,uint16,uint16,uint16,uint16,uint16,uint8,uint8,uint8,uint8) updateConfig, bool pause, uint128 flexibleData) returns()
func (_MintManager *MintManagerTransactorSession) UpdateAbridgedVector(vectorId *big.Int, _newVector IAbridgedMintVectorAbridgedVector, updateConfig IAbridgedMintVectorUpdateAbridgedVectorConfig, pause bool, flexibleData *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.UpdateAbridgedVector(&_MintManager.TransactOpts, vectorId, _newVector, updateConfig, pause, flexibleData)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 newPlatformMintFee) returns()
func (_MintManager *MintManagerTransactor) UpdatePlatformMintFee(opts *bind.TransactOpts, newPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "updatePlatformMintFee", newPlatformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 newPlatformMintFee) returns()
func (_MintManager *MintManagerSession) UpdatePlatformMintFee(newPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.UpdatePlatformMintFee(&_MintManager.TransactOpts, newPlatformMintFee)
}

// UpdatePlatformMintFee is a paid mutator transaction binding the contract method 0x61d2beae.
//
// Solidity: function updatePlatformMintFee(uint256 newPlatformMintFee) returns()
func (_MintManager *MintManagerTransactorSession) UpdatePlatformMintFee(newPlatformMintFee *big.Int) (*types.Transaction, error) {
	return _MintManager.Contract.UpdatePlatformMintFee(&_MintManager.TransactOpts, newPlatformMintFee)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MintManager *MintManagerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MintManager *MintManagerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.UpgradeTo(&_MintManager.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_MintManager *MintManagerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.UpgradeTo(&_MintManager.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintManager *MintManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintManager *MintManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintManager.Contract.UpgradeToAndCall(&_MintManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_MintManager *MintManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _MintManager.Contract.UpgradeToAndCall(&_MintManager.TransactOpts, newImplementation, data)
}

// VectorMintEdition721 is a paid mutator transaction binding the contract method 0x5bbd1326.
//
// Solidity: function vectorMintEdition721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactor) VectorMintEdition721(opts *bind.TransactOpts, vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "vectorMintEdition721", vectorId, numTokensToMint, mintRecipient)
}

// VectorMintEdition721 is a paid mutator transaction binding the contract method 0x5bbd1326.
//
// Solidity: function vectorMintEdition721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerSession) VectorMintEdition721(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintEdition721(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient)
}

// VectorMintEdition721 is a paid mutator transaction binding the contract method 0x5bbd1326.
//
// Solidity: function vectorMintEdition721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactorSession) VectorMintEdition721(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintEdition721(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient)
}

// VectorMintEdition721WithAllowlist is a paid mutator transaction binding the contract method 0x1ebbed86.
//
// Solidity: function vectorMintEdition721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerTransactor) VectorMintEdition721WithAllowlist(opts *bind.TransactOpts, vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "vectorMintEdition721WithAllowlist", vectorId, numTokensToMint, mintRecipient, proof)
}

// VectorMintEdition721WithAllowlist is a paid mutator transaction binding the contract method 0x1ebbed86.
//
// Solidity: function vectorMintEdition721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerSession) VectorMintEdition721WithAllowlist(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintEdition721WithAllowlist(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient, proof)
}

// VectorMintEdition721WithAllowlist is a paid mutator transaction binding the contract method 0x1ebbed86.
//
// Solidity: function vectorMintEdition721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerTransactorSession) VectorMintEdition721WithAllowlist(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintEdition721WithAllowlist(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient, proof)
}

// VectorMintSeries721 is a paid mutator transaction binding the contract method 0x705dcf34.
//
// Solidity: function vectorMintSeries721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactor) VectorMintSeries721(opts *bind.TransactOpts, vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "vectorMintSeries721", vectorId, numTokensToMint, mintRecipient)
}

// VectorMintSeries721 is a paid mutator transaction binding the contract method 0x705dcf34.
//
// Solidity: function vectorMintSeries721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerSession) VectorMintSeries721(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintSeries721(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient)
}

// VectorMintSeries721 is a paid mutator transaction binding the contract method 0x705dcf34.
//
// Solidity: function vectorMintSeries721(uint256 vectorId, uint48 numTokensToMint, address mintRecipient) payable returns()
func (_MintManager *MintManagerTransactorSession) VectorMintSeries721(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintSeries721(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient)
}

// VectorMintSeries721WithAllowlist is a paid mutator transaction binding the contract method 0x3497a0d2.
//
// Solidity: function vectorMintSeries721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerTransactor) VectorMintSeries721WithAllowlist(opts *bind.TransactOpts, vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "vectorMintSeries721WithAllowlist", vectorId, numTokensToMint, mintRecipient, proof)
}

// VectorMintSeries721WithAllowlist is a paid mutator transaction binding the contract method 0x3497a0d2.
//
// Solidity: function vectorMintSeries721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerSession) VectorMintSeries721WithAllowlist(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintSeries721WithAllowlist(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient, proof)
}

// VectorMintSeries721WithAllowlist is a paid mutator transaction binding the contract method 0x3497a0d2.
//
// Solidity: function vectorMintSeries721WithAllowlist(uint256 vectorId, uint48 numTokensToMint, address mintRecipient, bytes32[] proof) payable returns()
func (_MintManager *MintManagerTransactorSession) VectorMintSeries721WithAllowlist(vectorId *big.Int, numTokensToMint *big.Int, mintRecipient common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _MintManager.Contract.VectorMintSeries721WithAllowlist(&_MintManager.TransactOpts, vectorId, numTokensToMint, mintRecipient, proof)
}

// WithdrawNativeGasToken is a paid mutator transaction binding the contract method 0x513ea090.
//
// Solidity: function withdrawNativeGasToken() returns()
func (_MintManager *MintManagerTransactor) WithdrawNativeGasToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintManager.contract.Transact(opts, "withdrawNativeGasToken")
}

// WithdrawNativeGasToken is a paid mutator transaction binding the contract method 0x513ea090.
//
// Solidity: function withdrawNativeGasToken() returns()
func (_MintManager *MintManagerSession) WithdrawNativeGasToken() (*types.Transaction, error) {
	return _MintManager.Contract.WithdrawNativeGasToken(&_MintManager.TransactOpts)
}

// WithdrawNativeGasToken is a paid mutator transaction binding the contract method 0x513ea090.
//
// Solidity: function withdrawNativeGasToken() returns()
func (_MintManager *MintManagerTransactorSession) WithdrawNativeGasToken() (*types.Transaction, error) {
	return _MintManager.Contract.WithdrawNativeGasToken(&_MintManager.TransactOpts)
}

// MintManagerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the MintManager contract.
type MintManagerAdminChangedIterator struct {
	Event *MintManagerAdminChanged // Event containing the contract specifics and raw log

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
func (it *MintManagerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerAdminChanged)
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
		it.Event = new(MintManagerAdminChanged)
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
func (it *MintManagerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerAdminChanged represents a AdminChanged event raised by the MintManager contract.
type MintManagerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MintManager *MintManagerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*MintManagerAdminChangedIterator, error) {

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &MintManagerAdminChangedIterator{contract: _MintManager.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MintManager *MintManagerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *MintManagerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerAdminChanged)
				if err := _MintManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_MintManager *MintManagerFilterer) ParseAdminChanged(log types.Log) (*MintManagerAdminChanged, error) {
	event := new(MintManagerAdminChanged)
	if err := _MintManager.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the MintManager contract.
type MintManagerBeaconUpgradedIterator struct {
	Event *MintManagerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *MintManagerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerBeaconUpgraded)
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
		it.Event = new(MintManagerBeaconUpgraded)
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
func (it *MintManagerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerBeaconUpgraded represents a BeaconUpgraded event raised by the MintManager contract.
type MintManagerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MintManager *MintManagerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*MintManagerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerBeaconUpgradedIterator{contract: _MintManager.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MintManager *MintManagerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *MintManagerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerBeaconUpgraded)
				if err := _MintManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_MintManager *MintManagerFilterer) ParseBeaconUpgraded(log types.Log) (*MintManagerBeaconUpgraded, error) {
	event := new(MintManagerBeaconUpgraded)
	if err := _MintManager.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerChooseTokenMintIterator is returned from FilterChooseTokenMint and is used to iterate over the raw logs and unpacked data for ChooseTokenMint events raised by the MintManager contract.
type MintManagerChooseTokenMintIterator struct {
	Event *MintManagerChooseTokenMint // Event containing the contract specifics and raw log

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
func (it *MintManagerChooseTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerChooseTokenMint)
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
		it.Event = new(MintManagerChooseTokenMint)
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
func (it *MintManagerChooseTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerChooseTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerChooseTokenMint represents a ChooseTokenMint event raised by the MintManager contract.
type MintManagerChooseTokenMint struct {
	VectorId        [32]byte
	ContractAddress common.Address
	OnChainVector   bool
	TokenIds        []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterChooseTokenMint is a free log retrieval operation binding the contract event 0xe0bf8a4af82e2af496af5f4957e2767f8b52e51a77caedd2f30a1843872d1b7c.
//
// Solidity: event ChooseTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256[] tokenIds)
func (_MintManager *MintManagerFilterer) FilterChooseTokenMint(opts *bind.FilterOpts, vectorId [][32]byte, contractAddress []common.Address, onChainVector []bool) (*MintManagerChooseTokenMintIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var onChainVectorRule []interface{}
	for _, onChainVectorItem := range onChainVector {
		onChainVectorRule = append(onChainVectorRule, onChainVectorItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "ChooseTokenMint", vectorIdRule, contractAddressRule, onChainVectorRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerChooseTokenMintIterator{contract: _MintManager.contract, event: "ChooseTokenMint", logs: logs, sub: sub}, nil
}

// WatchChooseTokenMint is a free log subscription operation binding the contract event 0xe0bf8a4af82e2af496af5f4957e2767f8b52e51a77caedd2f30a1843872d1b7c.
//
// Solidity: event ChooseTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256[] tokenIds)
func (_MintManager *MintManagerFilterer) WatchChooseTokenMint(opts *bind.WatchOpts, sink chan<- *MintManagerChooseTokenMint, vectorId [][32]byte, contractAddress []common.Address, onChainVector []bool) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var onChainVectorRule []interface{}
	for _, onChainVectorItem := range onChainVector {
		onChainVectorRule = append(onChainVectorRule, onChainVectorItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "ChooseTokenMint", vectorIdRule, contractAddressRule, onChainVectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerChooseTokenMint)
				if err := _MintManager.contract.UnpackLog(event, "ChooseTokenMint", log); err != nil {
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

// ParseChooseTokenMint is a log parse operation binding the contract event 0xe0bf8a4af82e2af496af5f4957e2767f8b52e51a77caedd2f30a1843872d1b7c.
//
// Solidity: event ChooseTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256[] tokenIds)
func (_MintManager *MintManagerFilterer) ParseChooseTokenMint(log types.Log) (*MintManagerChooseTokenMint, error) {
	event := new(MintManagerChooseTokenMint)
	if err := _MintManager.contract.UnpackLog(event, "ChooseTokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerERC20PaymentIterator is returned from FilterERC20Payment and is used to iterate over the raw logs and unpacked data for ERC20Payment events raised by the MintManager contract.
type MintManagerERC20PaymentIterator struct {
	Event *MintManagerERC20Payment // Event containing the contract specifics and raw log

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
func (it *MintManagerERC20PaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerERC20Payment)
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
		it.Event = new(MintManagerERC20Payment)
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
func (it *MintManagerERC20PaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerERC20PaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerERC20Payment represents a ERC20Payment event raised by the MintManager contract.
type MintManagerERC20Payment struct {
	Currency             common.Address
	PaymentRecipient     common.Address
	VectorId             [32]byte
	Payer                common.Address
	AmountToCreator      *big.Int
	PercentageBPSOfTotal uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterERC20Payment is a free log retrieval operation binding the contract event 0xc899cbcc4511003ff90131e8b89605738e9a7f4925273377ae479a673cf5038c.
//
// Solidity: event ERC20Payment(address indexed currency, address indexed paymentRecipient, bytes32 indexed vectorId, address payer, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) FilterERC20Payment(opts *bind.FilterOpts, currency []common.Address, paymentRecipient []common.Address, vectorId [][32]byte) (*MintManagerERC20PaymentIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var paymentRecipientRule []interface{}
	for _, paymentRecipientItem := range paymentRecipient {
		paymentRecipientRule = append(paymentRecipientRule, paymentRecipientItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "ERC20Payment", currencyRule, paymentRecipientRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerERC20PaymentIterator{contract: _MintManager.contract, event: "ERC20Payment", logs: logs, sub: sub}, nil
}

// WatchERC20Payment is a free log subscription operation binding the contract event 0xc899cbcc4511003ff90131e8b89605738e9a7f4925273377ae479a673cf5038c.
//
// Solidity: event ERC20Payment(address indexed currency, address indexed paymentRecipient, bytes32 indexed vectorId, address payer, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) WatchERC20Payment(opts *bind.WatchOpts, sink chan<- *MintManagerERC20Payment, currency []common.Address, paymentRecipient []common.Address, vectorId [][32]byte) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var paymentRecipientRule []interface{}
	for _, paymentRecipientItem := range paymentRecipient {
		paymentRecipientRule = append(paymentRecipientRule, paymentRecipientItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "ERC20Payment", currencyRule, paymentRecipientRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerERC20Payment)
				if err := _MintManager.contract.UnpackLog(event, "ERC20Payment", log); err != nil {
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

// ParseERC20Payment is a log parse operation binding the contract event 0xc899cbcc4511003ff90131e8b89605738e9a7f4925273377ae479a673cf5038c.
//
// Solidity: event ERC20Payment(address indexed currency, address indexed paymentRecipient, bytes32 indexed vectorId, address payer, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) ParseERC20Payment(log types.Log) (*MintManagerERC20Payment, error) {
	event := new(MintManagerERC20Payment)
	if err := _MintManager.contract.UnpackLog(event, "ERC20Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerERC20PaymentMetaTxPacketsIterator is returned from FilterERC20PaymentMetaTxPackets and is used to iterate over the raw logs and unpacked data for ERC20PaymentMetaTxPackets events raised by the MintManager contract.
type MintManagerERC20PaymentMetaTxPacketsIterator struct {
	Event *MintManagerERC20PaymentMetaTxPackets // Event containing the contract specifics and raw log

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
func (it *MintManagerERC20PaymentMetaTxPacketsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerERC20PaymentMetaTxPackets)
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
		it.Event = new(MintManagerERC20PaymentMetaTxPackets)
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
func (it *MintManagerERC20PaymentMetaTxPacketsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerERC20PaymentMetaTxPacketsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerERC20PaymentMetaTxPackets represents a ERC20PaymentMetaTxPackets event raised by the MintManager contract.
type MintManagerERC20PaymentMetaTxPackets struct {
	Currency                 common.Address
	MsgSender                common.Address
	VectorId                 [32]byte
	PurchaseToCreatorPacket  MintManagerPurchaserMetaTxPacket
	PurchaseToPlatformPacket MintManagerPurchaserMetaTxPacket
	Amount                   *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterERC20PaymentMetaTxPackets is a free log retrieval operation binding the contract event 0xe415115991e960f57314b86d3f4e4db15b9ac36256043523adcb60f15c3ddefd.
//
// Solidity: event ERC20PaymentMetaTxPackets(address indexed currency, address indexed msgSender, bytes32 indexed vectorId, (bytes,bytes32,bytes32,uint8) purchaseToCreatorPacket, (bytes,bytes32,bytes32,uint8) purchaseToPlatformPacket, uint256 amount)
func (_MintManager *MintManagerFilterer) FilterERC20PaymentMetaTxPackets(opts *bind.FilterOpts, currency []common.Address, msgSender []common.Address, vectorId [][32]byte) (*MintManagerERC20PaymentMetaTxPacketsIterator, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "ERC20PaymentMetaTxPackets", currencyRule, msgSenderRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerERC20PaymentMetaTxPacketsIterator{contract: _MintManager.contract, event: "ERC20PaymentMetaTxPackets", logs: logs, sub: sub}, nil
}

// WatchERC20PaymentMetaTxPackets is a free log subscription operation binding the contract event 0xe415115991e960f57314b86d3f4e4db15b9ac36256043523adcb60f15c3ddefd.
//
// Solidity: event ERC20PaymentMetaTxPackets(address indexed currency, address indexed msgSender, bytes32 indexed vectorId, (bytes,bytes32,bytes32,uint8) purchaseToCreatorPacket, (bytes,bytes32,bytes32,uint8) purchaseToPlatformPacket, uint256 amount)
func (_MintManager *MintManagerFilterer) WatchERC20PaymentMetaTxPackets(opts *bind.WatchOpts, sink chan<- *MintManagerERC20PaymentMetaTxPackets, currency []common.Address, msgSender []common.Address, vectorId [][32]byte) (event.Subscription, error) {

	var currencyRule []interface{}
	for _, currencyItem := range currency {
		currencyRule = append(currencyRule, currencyItem)
	}
	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "ERC20PaymentMetaTxPackets", currencyRule, msgSenderRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerERC20PaymentMetaTxPackets)
				if err := _MintManager.contract.UnpackLog(event, "ERC20PaymentMetaTxPackets", log); err != nil {
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

// ParseERC20PaymentMetaTxPackets is a log parse operation binding the contract event 0xe415115991e960f57314b86d3f4e4db15b9ac36256043523adcb60f15c3ddefd.
//
// Solidity: event ERC20PaymentMetaTxPackets(address indexed currency, address indexed msgSender, bytes32 indexed vectorId, (bytes,bytes32,bytes32,uint8) purchaseToCreatorPacket, (bytes,bytes32,bytes32,uint8) purchaseToPlatformPacket, uint256 amount)
func (_MintManager *MintManagerFilterer) ParseERC20PaymentMetaTxPackets(log types.Log) (*MintManagerERC20PaymentMetaTxPackets, error) {
	event := new(MintManagerERC20PaymentMetaTxPackets)
	if err := _MintManager.contract.UnpackLog(event, "ERC20PaymentMetaTxPackets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerEditionVectorCreatedIterator is returned from FilterEditionVectorCreated and is used to iterate over the raw logs and unpacked data for EditionVectorCreated events raised by the MintManager contract.
type MintManagerEditionVectorCreatedIterator struct {
	Event *MintManagerEditionVectorCreated // Event containing the contract specifics and raw log

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
func (it *MintManagerEditionVectorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerEditionVectorCreated)
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
		it.Event = new(MintManagerEditionVectorCreated)
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
func (it *MintManagerEditionVectorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerEditionVectorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerEditionVectorCreated represents a EditionVectorCreated event raised by the MintManager contract.
type MintManagerEditionVectorCreated struct {
	VectorId        *big.Int
	EditionId       *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEditionVectorCreated is a free log retrieval operation binding the contract event 0xa712e8b25b3d4d043988e80f0a4087773b1c7e29e4115a4256e86aebe91c9be9.
//
// Solidity: event EditionVectorCreated(uint256 indexed vectorId, uint48 indexed editionId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) FilterEditionVectorCreated(opts *bind.FilterOpts, vectorId []*big.Int, editionId []*big.Int, contractAddress []common.Address) (*MintManagerEditionVectorCreatedIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "EditionVectorCreated", vectorIdRule, editionIdRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerEditionVectorCreatedIterator{contract: _MintManager.contract, event: "EditionVectorCreated", logs: logs, sub: sub}, nil
}

// WatchEditionVectorCreated is a free log subscription operation binding the contract event 0xa712e8b25b3d4d043988e80f0a4087773b1c7e29e4115a4256e86aebe91c9be9.
//
// Solidity: event EditionVectorCreated(uint256 indexed vectorId, uint48 indexed editionId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) WatchEditionVectorCreated(opts *bind.WatchOpts, sink chan<- *MintManagerEditionVectorCreated, vectorId []*big.Int, editionId []*big.Int, contractAddress []common.Address) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var editionIdRule []interface{}
	for _, editionIdItem := range editionId {
		editionIdRule = append(editionIdRule, editionIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "EditionVectorCreated", vectorIdRule, editionIdRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerEditionVectorCreated)
				if err := _MintManager.contract.UnpackLog(event, "EditionVectorCreated", log); err != nil {
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

// ParseEditionVectorCreated is a log parse operation binding the contract event 0xa712e8b25b3d4d043988e80f0a4087773b1c7e29e4115a4256e86aebe91c9be9.
//
// Solidity: event EditionVectorCreated(uint256 indexed vectorId, uint48 indexed editionId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) ParseEditionVectorCreated(log types.Log) (*MintManagerEditionVectorCreated, error) {
	event := new(MintManagerEditionVectorCreated)
	if err := _MintManager.contract.UnpackLog(event, "EditionVectorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MintManager contract.
type MintManagerInitializedIterator struct {
	Event *MintManagerInitialized // Event containing the contract specifics and raw log

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
func (it *MintManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerInitialized)
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
		it.Event = new(MintManagerInitialized)
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
func (it *MintManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerInitialized represents a Initialized event raised by the MintManager contract.
type MintManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintManager *MintManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*MintManagerInitializedIterator, error) {

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MintManagerInitializedIterator{contract: _MintManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_MintManager *MintManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MintManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerInitialized)
				if err := _MintManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_MintManager *MintManagerFilterer) ParseInitialized(log types.Log) (*MintManagerInitialized, error) {
	event := new(MintManagerInitialized)
	if err := _MintManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerNativeGasTokenPaymentIterator is returned from FilterNativeGasTokenPayment and is used to iterate over the raw logs and unpacked data for NativeGasTokenPayment events raised by the MintManager contract.
type MintManagerNativeGasTokenPaymentIterator struct {
	Event *MintManagerNativeGasTokenPayment // Event containing the contract specifics and raw log

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
func (it *MintManagerNativeGasTokenPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerNativeGasTokenPayment)
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
		it.Event = new(MintManagerNativeGasTokenPayment)
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
func (it *MintManagerNativeGasTokenPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerNativeGasTokenPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerNativeGasTokenPayment represents a NativeGasTokenPayment event raised by the MintManager contract.
type MintManagerNativeGasTokenPayment struct {
	PaymentRecipient     common.Address
	VectorId             [32]byte
	AmountToCreator      *big.Int
	PercentageBPSOfTotal uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNativeGasTokenPayment is a free log retrieval operation binding the contract event 0x9363885e28e7ba67b096932f9f00dff44742731d6cb4fa26ccd4424e78e41e13.
//
// Solidity: event NativeGasTokenPayment(address indexed paymentRecipient, bytes32 indexed vectorId, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) FilterNativeGasTokenPayment(opts *bind.FilterOpts, paymentRecipient []common.Address, vectorId [][32]byte) (*MintManagerNativeGasTokenPaymentIterator, error) {

	var paymentRecipientRule []interface{}
	for _, paymentRecipientItem := range paymentRecipient {
		paymentRecipientRule = append(paymentRecipientRule, paymentRecipientItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "NativeGasTokenPayment", paymentRecipientRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerNativeGasTokenPaymentIterator{contract: _MintManager.contract, event: "NativeGasTokenPayment", logs: logs, sub: sub}, nil
}

// WatchNativeGasTokenPayment is a free log subscription operation binding the contract event 0x9363885e28e7ba67b096932f9f00dff44742731d6cb4fa26ccd4424e78e41e13.
//
// Solidity: event NativeGasTokenPayment(address indexed paymentRecipient, bytes32 indexed vectorId, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) WatchNativeGasTokenPayment(opts *bind.WatchOpts, sink chan<- *MintManagerNativeGasTokenPayment, paymentRecipient []common.Address, vectorId [][32]byte) (event.Subscription, error) {

	var paymentRecipientRule []interface{}
	for _, paymentRecipientItem := range paymentRecipient {
		paymentRecipientRule = append(paymentRecipientRule, paymentRecipientItem)
	}
	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "NativeGasTokenPayment", paymentRecipientRule, vectorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerNativeGasTokenPayment)
				if err := _MintManager.contract.UnpackLog(event, "NativeGasTokenPayment", log); err != nil {
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

// ParseNativeGasTokenPayment is a log parse operation binding the contract event 0x9363885e28e7ba67b096932f9f00dff44742731d6cb4fa26ccd4424e78e41e13.
//
// Solidity: event NativeGasTokenPayment(address indexed paymentRecipient, bytes32 indexed vectorId, uint256 amountToCreator, uint32 percentageBPSOfTotal)
func (_MintManager *MintManagerFilterer) ParseNativeGasTokenPayment(log types.Log) (*MintManagerNativeGasTokenPayment, error) {
	event := new(MintManagerNativeGasTokenPayment)
	if err := _MintManager.contract.UnpackLog(event, "NativeGasTokenPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerNumTokenMintIterator is returned from FilterNumTokenMint and is used to iterate over the raw logs and unpacked data for NumTokenMint events raised by the MintManager contract.
type MintManagerNumTokenMintIterator struct {
	Event *MintManagerNumTokenMint // Event containing the contract specifics and raw log

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
func (it *MintManagerNumTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerNumTokenMint)
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
		it.Event = new(MintManagerNumTokenMint)
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
func (it *MintManagerNumTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerNumTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerNumTokenMint represents a NumTokenMint event raised by the MintManager contract.
type MintManagerNumTokenMint struct {
	VectorId        [32]byte
	ContractAddress common.Address
	OnChainVector   bool
	NumMinted       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNumTokenMint is a free log retrieval operation binding the contract event 0x981414aed4973b05aa301314dc13a5a4077f24490497b98bc270852581c1c578.
//
// Solidity: event NumTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256 numMinted)
func (_MintManager *MintManagerFilterer) FilterNumTokenMint(opts *bind.FilterOpts, vectorId [][32]byte, contractAddress []common.Address, onChainVector []bool) (*MintManagerNumTokenMintIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var onChainVectorRule []interface{}
	for _, onChainVectorItem := range onChainVector {
		onChainVectorRule = append(onChainVectorRule, onChainVectorItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "NumTokenMint", vectorIdRule, contractAddressRule, onChainVectorRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerNumTokenMintIterator{contract: _MintManager.contract, event: "NumTokenMint", logs: logs, sub: sub}, nil
}

// WatchNumTokenMint is a free log subscription operation binding the contract event 0x981414aed4973b05aa301314dc13a5a4077f24490497b98bc270852581c1c578.
//
// Solidity: event NumTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256 numMinted)
func (_MintManager *MintManagerFilterer) WatchNumTokenMint(opts *bind.WatchOpts, sink chan<- *MintManagerNumTokenMint, vectorId [][32]byte, contractAddress []common.Address, onChainVector []bool) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var onChainVectorRule []interface{}
	for _, onChainVectorItem := range onChainVector {
		onChainVectorRule = append(onChainVectorRule, onChainVectorItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "NumTokenMint", vectorIdRule, contractAddressRule, onChainVectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerNumTokenMint)
				if err := _MintManager.contract.UnpackLog(event, "NumTokenMint", log); err != nil {
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

// ParseNumTokenMint is a log parse operation binding the contract event 0x981414aed4973b05aa301314dc13a5a4077f24490497b98bc270852581c1c578.
//
// Solidity: event NumTokenMint(bytes32 indexed vectorId, address indexed contractAddress, bool indexed onChainVector, uint256 numMinted)
func (_MintManager *MintManagerFilterer) ParseNumTokenMint(log types.Log) (*MintManagerNumTokenMint, error) {
	event := new(MintManagerNumTokenMint)
	if err := _MintManager.contract.UnpackLog(event, "NumTokenMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MintManager contract.
type MintManagerOwnershipTransferredIterator struct {
	Event *MintManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerOwnershipTransferred)
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
		it.Event = new(MintManagerOwnershipTransferred)
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
func (it *MintManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerOwnershipTransferred represents a OwnershipTransferred event raised by the MintManager contract.
type MintManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerOwnershipTransferredIterator{contract: _MintManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MintManager *MintManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerOwnershipTransferred)
				if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MintManager *MintManagerFilterer) ParseOwnershipTransferred(log types.Log) (*MintManagerOwnershipTransferred, error) {
	event := new(MintManagerOwnershipTransferred)
	if err := _MintManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerPlatformExecutorChangedIterator is returned from FilterPlatformExecutorChanged and is used to iterate over the raw logs and unpacked data for PlatformExecutorChanged events raised by the MintManager contract.
type MintManagerPlatformExecutorChangedIterator struct {
	Event *MintManagerPlatformExecutorChanged // Event containing the contract specifics and raw log

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
func (it *MintManagerPlatformExecutorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerPlatformExecutorChanged)
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
		it.Event = new(MintManagerPlatformExecutorChanged)
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
func (it *MintManagerPlatformExecutorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerPlatformExecutorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerPlatformExecutorChanged represents a PlatformExecutorChanged event raised by the MintManager contract.
type MintManagerPlatformExecutorChanged struct {
	Executor common.Address
	Added    bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlatformExecutorChanged is a free log retrieval operation binding the contract event 0x88d9d369733d630a0dbd71ed25dcafecbae7bc29d54a925afc1d003acf2af385.
//
// Solidity: event PlatformExecutorChanged(address indexed executor, bool indexed added)
func (_MintManager *MintManagerFilterer) FilterPlatformExecutorChanged(opts *bind.FilterOpts, executor []common.Address, added []bool) (*MintManagerPlatformExecutorChangedIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "PlatformExecutorChanged", executorRule, addedRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerPlatformExecutorChangedIterator{contract: _MintManager.contract, event: "PlatformExecutorChanged", logs: logs, sub: sub}, nil
}

// WatchPlatformExecutorChanged is a free log subscription operation binding the contract event 0x88d9d369733d630a0dbd71ed25dcafecbae7bc29d54a925afc1d003acf2af385.
//
// Solidity: event PlatformExecutorChanged(address indexed executor, bool indexed added)
func (_MintManager *MintManagerFilterer) WatchPlatformExecutorChanged(opts *bind.WatchOpts, sink chan<- *MintManagerPlatformExecutorChanged, executor []common.Address, added []bool) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var addedRule []interface{}
	for _, addedItem := range added {
		addedRule = append(addedRule, addedItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "PlatformExecutorChanged", executorRule, addedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerPlatformExecutorChanged)
				if err := _MintManager.contract.UnpackLog(event, "PlatformExecutorChanged", log); err != nil {
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

// ParsePlatformExecutorChanged is a log parse operation binding the contract event 0x88d9d369733d630a0dbd71ed25dcafecbae7bc29d54a925afc1d003acf2af385.
//
// Solidity: event PlatformExecutorChanged(address indexed executor, bool indexed added)
func (_MintManager *MintManagerFilterer) ParsePlatformExecutorChanged(log types.Log) (*MintManagerPlatformExecutorChanged, error) {
	event := new(MintManagerPlatformExecutorChanged)
	if err := _MintManager.contract.UnpackLog(event, "PlatformExecutorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerSeriesVectorCreatedIterator is returned from FilterSeriesVectorCreated and is used to iterate over the raw logs and unpacked data for SeriesVectorCreated events raised by the MintManager contract.
type MintManagerSeriesVectorCreatedIterator struct {
	Event *MintManagerSeriesVectorCreated // Event containing the contract specifics and raw log

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
func (it *MintManagerSeriesVectorCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerSeriesVectorCreated)
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
		it.Event = new(MintManagerSeriesVectorCreated)
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
func (it *MintManagerSeriesVectorCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerSeriesVectorCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerSeriesVectorCreated represents a SeriesVectorCreated event raised by the MintManager contract.
type MintManagerSeriesVectorCreated struct {
	VectorId        *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSeriesVectorCreated is a free log retrieval operation binding the contract event 0x7258df9bfe0a9fb9cf1285396575e6472f56ca38b4851afcb725c82726fd67ff.
//
// Solidity: event SeriesVectorCreated(uint256 indexed vectorId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) FilterSeriesVectorCreated(opts *bind.FilterOpts, vectorId []*big.Int, contractAddress []common.Address) (*MintManagerSeriesVectorCreatedIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "SeriesVectorCreated", vectorIdRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerSeriesVectorCreatedIterator{contract: _MintManager.contract, event: "SeriesVectorCreated", logs: logs, sub: sub}, nil
}

// WatchSeriesVectorCreated is a free log subscription operation binding the contract event 0x7258df9bfe0a9fb9cf1285396575e6472f56ca38b4851afcb725c82726fd67ff.
//
// Solidity: event SeriesVectorCreated(uint256 indexed vectorId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) WatchSeriesVectorCreated(opts *bind.WatchOpts, sink chan<- *MintManagerSeriesVectorCreated, vectorId []*big.Int, contractAddress []common.Address) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "SeriesVectorCreated", vectorIdRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerSeriesVectorCreated)
				if err := _MintManager.contract.UnpackLog(event, "SeriesVectorCreated", log); err != nil {
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

// ParseSeriesVectorCreated is a log parse operation binding the contract event 0x7258df9bfe0a9fb9cf1285396575e6472f56ca38b4851afcb725c82726fd67ff.
//
// Solidity: event SeriesVectorCreated(uint256 indexed vectorId, address indexed contractAddress)
func (_MintManager *MintManagerFilterer) ParseSeriesVectorCreated(log types.Log) (*MintManagerSeriesVectorCreated, error) {
	event := new(MintManagerSeriesVectorCreated)
	if err := _MintManager.contract.UnpackLog(event, "SeriesVectorCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the MintManager contract.
type MintManagerUpgradedIterator struct {
	Event *MintManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *MintManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerUpgraded)
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
		it.Event = new(MintManagerUpgraded)
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
func (it *MintManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerUpgraded represents a Upgraded event raised by the MintManager contract.
type MintManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MintManager *MintManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MintManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerUpgradedIterator{contract: _MintManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MintManager *MintManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MintManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerUpgraded)
				if err := _MintManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_MintManager *MintManagerFilterer) ParseUpgraded(log types.Log) (*MintManagerUpgraded, error) {
	event := new(MintManagerUpgraded)
	if err := _MintManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerVectorDeletedIterator is returned from FilterVectorDeleted and is used to iterate over the raw logs and unpacked data for VectorDeleted events raised by the MintManager contract.
type MintManagerVectorDeletedIterator struct {
	Event *MintManagerVectorDeleted // Event containing the contract specifics and raw log

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
func (it *MintManagerVectorDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerVectorDeleted)
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
		it.Event = new(MintManagerVectorDeleted)
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
func (it *MintManagerVectorDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerVectorDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerVectorDeleted represents a VectorDeleted event raised by the MintManager contract.
type MintManagerVectorDeleted struct {
	VectorId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVectorDeleted is a free log retrieval operation binding the contract event 0xc838617e2997901e8e4856126ebd46593aef10fb97d78f88b4635c9420f67316.
//
// Solidity: event VectorDeleted(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) FilterVectorDeleted(opts *bind.FilterOpts, vectorId []*big.Int) (*MintManagerVectorDeletedIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "VectorDeleted", vectorIdRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerVectorDeletedIterator{contract: _MintManager.contract, event: "VectorDeleted", logs: logs, sub: sub}, nil
}

// WatchVectorDeleted is a free log subscription operation binding the contract event 0xc838617e2997901e8e4856126ebd46593aef10fb97d78f88b4635c9420f67316.
//
// Solidity: event VectorDeleted(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) WatchVectorDeleted(opts *bind.WatchOpts, sink chan<- *MintManagerVectorDeleted, vectorId []*big.Int) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "VectorDeleted", vectorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerVectorDeleted)
				if err := _MintManager.contract.UnpackLog(event, "VectorDeleted", log); err != nil {
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

// ParseVectorDeleted is a log parse operation binding the contract event 0xc838617e2997901e8e4856126ebd46593aef10fb97d78f88b4635c9420f67316.
//
// Solidity: event VectorDeleted(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) ParseVectorDeleted(log types.Log) (*MintManagerVectorDeleted, error) {
	event := new(MintManagerVectorDeleted)
	if err := _MintManager.contract.UnpackLog(event, "VectorDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerVectorMetadataSetIterator is returned from FilterVectorMetadataSet and is used to iterate over the raw logs and unpacked data for VectorMetadataSet events raised by the MintManager contract.
type MintManagerVectorMetadataSetIterator struct {
	Event *MintManagerVectorMetadataSet // Event containing the contract specifics and raw log

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
func (it *MintManagerVectorMetadataSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerVectorMetadataSet)
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
		it.Event = new(MintManagerVectorMetadataSet)
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
func (it *MintManagerVectorMetadataSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerVectorMetadataSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerVectorMetadataSet represents a VectorMetadataSet event raised by the MintManager contract.
type MintManagerVectorMetadataSet struct {
	VectorId     *big.Int
	Paused       bool
	FlexibleData *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVectorMetadataSet is a free log retrieval operation binding the contract event 0xaaca5cb46300e4b20595b143bc883119e775ef88ff77f45fa989ca323576f06b.
//
// Solidity: event VectorMetadataSet(uint256 indexed vectorId, bool indexed paused, uint128 indexed flexibleData)
func (_MintManager *MintManagerFilterer) FilterVectorMetadataSet(opts *bind.FilterOpts, vectorId []*big.Int, paused []bool, flexibleData []*big.Int) (*MintManagerVectorMetadataSetIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var pausedRule []interface{}
	for _, pausedItem := range paused {
		pausedRule = append(pausedRule, pausedItem)
	}
	var flexibleDataRule []interface{}
	for _, flexibleDataItem := range flexibleData {
		flexibleDataRule = append(flexibleDataRule, flexibleDataItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "VectorMetadataSet", vectorIdRule, pausedRule, flexibleDataRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerVectorMetadataSetIterator{contract: _MintManager.contract, event: "VectorMetadataSet", logs: logs, sub: sub}, nil
}

// WatchVectorMetadataSet is a free log subscription operation binding the contract event 0xaaca5cb46300e4b20595b143bc883119e775ef88ff77f45fa989ca323576f06b.
//
// Solidity: event VectorMetadataSet(uint256 indexed vectorId, bool indexed paused, uint128 indexed flexibleData)
func (_MintManager *MintManagerFilterer) WatchVectorMetadataSet(opts *bind.WatchOpts, sink chan<- *MintManagerVectorMetadataSet, vectorId []*big.Int, paused []bool, flexibleData []*big.Int) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}
	var pausedRule []interface{}
	for _, pausedItem := range paused {
		pausedRule = append(pausedRule, pausedItem)
	}
	var flexibleDataRule []interface{}
	for _, flexibleDataItem := range flexibleData {
		flexibleDataRule = append(flexibleDataRule, flexibleDataItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "VectorMetadataSet", vectorIdRule, pausedRule, flexibleDataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerVectorMetadataSet)
				if err := _MintManager.contract.UnpackLog(event, "VectorMetadataSet", log); err != nil {
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

// ParseVectorMetadataSet is a log parse operation binding the contract event 0xaaca5cb46300e4b20595b143bc883119e775ef88ff77f45fa989ca323576f06b.
//
// Solidity: event VectorMetadataSet(uint256 indexed vectorId, bool indexed paused, uint128 indexed flexibleData)
func (_MintManager *MintManagerFilterer) ParseVectorMetadataSet(log types.Log) (*MintManagerVectorMetadataSet, error) {
	event := new(MintManagerVectorMetadataSet)
	if err := _MintManager.contract.UnpackLog(event, "VectorMetadataSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintManagerVectorUpdatedIterator is returned from FilterVectorUpdated and is used to iterate over the raw logs and unpacked data for VectorUpdated events raised by the MintManager contract.
type MintManagerVectorUpdatedIterator struct {
	Event *MintManagerVectorUpdated // Event containing the contract specifics and raw log

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
func (it *MintManagerVectorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintManagerVectorUpdated)
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
		it.Event = new(MintManagerVectorUpdated)
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
func (it *MintManagerVectorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintManagerVectorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintManagerVectorUpdated represents a VectorUpdated event raised by the MintManager contract.
type MintManagerVectorUpdated struct {
	VectorId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVectorUpdated is a free log retrieval operation binding the contract event 0xe772ce44f6b7edf20d62f174efc62c5a18484d62a710bd48d57af1afd140811c.
//
// Solidity: event VectorUpdated(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) FilterVectorUpdated(opts *bind.FilterOpts, vectorId []*big.Int) (*MintManagerVectorUpdatedIterator, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.FilterLogs(opts, "VectorUpdated", vectorIdRule)
	if err != nil {
		return nil, err
	}
	return &MintManagerVectorUpdatedIterator{contract: _MintManager.contract, event: "VectorUpdated", logs: logs, sub: sub}, nil
}

// WatchVectorUpdated is a free log subscription operation binding the contract event 0xe772ce44f6b7edf20d62f174efc62c5a18484d62a710bd48d57af1afd140811c.
//
// Solidity: event VectorUpdated(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) WatchVectorUpdated(opts *bind.WatchOpts, sink chan<- *MintManagerVectorUpdated, vectorId []*big.Int) (event.Subscription, error) {

	var vectorIdRule []interface{}
	for _, vectorIdItem := range vectorId {
		vectorIdRule = append(vectorIdRule, vectorIdItem)
	}

	logs, sub, err := _MintManager.contract.WatchLogs(opts, "VectorUpdated", vectorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintManagerVectorUpdated)
				if err := _MintManager.contract.UnpackLog(event, "VectorUpdated", log); err != nil {
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

// ParseVectorUpdated is a log parse operation binding the contract event 0xe772ce44f6b7edf20d62f174efc62c5a18484d62a710bd48d57af1afd140811c.
//
// Solidity: event VectorUpdated(uint256 indexed vectorId)
func (_MintManager *MintManagerFilterer) ParseVectorUpdated(log types.Log) (*MintManagerVectorUpdated, error) {
	event := new(MintManagerVectorUpdated)
	if err := _MintManager.contract.UnpackLog(event, "VectorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
