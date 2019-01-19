package db_gateway

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func TestTable(t *testing.T) {
	db, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.Query("insert into users (id, name, mail, age) values (1, 'user', 'user@gmail.com', 20);")
	table := FindTable(db, "users")
	if !reflect.DeepEqual(table, Table{
		[]Row{
			Row{
				[]Column{
					Column{"id", 1},
					Column{"name", "user"},
					Column{"mail", "user@gmail.com"},
					Column{"age", 20},
				},
			},
		},
	}) {
		t.Error("table not create")
	}
	db.Query("truncate table users;")
	db.Close()
}
