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

// AavegotchiInfo is an auto generated low-level Go binding around an user-defined struct.
type AavegotchiInfo struct {
	TokenId               *big.Int
	Name                  string
	Owner                 common.Address
	RandomNumber          *big.Int
	Status                *big.Int
	NumericTraits         [6]int16
	ModifiedNumericTraits [6]int16
	EquippedWearables     [16]uint16
	Collateral            common.Address
	Escrow                common.Address
	StakedAmount          *big.Int
	MinimumStake          *big.Int
	Kinship               *big.Int
	LastInteracted        *big.Int
	Experience            *big.Int
	ToNextLevel           *big.Int
	UsedSkillPoints       *big.Int
	Level                 *big.Int
	HauntId               *big.Int
	BaseRarityScore       *big.Int
	ModifiedRarityScore   *big.Int
	Locked                bool
	Items                 []ItemTypeIO
}

// Dimensions is an auto generated low-level Go binding around an user-defined struct.
type Dimensions struct {
	X      uint8
	Y      uint8
	Width  uint8
	Height uint8
}

// ERC721Listing is an auto generated low-level Go binding around an user-defined struct.
type ERC721Listing struct {
	ListingId          *big.Int
	Seller             common.Address
	Erc721TokenAddress common.Address
	Erc721TokenId      *big.Int
	Category           *big.Int
	PriceInWei         *big.Int
	TimeCreated        *big.Int
	TimePurchased      *big.Int
	Cancelled          bool
	PrincipalSplit     [2]uint16
	Affiliate          common.Address
}

// ERC721MarketplaceFacetAavegotchiListing is an auto generated low-level Go binding around an user-defined struct.
type ERC721MarketplaceFacetAavegotchiListing struct {
	Listing        ERC721Listing
	AavegotchiInfo AavegotchiInfo
}

// ERC721MarketplaceFacetCategory is an auto generated low-level Go binding around an user-defined struct.
type ERC721MarketplaceFacetCategory struct {
	Erc721TokenAddress common.Address
	Category           *big.Int
}

// ItemType is an auto generated low-level Go binding around an user-defined struct.
type ItemType struct {
	Name                string
	Description         string
	Author              string
	TraitModifiers      [6]int8
	SlotPositions       [16]bool
	AllowedCollaterals  []uint8
	Dimensions          Dimensions
	GhstPrice           *big.Int
	MaxQuantity         *big.Int
	TotalQuantity       *big.Int
	SvgId               uint32
	RarityScoreModifier uint8
	CanPurchaseWithGhst bool
	MinLevel            uint16
	CanBeTransferred    bool
	Category            uint8
	KinshipBonus        int16
	ExperienceBonus     uint32
}

// ItemTypeIO is an auto generated low-level Go binding around an user-defined struct.
type ItemTypeIO struct {
	Balance  *big.Int
	ItemId   *big.Int
	ItemType ItemType
}

