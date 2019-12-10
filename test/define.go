package test

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db"
	. "github.com/osuthy/micro-test/http"
)

/*
_ = DefineK8sClustor(C{
	""
})
*/
// _ = DefineMinikubeClustorIp()
func init() {
	DefineHttpServer(C{
		"name": "test_server",
		"local": C{
			"host": "localhost",
			"port": 8080,
		},
		"k8s": C{
			"pod": "",
		},
	})
	DefineConnection(C{
		"name":     "conName",
		"driver":   "mysql",
		"database": "test_micro_test",
		"local": C{
			"host":     "localhost",
			"port":     "3306",
			"user":     "root",
			"password": "",
		},
		"k8s": C{
			"pod":      "",
			"user":     "root",
			"password": "",
		},
	})
}
