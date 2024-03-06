package oneinch

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Exchange https://etherscan.io/address/0x11111254369792b2Ca5d084aB5eEA397cA8fa48B
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Exchange.abi --pkg oneinch --type Exchange --out contract_exchange.go
// AggregationRouterV2 https://etherscan.io/address/0x111111125434b319222CdBf8C261674aDB56F3ae
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/AggregationRouterV2.abi --pkg oneinch --type AggregationRouterV2 --out contract_aggregation_router_v2.go
// AggregationRouterV3 https://etherscan.io/address/0x11111112542D85B3EF69AE05771c2dCCff4fAa26
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/AggregationRouterV3.abi --pkg oneinch --type AggregationRouterV3 --out contract_aggregation_router_v3.go
// AggregationRouterV4 https://etherscan.io/address/0x1111111254fb6c44bAC0beD2854e76F90643097d
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/AggregationRouterV4.abi --pkg oneinch --type AggregationRouterV4 --out contract_aggregation_router_v4.go
// AggregationRouterV5 https://etherscan.io/address/0x1111111254EEB25477B68fb85Ed929f73A960582
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/AggregationRouterV5.abi --pkg oneinch --type AggregationRouterV5 --out contract_aggregation_router_v5.go

var (
	AddressExchange2           = common.HexToAddress("0x11111254369792b2Ca5d084aB5eEA397cA8fa48B")
	AddressAggregationRouterV2 = common.HexToAddress("0x111111125434b319222CdBf8C261674aDB56F3ae")
	AddressAggregationRouterV3 = common.HexToAddress("0x11111112542D85B3EF69AE05771c2dCCff4fAa26")
	AddressAggregationRouterV4 = common.HexToAddress("0x1111111254fb6c44bAC0beD2854e76F90643097d")
	AddressAggregationRouterV5 = common.HexToAddress("0x1111111254EEB25477B68fb85Ed929f73A960582")
	AddressEther               = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	MethodIDAggregationRouterV4ClipperSwap               = contract.MethodID("clipperSwap(address,address,uint256,uint256)")
	MethodIDAggregationRouterV4ClipperSwapTo             = contract.MethodID("clipperSwapTo(address,address,address,uint256,uint256)")
	MethodIDAggregationRouterV4ClipperSwapToWithPermit   = contract.MethodID("clipperSwapToWithPermit(address,address,address,uint256,uint256,bytes)")
	MethodIDAggregationRouterV4FillOrderRFQ              = contract.MethodID("fillOrderRFQ((uint256,address,address,address,address,uint256,uint256),bytes,uint256,uint256)")
	MethodIDAggregationRouterV4FillOrderRFQTo            = contract.MethodID("fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256),bytes,uint256,uint256,address)")
	MethodIDAggregationRouterV4FillOrderRFQToWithPermit  = contract.MethodID("fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256),bytes,uint256,uint256,address,bytes)")
	MethodIDAggregationRouterV4Swap                      = contract.MethodID("swap(address,(address,address,address,address,uint256,uint256,uint256,bytes),bytes)")
	MethodIDAggregationRouterV4UniswapV3Swap             = contract.MethodID("uniswapV3Swap(uint256,uint256,uint256[])")
	MethodIDAggregationRouterV4UniswapV3SwapCallback     = contract.MethodID("uniswapV3SwapCallback(int256,int256,bytes)")
	MethodIDAggregationRouterV4UniswapV3SwapTo           = contract.MethodID("uniswapV3SwapTo(address,uint256,uint256,uint256[])")
	MethodIDAggregationRouterV4UniswapV3SwapToWithPermit = contract.MethodID("uniswapV3SwapToWithPermit(address,address,uint256,uint256,uint256[],bytes)")
	MethodIDAggregationRouterV4Unoswap                   = contract.MethodID("unoswap(address,uint256,uint256,bytes32[])")
	MethodIDAggregationRouterV4UnoswapWithPermit         = contract.MethodID("unoswapWithPermit(address,uint256,uint256,bytes32[],bytes)")

	MethodIDAggregationRouterV5ClipperSwap               = contract.MethodID("clipperSwap(address,address,address,uint256,uint256,uint256,bytes32,bytes32)")
	MethodIDAggregationRouterV5ClipperSwapTo             = contract.MethodID("clipperSwapTo(address,address,address,address,uint256,uint256,uint256,bytes32,bytes32)")
	MethodIDAggregationRouterV5ClipperSwapToWithPermit   = contract.MethodID("clipperSwapToWithPermit(address,address,address,address,uint256,uint256,uint256,bytes32,bytes32,bytes)")
	MethodIDAggregationRouterV5FillOrder                 = contract.MethodID("fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes),bytes,bytes,uint256,uint256,uint256)")
	MethodIDAggregationRouterV5FillOrderRFQ              = contract.MethodID("fillOrderRFQ((uint256,address,address,address,address,uint256,uint256),bytes,uint256)")
	MethodIDAggregationRouterV5FillOrderRFQCompact       = contract.MethodID("fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256),bytes32,bytes32,uint256)")
	MethodIDAggregationRouterV5FillOrderRFQTo            = contract.MethodID("fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256),bytes,uint256,address)")
	MethodIDAggregationRouterV5FillOrderRFQToWithPermit  = contract.MethodID("fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256),bytes,uint256,address,bytes)")
	MethodIDAggregationRouterV5FillOrderTo               = contract.MethodID("fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes),bytes,bytes,uint256,uint256,uint256,address)")
	MethodIDAggregationRouterV5FillOrderToWithPermit     = contract.MethodID("fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes),bytes,bytes,uint256,uint256,uint256,address,bytes)")
	MethodIDAggregationRouterV5Swap                      = contract.MethodID("swap(address,(address,address,address,address,uint256,uint256,uint256),bytes,bytes)")
	MethodIDAggregationRouterV5UniswapV3Swap             = contract.MethodID("uniswapV3Swap(uint256,uint256,uint256[])")
	MethodIDAggregationRouterV5UniswapV3SwapCallback     = contract.MethodID("uniswapV3SwapCallback(int256,int256,bytes)")
	MethodIDAggregationRouterV5UniswapV3SwapTo           = contract.MethodID("uniswapV3SwapTo(address,uint256,uint256,uint256[])")
	MethodIDAggregationRouterV5UniswapV3SwapToWithPermit = contract.MethodID("uniswapV3SwapToWithPermit(address,address,uint256,uint256,uint256[],bytes)")
	MethodIDAggregationRouterV5Unoswap                   = contract.MethodID("unoswap(address,uint256,uint256,uint256[])")
	MethodIDAggregationRouterV5UnoswapTo                 = contract.MethodID("unoswapTo(address,address,uint256,uint256,uint256[])")
	MethodIDAggregationRouterV5UnoswapToWithPermit       = contract.MethodID("unoswapToWithPermit(address,address,uint256,uint256,uint256[],bytes)")

	EventHashExchangeSwapped            = contract.EventHash("Swapped(address,address,address,uint256,uint256,uint256,uint256)")
	EventHashAggregationRouterV2Swapped = contract.EventHash("Swapped(address,address,address,address,uint256,uint256,uint256,uint256,uint256,address)")
	EventHashAggregationRouterV3Swapped = contract.EventHash("Swapped(address,address,address,address,uint256,uint256)")
)

