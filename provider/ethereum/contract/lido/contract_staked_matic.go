// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// IStMATICRequestWithdraw is an auto generated low-level Go binding around an user-defined struct.
type IStMATICRequestWithdraw struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}

// StakedMATICMetaData contains all meta data concerning the StakedMATIC contract.
var StakedMATICMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountClaimed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountBurned\",\"type\":\"uint256\"}],\"name\":\"ClaimTokensEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorShare\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amountClaimed\",\"type\":\"uint256\"}],\"name\":\"ClaimTotalDelegatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amountDelegated\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_remainder\",\"type\":\"uint256\"}],\"name\":\"DelegateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"DistributeRewardsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"RequestWithdrawEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldDaoAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDaoAddress\",\"type\":\"address\"}],\"name\":\"SetDaoAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_delegationLowerBound\",\"type\":\"uint256\"}],\"name\":\"SetDelegationLowerBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"daoFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"operatorsFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"insuranceFee\",\"type\":\"uint256\"}],\"name\":\"SetFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldFxStateRootTunnel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"SetFxStateRootTunnel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newInsuranceAddress\",\"type\":\"address\"}],\"name\":\"SetInsuranceAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLidoNFT\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newLidoNFT\",\"type\":\"address\"}],\"name\":\"SetLidoNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newNodeOperatorRegistryAddress\",\"type\":\"address\"}],\"name\":\"SetNodeOperatorRegistryAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"oldProtocolFee\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"newProtocolFee\",\"type\":\"uint8\"}],\"name\":\"SetProtocolFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldRewardDistributionLowerBound\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRewardDistributionLowerBound\",\"type\":\"uint256\"}],\"name\":\"SetRewardDistributionLowerBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"SubmitEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldVersion\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"}],\"name\":\"Version\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawTotalDelegatedEvent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAO\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calculatePendingBufferedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingBufferedTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"claimTokensFromValidatorToContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInMatic\",\"type\":\"uint256\"}],\"name\":\"convertMaticToStMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInStMatic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStMaticSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPooledMatic\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInStMatic\",\"type\":\"uint256\"}],\"name\":\"convertStMaticToMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountInMatic\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStMaticAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPooledMatic\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dao\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegationLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entityFees\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"dao\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"operators\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"insurance\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fxStateRootTunnel\",\"outputs\":[{\"internalType\":\"contractIFxStateRootTunnel\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getLiquidRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getMaticFromTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getToken2WithdrawRequests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalPooledMatic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIValidatorShare\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"getTotalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalStakeAcrossAllValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalWithdrawRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"internalType\":\"structIStMATIC.RequestWithdraw[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nodeOperatorRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dao\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_insurance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_poLidoNFT\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_fxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"insurance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastWithdrawnValidatorId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeOperatorRegistry\",\"outputs\":[{\"internalType\":\"contractINodeOperatorRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poLidoNFT\",\"outputs\":[{\"internalType\":\"contractIPoLidoNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebalanceDelegatedTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"requestWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardDistributionLowerBound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDAO\",\"type\":\"address\"}],\"name\":\"setDaoAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_delegationLowerBound\",\"type\":\"uint256\"}],\"name\":\"setDelegationLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_daoFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_operatorsFee\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_insuranceFee\",\"type\":\"uint8\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newFxStateRootTunnel\",\"type\":\"address\"}],\"name\":\"setFxStateRootTunnel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setInsuranceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setNodeOperatorRegistryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newLidoNFT\",\"type\":\"address\"}],\"name\":\"setPoLidoNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_newProtocolFee\",\"type\":\"uint8\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRewardDistributionLowerBound\",\"type\":\"uint256\"}],\"name\":\"setRewardDistributionLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newVersion\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stMaticWithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeManager\",\"outputs\":[{\"internalType\":\"contractIStakeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitHandler\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"token2WithdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount2WithdrawFromStMATIC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestEpoch\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuffered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_validatorShare\",\"type\":\"address\"}],\"name\":\"withdrawTotalDelegated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakedMATICABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedMATICMetaData.ABI instead.
var StakedMATICABI = StakedMATICMetaData.ABI

// StakedMATIC is an auto generated Go binding around an Ethereum contract.
type StakedMATIC struct {
	StakedMATICCaller     // Read-only binding to the contract
	StakedMATICTransactor // Write-only binding to the contract
	StakedMATICFilterer   // Log filterer for contract events
}

// StakedMATICCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedMATICCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedMATICTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedMATICTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedMATICFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedMATICFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedMATICSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedMATICSession struct {
	Contract     *StakedMATIC      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedMATICCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedMATICCallerSession struct {
	Contract *StakedMATICCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StakedMATICTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedMATICTransactorSession struct {
	Contract     *StakedMATICTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StakedMATICRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedMATICRaw struct {
	Contract *StakedMATIC // Generic contract binding to access the raw methods on
}

// StakedMATICCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedMATICCallerRaw struct {
	Contract *StakedMATICCaller // Generic read-only contract binding to access the raw methods on
}

// StakedMATICTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedMATICTransactorRaw struct {
	Contract *StakedMATICTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedMATIC creates a new instance of StakedMATIC, bound to a specific deployed contract.
func NewStakedMATIC(address common.Address, backend bind.ContractBackend) (*StakedMATIC, error) {
	contract, err := bindStakedMATIC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedMATIC{StakedMATICCaller: StakedMATICCaller{contract: contract}, StakedMATICTransactor: StakedMATICTransactor{contract: contract}, StakedMATICFilterer: StakedMATICFilterer{contract: contract}}, nil
}

// NewStakedMATICCaller creates a new read-only instance of StakedMATIC, bound to a specific deployed contract.
func NewStakedMATICCaller(address common.Address, caller bind.ContractCaller) (*StakedMATICCaller, error) {
	contract, err := bindStakedMATIC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedMATICCaller{contract: contract}, nil
}

// NewStakedMATICTransactor creates a new write-only instance of StakedMATIC, bound to a specific deployed contract.
func NewStakedMATICTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedMATICTransactor, error) {
	contract, err := bindStakedMATIC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedMATICTransactor{contract: contract}, nil
}

// NewStakedMATICFilterer creates a new log filterer instance of StakedMATIC, bound to a specific deployed contract.
func NewStakedMATICFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedMATICFilterer, error) {
	contract, err := bindStakedMATIC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedMATICFilterer{contract: contract}, nil
}

// bindStakedMATIC binds a generic wrapper to an already deployed contract.
func bindStakedMATIC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedMATICMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedMATIC *StakedMATICRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedMATIC.Contract.StakedMATICCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedMATIC *StakedMATICRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.Contract.StakedMATICTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedMATIC *StakedMATICRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedMATIC.Contract.StakedMATICTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedMATIC *StakedMATICCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedMATIC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedMATIC *StakedMATICTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedMATIC *StakedMATICTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedMATIC.Contract.contract.Transact(opts, method, params...)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StakedMATIC *StakedMATICCaller) DAO(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "DAO")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StakedMATIC *StakedMATICSession) DAO() ([32]byte, error) {
	return _StakedMATIC.Contract.DAO(&_StakedMATIC.CallOpts)
}

