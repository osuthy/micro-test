package infra

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/ShoichiroKitano/micro_test/db/domain"
)

func tearDown(driver *sql.DB) {
	tx, _ := driver.Begin()
	tx.Exec("truncate table test;")
	tx.Commit()
}

func TestFinedTable(t *testing.T) {
	driver := FindDBConnection("mysql", "root:@/test_connection").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A1', 'A2');")
	tx.Exec("insert into test (column1, column2) values ('B1', 'B2');")
	tx.Commit()

	table := NewConnection(driver).FindTable("test")
	assert.Equal(t, table.Name, "test")
	assert.Equal(t, table.Rows, []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
											 NewColumn("column2", "B2")})})
}

func TestTruncateTable(t *testing.T) {
	driver := FindDBConnection("mysql", "root:@/test_connection").Driver
	defer tearDown(driver)
	tx, _ := driver.Begin()
	tx.Exec("insert into test (column1, column2) values ('A1', 'A2');")
	tx.Commit()

	connection := NewConnection(driver)
	connection.TruncateTable("test")
	table := connection.FindTable("test")
	assert.Equal(t, table.Name, "test")
	assert.Equal(t, table.Rows, []*Row{})
}
