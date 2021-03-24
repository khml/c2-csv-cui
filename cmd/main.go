package main

import (
	"c2/pkg/c2"
	"c2/pkg/c2/c2csv"
	"github.com/nsf/termbox-go"
	"log"
	"os"
)

func pollEvent(v *c2.CsvViewer) {
	v.Render()
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc, termbox.KeyCtrlC:
				return
			case termbox.KeyArrowDown:
				v.Down()
				v.Render()
			case termbox.KeyArrowUp:
				v.Up()
				v.Render()
			case termbox.KeySpace:
				_, h := termbox.Size()
				v.DownN(h / 2)
				v.Render()
			default:
				v.Render()
			}
		case termbox.EventResize:
			v.Render()
		default:
			v.Render()
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s path/to/c2csv \n", os.Args[0])
	}

	filepath := os.Args[1]
	csvData, err := c2csv.ReadCsv(filepath, 0)
	if err != nil {
		log.Fatal(err)
	}

	err = termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	v := c2.NewCsvViewer(csvData)

	pollEvent(v)
}
