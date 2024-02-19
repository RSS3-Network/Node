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

// ProtocolRewardsMetaData contains all meta data concerning the ProtocolRewards contract.
var ProtocolRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_zoraERC721TransferHelper\",\"type\":\"address\"},{\"internalType\":\"contractIFactoryUpgradeGate\",\"name\":\"_factoryUpgradeGate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketFilterDAOAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mintFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_mintFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_protocolRewards\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"Access_MissingRoleOrAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Access_OnlyAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Access_WithdrawNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"Admin_InvalidUpgradeAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Admin_UnableToFinalizeNotOpenEdition\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CREATOR_FUNDS_RECIPIENT_NOT_SET\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExternalMetadataRenderer_CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ADDRESS_ZERO\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"INVALID_ETH_AMOUNT\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMintSchedule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketFilterDAOAddressNotSupportedForChain\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintFee_FundsSendFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Mint_SoldOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_CREATE_REFERRAL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_OWNER\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ONLY_PENDING_OWNER\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"OperatorNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_Inactive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_MerkleNotApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Presale_TooManyForAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolRewards_WithdrawSendFailure\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Purchase_TooManyForAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"correctPrice\",\"type\":\"uint256\"}],\"name\":\"Purchase_WrongPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemoteOperatorFilterRegistryCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Sale_Inactive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"maxRoyaltyBPS\",\"type\":\"uint16\"}],\"name\":\"Setup_RoyaltyPercentageTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Withdraw_FundsSendFailure\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"FundsRecipientChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawnBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawnTo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"MintComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"mintFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"MintFeePayout\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"numberOfMints\",\"type\":\"uint256\"}],\"name\":\"OpenMintFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"potentialNewOwner\",\"type\":\"address\"}],\"name\":\"OwnerCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"potentialNewOwner\",\"type\":\"address\"}],\"name\":\"OwnerPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"firstPurchasedTokenId\",\"type\":\"uint256\"}],\"name\":\"Sale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"}],\"name\":\"SalesConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractIMetadataRenderer\",\"name\":\"renderer\",\"type\":\"address\"}],\"name\":\"UpdatedMetadataRenderer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SALES_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"adminMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"}],\"name\":\"adminMintAirdrop\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"callMetadataRenderer\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computeFreeMintRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computePaidMintRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"creatorReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintReferralReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"firstMinterReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zoraReward\",\"type\":\"uint256\"}],\"internalType\":\"structRewardsSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"computeTotalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"metadataRenderer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"editionSize\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"royaltyBPS\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createReferral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factoryUpgradeGate\",\"outputs\":[{\"internalType\":\"contractIFactoryUpgradeGate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeOpenEdition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_contractName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractSymbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_fundsRecipient\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_editionSize\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"_royaltyBPS\",\"type\":\"uint16\"},{\"internalType\":\"bytes[]\",\"name\":\"_setupCalls\",\"type\":\"bytes[]\"},{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"_metadataRenderer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_metadataRendererInit\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_createReferral\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"manageMarketFilterDAOSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketFilterDAOAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataRenderer\",\"outputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"}],\"name\":\"mintWithRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"mintedPerAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalMints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"presaleMints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"publicMints\",\"type\":\"uint256\"}],\"internalType\":\"structIERC721Drop.AddressMintDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleMintsByAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"purchase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"purchasePresale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchasePresaleWithComment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mintReferral\",\"type\":\"address\"}],\"name\":\"purchasePresaleWithRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchaseWithComment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"purchaseWithRecipient\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyMintSchedule\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"publicSaleActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"presaleActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"publicSalePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalMinted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"}],\"internalType\":\"structIERC721Drop.SaleDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"salesConfig\",\"outputs\":[{\"internalType\":\"uint104\",\"name\":\"publicSalePrice\",\"type\":\"uint104\"},{\"internalType\":\"uint32\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newRecipientAddress\",\"type\":\"address\"}],\"name\":\"setFundsRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMetadataRenderer\",\"name\":\"newRenderer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"setupRenderer\",\"type\":\"bytes\"}],\"name\":\"setMetadataRenderer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint104\",\"name\":\"publicSalePrice\",\"type\":\"uint104\"},{\"internalType\":\"uint32\",\"name\":\"maxSalePurchasePerAddress\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"publicSaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleStart\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"presaleEnd\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"presaleMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"setSaleConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"updateCreateReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"args\",\"type\":\"bytes\"}],\"name\":\"updateMarketFilterSettings\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newSchedule\",\"type\":\"uint32\"}],\"name\":\"updateRoyaltyMintSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoraERC721TransferHelper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"zoraFeeForAmount\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ProtocolRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolRewardsMetaData.ABI instead.
var ProtocolRewardsABI = ProtocolRewardsMetaData.ABI

// ProtocolRewards is an auto generated Go binding around an Ethereum contract.
type ProtocolRewards struct {
	ProtocolRewardsCaller     // Read-only binding to the contract
	ProtocolRewardsTransactor // Write-only binding to the contract
	ProtocolRewardsFilterer   // Log filterer for contract events
}

// ProtocolRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolRewardsSession struct {
	Contract     *ProtocolRewards  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolRewardsCallerSession struct {
	Contract *ProtocolRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ProtocolRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolRewardsTransactorSession struct {
	Contract     *ProtocolRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ProtocolRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolRewardsRaw struct {
	Contract *ProtocolRewards // Generic contract binding to access the raw methods on
}

// ProtocolRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolRewardsCallerRaw struct {
	Contract *ProtocolRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolRewardsTransactorRaw struct {
	Contract *ProtocolRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolRewards creates a new instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewards(address common.Address, backend bind.ContractBackend) (*ProtocolRewards, error) {
	contract, err := bindProtocolRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewards{ProtocolRewardsCaller: ProtocolRewardsCaller{contract: contract}, ProtocolRewardsTransactor: ProtocolRewardsTransactor{contract: contract}, ProtocolRewardsFilterer: ProtocolRewardsFilterer{contract: contract}}, nil
}

// NewProtocolRewardsCaller creates a new read-only instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsCaller(address common.Address, caller bind.ContractCaller) (*ProtocolRewardsCaller, error) {
	contract, err := bindProtocolRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsCaller{contract: contract}, nil
}

// NewProtocolRewardsTransactor creates a new write-only instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolRewardsTransactor, error) {
	contract, err := bindProtocolRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsTransactor{contract: contract}, nil
}

// NewProtocolRewardsFilterer creates a new log filterer instance of ProtocolRewards, bound to a specific deployed contract.
func NewProtocolRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolRewardsFilterer, error) {
	contract, err := bindProtocolRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsFilterer{contract: contract}, nil
}

// bindProtocolRewards binds a generic wrapper to an already deployed contract.
func bindProtocolRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolRewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRewards *ProtocolRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRewards.Contract.ProtocolRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRewards *ProtocolRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ProtocolRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRewards *ProtocolRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ProtocolRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolRewards *ProtocolRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolRewards *ProtocolRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolRewards *ProtocolRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.DEFAULTADMINROLE(&_ProtocolRewards.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.DEFAULTADMINROLE(&_ProtocolRewards.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) MINTERROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.MINTERROLE(&_ProtocolRewards.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) MINTERROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.MINTERROLE(&_ProtocolRewards.CallOpts)
}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) SALESMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "SALES_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) SALESMANAGERROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.SALESMANAGERROLE(&_ProtocolRewards.CallOpts)
}

// SALESMANAGERROLE is a free data retrieval call binding the contract method 0xe26bd343.
//
// Solidity: function SALES_MANAGER_ROLE() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) SALESMANAGERROLE() ([32]byte, error) {
	return _ProtocolRewards.Contract.SALESMANAGERROLE(&_ProtocolRewards.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.BalanceOf(&_ProtocolRewards.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.BalanceOf(&_ProtocolRewards.CallOpts, owner)
}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCaller) ComputeFreeMintRewards(opts *bind.CallOpts, numTokens *big.Int) (RewardsSettings, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "computeFreeMintRewards", numTokens)

	if err != nil {
		return *new(RewardsSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardsSettings)).(*RewardsSettings)

	return out0, err

}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsSession) ComputeFreeMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _ProtocolRewards.Contract.ComputeFreeMintRewards(&_ProtocolRewards.CallOpts, numTokens)
}

