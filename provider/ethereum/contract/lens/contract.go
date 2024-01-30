package lens

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

var (
	BlockNumberLensV2 int64 = 49424881

	AddressLensProtocol               = common.HexToAddress("0xdb46d1dc155634fbc732f92e853b10b288ad5a1d")
	AddressV1ProfileCreationProxy     = common.HexToAddress("0x1eec6eccaa4625da3fa6cd6339dbcc2418710e8a")
	AddressLensPeriphery              = common.HexToAddress("0xeff187b4190E551FC25a7fA4dFC6cf7fDeF7194f")
	AddressV2LensHandle               = common.HexToAddress("0xe7E7EaD361f3AaCD73A61A9bD6C10cA17F38E945")
	AddressV2ProfileCreationProxy     = common.HexToAddress("0xdCB72aaB62d52aBC2E6be99BEEe535C2D1361fc0")
	AddressV2ProfileHandleRegistry    = common.HexToAddress("0xD4F2F33680FCCb36748FA9831851643781608844")
	AddressV2CollectPublicationAction = common.HexToAddress("0x0D90C58cBe787CD70B5Effe94Ce58185D72143fB")
	AddressProxyAction                = []common.Address{
		common.HexToAddress("0x5a84eC20F88e94dC3EB96cE77695997f8446a22D"),
		common.HexToAddress("0xdd3f6c22ecc68007cc9f76da18984995da4b7b82"),
		common.HexToAddress("0x772c1a3ae4e5425b59c89f2bd6e30228fe734bef"),
		common.HexToAddress("0xf8d491b0a732c910c5acf97e04c6cde41339f70a"),
		common.HexToAddress("0xcbea63064afbfab509c33f9843fd8e08336d5971"),
	}

	EventHashV1PostCreated           = contract.EventHash("PostCreated(uint256,uint256,string,address,bytes,address,bytes,uint256)")
	EventHashV1ProfileCreated        = contract.EventHash("ProfileCreated(uint256,address,address,string,string,address,bytes,string,uint256)")
	EventHashV1ProfileSet            = contract.EventHash("ProfileMetadataSet(uint256,string,uint256)")
	EventHashV1ProfileImageURISet    = contract.EventHash("ProfileImageURISet(uint256,string,uint256)")
	EventHashV1CommentCreated        = contract.EventHash("CommentCreated(uint256,uint256,string,uint256,uint256,bytes,address,bytes,address,bytes,uint256)")
	EventHashV1MirrorCreated         = contract.EventHash("MirrorCreated(uint256,uint256,uint256,uint256,bytes,address,bytes,uint256)")
	EventHashV1CollectNFTTransferred = contract.EventHash("CollectNFTTransferred(uint256,uint256,uint256,address,address,uint256)")

	EventHashV2PostCreated    = contract.EventHash("PostCreated((uint256,string,address[],bytes[],address,bytes),uint256,bytes[],bytes,address,uint256)")
	EventHashV2ProfileSet     = contract.EventHash("ProfileMetadataSet(uint256,string,address,uint256)")
	EventHashV2CommentCreated = contract.EventHash("CommentCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes),uint256,bytes,bytes[],bytes,address,uint256)")
	EventHashV2MirrorCreated  = contract.EventHash("MirrorCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes),uint256,bytes,address,uint256)")
	EventHashV2QuoteCreated   = contract.EventHash("QuoteCreated((uint256,string,uint256,uint256,uint256[],uint256[],bytes,address[],bytes[],address,bytes),uint256,bytes,bytes[],bytes,address,uint256)")
	EventHashV2Collected      = contract.EventHash("Collected(uint256,uint256,uint256,address,bytes,bytes,address,uint256,address,uint256)")
	EventHashV2ProfileCreated = contract.EventHash("ProfileCreated(uint256,address,address,uint256)")
)
