package taofp

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic",
			args: args{
				x: 3,
			},
			want: true,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := IsPrime(tt.args.x); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	t.Parallel()

	type args struct {
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
				n: 3,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuperFib(t *testing.T) {
	t.Parallel()

	type args struct {
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
				n: 3,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := SuperFib(tt.args.n); got != tt.want {
				t.Errorf("SuperFib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTwice(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		x int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				f: func(x int) int {
					return 2 * x
				},
				x: 1,
			},
			want: 4,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Twice(tt.args.f)(tt.args.x); got != tt.want {
				t.Errorf("Twice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompose(t *testing.T) {
	t.Parallel()

	type args struct {
		f, g func(int) int
		x    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				f: func(x int) int {
					return 2 * x
				},
				g: func(x int) int {
					return 3 * x
				},
				x: 1,
			},
			want: 6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Compose(tt.args.f, tt.args.g)(tt.args.x); got != tt.want {
				t.Errorf("Compose() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilteredAccumulate(t *testing.T) {
	t.Parallel()

	type args struct {
		combiner func(_, _ int) int
		init     int
		term     func(int) int
		p        func(int) bool
		m, n     int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				combiner: func(x, y int) int {
					return x + y
				},
				init: 0,
				term: func(x int) int {
					return x
				},
				p: func(x int) bool {
					return x%2 == 0
				},
				m: 1,
				n: 4,
			},
			want: 6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := FilteredAccumulate[int](tt.args.combiner, tt.args.init, tt.args.term, tt.args.p, tt.args.m, tt.args.n)

			if got != tt.want {
				t.Errorf("FilteredAccumulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
