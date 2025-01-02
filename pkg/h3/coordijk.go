package h3

import (
	"math"
)

type Direction int

type coordIJK struct {
	// i is the i component.
	i int
	// j is the j component.
	j int
	// k is the k component.
	k int
}

const (
	// CENTER_DIGIT is the digit of the center.
	CENTER_DIGIT = Direction(0)
	// K_AXES_DIGIT is the digit in the k-axes direction.
	K_AXES_DIGIT = Direction(1)
	// J_AXES_DIGIT is the digit in the j-axes direction.
	J_AXES_DIGIT = Direction(2)
	// JK_AXES_DIGIT is the digit in the j == k direction.
	JK_AXES_DIGIT = J_AXES_DIGIT | K_AXES_DIGIT // 3
	// I_AXES_DIGIT is the digit in the i-axes direction.
	I_AXES_DIGIT = Direction(4)
	// IK_AXES_DIGIT is the digit in the i == k direction.
	IK_AXES_DIGIT = I_AXES_DIGIT | K_AXES_DIGIT // 5
	// IJ_AXES_DIGIT is the digit in the i == j direction.
	IJ_AXES_DIGIT = I_AXES_DIGIT | J_AXES_DIGIT // 6
	// INVALID_DIGIT is the invalid direction.
	INVALID_DIGIT = Direction(7)
	// NUM_DIGITS is the number of valid digits.
	NUM_DIGITS = INVALID_DIGIT

	// NEXT_RING_DIRECTION is the direction used for traversing to the next outward hexagonal ring.
	NEXT_RING_DIRECTION = I_AXES_DIGIT

	// NUM_ICOSA_FACES is the number of icosahedral faces.
	NUM_ICOSA_FACES = 20

	// M_SQRT7 is the square root of 7.
	M_SQRT7 = 2.6457513110645905905016157536392604257102
	//M_RSQRT7 is the inverse square root of 7
	M_RSQRT7 = 0.37796447300922722721451653623418006081576

	// MAX_INT32_3 is the maximum value of a 32-bit integer divided by 3.
	MAX_INT32_3 = math.MaxInt32 / 3
)

