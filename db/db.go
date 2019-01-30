package db

import (
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
	fixtureTable := fixture.toTable()
	completedTable := fixtureTable.filledTableWith(defaultValues[0].defaultRow())
	this.connection.StoreTable(completedTable)
}

func (this DSL) ShouldHaveTable(expected TableInformation) {
	expectedTable := expected.toTable()
	resultTable := this.connection.FindTable(expectedTable.name)
  if expectedTable.isSame(resultTable) {
		runner.TestRunner.Result = "success"
	} else {
		runner.TestRunner.Result = ""
	}
}

func (this DSL) SetDefaultValue(defaultValue TableInformation) {
	defaultValues = append(defaultValues, defaultValue)
}
