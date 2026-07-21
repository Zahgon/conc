package conc

import (
	"sync"

	"github.com/sourcegraph/conc/panics"
)

func NewWaitGroup() *WaitGroup { _ = "STUB: not implemented"; return nil }

type WaitGroup struct {
	wg sync.WaitGroup
	pc panics.Catcher
}

func (h *WaitGroup) Go(f func()) { _ = "STUB: not implemented"; return }

func (h *WaitGroup) Wait() { _ = "STUB: not implemented"; return }

func (h *WaitGroup) WaitAndRecover() *panics.Recovered { _ = "STUB: not implemented"; return nil }
