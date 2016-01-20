package generation

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/habajca/simple-log-search/util"
	"io/ioutil"
	"time"
)

func GenerateTestData(
	directory string,
	fileCount, rowCount int,
	timeOrigin int64, timeDistance int,
	uidCount int,
	domainsFilename string,
	geoOrigin util.GeoPoint, geoDistance int,
) (filenames []string, err error) {
	domains, err := newDomains(domainsFilename)
	if err != nil {
		return nil, err
	}
	filenames = make([]string, fileCount)
	progressBar := pb.StartNew(fileCount)
	for i := 0; i < fileCount; i++ {
		filename, err := generateTestDataFile(
			directory,
			rowCount,
			timeOrigin, timeDistance,
			uidCount,
			domains,
			geoOrigin, geoDistance,
		)
		if err != nil {
			return nil, err
		}
		progressBar.Increment()
		filenames[i] = filename
	}
	progressBar.FinishPrint("All done!")
	return filenames, nil
}

func generateTestDataFile(
	directory string,
	rowCount int,
	timeOrigin int64, timeDistance int,
	uidCount int,
	domains domains,
	geoOrigin util.GeoPoint, geoDistance int,
) (filename string, err error) {
	output := ""
	for i := 0; i < rowCount; i++ {
		row := generateTestDataRow(
			timeOrigin, timeDistance,
			uidCount,
			domains,
			geoOrigin, geoDistance,
		)
		rowOutput := util.StructToString(row)
		output = output + rowOutput + "\n"
	}
	filename = generateFilename()
	err = ioutil.WriteFile(directory+"/"+filename, []byte(output), 0644)
	return filename, err
}

func generateTestDataRow(
	timeOrigin int64, timeDistance int,
	uidCount int,
	domains domains,
	geoOrigin util.GeoPoint, geoDistance int,
) util.LogRow {
	return util.LogRow{
		Timestamp: randomTimestamp(timeOrigin, timeDistance),
		Uid:       randomUid(uidCount),
		Domain:    domains.random(),
		Geo:       randomGeoPoint(geoOrigin, geoDistance),
	}
}

func generateFilename() string {
	return fmt.Sprintf("generated_data_%d.txt", time.Now().UnixNano())
}
