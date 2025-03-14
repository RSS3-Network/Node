package uniswap

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
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
// https://github.com/Uniswap/universal-router/blob/main/deploy-addresses/mainnet.json
// https://docs.uniswap.org/contracts/v3/reference/deployments/
var (
	AddressV1Factory                                   = common.HexToAddress("0xc0a47dFe034B400B47bDaD5FecDa2621de6c4d95")
	AddressV2Migrator                                  = common.HexToAddress("0x16D4F26C15f3658ec65B1126ff27DD3dF2a2996b")
	AddressV3Migrator                                  = common.HexToAddress("0xa5644e29708357803b5a882d272c41cc0df92b34")
	AddressV3MigratorArbitrum                          = common.HexToAddress("0xA5644E29708357803b5A882D272c41cC0dF92B34")
	AddressV3MigratorPolygon                           = common.HexToAddress("0xA5644E29708357803b5A882D272c41cC0dF92B34")
	AddressV3MigratorBase                              = common.HexToAddress("0x23cF10b1ee3AdfCA73B0eF17C07F7577e7ACd2d7")
	AddressV3MigratorBinanceSmartChain                 = common.HexToAddress("0x32681814957e0C13117ddc0c2aba232b5c9e760f")
	AddressV3MigratorAvalanche                         = common.HexToAddress("0x44f5f1f5E452ea8d29C890E8F6e893fC0f1f0f97")
	AddressV3MigratorLinea                             = common.HexToAddress("0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1")
	AddressV2SwapRouter01                              = common.HexToAddress("0xf164fC0Ec4E93095b804a4795bBe1e041497b92a")
	AddressV2SwapRouter02                              = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	AddressV3SwapRouter                                = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	AddressV3SwapRouter02                              = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	AddressV3SwapRouter02Arbitrum                      = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	AddressV3SwapRouter02Optimism                      = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	AddressV3SwapRouter02Polygon                       = common.HexToAddress("0x68b3465833fb72A70ecDF485E0e4C7bD8665Fc45")
	AddressV3SwapRouter02Base                          = common.HexToAddress("0x2626664c2603336E57B271c5C0b26F421741e481")
	AddressV3SwapRouter02BinanceSmartChain             = common.HexToAddress("0xB971eF87ede563556b2ED4b1C0b0019111Dd85d2")
	AddressV3SwapRouter02Avalanche                     = common.HexToAddress("0xbb00FF08d01D300023C629E8fFfFcb65A5a578cE")
	AddressV3SwapRouter02Linea                         = common.HexToAddress("0x3d4e44Eb1374240CE5F1B871ab261CD16335B76a")
	AddressV2Factory                                   = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	AddressV2FactorySAVM                               = common.HexToAddress("0x1842c9bD09bCba88b58776c7995A9A9bD220A925")
	AddressV3Factory                                   = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressV3FactoryArbitrum                           = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressV3FactoryOptimism                           = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressV3FactoryPolygon                            = common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984")
	AddressV3FactoryBase                               = common.HexToAddress("0x33128a8fC17869897dcE68Ed026d694621f6FDfD")
	AddressV3FactoryBinanceSmartChain                  = common.HexToAddress("0xdB1d10011AD0Ff90774D0C6Bb92e5C5c8b4461F7")
	AddressV3FactoryAvalanche                          = common.HexToAddress("0x740b1c1de25031C31FF4fC9A62f554A55cdC1baD")
	AddressV3FactoryLinea                              = common.HexToAddress("0x31FAfd4889FA1269F7a13A66eE0fB458f27D72A9")
	AddressUniversalRouter01                           = common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD")
	AddressUniversalRouter01Arbitrum                   = common.HexToAddress("0x5E325eDA8064b456f4781070C0738d849c824258")
	AddressUniversalRouter01Optimism                   = common.HexToAddress("0xCb1355ff08Ab38bBCE60111F1bb2B784bE25D7e8")
	AddressUniversalRouter01Polygon                    = common.HexToAddress("0xec7BE89e9d109e7e3Fec59c222CF297125FEFda2")
	AddressUniversalRouter01Base                       = common.HexToAddress("0x3fC91A3afd70395Cd496C647d5a6CC9D4B2b7FAD")
	AddressUniversalRouter01BinanceSmartChain          = common.HexToAddress("0x4Dae2f939ACf50408e13d58534Ff8c2776d45265")
	AddressUniversalRouter01Avalanche                  = common.HexToAddress("0x4Dae2f939ACf50408e13d58534Ff8c2776d45265")
	AddressUniversalRouter02                           = common.HexToAddress("0xEf1c6E67703c7BD7107eed8303Fbe6EC2554BF6B")
	AddressNonfungiblePositionManager                  = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	AddressNonfungiblePositionManagerArbitrum          = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	AddressNonfungiblePositionManagerOptimism          = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	AddressNonfungiblePositionManagerPolygon           = common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	AddressNonfungiblePositionManagerBase              = common.HexToAddress("0x03a520b32C04BF3bEEf7BEb72E919cf822Ed34f1")
	AddressNonfungiblePositionManagerBinanceSmartChain = common.HexToAddress("0x7b8A01B39D58278b5DE7e48c8449c9f4F5170613")
	AddressNonfungiblePositionManagerAvalanche         = common.HexToAddress("0x655C406EBFa14EE2006250925e54ec43AD184f8B")
	AddressNonfungiblePositionManagerLinea             = common.HexToAddress("0x4615C383F85D0a2BbED973d83ccecf5CB7121463")
	AddressV2SwapRouterSAVM                            = common.HexToAddress("0xC7c934E224e8567df50058A907904b451bD1c57D")

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
