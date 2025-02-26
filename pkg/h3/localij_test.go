package h3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell_toLocalIJK(t *testing.T) {
	type args struct {
		other Cell
	}
	tests := []struct {
		name    string
		c       Cell
		args    args
		want    coordIJK
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"identity",
			0x81283ffffffffff,
			args{0x81283ffffffffff},
			coordIJK{},
			assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.toLocalIJK(tt.args.other)
			if !tt.wantErr(t, err, fmt.Sprintf("toLocalIJK(%v)", tt.args.other)) {
				return
			}
			assert.Equalf(t, tt.want, got, "toLocalIJK(%v)", tt.args.other)
		})
	}
}