// ERC721MarketplaceMetaData contains all meta data concerning the ERC721Marketplace contract.
var ERC721MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERC721ExecutedListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721ExecutedToRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"ERC721ListingAdd\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"name\":\"ERC721ListingSplit\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"}],\"name\":\"addERC721Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint16[2]\",\"name\":\"_principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"_affiliate\",\"type\":\"address\"}],\"name\":\"addERC721ListingWithSplit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"cancelERC721Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"cancelERC721ListingByToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_listingIds\",\"type\":\"uint256[]\"}],\"name\":\"cancelERC721Listings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"executeERC721Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"executeERC721ListingToRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"getAavegotchiListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing\",\"name\":\"listing_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"int16[6]\",\"name\":\"numericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"int16[6]\",\"name\":\"modifiedNumericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"uint16[16]\",\"name\":\"equippedWearables\",\"type\":\"uint16[16]\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kinship\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastInteracted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"experience\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNextLevel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedSkillPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hauntId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modifiedRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"int8[6]\",\"name\":\"traitModifiers\",\"type\":\"int8[6]\"},{\"internalType\":\"bool[16]\",\"name\":\"slotPositions\",\"type\":\"bool[16]\"},{\"internalType\":\"uint8[]\",\"name\":\"allowedCollaterals\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"x\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"}],\"internalType\":\"structDimensions\",\"name\":\"dimensions\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"ghstPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"svgId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"rarityScoreModifier\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"canPurchaseWithGhst\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"minLevel\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"canBeTransferred\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"int16\",\"name\":\"kinshipBonus\",\"type\":\"int16\"},{\"internalType\":\"uint32\",\"name\":\"experienceBonus\",\"type\":\"uint32\"}],\"internalType\":\"structItemType\",\"name\":\"itemType\",\"type\":\"tuple\"}],\"internalType\":\"structItemTypeIO[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structAavegotchiInfo\",\"name\":\"aavegotchiInfo_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getAavegotchiListings\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing\",\"name\":\"listing_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"int16[6]\",\"name\":\"numericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"int16[6]\",\"name\":\"modifiedNumericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"uint16[16]\",\"name\":\"equippedWearables\",\"type\":\"uint16[16]\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kinship\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastInteracted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"experience\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNextLevel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedSkillPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hauntId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modifiedRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"int8[6]\",\"name\":\"traitModifiers\",\"type\":\"int8[6]\"},{\"internalType\":\"bool[16]\",\"name\":\"slotPositions\",\"type\":\"bool[16]\"},{\"internalType\":\"uint8[]\",\"name\":\"allowedCollaterals\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"x\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"}],\"internalType\":\"structDimensions\",\"name\":\"dimensions\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"ghstPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"svgId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"rarityScoreModifier\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"canPurchaseWithGhst\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"minLevel\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"canBeTransferred\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"int16\",\"name\":\"kinshipBonus\",\"type\":\"int16\"},{\"internalType\":\"uint32\",\"name\":\"experienceBonus\",\"type\":\"uint32\"}],\"internalType\":\"structItemType\",\"name\":\"itemType\",\"type\":\"tuple\"}],\"internalType\":\"structItemTypeIO[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structAavegotchiInfo\",\"name\":\"aavegotchiInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structERC721MarketplaceFacet.AavegotchiListing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"}],\"name\":\"getERC721Category\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"category_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"getERC721Listing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing\",\"name\":\"listing_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getERC721ListingFromToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing\",\"name\":\"listing_\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getERC721Listings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getOwnerAavegotchiListings\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing\",\"name\":\"listing_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"int16[6]\",\"name\":\"numericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"int16[6]\",\"name\":\"modifiedNumericTraits\",\"type\":\"int16[6]\"},{\"internalType\":\"uint16[16]\",\"name\":\"equippedWearables\",\"type\":\"uint16[16]\"},{\"internalType\":\"address\",\"name\":\"collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"escrow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minimumStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kinship\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastInteracted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"experience\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toNextLevel\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usedSkillPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hauntId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"modifiedRarityScore\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"locked\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"author\",\"type\":\"string\"},{\"internalType\":\"int8[6]\",\"name\":\"traitModifiers\",\"type\":\"int8[6]\"},{\"internalType\":\"bool[16]\",\"name\":\"slotPositions\",\"type\":\"bool[16]\"},{\"internalType\":\"uint8[]\",\"name\":\"allowedCollaterals\",\"type\":\"uint8[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"x\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"y\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"width\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"height\",\"type\":\"uint8\"}],\"internalType\":\"structDimensions\",\"name\":\"dimensions\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"ghstPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalQuantity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"svgId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"rarityScoreModifier\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"canPurchaseWithGhst\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"minLevel\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"canBeTransferred\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"category\",\"type\":\"uint8\"},{\"internalType\":\"int16\",\"name\":\"kinshipBonus\",\"type\":\"int16\"},{\"internalType\":\"uint32\",\"name\":\"experienceBonus\",\"type\":\"uint32\"}],\"internalType\":\"structItemType\",\"name\":\"itemType\",\"type\":\"tuple\"}],\"internalType\":\"structItemTypeIO[]\",\"name\":\"items\",\"type\":\"tuple[]\"}],\"internalType\":\"structAavegotchiInfo\",\"name\":\"aavegotchiInfo_\",\"type\":\"tuple\"}],\"internalType\":\"structERC721MarketplaceFacet.AavegotchiListing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_category\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_sort\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"getOwnerERC721Listings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeCreated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timePurchased\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"internalType\":\"uint16[2]\",\"name\":\"principalSplit\",\"type\":\"uint16[2]\"},{\"internalType\":\"address\",\"name\":\"affiliate\",\"type\":\"address\"}],\"internalType\":\"structERC721Listing[]\",\"name\":\"listings_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"category\",\"type\":\"uint256\"}],\"internalType\":\"structERC721MarketplaceFacet.Category[]\",\"name\":\"_categories\",\"type\":\"tuple[]\"}],\"name\":\"setERC721Categories\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc721TokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_erc721TokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"updateERC721Listing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ERC721MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC721MarketplaceMetaData.ABI instead.
var ERC721MarketplaceABI = ERC721MarketplaceMetaData.ABI

// ERC721Marketplace is an auto generated Go binding around an Ethereum contract.
type ERC721Marketplace struct {
	ERC721MarketplaceCaller     // Read-only binding to the contract
	ERC721MarketplaceTransactor // Write-only binding to the contract
	ERC721MarketplaceFilterer   // Log filterer for contract events
}

