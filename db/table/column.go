package table

type Column struct {
	Name string
	Value interface{}
}

func NewColumn(name string, value interface{}) *Column {
	return &Column{Name: name, Value: value}
}

func (this *Column) HasSameName(other *Column) bool {
	return this.Name == other.Name
}

func (this *Column) HigherOrderThan(other *Column) bool {
	return this.Name < other.Name
}
