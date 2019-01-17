package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
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

type R map[string]interface{}

func findRows() Rows {
}

func main() {
	db, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.Query("insert into users (name, mail, age) values ('userA', 'userA@gmail.com', '20');")
	db.Query("insert into users (name, mail, age) values ('userB', 'userB@gmail.com', '21');")
	rows, _ := db.Query("SELECT * FROM users;")
	expectedRows := []Row{
		Row{
			[]Column{
				Column{"id", 1},
				Column{"name", "userA"},
				Column{"mail", "userA@gmail.com"},
				Column{"age", 20},
			},
		},
		Row{
			[]Column{
				Column{"id", 2},
				Column{"name", "userB"},
				Column{"mail", "userB@gmail.com"},
				Column{"age", 21},
			},
		},
	}
	fmt.Println(expectedRows)
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

