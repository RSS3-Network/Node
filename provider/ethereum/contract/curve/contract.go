package curve

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Registry Exchange
// https://etherscan.io/address/0x99a58482BD75cbab83b27EC03CA68fF489b5788f
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/RegistryExchange.abi --pkg curve --type RegistryExchange --out contract_registry_exchange.go

// Stable Swap
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/StableSwap.abi --pkg curve --type StableSwap --out contract_stable_swap.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/StableSwapNCoins.abi --pkg curve --type StableSwapNCoins --out contract_stable_swap_n_coins.go

// Liquidity Gauge
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/LiquidityGauge.abi --pkg curve --type LiquidityGauge --out contract_liquidity_gauge.go

var (
	AddressETH                       = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	AddressRegistryExchangeEthereum  = common.HexToAddress("0x99a58482BD75cbab83b27EC03CA68fF489b5788f")
	AddressRegistryExchangeArbitrum  = common.HexToAddress("0x4c2Af2Df2a7E567B5155879720619EA06C5BB15D")
	AddressRegistryExchangeAvalanche = common.HexToAddress("0xBff334F8D5912AC5c4f2c590A2396d1C5d990123")
	AddressRegistryExchangeGnosis    = common.HexToAddress("0xE6358f6a45B502477e83CC1CDa759f540E4459ee")
	AddressRegistryExchangeOptimism  = common.HexToAddress("0x22D710931F01c1681Ca1570Ff016eD42EB7b7c2a")
	AddressRegistryExchangePolygon   = common.HexToAddress("0x2a426b3Bb4fa87488387545f15D01d81352732F9")

	EventHashRegistryExchangeExchangeMultiple         = contract.EventHash("ExchangeMultiple(address,address,address[9],uint256[3][4],address[4],uint256,uint256)")
	EventHashStableSwapTokenExchange                  = contract.EventHash("TokenExchange(address,int128,uint256,int128,uint256)")
	EventHashStableSwapAddLiquidity2Coins             = contract.EventHash("AddLiquidity(address,uint256[2],uint256[2],uint256,uint256)")
	EventHashStableSwapAddLiquidity3Coins4Params      = contract.EventHash("AddLiquidity(address,uint256[3],uint256,uint256)")
	EventHashStableSwapAddLiquidity3Coins             = contract.EventHash("AddLiquidity(address,uint256[3],uint256[3],uint256,uint256)")
	EventHashStableSwapAddLiquidity4Coins             = contract.EventHash("AddLiquidity(address,uint256[4],uint256[4],uint256,uint256)")
	EventHashStableSwapRemoveLiquidity2Coins          = contract.EventHash("RemoveLiquidity(address,uint256[2],uint256[2],uint256)")
	EventHashStableSwapRemoveLiquidity2Coins3Param    = contract.EventHash("RemoveLiquidity(address,uint256[2],uint256)")
	EventHashStableSwapRemoveLiquidity3Coins          = contract.EventHash("RemoveLiquidity(address,uint256[3],uint256[3],uint256)")
	EventHashStableSwapRemoveLiquidity4Coins          = contract.EventHash("RemoveLiquidity(address,uint256[4],uint256[4],uint256)")
	EventHashStableSwapRemoveLiquidityOne             = contract.EventHash("RemoveLiquidityOne(address,uint256,uint256)")
	EventHashStableSwapRemoveLiquidityOneFactory      = contract.EventHash("RemoveLiquidityOne(address,uint256,uint256,uint256)")
	EventHashStableSwapRemoveLiquidityImbalance2Coins = contract.EventHash("RemoveLiquidityImbalance(address,uint256[2],uint256[2],uint256,uint256)")
	EventHashStableSwapRemoveLiquidityImbalance3Coins = contract.EventHash("RemoveLiquidityImbalance(address,uint256[3],uint256[3],uint256,uint256)")
	EventHashStableSwapRemoveLiquidityImbalance4Coins = contract.EventHash("RemoveLiquidityImbalance(address,uint256[4],uint256[4],uint256,uint256)")
	EventHashLiquidityGaugeDeposit                    = contract.EventHash("Deposit(address,uint256)")
	EventHashLiquidityGaugeWithdraw                   = contract.EventHash("Withdraw(address,uint256)")

	MethodIDStableSwapAddLiquidity2Coins             = contract.MethodID("add_liquidity(uint256[2],uint256)")
	MethodIDStableSwapAddLiquidity3Coins             = contract.MethodID("add_liquidity(uint256[3],uint256)")
	MethodIDStableSwapAddLiquidity4Coins             = contract.MethodID("add_liquidity(uint256[4],uint256)")
	MethodIDStableSwapRemoveLiquidity2Coins          = contract.MethodID("remove_liquidity(uint256,uint256[2])")
	MethodIDStableSwapRemoveLiquidity3Coins          = contract.MethodID("remove_liquidity(uint256,uint256[3])")
	MethodIDStableSwapRemoveLiquidity4Coins          = contract.MethodID("remove_liquidity(uint256,uint256[4])")
	MethodIDStableSwapRemoveLiquidityOne             = contract.MethodID("remove_liquidity_one_coin(uint256,int128,uint256)")
	MethodIDStableSwapRemoveLiquidityImbalance2Coins = contract.MethodID("remove_liquidity_imbalance(uint256[2],uint256)")
	MethodIDStableSwapRemoveLiquidityImbalance3Coins = contract.MethodID("remove_liquidity_imbalance(uint256[3],uint256)")
	MethodIDStableSwapRemoveLiquidityImbalance4Coins = contract.MethodID("remove_liquidity_imbalance(uint256[4],uint256)")
)

func IsOfficeRegistryExchange(address common.Address) bool {
	return contract.MatchAddresses(
		address,
		AddressRegistryExchangeEthereum,
		AddressRegistryExchangeArbitrum,
		AddressRegistryExchangeAvalanche,
		AddressRegistryExchangeGnosis,
		AddressRegistryExchangeOptimism,
		AddressRegistryExchangePolygon,
	)
}
