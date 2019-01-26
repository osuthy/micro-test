package spec

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func tearDown(con *sql.DB) {
	con.Query("truncate table users;")
	con.Close()
}

func TestDBは検証できる(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into users (id, name, mail, age) values (1, 'userA', 'userA@gmail.com', 20);")
	con.Query("insert into users (id, name, mail, age) values (2, 'userB', 'userB@gmail.com', 21);")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "users").ShouldHaveTable(
		db.R{"id": 1, "name": "userA", "mail": "userA@gmail.com", "age": 20},
		db.R{"id": 2, "name": "userB", "mail": "userB@gmail.com", "age": 21},
	)
	assert.Equal(t, runner.TestRunner.Result, "success")
}

func TestDBは検証できる2(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into users (id, name, mail, age) values (1, 'userA', 'userA@gmail.com', 20);")
	con.Query("insert into users (id, name, mail, age) values (2, 'userB', 'userB@gmail.com', 21);")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "users").ShouldHaveTable(
		db.R{"id": 1, "name": "userA", "mail": "userA@gmail.com", "age": 20},
		db.R{"id": 2, "name": "userC", "mail": "userB@gmail.com", "age": 21},
	)
	assert.Equal(t, runner.TestRunner.Result, "")
}
