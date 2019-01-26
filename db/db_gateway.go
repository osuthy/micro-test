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

func FindTable(db *sql.DB, tableName string) Table {
	rows, _ := db.Query("SELECT * FROM " + tableName + ";")
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := make([]Row, 0)
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	return Table{newRows}
}