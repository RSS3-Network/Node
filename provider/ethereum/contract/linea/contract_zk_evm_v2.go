// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package linea

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

// IL1MessageServiceClaimMessageWithProofParams is an auto generated low-level Go binding around an user-defined struct.
type IL1MessageServiceClaimMessageWithProofParams struct {
	Proof         [][32]byte
	MessageNumber *big.Int
	LeafIndex     uint32
	From          common.Address
	To            common.Address
	Fee           *big.Int
	Value         *big.Int
	FeeRecipient  common.Address
	MerkleRoot    [32]byte
	Data          []byte
}

// ILineaRollupBlobSubmissionData is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupBlobSubmissionData struct {
	SubmissionData      ILineaRollupSupportingSubmissionDataV2
	DataEvaluationClaim *big.Int
	KzgCommitment       []byte
	KzgProof            []byte
}

// ILineaRollupFinalizationDataV2 is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupFinalizationDataV2 struct {
	ParentStateRootHash                     [32]byte
	LastFinalizedShnarf                     [32]byte
	FinalBlockInData                        *big.Int
	ShnarfData                              ILineaRollupShnarfData
	LastFinalizedTimestamp                  *big.Int
	FinalTimestamp                          *big.Int
	LastFinalizedL1RollingHash              [32]byte
	L1RollingHash                           [32]byte
	LastFinalizedL1RollingHashMessageNumber *big.Int
	L1RollingHashMessageNumber              *big.Int
	L2MerkleTreesDepth                      *big.Int
	L2MerkleRoots                           [][32]byte
	L2MessagingBlocksOffsets                []byte
}

// ILineaRollupShnarfData is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupShnarfData struct {
	ParentShnarf        [32]byte
	SnarkHash           [32]byte
	FinalStateRootHash  [32]byte
	DataEvaluationPoint [32]byte
	DataEvaluationClaim [32]byte
}

// ILineaRollupSubmissionDataV2 is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupSubmissionDataV2 struct {
	FinalStateRootHash [32]byte
	FirstBlockInData   *big.Int
	FinalBlockInData   *big.Int
	SnarkHash          [32]byte
	CompressedData     []byte
}

// ILineaRollupSupportingSubmissionDataV2 is an auto generated low-level Go binding around an user-defined struct.
type ILineaRollupSupportingSubmissionDataV2 struct {
	FinalStateRootHash [32]byte
	FirstBlockInData   *big.Int
	FinalBlockInData   *big.Int
	SnarkHash          [32]byte
}

// ZKEVMV2MetaData contains all meta data concerning the ZKEVMV2 contract.
var ZKEVMV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BlobSubmissionDataIsMissing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BytesLengthNotMultipleOf32\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bytesLength\",\"type\":\"uint256\"}],\"name\":\"BytesLengthNotMultipleOfTwo\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"currentDataHash\",\"type\":\"bytes32\"}],\"name\":\"DataAlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"DataStartingBlockDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBlobData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"EmptyBlobDataAtIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptySubmissionData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FeePaymentFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"FinalBlockDoesNotMatchShnarfFinalBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint256\"}],\"name\":\"FinalBlockNumberLessThanOrEqualToLastFinalizedBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FinalBlockStateEqualsZeroHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"FinalShnarfWrong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"l2BlockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBlockTimestamp\",\"type\":\"uint256\"}],\"name\":\"FinalizationInTheFuture\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"FinalizationStateIncorrect\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"firstBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockNumber\",\"type\":\"uint256\"}],\"name\":\"FirstBlockGreaterThanFinalBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"firstBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedBlock\",\"type\":\"uint256\"}],\"name\":\"FirstBlockLessThanOrEqualToLastFinalizedBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FirstByteIsNotZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMerkleProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProofType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"IsNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"IsPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"name\":\"L1RollingHashDoesNotExistOnL1\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"L2MerkleRootAlreadyAnchored\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"L2MerkleRootDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"LastFinalizedShnarfWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitIsZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageIndex\",\"type\":\"uint256\"}],\"name\":\"MessageAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageDoesNotExistOrHasAlreadyBeenClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"MessageSendingFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"name\":\"MissingMessageNumberForRollingHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"}],\"name\":\"MissingRollingHashForMessageNumber\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"ParentStateRootHashInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PeriodIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PointEvaluationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fieldElements\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blsCurveModulus\",\"type\":\"uint256\"}],\"name\":\"PointEvaluationResponseInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"}],\"name\":\"PrecompileReturnDataLengthWrong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProofIsEmpty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"ProofLengthDifferentThanMerkleDepth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shnarfsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockNumbers\",\"type\":\"uint256\"}],\"name\":\"ShnarfAndFinalBlockNumberLengthsMismatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SnarkHashIsZeroHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartingRootHashDoesNotMatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueSentTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressNotAllowed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"resettingAddress\",\"type\":\"address\"}],\"name\":\"AmountUsedInPeriodReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"stateRootHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"finalizedWithProof\",\"type\":\"bool\"}],\"name\":\"BlockFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastBlockFinalized\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"startingRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"finalRootHash\",\"type\":\"bytes32\"}],\"name\":\"BlocksVerificationDone\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"lastBlockFinalized\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"startingRootHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"finalRootHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withProof\",\"type\":\"bool\"}],\"name\":\"DataFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"DataSubmittedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"messageHashes\",\"type\":\"bytes32[]\"}],\"name\":\"L1L2MessagesReceivedOnL2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"L2L1MessageHashAddedToInbox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"l2MerkleRoot\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"treeDepth\",\"type\":\"uint256\"}],\"name\":\"L2MerkleRootAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"l2Block\",\"type\":\"uint256\"}],\"name\":\"L2MessagingBlockAnchored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"amountChangeBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"amountUsedLoweredToLimit\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"usedAmountResetToZero\",\"type\":\"bool\"}],\"name\":\"LimitAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_messageHash\",\"type\":\"bytes32\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"periodInSeconds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limitInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentPeriodEnd\",\"type\":\"uint256\"}],\"name\":\"RateLimitInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"RollingHashUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"systemMigrationBlock\",\"type\":\"uint256\"}],\"name\":\"SystemMigrationBlockInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageSender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pauseType\",\"type\":\"uint256\"}],\"name\":\"UnPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"verifierSetBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldVerifierAddress\",\"type\":\"address\"}],\"name\":\"VerifierAddressChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENERAL_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_SHNARF\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L1_L2_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"L2_L1_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_RECEIVED\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_SENT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OUTBOX_STATUS_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVING_SYSTEM_PAUSE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RATE_LIMIT_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_SETTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"claimMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"leafIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structIL1MessageService.ClaimMessageWithProofParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"claimMessageWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFinalizedShnarf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentFinalizedState\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2BlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2StoredL1MessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentL2StoredL1RollingHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodAmountInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataEndingBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"endingBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataFinalStateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataParents\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"parentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataShnarfHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarfHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"dataStartingBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startingBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_aggregatedProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastFinalizedShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockInData\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationPoint\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationClaim\",\"type\":\"bytes32\"}],\"internalType\":\"structILineaRollup.ShnarfData\",\"name\":\"shnarfData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lastFinalizedL1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedL1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2MerkleTreesDepth\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2MerkleRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"l2MessagingBlocksOffsets\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.FinalizationDataV2\",\"name\":\"_finalizationData\",\"type\":\"tuple\"}],\"name\":\"finalizeBlocksWithProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"lastFinalizedShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockInData\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationPoint\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataEvaluationClaim\",\"type\":\"bytes32\"}],\"internalType\":\"structILineaRollup.ShnarfData\",\"name\":\"shnarfData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"lastFinalizedL1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"l1RollingHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"lastFinalizedL1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l1RollingHashMessageNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"l2MerkleTreesDepth\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"l2MerkleRoots\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"l2MessagingBlocksOffsets\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.FinalizationDataV2\",\"name\":\"_finalizationData\",\"type\":\"tuple\"}],\"name\":\"finalizeBlocksWithoutProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"inboxL2L1MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageStatus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_initialStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_initialL2BlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_defaultVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_securityCouncil\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitPeriodInSeconds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rateLimitAmountInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_genesisTimestamp\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_shnarfs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_finalBlockNumbers\",\"type\":\"uint256[]\"}],\"name\":\"initializeParentShnarfsAndFinalizedState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_messageNumber\",\"type\":\"uint256\"}],\"name\":\"isMessageClaimed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"l2MerkleRootsDepths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"treeDepth\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextMessageNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"outboxL1L2MessageStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"messageStatus\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"pauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"pauseType\",\"type\":\"bytes32\"}],\"name\":\"pauseTypeStatuses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"pauseStatus\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"periodInSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetAmountUsedInPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"resetRateLimitAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"messageNumber\",\"type\":\"uint256\"}],\"name\":\"rollingHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollingHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newVerifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"setVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"shnarf\",\"type\":\"bytes32\"}],\"name\":\"shnarfFinalBlockNumbers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"finalBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"stateRootHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"stateRootHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"firstBlockInData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockInData\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"}],\"internalType\":\"structILineaRollup.SupportingSubmissionDataV2\",\"name\":\"submissionData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"dataEvaluationClaim\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"kzgCommitment\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"kzgProof\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.BlobSubmissionData[]\",\"name\":\"_blobSubmissionData\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_finalBlobShnarf\",\"type\":\"bytes32\"}],\"name\":\"submitBlobs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"finalStateRootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"firstBlockInData\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalBlockInData\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"snarkHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"compressedData\",\"type\":\"bytes\"}],\"internalType\":\"structILineaRollup.SubmissionDataV2\",\"name\":\"_submissionData\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_parentShnarf\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_expectedShnarf\",\"type\":\"bytes32\"}],\"name\":\"submitDataAsCalldata\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"systemMigrationBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_pauseType\",\"type\":\"uint8\"}],\"name\":\"unPauseByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proofType\",\"type\":\"uint256\"}],\"name\":\"unsetVerifierAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proofType\",\"type\":\"uint256\"}],\"name\":\"verifiers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"verifierAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZKEVMV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ZKEVMV2MetaData.ABI instead.
var ZKEVMV2ABI = ZKEVMV2MetaData.ABI

// ZKEVMV2 is an auto generated Go binding around an Ethereum contract.
type ZKEVMV2 struct {
	ZKEVMV2Caller     // Read-only binding to the contract
	ZKEVMV2Transactor // Write-only binding to the contract
	ZKEVMV2Filterer   // Log filterer for contract events
}

// ZKEVMV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ZKEVMV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ZKEVMV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZKEVMV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZKEVMV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZKEVMV2Session struct {
	Contract     *ZKEVMV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZKEVMV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZKEVMV2CallerSession struct {
	Contract *ZKEVMV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ZKEVMV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZKEVMV2TransactorSession struct {
	Contract     *ZKEVMV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ZKEVMV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ZKEVMV2Raw struct {
	Contract *ZKEVMV2 // Generic contract binding to access the raw methods on
}

// ZKEVMV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZKEVMV2CallerRaw struct {
	Contract *ZKEVMV2Caller // Generic read-only contract binding to access the raw methods on
}

// ZKEVMV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZKEVMV2TransactorRaw struct {
	Contract *ZKEVMV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewZKEVMV2 creates a new instance of ZKEVMV2, bound to a specific deployed contract.
func NewZKEVMV2(address common.Address, backend bind.ContractBackend) (*ZKEVMV2, error) {
	contract, err := bindZKEVMV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2{ZKEVMV2Caller: ZKEVMV2Caller{contract: contract}, ZKEVMV2Transactor: ZKEVMV2Transactor{contract: contract}, ZKEVMV2Filterer: ZKEVMV2Filterer{contract: contract}}, nil
}

// NewZKEVMV2Caller creates a new read-only instance of ZKEVMV2, bound to a specific deployed contract.
func NewZKEVMV2Caller(address common.Address, caller bind.ContractCaller) (*ZKEVMV2Caller, error) {
	contract, err := bindZKEVMV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2Caller{contract: contract}, nil
}

// NewZKEVMV2Transactor creates a new write-only instance of ZKEVMV2, bound to a specific deployed contract.
func NewZKEVMV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ZKEVMV2Transactor, error) {
	contract, err := bindZKEVMV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2Transactor{contract: contract}, nil
}

