package ens_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract/ens"
	"github.com/stretchr/testify/require"
)

func TestNameHash(t *testing.T) {
	t.Parallel()

	type arguments struct {
		name string
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      common.Hash
	}{
		{
			name: "vitalik.eth",
			arguments: arguments{
				name: "vitalik.eth",
			},
			want: common.HexToHash("0xee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"),
		},
		{
			name: "planetable.eth",
			arguments: arguments{
				name: "planetable.eth",
			},
			want: common.HexToHash("0x7b26d27abaced593c582148f2eb315e57ff6bf40bbf81da5326b3055c395a7bd"),
		},
		{
			name: "liuguo.eth",
			arguments: arguments{
				name: "liuguo.eth",
			},
			want: common.HexToHash("0xd9fe12401ccda681464d15a7094bf8f7efc16dcc56d4098c466dd7d0c84030c0"),
		},
		{
			name: "kallydev.eth",
			arguments: arguments{
				name: "kallydev.eth",
			},
			want: common.HexToHash("0x8236bd8890989e1af9188cbdad0776c6eb09640f5ab61a70d4171dc066e42f19"),
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			node := ens.NameHash(testcase.arguments.name)

			require.Equal(t, node, testcase.want)
		})
	}
}