// ERC721MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MarketplaceSession struct {
	Contract     *ERC721Marketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC721MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MarketplaceCallerSession struct {
	Contract *ERC721MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC721MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MarketplaceTransactorSession struct {
	Contract     *ERC721MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC721MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MarketplaceRaw struct {
	Contract *ERC721Marketplace // Generic contract binding to access the raw methods on
}

// ERC721MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MarketplaceCallerRaw struct {
	Contract *ERC721MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MarketplaceTransactorRaw struct {
	Contract *ERC721MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Marketplace creates a new instance of ERC721Marketplace, bound to a specific deployed contract.
func NewERC721Marketplace(address common.Address, backend bind.ContractBackend) (*ERC721Marketplace, error) {
	contract, err := bindERC721Marketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Marketplace{ERC721MarketplaceCaller: ERC721MarketplaceCaller{contract: contract}, ERC721MarketplaceTransactor: ERC721MarketplaceTransactor{contract: contract}, ERC721MarketplaceFilterer: ERC721MarketplaceFilterer{contract: contract}}, nil
}

// NewERC721MarketplaceCaller creates a new read-only instance of ERC721Marketplace, bound to a specific deployed contract.
func NewERC721MarketplaceCaller(address common.Address, caller bind.ContractCaller) (*ERC721MarketplaceCaller, error) {
	contract, err := bindERC721Marketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceCaller{contract: contract}, nil
}

// NewERC721MarketplaceTransactor creates a new write-only instance of ERC721Marketplace, bound to a specific deployed contract.
func NewERC721MarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MarketplaceTransactor, error) {
	contract, err := bindERC721Marketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceTransactor{contract: contract}, nil
}

// NewERC721MarketplaceFilterer creates a new log filterer instance of ERC721Marketplace, bound to a specific deployed contract.
func NewERC721MarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MarketplaceFilterer, error) {
	contract, err := bindERC721Marketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceFilterer{contract: contract}, nil
}

// bindERC721Marketplace binds a generic wrapper to an already deployed contract.
func bindERC721Marketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC721MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Marketplace *ERC721MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Marketplace.Contract.ERC721MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Marketplace *ERC721MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ERC721MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Marketplace *ERC721MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ERC721MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Marketplace *ERC721MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Marketplace *ERC721MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Marketplace *ERC721MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.contract.Transact(opts, method, params...)
}

// GetAavegotchiListing is a free data retrieval call binding the contract method 0xfb8b1c56.
//
// Solidity: function getAavegotchiListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_, (uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]) aavegotchiInfo_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetAavegotchiListing(opts *bind.CallOpts, _listingId *big.Int) (struct {
	Listing        ERC721Listing
	AavegotchiInfo AavegotchiInfo
}, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getAavegotchiListing", _listingId)

	outstruct := new(struct {
		Listing        ERC721Listing
		AavegotchiInfo AavegotchiInfo
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Listing = *abi.ConvertType(out[0], new(ERC721Listing)).(*ERC721Listing)
	outstruct.AavegotchiInfo = *abi.ConvertType(out[1], new(AavegotchiInfo)).(*AavegotchiInfo)

	return *outstruct, err

}

// GetAavegotchiListing is a free data retrieval call binding the contract method 0xfb8b1c56.
//
// Solidity: function getAavegotchiListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_, (uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]) aavegotchiInfo_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetAavegotchiListing(_listingId *big.Int) (struct {
	Listing        ERC721Listing
	AavegotchiInfo AavegotchiInfo
}, error) {
	return _ERC721Marketplace.Contract.GetAavegotchiListing(&_ERC721Marketplace.CallOpts, _listingId)
}

// GetAavegotchiListing is a free data retrieval call binding the contract method 0xfb8b1c56.
//
// Solidity: function getAavegotchiListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_, (uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]) aavegotchiInfo_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetAavegotchiListing(_listingId *big.Int) (struct {
	Listing        ERC721Listing
	AavegotchiInfo AavegotchiInfo
}, error) {
	return _ERC721Marketplace.Contract.GetAavegotchiListing(&_ERC721Marketplace.CallOpts, _listingId)
}

// GetAavegotchiListings is a free data retrieval call binding the contract method 0x107146a4.
//
// Solidity: function getAavegotchiListings(uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetAavegotchiListings(opts *bind.CallOpts, _category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getAavegotchiListings", _category, _sort, _length)

	if err != nil {
		return *new([]ERC721MarketplaceFacetAavegotchiListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC721MarketplaceFacetAavegotchiListing)).(*[]ERC721MarketplaceFacetAavegotchiListing)

	return out0, err

}

// GetAavegotchiListings is a free data retrieval call binding the contract method 0x107146a4.
//
// Solidity: function getAavegotchiListings(uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetAavegotchiListings(_category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	return _ERC721Marketplace.Contract.GetAavegotchiListings(&_ERC721Marketplace.CallOpts, _category, _sort, _length)
}

// GetAavegotchiListings is a free data retrieval call binding the contract method 0x107146a4.
//
// Solidity: function getAavegotchiListings(uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetAavegotchiListings(_category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	return _ERC721Marketplace.Contract.GetAavegotchiListings(&_ERC721Marketplace.CallOpts, _category, _sort, _length)
}

// GetERC721Category is a free data retrieval call binding the contract method 0xa676bafb.
//
// Solidity: function getERC721Category(address _erc721TokenAddress, uint256 _erc721TokenId) view returns(uint256 category_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetERC721Category(opts *bind.CallOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getERC721Category", _erc721TokenAddress, _erc721TokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetERC721Category is a free data retrieval call binding the contract method 0xa676bafb.
//
// Solidity: function getERC721Category(address _erc721TokenAddress, uint256 _erc721TokenId) view returns(uint256 category_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetERC721Category(_erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*big.Int, error) {
	return _ERC721Marketplace.Contract.GetERC721Category(&_ERC721Marketplace.CallOpts, _erc721TokenAddress, _erc721TokenId)
}

