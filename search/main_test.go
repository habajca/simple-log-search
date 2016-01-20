package search

import (
	"github.com/habajca/simple-log-search/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchFile(t *testing.T) {
	outputs, err := searchFilesInDirectory(
		"testdata",
		1453000000, 1000,
		util.GeoPoint{Latitude: 37.7576171, Longitude: -122.5776844}, 5000,
	)
	if !assert.Nil(t, err) {
		t.FailNow()
	}

	expected := []outputRow{
		outputRow{
			Domain: "test0.com",
			Count:  1,
		},
	}
	assert.Equal(t, expected, outputs)
}

func TestFilterRows(t *testing.T) {
	rows := []string{`{"Timestamp":1453000000,"Uid":"78","Domain":"test0.com","Geo":{"Latitude":37.7576171,"Longitude":-122.5776844}}`, `{"Timestamp":1453000000,"Uid":"78","Domain":"test0.com","Geo":{"Latitude":0,"Longitude":0}}`, `{"Timestamp":0,"Uid":"78","Domain":"test0.com","Geo":{"Latitude":37.7576171,"Longitude":-122.5776844}}`}

	outputs := filterRows(
		rows,
		1453000000, 1000,
		util.GeoPoint{Latitude: 37.7576171, Longitude: -122.5776844}, 5000,
	)
	assert.Len(t, outputs, 1)
	assert.Equal(t, rows[0], outputs[0])
}

func outputRowsToStrings(outputRows []outputRow) []string {
	rows := []string{}
	for _, or := range outputRows {
		rows = append(rows, util.StructToString(or))
	}
	return rows
}

func TestReduceToOutput(t *testing.T) {
	rows := []string{`{"Timestamp":1453000000,"Uid":"78","Domain":"test0.com","Geo":{"Latitude":37.7576171,"Longitude":-122.5776844}}`, `{"Timestamp":1453000000,"Uid":"78","Domain":"test0.com","Geo":{"Latitude":0,"Longitude":0}}`, `{"Timestamp":0,"Uid":"78","Domain":"test1.com","Geo":{"Latitude":37.7576171,"Longitude":-122.5776844}}`}

	outputs := reduceToOutput(rows)

	expectedOutputs := outputRowsToStrings([]outputRow{
		outputRow{Domain: "test0.com", Count: 2},
		outputRow{Domain: "test1.com", Count: 1},
	})

	assert.Equal(t, expectedOutputs, outputs)
}

func TestSortOutputRows(t *testing.T) {
	rows := outputRowsToStrings([]outputRow{
		outputRow{Domain: "test0.com", Count: 5},
		outputRow{Domain: "test1.com", Count: 10},
		outputRow{Domain: "test2.com", Count: 7},
	})

	outputs := sortOutputRows(rows)

	expectedOutputs := outputRowsToStrings([]outputRow{
		outputRow{Domain: "test1.com", Count: 10},
		outputRow{Domain: "test2.com", Count: 7},
		outputRow{Domain: "test0.com", Count: 5},
	})

	assert.Equal(t, expectedOutputs, outputs)
}
