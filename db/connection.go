package db

import (
	"database/sql"
)

type Connection struct {
	driver *sql.DB
	name string
}

