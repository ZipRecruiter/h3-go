package h3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constrainLng(t *testing.T) {
	type args struct {
		lng float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "lng 0",
			args: args{lng: 0},
			want: 0,
		},
		{
			name: "lng 1",
			args: args{lng: 1},
			want: 1,
		},
		{
			name: "lng pi",
			args: args{lng: math.Pi},
			want: math.Pi,
		},
		{
			name: "lng 2pi",
			args: args{lng: 2 * math.Pi},
			want: 0,
		},
		{
			name: "lng 3pi",
			args: args{lng: 3 * math.Pi},
			want: math.Pi,
		},
		{
			name: "lng 4pi",
			args: args{lng: 4 * math.Pi},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, constrainLng(tt.args.lng), "constrainLng(%v)", tt.args.lng)
		})
	}
}

func Test_constrainLat(t *testing.T) {
	type args struct {
		lat float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "lat 0",
			args: args{lat: 0},
			want: 0,
		},
		{
			name: "lat 1",
			args: args{lat: 1},
			want: 1,
		},
		{
			name: "lat pi/2",
			args: args{lat: M_PI_2},
			want: M_PI_2,
		},
		{
			name: "lat pi",
			args: args{lat: math.Pi},
			want: 0,
		},
		{
			name: "lat pi+1",
			args: args{lat: math.Pi + 1},
			want: 1,
		},
		{
			name: "lat 2pi+1",
			args: args{lat: 2*math.Pi + 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, constrainLat(tt.args.lat), "constrainLat(%v)", tt.args.lat)
		})
	}
}

func TestLatLng_geoAzimuthDistanceRads(t *testing.T) {
	type args struct {
		azimuth  float64
		distance float64
	}
	tests := []struct {
		name string
		l    LatLng
		args args
		want LatLng
	}{
		{
			name: "zero distance produces same point",
			l:    LatLng{15, 10},
			args: args{azimuth: 0, distance: 0},
			want: LatLng{15, 10},
		},
		{
			name: "due north to north pole produces north pole",
			l:    NewLatLng(45, 1),
			args: args{azimuth: 0, distance: deg2rad(45)},
			want: NewLatLng(90, 0),
		},
		{
			name: "due north to south pole produces south pole",
			l:    NewLatLng(45, 1),
			args: args{azimuth: 0, distance: deg2rad(45 + 180)},
			// (doesn't get wrapped properly, but that's known)
			want: NewLatLng(270, 1),
		},
		{
			name: "due south to south pole produces south pole",
			l:    NewLatLng(-45, 2),
			args: args{azimuth: deg2rad(180), distance: deg2rad(45)},
			want: NewLatLng(-90, 0),
		},
		{
			name: "due north produces expected result",
			l:    NewLatLng(-45, 10),
			args: args{azimuth: 0, distance: deg2rad(35)},
			want: NewLatLng(-10, 10),
		},
		{
			name: "some direction to south pole produces south pole",
			l:    NewLatLng(90, 0),
			args: args{azimuth: deg2rad(12), distance: deg2rad(180)},
			want: NewLatLng(-90, 0),
		},
		{
			name: "some direction to north pole produces north pole",
			l:    NewLatLng(-90, 0),
			args: args{azimuth: deg2rad(12), distance: deg2rad(180)},
			want: NewLatLng(90, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.l.geoAzimuthDistanceRads(tt.args.azimuth, tt.args.distance)
			if math.Abs(tt.want.Latitude()-actual.Latitude()) > EPSILON_RAD {
				t.Errorf("latitude: expected %f, got %f", tt.want.Latitude(), actual.Latitude())
			}
			if math.Abs(tt.want.Longitude()-actual.Longitude()) > EPSILON_RAD {
				t.Errorf("longitude: expected %f, got %f", tt.want.Longitude(), actual.Longitude())
			}
		})
	}
}

func TestLatLng_geoAzimuthDistanceRads_invertible(t *testing.T) {
	start := NewLatLng(15, 10)
	azimuth := deg2rad(20)
	degrees180 := deg2rad(180)
	distance := deg2rad(15)

	out := start.geoAzimuthDistanceRads(azimuth, distance)
	assert.InEpsilon(t, distance, start.greatCircleDistanceRads(out), EPSILON_RAD, "moved distance is as expected")

	start2 := out
	out = start2.geoAzimuthDistanceRads(azimuth+degrees180, distance)
	assert.InEpsilon(t, distance, start2.greatCircleDistanceRads(out), EPSILON_RAD, "moved distance is as expected")
}

func Test_deg2rad(t *testing.T) {
	type args struct {
		deg float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0 degrees",
			args: args{deg: 0},
			want: 0,
		},
		{
			name: "90 degrees",
			args: args{deg: 90},
			want: math.Pi / 2,
		},
		{
			name: "180 degrees",
			args: args{deg: 180},
			want: math.Pi,
		},
		{
			name: "270 degrees",
			args: args{deg: 270},
			want: 3 * math.Pi / 2,
		},
		{
			name: "360 degrees",
			args: args{deg: 360},
			want: 2 * math.Pi,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, deg2rad(tt.args.deg), "deg2rad(%v)", tt.args.deg)
		})
	}
}

func Test_deg2rad_invertible(t *testing.T) {
	assert.InDelta(t, 0, rad2deg(deg2rad(0)), EPSILON_RAD, "0 degrees")
	assert.InDelta(t, 90, rad2deg(deg2rad(90)), EPSILON_RAD, "90 degrees")
	assert.InDelta(t, 180, rad2deg(deg2rad(180)), EPSILON_RAD, "180 degrees")
	assert.InDelta(t, 270, rad2deg(deg2rad(270)), EPSILON_RAD, "270 degrees")
	assert.InDelta(t, 360, rad2deg(deg2rad(360)), EPSILON_RAD, "360 degrees")
}

func TestLatLng_greatCircleDistanceRads(t *testing.T) {
	type args struct {
		other LatLng
	}
	tests := []struct {
		name string
		l    LatLng
		args args
		want float64
	}{
		{
			name: "same point is zero",
			l:    NewLatLng(10, 10),
			args: args{other: NewLatLng(10, 10)},
			want: 0,
		},
		{
			name: "distance along longitude",
			l:    NewLatLng(10, 10),
			args: args{other: NewLatLng(0, 10)},
			want: deg2rad(10),
		},
		{
			name: "distance with wrapped longitude",
			l:    LatLng{0, -(math.Pi + M_PI_2)},
			args: args{other: LatLng{0, 0}},
			want: M_PI_2,
		},
		{
			name: "distance with wrapped longitude, swapped args",
			l:    LatLng{0, 0},
			args: args{other: LatLng{0, -(math.Pi + M_PI_2)}},
			want: M_PI_2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.InDelta(t, tt.want, tt.l.greatCircleDistanceRads(tt.args.other), EPSILON_RAD, "greatCircleDistanceRads(%v)", tt.args.other)
		})
	}
}
