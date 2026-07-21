package pool

import (
	"context"
)

type ResultErrorPool[T any] struct {
	errorPool      ErrorPool
	agg            resultAggregator[T]
	collectErrored bool
}

func (p *ResultErrorPool[T]) Go(f func() (T, error)) { _ = "STUB: not implemented"; return }

func (p *ResultErrorPool[T]) Wait() ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *ResultErrorPool[T]) WithCollectErrored() *ResultErrorPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultErrorPool[T]) WithContext(ctx context.Context) *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultErrorPool[T]) WithFirstError() *ResultErrorPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultErrorPool[T]) WithMaxGoroutines(n int) *ResultErrorPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultErrorPool[T]) panicIfInitialized() { _ = "STUB: not implemented"; return }
