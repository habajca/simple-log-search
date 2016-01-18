package data

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
		Latitude:  37.7576171,
		Longitude: -122.5776844,
	}
	rowAsString, err := originalRow.String()
	if assert.Nil(t, err) {
		endRow := LogRow{}
		err = (&endRow).FromString(rowAsString)
		if assert.Nil(t, err) {
			assert.Equal(t, originalRow, endRow)
		}
	}
}
