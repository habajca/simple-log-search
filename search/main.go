package search

import (
	"bufio"
	"fmt"
	"github.com/habajca/simple-log-search/map_reduce"
	"github.com/habajca/simple-log-search/util"
	"io/ioutil"
	"os"
)

// Note: This code was written to adhere to the map reduce paradigm. Some code may very well be written in an obtuse manor in order to adhere.

func OutputSearchResults(
	dirname string,
	timeOrigin int64, timeDistance int,
	geoOrigin util.GeoPoint, geoDistance int,
) error {
	outputs, err := searchFilesInDirectory(dirname, timeOrigin, timeDistance, geoOrigin, geoDistance)
	if err != nil {
		return err
	}
	for _, output := range outputs {
		fmt.Printf("%s, %d\n", output.Domain, output.Count)
	}
	return nil
}

func searchFilesInDirectory(
	dirname string,
	timeOrigin int64, timeDistance int,
	geoOrigin util.GeoPoint, geoDistance int,
) ([]outputRow, error) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	fmt.Println("Importing files...")
	inputs := []string{}
	for _, fileInfo := range fileInfos {
		fileStrings, err := openFile(dirname + "/" + fileInfo.Name())
		if err != nil {
			return nil, err
		}
		inputs = append(inputs, fileStrings...)
	}

	fmt.Println("Filtering rows...")
	filtered := filterRows(inputs, timeOrigin, timeDistance, geoOrigin, geoDistance)

	fmt.Println("Reducing to counts per domain...")
	unsortedOutput := reduceToOutput(filtered)

	fmt.Println("Sorting output...")
	sortedOutputStrings := sortOutputRows(unsortedOutput)
	outputs := make([]outputRow, len(sortedOutputStrings))
	for i, s := range sortedOutputStrings {
		output := outputRow{}
		util.StructFromString(s, &output)
		outputs[i] = output
	}
	return outputs, nil
}

func openFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	strings := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	return strings, scanner.Err()
}

func filterRows(
	rows []string,
	timeOrigin int64, timeDistance int,
	geoOrigin util.GeoPoint, geoDistance int,
) []string {
	withinTimeRange := func(time int64) bool {
		if time < timeOrigin-int64(timeDistance) {
			return false
		}
		if time > timeOrigin+int64(timeDistance) {
			return false
		}
		return true
	}

	pointOrigin := geoOrigin.Point()
	withinGeoRange := func(geo util.GeoPoint) bool {
		distance := pointOrigin.GreatCircleDistance(geo.Point())
		return distance <= float64(geoDistance)/1000
	}

	filterRow := func(s string) (string, bool) {
		row := &util.LogRow{}
		util.StructFromString(s, &row)
		return s, withinTimeRange(row.Timestamp) && withinGeoRange(row.Geo)
	}

	return map_reduce.DoMap(rows, filterRow)
}

type outputRow struct {
	Domain string
	Count  int
}

func domainPartition(s string) string {
	row := util.LogRow{}
	util.StructFromString(s, &row)

	return row.Domain
}

func lessDomain(l string, r string) bool {
	return l < r
}

func countDomainRows(acc []string, s string) []string {
	if len(acc) == 0 {
		row := util.LogRow{}
		util.StructFromString(s, &row)
		return []string{util.StructToString(outputRow{
			Domain: row.Domain,
			Count:  1,
		})}
	}
	row := outputRow{}
	util.StructFromString(acc[0], &row)
	row.Count = row.Count + 1
	return []string{util.StructToString(row)}
}

func reduceToOutput(rows []string) []string {
	return map_reduce.DoReduce(rows, domainPartition, lessDomain, countDomainRows)
}

func defaultPartition(s string) string {
	return s
}

func lessOutputRow(ls string, rs string) bool {
	if ls == rs {
		return false
	}
	l := outputRow{}
	r := outputRow{}
	util.StructFromString(ls, &l)
	util.StructFromString(rs, &r)

	if l.Count < r.Count {
		return false
	}
	if l.Domain < r.Domain {
		return false
	}
	return true
}

func defaultReduce(acc []string, s string) []string {
	return append(acc, s)
}

func sortOutputRows(rows []string) []string {
	return map_reduce.DoReduce(rows, defaultPartition, lessOutputRow, defaultReduce)
}
