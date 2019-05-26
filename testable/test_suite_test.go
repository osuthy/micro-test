package testable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func createTestSuite(tests ...Testable) *TestSuite {
	return NewTestSuite2(tests, nil)
}

func Test一番最初に入力されたテストケースのみを実行(t *testing.T) {
	t.Run("テストケースの場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }))
		suite.Execute()
		assert.Equal(t, []string{"test1"}, result)
	})

	t.Run("テストスイートの場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(NewTestCase(func() { result = append(result, "test1") })),
			NewTestCase(func() { result = append(result, "test2") }))
		suite.Execute()
		assert.Equal(t, []string{"test1"}, result)
	})
}

func Test次のテストケースの取り出し(t *testing.T) {
	t.Run("テストケースがある場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }),
			NewTestCase(func() { result = append(result, "test3") }))

		next := suite.NextTest()

		next.Execute()
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストケースの個数分取り出す場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }),
			NewTestCase(func() { result = append(result, "test3") }))

		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストスイートがある場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase(func() { result = append(result, "test1") }),
				NewTestCase(func() { result = append(result, "test2") }),
				NewTestCase(func() { result = append(result, "test3") })))

		next := suite.NextTest()

		next.Execute()
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストスイートからテストケースの個数分取り出す場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase(func() { result = append(result, "test1") }),
				NewTestCase(func() { result = append(result, "test2") }),
				NewTestCase(func() { result = append(result, "test3") })))
		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストケースの1個の場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		suite.Add(NewTestCase(func() { result = append(result, "test1") }))
		assert.Equal(t, nil, suite.NextTest())
	})
}

func Test次のテストケースがあるか判定(t *testing.T) {
	t.Run("複数テストケースがある場合", func(t *testing.T) {
		suite := createTestSuite(
			NewTestCase(func() {}),
			NewTestCase(func() {}))
		assert.True(t, suite.HasNextTest())
	})

	t.Run("テストケースが1つの場合", func(t *testing.T) {
		suite := createTestSuite(NewTestCase(func() {}))
		assert.False(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストケースが1つの場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(NewTestCase(func() {})))
		assert.False(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストケースが複数の場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(
				NewTestCase(func() {}),
				NewTestCase(func() {})))
		assert.True(t, suite.HasNextTest())
	})
}
