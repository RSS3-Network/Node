// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rss3

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

// DataTypesNode is an auto generated low-level Go binding around an user-defined struct.
type DataTypesNode struct {
	NodeId              *big.Int
	Account             common.Address
	TaxRateBasisPoints  uint64
	PublicGood          bool
	Alpha               bool
	Name                string
	Description         string
	OperationPoolTokens *big.Int
	StakingPoolTokens   *big.Int
	TotalShares         *big.Int
	SlashedTokens       *big.Int
}

// DataTypesUnstakeRequest is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUnstakeRequest struct {
	Owner         common.Address
	NodeAddr      common.Address
	Timestamp     *big.Int
	UnstakeAmount *big.Int
}

// DataTypesWithdrawalRequest is an auto generated low-level Go binding around an user-defined struct.
type DataTypesWithdrawalRequest struct {
	Owner     common.Address
	Timestamp *big.Int
	Amount    *big.Int
}

// StakingVSLMetaData contains all meta data concerning the StakingVSL contract.
var StakingVSLMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"treasury\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeUnbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositUnbondingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeSlashRateBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userSlashRateBasisPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTaxRateBasisPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlphaWithdrawNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchSizeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallerNotStaking\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ChipNotAuthorized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ChipNotPublicGood\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"ChipNotValid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChipsIdOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChipsNotSameOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimId\",\"type\":\"uint256\"}],\"name\":\"ClaimIdNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ClaimTimeNotReady\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CreateNodeToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyChipsIds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyNodeList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExcessWithdrawalAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"InvalidEpoch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"got\",\"type\":\"uint256\"}],\"name\":\"InvalidEpochNumber\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"traitId\",\"type\":\"uint256\"}],\"name\":\"InvalidTraitId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"NodeAlreadyPublicGood\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"NodeNotPublicGood\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeStakedOrDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationRewardsExceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicGoodNodeNotDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicGoodNodeNotInAlphaPhase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"RewardsAlreadyDistributed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SettlementPhase\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"StakeToPublicGoodNode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingRewardsExceed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubmissionIntervalNotElapsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TaxRateBasisPointsTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TaxRateBasisPointsTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"alpha\",\"type\":\"bool\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashedOperationPool\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"slashedStakingPool\",\"type\":\"uint256\"}],\"name\":\"NodeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"}],\"name\":\"NodeTaxRateBasisPointsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"NodeUpdated2PublicGood\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicPoolRewards\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicPoolTax\",\"type\":\"uint256\"}],\"name\":\"PublicGoodRewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"}],\"name\":\"PublicPoolTaxRateBasisPointsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nodeAddrs\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"operationRewards\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"stakingRewards\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"taxAmounts\",\"type\":\"uint256[]\"}],\"name\":\"RewardDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"name\":\"UnstakeClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"chipsIds\",\"type\":\"uint256[]\"}],\"name\":\"UnstakeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"WithdrawRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"WithdrawalClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_UNBONDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEPOSIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TAX_RATE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NODE_SLASH_RATE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHARES_PER_CHIP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKE_UNBONDING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TREASURY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_SLASH_RATE_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chipsContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"}],\"name\":\"createNode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableAlphaPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[3]\",\"name\":\"epochInfo\",\"type\":\"uint256[3]\"},{\"internalType\":\"address[]\",\"name\":\"nodeAddrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"operationRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakingRewards\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"publicPoolRewards\",\"type\":\"uint256\"}],\"name\":\"distributeRewards\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getChipsInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"getNode\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"alpha\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashedTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Component\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"getNodeAvatar\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodeAddrs\",\"type\":\"address[]\"}],\"name\":\"getNodes\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"alpha\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashedTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Component[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getNodesWithPagination\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"alpha\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashedTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Component[]\",\"name\":\"nodes\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getPendingUnstake\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UnstakeRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getPendingWithdrawal\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.WithdrawalRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalOperationPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStakingPoolTokens\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"publicGood\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"alpha\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"operationPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakingPoolTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashedTokens\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.Component\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"chips\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pauseAccount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAccount\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAlphaPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSettlementPhase\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"minTokensToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"chipsIds\",\"type\":\"uint256[]\"}],\"name\":\"requestUnstake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"requestWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"setSettlementPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"}],\"name\":\"setTaxRateBasisPoints4Node\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"taxRateBasisPoints\",\"type\":\"uint64\"}],\"name\":\"setTaxRateBasisPoints4PublicPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"nodeAddrs\",\"type\":\"address[]\"}],\"name\":\"slashNodes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nodeAddr\",\"type\":\"address\"}],\"name\":\"stakeToPublicPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateToPublicGood\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw2Treasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingVSLABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingVSLMetaData.ABI instead.
var StakingVSLABI = StakingVSLMetaData.ABI

// StakingVSL is an auto generated Go binding around an Ethereum contract.
type StakingVSL struct {
	StakingVSLCaller     // Read-only binding to the contract
	StakingVSLTransactor // Write-only binding to the contract
	StakingVSLFilterer   // Log filterer for contract events
}

// StakingVSLCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingVSLCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVSLTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingVSLTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVSLFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingVSLFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingVSLSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingVSLSession struct {
	Contract     *StakingVSL       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingVSLCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingVSLCallerSession struct {
	Contract *StakingVSLCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakingVSLTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingVSLTransactorSession struct {
	Contract     *StakingVSLTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakingVSLRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingVSLRaw struct {
	Contract *StakingVSL // Generic contract binding to access the raw methods on
}

// StakingVSLCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingVSLCallerRaw struct {
	Contract *StakingVSLCaller // Generic read-only contract binding to access the raw methods on
}

// StakingVSLTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingVSLTransactorRaw struct {
	Contract *StakingVSLTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingVSL creates a new instance of StakingVSL, bound to a specific deployed contract.
func NewStakingVSL(address common.Address, backend bind.ContractBackend) (*StakingVSL, error) {
	contract, err := bindStakingVSL(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingVSL{StakingVSLCaller: StakingVSLCaller{contract: contract}, StakingVSLTransactor: StakingVSLTransactor{contract: contract}, StakingVSLFilterer: StakingVSLFilterer{contract: contract}}, nil
}

// NewStakingVSLCaller creates a new read-only instance of StakingVSL, bound to a specific deployed contract.
func NewStakingVSLCaller(address common.Address, caller bind.ContractCaller) (*StakingVSLCaller, error) {
	contract, err := bindStakingVSL(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingVSLCaller{contract: contract}, nil
}

// NewStakingVSLTransactor creates a new write-only instance of StakingVSL, bound to a specific deployed contract.
func NewStakingVSLTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingVSLTransactor, error) {
	contract, err := bindStakingVSL(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingVSLTransactor{contract: contract}, nil
}

// NewStakingVSLFilterer creates a new log filterer instance of StakingVSL, bound to a specific deployed contract.
func NewStakingVSLFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingVSLFilterer, error) {
	contract, err := bindStakingVSL(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingVSLFilterer{contract: contract}, nil
}

// bindStakingVSL binds a generic wrapper to an already deployed contract.
func bindStakingVSL(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingVSLMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingVSL *StakingVSLRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingVSL.Contract.StakingVSLCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingVSL *StakingVSLRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.Contract.StakingVSLTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingVSL *StakingVSLRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingVSL.Contract.StakingVSLTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingVSL *StakingVSLCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingVSL.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingVSL *StakingVSLTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingVSL *StakingVSLTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingVSL.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingVSL.Contract.DEFAULTADMINROLE(&_StakingVSL.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingVSL.Contract.DEFAULTADMINROLE(&_StakingVSL.CallOpts)
}

// DEPOSITUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x6bdc11d5.
//
// Solidity: function DEPOSIT_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) DEPOSITUNBONDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "DEPOSIT_UNBONDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x6bdc11d5.
//
// Solidity: function DEPOSIT_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLSession) DEPOSITUNBONDINGPERIOD() (*big.Int, error) {
	return _StakingVSL.Contract.DEPOSITUNBONDINGPERIOD(&_StakingVSL.CallOpts)
}

// DEPOSITUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x6bdc11d5.
//
// Solidity: function DEPOSIT_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) DEPOSITUNBONDINGPERIOD() (*big.Int, error) {
	return _StakingVSL.Contract.DEPOSITUNBONDINGPERIOD(&_StakingVSL.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) MINDEPOSIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "MIN_DEPOSIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_StakingVSL *StakingVSLSession) MINDEPOSIT() (*big.Int, error) {
	return _StakingVSL.Contract.MINDEPOSIT(&_StakingVSL.CallOpts)
}

// MINDEPOSIT is a free data retrieval call binding the contract method 0xe1e158a5.
//
// Solidity: function MIN_DEPOSIT() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) MINDEPOSIT() (*big.Int, error) {
	return _StakingVSL.Contract.MINDEPOSIT(&_StakingVSL.CallOpts)
}

// MINTAXRATEBASISPOINTS is a free data retrieval call binding the contract method 0x2fe3a2a0.
//
// Solidity: function MIN_TAX_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) MINTAXRATEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "MIN_TAX_RATE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTAXRATEBASISPOINTS is a free data retrieval call binding the contract method 0x2fe3a2a0.
//
// Solidity: function MIN_TAX_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLSession) MINTAXRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.MINTAXRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// MINTAXRATEBASISPOINTS is a free data retrieval call binding the contract method 0x2fe3a2a0.
//
// Solidity: function MIN_TAX_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) MINTAXRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.MINTAXRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// NODESLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0x3daf051f.
//
// Solidity: function NODE_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) NODESLASHRATEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "NODE_SLASH_RATE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NODESLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0x3daf051f.
//
// Solidity: function NODE_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLSession) NODESLASHRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.NODESLASHRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// NODESLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0x3daf051f.
//
// Solidity: function NODE_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) NODESLASHRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.NODESLASHRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLSession) ORACLEROLE() ([32]byte, error) {
	return _StakingVSL.Contract.ORACLEROLE(&_StakingVSL.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCallerSession) ORACLEROLE() ([32]byte, error) {
	return _StakingVSL.Contract.ORACLEROLE(&_StakingVSL.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLSession) PAUSEROLE() ([32]byte, error) {
	return _StakingVSL.Contract.PAUSEROLE(&_StakingVSL.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakingVSL *StakingVSLCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StakingVSL.Contract.PAUSEROLE(&_StakingVSL.CallOpts)
}

// SHARESPERCHIP is a free data retrieval call binding the contract method 0x6b05f6dc.
//
// Solidity: function SHARES_PER_CHIP() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) SHARESPERCHIP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "SHARES_PER_CHIP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHARESPERCHIP is a free data retrieval call binding the contract method 0x6b05f6dc.
//
// Solidity: function SHARES_PER_CHIP() view returns(uint256)
func (_StakingVSL *StakingVSLSession) SHARESPERCHIP() (*big.Int, error) {
	return _StakingVSL.Contract.SHARESPERCHIP(&_StakingVSL.CallOpts)
}

// SHARESPERCHIP is a free data retrieval call binding the contract method 0x6b05f6dc.
//
// Solidity: function SHARES_PER_CHIP() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) SHARESPERCHIP() (*big.Int, error) {
	return _StakingVSL.Contract.SHARESPERCHIP(&_StakingVSL.CallOpts)
}

// STAKERATIO is a free data retrieval call binding the contract method 0x736fcdf6.
//
// Solidity: function STAKE_RATIO() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) STAKERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "STAKE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKERATIO is a free data retrieval call binding the contract method 0x736fcdf6.
//
// Solidity: function STAKE_RATIO() view returns(uint256)
func (_StakingVSL *StakingVSLSession) STAKERATIO() (*big.Int, error) {
	return _StakingVSL.Contract.STAKERATIO(&_StakingVSL.CallOpts)
}

// STAKERATIO is a free data retrieval call binding the contract method 0x736fcdf6.
//
// Solidity: function STAKE_RATIO() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) STAKERATIO() (*big.Int, error) {
	return _StakingVSL.Contract.STAKERATIO(&_StakingVSL.CallOpts)
}

// STAKEUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x2606a44a.
//
// Solidity: function STAKE_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) STAKEUNBONDINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "STAKE_UNBONDING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKEUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x2606a44a.
//
// Solidity: function STAKE_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLSession) STAKEUNBONDINGPERIOD() (*big.Int, error) {
	return _StakingVSL.Contract.STAKEUNBONDINGPERIOD(&_StakingVSL.CallOpts)
}

