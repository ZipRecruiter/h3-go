package h3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDimByCIIres(t *testing.T) {
	tests := []struct {
		name string
		res  int
		want int
	}{
		{name: "res 0", res: 0, want: 2},
		{name: "res 2", res: 2, want: 14},
		{name: "res 4", res: 4, want: 98},
		{name: "res 6", res: 6, want: 686},
		{name: "res 8", res: 8, want: 4802},
		{name: "res 10", res: 10, want: 33614},
		{name: "res 12", res: 12, want: 235298},
		{name: "res 14", res: 14, want: 1647086},
		{name: "res 16", res: 16, want: 11529602},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxDimByCIIres[tt.res])
		})
	}
}

func TestUnitScaleByCIIres(t *testing.T) {
	tests := []struct {
		name string
		res  int
		want int
	}{
		{name: "res 0", res: 0, want: 1},
		{name: "res 2", res: 2, want: 7},
		{name: "res 4", res: 4, want: 49},
		{name: "res 6", res: 6, want: 343},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, unitScaleByCIIres[tt.res])
		})
	}
}

func TestAdjustOverageClassII_noOverage(t *testing.T) {
	// A FaceIJK well within the face boundary should not be adjusted
	fijk := faceIJK{face: 0, coord: coordIJK{i: 0, j: 0, k: 0}}
	overage := adjustOverageClassII(&fijk, 0, false, false)
	assert.Equal(t, NO_OVERAGE, overage)
	assert.Equal(t, 0, fijk.face)
	assert.Equal(t, coordIJK{0, 0, 0}, fijk.coord)
}

func TestAdjustOverageClassII_atBoundary(t *testing.T) {
	// At res 0, maxDim=2. ijk sum of 2 is at boundary but not over.
	fijk := faceIJK{face: 0, coord: coordIJK{i: 1, j: 1, k: 0}}
	overage := adjustOverageClassII(&fijk, 0, false, false)
	assert.Equal(t, NO_OVERAGE, overage)
	assert.Equal(t, 0, fijk.face)
}

func TestAdjustOverageClassII_ijOverage(t *testing.T) {
	// At res 0, maxDim=2. ijk sum of 3 with k=0, j=0 → ij overage
	fijk := faceIJK{face: 0, coord: coordIJK{i: 3, j: 0, k: 0}}
	overage := adjustOverageClassII(&fijk, 0, false, false)
	assert.Equal(t, NEW_FACE, overage)
	// Should move to face 4 (IJ neighbor of face 0)
	assert.Equal(t, 4, fijk.face)
}

func TestAdjustOverageClassII_jkOverage(t *testing.T) {
	// At res 0, maxDim=2. Both j and k > 0 → jk overage
	fijk := faceIJK{face: 0, coord: coordIJK{i: 0, j: 2, k: 1}}
	overage := adjustOverageClassII(&fijk, 0, false, false)
	assert.Equal(t, NEW_FACE, overage)
	// Should move to face 5 (JK neighbor of face 0)
	assert.Equal(t, 5, fijk.face)
}

func TestAdjustOverageClassII_kiOverage(t *testing.T) {
	// At res 0, maxDim=2. k > 0, j == 0 → ki overage
	fijk := faceIJK{face: 0, coord: coordIJK{i: 0, j: 0, k: 3}}
	overage := adjustOverageClassII(&fijk, 0, false, false)
	assert.Equal(t, NEW_FACE, overage)
	// Should move to face 1 (KI neighbor of face 0)
	assert.Equal(t, 1, fijk.face)
}

func TestAdjustOverageClassII_substrate_faceEdge(t *testing.T) {
	// With substrate, maxDim=6 (2*3). Exact boundary = FACE_EDGE
	fijk := faceIJK{face: 0, coord: coordIJK{i: 3, j: 3, k: 0}}
	overage := adjustOverageClassII(&fijk, 0, false, true)
	assert.Equal(t, FACE_EDGE, overage)
	// Face should not change for edge case
	assert.Equal(t, 0, fijk.face)
}
