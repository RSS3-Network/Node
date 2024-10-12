package zerion

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi ./Router.abi --pkg zerion --type Router --out router.go
var (
	AddressRouter      = common.HexToAddress("0xd7F1Dd5D49206349CaE8b585fcB0Ce3D96f1696F")
	AddressNativeToken = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")

	EventHashExecuted = contract.EventHash("Executed(address,uint256,uint256,address,uint256,uint256,uint256,uint256,(uint8,(uint256,address),(uint256,address),address,address,bytes),address)")
)
