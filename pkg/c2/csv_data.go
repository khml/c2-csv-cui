package c2

import (
	"c2/pkg/c2/c2csv"
	"encoding/csv"
	"fmt"
	"os"
)

type CsvData struct {
	Header  *c2csv.Row
	Records []*c2csv.Row
}

func (c *CsvData) Print() {
	fmt.Printf("%v\n", c.Header)
	for _, record := range c.Records {
		fmt.Printf("%v\n", record)
	}
}

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

	var records []*c2csv.Row
	for {
		r, err := reader.Read()
		if err != nil {
			break
		}
		records = append(records, toCsvRecord(&r))
	}
	return &CsvData{Header: header, Records: records}, nil
}

func toCsvRecord(rows *[]string) *c2csv.Row {
	return c2csv.FromStrings(rows)
}
