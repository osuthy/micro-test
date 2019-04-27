package table

import (
	"sort"
)

type Row struct {
	Columns []*Column
}

func NewRow(columns []*Column) *Row {
	row := new(Row)
	row.Columns = columns
	return row
}

func (this *Row) Override(row *Row) *Row {
	columns := []*Column{}
	for _, column := range row.Columns {
		if !this.hasSameName(column) {
			columns = append(columns, column)
		}
	}
	return NewRow(append(this.Columns, columns...)).Sorted()
}

func (this *Row) Sorted() *Row {
	columns := []*Column{}
	for _, column := range this.Columns {
		columns = append(columns, column)
	}
	sort.Slice(columns, func(i, j int) bool { return columns[i].HigherOrderThan(columns[j]) })
	return NewRow(columns)
}

func (this *Row) hasSameName(column *Column) bool {
	for _, thisColumn := range this.Columns {
		if thisColumn.HasSameName(column) {
			return true
		}
	}
	return false
}

