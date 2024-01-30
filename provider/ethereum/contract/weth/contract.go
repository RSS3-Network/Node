package weth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/rss3-network/protocol-go/schema/filter"
)

// WETH9 https://etherscan.io/address/0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/WETH9.abi --pkg weth --type WETH9 --out contract_weth9.go

var (
	AddressMainnet     = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	AddressOptimism    = common.HexToAddress("0x4200000000000000000000000000000000000006")
	AddressGnosis      = common.HexToAddress("0x6A023CCd1ff6F2045C3309768eAd9E68F978f6e1")
	AddressPolygon     = common.HexToAddress("0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619")
	AddressFantom      = common.HexToAddress("0xA59982c7A272839cBd93e02Bd8978E9a78189AB5")
	AddressCelo        = common.HexToAddress("0x122013fd7dF1C6F636a5bb8f03108E876548b455")
	AddressArbitrumOne = common.HexToAddress("0x82aF49447D8a07e3bd95BD0d56f35241523fBab1")
	AddressAvalanche   = common.HexToAddress("0x49D5c2BdFfac6CE2BFdB6640F4F80f226bc10bAB")

	EventHashDeposit    = contract.EventHash("Deposit(address,uint256)")
	EventHashWithdrawal = contract.EventHash("Withdrawal(address,uint256)")
)

func IsWETH(chainID uint64, address common.Address) bool {
	switch chainID {
	case uint64(filter.EthereumChainIDMainnet):
		return address == AddressMainnet
	default:
		return false
	}
}
