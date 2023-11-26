package taofp

func Map[T, R any](f func(T) R) func([]T) []R {
	return func(l []T) []R {
		if len(l) == 0 {
			return []R{}
		}

		return append([]R{f(l[0])}, Map(f)(l[1:])...)
	}
}

func MapTree[T, R any](f func(T) R) func(*Node[T]) *Node[R] {
	return func(t *Node[T]) *Node[R] {
		if t == nil {
			return nil
		}

		return &Node[R]{
			l: MapTree(f)(t.l),
			r: MapTree(f)(t.r),
			v: f(t.v),
		}
	}
}

func MapOpt[T, R any](f func(T) R) func(Opt[T]) Opt[R] {
	return func(t Opt[T]) Opt[R] {
		if t.ok {
			return OptOfOk(f(t.v))
		}

		return Opt[R]{}
	}
}

func Filter[T any](p func(T) bool) func([]T) []T {
	return func(l []T) []T {
		if len(l) == 0 {
			return []T{}
		}

		tail := Filter(p)(l[1:])

		if p(l[0]) {
			return append([]T{l[0]}, tail...)
		}

		return tail
	}
}

func FoldRight[L ~[]T, T any](f func(T, T) T, init T) func(L) T {
	return func(l L) T {
		if len(l) == 0 {
			return init
		}

		return f(l[0], FoldRight[L](f, init)(l[1:]))
	}
}

var Any = FoldRight[[]bool](func(agg, curr bool) bool {
	return agg || curr
}, false)

var All = FoldRight[[]bool](func(agg, curr bool) bool {
	return agg && curr
}, true)
