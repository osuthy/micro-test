package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test複数のRecordをTableに変換する(t *testing.T) {
	records := []R{
		R{"column1": "A1", "column2": "A2"},
		R{"column1": "B1", "column2": "B2"},
	}
	assert.Equal(t, toTable(records), Table{
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
	})
}
