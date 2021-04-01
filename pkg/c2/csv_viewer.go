package c2

import (
	"c2/pkg/c2/c2csv"
	"c2/pkg/c2/util"
	"c2/pkg/c2/view"
	"github.com/nsf/termbox-go"
)

type CsvViewer struct {
	Data    *c2csv.CsvData
	ViewPos int
	cmdLine []rune
	view    view.View
}

func NewCsvViewer(c *c2csv.CsvData) *CsvViewer {
	v := new(CsvViewer)
	v.Data = c
	v.view = new(view.EqDividedView)
	return v
}

func (v *CsvViewer) render() {
	_, h := termbox.Size()

	var header []string
	for _, col := range v.Data.Header.Cells {
		header = append(header, col.String())
	}

	lines := *v.view.GetLines(v.Data, v.ViewPos, h-2)
	lines = append(lines, util.COLON+string(v.cmdLine))

	for i, row := range lines {
		RenderLine(0, i, row)
	}
}

func (v *CsvViewer) BackspaceToCmd() {
	l := len(v.cmdLine)
	if l == 0 {
		return
	}
	v.cmdLine = v.cmdLine[:l-1]
}

func (v *CsvViewer) Down() {
	v.ViewPos = util.MinInt(v.ViewPos+1, len(v.Data.Records))
}

func (v *CsvViewer) DownN(n int) {
	v.ViewPos = util.MinInt(v.ViewPos+n, len(v.Data.Records))
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
