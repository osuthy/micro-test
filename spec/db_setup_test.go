package spec

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
	//defer tearDown(con)

	db.DB("conName").HasRecords(
		db.TableName("test").
		Columns("column1", "column2").
		R("A1", "A2").
		R("B1", "B2"),
	)

	rows, _ := con.Query("SELECT * FROM test;")
	var colmun1 string
	var colmun2 string

	assert.True(t, rows.Next())
	rows.Scan(&colmun1, &colmun2)
	assert.Equal(t, colmun1, "A1")
	assert.Equal(t, colmun2, "A2")

	assert.True(t, rows.Next())
	rows.Scan(&colmun1, &colmun2)
	assert.Equal(t, colmun1, "B1")
	assert.Equal(t, colmun2, "B2")

	assert.False(t, rows.Next())
}
