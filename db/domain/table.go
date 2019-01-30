package domain

import (
	"reflect"
	"sort"
)

type Column struct {
	Name string
	Value interface{}
}

type Row struct {
	Columns []Column
}

type Table struct {
	Name string
	Rows []Row
}

func (this Table) IsSame(other Table) bool {
	for _, row := range this.Rows {
		sort.Slice(row.Columns, func(i, j int) bool { return row.Columns[i].Name < row.Columns[j].Name })
	}
	for _, row := range other.Rows {
		sort.Slice(row.Columns, func(i, j int) bool { return row.Columns[i].Name < row.Columns[j].Name })
	}
	return reflect.DeepEqual(this.Rows, other.Rows)
}

func (this Table) FilledTableWith(row Row) Table {
	newRows := []Row{}
	for _, thisRow := range this.Rows {
		newRow := thisRow.Override(row)
		newRows = append(newRows, newRow)
	}
	return Table{Rows: newRows, Name: this.Name}
}

func (this Row) contains(columnName string) bool {
	for _, thisColumn := range this.Columns {
		if thisColumn.Name == columnName {
			return true
		}
	}
	return false
}

func (this Row) Override(row Row) Row {
	columns := []Column{}
	for _, column := range row.Columns {
		if !this.contains(column.Name) {
			columns = append(columns, column)
		}
	}
	filledColumns := append(this.Columns, columns...)
	sort.Slice(filledColumns, func(i, j int) bool { return filledColumns[i].Name < filledColumns[j].Name })
	return Row {filledColumns}
}