// DAO is a free data retrieval call binding the contract method 0x98fabd3a.
//
// Solidity: function DAO() view returns(bytes32)
func (_StakedMATIC *StakedMATICCallerSession) DAO() ([32]byte, error) {
	return _StakedMATIC.Contract.DAO(&_StakedMATIC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.DEFAULTADMINROLE(&_StakedMATIC.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.DEFAULTADMINROLE(&_StakedMATIC.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICSession) PAUSEROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.PAUSEROLE(&_StakedMATIC.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.PAUSEROLE(&_StakedMATIC.CallOpts)
}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCaller) UNPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "UNPAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICSession) UNPAUSEROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.UNPAUSEROLE(&_StakedMATIC.CallOpts)
}

// UNPAUSEROLE is a free data retrieval call binding the contract method 0x309756fb.
//
// Solidity: function UNPAUSE_ROLE() view returns(bytes32)
func (_StakedMATIC *StakedMATICCallerSession) UNPAUSEROLE() ([32]byte, error) {
	return _StakedMATIC.Contract.UNPAUSEROLE(&_StakedMATIC.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedMATIC *StakedMATICSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.Allowance(&_StakedMATIC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.Allowance(&_StakedMATIC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedMATIC *StakedMATICSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.BalanceOf(&_StakedMATIC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.BalanceOf(&_StakedMATIC.CallOpts, account)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StakedMATIC *StakedMATICCaller) CalculatePendingBufferedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "calculatePendingBufferedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StakedMATIC *StakedMATICSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _StakedMATIC.Contract.CalculatePendingBufferedTokens(&_StakedMATIC.CallOpts)
}

// CalculatePendingBufferedTokens is a free data retrieval call binding the contract method 0xafd290a7.
//
// Solidity: function calculatePendingBufferedTokens() view returns(uint256 pendingBufferedTokens)
func (_StakedMATIC *StakedMATICCallerSession) CalculatePendingBufferedTokens() (*big.Int, error) {
	return _StakedMATIC.Contract.CalculatePendingBufferedTokens(&_StakedMATIC.CallOpts)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICCaller) ConvertMaticToStMatic(opts *bind.CallOpts, _amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "convertMaticToStMatic", _amountInMatic)

	outstruct := new(struct {
		AmountInStMatic    *big.Int
		TotalStMaticSupply *big.Int
		TotalPooledMatic   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInStMatic = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStMaticSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPooledMatic = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICSession) ConvertMaticToStMatic(_amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StakedMATIC.Contract.ConvertMaticToStMatic(&_StakedMATIC.CallOpts, _amountInMatic)
}

// ConvertMaticToStMatic is a free data retrieval call binding the contract method 0x917a52f5.
//
// Solidity: function convertMaticToStMatic(uint256 _amountInMatic) view returns(uint256 amountInStMatic, uint256 totalStMaticSupply, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICCallerSession) ConvertMaticToStMatic(_amountInMatic *big.Int) (struct {
	AmountInStMatic    *big.Int
	TotalStMaticSupply *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StakedMATIC.Contract.ConvertMaticToStMatic(&_StakedMATIC.CallOpts, _amountInMatic)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICCaller) ConvertStMaticToMatic(opts *bind.CallOpts, _amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "convertStMaticToMatic", _amountInStMatic)

	outstruct := new(struct {
		AmountInMatic      *big.Int
		TotalStMaticAmount *big.Int
		TotalPooledMatic   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInMatic = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStMaticAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalPooledMatic = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICSession) ConvertStMaticToMatic(_amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StakedMATIC.Contract.ConvertStMaticToMatic(&_StakedMATIC.CallOpts, _amountInStMatic)
}

// ConvertStMaticToMatic is a free data retrieval call binding the contract method 0xd968447c.
//
// Solidity: function convertStMaticToMatic(uint256 _amountInStMatic) view returns(uint256 amountInMatic, uint256 totalStMaticAmount, uint256 totalPooledMatic)
func (_StakedMATIC *StakedMATICCallerSession) ConvertStMaticToMatic(_amountInStMatic *big.Int) (struct {
	AmountInMatic      *big.Int
	TotalStMaticAmount *big.Int
	TotalPooledMatic   *big.Int
}, error) {
	return _StakedMATIC.Contract.ConvertStMaticToMatic(&_StakedMATIC.CallOpts, _amountInStMatic)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StakedMATIC *StakedMATICCaller) Dao(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "dao")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StakedMATIC *StakedMATICSession) Dao() (common.Address, error) {
	return _StakedMATIC.Contract.Dao(&_StakedMATIC.CallOpts)
}

// Dao is a free data retrieval call binding the contract method 0x4162169f.
//
// Solidity: function dao() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) Dao() (common.Address, error) {
	return _StakedMATIC.Contract.Dao(&_StakedMATIC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedMATIC *StakedMATICCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedMATIC *StakedMATICSession) Decimals() (uint8, error) {
	return _StakedMATIC.Contract.Decimals(&_StakedMATIC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StakedMATIC *StakedMATICCallerSession) Decimals() (uint8, error) {
	return _StakedMATIC.Contract.Decimals(&_StakedMATIC.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) DelegationLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "delegationLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) DelegationLowerBound() (*big.Int, error) {
	return _StakedMATIC.Contract.DelegationLowerBound(&_StakedMATIC.CallOpts)
}

// DelegationLowerBound is a free data retrieval call binding the contract method 0x0d7abc33.
//
// Solidity: function delegationLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) DelegationLowerBound() (*big.Int, error) {
	return _StakedMATIC.Contract.DelegationLowerBound(&_StakedMATIC.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StakedMATIC *StakedMATICCaller) EntityFees(opts *bind.CallOpts) (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "entityFees")

	outstruct := new(struct {
		Dao       uint8
		Operators uint8
		Insurance uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Dao = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Operators = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Insurance = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StakedMATIC *StakedMATICSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _StakedMATIC.Contract.EntityFees(&_StakedMATIC.CallOpts)
}

// EntityFees is a free data retrieval call binding the contract method 0x964a7596.
//
// Solidity: function entityFees() view returns(uint8 dao, uint8 operators, uint8 insurance)
func (_StakedMATIC *StakedMATICCallerSession) EntityFees() (struct {
	Dao       uint8
	Operators uint8
	Insurance uint8
}, error) {
	return _StakedMATIC.Contract.EntityFees(&_StakedMATIC.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StakedMATIC *StakedMATICCaller) FxStateRootTunnel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "fxStateRootTunnel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StakedMATIC *StakedMATICSession) FxStateRootTunnel() (common.Address, error) {
	return _StakedMATIC.Contract.FxStateRootTunnel(&_StakedMATIC.CallOpts)
}

// FxStateRootTunnel is a free data retrieval call binding the contract method 0xe062b10b.
//
// Solidity: function fxStateRootTunnel() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) FxStateRootTunnel() (common.Address, error) {
	return _StakedMATIC.Contract.FxStateRootTunnel(&_StakedMATIC.CallOpts)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) GetLiquidRewards(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getLiquidRewards", _validatorShare)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StakedMATIC *StakedMATICSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.GetLiquidRewards(&_StakedMATIC.CallOpts, _validatorShare)
}

// GetLiquidRewards is a free data retrieval call binding the contract method 0x676e5550.
//
// Solidity: function getLiquidRewards(address _validatorShare) view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) GetLiquidRewards(_validatorShare common.Address) (*big.Int, error) {
	return _StakedMATIC.Contract.GetLiquidRewards(&_StakedMATIC.CallOpts, _validatorShare)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) GetMaticFromTokenId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getMaticFromTokenId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StakedMATIC *StakedMATICSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _StakedMATIC.Contract.GetMaticFromTokenId(&_StakedMATIC.CallOpts, _tokenId)
}

// GetMaticFromTokenId is a free data retrieval call binding the contract method 0x720bcf1d.
//
// Solidity: function getMaticFromTokenId(uint256 _tokenId) view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) GetMaticFromTokenId(_tokenId *big.Int) (*big.Int, error) {
	return _StakedMATIC.Contract.GetMaticFromTokenId(&_StakedMATIC.CallOpts, _tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedMATIC *StakedMATICCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedMATIC *StakedMATICSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedMATIC.Contract.GetRoleAdmin(&_StakedMATIC.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedMATIC *StakedMATICCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedMATIC.Contract.GetRoleAdmin(&_StakedMATIC.CallOpts, role)
}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICCaller) GetToken2WithdrawRequests(opts *bind.CallOpts, _tokenId *big.Int) ([]IStMATICRequestWithdraw, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getToken2WithdrawRequests", _tokenId)

	if err != nil {
		return *new([]IStMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStMATICRequestWithdraw)).(*[]IStMATICRequestWithdraw)

	return out0, err

}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICSession) GetToken2WithdrawRequests(_tokenId *big.Int) ([]IStMATICRequestWithdraw, error) {
	return _StakedMATIC.Contract.GetToken2WithdrawRequests(&_StakedMATIC.CallOpts, _tokenId)
}

// GetToken2WithdrawRequests is a free data retrieval call binding the contract method 0x253d1735.
//
// Solidity: function getToken2WithdrawRequests(uint256 _tokenId) view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICCallerSession) GetToken2WithdrawRequests(_tokenId *big.Int) ([]IStMATICRequestWithdraw, error) {
	return _StakedMATIC.Contract.GetToken2WithdrawRequests(&_StakedMATIC.CallOpts, _tokenId)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) GetTotalPooledMatic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getTotalPooledMatic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StakedMATIC.Contract.GetTotalPooledMatic(&_StakedMATIC.CallOpts)
}

// GetTotalPooledMatic is a free data retrieval call binding the contract method 0xe00222a0.
//
// Solidity: function getTotalPooledMatic() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) GetTotalPooledMatic() (*big.Int, error) {
	return _StakedMATIC.Contract.GetTotalPooledMatic(&_StakedMATIC.CallOpts)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StakedMATIC *StakedMATICCaller) GetTotalStake(opts *bind.CallOpts, _validatorShare common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getTotalStake", _validatorShare)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StakedMATIC *StakedMATICSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StakedMATIC.Contract.GetTotalStake(&_StakedMATIC.CallOpts, _validatorShare)
}

// GetTotalStake is a free data retrieval call binding the contract method 0x1e7ff8f6.
//
// Solidity: function getTotalStake(address _validatorShare) view returns(uint256, uint256)
func (_StakedMATIC *StakedMATICCallerSession) GetTotalStake(_validatorShare common.Address) (*big.Int, *big.Int, error) {
	return _StakedMATIC.Contract.GetTotalStake(&_StakedMATIC.CallOpts, _validatorShare)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) GetTotalStakeAcrossAllValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getTotalStakeAcrossAllValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StakedMATIC.Contract.GetTotalStakeAcrossAllValidators(&_StakedMATIC.CallOpts)
}

// GetTotalStakeAcrossAllValidators is a free data retrieval call binding the contract method 0x7e978af8.
//
// Solidity: function getTotalStakeAcrossAllValidators() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) GetTotalStakeAcrossAllValidators() (*big.Int, error) {
	return _StakedMATIC.Contract.GetTotalStakeAcrossAllValidators(&_StakedMATIC.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICCaller) GetTotalWithdrawRequest(opts *bind.CallOpts) ([]IStMATICRequestWithdraw, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "getTotalWithdrawRequest")

	if err != nil {
		return *new([]IStMATICRequestWithdraw), err
	}

	out0 := *abi.ConvertType(out[0], new([]IStMATICRequestWithdraw)).(*[]IStMATICRequestWithdraw)

	return out0, err

}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _StakedMATIC.Contract.GetTotalWithdrawRequest(&_StakedMATIC.CallOpts)
}

// GetTotalWithdrawRequest is a free data retrieval call binding the contract method 0x916b9eba.
//
// Solidity: function getTotalWithdrawRequest() view returns((uint256,uint256,uint256,address)[])
func (_StakedMATIC *StakedMATICCallerSession) GetTotalWithdrawRequest() ([]IStMATICRequestWithdraw, error) {
	return _StakedMATIC.Contract.GetTotalWithdrawRequest(&_StakedMATIC.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedMATIC *StakedMATICCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedMATIC *StakedMATICSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedMATIC.Contract.HasRole(&_StakedMATIC.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedMATIC *StakedMATICCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedMATIC.Contract.HasRole(&_StakedMATIC.CallOpts, role, account)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StakedMATIC *StakedMATICCaller) Insurance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "insurance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StakedMATIC *StakedMATICSession) Insurance() (common.Address, error) {
	return _StakedMATIC.Contract.Insurance(&_StakedMATIC.CallOpts)
}

// Insurance is a free data retrieval call binding the contract method 0x89cf3204.
//
// Solidity: function insurance() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) Insurance() (common.Address, error) {
	return _StakedMATIC.Contract.Insurance(&_StakedMATIC.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) LastWithdrawnValidatorId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "lastWithdrawnValidatorId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _StakedMATIC.Contract.LastWithdrawnValidatorId(&_StakedMATIC.CallOpts)
}

// LastWithdrawnValidatorId is a free data retrieval call binding the contract method 0x71975a3e.
//
// Solidity: function lastWithdrawnValidatorId() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) LastWithdrawnValidatorId() (*big.Int, error) {
	return _StakedMATIC.Contract.LastWithdrawnValidatorId(&_StakedMATIC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedMATIC *StakedMATICCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedMATIC *StakedMATICSession) Name() (string, error) {
	return _StakedMATIC.Contract.Name(&_StakedMATIC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedMATIC *StakedMATICCallerSession) Name() (string, error) {
	return _StakedMATIC.Contract.Name(&_StakedMATIC.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StakedMATIC *StakedMATICCaller) NodeOperatorRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "nodeOperatorRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StakedMATIC *StakedMATICSession) NodeOperatorRegistry() (common.Address, error) {
	return _StakedMATIC.Contract.NodeOperatorRegistry(&_StakedMATIC.CallOpts)
}

// NodeOperatorRegistry is a free data retrieval call binding the contract method 0xe8f8708f.
//
// Solidity: function nodeOperatorRegistry() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) NodeOperatorRegistry() (common.Address, error) {
	return _StakedMATIC.Contract.NodeOperatorRegistry(&_StakedMATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakedMATIC *StakedMATICCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakedMATIC *StakedMATICSession) Paused() (bool, error) {
	return _StakedMATIC.Contract.Paused(&_StakedMATIC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakedMATIC *StakedMATICCallerSession) Paused() (bool, error) {
	return _StakedMATIC.Contract.Paused(&_StakedMATIC.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StakedMATIC *StakedMATICCaller) PoLidoNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "poLidoNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StakedMATIC *StakedMATICSession) PoLidoNFT() (common.Address, error) {
	return _StakedMATIC.Contract.PoLidoNFT(&_StakedMATIC.CallOpts)
}

// PoLidoNFT is a free data retrieval call binding the contract method 0x7029c90e.
//
// Solidity: function poLidoNFT() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) PoLidoNFT() (common.Address, error) {
	return _StakedMATIC.Contract.PoLidoNFT(&_StakedMATIC.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StakedMATIC *StakedMATICCaller) ProtocolFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StakedMATIC *StakedMATICSession) ProtocolFee() (uint8, error) {
	return _StakedMATIC.Contract.ProtocolFee(&_StakedMATIC.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint8)
func (_StakedMATIC *StakedMATICCallerSession) ProtocolFee() (uint8, error) {
	return _StakedMATIC.Contract.ProtocolFee(&_StakedMATIC.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) ReservedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "reservedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) ReservedFunds() (*big.Int, error) {
	return _StakedMATIC.Contract.ReservedFunds(&_StakedMATIC.CallOpts)
}

// ReservedFunds is a free data retrieval call binding the contract method 0x509c5df6.
//
// Solidity: function reservedFunds() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) ReservedFunds() (*big.Int, error) {
	return _StakedMATIC.Contract.ReservedFunds(&_StakedMATIC.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) RewardDistributionLowerBound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "rewardDistributionLowerBound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _StakedMATIC.Contract.RewardDistributionLowerBound(&_StakedMATIC.CallOpts)
}

// RewardDistributionLowerBound is a free data retrieval call binding the contract method 0xa2452947.
//
// Solidity: function rewardDistributionLowerBound() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) RewardDistributionLowerBound() (*big.Int, error) {
	return _StakedMATIC.Contract.RewardDistributionLowerBound(&_StakedMATIC.CallOpts)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCaller) StMaticWithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "stMaticWithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.StMaticWithdrawRequest(&_StakedMATIC.CallOpts, arg0)
}

