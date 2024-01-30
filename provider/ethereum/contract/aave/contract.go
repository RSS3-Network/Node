package aave

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// V1LendingPool
// https://etherscan.io/address/0xC1eC30dfD855c287084Bf6e14ae2FDD0246Baf0d
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V1LendingPool.abi --pkg aave --type V1LendingPool --out contract_v1_lendingpool.go

// V2LendingPool
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V2LendingPool.abi --pkg aave --type V2LendingPool --out contract_v2_lendingpool.go

// V3Pool
// https://etherscan.io/address/0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/V3Pool.abi --pkg aave --type V3Pool --out contract_v3_pool.go

var (
	AddressV1LendingPool          = common.HexToAddress("0x398eC7346DcD622eDc5ae82352F02bE94C62d119")
	AddressV2LendingPoolMainnet   = common.HexToAddress("0x7d2768dE32b0b80b7a3454c06BdAc94A69DDc7A9")
	AddressV2LendingPoolPolygon   = common.HexToAddress("0x8dFf5E27EA6b7AC08EbFdf9eB090F32ee9a30fcf")
	AddressV2LendingPoolAvalanche = common.HexToAddress("0x4F01AeD16D97E3aB5ab2B501154DC9bb0F1A5A2C")
	AddressV3PoolMainnet          = common.HexToAddress("0x87870Bca3F3fD6335C3F4ce8392D69350B4fA4E2")
	AddressV3PoolBase             = common.HexToAddress("0xA238Dd80C259a72e81d7e4664a9801593F98d1c5")
	AddressV3PoolOthers           = common.HexToAddress("0x794a61358D6845594F94dc1DB02A252b5b4814aD")

	EventHashV1LendingPoolDeposit  = contract.EventHash("Deposit(address,address,uint256,uint16,uint256)")
	EventHashV1LendingPoolBorrow   = contract.EventHash("Borrow(address,address,uint256,uint256,uint256,uint256,uint256,uint16,uint256)")
	EventHashV1LendingPoolRepay    = contract.EventHash("Repay(address,address,address,uint256,uint256,uint256,uint256)")
	EventHashV2LendingPoolDeposit  = contract.EventHash("Deposit(address,address,address,uint256,uint16)")
	EventHashV2LendingPoolWithdraw = contract.EventHash("Withdraw(address,address,address,uint256)")
	EventHashV2LendingPoolBorrow   = contract.EventHash("Borrow(address,address,address,uint256,uint256,uint256,uint16)")
	EventHashV2LendingPoolRepay    = contract.EventHash("Repay(address,address,address,uint256)")
	EventHashV3PoolSupply          = contract.EventHash("Supply(address,address,address,uint256,uint16)")
	EventHashV3PoolWithdraw        = contract.EventHash("Withdraw(address,address,address,uint256)")
	EventHashV3PoolBorrow          = contract.EventHash("Borrow(address,address,address,uint256,uint8,uint256,uint16)")
	EventHashV3PoolRepay           = contract.EventHash("Repay(address,address,address,uint256,bool)")
)
