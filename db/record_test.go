package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TableInformationのバリデーション
func TestTableInformationをTableに変換する(t *testing.T) {
	tableInfo := TableName("name").
	Columns("column1", "column2").
	R("A1", "A2").
	R("B1", "B2")
	assert.Equal(t, tableInfo.toTable(), Table{
		"name",
		[]Row {
			Row{
				[]Column{
					Column{"column1", "A1"}, Column{"column2", "A2"},
				},
			},
			Row{
				[]Column{
					Column{"column1", "B1"}, Column{"column2", "B2"},
				},
			},
		},
	}, tableInfo.toTable())
}

func TestTableInformationをTableに変換する際にカラムは名前順になる(t *testing.T) {
	tableInfo := TableName("name").
	Columns("column2", "column1").
	R("A2", "A1").
	R("B2", "B1")
	assert.Equal(t, Table{
		"name",
		[]Row {
			Row{
				[]Column{
					Column{"column1", "A1"}, Column{"column2", "A2"},
				},
			},
			Row{
				[]Column{
					Column{"column1", "B1"}, Column{"column2", "B2"},
				},
			},
		},
	}, tableInfo.toTable())
}

func TestTableInformationからデフォルトの行データを取得する(t *testing.T) {
	assert.Equal(t, Row{[]Column{Column{"column1", "d1"}, {"column2", "d2"}}},
	TableName("name").
		DefaultValue("column1", "d1").
		DefaultValue("column2", "d2").
		defaultRow())
}

func TestTableInformationからデフォルトの行データのカラムは名前順になる(t *testing.T) {
	assert.Equal(t, Row{[]Column{Column{"column1", "d1"}, {"column2", "d2"}}},
	TableName("name").
		DefaultValue("column2", "d2").
		DefaultValue("column1", "d1").
		defaultRow())
}
