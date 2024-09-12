package polymarket

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// CTF Exchange
// https://polygonscan.com/address/0x4bfb41d5b3570defd03c39a9a4d8de6bd8b8982e
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/CTFExchange.abi --pkg polymarket --type CTFExchange --out contract_ctf_exchange.go
// Neg Risk CTF Exchange
// https://polygonscan.com/address/0xc5d563a36ae78145c45a50134d48a1215220f80a
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NegRiskCTFExchange.abi --pkg polymarket --type NegRiskCTFExchange --out contract_neg_risk_ctf_exchange.go
// Condition Tokens
// https://polygonscan.com/address/0x4D97DCd97eC945f40cF65F87097ACe5EA0476045
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/ConditionTokens.abi --pkg polymarket --type ConditionTokens --out contract_condition_tokens.go

var (
	AddressPolyMarketCTFExchange        = common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E")
	AddressPolyMarketNegRiskCTFExchange = common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a")
	AddressPolyMarketConditionTokens    = common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045")

	// EventOrderMatched   = contract.EventHash("OrdersMatched(bytes32,address,uint256,uint256,uint256,uint256)")
	EventOrderFinalized = contract.EventHash("OrderFilled(bytes32,address,address,uint256,uint256,uint256,uint256,uint256)")
)