// NewZKEVMV2Filterer creates a new log filterer instance of ZKEVMV2, bound to a specific deployed contract.
func NewZKEVMV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ZKEVMV2Filterer, error) {
	contract, err := bindZKEVMV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2Filterer{contract: contract}, nil
}

// bindZKEVMV2 binds a generic wrapper to an already deployed contract.
func bindZKEVMV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZKEVMV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEVMV2 *ZKEVMV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEVMV2.Contract.ZKEVMV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEVMV2 *ZKEVMV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ZKEVMV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEVMV2 *ZKEVMV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ZKEVMV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZKEVMV2 *ZKEVMV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZKEVMV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZKEVMV2 *ZKEVMV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZKEVMV2 *ZKEVMV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.DEFAULTADMINROLE(&_ZKEVMV2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.DEFAULTADMINROLE(&_ZKEVMV2.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) GENERALPAUSETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "GENERAL_PAUSE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) GENERALPAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.GENERALPAUSETYPE(&_ZKEVMV2.CallOpts)
}

// GENERALPAUSETYPE is a free data retrieval call binding the contract method 0x6a637967.
//
// Solidity: function GENERAL_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) GENERALPAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.GENERALPAUSETYPE(&_ZKEVMV2.CallOpts)
}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) GENESISSHNARF(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "GENESIS_SHNARF")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) GENESISSHNARF() ([32]byte, error) {
	return _ZKEVMV2.Contract.GENESISSHNARF(&_ZKEVMV2.CallOpts)
}

// GENESISSHNARF is a free data retrieval call binding the contract method 0xe97a1e9e.
//
// Solidity: function GENESIS_SHNARF() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) GENESISSHNARF() ([32]byte, error) {
	return _ZKEVMV2.Contract.GENESISSHNARF(&_ZKEVMV2.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) INBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "INBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) INBOXSTATUSRECEIVED() (uint8, error) {
	return _ZKEVMV2.Contract.INBOXSTATUSRECEIVED(&_ZKEVMV2.CallOpts)
}

// INBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x48922ab7.
//
// Solidity: function INBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) INBOXSTATUSRECEIVED() (uint8, error) {
	return _ZKEVMV2.Contract.INBOXSTATUSRECEIVED(&_ZKEVMV2.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) INBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "INBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZKEVMV2.Contract.INBOXSTATUSUNKNOWN(&_ZKEVMV2.CallOpts)
}

// INBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x7d1e8c55.
//
// Solidity: function INBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) INBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZKEVMV2.Contract.INBOXSTATUSUNKNOWN(&_ZKEVMV2.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) L1L2PAUSETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "L1_L2_PAUSE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) L1L2PAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.L1L2PAUSETYPE(&_ZKEVMV2.CallOpts)
}

// L1L2PAUSETYPE is a free data retrieval call binding the contract method 0x11314d0f.
//
// Solidity: function L1_L2_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) L1L2PAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.L1L2PAUSETYPE(&_ZKEVMV2.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) L2L1PAUSETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "L2_L1_PAUSE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) L2L1PAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.L2L1PAUSETYPE(&_ZKEVMV2.CallOpts)
}

// L2L1PAUSETYPE is a free data retrieval call binding the contract method 0xabd6230d.
//
// Solidity: function L2_L1_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) L2L1PAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.L2L1PAUSETYPE(&_ZKEVMV2.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) OPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) OPERATORROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.OPERATORROLE(&_ZKEVMV2.CallOpts)
}

// OPERATORROLE is a free data retrieval call binding the contract method 0xf5b541a6.
//
// Solidity: function OPERATOR_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) OPERATORROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.OPERATORROLE(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) OUTBOXSTATUSRECEIVED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "OUTBOX_STATUS_RECEIVED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSRECEIVED(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSRECEIVED is a free data retrieval call binding the contract method 0x73bd07b7.
//
// Solidity: function OUTBOX_STATUS_RECEIVED() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) OUTBOXSTATUSRECEIVED() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSRECEIVED(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) OUTBOXSTATUSSENT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "OUTBOX_STATUS_SENT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) OUTBOXSTATUSSENT() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSSENT(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSSENT is a free data retrieval call binding the contract method 0x5b7eb4bd.
//
// Solidity: function OUTBOX_STATUS_SENT() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) OUTBOXSTATUSSENT() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSSENT(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) OUTBOXSTATUSUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "OUTBOX_STATUS_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSUNKNOWN(&_ZKEVMV2.CallOpts)
}

// OUTBOXSTATUSUNKNOWN is a free data retrieval call binding the contract method 0x986fcddd.
//
// Solidity: function OUTBOX_STATUS_UNKNOWN() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) OUTBOXSTATUSUNKNOWN() (uint8, error) {
	return _ZKEVMV2.Contract.OUTBOXSTATUSUNKNOWN(&_ZKEVMV2.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) PAUSEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "PAUSE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) PAUSEMANAGERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.PAUSEMANAGERROLE(&_ZKEVMV2.CallOpts)
}

// PAUSEMANAGERROLE is a free data retrieval call binding the contract method 0xd84f91e8.
//
// Solidity: function PAUSE_MANAGER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) PAUSEMANAGERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.PAUSEMANAGERROLE(&_ZKEVMV2.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Caller) PROVINGSYSTEMPAUSETYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "PROVING_SYSTEM_PAUSE_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2Session) PROVINGSYSTEMPAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.PROVINGSYSTEMPAUSETYPE(&_ZKEVMV2.CallOpts)
}

// PROVINGSYSTEMPAUSETYPE is a free data retrieval call binding the contract method 0xb4a5a4b7.
//
// Solidity: function PROVING_SYSTEM_PAUSE_TYPE() view returns(uint8)
func (_ZKEVMV2 *ZKEVMV2CallerSession) PROVINGSYSTEMPAUSETYPE() (uint8, error) {
	return _ZKEVMV2.Contract.PROVINGSYSTEMPAUSETYPE(&_ZKEVMV2.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) RATELIMITSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "RATE_LIMIT_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) RATELIMITSETTERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.RATELIMITSETTERROLE(&_ZKEVMV2.CallOpts)
}

// RATELIMITSETTERROLE is a free data retrieval call binding the contract method 0xbf3e7505.
//
// Solidity: function RATE_LIMIT_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) RATELIMITSETTERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.RATELIMITSETTERROLE(&_ZKEVMV2.CallOpts)
}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) VERIFIERSETTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "VERIFIER_SETTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) VERIFIERSETTERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.VERIFIERSETTERROLE(&_ZKEVMV2.CallOpts)
}

// VERIFIERSETTERROLE is a free data retrieval call binding the contract method 0x6e673843.
//
// Solidity: function VERIFIER_SETTER_ROLE() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) VERIFIERSETTERROLE() ([32]byte, error) {
	return _ZKEVMV2.Contract.VERIFIERSETTERROLE(&_ZKEVMV2.CallOpts)
}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentFinalizedShnarf(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentFinalizedShnarf")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentFinalizedShnarf() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentFinalizedShnarf(&_ZKEVMV2.CallOpts)
}

// CurrentFinalizedShnarf is a free data retrieval call binding the contract method 0xcd9b9e9a.
//
// Solidity: function currentFinalizedShnarf() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentFinalizedShnarf() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentFinalizedShnarf(&_ZKEVMV2.CallOpts)
}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentFinalizedState(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentFinalizedState")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentFinalizedState() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentFinalizedState(&_ZKEVMV2.CallOpts)
}

// CurrentFinalizedState is a free data retrieval call binding the contract method 0x921b278e.
//
// Solidity: function currentFinalizedState() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentFinalizedState() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentFinalizedState(&_ZKEVMV2.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentL2BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentL2BlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentL2BlockNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentL2BlockNumber(&_ZKEVMV2.CallOpts)
}

// CurrentL2BlockNumber is a free data retrieval call binding the contract method 0x695378f5.
//
// Solidity: function currentL2BlockNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentL2BlockNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentL2BlockNumber(&_ZKEVMV2.CallOpts)
}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentL2StoredL1MessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentL2StoredL1MessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentL2StoredL1MessageNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentL2StoredL1MessageNumber(&_ZKEVMV2.CallOpts)
}

// CurrentL2StoredL1MessageNumber is a free data retrieval call binding the contract method 0x05861180.
//
// Solidity: function currentL2StoredL1MessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentL2StoredL1MessageNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentL2StoredL1MessageNumber(&_ZKEVMV2.CallOpts)
}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentL2StoredL1RollingHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentL2StoredL1RollingHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentL2StoredL1RollingHash() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentL2StoredL1RollingHash(&_ZKEVMV2.CallOpts)
}

// CurrentL2StoredL1RollingHash is a free data retrieval call binding the contract method 0xd5d4b835.
//
// Solidity: function currentL2StoredL1RollingHash() view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentL2StoredL1RollingHash() ([32]byte, error) {
	return _ZKEVMV2.Contract.CurrentL2StoredL1RollingHash(&_ZKEVMV2.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentPeriodAmountInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentPeriodAmountInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentPeriodAmountInWei(&_ZKEVMV2.CallOpts)
}

// CurrentPeriodAmountInWei is a free data retrieval call binding the contract method 0xc0729ab1.
//
// Solidity: function currentPeriodAmountInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentPeriodAmountInWei() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentPeriodAmountInWei(&_ZKEVMV2.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentPeriodEnd() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentPeriodEnd(&_ZKEVMV2.CallOpts)
}

// CurrentPeriodEnd is a free data retrieval call binding the contract method 0x58794456.
//
// Solidity: function currentPeriodEnd() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentPeriodEnd() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentPeriodEnd(&_ZKEVMV2.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) CurrentTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "currentTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) CurrentTimestamp() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentTimestamp(&_ZKEVMV2.CallOpts)
}

// CurrentTimestamp is a free data retrieval call binding the contract method 0x1e2ff94f.
//
// Solidity: function currentTimestamp() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) CurrentTimestamp() (*big.Int, error) {
	return _ZKEVMV2.Contract.CurrentTimestamp(&_ZKEVMV2.CallOpts)
}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_ZKEVMV2 *ZKEVMV2Caller) DataEndingBlock(opts *bind.CallOpts, dataHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "dataEndingBlock", dataHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_ZKEVMV2 *ZKEVMV2Session) DataEndingBlock(dataHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.DataEndingBlock(&_ZKEVMV2.CallOpts, dataHash)
}

// DataEndingBlock is a free data retrieval call binding the contract method 0x5ed73ceb.
//
// Solidity: function dataEndingBlock(bytes32 dataHash) view returns(uint256 endingBlock)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DataEndingBlock(dataHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.DataEndingBlock(&_ZKEVMV2.CallOpts, dataHash)
}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_ZKEVMV2 *ZKEVMV2Caller) DataFinalStateRootHashes(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "dataFinalStateRootHashes", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_ZKEVMV2 *ZKEVMV2Session) DataFinalStateRootHashes(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataFinalStateRootHashes(&_ZKEVMV2.CallOpts, dataHash)
}

// DataFinalStateRootHashes is a free data retrieval call binding the contract method 0x6078bfd8.
//
// Solidity: function dataFinalStateRootHashes(bytes32 dataHash) view returns(bytes32 finalStateRootHash)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DataFinalStateRootHashes(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataFinalStateRootHashes(&_ZKEVMV2.CallOpts, dataHash)
}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_ZKEVMV2 *ZKEVMV2Caller) DataParents(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "dataParents", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_ZKEVMV2 *ZKEVMV2Session) DataParents(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataParents(&_ZKEVMV2.CallOpts, dataHash)
}

