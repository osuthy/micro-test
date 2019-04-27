package table

type Column struct {
	Name string
	Value interface{}
}

func NewColumn(name string, value interface{}) *Column {
	column := new(Column)
	column.Name = name
	column.Value = value
	return column
}

func (this *Column) HasSameName(other *Column) bool {
	return this.Name == other.Name
}

func (this *Column) HigherOrderThan(other *Column) bool {
	return this.Name < other.Name
}
