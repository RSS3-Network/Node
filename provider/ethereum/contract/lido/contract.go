package lido

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// StakedETH https://etherscan.io/address/0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/StakedETH.abi --pkg lido --type StakedETH --out contract_staked_eth.go
// WrappedStakedETH https://etherscan.io/address/0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/WrappedStakedETH.abi --pkg lido --type WrappedStakedETH --out contract_wrapped_staked_eth.go
// StakedETHWithdrawalNFT https://etherscan.io/address/0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/StakedETHWithdrawalNFT.abi --pkg lido --type StakedETHWithdrawalNFT --out contract_staked_eth_withdrawal_nft.go
// StakedMATIC https://etherscan.io/address/0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/StakedMATIC.abi --pkg lido --type StakedMATIC --out contract_staked_matic.go

var (
	AddressStakedETH                = common.HexToAddress("0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84")
	AddressWrappedStakedETH         = common.HexToAddress("0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0")
	AddressStakedETHWithdrawalNFT   = common.HexToAddress("0x889edC2eDab5f40e902b864aD4d7AdE8E412F9B1")
	AddressStakedMATIC              = common.HexToAddress("0x9ee91F9f426fA633d227f7a9b000E28b9dfd8599")
	AddressStakedMATICWithdrawalNFT = common.HexToAddress("0x60a91E2B7A1568f0848f3D43353C453730082E46")
	AddressMATIC                    = common.HexToAddress("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0")

	EventHashStakedETHSubmitted                        = contract.EventHash("Submitted(address,uint256,address)")
	EventHashStakedETHWithdrawalNFTWithdrawalRequested = contract.EventHash("WithdrawalRequested(uint256,address,address,uint256,uint256)")
	EventHashStakedETHWithdrawalNFTWithdrawalClaimed   = contract.EventHash("WithdrawalClaimed(uint256,address,address,uint256)")
	EventHashStakedMATICSubmitEvent                    = contract.EventHash("SubmitEvent(address,uint256,address)")
	EventHashStakedMATICRequestWithdrawEvent           = contract.EventHash("RequestWithdrawEvent(address,uint256,address)")
	EventHashStakedMATICClaimTokensEvent               = contract.EventHash("ClaimTokensEvent(address,uint256,uint256,uint256)")
	EventHashStakedETHTransferShares                   = contract.EventHash("TransferShares(address,address,uint256)")
)
