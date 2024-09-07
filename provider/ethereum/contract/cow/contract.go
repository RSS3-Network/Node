package cow

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// GPv2Settlement
// https://etherscan.io/address/0x9008D19f58AAbD9eD0D60971565AA8510560ab41
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/GPv2Settlement.abi --pkg cow --type Settlement --out contract_settlement.go

var (
	AddressSettlement = common.HexToAddress("0x9008D19f58AAbD9eD0D60971565AA8510560ab41")
	AddressETH        = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	EventHashSettlementTrade = contract.EventHash("Trade(address,address,address,uint256,uint256,uint256,bytes)")
)
