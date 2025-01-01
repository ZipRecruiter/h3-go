package h3

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustCellFromString(s string) Cell {
	c, err := NewCellFromString(s)
	if err != nil {
		panic(err)
	}
	return c
}

func TestCell_Valid(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want bool
	}{
		{
			name: "from isValidCell.txt - invalid",
			c:    mustCellFromString("85283473ffff"),
			want: false,
		},
		{
			name: "from isValidCell.txt - valid",
			c:    mustCellFromString("85283473fffffff"),
			want: true,
		},
		{
			name: "valid cell from the bay area",
			c:    mustCellFromString("872830829ffffff"),
			want: true,
		},
		{
			name: "invalid - high bit set",
			c:    mustCellFromString("f5283473fffffff"),
			want: false,
		},
		{
			name: "invalid - wrong mode",
			c:    mustCellFromString("85283473ffff"),
			want: false,
		},
		{
			name: "invalid - reserved bits set",
			c:    mustCellFromString("85283473ffefff"),
			want: false,
		},
		{
			name: "invalid - base cell out of range",
			c:    mustCellFromString("8f283473fffffff"),
			want: false,
		},
		{
			name: "invalid - resolution out of range",
			c:    mustCellFromString("872830829fffffff"),
			want: false,
		},
		{
			name: "invalid - invalid digit",
			c:    mustCellFromString("872830829fffffff"),
			want: false,
		},
		{
			name: "invalid - invalid digit in higher resolution",
			c:    mustCellFromString("872830829fffffff"),
			want: false,
		},
		{
			name: "invalid - deleted subsequence",
			c:    newCell(1, 4, K_AXES_DIGIT),
			want: false,
		},
		{
			name: "invalid - bad digit",
			c:    H3_INIT.setMode(H3_CELL_MODE).setResolution(1),
			want: false,
		},
		{
			name: "invalid - high bit",
			c:    H3_INIT.setMode(H3_CELL_MODE).setHighBit(1),
			want: false,
		},
		{
			name: "valid - reserved bit 0",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(0),
			want: true,
		},
		{
			name: "invalid - reserved bit 1",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(1),
			want: false,
		},
		{
			name: "invalid - reserved bit 2",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(2),
			want: false,
		},
		{
			name: "invalid - reserved bit 3",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(3),
			want: false,
		},
		{
			name: "invalid - reserved bit 4",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(4),
			want: false,
		},
		{
			name: "invalid - reserved bit 5",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(5),
			want: false,
		},
		{
			name: "invalid - reserved bit 6",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(6),
			want: false,
		},
		{
			name: "invalid - reserved bit 7",
			c:    H3_INIT.setMode(H3_CELL_MODE).setReservedBits(7),
			want: false,
		},
		{
			name: "invalid - h3 mode 0",
			c:    H3_INIT.setMode(0),
			want: false,
		},
		{
			name: "valid - h3 cell mode",
			c:    H3_INIT.setMode(H3_CELL_MODE),
			want: true,
		},
		{
			name: "invalid - mode 2",
			c:    H3_INIT.setMode(2),
			want: false,
		},
		{
			name: "invalid - mode 3",
			c:    H3_INIT.setMode(3),
			want: false,
		},
		{
			name: "invalid - mode 4",
			c:    H3_INIT.setMode(4),
			want: false,
		},
		{
			name: "invalid - mode 5",
			c:    H3_INIT.setMode(5),
			want: false,
		},
		{
			name: "invalid - mode 6",
			c:    H3_INIT.setMode(6),
			want: false,
		},
		{
			name: "invalid - mode 7",
			c:    H3_INIT.setMode(7),
			want: false,
		},
		{
			name: "invalid - mode 8",
			c:    H3_INIT.setMode(8),
			want: false,
		},
		{
			name: "invalid - mode 9",
			c:    H3_INIT.setMode(9),
			want: false,
		},
		{
			name: "invalid - mode 10",
			c:    H3_INIT.setMode(10),
			want: false,
		},
		{
			name: "invalid - mode 11",
			c:    H3_INIT.setMode(11),
			want: false,
		},
		{
			name: "invalid - mode 12",
			c:    H3_INIT.setMode(12),
			want: false,
		},
		{
			name: "invalid - mode 13",
			c:    H3_INIT.setMode(13),
			want: false,
		},
		{
			name: "invalid - mode 14",
			c:    H3_INIT.setMode(14),
			want: false,
		},
		{
			name: "invalid - mode 15",
			c:    H3_INIT.setMode(15),
			want: false,
		},
		{
			name: "invalid - wrong base cell",
			c:    H3_INIT.setMode(H3_CELL_MODE).setBaseCell(NUM_BASE_CELLS),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCell_Valid_baseCells(t *testing.T) {
	for i := 0; i < NUM_BASE_CELLS; i++ {
		c := H3_INIT.setMode(H3_CELL_MODE).setBaseCell(baseCell(i))
		assert.True(t, c.Valid(), "expected base cell %d to be valid", i)
		assert.Equal(t, baseCell(i), c.BaseCell(), "expected base cell %d", i)
	}
}

func TestCell_Valid_digits(t *testing.T) {
	ll := NewLatLng(0, 0)
	c, err := NewCellFromLatLng(ll, 1)
	assert.NoError(t, err)
	// Set a bit for an unused digit to something else
	c ^= 1
	assert.False(t, c.Valid(), "expected invalid cell")
}

func TestCell_Valid_atResolution(t *testing.T) {
	for i := 0; i <= MAX_H3_RES; i++ {
		ll := NewLatLng(0, 0)
		c, err := NewCellFromLatLng(ll, i)
		assert.NoError(t, err)
		assert.True(t, c.Valid(), "expected resolution %d to be valid", i)
	}
}

func TestCell_neighborRotations(t *testing.T) {
	t.Run("identity", func(t *testing.T) {
		origin := Cell(0x811d7ffffffffff)
		rotations := 0
		out, rotations, err := origin.neighborRotations(CENTER_DIGIT, rotations)
		assert.NoError(t, err)
		assert.Equal(t, origin, out)
		assert.Equal(t, 0, rotations)
	})

	t.Run("rotations overflow", func(t *testing.T) {
		origin := newCell(0, 0, CENTER_DIGIT)
		// A multiple of 6, so effectively no rotation. Very close to INT32_MAX.
		rotations := 2147483646
		out, rotations, err := origin.neighborRotations(K_AXES_DIGIT, rotations)
		assert.NoError(t, err)
		expected := newCell(0, 1, CENTER_DIGIT)
		assert.Equal(t, expected, out, "expected neighbor")
		assert.Equal(t, 5, rotations, "expected rotations")
	})

	t.Run("invalid rotations", func(t *testing.T) {
		origin := Cell(0x811d7ffffffffff)
		rotations := 0

		out, rotations, err := origin.neighborRotations(-1, rotations)
		assert.Error(t, err, "invalid direction should fail")
		assert.Nil(t, out)
		out, rotations, err = origin.neighborRotations(7, rotations)
		assert.Error(t, err, "invalid direction should fail")
		assert.Nil(t, out)
		out, rotations, err = origin.neighborRotations(100, rotations)
		assert.Error(t, err, "invalid direction should fail")
		assert.Nil(t, out)
	})
}

func Test_isResolutionClassIII(t *testing.T) {
	coord := NewLatLng(0, 0)
	for i := 0; i <= MAX_H3_RES; i++ {
		c, err := NewCellFromLatLng(coord, i)
		assert.NoError(t, err)
		assert.Equal(t, c.isResolutionClassIII(), isResolutionClassIII(i))
	}
}

func Test_newCell(t *testing.T) {
	c := newCell(5, 12, 1)
	assert.Equal(t, 5, c.Resolution())
	assert.Equal(t, baseCell(12), c.BaseCell())
	assert.Equal(t, H3_CELL_MODE, c.Mode())
	for i := 1; i <= 5; i++ {
		assert.Equal(t, K_AXES_DIGIT, c.getIndexDigit(i), "expected digit %d set", i)
	}
	for i := 6; i <= MAX_H3_RES; i++ {
		assert.Equal(t, INVALID_DIGIT, c.getIndexDigit(i), "expected digit %d blanked", i)
	}
	assert.Equal(t, Cell(0x85184927fffffff), c)
}

func TestNewCellFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Cell
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "no index from nothing",
			args:    args{s: ""},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "no index from junk",
			args:    args{s: "**"},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "valid index",
			args:    args{s: "ffffffffffffffff"},
			want:    Cell(0xffffffffffffffff),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCellFromString(tt.args.s)
			if !tt.wantErr(t, err, fmt.Sprintf("NewCellFromString(%v)", tt.args.s)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewCellFromString(%v)", tt.args.s)
		})
	}
}

