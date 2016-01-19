package map_reduce

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func panicAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func TestMap(t *testing.T) {
	filterNegative := func(s string) (string, bool) {
		i := panicAtoi(s)
		if i < 0 {
			return "", false
		}
		return strconv.Itoa(i), true
	}

	input := []string{"-1", "1"}
	output := DoMap(input, filterNegative)
	assert.Len(t, output, 1)
	assert.Equal(t, "1", output[0])
}

func TestReduce(t *testing.T) {
	oddOrEven := func(s string) string {
		i := panicAtoi(s)
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	}

	oddFirst := func(l string, r string) bool {
		if l == r {
			return false
		}
		if l == "odd" {
			return true
		}
		return false
	}

	sumUp := func(acc []string, s string) []string {
		if len(acc) == 0 {
			return []string{s}
		}
		i := panicAtoi(s)
		sum := i + panicAtoi(acc[0])
		return []string{strconv.Itoa(sum)}
	}

	inputs := []string{"0", "1", "2", "3"}
	outputs := DoReduce(inputs, oddOrEven, oddFirst, sumUp)
	assert.Len(t, outputs, 2)
	assert.Equal(t, "4", outputs[0])
	assert.Equal(t, "2", outputs[1])
}
