package testable

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test一番最初に入力されたテストケースのみを実行(t *testing.T) {
	t.Run("テストケースの場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){ result = append(result, "test1") }))
		suite.Add(NewTestCase(func(){ result = append(result, "test2") }))
		suite.Execute()
		assert.Equal(t, []string{"test1"}, result)
	})

	t.Run("テストスイートの場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		nested := NewTestSuite()
		nested.Add(NewTestCase(func(){ result = append(result, "test1") }))
		suite.Add(nested)
		suite.Add(NewTestCase(func(){ result = append(result, "test2") }))
		suite.Execute()
		assert.Equal(t, []string{"test1"}, result)
	})
}

func Test次のテストケースの取り出し(t *testing.T) {
	t.Run("テストケースがある場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){ result = append(result, "test1") }))
		suite.Add(NewTestCase(func(){ result = append(result, "test2") }))
		suite.Add(NewTestCase(func(){ result = append(result, "test3") }))
		next := suite.NextTest()

		next.Execute()
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストケースの個数分取り出す場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){ result = append(result, "test1") }))
		suite.Add(NewTestCase(func(){ result = append(result, "test2") }))
		suite.Add(NewTestCase(func(){ result = append(result, "test3") }))
		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストスイートがある場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		nested := NewTestSuite()
		nested.Add(NewTestCase(func(){ result = append(result, "test1") }))
		nested.Add(NewTestCase(func(){ result = append(result, "test2") }))
		nested.Add(NewTestCase(func(){ result = append(result, "test3") }))
		suite.Add(nested)
		next := suite.NextTest()

		next.Execute()
		assert.Equal(t, []string{"test2"}, result)
	})

	t.Run("テストスイートからテストケースの個数分取り出す場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		nested := NewTestSuite()
		nested.Add(NewTestCase(func(){ result = append(result, "test1") }))
		nested.Add(NewTestCase(func(){ result = append(result, "test2") }))
		nested.Add(NewTestCase(func(){ result = append(result, "test3") }))
		suite.Add(nested)
		last := suite.NextTest().NextTest()

		last.Execute()
		assert.Equal(t, []string{"test3"}, result)
	})

	t.Run("テストケースの1個の場合", func(t *testing.T) {
		var result []string
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){ result = append(result, "test1") }))
		assert.Equal(t, nil, suite.NextTest())
	})
}

func Test次のテストケースがあるか判定(t *testing.T) {
	t.Run("複数テストケースがある場合", func(t *testing.T) {
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){}))
		suite.Add(NewTestCase(func(){}))
		assert.True(t, suite.HasNextTest())
	})

	t.Run("テストケースが1つの場合", func(t *testing.T) {
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){}))
		assert.False(t, suite.HasNextTest())
	})

	t.Run("入れ子のテストケースが1つの場合", func(t *testing.T) {
		suite1 := NewTestSuite()
		suite2 := NewTestSuite()
		suite1.Add(suite2)
		suite2.Add(NewTestCase(func(){}))
		assert.False(t, suite1.HasNextTest())
	})

	t.Run("入れ子のテストケースが複数の場合", func(t *testing.T) {
		suite1 := NewTestSuite()
		suite2 := NewTestSuite()
		suite1.Add(suite2)
		suite2.Add(NewTestCase(func(){}))
		suite2.Add(NewTestCase(func(){}))
		assert.True(t, suite1.HasNextTest())
	})
}

