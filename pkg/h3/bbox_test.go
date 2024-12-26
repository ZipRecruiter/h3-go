package h3

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bboxNormalization(t *testing.T) {
	type args struct {
		a bbox
		b bbox
	}
	tests := []struct {
		name  string
		args  args
		want  longitudeNormalization
		want1 longitudeNormalization
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := bboxNormalization(tt.args.a, tt.args.b)
			assert.Equalf(t, tt.want, got, "bboxNormalization(%v, %v)", tt.args.a, tt.args.b)
			assert.Equalf(t, tt.want1, got1, "bboxNormalization(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

func Test_bbox_center(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   LatLng
	}{
		{
			name:   "quadrant pos/pos",
			fields: fields{north: 1, south: 0.8, east: 1, west: 0.8},
			want:   LatLng{0.9, 0.9},
		},
		{
			name:   "quadrant neg/pos",
			fields: fields{north: -0.8, south: -1.0, east: 1, west: 0.8},
			want:   LatLng{-0.9, 0.9},
		},
		{
			name:   "quadrant pos/neg",
			fields: fields{north: 1, south: 0.8, east: -0.8, west: -1.0},
			want:   LatLng{0.9, -0.9},
		},
		{
			name:   "quadrant neg/neg",
			fields: fields{north: -0.8, south: -1.0, east: -0.8, west: -1.0},
			want:   LatLng{-0.9, -0.9},
		},
		{
			name:   "transmeridian - skew west",
			fields: fields{north: 1, south: 0.8, east: -math.Pi + 0.3, west: math.Pi - 0.1},
			want:   LatLng{0.9, -math.Pi + 0.1},
		},
		{
			name:   "transmeridian - skew east",
			fields: fields{north: 1, south: 0.8, east: -math.Pi + 0.1, west: math.Pi - 0.3},
			want:   LatLng{0.9, math.Pi - 0.1},
		},
		{
			name:   "transmeridian - on antimeridian",
			fields: fields{north: 1, south: 0.8, east: -math.Pi + 0.1, west: math.Pi - 0.1},
			want:   LatLng{0.9, math.Pi},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			actualCenter := b.center()
			assert.InDelta(t, tt.want.Longitude(), actualCenter.Longitude(), EPSILON_RAD, "center() longitude not within epsilon")
			assert.InDelta(t, tt.want.Latitude(), actualCenter.Latitude(), EPSILON_RAD, "center() latitude not within epsilon")
		})
	}
}

func Test_bbox_containsBbox(t *testing.T) {
	t.Run("contains edge points", func(t *testing.T) {
		b := bbox{north: 0.1, south: -0.1, east: 0.2, west: -0.2}
		points := []LatLng{
			{0.1, 0.2}, {0.1, 0.0}, {0.1, -0.2}, {0.0, 0.2},
			{-0.1, 0.2}, {-0.1, 0.0}, {-0.1, -0.2}, {0.0, -0.2},
		}
		for _, p := range points {
			assert.Truef(t, b.containsPoint(p), "bbox should contain its edge point (%v)", p)
		}
	})

	t.Run("contains edge points across transmeridian", func(t *testing.T) {
		b := bbox{north: 0.1, south: -0.1, east: -math.Pi + 0.2, west: math.Pi - 0.2}
		points := []LatLng{
			{0.1, -math.Pi + 0.2}, {0.1, math.Pi}, {0.1, math.Pi - 0.2},
			{0.0, -math.Pi + 0.2}, {-0.1, -math.Pi + 0.2}, {-0.1, math.Pi},
			{-0.1, math.Pi - 0.2}, {0.0, math.Pi - 0.2},
		}
		for _, p := range points {
			assert.Truef(t, b.containsPoint(p), "bbox should contain its edge point (%v)", p)
		}
	})

	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	type args struct {
		o bbox
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.containsBbox(tt.args.o), "containsBbox(%v)", tt.args.o)
		})
	}
}

func Test_bbox_containsPoint(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	type args struct {
		p LatLng
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.containsPoint(tt.args.p), "containsPoint(%v)", tt.args.p)
		})
	}
}

func Test_bbox_equals(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	type args struct {
		o bbox
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "equals self",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   true,
		},
		{
			name:   "not equals different north",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1.1, south: 0, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "not equals different south",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0.1, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "not equals different east",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1.1, west: 0}},
			want:   false,
		},
		{
			name:   "not equals different west",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0.1}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.equals(tt.args.o), "equals(%v)", tt.args.o)
		})
	}
}

func Test_bbox_heightRads(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.heightRads(), "heightRads()")
		})
	}
}

func Test_bbox_isTransmeridian(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "normal box not transmeridian",
			fields: fields{north: 1, south: 0.8, east: 1, west: 0.8},
			want:   false,
		},
		{
			name:   "transmeridian box",
			fields: fields{north: 1, south: 0.8, east: -math.Pi + 0.3, west: math.Pi - 0.1},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.isTransmeridian(), "isTransmeridian()")
		})
	}
}

