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

// WithdrawalQueueBaseBatchesCalculationState is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalQueueBaseBatchesCalculationState struct {
	RemainingEthBudget *big.Int
	Finished           bool
	Batches            [36]*big.Int
	BatchesLength      *big.Int
}

// WithdrawalQueueBaseWithdrawalRequestStatus is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalQueueBaseWithdrawalRequestStatus struct {
	AmountOfStETH  *big.Int
	AmountOfShares *big.Int
	Owner          common.Address
	Timestamp      *big.Int
	IsFinalized    bool
	IsClaimed      bool
}

// WithdrawalQueuePermitInput is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalQueuePermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// StakedETHWithdrawalNFTMetaData contains all meta data concerning the StakedETHWithdrawalNFT contract.
var StakedETHWithdrawalNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wstETH\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdminZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_firstArrayLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondArrayLength\",\"type\":\"uint256\"}],\"name\":\"ArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BatchesAreNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CantSendValueRecipientMayHaveReverted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyBatches\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_hint\",\"type\":\"uint256\"}],\"name\":\"InvalidHint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"InvalidOwnerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportTimestamp\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"InvalidRequestId\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endId\",\"type\":\"uint256\"}],\"name\":\"InvalidRequestIdRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughEther\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOwnerOrApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"NotOwnerOrApprovedForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"RequestAlreadyClaimed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOfStETH\",\"type\":\"uint256\"}],\"name\":\"RequestAmountTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOfStETH\",\"type\":\"uint256\"}],\"name\":\"RequestAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestIdsNotSorted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"RequestNotFoundOrNotFinalized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExpected\",\"type\":\"uint256\"}],\"name\":\"TooMuchEtherToFinalize\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"realOwner\",\"type\":\"address\"}],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"TransferToNonIERC721Receiver\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToThemselves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmountOfETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMetadata\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroShareRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTimestamp\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURISet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"BunkerModeDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sinceTimestamp\",\"type\":\"uint256\"}],\"name\":\"BunkerModeEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"InitializedV1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftDescriptorAddress\",\"type\":\"address\"}],\"name\":\"NftDescriptorAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfETH\",\"type\":\"uint256\"}],\"name\":\"WithdrawalClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requestor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfStETH\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfShares\",\"type\":\"uint256\"}],\"name\":\"WithdrawalRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOfETHLocked\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesToBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"WithdrawalsFinalized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUNKER_MODE_DISABLED_TIMESTAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FINALIZE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_TOKEN_URI_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BATCHES_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_STETH_WITHDRAWAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WSTETH\",\"outputs\":[{\"internalType\":\"contractIWstETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bunkerModeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxShareRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxRequestsPerCall\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remainingEthBudget\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256[36]\",\"name\":\"batches\",\"type\":\"uint256[36]\"},{\"internalType\":\"uint256\",\"name\":\"batchesLength\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawalQueueBase.BatchesCalculationState\",\"name\":\"_state\",\"type\":\"tuple\"}],\"name\":\"calculateFinalizationBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"remainingEthBudget\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"finished\",\"type\":\"bool\"},{\"internalType\":\"uint256[36]\",\"name\":\"batches\",\"type\":\"uint256[36]\"},{\"internalType\":\"uint256\",\"name\":\"batchesLength\",\"type\":\"uint256\"}],\"internalType\":\"structWithdrawalQueueBase.BatchesCalculationState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"claimWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_hints\",\"type\":\"uint256[]\"}],\"name\":\"claimWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_hints\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"claimWithdrawalsTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lastRequestIdToBeFinalized\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxShareRate\",\"type\":\"uint256\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_firstIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lastIndex\",\"type\":\"uint256\"}],\"name\":\"findCheckpointHints\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"hintIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_hints\",\"type\":\"uint256[]\"}],\"name\":\"getClaimableEther\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"claimableEthValues\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastCheckpointIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastFinalizedRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLockedEtherAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTDescriptorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getWithdrawalRequests\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestsIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_requestIds\",\"type\":\"uint256[]\"}],\"name\":\"getWithdrawalStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amountOfStETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOfShares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isClaimed\",\"type\":\"bool\"}],\"internalType\":\"structWithdrawalQueueBase.WithdrawalRequestStatus[]\",\"name\":\"statuses\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isBunkerModeActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isBunkerModeNow\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_bunkerStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_currentReportTimestamp\",\"type\":\"uint256\"}],\"name\":\"onOracleReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseUntilInclusive\",\"type\":\"uint256\"}],\"name\":\"pauseUntil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_batches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxShareRate\",\"type\":\"uint256\"}],\"name\":\"prefinalize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethToLock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sharesToBurn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"requestWithdrawals\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawalQueue.PermitInput\",\"name\":\"_permit\",\"type\":\"tuple\"}],\"name\":\"requestWithdrawalsWithPermit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"requestWithdrawalsWstETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structWithdrawalQueue.PermitInput\",\"name\":\"_permit\",\"type\":\"tuple\"}],\"name\":\"requestWithdrawalsWstETHWithPermit\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"requestIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftDescriptorAddress\",\"type\":\"address\"}],\"name\":\"setNFTDescriptorAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unfinalizedRequestNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unfinalizedStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakedETHWithdrawalNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use StakedETHWithdrawalNFTMetaData.ABI instead.
var StakedETHWithdrawalNFTABI = StakedETHWithdrawalNFTMetaData.ABI

// StakedETHWithdrawalNFT is an auto generated Go binding around an Ethereum contract.
type StakedETHWithdrawalNFT struct {
	StakedETHWithdrawalNFTCaller     // Read-only binding to the contract
	StakedETHWithdrawalNFTTransactor // Write-only binding to the contract
	StakedETHWithdrawalNFTFilterer   // Log filterer for contract events
}

// StakedETHWithdrawalNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakedETHWithdrawalNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHWithdrawalNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakedETHWithdrawalNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHWithdrawalNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakedETHWithdrawalNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakedETHWithdrawalNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakedETHWithdrawalNFTSession struct {
	Contract     *StakedETHWithdrawalNFT // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakedETHWithdrawalNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakedETHWithdrawalNFTCallerSession struct {
	Contract *StakedETHWithdrawalNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StakedETHWithdrawalNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakedETHWithdrawalNFTTransactorSession struct {
	Contract     *StakedETHWithdrawalNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StakedETHWithdrawalNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakedETHWithdrawalNFTRaw struct {
	Contract *StakedETHWithdrawalNFT // Generic contract binding to access the raw methods on
}

// StakedETHWithdrawalNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakedETHWithdrawalNFTCallerRaw struct {
	Contract *StakedETHWithdrawalNFTCaller // Generic read-only contract binding to access the raw methods on
}

// StakedETHWithdrawalNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakedETHWithdrawalNFTTransactorRaw struct {
	Contract *StakedETHWithdrawalNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakedETHWithdrawalNFT creates a new instance of StakedETHWithdrawalNFT, bound to a specific deployed contract.
func NewStakedETHWithdrawalNFT(address common.Address, backend bind.ContractBackend) (*StakedETHWithdrawalNFT, error) {
	contract, err := bindStakedETHWithdrawalNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFT{StakedETHWithdrawalNFTCaller: StakedETHWithdrawalNFTCaller{contract: contract}, StakedETHWithdrawalNFTTransactor: StakedETHWithdrawalNFTTransactor{contract: contract}, StakedETHWithdrawalNFTFilterer: StakedETHWithdrawalNFTFilterer{contract: contract}}, nil
}

// NewStakedETHWithdrawalNFTCaller creates a new read-only instance of StakedETHWithdrawalNFT, bound to a specific deployed contract.
func NewStakedETHWithdrawalNFTCaller(address common.Address, caller bind.ContractCaller) (*StakedETHWithdrawalNFTCaller, error) {
	contract, err := bindStakedETHWithdrawalNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTCaller{contract: contract}, nil
}

