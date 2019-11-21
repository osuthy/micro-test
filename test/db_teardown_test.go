package test

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/osuthy/micro-test/db"
	. "github.com/osuthy/micro-test/db/infra"
	. "github.com/osuthy/micro-test"
)

func TestDBは指定したテーブルをtruncateできる(t *testing.T) {
	t.Skip()
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	DB(TC{}, "conName").HasRecords(
		Table("test").
			Columns("column1", "column2").
			Record("A1", "A2").
			Record("B1", "B2"),
	)

	DB(TC{}, "conName").Truncate("test")

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertEmpty(t, rows)
}

func AssertEmpty(t *testing.T, rows *sql.Rows) {
	assert.False(t, rows.Next())
}
