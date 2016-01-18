package generation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Note: generation.randomTimestamp is not tested as it is random by nature and testing could result in inconsistent results. generation.generateTimestamp, which contains all non-random logic is tested.

func TestGenerateTimestamp(t *testing.T) {
	assert.Equal(t, int64(-10), generateTimestamp(0, 10, 0))
	assert.Equal(t, int64(10), generateTimestamp(0, 10, 1))
}
