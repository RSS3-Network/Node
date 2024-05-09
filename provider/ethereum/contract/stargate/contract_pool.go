// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stargate

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

// PoolChainPath is an auto generated low-level Go binding around an user-defined struct.
type PoolChainPath struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}

// PoolCreditObj is an auto generated low-level Go binding around an user-defined struct.
type PoolCreditObj struct {
	Credits      *big.Int
	IdealBalance *big.Int
}

// PoolSwapObj is an auto generated low-level Go binding around an user-defined struct.
type PoolSwapObj struct {
	Amount      *big.Int
	EqFee       *big.Int
	EqReward    *big.Int
	LpFee       *big.Int
	ProtocolFee *big.Int
	LkbRemove   *big.Int
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_poolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sharedDecimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_localDecimals\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeLibrary\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ChainPathUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"name\":\"CreditChainPath\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"batched\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapDeltaBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpDeltaBP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"defaultSwapMode\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"defaultLPMode\",\"type\":\"bool\"}],\"name\":\"DeltaParamUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeLibraryAddr\",\"type\":\"address\"}],\"name\":\"FeeLibraryUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeBP\",\"type\":\"uint256\"}],\"name\":\"FeesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"InstantRedeemLocal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintFeeAmountSD\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"to\",\"type\":\"bytes\"}],\"name\":\"RedeemLocal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountToMintSD\",\"type\":\"uint256\"}],\"name\":\"RedeemLocalCallback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLP\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"RedeemRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"name\":\"SendCredits\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"swapStop\",\"type\":\"bool\"}],\"name\":\"StopSwapUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"chainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstFee\",\"type\":\"uint256\"}],\"name\":\"SwapRemote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"WithdrawMintFeeBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"name\":\"WithdrawProtocolFeeBalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"srcChainId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"srcPoolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"WithdrawRemote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BP_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"activateChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"}],\"name\":\"amountLPtoLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batched\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_fullMode\",\"type\":\"bool\"}],\"name\":\"callDelta\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainPathIndexLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainPaths\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"name\":\"createChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.CreditObj\",\"name\":\"_c\",\"type\":\"tuple\"}],\"name\":\"creditChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultLPMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultSwapMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deltaCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eqFeePool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeLibrary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"getChainPath\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"ready\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkb\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.ChainPath\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainPathsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"instantRedeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpDeltaBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_to\",\"type\":\"bytes\"}],\"name\":\"redeemLocal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountToMintSD\",\"type\":\"uint256\"}],\"name\":\"redeemLocalCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amountSD\",\"type\":\"uint256\"}],\"name\":\"redeemLocalCheckOnRemote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLP\",\"type\":\"uint256\"}],\"name\":\"redeemRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"}],\"name\":\"sendCredits\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"idealBalance\",\"type\":\"uint256\"}],\"internalType\":\"structPool.CreditObj\",\"name\":\"c\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_batched\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_swapDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lpDeltaBP\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_defaultSwapMode\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_defaultLPMode\",\"type\":\"bool\"}],\"name\":\"setDeltaParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_mintFeeBP\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeLibraryAddr\",\"type\":\"address\"}],\"name\":\"setFeeLibrary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_swapStop\",\"type\":\"bool\"}],\"name\":\"setSwapStop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_weight\",\"type\":\"uint16\"}],\"name\":\"setWeightForChainPath\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sharedDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_dstChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_dstPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountLD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minAmountLD\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"newLiquidity\",\"type\":\"bool\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkbRemove\",\"type\":\"uint256\"}],\"internalType\":\"structPool.SwapObj\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapDeltaBP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_srcChainId\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"_srcPoolId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eqReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lpFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lkbRemove\",\"type\":\"uint256\"}],\"internalType\":\"structPool.SwapObj\",\"name\":\"_s\",\"type\":\"tuple\"}],\"name\":\"swapRemote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountLD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawMintFeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdrawProtocolFeeBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_Pool *PoolCaller) BPDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "BP_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_Pool *PoolSession) BPDENOMINATOR() (*big.Int, error) {
	return _Pool.Contract.BPDENOMINATOR(&_Pool.CallOpts)
}

// BPDENOMINATOR is a free data retrieval call binding the contract method 0xabe685cd.
//
// Solidity: function BP_DENOMINATOR() view returns(uint256)
func (_Pool *PoolCallerSession) BPDENOMINATOR() (*big.Int, error) {
	return _Pool.Contract.BPDENOMINATOR(&_Pool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pool.Contract.DOMAINSEPARATOR(&_Pool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pool *PoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pool.Contract.DOMAINSEPARATOR(&_Pool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pool *PoolCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pool *PoolSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pool.Contract.PERMITTYPEHASH(&_Pool.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pool *PoolCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pool.Contract.PERMITTYPEHASH(&_Pool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pool *PoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pool.Contract.Allowance(&_Pool.CallOpts, arg0, arg1)
}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_Pool *PoolCaller) AmountLPtoLD(opts *bind.CallOpts, _amountLP *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "amountLPtoLD", _amountLP)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_Pool *PoolSession) AmountLPtoLD(_amountLP *big.Int) (*big.Int, error) {
	return _Pool.Contract.AmountLPtoLD(&_Pool.CallOpts, _amountLP)
}

// AmountLPtoLD is a free data retrieval call binding the contract method 0xf6cd35ee.
//
// Solidity: function amountLPtoLD(uint256 _amountLP) view returns(uint256)
func (_Pool *PoolCallerSession) AmountLPtoLD(_amountLP *big.Int) (*big.Int, error) {
	return _Pool.Contract.AmountLPtoLD(&_Pool.CallOpts, _amountLP)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pool *PoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pool *PoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pool *PoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, arg0)
}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_Pool *PoolCaller) Batched(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "batched")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_Pool *PoolSession) Batched() (bool, error) {
	return _Pool.Contract.Batched(&_Pool.CallOpts)
}

// Batched is a free data retrieval call binding the contract method 0x27f92376.
//
// Solidity: function batched() view returns(bool)
func (_Pool *PoolCallerSession) Batched() (bool, error) {
	return _Pool.Contract.Batched(&_Pool.CallOpts)
}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_Pool *PoolCaller) ChainPathIndexLookup(opts *bind.CallOpts, arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "chainPathIndexLookup", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_Pool *PoolSession) ChainPathIndexLookup(arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	return _Pool.Contract.ChainPathIndexLookup(&_Pool.CallOpts, arg0, arg1)
}

// ChainPathIndexLookup is a free data retrieval call binding the contract method 0x64c5f02d.
//
// Solidity: function chainPathIndexLookup(uint16 , uint256 ) view returns(uint256)
func (_Pool *PoolCallerSession) ChainPathIndexLookup(arg0 uint16, arg1 *big.Int) (*big.Int, error) {
	return _Pool.Contract.ChainPathIndexLookup(&_Pool.CallOpts, arg0, arg1)
}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_Pool *PoolCaller) ChainPaths(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "chainPaths", arg0)

	outstruct := new(struct {
		Ready        bool
		DstChainId   uint16
		DstPoolId    *big.Int
		Weight       *big.Int
		Balance      *big.Int
		Lkb          *big.Int
		Credits      *big.Int
		IdealBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ready = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DstChainId = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.DstPoolId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Weight = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Lkb = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Credits = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IdealBalance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_Pool *PoolSession) ChainPaths(arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	return _Pool.Contract.ChainPaths(&_Pool.CallOpts, arg0)
}

// ChainPaths is a free data retrieval call binding the contract method 0xa138ed6b.
//
// Solidity: function chainPaths(uint256 ) view returns(bool ready, uint16 dstChainId, uint256 dstPoolId, uint256 weight, uint256 balance, uint256 lkb, uint256 credits, uint256 idealBalance)
func (_Pool *PoolCallerSession) ChainPaths(arg0 *big.Int) (struct {
	Ready        bool
	DstChainId   uint16
	DstPoolId    *big.Int
	Weight       *big.Int
	Balance      *big.Int
	Lkb          *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
}, error) {
	return _Pool.Contract.ChainPaths(&_Pool.CallOpts, arg0)
}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_Pool *PoolCaller) ConvertRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "convertRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_Pool *PoolSession) ConvertRate() (*big.Int, error) {
	return _Pool.Contract.ConvertRate(&_Pool.CallOpts)
}

