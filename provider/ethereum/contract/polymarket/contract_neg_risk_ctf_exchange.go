// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package polymarket

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

// NegRiskCTFExchangeMetaData contains all meta data concerning the NegRiskCTFExchange contract.
var NegRiskCTFExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_ctf\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_negRiskAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proxyFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_safeFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidComplement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingGtRemaining\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MismatchedTokenIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotCrossing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTaker\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderFilledOrCancelled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Paused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleTokensReceived\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdminAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOperatorAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NewOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"takerOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"takerOrderMaker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAssetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makerAmountFilled\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"takerAmountFilled\",\"type\":\"uint256\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProxyFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProxyFactory\",\"type\":\"address\"}],\"name\":\"ProxyFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedOperator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"RemovedOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldSafeFactory\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSafeFactory\",\"type\":\"address\"}],\"name\":\"SafeFactoryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token0\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"token1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"TokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pauser\",\"type\":\"address\"}],\"name\":\"TradingUnpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"fillAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"fillOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCollateral\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getComplement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"}],\"name\":\"getConditionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCtf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getOrderStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"internalType\":\"structOrderStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPolyProxyFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getPolyProxyWalletAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getSafeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSafeFactoryImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"usr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"isValidNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"takerOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"makerOrders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"takerFillAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"makerFillAmounts\",\"type\":\"uint256[]\"}],\"name\":\"matchOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFilledOrCancelled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"parentCollectionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"name\":\"registerToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"conditionId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceAdminRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperatorRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"safeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newProxyFactory\",\"type\":\"address\"}],\"name\":\"setProxyFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSafeFactory\",\"type\":\"address\"}],\"name\":\"setSafeFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"token\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"complement\",\"type\":\"uint256\"}],\"name\":\"validateComplement\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRateBps\",\"type\":\"uint256\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"enumSignatureType\",\"name\":\"signatureType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"validateOrderSignature\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"validateTokenId\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NegRiskCTFExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use NegRiskCTFExchangeMetaData.ABI instead.
var NegRiskCTFExchangeABI = NegRiskCTFExchangeMetaData.ABI

// NegRiskCTFExchange is an auto generated Go binding around an Ethereum contract.
type NegRiskCTFExchange struct {
	NegRiskCTFExchangeCaller     // Read-only binding to the contract
	NegRiskCTFExchangeTransactor // Write-only binding to the contract
	NegRiskCTFExchangeFilterer   // Log filterer for contract events
}

// NegRiskCTFExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NegRiskCTFExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCTFExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NegRiskCTFExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCTFExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NegRiskCTFExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NegRiskCTFExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NegRiskCTFExchangeSession struct {
	Contract     *NegRiskCTFExchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NegRiskCTFExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NegRiskCTFExchangeCallerSession struct {
	Contract *NegRiskCTFExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NegRiskCTFExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NegRiskCTFExchangeTransactorSession struct {
	Contract     *NegRiskCTFExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NegRiskCTFExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NegRiskCTFExchangeRaw struct {
	Contract *NegRiskCTFExchange // Generic contract binding to access the raw methods on
}

// NegRiskCTFExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NegRiskCTFExchangeCallerRaw struct {
	Contract *NegRiskCTFExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// NegRiskCTFExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NegRiskCTFExchangeTransactorRaw struct {
	Contract *NegRiskCTFExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNegRiskCTFExchange creates a new instance of NegRiskCTFExchange, bound to a specific deployed contract.
func NewNegRiskCTFExchange(address common.Address, backend bind.ContractBackend) (*NegRiskCTFExchange, error) {
	contract, err := bindNegRiskCTFExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchange{NegRiskCTFExchangeCaller: NegRiskCTFExchangeCaller{contract: contract}, NegRiskCTFExchangeTransactor: NegRiskCTFExchangeTransactor{contract: contract}, NegRiskCTFExchangeFilterer: NegRiskCTFExchangeFilterer{contract: contract}}, nil
}

// NewNegRiskCTFExchangeCaller creates a new read-only instance of NegRiskCTFExchange, bound to a specific deployed contract.
func NewNegRiskCTFExchangeCaller(address common.Address, caller bind.ContractCaller) (*NegRiskCTFExchangeCaller, error) {
	contract, err := bindNegRiskCTFExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeCaller{contract: contract}, nil
}

// NewNegRiskCTFExchangeTransactor creates a new write-only instance of NegRiskCTFExchange, bound to a specific deployed contract.
func NewNegRiskCTFExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*NegRiskCTFExchangeTransactor, error) {
	contract, err := bindNegRiskCTFExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeTransactor{contract: contract}, nil
}

// NewNegRiskCTFExchangeFilterer creates a new log filterer instance of NegRiskCTFExchange, bound to a specific deployed contract.
func NewNegRiskCTFExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*NegRiskCTFExchangeFilterer, error) {
	contract, err := bindNegRiskCTFExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeFilterer{contract: contract}, nil
}

// bindNegRiskCTFExchange binds a generic wrapper to an already deployed contract.
func bindNegRiskCTFExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NegRiskCTFExchangeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCTFExchange *NegRiskCTFExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCTFExchange.Contract.NegRiskCTFExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCTFExchange *NegRiskCTFExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.NegRiskCTFExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCTFExchange *NegRiskCTFExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.NegRiskCTFExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NegRiskCTFExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Admins(&_NegRiskCTFExchange.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) Admins(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Admins(&_NegRiskCTFExchange.CallOpts, arg0)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) DomainSeparator() ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.DomainSeparator(&_NegRiskCTFExchange.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) DomainSeparator() ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.DomainSeparator(&_NegRiskCTFExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetCollateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getCollateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetCollateral() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetCollateral(&_NegRiskCTFExchange.CallOpts)
}

