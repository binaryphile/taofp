package taofp

import "testing"

func TestMaxOf3(t *testing.T) {
	t.Parallel()

	type args struct {
		x int
		y int
		z int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				x: 2,
				y: 3,
				z: 1,
			},
			want: 3,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := MaxOf3(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("MaxOf3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	t.Parallel()

	type args struct {
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
				x: -2,
			},
			want: 2,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
