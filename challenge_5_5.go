package taofp

import "cmp"

func MapEither[L, R, R2 any](f func(R) R2) func(Either[L, R]) Either[L, R2] {
	return func(e Either[L, R]) Either[L, R2] {
		if e.isRight {
			return RightOf[L, R2](f(e.r))
		}

		return LeftOf[L, R2](e.l)
	}
}

func FoldLeft[L ~[]T, T any](f func(T, T) T, acc T) func(L) T {
	return func(l L) T {
		if len(l) == 0 {
			return acc
		}

		return FoldLeft[L](f, f(acc, l[0]))(l[1:])
	}
}

func TreeToList[L ~[]T, T any](t *Node[T]) L {
	return TreeFold(
		func(l L, r L, v T) L {
			result := append(L{v}, l...)

			return append(result, r...)
		},
		nil,
	)(t)
}

func TreeFold[T, R any](f func(R, R, T) R, init R) func(*Node[T]) R {
	return func(t *Node[T]) R {
		if t == nil {
			return init
		}

		return f(TreeFold(f, init)(t.l), TreeFold(f, init)(t.r), t.v)
	}
}

func ZipWith[T, T2, R any](f func(T, T2) R) func([]T, []T2) []R {
	return func(ts []T, t2s []T2) []R {
		if len(ts) == 0 || len(t2s) == 0 {
			return []R{}
		}

		return append([]R{f(ts[0], t2s[0])}, ZipWith[T, T2, R](f)(ts[1:], t2s[1:])...)
	}
}

func lte[T cmp.Ordered](t, t2 T) bool {
	return t <= t2
}

func IsAscendingSorted[T cmp.Ordered](l []T) bool {
	if len(l) == 0 {
		return true
	}

	return All(ZipWith[T, T, bool](lte[T])(l, l[1:]))
}
