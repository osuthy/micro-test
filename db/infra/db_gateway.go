package infra

import (
	"database/sql"
	"reflect"
	. "github.com/ShoichiroKitano/micro_test/db/domain"
)

type Connection struct {
	driver *sql.DB
}

func NewConnection(driver *sql.DB) *Connection {
	this := new(Connection)
	this.driver = driver
	return this
}

func (this *Connection) FindTable(tableName string) Table {
	rows, _ := this.driver.Query("SELECT * FROM " + tableName + ";")
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := make([]Row, 0)
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	return Table{tableName, newRows}
}

func (this *Connection) StoreTable(table Table) {
	for _, row := range table.Rows {
		this.driver.Query("insert into test (" + row.Columns[0].Name + "," + row.Columns[1].Name + ") values (" + toLiteral(row.Columns[0]) + "," + toLiteral(row.Columns[1]) + ");")
	}
}

func (this *Connection) TruncateTable(tableName string) {
	this.driver.Query("truncate table " + tableName + ";")
}

func toLiteral(column Column) string {
	refv := reflect.ValueOf(column.Value)
	r, _ := refv.Interface().(string)
	return "'" + r +"'"
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

func rowFrom(types []*sql.ColumnType, dataPtrs []interface{}) Row {
	columns := make([]Column, len(types))
	for i := range types {
			refv := reflect.ValueOf(dataPtrs[i])
			if r, ok := refv.Interface().(*int); ok {
				columns[i] = Column{ types[i].Name(), *r }
			}
			if r, ok := refv.Interface().(*string); ok {
				columns[i] = Column{ types[i].Name(), *r }
			}
	}
	return Row{columns}
}
