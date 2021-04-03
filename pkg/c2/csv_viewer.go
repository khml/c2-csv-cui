package c2

import (
	"c2/pkg/c2/c2csv"
	"c2/pkg/c2/util"
	"github.com/nsf/termbox-go"
)

type CsvViewer struct {
	Data    *c2csv.CsvData
	ViewPos int
	cmdLine []rune
}

func NewCsvViewer(c *c2csv.CsvData) *CsvViewer {
	v := new(CsvViewer)
	v.Data = c
	return v
}

func (v *CsvViewer) render() {
	w, h := termbox.Size()
	ColMargin := 2
	colWidth := w/v.Data.ColSize - ColMargin

	for i, column := range v.Data.Columns {
		x := (colWidth + ColMargin) * i
		RenderLine(x, 0, v.Data.Header.Get(i).String())

		for y, cell := range column.Sub(v.ViewPos, v.ViewPos+h-2) {
			RenderLine(x, y+1, cell.Head(colWidth))
		}
	}
	RenderLine(0, h-1, util.COLON+string(v.cmdLine))
}

func (v *CsvViewer) BackspaceToCmd() {
	l := len(v.cmdLine)
	if l == 0 {
		return
	}
	v.cmdLine = v.cmdLine[:l-1]
}

func (v *CsvViewer) Down() {
	v.ViewPos = util.MinInt(v.ViewPos+1, v.Data.RowSize)
}

func (v *CsvViewer) DownN(n int) {
	v.ViewPos = util.MinInt(v.ViewPos+n, v.Data.RowSize)
}

func (v *CsvViewer) InputToCmd(r rune) {
	v.cmdLine = append(v.cmdLine, r)
}

func (v *CsvViewer) Up() {
	v.ViewPos = util.MaxInt(0, v.ViewPos-1)
}

func (v *CsvViewer) UpN(n int) {
	v.ViewPos = util.MaxInt(0, v.ViewPos-n)
}

func (v *CsvViewer) Render() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	v.render()

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
