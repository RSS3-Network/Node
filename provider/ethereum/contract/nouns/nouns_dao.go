// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nouns

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

// NounsDAOStorageV1AdjustedReceipt is an auto generated low-level Go binding around an user-defined struct.
type NounsDAOStorageV1AdjustedReceipt struct {
	HasVoted bool
	Support  uint8
	Votes    *big.Int
}

// NounsDAOStorageV2DynamicQuorumParams is an auto generated low-level Go binding around an user-defined struct.
type NounsDAOStorageV2DynamicQuorumParams struct {
	MinQuorumVotesBPS uint16
	MaxQuorumVotesBPS uint16
	QuorumCoefficient uint32
}

// NounsDAOStorageV2ProposalCondensed is an auto generated low-level Go binding around an user-defined struct.
type NounsDAOStorageV2ProposalCondensed struct {
	Id                *big.Int
	Proposer          common.Address
	ProposalThreshold *big.Int
	QuorumVotes       *big.Int
	Eta               *big.Int
	StartBlock        *big.Int
	EndBlock          *big.Int
	ForVotes          *big.Int
	AgainstVotes      *big.Int
	AbstainVotes      *big.Int
	Canceled          bool
	Vetoed            bool
	Executed          bool
	TotalSupply       *big.Int
	CreationBlock     *big.Int
}

// NounsDAOMetaData contains all meta data concerning the NounsDAO contract.
var NounsDAOMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AdminOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantCancelExecutedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantVetoExecutedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxQuorumVotesBPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinQuorumVotesBPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinQuorumBPSGreaterThanMaxQuorumBPS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingVetoerOnly\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeUint16Cast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VetoerBurned\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VetoerOnly\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldMaxQuorumVotesBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMaxQuorumVotesBPS\",\"type\":\"uint16\"}],\"name\":\"MaxQuorumVotesBPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oldMinQuorumVotesBPS\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"newMinQuorumVotesBPS\",\"type\":\"uint16\"}],\"name\":\"MinQuorumVotesBPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingVetoer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingVetoer\",\"type\":\"address\"}],\"name\":\"NewPendingVetoer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVetoer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVetoer\",\"type\":\"address\"}],\"name\":\"NewVetoer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quorumVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreatedWithRequirements\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThresholdBPS\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThresholdBPS\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdBPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalVetoed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldQuorumCoefficient\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newQuorumCoefficient\",\"type\":\"uint32\"}],\"name\":\"QuorumCoefficientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumVotesBPS\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumVotesBPS\",\"type\":\"uint256\"}],\"name\":\"QuorumVotesBPSSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"refundSent\",\"type\":\"bool\"}],\"name\":\"RefundableVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sent\",\"type\":\"bool\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PROPOSAL_THRESHOLD_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_QUORUM_VOTES_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_QUORUM_VOTES_BPS_UPPER_BOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REFUND_BASE_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REFUND_GAS_USED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REFUND_PRIORITY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VOTING_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_VOTING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_PROPOSAL_THRESHOLD_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_QUORUM_VOTES_BPS_LOWER_BOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_QUORUM_VOTES_BPS_UPPER_BOUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_VOTING_DELAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_VOTING_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REFUND_BASE_GAS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptVetoer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_burnVetoPower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMinQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newMaxQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"newQuorumCoefficient\",\"type\":\"uint32\"}],\"name\":\"_setDynamicQuorumParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMaxQuorumVotesBPS\",\"type\":\"uint16\"}],\"name\":\"_setMaxQuorumVotesBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newMinQuorumVotesBPS\",\"type\":\"uint16\"}],\"name\":\"_setMinQuorumVotesBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingVetoer\",\"type\":\"address\"}],\"name\":\"_setPendingVetoer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThresholdBPS\",\"type\":\"uint256\"}],\"name\":\"_setProposalThresholdBPS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newQuorumCoefficient\",\"type\":\"uint32\"}],\"name\":\"_setQuorumCoefficient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"_setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"_setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castRefundableVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castRefundableVoteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"castVoteBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"quorumCoefficient\",\"type\":\"uint32\"}],\"internalType\":\"structNounsDAOStorageV2.DynamicQuorumParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"dynamicQuorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getActions\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber_\",\"type\":\"uint256\"}],\"name\":\"getDynamicQuorumParamsAt\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"quorumCoefficient\",\"type\":\"uint32\"}],\"internalType\":\"structNounsDAOStorageV2.DynamicQuorumParams\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"getReceipt\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"uint96\",\"name\":\"votes\",\"type\":\"uint96\"}],\"internalType\":\"structNounsDAOStorageV1Adjusted.Receipt\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"timelock_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nouns_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vetoer_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votingPeriod_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"votingDelay_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalThresholdBPS_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"quorumCoefficient\",\"type\":\"uint32\"}],\"internalType\":\"structNounsDAOStorageV2.DynamicQuorumParams\",\"name\":\"dynamicQuorumParams_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"latestProposalIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxQuorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minQuorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nouns\",\"outputs\":[{\"internalType\":\"contractNounsTokenLike\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingVetoer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalMaxOperations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThresholdBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"proposalThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"canceled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"vetoed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creationBlock\",\"type\":\"uint256\"}],\"internalType\":\"structNounsDAOStorageV2.ProposalCondensed\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"queue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"quorumParamsCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"minQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"maxQuorumVotesBPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"quorumCoefficient\",\"type\":\"uint32\"}],\"internalType\":\"structNounsDAOStorageV2.DynamicQuorumParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"quorumVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumVotesBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumNounsDAOStorageV1Adjusted.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"contractINounsDAOExecutor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"veto\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vetoer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// NounsDAOABI is the input ABI used to generate the binding from.
// Deprecated: Use NounsDAOMetaData.ABI instead.
var NounsDAOABI = NounsDAOMetaData.ABI

// NounsDAO is an auto generated Go binding around an Ethereum contract.
type NounsDAO struct {
	NounsDAOCaller     // Read-only binding to the contract
	NounsDAOTransactor // Write-only binding to the contract
	NounsDAOFilterer   // Log filterer for contract events
}

// NounsDAOCaller is an auto generated read-only Go binding around an Ethereum contract.
type NounsDAOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsDAOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NounsDAOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsDAOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NounsDAOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NounsDAOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NounsDAOSession struct {
	Contract     *NounsDAO         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NounsDAOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NounsDAOCallerSession struct {
	Contract *NounsDAOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NounsDAOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NounsDAOTransactorSession struct {
	Contract     *NounsDAOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NounsDAORaw is an auto generated low-level Go binding around an Ethereum contract.
type NounsDAORaw struct {
	Contract *NounsDAO // Generic contract binding to access the raw methods on
}

// NounsDAOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NounsDAOCallerRaw struct {
	Contract *NounsDAOCaller // Generic read-only contract binding to access the raw methods on
}

// NounsDAOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NounsDAOTransactorRaw struct {
	Contract *NounsDAOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNounsDAO creates a new instance of NounsDAO, bound to a specific deployed contract.
func NewNounsDAO(address common.Address, backend bind.ContractBackend) (*NounsDAO, error) {
	contract, err := bindNounsDAO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NounsDAO{NounsDAOCaller: NounsDAOCaller{contract: contract}, NounsDAOTransactor: NounsDAOTransactor{contract: contract}, NounsDAOFilterer: NounsDAOFilterer{contract: contract}}, nil
}

// NewNounsDAOCaller creates a new read-only instance of NounsDAO, bound to a specific deployed contract.
func NewNounsDAOCaller(address common.Address, caller bind.ContractCaller) (*NounsDAOCaller, error) {
	contract, err := bindNounsDAO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NounsDAOCaller{contract: contract}, nil
}

// NewNounsDAOTransactor creates a new write-only instance of NounsDAO, bound to a specific deployed contract.
func NewNounsDAOTransactor(address common.Address, transactor bind.ContractTransactor) (*NounsDAOTransactor, error) {
	contract, err := bindNounsDAO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NounsDAOTransactor{contract: contract}, nil
}

// NewNounsDAOFilterer creates a new log filterer instance of NounsDAO, bound to a specific deployed contract.
func NewNounsDAOFilterer(address common.Address, filterer bind.ContractFilterer) (*NounsDAOFilterer, error) {
	contract, err := bindNounsDAO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NounsDAOFilterer{contract: contract}, nil
}

// bindNounsDAO binds a generic wrapper to an already deployed contract.
func bindNounsDAO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NounsDAOMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsDAO *NounsDAORaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsDAO.Contract.NounsDAOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsDAO *NounsDAORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.Contract.NounsDAOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsDAO *NounsDAORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsDAO.Contract.NounsDAOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NounsDAO *NounsDAOCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NounsDAO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NounsDAO *NounsDAOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NounsDAO *NounsDAOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NounsDAO.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _NounsDAO.Contract.BALLOTTYPEHASH(&_NounsDAO.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _NounsDAO.Contract.BALLOTTYPEHASH(&_NounsDAO.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NounsDAO.Contract.DOMAINTYPEHASH(&_NounsDAO.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_NounsDAO *NounsDAOCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _NounsDAO.Contract.DOMAINTYPEHASH(&_NounsDAO.CallOpts)
}

// MAXPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0x1e7b5d3a.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXPROPOSALTHRESHOLDBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_PROPOSAL_THRESHOLD_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0x1e7b5d3a.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXPROPOSALTHRESHOLDBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MAXPROPOSALTHRESHOLDBPS(&_NounsDAO.CallOpts)
}

// MAXPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0x1e7b5d3a.
//
// Solidity: function MAX_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXPROPOSALTHRESHOLDBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MAXPROPOSALTHRESHOLDBPS(&_NounsDAO.CallOpts)
}

// MAXQUORUMVOTESBPS is a free data retrieval call binding the contract method 0xbb677582.
//
// Solidity: function MAX_QUORUM_VOTES_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXQUORUMVOTESBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_QUORUM_VOTES_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXQUORUMVOTESBPS is a free data retrieval call binding the contract method 0xbb677582.
//
// Solidity: function MAX_QUORUM_VOTES_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXQUORUMVOTESBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MAXQUORUMVOTESBPS(&_NounsDAO.CallOpts)
}

// MAXQUORUMVOTESBPS is a free data retrieval call binding the contract method 0xbb677582.
//
// Solidity: function MAX_QUORUM_VOTES_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXQUORUMVOTESBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MAXQUORUMVOTESBPS(&_NounsDAO.CallOpts)
}

// MAXQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0x26e6dcb0.
//
// Solidity: function MAX_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXQUORUMVOTESBPSUPPERBOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_QUORUM_VOTES_BPS_UPPER_BOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0x26e6dcb0.
//
// Solidity: function MAX_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXQUORUMVOTESBPSUPPERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MAXQUORUMVOTESBPSUPPERBOUND(&_NounsDAO.CallOpts)
}

// MAXQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0x26e6dcb0.
//
// Solidity: function MAX_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXQUORUMVOTESBPSUPPERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MAXQUORUMVOTESBPSUPPERBOUND(&_NounsDAO.CallOpts)
}

// MAXREFUNDBASEFEE is a free data retrieval call binding the contract method 0xbc4cd084.
//
// Solidity: function MAX_REFUND_BASE_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXREFUNDBASEFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_REFUND_BASE_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFUNDBASEFEE is a free data retrieval call binding the contract method 0xbc4cd084.
//
// Solidity: function MAX_REFUND_BASE_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXREFUNDBASEFEE() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDBASEFEE(&_NounsDAO.CallOpts)
}

