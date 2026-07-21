package pool

import (
	"context"
)

type ContextPool struct {
	errorPool ErrorPool

	ctx    context.Context
	cancel context.CancelFunc

	cancelOnError bool
}

func (p *ContextPool) Go(f func(ctx context.Context) error) { _ = "STUB: not implemented"; return }

func (p *ContextPool) Wait() error { _ = "STUB: not implemented"; return nil }

func (p *ContextPool) WithFirstError() *ContextPool { _ = "STUB: not implemented"; return nil }

func (p *ContextPool) WithCancelOnError() *ContextPool { _ = "STUB: not implemented"; return nil }

func (p *ContextPool) WithFailFast() *ContextPool { _ = "STUB: not implemented"; return nil }

func (p *ContextPool) WithMaxGoroutines(n int) *ContextPool { _ = "STUB: not implemented"; return nil }

func (p *ContextPool) panicIfInitialized() { _ = "STUB: not implemented"; return }
