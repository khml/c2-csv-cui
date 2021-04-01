package view

import (
	"c2/pkg/c2/c2csv"
	"c2/pkg/c2/util"
	"github.com/nsf/termbox-go"
	"strings"
)

type EqDividedView struct{}

func (v *EqDividedView) makeLine(r *c2csv.Row, colSize int, margin string) string {
	var cols []string
	for _, c := range r.Cells {
		cols = append(cols, headWithPadding(&c, colSize))
	}
	return strings.Join(cols, margin)
}

func (v *EqDividedView) GetLines(data *c2csv.CsvData, viewPos, rowSize int) *[]string {
	w, _ := termbox.Size()
	colNum := len(data.Header.Cells)
	colSize := (w - ColMargin*colNum) / colNum
	margin := strings.Repeat(util.WHITESPACE, ColMargin)

	lines := []string{v.makeLine(data.Header, colSize, margin)}

	upperLimit := util.MinInt(viewPos+rowSize, len(data.Records))
	for i := viewPos; i < upperLimit; i++ {
		lines = append(lines, v.makeLine(data.Records[i], colSize, margin))
	}
	return &lines
}
