package infra

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/osuthy/micro-test/db/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

func InsertIntoTest(rdbms, connectionInformation, column1Value, column2Value string) {
	tx, _ := FindDBConnection(rdbms, connectionInformation).Driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('" + column1Value + "', '" + column2Value + "');")
	tx.Commit()
}

func TruncateTable(rdbms, connectionInformation, tableName string) {
	FindDBConnection(rdbms, connectionInformation).TruncateTable(NewTable(tableName, nil))
}

func Select(rdbms, connectionInformation, tableName string) *sql.Rows {
	rows, _ := FindDBConnection(rdbms, connectionInformation).Driver.Query("SELECT * FROM " + tableName + ";")
	return rows
}

func AssertNextRow(t *testing.T, rows *sql.Rows, columnValue1, columnValue2 string) {
	var column1, column2 string
	assert.True(t, rows.Next(), "fail next")
	rows.Scan(&column1, &column2)
	assert.Equal(t, columnValue1, column1, "fail column1")
	assert.Equal(t, columnValue2, column2, "fail column2")
}

func AssertNextIsNone(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