// GetERC721Category is a free data retrieval call binding the contract method 0xa676bafb.
//
// Solidity: function getERC721Category(address _erc721TokenAddress, uint256 _erc721TokenId) view returns(uint256 category_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetERC721Category(_erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*big.Int, error) {
	return _ERC721Marketplace.Contract.GetERC721Category(&_ERC721Marketplace.CallOpts, _erc721TokenAddress, _erc721TokenId)
}

// GetERC721Listing is a free data retrieval call binding the contract method 0x4eff5a3e.
//
// Solidity: function getERC721Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetERC721Listing(opts *bind.CallOpts, _listingId *big.Int) (ERC721Listing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getERC721Listing", _listingId)

	if err != nil {
		return *new(ERC721Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721Listing)).(*ERC721Listing)

	return out0, err

}

// GetERC721Listing is a free data retrieval call binding the contract method 0x4eff5a3e.
//
// Solidity: function getERC721Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetERC721Listing(_listingId *big.Int) (ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721Listing(&_ERC721Marketplace.CallOpts, _listingId)
}

// GetERC721Listing is a free data retrieval call binding the contract method 0x4eff5a3e.
//
// Solidity: function getERC721Listing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetERC721Listing(_listingId *big.Int) (ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721Listing(&_ERC721Marketplace.CallOpts, _listingId)
}

// GetERC721ListingFromToken is a free data retrieval call binding the contract method 0xd33a32df.
//
// Solidity: function getERC721ListingFromToken(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetERC721ListingFromToken(opts *bind.CallOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (ERC721Listing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getERC721ListingFromToken", _erc721TokenAddress, _erc721TokenId, _owner)

	if err != nil {
		return *new(ERC721Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(ERC721Listing)).(*ERC721Listing)

	return out0, err

}

// GetERC721ListingFromToken is a free data retrieval call binding the contract method 0xd33a32df.
//
// Solidity: function getERC721ListingFromToken(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetERC721ListingFromToken(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721ListingFromToken(&_ERC721Marketplace.CallOpts, _erc721TokenAddress, _erc721TokenId, _owner)
}

// GetERC721ListingFromToken is a free data retrieval call binding the contract method 0xd33a32df.
//
// Solidity: function getERC721ListingFromToken(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address) listing_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetERC721ListingFromToken(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721ListingFromToken(&_ERC721Marketplace.CallOpts, _erc721TokenAddress, _erc721TokenId, _owner)
}

// GetERC721Listings is a free data retrieval call binding the contract method 0xd90da106.
//
// Solidity: function getERC721Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetERC721Listings(opts *bind.CallOpts, _category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getERC721Listings", _category, _sort, _length)

	if err != nil {
		return *new([]ERC721Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC721Listing)).(*[]ERC721Listing)

	return out0, err

}

// GetERC721Listings is a free data retrieval call binding the contract method 0xd90da106.
//
// Solidity: function getERC721Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetERC721Listings(_category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721Listings(&_ERC721Marketplace.CallOpts, _category, _sort, _length)
}

// GetERC721Listings is a free data retrieval call binding the contract method 0xd90da106.
//
// Solidity: function getERC721Listings(uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetERC721Listings(_category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetERC721Listings(&_ERC721Marketplace.CallOpts, _category, _sort, _length)
}

// GetOwnerAavegotchiListings is a free data retrieval call binding the contract method 0x88358431.
//
// Solidity: function getOwnerAavegotchiListings(address _owner, uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetOwnerAavegotchiListings(opts *bind.CallOpts, _owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getOwnerAavegotchiListings", _owner, _category, _sort, _length)

	if err != nil {
		return *new([]ERC721MarketplaceFacetAavegotchiListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC721MarketplaceFacetAavegotchiListing)).(*[]ERC721MarketplaceFacetAavegotchiListing)

	return out0, err

}

// GetOwnerAavegotchiListings is a free data retrieval call binding the contract method 0x88358431.
//
// Solidity: function getOwnerAavegotchiListings(address _owner, uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetOwnerAavegotchiListings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	return _ERC721Marketplace.Contract.GetOwnerAavegotchiListings(&_ERC721Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// GetOwnerAavegotchiListings is a free data retrieval call binding the contract method 0x88358431.
//
// Solidity: function getOwnerAavegotchiListings(address _owner, uint256 _category, string _sort, uint256 _length) view returns(((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address),(uint256,string,address,uint256,uint256,int16[6],int16[6],uint16[16],address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,(uint256,uint256,(string,string,string,int8[6],bool[16],uint8[],(uint8,uint8,uint8,uint8),uint256,uint256,uint256,uint32,uint8,bool,uint16,bool,uint8,int16,uint32))[]))[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetOwnerAavegotchiListings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721MarketplaceFacetAavegotchiListing, error) {
	return _ERC721Marketplace.Contract.GetOwnerAavegotchiListings(&_ERC721Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// GetOwnerERC721Listings is a free data retrieval call binding the contract method 0x6c20642f.
//
// Solidity: function getOwnerERC721Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCaller) GetOwnerERC721Listings(opts *bind.CallOpts, _owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	var out []interface{}
	err := _ERC721Marketplace.contract.Call(opts, &out, "getOwnerERC721Listings", _owner, _category, _sort, _length)

	if err != nil {
		return *new([]ERC721Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]ERC721Listing)).(*[]ERC721Listing)

	return out0, err

}

