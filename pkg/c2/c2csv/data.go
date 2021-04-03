package c2csv

import (
	"c2/pkg/c2/util"
	"fmt"
	"strings"
)

type CsvData struct {
	Header  Cells
	Columns []Cells
	RowSize int
	ColSize int
}

func NewCsvData(h *Cells, cols *[]Cells) *CsvData {
	d := new(CsvData)
	d.Header = *h
	d.Columns = *cols
	d.RowSize = d.Columns[0].Size()
	d.ColSize = len(h.Values)
	return d
}

func (c *CsvData) Print() {
	fmt.Printf("%v\n", c.Header)

	for i := 0; i < c.RowSize; i++ {
		var row []string
		for _, col := range c.Columns {
			row = append(row, col.Get(i).String())
		}
		println(strings.Join(row, util.COMMA))
	}
}
