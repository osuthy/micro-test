package infra

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	. "github.com/ShoichiroKitano/micro_test/db/domain"
)

func TestFinedTable(t *testing.T) {
	driver, _ := sql.Open("mysql", "root:@/test_micro_test")
	driver.Query("insert into test (column1, column2) values ('A1', 'A2');")
	driver.Query("insert into test (column1, column2) values ('B1', 'B2');")
	defer driver.Query("truncate table test;")
	defer driver.Close()
	table := NewConnection(driver).FindTable("test")
	assert.Equal(t, table.Name, "test")
	assert.Equal(t, table.Rows, []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
											 NewColumn("column2", "B2")})})
}

func TestTruncateTable(t *testing.T) {
	driver, _ := sql.Open("mysql", "root:@/test_micro_test")
	driver.Query("insert into test (column1, column2) values ('A1', 'A2');")
	defer driver.Query("truncate table test;")
	defer driver.Close()
	connection := NewConnection(driver)
	connection.TruncateTable("test")
	table := connection.FindTable("test")
	assert.Equal(t, table.Name, "test")
	assert.Equal(t, table.Rows, []*Row{})
}
