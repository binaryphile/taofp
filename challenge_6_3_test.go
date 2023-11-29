package taofp

import (
	"reflect"
	"testing"
)

func TestMaxCircle(t *testing.T) {
	t.Parallel()

	type args struct {
		shapes []Shape
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "empty",
			args: args{
				shapes: []Shape{},
			},
			want: 0,
		},
		{
			name: "circle",
			args: args{
				shapes: []Shape{
					Circle{radius: 1.0},
					Rectangle{length: 1.0, width: 2.0},
					Circle{radius: 2.0},
					Rectangle{length: 2.0, width: 3.0},
				},
			},
			want: 12.56,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := MaxCircle(tt.args.shapes); got != tt.want {
				t.Errorf("MaxCircle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamMerge(t *testing.T) {
	t.Parallel()

	type args struct {
		first  *Stream[int]
		second *Stream[int]
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "evens then odds",
			args: args{
				first:  Evens(),
				second: Odds(),
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "odds then evens",
			args: args{
				first:  Odds(),
				second: Evens(),
			},
			want: []int{1, 0, 3, 2, 5, 4, 7, 6, 9, 8},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(len(tt.want), StreamMerge(tt.args.first, tt.args.second)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamFibonacci(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		want []int
	}{
		{
			name: "basic",
			want: []int{1, 1, 2, 3, 5, 8, 13, 21},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := StreamTake(len(tt.want), StreamFibonacci()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StreamFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
