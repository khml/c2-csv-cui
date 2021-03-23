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

func (v *CsvViewer) getLine(i int) string {
	return strings.Join(v.Data.Records[v.ViewPos+i], WHITESPACE)
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

	header := strings.Join(v.Data.Header, WHITESPACE)
	RenderLine(0, 0, header)

	for i := 0; i < h-2; i++ {
		RenderLine(0, i+1, v.getLine(i+1))
	}

	err = termbox.Flush()
	if err != nil {
		panic(err)
	}
}
