package h3

type baseCell int

const (
	// INVALID_BASE_CELL indicates an invalid base cell.
	INVALID_BASE_CELL = 127
	// INVALID_ROTATIONS indicates an invalid number of rotations.
	INVALID_ROTATIONS = -1
	// MAX_FACE_COORD is the maximum input for any component to face-to-base-cell lookup functions.
	MAX_FACE_COORD = 2
)

var (
	// baseCellData the resolution 0 base cell data table.
	//
	// For each base cell, gives the "home" face and ijk+ coordinates on that face,
	// whether the base cell is a pentagon. Additionally, if the base cell is a
	// pentagon, the two cw offset rotation adjacent faces are given (-1 indicates
	// that no cw offset rotation faces exist for this base cell).
	baseCellData = [NUM_BASE_CELLS]struct {
		// homeFijk is the 'home' face and normalized ijk coordinates on that face
		homeFijk faceIJK
		// isPentagon is true if this base cell is a pentagon
		isPentagon bool
		// cwOffsetPent is the two clockwise offset faces if the base cell is a pentagon
		cwOffsetPent [2]int
	}{
		{homeFijk: faceIJK{face: 1, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 0
		{homeFijk: faceIJK{face: 2, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 1
		{homeFijk: faceIJK{face: 1, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 2
		{homeFijk: faceIJK{face: 2, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 3
		{homeFijk: faceIJK{face: 0, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{-1, -1}},  // base cell 4
		{homeFijk: faceIJK{face: 1, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 5
		{homeFijk: faceIJK{face: 1, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 6
		{homeFijk: faceIJK{face: 2, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 7
		{homeFijk: faceIJK{face: 0, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 8
		{homeFijk: faceIJK{face: 2, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 9
		{homeFijk: faceIJK{face: 1, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 10
		{homeFijk: faceIJK{face: 1, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 11
		{homeFijk: faceIJK{face: 3, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 12
		{homeFijk: faceIJK{face: 3, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 13
		{homeFijk: faceIJK{face: 11, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{2, 6}},   // base cell 14
		{homeFijk: faceIJK{face: 4, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 15
		{homeFijk: faceIJK{face: 0, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 16
		{homeFijk: faceIJK{face: 6, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 17
		{homeFijk: faceIJK{face: 0, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 18
		{homeFijk: faceIJK{face: 2, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 19
		{homeFijk: faceIJK{face: 7, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 20
		{homeFijk: faceIJK{face: 2, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 21
		{homeFijk: faceIJK{face: 0, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 22
		{homeFijk: faceIJK{face: 6, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 23
		{homeFijk: faceIJK{face: 10, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{1, 5}},   // base cell 24
		{homeFijk: faceIJK{face: 6, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 25
		{homeFijk: faceIJK{face: 3, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 26
		{homeFijk: faceIJK{face: 11, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 27
		{homeFijk: faceIJK{face: 4, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 28
		{homeFijk: faceIJK{face: 3, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 29
		{homeFijk: faceIJK{face: 0, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 30
		{homeFijk: faceIJK{face: 4, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 31
		{homeFijk: faceIJK{face: 5, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 32
		{homeFijk: faceIJK{face: 0, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 33
		{homeFijk: faceIJK{face: 7, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 34
		{homeFijk: faceIJK{face: 11, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 35
		{homeFijk: faceIJK{face: 7, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 36
		{homeFijk: faceIJK{face: 10, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 37
		{homeFijk: faceIJK{face: 12, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{3, 7}},   // base cell 38
		{homeFijk: faceIJK{face: 6, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 39
		{homeFijk: faceIJK{face: 7, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 40
		{homeFijk: faceIJK{face: 4, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 41
		{homeFijk: faceIJK{face: 3, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 42
		{homeFijk: faceIJK{face: 3, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 43
		{homeFijk: faceIJK{face: 4, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 44
		{homeFijk: faceIJK{face: 6, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 45
		{homeFijk: faceIJK{face: 11, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 46
		{homeFijk: faceIJK{face: 8, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 47
		{homeFijk: faceIJK{face: 5, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 48
		{homeFijk: faceIJK{face: 14, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{0, 9}},   // base cell 49
		{homeFijk: faceIJK{face: 5, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 50
		{homeFijk: faceIJK{face: 12, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 51
		{homeFijk: faceIJK{face: 10, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 52
		{homeFijk: faceIJK{face: 4, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 53
		{homeFijk: faceIJK{face: 12, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 54
		{homeFijk: faceIJK{face: 7, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 55
		{homeFijk: faceIJK{face: 11, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 56
		{homeFijk: faceIJK{face: 10, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 57
		{homeFijk: faceIJK{face: 13, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{4, 8}},   // base cell 58
		{homeFijk: faceIJK{face: 10, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 59
		{homeFijk: faceIJK{face: 11, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 60
		{homeFijk: faceIJK{face: 9, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 61
		{homeFijk: faceIJK{face: 8, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 62
		{homeFijk: faceIJK{face: 6, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{11, 15}},  // base cell 63
		{homeFijk: faceIJK{face: 8, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 64
		{homeFijk: faceIJK{face: 9, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 65
		{homeFijk: faceIJK{face: 14, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 66
		{homeFijk: faceIJK{face: 5, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 67
		{homeFijk: faceIJK{face: 16, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 68
		{homeFijk: faceIJK{face: 8, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 69
		{homeFijk: faceIJK{face: 5, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 70
		{homeFijk: faceIJK{face: 12, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 71
		{homeFijk: faceIJK{face: 7, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{12, 16}},  // base cell 72
		{homeFijk: faceIJK{face: 12, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 73
		{homeFijk: faceIJK{face: 10, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 74
		{homeFijk: faceIJK{face: 9, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 75
		{homeFijk: faceIJK{face: 13, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 76
		{homeFijk: faceIJK{face: 16, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 77
		{homeFijk: faceIJK{face: 15, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 78
		{homeFijk: faceIJK{face: 15, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 79
		{homeFijk: faceIJK{face: 16, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 80
		{homeFijk: faceIJK{face: 14, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 81
		{homeFijk: faceIJK{face: 13, coord: coordIJK{1, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 82
		{homeFijk: faceIJK{face: 5, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{10, 19}},  // base cell 83
		{homeFijk: faceIJK{face: 8, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 84
		{homeFijk: faceIJK{face: 14, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 85
		{homeFijk: faceIJK{face: 9, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 86
		{homeFijk: faceIJK{face: 14, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 87
		{homeFijk: faceIJK{face: 17, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 88
		{homeFijk: faceIJK{face: 12, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 89
		{homeFijk: faceIJK{face: 16, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 90
		{homeFijk: faceIJK{face: 17, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 91
		{homeFijk: faceIJK{face: 15, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 92
		{homeFijk: faceIJK{face: 16, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 93
		{homeFijk: faceIJK{face: 9, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},   // base cell 94
		{homeFijk: faceIJK{face: 15, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 95
		{homeFijk: faceIJK{face: 13, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 96
		{homeFijk: faceIJK{face: 8, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{13, 17}},  // base cell 97
		{homeFijk: faceIJK{face: 13, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 98
		{homeFijk: faceIJK{face: 17, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 99
		{homeFijk: faceIJK{face: 19, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 100
		{homeFijk: faceIJK{face: 14, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 101
		{homeFijk: faceIJK{face: 19, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 102
		{homeFijk: faceIJK{face: 17, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 103
		{homeFijk: faceIJK{face: 13, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 104
		{homeFijk: faceIJK{face: 17, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 105
		{homeFijk: faceIJK{face: 16, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 106
		{homeFijk: faceIJK{face: 9, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{14, 18}},  // base cell 107
		{homeFijk: faceIJK{face: 15, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 108
		{homeFijk: faceIJK{face: 15, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 109
		{homeFijk: faceIJK{face: 18, coord: coordIJK{0, 1, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 110
		{homeFijk: faceIJK{face: 18, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 111
		{homeFijk: faceIJK{face: 19, coord: coordIJK{0, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 112
		{homeFijk: faceIJK{face: 17, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 113
		{homeFijk: faceIJK{face: 19, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 114
		{homeFijk: faceIJK{face: 18, coord: coordIJK{0, 1, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 115
		{homeFijk: faceIJK{face: 18, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 116
		{homeFijk: faceIJK{face: 19, coord: coordIJK{2, 0, 0}}, isPentagon: true, cwOffsetPent: [2]int{-1, -1}}, // base cell 117
		{homeFijk: faceIJK{face: 19, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 118
		{homeFijk: faceIJK{face: 18, coord: coordIJK{0, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 119
		{homeFijk: faceIJK{face: 19, coord: coordIJK{1, 0, 1}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 120
		{homeFijk: faceIJK{face: 18, coord: coordIJK{1, 0, 0}}, isPentagon: false, cwOffsetPent: [2]int{0, 0}},  // base cell 121
	}

	baseCellNeighbors = [NUM_BASE_CELLS][7]baseCell{
		{0, 1, 5, 2, 4, 3, 8},                             // base cell 0
		{1, 7, 6, 9, 0, 3, 2},                             // base cell 1
		{2, 6, 10, 11, 0, 1, 5},                           // base cell 2
		{3, 13, 1, 7, 4, 12, 0},                           // base cell 3
		{4, INVALID_BASE_CELL, 15, 8, 3, 0, 12},           // base cell 4 (pentagon)
		{5, 2, 18, 10, 8, 0, 16},                          // base cell 5
		{6, 14, 11, 17, 1, 9, 2},                          // base cell 6
		{7, 21, 9, 19, 3, 13, 1},                          // base cell 7
		{8, 5, 22, 16, 4, 0, 15},                          // base cell 8
		{9, 19, 14, 20, 1, 7, 6},                          // base cell 9
		{10, 11, 24, 23, 5, 2, 18},                        // base cell 10
		{11, 17, 23, 25, 2, 6, 10},                        // base cell 11
		{12, 28, 13, 26, 4, 15, 3},                        // base cell 12
		{13, 26, 21, 29, 3, 12, 7},                        // base cell 13
		{14, INVALID_BASE_CELL, 17, 27, 9, 20, 6},         // base cell 14 (pentagon)
		{15, 22, 28, 31, 4, 8, 12},                        // base cell 15
		{16, 18, 33, 30, 8, 5, 22},                        // base cell 16
		{17, 11, 14, 6, 35, 25, 27},                       // base cell 17
		{18, 24, 30, 32, 5, 10, 16},                       // base cell 18
		{19, 34, 20, 36, 7, 21, 9},                        // base cell 19
		{20, 14, 19, 9, 40, 27, 36},                       // base cell 20
		{21, 38, 19, 34, 13, 29, 7},                       // base cell 21
		{22, 16, 41, 33, 15, 8, 31},                       // base cell 22
		{23, 24, 11, 10, 39, 37, 25},                      // base cell 23
		{24, INVALID_BASE_CELL, 32, 37, 10, 23, 18},       // base cell 24 (pentagon)
		{25, 23, 17, 11, 45, 39, 35},                      // base cell 25
		{26, 42, 29, 43, 12, 28, 13},                      // base cell 26
		{27, 40, 35, 46, 14, 20, 17},                      // base cell 27
		{28, 31, 42, 44, 12, 15, 26},                      // base cell 28
		{29, 43, 38, 47, 13, 26, 21},                      // base cell 29
		{30, 32, 48, 50, 16, 18, 33},                      // base cell 30
		{31, 41, 44, 53, 15, 22, 28},                      // base cell 31
		{32, 30, 24, 18, 52, 50, 37},                      // base cell 32
		{33, 30, 49, 48, 22, 16, 41},                      // base cell 33
		{34, 19, 38, 21, 54, 36, 51},                      // base cell 34
		{35, 46, 45, 56, 17, 27, 25},                      // base cell 35
		{36, 20, 34, 19, 55, 40, 54},                      // base cell 36
		{37, 39, 52, 57, 24, 23, 32},                      // base cell 37
		{38, INVALID_BASE_CELL, 34, 51, 29, 47, 21},       // base cell 38 (pentagon)
		{39, 37, 25, 23, 59, 57, 45},                      // base cell 39
		{40, 27, 36, 20, 60, 46, 55},                      // base cell 40
		{41, 49, 53, 61, 22, 33, 31},                      // base cell 41
		{42, 58, 43, 62, 28, 44, 26},                      // base cell 42
		{43, 62, 47, 64, 26, 42, 29},                      // base cell 43
		{44, 53, 58, 65, 28, 31, 42},                      // base cell 44
		{45, 39, 35, 25, 63, 59, 56},                      // base cell 45
		{46, 60, 56, 68, 27, 40, 35},                      // base cell 46
		{47, 38, 43, 29, 69, 51, 64},                      // base cell 47
		{48, 49, 30, 33, 67, 66, 50},                      // base cell 48
		{49, INVALID_BASE_CELL, 61, 66, 33, 48, 41},       // base cell 49 (pentagon)
		{50, 48, 32, 30, 70, 67, 52},                      // base cell 50
		{51, 69, 54, 71, 38, 47, 34},                      // base cell 51
		{52, 57, 70, 74, 32, 37, 50},                      // base cell 52
		{53, 61, 65, 75, 31, 41, 44},                      // base cell 53
		{54, 71, 55, 73, 34, 51, 36},                      // base cell 54
		{55, 40, 54, 36, 72, 60, 73},                      // base cell 55
		{56, 68, 63, 77, 35, 46, 45},                      // base cell 56
		{57, 59, 74, 78, 37, 39, 52},                      // base cell 57
		{58, INVALID_BASE_CELL, 62, 76, 44, 65, 42},       // base cell 58 (pentagon)
		{59, 63, 78, 79, 39, 45, 57},                      // base cell 59
		{60, 72, 68, 80, 40, 55, 46},                      // base cell 60
		{61, 53, 49, 41, 81, 75, 66},                      // base cell 61
		{62, 43, 58, 42, 82, 64, 76},                      // base cell 62
		{63, INVALID_BASE_CELL, 56, 45, 79, 59, 77},       // base cell 63 (pentagon)
		{64, 47, 62, 43, 84, 69, 82},                      // base cell 64
		{65, 58, 53, 44, 86, 76, 75},                      // base cell 65
		{66, 67, 81, 85, 49, 48, 61},                      // base cell 66
		{67, 66, 50, 48, 87, 85, 70},                      // base cell 67
		{68, 56, 60, 46, 90, 77, 80},                      // base cell 68
		{69, 51, 64, 47, 89, 71, 84},                      // base cell 69
		{70, 67, 52, 50, 83, 87, 74},                      // base cell 70
		{71, 89, 73, 91, 51, 69, 54},                      // base cell 71
		{72, INVALID_BASE_CELL, 73, 55, 80, 60, 88},       // base cell 72 (pentagon)
		{73, 91, 72, 88, 54, 71, 55},                      // base cell 73
		{74, 78, 83, 92, 52, 57, 70},                      // base cell 74
		{75, 65, 61, 53, 94, 86, 81},                      // base cell 75
		{76, 86, 82, 96, 58, 65, 62},                      // base cell 76
		{77, 63, 68, 56, 93, 79, 90},                      // base cell 77
		{78, 74, 59, 57, 95, 92, 79},                      // base cell 78
		{79, 78, 63, 59, 93, 95, 77},                      // base cell 79
		{80, 68, 72, 60, 99, 90, 88},                      // base cell 80
		{81, 85, 94, 101, 61, 66, 75},                     // base cell 81
		{82, 96, 84, 98, 62, 76, 64},                      // base cell 82
		{83, INVALID_BASE_CELL, 74, 70, 100, 87, 92},      // base cell 83 (pentagon)
		{84, 69, 82, 64, 97, 89, 98},                      // base cell 84
		{85, 87, 101, 102, 66, 67, 81},                    // base cell 85
		{86, 76, 75, 65, 104, 96, 94},                     // base cell 86
		{87, 83, 102, 100, 67, 70, 85},                    // base cell 87
		{88, 72, 91, 73, 99, 80, 105},                     // base cell 88
		{89, 97, 91, 103, 69, 84, 71},                     // base cell 89
		{90, 77, 80, 68, 106, 93, 99},                     // base cell 90
		{91, 73, 89, 71, 105, 88, 103},                    // base cell 91
		{92, 83, 78, 74, 108, 100, 95},                    // base cell 92
		{93, 79, 90, 77, 109, 95, 106},                    // base cell 93
		{94, 86, 81, 75, 107, 104, 101},                   // base cell 94
		{95, 92, 79, 78, 109, 108, 93},                    // base cell 95
		{96, 104, 98, 110, 76, 86, 82},                    // base cell 96
		{97, INVALID_BASE_CELL, 98, 84, 103, 89, 111},     // base cell 97 (pentagon)
		{98, 110, 97, 111, 82, 96, 84},                    // base cell 98
		{99, 80, 105, 88, 106, 90, 113},                   // base cell 99
		{100, 102, 83, 87, 108, 114, 92},                  // base cell 100
		{101, 102, 107, 112, 81, 85, 94},                  // base cell 101
		{102, 101, 87, 85, 114, 112, 100},                 // base cell 102
		{103, 91, 97, 89, 116, 105, 111},                  // base cell 103
		{104, 107, 110, 115, 86, 94, 96},                  // base cell 104
		{105, 88, 103, 91, 113, 99, 116},                  // base cell 105
		{106, 93, 99, 90, 117, 109, 113},                  // base cell 106
		{107, INVALID_BASE_CELL, 101, 94, 115, 104, 112},  // base cell 107 (pentagon)
		{108, 100, 95, 92, 118, 114, 109},                 // base cell 108
		{109, 108, 93, 95, 117, 118, 106},                 // base cell 109
		{110, 98, 104, 96, 119, 111, 115},                 // base cell 110
		{111, 97, 110, 98, 116, 103, 119},                 // base cell 111
		{112, 107, 102, 101, 120, 115, 114},               // base cell 112
		{113, 99, 116, 105, 117, 106, 121},                // base cell 113
		{114, 112, 100, 102, 118, 120, 108},               // base cell 114
		{115, 110, 107, 104, 120, 119, 112},               // base cell 115
		{116, 103, 119, 111, 113, 105, 121},               // base cell 116
		{117, INVALID_BASE_CELL, 109, 118, 113, 121, 106}, // base cell 117 (pentagon)
		{118, 120, 108, 114, 117, 121, 109},               // base cell 118
		{119, 111, 115, 110, 121, 116, 120},               // base cell 119
		{120, 115, 114, 112, 121, 119, 118},               // base cell 120
		{121, 116, 120, 119, 117, 113, 118},               // base cell 121
	}

	// baseCellNeighbor60CCWRots lists neighboring base cell rotations in each IJK
	// direction.
	//
	// For each base cell, for each direction, the number of 60 degree CCW rotations
	// to the coordinate system of the neighbor is given. -1 indicates there is no
	// neighbor in that direction.
	baseCellNeighbor60CCWRots = [NUM_BASE_CELLS][7]int{
		{0, 5, 0, 0, 1, 5, 1},  // base cell 0
		{0, 0, 1, 0, 1, 0, 1},  // base cell 1
		{0, 0, 0, 0, 0, 5, 0},  // base cell 2
		{0, 5, 0, 0, 2, 5, 1},  // base cell 3
		{0, -1, 1, 0, 3, 4, 2}, // base cell 4 (pentagon)
		{0, 0, 1, 0, 1, 0, 1},  // base cell 5
		{0, 0, 0, 3, 5, 5, 0},  // base cell 6
		{0, 0, 0, 0, 0, 5, 0},  // base cell 7
		{0, 5, 0, 0, 0, 5, 1},  // base cell 8
		{0, 0, 1, 3, 0, 0, 1},  // base cell 9
		{0, 0, 1, 3, 0, 0, 1},  // base cell 10
		{0, 3, 3, 3, 0, 0, 0},  // base cell 11
		{0, 5, 0, 0, 3, 5, 1},  // base cell 12
		{0, 0, 1, 0, 1, 0, 1},  // base cell 13
		{0, -1, 3, 0, 5, 2, 0}, // base cell 14 (pentagon)
		{0, 5, 0, 0, 4, 5, 1},  // base cell 15
		{0, 0, 0, 0, 0, 5, 0},  // base cell 16
		{0, 3, 3, 3, 3, 0, 3},  // base cell 17
		{0, 0, 0, 3, 5, 5, 0},  // base cell 18
		{0, 3, 3, 3, 0, 0, 0},  // base cell 19
		{0, 3, 3, 3, 0, 3, 0},  // base cell 20
		{0, 0, 0, 3, 5, 5, 0},  // base cell 21
		{0, 0, 1, 0, 1, 0, 1},  // base cell 22
		{0, 3, 3, 3, 0, 3, 0},  // base cell 23
		{0, -1, 3, 0, 5, 2, 0}, // base cell 24 (pentagon)
		{0, 0, 0, 3, 0, 0, 3},  // base cell 25
		{0, 0, 0, 0, 0, 5, 0},  // base cell 26
		{0, 3, 0, 0, 0, 3, 3},  // base cell 27
		{0, 0, 1, 0, 1, 0, 1},  // base cell 28
		{0, 0, 1, 3, 0, 0, 1},  // base cell 29
		{0, 3, 3, 3, 0, 0, 0},  // base cell 30
		{0, 0, 0, 0, 0, 5, 0},  // base cell 31
		{0, 3, 3, 3, 3, 0, 3},  // base cell 32
		{0, 0, 1, 3, 0, 0, 1},  // base cell 33
		{0, 3, 3, 3, 3, 0, 3},  // base cell 34
		{0, 0, 3, 0, 3, 0, 3},  // base cell 35
		{0, 0, 0, 3, 0, 0, 3},  // base cell 36
		{0, 3, 0, 0, 0, 3, 3},  // base cell 37
		{0, -1, 3, 0, 5, 2, 0}, // base cell 38 (pentagon)
		{0, 3, 0, 0, 3, 3, 0},  // base cell 39
		{0, 3, 0, 0, 3, 3, 0},  // base cell 40
		{0, 0, 0, 3, 5, 5, 0},  // base cell 41
		{0, 0, 0, 3, 5, 5, 0},  // base cell 42
		{0, 3, 3, 3, 0, 0, 0},  // base cell 43
		{0, 0, 1, 3, 0, 0, 1},  // base cell 44
		{0, 0, 3, 0, 0, 3, 3},  // base cell 45
		{0, 0, 0, 3, 0, 3, 0},  // base cell 46
		{0, 3, 3, 3, 0, 3, 0},  // base cell 47
		{0, 3, 3, 3, 0, 3, 0},  // base cell 48
		{0, -1, 3, 0, 5, 2, 0}, // base cell 49 (pentagon)
		{0, 0, 0, 3, 0, 0, 3},  // base cell 50
		{0, 3, 0, 0, 0, 3, 3},  // base cell 51
		{0, 0, 3, 0, 3, 0, 3},  // base cell 52
		{0, 3, 3, 3, 0, 0, 0},  // base cell 53
		{0, 0, 3, 0, 3, 0, 3},  // base cell 54
		{0, 0, 3, 0, 0, 3, 3},  // base cell 55
		{0, 3, 3, 3, 0, 0, 3},  // base cell 56
		{0, 0, 0, 3, 0, 3, 0},  // base cell 57
		{0, -1, 3, 0, 5, 2, 0}, // base cell 58 (pentagon)
		{0, 3, 3, 3, 3, 3, 0},  // base cell 59
		{0, 3, 3, 3, 3, 3, 0},  // base cell 60
		{0, 3, 3, 3, 3, 0, 3},  // base cell 61
		{0, 3, 3, 3, 3, 0, 3},  // base cell 62
		{0, -1, 3, 0, 5, 2, 0}, // base cell 63 (pentagon)
		{0, 0, 0, 3, 0, 0, 3},  // base cell 64
		{0, 3, 3, 3, 0, 3, 0},  // base cell 65
		{0, 3, 0, 0, 0, 3, 3},  // base cell 66
		{0, 3, 0, 0, 3, 3, 0},  // base cell 67
		{0, 3, 3, 3, 0, 0, 0},  // base cell 68
		{0, 3, 0, 0, 3, 3, 0},  // base cell 69
		{0, 0, 3, 0, 0, 3, 3},  // base cell 70
		{0, 0, 0, 3, 0, 3, 0},  // base cell 71
		{0, -1, 3, 0, 5, 2, 0}, // base cell 72 (pentagon)
		{0, 3, 3, 3, 0, 0, 3},  // base cell 73
		{0, 3, 3, 3, 0, 0, 3},  // base cell 74
		{0, 0, 0, 3, 0, 0, 3},  // base cell 75
		{0, 3, 0, 0, 0, 3, 3},  // base cell 76
		{0, 0, 0, 3, 0, 5, 0},  // base cell 77
		{0, 3, 3, 3, 0, 0, 0},  // base cell 78
		{0, 0, 1, 3, 1, 0, 1},  // base cell 79
		{0, 0, 1, 3, 1, 0, 1},  // base cell 80
		{0, 0, 3, 0, 3, 0, 3},  // base cell 81
		{0, 0, 3, 0, 3, 0, 3},  // base cell 82
		{0, -1, 3, 0, 5, 2, 0}, // base cell 83 (pentagon)
		{0, 0, 3, 0, 0, 3, 3},  // base cell 84
		{0, 0, 0, 3, 0, 3, 0},  // base cell 85
		{0, 3, 0, 0, 3, 3, 0},  // base cell 86
		{0, 3, 3, 3, 3, 3, 0},  // base cell 87
		{0, 0, 0, 3, 0, 5, 0},  // base cell 88
		{0, 3, 3, 3, 3, 3, 0},  // base cell 89
		{0, 0, 0, 0, 0, 0, 1},  // base cell 90
		{0, 3, 3, 3, 0, 0, 0},  // base cell 91
		{0, 0, 0, 3, 0, 5, 0},  // base cell 92
		{0, 5, 0, 0, 5, 5, 0},  // base cell 93
		{0, 0, 3, 0, 0, 3, 3},  // base cell 94
		{0, 0, 0, 0, 0, 0, 1},  // base cell 95
		{0, 0, 0, 3, 0, 3, 0},  // base cell 96
		{0, -1, 3, 0, 5, 2, 0}, // base cell 97 (pentagon)
		{0, 3, 3, 3, 0, 0, 3},  // base cell 98
		{0, 5, 0, 0, 5, 5, 0},  // base cell 99
		{0, 0, 1, 3, 1, 0, 1},  // base cell 100
		{0, 3, 3, 3, 0, 0, 3},  // base cell 101
		{0, 3, 3, 3, 0, 0, 0},  // base cell 102
		{0, 0, 1, 3, 1, 0, 1},  // base cell 103
		{0, 3, 3, 3, 3, 3, 0},  // base cell 104
		{0, 0, 0, 0, 0, 0, 1},  // base cell 105
		{0, 0, 1, 0, 3, 5, 1},  // base cell 106
		{0, -1, 3, 0, 5, 2, 0}, // base cell 107 (pentagon)
		{0, 5, 0, 0, 5, 5, 0},  // base cell 108
		{0, 0, 1, 0, 4, 5, 1},  // base cell 109
		{0, 3, 3, 3, 0, 0, 0},  // base cell 110
		{0, 0, 0, 3, 0, 5, 0},  // base cell 111
		{0, 0, 0, 3, 0, 5, 0},  // base cell 112
		{0, 0, 1, 0, 2, 5, 1},  // base cell 113
		{0, 0, 0, 0, 0, 0, 1},  // base cell 114
		{0, 0, 1, 3, 1, 0, 1},  // base cell 115
		{0, 5, 0, 0, 5, 5, 0},  // base cell 116
		{0, -1, 1, 0, 3, 4, 2}, // base cell 117 (pentagon)
		{0, 0, 1, 0, 0, 5, 1},  // base cell 118
		{0, 0, 0, 0, 0, 0, 1},  // base cell 119
		{0, 5, 0, 0, 5, 5, 0},  // base cell 120
		{0, 0, 1, 0, 1, 5, 1},  // base cell 121
	}

	// faceIjkBaseCells is a resolution 0 base cell lookup table for each face.
	//
	// Given the face number and a resolution 0 ijk+ coordinate in that face's
	// face-centered ijk coordinate system, gives the base cell located at that
	// coordinate and the number of 60 ccw rotations to rotate into that base
	// cell's orientation.
	//
	// Valid lookup coordinates are from (0, 0, 0) to (2, 2, 2).
	faceIjkBaseCells = [NUM_ICOSA_FACES][3][3][3]struct {
		// bc is the base cell number
		bc baseCell
		// ccwRot60 is the number of ccw 60 degree rotations relative to current face
		ccwRot60 int
	}{
		{ // face 0
			{
				// i 0
				{{16, 0}, {18, 0}, {24, 0}}, // j 0
				{{33, 0}, {30, 0}, {32, 3}}, // j 1
				{{49, 1}, {48, 3}, {50, 3}}, // j 2
			},
			{
				// i 1
				{{8, 0}, {5, 5}, {10, 5}},   // j 0
				{{22, 0}, {16, 0}, {18, 0}}, // j 1
				{{41, 1}, {33, 0}, {30, 0}}, // j 2
			},
			{
				// i 2
				{{4, 0}, {0, 5}, {2, 5}},    // j 0
				{{15, 1}, {8, 0}, {5, 5}},   // j 1
				{{31, 1}, {22, 0}, {16, 0}}, // j 2
			}},
		{ // face 1
			{
				// i 0
				{{2, 0}, {6, 0}, {14, 0}},   // j 0
				{{10, 0}, {11, 0}, {17, 3}}, // j 1
				{{24, 1}, {23, 3}, {25, 3}}, // j 2
			},
			{
				// i 1
				{{0, 0}, {1, 5}, {9, 5}},    // j 0
				{{5, 0}, {2, 0}, {6, 0}},    // j 1
				{{18, 1}, {10, 0}, {11, 0}}, // j 2
			},
			{
				// i 2
				{{4, 1}, {3, 5}, {7, 5}},  // j 0
				{{8, 1}, {0, 0}, {1, 5}},  // j 1
				{{16, 1}, {5, 0}, {2, 0}}, // j 2
			}},
		{ // face 2
			{
				// i 0
				{{7, 0}, {21, 0}, {38, 0}},  // j 0
				{{9, 0}, {19, 0}, {34, 3}},  // j 1
				{{14, 1}, {20, 3}, {36, 3}}, // j 2
			},
			{
				// i 1
				{{3, 0}, {13, 5}, {29, 5}}, // j 0
				{{1, 0}, {7, 0}, {21, 0}},  // j 1
				{{6, 1}, {9, 0}, {19, 0}},  // j 2
			},
			{
				// i 2
				{{4, 2}, {12, 5}, {26, 5}}, // j 0
				{{0, 1}, {3, 0}, {13, 5}},  // j 1
				{{2, 1}, {1, 0}, {7, 0}},   // j 2
			}},
		{ // face 3
			{
				// i 0
				{{26, 0}, {42, 0}, {58, 0}}, // j 0
				{{29, 0}, {43, 0}, {62, 3}}, // j 1
				{{38, 1}, {47, 3}, {64, 3}}, // j 2
			},
			{
				// i 1
				{{12, 0}, {28, 5}, {44, 5}}, // j 0
				{{13, 0}, {26, 0}, {42, 0}}, // j 1
				{{21, 1}, {29, 0}, {43, 0}}, // j 2
			},
			{
				// i 2
				{{4, 3}, {15, 5}, {31, 5}}, // j 0
				{{3, 1}, {12, 0}, {28, 5}}, // j 1
				{{7, 1}, {13, 0}, {26, 0}}, // j 2
			}},
		{ // face 4
			{
				// i 0
				{{31, 0}, {41, 0}, {49, 0}}, // j 0
				{{44, 0}, {53, 0}, {61, 3}}, // j 1
				{{58, 1}, {65, 3}, {75, 3}}, // j 2
			},
			{
				// i 1
				{{15, 0}, {22, 5}, {33, 5}}, // j 0
				{{28, 0}, {31, 0}, {41, 0}}, // j 1
				{{42, 1}, {44, 0}, {53, 0}}, // j 2
			},
			{
				// i 2
				{{4, 4}, {8, 5}, {16, 5}},   // j 0
				{{12, 1}, {15, 0}, {22, 5}}, // j 1
				{{26, 1}, {28, 0}, {31, 0}}, // j 2
			}},
		{ // face 5
			{
				// i 0
				{{50, 0}, {48, 0}, {49, 3}}, // j 0
				{{32, 0}, {30, 3}, {33, 3}}, // j 1
				{{24, 3}, {18, 3}, {16, 3}}, // j 2
			},
			{
				// i 1
				{{70, 0}, {67, 0}, {66, 3}}, // j 0
				{{52, 3}, {50, 0}, {48, 0}}, // j 1
				{{37, 3}, {32, 0}, {30, 3}}, // j 2
			},
			{
				// i 2
				{{83, 0}, {87, 3}, {85, 3}}, // j 0
				{{74, 3}, {70, 0}, {67, 0}}, // j 1
				{{57, 1}, {52, 3}, {50, 0}}, // j 2
			}},
		{ // face 6
			{
				// i 0
				{{25, 0}, {23, 0}, {24, 3}}, // j 0
				{{17, 0}, {11, 3}, {10, 3}}, // j 1
				{{14, 3}, {6, 3}, {2, 3}},   // j 2
			},
			{
				// i 1
				{{45, 0}, {39, 0}, {37, 3}}, // j 0
				{{35, 3}, {25, 0}, {23, 0}}, // j 1
				{{27, 3}, {17, 0}, {11, 3}}, // j 2
			},
			{
				// i 2
				{{63, 0}, {59, 3}, {57, 3}}, // j 0
				{{56, 3}, {45, 0}, {39, 0}}, // j 1
				{{46, 3}, {35, 3}, {25, 0}}, // j 2
			}},
		{ // face 7
			{
				// i 0
				{{36, 0}, {20, 0}, {14, 3}}, // j 0
				{{34, 0}, {19, 3}, {9, 3}},  // j 1
				{{38, 3}, {21, 3}, {7, 3}},  // j 2
			},
			{
				// i 1
				{{55, 0}, {40, 0}, {27, 3}}, // j 0
				{{54, 3}, {36, 0}, {20, 0}}, // j 1
				{{51, 3}, {34, 0}, {19, 3}}, // j 2
			},
			{
				// i 2
				{{72, 0}, {60, 3}, {46, 3}}, // j 0
				{{73, 3}, {55, 0}, {40, 0}}, // j 1
				{{71, 3}, {54, 3}, {36, 0}}, // j 2
			}},
		{ // face 8
			{
				// i 0
				{{64, 0}, {47, 0}, {38, 3}}, // j 0
				{{62, 0}, {43, 3}, {29, 3}}, // j 1
				{{58, 3}, {42, 3}, {26, 3}}, // j 2
			},
			{
				// i 1
				{{84, 0}, {69, 0}, {51, 3}}, // j 0
				{{82, 3}, {64, 0}, {47, 0}}, // j 1
				{{76, 3}, {62, 0}, {43, 3}}, // j 2
			},
			{
				// i 2
				{{97, 0}, {89, 3}, {71, 3}}, // j 0
				{{98, 3}, {84, 0}, {69, 0}}, // j 1
				{{96, 3}, {82, 3}, {64, 0}}, // j 2
			}},
		{ // face 9
			{
				// i 0
				{{75, 0}, {65, 0}, {58, 3}}, // j 0
				{{61, 0}, {53, 3}, {44, 3}}, // j 1
				{{49, 3}, {41, 3}, {31, 3}}, // j 2
			},
			{
				// i 1
				{{94, 0}, {86, 0}, {76, 3}}, // j 0
				{{81, 3}, {75, 0}, {65, 0}}, // j 1
				{{66, 3}, {61, 0}, {53, 3}}, // j 2
			},
			{
				// i 2
				{{107, 0}, {104, 3}, {96, 3}}, // j 0
				{{101, 3}, {94, 0}, {86, 0}},  // j 1
				{{85, 3}, {81, 3}, {75, 0}},   // j 2
			}},
		{ // face 10
			{
				// i 0
				{{57, 0}, {59, 0}, {63, 3}}, // j 0
				{{74, 0}, {78, 3}, {79, 3}}, // j 1
				{{83, 3}, {92, 3}, {95, 3}}, // j 2
			},
			{
				// i 1
				{{37, 0}, {39, 3}, {45, 3}}, // j 0
				{{52, 0}, {57, 0}, {59, 0}}, // j 1
				{{70, 3}, {74, 0}, {78, 3}}, // j 2
			},
			{
				// i 2
				{{24, 0}, {23, 3}, {25, 3}}, // j 0
				{{32, 3}, {37, 0}, {39, 3}}, // j 1
				{{50, 3}, {52, 0}, {57, 0}}, // j 2
			}},
		{ // face 11
			{
				// i 0
				{{46, 0}, {60, 0}, {72, 3}}, // j 0
				{{56, 0}, {68, 3}, {80, 3}}, // j 1
				{{63, 3}, {77, 3}, {90, 3}}, // j 2
			},
			{
				// i 1
				{{27, 0}, {40, 3}, {55, 3}}, // j 0
				{{35, 0}, {46, 0}, {60, 0}}, // j 1
				{{45, 3}, {56, 0}, {68, 3}}, // j 2
			},
			{
				// i 2
				{{14, 0}, {20, 3}, {36, 3}}, // j 0
				{{17, 3}, {27, 0}, {40, 3}}, // j 1
				{{25, 3}, {35, 0}, {46, 0}}, // j 2
			}},
		{ // face 12
			{
				// i 0
				{{71, 0}, {89, 0}, {97, 3}},  // j 0
				{{73, 0}, {91, 3}, {103, 3}}, // j 1
				{{72, 3}, {88, 3}, {105, 3}}, // j 2
			},
			{
				// i 1
				{{51, 0}, {69, 3}, {84, 3}}, // j 0
				{{54, 0}, {71, 0}, {89, 0}}, // j 1
				{{55, 3}, {73, 0}, {91, 3}}, // j 2
			},
			{
				// i 2
				{{38, 0}, {47, 3}, {64, 3}}, // j 0
				{{34, 3}, {51, 0}, {69, 3}}, // j 1
				{{36, 3}, {54, 0}, {71, 0}}, // j 2
			}},
		{ // face 13
			{
				// i 0
				{{96, 0}, {104, 0}, {107, 3}}, // j 0
				{{98, 0}, {110, 3}, {115, 3}}, // j 1
				{{97, 3}, {111, 3}, {119, 3}}, // j 2
			},
			{
				// i 1
				{{76, 0}, {86, 3}, {94, 3}},  // j 0
				{{82, 0}, {96, 0}, {104, 0}}, // j 1
				{{84, 3}, {98, 0}, {110, 3}}, // j 2
			},
			{
				// i 2
				{{58, 0}, {65, 3}, {75, 3}}, // j 0
				{{62, 3}, {76, 0}, {86, 3}}, // j 1
				{{64, 3}, {82, 0}, {96, 0}}, // j 2
			}},
		{ // face 14
			{
				// i 0
				{{85, 0}, {87, 0}, {83, 3}},    // j 0
				{{101, 0}, {102, 3}, {100, 3}}, // j 1
				{{107, 3}, {112, 3}, {114, 3}}, // j 2
			},
			{
				// i 1
				{{66, 0}, {67, 3}, {70, 3}},   // j 0
				{{81, 0}, {85, 0}, {87, 0}},   // j 1
				{{94, 3}, {101, 0}, {102, 3}}, // j 2
			},
			{
				// i 2
				{{49, 0}, {48, 3}, {50, 3}}, // j 0
				{{61, 3}, {66, 0}, {67, 3}}, // j 1
				{{75, 3}, {81, 0}, {85, 0}}, // j 2
			}},
		{ // face 15
			{
				// i 0
				{{95, 0}, {92, 0}, {83, 0}}, // j 0
				{{79, 0}, {78, 0}, {74, 3}}, // j 1
				{{63, 1}, {59, 3}, {57, 3}}, // j 2
			},
			{
				// i 1
				{{109, 0}, {108, 0}, {100, 5}}, // j 0
				{{93, 1}, {95, 0}, {92, 0}},    // j 1
				{{77, 1}, {79, 0}, {78, 0}},    // j 2
			},
			{
				// i 2
				{{117, 4}, {118, 5}, {114, 5}}, // j 0
				{{106, 1}, {109, 0}, {108, 0}}, // j 1
				{{90, 1}, {93, 1}, {95, 0}},    // j 2
			}},
		{ // face 16
			{
				// i 0
				{{90, 0}, {77, 0}, {63, 0}}, // j 0
				{{80, 0}, {68, 0}, {56, 3}}, // j 1
				{{72, 1}, {60, 3}, {46, 3}}, // j 2
			},
			{
				// i 1
				{{106, 0}, {93, 0}, {79, 5}}, // j 0
				{{99, 1}, {90, 0}, {77, 0}},  // j 1
				{{88, 1}, {80, 0}, {68, 0}},  // j 2
			},
			{
				// i 2
				{{117, 3}, {109, 5}, {95, 5}}, // j 0
				{{113, 1}, {106, 0}, {93, 0}}, // j 1
				{{105, 1}, {99, 1}, {90, 0}},  // j 2
			}},
		{ // face 17
			{
				// i 0
				{{105, 0}, {88, 0}, {72, 0}}, // j 0
				{{103, 0}, {91, 0}, {73, 3}}, // j 1
				{{97, 1}, {89, 3}, {71, 3}},  // j 2
			},
			{
				// i 1
				{{113, 0}, {99, 0}, {80, 5}},  // j 0
				{{116, 1}, {105, 0}, {88, 0}}, // j 1
				{{111, 1}, {103, 0}, {91, 0}}, // j 2
			},
			{
				// i 2
				{{117, 2}, {106, 5}, {90, 5}},  // j 0
				{{121, 1}, {113, 0}, {99, 0}},  // j 1
				{{119, 1}, {116, 1}, {105, 0}}, // j 2
			}},
		{ // face 18
			{
				// i 0
				{{119, 0}, {111, 0}, {97, 0}}, // j 0
				{{115, 0}, {110, 0}, {98, 3}}, // j 1
				{{107, 1}, {104, 3}, {96, 3}}, // j 2
			},
			{
				// i 1
				{{121, 0}, {116, 0}, {103, 5}}, // j 0
				{{120, 1}, {119, 0}, {111, 0}}, // j 1
				{{112, 1}, {115, 0}, {110, 0}}, // j 2
			},
			{
				// i 2
				{{117, 1}, {113, 5}, {105, 5}}, // j 0
				{{118, 1}, {121, 0}, {116, 0}}, // j 1
				{{114, 1}, {120, 1}, {119, 0}}, // j 2
			}},
		{ // face 19
			{
				// i 0
				{{114, 0}, {112, 0}, {107, 0}}, // j 0
				{{100, 0}, {102, 0}, {101, 3}}, // j 1
				{{83, 1}, {87, 3}, {85, 3}},    // j 2
			},
			{
				// i 1
				{{118, 0}, {120, 0}, {115, 5}}, // j 0
				{{108, 1}, {114, 0}, {112, 0}}, // j 1
				{{92, 1}, {100, 0}, {102, 0}},  // j 2
			},
			{
				// i 2
				{{117, 0}, {121, 5}, {119, 5}}, // j 0
				{{109, 1}, {118, 0}, {120, 0}}, // j 1
				{{95, 1}, {108, 1}, {114, 0}},  // j 2
			}},
	}
)

// isPentagon returns whether the indicated base cell is a pentagon.
func (c baseCell) isPentagon() bool {
	if c < 0 || c >= NUM_BASE_CELLS {
		return false
	}
	return baseCellData[c].isPentagon
}

// isPolarPentagon returns whether the indicated base cell is a pentagon where
// all neighbors are oriented towards it.
func (c baseCell) isPolarPentagon() bool {
	return c == 4 || c == 117
}

// baseCellDirection returns the direction from the origin base cell to the
// neighbor base cell. Returns INVALID_DIGIT if the cells are not neighbors.
func (c baseCell) baseCellDirection(neighbor baseCell) Direction {
	for dir := CENTER_DIGIT; dir < NUM_DIGITS; dir++ {
		testBaseCell := c.getBaseCellNeighbor(dir)
		if testBaseCell == neighbor {
			return dir
		}
	}
	return INVALID_DIGIT
}

// getBaseCellNeighbor returns the neighboring base cell in the given direction.
func (c baseCell) getBaseCellNeighbor(dir Direction) baseCell {
	return baseCellNeighbors[c][dir]
}

// faceIjkToBaseCellCCWrot60 finds the base cell given FaceIJK.
//
// Given the face number and a resolution 0 ijk+ coordinate in that face's
// face-centered ijk coordinate system, return the number of 60' ccw rotations
// to rotate into the coordinate system of the base cell at that coordinates.
//
// Valid ijk+ lookup coordinates are from (0, 0, 0) to (2, 2, 2).
func faceIjkToBaseCellCCWrot60(bc faceIJK) int {
	return faceIjkBaseCells[bc.face][bc.coord.i][bc.coord.j][bc.coord.k].ccwRot60
}

// faceIjkToBaseCell finds the base cell given FaceIJK.
//
// Given the face number and a resolution 0 ijk+ coordinate in that face's
// face-centered ijk coordinate system, return the base cell located at that
// coordinate.
//
// Valid ijk+ lookup coordinates are from (0, 0, 0) to (2, 2, 2).
func faceIjkToBaseCell(fijk faceIJK) baseCell {
	return faceIjkBaseCells[fijk.face][fijk.coord.i][fijk.coord.j][fijk.coord.k].bc
}

// baseCellToFaceIJK finds the FaceIJK given base cell.
func baseCellToFaceIjk(bc baseCell) faceIJK {
	return baseCellData[bc].homeFijk
}

// Given a base cell and the face it appears on, return the number of 60' ccw
// rotations for that base cell's coordinate system.
//
// Returns the number of rotations, or INVALID_ROTATIONS if the base cell is not
// found on the given face
func baseCellToCCWrot60(bc baseCell, face int) int {
	if face < 0 || face > NUM_ICOSA_FACES {
		return INVALID_ROTATIONS
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if faceIjkBaseCells[face][i][j][k].bc == bc {
					return faceIjkBaseCells[face][i][j][k].ccwRot60
				}
			}
		}
	}

	return INVALID_ROTATIONS
}

// getRes0Cells returns the number of resolution 0 cells.
func getRes0CellCount() int {
	return NUM_BASE_CELLS
}

// getRes0Cells returns a list of all resolution 0 cells.
func getRes0Cells() []Cell {
	out := make([]Cell, NUM_BASE_CELLS)
	for i := baseCell(0); i < NUM_BASE_CELLS; i++ {
		out[i] = H3_INIT.setMode(H3_CELL_MODE).setBaseCell(i)
	}
	return out
}
