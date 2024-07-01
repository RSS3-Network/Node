package parameter

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/vsl"
	"github.com/rss3-network/protocol-go/schema/network"
)

// NumberOfMonthsToCover the number of months that a Node should cover data for
const NumberOfMonthsToCover = 3

type NetworkTolerance map[string]uint64
type NetworkStartBlock map[string]*big.Int
type NetworkCoreWorkerDiskSpacePerMonth map[string]uint

// NetworkParamsData contains the network parameters
type NetworkParamsData struct {
	NetworkTolerance                   NetworkTolerance                   `json:"NetworkTolerance"`
	NetworkStartBlock                  NetworkStartBlock                  `json:"NetworkStartBlock"`
	NetworkCoreWorkerDiskSpacePerMonth NetworkCoreWorkerDiskSpacePerMonth `json:"NetworkCoreWorkerDiskSpacePerMonth"`
}

// CurrentNetworkTolerance should be updated each epoch from vsl
var CurrentNetworkTolerance = NetworkTolerance{
	network.Arbitrum.String():          1000,
	network.Arweave.String():           100,
	network.Avalanche.String():         100,
	network.Base.String():              100,
	network.BinanceSmartChain.String(): 100,
	network.Crossbell.String():         500,
	network.Ethereum.String():          100,
	network.Farcaster.String():         3600000,
	network.Gnosis.String():            100,
	network.Linea.String():             100,
	network.Optimism.String():          100,
	network.Polygon.String():           100,
	network.SatoshiVM.String():         100,
	network.VSL.String():               100,
}

// CurrentNetworkStartBlock should be updated each epoch from vsl
var CurrentNetworkStartBlock = NetworkStartBlock{
	network.Arbitrum.String():          big.NewInt(185724972),
	network.Arweave.String():           big.NewInt(1374360),
	network.Avalanche.String():         big.NewInt(42301570),
	network.Base.String():              big.NewInt(11216527),
	network.BinanceSmartChain.String(): big.NewInt(36563564),
	network.Crossbell.String():         big.NewInt(58846671),
	network.Ethereum.String():          big.NewInt(19334220),
	network.Gnosis.String():            big.NewInt(32695982),
	network.Linea.String():             big.NewInt(2591120),
	network.Optimism.String():          big.NewInt(116811812),
	network.Polygon.String():           big.NewInt(54103805),
	network.SatoshiVM.String():         big.NewInt(60741),
	network.VSL.String():               big.NewInt(14192),
}

// CurrentNetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
// The data is calculated based on the average disk space usage during 2024 Q1.
// Actually usage may vary depending on the network's activity.
var CurrentNetworkCoreWorkerDiskSpacePerMonth = NetworkCoreWorkerDiskSpacePerMonth{
	network.Arbitrum.String():          26,
	network.Arweave.String():           0,
	network.Avalanche.String():         0,
	network.Base.String():              10,
	network.BinanceSmartChain.String(): 117,
	network.Crossbell.String():         0,
	network.Ethereum.String():          51,
	network.Gnosis.String():            9,
	network.Linea.String():             31,
	network.Optimism.String():          25,
	network.Polygon.String():           153,
	network.SatoshiVM.String():         1,
	network.VSL.String():               1,
	network.Farcaster.String():         50,
}

// PullNetworkParamsFromVSL pulls the network parameters from the VSL
func PullNetworkParamsFromVSL() error {
	vslClient, err := initVSLClient()
	if err != nil {
		return fmt.Errorf("init vsl client: %w", err)
	}

	networkParams, err := vsl.NewNetworkParamsCaller(vsl.AddressNetworkParams, vslClient)
	if err != nil {
		return fmt.Errorf("new network params caller: %w", err)
	}

	params, err := networkParams.GetParams(&bind.CallOpts{})
	if err != nil {
		return err
	}

	// unmarshal params to update CurrentNetworkTolerance, CurrentNetworkStartBlock and CurrentNetworkCoreWorkerDiskSpacePerMonth
	var paramsData NetworkParamsData

	err = json.Unmarshal([]byte(params), &paramsData)
	if err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}

	// CurrentNetworkTolerance should be updated each epoch from vsl
	CurrentNetworkTolerance = paramsData.NetworkTolerance

	// CurrentNetworkStartBlock should be updated each epoch from vsl
	CurrentNetworkStartBlock = paramsData.NetworkStartBlock

	// CurrentNetworkCoreWorkerDiskSpacePerMonth the disk space required for the network's core worker to store a month worth of data
	// The data is calculated based on the average disk space usage during 2024 Q1.
	// Actually usage may vary depending on the network's activity.
	CurrentNetworkCoreWorkerDiskSpacePerMonth = paramsData.NetworkCoreWorkerDiskSpacePerMonth

	return nil
}

func initVSLClient() (ethereum.Client, error) {
	// TODO should use vsl rpc url in prod
	// vslEndpoint := endpoint.MustGet(network.VSL)
	vslEndpoint := "https://rpc.testnet.rss3.io"

	// Initialize vsl ethereum client.
	vslClient, err := ethereum.Dial(context.Background(), vslEndpoint)
	if err != nil {
		return nil, err
	}

	return vslClient, nil
}
