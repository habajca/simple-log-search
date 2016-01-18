package generation

import "math/rand"

func randomTimestamp(origin int64, timeDistance int) int64 {
	return generateTimestamp(origin, rand.Intn(timeDistance), rand.Intn(2))
}

func generateTimestamp(origin int64, timeDistance int, positiveBoolAsInt int) int64 {
	if positiveBoolAsInt == 1 {
		return origin + int64(timeDistance)
	}
	if positiveBoolAsInt == 0 {
		return origin + -1*int64(timeDistance)
	}
	panic("positiveBoolAsInt must be 0 or 1.")
}