// GetOwnerERC721Listings is a free data retrieval call binding the contract method 0x6c20642f.
//
// Solidity: function getOwnerERC721Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceSession) GetOwnerERC721Listings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetOwnerERC721Listings(&_ERC721Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// GetOwnerERC721Listings is a free data retrieval call binding the contract method 0x6c20642f.
//
// Solidity: function getOwnerERC721Listings(address _owner, uint256 _category, string _sort, uint256 _length) view returns((uint256,address,address,uint256,uint256,uint256,uint256,uint256,bool,uint16[2],address)[] listings_)
func (_ERC721Marketplace *ERC721MarketplaceCallerSession) GetOwnerERC721Listings(_owner common.Address, _category *big.Int, _sort string, _length *big.Int) ([]ERC721Listing, error) {
	return _ERC721Marketplace.Contract.GetOwnerERC721Listings(&_ERC721Marketplace.CallOpts, _owner, _category, _sort, _length)
}

// AddERC721Listing is a paid mutator transaction binding the contract method 0x31b0b514.
//
// Solidity: function addERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) AddERC721Listing(opts *bind.TransactOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "addERC721Listing", _erc721TokenAddress, _erc721TokenId, _priceInWei)
}

// AddERC721Listing is a paid mutator transaction binding the contract method 0x31b0b514.
//
// Solidity: function addERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) AddERC721Listing(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.AddERC721Listing(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _priceInWei)
}

// AddERC721Listing is a paid mutator transaction binding the contract method 0x31b0b514.
//
// Solidity: function addERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) AddERC721Listing(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.AddERC721Listing(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _priceInWei)
}

// AddERC721ListingWithSplit is a paid mutator transaction binding the contract method 0x2170d391.
//
// Solidity: function addERC721ListingWithSplit(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) AddERC721ListingWithSplit(opts *bind.TransactOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "addERC721ListingWithSplit", _erc721TokenAddress, _erc721TokenId, _priceInWei, _principalSplit, _affiliate)
}

// AddERC721ListingWithSplit is a paid mutator transaction binding the contract method 0x2170d391.
//
// Solidity: function addERC721ListingWithSplit(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) AddERC721ListingWithSplit(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.AddERC721ListingWithSplit(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _priceInWei, _principalSplit, _affiliate)
}

// AddERC721ListingWithSplit is a paid mutator transaction binding the contract method 0x2170d391.
//
// Solidity: function addERC721ListingWithSplit(address _erc721TokenAddress, uint256 _erc721TokenId, uint256 _priceInWei, uint16[2] _principalSplit, address _affiliate) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) AddERC721ListingWithSplit(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _priceInWei *big.Int, _principalSplit [2]uint16, _affiliate common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.AddERC721ListingWithSplit(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _priceInWei, _principalSplit, _affiliate)
}

// CancelERC721Listing is a paid mutator transaction binding the contract method 0xdf3ad6d1.
//
// Solidity: function cancelERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) CancelERC721Listing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "cancelERC721Listing", _listingId)
}

// CancelERC721Listing is a paid mutator transaction binding the contract method 0xdf3ad6d1.
//
// Solidity: function cancelERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) CancelERC721Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721Listing(&_ERC721Marketplace.TransactOpts, _listingId)
}

// CancelERC721Listing is a paid mutator transaction binding the contract method 0xdf3ad6d1.
//
// Solidity: function cancelERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) CancelERC721Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721Listing(&_ERC721Marketplace.TransactOpts, _listingId)
}

// CancelERC721ListingByToken is a paid mutator transaction binding the contract method 0xba726be4.
//
// Solidity: function cancelERC721ListingByToken(address _erc721TokenAddress, uint256 _erc721TokenId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) CancelERC721ListingByToken(opts *bind.TransactOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "cancelERC721ListingByToken", _erc721TokenAddress, _erc721TokenId)
}

// CancelERC721ListingByToken is a paid mutator transaction binding the contract method 0xba726be4.
//
// Solidity: function cancelERC721ListingByToken(address _erc721TokenAddress, uint256 _erc721TokenId) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) CancelERC721ListingByToken(_erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721ListingByToken(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId)
}

// CancelERC721ListingByToken is a paid mutator transaction binding the contract method 0xba726be4.
//
// Solidity: function cancelERC721ListingByToken(address _erc721TokenAddress, uint256 _erc721TokenId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) CancelERC721ListingByToken(_erc721TokenAddress common.Address, _erc721TokenId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721ListingByToken(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId)
}

// CancelERC721Listings is a paid mutator transaction binding the contract method 0x0adfab17.
//
// Solidity: function cancelERC721Listings(uint256[] _listingIds) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) CancelERC721Listings(opts *bind.TransactOpts, _listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "cancelERC721Listings", _listingIds)
}

// CancelERC721Listings is a paid mutator transaction binding the contract method 0x0adfab17.
//
// Solidity: function cancelERC721Listings(uint256[] _listingIds) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) CancelERC721Listings(_listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721Listings(&_ERC721Marketplace.TransactOpts, _listingIds)
}

// CancelERC721Listings is a paid mutator transaction binding the contract method 0x0adfab17.
//
// Solidity: function cancelERC721Listings(uint256[] _listingIds) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) CancelERC721Listings(_listingIds []*big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.CancelERC721Listings(&_ERC721Marketplace.TransactOpts, _listingIds)
}

