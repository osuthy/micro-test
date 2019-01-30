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

func Test行の値で全てのカラムの値を補完する(t *testing.T) {
	result := Table{
		"name",
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
	}.filledTableWith(Row{[]Column{Column{"column3", "D3"}, Column{"column4", "D4"}}})
	assert.Equal(t, "name", result.name)
	assert.Equal(t, []Row{
			Row{
				[]Column{
					Column{"column1", "A1"}, Column{"column2", "A2"}, Column{"column3", "D3"}, Column{"column4", "D4"},
				},
			},
			Row{
				[]Column{
					Column{"column1", "B1"}, Column{"column2", "B2"}, Column{"column3", "D3"}, Column{"column4", "D4"},
				},
			}}, result.rows)
}

func Test補完対象の行にカラムの値がある場合は補完しない(t *testing.T) {
	result := Table{
		"name",
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
	}.filledTableWith(Row{[]Column{Column{"column1", "D1"}, Column{"column2", "D2"}, Column{"column3", "D3"},}})
	assert.Equal(t, "name", result.name)
	assert.Equal(t, []Row{
			Row{
				[]Column{
					Column{"column1", "A1"}, Column{"column2", "A2"}, Column{"column3", "D3"},
				},
			},
			Row{
				[]Column{
					Column{"column1", "B1"}, Column{"column2", "B2"}, Column{"column3", "D3"},
				},
			}}, result.rows)
}
