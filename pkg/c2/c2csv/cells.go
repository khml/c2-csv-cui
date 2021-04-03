package c2csv

import "c2/pkg/c2/util"

type Cells struct {
	Values []*Cell
}

func NewCells(strSlice *[]string) *Cells {
	var cells []*Cell
	for _, val := range *strSlice {
		cells = append(cells, NewCsvCell(val))
	}

	return &Cells{Values: cells}
}

func (c *Cells) addStr(s string) {
	c.Values = append(c.Values, NewCsvCell(s))
}

func (c *Cells) Get(index int) *Cell {
	return c.Values[index]
}

func (c *Cells) Size() int {
	return len(c.Values)
}

func (c *Cells) Sub(start, end int) []*Cell {
	end = util.MinInt(end, c.Size())
	return c.Values[start:end]
}

func (c *Cells) ToStrSlice() []string {
	var header []string
	for _, col := range c.Values {
		header = append(header, col.String())
	}
	return header
}
