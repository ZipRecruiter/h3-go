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
		var rotations int

		_, rotations, err := origin.neighborRotations(-1, rotations)
		assert.Error(t, err, "invalid direction should fail")
		_, rotations, err = origin.neighborRotations(7, rotations)
		assert.Error(t, err, "invalid direction should fail")
		_, rotations, err = origin.neighborRotations(100, rotations)
		assert.Error(t, err, "invalid direction should fail")
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

func TestCell_GridDistance(t *testing.T) {
	type args struct {
		other Cell
	}
	tests := []struct {
		name    string
		c       Cell
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "identity",
			c:       0x81283ffffffffff,
			args:    args{other: 0x81283ffffffffff},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name:    "next to each other",
			c:       0x872830874ffffff,
			args:    args{other: 0x872830876ffffff},
			want:    1,
			wantErr: assert.NoError,
		},
		{
			name:    "one further away",
			c:       0x872830874ffffff,
			args:    args{other: 0x87283080dffffff},
			want:    2,
			wantErr: assert.NoError,
		},
		{
			name:    "three away",
			c:       0x872830874ffffff,
			args:    args{other: 0x872830808ffffff},
			want:    3,
			wantErr: assert.NoError,
		},
		{
			name:    "from examples/distance.c",
			c:       0x8f2830828052d25,              // 1455 Market St @ resolution 15
			args:    args{other: 0x8f283082a30e623}, // 555 Market St @ resolution 15
			want:    2340,
			wantErr: assert.NoError,
		},
		{
			name:    "resolution mismatch",
			c:       0x832830fffffffff,
			args:    args{other: 0x822837fffffffff},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "distance from invalid cell",
			c:       0xffffffffffffffff,
			args:    args{other: 0xffffffffffffffff},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.GridDistance(tt.args.other)
			if !tt.wantErr(t, err, fmt.Sprintf("GridDistance(%v)", tt.args.other)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GridDistance(%v)", tt.args.other)
		})
	}
}

func TestCell_ToLatLng(t *testing.T) {
	// Ground truth values from h3-python (uber/h3 C library)
	tests := []struct {
		name    string
		cell    string
		wantLat float64 // degrees
		wantLng float64 // degrees
	}{
		{name: "SF res 5", cell: "85283473fffffff", wantLat: 37.345793375368480, wantLng: -121.976375972551196},
		{name: "LA cell 1 res 7", cell: "8729a19a0ffffff", wantLat: 33.955170370550128, wantLng: -118.460825162371194},
		{name: "LA cell 2 res 7", cell: "8729a1811ffffff", wantLat: 34.113888958192490, wantLng: -118.561069531613768},
		{name: "Bay Area res 7", cell: "872830829ffffff", wantLat: 37.782796841289489, wantLng: -122.443404411195274},
		{name: "NYC res 7", cell: "872a1072bffffff", wantLat: 40.685283975232352, wantLng: -74.030218323463984},
		{name: "London res 7", cell: "87195da49ffffff", wantLat: 51.509714188032547, wantLng: -0.142519940884726},
		{name: "Tokyo res 7", cell: "872f5a309ffffff", wantLat: 35.758394212166600, wantLng: 139.716969424166905},
		{name: "Near North Pole res 2", cell: "82001ffffffffff", wantLat: 81.029276156855076, wantLng: 50.504634769541724},
		{name: "origin res 1", cell: "81757ffffffffff", wantLat: 2.097907640587099, wantLng: 0.499473284285551},
		{name: "origin res 2", cell: "82754ffffffffff", wantLat: -0.459695642341287, wantLng: 0.534167862982357},
		{name: "origin res 15", cell: "8f754e64992d6d8", wantLat: 0.000002077039499, wantLng: 0.000002363229518},
		{name: "base cell 0 res 0", cell: "8021fffffffffff", wantLat: 46.041894318837713, wantLng: 71.527903299099236},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mustCellFromString(tt.cell)
			got, err := c.ToLatLng()
			assert.NoError(t, err)
			gotLat := rad2deg(got.Latitude())
			gotLng := rad2deg(got.Longitude())
			// Tolerance: 1e-7 degrees ≈ 1cm
			assert.InDelta(t, tt.wantLat, gotLat, 1e-7, "latitude")
			assert.InDelta(t, tt.wantLng, gotLng, 1e-7, "longitude")
		})
	}
}

