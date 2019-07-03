package test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/ShoichiroKitano/micro_test/db/infra"

	"github.com/ShoichiroKitano/micro_test/db"
)

func TestDBはデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.Table("test").
			Columns("column1", "column2").
			Record("A1", "A2").
			Record("B1", "B2"),
	)

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "A2")
	AssertNextRow(t, rows, "B1", "B2")
	AssertNextIsNone(t, rows)
}

func Test事前条件のデータの補完(t *testing.T) {
	TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")
	defer TruncateTable("mysql", "root:@/test_micro_test", "record_completion_all_type")

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.Table("record_completion_all_type").
			Columns("dummy").
			Record("dummy1").
			Record("dummy2"),
	)

	rows := Select("mysql", "root:@/test_micro_test", "record_completion_all_type")
	defer rows.Close()
	var (
		dummy string
		tinyintc int
		intc int
		datec string
	)
	rows.Next()
	rows.Scan(&dummy, &tinyintc, &intc, &datec)
	assert.Equal(t, "dummy1", dummy)
	assert.Equal(t, 0, tinyintc)
	assert.Equal(t, 0, intc)
	assert.Equal(t, "1970-01-01", datec)

	rows.Next()
	rows.Scan(&dummy, &tinyintc, &intc, &datec)
	assert.Equal(t, "dummy2", dummy)
	assert.Equal(t, 0, tinyintc)
	assert.Equal(t, 0, intc)
	assert.Equal(t, "1970-01-01", datec)

	AssertNextIsNone(t, rows)
}

