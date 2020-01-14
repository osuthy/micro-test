package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/osuthy/micro-test"
)

func Test定義したテストを実行できる(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	Describe("feature description", func() {
		It("test dscription1", func(c TC) {
			results = append(results, "test1")
		})

		It("test dscription2", func(c TC) {
			results = append(results, "test2")
		})
	})

	Run()

	assert.Equal(t, []string{"test1", "test2"}, results)
}

func TestImplicitSetUpをテスト毎に行える(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	Describe("feature description", func() {
		Before(func() {
			results = append(results, "setUp1")
		})

		Describe("sub feature description", func() {
			Before(func() {
				results = append(results, "setUp2")
			})

			It("test dscription1", func(c TC) {
				results = append(results, "description1")
			})

			It("test dscription2", func(c TC) {
				results = append(results, "description2")
			})
		})
	})

	Run()

	assert.Equal(t, []string{"setUp1", "setUp2", "description1", "setUp1", "setUp2", "description2"}, results)
}

func Testトップレベルの宣言ごとにテストSuiteが構築される(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	Describe("feature description1", func() {
		It("test dscription1", func(c TC) {
			results = append(results, "test1")
		})
	})

	Describe("feature description2", func() {
		It("test dscription2", func(c TC) {
			results = append(results, "test2")
		})
	})

	Run()

	assert.Equal(t, []string{"test1", "test2"}, results)
}

func TestImplicitTearDownをテスト毎に行える(t *testing.T) {
	defer resetSuites()
	resetSuites()

	results := []string{}
	Describe("feature description", func() {
		After(func() {
			results = append(results, "tearDown2")
		})

		Describe("sub feature description", func() {
			After(func() {
				results = append(results, "tearDown1")
			})

			It("test dscription1", func(c TC) {
				results = append(results, "description1")
			})

			It("test dscription2", func(c TC) {
				results = append(results, "description2")
			})
		})
	})

	Run()

	assert.Equal(t, []string{"description1", "tearDown1", "tearDown2", "description2", "tearDown1", "tearDown2"}, results)
}
