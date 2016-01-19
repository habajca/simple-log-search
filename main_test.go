package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeoSet(t *testing.T) {
	testGeo := geoPoint{}
	err := (&testGeo).Set("37.7576171,-122.5776844")
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	assert.Equal(
		t,
		geoPoint{
			Latitude:  37.7576171,
			Longitude: -122.5776844,
		},
		testGeo,
	)

	err = (&testGeo).Set("")
	assert.NotNil(t, err)

	err = (&testGeo).Set("37.7576171")
	assert.NotNil(t, err)

	err = (&testGeo).Set("word,-122.5776844")
	assert.NotNil(t, err)

	err = (&testGeo).Set("37.7576171,word")
	assert.NotNil(t, err)
}

func TestValid(t *testing.T) {
	test := func(f func()) {
		resetValid()
		f()
		assert.False(t, valid())
	}
	negate := func(val *int) func() {
		return func() {
			*val = -1
		}
	}
	for _, i := range []*int{&fileCount, &rowCount, &timeFrame, &uidCount, &distance} {
		test(negate(i))
	}
	test(func() {
		timeOrigin = -1
	})
	test(func() {
		geo = geoPoint{Latitude: 91, Longitude: 0}
	})
	test(func() {
		geo = geoPoint{Latitude: -91, Longitude: 0}
	})
	test(func() {
		geo = geoPoint{Latitude: 0, Longitude: 181}
	})
	test(func() {
		geo = geoPoint{Latitude: 0, Longitude: -181}
	})
	resetValid()
	assert.True(t, valid())
}

func resetValid() {
	fileCount = 0
	rowCount = 0
	timeOrigin = 0
	timeFrame = 0
	uidCount = 0
	distance = 0
	geo = geoPoint{}
}
