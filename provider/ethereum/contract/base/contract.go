package base

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
)

// OptimismPortal
// https://etherscan.io/address/0x49048044D57e1C92A77f79988d21Fa8fAF74E97e
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/OptimismPortal.abi --pkg base --type OptimismPortal --out contract_optimism_portal.go

// L1StandardBridge
// https://etherscan.io/address/0x3154Cf16ccdb4C6d922629664174b904d80F2C35
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1StandardBridge.abi --pkg base --type L1StandardBridge --out contract_l1_standard_bridge.go

var (
	AddressOptimismPortal         = common.HexToAddress("0x49048044D57e1C92A77f79988d21Fa8fAF74E97e")
	AddressL1StandardBridge       = common.HexToAddress("0x3154Cf16ccdb4C6d922629664174b904d80F2C35")
	AddressL2CrossDomainMessenger = common.HexToAddress("0x4200000000000000000000000000000000000007")

	EventHashOptimismPortalTransactionDeposited           = contract.EventHash("TransactionDeposited(address,address,uint256,bytes)")
	EventHashOptimismPortalTransactionWithdrawalFinalized = contract.EventHash("WithdrawalFinalized(bytes32,bool)")

	EventHashL1StandardBridgeETHDepositInitiated      = contract.EventHash("ETHDepositInitiated(address,address,uint256,bytes)")
	EventHashL1StandardBridgeERC20DepositInitiated    = contract.EventHash("ERC20DepositInitiated(address,address,address,address,uint256,bytes)")
	EventHashL1StandardBridgeETHWithdrawalFinalized   = contract.EventHash("ETHWithdrawalFinalized(address,address,uint256,bytes)")
	EventHashL1StandardBridgeERC20WithdrawalFinalized = contract.EventHash("ERC20WithdrawalFinalized(address,address,address,address,uint256,bytes)")
)
