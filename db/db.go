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
	definition := this.connection.FindColumnDefinition(fixtureTable)
	//defaultValue := findDefaultValueOf(fixtureTable.Name)
	//if !reflect.DeepEqual(defaultValue, TableInformation{}) {
	completedTable := definition.FillTableWithDefaultValue(fixtureTable)
	this.connection.StoreTable(completedTable)
	//} else {
	//	this.connection.StoreTable(fixtureTable)
	//}
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

// 以下は一旦不要
func (this DSL) DefineDefaultValue(defaultValue TableInformation) {
	defaultValues = append(defaultValues, defaultValue)
}

var defaultValues = []TableInformation{}

func ResetDefalultValue() {
	defaultValues = []TableInformation{}
}

