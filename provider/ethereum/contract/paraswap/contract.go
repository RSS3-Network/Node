package paraswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
)

// Paraswap V5 https://etherscan.io/address/0xdef171fe48cf0115b1d80b88dc8eab59176fee57
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V5ParaSwap.abi --pkg paraswap --type V5ParaSwap --out contract_v5_paraswap.go

// https://developers.paraswap.network/smart-contracts
var (
	AddressV5ParaSwap                  = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")
	AddressV5ParaSwapArbitrum          = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")
	AddressV5ParaSwapAvalanche         = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")
	AddressV5ParaSwapBase              = common.HexToAddress("0x59C7C832e96D2568bea6db468C1aAdcbbDa08A52")
	AddressV5ParaSwapBinanceSmartChain = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")
	AddressV5ParaSwapOptimism          = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")
	AddressV5ParaSwapPolygon           = common.HexToAddress("0xDEF171Fe48CF0115B1d80b88dc8eAB59176FEe57")

	AddressETH = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	EventHashV3Swapped     = contract.EventHash("SwappedV3(bytes16,address,uint256,address,address,address,address,uint256,uint256,uint256)")
	EventHashV3Bought      = contract.EventHash("BoughtV3(bytes16,address,uint256,address,address,address,address,uint256,uint256,uint256)")
	EventHashSwappedDirect = contract.EventHash("SwappedDirect(bytes16,address,uint256,address,uint8,address,address,address,uint256,uint256,uint256)")
)
