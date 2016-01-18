package generation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const TEST_RANDOM_TIMESTAMP_RUNS = 1000
const TEST_TIME_DISTANCE = 1
const TEST_TIME_ORIGIN = 0
const TEST_TIME_RANGE_START = TEST_TIME_ORIGIN - TEST_TIME_DISTANCE
const TEST_TIME_RANGE_END = TEST_TIME_ORIGIN + TEST_TIME_DISTANCE

// Note: this test involves randomness and therefore could in extremely rare cases have inconsistent results.
func TestRandomTimestamp(t *testing.T) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	covered := make(map[int64]bool)
	for i := 0; i < TEST_RANDOM_TIMESTAMP_RUNS; i++ {
		randTimestamp := randomTimestamp(TEST_TIME_ORIGIN, TEST_TIME_DISTANCE)
		if !assert.True(
			t,
			randTimestamp >= TEST_TIME_RANGE_START,
			fmt.Sprintf(
				"%d is not between %d and %d (inclusive). Seed: %d.",
				randTimestamp,
				TEST_TIME_RANGE_START,
				TEST_TIME_RANGE_END,
				seed,
			),
		) {
			t.FailNow()
		}
		if !assert.True(
			t,
			randTimestamp <= TEST_TIME_RANGE_END,
			fmt.Sprintf(
				"%d is not between %d and %d (inclusive). Seed: %d.",
				randTimestamp,
				TEST_TIME_RANGE_START,
				TEST_TIME_RANGE_END,
				seed,
			),
		) {
			t.FailNow()
		}
		covered[randTimestamp] = true
	}
	for i := -1; i <= 1; i++ {
		assert.True(
			t,
			covered[int64(i)],
			fmt.Sprintf(
				"%d was not covered in %d runs of randomTimestamp. Seed: %d.",
				i,
				TEST_RANDOM_TIMESTAMP_RUNS,
				seed,
			),
		)
	}
}

func TestGenerateTimestamp(t *testing.T) {
	assert.Equal(t, int64(5), generateTimestamp(0, 5, 10))
}