// STAKEUNBONDINGPERIOD is a free data retrieval call binding the contract method 0x2606a44a.
//
// Solidity: function STAKE_UNBONDING_PERIOD() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) STAKEUNBONDINGPERIOD() (*big.Int, error) {
	return _StakingVSL.Contract.STAKEUNBONDINGPERIOD(&_StakingVSL.CallOpts)
}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_StakingVSL *StakingVSLCaller) TREASURY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "TREASURY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_StakingVSL *StakingVSLSession) TREASURY() (common.Address, error) {
	return _StakingVSL.Contract.TREASURY(&_StakingVSL.CallOpts)
}

// TREASURY is a free data retrieval call binding the contract method 0x2d2c5565.
//
// Solidity: function TREASURY() view returns(address)
func (_StakingVSL *StakingVSLCallerSession) TREASURY() (common.Address, error) {
	return _StakingVSL.Contract.TREASURY(&_StakingVSL.CallOpts)
}

// USERSLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0xb47d343c.
//
// Solidity: function USER_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) USERSLASHRATEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "USER_SLASH_RATE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USERSLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0xb47d343c.
//
// Solidity: function USER_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLSession) USERSLASHRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.USERSLASHRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// USERSLASHRATEBASISPOINTS is a free data retrieval call binding the contract method 0xb47d343c.
//
// Solidity: function USER_SLASH_RATE_BASIS_POINTS() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) USERSLASHRATEBASISPOINTS() (*big.Int, error) {
	return _StakingVSL.Contract.USERSLASHRATEBASISPOINTS(&_StakingVSL.CallOpts)
}

// ChipsContract is a free data retrieval call binding the contract method 0xd13b19a3.
//
// Solidity: function chipsContract() view returns(address)
func (_StakingVSL *StakingVSLCaller) ChipsContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "chipsContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChipsContract is a free data retrieval call binding the contract method 0xd13b19a3.
//
// Solidity: function chipsContract() view returns(address)
func (_StakingVSL *StakingVSLSession) ChipsContract() (common.Address, error) {
	return _StakingVSL.Contract.ChipsContract(&_StakingVSL.CallOpts)
}

// ChipsContract is a free data retrieval call binding the contract method 0xd13b19a3.
//
// Solidity: function chipsContract() view returns(address)
func (_StakingVSL *StakingVSLCallerSession) ChipsContract() (common.Address, error) {
	return _StakingVSL.Contract.ChipsContract(&_StakingVSL.CallOpts)
}

// GetChipsInfo is a free data retrieval call binding the contract method 0x90d3f47c.
//
// Solidity: function getChipsInfo(uint256 tokenId) view returns(address nodeAddr, uint256 tokens)
func (_StakingVSL *StakingVSLCaller) GetChipsInfo(opts *bind.CallOpts, tokenId *big.Int) (struct {
	NodeAddr common.Address
	Tokens   *big.Int
}, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getChipsInfo", tokenId)

	outstruct := new(struct {
		NodeAddr common.Address
		Tokens   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NodeAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Tokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetChipsInfo is a free data retrieval call binding the contract method 0x90d3f47c.
//
// Solidity: function getChipsInfo(uint256 tokenId) view returns(address nodeAddr, uint256 tokens)
func (_StakingVSL *StakingVSLSession) GetChipsInfo(tokenId *big.Int) (struct {
	NodeAddr common.Address
	Tokens   *big.Int
}, error) {
	return _StakingVSL.Contract.GetChipsInfo(&_StakingVSL.CallOpts, tokenId)
}

// GetChipsInfo is a free data retrieval call binding the contract method 0x90d3f47c.
//
// Solidity: function getChipsInfo(uint256 tokenId) view returns(address nodeAddr, uint256 tokens)
func (_StakingVSL *StakingVSLCallerSession) GetChipsInfo(tokenId *big.Int) (struct {
	NodeAddr common.Address
	Tokens   *big.Int
}, error) {
	return _StakingVSL.Contract.GetChipsInfo(&_StakingVSL.CallOpts, tokenId)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddr) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLCaller) GetNode(opts *bind.CallOpts, nodeAddr common.Address) (DataTypesNode, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getNode", nodeAddr)

	if err != nil {
		return *new(DataTypesNode), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNode)).(*DataTypesNode)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddr) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLSession) GetNode(nodeAddr common.Address) (DataTypesNode, error) {
	return _StakingVSL.Contract.GetNode(&_StakingVSL.CallOpts, nodeAddr)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address nodeAddr) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLCallerSession) GetNode(nodeAddr common.Address) (DataTypesNode, error) {
	return _StakingVSL.Contract.GetNode(&_StakingVSL.CallOpts, nodeAddr)
}

// GetNodeAvatar is a free data retrieval call binding the contract method 0x1474deaa.
//
// Solidity: function getNodeAvatar(address nodeAddr) view returns(string)
func (_StakingVSL *StakingVSLCaller) GetNodeAvatar(opts *bind.CallOpts, nodeAddr common.Address) (string, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getNodeAvatar", nodeAddr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNodeAvatar is a free data retrieval call binding the contract method 0x1474deaa.
//
// Solidity: function getNodeAvatar(address nodeAddr) view returns(string)
func (_StakingVSL *StakingVSLSession) GetNodeAvatar(nodeAddr common.Address) (string, error) {
	return _StakingVSL.Contract.GetNodeAvatar(&_StakingVSL.CallOpts, nodeAddr)
}

// GetNodeAvatar is a free data retrieval call binding the contract method 0x1474deaa.
//
// Solidity: function getNodeAvatar(address nodeAddr) view returns(string)
func (_StakingVSL *StakingVSLCallerSession) GetNodeAvatar(nodeAddr common.Address) (string, error) {
	return _StakingVSL.Contract.GetNodeAvatar(&_StakingVSL.CallOpts, nodeAddr)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_StakingVSL *StakingVSLCaller) GetNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_StakingVSL *StakingVSLSession) GetNodeCount() (*big.Int, error) {
	return _StakingVSL.Contract.GetNodeCount(&_StakingVSL.CallOpts)
}

// GetNodeCount is a free data retrieval call binding the contract method 0x39bf397e.
//
// Solidity: function getNodeCount() view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) GetNodeCount() (*big.Int, error) {
	return _StakingVSL.Contract.GetNodeCount(&_StakingVSL.CallOpts)
}

// GetNodes is a free data retrieval call binding the contract method 0x38c96b14.
//
// Solidity: function getNodes(address[] nodeAddrs) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLCaller) GetNodes(opts *bind.CallOpts, nodeAddrs []common.Address) ([]DataTypesNode, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getNodes", nodeAddrs)

	if err != nil {
		return *new([]DataTypesNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesNode)).(*[]DataTypesNode)

	return out0, err

}

// GetNodes is a free data retrieval call binding the contract method 0x38c96b14.
//
// Solidity: function getNodes(address[] nodeAddrs) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLSession) GetNodes(nodeAddrs []common.Address) ([]DataTypesNode, error) {
	return _StakingVSL.Contract.GetNodes(&_StakingVSL.CallOpts, nodeAddrs)
}

// GetNodes is a free data retrieval call binding the contract method 0x38c96b14.
//
// Solidity: function getNodes(address[] nodeAddrs) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLCallerSession) GetNodes(nodeAddrs []common.Address) ([]DataTypesNode, error) {
	return _StakingVSL.Contract.GetNodes(&_StakingVSL.CallOpts, nodeAddrs)
}

// GetNodesWithPagination is a free data retrieval call binding the contract method 0xd995415b.
//
// Solidity: function getNodesWithPagination(uint256 offset, uint256 limit) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLCaller) GetNodesWithPagination(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]DataTypesNode, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getNodesWithPagination", offset, limit)

	if err != nil {
		return *new([]DataTypesNode), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataTypesNode)).(*[]DataTypesNode)

	return out0, err

}

// GetNodesWithPagination is a free data retrieval call binding the contract method 0xd995415b.
//
// Solidity: function getNodesWithPagination(uint256 offset, uint256 limit) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLSession) GetNodesWithPagination(offset *big.Int, limit *big.Int) ([]DataTypesNode, error) {
	return _StakingVSL.Contract.GetNodesWithPagination(&_StakingVSL.CallOpts, offset, limit)
}

// GetNodesWithPagination is a free data retrieval call binding the contract method 0xd995415b.
//
// Solidity: function getNodesWithPagination(uint256 offset, uint256 limit) view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256)[] nodes)
func (_StakingVSL *StakingVSLCallerSession) GetNodesWithPagination(offset *big.Int, limit *big.Int) ([]DataTypesNode, error) {
	return _StakingVSL.Contract.GetNodesWithPagination(&_StakingVSL.CallOpts, offset, limit)
}

