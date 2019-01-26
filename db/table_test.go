package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestIsSame(t *testing.T) {
	result := Table{ []Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{[]Row{Row{[]Column{Column{"column2", 2}, Column{"column1", 1}}}}})
	assert.True(t, result)
}

