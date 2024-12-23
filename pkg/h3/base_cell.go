package h3

type baseCell int

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
)

func (c baseCell) isPentagon() bool {
	if c < 0 || c >= NUM_BASE_CELLS {
		return false
	}
	return baseCellData[c].isPentagon
}