// GetPendingUnstake is a free data retrieval call binding the contract method 0xadfd065f.
//
// Solidity: function getPendingUnstake(uint256 requestId) view returns((address,address,uint256,uint256))
func (_StakingVSL *StakingVSLCaller) GetPendingUnstake(opts *bind.CallOpts, requestId *big.Int) (DataTypesUnstakeRequest, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getPendingUnstake", requestId)

	if err != nil {
		return *new(DataTypesUnstakeRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUnstakeRequest)).(*DataTypesUnstakeRequest)

	return out0, err

}

// GetPendingUnstake is a free data retrieval call binding the contract method 0xadfd065f.
//
// Solidity: function getPendingUnstake(uint256 requestId) view returns((address,address,uint256,uint256))
func (_StakingVSL *StakingVSLSession) GetPendingUnstake(requestId *big.Int) (DataTypesUnstakeRequest, error) {
	return _StakingVSL.Contract.GetPendingUnstake(&_StakingVSL.CallOpts, requestId)
}

// GetPendingUnstake is a free data retrieval call binding the contract method 0xadfd065f.
//
// Solidity: function getPendingUnstake(uint256 requestId) view returns((address,address,uint256,uint256))
func (_StakingVSL *StakingVSLCallerSession) GetPendingUnstake(requestId *big.Int) (DataTypesUnstakeRequest, error) {
	return _StakingVSL.Contract.GetPendingUnstake(&_StakingVSL.CallOpts, requestId)
}

// GetPendingWithdrawal is a free data retrieval call binding the contract method 0x38a3c878.
//
// Solidity: function getPendingWithdrawal(uint256 requestId) view returns((address,uint40,uint256))
func (_StakingVSL *StakingVSLCaller) GetPendingWithdrawal(opts *bind.CallOpts, requestId *big.Int) (DataTypesWithdrawalRequest, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getPendingWithdrawal", requestId)

	if err != nil {
		return *new(DataTypesWithdrawalRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesWithdrawalRequest)).(*DataTypesWithdrawalRequest)

	return out0, err

}

// GetPendingWithdrawal is a free data retrieval call binding the contract method 0x38a3c878.
//
// Solidity: function getPendingWithdrawal(uint256 requestId) view returns((address,uint40,uint256))
func (_StakingVSL *StakingVSLSession) GetPendingWithdrawal(requestId *big.Int) (DataTypesWithdrawalRequest, error) {
	return _StakingVSL.Contract.GetPendingWithdrawal(&_StakingVSL.CallOpts, requestId)
}