// NewStakedETHWithdrawalNFTTransactor creates a new write-only instance of StakedETHWithdrawalNFT, bound to a specific deployed contract.
func NewStakedETHWithdrawalNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*StakedETHWithdrawalNFTTransactor, error) {
	contract, err := bindStakedETHWithdrawalNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTTransactor{contract: contract}, nil
}

// NewStakedETHWithdrawalNFTFilterer creates a new log filterer instance of StakedETHWithdrawalNFT, bound to a specific deployed contract.
func NewStakedETHWithdrawalNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*StakedETHWithdrawalNFTFilterer, error) {
	contract, err := bindStakedETHWithdrawalNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTFilterer{contract: contract}, nil
}

// bindStakedETHWithdrawalNFT binds a generic wrapper to an already deployed contract.
func bindStakedETHWithdrawalNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakedETHWithdrawalNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedETHWithdrawalNFT.Contract.StakedETHWithdrawalNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.StakedETHWithdrawalNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.StakedETHWithdrawalNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakedETHWithdrawalNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.contract.Transact(opts, method, params...)
}

// BUNKERMODEDISABLEDTIMESTAMP is a free data retrieval call binding the contract method 0xe7c0835d.
//
// Solidity: function BUNKER_MODE_DISABLED_TIMESTAMP() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) BUNKERMODEDISABLEDTIMESTAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "BUNKER_MODE_DISABLED_TIMESTAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BUNKERMODEDISABLEDTIMESTAMP is a free data retrieval call binding the contract method 0xe7c0835d.
//
// Solidity: function BUNKER_MODE_DISABLED_TIMESTAMP() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) BUNKERMODEDISABLEDTIMESTAMP() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BUNKERMODEDISABLEDTIMESTAMP(&_StakedETHWithdrawalNFT.CallOpts)
}

// BUNKERMODEDISABLEDTIMESTAMP is a free data retrieval call binding the contract method 0xe7c0835d.
//
// Solidity: function BUNKER_MODE_DISABLED_TIMESTAMP() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) BUNKERMODEDISABLEDTIMESTAMP() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BUNKERMODEDISABLEDTIMESTAMP(&_StakedETHWithdrawalNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.DEFAULTADMINROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.DEFAULTADMINROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// FINALIZEROLE is a free data retrieval call binding the contract method 0x220ca2f4.
//
// Solidity: function FINALIZE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) FINALIZEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "FINALIZE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FINALIZEROLE is a free data retrieval call binding the contract method 0x220ca2f4.
//
// Solidity: function FINALIZE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) FINALIZEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.FINALIZEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// FINALIZEROLE is a free data retrieval call binding the contract method 0x220ca2f4.
//
// Solidity: function FINALIZE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) FINALIZEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.FINALIZEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// MANAGETOKENURIROLE is a free data retrieval call binding the contract method 0xb7bdf748.
//
// Solidity: function MANAGE_TOKEN_URI_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) MANAGETOKENURIROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "MANAGE_TOKEN_URI_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGETOKENURIROLE is a free data retrieval call binding the contract method 0xb7bdf748.
//
// Solidity: function MANAGE_TOKEN_URI_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) MANAGETOKENURIROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.MANAGETOKENURIROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// MANAGETOKENURIROLE is a free data retrieval call binding the contract method 0xb7bdf748.
//
// Solidity: function MANAGE_TOKEN_URI_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) MANAGETOKENURIROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.MANAGETOKENURIROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// MAXBATCHESLENGTH is a free data retrieval call binding the contract method 0x29fd065d.
//
// Solidity: function MAX_BATCHES_LENGTH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) MAXBATCHESLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "MAX_BATCHES_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBATCHESLENGTH is a free data retrieval call binding the contract method 0x29fd065d.
//
// Solidity: function MAX_BATCHES_LENGTH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) MAXBATCHESLENGTH() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MAXBATCHESLENGTH(&_StakedETHWithdrawalNFT.CallOpts)
}

// MAXBATCHESLENGTH is a free data retrieval call binding the contract method 0x29fd065d.
//
// Solidity: function MAX_BATCHES_LENGTH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) MAXBATCHESLENGTH() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MAXBATCHESLENGTH(&_StakedETHWithdrawalNFT.CallOpts)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) MAXSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "MAX_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MAXSTETHWITHDRAWALAMOUNT(&_StakedETHWithdrawalNFT.CallOpts)
}

// MAXSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0xdb2296cd.
//
// Solidity: function MAX_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) MAXSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MAXSTETHWITHDRAWALAMOUNT(&_StakedETHWithdrawalNFT.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) MINSTETHWITHDRAWALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "MIN_STETH_WITHDRAWAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MINSTETHWITHDRAWALAMOUNT(&_StakedETHWithdrawalNFT.CallOpts)
}

// MINSTETHWITHDRAWALAMOUNT is a free data retrieval call binding the contract method 0x0d25a957.
//
// Solidity: function MIN_STETH_WITHDRAWAL_AMOUNT() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) MINSTETHWITHDRAWALAMOUNT() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.MINSTETHWITHDRAWALAMOUNT(&_StakedETHWithdrawalNFT.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) ORACLEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "ORACLE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) ORACLEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.ORACLEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// ORACLEROLE is a free data retrieval call binding the contract method 0x07e2cea5.
//
// Solidity: function ORACLE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) ORACLEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.ORACLEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) PAUSEINFINITELY() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.PAUSEINFINITELY(&_StakedETHWithdrawalNFT.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.PAUSEINFINITELY(&_StakedETHWithdrawalNFT.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) PAUSEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.PAUSEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) PAUSEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.PAUSEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RESUMEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.RESUMEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) RESUMEROLE() ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.RESUMEROLE(&_StakedETHWithdrawalNFT.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) STETH() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.STETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) STETH() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.STETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) WSTETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "WSTETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) WSTETH() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.WSTETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) WSTETH() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.WSTETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BalanceOf(&_StakedETHWithdrawalNFT.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BalanceOf(&_StakedETHWithdrawalNFT.CallOpts, _owner)
}

// BunkerModeSinceTimestamp is a free data retrieval call binding the contract method 0x9b36be58.
//
// Solidity: function bunkerModeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) BunkerModeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "bunkerModeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BunkerModeSinceTimestamp is a free data retrieval call binding the contract method 0x9b36be58.
//
// Solidity: function bunkerModeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) BunkerModeSinceTimestamp() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BunkerModeSinceTimestamp(&_StakedETHWithdrawalNFT.CallOpts)
}

// BunkerModeSinceTimestamp is a free data retrieval call binding the contract method 0x9b36be58.
//
// Solidity: function bunkerModeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) BunkerModeSinceTimestamp() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.BunkerModeSinceTimestamp(&_StakedETHWithdrawalNFT.CallOpts)
}

// CalculateFinalizationBatches is a free data retrieval call binding the contract method 0xeed53bf5.
//
// Solidity: function calculateFinalizationBatches(uint256 _maxShareRate, uint256 _maxTimestamp, uint256 _maxRequestsPerCall, (uint256,bool,uint256[36],uint256) _state) view returns((uint256,bool,uint256[36],uint256))
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) CalculateFinalizationBatches(opts *bind.CallOpts, _maxShareRate *big.Int, _maxTimestamp *big.Int, _maxRequestsPerCall *big.Int, _state WithdrawalQueueBaseBatchesCalculationState) (WithdrawalQueueBaseBatchesCalculationState, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "calculateFinalizationBatches", _maxShareRate, _maxTimestamp, _maxRequestsPerCall, _state)

	if err != nil {
		return *new(WithdrawalQueueBaseBatchesCalculationState), err
	}

	out0 := *abi.ConvertType(out[0], new(WithdrawalQueueBaseBatchesCalculationState)).(*WithdrawalQueueBaseBatchesCalculationState)

	return out0, err

}

