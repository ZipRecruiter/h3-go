package h3

import "math"

const (
	// EPSILON_DEG is an epsilon of ~0.1mm in degrees
	EPSILON_DEG = .000000001
	// EPSILON_RAD is an epsilon of ~0.1mm in radians
	EPSILON_RAD = EPSILON_DEG * M_PI_180

	// NORMALIZE_NONE means do not normalize
	NORMALIZE_NONE = longitudeNormalization(0)
	// NORMALIZE_EAST means normalize negative numbers to the east
	NORMALIZE_EAST = longitudeNormalization(1)
	// NORMALIZE_WEST means normalize positive numbers to the west
	NORMALIZE_WEST = longitudeNormalization(2)
)

type longitudeNormalization int

type LatLng [2]float64

func (l LatLng) Latitude() float64 {
	return l[0]
}

func (l LatLng) Longitude() float64 {
	return l[1]
}

// NewLatLng creates a new LatLng from the given latitude and longitude in degrees.
func NewLatLng(lat float64, lng float64) LatLng {
	return LatLng{
		deg2rad(lat),
		deg2rad(lng),
	}
}

// deg2rad converts degrees to radians.
func deg2rad(deg float64) float64 {
	return (deg * math.Pi) / 180.0
}

// rad2deg converts radians to degrees.
func rad2deg(rad float64) float64 {
	return (rad * 180.0) / math.Pi
}

// toFaceIJK returns the corresponding icosahedral face and containing 2D hex
// coordinates relative to that face center.
func (l LatLng) toFaceIJK(res int) faceIJK {
	return geoToFaceIJK(l, res)
}

// posAngleRads normalizes radians to a value between 0.0 and two*PI.
func posAngleRads(rads float64) float64 {
	tmp := rads
	if rads < 0.0 {
		tmp = rads + M_2PI
	}

	if rads >= M_2PI {
		tmp -= M_2PI
	}

	return tmp
}

// geoAzimuthRads calculates the azimuth to other from this point in radians.
func (l LatLng) geoAzimuthRads(other LatLng) float64 {
	return math.Atan2(math.Cos(other.Latitude())*math.Sin(other.Longitude()-l.Longitude()),
		math.Cos(l.Latitude())*math.Sin(other.Latitude())-
			math.Sin(l.Latitude())*math.Cos(other.Latitude())*math.Cos(other.Longitude()-l.Longitude()))
}

// geoAzimuthDistanceRads computes the point on the sphere a specified azimuth and distance from this point.
func (l LatLng) geoAzimuthDistanceRads(azimuth float64, distance float64) LatLng {
	if distance < EPSILON {
		return l
	}

	var sinlat, sinlng, coslng float64

	az := posAngleRads(azimuth)

	p2 := LatLng{}

	// check for due north/south azimuth
	if az < EPSILON || math.Abs(az-math.Pi) < EPSILON {
		if az < EPSILON {
			// due north
			p2[0] = l[0] + distance
		} else {
			// due south
			p2[0] = l[0] - distance
		}

		if math.Abs(p2[0]-M_PI_2) < EPSILON {
			// north pole
			p2[0] = M_PI_2
			p2[1] = 0.0
		} else if math.Abs(p2[0]+M_PI_2) < EPSILON {
			// south pole
			p2[0] = -M_PI_2
			p2[1] = 0.0
		} else {
			p2[1] = constrainLng(l[1])
		}
	} else {
		// not due north or south
		sinlat = math.Sin(l[0])*math.Cos(distance) + math.Cos(l[0])*math.Sin(distance)*math.Cos(az)
		if sinlat > 1.0 {
			sinlat = 1.0
		}
		if sinlat < -1.0 {
			sinlat = -1.0
		}
		p2[0] = math.Asin(sinlat)
		if math.Abs(p2[0]-M_PI_2) < EPSILON {
			// north pole
			p2[0] = M_PI_2
			p2[1] = 0.0
		} else if math.Abs(p2[0]+M_PI_2) < EPSILON {
			// south pole
			p2[0] = -M_PI_2
			p2[1] = 0.0
		} else {
			invcosp2lat := 1.0 / math.Cos(p2[0])
			sinlng = math.Sin(az) * math.Sin(distance) * invcosp2lat
			coslng = (math.Cos(distance) - math.Sin(l[0])*math.Sin(p2[0])) /
				math.Cos(l[0]) * invcosp2lat
			if sinlng > 1.0 {
				sinlng = 1.0
			}
			if sinlng < -1.0 {
				sinlng = -1.0
			}
			if coslng > 1.0 {
				coslng = 1.0
			}
			if coslng < -1.0 {
				coslng = -1.0
			}
			p2[1] = constrainLng(l[1] + math.Atan2(sinlng, coslng))
		}
	}

	return p2
}

// greatCircleDistanceRads calculates the great-circle distance between two
// lat/lng points (in radians). Returns the great circle distance between the two
// points in radians.
func (l LatLng) greatCircleDistanceRads(other LatLng) float64 {
	sinLat := math.Sin((other.Latitude() - l.Latitude()) * 0.5)
	sinLng := math.Sin((other.Longitude() - l.Longitude()) * 0.5)

	a := sinLat*sinLat + math.Cos(l.Latitude())*math.Cos(other.Latitude())*sinLng*sinLng

	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

// GetNumCells returns the number of cells (hexagons) at the given resolution.
func GetNumCells(res int) (uint64, error) {
	if res < 0 || res > MAX_H3_RES {
		return 0, ErrInvalidArgument
	}

	return uint64(2 + 120*ipow(7, res)), nil
}

// contrainLat makes sure latitudes are in the proper bounds
func constrainLat(lat float64) float64 {
	for lat > M_PI_2 {
		lat = lat - math.Pi
	}
	return lat
}

// constrainLng makes sure longitudes are in the proper bounds
func constrainLng(lng float64) float64 {
	for lng > math.Pi {
		lng = lng - (2 * math.Pi)
	}

	for lng < -math.Pi {
		lng = lng + (2 * math.Pi)
	}

	return lng
}

// normalizeLng normalizes an input longitude according to the specified
// normalization
func normalizeLng(lng float64, normalization longitudeNormalization) float64 {
	switch normalization {
	case NORMALIZE_EAST:
		if lng < 0 {
			return lng + M_2PI
		} else {
			return lng
		}
	case NORMALIZE_WEST:
		if lng > 0 {
			return lng - M_2PI
		} else {
			return lng
		}
	}

	return lng
}
