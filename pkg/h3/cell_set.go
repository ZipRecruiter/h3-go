package h3

import "fmt"

// CellSet represents a set of H3 cells.
type CellSet map[Cell]struct{}

// NewCellSetFromStrings creates a new cell set from a list of hex-encoded strings.
func NewCellSetFromStrings(ss []string) (CellSet, error) {
	cs := make(CellSet, len(ss))
	for _, s := range ss {
		c, err := NewCellFromString(s)
		if err != nil {
			return nil, fmt.Errorf("error converting string %s to cell: %w", s, err)
		}
		cs[c] = struct{}{}
	}
	return cs, nil
}

// NewCellSetFromCells creates a new cell set from a list of cells.
func NewCellSetFromCells(cells []Cell) CellSet {
	cs := make(CellSet, len(cells))
	for _, c := range cells {
		cs[c] = struct{}{}
	}
	return cs
}

// Cells returns the cells in the set as a list.
func (cs CellSet) Cells() []Cell {
	cells := make([]Cell, 0, len(cs))
	for c := range cs {
		cells = append(cells, c)
	}
	return cells
}

func (cs CellSet) String() string {
	return fmt.Sprintf("%v", cs.Cells())
}

// Strings returns the cells in the set as a list of hex-encoded strings.
func (cs CellSet) Strings() []string {
	ss := make([]string, 0, len(cs))
	for c := range cs {
		ss = append(ss, c.String())
	}
	return ss
}

// Contains returns whether the set contains the given cell.
func (cs CellSet) Contains(c Cell) bool {
	_, ok := cs[c]
	return ok
}

// Add adds a cell to the set.
func (cs CellSet) Add(c Cell) {
	cs[c] = struct{}{}
}

// Union returns the union of two cell sets.
func (cs CellSet) Union(other CellSet) CellSet {
	newSet := make(CellSet, len(cs)+len(other))
	for c := range cs {
		newSet[c] = struct{}{}
	}
	for c := range other {
		newSet[c] = struct{}{}
	}
	return newSet
}

// GridDisk returns the cells in the set's grid disk of radius k. The grid disk
// is the set of cells within k grid steps of the cells in the set. k=0 returns
// the set itself.
func (cs CellSet) GridDisk(k int) (CellSet, error) {
	if k == 0 {
		return cs, nil
	}

	newSet := make(CellSet)
	for cell := range cs {
		gridDiskListForCell, err := cell.GridDisk(k)
		if err != nil {
			return nil, fmt.Errorf("error computing grid disk for cell %s: %w", cell, err)
		}

		gridDiskSetForCell := NewCellSetFromCells(gridDiskListForCell)

		newSet = newSet.Union(gridDiskSetForCell)
	}

	return newSet, nil
}

// GridDistance returns the minimum grid distance between cells in the two sets.
// All cells in the sets must have the same resolution. Distance is zero if any
// of the cells overlap. The function will return an error if either set is
// empty.
func (cs CellSet) GridDistance(other CellSet) (int, error) {
	// If either set is empty, return an error
	if len(cs) == 0 || len(other) == 0 {
		return 0, fmt.Errorf("cannot compute grid distance between empty cell sets")
	}

	// Both sets must contain cells with the same resolution.
	thisResolution, err := cs.Resolution()
	if err != nil {
		return 0, fmt.Errorf("this cell set is not consistent resolution: %w", err)
	}

	otherResolution, err := other.Resolution()
	if err != nil {
		return 0, fmt.Errorf("other cell set is not consistent resolution: %w", err)
	}

	if thisResolution != otherResolution {
		return 0, fmt.Errorf("cell sets have different resolutions: %d and %d", thisResolution, otherResolution)
	}

	// If any cells overlap, the distance is zero.
	if cs.Intersects(other) {
		return 0, nil
	}

	// Check all pairs of cells to find the minimum distance.
	minDistance := -1

	for c1 := range cs {
		for c2 := range other {
			d, err := c1.GridDistance(c2)
			if err != nil {
				return 0, fmt.Errorf("error computing grid distance between cells %s and %s: %w", c1, c2, err)
			}

			if minDistance == -1 || d < minDistance {
				minDistance = d
			}
		}
	}

	return minDistance, nil
}

// Resolution returns the resolution of the cells in the set. The function will
// return an error if the set is empty or contains cells of different
// resolutions.
func (cs CellSet) Resolution() (int, error) {
	// If the set is empty, return an error
	if len(cs) == 0 {
		return 0, fmt.Errorf("empty cell set")
	}

	// Check if all cells have the same resolution
	resolution := -1
	for c := range cs {
		if resolution == -1 {
			resolution = c.Resolution()
		} else if c.Resolution() != resolution {
			return 0, fmt.Errorf("cell set contains cells of different resolutions")
		}
	}

	return resolution, nil
}

// Intersects returns whether the set intersects with another set.
func (cs CellSet) Intersects(other CellSet) bool {
	for c := range cs {
		if other.Contains(c) {
			return true
		}
	}
	return false
}