// GetCollateral is a free data retrieval call binding the contract method 0x5c1548fb.
//
// Solidity: function getCollateral() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetCollateral() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetCollateral(&_NegRiskCTFExchange.CallOpts)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetComplement(opts *bind.CallOpts, token *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getComplement", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.GetComplement(&_NegRiskCTFExchange.CallOpts, token)
}

// GetComplement is a free data retrieval call binding the contract method 0xa10f3dce.
//
// Solidity: function getComplement(uint256 token) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetComplement(token *big.Int) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.GetComplement(&_NegRiskCTFExchange.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetConditionId(opts *bind.CallOpts, token *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getConditionId", token)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.GetConditionId(&_NegRiskCTFExchange.CallOpts, token)
}

// GetConditionId is a free data retrieval call binding the contract method 0xd7fb272f.
//
// Solidity: function getConditionId(uint256 token) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetConditionId(token *big.Int) ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.GetConditionId(&_NegRiskCTFExchange.CallOpts, token)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetCtf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getCtf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetCtf() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetCtf(&_NegRiskCTFExchange.CallOpts)
}

// GetCtf is a free data retrieval call binding the contract method 0x3b521d78.
//
// Solidity: function getCtf() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetCtf() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetCtf(&_NegRiskCTFExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetMaxFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getMaxFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.GetMaxFeeRate(&_NegRiskCTFExchange.CallOpts)
}

// GetMaxFeeRate is a free data retrieval call binding the contract method 0x4a2a11f5.
//
// Solidity: function getMaxFeeRate() pure returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetMaxFeeRate() (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.GetMaxFeeRate(&_NegRiskCTFExchange.CallOpts)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetOrderStatus(opts *bind.CallOpts, orderHash [32]byte) (OrderStatus, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getOrderStatus", orderHash)

	if err != nil {
		return *new(OrderStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(OrderStatus)).(*OrderStatus)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskCTFExchange.Contract.GetOrderStatus(&_NegRiskCTFExchange.CallOpts, orderHash)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderHash) view returns((bool,uint256))
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetOrderStatus(orderHash [32]byte) (OrderStatus, error) {
	return _NegRiskCTFExchange.Contract.GetOrderStatus(&_NegRiskCTFExchange.CallOpts, orderHash)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetPolyProxyFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getPolyProxyFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetPolyProxyFactoryImplementation(&_NegRiskCTFExchange.CallOpts)
}

// GetPolyProxyFactoryImplementation is a free data retrieval call binding the contract method 0x06b9d691.
//
// Solidity: function getPolyProxyFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetPolyProxyFactoryImplementation() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetPolyProxyFactoryImplementation(&_NegRiskCTFExchange.CallOpts)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetPolyProxyWalletAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getPolyProxyWalletAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetPolyProxyWalletAddress(&_NegRiskCTFExchange.CallOpts, _addr)
}

// GetPolyProxyWalletAddress is a free data retrieval call binding the contract method 0xedef7d8e.
//
// Solidity: function getPolyProxyWalletAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetPolyProxyWalletAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetPolyProxyWalletAddress(&_NegRiskCTFExchange.CallOpts, _addr)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetProxyFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetProxyFactory(&_NegRiskCTFExchange.CallOpts)
}

// GetProxyFactory is a free data retrieval call binding the contract method 0xb28c51c0.
//
// Solidity: function getProxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetProxyFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetProxyFactory(&_NegRiskCTFExchange.CallOpts)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetSafeAddress(opts *bind.CallOpts, _addr common.Address) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getSafeAddress", _addr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeAddress(&_NegRiskCTFExchange.CallOpts, _addr)
}

// GetSafeAddress is a free data retrieval call binding the contract method 0xa287bdf1.
//
// Solidity: function getSafeAddress(address _addr) view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetSafeAddress(_addr common.Address) (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeAddress(&_NegRiskCTFExchange.CallOpts, _addr)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetSafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getSafeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetSafeFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeFactory(&_NegRiskCTFExchange.CallOpts)
}

// GetSafeFactory is a free data retrieval call binding the contract method 0x75d7370a.
//
// Solidity: function getSafeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetSafeFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeFactory(&_NegRiskCTFExchange.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) GetSafeFactoryImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "getSafeFactoryImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeFactoryImplementation(&_NegRiskCTFExchange.CallOpts)
}

// GetSafeFactoryImplementation is a free data retrieval call binding the contract method 0xe03ac3d0.
//
// Solidity: function getSafeFactoryImplementation() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) GetSafeFactoryImplementation() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.GetSafeFactoryImplementation(&_NegRiskCTFExchange.CallOpts)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) HashOrder(opts *bind.CallOpts, order Order) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.HashOrder(&_NegRiskCTFExchange.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0xe50e4f97.
//
// Solidity: function hashOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) HashOrder(order Order) ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.HashOrder(&_NegRiskCTFExchange.CallOpts, order)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) IsAdmin(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "isAdmin", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) IsAdmin(usr common.Address) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsAdmin(&_NegRiskCTFExchange.CallOpts, usr)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) IsAdmin(usr common.Address) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsAdmin(&_NegRiskCTFExchange.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) IsOperator(opts *bind.CallOpts, usr common.Address) (bool, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "isOperator", usr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) IsOperator(usr common.Address) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsOperator(&_NegRiskCTFExchange.CallOpts, usr)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address usr) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) IsOperator(usr common.Address) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsOperator(&_NegRiskCTFExchange.CallOpts, usr)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) IsValidNonce(opts *bind.CallOpts, usr common.Address, nonce *big.Int) (bool, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "isValidNonce", usr, nonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsValidNonce(&_NegRiskCTFExchange.CallOpts, usr, nonce)
}

