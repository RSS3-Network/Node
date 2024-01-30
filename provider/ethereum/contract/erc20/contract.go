package erc20

import "github.com/rss3-network/node/provider/ethereum/contract"

// https://eips.ethereum.org/EIPS/eip-20
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC20.abi --pkg erc20 --type ERC20 --out contract_erc20.go

var (
	MethodIDTransfer     = contract.MethodID("transfer(address,uint256)")
	MethodIDTransferFrom = contract.MethodID("transferFrom(address,address,uint256)")
	MethodIDApprove      = contract.MethodID("approve(address,uint256)")

	EventHashTransfer = contract.EventHash("Transfer(address,address,uint256)")
	EventHashApproval = contract.EventHash("Approval(address,address,uint256)")
)
