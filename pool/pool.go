package pool

import (
	"context"
	"sync"

	"github.com/sourcegraph/conc"
)

func New() *Pool { _ = "STUB: not implemented"; return nil }

type Pool struct {
	handle   conc.WaitGroup
	limiter  limiter
	tasks    chan func()
	initOnce sync.Once
}

func (p *Pool) Go(f func()) { _ = "STUB: not implemented"; return }

func (p *Pool) Wait() { _ = "STUB: not implemented"; return }

func (p *Pool) MaxGoroutines() int { _ = "STUB: not implemented"; return 0 }

func (p *Pool) WithMaxGoroutines(n int) *Pool { _ = "STUB: not implemented"; return nil }

func (p *Pool) init() {
	p.initOnce.Do(func() {
		p.tasks = make(chan func())
	})
}

func (p *Pool) panicIfInitialized() { _ = "STUB: not implemented"; return }

func (p *Pool) WithErrors() *ErrorPool { _ = "STUB: not implemented"; return nil }

func (p *Pool) deref() Pool { _ = "STUB: not implemented"; return *new(Pool) }

func (p *Pool) WithContext(ctx context.Context) *ContextPool { _ = "STUB: not implemented"; return nil }

func (p *Pool) worker(initialFunc func()) { _ = "STUB: not implemented"; return }

type limiter chan struct{}

func (l limiter) limit() int { _ = "STUB: not implemented"; return 0 }

func (l limiter) release() { _ = "STUB: not implemented"; return }
