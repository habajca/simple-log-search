package util

import (
	"encoding/json"
	"fmt"
	"github.com/kellydunn/golang-geo"
	"strconv"
)

type LogRow struct {
	Timestamp int64
	Uid       string
	Domain    string
	Geo       GeoPoint
}

func StructToString(o interface{}) string {
	bytes, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func StructFromString(s string, o interface{}) {
	err := json.Unmarshal([]byte(s), o)
	if err != nil {
		panic(fmt.Sprintf("%s\n%s", s, err))
	}
}

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

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
