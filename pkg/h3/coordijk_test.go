package h3

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coordIJK_toDigit(t *testing.T) {
	type args struct {
		ijk coordIJK
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{
			name: "unit ijk to zero",
			args: args{ijk: coordIJK{0, 0, 0}},
			want: CENTER_DIGIT,
		},
		{
			name: "unit ijk to i axis",
			args: args{ijk: coordIJK{1, 0, 0}},
			want: I_AXES_DIGIT,
		},
		{
			name: "out of bounds",
			args: args{ijk: coordIJK{2, 0, 0}},
			want: INVALID_DIGIT,
		},
		{
			name: "unnormalize unit ijk to zero",
			args: args{ijk: coordIJK{2, 2, 2}},
			want: CENTER_DIGIT,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.args.ijk.toDigit(), "toDigit(%v)", tt.args.ijk)
		})
	}
}

func Test_coordIJK_neighbor(t *testing.T) {
	type fields struct {
		i int
		j int
		k int
	}
	type args struct {
		dir Direction
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   coordIJK
	}{
		{
			name:   "center neighbor is self",
			fields: fields{i: 0, j: 0, k: 0},
			args:   args{dir: CENTER_DIGIT},
			want:   coordIJK{0, 0, 0},
		},
		{
			name:   "i neighbor",
			fields: fields{i: 0, j: 0, k: 0},
			args:   args{dir: I_AXES_DIGIT},
			want:   coordIJK{1, 0, 0},
		},
		{
			name:   "invalid neighbor is self",
			fields: fields{i: 0, j: 0, k: 0},
			args:   args{dir: INVALID_DIGIT},
			want:   coordIJK{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := coordIJK{
				i: tt.fields.i,
				j: tt.fields.j,
				k: tt.fields.k,
			}
			assert.Equalf(t, tt.want, u.neighbor(tt.args.dir), "neighbor(%v)", tt.args.dir)
		})
	}
}

func Test_coordIJK_toIj(t *testing.T) {
	type fields struct {
		i int
		j int
		k int
	}
	tests := []struct {
		name   string
		fields fields
		want   coordIJ
	}{
		{
			name:   "zero",
			fields: fields{i: 0, j: 0, k: 0},
			want:   coordIJ{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coordIJK{
				i: tt.fields.i,
				j: tt.fields.j,
				k: tt.fields.k,
			}
			assert.Equalf(t, tt.want, c.toIj(), "toIj()")
		})
	}
}

func Test_coordIJ_toIjk(t *testing.T) {
	type fields struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		fields  fields
		want    coordIJK
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "zero",
			fields:  fields{i: 0, j: 0},
			want:    coordIJK{0, 0, 0},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := coordIJ{
				i: tt.fields.i,
				j: tt.fields.j,
			}
			got, err := c.toIjk()
			if !tt.wantErr(t, err, fmt.Sprintf("toIjk()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "toIjk()")
		})
	}
}

func Test_coordIJK_toIj_roundtrip(t *testing.T) {
	for d := CENTER_DIGIT; d < NUM_DIGITS; d++ {
		ijk := coordIJK{0, 0, 0}
		n := ijk.neighbor(d)

		ij := n.toIj()

		recoveredIJK, err := ij.toIjk()
		assert.NoError(t, err)
		assert.True(t, n.matches(recoveredIJK), "expected %v, got %v", ijk, recoveredIJK)
	}
}

func Test_coordIJK_toCube_roundtrip(t *testing.T) {
	for d := CENTER_DIGIT; d < NUM_DIGITS; d++ {
		ijk := coordIJK{0, 0, 0}
		n := ijk.neighbor(d)

		cubeIJK := n.toCube()
		recoveredIJK := NewCoordIJKFromCube(cubeIJK)

		assert.True(t, n.matches(recoveredIJK), "expected %v, got %v", ijk, recoveredIJK)
	}
}

func Test_coordIJK_upAp7Checked(t *testing.T) {
	t.Run("unit ijk", func(t *testing.T) {
		ijk := coordIJK{0, 0, 0}
		_, err := ijk.upAp7Checked()
		assert.NoError(t, err)
	})

	t.Run("i + i overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32, 0, 0}
		_, err := ijk.upAp7Checked()
		assert.Error(t, err)
	})

	t.Run("i * 3 overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32 / 2, 0, 0}
		_, err := ijk.upAp7Checked()
		assert.Error(t, err)
	})

	t.Run("j + j overflwos", func(t *testing.T) {
		ijk := coordIJK{0, math.MaxInt32, 0}
		_, err := ijk.upAp7Checked()
		assert.Error(t, err)
	})

	t.Run("(i * 3) - j overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32 / 3, -2, 0}
		_, err := ijk.upAp7Checked()
		assert.Error(t, err)
	})

	t.Run("i + (j * 2) overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32 / 3, math.MaxInt32 / 2, 0}
		_, err := ijk.upAp7Checked()
		assert.Error(t, err)
	})

	t.Run("i < 0 succeeds", func(t *testing.T) {
		ijk := coordIJK{-1, 0, 0}
		_, err := ijk.upAp7Checked()
		assert.NoError(t, err)
	})
}

func Test_coordIJK_upAp7rChecked(t *testing.T) {
	t.Run("unit ijk", func(t *testing.T) {
		ijk := coordIJK{0, 0, 0}
		_, err := ijk.upAp7rChecked()
		assert.NoError(t, err)
	})

	t.Run("i + i overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32, 0, 0}
		_, err := ijk.upAp7rChecked()
		assert.Error(t, err)
	})

	t.Run("j + j overflows", func(t *testing.T) {
		ijk := coordIJK{0, math.MaxInt32, 0}
		_, err := ijk.upAp7rChecked()
		assert.Error(t, err)
	})

	t.Run("3 * j overflwos", func(t *testing.T) {
		ijk := coordIJK{0, math.MaxInt32 / 2, 0}
		_, err := ijk.upAp7rChecked()
		assert.Error(t, err)
	})

	t.Run("(i * 2) + j overflows", func(t *testing.T) {
		ijk := coordIJK{math.MaxInt32 / 2, math.MaxInt32 / 3, 0}
		_, err := ijk.upAp7rChecked()
		assert.Error(t, err)
	})

	t.Run("(j * 3) - 1 overflows", func(t *testing.T) {
		ijk := coordIJK{-2, math.MaxInt32 / 3, 0}
		_, err := ijk.upAp7rChecked()
		assert.Error(t, err)
	})

	t.Run("i < 0 succeeds", func(t *testing.T) {
		ijk := coordIJK{-1, 0, 0}
		_, err := ijk.upAp7rChecked()
		assert.NoError(t, err)
	})
}
