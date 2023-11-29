package taofp

import (
	"reflect"
	"testing"
)

func TestNaturalsFrom(t *testing.T) {
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
				n: 1,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := NaturalsFrom(tt.args.n).Next().Value; got != tt.want {
				t.Errorf("NaturalsFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamTake(t *testing.T) {
	t.Parallel()

	type args struct {
		n int
		s *Stream[int]
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				n: 2,
				s: StreamFromSlice([]int{1, 2, 3}),
			},
			want: []int{1, 2},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(tt.args.n, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamTake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamMap(t *testing.T) {
	t.Parallel()

	type args struct {
		f func(int) int
		s *Stream[int]
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				f: func(x int) int {
					return x * x
				},
				s: StreamFromSlice([]int{1, 2, 3}),
			},
			want: []int{1, 4, 9},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(3, StreamMap(tt.args.f)(tt.args.s)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamFilter(t *testing.T) {
	t.Parallel()

	type args struct {
		p func(int) bool
		s *Stream[int]
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic",
			args: args{
				p: func(x int) bool {
					return x == 2
				},
				s: StreamFromSlice([]int{1, 2, 3}),
			},
			want: []int{2},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(len(tt.want), StreamFilter(tt.args.p)(tt.args.s)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamZipWith(t *testing.T) {
	t.Parallel()

	type args struct {
		f    func(int, int) bool
		s, t *Stream[int]
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "basic",
			args: args{
				f: func(x, y int) bool {
					return x == y
				},
				s: StreamFromSlice([]int{1, 2, 3}),
				t: StreamFromSlice([]int{0, 2, 2}),
			},
			want: []bool{false, true, false},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(len(tt.want), StreamZipWith(tt.args.f)(tt.args.s, tt.args.t)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamZipWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
