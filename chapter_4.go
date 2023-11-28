package taofp

import (
	"cmp"
	"errors"
)

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
	l       L
	r       R
	isRight bool
}

func EitherOf[L, R any](l L, r R, isRight bool) Either[L, R] {
	return Either[L, R]{
		l:       l,
		r:       r,
		isRight: isRight,
	}
}

func LeftOf[L, R any](l L) Either[L, R] {
	return Either[L, R]{
		l: l,
	}
}

func RightOf[L, R any](r R) Either[L, R] {
	return Either[L, R]{
		r:       r,
		isRight: true,
	}
}

var DivisionError = errors.New("division by zero")

func SafeDiv[T ~int](a, b T) Either[error, T] {
	zero := Either[error, T]{}

	if b != 0 {
		return EitherOf[error, T](zero.l, a/b, true)
	}

	return EitherOf[error, T](DivisionError, zero.r, false)
}

type Opt[T any] struct {
	ok bool
	v  T
}

func OptOfOk[T any](v T) Opt[T] {
	return Opt[T]{
		ok: true,
		v:  v,
	}
}

func ListMax[L ~[]T, T cmp.Ordered](l L) Opt[T] {
	if len(l) == 0 {
		return Opt[T]{}
	}

	tailMax := ListMax[L](l[1:])

	if !tailMax.ok {
		return OptOfOk(l[0])
	}

	return OptOfOk(max(l[0], tailMax.v))
}