// ExecuteERC721Listing is a paid mutator transaction binding the contract method 0xf17a5c6c.
//
// Solidity: function executeERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) ExecuteERC721Listing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "executeERC721Listing", _listingId)
}

// ExecuteERC721Listing is a paid mutator transaction binding the contract method 0xf17a5c6c.
//
// Solidity: function executeERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) ExecuteERC721Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ExecuteERC721Listing(&_ERC721Marketplace.TransactOpts, _listingId)
}

// ExecuteERC721Listing is a paid mutator transaction binding the contract method 0xf17a5c6c.
//
// Solidity: function executeERC721Listing(uint256 _listingId) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) ExecuteERC721Listing(_listingId *big.Int) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ExecuteERC721Listing(&_ERC721Marketplace.TransactOpts, _listingId)
}

// ExecuteERC721ListingToRecipient is a paid mutator transaction binding the contract method 0x66609c88.
//
// Solidity: function executeERC721ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _priceInWei, uint256 _tokenId, address _recipient) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) ExecuteERC721ListingToRecipient(opts *bind.TransactOpts, _listingId *big.Int, _contractAddress common.Address, _priceInWei *big.Int, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "executeERC721ListingToRecipient", _listingId, _contractAddress, _priceInWei, _tokenId, _recipient)
}

// ExecuteERC721ListingToRecipient is a paid mutator transaction binding the contract method 0x66609c88.
//
// Solidity: function executeERC721ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _priceInWei, uint256 _tokenId, address _recipient) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) ExecuteERC721ListingToRecipient(_listingId *big.Int, _contractAddress common.Address, _priceInWei *big.Int, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ExecuteERC721ListingToRecipient(&_ERC721Marketplace.TransactOpts, _listingId, _contractAddress, _priceInWei, _tokenId, _recipient)
}

// ExecuteERC721ListingToRecipient is a paid mutator transaction binding the contract method 0x66609c88.
//
// Solidity: function executeERC721ListingToRecipient(uint256 _listingId, address _contractAddress, uint256 _priceInWei, uint256 _tokenId, address _recipient) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) ExecuteERC721ListingToRecipient(_listingId *big.Int, _contractAddress common.Address, _priceInWei *big.Int, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.ExecuteERC721ListingToRecipient(&_ERC721Marketplace.TransactOpts, _listingId, _contractAddress, _priceInWei, _tokenId, _recipient)
}

// SetERC721Categories is a paid mutator transaction binding the contract method 0x53922ca5.
//
// Solidity: function setERC721Categories((address,uint256)[] _categories) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) SetERC721Categories(opts *bind.TransactOpts, _categories []ERC721MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "setERC721Categories", _categories)
}

// SetERC721Categories is a paid mutator transaction binding the contract method 0x53922ca5.
//
// Solidity: function setERC721Categories((address,uint256)[] _categories) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) SetERC721Categories(_categories []ERC721MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.SetERC721Categories(&_ERC721Marketplace.TransactOpts, _categories)
}

// SetERC721Categories is a paid mutator transaction binding the contract method 0x53922ca5.
//
// Solidity: function setERC721Categories((address,uint256)[] _categories) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) SetERC721Categories(_categories []ERC721MarketplaceFacetCategory) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.SetERC721Categories(&_ERC721Marketplace.TransactOpts, _categories)
}

// UpdateERC721Listing is a paid mutator transaction binding the contract method 0x3eed635a.
//
// Solidity: function updateERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactor) UpdateERC721Listing(opts *bind.TransactOpts, _erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.contract.Transact(opts, "updateERC721Listing", _erc721TokenAddress, _erc721TokenId, _owner)
}

// UpdateERC721Listing is a paid mutator transaction binding the contract method 0x3eed635a.
//
// Solidity: function updateERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) returns()
func (_ERC721Marketplace *ERC721MarketplaceSession) UpdateERC721Listing(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.UpdateERC721Listing(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _owner)
}

// UpdateERC721Listing is a paid mutator transaction binding the contract method 0x3eed635a.
//
// Solidity: function updateERC721Listing(address _erc721TokenAddress, uint256 _erc721TokenId, address _owner) returns()
func (_ERC721Marketplace *ERC721MarketplaceTransactorSession) UpdateERC721Listing(_erc721TokenAddress common.Address, _erc721TokenId *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _ERC721Marketplace.Contract.UpdateERC721Listing(&_ERC721Marketplace.TransactOpts, _erc721TokenAddress, _erc721TokenId, _owner)
}

