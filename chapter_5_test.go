package taofp

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		l []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "square",
			args: args{
				f: func(x int) int {
					return x * x
				},
				l: []int{1, 2, 3},
			},
			want: []int{1, 4, 9},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Map(tt.args.f)(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapTree(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		t *Node[int]
	}

	tests := []struct {
		name string
		args args
		want *Node[int]
	}{
		{
			name: "basic",
			args: args{
				f: func(x int) int {
					return x * x
				},
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
			want: &Node[int]{
				l: &Node[int]{
					v: 4,
				},
				r: &Node[int]{
					l: &Node[int]{
						v: 16,
					},
					v: 9,
				},
				v: 1,
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := MapTree(tt.args.f)(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapOpt(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		o Opt[int]
	}

	tests := []struct {
		name string
		args args
		want Opt[int]
	}{
		{
			name: "ok",
			args: args{
				f: func(x int) int {
					return x * x
				},
				o: OptOfOk(2),
			},
			want: OptOfOk(4),
		},
		{
			name: "not ok",
			args: args{
				f: func(x int) int {
					return x * x
				},
				o: Opt[int]{},
			},
			want: Opt[int]{},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := MapOpt(tt.args.f)(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	type args struct {
		l []int
		p func(int) bool
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "evens",
			args: args{
				l: []int{1, 2, 3, 4, 5},
				p: func(x int) bool {
					return x%2 == 0
				},
			},
			want: []int{2, 4},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Filter(tt.args.p)(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFoldRight(t *testing.T) {
	t.Parallel()

	type args struct {
		f    func(int, int) int
		init int
		l    []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "product",
			args: args{
				f: func(x, y int) int {
					return x * y
				},
				init: 1,
				l:    []int{1, 2, 3},
			},
			want: 6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := FoldRight[[]int](tt.args.f, tt.args.init)(tt.args.l); got != tt.want {
				t.Errorf("FoldRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny(t *testing.T) {
	t.Parallel()

	type args struct {
		l []bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				l: []bool{},
			},
			want: false,
		},
		{
			name: "false",
			args: args{
				l: []bool{false, false},
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				l: []bool{false, true},
			},
			want: true,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Any(tt.args.l); got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	type args struct {
		l []bool
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty",
			args: args{
				l: []bool{},
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				l: []bool{false, true},
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				l: []bool{true, true},
			},
			want: true,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := All(tt.args.l); got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
