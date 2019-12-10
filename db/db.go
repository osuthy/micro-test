package db

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db/infra"
)

type DSL struct {
	connection  *Connection
	differences *Differences
}

func DB(tc TC, connectionName string) DSL {
	con := tc[connectionName].(*Connection)
	diffs := tc["differences"].(*Differences)
	return DSL{connection: con, differences: diffs}
}

func (this DSL) HasRecords(fixture TableInformation) {
	fixtureTable := fixture.ToTable()
	definition := this.connection.FindColumnDefinition(fixtureTable)
	completedTable := definition.FillTableWithDefaultValue(fixtureTable)
	this.connection.StoreTable(completedTable)
}

func (this DSL) ShouldHaveTable(expected TableInformation) {
	expectedTable := expected.ToTable()
	resultTable := this.connection.FindTable(expectedTable)
	if !expectedTable.IsSameAsTable(resultTable) {
		this.differences.Push("assert is fail")
	}
}

func (this DSL) Truncate(tableName string) {
	this.connection.TruncateTable(Table(tableName).ToTable())
}
