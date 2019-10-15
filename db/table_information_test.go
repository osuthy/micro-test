package db

import (
	tbl "github.com/osuthy/micro-test/db/table"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
