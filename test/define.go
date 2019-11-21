package test

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db"
)

/*
_ = DefineK8sClustor(C{
	""
})
*/
// _ = DefineMinikubeClustorIp()
func init() {
	DefineConnection(C{
		"name": "conName",
		"driver": "mysql",
		"database": "test_micro_test",
		"local": C{
			"host": "localhost",
			"port": "3306",
			"user": "root",
			"password": "",
		},
		"k8s": C{
			"pod": "",
			"user": "root",
			"password": "",
		},
	})
	// DefineApiServer()
}
/*
connection.MakeContextForLocal(tc TestContext) tc {
	tc["conName"] = Conn.New(sql.Open(, ))
	tc["assertionErrors"] = 
	return tc 
}
*/
