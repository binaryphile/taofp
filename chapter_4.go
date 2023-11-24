package taofp

import "errors"

func EnumerateIntegers(a, b int) []int {
	enumerateRest := func() []int {
		return append([]int{a}, EnumerateIntegers(a+1, b)...)
	}

	return ifThenElse(a > b, []int{}, enumerateRest)
}

func Hd[L ~[]T, T any](l L) T {
	return l[0]
}

func Tl[L ~[]T, T any](l L) L {
	if len(l) < 1 {
		return nil
	}

	return append(L{}, l[1:]...)
}

func Nth[L ~[]T, T any](l L, n int) T {
	hd := lazy(Hd[L])(l)

	next := lazy2(Nth[L])(Tl(l), n-1)

	return ifThenElse2(n <= 0, hd, next)
}

func Size[T any](node *Node[T]) int {
	size := func() int {
		return 1 + Size(node.l) + Size(node.r)
	}

	return ifThenElse(node == nil, 0, size)
}

type Node[T any] struct {
	l, r *Node[T]
	v    T
}

func SumTree[T ~int](node *Node[T]) T {
	currentSum := func() T {
		return node.v + SumTree(node.l) + SumTree(node.r)
	}

	return ifThenElse(node == nil, 0, currentSum)
}

type Either[L, R any] struct {
	l     L
	r     R
	right bool
}

func NewEither[L, R any](l L, r R, right bool) Either[L, R] {
	return Either[L, R]{
		l:     l,
		r:     r,
		right: right,
	}
}

var DivisionError = errors.New("division by zero")

func SafeDiv[T ~int](a, b T) Either[error, T] {
	zero := Either[error, T]{}

	if b != 0 {
		return NewEither[error, T](zero.l, a/b, true)
	}

	return NewEither[error, T](DivisionError, zero.r, false)
}
