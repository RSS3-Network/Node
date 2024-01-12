package ens_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	utils "github.com/naturalselectionlabs/rss3-node/internal/engine/worker/contract/ens"
	"github.com/stretchr/testify/require"
)

func TestNamehash(t *testing.T) {
	t.Parallel()

	type arguments struct {
		name string
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      string
	}{
		{
			name: "Empty",
			arguments: arguments{
				name: "",
			},
			want: "0x0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			name: "ETH root",
			arguments: arguments{
				name: "eth",
			},
			want: "0x93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae",
		},
		{
			name: "ETH 1st level domain name",
			arguments: arguments{
				name: "foo.eth",
			},
			want: "0xde9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f",
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, utils.Namehash(testcase.arguments.name), common.HexToHash(testcase.want))
		})
	}
}
