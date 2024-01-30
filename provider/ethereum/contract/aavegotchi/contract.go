package aavegotchi

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC1155MarketplaceFacet.abi --pkg aavegotchi --type ERC1155Marketplace --out ./contract_erc1155_marketplace_facet.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC721MarketplaceFacet.abi --pkg aavegotchi --type ERC721Marketplace --out ./contract_erc721_marketplace_facet.go

var (
	AddressAavegotchi = common.HexToAddress("0x86935F11C86623deC8a25696E1C19a8659CbF95d")

	EventHashERC721ListingAdd       = contract.EventHash("ERC721ListingAdd(uint256,address,address,uint256,uint256,uint256)")
	EventHashERC721ExecutedListing  = contract.EventHash("ERC721ExecutedListing(uint256,address,address,address,uint256,uint256,uint256,uint256)")
	EventHashERC1155ExecutedListing = contract.EventHash("ERC1155ExecutedListing(uint256,address,address,address,uint256,uint256,uint256,uint256,uint256)")
	EventHashERC1155ListingAdd      = contract.EventHash("ERC1155ListingAdd(uint256,address,address,uint256,uint256,uint256,uint256,uint256)")
)
