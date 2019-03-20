package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/db/infra"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func tearDown(driver *sql.DB) {
	tx, _ := driver.Begin()
	tx.Exec("truncate table test;")
	tx.Commit()
}

func TestDBはカラムの値を正しいと判定する(t *testing.T) {
	runner.TestRunner.Result = "init"
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A1', 'A2');")
	tx.Exec("insert into test (column1, column2) values ('B1', 'B2');")
	tx.Commit()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").ShouldHaveTable(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestDBはカラムの値の誤りを検出する(t *testing.T) {
	runner.TestRunner.Result = "init"
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A1', 'A2');")
	tx.Exec("insert into test (column1, column2) values ('B1', 'B2');")
	tx.Commit()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").ShouldHaveTable(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("BUG", "B2"),
	)
	assert.Equal(t, runner.TestRunner.Result, "")
}

func TestDBはカラム順序は無視して検証する(t *testing.T) {
	runner.TestRunner.Result = "init"
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A', 'B');")
	tx.Commit()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").ShouldHaveTable(
		db.TableName("test").
		Columns("column2", "column1").
		Record("B", "A"),
	)
	assert.Equal(t, "success", runner.TestRunner.Result)
}

func TestDBは行の順序が期待値と異なる場合はテストを失敗させる(t *testing.T) {
	runner.TestRunner.Result = "init"
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A1', 'A2');")
	tx.Exec("insert into test (column1, column2) values ('B1', 'B2');")
	tx.Commit()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").ShouldHaveTable(
		db.TableName("test").
		Columns("column1", "column2").
		Record("B1", "B2").
		Record("A1", "A2"),
	)
	assert.Equal(t, "", runner.TestRunner.Result)
}
