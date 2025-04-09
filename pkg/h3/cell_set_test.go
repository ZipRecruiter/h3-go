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
	// L7 cells for San Francisco
	sfCells := NewCellSetFromCells([]Cell{0x872830876ffffff, 0x87283082bffffff, 0x87283082affffff, 0x872830874ffffff, 0x872830829ffffff, 0x872830828ffffff, 0x87283082effffff, 0x87283095bffffff, 0x87283095affffff, 0x87283082dffffff, 0x87283082cffffff, 0x872830821ffffff, 0x872830820ffffff, 0x872830958ffffff, 0x87283095effffff, 0x872830953ffffff, 0x872830952ffffff, 0x872830825ffffff, 0x87283095cffffff, 0x872830951ffffff, 0x872830950ffffff, 0x872830942ffffff})

	// L7 cells for Vallejo
	vallejoCells := NewCellSetFromCells([]Cell{0x87283002dffffff, 0x87283002cffffff, 0x87283002effffff, 0x872830023ffffff, 0x872830004ffffff, 0x872830153ffffff, 0x872830152ffffff, 0x872830021ffffff, 0x872830020ffffff, 0x872830022ffffff, 0x872830026ffffff, 0x872830024ffffff, 0x872830025ffffff, 0x872830156ffffff, 0x872830150ffffff, 0x872830154ffffff, 0x872830109ffffff, 0x87283010bffffff, 0x87283010affffff, 0x872830119ffffff, 0x87283011dffffff, 0x87283010effffff, 0x872830108ffffff})

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
		{
			"sf to vallejo",
			sfCells,
			args{vallejoCells},
			16,
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

func BenchmarkCellSet_GridDisk(b *testing.B) {
	// L7 cells for San Francisco
	cells := NewCellSetFromCells([]Cell{
		0x872830876ffffff,
		0x87283082bffffff,
		0x87283082affffff,
		0x872830874ffffff,
		0x872830829ffffff,
		0x872830828ffffff,
		0x87283082effffff,
		0x87283095bffffff,
		0x87283095affffff,
		0x87283082dffffff,
		0x87283082cffffff,
		0x872830821ffffff,
		0x872830820ffffff,
		0x872830958ffffff,
		0x87283095effffff,
		0x872830953ffffff,
		0x872830952ffffff,
		0x872830825ffffff,
		0x87283095cffffff,
		0x872830951ffffff,
		0x872830950ffffff,
		0x872830942ffffff,
	})

	for k := 0; k < 100; k++ {
		b.Run(fmt.Sprintf("k=%d", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := cells.GridDisk(k)
				if err != nil {
					b.Fatalf("GridDisk() error = %v", err)
				}
			}
		})
	}
}

func TestCellSet_BoundaryCells(t *testing.T) {
	tests := []struct {
		name    string
		cs      CellSet
		want    CellSet
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"empty",
			CellSet{},
			CellSet{},
			assert.NoError,
		},
		{
			"single cell",
			CellSet{0x87283082affffff: {}},
			CellSet{0x87283082affffff: {}},
			assert.NoError,
		},
		{
			"two cells",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}},
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}},
			assert.NoError,
		},
		{
			"6 cells around a center cell should return the outer cells",
			CellSet{
				0x872830876ffffff: {},
				0x87283082bffffff: {},
				0x872830874ffffff: {},
				0x872830829ffffff: {},
				0x872830828ffffff: {},
				0x87283082dffffff: {},
				0x87283095affffff: {},
			},
			CellSet{
				0x872830876ffffff: {},
				0x87283082bffffff: {},
				0x872830828ffffff: {},
				0x87283082dffffff: {},
				0x87283095affffff: {},
				0x872830874ffffff: {},
			},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.BoundaryCells()
			if !tt.wantErr(t, err, fmt.Sprintf("BoundaryCells()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "BoundaryCells()")
		})
	}
}