// ERC721MarketplaceERC721ExecutedListingIterator is returned from FilterERC721ExecutedListing and is used to iterate over the raw logs and unpacked data for ERC721ExecutedListing events raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ExecutedListingIterator struct {
	Event *ERC721MarketplaceERC721ExecutedListing // Event containing the contract specifics and raw log

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
func (it *ERC721MarketplaceERC721ExecutedListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MarketplaceERC721ExecutedListing)
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
		it.Event = new(ERC721MarketplaceERC721ExecutedListing)
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
func (it *ERC721MarketplaceERC721ExecutedListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MarketplaceERC721ExecutedListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MarketplaceERC721ExecutedListing represents a ERC721ExecutedListing event raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ExecutedListing struct {
	ListingId          *big.Int
	Seller             common.Address
	Buyer              common.Address
	Erc721TokenAddress common.Address
	Erc721TokenId      *big.Int
	Category           *big.Int
	PriceInWei         *big.Int
	Time               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterERC721ExecutedListing is a free log retrieval operation binding the contract event 0x8abb27e53fc4ab514e7e5521910bdf2d1e48a69f150dbb671efa7ccc8f44bb28.
//
// Solidity: event ERC721ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 priceInWei, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) FilterERC721ExecutedListing(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, category []*big.Int) (*ERC721MarketplaceERC721ExecutedListingIterator, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.FilterLogs(opts, "ERC721ExecutedListing", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceERC721ExecutedListingIterator{contract: _ERC721Marketplace.contract, event: "ERC721ExecutedListing", logs: logs, sub: sub}, nil
}

// WatchERC721ExecutedListing is a free log subscription operation binding the contract event 0x8abb27e53fc4ab514e7e5521910bdf2d1e48a69f150dbb671efa7ccc8f44bb28.
//
// Solidity: event ERC721ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 priceInWei, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) WatchERC721ExecutedListing(opts *bind.WatchOpts, sink chan<- *ERC721MarketplaceERC721ExecutedListing, listingId []*big.Int, seller []common.Address, category []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.WatchLogs(opts, "ERC721ExecutedListing", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MarketplaceERC721ExecutedListing)
				if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ExecutedListing", log); err != nil {
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

// ParseERC721ExecutedListing is a log parse operation binding the contract event 0x8abb27e53fc4ab514e7e5521910bdf2d1e48a69f150dbb671efa7ccc8f44bb28.
//
// Solidity: event ERC721ExecutedListing(uint256 indexed listingId, address indexed seller, address buyer, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 priceInWei, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) ParseERC721ExecutedListing(log types.Log) (*ERC721MarketplaceERC721ExecutedListing, error) {
	event := new(ERC721MarketplaceERC721ExecutedListing)
	if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ExecutedListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MarketplaceERC721ExecutedToRecipientIterator is returned from FilterERC721ExecutedToRecipient and is used to iterate over the raw logs and unpacked data for ERC721ExecutedToRecipient events raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ExecutedToRecipientIterator struct {
	Event *ERC721MarketplaceERC721ExecutedToRecipient // Event containing the contract specifics and raw log

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
func (it *ERC721MarketplaceERC721ExecutedToRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MarketplaceERC721ExecutedToRecipient)
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
		it.Event = new(ERC721MarketplaceERC721ExecutedToRecipient)
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
func (it *ERC721MarketplaceERC721ExecutedToRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MarketplaceERC721ExecutedToRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MarketplaceERC721ExecutedToRecipient represents a ERC721ExecutedToRecipient event raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ExecutedToRecipient struct {
	ListingId *big.Int
	Buyer     common.Address
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721ExecutedToRecipient is a free log retrieval operation binding the contract event 0x0f52fde6171d45c02579364597406652690bf2264e564ba1a5b88ffa98d8815b.
//
// Solidity: event ERC721ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) FilterERC721ExecutedToRecipient(opts *bind.FilterOpts, listingId []*big.Int, buyer []common.Address, recipient []common.Address) (*ERC721MarketplaceERC721ExecutedToRecipientIterator, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.FilterLogs(opts, "ERC721ExecutedToRecipient", listingIdRule, buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceERC721ExecutedToRecipientIterator{contract: _ERC721Marketplace.contract, event: "ERC721ExecutedToRecipient", logs: logs, sub: sub}, nil
}

// WatchERC721ExecutedToRecipient is a free log subscription operation binding the contract event 0x0f52fde6171d45c02579364597406652690bf2264e564ba1a5b88ffa98d8815b.
//
// Solidity: event ERC721ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) WatchERC721ExecutedToRecipient(opts *bind.WatchOpts, sink chan<- *ERC721MarketplaceERC721ExecutedToRecipient, listingId []*big.Int, buyer []common.Address, recipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.WatchLogs(opts, "ERC721ExecutedToRecipient", listingIdRule, buyerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MarketplaceERC721ExecutedToRecipient)
				if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ExecutedToRecipient", log); err != nil {
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

// ParseERC721ExecutedToRecipient is a log parse operation binding the contract event 0x0f52fde6171d45c02579364597406652690bf2264e564ba1a5b88ffa98d8815b.
//
// Solidity: event ERC721ExecutedToRecipient(uint256 indexed listingId, address indexed buyer, address indexed recipient)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) ParseERC721ExecutedToRecipient(log types.Log) (*ERC721MarketplaceERC721ExecutedToRecipient, error) {
	event := new(ERC721MarketplaceERC721ExecutedToRecipient)
	if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ExecutedToRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MarketplaceERC721ListingAddIterator is returned from FilterERC721ListingAdd and is used to iterate over the raw logs and unpacked data for ERC721ListingAdd events raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ListingAddIterator struct {
	Event *ERC721MarketplaceERC721ListingAdd // Event containing the contract specifics and raw log

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
func (it *ERC721MarketplaceERC721ListingAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MarketplaceERC721ListingAdd)
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
		it.Event = new(ERC721MarketplaceERC721ListingAdd)
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
func (it *ERC721MarketplaceERC721ListingAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MarketplaceERC721ListingAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MarketplaceERC721ListingAdd represents a ERC721ListingAdd event raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ListingAdd struct {
	ListingId          *big.Int
	Seller             common.Address
	Erc721TokenAddress common.Address
	Erc721TokenId      *big.Int
	Category           *big.Int
	Time               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterERC721ListingAdd is a free log retrieval operation binding the contract event 0xd9b541cd1d9650dbbcb9b244efe4aa9455a9e7bb329f80cd30ab25275534c72d.
//
// Solidity: event ERC721ListingAdd(uint256 indexed listingId, address indexed seller, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) FilterERC721ListingAdd(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address, category []*big.Int) (*ERC721MarketplaceERC721ListingAddIterator, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.FilterLogs(opts, "ERC721ListingAdd", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceERC721ListingAddIterator{contract: _ERC721Marketplace.contract, event: "ERC721ListingAdd", logs: logs, sub: sub}, nil
}

// WatchERC721ListingAdd is a free log subscription operation binding the contract event 0xd9b541cd1d9650dbbcb9b244efe4aa9455a9e7bb329f80cd30ab25275534c72d.
//
// Solidity: event ERC721ListingAdd(uint256 indexed listingId, address indexed seller, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) WatchERC721ListingAdd(opts *bind.WatchOpts, sink chan<- *ERC721MarketplaceERC721ListingAdd, listingId []*big.Int, seller []common.Address, category []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _ERC721Marketplace.contract.WatchLogs(opts, "ERC721ListingAdd", listingIdRule, sellerRule, categoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MarketplaceERC721ListingAdd)
				if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ListingAdd", log); err != nil {
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

// ParseERC721ListingAdd is a log parse operation binding the contract event 0xd9b541cd1d9650dbbcb9b244efe4aa9455a9e7bb329f80cd30ab25275534c72d.
//
// Solidity: event ERC721ListingAdd(uint256 indexed listingId, address indexed seller, address erc721TokenAddress, uint256 erc721TokenId, uint256 indexed category, uint256 time)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) ParseERC721ListingAdd(log types.Log) (*ERC721MarketplaceERC721ListingAdd, error) {
	event := new(ERC721MarketplaceERC721ListingAdd)
	if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ListingAdd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721MarketplaceERC721ListingSplitIterator is returned from FilterERC721ListingSplit and is used to iterate over the raw logs and unpacked data for ERC721ListingSplit events raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ListingSplitIterator struct {
	Event *ERC721MarketplaceERC721ListingSplit // Event containing the contract specifics and raw log

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
func (it *ERC721MarketplaceERC721ListingSplitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721MarketplaceERC721ListingSplit)
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
		it.Event = new(ERC721MarketplaceERC721ListingSplit)
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
func (it *ERC721MarketplaceERC721ListingSplitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721MarketplaceERC721ListingSplitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721MarketplaceERC721ListingSplit represents a ERC721ListingSplit event raised by the ERC721Marketplace contract.
type ERC721MarketplaceERC721ListingSplit struct {
	ListingId      *big.Int
	PrincipalSplit [2]uint16
	Affiliate      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterERC721ListingSplit is a free log retrieval operation binding the contract event 0x578ffc272d43e4ca7465fd44d471e0494a2b32e0caa226e1b9e61bbdc54fc96b.
//
// Solidity: event ERC721ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) FilterERC721ListingSplit(opts *bind.FilterOpts, listingId []*big.Int) (*ERC721MarketplaceERC721ListingSplitIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC721Marketplace.contract.FilterLogs(opts, "ERC721ListingSplit", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721MarketplaceERC721ListingSplitIterator{contract: _ERC721Marketplace.contract, event: "ERC721ListingSplit", logs: logs, sub: sub}, nil
}

// WatchERC721ListingSplit is a free log subscription operation binding the contract event 0x578ffc272d43e4ca7465fd44d471e0494a2b32e0caa226e1b9e61bbdc54fc96b.
//
// Solidity: event ERC721ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) WatchERC721ListingSplit(opts *bind.WatchOpts, sink chan<- *ERC721MarketplaceERC721ListingSplit, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _ERC721Marketplace.contract.WatchLogs(opts, "ERC721ListingSplit", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721MarketplaceERC721ListingSplit)
				if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ListingSplit", log); err != nil {
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

// ParseERC721ListingSplit is a log parse operation binding the contract event 0x578ffc272d43e4ca7465fd44d471e0494a2b32e0caa226e1b9e61bbdc54fc96b.
//
// Solidity: event ERC721ListingSplit(uint256 indexed listingId, uint16[2] principalSplit, address affiliate)
func (_ERC721Marketplace *ERC721MarketplaceFilterer) ParseERC721ListingSplit(log types.Log) (*ERC721MarketplaceERC721ListingSplit, error) {
	event := new(ERC721MarketplaceERC721ListingSplit)
	if err := _ERC721Marketplace.contract.UnpackLog(event, "ERC721ListingSplit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
