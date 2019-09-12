package testable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Testテストケースの実行(t *testing.T) {
	t.Run("テストケースの場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }))
		suite.Execute()
		assert.Equal(t, []string{"test1"}, result)
	})

	t.Run("複数回実行する場合も常に同じテストケースが実行される", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }))
		suite.Execute()
		suite.Execute()
		assert.Equal(t, []string{"test1", "test1"}, result)
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
	t.Run("次のテストケースがある場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }),
			NewTestCase(func() { result = append(result, "test3") }))

		next := suite.NextTest()

		next.Execute()
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストケースを全て取り出す場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }),
			NewTestCase(func() { result = append(result, "test2") }),
			NewTestCase(func() { result = append(result, "test3") }))

		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストスイートをネストしている場合", func(t *testing.T) {
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

	t.Run("複数のテストスイートが存在する場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			createTestSuite(
				NewTestCase(func() { result = append(result, "test1") }),
				NewTestCase(func() { result = append(result, "test2") })),
			createTestSuite(
				NewTestCase(func() { result = append(result, "test3") })),
			)
		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("次のテストケースが無い場合", func(t *testing.T) {
		var result []string
		suite := createTestSuite(
			NewTestCase(func() { result = append(result, "test1") }))
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

	t.Run("入れ子のテストスイートが複数の場合", func(t *testing.T) {
		suite := createTestSuite(
			createTestSuite(NewTestCase(func() {})),
			createTestSuite(NewTestCase(func() {})))
		assert.True(t, suite.HasNextTest())
	})
}

//func Test(t *testing.T) {
//	t.Run("テストケースの場合", func(t *testing.T) {
//	})
//}

func createTestSuite(tests ...Testable) *TestSuite {
	return NewTestSuite(tests, nil)
}