func BenchmarkCellSet_GridDistance(b *testing.B) {
	// L7 cells for San Francisco
	sfCells := NewCellSetFromCells([]Cell{0x872830876ffffff, 0x87283082bffffff, 0x87283082affffff, 0x872830874ffffff, 0x872830829ffffff, 0x872830828ffffff, 0x87283082effffff, 0x87283095bffffff, 0x87283095affffff, 0x87283082dffffff, 0x87283082cffffff, 0x872830821ffffff, 0x872830820ffffff, 0x872830958ffffff, 0x87283095effffff, 0x872830953ffffff, 0x872830952ffffff, 0x872830825ffffff, 0x87283095cffffff, 0x872830951ffffff, 0x872830950ffffff, 0x872830942ffffff})

	// L7 cells for Vallejo
	vallejoCells := NewCellSetFromCells([]Cell{0x87283002dffffff, 0x87283002cffffff, 0x87283002effffff, 0x872830023ffffff, 0x872830004ffffff, 0x872830153ffffff, 0x872830152ffffff, 0x872830021ffffff, 0x872830020ffffff, 0x872830022ffffff, 0x872830026ffffff, 0x872830024ffffff, 0x872830025ffffff, 0x872830156ffffff, 0x872830150ffffff, 0x872830154ffffff, 0x872830109ffffff, 0x87283010bffffff, 0x87283010affffff, 0x872830119ffffff, 0x87283011dffffff, 0x87283010effffff, 0x872830108ffffff})

	// L7 cells for Chicago
	chicagoCells := NewCellSetFromCells([]Cell{0x872664523ffffff, 0x872664c84ffffff, 0x87266452dffffff, 0x872664c89ffffff, 0x872664193ffffff, 0x872664ca2ffffff, 0x872664cacffffff, 0x872664cb1ffffff, 0x872664564ffffff, 0x872664cc0ffffff, 0x872664cc5ffffff, 0x872664ccaffffff, 0x872664cd4ffffff, 0x872664cd9ffffff, 0x872664cdeffffff, 0x872664cf2ffffff, 0x872664d88ffffff, 0x87275934cffffff, 0x872664c1effffff, 0x872664521ffffff, 0x872664c8cffffff, 0x872664191ffffff, 0x872664196ffffff, 0x872664ca0ffffff, 0x872664ca5ffffff, 0x872664caaffffff, 0x872664562ffffff, 0x87266456cffffff, 0x872664cc3ffffff, 0x872664cc8ffffff, 0x872664ccdffffff, 0x872664cd2ffffff, 0x872664cdcffffff, 0x872664cebffffff, 0x872664cf0ffffff, 0x872664cf5ffffff, 0x872664d8bffffff, 0x872664d9affffff, 0x872664c12ffffff, 0x872664c80ffffff, 0x872664c85ffffff, 0x872664194ffffff, 0x872664ca3ffffff, 0x872664ca8ffffff, 0x872664cadffffff, 0x872664cb2ffffff, 0x8726641b2ffffff, 0x872664560ffffff, 0x872664565ffffff, 0x872664cc1ffffff, 0x872664cc6ffffff, 0x872664ccbffffff, 0x872664cd0ffffff, 0x872664cd5ffffff, 0x872664cf3ffffff, 0x872664d89ffffff, 0x872664d8effffff, 0x872664d98ffffff, 0x872664d9dffffff, 0x872759343ffffff, 0x87275934dffffff, 0x87275936bffffff, 0x872664c10ffffff, 0x872664c1affffff, 0x872664c83ffffff, 0x87266452cffffff, 0x872664c88ffffff, 0x872664c8dffffff, 0x872664192ffffff, 0x872664ca1ffffff, 0x872664ca6ffffff, 0x872664cabffffff, 0x872664cb0ffffff, 0x872664cb5ffffff, 0x872664563ffffff, 0x872664cc4ffffff, 0x872664cc9ffffff, 0x872664cceffffff, 0x872664cd3ffffff, 0x872664cd8ffffff, 0x872664cddffffff, 0x872664ce2ffffff, 0x872664cf1ffffff, 0x872664cf6ffffff, 0x872664d8cffffff, 0x872664d9bffffff, 0x872759341ffffff, 0x872759369ffffff, 0x872664c13ffffff, 0x872664c18ffffff, 0x872664520ffffff, 0x872664525ffffff, 0x872664c81ffffff, 0x872664c86ffffff, 0x872664c8bffffff, 0x872664190ffffff, 0x872664ca4ffffff, 0x872664ca9ffffff, 0x872664caeffffff, 0x872664561ffffff, 0x872664566ffffff, 0x872664cc2ffffff, 0x872664570ffffff, 0x872664575ffffff, 0x872664cccffffff, 0x872664cd1ffffff, 0x872664cd6ffffff, 0x872664cdbffffff, 0x872664ceaffffff, 0x872664cf4ffffff, 0x872664d8affffff, 0x872664d99ffffff, 0x872664d9effffff, 0x87275934effffff, 0x87275935dffffff, 0x872664c16ffffff, 0x872664c1bffffff})

	// L7 cells for Milwaukee
	mkeCells := NewCellSetFromCells([]Cell{0x87275d745ffffff, 0x87275d39affffff, 0x87275d0dcffffff, 0x87275d764ffffff, 0x87275d396ffffff, 0x87275d286ffffff, 0x87275d0d8ffffff, 0x87275d76dffffff, 0x87275d2b2ffffff, 0x87275d0cbffffff, 0x87275d760ffffff, 0x87275d392ffffff, 0x87275d282ffffff, 0x87275d769ffffff, 0x87275d39bffffff, 0x87275d0ddffffff, 0x87275d765ffffff, 0x87275d294ffffff, 0x87275d0c3ffffff, 0x87275d0d9ffffff, 0x87275d76effffff, 0x87275d2b3ffffff, 0x87275d761ffffff, 0x87275d2a6ffffff, 0x87275d393ffffff, 0x87275d283ffffff, 0x87275d39cffffff, 0x87275d2a2ffffff, 0x87275d766ffffff, 0x87275d398ffffff, 0x87275d2b4ffffff, 0x87275d74cffffff, 0x87275d762ffffff, 0x87275d284ffffff, 0x87275d76bffffff, 0x87275d2b0ffffff, 0x87275d390ffffff, 0x87275d280ffffff, 0x87275d296ffffff, 0x87275d0dbffffff, 0x87275d2b5ffffff, 0x87275d763ffffff, 0x87275d76cffffff, 0x87275d2b1ffffff, 0x87275d39effffff, 0x87275d0caffffff, 0x87275d775ffffff, 0x87275d391ffffff, 0x87275d281ffffff, 0x87275d768ffffff, 0x87275d0d1ffffff})

	// L7 cells for Evanston (adjacent to Chicago)
	evanstonCells := NewCellSetFromCells([]Cell{0x872664d81ffffff, 0x872664d85ffffff, 0x872664d80ffffff, 0x872664d84ffffff})

	for i := 0; i < b.N; i++ {
		// Distance between smaller, closer sets
		_, err := sfCells.GridDistance(vallejoCells)
		if err != nil {
			b.Fatalf("GridDistance() error = %v", err)
		}

		// Distance between larger, farther sets
		_, err = chicagoCells.GridDistance(mkeCells)
		if err != nil {
			b.Fatalf("GridDistance() error = %v", err)
		}

		// Distance between two adjacent (but not overlapping) sets
		_, err = chicagoCells.GridDistance(evanstonCells)
		if err != nil {
			b.Fatalf("GridDistance() error = %v", err)
		}
	}
}

