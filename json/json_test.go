package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test各要素をJSONに変換する(t *testing.T) {
	t.Run("文字列の場合", func(t *testing.T) {
		actual := A{"aho"}.ToJson()
		assert.Equal(t, `["aho"]`, string(actual))
	})

	t.Run("文字列の配列の場合", func(t *testing.T) {
		actual := A{"dog", "bird", "cat"}.ToJson()
		assert.Equal(t, `["dog","bird","cat"]`, string(actual))
	})

	t.Run("整数の場合", func(t *testing.T) {
		actual := A{100}.ToJson()
		assert.Equal(t, `[100]`, string(actual))
	})

	t.Run("整数の配列の場合", func(t *testing.T) {
		actual := A{1, 2, 3}.ToJson()
		assert.Equal(t, `[1,2,3]`, string(actual))
	})

	t.Run("入れ子構造のオブジェクトの場合", func(t *testing.T) {
		actual := O{"string": "str", "object": O{"nestedObject": O{"int": 100}}}.ToJson()
		assert.Equal(t, `{"object":{"nestedObject":{"int":100}},"string":"str"}`, string(actual))
	})

	t.Run("配列を持つ入れ子構造のオブジェクトの場合", func(t *testing.T) {
		actual := O{"string": "str", "intArrays": A{O{"one": 1}, O{"two": 2}}}.ToJson()
		assert.Equal(t, `{"intArrays":[{"one":1},{"two":2}],"string":"str"}`, string(actual))
	})

}

func TestJsonの要素を上書きする(t *testing.T) {
	t.Run("要素がオブジェクトの場合", func(t *testing.T) {
		actual := O{"key1": "value1", "key2": "value2"}.Override("key1", "value100")
		assert.Equal(t, O{"key1": "value100", "key2": "value2"}, actual)
	})
}

func TestJsonの各要素をインクリメントして返却する(t *testing.T) {
	t.Run("int型の場合", func(t *testing.T) {
		actual := O{"o": 10}.Generate(2)
		assert.Equal(t, A{O{"o": 11}, O{"o": 12}}, actual)
	})

	t.Run("float64型の場合", func(t *testing.T) {
		actual := O{"o": 10.1}.Generate(2)
		assert.Equal(t, A{O{"o": 11.1}, O{"o": 12.1}}, actual)
	})

	t.Run("T型の場合", func(t *testing.T) {
		actual := O{"o": T("2001-12-20 23:59:59")}.Generate(2)
		assert.Equal(t, A{O{"o": T("2001-12-20 23:59:59")}, O{"o": T("2001-12-20 23:59:59")}}, actual)
	})
}
