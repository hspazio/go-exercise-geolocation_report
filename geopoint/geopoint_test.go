package geopoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeopoint(t *testing.T) {
	assert := assert.New(t)

	t.Run("converts degrees to radians", func(t *testing.T) {
		point := Degrees{53.3393, -6.2576841}

		radiansPoint := point.ToRadians()

		assert.InDelta(0.9309464057, radiansPoint.Latitude, 1.0e-10)
		assert.InDelta(-0.109217191094, radiansPoint.Longitude, 1.0e-10)
	})

	t.Run("converts radians to degrees", func(t *testing.T) {
		point := Radians{0.9309464057, -0.109217191094}

		degreesPoint := point.ToDegrees()

		assert.InDelta(53.3393, degreesPoint.Latitude, 0.0001)
		assert.InDelta(-6.2576841, degreesPoint.Longitude, 0.0001)
	})

	t.Run("returns distance in km between Degrees points", func(t *testing.T) {
		point1 := Degrees{53.3393, -6.2576841}
		point2 := Degrees{51.8856167, -10.4240951}

		distance := DistanceInKm(point1, point2)

		assert.InDelta(324.36, distance, 0.01)
	})

	t.Run("returns distance in km between Radians points", func(t *testing.T) {
		point1 := Degrees{53.3393, -6.2576841}.ToRadians()
		point2 := Degrees{51.8856167, -10.4240951}.ToRadians()

		distance := DistanceInKm(point1, point2)

		assert.InDelta(324.36, distance, 0.01)
	})
}
