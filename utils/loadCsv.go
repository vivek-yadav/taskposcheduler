package utils

import (
	"encoding/csv"
	"os"
)

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
