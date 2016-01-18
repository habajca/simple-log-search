package generation

import (
	"math/rand"
	"strconv"
)

func randomUid(uidCount int) string {
	return strconv.Itoa(rand.Intn(uidCount))
}
