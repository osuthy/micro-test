package table

import (
	"reflect"
	"sort"
)

type Table struct {
	Name string
	Rows []*Row
}

func NewTable(name string, rows []*Row) *Table {
	table := new(Table)
	table.Name = name
	table.Rows = rows
	return table
}

// ここが破壊的な操作している
func (this *Table) IsSameAsTable(other *Table) bool {
	if(this.Name != other.Name) { return false }
	for _, row := range this.Rows {
		sort.Slice(row.Columns, func(i, j int) bool { return row.Columns[i].Name < row.Columns[j].Name })
	}
	for _, row := range other.Rows {
		sort.Slice(row.Columns, func(i, j int) bool { return row.Columns[i].Name < row.Columns[j].Name })
	}
	return reflect.DeepEqual(this.Rows, other.Rows)
}

func (this *Table) FilledTableWith(row *Row) *Table {
	newRows := []*Row{}
	for _, thisRow := range this.Rows {
		newRow := thisRow.Override(row)
		newRows = append(newRows, newRow)
	}
	return NewTable(this.Name, newRows)
}

