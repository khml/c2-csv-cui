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

	header := NewCells(&csvHeader)

	var columns []Cells
	for i := 0; i < header.Size(); i++ {
		columns = append(columns, Cells{})
	}

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		for i, s := range row {
			columns[i].addStr(s)
		}
	}

	return NewCsvData(header, &columns), nil
}
