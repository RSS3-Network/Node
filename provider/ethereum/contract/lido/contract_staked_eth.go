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

// StakedETHMetaData contains all meta data concerning the StakedETH contract.
var StakedETHMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_CONTROL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_ethAmount\",\"type\":\"uint256\"}],\"name\":\"getSharesByPooledEth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStakingPaused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"_stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"setStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalPooledEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newDepositedValidators\",\"type\":\"uint256\"}],\"name\":\"unsafeChangeDepositedValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTreasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isStopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBufferedEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_lidoLocator\",\"type\":\"address\"},{\"name\":\"_eip712StETH\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveELRewards\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentStakeLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeLimitFullInfo\",\"outputs\":[{\"name\":\"isStakingPaused\",\"type\":\"bool\"},{\"name\":\"isStakingLimitSet\",\"type\":\"bool\"},{\"name\":\"currentStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"name\":\"maxStakeLimitGrowthBlocks\",\"type\":\"uint256\"},{\"name\":\"prevStakeLimit\",\"type\":\"uint256\"},{\"name\":\"prevStakeBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferSharesFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFeeDistribution\",\"outputs\":[{\"name\":\"treasuryFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"insuranceFeeBasisPoints\",\"type\":\"uint16\"},{\"name\":\"operatorsFeeBasisPoints\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"receiveWithdrawals\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"getPooledEthByShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOracle\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"version\",\"type\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_sharesAmount\",\"type\":\"uint256\"}],\"name\":\"transferShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEIP712StETH\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_referral\",\"type\":\"address\"}],\"name\":\"submit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_maxDepositsCount\",\"type\":\"uint256\"},{\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBeaconStat\",\"outputs\":[{\"name\":\"depositedValidators\",\"type\":\"uint256\"},{\"name\":\"beaconValidators\",\"type\":\"uint256\"},{\"name\":\"beaconBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeStakingLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reportTimestamp\",\"type\":\"uint256\"},{\"name\":\"_timeElapsed\",\"type\":\"uint256\"},{\"name\":\"_clValidators\",\"type\":\"uint256\"},{\"name\":\"_clBalance\",\"type\":\"uint256\"},{\"name\":\"_withdrawalVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_elRewardsVaultBalance\",\"type\":\"uint256\"},{\"name\":\"_sharesRequestedToBurn\",\"type\":\"uint256\"},{\"name\":\"_withdrawalFinalizationBatches\",\"type\":\"uint256[]\"},{\"name\":\"_simulatedShareRate\",\"type\":\"uint256\"}],\"name\":\"handleOracleReport\",\"outputs\":[{\"name\":\"postRebaseAmounts\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFee\",\"outputs\":[{\"name\":\"totalFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalShares\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_deadline\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLidoLocator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canDeposit\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"STAKING_PAUSE_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositableEther\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"sharesOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalELRewardsCollected\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"maxStakeLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"stakeLimitIncreasePerBlock\",\"type\":\"uint256\"}],\"name\":\"StakingLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"StakingLimitRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLValidators\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLValidators\",\"type\":\"uint256\"}],\"name\":\"CLValidatorsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositedValidators\",\"type\":\"uint256\"}],\"name\":\"DepositedValidatorsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postCLBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"withdrawalsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"executionLayerRewardsWithdrawn\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postBufferedEther\",\"type\":\"uint256\"}],\"name\":\"ETHDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"reportTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeElapsed\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"preTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalShares\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postTotalEther\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesMintedAsFees\",\"type\":\"uint256\"}],\"name\":\"TokenRebased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lidoLocator\",\"type\":\"address\"}],\"name\":\"LidoLocatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"referral\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unbuffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eip712StETH\",\"type\":\"address\"}],\"name\":\"EIP712StETHInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sharesValue\",\"type\":\"uint256\"}],\"name\":\"TransferShares\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"postRebaseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sharesAmount\",\"type\":\"uint256\"}],\"name\":\"SharesBurnt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Stopped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"}]",
}

// StakedETHABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedETHMetaData.ABI instead.
var StakedETHABI = StakedETHMetaData.ABI

// StakedETH is an auto generated Go binding around an Ethereum contract.
type StakedETH struct {
	StakedETHCaller     // Read-only binding to the contract
	StakedETHTransactor // Write-only binding to the contract
	StakedETHFilterer   // Log filterer for contract events
}

// StakedETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedETHSession struct {
	Contract     *StakedETH        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakedETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedETHCallerSession struct {
	Contract *StakedETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StakedETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedETHTransactorSession struct {
	Contract     *StakedETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StakedETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedETHRaw struct {
	Contract *StakedETH // Generic contract binding to access the raw methods on
}

// StakedETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedETHCallerRaw struct {
	Contract *StakedETHCaller // Generic read-only contract binding to access the raw methods on
}

// StakedETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedETHTransactorRaw struct {
	Contract *StakedETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedETH creates a new instance of StakedETH, bound to a specific deployed contract.
func NewStakedETH(address common.Address, backend bind.ContractBackend) (*StakedETH, error) {
	contract, err := bindStakedETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedETH{StakedETHCaller: StakedETHCaller{contract: contract}, StakedETHTransactor: StakedETHTransactor{contract: contract}, StakedETHFilterer: StakedETHFilterer{contract: contract}}, nil
}

// NewStakedETHCaller creates a new read-only instance of StakedETH, bound to a specific deployed contract.
func NewStakedETHCaller(address common.Address, caller bind.ContractCaller) (*StakedETHCaller, error) {
	contract, err := bindStakedETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedETHCaller{contract: contract}, nil
}

// NewStakedETHTransactor creates a new write-only instance of StakedETH, bound to a specific deployed contract.
func NewStakedETHTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedETHTransactor, error) {
	contract, err := bindStakedETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedETHTransactor{contract: contract}, nil
}

// NewStakedETHFilterer creates a new log filterer instance of StakedETH, bound to a specific deployed contract.
func NewStakedETHFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedETHFilterer, error) {
	contract, err := bindStakedETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedETHFilterer{contract: contract}, nil
}