var (
	// DIRECTIONS is a list of directions used for traversing a hexagonal ring
	// counterclockwise around {1, 0, 0}
	//
	// <pre>
	//      _
	//    _/ \\_
	//   / \\5/ \\
	//   \\0/ \\4/
	//   / \\_/ \\
	//   \\1/ \\3/
	//     \\2/
	// </pre>
	DIRECTIONS = [6]Direction{
		J_AXES_DIGIT,
		JK_AXES_DIGIT,
		K_AXES_DIGIT,
		IK_AXES_DIGIT,
		I_AXES_DIGIT,
		IJ_AXES_DIGIT,
	}

	// NEW_DIGIT_II lists new when traversing along class II grids.
	//
	// Current digit -> direction -> new digit.
	NEW_DIGIT_II = [7][7]Direction{
		{CENTER_DIGIT, K_AXES_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT, I_AXES_DIGIT,
			IK_AXES_DIGIT, IJ_AXES_DIGIT},
		{K_AXES_DIGIT, I_AXES_DIGIT, JK_AXES_DIGIT, IJ_AXES_DIGIT, IK_AXES_DIGIT,
			J_AXES_DIGIT, CENTER_DIGIT},
		{J_AXES_DIGIT, JK_AXES_DIGIT, K_AXES_DIGIT, I_AXES_DIGIT, IJ_AXES_DIGIT,
			CENTER_DIGIT, IK_AXES_DIGIT},
		{JK_AXES_DIGIT, IJ_AXES_DIGIT, I_AXES_DIGIT, IK_AXES_DIGIT, CENTER_DIGIT,
			K_AXES_DIGIT, J_AXES_DIGIT},
		{I_AXES_DIGIT, IK_AXES_DIGIT, IJ_AXES_DIGIT, CENTER_DIGIT, J_AXES_DIGIT,
			JK_AXES_DIGIT, K_AXES_DIGIT},
		{IK_AXES_DIGIT, J_AXES_DIGIT, CENTER_DIGIT, K_AXES_DIGIT, JK_AXES_DIGIT,
			IJ_AXES_DIGIT, I_AXES_DIGIT},
		{IJ_AXES_DIGIT, CENTER_DIGIT, IK_AXES_DIGIT, J_AXES_DIGIT, K_AXES_DIGIT,
			I_AXES_DIGIT, JK_AXES_DIGIT},
	}

	// NEW_ADJUSTMENT_II is the new traversal direction when traversing along class II grids.
	//
	// Current digit -> direction -> new ap7 move (at coarser level).
	NEW_ADJUSTMENT_II = [7][7]Direction{
		{CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, K_AXES_DIGIT, CENTER_DIGIT, K_AXES_DIGIT, CENTER_DIGIT,
			IK_AXES_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, J_AXES_DIGIT},
		{CENTER_DIGIT, K_AXES_DIGIT, JK_AXES_DIGIT, JK_AXES_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, I_AXES_DIGIT,
			I_AXES_DIGIT, IJ_AXES_DIGIT},
		{CENTER_DIGIT, IK_AXES_DIGIT, CENTER_DIGIT, CENTER_DIGIT, I_AXES_DIGIT,
			IK_AXES_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, J_AXES_DIGIT, CENTER_DIGIT, IJ_AXES_DIGIT,
			CENTER_DIGIT, IJ_AXES_DIGIT},
	}

	// NEW_DIGIT_III is the new traversal direction when traversing along class III grids.
	//
	// Current digit -> direction -> new ap7 move (at coarser level).
	NEW_DIGIT_III = [7][7]Direction{
		{CENTER_DIGIT, K_AXES_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT, I_AXES_DIGIT,
			IK_AXES_DIGIT, IJ_AXES_DIGIT},
		{K_AXES_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT, I_AXES_DIGIT, IK_AXES_DIGIT,
			IJ_AXES_DIGIT, CENTER_DIGIT},
		{J_AXES_DIGIT, JK_AXES_DIGIT, I_AXES_DIGIT, IK_AXES_DIGIT, IJ_AXES_DIGIT,
			CENTER_DIGIT, K_AXES_DIGIT},
		{JK_AXES_DIGIT, I_AXES_DIGIT, IK_AXES_DIGIT, IJ_AXES_DIGIT, CENTER_DIGIT,
			K_AXES_DIGIT, J_AXES_DIGIT},
		{I_AXES_DIGIT, IK_AXES_DIGIT, IJ_AXES_DIGIT, CENTER_DIGIT, K_AXES_DIGIT,
			J_AXES_DIGIT, JK_AXES_DIGIT},
		{IK_AXES_DIGIT, IJ_AXES_DIGIT, CENTER_DIGIT, K_AXES_DIGIT, J_AXES_DIGIT,
			JK_AXES_DIGIT, I_AXES_DIGIT},
		{IJ_AXES_DIGIT, CENTER_DIGIT, K_AXES_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT,
			I_AXES_DIGIT, IK_AXES_DIGIT},
	}

	// NEW_ADJUSTMENT_III is the new traversal direction when traversing along class III grids.
	//
	// Current digit -> direction -> new ap7 move (at coarser level).
	NEW_ADJUSTMENT_III = [7][7]Direction{
		{CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, K_AXES_DIGIT, CENTER_DIGIT, JK_AXES_DIGIT, CENTER_DIGIT,
			K_AXES_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, J_AXES_DIGIT, J_AXES_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, IJ_AXES_DIGIT},
		{CENTER_DIGIT, JK_AXES_DIGIT, J_AXES_DIGIT, JK_AXES_DIGIT, CENTER_DIGIT,
			CENTER_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, CENTER_DIGIT, I_AXES_DIGIT,
			IK_AXES_DIGIT, I_AXES_DIGIT},
		{CENTER_DIGIT, K_AXES_DIGIT, CENTER_DIGIT, CENTER_DIGIT, IK_AXES_DIGIT,
			IK_AXES_DIGIT, CENTER_DIGIT},
		{CENTER_DIGIT, CENTER_DIGIT, IJ_AXES_DIGIT, CENTER_DIGIT, I_AXES_DIGIT,
			CENTER_DIGIT, IJ_AXES_DIGIT},
	}

	faceCenterPoint = [NUM_ICOSA_FACES]vec3d{
		{0.2199307791404606, 0.6583691780274996, 0.7198475378926182},    // face  0
		{-0.2139234834501421, 0.1478171829550703, 0.9656017935214205},   // face  1
		{0.1092625278784797, -0.4811951572873210, 0.8697775121287253},   // face  2
		{0.7428567301586791, -0.3593941678278028, 0.5648005936517033},   // face  3
		{0.8112534709140969, 0.3448953237639384, 0.4721387736413930},    // face  4
		{-0.1055498149613921, 0.9794457296411413, 0.1718874610009365},   // face  5
		{-0.8075407579970092, 0.1533552485898818, 0.5695261994882688},   // face  6
		{-0.2846148069787907, -0.8644080972654206, 0.4144792552473539},  // face  7
		{0.7405621473854482, -0.6673299564565524, -0.0789837646326737},  // face  8
		{0.8512303986474293, 0.4722343788582681, -0.2289137388687808},   // face  9
		{-0.7405621473854481, 0.6673299564565524, 0.0789837646326737},   // face 10
		{-0.8512303986474292, -0.4722343788582682, 0.2289137388687808},  // face 11
		{0.1055498149613919, -0.9794457296411413, -0.1718874610009365},  // face 12
		{0.8075407579970092, -0.1533552485898819, -0.5695261994882688},  // face 13
		{0.2846148069787908, 0.8644080972654204, -0.4144792552473539},   // face 14
		{-0.7428567301586791, 0.3593941678278027, -0.5648005936517033},  // face 15
		{-0.8112534709140971, -0.3448953237639382, -0.4721387736413930}, // face 16
		{-0.2199307791404607, -0.6583691780274996, -0.7198475378926182}, // face 17
		{0.2139234834501420, -0.1478171829550704, -0.9656017935214205},  // face 18
		{-0.1092625278784796, 0.4811951572873210, -0.8697775121287253},  // face 19
	}

	// faceAxesAzRadsCII is icosahedron face ijk axes as azimuth in radians from face
	// center to vertex 0/1/2 respectively
	faceAxesAzRadsCII = [NUM_ICOSA_FACES][3]float64{
		{5.619958268523939882, 3.525563166130744542,
			1.431168063737548730}, // face  0
		{5.760339081714187279, 3.665943979320991689,
			1.571548876927796127}, // face  1
		{0.780213654393430055, 4.969003859179821079,
			2.874608756786625655}, // face  2
		{0.430469363979999913, 4.619259568766391033,
			2.524864466373195467}, // face  3
		{6.130269123335111400, 4.035874020941915804,
			1.941478918548720291}, // face  4
		{2.692877706530642877, 0.598482604137447119,
			4.787272808923838195}, // face  5
		{2.982963003477243874, 0.888567901084048369,
			5.077358105870439581}, // face  6
		{3.532912002790141181, 1.438516900396945656,
			5.627307105183336758}, // face  7
		{3.494305004259568154, 1.399909901866372864,
			5.588700106652763840}, // face  8
		{3.003214169499538391, 0.908819067106342928,
			5.097609271892733906}, // face  9
		{5.930472956509811562, 3.836077854116615875,
			1.741682751723420374}, // face 10
		{0.138378484090254847, 4.327168688876645809,
			2.232773586483450311}, // face 11
		{0.448714947059150361, 4.637505151845541521,
			2.543110049452346120}, // face 12
		{0.158629650112549365, 4.347419854898940135,
			2.253024752505744869}, // face 13
		{5.891865957979238535, 3.797470855586042958,
			1.703075753192847583}, // face 14
		{2.711123289609793325, 0.616728187216597771,
			4.805518392002988683}, // face 15
		{3.294508837434268316, 1.200113735041072948,
			5.388903939827463911}, // face 16
		{3.804819692245439833, 1.710424589852244509,
			5.899214794638635174}, // face 17
		{3.664438879055192436, 1.570043776661997111,
			5.758833981448388027}, // face 18
		{2.361378999196363184, 0.266983896803167583,
			4.455774101589558636}, // face 19
	}

	faceCenterGeo = [NUM_ICOSA_FACES]LatLng{
		{0.803582649718989942, 1.248397419617396099},   // face  0
		{1.307747883455638156, 2.536945009877921159},   // face  1
		{1.054751253523952054, -1.347517358900396623},  // face  2
		{0.600191595538186799, -0.450603909469755746},  // face  3
		{0.491715428198773866, 0.401988202911306943},   // face  4
		{0.172745327415618701, 1.678146885280433686},   // face  5
		{0.605929321571350690, 2.953923329812411617},   // face  6
		{0.427370518328979641, -1.888876200336285401},  // face  7
		{-0.079066118549212831, -0.733429513380867741}, // face  8
		{-0.230961644455383637, 0.506495587332349035},  // face  9
		{0.079066118549212831, 2.408163140208925497},   // face 10
		{0.230961644455383637, -2.635097066257444203},  // face 11
		{-0.172745327415618701, -1.463445768309359553}, // face 12
		{-0.605929321571350690, -0.187669323777381622}, // face 13
		{-0.427370518328979641, 1.252716453253507838},  // face 14
		{-0.600191595538186799, 2.690988744120037492},  // face 15
		{-0.491715428198773866, -2.739604450678486295}, // face 16
		{-0.803582649718989942, -1.893195233972397139}, // face 17
		{-1.307747883455638156, -0.604647643711872080}, // face 18
		{-1.054751253523952054, 1.794075294689396615},  // face 19
	}

	// UNIT_VECS is a list of the CoordIJK unit vectors corresponding to the 7 H3
	// digits.
	UNIT_VECS = [7]coordIJK{
		{0, 0, 0}, // direction 0
		{0, 0, 1}, // direction 1
		{0, 1, 0}, // direction 2
		{0, 1, 1}, // direction 3
		{1, 0, 0}, // direction 4
		{1, 0, 1}, // direction 5
		{1, 1, 0}, // direction 6
	}
)

