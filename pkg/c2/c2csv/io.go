package c2csv

import (
	"encoding/csv"
	"os"
)

func ReadCsv(filepath string, skipRowNum int) (*CsvData, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// skip rows
	for i := skipRowNum; i > 0; i-- {
		_, err := reader.Read()
		if err != nil {
			return nil, err
		}
	}

	csvHeader, err := reader.Read()
	if err != nil {
		return nil, err
	}

	header := toCsvRecord(&csvHeader)

	var records []*Row
	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, toCsvRecord(&r))
	}
	return &CsvData{Header: header, Records: records}, nil
}

func toCsvRecord(rows *[]string) *Row {
	return FromStrings(rows)
}
