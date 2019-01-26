package db

import (
	"sort"
)

type R map[string]interface{}

func toTable(records []R) Table {
	rows := make([]Row, 0)
	for _, record := range records {
		columns := make([]Column, 0)
		for column, value := range record {
			columns = append(columns, Column{column, value})
		}
		sort.Slice(columns, func(i, j int) bool { return columns[i].name < columns[j].name })
		rows = append(rows, Row{columns})
	}
	return Table{rows}
}

func toColumn() {
}