// normalize normalizes ijk coordinates by setting the components to the smallest
// possible values and returns the result.
//
// This function does not protect against signed integer overflow. The caller
// must ensure that none of (i - j), (i - k), (j - i), (j - k), (k - i), (k - j)
// will overflow. This function may be changed in the future to make that check
// itself and return an error code.
func (c coordIJK) normalize() coordIJK {
	out := c

	// remove any negative values
	if out.i < 0 {
		out.j -= out.i
		out.k -= out.i
		out.i = 0
	}

	if out.j < 0 {
		out.i -= out.j
		out.k -= out.j
		out.j = 0
	}

	if out.k < 0 {
		out.i -= out.k
		out.j -= out.k
		out.k = 0
	}

	// remove the min value if needed
	minVal := out.i
	if out.j < minVal {
		minVal = out.j
	}
	if out.k < minVal {
		minVal = out.k
	}
	if minVal > 0 {
		out.i -= minVal
		out.j -= minVal
		out.k -= minVal
	}

	return out
}

// normalizeCouldOverflow returns true if normalize with the given input could
// have a signed integer overflow. Assumes k is set to 0.
func (c coordIJK) normalizeCouldOverflow() bool {
	// Check for the possibility of overflow
	var maxVal, minVal int
	if c.i > c.j {
		maxVal = c.i
		minVal = c.j
	} else {
		maxVal = c.j
		minVal = c.i
	}
	if minVal < 0 {
		// Only if the minVal is less than 0 will the resulting number be larger
		// than maxVal. If minVal is positive, then maxVal is also positive, and a
		// positive signed integer minus another positive signed integer will
		// not overflow.
		if addInt32sWouldOverflow(maxVal, minVal) {
			// maxVal + minVal would overflow
			return true
		}
		if subInt32sWouldOverflow(0, minVal) {
			// 0 - INT32_MIN would overflow
			return true
		}
		if subInt32sWouldOverflow(maxVal, minVal) {
			// maxVal - minVal would overflow
			return true
		}
	}
	return false
}

