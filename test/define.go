package test

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db"
	. "github.com/osuthy/micro-test/http"
)

func init() {
	DefineHttpServer(C{
		"name": "test_server",
		"local": C{
			"host": "localhost",
			"port": 8080,
		},
		"k8s": C{
			"svn": "",
		},
	})
	DefineRDB(C{
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
			"svn": "",
			"user":     "root",
			"password": "",
		},
	})
}
