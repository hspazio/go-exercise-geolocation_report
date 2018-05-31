package geopoint

import "math"

// Point is the interface for any geopoints
type Point interface {
	ToRadians() Radians
	ToDegrees() Degrees
}

// Degrees is the point in degrees
type Degrees struct {
	Latitude  float64
	Longitude float64
}

// ToRadians converts a Degrees point to Radians
func (p Degrees) ToRadians() Radians {
	return Radians{
		degreesToRadians(p.Latitude),
		degreesToRadians(p.Longitude),
	}
}

// ToDegrees converts a Radians point to Degrees
func (p Degrees) ToDegrees() Degrees {
	return p
}

// Radians is the point in radians
type Radians struct {
	Latitude  float64
	Longitude float64
}

// ToRadians converts a Degrees point to Radians
func (p Radians) ToRadians() Radians {
	return p
}

// ToDegrees converts a Radians point to Degrees
func (p Radians) ToDegrees() Degrees {
	return Degrees{
		radiansToDegrees(p.Latitude),
		radiansToDegrees(p.Longitude),
	}
}

// DistanceInKm calculates the distance between points
func DistanceInKm(p1, p2 Point) float64 {
	h := haversine(p1.ToRadians(), p2.ToRadians())
	hearthRadiusInMeters := 6371e3
	return hearthRadiusInMeters * h / 1000
}

func haversine(p1, p2 Radians) float64 {
	diffLatitude := p1.Latitude - p2.Latitude
	diffLongitude := p1.Longitude - p2.Longitude

	a := math.Sin(diffLatitude/2)*math.Sin(diffLatitude/2) +
		math.Cos(p1.Latitude)*
			math.Cos(p2.Latitude)*
			math.Sin(diffLongitude/2)*math.Sin(diffLongitude/2)

	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func degreesToRadians(value float64) float64 {
	return value * math.Pi / 180
}

func radiansToDegrees(value float64) float64 {
	return value * 180 / math.Pi
}
