package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ConnectionInfo struct {
	userName	string
	password	string
	protocol	string
	host		string
	port		int
}

func (this ConnectionInfo) builder(userName string, password string, protocol string, host string, port int) ConnectionInfo{
	return ConnectionInfo{userName, password, protocol, host, port}
}

func conncetion_builder(driverName string, connectionInfo ConnectionInfo) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@%s(%s:%d)/",
		connectionInfo.userName,
		connectionInfo.password,
		connectionInfo.protocol,
		connectionInfo.host,
		connectionInfo.port)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		fmt.Println("database connetion error!!")
		fmt.Println(err)
	}
	return db
}

func main() {
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	connectionInfo := ConnectionInfo{}.builder("root", "micro-test", "tcp", "127.0.0.1", 3306)
	db := conncetion_builder("mysql", connectionInfo)

	err := db.Ping()
    if err != nil {
        panic(err)
	}

	fmt.Println("Hello World.")
}
