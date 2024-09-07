package arbitrum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Bridge
// https://etherscan.io/address/0x8315177ab297ba92a06054ce80a67ed4dbd7ed3a
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Bridge.abi --pkg arbitrum --type Bridge --out contract_bridge.go
// Inbox
// https://etherscan.io/address/0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Inbox.abi --pkg arbitrum --type Inbox --out contract_inbox.go
// L1CustomGateway
// https://etherscan.io/address/0xcEe284F754E854890e311e3280b767F80797180d
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1CustomGateway.abi --pkg arbitrum --type L1CustomGateway --out contract_l1_custom_gateway.go
// L1ERC20Gateway
// https://etherscan.io/address/0xa3A7B6F88361F48403514059F1F16C8E78d60EeC
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1ERC20Gateway.abi --pkg arbitrum --type L1ERC20Gateway --out contract_l1_erc20_gateway.go
// L1GatewayRouter
// https://arbiscan.io/address/0x5288c571fd7ad117bea99bf60fe0846c4e84f933
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1GatewayRouter.abi --pkg arbitrum --type L1GatewayRouter --out contract_l1_gateway_router.go
// L2GatewayRouter
// https://arbiscan.io/address/0x5288c571fd7ad117bea99bf60fe0846c4e84f933
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L2GatewayRouter.abi --pkg arbitrum --type L2GatewayRouter --out contract_l2_gateway_router.go
// L2ReverseCustomGateway
// https://arbiscan.io/address/0xCaD7828a19b363A2B44717AFB1786B5196974D8E
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L2ReverseCustomGateway.abi --pkg arbitrum --type L2ReverseCustomGateway --out contract_l2_reverse_custom_gateway.go
// ArbSys
// https://arbiscan.io/address/0x0000000000000000000000000000000000000064
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/ArbSys.abi --pkg arbitrum --type ArbSys --out contract_arb_sys.go

var (
	AddressARB                    = common.HexToAddress("0x912CE59144191C1204E64559FE8253a0e49E6548")
	AddressBridgeOne              = common.HexToAddress("0x8315177ab297ba92a06054ce80a67ed4dbd7ed3a")
	AddressInboxOne               = common.HexToAddress("0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f")
	AddressL1CustomGatewayOne     = common.HexToAddress("0xcEe284F754E854890e311e3280b767F80797180d")
	AddressL1ERC20GatewayOne      = common.HexToAddress("0xa3A7B6F88361F48403514059F1F16C8E78d60EeC")
	AddressL1GatewayRouterOne     = common.HexToAddress("0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef")
	AddressL2GatewayRouter        = common.HexToAddress("0x5288c571fd7ad117bea99bf60fe0846c4e84f933")
	AddressL2ReverseCustomGateway = common.HexToAddress("0xCaD7828a19b363A2B44717AFB1786B5196974D8E")
	AddressArbSys                 = common.HexToAddress("0x0000000000000000000000000000000000000064")

	EventHashBridgeMessageDelivered                    = contract.EventHash("MessageDelivered(uint256,bytes32,address,uint8,address,bytes32,uint256,uint64)")
	EventHashL1CustomGatewayDepositInitiated           = contract.EventHash("DepositInitiated(address,address,address,uint256,uint256)")
	EventHashL1CustomGatewayWithdrawalFinalized        = contract.EventHash("WithdrawalFinalized(address,address,address,uint256,uint256)")
	EventHashL2ReverseCustomGatewayWithdrawalInitiated = contract.EventHash("WithdrawalInitiated(address,address,address,uint256,uint256,uint256)")
	EventHashL2ReverseCustomGatewayDepositFinalized    = contract.EventHash("DepositFinalized(address,address,address,uint256)")
	EventHashArbSysL2ToL1Tx                            = contract.EventHash("L2ToL1Tx(address,address,uint256,uint256,uint256,uint256,uint256,uint256,bytes)")
)

const (
	L1MessageTypeTransactCall uint8 = 9
	L1MessageTypeETHDeposit   uint8 = 12
)
