package table

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"reflect"
)

type Table struct {
	name string
	rows []*Row
}

func NewTable(name string, rows []*Row) *Table {
	return &Table{name: name, rows: rows}
}

func (this *Table) Name() string {
	return this.name
}

func (this *Table) IsSameAsTable(other *Table) bool {
	if this.name != other.name {
		return false
	}
	return reflect.DeepEqual(this.rowsColumnsSorted(), other.rowsColumnsSorted())
}

func (this *Table) FilledTableWith(row *Row) *Table {
	newRows := []*Row{}
	for _, thisRow := range this.rows {
		newRow := thisRow.Override(row)
		newRows = append(newRows, newRow)
	}
	return NewTable(this.name, newRows)
}

func (this *Table) rowsColumnsSorted() []*Row {
	sortedRows := []*Row{}
	for _, row := range this.rows {
		sortedRows = append(sortedRows, row.Sorted())
	}
	return sortedRows
}

func (this *Table) SelectAllQuery() string {
	sql, _, _ := sq.Select("*").From(this.name).ToSql()
	return sql
}

func (this *Table) InsertQuery() string {
	sql, _, _ := sq.Insert(this.name).
		Columns(this.rows[0].ColumnNames()...).
		Values(this.rows[0].ColumnValues()...).ToSql()
	return sql
}

func (this *Table) TruncateQuery() string {
	return fmt.Sprintf("truncate table %s;", this.name)
}

func (this *Table) AllValues() [][]interface{} {
	allValue := make([][]interface{}, 0, len(this.rows))
	for _, row := range this.rows {
		allValue = append(allValue, row.ColumnValues())
	}
	return allValue
}

func (this *Table) MysqlColumnDefinitionQuery() string {
	return fmt.Sprintf("show columns from %s;", this.name)
}
