package contract_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/provider/ethereum/contract"
	"github.com/stretchr/testify/require"
)

func TestMethodID(t *testing.T) {
	t.Parallel()

	type arguments struct {
		value string
	}

	testcases := []struct {
		name      string
		arguments arguments
		want      require.ValueAssertionFunc
		wantError require.ErrorAssertionFunc
	}{
		{
			name: "implementation()",
			arguments: arguments{
				value: "implementation()",
			},
			want: func(t require.TestingT, value interface{}, msgAndArgs ...interface{}) {
				methodID, ok := value.(string)
				require.True(t, ok)

				require.Equal(t, "0x5c60da1b", methodID)
			},
		},
		{
			name: "name()",
			arguments: arguments{
				value: "name()",
			},
			want: func(t require.TestingT, value interface{}, msgAndArgs ...interface{}) {
				methodID, ok := value.(string)
				require.True(t, ok)

				require.Equal(t, "0x06fdde03", methodID)
			},
		},
	}

	for _, testcase := range testcases {
		testcase := testcase

		t.Run(testcase.name, func(t *testing.T) {
			t.Parallel()

			result := contract.MethodID(testcase.arguments.value)
			testcase.want(t, hexutil.Encode(result[:]))
		})
	}
}
