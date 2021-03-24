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

func (v *CsvViewer) getLine(rowNumber int) string {
	var line string
	record := *v.Data.Records[rowNumber]
	for i, headerCol := range *v.Data.Header {
		col := record[i]
		line += col.Slice(0, minInt(col.RuneCount(), headerCol.RuneCount()))
		p := maxInt(headerCol.RuneCount()-col.RuneCount(), 0) + 1
		line += strings.Repeat(WHITESPACE, p)
	}
	return line
}

func (v *CsvViewer) getLines(rowSize int) *[]string {
	var lines []string

	upperLimit := minInt(v.ViewPos+rowSize, len(v.Data.Records))
	for i := v.ViewPos; i < upperLimit; i++ {
		lines = append(lines, v.getLine(i))
	}
	return &lines
}

func (v *CsvViewer) headerSizeView() *[]string {
	_, h := termbox.Size()

	var header []string
	for _, col := range *v.Data.Header {
		header = append(header, col.String())
	}
	lines := append([]string{strings.Join(header, WHITESPACE)}, *v.getLines(h - 2)...)

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

	for i, row := range *v.headerSizeView() {
		RenderLine(0, i, row)
	}

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
