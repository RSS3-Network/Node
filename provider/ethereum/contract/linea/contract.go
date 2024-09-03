package linea

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// ZkEvmV2
// https://etherscan.io/address/0xd19d4B5d358258f05D7B411E21A1460D11B0876F
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/ZkEvmV2.abi --pkg linea --type ZKEVMV2 --out contract_zk_evm_v2.go

// TokenBridge
// https://etherscan.io/address/0x051F1D88f0aF5763fB888eC4378b4D8B29ea3319
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/TokenBridge.abi --pkg linea --type TokenBridge --out contract_token_bridge.go

var (
	AddressZKEVMV2     = common.HexToAddress("0xd19d4B5d358258f05D7B411E21A1460D11B0876F")
	AddressTokenBridge = common.HexToAddress("0x051F1D88f0aF5763fB888eC4378b4D8B29ea3319")

	EventHashZKEVMV2MessageSent           = contract.EventHash("MessageSent(address,address,uint256,uint256,uint256,bytes,bytes32)")
	EventHashZKEVMV2MessageClaimed        = contract.EventHash("MessageClaimed(bytes32)")
	EventHashTokenBridgeBridgingInitiated = contract.EventHash("BridgingInitiatedV2(address,address,address,uint256)")
	EventHashTokenBridgeBridgingFinalized = contract.EventHash("BridgingFinalizedV2(address,address,uint256,address)")
	MethodIDZKEVMV2ClaimMessage           = contract.MethodID("claimMessage(address,address,uint256,uint256,address,bytes,uint256)")
)

type ZKEVMV2ClaimMessageInput struct {
	From         common.Address
	To           common.Address
	Fee          *big.Int
	Value        *big.Int
	FeeRecipient common.Address
	Calldata     []byte
	Nonce        *big.Int
}
