package erc1155

import "github.com/rss3-network/node/provider/ethereum/contract"

// https://eips.ethereum.org/EIPS/eip-1155
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC1155.abi --pkg erc1155 --type ERC1155 --out contract_erc1155.go

var (
	MethodIDSafeTransferFrom      = contract.MethodID("safeTransferFrom(address,address,uint256,uint256,bytes)")
	MethodIDSafeBatchTransferFrom = contract.MethodID("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)")
	MethodIDSetApprovalForAll     = contract.MethodID("setApprovalForAll(address,bool)")

	EventHashTransferSingle = contract.EventHash("TransferSingle(address,address,address,uint256,uint256)")
	EventHashTransferBatch  = contract.EventHash("TransferBatch(address,address,address,uint256[],uint256[])")
	EventHashApprovalForAll = contract.EventHash("ApprovalForAll(address,address,bool)")
	EventHashURI            = contract.EventHash("URI(string,uint256)")
)
