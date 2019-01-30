package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

//テーブル名が異なる
func TestRow2(t *testing.T) {
	result := Table{
		"name",
		[]Row{
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
	}.isSame(Table{
		"name",
		[]Row{
			Row{[]Column{Column{"column", 2}}},
			Row{[]Column{Column{"column", 1}}},
	}})
	assert.False(t, result)
}

func TestRow(t *testing.T) {
	result := Table{
		"name",
		[]Row{
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
	}.isSame(Table{
		"name",
		[]Row{
		Row{[]Column{Column{"column", 1}}},
		Row{[]Column{Column{"column", 2}}},
	}})
	assert.True(t, result)
}

func TestIgnoreOrder(t *testing.T) {
	result := Table{
		"name",
		[]Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{"name", []Row{Row{[]Column{Column{"column2", 2}, Column{"column1", 1}}}}})
	assert.True(t, result)
}

func TestIsSameTrue(t *testing.T) {
	result := Table{
		"name",
		[]Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{"name", []Row{Row{[]Column{Column{"column1", 1}, Column{"column2", 2}}}}})
	assert.True(t, result)
}

func TestIsSameFalse(t *testing.T) {
	result := Table{
		"name",
		[]Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.isSame(Table{"name", []Row{Row{[]Column{Column{"column1", 100}, Column{"column2", 2}}}}})
	assert.False(t, result)
}

func Test行方法で全ての行の値を補完する(t *testing.T) {
	result := Table{
		"name",
		[]Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2},
				},
			},
		},
	}.filledTableWith(Row{[]Column{Column{"column3", 3}, Column{"column4", 4}}})
	assert.Equal(t, result.rows, []Row{
			Row{
				[]Column{
					Column{"column1", 1}, Column{"column2", 2}, Column{"column3", 3}, Column{"column4", 4},
				}}})
}