// StMaticWithdrawRequest is a free data retrieval call binding the contract method 0xf1a13fce.
//
// Solidity: function stMaticWithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCallerSession) StMaticWithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.StMaticWithdrawRequest(&_StakedMATIC.CallOpts, arg0)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StakedMATIC *StakedMATICCaller) StakeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "stakeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StakedMATIC *StakedMATICSession) StakeManager() (common.Address, error) {
	return _StakedMATIC.Contract.StakeManager(&_StakedMATIC.CallOpts)
}

// StakeManager is a free data retrieval call binding the contract method 0x7542ff95.
//
// Solidity: function stakeManager() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) StakeManager() (common.Address, error) {
	return _StakedMATIC.Contract.StakeManager(&_StakedMATIC.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StakedMATIC *StakedMATICCaller) SubmitHandler(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "submitHandler")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StakedMATIC *StakedMATICSession) SubmitHandler() (bool, error) {
	return _StakedMATIC.Contract.SubmitHandler(&_StakedMATIC.CallOpts)
}

// SubmitHandler is a free data retrieval call binding the contract method 0xe259faf7.
//
// Solidity: function submitHandler() view returns(bool)
func (_StakedMATIC *StakedMATICCallerSession) SubmitHandler() (bool, error) {
	return _StakedMATIC.Contract.SubmitHandler(&_StakedMATIC.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) SubmitThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "submitThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) SubmitThreshold() (*big.Int, error) {
	return _StakedMATIC.Contract.SubmitThreshold(&_StakedMATIC.CallOpts)
}

