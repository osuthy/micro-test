package infra

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testコネクションへのインターフェースの生成は一回のみ行う(t *testing.T) {
	con1 := FindDBConnection("mysql", "root:@/test_connection_repository")
	con2 := FindDBConnection("mysql", "root:@/test_connection_repository")
	assert.True(t, con1 == con2)
	assert.True(t, con1.Driver == con2.Driver)
}
