package c2

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvHeader = []string
type CsvRecords = [][]string
type TableMap = map[string]CsvRecords

type CsvData struct {
	Header  CsvHeader
	Records CsvRecords
}

func ReadCsv(filepath string, skipRowNum int) (CsvData, error) {
	var records CsvRecords

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

	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, r)
	}
	return CsvData{Header: csvHeader, Records: records}, nil
}

func (c *CsvData) Print() {
	fmt.Printf("%v\n", c.Header)
	for _, record := range c.Records {
		fmt.Printf("%v\n", record)
	}
}
