package nouns

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// NounsAuctionHouseProxy https://etherscan.io/address/0x830BD73E4184ceF73443C15111a1DF14e495C706
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NounsAuctionHouse.abi --pkg nouns --type NounsAuctionHouse --out nouns_auction_house.go

// NounsDAOProxy https://etherscan.io/address/0x6f3E6272A167e8AcCb32072d08E0957F9c79223d
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NounsDAO.abi --pkg nouns --type NounsDAO --out nouns_dao.go

// NounsToken https://etherscan.io/address/0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NounsToken.abi --pkg nouns --type NounsToken --out nouns_token.go

var (
	AddressNouns        = common.HexToAddress("0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03")
	AddressNounsAuction = common.HexToAddress("0x830BD73E4184ceF73443C15111a1DF14e495C706")
	AddressNounsDAO     = common.HexToAddress("0x6f3E6272A167e8AcCb32072d08E0957F9c79223d")

	EventAuctionBid     = contract.EventHash("AuctionBid(uint256,address,uint256,bool)")
	EventAuctionSettled = contract.EventHash("AuctionSettled(uint256,address,uint256)")
	EventAuctionCreated = contract.EventHash("AuctionCreated(uint256,uint256,uint256)")
	EventNounCreated    = contract.EventHash("NounCreated(uint256,(uint48,uint48,uint48,uint48,uint48))")
	EventNounsProposal  = contract.EventHash("ProposalCreatedWithRequirements(uint256,address,address[],uint256[],string[],bytes[],uint256,uint256,uint256,uint256,string)")
	EventNounsVote      = contract.EventHash("VoteCast(address,uint256,uint8,uint256,string)")
)
