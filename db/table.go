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
	row2 := this.rows[0]
	row3 := row2.override(row)
	return Table{rows: []Row{row3}}
}

func (this Row) override(row Row) Row {
	columns := []Column{}
	for _, column := range row.columns {
		contains := false
		for _, column2 := range this.columns {
			if reflect.DeepEqual(column, column2) {
				contains = true
			}
		}
		if !contains {
			columns = append(columns, column)
		}
	}
	filledColumns := append(this.columns, columns...)
	sort.Slice(filledColumns, func(i, j int) bool { return filledColumns[i].name < filledColumns[j].name })
	return Row {filledColumns}
}
