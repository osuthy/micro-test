package db

import (
	"reflect"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/runner"
)

var defaultValues = make([]TableInformation, 0)

// 本当はkubernetesのnamespace単位でコネクション作るからここでOpenはしない予定
func DefineConnection(connectionName, rdbms, information string) {
	db, _ := sql.Open(rdbms, information)
	con := new(Connection)
	con.name = connectionName
	con.driver = db
	connections = append(connections, con)
}

type DSL struct {
	connection *Connection
}

// コネクションがきれていた場合は再接続を行う
func DB(connectionName string) DSL {
	return DSL{connections[0]}
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
  if expectedTable.IsSame(resultTable) {
		runner.TestRunner.Result = "success"
	} else {
		runner.TestRunner.Result = ""
	}
}

func (this DSL) SetDefaultValue(defaultValue TableInformation) {
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

