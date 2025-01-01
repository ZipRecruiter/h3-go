package h3

import (
	"math"
	"strconv"
)

const (
	// H3_NUM_BITS is the number of bits in an H3 index.
	H3_NUM_BITS = 64
	// H3_MAX_OFFSET is the max resolution digit in an H3 index.
	H3_MAX_OFFSET = 63
	// H3_MODE_OFFSET is the bit offset of the mode in an H3 index.
	H3_MODE_OFFSET = 59
	// H3_BC_OFFSET is the bit offset of the base cell in an H3 index.
	H3_BC_OFFSET = 45
	// H3_RES_OFFSET is the bit offset of the resolution in an H3 index.
	H3_RES_OFFSET = 52
	// H3_RESERVED_OFFSET is the bit offset of the reserved bits in an H3 index.
	H3_RESERVED_OFFSET = 56
	// H3_PER_DIGIT_OFFSET is the number of bits in a single H3 resolution digit.
	H3_PER_DIGIT_OFFSET = 3
	// H3_HIGH_BIT_MASK is a mask with 1 in the highest bit, 0 elsewhere.
	H3_HIGH_BIT_MASK = uint64(1 << H3_MAX_OFFSET)
	// H3_HIGH_BIT_MASK_NEGATIVE is a mask with 0 in the highest bit, 1 elsewhere.
	H3_HIGH_BIT_MASK_NEGATIVE = ^H3_HIGH_BIT_MASK
	// H3_MODE_MASK is a mask with 1's in the four mode bits, 0 elsewhere.
	H3_MODE_MASK = uint64(15 << H3_MODE_OFFSET)
	// H3_MODE_MASK_NEGATIVE is a mask with 0's in the four mode bits, 1 elsewhere.
	H3_MODE_MASK_NEGATIVE = ^H3_MODE_MASK
	// H3_BC_MASK is a mask with 1's in the seven base cell bits, 0 elsewhere.
	H3_BC_MASK = uint64(127 << H3_BC_OFFSET)
	// H3_BC_MASK_NEGATIVE is a mask with 0's in the seven base cell bits, 1 elsewhere.
	H3_BC_MASK_NEGATIVE = ^H3_BC_MASK
	// H3_RES_MASK is a mask with 1's in the four resolution bits, 0 elsewhere.
	H3_RES_MASK = uint64(15 << H3_RES_OFFSET)
	// H3_RES_MASK_NEGATIVE is a mask with 0's in the four resolution bits, 1 elsewhere.
	H3_RES_MASK_NEGATIVE = ^H3_RES_MASK
	// H3_RESERVED_MASK is a mask with 1's in the three reserved bits, 0 elsewhere.
	H3_RESERVED_MASK = uint64(7 << H3_RESERVED_OFFSET)
	// H3_RESERVED_MASK_NEGATIVE is a mask with 0's in the three reserved bits, 1 elsewhere.
	H3_RESERVED_MASK_NEGATIVE = ^H3_RESERVED_MASK
	// H3_DIGIT_MASK is a mask with 1's in the three bits of resolution 15 digit bits, 0 elsewhere.
	H3_DIGIT_MASK = 7
	// H3_DIGIT_MASK_NEGATIVE is a mask with 0's in the three bits of resolution 15 digit bits, 1 elsewhere.
	H3_DIGIT_MASK_NEGATIVE = ^H3_DIGIT_MASK

	H3_CELL_MODE         = 1
	H3_DIRECTEDEDGE_MODE = 2
	H3_EDGE_MODE         = 3
	H3_VERTEX_MODE       = 4

	// H3_INIT is an H3 index with mode 0, res 0, base cell 0, and 7 for all index
	// digits. Typically used to initialize the creation of an H3 cell index, which
	// expects all direction digits to be 7 beyond the cell's resolution.
	H3_INIT = Cell(35184372088831)
)

// Cell represents a single H3 cell/index.
type Cell uint64

// NewCellFromString creates a new cell from a hex-encoded string.
func NewCellFromString(s string) (Cell, error) {
	i, err := strconv.ParseUint(s, 16, 64)
	if err != nil {
		return 0, err
	}

	return Cell(i), nil
}

