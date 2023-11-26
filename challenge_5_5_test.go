package taofp

import (
	"errors"
	"reflect"
	"testing"
)

func TestMapEither(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		e Either[error, int]
	}

	tests := []struct {
		name string
		args args
		want Either[error, int]
	}{
		{
			name: "right",
			args: args{
				f: func(x int) int {
					return x * x
				},
				e: RightOf[error, int](2),
			},
			want: RightOf[error, int](4),
		},
		{
			name: "left",
			args: args{
				f: func(x int) int {
					return x * x
				},
				e: LeftOf[error, int](errors.New("error")),
			},
			want: LeftOf[error, int](errors.New("error")),
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := MapEither[error, int](tt.args.f)(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapEither() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldLeft(t *testing.T) {
	t.Parallel()

	type args struct {
		f    func(int, int) int
		l    []int
		init int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				f: func(x, y int) int {
					return x - y
				},
				l:    []int{1},
				init: 0,
			},
			want: -1,
		},
		{
			name: "minus",
			args: args{
				f: func(x, y int) int {
					return x - y
				},
				l:    []int{1, 2, 3},
				init: 0,
			},
			want: -6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := FoldLeft[[]int](tt.args.f, tt.args.init)(tt.args.l); got != tt.want {
				t.Errorf("FoldLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeToList(t *testing.T) {
	t.Parallel()

	type args struct {
		t *Node[int]
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				t: &Node[int]{
					l: &Node[int]{
						v: 2,
					},
					r: &Node[int]{
						l: &Node[int]{
							v: 4,
						},
						v: 3,
					},
					v: 1,
				},
			},
			want: []int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := TreeToList[[]int](tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TreeToList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAscendingSorted(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "sorted",
			args: args{
				l: []int{1, 2, 3},
			},
			want: true,
		},
		{
			name: "sorted same",
			args: args{
				l: []int{1, 1, 3},
			},
			want: true,
		},
		{
			name: "unsorted",
			args: args{
				l: []int{3, 1, 2},
			},
			want: false,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsAscendingSorted(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IsAscendingSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
