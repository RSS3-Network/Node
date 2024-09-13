package multicall3

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/protocol-go/schema/network"
	"github.com/samber/lo"
	"github.com/sourcegraph/conc/pool"
)

// Multicall https://github.com/mds1/multicall
// https://etherscan.io/address/0xcA11bde05977b3631167028862bE2a173976CA11
//go:generate go run -mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/Multicall3.abi --pkg multicall3 --type Multicall3 --out contract_multicall3.go

var (
	AddressMulticall3     = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")
	AddressMulticall3SAVM = common.HexToAddress("0xd81C066cbb9f0BE7d33f5e504B62fa4f2D587DD3")
)

var deployedAtMap = map[uint64]uint64{
	uint64(network.EthereumChainIDMainnet):           14353601, // https://etherscan.io/tx/0x00d9fcb7848f6f6b0aae4fb709c133d69262b902156c85a473ef23faa60760bd
	uint64(network.EthereumChainIDOptimism):          4286263,  // https://optimistic.etherscan.io/tx/0xb62f9191a2cf399c0d2afd33f5b8baf7c6b52af6dd2386e44121b1bab91b80e5
	uint64(network.EthereumChainIDPolygon):           25770160, // https://polygonscan.com/tx/0x25d385667b12d6992742127dc7682e570136397806e2773dc47922eba0001989
	uint64(network.EthereumChainIDArbitrum):          7654707,  // https://arbiscan.io/tx/0x211f6689adbb0f3fba7392e899d23bde029cef532cbd0ae900920cc09f7d1f32
	uint64(network.EthereumChainIDBase):              5022,     // https://basescan.org/tx/0x07471adfe8f4ec553c1199f495be97fc8be8e0626ae307281c22534460184ed1
	uint64(network.EthereumChainIDCrossbell):         38246031, // https://scan.crossbell.io/tx/0x0d7367f09c151993f1a7ac32780948fc07d232806e81d0d0c3e7058e4538d7f5
	uint64(network.EthereumChainIDVSL):               14193,    // https://scan.rss3.io/tx/0x07471adfe8f4ec553c1199f495be97fc8be8e0626ae307281c22534460184ed1
	uint64(network.EthereumChainIDSatoshiVM):         60740,    // https://svmscan.io/tx/0x3349d0d487c2ba43d19f7ec6a1e2176e3f72ccbfa647ac719a074682346dd40c
	uint64(network.EthereumChainIDBinanceSmartChain): 15921452, // https://bscscan.com/tx/0xcc0ddf5f791617ba9befce57995dbcb3a202946a1eefa3469742b01a0decdaf2
	uint64(network.EthereumChainIDGnosis):            21022491, // https://gnosis.blockscout.com/tx/0xf528b4398b2961bab9404a943d1da24d6c82b7367b0cc233629fccc94f1ababa
	uint64(network.EthereumChainIDLinea):             42,       // https://lineascan.build/tx/0x369305bf856786a160151d3897de2540fc700cf3d6a289b954ca93460f1038b6
	uint64(network.EthereumChainIDXLayer):            47416,    // https://www.okx.com/web3/explorer/xlayer/tx/0x07471adfe8f4ec553c1199f495be97fc8be8e0626ae307281c22534460184ed1
}

func IsDeployed(chainID uint64, blockNumber *big.Int) bool {
	deployedAt, exists := deployedAtMap[chainID]
	if !exists {
		return false
	}

	if blockNumber == nil {
		return true
	}

	return deployedAt < blockNumber.Uint64()
}

func Aggregate3(ctx context.Context, chainID uint64, calls []Multicall3Call3, blockNumber *big.Int, contractBackend bind.ContractCaller) ([]*Multicall3Result, error) {
	// If the Multicall3 contract is not yet deployed, split it into multiple regular requests.
	if !IsDeployed(chainID, blockNumber) {
		resultPool := pool.NewWithResults[*Multicall3Result]().WithContext(ctx).WithCancelOnError()

		for _, call := range calls {
			call := call

			resultPool.Go(func(ctx context.Context) (*Multicall3Result, error) {
				message := ethereum.CallMsg{
					To:   lo.ToPtr(call.Target),
					Data: call.CallData,
				}

				data, err := contractBackend.CallContract(ctx, message, blockNumber)
				if err != nil && !call.AllowFailure {
					return nil, err
				}

				result := Multicall3Result{
					Success:    err == nil,
					ReturnData: data,
				}

				return &result, nil
			})
		}

		return resultPool.Wait()
	}

	abi, err := Multicall3MetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("load abi: %w", err)
	}

	callData, err := abi.Pack("aggregate3", calls)
	if err != nil {
		return nil, fmt.Errorf("pack data: %w", err)
	}

	var contractAt common.Address

	switch chainID {
	case uint64(network.EthereumChainIDSatoshiVM):
		contractAt = AddressMulticall3SAVM
	default:
		contractAt = AddressMulticall3
	}

	message := ethereum.CallMsg{
		To:   &contractAt,
		Data: callData,
	}

	results := make([]Multicall3Result, 0, len(calls))

	data, err := contractBackend.CallContract(ctx, message, blockNumber)
	if err != nil {
		return nil, fmt.Errorf("call contract: %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data in empty")
	}

	if err := abi.UnpackIntoInterface(&results, "aggregate3", data); err != nil {
		return nil, fmt.Errorf("unpack result: %w", err)
	}

	return lo.ToSlicePtr(results), nil
}
