package c2csv

type Row struct {
	Cells []Cell
}

func (r *Row) ToStrSlice() []string {
	var header []string
	for _, col := range r.Cells {
		header = append(header, col.String())
	}
	return header
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
