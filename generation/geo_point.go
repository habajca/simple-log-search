package generation

import (
	"github.com/habajca/simple-log-search/data"
	"math/rand"
)

func randomGeoPoint(origin data.GeoPoint, distance int) data.GeoPoint {
	randDistance := rand.Float64() * float64(distance) / 1000
	randBearing := rand.Float64()
	point := origin.Point()
	point.PointAtDistanceAndBearing(randDistance, randBearing)
	return data.NewGeoPoint(point)
}
