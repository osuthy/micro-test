package db

import (
	. "github.com/ShoichiroKitano/micro_test/db/infra"
	"github.com/ShoichiroKitano/micro_test/runner"
	_ "github.com/go-sql-driver/mysql"
)

var connectionInformations = [](*ConnectionInformation){}

type ConnectionInformation struct {
	name        string
	rdbms       string
	information string
}

func DefineConnection(connectionName, rdbms, information string) {
	c := ConnectionInformation{
		name:        connectionName,
		rdbms:       rdbms,
		information: information}
	connectionInformations = append(connectionInformations, &c)
}

func findConnectionInformation(connectionName string) *ConnectionInformation {
	for _, connectionInformation := range connectionInformations {
		if connectionInformation.name == connectionName {
			return connectionInformation
		}
	}
	return nil
}

type DSL struct {
	connection *Connection
}

func DB(connectionName string) DSL {
	info := findConnectionInformation(connectionName)
	con := FindDBConnection(info.rdbms, info.information)
	return DSL{con}
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
		runner.Queue.Push("assert is fail")
	}
}

func (this DSL) Truncate(tableName string) {
	this.connection.TruncateTable(Table(tableName).ToTable())
}