func NewCellFromLatLng(ll LatLng, res int) (Cell, error) {
	// Check for valid resolution
	if res < 0 || res > MAX_H3_RES {
		return 0, ErrInvalidArgument
	}

	// Check for valid lat/lng
	if math.IsInf(ll.Latitude(), 0) || math.IsInf(ll.Longitude(), 0) {
		return 0, ErrInvalidArgument
	}

	fijk := geoToFaceIJK(ll, res)
	cell, err := faceIJKToH3(fijk, res)
	if err != nil {
		return 0, err
	}

	return cell, nil
}

func faceIJKToH3(fijk faceIJK, res int) (Cell, error) {
	// Initialize the new H3 index
	h := H3_INIT.setMode(H3_CELL_MODE).setResolution(res)

	// Check for res 0/base cell
	if res == 0 {
		if fijk.coord.i > MAX_FACE_COORD || fijk.coord.j > MAX_FACE_COORD || fijk.coord.k > MAX_FACE_COORD {
			// out of range input
			return 0, ErrInvalidArgument
		}

		h = h.setBaseCell(faceIjkToBaseCell(fijk))
		return h, nil
	}

	// We need to find the correct base cell FaceIJK for this H3 index. Start with
	// the passed in face and resolution res ijk coordinates in that face's
	// coordinate system
	fijkBC := fijk

	// Build the H3Index from the finest res up. Adjust r for the fact that the res 0
	// base cell offsets the indexing digits.
	ijk := fijkBC.coord
	for r := res - 1; r >= 0; r-- {
		lastIJK := ijk
		var lastCenter coordIJK
		if isResolutionClassIII(r + 1) {
			// rotate ccw
			ijk = ijk.upAp7()
			lastCenter = ijk
			lastCenter = lastCenter.downAp7()
		} else {
			// rotate cw
			ijk = ijk.upAp7r()
			lastCenter = ijk
			lastCenter = lastCenter.downAp7r()
		}

		diff := lastIJK.subtract(lastCenter)
		diff = diff.normalize()

		h = h.setIndexDigit(r+1, diff.toDigit())
	}
	fijkBC.coord = ijk

	// fijkBC should now hold the IJK of the base cell in the coordinate system of
	// the current face

	if fijkBC.coord.i > MAX_FACE_COORD || fijkBC.coord.j > MAX_FACE_COORD || fijkBC.coord.k > MAX_FACE_COORD {
		// out of range input
		return 0, ErrInvalidArgument
	}

	// lookup the correct base cell
	bc := faceIjkToBaseCell(fijkBC)
	h = h.setBaseCell(bc)

	// rotate if necessary to get the canonical base cell orientation for this base cell
	numRots := faceIjkToBaseCellCCWrot60(fijkBC)
	if bc.isPentagon() {
		// force rotation out of missing k-axes sub-sequence
		if h.leadingNonZeroDigit() == K_AXES_DIGIT {
			// check for a cw/ccw offset face; default is ccw
			if baseCellIsCwOffset(bc, fijkBC.face) {
				h = h.rotate60cw()
			} else {
				h = h.rotate60ccw()
			}
		}

		for i := 0; i < numRots; i++ {
			h = h.rotatePentagon60ccw()
		}
	} else {
		for i := 0; i < numRots; i++ {
			h = h.rotate60ccw()
		}
	}

	return h, nil
}

func newCell(res int, bc baseCell, initDigit Direction) Cell {
	return H3_INIT.setMode(H3_CELL_MODE).setBaseCell(bc).setResolution(res).setIndexDigit(res, initDigit)
}

// String returns the string representation of the cell as a hex-encoded string.
func (c Cell) String() string {
	return strconv.FormatUint(uint64(c), 16)
}