// DataParents is a free data retrieval call binding the contract method 0x4cdd389b.
//
// Solidity: function dataParents(bytes32 dataHash) view returns(bytes32 parentHash)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DataParents(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataParents(&_ZKEVMV2.CallOpts, dataHash)
}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_ZKEVMV2 *ZKEVMV2Caller) DataShnarfHashes(opts *bind.CallOpts, dataHash [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "dataShnarfHashes", dataHash)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_ZKEVMV2 *ZKEVMV2Session) DataShnarfHashes(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataShnarfHashes(&_ZKEVMV2.CallOpts, dataHash)
}

// DataShnarfHashes is a free data retrieval call binding the contract method 0x66f96e98.
//
// Solidity: function dataShnarfHashes(bytes32 dataHash) view returns(bytes32 shnarfHash)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DataShnarfHashes(dataHash [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.DataShnarfHashes(&_ZKEVMV2.CallOpts, dataHash)
}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_ZKEVMV2 *ZKEVMV2Caller) DataStartingBlock(opts *bind.CallOpts, dataHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "dataStartingBlock", dataHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_ZKEVMV2 *ZKEVMV2Session) DataStartingBlock(dataHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.DataStartingBlock(&_ZKEVMV2.CallOpts, dataHash)
}

// DataStartingBlock is a free data retrieval call binding the contract method 0x1f443da0.
//
// Solidity: function dataStartingBlock(bytes32 dataHash) view returns(uint256 startingBlock)
func (_ZKEVMV2 *ZKEVMV2CallerSession) DataStartingBlock(dataHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.DataStartingBlock(&_ZKEVMV2.CallOpts, dataHash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.GetRoleAdmin(&_ZKEVMV2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ZKEVMV2 *ZKEVMV2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ZKEVMV2.Contract.GetRoleAdmin(&_ZKEVMV2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEVMV2.Contract.HasRole(&_ZKEVMV2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ZKEVMV2.Contract.HasRole(&_ZKEVMV2.CallOpts, role, account)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2Caller) InboxL2L1MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "inboxL2L1MessageStatus", messageHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2Session) InboxL2L1MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.InboxL2L1MessageStatus(&_ZKEVMV2.CallOpts, messageHash)
}

// InboxL2L1MessageStatus is a free data retrieval call binding the contract method 0x5c721a0c.
//
// Solidity: function inboxL2L1MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2CallerSession) InboxL2L1MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.InboxL2L1MessageStatus(&_ZKEVMV2.CallOpts, messageHash)
}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Caller) IsMessageClaimed(opts *bind.CallOpts, _messageNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "isMessageClaimed", _messageNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Session) IsMessageClaimed(_messageNumber *big.Int) (bool, error) {
	return _ZKEVMV2.Contract.IsMessageClaimed(&_ZKEVMV2.CallOpts, _messageNumber)
}

// IsMessageClaimed is a free data retrieval call binding the contract method 0x9ee8b211.
//
// Solidity: function isMessageClaimed(uint256 _messageNumber) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2CallerSession) IsMessageClaimed(_messageNumber *big.Int) (bool, error) {
	return _ZKEVMV2.Contract.IsMessageClaimed(&_ZKEVMV2.CallOpts, _messageNumber)
}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Caller) IsPaused(opts *bind.CallOpts, _pauseType uint8) (bool, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "isPaused", _pauseType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Session) IsPaused(_pauseType uint8) (bool, error) {
	return _ZKEVMV2.Contract.IsPaused(&_ZKEVMV2.CallOpts, _pauseType)
}

// IsPaused is a free data retrieval call binding the contract method 0xbc61e733.
//
// Solidity: function isPaused(uint8 _pauseType) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2CallerSession) IsPaused(_pauseType uint8) (bool, error) {
	return _ZKEVMV2.Contract.IsPaused(&_ZKEVMV2.CallOpts, _pauseType)
}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_ZKEVMV2 *ZKEVMV2Caller) L2MerkleRootsDepths(opts *bind.CallOpts, merkleRoot [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "l2MerkleRootsDepths", merkleRoot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_ZKEVMV2 *ZKEVMV2Session) L2MerkleRootsDepths(merkleRoot [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.L2MerkleRootsDepths(&_ZKEVMV2.CallOpts, merkleRoot)
}

// L2MerkleRootsDepths is a free data retrieval call binding the contract method 0x60e83cf3.
//
// Solidity: function l2MerkleRootsDepths(bytes32 merkleRoot) view returns(uint256 treeDepth)
func (_ZKEVMV2 *ZKEVMV2CallerSession) L2MerkleRootsDepths(merkleRoot [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.L2MerkleRootsDepths(&_ZKEVMV2.CallOpts, merkleRoot)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) LimitInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "limitInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) LimitInWei() (*big.Int, error) {
	return _ZKEVMV2.Contract.LimitInWei(&_ZKEVMV2.CallOpts)
}

// LimitInWei is a free data retrieval call binding the contract method 0xad422ff0.
//
// Solidity: function limitInWei() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) LimitInWei() (*big.Int, error) {
	return _ZKEVMV2.Contract.LimitInWei(&_ZKEVMV2.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) NextMessageNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "nextMessageNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) NextMessageNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.NextMessageNumber(&_ZKEVMV2.CallOpts)
}

// NextMessageNumber is a free data retrieval call binding the contract method 0xb837dbe9.
//
// Solidity: function nextMessageNumber() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) NextMessageNumber() (*big.Int, error) {
	return _ZKEVMV2.Contract.NextMessageNumber(&_ZKEVMV2.CallOpts)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2Caller) OutboxL1L2MessageStatus(opts *bind.CallOpts, messageHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "outboxL1L2MessageStatus", messageHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2Session) OutboxL1L2MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.OutboxL1L2MessageStatus(&_ZKEVMV2.CallOpts, messageHash)
}

// OutboxL1L2MessageStatus is a free data retrieval call binding the contract method 0x3fc08b65.
//
// Solidity: function outboxL1L2MessageStatus(bytes32 messageHash) view returns(uint256 messageStatus)
func (_ZKEVMV2 *ZKEVMV2CallerSession) OutboxL1L2MessageStatus(messageHash [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.OutboxL1L2MessageStatus(&_ZKEVMV2.CallOpts, messageHash)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_ZKEVMV2 *ZKEVMV2Caller) PauseTypeStatuses(opts *bind.CallOpts, pauseType [32]byte) (bool, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "pauseTypeStatuses", pauseType)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_ZKEVMV2 *ZKEVMV2Session) PauseTypeStatuses(pauseType [32]byte) (bool, error) {
	return _ZKEVMV2.Contract.PauseTypeStatuses(&_ZKEVMV2.CallOpts, pauseType)
}

// PauseTypeStatuses is a free data retrieval call binding the contract method 0xcc5782f6.
//
// Solidity: function pauseTypeStatuses(bytes32 pauseType) view returns(bool pauseStatus)
func (_ZKEVMV2 *ZKEVMV2CallerSession) PauseTypeStatuses(pauseType [32]byte) (bool, error) {
	return _ZKEVMV2.Contract.PauseTypeStatuses(&_ZKEVMV2.CallOpts, pauseType)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) PeriodInSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "periodInSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) PeriodInSeconds() (*big.Int, error) {
	return _ZKEVMV2.Contract.PeriodInSeconds(&_ZKEVMV2.CallOpts)
}

// PeriodInSeconds is a free data retrieval call binding the contract method 0xc1dc0f07.
//
// Solidity: function periodInSeconds() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) PeriodInSeconds() (*big.Int, error) {
	return _ZKEVMV2.Contract.PeriodInSeconds(&_ZKEVMV2.CallOpts)
}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_ZKEVMV2 *ZKEVMV2Caller) RollingHashes(opts *bind.CallOpts, messageNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "rollingHashes", messageNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_ZKEVMV2 *ZKEVMV2Session) RollingHashes(messageNumber *big.Int) ([32]byte, error) {
	return _ZKEVMV2.Contract.RollingHashes(&_ZKEVMV2.CallOpts, messageNumber)
}

// RollingHashes is a free data retrieval call binding the contract method 0x914e57eb.
//
// Solidity: function rollingHashes(uint256 messageNumber) view returns(bytes32 rollingHash)
func (_ZKEVMV2 *ZKEVMV2CallerSession) RollingHashes(messageNumber *big.Int) ([32]byte, error) {
	return _ZKEVMV2.Contract.RollingHashes(&_ZKEVMV2.CallOpts, messageNumber)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address addr)
func (_ZKEVMV2 *ZKEVMV2Caller) Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address addr)
func (_ZKEVMV2 *ZKEVMV2Session) Sender() (common.Address, error) {
	return _ZKEVMV2.Contract.Sender(&_ZKEVMV2.CallOpts)
}

// Sender is a free data retrieval call binding the contract method 0x67e404ce.
//
// Solidity: function sender() view returns(address addr)
func (_ZKEVMV2 *ZKEVMV2CallerSession) Sender() (common.Address, error) {
	return _ZKEVMV2.Contract.Sender(&_ZKEVMV2.CallOpts)
}

// ShnarfFinalBlockNumbers is a free data retrieval call binding the contract method 0xf93e9857.
//
// Solidity: function shnarfFinalBlockNumbers(bytes32 shnarf) view returns(uint256 finalBlockNumber)
func (_ZKEVMV2 *ZKEVMV2Caller) ShnarfFinalBlockNumbers(opts *bind.CallOpts, shnarf [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "shnarfFinalBlockNumbers", shnarf)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShnarfFinalBlockNumbers is a free data retrieval call binding the contract method 0xf93e9857.
//
// Solidity: function shnarfFinalBlockNumbers(bytes32 shnarf) view returns(uint256 finalBlockNumber)
func (_ZKEVMV2 *ZKEVMV2Session) ShnarfFinalBlockNumbers(shnarf [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.ShnarfFinalBlockNumbers(&_ZKEVMV2.CallOpts, shnarf)
}

// ShnarfFinalBlockNumbers is a free data retrieval call binding the contract method 0xf93e9857.
//
// Solidity: function shnarfFinalBlockNumbers(bytes32 shnarf) view returns(uint256 finalBlockNumber)
func (_ZKEVMV2 *ZKEVMV2CallerSession) ShnarfFinalBlockNumbers(shnarf [32]byte) (*big.Int, error) {
	return _ZKEVMV2.Contract.ShnarfFinalBlockNumbers(&_ZKEVMV2.CallOpts, shnarf)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_ZKEVMV2 *ZKEVMV2Caller) StateRootHashes(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "stateRootHashes", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_ZKEVMV2 *ZKEVMV2Session) StateRootHashes(blockNumber *big.Int) ([32]byte, error) {
	return _ZKEVMV2.Contract.StateRootHashes(&_ZKEVMV2.CallOpts, blockNumber)
}

// StateRootHashes is a free data retrieval call binding the contract method 0x8be745d1.
//
// Solidity: function stateRootHashes(uint256 blockNumber) view returns(bytes32 stateRootHash)
func (_ZKEVMV2 *ZKEVMV2CallerSession) StateRootHashes(blockNumber *big.Int) ([32]byte, error) {
	return _ZKEVMV2.Contract.StateRootHashes(&_ZKEVMV2.CallOpts, blockNumber)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEVMV2.Contract.SupportsInterface(&_ZKEVMV2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ZKEVMV2 *ZKEVMV2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ZKEVMV2.Contract.SupportsInterface(&_ZKEVMV2.CallOpts, interfaceId)
}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Caller) SystemMigrationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "systemMigrationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2Session) SystemMigrationBlock() (*big.Int, error) {
	return _ZKEVMV2.Contract.SystemMigrationBlock(&_ZKEVMV2.CallOpts)
}

// SystemMigrationBlock is a free data retrieval call binding the contract method 0x2c70645c.
//
// Solidity: function systemMigrationBlock() view returns(uint256)
func (_ZKEVMV2 *ZKEVMV2CallerSession) SystemMigrationBlock() (*big.Int, error) {
	return _ZKEVMV2.Contract.SystemMigrationBlock(&_ZKEVMV2.CallOpts)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_ZKEVMV2 *ZKEVMV2Caller) Verifiers(opts *bind.CallOpts, proofType *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZKEVMV2.contract.Call(opts, &out, "verifiers", proofType)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_ZKEVMV2 *ZKEVMV2Session) Verifiers(proofType *big.Int) (common.Address, error) {
	return _ZKEVMV2.Contract.Verifiers(&_ZKEVMV2.CallOpts, proofType)
}

// Verifiers is a free data retrieval call binding the contract method 0xac1eff68.
//
// Solidity: function verifiers(uint256 proofType) view returns(address verifierAddress)
func (_ZKEVMV2 *ZKEVMV2CallerSession) Verifiers(proofType *big.Int) (common.Address, error) {
	return _ZKEVMV2.Contract.Verifiers(&_ZKEVMV2.CallOpts, proofType)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) ClaimMessage(opts *bind.TransactOpts, _from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "claimMessage", _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZKEVMV2 *ZKEVMV2Session) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ClaimMessage(&_ZKEVMV2.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessage is a paid mutator transaction binding the contract method 0x491e0936.
//
// Solidity: function claimMessage(address _from, address _to, uint256 _fee, uint256 _value, address _feeRecipient, bytes _calldata, uint256 _nonce) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) ClaimMessage(_from common.Address, _to common.Address, _fee *big.Int, _value *big.Int, _feeRecipient common.Address, _calldata []byte, _nonce *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ClaimMessage(&_ZKEVMV2.TransactOpts, _from, _to, _fee, _value, _feeRecipient, _calldata, _nonce)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) ClaimMessageWithProof(opts *bind.TransactOpts, _params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "claimMessageWithProof", _params)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_ZKEVMV2 *ZKEVMV2Session) ClaimMessageWithProof(_params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ClaimMessageWithProof(&_ZKEVMV2.TransactOpts, _params)
}

// ClaimMessageWithProof is a paid mutator transaction binding the contract method 0x6463fb2a.
//
// Solidity: function claimMessageWithProof((bytes32[],uint256,uint32,address,address,uint256,uint256,address,bytes32,bytes) _params) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) ClaimMessageWithProof(_params IL1MessageServiceClaimMessageWithProofParams) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ClaimMessageWithProof(&_ZKEVMV2.TransactOpts, _params)
}

