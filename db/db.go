package db

import (
	"reflect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/runner"
	. "github.com/ShoichiroKitano/micro_test/db/infra"
)

var connectionInformations = [](*ConnectionInformation){}

type ConnectionInformation struct {
	name string
	rdbms string
	information string
}

func DefineConnection(connectionName, rdbms, information string) {
	connection := new(ConnectionInformation)
	connection.rdbms = rdbms
	connection.name = connectionName
	connection.information = information
	connectionInformations = append(connectionInformations, connection)
}

func findConnectionInformation(connectionName string) *ConnectionInformation {
	for _, connectionInformation := range connectionInformations {
		if connectionInformation.name == connectionName {
			return connectionInformation
		}
	}
	return nil
}

var defaultValues = make([]TableInformation, 0)

func (this DSL) DefineDefaultValue(defaultValue TableInformation) {
	defaultValues = append(defaultValues, defaultValue)
}

func findDefaultValueOf(tableName string) TableInformation {
	for _, tableInformation := range defaultValues {
		if tableInformation.tableName == tableName {
			return tableInformation
		}
	}
	return TableInformation{}
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
	defaultValue := findDefaultValueOf(fixtureTable.Name)
	if !reflect.DeepEqual(defaultValue, TableInformation{}) {
		completedTable := fixtureTable.FilledTableWith(defaultValue.DefaultRow())
		this.connection.StoreTable(completedTable)
	} else {
		this.connection.StoreTable(fixtureTable)
	}
}

func (this DSL) ShouldHaveTable(expected TableInformation) {
	expectedTable := expected.ToTable()
	resultTable := this.connection.FindTable(expectedTable.Name)
  if expectedTable.IsSameAsTable(resultTable) {
		runner.TestRunner.Result = "success"
	} else {
		runner.TestRunner.Result = ""
	}
}

func (this DSL) Truncate(tableName string) {
	this.connection.TruncateTable(tableName)
}