type coordIJ struct {
	// i is the i component.
	i int
	// j is the j component.
	j int
}

// toIjk converts the i, j coordinates to ijk coordinates and returns the result.
// It will return an error if signed integer overflow would have occurred.
func (c coordIJ) toIjk() (coordIJK, error) {
	ijk := coordIJK{
		i: c.i,
		j: c.j,
		k: 0,
	}

	if ijk.normalizeCouldOverflow() {
		return coordIJK{}, ErrInvalidArgument
	}

	return ijk.normalize(), nil
}

// toCube converts the ijk coordinates to cube coordinates and returns the result.
func (c coordIJK) toCube() coordIJK {
	return coordIJK{
		i: -c.i + c.k,
		j: c.j - c.k,
		k: -c.i - c.j,
	}
}

// NewCoordIJKFromCube converts the cube coordinates to ijk coordinates and
// returns the result.
func NewCoordIJKFromCube(c coordIJK) coordIJK {
	ijk := coordIJK{
		i: -c.i,
		j: c.j,
		k: 0,
	}
	return ijk.normalize()
}

// toIj converts the ijk coordinates to i, j coordinates and returns the result.
func (c coordIJK) toIj() coordIJ {
	return coordIJ{
		i: c.i - c.k,
		j: c.j - c.k,
	}
}

