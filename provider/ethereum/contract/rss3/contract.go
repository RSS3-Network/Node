package rss3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// Staking https://etherscan.io/address/0x5301cbbedc048abac7e213184132cf982d593563
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Staking.abi --pkg rss3 --type Staking --out contract_staking.go

var (
	AddressToken   = common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")
	AddressStaking = common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")

	EventHashStakingDeposited = contract.EventHash("Deposited(address,uint256,uint256,uint256)")
	EventHashStakingWithdrawn = contract.EventHash("Withdrawn(uint256,address,address,uint256)")
	EventHashRewardsClaimed   = contract.EventHash("RewardsClaimed(address,address,uint256)")
)
