package table

type ColumnMetaInformation struct {
	name string
	dataType string
	preferredKeyConstraint string
	additional string
	isNullable bool
	hasDefaultValue bool
}

func NewColumnMetaInformation(
	name string,
	dataType string,
	preferredKeyConstraint string,
	additional string,
	isNullable bool,
	hasDefaultValue bool) *ColumnMetaInformation {
	return &ColumnMetaInformation {
		name: name,
		dataType: dataType,
		preferredKeyConstraint: preferredKeyConstraint,
		additional: additional,
		hasDefaultValue: hasDefaultValue,
		isNullable: isNullable,
	}
}

func (this *ColumnMetaInformation) defaultColumn() *Column {
	var value interface{}
	if this.dataType == "string" {
		value = ""
	} else if this.dataType == "int" {
		value = 0
	}
	return NewColumn(this.name, value)
}