// Valid returns whether an H3 cell is valid (hexagon or pentagon).
func (c Cell) Valid() bool {
	if c.getHighBit() != 0 {
		return false
	}

	if c.Mode() != H3_CELL_MODE {
		return false
	}

	if c.getReservedBits() != 0 {
		return false
	}

	bc := c.BaseCell()
	if bc < 0 || bc >= NUM_BASE_CELLS {
		return false
	}

	res := c.Resolution()
	if res < 0 || res > MAX_H3_RES {
		return false
	}

	foundFirstNonZeroDigit := false
	for r := 1; r <= res; r++ {
		digit := c.getIndexDigit(r)
		if !foundFirstNonZeroDigit && digit != CENTER_DIGIT {
			foundFirstNonZeroDigit = true

			if bc.isPentagon() && digit == K_AXES_DIGIT {
				return false
			}
		}

		if digit < CENTER_DIGIT || digit >= NUM_DIGITS {
			return false
		}
	}

	for r := res + 1; r <= MAX_H3_RES; r++ {
		if c.getIndexDigit(r) != INVALID_DIGIT {
			return false
		}
	}

	return true
}

// Resolution gets the integer resolution of the H3 index.
func (c Cell) Resolution() int {
	return int((uint64(c) & H3_RES_MASK) >> H3_RES_OFFSET)
}

// getHighBit returns the highest bit of the H3 index.
func (c Cell) getHighBit() int {
	return int((uint64(c) & H3_HIGH_BIT_MASK) >> H3_MAX_OFFSET)
}

// Mode returns the mode of the H3 index.
func (c Cell) Mode() int {
	return int((uint64(c) & H3_MODE_MASK) >> H3_MODE_OFFSET)
}

// getReservedBits returns the reserved bits of the H3 index.
func (c Cell) getReservedBits() int {
	return int((uint64(c) & H3_RESERVED_MASK) >> H3_RESERVED_OFFSET)
}

// BaseCell returns the base cell of the H3 index.
func (c Cell) BaseCell() baseCell {
	return baseCell((uint64(c) & H3_BC_MASK) >> H3_BC_OFFSET)
}

// getIndexDigit gets the `r` integer digit (0-7) of the H3 index.
func (c Cell) getIndexDigit(r int) Direction {
	return Direction((uint64(c) >> ((MAX_H3_RES - r) * H3_PER_DIGIT_OFFSET)) & H3_DIGIT_MASK)
}

// isPentagon returns whether the H3 index is a pentagon.
func (c Cell) isPentagon() bool {
	return c.BaseCell().isPentagon() && c.leadingNonZeroDigit() == 0
}

// leadingNonZeroDigit returns the highest resolution non-zero digit on the cell.
func (c Cell) leadingNonZeroDigit() Direction {
	for r := 1; r <= c.Resolution(); r++ {
		digit := c.getIndexDigit(r)
		if digit != CENTER_DIGIT {
			return digit
		}
	}

	return CENTER_DIGIT
}

