package proxy

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/common/sync"
	"github.com/rss3-network/node/provider/ethereum"
)

var (
	// SlotERC1967 key from https://eips.ethereum.org/EIPS/eip-1967#:~:text=0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc
	SlotERC1967 = common.HexToHash("0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc")
	// SlotOpenZeppelin key from https://github.com/OpenZeppelin/openzeppelin-labs/blob/54ad91472fdd0ac4c34aa97d3a3da45c28245510/initializer_with_sol_editing/contracts/UpgradeabilityProxy.sol#L24
	SlotOpenZeppelin = common.HexToHash("0x7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3") // https://github.com/OpenZeppelin/openzeppelin-labs/blob/54ad91472fdd0ac4c34aa97d3a3da45c28245510/initializer_with_sol_editing/contracts/UpgradeabilityProxy.sol#L19-L24

	slots = []common.Hash{
		SlotERC1967,
		SlotOpenZeppelin,
	}
)

var ErrorNotAProxyContract = errors.New("not a proxy contract")

// GetImplementation returns the implementation address of the proxy contract,
// it used sync.QuickGroup to get the result as soon as possible.
func GetImplementation(ctx context.Context, address common.Address, blockNumber *big.Int, ethereumClient ethereum.Client) (*common.Address, error) {
	quickGroup := sync.NewQuickGroup[*common.Address](ctx)

	for _, slot := range slots {
		slot := slot

		quickGroup.Go(func(ctx context.Context) (*common.Address, error) {
			data, err := ethereumClient.StorageAt(ctx, address, slot, blockNumber)
			if err != nil {
				return nil, fmt.Errorf("get storage at slot %s: %w", slot.String(), err)
			}

			result := common.BytesToAddress(data)

			if result == ethereum.AddressGenesis {
				return nil, fmt.Errorf("invalid result %s", common.Bytes2Hex(data))
			}

			return &result, nil
		})
	}

	result, err := quickGroup.Wait()
	if err != nil {
		// If all slots is empty, then it's not a proxy contract.
		if errors.Is(err, sync.ErrorNoResult) {
			return nil, ErrorNotAProxyContract
		}

		return nil, err
	}

	return result, nil
}
