package generation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

const TEST_RANDOM_DOMAIN_RUNS = 1000
const TEST_DOMAINS_COUNT = 3

// Note: this test involves randomness and therefore could in extremely rare cases have inconsistent results.
func TestRandomDomain(t *testing.T) {
	seed := time.Now().Unix()
	rand.Seed(seed)
	domains, err := newDomains("testdata/domains.txt")
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	covered := make(map[string]bool)
	for i := 0; i < TEST_RANDOM_DOMAIN_RUNS; i++ {
		covered[domains.random()] = true
	}
	for i := 0; i < TEST_DOMAINS_COUNT; i++ {
		if !assert.True(
			t,
			covered[fmt.Sprintf("test%d.com", i)],
			fmt.Sprintf(
				"%d should be covered by %d runs of domains.random. Seed: %d.",
				i,
				TEST_RANDOM_DOMAIN_RUNS,
				seed,
			),
		) {
			t.FailNow()
		}
	}
}

func TestNewDomains(t *testing.T) {
	_, err := newDomains("testdata/shouldnt_exist")
	assert.NotNil(t, err)

	domains, err := newDomains("testdata/domains.txt")
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	for i := 0; i < TEST_DOMAINS_COUNT; i++ {
		assert.Equal(t, fmt.Sprintf("test%d.com", i), domains[i])
	}
}
