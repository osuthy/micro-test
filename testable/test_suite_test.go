package testable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testテストケースの実行(t *testing.T) {
	t.Run("ネストされていないテストケースの場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test2") }))
		suite.Execute(TestContext{})
		assert.Equal(t, []string{"test1"}, result)
	})

	t.Run("複数回実行する場合も常に同じテストケースが実行される", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test2") }))
		suite.Execute(TestContext{})
		suite.Execute(TestContext{})
		assert.Equal(t, []string{"test1", "test1"}, result)
	})

	t.Run("他のテストスイートにネストされたテストケースの場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(NewTestCase("", func(c TestContext) { result = append(result, "test1") })),
			NewTestCase("", func(c TestContext) { result = append(result, "test2") }))
		suite.Execute(TestContext{})
		assert.Equal(t, []string{"test1"}, result)
	})
}

func Test実行対象のテストケースの説明を取得(t *testing.T) {
	t.Run("ネストされていないテストケースの場合", func(t *testing.T) {
		result := createTestSuiteWithDescription("A", NewTestCase("B", func(c TestContext) {})).Descriptions()
		assert.Equal(t, []string{"A", "B"}, result)
	})

	t.Run("実行対象を次のテストケースに変更した場合", func(t *testing.T) {
		result := createTestSuiteWithDescription("A",
			NewTestCase("B1", func(c TestContext) {}),
			NewTestCase("B2", func(c TestContext) {}),
		).NextTest().Descriptions()
		assert.Equal(t, []string{"A", "B2"}, result)
	})

	t.Run("ネストされたテストケースの場合", func(t *testing.T) {
		result := createTestSuiteWithDescription("A",
			createTestSuiteWithDescription("B",
				NewTestCase("C", func(c TestContext) {}),
			)).Descriptions()
		assert.Equal(t, []string{"A", "B", "C"}, result)
	})

	t.Run("実行対象を次のテストスイートに変更した場合", func(t *testing.T) {
		result := createTestSuiteWithDescription("A",
			createTestSuiteWithDescription("B1",
				NewTestCase("C1", func(c TestContext) {})),
			createTestSuiteWithDescription("B2",
				NewTestCase("C2", func(c TestContext) {}),
				NewTestCase("C3", func(c TestContext) {}),
			)).NextTest().Descriptions()
		assert.Equal(t, []string{"A", "B2", "C2"}, result)
	})
}

func Test次のテストケースの取り出し(t *testing.T) {
	t.Run("次のテストケースがある場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test2") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test3") }))

		next := suite.NextTest()

		next.Execute(TestContext{})
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストケースを全て取り出す場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test2") }),
			NewTestCase("", func(c TestContext) { result = append(result, "test3") }))

		last := suite.NextTest().NextTest()

		last.Execute(TestContext{})
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストスイートをネストしている場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
				NewTestCase("", func(c TestContext) { result = append(result, "test2") }),
				NewTestCase("", func(c TestContext) { result = append(result, "test3") })))

		next := suite.NextTest()

		next.Execute(TestContext{})
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストスイートから最後のテストケースを取り出す場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
				NewTestCase("", func(c TestContext) { result = append(result, "test2") }),
				NewTestCase("", func(c TestContext) { result = append(result, "test3") })))
		last := suite.NextTest().NextTest()

		last.Execute(TestContext{})
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("複数のテストスイートが存在する場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase("", func(c TestContext) { result = append(result, "test1") }),
				NewTestCase("", func(c TestContext) { result = append(result, "test2") })),
			createTestSuite(
				NewTestCase("", func(c TestContext) { result = append(result, "test3") })),
			)
		last := suite.NextTest().NextTest()

		last.Execute(TestContext{})
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("次のテストケースが無い場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) { result = append(result, "test1") }))
		assert.Equal(t, nil, suite.NextTest())
	})
}

func Test次のテストケースがあるか判定(t *testing.T) {
	t.Run("複数テストケースがある場合", func(t *testing.T) {
		suite := createTestSuite(
			NewTestCase("", func(c TestContext) {}),
			NewTestCase("", func(c TestContext) {}))
		assert.True(t, suite.HasNextTest())
	})

	t.Run("テストケースが1つの場合", func(t *testing.T) {
		suite := createTestSuite(NewTestCase("", func(c TestContext) {}))
		assert.False(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストケースが1つの場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(NewTestCase("", func(c TestContext) {})))
		assert.False(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストケースが複数の場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(
				NewTestCase("", func(c TestContext) {}),
				NewTestCase("", func(c TestContext) {})))
		assert.True(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストスイートが複数の場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(NewTestCase("", func(c TestContext) {})),
			createTestSuite(NewTestCase("", func(c TestContext) {})))
		assert.True(t, suite.HasNextTest())
	})
}

func createTestSuite(tests ...Testable) *TestSuite {
	return createTestSuiteWithDescription("", tests...)
}

func createTestSuiteWithDescription(description string, tests ...Testable) *TestSuite {
	return NewTestSuite(description, tests, nil, nil)
}
