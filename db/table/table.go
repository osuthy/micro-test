package table

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
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
	sql, _, _ := sq.Select("*").From(this.Name).ToSql()
	return sql
}

func (this *Table) InsertQuery() string {
	sql, _, _ := sq.Insert(this.Name).
		Columns(this.Rows[0].ColumnNames()...).
		Values(this.Rows[0].ColumnValues()...).ToSql()
	return sql
}

func (this *Table) TruncateQuery() string {
	return fmt.Sprintf("truncate table %s;", this.Name)
}

func (this *Table) AllValues() [][]interface{} {
	allValue := make([][]interface{}, 0, len(this.Rows))
	for _, row := range this.Rows {
		allValue = append(allValue, row.ColumnValues())
	}
	return allValue
}

func (this *Table) MysqlColumnDefinitionQuery() string {
	return fmt.Sprintf("show columns from %s;", this.Name)
}

