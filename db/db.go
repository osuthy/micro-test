package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/runner"
)

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

func (this DSL) HasRecords(fixture TableInformation) {
	//fixtureTable := fixture.toTable()
	//this.connection.StoreTable(fixtureTable)
	con, _ := sql.Open("mysql", "root:@/test_micro_test")
	con.Query("insert into test (column1, column2) values ('A1', 'A2');")
	con.Query("insert into test (column1, column2) values ('B1', 'B2');")
	con.Close()
}

func (this DSL) ShouldHaveTable(expected TableInformation) {
	expectedTable := expected.toTable()
	resultTable := FindTable(this.connection.driver, expectedTable.name)
  if expectedTable.isSame(resultTable) {
		runner.TestRunner.Result = "success"
	} else {
		runner.TestRunner.Result = ""
	}
}

// コネクションがきれていた場合は再接続を行う
func DB(connectionName string) DSL {
	return DSL{connections[0]}
}
