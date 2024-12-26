package h3

import "math"

type vec2d struct {
	x, y float64
}

func (d vec2d) X() float64 {
	return d.x
}

func (d vec2d) Y() float64 {
	return d.y
}

// Mag returns the magnitude of the vector.
func (d vec2d) Mag() float64 {
	return math.Sqrt(d.x*d.x + d.y*d.y)
}

// almostEqualThreshold returns true if the two vectors are almost equal, within some threshold.
func (d vec2d) almostEqualThreshold(other vec2d, threshold float64) bool {
	return math.Abs(d.x-other.x) < threshold && math.Abs(d.y-other.y) < threshold
}

// almostEqual returns true if the two vectors are almost equal.
func (d vec2d) almostEqual(other vec2d) bool {
	return d.almostEqualThreshold(other, 0.0000001)
}

// vec2dIntersect finds the intersection between two lines, assuming the lines intersect and that the intersection is not at an endpoint of either line.
func vec2dIntersect(p0, p1, p2, p3 vec2d) vec2d {
	var s1, s2 vec2d
	s1.x = p1.x - p0.x
	s1.y = p1.y - p0.y
	s2.x = p3.x - p2.x
	s2.y = p3.y - p2.y

	t := (s2.x*(p0.y-p2.y) - s2.y*(p0.x-p2.x)) /
		(-s2.x*s1.y + s1.x*s2.y)

	return vec2d{
		x: p0.x + (t * s1.x),
		y: p0.y + (t * s1.y),
	}
}