// ijkToHex2d finds the center point in 2D cartesian coordinates of a hex.
func ijkToHex2d(h coordIJK) vec2d {
	i := h.i - h.k
	j := h.j - h.k

	return vec2d{
		x: float64(i) - 0.5*float64(j),
		y: float64(j) * M_SQRT3_2,
	}
}

// matches returns whether two ijk coordinates contain exactly the same component
// values.
func (c coordIJK) matches(o coordIJK) bool {
	return c.i == o.i && c.j == o.j && c.k == o.k
}

// subtract subtracts the other ijk coordinates from the current ijk coordinates
// and returns the difference.
func (c coordIJK) subtract(other coordIJK) coordIJK {
	return coordIJK{
		i: c.i - other.i,
		j: c.j - other.j,
		k: c.k - other.k,
	}
}

// scale uniformly scales current ijk coordinates by a scalar `s` and returns the result.
func (c coordIJK) scale(s int) coordIJK {
	return coordIJK{
		i: c.i * s,
		j: c.j * s,
		k: c.k * s,
	}
}

// add adds the other ijk coordinates to the current ijk coordinates and returns the sum.
func (c coordIJK) add(other coordIJK) coordIJK {
	return coordIJK{
		i: c.i + other.i,
		j: c.j + other.j,
		k: c.k + other.k,
	}
}

type faceIJK struct {
	// face is the face number.
	face int
	// coord is the ijk coordinates on the face.
	coord coordIJK
}

// geoToFaceIJK encodes a coordinate on the sphere to the FaceIJK address
// of the containing cell at the specified resolution.
func geoToFaceIJK(ll LatLng, res int) faceIJK {
	v := geoToHex2d(ll, res)
	c := hex2dToCoordIJK(v)

	closestFace, _ := geoToClosestFace(ll)

	f := faceIJK{
		face:  closestFace,
		coord: c,
	}

	return f
}

// hex2dToCoordIJK returns the containing hex in ijk+ coordinates for a 2D
// cartesian coordinate vector (from DGGRID).
func hex2dToCoordIJK(v vec2d) coordIJK {
	h := coordIJK{}
	h.k = 0.0

	a1 := math.Abs(v.x)
	a2 := math.Abs(v.y)

	// first do a reverse conversion
	x2 := a2 * M_RSIN60
	x1 := a1 + x2/2.0

	// check if we have the center of a hex
	m1 := int(x1)
	m2 := int(x2)

	// otherwise round correctly
	r1 := x1 - float64(m1)
	r2 := x2 - float64(m2)

	if r1 < 0.5 {
		if r1 < 1.0/3.0 {
			if r2 < (1.0+r1)/2.0 {
				h.i = m1
				h.j = m2
			} else {
				h.i = m1
				h.j = m2 + 1
			}
		} else {
			if r2 < (1.0 - r1) {
				h.j = m2
			} else {
				h.j = m2 + 1
			}

			if (1.0-r1) <= r2 && r2 < (2.0*r1) {
				h.i = m1 + 1
			} else {
				h.i = m1
			}
		}
	} else {
		if r1 < 2.0/3.0 {
			if r2 < (1.0 - r1) {
				h.j = m2
			} else {
				h.j = m2 + 1
			}

			if (2.0*r1-1.0) < r2 && r2 < (1.0-r1) {
				h.i = m1
			} else {
				h.i = m1 + 1
			}
		} else {
			if r2 < (r1 / 2.0) {
				h.i = m1 + 1
				h.j = m2
			} else {
				h.i = m1 + 1
				h.j = m2 + 1
			}
		}
	}

	// now fold across the axes if necessary

	if v.x < 0.0 {
		if h.j%2 == 0 {
			// even
			axisI := h.j / 2
			diff := h.i - axisI
			h.i = h.i - 2.0*diff
		} else {
			// odd
			axisI := (h.j + 1) / 2
			diff := h.i - axisI
			h.i = h.i - (2.0*diff + 1)
		}
	}

	if v.y < 0.0 {
		h.i = h.i - (2*h.j+1)/2
		h.j = -1 * h.j
	}

	h = h.normalize()

	return h
}

