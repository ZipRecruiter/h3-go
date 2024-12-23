package h3

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
