package test

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/db"
)

func TestDBは指定したテーブルをtruncateできる(t *testing.T) {
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	defer tearDown(con)

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		Record("A1", "A2").
		Record("B1", "B2"),
	)

	db.DB("conName").Truncate("test")

	rows, _ := con.Query("SELECT * FROM test;")
	assert.False(t, rows.Next())
}