// CalculateFinalizationBatches is a free data retrieval call binding the contract method 0xeed53bf5.
//
// Solidity: function calculateFinalizationBatches(uint256 _maxShareRate, uint256 _maxTimestamp, uint256 _maxRequestsPerCall, (uint256,bool,uint256[36],uint256) _state) view returns((uint256,bool,uint256[36],uint256))
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) CalculateFinalizationBatches(_maxShareRate *big.Int, _maxTimestamp *big.Int, _maxRequestsPerCall *big.Int, _state WithdrawalQueueBaseBatchesCalculationState) (WithdrawalQueueBaseBatchesCalculationState, error) {
	return _StakedETHWithdrawalNFT.Contract.CalculateFinalizationBatches(&_StakedETHWithdrawalNFT.CallOpts, _maxShareRate, _maxTimestamp, _maxRequestsPerCall, _state)
}

// CalculateFinalizationBatches is a free data retrieval call binding the contract method 0xeed53bf5.
//
// Solidity: function calculateFinalizationBatches(uint256 _maxShareRate, uint256 _maxTimestamp, uint256 _maxRequestsPerCall, (uint256,bool,uint256[36],uint256) _state) view returns((uint256,bool,uint256[36],uint256))
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) CalculateFinalizationBatches(_maxShareRate *big.Int, _maxTimestamp *big.Int, _maxRequestsPerCall *big.Int, _state WithdrawalQueueBaseBatchesCalculationState) (WithdrawalQueueBaseBatchesCalculationState, error) {
	return _StakedETHWithdrawalNFT.Contract.CalculateFinalizationBatches(&_StakedETHWithdrawalNFT.CallOpts, _maxShareRate, _maxTimestamp, _maxRequestsPerCall, _state)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) FindCheckpointHints(opts *bind.CallOpts, _requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "findCheckpointHints", _requestIds, _firstIndex, _lastIndex)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.FindCheckpointHints(&_StakedETHWithdrawalNFT.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// FindCheckpointHints is a free data retrieval call binding the contract method 0x62abe3fa.
//
// Solidity: function findCheckpointHints(uint256[] _requestIds, uint256 _firstIndex, uint256 _lastIndex) view returns(uint256[] hintIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) FindCheckpointHints(_requestIds []*big.Int, _firstIndex *big.Int, _lastIndex *big.Int) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.FindCheckpointHints(&_StakedETHWithdrawalNFT.CallOpts, _requestIds, _firstIndex, _lastIndex)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetApproved(opts *bind.CallOpts, _requestId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getApproved", _requestId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetApproved(_requestId *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetApproved(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetApproved(_requestId *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetApproved(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetBaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getBaseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetBaseURI() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.GetBaseURI(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetBaseURI is a free data retrieval call binding the contract method 0x714c5398.
//
// Solidity: function getBaseURI() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetBaseURI() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.GetBaseURI(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetClaimableEther is a free data retrieval call binding the contract method 0xc97912d8.
//
// Solidity: function getClaimableEther(uint256[] _requestIds, uint256[] _hints) view returns(uint256[] claimableEthValues)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetClaimableEther(opts *bind.CallOpts, _requestIds []*big.Int, _hints []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getClaimableEther", _requestIds, _hints)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetClaimableEther is a free data retrieval call binding the contract method 0xc97912d8.
//
// Solidity: function getClaimableEther(uint256[] _requestIds, uint256[] _hints) view returns(uint256[] claimableEthValues)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetClaimableEther(_requestIds []*big.Int, _hints []*big.Int) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetClaimableEther(&_StakedETHWithdrawalNFT.CallOpts, _requestIds, _hints)
}

// GetClaimableEther is a free data retrieval call binding the contract method 0xc97912d8.
//
// Solidity: function getClaimableEther(uint256[] _requestIds, uint256[] _hints) view returns(uint256[] claimableEthValues)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetClaimableEther(_requestIds []*big.Int, _hints []*big.Int) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetClaimableEther(&_StakedETHWithdrawalNFT.CallOpts, _requestIds, _hints)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetContractVersion() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetContractVersion(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetContractVersion() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetContractVersion(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetLastCheckpointIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getLastCheckpointIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetLastCheckpointIndex() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastCheckpointIndex(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastCheckpointIndex is a free data retrieval call binding the contract method 0x526eae3e.
//
// Solidity: function getLastCheckpointIndex() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetLastCheckpointIndex() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastCheckpointIndex(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastFinalizedRequestId is a free data retrieval call binding the contract method 0x4f069a13.
//
// Solidity: function getLastFinalizedRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetLastFinalizedRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getLastFinalizedRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastFinalizedRequestId is a free data retrieval call binding the contract method 0x4f069a13.
//
// Solidity: function getLastFinalizedRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetLastFinalizedRequestId() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastFinalizedRequestId(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastFinalizedRequestId is a free data retrieval call binding the contract method 0x4f069a13.
//
// Solidity: function getLastFinalizedRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetLastFinalizedRequestId() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastFinalizedRequestId(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetLastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getLastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetLastRequestId() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastRequestId(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLastRequestId is a free data retrieval call binding the contract method 0x19c2b4c3.
//
// Solidity: function getLastRequestId() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetLastRequestId() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLastRequestId(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLockedEtherAmount is a free data retrieval call binding the contract method 0xf6fa8a47.
//
// Solidity: function getLockedEtherAmount() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetLockedEtherAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getLockedEtherAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLockedEtherAmount is a free data retrieval call binding the contract method 0xf6fa8a47.
//
// Solidity: function getLockedEtherAmount() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetLockedEtherAmount() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLockedEtherAmount(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetLockedEtherAmount is a free data retrieval call binding the contract method 0xf6fa8a47.
//
// Solidity: function getLockedEtherAmount() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetLockedEtherAmount() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetLockedEtherAmount(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetNFTDescriptorAddress is a free data retrieval call binding the contract method 0x46a086b4.
//
// Solidity: function getNFTDescriptorAddress() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetNFTDescriptorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getNFTDescriptorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTDescriptorAddress is a free data retrieval call binding the contract method 0x46a086b4.
//
// Solidity: function getNFTDescriptorAddress() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetNFTDescriptorAddress() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetNFTDescriptorAddress(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetNFTDescriptorAddress is a free data retrieval call binding the contract method 0x46a086b4.
//
// Solidity: function getNFTDescriptorAddress() view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetNFTDescriptorAddress() (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetNFTDescriptorAddress(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetResumeSinceTimestamp(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetResumeSinceTimestamp(&_StakedETHWithdrawalNFT.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleAdmin(&_StakedETHWithdrawalNFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleAdmin(&_StakedETHWithdrawalNFT.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleMember(&_StakedETHWithdrawalNFT.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleMember(&_StakedETHWithdrawalNFT.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleMemberCount(&_StakedETHWithdrawalNFT.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetRoleMemberCount(&_StakedETHWithdrawalNFT.CallOpts, role)
}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0x7d031b65.
//
// Solidity: function getWithdrawalRequests(address _owner) view returns(uint256[] requestsIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetWithdrawalRequests(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getWithdrawalRequests", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0x7d031b65.
//
// Solidity: function getWithdrawalRequests(address _owner) view returns(uint256[] requestsIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetWithdrawalRequests(_owner common.Address) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetWithdrawalRequests(&_StakedETHWithdrawalNFT.CallOpts, _owner)
}

// GetWithdrawalRequests is a free data retrieval call binding the contract method 0x7d031b65.
//
// Solidity: function getWithdrawalRequests(address _owner) view returns(uint256[] requestsIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetWithdrawalRequests(_owner common.Address) ([]*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.GetWithdrawalRequests(&_StakedETHWithdrawalNFT.CallOpts, _owner)
}

// GetWithdrawalStatus is a free data retrieval call binding the contract method 0xb8c4b85a.
//
// Solidity: function getWithdrawalStatus(uint256[] _requestIds) view returns((uint256,uint256,address,uint256,bool,bool)[] statuses)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) GetWithdrawalStatus(opts *bind.CallOpts, _requestIds []*big.Int) ([]WithdrawalQueueBaseWithdrawalRequestStatus, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "getWithdrawalStatus", _requestIds)

	if err != nil {
		return *new([]WithdrawalQueueBaseWithdrawalRequestStatus), err
	}

	out0 := *abi.ConvertType(out[0], new([]WithdrawalQueueBaseWithdrawalRequestStatus)).(*[]WithdrawalQueueBaseWithdrawalRequestStatus)

	return out0, err

}

