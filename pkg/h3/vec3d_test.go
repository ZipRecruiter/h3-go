package h3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vec3d_pointSquareDistance(t *testing.T) {
	type args struct {
		v2 vec3d
	}
	tests := []struct {
		name string
		v    vec3d
		args args
		want float64
	}{
		{
			name: "distance to self is zero",
			v:    vec3d{0, 0, 0},
			args: args{v2: vec3d{0, 0, 0}},
			want: 0,
		},
		{
			name: "distance to (1, 0, 0) is 1",
			v:    vec3d{0, 0, 0},
			args: args{v2: vec3d{1, 0, 0}},
			want: 1,
		},
		{
			name: "distance to (0, 1, 1) is 2",
			v:    vec3d{0, 0, 0},
			args: args{v2: vec3d{0, 1, 1}},
			want: 2,
		},
		{
			name: "distance to (1, 1, 1) is 3",
			v:    vec3d{0, 0, 0},
			args: args{v2: vec3d{1, 1, 1}},
			want: 3,
		},
		{
			name: "distance to (1, 1, 2) is 6",
			v:    vec3d{0, 0, 0},
			args: args{v2: vec3d{1, 1, 2}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.v.pointSquareDistance(tt.args.v2), "pointSquareDistance(%v)", tt.args.v2)
		})
	}
}

func Test_newVec3dFromLatLng(t *testing.T) {
	origin := vec3d{0, 0, 0}

	c1 := LatLng{0, 0}
	p1 := newVec3dFromLatLng(c1)
	assert.InEpsilon(t, 1.0, origin.pointSquareDistance(p1), EPSILON_RAD, "geo point is on the unit sphere")
	c2 := LatLng{M_PI_2, 0}
	p2 := newVec3dFromLatLng(c2)
	assert.InEpsilon(t, 2.0, p1.pointSquareDistance(p2), EPSILON_RAD, "geo point is on another axis")
	c3 := LatLng{math.Pi, 0}
	p3 := newVec3dFromLatLng(c3)
	assert.InEpsilon(t, 4.0, p1.pointSquareDistance(p3), EPSILON_RAD, "geo point is on the opposite side of the sphere")
}
