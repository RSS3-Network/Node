package optimism

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/rss3-node/provider/ethereum/contract"
)

// L1StandardBridge https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1StandardBridge.abi --pkg optimism --type L1StandardBridge --out contract_l1_standard_bridge.go
// L2StandardBridge https://optimistic.etherscan.io/address/0x4200000000000000000000000000000000000010
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L2StandardBridge.abi --pkg optimism --type L2StandardBridge --out contract_l2_standard_bridge.go

var (
	AddressETH              = common.HexToAddress("0xDeadDeAddeAddEAddeadDEaDDEAdDeaDDeAD0000")
	AddressL1StandardBridge = common.HexToAddress("0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1")
	AddressL2StandardBridge = common.HexToAddress("0x4200000000000000000000000000000000000010")

	EventHashAddressL1StandardBridgeETHDepositInitiated   = contract.EventHash("ETHDepositInitiated(address,address,uint256,bytes)")
	EventHashAddressL1StandardBridgeERC20DepositInitiated = contract.EventHash("ERC20DepositInitiated(address,address,address,address,uint256,bytes)")
	EventHashAddressL2StandardBridgeWithdrawalInitiated   = contract.EventHash("WithdrawalInitiated(address,address,address,address,uint256,bytes)")
)
