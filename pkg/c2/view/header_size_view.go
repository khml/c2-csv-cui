package view

import (
	"c2/pkg/c2/c2csv"
	"c2/pkg/c2/util"
	"strings"
)

type HeaderSizeView struct{}

func (v *HeaderSizeView) getLine(data *c2csv.CsvData, rowNumber int) string {
	var line string
	record := *data.Records[rowNumber]
	for i, headerCell := range data.Header.Cells {
		col := record.Cells[i]
		line += col.Head(headerCell.Len())
		p := util.MaxInt(headerCell.Len()-col.Len(), 0) + 1
		line += strings.Repeat(util.WHITESPACE, p)
	}
	return line
}

func (v *HeaderSizeView) GetLines(data *c2csv.CsvData, viewPos, rowSize int) *[]string {
	var lines []string

	upperLimit := util.MinInt(viewPos+rowSize, len(data.Records))
	for i := viewPos; i < upperLimit; i++ {
		lines = append(lines, v.getLine(data, i))
	}
	return &lines
}
