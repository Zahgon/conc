package stream

import (
	"sync"

	"github.com/sourcegraph/conc"
	"github.com/sourcegraph/conc/pool"
)

func New() *Stream { _ = "STUB: not implemented"; return nil }

type Stream struct {
	pool             pool.Pool
	callbackerHandle conc.WaitGroup
	queue            chan callbackCh

	initOnce sync.Once
}

type Task func() Callback

type Callback func()

func (s *Stream) Go(f Task) { _ = "STUB: not implemented"; return }

func (s *Stream) Wait() { _ = "STUB: not implemented"; return }

func (s *Stream) WithMaxGoroutines(n int) *Stream { _ = "STUB: not implemented"; return nil }

func (s *Stream) init() {
	s.initOnce.Do(func() {
		s.queue = make(chan callbackCh, s.pool.MaxGoroutines()+1)

		s.callbackerHandle.Go(s.callbacker)
	})
}

func (s *Stream) callbacker() { _ = "STUB: not implemented"; return }

type callbackCh chan func()

var callbackChPool = sync.Pool{
	New: func() any {
		return make(callbackCh, 1)
	},
}

func getCh() callbackCh { _ = "STUB: not implemented"; return *new(callbackCh) }

func putCh(ch callbackCh) { _ = "STUB: not implemented"; return }