// IsValidNonce is a free data retrieval call binding the contract method 0x0647ee20.
//
// Solidity: function isValidNonce(address usr, uint256 nonce) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) IsValidNonce(usr common.Address, nonce *big.Int) (bool, error) {
	return _NegRiskCTFExchange.Contract.IsValidNonce(&_NegRiskCTFExchange.CallOpts, usr, nonce)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Nonces(&_NegRiskCTFExchange.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Nonces(&_NegRiskCTFExchange.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) Operators(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "operators", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Operators(&_NegRiskCTFExchange.CallOpts, arg0)
}

// Operators is a free data retrieval call binding the contract method 0x13e7c9d8.
//
// Solidity: function operators(address ) view returns(uint256)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) Operators(arg0 common.Address) (*big.Int, error) {
	return _NegRiskCTFExchange.Contract.Operators(&_NegRiskCTFExchange.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) OrderStatus(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "orderStatus", arg0)

	outstruct := new(struct {
		IsFilledOrCancelled bool
		Remaining           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFilledOrCancelled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Remaining = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _NegRiskCTFExchange.Contract.OrderStatus(&_NegRiskCTFExchange.CallOpts, arg0)
}

// OrderStatus is a free data retrieval call binding the contract method 0x2dff692d.
//
// Solidity: function orderStatus(bytes32 ) view returns(bool isFilledOrCancelled, uint256 remaining)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) OrderStatus(arg0 [32]byte) (struct {
	IsFilledOrCancelled bool
	Remaining           *big.Int
}, error) {
	return _NegRiskCTFExchange.Contract.OrderStatus(&_NegRiskCTFExchange.CallOpts, arg0)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ParentCollectionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "parentCollectionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ParentCollectionId() ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.ParentCollectionId(&_NegRiskCTFExchange.CallOpts)
}

// ParentCollectionId is a free data retrieval call binding the contract method 0x44bea37e.
//
// Solidity: function parentCollectionId() view returns(bytes32)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ParentCollectionId() ([32]byte, error) {
	return _NegRiskCTFExchange.Contract.ParentCollectionId(&_NegRiskCTFExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) Paused() (bool, error) {
	return _NegRiskCTFExchange.Contract.Paused(&_NegRiskCTFExchange.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) Paused() (bool, error) {
	return _NegRiskCTFExchange.Contract.Paused(&_NegRiskCTFExchange.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "proxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ProxyFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.ProxyFactory(&_NegRiskCTFExchange.CallOpts)
}

// ProxyFactory is a free data retrieval call binding the contract method 0xc10f1a75.
//
// Solidity: function proxyFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ProxyFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.ProxyFactory(&_NegRiskCTFExchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) Registry(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "registry", arg0)

	outstruct := new(struct {
		Complement  *big.Int
		ConditionId [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Complement = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ConditionId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _NegRiskCTFExchange.Contract.Registry(&_NegRiskCTFExchange.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x5893253c.
//
// Solidity: function registry(uint256 ) view returns(uint256 complement, bytes32 conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) Registry(arg0 *big.Int) (struct {
	Complement  *big.Int
	ConditionId [32]byte
}, error) {
	return _NegRiskCTFExchange.Contract.Registry(&_NegRiskCTFExchange.CallOpts, arg0)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) SafeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "safeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) SafeFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.SafeFactory(&_NegRiskCTFExchange.CallOpts)
}

// SafeFactory is a free data retrieval call binding the contract method 0x131e7e1c.
//
// Solidity: function safeFactory() view returns(address)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) SafeFactory() (common.Address, error) {
	return _NegRiskCTFExchange.Contract.SafeFactory(&_NegRiskCTFExchange.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCTFExchange.Contract.SupportsInterface(&_NegRiskCTFExchange.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NegRiskCTFExchange.Contract.SupportsInterface(&_NegRiskCTFExchange.CallOpts, interfaceId)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ValidateComplement(opts *bind.CallOpts, token *big.Int, complement *big.Int) error {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "validateComplement", token, complement)

	if err != nil {
		return err
	}

	return err

}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _NegRiskCTFExchange.Contract.ValidateComplement(&_NegRiskCTFExchange.CallOpts, token, complement)
}

// ValidateComplement is a free data retrieval call binding the contract method 0xd82da838.
//
// Solidity: function validateComplement(uint256 token, uint256 complement) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ValidateComplement(token *big.Int, complement *big.Int) error {
	return _NegRiskCTFExchange.Contract.ValidateComplement(&_NegRiskCTFExchange.CallOpts, token, complement)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ValidateOrder(opts *bind.CallOpts, order Order) error {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "validateOrder", order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ValidateOrder(order Order) error {
	return _NegRiskCTFExchange.Contract.ValidateOrder(&_NegRiskCTFExchange.CallOpts, order)
}

// ValidateOrder is a free data retrieval call binding the contract method 0x654f0ce4.
//
// Solidity: function validateOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ValidateOrder(order Order) error {
	return _NegRiskCTFExchange.Contract.ValidateOrder(&_NegRiskCTFExchange.CallOpts, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ValidateOrderSignature(opts *bind.CallOpts, orderHash [32]byte, order Order) error {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "validateOrderSignature", orderHash, order)

	if err != nil {
		return err
	}

	return err

}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskCTFExchange.Contract.ValidateOrderSignature(&_NegRiskCTFExchange.CallOpts, orderHash, order)
}

// ValidateOrderSignature is a free data retrieval call binding the contract method 0xe2eec405.
//
// Solidity: function validateOrderSignature(bytes32 orderHash, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ValidateOrderSignature(orderHash [32]byte, order Order) error {
	return _NegRiskCTFExchange.Contract.ValidateOrderSignature(&_NegRiskCTFExchange.CallOpts, orderHash, order)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCaller) ValidateTokenId(opts *bind.CallOpts, tokenId *big.Int) error {
	var out []interface{}
	err := _NegRiskCTFExchange.contract.Call(opts, &out, "validateTokenId", tokenId)

	if err != nil {
		return err
	}

	return err

}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) ValidateTokenId(tokenId *big.Int) error {
	return _NegRiskCTFExchange.Contract.ValidateTokenId(&_NegRiskCTFExchange.CallOpts, tokenId)
}

// ValidateTokenId is a free data retrieval call binding the contract method 0x34600901.
//
// Solidity: function validateTokenId(uint256 tokenId) view returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeCallerSession) ValidateTokenId(tokenId *big.Int) error {
	return _NegRiskCTFExchange.Contract.ValidateTokenId(&_NegRiskCTFExchange.CallOpts, tokenId)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) AddAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "addAdmin", admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.AddAdmin(&_NegRiskCTFExchange.TransactOpts, admin_)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) AddAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.AddAdmin(&_NegRiskCTFExchange.TransactOpts, admin_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) AddOperator(opts *bind.TransactOpts, operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "addOperator", operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.AddOperator(&_NegRiskCTFExchange.TransactOpts, operator_)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address operator_) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) AddOperator(operator_ common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.AddOperator(&_NegRiskCTFExchange.TransactOpts, operator_)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.CancelOrder(&_NegRiskCTFExchange.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xa6dfcf86.
//
// Solidity: function cancelOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.CancelOrder(&_NegRiskCTFExchange.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.CancelOrders(&_NegRiskCTFExchange.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xfa950b48.
//
// Solidity: function cancelOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.CancelOrders(&_NegRiskCTFExchange.TransactOpts, orders)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) FillOrder(opts *bind.TransactOpts, order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "fillOrder", order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.FillOrder(&_NegRiskCTFExchange.TransactOpts, order, fillAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0xfe729aaf.
//
// Solidity: function fillOrder((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) order, uint256 fillAmount) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) FillOrder(order Order, fillAmount *big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.FillOrder(&_NegRiskCTFExchange.TransactOpts, order, fillAmount)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) FillOrders(opts *bind.TransactOpts, orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "fillOrders", orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.FillOrders(&_NegRiskCTFExchange.TransactOpts, orders, fillAmounts)
}

// FillOrders is a paid mutator transaction binding the contract method 0xd798eff6.
//
// Solidity: function fillOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] orders, uint256[] fillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) FillOrders(orders []Order, fillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.FillOrders(&_NegRiskCTFExchange.TransactOpts, orders, fillAmounts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) IncrementNonce() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.IncrementNonce(&_NegRiskCTFExchange.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.IncrementNonce(&_NegRiskCTFExchange.TransactOpts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) MatchOrders(opts *bind.TransactOpts, takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "matchOrders", takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.MatchOrders(&_NegRiskCTFExchange.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// MatchOrders is a paid mutator transaction binding the contract method 0xe60f0c05.
//
// Solidity: function matchOrders((uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes) takerOrder, (uint256,address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint8,uint8,bytes)[] makerOrders, uint256 takerFillAmount, uint256[] makerFillAmounts) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) MatchOrders(takerOrder Order, makerOrders []Order, takerFillAmount *big.Int, makerFillAmounts []*big.Int) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.MatchOrders(&_NegRiskCTFExchange.TransactOpts, takerOrder, makerOrders, takerFillAmount, makerFillAmounts)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.OnERC1155BatchReceived(&_NegRiskCTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.OnERC1155BatchReceived(&_NegRiskCTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.OnERC1155Received(&_NegRiskCTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.OnERC1155Received(&_NegRiskCTFExchange.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) PauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "pauseTrading")
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) PauseTrading() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.PauseTrading(&_NegRiskCTFExchange.TransactOpts)
}

