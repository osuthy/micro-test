package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	// root:micro-tes@tcp(127.0.0.1:3306)/sample
	db, err := sql.Open("mysql", "root:micro-test@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
    if err != nil {
        panic(err)
	}

	fmt.Println("Hello World.")
}
