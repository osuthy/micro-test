package test

import (
	"testing"

	. "github.com/ShoichiroKitano/micro_test/db/infra"

	"github.com/ShoichiroKitano/micro_test/db"
)

func TestDBはデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	db.ResetDefalultValue()

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

func TestDBはデフォルト値をつかってMysqlのデータのセットアップができる(t *testing.T) {
	defer TruncateTable("mysql", "root:@/test_micro_test", "test")
	db.ResetDefalultValue()

	db.DefineConnection("conName", "mysql", "root:@/test_micro_test")
	db.DB("conName").HasRecords(
		db.Table("test").
			Columns("column1").
			Record("A1").
			Record("B1"),
	)

	rows := Select("mysql", "root:@/test_micro_test", "test")
	defer rows.Close()
	AssertNextRow(t, rows, "A1", "")
	AssertNextRow(t, rows, "B1", "")
	AssertNextIsNone(t, rows)
}