// PauseTrading is a paid mutator transaction binding the contract method 0x1031e36e.
//
// Solidity: function pauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) PauseTrading() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.PauseTrading(&_NegRiskCTFExchange.TransactOpts)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) RegisterToken(opts *bind.TransactOpts, token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "registerToken", token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RegisterToken(&_NegRiskCTFExchange.TransactOpts, token, complement, conditionId)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x68c7450f.
//
// Solidity: function registerToken(uint256 token, uint256 complement, bytes32 conditionId) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) RegisterToken(token *big.Int, complement *big.Int, conditionId [32]byte) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RegisterToken(&_NegRiskCTFExchange.TransactOpts, token, complement, conditionId)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) RemoveAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "removeAdmin", admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RemoveAdmin(&_NegRiskCTFExchange.TransactOpts, admin)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address admin) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) RemoveAdmin(admin common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RemoveAdmin(&_NegRiskCTFExchange.TransactOpts, admin)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RemoveOperator(&_NegRiskCTFExchange.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(address operator) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RemoveOperator(&_NegRiskCTFExchange.TransactOpts, operator)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) RenounceAdminRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "renounceAdminRole")
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) RenounceAdminRole() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RenounceAdminRole(&_NegRiskCTFExchange.TransactOpts)
}

// RenounceAdminRole is a paid mutator transaction binding the contract method 0x83b8a5ae.
//
// Solidity: function renounceAdminRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) RenounceAdminRole() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RenounceAdminRole(&_NegRiskCTFExchange.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) RenounceOperatorRole(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "renounceOperatorRole")
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RenounceOperatorRole(&_NegRiskCTFExchange.TransactOpts)
}

// RenounceOperatorRole is a paid mutator transaction binding the contract method 0x3d6d3598.
//
// Solidity: function renounceOperatorRole() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) RenounceOperatorRole() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.RenounceOperatorRole(&_NegRiskCTFExchange.TransactOpts)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) SetProxyFactory(opts *bind.TransactOpts, _newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "setProxyFactory", _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.SetProxyFactory(&_NegRiskCTFExchange.TransactOpts, _newProxyFactory)
}

