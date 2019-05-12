package test

import(
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test/dsl"
	"github.com/ShoichiroKitano/micro_test/runner"
	. "github.com/ShoichiroKitano/micro_test/dsl/testable"
)

func resetSuites() {
	dsl.Suites = []Testable{}
}

func Test定義したテストを実行できる(t *testing.T) {
	defer resetSuites()
	resetSuites()

	result1 := 0
	result2 := 0
	dsl.Feature("feature description", func() {
		dsl.Test("test dscription1", func(){
			result1++
		})

		dsl.Test("test dscription2", func(){
			result2++
		})
	})

	runner.Run()

	assert.Equal(t, 1, result1)
	assert.Equal(t, 1, result2)
}

func TestImplicitSetUpをテスト毎に行える(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	dsl.Feature("feature description", func() {
		dsl.Before(func() {
			results = append(results, "setUp1")
		})

		dsl.Feature("sub feature description", func() {
			dsl.Before(func() {
				results = append(results, "setUp2")
			})

			dsl.Test("test dscription1", func(){
				results = append(results, "description1")
			})

			dsl.Test("test dscription2", func(){
				results = append(results, "description2")
			})
		})
	})

	runner.Run()

	assert.Equal(t, []string{"setUp1", "setUp2", "description1", "setUp1", "setUp2", "description2"}, results)
}

func Testトップレベルの宣言ごとにテストSuiteが構築される(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	dsl.Feature("feature description1", func() {
		dsl.Test("test dscription1", func(){
			results = append(results, "description1")
		})
	})

	dsl.Feature("feature description2", func() {
		dsl.Test("test dscription2", func(){
			results = append(results, "description2")
		})
	})

	runner.Run()

	assert.Equal(t, []string{"description1", "description2"}, results)
}
