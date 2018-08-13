package utils

import (
	"encoding/csv"
	"os"
)

var DT_FORMAT = "2006-01-02T15:04:05"
var FIND_DT_FORMAT = "2006-01-02"

func LoadCsv(file string) (records [][]string, err error) {
	var f *os.File
	f, err = os.Open(file)
	if err != nil {
		return
	}

	reader := csv.NewReader(f)
	records, err = reader.ReadAll()
	if err != nil {
		return
	}
	return
}
