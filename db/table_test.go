package db

import (
	"testing"
)

func TestIsSame(t *testing.T) {
	result := Table{
		[]Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{[]Row{Row{[]Column{Column{"column2", 2}, Column{"column1", 1}}}}})
	if !result { t.Fail() }
}

