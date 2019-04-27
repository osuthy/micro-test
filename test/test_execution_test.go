package test

import(
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/dsl"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func Test定義したテストを実行できる(t *testing.T) {
	result := ""
	dsl.Test("test dscription", func(){
		result = "test ran"
	})
	runner.Run()
	assert.Equal(t, "test ran", result)
}
