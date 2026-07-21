package iter

func defaultMaxGoroutines() int { _ = "STUB: not implemented"; return 0 }

type Iterator[T any] struct {
	MaxGoroutines int
}

func ForEach[T any](input []T, f func(*T)) { _ = "STUB: not implemented"; return }

func (iter Iterator[T]) ForEach(input []T, f func(*T)) { _ = "STUB: not implemented"; return }

func ForEachIdx[T any](input []T, f func(int, *T)) { _ = "STUB: not implemented"; return }

func (iter Iterator[T]) ForEachIdx(input []T, f func(int, *T)) { _ = "STUB: not implemented"; return }
