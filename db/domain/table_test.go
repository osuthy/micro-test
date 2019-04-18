package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type TableBuilder struct {
	tableName string
	rows []*Row
}

func BuildTable(tableName string) TableBuilder {
	return TableBuilder{tableName, []*Row{}}
}

func (builder TableBuilder) WithRow(columns ...*Column) TableBuilder {
	builder.rows = append(builder.rows, NewRow(columns))
	return builder
}

func (builder TableBuilder) Build() *Table {
	return NewTable(builder.tableName, builder.rows)
}

func Testテブールとしての等価性の判定(t *testing.T) {
	t.Run("名前と行が完全に一致", func(t *testing.T) {
		table1 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		assert.True(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの順序が不一致", func(t *testing.T) {
		table1 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable("name").
								WithRow(NewColumn("c2", 12), NewColumn("c1", 11)).
								WithRow(NewColumn("c2", 22), NewColumn("c1", 21)).Build()
		assert.True(t, table1.IsSameAsTable(table2))
	})

	t.Run("名前が不一致", func(t *testing.T) {
		table1 := BuildTable("name").Build()
		table2 := BuildTable("different name").Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの値が不一致", func(t *testing.T) {
		table1 := BuildTable("name").WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable("name").WithRow(NewColumn("c1", 100), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの数が多い", func(t *testing.T) {
		table1 := BuildTable("name").WithRow(NewColumn("c1", 1)).Build()
		table2 := BuildTable("name").WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラムの数が少ない", func(t *testing.T) {
		table1 := BuildTable("name").WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable("name").WithRow(NewColumn("c1", 1)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("カラム名が不一致", func(t *testing.T) {
		table1 := BuildTable("name").WithRow(NewColumn("c1", 1), NewColumn("c2", 2)).Build()
		table2 := BuildTable("name").WithRow(NewColumn("cA", 1), NewColumn("c2", 2)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の順序が不一致", func(t *testing.T) {
		table1 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable("name").
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の数が少ない", func(t *testing.T) {
		table1 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		table2 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})

	t.Run("行の数が多い", func(t *testing.T) {
		table1 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).Build()
		table2 := BuildTable("name").
								WithRow(NewColumn("c1", 11), NewColumn("c2", 12)).
								WithRow(NewColumn("c1", 21), NewColumn("c2", 22)).Build()
		assert.False(t, table1.IsSameAsTable(table2))
	})
}

func Test行の値で全てのカラムの値を補完する(t *testing.T) {
	result := NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
											 NewColumn("column2", "B2")})},
		  ).FilledTableWith(NewRow([]*Column{NewColumn("column3", "D3"),
																				 NewColumn("column4", "D4")}))
	assert.Equal(t, "name", result.Name)
	assert.Equal(t, []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2"),
											 NewColumn("column3", "D3"),
											 NewColumn("column4", "D4")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
											 NewColumn("column2", "B2"),
											 NewColumn("column3", "D3"),
											 NewColumn("column4", "D4")})}, result.Rows)
}

func Test補完対象の行にカラムの値がある場合は補完しない(t *testing.T) {
	result := NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
			                 NewColumn("column2", "B2")})},
		  ).FilledTableWith(NewRow([]*Column{NewColumn("column1", "D1"),
																		     NewColumn("column2", "D2"),
																		     NewColumn("column3", "D3"),}))
	assert.Equal(t, "name", result.Name)
	assert.Equal(t, []*Row{
			NewRow([]*Column{NewColumn("column1", "A1"),
											 NewColumn("column2", "A2"),
											 NewColumn("column3", "D3")}),
			NewRow([]*Column{NewColumn("column1", "B1"),
											 NewColumn("column2", "B2"),
											 NewColumn("column3", "D3")})}, result.Rows)
}
