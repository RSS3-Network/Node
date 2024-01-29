package matters

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/serving-node/provider/ethereum/contract"
)

//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen --abi abi/curation.abi --pkg matters --type Matters --out contract_curation.go

var (
	AddressCuration = common.HexToAddress("0x5edebbdae7B5C79a69AaCF7873796bb1Ec664DB8")

	EventCuration = contract.EventHash("Curation(address,address,address,string,uint256)")
)