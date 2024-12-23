package h3

import "strconv"

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

// String returns the string representation of the cell as a hex-encoded string.
func (c Cell) String() string {
	return strconv.FormatUint(uint64(c), 16)
}

// Valid returns whether an H3 cell is valid (hexagon or pentagon).
func (c Cell) Valid() bool {
	if c.getHighBit() != 0 {
		return false
	}

	if c.getMode() != H3_CELL_MODE {
		return false
	}

	if c.getReservedBits() != 0 {
		return false
	}

	bc := c.getBaseCell()
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

// getMode returns the mode of the H3 index.
func (c Cell) getMode() int {
	return int((uint64(c) & H3_MODE_MASK) >> H3_MODE_OFFSET)
}

// getReservedBits returns the reserved bits of the H3 index.
func (c Cell) getReservedBits() int {
	return int((uint64(c) & H3_RESERVED_MASK) >> H3_RESERVED_OFFSET)
}

// getBaseCell returns the base cell of the H3 index.
func (c Cell) getBaseCell() baseCell {
	return baseCell((uint64(c) & H3_BC_MASK) >> H3_BC_OFFSET)
}

// getIndexDigit gets the `r` integer digit (0-7) of the H3 index.
func (c Cell) getIndexDigit(r int) int {
	return int((uint64(c) >> ((MAX_H3_RES - r) * H3_PER_DIGIT_OFFSET)) & H3_DIGIT_MASK)
}
