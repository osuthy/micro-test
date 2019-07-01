package table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// デフォルト値がある場合、auto_incの場合、主キーの場合、
func Test足りないカラムの値の補完(t *testing.T) {
	t.Run("メタ情報のカラムの値が全てテーブルにある場合", func(t *testing.T){
		table :=  BuildTable().
		WithRow(NewColumn("c1", 11), NewColumn("c2", 21)).
		WithRow(NewColumn("c1", 12), NewColumn("c2", 22)).Build()
		result := NewColumnDefinition(
			[]*ColumnMetaInformation{
				NewColumnMetaInformation("c1", "int", "", "", false, false),
				NewColumnMetaInformation("c2", "int", "", "", false, false),
			}).FillTableWithDefaultValue(table)
		assert.Equal(t, table, result)
	})

	t.Run("整数の場合", func(t *testing.T) {
		result := NewColumnDefinition(
			[]*ColumnMetaInformation{
				NewColumnMetaInformation("c1", "int", "", "", false, false),
				NewColumnMetaInformation("c2", "int", "", "", false, false),
			}).FillTableWithDefaultValue(BuildTable().WithRow(NewColumn("c1", 1)).Build())
		expected := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", 0)).Build()
		assert.Equal(t, expected, result)
	})

	t.Run("文字列の場合", func(t *testing.T) {
		result := NewColumnDefinition(
			[]*ColumnMetaInformation{
				NewColumnMetaInformation("c1", "int", "", "", false, false),
				NewColumnMetaInformation("c2", "string", "", "", false, false),
			}).FillTableWithDefaultValue(BuildTable().WithRow(NewColumn("c1", 1)).Build())
		expected := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", "")).Build()
		assert.Equal(t, expected, result)
	})

	t.Run("複数の行で補完すべきカラムがある場合", func(t *testing.T) {
		result := NewColumnDefinition(
			[]*ColumnMetaInformation{
				NewColumnMetaInformation("c1", "int", "", "", false, false),
				NewColumnMetaInformation("c2", "int", "", "", false, false),
			}).FillTableWithDefaultValue(BuildTable().
			WithRow(NewColumn("c1", 11)).
			WithRow(NewColumn("c1", 21)).Build())
		expected := BuildTable().
		WithRow(NewColumn("c1", 11), NewColumn("c2", 0)).
		WithRow(NewColumn("c1", 21), NewColumn("c2", 0)).Build()
		assert.Equal(t, expected, result)
	})
}
