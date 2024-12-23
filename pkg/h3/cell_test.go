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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
