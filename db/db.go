package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ShoichiroKitano/micro_test/runner"
)

var (
	connections = make([](*Connection), 0)
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
	tableName string
}

func (this DSL) ShouldHaveTable(expected TableInformation) {
	expectedTable := expected.toTable()
	resultTable := FindTable(this.connection.driver, this.tableName)
  if expectedTable.isSame(resultTable) {
		runner.TestRunner.Result = "success"
	} else {
		runner.TestRunner.Result = ""
	}
}

func DB(connectionName, tableName string) DSL {
	return DSL{connections[0], tableName}
}
