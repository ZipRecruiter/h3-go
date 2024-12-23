package h3

// CellSet represents a set of H3 cells.
type CellSet map[Cell]struct{}

// NewCellSetFromStrings creates a new cell set from a list of hex-encoded strings.
func NewCellSetFromStrings(ss []string) (CellSet, error) {
	cs := make(CellSet, len(ss))
	for _, s := range ss {
		c, err := NewCellFromString(s)
		if err != nil {
			return nil, err
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
