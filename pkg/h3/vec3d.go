package h3

import "math"

type vec3d struct {
	x, y, z float64
}

// pointSquareDistance returns the square of the distance between two 3D coordinates.
func (v vec3d) pointSquareDistance(v2 vec3d) float64 {
	return square(v.x-v2.x) + square(v.y-v2.y) + square(v.z-v2.z)
}

// square returns the square of the input f.
func square(f float64) float64 {
	return f * f
}

// newVec3dFromLatLng calculates the 3D coordinate on unit sphere from the latitude and longitude.
func newVec3dFromLatLng(l LatLng) vec3d {
	r := math.Cos(l.Latitude())

	return vec3d{
		z: math.Sin(l.Latitude()),
		x: math.Cos(l.Longitude()) * r,
		y: math.Sin(l.Longitude()) * r,
	}
}
