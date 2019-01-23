package db

import (
	"testing"
	"reflect"
)

func TestRecordToTable(t *testing.T) {
	records := []R{
		R{"column1": "A1", "column2": "A2"},
		R{"column1": "B1", "column2": "B2"},
	}
	if !reflect.DeepEqual(toTable(records), Table{
		[]Row {
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
	}) {
		t.Error("toTable is fail")
	}
}
