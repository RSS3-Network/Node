package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// V1Exchange
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V1Factory.abi --pkg uniswap --type V1Factory --out contract_v1_factory.go
// V1Factory https://etherscan.io/address/0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V1Exchange.abi --pkg uniswap --type V1Exchange --out contract_v1_exchange.go
// V2Factory https://etherscan.io/address/0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V2Factory.abi --pkg uniswap --type V2Factory --out contract_v2_factory.go
// V2Pair
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V2Pair.abi --pkg uniswap --type V2Pair --out contract_v2_pair.go
// V3Factory https://etherscan.io/address/0x1F98431c8aD98523631AE4a59f267346ea31F984
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V3Factory.abi --pkg uniswap --type V3Factory --out contract_v3_factory.go
// V3Pool
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V3Pool.abi --pkg uniswap --type V3Pool --out contract_v3_pool.go
// NonfungiblePositionManager https://etherscan.io/address/0xc36442b4a4522e871399cd717abdd847ab11fe88#code
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NonfungiblePositionManager.abi --pkg uniswap --type NonfungiblePositionManager --out contract_nonfungible_position_manager.go

// https://docs.uniswap.org/
var (
	AddressV1Factory                             = common.HexToAddress("0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95")
	AddressV2Migrator                            = common.HexToAddress("0x16D4F26C15f3658ec65B1126ff27DD3dF2a2996b")
	AddressV2SwapRouter01                        = common.HexToAddress("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a")
	AddressV2SwapRouter02                        = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	AddressV3Migrator                            = common.HexToAddress("0xa5644e29708357803b5a882d272c41cc0df92b34")
	AddressV3SwapRouter                          = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	AddressV3SwapRouterCelo                      = common.HexToAddress("0x5615CDAb10dc425a742d643d949a7F474C01abc4")
	AddressV3SwapRouterRSS3Testnet               = common.HexToAddress("0xe8A7aAbb64B446505c7722c4347C6Cc67E0A9a8d")
	AddressV3SwapRouter02                        = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	AddressV3SwapRouter02Celo                    = common.HexToAddress("0x5615CDAb10dc425a742d643d949a7F474C01abc4")
	AddressV3SwapRouter02BinanceSmartChain       = common.HexToAddress("0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2")
	AddressV2Factory                             = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	AddressV3Factory                             = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressV3FactoryCelo                         = common.HexToAddress("0xAfE208a311B21f13EF87E33A90049fC17A7acDEc")
	AddressV3FactoryBinanceSmartChain            = common.HexToAddress("0xdB1d10011AD0Ff90774D0C6Bb92e5C5c8b4461F7")
	AddressV3FactoryRSS3Testnet                  = common.HexToAddress("0xCA9a021A7434dA20f7C260325856E1478148dF67")
	AddressUniversalRouter                       = common.HexToAddress("0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B")
	AddressNonfungiblePositionManager            = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	AddressNonfungiblePositionManagerRSS3Testnet = common.HexToAddress("0x0F1d392Fc4a2652dc8224E0BB919643AB3d67370")

	EventHashV1ExchangeTokenPurchase                     = contract.EventHash("TokenPurchase(address,uint256,uint256)")
	EventHashV1ExchangeEthPurchase                       = contract.EventHash("EthPurchase(address,uint256,uint256)")
	EventHashV1ExchangeAddLiquidity                      = contract.EventHash("AddLiquidity(address,uint256,uint256)")
	EventHashV1ExchangeRemoveLiquidity                   = contract.EventHash("RemoveLiquidity(address,uint256,uint256)")
	EventHashV2FactoryPairCreated                        = contract.EventHash("PairCreated(address,address,address,uint256)")
	EventHashV3FactoryPoolCreated                        = contract.EventHash("PoolCreated(address,address,uint24,int24,address)")
	EventHashV2PairSwap                                  = contract.EventHash("Swap(address,uint256,uint256,uint256,uint256,address)")
	EventHashV2PairMint                                  = contract.EventHash("Mint(address,uint256,uint256)")
	EventHashV2PairBurn                                  = contract.EventHash("Burn(address,uint256,uint256,address)")
	EventHashV3PoolSwap                                  = contract.EventHash("Swap(address,address,int256,int256,uint160,uint128,int24)")
	EventHashNonfungiblePositionManagerCollect           = contract.EventHash("Collect(uint256,address,uint256,uint256)")
	EventHashNonfungiblePositionManagerDecreaseLiquidity = contract.EventHash("DecreaseLiquidity(uint256,uint128,uint256,uint256)")
	EventHashNonfungiblePositionManagerIncreaseLiquidity = contract.EventHash("IncreaseLiquidity(uint256,uint128,uint256,uint256)")
)
