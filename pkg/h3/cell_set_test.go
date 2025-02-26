package h3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestCellSet_Union(t *testing.T) {
	type args struct {
		other CellSet
	}
	tests := []struct {
		name string
		cs   CellSet
		args args
		want CellSet
	}{
		{
			"empty",
			CellSet{},
			args{CellSet{}},
			CellSet{},
		},
		{
			"non-empty",
			CellSet{0x8f283473fffffff: {}},
			args{CellSet{0x872830829fffffff: {}}},
			CellSet{0x8f283473fffffff: {}, 0x872830829fffffff: {}},
		},
		{
			"non-empty with overlap",
			CellSet{0x8f283473fffffff: {}},
			args{CellSet{0x8f283473fffffff: {}}},
			CellSet{0x8f283473fffffff: {}},
		},
		{
			"non-empty with empty",
			CellSet{0x8f283473fffffff: {}},
			args{CellSet{}},
			CellSet{0x8f283473fffffff: {}},
		},
		{
			"empty with non-empty",
			CellSet{},
			args{CellSet{0x8f283473fffffff: {}}},
			CellSet{0x8f283473fffffff: {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.cs.Union(tt.args.other), "Union(%v)", tt.args.other)
		})
	}
}

func TestCellSet_GridDisk(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		cs      CellSet
		args    args
		want    CellSet
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty",
			CellSet{},
			args{0},
			CellSet{},
			assert.NoError,
		},
		{
			"non-empty",
			CellSet{0x87283082affffff: {}},
			args{0},
			CellSet{0x87283082affffff: {}},
			assert.NoError,
		},
		{
			"non-empty k=1",
			CellSet{0x87283082affffff: {}},
			args{1},
			CellSet{
				0x87283082affffff: {},
				0x87283082bffffff: {},
				0x87283080cffffff: {},
				0x872830801ffffff: {},
				0x872830805ffffff: {},
				0x87283082effffff: {},
				0x872830828ffffff: {},
			},
			assert.NoError,
		},
		{
			"two nearby cells, k=1",
			CellSet{
				0x87283082affffff: {},
				0x872830823ffffff: {},
			},
			args{1},
			CellSet{
				0x872830801ffffff: {},
				0x872830804ffffff: {},
				0x872830805ffffff: {},
				0x87283080cffffff: {},
				0x872830820ffffff: {},
				0x872830821ffffff: {},
				0x872830822ffffff: {},
				0x872830823ffffff: {},
				0x872830828ffffff: {},
				0x87283082affffff: {},
				0x87283082bffffff: {},
				0x87283082effffff: {},
			},
			assert.NoError,
		},
		{
			"invalid k",
			CellSet{0x87283082affffff: {}},
			args{-1},
			nil,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.GridDisk(tt.args.k)
			if !tt.wantErr(t, err, fmt.Sprintf("GridDisk(%v)", tt.args.k)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GridDisk(%v)", tt.args.k)
		})
	}
}

func TestCellSet_Intersects(t *testing.T) {
	type args struct {
		other CellSet
	}
	tests := []struct {
		name string
		cs   CellSet
		args args
		want bool
	}{
		{
			"empty",
			CellSet{},
			args{CellSet{}},
			false,
		},
		{
			"non-empty intersects",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082affffff: {}}},
			true,
		},
		{
			"non-empty no intersection",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082bffffff: {}}},
			false,
		},
		{
			"empty other",
			CellSet{0x87283082affffff: {}},
			args{CellSet{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.cs.Intersects(tt.args.other), "Intersects(%v)", tt.args.other)
		})
	}
}

func TestCellSet_Resolution(t *testing.T) {
	tests := []struct {
		name    string
		cs      CellSet
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty",
			CellSet{},
			0,
			assert.Error,
		},
		{
			"single cell",
			CellSet{0x87283082affffff: {}},
			7,
			assert.NoError,
		},
		{
			"multiple cells",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}},
			7,
			assert.NoError,
		},
		{
			"multiple cells different resolutions",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x86283080fffffff: {}}, // 7, 7, 6
			0,
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.Resolution()
			if !tt.wantErr(t, err, fmt.Sprintf("Resolution()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "Resolution()")
		})
	}
}

func TestCellSet_GridDistance(t *testing.T) {
	type args struct {
		other CellSet
	}
	tests := []struct {
		name    string
		cs      CellSet
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty",
			CellSet{},
			args{CellSet{}},
			0,
			assert.Error,
		},
		{
			"single identical cell",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082affffff: {}}},
			0,
			assert.NoError,
		},
		{
			"different resolutions",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x86283080fffffff: {}}},
			0,
			assert.Error,
		},
		{
			"single cell distance 1",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082bffffff: {}}},
			1,
			assert.NoError,
		},
		{
			"overlapping cell sets",
			CellSet{0x872830876ffffff: {}, 0x872830874ffffff: {}},
			args{CellSet{0x872830876ffffff: {}, 0x872830808ffffff: {}}},
			0,
			assert.NoError,
		},
		{
			"non-zero distance",
			CellSet{0x872830876ffffff: {}, 0x872830874ffffff: {}},
			args{CellSet{0x872830808ffffff: {}}},
			2,
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.GridDistance(tt.args.other)
			if !tt.wantErr(t, err, fmt.Sprintf("GridDistance(%v)", tt.args.other)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GridDistance(%v)", tt.args.other)
		})
	}
}
