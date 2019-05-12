package test

import(
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ShoichiroKitano/micro_test"
	"github.com/ShoichiroKitano/micro_test/runner"
	. "github.com/ShoichiroKitano/micro_test/testable"
)

func resetSuites() {
	micro_test.Suites = []Testable{}
}

func Test定義したテストを実行できる(t *testing.T) {
	defer resetSuites()
	resetSuites()

	result1 := 0
	result2 := 0
	micro_test.Feature("feature description", func() {
		micro_test.Test("test dscription1", func(){
			result1++
		})

		micro_test.Test("test dscription2", func(){
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
	micro_test.Feature("feature description", func() {
		micro_test.Before(func() {
			results = append(results, "setUp1")
		})

		micro_test.Feature("sub feature description", func() {
			micro_test.Before(func() {
				results = append(results, "setUp2")
			})

			micro_test.Test("test dscription1", func(){
				results = append(results, "description1")
			})

			micro_test.Test("test dscription2", func(){
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
	micro_test.Feature("feature description1", func() {
		micro_test.Test("test dscription1", func(){
			results = append(results, "description1")
		})
	})

	micro_test.Feature("feature description2", func() {
		micro_test.Test("test dscription2", func(){
			results = append(results, "description2")
		})
	})

	runner.Run()

	assert.Equal(t, []string{"description1", "description2"}, results)
}
