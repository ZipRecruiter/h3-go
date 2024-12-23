package h3

import (
	"testing"
)

func TestNewCellSetFromStrings(t *testing.T) {
	strs := []string{"8f283473fffffff", "872830829fffffff"}
	cs, err := NewCellSetFromStrings(strs)
	if err != nil {
		t.Fatalf("NewCellSetFromStrings() error = %v", err)
	}
	if len(cs) != len(strs) {
		t.Errorf("NewCellSetFromStrings() = %v, want %v", len(cs), len(strs))
	}
	for _, s := range strs {
		c, _ := NewCellFromString(s)
		if !cs.Contains(c) {
			t.Errorf("NewCellSetFromStrings() missing cell %v", c)
		}
	}
}

func TestNewCellSetFromCells(t *testing.T) {
	cells := []Cell{0x8f283473fffffff, 0x872830829fffffff}
	cs := NewCellSetFromCells(cells)
	if len(cs) != len(cells) {
		t.Errorf("NewCellSetFromCells() = %v, want %v", len(cs), len(cells))
	}
	for _, c := range cells {
		if !cs.Contains(c) {
			t.Errorf("NewCellSetFromCells() missing cell %v", c)
		}
	}
}

func TestCellSet_Contains(t *testing.T) {
	cs := CellSet{0x8f283473fffffff: {}, 0x872830829fffffff: {}}
	tests := []struct {
		cell Cell
		want bool
	}{
		{0x8f283473fffffff, true},
		{0x872830829fffffff, true},
		{0x8f2834740000000, false},
	}
	for _, tt := range tests {
		if got := cs.Contains(tt.cell); got != tt.want {
			t.Errorf("Contains() = %v, want %v", got, tt.want)
		}
	}
}

func TestCellSet_Add(t *testing.T) {
	cs := CellSet{}
	cell := Cell(0x8f283473fffffff)
	cs.Add(cell)
	if !cs.Contains(cell) {
		t.Errorf("Add() did not add cell %v", cell)
	}
}

func TestCellSet_Cells(t *testing.T) {
	cs := CellSet{0x8f283473fffffff: {}, 0x872830829fffffff: {}}
	cells := cs.Cells()
	if len(cells) != len(cs) {
		t.Errorf("Cells() = %v, want %v", len(cells), len(cs))
	}
	for _, c := range cells {
		if !cs.Contains(c) {
			t.Errorf("Cells() missing cell %v", c)
		}
	}
}

func TestCellSet_Strings(t *testing.T) {
	cs := CellSet{0x8f283473fffffff: {}, 0x872830829fffffff: {}}
	strs := cs.Strings()
	if len(strs) != len(cs) {
		t.Errorf("Strings() = %v, want %v", len(strs), len(cs))
	}
	for _, s := range strs {
		c, _ := NewCellFromString(s)
		if !cs.Contains(c) {
			t.Errorf("Strings() missing cell %v", c)
		}
	}
}
