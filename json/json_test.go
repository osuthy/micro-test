package json

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test各要素をJSONに変換する(t *testing.T) {
	t.Run("文字列の場合", func(t *testing.T) {
		var obj Object = A{"aho"}
		actual := obj.ToJson()
		assert.Equal(t, `["aho"]`, string(actual))
	})

	t.Run("文字列の配列の場合", func(t *testing.T) {
		a := A{"dog", "bird", "cat"}
		actual := a.ToJson()
		assert.Equal(t, `["dog","bird","cat"]`, string(actual))
	})

	t.Run("整数の場合", func(t *testing.T) {
		a := A{100}
		actual := a.ToJson()
		assert.Equal(t, `[100]`, string(actual))
	})

	t.Run("整数の配列の場合", func(t *testing.T) {
		a := A{1, 2, 3}
		actual := a.ToJson()
		assert.Equal(t, `[1,2,3]`, string(actual))
	})

	t.Run("入れ子構造のオブジェクトの場合", func(t *testing.T) {
		o := O{"string": "str", "object": O{"nestedObject": O{"int": 100}}}
		actual := o.ToJson()
		assert.Equal(t, `{"object":{"nestedObject":{"int":100}},"string":"str"}`, string(actual))
	})

	t.Run("配列を持つ入れ子構造のオブジェクトの場合", func(t *testing.T) {
		o := O{"string": "str", "intArrays": A{O{"one": 1}, O{"two": 2}}}
		actual := o.ToJson()
		assert.Equal(t, `{"intArrays":[{"one":1},{"two":2}],"string":"str"}`, string(actual))
	})

}

func TestJsonの要素を上書きする(t *testing.T) {
	t.Run("要素がオブジェクトの場合", func(t *testing.T) {
		o := O{"key1": "value1", "key2": "value2"}
		actual := o.Override("key1", "value100")
		assert.Equal(t, O{"key1": "value100", "key2": "value2"}, actual)
	})
}

func TestJsonの各要素をインクリメントして返却する(t *testing.T) {
	t.Run("int型の場合", func(t *testing.T) {
		o := O{"o": 10}
		actual := o.Generate(2)
		assert.Equal(t, A{O{"o": 11}, O{"o": 12}}, actual)
	})

	t.Run("float64型の場合", func(t *testing.T) {
		o := O{"o": 10.1}
		actual := o.Generate(2)
		assert.Equal(t, A{O{"o": 11.1}, O{"o": 12.1}}, actual)
	})

	t.Run("T型の場合", func(t *testing.T) {
		o := O{"o": T("2001-12-20 23:59:59")}
		actual := o.Generate(2)
		assert.Equal(t, A{O{"o": T("2001-12-20 23:59:59")}, O{"o": T("2001-12-20 23:59:59")}}, actual)
	})
}
