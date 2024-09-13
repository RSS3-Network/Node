package vsl

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

// L1StandardBridge https://etherscan.io/address/0x99C9fc46f92E8a1c0deC1b1747d010903E884bE1
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L1StandardBridge.abi --pkg vsl --type L1StandardBridge --out contract_l1_standard_bridge.go
// L2StandardBridge https://optimistic.etherscan.io/address/0x4200000000000000000000000000000000000010
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/L2StandardBridge.abi --pkg vsl --type L2StandardBridge --out contract_l2_standard_bridge.go
// NetworkParams https://scan.rss3.io/address/0x15176Aabdc4836c38947a67313d209204051C502
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/NetworkParams.abi --pkg vsl --type NetworkParams --out contract_network_params.go
// Settlement https://scan.rss3.io/address/0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./abi/Settlement.abi --pkg vsl --type Settlement --out contract_settlement.go

const (
	ChainIDMainnet = 12553
	ChainIDTestnet = 2331
)

var (
	AddressL1RSS3           = common.HexToAddress("0xc98D64DA73a6616c42117b582e832812e7B8D57F")
	AddressL2RSS3           = common.HexToAddress("0x4200000000000000000000000000000000000042")
	AddressL1StandardBridge = common.HexToAddress("0x4cbab69108Aa72151EDa5A3c164eA86845f18438")
	AddressL2StandardBridge = common.HexToAddress("0x4200000000000000000000000000000000000010")
	AddressL1OptimismPortal = common.HexToAddress("0x6A12432491bbbE8d3babf75F759766774C778Db4")

	AddressNetworkParams = map[int64]common.Address{
		ChainIDMainnet: common.HexToAddress("0x15176Aabdc4836c38947a67313d209204051C502"),
		ChainIDTestnet: common.HexToAddress("0x5d768cAef86d3DA8eC6009eE4B3d9b7Fe26A43CB"),
	}

	AddressSettlement = map[int64]common.Address{
		ChainIDMainnet: common.HexToAddress("0x0cE3159BF19F3C55B648D04E8f0Ae1Ae118D2A0B"),
		ChainIDTestnet: common.HexToAddress("0xA37a6Ef0c3635824be2b6c87A23F6Df5d0E2ba1b"),
	}

	EventHashAddressL1StandardBridgeETHDepositInitiated      = contract.EventHash("ETHDepositInitiated(address,address,uint256,bytes)")
	EventHashAddressL1StandardBridgeERC20DepositInitiated    = contract.EventHash("ERC20DepositInitiated(address,address,address,address,uint256,bytes)")
	EventHashAddressL1StandardBridgeETHWithdrawalFinalized   = contract.EventHash("ETHWithdrawalFinalized(address,address,uint256,bytes)")
	EventHashAddressL1StandardBridgeERC20WithdrawalFinalized = contract.EventHash("ERC20WithdrawalFinalized(address,address,address,address,uint256,bytes)")
	EventHashAddressL2StandardBridgeWithdrawalInitiated      = contract.EventHash("WithdrawalInitiated(address,address,address,address,uint256,bytes)")
	EventHashAddressL2StandardBridgeDepositFinalized         = contract.EventHash("DepositFinalized(address,address,address,address,uint256,bytes)")
)
