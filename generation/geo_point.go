package generation

import (
	"github.com/habajca/simple-log-search/util"
	"math/rand"
)

func randomGeoPoint(origin util.GeoPoint, distance int) util.GeoPoint {
	randDistance := rand.Float64() * float64(distance) / 1000
	randBearing := rand.Float64()
	point := origin.Point()
	point.PointAtDistanceAndBearing(randDistance, randBearing)
	return util.NewGeoPoint(point)
}
