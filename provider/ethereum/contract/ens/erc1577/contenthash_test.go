package erc1577_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/provider/ethereum/contract/ens/erc1577"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	t.Parallel()

	type arguments struct {
		contenthash []byte
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      string
		wantError error
	}{
		{
			name: "vitalik.eth",
			arguments: arguments{
				contenthash: lo.Must(hexutil.Decode("0xe3010170122069e1ae368014fdf33ee9b7bc560fbb14ff64772b7769b14f0696b7334f66a765")),
			},
			want: "/ipfs/bafybeidj4gxdnaau7xzt52nxxrla7oyu75shok3xngyu6buww4zu6zvhmu",
		},
		{
			name: "planetable.eth",
			arguments: arguments{
				contenthash: lo.Must(hexutil.Decode("0xe50101720024080112201b78b16d294d40117c9029caeab5e742c6adbed2233bf6369d7d94254faf1f8a")),
			},
			want: "/ipns/bafzaajaiaejcag3ywfwsstkacf6jakok5k26oqwgvw7neiz36y3j27muevh26h4k",
		},
		{
			name: "liuguo.eth",
			arguments: arguments{
				contenthash: lo.Must(hexutil.Decode("0xe50101720024080112203560641da3d672f2de69cef6d49bf01890af66826d8c8a7e8e9e93c172bfc9bd")),
			},
			want: "/ipns/bafzaajaiaejcanlamqo2hvts6lpgttxw2sn7ageqv5tie3mmrj7i5hutyfzl7sn5",
		},
		{
			name: "kallydev.eth",
			arguments: arguments{
				contenthash: lo.Must(hexutil.Decode("0xe50101720024080112208bf98ae9ae12aeca905ae50cd8d99f46bd633c5f702da516433507aefe788650")),
			},
			want: "/ipns/bafzaajaiaejcbc7zrlu24evozkifvzim3dmz6rv5mm6f64bnuulegnihv37hrbsq",
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			path, err := erc1577.Parse(testcase.arguments.contenthash)
			require.ErrorIs(t, err, testcase.wantError)

			require.Equal(t, path, testcase.want)
		})
	}
}
