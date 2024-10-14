package rainbow

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Router https://etherscan.io/address/0x00000000009726632680fb29d3f7a9734e3010e2
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Router.abi --pkg rainbow --type Router --out contract_router.go

var (
	AddressRouter = common.HexToAddress("0x00000000009726632680FB29d3F7A9734E3010E2")

	MethodIDRouterFillQuoteEthToToken   = contract.MethodID("fillQuoteEthToToken(address,address,bytes,uint256)")
	MethodIDRouterFillQuoteTokenToEth   = contract.MethodID("fillQuoteTokenToEth(address,address,bytes,uint256,uint256)")
	MethodIDRouterFillQuoteTokenToToken = contract.MethodID("fillQuoteTokenToToken(address,address,address,bytes,uint256,uint256)")
)

type RouterFillQuoteEthToTokenInput struct {
	BuyTokenAddress common.Address
	Target          common.Address
	SwapCallData    []byte
	FeeAmount       *big.Int
}

type RouterFillQuoteTokenToEthInput struct {
	SellTokenAddress         common.Address
	Target                   common.Address
	SwapCallData             []byte
	SellAmount               *big.Int
	FeePercentageBasisPoints *big.Int
}

type RouterFillQuoteTokenToTokenInput struct {
	SellTokenAddress common.Address
	BuyTokenAddress  common.Address
	Target           common.Address
	SwapCallData     []byte
	SellAmount       *big.Int
	FeeAmount        *big.Int
}
