package h3

import "math"

type bbox struct {
	north, south, east, west float64
}

// widthRads returns the width of the bounding box in radians.
func (b bbox) widthRads() float64 {
	if b.isTransmeridian() {
		return b.east - b.west + M_2PI
	} else {
		return b.east - b.west
	}
}

// heightRads returns the height of the bounding box in radians.
func (b bbox) heightRads() float64 {
	return b.north - b.south
}

// isTransmeridian returns true if the bounding box crosses the antimeridian.
func (b bbox) isTransmeridian() bool {
	return b.east < b.west
}

// center returns the center of the bounding box.
func (b bbox) center() LatLng {
	east := b.east
	if b.isTransmeridian() {
		east += M_2PI
	}

	return LatLng{
		(b.north + b.south) * 0.5,
		constrainLng((east + b.west) * 0.5),
	}
}

// overlaps returns true if the bounding box overlaps the given bounding box.
func (b bbox) overlaps(o bbox) bool {
	// check latitudes first
	if b.north < o.south || b.south > o.north {
		return false
	}

	// check for overlapping longitude, accounting for transmeridian case
	aNormalization, bNormalization := bboxNormalization(b, o)

	if normalizeLng(b.east, aNormalization) <
		normalizeLng(o.west, bNormalization) ||
		normalizeLng(b.west, aNormalization) >
			normalizeLng(o.east, bNormalization) {
		return false
	}

	return true
}

// contains returns true if the bounding box contains the given point.
func (b bbox) containsPoint(p LatLng) bool {
	if p.Latitude() < b.south || p.Latitude() > b.north {
		return false
	}

	if b.isTransmeridian() {
		// transmeridian case
		return p.Longitude() >= b.west || p.Longitude() <= b.east
	} else {
		// standard case
		return p.Longitude() >= b.west && p.Longitude() <= b.east
	}
}

// containsBbox returns true if the bounding box contains the given bounding box.
func (b bbox) containsBbox(o bbox) bool {
	// Check whether latitude coords are contained
	if b.north < o.north || b.south > o.south {
		return false
	}
	// Check whether longitude coords are contained
	// Account for transmeridian bboxes
	aNormalization, bNormalization := bboxNormalization(b, o)
	return normalizeLng(b.west, aNormalization) <=
		normalizeLng(o.west, bNormalization) &&
		normalizeLng(b.east, aNormalization) >=
			normalizeLng(o.east, bNormalization)
}

// equals returns true if the bounding box is equal to the given bounding box.
func (b bbox) equals(o bbox) bool {
	return b.north == o.north && b.south == o.south && b.east == o.east && b.west == o.west
}

// scale scales the bounding box by a given factor and returns the new bounding box.
func (b bbox) scale(factor float64) bbox {
	width := b.widthRads()
	height := b.heightRads()

	widthBuffer := (width*factor - width) * 0.5
	heightBuffer := (height*factor - height) * 0.5

	newBbox := bbox{
		north: b.north + heightBuffer,
		south: b.south - heightBuffer,
		east:  b.east + widthBuffer,
		west:  b.west - widthBuffer,
	}

	// scale north and south, clamping to latitude domain
	if newBbox.north > M_PI_2 {
		newBbox.north = M_PI_2
	}
	if newBbox.south < -M_PI_2 {
		newBbox.south = -M_PI_2
	}
	// scale east and west, clamping to longitude domain
	if newBbox.east > math.Pi {
		newBbox.east -= M_2PI
	}
	if newBbox.east < -math.Pi {
		newBbox.east += M_2PI
	}
	if newBbox.west > math.Pi {
		newBbox.west -= M_2PI
	}
	if newBbox.west < -math.Pi {
		newBbox.west += M_2PI
	}

	return newBbox
}

// bboxNormalization determines the longitude normalization scheme for two
// bounding boxes, either or both of which might cross the antimeridian. The goal
// is to transform latitudes in one or both boxes so that they are in the same
// frame of reference and can be operated on with standard Cartesian functions.
func bboxNormalization(a bbox, b bbox) (longitudeNormalization, longitudeNormalization) {
	aIsTransmeridian := a.isTransmeridian()
	bIsTransmeridian := b.isTransmeridian()
	aToBTrendsEast := a.west-b.east < b.west-a.east

	// If neither is transmeridian, no normalization.
	// If both are transmeridian, normalize east by convention.
	// If one is transmeridian and one is not, normalize toward the other.

	var aNormalization, bNormalization longitudeNormalization

	if !aIsTransmeridian {
		aNormalization = NORMALIZE_NONE
	} else if bIsTransmeridian {
		aNormalization = NORMALIZE_EAST
	} else if aToBTrendsEast {
		aNormalization = NORMALIZE_EAST
	} else {
		aNormalization = NORMALIZE_WEST
	}

	if !bIsTransmeridian {
		bNormalization = NORMALIZE_NONE
	} else if aIsTransmeridian {
		bNormalization = NORMALIZE_EAST
	} else if aToBTrendsEast {
		bNormalization = NORMALIZE_WEST
	} else {
		bNormalization = NORMALIZE_EAST
	}

	return aNormalization, bNormalization
}
