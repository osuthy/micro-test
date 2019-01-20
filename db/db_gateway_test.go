package db

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

func TestFinedTable(t *testing.T) {
	db, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.Query("insert into users (id, name, mail, age) values (1, 'userA', 'userA@gmail.com', 20);")
	db.Query("insert into users (id, name, mail, age) values (2, 'userB', 'userB@gmail.com', 21);")
	defer db.Query("truncate table users;")
	defer db.Close()

	table := FindTable(db, "users")
	if !reflect.DeepEqual(table, Table{
		[]Row{
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
		},
	}) {
		t.Error("table not create")
	}
}
