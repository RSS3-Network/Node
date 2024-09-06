package benddao

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/BendExchange.abi --pkg benddao --type BendExchange --out ./contract_bend_exchange.go
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/LendPool.abi --pkg benddao --type LendPool --out ./contract_lend_pool.go

var (
	AddressBendExchange = common.HexToAddress("0x7e832eC8ad6F66E6C9ECE63acD94516Dd7fC537A")
	AddressLendPool     = common.HexToAddress("0x70b97A0da65C15dfb0FFA02aEE6FA36e507C2762")

	EventTakerAsk = contract.EventHash("TakerAsk(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")
	EventTakerBid = contract.EventHash("TakerBid(bytes32,uint256,address,address,address,address,address,uint256,uint256,uint256)")

	EventDeposit   = contract.EventHash("Deposit(address,address,uint256,address,uint16)")
	EventWithdraw  = contract.EventHash("Withdraw(address,address,uint256,address)")
	EventBorrow    = contract.EventHash("Borrow(address,address,uint256,address,uint256,address,uint256,uint256,uint16)")
	EventRepay     = contract.EventHash("Repay(address,address,uint256,address,uint256,address,uint256)")
	EventAuction   = contract.EventHash("Auction(address,address,uint256,address,uint256,address,address,uint256)")
	EventRedeem    = contract.EventHash("Redeem(address,address,uint256,uint256,address,uint256,address,uint256)")
	EventLiquidate = contract.EventHash("Liquidate(address,address,uint256,uint256,address,uint256,address,uint256)")
)