// bindStakedETH binds a generic wrapper to an already deployed contract.
func bindStakedETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedETH *StakedETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedETH.Contract.StakedETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedETH *StakedETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.Contract.StakedETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedETH *StakedETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedETH.Contract.StakedETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedETH *StakedETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedETH *StakedETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedETH *StakedETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedETH.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StakedETH *StakedETHCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StakedETH *StakedETHSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StakedETH.Contract.DOMAINSEPARATOR(&_StakedETH.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StakedETH.Contract.DOMAINSEPARATOR(&_StakedETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHSession) PAUSEROLE() ([32]byte, error) {
	return _StakedETH.Contract.PAUSEROLE(&_StakedETH.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StakedETH.Contract.PAUSEROLE(&_StakedETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHSession) RESUMEROLE() ([32]byte, error) {
	return _StakedETH.Contract.RESUMEROLE(&_StakedETH.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) RESUMEROLE() ([32]byte, error) {
	return _StakedETH.Contract.RESUMEROLE(&_StakedETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCaller) STAKINGCONTROLROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "STAKING_CONTROL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _StakedETH.Contract.STAKINGCONTROLROLE(&_StakedETH.CallOpts)
}

// STAKINGCONTROLROLE is a free data retrieval call binding the contract method 0x136dd43c.
//
// Solidity: function STAKING_CONTROL_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) STAKINGCONTROLROLE() ([32]byte, error) {
	return _StakedETH.Contract.STAKINGCONTROLROLE(&_StakedETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCaller) STAKINGPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "STAKING_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _StakedETH.Contract.STAKINGPAUSEROLE(&_StakedETH.CallOpts)
}

// STAKINGPAUSEROLE is a free data retrieval call binding the contract method 0xeb85262f.
//
// Solidity: function STAKING_PAUSE_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) STAKINGPAUSEROLE() ([32]byte, error) {
	return _StakedETH.Contract.STAKINGPAUSEROLE(&_StakedETH.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCaller) UNSAFECHANGEDEPOSITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _StakedETH.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_StakedETH.CallOpts)
}

// UNSAFECHANGEDEPOSITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xad1394e9.
//
// Solidity: function UNSAFE_CHANGE_DEPOSITED_VALIDATORS_ROLE() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) UNSAFECHANGEDEPOSITEDVALIDATORSROLE() ([32]byte, error) {
	return _StakedETH.Contract.UNSAFECHANGEDEPOSITEDVALIDATORSROLE(&_StakedETH.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StakedETH *StakedETHCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StakedETH *StakedETHSession) AllowRecoverability(token common.Address) (bool, error) {
	return _StakedETH.Contract.AllowRecoverability(&_StakedETH.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_StakedETH *StakedETHCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _StakedETH.Contract.AllowRecoverability(&_StakedETH.CallOpts, token)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedETH *StakedETHCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedETH *StakedETHSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StakedETH.Contract.Allowance(&_StakedETH.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StakedETH.Contract.Allowance(&_StakedETH.CallOpts, _owner, _spender)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StakedETH *StakedETHCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StakedETH *StakedETHSession) AppId() ([32]byte, error) {
	return _StakedETH.Contract.AppId(&_StakedETH.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) AppId() ([32]byte, error) {
	return _StakedETH.Contract.AppId(&_StakedETH.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StakedETH.Contract.BalanceOf(&_StakedETH.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _StakedETH.Contract.BalanceOf(&_StakedETH.CallOpts, _account)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StakedETH *StakedETHCaller) CanDeposit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "canDeposit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StakedETH *StakedETHSession) CanDeposit() (bool, error) {
	return _StakedETH.Contract.CanDeposit(&_StakedETH.CallOpts)
}

// CanDeposit is a free data retrieval call binding the contract method 0xe78a5875.
//
// Solidity: function canDeposit() view returns(bool)
func (_StakedETH *StakedETHCallerSession) CanDeposit() (bool, error) {
	return _StakedETH.Contract.CanDeposit(&_StakedETH.CallOpts)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StakedETH *StakedETHCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StakedETH *StakedETHSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _StakedETH.Contract.CanPerform(&_StakedETH.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_StakedETH *StakedETHCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _StakedETH.Contract.CanPerform(&_StakedETH.CallOpts, _sender, _role, _params)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedETH *StakedETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedETH *StakedETHSession) Decimals() (uint8, error) {
	return _StakedETH.Contract.Decimals(&_StakedETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_StakedETH *StakedETHCallerSession) Decimals() (uint8, error) {
	return _StakedETH.Contract.Decimals(&_StakedETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StakedETH *StakedETHCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StakedETH *StakedETHSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _StakedETH.Contract.Eip712Domain(&_StakedETH.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(string name, string version, uint256 chainId, address verifyingContract)
func (_StakedETH *StakedETHCallerSession) Eip712Domain() (struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}, error) {
	return _StakedETH.Contract.Eip712Domain(&_StakedETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StakedETH *StakedETHCaller) GetBeaconStat(opts *bind.CallOpts) (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getBeaconStat")

	outstruct := new(struct {
		DepositedValidators *big.Int
		BeaconValidators    *big.Int
		BeaconBalance       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DepositedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BeaconValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BeaconBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StakedETH *StakedETHSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _StakedETH.Contract.GetBeaconStat(&_StakedETH.CallOpts)
}

// GetBeaconStat is a free data retrieval call binding the contract method 0xae2e3538.
//
// Solidity: function getBeaconStat() view returns(uint256 depositedValidators, uint256 beaconValidators, uint256 beaconBalance)
func (_StakedETH *StakedETHCallerSession) GetBeaconStat() (struct {
	DepositedValidators *big.Int
	BeaconValidators    *big.Int
	BeaconBalance       *big.Int
}, error) {
	return _StakedETH.Contract.GetBeaconStat(&_StakedETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetBufferedEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getBufferedEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StakedETH *StakedETHSession) GetBufferedEther() (*big.Int, error) {
	return _StakedETH.Contract.GetBufferedEther(&_StakedETH.CallOpts)
}

// GetBufferedEther is a free data retrieval call binding the contract method 0x47b714e0.
//
// Solidity: function getBufferedEther() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetBufferedEther() (*big.Int, error) {
	return _StakedETH.Contract.GetBufferedEther(&_StakedETH.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETH *StakedETHSession) GetContractVersion() (*big.Int, error) {
	return _StakedETH.Contract.GetContractVersion(&_StakedETH.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetContractVersion() (*big.Int, error) {
	return _StakedETH.Contract.GetContractVersion(&_StakedETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetCurrentStakeLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getCurrentStakeLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StakedETH *StakedETHSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _StakedETH.Contract.GetCurrentStakeLimit(&_StakedETH.CallOpts)
}

// GetCurrentStakeLimit is a free data retrieval call binding the contract method 0x609c4c6c.
//
// Solidity: function getCurrentStakeLimit() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetCurrentStakeLimit() (*big.Int, error) {
	return _StakedETH.Contract.GetCurrentStakeLimit(&_StakedETH.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetDepositableEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getDepositableEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StakedETH *StakedETHSession) GetDepositableEther() (*big.Int, error) {
	return _StakedETH.Contract.GetDepositableEther(&_StakedETH.CallOpts)
}

// GetDepositableEther is a free data retrieval call binding the contract method 0xf2cfa87d.
//
// Solidity: function getDepositableEther() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetDepositableEther() (*big.Int, error) {
	return _StakedETH.Contract.GetDepositableEther(&_StakedETH.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StakedETH *StakedETHCaller) GetEIP712StETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getEIP712StETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StakedETH *StakedETHSession) GetEIP712StETH() (common.Address, error) {
	return _StakedETH.Contract.GetEIP712StETH(&_StakedETH.CallOpts)
}

// GetEIP712StETH is a free data retrieval call binding the contract method 0x9861f8e5.
//
// Solidity: function getEIP712StETH() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetEIP712StETH() (common.Address, error) {
	return _StakedETH.Contract.GetEIP712StETH(&_StakedETH.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StakedETH *StakedETHCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StakedETH *StakedETHSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _StakedETH.Contract.GetEVMScriptExecutor(&_StakedETH.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_StakedETH *StakedETHCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _StakedETH.Contract.GetEVMScriptExecutor(&_StakedETH.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StakedETH *StakedETHCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StakedETH *StakedETHSession) GetEVMScriptRegistry() (common.Address, error) {
	return _StakedETH.Contract.GetEVMScriptRegistry(&_StakedETH.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _StakedETH.Contract.GetEVMScriptRegistry(&_StakedETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StakedETH *StakedETHCaller) GetFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StakedETH *StakedETHSession) GetFee() (uint16, error) {
	return _StakedETH.Contract.GetFee(&_StakedETH.CallOpts)
}

// GetFee is a free data retrieval call binding the contract method 0xced72f87.
//
// Solidity: function getFee() view returns(uint16 totalFee)
func (_StakedETH *StakedETHCallerSession) GetFee() (uint16, error) {
	return _StakedETH.Contract.GetFee(&_StakedETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StakedETH *StakedETHCaller) GetFeeDistribution(opts *bind.CallOpts) (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getFeeDistribution")

	outstruct := new(struct {
		TreasuryFeeBasisPoints  uint16
		InsuranceFeeBasisPoints uint16
		OperatorsFeeBasisPoints uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TreasuryFeeBasisPoints = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.InsuranceFeeBasisPoints = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.OperatorsFeeBasisPoints = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StakedETH *StakedETHSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _StakedETH.Contract.GetFeeDistribution(&_StakedETH.CallOpts)
}

// GetFeeDistribution is a free data retrieval call binding the contract method 0x752f77f1.
//
// Solidity: function getFeeDistribution() view returns(uint16 treasuryFeeBasisPoints, uint16 insuranceFeeBasisPoints, uint16 operatorsFeeBasisPoints)
func (_StakedETH *StakedETHCallerSession) GetFeeDistribution() (struct {
	TreasuryFeeBasisPoints  uint16
	InsuranceFeeBasisPoints uint16
	OperatorsFeeBasisPoints uint16
}, error) {
	return _StakedETH.Contract.GetFeeDistribution(&_StakedETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StakedETH *StakedETHSession) GetInitializationBlock() (*big.Int, error) {
	return _StakedETH.Contract.GetInitializationBlock(&_StakedETH.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _StakedETH.Contract.GetInitializationBlock(&_StakedETH.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StakedETH *StakedETHCaller) GetLidoLocator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getLidoLocator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StakedETH *StakedETHSession) GetLidoLocator() (common.Address, error) {
	return _StakedETH.Contract.GetLidoLocator(&_StakedETH.CallOpts)
}

// GetLidoLocator is a free data retrieval call binding the contract method 0xe654ff17.
//
// Solidity: function getLidoLocator() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetLidoLocator() (common.Address, error) {
	return _StakedETH.Contract.GetLidoLocator(&_StakedETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StakedETH *StakedETHCaller) GetOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StakedETH *StakedETHSession) GetOracle() (common.Address, error) {
	return _StakedETH.Contract.GetOracle(&_StakedETH.CallOpts)
}

// GetOracle is a free data retrieval call binding the contract method 0x833b1fce.
//
// Solidity: function getOracle() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetOracle() (common.Address, error) {
	return _StakedETH.Contract.GetOracle(&_StakedETH.CallOpts)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StakedETH *StakedETHCaller) GetPooledEthByShares(opts *bind.CallOpts, _sharesAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getPooledEthByShares", _sharesAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StakedETH *StakedETHSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _StakedETH.Contract.GetPooledEthByShares(&_StakedETH.CallOpts, _sharesAmount)
}

// GetPooledEthByShares is a free data retrieval call binding the contract method 0x7a28fb88.
//
// Solidity: function getPooledEthByShares(uint256 _sharesAmount) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetPooledEthByShares(_sharesAmount *big.Int) (*big.Int, error) {
	return _StakedETH.Contract.GetPooledEthByShares(&_StakedETH.CallOpts, _sharesAmount)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StakedETH *StakedETHCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StakedETH *StakedETHSession) GetRecoveryVault() (common.Address, error) {
	return _StakedETH.Contract.GetRecoveryVault(&_StakedETH.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetRecoveryVault() (common.Address, error) {
	return _StakedETH.Contract.GetRecoveryVault(&_StakedETH.CallOpts)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StakedETH *StakedETHCaller) GetSharesByPooledEth(opts *bind.CallOpts, _ethAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getSharesByPooledEth", _ethAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StakedETH *StakedETHSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _StakedETH.Contract.GetSharesByPooledEth(&_StakedETH.CallOpts, _ethAmount)
}

// GetSharesByPooledEth is a free data retrieval call binding the contract method 0x19208451.
//
// Solidity: function getSharesByPooledEth(uint256 _ethAmount) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetSharesByPooledEth(_ethAmount *big.Int) (*big.Int, error) {
	return _StakedETH.Contract.GetSharesByPooledEth(&_StakedETH.CallOpts, _ethAmount)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StakedETH *StakedETHCaller) GetStakeLimitFullInfo(opts *bind.CallOpts) (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getStakeLimitFullInfo")

	outstruct := new(struct {
		IsStakingPaused           bool
		IsStakingLimitSet         bool
		CurrentStakeLimit         *big.Int
		MaxStakeLimit             *big.Int
		MaxStakeLimitGrowthBlocks *big.Int
		PrevStakeLimit            *big.Int
		PrevStakeBlockNumber      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsStakingPaused = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.IsStakingLimitSet = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.CurrentStakeLimit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxStakeLimitGrowthBlocks = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeLimit = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.PrevStakeBlockNumber = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StakedETH *StakedETHSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _StakedETH.Contract.GetStakeLimitFullInfo(&_StakedETH.CallOpts)
}

// GetStakeLimitFullInfo is a free data retrieval call binding the contract method 0x665b4b0b.
//
// Solidity: function getStakeLimitFullInfo() view returns(bool isStakingPaused, bool isStakingLimitSet, uint256 currentStakeLimit, uint256 maxStakeLimit, uint256 maxStakeLimitGrowthBlocks, uint256 prevStakeLimit, uint256 prevStakeBlockNumber)
func (_StakedETH *StakedETHCallerSession) GetStakeLimitFullInfo() (struct {
	IsStakingPaused           bool
	IsStakingLimitSet         bool
	CurrentStakeLimit         *big.Int
	MaxStakeLimit             *big.Int
	MaxStakeLimitGrowthBlocks *big.Int
	PrevStakeLimit            *big.Int
	PrevStakeBlockNumber      *big.Int
}, error) {
	return _StakedETH.Contract.GetStakeLimitFullInfo(&_StakedETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetTotalELRewardsCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getTotalELRewardsCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StakedETH *StakedETHSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalELRewardsCollected(&_StakedETH.CallOpts)
}

// GetTotalELRewardsCollected is a free data retrieval call binding the contract method 0xfa64ebac.
//
// Solidity: function getTotalELRewardsCollected() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetTotalELRewardsCollected() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalELRewardsCollected(&_StakedETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetTotalPooledEther(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getTotalPooledEther")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StakedETH *StakedETHSession) GetTotalPooledEther() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalPooledEther(&_StakedETH.CallOpts)
}

// GetTotalPooledEther is a free data retrieval call binding the contract method 0x37cfdaca.
//
// Solidity: function getTotalPooledEther() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetTotalPooledEther() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalPooledEther(&_StakedETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StakedETH *StakedETHCaller) GetTotalShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getTotalShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StakedETH *StakedETHSession) GetTotalShares() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalShares(&_StakedETH.CallOpts)
}

// GetTotalShares is a free data retrieval call binding the contract method 0xd5002f2e.
//
// Solidity: function getTotalShares() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) GetTotalShares() (*big.Int, error) {
	return _StakedETH.Contract.GetTotalShares(&_StakedETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakedETH *StakedETHCaller) GetTreasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getTreasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakedETH *StakedETHSession) GetTreasury() (common.Address, error) {
	return _StakedETH.Contract.GetTreasury(&_StakedETH.CallOpts)
}

// GetTreasury is a free data retrieval call binding the contract method 0x3b19e84a.
//
// Solidity: function getTreasury() view returns(address)
func (_StakedETH *StakedETHCallerSession) GetTreasury() (common.Address, error) {
	return _StakedETH.Contract.GetTreasury(&_StakedETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StakedETH *StakedETHCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StakedETH *StakedETHSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _StakedETH.Contract.GetWithdrawalCredentials(&_StakedETH.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_StakedETH *StakedETHCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _StakedETH.Contract.GetWithdrawalCredentials(&_StakedETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StakedETH *StakedETHCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StakedETH *StakedETHSession) HasInitialized() (bool, error) {
	return _StakedETH.Contract.HasInitialized(&_StakedETH.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_StakedETH *StakedETHCallerSession) HasInitialized() (bool, error) {
	return _StakedETH.Contract.HasInitialized(&_StakedETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StakedETH *StakedETHCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StakedETH *StakedETHSession) IsPetrified() (bool, error) {
	return _StakedETH.Contract.IsPetrified(&_StakedETH.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_StakedETH *StakedETHCallerSession) IsPetrified() (bool, error) {
	return _StakedETH.Contract.IsPetrified(&_StakedETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StakedETH *StakedETHCaller) IsStakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "isStakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StakedETH *StakedETHSession) IsStakingPaused() (bool, error) {
	return _StakedETH.Contract.IsStakingPaused(&_StakedETH.CallOpts)
}

// IsStakingPaused is a free data retrieval call binding the contract method 0x1ea7ca89.
//
// Solidity: function isStakingPaused() view returns(bool)
func (_StakedETH *StakedETHCallerSession) IsStakingPaused() (bool, error) {
	return _StakedETH.Contract.IsStakingPaused(&_StakedETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StakedETH *StakedETHCaller) IsStopped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "isStopped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StakedETH *StakedETHSession) IsStopped() (bool, error) {
	return _StakedETH.Contract.IsStopped(&_StakedETH.CallOpts)
}

// IsStopped is a free data retrieval call binding the contract method 0x3f683b6a.
//
// Solidity: function isStopped() view returns(bool)
func (_StakedETH *StakedETHCallerSession) IsStopped() (bool, error) {
	return _StakedETH.Contract.IsStopped(&_StakedETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StakedETH *StakedETHCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StakedETH *StakedETHSession) Kernel() (common.Address, error) {
	return _StakedETH.Contract.Kernel(&_StakedETH.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_StakedETH *StakedETHCallerSession) Kernel() (common.Address, error) {
	return _StakedETH.Contract.Kernel(&_StakedETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedETH *StakedETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedETH *StakedETHSession) Name() (string, error) {
	return _StakedETH.Contract.Name(&_StakedETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_StakedETH *StakedETHCallerSession) Name() (string, error) {
	return _StakedETH.Contract.Name(&_StakedETH.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedETH *StakedETHCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedETH *StakedETHSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StakedETH.Contract.Nonces(&_StakedETH.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _StakedETH.Contract.Nonces(&_StakedETH.CallOpts, owner)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHCaller) SharesOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "sharesOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _StakedETH.Contract.SharesOf(&_StakedETH.CallOpts, _account)
}

// SharesOf is a free data retrieval call binding the contract method 0xf5eb42dc.
//
// Solidity: function sharesOf(address _account) view returns(uint256)
func (_StakedETH *StakedETHCallerSession) SharesOf(_account common.Address) (*big.Int, error) {
	return _StakedETH.Contract.SharesOf(&_StakedETH.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedETH *StakedETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedETH *StakedETHSession) Symbol() (string, error) {
	return _StakedETH.Contract.Symbol(&_StakedETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_StakedETH *StakedETHCallerSession) Symbol() (string, error) {
	return _StakedETH.Contract.Symbol(&_StakedETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedETH *StakedETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedETH *StakedETHSession) TotalSupply() (*big.Int, error) {
	return _StakedETH.Contract.TotalSupply(&_StakedETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakedETH *StakedETHCallerSession) TotalSupply() (*big.Int, error) {
	return _StakedETH.Contract.TotalSupply(&_StakedETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.Approve(&_StakedETH.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.Approve(&_StakedETH.TransactOpts, _spender, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StakedETH *StakedETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "decreaseAllowance", _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StakedETH *StakedETHSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.DecreaseAllowance(&_StakedETH.TransactOpts, _spender, _subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address _spender, uint256 _subtractedValue) returns(bool)
func (_StakedETH *StakedETHTransactorSession) DecreaseAllowance(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.DecreaseAllowance(&_StakedETH.TransactOpts, _spender, _subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StakedETH *StakedETHTransactor) Deposit(opts *bind.TransactOpts, _maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "deposit", _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StakedETH *StakedETHSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Deposit(&_StakedETH.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _maxDepositsCount, uint256 _stakingModuleId, bytes _depositCalldata) returns()
func (_StakedETH *StakedETHTransactorSession) Deposit(_maxDepositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Deposit(&_StakedETH.TransactOpts, _maxDepositsCount, _stakingModuleId, _depositCalldata)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StakedETH *StakedETHTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "finalizeUpgrade_v2", _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StakedETH *StakedETHSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.FinalizeUpgradeV2(&_StakedETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x2f85e57c.
//
// Solidity: function finalizeUpgrade_v2(address _lidoLocator, address _eip712StETH) returns()
func (_StakedETH *StakedETHTransactorSession) FinalizeUpgradeV2(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.FinalizeUpgradeV2(&_StakedETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StakedETH *StakedETHTransactor) HandleOracleReport(opts *bind.TransactOpts, _reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "handleOracleReport", _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StakedETH *StakedETHSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.HandleOracleReport(&_StakedETH.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// HandleOracleReport is a paid mutator transaction binding the contract method 0xbac3f3c5.
//
// Solidity: function handleOracleReport(uint256 _reportTimestamp, uint256 _timeElapsed, uint256 _clValidators, uint256 _clBalance, uint256 _withdrawalVaultBalance, uint256 _elRewardsVaultBalance, uint256 _sharesRequestedToBurn, uint256[] _withdrawalFinalizationBatches, uint256 _simulatedShareRate) returns(uint256[4] postRebaseAmounts)
func (_StakedETH *StakedETHTransactorSession) HandleOracleReport(_reportTimestamp *big.Int, _timeElapsed *big.Int, _clValidators *big.Int, _clBalance *big.Int, _withdrawalVaultBalance *big.Int, _elRewardsVaultBalance *big.Int, _sharesRequestedToBurn *big.Int, _withdrawalFinalizationBatches []*big.Int, _simulatedShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.HandleOracleReport(&_StakedETH.TransactOpts, _reportTimestamp, _timeElapsed, _clValidators, _clBalance, _withdrawalVaultBalance, _elRewardsVaultBalance, _sharesRequestedToBurn, _withdrawalFinalizationBatches, _simulatedShareRate)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StakedETH *StakedETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "increaseAllowance", _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StakedETH *StakedETHSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.IncreaseAllowance(&_StakedETH.TransactOpts, _spender, _addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address _spender, uint256 _addedValue) returns(bool)
func (_StakedETH *StakedETHTransactorSession) IncreaseAllowance(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.IncreaseAllowance(&_StakedETH.TransactOpts, _spender, _addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StakedETH *StakedETHTransactor) Initialize(opts *bind.TransactOpts, _lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "initialize", _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StakedETH *StakedETHSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.Initialize(&_StakedETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _lidoLocator, address _eip712StETH) payable returns()
func (_StakedETH *StakedETHTransactorSession) Initialize(_lidoLocator common.Address, _eip712StETH common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.Initialize(&_StakedETH.TransactOpts, _lidoLocator, _eip712StETH)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StakedETH *StakedETHTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StakedETH *StakedETHSession) PauseStaking() (*types.Transaction, error) {
	return _StakedETH.Contract.PauseStaking(&_StakedETH.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_StakedETH *StakedETHTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _StakedETH.Contract.PauseStaking(&_StakedETH.TransactOpts)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StakedETH *StakedETHTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "permit", _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StakedETH *StakedETHSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Permit(&_StakedETH.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _value, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StakedETH *StakedETHTransactorSession) Permit(_owner common.Address, _spender common.Address, _value *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Permit(&_StakedETH.TransactOpts, _owner, _spender, _value, _deadline, _v, _r, _s)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StakedETH *StakedETHTransactor) ReceiveELRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "receiveELRewards")
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StakedETH *StakedETHSession) ReceiveELRewards() (*types.Transaction, error) {
	return _StakedETH.Contract.ReceiveELRewards(&_StakedETH.TransactOpts)
}

// ReceiveELRewards is a paid mutator transaction binding the contract method 0x4ad509b2.
//
// Solidity: function receiveELRewards() payable returns()
func (_StakedETH *StakedETHTransactorSession) ReceiveELRewards() (*types.Transaction, error) {
	return _StakedETH.Contract.ReceiveELRewards(&_StakedETH.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StakedETH *StakedETHTransactor) ReceiveWithdrawals(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "receiveWithdrawals")
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StakedETH *StakedETHSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _StakedETH.Contract.ReceiveWithdrawals(&_StakedETH.TransactOpts)
}

// ReceiveWithdrawals is a paid mutator transaction binding the contract method 0x78ffcfe2.
//
// Solidity: function receiveWithdrawals() payable returns()
func (_StakedETH *StakedETHTransactorSession) ReceiveWithdrawals() (*types.Transaction, error) {
	return _StakedETH.Contract.ReceiveWithdrawals(&_StakedETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StakedETH *StakedETHTransactor) RemoveStakingLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "removeStakingLimit")
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StakedETH *StakedETHSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _StakedETH.Contract.RemoveStakingLimit(&_StakedETH.TransactOpts)
}

// RemoveStakingLimit is a paid mutator transaction binding the contract method 0xb3320d9a.
//
// Solidity: function removeStakingLimit() returns()
func (_StakedETH *StakedETHTransactorSession) RemoveStakingLimit() (*types.Transaction, error) {
	return _StakedETH.Contract.RemoveStakingLimit(&_StakedETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETH *StakedETHTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETH *StakedETHSession) Resume() (*types.Transaction, error) {
	return _StakedETH.Contract.Resume(&_StakedETH.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETH *StakedETHTransactorSession) Resume() (*types.Transaction, error) {
	return _StakedETH.Contract.Resume(&_StakedETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StakedETH *StakedETHTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StakedETH *StakedETHSession) ResumeStaking() (*types.Transaction, error) {
	return _StakedETH.Contract.ResumeStaking(&_StakedETH.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_StakedETH *StakedETHTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _StakedETH.Contract.ResumeStaking(&_StakedETH.TransactOpts)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StakedETH *StakedETHTransactor) SetStakingLimit(opts *bind.TransactOpts, _maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "setStakingLimit", _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StakedETH *StakedETHSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.SetStakingLimit(&_StakedETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// SetStakingLimit is a paid mutator transaction binding the contract method 0x2cb5f784.
//
// Solidity: function setStakingLimit(uint256 _maxStakeLimit, uint256 _stakeLimitIncreasePerBlock) returns()
func (_StakedETH *StakedETHTransactorSession) SetStakingLimit(_maxStakeLimit *big.Int, _stakeLimitIncreasePerBlock *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.SetStakingLimit(&_StakedETH.TransactOpts, _maxStakeLimit, _stakeLimitIncreasePerBlock)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StakedETH *StakedETHTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StakedETH *StakedETHSession) Stop() (*types.Transaction, error) {
	return _StakedETH.Contract.Stop(&_StakedETH.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_StakedETH *StakedETHTransactorSession) Stop() (*types.Transaction, error) {
	return _StakedETH.Contract.Stop(&_StakedETH.TransactOpts)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StakedETH *StakedETHTransactor) Submit(opts *bind.TransactOpts, _referral common.Address) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "submit", _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StakedETH *StakedETHSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.Submit(&_StakedETH.TransactOpts, _referral)
}

// Submit is a paid mutator transaction binding the contract method 0xa1903eab.
//
// Solidity: function submit(address _referral) payable returns(uint256)
func (_StakedETH *StakedETHTransactorSession) Submit(_referral common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.Submit(&_StakedETH.TransactOpts, _referral)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.Transfer(&_StakedETH.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.Transfer(&_StakedETH.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferFrom(&_StakedETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_StakedETH *StakedETHTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferFrom(&_StakedETH.TransactOpts, _sender, _recipient, _amount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHTransactor) TransferShares(opts *bind.TransactOpts, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "transferShares", _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferShares(&_StakedETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferShares is a paid mutator transaction binding the contract method 0x8fcb4e5b.
//
// Solidity: function transferShares(address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHTransactorSession) TransferShares(_recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferShares(&_StakedETH.TransactOpts, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHTransactor) TransferSharesFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "transferSharesFrom", _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferSharesFrom(&_StakedETH.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferSharesFrom is a paid mutator transaction binding the contract method 0x6d780459.
//
// Solidity: function transferSharesFrom(address _sender, address _recipient, uint256 _sharesAmount) returns(uint256)
func (_StakedETH *StakedETHTransactorSession) TransferSharesFrom(_sender common.Address, _recipient common.Address, _sharesAmount *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferSharesFrom(&_StakedETH.TransactOpts, _sender, _recipient, _sharesAmount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StakedETH *StakedETHTransactor) TransferToVault(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "transferToVault", arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StakedETH *StakedETHSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferToVault(&_StakedETH.TransactOpts, arg0)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address ) returns()
func (_StakedETH *StakedETHTransactorSession) TransferToVault(arg0 common.Address) (*types.Transaction, error) {
	return _StakedETH.Contract.TransferToVault(&_StakedETH.TransactOpts, arg0)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StakedETH *StakedETHTransactor) UnsafeChangeDepositedValidators(opts *bind.TransactOpts, _newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StakedETH.contract.Transact(opts, "unsafeChangeDepositedValidators", _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StakedETH *StakedETHSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.UnsafeChangeDepositedValidators(&_StakedETH.TransactOpts, _newDepositedValidators)
}

// UnsafeChangeDepositedValidators is a paid mutator transaction binding the contract method 0x38998624.
//
// Solidity: function unsafeChangeDepositedValidators(uint256 _newDepositedValidators) returns()
func (_StakedETH *StakedETHTransactorSession) UnsafeChangeDepositedValidators(_newDepositedValidators *big.Int) (*types.Transaction, error) {
	return _StakedETH.Contract.UnsafeChangeDepositedValidators(&_StakedETH.TransactOpts, _newDepositedValidators)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedETH *StakedETHTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StakedETH.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedETH *StakedETHSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Fallback(&_StakedETH.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StakedETH *StakedETHTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StakedETH.Contract.Fallback(&_StakedETH.TransactOpts, calldata)
}

// StakedETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedETH contract.
type StakedETHApprovalIterator struct {
	Event *StakedETHApproval // Event containing the contract specifics and raw log

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
func (it *StakedETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHApproval)
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
		it.Event = new(StakedETHApproval)
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
func (it *StakedETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHApproval represents a Approval event raised by the StakedETH contract.
type StakedETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedETH *StakedETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StakedETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHApprovalIterator{contract: _StakedETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StakedETH *StakedETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHApproval)
				if err := _StakedETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StakedETH *StakedETHFilterer) ParseApproval(log types.Log) (*StakedETHApproval, error) {
	event := new(StakedETHApproval)
	if err := _StakedETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHCLValidatorsUpdatedIterator is returned from FilterCLValidatorsUpdated and is used to iterate over the raw logs and unpacked data for CLValidatorsUpdated events raised by the StakedETH contract.
type StakedETHCLValidatorsUpdatedIterator struct {
	Event *StakedETHCLValidatorsUpdated // Event containing the contract specifics and raw log

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
func (it *StakedETHCLValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHCLValidatorsUpdated)
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
		it.Event = new(StakedETHCLValidatorsUpdated)
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
func (it *StakedETHCLValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHCLValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHCLValidatorsUpdated represents a CLValidatorsUpdated event raised by the StakedETH contract.
type StakedETHCLValidatorsUpdated struct {
	ReportTimestamp  *big.Int
	PreCLValidators  *big.Int
	PostCLValidators *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCLValidatorsUpdated is a free log retrieval operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StakedETH *StakedETHFilterer) FilterCLValidatorsUpdated(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StakedETHCLValidatorsUpdatedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHCLValidatorsUpdatedIterator{contract: _StakedETH.contract, event: "CLValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchCLValidatorsUpdated is a free log subscription operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StakedETH *StakedETHFilterer) WatchCLValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *StakedETHCLValidatorsUpdated, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "CLValidatorsUpdated", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHCLValidatorsUpdated)
				if err := _StakedETH.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
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

// ParseCLValidatorsUpdated is a log parse operation binding the contract event 0x1252331d4f3ee8a9f0a3484c4c2fb059c70a047b5dc5482a3ee6415f742d9f2e.
//
// Solidity: event CLValidatorsUpdated(uint256 indexed reportTimestamp, uint256 preCLValidators, uint256 postCLValidators)
func (_StakedETH *StakedETHFilterer) ParseCLValidatorsUpdated(log types.Log) (*StakedETHCLValidatorsUpdated, error) {
	event := new(StakedETHCLValidatorsUpdated)
	if err := _StakedETH.contract.UnpackLog(event, "CLValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the StakedETH contract.
type StakedETHContractVersionSetIterator struct {
	Event *StakedETHContractVersionSet // Event containing the contract specifics and raw log

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
func (it *StakedETHContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHContractVersionSet)
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
		it.Event = new(StakedETHContractVersionSet)
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
func (it *StakedETHContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHContractVersionSet represents a ContractVersionSet event raised by the StakedETH contract.
type StakedETHContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StakedETH *StakedETHFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*StakedETHContractVersionSetIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &StakedETHContractVersionSetIterator{contract: _StakedETH.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StakedETH *StakedETHFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *StakedETHContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHContractVersionSet)
				if err := _StakedETH.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StakedETH *StakedETHFilterer) ParseContractVersionSet(log types.Log) (*StakedETHContractVersionSet, error) {
	event := new(StakedETHContractVersionSet)
	if err := _StakedETH.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHDepositedValidatorsChangedIterator is returned from FilterDepositedValidatorsChanged and is used to iterate over the raw logs and unpacked data for DepositedValidatorsChanged events raised by the StakedETH contract.
type StakedETHDepositedValidatorsChangedIterator struct {
	Event *StakedETHDepositedValidatorsChanged // Event containing the contract specifics and raw log

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
func (it *StakedETHDepositedValidatorsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHDepositedValidatorsChanged)
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
		it.Event = new(StakedETHDepositedValidatorsChanged)
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
func (it *StakedETHDepositedValidatorsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHDepositedValidatorsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHDepositedValidatorsChanged represents a DepositedValidatorsChanged event raised by the StakedETH contract.
type StakedETHDepositedValidatorsChanged struct {
	DepositedValidators *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDepositedValidatorsChanged is a free log retrieval operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StakedETH *StakedETHFilterer) FilterDepositedValidatorsChanged(opts *bind.FilterOpts) (*StakedETHDepositedValidatorsChangedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return &StakedETHDepositedValidatorsChangedIterator{contract: _StakedETH.contract, event: "DepositedValidatorsChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedValidatorsChanged is a free log subscription operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StakedETH *StakedETHFilterer) WatchDepositedValidatorsChanged(opts *bind.WatchOpts, sink chan<- *StakedETHDepositedValidatorsChanged) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "DepositedValidatorsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHDepositedValidatorsChanged)
				if err := _StakedETH.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
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

// ParseDepositedValidatorsChanged is a log parse operation binding the contract event 0xe0aacfc334457703148118055ec794ac17654c6f918d29638ba3b18003cee5ff.
//
// Solidity: event DepositedValidatorsChanged(uint256 depositedValidators)
func (_StakedETH *StakedETHFilterer) ParseDepositedValidatorsChanged(log types.Log) (*StakedETHDepositedValidatorsChanged, error) {
	event := new(StakedETHDepositedValidatorsChanged)
	if err := _StakedETH.contract.UnpackLog(event, "DepositedValidatorsChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHEIP712StETHInitializedIterator is returned from FilterEIP712StETHInitialized and is used to iterate over the raw logs and unpacked data for EIP712StETHInitialized events raised by the StakedETH contract.
type StakedETHEIP712StETHInitializedIterator struct {
	Event *StakedETHEIP712StETHInitialized // Event containing the contract specifics and raw log

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
func (it *StakedETHEIP712StETHInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHEIP712StETHInitialized)
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
		it.Event = new(StakedETHEIP712StETHInitialized)
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
func (it *StakedETHEIP712StETHInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHEIP712StETHInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHEIP712StETHInitialized represents a EIP712StETHInitialized event raised by the StakedETH contract.
type StakedETHEIP712StETHInitialized struct {
	Eip712StETH common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEIP712StETHInitialized is a free log retrieval operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StakedETH *StakedETHFilterer) FilterEIP712StETHInitialized(opts *bind.FilterOpts) (*StakedETHEIP712StETHInitializedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return &StakedETHEIP712StETHInitializedIterator{contract: _StakedETH.contract, event: "EIP712StETHInitialized", logs: logs, sub: sub}, nil
}

// WatchEIP712StETHInitialized is a free log subscription operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StakedETH *StakedETHFilterer) WatchEIP712StETHInitialized(opts *bind.WatchOpts, sink chan<- *StakedETHEIP712StETHInitialized) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "EIP712StETHInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHEIP712StETHInitialized)
				if err := _StakedETH.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
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

// ParseEIP712StETHInitialized is a log parse operation binding the contract event 0xb80a5409082a3729c9fc139f8b41192c40e85252752df2c07caebd613086ca83.
//
// Solidity: event EIP712StETHInitialized(address eip712StETH)
func (_StakedETH *StakedETHFilterer) ParseEIP712StETHInitialized(log types.Log) (*StakedETHEIP712StETHInitialized, error) {
	event := new(StakedETHEIP712StETHInitialized)
	if err := _StakedETH.contract.UnpackLog(event, "EIP712StETHInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHELRewardsReceivedIterator is returned from FilterELRewardsReceived and is used to iterate over the raw logs and unpacked data for ELRewardsReceived events raised by the StakedETH contract.
type StakedETHELRewardsReceivedIterator struct {
	Event *StakedETHELRewardsReceived // Event containing the contract specifics and raw log

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
func (it *StakedETHELRewardsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHELRewardsReceived)
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
		it.Event = new(StakedETHELRewardsReceived)
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
func (it *StakedETHELRewardsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHELRewardsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHELRewardsReceived represents a ELRewardsReceived event raised by the StakedETH contract.
type StakedETHELRewardsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterELRewardsReceived is a free log retrieval operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) FilterELRewardsReceived(opts *bind.FilterOpts) (*StakedETHELRewardsReceivedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return &StakedETHELRewardsReceivedIterator{contract: _StakedETH.contract, event: "ELRewardsReceived", logs: logs, sub: sub}, nil
}

// WatchELRewardsReceived is a free log subscription operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) WatchELRewardsReceived(opts *bind.WatchOpts, sink chan<- *StakedETHELRewardsReceived) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "ELRewardsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHELRewardsReceived)
				if err := _StakedETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
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

// ParseELRewardsReceived is a log parse operation binding the contract event 0xd27f9b0c98bdee27044afa149eadcd2047d6399cb6613a45c5b87e6aca76e6b5.
//
// Solidity: event ELRewardsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) ParseELRewardsReceived(log types.Log) (*StakedETHELRewardsReceived, error) {
	event := new(StakedETHELRewardsReceived)
	if err := _StakedETH.contract.UnpackLog(event, "ELRewardsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHETHDistributedIterator is returned from FilterETHDistributed and is used to iterate over the raw logs and unpacked data for ETHDistributed events raised by the StakedETH contract.
type StakedETHETHDistributedIterator struct {
	Event *StakedETHETHDistributed // Event containing the contract specifics and raw log

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
func (it *StakedETHETHDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHETHDistributed)
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
		it.Event = new(StakedETHETHDistributed)
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
func (it *StakedETHETHDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHETHDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHETHDistributed represents a ETHDistributed event raised by the StakedETH contract.
type StakedETHETHDistributed struct {
	ReportTimestamp                *big.Int
	PreCLBalance                   *big.Int
	PostCLBalance                  *big.Int
	WithdrawalsWithdrawn           *big.Int
	ExecutionLayerRewardsWithdrawn *big.Int
	PostBufferedEther              *big.Int
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterETHDistributed is a free log retrieval operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StakedETH *StakedETHFilterer) FilterETHDistributed(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StakedETHETHDistributedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHETHDistributedIterator{contract: _StakedETH.contract, event: "ETHDistributed", logs: logs, sub: sub}, nil
}

// WatchETHDistributed is a free log subscription operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StakedETH *StakedETHFilterer) WatchETHDistributed(opts *bind.WatchOpts, sink chan<- *StakedETHETHDistributed, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "ETHDistributed", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHETHDistributed)
				if err := _StakedETH.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
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

// ParseETHDistributed is a log parse operation binding the contract event 0x92dd3cb149a1eebd51fd8c2a3653fd96f30c4ac01d4f850fc16d46abd6c3e92f.
//
// Solidity: event ETHDistributed(uint256 indexed reportTimestamp, uint256 preCLBalance, uint256 postCLBalance, uint256 withdrawalsWithdrawn, uint256 executionLayerRewardsWithdrawn, uint256 postBufferedEther)
func (_StakedETH *StakedETHFilterer) ParseETHDistributed(log types.Log) (*StakedETHETHDistributed, error) {
	event := new(StakedETHETHDistributed)
	if err := _StakedETH.contract.UnpackLog(event, "ETHDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHLidoLocatorSetIterator is returned from FilterLidoLocatorSet and is used to iterate over the raw logs and unpacked data for LidoLocatorSet events raised by the StakedETH contract.
type StakedETHLidoLocatorSetIterator struct {
	Event *StakedETHLidoLocatorSet // Event containing the contract specifics and raw log

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
func (it *StakedETHLidoLocatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHLidoLocatorSet)
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
		it.Event = new(StakedETHLidoLocatorSet)
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
func (it *StakedETHLidoLocatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHLidoLocatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHLidoLocatorSet represents a LidoLocatorSet event raised by the StakedETH contract.
type StakedETHLidoLocatorSet struct {
	LidoLocator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLidoLocatorSet is a free log retrieval operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StakedETH *StakedETHFilterer) FilterLidoLocatorSet(opts *bind.FilterOpts) (*StakedETHLidoLocatorSetIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return &StakedETHLidoLocatorSetIterator{contract: _StakedETH.contract, event: "LidoLocatorSet", logs: logs, sub: sub}, nil
}

// WatchLidoLocatorSet is a free log subscription operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StakedETH *StakedETHFilterer) WatchLidoLocatorSet(opts *bind.WatchOpts, sink chan<- *StakedETHLidoLocatorSet) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "LidoLocatorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHLidoLocatorSet)
				if err := _StakedETH.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
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

// ParseLidoLocatorSet is a log parse operation binding the contract event 0x61f9416d3c29deb4e424342445a2b132738430becd9fa275e11297c90668b22e.
//
// Solidity: event LidoLocatorSet(address lidoLocator)
func (_StakedETH *StakedETHFilterer) ParseLidoLocatorSet(log types.Log) (*StakedETHLidoLocatorSet, error) {
	event := new(StakedETHLidoLocatorSet)
	if err := _StakedETH.contract.UnpackLog(event, "LidoLocatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the StakedETH contract.
type StakedETHRecoverToVaultIterator struct {
	Event *StakedETHRecoverToVault // Event containing the contract specifics and raw log

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
func (it *StakedETHRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHRecoverToVault)
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
		it.Event = new(StakedETHRecoverToVault)
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
func (it *StakedETHRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHRecoverToVault represents a RecoverToVault event raised by the StakedETH contract.
type StakedETHRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StakedETH *StakedETHFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*StakedETHRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHRecoverToVaultIterator{contract: _StakedETH.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StakedETH *StakedETHFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *StakedETHRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHRecoverToVault)
				if err := _StakedETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
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

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_StakedETH *StakedETHFilterer) ParseRecoverToVault(log types.Log) (*StakedETHRecoverToVault, error) {
	event := new(StakedETHRecoverToVault)
	if err := _StakedETH.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the StakedETH contract.
type StakedETHResumedIterator struct {
	Event *StakedETHResumed // Event containing the contract specifics and raw log

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
func (it *StakedETHResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHResumed)
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
		it.Event = new(StakedETHResumed)
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
func (it *StakedETHResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHResumed represents a Resumed event raised by the StakedETH contract.
type StakedETHResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StakedETH *StakedETHFilterer) FilterResumed(opts *bind.FilterOpts) (*StakedETHResumedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &StakedETHResumedIterator{contract: _StakedETH.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StakedETH *StakedETHFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *StakedETHResumed) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHResumed)
				if err := _StakedETH.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StakedETH *StakedETHFilterer) ParseResumed(log types.Log) (*StakedETHResumed, error) {
	event := new(StakedETHResumed)
	if err := _StakedETH.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the StakedETH contract.
type StakedETHScriptResultIterator struct {
	Event *StakedETHScriptResult // Event containing the contract specifics and raw log

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
func (it *StakedETHScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHScriptResult)
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
		it.Event = new(StakedETHScriptResult)
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
func (it *StakedETHScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHScriptResult represents a ScriptResult event raised by the StakedETH contract.
type StakedETHScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StakedETH *StakedETHFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*StakedETHScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHScriptResultIterator{contract: _StakedETH.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StakedETH *StakedETHFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *StakedETHScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHScriptResult)
				if err := _StakedETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
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

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_StakedETH *StakedETHFilterer) ParseScriptResult(log types.Log) (*StakedETHScriptResult, error) {
	event := new(StakedETHScriptResult)
	if err := _StakedETH.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHSharesBurntIterator is returned from FilterSharesBurnt and is used to iterate over the raw logs and unpacked data for SharesBurnt events raised by the StakedETH contract.
type StakedETHSharesBurntIterator struct {
	Event *StakedETHSharesBurnt // Event containing the contract specifics and raw log

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
func (it *StakedETHSharesBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHSharesBurnt)
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
		it.Event = new(StakedETHSharesBurnt)
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
func (it *StakedETHSharesBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHSharesBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHSharesBurnt represents a SharesBurnt event raised by the StakedETH contract.
type StakedETHSharesBurnt struct {
	Account               common.Address
	PreRebaseTokenAmount  *big.Int
	PostRebaseTokenAmount *big.Int
	SharesAmount          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSharesBurnt is a free log retrieval operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StakedETH *StakedETHFilterer) FilterSharesBurnt(opts *bind.FilterOpts, account []common.Address) (*StakedETHSharesBurntIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHSharesBurntIterator{contract: _StakedETH.contract, event: "SharesBurnt", logs: logs, sub: sub}, nil
}

// WatchSharesBurnt is a free log subscription operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StakedETH *StakedETHFilterer) WatchSharesBurnt(opts *bind.WatchOpts, sink chan<- *StakedETHSharesBurnt, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "SharesBurnt", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHSharesBurnt)
				if err := _StakedETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
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

// ParseSharesBurnt is a log parse operation binding the contract event 0x8b2a1e1ad5e0578c3dd82494156e985dade827a87c573b5c1c7716a32162ad64.
//
// Solidity: event SharesBurnt(address indexed account, uint256 preRebaseTokenAmount, uint256 postRebaseTokenAmount, uint256 sharesAmount)
func (_StakedETH *StakedETHFilterer) ParseSharesBurnt(log types.Log) (*StakedETHSharesBurnt, error) {
	event := new(StakedETHSharesBurnt)
	if err := _StakedETH.contract.UnpackLog(event, "SharesBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHStakingLimitRemovedIterator is returned from FilterStakingLimitRemoved and is used to iterate over the raw logs and unpacked data for StakingLimitRemoved events raised by the StakedETH contract.
type StakedETHStakingLimitRemovedIterator struct {
	Event *StakedETHStakingLimitRemoved // Event containing the contract specifics and raw log

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
func (it *StakedETHStakingLimitRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHStakingLimitRemoved)
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
		it.Event = new(StakedETHStakingLimitRemoved)
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
func (it *StakedETHStakingLimitRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHStakingLimitRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHStakingLimitRemoved represents a StakingLimitRemoved event raised by the StakedETH contract.
type StakedETHStakingLimitRemoved struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitRemoved is a free log retrieval operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StakedETH *StakedETHFilterer) FilterStakingLimitRemoved(opts *bind.FilterOpts) (*StakedETHStakingLimitRemovedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return &StakedETHStakingLimitRemovedIterator{contract: _StakedETH.contract, event: "StakingLimitRemoved", logs: logs, sub: sub}, nil
}

// WatchStakingLimitRemoved is a free log subscription operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StakedETH *StakedETHFilterer) WatchStakingLimitRemoved(opts *bind.WatchOpts, sink chan<- *StakedETHStakingLimitRemoved) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "StakingLimitRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHStakingLimitRemoved)
				if err := _StakedETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
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

// ParseStakingLimitRemoved is a log parse operation binding the contract event 0x9b2a687c198898fcc32a33bbc610d478f177a73ab7352023e6cc1de5bf12a3df.
//
// Solidity: event StakingLimitRemoved()
func (_StakedETH *StakedETHFilterer) ParseStakingLimitRemoved(log types.Log) (*StakedETHStakingLimitRemoved, error) {
	event := new(StakedETHStakingLimitRemoved)
	if err := _StakedETH.contract.UnpackLog(event, "StakingLimitRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHStakingLimitSetIterator is returned from FilterStakingLimitSet and is used to iterate over the raw logs and unpacked data for StakingLimitSet events raised by the StakedETH contract.
type StakedETHStakingLimitSetIterator struct {
	Event *StakedETHStakingLimitSet // Event containing the contract specifics and raw log

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
func (it *StakedETHStakingLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHStakingLimitSet)
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
		it.Event = new(StakedETHStakingLimitSet)
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
func (it *StakedETHStakingLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHStakingLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHStakingLimitSet represents a StakingLimitSet event raised by the StakedETH contract.
type StakedETHStakingLimitSet struct {
	MaxStakeLimit              *big.Int
	StakeLimitIncreasePerBlock *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingLimitSet is a free log retrieval operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StakedETH *StakedETHFilterer) FilterStakingLimitSet(opts *bind.FilterOpts) (*StakedETHStakingLimitSetIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return &StakedETHStakingLimitSetIterator{contract: _StakedETH.contract, event: "StakingLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingLimitSet is a free log subscription operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StakedETH *StakedETHFilterer) WatchStakingLimitSet(opts *bind.WatchOpts, sink chan<- *StakedETHStakingLimitSet) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "StakingLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHStakingLimitSet)
				if err := _StakedETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
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

// ParseStakingLimitSet is a log parse operation binding the contract event 0xce9fddf6179affa1ea7bf36d80a6bf0284e0f3b91f4b2fa6eea2af923e7fac2d.
//
// Solidity: event StakingLimitSet(uint256 maxStakeLimit, uint256 stakeLimitIncreasePerBlock)
func (_StakedETH *StakedETHFilterer) ParseStakingLimitSet(log types.Log) (*StakedETHStakingLimitSet, error) {
	event := new(StakedETHStakingLimitSet)
	if err := _StakedETH.contract.UnpackLog(event, "StakingLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHStakingPausedIterator is returned from FilterStakingPaused and is used to iterate over the raw logs and unpacked data for StakingPaused events raised by the StakedETH contract.
type StakedETHStakingPausedIterator struct {
	Event *StakedETHStakingPaused // Event containing the contract specifics and raw log

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
func (it *StakedETHStakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHStakingPaused)
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
		it.Event = new(StakedETHStakingPaused)
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
func (it *StakedETHStakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHStakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHStakingPaused represents a StakingPaused event raised by the StakedETH contract.
type StakedETHStakingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingPaused is a free log retrieval operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StakedETH *StakedETHFilterer) FilterStakingPaused(opts *bind.FilterOpts) (*StakedETHStakingPausedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return &StakedETHStakingPausedIterator{contract: _StakedETH.contract, event: "StakingPaused", logs: logs, sub: sub}, nil
}

// WatchStakingPaused is a free log subscription operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StakedETH *StakedETHFilterer) WatchStakingPaused(opts *bind.WatchOpts, sink chan<- *StakedETHStakingPaused) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "StakingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHStakingPaused)
				if err := _StakedETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
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

// ParseStakingPaused is a log parse operation binding the contract event 0x26d1807b479eaba249c1214b82e4b65bbb0cc73ee8a17901324b1ef1b5904e49.
//
// Solidity: event StakingPaused()
func (_StakedETH *StakedETHFilterer) ParseStakingPaused(log types.Log) (*StakedETHStakingPaused, error) {
	event := new(StakedETHStakingPaused)
	if err := _StakedETH.contract.UnpackLog(event, "StakingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHStakingResumedIterator is returned from FilterStakingResumed and is used to iterate over the raw logs and unpacked data for StakingResumed events raised by the StakedETH contract.
type StakedETHStakingResumedIterator struct {
	Event *StakedETHStakingResumed // Event containing the contract specifics and raw log

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
func (it *StakedETHStakingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHStakingResumed)
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
		it.Event = new(StakedETHStakingResumed)
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
func (it *StakedETHStakingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHStakingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHStakingResumed represents a StakingResumed event raised by the StakedETH contract.
type StakedETHStakingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStakingResumed is a free log retrieval operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StakedETH *StakedETHFilterer) FilterStakingResumed(opts *bind.FilterOpts) (*StakedETHStakingResumedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return &StakedETHStakingResumedIterator{contract: _StakedETH.contract, event: "StakingResumed", logs: logs, sub: sub}, nil
}

// WatchStakingResumed is a free log subscription operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StakedETH *StakedETHFilterer) WatchStakingResumed(opts *bind.WatchOpts, sink chan<- *StakedETHStakingResumed) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "StakingResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHStakingResumed)
				if err := _StakedETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
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

// ParseStakingResumed is a log parse operation binding the contract event 0xedaeeae9aed70c4545d3ab0065713261c9cee8d6cf5c8b07f52f0a65fd91efda.
//
// Solidity: event StakingResumed()
func (_StakedETH *StakedETHFilterer) ParseStakingResumed(log types.Log) (*StakedETHStakingResumed, error) {
	event := new(StakedETHStakingResumed)
	if err := _StakedETH.contract.UnpackLog(event, "StakingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHStoppedIterator is returned from FilterStopped and is used to iterate over the raw logs and unpacked data for Stopped events raised by the StakedETH contract.
type StakedETHStoppedIterator struct {
	Event *StakedETHStopped // Event containing the contract specifics and raw log

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
func (it *StakedETHStoppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHStopped)
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
		it.Event = new(StakedETHStopped)
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
func (it *StakedETHStoppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHStoppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHStopped represents a Stopped event raised by the StakedETH contract.
type StakedETHStopped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterStopped is a free log retrieval operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StakedETH *StakedETHFilterer) FilterStopped(opts *bind.FilterOpts) (*StakedETHStoppedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return &StakedETHStoppedIterator{contract: _StakedETH.contract, event: "Stopped", logs: logs, sub: sub}, nil
}

// WatchStopped is a free log subscription operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StakedETH *StakedETHFilterer) WatchStopped(opts *bind.WatchOpts, sink chan<- *StakedETHStopped) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Stopped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHStopped)
				if err := _StakedETH.contract.UnpackLog(event, "Stopped", log); err != nil {
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

// ParseStopped is a log parse operation binding the contract event 0x7acc84e34091ae817647a4c49116f5cc07f319078ba80f8f5fde37ea7e25cbd6.
//
// Solidity: event Stopped()
func (_StakedETH *StakedETHFilterer) ParseStopped(log types.Log) (*StakedETHStopped, error) {
	event := new(StakedETHStopped)
	if err := _StakedETH.contract.UnpackLog(event, "Stopped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the StakedETH contract.
type StakedETHSubmittedIterator struct {
	Event *StakedETHSubmitted // Event containing the contract specifics and raw log

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
func (it *StakedETHSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHSubmitted)
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
		it.Event = new(StakedETHSubmitted)
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
func (it *StakedETHSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHSubmitted represents a Submitted event raised by the StakedETH contract.
type StakedETHSubmitted struct {
	Sender   common.Address
	Amount   *big.Int
	Referral common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StakedETH *StakedETHFilterer) FilterSubmitted(opts *bind.FilterOpts, sender []common.Address) (*StakedETHSubmittedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHSubmittedIterator{contract: _StakedETH.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StakedETH *StakedETHFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *StakedETHSubmitted, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Submitted", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHSubmitted)
				if err := _StakedETH.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x96a25c8ce0baabc1fdefd93e9ed25d8e092a3332f3aa9a41722b5697231d1d1a.
//
// Solidity: event Submitted(address indexed sender, uint256 amount, address referral)
func (_StakedETH *StakedETHFilterer) ParseSubmitted(log types.Log) (*StakedETHSubmitted, error) {
	event := new(StakedETHSubmitted)
	if err := _StakedETH.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHTokenRebasedIterator is returned from FilterTokenRebased and is used to iterate over the raw logs and unpacked data for TokenRebased events raised by the StakedETH contract.
type StakedETHTokenRebasedIterator struct {
	Event *StakedETHTokenRebased // Event containing the contract specifics and raw log

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
func (it *StakedETHTokenRebasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHTokenRebased)
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
		it.Event = new(StakedETHTokenRebased)
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
func (it *StakedETHTokenRebasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHTokenRebasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHTokenRebased represents a TokenRebased event raised by the StakedETH contract.
type StakedETHTokenRebased struct {
	ReportTimestamp    *big.Int
	TimeElapsed        *big.Int
	PreTotalShares     *big.Int
	PreTotalEther      *big.Int
	PostTotalShares    *big.Int
	PostTotalEther     *big.Int
	SharesMintedAsFees *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTokenRebased is a free log retrieval operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StakedETH *StakedETHFilterer) FilterTokenRebased(opts *bind.FilterOpts, reportTimestamp []*big.Int) (*StakedETHTokenRebasedIterator, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHTokenRebasedIterator{contract: _StakedETH.contract, event: "TokenRebased", logs: logs, sub: sub}, nil
}

// WatchTokenRebased is a free log subscription operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StakedETH *StakedETHFilterer) WatchTokenRebased(opts *bind.WatchOpts, sink chan<- *StakedETHTokenRebased, reportTimestamp []*big.Int) (event.Subscription, error) {

	var reportTimestampRule []interface{}
	for _, reportTimestampItem := range reportTimestamp {
		reportTimestampRule = append(reportTimestampRule, reportTimestampItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "TokenRebased", reportTimestampRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHTokenRebased)
				if err := _StakedETH.contract.UnpackLog(event, "TokenRebased", log); err != nil {
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

// ParseTokenRebased is a log parse operation binding the contract event 0xff08c3ef606d198e316ef5b822193c489965899eb4e3c248cea1a4626c3eda50.
//
// Solidity: event TokenRebased(uint256 indexed reportTimestamp, uint256 timeElapsed, uint256 preTotalShares, uint256 preTotalEther, uint256 postTotalShares, uint256 postTotalEther, uint256 sharesMintedAsFees)
func (_StakedETH *StakedETHFilterer) ParseTokenRebased(log types.Log) (*StakedETHTokenRebased, error) {
	event := new(StakedETHTokenRebased)
	if err := _StakedETH.contract.UnpackLog(event, "TokenRebased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakedETH contract.
type StakedETHTransferIterator struct {
	Event *StakedETHTransfer // Event containing the contract specifics and raw log

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
func (it *StakedETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHTransfer)
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
		it.Event = new(StakedETHTransfer)
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
func (it *StakedETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHTransfer represents a Transfer event raised by the StakedETH contract.
type StakedETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedETH *StakedETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakedETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHTransferIterator{contract: _StakedETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StakedETH *StakedETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakedETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHTransfer)
				if err := _StakedETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakedETH *StakedETHFilterer) ParseTransfer(log types.Log) (*StakedETHTransfer, error) {
	event := new(StakedETHTransfer)
	if err := _StakedETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHTransferSharesIterator is returned from FilterTransferShares and is used to iterate over the raw logs and unpacked data for TransferShares events raised by the StakedETH contract.
type StakedETHTransferSharesIterator struct {
	Event *StakedETHTransferShares // Event containing the contract specifics and raw log

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
func (it *StakedETHTransferSharesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHTransferShares)
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
		it.Event = new(StakedETHTransferShares)
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
func (it *StakedETHTransferSharesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHTransferSharesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHTransferShares represents a TransferShares event raised by the StakedETH contract.
type StakedETHTransferShares struct {
	From        common.Address
	To          common.Address
	SharesValue *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferShares is a free log retrieval operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StakedETH *StakedETHFilterer) FilterTransferShares(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakedETHTransferSharesIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHTransferSharesIterator{contract: _StakedETH.contract, event: "TransferShares", logs: logs, sub: sub}, nil
}

// WatchTransferShares is a free log subscription operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StakedETH *StakedETHFilterer) WatchTransferShares(opts *bind.WatchOpts, sink chan<- *StakedETHTransferShares, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "TransferShares", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHTransferShares)
				if err := _StakedETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
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

// ParseTransferShares is a log parse operation binding the contract event 0x9d9c909296d9c674451c0c24f02cb64981eb3b727f99865939192f880a755dcb.
//
// Solidity: event TransferShares(address indexed from, address indexed to, uint256 sharesValue)
func (_StakedETH *StakedETHFilterer) ParseTransferShares(log types.Log) (*StakedETHTransferShares, error) {
	event := new(StakedETHTransferShares)
	if err := _StakedETH.contract.UnpackLog(event, "TransferShares", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHUnbufferedIterator is returned from FilterUnbuffered and is used to iterate over the raw logs and unpacked data for Unbuffered events raised by the StakedETH contract.
type StakedETHUnbufferedIterator struct {
	Event *StakedETHUnbuffered // Event containing the contract specifics and raw log

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
func (it *StakedETHUnbufferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHUnbuffered)
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
		it.Event = new(StakedETHUnbuffered)
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
func (it *StakedETHUnbufferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHUnbufferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHUnbuffered represents a Unbuffered event raised by the StakedETH contract.
type StakedETHUnbuffered struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbuffered is a free log retrieval operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StakedETH *StakedETHFilterer) FilterUnbuffered(opts *bind.FilterOpts) (*StakedETHUnbufferedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return &StakedETHUnbufferedIterator{contract: _StakedETH.contract, event: "Unbuffered", logs: logs, sub: sub}, nil
}

// WatchUnbuffered is a free log subscription operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StakedETH *StakedETHFilterer) WatchUnbuffered(opts *bind.WatchOpts, sink chan<- *StakedETHUnbuffered) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "Unbuffered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHUnbuffered)
				if err := _StakedETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
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

// ParseUnbuffered is a log parse operation binding the contract event 0x76a397bea5768d4fca97ef47792796e35f98dc81b16c1de84e28a818e1f97108.
//
// Solidity: event Unbuffered(uint256 amount)
func (_StakedETH *StakedETHFilterer) ParseUnbuffered(log types.Log) (*StakedETHUnbuffered, error) {
	event := new(StakedETHUnbuffered)
	if err := _StakedETH.contract.UnpackLog(event, "Unbuffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalsReceivedIterator is returned from FilterWithdrawalsReceived and is used to iterate over the raw logs and unpacked data for WithdrawalsReceived events raised by the StakedETH contract.
type StakedETHWithdrawalsReceivedIterator struct {
	Event *StakedETHWithdrawalsReceived // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalsReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalsReceived)
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
		it.Event = new(StakedETHWithdrawalsReceived)
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
func (it *StakedETHWithdrawalsReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalsReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalsReceived represents a WithdrawalsReceived event raised by the StakedETH contract.
type StakedETHWithdrawalsReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsReceived is a free log retrieval operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) FilterWithdrawalsReceived(opts *bind.FilterOpts) (*StakedETHWithdrawalsReceivedIterator, error) {

	logs, sub, err := _StakedETH.contract.FilterLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalsReceivedIterator{contract: _StakedETH.contract, event: "WithdrawalsReceived", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsReceived is a free log subscription operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) WatchWithdrawalsReceived(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalsReceived) (event.Subscription, error) {

	logs, sub, err := _StakedETH.contract.WatchLogs(opts, "WithdrawalsReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalsReceived)
				if err := _StakedETH.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
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

// ParseWithdrawalsReceived is a log parse operation binding the contract event 0x6e5086f7e1ab04bd826e77faae35b1bcfe31bd144623361a40ea4af51670b1c3.
//
// Solidity: event WithdrawalsReceived(uint256 amount)
func (_StakedETH *StakedETHFilterer) ParseWithdrawalsReceived(log types.Log) (*StakedETHWithdrawalsReceived, error) {
	event := new(StakedETHWithdrawalsReceived)
	if err := _StakedETH.contract.UnpackLog(event, "WithdrawalsReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
