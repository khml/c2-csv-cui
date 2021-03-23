package c2

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/exp/utf8string"
	"os"
)

type CsvRecord = *[]utf8string.String
type CsvRecords = []CsvRecord

type CsvData struct {
	Header  CsvRecord
	Records CsvRecords
}

func (c *CsvData) Print() {
	fmt.Printf("%v\n", c.Header)
	for _, record := range c.Records {
		fmt.Printf("%v\n", record)
	}
}

func ReadCsv(filepath string, skipRowNum int) (CsvData, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return CsvData{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// skip rows
	for i := skipRowNum; i > 0; i-- {
		_, err := reader.Read()
		if err != nil {
			return CsvData{}, err
		}
	}

	csvHeader, err := reader.Read()
	if err != nil {
		return CsvData{}, err
	}

	header := toCsvRecord(&csvHeader)

	var records CsvRecords
	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, toCsvRecord(&r))
	}
	return CsvData{Header: header, Records: records}, nil
}

func toCsvRecord(rows *[]string) CsvRecord {
	var utf8Rows []utf8string.String
	for _, val := range *rows {
		utf8Rows = append(utf8Rows, *utf8string.NewString(val))
	}
	return &utf8Rows
}
