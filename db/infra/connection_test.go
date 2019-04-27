package infra

import (
	"testing"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/ShoichiroKitano/micro_test/db/table"
)

func TestDBのデータからTableオブジェクトを構築(t *testing.T) {
	defer tearDown()
	InsertIntoTest("mysql", "root:@/test_connection", "A1", "A2")
	InsertIntoTest("mysql", "root:@/test_connection", "B1", "B2")

	table := FindDBConnection("mysql", "root:@/test_connection").FindTable("test")

	expected := BuildTable("test").
							WithRow(NewColumn("column1", "A1"), NewColumn("column2", "A2")).
							WithRow(NewColumn("column1", "B1"), NewColumn("column2", "B2")).Build()
	assert.Equal(t, expected, table)
}

func Testテーブルのトランケート(t *testing.T) {
	defer tearDown()
	InsertIntoTest("mysql", "root:@/test_connection", "A1", "A2")

	connection := FindDBConnection("mysql", "root:@/test_connection")
	connection.TruncateTable("test")

	verifyTableIsEmpty(t, connection, "test")
}

func verifyTableIsEmpty(t *testing.T, connection *Connection, tableName string) {
	table := connection.FindTable(tableName)
	expected := BuildTable(tableName).Build()
	assert.Equal(t, expected, table)
}

func tearDown() {
	tx, _ := FindDBConnection("mysql", "root:@/test_connection").Driver.Begin()
	tx.Exec("truncate table test;")
	tx.Commit()
}