// ComputeFreeMintRewards is a free data retrieval call binding the contract method 0x722933f7.
//
// Solidity: function computeFreeMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCallerSession) ComputeFreeMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _ProtocolRewards.Contract.ComputeFreeMintRewards(&_ProtocolRewards.CallOpts, numTokens)
}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCaller) ComputePaidMintRewards(opts *bind.CallOpts, numTokens *big.Int) (RewardsSettings, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "computePaidMintRewards", numTokens)

	if err != nil {
		return *new(RewardsSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(RewardsSettings)).(*RewardsSettings)

	return out0, err

}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsSession) ComputePaidMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _ProtocolRewards.Contract.ComputePaidMintRewards(&_ProtocolRewards.CallOpts, numTokens)
}

// ComputePaidMintRewards is a free data retrieval call binding the contract method 0x5c046084.
//
// Solidity: function computePaidMintRewards(uint256 numTokens) pure returns((uint256,uint256,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCallerSession) ComputePaidMintRewards(numTokens *big.Int) (RewardsSettings, error) {
	return _ProtocolRewards.Contract.ComputePaidMintRewards(&_ProtocolRewards.CallOpts, numTokens)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) ComputeTotalReward(opts *bind.CallOpts, numTokens *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "computeTotalReward", numTokens)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) ComputeTotalReward(numTokens *big.Int) (*big.Int, error) {
	return _ProtocolRewards.Contract.ComputeTotalReward(&_ProtocolRewards.CallOpts, numTokens)
}

// ComputeTotalReward is a free data retrieval call binding the contract method 0x4132239b.
//
// Solidity: function computeTotalReward(uint256 numTokens) pure returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ComputeTotalReward(numTokens *big.Int) (*big.Int, error) {
	return _ProtocolRewards.Contract.ComputeTotalReward(&_ProtocolRewards.CallOpts, numTokens)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address metadataRenderer, uint64 editionSize, uint16 royaltyBPS, address fundsRecipient)