// neighborRotations returns the neighbor of the cell in the specified direction.
func (c Cell) neighborRotations(direction Direction, rotations int) (Cell, int, error) {
	current := c
	outRotations := rotations

	if direction < CENTER_DIGIT || direction >= INVALID_DIGIT {
		return current, outRotations, ErrInvalidArgument
	}

	outRotations = outRotations % 6
	for i := 0; i < outRotations; i++ {
		direction = rotate60ccw(direction)
	}

	newRotations := 0
	oldBaseCell := current.BaseCell()
	if oldBaseCell < 0 || oldBaseCell >= NUM_BASE_CELLS {
		return current, newRotations, ErrInvalidArgument
	}
	oldLeadingDigit := current.leadingNonZeroDigit()

	r := current.Resolution() - 1
	for {
		if r == -1 {
			current = current.setBaseCell(baseCellNeighbors[oldBaseCell][direction])
			newRotations = baseCellNeighbor60CCWRots[oldBaseCell][direction]

			if current.BaseCell() == INVALID_BASE_CELL {
				current = current.setBaseCell(baseCellNeighbors[oldBaseCell][IK_AXES_DIGIT])
				newRotations = baseCellNeighbor60CCWRots[oldBaseCell][IK_AXES_DIGIT]

				current = current.rotate60ccw()
				outRotations = outRotations + 1
			}

			break
		} else {
			oldDigit := current.getIndexDigit(r + 1)
			var nextDirection Direction
			if oldDigit == INVALID_DIGIT {
				return current, newRotations, ErrInvalidArgument
			} else if isResolutionClassIII(r + 1) {
				current = current.setIndexDigit(r+1, NEW_DIGIT_II[oldDigit][direction])
				nextDirection = NEW_ADJUSTMENT_II[oldDigit][direction]
			} else {
				current = current.setIndexDigit(r+1, NEW_DIGIT_III[oldDigit][direction])
				nextDirection = NEW_ADJUSTMENT_III[oldDigit][direction]
			}

			if nextDirection != CENTER_DIGIT {
				direction = nextDirection
				r--
			} else {
				// No more adjustments to perform
				break
			}
		}
	}

	newBaseCell := current.BaseCell()
	if newBaseCell.isPentagon() {
		alreadyAdjustedKSubsequence := false

		if current.leadingNonZeroDigit() == K_AXES_DIGIT {
			if oldBaseCell != newBaseCell {
				// in this case, we traversed into the deleted k subsequence of a pentagon base
				// cell. We need to rotate out of that case depending on how we got here. check
				// for a cw/ccw offset face; default is ccw
				if baseCellIsCwOffset(newBaseCell, baseCellData[oldBaseCell].homeFijk.face) {
					current = current.rotate60cw()
				} else {
					current = current.rotate60ccw()
				}
				alreadyAdjustedKSubsequence = true
			} else {
				if oldLeadingDigit == CENTER_DIGIT {
					// Undefined: the k direction is deleted from here
					return current, newRotations, ErrPentagonEncountered
				} else if oldLeadingDigit == JK_AXES_DIGIT {
					// Rotate out of the deleted k subsequence. We also need an additional change to
					// the direction we're moving in.
					current = current.rotate60ccw()
					outRotations = outRotations + 1
				} else if oldLeadingDigit == IK_AXES_DIGIT {
					// Rotate out of the deleted k subsequence. We also need an additional change to
					// the direction we're moving in
					current = current.rotate60cw()
					outRotations = outRotations + 5
				} else {
					// TODO: Should never occur, but is reachable by fuzzer
					return current, newRotations, ErrInvalidArgument
				}
			}
		}

		for i := 0; i < newRotations; i++ {
			current = current.rotate60ccw()
		}

		if oldBaseCell != newBaseCell {
			if newBaseCell.isPolarPentagon() {
				// 'polar' base cells behave differently because they have all i neighbors
				if oldBaseCell != 118 && oldBaseCell != 8 && current.leadingNonZeroDigit() != JK_AXES_DIGIT {
					outRotations = outRotations + 1
				}
			} else if current.leadingNonZeroDigit() == IK_AXES_DIGIT && !alreadyAdjustedKSubsequence {
				// account for distortion introduced to the 5 neighbor by the deleted k
				// subsequence.
				outRotations = outRotations + 1
			}
		}
	} else {
		for i := 0; i < newRotations; i++ {

			current = current.rotate60ccw()
		}
	}

	outRotations = (outRotations + newRotations) % 6
	return current, outRotations, nil
}

// baseCellIsCwOffset returns whether the tested face is a cw offset face.
func baseCellIsCwOffset(bc baseCell, testFace int) bool {
	return baseCellData[bc].cwOffsetPent[0] == testFace || baseCellData[bc].cwOffsetPent[1] == testFace
}

// isResolutionClassIII returns whether the resolution is a class III grid. Odd
// resolutions are class III and even resolutions are class II.
func isResolutionClassIII(r int) bool {
	return r%2 > 0
}

// rotate60ccw rotates the given cell 60 degrees counter-clockwise and returns the new cell.
func (c Cell) rotate60ccw() Cell {
	res := c.Resolution()
	out := c
	for r := 1; r <= res; r++ {
		oldDigit := out.getIndexDigit(r)
		out = out.setIndexDigit(r, rotate60ccw(oldDigit))
	}
	return out
}

// rotate60cw rotates the given cell 60 degrees clockwise and returns the new cell.
func (c Cell) rotate60cw() Cell {
	res := c.Resolution()
	newCell := c
	for r := 1; r <= res; r++ {
		newCell = newCell.setIndexDigit(r, rotate60cw(c.getIndexDigit(r)))
	}
	return newCell
}

