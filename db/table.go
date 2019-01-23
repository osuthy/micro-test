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

