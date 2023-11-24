package taofp

func IsPrime(x int) bool {
	var factorXStartingFrom func(int) bool

	factorXStartingFrom = func(fact int) bool {
		tryNextFactor := lazy(factorXStartingFrom)(fact + 1)

		outOfFactorsOrTryNext := lazy3(ifThenElse[bool])(fact*fact > x, true, tryNextFactor)

		return ifThenElse(x%fact == 0, false, outOfFactorsOrTryNext)
	}

	factorX := lazy(factorXStartingFrom)(2)

	return ifThenElse(x < 2, false, factorX)
}

func Fib(n int) int {
	nextFib := func() int {
		return Fib(n-1) + Fib(n-2)
	}

	baseCaseOrNextFib := lazy3(ifThenElse[int])(n <= 2, 1, nextFib)

	return ifThenElse(n <= 0, 0, baseCaseOrNextFib)
}

func SuperFib(n int) int {
	var helper func(_, _, _ int) int

	helper = func(curr, next, m int) int {
		nextFib := lazy3(helper)(next, curr+next, m+1)

		return ifThenElse(m >= n, curr, nextFib)
	}

	return helper(0, 1, 0)
}

func Twice[T any](f func(T) T) func(T) T {
	return func(t T) T {
		return f(f(t))
	}
}

func Compose[T, T2, T3 any](f func(T2) T3, g func(T) T2) func(T) T3 {
	return func(t T) T3 {
		return f(g(t))
	}
}

func FilteredAccumulate[T any](
	combiner func(T, T) T,
	init T,
	term func(int) T,
	p func(T) bool,
	m, n int,
) T {
	filterOnTerm := func() T {
		tail := lazy6(FilteredAccumulate[T])(combiner, init, term, p, m+1, n)

		mthTerm := term(m)

		combined := func() T {
			return combiner(mthTerm, tail())
		}

		return ifThenElse2(p(mthTerm), combined, tail)
	}

	return ifThenElse(m > n, init, filterOnTerm)
}

func lazy[T, R any](f func(T) R) func(T) func() R {
	return func(t T) func() R {
		return func() R {
			return f(t)
		}
	}
}

func lazy2[T, T2, R any](f func(T, T2) R) func(T, T2) func() R {
	return func(t T, t2 T2) func() R {
		return func() R {
			return f(t, t2)
		}
	}
}

func lazy3[T, T2, T3, R any](f func(T, T2, T3) R) func(T, T2, T3) func() R {
	return func(t T, t2 T2, t3 T3) func() R {
		return func() R {
			return f(t, t2, t3)
		}
	}
}

func lazy6[T, T2, T3, T4, T5, T6, R any](f func(T, T2, T3, T4, T5, T6) R) func(T, T2, T3, T4, T5, T6) func() R {
	return func(t T, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) func() R {
		return func() R {
			return f(t, t2, t3, t4, t5, t6)
		}
	}
}
