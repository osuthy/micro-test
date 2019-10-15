package infra

import (
	. "github.com/osuthy/micro-test/db/table"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDBのデータからTableオブジェクトを構築(t *testing.T) {
	tearDown()
	defer tearDown()
	InsertIntoTest("mysql", "root:@/test_connection", "A1", "A2")
	InsertIntoTest("mysql", "root:@/test_connection", "B1", "B2")

	table := FindDBConnection("mysql", "root:@/test_connection").FindTable(BuildTable().WithName("test").Build())

	expected := BuildTable().WithName("test").
		WithRow(NewColumn("column1", "A1"), NewColumn("column2", "A2")).
		WithRow(NewColumn("column1", "B1"), NewColumn("column2", "B2")).Build()
	assert.Equal(t, expected, table)
}

func Testテーブルのトランケート(t *testing.T) {
	tearDown()
	defer tearDown()
	InsertIntoTest("mysql", "root:@/test_connection", "A1", "A2")

	connection := FindDBConnection("mysql", "root:@/test_connection")
	connection.TruncateTable(BuildTable().WithName("test").Build())

	verifyTableIsEmpty(t, connection, "test")
}

func verifyTableIsEmpty(t *testing.T, connection *Connection, tableName string) {
	expected := BuildTable().WithName(tableName).Build()
	assert.Equal(t, expected, connection.FindTable(expected))
}

func tearDown() {
	tx, _ := FindDBConnection("mysql", "root:@/test_connection").Driver.Begin()
	tx.Exec("truncate table test;")
	tx.Commit()
}
