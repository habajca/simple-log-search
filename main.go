package main

import (
	"errors"
	"fmt"
	"github.com/habajca/simple-log-search/data"
	"github.com/habajca/simple-log-search/generation"
	flag "github.com/ogier/pflag"
	"os"
	"strconv"
	"strings"
	"time"
)

var generateTestFiles bool

func init() {
	flag.BoolVar(&generateTestFiles, "generate", false, "Indicates that test files should be generated.")
}

var timeOrigin int64

func init() {
	flag.Int64VarP(&timeOrigin, "time", "t", time.Now().Unix(), "The origin of the search in the time dimension as a unix timestamp (in seconds). (defaults to now)")
}

var timeFrame int

func init() {
	flag.IntVarP(&timeFrame, "timeframe", "m", 3600, "The search space in the time dimension in seconds.")
}

type geoPoint data.GeoPoint

const defaultGeoPointString = "37.7576171,-122.5776844"

func (geo *geoPoint) String() string {
	return defaultGeoPointString
}

func (geo *geoPoint) Set(value string) error {
	values := strings.SplitN(value, ",", 2)
	if len(values) < 2 {
		return errors.New(fmt.Sprintf("Expecting something like %s.", defaultGeoPointString))
	}

	latitude, latErr := strconv.ParseFloat(values[0], 64)
	if latErr != nil {
		return latErr
	}
	longitude, lonErr := strconv.ParseFloat(values[1], 64)
	if lonErr != nil {
		return lonErr
	}

	geo.Latitude = latitude
	geo.Longitude = longitude
	return nil
}

var geo geoPoint

func init() {
	(&geo).Set(defaultGeoPointString)
	flag.VarP(&geo, "geo", "g", "The origin of the search in the geo dimensions as a latitude, longitude tuple.")
}

var distance int

func init() {
	flag.IntVarP(&distance, "distance", "d", 5000, "The search space in the geo dimensions (in meters).")
}

var fileCount int

func init() {
	flag.IntVarP(&fileCount, "generation-files", "f", 1000, "The number of test files to create. (generation only)")
}

var rowCount int

func init() {
	flag.IntVarP(&rowCount, "generation-rows", "r", 1000, "The number of rows per test file. (generation only)")
}

var uidCount int

func init() {
	flag.IntVarP(&uidCount, "generation-uids", "u", 1000, "The number of uids in all test files. (generation only)")
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`Usage:
%s log_directory [options]
or
%s log_directory domains_file --generate [options]
with options:
`,
			os.Args[0],
			os.Args[0],
		)
		flag.PrintDefaults()
	}
}

func positive(i int64) bool {
	if i < 0 {
		flag.Usage()
		return false
	}
	return true
}

func valid() bool {
	for _, i := range []int{fileCount, rowCount, timeFrame, uidCount, distance} {
		fmt.Println(i)
		if !positive(int64(i)) {
			return false
		}
	}
	if !positive(timeOrigin) {
		return false
	}
	if geo.Latitude > 90 || geo.Latitude < -90 {
		return false
	}
	if geo.Longitude > 180 || geo.Longitude < -180 {
		return false
	}

	return true
}

func main() {
	flag.Parse()
	if !valid() {
		flag.Usage()
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	directory := os.Args[1]
	if generateTestFiles {
		if len(os.Args) < 3 {
			flag.Usage()
			os.Exit(1)
		}
		domainsFilename := os.Args[2]
		_, err := generation.GenerateTestData(
			directory,
			fileCount, rowCount,
			timeOrigin, timeFrame,
			uidCount,
			domainsFilename,
			data.GeoPoint(geo), distance,
		)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}
}
