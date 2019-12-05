package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/osuthy/micro-test/runner"
	. "github.com/osuthy/micro-test/db/infra"
	. "github.com/osuthy/micro-test"
)

type RDBDef struct {
	config C
}

func DefineConnection(config C) {
	runner.AppendConnectionDefinable(&RDBDef{
		config: config,
	})
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

type DSL struct {
	connection *Connection
	differences *runner.Differences
}

func DB(tc TC, connectionName string) DSL {
	con := tc[connectionName].(*Connection)
	diffs := tc["differences"].(*runner.Differences)
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