// ConvertRate is a free data retrieval call binding the contract method 0xfeb56b15.
//
// Solidity: function convertRate() view returns(uint256)
func (_Pool *PoolCallerSession) ConvertRate() (*big.Int, error) {
	return _Pool.Contract.ConvertRate(&_Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pool *PoolCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pool *PoolSession) Decimals() (*big.Int, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Pool *PoolCallerSession) Decimals() (*big.Int, error) {
	return _Pool.Contract.Decimals(&_Pool.CallOpts)
}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_Pool *PoolCaller) DefaultLPMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "defaultLPMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_Pool *PoolSession) DefaultLPMode() (bool, error) {
	return _Pool.Contract.DefaultLPMode(&_Pool.CallOpts)
}

// DefaultLPMode is a free data retrieval call binding the contract method 0x28f079c2.
//
// Solidity: function defaultLPMode() view returns(bool)
func (_Pool *PoolCallerSession) DefaultLPMode() (bool, error) {
	return _Pool.Contract.DefaultLPMode(&_Pool.CallOpts)
}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_Pool *PoolCaller) DefaultSwapMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "defaultSwapMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_Pool *PoolSession) DefaultSwapMode() (bool, error) {
	return _Pool.Contract.DefaultSwapMode(&_Pool.CallOpts)
}

// DefaultSwapMode is a free data retrieval call binding the contract method 0x99a22d68.
//
// Solidity: function defaultSwapMode() view returns(bool)
func (_Pool *PoolCallerSession) DefaultSwapMode() (bool, error) {
	return _Pool.Contract.DefaultSwapMode(&_Pool.CallOpts)
}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_Pool *PoolCaller) DeltaCredit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "deltaCredit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_Pool *PoolSession) DeltaCredit() (*big.Int, error) {
	return _Pool.Contract.DeltaCredit(&_Pool.CallOpts)
}

// DeltaCredit is a free data retrieval call binding the contract method 0x1e8e51da.
//
// Solidity: function deltaCredit() view returns(uint256)
func (_Pool *PoolCallerSession) DeltaCredit() (*big.Int, error) {
	return _Pool.Contract.DeltaCredit(&_Pool.CallOpts)
}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_Pool *PoolCaller) EqFeePool(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "eqFeePool")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_Pool *PoolSession) EqFeePool() (*big.Int, error) {
	return _Pool.Contract.EqFeePool(&_Pool.CallOpts)
}

// EqFeePool is a free data retrieval call binding the contract method 0x9bb81119.
//
// Solidity: function eqFeePool() view returns(uint256)
func (_Pool *PoolCallerSession) EqFeePool() (*big.Int, error) {
	return _Pool.Contract.EqFeePool(&_Pool.CallOpts)
}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_Pool *PoolCaller) FeeLibrary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "feeLibrary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_Pool *PoolSession) FeeLibrary() (common.Address, error) {
	return _Pool.Contract.FeeLibrary(&_Pool.CallOpts)
}

// FeeLibrary is a free data retrieval call binding the contract method 0x001edfab.
//
// Solidity: function feeLibrary() view returns(address)
func (_Pool *PoolCallerSession) FeeLibrary() (common.Address, error) {
	return _Pool.Contract.FeeLibrary(&_Pool.CallOpts)
}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolCaller) GetChainPath(opts *bind.CallOpts, _dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getChainPath", _dstChainId, _dstPoolId)

	if err != nil {
		return *new(PoolChainPath), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolChainPath)).(*PoolChainPath)

	return out0, err

}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolSession) GetChainPath(_dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	return _Pool.Contract.GetChainPath(&_Pool.CallOpts, _dstChainId, _dstPoolId)
}

// GetChainPath is a free data retrieval call binding the contract method 0x159f6add.
//
// Solidity: function getChainPath(uint16 _dstChainId, uint256 _dstPoolId) view returns((bool,uint16,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolCallerSession) GetChainPath(_dstChainId uint16, _dstPoolId *big.Int) (PoolChainPath, error) {
	return _Pool.Contract.GetChainPath(&_Pool.CallOpts, _dstChainId, _dstPoolId)
}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_Pool *PoolCaller) GetChainPathsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getChainPathsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_Pool *PoolSession) GetChainPathsLength() (*big.Int, error) {
	return _Pool.Contract.GetChainPathsLength(&_Pool.CallOpts)
}

// GetChainPathsLength is a free data retrieval call binding the contract method 0x163ef490.
//
// Solidity: function getChainPathsLength() view returns(uint256)
func (_Pool *PoolCallerSession) GetChainPathsLength() (*big.Int, error) {
	return _Pool.Contract.GetChainPathsLength(&_Pool.CallOpts)
}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_Pool *PoolCaller) LocalDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "localDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_Pool *PoolSession) LocalDecimals() (*big.Int, error) {
	return _Pool.Contract.LocalDecimals(&_Pool.CallOpts)
}

// LocalDecimals is a free data retrieval call binding the contract method 0xe46e7058.
//
// Solidity: function localDecimals() view returns(uint256)
func (_Pool *PoolCallerSession) LocalDecimals() (*big.Int, error) {
	return _Pool.Contract.LocalDecimals(&_Pool.CallOpts)
}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_Pool *PoolCaller) LpDeltaBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "lpDeltaBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_Pool *PoolSession) LpDeltaBP() (*big.Int, error) {
	return _Pool.Contract.LpDeltaBP(&_Pool.CallOpts)
}

// LpDeltaBP is a free data retrieval call binding the contract method 0x36448777.
//
// Solidity: function lpDeltaBP() view returns(uint256)
func (_Pool *PoolCallerSession) LpDeltaBP() (*big.Int, error) {
	return _Pool.Contract.LpDeltaBP(&_Pool.CallOpts)
}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_Pool *PoolCaller) MintFeeBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "mintFeeBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_Pool *PoolSession) MintFeeBP() (*big.Int, error) {
	return _Pool.Contract.MintFeeBP(&_Pool.CallOpts)
}

// MintFeeBP is a free data retrieval call binding the contract method 0xfaa24f07.
//
// Solidity: function mintFeeBP() view returns(uint256)
func (_Pool *PoolCallerSession) MintFeeBP() (*big.Int, error) {
	return _Pool.Contract.MintFeeBP(&_Pool.CallOpts)
}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_Pool *PoolCaller) MintFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "mintFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_Pool *PoolSession) MintFeeBalance() (*big.Int, error) {
	return _Pool.Contract.MintFeeBalance(&_Pool.CallOpts)
}