// SubmitThreshold is a free data retrieval call binding the contract method 0x893818a3.
//
// Solidity: function submitThreshold() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) SubmitThreshold() (*big.Int, error) {
	return _StakedMATIC.Contract.SubmitThreshold(&_StakedMATIC.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedMATIC *StakedMATICCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedMATIC *StakedMATICSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedMATIC.Contract.SupportsInterface(&_StakedMATIC.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedMATIC *StakedMATICCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedMATIC.Contract.SupportsInterface(&_StakedMATIC.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedMATIC *StakedMATICCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedMATIC *StakedMATICSession) Symbol() (string, error) {
	return _StakedMATIC.Contract.Symbol(&_StakedMATIC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedMATIC *StakedMATICCallerSession) Symbol() (string, error) {
	return _StakedMATIC.Contract.Symbol(&_StakedMATIC.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakedMATIC *StakedMATICCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakedMATIC *StakedMATICSession) Token() (common.Address, error) {
	return _StakedMATIC.Contract.Token(&_StakedMATIC.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_StakedMATIC *StakedMATICCallerSession) Token() (common.Address, error) {
	return _StakedMATIC.Contract.Token(&_StakedMATIC.CallOpts)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCaller) Token2WithdrawRequest(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "token2WithdrawRequest", arg0)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.Token2WithdrawRequest(&_StakedMATIC.CallOpts, arg0)
}

// Token2WithdrawRequest is a free data retrieval call binding the contract method 0xf08711fe.
//
// Solidity: function token2WithdrawRequest(uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCallerSession) Token2WithdrawRequest(arg0 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.Token2WithdrawRequest(&_StakedMATIC.CallOpts, arg0)
}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCaller) Token2WithdrawRequests(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "token2WithdrawRequests", arg0, arg1)

	outstruct := new(struct {
		Amount2WithdrawFromStMATIC *big.Int
		ValidatorNonce             *big.Int
		RequestEpoch               *big.Int
		ValidatorAddress           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount2WithdrawFromStMATIC = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ValidatorNonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RequestEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValidatorAddress = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICSession) Token2WithdrawRequests(arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.Token2WithdrawRequests(&_StakedMATIC.CallOpts, arg0, arg1)
}

// Token2WithdrawRequests is a free data retrieval call binding the contract method 0xc697d2c7.
//
// Solidity: function token2WithdrawRequests(uint256 , uint256 ) view returns(uint256 amount2WithdrawFromStMATIC, uint256 validatorNonce, uint256 requestEpoch, address validatorAddress)
func (_StakedMATIC *StakedMATICCallerSession) Token2WithdrawRequests(arg0 *big.Int, arg1 *big.Int) (struct {
	Amount2WithdrawFromStMATIC *big.Int
	ValidatorNonce             *big.Int
	RequestEpoch               *big.Int
	ValidatorAddress           common.Address
}, error) {
	return _StakedMATIC.Contract.Token2WithdrawRequests(&_StakedMATIC.CallOpts, arg0, arg1)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) TotalBuffered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "totalBuffered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) TotalBuffered() (*big.Int, error) {
	return _StakedMATIC.Contract.TotalBuffered(&_StakedMATIC.CallOpts)
}

// TotalBuffered is a free data retrieval call binding the contract method 0x52349b17.
//
// Solidity: function totalBuffered() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) TotalBuffered() (*big.Int, error) {
	return _StakedMATIC.Contract.TotalBuffered(&_StakedMATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedMATIC *StakedMATICCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedMATIC *StakedMATICSession) TotalSupply() (*big.Int, error) {
	return _StakedMATIC.Contract.TotalSupply(&_StakedMATIC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedMATIC *StakedMATICCallerSession) TotalSupply() (*big.Int, error) {
	return _StakedMATIC.Contract.TotalSupply(&_StakedMATIC.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StakedMATIC *StakedMATICCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedMATIC.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StakedMATIC *StakedMATICSession) Version() (string, error) {
	return _StakedMATIC.Contract.Version(&_StakedMATIC.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_StakedMATIC *StakedMATICCallerSession) Version() (string, error) {
	return _StakedMATIC.Contract.Version(&_StakedMATIC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Approve(&_StakedMATIC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Approve(&_StakedMATIC.TransactOpts, spender, amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StakedMATIC *StakedMATICTransactor) ClaimTokens(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "claimTokens", _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StakedMATIC *StakedMATICSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.ClaimTokens(&_StakedMATIC.TransactOpts, _tokenId)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x46e04a2f.
//
// Solidity: function claimTokens(uint256 _tokenId) returns()
func (_StakedMATIC *StakedMATICTransactorSession) ClaimTokens(_tokenId *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.ClaimTokens(&_StakedMATIC.TransactOpts, _tokenId)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StakedMATIC *StakedMATICTransactor) ClaimTokensFromValidatorToContract(opts *bind.TransactOpts, _index *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "claimTokensFromValidatorToContract", _index)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StakedMATIC *StakedMATICSession) ClaimTokensFromValidatorToContract(_index *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.ClaimTokensFromValidatorToContract(&_StakedMATIC.TransactOpts, _index)
}

// ClaimTokensFromValidatorToContract is a paid mutator transaction binding the contract method 0x4cfeb862.
//
// Solidity: function claimTokensFromValidatorToContract(uint256 _index) returns()
func (_StakedMATIC *StakedMATICTransactorSession) ClaimTokensFromValidatorToContract(_index *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.ClaimTokensFromValidatorToContract(&_StakedMATIC.TransactOpts, _index)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedMATIC *StakedMATICTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedMATIC *StakedMATICSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.DecreaseAllowance(&_StakedMATIC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StakedMATIC *StakedMATICTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.DecreaseAllowance(&_StakedMATIC.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StakedMATIC *StakedMATICTransactor) Delegate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "delegate")
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StakedMATIC *StakedMATICSession) Delegate() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Delegate(&_StakedMATIC.TransactOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0xc89e4361.
//
// Solidity: function delegate() returns()
func (_StakedMATIC *StakedMATICTransactorSession) Delegate() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Delegate(&_StakedMATIC.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StakedMATIC *StakedMATICTransactor) DistributeRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "distributeRewards")
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StakedMATIC *StakedMATICSession) DistributeRewards() (*types.Transaction, error) {
	return _StakedMATIC.Contract.DistributeRewards(&_StakedMATIC.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x6f4a2cd0.
//
// Solidity: function distributeRewards() returns()
func (_StakedMATIC *StakedMATICTransactorSession) DistributeRewards() (*types.Transaction, error) {
	return _StakedMATIC.Contract.DistributeRewards(&_StakedMATIC.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.GrantRole(&_StakedMATIC.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.GrantRole(&_StakedMATIC.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedMATIC *StakedMATICTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedMATIC *StakedMATICSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.IncreaseAllowance(&_StakedMATIC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StakedMATIC *StakedMATICTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.IncreaseAllowance(&_StakedMATIC.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICTransactor) Initialize(opts *bind.TransactOpts, _nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "initialize", _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Initialize(&_StakedMATIC.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address _nodeOperatorRegistry, address _token, address _dao, address _insurance, address _stakeManager, address _poLidoNFT, address _fxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICTransactorSession) Initialize(_nodeOperatorRegistry common.Address, _token common.Address, _dao common.Address, _insurance common.Address, _stakeManager common.Address, _poLidoNFT common.Address, _fxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Initialize(&_StakedMATIC.TransactOpts, _nodeOperatorRegistry, _token, _dao, _insurance, _stakeManager, _poLidoNFT, _fxStateRootTunnel)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakedMATIC *StakedMATICTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakedMATIC *StakedMATICSession) Pause() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Pause(&_StakedMATIC.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakedMATIC *StakedMATICTransactorSession) Pause() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Pause(&_StakedMATIC.TransactOpts)
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StakedMATIC *StakedMATICTransactor) RebalanceDelegatedTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "rebalanceDelegatedTokens")
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StakedMATIC *StakedMATICSession) RebalanceDelegatedTokens() (*types.Transaction, error) {
	return _StakedMATIC.Contract.RebalanceDelegatedTokens(&_StakedMATIC.TransactOpts)
}

// RebalanceDelegatedTokens is a paid mutator transaction binding the contract method 0xd280f14f.
//
// Solidity: function rebalanceDelegatedTokens() returns()
func (_StakedMATIC *StakedMATICTransactorSession) RebalanceDelegatedTokens() (*types.Transaction, error) {
	return _StakedMATIC.Contract.RebalanceDelegatedTokens(&_StakedMATIC.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RenounceRole(&_StakedMATIC.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RenounceRole(&_StakedMATIC.TransactOpts, role, account)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICTransactor) RequestWithdraw(opts *bind.TransactOpts, _amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "requestWithdraw", _amount, _referral)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICSession) RequestWithdraw(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RequestWithdraw(&_StakedMATIC.TransactOpts, _amount, _referral)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xccc143b8.
//
// Solidity: function requestWithdraw(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICTransactorSession) RequestWithdraw(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RequestWithdraw(&_StakedMATIC.TransactOpts, _amount, _referral)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RevokeRole(&_StakedMATIC.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedMATIC *StakedMATICTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.RevokeRole(&_StakedMATIC.TransactOpts, role, account)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StakedMATIC *StakedMATICTransactor) SetDaoAddress(opts *bind.TransactOpts, _newDAO common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setDaoAddress", _newDAO)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StakedMATIC *StakedMATICSession) SetDaoAddress(_newDAO common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetDaoAddress(&_StakedMATIC.TransactOpts, _newDAO)
}

// SetDaoAddress is a paid mutator transaction binding the contract method 0x9a3cac6a.
//
// Solidity: function setDaoAddress(address _newDAO) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetDaoAddress(_newDAO common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetDaoAddress(&_StakedMATIC.TransactOpts, _newDAO)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StakedMATIC *StakedMATICTransactor) SetDelegationLowerBound(opts *bind.TransactOpts, _delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setDelegationLowerBound", _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StakedMATIC *StakedMATICSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetDelegationLowerBound(&_StakedMATIC.TransactOpts, _delegationLowerBound)
}

// SetDelegationLowerBound is a paid mutator transaction binding the contract method 0x7682c902.
//
// Solidity: function setDelegationLowerBound(uint256 _delegationLowerBound) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetDelegationLowerBound(_delegationLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetDelegationLowerBound(&_StakedMATIC.TransactOpts, _delegationLowerBound)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StakedMATIC *StakedMATICTransactor) SetFees(opts *bind.TransactOpts, _daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setFees", _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StakedMATIC *StakedMATICSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetFees(&_StakedMATIC.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFees is a paid mutator transaction binding the contract method 0xf6794fdb.
//
// Solidity: function setFees(uint8 _daoFee, uint8 _operatorsFee, uint8 _insuranceFee) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetFees(_daoFee uint8, _operatorsFee uint8, _insuranceFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetFees(&_StakedMATIC.TransactOpts, _daoFee, _operatorsFee, _insuranceFee)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICTransactor) SetFxStateRootTunnel(opts *bind.TransactOpts, _newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setFxStateRootTunnel", _newFxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICSession) SetFxStateRootTunnel(_newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetFxStateRootTunnel(&_StakedMATIC.TransactOpts, _newFxStateRootTunnel)
}

// SetFxStateRootTunnel is a paid mutator transaction binding the contract method 0x70bf9fe9.
//
// Solidity: function setFxStateRootTunnel(address _newFxStateRootTunnel) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetFxStateRootTunnel(_newFxStateRootTunnel common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetFxStateRootTunnel(&_StakedMATIC.TransactOpts, _newFxStateRootTunnel)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StakedMATIC *StakedMATICTransactor) SetInsuranceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setInsuranceAddress", _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StakedMATIC *StakedMATICSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetInsuranceAddress(&_StakedMATIC.TransactOpts, _address)
}

// SetInsuranceAddress is a paid mutator transaction binding the contract method 0xbb208f55.
//
// Solidity: function setInsuranceAddress(address _address) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetInsuranceAddress(_address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetInsuranceAddress(&_StakedMATIC.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StakedMATIC *StakedMATICTransactor) SetNodeOperatorRegistryAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setNodeOperatorRegistryAddress", _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StakedMATIC *StakedMATICSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetNodeOperatorRegistryAddress(&_StakedMATIC.TransactOpts, _address)
}

// SetNodeOperatorRegistryAddress is a paid mutator transaction binding the contract method 0x0f2b2639.
//
// Solidity: function setNodeOperatorRegistryAddress(address _address) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetNodeOperatorRegistryAddress(_address common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetNodeOperatorRegistryAddress(&_StakedMATIC.TransactOpts, _address)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StakedMATIC *StakedMATICTransactor) SetPoLidoNFT(opts *bind.TransactOpts, _newLidoNFT common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setPoLidoNFT", _newLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StakedMATIC *StakedMATICSession) SetPoLidoNFT(_newLidoNFT common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetPoLidoNFT(&_StakedMATIC.TransactOpts, _newLidoNFT)
}

// SetPoLidoNFT is a paid mutator transaction binding the contract method 0x15539d3f.
//
// Solidity: function setPoLidoNFT(address _newLidoNFT) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetPoLidoNFT(_newLidoNFT common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetPoLidoNFT(&_StakedMATIC.TransactOpts, _newLidoNFT)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StakedMATIC *StakedMATICTransactor) SetProtocolFee(opts *bind.TransactOpts, _newProtocolFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setProtocolFee", _newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StakedMATIC *StakedMATICSession) SetProtocolFee(_newProtocolFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetProtocolFee(&_StakedMATIC.TransactOpts, _newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x4e91f811.
//
// Solidity: function setProtocolFee(uint8 _newProtocolFee) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetProtocolFee(_newProtocolFee uint8) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetProtocolFee(&_StakedMATIC.TransactOpts, _newProtocolFee)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StakedMATIC *StakedMATICTransactor) SetRewardDistributionLowerBound(opts *bind.TransactOpts, _newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setRewardDistributionLowerBound", _newRewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StakedMATIC *StakedMATICSession) SetRewardDistributionLowerBound(_newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetRewardDistributionLowerBound(&_StakedMATIC.TransactOpts, _newRewardDistributionLowerBound)
}

// SetRewardDistributionLowerBound is a paid mutator transaction binding the contract method 0x3b573c4a.
//
// Solidity: function setRewardDistributionLowerBound(uint256 _newRewardDistributionLowerBound) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetRewardDistributionLowerBound(_newRewardDistributionLowerBound *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetRewardDistributionLowerBound(&_StakedMATIC.TransactOpts, _newRewardDistributionLowerBound)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StakedMATIC *StakedMATICTransactor) SetVersion(opts *bind.TransactOpts, _newVersion string) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "setVersion", _newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StakedMATIC *StakedMATICSession) SetVersion(_newVersion string) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetVersion(&_StakedMATIC.TransactOpts, _newVersion)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string _newVersion) returns()
func (_StakedMATIC *StakedMATICTransactorSession) SetVersion(_newVersion string) (*types.Transaction, error) {
	return _StakedMATIC.Contract.SetVersion(&_StakedMATIC.TransactOpts, _newVersion)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICTransactor) Submit(opts *bind.TransactOpts, _amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "submit", _amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICSession) Submit(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Submit(&_StakedMATIC.TransactOpts, _amount, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xf532e86a.
//
// Solidity: function submit(uint256 _amount, address _referral) returns(uint256)
func (_StakedMATIC *StakedMATICTransactorSession) Submit(_amount *big.Int, _referral common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Submit(&_StakedMATIC.TransactOpts, _amount, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Transfer(&_StakedMATIC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.Transfer(&_StakedMATIC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.TransferFrom(&_StakedMATIC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_StakedMATIC *StakedMATICTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakedMATIC.Contract.TransferFrom(&_StakedMATIC.TransactOpts, from, to, amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakedMATIC *StakedMATICTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakedMATIC *StakedMATICSession) Unpause() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Unpause(&_StakedMATIC.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakedMATIC *StakedMATICTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakedMATIC.Contract.Unpause(&_StakedMATIC.TransactOpts)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StakedMATIC *StakedMATICTransactor) WithdrawTotalDelegated(opts *bind.TransactOpts, _validatorShare common.Address) (*types.Transaction, error) {
	return _StakedMATIC.contract.Transact(opts, "withdrawTotalDelegated", _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StakedMATIC *StakedMATICSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.WithdrawTotalDelegated(&_StakedMATIC.TransactOpts, _validatorShare)
}

// WithdrawTotalDelegated is a paid mutator transaction binding the contract method 0xc75e7832.
//
// Solidity: function withdrawTotalDelegated(address _validatorShare) returns()
func (_StakedMATIC *StakedMATICTransactorSession) WithdrawTotalDelegated(_validatorShare common.Address) (*types.Transaction, error) {
	return _StakedMATIC.Contract.WithdrawTotalDelegated(&_StakedMATIC.TransactOpts, _validatorShare)
}

// StakedMATICApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedMATIC contract.
type StakedMATICApprovalIterator struct {
	Event *StakedMATICApproval // Event containing the contract specifics and raw log

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
func (it *StakedMATICApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICApproval)
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
		it.Event = new(StakedMATICApproval)
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
func (it *StakedMATICApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICApproval represents a Approval event raised by the StakedMATIC contract.
type StakedMATICApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedMATIC *StakedMATICFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakedMATICApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICApprovalIterator{contract: _StakedMATIC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedMATIC *StakedMATICFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedMATICApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICApproval)
				if err := _StakedMATIC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StakedMATIC *StakedMATICFilterer) ParseApproval(log types.Log) (*StakedMATICApproval, error) {
	event := new(StakedMATICApproval)
	if err := _StakedMATIC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICClaimTokensEventIterator is returned from FilterClaimTokensEvent and is used to iterate over the raw logs and unpacked data for ClaimTokensEvent events raised by the StakedMATIC contract.
type StakedMATICClaimTokensEventIterator struct {
	Event *StakedMATICClaimTokensEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICClaimTokensEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICClaimTokensEvent)
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
		it.Event = new(StakedMATICClaimTokensEvent)
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
func (it *StakedMATICClaimTokensEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICClaimTokensEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICClaimTokensEvent represents a ClaimTokensEvent event raised by the StakedMATIC contract.
type StakedMATICClaimTokensEvent struct {
	From          common.Address
	Id            *big.Int
	AmountClaimed *big.Int
	AmountBurned  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterClaimTokensEvent is a free log retrieval operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_StakedMATIC *StakedMATICFilterer) FilterClaimTokensEvent(opts *bind.FilterOpts, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (*StakedMATICClaimTokensEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICClaimTokensEventIterator{contract: _StakedMATIC.contract, event: "ClaimTokensEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTokensEvent is a free log subscription operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_StakedMATIC *StakedMATICFilterer) WatchClaimTokensEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICClaimTokensEvent, _from []common.Address, _id []*big.Int, _amountClaimed []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _amountClaimedRule []interface{}
	for _, _amountClaimedItem := range _amountClaimed {
		_amountClaimedRule = append(_amountClaimedRule, _amountClaimedItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "ClaimTokensEvent", _fromRule, _idRule, _amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICClaimTokensEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
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

// ParseClaimTokensEvent is a log parse operation binding the contract event 0xaca94a3466fab333b79851ab29b0715612740e4ae0d891ef8e9bd2a1bf5e24dd.
//
// Solidity: event ClaimTokensEvent(address indexed _from, uint256 indexed _id, uint256 indexed _amountClaimed, uint256 _amountBurned)
func (_StakedMATIC *StakedMATICFilterer) ParseClaimTokensEvent(log types.Log) (*StakedMATICClaimTokensEvent, error) {
	event := new(StakedMATICClaimTokensEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "ClaimTokensEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICClaimTotalDelegatedEventIterator is returned from FilterClaimTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for ClaimTotalDelegatedEvent events raised by the StakedMATIC contract.
type StakedMATICClaimTotalDelegatedEventIterator struct {
	Event *StakedMATICClaimTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICClaimTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICClaimTotalDelegatedEvent)
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
		it.Event = new(StakedMATICClaimTotalDelegatedEvent)
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
func (it *StakedMATICClaimTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICClaimTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICClaimTotalDelegatedEvent represents a ClaimTotalDelegatedEvent event raised by the StakedMATIC contract.
type StakedMATICClaimTotalDelegatedEvent struct {
	ValidatorShare common.Address
	AmountClaimed  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterClaimTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed validatorShare, uint256 indexed amountClaimed)
func (_StakedMATIC *StakedMATICFilterer) FilterClaimTotalDelegatedEvent(opts *bind.FilterOpts, validatorShare []common.Address, amountClaimed []*big.Int) (*StakedMATICClaimTotalDelegatedEventIterator, error) {

	var validatorShareRule []interface{}
	for _, validatorShareItem := range validatorShare {
		validatorShareRule = append(validatorShareRule, validatorShareItem)
	}
	var amountClaimedRule []interface{}
	for _, amountClaimedItem := range amountClaimed {
		amountClaimedRule = append(amountClaimedRule, amountClaimedItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "ClaimTotalDelegatedEvent", validatorShareRule, amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICClaimTotalDelegatedEventIterator{contract: _StakedMATIC.contract, event: "ClaimTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchClaimTotalDelegatedEvent is a free log subscription operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed validatorShare, uint256 indexed amountClaimed)
func (_StakedMATIC *StakedMATICFilterer) WatchClaimTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICClaimTotalDelegatedEvent, validatorShare []common.Address, amountClaimed []*big.Int) (event.Subscription, error) {

	var validatorShareRule []interface{}
	for _, validatorShareItem := range validatorShare {
		validatorShareRule = append(validatorShareRule, validatorShareItem)
	}
	var amountClaimedRule []interface{}
	for _, amountClaimedItem := range amountClaimed {
		amountClaimedRule = append(amountClaimedRule, amountClaimedItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "ClaimTotalDelegatedEvent", validatorShareRule, amountClaimedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICClaimTotalDelegatedEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
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

// ParseClaimTotalDelegatedEvent is a log parse operation binding the contract event 0x4c42a3bec298a4d82d41b7a540d8ebc22d91ee8a61459bce23849ff470d31dea.
//
// Solidity: event ClaimTotalDelegatedEvent(address indexed validatorShare, uint256 indexed amountClaimed)
func (_StakedMATIC *StakedMATICFilterer) ParseClaimTotalDelegatedEvent(log types.Log) (*StakedMATICClaimTotalDelegatedEvent, error) {
	event := new(StakedMATICClaimTotalDelegatedEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "ClaimTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICDelegateEventIterator is returned from FilterDelegateEvent and is used to iterate over the raw logs and unpacked data for DelegateEvent events raised by the StakedMATIC contract.
type StakedMATICDelegateEventIterator struct {
	Event *StakedMATICDelegateEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICDelegateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICDelegateEvent)
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
		it.Event = new(StakedMATICDelegateEvent)
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
func (it *StakedMATICDelegateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICDelegateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICDelegateEvent represents a DelegateEvent event raised by the StakedMATIC contract.
type StakedMATICDelegateEvent struct {
	AmountDelegated *big.Int
	Remainder       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateEvent is a free log retrieval operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_StakedMATIC *StakedMATICFilterer) FilterDelegateEvent(opts *bind.FilterOpts, _amountDelegated []*big.Int, _remainder []*big.Int) (*StakedMATICDelegateEventIterator, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICDelegateEventIterator{contract: _StakedMATIC.contract, event: "DelegateEvent", logs: logs, sub: sub}, nil
}

// WatchDelegateEvent is a free log subscription operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_StakedMATIC *StakedMATICFilterer) WatchDelegateEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICDelegateEvent, _amountDelegated []*big.Int, _remainder []*big.Int) (event.Subscription, error) {

	var _amountDelegatedRule []interface{}
	for _, _amountDelegatedItem := range _amountDelegated {
		_amountDelegatedRule = append(_amountDelegatedRule, _amountDelegatedItem)
	}
	var _remainderRule []interface{}
	for _, _remainderItem := range _remainder {
		_remainderRule = append(_remainderRule, _remainderItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "DelegateEvent", _amountDelegatedRule, _remainderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICDelegateEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
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

// ParseDelegateEvent is a log parse operation binding the contract event 0x421adba60af7a6b11679e2ac133b1bc91d3de91d56866ec19703d9d60cf950c8.
//
// Solidity: event DelegateEvent(uint256 indexed _amountDelegated, uint256 indexed _remainder)
func (_StakedMATIC *StakedMATICFilterer) ParseDelegateEvent(log types.Log) (*StakedMATICDelegateEvent, error) {
	event := new(StakedMATICDelegateEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "DelegateEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICDistributeRewardsEventIterator is returned from FilterDistributeRewardsEvent and is used to iterate over the raw logs and unpacked data for DistributeRewardsEvent events raised by the StakedMATIC contract.
type StakedMATICDistributeRewardsEventIterator struct {
	Event *StakedMATICDistributeRewardsEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICDistributeRewardsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICDistributeRewardsEvent)
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
		it.Event = new(StakedMATICDistributeRewardsEvent)
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
func (it *StakedMATICDistributeRewardsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICDistributeRewardsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICDistributeRewardsEvent represents a DistributeRewardsEvent event raised by the StakedMATIC contract.
type StakedMATICDistributeRewardsEvent struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributeRewardsEvent is a free log retrieval operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) FilterDistributeRewardsEvent(opts *bind.FilterOpts, _amount []*big.Int) (*StakedMATICDistributeRewardsEventIterator, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICDistributeRewardsEventIterator{contract: _StakedMATIC.contract, event: "DistributeRewardsEvent", logs: logs, sub: sub}, nil
}

// WatchDistributeRewardsEvent is a free log subscription operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) WatchDistributeRewardsEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICDistributeRewardsEvent, _amount []*big.Int) (event.Subscription, error) {

	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "DistributeRewardsEvent", _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICDistributeRewardsEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
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

// ParseDistributeRewardsEvent is a log parse operation binding the contract event 0x4e3c6a1e602996ae70905ac6165ed2434753246e3bfa52b6ca6852b40e2d4408.
//
// Solidity: event DistributeRewardsEvent(uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) ParseDistributeRewardsEvent(log types.Log) (*StakedMATICDistributeRewardsEvent, error) {
	event := new(StakedMATICDistributeRewardsEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "DistributeRewardsEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakedMATIC contract.
type StakedMATICPausedIterator struct {
	Event *StakedMATICPaused // Event containing the contract specifics and raw log

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
func (it *StakedMATICPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICPaused)
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
		it.Event = new(StakedMATICPaused)
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
func (it *StakedMATICPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICPaused represents a Paused event raised by the StakedMATIC contract.
type StakedMATICPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakedMATIC *StakedMATICFilterer) FilterPaused(opts *bind.FilterOpts) (*StakedMATICPausedIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakedMATICPausedIterator{contract: _StakedMATIC.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakedMATIC *StakedMATICFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakedMATICPaused) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICPaused)
				if err := _StakedMATIC.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakedMATIC *StakedMATICFilterer) ParsePaused(log types.Log) (*StakedMATICPaused, error) {
	event := new(StakedMATICPaused)
	if err := _StakedMATIC.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICRequestWithdrawEventIterator is returned from FilterRequestWithdrawEvent and is used to iterate over the raw logs and unpacked data for RequestWithdrawEvent events raised by the StakedMATIC contract.
type StakedMATICRequestWithdrawEventIterator struct {
	Event *StakedMATICRequestWithdrawEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICRequestWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICRequestWithdrawEvent)
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
		it.Event = new(StakedMATICRequestWithdrawEvent)
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
func (it *StakedMATICRequestWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICRequestWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICRequestWithdrawEvent represents a RequestWithdrawEvent event raised by the StakedMATIC contract.
type StakedMATICRequestWithdrawEvent struct {
	From     common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestWithdrawEvent is a free log retrieval operation binding the contract event 0x4318b22a7b774533f1c9cd7102530d96faffc18ef44a1ecb56abc9a55d49fd8b.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) FilterRequestWithdrawEvent(opts *bind.FilterOpts, _from []common.Address, _referral []common.Address) (*StakedMATICRequestWithdrawEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "RequestWithdrawEvent", _fromRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICRequestWithdrawEventIterator{contract: _StakedMATIC.contract, event: "RequestWithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchRequestWithdrawEvent is a free log subscription operation binding the contract event 0x4318b22a7b774533f1c9cd7102530d96faffc18ef44a1ecb56abc9a55d49fd8b.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) WatchRequestWithdrawEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICRequestWithdrawEvent, _from []common.Address, _referral []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "RequestWithdrawEvent", _fromRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICRequestWithdrawEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
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

// ParseRequestWithdrawEvent is a log parse operation binding the contract event 0x4318b22a7b774533f1c9cd7102530d96faffc18ef44a1ecb56abc9a55d49fd8b.
//
// Solidity: event RequestWithdrawEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) ParseRequestWithdrawEvent(log types.Log) (*StakedMATICRequestWithdrawEvent, error) {
	event := new(StakedMATICRequestWithdrawEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "RequestWithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakedMATIC contract.
type StakedMATICRoleAdminChangedIterator struct {
	Event *StakedMATICRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakedMATICRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICRoleAdminChanged)
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
		it.Event = new(StakedMATICRoleAdminChanged)
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
func (it *StakedMATICRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICRoleAdminChanged represents a RoleAdminChanged event raised by the StakedMATIC contract.
type StakedMATICRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedMATIC *StakedMATICFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakedMATICRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICRoleAdminChangedIterator{contract: _StakedMATIC.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedMATIC *StakedMATICFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakedMATICRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICRoleAdminChanged)
				if err := _StakedMATIC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakedMATIC *StakedMATICFilterer) ParseRoleAdminChanged(log types.Log) (*StakedMATICRoleAdminChanged, error) {
	event := new(StakedMATICRoleAdminChanged)
	if err := _StakedMATIC.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakedMATIC contract.
type StakedMATICRoleGrantedIterator struct {
	Event *StakedMATICRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakedMATICRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICRoleGranted)
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
		it.Event = new(StakedMATICRoleGranted)
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
func (it *StakedMATICRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICRoleGranted represents a RoleGranted event raised by the StakedMATIC contract.
type StakedMATICRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedMATIC *StakedMATICFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedMATICRoleGrantedIterator, error) {

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

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICRoleGrantedIterator{contract: _StakedMATIC.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedMATIC *StakedMATICFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakedMATICRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICRoleGranted)
				if err := _StakedMATIC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakedMATIC *StakedMATICFilterer) ParseRoleGranted(log types.Log) (*StakedMATICRoleGranted, error) {
	event := new(StakedMATICRoleGranted)
	if err := _StakedMATIC.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakedMATIC contract.
type StakedMATICRoleRevokedIterator struct {
	Event *StakedMATICRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakedMATICRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICRoleRevoked)
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
		it.Event = new(StakedMATICRoleRevoked)
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
func (it *StakedMATICRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICRoleRevoked represents a RoleRevoked event raised by the StakedMATIC contract.
type StakedMATICRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedMATIC *StakedMATICFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedMATICRoleRevokedIterator, error) {

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

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICRoleRevokedIterator{contract: _StakedMATIC.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedMATIC *StakedMATICFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakedMATICRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICRoleRevoked)
				if err := _StakedMATIC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakedMATIC *StakedMATICFilterer) ParseRoleRevoked(log types.Log) (*StakedMATICRoleRevoked, error) {
	event := new(StakedMATICRoleRevoked)
	if err := _StakedMATIC.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetDaoAddressIterator is returned from FilterSetDaoAddress and is used to iterate over the raw logs and unpacked data for SetDaoAddress events raised by the StakedMATIC contract.
type StakedMATICSetDaoAddressIterator struct {
	Event *StakedMATICSetDaoAddress // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetDaoAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetDaoAddress)
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
		it.Event = new(StakedMATICSetDaoAddress)
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
func (it *StakedMATICSetDaoAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetDaoAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetDaoAddress represents a SetDaoAddress event raised by the StakedMATIC contract.
type StakedMATICSetDaoAddress struct {
	OldDaoAddress common.Address
	NewDaoAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetDaoAddress is a free log retrieval operation binding the contract event 0x3b8dbc80bf27331221431f1c8b2c6fe358eadfdb3c3d7085e6a2521eeb775b07.
//
// Solidity: event SetDaoAddress(address oldDaoAddress, address newDaoAddress)
func (_StakedMATIC *StakedMATICFilterer) FilterSetDaoAddress(opts *bind.FilterOpts) (*StakedMATICSetDaoAddressIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetDaoAddress")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetDaoAddressIterator{contract: _StakedMATIC.contract, event: "SetDaoAddress", logs: logs, sub: sub}, nil
}

// WatchSetDaoAddress is a free log subscription operation binding the contract event 0x3b8dbc80bf27331221431f1c8b2c6fe358eadfdb3c3d7085e6a2521eeb775b07.
//
// Solidity: event SetDaoAddress(address oldDaoAddress, address newDaoAddress)
func (_StakedMATIC *StakedMATICFilterer) WatchSetDaoAddress(opts *bind.WatchOpts, sink chan<- *StakedMATICSetDaoAddress) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetDaoAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetDaoAddress)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetDaoAddress", log); err != nil {
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

// ParseSetDaoAddress is a log parse operation binding the contract event 0x3b8dbc80bf27331221431f1c8b2c6fe358eadfdb3c3d7085e6a2521eeb775b07.
//
// Solidity: event SetDaoAddress(address oldDaoAddress, address newDaoAddress)
func (_StakedMATIC *StakedMATICFilterer) ParseSetDaoAddress(log types.Log) (*StakedMATICSetDaoAddress, error) {
	event := new(StakedMATICSetDaoAddress)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetDaoAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetDelegationLowerBoundIterator is returned from FilterSetDelegationLowerBound and is used to iterate over the raw logs and unpacked data for SetDelegationLowerBound events raised by the StakedMATIC contract.
type StakedMATICSetDelegationLowerBoundIterator struct {
	Event *StakedMATICSetDelegationLowerBound // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetDelegationLowerBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetDelegationLowerBound)
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
		it.Event = new(StakedMATICSetDelegationLowerBound)
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
func (it *StakedMATICSetDelegationLowerBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetDelegationLowerBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetDelegationLowerBound represents a SetDelegationLowerBound event raised by the StakedMATIC contract.
type StakedMATICSetDelegationLowerBound struct {
	DelegationLowerBound *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetDelegationLowerBound is a free log retrieval operation binding the contract event 0x1cabc2f7b706218bb8613769cd658789cd6f1860310413b841020a2c7b7a0e32.
//
// Solidity: event SetDelegationLowerBound(uint256 indexed _delegationLowerBound)
func (_StakedMATIC *StakedMATICFilterer) FilterSetDelegationLowerBound(opts *bind.FilterOpts, _delegationLowerBound []*big.Int) (*StakedMATICSetDelegationLowerBoundIterator, error) {

	var _delegationLowerBoundRule []interface{}
	for _, _delegationLowerBoundItem := range _delegationLowerBound {
		_delegationLowerBoundRule = append(_delegationLowerBoundRule, _delegationLowerBoundItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetDelegationLowerBound", _delegationLowerBoundRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetDelegationLowerBoundIterator{contract: _StakedMATIC.contract, event: "SetDelegationLowerBound", logs: logs, sub: sub}, nil
}

// WatchSetDelegationLowerBound is a free log subscription operation binding the contract event 0x1cabc2f7b706218bb8613769cd658789cd6f1860310413b841020a2c7b7a0e32.
//
// Solidity: event SetDelegationLowerBound(uint256 indexed _delegationLowerBound)
func (_StakedMATIC *StakedMATICFilterer) WatchSetDelegationLowerBound(opts *bind.WatchOpts, sink chan<- *StakedMATICSetDelegationLowerBound, _delegationLowerBound []*big.Int) (event.Subscription, error) {

	var _delegationLowerBoundRule []interface{}
	for _, _delegationLowerBoundItem := range _delegationLowerBound {
		_delegationLowerBoundRule = append(_delegationLowerBoundRule, _delegationLowerBoundItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetDelegationLowerBound", _delegationLowerBoundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetDelegationLowerBound)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetDelegationLowerBound", log); err != nil {
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

// ParseSetDelegationLowerBound is a log parse operation binding the contract event 0x1cabc2f7b706218bb8613769cd658789cd6f1860310413b841020a2c7b7a0e32.
//
// Solidity: event SetDelegationLowerBound(uint256 indexed _delegationLowerBound)
func (_StakedMATIC *StakedMATICFilterer) ParseSetDelegationLowerBound(log types.Log) (*StakedMATICSetDelegationLowerBound, error) {
	event := new(StakedMATICSetDelegationLowerBound)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetDelegationLowerBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetFeesIterator is returned from FilterSetFees and is used to iterate over the raw logs and unpacked data for SetFees events raised by the StakedMATIC contract.
type StakedMATICSetFeesIterator struct {
	Event *StakedMATICSetFees // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetFees)
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
		it.Event = new(StakedMATICSetFees)
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
func (it *StakedMATICSetFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetFees represents a SetFees event raised by the StakedMATIC contract.
type StakedMATICSetFees struct {
	DaoFee       *big.Int
	OperatorsFee *big.Int
	InsuranceFee *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetFees is a free log retrieval operation binding the contract event 0x37322890d66d781059d797be5e2f27dc160a34d8bc0a8e09116cb9a773ce88ef.
//
// Solidity: event SetFees(uint256 daoFee, uint256 operatorsFee, uint256 insuranceFee)
func (_StakedMATIC *StakedMATICFilterer) FilterSetFees(opts *bind.FilterOpts) (*StakedMATICSetFeesIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetFees")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetFeesIterator{contract: _StakedMATIC.contract, event: "SetFees", logs: logs, sub: sub}, nil
}

// WatchSetFees is a free log subscription operation binding the contract event 0x37322890d66d781059d797be5e2f27dc160a34d8bc0a8e09116cb9a773ce88ef.
//
// Solidity: event SetFees(uint256 daoFee, uint256 operatorsFee, uint256 insuranceFee)
func (_StakedMATIC *StakedMATICFilterer) WatchSetFees(opts *bind.WatchOpts, sink chan<- *StakedMATICSetFees) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetFees)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetFees", log); err != nil {
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

// ParseSetFees is a log parse operation binding the contract event 0x37322890d66d781059d797be5e2f27dc160a34d8bc0a8e09116cb9a773ce88ef.
//
// Solidity: event SetFees(uint256 daoFee, uint256 operatorsFee, uint256 insuranceFee)
func (_StakedMATIC *StakedMATICFilterer) ParseSetFees(log types.Log) (*StakedMATICSetFees, error) {
	event := new(StakedMATICSetFees)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetFxStateRootTunnelIterator is returned from FilterSetFxStateRootTunnel and is used to iterate over the raw logs and unpacked data for SetFxStateRootTunnel events raised by the StakedMATIC contract.
type StakedMATICSetFxStateRootTunnelIterator struct {
	Event *StakedMATICSetFxStateRootTunnel // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetFxStateRootTunnelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetFxStateRootTunnel)
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
		it.Event = new(StakedMATICSetFxStateRootTunnel)
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
func (it *StakedMATICSetFxStateRootTunnelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetFxStateRootTunnelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetFxStateRootTunnel represents a SetFxStateRootTunnel event raised by the StakedMATIC contract.
type StakedMATICSetFxStateRootTunnel struct {
	OldFxStateRootTunnel common.Address
	NewFxStateRootTunnel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetFxStateRootTunnel is a free log retrieval operation binding the contract event 0x8f8196a0718fc814ab20b90fdc8c2024046578345148e1b903b20667f693a6ba.
//
// Solidity: event SetFxStateRootTunnel(address oldFxStateRootTunnel, address newFxStateRootTunnel)
func (_StakedMATIC *StakedMATICFilterer) FilterSetFxStateRootTunnel(opts *bind.FilterOpts) (*StakedMATICSetFxStateRootTunnelIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetFxStateRootTunnel")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetFxStateRootTunnelIterator{contract: _StakedMATIC.contract, event: "SetFxStateRootTunnel", logs: logs, sub: sub}, nil
}

// WatchSetFxStateRootTunnel is a free log subscription operation binding the contract event 0x8f8196a0718fc814ab20b90fdc8c2024046578345148e1b903b20667f693a6ba.
//
// Solidity: event SetFxStateRootTunnel(address oldFxStateRootTunnel, address newFxStateRootTunnel)
func (_StakedMATIC *StakedMATICFilterer) WatchSetFxStateRootTunnel(opts *bind.WatchOpts, sink chan<- *StakedMATICSetFxStateRootTunnel) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetFxStateRootTunnel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetFxStateRootTunnel)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetFxStateRootTunnel", log); err != nil {
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

// ParseSetFxStateRootTunnel is a log parse operation binding the contract event 0x8f8196a0718fc814ab20b90fdc8c2024046578345148e1b903b20667f693a6ba.
//
// Solidity: event SetFxStateRootTunnel(address oldFxStateRootTunnel, address newFxStateRootTunnel)
func (_StakedMATIC *StakedMATICFilterer) ParseSetFxStateRootTunnel(log types.Log) (*StakedMATICSetFxStateRootTunnel, error) {
	event := new(StakedMATICSetFxStateRootTunnel)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetFxStateRootTunnel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetInsuranceAddressIterator is returned from FilterSetInsuranceAddress and is used to iterate over the raw logs and unpacked data for SetInsuranceAddress events raised by the StakedMATIC contract.
type StakedMATICSetInsuranceAddressIterator struct {
	Event *StakedMATICSetInsuranceAddress // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetInsuranceAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetInsuranceAddress)
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
		it.Event = new(StakedMATICSetInsuranceAddress)
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
func (it *StakedMATICSetInsuranceAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetInsuranceAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetInsuranceAddress represents a SetInsuranceAddress event raised by the StakedMATIC contract.
type StakedMATICSetInsuranceAddress struct {
	NewInsuranceAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetInsuranceAddress is a free log retrieval operation binding the contract event 0x4029fa39dede0b39dd254005137e18c9116d6b2620b2037632c09bf89404c6d7.
//
// Solidity: event SetInsuranceAddress(address indexed _newInsuranceAddress)
func (_StakedMATIC *StakedMATICFilterer) FilterSetInsuranceAddress(opts *bind.FilterOpts, _newInsuranceAddress []common.Address) (*StakedMATICSetInsuranceAddressIterator, error) {

	var _newInsuranceAddressRule []interface{}
	for _, _newInsuranceAddressItem := range _newInsuranceAddress {
		_newInsuranceAddressRule = append(_newInsuranceAddressRule, _newInsuranceAddressItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetInsuranceAddress", _newInsuranceAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetInsuranceAddressIterator{contract: _StakedMATIC.contract, event: "SetInsuranceAddress", logs: logs, sub: sub}, nil
}

// WatchSetInsuranceAddress is a free log subscription operation binding the contract event 0x4029fa39dede0b39dd254005137e18c9116d6b2620b2037632c09bf89404c6d7.
//
// Solidity: event SetInsuranceAddress(address indexed _newInsuranceAddress)
func (_StakedMATIC *StakedMATICFilterer) WatchSetInsuranceAddress(opts *bind.WatchOpts, sink chan<- *StakedMATICSetInsuranceAddress, _newInsuranceAddress []common.Address) (event.Subscription, error) {

	var _newInsuranceAddressRule []interface{}
	for _, _newInsuranceAddressItem := range _newInsuranceAddress {
		_newInsuranceAddressRule = append(_newInsuranceAddressRule, _newInsuranceAddressItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetInsuranceAddress", _newInsuranceAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetInsuranceAddress)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetInsuranceAddress", log); err != nil {
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

// ParseSetInsuranceAddress is a log parse operation binding the contract event 0x4029fa39dede0b39dd254005137e18c9116d6b2620b2037632c09bf89404c6d7.
//
// Solidity: event SetInsuranceAddress(address indexed _newInsuranceAddress)
func (_StakedMATIC *StakedMATICFilterer) ParseSetInsuranceAddress(log types.Log) (*StakedMATICSetInsuranceAddress, error) {
	event := new(StakedMATICSetInsuranceAddress)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetInsuranceAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetLidoNFTIterator is returned from FilterSetLidoNFT and is used to iterate over the raw logs and unpacked data for SetLidoNFT events raised by the StakedMATIC contract.
type StakedMATICSetLidoNFTIterator struct {
	Event *StakedMATICSetLidoNFT // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetLidoNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetLidoNFT)
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
		it.Event = new(StakedMATICSetLidoNFT)
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
func (it *StakedMATICSetLidoNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetLidoNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetLidoNFT represents a SetLidoNFT event raised by the StakedMATIC contract.
type StakedMATICSetLidoNFT struct {
	OldLidoNFT common.Address
	NewLidoNFT common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetLidoNFT is a free log retrieval operation binding the contract event 0x3879f4996438c287cec42ea09f698215be51d075ec446a9f5bc83c98ccc19101.
//
// Solidity: event SetLidoNFT(address oldLidoNFT, address newLidoNFT)
func (_StakedMATIC *StakedMATICFilterer) FilterSetLidoNFT(opts *bind.FilterOpts) (*StakedMATICSetLidoNFTIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetLidoNFT")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetLidoNFTIterator{contract: _StakedMATIC.contract, event: "SetLidoNFT", logs: logs, sub: sub}, nil
}

// WatchSetLidoNFT is a free log subscription operation binding the contract event 0x3879f4996438c287cec42ea09f698215be51d075ec446a9f5bc83c98ccc19101.
//
// Solidity: event SetLidoNFT(address oldLidoNFT, address newLidoNFT)
func (_StakedMATIC *StakedMATICFilterer) WatchSetLidoNFT(opts *bind.WatchOpts, sink chan<- *StakedMATICSetLidoNFT) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetLidoNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetLidoNFT)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetLidoNFT", log); err != nil {
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

// ParseSetLidoNFT is a log parse operation binding the contract event 0x3879f4996438c287cec42ea09f698215be51d075ec446a9f5bc83c98ccc19101.
//
// Solidity: event SetLidoNFT(address oldLidoNFT, address newLidoNFT)
func (_StakedMATIC *StakedMATICFilterer) ParseSetLidoNFT(log types.Log) (*StakedMATICSetLidoNFT, error) {
	event := new(StakedMATICSetLidoNFT)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetLidoNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetNodeOperatorRegistryAddressIterator is returned from FilterSetNodeOperatorRegistryAddress and is used to iterate over the raw logs and unpacked data for SetNodeOperatorRegistryAddress events raised by the StakedMATIC contract.
type StakedMATICSetNodeOperatorRegistryAddressIterator struct {
	Event *StakedMATICSetNodeOperatorRegistryAddress // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetNodeOperatorRegistryAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetNodeOperatorRegistryAddress)
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
		it.Event = new(StakedMATICSetNodeOperatorRegistryAddress)
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
func (it *StakedMATICSetNodeOperatorRegistryAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetNodeOperatorRegistryAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetNodeOperatorRegistryAddress represents a SetNodeOperatorRegistryAddress event raised by the StakedMATIC contract.
type StakedMATICSetNodeOperatorRegistryAddress struct {
	NewNodeOperatorRegistryAddress common.Address
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterSetNodeOperatorRegistryAddress is a free log retrieval operation binding the contract event 0xb8e1a40638c48c0ebe9679e0b5b032f2066cf44139d775cc09c945b5df070c4e.
//
// Solidity: event SetNodeOperatorRegistryAddress(address indexed _newNodeOperatorRegistryAddress)
func (_StakedMATIC *StakedMATICFilterer) FilterSetNodeOperatorRegistryAddress(opts *bind.FilterOpts, _newNodeOperatorRegistryAddress []common.Address) (*StakedMATICSetNodeOperatorRegistryAddressIterator, error) {

	var _newNodeOperatorRegistryAddressRule []interface{}
	for _, _newNodeOperatorRegistryAddressItem := range _newNodeOperatorRegistryAddress {
		_newNodeOperatorRegistryAddressRule = append(_newNodeOperatorRegistryAddressRule, _newNodeOperatorRegistryAddressItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetNodeOperatorRegistryAddress", _newNodeOperatorRegistryAddressRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetNodeOperatorRegistryAddressIterator{contract: _StakedMATIC.contract, event: "SetNodeOperatorRegistryAddress", logs: logs, sub: sub}, nil
}

// WatchSetNodeOperatorRegistryAddress is a free log subscription operation binding the contract event 0xb8e1a40638c48c0ebe9679e0b5b032f2066cf44139d775cc09c945b5df070c4e.
//
// Solidity: event SetNodeOperatorRegistryAddress(address indexed _newNodeOperatorRegistryAddress)
func (_StakedMATIC *StakedMATICFilterer) WatchSetNodeOperatorRegistryAddress(opts *bind.WatchOpts, sink chan<- *StakedMATICSetNodeOperatorRegistryAddress, _newNodeOperatorRegistryAddress []common.Address) (event.Subscription, error) {

	var _newNodeOperatorRegistryAddressRule []interface{}
	for _, _newNodeOperatorRegistryAddressItem := range _newNodeOperatorRegistryAddress {
		_newNodeOperatorRegistryAddressRule = append(_newNodeOperatorRegistryAddressRule, _newNodeOperatorRegistryAddressItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetNodeOperatorRegistryAddress", _newNodeOperatorRegistryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetNodeOperatorRegistryAddress)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetNodeOperatorRegistryAddress", log); err != nil {
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

// ParseSetNodeOperatorRegistryAddress is a log parse operation binding the contract event 0xb8e1a40638c48c0ebe9679e0b5b032f2066cf44139d775cc09c945b5df070c4e.
//
// Solidity: event SetNodeOperatorRegistryAddress(address indexed _newNodeOperatorRegistryAddress)
func (_StakedMATIC *StakedMATICFilterer) ParseSetNodeOperatorRegistryAddress(log types.Log) (*StakedMATICSetNodeOperatorRegistryAddress, error) {
	event := new(StakedMATICSetNodeOperatorRegistryAddress)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetNodeOperatorRegistryAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetProtocolFeeIterator is returned from FilterSetProtocolFee and is used to iterate over the raw logs and unpacked data for SetProtocolFee events raised by the StakedMATIC contract.
type StakedMATICSetProtocolFeeIterator struct {
	Event *StakedMATICSetProtocolFee // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetProtocolFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetProtocolFee)
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
		it.Event = new(StakedMATICSetProtocolFee)
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
func (it *StakedMATICSetProtocolFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetProtocolFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetProtocolFee represents a SetProtocolFee event raised by the StakedMATIC contract.
type StakedMATICSetProtocolFee struct {
	OldProtocolFee uint8
	NewProtocolFee uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSetProtocolFee is a free log retrieval operation binding the contract event 0x6b1719571aee7af62357ac4d4c98cc35155a52a5fcf0c09198874443c0fe430d.
//
// Solidity: event SetProtocolFee(uint8 oldProtocolFee, uint8 newProtocolFee)
func (_StakedMATIC *StakedMATICFilterer) FilterSetProtocolFee(opts *bind.FilterOpts) (*StakedMATICSetProtocolFeeIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetProtocolFee")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetProtocolFeeIterator{contract: _StakedMATIC.contract, event: "SetProtocolFee", logs: logs, sub: sub}, nil
}

// WatchSetProtocolFee is a free log subscription operation binding the contract event 0x6b1719571aee7af62357ac4d4c98cc35155a52a5fcf0c09198874443c0fe430d.
//
// Solidity: event SetProtocolFee(uint8 oldProtocolFee, uint8 newProtocolFee)
func (_StakedMATIC *StakedMATICFilterer) WatchSetProtocolFee(opts *bind.WatchOpts, sink chan<- *StakedMATICSetProtocolFee) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetProtocolFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetProtocolFee)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetProtocolFee", log); err != nil {
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

// ParseSetProtocolFee is a log parse operation binding the contract event 0x6b1719571aee7af62357ac4d4c98cc35155a52a5fcf0c09198874443c0fe430d.
//
// Solidity: event SetProtocolFee(uint8 oldProtocolFee, uint8 newProtocolFee)
func (_StakedMATIC *StakedMATICFilterer) ParseSetProtocolFee(log types.Log) (*StakedMATICSetProtocolFee, error) {
	event := new(StakedMATICSetProtocolFee)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetProtocolFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSetRewardDistributionLowerBoundIterator is returned from FilterSetRewardDistributionLowerBound and is used to iterate over the raw logs and unpacked data for SetRewardDistributionLowerBound events raised by the StakedMATIC contract.
type StakedMATICSetRewardDistributionLowerBoundIterator struct {
	Event *StakedMATICSetRewardDistributionLowerBound // Event containing the contract specifics and raw log

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
func (it *StakedMATICSetRewardDistributionLowerBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSetRewardDistributionLowerBound)
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
		it.Event = new(StakedMATICSetRewardDistributionLowerBound)
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
func (it *StakedMATICSetRewardDistributionLowerBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSetRewardDistributionLowerBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSetRewardDistributionLowerBound represents a SetRewardDistributionLowerBound event raised by the StakedMATIC contract.
type StakedMATICSetRewardDistributionLowerBound struct {
	OldRewardDistributionLowerBound *big.Int
	NewRewardDistributionLowerBound *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterSetRewardDistributionLowerBound is a free log retrieval operation binding the contract event 0xbf99ce7c5a72f4c7135bb72a36193d147a37a54bfcf63ec29505b5e6d5e2921d.
//
// Solidity: event SetRewardDistributionLowerBound(uint256 oldRewardDistributionLowerBound, uint256 newRewardDistributionLowerBound)
func (_StakedMATIC *StakedMATICFilterer) FilterSetRewardDistributionLowerBound(opts *bind.FilterOpts) (*StakedMATICSetRewardDistributionLowerBoundIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SetRewardDistributionLowerBound")
	if err != nil {
		return nil, err
	}
	return &StakedMATICSetRewardDistributionLowerBoundIterator{contract: _StakedMATIC.contract, event: "SetRewardDistributionLowerBound", logs: logs, sub: sub}, nil
}

// WatchSetRewardDistributionLowerBound is a free log subscription operation binding the contract event 0xbf99ce7c5a72f4c7135bb72a36193d147a37a54bfcf63ec29505b5e6d5e2921d.
//
// Solidity: event SetRewardDistributionLowerBound(uint256 oldRewardDistributionLowerBound, uint256 newRewardDistributionLowerBound)
func (_StakedMATIC *StakedMATICFilterer) WatchSetRewardDistributionLowerBound(opts *bind.WatchOpts, sink chan<- *StakedMATICSetRewardDistributionLowerBound) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SetRewardDistributionLowerBound")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSetRewardDistributionLowerBound)
				if err := _StakedMATIC.contract.UnpackLog(event, "SetRewardDistributionLowerBound", log); err != nil {
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

// ParseSetRewardDistributionLowerBound is a log parse operation binding the contract event 0xbf99ce7c5a72f4c7135bb72a36193d147a37a54bfcf63ec29505b5e6d5e2921d.
//
// Solidity: event SetRewardDistributionLowerBound(uint256 oldRewardDistributionLowerBound, uint256 newRewardDistributionLowerBound)
func (_StakedMATIC *StakedMATICFilterer) ParseSetRewardDistributionLowerBound(log types.Log) (*StakedMATICSetRewardDistributionLowerBound, error) {
	event := new(StakedMATICSetRewardDistributionLowerBound)
	if err := _StakedMATIC.contract.UnpackLog(event, "SetRewardDistributionLowerBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICSubmitEventIterator is returned from FilterSubmitEvent and is used to iterate over the raw logs and unpacked data for SubmitEvent events raised by the StakedMATIC contract.
type StakedMATICSubmitEventIterator struct {
	Event *StakedMATICSubmitEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICSubmitEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICSubmitEvent)
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
		it.Event = new(StakedMATICSubmitEvent)
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
func (it *StakedMATICSubmitEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICSubmitEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICSubmitEvent represents a SubmitEvent event raised by the StakedMATIC contract.
type StakedMATICSubmitEvent struct {
	From     common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitEvent is a free log retrieval operation binding the contract event 0x98d2bc018caf34c71a8f920d9d93d4ed62e9789506b74087b48570c17b28ed99.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) FilterSubmitEvent(opts *bind.FilterOpts, _from []common.Address, _referral []common.Address) (*StakedMATICSubmitEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "SubmitEvent", _fromRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICSubmitEventIterator{contract: _StakedMATIC.contract, event: "SubmitEvent", logs: logs, sub: sub}, nil
}

// WatchSubmitEvent is a free log subscription operation binding the contract event 0x98d2bc018caf34c71a8f920d9d93d4ed62e9789506b74087b48570c17b28ed99.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) WatchSubmitEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICSubmitEvent, _from []common.Address, _referral []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "SubmitEvent", _fromRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICSubmitEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
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

// ParseSubmitEvent is a log parse operation binding the contract event 0x98d2bc018caf34c71a8f920d9d93d4ed62e9789506b74087b48570c17b28ed99.
//
// Solidity: event SubmitEvent(address indexed _from, uint256 _amount, address indexed _referral)
func (_StakedMATIC *StakedMATICFilterer) ParseSubmitEvent(log types.Log) (*StakedMATICSubmitEvent, error) {
	event := new(StakedMATICSubmitEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "SubmitEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakedMATIC contract.
type StakedMATICTransferIterator struct {
	Event *StakedMATICTransfer // Event containing the contract specifics and raw log

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
func (it *StakedMATICTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICTransfer)
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
		it.Event = new(StakedMATICTransfer)
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
func (it *StakedMATICTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICTransfer represents a Transfer event raised by the StakedMATIC contract.
type StakedMATICTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedMATIC *StakedMATICFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakedMATICTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICTransferIterator{contract: _StakedMATIC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedMATIC *StakedMATICFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakedMATICTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICTransfer)
				if err := _StakedMATIC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakedMATIC *StakedMATICFilterer) ParseTransfer(log types.Log) (*StakedMATICTransfer, error) {
	event := new(StakedMATICTransfer)
	if err := _StakedMATIC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakedMATIC contract.
type StakedMATICUnpausedIterator struct {
	Event *StakedMATICUnpaused // Event containing the contract specifics and raw log

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
func (it *StakedMATICUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICUnpaused)
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
		it.Event = new(StakedMATICUnpaused)
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
func (it *StakedMATICUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICUnpaused represents a Unpaused event raised by the StakedMATIC contract.
type StakedMATICUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakedMATIC *StakedMATICFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakedMATICUnpausedIterator, error) {

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakedMATICUnpausedIterator{contract: _StakedMATIC.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakedMATIC *StakedMATICFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakedMATICUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICUnpaused)
				if err := _StakedMATIC.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakedMATIC *StakedMATICFilterer) ParseUnpaused(log types.Log) (*StakedMATICUnpaused, error) {
	event := new(StakedMATICUnpaused)
	if err := _StakedMATIC.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICVersionIterator is returned from FilterVersion and is used to iterate over the raw logs and unpacked data for Version events raised by the StakedMATIC contract.
type StakedMATICVersionIterator struct {
	Event *StakedMATICVersion // Event containing the contract specifics and raw log

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
func (it *StakedMATICVersionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICVersion)
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
		it.Event = new(StakedMATICVersion)
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
func (it *StakedMATICVersionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICVersionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICVersion represents a Version event raised by the StakedMATIC contract.
type StakedMATICVersion struct {
	OldVersion string
	NewVersion common.Hash
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVersion is a free log retrieval operation binding the contract event 0xa22d531a51c0ad90c971d36d779910f19a3b85d5e5005072f84700bd68b6f0c5.
//
// Solidity: event Version(string oldVersion, string indexed newVersion)
func (_StakedMATIC *StakedMATICFilterer) FilterVersion(opts *bind.FilterOpts, newVersion []string) (*StakedMATICVersionIterator, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "Version", newVersionRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICVersionIterator{contract: _StakedMATIC.contract, event: "Version", logs: logs, sub: sub}, nil
}

// WatchVersion is a free log subscription operation binding the contract event 0xa22d531a51c0ad90c971d36d779910f19a3b85d5e5005072f84700bd68b6f0c5.
//
// Solidity: event Version(string oldVersion, string indexed newVersion)
func (_StakedMATIC *StakedMATICFilterer) WatchVersion(opts *bind.WatchOpts, sink chan<- *StakedMATICVersion, newVersion []string) (event.Subscription, error) {

	var newVersionRule []interface{}
	for _, newVersionItem := range newVersion {
		newVersionRule = append(newVersionRule, newVersionItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "Version", newVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICVersion)
				if err := _StakedMATIC.contract.UnpackLog(event, "Version", log); err != nil {
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

// ParseVersion is a log parse operation binding the contract event 0xa22d531a51c0ad90c971d36d779910f19a3b85d5e5005072f84700bd68b6f0c5.
//
// Solidity: event Version(string oldVersion, string indexed newVersion)
func (_StakedMATIC *StakedMATICFilterer) ParseVersion(log types.Log) (*StakedMATICVersion, error) {
	event := new(StakedMATICVersion)
	if err := _StakedMATIC.contract.UnpackLog(event, "Version", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedMATICWithdrawTotalDelegatedEventIterator is returned from FilterWithdrawTotalDelegatedEvent and is used to iterate over the raw logs and unpacked data for WithdrawTotalDelegatedEvent events raised by the StakedMATIC contract.
type StakedMATICWithdrawTotalDelegatedEventIterator struct {
	Event *StakedMATICWithdrawTotalDelegatedEvent // Event containing the contract specifics and raw log

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
func (it *StakedMATICWithdrawTotalDelegatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedMATICWithdrawTotalDelegatedEvent)
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
		it.Event = new(StakedMATICWithdrawTotalDelegatedEvent)
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
func (it *StakedMATICWithdrawTotalDelegatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedMATICWithdrawTotalDelegatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedMATICWithdrawTotalDelegatedEvent represents a WithdrawTotalDelegatedEvent event raised by the StakedMATIC contract.
type StakedMATICWithdrawTotalDelegatedEvent struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTotalDelegatedEvent is a free log retrieval operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) FilterWithdrawTotalDelegatedEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int) (*StakedMATICWithdrawTotalDelegatedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _StakedMATIC.contract.FilterLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &StakedMATICWithdrawTotalDelegatedEventIterator{contract: _StakedMATIC.contract, event: "WithdrawTotalDelegatedEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawTotalDelegatedEvent is a free log subscription operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) WatchWithdrawTotalDelegatedEvent(opts *bind.WatchOpts, sink chan<- *StakedMATICWithdrawTotalDelegatedEvent, _from []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _StakedMATIC.contract.WatchLogs(opts, "WithdrawTotalDelegatedEvent", _fromRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedMATICWithdrawTotalDelegatedEvent)
				if err := _StakedMATIC.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
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

// ParseWithdrawTotalDelegatedEvent is a log parse operation binding the contract event 0x65fcdf1cdc99352d178d6d953d52e01307cde7a592027b09c9e1d9ac8eb09ab7.
//
// Solidity: event WithdrawTotalDelegatedEvent(address indexed _from, uint256 indexed _amount)
func (_StakedMATIC *StakedMATICFilterer) ParseWithdrawTotalDelegatedEvent(log types.Log) (*StakedMATICWithdrawTotalDelegatedEvent, error) {
	event := new(StakedMATICWithdrawTotalDelegatedEvent)
	if err := _StakedMATIC.contract.UnpackLog(event, "WithdrawTotalDelegatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