// MAXREFUNDBASEFEE is a free data retrieval call binding the contract method 0xbc4cd084.
//
// Solidity: function MAX_REFUND_BASE_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXREFUNDBASEFEE() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDBASEFEE(&_NounsDAO.CallOpts)
}

// MAXREFUNDGASUSED is a free data retrieval call binding the contract method 0x042bc3de.
//
// Solidity: function MAX_REFUND_GAS_USED() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXREFUNDGASUSED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_REFUND_GAS_USED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFUNDGASUSED is a free data retrieval call binding the contract method 0x042bc3de.
//
// Solidity: function MAX_REFUND_GAS_USED() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXREFUNDGASUSED() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDGASUSED(&_NounsDAO.CallOpts)
}

// MAXREFUNDGASUSED is a free data retrieval call binding the contract method 0x042bc3de.
//
// Solidity: function MAX_REFUND_GAS_USED() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXREFUNDGASUSED() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDGASUSED(&_NounsDAO.CallOpts)
}

// MAXREFUNDPRIORITYFEE is a free data retrieval call binding the contract method 0xfbfee876.
//
// Solidity: function MAX_REFUND_PRIORITY_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXREFUNDPRIORITYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_REFUND_PRIORITY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXREFUNDPRIORITYFEE is a free data retrieval call binding the contract method 0xfbfee876.
//
// Solidity: function MAX_REFUND_PRIORITY_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXREFUNDPRIORITYFEE() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDPRIORITYFEE(&_NounsDAO.CallOpts)
}

// MAXREFUNDPRIORITYFEE is a free data retrieval call binding the contract method 0xfbfee876.
//
// Solidity: function MAX_REFUND_PRIORITY_FEE() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXREFUNDPRIORITYFEE() (*big.Int, error) {
	return _NounsDAO.Contract.MAXREFUNDPRIORITYFEE(&_NounsDAO.CallOpts)
}

// MAXVOTINGDELAY is a free data retrieval call binding the contract method 0xb1126263.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXVOTINGDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_VOTING_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVOTINGDELAY is a free data retrieval call binding the contract method 0xb1126263.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXVOTINGDELAY() (*big.Int, error) {
	return _NounsDAO.Contract.MAXVOTINGDELAY(&_NounsDAO.CallOpts)
}

// MAXVOTINGDELAY is a free data retrieval call binding the contract method 0xb1126263.
//
// Solidity: function MAX_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXVOTINGDELAY() (*big.Int, error) {
	return _NounsDAO.Contract.MAXVOTINGDELAY(&_NounsDAO.CallOpts)
}

// MAXVOTINGPERIOD is a free data retrieval call binding the contract method 0xa64e024a.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MAXVOTINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MAX_VOTING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXVOTINGPERIOD is a free data retrieval call binding the contract method 0xa64e024a.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MAXVOTINGPERIOD() (*big.Int, error) {
	return _NounsDAO.Contract.MAXVOTINGPERIOD(&_NounsDAO.CallOpts)
}

// MAXVOTINGPERIOD is a free data retrieval call binding the contract method 0xa64e024a.
//
// Solidity: function MAX_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MAXVOTINGPERIOD() (*big.Int, error) {
	return _NounsDAO.Contract.MAXVOTINGPERIOD(&_NounsDAO.CallOpts)
}

// MINPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0xc82fbd08.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MINPROPOSALTHRESHOLDBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MIN_PROPOSAL_THRESHOLD_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0xc82fbd08.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MINPROPOSALTHRESHOLDBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MINPROPOSALTHRESHOLDBPS(&_NounsDAO.CallOpts)
}

// MINPROPOSALTHRESHOLDBPS is a free data retrieval call binding the contract method 0xc82fbd08.
//
// Solidity: function MIN_PROPOSAL_THRESHOLD_BPS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MINPROPOSALTHRESHOLDBPS() (*big.Int, error) {
	return _NounsDAO.Contract.MINPROPOSALTHRESHOLDBPS(&_NounsDAO.CallOpts)
}

// MINQUORUMVOTESBPSLOWERBOUND is a free data retrieval call binding the contract method 0xccfa51c9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_LOWER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MINQUORUMVOTESBPSLOWERBOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MIN_QUORUM_VOTES_BPS_LOWER_BOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINQUORUMVOTESBPSLOWERBOUND is a free data retrieval call binding the contract method 0xccfa51c9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_LOWER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MINQUORUMVOTESBPSLOWERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MINQUORUMVOTESBPSLOWERBOUND(&_NounsDAO.CallOpts)
}

// MINQUORUMVOTESBPSLOWERBOUND is a free data retrieval call binding the contract method 0xccfa51c9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_LOWER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MINQUORUMVOTESBPSLOWERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MINQUORUMVOTESBPSLOWERBOUND(&_NounsDAO.CallOpts)
}

// MINQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0xe7951bb9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MINQUORUMVOTESBPSUPPERBOUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MIN_QUORUM_VOTES_BPS_UPPER_BOUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0xe7951bb9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MINQUORUMVOTESBPSUPPERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MINQUORUMVOTESBPSUPPERBOUND(&_NounsDAO.CallOpts)
}

// MINQUORUMVOTESBPSUPPERBOUND is a free data retrieval call binding the contract method 0xe7951bb9.
//
// Solidity: function MIN_QUORUM_VOTES_BPS_UPPER_BOUND() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MINQUORUMVOTESBPSUPPERBOUND() (*big.Int, error) {
	return _NounsDAO.Contract.MINQUORUMVOTESBPSUPPERBOUND(&_NounsDAO.CallOpts)
}

// MINVOTINGDELAY is a free data retrieval call binding the contract method 0xe48083fe.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MINVOTINGDELAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MIN_VOTING_DELAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINVOTINGDELAY is a free data retrieval call binding the contract method 0xe48083fe.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MINVOTINGDELAY() (*big.Int, error) {
	return _NounsDAO.Contract.MINVOTINGDELAY(&_NounsDAO.CallOpts)
}

// MINVOTINGDELAY is a free data retrieval call binding the contract method 0xe48083fe.
//
// Solidity: function MIN_VOTING_DELAY() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MINVOTINGDELAY() (*big.Int, error) {
	return _NounsDAO.Contract.MINVOTINGDELAY(&_NounsDAO.CallOpts)
}

// MINVOTINGPERIOD is a free data retrieval call binding the contract method 0x215809ca.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MINVOTINGPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "MIN_VOTING_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINVOTINGPERIOD is a free data retrieval call binding the contract method 0x215809ca.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MINVOTINGPERIOD() (*big.Int, error) {
	return _NounsDAO.Contract.MINVOTINGPERIOD(&_NounsDAO.CallOpts)
}

// MINVOTINGPERIOD is a free data retrieval call binding the contract method 0x215809ca.
//
// Solidity: function MIN_VOTING_PERIOD() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MINVOTINGPERIOD() (*big.Int, error) {
	return _NounsDAO.Contract.MINVOTINGPERIOD(&_NounsDAO.CallOpts)
}

// REFUNDBASEGAS is a free data retrieval call binding the contract method 0x3be8ef3f.
//
// Solidity: function REFUND_BASE_GAS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) REFUNDBASEGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "REFUND_BASE_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REFUNDBASEGAS is a free data retrieval call binding the contract method 0x3be8ef3f.
//
// Solidity: function REFUND_BASE_GAS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) REFUNDBASEGAS() (*big.Int, error) {
	return _NounsDAO.Contract.REFUNDBASEGAS(&_NounsDAO.CallOpts)
}

// REFUNDBASEGAS is a free data retrieval call binding the contract method 0x3be8ef3f.
//
// Solidity: function REFUND_BASE_GAS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) REFUNDBASEGAS() (*big.Int, error) {
	return _NounsDAO.Contract.REFUNDBASEGAS(&_NounsDAO.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NounsDAO *NounsDAOCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NounsDAO *NounsDAOSession) Admin() (common.Address, error) {
	return _NounsDAO.Contract.Admin(&_NounsDAO.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) Admin() (common.Address, error) {
	return _NounsDAO.Contract.Admin(&_NounsDAO.CallOpts)
}

// DynamicQuorumVotes is a free data retrieval call binding the contract method 0x9a0dfb53.
//
// Solidity: function dynamicQuorumVotes(uint256 againstVotes, uint256 totalSupply, (uint16,uint16,uint32) params) pure returns(uint256)
func (_NounsDAO *NounsDAOCaller) DynamicQuorumVotes(opts *bind.CallOpts, againstVotes *big.Int, totalSupply *big.Int, params NounsDAOStorageV2DynamicQuorumParams) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "dynamicQuorumVotes", againstVotes, totalSupply, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DynamicQuorumVotes is a free data retrieval call binding the contract method 0x9a0dfb53.
//
// Solidity: function dynamicQuorumVotes(uint256 againstVotes, uint256 totalSupply, (uint16,uint16,uint32) params) pure returns(uint256)
func (_NounsDAO *NounsDAOSession) DynamicQuorumVotes(againstVotes *big.Int, totalSupply *big.Int, params NounsDAOStorageV2DynamicQuorumParams) (*big.Int, error) {
	return _NounsDAO.Contract.DynamicQuorumVotes(&_NounsDAO.CallOpts, againstVotes, totalSupply, params)
}

// DynamicQuorumVotes is a free data retrieval call binding the contract method 0x9a0dfb53.
//
// Solidity: function dynamicQuorumVotes(uint256 againstVotes, uint256 totalSupply, (uint16,uint16,uint32) params) pure returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) DynamicQuorumVotes(againstVotes *big.Int, totalSupply *big.Int, params NounsDAOStorageV2DynamicQuorumParams) (*big.Int, error) {
	return _NounsDAO.Contract.DynamicQuorumVotes(&_NounsDAO.CallOpts, againstVotes, totalSupply, params)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_NounsDAO *NounsDAOCaller) GetActions(opts *bind.CallOpts, proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "getActions", proposalId)

	outstruct := new(struct {
		Targets    []common.Address
		Values     []*big.Int
		Signatures []string
		Calldatas  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Targets = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Signatures = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Calldatas = *abi.ConvertType(out[3], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_NounsDAO *NounsDAOSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _NounsDAO.Contract.GetActions(&_NounsDAO.CallOpts, proposalId)
}

// GetActions is a free data retrieval call binding the contract method 0x328dd982.
//
// Solidity: function getActions(uint256 proposalId) view returns(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas)
func (_NounsDAO *NounsDAOCallerSession) GetActions(proposalId *big.Int) (struct {
	Targets    []common.Address
	Values     []*big.Int
	Signatures []string
	Calldatas  [][]byte
}, error) {
	return _NounsDAO.Contract.GetActions(&_NounsDAO.CallOpts, proposalId)
}

// GetDynamicQuorumParamsAt is a free data retrieval call binding the contract method 0x47f4a077.
//
// Solidity: function getDynamicQuorumParamsAt(uint256 blockNumber_) view returns((uint16,uint16,uint32))
func (_NounsDAO *NounsDAOCaller) GetDynamicQuorumParamsAt(opts *bind.CallOpts, blockNumber_ *big.Int) (NounsDAOStorageV2DynamicQuorumParams, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "getDynamicQuorumParamsAt", blockNumber_)

	if err != nil {
		return *new(NounsDAOStorageV2DynamicQuorumParams), err
	}

	out0 := *abi.ConvertType(out[0], new(NounsDAOStorageV2DynamicQuorumParams)).(*NounsDAOStorageV2DynamicQuorumParams)

	return out0, err

}

// GetDynamicQuorumParamsAt is a free data retrieval call binding the contract method 0x47f4a077.
//
// Solidity: function getDynamicQuorumParamsAt(uint256 blockNumber_) view returns((uint16,uint16,uint32))
func (_NounsDAO *NounsDAOSession) GetDynamicQuorumParamsAt(blockNumber_ *big.Int) (NounsDAOStorageV2DynamicQuorumParams, error) {
	return _NounsDAO.Contract.GetDynamicQuorumParamsAt(&_NounsDAO.CallOpts, blockNumber_)
}

// GetDynamicQuorumParamsAt is a free data retrieval call binding the contract method 0x47f4a077.
//
// Solidity: function getDynamicQuorumParamsAt(uint256 blockNumber_) view returns((uint16,uint16,uint32))
func (_NounsDAO *NounsDAOCallerSession) GetDynamicQuorumParamsAt(blockNumber_ *big.Int) (NounsDAOStorageV2DynamicQuorumParams, error) {
	return _NounsDAO.Contract.GetDynamicQuorumParamsAt(&_NounsDAO.CallOpts, blockNumber_)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_NounsDAO *NounsDAOCaller) GetReceipt(opts *bind.CallOpts, proposalId *big.Int, voter common.Address) (NounsDAOStorageV1AdjustedReceipt, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "getReceipt", proposalId, voter)

	if err != nil {
		return *new(NounsDAOStorageV1AdjustedReceipt), err
	}

	out0 := *abi.ConvertType(out[0], new(NounsDAOStorageV1AdjustedReceipt)).(*NounsDAOStorageV1AdjustedReceipt)

	return out0, err

}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_NounsDAO *NounsDAOSession) GetReceipt(proposalId *big.Int, voter common.Address) (NounsDAOStorageV1AdjustedReceipt, error) {
	return _NounsDAO.Contract.GetReceipt(&_NounsDAO.CallOpts, proposalId, voter)
}