func (_ProtocolRewards *ProtocolRewardsCaller) Config(opts *bind.CallOpts) (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "config")

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
func (_ProtocolRewards *ProtocolRewardsSession) Config() (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	return _ProtocolRewards.Contract.Config(&_ProtocolRewards.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address metadataRenderer, uint64 editionSize, uint16 royaltyBPS, address fundsRecipient)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Config() (struct {
	MetadataRenderer common.Address
	EditionSize      uint64
	RoyaltyBPS       uint16
	FundsRecipient   common.Address
}, error) {
	return _ProtocolRewards.Contract.Config(&_ProtocolRewards.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ProtocolRewards *ProtocolRewardsSession) ContractURI() (string, error) {
	return _ProtocolRewards.Contract.ContractURI(&_ProtocolRewards.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ContractURI() (string, error) {
	return _ProtocolRewards.Contract.ContractURI(&_ProtocolRewards.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsCaller) ContractVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsSession) ContractVersion() (uint32, error) {
	return _ProtocolRewards.Contract.ContractVersion(&_ProtocolRewards.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ContractVersion() (uint32, error) {
	return _ProtocolRewards.Contract.ContractVersion(&_ProtocolRewards.CallOpts)
}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) CreateReferral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "createReferral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) CreateReferral() (common.Address, error) {
	return _ProtocolRewards.Contract.CreateReferral(&_ProtocolRewards.CallOpts)
}

// CreateReferral is a free data retrieval call binding the contract method 0x62bf43f0.
//
// Solidity: function createReferral() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) CreateReferral() (common.Address, error) {
	return _ProtocolRewards.Contract.CreateReferral(&_ProtocolRewards.CallOpts)
}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) FactoryUpgradeGate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "factoryUpgradeGate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) FactoryUpgradeGate() (common.Address, error) {
	return _ProtocolRewards.Contract.FactoryUpgradeGate(&_ProtocolRewards.CallOpts)
}

// FactoryUpgradeGate is a free data retrieval call binding the contract method 0x25eb54c6.
//
// Solidity: function factoryUpgradeGate() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) FactoryUpgradeGate() (common.Address, error) {
	return _ProtocolRewards.Contract.FactoryUpgradeGate(&_ProtocolRewards.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ProtocolRewards.Contract.GetApproved(&_ProtocolRewards.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ProtocolRewards.Contract.GetApproved(&_ProtocolRewards.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProtocolRewards.Contract.GetRoleAdmin(&_ProtocolRewards.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ProtocolRewards.Contract.GetRoleAdmin(&_ProtocolRewards.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProtocolRewards.Contract.HasRole(&_ProtocolRewards.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ProtocolRewards.Contract.HasRole(&_ProtocolRewards.CallOpts, role, account)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCaller) IsAdmin(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "isAdmin", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsSession) IsAdmin(user common.Address) (bool, error) {
	return _ProtocolRewards.Contract.IsAdmin(&_ProtocolRewards.CallOpts, user)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address user) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCallerSession) IsAdmin(user common.Address) (bool, error) {
	return _ProtocolRewards.Contract.IsAdmin(&_ProtocolRewards.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCaller) IsApprovedForAll(opts *bind.CallOpts, nftOwner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "isApprovedForAll", nftOwner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsSession) IsApprovedForAll(nftOwner common.Address, operator common.Address) (bool, error) {
	return _ProtocolRewards.Contract.IsApprovedForAll(&_ProtocolRewards.CallOpts, nftOwner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address nftOwner, address operator) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCallerSession) IsApprovedForAll(nftOwner common.Address, operator common.Address) (bool, error) {
	return _ProtocolRewards.Contract.IsApprovedForAll(&_ProtocolRewards.CallOpts, nftOwner, operator)
}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) MarketFilterDAOAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "marketFilterDAOAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) MarketFilterDAOAddress() (common.Address, error) {
	return _ProtocolRewards.Contract.MarketFilterDAOAddress(&_ProtocolRewards.CallOpts)
}

// MarketFilterDAOAddress is a free data retrieval call binding the contract method 0xce3ca396.
//
// Solidity: function marketFilterDAOAddress() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) MarketFilterDAOAddress() (common.Address, error) {
	return _ProtocolRewards.Contract.MarketFilterDAOAddress(&_ProtocolRewards.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) MetadataRenderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "metadataRenderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) MetadataRenderer() (common.Address, error) {
	return _ProtocolRewards.Contract.MetadataRenderer(&_ProtocolRewards.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) MetadataRenderer() (common.Address, error) {
	return _ProtocolRewards.Contract.MetadataRenderer(&_ProtocolRewards.CallOpts)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCaller) MintedPerAddress(opts *bind.CallOpts, minter common.Address) (IERC721DropAddressMintDetails, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "mintedPerAddress", minter)

	if err != nil {
		return *new(IERC721DropAddressMintDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721DropAddressMintDetails)).(*IERC721DropAddressMintDetails)

	return out0, err

}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsSession) MintedPerAddress(minter common.Address) (IERC721DropAddressMintDetails, error) {
	return _ProtocolRewards.Contract.MintedPerAddress(&_ProtocolRewards.CallOpts, minter)
}

// MintedPerAddress is a free data retrieval call binding the contract method 0xd445b978.
//
// Solidity: function mintedPerAddress(address minter) view returns((uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCallerSession) MintedPerAddress(minter common.Address) (IERC721DropAddressMintDetails, error) {
	return _ProtocolRewards.Contract.MintedPerAddress(&_ProtocolRewards.CallOpts, minter)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProtocolRewards *ProtocolRewardsSession) Name() (string, error) {
	return _ProtocolRewards.Contract.Name(&_ProtocolRewards.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Name() (string, error) {
	return _ProtocolRewards.Contract.Name(&_ProtocolRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) Owner() (common.Address, error) {
	return _ProtocolRewards.Contract.Owner(&_ProtocolRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Owner() (common.Address, error) {
	return _ProtocolRewards.Contract.Owner(&_ProtocolRewards.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ProtocolRewards.Contract.OwnerOf(&_ProtocolRewards.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ProtocolRewards.Contract.OwnerOf(&_ProtocolRewards.CallOpts, tokenId)
}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) PresaleMintsByAddress(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "presaleMintsByAddress", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PresaleMintsByAddress(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.PresaleMintsByAddress(&_ProtocolRewards.CallOpts, arg0)
}

// PresaleMintsByAddress is a free data retrieval call binding the contract method 0x61fea768.
//
// Solidity: function presaleMintsByAddress(address ) view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) PresaleMintsByAddress(arg0 common.Address) (*big.Int, error) {
	return _ProtocolRewards.Contract.PresaleMintsByAddress(&_ProtocolRewards.CallOpts, arg0)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsSession) ProxiableUUID() ([32]byte, error) {
	return _ProtocolRewards.Contract.ProxiableUUID(&_ProtocolRewards.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ProtocolRewards.Contract.ProxiableUUID(&_ProtocolRewards.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_ProtocolRewards *ProtocolRewardsCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "royaltyInfo", arg0, _salePrice)

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
func (_ProtocolRewards *ProtocolRewardsSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ProtocolRewards.Contract.RoyaltyInfo(&_ProtocolRewards.CallOpts, arg0, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 _salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_ProtocolRewards *ProtocolRewardsCallerSession) RoyaltyInfo(arg0 *big.Int, _salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ProtocolRewards.Contract.RoyaltyInfo(&_ProtocolRewards.CallOpts, arg0, _salePrice)
}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsCaller) RoyaltyMintSchedule(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "royaltyMintSchedule")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsSession) RoyaltyMintSchedule() (uint32, error) {
	return _ProtocolRewards.Contract.RoyaltyMintSchedule(&_ProtocolRewards.CallOpts)
}

// RoyaltyMintSchedule is a free data retrieval call binding the contract method 0x6dc45b22.
//
// Solidity: function royaltyMintSchedule() view returns(uint32)
func (_ProtocolRewards *ProtocolRewardsCallerSession) RoyaltyMintSchedule() (uint32, error) {
	return _ProtocolRewards.Contract.RoyaltyMintSchedule(&_ProtocolRewards.CallOpts)
}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCaller) SaleDetails(opts *bind.CallOpts) (IERC721DropSaleDetails, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "saleDetails")

	if err != nil {
		return *new(IERC721DropSaleDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(IERC721DropSaleDetails)).(*IERC721DropSaleDetails)

	return out0, err

}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsSession) SaleDetails() (IERC721DropSaleDetails, error) {
	return _ProtocolRewards.Contract.SaleDetails(&_ProtocolRewards.CallOpts)
}

// SaleDetails is a free data retrieval call binding the contract method 0x3474a4a6.
//
// Solidity: function saleDetails() view returns((bool,bool,uint256,uint64,uint64,uint64,uint64,bytes32,uint256,uint256,uint256))
func (_ProtocolRewards *ProtocolRewardsCallerSession) SaleDetails() (IERC721DropSaleDetails, error) {
	return _ProtocolRewards.Contract.SaleDetails(&_ProtocolRewards.CallOpts)
}

// SalesConfig is a free data retrieval call binding the contract method 0x1d2c0b38.
//
// Solidity: function salesConfig() view returns(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot)
func (_ProtocolRewards *ProtocolRewardsCaller) SalesConfig(opts *bind.CallOpts) (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "salesConfig")

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
func (_ProtocolRewards *ProtocolRewardsSession) SalesConfig() (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	return _ProtocolRewards.Contract.SalesConfig(&_ProtocolRewards.CallOpts)
}

// SalesConfig is a free data retrieval call binding the contract method 0x1d2c0b38.
//
// Solidity: function salesConfig() view returns(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot)
func (_ProtocolRewards *ProtocolRewardsCallerSession) SalesConfig() (struct {
	PublicSalePrice           *big.Int
	MaxSalePurchasePerAddress uint32
	PublicSaleStart           uint64
	PublicSaleEnd             uint64
	PresaleStart              uint64
	PresaleEnd                uint64
	PresaleMerkleRoot         [32]byte
}, error) {
	return _ProtocolRewards.Contract.SalesConfig(&_ProtocolRewards.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProtocolRewards.Contract.SupportsInterface(&_ProtocolRewards.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ProtocolRewards *ProtocolRewardsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ProtocolRewards.Contract.SupportsInterface(&_ProtocolRewards.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProtocolRewards *ProtocolRewardsSession) Symbol() (string, error) {
	return _ProtocolRewards.Contract.Symbol(&_ProtocolRewards.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ProtocolRewards *ProtocolRewardsCallerSession) Symbol() (string, error) {
	return _ProtocolRewards.Contract.Symbol(&_ProtocolRewards.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ProtocolRewards *ProtocolRewardsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ProtocolRewards *ProtocolRewardsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ProtocolRewards.Contract.TokenURI(&_ProtocolRewards.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ProtocolRewards *ProtocolRewardsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ProtocolRewards.Contract.TokenURI(&_ProtocolRewards.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) TotalSupply() (*big.Int, error) {
	return _ProtocolRewards.Contract.TotalSupply(&_ProtocolRewards.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ProtocolRewards *ProtocolRewardsCallerSession) TotalSupply() (*big.Int, error) {
	return _ProtocolRewards.Contract.TotalSupply(&_ProtocolRewards.CallOpts)
}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCaller) ZoraERC721TransferHelper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "zoraERC721TransferHelper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_ProtocolRewards *ProtocolRewardsSession) ZoraERC721TransferHelper() (common.Address, error) {
	return _ProtocolRewards.Contract.ZoraERC721TransferHelper(&_ProtocolRewards.CallOpts)
}

// ZoraERC721TransferHelper is a free data retrieval call binding the contract method 0x00cd4b5e.
//
// Solidity: function zoraERC721TransferHelper() view returns(address)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ZoraERC721TransferHelper() (common.Address, error) {
	return _ProtocolRewards.Contract.ZoraERC721TransferHelper(&_ProtocolRewards.CallOpts)
}

// ZoraFeeForAmount is a free data retrieval call binding the contract method 0xee37be39.
//
// Solidity: function zoraFeeForAmount(uint256 quantity) view returns(address recipient, uint256 fee)
func (_ProtocolRewards *ProtocolRewardsCaller) ZoraFeeForAmount(opts *bind.CallOpts, quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _ProtocolRewards.contract.Call(opts, &out, "zoraFeeForAmount", quantity)

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
func (_ProtocolRewards *ProtocolRewardsSession) ZoraFeeForAmount(quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	return _ProtocolRewards.Contract.ZoraFeeForAmount(&_ProtocolRewards.CallOpts, quantity)
}

// ZoraFeeForAmount is a free data retrieval call binding the contract method 0xee37be39.
//
// Solidity: function zoraFeeForAmount(uint256 quantity) view returns(address recipient, uint256 fee)
func (_ProtocolRewards *ProtocolRewardsCallerSession) ZoraFeeForAmount(quantity *big.Int) (struct {
	Recipient common.Address
	Fee       *big.Int
}, error) {
	return _ProtocolRewards.Contract.ZoraFeeForAmount(&_ProtocolRewards.CallOpts, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) AdminMint(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "adminMint", recipient, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) AdminMint(recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.AdminMint(&_ProtocolRewards.TransactOpts, recipient, quantity)
}

// AdminMint is a paid mutator transaction binding the contract method 0xe58306f9.
//
// Solidity: function adminMint(address recipient, uint256 quantity) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) AdminMint(recipient common.Address, quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.AdminMint(&_ProtocolRewards.TransactOpts, recipient, quantity)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) AdminMintAirdrop(opts *bind.TransactOpts, recipients []common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "adminMintAirdrop", recipients)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) AdminMintAirdrop(recipients []common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.AdminMintAirdrop(&_ProtocolRewards.TransactOpts, recipients)
}

// AdminMintAirdrop is a paid mutator transaction binding the contract method 0xb8ae5a2c.
//
// Solidity: function adminMintAirdrop(address[] recipients) returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) AdminMintAirdrop(recipients []common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.AdminMintAirdrop(&_ProtocolRewards.TransactOpts, recipients)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Approve(&_ProtocolRewards.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Approve(&_ProtocolRewards.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Burn(&_ProtocolRewards.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Burn(&_ProtocolRewards.TransactOpts, tokenId)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsTransactor) CallMetadataRenderer(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "callMetadataRenderer", data)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsSession) CallMetadataRenderer(data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.CallMetadataRenderer(&_ProtocolRewards.TransactOpts, data)
}

// CallMetadataRenderer is a paid mutator transaction binding the contract method 0xb9e7a584.
//
// Solidity: function callMetadataRenderer(bytes data) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) CallMetadataRenderer(data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.CallMetadataRenderer(&_ProtocolRewards.TransactOpts, data)
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) FinalizeOpenEdition(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "finalizeOpenEdition")
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_ProtocolRewards *ProtocolRewardsSession) FinalizeOpenEdition() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.FinalizeOpenEdition(&_ProtocolRewards.TransactOpts)
}

// FinalizeOpenEdition is a paid mutator transaction binding the contract method 0x41e96eb1.
//
// Solidity: function finalizeOpenEdition() returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) FinalizeOpenEdition() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.FinalizeOpenEdition(&_ProtocolRewards.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.GrantRole(&_ProtocolRewards.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.GrantRole(&_ProtocolRewards.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Initialize(opts *bind.TransactOpts, _contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "initialize", _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_ProtocolRewards *ProtocolRewardsSession) Initialize(_contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Initialize(&_ProtocolRewards.TransactOpts, _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// Initialize is a paid mutator transaction binding the contract method 0x5b94e902.
//
// Solidity: function initialize(string _contractName, string _contractSymbol, address _initialOwner, address _fundsRecipient, uint64 _editionSize, uint16 _royaltyBPS, bytes[] _setupCalls, address _metadataRenderer, bytes _metadataRendererInit, address _createReferral) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Initialize(_contractName string, _contractSymbol string, _initialOwner common.Address, _fundsRecipient common.Address, _editionSize uint64, _royaltyBPS uint16, _setupCalls [][]byte, _metadataRenderer common.Address, _metadataRendererInit []byte, _createReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Initialize(&_ProtocolRewards.TransactOpts, _contractName, _contractSymbol, _initialOwner, _fundsRecipient, _editionSize, _royaltyBPS, _setupCalls, _metadataRenderer, _metadataRendererInit, _createReferral)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) ManageMarketFilterDAOSubscription(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "manageMarketFilterDAOSubscription", enable)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_ProtocolRewards *ProtocolRewardsSession) ManageMarketFilterDAOSubscription(enable bool) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ManageMarketFilterDAOSubscription(&_ProtocolRewards.TransactOpts, enable)
}

// ManageMarketFilterDAOSubscription is a paid mutator transaction binding the contract method 0x9bdb89e5.
//
// Solidity: function manageMarketFilterDAOSubscription(bool enable) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) ManageMarketFilterDAOSubscription(enable bool) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.ManageMarketFilterDAOSubscription(&_ProtocolRewards.TransactOpts, enable)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) MintWithRewards(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "mintWithRewards", recipient, quantity, comment, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) MintWithRewards(recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.MintWithRewards(&_ProtocolRewards.TransactOpts, recipient, quantity, comment, mintReferral)
}

// MintWithRewards is a paid mutator transaction binding the contract method 0x45368181.
//
// Solidity: function mintWithRewards(address recipient, uint256 quantity, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) MintWithRewards(recipient common.Address, quantity *big.Int, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.MintWithRewards(&_ProtocolRewards.TransactOpts, recipient, quantity, comment, mintReferral)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ProtocolRewards *ProtocolRewardsTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ProtocolRewards *ProtocolRewardsSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Multicall(&_ProtocolRewards.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Multicall(&_ProtocolRewards.TransactOpts, data)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) Purchase(opts *bind.TransactOpts, quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchase", quantity)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) Purchase(quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Purchase(&_ProtocolRewards.TransactOpts, quantity)
}

// Purchase is a paid mutator transaction binding the contract method 0xefef39a1.
//
// Solidity: function purchase(uint256 quantity) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Purchase(quantity *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Purchase(&_ProtocolRewards.TransactOpts, quantity)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) PurchasePresale(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchasePresale", quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PurchasePresale(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresale(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresale is a paid mutator transaction binding the contract method 0x25024a2b.
//
// Solidity: function purchasePresale(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) PurchasePresale(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresale(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) PurchasePresaleWithComment(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchasePresaleWithComment", quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PurchasePresaleWithComment(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresaleWithComment(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithComment is a paid mutator transaction binding the contract method 0x2e706b5a.
//
// Solidity: function purchasePresaleWithComment(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) PurchasePresaleWithComment(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresaleWithComment(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) PurchasePresaleWithRewards(opts *bind.TransactOpts, quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchasePresaleWithRewards", quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PurchasePresaleWithRewards(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresaleWithRewards(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchasePresaleWithRewards is a paid mutator transaction binding the contract method 0xae6e7875.
//
// Solidity: function purchasePresaleWithRewards(uint256 quantity, uint256 maxQuantity, uint256 pricePerToken, bytes32[] merkleProof, string comment, address mintReferral) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) PurchasePresaleWithRewards(quantity *big.Int, maxQuantity *big.Int, pricePerToken *big.Int, merkleProof [][32]byte, comment string, mintReferral common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchasePresaleWithRewards(&_ProtocolRewards.TransactOpts, quantity, maxQuantity, pricePerToken, merkleProof, comment, mintReferral)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) PurchaseWithComment(opts *bind.TransactOpts, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchaseWithComment", quantity, comment)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PurchaseWithComment(quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchaseWithComment(&_ProtocolRewards.TransactOpts, quantity, comment)
}

// PurchaseWithComment is a paid mutator transaction binding the contract method 0x03ee2733.
//
// Solidity: function purchaseWithComment(uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) PurchaseWithComment(quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchaseWithComment(&_ProtocolRewards.TransactOpts, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactor) PurchaseWithRecipient(opts *bind.TransactOpts, recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "purchaseWithRecipient", recipient, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsSession) PurchaseWithRecipient(recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchaseWithRecipient(&_ProtocolRewards.TransactOpts, recipient, quantity, comment)
}

// PurchaseWithRecipient is a paid mutator transaction binding the contract method 0xd234255c.
//
// Solidity: function purchaseWithRecipient(address recipient, uint256 quantity, string comment) payable returns(uint256)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) PurchaseWithRecipient(recipient common.Address, quantity *big.Int, comment string) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.PurchaseWithRecipient(&_ProtocolRewards.TransactOpts, recipient, quantity, comment)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.RenounceRole(&_ProtocolRewards.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.RenounceRole(&_ProtocolRewards.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.RevokeRole(&_ProtocolRewards.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.RevokeRole(&_ProtocolRewards.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SafeTransferFrom(&_ProtocolRewards.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SafeTransferFrom(&_ProtocolRewards.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SafeTransferFrom0(&_ProtocolRewards.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SafeTransferFrom0(&_ProtocolRewards.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetApprovalForAll(&_ProtocolRewards.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetApprovalForAll(&_ProtocolRewards.TransactOpts, operator, approved)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SetFundsRecipient(opts *bind.TransactOpts, newRecipientAddress common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "setFundsRecipient", newRecipientAddress)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SetFundsRecipient(newRecipientAddress common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetFundsRecipient(&_ProtocolRewards.TransactOpts, newRecipientAddress)
}

// SetFundsRecipient is a paid mutator transaction binding the contract method 0x10a7eb5d.
//
// Solidity: function setFundsRecipient(address newRecipientAddress) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SetFundsRecipient(newRecipientAddress common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetFundsRecipient(&_ProtocolRewards.TransactOpts, newRecipientAddress)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SetMetadataRenderer(opts *bind.TransactOpts, newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "setMetadataRenderer", newRenderer, setupRenderer)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SetMetadataRenderer(newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetMetadataRenderer(&_ProtocolRewards.TransactOpts, newRenderer, setupRenderer)
}

// SetMetadataRenderer is a paid mutator transaction binding the contract method 0x3bcdcc87.
//
// Solidity: function setMetadataRenderer(address newRenderer, bytes setupRenderer) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SetMetadataRenderer(newRenderer common.Address, setupRenderer []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetMetadataRenderer(&_ProtocolRewards.TransactOpts, newRenderer, setupRenderer)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetOwner(&_ProtocolRewards.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetOwner(&_ProtocolRewards.TransactOpts, newOwner)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) SetSaleConfiguration(opts *bind.TransactOpts, publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "setSaleConfiguration", publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_ProtocolRewards *ProtocolRewardsSession) SetSaleConfiguration(publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetSaleConfiguration(&_ProtocolRewards.TransactOpts, publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// SetSaleConfiguration is a paid mutator transaction binding the contract method 0xffdb7163.
//
// Solidity: function setSaleConfiguration(uint104 publicSalePrice, uint32 maxSalePurchasePerAddress, uint64 publicSaleStart, uint64 publicSaleEnd, uint64 presaleStart, uint64 presaleEnd, bytes32 presaleMerkleRoot) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) SetSaleConfiguration(publicSalePrice *big.Int, maxSalePurchasePerAddress uint32, publicSaleStart uint64, publicSaleEnd uint64, presaleStart uint64, presaleEnd uint64, presaleMerkleRoot [32]byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.SetSaleConfiguration(&_ProtocolRewards.TransactOpts, publicSalePrice, maxSalePurchasePerAddress, publicSaleStart, publicSaleEnd, presaleStart, presaleEnd, presaleMerkleRoot)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.TransferFrom(&_ProtocolRewards.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.TransferFrom(&_ProtocolRewards.TransactOpts, from, to, tokenId)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) UpdateCreateReferral(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "updateCreateReferral", recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_ProtocolRewards *ProtocolRewardsSession) UpdateCreateReferral(recipient common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateCreateReferral(&_ProtocolRewards.TransactOpts, recipient)
}

