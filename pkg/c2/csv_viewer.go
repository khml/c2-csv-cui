package c2

import (
	"github.com/nsf/termbox-go"
	"strings"
)

type CsvViewer struct {
	Data    *CsvData
	ViewPos int
}

func NewCsvViewer(c *CsvData) *CsvViewer {
	v := new(CsvViewer)
	v.Data = c
	return v
}

func (v *CsvViewer) getLine(n int) string {
	var line string
	for _, col := range *v.Data.Records[n] {
		line += col.String() + WHITESPACE
	}
	return line
}

func (v *CsvViewer) getLines(n int) *[]string {
	var lines []string

	upperLimit := minInt(v.ViewPos+n, len(v.Data.Records))
	for i := v.ViewPos; i < upperLimit; i++ {
		lines = append(lines, v.getLine(i))
	}
	return &lines
}

func (v *CsvViewer) Down() {
	v.ViewPos = minInt(v.ViewPos+1, len(v.Data.Records))
}

func (v *CsvViewer) DownN(n int) {
	v.ViewPos = minInt(v.ViewPos+n, len(v.Data.Records))
}

func (v *CsvViewer) Up() {
	v.ViewPos = maxInt(0, v.ViewPos-1)
}

func (v *CsvViewer) UpN(n int) {
	v.ViewPos = maxInt(0, v.ViewPos-n)
}

func (v *CsvViewer) Render() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		panic(err)
	}

	_, h := termbox.Size()

	var header []string
	for _, col := range *v.Data.Header {
		header = append(header, col.String())
	}

	headerLine := strings.Join(header, WHITESPACE)
	RenderLine(0, 0, headerLine)

	for i, row := range *v.getLines(h - 2) {
		RenderLine(0, i+1, row)
	}

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
