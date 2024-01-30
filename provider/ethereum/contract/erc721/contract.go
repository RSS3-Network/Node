package erc721

import "github.com/rss3-network/node/provider/ethereum/contract"

// https://eips.ethereum.org/EIPS/eip-721
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC721.abi --pkg erc721 --type ERC721 --out contract_erc721.go

var (
	MethodIDSafeTransferFrom                = contract.MethodID("safeTransferFrom(address,address,uint256)")
	MethodIDSafeTransferFromAndCallReceiver = contract.MethodID("safeTransferFrom(address,address,uint256,bytes)")
	MethodIDApprove                         = contract.MethodID("approve(address,uint256)")
	MethodIDSetApprovalForAll               = contract.MethodID("setApprovalForAll(address,bool)")

	EventHashTransfer       = contract.EventHash("Transfer(address,address,uint256)")
	EventHashApproval       = contract.EventHash("Approval(address,address,uint256)")
	EventHashApprovalForAll = contract.EventHash("ApprovalForAll(address,address,bool)")
)
