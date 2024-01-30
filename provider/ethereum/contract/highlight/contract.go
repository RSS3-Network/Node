package highlight

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// exchange https://etherscan.io/address/0x1bf979282181f2b7a640d17ab5d2e25125f2de5e
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/MintManager.abi --pkg highlight --type mint_manager --out mint_manager.go

var (
	AddressMintManagerMainnet  = common.HexToAddress("0x1bf979282181f2b7a640d17ab5d2e25125f2de5e")
	AddressMintManagerPolygon  = common.HexToAddress("0xfbb65C52f439B762F712026CF6DD7D8E82F81eb9")
	AddressMintManagerArbitrum = common.HexToAddress("0x41cbab1028984A34C1338F437C726de791695AE8")
	AddressMintManagerBase     = common.HexToAddress("0x8087039152c472Fa74F47398628fF002994056EA")
	AddressMintManagerOptimism = common.HexToAddress("0xFafd47bb399d570b5AC95694c5D2a1fb5EA030bB")

	EventHashNativeGasTokenPayment = contract.EventHash("NativeGasTokenPayment(address,bytes32,uint256,uint32)")
	EventHashNumTokenMint          = contract.EventHash("NumTokenMint(bytes32,address,bool,uint256)")
)