// GetPendingWithdrawal is a free data retrieval call binding the contract method 0x38a3c878.
//
// Solidity: function getPendingWithdrawal(uint256 requestId) view returns((address,uint40,uint256))
func (_StakingVSL *StakingVSLCallerSession) GetPendingWithdrawal(requestId *big.Int) (DataTypesWithdrawalRequest, error) {
	return _StakingVSL.Contract.GetPendingWithdrawal(&_StakingVSL.CallOpts, requestId)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns(uint256 totalOperationPoolTokens, uint256 totalStakingPoolTokens)
func (_StakingVSL *StakingVSLCaller) GetPoolInfo(opts *bind.CallOpts) (struct {
	TotalOperationPoolTokens *big.Int
	TotalStakingPoolTokens   *big.Int
}, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getPoolInfo")

	outstruct := new(struct {
		TotalOperationPoolTokens *big.Int
		TotalStakingPoolTokens   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalOperationPoolTokens = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStakingPoolTokens = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns(uint256 totalOperationPoolTokens, uint256 totalStakingPoolTokens)
func (_StakingVSL *StakingVSLSession) GetPoolInfo() (struct {
	TotalOperationPoolTokens *big.Int
	TotalStakingPoolTokens   *big.Int
}, error) {
	return _StakingVSL.Contract.GetPoolInfo(&_StakingVSL.CallOpts)
}

// GetPoolInfo is a free data retrieval call binding the contract method 0x60246c88.
//
// Solidity: function getPoolInfo() view returns(uint256 totalOperationPoolTokens, uint256 totalStakingPoolTokens)
func (_StakingVSL *StakingVSLCallerSession) GetPoolInfo() (struct {
	TotalOperationPoolTokens *big.Int
	TotalStakingPoolTokens   *big.Int
}, error) {
	return _StakingVSL.Contract.GetPoolInfo(&_StakingVSL.CallOpts)
}

// GetPublicPool is a free data retrieval call binding the contract method 0xc84c42a3.
//
// Solidity: function getPublicPool() view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLCaller) GetPublicPool(opts *bind.CallOpts) (DataTypesNode, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getPublicPool")

	if err != nil {
		return *new(DataTypesNode), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesNode)).(*DataTypesNode)

	return out0, err

}

// GetPublicPool is a free data retrieval call binding the contract method 0xc84c42a3.
//
// Solidity: function getPublicPool() view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLSession) GetPublicPool() (DataTypesNode, error) {
	return _StakingVSL.Contract.GetPublicPool(&_StakingVSL.CallOpts)
}

// GetPublicPool is a free data retrieval call binding the contract method 0xc84c42a3.
//
// Solidity: function getPublicPool() view returns((uint256,address,uint64,bool,bool,string,string,uint256,uint256,uint256,uint256))
func (_StakingVSL *StakingVSLCallerSession) GetPublicPool() (DataTypesNode, error) {
	return _StakingVSL.Contract.GetPublicPool(&_StakingVSL.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingVSL *StakingVSLCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingVSL *StakingVSLSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingVSL.Contract.GetRoleAdmin(&_StakingVSL.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingVSL *StakingVSLCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingVSL.Contract.GetRoleAdmin(&_StakingVSL.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingVSL *StakingVSLCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingVSL *StakingVSLSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingVSL.Contract.GetRoleMember(&_StakingVSL.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakingVSL *StakingVSLCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakingVSL.Contract.GetRoleMember(&_StakingVSL.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingVSL *StakingVSLCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingVSL *StakingVSLSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingVSL.Contract.GetRoleMemberCount(&_StakingVSL.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakingVSL.Contract.GetRoleMemberCount(&_StakingVSL.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingVSL *StakingVSLCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingVSL *StakingVSLSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingVSL.Contract.HasRole(&_StakingVSL.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingVSL *StakingVSLCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingVSL.Contract.HasRole(&_StakingVSL.CallOpts, role, account)
}

// IsAlphaPhase is a free data retrieval call binding the contract method 0x69ff71f6.
//
// Solidity: function isAlphaPhase() view returns(bool)
func (_StakingVSL *StakingVSLCaller) IsAlphaPhase(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "isAlphaPhase")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAlphaPhase is a free data retrieval call binding the contract method 0x69ff71f6.
//
// Solidity: function isAlphaPhase() view returns(bool)
func (_StakingVSL *StakingVSLSession) IsAlphaPhase() (bool, error) {
	return _StakingVSL.Contract.IsAlphaPhase(&_StakingVSL.CallOpts)
}

// IsAlphaPhase is a free data retrieval call binding the contract method 0x69ff71f6.
//
// Solidity: function isAlphaPhase() view returns(bool)
func (_StakingVSL *StakingVSLCallerSession) IsAlphaPhase() (bool, error) {
	return _StakingVSL.Contract.IsAlphaPhase(&_StakingVSL.CallOpts)
}

// IsSettlementPhase is a free data retrieval call binding the contract method 0x2e75fd59.
//
// Solidity: function isSettlementPhase() view returns(bool)
func (_StakingVSL *StakingVSLCaller) IsSettlementPhase(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "isSettlementPhase")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSettlementPhase is a free data retrieval call binding the contract method 0x2e75fd59.
//
// Solidity: function isSettlementPhase() view returns(bool)
func (_StakingVSL *StakingVSLSession) IsSettlementPhase() (bool, error) {
	return _StakingVSL.Contract.IsSettlementPhase(&_StakingVSL.CallOpts)
}

// IsSettlementPhase is a free data retrieval call binding the contract method 0x2e75fd59.
//
// Solidity: function isSettlementPhase() view returns(bool)
func (_StakingVSL *StakingVSLCallerSession) IsSettlementPhase() (bool, error) {
	return _StakingVSL.Contract.IsSettlementPhase(&_StakingVSL.CallOpts)
}

// MinTokensToStake is a free data retrieval call binding the contract method 0x14936b13.
//
// Solidity: function minTokensToStake(address nodeAddr) view returns(uint256)
func (_StakingVSL *StakingVSLCaller) MinTokensToStake(opts *bind.CallOpts, nodeAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "minTokensToStake", nodeAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTokensToStake is a free data retrieval call binding the contract method 0x14936b13.
//
// Solidity: function minTokensToStake(address nodeAddr) view returns(uint256)
func (_StakingVSL *StakingVSLSession) MinTokensToStake(nodeAddr common.Address) (*big.Int, error) {
	return _StakingVSL.Contract.MinTokensToStake(&_StakingVSL.CallOpts, nodeAddr)
}

// MinTokensToStake is a free data retrieval call binding the contract method 0x14936b13.
//
// Solidity: function minTokensToStake(address nodeAddr) view returns(uint256)
func (_StakingVSL *StakingVSLCallerSession) MinTokensToStake(nodeAddr common.Address) (*big.Int, error) {
	return _StakingVSL.Contract.MinTokensToStake(&_StakingVSL.CallOpts, nodeAddr)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingVSL *StakingVSLCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingVSL *StakingVSLSession) Paused() (bool, error) {
	return _StakingVSL.Contract.Paused(&_StakingVSL.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingVSL *StakingVSLCallerSession) Paused() (bool, error) {
	return _StakingVSL.Contract.Paused(&_StakingVSL.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingVSL *StakingVSLCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingVSL.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingVSL *StakingVSLSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingVSL.Contract.SupportsInterface(&_StakingVSL.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingVSL *StakingVSLCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingVSL.Contract.SupportsInterface(&_StakingVSL.CallOpts, interfaceId)
}

// ClaimUnstake is a paid mutator transaction binding the contract method 0x04a4fb10.
//
// Solidity: function claimUnstake(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLTransactor) ClaimUnstake(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "claimUnstake", requestIds)
}

// ClaimUnstake is a paid mutator transaction binding the contract method 0x04a4fb10.
//
// Solidity: function claimUnstake(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLSession) ClaimUnstake(requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.ClaimUnstake(&_StakingVSL.TransactOpts, requestIds)
}

// ClaimUnstake is a paid mutator transaction binding the contract method 0x04a4fb10.
//
// Solidity: function claimUnstake(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLTransactorSession) ClaimUnstake(requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.ClaimUnstake(&_StakingVSL.TransactOpts, requestIds)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0x3c256b98.
//
// Solidity: function claimWithdrawal(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLTransactor) ClaimWithdrawal(opts *bind.TransactOpts, requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "claimWithdrawal", requestIds)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0x3c256b98.
//
// Solidity: function claimWithdrawal(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLSession) ClaimWithdrawal(requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.ClaimWithdrawal(&_StakingVSL.TransactOpts, requestIds)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0x3c256b98.
//
// Solidity: function claimWithdrawal(uint256[] requestIds) returns()
func (_StakingVSL *StakingVSLTransactorSession) ClaimWithdrawal(requestIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.ClaimWithdrawal(&_StakingVSL.TransactOpts, requestIds)
}

// CreateNode is a paid mutator transaction binding the contract method 0x96531623.
//
// Solidity: function createNode(string name, string description, uint64 taxRateBasisPoints, bool publicGood) payable returns()
func (_StakingVSL *StakingVSLTransactor) CreateNode(opts *bind.TransactOpts, name string, description string, taxRateBasisPoints uint64, publicGood bool) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "createNode", name, description, taxRateBasisPoints, publicGood)
}

// CreateNode is a paid mutator transaction binding the contract method 0x96531623.
//
// Solidity: function createNode(string name, string description, uint64 taxRateBasisPoints, bool publicGood) payable returns()
func (_StakingVSL *StakingVSLSession) CreateNode(name string, description string, taxRateBasisPoints uint64, publicGood bool) (*types.Transaction, error) {
	return _StakingVSL.Contract.CreateNode(&_StakingVSL.TransactOpts, name, description, taxRateBasisPoints, publicGood)
}

// CreateNode is a paid mutator transaction binding the contract method 0x96531623.
//
// Solidity: function createNode(string name, string description, uint64 taxRateBasisPoints, bool publicGood) payable returns()
func (_StakingVSL *StakingVSLTransactorSession) CreateNode(name string, description string, taxRateBasisPoints uint64, publicGood bool) (*types.Transaction, error) {
	return _StakingVSL.Contract.CreateNode(&_StakingVSL.TransactOpts, name, description, taxRateBasisPoints, publicGood)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingVSL *StakingVSLTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingVSL *StakingVSLSession) Deposit() (*types.Transaction, error) {
	return _StakingVSL.Contract.Deposit(&_StakingVSL.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_StakingVSL *StakingVSLTransactorSession) Deposit() (*types.Transaction, error) {
	return _StakingVSL.Contract.Deposit(&_StakingVSL.TransactOpts)
}

// DisableAlphaPhase is a paid mutator transaction binding the contract method 0x81001e60.
//
// Solidity: function disableAlphaPhase() returns()
func (_StakingVSL *StakingVSLTransactor) DisableAlphaPhase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "disableAlphaPhase")
}

// DisableAlphaPhase is a paid mutator transaction binding the contract method 0x81001e60.
//
// Solidity: function disableAlphaPhase() returns()
func (_StakingVSL *StakingVSLSession) DisableAlphaPhase() (*types.Transaction, error) {
	return _StakingVSL.Contract.DisableAlphaPhase(&_StakingVSL.TransactOpts)
}

// DisableAlphaPhase is a paid mutator transaction binding the contract method 0x81001e60.
//
// Solidity: function disableAlphaPhase() returns()
func (_StakingVSL *StakingVSLTransactorSession) DisableAlphaPhase() (*types.Transaction, error) {
	return _StakingVSL.Contract.DisableAlphaPhase(&_StakingVSL.TransactOpts)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x7a142d5e.
//
// Solidity: function distributeRewards(uint256[3] epochInfo, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256 publicPoolRewards) payable returns()
func (_StakingVSL *StakingVSLTransactor) DistributeRewards(opts *bind.TransactOpts, epochInfo [3]*big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, stakingRewards []*big.Int, publicPoolRewards *big.Int) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "distributeRewards", epochInfo, nodeAddrs, operationRewards, stakingRewards, publicPoolRewards)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x7a142d5e.
//
// Solidity: function distributeRewards(uint256[3] epochInfo, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256 publicPoolRewards) payable returns()
func (_StakingVSL *StakingVSLSession) DistributeRewards(epochInfo [3]*big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, stakingRewards []*big.Int, publicPoolRewards *big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.DistributeRewards(&_StakingVSL.TransactOpts, epochInfo, nodeAddrs, operationRewards, stakingRewards, publicPoolRewards)
}

// DistributeRewards is a paid mutator transaction binding the contract method 0x7a142d5e.
//
// Solidity: function distributeRewards(uint256[3] epochInfo, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256 publicPoolRewards) payable returns()
func (_StakingVSL *StakingVSLTransactorSession) DistributeRewards(epochInfo [3]*big.Int, nodeAddrs []common.Address, operationRewards []*big.Int, stakingRewards []*big.Int, publicPoolRewards *big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.DistributeRewards(&_StakingVSL.TransactOpts, epochInfo, nodeAddrs, operationRewards, stakingRewards, publicPoolRewards)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.GrantRole(&_StakingVSL.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.GrantRole(&_StakingVSL.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address chips, address pauseAccount, address oracleAccount) returns()
func (_StakingVSL *StakingVSLTransactor) Initialize(opts *bind.TransactOpts, chips common.Address, pauseAccount common.Address, oracleAccount common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "initialize", chips, pauseAccount, oracleAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address chips, address pauseAccount, address oracleAccount) returns()
func (_StakingVSL *StakingVSLSession) Initialize(chips common.Address, pauseAccount common.Address, oracleAccount common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.Initialize(&_StakingVSL.TransactOpts, chips, pauseAccount, oracleAccount)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address chips, address pauseAccount, address oracleAccount) returns()
func (_StakingVSL *StakingVSLTransactorSession) Initialize(chips common.Address, pauseAccount common.Address, oracleAccount common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.Initialize(&_StakingVSL.TransactOpts, chips, pauseAccount, oracleAccount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingVSL *StakingVSLTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingVSL *StakingVSLSession) Pause() (*types.Transaction, error) {
	return _StakingVSL.Contract.Pause(&_StakingVSL.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingVSL *StakingVSLTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingVSL.Contract.Pause(&_StakingVSL.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingVSL *StakingVSLTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingVSL *StakingVSLSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.RenounceRole(&_StakingVSL.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingVSL *StakingVSLTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.RenounceRole(&_StakingVSL.TransactOpts, role, callerConfirmation)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0xbcdd4190.
//
// Solidity: function requestUnstake(address nodeAddr, uint256[] chipsIds) returns(uint256 requestId)
func (_StakingVSL *StakingVSLTransactor) RequestUnstake(opts *bind.TransactOpts, nodeAddr common.Address, chipsIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "requestUnstake", nodeAddr, chipsIds)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0xbcdd4190.
//
// Solidity: function requestUnstake(address nodeAddr, uint256[] chipsIds) returns(uint256 requestId)
func (_StakingVSL *StakingVSLSession) RequestUnstake(nodeAddr common.Address, chipsIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.RequestUnstake(&_StakingVSL.TransactOpts, nodeAddr, chipsIds)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0xbcdd4190.
//
// Solidity: function requestUnstake(address nodeAddr, uint256[] chipsIds) returns(uint256 requestId)
func (_StakingVSL *StakingVSLTransactorSession) RequestUnstake(nodeAddr common.Address, chipsIds []*big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.RequestUnstake(&_StakingVSL.TransactOpts, nodeAddr, chipsIds)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns(uint256 requestId)
func (_StakingVSL *StakingVSLTransactor) RequestWithdrawal(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "requestWithdrawal", amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns(uint256 requestId)
func (_StakingVSL *StakingVSLSession) RequestWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.RequestWithdrawal(&_StakingVSL.TransactOpts, amount)
}

// RequestWithdrawal is a paid mutator transaction binding the contract method 0x9ee679e8.
//
// Solidity: function requestWithdrawal(uint256 amount) returns(uint256 requestId)
func (_StakingVSL *StakingVSLTransactorSession) RequestWithdrawal(amount *big.Int) (*types.Transaction, error) {
	return _StakingVSL.Contract.RequestWithdrawal(&_StakingVSL.TransactOpts, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.RevokeRole(&_StakingVSL.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingVSL *StakingVSLTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.RevokeRole(&_StakingVSL.TransactOpts, role, account)
}

// SetSettlementPhase is a paid mutator transaction binding the contract method 0x4e7a1286.
//
// Solidity: function setSettlementPhase(bool enabled) returns()
func (_StakingVSL *StakingVSLTransactor) SetSettlementPhase(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "setSettlementPhase", enabled)
}

// SetSettlementPhase is a paid mutator transaction binding the contract method 0x4e7a1286.
//
// Solidity: function setSettlementPhase(bool enabled) returns()
func (_StakingVSL *StakingVSLSession) SetSettlementPhase(enabled bool) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetSettlementPhase(&_StakingVSL.TransactOpts, enabled)
}

// SetSettlementPhase is a paid mutator transaction binding the contract method 0x4e7a1286.
//
// Solidity: function setSettlementPhase(bool enabled) returns()
func (_StakingVSL *StakingVSLTransactorSession) SetSettlementPhase(enabled bool) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetSettlementPhase(&_StakingVSL.TransactOpts, enabled)
}

// SetTaxRateBasisPoints4Node is a paid mutator transaction binding the contract method 0xc7057c1f.
//
// Solidity: function setTaxRateBasisPoints4Node(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLTransactor) SetTaxRateBasisPoints4Node(opts *bind.TransactOpts, taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "setTaxRateBasisPoints4Node", taxRateBasisPoints)
}

// SetTaxRateBasisPoints4Node is a paid mutator transaction binding the contract method 0xc7057c1f.
//
// Solidity: function setTaxRateBasisPoints4Node(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLSession) SetTaxRateBasisPoints4Node(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetTaxRateBasisPoints4Node(&_StakingVSL.TransactOpts, taxRateBasisPoints)
}

// SetTaxRateBasisPoints4Node is a paid mutator transaction binding the contract method 0xc7057c1f.
//
// Solidity: function setTaxRateBasisPoints4Node(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLTransactorSession) SetTaxRateBasisPoints4Node(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetTaxRateBasisPoints4Node(&_StakingVSL.TransactOpts, taxRateBasisPoints)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLTransactor) SetTaxRateBasisPoints4PublicPool(opts *bind.TransactOpts, taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "setTaxRateBasisPoints4PublicPool", taxRateBasisPoints)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLSession) SetTaxRateBasisPoints4PublicPool(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetTaxRateBasisPoints4PublicPool(&_StakingVSL.TransactOpts, taxRateBasisPoints)
}

// SetTaxRateBasisPoints4PublicPool is a paid mutator transaction binding the contract method 0xe3fb8dca.
//
// Solidity: function setTaxRateBasisPoints4PublicPool(uint64 taxRateBasisPoints) returns()
func (_StakingVSL *StakingVSLTransactorSession) SetTaxRateBasisPoints4PublicPool(taxRateBasisPoints uint64) (*types.Transaction, error) {
	return _StakingVSL.Contract.SetTaxRateBasisPoints4PublicPool(&_StakingVSL.TransactOpts, taxRateBasisPoints)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_StakingVSL *StakingVSLTransactor) SlashNodes(opts *bind.TransactOpts, nodeAddrs []common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "slashNodes", nodeAddrs)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_StakingVSL *StakingVSLSession) SlashNodes(nodeAddrs []common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.SlashNodes(&_StakingVSL.TransactOpts, nodeAddrs)
}

// SlashNodes is a paid mutator transaction binding the contract method 0xa2f641c3.
//
// Solidity: function slashNodes(address[] nodeAddrs) returns()
func (_StakingVSL *StakingVSLTransactorSession) SlashNodes(nodeAddrs []common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.SlashNodes(&_StakingVSL.TransactOpts, nodeAddrs)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLTransactor) Stake(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "stake", nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.Stake(&_StakingVSL.TransactOpts, nodeAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLTransactorSession) Stake(nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.Stake(&_StakingVSL.TransactOpts, nodeAddr)
}

// StakeToPublicPool is a paid mutator transaction binding the contract method 0x379f8100.
//
// Solidity: function stakeToPublicPool(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLTransactor) StakeToPublicPool(opts *bind.TransactOpts, nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "stakeToPublicPool", nodeAddr)
}

// StakeToPublicPool is a paid mutator transaction binding the contract method 0x379f8100.
//
// Solidity: function stakeToPublicPool(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLSession) StakeToPublicPool(nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.StakeToPublicPool(&_StakingVSL.TransactOpts, nodeAddr)
}

// StakeToPublicPool is a paid mutator transaction binding the contract method 0x379f8100.
//
// Solidity: function stakeToPublicPool(address nodeAddr) payable returns(uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLTransactorSession) StakeToPublicPool(nodeAddr common.Address) (*types.Transaction, error) {
	return _StakingVSL.Contract.StakeToPublicPool(&_StakingVSL.TransactOpts, nodeAddr)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingVSL *StakingVSLTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingVSL *StakingVSLSession) Unpause() (*types.Transaction, error) {
	return _StakingVSL.Contract.Unpause(&_StakingVSL.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingVSL *StakingVSLTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingVSL.Contract.Unpause(&_StakingVSL.TransactOpts)
}

// UpdateToPublicGood is a paid mutator transaction binding the contract method 0xc9af094c.
//
// Solidity: function updateToPublicGood() returns()
func (_StakingVSL *StakingVSLTransactor) UpdateToPublicGood(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "updateToPublicGood")
}

// UpdateToPublicGood is a paid mutator transaction binding the contract method 0xc9af094c.
//
// Solidity: function updateToPublicGood() returns()
func (_StakingVSL *StakingVSLSession) UpdateToPublicGood() (*types.Transaction, error) {
	return _StakingVSL.Contract.UpdateToPublicGood(&_StakingVSL.TransactOpts)
}

// UpdateToPublicGood is a paid mutator transaction binding the contract method 0xc9af094c.
//
// Solidity: function updateToPublicGood() returns()
func (_StakingVSL *StakingVSLTransactorSession) UpdateToPublicGood() (*types.Transaction, error) {
	return _StakingVSL.Contract.UpdateToPublicGood(&_StakingVSL.TransactOpts)
}

// Withdraw2Treasury is a paid mutator transaction binding the contract method 0x4a7dfc90.
//
// Solidity: function withdraw2Treasury() returns()
func (_StakingVSL *StakingVSLTransactor) Withdraw2Treasury(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingVSL.contract.Transact(opts, "withdraw2Treasury")
}

// Withdraw2Treasury is a paid mutator transaction binding the contract method 0x4a7dfc90.
//
// Solidity: function withdraw2Treasury() returns()
func (_StakingVSL *StakingVSLSession) Withdraw2Treasury() (*types.Transaction, error) {
	return _StakingVSL.Contract.Withdraw2Treasury(&_StakingVSL.TransactOpts)
}

// Withdraw2Treasury is a paid mutator transaction binding the contract method 0x4a7dfc90.
//
// Solidity: function withdraw2Treasury() returns()
func (_StakingVSL *StakingVSLTransactorSession) Withdraw2Treasury() (*types.Transaction, error) {
	return _StakingVSL.Contract.Withdraw2Treasury(&_StakingVSL.TransactOpts)
}

// StakingVSLDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the StakingVSL contract.
type StakingVSLDepositedIterator struct {
	Event *StakingVSLDeposited // Event containing the contract specifics and raw log

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
func (it *StakingVSLDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLDeposited)
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
		it.Event = new(StakingVSLDeposited)
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
func (it *StakingVSLDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLDeposited represents a Deposited event raised by the StakingVSL contract.
type StakingVSLDeposited struct {
	NodeAddr common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed nodeAddr, uint256 indexed amount)
func (_StakingVSL *StakingVSLFilterer) FilterDeposited(opts *bind.FilterOpts, nodeAddr []common.Address, amount []*big.Int) (*StakingVSLDepositedIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "Deposited", nodeAddrRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLDepositedIterator{contract: _StakingVSL.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed nodeAddr, uint256 indexed amount)
func (_StakingVSL *StakingVSLFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *StakingVSLDeposited, nodeAddr []common.Address, amount []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "Deposited", nodeAddrRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLDeposited)
				if err := _StakingVSL.contract.UnpackLog(event, "Deposited", log); err != nil {
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

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address indexed nodeAddr, uint256 indexed amount)
func (_StakingVSL *StakingVSLFilterer) ParseDeposited(log types.Log) (*StakingVSLDeposited, error) {
	event := new(StakingVSLDeposited)
	if err := _StakingVSL.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingVSL contract.
type StakingVSLInitializedIterator struct {
	Event *StakingVSLInitialized // Event containing the contract specifics and raw log

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
func (it *StakingVSLInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLInitialized)
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
		it.Event = new(StakingVSLInitialized)
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
func (it *StakingVSLInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLInitialized represents a Initialized event raised by the StakingVSL contract.
type StakingVSLInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingVSL *StakingVSLFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingVSLInitializedIterator, error) {

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingVSLInitializedIterator{contract: _StakingVSL.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingVSL *StakingVSLFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingVSLInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLInitialized)
				if err := _StakingVSL.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingVSL *StakingVSLFilterer) ParseInitialized(log types.Log) (*StakingVSLInitialized, error) {
	event := new(StakingVSLInitialized)
	if err := _StakingVSL.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the StakingVSL contract.
type StakingVSLNodeCreatedIterator struct {
	Event *StakingVSLNodeCreated // Event containing the contract specifics and raw log

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
func (it *StakingVSLNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLNodeCreated)
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
		it.Event = new(StakingVSLNodeCreated)
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
func (it *StakingVSLNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLNodeCreated represents a NodeCreated event raised by the StakingVSL contract.
type StakingVSLNodeCreated struct {
	NodeId             *big.Int
	NodeAddr           common.Address
	Name               string
	Description        string
	TaxRateBasisPoints uint64
	PublicGood         bool
	Alpha              bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x37570f68d94fd46cd4009b3823da2b2bc1a9a7e38f824f311ede9e876816e321.
//
// Solidity: event NodeCreated(uint256 indexed nodeId, address indexed nodeAddr, string name, string description, uint64 taxRateBasisPoints, bool publicGood, bool alpha)
func (_StakingVSL *StakingVSLFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeId []*big.Int, nodeAddr []common.Address) (*StakingVSLNodeCreatedIterator, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "NodeCreated", nodeIdRule, nodeAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLNodeCreatedIterator{contract: _StakingVSL.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x37570f68d94fd46cd4009b3823da2b2bc1a9a7e38f824f311ede9e876816e321.
//
// Solidity: event NodeCreated(uint256 indexed nodeId, address indexed nodeAddr, string name, string description, uint64 taxRateBasisPoints, bool publicGood, bool alpha)
func (_StakingVSL *StakingVSLFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *StakingVSLNodeCreated, nodeId []*big.Int, nodeAddr []common.Address) (event.Subscription, error) {

	var nodeIdRule []interface{}
	for _, nodeIdItem := range nodeId {
		nodeIdRule = append(nodeIdRule, nodeIdItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "NodeCreated", nodeIdRule, nodeAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLNodeCreated)
				if err := _StakingVSL.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x37570f68d94fd46cd4009b3823da2b2bc1a9a7e38f824f311ede9e876816e321.
//
// Solidity: event NodeCreated(uint256 indexed nodeId, address indexed nodeAddr, string name, string description, uint64 taxRateBasisPoints, bool publicGood, bool alpha)
func (_StakingVSL *StakingVSLFilterer) ParseNodeCreated(log types.Log) (*StakingVSLNodeCreated, error) {
	event := new(StakingVSLNodeCreated)
	if err := _StakingVSL.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLNodeSlashedIterator is returned from FilterNodeSlashed and is used to iterate over the raw logs and unpacked data for NodeSlashed events raised by the StakingVSL contract.
type StakingVSLNodeSlashedIterator struct {
	Event *StakingVSLNodeSlashed // Event containing the contract specifics and raw log

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
func (it *StakingVSLNodeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLNodeSlashed)
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
		it.Event = new(StakingVSLNodeSlashed)
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
func (it *StakingVSLNodeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLNodeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLNodeSlashed represents a NodeSlashed event raised by the StakingVSL contract.
type StakingVSLNodeSlashed struct {
	NodeAddr             common.Address
	SlashedOperationPool *big.Int
	SlashedStakingPool   *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNodeSlashed is a free log retrieval operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeAddr, uint256 indexed slashedOperationPool, uint256 indexed slashedStakingPool)
func (_StakingVSL *StakingVSLFilterer) FilterNodeSlashed(opts *bind.FilterOpts, nodeAddr []common.Address, slashedOperationPool []*big.Int, slashedStakingPool []*big.Int) (*StakingVSLNodeSlashedIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var slashedOperationPoolRule []interface{}
	for _, slashedOperationPoolItem := range slashedOperationPool {
		slashedOperationPoolRule = append(slashedOperationPoolRule, slashedOperationPoolItem)
	}
	var slashedStakingPoolRule []interface{}
	for _, slashedStakingPoolItem := range slashedStakingPool {
		slashedStakingPoolRule = append(slashedStakingPoolRule, slashedStakingPoolItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "NodeSlashed", nodeAddrRule, slashedOperationPoolRule, slashedStakingPoolRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLNodeSlashedIterator{contract: _StakingVSL.contract, event: "NodeSlashed", logs: logs, sub: sub}, nil
}

// WatchNodeSlashed is a free log subscription operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeAddr, uint256 indexed slashedOperationPool, uint256 indexed slashedStakingPool)
func (_StakingVSL *StakingVSLFilterer) WatchNodeSlashed(opts *bind.WatchOpts, sink chan<- *StakingVSLNodeSlashed, nodeAddr []common.Address, slashedOperationPool []*big.Int, slashedStakingPool []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var slashedOperationPoolRule []interface{}
	for _, slashedOperationPoolItem := range slashedOperationPool {
		slashedOperationPoolRule = append(slashedOperationPoolRule, slashedOperationPoolItem)
	}
	var slashedStakingPoolRule []interface{}
	for _, slashedStakingPoolItem := range slashedStakingPool {
		slashedStakingPoolRule = append(slashedStakingPoolRule, slashedStakingPoolItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "NodeSlashed", nodeAddrRule, slashedOperationPoolRule, slashedStakingPoolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLNodeSlashed)
				if err := _StakingVSL.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
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

// ParseNodeSlashed is a log parse operation binding the contract event 0xa8d720d0a0a2e7c96bf9eb87433901ebb6331356c8f3283b2568de34478703cc.
//
// Solidity: event NodeSlashed(address indexed nodeAddr, uint256 indexed slashedOperationPool, uint256 indexed slashedStakingPool)
func (_StakingVSL *StakingVSLFilterer) ParseNodeSlashed(log types.Log) (*StakingVSLNodeSlashed, error) {
	event := new(StakingVSLNodeSlashed)
	if err := _StakingVSL.contract.UnpackLog(event, "NodeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLNodeTaxRateBasisPointsSetIterator is returned from FilterNodeTaxRateBasisPointsSet and is used to iterate over the raw logs and unpacked data for NodeTaxRateBasisPointsSet events raised by the StakingVSL contract.
type StakingVSLNodeTaxRateBasisPointsSetIterator struct {
	Event *StakingVSLNodeTaxRateBasisPointsSet // Event containing the contract specifics and raw log

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
func (it *StakingVSLNodeTaxRateBasisPointsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLNodeTaxRateBasisPointsSet)
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
		it.Event = new(StakingVSLNodeTaxRateBasisPointsSet)
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
func (it *StakingVSLNodeTaxRateBasisPointsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLNodeTaxRateBasisPointsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLNodeTaxRateBasisPointsSet represents a NodeTaxRateBasisPointsSet event raised by the StakingVSL contract.
type StakingVSLNodeTaxRateBasisPointsSet struct {
	NodeAddr           common.Address
	TaxRateBasisPoints uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeTaxRateBasisPointsSet is a free log retrieval operation binding the contract event 0xb8e5551053b871a40f7c7382e5bd3af5a62dd737d059d3838cf3aa7c325bd479.
//
// Solidity: event NodeTaxRateBasisPointsSet(address indexed nodeAddr, uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) FilterNodeTaxRateBasisPointsSet(opts *bind.FilterOpts, nodeAddr []common.Address, taxRateBasisPoints []uint64) (*StakingVSLNodeTaxRateBasisPointsSetIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var taxRateBasisPointsRule []interface{}
	for _, taxRateBasisPointsItem := range taxRateBasisPoints {
		taxRateBasisPointsRule = append(taxRateBasisPointsRule, taxRateBasisPointsItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "NodeTaxRateBasisPointsSet", nodeAddrRule, taxRateBasisPointsRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLNodeTaxRateBasisPointsSetIterator{contract: _StakingVSL.contract, event: "NodeTaxRateBasisPointsSet", logs: logs, sub: sub}, nil
}

// WatchNodeTaxRateBasisPointsSet is a free log subscription operation binding the contract event 0xb8e5551053b871a40f7c7382e5bd3af5a62dd737d059d3838cf3aa7c325bd479.
//
// Solidity: event NodeTaxRateBasisPointsSet(address indexed nodeAddr, uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) WatchNodeTaxRateBasisPointsSet(opts *bind.WatchOpts, sink chan<- *StakingVSLNodeTaxRateBasisPointsSet, nodeAddr []common.Address, taxRateBasisPoints []uint64) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var taxRateBasisPointsRule []interface{}
	for _, taxRateBasisPointsItem := range taxRateBasisPoints {
		taxRateBasisPointsRule = append(taxRateBasisPointsRule, taxRateBasisPointsItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "NodeTaxRateBasisPointsSet", nodeAddrRule, taxRateBasisPointsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLNodeTaxRateBasisPointsSet)
				if err := _StakingVSL.contract.UnpackLog(event, "NodeTaxRateBasisPointsSet", log); err != nil {
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

// ParseNodeTaxRateBasisPointsSet is a log parse operation binding the contract event 0xb8e5551053b871a40f7c7382e5bd3af5a62dd737d059d3838cf3aa7c325bd479.
//
// Solidity: event NodeTaxRateBasisPointsSet(address indexed nodeAddr, uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) ParseNodeTaxRateBasisPointsSet(log types.Log) (*StakingVSLNodeTaxRateBasisPointsSet, error) {
	event := new(StakingVSLNodeTaxRateBasisPointsSet)
	if err := _StakingVSL.contract.UnpackLog(event, "NodeTaxRateBasisPointsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLNodeUpdated2PublicGoodIterator is returned from FilterNodeUpdated2PublicGood and is used to iterate over the raw logs and unpacked data for NodeUpdated2PublicGood events raised by the StakingVSL contract.
type StakingVSLNodeUpdated2PublicGoodIterator struct {
	Event *StakingVSLNodeUpdated2PublicGood // Event containing the contract specifics and raw log

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
func (it *StakingVSLNodeUpdated2PublicGoodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLNodeUpdated2PublicGood)
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
		it.Event = new(StakingVSLNodeUpdated2PublicGood)
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
func (it *StakingVSLNodeUpdated2PublicGoodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLNodeUpdated2PublicGoodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLNodeUpdated2PublicGood represents a NodeUpdated2PublicGood event raised by the StakingVSL contract.
type StakingVSLNodeUpdated2PublicGood struct {
	NodeAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated2PublicGood is a free log retrieval operation binding the contract event 0x86538b79bef9c52dbe4e888742cfbd70114655c47ef30ede4997791fe79a9376.
//
// Solidity: event NodeUpdated2PublicGood(address indexed nodeAddr)
func (_StakingVSL *StakingVSLFilterer) FilterNodeUpdated2PublicGood(opts *bind.FilterOpts, nodeAddr []common.Address) (*StakingVSLNodeUpdated2PublicGoodIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "NodeUpdated2PublicGood", nodeAddrRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLNodeUpdated2PublicGoodIterator{contract: _StakingVSL.contract, event: "NodeUpdated2PublicGood", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated2PublicGood is a free log subscription operation binding the contract event 0x86538b79bef9c52dbe4e888742cfbd70114655c47ef30ede4997791fe79a9376.
//
// Solidity: event NodeUpdated2PublicGood(address indexed nodeAddr)
func (_StakingVSL *StakingVSLFilterer) WatchNodeUpdated2PublicGood(opts *bind.WatchOpts, sink chan<- *StakingVSLNodeUpdated2PublicGood, nodeAddr []common.Address) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "NodeUpdated2PublicGood", nodeAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLNodeUpdated2PublicGood)
				if err := _StakingVSL.contract.UnpackLog(event, "NodeUpdated2PublicGood", log); err != nil {
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

// ParseNodeUpdated2PublicGood is a log parse operation binding the contract event 0x86538b79bef9c52dbe4e888742cfbd70114655c47ef30ede4997791fe79a9376.
//
// Solidity: event NodeUpdated2PublicGood(address indexed nodeAddr)
func (_StakingVSL *StakingVSLFilterer) ParseNodeUpdated2PublicGood(log types.Log) (*StakingVSLNodeUpdated2PublicGood, error) {
	event := new(StakingVSLNodeUpdated2PublicGood)
	if err := _StakingVSL.contract.UnpackLog(event, "NodeUpdated2PublicGood", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingVSL contract.
type StakingVSLPausedIterator struct {
	Event *StakingVSLPaused // Event containing the contract specifics and raw log

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
func (it *StakingVSLPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLPaused)
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
		it.Event = new(StakingVSLPaused)
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
func (it *StakingVSLPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLPaused represents a Paused event raised by the StakingVSL contract.
type StakingVSLPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingVSL *StakingVSLFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingVSLPausedIterator, error) {

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingVSLPausedIterator{contract: _StakingVSL.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingVSL *StakingVSLFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingVSLPaused) (event.Subscription, error) {

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLPaused)
				if err := _StakingVSL.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingVSL *StakingVSLFilterer) ParsePaused(log types.Log) (*StakingVSLPaused, error) {
	event := new(StakingVSLPaused)
	if err := _StakingVSL.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLPublicGoodRewardDistributedIterator is returned from FilterPublicGoodRewardDistributed and is used to iterate over the raw logs and unpacked data for PublicGoodRewardDistributed events raised by the StakingVSL contract.
type StakingVSLPublicGoodRewardDistributedIterator struct {
	Event *StakingVSLPublicGoodRewardDistributed // Event containing the contract specifics and raw log

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
func (it *StakingVSLPublicGoodRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLPublicGoodRewardDistributed)
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
		it.Event = new(StakingVSLPublicGoodRewardDistributed)
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
func (it *StakingVSLPublicGoodRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLPublicGoodRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLPublicGoodRewardDistributed represents a PublicGoodRewardDistributed event raised by the StakingVSL contract.
type StakingVSLPublicGoodRewardDistributed struct {
	Epoch             *big.Int
	StartTimestamp    *big.Int
	EndTimestamp      *big.Int
	PublicPoolRewards *big.Int
	PublicPoolTax     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPublicGoodRewardDistributed is a free log retrieval operation binding the contract event 0xab7d25a2f6206ef56c88807f2474ddcd97e1a6323cb25149cde3a607fed6f2d7.
//
// Solidity: event PublicGoodRewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, uint256 publicPoolRewards, uint256 publicPoolTax)
func (_StakingVSL *StakingVSLFilterer) FilterPublicGoodRewardDistributed(opts *bind.FilterOpts, epoch []*big.Int) (*StakingVSLPublicGoodRewardDistributedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "PublicGoodRewardDistributed", epochRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLPublicGoodRewardDistributedIterator{contract: _StakingVSL.contract, event: "PublicGoodRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchPublicGoodRewardDistributed is a free log subscription operation binding the contract event 0xab7d25a2f6206ef56c88807f2474ddcd97e1a6323cb25149cde3a607fed6f2d7.
//
// Solidity: event PublicGoodRewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, uint256 publicPoolRewards, uint256 publicPoolTax)
func (_StakingVSL *StakingVSLFilterer) WatchPublicGoodRewardDistributed(opts *bind.WatchOpts, sink chan<- *StakingVSLPublicGoodRewardDistributed, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "PublicGoodRewardDistributed", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLPublicGoodRewardDistributed)
				if err := _StakingVSL.contract.UnpackLog(event, "PublicGoodRewardDistributed", log); err != nil {
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

// ParsePublicGoodRewardDistributed is a log parse operation binding the contract event 0xab7d25a2f6206ef56c88807f2474ddcd97e1a6323cb25149cde3a607fed6f2d7.
//
// Solidity: event PublicGoodRewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, uint256 publicPoolRewards, uint256 publicPoolTax)
func (_StakingVSL *StakingVSLFilterer) ParsePublicGoodRewardDistributed(log types.Log) (*StakingVSLPublicGoodRewardDistributed, error) {
	event := new(StakingVSLPublicGoodRewardDistributed)
	if err := _StakingVSL.contract.UnpackLog(event, "PublicGoodRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLPublicPoolTaxRateBasisPointsSetIterator is returned from FilterPublicPoolTaxRateBasisPointsSet and is used to iterate over the raw logs and unpacked data for PublicPoolTaxRateBasisPointsSet events raised by the StakingVSL contract.
type StakingVSLPublicPoolTaxRateBasisPointsSetIterator struct {
	Event *StakingVSLPublicPoolTaxRateBasisPointsSet // Event containing the contract specifics and raw log

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
func (it *StakingVSLPublicPoolTaxRateBasisPointsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLPublicPoolTaxRateBasisPointsSet)
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
		it.Event = new(StakingVSLPublicPoolTaxRateBasisPointsSet)
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
func (it *StakingVSLPublicPoolTaxRateBasisPointsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLPublicPoolTaxRateBasisPointsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLPublicPoolTaxRateBasisPointsSet represents a PublicPoolTaxRateBasisPointsSet event raised by the StakingVSL contract.
type StakingVSLPublicPoolTaxRateBasisPointsSet struct {
	TaxRateBasisPoints uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterPublicPoolTaxRateBasisPointsSet is a free log retrieval operation binding the contract event 0x948cf2302b029d76db2ac06e4ef2625e6c687335de349317468f47942a44e8b0.
//
// Solidity: event PublicPoolTaxRateBasisPointsSet(uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) FilterPublicPoolTaxRateBasisPointsSet(opts *bind.FilterOpts, taxRateBasisPoints []uint64) (*StakingVSLPublicPoolTaxRateBasisPointsSetIterator, error) {

	var taxRateBasisPointsRule []interface{}
	for _, taxRateBasisPointsItem := range taxRateBasisPoints {
		taxRateBasisPointsRule = append(taxRateBasisPointsRule, taxRateBasisPointsItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "PublicPoolTaxRateBasisPointsSet", taxRateBasisPointsRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLPublicPoolTaxRateBasisPointsSetIterator{contract: _StakingVSL.contract, event: "PublicPoolTaxRateBasisPointsSet", logs: logs, sub: sub}, nil
}

// WatchPublicPoolTaxRateBasisPointsSet is a free log subscription operation binding the contract event 0x948cf2302b029d76db2ac06e4ef2625e6c687335de349317468f47942a44e8b0.
//
// Solidity: event PublicPoolTaxRateBasisPointsSet(uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) WatchPublicPoolTaxRateBasisPointsSet(opts *bind.WatchOpts, sink chan<- *StakingVSLPublicPoolTaxRateBasisPointsSet, taxRateBasisPoints []uint64) (event.Subscription, error) {

	var taxRateBasisPointsRule []interface{}
	for _, taxRateBasisPointsItem := range taxRateBasisPoints {
		taxRateBasisPointsRule = append(taxRateBasisPointsRule, taxRateBasisPointsItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "PublicPoolTaxRateBasisPointsSet", taxRateBasisPointsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLPublicPoolTaxRateBasisPointsSet)
				if err := _StakingVSL.contract.UnpackLog(event, "PublicPoolTaxRateBasisPointsSet", log); err != nil {
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

// ParsePublicPoolTaxRateBasisPointsSet is a log parse operation binding the contract event 0x948cf2302b029d76db2ac06e4ef2625e6c687335de349317468f47942a44e8b0.
//
// Solidity: event PublicPoolTaxRateBasisPointsSet(uint64 indexed taxRateBasisPoints)
func (_StakingVSL *StakingVSLFilterer) ParsePublicPoolTaxRateBasisPointsSet(log types.Log) (*StakingVSLPublicPoolTaxRateBasisPointsSet, error) {
	event := new(StakingVSLPublicPoolTaxRateBasisPointsSet)
	if err := _StakingVSL.contract.UnpackLog(event, "PublicPoolTaxRateBasisPointsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLRewardDistributedIterator is returned from FilterRewardDistributed and is used to iterate over the raw logs and unpacked data for RewardDistributed events raised by the StakingVSL contract.
type StakingVSLRewardDistributedIterator struct {
	Event *StakingVSLRewardDistributed // Event containing the contract specifics and raw log

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
func (it *StakingVSLRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLRewardDistributed)
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
		it.Event = new(StakingVSLRewardDistributed)
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
func (it *StakingVSLRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLRewardDistributed represents a RewardDistributed event raised by the StakingVSL contract.
type StakingVSLRewardDistributed struct {
	Epoch            *big.Int
	StartTimestamp   *big.Int
	EndTimestamp     *big.Int
	NodeAddrs        []common.Address
	OperationRewards []*big.Int
	StakingRewards   []*big.Int
	TaxAmounts       []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardDistributed is a free log retrieval operation binding the contract event 0x255edf64a597e0852bdcb216fc8c3e79febaf9a041e1f6e546215fbadc96c338.
//
// Solidity: event RewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256[] taxAmounts)
func (_StakingVSL *StakingVSLFilterer) FilterRewardDistributed(opts *bind.FilterOpts, epoch []*big.Int) (*StakingVSLRewardDistributedIterator, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "RewardDistributed", epochRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLRewardDistributedIterator{contract: _StakingVSL.contract, event: "RewardDistributed", logs: logs, sub: sub}, nil
}

// WatchRewardDistributed is a free log subscription operation binding the contract event 0x255edf64a597e0852bdcb216fc8c3e79febaf9a041e1f6e546215fbadc96c338.
//
// Solidity: event RewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256[] taxAmounts)
func (_StakingVSL *StakingVSLFilterer) WatchRewardDistributed(opts *bind.WatchOpts, sink chan<- *StakingVSLRewardDistributed, epoch []*big.Int) (event.Subscription, error) {

	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "RewardDistributed", epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLRewardDistributed)
				if err := _StakingVSL.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
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

// ParseRewardDistributed is a log parse operation binding the contract event 0x255edf64a597e0852bdcb216fc8c3e79febaf9a041e1f6e546215fbadc96c338.
//
// Solidity: event RewardDistributed(uint256 indexed epoch, uint256 startTimestamp, uint256 endTimestamp, address[] nodeAddrs, uint256[] operationRewards, uint256[] stakingRewards, uint256[] taxAmounts)
func (_StakingVSL *StakingVSLFilterer) ParseRewardDistributed(log types.Log) (*StakingVSLRewardDistributed, error) {
	event := new(StakingVSLRewardDistributed)
	if err := _StakingVSL.contract.UnpackLog(event, "RewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingVSL contract.
type StakingVSLRoleAdminChangedIterator struct {
	Event *StakingVSLRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingVSLRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLRoleAdminChanged)
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
		it.Event = new(StakingVSLRoleAdminChanged)
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
func (it *StakingVSLRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLRoleAdminChanged represents a RoleAdminChanged event raised by the StakingVSL contract.
type StakingVSLRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingVSL *StakingVSLFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingVSLRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLRoleAdminChangedIterator{contract: _StakingVSL.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingVSL *StakingVSLFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingVSLRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLRoleAdminChanged)
				if err := _StakingVSL.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingVSL *StakingVSLFilterer) ParseRoleAdminChanged(log types.Log) (*StakingVSLRoleAdminChanged, error) {
	event := new(StakingVSLRoleAdminChanged)
	if err := _StakingVSL.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingVSL contract.
type StakingVSLRoleGrantedIterator struct {
	Event *StakingVSLRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingVSLRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLRoleGranted)
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
		it.Event = new(StakingVSLRoleGranted)
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
func (it *StakingVSLRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLRoleGranted represents a RoleGranted event raised by the StakingVSL contract.
type StakingVSLRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingVSL *StakingVSLFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingVSLRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLRoleGrantedIterator{contract: _StakingVSL.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingVSL *StakingVSLFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingVSLRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLRoleGranted)
				if err := _StakingVSL.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingVSL *StakingVSLFilterer) ParseRoleGranted(log types.Log) (*StakingVSLRoleGranted, error) {
	event := new(StakingVSLRoleGranted)
	if err := _StakingVSL.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingVSL contract.
type StakingVSLRoleRevokedIterator struct {
	Event *StakingVSLRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingVSLRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLRoleRevoked)
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
		it.Event = new(StakingVSLRoleRevoked)
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
func (it *StakingVSLRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLRoleRevoked represents a RoleRevoked event raised by the StakingVSL contract.
type StakingVSLRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingVSL *StakingVSLFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingVSLRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLRoleRevokedIterator{contract: _StakingVSL.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingVSL *StakingVSLFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingVSLRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLRoleRevoked)
				if err := _StakingVSL.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingVSL *StakingVSLFilterer) ParseRoleRevoked(log types.Log) (*StakingVSLRoleRevoked, error) {
	event := new(StakingVSLRoleRevoked)
	if err := _StakingVSL.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the StakingVSL contract.
type StakingVSLStakedIterator struct {
	Event *StakingVSLStaked // Event containing the contract specifics and raw log

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
func (it *StakingVSLStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLStaked)
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
		it.Event = new(StakingVSLStaked)
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
func (it *StakingVSLStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLStaked represents a Staked event raised by the StakingVSL contract.
type StakingVSLStaked struct {
	User         common.Address
	NodeAddr     common.Address
	Amount       *big.Int
	StartTokenId *big.Int
	EndTokenId   *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address indexed nodeAddr, uint256 indexed amount, uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, nodeAddr []common.Address, amount []*big.Int) (*StakingVSLStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "Staked", userRule, nodeAddrRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLStakedIterator{contract: _StakingVSL.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address indexed nodeAddr, uint256 indexed amount, uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingVSLStaked, user []common.Address, nodeAddr []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "Staked", userRule, nodeAddrRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLStaked)
				if err := _StakingVSL.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0xad3fa07f4195b47e64892eb944ecbfc253384053c119852bb2bcae484c2fcb69.
//
// Solidity: event Staked(address indexed user, address indexed nodeAddr, uint256 indexed amount, uint256 startTokenId, uint256 endTokenId)
func (_StakingVSL *StakingVSLFilterer) ParseStaked(log types.Log) (*StakingVSLStaked, error) {
	event := new(StakingVSLStaked)
	if err := _StakingVSL.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingVSL contract.
type StakingVSLUnpausedIterator struct {
	Event *StakingVSLUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingVSLUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLUnpaused)
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
		it.Event = new(StakingVSLUnpaused)
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
func (it *StakingVSLUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLUnpaused represents a Unpaused event raised by the StakingVSL contract.
type StakingVSLUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingVSL *StakingVSLFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingVSLUnpausedIterator, error) {

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingVSLUnpausedIterator{contract: _StakingVSL.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingVSL *StakingVSLFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingVSLUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLUnpaused)
				if err := _StakingVSL.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingVSL *StakingVSLFilterer) ParseUnpaused(log types.Log) (*StakingVSLUnpaused, error) {
	event := new(StakingVSLUnpaused)
	if err := _StakingVSL.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLUnstakeClaimedIterator is returned from FilterUnstakeClaimed and is used to iterate over the raw logs and unpacked data for UnstakeClaimed events raised by the StakingVSL contract.
type StakingVSLUnstakeClaimedIterator struct {
	Event *StakingVSLUnstakeClaimed // Event containing the contract specifics and raw log

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
func (it *StakingVSLUnstakeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLUnstakeClaimed)
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
		it.Event = new(StakingVSLUnstakeClaimed)
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
func (it *StakingVSLUnstakeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLUnstakeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLUnstakeClaimed represents a UnstakeClaimed event raised by the StakingVSL contract.
type StakingVSLUnstakeClaimed struct {
	RequestId     *big.Int
	NodeAddr      common.Address
	User          common.Address
	UnstakeAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeClaimed is a free log retrieval operation binding the contract event 0x2769ece66eadb650afd8c08c7a8772e39381dddd7230f9e039669e631044d47c.
//
// Solidity: event UnstakeClaimed(uint256 indexed requestId, address indexed nodeAddr, address indexed user, uint256 unstakeAmount)
func (_StakingVSL *StakingVSLFilterer) FilterUnstakeClaimed(opts *bind.FilterOpts, requestId []*big.Int, nodeAddr []common.Address, user []common.Address) (*StakingVSLUnstakeClaimedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "UnstakeClaimed", requestIdRule, nodeAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLUnstakeClaimedIterator{contract: _StakingVSL.contract, event: "UnstakeClaimed", logs: logs, sub: sub}, nil
}

// WatchUnstakeClaimed is a free log subscription operation binding the contract event 0x2769ece66eadb650afd8c08c7a8772e39381dddd7230f9e039669e631044d47c.
//
// Solidity: event UnstakeClaimed(uint256 indexed requestId, address indexed nodeAddr, address indexed user, uint256 unstakeAmount)
func (_StakingVSL *StakingVSLFilterer) WatchUnstakeClaimed(opts *bind.WatchOpts, sink chan<- *StakingVSLUnstakeClaimed, requestId []*big.Int, nodeAddr []common.Address, user []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "UnstakeClaimed", requestIdRule, nodeAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLUnstakeClaimed)
				if err := _StakingVSL.contract.UnpackLog(event, "UnstakeClaimed", log); err != nil {
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

// ParseUnstakeClaimed is a log parse operation binding the contract event 0x2769ece66eadb650afd8c08c7a8772e39381dddd7230f9e039669e631044d47c.
//
// Solidity: event UnstakeClaimed(uint256 indexed requestId, address indexed nodeAddr, address indexed user, uint256 unstakeAmount)
func (_StakingVSL *StakingVSLFilterer) ParseUnstakeClaimed(log types.Log) (*StakingVSLUnstakeClaimed, error) {
	event := new(StakingVSLUnstakeClaimed)
	if err := _StakingVSL.contract.UnpackLog(event, "UnstakeClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the StakingVSL contract.
type StakingVSLUnstakeRequestedIterator struct {
	Event *StakingVSLUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *StakingVSLUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLUnstakeRequested)
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
		it.Event = new(StakingVSLUnstakeRequested)
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
func (it *StakingVSLUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLUnstakeRequested represents a UnstakeRequested event raised by the StakingVSL contract.
type StakingVSLUnstakeRequested struct {
	User          common.Address
	NodeAddr      common.Address
	RequestId     *big.Int
	UnstakeAmount *big.Int
	ChipsIds      []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0x2808f92d5a0fada467cbe4e766f62f323e78271a7471058a87ef63a9e3e4c5c5.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed nodeAddr, uint256 indexed requestId, uint256 unstakeAmount, uint256[] chipsIds)
func (_StakingVSL *StakingVSLFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, user []common.Address, nodeAddr []common.Address, requestId []*big.Int) (*StakingVSLUnstakeRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "UnstakeRequested", userRule, nodeAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLUnstakeRequestedIterator{contract: _StakingVSL.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0x2808f92d5a0fada467cbe4e766f62f323e78271a7471058a87ef63a9e3e4c5c5.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed nodeAddr, uint256 indexed requestId, uint256 unstakeAmount, uint256[] chipsIds)
func (_StakingVSL *StakingVSLFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *StakingVSLUnstakeRequested, user []common.Address, nodeAddr []common.Address, requestId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "UnstakeRequested", userRule, nodeAddrRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLUnstakeRequested)
				if err := _StakingVSL.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0x2808f92d5a0fada467cbe4e766f62f323e78271a7471058a87ef63a9e3e4c5c5.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed nodeAddr, uint256 indexed requestId, uint256 unstakeAmount, uint256[] chipsIds)
func (_StakingVSL *StakingVSLFilterer) ParseUnstakeRequested(log types.Log) (*StakingVSLUnstakeRequested, error) {
	event := new(StakingVSLUnstakeRequested)
	if err := _StakingVSL.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLWithdrawRequestedIterator is returned from FilterWithdrawRequested and is used to iterate over the raw logs and unpacked data for WithdrawRequested events raised by the StakingVSL contract.
type StakingVSLWithdrawRequestedIterator struct {
	Event *StakingVSLWithdrawRequested // Event containing the contract specifics and raw log

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
func (it *StakingVSLWithdrawRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLWithdrawRequested)
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
		it.Event = new(StakingVSLWithdrawRequested)
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
func (it *StakingVSLWithdrawRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLWithdrawRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLWithdrawRequested represents a WithdrawRequested event raised by the StakingVSL contract.
type StakingVSLWithdrawRequested struct {
	NodeAddr  common.Address
	Amount    *big.Int
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawRequested is a free log retrieval operation binding the contract event 0xd72eb5d043f24a0168ae744d5c44f9596fd673a26bf74d9646bff4b844882d14.
//
// Solidity: event WithdrawRequested(address indexed nodeAddr, uint256 indexed amount, uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) FilterWithdrawRequested(opts *bind.FilterOpts, nodeAddr []common.Address, amount []*big.Int, requestId []*big.Int) (*StakingVSLWithdrawRequestedIterator, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "WithdrawRequested", nodeAddrRule, amountRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLWithdrawRequestedIterator{contract: _StakingVSL.contract, event: "WithdrawRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawRequested is a free log subscription operation binding the contract event 0xd72eb5d043f24a0168ae744d5c44f9596fd673a26bf74d9646bff4b844882d14.
//
// Solidity: event WithdrawRequested(address indexed nodeAddr, uint256 indexed amount, uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) WatchWithdrawRequested(opts *bind.WatchOpts, sink chan<- *StakingVSLWithdrawRequested, nodeAddr []common.Address, amount []*big.Int, requestId []*big.Int) (event.Subscription, error) {

	var nodeAddrRule []interface{}
	for _, nodeAddrItem := range nodeAddr {
		nodeAddrRule = append(nodeAddrRule, nodeAddrItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}
	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "WithdrawRequested", nodeAddrRule, amountRule, requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLWithdrawRequested)
				if err := _StakingVSL.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
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

// ParseWithdrawRequested is a log parse operation binding the contract event 0xd72eb5d043f24a0168ae744d5c44f9596fd673a26bf74d9646bff4b844882d14.
//
// Solidity: event WithdrawRequested(address indexed nodeAddr, uint256 indexed amount, uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) ParseWithdrawRequested(log types.Log) (*StakingVSLWithdrawRequested, error) {
	event := new(StakingVSLWithdrawRequested)
	if err := _StakingVSL.contract.UnpackLog(event, "WithdrawRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingVSLWithdrawalClaimedIterator is returned from FilterWithdrawalClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalClaimed events raised by the StakingVSL contract.
type StakingVSLWithdrawalClaimedIterator struct {
	Event *StakingVSLWithdrawalClaimed // Event containing the contract specifics and raw log

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
func (it *StakingVSLWithdrawalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingVSLWithdrawalClaimed)
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
		it.Event = new(StakingVSLWithdrawalClaimed)
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
func (it *StakingVSLWithdrawalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingVSLWithdrawalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingVSLWithdrawalClaimed represents a WithdrawalClaimed event raised by the StakingVSL contract.
type StakingVSLWithdrawalClaimed struct {
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalClaimed is a free log retrieval operation binding the contract event 0x8772d6f79a1845a0c0e90ef18d99f91242bbc0ba98c9ca780feaad42b81f02ba.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) FilterWithdrawalClaimed(opts *bind.FilterOpts, requestId []*big.Int) (*StakingVSLWithdrawalClaimedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.FilterLogs(opts, "WithdrawalClaimed", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingVSLWithdrawalClaimedIterator{contract: _StakingVSL.contract, event: "WithdrawalClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalClaimed is a free log subscription operation binding the contract event 0x8772d6f79a1845a0c0e90ef18d99f91242bbc0ba98c9ca780feaad42b81f02ba.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) WatchWithdrawalClaimed(opts *bind.WatchOpts, sink chan<- *StakingVSLWithdrawalClaimed, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _StakingVSL.contract.WatchLogs(opts, "WithdrawalClaimed", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingVSLWithdrawalClaimed)
				if err := _StakingVSL.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
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

// ParseWithdrawalClaimed is a log parse operation binding the contract event 0x8772d6f79a1845a0c0e90ef18d99f91242bbc0ba98c9ca780feaad42b81f02ba.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId)
func (_StakingVSL *StakingVSLFilterer) ParseWithdrawalClaimed(log types.Log) (*StakingVSLWithdrawalClaimed, error) {
	event := new(StakingVSLWithdrawalClaimed)
	if err := _StakingVSL.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
