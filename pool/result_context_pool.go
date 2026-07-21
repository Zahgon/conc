package pool

import (
	"context"
)

type ResultContextPool[T any] struct {
	contextPool    ContextPool
	agg            resultAggregator[T]
	collectErrored bool
}

func (p *ResultContextPool[T]) Go(f func(context.Context) (T, error)) {
	_ = "STUB: not implemented"
	return
}

func (p *ResultContextPool[T]) Wait() ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *ResultContextPool[T]) WithCollectErrored() *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultContextPool[T]) WithFirstError() *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultContextPool[T]) WithCancelOnError() *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultContextPool[T]) WithFailFast() *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultContextPool[T]) WithMaxGoroutines(n int) *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultContextPool[T]) panicIfInitialized() { _ = "STUB: not implemented"; return }
