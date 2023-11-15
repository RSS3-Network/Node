package ethereum

import "github.com/ethereum/go-ethereum/common"

var (
	AddressGenesis = common.HexToAddress("0x0000000000000000000000000000000000000000")
	HashGenesis    = common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000")
)

// IsBurnAddress will return whether the current address is a black hole address,
// most of which are provided by https://etherscan.io/accounts/label/burn.
func IsBurnAddress(address common.Address) bool {
	switch address {
	case
		common.HexToAddress("0x000000000000000000000000000000000000dEaD"),
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		common.HexToAddress("0x0000000000000000000000000000000000000001"),
		common.HexToAddress("0x0000000000000000000000000000000000000002"),
		common.HexToAddress("0x0000000000000000000000000000000000000003"),
		common.HexToAddress("0x0000000000000000000000000000000000000004"),
		common.HexToAddress("0x0000000000000000000000000000000000000005"),
		common.HexToAddress("0x0000000000000000000000000000000000000006"),
		common.HexToAddress("0x0000000000000000000000000000000000000007"),
		common.HexToAddress("0x0000000000000000000000000000000000000008"),
		common.HexToAddress("0x0000000000000000000000000000000000000009"),
		common.HexToAddress("0x00000000000000000000045261D4Ee77acdb3286"),
		common.HexToAddress("0x0123456789012345678901234567890123456789"),
		common.HexToAddress("0x1111111111111111111111111111111111111111"),
		common.HexToAddress("0x1234567890123456789012345678901234567890"),
		common.HexToAddress("0x2222222222222222222222222222222222222222"),
		common.HexToAddress("0x3333333333333333333333333333333333333333"),
		common.HexToAddress("0x4444444444444444444444444444444444444444"),
		common.HexToAddress("0x6666666666666666666666666666666666666666"),
		common.HexToAddress("0x8888888888888888888888888888888888888888"),
		common.HexToAddress("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB"),
		common.HexToAddress("0xdEAD000000000000000042069420694206942069"),
		common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
		common.HexToAddress("0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF"),
		common.HexToAddress("0xaAaAaAaaAaAaAaaAaAAAAAAAAaaaAaAaAaaAaaAa"):
		return true
	default:
		return false
	}
}
