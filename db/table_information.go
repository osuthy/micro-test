package db

import (
	tbl "github.com/osuthy/micro-test/db/table"
)

type TableInformation struct {
	tableName   string
	columnNames []string
	records     [][]interface{}
}

func Table(tableName string) TableInformation {
	return TableInformation{
		tableName: tableName,
		records:   [][]interface{}{},
	}
}

func (this TableInformation) Columns(columnNames ...string) TableInformation {
	this.columnNames = columnNames
	return this
}

func (this TableInformation) Record(values ...interface{}) TableInformation {
	this.records = append(this.records, values)
	return this
}

func (this TableInformation) ToTable() *tbl.Table {
	rows := make([]*tbl.Row, 0, len(this.records))
	for _, record := range this.records {
		columns := make([]*tbl.Column, 0, len(this.columnNames))
		for i, name := range this.columnNames {
			columns = append(columns, tbl.NewColumn(name, record[i]))
		}
		rows = append(rows, tbl.NewRow(columns).Sorted())
	}
	return tbl.NewTable(this.tableName, rows)
}
