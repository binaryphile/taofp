package taofp

import "github.com/mariomac/gostream/stream"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func MaxCircle(shapes []Shape) float64 {
	isCircle := func(shape Shape) bool {
		_, ok := shape.(Circle)

		return ok
	}

	maxFloat := func(x, y float64) float64 {
		return max(x, y)
	}

	circles := stream.OfSlice(shapes).Filter(isCircle)

	maxArea, ok := stream.Map(circles, Shape.Area).Reduce(maxFloat)
	if !ok {
		maxArea = 0
	}

	return maxArea
}

func Evens() *Stream[int] {
	return StreamFilter(func(x int) bool { return x%2 == 0 })(NaturalsFrom(0))
}

func Odds() *Stream[int] {
	return StreamFilter(func(x int) bool { return x%2 != 0 })(NaturalsFrom(0))
}

func StreamMerge[T any](first, second *Stream[T]) *Stream[T] {
	if first == nil && second == nil {
		return nil
	}

	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	return &Stream[T]{
		Value: first.Value,
		Next: func() *Stream[T] {
			return &Stream[T]{
				Value: second.Value,
				Next: func() *Stream[T] {
					return StreamMerge(first.Next(), second.Next())
				},
			}
		},
	}
}

func add(x, y int) int {
	return x + y
}

func StreamFibonacci() *Stream[int] {
	return &Stream[int]{
		Value: 1,
		Next: func() *Stream[int] {
			return &Stream[int]{
				Value: 1,
				Next: func() *Stream[int] {
					return StreamZipWith(add)(StreamFibonacci(), StreamFibonacci().Next())
				},
			}
		}}
}
