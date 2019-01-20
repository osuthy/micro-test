package spec

import(
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func TestDBは検証できる(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into users (id, name, mail, age) values (1, 'userA', 'userA@gmail.com', 20);")
	con.Query("insert into users (id, name, mail, age) values (2, 'userB', 'userB@gmail.com', 21);")
	defer con.Query("truncate table users;")
	defer con.Close()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "users").ShouldHaveTable(
		db.R{"id": 1, "name": "userA", "mail": "userA@gmail.com", "age": 20},
		db.R{"id": 1, "name": "userA", "mail": "userA@gmail.com", "age": 21},
	)
	if !reflect.DeepEqual(runner.TestRunner.Result, "success") {
		t.Error("verify is fail")
	}
}
