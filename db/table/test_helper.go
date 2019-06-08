package table

type TableBuilder struct {
	tableName string
	rows      []*Row
}

func BuildTable() TableBuilder {
	return TableBuilder{"", []*Row{}}
}

func (this TableBuilder) WithName(tableName string) TableBuilder {
	this.tableName = tableName
	return this
}

func (this TableBuilder) WithRow(columns ...*Column) TableBuilder {
	this.rows = append(this.rows, CreateRow(columns...))
	return this
}

func (this TableBuilder) Build() *Table {
	return NewTable(this.tableName, this.rows)
}

func CreateRow(columns ...*Column) *Row {
	return NewRow(columns)
}
