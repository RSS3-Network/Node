// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package benddao

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

// DataTypesNftConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftConfigurationMap struct {
	Data *big.Int
}

// DataTypesNftData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNftData struct {
	Configuration DataTypesNftConfigurationMap
	BNftAddress   common.Address
	Id            uint8
	MaxSupply     *big.Int
	MaxTokenId    *big.Int
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration             DataTypesReserveConfigurationMap
	LiquidityIndex            *big.Int
	VariableBorrowIndex       *big.Int
	CurrentLiquidityRate      *big.Int
	CurrentVariableBorrowRate *big.Int
	LastUpdateTimestamp       *big.Int
	BTokenAddress             common.Address
	DebtTokenAddress          common.Address
	InterestRateAddress       common.Address
	Id                        uint8
}

// LendPoolMetaData contains all meta data concerning the LendPool contract.
var LendPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Auction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Liquidate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"PausedTimeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fineAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"auction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"batchBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nftAssets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"nftTokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchRepay\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromBefore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvider\",\"outputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfNfts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxNumberOfReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidderAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidBorrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftAuctionEndTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemEndTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"}],\"name\":\"getNftCollateralData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsInReserve\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getNftData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.NftData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftDebtData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserveAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"name\":\"getNftLiquidatePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidatePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"paybackAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNftsList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPausedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bNftAddress\",\"type\":\"address\"}],\"name\":\"initNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILendPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFine\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAsset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfNfts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setMaxNumberOfReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setNftConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTokenId\",\"type\":\"uint256\"}],\"name\":\"setNftMaxSupplyAndTokenId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"durationTime\",\"type\":\"uint256\"}],\"name\":\"setPausedTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setReserveConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LendPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use LendPoolMetaData.ABI instead.
var LendPoolABI = LendPoolMetaData.ABI

// LendPool is an auto generated Go binding around an Ethereum contract.
type LendPool struct {
	LendPoolCaller     // Read-only binding to the contract
	LendPoolTransactor // Write-only binding to the contract
	LendPoolFilterer   // Log filterer for contract events
}

// LendPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendPoolSession struct {
	Contract     *LendPool         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendPoolCallerSession struct {
	Contract *LendPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// LendPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendPoolTransactorSession struct {
	Contract     *LendPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// LendPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendPoolRaw struct {
	Contract *LendPool // Generic contract binding to access the raw methods on
}

// LendPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendPoolCallerRaw struct {
	Contract *LendPoolCaller // Generic read-only contract binding to access the raw methods on
}

// LendPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendPoolTransactorRaw struct {
	Contract *LendPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendPool creates a new instance of LendPool, bound to a specific deployed contract.
func NewLendPool(address common.Address, backend bind.ContractBackend) (*LendPool, error) {
	contract, err := bindLendPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LendPool{LendPoolCaller: LendPoolCaller{contract: contract}, LendPoolTransactor: LendPoolTransactor{contract: contract}, LendPoolFilterer: LendPoolFilterer{contract: contract}}, nil
}

// NewLendPoolCaller creates a new read-only instance of LendPool, bound to a specific deployed contract.
func NewLendPoolCaller(address common.Address, caller bind.ContractCaller) (*LendPoolCaller, error) {
	contract, err := bindLendPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendPoolCaller{contract: contract}, nil
}

// NewLendPoolTransactor creates a new write-only instance of LendPool, bound to a specific deployed contract.
func NewLendPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LendPoolTransactor, error) {
	contract, err := bindLendPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendPoolTransactor{contract: contract}, nil
}

// NewLendPoolFilterer creates a new log filterer instance of LendPool, bound to a specific deployed contract.
func NewLendPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LendPoolFilterer, error) {
	contract, err := bindLendPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendPoolFilterer{contract: contract}, nil
}

