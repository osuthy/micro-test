package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
	. "github.com/ShoichiroKitano/micro_test/db/domain"
)

// TableInformationのバリデーション
func TestTableInformationをTableに変換する(t *testing.T) {
	tableInfo := TableName("name").
	Columns("column1", "column2").
	Record("A1", "A2").
	Record("B1", "B2")
	assert.Equal(t, tableInfo.ToTable(), Table{
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
	}, tableInfo.ToTable())
}

func TestTableInformationをTableに変換する際にカラムは名前順になる(t *testing.T) {
	tableInfo := TableName("name").
	Columns("column2", "column1").
	Record("A2", "A1").
	Record("B2", "B1")
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
	}, tableInfo.ToTable())
}

func TestTableInformationからデフォルトの行データを取得する(t *testing.T) {
	assert.Equal(t, Row{[]Column{Column{"column1", "d1"}, {"column2", "d2"}}},
	TableName("name").
		DefaultValue("column1", "d1").
		DefaultValue("column2", "d2").
		DefaultRow())
}

func TestTableInformationからデフォルトの行データのカラムは名前順になる(t *testing.T) {
	assert.Equal(t, Row{[]Column{Column{"column1", "d1"}, {"column2", "d2"}}},
	TableName("name").
		DefaultValue("column2", "d2").
		DefaultValue("column1", "d1").
		DefaultRow())
}