// setIndexDigit sets the resolution `res` digit of this cell to the integer `digit` (0-7) and returns the new cell.
func (c Cell) setIndexDigit(res int, digit Direction) Cell {
	return Cell(uint64(c & ^(H3_DIGIT_MASK<<((MAX_H3_RES-(res))*H3_PER_DIGIT_OFFSET))) | (((uint64)(digit)) << ((MAX_H3_RES - (res)) * H3_PER_DIGIT_OFFSET)))
}

// setBaseCell sets the base cell of this cell to the given base cell and returns the new cell.
func (c Cell) setBaseCell(base baseCell) Cell {
	return Cell((uint64(c) & H3_BC_MASK_NEGATIVE) | (uint64(base) << H3_BC_OFFSET))
}

// setMode sets the mode of this cell to the given mode and returns the new cell.
func (c Cell) setMode(mode int) Cell {
	return Cell((uint64(c) & H3_MODE_MASK_NEGATIVE) | (((uint64)(mode)) << H3_MODE_OFFSET))
}

// setResolution sets the resolution of this cell to the given resolution and returns the new cell.
func (c Cell) setResolution(res int) Cell {
	return Cell((uint64(c) & H3_RES_MASK_NEGATIVE) | (((uint64)(res)) << H3_RES_OFFSET))
}

// setHighBit sets the highest bit of this cell to `i` and returns the new cell.
func (c Cell) setHighBit(i int) Cell {
	return Cell(uint64(c)&H3_HIGH_BIT_MASK_NEGATIVE | (uint64(i) << H3_MAX_OFFSET))
}

// setReservedBits sets the reserved bits of this cell to `i` and returns the new cell.
func (c Cell) setReservedBits(i int) Cell {
	return Cell(uint64(c)&H3_RESERVED_MASK_NEGATIVE | (uint64(i) << H3_RESERVED_OFFSET))
}

// rotatePentagon60ccw rotates an H3Index 60 degrees counter-clockwise about a
// pentagonal center and returns the new index.
func (c Cell) rotatePentagon60ccw() Cell {
	out := c
	foundFirstNonZeroDigit := false
	res := c.Resolution()

	for r := 1; r <= res; r++ {
		// rotate this digit
		out = out.setIndexDigit(r, rotate60ccw(out.getIndexDigit(r)))

		// look for the first non-zero digit so we can adjust for deleted k-axes sequence
		// if necessary
		if !foundFirstNonZeroDigit && out.getIndexDigit(r) != 0 {
			foundFirstNonZeroDigit = true

			// adjust for deleted k-axes sequence
			if out.leadingNonZeroDigit() == K_AXES_DIGIT {
				out = out.rotate60ccw()
			}
		}
	}

	return out
}

func (c Cell) isResolutionClassIII() bool {
	return c.Resolution()%2 > 0
}

// rotate60ccw rotates the given digit 60 degrees counter-clockwise and returns the new digit.
func rotate60ccw(digit Direction) Direction {
	switch digit {
	case K_AXES_DIGIT:
		return IK_AXES_DIGIT
	case IK_AXES_DIGIT:
		return I_AXES_DIGIT
	case I_AXES_DIGIT:
		return IJ_AXES_DIGIT
	case IJ_AXES_DIGIT:
		return J_AXES_DIGIT
	case J_AXES_DIGIT:
		return JK_AXES_DIGIT
	case JK_AXES_DIGIT:
		return K_AXES_DIGIT
	default:
		return digit
	}
}

// rotate60cw rotates indexing digit 60 degrees clockwise and returns result.
func rotate60cw(digit Direction) Direction {
	switch digit {
	case K_AXES_DIGIT:
		return JK_AXES_DIGIT
	case JK_AXES_DIGIT:
		return J_AXES_DIGIT
	case J_AXES_DIGIT:
		return IJ_AXES_DIGIT
	case IJ_AXES_DIGIT:
		return I_AXES_DIGIT
	case I_AXES_DIGIT:
		return IK_AXES_DIGIT
	case IK_AXES_DIGIT:
		return K_AXES_DIGIT
	default:
		return digit
	}
}
