package syncx

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
)

var ErrorNoResult = errors.New("no result")

type QuickGroup[R any] struct {
	waitGroup  sync.WaitGroup
	ctx        context.Context
	cancel     []context.CancelFunc
	resultOnce sync.Once
	result     R
	hasResult  bool
	mu         sync.RWMutex
	done       atomic.Bool
}

func (q *QuickGroup[R]) Go(f func(ctx context.Context) (R, error)) {
	if q.done.Load() {
		return
	}

	q.waitGroup.Add(1)

	q.mu.Lock()
	ctx, cancel := context.WithCancel(q.ctx)
	q.cancel = append(q.cancel, cancel)
	id := len(q.cancel) - 1
	q.mu.Unlock()

	go func() {
		defer q.waitGroup.Done()

		if result, err := f(ctx); err == nil {
			q.resultOnce.Do(func() {
				q.result = result
				q.hasResult = true
				q.done.Store(true)

				q.mu.Lock()
				defer q.mu.Unlock()

				for index := range q.cancel {
					if index == id {
						continue
					}

					if q.cancel[index] != nil {
						q.cancel[index]()
					}
				}
			})
		}
	}()
}

func (q *QuickGroup[R]) Wait() (result R, err error) {
	q.waitGroup.Wait()

	if q.hasResult {
		return q.result, nil
	}

	return result, ErrorNoResult
}

func NewQuickGroup[R any](ctx context.Context) *QuickGroup[R] {
	quickGroup := &QuickGroup[R]{
		ctx: ctx,
	}

	return quickGroup
}
