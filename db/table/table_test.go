package table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testテーブルとしての等価性の判定(t *testing.T) {
	t.Run("名前と行が完全に一致", func(t *testing.T) {
		table1 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		assert.True(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの順序が不一致", func(t *testing.T) {
		table1 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable().
			WithRow(NewColumn("c2", 12), NewColumn("c1", 11)).
			WithRow(NewColumn("c2", 22), NewColumn("c1", 21)).Build()
		assert.True(t, table1.IsSameAsTable(table2))
	})

	t.Run("名前が不一致", func(t *testing.T) {
		table1 := BuildTable().WithName("name").Build()
		table2 := BuildTable().WithName("different name").Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの値が不一致", func(t *testing.T) {
		table1 := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable().WithRow(NewColumn("c1", 100), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの数が多い", func(t *testing.T) {
		table1 := BuildTable().WithRow(NewColumn("c1", 1)).Build()
		table2 := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの数が少ない", func(t *testing.T) {
		table1 := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable().WithRow(NewColumn("c1", 1)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラム名が不一致", func(t *testing.T) {
		table1 := BuildTable().WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable().WithRow(NewColumn("cA", 1), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の順序が不一致", func(t *testing.T) {
		table1 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable().
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の数が少ない", func(t *testing.T) {
		table1 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の数が多い", func(t *testing.T) {
		table1 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		table2 := BuildTable().
			WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
			WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})
}

func Test行によるカラムの補完(t *testing.T) {
	t.Run("テーブルに存在していないカラムを行から補完する", func(t *testing.T) {
		table := BuildTable().
			WithRow(NewColumn("c1", "A1"), NewColumn("c2", "A2")).
			WithRow(NewColumn("c1", "B1"), NewColumn("c2", "B2")).Build()
		row := CreateRow(NewColumn("c3", "D3"), NewColumn("c4", "D4"))

		expected := BuildTable().
			WithRow(NewColumn("c1", "A1"), NewColumn("c2", "A2"), NewColumn("c3", "D3"), NewColumn("c4", "D4")).
			WithRow(NewColumn("c1", "B1"), NewColumn("c2", "B2"), NewColumn("c3", "D3"), NewColumn("c4", "D4")).
			Build()
		assert.Equal(t, expected, table.FilledTableWith(row))
	})

	t.Run("行が持つカラムをテーブルが持っているなら補完しない", func(t *testing.T) {
		table := BuildTable().
			WithRow(NewColumn("c1", "A1"), NewColumn("c2", "A2")).
			WithRow(NewColumn("c1", "B1"), NewColumn("c2", "B2")).Build()
		row := CreateRow(NewColumn("c1", "not completed1"), NewColumn("c2", "not completed2"), NewColumn("c3", "completed"))

		expected := BuildTable().
			WithRow(NewColumn("c1", "A1"), NewColumn("c2", "A2"), NewColumn("c3", "completed")).
			WithRow(NewColumn("c1", "B1"), NewColumn("c2", "B2"), NewColumn("c3", "completed")).Build()
		assert.Equal(t, expected, table.FilledTableWith(row))
	})
}
