package test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/ShoichiroKitano/micro_test/db/infra"

	"github.com/ShoichiroKitano/micro_test/db"
)

func TestDBは指定したテーブルをtruncateできる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.TableName("test").
			Columns("column1", "column2").
			Record("A1", "A2").
			Record("B1", "B2"),
	)

	db.DB("conName").Truncate("test")

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertEmpty(t, rows)
}

func AssertEmpty(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