type AggregationRouterV4ClipperSwapToInput struct {
	Recipient common.Address
	SrcToken  common.Address
	DstToken  common.Address
	Amount    *big.Int
	MinReturn *big.Int
}

type AggregationRouterV4ClipperSwapToWithPermitInput struct {
	AggregationRouterV4ClipperSwapToInput

	Permit *big.Int
}

type AggregationRouterV4SwapInput struct {
	Caller common.Address
	Desc   AggregationRouterV4SwapDescription
	Data   []byte
}

type AggregationRouterV4FillOrderRFQToInput struct {
	Order        OrderRFQLibOrderRFQ
	Signature    []byte
	MakingAmount *big.Int
	TakingAmount *big.Int
	Target       common.Address
}

type AggregationRouterV4FillOrderRFQToWithPermitInput struct {
	AggregationRouterV4FillOrderRFQToInput

	Permit []byte
}

type AggregationRouterV4UniswapV3SwapToInput struct {
	Recipient common.Address
	Amount    *big.Int
	MinReturn *big.Int
	Pools     []*big.Int
}

type AggregationRouterV4UniswapV3SwapToWithPermitInput struct {
	AggregationRouterV4UniswapV3SwapToInput

	Permit []byte
}

type AggregationRouterV5ClipperSwapToInput struct {
	ClipperExchange common.Address
	Recipient       common.Address
	SrcToken        common.Address
	DstToken        common.Address
	InputAmount     *big.Int
	OutputAmount    *big.Int
	GoodUntil       *big.Int
	R               []byte
	VS              []byte
}

type AggregationRouterV5ClipperSwapToWithPermitInput struct {
	AggregationRouterV5ClipperSwapToInput

	Permit []byte
}

type AggregationRouterV5FillOrderRFQToInput struct {
	OrderRFQLibOrderRFQ
	Signature      []byte
	FlagsAndAmount *big.Int
	Target         common.Address
}

type AggregationRouterV5FillOrderRFQToWithPermitInput struct {
	AggregationRouterV5FillOrderRFQToInput

	Permit []byte
}

type AggregationRouterV5FillOrderToInput struct {
	OrderLibOrder

	Signature                    []byte
	Interaction                  []byte
	MakingAmount                 *big.Int
	TakingAmount                 *big.Int
	SkipPermitAndThresholdAmount *big.Int
	Target                       common.Address
}

type AggregationRouterV5FillOrderToWithPermitInput struct {
	AggregationRouterV5FillOrderToInput

	Permit []byte
}

type AggregationRouterV5SwapInput struct {
	Executor common.Address
	Desc     GenericRouterSwapDescription
	Permit   []byte
	Data     []byte
}

type AggregationRouterV5UniswapV3SwapToInput struct {
	Recipient common.Address
	Amount    *big.Int
	MinReturn *big.Int
	Pools     []*big.Int
}

type AggregationRouterV5UniswapV3SwapToWithPermitInput struct {
	AggregationRouterV5UniswapV3SwapToInput

	Permit []byte
}

type AggregationRouterV5UnoswapToInput struct {
	Recipient common.Address
	SrcToken  common.Address
	Amount    *big.Int
	MinReturn *big.Int
	Pools     []*big.Int
}

type AggregationRouterV5UnoswapToWithPermitInput struct {
	AggregationRouterV5UnoswapToInput

	Permit []byte
}
