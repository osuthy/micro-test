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
