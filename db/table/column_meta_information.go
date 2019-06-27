package table

type ColumnMetaInformation struct {
	name string
	dataType string
	isNullable bool
	preferredKeyConstraint string
	defaultValue interface{}
	additional string
}

func NewColumnMetaInformation(
	name string,
	dataType string,
	preferredKeyConstraint string,
	isNullable bool,
	defaultValue interface{},
	additional string) *ColumnMetaInformation {
	return &ColumnMetaInformation {
		name: name,
		column: column,
		dataType: dataType,
		preferredKeyConstraint: preferredKeyConstraint,
		defaultValue: defaultValue,
		isNullable: isNullable,
		additional: additional,
	}
}