// GetWithdrawalStatus is a free data retrieval call binding the contract method 0xb8c4b85a.
//
// Solidity: function getWithdrawalStatus(uint256[] _requestIds) view returns((uint256,uint256,address,uint256,bool,bool)[] statuses)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GetWithdrawalStatus(_requestIds []*big.Int) ([]WithdrawalQueueBaseWithdrawalRequestStatus, error) {
	return _StakedETHWithdrawalNFT.Contract.GetWithdrawalStatus(&_StakedETHWithdrawalNFT.CallOpts, _requestIds)
}

// GetWithdrawalStatus is a free data retrieval call binding the contract method 0xb8c4b85a.
//
// Solidity: function getWithdrawalStatus(uint256[] _requestIds) view returns((uint256,uint256,address,uint256,bool,bool)[] statuses)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) GetWithdrawalStatus(_requestIds []*big.Int) ([]WithdrawalQueueBaseWithdrawalRequestStatus, error) {
	return _StakedETHWithdrawalNFT.Contract.GetWithdrawalStatus(&_StakedETHWithdrawalNFT.CallOpts, _requestIds)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.HasRole(&_StakedETHWithdrawalNFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.HasRole(&_StakedETHWithdrawalNFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsApprovedForAll(&_StakedETHWithdrawalNFT.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsApprovedForAll(&_StakedETHWithdrawalNFT.CallOpts, _owner, _operator)
}

// IsBunkerModeActive is a free data retrieval call binding the contract method 0x2b95b781.
//
// Solidity: function isBunkerModeActive() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) IsBunkerModeActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "isBunkerModeActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBunkerModeActive is a free data retrieval call binding the contract method 0x2b95b781.
//
// Solidity: function isBunkerModeActive() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) IsBunkerModeActive() (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsBunkerModeActive(&_StakedETHWithdrawalNFT.CallOpts)
}

// IsBunkerModeActive is a free data retrieval call binding the contract method 0x2b95b781.
//
// Solidity: function isBunkerModeActive() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) IsBunkerModeActive() (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsBunkerModeActive(&_StakedETHWithdrawalNFT.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) IsPaused() (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsPaused(&_StakedETHWithdrawalNFT.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) IsPaused() (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.IsPaused(&_StakedETHWithdrawalNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Name() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.Name(&_StakedETHWithdrawalNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) Name() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.Name(&_StakedETHWithdrawalNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) OwnerOf(opts *bind.CallOpts, _requestId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "ownerOf", _requestId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) OwnerOf(_requestId *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.OwnerOf(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _requestId) view returns(address)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) OwnerOf(_requestId *big.Int) (common.Address, error) {
	return _StakedETHWithdrawalNFT.Contract.OwnerOf(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// Prefinalize is a free data retrieval call binding the contract method 0xa52e9c9f.
//
// Solidity: function prefinalize(uint256[] _batches, uint256 _maxShareRate) view returns(uint256 ethToLock, uint256 sharesToBurn)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) Prefinalize(opts *bind.CallOpts, _batches []*big.Int, _maxShareRate *big.Int) (struct {
	EthToLock    *big.Int
	SharesToBurn *big.Int
}, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "prefinalize", _batches, _maxShareRate)

	outstruct := new(struct {
		EthToLock    *big.Int
		SharesToBurn *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthToLock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SharesToBurn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Prefinalize is a free data retrieval call binding the contract method 0xa52e9c9f.
//
// Solidity: function prefinalize(uint256[] _batches, uint256 _maxShareRate) view returns(uint256 ethToLock, uint256 sharesToBurn)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Prefinalize(_batches []*big.Int, _maxShareRate *big.Int) (struct {
	EthToLock    *big.Int
	SharesToBurn *big.Int
}, error) {
	return _StakedETHWithdrawalNFT.Contract.Prefinalize(&_StakedETHWithdrawalNFT.CallOpts, _batches, _maxShareRate)
}

// Prefinalize is a free data retrieval call binding the contract method 0xa52e9c9f.
//
// Solidity: function prefinalize(uint256[] _batches, uint256 _maxShareRate) view returns(uint256 ethToLock, uint256 sharesToBurn)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) Prefinalize(_batches []*big.Int, _maxShareRate *big.Int) (struct {
	EthToLock    *big.Int
	SharesToBurn *big.Int
}, error) {
	return _StakedETHWithdrawalNFT.Contract.Prefinalize(&_StakedETHWithdrawalNFT.CallOpts, _batches, _maxShareRate)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.SupportsInterface(&_StakedETHWithdrawalNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakedETHWithdrawalNFT.Contract.SupportsInterface(&_StakedETHWithdrawalNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Symbol() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.Symbol(&_StakedETHWithdrawalNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) Symbol() (string, error) {
	return _StakedETHWithdrawalNFT.Contract.Symbol(&_StakedETHWithdrawalNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _requestId) view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) TokenURI(opts *bind.CallOpts, _requestId *big.Int) (string, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "tokenURI", _requestId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _requestId) view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) TokenURI(_requestId *big.Int) (string, error) {
	return _StakedETHWithdrawalNFT.Contract.TokenURI(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _requestId) view returns(string)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) TokenURI(_requestId *big.Int) (string, error) {
	return _StakedETHWithdrawalNFT.Contract.TokenURI(&_StakedETHWithdrawalNFT.CallOpts, _requestId)
}

// UnfinalizedRequestNumber is a free data retrieval call binding the contract method 0xc2fc7aff.
//
// Solidity: function unfinalizedRequestNumber() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) UnfinalizedRequestNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "unfinalizedRequestNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnfinalizedRequestNumber is a free data retrieval call binding the contract method 0xc2fc7aff.
//
// Solidity: function unfinalizedRequestNumber() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) UnfinalizedRequestNumber() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.UnfinalizedRequestNumber(&_StakedETHWithdrawalNFT.CallOpts)
}

// UnfinalizedRequestNumber is a free data retrieval call binding the contract method 0xc2fc7aff.
//
// Solidity: function unfinalizedRequestNumber() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) UnfinalizedRequestNumber() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.UnfinalizedRequestNumber(&_StakedETHWithdrawalNFT.CallOpts)
}

// UnfinalizedStETH is a free data retrieval call binding the contract method 0xd0fb84e8.
//
// Solidity: function unfinalizedStETH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCaller) UnfinalizedStETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakedETHWithdrawalNFT.contract.Call(opts, &out, "unfinalizedStETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnfinalizedStETH is a free data retrieval call binding the contract method 0xd0fb84e8.
//
// Solidity: function unfinalizedStETH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) UnfinalizedStETH() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.UnfinalizedStETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// UnfinalizedStETH is a free data retrieval call binding the contract method 0xd0fb84e8.
//
// Solidity: function unfinalizedStETH() view returns(uint256)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTCallerSession) UnfinalizedStETH() (*big.Int, error) {
	return _StakedETHWithdrawalNFT.Contract.UnfinalizedStETH(&_StakedETHWithdrawalNFT.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "approve", _to, _requestId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Approve(_to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Approve(&_StakedETHWithdrawalNFT.TransactOpts, _to, _requestId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) Approve(_to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Approve(&_StakedETHWithdrawalNFT.TransactOpts, _to, _requestId)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) ClaimWithdrawal(opts *bind.TransactOpts, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "claimWithdrawal", _requestId)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) ClaimWithdrawal(_requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawal(&_StakedETHWithdrawalNFT.TransactOpts, _requestId)
}

// ClaimWithdrawal is a paid mutator transaction binding the contract method 0xf8444436.
//
// Solidity: function claimWithdrawal(uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) ClaimWithdrawal(_requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawal(&_StakedETHWithdrawalNFT.TransactOpts, _requestId)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) ClaimWithdrawals(opts *bind.TransactOpts, _requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "claimWithdrawals", _requestIds, _hints)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) ClaimWithdrawals(_requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawals(&_StakedETHWithdrawalNFT.TransactOpts, _requestIds, _hints)
}

