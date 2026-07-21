package panics

import (
	"sync/atomic"
)

type Catcher struct {
	recovered atomic.Pointer[Recovered]
}

func (p *Catcher) Try(f func()) { _ = "STUB: not implemented"; return }

func (p *Catcher) tryRecover() { _ = "STUB: not implemented"; return }

func (p *Catcher) Repanic() { _ = "STUB: not implemented"; return }

func (p *Catcher) Recovered() *Recovered { _ = "STUB: not implemented"; return nil }

func NewRecovered(skip int, value any) Recovered { _ = "STUB: not implemented"; return *new(Recovered) }

type Recovered struct {
	Value any

	Callers []uintptr

	Stack []byte
}

func (p *Recovered) String() string { _ = "STUB: not implemented"; return "" }

func (p *Recovered) AsError() error { _ = "STUB: not implemented"; return nil }

type ErrRecovered struct{ Recovered }

var _ error = (*ErrRecovered)(nil)

func (p *ErrRecovered) Error() string { _ = "STUB: not implemented"; return "" }

func (p *ErrRecovered) Unwrap() error { _ = "STUB: not implemented"; return nil }
