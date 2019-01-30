package db

import (
	"reflect"
	"sort"
)

type Column struct {
	name string
	value interface{}
}

type Row struct {
	columns []Column
}

type Table struct {
	name string
	rows []Row
}

func (this Table) isSame(other Table) bool {
	for _, row := range this.rows {
		sort.Slice(row.columns, func(i, j int) bool { return row.columns[i].name < row.columns[j].name })
	}
	for _, row := range other.rows {
		sort.Slice(row.columns, func(i, j int) bool { return row.columns[i].name < row.columns[j].name })
	}
	return reflect.DeepEqual(this.rows, other.rows)
}

func (this Table) filledTableWith(row Row) Table {
	newRows := []Row{}
	for _, thisRow := range this.rows {
		newRow := thisRow.override(row)
		newRows = append(newRows, newRow)
	}
	return Table{rows: newRows, name: this.name}
}

func (this Row) contains(columnName string) bool {
	for _, thisColumn := range this.columns {
		if thisColumn.name == columnName {
			return true
		}
	}
	return false
}

func (this Row) override(row Row) Row {
	columns := []Column{}
	for _, column := range row.columns {
		if !this.contains(column.name) {
			columns = append(columns, column)
		}
	}
	filledColumns := append(this.columns, columns...)
	sort.Slice(filledColumns, func(i, j int) bool { return filledColumns[i].name < filledColumns[j].name })
	return Row {filledColumns}
}
