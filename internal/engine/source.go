package engine

import "context"

type SourceContext[T any] interface {
	Collect(elements ...T)
}

type Source[T any] interface {
	Run(ctx context.Context, sourceCtx SourceContext[T]) error

	// SaveCheckpoint()
	// LoadCheckpoint()
}
