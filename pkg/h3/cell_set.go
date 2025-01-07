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
		gridDiskListForCell, err := cell.gridDisk(k)
		if err != nil {
			return nil, fmt.Errorf("error computing grid disk for cell %s: %w", cell, err)
		}

		gridDiskSetForCell := NewCellSetFromCells(gridDiskListForCell)

		newSet = newSet.Union(gridDiskSetForCell)
	}

	return newSet, nil
}
