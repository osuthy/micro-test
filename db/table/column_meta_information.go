package table

import (
	"time"
)

type ColumnMetaInformation struct {
	name string
	dataType string
	preferredKeyConstraint string
	additional string
	isNotNull bool
	hasDefaultValue bool
}

func NewColumnMetaInformation(
	name string,
	dataType string,
	preferredKeyConstraint string,
	additional string,
	isNotNull bool,
	hasDefaultValue bool) *ColumnMetaInformation {
	return &ColumnMetaInformation {
		name: name,
		dataType: dataType,
		preferredKeyConstraint: preferredKeyConstraint,
		additional: additional,
		hasDefaultValue: hasDefaultValue,
		isNotNull: isNotNull,
	}
}

// mysql専用にした方が良い？
func (this *ColumnMetaInformation) defaultColumn() *Column {
	var value interface{}
	if this.dataType == "string" {
		value = ""
	} else if this.dataType == "int" {
		value = 0
	} else if this.dataType == "time" {
		value = time.Unix(0, 0).Add(time.Hour * 10) //timestampがエラーを吐くので10時間追加
	} else if this.dataType == "year" {
		value = "1901"
	} else if this.dataType == "decimal" {
		value = 0.0
	} else if this.dataType == "binary" {
		value = 0x0
	}
	return NewColumn(this.name, value)
}
