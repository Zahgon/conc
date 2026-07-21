package pool

import (
	"context"
	"sync"
)

func NewWithResults[T any]() *ResultPool[T] { _ = "STUB: not implemented"; return nil }

type ResultPool[T any] struct {
	pool Pool
	agg  resultAggregator[T]
}

func (p *ResultPool[T]) Go(f func() T) { _ = "STUB: not implemented"; return }

func (p *ResultPool[T]) Wait() []T { _ = "STUB: not implemented"; return nil }

func (p *ResultPool[T]) MaxGoroutines() int { _ = "STUB: not implemented"; return 0 }

func (p *ResultPool[T]) WithErrors() *ResultErrorPool[T] { _ = "STUB: not implemented"; return nil }

func (p *ResultPool[T]) WithContext(ctx context.Context) *ResultContextPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultPool[T]) WithMaxGoroutines(n int) *ResultPool[T] {
	_ = "STUB: not implemented"
	return nil
}

func (p *ResultPool[T]) panicIfInitialized() { _ = "STUB: not implemented"; return }

type resultAggregator[T any] struct {
	mu      sync.Mutex
	len     int
	results []T
	errored []int
}

func (r *resultAggregator[T]) nextIndex() int { _ = "STUB: not implemented"; return 0 }

func (r *resultAggregator[T]) save(i int, res T, errored bool) { _ = "STUB: not implemented"; return }

func (r *resultAggregator[T]) collect(collectErrored bool) []T {
	_ = "STUB: not implemented"
	return nil
}