// FinalizeBlocksWithProof is a paid mutator transaction binding the contract method 0xabffac32.
//
// Solidity: function finalizeBlocksWithProof(bytes _aggregatedProof, uint256 _proofType, (bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) FinalizeBlocksWithProof(opts *bind.TransactOpts, _aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "finalizeBlocksWithProof", _aggregatedProof, _proofType, _finalizationData)
}

// FinalizeBlocksWithProof is a paid mutator transaction binding the contract method 0xabffac32.
//
// Solidity: function finalizeBlocksWithProof(bytes _aggregatedProof, uint256 _proofType, (bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2Session) FinalizeBlocksWithProof(_aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.FinalizeBlocksWithProof(&_ZKEVMV2.TransactOpts, _aggregatedProof, _proofType, _finalizationData)
}

// FinalizeBlocksWithProof is a paid mutator transaction binding the contract method 0xabffac32.
//
// Solidity: function finalizeBlocksWithProof(bytes _aggregatedProof, uint256 _proofType, (bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) FinalizeBlocksWithProof(_aggregatedProof []byte, _proofType *big.Int, _finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.FinalizeBlocksWithProof(&_ZKEVMV2.TransactOpts, _aggregatedProof, _proofType, _finalizationData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x63a28393.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) FinalizeBlocksWithoutProof(opts *bind.TransactOpts, _finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "finalizeBlocksWithoutProof", _finalizationData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x63a28393.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2Session) FinalizeBlocksWithoutProof(_finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.FinalizeBlocksWithoutProof(&_ZKEVMV2.TransactOpts, _finalizationData)
}

// FinalizeBlocksWithoutProof is a paid mutator transaction binding the contract method 0x63a28393.
//
// Solidity: function finalizeBlocksWithoutProof((bytes32,bytes32,uint256,(bytes32,bytes32,bytes32,bytes32,bytes32),uint256,uint256,bytes32,bytes32,uint256,uint256,uint256,bytes32[],bytes) _finalizationData) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) FinalizeBlocksWithoutProof(_finalizationData ILineaRollupFinalizationDataV2) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.FinalizeBlocksWithoutProof(&_ZKEVMV2.TransactOpts, _finalizationData)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.GrantRole(&_ZKEVMV2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.GrantRole(&_ZKEVMV2.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x5355420e.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei, uint256 _genesisTimestamp) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) Initialize(opts *bind.TransactOpts, _initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int, _genesisTimestamp *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "initialize", _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei, _genesisTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x5355420e.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei, uint256 _genesisTimestamp) returns()
func (_ZKEVMV2 *ZKEVMV2Session) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int, _genesisTimestamp *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.Initialize(&_ZKEVMV2.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei, _genesisTimestamp)
}

// Initialize is a paid mutator transaction binding the contract method 0x5355420e.
//
// Solidity: function initialize(bytes32 _initialStateRootHash, uint256 _initialL2BlockNumber, address _defaultVerifier, address _securityCouncil, address[] _operators, uint256 _rateLimitPeriodInSeconds, uint256 _rateLimitAmountInWei, uint256 _genesisTimestamp) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) Initialize(_initialStateRootHash [32]byte, _initialL2BlockNumber *big.Int, _defaultVerifier common.Address, _securityCouncil common.Address, _operators []common.Address, _rateLimitPeriodInSeconds *big.Int, _rateLimitAmountInWei *big.Int, _genesisTimestamp *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.Initialize(&_ZKEVMV2.TransactOpts, _initialStateRootHash, _initialL2BlockNumber, _defaultVerifier, _securityCouncil, _operators, _rateLimitPeriodInSeconds, _rateLimitAmountInWei, _genesisTimestamp)
}

// InitializeParentShnarfsAndFinalizedState is a paid mutator transaction binding the contract method 0x8116d995.
//
// Solidity: function initializeParentShnarfsAndFinalizedState(bytes32[] _shnarfs, uint256[] _finalBlockNumbers) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) InitializeParentShnarfsAndFinalizedState(opts *bind.TransactOpts, _shnarfs [][32]byte, _finalBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "initializeParentShnarfsAndFinalizedState", _shnarfs, _finalBlockNumbers)
}

// InitializeParentShnarfsAndFinalizedState is a paid mutator transaction binding the contract method 0x8116d995.
//
// Solidity: function initializeParentShnarfsAndFinalizedState(bytes32[] _shnarfs, uint256[] _finalBlockNumbers) returns()
func (_ZKEVMV2 *ZKEVMV2Session) InitializeParentShnarfsAndFinalizedState(_shnarfs [][32]byte, _finalBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.InitializeParentShnarfsAndFinalizedState(&_ZKEVMV2.TransactOpts, _shnarfs, _finalBlockNumbers)
}

// InitializeParentShnarfsAndFinalizedState is a paid mutator transaction binding the contract method 0x8116d995.
//
// Solidity: function initializeParentShnarfsAndFinalizedState(bytes32[] _shnarfs, uint256[] _finalBlockNumbers) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) InitializeParentShnarfsAndFinalizedState(_shnarfs [][32]byte, _finalBlockNumbers []*big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.InitializeParentShnarfsAndFinalizedState(&_ZKEVMV2.TransactOpts, _shnarfs, _finalBlockNumbers)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) PauseByType(opts *bind.TransactOpts, _pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "pauseByType", _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2Session) PauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.PauseByType(&_ZKEVMV2.TransactOpts, _pauseType)
}

// PauseByType is a paid mutator transaction binding the contract method 0xe196fb5d.
//
// Solidity: function pauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) PauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.PauseByType(&_ZKEVMV2.TransactOpts, _pauseType)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.RenounceRole(&_ZKEVMV2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.RenounceRole(&_ZKEVMV2.TransactOpts, role, account)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) ResetAmountUsedInPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "resetAmountUsedInPeriod")
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZKEVMV2 *ZKEVMV2Session) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ResetAmountUsedInPeriod(&_ZKEVMV2.TransactOpts)
}

// ResetAmountUsedInPeriod is a paid mutator transaction binding the contract method 0xaea4f745.
//
// Solidity: function resetAmountUsedInPeriod() returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) ResetAmountUsedInPeriod() (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ResetAmountUsedInPeriod(&_ZKEVMV2.TransactOpts)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) ResetRateLimitAmount(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "resetRateLimitAmount", _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZKEVMV2 *ZKEVMV2Session) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ResetRateLimitAmount(&_ZKEVMV2.TransactOpts, _amount)
}

// ResetRateLimitAmount is a paid mutator transaction binding the contract method 0x557eac73.
//
// Solidity: function resetRateLimitAmount(uint256 _amount) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) ResetRateLimitAmount(_amount *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.ResetRateLimitAmount(&_ZKEVMV2.TransactOpts, _amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.RevokeRole(&_ZKEVMV2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.RevokeRole(&_ZKEVMV2.TransactOpts, role, account)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) SendMessage(opts *bind.TransactOpts, _to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "sendMessage", _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZKEVMV2 *ZKEVMV2Session) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SendMessage(&_ZKEVMV2.TransactOpts, _to, _fee, _calldata)
}

// SendMessage is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _to, uint256 _fee, bytes _calldata) payable returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) SendMessage(_to common.Address, _fee *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SendMessage(&_ZKEVMV2.TransactOpts, _to, _fee, _calldata)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) SetVerifierAddress(opts *bind.TransactOpts, _newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "setVerifierAddress", _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2Session) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SetVerifierAddress(&_ZKEVMV2.TransactOpts, _newVerifierAddress, _proofType)
}

// SetVerifierAddress is a paid mutator transaction binding the contract method 0xc2116974.
//
// Solidity: function setVerifierAddress(address _newVerifierAddress, uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) SetVerifierAddress(_newVerifierAddress common.Address, _proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SetVerifierAddress(&_ZKEVMV2.TransactOpts, _newVerifierAddress, _proofType)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x42fbe842.
//
// Solidity: function submitBlobs(((bytes32,uint256,uint256,bytes32),uint256,bytes,bytes)[] _blobSubmissionData, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) SubmitBlobs(opts *bind.TransactOpts, _blobSubmissionData []ILineaRollupBlobSubmissionData, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "submitBlobs", _blobSubmissionData, _parentShnarf, _finalBlobShnarf)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x42fbe842.
//
// Solidity: function submitBlobs(((bytes32,uint256,uint256,bytes32),uint256,bytes,bytes)[] _blobSubmissionData, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2Session) SubmitBlobs(_blobSubmissionData []ILineaRollupBlobSubmissionData, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SubmitBlobs(&_ZKEVMV2.TransactOpts, _blobSubmissionData, _parentShnarf, _finalBlobShnarf)
}

// SubmitBlobs is a paid mutator transaction binding the contract method 0x42fbe842.
//
// Solidity: function submitBlobs(((bytes32,uint256,uint256,bytes32),uint256,bytes,bytes)[] _blobSubmissionData, bytes32 _parentShnarf, bytes32 _finalBlobShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) SubmitBlobs(_blobSubmissionData []ILineaRollupBlobSubmissionData, _parentShnarf [32]byte, _finalBlobShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SubmitBlobs(&_ZKEVMV2.TransactOpts, _blobSubmissionData, _parentShnarf, _finalBlobShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0xd05d9c3d.
//
// Solidity: function submitDataAsCalldata((bytes32,uint256,uint256,bytes32,bytes) _submissionData, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) SubmitDataAsCalldata(opts *bind.TransactOpts, _submissionData ILineaRollupSubmissionDataV2, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "submitDataAsCalldata", _submissionData, _parentShnarf, _expectedShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0xd05d9c3d.
//
// Solidity: function submitDataAsCalldata((bytes32,uint256,uint256,bytes32,bytes) _submissionData, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2Session) SubmitDataAsCalldata(_submissionData ILineaRollupSubmissionDataV2, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SubmitDataAsCalldata(&_ZKEVMV2.TransactOpts, _submissionData, _parentShnarf, _expectedShnarf)
}