// geoToHex2d encodes a coordinate on the sphere to the corresponding icosahedral
// face and containing 2D hex coordinates relative to that face center.
func geoToHex2d(ll LatLng, res int) vec2d {
	// determine the icosahedron face
	face, sqd := geoToClosestFace(ll)

	r := math.Acos(1.0 - sqd*0.5)

	if r < EPSILON {
		return vec2d{0.0, 0.0}
	}

	// now have face and r, now find CCW theta from CII i-axis
	theta := posAngleRads(faceAxesAzRadsCII[face][0] - posAngleRads(faceCenterGeo[face].geoAzimuthRads(ll)))

	// adjust theta for Class III (odd resolutions)
	if isResolutionClassIII(res) {
		theta = posAngleRads(theta - M_AP7_ROT_RADS)
	}

	// perform gnomonic scaling of r
	r = math.Tan(r)

	// scale for current resolution length u
	r *= INV_RES0_U_GNOMONIC
	for i := 0; i < res; i++ {
		r *= M_SQRT7
	}

	// we now have (r, theta) in hex2d with theta ccw from x-axes

	return vec2d{
		x: r * math.Cos(theta),
		y: r * math.Sin(theta),
	}
}

// geoToClosestFace encodes a coordinate on the sphere to the corresponding
// icosahedral face and containing the squared Euclidean distance to that face
// center.
//
// Returns the icosahedral face number and the squared Euclidean distance to the
// face center.
func geoToClosestFace(ll LatLng) (int, float64) {
	v3d := newVec3dFromLatLng(ll)

	// determine the icosahedron face
	face := 0
	// The distance between two farthest points is 2.0, therefore the square of the
	// distance between two points should always be less or equal than 4.0.
	sqd := 5.0
	for f := 0; f < NUM_ICOSA_FACES; f++ {
		sqdT := faceCenterPoint[f].pointSquareDistance(v3d)
		if sqdT < sqd {
			face = f
			sqd = sqdT
		}
	}

	return face, sqd
}

// toDigit determines the H3 digit corresponding to a unit vector or the
// zero vector in ijk coordinates.
//
// Accepts the ijk coordinates; must be a unit vector or zero vector. Returns the
// H3 digit (0-6) corresponding to the ijk unit vector, zero vector, or
// INVALID_DIGIT (7) on failure.
func (c coordIJK) toDigit() Direction {
	u := c.normalize()

	digit := INVALID_DIGIT
	for i := CENTER_DIGIT; i < NUM_DIGITS; i++ {
		if u.matches(UNIT_VECS[i]) {
			digit = i
			break
		}
	}

	return digit
}

// neighbor returns the neighbor of this coordIJK in the direction dir.
func (c coordIJK) neighbor(dir Direction) coordIJK {
	if dir <= CENTER_DIGIT || dir >= NUM_DIGITS {
		return c
	}

	return c.add(UNIT_VECS[dir]).normalize()
}

// downAp7r finds the normalized ijk coordinates of the hex centered on the
// indicated hex at the next finer aperture 7 clockwise resolution.
func (c coordIJK) downAp7r() coordIJK {
	// res r unit vectors in res r+1
	iVec := coordIJK{3, 1, 0}
	jVec := coordIJK{0, 3, 1}
	kVec := coordIJK{1, 0, 3}

	iVec = iVec.scale(c.i)
	jVec = jVec.scale(c.j)
	kVec = kVec.scale(c.k)

	out := iVec.add(jVec)
	out = out.add(kVec)

	out = out.normalize()
	return out
}

