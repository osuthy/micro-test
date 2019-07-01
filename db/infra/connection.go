package infra

import (
	"database/sql"
	. "github.com/ShoichiroKitano/micro_test/db/table"
	"reflect"
)

type Connection struct {
	Driver *sql.DB
	rdbms string
}

func NewConnection(driver *sql.DB, rdbms string) *Connection {
	return &Connection{Driver: driver, rdbms: rdbms}
}

func (this *Connection) FindTable(table *Table) *Table {
	rows, _ := this.Driver.Query(table.SelectAllQuery())
	defer rows.Close()
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := []*Row{}
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	return NewTable(table.Name, newRows)
}

func (this *Connection) FindColumnDefinition(table *Table) *ColumnDefinition {
	return this.mysqlColumnDefinition(table)
}

func (this *Connection) mysqlColumnDefinition(table *Table) *ColumnDefinition {
	rows, _ := this.Driver.Query(table.MysqlColumnDefinitionQuery())
	defer rows.Close()
	infos := []*ColumnMetaInformation{}
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := []*Row{}
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	for _, row := range newRows {
		infos = append(infos, NewColumnMetaInformation(
			row.ColumnValueOf("Field").(string),
			row.ColumnValueOf("Type").(string),
			row.ColumnValueOf("Key").(string),
			isNullable(row.ColumnValueOf("Null").(string)),
			row.ColumnValueOf("Default"),
			row.ColumnValueOf("Extra").(string),
		))
	}
	return NewColumnDefinition(infos)
}

func isNullable(value string) bool {
	if value == "YES" {
		return true
	}
	return false // else if "NO"
}

func (this *Connection) StoreTable(table *Table) {
	tx, _ := this.Driver.Begin()
	stmt, _ := tx.Prepare(table.InsertQuery())
	for _, values := range table.AllValues() {
		stmt.Exec(values...)
	}
	tx.Commit()
}

func (this *Connection) TruncateTable(table *Table) {
	this.Driver.Exec(table.TruncateQuery())
}

func scanArgs(types []*sql.ColumnType) []interface{} {
	dataPtrs := make([]interface{}, len(types))
	for i := range types {
		if types[i].DatabaseTypeName() == "INT" {
			dataPtrs[i] = new(int)
		} else if types[i].DatabaseTypeName() == "CHAR" {
			dataPtrs[i] = new(string)
		}
	}
	return dataPtrs
}

func rowFrom(types []*sql.ColumnType, dataPtrs []interface{}) *Row {
	columns := make([]*Column, len(types))
	for i := range types {
		refv := reflect.ValueOf(dataPtrs[i])
		if r, ok := refv.Interface().(*int); ok {
			columns[i] = NewColumn(types[i].Name(), *r)
		}
		if r, ok := refv.Interface().(*string); ok {
			columns[i] = NewColumn(types[i].Name(), *r)
		}
	}
	return NewRow(columns)
}
