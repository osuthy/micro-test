package domain

import (
	"reflect"
	"sort"
)

type Column struct {
	Name string
	Value interface{}
}

func NewColumn(name string, value interface{}) *Column {
	column := new(Column)
	column.Name = name
	column.Value = value
	return column
}

type Row struct {
	Columns []*Column
}

func NewRow(columns []*Column) *Row {
	row := new(Row)
	row.Columns = columns
	return row
}

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

func (this *Row) contains(columnName string) bool {
	for _, thisColumn := range this.Columns {
		if thisColumn.Name == columnName {
			return true
		}
	}
	return false
}

func (this *Row) Override(row *Row) *Row {
	columns := []*Column{}
	for _, column := range row.Columns {
		if !this.contains(column.Name) {
			columns = append(columns, column)
		}
	}
	filledColumns := append(this.Columns, columns...)
	sort.Slice(filledColumns, func(i, j int) bool { return filledColumns[i].Name < filledColumns[j].Name })
	return NewRow(filledColumns)
}
