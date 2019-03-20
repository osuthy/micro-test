package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/db/infra"
)

func TestDBはデータのセットアップができる(t *testing.T) {
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)

	rows, _ := driver.Query("SELECT * FROM test;")
	defer rows.Close()
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
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)

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

	rows, _ := driver.Query("SELECT * FROM test;")
	defer rows.Close()
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
