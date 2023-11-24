package taofp

type (
	lazyBool = func() bool
)

func MaxOf3(x, y, z int) int {
	var m int

	if x > y {
		m = x
	} else {
		m = y
	}

	if z > m {
		m = z
	}

	return m
}

func Abs(x int) int {
	if x > 0 {
		return x
	}

	return -x
}

func ifThenElse[T any](pred bool, then T, elseDo func() T) T {
	if pred {
		return then
	}

	return elseDo()
}

func ifThenElse2[T any](pred bool, thenDo func() T, elseDo func() T) T {
	if pred {
		return thenDo()
	}

	return elseDo()
}

//func ifThenElse3[T any](predDo lazyBool, thenDo func() T, elseDo func() T) T {
//	if predDo() {
//		return thenDo()
//	}
//
//	return elseDo()
//}
