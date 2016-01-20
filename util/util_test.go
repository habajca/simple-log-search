package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Test to make marshalling and unmarshalling lines up for sanity's sake
func TestLogRowToStringRoundTrip(t *testing.T) {
	originalRow := LogRow{
		Timestamp: time.Now().Unix(),
		Uid:       "uid",
		Domain:    "test.com",
		Geo: GeoPoint{
			Latitude:  37.7576171,
			Longitude: -122.5776844,
		},
	}
	rowAsString := StructToString(originalRow)
	endRow := LogRow{}
	StructFromString(rowAsString, &endRow)
	assert.Equal(t, originalRow, endRow)
}