// GetReceipt is a free data retrieval call binding the contract method 0xe23a9a52.
//
// Solidity: function getReceipt(uint256 proposalId, address voter) view returns((bool,uint8,uint96))
func (_NounsDAO *NounsDAOCallerSession) GetReceipt(proposalId *big.Int, voter common.Address) (NounsDAOStorageV1AdjustedReceipt, error) {
	return _NounsDAO.Contract.GetReceipt(&_NounsDAO.CallOpts, proposalId, voter)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NounsDAO *NounsDAOCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NounsDAO *NounsDAOSession) Implementation() (common.Address, error) {
	return _NounsDAO.Contract.Implementation(&_NounsDAO.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) Implementation() (common.Address, error) {
	return _NounsDAO.Contract.Implementation(&_NounsDAO.CallOpts)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_NounsDAO *NounsDAOCaller) LatestProposalIds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "latestProposalIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_NounsDAO *NounsDAOSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _NounsDAO.Contract.LatestProposalIds(&_NounsDAO.CallOpts, arg0)
}

// LatestProposalIds is a free data retrieval call binding the contract method 0x17977c61.
//
// Solidity: function latestProposalIds(address ) view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) LatestProposalIds(arg0 common.Address) (*big.Int, error) {
	return _NounsDAO.Contract.LatestProposalIds(&_NounsDAO.CallOpts, arg0)
}

// MaxQuorumVotes is a free data retrieval call binding the contract method 0x7fa230bd.
//
// Solidity: function maxQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MaxQuorumVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "maxQuorumVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxQuorumVotes is a free data retrieval call binding the contract method 0x7fa230bd.
//
// Solidity: function maxQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MaxQuorumVotes() (*big.Int, error) {
	return _NounsDAO.Contract.MaxQuorumVotes(&_NounsDAO.CallOpts)
}

// MaxQuorumVotes is a free data retrieval call binding the contract method 0x7fa230bd.
//
// Solidity: function maxQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MaxQuorumVotes() (*big.Int, error) {
	return _NounsDAO.Contract.MaxQuorumVotes(&_NounsDAO.CallOpts)
}

// MinQuorumVotes is a free data retrieval call binding the contract method 0x3e273334.
//
// Solidity: function minQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) MinQuorumVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "minQuorumVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumVotes is a free data retrieval call binding the contract method 0x3e273334.
//
// Solidity: function minQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOSession) MinQuorumVotes() (*big.Int, error) {
	return _NounsDAO.Contract.MinQuorumVotes(&_NounsDAO.CallOpts)
}

// MinQuorumVotes is a free data retrieval call binding the contract method 0x3e273334.
//
// Solidity: function minQuorumVotes() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) MinQuorumVotes() (*big.Int, error) {
	return _NounsDAO.Contract.MinQuorumVotes(&_NounsDAO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsDAO *NounsDAOCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsDAO *NounsDAOSession) Name() (string, error) {
	return _NounsDAO.Contract.Name(&_NounsDAO.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NounsDAO *NounsDAOCallerSession) Name() (string, error) {
	return _NounsDAO.Contract.Name(&_NounsDAO.CallOpts)
}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsDAO *NounsDAOCaller) Nouns(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "nouns")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsDAO *NounsDAOSession) Nouns() (common.Address, error) {
	return _NounsDAO.Contract.Nouns(&_NounsDAO.CallOpts)
}

// Nouns is a free data retrieval call binding the contract method 0x2de45f18.
//
// Solidity: function nouns() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) Nouns() (common.Address, error) {
	return _NounsDAO.Contract.Nouns(&_NounsDAO.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_NounsDAO *NounsDAOCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_NounsDAO *NounsDAOSession) PendingAdmin() (common.Address, error) {
	return _NounsDAO.Contract.PendingAdmin(&_NounsDAO.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) PendingAdmin() (common.Address, error) {
	return _NounsDAO.Contract.PendingAdmin(&_NounsDAO.CallOpts)
}

// PendingVetoer is a free data retrieval call binding the contract method 0xa67d0635.
//
// Solidity: function pendingVetoer() view returns(address)
func (_NounsDAO *NounsDAOCaller) PendingVetoer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "pendingVetoer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingVetoer is a free data retrieval call binding the contract method 0xa67d0635.
//
// Solidity: function pendingVetoer() view returns(address)
func (_NounsDAO *NounsDAOSession) PendingVetoer() (common.Address, error) {
	return _NounsDAO.Contract.PendingVetoer(&_NounsDAO.CallOpts)
}

// PendingVetoer is a free data retrieval call binding the contract method 0xa67d0635.
//
// Solidity: function pendingVetoer() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) PendingVetoer() (common.Address, error) {
	return _NounsDAO.Contract.PendingVetoer(&_NounsDAO.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_NounsDAO *NounsDAOSession) ProposalCount() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalCount(&_NounsDAO.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) ProposalCount() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalCount(&_NounsDAO.CallOpts)
}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) ProposalMaxOperations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "proposalMaxOperations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (_NounsDAO *NounsDAOSession) ProposalMaxOperations() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalMaxOperations(&_NounsDAO.CallOpts)
}

// ProposalMaxOperations is a free data retrieval call binding the contract method 0x7bdbe4d0.
//
// Solidity: function proposalMaxOperations() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) ProposalMaxOperations() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalMaxOperations(&_NounsDAO.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_NounsDAO *NounsDAOSession) ProposalThreshold() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalThreshold(&_NounsDAO.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) ProposalThreshold() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalThreshold(&_NounsDAO.CallOpts)
}

// ProposalThresholdBPS is a free data retrieval call binding the contract method 0x14a67ea4.
//
// Solidity: function proposalThresholdBPS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) ProposalThresholdBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "proposalThresholdBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThresholdBPS is a free data retrieval call binding the contract method 0x14a67ea4.
//
// Solidity: function proposalThresholdBPS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) ProposalThresholdBPS() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalThresholdBPS(&_NounsDAO.CallOpts)
}

// ProposalThresholdBPS is a free data retrieval call binding the contract method 0x14a67ea4.
//
// Solidity: function proposalThresholdBPS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) ProposalThresholdBPS() (*big.Int, error) {
	return _NounsDAO.Contract.ProposalThresholdBPS(&_NounsDAO.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,uint256,uint256))
func (_NounsDAO *NounsDAOCaller) Proposals(opts *bind.CallOpts, proposalId *big.Int) (NounsDAOStorageV2ProposalCondensed, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "proposals", proposalId)

	if err != nil {
		return *new(NounsDAOStorageV2ProposalCondensed), err
	}

	out0 := *abi.ConvertType(out[0], new(NounsDAOStorageV2ProposalCondensed)).(*NounsDAOStorageV2ProposalCondensed)

	return out0, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,uint256,uint256))
func (_NounsDAO *NounsDAOSession) Proposals(proposalId *big.Int) (NounsDAOStorageV2ProposalCondensed, error) {
	return _NounsDAO.Contract.Proposals(&_NounsDAO.CallOpts, proposalId)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 proposalId) view returns((uint256,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool,uint256,uint256))
func (_NounsDAO *NounsDAOCallerSession) Proposals(proposalId *big.Int) (NounsDAOStorageV2ProposalCondensed, error) {
	return _NounsDAO.Contract.Proposals(&_NounsDAO.CallOpts, proposalId)
}

