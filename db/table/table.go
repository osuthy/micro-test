package table

import (
	"reflect"
)

type Table struct {
	Name string
	Rows []*Row
}

func NewTable(name string, rows []*Row) *Table {
	return &Table{Name: name, Rows: rows}
}

func (this *Table) IsSameAsTable(other *Table) bool {
	if this.Name != other.Name {
		return false
	}
	return reflect.DeepEqual(this.rowsColumnsSorted(), other.rowsColumnsSorted())
}

func (this *Table) FilledTableWith(row *Row) *Table {
	newRows := []*Row{}
	for _, thisRow := range this.Rows {
		newRow := thisRow.Override(row)
		newRows = append(newRows, newRow)
	}
	return NewTable(this.Name, newRows)
}

func (this *Table) rowsColumnsSorted() []*Row {
	sortedRows := []*Row{}
	for _, row := range this.Rows {
		sortedRows = append(sortedRows, row.Sorted())
	}
	return sortedRows
}

func (this *Table) SelectAllQuery() string {
	return "select * from " + this.Name + ";"
}
