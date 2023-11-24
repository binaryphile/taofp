package taofp

func Sum(n int) int {
	nextTerm := func() int {
		return n + Sum(n-1)
	}

	return ifThenElse(n <= 0, 0, nextTerm)
}

func SumIter(s, c, n int) int {
	remainingSum := lazy3(SumIter)(s+c, c+1, n)

	return ifThenElse(c > n, s, remainingSum)
}

func SumGeneral(term func(int) int, m, n int) int {
	currentSum := func() int {
		return term(m) + SumGeneral(term, m+1, n)
	}

	return ifThenElse(m > n, 0, currentSum)
}

func square(x int) int {
	return x * x
}

func SumIntegerSquares(m, n int) int {
	return SumGeneral(square, m, n)
}

var sumIntegerSquares = PartiallyApply3(SumGeneral, square)

func PartiallyApply3[T, T2, T3, R any](f func(T, T2, T3) R, t T) func(T2, T3) R {
	return func(t2 T2, t3 T3) R {
		return f(t, t2, t3)
	}
}

func Accumulate[T any](combiner func(T, T) T, init T, term func(int) T, m, n int) T {
	rest := func() T {
		return combiner(term(m), Accumulate(combiner, init, term, m+1, n))
	}

	return ifThenElse(m > n, init, rest)
}
