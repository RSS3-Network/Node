package looksrare

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// exchange https://etherscan.io/address/0x59728544B08AB483533076417FbBB2fD0B17CE3a
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/exchange.abi --pkg looksrare --type exchange --out exchange.go
// exchange v2 https://etherscan.io/address/0x0000000000E655fAe4d56241588680F86E3b2377
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/exchange_v2.abi --pkg looksrare --type exchange_v2 --out exchange_v2.go
// aggregator https://etherscan.io/address/0x00000000005228B791a99a61f36A130d50600106
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/aggregator.abi --pkg looksrare --type aggregator --out aggregator.go
// seaport https://etherscan.io/address/0x000000000055d65008F1dFf7167f24E70DB431F6
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/seaport.abi --pkg looksrare --type seaport --out seaport.go

var (
	// V1
	AddressExchange         = common.HexToAddress("0x59728544B08AB483533076417FbBB2fD0B17CE3a")
	AddressAggregator       = common.HexToAddress("0x00000000005228B791a99a61f36A130d50600106")
	AddressFeeSharingSetter = common.HexToAddress("0x5924A28caAF1cc016617874a2f0C3710d881f3c1")

	EventHashRoyaltyPayment = common.BytesToHash(crypto.Keccak256([]byte("RoyaltyPayment(address,uint256,address,address,uint256)")))
	EventHashTakerBid       = common.BytesToHash(crypto.Keccak256([]byte("TakerBid(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")))
	EventHashTakerAsk       = common.BytesToHash(crypto.Keccak256([]byte("TakerAsk(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")))

	// V2
	AddressExchangeV2           = common.HexToAddress("0x0000000000E655fAe4d56241588680F86E3b2377")
	AddressProtocolFeeRecipient = common.HexToAddress("0x1838De7d4e4e42c8eB7b204A91e28E9fad14F536")

	EventHashTakerBidV2      = common.HexToHash("0x3ee3de4684413690dee6fff1a0a4f92916a1b97d1c5a83cdf24671844306b2e3")
	EventHashTakerAskV2      = common.HexToHash("0x9aaa45d6db2ef74ead0751ea9113263d1dec1b50cea05f0ca2002cb8063564a4")
	EventHashAggregatorSweep = common.HexToHash("0x807273efecfbeb7ae7d3a2189d1ed5a7db80074eed86e7d80b10bb925cd1db73")
	EventHashOrderFulfilled  = common.HexToHash("0x9d9af8e38d66c62e2c12f0225249fd9d721c54b83f48d9352c97c6cacdcb6f31")
)
