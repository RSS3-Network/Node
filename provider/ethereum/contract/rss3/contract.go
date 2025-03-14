package rss3

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/v2/provider/ethereum/contract"
)

// Staking https://etherscan.io/address/0x5301cbbedc048abac7e213184132cf982d593563
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Staking.abi --pkg rss3 --type Staking --out contract_staking.go

// Staking VSL https://scan.rss3.io/address/0x28F14d917fddbA0c1f2923C406952478DfDA5578
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/StakingVSL.abi --pkg rss3 --type StakingVSL --out contract_staking_vsl.go

// Settlement https://scan.rss3.io/address/0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Settlement.abi --pkg rss3 --type Settlement --out contract_settlement.go

// Chips https://scan.rss3.io/address/0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Chips.abi --pkg rss3 --type Chips --out contract_chips.go

var (
	AddressToken         = common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")
	AddressStaking       = common.HexToAddress("0x5301CbBeDc048AbaC7e213184132cf982d593563")
	AddressStakingVSL    = common.HexToAddress("0x28F14d917fddbA0c1f2923C406952478DfDA5578")
	AddressSettlementVSL = common.HexToAddress("0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B")
	AddressChipsVSL      = common.HexToAddress("0x849f8F55078dCc69dD857b58Cc04631EBA54E4DE")

	EventHashStakingDeposited = contract.EventHash("Deposited(address,uint256,uint256,uint256)")
	EventHashStakingWithdrawn = contract.EventHash("Withdrawn(uint256,address,address,uint256)")
	EventHashRewardsClaimed   = contract.EventHash("RewardsClaimed(address,address,uint256)")

	EventHashStakingVSLDeposited = contract.EventHash("Deposited(address,uint256)")
	EventHashStakingVSLStaked    = contract.EventHash("Staked(address,address,uint256,uint256,uint256)")
	EventHashTransfer            = contract.EventHash("Transfer(address,address,uint256)")
)
