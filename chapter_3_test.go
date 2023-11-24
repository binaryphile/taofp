package taofp

import "testing"

func TestSum(t *testing.T) {
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
			want: 6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Sum(tt.args.n)

			if got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumIter(t *testing.T) {
	type args struct {
		s, c, n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				s: 0,
				c: 1,
				n: 3,
			},
			want: 6,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := SumIter(tt.args.s, tt.args.c, tt.args.n)

			if got != tt.want {
				t.Errorf("SumIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumGeneral(t *testing.T) {
	type args struct {
		term func(int) int
		m, n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				term: func(x int) int {
					return x * 2
				},
				m: 1,
				n: 3,
			},
			want: 12,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := SumGeneral(tt.args.term, tt.args.m, tt.args.n)

			if got != tt.want {
				t.Errorf("SumGeneral() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccumulate(t *testing.T) {
	type args struct {
		combiner func(int, int) int
		init     int
		term     func(int) int
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
					return x * y
				},
				init: 1,
				term: func(x int) int {
					return x
				},
				m: 1,
				n: 4,
			},
			want: 24,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Accumulate(tt.args.combiner, tt.args.init, tt.args.term, tt.args.m, tt.args.n)

			if got != tt.want {
				t.Errorf("Accumulate() = %v, want %v", got, tt.want)
			}
		})
	}
}
