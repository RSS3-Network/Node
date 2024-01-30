package opensea

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// WyvernExchangeV1 https://etherscan.io/address/0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/WyvernExchangeV1.abi --pkg opensea --type WyvernExchangeV1 --out contract_wyvern_exchange_v1.go
// WyvernExchangeV2 https://etherscan.io/address/0x7f268357A8c2552623316e2562D90e642bB538E5
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/WyvernExchangeV2.abi --pkg opensea --type WyvernExchangeV2 --out contract_wyvern_exchange_v2.go
// SeaportV1
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/SeaportV1.abi --pkg opensea --type SeaportV1 --out contract_seaport_v1.go

var (
	AddressWyvernExchangeV1 = common.HexToAddress("0x7Be8076f4EA4A4AD08075C2508e481d6C946D12b")
	AddressWyvernExchangeV2 = common.HexToAddress("0x7f268357A8c2552623316e2562D90e642bB538E5")
	AddressSeaportV1Dot0    = common.HexToAddress("0x00000000006CEE72100D161c57ADA5Bb2be1CA79")
	AddressSeaportV1Dot1    = common.HexToAddress("0x00000000006c3852cbEf3e08E8dF289169EdE581")
	AddressSeaportV1Dot2    = common.HexToAddress("0x00000000000006c7676171937C444f6BDe3D6282")
	AddressSeaportV1Dot3    = common.HexToAddress("0x0000000000000aD24e80fd803C6ac37206a45f15")
	AddressSeaportV1Dot4    = common.HexToAddress("0x00000000000001ad428e4906aE43D8F9852d0dD6")
	AddressSeaportV1Dot5    = common.HexToAddress("0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC")

	EventHashWyvernExchangeV1OrdersMatched = contract.EventHash("OrdersMatched(bytes32,bytes32,address,address,uint256,bytes32)")
	EventHashWyvernExchangeV2OrdersMatched = EventHashWyvernExchangeV1OrdersMatched
	EventHashSeaportV1OrderFulfilled       = contract.EventHash("OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])")
)

const (
	ItemTypeNative  uint8 = 0
	ItemTypeERC20   uint8 = 1
	ItemTypeERC721  uint8 = 2
	ItemTypeERC1155 uint8 = 3
)
