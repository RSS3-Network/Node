package savm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
)

// SAVMBridge
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/SAVMBridge.abi --pkg savm --type SAVMBridge --out contract_savm_bridge.go
// BTCBridge
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/BTCBridge.abi --pkg savm --type BTCBridge --out contract_btc_bridge.go

var (
	AddressSAVMBridge = common.HexToAddress("0x1109F6221F684BCb9B2529b8803a7b8c3411d45f")
	AddressBTCBridge  = common.HexToAddress("0xF70Af817B07118CBF7ACCC38767899598e045408")
	AddressSAVMToken  = common.HexToAddress("0x0E02765992f946397E6d2e65642eABb9cc674928")
	AddressWBTCToken  = common.HexToAddress("0x5db252ead05C54B08A83414adCAbF46Eaa9E0337")

	EventHashBTCBridgeDeposit  = contract.EventHash("Deposit(bytes32,uint256,address)")
	EventHashBTCBridgeWithdraw = contract.EventHash("Withdraw(bytes32,uint256,address)")
	EventHashSAVMTransfer      = contract.EventHash("Transfer(address,address,uint256)")
)
