// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kiwistand

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

// IERC721DropAddressMintDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC721DropAddressMintDetails struct {
	TotalMints   *big.Int
	PresaleMints *big.Int
	PublicMints  *big.Int
}

// IERC721DropSaleDetails is an auto generated low-level Go binding around an user-defined struct.
type IERC721DropSaleDetails struct {
	PublicSaleActive          bool
	PresaleActive             bool
	PublicSalePrice           *big.Int
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
	MaxSalePurchasePerAddress *big.Int
	TotalMinted               *big.Int
	MaxSupply                 *big.Int
}

// RewardsSettings is an auto generated low-level Go binding around an user-defined struct.
type RewardsSettings struct {
	CreatorReward        *big.Int
	CreateReferralReward *big.Int
	MintReferralReward   *big.Int
	FirstMinterReward    *big.Int
	ZoraReward           *big.Int
}

// KiwiMetaData contains all meta data concerning the Kiwi contract.
var KiwiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zoraERC721TransferHelper\",\"type\":\"address\"},{\"internalType\":\"contractIFactoryUpgradeGate\",\"name\":\"_factoryUpgradeGate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketFilterDAOAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_mintFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolRewards\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"Access_MissingRoleOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Access_OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Access_WithdrawNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"Admin_InvalidUpgradeAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Admin_UnableToFinalizeNotOpenEdition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CREATOR_FUNDS_RECIPIENT_NOT_SET\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExternalMetadataRenderer_CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ETH_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintSchedule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketFilterDAOAddressNotSupportedForChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintFee_FundsSendFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_SoldOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_CREATE_REFERRAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_OWNER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_PENDING_OWNER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_Inactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_MerkleNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_TooManyForAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolRewards_WithdrawSendFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Purchase_TooManyForAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"correctPrice\",\"type\":\"uint256\"}],\"name\":\"Purchase_WrongPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteOperatorFilterRegistryCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Sale_Inactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"maxRoyaltyBPS\",\"type\":\"uint16\"}],\"name\":\"Setup_RoyaltyPercentageTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Withdraw_FundsSendFailure\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"FundsRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawnBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawnTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"MintComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"MintFeePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfMints\",\"type\":\"uint256\"}],\"name\":\"OpenMintFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"potentialNewOwner\",\"type\":\"address\"}],\"name\":\"OwnerCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"potentialNewOwner\",\"type\":\"address\"}],\"name\":\"OwnerPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPurchasedTokenId\",\"type\":\"uint256\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"SalesConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMetadataRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"UpdatedMetadataRenderer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SALES_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"adminMintAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callMetadataRenderer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computeFreeMintRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computePaidMintRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computeTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"metadataRenderer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"editionSize\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"royaltyBPS\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createReferral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryUpgradeGate\",\"outputs\":[{\"internalType\":\"contractIFactoryUpgradeGate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeOpenEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_editionSize\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyBPS\",\"type\":\"uint16\"},{\"internalType\":\"bytes[]\",\"name\":\"_setupCalls\",\"type\":\"bytes[]\"},{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"_metadataRenderer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metadataRendererInit\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_createReferral\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"manageMarketFilterDAOSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketFilterDAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataRenderer\",\"outputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"}],\"name\":\"mintWithRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"mintedPerAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalMints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"presaleMints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicMints\",\"type\":\"uint256\"}],\"internalType\":\"structIERC721Drop.AddressMintDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleMintsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"purchasePresale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchasePresaleWithComment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"}],\"name\":\"purchasePresaleWithRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchaseWithComment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchaseWithRecipient\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyMintSchedule\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"publicSaleActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"presaleActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"publicSalePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"internalType\":\"structIERC721Drop.SaleDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salesConfig\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"publicSalePrice\",\"type\":\"uint104\"},{\"internalType\":\"uint32\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newRecipientAddress\",\"type\":\"address\"}],\"name\":\"setFundsRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"newRenderer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"setupRenderer\",\"type\":\"bytes\"}],\"name\":\"setMetadataRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint104\",\"name\":\"publicSalePrice\",\"type\":\"uint104\"},{\"internalType\":\"uint32\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setSaleConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"updateCreateReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"updateMarketFilterSettings\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newSchedule\",\"type\":\"uint32\"}],\"name\":\"updateRoyaltyMintSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoraERC721TransferHelper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"zoraFeeForAmount\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KiwiABI is the input ABI used to generate the binding from.
// Deprecated: Use KiwiMetaData.ABI instead.
var KiwiABI = KiwiMetaData.ABI

// Kiwi is an auto generated Go binding around an Ethereum contract.
type Kiwi struct {
	KiwiCaller     // Read-only binding to the contract
	KiwiTransactor // Write-only binding to the contract
	KiwiFilterer   // Log filterer for contract events
}

// KiwiCaller is an auto generated read-only Go binding around an Ethereum contract.
type KiwiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KiwiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KiwiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KiwiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KiwiSession struct {
	Contract     *Kiwi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KiwiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KiwiCallerSession struct {
	Contract *KiwiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KiwiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KiwiTransactorSession struct {
	Contract     *KiwiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KiwiRaw is an auto generated low-level Go binding around an Ethereum contract.
type KiwiRaw struct {
	Contract *Kiwi // Generic contract binding to access the raw methods on
}

// KiwiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KiwiCallerRaw struct {
	Contract *KiwiCaller // Generic read-only contract binding to access the raw methods on
}

// KiwiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KiwiTransactorRaw struct {
	Contract *KiwiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKiwi creates a new instance of Kiwi, bound to a specific deployed contract.
func NewKiwi(address common.Address, backend bind.ContractBackend) (*Kiwi, error) {
	contract, err := bindKiwi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kiwi{KiwiCaller: KiwiCaller{contract: contract}, KiwiTransactor: KiwiTransactor{contract: contract}, KiwiFilterer: KiwiFilterer{contract: contract}}, nil
}

// NewKiwiCaller creates a new read-only instance of Kiwi, bound to a specific deployed contract.
func NewKiwiCaller(address common.Address, caller bind.ContractCaller) (*KiwiCaller, error) {
	contract, err := bindKiwi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KiwiCaller{contract: contract}, nil
}

// NewKiwiTransactor creates a new write-only instance of Kiwi, bound to a specific deployed contract.
func NewKiwiTransactor(address common.Address, transactor bind.ContractTransactor) (*KiwiTransactor, error) {
	contract, err := bindKiwi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KiwiTransactor{contract: contract}, nil
}

// NewKiwiFilterer creates a new log filterer instance of Kiwi, bound to a specific deployed contract.
func NewKiwiFilterer(address common.Address, filterer bind.ContractFilterer) (*KiwiFilterer, error) {
	contract, err := bindKiwi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KiwiFilterer{contract: contract}, nil
}

// bindKiwi binds a generic wrapper to an already deployed contract.
func bindKiwi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := KiwiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kiwi *KiwiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kiwi.Contract.KiwiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kiwi *KiwiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.Contract.KiwiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kiwi *KiwiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kiwi.Contract.KiwiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kiwi *KiwiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kiwi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kiwi *KiwiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kiwi *KiwiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kiwi.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Kiwi *KiwiSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Kiwi.Contract.DEFAULTADMINROLE(&_Kiwi.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Kiwi.Contract.DEFAULTADMINROLE(&_Kiwi.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiSession) MINTERROLE() ([32]byte, error) {
	return _Kiwi.Contract.MINTERROLE(&_Kiwi.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCallerSession) MINTERROLE() ([32]byte, error) {
	return _Kiwi.Contract.MINTERROLE(&_Kiwi.CallOpts)
}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCaller) SALESMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "SALES_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiSession) SALESMANAGERROLE() ([32]byte, error) {
	return _Kiwi.Contract.SALESMANAGERROLE(&_Kiwi.CallOpts)
}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_Kiwi *KiwiCallerSession) SALESMANAGERROLE() ([32]byte, error) {
	return _Kiwi.Contract.SALESMANAGERROLE(&_Kiwi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kiwi *KiwiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kiwi *KiwiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kiwi.Contract.BalanceOf(&_Kiwi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Kiwi *KiwiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Kiwi.Contract.BalanceOf(&_Kiwi.CallOpts, owner)
}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiCaller) ComputeFreeMintRewards(opts *bind.CallOpts, numTokens *big.Int) (RewardsSettings, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "computeFreeMintRewards", numTokens)

	if err != nil {
		return *new(RewardsSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardsSettings)).(*RewardsSettings)

	return out0, err

}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiSession) ComputeFreeMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _Kiwi.Contract.ComputeFreeMintRewards(&_Kiwi.CallOpts, numTokens)
}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiCallerSession) ComputeFreeMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _Kiwi.Contract.ComputeFreeMintRewards(&_Kiwi.CallOpts, numTokens)
}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiCaller) ComputePaidMintRewards(opts *bind.CallOpts, numTokens *big.Int) (RewardsSettings, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "computePaidMintRewards", numTokens)

	if err != nil {
		return *new(RewardsSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardsSettings)).(*RewardsSettings)

	return out0, err

}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiSession) ComputePaidMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _Kiwi.Contract.ComputePaidMintRewards(&_Kiwi.CallOpts, numTokens)
}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_Kiwi *KiwiCallerSession) ComputePaidMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _Kiwi.Contract.ComputePaidMintRewards(&_Kiwi.CallOpts, numTokens)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_Kiwi *KiwiCaller) ComputeTotalReward(opts *bind.CallOpts, numTokens *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "computeTotalReward", numTokens)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_Kiwi *KiwiSession) ComputeTotalReward(numTokens *big.Int) (*big.Int, error) {
	return _Kiwi.Contract.ComputeTotalReward(&_Kiwi.CallOpts, numTokens)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_Kiwi *KiwiCallerSession) ComputeTotalReward(numTokens *big.Int) (*big.Int, error) {
	return _Kiwi.Contract.ComputeTotalReward(&_Kiwi.CallOpts, numTokens)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address metadataRenderer, uint64 editionSize, uint16 royaltyBPS, address fundsRecipient)
