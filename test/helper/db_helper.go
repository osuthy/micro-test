package helper

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ShoichiroKitano/micro_test/db/infra"
)

func TruncateTable(rdbms, connectionInformation, tableName string) {
	infra.FindDBConnection(rdbms, connectionInformation).TruncateTable(tableName)
}

func Select(rdbms, connectionInformation, tableName string) *sql.Rows {
	rows, _ := infra.FindDBConnection(rdbms, connectionInformation).Driver.Query("SELECT * FROM " + tableName + ";")
	return rows
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
