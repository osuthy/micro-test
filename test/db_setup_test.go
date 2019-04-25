package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/db/infra"
)

func TruncateTable(rdbms, connectionInformation, tableName string) {
	infra.FindDBConnection(rdbms, connectionInformation).TruncateTable(tableName)
}

func FindDriver(rdbms, connectionInformation string) *sql.DB {
	return infra.FindDBConnection(rdbms, connectionInformation).Driver
}

func TestDBはデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)

	rows, _ := FindDriver("mysql", "root:@/test_micro_test").Query("SELECT * FROM test;")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "A2")
	AssertNextRow(t, rows, "B1", "B2")
	AssertNextIsNone(t, rows)
}

func TestDBはデフォルト値をつかってデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").DefineDefaultValue(
		db.TableName("test").
		DefaultValue("column1", "default1").
		DefaultValue("column2", "default2"),
	)

	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1").
		Record("A1").
		Record("B1"),
	)

	rows, _ := FindDriver("mysql", "root:@/test_micro_test").Query("SELECT * FROM test;")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "default2")
	AssertNextRow(t, rows, "B1", "default2")
	AssertNextIsNone(t, rows)
}

func AssertNextRow(t *testing.T, rows *sql.Rows, columnValue1, columnValue2 string) {
	var column1, column2 string
	assert.True(t, rows.Next(), "fail next")
	rows.Scan(&column1, &column2)
	assert.Equal(t, columnValue1, column1, "fail column1")
	assert.Equal(t, columnValue2, column2, "fail column1")
}

func AssertNextIsNone(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
