package c2csv

import (
	"c2/pkg/c2/util"
	"golang.org/x/exp/utf8string"
)

type Cell struct {
	Str *utf8string.String
}

func NewCsvCell(str string) *Cell {
	cell := new(Cell)
	cell.Str = utf8string.NewString(str)
	return cell
}

func (c *Cell) Head(n int) string {
	return c.Str.Slice(0, util.MinInt(c.Len(), n))
}

func (c *Cell) Len() int {
	return c.Str.RuneCount()
}

func (c *Cell) String() string {
	return c.Str.String()
}