// QuorumParamsCheckpoints is a free data retrieval call binding the contract method 0xabb308b2.
//
// Solidity: function quorumParamsCheckpoints(uint256 ) view returns(uint32 fromBlock, (uint16,uint16,uint32) params)
func (_NounsDAO *NounsDAOCaller) QuorumParamsCheckpoints(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FromBlock uint32
	Params    NounsDAOStorageV2DynamicQuorumParams
}, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "quorumParamsCheckpoints", arg0)

	outstruct := new(struct {
		FromBlock uint32
		Params    NounsDAOStorageV2DynamicQuorumParams
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Params = *abi.ConvertType(out[1], new(NounsDAOStorageV2DynamicQuorumParams)).(*NounsDAOStorageV2DynamicQuorumParams)

	return *outstruct, err

}

// QuorumParamsCheckpoints is a free data retrieval call binding the contract method 0xabb308b2.
//
// Solidity: function quorumParamsCheckpoints(uint256 ) view returns(uint32 fromBlock, (uint16,uint16,uint32) params)
func (_NounsDAO *NounsDAOSession) QuorumParamsCheckpoints(arg0 *big.Int) (struct {
	FromBlock uint32
	Params    NounsDAOStorageV2DynamicQuorumParams
}, error) {
	return _NounsDAO.Contract.QuorumParamsCheckpoints(&_NounsDAO.CallOpts, arg0)
}

// QuorumParamsCheckpoints is a free data retrieval call binding the contract method 0xabb308b2.
//
// Solidity: function quorumParamsCheckpoints(uint256 ) view returns(uint32 fromBlock, (uint16,uint16,uint32) params)
func (_NounsDAO *NounsDAOCallerSession) QuorumParamsCheckpoints(arg0 *big.Int) (struct {
	FromBlock uint32
	Params    NounsDAOStorageV2DynamicQuorumParams
}, error) {
	return _NounsDAO.Contract.QuorumParamsCheckpoints(&_NounsDAO.CallOpts, arg0)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x0f7b1f08.
//
// Solidity: function quorumVotes(uint256 proposalId) view returns(uint256)
func (_NounsDAO *NounsDAOCaller) QuorumVotes(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "quorumVotes", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumVotes is a free data retrieval call binding the contract method 0x0f7b1f08.
//
// Solidity: function quorumVotes(uint256 proposalId) view returns(uint256)
func (_NounsDAO *NounsDAOSession) QuorumVotes(proposalId *big.Int) (*big.Int, error) {
	return _NounsDAO.Contract.QuorumVotes(&_NounsDAO.CallOpts, proposalId)
}

// QuorumVotes is a free data retrieval call binding the contract method 0x0f7b1f08.
//
// Solidity: function quorumVotes(uint256 proposalId) view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) QuorumVotes(proposalId *big.Int) (*big.Int, error) {
	return _NounsDAO.Contract.QuorumVotes(&_NounsDAO.CallOpts, proposalId)
}

// QuorumVotesBPS is a free data retrieval call binding the contract method 0x83cce0e1.
//
// Solidity: function quorumVotesBPS() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) QuorumVotesBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "quorumVotesBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumVotesBPS is a free data retrieval call binding the contract method 0x83cce0e1.
//
// Solidity: function quorumVotesBPS() view returns(uint256)
func (_NounsDAO *NounsDAOSession) QuorumVotesBPS() (*big.Int, error) {
	return _NounsDAO.Contract.QuorumVotesBPS(&_NounsDAO.CallOpts)
}

// QuorumVotesBPS is a free data retrieval call binding the contract method 0x83cce0e1.
//
// Solidity: function quorumVotesBPS() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) QuorumVotesBPS() (*big.Int, error) {
	return _NounsDAO.Contract.QuorumVotesBPS(&_NounsDAO.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_NounsDAO *NounsDAOCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_NounsDAO *NounsDAOSession) State(proposalId *big.Int) (uint8, error) {
	return _NounsDAO.Contract.State(&_NounsDAO.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_NounsDAO *NounsDAOCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _NounsDAO.Contract.State(&_NounsDAO.CallOpts, proposalId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_NounsDAO *NounsDAOCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_NounsDAO *NounsDAOSession) Timelock() (common.Address, error) {
	return _NounsDAO.Contract.Timelock(&_NounsDAO.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) Timelock() (common.Address, error) {
	return _NounsDAO.Contract.Timelock(&_NounsDAO.CallOpts)
}

// Vetoer is a free data retrieval call binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() view returns(address)
func (_NounsDAO *NounsDAOCaller) Vetoer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "vetoer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vetoer is a free data retrieval call binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() view returns(address)
func (_NounsDAO *NounsDAOSession) Vetoer() (common.Address, error) {
	return _NounsDAO.Contract.Vetoer(&_NounsDAO.CallOpts)
}

// Vetoer is a free data retrieval call binding the contract method 0xd8bff440.
//
// Solidity: function vetoer() view returns(address)
func (_NounsDAO *NounsDAOCallerSession) Vetoer() (common.Address, error) {
	return _NounsDAO.Contract.Vetoer(&_NounsDAO.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_NounsDAO *NounsDAOSession) VotingDelay() (*big.Int, error) {
	return _NounsDAO.Contract.VotingDelay(&_NounsDAO.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) VotingDelay() (*big.Int, error) {
	return _NounsDAO.Contract.VotingDelay(&_NounsDAO.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_NounsDAO *NounsDAOCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NounsDAO.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_NounsDAO *NounsDAOSession) VotingPeriod() (*big.Int, error) {
	return _NounsDAO.Contract.VotingPeriod(&_NounsDAO.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_NounsDAO *NounsDAOCallerSession) VotingPeriod() (*big.Int, error) {
	return _NounsDAO.Contract.VotingPeriod(&_NounsDAO.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_NounsDAO *NounsDAOTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_NounsDAO *NounsDAOSession) AcceptAdmin() (*types.Transaction, error) {
	return _NounsDAO.Contract.AcceptAdmin(&_NounsDAO.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_NounsDAO *NounsDAOTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _NounsDAO.Contract.AcceptAdmin(&_NounsDAO.TransactOpts)
}

// AcceptVetoer is a paid mutator transaction binding the contract method 0x2cfc81c6.
//
// Solidity: function _acceptVetoer() returns()
func (_NounsDAO *NounsDAOTransactor) AcceptVetoer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_acceptVetoer")
}

// AcceptVetoer is a paid mutator transaction binding the contract method 0x2cfc81c6.
//
// Solidity: function _acceptVetoer() returns()
func (_NounsDAO *NounsDAOSession) AcceptVetoer() (*types.Transaction, error) {
	return _NounsDAO.Contract.AcceptVetoer(&_NounsDAO.TransactOpts)
}

// AcceptVetoer is a paid mutator transaction binding the contract method 0x2cfc81c6.
//
// Solidity: function _acceptVetoer() returns()
func (_NounsDAO *NounsDAOTransactorSession) AcceptVetoer() (*types.Transaction, error) {
	return _NounsDAO.Contract.AcceptVetoer(&_NounsDAO.TransactOpts)
}

// BurnVetoPower is a paid mutator transaction binding the contract method 0xbf7a2963.
//
// Solidity: function _burnVetoPower() returns()
func (_NounsDAO *NounsDAOTransactor) BurnVetoPower(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_burnVetoPower")
}

// BurnVetoPower is a paid mutator transaction binding the contract method 0xbf7a2963.
//
// Solidity: function _burnVetoPower() returns()
func (_NounsDAO *NounsDAOSession) BurnVetoPower() (*types.Transaction, error) {
	return _NounsDAO.Contract.BurnVetoPower(&_NounsDAO.TransactOpts)
}

// BurnVetoPower is a paid mutator transaction binding the contract method 0xbf7a2963.
//
// Solidity: function _burnVetoPower() returns()
func (_NounsDAO *NounsDAOTransactorSession) BurnVetoPower() (*types.Transaction, error) {
	return _NounsDAO.Contract.BurnVetoPower(&_NounsDAO.TransactOpts)
}

// SetDynamicQuorumParams is a paid mutator transaction binding the contract method 0xec91deda.
//
// Solidity: function _setDynamicQuorumParams(uint16 newMinQuorumVotesBPS, uint16 newMaxQuorumVotesBPS, uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOTransactor) SetDynamicQuorumParams(opts *bind.TransactOpts, newMinQuorumVotesBPS uint16, newMaxQuorumVotesBPS uint16, newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setDynamicQuorumParams", newMinQuorumVotesBPS, newMaxQuorumVotesBPS, newQuorumCoefficient)
}

// SetDynamicQuorumParams is a paid mutator transaction binding the contract method 0xec91deda.
//
// Solidity: function _setDynamicQuorumParams(uint16 newMinQuorumVotesBPS, uint16 newMaxQuorumVotesBPS, uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOSession) SetDynamicQuorumParams(newMinQuorumVotesBPS uint16, newMaxQuorumVotesBPS uint16, newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetDynamicQuorumParams(&_NounsDAO.TransactOpts, newMinQuorumVotesBPS, newMaxQuorumVotesBPS, newQuorumCoefficient)
}

// SetDynamicQuorumParams is a paid mutator transaction binding the contract method 0xec91deda.
//
// Solidity: function _setDynamicQuorumParams(uint16 newMinQuorumVotesBPS, uint16 newMaxQuorumVotesBPS, uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetDynamicQuorumParams(newMinQuorumVotesBPS uint16, newMaxQuorumVotesBPS uint16, newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetDynamicQuorumParams(&_NounsDAO.TransactOpts, newMinQuorumVotesBPS, newMaxQuorumVotesBPS, newQuorumCoefficient)
}

// SetMaxQuorumVotesBPS is a paid mutator transaction binding the contract method 0x50196db3.
//
// Solidity: function _setMaxQuorumVotesBPS(uint16 newMaxQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOTransactor) SetMaxQuorumVotesBPS(opts *bind.TransactOpts, newMaxQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setMaxQuorumVotesBPS", newMaxQuorumVotesBPS)
}

// SetMaxQuorumVotesBPS is a paid mutator transaction binding the contract method 0x50196db3.
//
// Solidity: function _setMaxQuorumVotesBPS(uint16 newMaxQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOSession) SetMaxQuorumVotesBPS(newMaxQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetMaxQuorumVotesBPS(&_NounsDAO.TransactOpts, newMaxQuorumVotesBPS)
}

// SetMaxQuorumVotesBPS is a paid mutator transaction binding the contract method 0x50196db3.
//
// Solidity: function _setMaxQuorumVotesBPS(uint16 newMaxQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetMaxQuorumVotesBPS(newMaxQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetMaxQuorumVotesBPS(&_NounsDAO.TransactOpts, newMaxQuorumVotesBPS)
}

// SetMinQuorumVotesBPS is a paid mutator transaction binding the contract method 0x7a3da691.
//
// Solidity: function _setMinQuorumVotesBPS(uint16 newMinQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOTransactor) SetMinQuorumVotesBPS(opts *bind.TransactOpts, newMinQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setMinQuorumVotesBPS", newMinQuorumVotesBPS)
}

// SetMinQuorumVotesBPS is a paid mutator transaction binding the contract method 0x7a3da691.
//
// Solidity: function _setMinQuorumVotesBPS(uint16 newMinQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOSession) SetMinQuorumVotesBPS(newMinQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetMinQuorumVotesBPS(&_NounsDAO.TransactOpts, newMinQuorumVotesBPS)
}

// SetMinQuorumVotesBPS is a paid mutator transaction binding the contract method 0x7a3da691.
//
// Solidity: function _setMinQuorumVotesBPS(uint16 newMinQuorumVotesBPS) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetMinQuorumVotesBPS(newMinQuorumVotesBPS uint16) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetMinQuorumVotesBPS(&_NounsDAO.TransactOpts, newMinQuorumVotesBPS)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_NounsDAO *NounsDAOTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_NounsDAO *NounsDAOSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetPendingAdmin(&_NounsDAO.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetPendingAdmin(&_NounsDAO.TransactOpts, newPendingAdmin)
}

// SetPendingVetoer is a paid mutator transaction binding the contract method 0xd3f662e1.
//
// Solidity: function _setPendingVetoer(address newPendingVetoer) returns()
func (_NounsDAO *NounsDAOTransactor) SetPendingVetoer(opts *bind.TransactOpts, newPendingVetoer common.Address) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setPendingVetoer", newPendingVetoer)
}

// SetPendingVetoer is a paid mutator transaction binding the contract method 0xd3f662e1.
//
// Solidity: function _setPendingVetoer(address newPendingVetoer) returns()
func (_NounsDAO *NounsDAOSession) SetPendingVetoer(newPendingVetoer common.Address) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetPendingVetoer(&_NounsDAO.TransactOpts, newPendingVetoer)
}

// SetPendingVetoer is a paid mutator transaction binding the contract method 0xd3f662e1.
//
// Solidity: function _setPendingVetoer(address newPendingVetoer) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetPendingVetoer(newPendingVetoer common.Address) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetPendingVetoer(&_NounsDAO.TransactOpts, newPendingVetoer)
}

// SetProposalThresholdBPS is a paid mutator transaction binding the contract method 0x97d048e5.
//
// Solidity: function _setProposalThresholdBPS(uint256 newProposalThresholdBPS) returns()
func (_NounsDAO *NounsDAOTransactor) SetProposalThresholdBPS(opts *bind.TransactOpts, newProposalThresholdBPS *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setProposalThresholdBPS", newProposalThresholdBPS)
}

// SetProposalThresholdBPS is a paid mutator transaction binding the contract method 0x97d048e5.
//
// Solidity: function _setProposalThresholdBPS(uint256 newProposalThresholdBPS) returns()
func (_NounsDAO *NounsDAOSession) SetProposalThresholdBPS(newProposalThresholdBPS *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetProposalThresholdBPS(&_NounsDAO.TransactOpts, newProposalThresholdBPS)
}

