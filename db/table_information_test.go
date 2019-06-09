package db

import (
	tbl "github.com/ShoichiroKitano/micro_test/db/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TableInformationのバリデーション
func TestTableInformationをTableに変換する(t *testing.T) {
	tableInfo := Table("name").
		Columns("c1", "c2").
		Record(11, 12).
		Record(21, 22)
	expected := tbl.BuildTable().WithName("name").
		WithRow(tbl.NewColumn("c1", 11), tbl.NewColumn("c2", 12)).
		WithRow(tbl.NewColumn("c1", 21), tbl.NewColumn("c2", 22)).Build()
	assert.Equal(t, expected, tableInfo.ToTable())
}

func TestTableInformationをTableに変換する際にカラムは名前順になる(t *testing.T) {
	tableInfo := Table("name").
		Columns("c2", "c1").
		Record(12, 11).
		Record(22, 21)
	expected := tbl.BuildTable().WithName("name").
		WithRow(tbl.NewColumn("c1", 11), tbl.NewColumn("c2", 12)).
		WithRow(tbl.NewColumn("c1", 21), tbl.NewColumn("c2", 22)).Build()
	assert.Equal(t, expected, tableInfo.ToTable())
}

func TestTableInformationからデフォルトの行データを取得する(t *testing.T) {
	assert.Equal(t, tbl.CreateRow(tbl.NewColumn("c1", "d1"), tbl.NewColumn("c2", "d2")),
		Table("").
			DefaultValue("c1", "d1").
			DefaultValue("c2", "d2").
			DefaultRow())
}

func TestTableInformationからデフォルトの行データのカラムは名前順になる(t *testing.T) {
	assert.Equal(t, tbl.CreateRow(tbl.NewColumn("c1", "d1"), tbl.NewColumn("c2", "d2")),
		Table("").
			DefaultValue("c2", "d2").
			DefaultValue("c1", "d1").
			DefaultRow())
}
