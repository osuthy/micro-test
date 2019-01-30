package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
)
// 自動採番、テンプレート、時間

func TestDBはデータのセットアップができる(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	defer tearDown(con)

	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)

	rows, _ := con.Query("SELECT * FROM test;")
	var column1 string
	var column2 string

	assert.True(t, rows.Next())
	rows.Scan(&column1, &column2)
	assert.Equal(t, "A1", column1)
	assert.Equal(t, "A2", column2)

	assert.True(t, rows.Next())
	rows.Scan(&column1, &column2)
	assert.Equal(t, "B1", column1)
	assert.Equal(t, "B2", column2)

	assert.False(t, rows.Next())
}

func TestDBはデフォルト値をつかってデータのセットアップができる(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").SetDefaultValue(
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

	rows, _ := con.Query("SELECT * FROM test;")
	var column1 string
	var column2 string

	assert.True(t, rows.Next())
	rows.Scan(&column1, &column2)
	assert.Equal(t, "A1", column1)
	assert.Equal(t, "default2", column2)

	assert.True(t, rows.Next())
	rows.Scan(&column1, &column2)
	assert.Equal(t, "B1", column1)
	assert.Equal(t, "default2", column2)

	assert.False(t, rows.Next())
}
