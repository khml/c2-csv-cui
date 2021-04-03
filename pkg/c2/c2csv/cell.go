package c2csv

import (
	"c2/pkg/c2/util"
)

type Cell struct {
	val []rune
}

func NewCsvCell(str string) *Cell {
	cell := new(Cell)
	cell.val = str2rune(str)
	return cell
}

func (c *Cell) Head(n int) string {
	return string(c.val[0:util.MinInt(c.Len(), n)])
}

func (c *Cell) Len() int {
	return len(c.val)
}

func (c *Cell) String() string {
	return string(c.val)
}

func str2rune(s string) []rune {
	return []rune(s)
}
