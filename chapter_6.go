package taofp

type (
	Stream[T any] struct {
		Value T
		Next  func() *Stream[T]
	}
)

func NaturalsFrom(n int) *Stream[int] {
	return &Stream[int]{
		Value: n,
		Next: func() *Stream[int] {
			return NaturalsFrom(n + 1)
		},
	}
}

func StreamFromSlice[T any](ts []T) *Stream[T] {
	if len(ts) == 0 {
		return nil
	}

	return &Stream[T]{
		Value: ts[0],
		Next: func() *Stream[T] {
			return StreamFromSlice(ts[1:])
		},
	}
}

func StreamTake[T any](n int, s *Stream[T]) []T {
	if n == 0 {
		return nil
	}

	return append([]T{s.Value}, StreamTake(n-1, s.Next())...)
}

func StreamMap[T, R any](f func(T) R) func(*Stream[T]) *Stream[R] {
	return func(s *Stream[T]) *Stream[R] {
		if s == nil {
			return nil
		}

		return &Stream[R]{
			Value: f(s.Value),
			Next: func() *Stream[R] {
				return StreamMap[T, R](f)(s.Next())
			},
		}
	}
}

func StreamFilter[T any](p func(T) bool) func(*Stream[T]) *Stream[T] {
	return func(s *Stream[T]) *Stream[T] {
		if s == nil {
			return nil
		}

		if p(s.Value) {
			return &Stream[T]{
				Value: s.Value,
				Next: func() *Stream[T] {
					return StreamFilter(p)(s.Next())
				},
			}
		}

		return StreamFilter(p)(s.Next())
	}
}

func StreamZipWith[T, T2, R any](f func(T, T2) R) func(*Stream[T], *Stream[T2]) *Stream[R] {
	return func(t *Stream[T], t2 *Stream[T2]) *Stream[R] {
		if t == nil || t2 == nil {
			return nil
		}

		return &Stream[R]{
			Value: f(t.Value, t2.Value),
			Next: func() *Stream[R] {
				return StreamZipWith(f)(t.Next(), t2.Next())
			},
		}
	}
}