// SubmitDataAsCalldata is a paid mutator transaction binding the contract method 0xd05d9c3d.
//
// Solidity: function submitDataAsCalldata((bytes32,uint256,uint256,bytes32,bytes) _submissionData, bytes32 _parentShnarf, bytes32 _expectedShnarf) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) SubmitDataAsCalldata(_submissionData ILineaRollupSubmissionDataV2, _parentShnarf [32]byte, _expectedShnarf [32]byte) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.SubmitDataAsCalldata(&_ZKEVMV2.TransactOpts, _submissionData, _parentShnarf, _expectedShnarf)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) UnPauseByType(opts *bind.TransactOpts, _pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "unPauseByType", _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2Session) UnPauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.UnPauseByType(&_ZKEVMV2.TransactOpts, _pauseType)
}

// UnPauseByType is a paid mutator transaction binding the contract method 0x1065a399.
//
// Solidity: function unPauseByType(uint8 _pauseType) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) UnPauseByType(_pauseType uint8) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.UnPauseByType(&_ZKEVMV2.TransactOpts, _pauseType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2Transactor) UnsetVerifierAddress(opts *bind.TransactOpts, _proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.contract.Transact(opts, "unsetVerifierAddress", _proofType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2Session) UnsetVerifierAddress(_proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.UnsetVerifierAddress(&_ZKEVMV2.TransactOpts, _proofType)
}

// UnsetVerifierAddress is a paid mutator transaction binding the contract method 0x28958174.
//
// Solidity: function unsetVerifierAddress(uint256 _proofType) returns()
func (_ZKEVMV2 *ZKEVMV2TransactorSession) UnsetVerifierAddress(_proofType *big.Int) (*types.Transaction, error) {
	return _ZKEVMV2.Contract.UnsetVerifierAddress(&_ZKEVMV2.TransactOpts, _proofType)
}