// MintFeeBalance is a free data retrieval call binding the contract method 0x65152f2b.
//
// Solidity: function mintFeeBalance() view returns(uint256)
func (_Pool *PoolCallerSession) MintFeeBalance() (*big.Int, error) {
	return _Pool.Contract.MintFeeBalance(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pool *PoolCallerSession) Name() (string, error) {
	return _Pool.Contract.Name(&_Pool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pool *PoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pool *PoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.Nonces(&_Pool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pool *PoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pool.Contract.Nonces(&_Pool.CallOpts, arg0)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Pool *PoolCaller) PoolId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "poolId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Pool *PoolSession) PoolId() (*big.Int, error) {
	return _Pool.Contract.PoolId(&_Pool.CallOpts)
}

// PoolId is a free data retrieval call binding the contract method 0x3e0dc34e.
//
// Solidity: function poolId() view returns(uint256)
func (_Pool *PoolCallerSession) PoolId() (*big.Int, error) {
	return _Pool.Contract.PoolId(&_Pool.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_Pool *PoolCaller) ProtocolFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "protocolFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_Pool *PoolSession) ProtocolFeeBalance() (*big.Int, error) {
	return _Pool.Contract.ProtocolFeeBalance(&_Pool.CallOpts)
}

// ProtocolFeeBalance is a free data retrieval call binding the contract method 0x0a22d68c.
//
// Solidity: function protocolFeeBalance() view returns(uint256)
func (_Pool *PoolCallerSession) ProtocolFeeBalance() (*big.Int, error) {
	return _Pool.Contract.ProtocolFeeBalance(&_Pool.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Pool *PoolCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Pool *PoolSession) Router() (common.Address, error) {
	return _Pool.Contract.Router(&_Pool.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Pool *PoolCallerSession) Router() (common.Address, error) {
	return _Pool.Contract.Router(&_Pool.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_Pool *PoolCaller) SharedDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "sharedDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_Pool *PoolSession) SharedDecimals() (*big.Int, error) {
	return _Pool.Contract.SharedDecimals(&_Pool.CallOpts)
}

// SharedDecimals is a free data retrieval call binding the contract method 0x857749b0.
//
// Solidity: function sharedDecimals() view returns(uint256)
func (_Pool *PoolCallerSession) SharedDecimals() (*big.Int, error) {
	return _Pool.Contract.SharedDecimals(&_Pool.CallOpts)
}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_Pool *PoolCaller) StopSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "stopSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_Pool *PoolSession) StopSwap() (bool, error) {
	return _Pool.Contract.StopSwap(&_Pool.CallOpts)
}

// StopSwap is a free data retrieval call binding the contract method 0xb633b364.
//
// Solidity: function stopSwap() view returns(bool)
func (_Pool *PoolCallerSession) StopSwap() (bool, error) {
	return _Pool.Contract.StopSwap(&_Pool.CallOpts)
}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_Pool *PoolCaller) SwapDeltaBP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "swapDeltaBP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_Pool *PoolSession) SwapDeltaBP() (*big.Int, error) {
	return _Pool.Contract.SwapDeltaBP(&_Pool.CallOpts)
}

// SwapDeltaBP is a free data retrieval call binding the contract method 0xcdfed0ab.
//
// Solidity: function swapDeltaBP() view returns(uint256)
func (_Pool *PoolCallerSession) SwapDeltaBP() (*big.Int, error) {
	return _Pool.Contract.SwapDeltaBP(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pool *PoolCallerSession) Symbol() (string, error) {
	return _Pool.Contract.Symbol(&_Pool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pool *PoolCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pool *PoolSession) Token() (common.Address, error) {
	return _Pool.Contract.Token(&_Pool.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Pool *PoolCallerSession) Token() (common.Address, error) {
	return _Pool.Contract.Token(&_Pool.CallOpts)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_Pool *PoolCaller) TotalLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_Pool *PoolSession) TotalLiquidity() (*big.Int, error) {
	return _Pool.Contract.TotalLiquidity(&_Pool.CallOpts)
}

// TotalLiquidity is a free data retrieval call binding the contract method 0x15770f92.
//
// Solidity: function totalLiquidity() view returns(uint256)
func (_Pool *PoolCallerSession) TotalLiquidity() (*big.Int, error) {
	return _Pool.Contract.TotalLiquidity(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pool *PoolCallerSession) TotalSupply() (*big.Int, error) {
	return _Pool.Contract.TotalSupply(&_Pool.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Pool *PoolCaller) TotalWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Pool *PoolSession) TotalWeight() (*big.Int, error) {
	return _Pool.Contract.TotalWeight(&_Pool.CallOpts)
}

// TotalWeight is a free data retrieval call binding the contract method 0x96c82e57.
//
// Solidity: function totalWeight() view returns(uint256)
func (_Pool *PoolCallerSession) TotalWeight() (*big.Int, error) {
	return _Pool.Contract.TotalWeight(&_Pool.CallOpts)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Pool *PoolTransactor) ActivateChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "activateChainPath", _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Pool *PoolSession) ActivateChainPath(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ActivateChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId)
}

// ActivateChainPath is a paid mutator transaction binding the contract method 0x8bd86d0a.
//
// Solidity: function activateChainPath(uint16 _dstChainId, uint256 _dstPoolId) returns()
func (_Pool *PoolTransactorSession) ActivateChainPath(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.ActivateChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pool *PoolTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pool *PoolSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pool *PoolTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Approve(&_Pool.TransactOpts, spender, value)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_Pool *PoolTransactor) CallDelta(opts *bind.TransactOpts, _fullMode bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "callDelta", _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_Pool *PoolSession) CallDelta(_fullMode bool) (*types.Transaction, error) {
	return _Pool.Contract.CallDelta(&_Pool.TransactOpts, _fullMode)
}

// CallDelta is a paid mutator transaction binding the contract method 0x7fb65265.
//
// Solidity: function callDelta(bool _fullMode) returns()
func (_Pool *PoolTransactorSession) CallDelta(_fullMode bool) (*types.Transaction, error) {
	return _Pool.Contract.CallDelta(&_Pool.TransactOpts, _fullMode)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Pool *PoolTransactor) CreateChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "createChainPath", _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Pool *PoolSession) CreateChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.CreateChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// CreateChainPath is a paid mutator transaction binding the contract method 0x20d6bc75.
//
// Solidity: function createChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint256 _weight) returns()
func (_Pool *PoolTransactorSession) CreateChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.CreateChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_Pool *PoolTransactor) CreditChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "creditChainPath", _dstChainId, _dstPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_Pool *PoolSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Pool.Contract.CreditChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _c)
}

// CreditChainPath is a paid mutator transaction binding the contract method 0xb6addec7.
//
// Solidity: function creditChainPath(uint16 _dstChainId, uint256 _dstPoolId, (uint256,uint256) _c) returns()
func (_Pool *PoolTransactorSession) CreditChainPath(_dstChainId uint16, _dstPoolId *big.Int, _c PoolCreditObj) (*types.Transaction, error) {
	return _Pool.Contract.CreditChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _c)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Pool *PoolTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.DecreaseAllowance(&_Pool.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Pool *PoolTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.IncreaseAllowance(&_Pool.TransactOpts, spender, addedValue)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Pool *PoolTransactor) InstantRedeemLocal(opts *bind.TransactOpts, _from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "instantRedeemLocal", _from, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Pool *PoolSession) InstantRedeemLocal(_from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.InstantRedeemLocal(&_Pool.TransactOpts, _from, _amountLP, _to)
}

// InstantRedeemLocal is a paid mutator transaction binding the contract method 0x0986b61a.
//
// Solidity: function instantRedeemLocal(address _from, uint256 _amountLP, address _to) returns(uint256 amountSD)
func (_Pool *PoolTransactorSession) InstantRedeemLocal(_from common.Address, _amountLP *big.Int, _to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.InstantRedeemLocal(&_Pool.TransactOpts, _from, _amountLP, _to)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_Pool *PoolTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "mint", _to, _amountLD)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_Pool *PoolSession) Mint(_to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, _to, _amountLD)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amountLD) returns(uint256)
func (_Pool *PoolTransactorSession) Mint(_to common.Address, _amountLD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Mint(&_Pool.TransactOpts, _to, _amountLD)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pool *PoolTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pool.Contract.Permit(&_Pool.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_Pool *PoolTransactor) RedeemLocal(opts *bind.TransactOpts, _from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeemLocal", _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_Pool *PoolSession) RedeemLocal(_from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocal(&_Pool.TransactOpts, _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocal is a paid mutator transaction binding the contract method 0xb0fab0bc.
//
// Solidity: function redeemLocal(address _from, uint256 _amountLP, uint16 _dstChainId, uint256 _dstPoolId, bytes _to) returns(uint256 amountSD)
func (_Pool *PoolTransactorSession) RedeemLocal(_from common.Address, _amountLP *big.Int, _dstChainId uint16, _dstPoolId *big.Int, _to []byte) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocal(&_Pool.TransactOpts, _from, _amountLP, _dstChainId, _dstPoolId, _to)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_Pool *PoolTransactor) RedeemLocalCallback(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeemLocalCallback", _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_Pool *PoolSession) RedeemLocalCallback(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocalCallback(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCallback is a paid mutator transaction binding the contract method 0xb30daeac.
//
// Solidity: function redeemLocalCallback(uint16 _srcChainId, uint256 _srcPoolId, address _to, uint256 _amountSD, uint256 _amountToMintSD) returns()
func (_Pool *PoolTransactorSession) RedeemLocalCallback(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _amountSD *big.Int, _amountToMintSD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocalCallback(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _to, _amountSD, _amountToMintSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolTransactor) RedeemLocalCheckOnRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeemLocalCheckOnRemote", _srcChainId, _srcPoolId, _amountSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocalCheckOnRemote(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _amountSD)
}

// RedeemLocalCheckOnRemote is a paid mutator transaction binding the contract method 0xea89e2aa.
//
// Solidity: function redeemLocalCheckOnRemote(uint16 _srcChainId, uint256 _srcPoolId, uint256 _amountSD) returns(uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolTransactorSession) RedeemLocalCheckOnRemote(_srcChainId uint16, _srcPoolId *big.Int, _amountSD *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemLocalCheckOnRemote(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _amountSD)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_Pool *PoolTransactor) RedeemRemote(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeemRemote", _dstChainId, _dstPoolId, _from, _amountLP)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_Pool *PoolSession) RedeemRemote(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemRemote(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLP)
}

// RedeemRemote is a paid mutator transaction binding the contract method 0x7298a5dc.
//
// Solidity: function redeemRemote(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLP) returns()
func (_Pool *PoolTransactorSession) RedeemRemote(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLP *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.RedeemRemote(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLP)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_Pool *PoolTransactor) SendCredits(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "sendCredits", _dstChainId, _dstPoolId)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_Pool *PoolSession) SendCredits(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SendCredits(&_Pool.TransactOpts, _dstChainId, _dstPoolId)
}

