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
	con.Query("truncate table test;")
	con.Close()
}

func TestDBはカラムの値を正しいと判定する(t *testing.T) {
	runner.TestRunner.Result = "init"
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into test (column1, column2) values ('A1', 'A2');")
	con.Query("insert into test (column1, column2) values ('B1', 'B2');")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "test").ShouldHaveTable(
		db.R{"column1": "A1", "column2": "A2"},
		db.R{"column1": "B1", "column2": "B2"},
	)
	assert.Equal(t, runner.TestRunner.Result, "success")
}

func TestDBはカラムの値の誤りを検出する(t *testing.T) {
	runner.TestRunner.Result = "init"
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into test (column1, column2) values ('A1', 'A2');")
	con.Query("insert into test (column1, column2) values ('B1', 'B2');")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "test").ShouldHaveTable(
		db.R{"column1": "A1", "column2": "A2"},
		db.R{"column1": "fail", "column2": "B2"},
	)
	assert.Equal(t, runner.TestRunner.Result, "")
}

func TestDBはカラム順序は無視して検証する(t *testing.T) {
	runner.TestRunner.Result = "init"
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into test (column1, column2) values ('1', '2');")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "test").ShouldHaveTable(
		db.R{"column2": "2", "column1": "1"},
	)
	assert.Equal(t, runner.TestRunner.Result, "success")
}

func TestDBは行の順序が期待値と異なる場合はテストを失敗させる(t *testing.T) {
	runner.TestRunner.Result = "init"
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into test (column1, column2) values ('A1', 'A2');")
	con.Query("insert into test (column1, column2) values ('B1', 'B2');")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName", "test").ShouldHaveTable(
		db.R{"column1": "B1", "column2": "B2"},
		db.R{"column1": "A1", "column2": "A2"},
	)
	assert.Equal(t, runner.TestRunner.Result, "")
}
