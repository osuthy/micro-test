package db

import (
	"sort"
	. "github.com/ShoichiroKitano/micro_test/db/table"
)

type TableInformation struct {
	tableName string
	columnNames []string
	records [][]interface{}
	defaultRecord map[string]interface{}
}

func TableName(tableName string) TableInformation {
	return TableInformation{
		tableName: tableName,
		records: make([][]interface{}, 0),
		defaultRecord: make(map[string]interface{})}
}

func (this TableInformation) Columns(columnNames...string) TableInformation {
	this.columnNames = columnNames
	return this
}

func (this TableInformation) Record(values...interface{}) TableInformation {
	this.records = append(this.records, values)
	return this
}

func (this TableInformation) DefaultValue(columnName string, value interface{}) TableInformation {
	this.defaultRecord[columnName] = value
	return this
}

func (this TableInformation) DefaultRow() *Row {
	columns := make([]*Column, 0)
	for columnName, value := range this.defaultRecord {
		columns = append(columns, NewColumn(columnName, value))
	}
	sort.Slice(columns, func(i, j int) bool { return columns[i].Name < columns[j].Name })
	return NewRow(columns)
}

func (this TableInformation) ToTable() *Table {
	rows := make([]*Row, 0)
	for _, record := range this.records {
		columns := make([]*Column, 0)
		for i, name := range this.columnNames {
			columns = append(columns, NewColumn(name, record[i]))
		}
		sort.Slice(columns, func(i, j int) bool { return columns[i].Name < columns[j].Name })
		rows = append(rows, NewRow(columns))
	}
	return NewTable(this.tableName, rows)
}