// ZKEVMV2AmountUsedInPeriodResetIterator is returned from FilterAmountUsedInPeriodReset and is used to iterate over the raw logs and unpacked data for AmountUsedInPeriodReset events raised by the ZKEVMV2 contract.
type ZKEVMV2AmountUsedInPeriodResetIterator struct {
	Event *ZKEVMV2AmountUsedInPeriodReset // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2AmountUsedInPeriodResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2AmountUsedInPeriodReset)
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
		it.Event = new(ZKEVMV2AmountUsedInPeriodReset)
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
func (it *ZKEVMV2AmountUsedInPeriodResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2AmountUsedInPeriodResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2AmountUsedInPeriodReset represents a AmountUsedInPeriodReset event raised by the ZKEVMV2 contract.
type ZKEVMV2AmountUsedInPeriodReset struct {
	ResettingAddress common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAmountUsedInPeriodReset is a free log retrieval operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterAmountUsedInPeriodReset(opts *bind.FilterOpts, resettingAddress []common.Address) (*ZKEVMV2AmountUsedInPeriodResetIterator, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2AmountUsedInPeriodResetIterator{contract: _ZKEVMV2.contract, event: "AmountUsedInPeriodReset", logs: logs, sub: sub}, nil
}

// WatchAmountUsedInPeriodReset is a free log subscription operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchAmountUsedInPeriodReset(opts *bind.WatchOpts, sink chan<- *ZKEVMV2AmountUsedInPeriodReset, resettingAddress []common.Address) (event.Subscription, error) {

	var resettingAddressRule []interface{}
	for _, resettingAddressItem := range resettingAddress {
		resettingAddressRule = append(resettingAddressRule, resettingAddressItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "AmountUsedInPeriodReset", resettingAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2AmountUsedInPeriodReset)
				if err := _ZKEVMV2.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
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

// ParseAmountUsedInPeriodReset is a log parse operation binding the contract event 0xba88c025b0cbb77022c0c487beef24f759f1e4be2f51a205bc427cee19c2eaa6.
//
// Solidity: event AmountUsedInPeriodReset(address indexed resettingAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseAmountUsedInPeriodReset(log types.Log) (*ZKEVMV2AmountUsedInPeriodReset, error) {
	event := new(ZKEVMV2AmountUsedInPeriodReset)
	if err := _ZKEVMV2.contract.UnpackLog(event, "AmountUsedInPeriodReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2BlockFinalizedIterator is returned from FilterBlockFinalized and is used to iterate over the raw logs and unpacked data for BlockFinalized events raised by the ZKEVMV2 contract.
type ZKEVMV2BlockFinalizedIterator struct {
	Event *ZKEVMV2BlockFinalized // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2BlockFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2BlockFinalized)
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
		it.Event = new(ZKEVMV2BlockFinalized)
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
func (it *ZKEVMV2BlockFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2BlockFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2BlockFinalized represents a BlockFinalized event raised by the ZKEVMV2 contract.
type ZKEVMV2BlockFinalized struct {
	BlockNumber        *big.Int
	StateRootHash      [32]byte
	FinalizedWithProof bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBlockFinalized is a free log retrieval operation binding the contract event 0x047c6ce79802b16b6527cedd89156bb59f2da26867b4f218fa60c9521ddcce55.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash, bool indexed finalizedWithProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterBlockFinalized(opts *bind.FilterOpts, blockNumber []*big.Int, stateRootHash [][32]byte, finalizedWithProof []bool) (*ZKEVMV2BlockFinalizedIterator, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var stateRootHashRule []interface{}
	for _, stateRootHashItem := range stateRootHash {
		stateRootHashRule = append(stateRootHashRule, stateRootHashItem)
	}
	var finalizedWithProofRule []interface{}
	for _, finalizedWithProofItem := range finalizedWithProof {
		finalizedWithProofRule = append(finalizedWithProofRule, finalizedWithProofItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "BlockFinalized", blockNumberRule, stateRootHashRule, finalizedWithProofRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2BlockFinalizedIterator{contract: _ZKEVMV2.contract, event: "BlockFinalized", logs: logs, sub: sub}, nil
}

// WatchBlockFinalized is a free log subscription operation binding the contract event 0x047c6ce79802b16b6527cedd89156bb59f2da26867b4f218fa60c9521ddcce55.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash, bool indexed finalizedWithProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchBlockFinalized(opts *bind.WatchOpts, sink chan<- *ZKEVMV2BlockFinalized, blockNumber []*big.Int, stateRootHash [][32]byte, finalizedWithProof []bool) (event.Subscription, error) {

	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}
	var stateRootHashRule []interface{}
	for _, stateRootHashItem := range stateRootHash {
		stateRootHashRule = append(stateRootHashRule, stateRootHashItem)
	}
	var finalizedWithProofRule []interface{}
	for _, finalizedWithProofItem := range finalizedWithProof {
		finalizedWithProofRule = append(finalizedWithProofRule, finalizedWithProofItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "BlockFinalized", blockNumberRule, stateRootHashRule, finalizedWithProofRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2BlockFinalized)
				if err := _ZKEVMV2.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
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

// ParseBlockFinalized is a log parse operation binding the contract event 0x047c6ce79802b16b6527cedd89156bb59f2da26867b4f218fa60c9521ddcce55.
//
// Solidity: event BlockFinalized(uint256 indexed blockNumber, bytes32 indexed stateRootHash, bool indexed finalizedWithProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseBlockFinalized(log types.Log) (*ZKEVMV2BlockFinalized, error) {
	event := new(ZKEVMV2BlockFinalized)
	if err := _ZKEVMV2.contract.UnpackLog(event, "BlockFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2BlocksVerificationDoneIterator is returned from FilterBlocksVerificationDone and is used to iterate over the raw logs and unpacked data for BlocksVerificationDone events raised by the ZKEVMV2 contract.
type ZKEVMV2BlocksVerificationDoneIterator struct {
	Event *ZKEVMV2BlocksVerificationDone // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2BlocksVerificationDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2BlocksVerificationDone)
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
		it.Event = new(ZKEVMV2BlocksVerificationDone)
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
func (it *ZKEVMV2BlocksVerificationDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2BlocksVerificationDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2BlocksVerificationDone represents a BlocksVerificationDone event raised by the ZKEVMV2 contract.
type ZKEVMV2BlocksVerificationDone struct {
	LastBlockFinalized *big.Int
	StartingRootHash   [32]byte
	FinalRootHash      [32]byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBlocksVerificationDone is a free log retrieval operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterBlocksVerificationDone(opts *bind.FilterOpts, lastBlockFinalized []*big.Int) (*ZKEVMV2BlocksVerificationDoneIterator, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2BlocksVerificationDoneIterator{contract: _ZKEVMV2.contract, event: "BlocksVerificationDone", logs: logs, sub: sub}, nil
}

// WatchBlocksVerificationDone is a free log subscription operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchBlocksVerificationDone(opts *bind.WatchOpts, sink chan<- *ZKEVMV2BlocksVerificationDone, lastBlockFinalized []*big.Int) (event.Subscription, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "BlocksVerificationDone", lastBlockFinalizedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2BlocksVerificationDone)
				if err := _ZKEVMV2.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
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

// ParseBlocksVerificationDone is a log parse operation binding the contract event 0x5c885a794662ebe3b08ae0874fc2c88b5343b0223ba9cd2cad92b69c0d0c901f.
//
// Solidity: event BlocksVerificationDone(uint256 indexed lastBlockFinalized, bytes32 startingRootHash, bytes32 finalRootHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseBlocksVerificationDone(log types.Log) (*ZKEVMV2BlocksVerificationDone, error) {
	event := new(ZKEVMV2BlocksVerificationDone)
	if err := _ZKEVMV2.contract.UnpackLog(event, "BlocksVerificationDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2DataFinalizedIterator is returned from FilterDataFinalized and is used to iterate over the raw logs and unpacked data for DataFinalized events raised by the ZKEVMV2 contract.
type ZKEVMV2DataFinalizedIterator struct {
	Event *ZKEVMV2DataFinalized // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2DataFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2DataFinalized)
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
		it.Event = new(ZKEVMV2DataFinalized)
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
func (it *ZKEVMV2DataFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2DataFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2DataFinalized represents a DataFinalized event raised by the ZKEVMV2 contract.
type ZKEVMV2DataFinalized struct {
	LastBlockFinalized *big.Int
	StartingRootHash   [32]byte
	FinalRootHash      [32]byte
	WithProof          bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDataFinalized is a free log retrieval operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterDataFinalized(opts *bind.FilterOpts, lastBlockFinalized []*big.Int, startingRootHash [][32]byte, finalRootHash [][32]byte) (*ZKEVMV2DataFinalizedIterator, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}
	var startingRootHashRule []interface{}
	for _, startingRootHashItem := range startingRootHash {
		startingRootHashRule = append(startingRootHashRule, startingRootHashItem)
	}
	var finalRootHashRule []interface{}
	for _, finalRootHashItem := range finalRootHash {
		finalRootHashRule = append(finalRootHashRule, finalRootHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "DataFinalized", lastBlockFinalizedRule, startingRootHashRule, finalRootHashRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2DataFinalizedIterator{contract: _ZKEVMV2.contract, event: "DataFinalized", logs: logs, sub: sub}, nil
}

// WatchDataFinalized is a free log subscription operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchDataFinalized(opts *bind.WatchOpts, sink chan<- *ZKEVMV2DataFinalized, lastBlockFinalized []*big.Int, startingRootHash [][32]byte, finalRootHash [][32]byte) (event.Subscription, error) {

	var lastBlockFinalizedRule []interface{}
	for _, lastBlockFinalizedItem := range lastBlockFinalized {
		lastBlockFinalizedRule = append(lastBlockFinalizedRule, lastBlockFinalizedItem)
	}
	var startingRootHashRule []interface{}
	for _, startingRootHashItem := range startingRootHash {
		startingRootHashRule = append(startingRootHashRule, startingRootHashItem)
	}
	var finalRootHashRule []interface{}
	for _, finalRootHashItem := range finalRootHash {
		finalRootHashRule = append(finalRootHashRule, finalRootHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "DataFinalized", lastBlockFinalizedRule, startingRootHashRule, finalRootHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2DataFinalized)
				if err := _ZKEVMV2.contract.UnpackLog(event, "DataFinalized", log); err != nil {
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

// ParseDataFinalized is a log parse operation binding the contract event 0x1335f1a2b3ff25f07f5fef07dd35d8fb4312c3c73b138e2fad9347b3319ab53c.
//
// Solidity: event DataFinalized(uint256 indexed lastBlockFinalized, bytes32 indexed startingRootHash, bytes32 indexed finalRootHash, bool withProof)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseDataFinalized(log types.Log) (*ZKEVMV2DataFinalized, error) {
	event := new(ZKEVMV2DataFinalized)
	if err := _ZKEVMV2.contract.UnpackLog(event, "DataFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2DataSubmittedV2Iterator is returned from FilterDataSubmittedV2 and is used to iterate over the raw logs and unpacked data for DataSubmittedV2 events raised by the ZKEVMV2 contract.
type ZKEVMV2DataSubmittedV2Iterator struct {
	Event *ZKEVMV2DataSubmittedV2 // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2DataSubmittedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2DataSubmittedV2)
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
		it.Event = new(ZKEVMV2DataSubmittedV2)
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
func (it *ZKEVMV2DataSubmittedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2DataSubmittedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2DataSubmittedV2 represents a DataSubmittedV2 event raised by the ZKEVMV2 contract.
type ZKEVMV2DataSubmittedV2 struct {
	Shnarf     [32]byte
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataSubmittedV2 is a free log retrieval operation binding the contract event 0x52bd912b0041a6ec8554b0ffae3cc1cb7b137bade0dc9e3f8f0b6758466d5df9.
//
// Solidity: event DataSubmittedV2(bytes32 indexed shnarf, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterDataSubmittedV2(opts *bind.FilterOpts, shnarf [][32]byte, startBlock []*big.Int, endBlock []*big.Int) (*ZKEVMV2DataSubmittedV2Iterator, error) {

	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "DataSubmittedV2", shnarfRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2DataSubmittedV2Iterator{contract: _ZKEVMV2.contract, event: "DataSubmittedV2", logs: logs, sub: sub}, nil
}

// WatchDataSubmittedV2 is a free log subscription operation binding the contract event 0x52bd912b0041a6ec8554b0ffae3cc1cb7b137bade0dc9e3f8f0b6758466d5df9.
//
// Solidity: event DataSubmittedV2(bytes32 indexed shnarf, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchDataSubmittedV2(opts *bind.WatchOpts, sink chan<- *ZKEVMV2DataSubmittedV2, shnarf [][32]byte, startBlock []*big.Int, endBlock []*big.Int) (event.Subscription, error) {

	var shnarfRule []interface{}
	for _, shnarfItem := range shnarf {
		shnarfRule = append(shnarfRule, shnarfItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "DataSubmittedV2", shnarfRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2DataSubmittedV2)
				if err := _ZKEVMV2.contract.UnpackLog(event, "DataSubmittedV2", log); err != nil {
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

// ParseDataSubmittedV2 is a log parse operation binding the contract event 0x52bd912b0041a6ec8554b0ffae3cc1cb7b137bade0dc9e3f8f0b6758466d5df9.
//
// Solidity: event DataSubmittedV2(bytes32 indexed shnarf, uint256 indexed startBlock, uint256 indexed endBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseDataSubmittedV2(log types.Log) (*ZKEVMV2DataSubmittedV2, error) {
	event := new(ZKEVMV2DataSubmittedV2)
	if err := _ZKEVMV2.contract.UnpackLog(event, "DataSubmittedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ZKEVMV2 contract.
type ZKEVMV2InitializedIterator struct {
	Event *ZKEVMV2Initialized // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2Initialized)
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
		it.Event = new(ZKEVMV2Initialized)
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
func (it *ZKEVMV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2Initialized represents a Initialized event raised by the ZKEVMV2 contract.
type ZKEVMV2Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*ZKEVMV2InitializedIterator, error) {

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2InitializedIterator{contract: _ZKEVMV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ZKEVMV2Initialized) (event.Subscription, error) {

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2Initialized)
				if err := _ZKEVMV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseInitialized(log types.Log) (*ZKEVMV2Initialized, error) {
	event := new(ZKEVMV2Initialized)
	if err := _ZKEVMV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2L1L2MessagesReceivedOnL2Iterator is returned from FilterL1L2MessagesReceivedOnL2 and is used to iterate over the raw logs and unpacked data for L1L2MessagesReceivedOnL2 events raised by the ZKEVMV2 contract.
type ZKEVMV2L1L2MessagesReceivedOnL2Iterator struct {
	Event *ZKEVMV2L1L2MessagesReceivedOnL2 // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2L1L2MessagesReceivedOnL2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2L1L2MessagesReceivedOnL2)
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
		it.Event = new(ZKEVMV2L1L2MessagesReceivedOnL2)
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
func (it *ZKEVMV2L1L2MessagesReceivedOnL2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2L1L2MessagesReceivedOnL2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2L1L2MessagesReceivedOnL2 represents a L1L2MessagesReceivedOnL2 event raised by the ZKEVMV2 contract.
type ZKEVMV2L1L2MessagesReceivedOnL2 struct {
	MessageHashes [][32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterL1L2MessagesReceivedOnL2 is a free log retrieval operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterL1L2MessagesReceivedOnL2(opts *bind.FilterOpts) (*ZKEVMV2L1L2MessagesReceivedOnL2Iterator, error) {

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2L1L2MessagesReceivedOnL2Iterator{contract: _ZKEVMV2.contract, event: "L1L2MessagesReceivedOnL2", logs: logs, sub: sub}, nil
}

// WatchL1L2MessagesReceivedOnL2 is a free log subscription operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchL1L2MessagesReceivedOnL2(opts *bind.WatchOpts, sink chan<- *ZKEVMV2L1L2MessagesReceivedOnL2) (event.Subscription, error) {

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "L1L2MessagesReceivedOnL2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2L1L2MessagesReceivedOnL2)
				if err := _ZKEVMV2.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
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

// ParseL1L2MessagesReceivedOnL2 is a log parse operation binding the contract event 0x95e84bb4317676921a29fd1d13f8f0153508473b899c12b3cd08314348801d64.
//
// Solidity: event L1L2MessagesReceivedOnL2(bytes32[] messageHashes)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseL1L2MessagesReceivedOnL2(log types.Log) (*ZKEVMV2L1L2MessagesReceivedOnL2, error) {
	event := new(ZKEVMV2L1L2MessagesReceivedOnL2)
	if err := _ZKEVMV2.contract.UnpackLog(event, "L1L2MessagesReceivedOnL2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2L2L1MessageHashAddedToInboxIterator is returned from FilterL2L1MessageHashAddedToInbox and is used to iterate over the raw logs and unpacked data for L2L1MessageHashAddedToInbox events raised by the ZKEVMV2 contract.
type ZKEVMV2L2L1MessageHashAddedToInboxIterator struct {
	Event *ZKEVMV2L2L1MessageHashAddedToInbox // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2L2L1MessageHashAddedToInboxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2L2L1MessageHashAddedToInbox)
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
		it.Event = new(ZKEVMV2L2L1MessageHashAddedToInbox)
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
func (it *ZKEVMV2L2L1MessageHashAddedToInboxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2L2L1MessageHashAddedToInboxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2L2L1MessageHashAddedToInbox represents a L2L1MessageHashAddedToInbox event raised by the ZKEVMV2 contract.
type ZKEVMV2L2L1MessageHashAddedToInbox struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterL2L1MessageHashAddedToInbox is a free log retrieval operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterL2L1MessageHashAddedToInbox(opts *bind.FilterOpts, messageHash [][32]byte) (*ZKEVMV2L2L1MessageHashAddedToInboxIterator, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2L2L1MessageHashAddedToInboxIterator{contract: _ZKEVMV2.contract, event: "L2L1MessageHashAddedToInbox", logs: logs, sub: sub}, nil
}

// WatchL2L1MessageHashAddedToInbox is a free log subscription operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchL2L1MessageHashAddedToInbox(opts *bind.WatchOpts, sink chan<- *ZKEVMV2L2L1MessageHashAddedToInbox, messageHash [][32]byte) (event.Subscription, error) {

	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "L2L1MessageHashAddedToInbox", messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2L2L1MessageHashAddedToInbox)
				if err := _ZKEVMV2.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
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

// ParseL2L1MessageHashAddedToInbox is a log parse operation binding the contract event 0x810484e22f73d8f099aaee1edb851ec6be6d84d43045d0a7803e5f7b3612edce.
//
// Solidity: event L2L1MessageHashAddedToInbox(bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseL2L1MessageHashAddedToInbox(log types.Log) (*ZKEVMV2L2L1MessageHashAddedToInbox, error) {
	event := new(ZKEVMV2L2L1MessageHashAddedToInbox)
	if err := _ZKEVMV2.contract.UnpackLog(event, "L2L1MessageHashAddedToInbox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2L2MerkleRootAddedIterator is returned from FilterL2MerkleRootAdded and is used to iterate over the raw logs and unpacked data for L2MerkleRootAdded events raised by the ZKEVMV2 contract.
type ZKEVMV2L2MerkleRootAddedIterator struct {
	Event *ZKEVMV2L2MerkleRootAdded // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2L2MerkleRootAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2L2MerkleRootAdded)
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
		it.Event = new(ZKEVMV2L2MerkleRootAdded)
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
func (it *ZKEVMV2L2MerkleRootAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2L2MerkleRootAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2L2MerkleRootAdded represents a L2MerkleRootAdded event raised by the ZKEVMV2 contract.
type ZKEVMV2L2MerkleRootAdded struct {
	L2MerkleRoot [32]byte
	TreeDepth    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterL2MerkleRootAdded is a free log retrieval operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterL2MerkleRootAdded(opts *bind.FilterOpts, l2MerkleRoot [][32]byte, treeDepth []*big.Int) (*ZKEVMV2L2MerkleRootAddedIterator, error) {

	var l2MerkleRootRule []interface{}
	for _, l2MerkleRootItem := range l2MerkleRoot {
		l2MerkleRootRule = append(l2MerkleRootRule, l2MerkleRootItem)
	}
	var treeDepthRule []interface{}
	for _, treeDepthItem := range treeDepth {
		treeDepthRule = append(treeDepthRule, treeDepthItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "L2MerkleRootAdded", l2MerkleRootRule, treeDepthRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2L2MerkleRootAddedIterator{contract: _ZKEVMV2.contract, event: "L2MerkleRootAdded", logs: logs, sub: sub}, nil
}

// WatchL2MerkleRootAdded is a free log subscription operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchL2MerkleRootAdded(opts *bind.WatchOpts, sink chan<- *ZKEVMV2L2MerkleRootAdded, l2MerkleRoot [][32]byte, treeDepth []*big.Int) (event.Subscription, error) {

	var l2MerkleRootRule []interface{}
	for _, l2MerkleRootItem := range l2MerkleRoot {
		l2MerkleRootRule = append(l2MerkleRootRule, l2MerkleRootItem)
	}
	var treeDepthRule []interface{}
	for _, treeDepthItem := range treeDepth {
		treeDepthRule = append(treeDepthRule, treeDepthItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "L2MerkleRootAdded", l2MerkleRootRule, treeDepthRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2L2MerkleRootAdded)
				if err := _ZKEVMV2.contract.UnpackLog(event, "L2MerkleRootAdded", log); err != nil {
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

// ParseL2MerkleRootAdded is a log parse operation binding the contract event 0x300e6f978eee6a4b0bba78dd8400dc64fd5652dbfc868a2258e16d0977be222b.
//
// Solidity: event L2MerkleRootAdded(bytes32 indexed l2MerkleRoot, uint256 indexed treeDepth)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseL2MerkleRootAdded(log types.Log) (*ZKEVMV2L2MerkleRootAdded, error) {
	event := new(ZKEVMV2L2MerkleRootAdded)
	if err := _ZKEVMV2.contract.UnpackLog(event, "L2MerkleRootAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2L2MessagingBlockAnchoredIterator is returned from FilterL2MessagingBlockAnchored and is used to iterate over the raw logs and unpacked data for L2MessagingBlockAnchored events raised by the ZKEVMV2 contract.
type ZKEVMV2L2MessagingBlockAnchoredIterator struct {
	Event *ZKEVMV2L2MessagingBlockAnchored // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2L2MessagingBlockAnchoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2L2MessagingBlockAnchored)
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
		it.Event = new(ZKEVMV2L2MessagingBlockAnchored)
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
func (it *ZKEVMV2L2MessagingBlockAnchoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2L2MessagingBlockAnchoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2L2MessagingBlockAnchored represents a L2MessagingBlockAnchored event raised by the ZKEVMV2 contract.
type ZKEVMV2L2MessagingBlockAnchored struct {
	L2Block *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterL2MessagingBlockAnchored is a free log retrieval operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterL2MessagingBlockAnchored(opts *bind.FilterOpts, l2Block []*big.Int) (*ZKEVMV2L2MessagingBlockAnchoredIterator, error) {

	var l2BlockRule []interface{}
	for _, l2BlockItem := range l2Block {
		l2BlockRule = append(l2BlockRule, l2BlockItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "L2MessagingBlockAnchored", l2BlockRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2L2MessagingBlockAnchoredIterator{contract: _ZKEVMV2.contract, event: "L2MessagingBlockAnchored", logs: logs, sub: sub}, nil
}

// WatchL2MessagingBlockAnchored is a free log subscription operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchL2MessagingBlockAnchored(opts *bind.WatchOpts, sink chan<- *ZKEVMV2L2MessagingBlockAnchored, l2Block []*big.Int) (event.Subscription, error) {

	var l2BlockRule []interface{}
	for _, l2BlockItem := range l2Block {
		l2BlockRule = append(l2BlockRule, l2BlockItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "L2MessagingBlockAnchored", l2BlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2L2MessagingBlockAnchored)
				if err := _ZKEVMV2.contract.UnpackLog(event, "L2MessagingBlockAnchored", log); err != nil {
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

// ParseL2MessagingBlockAnchored is a log parse operation binding the contract event 0x3c116827db9db3a30c1a25db8b0ee4bab9d2b223560209cfd839601b621c726d.
//
// Solidity: event L2MessagingBlockAnchored(uint256 indexed l2Block)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseL2MessagingBlockAnchored(log types.Log) (*ZKEVMV2L2MessagingBlockAnchored, error) {
	event := new(ZKEVMV2L2MessagingBlockAnchored)
	if err := _ZKEVMV2.contract.UnpackLog(event, "L2MessagingBlockAnchored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2LimitAmountChangedIterator is returned from FilterLimitAmountChanged and is used to iterate over the raw logs and unpacked data for LimitAmountChanged events raised by the ZKEVMV2 contract.
type ZKEVMV2LimitAmountChangedIterator struct {
	Event *ZKEVMV2LimitAmountChanged // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2LimitAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2LimitAmountChanged)
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
		it.Event = new(ZKEVMV2LimitAmountChanged)
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
func (it *ZKEVMV2LimitAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2LimitAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2LimitAmountChanged represents a LimitAmountChanged event raised by the ZKEVMV2 contract.
type ZKEVMV2LimitAmountChanged struct {
	AmountChangeBy           common.Address
	Amount                   *big.Int
	AmountUsedLoweredToLimit bool
	UsedAmountResetToZero    bool
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterLimitAmountChanged is a free log retrieval operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterLimitAmountChanged(opts *bind.FilterOpts, amountChangeBy []common.Address) (*ZKEVMV2LimitAmountChangedIterator, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2LimitAmountChangedIterator{contract: _ZKEVMV2.contract, event: "LimitAmountChanged", logs: logs, sub: sub}, nil
}

// WatchLimitAmountChanged is a free log subscription operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchLimitAmountChanged(opts *bind.WatchOpts, sink chan<- *ZKEVMV2LimitAmountChanged, amountChangeBy []common.Address) (event.Subscription, error) {

	var amountChangeByRule []interface{}
	for _, amountChangeByItem := range amountChangeBy {
		amountChangeByRule = append(amountChangeByRule, amountChangeByItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "LimitAmountChanged", amountChangeByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2LimitAmountChanged)
				if err := _ZKEVMV2.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
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

// ParseLimitAmountChanged is a log parse operation binding the contract event 0xbc3dc0cb5c15c51c81316450d44048838bb478b9809447d01c766a06f3e9f2c8.
//
// Solidity: event LimitAmountChanged(address indexed amountChangeBy, uint256 amount, bool amountUsedLoweredToLimit, bool usedAmountResetToZero)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseLimitAmountChanged(log types.Log) (*ZKEVMV2LimitAmountChanged, error) {
	event := new(ZKEVMV2LimitAmountChanged)
	if err := _ZKEVMV2.contract.UnpackLog(event, "LimitAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2MessageClaimedIterator is returned from FilterMessageClaimed and is used to iterate over the raw logs and unpacked data for MessageClaimed events raised by the ZKEVMV2 contract.
type ZKEVMV2MessageClaimedIterator struct {
	Event *ZKEVMV2MessageClaimed // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2MessageClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2MessageClaimed)
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
		it.Event = new(ZKEVMV2MessageClaimed)
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
func (it *ZKEVMV2MessageClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2MessageClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2MessageClaimed represents a MessageClaimed event raised by the ZKEVMV2 contract.
type ZKEVMV2MessageClaimed struct {
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageClaimed is a free log retrieval operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterMessageClaimed(opts *bind.FilterOpts, _messageHash [][32]byte) (*ZKEVMV2MessageClaimedIterator, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2MessageClaimedIterator{contract: _ZKEVMV2.contract, event: "MessageClaimed", logs: logs, sub: sub}, nil
}

// WatchMessageClaimed is a free log subscription operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchMessageClaimed(opts *bind.WatchOpts, sink chan<- *ZKEVMV2MessageClaimed, _messageHash [][32]byte) (event.Subscription, error) {

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "MessageClaimed", _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2MessageClaimed)
				if err := _ZKEVMV2.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
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

// ParseMessageClaimed is a log parse operation binding the contract event 0xa4c827e719e911e8f19393ccdb85b5102f08f0910604d340ba38390b7ff2ab0e.
//
// Solidity: event MessageClaimed(bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseMessageClaimed(log types.Log) (*ZKEVMV2MessageClaimed, error) {
	event := new(ZKEVMV2MessageClaimed)
	if err := _ZKEVMV2.contract.UnpackLog(event, "MessageClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2MessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the ZKEVMV2 contract.
type ZKEVMV2MessageSentIterator struct {
	Event *ZKEVMV2MessageSent // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2MessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2MessageSent)
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
		it.Event = new(ZKEVMV2MessageSent)
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
func (it *ZKEVMV2MessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2MessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2MessageSent represents a MessageSent event raised by the ZKEVMV2 contract.
type ZKEVMV2MessageSent struct {
	From        common.Address
	To          common.Address
	Fee         *big.Int
	Value       *big.Int
	Nonce       *big.Int
	Calldata    []byte
	MessageHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterMessageSent(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (*ZKEVMV2MessageSentIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2MessageSentIterator{contract: _ZKEVMV2.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *ZKEVMV2MessageSent, _from []common.Address, _to []common.Address, _messageHash [][32]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	var _messageHashRule []interface{}
	for _, _messageHashItem := range _messageHash {
		_messageHashRule = append(_messageHashRule, _messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "MessageSent", _fromRule, _toRule, _messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2MessageSent)
				if err := _ZKEVMV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0xe856c2b8bd4eb0027ce32eeaf595c21b0b6b4644b326e5b7bd80a1cf8db72e6c.
//
// Solidity: event MessageSent(address indexed _from, address indexed _to, uint256 _fee, uint256 _value, uint256 _nonce, bytes _calldata, bytes32 indexed _messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseMessageSent(log types.Log) (*ZKEVMV2MessageSent, error) {
	event := new(ZKEVMV2MessageSent)
	if err := _ZKEVMV2.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ZKEVMV2 contract.
type ZKEVMV2PausedIterator struct {
	Event *ZKEVMV2Paused // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2Paused)
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
		it.Event = new(ZKEVMV2Paused)
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
func (it *ZKEVMV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2Paused represents a Paused event raised by the ZKEVMV2 contract.
type ZKEVMV2Paused struct {
	MessageSender common.Address
	PauseType     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterPaused(opts *bind.FilterOpts, pauseType []*big.Int) (*ZKEVMV2PausedIterator, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "Paused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2PausedIterator{contract: _ZKEVMV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ZKEVMV2Paused, pauseType []*big.Int) (event.Subscription, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "Paused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2Paused)
				if err := _ZKEVMV2.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0xab40a374bc51de372200a8bc981af8c9ecdc08dfdaef0bb6e09f88f3c616ef3d.
//
// Solidity: event Paused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParsePaused(log types.Log) (*ZKEVMV2Paused, error) {
	event := new(ZKEVMV2Paused)
	if err := _ZKEVMV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2RateLimitInitializedIterator is returned from FilterRateLimitInitialized and is used to iterate over the raw logs and unpacked data for RateLimitInitialized events raised by the ZKEVMV2 contract.
type ZKEVMV2RateLimitInitializedIterator struct {
	Event *ZKEVMV2RateLimitInitialized // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2RateLimitInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2RateLimitInitialized)
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
		it.Event = new(ZKEVMV2RateLimitInitialized)
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
func (it *ZKEVMV2RateLimitInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2RateLimitInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2RateLimitInitialized represents a RateLimitInitialized event raised by the ZKEVMV2 contract.
type ZKEVMV2RateLimitInitialized struct {
	PeriodInSeconds  *big.Int
	LimitInWei       *big.Int
	CurrentPeriodEnd *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRateLimitInitialized is a free log retrieval operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterRateLimitInitialized(opts *bind.FilterOpts) (*ZKEVMV2RateLimitInitializedIterator, error) {

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "RateLimitInitialized")
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2RateLimitInitializedIterator{contract: _ZKEVMV2.contract, event: "RateLimitInitialized", logs: logs, sub: sub}, nil
}

// WatchRateLimitInitialized is a free log subscription operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchRateLimitInitialized(opts *bind.WatchOpts, sink chan<- *ZKEVMV2RateLimitInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "RateLimitInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2RateLimitInitialized)
				if err := _ZKEVMV2.contract.UnpackLog(event, "RateLimitInitialized", log); err != nil {
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

// ParseRateLimitInitialized is a log parse operation binding the contract event 0x8f805c372b66240792580418b7328c0c554ae235f0932475c51b026887fe26a9.
//
// Solidity: event RateLimitInitialized(uint256 periodInSeconds, uint256 limitInWei, uint256 currentPeriodEnd)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseRateLimitInitialized(log types.Log) (*ZKEVMV2RateLimitInitialized, error) {
	event := new(ZKEVMV2RateLimitInitialized)
	if err := _ZKEVMV2.contract.UnpackLog(event, "RateLimitInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ZKEVMV2 contract.
type ZKEVMV2RoleAdminChangedIterator struct {
	Event *ZKEVMV2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2RoleAdminChanged)
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
		it.Event = new(ZKEVMV2RoleAdminChanged)
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
func (it *ZKEVMV2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2RoleAdminChanged represents a RoleAdminChanged event raised by the ZKEVMV2 contract.
type ZKEVMV2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ZKEVMV2RoleAdminChangedIterator, error) {

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

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2RoleAdminChangedIterator{contract: _ZKEVMV2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ZKEVMV2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2RoleAdminChanged)
				if err := _ZKEVMV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseRoleAdminChanged(log types.Log) (*ZKEVMV2RoleAdminChanged, error) {
	event := new(ZKEVMV2RoleAdminChanged)
	if err := _ZKEVMV2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ZKEVMV2 contract.
type ZKEVMV2RoleGrantedIterator struct {
	Event *ZKEVMV2RoleGranted // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2RoleGranted)
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
		it.Event = new(ZKEVMV2RoleGranted)
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
func (it *ZKEVMV2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2RoleGranted represents a RoleGranted event raised by the ZKEVMV2 contract.
type ZKEVMV2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEVMV2RoleGrantedIterator, error) {

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

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2RoleGrantedIterator{contract: _ZKEVMV2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ZKEVMV2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2RoleGranted)
				if err := _ZKEVMV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseRoleGranted(log types.Log) (*ZKEVMV2RoleGranted, error) {
	event := new(ZKEVMV2RoleGranted)
	if err := _ZKEVMV2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ZKEVMV2 contract.
type ZKEVMV2RoleRevokedIterator struct {
	Event *ZKEVMV2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2RoleRevoked)
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
		it.Event = new(ZKEVMV2RoleRevoked)
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
func (it *ZKEVMV2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2RoleRevoked represents a RoleRevoked event raised by the ZKEVMV2 contract.
type ZKEVMV2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ZKEVMV2RoleRevokedIterator, error) {

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

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2RoleRevokedIterator{contract: _ZKEVMV2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ZKEVMV2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2RoleRevoked)
				if err := _ZKEVMV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseRoleRevoked(log types.Log) (*ZKEVMV2RoleRevoked, error) {
	event := new(ZKEVMV2RoleRevoked)
	if err := _ZKEVMV2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2RollingHashUpdatedIterator is returned from FilterRollingHashUpdated and is used to iterate over the raw logs and unpacked data for RollingHashUpdated events raised by the ZKEVMV2 contract.
type ZKEVMV2RollingHashUpdatedIterator struct {
	Event *ZKEVMV2RollingHashUpdated // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2RollingHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2RollingHashUpdated)
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
		it.Event = new(ZKEVMV2RollingHashUpdated)
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
func (it *ZKEVMV2RollingHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2RollingHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2RollingHashUpdated represents a RollingHashUpdated event raised by the ZKEVMV2 contract.
type ZKEVMV2RollingHashUpdated struct {
	MessageNumber *big.Int
	RollingHash   [32]byte
	MessageHash   [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollingHashUpdated is a free log retrieval operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterRollingHashUpdated(opts *bind.FilterOpts, messageNumber []*big.Int, rollingHash [][32]byte, messageHash [][32]byte) (*ZKEVMV2RollingHashUpdatedIterator, error) {

	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var rollingHashRule []interface{}
	for _, rollingHashItem := range rollingHash {
		rollingHashRule = append(rollingHashRule, rollingHashItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "RollingHashUpdated", messageNumberRule, rollingHashRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2RollingHashUpdatedIterator{contract: _ZKEVMV2.contract, event: "RollingHashUpdated", logs: logs, sub: sub}, nil
}

// WatchRollingHashUpdated is a free log subscription operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchRollingHashUpdated(opts *bind.WatchOpts, sink chan<- *ZKEVMV2RollingHashUpdated, messageNumber []*big.Int, rollingHash [][32]byte, messageHash [][32]byte) (event.Subscription, error) {

	var messageNumberRule []interface{}
	for _, messageNumberItem := range messageNumber {
		messageNumberRule = append(messageNumberRule, messageNumberItem)
	}
	var rollingHashRule []interface{}
	for _, rollingHashItem := range rollingHash {
		rollingHashRule = append(rollingHashRule, rollingHashItem)
	}
	var messageHashRule []interface{}
	for _, messageHashItem := range messageHash {
		messageHashRule = append(messageHashRule, messageHashItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "RollingHashUpdated", messageNumberRule, rollingHashRule, messageHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2RollingHashUpdated)
				if err := _ZKEVMV2.contract.UnpackLog(event, "RollingHashUpdated", log); err != nil {
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

// ParseRollingHashUpdated is a log parse operation binding the contract event 0xea3b023b4c8680d4b4824f0143132c95476359a2bb70a81d6c5a36f6918f6339.
//
// Solidity: event RollingHashUpdated(uint256 indexed messageNumber, bytes32 indexed rollingHash, bytes32 indexed messageHash)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseRollingHashUpdated(log types.Log) (*ZKEVMV2RollingHashUpdated, error) {
	event := new(ZKEVMV2RollingHashUpdated)
	if err := _ZKEVMV2.contract.UnpackLog(event, "RollingHashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2SystemMigrationBlockInitializedIterator is returned from FilterSystemMigrationBlockInitialized and is used to iterate over the raw logs and unpacked data for SystemMigrationBlockInitialized events raised by the ZKEVMV2 contract.
type ZKEVMV2SystemMigrationBlockInitializedIterator struct {
	Event *ZKEVMV2SystemMigrationBlockInitialized // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2SystemMigrationBlockInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2SystemMigrationBlockInitialized)
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
		it.Event = new(ZKEVMV2SystemMigrationBlockInitialized)
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
func (it *ZKEVMV2SystemMigrationBlockInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2SystemMigrationBlockInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2SystemMigrationBlockInitialized represents a SystemMigrationBlockInitialized event raised by the ZKEVMV2 contract.
type ZKEVMV2SystemMigrationBlockInitialized struct {
	SystemMigrationBlock *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSystemMigrationBlockInitialized is a free log retrieval operation binding the contract event 0x405b3b16b9190c1e995514c13ab4e8e7d895d9103e91c3a8c8f12df6cd50aa2c.
//
// Solidity: event SystemMigrationBlockInitialized(uint256 systemMigrationBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterSystemMigrationBlockInitialized(opts *bind.FilterOpts) (*ZKEVMV2SystemMigrationBlockInitializedIterator, error) {

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "SystemMigrationBlockInitialized")
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2SystemMigrationBlockInitializedIterator{contract: _ZKEVMV2.contract, event: "SystemMigrationBlockInitialized", logs: logs, sub: sub}, nil
}

// WatchSystemMigrationBlockInitialized is a free log subscription operation binding the contract event 0x405b3b16b9190c1e995514c13ab4e8e7d895d9103e91c3a8c8f12df6cd50aa2c.
//
// Solidity: event SystemMigrationBlockInitialized(uint256 systemMigrationBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchSystemMigrationBlockInitialized(opts *bind.WatchOpts, sink chan<- *ZKEVMV2SystemMigrationBlockInitialized) (event.Subscription, error) {

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "SystemMigrationBlockInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2SystemMigrationBlockInitialized)
				if err := _ZKEVMV2.contract.UnpackLog(event, "SystemMigrationBlockInitialized", log); err != nil {
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

// ParseSystemMigrationBlockInitialized is a log parse operation binding the contract event 0x405b3b16b9190c1e995514c13ab4e8e7d895d9103e91c3a8c8f12df6cd50aa2c.
//
// Solidity: event SystemMigrationBlockInitialized(uint256 systemMigrationBlock)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseSystemMigrationBlockInitialized(log types.Log) (*ZKEVMV2SystemMigrationBlockInitialized, error) {
	event := new(ZKEVMV2SystemMigrationBlockInitialized)
	if err := _ZKEVMV2.contract.UnpackLog(event, "SystemMigrationBlockInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2UnPausedIterator is returned from FilterUnPaused and is used to iterate over the raw logs and unpacked data for UnPaused events raised by the ZKEVMV2 contract.
type ZKEVMV2UnPausedIterator struct {
	Event *ZKEVMV2UnPaused // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2UnPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2UnPaused)
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
		it.Event = new(ZKEVMV2UnPaused)
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
func (it *ZKEVMV2UnPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2UnPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2UnPaused represents a UnPaused event raised by the ZKEVMV2 contract.
type ZKEVMV2UnPaused struct {
	MessageSender common.Address
	PauseType     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUnPaused is a free log retrieval operation binding the contract event 0xef04ba2036ccaeab3a59717b51d2b9146b0b0904077177f1148a5418bf1eae23.
//
// Solidity: event UnPaused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterUnPaused(opts *bind.FilterOpts, pauseType []*big.Int) (*ZKEVMV2UnPausedIterator, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "UnPaused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2UnPausedIterator{contract: _ZKEVMV2.contract, event: "UnPaused", logs: logs, sub: sub}, nil
}

// WatchUnPaused is a free log subscription operation binding the contract event 0xef04ba2036ccaeab3a59717b51d2b9146b0b0904077177f1148a5418bf1eae23.
//
// Solidity: event UnPaused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchUnPaused(opts *bind.WatchOpts, sink chan<- *ZKEVMV2UnPaused, pauseType []*big.Int) (event.Subscription, error) {

	var pauseTypeRule []interface{}
	for _, pauseTypeItem := range pauseType {
		pauseTypeRule = append(pauseTypeRule, pauseTypeItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "UnPaused", pauseTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2UnPaused)
				if err := _ZKEVMV2.contract.UnpackLog(event, "UnPaused", log); err != nil {
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

// ParseUnPaused is a log parse operation binding the contract event 0xef04ba2036ccaeab3a59717b51d2b9146b0b0904077177f1148a5418bf1eae23.
//
// Solidity: event UnPaused(address messageSender, uint256 indexed pauseType)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseUnPaused(log types.Log) (*ZKEVMV2UnPaused, error) {
	event := new(ZKEVMV2UnPaused)
	if err := _ZKEVMV2.contract.UnpackLog(event, "UnPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZKEVMV2VerifierAddressChangedIterator is returned from FilterVerifierAddressChanged and is used to iterate over the raw logs and unpacked data for VerifierAddressChanged events raised by the ZKEVMV2 contract.
type ZKEVMV2VerifierAddressChangedIterator struct {
	Event *ZKEVMV2VerifierAddressChanged // Event containing the contract specifics and raw log

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
func (it *ZKEVMV2VerifierAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZKEVMV2VerifierAddressChanged)
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
		it.Event = new(ZKEVMV2VerifierAddressChanged)
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
func (it *ZKEVMV2VerifierAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZKEVMV2VerifierAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZKEVMV2VerifierAddressChanged represents a VerifierAddressChanged event raised by the ZKEVMV2 contract.
type ZKEVMV2VerifierAddressChanged struct {
	VerifierAddress    common.Address
	ProofType          *big.Int
	VerifierSetBy      common.Address
	OldVerifierAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterVerifierAddressChanged is a free log retrieval operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) FilterVerifierAddressChanged(opts *bind.FilterOpts, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (*ZKEVMV2VerifierAddressChangedIterator, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _ZKEVMV2.contract.FilterLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return &ZKEVMV2VerifierAddressChangedIterator{contract: _ZKEVMV2.contract, event: "VerifierAddressChanged", logs: logs, sub: sub}, nil
}

// WatchVerifierAddressChanged is a free log subscription operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) WatchVerifierAddressChanged(opts *bind.WatchOpts, sink chan<- *ZKEVMV2VerifierAddressChanged, verifierAddress []common.Address, proofType []*big.Int, verifierSetBy []common.Address) (event.Subscription, error) {

	var verifierAddressRule []interface{}
	for _, verifierAddressItem := range verifierAddress {
		verifierAddressRule = append(verifierAddressRule, verifierAddressItem)
	}
	var proofTypeRule []interface{}
	for _, proofTypeItem := range proofType {
		proofTypeRule = append(proofTypeRule, proofTypeItem)
	}
	var verifierSetByRule []interface{}
	for _, verifierSetByItem := range verifierSetBy {
		verifierSetByRule = append(verifierSetByRule, verifierSetByItem)
	}

	logs, sub, err := _ZKEVMV2.contract.WatchLogs(opts, "VerifierAddressChanged", verifierAddressRule, proofTypeRule, verifierSetByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZKEVMV2VerifierAddressChanged)
				if err := _ZKEVMV2.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
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

// ParseVerifierAddressChanged is a log parse operation binding the contract event 0x4a29db3fc6b42bda201e4b4d69ce8d575eeeba5f153509c0d0a342af0f1bd021.
//
// Solidity: event VerifierAddressChanged(address indexed verifierAddress, uint256 indexed proofType, address indexed verifierSetBy, address oldVerifierAddress)
func (_ZKEVMV2 *ZKEVMV2Filterer) ParseVerifierAddressChanged(log types.Log) (*ZKEVMV2VerifierAddressChanged, error) {
	event := new(ZKEVMV2VerifierAddressChanged)
	if err := _ZKEVMV2.contract.UnpackLog(event, "VerifierAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
