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
	rows, _ := db.Query("SELECT * FROM users;")
	dataPtrs := make([]interface{}, 4)
	for rows.Next() {
		types, _ := rows.ColumnTypes()
		for i := range types {
			fmt.Println(types[i].DatabaseTypeName())
		}
		dataPtrs[0] = new(int)
		//dataPtrs[1] = new(string)
		//dataPtrs[2] = new(string)
		//dataPtrs[3] = new(int)
		rows.Scan(dataPtrs...)
		refv := reflect.ValueOf(dataPtrs[0])
		fmt.Println(refv.Kind())
		if r, ok := refv.Interface().(*int); ok {
			fmt.Println(*r)
		}
	}
	db.Query("truncate table users;")
}

