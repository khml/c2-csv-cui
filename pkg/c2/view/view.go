package view

import (
	"c2/pkg/c2/c2csv"
)

type View interface {
	GetLines(data *c2csv.CsvData, viewPos, rowSize int) *[]string
}
