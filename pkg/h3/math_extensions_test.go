package h3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addInt32sWouldOverflow(t *testing.T) {
	assertAddDoesNotOverflow := func(a, b int) {
		if addInt32sWouldOverflow(a, b) {
			t.Errorf("%d + %d should not overflow", a, b)
		}
	}

	assertAddOverflow := func(a, b int) {
		if !addInt32sWouldOverflow(a, b) {
			t.Errorf("%d + %d should overflow", a, b)
		}
	}

	assertAddDoesNotOverflow(0, 0)
	assertAddDoesNotOverflow(math.MinInt32, 0)
	assertAddDoesNotOverflow(math.MinInt32, 1)
	assertAddOverflow(math.MinInt32, -1)
	assertAddDoesNotOverflow(math.MinInt32+1, 0)
	assertAddDoesNotOverflow(math.MinInt32+1, 1)
	assertAddDoesNotOverflow(math.MinInt32+1, -1)
	assertAddDoesNotOverflow(math.MinInt32+1, 2)
	assertAddOverflow(math.MinInt32+1, -2)
	assertAddDoesNotOverflow(100, 10)
	assertAddDoesNotOverflow(math.MaxInt32, 0)
	assertAddOverflow(math.MaxInt32, 1)
	assertAddDoesNotOverflow(math.MaxInt32, -1)
	assertAddDoesNotOverflow(math.MaxInt32-1, 1)
	assertAddDoesNotOverflow(math.MaxInt32-1, -1)
	assertAddDoesNotOverflow(math.MaxInt32-1, -2)
	assertAddOverflow(math.MaxInt32-1, 2)
	assertAddDoesNotOverflow(math.MinInt32, math.MaxInt32)
	assertAddDoesNotOverflow(math.MaxInt32, math.MinInt32)
	assertAddOverflow(math.MaxInt32, math.MaxInt32)
	assertAddOverflow(math.MinInt32, math.MinInt32)
	assertAddDoesNotOverflow(-1, 0)
	assertAddDoesNotOverflow(-1, 10)
	assertAddDoesNotOverflow(-1, -10)
	assertAddDoesNotOverflow(-1, math.MaxInt32)
	assertAddDoesNotOverflow(-2, math.MaxInt32)
	assertAddOverflow(-1, math.MinInt32)
	assertAddDoesNotOverflow(0, math.MinInt32)
}

func Test_ipow(t *testing.T) {
	type args struct {
		base int
		exp  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "7 ** 0 == 1",
			args: args{7, 0},
			want: 1,
		},
		{
			name: "7 ** 1 == 7",
			args: args{7, 1},
			want: 7,
		},
		{
			name: "7 ** 2 == 49",
			args: args{7, 2},
			want: 49,
		},
		{
			name: "1 ** 20 == 1",
			args: args{1, 20},
			want: 1,
		},
		{
			name: "2 ** 5 == 32",
			args: args{2, 5},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ipow(tt.args.base, tt.args.exp), "ipow(%v, %v)", tt.args.base, tt.args.exp)
		})
	}
}

func Test_subInt32sWouldOverflow(t *testing.T) {
	assertSubDoesNotOverflow := func(a, b int) {
		if subInt32sWouldOverflow(a, b) {
			t.Errorf("%d - %d should not overflow", a, b)
		}
	}

	assertSubOverflow := func(a, b int) {
		if !subInt32sWouldOverflow(a, b) {
			t.Errorf("%d - %d should overflow", a, b)
		}
	}

	assertSubDoesNotOverflow(0, 0)
	assertSubDoesNotOverflow(math.MinInt32, 0)
	assertSubOverflow(math.MinInt32, 1)
	assertSubDoesNotOverflow(math.MinInt32, -1)
	assertSubDoesNotOverflow(math.MinInt32+1, 0)
	assertSubDoesNotOverflow(math.MinInt32+1, 1)
	assertSubDoesNotOverflow(math.MinInt32+1, -1)
	assertSubOverflow(math.MinInt32+1, 2)
	assertSubDoesNotOverflow(math.MinInt32+1, -2)
	assertSubDoesNotOverflow(100, 10)
	assertSubDoesNotOverflow(math.MaxInt32, 0)
	assertSubDoesNotOverflow(math.MaxInt32, 1)
	assertSubOverflow(math.MaxInt32, -1)
	assertSubDoesNotOverflow(math.MaxInt32-1, 1)
	assertSubDoesNotOverflow(math.MaxInt32-1, -1)
	assertSubOverflow(math.MaxInt32-1, -2)
	assertSubOverflow(math.MaxInt32-1, -2)
	assertSubOverflow(math.MinInt32, math.MaxInt32)
	assertSubOverflow(math.MaxInt32, math.MinInt32)
	assertSubDoesNotOverflow(math.MinInt32, math.MinInt32)
	assertSubDoesNotOverflow(math.MaxInt32, math.MaxInt32)
	assertSubDoesNotOverflow(-1, 0)
	assertSubDoesNotOverflow(-1, 10)
	assertSubDoesNotOverflow(-1, -10)
	assertSubDoesNotOverflow(-1, math.MaxInt32)
	assertSubOverflow(-2, math.MaxInt32)
	assertSubDoesNotOverflow(-1, math.MinInt32)
	assertSubOverflow(0, math.MinInt32)
}
