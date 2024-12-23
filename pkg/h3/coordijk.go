package h3

const (
	// CENTER_DIGIT is the digit of the center.
	CENTER_DIGIT = 0
	// K_AXES_DIGIT is the digit in the k-axes direction.
	K_AXES_DIGIT = 1
	// J_AXES_DIGIT is the digit in the j-axes direction.
	J_AXES_DIGIT = 2
	// JK_AXES_DIGIT is the digit in the j == k direction.
	JK_AXES_DIGIT = J_AXES_DIGIT | K_AXES_DIGIT // 3
	// I_AXES_DIGIT is the digit in the i-axes direction.
	I_AXES_DIGIT = 4
	// IK_AXES_DIGIT is the digit in the i == k direction.
	IK_AXES_DIGIT = I_AXES_DIGIT | K_AXES_DIGIT // 5
	// IJ_AXES_DIGIT is the digit in the i == j direction.
	IJ_AXES_DIGIT = I_AXES_DIGIT | J_AXES_DIGIT // 6
	// INVALID_DIGIT is the invalid direction.
	INVALID_DIGIT = 7
	// NUM_DIGITS is the number of valid digits.
	NUM_DIGITS = INVALID_DIGIT
)

type coordIJK struct {
	// i is the i component.
	i int
	// j is the j component.
	j int
	// k is the k component.
	k int
}

type faceIJK struct {
	// face is the face number.
	face int
	// coord is the ijk coordinates on the face.
	coord coordIJK
}
