package table

import (
	"sort"
)

type Row struct {
	columns []*Column
}

func NewRow(columns []*Column) *Row {
	return &Row{columns: columns}
}

func (this *Row) Override(row *Row) *Row {
	columns := []*Column{}
	for _, column := range row.columns {
		if !this.hasSameName(column) {
			columns = append(columns, column)
		}
	}
	return NewRow(append(this.columns, columns...)).Sorted()
}

func (this *Row) Sorted() *Row {
	columns := []*Column{}
	for _, column := range this.columns {
		columns = append(columns, column)
	}
	sort.Slice(columns, func(i, j int) bool { return columns[i].HigherOrderThan(columns[j]) })
	return NewRow(columns)
}

func (this *Row) hasSameName(column *Column) bool {
	for _, thisColumn := range this.columns {
		if thisColumn.HasSameName(column) {
			return true
		}
	}
	return false
}
func (this *Row) ColumnNames() []string {
	names := make([]string, 0, len(this.columns))
	for _, column := range this.columns {
		names = append(names, column.Name)
	}
	return names
}

func (this *Row) ColumnValues() []interface{} {
	values := make([]interface{}, 0, len(this.columns))
	for _, column := range this.columns {
		values = append(values, column.Value)
	}
	return values
}

func (this *Row) ColumnValueOf(columnName string) interface{} {
	for _, column := range this.columns {
		if column.Name == columnName {
			return column.Value
		}
	}
	return nil
}
