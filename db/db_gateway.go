package db

import (
	"database/sql"
	"reflect"
)

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
	for _, row := range table.rows {
		this.driver.Query("insert into test (" + row.columns[0].name + "," + row.columns[1].name + ") values (" + toLiteral(row.columns[0]) + "," + toLiteral(row.columns[1]) + ");")
	}
}

func toLiteral(column Column) string {
	refv := reflect.ValueOf(column.value)
	r, _ := refv.Interface().(string)
	return "'" + r +"'"
}