// UpdateCreateReferral is a paid mutator transaction binding the contract method 0x8b338c7c.
//
// Solidity: function updateCreateReferral(address recipient) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) UpdateCreateReferral(recipient common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateCreateReferral(&_ProtocolRewards.TransactOpts, recipient)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsTransactor) UpdateMarketFilterSettings(opts *bind.TransactOpts, args []byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "updateMarketFilterSettings", args)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsSession) UpdateMarketFilterSettings(args []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateMarketFilterSettings(&_ProtocolRewards.TransactOpts, args)
}

// UpdateMarketFilterSettings is a paid mutator transaction binding the contract method 0xbb20d857.
//
// Solidity: function updateMarketFilterSettings(bytes args) returns(bytes)
func (_ProtocolRewards *ProtocolRewardsTransactorSession) UpdateMarketFilterSettings(args []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateMarketFilterSettings(&_ProtocolRewards.TransactOpts, args)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) UpdateRoyaltyMintSchedule(opts *bind.TransactOpts, newSchedule uint32) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "updateRoyaltyMintSchedule", newSchedule)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_ProtocolRewards *ProtocolRewardsSession) UpdateRoyaltyMintSchedule(newSchedule uint32) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateRoyaltyMintSchedule(&_ProtocolRewards.TransactOpts, newSchedule)
}

