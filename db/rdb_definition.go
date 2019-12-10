package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db/infra"
)

type RDBDefinition struct {
	config C
}

func (this *RDBDefinition) SetConnectionForLocal(tc TC) TC {
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

func (this *RDBDefinition) SetConnectionForK8S(tc TC, namespace string) TC {
	return tc
}

func DefineRDB(config C) {
	AppendConnectionDefinable(&RDBDefinition{
		config: config,
	})
}
