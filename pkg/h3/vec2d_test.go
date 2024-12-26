package h3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vec2d_Mag(t *testing.T) {
	type fields struct {
		x float64
		y float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "three four five",
			fields: fields{
				x: 3,
				y: 4,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := vec2d{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			assert.Equalf(t, tt.want, d.Mag(), "Mag()")
		})
	}
}

func Test_vec2dIntersect(t *testing.T) {
	type args struct {
		p0 vec2d
		p1 vec2d
		p2 vec2d
		p3 vec2d
	}
	tests := []struct {
		name string
		args args
		want vec2d
	}{
		{
			name: "intersect",
			args: args{
				p0: vec2d{x: 2, y: 2},
				p1: vec2d{x: 6, y: 6},
				p2: vec2d{x: 0, y: 4},
				p3: vec2d{x: 10, y: 4},
			},
			want: vec2d{x: 4, y: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, vec2dIntersect(tt.args.p0, tt.args.p1, tt.args.p2, tt.args.p3), "vec2dIntersect(%v, %v, %v, %v)", tt.args.p0, tt.args.p1, tt.args.p2, tt.args.p3)
		})
	}
}

func Test_vec2d_almostEqual(t *testing.T) {
	type fields struct {
		x float64
		y float64
	}
	type args struct {
		other vec2d
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true for equal vectors",
			fields: fields{
				x: 3,
				y: 4,
			},
			args: args{
				other: vec2d{
					x: 3,
					y: 4,
				},
			},
			want: true,
		},
		{
			name: "false for a different x",
			fields: fields{
				x: 3,
				y: 4,
			},
			args: args{
				other: vec2d{
					x: 3.5,
					y: 4,
				},
			},
			want: false,
		},
		{
			name: "false for a different y",
			fields: fields{
				x: 3,
				y: 4,
			},
			args: args{
				other: vec2d{
					x: 3,
					y: 4.5,
				},
			},
			want: false,
		},
		{
			name: "true for almost equal",
			fields: fields{
				x: 3,
				y: 4,
			},
			args: args{
				other: vec2d{
					x: 3.0000000001,
					y: 4,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := vec2d{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			assert.Equalf(t, tt.want, d.almostEqual(tt.args.other), "almostEqual(%v)", tt.args.other)
		})
	}
}
