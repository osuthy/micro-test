package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
	"github.com/ShoichiroKitano/micro_test/db/infra"
)

func TestDBは指定したテーブルをtruncateできる(t *testing.T) {
	driver := infra.FindDBConnection("mysql", "root:@/test_micro_test").Driver
	defer tearDown(driver)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)

	db.DB("conName").Truncate("test")

	rows, _ := driver.Query("SELECT * FROM test;")

	defer rows.Close()
	assert.False(t, rows.Next())
}