// Find the normalized ijk coordinates of the hex centered on the indicated hex
// at the next finer aperture 7 counter-clockwise resolution.
func (c coordIJK) downAp7() coordIJK {
	// res r unit vectors in res r+1
	iVec := coordIJK{3, 0, 1}
	jVec := coordIJK{1, 3, 0}
	kVec := coordIJK{0, 1, 3}

	iVec = iVec.scale(c.i)
	jVec = jVec.scale(c.j)
	kVec = kVec.scale(c.k)

	out := iVec.add(jVec)
	out = out.add(kVec)

	out = out.normalize()
	return out
}

// upAp7r finds the normalized ijk coordinates of the indexing parent of a cell
// in a clockwise aperture 7 grid.
func (c coordIJK) upAp7r() coordIJK {
	// convert to coordIJ
	i := c.i - c.k
	j := c.j - c.k

	o := coordIJK{
		i: int(math.Round(float64(2*i+j) * M_ONESEVENTH)),
		j: int(math.Round(float64(3*j-i) * M_ONESEVENTH)),
		k: 0,
	}
	return o.normalize()
}

// Find the normalized ijk coordinates of the indexing parent of a cell in a
// counter-clockwise aperture 7 grid.
func (c coordIJK) upAp7() coordIJK {
	// convert to CoordIJ
	i := c.i - c.k
	j := c.j - c.k

	o := coordIJK{
		i: int(math.Round(float64(3*i-j) * M_ONESEVENTH)),
		j: int(math.Round(float64(i+2*j) * M_ONESEVENTH)),
		k: 0,
	}
	return o.normalize()
}

func (c coordIJK) upAp7Checked() (coordIJK, error) {
	i := c.i - c.k
	j := c.j - c.k

	// <0 is checked because the input must all be non-negative, but some
	// negative inputs are used in unit tests to exercise the below.
	if i >= MAX_INT32_3 || j >= MAX_INT32_3 || i < 0 || j < 0 {
		if addInt32sWouldOverflow(i, i) {
			return coordIJK{}, ErrInvalidArgument
		}
		i2 := i + i
		if addInt32sWouldOverflow(i2, i) {
			return coordIJK{}, ErrInvalidArgument
		}
		i3 := i2 + i
		if addInt32sWouldOverflow(j, j) {
			return coordIJK{}, ErrInvalidArgument
		}
		j2 := j + j

		if subInt32sWouldOverflow(i3, j) {
			return coordIJK{}, ErrInvalidArgument
		}
		if addInt32sWouldOverflow(i, j2) {
			return coordIJK{}, ErrInvalidArgument
		}
	}

	o := coordIJK{
		i: int(math.Round(((float64(i) * 3.0) - float64(j)) * M_ONESEVENTH)),
		j: int(math.Round((float64(i) + (float64(j) * 2)) * M_ONESEVENTH)),
		k: 0,
	}

	if o.normalizeCouldOverflow() {
		return coordIJK{}, ErrInvalidArgument
	}

	return o.normalize(), nil
}

func (c coordIJK) upAp7rChecked() (coordIJK, error) {
	i := c.i - c.k
	j := c.j - c.k

	if i >= MAX_INT32_3 || j >= MAX_INT32_3 || i < 0 || j < 0 {
		if addInt32sWouldOverflow(i, i) {
			return coordIJK{}, ErrInvalidArgument
		}
		i2 := i + i
		if addInt32sWouldOverflow(j, j) {
			return coordIJK{}, ErrInvalidArgument
		}
		j2 := j + j
		if addInt32sWouldOverflow(j2, j) {
			return coordIJK{}, ErrInvalidArgument
		}
		j3 := j2 + j

		if addInt32sWouldOverflow(i2, j) {
			return coordIJK{}, ErrInvalidArgument
		}
		if subInt32sWouldOverflow(j3, i) {
			return coordIJK{}, ErrInvalidArgument
		}
	}

	o := coordIJK{
		i: int(math.Round(((float64(i) * 2.0) + float64(j)) * M_ONESEVENTH)),
		j: int(math.Round(((float64(j) * 3) - float64(i)) * M_ONESEVENTH)),
		k: 0,
	}

	if o.normalizeCouldOverflow() {
		return coordIJK{}, ErrInvalidArgument
	}

	return o.normalize(), nil
}
