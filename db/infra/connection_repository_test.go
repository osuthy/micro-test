package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"github.com/stretchr/testify/assert"
)

func Testコネクションへのインターフェースの生成は一回のみ行う(t *testing.T) {
	con1 := FindDBConnection("mysql", "root:@/test_connection")
	con2 := FindDBConnection("mysql", "root:@/test_connection")
	assert.True(t, con1 == con2)
	assert.True(t, con1.Driver == con2.Driver)
}
