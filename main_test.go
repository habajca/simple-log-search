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
			latitude:  37.7576171,
			longitude: -122.5776844,
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
