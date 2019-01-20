package fixture

import()

DefineConnection("mysql", "root:@/db_name", "MYSQL_DB")
DefineConnection("postgresql", "root:@/db_name=schema", "POSTGRESQL_DB")

DB("db_name", "table_name").Default(
	R{"column1": "user", "column2", "id"}
)