func TestCell_String(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want string
	}{
		{
			name: "cafe",
			c:    0xCAFE,
			want: "cafe",
		},
		{
			name: "large",
			c:    0xffffffffffffffff,
			want: "ffffffffffffffff",
		},
		{
			name: "standard",
			c:    0x85184927fffffff,
			want: "85184927fffffff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.String(), "String()")
		})
	}
}

func TestNewCellFromLatLng(t *testing.T) {
	type args struct {
		ll  LatLng
		res int
	}
	tests := []struct {
		name    string
		args    args
		want    Cell
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "normal",
			args:    args{ll: NewLatLng(20, 123), res: 2},
			want:    Cell(0x824b9ffffffffff),
			wantErr: assert.NoError,
		},
		{
			name:    "extreme longitude",
			args:    args{ll: NewLatLng(0, 1e45), res: 14},
			want:    Cell(0x8e7b2b95e164cd7),
			wantErr: assert.NoError,
		},
		{
			name:    "extreme latitude",
			args:    args{ll: NewLatLng(1e46, 1e45), res: 15},
			want:    Cell(0x8ff3a922bc1a65b),
			wantErr: assert.NoError,
		},
		{
			name:    "invalid low resolution",
			args:    args{ll: NewLatLng(0, 0), res: -1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "invalid high resolution",
			args:    args{ll: NewLatLng(0, 0), res: MAX_H3_RES + 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "infinite latitude",
			args:    args{ll: NewLatLng(math.Inf(-1), 0), res: 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "valid at 0,0 res 1",
			args:    args{ll: NewLatLng(0, 0), res: 1},
			want:    0x81757ffffffffff,
			wantErr: assert.NoError,
		},
		{
			name:    "valid at 0,0 res 2",
			args:    args{ll: NewLatLng(0, 0), res: 2},
			want:    0x82754ffffffffff,
			wantErr: assert.NoError,
		},
		{
			name:    "valid at 0,0 res 15",
			args:    args{ll: NewLatLng(0, 0), res: 15},
			want:    0x8f754e64992d6d8,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCellFromLatLng(tt.args.ll, tt.args.res)
			if !tt.wantErr(t, err, fmt.Sprintf("NewCellFromLatLng(%v, %v)", tt.args.ll, tt.args.res)) {
				return
			}
			assert.Equalf(t, tt.want, got, "NewCellFromLatLng(%v, %v)", tt.args.ll, tt.args.res)
		})
	}
}

func Test_faceIJKToH3(t *testing.T) {
	type args struct {
		fijk faceIJK
		res  int
	}
	tests := []struct {
		name    string
		args    args
		want    Cell
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "i out of bounds at res 0",
			args:    args{fijk: faceIJK{face: 0, coord: coordIJK{i: 3, j: 0, k: 0}}, res: 0},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "j out of bounds at res 0",
			args:    args{fijk: faceIJK{face: 1, coord: coordIJK{i: 0, j: 4, k: 0}}, res: 0},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "k out of bounds at res 0",
			args:    args{fijk: faceIJK{face: 2, coord: coordIJK{i: 2, j: 0, k: 5}}, res: 0},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "i out of bounds at res 1",
			args:    args{fijk: faceIJK{face: 3, coord: coordIJK{i: 6, j: 0, k: 0}}, res: 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "j out of bounds at res 1",
			args:    args{fijk: faceIJK{face: 4, coord: coordIJK{i: 0, j: 7, k: 1}}, res: 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "k out of bounds at res 1",
			args:    args{fijk: faceIJK{face: 5, coord: coordIJK{i: 2, j: 0, k: 8}}, res: 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "i out of bounds at res 2",
			args:    args{fijk: faceIJK{face: 6, coord: coordIJK{i: 18, j: 0, k: 0}}, res: 2},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "j out of bounds at res 2",
			args:    args{fijk: faceIJK{face: 7, coord: coordIJK{i: 0, j: 19, k: 1}}, res: 2},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "k out of bounds at res 2",
			args:    args{fijk: faceIJK{face: 8, coord: coordIJK{i: 2, j: 0, k: 20}}, res: 2},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "valid res 0",
			args:    args{fijk: faceIJK{face: 0, coord: coordIJK{i: 0, j: 0, k: 0}}, res: 0},
			want:    Cell(0x8021fffffffffff),
			wantErr: assert.NoError,
		},
		{
			name:    "valid res 1",
			args:    args{fijk: faceIJK{face: 1, coord: coordIJK{i: 1, j: 0, k: 0}}, res: 1},
			want:    Cell(0x81053ffffffffff),
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := faceIJKToH3(tt.args.fijk, tt.args.res)
			if !tt.wantErr(t, err, fmt.Sprintf("faceIJKToH3(%v, %v)", tt.args.fijk, tt.args.res)) {
				return
			}
			assert.Equalf(t, tt.want, got, "faceIJKToH3(%v, %v)", tt.args.fijk, tt.args.res)
		})
	}
}

func TestCell_isPentagon(t *testing.T) {
	tests := []struct {
		name string
		c    Cell
		want bool
	}{
		{
			name: "pentagon that comes from sf test",
			c:    mustCellFromString("801dfffffffffff"),
			want: true,
		},
		{
			name: "from isPentagon.txt - not a pentagon",
			c:    mustCellFromString("85283473fffffff"),
			want: false,
		},
		{
			name: "0 is not a pentagon",
			c:    0,
			want: false,
		},
		{
			name: "all but high bit is not a pentagon",
			c:    0x7fffffffffffffff,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.c.isPentagon(), "isPentagon()")
		})
	}
}