// SendCredits is a paid mutator transaction binding the contract method 0x08e9d8c2.
//
// Solidity: function sendCredits(uint16 _dstChainId, uint256 _dstPoolId) returns((uint256,uint256) c)
func (_Pool *PoolTransactorSession) SendCredits(_dstChainId uint16, _dstPoolId *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SendCredits(&_Pool.TransactOpts, _dstChainId, _dstPoolId)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Pool *PoolTransactor) SetDeltaParam(opts *bind.TransactOpts, _batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setDeltaParam", _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Pool *PoolSession) SetDeltaParam(_batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Pool.Contract.SetDeltaParam(&_Pool.TransactOpts, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetDeltaParam is a paid mutator transaction binding the contract method 0xe065608b.
//
// Solidity: function setDeltaParam(bool _batched, uint256 _swapDeltaBP, uint256 _lpDeltaBP, bool _defaultSwapMode, bool _defaultLPMode) returns()
func (_Pool *PoolTransactorSession) SetDeltaParam(_batched bool, _swapDeltaBP *big.Int, _lpDeltaBP *big.Int, _defaultSwapMode bool, _defaultLPMode bool) (*types.Transaction, error) {
	return _Pool.Contract.SetDeltaParam(&_Pool.TransactOpts, _batched, _swapDeltaBP, _lpDeltaBP, _defaultSwapMode, _defaultLPMode)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_Pool *PoolTransactor) SetFee(opts *bind.TransactOpts, _mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFee", _mintFeeBP)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_Pool *PoolSession) SetFee(_mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, _mintFeeBP)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 _mintFeeBP) returns()
func (_Pool *PoolTransactorSession) SetFee(_mintFeeBP *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, _mintFeeBP)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_Pool *PoolTransactor) SetFeeLibrary(opts *bind.TransactOpts, _feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFeeLibrary", _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_Pool *PoolSession) SetFeeLibrary(_feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFeeLibrary(&_Pool.TransactOpts, _feeLibraryAddr)
}

// SetFeeLibrary is a paid mutator transaction binding the contract method 0x4b5cacbc.
//
// Solidity: function setFeeLibrary(address _feeLibraryAddr) returns()
func (_Pool *PoolTransactorSession) SetFeeLibrary(_feeLibraryAddr common.Address) (*types.Transaction, error) {
	return _Pool.Contract.SetFeeLibrary(&_Pool.TransactOpts, _feeLibraryAddr)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_Pool *PoolTransactor) SetSwapStop(opts *bind.TransactOpts, _swapStop bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setSwapStop", _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_Pool *PoolSession) SetSwapStop(_swapStop bool) (*types.Transaction, error) {
	return _Pool.Contract.SetSwapStop(&_Pool.TransactOpts, _swapStop)
}

// SetSwapStop is a paid mutator transaction binding the contract method 0xac2cc36b.
//
// Solidity: function setSwapStop(bool _swapStop) returns()
func (_Pool *PoolTransactorSession) SetSwapStop(_swapStop bool) (*types.Transaction, error) {
	return _Pool.Contract.SetSwapStop(&_Pool.TransactOpts, _swapStop)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Pool *PoolTransactor) SetWeightForChainPath(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setWeightForChainPath", _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Pool *PoolSession) SetWeightForChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Pool.Contract.SetWeightForChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// SetWeightForChainPath is a paid mutator transaction binding the contract method 0xa985565f.
//
// Solidity: function setWeightForChainPath(uint16 _dstChainId, uint256 _dstPoolId, uint16 _weight) returns()
func (_Pool *PoolTransactorSession) SetWeightForChainPath(_dstChainId uint16, _dstPoolId *big.Int, _weight uint16) (*types.Transaction, error) {
	return _Pool.Contract.SetWeightForChainPath(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _weight)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolTransactor) Swap(opts *bind.TransactOpts, _dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swap", _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolSession) Swap(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// Swap is a paid mutator transaction binding the contract method 0x1b7319b6.
//
// Solidity: function swap(uint16 _dstChainId, uint256 _dstPoolId, address _from, uint256 _amountLD, uint256 _minAmountLD, bool newLiquidity) returns((uint256,uint256,uint256,uint256,uint256,uint256))
func (_Pool *PoolTransactorSession) Swap(_dstChainId uint16, _dstPoolId *big.Int, _from common.Address, _amountLD *big.Int, _minAmountLD *big.Int, newLiquidity bool) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, _dstChainId, _dstPoolId, _from, _amountLD, _minAmountLD, newLiquidity)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_Pool *PoolTransactor) SwapRemote(opts *bind.TransactOpts, _srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swapRemote", _srcChainId, _srcPoolId, _to, _s)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_Pool *PoolSession) SwapRemote(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _Pool.Contract.SwapRemote(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _to, _s)
}

// SwapRemote is a paid mutator transaction binding the contract method 0x902b8ab7.
//
// Solidity: function swapRemote(uint16 _srcChainId, uint256 _srcPoolId, address _to, (uint256,uint256,uint256,uint256,uint256,uint256) _s) returns(uint256 amountLD)
func (_Pool *PoolTransactorSession) SwapRemote(_srcChainId uint16, _srcPoolId *big.Int, _to common.Address, _s PoolSwapObj) (*types.Transaction, error) {
	return _Pool.Contract.SwapRemote(&_Pool.TransactOpts, _srcChainId, _srcPoolId, _to, _s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pool *PoolTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pool *PoolSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pool *PoolTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Transfer(&_Pool.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pool *PoolTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pool *PoolSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pool *PoolTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.TransferFrom(&_Pool.TransactOpts, from, to, value)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_Pool *PoolTransactor) WithdrawMintFeeBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdrawMintFeeBalance", _to)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_Pool *PoolSession) WithdrawMintFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawMintFeeBalance(&_Pool.TransactOpts, _to)
}

// WithdrawMintFeeBalance is a paid mutator transaction binding the contract method 0x476efe40.
//
// Solidity: function withdrawMintFeeBalance(address _to) returns()
func (_Pool *PoolTransactorSession) WithdrawMintFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawMintFeeBalance(&_Pool.TransactOpts, _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_Pool *PoolTransactor) WithdrawProtocolFeeBalance(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "withdrawProtocolFeeBalance", _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_Pool *PoolSession) WithdrawProtocolFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawProtocolFeeBalance(&_Pool.TransactOpts, _to)
}

// WithdrawProtocolFeeBalance is a paid mutator transaction binding the contract method 0xbe310294.
//
// Solidity: function withdrawProtocolFeeBalance(address _to) returns()
func (_Pool *PoolTransactorSession) WithdrawProtocolFeeBalance(_to common.Address) (*types.Transaction, error) {
	return _Pool.Contract.WithdrawProtocolFeeBalance(&_Pool.TransactOpts, _to)
}

// PoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pool contract.
type PoolApprovalIterator struct {
	Event *PoolApproval // Event containing the contract specifics and raw log

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
func (it *PoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolApproval)
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
		it.Event = new(PoolApproval)
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
func (it *PoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolApproval represents a Approval event raised by the Pool contract.
type PoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolApprovalIterator{contract: _Pool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolApproval)
				if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pool *PoolFilterer) ParseApproval(log types.Log) (*PoolApproval, error) {
	event := new(PoolApproval)
	if err := _Pool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pool contract.
type PoolBurnIterator struct {
	Event *PoolBurn // Event containing the contract specifics and raw log

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
func (it *PoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBurn)
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
		it.Event = new(PoolBurn)
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
func (it *PoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBurn represents a Burn event raised by the Pool contract.
type PoolBurn struct {
	From     common.Address
	AmountLP *big.Int
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) FilterBurn(opts *bind.FilterOpts) (*PoolBurnIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &PoolBurnIterator{contract: _Pool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PoolBurn) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBurn)
				if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) ParseBurn(log types.Log) (*PoolBurn, error) {
	event := new(PoolBurn)
	if err := _Pool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolChainPathUpdateIterator is returned from FilterChainPathUpdate and is used to iterate over the raw logs and unpacked data for ChainPathUpdate events raised by the Pool contract.
type PoolChainPathUpdateIterator struct {
	Event *PoolChainPathUpdate // Event containing the contract specifics and raw log

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
func (it *PoolChainPathUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolChainPathUpdate)
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
		it.Event = new(PoolChainPathUpdate)
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
func (it *PoolChainPathUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolChainPathUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolChainPathUpdate represents a ChainPathUpdate event raised by the Pool contract.
type PoolChainPathUpdate struct {
	DstChainId uint16
	DstPoolId  *big.Int
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChainPathUpdate is a free log retrieval operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_Pool *PoolFilterer) FilterChainPathUpdate(opts *bind.FilterOpts) (*PoolChainPathUpdateIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ChainPathUpdate")
	if err != nil {
		return nil, err
	}
	return &PoolChainPathUpdateIterator{contract: _Pool.contract, event: "ChainPathUpdate", logs: logs, sub: sub}, nil
}

// WatchChainPathUpdate is a free log subscription operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_Pool *PoolFilterer) WatchChainPathUpdate(opts *bind.WatchOpts, sink chan<- *PoolChainPathUpdate) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ChainPathUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolChainPathUpdate)
				if err := _Pool.contract.UnpackLog(event, "ChainPathUpdate", log); err != nil {
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

// ParseChainPathUpdate is a log parse operation binding the contract event 0x8fb3b21a941c2361df46475f9ae2f7b5dac5de7bd085fa22415ec0bb30c77e22.
//
// Solidity: event ChainPathUpdate(uint16 dstChainId, uint256 dstPoolId, uint256 weight)
func (_Pool *PoolFilterer) ParseChainPathUpdate(log types.Log) (*PoolChainPathUpdate, error) {
	event := new(PoolChainPathUpdate)
	if err := _Pool.contract.UnpackLog(event, "ChainPathUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolCreditChainPathIterator is returned from FilterCreditChainPath and is used to iterate over the raw logs and unpacked data for CreditChainPath events raised by the Pool contract.
type PoolCreditChainPathIterator struct {
	Event *PoolCreditChainPath // Event containing the contract specifics and raw log

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
func (it *PoolCreditChainPathIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCreditChainPath)
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
		it.Event = new(PoolCreditChainPath)
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
func (it *PoolCreditChainPathIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCreditChainPathIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCreditChainPath represents a CreditChainPath event raised by the Pool contract.
type PoolCreditChainPath struct {
	ChainId      uint16
	SrcPoolId    *big.Int
	AmountSD     *big.Int
	IdealBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCreditChainPath is a free log retrieval operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_Pool *PoolFilterer) FilterCreditChainPath(opts *bind.FilterOpts) (*PoolCreditChainPathIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "CreditChainPath")
	if err != nil {
		return nil, err
	}
	return &PoolCreditChainPathIterator{contract: _Pool.contract, event: "CreditChainPath", logs: logs, sub: sub}, nil
}

// WatchCreditChainPath is a free log subscription operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_Pool *PoolFilterer) WatchCreditChainPath(opts *bind.WatchOpts, sink chan<- *PoolCreditChainPath) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "CreditChainPath")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCreditChainPath)
				if err := _Pool.contract.UnpackLog(event, "CreditChainPath", log); err != nil {
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

// ParseCreditChainPath is a log parse operation binding the contract event 0xdbdd25248751feb2f3b66721dfdd11662a68bc155af3771e661aabec92fba814.
//
// Solidity: event CreditChainPath(uint16 chainId, uint256 srcPoolId, uint256 amountSD, uint256 idealBalance)
func (_Pool *PoolFilterer) ParseCreditChainPath(log types.Log) (*PoolCreditChainPath, error) {
	event := new(PoolCreditChainPath)
	if err := _Pool.contract.UnpackLog(event, "CreditChainPath", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolDeltaParamUpdatedIterator is returned from FilterDeltaParamUpdated and is used to iterate over the raw logs and unpacked data for DeltaParamUpdated events raised by the Pool contract.
type PoolDeltaParamUpdatedIterator struct {
	Event *PoolDeltaParamUpdated // Event containing the contract specifics and raw log

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
func (it *PoolDeltaParamUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolDeltaParamUpdated)
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
		it.Event = new(PoolDeltaParamUpdated)
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
func (it *PoolDeltaParamUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolDeltaParamUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolDeltaParamUpdated represents a DeltaParamUpdated event raised by the Pool contract.
type PoolDeltaParamUpdated struct {
	Batched         bool
	SwapDeltaBP     *big.Int
	LpDeltaBP       *big.Int
	DefaultSwapMode bool
	DefaultLPMode   bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeltaParamUpdated is a free log retrieval operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_Pool *PoolFilterer) FilterDeltaParamUpdated(opts *bind.FilterOpts) (*PoolDeltaParamUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "DeltaParamUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolDeltaParamUpdatedIterator{contract: _Pool.contract, event: "DeltaParamUpdated", logs: logs, sub: sub}, nil
}

// WatchDeltaParamUpdated is a free log subscription operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_Pool *PoolFilterer) WatchDeltaParamUpdated(opts *bind.WatchOpts, sink chan<- *PoolDeltaParamUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "DeltaParamUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolDeltaParamUpdated)
				if err := _Pool.contract.UnpackLog(event, "DeltaParamUpdated", log); err != nil {
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

// ParseDeltaParamUpdated is a log parse operation binding the contract event 0x7cc11124872dc29ed41dd447ee7ab07d9eee5d8ebb55f65dd92bce19bb20224a.
//
// Solidity: event DeltaParamUpdated(bool batched, uint256 swapDeltaBP, uint256 lpDeltaBP, bool defaultSwapMode, bool defaultLPMode)
func (_Pool *PoolFilterer) ParseDeltaParamUpdated(log types.Log) (*PoolDeltaParamUpdated, error) {
	event := new(PoolDeltaParamUpdated)
	if err := _Pool.contract.UnpackLog(event, "DeltaParamUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFeeLibraryUpdatedIterator is returned from FilterFeeLibraryUpdated and is used to iterate over the raw logs and unpacked data for FeeLibraryUpdated events raised by the Pool contract.
type PoolFeeLibraryUpdatedIterator struct {
	Event *PoolFeeLibraryUpdated // Event containing the contract specifics and raw log

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
func (it *PoolFeeLibraryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFeeLibraryUpdated)
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
		it.Event = new(PoolFeeLibraryUpdated)
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
func (it *PoolFeeLibraryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFeeLibraryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFeeLibraryUpdated represents a FeeLibraryUpdated event raised by the Pool contract.
type PoolFeeLibraryUpdated struct {
	FeeLibraryAddr common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeLibraryUpdated is a free log retrieval operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_Pool *PoolFilterer) FilterFeeLibraryUpdated(opts *bind.FilterOpts) (*PoolFeeLibraryUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "FeeLibraryUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolFeeLibraryUpdatedIterator{contract: _Pool.contract, event: "FeeLibraryUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeLibraryUpdated is a free log subscription operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_Pool *PoolFilterer) WatchFeeLibraryUpdated(opts *bind.WatchOpts, sink chan<- *PoolFeeLibraryUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "FeeLibraryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFeeLibraryUpdated)
				if err := _Pool.contract.UnpackLog(event, "FeeLibraryUpdated", log); err != nil {
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

// ParseFeeLibraryUpdated is a log parse operation binding the contract event 0x5138b884a20454b6db937b9e11c8534e02e708750e0c465df6cd9701622952ce.
//
// Solidity: event FeeLibraryUpdated(address feeLibraryAddr)
func (_Pool *PoolFilterer) ParseFeeLibraryUpdated(log types.Log) (*PoolFeeLibraryUpdated, error) {
	event := new(PoolFeeLibraryUpdated)
	if err := _Pool.contract.UnpackLog(event, "FeeLibraryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolFeesUpdatedIterator is returned from FilterFeesUpdated and is used to iterate over the raw logs and unpacked data for FeesUpdated events raised by the Pool contract.
type PoolFeesUpdatedIterator struct {
	Event *PoolFeesUpdated // Event containing the contract specifics and raw log

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
func (it *PoolFeesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolFeesUpdated)
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
		it.Event = new(PoolFeesUpdated)
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
func (it *PoolFeesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolFeesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolFeesUpdated represents a FeesUpdated event raised by the Pool contract.
type PoolFeesUpdated struct {
	MintFeeBP *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeesUpdated is a free log retrieval operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_Pool *PoolFilterer) FilterFeesUpdated(opts *bind.FilterOpts) (*PoolFeesUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolFeesUpdatedIterator{contract: _Pool.contract, event: "FeesUpdated", logs: logs, sub: sub}, nil
}

// WatchFeesUpdated is a free log subscription operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_Pool *PoolFilterer) WatchFeesUpdated(opts *bind.WatchOpts, sink chan<- *PoolFeesUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "FeesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolFeesUpdated)
				if err := _Pool.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
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

// ParseFeesUpdated is a log parse operation binding the contract event 0x9fe6eeb0f0541c644a56c67efeb872dbadd803a60b909d7dde1b35a3fe230b0e.
//
// Solidity: event FeesUpdated(uint256 mintFeeBP)
func (_Pool *PoolFilterer) ParseFeesUpdated(log types.Log) (*PoolFeesUpdated, error) {
	event := new(PoolFeesUpdated)
	if err := _Pool.contract.UnpackLog(event, "FeesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolInstantRedeemLocalIterator is returned from FilterInstantRedeemLocal and is used to iterate over the raw logs and unpacked data for InstantRedeemLocal events raised by the Pool contract.
type PoolInstantRedeemLocalIterator struct {
	Event *PoolInstantRedeemLocal // Event containing the contract specifics and raw log

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
func (it *PoolInstantRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolInstantRedeemLocal)
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
		it.Event = new(PoolInstantRedeemLocal)
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
func (it *PoolInstantRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolInstantRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolInstantRedeemLocal represents a InstantRedeemLocal event raised by the Pool contract.
type PoolInstantRedeemLocal struct {
	From     common.Address
	AmountLP *big.Int
	AmountSD *big.Int
	To       common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInstantRedeemLocal is a free log retrieval operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_Pool *PoolFilterer) FilterInstantRedeemLocal(opts *bind.FilterOpts) (*PoolInstantRedeemLocalIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "InstantRedeemLocal")
	if err != nil {
		return nil, err
	}
	return &PoolInstantRedeemLocalIterator{contract: _Pool.contract, event: "InstantRedeemLocal", logs: logs, sub: sub}, nil
}

// WatchInstantRedeemLocal is a free log subscription operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_Pool *PoolFilterer) WatchInstantRedeemLocal(opts *bind.WatchOpts, sink chan<- *PoolInstantRedeemLocal) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "InstantRedeemLocal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolInstantRedeemLocal)
				if err := _Pool.contract.UnpackLog(event, "InstantRedeemLocal", log); err != nil {
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

// ParseInstantRedeemLocal is a log parse operation binding the contract event 0x2125a70154569bd1686edd3cf981bb23dea7c1fa1637909dbb3c9a967cb0c2f2.
//
// Solidity: event InstantRedeemLocal(address from, uint256 amountLP, uint256 amountSD, address to)
func (_Pool *PoolFilterer) ParseInstantRedeemLocal(log types.Log) (*PoolInstantRedeemLocal, error) {
	event := new(PoolInstantRedeemLocal)
	if err := _Pool.contract.UnpackLog(event, "InstantRedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pool contract.
type PoolMintIterator struct {
	Event *PoolMint // Event containing the contract specifics and raw log

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
func (it *PoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMint)
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
		it.Event = new(PoolMint)
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
func (it *PoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMint represents a Mint event raised by the Pool contract.
type PoolMint struct {
	To              common.Address
	AmountLP        *big.Int
	AmountSD        *big.Int
	MintFeeAmountSD *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_Pool *PoolFilterer) FilterMint(opts *bind.FilterOpts) (*PoolMintIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &PoolMintIterator{contract: _Pool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_Pool *PoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PoolMint) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMint)
				if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address to, uint256 amountLP, uint256 amountSD, uint256 mintFeeAmountSD)
func (_Pool *PoolFilterer) ParseMint(log types.Log) (*PoolMint, error) {
	event := new(PoolMint)
	if err := _Pool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRedeemLocalIterator is returned from FilterRedeemLocal and is used to iterate over the raw logs and unpacked data for RedeemLocal events raised by the Pool contract.
type PoolRedeemLocalIterator struct {
	Event *PoolRedeemLocal // Event containing the contract specifics and raw log

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
func (it *PoolRedeemLocalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRedeemLocal)
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
		it.Event = new(PoolRedeemLocal)
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
func (it *PoolRedeemLocalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRedeemLocalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRedeemLocal represents a RedeemLocal event raised by the Pool contract.
type PoolRedeemLocal struct {
	From      common.Address
	AmountLP  *big.Int
	AmountSD  *big.Int
	ChainId   uint16
	DstPoolId *big.Int
	To        []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocal is a free log retrieval operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_Pool *PoolFilterer) FilterRedeemLocal(opts *bind.FilterOpts) (*PoolRedeemLocalIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RedeemLocal")
	if err != nil {
		return nil, err
	}
	return &PoolRedeemLocalIterator{contract: _Pool.contract, event: "RedeemLocal", logs: logs, sub: sub}, nil
}

// WatchRedeemLocal is a free log subscription operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_Pool *PoolFilterer) WatchRedeemLocal(opts *bind.WatchOpts, sink chan<- *PoolRedeemLocal) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RedeemLocal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRedeemLocal)
				if err := _Pool.contract.UnpackLog(event, "RedeemLocal", log); err != nil {
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

// ParseRedeemLocal is a log parse operation binding the contract event 0x53c03ee0722b52efeb42444f48d90173854501b3de3c590fcb445743377115c2.
//
// Solidity: event RedeemLocal(address from, uint256 amountLP, uint256 amountSD, uint16 chainId, uint256 dstPoolId, bytes to)
func (_Pool *PoolFilterer) ParseRedeemLocal(log types.Log) (*PoolRedeemLocal, error) {
	event := new(PoolRedeemLocal)
	if err := _Pool.contract.UnpackLog(event, "RedeemLocal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRedeemLocalCallbackIterator is returned from FilterRedeemLocalCallback and is used to iterate over the raw logs and unpacked data for RedeemLocalCallback events raised by the Pool contract.
type PoolRedeemLocalCallbackIterator struct {
	Event *PoolRedeemLocalCallback // Event containing the contract specifics and raw log

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
func (it *PoolRedeemLocalCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRedeemLocalCallback)
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
		it.Event = new(PoolRedeemLocalCallback)
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
func (it *PoolRedeemLocalCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRedeemLocalCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRedeemLocalCallback represents a RedeemLocalCallback event raised by the Pool contract.
type PoolRedeemLocalCallback struct {
	To             common.Address
	AmountSD       *big.Int
	AmountToMintSD *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRedeemLocalCallback is a free log retrieval operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_Pool *PoolFilterer) FilterRedeemLocalCallback(opts *bind.FilterOpts) (*PoolRedeemLocalCallbackIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RedeemLocalCallback")
	if err != nil {
		return nil, err
	}
	return &PoolRedeemLocalCallbackIterator{contract: _Pool.contract, event: "RedeemLocalCallback", logs: logs, sub: sub}, nil
}

// WatchRedeemLocalCallback is a free log subscription operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_Pool *PoolFilterer) WatchRedeemLocalCallback(opts *bind.WatchOpts, sink chan<- *PoolRedeemLocalCallback) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RedeemLocalCallback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRedeemLocalCallback)
				if err := _Pool.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
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

// ParseRedeemLocalCallback is a log parse operation binding the contract event 0xa97166013ecf5305dd9a58d6d867f05e646d4275f52d2bd52a5c7f00a690ad1b.
//
// Solidity: event RedeemLocalCallback(address _to, uint256 _amountSD, uint256 _amountToMintSD)
func (_Pool *PoolFilterer) ParseRedeemLocalCallback(log types.Log) (*PoolRedeemLocalCallback, error) {
	event := new(PoolRedeemLocalCallback)
	if err := _Pool.contract.UnpackLog(event, "RedeemLocalCallback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRedeemRemoteIterator is returned from FilterRedeemRemote and is used to iterate over the raw logs and unpacked data for RedeemRemote events raised by the Pool contract.
type PoolRedeemRemoteIterator struct {
	Event *PoolRedeemRemote // Event containing the contract specifics and raw log

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
func (it *PoolRedeemRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRedeemRemote)
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
		it.Event = new(PoolRedeemRemote)
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
func (it *PoolRedeemRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRedeemRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRedeemRemote represents a RedeemRemote event raised by the Pool contract.
type PoolRedeemRemote struct {
	ChainId   uint16
	DstPoolId *big.Int
	From      common.Address
	AmountLP  *big.Int
	AmountSD  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemRemote is a free log retrieval operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) FilterRedeemRemote(opts *bind.FilterOpts) (*PoolRedeemRemoteIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RedeemRemote")
	if err != nil {
		return nil, err
	}
	return &PoolRedeemRemoteIterator{contract: _Pool.contract, event: "RedeemRemote", logs: logs, sub: sub}, nil
}

// WatchRedeemRemote is a free log subscription operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) WatchRedeemRemote(opts *bind.WatchOpts, sink chan<- *PoolRedeemRemote) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RedeemRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRedeemRemote)
				if err := _Pool.contract.UnpackLog(event, "RedeemRemote", log); err != nil {
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

// ParseRedeemRemote is a log parse operation binding the contract event 0xa33f5c0b76f00f6737b1780a8a7f18e19c3fe8fe9ee01a6c1b8ce1eae5ed54f9.
//
// Solidity: event RedeemRemote(uint16 chainId, uint256 dstPoolId, address from, uint256 amountLP, uint256 amountSD)
func (_Pool *PoolFilterer) ParseRedeemRemote(log types.Log) (*PoolRedeemRemote, error) {
	event := new(PoolRedeemRemote)
	if err := _Pool.contract.UnpackLog(event, "RedeemRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSendCreditsIterator is returned from FilterSendCredits and is used to iterate over the raw logs and unpacked data for SendCredits events raised by the Pool contract.
type PoolSendCreditsIterator struct {
	Event *PoolSendCredits // Event containing the contract specifics and raw log

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
func (it *PoolSendCreditsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSendCredits)
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
		it.Event = new(PoolSendCredits)
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
func (it *PoolSendCreditsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSendCreditsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSendCredits represents a SendCredits event raised by the Pool contract.
type PoolSendCredits struct {
	DstChainId   uint16
	DstPoolId    *big.Int
	Credits      *big.Int
	IdealBalance *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSendCredits is a free log retrieval operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_Pool *PoolFilterer) FilterSendCredits(opts *bind.FilterOpts) (*PoolSendCreditsIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SendCredits")
	if err != nil {
		return nil, err
	}
	return &PoolSendCreditsIterator{contract: _Pool.contract, event: "SendCredits", logs: logs, sub: sub}, nil
}

// WatchSendCredits is a free log subscription operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_Pool *PoolFilterer) WatchSendCredits(opts *bind.WatchOpts, sink chan<- *PoolSendCredits) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SendCredits")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSendCredits)
				if err := _Pool.contract.UnpackLog(event, "SendCredits", log); err != nil {
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

// ParseSendCredits is a log parse operation binding the contract event 0x6939f93e3f21cf1362eb17155b740277de5687dae9a83a85909fd71da95944e7.
//
// Solidity: event SendCredits(uint16 dstChainId, uint256 dstPoolId, uint256 credits, uint256 idealBalance)
func (_Pool *PoolFilterer) ParseSendCredits(log types.Log) (*PoolSendCredits, error) {
	event := new(PoolSendCredits)
	if err := _Pool.contract.UnpackLog(event, "SendCredits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolStopSwapUpdatedIterator is returned from FilterStopSwapUpdated and is used to iterate over the raw logs and unpacked data for StopSwapUpdated events raised by the Pool contract.
type PoolStopSwapUpdatedIterator struct {
	Event *PoolStopSwapUpdated // Event containing the contract specifics and raw log

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
func (it *PoolStopSwapUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolStopSwapUpdated)
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
		it.Event = new(PoolStopSwapUpdated)
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
func (it *PoolStopSwapUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolStopSwapUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolStopSwapUpdated represents a StopSwapUpdated event raised by the Pool contract.
type PoolStopSwapUpdated struct {
	SwapStop bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopSwapUpdated is a free log retrieval operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_Pool *PoolFilterer) FilterStopSwapUpdated(opts *bind.FilterOpts) (*PoolStopSwapUpdatedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "StopSwapUpdated")
	if err != nil {
		return nil, err
	}
	return &PoolStopSwapUpdatedIterator{contract: _Pool.contract, event: "StopSwapUpdated", logs: logs, sub: sub}, nil
}

// WatchStopSwapUpdated is a free log subscription operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_Pool *PoolFilterer) WatchStopSwapUpdated(opts *bind.WatchOpts, sink chan<- *PoolStopSwapUpdated) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "StopSwapUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolStopSwapUpdated)
				if err := _Pool.contract.UnpackLog(event, "StopSwapUpdated", log); err != nil {
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

// ParseStopSwapUpdated is a log parse operation binding the contract event 0x59a9350977452c5240699f57f18b5915cd0440a56f08820a38b9f2432a82ba3e.
//
// Solidity: event StopSwapUpdated(bool swapStop)
func (_Pool *PoolFilterer) ParseStopSwapUpdated(log types.Log) (*PoolStopSwapUpdated, error) {
	event := new(PoolStopSwapUpdated)
	if err := _Pool.contract.UnpackLog(event, "StopSwapUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pool contract.
type PoolSwapIterator struct {
	Event *PoolSwap // Event containing the contract specifics and raw log

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
func (it *PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwap)
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
		it.Event = new(PoolSwap)
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
func (it *PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwap represents a Swap event raised by the Pool contract.
type PoolSwap struct {
	ChainId     uint16
	DstPoolId   *big.Int
	From        common.Address
	AmountSD    *big.Int
	EqReward    *big.Int
	EqFee       *big.Int
	ProtocolFee *big.Int
	LpFee       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_Pool *PoolFilterer) FilterSwap(opts *bind.FilterOpts) (*PoolSwapIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &PoolSwapIterator{contract: _Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_Pool *PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolSwap) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwap)
				if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x34660fc8af304464529f48a778e03d03e4d34bcd5f9b6f0cfbf3cd238c642f7f.
//
// Solidity: event Swap(uint16 chainId, uint256 dstPoolId, address from, uint256 amountSD, uint256 eqReward, uint256 eqFee, uint256 protocolFee, uint256 lpFee)
func (_Pool *PoolFilterer) ParseSwap(log types.Log) (*PoolSwap, error) {
	event := new(PoolSwap)
	if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapRemoteIterator is returned from FilterSwapRemote and is used to iterate over the raw logs and unpacked data for SwapRemote events raised by the Pool contract.
type PoolSwapRemoteIterator struct {
	Event *PoolSwapRemote // Event containing the contract specifics and raw log

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
func (it *PoolSwapRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwapRemote)
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
		it.Event = new(PoolSwapRemote)
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
func (it *PoolSwapRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwapRemote represents a SwapRemote event raised by the Pool contract.
type PoolSwapRemote struct {
	To          common.Address
	AmountSD    *big.Int
	ProtocolFee *big.Int
	DstFee      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwapRemote is a free log retrieval operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_Pool *PoolFilterer) FilterSwapRemote(opts *bind.FilterOpts) (*PoolSwapRemoteIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SwapRemote")
	if err != nil {
		return nil, err
	}
	return &PoolSwapRemoteIterator{contract: _Pool.contract, event: "SwapRemote", logs: logs, sub: sub}, nil
}

// WatchSwapRemote is a free log subscription operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_Pool *PoolFilterer) WatchSwapRemote(opts *bind.WatchOpts, sink chan<- *PoolSwapRemote) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SwapRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwapRemote)
				if err := _Pool.contract.UnpackLog(event, "SwapRemote", log); err != nil {
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

// ParseSwapRemote is a log parse operation binding the contract event 0xfb2b592367452f1c437675bed47f5e1e6c25188c17d7ba01a12eb030bc41ccef.
//
// Solidity: event SwapRemote(address to, uint256 amountSD, uint256 protocolFee, uint256 dstFee)
func (_Pool *PoolFilterer) ParseSwapRemote(log types.Log) (*PoolSwapRemote, error) {
	event := new(PoolSwapRemote)
	if err := _Pool.contract.UnpackLog(event, "SwapRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pool contract.
type PoolTransferIterator struct {
	Event *PoolTransfer // Event containing the contract specifics and raw log

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
func (it *PoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransfer)
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
		it.Event = new(PoolTransfer)
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
func (it *PoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransfer represents a Transfer event raised by the Pool contract.
type PoolTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolTransferIterator{contract: _Pool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransfer)
				if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pool *PoolFilterer) ParseTransfer(log types.Log) (*PoolTransfer, error) {
	event := new(PoolTransfer)
	if err := _Pool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawMintFeeBalanceIterator is returned from FilterWithdrawMintFeeBalance and is used to iterate over the raw logs and unpacked data for WithdrawMintFeeBalance events raised by the Pool contract.
type PoolWithdrawMintFeeBalanceIterator struct {
	Event *PoolWithdrawMintFeeBalance // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawMintFeeBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawMintFeeBalance)
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
		it.Event = new(PoolWithdrawMintFeeBalance)
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
func (it *PoolWithdrawMintFeeBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawMintFeeBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawMintFeeBalance represents a WithdrawMintFeeBalance event raised by the Pool contract.
type PoolWithdrawMintFeeBalance struct {
	To       common.Address
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawMintFeeBalance is a free log retrieval operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) FilterWithdrawMintFeeBalance(opts *bind.FilterOpts) (*PoolWithdrawMintFeeBalanceIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawMintFeeBalance")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawMintFeeBalanceIterator{contract: _Pool.contract, event: "WithdrawMintFeeBalance", logs: logs, sub: sub}, nil
}

// WatchWithdrawMintFeeBalance is a free log subscription operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) WatchWithdrawMintFeeBalance(opts *bind.WatchOpts, sink chan<- *PoolWithdrawMintFeeBalance) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawMintFeeBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawMintFeeBalance)
				if err := _Pool.contract.UnpackLog(event, "WithdrawMintFeeBalance", log); err != nil {
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

// ParseWithdrawMintFeeBalance is a log parse operation binding the contract event 0x87b3b2749102aa96f2d08396e34cd47673e57148af9cfff965d99bc0378a87dc.
//
// Solidity: event WithdrawMintFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) ParseWithdrawMintFeeBalance(log types.Log) (*PoolWithdrawMintFeeBalance, error) {
	event := new(PoolWithdrawMintFeeBalance)
	if err := _Pool.contract.UnpackLog(event, "WithdrawMintFeeBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawProtocolFeeBalanceIterator is returned from FilterWithdrawProtocolFeeBalance and is used to iterate over the raw logs and unpacked data for WithdrawProtocolFeeBalance events raised by the Pool contract.
type PoolWithdrawProtocolFeeBalanceIterator struct {
	Event *PoolWithdrawProtocolFeeBalance // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawProtocolFeeBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawProtocolFeeBalance)
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
		it.Event = new(PoolWithdrawProtocolFeeBalance)
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
func (it *PoolWithdrawProtocolFeeBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawProtocolFeeBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawProtocolFeeBalance represents a WithdrawProtocolFeeBalance event raised by the Pool contract.
type PoolWithdrawProtocolFeeBalance struct {
	To       common.Address
	AmountSD *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProtocolFeeBalance is a free log retrieval operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) FilterWithdrawProtocolFeeBalance(opts *bind.FilterOpts) (*PoolWithdrawProtocolFeeBalanceIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawProtocolFeeBalance")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawProtocolFeeBalanceIterator{contract: _Pool.contract, event: "WithdrawProtocolFeeBalance", logs: logs, sub: sub}, nil
}

// WatchWithdrawProtocolFeeBalance is a free log subscription operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) WatchWithdrawProtocolFeeBalance(opts *bind.WatchOpts, sink chan<- *PoolWithdrawProtocolFeeBalance) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawProtocolFeeBalance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawProtocolFeeBalance)
				if err := _Pool.contract.UnpackLog(event, "WithdrawProtocolFeeBalance", log); err != nil {
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

// ParseWithdrawProtocolFeeBalance is a log parse operation binding the contract event 0x70dc5a44816033bea80f836440f4b1fe1b3bb06b568c8dc2301901f03bf237c7.
//
// Solidity: event WithdrawProtocolFeeBalance(address to, uint256 amountSD)
func (_Pool *PoolFilterer) ParseWithdrawProtocolFeeBalance(log types.Log) (*PoolWithdrawProtocolFeeBalance, error) {
	event := new(PoolWithdrawProtocolFeeBalance)
	if err := _Pool.contract.UnpackLog(event, "WithdrawProtocolFeeBalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolWithdrawRemoteIterator is returned from FilterWithdrawRemote and is used to iterate over the raw logs and unpacked data for WithdrawRemote events raised by the Pool contract.
type PoolWithdrawRemoteIterator struct {
	Event *PoolWithdrawRemote // Event containing the contract specifics and raw log

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
func (it *PoolWithdrawRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolWithdrawRemote)
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
		it.Event = new(PoolWithdrawRemote)
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
func (it *PoolWithdrawRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolWithdrawRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolWithdrawRemote represents a WithdrawRemote event raised by the Pool contract.
type PoolWithdrawRemote struct {
	SrcChainId uint16
	SrcPoolId  *big.Int
	SwapAmount *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRemote is a free log retrieval operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolFilterer) FilterWithdrawRemote(opts *bind.FilterOpts) (*PoolWithdrawRemoteIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "WithdrawRemote")
	if err != nil {
		return nil, err
	}
	return &PoolWithdrawRemoteIterator{contract: _Pool.contract, event: "WithdrawRemote", logs: logs, sub: sub}, nil
}

// WatchWithdrawRemote is a free log subscription operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolFilterer) WatchWithdrawRemote(opts *bind.WatchOpts, sink chan<- *PoolWithdrawRemote) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "WithdrawRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolWithdrawRemote)
				if err := _Pool.contract.UnpackLog(event, "WithdrawRemote", log); err != nil {
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

// ParseWithdrawRemote is a log parse operation binding the contract event 0x44d3575fd94f9e0a41d7ebbc7e952f9b615c3f8d1faf924e1e9e98c0edf0d380.
//
// Solidity: event WithdrawRemote(uint16 srcChainId, uint256 srcPoolId, uint256 swapAmount, uint256 mintAmount)
func (_Pool *PoolFilterer) ParseWithdrawRemote(log types.Log) (*PoolWithdrawRemote, error) {
	event := new(PoolWithdrawRemote)
	if err := _Pool.contract.UnpackLog(event, "WithdrawRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
