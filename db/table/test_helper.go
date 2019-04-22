package table

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
	builder.rows = append(builder.rows, CreateRow(columns...))
	return builder
}

func (builder TableBuilder) Build() *Table {
	return NewTable(builder.tableName, builder.rows)
}

func CreateRow(columns ...*Column) *Row {
	return NewRow(columns)
}