func TestCell_ToLatLng_pentagons(t *testing.T) {
	tests := []struct {
		name    string
		cell    string
		wantLat float64
		wantLng float64
	}{
		{name: "pentagon bc=4", cell: "8009fffffffffff", wantLat: 64.700000127934871, wantLng: 10.536199075467669},
		{name: "pentagon bc=14", cell: "801dfffffffffff", wantLat: 50.103201482241332, wantLng: -143.478490015025159},
		{name: "pentagon bc=24", cell: "8031fffffffffff", wantLat: 39.100000033975917, wantLng: 122.300000407787039},
		{name: "pentagon bc=38", cell: "804dfffffffffff", wantLat: 23.717925271222960, wantLng: -67.132326366435663},
		{name: "pentagon bc=49", cell: "8063fffffffffff", wantLat: 10.447345187511031, wantLng: 58.157705839572586},
		{name: "pentagon bc=58", cell: "8075fffffffffff", wantLat: 2.300882111626741, wantLng: -5.245390296777326},
		{name: "pentagon bc=63", cell: "807ffffffffffff", wantLat: -2.300882111626741, wantLng: 174.754609703222684},
		{name: "pentagon bc=72", cell: "8091fffffffffff", wantLat: -10.447345187511036, wantLng: -121.842294160427400},
		{name: "pentagon bc=83", cell: "80a7fffffffffff", wantLat: -23.717925271222960, wantLng: 112.867673633564351},
		{name: "pentagon bc=97", cell: "80c3fffffffffff", wantLat: -39.100000033975910, wantLng: -57.699999592212968},
		{name: "pentagon bc=107", cell: "80d7fffffffffff", wantLat: -50.103201482241332, wantLng: 36.521509984974834},
		{name: "pentagon bc=117", cell: "80ebfffffffffff", wantLat: -64.700000127934885, wantLng: -169.463800924532393},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := mustCellFromString(tt.cell)
			assert.True(t, c.isPentagon(), "expected pentagon")
			got, err := c.ToLatLng()
			assert.NoError(t, err)
			gotLat := rad2deg(got.Latitude())
			gotLng := rad2deg(got.Longitude())
			assert.InDelta(t, tt.wantLat, gotLat, 1e-7, "latitude")
			assert.InDelta(t, tt.wantLng, gotLng, 1e-7, "longitude")
		})
	}
}

func TestCell_ToLatLng_allBaseCells(t *testing.T) {
	for i := 0; i < NUM_BASE_CELLS; i++ {
		c := H3_INIT.setMode(H3_CELL_MODE).setBaseCell(baseCell(i))
		if !c.Valid() {
			continue
		}
		t.Run(fmt.Sprintf("bc_%d", i), func(t *testing.T) {
			ll, err := c.ToLatLng()
			assert.NoError(t, err, "ToLatLng should not error for valid base cell %d", i)

			// Round-trip: the cell containing ll at res 0 should be c
			roundTrip, err := NewCellFromLatLng(ll, 0)
			assert.NoError(t, err)
			assert.Equal(t, c, roundTrip, "round-trip failed for base cell %d", i)
		})
	}
}

func TestCell_ToLatLng_roundTrip(t *testing.T) {
	testPoints := []struct{ lat, lng float64 }{
		{37.7749, -122.4194},
		{34.0522, -118.2437},
		{40.7128, -74.0060},
		{51.5074, -0.1278},
		{35.6762, 139.6503},
		{0.0, 0.0},
		{89.9, 1.0},
		{-89.9, 1.0},
		{0.0, 179.99},
		{0.0, -179.99},
		{-33.8688, 151.2093},
		{55.7558, 37.6173},
		{-22.9068, -43.1729},
	}
	for _, tp := range testPoints {
		for _, res := range []int{0, 1, 3, 5, 7, 10, 12, 15} {
			name := fmt.Sprintf("lat%.1f_lng%.1f_res%d", tp.lat, tp.lng, res)
			t.Run(name, func(t *testing.T) {
				ll := NewLatLng(tp.lat, tp.lng)
				cell, err := NewCellFromLatLng(ll, res)
				assert.NoError(t, err)

				center, err := cell.ToLatLng()
				assert.NoError(t, err)

				cell2, err := NewCellFromLatLng(center, res)
				assert.NoError(t, err)
				assert.Equal(t, cell, cell2, "round-trip cell mismatch")
			})
		}
	}
}

func TestCell_ToLatLng_invalidCells(t *testing.T) {
	tests := []struct {
		name string
		cell Cell
	}{
		{name: "zero cell", cell: Cell(0)},
		{name: "wrong mode", cell: H3_INIT.setMode(2)},
		{name: "reserved bits set", cell: H3_INIT.setMode(H3_CELL_MODE).setReservedBits(1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.cell.ToLatLng()
			assert.Error(t, err, "expected error for invalid cell")
		})
	}
}

func TestCell_Parent(t *testing.T) {
	type args struct {
		res int
	}
	tests := []struct {
		name    string
		c       Cell
		args    args
		want    Cell
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "valid parent",
			c:       0x872830829ffffff,
			args:    args{res: 2},
			want:    0x822837fffffffff,
			wantErr: assert.NoError,
		},
		{
			name:    "res is too low",
			c:       0x872830829ffffff,
			args:    args{res: -1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "res is too high",
			c:       0x872830829ffffff,
			args:    args{res: MAX_H3_RES + 1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "res is same",
			c:       0x872830829ffffff,
			args:    args{res: 7},
			want:    0x872830829ffffff,
			wantErr: assert.NoError,
		},
		{
			name:    "res is higher",
			c:       0x872830829ffffff,
			args:    args{res: 16},
			want:    0,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Parent(tt.args.res)
			if !tt.wantErr(t, err, fmt.Sprintf("Parent(%v)", tt.args.res)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Parent(%v)", tt.args.res)
		})
	}
}
