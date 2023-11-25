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
				t.Errorf("LongestString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	t.Parallel()

	type args struct {
		l         []string
		delimiter string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				l:         []string{},
				delimiter: ",",
			},
			want: "",
		},
		{
			name: "single",
			args: args{
				l:         []string{"a"},
				delimiter: ",",
			},
			want: "a",
		},
		{
			name: "double",
			args: args{
				l:         []string{"a", "b"},
				delimiter: ",",
			},
			want: "a,b",
		},
		{
			name: "triple",
			args: args{
				l:         []string{"a", "b", "c"},
				delimiter: "--",
			},
			want: "a--b--c",
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Concat(tt.args.l, tt.args.delimiter); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeight(t *testing.T) {
	t.Parallel()

	type args struct {
		n *Node[int]
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic",
			args: args{
				n: &Node[int]{
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
			want: 3,
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Height(tt.args.n); got != tt.want {
				t.Errorf("Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPred(t *testing.T) {
	t.Parallel()

	type args struct {
		n *Nat
	}

	tests := []struct {
		name string
		args args
		want Opt[*Nat]
	}{
		{
			name: "of zero",
			args: args{
				n: nil,
			},
			want: Opt[*Nat]{},
		},
		{
			name: "of one",
			args: args{
				n: &Nat{},
			},
			want: OptOfOk[*Nat](nil),
		},
		{
			name: "of two",
			args: args{
				n: &Nat{
					Nat: &Nat{},
				},
			},
			want: OptOfOk(&Nat{}),
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Pred(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pred() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	type args struct {
		m, n *Nat
	}

	tests := []struct {
		name string
		args args
		want *Nat
	}{
		{
			name: "zero and zero",
			args: args{
				m: nil,
				n: nil,
			},
			want: nil,
		},
		{
			name: "zero and one",
			args: args{
				m: nil,
				n: &Nat{},
			},
			want: &Nat{},
		},
		{
			name: "one and one",
			args: args{
				m: &Nat{},
				n: &Nat{},
			},
			want: &Nat{
				Nat: &Nat{},
			},
		},
	}

	for _, test := range tests {
		tt := test

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := Add(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
