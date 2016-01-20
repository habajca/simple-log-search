package generation

import (
	"github.com/habajca/simple-log-search/util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestFileGeneration(t *testing.T) {
	filenames, err := GenerateTestData(
		"tmpdata",
		1, 1,
		time.Now().Unix(), 3600,
		1,
		"testdata/domains.txt",
		util.GeoPoint{
			Latitude:  37.7576171,
			Longitude: -122.5776844,
		}, 5000,
	)
	if !assert.Nil(t, err) {
		t.FailNow()
	}
	if !assert.Len(t, filenames, 1) {
		t.FailNow()
	}
	_, err = os.Stat("tmpdata/" + filenames[0])
	if !assert.Nil(t, err) {
		t.FailNow()
	}
}
