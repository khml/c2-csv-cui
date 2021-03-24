package c2csv

import (
	"fmt"
)

type CsvData struct {
	Header  *Row
	Records []*Row
}

func (c *CsvData) Print() {
	fmt.Printf("%v\n", c.Header)
	for _, record := range c.Records {
		fmt.Printf("%v\n", record)
	}
}
