package main

import (
	"errors"
	"fmt"
	"github.com/habajca/simple-log-search/data"
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
	flag.Int64VarP(&timeOrigin, "search-time", "t", time.Now().Unix(), "The origin of the search in the time dimension as a unix timestamp (in seconds). (defaults to now) (search only)")
}

var timeFrame int

func init() {
	flag.IntVarP(&timeFrame, "search-timeframe", "m", 3600, "The search space in the time dimension in seconds. (search only)")
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
	flag.VarP(&geo, "search-geo", "g", "The origin of the search in the geo dimensions as a latitude, longitude tuple. (search only)")
}

var distance int

func init() {
	flag.IntVarP(&distance, "search-distance", "d", 5000, "The search space in the geo dimensions (in meters). (search only)")
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
	flag.IntVarP(&uidCount, "generation-uids", "u", 1000, "The number of uids in all test files.")
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			os.Stderr,
			`Usage:
%s [log_directory] [options]
or
%s [log_directory] [domains_file] --generate [options]
with options:
`,
			os.Args[0],
			os.Args[0],
		)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	fmt.Println(generateTestFiles, timeOrigin, timeFrame, geo, distance, fileCount, rowCount, uidCount)
}