// ClaimWithdrawals is a paid mutator transaction binding the contract method 0xe3afe0a3.
//
// Solidity: function claimWithdrawals(uint256[] _requestIds, uint256[] _hints) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) ClaimWithdrawals(_requestIds []*big.Int, _hints []*big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawals(&_StakedETHWithdrawalNFT.TransactOpts, _requestIds, _hints)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) ClaimWithdrawalsTo(opts *bind.TransactOpts, _requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "claimWithdrawalsTo", _requestIds, _hints, _recipient)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) ClaimWithdrawalsTo(_requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawalsTo(&_StakedETHWithdrawalNFT.TransactOpts, _requestIds, _hints, _recipient)
}

// ClaimWithdrawalsTo is a paid mutator transaction binding the contract method 0x5e7eead9.
//
// Solidity: function claimWithdrawalsTo(uint256[] _requestIds, uint256[] _hints, address _recipient) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) ClaimWithdrawalsTo(_requestIds []*big.Int, _hints []*big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.ClaimWithdrawalsTo(&_StakedETHWithdrawalNFT.TransactOpts, _requestIds, _hints, _recipient)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) Finalize(opts *bind.TransactOpts, _lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "finalize", _lastRequestIdToBeFinalized, _maxShareRate)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Finalize(_lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Finalize(&_StakedETHWithdrawalNFT.TransactOpts, _lastRequestIdToBeFinalized, _maxShareRate)
}

// Finalize is a paid mutator transaction binding the contract method 0xb6013cef.
//
// Solidity: function finalize(uint256 _lastRequestIdToBeFinalized, uint256 _maxShareRate) payable returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) Finalize(_lastRequestIdToBeFinalized *big.Int, _maxShareRate *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Finalize(&_StakedETHWithdrawalNFT.TransactOpts, _lastRequestIdToBeFinalized, _maxShareRate)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.GrantRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.GrantRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "initialize", _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Initialize(&_StakedETHWithdrawalNFT.TransactOpts, _admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _admin) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) Initialize(_admin common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Initialize(&_StakedETHWithdrawalNFT.TransactOpts, _admin)
}

// OnOracleReport is a paid mutator transaction binding the contract method 0x96992fed.
//
// Solidity: function onOracleReport(bool _isBunkerModeNow, uint256 _bunkerStartTimestamp, uint256 _currentReportTimestamp) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) OnOracleReport(opts *bind.TransactOpts, _isBunkerModeNow bool, _bunkerStartTimestamp *big.Int, _currentReportTimestamp *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "onOracleReport", _isBunkerModeNow, _bunkerStartTimestamp, _currentReportTimestamp)
}

// OnOracleReport is a paid mutator transaction binding the contract method 0x96992fed.
//
// Solidity: function onOracleReport(bool _isBunkerModeNow, uint256 _bunkerStartTimestamp, uint256 _currentReportTimestamp) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) OnOracleReport(_isBunkerModeNow bool, _bunkerStartTimestamp *big.Int, _currentReportTimestamp *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.OnOracleReport(&_StakedETHWithdrawalNFT.TransactOpts, _isBunkerModeNow, _bunkerStartTimestamp, _currentReportTimestamp)
}

// OnOracleReport is a paid mutator transaction binding the contract method 0x96992fed.
//
// Solidity: function onOracleReport(bool _isBunkerModeNow, uint256 _bunkerStartTimestamp, uint256 _currentReportTimestamp) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) OnOracleReport(_isBunkerModeNow bool, _bunkerStartTimestamp *big.Int, _currentReportTimestamp *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.OnOracleReport(&_StakedETHWithdrawalNFT.TransactOpts, _isBunkerModeNow, _bunkerStartTimestamp, _currentReportTimestamp)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) PauseFor(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "pauseFor", _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.PauseFor(&_StakedETHWithdrawalNFT.TransactOpts, _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.PauseFor(&_StakedETHWithdrawalNFT.TransactOpts, _duration)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) PauseUntil(opts *bind.TransactOpts, _pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "pauseUntil", _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.PauseUntil(&_StakedETHWithdrawalNFT.TransactOpts, _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.PauseUntil(&_StakedETHWithdrawalNFT.TransactOpts, _pauseUntilInclusive)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RenounceRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RenounceRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RequestWithdrawals(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "requestWithdrawals", _amounts, _owner)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RequestWithdrawals(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawals(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner)
}

// RequestWithdrawals is a paid mutator transaction binding the contract method 0xd6681042.
//
// Solidity: function requestWithdrawals(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RequestWithdrawals(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawals(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RequestWithdrawalsWithPermit(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "requestWithdrawalsWithPermit", _amounts, _owner, _permit)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RequestWithdrawalsWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWithPermit(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner, _permit)
}

// RequestWithdrawalsWithPermit is a paid mutator transaction binding the contract method 0xacf41e4d.
//
// Solidity: function requestWithdrawalsWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RequestWithdrawalsWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWithPermit(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner, _permit)
}

// RequestWithdrawalsWstETH is a paid mutator transaction binding the contract method 0x19aa6257.
//
// Solidity: function requestWithdrawalsWstETH(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RequestWithdrawalsWstETH(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "requestWithdrawalsWstETH", _amounts, _owner)
}

// RequestWithdrawalsWstETH is a paid mutator transaction binding the contract method 0x19aa6257.
//
// Solidity: function requestWithdrawalsWstETH(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RequestWithdrawalsWstETH(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWstETH(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner)
}

// RequestWithdrawalsWstETH is a paid mutator transaction binding the contract method 0x19aa6257.
//
// Solidity: function requestWithdrawalsWstETH(uint256[] _amounts, address _owner) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RequestWithdrawalsWstETH(_amounts []*big.Int, _owner common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWstETH(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner)
}

// RequestWithdrawalsWstETHWithPermit is a paid mutator transaction binding the contract method 0x7951b76f.
//
// Solidity: function requestWithdrawalsWstETHWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RequestWithdrawalsWstETHWithPermit(opts *bind.TransactOpts, _amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "requestWithdrawalsWstETHWithPermit", _amounts, _owner, _permit)
}

// RequestWithdrawalsWstETHWithPermit is a paid mutator transaction binding the contract method 0x7951b76f.
//
// Solidity: function requestWithdrawalsWstETHWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RequestWithdrawalsWstETHWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWstETHWithPermit(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner, _permit)
}

// RequestWithdrawalsWstETHWithPermit is a paid mutator transaction binding the contract method 0x7951b76f.
//
// Solidity: function requestWithdrawalsWstETHWithPermit(uint256[] _amounts, address _owner, (uint256,uint256,uint8,bytes32,bytes32) _permit) returns(uint256[] requestIds)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RequestWithdrawalsWstETHWithPermit(_amounts []*big.Int, _owner common.Address, _permit WithdrawalQueuePermitInput) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RequestWithdrawalsWstETHWithPermit(&_StakedETHWithdrawalNFT.TransactOpts, _amounts, _owner, _permit)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) Resume() (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Resume(&_StakedETHWithdrawalNFT.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) Resume() (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.Resume(&_StakedETHWithdrawalNFT.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RevokeRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.RevokeRole(&_StakedETHWithdrawalNFT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "safeTransferFrom", _from, _to, _requestId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SafeTransferFrom(_from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SafeTransferFrom(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SafeTransferFrom(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId, bytes _data) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _requestId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "safeTransferFrom0", _from, _to, _requestId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId, bytes _data) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SafeTransferFrom0(_from common.Address, _to common.Address, _requestId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SafeTransferFrom0(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _requestId, bytes _data) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _requestId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SafeTransferFrom0(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetApprovalForAll(&_StakedETHWithdrawalNFT.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetApprovalForAll(&_StakedETHWithdrawalNFT.TransactOpts, _operator, _approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetBaseURI(&_StakedETHWithdrawalNFT.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetBaseURI(&_StakedETHWithdrawalNFT.TransactOpts, _baseURI)
}

