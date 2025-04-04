package h3

import (
	"fmt"
)

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
	// k<0 returns an error
	if k < 0 {
		return nil, fmt.Errorf("k must be >= 0")
	}

	// k=0 returns the set itself
	if k == 0 {
		return cs, nil
	}

	// If the set is empty, return an error
	if len(cs) == 0 {
		return nil, fmt.Errorf("empty cell set")
	}

	// Start with the original set
	result := make(CellSet, len(cs))
	for c := range cs {
		result.Add(c)
	}

	// For each k, expand only the cells added in the previous step
	currentShell := cs
	for i := 0; i < k; i++ {
		// Find boundary cells of the current shell
		nextShell := make(CellSet, len(currentShell))

		for c := range currentShell {
			// Get immediate neighbors of the current shell
			neighbors, err := c.GridDisk(1)
			if err != nil {
				return nil, fmt.Errorf("error getting neighbors for cell %s: %w", c, err)
			}

			// Add new neighbors to the next shell
			for _, n := range neighbors {
				if !result.Contains(n) {
					nextShell.Add(n)
					result.Add(n)
				}
			}
		}

		if len(nextShell) == 0 {
			break // No new cells added, exit the loop
		}

		// Move to the next shell
		currentShell = nextShell
	}

	return result, nil
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

	// Get boundary cells of both sets
	selfBoundaryCells, err := cs.BoundaryCells()
	if err != nil {
		return 0, fmt.Errorf("error getting boundary cells for this set: %w", err)
	}

	otherBoundaryCells, err := other.BoundaryCells()
	if err != nil {
		return 0, fmt.Errorf("error getting boundary cells for other set: %w", err)
	}

	// Group cells by their lower-resolution parent to reduce search space. We'll use 2 levels of parent to create the groups.
	const parentReduction = 2
	selfGroups := make(map[Cell][]Cell)
	otherGroups := make(map[Cell][]Cell)

	// Group cells by their parent
	for c := range selfBoundaryCells {
		parent, err := c.Parent(thisResolution - parentReduction)
		if err != nil {
			return 0, fmt.Errorf("error getting parent for cell %s: %w", c, err)
		}
		selfGroups[parent] = append(selfGroups[parent], c)
	}

	for c := range otherBoundaryCells {
		parent, err := c.Parent(thisResolution - parentReduction)
		if err != nil {
			return 0, fmt.Errorf("error getting parent for cell %s: %w", c, err)
		}
		otherGroups[parent] = append(otherGroups[parent], c)
	}

	// First find minimum distance between parent cells
	minParentDistance := -1
	relevantPairs := make([][2]Cell, 0, len(selfGroups)*len(otherGroups))
	for parent1 := range selfGroups {
		for parent2 := range otherGroups {
			d, err := parent1.GridDistance(parent2)
			if err != nil {
				return 0, fmt.Errorf("error computing grid distance between parent cells %s and %s: %w", parent1, parent2, err)
			}

			if minParentDistance == -1 || d < minParentDistance {
				minParentDistance = d
				relevantPairs = relevantPairs[:0]
			}

			// This pair is among the closest, so we'll check its children later
			if d <= minParentDistance+parentReduction {
				relevantPairs = append(relevantPairs, [2]Cell{parent1, parent2})
			}
		}
	}

	// Now calculate the actual minimum distance between the children of the relevant pairs
	minDistance := -1

	for _, pair := range relevantPairs {
		parent1, parent2 := pair[0], pair[1]
		childrenPairs1 := selfGroups[parent1]
		childrenPairs2 := otherGroups[parent2]

		for _, c1 := range childrenPairs1 {
			for _, c2 := range childrenPairs2 {
				d, err := c1.GridDistance(c2)
				if err != nil {
					return 0, fmt.Errorf("error computing grid distance between cells %s and %s: %w", c1, c2, err)
				}

				if minDistance == -1 || d < minDistance {
					minDistance = d
				}
			}
		}
	}

	// If no distance was found, return an error
	if minDistance == -1 {
		return 0, fmt.Errorf("no distance found between cell sets")
	}

	return minDistance, nil
}

// BoundaryCells returns the cells on the outer boundary of the set. A boundary
// cell is one that has at least one neighboring cell that's not in the set.
func (cs CellSet) BoundaryCells() (CellSet, error) {
	// If the set has < 7 cells, return the set itself because there aren't enough
	// cells to enclose one cell.
	if len(cs) < 7 {
		return cs, nil
	}

	boundaryCells := make(CellSet, len(cs))
	for c := range cs {
		neighbors, err := c.GridDisk(1)
		if err != nil {
			return nil, fmt.Errorf("error getting neighbors for cell %s: %w", c, err)
		}

		for _, n := range neighbors {
			if !cs.Contains(n) {
				boundaryCells.Add(c)
				break
			}
		}
	}

	return boundaryCells, nil
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
