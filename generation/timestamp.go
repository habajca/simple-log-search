package generation

import "math/rand"

func randomTimestamp(origin int64, timeDistance int) int64 {
	return generateTimestamp(origin, timeDistance, rand.Int63n(2*int64(timeDistance)+1))
}

func generateTimestamp(origin int64, timeDistance int, randOffset int64) int64 {
	return origin - int64(timeDistance) + randOffset
}
