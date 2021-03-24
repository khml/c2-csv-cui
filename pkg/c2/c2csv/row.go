package c2csv

type Row struct {
	Cells []Cell
}

func FromStrings(strSlice *[]string) *Row {
	r := new(Row)

	var cells []Cell
	for _, val := range *strSlice {
		cells = append(cells, *NewCsvCell(val))
	}

	r.Cells = cells
	return r
}
