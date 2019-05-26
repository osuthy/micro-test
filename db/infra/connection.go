package infra

import (
	"database/sql"
	. "github.com/ShoichiroKitano/micro_test/db/table"
	"reflect"
)

type Connection struct {
	Driver *sql.DB
}

func NewConnection(driver *sql.DB) *Connection {
	return &Connection{Driver: driver}
}

func (this *Connection) FindTable(tableName string) *Table {
	rows, _ := this.Driver.Query("SELECT * FROM " + tableName + ";")
	defer rows.Close()
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := make([]*Row, 0)
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	return NewTable(tableName, newRows)
}

func (this *Connection) StoreTable(table *Table) {
	tx, _ := this.Driver.Begin()
	for _, row := range table.Rows {
		tx.Exec("insert into test (" + row.Columns[0].Name + "," + row.Columns[1].Name + ") values (" + toLiteral(row.Columns[0]) + "," + toLiteral(row.Columns[1]) + ");")
	}
	tx.Commit()
}

func (this *Connection) TruncateTable(tableName string) {
	tx, _ := this.Driver.Begin()
	tx.Exec("truncate table " + tableName + ";")
	tx.Commit()
}

func toLiteral(column *Column) string {
	refv := reflect.ValueOf(column.Value)
	r, _ := refv.Interface().(string)
	return "'" + r + "'"
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
