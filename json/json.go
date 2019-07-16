package json

import (
	"encoding/json"
	"fmt"
	// "reflect"
	"bytes"
)

type Object interface {
	ToJson() []byte
}

type O map[string]interface{}

func (this O) Print() {
	println("O")
}

func (this O) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}

func (this O) OverrideByStrings(elements ...string) O {
	o := make(O)
	for k, v := range this {
		o[k] = v
	}

	first := elements[0]

	if len(elements) == 2 {
		o[first] = elements[1]
		return o
	}

	x := o[first].(O).OverrideByStrings(elements[1:]...)
	o[first] = x

	return o

}

func (this O) overrideByObject(other O) O {

	for k, v := range other {
		if json, ok := v.(O); ok {
			newO := this[k].(O).overrideByObject(json)
			this[k] = newO
		} else {
			this[k] = v
		}
	}

	return this
}

func (this O) Override(elements ...interface{}) O {
	if _, ok := elements[0].(string); ok {
		r := []string{}
		for _, e := range elements {
			r = append(r, e.(string))
		}
		return this.OverrideByStrings(r...)
	}

	return this.overrideByObject(elements[0].(O))
}

type A []interface{}

func (this A) Print() {
	println("A")
}

func (this A) ToJson() []byte {
	json, _ := json.Marshal(this)
	return json
}

type T string

func (this O) Generate(num int) A {
	a := A{}

	for i := 1; i <= num; i++ {
		o := O{}

		for k, v := range this {
			switch value := v.(type) {
			case string:
				var buffer bytes.Buffer
				buffer.WriteString(value)
				buffer.WriteString(fmt.Sprint(i))
				o[k] = buffer.String()
			case int:
				o[k] = value + i
			case float64:
				o[k] = value + float64(i)
			case T:
				o[k] = value
			}
		}
		a = append(a, o)
	}
	return a
}
