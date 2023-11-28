package taofp

import (
	"reflect"
	"testing"
)

func TestEnumerateIntegers(t *testing.T) {
	t.Parallel()

	type args struct {
		a, b int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				a: 1,
				b: 10,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := EnumerateIntegers(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnumerateIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHd(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				l: []int{1, 2, 3},
			},
			want: 1,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Hd(tt.args.l); got != tt.want {
				t.Errorf("Hd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTl(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				l: []int{1, 2, 3},
			},
			want: []int{2, 3},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Tl(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNth(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				l: []int{1, 2, 3},
				n: 1,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Nth(tt.args.l, tt.args.n); got != tt.want {
				t.Errorf("Nth() = %v, want %v", got, tt.want)
			}
		})
	}
}

type IntNode = Node[int]

func TestSize(t *testing.T) {
	t.Parallel()

	type args struct {
		node *IntNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				node: &IntNode{
					l: &IntNode{
						v: 2,
					},
					r: &IntNode{
						l: &IntNode{
							v: 4,
						},
						v: 3,
					},
					v: 1,
				},
			},
			want: 4,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Size(tt.args.node); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumTree(t *testing.T) {
	t.Parallel()

	type args struct {
		node *IntNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				node: &IntNode{
					l: &IntNode{
						v: 2,
					},
					r: &IntNode{
						l: &IntNode{
							v: 4,
						},
						v: 3,
					},
					v: 1,
				},
			},
			want: 10,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := SumTree(tt.args.node); got != tt.want {
				t.Errorf("SumTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEitherOf(t *testing.T) {
	t.Parallel()

	type args struct {
		l, r    int
		isRight bool
	}

	tests := []struct {
		name string
		args args
		want Either[int, int]
	}{
		{
			name: "basic",
			args: args{
				r:       2,
				isRight: true,
			},
			want: Either[int, int]{
				r:       2,
				isRight: true,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := EitherOf(tt.args.l, tt.args.r, tt.args.isRight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EitherOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLeftOf(t *testing.T) {
	t.Parallel()

	type args struct {
		l int
	}

	tests := []struct {
		name string
		args args
		want Either[int, int]
	}{
		{
			name: "basic",
			args: args{
				l: 2,
			},
			want: Either[int, int]{
				l: 2,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := LeftOf[int, int](tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeftOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRightOf(t *testing.T) {
	t.Parallel()

	type args struct {
		r int
	}

	tests := []struct {
		name string
		args args
		want Either[int, int]
	}{
		{
			name: "basic",
			args: args{
				r: 2,
			},
			want: Either[int, int]{
				r:       2,
				isRight: true,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := RightOf[int, int](tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeDiv(t *testing.T) {
	t.Parallel()

	type args struct {
		a, b int
	}

	tests := []struct {
		name string
		args args
		want Either[error, int]
	}{
		{
			name: "basic",
			args: args{
				a: 2,
				b: 2,
			},
			want: Either[error, int]{
				r:       1,
				isRight: true,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := SafeDiv(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOptOfOk(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want Opt[int]
	}{
		{
			name: "basic",
			args: args{
				x: 2,
			},
			want: Opt[int]{
				ok: true,
				v:  2,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := OptOfOk(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OptOfOk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListMax(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
	}

	tests := []struct {
		name string
		args args
		want Opt[int]
	}{
		{
			name: "ok",
			args: args{
				l: []int{1, 2, 3},
			},
			want: Opt[int]{
				ok: true,
				v:  3,
			},
		},
		{
			name: "not ok",
			args: args{
				[]int{},
			},
			want: Opt[int]{},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := ListMax(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
