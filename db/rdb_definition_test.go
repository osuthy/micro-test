package db

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db/infra"
)

func Testネームスペースからコネクションを作成(t *testing.T) {
	definition := &RDBDefinition{
		config: C{
			"name":     "conName",
			"driver":   "mysql",
			"database": "test_db",
			"k8s": C{
				"svc":     "micro-test-mysql",
				"user":     "root",
				"password": "",
			},
		},
	}
	tc := TC{}
	definition.SetConnectionForK8S(tc, "test-namespace")
	con := tc["conName"].(*Connection)
	assert.Nil(t, con.Ping())
}
