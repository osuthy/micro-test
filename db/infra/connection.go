package infra

import (
	"database/sql"
	. "github.com/ShoichiroKitano/micro_test/db/table"
	"reflect"
	"strings"
)

type Connection struct {
	Driver *sql.DB
	rdbms string
}

func NewConnection(driver *sql.DB, rdbms string) *Connection {
	return &Connection{Driver: driver, rdbms: rdbms}
}

func (this *Connection) FindTable(table *Table) *Table {
	rows, _ := this.Driver.Query(table.SelectAllQuery())
	defer rows.Close()
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := []*Row{}
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	return NewTable(table.Name(), newRows)
}

func (this *Connection) FindColumnDefinition(table *Table) *ColumnDefinition {
	return this.mysqlColumnDefinition(table)
}

func (this *Connection) StoreTable(table *Table) {
	tx, _ := this.Driver.Begin()
	stmt, _ := tx.Prepare(table.InsertQuery())
	p(table.InsertQuery())
	for _, values := range table.AllValues() {
		_, err :=  stmt.Exec(values...)
		p(err)
	}
	tx.Commit()
}

func (this *Connection) TruncateTable(table *Table) {
	this.Driver.Exec(table.TruncateQuery())
}

func rowFrom(types []*sql.ColumnType, dataPtrs []interface{}) *Row {
	columns := make([]*Column, len(types))
	for i := range types {
		refv := reflect.ValueOf(dataPtrs[i])
		if r, ok := refv.Interface().(*int); ok {
			columns[i] = NewColumn(types[i].Name(), *r)
		}
		if r, ok := refv.Interface().(*string); ok {
			columns[i] = NewColumn(types[i].Name(), *r)
		}
		if r, ok := refv.Interface().(*sql.NullString); ok {
			columns[i] = NewColumn(types[i].Name(), r.String)
		}
	}
	return NewRow(columns)
}

// for mysql
func (this *Connection) mysqlColumnDefinition(table *Table) *ColumnDefinition {
	rows, _ := this.Driver.Query(table.MysqlColumnDefinitionQuery())
	defer rows.Close()
	types, _ := rows.ColumnTypes()
	dataPtrs := scanArgs(types)
	newRows := []*Row{}
	for rows.Next() {
		rows.Scan(dataPtrs...)
		newRows = append(newRows, rowFrom(types, dataPtrs))
	}
	infos := []*ColumnMetaInformation{}
	for _, row := range newRows {
		infos = append(infos, NewColumnMetaInformation(
			row.ColumnValueOf("Field").(string),
			toDataTypeNameForMysql(row.ColumnValueOf("Type").(string)),
			"",
			"",
			true,
			false,
			//row.ColumnValueOf("Key").(string),
			//row.ColumnValueOf("Extra").(string),
			//isNotNullFlag(row.ColumnValueOf("Null").(string)),
			//row.ColumnValueOf("Default"),
		))
	}
	return NewColumnDefinition(infos)
}

func toDataTypeNameForMysql(mysqlDataType string) string {
	p(mysqlDataType)
	if strings.Contains(mysqlDataType, "int") {
		return "int"
	} else if strings.Contains(mysqlDataType, "char") || strings.Contains(mysqlDataType, "text") {
		return "string"
	} else if strings.Contains(mysqlDataType, "year") {
		return "year"
	} else if strings.Contains(mysqlDataType, "float") ||
	strings.Contains(mysqlDataType, "double") ||
	strings.Contains(mysqlDataType, "decimal") {
		return "decimal"
	} else if strings.Contains(mysqlDataType, "bit") {
		return "binary"
	}
	return "time"
}

func isNotNullFlagForMysql(value string) bool {
	if value == "YES" {
		return false
	}
	return true // else if "NO"
}

func scanArgs(types []*sql.ColumnType) []interface{} {
	dataPtrs := make([]interface{}, len(types))
	for i := range types {
		if types[i].DatabaseTypeName() == "INT" {
			dataPtrs[i] = new(int)
		} else if isStringTypeName(types[i].DatabaseTypeName()) {
			if _, ok := types[i].Nullable(); ok {
				dataPtrs[i] = new(sql.NullString)
			} else {
				dataPtrs[i] = new(string)
			}
		}
	}
	return dataPtrs
}

func isStringTypeName(typeName string) bool {
	switch typeName {
		case "CHAR", "VARCHAR", "TEXT": return true
		default: return false
	}
}