// SetProxyFactory is a paid mutator transaction binding the contract method 0xfbddd751.
//
// Solidity: function setProxyFactory(address _newProxyFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) SetProxyFactory(_newProxyFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.SetProxyFactory(&_NegRiskCTFExchange.TransactOpts, _newProxyFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) SetSafeFactory(opts *bind.TransactOpts, _newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "setSafeFactory", _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.SetSafeFactory(&_NegRiskCTFExchange.TransactOpts, _newSafeFactory)
}

// SetSafeFactory is a paid mutator transaction binding the contract method 0x4544f055.
//
// Solidity: function setSafeFactory(address _newSafeFactory) returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) SetSafeFactory(_newSafeFactory common.Address) (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.SetSafeFactory(&_NegRiskCTFExchange.TransactOpts, _newSafeFactory)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactor) UnpauseTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NegRiskCTFExchange.contract.Transact(opts, "unpauseTrading")
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeSession) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.UnpauseTrading(&_NegRiskCTFExchange.TransactOpts)
}

// UnpauseTrading is a paid mutator transaction binding the contract method 0x456068d2.
//
// Solidity: function unpauseTrading() returns()
func (_NegRiskCTFExchange *NegRiskCTFExchangeTransactorSession) UnpauseTrading() (*types.Transaction, error) {
	return _NegRiskCTFExchange.Contract.UnpauseTrading(&_NegRiskCTFExchange.TransactOpts)
}

