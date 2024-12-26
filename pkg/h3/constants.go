package h3

import "math"

const (
	// MAX_H3_RES is the maximum resolution of an H3 cell.
	MAX_H3_RES = 15
	// NUM_BASE_CELLS is the number of H3 base cells.
	NUM_BASE_CELLS = 122
	// M_PI_2 is pi/2.
	M_PI_2 = math.Pi / 2.0
	// M_2PI is 2*pi.
	M_2PI = 2.0 * math.Pi
	// M_PI_180 is pi/180.
	M_PI_180 = math.Pi / 180.0
	// M_180_PI is 180/pi.
	M_180_PI = 180.0 / math.Pi
	// EPSILON is a floating point difference threshold.
	EPSILON = 0.0000000000000001
	// M_AP7_ROT_RADS is the rotation angle between Class II and Class III resolution axes
	// (asin(sqrt(3.0 / 28.0)))
	M_AP7_ROT_RADS = 0.333473172251832115336090755351601070065900389
	// scaling factor from hex2d resolution 0 unit length (or distance between
	// adjacent cell center points on the plane) to gnomonic unit length.
	RES0_U_GNOMONIC     = 0.38196601125010500003
	INV_RES0_U_GNOMONIC = 2.61803398874989588842
	// M_RSIN60 is 1/sin(60').
	M_RSIN60 = 1.1547005383792515290182975610039149112953
	// M_ONESEVENTH is 1/7.
	M_ONESEVENTH = 1.0 / 7.0
	// M_SQRT3_2 is sqrt(3)/2.
	M_SQRT3_2 = 0.8660254037844386467637231707529361834714
)
