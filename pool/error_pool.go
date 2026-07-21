package pool

import (
	"context"
	"sync"
)

type ErrorPool struct {
	pool Pool

	onlyFirstError bool

	mu   sync.Mutex
	errs []error
}

func (p *ErrorPool) Go(f func() error) { _ = "STUB: not implemented"; return }

func (p *ErrorPool) Wait() error { _ = "STUB: not implemented"; return nil }

func (p *ErrorPool) WithContext(ctx context.Context) *ContextPool {
	_ = "STUB: not implemented"
	return nil
}

func (p *ErrorPool) WithFirstError() *ErrorPool { _ = "STUB: not implemented"; return nil }

func (p *ErrorPool) WithMaxGoroutines(n int) *ErrorPool { _ = "STUB: not implemented"; return nil }

func (p *ErrorPool) deref() ErrorPool { _ = "STUB: not implemented"; return *new(ErrorPool) }

func (p *ErrorPool) panicIfInitialized() { _ = "STUB: not implemented"; return }

func (p *ErrorPool) addErr(err error) { _ = "STUB: not implemented"; return }
