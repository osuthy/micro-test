package db

func SpecDB() {
	// insert into connection_name.table_name (colum1, colum2) values
	// (1, "A"),
	// (2, "B");
	DB("connection_name", "table_name").HasRecords(
		R{"colum1": 1, "colum2": "A"},
		R{"colum1": 2, "colum2": "B"}
  )

	Exercise()

	// verify only soecified columns and ignore not specified columns and order
	DB("connection_name", "table_name").ShouldHaveColumns(
		R{"colum2": "B"}
		R{"colum2": "A"},
  )

	// verify all columns with order
	DB("connection_name", "table_name").ShouldHaveTable(
		R{"colum1": 1, "colum2": "A"},
		R{"colum1": 2, "colum2": "B"}
  )
}