// NegRiskCTFExchangeFeeChargedIterator is returned from FilterFeeCharged and is used to iterate over the raw logs and unpacked data for FeeCharged events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeFeeChargedIterator struct {
	Event *NegRiskCTFExchangeFeeCharged // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeFeeChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeFeeCharged)
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
		it.Event = new(NegRiskCTFExchangeFeeCharged)
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
func (it *NegRiskCTFExchangeFeeChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeFeeChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeFeeCharged represents a FeeCharged event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeFeeCharged struct {
	Receiver common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeCharged is a free log retrieval operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterFeeCharged(opts *bind.FilterOpts, receiver []common.Address) (*NegRiskCTFExchangeFeeChargedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeFeeChargedIterator{contract: _NegRiskCTFExchange.contract, event: "FeeCharged", logs: logs, sub: sub}, nil
}

// WatchFeeCharged is a free log subscription operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchFeeCharged(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeFeeCharged, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "FeeCharged", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeFeeCharged)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
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

// ParseFeeCharged is a log parse operation binding the contract event 0xacffcc86834d0f1a64b0d5a675798deed6ff0bcfc2231edd3480e7288dba7ff4.
//
// Solidity: event FeeCharged(address indexed receiver, uint256 tokenId, uint256 amount)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseFeeCharged(log types.Log) (*NegRiskCTFExchangeFeeCharged, error) {
	event := new(NegRiskCTFExchangeFeeCharged)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "FeeCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeNewAdminIterator struct {
	Event *NegRiskCTFExchangeNewAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeNewAdmin)
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
		it.Event = new(NegRiskCTFExchangeNewAdmin)
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
func (it *NegRiskCTFExchangeNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeNewAdmin represents a NewAdmin event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeNewAdmin struct {
	NewAdminAddress common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterNewAdmin(opts *bind.FilterOpts, newAdminAddress []common.Address, admin []common.Address) (*NegRiskCTFExchangeNewAdminIterator, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeNewAdminIterator{contract: _NegRiskCTFExchange.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeNewAdmin, newAdminAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newAdminAddressRule []interface{}
	for _, newAdminAddressItem := range newAdminAddress {
		newAdminAddressRule = append(newAdminAddressRule, newAdminAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "NewAdmin", newAdminAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeNewAdmin)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address indexed newAdminAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseNewAdmin(log types.Log) (*NegRiskCTFExchangeNewAdmin, error) {
	event := new(NegRiskCTFExchangeNewAdmin)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeNewOperatorIterator is returned from FilterNewOperator and is used to iterate over the raw logs and unpacked data for NewOperator events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeNewOperatorIterator struct {
	Event *NegRiskCTFExchangeNewOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeNewOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeNewOperator)
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
		it.Event = new(NegRiskCTFExchangeNewOperator)
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
func (it *NegRiskCTFExchangeNewOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeNewOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeNewOperator represents a NewOperator event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeNewOperator struct {
	NewOperatorAddress common.Address
	Admin              common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewOperator is a free log retrieval operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterNewOperator(opts *bind.FilterOpts, newOperatorAddress []common.Address, admin []common.Address) (*NegRiskCTFExchangeNewOperatorIterator, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeNewOperatorIterator{contract: _NegRiskCTFExchange.contract, event: "NewOperator", logs: logs, sub: sub}, nil
}

// WatchNewOperator is a free log subscription operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchNewOperator(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeNewOperator, newOperatorAddress []common.Address, admin []common.Address) (event.Subscription, error) {

	var newOperatorAddressRule []interface{}
	for _, newOperatorAddressItem := range newOperatorAddress {
		newOperatorAddressRule = append(newOperatorAddressRule, newOperatorAddressItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "NewOperator", newOperatorAddressRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeNewOperator)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
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

// ParseNewOperator is a log parse operation binding the contract event 0xf1e04d73c4304b5ff164f9d10c7473e2a1593b740674a6107975e2a7001c1e5c.
//
// Solidity: event NewOperator(address indexed newOperatorAddress, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseNewOperator(log types.Log) (*NegRiskCTFExchangeNewOperator, error) {
	event := new(NegRiskCTFExchangeNewOperator)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "NewOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrderCancelledIterator struct {
	Event *NegRiskCTFExchangeOrderCancelled // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeOrderCancelled)
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
		it.Event = new(NegRiskCTFExchangeOrderCancelled)
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
func (it *NegRiskCTFExchangeOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeOrderCancelled represents a OrderCancelled event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrderCancelled struct {
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderHash [][32]byte) (*NegRiskCTFExchangeOrderCancelledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeOrderCancelledIterator{contract: _NegRiskCTFExchange.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeOrderCancelled, orderHash [][32]byte) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "OrderCancelled", orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeOrderCancelled)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed orderHash)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseOrderCancelled(log types.Log) (*NegRiskCTFExchangeOrderCancelled, error) {
	event := new(NegRiskCTFExchangeOrderCancelled)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrderFilledIterator struct {
	Event *NegRiskCTFExchangeOrderFilled // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeOrderFilled)
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
		it.Event = new(NegRiskCTFExchangeOrderFilled)
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
func (it *NegRiskCTFExchangeOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeOrderFilled represents a OrderFilled event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrderFilled struct {
	OrderHash         [32]byte
	Maker             common.Address
	Taker             common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Fee               *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterOrderFilled(opts *bind.FilterOpts, orderHash [][32]byte, maker []common.Address, taker []common.Address) (*NegRiskCTFExchangeOrderFilledIterator, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeOrderFilledIterator{contract: _NegRiskCTFExchange.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeOrderFilled, orderHash [][32]byte, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "OrderFilled", orderHashRule, makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeOrderFilled)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xd0a08e8c493f9c94f29311604c9de1b4e8c8d4c06bd0c789af57f2d65bfec0f6.
//
// Solidity: event OrderFilled(bytes32 indexed orderHash, address indexed maker, address indexed taker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled, uint256 fee)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseOrderFilled(log types.Log) (*NegRiskCTFExchangeOrderFilled, error) {
	event := new(NegRiskCTFExchangeOrderFilled)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeOrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrdersMatchedIterator struct {
	Event *NegRiskCTFExchangeOrdersMatched // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeOrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeOrdersMatched)
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
		it.Event = new(NegRiskCTFExchangeOrdersMatched)
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
func (it *NegRiskCTFExchangeOrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeOrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeOrdersMatched represents a OrdersMatched event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeOrdersMatched struct {
	TakerOrderHash    [32]byte
	TakerOrderMaker   common.Address
	MakerAssetId      *big.Int
	TakerAssetId      *big.Int
	MakerAmountFilled *big.Int
	TakerAmountFilled *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterOrdersMatched(opts *bind.FilterOpts, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (*NegRiskCTFExchangeOrdersMatchedIterator, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeOrdersMatchedIterator{contract: _NegRiskCTFExchange.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeOrdersMatched, takerOrderHash [][32]byte, takerOrderMaker []common.Address) (event.Subscription, error) {

	var takerOrderHashRule []interface{}
	for _, takerOrderHashItem := range takerOrderHash {
		takerOrderHashRule = append(takerOrderHashRule, takerOrderHashItem)
	}
	var takerOrderMakerRule []interface{}
	for _, takerOrderMakerItem := range takerOrderMaker {
		takerOrderMakerRule = append(takerOrderMakerRule, takerOrderMakerItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "OrdersMatched", takerOrderHashRule, takerOrderMakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeOrdersMatched)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x63bf4d16b7fa898ef4c4b2b6d90fd201e9c56313b65638af6088d149d2ce956c.
//
// Solidity: event OrdersMatched(bytes32 indexed takerOrderHash, address indexed takerOrderMaker, uint256 makerAssetId, uint256 takerAssetId, uint256 makerAmountFilled, uint256 takerAmountFilled)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseOrdersMatched(log types.Log) (*NegRiskCTFExchangeOrdersMatched, error) {
	event := new(NegRiskCTFExchangeOrdersMatched)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeProxyFactoryUpdatedIterator is returned from FilterProxyFactoryUpdated and is used to iterate over the raw logs and unpacked data for ProxyFactoryUpdated events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeProxyFactoryUpdatedIterator struct {
	Event *NegRiskCTFExchangeProxyFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeProxyFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeProxyFactoryUpdated)
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
		it.Event = new(NegRiskCTFExchangeProxyFactoryUpdated)
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
func (it *NegRiskCTFExchangeProxyFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeProxyFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeProxyFactoryUpdated represents a ProxyFactoryUpdated event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeProxyFactoryUpdated struct {
	OldProxyFactory common.Address
	NewProxyFactory common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterProxyFactoryUpdated is a free log retrieval operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterProxyFactoryUpdated(opts *bind.FilterOpts, oldProxyFactory []common.Address, newProxyFactory []common.Address) (*NegRiskCTFExchangeProxyFactoryUpdatedIterator, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeProxyFactoryUpdatedIterator{contract: _NegRiskCTFExchange.contract, event: "ProxyFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchProxyFactoryUpdated is a free log subscription operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchProxyFactoryUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeProxyFactoryUpdated, oldProxyFactory []common.Address, newProxyFactory []common.Address) (event.Subscription, error) {

	var oldProxyFactoryRule []interface{}
	for _, oldProxyFactoryItem := range oldProxyFactory {
		oldProxyFactoryRule = append(oldProxyFactoryRule, oldProxyFactoryItem)
	}
	var newProxyFactoryRule []interface{}
	for _, newProxyFactoryItem := range newProxyFactory {
		newProxyFactoryRule = append(newProxyFactoryRule, newProxyFactoryItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "ProxyFactoryUpdated", oldProxyFactoryRule, newProxyFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeProxyFactoryUpdated)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
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

// ParseProxyFactoryUpdated is a log parse operation binding the contract event 0x3053c6252a932554235c173caffc1913604dba3a41cee89516f631c4a1a50a37.
//
// Solidity: event ProxyFactoryUpdated(address indexed oldProxyFactory, address indexed newProxyFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseProxyFactoryUpdated(log types.Log) (*NegRiskCTFExchangeProxyFactoryUpdated, error) {
	event := new(NegRiskCTFExchangeProxyFactoryUpdated)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "ProxyFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeRemovedAdminIterator is returned from FilterRemovedAdmin and is used to iterate over the raw logs and unpacked data for RemovedAdmin events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeRemovedAdminIterator struct {
	Event *NegRiskCTFExchangeRemovedAdmin // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeRemovedAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeRemovedAdmin)
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
		it.Event = new(NegRiskCTFExchangeRemovedAdmin)
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
func (it *NegRiskCTFExchangeRemovedAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeRemovedAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeRemovedAdmin represents a RemovedAdmin event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeRemovedAdmin struct {
	RemovedAdmin common.Address
	Admin        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRemovedAdmin is a free log retrieval operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterRemovedAdmin(opts *bind.FilterOpts, removedAdmin []common.Address, admin []common.Address) (*NegRiskCTFExchangeRemovedAdminIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeRemovedAdminIterator{contract: _NegRiskCTFExchange.contract, event: "RemovedAdmin", logs: logs, sub: sub}, nil
}

// WatchRemovedAdmin is a free log subscription operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchRemovedAdmin(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeRemovedAdmin, removedAdmin []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "RemovedAdmin", removedAdminRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeRemovedAdmin)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
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

// ParseRemovedAdmin is a log parse operation binding the contract event 0x787a2e12f4a55b658b8f573c32432ee11a5e8b51677d1e1e937aaf6a0bb5776e.
//
// Solidity: event RemovedAdmin(address indexed removedAdmin, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseRemovedAdmin(log types.Log) (*NegRiskCTFExchangeRemovedAdmin, error) {
	event := new(NegRiskCTFExchangeRemovedAdmin)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "RemovedAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeRemovedOperatorIterator is returned from FilterRemovedOperator and is used to iterate over the raw logs and unpacked data for RemovedOperator events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeRemovedOperatorIterator struct {
	Event *NegRiskCTFExchangeRemovedOperator // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeRemovedOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeRemovedOperator)
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
		it.Event = new(NegRiskCTFExchangeRemovedOperator)
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
func (it *NegRiskCTFExchangeRemovedOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeRemovedOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeRemovedOperator represents a RemovedOperator event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeRemovedOperator struct {
	RemovedOperator common.Address
	Admin           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRemovedOperator is a free log retrieval operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterRemovedOperator(opts *bind.FilterOpts, removedOperator []common.Address, admin []common.Address) (*NegRiskCTFExchangeRemovedOperatorIterator, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeRemovedOperatorIterator{contract: _NegRiskCTFExchange.contract, event: "RemovedOperator", logs: logs, sub: sub}, nil
}

// WatchRemovedOperator is a free log subscription operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchRemovedOperator(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeRemovedOperator, removedOperator []common.Address, admin []common.Address) (event.Subscription, error) {

	var removedOperatorRule []interface{}
	for _, removedOperatorItem := range removedOperator {
		removedOperatorRule = append(removedOperatorRule, removedOperatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "RemovedOperator", removedOperatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeRemovedOperator)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
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

// ParseRemovedOperator is a log parse operation binding the contract event 0xf7262ed0443cc211121ceb1a80d69004f319245615a7488f951f1437fd91642c.
//
// Solidity: event RemovedOperator(address indexed removedOperator, address indexed admin)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseRemovedOperator(log types.Log) (*NegRiskCTFExchangeRemovedOperator, error) {
	event := new(NegRiskCTFExchangeRemovedOperator)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "RemovedOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeSafeFactoryUpdatedIterator is returned from FilterSafeFactoryUpdated and is used to iterate over the raw logs and unpacked data for SafeFactoryUpdated events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeSafeFactoryUpdatedIterator struct {
	Event *NegRiskCTFExchangeSafeFactoryUpdated // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeSafeFactoryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeSafeFactoryUpdated)
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
		it.Event = new(NegRiskCTFExchangeSafeFactoryUpdated)
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
func (it *NegRiskCTFExchangeSafeFactoryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeSafeFactoryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeSafeFactoryUpdated represents a SafeFactoryUpdated event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeSafeFactoryUpdated struct {
	OldSafeFactory common.Address
	NewSafeFactory common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSafeFactoryUpdated is a free log retrieval operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterSafeFactoryUpdated(opts *bind.FilterOpts, oldSafeFactory []common.Address, newSafeFactory []common.Address) (*NegRiskCTFExchangeSafeFactoryUpdatedIterator, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeSafeFactoryUpdatedIterator{contract: _NegRiskCTFExchange.contract, event: "SafeFactoryUpdated", logs: logs, sub: sub}, nil
}

// WatchSafeFactoryUpdated is a free log subscription operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchSafeFactoryUpdated(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeSafeFactoryUpdated, oldSafeFactory []common.Address, newSafeFactory []common.Address) (event.Subscription, error) {

	var oldSafeFactoryRule []interface{}
	for _, oldSafeFactoryItem := range oldSafeFactory {
		oldSafeFactoryRule = append(oldSafeFactoryRule, oldSafeFactoryItem)
	}
	var newSafeFactoryRule []interface{}
	for _, newSafeFactoryItem := range newSafeFactory {
		newSafeFactoryRule = append(newSafeFactoryRule, newSafeFactoryItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "SafeFactoryUpdated", oldSafeFactoryRule, newSafeFactoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeSafeFactoryUpdated)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
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

// ParseSafeFactoryUpdated is a log parse operation binding the contract event 0x9726d7faf7429d6b059560dc858ed769377ccdf8b7541eabe12b22548719831f.
//
// Solidity: event SafeFactoryUpdated(address indexed oldSafeFactory, address indexed newSafeFactory)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseSafeFactoryUpdated(log types.Log) (*NegRiskCTFExchangeSafeFactoryUpdated, error) {
	event := new(NegRiskCTFExchangeSafeFactoryUpdated)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "SafeFactoryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeTokenRegisteredIterator is returned from FilterTokenRegistered and is used to iterate over the raw logs and unpacked data for TokenRegistered events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTokenRegisteredIterator struct {
	Event *NegRiskCTFExchangeTokenRegistered // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeTokenRegistered)
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
		it.Event = new(NegRiskCTFExchangeTokenRegistered)
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
func (it *NegRiskCTFExchangeTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeTokenRegistered represents a TokenRegistered event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTokenRegistered struct {
	Token0      *big.Int
	Token1      *big.Int
	ConditionId [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenRegistered is a free log retrieval operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterTokenRegistered(opts *bind.FilterOpts, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (*NegRiskCTFExchangeTokenRegisteredIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeTokenRegisteredIterator{contract: _NegRiskCTFExchange.contract, event: "TokenRegistered", logs: logs, sub: sub}, nil
}

// WatchTokenRegistered is a free log subscription operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchTokenRegistered(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeTokenRegistered, token0 []*big.Int, token1 []*big.Int, conditionId [][32]byte) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}
	var conditionIdRule []interface{}
	for _, conditionIdItem := range conditionId {
		conditionIdRule = append(conditionIdRule, conditionIdItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "TokenRegistered", token0Rule, token1Rule, conditionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeTokenRegistered)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
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

// ParseTokenRegistered is a log parse operation binding the contract event 0xbc9a2432e8aeb48327246cddd6e872ef452812b4243c04e6bfb786a2cd8faf0d.
//
// Solidity: event TokenRegistered(uint256 indexed token0, uint256 indexed token1, bytes32 indexed conditionId)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseTokenRegistered(log types.Log) (*NegRiskCTFExchangeTokenRegistered, error) {
	event := new(NegRiskCTFExchangeTokenRegistered)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeTradingPausedIterator is returned from FilterTradingPaused and is used to iterate over the raw logs and unpacked data for TradingPaused events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTradingPausedIterator struct {
	Event *NegRiskCTFExchangeTradingPaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeTradingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeTradingPaused)
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
		it.Event = new(NegRiskCTFExchangeTradingPaused)
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
func (it *NegRiskCTFExchangeTradingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeTradingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeTradingPaused represents a TradingPaused event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTradingPaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingPaused is a free log retrieval operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterTradingPaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskCTFExchangeTradingPausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeTradingPausedIterator{contract: _NegRiskCTFExchange.contract, event: "TradingPaused", logs: logs, sub: sub}, nil
}

// WatchTradingPaused is a free log subscription operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchTradingPaused(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeTradingPaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "TradingPaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeTradingPaused)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
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

// ParseTradingPaused is a log parse operation binding the contract event 0x203c4bd3e526634f661575359ff30de3b0edaba6c2cb1eac60f730b6d2d9d536.
//
// Solidity: event TradingPaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseTradingPaused(log types.Log) (*NegRiskCTFExchangeTradingPaused, error) {
	event := new(NegRiskCTFExchangeTradingPaused)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TradingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NegRiskCTFExchangeTradingUnpausedIterator is returned from FilterTradingUnpaused and is used to iterate over the raw logs and unpacked data for TradingUnpaused events raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTradingUnpausedIterator struct {
	Event *NegRiskCTFExchangeTradingUnpaused // Event containing the contract specifics and raw log

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
func (it *NegRiskCTFExchangeTradingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NegRiskCTFExchangeTradingUnpaused)
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
		it.Event = new(NegRiskCTFExchangeTradingUnpaused)
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
func (it *NegRiskCTFExchangeTradingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NegRiskCTFExchangeTradingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NegRiskCTFExchangeTradingUnpaused represents a TradingUnpaused event raised by the NegRiskCTFExchange contract.
type NegRiskCTFExchangeTradingUnpaused struct {
	Pauser common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingUnpaused is a free log retrieval operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) FilterTradingUnpaused(opts *bind.FilterOpts, pauser []common.Address) (*NegRiskCTFExchangeTradingUnpausedIterator, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.FilterLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return &NegRiskCTFExchangeTradingUnpausedIterator{contract: _NegRiskCTFExchange.contract, event: "TradingUnpaused", logs: logs, sub: sub}, nil
}

// WatchTradingUnpaused is a free log subscription operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) WatchTradingUnpaused(opts *bind.WatchOpts, sink chan<- *NegRiskCTFExchangeTradingUnpaused, pauser []common.Address) (event.Subscription, error) {

	var pauserRule []interface{}
	for _, pauserItem := range pauser {
		pauserRule = append(pauserRule, pauserItem)
	}

	logs, sub, err := _NegRiskCTFExchange.contract.WatchLogs(opts, "TradingUnpaused", pauserRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NegRiskCTFExchangeTradingUnpaused)
				if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
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

// ParseTradingUnpaused is a log parse operation binding the contract event 0xa1e8a54850dbd7f520bcc09f47bff152294b77b2081da545a7adf531b7ea283b.
//
// Solidity: event TradingUnpaused(address indexed pauser)
func (_NegRiskCTFExchange *NegRiskCTFExchangeFilterer) ParseTradingUnpaused(log types.Log) (*NegRiskCTFExchangeTradingUnpaused, error) {
	event := new(NegRiskCTFExchangeTradingUnpaused)
	if err := _NegRiskCTFExchange.contract.UnpackLog(event, "TradingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
