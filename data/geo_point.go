package data

import (
	"github.com/kellydunn/golang-geo"
)

type GeoPoint struct {
	Latitude  float64
	Longitude float64
}

type GeoPointBase interface {
	Lat() float64
	Lng() float64
}

func NewGeoPoint(base GeoPointBase) GeoPoint {
	return GeoPoint{
		Latitude:  base.Lat(),
		Longitude: base.Lng(),
	}
}

func (geoPoint GeoPoint) Point() *geo.Point {
	return geo.NewPoint(geoPoint.Latitude, geoPoint.Longitude)
}
