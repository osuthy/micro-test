package json

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test文字をJSONに変換する(t *testing.T) {
	var obj Object = &A{"aho"}
	expected := `["aho"]`
	actual := obj.ToJson()
	assert.Equal(t, expected, string(actual))
}

func Test整数をJSONに変換する(t *testing.T) {
	aho := A{100}
	actual := aho.ToJson()
	expected := `[100]`

	assert.Equal(t, expected, string(actual))
}

func Test整数の配列をJSONに変換する(t *testing.T) {
	aho := A{1, 2, 3}
	actual := aho.ToJson()
	expected := "[1,2,3]"

	assert.Equal(t, expected, string(actual))
}

func Test文字列の配列をJSONに変換する(t *testing.T) {
	aho := A{"dog", "bird", "cat"}
	actual := aho.ToJson()
	expected := `["dog","bird","cat"]`

	assert.Equal(t, expected, string(actual))
}

func Test入れ子構造のオブジェクトをJSONに変換する(t *testing.T) {
	o := O{"string": "str", "object": O{"nestedObject": O{"int": 100}}}
	actual := o.ToJson()
	expected := `{"object":{"nestedObject":{"int":100}},"string":"str"}`

	assert.Equal(t, expected, string(actual))
}

func Test配列を持つ入れ子構造のオブジェクトをJSONに変換する(t *testing.T) {
	o := O{"string": "str", "intArrays": A{O{"one": 1}, O{"two": 2}}}
	actual := o.ToJson()
	expected := `{"intArrays":[{"one":1},{"two":2}],"string":"str"}`

	assert.Equal(t, expected, string(actual))
}

func TestJsonの要素を上書きする(t *testing.T) {
	o := O{"key1": "value1", "key2": "value2"}
	actual := o.Override("key1", "value100")
	expected := O{"key1": "value100", "key2": "value2"}

	assert.Equal(t, expected, actual)
}

func Test要素がint型のjsonをインクリメントして返却する(t *testing.T) {
	o := O{"o": 10}
	actual := o.Generate(2)
	expected := A{O{"o":11}, O{"o":12}}

	assert.Equal(t, expected, actual)
}

func Test要素がfloat型のjsonをインクリメントして返却する(t *testing.T) {
	o := O{"o": 10.1}
	actual := o.Generate(2)
	expected := A{O{"o":11.1}, O{"o":12.1}}

	assert.Equal(t, expected, actual)
}

func Test要素がT型のjsonをインクリメントせずにそのまま返却する(t *testing.T) {
	o := O{"o": T("2001-12-20 23:59:59")}
	actual := o.Generate(2)
	expected := A{O{"o":T("2001-12-20 23:59:59")}, O{"o":T("2001-12-20 23:59:59")}}

	assert.Equal(t, expected, actual)
}