package h3

// Overage indicates whether a FaceIJK address is on the original face or on an
// adjacent face.
type Overage int

const (
	// NO_OVERAGE indicates coordinates are on the original face.
	NO_OVERAGE Overage = iota
	// FACE_EDGE indicates coordinates are on an edge of the original face.
	FACE_EDGE
	// NEW_FACE indicates coordinates are on an adjacent face.
	NEW_FACE
)

// Direction indices for faceNeighbors lookup. Entry 0 is the central face
// (self), 1 = IJ quadrant, 2 = KI quadrant, 3 = JK quadrant.
const (
	ovIJ = 1
	ovKI = 2
	ovJK = 3
)

// faceOrientIJK describes the orientation of a neighboring face relative to
// the current face in IJK coordinates.
type faceOrientIJK struct {
	face      int
	translate coordIJK
	ccwRot60  int
}

// faceNeighbors defines, for each icosahedral face, the neighboring faces
// along the IJ, KI, and JK axes, including the translation vector and
// rotation count for coordinate transformation. Index 0 is the central
// face (self). Ported from the H3 C reference (faceijk.c).
var faceNeighbors = [NUM_ICOSA_FACES][4]faceOrientIJK{
	{ // face 0
		{0, coordIJK{0, 0, 0}, 0},  // central
		{4, coordIJK{2, 0, 2}, 1},  // ij
		{1, coordIJK{2, 2, 0}, 5},  // ki
		{5, coordIJK{0, 2, 2}, 3},  // jk
	},
	{ // face 1
		{1, coordIJK{0, 0, 0}, 0},
		{0, coordIJK{2, 0, 2}, 1},
		{2, coordIJK{2, 2, 0}, 5},
		{6, coordIJK{0, 2, 2}, 3},
	},
	{ // face 2
		{2, coordIJK{0, 0, 0}, 0},
		{1, coordIJK{2, 0, 2}, 1},
		{3, coordIJK{2, 2, 0}, 5},
		{7, coordIJK{0, 2, 2}, 3},
	},
	{ // face 3
		{3, coordIJK{0, 0, 0}, 0},
		{2, coordIJK{2, 0, 2}, 1},
		{4, coordIJK{2, 2, 0}, 5},
		{8, coordIJK{0, 2, 2}, 3},
	},
	{ // face 4
		{4, coordIJK{0, 0, 0}, 0},
		{3, coordIJK{2, 0, 2}, 1},
		{0, coordIJK{2, 2, 0}, 5},
		{9, coordIJK{0, 2, 2}, 3},
	},
	{ // face 5
		{5, coordIJK{0, 0, 0}, 0},
		{10, coordIJK{2, 2, 0}, 3},
		{14, coordIJK{2, 0, 2}, 3},
		{0, coordIJK{0, 2, 2}, 3},
	},
	{ // face 6
		{6, coordIJK{0, 0, 0}, 0},
		{11, coordIJK{2, 2, 0}, 3},
		{10, coordIJK{2, 0, 2}, 3},
		{1, coordIJK{0, 2, 2}, 3},
	},
	{ // face 7
		{7, coordIJK{0, 0, 0}, 0},
		{12, coordIJK{2, 2, 0}, 3},
		{11, coordIJK{2, 0, 2}, 3},
		{2, coordIJK{0, 2, 2}, 3},
	},
	{ // face 8
		{8, coordIJK{0, 0, 0}, 0},
		{13, coordIJK{2, 2, 0}, 3},
		{12, coordIJK{2, 0, 2}, 3},
		{3, coordIJK{0, 2, 2}, 3},
	},
	{ // face 9
		{9, coordIJK{0, 0, 0}, 0},
		{14, coordIJK{2, 2, 0}, 3},
		{13, coordIJK{2, 0, 2}, 3},
		{4, coordIJK{0, 2, 2}, 3},
	},
	{ // face 10
		{10, coordIJK{0, 0, 0}, 0},
		{5, coordIJK{2, 2, 0}, 3},
		{6, coordIJK{2, 0, 2}, 3},
		{15, coordIJK{0, 2, 2}, 3},
	},
	{ // face 11
		{11, coordIJK{0, 0, 0}, 0},
		{6, coordIJK{2, 2, 0}, 3},
		{7, coordIJK{2, 0, 2}, 3},
		{16, coordIJK{0, 2, 2}, 3},
	},
	{ // face 12
		{12, coordIJK{0, 0, 0}, 0},
		{7, coordIJK{2, 2, 0}, 3},
		{8, coordIJK{2, 0, 2}, 3},
		{17, coordIJK{0, 2, 2}, 3},
	},
	{ // face 13
		{13, coordIJK{0, 0, 0}, 0},
		{8, coordIJK{2, 2, 0}, 3},
		{9, coordIJK{2, 0, 2}, 3},
		{18, coordIJK{0, 2, 2}, 3},
	},
	{ // face 14
		{14, coordIJK{0, 0, 0}, 0},
		{9, coordIJK{2, 2, 0}, 3},
		{5, coordIJK{2, 0, 2}, 3},
		{19, coordIJK{0, 2, 2}, 3},
	},
	{ // face 15
		{15, coordIJK{0, 0, 0}, 0},
		{16, coordIJK{2, 0, 2}, 1},
		{19, coordIJK{2, 2, 0}, 5},
		{10, coordIJK{0, 2, 2}, 3},
	},
	{ // face 16
		{16, coordIJK{0, 0, 0}, 0},
		{17, coordIJK{2, 0, 2}, 1},
		{15, coordIJK{2, 2, 0}, 5},
		{11, coordIJK{0, 2, 2}, 3},
	},
	{ // face 17
		{17, coordIJK{0, 0, 0}, 0},
		{18, coordIJK{2, 0, 2}, 1},
		{16, coordIJK{2, 2, 0}, 5},
		{12, coordIJK{0, 2, 2}, 3},
	},
	{ // face 18
		{18, coordIJK{0, 0, 0}, 0},
		{19, coordIJK{2, 0, 2}, 1},
		{17, coordIJK{2, 2, 0}, 5},
		{13, coordIJK{0, 2, 2}, 3},
	},
	{ // face 19
		{19, coordIJK{0, 0, 0}, 0},
		{15, coordIJK{2, 0, 2}, 1},
		{18, coordIJK{2, 2, 0}, 5},
		{14, coordIJK{0, 2, 2}, 3},
	},
}

