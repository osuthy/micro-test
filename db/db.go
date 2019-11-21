package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/osuthy/micro-test/runner"
	. "github.com/osuthy/micro-test/db/infra"
	. "github.com/osuthy/micro-test"
)

var connectionInformations = [](*ConnectionInformation){}

type ConnectionInformation struct {
	name        string
	rdbms       string
	information string
}

type RDBDef struct {
	config C
}

func (this *RDBDef) SetConnectionForLocal(tc TC) TC {
	driver := this.config["driver"].(string)
	localConfig := this.config["local"].(C)
	source := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		localConfig["user"].(string),
		localConfig["password"].(string),
		localConfig["host"].(string),
		localConfig["port"].(string),
		this.config["database"].(string),
	)
	c, _ := sql.Open(driver, source)
	tc[this.config["name"].(string)] = NewConnection(c, driver)
	return tc
}

func (this *RDBDef) SetConnectionForK8S(tc TC, namespace string) TC {
	return tc
}

func DefineConnection(config C) {
	localConfig := config["local"].(C)
	c := ConnectionInformation{
		name: config["name"].(string),
		rdbms: config["driver"].(string),
		information: fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			localConfig["user"].(string),
			localConfig["password"].(string),
			localConfig["host"].(string),
			localConfig["port"].(string),
			config["database"].(string)),
	}
	runner.AppendConnectionDefinable(&RDBDef{
		config: config,
	})
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

func DB(tc TC, connectionName string) DSL {
	//info := findConnectionInformation(connectionName)
	//con := FindDBConnection(info.rdbms, info.information)
	con := tc[connectionName].(*Connection)
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
		runner.Diffs.Push("assert is fail")
	}
}

func (this DSL) Truncate(tableName string) {
	this.connection.TruncateTable(Table(tableName).ToTable())
}