// bindLendPool binds a generic wrapper to an already deployed contract.
func bindLendPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LendPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendPool *LendPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendPool.Contract.LendPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendPool *LendPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendPool.Contract.LendPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendPool *LendPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendPool.Contract.LendPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LendPool *LendPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LendPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LendPool *LendPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LendPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LendPool *LendPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LendPool.Contract.contract.Transact(opts, method, params...)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_LendPool *LendPoolCaller) FinalizeTransfer(opts *bind.CallOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "finalizeTransfer", asset, from, to, amount, balanceFromBefore, balanceToBefore)

	if err != nil {
		return err
	}

	return err

}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_LendPool *LendPoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _LendPool.Contract.FinalizeTransfer(&_LendPool.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// FinalizeTransfer is a free data retrieval call binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromBefore, uint256 balanceToBefore) view returns()
func (_LendPool *LendPoolCallerSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromBefore *big.Int, balanceToBefore *big.Int) error {
	return _LendPool.Contract.FinalizeTransfer(&_LendPool.CallOpts, asset, from, to, amount, balanceFromBefore, balanceToBefore)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_LendPool *LendPoolCaller) GetAddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getAddressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_LendPool *LendPoolSession) GetAddressesProvider() (common.Address, error) {
	return _LendPool.Contract.GetAddressesProvider(&_LendPool.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_LendPool *LendPoolCallerSession) GetAddressesProvider() (common.Address, error) {
	return _LendPool.Contract.GetAddressesProvider(&_LendPool.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_LendPool *LendPoolCaller) GetMaxNumberOfNfts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getMaxNumberOfNfts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_LendPool *LendPoolSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _LendPool.Contract.GetMaxNumberOfNfts(&_LendPool.CallOpts)
}

// GetMaxNumberOfNfts is a free data retrieval call binding the contract method 0xdd90ff38.
//
// Solidity: function getMaxNumberOfNfts() view returns(uint256)
func (_LendPool *LendPoolCallerSession) GetMaxNumberOfNfts() (*big.Int, error) {
	return _LendPool.Contract.GetMaxNumberOfNfts(&_LendPool.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_LendPool *LendPoolCaller) GetMaxNumberOfReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getMaxNumberOfReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_LendPool *LendPoolSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _LendPool.Contract.GetMaxNumberOfReserves(&_LendPool.CallOpts)
}

// GetMaxNumberOfReserves is a free data retrieval call binding the contract method 0x08ac08b9.
//
// Solidity: function getMaxNumberOfReserves() view returns(uint256)
func (_LendPool *LendPoolCallerSession) GetMaxNumberOfReserves() (*big.Int, error) {
	return _LendPool.Contract.GetMaxNumberOfReserves(&_LendPool.CallOpts)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_LendPool *LendPoolCaller) GetNftAuctionData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftAuctionData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId          *big.Int
		BidderAddress   common.Address
		BidPrice        *big.Int
		BidBorrowAmount *big.Int
		BidFine         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidderAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BidPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.BidBorrowAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.BidFine = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_LendPool *LendPoolSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _LendPool.Contract.GetNftAuctionData(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionData is a free data retrieval call binding the contract method 0xe5bceca5.
//
// Solidity: function getNftAuctionData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address bidderAddress, uint256 bidPrice, uint256 bidBorrowAmount, uint256 bidFine)
func (_LendPool *LendPoolCallerSession) GetNftAuctionData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId          *big.Int
	BidderAddress   common.Address
	BidPrice        *big.Int
	BidBorrowAmount *big.Int
	BidFine         *big.Int
}, error) {
	return _LendPool.Contract.GetNftAuctionData(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_LendPool *LendPoolCaller) GetNftAuctionEndTime(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftAuctionEndTime", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId             *big.Int
		BidStartTimestamp  *big.Int
		BidEndTimestamp    *big.Int
		RedeemEndTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BidStartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BidEndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RedeemEndTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_LendPool *LendPoolSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _LendPool.Contract.GetNftAuctionEndTime(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftAuctionEndTime is a free data retrieval call binding the contract method 0x17c8a456.
//
// Solidity: function getNftAuctionEndTime(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, uint256 bidStartTimestamp, uint256 bidEndTimestamp, uint256 redeemEndTimestamp)
func (_LendPool *LendPoolCallerSession) GetNftAuctionEndTime(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId             *big.Int
	BidStartTimestamp  *big.Int
	BidEndTimestamp    *big.Int
	RedeemEndTimestamp *big.Int
}, error) {
	return _LendPool.Contract.GetNftAuctionEndTime(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_LendPool *LendPoolCaller) GetNftCollateralData(opts *bind.CallOpts, nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftCollateralData", nftAsset, reserveAsset)

	outstruct := new(struct {
		TotalCollateralInETH      *big.Int
		TotalCollateralInReserve  *big.Int
		AvailableBorrowsInETH     *big.Int
		AvailableBorrowsInReserve *big.Int
		Ltv                       *big.Int
		LiquidationThreshold      *big.Int
		LiquidationBonus          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralInETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalCollateralInReserve = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsInReserve = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_LendPool *LendPoolSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _LendPool.Contract.GetNftCollateralData(&_LendPool.CallOpts, nftAsset, reserveAsset)
}

// GetNftCollateralData is a free data retrieval call binding the contract method 0xcc8ccdf2.
//
// Solidity: function getNftCollateralData(address nftAsset, address reserveAsset) view returns(uint256 totalCollateralInETH, uint256 totalCollateralInReserve, uint256 availableBorrowsInETH, uint256 availableBorrowsInReserve, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus)
func (_LendPool *LendPoolCallerSession) GetNftCollateralData(nftAsset common.Address, reserveAsset common.Address) (struct {
	TotalCollateralInETH      *big.Int
	TotalCollateralInReserve  *big.Int
	AvailableBorrowsInETH     *big.Int
	AvailableBorrowsInReserve *big.Int
	Ltv                       *big.Int
	LiquidationThreshold      *big.Int
	LiquidationBonus          *big.Int
}, error) {
	return _LendPool.Contract.GetNftCollateralData(&_LendPool.CallOpts, nftAsset, reserveAsset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolCaller) GetNftConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesNftConfigurationMap, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftConfiguration", asset)

	if err != nil {
		return *new(DataTypesNftConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftConfigurationMap)).(*DataTypesNftConfigurationMap)

	return out0, err

}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _LendPool.Contract.GetNftConfiguration(&_LendPool.CallOpts, asset)
}

// GetNftConfiguration is a free data retrieval call binding the contract method 0x87c32dec.
//
// Solidity: function getNftConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolCallerSession) GetNftConfiguration(asset common.Address) (DataTypesNftConfigurationMap, error) {
	return _LendPool.Contract.GetNftConfiguration(&_LendPool.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_LendPool *LendPoolCaller) GetNftData(opts *bind.CallOpts, asset common.Address) (DataTypesNftData, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftData", asset)

	if err != nil {
		return *new(DataTypesNftData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNftData)).(*DataTypesNftData)

	return out0, err

}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_LendPool *LendPoolSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _LendPool.Contract.GetNftData(&_LendPool.CallOpts, asset)
}

// GetNftData is a free data retrieval call binding the contract method 0x77bdc0c3.
//
// Solidity: function getNftData(address asset) view returns(((uint256),address,uint8,uint256,uint256))
func (_LendPool *LendPoolCallerSession) GetNftData(asset common.Address) (DataTypesNftData, error) {
	return _LendPool.Contract.GetNftData(&_LendPool.CallOpts, asset)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_LendPool *LendPoolCaller) GetNftDebtData(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftDebtData", nftAsset, nftTokenId)

	outstruct := new(struct {
		LoanId           *big.Int
		ReserveAsset     common.Address
		TotalCollateral  *big.Int
		TotalDebt        *big.Int
		AvailableBorrows *big.Int
		HealthFactor     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LoanId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveAsset = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalCollateral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrows = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_LendPool *LendPoolSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _LendPool.Contract.GetNftDebtData(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftDebtData is a free data retrieval call binding the contract method 0xec765d3d.
//
// Solidity: function getNftDebtData(address nftAsset, uint256 nftTokenId) view returns(uint256 loanId, address reserveAsset, uint256 totalCollateral, uint256 totalDebt, uint256 availableBorrows, uint256 healthFactor)
func (_LendPool *LendPoolCallerSession) GetNftDebtData(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LoanId           *big.Int
	ReserveAsset     common.Address
	TotalCollateral  *big.Int
	TotalDebt        *big.Int
	AvailableBorrows *big.Int
	HealthFactor     *big.Int
}, error) {
	return _LendPool.Contract.GetNftDebtData(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_LendPool *LendPoolCaller) GetNftLiquidatePrice(opts *bind.CallOpts, nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftLiquidatePrice", nftAsset, nftTokenId)

	outstruct := new(struct {
		LiquidatePrice *big.Int
		PaybackAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LiquidatePrice = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaybackAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_LendPool *LendPoolSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _LendPool.Contract.GetNftLiquidatePrice(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftLiquidatePrice is a free data retrieval call binding the contract method 0x798b9e3d.
//
// Solidity: function getNftLiquidatePrice(address nftAsset, uint256 nftTokenId) view returns(uint256 liquidatePrice, uint256 paybackAmount)
func (_LendPool *LendPoolCallerSession) GetNftLiquidatePrice(nftAsset common.Address, nftTokenId *big.Int) (struct {
	LiquidatePrice *big.Int
	PaybackAmount  *big.Int
}, error) {
	return _LendPool.Contract.GetNftLiquidatePrice(&_LendPool.CallOpts, nftAsset, nftTokenId)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_LendPool *LendPoolCaller) GetNftsList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getNftsList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_LendPool *LendPoolSession) GetNftsList() ([]common.Address, error) {
	return _LendPool.Contract.GetNftsList(&_LendPool.CallOpts)
}

// GetNftsList is a free data retrieval call binding the contract method 0x6b25c835.
//
// Solidity: function getNftsList() view returns(address[])
func (_LendPool *LendPoolCallerSession) GetNftsList() ([]common.Address, error) {
	return _LendPool.Contract.GetNftsList(&_LendPool.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_LendPool *LendPoolCaller) GetPausedTime(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getPausedTime")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_LendPool *LendPoolSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _LendPool.Contract.GetPausedTime(&_LendPool.CallOpts)
}

// GetPausedTime is a free data retrieval call binding the contract method 0x8fc42188.
//
// Solidity: function getPausedTime() view returns(uint256, uint256)
func (_LendPool *LendPoolCallerSession) GetPausedTime() (*big.Int, *big.Int, error) {
	return _LendPool.Contract.GetPausedTime(&_LendPool.CallOpts)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolCaller) GetReserveConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getReserveConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _LendPool.Contract.GetReserveConfiguration(&_LendPool.CallOpts, asset)
}

// GetReserveConfiguration is a free data retrieval call binding the contract method 0x5fc526ff.
//
// Solidity: function getReserveConfiguration(address asset) view returns((uint256))
func (_LendPool *LendPoolCallerSession) GetReserveConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _LendPool.Contract.GetReserveConfiguration(&_LendPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_LendPool *LendPoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_LendPool *LendPoolSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _LendPool.Contract.GetReserveData(&_LendPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint40,address,address,address,uint8))
func (_LendPool *LendPoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _LendPool.Contract.GetReserveData(&_LendPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_LendPool *LendPoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_LendPool *LendPoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _LendPool.Contract.GetReserveNormalizedIncome(&_LendPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_LendPool *LendPoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _LendPool.Contract.GetReserveNormalizedIncome(&_LendPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_LendPool *LendPoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_LendPool *LendPoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _LendPool.Contract.GetReserveNormalizedVariableDebt(&_LendPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_LendPool *LendPoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _LendPool.Contract.GetReserveNormalizedVariableDebt(&_LendPool.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_LendPool *LendPoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_LendPool *LendPoolSession) GetReservesList() ([]common.Address, error) {
	return _LendPool.Contract.GetReservesList(&_LendPool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_LendPool *LendPoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _LendPool.Contract.GetReservesList(&_LendPool.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_LendPool *LendPoolCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_LendPool *LendPoolSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _LendPool.Contract.OnERC721Received(&_LendPool.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) pure returns(bytes4)
func (_LendPool *LendPoolCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _LendPool.Contract.OnERC721Received(&_LendPool.CallOpts, operator, from, tokenId, data)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendPool *LendPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LendPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendPool *LendPoolSession) Paused() (bool, error) {
	return _LendPool.Contract.Paused(&_LendPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LendPool *LendPoolCallerSession) Paused() (bool, error) {
	return _LendPool.Contract.Paused(&_LendPool.CallOpts)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_LendPool *LendPoolTransactor) Auction(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "auction", nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_LendPool *LendPoolSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Auction(&_LendPool.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// Auction is a paid mutator transaction binding the contract method 0xa4c0166b.
//
// Solidity: function auction(address nftAsset, uint256 nftTokenId, uint256 bidPrice, address onBehalfOf) returns()
func (_LendPool *LendPoolTransactorSession) Auction(nftAsset common.Address, nftTokenId *big.Int, bidPrice *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Auction(&_LendPool.TransactOpts, nftAsset, nftTokenId, bidPrice, onBehalfOf)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactor) BatchBorrow(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "batchBorrow", assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.BatchBorrow(&_LendPool.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchBorrow is a paid mutator transaction binding the contract method 0xc9fef2fe.
//
// Solidity: function batchBorrow(address[] assets, uint256[] amounts, address[] nftAssets, uint256[] nftTokenIds, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactorSession) BatchBorrow(assets []common.Address, amounts []*big.Int, nftAssets []common.Address, nftTokenIds []*big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.BatchBorrow(&_LendPool.TransactOpts, assets, amounts, nftAssets, nftTokenIds, onBehalfOf, referralCode)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_LendPool *LendPoolTransactor) BatchRepay(opts *bind.TransactOpts, nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "batchRepay", nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_LendPool *LendPoolSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.BatchRepay(&_LendPool.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// BatchRepay is a paid mutator transaction binding the contract method 0x5bc5bbf1.
//
// Solidity: function batchRepay(address[] nftAssets, uint256[] nftTokenIds, uint256[] amounts) returns(uint256[], bool[])
func (_LendPool *LendPoolTransactorSession) BatchRepay(nftAssets []common.Address, nftTokenIds []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.BatchRepay(&_LendPool.TransactOpts, nftAssets, nftTokenIds, amounts)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "borrow", asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.Borrow(&_LendPool.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xb6529aee.
//
// Solidity: function borrow(address asset, uint256 amount, address nftAsset, uint256 nftTokenId, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactorSession) Borrow(asset common.Address, amount *big.Int, nftAsset common.Address, nftTokenId *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.Borrow(&_LendPool.TransactOpts, asset, amount, nftAsset, nftTokenId, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.Deposit(&_LendPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_LendPool *LendPoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _LendPool.Contract.Deposit(&_LendPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_LendPool *LendPoolTransactor) InitNft(opts *bind.TransactOpts, asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "initNft", asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_LendPool *LendPoolSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.InitNft(&_LendPool.TransactOpts, asset, bNftAddress)
}

// InitNft is a paid mutator transaction binding the contract method 0x873e4dab.
//
// Solidity: function initNft(address asset, address bNftAddress) returns()
func (_LendPool *LendPoolTransactorSession) InitNft(asset common.Address, bNftAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.InitNft(&_LendPool.TransactOpts, asset, bNftAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_LendPool *LendPoolTransactor) InitReserve(opts *bind.TransactOpts, asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "initReserve", asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_LendPool *LendPoolSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.InitReserve(&_LendPool.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x8bd25677.
//
// Solidity: function initReserve(address asset, address bTokenAddress, address debtTokenAddress, address interestRateAddress) returns()
func (_LendPool *LendPoolTransactorSession) InitReserve(asset common.Address, bTokenAddress common.Address, debtTokenAddress common.Address, interestRateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.InitReserve(&_LendPool.TransactOpts, asset, bTokenAddress, debtTokenAddress, interestRateAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_LendPool *LendPoolTransactor) Initialize(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "initialize", provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_LendPool *LendPoolSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Initialize(&_LendPool.TransactOpts, provider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address provider) returns()
func (_LendPool *LendPoolTransactorSession) Initialize(provider common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Initialize(&_LendPool.TransactOpts, provider)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_LendPool *LendPoolTransactor) Liquidate(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "liquidate", nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_LendPool *LendPoolSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Liquidate(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x0710285c.
//
// Solidity: function liquidate(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256)
func (_LendPool *LendPoolTransactorSession) Liquidate(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Liquidate(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_LendPool *LendPoolTransactor) Redeem(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "redeem", nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_LendPool *LendPoolSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Redeem(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Redeem is a paid mutator transaction binding the contract method 0xea2092f3.
//
// Solidity: function redeem(address nftAsset, uint256 nftTokenId, uint256 amount, uint256 bidFine) returns(uint256)
func (_LendPool *LendPoolTransactorSession) Redeem(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int, bidFine *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Redeem(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount, bidFine)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_LendPool *LendPoolTransactor) Repay(opts *bind.TransactOpts, nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "repay", nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_LendPool *LendPoolSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Repay(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount)
}

// Repay is a paid mutator transaction binding the contract method 0x8cd2e0c7.
//
// Solidity: function repay(address nftAsset, uint256 nftTokenId, uint256 amount) returns(uint256, bool)
func (_LendPool *LendPoolTransactorSession) Repay(nftAsset common.Address, nftTokenId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.Repay(&_LendPool.TransactOpts, nftAsset, nftTokenId, amount)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_LendPool *LendPoolTransactor) SetMaxNumberOfNfts(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setMaxNumberOfNfts", val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_LendPool *LendPoolSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetMaxNumberOfNfts(&_LendPool.TransactOpts, val)
}

// SetMaxNumberOfNfts is a paid mutator transaction binding the contract method 0xfdff6f26.
//
// Solidity: function setMaxNumberOfNfts(uint256 val) returns()
func (_LendPool *LendPoolTransactorSession) SetMaxNumberOfNfts(val *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetMaxNumberOfNfts(&_LendPool.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_LendPool *LendPoolTransactor) SetMaxNumberOfReserves(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setMaxNumberOfReserves", val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_LendPool *LendPoolSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetMaxNumberOfReserves(&_LendPool.TransactOpts, val)
}

// SetMaxNumberOfReserves is a paid mutator transaction binding the contract method 0x746c35a2.
//
// Solidity: function setMaxNumberOfReserves(uint256 val) returns()
func (_LendPool *LendPoolTransactorSession) SetMaxNumberOfReserves(val *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetMaxNumberOfReserves(&_LendPool.TransactOpts, val)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolTransactor) SetNftConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setNftConfiguration", asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetNftConfiguration(&_LendPool.TransactOpts, asset, configuration)
}

// SetNftConfiguration is a paid mutator transaction binding the contract method 0x83c8afd7.
//
// Solidity: function setNftConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolTransactorSession) SetNftConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetNftConfiguration(&_LendPool.TransactOpts, asset, configuration)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_LendPool *LendPoolTransactor) SetNftMaxSupplyAndTokenId(opts *bind.TransactOpts, asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setNftMaxSupplyAndTokenId", asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_LendPool *LendPoolSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetNftMaxSupplyAndTokenId(&_LendPool.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetNftMaxSupplyAndTokenId is a paid mutator transaction binding the contract method 0xdb78f216.
//
// Solidity: function setNftMaxSupplyAndTokenId(address asset, uint256 maxSupply, uint256 maxTokenId) returns()
func (_LendPool *LendPoolTransactorSession) SetNftMaxSupplyAndTokenId(asset common.Address, maxSupply *big.Int, maxTokenId *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetNftMaxSupplyAndTokenId(&_LendPool.TransactOpts, asset, maxSupply, maxTokenId)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_LendPool *LendPoolTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_LendPool *LendPoolSession) SetPause(val bool) (*types.Transaction, error) {
	return _LendPool.Contract.SetPause(&_LendPool.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_LendPool *LendPoolTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _LendPool.Contract.SetPause(&_LendPool.TransactOpts, val)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_LendPool *LendPoolTransactor) SetPausedTime(opts *bind.TransactOpts, startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setPausedTime", startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_LendPool *LendPoolSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetPausedTime(&_LendPool.TransactOpts, startTime, durationTime)
}

// SetPausedTime is a paid mutator transaction binding the contract method 0x2f923ff7.
//
// Solidity: function setPausedTime(uint256 startTime, uint256 durationTime) returns()
func (_LendPool *LendPoolTransactorSession) SetPausedTime(startTime *big.Int, durationTime *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetPausedTime(&_LendPool.TransactOpts, startTime, durationTime)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolTransactor) SetReserveConfiguration(opts *bind.TransactOpts, asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setReserveConfiguration", asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetReserveConfiguration(&_LendPool.TransactOpts, asset, configuration)
}

// SetReserveConfiguration is a paid mutator transaction binding the contract method 0x43f0f733.
//
// Solidity: function setReserveConfiguration(address asset, uint256 configuration) returns()
func (_LendPool *LendPoolTransactorSession) SetReserveConfiguration(asset common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _LendPool.Contract.SetReserveConfiguration(&_LendPool.TransactOpts, asset, configuration)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_LendPool *LendPoolTransactor) SetReserveInterestRateAddress(opts *bind.TransactOpts, asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "setReserveInterestRateAddress", asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_LendPool *LendPoolSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.SetReserveInterestRateAddress(&_LendPool.TransactOpts, asset, rateAddress)
}

// SetReserveInterestRateAddress is a paid mutator transaction binding the contract method 0x83b1555f.
//
// Solidity: function setReserveInterestRateAddress(address asset, address rateAddress) returns()
func (_LendPool *LendPoolTransactorSession) SetReserveInterestRateAddress(asset common.Address, rateAddress common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.SetReserveInterestRateAddress(&_LendPool.TransactOpts, asset, rateAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_LendPool *LendPoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LendPool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_LendPool *LendPoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Withdraw(&_LendPool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_LendPool *LendPoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _LendPool.Contract.Withdraw(&_LendPool.TransactOpts, asset, amount, to)
}

// LendPoolAuctionIterator is returned from FilterAuction and is used to iterate over the raw logs and unpacked data for Auction events raised by the LendPool contract.
type LendPoolAuctionIterator struct {
	Event *LendPoolAuction // Event containing the contract specifics and raw log

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
func (it *LendPoolAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolAuction)
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
		it.Event = new(LendPoolAuction)
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
func (it *LendPoolAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolAuction represents a Auction event raised by the LendPool contract.
type LendPoolAuction struct {
	User       common.Address
	Reserve    common.Address
	BidPrice   *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuction is a free log retrieval operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) FilterAuction(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*LendPoolAuctionIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolAuctionIterator{contract: _LendPool.contract, event: "Auction", logs: logs, sub: sub}, nil
}

// WatchAuction is a free log subscription operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) WatchAuction(opts *bind.WatchOpts, sink chan<- *LendPoolAuction, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Auction", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolAuction)
				if err := _LendPool.contract.UnpackLog(event, "Auction", log); err != nil {
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

// ParseAuction is a log parse operation binding the contract event 0xd4c7449fa0ea241233dd0a9e78a940879918f95e0caa34e0399a7d2813c8efba.
//
// Solidity: event Auction(address user, address indexed reserve, uint256 bidPrice, address indexed nftAsset, uint256 nftTokenId, address onBehalfOf, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) ParseAuction(log types.Log) (*LendPoolAuction, error) {
	event := new(LendPoolAuction)
	if err := _LendPool.contract.UnpackLog(event, "Auction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LendPool contract.
type LendPoolBorrowIterator struct {
	Event *LendPoolBorrow // Event containing the contract specifics and raw log

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
func (it *LendPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolBorrow)
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
		it.Event = new(LendPoolBorrow)
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
func (it *LendPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolBorrow represents a Borrow event raised by the LendPool contract.
type LendPoolBorrow struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	OnBehalfOf common.Address
	BorrowRate *big.Int
	LoanId     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*LendPoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolBorrowIterator{contract: _LendPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LendPoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolBorrow)
				if err := _LendPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xcfb3a669117d9dc90f0d3521228dc9fe67c5102dde205ef16fe9a1f81be698d5.
//
// Solidity: event Borrow(address user, address indexed reserve, uint256 amount, address nftAsset, uint256 nftTokenId, address indexed onBehalfOf, uint256 borrowRate, uint256 loanId, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) ParseBorrow(log types.Log) (*LendPoolBorrow, error) {
	event := new(LendPoolBorrow)
	if err := _LendPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the LendPool contract.
type LendPoolDepositIterator struct {
	Event *LendPoolDeposit // Event containing the contract specifics and raw log

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
func (it *LendPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolDeposit)
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
		it.Event = new(LendPoolDeposit)
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
func (it *LendPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolDeposit represents a Deposit event raised by the LendPool contract.
type LendPoolDeposit struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	OnBehalfOf common.Address
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*LendPoolDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolDepositIterator{contract: _LendPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LendPoolDeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}
	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolDeposit)
				if err := _LendPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x443ff2d25883a4800d36062db52ca3dd7ced05bd8627c8a6a37f8699715b5431.
//
// Solidity: event Deposit(address user, address indexed reserve, uint256 amount, address indexed onBehalfOf, uint16 indexed referral)
func (_LendPool *LendPoolFilterer) ParseDeposit(log types.Log) (*LendPoolDeposit, error) {
	event := new(LendPoolDeposit)
	if err := _LendPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolLiquidateIterator is returned from FilterLiquidate and is used to iterate over the raw logs and unpacked data for Liquidate events raised by the LendPool contract.
type LendPoolLiquidateIterator struct {
	Event *LendPoolLiquidate // Event containing the contract specifics and raw log

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
func (it *LendPoolLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolLiquidate)
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
		it.Event = new(LendPoolLiquidate)
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
func (it *LendPoolLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolLiquidate represents a Liquidate event raised by the LendPool contract.
type LendPoolLiquidate struct {
	User         common.Address
	Reserve      common.Address
	RepayAmount  *big.Int
	RemainAmount *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLiquidate is a free log retrieval operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) FilterLiquidate(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*LendPoolLiquidateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolLiquidateIterator{contract: _LendPool.contract, event: "Liquidate", logs: logs, sub: sub}, nil
}

// WatchLiquidate is a free log subscription operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) WatchLiquidate(opts *bind.WatchOpts, sink chan<- *LendPoolLiquidate, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Liquidate", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolLiquidate)
				if err := _LendPool.contract.UnpackLog(event, "Liquidate", log); err != nil {
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

// ParseLiquidate is a log parse operation binding the contract event 0xf028795898a18c6fc88094dc5671c6a79d5dc3458c44015e9299fbc6c6268cf8.
//
// Solidity: event Liquidate(address user, address indexed reserve, uint256 repayAmount, uint256 remainAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) ParseLiquidate(log types.Log) (*LendPoolLiquidate, error) {
	event := new(LendPoolLiquidate)
	if err := _LendPool.contract.UnpackLog(event, "Liquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LendPool contract.
type LendPoolPausedIterator struct {
	Event *LendPoolPaused // Event containing the contract specifics and raw log

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
func (it *LendPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolPaused)
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
		it.Event = new(LendPoolPaused)
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
func (it *LendPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolPaused represents a Paused event raised by the LendPool contract.
type LendPoolPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_LendPool *LendPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*LendPoolPausedIterator, error) {

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LendPoolPausedIterator{contract: _LendPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_LendPool *LendPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LendPoolPaused) (event.Subscription, error) {

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolPaused)
				if err := _LendPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_LendPool *LendPoolFilterer) ParsePaused(log types.Log) (*LendPoolPaused, error) {
	event := new(LendPoolPaused)
	if err := _LendPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolPausedTimeUpdatedIterator is returned from FilterPausedTimeUpdated and is used to iterate over the raw logs and unpacked data for PausedTimeUpdated events raised by the LendPool contract.
type LendPoolPausedTimeUpdatedIterator struct {
	Event *LendPoolPausedTimeUpdated // Event containing the contract specifics and raw log

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
func (it *LendPoolPausedTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolPausedTimeUpdated)
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
		it.Event = new(LendPoolPausedTimeUpdated)
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
func (it *LendPoolPausedTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolPausedTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolPausedTimeUpdated represents a PausedTimeUpdated event raised by the LendPool contract.
type LendPoolPausedTimeUpdated struct {
	StartTime    *big.Int
	DurationTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPausedTimeUpdated is a free log retrieval operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_LendPool *LendPoolFilterer) FilterPausedTimeUpdated(opts *bind.FilterOpts) (*LendPoolPausedTimeUpdatedIterator, error) {

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &LendPoolPausedTimeUpdatedIterator{contract: _LendPool.contract, event: "PausedTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchPausedTimeUpdated is a free log subscription operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_LendPool *LendPoolFilterer) WatchPausedTimeUpdated(opts *bind.WatchOpts, sink chan<- *LendPoolPausedTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "PausedTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolPausedTimeUpdated)
				if err := _LendPool.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
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

// ParsePausedTimeUpdated is a log parse operation binding the contract event 0xd897a722b1c0a957941f99a13c0ea24d7d4ffafe0953658f68f49e13ccba5c5a.
//
// Solidity: event PausedTimeUpdated(uint256 startTime, uint256 durationTime)
func (_LendPool *LendPoolFilterer) ParsePausedTimeUpdated(log types.Log) (*LendPoolPausedTimeUpdated, error) {
	event := new(LendPoolPausedTimeUpdated)
	if err := _LendPool.contract.UnpackLog(event, "PausedTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the LendPool contract.
type LendPoolRedeemIterator struct {
	Event *LendPoolRedeem // Event containing the contract specifics and raw log

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
func (it *LendPoolRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolRedeem)
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
		it.Event = new(LendPoolRedeem)
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
func (it *LendPoolRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolRedeem represents a Redeem event raised by the LendPool contract.
type LendPoolRedeem struct {
	User         common.Address
	Reserve      common.Address
	BorrowAmount *big.Int
	FineAmount   *big.Int
	NftAsset     common.Address
	NftTokenId   *big.Int
	Borrower     common.Address
	LoanId       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) FilterRedeem(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*LendPoolRedeemIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolRedeemIterator{contract: _LendPool.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *LendPoolRedeem, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Redeem", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolRedeem)
				if err := _LendPool.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x0fcfe1a3f2afab13e32fa3c091795159ed5dfe66dc078e21c7f521f42e163afc.
//
// Solidity: event Redeem(address user, address indexed reserve, uint256 borrowAmount, uint256 fineAmount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) ParseRedeem(log types.Log) (*LendPoolRedeem, error) {
	event := new(LendPoolRedeem)
	if err := _LendPool.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the LendPool contract.
type LendPoolRepayIterator struct {
	Event *LendPoolRepay // Event containing the contract specifics and raw log

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
func (it *LendPoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolRepay)
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
		it.Event = new(LendPoolRepay)
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
func (it *LendPoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolRepay represents a Repay event raised by the LendPool contract.
type LendPoolRepay struct {
	User       common.Address
	Reserve    common.Address
	Amount     *big.Int
	NftAsset   common.Address
	NftTokenId *big.Int
	Borrower   common.Address
	LoanId     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (*LendPoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolRepayIterator{contract: _LendPool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LendPoolRepay, reserve []common.Address, nftAsset []common.Address, borrower []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var nftAssetRule []interface{}
	for _, nftAssetItem := range nftAsset {
		nftAssetRule = append(nftAssetRule, nftAssetItem)
	}

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Repay", reserveRule, nftAssetRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolRepay)
				if err := _LendPool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x50e03867c1178391f204f7bf0eb2f52d5167dc65a99a9650a95abe55d17be17e.
//
// Solidity: event Repay(address user, address indexed reserve, uint256 amount, address indexed nftAsset, uint256 nftTokenId, address indexed borrower, uint256 loanId)
func (_LendPool *LendPoolFilterer) ParseRepay(log types.Log) (*LendPoolRepay, error) {
	event := new(LendPoolRepay)
	if err := _LendPool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the LendPool contract.
type LendPoolReserveDataUpdatedIterator struct {
	Event *LendPoolReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *LendPoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolReserveDataUpdated)
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
		it.Event = new(LendPoolReserveDataUpdated)
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
func (it *LendPoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolReserveDataUpdated represents a ReserveDataUpdated event raised by the LendPool contract.
type LendPoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendPool *LendPoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*LendPoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolReserveDataUpdatedIterator{contract: _LendPool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendPool *LendPoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *LendPoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolReserveDataUpdated)
				if err := _LendPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x4063a2df84b66bb796eb32622851d833e57b2c4292900c18f963af8808b13e35.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_LendPool *LendPoolFilterer) ParseReserveDataUpdated(log types.Log) (*LendPoolReserveDataUpdated, error) {
	event := new(LendPoolReserveDataUpdated)
	if err := _LendPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LendPool contract.
type LendPoolUnpausedIterator struct {
	Event *LendPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *LendPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolUnpaused)
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
		it.Event = new(LendPoolUnpaused)
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
func (it *LendPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolUnpaused represents a Unpaused event raised by the LendPool contract.
type LendPoolUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_LendPool *LendPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LendPoolUnpausedIterator, error) {

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LendPoolUnpausedIterator{contract: _LendPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_LendPool *LendPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LendPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolUnpaused)
				if err := _LendPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_LendPool *LendPoolFilterer) ParseUnpaused(log types.Log) (*LendPoolUnpaused, error) {
	event := new(LendPoolUnpaused)
	if err := _LendPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LendPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the LendPool contract.
type LendPoolWithdrawIterator struct {
	Event *LendPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *LendPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendPoolWithdraw)
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
		it.Event = new(LendPoolWithdraw)
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
func (it *LendPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendPoolWithdraw represents a Withdraw event raised by the LendPool contract.
type LendPoolWithdraw struct {
	User    common.Address
	Reserve common.Address
	Amount  *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_LendPool *LendPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, reserve []common.Address, to []common.Address) (*LendPoolWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LendPool.contract.FilterLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LendPoolWithdrawIterator{contract: _LendPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_LendPool *LendPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *LendPoolWithdraw, user []common.Address, reserve []common.Address, to []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LendPool.contract.WatchLogs(opts, "Withdraw", userRule, reserveRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendPoolWithdraw)
				if err := _LendPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3ed4ee04a905a278b050a856bbe7ddaaf327a30514373e65aa6103beeae488c3.
//
// Solidity: event Withdraw(address indexed user, address indexed reserve, uint256 amount, address indexed to)
func (_LendPool *LendPoolFilterer) ParseWithdraw(log types.Log) (*LendPoolWithdraw, error) {
	event := new(LendPoolWithdraw)
	if err := _LendPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
