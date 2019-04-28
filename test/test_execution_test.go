package test

import(
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/dsl"
	"github.com/ShoichiroKitano/micro_test/runner"
)

func Test定義したテストを実行できる(t *testing.T) {
	result1 := 0
	result2 := 0
	dsl.Feature("feature description", func() {
		dsl.Feature("sub feature description", func() {
			dsl.Test("test dscription1", func(){
				result1++
			})

			dsl.Test("test dscription2", func(){
				result2++
			})
		})
	})

	runner.Run()

	assert.Equal(t, 1, result1)
	assert.Equal(t, 1, result2)
}