func Test_bbox_overlaps(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	type args struct {
		o bbox
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "no intersection to the west",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: -1, west: -1.5}},
			want:   false,
		},
		{
			name:   "no intersection to the west, reverse",
			fields: fields{north: 1, south: 0, east: -1, west: -1.5},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "no intersection to the east",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 2, west: 1.5}},
			want:   false,
		},
		{
			name:   "no intersection to the east, reverse",
			fields: fields{north: 1, south: 0, east: 2, west: 1.5},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "no intersection to the south",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: -1, south: -1.5, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "no intersection to the south, reverse",
			fields: fields{north: -1, south: -1.5, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "no intersection to the north",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 2, south: 1.5, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "no intersection to the north, reverse",
			fields: fields{north: 2, south: 1.5, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   false,
		},
		{
			name:   "intersection to the west",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 0.5, west: -1.5}},
			want:   true,
		},
		{
			name:   "intersection to the east",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 2, west: 0.5}},
			want:   true,
		},
		{
			name:   "intersection to the south",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 0.5, south: -1.5, east: 1, west: 0}},
			want:   true,
		},
		{
			name:   "intersection to the north",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 2, south: 0.5, east: 1, west: 0}},
			want:   true,
		},
		{
			name:   "intersection - b contains a",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1.5, south: -0.5, east: 1.5, west: -0.5}},
			want:   true,
		},
		{
			name:   "intersection - a contains b",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 0.5, south: 0.25, east: 0.5, west: 0.25}},
			want:   true,
		},
		{
			name:   "intersection - a equals b",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{o: bbox{north: 1, south: 0, east: 1, west: 0}},
			want:   true,
		},
		{
			name:   "intersection, transmeridian to the east",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.9, west: math.Pi - 0.4}},
			want:   true,
		},
		{
			name:   "intersection, transmeridian to the east, reverse",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.9, west: math.Pi - 0.4},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   true,
		},
		{
			name:   "intersection, transmeridian to the west",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.4, west: math.Pi - 0.9}},
			want:   true,
		},
		{
			name:   "intersection, transmeridian to the west, reverse",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.4, west: math.Pi - 0.9},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   true,
		},
		{
			name:   "transmeridian, no intersection to the west",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: math.Pi - 0.7, west: math.Pi - 0.9}},
			want:   false,
		},
		{
			name:   "transmeridian, no intersection to the west, reverse",
			fields: fields{north: 1, south: 0, east: math.Pi - 0.7, west: math.Pi - 0.9},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   false,
		},
		{
			name:   "transmeridian, no intersection to the east",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.9, west: -math.Pi + 0.7}},
			want:   false,
		},
		{
			name:   "transmeridian, no intersection to the east, reverse",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.9, west: -math.Pi + 0.7},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   false,
		},
		{
			name:   "transmeridian, intersection to the west",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: math.Pi - 0.4, west: math.Pi - 0.9}},
			want:   true,
		},
		{
			name:   "transmeridian, intersection to the west, reverse",
			fields: fields{north: 1, south: 0, east: math.Pi - 0.4, west: math.Pi - 0.9},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   true,
		},
		{
			name:   "transmeridian, intersection to the east",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.9, west: -math.Pi + 0.4}},
			want:   true,
		},
		{
			name:   "transmeridian, intersection to the east, reverse",
			fields: fields{north: 1, south: 0, east: -math.Pi + 0.9, west: -math.Pi + 0.4},
			args:   args{o: bbox{north: 1, south: 0, east: -math.Pi + 0.5, west: math.Pi - 0.5}},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.overlaps(tt.args.o), "overlaps(%v)", tt.args.o)
		})
	}
}

func Test_bbox_scale(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	type args struct {
		factor float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bbox
	}{
		{
			name:   "noop",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{factor: 1},
			want:   bbox{north: 1, south: 0, east: 1, west: 0},
		},
		{
			name:   "basic grow",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{factor: 2},
			want:   bbox{north: 1.5, south: -0.5, east: 1.5, west: -0.5},
		},
		{
			name:   "basic shrink",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			args:   args{factor: 0.5},
			want:   bbox{north: 0.75, south: 0.25, east: 0.75, west: 0.25},
		},
		{
			name:   "clamp north south",
			fields: fields{north: M_PI_2 * 0.9, south: -M_PI_2 * 0.9, east: 1.0, west: 0.0},
			args:   args{factor: 2},
			want:   bbox{north: M_PI_2, south: -M_PI_2, east: 1.5, west: -0.5},
		},
		{
			name:   "clamp east positive",
			fields: fields{north: 1.0, south: 0.0, east: math.Pi - 0.1, west: math.Pi - 1.1},
			args:   args{factor: 2},
			want:   bbox{north: 1.5, south: -0.5, east: -math.Pi + 0.4, west: math.Pi - 1.6},
		},
		{
			name:   "clamp east negative",
			fields: fields{north: 1.5, south: -0.5, east: -math.Pi + 0.4, west: math.Pi - 1.6},
			args:   args{factor: 2},
			want:   bbox{north: 1.0, south: 0.0, east: math.Pi - 0.1, west: math.Pi - 1.1},
		},
		{
			name:   "clamp west positive",
			fields: fields{north: 1.0, south: 0.0, east: -math.Pi + 0.9, west: math.Pi - 0.1},
			args:   args{factor: 0.5},
			want:   bbox{north: 0.75, south: 0.25, east: -math.Pi + 0.65, west: -math.Pi + 0.15},
		},
		{
			name:   "clamp west negative",
			fields: fields{north: 0.75, south: 0.25, east: -math.Pi + 0.65, west: -math.Pi + 0.15},
			args:   args{factor: 2},
			want:   bbox{north: 1.0, south: 0.0, east: -math.Pi + 0.9, west: math.Pi - 0.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.scale(tt.args.factor), "scale(%v)", tt.args.factor)
		})
	}
}

func Test_bbox_widthRads(t *testing.T) {
	type fields struct {
		north float64
		south float64
		east  float64
		west  float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "normal",
			fields: fields{north: 1, south: 0, east: 1, west: 0},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bbox{north: tt.fields.north, south: tt.fields.south,
				east: tt.fields.east, west: tt.fields.west,
			}
			assert.Equalf(t, tt.want, b.widthRads(), "widthRads()")
		})
	}
}
