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
			"name":     "micro-test-mysql",
			"driver":   "mysql",
			"database": "test_db",
			"k8s": C{
				"user":     "root",
				"password": "",
			},
		},
	}
	tc := TC{}
	definition.SetConnectionForK8S(tc, "test-namespace")
	con := tc["micro-test-mysql"].(*Connection)
	assert.Nil(t, con.Ping())
}
