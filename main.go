package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
)
type R map[string]interface{}

func main() {
	db, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.Query("insert into users (name, mail, age) values ('userA', 'userA@gmail.com', '20');")
	db.Query("insert into users (name, mail, age) values ('userB', 'userB@gmail.com', '21');")
	rows, _ := db.Query("SELECT * FROM users;")
	dataPtrs := make([]interface{}, 4)
	for rows.Next() {
		types, _ := rows.ColumnTypes()
		for i := range types {
			if types[i].DatabaseTypeName() == "INT" {
				dataPtrs[i] = new(int)
			} else if types[i].DatabaseTypeName() == "CHAR" {
				dataPtrs[i] = new(string)
			}
		}
		rows.Scan(dataPtrs...)
		for i := range types {
			refv := reflect.ValueOf(dataPtrs[i])
			fmt.Println(refv.Kind())
			if r, ok := refv.Interface().(*int); ok {
				fmt.Println(*r)
			}
			if r, ok := refv.Interface().(*string); ok {
				fmt.Println(*r)
			}
		}
	}
	db.Query("truncate table users;")
}

