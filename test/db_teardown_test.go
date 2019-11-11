package test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/osuthy/micro-test/db"
	. "github.com/osuthy/micro-test/db/infra"
	. "github.com/osuthy/micro-test/testable"
)

func TestDBは指定したテーブルをtruncateできる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	DB(TestContext{}, "conName").HasRecords(
		Table("test").
			Columns("column1", "column2").
			Record("A1", "A2").
			Record("B1", "B2"),
	)

	DB(TestContext{}, "conName").Truncate("test")

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertEmpty(t, rows)
}

func AssertEmpty(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
