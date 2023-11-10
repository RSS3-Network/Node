package message

import (
	"time"

	"github.com/ethereum/go-ethereum/core/types"
)

var _ Task = (*EthereumTask)(nil)

type EthereumTask struct {
	Header      types.Header
	Transaction types.Transaction
}

func (e EthereumTask) ID() string {
	return e.Transaction.Hash().String()
}

func (e EthereumTask) Timestamp() time.Time {
	return time.Unix(int64(e.Header.Time), 0)
}