// UpdateRoyaltyMintSchedule is a paid mutator transaction binding the contract method 0xd0bd3c6b.
//
// Solidity: function updateRoyaltyMintSchedule(uint32 newSchedule) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) UpdateRoyaltyMintSchedule(newSchedule uint32) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpdateRoyaltyMintSchedule(&_ProtocolRewards.TransactOpts, newSchedule)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolRewards *ProtocolRewardsSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpgradeTo(&_ProtocolRewards.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpgradeTo(&_ProtocolRewards.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolRewards *ProtocolRewardsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpgradeToAndCall(&_ProtocolRewards.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.UpgradeToAndCall(&_ProtocolRewards.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolRewards *ProtocolRewardsSession) Withdraw() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Withdraw(&_ProtocolRewards.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Withdraw(&_ProtocolRewards.TransactOpts)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) WithdrawRewards(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.contract.Transact(opts, "withdrawRewards", to, amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsSession) WithdrawRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawRewards(&_ProtocolRewards.TransactOpts, to, amount)
}

// WithdrawRewards is a paid mutator transaction binding the contract method 0xd6ef7af0.
//
// Solidity: function withdrawRewards(address to, uint256 amount) returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) WithdrawRewards(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ProtocolRewards.Contract.WithdrawRewards(&_ProtocolRewards.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolRewards.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolRewards *ProtocolRewardsSession) Receive() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Receive(&_ProtocolRewards.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolRewards *ProtocolRewardsTransactorSession) Receive() (*types.Transaction, error) {
	return _ProtocolRewards.Contract.Receive(&_ProtocolRewards.TransactOpts)
}

// ProtocolRewardsAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ProtocolRewards contract.
type ProtocolRewardsAdminChangedIterator struct {
	Event *ProtocolRewardsAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsAdminChanged)
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
		it.Event = new(ProtocolRewardsAdminChanged)
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
func (it *ProtocolRewardsAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsAdminChanged represents a AdminChanged event raised by the ProtocolRewards contract.
type ProtocolRewardsAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ProtocolRewardsAdminChangedIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsAdminChangedIterator{contract: _ProtocolRewards.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsAdminChanged)
				if err := _ProtocolRewards.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseAdminChanged(log types.Log) (*ProtocolRewardsAdminChanged, error) {
	event := new(ProtocolRewardsAdminChanged)
	if err := _ProtocolRewards.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ProtocolRewards contract.
type ProtocolRewardsApprovalIterator struct {
	Event *ProtocolRewardsApproval // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsApproval)
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
		it.Event = new(ProtocolRewardsApproval)
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
func (it *ProtocolRewardsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsApproval represents a Approval event raised by the ProtocolRewards contract.
type ProtocolRewardsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ProtocolRewardsApprovalIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsApprovalIterator{contract: _ProtocolRewards.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsApproval)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseApproval(log types.Log) (*ProtocolRewardsApproval, error) {
	event := new(ProtocolRewardsApproval)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ProtocolRewards contract.
type ProtocolRewardsApprovalForAllIterator struct {
	Event *ProtocolRewardsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsApprovalForAll)
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
		it.Event = new(ProtocolRewardsApprovalForAll)
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
func (it *ProtocolRewardsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsApprovalForAll represents a ApprovalForAll event raised by the ProtocolRewards contract.
type ProtocolRewardsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ProtocolRewardsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsApprovalForAllIterator{contract: _ProtocolRewards.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsApprovalForAll)
				if err := _ProtocolRewards.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseApprovalForAll(log types.Log) (*ProtocolRewardsApprovalForAll, error) {
	event := new(ProtocolRewardsApprovalForAll)
	if err := _ProtocolRewards.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the ProtocolRewards contract.
type ProtocolRewardsBatchMetadataUpdateIterator struct {
	Event *ProtocolRewardsBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsBatchMetadataUpdate)
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
		it.Event = new(ProtocolRewardsBatchMetadataUpdate)
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
func (it *ProtocolRewardsBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the ProtocolRewards contract.
type ProtocolRewardsBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ProtocolRewardsBatchMetadataUpdateIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsBatchMetadataUpdateIterator{contract: _ProtocolRewards.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsBatchMetadataUpdate)
				if err := _ProtocolRewards.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseBatchMetadataUpdate(log types.Log) (*ProtocolRewardsBatchMetadataUpdate, error) {
	event := new(ProtocolRewardsBatchMetadataUpdate)
	if err := _ProtocolRewards.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ProtocolRewards contract.
type ProtocolRewardsBeaconUpgradedIterator struct {
	Event *ProtocolRewardsBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsBeaconUpgraded)
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
		it.Event = new(ProtocolRewardsBeaconUpgraded)
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
func (it *ProtocolRewardsBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsBeaconUpgraded represents a BeaconUpgraded event raised by the ProtocolRewards contract.
type ProtocolRewardsBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ProtocolRewardsBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsBeaconUpgradedIterator{contract: _ProtocolRewards.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsBeaconUpgraded)
				if err := _ProtocolRewards.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseBeaconUpgraded(log types.Log) (*ProtocolRewardsBeaconUpgraded, error) {
	event := new(ProtocolRewardsBeaconUpgraded)
	if err := _ProtocolRewards.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsFundsReceivedIterator is returned from FilterFundsReceived and is used to iterate over the raw logs and unpacked data for FundsReceived events raised by the ProtocolRewards contract.
type ProtocolRewardsFundsReceivedIterator struct {
	Event *ProtocolRewardsFundsReceived // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsFundsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsFundsReceived)
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
		it.Event = new(ProtocolRewardsFundsReceived)
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
func (it *ProtocolRewardsFundsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsFundsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsFundsReceived represents a FundsReceived event raised by the ProtocolRewards contract.
type ProtocolRewardsFundsReceived struct {
	Source common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsReceived is a free log retrieval operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed source, uint256 amount)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterFundsReceived(opts *bind.FilterOpts, source []common.Address) (*ProtocolRewardsFundsReceivedIterator, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "FundsReceived", sourceRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsFundsReceivedIterator{contract: _ProtocolRewards.contract, event: "FundsReceived", logs: logs, sub: sub}, nil
}

// WatchFundsReceived is a free log subscription operation binding the contract event 0x8e47b87b0ef542cdfa1659c551d88bad38aa7f452d2bbb349ab7530dfec8be8f.
//
// Solidity: event FundsReceived(address indexed source, uint256 amount)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchFundsReceived(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsFundsReceived, source []common.Address) (event.Subscription, error) {

	var sourceRule []interface{}
	for _, sourceItem := range source {
		sourceRule = append(sourceRule, sourceItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "FundsReceived", sourceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsFundsReceived)
				if err := _ProtocolRewards.contract.UnpackLog(event, "FundsReceived", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseFundsReceived(log types.Log) (*ProtocolRewardsFundsReceived, error) {
	event := new(ProtocolRewardsFundsReceived)
	if err := _ProtocolRewards.contract.UnpackLog(event, "FundsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsFundsRecipientChangedIterator is returned from FilterFundsRecipientChanged and is used to iterate over the raw logs and unpacked data for FundsRecipientChanged events raised by the ProtocolRewards contract.
type ProtocolRewardsFundsRecipientChangedIterator struct {
	Event *ProtocolRewardsFundsRecipientChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsFundsRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsFundsRecipientChanged)
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
		it.Event = new(ProtocolRewardsFundsRecipientChanged)
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
func (it *ProtocolRewardsFundsRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsFundsRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsFundsRecipientChanged represents a FundsRecipientChanged event raised by the ProtocolRewards contract.
type ProtocolRewardsFundsRecipientChanged struct {
	NewAddress common.Address
	ChangedBy  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFundsRecipientChanged is a free log retrieval operation binding the contract event 0x70a7ea5c664ab9c21baf3da59bb2f1e1ca33557b08a0031fab4f170767449951.
//
// Solidity: event FundsRecipientChanged(address indexed newAddress, address indexed changedBy)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterFundsRecipientChanged(opts *bind.FilterOpts, newAddress []common.Address, changedBy []common.Address) (*ProtocolRewardsFundsRecipientChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "FundsRecipientChanged", newAddressRule, changedByRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsFundsRecipientChangedIterator{contract: _ProtocolRewards.contract, event: "FundsRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchFundsRecipientChanged is a free log subscription operation binding the contract event 0x70a7ea5c664ab9c21baf3da59bb2f1e1ca33557b08a0031fab4f170767449951.
//
// Solidity: event FundsRecipientChanged(address indexed newAddress, address indexed changedBy)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchFundsRecipientChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsFundsRecipientChanged, newAddress []common.Address, changedBy []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "FundsRecipientChanged", newAddressRule, changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsFundsRecipientChanged)
				if err := _ProtocolRewards.contract.UnpackLog(event, "FundsRecipientChanged", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseFundsRecipientChanged(log types.Log) (*ProtocolRewardsFundsRecipientChanged, error) {
	event := new(ProtocolRewardsFundsRecipientChanged)
	if err := _ProtocolRewards.contract.UnpackLog(event, "FundsRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the ProtocolRewards contract.
type ProtocolRewardsFundsWithdrawnIterator struct {
	Event *ProtocolRewardsFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsFundsWithdrawn)
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
		it.Event = new(ProtocolRewardsFundsWithdrawn)
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
func (it *ProtocolRewardsFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsFundsWithdrawn represents a FundsWithdrawn event raised by the ProtocolRewards contract.
type ProtocolRewardsFundsWithdrawn struct {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, withdrawnBy []common.Address, withdrawnTo []common.Address) (*ProtocolRewardsFundsWithdrawnIterator, error) {

	var withdrawnByRule []interface{}
	for _, withdrawnByItem := range withdrawnBy {
		withdrawnByRule = append(withdrawnByRule, withdrawnByItem)
	}
	var withdrawnToRule []interface{}
	for _, withdrawnToItem := range withdrawnTo {
		withdrawnToRule = append(withdrawnToRule, withdrawnToItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "FundsWithdrawn", withdrawnByRule, withdrawnToRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsFundsWithdrawnIterator{contract: _ProtocolRewards.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0x8a95554e4c9dcaaf33f247387f2ee77390780487d3365e3a804788791a1df500.
//
// Solidity: event FundsWithdrawn(address indexed withdrawnBy, address indexed withdrawnTo, uint256 amount, address feeRecipient, uint256 feeAmount)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsFundsWithdrawn, withdrawnBy []common.Address, withdrawnTo []common.Address) (event.Subscription, error) {

	var withdrawnByRule []interface{}
	for _, withdrawnByItem := range withdrawnBy {
		withdrawnByRule = append(withdrawnByRule, withdrawnByItem)
	}
	var withdrawnToRule []interface{}
	for _, withdrawnToItem := range withdrawnTo {
		withdrawnToRule = append(withdrawnToRule, withdrawnToItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "FundsWithdrawn", withdrawnByRule, withdrawnToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsFundsWithdrawn)
				if err := _ProtocolRewards.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseFundsWithdrawn(log types.Log) (*ProtocolRewardsFundsWithdrawn, error) {
	event := new(ProtocolRewardsFundsWithdrawn)
	if err := _ProtocolRewards.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the ProtocolRewards contract.
type ProtocolRewardsMetadataUpdateIterator struct {
	Event *ProtocolRewardsMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsMetadataUpdate)
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
		it.Event = new(ProtocolRewardsMetadataUpdate)
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
func (it *ProtocolRewardsMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsMetadataUpdate represents a MetadataUpdate event raised by the ProtocolRewards contract.
type ProtocolRewardsMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ProtocolRewardsMetadataUpdateIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsMetadataUpdateIterator{contract: _ProtocolRewards.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsMetadataUpdate)
				if err := _ProtocolRewards.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseMetadataUpdate(log types.Log) (*ProtocolRewardsMetadataUpdate, error) {
	event := new(ProtocolRewardsMetadataUpdate)
	if err := _ProtocolRewards.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsMintCommentIterator is returned from FilterMintComment and is used to iterate over the raw logs and unpacked data for MintComment events raised by the ProtocolRewards contract.
type ProtocolRewardsMintCommentIterator struct {
	Event *ProtocolRewardsMintComment // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsMintCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsMintComment)
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
		it.Event = new(ProtocolRewardsMintComment)
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
func (it *ProtocolRewardsMintCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsMintCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsMintComment represents a MintComment event raised by the ProtocolRewards contract.
type ProtocolRewardsMintComment struct {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterMintComment(opts *bind.FilterOpts, sender []common.Address, tokenContract []common.Address, tokenId []*big.Int) (*ProtocolRewardsMintCommentIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "MintComment", senderRule, tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsMintCommentIterator{contract: _ProtocolRewards.contract, event: "MintComment", logs: logs, sub: sub}, nil
}

// WatchMintComment is a free log subscription operation binding the contract event 0xb9490aee663998179ad13f9e1c1eb6189c71ad1a9ec87f33ad2766f98d9a268a.
//
// Solidity: event MintComment(address indexed sender, address indexed tokenContract, uint256 indexed tokenId, uint256 quantity, string comment)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchMintComment(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsMintComment, sender []common.Address, tokenContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "MintComment", senderRule, tokenContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsMintComment)
				if err := _ProtocolRewards.contract.UnpackLog(event, "MintComment", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseMintComment(log types.Log) (*ProtocolRewardsMintComment, error) {
	event := new(ProtocolRewardsMintComment)
	if err := _ProtocolRewards.contract.UnpackLog(event, "MintComment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsMintFeePayoutIterator is returned from FilterMintFeePayout and is used to iterate over the raw logs and unpacked data for MintFeePayout events raised by the ProtocolRewards contract.
type ProtocolRewardsMintFeePayoutIterator struct {
	Event *ProtocolRewardsMintFeePayout // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsMintFeePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsMintFeePayout)
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
		it.Event = new(ProtocolRewardsMintFeePayout)
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
func (it *ProtocolRewardsMintFeePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsMintFeePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsMintFeePayout represents a MintFeePayout event raised by the ProtocolRewards contract.
type ProtocolRewardsMintFeePayout struct {
	MintFeeAmount    *big.Int
	MintFeeRecipient common.Address
	Success          bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMintFeePayout is a free log retrieval operation binding the contract event 0x6f8da53cfedb8cc4f7935c3629624e50b63053c93bb2cad246aa4d3a2ba7d4ce.
//
// Solidity: event MintFeePayout(uint256 mintFeeAmount, address mintFeeRecipient, bool success)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterMintFeePayout(opts *bind.FilterOpts) (*ProtocolRewardsMintFeePayoutIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "MintFeePayout")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsMintFeePayoutIterator{contract: _ProtocolRewards.contract, event: "MintFeePayout", logs: logs, sub: sub}, nil
}

// WatchMintFeePayout is a free log subscription operation binding the contract event 0x6f8da53cfedb8cc4f7935c3629624e50b63053c93bb2cad246aa4d3a2ba7d4ce.
//
// Solidity: event MintFeePayout(uint256 mintFeeAmount, address mintFeeRecipient, bool success)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchMintFeePayout(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsMintFeePayout) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "MintFeePayout")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsMintFeePayout)
				if err := _ProtocolRewards.contract.UnpackLog(event, "MintFeePayout", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseMintFeePayout(log types.Log) (*ProtocolRewardsMintFeePayout, error) {
	event := new(ProtocolRewardsMintFeePayout)
	if err := _ProtocolRewards.contract.UnpackLog(event, "MintFeePayout", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsOpenMintFinalizedIterator is returned from FilterOpenMintFinalized and is used to iterate over the raw logs and unpacked data for OpenMintFinalized events raised by the ProtocolRewards contract.
type ProtocolRewardsOpenMintFinalizedIterator struct {
	Event *ProtocolRewardsOpenMintFinalized // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsOpenMintFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsOpenMintFinalized)
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
		it.Event = new(ProtocolRewardsOpenMintFinalized)
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
func (it *ProtocolRewardsOpenMintFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsOpenMintFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsOpenMintFinalized represents a OpenMintFinalized event raised by the ProtocolRewards contract.
type ProtocolRewardsOpenMintFinalized struct {
	Sender        common.Address
	NumberOfMints *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOpenMintFinalized is a free log retrieval operation binding the contract event 0xb6cc1e876b8d7479c8afe025a58658b0f3c3ec5bd0f28cb4261326b162069bf8.
//
// Solidity: event OpenMintFinalized(address indexed sender, uint256 numberOfMints)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterOpenMintFinalized(opts *bind.FilterOpts, sender []common.Address) (*ProtocolRewardsOpenMintFinalizedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "OpenMintFinalized", senderRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsOpenMintFinalizedIterator{contract: _ProtocolRewards.contract, event: "OpenMintFinalized", logs: logs, sub: sub}, nil
}

// WatchOpenMintFinalized is a free log subscription operation binding the contract event 0xb6cc1e876b8d7479c8afe025a58658b0f3c3ec5bd0f28cb4261326b162069bf8.
//
// Solidity: event OpenMintFinalized(address indexed sender, uint256 numberOfMints)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchOpenMintFinalized(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsOpenMintFinalized, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "OpenMintFinalized", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsOpenMintFinalized)
				if err := _ProtocolRewards.contract.UnpackLog(event, "OpenMintFinalized", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseOpenMintFinalized(log types.Log) (*ProtocolRewardsOpenMintFinalized, error) {
	event := new(ProtocolRewardsOpenMintFinalized)
	if err := _ProtocolRewards.contract.UnpackLog(event, "OpenMintFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsOwnerCanceledIterator is returned from FilterOwnerCanceled and is used to iterate over the raw logs and unpacked data for OwnerCanceled events raised by the ProtocolRewards contract.
type ProtocolRewardsOwnerCanceledIterator struct {
	Event *ProtocolRewardsOwnerCanceled // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsOwnerCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsOwnerCanceled)
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
		it.Event = new(ProtocolRewardsOwnerCanceled)
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
func (it *ProtocolRewardsOwnerCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsOwnerCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsOwnerCanceled represents a OwnerCanceled event raised by the ProtocolRewards contract.
type ProtocolRewardsOwnerCanceled struct {
	PreviousOwner     common.Address
	PotentialNewOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOwnerCanceled is a free log retrieval operation binding the contract event 0x682679deecef4dcd49674845cc1e3a075fea9073680aa445a8207d5a4bdea3da.
//
// Solidity: event OwnerCanceled(address indexed previousOwner, address indexed potentialNewOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterOwnerCanceled(opts *bind.FilterOpts, previousOwner []common.Address, potentialNewOwner []common.Address) (*ProtocolRewardsOwnerCanceledIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "OwnerCanceled", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsOwnerCanceledIterator{contract: _ProtocolRewards.contract, event: "OwnerCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnerCanceled is a free log subscription operation binding the contract event 0x682679deecef4dcd49674845cc1e3a075fea9073680aa445a8207d5a4bdea3da.
//
// Solidity: event OwnerCanceled(address indexed previousOwner, address indexed potentialNewOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchOwnerCanceled(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsOwnerCanceled, previousOwner []common.Address, potentialNewOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "OwnerCanceled", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsOwnerCanceled)
				if err := _ProtocolRewards.contract.UnpackLog(event, "OwnerCanceled", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseOwnerCanceled(log types.Log) (*ProtocolRewardsOwnerCanceled, error) {
	event := new(ProtocolRewardsOwnerCanceled)
	if err := _ProtocolRewards.contract.UnpackLog(event, "OwnerCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsOwnerPendingIterator is returned from FilterOwnerPending and is used to iterate over the raw logs and unpacked data for OwnerPending events raised by the ProtocolRewards contract.
type ProtocolRewardsOwnerPendingIterator struct {
	Event *ProtocolRewardsOwnerPending // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsOwnerPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsOwnerPending)
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
		it.Event = new(ProtocolRewardsOwnerPending)
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
func (it *ProtocolRewardsOwnerPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsOwnerPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsOwnerPending represents a OwnerPending event raised by the ProtocolRewards contract.
type ProtocolRewardsOwnerPending struct {
	PreviousOwner     common.Address
	PotentialNewOwner common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOwnerPending is a free log retrieval operation binding the contract event 0x4f2638f5949b9614ef8d5e268cb51348ad7f434a34812bf64b6e95014fbd357e.
//
// Solidity: event OwnerPending(address indexed previousOwner, address indexed potentialNewOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterOwnerPending(opts *bind.FilterOpts, previousOwner []common.Address, potentialNewOwner []common.Address) (*ProtocolRewardsOwnerPendingIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "OwnerPending", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsOwnerPendingIterator{contract: _ProtocolRewards.contract, event: "OwnerPending", logs: logs, sub: sub}, nil
}

// WatchOwnerPending is a free log subscription operation binding the contract event 0x4f2638f5949b9614ef8d5e268cb51348ad7f434a34812bf64b6e95014fbd357e.
//
// Solidity: event OwnerPending(address indexed previousOwner, address indexed potentialNewOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchOwnerPending(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsOwnerPending, previousOwner []common.Address, potentialNewOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var potentialNewOwnerRule []interface{}
	for _, potentialNewOwnerItem := range potentialNewOwner {
		potentialNewOwnerRule = append(potentialNewOwnerRule, potentialNewOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "OwnerPending", previousOwnerRule, potentialNewOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsOwnerPending)
				if err := _ProtocolRewards.contract.UnpackLog(event, "OwnerPending", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseOwnerPending(log types.Log) (*ProtocolRewardsOwnerPending, error) {
	event := new(ProtocolRewardsOwnerPending)
	if err := _ProtocolRewards.contract.UnpackLog(event, "OwnerPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProtocolRewards contract.
type ProtocolRewardsOwnershipTransferredIterator struct {
	Event *ProtocolRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsOwnershipTransferred)
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
		it.Event = new(ProtocolRewardsOwnershipTransferred)
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
func (it *ProtocolRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the ProtocolRewards contract.
type ProtocolRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProtocolRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsOwnershipTransferredIterator{contract: _ProtocolRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsOwnershipTransferred)
				if err := _ProtocolRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*ProtocolRewardsOwnershipTransferred, error) {
	event := new(ProtocolRewardsOwnershipTransferred)
	if err := _ProtocolRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ProtocolRewards contract.
type ProtocolRewardsRoleAdminChangedIterator struct {
	Event *ProtocolRewardsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsRoleAdminChanged)
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
		it.Event = new(ProtocolRewardsRoleAdminChanged)
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
func (it *ProtocolRewardsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsRoleAdminChanged represents a RoleAdminChanged event raised by the ProtocolRewards contract.
type ProtocolRewardsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ProtocolRewardsRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsRoleAdminChangedIterator{contract: _ProtocolRewards.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsRoleAdminChanged)
				if err := _ProtocolRewards.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseRoleAdminChanged(log types.Log) (*ProtocolRewardsRoleAdminChanged, error) {
	event := new(ProtocolRewardsRoleAdminChanged)
	if err := _ProtocolRewards.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ProtocolRewards contract.
type ProtocolRewardsRoleGrantedIterator struct {
	Event *ProtocolRewardsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsRoleGranted)
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
		it.Event = new(ProtocolRewardsRoleGranted)
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
func (it *ProtocolRewardsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsRoleGranted represents a RoleGranted event raised by the ProtocolRewards contract.
type ProtocolRewardsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProtocolRewardsRoleGrantedIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsRoleGrantedIterator{contract: _ProtocolRewards.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsRoleGranted)
				if err := _ProtocolRewards.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseRoleGranted(log types.Log) (*ProtocolRewardsRoleGranted, error) {
	event := new(ProtocolRewardsRoleGranted)
	if err := _ProtocolRewards.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ProtocolRewards contract.
type ProtocolRewardsRoleRevokedIterator struct {
	Event *ProtocolRewardsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsRoleRevoked)
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
		it.Event = new(ProtocolRewardsRoleRevoked)
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
func (it *ProtocolRewardsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsRoleRevoked represents a RoleRevoked event raised by the ProtocolRewards contract.
type ProtocolRewardsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ProtocolRewardsRoleRevokedIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsRoleRevokedIterator{contract: _ProtocolRewards.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsRoleRevoked)
				if err := _ProtocolRewards.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseRoleRevoked(log types.Log) (*ProtocolRewardsRoleRevoked, error) {
	event := new(ProtocolRewardsRoleRevoked)
	if err := _ProtocolRewards.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsSaleIterator is returned from FilterSale and is used to iterate over the raw logs and unpacked data for Sale events raised by the ProtocolRewards contract.
type ProtocolRewardsSaleIterator struct {
	Event *ProtocolRewardsSale // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsSale)
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
		it.Event = new(ProtocolRewardsSale)
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
func (it *ProtocolRewardsSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsSale represents a Sale event raised by the ProtocolRewards contract.
type ProtocolRewardsSale struct {
	To                    common.Address
	Quantity              *big.Int
	PricePerToken         *big.Int
	FirstPurchasedTokenId *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSale is a free log retrieval operation binding the contract event 0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90.
//
// Solidity: event Sale(address indexed to, uint256 indexed quantity, uint256 indexed pricePerToken, uint256 firstPurchasedTokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterSale(opts *bind.FilterOpts, to []common.Address, quantity []*big.Int, pricePerToken []*big.Int) (*ProtocolRewardsSaleIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Sale", toRule, quantityRule, pricePerTokenRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsSaleIterator{contract: _ProtocolRewards.contract, event: "Sale", logs: logs, sub: sub}, nil
}

// WatchSale is a free log subscription operation binding the contract event 0x4e26b0356a15833a75d497ecc40ebbb716b99466ed0dba9454f1fff451e25a90.
//
// Solidity: event Sale(address indexed to, uint256 indexed quantity, uint256 indexed pricePerToken, uint256 firstPurchasedTokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchSale(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsSale, to []common.Address, quantity []*big.Int, pricePerToken []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Sale", toRule, quantityRule, pricePerTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsSale)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Sale", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseSale(log types.Log) (*ProtocolRewardsSale, error) {
	event := new(ProtocolRewardsSale)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Sale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsSalesConfigChangedIterator is returned from FilterSalesConfigChanged and is used to iterate over the raw logs and unpacked data for SalesConfigChanged events raised by the ProtocolRewards contract.
type ProtocolRewardsSalesConfigChangedIterator struct {
	Event *ProtocolRewardsSalesConfigChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsSalesConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsSalesConfigChanged)
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
		it.Event = new(ProtocolRewardsSalesConfigChanged)
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
func (it *ProtocolRewardsSalesConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsSalesConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsSalesConfigChanged represents a SalesConfigChanged event raised by the ProtocolRewards contract.
type ProtocolRewardsSalesConfigChanged struct {
	ChangedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSalesConfigChanged is a free log retrieval operation binding the contract event 0xc1ff5e4744ac8dd2b8027a10e3723b165975297501c71c4e7dcb8796d96375db.
//
// Solidity: event SalesConfigChanged(address indexed changedBy)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterSalesConfigChanged(opts *bind.FilterOpts, changedBy []common.Address) (*ProtocolRewardsSalesConfigChangedIterator, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "SalesConfigChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsSalesConfigChangedIterator{contract: _ProtocolRewards.contract, event: "SalesConfigChanged", logs: logs, sub: sub}, nil
}

// WatchSalesConfigChanged is a free log subscription operation binding the contract event 0xc1ff5e4744ac8dd2b8027a10e3723b165975297501c71c4e7dcb8796d96375db.
//
// Solidity: event SalesConfigChanged(address indexed changedBy)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchSalesConfigChanged(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsSalesConfigChanged, changedBy []common.Address) (event.Subscription, error) {

	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "SalesConfigChanged", changedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsSalesConfigChanged)
				if err := _ProtocolRewards.contract.UnpackLog(event, "SalesConfigChanged", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseSalesConfigChanged(log types.Log) (*ProtocolRewardsSalesConfigChanged, error) {
	event := new(ProtocolRewardsSalesConfigChanged)
	if err := _ProtocolRewards.contract.UnpackLog(event, "SalesConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ProtocolRewards contract.
type ProtocolRewardsTransferIterator struct {
	Event *ProtocolRewardsTransfer // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsTransfer)
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
		it.Event = new(ProtocolRewardsTransfer)
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
func (it *ProtocolRewardsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsTransfer represents a Transfer event raised by the ProtocolRewards contract.
type ProtocolRewardsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ProtocolRewardsTransferIterator, error) {

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

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsTransferIterator{contract: _ProtocolRewards.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsTransfer)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseTransfer(log types.Log) (*ProtocolRewardsTransfer, error) {
	event := new(ProtocolRewardsTransfer)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsUpdatedMetadataRendererIterator is returned from FilterUpdatedMetadataRenderer and is used to iterate over the raw logs and unpacked data for UpdatedMetadataRenderer events raised by the ProtocolRewards contract.
type ProtocolRewardsUpdatedMetadataRendererIterator struct {
	Event *ProtocolRewardsUpdatedMetadataRenderer // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsUpdatedMetadataRendererIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsUpdatedMetadataRenderer)
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
		it.Event = new(ProtocolRewardsUpdatedMetadataRenderer)
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
func (it *ProtocolRewardsUpdatedMetadataRendererIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsUpdatedMetadataRendererIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsUpdatedMetadataRenderer represents a UpdatedMetadataRenderer event raised by the ProtocolRewards contract.
type ProtocolRewardsUpdatedMetadataRenderer struct {
	Sender   common.Address
	Renderer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpdatedMetadataRenderer is a free log retrieval operation binding the contract event 0x046c5d913c35948c3e0e44c3599eb14bf33b73f141fa8bb282b300414998b868.
//
// Solidity: event UpdatedMetadataRenderer(address sender, address renderer)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterUpdatedMetadataRenderer(opts *bind.FilterOpts) (*ProtocolRewardsUpdatedMetadataRendererIterator, error) {

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "UpdatedMetadataRenderer")
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsUpdatedMetadataRendererIterator{contract: _ProtocolRewards.contract, event: "UpdatedMetadataRenderer", logs: logs, sub: sub}, nil
}

// WatchUpdatedMetadataRenderer is a free log subscription operation binding the contract event 0x046c5d913c35948c3e0e44c3599eb14bf33b73f141fa8bb282b300414998b868.
//
// Solidity: event UpdatedMetadataRenderer(address sender, address renderer)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchUpdatedMetadataRenderer(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsUpdatedMetadataRenderer) (event.Subscription, error) {

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "UpdatedMetadataRenderer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsUpdatedMetadataRenderer)
				if err := _ProtocolRewards.contract.UnpackLog(event, "UpdatedMetadataRenderer", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseUpdatedMetadataRenderer(log types.Log) (*ProtocolRewardsUpdatedMetadataRenderer, error) {
	event := new(ProtocolRewardsUpdatedMetadataRenderer)
	if err := _ProtocolRewards.contract.UnpackLog(event, "UpdatedMetadataRenderer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolRewardsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ProtocolRewards contract.
type ProtocolRewardsUpgradedIterator struct {
	Event *ProtocolRewardsUpgraded // Event containing the contract specifics and raw log

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
func (it *ProtocolRewardsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolRewardsUpgraded)
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
		it.Event = new(ProtocolRewardsUpgraded)
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
func (it *ProtocolRewardsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolRewardsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolRewardsUpgraded represents a Upgraded event raised by the ProtocolRewards contract.
type ProtocolRewardsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProtocolRewards *ProtocolRewardsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProtocolRewardsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProtocolRewards.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolRewardsUpgradedIterator{contract: _ProtocolRewards.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProtocolRewards *ProtocolRewardsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProtocolRewardsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProtocolRewards.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolRewardsUpgraded)
				if err := _ProtocolRewards.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ProtocolRewards *ProtocolRewardsFilterer) ParseUpgraded(log types.Log) (*ProtocolRewardsUpgraded, error) {
	event := new(ProtocolRewardsUpgraded)
	if err := _ProtocolRewards.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