func TestCellSet_Subtract(t *testing.T) {
	type args struct {
		other CellSet
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
			args{CellSet{}},
			CellSet{},
			assert.NoError,
		},
		{
			"different resolutions",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x86283080fffffff: {}}},
			nil,
			assert.Error,
		},
		{
			"single cell removed",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082affffff: {}}},
			CellSet{},
			assert.NoError,
		},
		{
			"single cell not removed",
			CellSet{0x87283082affffff: {}},
			args{CellSet{0x87283082bffffff: {}}},
			CellSet{0x87283082affffff: {}},
			assert.NoError,
		},
		{
			"multiple cells, some removed",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			args{CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}}},
			CellSet{0x87283082cffffff: {}},
			assert.NoError,
		},
		{
			"multiple cells, all removed",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			args{CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}}},
			CellSet{},
			assert.NoError,
		},
		{
			"multiple cells, none removed",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			args{CellSet{0x87283082dffffff: {}}},
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			assert.NoError,
		},
		{
			"multiple cells, empty removal set",
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			args{CellSet{}},
			CellSet{0x87283082affffff: {}, 0x87283082bffffff: {}, 0x87283082cffffff: {}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.Subtract(tt.args.other)
			if !tt.wantErr(t, err, fmt.Sprintf("Subtract(%v)", tt.args.other)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Subtract(%v)", tt.args.other)
		})
	}
}

func TestCellSet_Parent(t *testing.T) {
	type args struct {
		resolution int
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
			nil,
			assert.Error,
		},
		{
			"single cell, same resolution",
			CellSet{0x87283082affffff: {}},
			args{7},
			CellSet{0x87283082affffff: {}},
			assert.NoError,
		},
		{
			"single cell, different resolution",
			CellSet{0x87283082affffff: {}},
			args{5},
			CellSet{0x85283083fffffff: {}},
			assert.NoError,
		},
		{
			"single cell, greater resolution",
			CellSet{0x87283082affffff: {}},
			args{8},
			nil,
			assert.Error,
		},
		{
			"single cell, invalid resolution",
			CellSet{0x87283082affffff: {}},
			args{-1},
			nil,
			assert.Error,
		},
		{
			"multiple cells, different resolution",
			CellSet{0x872830876ffffff: {}, 0x87283082bffffff: {}, 0x87283082affffff: {}},
			args{5},
			CellSet{0x85283083fffffff: {}, 0x85283087fffffff: {}},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cs.Parent(tt.args.resolution)
			if !tt.wantErr(t, err, fmt.Sprintf("Parent(%v)", tt.args.resolution)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Parent(%v)", tt.args.resolution)
		})
	}
}
