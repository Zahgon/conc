package iter

type Mapper[T, R any] Iterator[T]

func Map[T, R any](input []T, f func(*T) R) []R { _ = "STUB: not implemented"; return nil }

func (m Mapper[T, R]) Map(input []T, f func(*T) R) []R { _ = "STUB: not implemented"; return nil }

func MapErr[T, R any](input []T, f func(*T) (R, error)) ([]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m Mapper[T, R]) MapErr(input []T, f func(*T) (R, error)) ([]R, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
