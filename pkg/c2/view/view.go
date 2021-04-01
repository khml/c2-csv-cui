package view

import (
	"c2/pkg/c2/c2csv"
	"c2/pkg/c2/util"
	"strings"
)

type View interface {
	GetLines(data *c2csv.CsvData, viewPos, rowSize int) *[]string
}

const ColMargin = 2

func headWithPadding(c *c2csv.Cell, colSize int) string {
	wordSize := util.MinInt(c.Len(), colSize)
	s := c.Head(wordSize)

	paddingSize := colSize - wordSize
	if paddingSize > 0 {
		s += strings.Repeat(util.WHITESPACE, paddingSize)
	}

	return s
}
