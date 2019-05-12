package testable

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test実行可能なテストがあるか判定(t *testing.T) {
	t.Run("テストケースが未実行の場合", func(t *testing.T) {
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){}))
		assert.True(t, suite.HasUnexecutedTest())
	})

	t.Run("1つだけテストケースが実行済みの場合", func(t *testing.T) {
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){}))
		suite.Add(NewTestCase(func(){}))
		suite.Execute()
		assert.True(t, suite.HasUnexecutedTest())
	})

	t.Run("1つだけ入れ子のテストケースが実行済みの場合", func(t *testing.T) {
		suite1 := NewTestSuite()
		suite2 := NewTestSuite()
		suite1.Add(suite2)
		suite2.Add(NewTestCase(func(){}))
		suite2.Add(NewTestCase(func(){}))
		suite1.Execute()
		assert.True(t, suite1.HasUnexecutedTest())
	})

	t.Run("テストケースが実行済みの場合", func(t *testing.T) {
		suite := NewTestSuite()
		suite.Add(NewTestCase(func(){}))
		suite.Execute()
		assert.False(t, suite.HasUnexecutedTest())
	})


	t.Run("入れ子のテストケースが実行済みの場合", func(t *testing.T) {
		suite1 := NewTestSuite()
		suite2 := NewTestSuite()
		suite1.Add(suite2)
		suite2.Add(NewTestCase(func(){}))
		suite2.Add(NewTestCase(func(){}))
		suite1.Execute()
		suite1.Execute()
		assert.False(t, suite1.HasUnexecutedTest())
	})
}

