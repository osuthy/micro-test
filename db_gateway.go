package db_gateway

import (
	"database/sql"
	"reflect"
	"fmt"
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

func FindTable(db *sql.DB, tableName string) Table {
	rows, _ := db.Query("SELECT * FROM users;")
	types, _ := rows.ColumnTypes()
	dataPtrs := make([]interface{}, len(types))

	for i := range types {
		if types[i].DatabaseTypeName() == "INT" {
			dataPtrs[i] = new(int)
		} else if types[i].DatabaseTypeName() == "CHAR" {
			dataPtrs[i] = new(string)
		}
	}

	rows.Next()
	rows.Scan(dataPtrs...)
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
	fmt.Println(columns)
	return Table{
								[]Row{
										Row{
											columns,
										},
								},
							}
}
