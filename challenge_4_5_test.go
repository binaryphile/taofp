package taofp

import (
	"reflect"
	"testing"
)

func TestLongestString(t *testing.T) {
	t.Parallel()

	type args struct {
		l []string
	}

	tests := []struct {
		name string
		args args
		want Opt[string]
	}{
		{
			name: "ok",
			args: args{
				l: []string{"a", "abc", "ab"},
			},
			want: OptOfOk[string]("abc"),
		},
		{
			name: "not ok",
			args: args{
				l: []string{},
			},
			want: Opt[string]{},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := LongestString(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