// SetProposalThresholdBPS is a paid mutator transaction binding the contract method 0x97d048e5.
//
// Solidity: function _setProposalThresholdBPS(uint256 newProposalThresholdBPS) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetProposalThresholdBPS(newProposalThresholdBPS *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetProposalThresholdBPS(&_NounsDAO.TransactOpts, newProposalThresholdBPS)
}

// SetQuorumCoefficient is a paid mutator transaction binding the contract method 0x2b5ca189.
//
// Solidity: function _setQuorumCoefficient(uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOTransactor) SetQuorumCoefficient(opts *bind.TransactOpts, newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setQuorumCoefficient", newQuorumCoefficient)
}

// SetQuorumCoefficient is a paid mutator transaction binding the contract method 0x2b5ca189.
//
// Solidity: function _setQuorumCoefficient(uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOSession) SetQuorumCoefficient(newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetQuorumCoefficient(&_NounsDAO.TransactOpts, newQuorumCoefficient)
}

// SetQuorumCoefficient is a paid mutator transaction binding the contract method 0x2b5ca189.
//
// Solidity: function _setQuorumCoefficient(uint32 newQuorumCoefficient) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetQuorumCoefficient(newQuorumCoefficient uint32) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetQuorumCoefficient(&_NounsDAO.TransactOpts, newQuorumCoefficient)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x1dfb1b5a.
//
// Solidity: function _setVotingDelay(uint256 newVotingDelay) returns()
func (_NounsDAO *NounsDAOTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x1dfb1b5a.
//
// Solidity: function _setVotingDelay(uint256 newVotingDelay) returns()
func (_NounsDAO *NounsDAOSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetVotingDelay(&_NounsDAO.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x1dfb1b5a.
//
// Solidity: function _setVotingDelay(uint256 newVotingDelay) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetVotingDelay(&_NounsDAO.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0x0ea2d98c.
//
// Solidity: function _setVotingPeriod(uint256 newVotingPeriod) returns()
func (_NounsDAO *NounsDAOTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0x0ea2d98c.
//
// Solidity: function _setVotingPeriod(uint256 newVotingPeriod) returns()
func (_NounsDAO *NounsDAOSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetVotingPeriod(&_NounsDAO.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0x0ea2d98c.
//
// Solidity: function _setVotingPeriod(uint256 newVotingPeriod) returns()
func (_NounsDAO *NounsDAOTransactorSession) SetVotingPeriod(newVotingPeriod *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.SetVotingPeriod(&_NounsDAO.TransactOpts, newVotingPeriod)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc10eb14d.
//
// Solidity: function _withdraw() returns(uint256, bool)
func (_NounsDAO *NounsDAOTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "_withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0xc10eb14d.
//
// Solidity: function _withdraw() returns(uint256, bool)
func (_NounsDAO *NounsDAOSession) Withdraw() (*types.Transaction, error) {
	return _NounsDAO.Contract.Withdraw(&_NounsDAO.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xc10eb14d.
//
// Solidity: function _withdraw() returns(uint256, bool)
func (_NounsDAO *NounsDAOTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NounsDAO.Contract.Withdraw(&_NounsDAO.TransactOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactor) Cancel(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "cancel", proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Cancel(&_NounsDAO.TransactOpts, proposalId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactorSession) Cancel(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Cancel(&_NounsDAO.TransactOpts, proposalId)
}

// CastRefundableVote is a paid mutator transaction binding the contract method 0x44fac8f6.
//
// Solidity: function castRefundableVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOTransactor) CastRefundableVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "castRefundableVote", proposalId, support)
}

// CastRefundableVote is a paid mutator transaction binding the contract method 0x44fac8f6.
//
// Solidity: function castRefundableVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOSession) CastRefundableVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastRefundableVote(&_NounsDAO.TransactOpts, proposalId, support)
}

// CastRefundableVote is a paid mutator transaction binding the contract method 0x44fac8f6.
//
// Solidity: function castRefundableVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOTransactorSession) CastRefundableVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastRefundableVote(&_NounsDAO.TransactOpts, proposalId, support)
}

// CastRefundableVoteWithReason is a paid mutator transaction binding the contract method 0x64c05995.
//
// Solidity: function castRefundableVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOTransactor) CastRefundableVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "castRefundableVoteWithReason", proposalId, support, reason)
}

// CastRefundableVoteWithReason is a paid mutator transaction binding the contract method 0x64c05995.
//
// Solidity: function castRefundableVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOSession) CastRefundableVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastRefundableVoteWithReason(&_NounsDAO.TransactOpts, proposalId, support, reason)
}

// CastRefundableVoteWithReason is a paid mutator transaction binding the contract method 0x64c05995.
//
// Solidity: function castRefundableVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOTransactorSession) CastRefundableVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastRefundableVoteWithReason(&_NounsDAO.TransactOpts, proposalId, support, reason)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVote(&_NounsDAO.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns()
func (_NounsDAO *NounsDAOTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVote(&_NounsDAO.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsDAO *NounsDAOTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "castVoteBySig", proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsDAO *NounsDAOSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVoteBySig(&_NounsDAO.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x3bccf4fd.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, uint8 v, bytes32 r, bytes32 s) returns()
func (_NounsDAO *NounsDAOTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVoteBySig(&_NounsDAO.TransactOpts, proposalId, support, v, r, s)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVoteWithReason(&_NounsDAO.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns()
func (_NounsDAO *NounsDAOTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _NounsDAO.Contract.CastVoteWithReason(&_NounsDAO.TransactOpts, proposalId, support, reason)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Execute(&_NounsDAO.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Execute(&_NounsDAO.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d183f66.
//
// Solidity: function initialize(address timelock_, address nouns_, address vetoer_, uint256 votingPeriod_, uint256 votingDelay_, uint256 proposalThresholdBPS_, (uint16,uint16,uint32) dynamicQuorumParams_) returns()
func (_NounsDAO *NounsDAOTransactor) Initialize(opts *bind.TransactOpts, timelock_ common.Address, nouns_ common.Address, vetoer_ common.Address, votingPeriod_ *big.Int, votingDelay_ *big.Int, proposalThresholdBPS_ *big.Int, dynamicQuorumParams_ NounsDAOStorageV2DynamicQuorumParams) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "initialize", timelock_, nouns_, vetoer_, votingPeriod_, votingDelay_, proposalThresholdBPS_, dynamicQuorumParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d183f66.
//
// Solidity: function initialize(address timelock_, address nouns_, address vetoer_, uint256 votingPeriod_, uint256 votingDelay_, uint256 proposalThresholdBPS_, (uint16,uint16,uint32) dynamicQuorumParams_) returns()
func (_NounsDAO *NounsDAOSession) Initialize(timelock_ common.Address, nouns_ common.Address, vetoer_ common.Address, votingPeriod_ *big.Int, votingDelay_ *big.Int, proposalThresholdBPS_ *big.Int, dynamicQuorumParams_ NounsDAOStorageV2DynamicQuorumParams) (*types.Transaction, error) {
	return _NounsDAO.Contract.Initialize(&_NounsDAO.TransactOpts, timelock_, nouns_, vetoer_, votingPeriod_, votingDelay_, proposalThresholdBPS_, dynamicQuorumParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x0d183f66.
//
// Solidity: function initialize(address timelock_, address nouns_, address vetoer_, uint256 votingPeriod_, uint256 votingDelay_, uint256 proposalThresholdBPS_, (uint16,uint16,uint32) dynamicQuorumParams_) returns()
func (_NounsDAO *NounsDAOTransactorSession) Initialize(timelock_ common.Address, nouns_ common.Address, vetoer_ common.Address, votingPeriod_ *big.Int, votingDelay_ *big.Int, proposalThresholdBPS_ *big.Int, dynamicQuorumParams_ NounsDAOStorageV2DynamicQuorumParams) (*types.Transaction, error) {
	return _NounsDAO.Contract.Initialize(&_NounsDAO.TransactOpts, timelock_, nouns_, vetoer_, votingPeriod_, votingDelay_, proposalThresholdBPS_, dynamicQuorumParams_)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_NounsDAO *NounsDAOTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "propose", targets, values, signatures, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_NounsDAO *NounsDAOSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _NounsDAO.Contract.Propose(&_NounsDAO.TransactOpts, targets, values, signatures, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0xda95691a.
//
// Solidity: function propose(address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, string description) returns(uint256)
func (_NounsDAO *NounsDAOTransactorSession) Propose(targets []common.Address, values []*big.Int, signatures []string, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _NounsDAO.Contract.Propose(&_NounsDAO.TransactOpts, targets, values, signatures, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactor) Queue(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "queue", proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Queue(&_NounsDAO.TransactOpts, proposalId)
}

// Queue is a paid mutator transaction binding the contract method 0xddf0b009.
//
// Solidity: function queue(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactorSession) Queue(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Queue(&_NounsDAO.TransactOpts, proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactor) Veto(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.contract.Transact(opts, "veto", proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOSession) Veto(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Veto(&_NounsDAO.TransactOpts, proposalId)
}

// Veto is a paid mutator transaction binding the contract method 0x1d28dec7.
//
// Solidity: function veto(uint256 proposalId) returns()
func (_NounsDAO *NounsDAOTransactorSession) Veto(proposalId *big.Int) (*types.Transaction, error) {
	return _NounsDAO.Contract.Veto(&_NounsDAO.TransactOpts, proposalId)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NounsDAO *NounsDAOTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NounsDAO.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NounsDAO *NounsDAOSession) Receive() (*types.Transaction, error) {
	return _NounsDAO.Contract.Receive(&_NounsDAO.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_NounsDAO *NounsDAOTransactorSession) Receive() (*types.Transaction, error) {
	return _NounsDAO.Contract.Receive(&_NounsDAO.TransactOpts)
}

// NounsDAOMaxQuorumVotesBPSSetIterator is returned from FilterMaxQuorumVotesBPSSet and is used to iterate over the raw logs and unpacked data for MaxQuorumVotesBPSSet events raised by the NounsDAO contract.
type NounsDAOMaxQuorumVotesBPSSetIterator struct {
	Event *NounsDAOMaxQuorumVotesBPSSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOMaxQuorumVotesBPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOMaxQuorumVotesBPSSet)
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
		it.Event = new(NounsDAOMaxQuorumVotesBPSSet)
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
func (it *NounsDAOMaxQuorumVotesBPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOMaxQuorumVotesBPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOMaxQuorumVotesBPSSet represents a MaxQuorumVotesBPSSet event raised by the NounsDAO contract.
type NounsDAOMaxQuorumVotesBPSSet struct {
	OldMaxQuorumVotesBPS uint16
	NewMaxQuorumVotesBPS uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMaxQuorumVotesBPSSet is a free log retrieval operation binding the contract event 0x4bfb1235074b38f02e5cf8ba90f535905417c196a12654f65ee0584512d70642.
//
// Solidity: event MaxQuorumVotesBPSSet(uint16 oldMaxQuorumVotesBPS, uint16 newMaxQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) FilterMaxQuorumVotesBPSSet(opts *bind.FilterOpts) (*NounsDAOMaxQuorumVotesBPSSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "MaxQuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOMaxQuorumVotesBPSSetIterator{contract: _NounsDAO.contract, event: "MaxQuorumVotesBPSSet", logs: logs, sub: sub}, nil
}

// WatchMaxQuorumVotesBPSSet is a free log subscription operation binding the contract event 0x4bfb1235074b38f02e5cf8ba90f535905417c196a12654f65ee0584512d70642.
//
// Solidity: event MaxQuorumVotesBPSSet(uint16 oldMaxQuorumVotesBPS, uint16 newMaxQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) WatchMaxQuorumVotesBPSSet(opts *bind.WatchOpts, sink chan<- *NounsDAOMaxQuorumVotesBPSSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "MaxQuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOMaxQuorumVotesBPSSet)
				if err := _NounsDAO.contract.UnpackLog(event, "MaxQuorumVotesBPSSet", log); err != nil {
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

// ParseMaxQuorumVotesBPSSet is a log parse operation binding the contract event 0x4bfb1235074b38f02e5cf8ba90f535905417c196a12654f65ee0584512d70642.
//
// Solidity: event MaxQuorumVotesBPSSet(uint16 oldMaxQuorumVotesBPS, uint16 newMaxQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) ParseMaxQuorumVotesBPSSet(log types.Log) (*NounsDAOMaxQuorumVotesBPSSet, error) {
	event := new(NounsDAOMaxQuorumVotesBPSSet)
	if err := _NounsDAO.contract.UnpackLog(event, "MaxQuorumVotesBPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOMinQuorumVotesBPSSetIterator is returned from FilterMinQuorumVotesBPSSet and is used to iterate over the raw logs and unpacked data for MinQuorumVotesBPSSet events raised by the NounsDAO contract.
type NounsDAOMinQuorumVotesBPSSetIterator struct {
	Event *NounsDAOMinQuorumVotesBPSSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOMinQuorumVotesBPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOMinQuorumVotesBPSSet)
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
		it.Event = new(NounsDAOMinQuorumVotesBPSSet)
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
func (it *NounsDAOMinQuorumVotesBPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOMinQuorumVotesBPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOMinQuorumVotesBPSSet represents a MinQuorumVotesBPSSet event raised by the NounsDAO contract.
type NounsDAOMinQuorumVotesBPSSet struct {
	OldMinQuorumVotesBPS uint16
	NewMinQuorumVotesBPS uint16
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMinQuorumVotesBPSSet is a free log retrieval operation binding the contract event 0xfaeebe30d875e399189096ea49fea81bd41fe6dfc86ad3550639063219e60778.
//
// Solidity: event MinQuorumVotesBPSSet(uint16 oldMinQuorumVotesBPS, uint16 newMinQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) FilterMinQuorumVotesBPSSet(opts *bind.FilterOpts) (*NounsDAOMinQuorumVotesBPSSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "MinQuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOMinQuorumVotesBPSSetIterator{contract: _NounsDAO.contract, event: "MinQuorumVotesBPSSet", logs: logs, sub: sub}, nil
}

// WatchMinQuorumVotesBPSSet is a free log subscription operation binding the contract event 0xfaeebe30d875e399189096ea49fea81bd41fe6dfc86ad3550639063219e60778.
//
// Solidity: event MinQuorumVotesBPSSet(uint16 oldMinQuorumVotesBPS, uint16 newMinQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) WatchMinQuorumVotesBPSSet(opts *bind.WatchOpts, sink chan<- *NounsDAOMinQuorumVotesBPSSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "MinQuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOMinQuorumVotesBPSSet)
				if err := _NounsDAO.contract.UnpackLog(event, "MinQuorumVotesBPSSet", log); err != nil {
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

// ParseMinQuorumVotesBPSSet is a log parse operation binding the contract event 0xfaeebe30d875e399189096ea49fea81bd41fe6dfc86ad3550639063219e60778.
//
// Solidity: event MinQuorumVotesBPSSet(uint16 oldMinQuorumVotesBPS, uint16 newMinQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) ParseMinQuorumVotesBPSSet(log types.Log) (*NounsDAOMinQuorumVotesBPSSet, error) {
	event := new(NounsDAOMinQuorumVotesBPSSet)
	if err := _NounsDAO.contract.UnpackLog(event, "MinQuorumVotesBPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAONewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the NounsDAO contract.
type NounsDAONewAdminIterator struct {
	Event *NounsDAONewAdmin // Event containing the contract specifics and raw log

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
func (it *NounsDAONewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAONewAdmin)
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
		it.Event = new(NounsDAONewAdmin)
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
func (it *NounsDAONewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAONewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAONewAdmin represents a NewAdmin event raised by the NounsDAO contract.
type NounsDAONewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_NounsDAO *NounsDAOFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*NounsDAONewAdminIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &NounsDAONewAdminIterator{contract: _NounsDAO.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_NounsDAO *NounsDAOFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *NounsDAONewAdmin) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAONewAdmin)
				if err := _NounsDAO.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_NounsDAO *NounsDAOFilterer) ParseNewAdmin(log types.Log) (*NounsDAONewAdmin, error) {
	event := new(NounsDAONewAdmin)
	if err := _NounsDAO.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAONewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the NounsDAO contract.
type NounsDAONewImplementationIterator struct {
	Event *NounsDAONewImplementation // Event containing the contract specifics and raw log

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
func (it *NounsDAONewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAONewImplementation)
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
		it.Event = new(NounsDAONewImplementation)
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
func (it *NounsDAONewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAONewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAONewImplementation represents a NewImplementation event raised by the NounsDAO contract.
type NounsDAONewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_NounsDAO *NounsDAOFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*NounsDAONewImplementationIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &NounsDAONewImplementationIterator{contract: _NounsDAO.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_NounsDAO *NounsDAOFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *NounsDAONewImplementation) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAONewImplementation)
				if err := _NounsDAO.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_NounsDAO *NounsDAOFilterer) ParseNewImplementation(log types.Log) (*NounsDAONewImplementation, error) {
	event := new(NounsDAONewImplementation)
	if err := _NounsDAO.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAONewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the NounsDAO contract.
type NounsDAONewPendingAdminIterator struct {
	Event *NounsDAONewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *NounsDAONewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAONewPendingAdmin)
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
		it.Event = new(NounsDAONewPendingAdmin)
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
func (it *NounsDAONewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAONewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAONewPendingAdmin represents a NewPendingAdmin event raised by the NounsDAO contract.
type NounsDAONewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_NounsDAO *NounsDAOFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*NounsDAONewPendingAdminIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &NounsDAONewPendingAdminIterator{contract: _NounsDAO.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_NounsDAO *NounsDAOFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *NounsDAONewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAONewPendingAdmin)
				if err := _NounsDAO.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_NounsDAO *NounsDAOFilterer) ParseNewPendingAdmin(log types.Log) (*NounsDAONewPendingAdmin, error) {
	event := new(NounsDAONewPendingAdmin)
	if err := _NounsDAO.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAONewPendingVetoerIterator is returned from FilterNewPendingVetoer and is used to iterate over the raw logs and unpacked data for NewPendingVetoer events raised by the NounsDAO contract.
type NounsDAONewPendingVetoerIterator struct {
	Event *NounsDAONewPendingVetoer // Event containing the contract specifics and raw log

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
func (it *NounsDAONewPendingVetoerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAONewPendingVetoer)
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
		it.Event = new(NounsDAONewPendingVetoer)
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
func (it *NounsDAONewPendingVetoerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAONewPendingVetoerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAONewPendingVetoer represents a NewPendingVetoer event raised by the NounsDAO contract.
type NounsDAONewPendingVetoer struct {
	OldPendingVetoer common.Address
	NewPendingVetoer common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPendingVetoer is a free log retrieval operation binding the contract event 0x7ad92e57a52c4e3e83ba624622b14e3a5efa0160dd6f9a7975c43ea66bad79ea.
//
// Solidity: event NewPendingVetoer(address oldPendingVetoer, address newPendingVetoer)
func (_NounsDAO *NounsDAOFilterer) FilterNewPendingVetoer(opts *bind.FilterOpts) (*NounsDAONewPendingVetoerIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "NewPendingVetoer")
	if err != nil {
		return nil, err
	}
	return &NounsDAONewPendingVetoerIterator{contract: _NounsDAO.contract, event: "NewPendingVetoer", logs: logs, sub: sub}, nil
}

// WatchNewPendingVetoer is a free log subscription operation binding the contract event 0x7ad92e57a52c4e3e83ba624622b14e3a5efa0160dd6f9a7975c43ea66bad79ea.
//
// Solidity: event NewPendingVetoer(address oldPendingVetoer, address newPendingVetoer)
func (_NounsDAO *NounsDAOFilterer) WatchNewPendingVetoer(opts *bind.WatchOpts, sink chan<- *NounsDAONewPendingVetoer) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "NewPendingVetoer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAONewPendingVetoer)
				if err := _NounsDAO.contract.UnpackLog(event, "NewPendingVetoer", log); err != nil {
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

// ParseNewPendingVetoer is a log parse operation binding the contract event 0x7ad92e57a52c4e3e83ba624622b14e3a5efa0160dd6f9a7975c43ea66bad79ea.
//
// Solidity: event NewPendingVetoer(address oldPendingVetoer, address newPendingVetoer)
func (_NounsDAO *NounsDAOFilterer) ParseNewPendingVetoer(log types.Log) (*NounsDAONewPendingVetoer, error) {
	event := new(NounsDAONewPendingVetoer)
	if err := _NounsDAO.contract.UnpackLog(event, "NewPendingVetoer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAONewVetoerIterator is returned from FilterNewVetoer and is used to iterate over the raw logs and unpacked data for NewVetoer events raised by the NounsDAO contract.
type NounsDAONewVetoerIterator struct {
	Event *NounsDAONewVetoer // Event containing the contract specifics and raw log

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
func (it *NounsDAONewVetoerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAONewVetoer)
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
		it.Event = new(NounsDAONewVetoer)
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
func (it *NounsDAONewVetoerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAONewVetoerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAONewVetoer represents a NewVetoer event raised by the NounsDAO contract.
type NounsDAONewVetoer struct {
	OldVetoer common.Address
	NewVetoer common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewVetoer is a free log retrieval operation binding the contract event 0xc5644f3588a066b15dcf6b636b74aadca57cfaccf608d9de7d8786364b7a8d02.
//
// Solidity: event NewVetoer(address oldVetoer, address newVetoer)
func (_NounsDAO *NounsDAOFilterer) FilterNewVetoer(opts *bind.FilterOpts) (*NounsDAONewVetoerIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "NewVetoer")
	if err != nil {
		return nil, err
	}
	return &NounsDAONewVetoerIterator{contract: _NounsDAO.contract, event: "NewVetoer", logs: logs, sub: sub}, nil
}

// WatchNewVetoer is a free log subscription operation binding the contract event 0xc5644f3588a066b15dcf6b636b74aadca57cfaccf608d9de7d8786364b7a8d02.
//
// Solidity: event NewVetoer(address oldVetoer, address newVetoer)
func (_NounsDAO *NounsDAOFilterer) WatchNewVetoer(opts *bind.WatchOpts, sink chan<- *NounsDAONewVetoer) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "NewVetoer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAONewVetoer)
				if err := _NounsDAO.contract.UnpackLog(event, "NewVetoer", log); err != nil {
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

// ParseNewVetoer is a log parse operation binding the contract event 0xc5644f3588a066b15dcf6b636b74aadca57cfaccf608d9de7d8786364b7a8d02.
//
// Solidity: event NewVetoer(address oldVetoer, address newVetoer)
func (_NounsDAO *NounsDAOFilterer) ParseNewVetoer(log types.Log) (*NounsDAONewVetoer, error) {
	event := new(NounsDAONewVetoer)
	if err := _NounsDAO.contract.UnpackLog(event, "NewVetoer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the NounsDAO contract.
type NounsDAOProposalCanceledIterator struct {
	Event *NounsDAOProposalCanceled // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalCanceled)
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
		it.Event = new(NounsDAOProposalCanceled)
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
func (it *NounsDAOProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalCanceled represents a ProposalCanceled event raised by the NounsDAO contract.
type NounsDAOProposalCanceled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_NounsDAO *NounsDAOFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*NounsDAOProposalCanceledIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalCanceledIterator{contract: _NounsDAO.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_NounsDAO *NounsDAOFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalCanceled)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 id)
func (_NounsDAO *NounsDAOFilterer) ParseProposalCanceled(log types.Log) (*NounsDAOProposalCanceled, error) {
	event := new(NounsDAOProposalCanceled)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the NounsDAO contract.
type NounsDAOProposalCreatedIterator struct {
	Event *NounsDAOProposalCreated // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalCreated)
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
		it.Event = new(NounsDAOProposalCreated)
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
func (it *NounsDAOProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalCreated represents a ProposalCreated event raised by the NounsDAO contract.
type NounsDAOProposalCreated struct {
	Id          *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	StartBlock  *big.Int
	EndBlock    *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_NounsDAO *NounsDAOFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*NounsDAOProposalCreatedIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalCreatedIterator{contract: _NounsDAO.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_NounsDAO *NounsDAOFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalCreated) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalCreated)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, string description)
func (_NounsDAO *NounsDAOFilterer) ParseProposalCreated(log types.Log) (*NounsDAOProposalCreated, error) {
	event := new(NounsDAOProposalCreated)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalCreatedWithRequirementsIterator is returned from FilterProposalCreatedWithRequirements and is used to iterate over the raw logs and unpacked data for ProposalCreatedWithRequirements events raised by the NounsDAO contract.
type NounsDAOProposalCreatedWithRequirementsIterator struct {
	Event *NounsDAOProposalCreatedWithRequirements // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalCreatedWithRequirementsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalCreatedWithRequirements)
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
		it.Event = new(NounsDAOProposalCreatedWithRequirements)
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
func (it *NounsDAOProposalCreatedWithRequirementsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalCreatedWithRequirementsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalCreatedWithRequirements represents a ProposalCreatedWithRequirements event raised by the NounsDAO contract.
type NounsDAOProposalCreatedWithRequirements struct {
	Id                *big.Int
	Proposer          common.Address
	Targets           []common.Address
	Values            []*big.Int
	Signatures        []string
	Calldatas         [][]byte
	StartBlock        *big.Int
	EndBlock          *big.Int
	ProposalThreshold *big.Int
	QuorumVotes       *big.Int
	Description       string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterProposalCreatedWithRequirements is a free log retrieval operation binding the contract event 0x6af0134faa0f9290c1d686d55012aca80302d31d5c856e4bc7954f7613dc7f87.
//
// Solidity: event ProposalCreatedWithRequirements(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, uint256 proposalThreshold, uint256 quorumVotes, string description)
func (_NounsDAO *NounsDAOFilterer) FilterProposalCreatedWithRequirements(opts *bind.FilterOpts) (*NounsDAOProposalCreatedWithRequirementsIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalCreatedWithRequirements")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalCreatedWithRequirementsIterator{contract: _NounsDAO.contract, event: "ProposalCreatedWithRequirements", logs: logs, sub: sub}, nil
}

// WatchProposalCreatedWithRequirements is a free log subscription operation binding the contract event 0x6af0134faa0f9290c1d686d55012aca80302d31d5c856e4bc7954f7613dc7f87.
//
// Solidity: event ProposalCreatedWithRequirements(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, uint256 proposalThreshold, uint256 quorumVotes, string description)
func (_NounsDAO *NounsDAOFilterer) WatchProposalCreatedWithRequirements(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalCreatedWithRequirements) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalCreatedWithRequirements")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalCreatedWithRequirements)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalCreatedWithRequirements", log); err != nil {
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

// ParseProposalCreatedWithRequirements is a log parse operation binding the contract event 0x6af0134faa0f9290c1d686d55012aca80302d31d5c856e4bc7954f7613dc7f87.
//
// Solidity: event ProposalCreatedWithRequirements(uint256 id, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 startBlock, uint256 endBlock, uint256 proposalThreshold, uint256 quorumVotes, string description)
func (_NounsDAO *NounsDAOFilterer) ParseProposalCreatedWithRequirements(log types.Log) (*NounsDAOProposalCreatedWithRequirements, error) {
	event := new(NounsDAOProposalCreatedWithRequirements)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalCreatedWithRequirements", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the NounsDAO contract.
type NounsDAOProposalExecutedIterator struct {
	Event *NounsDAOProposalExecuted // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalExecuted)
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
		it.Event = new(NounsDAOProposalExecuted)
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
func (it *NounsDAOProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalExecuted represents a ProposalExecuted event raised by the NounsDAO contract.
type NounsDAOProposalExecuted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_NounsDAO *NounsDAOFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*NounsDAOProposalExecutedIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalExecutedIterator{contract: _NounsDAO.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_NounsDAO *NounsDAOFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalExecuted)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 id)
func (_NounsDAO *NounsDAOFilterer) ParseProposalExecuted(log types.Log) (*NounsDAOProposalExecuted, error) {
	event := new(NounsDAOProposalExecuted)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the NounsDAO contract.
type NounsDAOProposalQueuedIterator struct {
	Event *NounsDAOProposalQueued // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalQueued)
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
		it.Event = new(NounsDAOProposalQueued)
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
func (it *NounsDAOProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalQueued represents a ProposalQueued event raised by the NounsDAO contract.
type NounsDAOProposalQueued struct {
	Id  *big.Int
	Eta *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_NounsDAO *NounsDAOFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*NounsDAOProposalQueuedIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalQueuedIterator{contract: _NounsDAO.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_NounsDAO *NounsDAOFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalQueued) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalQueued)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 id, uint256 eta)
func (_NounsDAO *NounsDAOFilterer) ParseProposalQueued(log types.Log) (*NounsDAOProposalQueued, error) {
	event := new(NounsDAOProposalQueued)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalThresholdBPSSetIterator is returned from FilterProposalThresholdBPSSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdBPSSet events raised by the NounsDAO contract.
type NounsDAOProposalThresholdBPSSetIterator struct {
	Event *NounsDAOProposalThresholdBPSSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalThresholdBPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalThresholdBPSSet)
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
		it.Event = new(NounsDAOProposalThresholdBPSSet)
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
func (it *NounsDAOProposalThresholdBPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalThresholdBPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalThresholdBPSSet represents a ProposalThresholdBPSSet event raised by the NounsDAO contract.
type NounsDAOProposalThresholdBPSSet struct {
	OldProposalThresholdBPS *big.Int
	NewProposalThresholdBPS *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdBPSSet is a free log retrieval operation binding the contract event 0xfc216faa269bf440fb06aa490693f409461bde9cdcb949c7b9f2cb79589e7a58.
//
// Solidity: event ProposalThresholdBPSSet(uint256 oldProposalThresholdBPS, uint256 newProposalThresholdBPS)
func (_NounsDAO *NounsDAOFilterer) FilterProposalThresholdBPSSet(opts *bind.FilterOpts) (*NounsDAOProposalThresholdBPSSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalThresholdBPSSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalThresholdBPSSetIterator{contract: _NounsDAO.contract, event: "ProposalThresholdBPSSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdBPSSet is a free log subscription operation binding the contract event 0xfc216faa269bf440fb06aa490693f409461bde9cdcb949c7b9f2cb79589e7a58.
//
// Solidity: event ProposalThresholdBPSSet(uint256 oldProposalThresholdBPS, uint256 newProposalThresholdBPS)
func (_NounsDAO *NounsDAOFilterer) WatchProposalThresholdBPSSet(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalThresholdBPSSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalThresholdBPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalThresholdBPSSet)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalThresholdBPSSet", log); err != nil {
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

// ParseProposalThresholdBPSSet is a log parse operation binding the contract event 0xfc216faa269bf440fb06aa490693f409461bde9cdcb949c7b9f2cb79589e7a58.
//
// Solidity: event ProposalThresholdBPSSet(uint256 oldProposalThresholdBPS, uint256 newProposalThresholdBPS)
func (_NounsDAO *NounsDAOFilterer) ParseProposalThresholdBPSSet(log types.Log) (*NounsDAOProposalThresholdBPSSet, error) {
	event := new(NounsDAOProposalThresholdBPSSet)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalThresholdBPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOProposalVetoedIterator is returned from FilterProposalVetoed and is used to iterate over the raw logs and unpacked data for ProposalVetoed events raised by the NounsDAO contract.
type NounsDAOProposalVetoedIterator struct {
	Event *NounsDAOProposalVetoed // Event containing the contract specifics and raw log

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
func (it *NounsDAOProposalVetoedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOProposalVetoed)
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
		it.Event = new(NounsDAOProposalVetoed)
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
func (it *NounsDAOProposalVetoedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOProposalVetoedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOProposalVetoed represents a ProposalVetoed event raised by the NounsDAO contract.
type NounsDAOProposalVetoed struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalVetoed is a free log retrieval operation binding the contract event 0xde0cea2a3a0097cc3d981d40c375407760e85bc9c5e69aea449ac3885f8615c6.
//
// Solidity: event ProposalVetoed(uint256 id)
func (_NounsDAO *NounsDAOFilterer) FilterProposalVetoed(opts *bind.FilterOpts) (*NounsDAOProposalVetoedIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "ProposalVetoed")
	if err != nil {
		return nil, err
	}
	return &NounsDAOProposalVetoedIterator{contract: _NounsDAO.contract, event: "ProposalVetoed", logs: logs, sub: sub}, nil
}

// WatchProposalVetoed is a free log subscription operation binding the contract event 0xde0cea2a3a0097cc3d981d40c375407760e85bc9c5e69aea449ac3885f8615c6.
//
// Solidity: event ProposalVetoed(uint256 id)
func (_NounsDAO *NounsDAOFilterer) WatchProposalVetoed(opts *bind.WatchOpts, sink chan<- *NounsDAOProposalVetoed) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "ProposalVetoed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOProposalVetoed)
				if err := _NounsDAO.contract.UnpackLog(event, "ProposalVetoed", log); err != nil {
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

// ParseProposalVetoed is a log parse operation binding the contract event 0xde0cea2a3a0097cc3d981d40c375407760e85bc9c5e69aea449ac3885f8615c6.
//
// Solidity: event ProposalVetoed(uint256 id)
func (_NounsDAO *NounsDAOFilterer) ParseProposalVetoed(log types.Log) (*NounsDAOProposalVetoed, error) {
	event := new(NounsDAOProposalVetoed)
	if err := _NounsDAO.contract.UnpackLog(event, "ProposalVetoed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOQuorumCoefficientSetIterator is returned from FilterQuorumCoefficientSet and is used to iterate over the raw logs and unpacked data for QuorumCoefficientSet events raised by the NounsDAO contract.
type NounsDAOQuorumCoefficientSetIterator struct {
	Event *NounsDAOQuorumCoefficientSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOQuorumCoefficientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOQuorumCoefficientSet)
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
		it.Event = new(NounsDAOQuorumCoefficientSet)
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
func (it *NounsDAOQuorumCoefficientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOQuorumCoefficientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOQuorumCoefficientSet represents a QuorumCoefficientSet event raised by the NounsDAO contract.
type NounsDAOQuorumCoefficientSet struct {
	OldQuorumCoefficient uint32
	NewQuorumCoefficient uint32
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterQuorumCoefficientSet is a free log retrieval operation binding the contract event 0x5e3adb1066359dafa23c629f245d93543856115700821dcb4debc416f393c469.
//
// Solidity: event QuorumCoefficientSet(uint32 oldQuorumCoefficient, uint32 newQuorumCoefficient)
func (_NounsDAO *NounsDAOFilterer) FilterQuorumCoefficientSet(opts *bind.FilterOpts) (*NounsDAOQuorumCoefficientSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "QuorumCoefficientSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOQuorumCoefficientSetIterator{contract: _NounsDAO.contract, event: "QuorumCoefficientSet", logs: logs, sub: sub}, nil
}

// WatchQuorumCoefficientSet is a free log subscription operation binding the contract event 0x5e3adb1066359dafa23c629f245d93543856115700821dcb4debc416f393c469.
//
// Solidity: event QuorumCoefficientSet(uint32 oldQuorumCoefficient, uint32 newQuorumCoefficient)
func (_NounsDAO *NounsDAOFilterer) WatchQuorumCoefficientSet(opts *bind.WatchOpts, sink chan<- *NounsDAOQuorumCoefficientSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "QuorumCoefficientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOQuorumCoefficientSet)
				if err := _NounsDAO.contract.UnpackLog(event, "QuorumCoefficientSet", log); err != nil {
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

// ParseQuorumCoefficientSet is a log parse operation binding the contract event 0x5e3adb1066359dafa23c629f245d93543856115700821dcb4debc416f393c469.
//
// Solidity: event QuorumCoefficientSet(uint32 oldQuorumCoefficient, uint32 newQuorumCoefficient)
func (_NounsDAO *NounsDAOFilterer) ParseQuorumCoefficientSet(log types.Log) (*NounsDAOQuorumCoefficientSet, error) {
	event := new(NounsDAOQuorumCoefficientSet)
	if err := _NounsDAO.contract.UnpackLog(event, "QuorumCoefficientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOQuorumVotesBPSSetIterator is returned from FilterQuorumVotesBPSSet and is used to iterate over the raw logs and unpacked data for QuorumVotesBPSSet events raised by the NounsDAO contract.
type NounsDAOQuorumVotesBPSSetIterator struct {
	Event *NounsDAOQuorumVotesBPSSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOQuorumVotesBPSSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOQuorumVotesBPSSet)
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
		it.Event = new(NounsDAOQuorumVotesBPSSet)
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
func (it *NounsDAOQuorumVotesBPSSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOQuorumVotesBPSSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOQuorumVotesBPSSet represents a QuorumVotesBPSSet event raised by the NounsDAO contract.
type NounsDAOQuorumVotesBPSSet struct {
	OldQuorumVotesBPS *big.Int
	NewQuorumVotesBPS *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterQuorumVotesBPSSet is a free log retrieval operation binding the contract event 0xd73ab1b53ca7a080713bcecd1a0acb2066a6a6c3d2fd6d78b67ae5005e652d9b.
//
// Solidity: event QuorumVotesBPSSet(uint256 oldQuorumVotesBPS, uint256 newQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) FilterQuorumVotesBPSSet(opts *bind.FilterOpts) (*NounsDAOQuorumVotesBPSSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "QuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOQuorumVotesBPSSetIterator{contract: _NounsDAO.contract, event: "QuorumVotesBPSSet", logs: logs, sub: sub}, nil
}

// WatchQuorumVotesBPSSet is a free log subscription operation binding the contract event 0xd73ab1b53ca7a080713bcecd1a0acb2066a6a6c3d2fd6d78b67ae5005e652d9b.
//
// Solidity: event QuorumVotesBPSSet(uint256 oldQuorumVotesBPS, uint256 newQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) WatchQuorumVotesBPSSet(opts *bind.WatchOpts, sink chan<- *NounsDAOQuorumVotesBPSSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "QuorumVotesBPSSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOQuorumVotesBPSSet)
				if err := _NounsDAO.contract.UnpackLog(event, "QuorumVotesBPSSet", log); err != nil {
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

// ParseQuorumVotesBPSSet is a log parse operation binding the contract event 0xd73ab1b53ca7a080713bcecd1a0acb2066a6a6c3d2fd6d78b67ae5005e652d9b.
//
// Solidity: event QuorumVotesBPSSet(uint256 oldQuorumVotesBPS, uint256 newQuorumVotesBPS)
func (_NounsDAO *NounsDAOFilterer) ParseQuorumVotesBPSSet(log types.Log) (*NounsDAOQuorumVotesBPSSet, error) {
	event := new(NounsDAOQuorumVotesBPSSet)
	if err := _NounsDAO.contract.UnpackLog(event, "QuorumVotesBPSSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAORefundableVoteIterator is returned from FilterRefundableVote and is used to iterate over the raw logs and unpacked data for RefundableVote events raised by the NounsDAO contract.
type NounsDAORefundableVoteIterator struct {
	Event *NounsDAORefundableVote // Event containing the contract specifics and raw log

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
func (it *NounsDAORefundableVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAORefundableVote)
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
		it.Event = new(NounsDAORefundableVote)
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
func (it *NounsDAORefundableVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAORefundableVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAORefundableVote represents a RefundableVote event raised by the NounsDAO contract.
type NounsDAORefundableVote struct {
	Voter        common.Address
	RefundAmount *big.Int
	RefundSent   bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRefundableVote is a free log retrieval operation binding the contract event 0xfabef36fd46c4c3a6ad676521be5367a4dfdbf3faa68d8e826003b1752d68f4f.
//
// Solidity: event RefundableVote(address indexed voter, uint256 refundAmount, bool refundSent)
func (_NounsDAO *NounsDAOFilterer) FilterRefundableVote(opts *bind.FilterOpts, voter []common.Address) (*NounsDAORefundableVoteIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "RefundableVote", voterRule)
	if err != nil {
		return nil, err
	}
	return &NounsDAORefundableVoteIterator{contract: _NounsDAO.contract, event: "RefundableVote", logs: logs, sub: sub}, nil
}

// WatchRefundableVote is a free log subscription operation binding the contract event 0xfabef36fd46c4c3a6ad676521be5367a4dfdbf3faa68d8e826003b1752d68f4f.
//
// Solidity: event RefundableVote(address indexed voter, uint256 refundAmount, bool refundSent)
func (_NounsDAO *NounsDAOFilterer) WatchRefundableVote(opts *bind.WatchOpts, sink chan<- *NounsDAORefundableVote, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "RefundableVote", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAORefundableVote)
				if err := _NounsDAO.contract.UnpackLog(event, "RefundableVote", log); err != nil {
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

// ParseRefundableVote is a log parse operation binding the contract event 0xfabef36fd46c4c3a6ad676521be5367a4dfdbf3faa68d8e826003b1752d68f4f.
//
// Solidity: event RefundableVote(address indexed voter, uint256 refundAmount, bool refundSent)
func (_NounsDAO *NounsDAOFilterer) ParseRefundableVote(log types.Log) (*NounsDAORefundableVote, error) {
	event := new(NounsDAORefundableVote)
	if err := _NounsDAO.contract.UnpackLog(event, "RefundableVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the NounsDAO contract.
type NounsDAOVoteCastIterator struct {
	Event *NounsDAOVoteCast // Event containing the contract specifics and raw log

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
func (it *NounsDAOVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOVoteCast)
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
		it.Event = new(NounsDAOVoteCast)
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
func (it *NounsDAOVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOVoteCast represents a VoteCast event raised by the NounsDAO contract.
type NounsDAOVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Votes      *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 votes, string reason)
func (_NounsDAO *NounsDAOFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*NounsDAOVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &NounsDAOVoteCastIterator{contract: _NounsDAO.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 votes, string reason)
func (_NounsDAO *NounsDAOFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *NounsDAOVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOVoteCast)
				if err := _NounsDAO.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 votes, string reason)
func (_NounsDAO *NounsDAOFilterer) ParseVoteCast(log types.Log) (*NounsDAOVoteCast, error) {
	event := new(NounsDAOVoteCast)
	if err := _NounsDAO.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the NounsDAO contract.
type NounsDAOVotingDelaySetIterator struct {
	Event *NounsDAOVotingDelaySet // Event containing the contract specifics and raw log

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
func (it *NounsDAOVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOVotingDelaySet)
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
		it.Event = new(NounsDAOVotingDelaySet)
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
func (it *NounsDAOVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOVotingDelaySet represents a VotingDelaySet event raised by the NounsDAO contract.
type NounsDAOVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_NounsDAO *NounsDAOFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*NounsDAOVotingDelaySetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOVotingDelaySetIterator{contract: _NounsDAO.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_NounsDAO *NounsDAOFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *NounsDAOVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOVotingDelaySet)
				if err := _NounsDAO.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
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

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_NounsDAO *NounsDAOFilterer) ParseVotingDelaySet(log types.Log) (*NounsDAOVotingDelaySet, error) {
	event := new(NounsDAOVotingDelaySet)
	if err := _NounsDAO.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the NounsDAO contract.
type NounsDAOVotingPeriodSetIterator struct {
	Event *NounsDAOVotingPeriodSet // Event containing the contract specifics and raw log

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
func (it *NounsDAOVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOVotingPeriodSet)
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
		it.Event = new(NounsDAOVotingPeriodSet)
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
func (it *NounsDAOVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOVotingPeriodSet represents a VotingPeriodSet event raised by the NounsDAO contract.
type NounsDAOVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_NounsDAO *NounsDAOFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*NounsDAOVotingPeriodSetIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &NounsDAOVotingPeriodSetIterator{contract: _NounsDAO.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_NounsDAO *NounsDAOFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *NounsDAOVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOVotingPeriodSet)
				if err := _NounsDAO.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
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

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_NounsDAO *NounsDAOFilterer) ParseVotingPeriodSet(log types.Log) (*NounsDAOVotingPeriodSet, error) {
	event := new(NounsDAOVotingPeriodSet)
	if err := _NounsDAO.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NounsDAOWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NounsDAO contract.
type NounsDAOWithdrawIterator struct {
	Event *NounsDAOWithdraw // Event containing the contract specifics and raw log

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
func (it *NounsDAOWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NounsDAOWithdraw)
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
		it.Event = new(NounsDAOWithdraw)
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
func (it *NounsDAOWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NounsDAOWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NounsDAOWithdraw represents a Withdraw event raised by the NounsDAO contract.
type NounsDAOWithdraw struct {
	Amount *big.Int
	Sent   bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x2aeb20ed0ead73e7bc740154a0b979547bc9e00691d84a700e6454ada9fe4679.
//
// Solidity: event Withdraw(uint256 amount, bool sent)
func (_NounsDAO *NounsDAOFilterer) FilterWithdraw(opts *bind.FilterOpts) (*NounsDAOWithdrawIterator, error) {

	logs, sub, err := _NounsDAO.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &NounsDAOWithdrawIterator{contract: _NounsDAO.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x2aeb20ed0ead73e7bc740154a0b979547bc9e00691d84a700e6454ada9fe4679.
//
// Solidity: event Withdraw(uint256 amount, bool sent)
func (_NounsDAO *NounsDAOFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NounsDAOWithdraw) (event.Subscription, error) {

	logs, sub, err := _NounsDAO.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NounsDAOWithdraw)
				if err := _NounsDAO.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x2aeb20ed0ead73e7bc740154a0b979547bc9e00691d84a700e6454ada9fe4679.
//
// Solidity: event Withdraw(uint256 amount, bool sent)
func (_NounsDAO *NounsDAOFilterer) ParseWithdraw(log types.Log) (*NounsDAOWithdraw, error) {
	event := new(NounsDAOWithdraw)
	if err := _NounsDAO.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