// SetNFTDescriptorAddress is a paid mutator transaction binding the contract method 0x92b18a47.
//
// Solidity: function setNFTDescriptorAddress(address _nftDescriptorAddress) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) SetNFTDescriptorAddress(opts *bind.TransactOpts, _nftDescriptorAddress common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "setNFTDescriptorAddress", _nftDescriptorAddress)
}

// SetNFTDescriptorAddress is a paid mutator transaction binding the contract method 0x92b18a47.
//
// Solidity: function setNFTDescriptorAddress(address _nftDescriptorAddress) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) SetNFTDescriptorAddress(_nftDescriptorAddress common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetNFTDescriptorAddress(&_StakedETHWithdrawalNFT.TransactOpts, _nftDescriptorAddress)
}

// SetNFTDescriptorAddress is a paid mutator transaction binding the contract method 0x92b18a47.
//
// Solidity: function setNFTDescriptorAddress(address _nftDescriptorAddress) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) SetNFTDescriptorAddress(_nftDescriptorAddress common.Address) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.SetNFTDescriptorAddress(&_StakedETHWithdrawalNFT.TransactOpts, _nftDescriptorAddress)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.contract.Transact(opts, "transferFrom", _from, _to, _requestId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTSession) TransferFrom(_from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.TransferFrom(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _requestId) returns()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTTransactorSession) TransferFrom(_from common.Address, _to common.Address, _requestId *big.Int) (*types.Transaction, error) {
	return _StakedETHWithdrawalNFT.Contract.TransferFrom(&_StakedETHWithdrawalNFT.TransactOpts, _from, _to, _requestId)
}

// StakedETHWithdrawalNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTApprovalIterator struct {
	Event *StakedETHWithdrawalNFTApproval // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTApproval)
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
		it.Event = new(StakedETHWithdrawalNFTApproval)
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
func (it *StakedETHWithdrawalNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTApproval represents a Approval event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StakedETHWithdrawalNFTApprovalIterator, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTApprovalIterator{contract: _StakedETHWithdrawalNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTApproval)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseApproval(log types.Log) (*StakedETHWithdrawalNFTApproval, error) {
	event := new(StakedETHWithdrawalNFTApproval)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTApprovalForAllIterator struct {
	Event *StakedETHWithdrawalNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTApprovalForAll)
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
		it.Event = new(StakedETHWithdrawalNFTApprovalForAll)
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
func (it *StakedETHWithdrawalNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTApprovalForAll represents a ApprovalForAll event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StakedETHWithdrawalNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTApprovalForAllIterator{contract: _StakedETHWithdrawalNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTApprovalForAll)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseApprovalForAll(log types.Log) (*StakedETHWithdrawalNFTApprovalForAll, error) {
	event := new(StakedETHWithdrawalNFTApprovalForAll)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTBaseURISetIterator is returned from FilterBaseURISet and is used to iterate over the raw logs and unpacked data for BaseURISet events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBaseURISetIterator struct {
	Event *StakedETHWithdrawalNFTBaseURISet // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTBaseURISetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTBaseURISet)
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
		it.Event = new(StakedETHWithdrawalNFTBaseURISet)
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
func (it *StakedETHWithdrawalNFTBaseURISetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTBaseURISetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTBaseURISet represents a BaseURISet event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBaseURISet struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURISet is a free log retrieval operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string baseURI)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterBaseURISet(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTBaseURISetIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "BaseURISet")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTBaseURISetIterator{contract: _StakedETHWithdrawalNFT.contract, event: "BaseURISet", logs: logs, sub: sub}, nil
}

// WatchBaseURISet is a free log subscription operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string baseURI)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchBaseURISet(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTBaseURISet) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "BaseURISet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTBaseURISet)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BaseURISet", log); err != nil {
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

// ParseBaseURISet is a log parse operation binding the contract event 0xf9c7803e94e0d3c02900d8a90893a6d5e90dd04d32a4cfe825520f82bf9f32f6.
//
// Solidity: event BaseURISet(string baseURI)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseBaseURISet(log types.Log) (*StakedETHWithdrawalNFTBaseURISet, error) {
	event := new(StakedETHWithdrawalNFTBaseURISet)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BaseURISet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBatchMetadataUpdateIterator struct {
	Event *StakedETHWithdrawalNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTBatchMetadataUpdate)
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
		it.Event = new(StakedETHWithdrawalNFTBatchMetadataUpdate)
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
func (it *StakedETHWithdrawalNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTBatchMetadataUpdateIterator{contract: _StakedETHWithdrawalNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTBatchMetadataUpdate)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*StakedETHWithdrawalNFTBatchMetadataUpdate, error) {
	event := new(StakedETHWithdrawalNFTBatchMetadataUpdate)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTBunkerModeDisabledIterator is returned from FilterBunkerModeDisabled and is used to iterate over the raw logs and unpacked data for BunkerModeDisabled events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBunkerModeDisabledIterator struct {
	Event *StakedETHWithdrawalNFTBunkerModeDisabled // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTBunkerModeDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTBunkerModeDisabled)
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
		it.Event = new(StakedETHWithdrawalNFTBunkerModeDisabled)
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
func (it *StakedETHWithdrawalNFTBunkerModeDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTBunkerModeDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTBunkerModeDisabled represents a BunkerModeDisabled event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBunkerModeDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBunkerModeDisabled is a free log retrieval operation binding the contract event 0xd1f8a2998c0caf73e09434aa93d273a599060d789407c6f70ccd4c9c9f32c8f4.
//
// Solidity: event BunkerModeDisabled()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterBunkerModeDisabled(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTBunkerModeDisabledIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "BunkerModeDisabled")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTBunkerModeDisabledIterator{contract: _StakedETHWithdrawalNFT.contract, event: "BunkerModeDisabled", logs: logs, sub: sub}, nil
}

// WatchBunkerModeDisabled is a free log subscription operation binding the contract event 0xd1f8a2998c0caf73e09434aa93d273a599060d789407c6f70ccd4c9c9f32c8f4.
//
// Solidity: event BunkerModeDisabled()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchBunkerModeDisabled(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTBunkerModeDisabled) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "BunkerModeDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTBunkerModeDisabled)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BunkerModeDisabled", log); err != nil {
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

// ParseBunkerModeDisabled is a log parse operation binding the contract event 0xd1f8a2998c0caf73e09434aa93d273a599060d789407c6f70ccd4c9c9f32c8f4.
//
// Solidity: event BunkerModeDisabled()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseBunkerModeDisabled(log types.Log) (*StakedETHWithdrawalNFTBunkerModeDisabled, error) {
	event := new(StakedETHWithdrawalNFTBunkerModeDisabled)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BunkerModeDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTBunkerModeEnabledIterator is returned from FilterBunkerModeEnabled and is used to iterate over the raw logs and unpacked data for BunkerModeEnabled events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBunkerModeEnabledIterator struct {
	Event *StakedETHWithdrawalNFTBunkerModeEnabled // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTBunkerModeEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTBunkerModeEnabled)
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
		it.Event = new(StakedETHWithdrawalNFTBunkerModeEnabled)
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
func (it *StakedETHWithdrawalNFTBunkerModeEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTBunkerModeEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTBunkerModeEnabled represents a BunkerModeEnabled event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTBunkerModeEnabled struct {
	SinceTimestamp *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBunkerModeEnabled is a free log retrieval operation binding the contract event 0x47f03b07e5b5377f871539bb2942f5ecb72733be9fc9d55a17b6d6a05d418345.
//
// Solidity: event BunkerModeEnabled(uint256 _sinceTimestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterBunkerModeEnabled(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTBunkerModeEnabledIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "BunkerModeEnabled")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTBunkerModeEnabledIterator{contract: _StakedETHWithdrawalNFT.contract, event: "BunkerModeEnabled", logs: logs, sub: sub}, nil
}

// WatchBunkerModeEnabled is a free log subscription operation binding the contract event 0x47f03b07e5b5377f871539bb2942f5ecb72733be9fc9d55a17b6d6a05d418345.
//
// Solidity: event BunkerModeEnabled(uint256 _sinceTimestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchBunkerModeEnabled(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTBunkerModeEnabled) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "BunkerModeEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTBunkerModeEnabled)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BunkerModeEnabled", log); err != nil {
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

// ParseBunkerModeEnabled is a log parse operation binding the contract event 0x47f03b07e5b5377f871539bb2942f5ecb72733be9fc9d55a17b6d6a05d418345.
//
// Solidity: event BunkerModeEnabled(uint256 _sinceTimestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseBunkerModeEnabled(log types.Log) (*StakedETHWithdrawalNFTBunkerModeEnabled, error) {
	event := new(StakedETHWithdrawalNFTBunkerModeEnabled)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "BunkerModeEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTContractVersionSetIterator struct {
	Event *StakedETHWithdrawalNFTContractVersionSet // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTContractVersionSet)
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
		it.Event = new(StakedETHWithdrawalNFTContractVersionSet)
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
func (it *StakedETHWithdrawalNFTContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTContractVersionSet represents a ContractVersionSet event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTContractVersionSetIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTContractVersionSetIterator{contract: _StakedETHWithdrawalNFT.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTContractVersionSet)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseContractVersionSet(log types.Log) (*StakedETHWithdrawalNFTContractVersionSet, error) {
	event := new(StakedETHWithdrawalNFTContractVersionSet)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTInitializedV1Iterator is returned from FilterInitializedV1 and is used to iterate over the raw logs and unpacked data for InitializedV1 events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTInitializedV1Iterator struct {
	Event *StakedETHWithdrawalNFTInitializedV1 // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTInitializedV1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTInitializedV1)
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
		it.Event = new(StakedETHWithdrawalNFTInitializedV1)
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
func (it *StakedETHWithdrawalNFTInitializedV1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTInitializedV1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTInitializedV1 represents a InitializedV1 event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTInitializedV1 struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterInitializedV1 is a free log retrieval operation binding the contract event 0x20b34d2aaaf6acb4fbbc9c4846858bb824053ab11ff44a59dfba1e22ceb8a509.
//
// Solidity: event InitializedV1(address _admin)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterInitializedV1(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTInitializedV1Iterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "InitializedV1")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTInitializedV1Iterator{contract: _StakedETHWithdrawalNFT.contract, event: "InitializedV1", logs: logs, sub: sub}, nil
}

// WatchInitializedV1 is a free log subscription operation binding the contract event 0x20b34d2aaaf6acb4fbbc9c4846858bb824053ab11ff44a59dfba1e22ceb8a509.
//
// Solidity: event InitializedV1(address _admin)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchInitializedV1(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTInitializedV1) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "InitializedV1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTInitializedV1)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "InitializedV1", log); err != nil {
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

// ParseInitializedV1 is a log parse operation binding the contract event 0x20b34d2aaaf6acb4fbbc9c4846858bb824053ab11ff44a59dfba1e22ceb8a509.
//
// Solidity: event InitializedV1(address _admin)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseInitializedV1(log types.Log) (*StakedETHWithdrawalNFTInitializedV1, error) {
	event := new(StakedETHWithdrawalNFTInitializedV1)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "InitializedV1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTMetadataUpdateIterator struct {
	Event *StakedETHWithdrawalNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTMetadataUpdate)
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
		it.Event = new(StakedETHWithdrawalNFTMetadataUpdate)
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
func (it *StakedETHWithdrawalNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTMetadataUpdate represents a MetadataUpdate event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTMetadataUpdateIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTMetadataUpdateIterator{contract: _StakedETHWithdrawalNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTMetadataUpdate)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseMetadataUpdate(log types.Log) (*StakedETHWithdrawalNFTMetadataUpdate, error) {
	event := new(StakedETHWithdrawalNFTMetadataUpdate)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTNftDescriptorAddressSetIterator is returned from FilterNftDescriptorAddressSet and is used to iterate over the raw logs and unpacked data for NftDescriptorAddressSet events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTNftDescriptorAddressSetIterator struct {
	Event *StakedETHWithdrawalNFTNftDescriptorAddressSet // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTNftDescriptorAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTNftDescriptorAddressSet)
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
		it.Event = new(StakedETHWithdrawalNFTNftDescriptorAddressSet)
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
func (it *StakedETHWithdrawalNFTNftDescriptorAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTNftDescriptorAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTNftDescriptorAddressSet represents a NftDescriptorAddressSet event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTNftDescriptorAddressSet struct {
	NftDescriptorAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNftDescriptorAddressSet is a free log retrieval operation binding the contract event 0x4ec04ac71c49eea0a94dc5967b493412a8cdb2934b367713019d3b110e9f0ba8.
//
// Solidity: event NftDescriptorAddressSet(address nftDescriptorAddress)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterNftDescriptorAddressSet(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTNftDescriptorAddressSetIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "NftDescriptorAddressSet")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTNftDescriptorAddressSetIterator{contract: _StakedETHWithdrawalNFT.contract, event: "NftDescriptorAddressSet", logs: logs, sub: sub}, nil
}

// WatchNftDescriptorAddressSet is a free log subscription operation binding the contract event 0x4ec04ac71c49eea0a94dc5967b493412a8cdb2934b367713019d3b110e9f0ba8.
//
// Solidity: event NftDescriptorAddressSet(address nftDescriptorAddress)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchNftDescriptorAddressSet(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTNftDescriptorAddressSet) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "NftDescriptorAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTNftDescriptorAddressSet)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "NftDescriptorAddressSet", log); err != nil {
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

// ParseNftDescriptorAddressSet is a log parse operation binding the contract event 0x4ec04ac71c49eea0a94dc5967b493412a8cdb2934b367713019d3b110e9f0ba8.
//
// Solidity: event NftDescriptorAddressSet(address nftDescriptorAddress)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseNftDescriptorAddressSet(log types.Log) (*StakedETHWithdrawalNFTNftDescriptorAddressSet, error) {
	event := new(StakedETHWithdrawalNFTNftDescriptorAddressSet)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "NftDescriptorAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTPausedIterator struct {
	Event *StakedETHWithdrawalNFTPaused // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTPaused)
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
		it.Event = new(StakedETHWithdrawalNFTPaused)
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
func (it *StakedETHWithdrawalNFTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTPaused represents a Paused event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterPaused(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTPausedIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTPausedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTPaused) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTPaused)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParsePaused(log types.Log) (*StakedETHWithdrawalNFTPaused, error) {
	event := new(StakedETHWithdrawalNFTPaused)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTResumedIterator struct {
	Event *StakedETHWithdrawalNFTResumed // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTResumed)
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
		it.Event = new(StakedETHWithdrawalNFTResumed)
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
func (it *StakedETHWithdrawalNFTResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTResumed represents a Resumed event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterResumed(opts *bind.FilterOpts) (*StakedETHWithdrawalNFTResumedIterator, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTResumedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTResumed) (event.Subscription, error) {

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTResumed)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseResumed(log types.Log) (*StakedETHWithdrawalNFTResumed, error) {
	event := new(StakedETHWithdrawalNFTResumed)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleAdminChangedIterator struct {
	Event *StakedETHWithdrawalNFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTRoleAdminChanged)
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
		it.Event = new(StakedETHWithdrawalNFTRoleAdminChanged)
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
func (it *StakedETHWithdrawalNFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTRoleAdminChanged represents a RoleAdminChanged event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakedETHWithdrawalNFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTRoleAdminChangedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTRoleAdminChanged)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseRoleAdminChanged(log types.Log) (*StakedETHWithdrawalNFTRoleAdminChanged, error) {
	event := new(StakedETHWithdrawalNFTRoleAdminChanged)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleGrantedIterator struct {
	Event *StakedETHWithdrawalNFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTRoleGranted)
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
		it.Event = new(StakedETHWithdrawalNFTRoleGranted)
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
func (it *StakedETHWithdrawalNFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTRoleGranted represents a RoleGranted event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedETHWithdrawalNFTRoleGrantedIterator, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTRoleGrantedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTRoleGranted)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseRoleGranted(log types.Log) (*StakedETHWithdrawalNFTRoleGranted, error) {
	event := new(StakedETHWithdrawalNFTRoleGranted)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleRevokedIterator struct {
	Event *StakedETHWithdrawalNFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTRoleRevoked)
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
		it.Event = new(StakedETHWithdrawalNFTRoleRevoked)
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
func (it *StakedETHWithdrawalNFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTRoleRevoked represents a RoleRevoked event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakedETHWithdrawalNFTRoleRevokedIterator, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTRoleRevokedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTRoleRevoked)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseRoleRevoked(log types.Log) (*StakedETHWithdrawalNFTRoleRevoked, error) {
	event := new(StakedETHWithdrawalNFTRoleRevoked)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTTransferIterator struct {
	Event *StakedETHWithdrawalNFTTransfer // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTTransfer)
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
		it.Event = new(StakedETHWithdrawalNFTTransfer)
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
func (it *StakedETHWithdrawalNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTTransfer represents a Transfer event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StakedETHWithdrawalNFTTransferIterator, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTTransferIterator{contract: _StakedETHWithdrawalNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTTransfer)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseTransfer(log types.Log) (*StakedETHWithdrawalNFTTransfer, error) {
	event := new(StakedETHWithdrawalNFTTransfer)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTWithdrawalClaimedIterator is returned from FilterWithdrawalClaimed and is used to iterate over the raw logs and unpacked data for WithdrawalClaimed events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalClaimedIterator struct {
	Event *StakedETHWithdrawalNFTWithdrawalClaimed // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTWithdrawalClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTWithdrawalClaimed)
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
		it.Event = new(StakedETHWithdrawalNFTWithdrawalClaimed)
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
func (it *StakedETHWithdrawalNFTWithdrawalClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTWithdrawalClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTWithdrawalClaimed represents a WithdrawalClaimed event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalClaimed struct {
	RequestId   *big.Int
	Owner       common.Address
	Receiver    common.Address
	AmountOfETH *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalClaimed is a free log retrieval operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterWithdrawalClaimed(opts *bind.FilterOpts, requestId []*big.Int, owner []common.Address, receiver []common.Address) (*StakedETHWithdrawalNFTWithdrawalClaimedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "WithdrawalClaimed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTWithdrawalClaimedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "WithdrawalClaimed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalClaimed is a free log subscription operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchWithdrawalClaimed(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTWithdrawalClaimed, requestId []*big.Int, owner []common.Address, receiver []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "WithdrawalClaimed", requestIdRule, ownerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTWithdrawalClaimed)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
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

// ParseWithdrawalClaimed is a log parse operation binding the contract event 0x6ad26c5e238e7d002799f9a5db07e81ef14e37386ae03496d7a7ef04713e145b.
//
// Solidity: event WithdrawalClaimed(uint256 indexed requestId, address indexed owner, address indexed receiver, uint256 amountOfETH)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseWithdrawalClaimed(log types.Log) (*StakedETHWithdrawalNFTWithdrawalClaimed, error) {
	event := new(StakedETHWithdrawalNFTWithdrawalClaimed)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTWithdrawalRequestedIterator is returned from FilterWithdrawalRequested and is used to iterate over the raw logs and unpacked data for WithdrawalRequested events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalRequestedIterator struct {
	Event *StakedETHWithdrawalNFTWithdrawalRequested // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTWithdrawalRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTWithdrawalRequested)
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
		it.Event = new(StakedETHWithdrawalNFTWithdrawalRequested)
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
func (it *StakedETHWithdrawalNFTWithdrawalRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTWithdrawalRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTWithdrawalRequested represents a WithdrawalRequested event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalRequested struct {
	RequestId      *big.Int
	Requestor      common.Address
	Owner          common.Address
	AmountOfStETH  *big.Int
	AmountOfShares *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalRequested is a free log retrieval operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterWithdrawalRequested(opts *bind.FilterOpts, requestId []*big.Int, requestor []common.Address, owner []common.Address) (*StakedETHWithdrawalNFTWithdrawalRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestorRule []interface{}
	for _, requestorItem := range requestor {
		requestorRule = append(requestorRule, requestorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "WithdrawalRequested", requestIdRule, requestorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTWithdrawalRequestedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "WithdrawalRequested", logs: logs, sub: sub}, nil
}

// WatchWithdrawalRequested is a free log subscription operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchWithdrawalRequested(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTWithdrawalRequested, requestId []*big.Int, requestor []common.Address, owner []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}
	var requestorRule []interface{}
	for _, requestorItem := range requestor {
		requestorRule = append(requestorRule, requestorItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "WithdrawalRequested", requestIdRule, requestorRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTWithdrawalRequested)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
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

// ParseWithdrawalRequested is a log parse operation binding the contract event 0xf0cb471f23fb74ea44b8252eb1881a2dca546288d9f6e90d1a0e82fe0ed342ab.
//
// Solidity: event WithdrawalRequested(uint256 indexed requestId, address indexed requestor, address indexed owner, uint256 amountOfStETH, uint256 amountOfShares)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseWithdrawalRequested(log types.Log) (*StakedETHWithdrawalNFTWithdrawalRequested, error) {
	event := new(StakedETHWithdrawalNFTWithdrawalRequested)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakedETHWithdrawalNFTWithdrawalsFinalizedIterator is returned from FilterWithdrawalsFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalsFinalized events raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalsFinalizedIterator struct {
	Event *StakedETHWithdrawalNFTWithdrawalsFinalized // Event containing the contract specifics and raw log

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
func (it *StakedETHWithdrawalNFTWithdrawalsFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakedETHWithdrawalNFTWithdrawalsFinalized)
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
		it.Event = new(StakedETHWithdrawalNFTWithdrawalsFinalized)
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
func (it *StakedETHWithdrawalNFTWithdrawalsFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakedETHWithdrawalNFTWithdrawalsFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakedETHWithdrawalNFTWithdrawalsFinalized represents a WithdrawalsFinalized event raised by the StakedETHWithdrawalNFT contract.
type StakedETHWithdrawalNFTWithdrawalsFinalized struct {
	From              *big.Int
	To                *big.Int
	AmountOfETHLocked *big.Int
	SharesToBurn      *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsFinalized is a free log retrieval operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) FilterWithdrawalsFinalized(opts *bind.FilterOpts, from []*big.Int, to []*big.Int) (*StakedETHWithdrawalNFTWithdrawalsFinalizedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.FilterLogs(opts, "WithdrawalsFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakedETHWithdrawalNFTWithdrawalsFinalizedIterator{contract: _StakedETHWithdrawalNFT.contract, event: "WithdrawalsFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsFinalized is a free log subscription operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) WatchWithdrawalsFinalized(opts *bind.WatchOpts, sink chan<- *StakedETHWithdrawalNFTWithdrawalsFinalized, from []*big.Int, to []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakedETHWithdrawalNFT.contract.WatchLogs(opts, "WithdrawalsFinalized", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakedETHWithdrawalNFTWithdrawalsFinalized)
				if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalsFinalized", log); err != nil {
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

// ParseWithdrawalsFinalized is a log parse operation binding the contract event 0x197874c72af6a06fb0aa4fab45fd39c7cb61ac0992159872dc3295207da7e9eb.
//
// Solidity: event WithdrawalsFinalized(uint256 indexed from, uint256 indexed to, uint256 amountOfETHLocked, uint256 sharesToBurn, uint256 timestamp)
func (_StakedETHWithdrawalNFT *StakedETHWithdrawalNFTFilterer) ParseWithdrawalsFinalized(log types.Log) (*StakedETHWithdrawalNFTWithdrawalsFinalized, error) {
	event := new(StakedETHWithdrawalNFTWithdrawalsFinalized)
	if err := _StakedETHWithdrawalNFT.contract.UnpackLog(event, "WithdrawalsFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
