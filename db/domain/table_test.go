package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//テーブル名が異なる場合

func TestRow2(t *testing.T) {
	result := NewTable(
		"name",
		[]*Row{NewRow([]*Column{NewColumn("column", 1)}), NewRow([]*Column{NewColumn("column", 2)})},
	).IsSame(NewTable(
		"name",
		[]*Row{NewRow([]*Column{NewColumn("column", 2)}), NewRow([]*Column{NewColumn("column", 1)})},
	))
	assert.False(t, result)
}

func TestRow(t *testing.T) {
	result := NewTable(
		"name",
		[]*Row{NewRow([]*Column{NewColumn("column", 1)}), NewRow([]*Column{NewColumn("column", 2)})},
	).IsSame(NewTable(
		"name",
		[]*Row{NewRow([]*Column{NewColumn("column", 1)}), NewRow([]*Column{NewColumn("column", 2)})},
	))
	assert.True(t, result)
}

func TestIgnoreOrder(t *testing.T) {
	result := NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", 1),
											 NewColumn("column2", 2)})},
		).IsSame(NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column2", 2),
											 NewColumn("column1", 1)})}))
	assert.True(t, result)
}

func TestIsSameTrue(t *testing.T) {
	result := NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", 1),
											 NewColumn("column2", 2)})},
		).IsSame(NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", 1),
											 NewColumn("column2", 2)})}))
	assert.True(t, result)
}

func TestIsSameFalse(t *testing.T) {
	result := NewTable("name", []*Row{
			NewRow([]*Column{NewColumn("column1", 1),
											 NewColumn("column2", 2)})},
	    ).IsSame(NewTable("name", []*Row{
				NewRow([]*Column{NewColumn("column1", 100),
												 NewColumn("column2", 2)})}))
	assert.False(t, result)
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
