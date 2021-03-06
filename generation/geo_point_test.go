package generation

import (
	"fmt"
	"github.com/habajca/simple-log-search/util"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const TEST_RANDOM_GEO_RUNS = 1000
const TEST_LATITUDE = 37.7576171
const TEST_LONGITUDE = -122.5776844
const TEST_DISTANCE_M = 5000
const TEST_DISTNACE_KM = float64(TEST_DISTANCE_M) / 1000

// Note: This test involves random logic and therefore may be inconsistent. Reference the seed for repeatability.
func TestRandomGeo(t *testing.T) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	origin := util.GeoPoint{
		Latitude:  TEST_LATITUDE,
		Longitude: TEST_LONGITUDE,
	}
	for i := 0; i < TEST_RANDOM_GEO_RUNS; i++ {
		randGeoPoint := randomGeoPoint(origin, TEST_DISTANCE_M)
		distance := origin.Point().GreatCircleDistance(randGeoPoint.Point())
		if !assert.True(
			t,
			distance <= TEST_DISTNACE_KM,
			fmt.Sprintf(
				"The random point %v is %f kms away from the test origin. Seed: %d.",
				randGeoPoint,
				distance,
				seed,
			),
		) {
			t.FailNow()
		}
	}
	assert.NotEqual(
		t,
		randomGeoPoint(origin, TEST_DISTANCE_M),
		randomGeoPoint(origin, TEST_DISTANCE_M),
		fmt.Sprintf("Two random points should not be the same. Seed: %d.", seed),
	)
}
