package contract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/naturalselectionlabs/sakuin/common/ethereum/contract"
)

// L1GatewayRouter
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1GatewayRouter.abi --pkg scroll --type L1GatewayRouter --out contract_l1_gateway_router.go
// L2GatewayRouter
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L2GatewayRouter.abi --pkg scroll --type L2GatewayRouter --out contract_l2_gateway_router.go

var (
	AddressL1GatewayRouter = common.HexToAddress("0xF8B1378579659D8F7EE5f3C929c2f3E332E41Fd6")
	AddressL2GatewayRouter = common.HexToAddress("0x4C0926FF5252A435FD19e10ED15e5a249Ba19d79")

	EventHashL1GatewayRouterDepositETH    = contract.EventHash("DepositETH(address,address,uint256,bytes)")
	EventHashL1GatewayRouterDepositERC20  = contract.EventHash("DepositERC20(address,address,address,address,uint256,bytes)")
	EventHashL2GatewayRouterWithdrawETH   = contract.EventHash("WithdrawETH(address,address,uint256,bytes)")
	EventHashL2GatewayRouterWithdrawERC20 = contract.EventHash("WithdrawERC20(address,address,address,address,uint256,bytes)")
	// EventHashFinalizeWithdrawETH   = contract.EventHash("FinalizeWithdrawETH(address,address,uint256,bytes)")
	// EventHashFinalizeWithdrawERC20 = contract.EventHash("FinalizeWithdrawERC20(address,address,address,address,uint256,bytes)")
)
