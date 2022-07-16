package tool

import (
	"testing"
)

func TestMult(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test normal",
			args: args{
				a: 4,
				b: 5,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mult(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlus(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test normal",
			args: args{
				a: 5,
				b: 4,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plus(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Plus() = %v, want %v", got, tt.want)
			}
		})
	}
}
