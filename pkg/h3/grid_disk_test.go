package h3

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxGridDiskSize(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "invalid",
			args:    args{k: -1},
			want:    0,
			wantErr: assert.Error,
		},
		{
			name:    "large",
			args:    args{k: 26755},
			want:    2147570341,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := maxGridDiskSize(tt.args.k)
			if !tt.wantErr(t, err, fmt.Sprintf("maxGridDiskSize(%v)", tt.args.k)) {
				return
			}
			assert.Equalf(t, tt.want, got, "maxGridDiskSize(%v)", tt.args.k)
		})
	}
}

func Test_maxGridDiskSize_numCells(t *testing.T) {
	maxCells, err := GetNumCells(15)
	assert.NoError(t, err)

	// 13780510 will produce values above max
	prev := 0
	for k := 13780510 - 100; k < 13780510+100; k++ {
		size, err := maxGridDiskSize(k)
		assert.NoError(t, err)
		assert.LessOrEqual(t, uint64(size), maxCells)
		assert.LessOrEqual(t, prev, size)
		prev = size
	}

	size, err := maxGridDiskSize(math.MaxInt32)
	assert.NoError(t, err)
	assert.Equal(t, maxCells, uint64(size))
}

func TestCell_gridDisk(t *testing.T) {
	t.Run("invalid k", func(t *testing.T) {
		sfLL := NewLatLng(37.813318, -122.40929)
		sf, err := NewCellFromLatLng(sfLL, 0)
		assert.NoError(t, err)

		cells, distances, err := sf.gridDiskDistancesUnsafe(-1)
		assert.Error(t, err)
		assert.Nil(t, cells)
		assert.Nil(t, distances)

		cells, distances, err = sf.gridDiskDistancesSafe(-1)
		assert.Error(t, err)
		assert.Nil(t, cells)
		assert.Nil(t, distances)
	})

	t.Run("invalid cell", func(t *testing.T) {
		_, err := Cell(0x7fffffffffffffff).GridDisk(1000)
		assert.Error(t, err, "should not be able to create a grid disk from an invalid index")
	})

	t.Run("san francisco k=0", func(t *testing.T) {
		// This should run into a pentagon and go into the slower, recursive "safe" path.
		sfLL := NewLatLng(37.813318, -122.40929)
		sf, err := NewCellFromLatLng(sfLL, 0)
		assert.NoError(t, err)

		cells, distances, err := sf.GridDiskDistances(0)
		assert.NoError(t, err)
		assert.Len(t, cells, 1)
		assert.Len(t, distances, 1)

		expectedDistances := []int{0}
		assert.Equal(t, expectedDistances, distances)
		expectedCells := []Cell{sf}
		assert.Equal(t, expectedCells, cells)
	})

	t.Run("san francisco res=1 k=1", func(t *testing.T) {
		// This should not run into a pentagon and should be able to use the faster, "unsafe" path.
		sfLL := NewLatLng(37.813318, -122.40929)
		sf, err := NewCellFromLatLng(sfLL, 1)
		assert.NoError(t, err)

		cells, distances, err := sf.GridDiskDistances(1)
		assert.NoError(t, err)
		assert.Len(t, cells, 7)
		assert.Len(t, distances, 7)

		// First cell is the origin, so it should have a distance of 0. The rest should be 1.
		expectedDistances := []int{0, 1, 1, 1, 1, 1, 1}
		assert.Equal(t, expectedDistances, distances)

		expectedCells := []Cell{
			sf,
			mustCellFromString("8129bffffffffff"),
			mustCellFromString("8128bffffffffff"),
			mustCellFromString("8128fffffffffff"),
			mustCellFromString("81287ffffffffff"),
			mustCellFromString("81297ffffffffff"),
			mustCellFromString("81293ffffffffff"),
		}
		assert.Equal(t, expectedCells, cells)
	})

	t.Run("san francisco res=0 k=1", func(t *testing.T) {
		sfLL := NewLatLng(37.813318, -122.40929)
		sf, err := NewCellFromLatLng(sfLL, 0)
		assert.NoError(t, err)

		cells, distances, err := sf.GridDiskDistances(1)
		assert.NoError(t, err)
		assert.Len(t, cells, 7)
		assert.Len(t, distances, 7)

		// First cell is the origin, so it should have a distance of 0. The rest should be 1.
		expectedDistances := []int{1, 1, 1, 0, 1, 1, 1}
		assert.Equal(t, expectedDistances, distances)

		expectedCells := []Cell{
			mustCellFromString("8051fffffffffff"),
			mustCellFromString("8049fffffffffff"),
			mustCellFromString("8027fffffffffff"),
			sf,
			mustCellFromString("801dfffffffffff"),
			mustCellFromString("8037fffffffffff"),
			mustCellFromString("8013fffffffffff"),
		}
		assert.Equal(t, expectedCells, cells)
	})
}
