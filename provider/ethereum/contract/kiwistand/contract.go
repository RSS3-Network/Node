package kiwistand

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// KIWI
//https://optimistic.etherscan.io/address/0x66747bdC903d17C586fA09eE5D6b54CC85bBEA45
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/kiwi.abi --pkg kiwistand --type Kiwi --out contract_kiwi.go
// ProtocolRewards
// https://optimistic.etherscan.io/address/0x7777777f279eba3d3ad8f4e708545291a6fdba8b
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/ProtocolRewards.abi --pkg kiwistand --type ProtocolRewards --out contract_protocol_rewards.go

var (
	AddressKIWI            = common.HexToAddress("0x66747bdc903d17c586fa09ee5d6b54cc85bbea45")
	AddressProtocolRewards = common.HexToAddress("0x7777777f279eba3d3ad8f4e708545291a6fdba8b")

	EventHashRewardsDeposit = contract.EventHash("RewardsDeposit(address,address,address,address,address,address,uint256,uint256,uint256,uint256,uint256)")
	EventHashTransfer       = contract.EventHash("Transfer(address,address,uint256)")
	EventHashSale           = contract.EventHash("Sale(address,uint256,uint256,uint256)")
	EventHashMintComment    = contract.EventHash("MintComment(address,address,uint256,uint256,string)")
)
