package data

import (
	"encoding/json"
)

type LogRow struct {
	Timestamp int64
	Uid       string
	Domain    string
	Geo       GeoPoint
}

func (r LogRow) String() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *LogRow) FromString(s string) error {
	return json.Unmarshal([]byte(s), r)
}
