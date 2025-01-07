package h3

import (
	"errors"
)

const (
	// K_ALL_CELLS_AT_RES_15 is the k value that encompass all cells at resolution
	// 15. This is the largest possible k in the H3 grid system.
	K_ALL_CELLS_AT_RES_15 = 13780510
)

// Maximum number of cells that result from the GridDisk algorithm with the
// given k. Formula source and proof: https://oeis.org/A003215
func maxGridDiskSize(k int) (int, error) {
	if k < 0 {
		return 0, ErrInvalidArgument
	}

	if k >= K_ALL_CELLS_AT_RES_15 {
		// If a k value of this value or above is provided, this function will estimate
		// more cells than exist in the H3 grid at the finest resolution. This is a
		// problem since the function does signed integer arithmetic on `k`, which could
		// overflow. To prevent that, instead substitute the maximum number of cells in
		// the grid, as it should not be possible for the GridDisk functions to exceed
		// that. Note this is not resolution specific. So, when resolution < 15, this
		// function may still estimate a size larger than the number of cells in the
		// grid.
		return getNumCellsAtResolution(MAX_H3_RES), nil
	}

	return 3*k*(k+1) + 1, nil
}

// gridDiskDistancesSafe is the safe but slow version of GridDiskDistances (also
// called by it when needed).
//
// Includes the origin cell in the output list (treating it as a hash set) and
// recurses to its neighbors, if needed.
func (c Cell) gridDiskDistancesSafe(k int) ([]Cell, []int, error) {
	maxIdx, err := maxGridDiskSize(k)
	if err != nil {
		return nil, nil, err
	}

	out := make([]Cell, maxIdx)
	distances := make([]int, maxIdx)
	err = c.gridDiskDistancesInternal(k, &out, &distances, maxIdx, 0)
	return out, distances, err
}

// gridDiskDistancesUnsafe produces indexes within k distance of the origin
// index. Output behavior is undefined when one of the indexes returned by this
// function is a pentagon or is in the pentagon distortion area.
//
// k-ring 0 is defined as the origin index, k-ring 1 is defined as k-ring 0 and
// all neighboring indexes, and so on.
//
// The cells are returned in order of increasing distance from the origin. The
// second return value is a list of distances from the origin index.
func (c Cell) gridDiskDistancesUnsafe(k int) ([]Cell, []int, error) {
	if k < 0 {
		return nil, nil, ErrInvalidArgument
	}

	gridDiskSize, err := maxGridDiskSize(k)
	if err != nil {
		return nil, nil, err
	}

	origin := c
	cells := make([]Cell, gridDiskSize)
	distances := make([]int, gridDiskSize)

	idx := 0
	cells[idx] = origin
	distances[idx] = 0
	idx++

	if origin.isPentagon() {
		return cells, distances, ErrPentagonEncountered
	}

	// current ring, 0 < ring <= k
	ring := 1
	// current side of the ring, 0 <= direction < 6
	direction := 0
	// current position on the side of the ring, 0 <= i < ring
	i := 0
	// number of 60 degree counterclockwise rotations to perform on the direction
	// (based on which faces are crossed)
	rotations := 0

	for ring <= k {
		var neighborErr error
		if direction == 0 && i == 0 {
			origin, rotations, neighborErr = origin.neighborRotations(NEXT_RING_DIRECTION, rotations)
			if neighborErr != nil {
				return cells, distances, neighborErr
			}

			if origin.isPentagon() {
				return cells, distances, ErrPentagonEncountered
			}
		}

		origin, rotations, neighborErr = origin.neighborRotations(DIRECTIONS[direction], rotations)
		if neighborErr != nil {
			return cells, distances, neighborErr
		}

		cells[idx] = origin
		distances[idx] = ring
		idx++

		i++
		if i == ring {
			i = 0
			direction++
			if direction == 6 {
				direction = 0
				ring++
			}
		}

		if origin.isPentagon() {
			return cells, distances, ErrPentagonEncountered
		}
	}

	return cells, distances, nil
}

// gridDiskUnsafe produces indexes within k distance of the origin index.
// Output behavior is undefined when one of the indexes returned by this
// function is a pentagon or is in the pentagon distortion area.
func (c Cell) gridDiskUnsafe(k int) ([]Cell, error) {
	cells, _, err := c.gridDiskDistancesUnsafe(k)
	return cells, err
}

// GridDisk produces cells within k distance of the origin cell.
//
// k-ring 0 is defined as the origin cell, k-ring 1 is defined as k-ring 0 and
// all neighboring cells, and so on.
func (c Cell) GridDisk(k int) ([]Cell, error) {
	cells, _, err := c.GridDiskDistances(k)
	return cells, err
}

// GridDiskDistances produces cells and their distances from the given origin
// cell, up to distance k.
//
// k-ring 0 is defined as the origin cell, k-ring 1 is defined as k-ring 0 and
// all neighboring cells, and so on.
func (c Cell) GridDiskDistances(k int) ([]Cell, []int, error) {
	// Try the faster, unsafe version first
	cells, distances, err := c.gridDiskDistancesUnsafe(k)
	if err == nil {
		return cells, distances, nil
	}

	// If the unsafe version failed, fall back to the safe version
	maxIdx, err := maxGridDiskSize(k)
	if err != nil {
		return nil, nil, err
	}
	out := make([]Cell, maxIdx)
	distances = make([]int, maxIdx)
	err = c.gridDiskDistancesInternal(k, &out, &distances, maxIdx, 0)
	return out, distances, err
}

func (c Cell) gridDiskDistancesInternal(k int, out *[]Cell, distances *[]int, maxIdx int, curK int) error {
	// Put origin in the output array, which is used as a hash set
	off := uint64(c) % uint64(maxIdx)
	for (*out)[off] != 0 && (*out)[off] != c {
		off = (off + 1) % uint64(maxIdx)
	}

	// We either got a free slot in the hash set or hit a duplicate
	// We might need to process the duplicate anyways because we got
	// here on a longer path before.
	if (*out)[off] == c && (*distances)[off] <= curK {
		return nil
	}

	(*out)[off] = c
	(*distances)[off] = curK

	// Base case: reached an index k away from the origin.
	if curK >= k {
		return nil
	}

	// Recurse to all neighbors in no particular order
	for i := 0; i < 6; i++ {
		rotations := 0
		nextNeighbor, rotations, neighborResult := c.neighborRotations(DIRECTIONS[i], rotations)
		if !errors.Is(neighborResult, ErrPentagonEncountered) {
			// ErrPentagonEncountered is an expected case when trying to traverse off of
			// pentagons.
			if neighborResult != nil {
				return neighborResult
			}

			neighborResult = nextNeighbor.gridDiskDistancesInternal(
				k, out, distances, maxIdx, curK+1)

			if neighborResult != nil {
				return neighborResult
			}
		}
	}

	return nil
}

func getNumCellsAtResolution(res int) int {
	if res < 0 || res > MAX_H3_RES {
		return 0
	}

	return 2 + 120*ipow(7, res)
}
