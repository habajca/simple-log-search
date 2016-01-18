package generation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

const TEST_RANDOM_UID_RUNS = 1000
const TEST_UIDS_COUNT = 2

// Note: this test involves randomness and therefore could in extremely rare cases have inconsistent results.
func TestRandomUid(t *testing.T) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	covered := make(map[string]bool)
	for i := 0; i < TEST_RANDOM_UID_RUNS; i++ {
		covered[randomUid(TEST_UIDS_COUNT)] = true
	}
	for i := 0; i < TEST_UIDS_COUNT; i++ {
		assert.True(
			t,
			covered[strconv.Itoa(i)],
			fmt.Sprintf(
				"%d should be covered by %d runs of randomUid. Seed: %d.",
				i,
				TEST_RANDOM_UID_RUNS,
				seed,
			),
		)
	}
}