// maxDimByCIIres is the overage distance table. Entry at index r gives the
// maximum IJK sum for a Class II resolution r. Odd resolutions are -1
// (Class III; not directly used).
var maxDimByCIIres = [17]int{
	2,        // res  0
	-1,       // res  1
	14,       // res  2
	-1,       // res  3
	98,       // res  4
	-1,       // res  5
	686,      // res  6
	-1,       // res  7
	4802,     // res  8
	-1,       // res  9
	33614,    // res 10
	-1,       // res 11
	235298,   // res 12
	-1,       // res 13
	1647086,  // res 14
	-1,       // res 15
	11529602, // res 16
}

// unitScaleByCIIres is the unit scale distance table. Entry at index r gives
// the scale factor for translating to a neighboring face at Class II
// resolution r. Odd resolutions are -1 (Class III; not directly used).
var unitScaleByCIIres = [17]int{
	1,       // res  0
	-1,      // res  1
	7,       // res  2
	-1,      // res  3
	49,      // res  4
	-1,      // res  5
	343,     // res  6
	-1,      // res  7
	2401,    // res  8
	-1,      // res  9
	16807,   // res 10
	-1,      // res 11
	117649,  // res 12
	-1,      // res 13
	823543,  // res 14
	-1,      // res 15
	5764801, // res 16
}

// adjustOverageClassII adjusts a FaceIJK address for a Class II resolution
// when the coordinates exceed the face boundary. It detects overage by
// checking the IJK coordinate sum against the maximum for the resolution,
// then translates to the correct neighboring face if necessary.
//
// pentLeading4 should be true if the cell is a pentagon with a leading digit 4.
// substrate should be true if working with substrate (finer than cell res) grids.
//
// Ported from the H3 C reference (_adjustOverageClassII in faceijk.c).
func adjustOverageClassII(fijk *faceIJK, res int, pentLeading4 bool, substrate bool) Overage {
	overage := NO_OVERAGE

	ijk := &fijk.coord

	// get the maximum dimension value; scale if a substrate grid
	maxDim := maxDimByCIIres[res]
	if substrate {
		maxDim *= 3
	}

	// check for overage
	if substrate && ijk.i+ijk.j+ijk.k == maxDim {
		// on edge
		overage = FACE_EDGE
	} else if ijk.i+ijk.j+ijk.k > maxDim {
		// overage
		overage = NEW_FACE

		var fijkOrient faceOrientIJK
		if ijk.k > 0 {
			if ijk.j > 0 {
				// jk "quadrant"
				fijkOrient = faceNeighbors[fijk.face][ovJK]
			} else {
				// ik "quadrant"
				fijkOrient = faceNeighbors[fijk.face][ovKI]

				// adjust for the pentagonal missing sequence
				if pentLeading4 {
					// translate origin to center of pentagon
					origin := coordIJK{maxDim, 0, 0}
					tmp := ijk.subtract(origin)
					// rotate to adjust for the missing sequence
					tmp = tmp.rotate60cw()
					// translate the origin back to the center of the triangle
					*ijk = tmp.add(origin)
				}
			}
		} else {
			// ij "quadrant"
			fijkOrient = faceNeighbors[fijk.face][ovIJ]
		}

		fijk.face = fijkOrient.face

		// rotate and translate for adjacent face
		for i := 0; i < fijkOrient.ccwRot60; i++ {
			*ijk = ijk.rotate60ccw()
		}

		transVec := fijkOrient.translate
		unitScale := unitScaleByCIIres[res]
		if substrate {
			unitScale *= 3
		}
		transVec = transVec.scale(unitScale)
		*ijk = ijk.add(transVec)
		*ijk = ijk.normalize()

		// overage points on pentagon boundaries can end up on edges
		if substrate && ijk.i+ijk.j+ijk.k == maxDim {
			overage = FACE_EDGE
		}
	}

	return overage
}
