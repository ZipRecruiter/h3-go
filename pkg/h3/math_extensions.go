package h3

import "math"

// addInt32sWouldOverflow evaluates to true if a + b would overflow for int32.
func addInt32sWouldOverflow(a int, b int) bool {
	if a > 0 {
		return math.MaxInt32-a < b
	} else {
		return math.MinInt32-a > b
	}
}

// subInt32sWouldOverflow evaluates to true if a - b would overflow for int32.
func subInt32sWouldOverflow(a int, b int) bool {
	if a >= 0 {
		return math.MinInt32+a >= b
	} else {
		return math.MaxInt32+a+1 < b
	}
}

// ipow does integer exponentiation efficiently. Taken from StackOverflow.
//
// @param base the integer base (can be positive or negative)
// @param exp the integer exponent (should be non-negative)
//
// @return the exponentiated value
func ipow(base int, exp int) int {
	result := 1
	for exp != 0 {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		base *= base
	}
	return result
}
