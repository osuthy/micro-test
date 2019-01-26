package db

import (
	"testing"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestFinedTable(t *testing.T) {
	db, _ := sql.Open("mysql", "root:@/test_micro_test")
	db.Query("insert into test (column1, column2) values ('A1', 'A2');")
	db.Query("insert into test (column1, column2) values ('B1', 'B2');")
	defer db.Query("truncate table test;")
	defer db.Close()

	table := FindTable(db, "test")
	assert.Equal(t, table, Table{
		[]Row{
			Row{
				[]Column{
					Column{"column1", "A1"}, Column{"column2", "A2"},
				},
			},
			Row{
				[]Column{
					Column{"column1", "B1"}, Column{"column2", "B2"},
				},
			},
		},
	})
}