func (_Kiwi *KiwiCaller) Config(opts *bind.CallOpts) (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		MetadataRenderer common.Address
		EditionSize      uint64
		RoyaltyBPS       uint16
		FundsRecipient   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MetadataRenderer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EditionSize = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RoyaltyBPS = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.FundsRecipient = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address metadataRenderer, uint64 editionSize, uint16 royaltyBPS, address fundsRecipient)
func (_Kiwi *KiwiSession) Config() (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	return _Kiwi.Contract.Config(&_Kiwi.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address metadataRenderer, uint64 editionSize, uint16 royaltyBPS, address fundsRecipient)
func (_Kiwi *KiwiCallerSession) Config() (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	return _Kiwi.Contract.Config(&_Kiwi.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Kiwi *KiwiCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Kiwi *KiwiSession) ContractURI() (string, error) {
	return _Kiwi.Contract.ContractURI(&_Kiwi.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Kiwi *KiwiCallerSession) ContractURI() (string, error) {
	return _Kiwi.Contract.ContractURI(&_Kiwi.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_Kiwi *KiwiCaller) ContractVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_Kiwi *KiwiSession) ContractVersion() (uint32, error) {
	return _Kiwi.Contract.ContractVersion(&_Kiwi.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_Kiwi *KiwiCallerSession) ContractVersion() (uint32, error) {
	return _Kiwi.Contract.ContractVersion(&_Kiwi.CallOpts)
}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_Kiwi *KiwiCaller) CreateReferral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "createReferral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_Kiwi *KiwiSession) CreateReferral() (common.Address, error) {
	return _Kiwi.Contract.CreateReferral(&_Kiwi.CallOpts)
}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_Kiwi *KiwiCallerSession) CreateReferral() (common.Address, error) {
	return _Kiwi.Contract.CreateReferral(&_Kiwi.CallOpts)
}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_Kiwi *KiwiCaller) FactoryUpgradeGate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "factoryUpgradeGate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_Kiwi *KiwiSession) FactoryUpgradeGate() (common.Address, error) {
	return _Kiwi.Contract.FactoryUpgradeGate(&_Kiwi.CallOpts)
}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_Kiwi *KiwiCallerSession) FactoryUpgradeGate() (common.Address, error) {
	return _Kiwi.Contract.FactoryUpgradeGate(&_Kiwi.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kiwi.Contract.GetApproved(&_Kiwi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Kiwi.Contract.GetApproved(&_Kiwi.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Kiwi *KiwiCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Kiwi *KiwiSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Kiwi.Contract.GetRoleAdmin(&_Kiwi.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Kiwi *KiwiCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Kiwi.Contract.GetRoleAdmin(&_Kiwi.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Kiwi *KiwiCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Kiwi *KiwiSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Kiwi.Contract.HasRole(&_Kiwi.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Kiwi *KiwiCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Kiwi.Contract.HasRole(&_Kiwi.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_Kiwi *KiwiCaller) IsAdmin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "isAdmin", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_Kiwi *KiwiSession) IsAdmin(user common.Address) (bool, error) {
	return _Kiwi.Contract.IsAdmin(&_Kiwi.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_Kiwi *KiwiCallerSession) IsAdmin(user common.Address) (bool, error) {
	return _Kiwi.Contract.IsAdmin(&_Kiwi.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_Kiwi *KiwiCaller) IsApprovedForAll(opts *bind.CallOpts, nftOwner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "isApprovedForAll", nftOwner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_Kiwi *KiwiSession) IsApprovedForAll(nftOwner common.Address, operator common.Address) (bool, error) {
	return _Kiwi.Contract.IsApprovedForAll(&_Kiwi.CallOpts, nftOwner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_Kiwi *KiwiCallerSession) IsApprovedForAll(nftOwner common.Address, operator common.Address) (bool, error) {
	return _Kiwi.Contract.IsApprovedForAll(&_Kiwi.CallOpts, nftOwner, operator)
}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_Kiwi *KiwiCaller) MarketFilterDAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "marketFilterDAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_Kiwi *KiwiSession) MarketFilterDAOAddress() (common.Address, error) {
	return _Kiwi.Contract.MarketFilterDAOAddress(&_Kiwi.CallOpts)
}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_Kiwi *KiwiCallerSession) MarketFilterDAOAddress() (common.Address, error) {
	return _Kiwi.Contract.MarketFilterDAOAddress(&_Kiwi.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_Kiwi *KiwiCaller) MetadataRenderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "metadataRenderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_Kiwi *KiwiSession) MetadataRenderer() (common.Address, error) {
	return _Kiwi.Contract.MetadataRenderer(&_Kiwi.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_Kiwi *KiwiCallerSession) MetadataRenderer() (common.Address, error) {
	return _Kiwi.Contract.MetadataRenderer(&_Kiwi.CallOpts)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_Kiwi *KiwiCaller) MintedPerAddress(opts *bind.CallOpts, minter common.Address) (IERC721DropAddressMintDetails, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "mintedPerAddress", minter)

	if err != nil {
		return *new(IERC721DropAddressMintDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721DropAddressMintDetails)).(*IERC721DropAddressMintDetails)

	return out0, err

}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_Kiwi *KiwiSession) MintedPerAddress(minter common.Address) (IERC721DropAddressMintDetails, error) {
	return _Kiwi.Contract.MintedPerAddress(&_Kiwi.CallOpts, minter)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_Kiwi *KiwiCallerSession) MintedPerAddress(minter common.Address) (IERC721DropAddressMintDetails, error) {
	return _Kiwi.Contract.MintedPerAddress(&_Kiwi.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kiwi *KiwiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kiwi *KiwiSession) Name() (string, error) {
	return _Kiwi.Contract.Name(&_Kiwi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kiwi *KiwiCallerSession) Name() (string, error) {
	return _Kiwi.Contract.Name(&_Kiwi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kiwi *KiwiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kiwi *KiwiSession) Owner() (common.Address, error) {
	return _Kiwi.Contract.Owner(&_Kiwi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Kiwi *KiwiCallerSession) Owner() (common.Address, error) {
	return _Kiwi.Contract.Owner(&_Kiwi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kiwi.Contract.OwnerOf(&_Kiwi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Kiwi *KiwiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Kiwi.Contract.OwnerOf(&_Kiwi.CallOpts, tokenId)
}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_Kiwi *KiwiCaller) PresaleMintsByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "presaleMintsByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_Kiwi *KiwiSession) PresaleMintsByAddress(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.PresaleMintsByAddress(&_Kiwi.CallOpts, arg0)
}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_Kiwi *KiwiCallerSession) PresaleMintsByAddress(arg0 common.Address) (*big.Int, error) {
	return _Kiwi.Contract.PresaleMintsByAddress(&_Kiwi.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Kiwi *KiwiCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Kiwi *KiwiSession) ProxiableUUID() ([32]byte, error) {
	return _Kiwi.Contract.ProxiableUUID(&_Kiwi.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Kiwi *KiwiCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Kiwi.Contract.ProxiableUUID(&_Kiwi.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Kiwi *KiwiCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Kiwi *KiwiSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Kiwi.Contract.RoyaltyInfo(&_Kiwi.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Kiwi *KiwiCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Kiwi.Contract.RoyaltyInfo(&_Kiwi.CallOpts, arg0, _salePrice)
}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_Kiwi *KiwiCaller) RoyaltyMintSchedule(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "royaltyMintSchedule")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_Kiwi *KiwiSession) RoyaltyMintSchedule() (uint32, error) {
	return _Kiwi.Contract.RoyaltyMintSchedule(&_Kiwi.CallOpts)
}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_Kiwi *KiwiCallerSession) RoyaltyMintSchedule() (uint32, error) {
	return _Kiwi.Contract.RoyaltyMintSchedule(&_Kiwi.CallOpts)
}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_Kiwi *KiwiCaller) SaleDetails(opts *bind.CallOpts) (IERC721DropSaleDetails, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "saleDetails")

	if err != nil {
		return *new(IERC721DropSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721DropSaleDetails)).(*IERC721DropSaleDetails)

	return out0, err

}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_Kiwi *KiwiSession) SaleDetails() (IERC721DropSaleDetails, error) {
	return _Kiwi.Contract.SaleDetails(&_Kiwi.CallOpts)
}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_Kiwi *KiwiCallerSession) SaleDetails() (IERC721DropSaleDetails, error) {
	return _Kiwi.Contract.SaleDetails(&_Kiwi.CallOpts)
}

// SalesConfig is a free data retrieval call binding the contract method 0x1d2c0b38.
//
// Solidity: function salesConfig() view returns(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot)
func (_Kiwi *KiwiCaller) SalesConfig(opts *bind.CallOpts) (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "salesConfig")

	outstruct := new(struct {
		PublicSalePrice           *big.Int
		MaxSalePurchasePerAddress uint32
		PublicSaleStart           uint64
		PublicSaleEnd             uint64
		PresaleStart              uint64
		PresaleEnd                uint64
		PresaleMerkleRoot         [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicSalePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaxSalePurchasePerAddress = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.PublicSaleStart = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.PublicSaleEnd = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.PresaleStart = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.PresaleEnd = *abi.ConvertType(out[5], new(uint64)).(*uint64)
	outstruct.PresaleMerkleRoot = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SalesConfig is a free data retrieval call binding the contract method 0x1d2c0b38.
//
// Solidity: function salesConfig() view returns(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot)
func (_Kiwi *KiwiSession) SalesConfig() (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	return _Kiwi.Contract.SalesConfig(&_Kiwi.CallOpts)
}

// SalesConfig is a free data retrieval call binding the contract method 0x1d2c0b38.
//
// Solidity: function salesConfig() view returns(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot)
func (_Kiwi *KiwiCallerSession) SalesConfig() (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	return _Kiwi.Contract.SalesConfig(&_Kiwi.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kiwi *KiwiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kiwi *KiwiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kiwi.Contract.SupportsInterface(&_Kiwi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Kiwi *KiwiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Kiwi.Contract.SupportsInterface(&_Kiwi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kiwi *KiwiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kiwi *KiwiSession) Symbol() (string, error) {
	return _Kiwi.Contract.Symbol(&_Kiwi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Kiwi *KiwiCallerSession) Symbol() (string, error) {
	return _Kiwi.Contract.Symbol(&_Kiwi.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kiwi *KiwiCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kiwi *KiwiSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kiwi.Contract.TokenURI(&_Kiwi.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Kiwi *KiwiCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Kiwi.Contract.TokenURI(&_Kiwi.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiSession) TotalSupply() (*big.Int, error) {
	return _Kiwi.Contract.TotalSupply(&_Kiwi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Kiwi *KiwiCallerSession) TotalSupply() (*big.Int, error) {
	return _Kiwi.Contract.TotalSupply(&_Kiwi.CallOpts)
}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_Kiwi *KiwiCaller) ZoraERC721TransferHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "zoraERC721TransferHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_Kiwi *KiwiSession) ZoraERC721TransferHelper() (common.Address, error) {
	return _Kiwi.Contract.ZoraERC721TransferHelper(&_Kiwi.CallOpts)
}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_Kiwi *KiwiCallerSession) ZoraERC721TransferHelper() (common.Address, error) {
	return _Kiwi.Contract.ZoraERC721TransferHelper(&_Kiwi.CallOpts)
}

// ZoraFeeForAmount is a free data retrieval call binding the contract method 0xee37be39.
//
// Solidity: function zoraFeeForAmount(uint256 quantity) view returns(address recipient, uint256 fee)
func (_Kiwi *KiwiCaller) ZoraFeeForAmount(opts *bind.CallOpts, quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _Kiwi.contract.Call(opts, &out, "zoraFeeForAmount", quantity)

	outstruct := new(struct {
		Recipient common.Address
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipient = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ZoraFeeForAmount is a free data retrieval call binding the contract method 0xee37be39.
//
// Solidity: function zoraFeeForAmount(uint256 quantity) view returns(address recipient, uint256 fee)
func (_Kiwi *KiwiSession) ZoraFeeForAmount(quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	return _Kiwi.Contract.ZoraFeeForAmount(&_Kiwi.CallOpts, quantity)
}

// ZoraFeeForAmount is a free data retrieval call binding the contract method 0xee37be39.
//
// Solidity: function zoraFeeForAmount(uint256 quantity) view returns(address recipient, uint256 fee)
func (_Kiwi *KiwiCallerSession) ZoraFeeForAmount(quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	return _Kiwi.Contract.ZoraFeeForAmount(&_Kiwi.CallOpts, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_Kiwi *KiwiTransactor) AdminMint(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "adminMint", recipient, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_Kiwi *KiwiSession) AdminMint(recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.AdminMint(&_Kiwi.TransactOpts, recipient, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_Kiwi *KiwiTransactorSession) AdminMint(recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.AdminMint(&_Kiwi.TransactOpts, recipient, quantity)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_Kiwi *KiwiTransactor) AdminMintAirdrop(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "adminMintAirdrop", recipients)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_Kiwi *KiwiSession) AdminMintAirdrop(recipients []common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.AdminMintAirdrop(&_Kiwi.TransactOpts, recipients)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_Kiwi *KiwiTransactorSession) AdminMintAirdrop(recipients []common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.AdminMintAirdrop(&_Kiwi.TransactOpts, recipients)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kiwi *KiwiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Approve(&_Kiwi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Approve(&_Kiwi.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Kiwi *KiwiTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Kiwi *KiwiSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Burn(&_Kiwi.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_Kiwi *KiwiTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Burn(&_Kiwi.TransactOpts, tokenId)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_Kiwi *KiwiTransactor) CallMetadataRenderer(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "callMetadataRenderer", data)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_Kiwi *KiwiSession) CallMetadataRenderer(data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.CallMetadataRenderer(&_Kiwi.TransactOpts, data)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_Kiwi *KiwiTransactorSession) CallMetadataRenderer(data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.CallMetadataRenderer(&_Kiwi.TransactOpts, data)
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_Kiwi *KiwiTransactor) FinalizeOpenEdition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "finalizeOpenEdition")
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_Kiwi *KiwiSession) FinalizeOpenEdition() (*types.Transaction, error) {
	return _Kiwi.Contract.FinalizeOpenEdition(&_Kiwi.TransactOpts)
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_Kiwi *KiwiTransactorSession) FinalizeOpenEdition() (*types.Transaction, error) {
	return _Kiwi.Contract.FinalizeOpenEdition(&_Kiwi.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.GrantRole(&_Kiwi.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.GrantRole(&_Kiwi.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_Kiwi *KiwiTransactor) Initialize(opts *bind.TransactOpts, _contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "initialize", _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_Kiwi *KiwiSession) Initialize(_contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.Initialize(&_Kiwi.TransactOpts, _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_Kiwi *KiwiTransactorSession) Initialize(_contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.Initialize(&_Kiwi.TransactOpts, _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_Kiwi *KiwiTransactor) ManageMarketFilterDAOSubscription(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "manageMarketFilterDAOSubscription", enable)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_Kiwi *KiwiSession) ManageMarketFilterDAOSubscription(enable bool) (*types.Transaction, error) {
	return _Kiwi.Contract.ManageMarketFilterDAOSubscription(&_Kiwi.TransactOpts, enable)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_Kiwi *KiwiTransactorSession) ManageMarketFilterDAOSubscription(enable bool) (*types.Transaction, error) {
	return _Kiwi.Contract.ManageMarketFilterDAOSubscription(&_Kiwi.TransactOpts, enable)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiTransactor) MintWithRewards(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "mintWithRewards", recipient, quantity, comment, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiSession) MintWithRewards(recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.MintWithRewards(&_Kiwi.TransactOpts, recipient, quantity, comment, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) MintWithRewards(recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.MintWithRewards(&_Kiwi.TransactOpts, recipient, quantity, comment, mintReferral)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Kiwi *KiwiTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Kiwi *KiwiSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.Multicall(&_Kiwi.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Kiwi *KiwiTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.Multicall(&_Kiwi.TransactOpts, data)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_Kiwi *KiwiTransactor) Purchase(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchase", quantity)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_Kiwi *KiwiSession) Purchase(quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Purchase(&_Kiwi.TransactOpts, quantity)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) Purchase(quantity *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.Purchase(&_Kiwi.TransactOpts, quantity)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_Kiwi *KiwiTransactor) PurchasePresale(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchasePresale", quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_Kiwi *KiwiSession) PurchasePresale(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresale(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) PurchasePresale(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresale(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactor) PurchasePresaleWithComment(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchasePresaleWithComment", quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_Kiwi *KiwiSession) PurchasePresaleWithComment(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresaleWithComment(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) PurchasePresaleWithComment(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresaleWithComment(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiTransactor) PurchasePresaleWithRewards(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchasePresaleWithRewards", quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiSession) PurchasePresaleWithRewards(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresaleWithRewards(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) PurchasePresaleWithRewards(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchasePresaleWithRewards(&_Kiwi.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactor) PurchaseWithComment(opts *bind.TransactOpts, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchaseWithComment", quantity, comment)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiSession) PurchaseWithComment(quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchaseWithComment(&_Kiwi.TransactOpts, quantity, comment)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) PurchaseWithComment(quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchaseWithComment(&_Kiwi.TransactOpts, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactor) PurchaseWithRecipient(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "purchaseWithRecipient", recipient, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiSession) PurchaseWithRecipient(recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchaseWithRecipient(&_Kiwi.TransactOpts, recipient, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_Kiwi *KiwiTransactorSession) PurchaseWithRecipient(recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _Kiwi.Contract.PurchaseWithRecipient(&_Kiwi.TransactOpts, recipient, quantity, comment)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.RenounceRole(&_Kiwi.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.RenounceRole(&_Kiwi.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.RevokeRole(&_Kiwi.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Kiwi *KiwiTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.RevokeRole(&_Kiwi.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.SafeTransferFrom(&_Kiwi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.SafeTransferFrom(&_Kiwi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Kiwi *KiwiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Kiwi *KiwiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SafeTransferFrom0(&_Kiwi.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Kiwi *KiwiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SafeTransferFrom0(&_Kiwi.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kiwi *KiwiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kiwi *KiwiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kiwi.Contract.SetApprovalForAll(&_Kiwi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Kiwi *KiwiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Kiwi.Contract.SetApprovalForAll(&_Kiwi.TransactOpts, operator, approved)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_Kiwi *KiwiTransactor) SetFundsRecipient(opts *bind.TransactOpts, newRecipientAddress common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "setFundsRecipient", newRecipientAddress)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_Kiwi *KiwiSession) SetFundsRecipient(newRecipientAddress common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.SetFundsRecipient(&_Kiwi.TransactOpts, newRecipientAddress)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_Kiwi *KiwiTransactorSession) SetFundsRecipient(newRecipientAddress common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.SetFundsRecipient(&_Kiwi.TransactOpts, newRecipientAddress)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_Kiwi *KiwiTransactor) SetMetadataRenderer(opts *bind.TransactOpts, newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "setMetadataRenderer", newRenderer, setupRenderer)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_Kiwi *KiwiSession) SetMetadataRenderer(newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SetMetadataRenderer(&_Kiwi.TransactOpts, newRenderer, setupRenderer)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_Kiwi *KiwiTransactorSession) SetMetadataRenderer(newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SetMetadataRenderer(&_Kiwi.TransactOpts, newRenderer, setupRenderer)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Kiwi *KiwiTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Kiwi *KiwiSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.SetOwner(&_Kiwi.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_Kiwi *KiwiTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.SetOwner(&_Kiwi.TransactOpts, newOwner)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_Kiwi *KiwiTransactor) SetSaleConfiguration(opts *bind.TransactOpts, publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "setSaleConfiguration", publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_Kiwi *KiwiSession) SetSaleConfiguration(publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SetSaleConfiguration(&_Kiwi.TransactOpts, publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_Kiwi *KiwiTransactorSession) SetSaleConfiguration(publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Kiwi.Contract.SetSaleConfiguration(&_Kiwi.TransactOpts, publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.TransferFrom(&_Kiwi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Kiwi *KiwiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.TransferFrom(&_Kiwi.TransactOpts, from, to, tokenId)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_Kiwi *KiwiTransactor) UpdateCreateReferral(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "updateCreateReferral", recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_Kiwi *KiwiSession) UpdateCreateReferral(recipient common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateCreateReferral(&_Kiwi.TransactOpts, recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_Kiwi *KiwiTransactorSession) UpdateCreateReferral(recipient common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateCreateReferral(&_Kiwi.TransactOpts, recipient)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_Kiwi *KiwiTransactor) UpdateMarketFilterSettings(opts *bind.TransactOpts, args []byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "updateMarketFilterSettings", args)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_Kiwi *KiwiSession) UpdateMarketFilterSettings(args []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateMarketFilterSettings(&_Kiwi.TransactOpts, args)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_Kiwi *KiwiTransactorSession) UpdateMarketFilterSettings(args []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateMarketFilterSettings(&_Kiwi.TransactOpts, args)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_Kiwi *KiwiTransactor) UpdateRoyaltyMintSchedule(opts *bind.TransactOpts, newSchedule uint32) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "updateRoyaltyMintSchedule", newSchedule)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_Kiwi *KiwiSession) UpdateRoyaltyMintSchedule(newSchedule uint32) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateRoyaltyMintSchedule(&_Kiwi.TransactOpts, newSchedule)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_Kiwi *KiwiTransactorSession) UpdateRoyaltyMintSchedule(newSchedule uint32) (*types.Transaction, error) {
	return _Kiwi.Contract.UpdateRoyaltyMintSchedule(&_Kiwi.TransactOpts, newSchedule)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Kiwi *KiwiTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Kiwi *KiwiSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.UpgradeTo(&_Kiwi.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Kiwi *KiwiTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Kiwi.Contract.UpgradeTo(&_Kiwi.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Kiwi *KiwiTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Kiwi *KiwiSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.UpgradeToAndCall(&_Kiwi.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Kiwi *KiwiTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Kiwi.Contract.UpgradeToAndCall(&_Kiwi.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kiwi *KiwiTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kiwi *KiwiSession) Withdraw() (*types.Transaction, error) {
	return _Kiwi.Contract.Withdraw(&_Kiwi.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Kiwi *KiwiTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Kiwi.Contract.Withdraw(&_Kiwi.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactor) WithdrawRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.contract.Transact(opts, "withdrawRewards", to, amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_Kiwi *KiwiSession) WithdrawRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawRewards(&_Kiwi.TransactOpts, to, amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_Kiwi *KiwiTransactorSession) WithdrawRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Kiwi.Contract.WithdrawRewards(&_Kiwi.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kiwi *KiwiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kiwi.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kiwi *KiwiSession) Receive() (*types.Transaction, error) {
	return _Kiwi.Contract.Receive(&_Kiwi.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kiwi *KiwiTransactorSession) Receive() (*types.Transaction, error) {
	return _Kiwi.Contract.Receive(&_Kiwi.TransactOpts)
}

// KiwiAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Kiwi contract.
type KiwiAdminChangedIterator struct {
	Event *KiwiAdminChanged // Event containing the contract specifics and raw log

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
func (it *KiwiAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiAdminChanged)
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
		it.Event = new(KiwiAdminChanged)
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
func (it *KiwiAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiAdminChanged represents a AdminChanged event raised by the Kiwi contract.
type KiwiAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Kiwi *KiwiFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*KiwiAdminChangedIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &KiwiAdminChangedIterator{contract: _Kiwi.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Kiwi *KiwiFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *KiwiAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiAdminChanged)
				if err := _Kiwi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseAdminChanged(log types.Log) (*KiwiAdminChanged, error) {
	event := new(KiwiAdminChanged)
	if err := _Kiwi.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Kiwi contract.
type KiwiApprovalIterator struct {
	Event *KiwiApproval // Event containing the contract specifics and raw log

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
func (it *KiwiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiApproval)
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
		it.Event = new(KiwiApproval)
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
func (it *KiwiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiApproval represents a Approval event raised by the Kiwi contract.
type KiwiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kiwi *KiwiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*KiwiApprovalIterator, error) {

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

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KiwiApprovalIterator{contract: _Kiwi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Kiwi *KiwiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KiwiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiApproval)
				if err := _Kiwi.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseApproval(log types.Log) (*KiwiApproval, error) {
	event := new(KiwiApproval)
	if err := _Kiwi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Kiwi contract.
type KiwiApprovalForAllIterator struct {
	Event *KiwiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *KiwiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiApprovalForAll)
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
		it.Event = new(KiwiApprovalForAll)
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
func (it *KiwiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiApprovalForAll represents a ApprovalForAll event raised by the Kiwi contract.
type KiwiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kiwi *KiwiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*KiwiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &KiwiApprovalForAllIterator{contract: _Kiwi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Kiwi *KiwiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *KiwiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiApprovalForAll)
				if err := _Kiwi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseApprovalForAll(log types.Log) (*KiwiApprovalForAll, error) {
	event := new(KiwiApprovalForAll)
	if err := _Kiwi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Kiwi contract.
type KiwiBatchMetadataUpdateIterator struct {
	Event *KiwiBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *KiwiBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiBatchMetadataUpdate)
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
		it.Event = new(KiwiBatchMetadataUpdate)
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
func (it *KiwiBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Kiwi contract.
type KiwiBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Kiwi *KiwiFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*KiwiBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &KiwiBatchMetadataUpdateIterator{contract: _Kiwi.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Kiwi *KiwiFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *KiwiBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiBatchMetadataUpdate)
				if err := _Kiwi.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Kiwi *KiwiFilterer) ParseBatchMetadataUpdate(log types.Log) (*KiwiBatchMetadataUpdate, error) {
	event := new(KiwiBatchMetadataUpdate)
	if err := _Kiwi.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Kiwi contract.
type KiwiBeaconUpgradedIterator struct {
	Event *KiwiBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *KiwiBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiBeaconUpgraded)
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
		it.Event = new(KiwiBeaconUpgraded)
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
func (it *KiwiBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiBeaconUpgraded represents a BeaconUpgraded event raised by the Kiwi contract.
type KiwiBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Kiwi *KiwiFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*KiwiBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &KiwiBeaconUpgradedIterator{contract: _Kiwi.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Kiwi *KiwiFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *KiwiBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiBeaconUpgraded)
				if err := _Kiwi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseBeaconUpgraded(log types.Log) (*KiwiBeaconUpgraded, error) {
	event := new(KiwiBeaconUpgraded)
	if err := _Kiwi.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiFundsReceivedIterator is returned from FilterFundsReceived and is used to iterate over the raw logs and unpacked data for FundsReceived events raised by the Kiwi contract.
type KiwiFundsReceivedIterator struct {
	Event *KiwiFundsReceived // Event containing the contract specifics and raw log

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
func (it *KiwiFundsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiFundsReceived)
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
		it.Event = new(KiwiFundsReceived)
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
func (it *KiwiFundsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiFundsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiFundsReceived represents a FundsReceived event raised by the Kiwi contract.
type KiwiFundsReceived struct {
	Source common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReceived is a free log retrieval operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed source, uint256 amount)
func (_Kiwi *KiwiFilterer) FilterFundsReceived(opts *bind.FilterOpts, source []common.Address) (*KiwiFundsReceivedIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "FundsReceived", sourceRule)
	if err != nil {
		return nil, err
	}
	return &KiwiFundsReceivedIterator{contract: _Kiwi.contract, event: "FundsReceived", logs: logs, sub: sub}, nil
}

// WatchFundsReceived is a free log subscription operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed source, uint256 amount)
func (_Kiwi *KiwiFilterer) WatchFundsReceived(opts *bind.WatchOpts, sink chan<- *KiwiFundsReceived, source []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "FundsReceived", sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiFundsReceived)
				if err := _Kiwi.contract.UnpackLog(event, "FundsReceived", log); err != nil {
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

// ParseFundsReceived is a log parse operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed source, uint256 amount)
func (_Kiwi *KiwiFilterer) ParseFundsReceived(log types.Log) (*KiwiFundsReceived, error) {
	event := new(KiwiFundsReceived)
	if err := _Kiwi.contract.UnpackLog(event, "FundsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiFundsRecipientChangedIterator is returned from FilterFundsRecipientChanged and is used to iterate over the raw logs and unpacked data for FundsRecipientChanged events raised by the Kiwi contract.
type KiwiFundsRecipientChangedIterator struct {
	Event *KiwiFundsRecipientChanged // Event containing the contract specifics and raw log

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
func (it *KiwiFundsRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiFundsRecipientChanged)
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
		it.Event = new(KiwiFundsRecipientChanged)
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
func (it *KiwiFundsRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiFundsRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiFundsRecipientChanged represents a FundsRecipientChanged event raised by the Kiwi contract.
type KiwiFundsRecipientChanged struct {
	NewAddress common.Address
	ChangedBy  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundsRecipientChanged is a free log retrieval operation binding the contract event 0x70a7ea5c664ab9c21baf3da59bb2f1e1ca33557b08a0031fab4f170767449951.
//
// Solidity: event FundsRecipientChanged(address indexed newAddress, address indexed changedBy)
func (_Kiwi *KiwiFilterer) FilterFundsRecipientChanged(opts *bind.FilterOpts, newAddress []common.Address, changedBy []common.Address) (*KiwiFundsRecipientChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "FundsRecipientChanged", newAddressRule, changedByRule)
	if err != nil {
		return nil, err
	}
	return &KiwiFundsRecipientChangedIterator{contract: _Kiwi.contract, event: "FundsRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchFundsRecipientChanged is a free log subscription operation binding the contract event 0x70a7ea5c664ab9c21baf3da59bb2f1e1ca33557b08a0031fab4f170767449951.
//
// Solidity: event FundsRecipientChanged(address indexed newAddress, address indexed changedBy)
func (_Kiwi *KiwiFilterer) WatchFundsRecipientChanged(opts *bind.WatchOpts, sink chan<- *KiwiFundsRecipientChanged, newAddress []common.Address, changedBy []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "FundsRecipientChanged", newAddressRule, changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiFundsRecipientChanged)
				if err := _Kiwi.contract.UnpackLog(event, "FundsRecipientChanged", log); err != nil {
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

// ParseFundsRecipientChanged is a log parse operation binding the contract event 0x70a7ea5c664ab9c21baf3da59bb2f1e1ca33557b08a0031fab4f170767449951.
//
// Solidity: event FundsRecipientChanged(address indexed newAddress, address indexed changedBy)
func (_Kiwi *KiwiFilterer) ParseFundsRecipientChanged(log types.Log) (*KiwiFundsRecipientChanged, error) {
	event := new(KiwiFundsRecipientChanged)
	if err := _Kiwi.contract.UnpackLog(event, "FundsRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Kiwi contract.
type KiwiFundsWithdrawnIterator struct {
	Event *KiwiFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *KiwiFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiFundsWithdrawn)
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
		it.Event = new(KiwiFundsWithdrawn)
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
func (it *KiwiFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiFundsWithdrawn represents a FundsWithdrawn event raised by the Kiwi contract.
type KiwiFundsWithdrawn struct {
	WithdrawnBy  common.Address
	WithdrawnTo  common.Address
	Amount       *big.Int
	FeeRecipient common.Address
	FeeAmount    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0x8a95554e4c9dcaaf33f247387f2ee77390780487d3365e3a804788791a1df500.
//
// Solidity: event FundsWithdrawn(address indexed withdrawnBy, address indexed withdrawnTo, uint256 amount, address feeRecipient, uint256 feeAmount)
func (_Kiwi *KiwiFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, withdrawnBy []common.Address, withdrawnTo []common.Address) (*KiwiFundsWithdrawnIterator, error) {

	var withdrawnByRule []interface{}
	for _, withdrawnByItem := range withdrawnBy {
		withdrawnByRule = append(withdrawnByRule, withdrawnByItem)
	}
	var withdrawnToRule []interface{}
	for _, withdrawnToItem := range withdrawnTo {
		withdrawnToRule = append(withdrawnToRule, withdrawnToItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "FundsWithdrawn", withdrawnByRule, withdrawnToRule)
	if err != nil {
		return nil, err
	}
	return &KiwiFundsWithdrawnIterator{contract: _Kiwi.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0x8a95554e4c9dcaaf33f247387f2ee77390780487d3365e3a804788791a1df500.
//
// Solidity: event FundsWithdrawn(address indexed withdrawnBy, address indexed withdrawnTo, uint256 amount, address feeRecipient, uint256 feeAmount)
func (_Kiwi *KiwiFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *KiwiFundsWithdrawn, withdrawnBy []common.Address, withdrawnTo []common.Address) (event.Subscription, error) {

	var withdrawnByRule []interface{}
	for _, withdrawnByItem := range withdrawnBy {
		withdrawnByRule = append(withdrawnByRule, withdrawnByItem)
	}
	var withdrawnToRule []interface{}
	for _, withdrawnToItem := range withdrawnTo {
		withdrawnToRule = append(withdrawnToRule, withdrawnToItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "FundsWithdrawn", withdrawnByRule, withdrawnToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiFundsWithdrawn)
				if err := _Kiwi.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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

// ParseFundsWithdrawn is a log parse operation binding the contract event 0x8a95554e4c9dcaaf33f247387f2ee77390780487d3365e3a804788791a1df500.
//
// Solidity: event FundsWithdrawn(address indexed withdrawnBy, address indexed withdrawnTo, uint256 amount, address feeRecipient, uint256 feeAmount)
func (_Kiwi *KiwiFilterer) ParseFundsWithdrawn(log types.Log) (*KiwiFundsWithdrawn, error) {
	event := new(KiwiFundsWithdrawn)
	if err := _Kiwi.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Kiwi contract.
type KiwiMetadataUpdateIterator struct {
	Event *KiwiMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *KiwiMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiMetadataUpdate)
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
		it.Event = new(KiwiMetadataUpdate)
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
func (it *KiwiMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiMetadataUpdate represents a MetadataUpdate event raised by the Kiwi contract.
type KiwiMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Kiwi *KiwiFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*KiwiMetadataUpdateIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &KiwiMetadataUpdateIterator{contract: _Kiwi.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Kiwi *KiwiFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *KiwiMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiMetadataUpdate)
				if err := _Kiwi.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Kiwi *KiwiFilterer) ParseMetadataUpdate(log types.Log) (*KiwiMetadataUpdate, error) {
	event := new(KiwiMetadataUpdate)
	if err := _Kiwi.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiMintCommentIterator is returned from FilterMintComment and is used to iterate over the raw logs and unpacked data for MintComment events raised by the Kiwi contract.
type KiwiMintCommentIterator struct {
	Event *KiwiMintComment // Event containing the contract specifics and raw log

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
func (it *KiwiMintCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiMintComment)
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
		it.Event = new(KiwiMintComment)
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
func (it *KiwiMintCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiMintCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiMintComment represents a MintComment event raised by the Kiwi contract.
type KiwiMintComment struct {
	Sender        common.Address
	TokenContract common.Address
	TokenId       *big.Int
	Quantity      *big.Int
	Comment       string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintComment is a free log retrieval operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed tokenContract, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Kiwi *KiwiFilterer) FilterMintComment(opts *bind.FilterOpts, sender []common.Address, tokenContract []common.Address, tokenId []*big.Int) (*KiwiMintCommentIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "MintComment", senderRule, tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KiwiMintCommentIterator{contract: _Kiwi.contract, event: "MintComment", logs: logs, sub: sub}, nil
}

// WatchMintComment is a free log subscription operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed tokenContract, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Kiwi *KiwiFilterer) WatchMintComment(opts *bind.WatchOpts, sink chan<- *KiwiMintComment, sender []common.Address, tokenContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "MintComment", senderRule, tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiMintComment)
				if err := _Kiwi.contract.UnpackLog(event, "MintComment", log); err != nil {
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

// ParseMintComment is a log parse operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed tokenContract, uint256 indexed tokenId, uint256 quantity, string comment)
func (_Kiwi *KiwiFilterer) ParseMintComment(log types.Log) (*KiwiMintComment, error) {
	event := new(KiwiMintComment)
	if err := _Kiwi.contract.UnpackLog(event, "MintComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiMintFeePayoutIterator is returned from FilterMintFeePayout and is used to iterate over the raw logs and unpacked data for MintFeePayout events raised by the Kiwi contract.
type KiwiMintFeePayoutIterator struct {
	Event *KiwiMintFeePayout // Event containing the contract specifics and raw log

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
func (it *KiwiMintFeePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiMintFeePayout)
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
		it.Event = new(KiwiMintFeePayout)
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
func (it *KiwiMintFeePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiMintFeePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiMintFeePayout represents a MintFeePayout event raised by the Kiwi contract.
type KiwiMintFeePayout struct {
	MintFeeAmount    *big.Int
	MintFeeRecipient common.Address
	Success          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMintFeePayout is a free log retrieval operation binding the contract event 0x6f8da53cfedb8cc4f7935c3629624e50b63053c93bb2cad246aa4d3a2ba7d4ce.
//
// Solidity: event MintFeePayout(uint256 mintFeeAmount, address mintFeeRecipient, bool success)
func (_Kiwi *KiwiFilterer) FilterMintFeePayout(opts *bind.FilterOpts) (*KiwiMintFeePayoutIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "MintFeePayout")
	if err != nil {
		return nil, err
	}
	return &KiwiMintFeePayoutIterator{contract: _Kiwi.contract, event: "MintFeePayout", logs: logs, sub: sub}, nil
}

// WatchMintFeePayout is a free log subscription operation binding the contract event 0x6f8da53cfedb8cc4f7935c3629624e50b63053c93bb2cad246aa4d3a2ba7d4ce.
//
// Solidity: event MintFeePayout(uint256 mintFeeAmount, address mintFeeRecipient, bool success)
func (_Kiwi *KiwiFilterer) WatchMintFeePayout(opts *bind.WatchOpts, sink chan<- *KiwiMintFeePayout) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "MintFeePayout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiMintFeePayout)
				if err := _Kiwi.contract.UnpackLog(event, "MintFeePayout", log); err != nil {
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

// ParseMintFeePayout is a log parse operation binding the contract event 0x6f8da53cfedb8cc4f7935c3629624e50b63053c93bb2cad246aa4d3a2ba7d4ce.
//
// Solidity: event MintFeePayout(uint256 mintFeeAmount, address mintFeeRecipient, bool success)
func (_Kiwi *KiwiFilterer) ParseMintFeePayout(log types.Log) (*KiwiMintFeePayout, error) {
	event := new(KiwiMintFeePayout)
	if err := _Kiwi.contract.UnpackLog(event, "MintFeePayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiOpenMintFinalizedIterator is returned from FilterOpenMintFinalized and is used to iterate over the raw logs and unpacked data for OpenMintFinalized events raised by the Kiwi contract.
type KiwiOpenMintFinalizedIterator struct {
	Event *KiwiOpenMintFinalized // Event containing the contract specifics and raw log

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
func (it *KiwiOpenMintFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiOpenMintFinalized)
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
		it.Event = new(KiwiOpenMintFinalized)
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
func (it *KiwiOpenMintFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiOpenMintFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiOpenMintFinalized represents a OpenMintFinalized event raised by the Kiwi contract.
type KiwiOpenMintFinalized struct {
	Sender        common.Address
	NumberOfMints *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenMintFinalized is a free log retrieval operation binding the contract event 0xb6cc1e876b8d7479c8afe025a58658b0f3c3ec5bd0f28cb4261326b162069bf8.
//
// Solidity: event OpenMintFinalized(address indexed sender, uint256 numberOfMints)
func (_Kiwi *KiwiFilterer) FilterOpenMintFinalized(opts *bind.FilterOpts, sender []common.Address) (*KiwiOpenMintFinalizedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "OpenMintFinalized", senderRule)
	if err != nil {
		return nil, err
	}
	return &KiwiOpenMintFinalizedIterator{contract: _Kiwi.contract, event: "OpenMintFinalized", logs: logs, sub: sub}, nil
}

// WatchOpenMintFinalized is a free log subscription operation binding the contract event 0xb6cc1e876b8d7479c8afe025a58658b0f3c3ec5bd0f28cb4261326b162069bf8.
//
// Solidity: event OpenMintFinalized(address indexed sender, uint256 numberOfMints)
func (_Kiwi *KiwiFilterer) WatchOpenMintFinalized(opts *bind.WatchOpts, sink chan<- *KiwiOpenMintFinalized, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "OpenMintFinalized", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiOpenMintFinalized)
				if err := _Kiwi.contract.UnpackLog(event, "OpenMintFinalized", log); err != nil {
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

// ParseOpenMintFinalized is a log parse operation binding the contract event 0xb6cc1e876b8d7479c8afe025a58658b0f3c3ec5bd0f28cb4261326b162069bf8.
//
// Solidity: event OpenMintFinalized(address indexed sender, uint256 numberOfMints)
func (_Kiwi *KiwiFilterer) ParseOpenMintFinalized(log types.Log) (*KiwiOpenMintFinalized, error) {
	event := new(KiwiOpenMintFinalized)
	if err := _Kiwi.contract.UnpackLog(event, "OpenMintFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiOwnerCanceledIterator is returned from FilterOwnerCanceled and is used to iterate over the raw logs and unpacked data for OwnerCanceled events raised by the Kiwi contract.
type KiwiOwnerCanceledIterator struct {
	Event *KiwiOwnerCanceled // Event containing the contract specifics and raw log

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
func (it *KiwiOwnerCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiOwnerCanceled)
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
		it.Event = new(KiwiOwnerCanceled)
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
func (it *KiwiOwnerCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiOwnerCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiOwnerCanceled represents a OwnerCanceled event raised by the Kiwi contract.
type KiwiOwnerCanceled struct {
	PreviousOwner     common.Address
	PotentialNewOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOwnerCanceled is a free log retrieval operation binding the contract event 0x682679deecef4dcd49674845cc1e3a075fea9073680aa445a8207d5a4bdea3da.
//
// Solidity: event OwnerCanceled(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) FilterOwnerCanceled(opts *bind.FilterOpts, previousOwner []common.Address, potentialNewOwner []common.Address) (*KiwiOwnerCanceledIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "OwnerCanceled", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KiwiOwnerCanceledIterator{contract: _Kiwi.contract, event: "OwnerCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnerCanceled is a free log subscription operation binding the contract event 0x682679deecef4dcd49674845cc1e3a075fea9073680aa445a8207d5a4bdea3da.
//
// Solidity: event OwnerCanceled(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) WatchOwnerCanceled(opts *bind.WatchOpts, sink chan<- *KiwiOwnerCanceled, previousOwner []common.Address, potentialNewOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "OwnerCanceled", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiOwnerCanceled)
				if err := _Kiwi.contract.UnpackLog(event, "OwnerCanceled", log); err != nil {
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

// ParseOwnerCanceled is a log parse operation binding the contract event 0x682679deecef4dcd49674845cc1e3a075fea9073680aa445a8207d5a4bdea3da.
//
// Solidity: event OwnerCanceled(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) ParseOwnerCanceled(log types.Log) (*KiwiOwnerCanceled, error) {
	event := new(KiwiOwnerCanceled)
	if err := _Kiwi.contract.UnpackLog(event, "OwnerCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiOwnerPendingIterator is returned from FilterOwnerPending and is used to iterate over the raw logs and unpacked data for OwnerPending events raised by the Kiwi contract.
type KiwiOwnerPendingIterator struct {
	Event *KiwiOwnerPending // Event containing the contract specifics and raw log

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
func (it *KiwiOwnerPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiOwnerPending)
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
		it.Event = new(KiwiOwnerPending)
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
func (it *KiwiOwnerPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiOwnerPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiOwnerPending represents a OwnerPending event raised by the Kiwi contract.
type KiwiOwnerPending struct {
	PreviousOwner     common.Address
	PotentialNewOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOwnerPending is a free log retrieval operation binding the contract event 0x4f2638f5949b9614ef8d5e268cb51348ad7f434a34812bf64b6e95014fbd357e.
//
// Solidity: event OwnerPending(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) FilterOwnerPending(opts *bind.FilterOpts, previousOwner []common.Address, potentialNewOwner []common.Address) (*KiwiOwnerPendingIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "OwnerPending", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KiwiOwnerPendingIterator{contract: _Kiwi.contract, event: "OwnerPending", logs: logs, sub: sub}, nil
}

// WatchOwnerPending is a free log subscription operation binding the contract event 0x4f2638f5949b9614ef8d5e268cb51348ad7f434a34812bf64b6e95014fbd357e.
//
// Solidity: event OwnerPending(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) WatchOwnerPending(opts *bind.WatchOpts, sink chan<- *KiwiOwnerPending, previousOwner []common.Address, potentialNewOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "OwnerPending", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiOwnerPending)
				if err := _Kiwi.contract.UnpackLog(event, "OwnerPending", log); err != nil {
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

// ParseOwnerPending is a log parse operation binding the contract event 0x4f2638f5949b9614ef8d5e268cb51348ad7f434a34812bf64b6e95014fbd357e.
//
// Solidity: event OwnerPending(address indexed previousOwner, address indexed potentialNewOwner)
func (_Kiwi *KiwiFilterer) ParseOwnerPending(log types.Log) (*KiwiOwnerPending, error) {
	event := new(KiwiOwnerPending)
	if err := _Kiwi.contract.UnpackLog(event, "OwnerPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Kiwi contract.
type KiwiOwnershipTransferredIterator struct {
	Event *KiwiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KiwiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiOwnershipTransferred)
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
		it.Event = new(KiwiOwnershipTransferred)
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
func (it *KiwiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiOwnershipTransferred represents a OwnershipTransferred event raised by the Kiwi contract.
type KiwiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kiwi *KiwiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KiwiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KiwiOwnershipTransferredIterator{contract: _Kiwi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Kiwi *KiwiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KiwiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiOwnershipTransferred)
				if err := _Kiwi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseOwnershipTransferred(log types.Log) (*KiwiOwnershipTransferred, error) {
	event := new(KiwiOwnershipTransferred)
	if err := _Kiwi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Kiwi contract.
type KiwiRoleAdminChangedIterator struct {
	Event *KiwiRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KiwiRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiRoleAdminChanged)
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
		it.Event = new(KiwiRoleAdminChanged)
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
func (it *KiwiRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiRoleAdminChanged represents a RoleAdminChanged event raised by the Kiwi contract.
type KiwiRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Kiwi *KiwiFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KiwiRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KiwiRoleAdminChangedIterator{contract: _Kiwi.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Kiwi *KiwiFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KiwiRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiRoleAdminChanged)
				if err := _Kiwi.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Kiwi *KiwiFilterer) ParseRoleAdminChanged(log types.Log) (*KiwiRoleAdminChanged, error) {
	event := new(KiwiRoleAdminChanged)
	if err := _Kiwi.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Kiwi contract.
type KiwiRoleGrantedIterator struct {
	Event *KiwiRoleGranted // Event containing the contract specifics and raw log

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
func (it *KiwiRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiRoleGranted)
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
		it.Event = new(KiwiRoleGranted)
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
func (it *KiwiRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiRoleGranted represents a RoleGranted event raised by the Kiwi contract.
type KiwiRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KiwiRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KiwiRoleGrantedIterator{contract: _Kiwi.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KiwiRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiRoleGranted)
				if err := _Kiwi.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) ParseRoleGranted(log types.Log) (*KiwiRoleGranted, error) {
	event := new(KiwiRoleGranted)
	if err := _Kiwi.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Kiwi contract.
type KiwiRoleRevokedIterator struct {
	Event *KiwiRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KiwiRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiRoleRevoked)
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
		it.Event = new(KiwiRoleRevoked)
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
func (it *KiwiRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiRoleRevoked represents a RoleRevoked event raised by the Kiwi contract.
type KiwiRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KiwiRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KiwiRoleRevokedIterator{contract: _Kiwi.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KiwiRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiRoleRevoked)
				if err := _Kiwi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Kiwi *KiwiFilterer) ParseRoleRevoked(log types.Log) (*KiwiRoleRevoked, error) {
	event := new(KiwiRoleRevoked)
	if err := _Kiwi.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiSaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the Kiwi contract.
type KiwiSaleIterator struct {
	Event *KiwiSale // Event containing the contract specifics and raw log

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
func (it *KiwiSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiSale)
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
		it.Event = new(KiwiSale)
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
func (it *KiwiSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiSale represents a Sale event raised by the Kiwi contract.
type KiwiSale struct {
	To                    common.Address
	Quantity              *big.Int
	PricePerToken         *big.Int
	FirstPurchasedTokenId *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90.
//
// Solidity: event Sale(address indexed to, uint256 indexed quantity, uint256 indexed pricePerToken, uint256 firstPurchasedTokenId)
func (_Kiwi *KiwiFilterer) FilterSale(opts *bind.FilterOpts, to []common.Address, quantity []*big.Int, pricePerToken []*big.Int) (*KiwiSaleIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var quantityRule []interface{}
	for _, quantityItem := range quantity {
		quantityRule = append(quantityRule, quantityItem)
	}
	var pricePerTokenRule []interface{}
	for _, pricePerTokenItem := range pricePerToken {
		pricePerTokenRule = append(pricePerTokenRule, pricePerTokenItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Sale", toRule, quantityRule, pricePerTokenRule)
	if err != nil {
		return nil, err
	}
	return &KiwiSaleIterator{contract: _Kiwi.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90.
//
// Solidity: event Sale(address indexed to, uint256 indexed quantity, uint256 indexed pricePerToken, uint256 firstPurchasedTokenId)
func (_Kiwi *KiwiFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *KiwiSale, to []common.Address, quantity []*big.Int, pricePerToken []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var quantityRule []interface{}
	for _, quantityItem := range quantity {
		quantityRule = append(quantityRule, quantityItem)
	}
	var pricePerTokenRule []interface{}
	for _, pricePerTokenItem := range pricePerToken {
		pricePerTokenRule = append(pricePerTokenRule, pricePerTokenItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Sale", toRule, quantityRule, pricePerTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiSale)
				if err := _Kiwi.contract.UnpackLog(event, "Sale", log); err != nil {
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

// ParseSale is a log parse operation binding the contract event 0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90.
//
// Solidity: event Sale(address indexed to, uint256 indexed quantity, uint256 indexed pricePerToken, uint256 firstPurchasedTokenId)
func (_Kiwi *KiwiFilterer) ParseSale(log types.Log) (*KiwiSale, error) {
	event := new(KiwiSale)
	if err := _Kiwi.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiSalesConfigChangedIterator is returned from FilterSalesConfigChanged and is used to iterate over the raw logs and unpacked data for SalesConfigChanged events raised by the Kiwi contract.
type KiwiSalesConfigChangedIterator struct {
	Event *KiwiSalesConfigChanged // Event containing the contract specifics and raw log

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
func (it *KiwiSalesConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiSalesConfigChanged)
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
		it.Event = new(KiwiSalesConfigChanged)
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
func (it *KiwiSalesConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiSalesConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiSalesConfigChanged represents a SalesConfigChanged event raised by the Kiwi contract.
type KiwiSalesConfigChanged struct {
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSalesConfigChanged is a free log retrieval operation binding the contract event 0xc1ff5e4744ac8dd2b8027a10e3723b165975297501c71c4e7dcb8796d96375db.
//
// Solidity: event SalesConfigChanged(address indexed changedBy)
func (_Kiwi *KiwiFilterer) FilterSalesConfigChanged(opts *bind.FilterOpts, changedBy []common.Address) (*KiwiSalesConfigChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "SalesConfigChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &KiwiSalesConfigChangedIterator{contract: _Kiwi.contract, event: "SalesConfigChanged", logs: logs, sub: sub}, nil
}

// WatchSalesConfigChanged is a free log subscription operation binding the contract event 0xc1ff5e4744ac8dd2b8027a10e3723b165975297501c71c4e7dcb8796d96375db.
//
// Solidity: event SalesConfigChanged(address indexed changedBy)
func (_Kiwi *KiwiFilterer) WatchSalesConfigChanged(opts *bind.WatchOpts, sink chan<- *KiwiSalesConfigChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "SalesConfigChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiSalesConfigChanged)
				if err := _Kiwi.contract.UnpackLog(event, "SalesConfigChanged", log); err != nil {
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

// ParseSalesConfigChanged is a log parse operation binding the contract event 0xc1ff5e4744ac8dd2b8027a10e3723b165975297501c71c4e7dcb8796d96375db.
//
// Solidity: event SalesConfigChanged(address indexed changedBy)
func (_Kiwi *KiwiFilterer) ParseSalesConfigChanged(log types.Log) (*KiwiSalesConfigChanged, error) {
	event := new(KiwiSalesConfigChanged)
	if err := _Kiwi.contract.UnpackLog(event, "SalesConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Kiwi contract.
type KiwiTransferIterator struct {
	Event *KiwiTransfer // Event containing the contract specifics and raw log

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
func (it *KiwiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiTransfer)
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
		it.Event = new(KiwiTransfer)
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
func (it *KiwiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiTransfer represents a Transfer event raised by the Kiwi contract.
type KiwiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kiwi *KiwiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*KiwiTransferIterator, error) {

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

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &KiwiTransferIterator{contract: _Kiwi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Kiwi *KiwiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KiwiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiTransfer)
				if err := _Kiwi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseTransfer(log types.Log) (*KiwiTransfer, error) {
	event := new(KiwiTransfer)
	if err := _Kiwi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiUpdatedMetadataRendererIterator is returned from FilterUpdatedMetadataRenderer and is used to iterate over the raw logs and unpacked data for UpdatedMetadataRenderer events raised by the Kiwi contract.
type KiwiUpdatedMetadataRendererIterator struct {
	Event *KiwiUpdatedMetadataRenderer // Event containing the contract specifics and raw log

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
func (it *KiwiUpdatedMetadataRendererIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiUpdatedMetadataRenderer)
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
		it.Event = new(KiwiUpdatedMetadataRenderer)
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
func (it *KiwiUpdatedMetadataRendererIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiUpdatedMetadataRendererIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiUpdatedMetadataRenderer represents a UpdatedMetadataRenderer event raised by the Kiwi contract.
type KiwiUpdatedMetadataRenderer struct {
	Sender   common.Address
	Renderer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMetadataRenderer is a free log retrieval operation binding the contract event 0x046c5d913c35948c3e0e44c3599eb14bf33b73f141fa8bb282b300414998b868.
//
// Solidity: event UpdatedMetadataRenderer(address sender, address renderer)
func (_Kiwi *KiwiFilterer) FilterUpdatedMetadataRenderer(opts *bind.FilterOpts) (*KiwiUpdatedMetadataRendererIterator, error) {

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "UpdatedMetadataRenderer")
	if err != nil {
		return nil, err
	}
	return &KiwiUpdatedMetadataRendererIterator{contract: _Kiwi.contract, event: "UpdatedMetadataRenderer", logs: logs, sub: sub}, nil
}

// WatchUpdatedMetadataRenderer is a free log subscription operation binding the contract event 0x046c5d913c35948c3e0e44c3599eb14bf33b73f141fa8bb282b300414998b868.
//
// Solidity: event UpdatedMetadataRenderer(address sender, address renderer)
func (_Kiwi *KiwiFilterer) WatchUpdatedMetadataRenderer(opts *bind.WatchOpts, sink chan<- *KiwiUpdatedMetadataRenderer) (event.Subscription, error) {

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "UpdatedMetadataRenderer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiUpdatedMetadataRenderer)
				if err := _Kiwi.contract.UnpackLog(event, "UpdatedMetadataRenderer", log); err != nil {
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

// ParseUpdatedMetadataRenderer is a log parse operation binding the contract event 0x046c5d913c35948c3e0e44c3599eb14bf33b73f141fa8bb282b300414998b868.
//
// Solidity: event UpdatedMetadataRenderer(address sender, address renderer)
func (_Kiwi *KiwiFilterer) ParseUpdatedMetadataRenderer(log types.Log) (*KiwiUpdatedMetadataRenderer, error) {
	event := new(KiwiUpdatedMetadataRenderer)
	if err := _Kiwi.contract.UnpackLog(event, "UpdatedMetadataRenderer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KiwiUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Kiwi contract.
type KiwiUpgradedIterator struct {
	Event *KiwiUpgraded // Event containing the contract specifics and raw log

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
func (it *KiwiUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KiwiUpgraded)
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
		it.Event = new(KiwiUpgraded)
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
func (it *KiwiUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KiwiUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KiwiUpgraded represents a Upgraded event raised by the Kiwi contract.
type KiwiUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Kiwi *KiwiFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*KiwiUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Kiwi.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &KiwiUpgradedIterator{contract: _Kiwi.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Kiwi *KiwiFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *KiwiUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Kiwi.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KiwiUpgraded)
				if err := _Kiwi.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Kiwi *KiwiFilterer) ParseUpgraded(log types.Log) (*KiwiUpgraded, error) {
	event := new(KiwiUpgraded)
	if err := _Kiwi.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
