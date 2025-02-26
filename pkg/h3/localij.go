package h3

import "fmt"

var (
	ErrResolutionMismatch = fmt.Errorf("cells have different resolutions")

	// FAILED_DIRECTIONS is a lookup table that indicates prohibited directions when unfolding a pentagon.
	//
	// Indexes by two directions, both relative to the pentagon base cell. The first
	// is the direction of the origin index and the second is the direction of the
	// index to unfold. Direction refers to the direction from base cell to base cell
	// if the indexes are on different base cells, or the leading digit if within the
	// pentagon base cell.
	//
	// This previously included a Class II/Class III check but these were removed due
	// to failure cases. It's possible this could be restricted to a narrower set of
	// a failure cases. Currently, the logic is any unfolding across more than one
	// icosahedron face is not permitted.
	FAILED_DIRECTIONS = [7][7]bool{
		{false, false, false, false, false, false, false}, // 0
		{false, false, false, false, false, false, false}, // 1
		{false, false, false, false, true, true, false},   // 2
		{false, false, false, false, true, false, true},   // 3
		{false, false, true, true, false, false, false},   // 4
		{false, false, true, false, false, false, true},   // 5
		{false, false, false, true, false, true, false},   // 6
	}

	// PENTAGON_ROTATIONS is origin leading digit -> index leading digit -> rotations 60 clockwise.
	PENTAGON_ROTATIONS = [7][7]int{
		{0, -1, 0, 0, 0, 0, 0},       // 0
		{-1, -1, -1, -1, -1, -1, -1}, // 1
		{0, -1, 0, 0, 0, 1, 0},       // 2
		{0, -1, 0, 0, 1, 1, 0},       // 3
		{0, -1, 0, 5, 0, 0, 0},       // 4
		{0, -1, 5, 5, 0, 0, 0},       // 5
		{0, -1, 0, 0, 0, 0, 0},       // 6
	}
)

func (c Cell) toLocalIJK(other Cell) (coordIJK, error) {
	originRes := c.Resolution()
	otherRes := other.Resolution()

	if originRes != otherRes {
		return coordIJK{}, ErrResolutionMismatch
	}

	originBaseCell := c.BaseCell()

	if originBaseCell < 0 || originBaseCell >= NUM_BASE_CELLS {
		return coordIJK{}, ErrInvalidArgument
	}

	otherBaseCell := other.BaseCell()
	if otherBaseCell < 0 || otherBaseCell >= NUM_BASE_CELLS {
		return coordIJK{}, ErrInvalidArgument
	}

	// Direction from origin base cell to index base cell
	dir := CENTER_DIGIT
	revDir := CENTER_DIGIT
	if originBaseCell != otherBaseCell {
		dir = originBaseCell.baseCellDirection(otherBaseCell)
		if dir == INVALID_DIGIT {
			return coordIJK{}, ErrInvalidArgument
		}

		revDir = otherBaseCell.baseCellDirection(originBaseCell)
		if revDir == INVALID_DIGIT {
			return coordIJK{}, ErrInvalidArgument
		}
	}

	originOnPentagon := originBaseCell.isPentagon()
	indexOnPentagon := otherBaseCell.isPentagon()

	cellOut := other

	indexFijk := faceIJK{}
	if dir != CENTER_DIGIT {
		// Rotate index into the orientation of the origin base cell.
		// Clockwise because we are undoing the rotation into that base cell.
		baseCellRotations := baseCellNeighbor60CCWRots[originBaseCell][dir]
		if indexOnPentagon {
			for i := 0; i < baseCellRotations; i++ {
				cellOut = cellOut.rotatePentagon60cw()

				revDir = rotate60cw(revDir)
				if revDir == K_AXES_DIGIT {
					revDir = rotate60cw(revDir)
				}
			}
		} else {
			for i := 0; i < baseCellRotations; i++ {
				cellOut = cellOut.rotate60cw()
				revDir = rotate60cw(revDir)
			}
		}
	}
	// Face is unused. This produces coordinates in base cell coordinate space.
	indexFijk, _ = cellOut.toFaceIjkWithInitializedFijk(indexFijk)

	if dir != CENTER_DIGIT {
		if otherBaseCell == originBaseCell {
			return coordIJK{}, ErrInvalidArgument
		}

		if originOnPentagon && indexOnPentagon {
			return coordIJK{}, ErrInvalidArgument
		}

		pentagonRotations := 0
		directionRotations := 0

		if originOnPentagon {
			originLeadingDigit := c.leadingNonZeroDigit()

			if originLeadingDigit == INVALID_DIGIT {
				return coordIJK{}, ErrInvalidArgument
			}

			if FAILED_DIRECTIONS[originLeadingDigit][dir] {
				// TODO: We may be unfolding the pentagon incorrectly in this case; return an
				// error code until this is guaranteed to be correct.
				return coordIJK{}, ErrInvalidArgument
			}

			directionRotations = PENTAGON_ROTATIONS[originLeadingDigit][dir]
			pentagonRotations = directionRotations
		} else if indexOnPentagon {
			indexLeadingDigit := cellOut.leadingNonZeroDigit()

			if indexLeadingDigit == INVALID_DIGIT {
				return coordIJK{}, ErrInvalidArgument
			}

			if FAILED_DIRECTIONS[indexLeadingDigit][revDir] {
				// TODO: We may be unfolding the pentagon incorrectly in this case; return an
				// error code until this is guaranteed to be correct.
				return coordIJK{}, ErrInvalidArgument
			}

			pentagonRotations = PENTAGON_ROTATIONS[revDir][indexLeadingDigit]
		}

		if pentagonRotations < 0 || directionRotations < 0 {
			// This occurs when an invalid K axis digit is present
			return coordIJK{}, ErrInvalidArgument
		}

		for i := 0; i < pentagonRotations; i++ {
			indexFijk.coord = indexFijk.coord.rotate60cw()
		}

		offset := coordIJK{}
		offset = offset.neighbor(dir)

		// Scale offset based on resolution
		for r := originRes - 1; r >= 0; r-- {
			if isResolutionClassIII(r + 1) {
				// rotate ccw
				offset = offset.downAp7()
			} else {
				// rotate cw
				offset = offset.downAp7r()
			}
		}

		for i := 0; i < directionRotations; i++ {
			offset = offset.rotate60cw()
		}

		// Perform necessary translation
		indexFijk.coord = indexFijk.coord.add(offset)
		indexFijk.coord = indexFijk.coord.normalize()
	} else if originOnPentagon && indexOnPentagon {
		// If the origin and index are on pentagon, and we checked that the base cells
		// are the same or neighboring, then they must be the same base cell.
		if originBaseCell != otherBaseCell {
			return coordIJK{}, ErrInvalidArgument
		}

		originLeadingDigit := c.leadingNonZeroDigit()
		indexLeadingDigit := cellOut.leadingNonZeroDigit()

		if originLeadingDigit == INVALID_DIGIT || indexLeadingDigit == INVALID_DIGIT {
			return coordIJK{}, ErrInvalidArgument
		}

		if FAILED_DIRECTIONS[originLeadingDigit][indexLeadingDigit] {
			// TODO We may be unfolding the pentagon incorrectly in this case; return an
			// error code until this is guaranteed to be correct.
			return coordIJK{}, ErrInvalidArgument
		}

		withinPentagonRotations := PENTAGON_ROTATIONS[originLeadingDigit][indexLeadingDigit]

		for i := 0; i < withinPentagonRotations; i++ {
			indexFijk.coord = indexFijk.coord.rotate60cw()
		}
	}

	return indexFijk.coord, nil
}
