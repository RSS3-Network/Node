package ethereum

import (
	"context"
	"fmt"
	"github.com/naturalselectionlabs/framework/internal/engine"
	"github.com/naturalselectionlabs/framework/internal/message"
	"math/big"
)

var _ engine.Source[message.EthereumTask] = (*Source)(nil)

type Source struct {
	state *big.Int
}

func (s *Source) Initialize() error {
	s.state = big.NewInt(0) // Load checkpoint

	return nil
}

func (s *Source) Run(ctx context.Context, sourceCtx engine.SourceContext[message.EthereumTask]) error {
	for {
		var task message.EthereumTask

		state := task.Header.Number

		sourceCtx.Collect(task)

		fmt.Println(state) // Save checkpoint
	}
}

func init() {
	// 1, 2, 3 / 1.0.0
	// 1, 2.1, 3, 4 / 2.0.0
	// 2.1, 4
}

// file://
