package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TableInformationのバリデーション
func Test複数のRecordをTableに変換する(t *testing.T) {
	tableInfo := Columns("column1", "column2").
	R("A1", "A2").
	R("B1", "B2")
	assert.Equal(t, tableInfo.toTable(), Table{
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
	})
}
