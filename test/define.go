package test

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/db"
)

func init() {
	DefineConnection2(C{
		"name": "conName",
		"driver": "mysql",
		"database": "test_micro_test",
		"local": C{
			"host": "localhost",
			"port": "3306",
			"user": "root",
			"password": "",
		},
		"k8s": C{ //k8s用のrepositoryに格納する
			"clusterIp": MinikubeIp(), // 動的にminikubeとその他のクラスタを変更できるようにする
			"user": "root",
			"password": "",
		},
	})
}
