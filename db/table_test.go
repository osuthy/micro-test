package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//順不同、複数カラム、複数行
func TestRow2(t *testing.T) {
	result := Table{ []Row{
			Row{
				[]Column{
					Column{"column", 1},
				},
			},
			Row{
				[]Column{
					Column{"column", 2},
				},
			},
		},
	}.isSame(Table{[]Row{
		Row{[]Column{Column{"column", 2}}},
		Row{[]Column{Column{"column", 1}}},
	}})
	assert.False(t, result)
}

func TestRow(t *testing.T) {
	result := Table{ []Row{
			Row{
				[]Column{
					Column{"column", 1},
				},
			},
			Row{
				[]Column{
					Column{"column", 2},
				},
			},
		},
	}.isSame(Table{[]Row{
		Row{[]Column{Column{"column", 1}}},
		Row{[]Column{Column{"column", 2}}},
	}})
	assert.True(t, result)
}

func TestIgnoreOrder(t *testing.T) {
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

func TestIsSameTrue(t *testing.T) {
	result := Table{ []Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{[]Row{Row{[]Column{Column{"column1", 1}, Column{"column2", 2}}}}})
	assert.True(t, result)
}

func TestIsSameFalse(t *testing.T) {
	result := Table{ []Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{[]Row{Row{[]Column{Column{"column1", 100}, Column{"column2", 2}}}}})
	assert.False(t, result)
}
